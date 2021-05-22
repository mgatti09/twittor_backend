package bd

import (
	"context"
	"time"

	"github.com/mgatti09/twittor_backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*FollowingUser verifica si se está siguiendo al usuario */
func FollowingUser(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	coll := db.Collection("relations")

	condition := bson.M{
		"userId":          t.UserID,
		"UserFollowingID": t.UserFollowingID,
	}

	var result models.Relation

	err := coll.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return false, err
	}

	return true, nil
}
