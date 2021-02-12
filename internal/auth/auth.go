package auth

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const (
	whoAmIPath    = "/apps/users/whoami/"
	whoamITimeout = 2 * time.Second
	saltPrefix    = "openslides.utils.sessions.SessionStore"
)

// Backend gives the sessionData for a sessionID
type Backend interface {
	GetSession(sessionID string) ([]byte, error)
}

// Configer returns the value for a config name.
type Configer interface {
	ConfigValue(key string, v interface{}) error
}

// Auth authentivates a request using the whoami view.
type Auth struct {
	cookieName string
	key        []byte
	backend    Backend
	configer   Configer
}

// New creates a new Auth instance.
func New(cookieName string, secret string, backend Backend, configer Configer) *Auth {
	bs := sha1.Sum([]byte(saltPrefix + secret))
	return &Auth{
		cookieName: cookieName,
		backend:    backend,
		configer:   configer,
		key:        []byte(bs[:]),
	}
}

func (a *Auth) validateSessionData(original, value []byte) (bool, error) {
	mac := hmac.New(sha1.New, a.key)
	if _, err := mac.Write(value); err != nil {
		return false, fmt.Errorf("writing hmac")
	}

	originalDecoded := make([]byte, hex.DecodedLen(len(original)))
	n, err := hex.Decode(originalDecoded, original)
	if err != nil {
		return false, err
	}

	originalDecoded = originalDecoded[:n]
	expectedHash := mac.Sum(nil)
	return hmac.Equal(originalDecoded, expectedHash), nil
}

func (a *Auth) userID(r *http.Request) (int, error) {
	cookie, err := r.Cookie(a.cookieName)
	if err != nil {
		if err == http.ErrNoCookie {
			// Anonymous.
			return 0, nil
		}
		return 0, fmt.Errorf("loading session cookie: %v", err)
	}

	encodedSessionData, err := a.backend.GetSession(cookie.Value)
	if encodedSessionData == nil {
		return 0, Error("Invalid session data")
	}

	b64decoded := make([]byte, base64.StdEncoding.DecodedLen(len(encodedSessionData)))
	n, err := base64.StdEncoding.Decode(b64decoded, encodedSessionData)
	if err != nil {
		return 0, fmt.Errorf("base64 decode session data: %v", err)
	}
	b64decoded = b64decoded[:n]

	parts := bytes.SplitN(b64decoded, []byte(":"), 2)
	valid, err := a.validateSessionData(parts[0], parts[1])
	if !valid {
		return 0, Error("Invalid session data")
	}

	var sessionData struct {
		UserID string `json:"_auth_user_id"`
	}
	if err := json.Unmarshal(parts[1], &sessionData); err != nil {
		return 0, fmt.Errorf("json decoding session data: %v", err)
	}

	uid, err := strconv.Atoi(sessionData.UserID)
	if err != nil {
		return 0, fmt.Errorf("invalid user id: %s", sessionData.UserID)
	}

	return uid, nil
}

// Authenticate returns an context that can be used by auth.FromContext().
func (a *Auth) Authenticate(r *http.Request) (context.Context, error) {
	uid, err := a.userID(r)
	if err != nil {
		return nil, fmt.Errorf("getting user id: %v", err)
	}

	if uid == 0 {
		var enabled bool
		if err := a.configer.ConfigValue("general_system_enable_anonymous", &enabled); err != nil {
			return nil, fmt.Errorf("getting config value for anonymous: %v", err)
		}
		if !enabled {
			return nil, Error("Anonymous is not enabled.")
		}
	}

	return context.WithValue(r.Context(), userIDType, uid), nil
}

// FromContext returnes the user id from a context that was called insinde the
// Middleware.
func FromContext(ctx context.Context) int {
	v := ctx.Value(userIDType)
	if v == nil {
		return 0
	}

	return v.(int)
}

type authString string

const userIDType authString = "user_id"
