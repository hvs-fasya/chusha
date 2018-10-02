package api

import (
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
)

// Server is api access server
type Server struct {
	//DB         engine.DBInterface
	httpServer *http.Server
}

// Run router
func (s *Server) Run(connstr string) {
	log.Info().Msg("Запуск сервера на " + connstr)
	e := http.ListenAndServe(connstr, NewRouter(Routes))
	if e != nil {
		log.Fatal().Err(e).Msg("Ошибка запуска")
	}
}

// Endpoint эндоинт.
type Endpoint struct {
	Name        string   // Имя ендпоинта.
	Description string   // Краткое описане ендпоинта.
	Path        string   // URI Ендпоинта - не должен быть пустым. Иначе не добавиться при инициализации. Если всё-же пустой, то будет попытка добаления по PathPrefix.
	PathPrefix  string   // Prefix URI Ендпоинта - не должен быть пустым. Иначе не добавиться при инициализации.
	Methods     []string // HTTP Методы, которые обслуживает ендпоинт.
	Handler     Handler  // HTTP обработчик
}

// Handler Обработчик ендпоинта.
type Handler func(http.ResponseWriter, *http.Request)

// ServeHTTP Вызов обработчика HTTP запроса.
func (endpoint Endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Вызов обработчика HTTP запросов
	if endpoint.Handler != nil {
		endpoint.Handler(w, r)
	}
}

// Router - роутер вебсервера
type Router struct {
	mux *mux.Router
}

// NewRouter Создать - новый роутер
func NewRouter(endpoints []Endpoint) *Router {
	rt := Router{
		mux: mux.NewRouter().StrictSlash(true),
	}

	for i := range endpoints {
		if "" != endpoints[i].Path {
			rt.mux.Methods(endpoints[i].Methods...).Name(endpoints[i].Name).Path(endpoints[i].Path).Handler(endpoints[i])
			continue
		}

		if "" != endpoints[i].PathPrefix {
			rt.mux.Methods(endpoints[i].Methods...).Name(endpoints[i].Name).PathPrefix(endpoints[i].PathPrefix).Handler(http.StripPrefix(endpoints[i].PathPrefix, endpoints[i]))
			continue
		}
	}

	return &rt
}

// ServeHTTP Вызов обработчика HTTP запроса.
func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "private, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "-1")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept,Authorization,Cache-Control,Content-Type,DNT,Expires,If-Modified-Since,Keep-Alive,Origin,Pragma,User-Agent,X-Requested-With,X-Initiator-User")

	rt.mux.ServeHTTP(w, r)
}
