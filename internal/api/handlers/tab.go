package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/rs/zerolog/log"

	"github.com/hvs-fasya/chusha/internal/engine"
)

//TabsGet get tabs list
func TabsGet(w http.ResponseWriter, r *http.Request) {
	var enabled bool
	enabled, err := strconv.ParseBool(r.URL.Query().Get("enabled"))
	if err != nil {
		enabled = true
	}
	tabs, err := engine.DB.TabsGet(enabled)
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
