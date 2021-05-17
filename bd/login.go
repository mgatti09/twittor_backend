package bd

import (
	"github.com/mgatti09/twittor_backend/models"
	"golang.org/x/crypto/bcrypt"
)

/* Login realiza el logueo. Si es exitoso retorna el User y True, de lo contrario False de logueo no exitoso*/
func Login(email string, pwd string) (models.User, bool) {
	usu, userFound, _ := UserExists(email)
	if !userFound {
		return usu, false
	}

	pwdBytes := []byte(pwd)
	pwdBD := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(pwdBD, pwdBytes)
	if err != nil {
		return usu, false
	}

	return usu, true

}
