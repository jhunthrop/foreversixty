package creators

import (
	"os"
	"strings"
	"testing"
)

// The payouts key has platform-account permissions only (Balance read; Accounts, Login Links and
// Transfers write) and none in connected accounts, so the gateway must never act on a connected
// account's behalf: no Stripe-Account header, which stripe-go sets from Params.StripeAccount.
func TestTheGatewayNeverActsAsAConnectedAccount(t *testing.T) {
	src, err := os.ReadFile("stripe_gateway.go")
	if err != nil {
		t.Fatal(err)
	}
	for _, banned := range []string{"StripeAccount", "SetStripeAccount", "Stripe-Account"} {
		if strings.Contains(string(src), banned) {
			t.Errorf("stripe_gateway.go uses %s: a call on a connected account's behalf needs permissions the payouts key does not have", banned)
		}
	}
}
