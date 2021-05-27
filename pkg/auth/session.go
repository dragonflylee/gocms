package auth

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/gob"
	"io"
	"log"
	"net/http"
	"os"

	"gocms/model"
	"gocms/pkg/config"

	"github.com/gorilla/sessions"
)

type sessKey int

const (
	_ sessKey = iota
	userKey
)

// Session checks if there is a user uid stored in the session and returns the user
// object for that uid.
type Session struct {
	name  string
	store sessions.Store
}

// Init does nothing as the Session implementation does not need to allocate any resources
func (s *Session) Init() error {
	cfg := config.Session()

	secret, err := base64.StdEncoding.DecodeString(cfg.Secret)
	if err != nil {
		secret = make([]byte, 32)
		_, err = io.ReadFull(rand.Reader, secret)
		if err != nil {
			return err
		}
		cfg.Secret = base64.StdEncoding.EncodeToString(secret)
	}

	switch cfg.Type {
	default:
		s.store = sessions.NewFilesystemStore(os.TempDir(), secret)
	}

	s.name = cfg.Name
	return nil
}

// IsEnabled returns true as this plugin is enabled by default and its not possible to disable
func (s *Session) IsEnabled() bool {
	return true
}

// Verify checks if there is a user uid stored in the session and returns the user
func (s *Session) Verify(r *http.Request, w http.ResponseWriter) *model.Admin {
	sess, err := s.store.Get(r, s.name)
	if err != nil {
		return nil
	}
	cookie, exist := sess.Values[userKey]
	if !exist {
		return nil
	}
	user, ok := cookie.(*model.Admin)
	if !ok {
		return nil
	}
	return user
}

// Set sets value in session.
func (s *Session) Set(r *http.Request, w http.ResponseWriter, user *model.Admin) error {
	sess, err := s.store.Get(r, s.name)
	if err != nil {
		log.Printf("Get Session %s failed: %v", s.name, err)
	}
	sess.Values[userKey] = user
	return sess.Save(r, w)
}

// Flush deletes all session data.
func (s *Session) Flush(r *http.Request, w http.ResponseWriter) error {
	sess, err := s.store.Get(r, s.name)
	if err != nil {
		return err
	}
	sess.Options.MaxAge = -1
	return sess.Save(r, w)
}

func init() {
	gob.Register(new(model.Admin))
	gob.Register(userKey)
}
