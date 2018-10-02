package api

import (
	"github.com/hvs-fasya/chusha/internal/api/handlers"
)

var (
	// Routes - Роуты - коллекция ендпоинтов для роутера веб сервера
	Routes = []Endpoint{
		Endpoint{
			Description: "Простой ответ от сервера. Может быть использован для проверки жив - не жив.",
			Path:        "/alive",
			Methods:     []string{"GET"},
			Handler:     handlers.Alive,
		},
		Endpoint{
			Description: "Получение перечня вкладок",
			Path:        "/api/v1/tabs",
			Methods:     []string{"GET", "OPTIONS"},
			Handler:     handlers.TabsGet,
		},
	}
)
