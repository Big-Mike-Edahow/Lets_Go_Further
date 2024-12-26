// healthcheck.go

package main

import (
	"fmt"
	"net/http"
	"time"
)

func (app *application) indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	fmt.Fprintf(w, "Greenlight movie database.\n")
}

func (app *application) viewHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
    if err != nil {
        // Use the new notFoundResponse() helper.
        app.notFoundResponse(w, r)
        return
    }

    movie := Movie{
        Id:        id,
        Created:   time.Now(),
        Title:     "Casablanca",
        Runtime:   102,
        Genres:    []string{"drama", "romance", "war"},
        Version:   1,
    }

    err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
    if err != nil {
        // Use the new serverErrorResponse() helper.
        app.serverErrorResponse(w, r, err)
    }
}

func (app *application) createHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new movie.")
}

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
    env := envelope{
        "status": "available",
        "system_info": map[string]string{
            "environment": app.config.env,
            "version":     version,
        },
    }

    err := app.writeJSON(w, http.StatusOK, env, nil)
    if err != nil {
        // Use the new serverErrorResponse() helper.
        app.serverErrorResponse(w, r, err)
    }
}
