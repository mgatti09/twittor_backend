package routers

import (
	"encoding/json"
	"net/http"

	"github.com/mgatti09/twittor_backend/bd"
	"github.com/mgatti09/twittor_backend/models"
)

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Invalid Data: "+err.Error(), 400)
		return
	}

	var status bool
	//UserID recordar que este dato se almacena en router ProcessToken
	status, err = bd.UpdateUser(t, UserID)
	if err != nil {
		http.Error(w, "An error occurred while trying to modify the registry, please try again: "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "Unable to modify the user's registry", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
