package traefik_header_setter

import (
	"context"
	"fmt"
	"net/http"
)

// Config the plugin configuration.
type Config struct {
	Create string `json:"create,omitempty"` // creating a new header for store extracted data from the host
}

// CreateConfig creates and initializes the plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

// New creates and returns a plugin instance.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if len(config.Create) == 0 {
		return nil, fmt.Errorf("Create can't be empty")
	}

	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		req.Header.Set("Host", config.Create)
		next.ServeHTTP(rw, req)
	}), nil
