// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

//go:generate go run ../gen-accessors.go

// Response represents class PaymentResponse
type Response struct {
	HostedCheckoutSpecificOutput *HostedCheckoutSpecificOutput `json:"hostedCheckoutSpecificOutput,omitempty"`
	ID                           *string                       `json:"id,omitempty"`
	PaymentOutput                *Output                       `json:"paymentOutput,omitempty"`
	Status                       *string                       `json:"status,omitempty"`
	StatusOutput                 *StatusOutput                 `json:"statusOutput,omitempty"`
}

// NewResponse constructs a new Response
func NewResponse() *Response {
	return &Response{}
}
