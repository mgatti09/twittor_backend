package bd

import (
	"context"
	"log"
	"time"

	"github.com/mgatti09/twittor_backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*GetTweets obtiene los tweets de la BD de un perfil */
func GetTweets(ID string, page int64) ([]*models.GetTweetsBD, bool) {
	var tweetsPerPage int64 = 20

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	coll := db.Collection("tweets")

	var results []*models.GetTweetsBD

	condition := bson.M{
		"userid": ID,
	}

	// opciones para poder filtrar y darle un comportamiento a mi consulta de BD
	opt := options.Find()
	//Propiedad que define la cantidad maxima de docs a traer para poder paginar
	opt.SetLimit(tweetsPerPage)
	//Propiedad que ordena los datos por fecha de forma descendente
	opt.SetSort(bson.D{{Key: "date", Value: -1}})
	//Para la paginación
	opt.SetSkip((page - 1) * tweetsPerPage)

	cursor, err := coll.Find(ctx, condition, opt)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	for cursor.Next(context.TODO()) {
		var record models.GetTweetsBD
		err := cursor.Decode(&record)
		if err != nil {
			return results, false
		}
		results = append(results, &record)
	}

	return results, true
}
