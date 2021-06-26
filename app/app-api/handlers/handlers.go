package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/bencodesall/ardanlabs-service-2.0/business/middleware"
	"github.com/bencodesall/ardanlabs-service-2.0/foundation/web"
)

// API constructs an http.Handler with all application routes defined
func API(build string, shutdown chan os.Signal, log *log.Logger) *web.App {

	app := web.NewApp(shutdown, middleware.Logger(log), middleware.Errors(log), middleware.Metrics(), middleware.Panics(log))

	check := check{
		log: log,
	}
	app.Handle(http.MethodGet, "/readiness", check.readiness)

	return app
}
