package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/mgatti09/twittor_backend/bd"
)

func GetListUsers(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")

	userType := r.URL.Query().Get("type")
	if userType != "new" && userType != "follow" {
		http.Error(w, "userType value is invalid. Must be 'new' or 'follow'", http.StatusUnprocessableEntity)
		return
	}
	pageParam := r.URL.Query().Get("page")

	pag, err := strconv.Atoi(pageParam)
	if err != nil || pag <= 0 {
		http.Error(w, "page must be an integer greater than 0", http.StatusUnprocessableEntity)
		return
	}
	page := int64(pag)

	result, status := bd.GetListUsers(UserID, page, search, userType)
	if !status {
		http.Error(w, "Error fetching users", http.StatusUnprocessableEntity)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}
