package response

type BasicVendorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}
