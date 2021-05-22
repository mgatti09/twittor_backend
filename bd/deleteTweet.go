package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*DeleteTweet función que recibe el ID del tweet y del usuario que lo publico para borrarlo */
func DeleteTweet(ID string, UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	coll := db.Collection("tweets")

	objId, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id":    objId,
		"userid": UserID,
	}

	_, err := coll.DeleteOne(ctx, condition)

	return err
}
