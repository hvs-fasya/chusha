package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/satori/go.uuid"

	"github.com/hvs-fasya/chusha/internal/models"
)

//UserCreate register user
func UserCreate(w http.ResponseWriter, r *http.Request) {
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte{})
		return
	}
	user := new(models.User)
	err = json.Unmarshal(payload, &user)
	sessionToken := uuid.NewV4().String()
	http.SetCookie(w, &http.Cookie{
		Name:     SessionCookieName,
		Value:    sessionToken,
		Expires:  time.Now().Add(time.Duration(SessionCookieExpirationTime) * time.Minute),
		Secure:   true,
		HttpOnly: true,
	})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte{})
}
