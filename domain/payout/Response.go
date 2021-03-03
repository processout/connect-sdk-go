// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payout

import (
	"github.com/processout/connect-sdk-go/domain/definitions"
	"github.com/processout/connect-sdk-go/domain/payment"
)

//go:generate go run ../gen-accessors.go

// Response represents class PayoutResponse
type Response struct {
	ID           *string                        `json:"id,omitempty"`
	PayoutOutput *payment.OrderOutput           `json:"payoutOutput,omitempty"`
	Status       *string                        `json:"status,omitempty"`
	StatusOutput *definitions.OrderStatusOutput `json:"statusOutput,omitempty"`
}

// NewResponse constructs a new Response
func NewResponse() *Response {
	return &Response{}
}
