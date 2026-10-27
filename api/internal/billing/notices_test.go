// api/internal/billing/notices_test.go
package billing

import (
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
	"github.com/jhunthrop/foreversixty/api/internal/mail"
)

var noticeNow = time.Date(2027, 11, 10, 20, 0, 0, 0, time.UTC)

type noticeHarness struct {
	*harness
	mailer *mail.Fake
}

func newNoticeHarness(t *testing.T) *noticeHarness {
	t.Helper()
	h := newHarness(t)
	m := &mail.Fake{}
	h.svc.Mailer = m
	return &noticeHarness{harness: h, mailer: m}
}

// seedRow inserts a Stripe row for billingUser; guildID 0 makes a Premium row.
func (h *noticeHarness) seedRow(t *testing.T, guildID, billingUser int64, tier, lookupKey string, periodEnd time.Time, canceling bool) int64 {
	t.Helper()
	var id int64
	var err error
	if guildID == 0 {
		err = h.pool.QueryRow(t.Context(),
			`insert into entitlements (user_id, plan, source, status, billing_user_id, stripe_subscription_id,
			   price_lookup_key, current_period_end, cancel_at_period_end)
			 values ($1, 'premium', 'stripe', 'active', $1, $2, $3, $4, $5) returning id`,
			billingUser, fmt.Sprintf("sub_%d_%s", billingUser, lookupKey), lookupKey, periodEnd, canceling).Scan(&id)
	} else {
		err = h.pool.QueryRow(t.Context(),
			`insert into entitlements (guild_id, plan, tier, source, status, billing_user_id, stripe_subscription_id,
			   price_lookup_key, current_period_end, cancel_at_period_end)
			 values ($1, 'guild', $2, 'stripe', 'active', $3, $4, $5, $6, $7) returning id`,
			guildID, tier, billingUser, fmt.Sprintf("sub_g%d", guildID), lookupKey, periodEnd, canceling).Scan(&id)
	}
	if err != nil {
		t.Fatal(err)
	}
	return id
}

func (h *noticeHarness) seedMembers(t *testing.T, guildID int64, n int) {
	t.Helper()
	for i := 0; i < n; i++ {
		h.seedGuildMember(t, guildID, h.seedUser(t, fmt.Sprintf("m%d-%d-%d@example.com", guildID, n, i)), "member")
	}
}

func (h *noticeHarness) send(t *testing.T) int {
	t.Helper()
	n, err := h.svc.SendNotices(t.Context(), noticeNow)
	if err != nil {
		t.Fatal(err)
	}
	return n
}

func inDays(d int) time.Time { return noticeNow.Add(time.Duration(d) * 24 * time.Hour) }

func TestNotices_RenewalYearlyOnly(t *testing.T) {
	h := newNoticeHarness(t)
	u := h.seedUser(t, "monthly@example.com")
	h.seedRow(t, 0, u, "", catalog.LookupKeyPremiumMonthlyLaunch, inDays(3), false)
	if n := h.send(t); n != 0 || len(h.mailer.Sent) != 0 {
		t.Fatalf("monthly row sent %d", n)
	}
}

func TestNotices_RenewalPremiumAndGuild(t *testing.T) {
	h := newNoticeHarness(t)
	u := h.seedUser(t, "payer@example.com")
	h.seedRow(t, 0, u, "", catalog.LookupKeyPremiumFounderYearly, inDays(7), false)
	g := h.seedGuild(t, "Olympus XXVII", true)
	h.seedRow(t, g, u, catalog.TierGuild, catalog.LookupKeyGuildFounderYearly, inDays(2), false)
	h.seedRow(t, h.seedGuild(t, "Later", true), u, catalog.TierGuild, catalog.LookupKeyGuildYearly, inDays(8), false)
	if n := h.send(t); n != 2 {
		t.Fatalf("sent %d, want 2", n)
	}
	var subjects, bodies []string
	for _, m := range h.mailer.Sent {
		subjects, bodies = append(subjects, m.Subject), append(bodies, m.Text)
	}
	all := strings.Join(subjects, "|") + strings.Join(bodies, "|")
	for _, want := range []string{
		"Your Premium plan renews on 17 November", "Your Guild plan renews on 12 November",
		"https://foreversixty.gg/account#plan", "https://foreversixty.gg/guild/us/hardcore/olympus-xxvii/settings#billing",
		"$120 a year, your founder price",
	} {
		if !strings.Contains(all, want) {
			t.Errorf("missing %q in %q", want, all)
		}
	}
}

// A transfer's incoming subscription trials until the outgoing period ends;
// its first yearly charge gets the 7-day email, and its size email, too.
func TestNotices_TrialingRowsAreNoticed(t *testing.T) {
	h := newNoticeHarness(t)
	u := h.seedUser(t, "trialing@example.com")
	g := h.seedGuild(t, "Trialing", true)
	id := h.seedRow(t, g, u, catalog.TierGuild, catalog.LookupKeyGuildYearly, inDays(3), false)
	h.seedMembers(t, g, catalog.GuildTierMemberLimit+1)
	if _, err := h.pool.Exec(t.Context(), `update entitlements set status = 'trialing' where id = $1`, id); err != nil {
		t.Fatal(err)
	}
	if n := h.send(t); n != 2 {
		t.Fatalf("sent %d, want the renewal and the size email", n)
	}
}

