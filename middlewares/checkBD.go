package middlewares

/*
Middlewares es como un pasamano, debe recibir y devolver lo mismo
*/

import (
	"net/http"

	"github.com/mgatti09/twittor_backend/bd"
)

/* CheckBD middleware que permite conocer el estado de la BD */
func CheckBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !bd.CheckConn() {
			http.Error(w, "Connection lost with the DB", 5000)
			return
		}
		next.ServeHTTP(w, r)
	}
}
