package bd

import (
	"context"
	"log"
	"time"

	"github.com/mgatti09/twittor_backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*SearchProfile busca un perfil en la BD */
func SearchProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoCN.Database("twittor")
	coll := db.Collection("users")

	var profile models.User

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	//Aca ya se tiene la información del perfil
	err := coll.FindOne(ctx, condition).Decode(&profile)
	if err != nil {
		log.Println("User not found: " + err.Error())
		return profile, err
	}
	//Se limpia el password y como tiene un omit no lo devolvera en los resultados
	profile.Password = ""

	return profile, nil
}
