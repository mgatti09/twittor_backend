package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*TweetsIFollow estructura a devolver de los tweets de mis seguidores*/
type TweetsIFollow struct {
	ID              primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID          string             `bson:"userId" json:"userId,omitempty"`
	UserFollowingID string             `bson:"userFollowingID" json:"userFollowingID,omitempty"`
	Tweet           struct {
		Message string    `message:"userFollowingID" json:"message,omitempty"`
		Date    time.Time `date:"userFollowingID" json:"date,omitempty"`
		ID      string    `_id:"userFollowingID" json:"_id,omitempty"`
	}
}
