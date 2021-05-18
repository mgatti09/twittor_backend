package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/mgatti09/twittor_backend/bd"
	"github.com/mgatti09/twittor_backend/models"
)

/*InsertTweet Permite grabar el tweet en la BD*/
func InsertTweet(w http.ResponseWriter, r *http.Request) {
	var message models.Tweet
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, "Invalid Tweet Data: "+err.Error(), 400)
		return
	}

	doc := models.TweetBD{
		UserID:  UserID,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := bd.InsertTweet(doc)
	if err != nil {
		http.Error(w, "An error occurred while trying to insert the tweet, please try again: "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "Unable to insert the tweet", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
