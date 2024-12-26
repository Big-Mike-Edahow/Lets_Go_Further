// routes.go

package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
    router := httprouter.New()

    // Convert the notFoundResponse() helper to a http.Handler using the 
    // http.HandlerFunc() adapter, and then set it as the custom error handler for 404
    // Not Found responses.
    router.NotFound = http.HandlerFunc(app.notFoundResponse)

    // Likewise, convert the methodNotAllowedResponse() helper to a http.Handler and set
    // it as the custom error handler for 405 Method Not Allowed responses.
    router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/", app.indexHandler)
    router.HandlerFunc(http.MethodGet, "/healthcheck", app.healthcheckHandler)
    router.HandlerFunc(http.MethodGet, "/view/:id", app.viewHandler)
	router.HandlerFunc(http.MethodPost, "/create", app.createHandler)

	// Wrap the router with the panic recovery middleware.
    return app.recoverPanic(router)
}
