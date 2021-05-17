package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*User entidad de usuario en MongoDB */
type User struct {
	// Field appears in JSON as key "_id" and the field is omitted from the object if its value is empty.
	// https://golang.org/pkg/encoding/json/
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Surname   string             `bson:"surname" json:"surname,omitempty"`
	Lastname  string             `bson:"lastname" json:"lastname,omitempty"`
	Birthdate time.Time          `bson:"birthdate" json:"birthdate,omitempty"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar,omitempty"`
	Banner    string             `bson:"banner" json:"banner,omitempty"`
	Biography string             `bson:"biography" json:"biography,omitempty"`
	Location  string             `bson:"location" json:"location,omitempty"`
	Website   string             `bson:"website" json:"website,omitempty"`
}