func TestNotices_RenewalSkipsCanceling(t *testing.T) {
	h := newNoticeHarness(t)
	u := h.seedUser(t, "leaving@example.com")
	h.seedRow(t, 0, u, "", catalog.LookupKeyPremiumYearlyLaunch, inDays(3), true)
	if n := h.send(t); n != 0 {
		t.Fatalf("canceling row sent %d", n)
	}
}

func TestNotices_RenewalOncePerPeriod(t *testing.T) {
	h := newNoticeHarness(t)
	u := h.seedUser(t, "once@example.com")
	id := h.seedRow(t, 0, u, "", catalog.LookupKeyPremiumYearlyLaunch, inDays(3), false)
	if h.send(t) != 1 || h.send(t) != 0 || len(h.mailer.Sent) != 1 {
		t.Fatalf("sent %d messages", len(h.mailer.Sent))
	}
	// A renewed period is a new notice.
	if _, err := h.pool.Exec(t.Context(),
		`update entitlements set current_period_end = $2 where id = $1`, id, inDays(4)); err != nil {
		t.Fatal(err)
	}
	if h.send(t) != 1 {
		t.Fatal("new period not sent")
	}
}

func TestNotices_SendFailureReleasesClaim(t *testing.T) {
	h := newNoticeHarness(t)
	u := h.seedUser(t, "retry@example.com")
	h.seedRow(t, 0, u, "", catalog.LookupKeyPremiumYearlyLaunch, inDays(3), false)
	h.mailer.Err = errors.New("resend down")
	if n, err := h.svc.SendNotices(t.Context(), noticeNow); n != 0 || err == nil {
		t.Fatalf("n=%d err=%v", n, err)
	}
	h.mailer.Err = nil
	if h.send(t) != 1 {
		t.Fatal("retry did not send")
	}
}

func TestNotices_SizeAbove25Only(t *testing.T) {
	h := newNoticeHarness(t)
	u := h.seedUser(t, "boss@example.com")
	g := h.seedGuild(t, "Olympus XXVII", true)
	h.seedRow(t, g, u, catalog.TierGuild, catalog.LookupKeyGuildFounderMonthly, inDays(12), false)
	h.seedMembers(t, g, 25)
	if n := h.send(t); n != 0 {
		t.Fatalf("25 members sent %d", n)
	}
	h.seedMembers(t, g, 1)
	if n := h.send(t); n != 1 {
		t.Fatalf("26 members sent %d", n)
	}
	m := h.mailer.Sent[0]
	if m.Subject != "Olympus XXVII has 26 members; Guild covers 25" {
		t.Errorf("subject %q", m.Subject)
	}
	if !strings.Contains(m.Text, "Raid guild at $20 a month") || !strings.Contains(m.Text, "/guild/us/hardcore/olympus-xxvii/settings#billing") {
		t.Errorf("body %q", m.Text)
	}
}

func TestNotices_SizeOncePerPeriod(t *testing.T) {
	h := newNoticeHarness(t)
	u := h.seedUser(t, "boss2@example.com")
	g := h.seedGuild(t, "Olympus XXVII", true)
	h.seedRow(t, g, u, catalog.TierGuild, catalog.LookupKeyGuildYearly, inDays(12), false)
	h.seedMembers(t, g, 26)
	if h.send(t) != 1 || h.send(t) != 0 || len(h.mailer.Sent) != 1 {
		t.Fatalf("sent %d messages", len(h.mailer.Sent))
	}
}

func TestNotices_SizeRaidGuildTierNever(t *testing.T) {
	h := newNoticeHarness(t)
	u := h.seedUser(t, "raid@example.com")
	g := h.seedGuild(t, "Big", true)
	h.seedRow(t, g, u, catalog.TierRaidGuild, catalog.LookupKeyRaidGuildMonthly, inDays(5), false)
	h.seedMembers(t, g, 40)
	if n := h.send(t); n != 0 {
		t.Fatalf("raid guild tier sent %d", n)
	}
}

func TestNotices_NoEmailSkipped(t *testing.T) {
	h := newNoticeHarness(t)
	var u int64
	if err := h.pool.QueryRow(t.Context(), `insert into users (email) values (null) returning id`).Scan(&u); err != nil {
		t.Skip("users.email is not nullable: " + err.Error())
	}
	h.seedRow(t, 0, u, "", catalog.LookupKeyPremiumYearlyLaunch, inDays(3), false)
	if n := h.send(t); n != 0 || len(h.mailer.Sent) != 0 {
		t.Fatalf("sent %d", n)
	}
}
