package request

import "backend_capstone/services/transaction/dto"

type MidtransReq struct {
	Code          string `json:"status_code,omitempty"`
	TranasctionId string `json:"order_id,omitempty"`
	Status        string `json:"transaction_status,omitempty"`
}

func (req *MidtransReq) DtoReq() dto.MidtransAfterPayment {
	return dto.MidtransAfterPayment{
		TransactionId: req.TranasctionId,
		Status:        req.Status,
		Code:          req.Code,
	}
}
