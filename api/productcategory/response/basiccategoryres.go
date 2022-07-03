package response

type BasicCategoryResponse struct {
	Status  string `json:"status,omitempty" example:"fail"`
	Message string `json:"message,omitempty" example:"error message"`
}

type BasicCategorySuccessResponse struct {
	Status string `json:"status,omitempty" example:"success"`
}
