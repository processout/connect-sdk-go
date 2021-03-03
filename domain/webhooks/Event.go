package webhooks

import (
	"github.com/processout/connect-sdk-go/domain/dispute"
	"github.com/processout/connect-sdk-go/domain/payment"
	"github.com/processout/connect-sdk-go/domain/payout"
	"github.com/processout/connect-sdk-go/domain/refund"
	"github.com/processout/connect-sdk-go/domain/token"
)

//go:generate go run ../gen-accessors.go

// Event represents an event that is sent by webhooks
type Event struct {
	APIVersion *string           `json:"apiVersion,omitempty"`
	ID         *string           `json:"id,omitempty"`
	Created    *string           `json:"created,omitempty"`
	MerchantID *string           `json:"merchantId,omitempty"`
	Type       *string           `json:"type,omitempty"`
	Payment    *payment.Response `json:"payment,omitempty"`
	Refund     *refund.Response  `json:"refund,omitempty"`
	Payout     *payout.Response  `json:"payout,omitempty"`
	Token      *token.Response   `json:"token,omitempty"`
	Dispute    *dispute.Response `json:"dispute,omitempty"`
}

// NewEvent constructs a new Event
func NewEvent() (*Event, error) {
	return &Event{}, nil
}
