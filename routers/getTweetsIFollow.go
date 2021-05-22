package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/mgatti09/twittor_backend/bd"
)

func GetTweetsIFollow(w http.ResponseWriter, r *http.Request) {
	pageParam := r.URL.Query().Get("page")
	pag, err := strconv.Atoi(pageParam)
	if err != nil || pag <= 0 {
		http.Error(w, "page must be an integer greater than 0", http.StatusUnprocessableEntity)
		return
	}
	page := int64(pag)

	response, statusOk := bd.GetTweetsIFollow(UserID, page)
	if !statusOk {
		http.Error(w, "Error reading tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
