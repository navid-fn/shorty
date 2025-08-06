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
  r.Post("/register/", app.UserHandler.Register)
  r.Post("/login/", app.UserHandler.Login)

  r.Get("/url/redirect/{code}", app.UrlHandler.HandleRedirectUrl)


	r.Group(func(r chi.Router) {
		// middlewares
		r.Use(authmiddleware.AuthMiddleware)

  	r.Post("/url/create/", app.UrlHandler.HandleCreateUrl)
	})



	return r

}
