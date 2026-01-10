package dto

type RegisterLoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenDto struct {
	Token string `json:"token"`
}
