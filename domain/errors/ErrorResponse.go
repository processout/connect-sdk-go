// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package errors

//go:generate go run ../gen-accessors.go

// ErrorResponse represents class ErrorResponse
type ErrorResponse struct {
	ErrorID *string     `json:"errorId,omitempty"`
	Errors  *[]APIError `json:"errors,omitempty"`
}

// NewErrorResponse constructs a new ErrorResponse
func NewErrorResponse() *ErrorResponse {
	return &ErrorResponse{}
}
