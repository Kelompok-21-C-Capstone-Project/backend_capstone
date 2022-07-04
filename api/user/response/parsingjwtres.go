package response

type JwtPayload struct {
	Id       string   `json:"id"`
	Username string   `json:"username"`
	Name     string   `json:"name"`
	Role     []string `json:"role"`
	Email    string   `json:"email"`
	Phone    string   `json:"phone"`
}
