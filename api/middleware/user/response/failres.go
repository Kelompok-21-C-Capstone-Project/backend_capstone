package response

type FailResponse struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}
