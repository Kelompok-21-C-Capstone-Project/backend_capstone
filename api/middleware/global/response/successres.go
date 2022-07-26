package response

type SuccessResponse struct {
	Status string `json:"status,omitempty"`
	Token  string `json:"token,omitempty"`
}
