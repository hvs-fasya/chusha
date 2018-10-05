package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/satori/go.uuid"

	"database/sql"
	"github.com/hvs-fasya/chusha/internal/engine"
)

//SessionCreate login user
func SessionCreate(w http.ResponseWriter, r *http.Request) {
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte{})
		return
	}
	loginData := struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}{}
	err = json.Unmarshal(payload, &loginData)
	user, err := engine.DB.UserCheck(loginData.Login, loginData.Password)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			w.WriteHeader(http.StatusUnauthorized)
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	fmt.Printf("%+v", user)
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

//SessionDestroy logout user
func SessionDestroy(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(SessionCookieName)
	if err != nil {
		switch err {
		case http.ErrNoCookie:
			w.WriteHeader(http.StatusUnauthorized)
			return
		default:
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	fmt.Println(c.Value)
	sessionToken := c.Value
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(sessionToken))
}
