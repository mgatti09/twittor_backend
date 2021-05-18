package bd

import (
	"context"
	"time"

	"github.com/mgatti09/twittor_backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateUser(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	coll := db.Collection("users")

	//make permite crear slices o mapas
	record := make(map[string]interface{})
	if len(u.Surname) > 0 {
		record["surname"] = u.Surname
	}
	if len(u.Lastname) > 0 {
		record["lastname"] = u.Lastname
	}
	record["birthdate"] = u.Birthdate
	if len(u.Avatar) > 0 {
		record["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		record["banner"] = u.Banner
	}
	if len(u.Biography) > 0 {
		record["biography"] = u.Biography
	}
	if len(u.Location) > 0 {
		record["location"] = u.Location
	}
	if len(u.Website) > 0 {
		record["website"] = u.Website
	}

	//Armando el registro de actualización. "$set" instrucción de mongo para actualizar
	updtString := bson.M{
		"$set": record,
	}

	//Esto transforma el ID en un object ID
	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{"_id": bson.M{"$eq": objID}}

	//Ahora se realiza el UPDATE
	_, err := coll.UpdateOne(ctx, condition, updtString)
	if err != nil {
		return false, err
	}
	return true, nil
}
