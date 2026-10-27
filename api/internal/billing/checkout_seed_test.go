// api/internal/billing/checkout_seed_test.go
package billing

import (
	"testing"
	"time"
)

func (h *harness) seedGuildTrial(t *testing.T, guildID int64) {
	t.Helper()
	if _, err := h.pool.Exec(t.Context(),
		`insert into guild_trials (guild_id, stripe_subscription_id) values ($1, 'sub_trial')`, guildID); err != nil {
		t.Fatal(err)
	}
}

func (h *harness) seedTransferOffer(t *testing.T, guildID, fromUserID, toUserID int64) {
	t.Helper()
	if _, err := h.pool.Exec(t.Context(),
		`insert into plan_transfers (guild_id, from_user_id, to_user_id, from_subscription_id, state)
		 values ($1, $2, $3, 'sub_old', 'offered')`, guildID, fromUserID, toUserID); err != nil {
		t.Fatal(err)
	}
}

// seedStripeGuildPlan inserts a live Stripe guild plan on lookupKey whose
// period ends at periodEnd.
func (h *harness) seedStripeGuildPlan(t *testing.T, guildID int64, lookupKey string, periodEnd time.Time) {
	t.Helper()
	if _, err := h.pool.Exec(t.Context(),
		`insert into entitlements (guild_id, plan, tier, source, status, stripe_subscription_id, price_lookup_key, current_period_end)
		 values ($1, 'guild', 'guild', 'stripe', 'active', 'sub_old', $2, $3)`, guildID, lookupKey, periodEnd); err != nil {
		t.Fatal(err)
	}
}

func (h *harness) seedComeback(t *testing.T, userID int64, lookupKey string, expiresAt time.Time) int64 {
	t.Helper()
	var id int64
	if err := h.pool.QueryRow(t.Context(),
		`insert into seat_comebacks (user_id, guild_name, ended_at, price_lookup_key, expires_at)
		 values ($1, 'Old Guild', now() - interval '1 day', $2, $3) returning id`,
		userID, lookupKey, expiresAt).Scan(&id); err != nil {
		t.Fatal(err)
	}
	return id
}

// officerOf seeds a verified leader of a fresh claimed guild.
func (h *harness) officerOf(t *testing.T, name string) (guildID, officerID int64) {
	t.Helper()
	guildID = h.seedGuild(t, name, true)
	officerID = h.seedUser(t, name+"-officer@example.com")
	h.seedGuildMember(t, guildID, officerID, "leader")
	return guildID, officerID
}
