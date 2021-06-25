package handlers

import (
	"context"
	"github.com/bencodesall/ardanlabs-service-2.0/foundation/web"
	"github.com/pkg/errors"
	"log"
	"math/rand"
	"net/http"
)

type check struct {
	log *log.Logger
}

func (c check) readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if n := rand.Intn(100); n%100 == 0 {
		return errors.New("untrusted error")
	}
	status := struct {
		Status string
	}{
		Status: "Ok",
	}
	log.Println(r, status)
	//return json.NewEncoder(w).Encode(status)
	return web.Respond(ctx, w, status, http.StatusOK)
}
