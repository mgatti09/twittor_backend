package main

import (
	"log"

	"github.com/mgatti09/twittor_backend/bd"
	"github.com/mgatti09/twittor_backend/handlers"
)

func main() {
	if !bd.CheckConn() {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Handlers()
}
