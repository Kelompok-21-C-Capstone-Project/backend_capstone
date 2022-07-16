package response

type PaymentMethods struct {
	Id       int              `json:"id" example:"1"`
	Type     string           `json:"type" example:"Virtual Account"`
	Services []PaymentService `json:"services"`
}

type PaymentService struct {
	Id    int    `json:"id" example:"1"`
	Label string `json:"label" example:"Mandiri Virtual Account"`
	Icon  string `json:"icon,omitempty" example:"mdi-va-mandiri"`
}
