package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/mgatti09/twittor_backend/bd"
)

/*GetBanner envia la imagen del Banner al HTTP */
func GetBanner(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "id key is required", http.StatusUnprocessableEntity)
		return
	}

	profile, err := bd.SearchProfile(ID)
	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	openFile, err := os.Open("uploads/banners/" + profile.Banner)
	if err != nil {
		http.Error(w, "Banner image not found", http.StatusBadRequest)
		return
	}

	//Se copia la imagen al response writter en modo binario
	_, err = io.Copy(w, openFile)
	if err != nil {
		http.Error(w, "Error copying the Banner image", http.StatusBadRequest)
		return
	}

}
