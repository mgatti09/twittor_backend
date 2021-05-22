package routers

import (
	"encoding/json"
	"net/http"

	"github.com/mgatti09/twittor_backend/bd"
	"github.com/mgatti09/twittor_backend/models"
)

/*FollowingUser retorna el response que indica si un usuario sigue a otro */
func FollowingUser(w http.ResponseWriter, r *http.Request) {
	followID := r.URL.Query().Get("id")
	if len(followID) < 1 {
		http.Error(w, "id key is required", http.StatusUnprocessableEntity)
		return
	}

	var t models.Relation

	t.UserID = UserID
	t.UserFollowingID = followID

	var resp models.FollowingUserResponse
	resp.Status = true

	status, err := bd.FollowingUser(t)
	if err != nil || !status {
		resp.Status = false
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

}
