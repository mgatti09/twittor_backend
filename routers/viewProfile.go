package routers

import (
	"encoding/json"
	"net/http"

	"github.com/mgatti09/twittor_backend/bd"
)

/*ViewProfile extraer los valores del perfil */
func ViewProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "ID key is required", http.StatusUnprocessableEntity)
		return
	}

	profile, err := bd.SearchProfile(ID)
	if err != nil {
		http.Error(w, "An error has ocurr while searching the record: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profile)
}
