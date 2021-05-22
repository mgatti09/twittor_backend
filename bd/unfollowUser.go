package bd

import (
	"context"
	"time"

	"github.com/mgatti09/twittor_backend/models"
)

/*UnfollowUser elimina la relación en la BD */
func UnfollowUser(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	coll := db.Collection("relations")

	_, err := coll.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}
