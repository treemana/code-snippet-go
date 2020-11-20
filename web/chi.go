package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func CHI() {
	app := chi.NewRouter()
	app.Use(middleware.Logger)
	app.Get("/chi", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("welcome"))
	})
	_ = http.ListenAndServe(":8080", app)
}
