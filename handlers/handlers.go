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

	router.HandleFunc("/singup", middlewares.CheckBD(routers.SignUp)).Methods("POST")
	router.HandleFunc("/login", middlewares.CheckBD(routers.Login)).Methods("POST")
	router.HandleFunc("/viewprofile", middlewares.CheckBD(middlewares.CheckJWT(routers.ViewProfile))).Methods("GET")
	router.HandleFunc("/updateprofile", middlewares.CheckBD(middlewares.CheckJWT(routers.UpdateProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middlewares.CheckBD(middlewares.CheckJWT(routers.InsertTweet))).Methods("POST")
	router.HandleFunc("/tweets", middlewares.CheckBD(middlewares.CheckJWT(routers.GetTweets))).Methods("GET")
	router.HandleFunc("/tweet", middlewares.CheckBD(middlewares.CheckJWT(routers.DeleteTweet))).Methods("DELETE")

	router.HandleFunc("/avatar", middlewares.CheckBD(middlewares.CheckJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/avatar", middlewares.CheckBD(routers.GetAvatar)).Methods("GET")
	router.HandleFunc("/banner", middlewares.CheckBD(middlewares.CheckJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/banner", middlewares.CheckBD(routers.GetBanner)).Methods("GET")

	router.HandleFunc("/follow", middlewares.CheckBD(middlewares.CheckJWT(routers.FollowUser))).Methods("POST")
	router.HandleFunc("/follow", middlewares.CheckBD(middlewares.CheckJWT(routers.UnfollowUser))).Methods("DELETE")
	router.HandleFunc("/follow", middlewares.CheckBD(middlewares.CheckJWT(routers.FollowingUser))).Methods("GET")

	router.HandleFunc("/listUsers", middlewares.CheckBD(middlewares.CheckJWT(routers.GetListUsers))).Methods("GET")

	router.HandleFunc("/tweetsIFollow", middlewares.CheckBD(middlewares.CheckJWT(routers.GetTweetsIFollow))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
