// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

//go:generate go run ../gen-accessors.go

// CancelApprovalPaymentResponse represents class CancelApprovalPaymentResponse
type CancelApprovalPaymentResponse struct {
	Payment *Payment `json:"payment,omitempty"`
}

// NewCancelApprovalPaymentResponse constructs a new CancelApprovalPaymentResponse
func NewCancelApprovalPaymentResponse() *CancelApprovalPaymentResponse {
	return &CancelApprovalPaymentResponse{}
}
