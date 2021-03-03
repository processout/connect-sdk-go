// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package dispute

import "github.com/processout/connect-sdk-go/domain/definitions"

// CreateRequest represents class CreateDisputeRequest
type CreateRequest struct {
	AmountOfMoney  *definitions.AmountOfMoney `json:"amountOfMoney,omitempty"`
	ContactPerson  *string                    `json:"contactPerson,omitempty"`
	EmailAddress   *string                    `json:"emailAddress,omitempty"`
	ReplyTo        *string                    `json:"replyTo,omitempty"`
	RequestMessage *string                    `json:"requestMessage,omitempty"`
}

// NewCreateRequest constructs a new CreateRequest
func NewCreateRequest() *CreateRequest {
	return &CreateRequest{}
}
