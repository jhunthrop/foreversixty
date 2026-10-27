// api/internal/billing/guild_view_test.go
package billing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
)

// payerGuild seeds a claimed guild whose verified leader holds the claim.
func (h *harness) payerGuild(t *testing.T, name string) (guildID, leaderID int64) {
	t.Helper()
	guildID, leaderID = h.officerOf(t, name)
	if _, err := h.pool.Exec(t.Context(), `update guilds set claimed_by = $2 where id = $1`, guildID, leaderID); err != nil {
		t.Fatal(err)
	}
	return guildID, leaderID
}

// guildPlan is a guild entitlements row to seed.
type guildPlan struct {
	subID     string
	billedBy  int64
	tier      string
	key       string
	cents     int64
	status    string
	trialEnd  *time.Time
	periodEnd time.Time
}

func (h *harness) seedPlan(t *testing.T, guildID int64, p guildPlan) {
	t.Helper()
	if p.status == "" {
		p.status = "active"
	}
	if p.tier == "" {
		p.tier = catalog.TierGuild
	}
	if p.periodEnd.IsZero() {
		p.periodEnd = time.Now().Add(20 * 24 * time.Hour)
	}
	if _, err := h.pool.Exec(t.Context(), `
		insert into entitlements (guild_id, plan, tier, source, status, stripe_subscription_id, billing_user_id,
		                          price_lookup_key, unit_amount_cents, trial_end, current_period_end)
		values ($1, 'guild', $2, 'stripe', $3, $4, $5, $6, $7, $8, $9)`,
		guildID, p.tier, p.status, p.subID, p.billedBy, p.key, p.cents, p.trialEnd, p.periodEnd); err != nil {
		t.Fatal(err)
	}
}

func (h *harness) seedGrantPlan(t *testing.T, guildID int64) {
	t.Helper()
	if _, err := h.pool.Exec(t.Context(), `
		insert into entitlements (guild_id, plan, tier, source, status, current_period_end)
		values ($1, 'guild', 'raid_guild', 'grant', 'active', now() + interval '30 days')`, guildID); err != nil {
		t.Fatal(err)
	}
}

func (h *harness) member(t *testing.T, guildID int64, name, rank string) int64 {
	t.Helper()
	id := h.seedUser(t, fmt.Sprintf("%s-%d@example.com", name, guildID))
	h.seedGuildMember(t, guildID, id, rank)
	return id
}

func billingPath(guildID int64, tail string) string {
	return fmt.Sprintf("/v1/guilds/%d/billing%s", guildID, tail)
}

// readBilling GETs the guild billing read as userID, failing on any status
// other than want; it returns the typed read and the raw data object.
func (h *harness) readBilling(t *testing.T, guildID, userID int64) (GuildBillingRead, map[string]json.RawMessage) {
	t.Helper()
	w := h.do(t, h.sessionRequest(t, http.MethodGet, billingPath(guildID, ""), userID, nil))
	if w.Code != http.StatusOK {
		t.Fatalf("billing read status = %d, body = %s", w.Code, w.Body.String())
	}
	return decodeBillingRead(t, w.Body.Bytes())
}

