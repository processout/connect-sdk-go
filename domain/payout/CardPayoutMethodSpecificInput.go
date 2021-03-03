// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payout

import "github.com/processout/connect-sdk-go/domain/definitions"

// CardPayoutMethodSpecificInput represents class CardPayoutMethodSpecificInput
type CardPayoutMethodSpecificInput struct {
	Card             *definitions.Card `json:"card,omitempty"`
	PaymentProductID *int32            `json:"paymentProductId,omitempty"`
	Token            *string           `json:"token,omitempty"`
}

// NewCardPayoutMethodSpecificInput constructs a new CardPayoutMethodSpecificInput
func NewCardPayoutMethodSpecificInput() *CardPayoutMethodSpecificInput {
	return &CardPayoutMethodSpecificInput{}
}
