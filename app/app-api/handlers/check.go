package handlers

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/bencodesall/ardanlabs-service-2.0/foundation/web"
)

type check struct {
	log *log.Logger
	build string
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

func (c check) liveness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
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
		Build:     c.build,
		Host:      host,
		Pod:       os.Getenv("KUBERNETES_PODNAME"),
		PodIP:     os.Getenv("KUBERNETES_NAMESPACE_POD_IP"),
		Node:      os.Getenv("KUBERNETES_NODENAME"),
		Namespace: os.Getenv("KUBERNETES_NAMESPACE"),
	}

	return web.Respond(ctx, w, info, http.StatusOK)
}
