package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/hvs-fasya/chusha/internal/engine"
)

//TabsGet get tabs list
func TabsGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	tabs, err := engine.DB.TabsGet()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error().Msgf("[DB] TabsGet error %v", err)
		errResp := ErrResponse{
			Errors: []string{http.StatusText(http.StatusInternalServerError)},
		}
		resp, _ := json.Marshal(errResp)
		w.Write([]byte(resp))
		return
	}
	w.WriteHeader(http.StatusOK)
	resp, _ := json.Marshal(tabs)
	w.Write(resp)
}
