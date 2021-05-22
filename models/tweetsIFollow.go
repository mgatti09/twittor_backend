package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*TweetsIFollow estructura a devolver de los tweets de mis seguidores*/
type TweetsIFollow struct {
	ID              primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID          string             `bson:"userId" json:"userId,omitempty"`
	UserFollowingID string             `bson:"UserFollowingID" json:"userFollowingID,omitempty"`
	Tweet           struct {
		Message string    `bson:"message" json:"message,omitempty"`
		Date    time.Time `bson:"date" json:"date,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitempty"`
	}
}
