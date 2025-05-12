package dto

type LoginBody struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}
