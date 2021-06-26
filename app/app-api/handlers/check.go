package handlers

import (
	"context"
	"log"
	"math/rand"
	"net/http"

	"github.com/bencodesall/ardanlabs-service-2.0/foundation/web"
	"github.com/pkg/errors"
)

type check struct {
	log *log.Logger
}

func (c check) readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	if n := rand.Intn(100); n%100 == 0 {
		// TESTS FOR MIDDLEWARE LOGGING
		// return errors.New("untrusted error")
		return web.NewRequestError(errors.New("trusted error"), http.StatusBadRequest)
		// panic("forcing panic")
		// return web.NewShutdownError("forcing shutdown")
	}
	status := struct {
		Status string
	}{
		Status: "Ok",
	}

	//return json.NewEncoder(w).Encode(status)
	return web.Respond(ctx, w, status, http.StatusOK)
}
