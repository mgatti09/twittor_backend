package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mgatti09/twittor_backend/bd"
	"github.com/mgatti09/twittor_backend/models"
)

/*Email valor del email usado en todos los endpoints*/
var Email string

/*UserID es el ID devuelto del modelo que se usará en todos los endpoints */
var UserID string

/*ProcessToken procesamiento del token para extraer sus valores */
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	//Para poder decodificar el token
	myPwd := []byte("MiClaveUltraSECRETA")

	//jwt exige que sea un puntero
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("token format invalid")
	}

	tk = strings.TrimSpace(splitToken[1])

	//Sintaxis para verificar si el token es valido y mapear el token dentro de claims
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myPwd, nil
	})
	if err == nil {
		//Si el token es válido lo primero que validamos es si el email existe en la BD
		_, userFound, _ := bd.UserExists(claims.Email)
		if userFound {
			Email = claims.Email
			UserID = claims.ID.Hex()
		}
		return claims, userFound, UserID, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}

	return claims, false, string(""), err
}
