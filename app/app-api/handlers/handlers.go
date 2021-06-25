package handlers

import (
	"encoding/json"
	"github.com/dimfeld/httptreemux/v5"
	"log"
	"net/http"
	"os"
)

// API constructs an http.Handler with all application routes defined
func API(build string, shutdown chan os.Signal, log *log.Logger) *httptreemux.ContextMux {

	tm := httptreemux.NewContextMux()
	h := func(w http.ResponseWriter, r *http.Request) {
		status := struct {
			Status string
		}{
			Status: "Ok",
		}
		json.NewEncoder(w).Encode(status)
	}
	tm.Handle(http.MethodGet, "/test", h)
	return tm
}