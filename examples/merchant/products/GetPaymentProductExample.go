// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package examples

import (
	"fmt"

	"github.com/Ingenico-ePayments/connect-sdk-go/merchant/products"
)

func getPaymentProductExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper functions newBool, newInt64 and newString to overcome this issue.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.
	// The code for these helper functions can be found in file Helper.go

	var query products.GetParams
	query.Amount = newInt64(1000)
	query.IsRecurring = newBool(true)
	query.CountryCode = newString("US")
	query.Locale = newString("en_US")
	query.CurrencyCode = newString("USD")
	query.AddHide("fields")

	response, err := client.Merchant("merchantId").Products().Get(1, query, nil)

	fmt.Println(response, err)
}
