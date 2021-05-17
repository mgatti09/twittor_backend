package handlers

/*
Función Handlers que sera la que se ejecuta cuando se llama a la API.
Acá se van a definir las rutas
*/

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mgatti09/twittor_backend/middlewares"
	"github.com/mgatti09/twittor_backend/routers"
	"github.com/rs/cors" //permisos que le doy a mi API para que sea accesible desde cualquier lugar o limitar el acceso
)

/*Handlers() seteo del puerto, el handler y pongo a escuchar al servidor */
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlewares.CheckBD(routers.SignUp)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8081"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
