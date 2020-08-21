package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/apps/agenda"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/apps/assignment"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/apps/core"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/apps/mediafile"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/apps/motion"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/apps/poll"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/apps/topic"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/apps/user"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/auth"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/autoupdate"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/datastore"
	autoupdatehttp "github.com/OpenSlides/openslides3-autoupdate-service/internal/http"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/notify"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/redis"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/restricter"
)

func main() {
	workerAddr := getEnv("WORKER_PROTOCOL", "http") + "://" + getEnv("WORKER_HOST", "localhost") + ":" + getEnv("WORKER_PORT", "8000")
	redisAddr := getEnv("MESSAGE_BUS_HOST", "localhost") + ":" + getEnv("MESSAGE_BUS_PORT", "6379")

	redisConn := redis.New(redisAddr)
	for {
		if err := redisConn.TestConn(); err == nil {
			break
		}
		log.Printf("Can not connect to redis at %s. Retry...", redisAddr)
		time.Sleep(time.Second)
	}
	fmt.Printf("Connected to Redis at %s\n", redisAddr)

	requiredUserCallables := openslidesRequiredUsers()
	projectorCallables := openslidesProjectorCallables()
	closed := make(chan struct{})
	ds, err := datastore.New(workerAddr, redisConn, requiredUserCallables, projectorCallables, closed)
	if err != nil {
		log.Fatalf("Can not initialize data: %v", err)
	}
	fmt.Printf("Connected to OpenSlides at %s\n", workerAddr)

	osRestricters := openslidesRestricters(ds)
	restricter := restricter.New(ds, osRestricters)
	auth := auth.New(workerAddr)

	a, err := autoupdate.New(ds, restricter, closed)
	if err != nil {
		log.Fatalf("Can not create autoupdate service: %v", err)
	}

	n := notify.New(redisConn, closed)

	var forceHTTP bool
	if getEnv("FORCE_HTTP2", "") != "" {
		forceHTTP = true
	}

	handler := autoupdatehttp.New(auth, autoupdatehttp.WithForceHTTP2(forceHTTP), autoupdatehttp.WithAutoupdate(a), autoupdatehttp.WithNotify(n))

	// Create tls http2 server.
	listenAddr := getEnv("AUTOUPDATE_HOST", "") + ":" + getEnv("AUTOUPDATE_PORT", "8002")
	srv := &http.Server{Handler: handler}
	ln, err := tlsListener(listenAddr)
	if err != nil {
		log.Fatalf("Can not create tls listener: %v", err)
	}
	defer ln.Close()

	go func() {
		fmt.Printf("Listen on %s\n", listenAddr)
		if err := srv.Serve(ln); err != http.ErrServerClosed {
			log.Fatalf("HTTP Server failed: %v", err)
		}
	}()

	waitForShutdown()

	close(closed)
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("Error on HTTP server shutdown: %v", err)
	}
}

// waitForShutdown blocks until the service exists.
//
// It listens on SIGINT and SIGTERM. If the signal is received for a second
// time, the process is killed with statuscode 1.
func waitForShutdown() {
	sigint := make(chan os.Signal, 1)
	// syscall.SIGTERM is not pressent on all plattforms. Since the autoupdate
	// service is only run on linux, this is ok. If other plattforms should be
	// supported, os.Interrupt should be used instead.
	signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
	<-sigint
	go func() {
		<-sigint
		os.Exit(1)
	}()
}

// getEnv returns the value of the environment variable env. If it is empty, the
// defaultValue is used.
func getEnv(env, devaultValue string) string {
	value := os.Getenv(env)
	if value == "" {
		return devaultValue
	}
	return value
}

const (
	generalCertName = "cert.pem"
	generalKeyName  = "key.pem"
	specialCertName = "autoupdate.pem"
	specialKeyName  = "autoupdate-key.pem"
)

func getCert() (tls.Certificate, error) {
	certDir := getEnv("CERT_DIR", "")
	if certDir == "" {
		cert, err := autoupdatehttp.GenerateCert()
		if err != nil {
			return tls.Certificate{}, fmt.Errorf("creating new certificate: %w", err)
		}
		fmt.Println("Use inmemory self signed certificate")
		return cert, nil
	}
	certFile := path.Join(certDir, specialCertName)
	if _, err := os.Stat(certFile); os.IsNotExist(err) {
		certFile2 := path.Join(certDir, generalCertName)
		if _, err := os.Stat(certFile); os.IsNotExist(err) {
			return tls.Certificate{}, fmt.Errorf("%s or %s has to exist", certFile, certFile2)
		}
		certFile = certFile2
	}

	keyFile := path.Join(certDir, specialKeyName)
	if _, err := os.Stat(keyFile); os.IsNotExist(err) {
		keyFile2 := path.Join(certDir, generalKeyName)
		if _, err := os.Stat(keyFile); os.IsNotExist(err) {
			return tls.Certificate{}, fmt.Errorf("%s or %s has to exist", keyFile, keyFile2)
		}
		keyFile = keyFile2
	}

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("loading certificates from %s and %s: %w", certFile, keyFile, err)
	}
	fmt.Printf("Use certificate %s with key %s\n", certFile, keyFile)

	return cert, nil
}

