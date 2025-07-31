package routes

import (
	"github.com/navid-fn/shorty/internal/app"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)
  r.Post("/url/create/", app.UrlHandler.HandleCreateUrl)

	return r

}
