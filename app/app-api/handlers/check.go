package handlers

import (
	"context"
	"github.com/bencodesall/ardanlabs-service-2.0/foundation/database"
	"github.com/jmoiron/sqlx"
	"net/http"
	"os"

	"github.com/bencodesall/ardanlabs-service-2.0/foundation/web"
)

type checkGroup struct {
	build string
	db *sqlx.DB
}

func (cg checkGroup) readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	// RANDOM FAILURE TESTING
	//if n := rand.Intn(100); n%100 == 0 {
	//	// TESTS FOR MIDDLEWARE LOGGING
	//	// return errors.New("untrusted error")
	//	return web.NewRequestError(errors.New("trusted error"), http.StatusBadRequest)
	//	// panic("forcing panic")
	//	// return web.NewShutdownError("forcing shutdown")
	//}
	status := "ok"
	statusCode := http.StatusOK
	if err := database.StatusCheck(ctx, cg.db); err != nil {
		status = "db not ready"
		statusCode = http.StatusInternalServerError
	}

	health := struct {
		Status string `json:"status"`
	}{
		Status: status,
	}

	return web.Respond(ctx, w, health, statusCode)
}

func (cg checkGroup) liveness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	host, err := os.Hostname()
	if err != nil {
		host = "unavailable"
	}

	info := struct {
		Status    string `json:"status,omitempty"`
		Build     string `json:"build,omitempty"`
		Host      string `json:"host,omitempty"`
		Pod       string `json:"pod,omitempty"`
		PodIP     string `json:"podIP,omitempty"`
		Node      string `json:"node,omitempty"`
		Namespace string `json:"namespace,omitempty"`
	}{
		Status:    "up",
		Build:     cg.build,
		Host:      host,
		Pod:       os.Getenv("KUBERNETES_PODNAME"),
		PodIP:     os.Getenv("KUBERNETES_NAMESPACE_POD_IP"),
		Node:      os.Getenv("KUBERNETES_NODENAME"),
		Namespace: os.Getenv("KUBERNETES_NAMESPACE"),
	}

	return web.Respond(ctx, w, info, http.StatusOK)
}
