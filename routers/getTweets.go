package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/mgatti09/twittor_backend/bd"
)

/*GetTweets request */
func GetTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "id key is required", http.StatusUnprocessableEntity)
		return
	}

	pageParam := r.URL.Query().Get("page")
	if len(pageParam) < 1 {
		http.Error(w, "page key is required", http.StatusUnprocessableEntity)
		return
	}

	pag, err := strconv.Atoi(pageParam)
	if err != nil {
		http.Error(w, "page must be an integer greater than 0", http.StatusUnprocessableEntity)
		return
	}

	page := int64(pag)
	response, getOK := bd.GetTweets(ID, page)
	if !getOK {
		http.Error(w, "Error getting tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
