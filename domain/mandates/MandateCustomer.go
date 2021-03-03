// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package mandates

import "github.com/processout/connect-sdk-go/domain/definitions"

// MandateCustomer represents class MandateCustomer
type MandateCustomer struct {
	BankAccountIban     *definitions.BankAccountIban `json:"bankAccountIban,omitempty"`
	CompanyName         *string                      `json:"companyName,omitempty"`
	ContactDetails      *MandateContactDetails       `json:"contactDetails,omitempty"`
	MandateAddress      *MandateAddress              `json:"mandateAddress,omitempty"`
	PersonalInformation *MandatePersonalInformation  `json:"personalInformation,omitempty"`
}

// NewMandateCustomer constructs a new MandateCustomer
func NewMandateCustomer() *MandateCustomer {
	return &MandateCustomer{}
}
