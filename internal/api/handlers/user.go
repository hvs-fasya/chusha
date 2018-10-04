package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"fmt"
	"github.com/hvs-fasya/chusha/internal/models"
)

//UserLogin login user
func UserLogin(w http.ResponseWriter, r *http.Request) {
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte{})
		return
	}
	user := new(models.User)
	err = json.Unmarshal(payload, &user)
	sessionToken := `maqhjzdysjxnvrct1dpjwi1t`
	http.SetCookie(w, &http.Cookie{
		Name:    SessionCookieName,
		Value:   sessionToken,
		Expires: time.Now().Add(time.Duration(SessionCookieExpirationTime) * time.Minute),
		//Secure:   true,
		//HttpOnly: true,
	})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte{})
}

//UserLogout logout user
func UserLogout(w http.ResponseWriter, r *http.Request) {
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
