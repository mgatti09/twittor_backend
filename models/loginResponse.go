package models

/*LoginResponse contiene el token que se obtiene con el Login */
type LoginResponse struct {
	Token string `json:"token,omitempty"`
}
