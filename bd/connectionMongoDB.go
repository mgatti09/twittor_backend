package bd

/* Go Concurrency Patterns: Context. https://blog.golang.org/context
The Context package defines an API which provides support for deadlines, cancelation signals,
and request-scoped values that can be passed across API boundaries and between goroutines.
This API is an essential part of any application you will write in Go.

Context es un entorno de ejecución donde se podra setear un contexto de ejecución de manera de
evitar los cuelgues. Permite para comunicar info entre ejecución, ademas permite setear valores
como timeouts; por ejemplo si llega a haber un problema en la BD no cuelga la API completa, sino
solo el contexto que realizo la petición
*/

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN objeto de conexion a la BD*/
var MongoCN = ConnBD()

/*ConnBD función que conecta a la BD  */
func ConnBD() *mongo.Client {
	uri := "mongodb+srv://user:pass@mongodb.net/test?w=majority"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Could not connect to MongoDB\n" + err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Could not Ping to MongoDB\n" + err.Error())
		return client
	}

	log.Println("Successfully connected and pinged.")
	return client
}

/*CheckConn ping a la BD para verificar si está arriba*/
func CheckConn() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	return err == nil
}
