// api/internal/billing/trial_harness_test.go
package billing

import (
	"bytes"
	"fmt"
	"log/slog"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
	"github.com/jhunthrop/foreversixty/api/internal/mail"
)

// trialHarness is the notice harness with captured logs and in-app events, for the trial
// reminder, cancel link and receipt tests.
type trialHarness struct {
	*noticeHarness
	logs   *bytes.Buffer
	events *eventRecorder
}

func newTrialHarness(t *testing.T) *trialHarness {
	t.Helper()
	h := &trialHarness{noticeHarness: newNoticeHarness(t), logs: &bytes.Buffer{}, events: &eventRecorder{}}
	h.svc.Log = slog.New(slog.NewTextHandler(h.logs, nil))
	h.svc.Events = h.events
	h.svc.RenewalEmails = false
	h.gateway.CustomerEmails = map[string]string{}
	return h
}

// seededTrial is one trialing subscription the harness wrote.
type seededTrial struct {
	entitlementID  int64
	subscriptionID string
	end            time.Time
}

// seedTrial writes a trialing Stripe row from start to end and its trial record (premium_trials,
// or guild_trials for a guild); guildID 0 is Premium.
func (h *trialHarness) seedTrial(t *testing.T, guildID, payer int64, lookupKey string, start, end time.Time) seededTrial {
	t.Helper()
	price, _ := catalog.Lookup(lookupKey)
	tr := seededTrial{subscriptionID: fmt.Sprintf("sub_trial_%d_%d", payer, guildID), end: end}
	var err error
	if guildID == 0 {
		err = h.pool.QueryRow(t.Context(),
			`insert into entitlements (user_id, plan, source, status, billing_user_id, stripe_subscription_id,
			   price_lookup_key, unit_amount_cents, current_period_end, trial_end)
			 values ($1, 'premium', 'stripe', 'trialing', $1, $2, $3, $4, $5, $5) returning id`,
			payer, tr.subscriptionID, lookupKey, price.AmountCents, end).Scan(&tr.entitlementID)
		if err == nil {
			_, err = h.pool.Exec(t.Context(), `insert into premium_trials (user_id, stripe_subscription_id, created_at)
				values ($1, $2, $3)`, payer, tr.subscriptionID, start)
		}
	} else {
		err = h.pool.QueryRow(t.Context(),
			`insert into entitlements (guild_id, plan, tier, source, status, billing_user_id, stripe_subscription_id,
			   price_lookup_key, unit_amount_cents, current_period_end, trial_end)
			 values ($1, 'guild', $2, 'stripe', 'trialing', $3, $4, $5, $6, $7, $7) returning id`,
			guildID, price.Offer.Tier(), payer, tr.subscriptionID, lookupKey, price.AmountCents, end).Scan(&tr.entitlementID)
		if err == nil {
			_, err = h.pool.Exec(t.Context(), `insert into guild_trials (guild_id, stripe_subscription_id, started_at)
				values ($1, $2, $3)`, guildID, tr.subscriptionID, start)
		}
	}
	if err != nil {
		t.Fatal(err)
	}
	return tr
}

// seedCustomer gives the payer a Stripe customer with email ("" for a customer without one).
func (h *trialHarness) seedCustomer(t *testing.T, userID int64, email string) string {
	t.Helper()
	id := fmt.Sprintf("cus_trial_%d", userID)
	if err := h.svc.Store.SaveCustomerID(t.Context(), userID, id); err != nil {
		t.Fatal(err)
	}
	h.gateway.CustomerEmails[id] = email
	return id
}

// sendAt runs the notices at now and answers what was sent and published this run.
func (h *trialHarness) sendAt(t *testing.T, now time.Time) ([]mail.Message, int) {
	t.Helper()
	mails, published := len(h.mailer.Sent), len(h.events.got)
	if _, err := h.svc.SendNotices(t.Context(), now); err != nil {
		t.Fatal(err)
	}
	return h.mailer.Sent[mails:], len(h.events.got) - published
}

// tokenOf reads the cancel token out of a reminder's link.
func tokenOf(t *testing.T, m mail.Message) string {
	t.Helper()
	const marker = "/premium/cancel?t="
	i := strings.Index(m.Text, marker)
	if i < 0 {
		t.Fatalf("no cancel link in %q", m.Text)
	}
	rest := m.Text[i+len(marker):]
	if j := strings.IndexAny(rest, " \n"); j >= 0 {
		rest = rest[:j]
	}
	return rest
}
