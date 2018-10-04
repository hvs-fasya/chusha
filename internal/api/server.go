package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hvs-fasya/chusha/internal/api/handlers"
	"github.com/hvs-fasya/chusha/internal/api/handlers/front"
	"github.com/rs/zerolog/log"
)

// Server is http server
type Server struct {
	httpServer *http.Server
}

// Run server
func (s *Server) Run(connstr string) {
	log.Info().Msg("Запуск сервера на " + connstr)
	e := http.ListenAndServe(connstr, NewRouter())
	if e != nil {
		log.Fatal().Err(e).Msg("Ошибка запуска")
	}
}

// NewRouter Создать - новый роутер
func NewRouter() *mux.Router {
	rt := new(mux.Router)
	//front
	rt.HandleFunc("/alive", handlers.Alive).Methods("GET")
	rt.HandleFunc("/", front.IndexHandler).Methods("GET")
	rt.HandleFunc(`/{file:favicon.+}`, front.FaviconHandler).Methods("GET")
	rt.HandleFunc(`/{js:.+\.js}`, front.JSHandler).Methods("GET")
	rt.HandleFunc("/fonts/{font}", front.FontsHandler).Methods("GET")
	//api
	apiRouter := rt.PathPrefix(("/api/v1")).Subrouter()
	apiRouter.Use(setApiHeaders)
	apiRouter.Use(respondOptions)

	apiRouter.HandleFunc("/tabs", handlers.TabsGet).Methods("GET", "OPTIONS")
	apiRouter.HandleFunc("/login", handlers.UserLogin).Methods("POST", "OPTIONS")
	apiRouter.HandleFunc("/logout", handlers.UserLogout).Methods("GET", "OPTIONS")
	return rt
}

func setApiHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Cache-Control", "private, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "-1")

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept,Authorization,Cache-Control,Content-Type,DNT,Expires,If-Modified-Since,Keep-Alive,Origin,Pragma,User-Agent,X-Requested-With,X-Initiator-User")
		next.ServeHTTP(w, r)
	})
}

func respondOptions(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte{})
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
