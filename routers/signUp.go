package routers

import (
	"encoding/json" //Paquete propio de go que nos permite manipular JSONs
	"net/http"

	"github.com/mgatti09/twittor_backend/bd"
	"github.com/mgatti09/twittor_backend/models"
)

/* Registry es la función para crear en BD el registro del usuario*/
func SignUp(w http.ResponseWriter, r *http.Request) {
	var t models.User

	//El Body de un http.Request es un STREAM, es decir se lee una vez y se destruye.
	//.Decode(&t) con esta instrucción llenamos la data en la var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "There was a problem with the data received "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Password lenght must be greater than 6 characters", 400)
		return
	}

	_, userExists, _ := bd.UserExists(t.Email)
	if userExists {
		http.Error(w, "User already exists with the email provided", 400)
		return
	}

	_, status, err := bd.CreateUser(t)
	if err != nil {
		http.Error(w, "There was a problem creating the user "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "It was not possible to create the user", 400)
	}

	w.WriteHeader(http.StatusCreated)
}
