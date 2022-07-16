package response

type BasicProductResponse struct {
	Status  string `json:"status" example:"fail"`
	Message string `json:"message,omitempty" example:"error message"`
}

type BasicProductSuccessResponse struct {
	Status string `json:"status,omitempty" example:"success"`
}
