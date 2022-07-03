package response

type BasicUserResponse struct {
	Status  string `json:"status,omitempty" example:"fail"`
	Message string `json:"message,omitempty" example:"error message"`
}

type BasicUserSuccessResponse struct {
	Status string `json:"status,omitempty" example:"success"`
}

type SuccessLoginResponse struct {
	Status string `json:"status,omitempty" example:"success"`
	Token  string `json:"token,omitempty" example:"xxxxxxxxxxxxxxxxxxxxxxxxxxxxx.eyJjcmVhdGVkX2F0IjoiMjAyMi0wNy0wxxxxxxxxxxxxxxxxxxxxxxxxxxxxF9hdCI6IjIwMjItMDctMDNUMTA6MjA6NDYuMzAzNDA1OSswNzowMCIsImlkIjoiYWZjNxxxxxxxxxxxxxxxxxxxxxxxxxxxxbmFtZSI6Iml6YXFpIiwicm9xxxxxxxxxxxxxxxxxxxxxxxxxxxxJ9.-VKi0DWLKT1SxxxxxxxxxxxxxY2UbxbtFOj7cPA"`
}
