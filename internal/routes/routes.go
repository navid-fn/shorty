package routes

import (
	"github.com/navid-fn/shorty/internal/app"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()
	
	// url
	r.Get("/health", app.HealthCheck)
  r.Post("/url/create/", app.UrlHandler.HandleCreateUrl)
  r.Get("/url/redirect/{code}", app.UrlHandler.HandleRedirectUrl)


  // user

	return r

}
