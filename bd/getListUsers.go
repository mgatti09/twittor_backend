package bd

import (
	"context"
	"log"
	"time"

	"github.com/mgatti09/twittor_backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetListUsers(ID string, page int64, search string, userType string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	coll := db.Collection("users")

	var results []*models.User
	var usersPerPage int64 = 20

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * usersPerPage)
	findOptions.SetLimit(usersPerPage)

	//regex (?i) para no tormar en cuenta si es mayus o min
	query := bson.M{
		"surname": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := coll.Find(ctx, query, findOptions)
	if err != nil {
		log.Println("coll.Find " + err.Error())
		return results, false
	}

	for cur.Next(ctx) {
		var s models.User
		err := cur.Decode(&s)
		if err != nil {
			log.Println("cur.Decode(&s) " + err.Error())
			return results, false
		}
		var r models.Relation

		//Construyendo el objeto relacion para verificar si sigo al usuario
		r.UserID = ID
		r.UserFollowingID = s.ID.Hex()

		include := false
		following, _ := FollowingUser(r)

		//Para obtener listado de los usuarios que no sigo
		if userType == "new" && !following {
			include = true
		}

		//Para obtener listado de los usuarios que sigo
		if userType == "follow" && following {
			include = true
		}

		// Para evitar si me sigo a mi mismo
		if r.UserFollowingID == ID {
			include = false
		}

		if include {
			s.Password = ""
			s.Biography = ""
			s.Website = ""
			s.Location = ""
			s.Banner = ""
			s.Email = ""

			results = append(results, &s)
		}
	}

	err = cur.Err()
	if err != nil {
		log.Println("cur.Err() " + err.Error())
		return results, false
	}
	cur.Close(ctx)
	return results, true
}
