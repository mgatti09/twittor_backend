package bd

import (
	"context"
	"time"

	"github.com/mgatti09/twittor_backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTweetsIFollow(ID string, page int64) ([]models.TweetsIFollow, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	coll := db.Collection("relations")

	var tweetsPerPage int64 = 20
	skip := (page - 1) * tweetsPerPage

	//Creando un slice (array) de condiciones
	conditions := make([]bson.M, 0)
	//A continuación se hacen los append de condiciones con el Aggregate de Mongo
	//$match: busca el usuario ID en la relación
	conditions = append(conditions, bson.M{"$match": bson.M{"userId": ID}})
	//Filtrada la tabla relations por el userId procedo a hacer un join con la tabla tweets.
	//$lookup: permite unir dos tablas. Hace uso de 4 parametros para poder unir tablas
	//from: se indica la tabla con la cual se hará el join, en este caso es tweets
	//localField: campo por el cual se hará el join, en este caso es UserFollowingID
	//foreingField: campo de la tabla a unir por el cual se hará el join, en este caso es el userId de tweets
	//as: Alias de como llamalos la tabla a unir, en este caso lo mantenemos tweets
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "tweets",
			"localField":   "UserFollowingID",
			"foreingField": "userid",
			"as":           "tweets",
		}})
	//$unwind para evitar que la info de la union venga en maestro detalle
	conditions = append(conditions, bson.M{"$unwind": "$tweets"})
	//$sort ordenando por campo fecha descendente
	conditions = append(conditions, bson.M{"$sort": bson.M{"tweets.date": -1}})
	//$skip ignorando registro de acuerdo a la variable  skip
	conditions = append(conditions, bson.M{"$skip": skip})
	//$limit limitando a 20 tweets por pagina
	conditions = append(conditions, bson.M{"$limit": tweetsPerPage})

	cursor, _ := coll.Aggregate(ctx, conditions)
	var result []models.TweetsIFollow

	//Ejecutando todo el cursor para que arme todo el documento. esto se hace con All
	//basando en la estructura TweetsIFollow
	err := cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true
}
