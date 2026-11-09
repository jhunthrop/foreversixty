package main

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/mail"
)

// The daily job's sweep gets the Stripe gateway as its Canceller, so a seat
// released before it starts resumes its holder's own Premium; without Stripe
// it stays nil and the sweep falls back to the comeback.
func TestBillingDaily_SeatsCancellerFollowsStripeConfig(t *testing.T) {
	configured := config.Config{StripeSecretKey: "sk_test_x", StripeWebhookSecret: "whsec_x"}
	if svc := newBillingDailyServices(configured, nil, &mail.Fake{}, slog.Default()); svc.Seats.Canceller == nil {
		t.Fatal("configured Stripe left the seats Canceller nil")
	}
	if svc := newBillingDailyServices(config.Config{}, nil, &mail.Fake{}, slog.Default()); svc.Seats.Canceller != nil {
		t.Fatalf("unconfigured Stripe set a Canceller: %#v", svc.Seats.Canceller)
	}
}

func TestBillingDaily_SendsRenewalOnce(t *testing.T) {
	if sent := dailyRenewals(t, true); sent != 1 {
		t.Fatalf("sent %d, want 1", sent)
	}
}

// OUR_RENEWAL_EMAILS is off by default: Stripe sends the upcoming-renewal email, and ours stays quiet.
func TestBillingDaily_NoRenewalEmailWhileTheSwitchIsOff(t *testing.T) {
	if sent := dailyRenewals(t, false); sent != 0 {
		t.Fatalf("sent %d with OUR_RENEWAL_EMAILS off, want 0", sent)
	}
}

// dailyRenewals runs the job twice over one yearly row two days from renewal and counts the emails.
func dailyRenewals(t *testing.T, ourRenewalEmails bool) int {
	t.Helper()
	pool := grandfatherPool(t)
	now := time.Date(2027, 11, 10, 20, 0, 0, 0, time.UTC)
	var uid int64
	if err := pool.QueryRow(context.Background(),
		`insert into users (email) values ('daily@example.com') returning id`).Scan(&uid); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(context.Background(),
		`insert into entitlements (user_id, plan, source, status, billing_user_id, stripe_subscription_id,
		   price_lookup_key, current_period_end)
		 values ($1, 'premium', 'stripe', 'active', $1, 'sub_daily', 'premium_yearly_launch', $2)`,
		uid, now.Add(48*time.Hour)); err != nil {
		t.Fatal(err)
	}
	fake := &mail.Fake{}
	cfg := config.Config{PublicBaseURL: "https://foreversixty.gg", OurRenewalEmails: ourRenewalEmails}
	svc := newBillingDailyServices(cfg, pool, fake, slog.Default())
	for i := 0; i < 2; i++ {
		if err := billingDaily(context.Background(), slog.Default(), svc, now); err != nil {
			t.Fatal(err)
		}
	}
	return len(fake.Sent)
}
