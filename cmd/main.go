package main

import (
	"context"
	"jobreport/internal/common"
	"jobreport/internal/database"
	"jobreport/internal/user"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

func main() {

	app := NewApp()

	app.Run()
}
func NewApp() *fx.App {

	return fx.New(
		fx.Provide(
			common.LoadPostgresDatabaseConfig,
			database.NewPostgresDatabase,
			user.NewUserService,
			user.NewUserDatabase,
			NewMux,
		),
		fx.Invoke(
			user.MakeLoginHandler,
		),
	)
}
func NewMux(lc fx.Lifecycle) *mux.Router {
	logrus.Info("creating mux")
	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("../static/jobreport"))))
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "PATCH", "OPTIONS", "DELETE"})
	cors := handlers.CORS(originsOk, headersOk, methodsOk)

	router.Use(cors)
	handler := (cors)((router))

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logrus.Info("starting server")
			go http.ListenAndServe(":8080", handler)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logrus.Info("stopping server")
			return nil
		},
	})
	return router
}
