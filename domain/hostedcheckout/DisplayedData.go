// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package hostedcheckout

import "github.com/processout/connect-sdk-go/domain/definitions"

// DisplayedData represents class DisplayedData
type DisplayedData struct {
	DisplayedDataType *string                     `json:"displayedDataType,omitempty"`
	RenderingData     *string                     `json:"renderingData,omitempty"`
	ShowData          *[]definitions.KeyValuePair `json:"showData,omitempty"`
}

// NewDisplayedData constructs a new DisplayedData
func NewDisplayedData() *DisplayedData {
	return &DisplayedData{}
}
