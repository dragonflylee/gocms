package auth

import (
	"log"
	"net/http"
	"reflect"

	"gocms/model"
)

// Auth represents a authentication method (plugin) for HTTP requests.
type Auth interface {
	// Init should be called exactly once before using any of the other methods,
	// in order to allow the plugin to allocate necessary resources
	Init() error

	// IsEnabled checks if the current method has been enabled in settings.
	IsEnabled() bool

	// Verify tries to verify the authentication data contained in the request.
	// If verification is successful returns either an existing user object (with id > 0)
	// or a new user object (with id = 0) populated with the information that was found
	// in the authentication data (username or email).
	// Returns nil if verification fails.
	Verify(r *http.Request, w http.ResponseWriter) *model.Admin
}

var authMethods = make([]Auth, 0)

// SignedIn returns the user object of signed user.
// It returns a bool value to indicate whether user uses basic auth or not.
func SignedIn(r *http.Request, w http.ResponseWriter) *model.Admin {
	if !model.IsOpen() {
		return nil
	}
	// Try to sign in with each of the enabled plugins
	for _, auth := range authMethods {
		if !auth.IsEnabled() {
			continue
		}
		user := auth.Verify(r, w)
		if user != nil {
			return user
		}
	}
	return nil
}

// Register adds the specified instance to the list of available Auth methods
func Register(method Auth) {
	authMethods = append(authMethods, method)
}

// Init should be called exactly once when the application starts to allow Auth plugins
// to allocate necessary resources
func Init() {
	for _, method := range authMethods {
		if err := method.Init(); err != nil {
			log.Fatalf("init '%s' auth method failed: %s", reflect.TypeOf(method), err)
		}
	}
}
