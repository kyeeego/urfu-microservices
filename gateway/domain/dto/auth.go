package dto

type RegisterLoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenDto struct {
	Token string `json:"token"`
}

type AuthorizeDto struct {
	UserID int `json:""user_id`
}
