package routes

import (
	"github.com/navid-fn/shorty/internal/app"
	authmiddleware "github.com/navid-fn/shorty/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()
	
	r.Get("/health", app.HealthCheck)

  // user
  r.Post("/api/v1/register/", app.UserHandler.Register)
  r.Post("/api/v1/login/", app.UserHandler.Login)

  r.Get("{code}", app.UrlHandler.HandleRedirectUrl)


	r.Group(func(r chi.Router) {
		// middlewares
		r.Use(authmiddleware.AuthMiddleware)

  	r.Post("/api/v1/shorten/", app.UrlHandler.HandleCreateUrl)
	})



	return r

}
