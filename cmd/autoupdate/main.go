package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/auth"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/autoupdate"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/datastore"
	ahttp "github.com/OpenSlides/openslides3-autoupdate-service/internal/http"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/redis"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/restricter"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/restricter/agenda"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/restricter/assignment"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/restricter/mediafile"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/restricter/motion"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/restricter/user"
)

func main() {
	listenAddr := getEnv("LISTEN_HTTP_ADDR", ":8002")
	workerAddr := getEnv("WORKER_ADDR", "http://localhost:8000")

	redisConn := redis.New(getEnv("REDIS_ADDR", "localhost:6379"))

	requiredUserCallable := map[string]func(json.RawMessage) ([]int, error){
		"agenda/list-of-speakers":       agenda.RequiredSpeakers,
		"assignments/assignment":        assignment.RequiredAssignments,
		"assignments/assignment-option": assignment.RequiredPollOption,
		"motions/motion":                motion.RequiredMotions,
		"motions/motion-option":         motion.RequiredPollOption,
	}

	ds, err := datastore.New(workerAddr, redisConn, requiredUserCallable)
	if err != nil {
		log.Fatalf("Can not initialize data: %v", err)
	}

	basePerm := restricter.BasePermission(ds)

	restricter := restricter.New(ds, map[string]restricter.Element{
		"agenda/item":             agenda.Restrict(ds),
		"agenda/list-of-speakers": basePerm("agenda.can_see_list_of_speakers"),

		"assignments/assignment":        basePerm("assignments.can_see"),
		"assignments/assignment-poll":   basePerm("assignments.can_see"),
		"assignments/assignment-option": basePerm("assignments.can_see"),
		"assignments/assignment-vote":   basePerm("assignments.can_see"),

		"core/projector":          basePerm("core.can_see_projector"),
		"core/projection-default": basePerm("core.can_see_projector"),
		"core/projector-message":  basePerm("core.can_see_projector"),
		"core/countdown":          basePerm("core.can_see_projector"),
		"core/tag":                restricter.ForAll,
		"core/config":             restricter.ForAll,

		"mediafiles/mediafile": mediafile.Restrict(ds),

		"motions/category":                     basePerm("motions.can_see"),
		"motions/statute-paragraph":            basePerm("motions.can_see"),
		"motions/motion":                       motion.Restrict(ds),
		"motions/motion-block":                 motion.BlockRestrict(ds),
		"motions/motion-comment-section":       motion.CommentSectionRestrict(ds),
		"motions/workflow":                     basePerm("motions.can_see"),
		"motions/motion-change-recommendation": motion.ChangeRecommendationRestrict(ds),
		"motions/motion-poll":                  basePerm("motions.can_see"),
		"motions/motion-option":                basePerm("motions.can_see"),
		"motions/motion-vote":                  basePerm("motions.can_see"),
		"motions/state":                        basePerm("motions.can_see"),

		"topics/topic": basePerm("agenda.can_see"),

		"users/user":          user.Restrict(ds),
		"users/group":         restricter.ForAll,
		"users/personal-note": restricter.ElementFunc(user.PersonalNoteRestrict),
	})

	service, err := autoupdate.New(ds, restricter)
	if err != nil {
		log.Fatalf("Can not create autoupdate service: %v", err)
	}

	auth := auth.New(workerAddr)
	handler := ahttp.New(service, auth)

	srv := &http.Server{Addr: listenAddr, Handler: handler}
	defer func() {
		service.Close()
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("Error on HTTP server shutdown: %v", err)
		}
	}()

	go func() {
		fmt.Printf("Listen on %s\n", listenAddr)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("HTTP server failed: %v", err)
		}
	}()

	waitForShutdown()
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
