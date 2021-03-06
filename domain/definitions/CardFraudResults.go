// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package definitions

//go:generate go run ../gen-accessors.go

// CardFraudResults represents class CardFraudResults
type CardFraudResults struct {
	AvsResult          *string                      `json:"avsResult,omitempty"`
	CvvResult          *string                      `json:"cvvResult,omitempty"`
	FraudServiceResult *string                      `json:"fraudServiceResult,omitempty"`
	Fraugster          *FraugsterResults            `json:"fraugster,omitempty"`
	InAuth             *InAuth                      `json:"inAuth,omitempty"`
	RetailDecisions    *FraudResultsRetailDecisions `json:"retailDecisions,omitempty"`
}

// NewCardFraudResults constructs a new CardFraudResults
func NewCardFraudResults() *CardFraudResults {
	return &CardFraudResults{}
}
