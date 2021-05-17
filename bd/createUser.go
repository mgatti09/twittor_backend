package bd

import (
	"context"
	"time"

	"github.com/mgatti09/twittor_backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*CreateUser crea un usuario en la BD */
func CreateUser(u models.User) (string, bool, error) {

	//Contexto que controla que no tarde mas de 15 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	//Al final de la funcion se cancela el contexto WithTimeout. Esto se hace para evitar
	//dejar contextos vivos que no se están usando
	defer cancel()

	db := MongoCN.Database("twittor")
	coll := db.Collection("users")

	//Encriptando el password
	u.Password, _ = EncryptPassword(u.Password)

	result, err := coll.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
