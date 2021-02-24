// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package dispute

//go:generate go run ../gen-accessors.go

// Response represents class DisputeResponse
type Response struct {
	DisputeOutput *Output       `json:"disputeOutput,omitempty"`
	ID            *string       `json:"id,omitempty"`
	PaymentID     *string       `json:"paymentId,omitempty"`
	Status        *string       `json:"status,omitempty"`
	StatusOutput  *StatusOutput `json:"statusOutput,omitempty"`
}

// NewResponse constructs a new Response
func NewResponse() *Response {
	return &Response{}
}
