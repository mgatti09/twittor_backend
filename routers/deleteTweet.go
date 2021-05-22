package routers

import (
	"net/http"

	"github.com/mgatti09/twittor_backend/bd"
)

func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	tweetID := r.URL.Query().Get("id")
	if len(tweetID) < 1 {
		http.Error(w, "id key is required", http.StatusUnprocessableEntity)
		return
	}

	//UserID recordar que este dato se almacena en router ProcessToken
	err := bd.DeleteTweet(tweetID, UserID)
	if err != nil {
		http.Error(w, "An error occurred while trying to delete the tweet: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}
