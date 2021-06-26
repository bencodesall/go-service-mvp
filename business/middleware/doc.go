// Package middleware contains the set of middleware functions.
package middleware

/*
package middleware

import (
	"context"
	"github.com/bencodesall/ardanlabs-service-2.0/foundation/web"
	"log"
	"net/http"
)

// Errors ...
func Errors(log *log.Logger) web.Middleware {
	// This is the actual middleware function to be executed.
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			if err := handler(ctx, w, r); err != nil {

				// TRACE ID

				// HANDLE ERROR
				log.Println(err)

				// HANDLE ERROR RESPONSE

				// SHUTDOWN SIGNAL

				// STOP HERE
				return nil
			}

			return nil
		}

		return h
	}

	return m
}

*/
