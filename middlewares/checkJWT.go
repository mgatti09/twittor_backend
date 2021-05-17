package middlewares

import (
	"net/http"

	"github.com/mgatti09/twittor_backend/routers"
)

/* checkJWT valida el token que viene desde la petición*/
func CheckJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Procesamiento del Key Authorization que vendra en el Header de las peticiones
		_, _, _, err := routers.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error at Token!: "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
