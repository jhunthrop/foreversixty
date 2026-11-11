package billing

import (
	"strings"
	"testing"
	"time"
)

func TestCheckoutSubscriptionData(t *testing.T) {
	end := time.Date(2026, 12, 1, 0, 0, 0, 0, time.UTC)
	cases := []struct {
		name     string
		params   CheckoutParams
		wantDays *int64
		wantEnd  *int64
		wantErr  bool
	}{
		{name: "no trial"},
		{name: "days", params: CheckoutParams{TrialPeriodDays: 7}, wantDays: ptr(int64(7))},
		{name: "end", params: CheckoutParams{TrialEnd: &end}, wantEnd: ptr(end.Unix())},
		{name: "both is an error", params: CheckoutParams{TrialPeriodDays: 7, TrialEnd: &end}, wantErr: true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.params.Metadata = map[string]string{"plan": "guild"}
			data, err := checkoutSubscriptionData(c.params)
			if (err != nil) != c.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, c.wantErr)
			}
			if c.wantErr {
				return
			}
			if data.Metadata["plan"] != "guild" {
				t.Errorf("metadata = %v", data.Metadata)
			}
			if !equalInt64Ptr(data.TrialPeriodDays, c.wantDays) || !equalInt64Ptr(data.TrialEnd, c.wantEnd) {
				t.Errorf("trial = %v / %v, want %v / %v", data.TrialPeriodDays, data.TrialEnd, c.wantDays, c.wantEnd)
			}
		})
	}
}

func ptr[T any](v T) *T { return &v }

func equalInt64Ptr(a, b *int64) bool {
	if a == nil || b == nil {
		return a == b
	}
	return *a == *b
}

// TestCustomerCreateParamsCarryAPerAccountIdempotencyKey is issue #3 L4: two checkouts racing
// for an account with no Customer yet (two tabs, two guilds) get one Customer back from Stripe,
// not two.
func TestCustomerCreateParamsCarryAPerAccountIdempotencyKey(t *testing.T) {
	a := customerCreateParams(42, "a@example.com")
	if a.IdempotencyKey == nil || !strings.HasPrefix(*a.IdempotencyKey, "fs-customer-42-") {
		t.Fatalf("idempotency key = %v, want fs-customer-42-<email digest>", a.IdempotencyKey)
	}
	if again := customerCreateParams(42, "a@example.com"); *again.IdempotencyKey != *a.IdempotencyKey {
		t.Fatalf("keys differ for the same account and email: %s, %s", *again.IdempotencyKey, *a.IdempotencyKey)
	}
	if moved := customerCreateParams(42, "b@example.com"); *moved.IdempotencyKey == *a.IdempotencyKey {
		t.Fatal("a changed email reuses the key Stripe would refuse")
	}
	if a.Email == nil || *a.Email != "a@example.com" {
		t.Fatalf("email = %v", a.Email)
	}
	if b := customerCreateParams(43, ""); b.Email != nil || *b.IdempotencyKey == *a.IdempotencyKey {
		t.Fatalf("params = %+v, want no email and a key of its own", b)
	}
}
