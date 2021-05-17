package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mgatti09/twittor_backend/models"
)

func GetToken(t models.User) (string, error) {
	myPwd := []byte("MiClaveUltraSECRETA")

	//Lista de privilegios que se guarda en el payload. Se arma manual porque no se debe guardar la
	//password en el JWT ya que es muy fácil decodificarlo por ejemplo en la página jwt.io
	payload := jwt.MapClaims{
		"email":     t.Email,
		"surname":   t.Surname,
		"lastname":  t.Lastname,
		"birthdate": t.Birthdate,
		"biography": t.Biography,
		"location":  t.Location,
		"website":   t.Website,
		"_id":       t.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myPwd)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil

}
