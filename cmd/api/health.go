package main

import (
	"net/http"
)

func (app *application) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
	app.store.Posts.Create(r.Context(), nil)
}
