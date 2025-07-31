package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/navid-fn/shorty/internal/api"
	"github.com/navid-fn/shorty/internal/store"
	"github.com/navid-fn/shorty/migrations"
)

type Application struct {
	Logger      *log.Logger
	UserHandler *api.UserHandler
	UrlHandler  *api.UrlHandler
	DB          *sql.DB
}

func NewApplication() (*Application, error) {
	pgxdb, err := store.Open()
	if err != nil {
		return nil, err
	}
	err = store.MigrateFS(pgxdb, migrations.FS, ".")
	if err != nil {
		panic(err)
	}
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// store here
	urlStore := store.NewPostgresUrlStore(pgxdb)

	// handler here
	urlHandler := api.NewUrlHandler(urlStore)

	app := &Application{
		Logger:     logger,
		DB:         pgxdb,
		UrlHandler: urlHandler,
	}
	return app, nil
}

func (app *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "status is Available\n")
}
