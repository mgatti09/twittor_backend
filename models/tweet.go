package models

/* Tweet captura del Body, el mensaje que nos llega desde la petición*/
type Tweet struct {
	Message string `bson:"message" json:"message"`
}