func decodeBillingRead(t *testing.T, body []byte) (GuildBillingRead, map[string]json.RawMessage) {
	t.Helper()
	var env struct {
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal(body, &env); err != nil {
		t.Fatal(err)
	}
	var read GuildBillingRead
	if err := json.Unmarshal(env.Data, &read); err != nil {
		t.Fatal(err)
	}
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(env.Data, &raw); err != nil {
		t.Fatal(err)
	}
	return read, raw
}

func TestGuildBilling_NotOnPlan(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.payerGuild(t, "noplan")
	read, raw := h.readBilling(t, gid, leader)
	if string(raw["billing"]) != "null" || read.TrialUsed || len(read.Seats) != 0 || string(raw["seats"]) != "[]" {
		t.Fatalf("read = %+v", read)
	}
	if string(raw["transfer"]) != "null" || string(raw["transfer_candidates"]) != "[]" {
		t.Fatalf("transfer = %s, candidates = %s", raw["transfer"], raw["transfer_candidates"])
	}
	if !slices.Equal(read.Can, []string{canBilling, canSeats}) {
		t.Fatalf("can = %v", read.Can)
	}
}

func TestGuildBilling_Trial(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.payerGuild(t, "trial")
	trialEnd := time.Now().Add(5 * 24 * time.Hour).UTC().Truncate(time.Second)
	h.seedPlan(t, gid, guildPlan{subID: "sub_trial", billedBy: leader, key: catalog.LookupKeyGuildFounderMonthly,
		cents: 1200, status: "trialing", trialEnd: &trialEnd})
	h.seedGuildTrial(t, gid)
	read, raw := h.readBilling(t, gid, leader)
	b := read.Billing
	if b == nil || b.Status != "trialing" || b.TrialEnd == nil || *b.TrialEnd != trialEnd.Format(time.RFC3339) {
		t.Fatalf("billing = %+v", b)
	}
	if !read.TrialUsed {
		t.Fatal("trial_used = false")
	}
	if string(raw["trial_used"]) != "true" {
		t.Fatalf("trial_used must be top-level, got %s", raw["trial_used"])
	}
}

func TestGuildBilling_FounderActive(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.payerGuild(t, "founder")
	h.seedPlan(t, gid, guildPlan{subID: "sub_f", billedBy: leader, key: catalog.LookupKeyGuildFounderYearly, cents: 12000})
	read, _ := h.readBilling(t, gid, leader)
	b := read.Billing
	if b == nil || !b.Founder || b.Interval == nil || *b.Interval != "yearly" || b.AmountCents == nil || *b.AmountCents != 12000 {
		t.Fatalf("billing = %+v", b)
	}
	if b.Scope != "plan" || b.Tier != catalog.TierGuild || !b.YouAreBillingContact || b.BilledBy == "" {
		t.Fatalf("billing = %+v", b)
	}
	if len(read.Seats) != catalog.SeatsGuild {
		t.Fatalf("seats = %+v", read.Seats)
	}
	if !slices.Equal(read.Can, []string{canBilling, canSeats, canTransfer}) {
		t.Fatalf("can = %v", read.Can)
	}
}

func TestGuildBilling_GrantPriceFactsAreNull(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.payerGuild(t, "grant")
	h.seedGrantPlan(t, gid)
	read, raw := h.readBilling(t, gid, leader)
	var billing map[string]json.RawMessage
	if err := json.Unmarshal(raw["billing"], &billing); err != nil {
		t.Fatal(err)
	}
	if string(billing["interval"]) != "null" || string(billing["amount_cents"]) != "null" || string(billing["billed_by"]) != `""` {
		t.Fatalf("billing = %s", raw["billing"])
	}
	if len(read.Seats) != catalog.SeatsRaidGuild || slices.Contains(read.Can, canTransfer) {
		t.Fatalf("seats = %d, can = %v", len(read.Seats), read.Can)
	}
}

func TestGuildBilling_ReadOnlyOfficerHasNoCan(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.payerGuild(t, "readonly")
	officer := h.member(t, gid, "officer", "officer")
	h.seedPlan(t, gid, guildPlan{subID: "sub_ro", billedBy: leader, key: catalog.LookupKeyGuildMonthly, cents: 1500})
	read, raw := h.readBilling(t, gid, officer)
	if string(raw["can"]) != "[]" || read.Billing == nil || read.Billing.YouAreBillingContact {
		t.Fatalf("can = %s, billing = %+v", raw["can"], read.Billing)
	}
	if string(raw["transfer_candidates"]) != "[]" {
		t.Fatalf("candidates = %s", raw["transfer_candidates"])
	}
}

func TestGuildBilling_FrozenClaimantHasNoCan(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.payerGuild(t, "frozen")
	h.svc.Guilds = fakeFrozenClaimant{frozenUserID: leader}
	_, raw := h.readBilling(t, gid, leader)
	if string(raw["can"]) != "[]" {
		t.Fatalf("can = %s", raw["can"])
	}
}

func TestGuildBilling_MemberGets404(t *testing.T) {
	h := newHarness(t)
	gid, _ := h.payerGuild(t, "members")
	member := h.member(t, gid, "member", "member")
	stranger := h.seedUser(t, "stranger@example.com")
	for _, uid := range []int64{member, stranger} {
		w := h.do(t, h.sessionRequest(t, http.MethodGet, billingPath(gid, ""), uid, nil))
		if w.Code != http.StatusNotFound {
			t.Fatalf("user %d: status = %d", uid, w.Code)
		}
	}
	w := h.do(t, h.sessionRequest(t, http.MethodGet, "/v1/guilds/nope/billing", member, nil))
	if w.Code != http.StatusNotFound {
		t.Fatalf("bad id: status = %d", w.Code)
	}
}

func TestGuildBilling_VerifiedCountIncludesLeader(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.payerGuild(t, "count")
	h.member(t, gid, "a", "member")
	h.member(t, gid, "b", "officer")
	unverified := h.seedUser(t, "unverified@example.com")
	if _, err := h.pool.Exec(t.Context(), `insert into guild_members (guild_id, user_id, rank, refreshed_at)
		values ($1, $2, 'member', now())`, gid, unverified); err != nil {
		t.Fatal(err)
	}
	read, _ := h.readBilling(t, gid, leader)
	if read.VerifiedCount != 3 {
		t.Fatalf("verified_count = %d", read.VerifiedCount)
	}
}

func TestGuildBilling_SweepsDepartedHolderFirst(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.payerGuild(t, "sweep")
	h.seedPlan(t, gid, guildPlan{subID: "sub_sw", billedBy: leader, key: catalog.LookupKeyGuildMonthly, cents: 1500})
	gone := h.seedUser(t, "gone@example.com")
	if _, err := h.pool.Exec(t.Context(), `insert into premium_seats (guild_id, slot, state, holder_user_id, given_by, given_at, starts_at, movable_at)
		values ($1, 1, 'held', $2, $3, now(), now(), now() + interval '30 days')`, gid, gone, leader); err != nil {
		t.Fatal(err)
	}
	read, _ := h.readBilling(t, gid, leader)
	if len(read.Seats) != 1 || read.Seats[0].State != "empty" || read.Seats[0].Holder != nil {
		t.Fatalf("seats = %+v", read.Seats)
	}
}

// seedCharacter gives userID a verified character named name in guildID.
func (h *harness) seedCharacter(t *testing.T, guildID, userID int64, name string) {
	t.Helper()
	key := fmt.Sprintf("us/hardcore/%s-%d", name, userID)
	if _, err := h.pool.Exec(t.Context(),
		`insert into characters (key, region, ruleset, name, class, user_id) values ($1, 'us', 'hardcore', $2, 'warrior', $3)`,
		key, name, userID); err != nil {
		t.Fatal(err)
	}
	if _, err := h.pool.Exec(t.Context(),
		`insert into guild_characters (guild_id, character_key, user_id, verified_at) values ($1, $2, $3, now())`,
		guildID, key, userID); err != nil {
		t.Fatal(err)
	}
}

func TestGuildBilling_BilledByNamesThePayersCharacter(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.payerGuild(t, "billedby")
	if _, err := h.pool.Exec(t.Context(), `update users set battletag = 'Kraggor#1234' where id = $1`, leader); err != nil {
		t.Fatal(err)
	}
	h.seedCharacter(t, gid, leader, "Kraggor")
	officer := h.member(t, gid, "officer", "officer")
	h.seedPlan(t, gid, guildPlan{subID: "sub_bb", billedBy: leader, key: catalog.LookupKeyGuildMonthly, cents: 1500})
	read, _ := h.readBilling(t, gid, officer)
	if read.Billing == nil || read.Billing.BilledBy != "Kraggor" {
		t.Fatalf("billed_by = %+v, want the character name Kraggor", read.Billing)
	}
}

func TestGuildBilling_BilledByFallsBackToPublicName(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.payerGuild(t, "billedbyfallback")
	officer := h.member(t, gid, "officer", "officer")
	h.seedPlan(t, gid, guildPlan{subID: "sub_bbf", billedBy: leader, key: catalog.LookupKeyGuildMonthly, cents: 1500})
	read, _ := h.readBilling(t, gid, officer)
	if read.Billing == nil || read.Billing.BilledBy != fmt.Sprintf("user-%d", leader) {
		t.Fatalf("billed_by = %+v, want the public name", read.Billing)
	}
}
