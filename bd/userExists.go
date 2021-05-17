package bd

import (
	"context"
	"time"

	"github.com/mgatti09/twittor_backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* UserExists recibe por parametro un email y busca si el usuario se encuentra registrado en la BD
Si existe retorna la información del usuario y la devuelve en el primer valor de retorno y coloca en true
el segundo valor de retorno */
func UserExists(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	coll := db.Collection("users")

	//Condición para realizar la búsqueda en Mongo, para eso creamos el "JSON" con M
	condition := bson.M{"email": email}

	var result models.User
	err := coll.FindOne(ctx, condition).Decode(&result)

	//en result se guardara el ObjectId pero tratar ese tipo de dato es complejo, por ello se declara
	//variable ID para devolver en el tercer parámetro este valor. Se convierte a hexadecimal de string
	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}
