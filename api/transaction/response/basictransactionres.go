package response

type BasicTransactionResponse struct {
	Status  string `json:"status" example:"fail"`
	Message string `json:"message,omitempty" example:"error message"`
}

type BasicTransactionSuccessResponse struct {
	Status string `json:"status,omitempty" example:"success"`
}
