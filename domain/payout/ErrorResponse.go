// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payout

import "github.com/processout/connect-sdk-go/domain/errors"

// ErrorResponse represents class PayoutErrorResponse
type ErrorResponse struct {
	ErrorID      *string            `json:"errorId,omitempty"`
	Errors       *[]errors.APIError `json:"errors,omitempty"`
	PayoutResult *Result            `json:"payoutResult,omitempty"`
}

// NewErrorResponse constructs a new ErrorResponse
func NewErrorResponse() *ErrorResponse {
	return &ErrorResponse{}
}
