package response

type BasicMethodResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}
