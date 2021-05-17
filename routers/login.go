package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/mgatti09/twittor_backend/bd"
	"github.com/mgatti09/twittor_backend/jwt"
	"github.com/mgatti09/twittor_backend/models"
)

/* Login realiza el Login*/
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "User or Password invalid: "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}

	doc, exists := bd.Login(t.Email, t.Password)
	if !exists {
		http.Error(w, "User or Password are invalid", 400)
	}

	jwtToken, err := jwt.GetToken(doc)
	if err != nil {
		http.Error(w, "There was a problem with the Token generation: "+err.Error(), 400)
		return
	}

	resp := models.LoginResponse{
		Token: jwtToken,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

	//Seteando una cookie para el token
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtToken,
		Expires: expirationTime,
	})
}
