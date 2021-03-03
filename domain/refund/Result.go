// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package refund

import (
	"github.com/processout/connect-sdk-go/domain/definitions"
	"github.com/processout/connect-sdk-go/domain/payment"
)

//go:generate go run ../gen-accessors.go

// Result represents class RefundResult
type Result struct {
	ID           *string                        `json:"id,omitempty"`
	RefundOutput *payment.RefundOutput          `json:"refundOutput,omitempty"`
	Status       *string                        `json:"status,omitempty"`
	StatusOutput *definitions.OrderStatusOutput `json:"statusOutput,omitempty"`
}

// NewResult constructs a new Result
func NewResult() *Result {
	return &Result{}
}
