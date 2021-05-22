package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/mgatti09/twittor_backend/bd"
	"github.com/mgatti09/twittor_backend/models"
)

func UploadBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("banner")
	if err != nil {
		http.Error(w, "An error occurred with the avatar image file: "+err.Error(), http.StatusBadRequest)
		return
	}
	var extension = strings.Split(handler.Filename, ".")[1]

	var fileName string = "uploads/banners/" + UserID + "." + extension

	//Se abre un archivo nuevo con un nombre unico para el banner del usuario
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "An error occurred uploading the banner image: "+err.Error(), http.StatusBadRequest)
		return
	}

	//Copiamos la data del archivo que viene del request en el archivo nuevo
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "An error occurred copying the banner image: "+err.Error(), http.StatusBadRequest)
		return
	}

	var userBanner models.User

	userBanner.Banner = UserID + "." + extension
	status, err := bd.UpdateUser(userBanner, UserID)
	if err != nil || !status {
		http.Error(w, "An error occurred storing the banner in the BD: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
