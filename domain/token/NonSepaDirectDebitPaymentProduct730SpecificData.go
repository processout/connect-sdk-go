// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package token

import "github.com/processout/connect-sdk-go/domain/definitions"

// NonSepaDirectDebitPaymentProduct730SpecificData represents class TokenNonSepaDirectDebitPaymentProduct730SpecificData
type NonSepaDirectDebitPaymentProduct730SpecificData struct {
	BankAccountBban *definitions.BankAccountBban `json:"bankAccountBban,omitempty"`
}

// NewNonSepaDirectDebitPaymentProduct730SpecificData constructs a new NonSepaDirectDebitPaymentProduct730SpecificData
func NewNonSepaDirectDebitPaymentProduct730SpecificData() *NonSepaDirectDebitPaymentProduct730SpecificData {
	return &NonSepaDirectDebitPaymentProduct730SpecificData{}
}
