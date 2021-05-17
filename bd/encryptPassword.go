package bd

import "golang.org/x/crypto/bcrypt"

/* EncryptPassword rutina para encriptar el password*/
func EncryptPassword(pwd string) (string, error) {
	//cost es la cantidad de pasadas para encriptar a mayor costo mayor seguridad pero mayor tiempo de procesamiento
	//En este caso hace 2^8 encriptaciones
	cost := 8
	//[]byte(pwd) transformando el string a un array de byte
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), cost)

	return string(bytes), err
}
