package main

import (
	"context"
	"fmt"
	"jobreport/internal/common"
	"jobreport/internal/database"
	"jobreport/internal/reports"
	"jobreport/internal/user"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
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
			reports.NewReportService,
			reports.NewReportDatabase,
			NewMux,
		),
		fx.Invoke(
			user.MakeLoginHandler,
			reports.MakeReportHandler,
		),
	)
}
func NewMux(lc fx.Lifecycle) *mux.Router {
	logrus.Info("creating mux")
	router := mux.NewRouter()
	//router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("../static/jobreport"))))
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "PATCH", "OPTIONS", "DELETE"})
	cors := handlers.CORS(originsOk, headersOk, methodsOk)

	//router.Use(cors, authenticate)
	router.Use(cors)
	//handler := (cors)((authenticate)(router))
	handler := (cors)((router))

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logrus.Info("starting server")
			port := os.Getenv("PORT")
			if port == "" {
				port = "8080"
			}
			logrus.Info("starting server %S", port)
			go http.ListenAndServe(":"+port, handler)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logrus.Info("stopping server")
			return nil
		},
	})
	return router
}

func authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "login") && r.Method == "POST" {
			next.ServeHTTP(w, r) // call original
			return
		}
		cookie := r.Cookies()
		tokenString := cookie[0].Value
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte("secretKey"), nil
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			next.ServeHTTP(w, r) // call original
		} else {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}
	})
}
