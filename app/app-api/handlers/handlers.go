package handlers

import (
	"github.com/bencodesall/ardanlabs-service-2.0/business/auth"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"

	"github.com/bencodesall/ardanlabs-service-2.0/business/middleware"
	"github.com/bencodesall/ardanlabs-service-2.0/foundation/web"
)

// API constructs an http.Handler with all application routes defined
func API(build string, shutdown chan os.Signal, log *log.Logger, a *auth.Auth, db *sqlx.DB) *web.App {

	app := web.NewApp(shutdown, middleware.Logger(log), middleware.Errors(log), middleware.Metrics(), middleware.Panics(log))

	cg := checkGroup{
		build: build,
		db: db,
	}
	//app.Handle(http.MethodGet, "/readiness", check.readiness, middleware.Authenticate(a), middleware.Authorize(auth.RoleUser))
	//app.Handle(http.MethodGet, "/liveness", check.liveness, middleware.Authenticate(a), middleware.Authorize(auth.RoleUser))
	app.Handle(http.MethodGet, "/readiness", cg.readiness)
	app.Handle(http.MethodGet, "/liveness", cg.liveness)

	return app
}
