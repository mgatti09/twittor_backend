package bd

import (
	"context"
	"time"

	"github.com/mgatti09/twittor_backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertTweet graba el tweet en la BD*/
func InsertTweet(t models.TweetBD) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	coll := db.Collection("tweets")

	doc := bson.M{
		"userid":  t.UserID,
		"message": t.Message,
		"date":    t.Date,
	}

	result, err := coll.InsertOne(ctx, doc)
	if err != nil {
		return string(""), false, err
	}

	//Extrae la clave del último doc insertado
	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.Hex(), true, nil
}