func tlsListener(listenAddr string) (net.Listener, error) {
	cert, err := getCert()
	if err != nil {
		return nil, fmt.Errorf("get certificate: %w", err)
	}

	tlsConf := new(tls.Config)
	tlsConf.NextProtos = []string{"h2"}
	tlsConf.Certificates = []tls.Certificate{cert}

	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return nil, fmt.Errorf("create tcp listener: %w", err)
	}
	return tls.NewListener(ln, tlsConf), nil
}

func openslidesRequiredUsers() map[string]func(json.RawMessage) ([]int, string, error) {
	return map[string]func(json.RawMessage) ([]int, string, error){
		"agenda/list-of-speakers":       agenda.RequiredSpeakers,
		"assignments/assignment":        assignment.RequiredAssignments,
		"assignments/assignment-poll":   assignment.RequiredPoll,
		"assignments/assignment-option": assignment.RequiredPollOption,
		"motions/motion":                motion.RequiredMotions,
		"motions/motion-option":         motion.RequiredPollOption,
	}
}

func openslidesRestricters(ds restricter.HasPermer) map[string]restricter.Element {
	basePerm := restricter.BasePermission(ds)
	return map[string]restricter.Element{
		"agenda/item":             agenda.Restrict(ds),
		"agenda/list-of-speakers": basePerm(agenda.CanSeeListOfSpeakers),

		"assignments/assignment":        basePerm(assignment.CanSee),
		"assignments/assignment-poll":   poll.RestrictPoll(ds, assignment.CanSee, assignment.CanManage, []string{"amount_global_no", "amount_global_abstain"}),
		"assignments/assignment-option": poll.RestrictOption(ds, assignment.CanSee, assignment.CanManage),
		"assignments/assignment-vote":   poll.RestrictVote(ds, assignment.CanSee, assignment.CanManage),

		"core/projector":          basePerm(core.CanSeeProjector),
		"core/projection-default": basePerm(core.CanSeeProjector),
		"core/projector-message":  basePerm(core.CanSeeProjector),
		"core/countdown":          basePerm(core.CanSeeProjector),
		"core/tag":                restricter.ForAll,
		"core/config":             restricter.ForAll,

		"mediafiles/mediafile": mediafile.Restrict(ds),

		"motions/category":                     basePerm(motion.CanSee),
		"motions/statute-paragraph":            basePerm(motion.CanSee),
		"motions/motion":                       motion.Restrict(ds),
		"motions/motion-block":                 motion.BlockRestrict(ds),
		"motions/motion-comment-section":       motion.CommentSectionRestrict(ds),
		"motions/workflow":                     basePerm(motion.CanSee),
		"motions/motion-change-recommendation": motion.ChangeRecommendationRestrict(ds),
		"motions/motion-poll":                  poll.RestrictPoll(ds, motion.CanSee, motion.CanManage, nil),
		"motions/motion-option":                poll.RestrictOption(ds, motion.CanSee, motion.CanManage),
		"motions/motion-vote":                  poll.RestrictVote(ds, motion.CanSee, motion.CanManage),
		"motions/state":                        basePerm(motion.CanSee),

		"topics/topic": basePerm("agenda.can_see"),

		"users/user":          user.Restrict(ds),
		"users/group":         restricter.ForAll,
		"users/personal-note": restricter.ElementFunc(user.PersonalNoteRestrict),
	}
}

func openslidesProjectorCallables() map[string]projector.Callable {
	return map[string]projector.Callable{
		"agenda/item-list":                        agenda.ItemListSlide(),
		"agenda/list-of-speakers":                 agenda.ListOfSpeakersSlide(),
		"agenda/current-list-of-speakers":         agenda.CurrentListOfSpeakersSlide(),
		"agenda/current-list-of-speakers-overlay": agenda.CurrentListOfSpeakersSlide(),
		"agenda/current-speaker-chyron":           agenda.CurrentSpeakerChyronSlide(),

		"assignments/assignment":      assignment.Slide(),
		"assignments/assignment-poll": assignment.PollSlide(),

		"core/countdown":         core.CountdownSlide(),
		"core/projector-message": core.MessageSlide(),
		"core/clock":             core.ClockSlide(),

		"mediafiles/mediafile": mediafile.Slide(),

		"motions/motion":       motion.Slide(),
		"motions/motion-block": motion.SlideMotionBlock(),
		"motions/motion-poll":  motion.SlideMotionPoll(),

		"topics/topic": topic.Slide(),
		"users/user":   user.Slide(),
	}
}
