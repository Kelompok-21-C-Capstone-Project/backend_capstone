package response

type BasicUserResponse struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}
