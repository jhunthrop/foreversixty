package main

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/mail"
)

func TestBillingDaily_SendsRenewalOnce(t *testing.T) {
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
	svc := newBillingDailyServices(config.Config{PublicBaseURL: "https://foreversixty.gg"}, pool, fake, slog.Default())
	for i := 0; i < 2; i++ {
		if err := billingDaily(context.Background(), slog.Default(), svc, now); err != nil {
			t.Fatal(err)
		}
	}
	if len(fake.Sent) != 1 {
		t.Fatalf("sent %d, want 1", len(fake.Sent))
	}
}
