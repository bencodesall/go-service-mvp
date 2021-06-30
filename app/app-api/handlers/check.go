package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/bencodesall/ardanlabs-service-2.0/foundation/web"
)

type check struct {
	log *log.Logger
}

func (c check) readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	// RANDOM FAILURE TESTING
	//if n := rand.Intn(100); n%100 == 0 {
	//	// TESTS FOR MIDDLEWARE LOGGING
	//	// return errors.New("untrusted error")
	//	return web.NewRequestError(errors.New("trusted error"), http.StatusBadRequest)
	//	// panic("forcing panic")
	//	// return web.NewShutdownError("forcing shutdown")
	//}
	status := struct {
		Status string
	}{
		Status: "Ok",
	}

	//return json.NewEncoder(w).Encode(status)
	return web.Respond(ctx, w, status, http.StatusOK)
}
