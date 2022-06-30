package response

type BasicProductResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}
