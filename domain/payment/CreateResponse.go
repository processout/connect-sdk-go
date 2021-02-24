// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

//go:generate go run ../gen-accessors.go

// CreateResponse represents class CreatePaymentResponse
type CreateResponse struct {
	CreationOutput *CreationOutput `json:"creationOutput,omitempty"`
	MerchantAction *MerchantAction `json:"merchantAction,omitempty"`
	Payment        *Payment        `json:"payment,omitempty"`
}

// NewCreateResponse constructs a new CreateResponse
func NewCreateResponse() *CreateResponse {
	return &CreateResponse{}
}
