package response

type JwtPayload struct {
	Id       string   `json:"id" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
	Username string   `json:"username" example:"username"`
	Name     string   `json:"name" example:"some name"`
	Role     []string `json:"role" example:"admin"`
	Email    string   `json:"email" example:"somemail@mail.com"`
	Phone    string   `json:"phone" example:"08XXXXXXXX"`
	Iat      int64    `json:"iat" example:"1257894000000"`
	Exp      int64    `json:"ext" example:"1257894000000"`
}
