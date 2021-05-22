package routers

import (
	"net/http"

	"github.com/mgatti09/twittor_backend/bd"
	"github.com/mgatti09/twittor_backend/models"
)

/*UnfollowUser realiza el borrado de la relación entre usuarios */
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followID := r.URL.Query().Get("id")
	if len(followID) < 1 {
		http.Error(w, "id key is required", http.StatusUnprocessableEntity)
		return
	}

	var t models.Relation

	t.UserID = UserID
	t.UserFollowingID = followID

	status, err := bd.UnfollowUser(t)
	if err != nil {
		http.Error(w, "An error occurred while trying to unfollow the user, please try again: "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "Unable to unfollow the user", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
