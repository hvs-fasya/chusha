package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"github.com/satori/go.uuid"

	"github.com/hvs-fasya/chusha/internal/engine"
	"github.com/hvs-fasya/chusha/internal/models"
	"github.com/hvs-fasya/chusha/internal/utils"
)

//UserRegister register user
func UserRegister(w http.ResponseWriter, r *http.Request) {
	//todo: all user attributes at frontend
	var err error
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error().Msgf("http USER CREATE read body error: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		errorToResponse("failed to read request body", HumanInternalError)
		return
	}
	user := new(models.UserNewInput)
	err = json.Unmarshal(payload, &user)
	if err != nil {
		log.Error().Msgf("http USER CREATE unmarshal body error: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		errorToResponse("failed to unmarshal request body", HumanInternalError)
		return
	}
	errs := user.Validate()
	if len(errs) != 0 {
		w.WriteHeader(http.StatusBadRequest)
		errResp := ErrResponse{
			Errors:      []string{},
			HumanErrors: []string{},
		}
		for _, e := range errs {
			errResp.Errors = append(errResp.Errors, e)
			errResp.HumanErrors = append(errResp.HumanErrors, e)
		}
		resp, _ := json.Marshal(errResp)
		w.Write(resp)
		return
	}
	user.Role = new(models.RoleDB)
	user.Role, err = engine.DB.RoleGetByName(utils.ClientRoleName)
	if err != nil {
		log.Error().Msgf("database GET ROLE BY NAME request error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		respond500(w)
		return
	}
	err = engine.DB.UserCreate(user)
	if err != nil {
		if pgerr, ok := err.(*pq.Error); ok {
			if pgerr.Code == "23503" { //unique key login or email violation
				w.WriteHeader(http.StatusBadRequest)
				w.Write(errorToResponse("login or email "+ErrAlreadyExist, "логин или email "+ErrAlreadyExistRus))
				return
			}
			log.Error().Msgf("database USER CREATE request error: %s", err)
			respond500(w)
			return
		}
		log.Error().Msgf("database USER CREATE request error: %s", err)
		respond500(w)
		return
	}
	sessionToken := uuid.NewV4().String()
	http.SetCookie(w, &http.Cookie{
		Name:     SessionCookieName,
		Value:    sessionToken,
		Expires:  time.Now().Add(time.Duration(SessionCookieExpirationTime) * time.Minute),
		Secure:   true,
		HttpOnly: true,
	})
	//todo: put session token to redis with key=new user login????
	w.WriteHeader(http.StatusOK)
	w.Write([]byte{})
}
