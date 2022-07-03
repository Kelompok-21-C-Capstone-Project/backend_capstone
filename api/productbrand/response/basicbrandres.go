package response

type BasicBrandResponse struct {
	Status  string `json:"status" example:"fail"`
	Message string `json:"message,omitempty" example:"error message"`
}

type BasicBrandSuccessResponse struct {
	Status string `json:"status,omitempty" example:"success"`
}
