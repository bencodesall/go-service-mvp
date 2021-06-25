package web

import (
	"context"
	"net/http"
	"os"
	"syscall"

	"github.com/dimfeld/httptreemux/v5"
)

// Handler is a type that handles an http request within our own mini framework
// using an "onion" method
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

// App is the entrypoint into our applicatino and what configures our context
// object for each of our http handlers. Feel free to add any configuration
// data/logic on this App struct. App is everything that a ContextMux is by
// virtue of promotion.
type App struct {
	*httptreemux.ContextMux
	shutdown chan os.Signal
}

// NewApp creates an App value that handles a set of routes for the application.
func NewApp(shutdown chan os.Signal) *App {
	app := App{
		ContextMux: httptreemux.NewContextMux(),
		shutdown:   shutdown,
	}

	return &app
}

// SignalShutdown is used to gracefully shutdown the app when an integrity
// issue is identified.
func (a *App) SignalShutdown() {
	a.shutdown <- syscall.SIGTERM
}

// Handle ...
func (a *App) Handle(method string, path string, handler Handler) {
	h := func(w http.ResponseWriter, r *http.Request) {

		// BOILERPLATE

		if err := handler(r.Context(), w, r); err != nil {
			a.SignalShutdown()
			return
		}

		// BOILERPLATE
	}
	a.ContextMux.Handle(method, path, h)
}
