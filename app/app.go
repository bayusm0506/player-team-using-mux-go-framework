package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bayusm0506/player-team/config"
	"github.com/gorilla/mux"
)

// App has router
type App struct {
	Router *mux.Router
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.setRouters()
}

// SetRouters sets the all required routers
func (a *App) setRouters() {

}

// GET wraps the routers for GET method
func (a *App) GET(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the routers for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// PUT wraps the routers for PUT method
func (a *App) PUT(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// DELETE wraps the routers for DELETE method
func (a *App) DELETE(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the app on it's router
func (a *App) Run(config *config.Config) {
	port := fmt.Sprintf(":%s",
		config.APPS.Port,
	)

	log.Fatal(http.ListenAndServe(port, a.Router))
}

type RequestHandlerFunction func(w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}
}
