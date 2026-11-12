package creators

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestAttributionFirstTouchWindowAndSelf(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	a := h.creator("grimbrew", "2026-11-08")
	b := h.creator("lunawick", "2026-11-08")
	signup := at("2026-12-01T10:00:00Z")

	first := h.user(signup)
	if ok, err := h.svc.Store.AttributeLink(ctx, first, a.Slug, signup.Add(-48*time.Hour)); err != nil || !ok {
		t.Fatalf("first link: %v %v", ok, err)
	}
	// One account, one creator: a later link or code never moves it.
	if ok, _ := h.svc.Store.AttributeLink(ctx, first, b.Slug, signup.Add(-time.Hour)); ok {
		t.Fatal("a second link moved the account")
	}
	if ok, _ := h.svc.Store.AttributeCode(ctx, first, b.Code, signup.Add(time.Hour)); ok {
		t.Fatal("a code moved a link player")
	}
	if id, ok, _ := h.svc.Store.PlayerCreator(ctx, first); !ok || id != a.ID {
		t.Fatalf("player of %d, want %d", id, a.ID)
	}

	late := h.user(signup)
	if ok, _ := h.svc.Store.AttributeLink(ctx, late, a.Slug, signup.Add(-LinkWindow-time.Minute)); ok {
		t.Fatal("a click older than 30 days counted")
	}
	// The code still counts for that account: the rules are link OR code.
	if ok, err := h.svc.Store.AttributeCode(ctx, late, b.Code, signup.Add(time.Hour)); err != nil || !ok {
		t.Fatalf("code after a stale link: %v %v", ok, err)
	}

	old := h.user(at("2026-11-01T00:00:00Z"))
	if ok, _ := h.svc.Store.AttributeLink(ctx, old, a.Slug, at("2026-10-31T00:00:00Z")); ok {
		t.Fatal("an account older than the creator's start counted")
	}

	// The creator's own account never counts, and binding it drops it if it was counted.
	self := h.user(signup)
	if ok, _ := h.svc.Store.AttributeLink(ctx, self, a.Slug, signup.Add(-time.Hour)); !ok {
		t.Fatal("setup: the account should count before it is bound")
	}
	if err := h.svc.bind(ctx, a.ID, self); err != nil {
		t.Fatal(err)
	}
	if _, ok, _ := h.svc.Store.PlayerCreator(ctx, self); ok {
		t.Fatal("the creator's own account is still a player")
	}
	again := h.user(signup)
	h.exec(`update creators set user_id = $1 where id = $2`, again, b.ID)
	if ok, _ := h.svc.Store.AttributeCode(ctx, again, b.Code, signup); ok {
		t.Fatal("a creator redeemed their own code")
	}

	// A removed creator takes no new players.
	if err := h.svc.Remove(ctx, a.ID); err != nil {
		t.Fatal(err)
	}
	if ok, _ := h.svc.Store.AttributeLink(ctx, h.user(signup), a.Slug, signup.Add(-time.Hour)); ok {
		t.Fatal("a removed creator's link counted")
	}
}

func TestLinkArrivalThroughTheHook(t *testing.T) {
	h := newHarness(t)
	c := h.creator("grimbrew", "2026-11-08")
	uid := h.user(h.now.Add(-time.Minute))
	if err := h.svc.LinkArrival(context.Background(), uid, c.Slug, h.now.Add(-10*24*time.Hour)); err != nil {
		t.Fatal(err)
	}
	if _, ok, _ := h.svc.Store.PlayerCreator(context.Background(), uid); !ok {
		t.Fatal("the arrival did not attribute")
	}
}

func TestReturningPlayerRule(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	c := h.ready(h.creator("grimbrew", "2026-11-08"))
	signup := at("2026-11-10T22:00:00Z")
	count := func() int64 {
		t.Helper()
		st := h.statement(c, "2026-11")
		return st.BountyCount
	}
	uid := h.user(signup)
	if ok, _ := h.svc.Store.AttributeLink(ctx, uid, c.Slug, signup.Add(-time.Hour)); !ok {
		t.Fatal("setup")
	}
	if ok, _ := h.svc.Store.MarkVisit(ctx, uid, signup.Add(time.Hour)); ok {
		t.Fatal("the same UTC day counted as a second day")
	}
	if ok, _ := h.svc.Store.MarkVisit(ctx, uid, signup.Add(3*time.Hour)); !ok {
		t.Fatal("the next UTC day did not count")
	}
	if count() != 0 {
		t.Fatal("a player with no Forever character counted")
	}
	h.character(uid, "us-normal-thrall")
	if count() != 1 {
		t.Fatal("a returning player did not count")
	}
	h.exec(`update users set bnet_sub = null where id = $1`, uid)
	if count() != 0 {
		t.Fatal("a player without Battle.net counted")
	}
	h.exec(`update users set bnet_sub = 'restored' where id = $1`, uid)

	late := h.user(signup)
	h.svc.Store.AttributeLink(ctx, late, c.Slug, signup.Add(-time.Hour))
	h.character(late, "us-normal-late")
	if ok, _ := h.svc.Store.MarkVisit(ctx, late, signup.Add(ReturnWindow+time.Hour)); ok {
		t.Fatal("a visit after day 14 counted")
	}
	if count() != 1 {
		t.Fatalf("returning = %d, want 1", count())
	}
}

func TestACharacterQualifiesOneAccount(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	c := h.ready(h.creator("grimbrew", "2026-11-08"))
	first := h.returning(c, at("2026-11-10T10:00:00Z"), "us-normal-shared")
	h.pays(first, "in_shared", at("2026-11-12T10:00:00Z"))
	if _, err := h.svc.Approve(ctx, c.ID, mustMonth("2026-11")); err != nil {
		t.Fatal(err)
	}
	// The character moves to an alt account, which signs up through the same link.
	h.exec(`update characters set user_id = null where key = 'us-normal-shared'`)
	alt := h.user(at("2026-12-02T10:00:00Z"))
	h.svc.Store.AttributeLink(ctx, alt, c.Slug, at("2026-12-02T09:00:00Z"))
	h.exec(`update characters set user_id = $1 where key = 'us-normal-shared'`, alt)
	h.svc.Store.MarkVisit(ctx, alt, at("2026-12-03T10:00:00Z"))
	if st := h.statement(c, "2026-12"); st.BountyCount != 0 {
		t.Fatalf("an alt qualified with a character already paid for (first %d)", first)
	}
}

func mustMonth(s string) Month {
	m, err := ParseMonth(s)
	if err != nil {
		panic(err)
	}
	return m
}

func TestRevenueShareWindowRefundsAndCap(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	c := h.ready(h.creator("grimbrew", "2026-11-08"))
	uid := h.returning(c, at("2026-11-10T10:00:00Z"), "us-normal-payer")
	h.gw.invoices["in_1"] = InvoiceMoney{ChargeID: "ch_1", PaidAt: at("2026-11-11T10:00:00Z"), AmountCents: 450, FeeCents: 43, NetCents: 407}
	h.gw.invoices["in_trial"] = InvoiceMoney{ChargeID: "ch_0", PaidAt: at("2026-11-11T10:00:00Z")}
	h.gw.invoices["in_late"] = InvoiceMoney{ChargeID: "ch_9", PaidAt: at("2027-11-12T10:00:00Z"), AmountCents: 500, NetCents: 470}
	for _, inv := range []string{"in_1", "in_1", "in_trial"} {
		if err := h.svc.InvoicePaid(ctx, uid, inv); err != nil {
			t.Fatal(err)
		}
	}
	if err := h.svc.InvoicePaid(ctx, h.user(at("2026-11-10T00:00:00Z")), "in_other"); err != nil {
		t.Fatal("a non-player's invoice should cost nothing and fail nothing")
	}
	h.now = at("2027-12-01T00:00:00Z")
	if err := h.svc.InvoicePaid(ctx, uid, "in_late"); err != nil {
		t.Fatal(err)
	}
	h.now = at("2026-12-09T12:00:00Z")
	st := h.statement(c, "2026-11")
	if st.ShareCents != 203 || st.BountyCents != BountyCents || st.TotalCents != 403 {
		t.Fatalf("november = %+v", st)
	}
	var rows int
	h.pool.QueryRow(ctx, `select count(*) from creator_revenue`).Scan(&rows)
	if rows != 1 {
		t.Fatalf("revenue rows = %d, want 1 (no duplicate, no trial, nothing past 12 months)", rows)
	}

	// A refund in December is a December line at half its amount.
	h.gw.refunds["ch_1"] = []Refund{{ID: "re_1", AmountCents: 200, Created: at("2026-12-03T00:00:00Z")}}
	for range 2 {
		if err := h.svc.ChargeRefunded(ctx, "ch_1"); err != nil {
			t.Fatal(err)
		}
	}
	h.now = at("2027-01-02T00:00:00Z")
	dec := h.statement(c, "2026-12")
	if dec.RefundCents != 100 || dec.TotalCents != 403-100 {
		t.Fatalf("december (november unpaid, so it rolls in) = %+v", dec)
	}

	// The cap: a $1 cap pays $1 and shows what it took off.
	one := int64(100)
	if _, err := h.svc.Edit(ctx, c.ID, Input{Name: c.Name, ChannelURL: c.ChannelURL, StartsOn: c.StartsOn, CapCents: &one}); err != nil {
		t.Fatal(err)
	}
	capped := h.statement(c, "2026-12")
	if capped.TotalCents != 100 || capped.CappedCents != 203 {
		t.Fatalf("capped = %+v", capped)
	}
}

func TestApproveIsIdempotentAndNeverAutomatic(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	c := h.creator("grimbrew", "2026-11-08")
	h.pays(h.returning(c, at("2026-11-10T10:00:00Z"), "us-normal-a"), "in_a", at("2026-11-12T10:00:00Z"))
	nov := mustMonth("2026-11")

	if _, err := h.svc.Approve(ctx, c.ID, nov); !errors.Is(err, ErrNotPayable) {
		t.Fatalf("approve before onboarding = %v, want ErrNotPayable", err)
	}
	c = h.ready(c)
	if _, err := h.svc.Approve(ctx, c.ID, mustMonth("2026-12")); !errors.Is(err, ErrMonthOpen) {
		t.Fatalf("approve an open month = %v", err)
	}
	h.svc.PayoutsEnabled = false
	if _, err := h.svc.Approve(ctx, c.ID, nov); !errors.Is(err, ErrPayoutsOff) {
		t.Fatalf("approve with payouts off = %v", err)
	}
	h.svc.PayoutsEnabled = true

	// The transfer fails once: the payout stays pending, and the next press finishes it.
	h.gw.failNext = errors.New("stripe: balance_insufficient")
	if _, err := h.svc.Approve(ctx, c.ID, nov); err == nil {
		t.Fatal("a failed transfer was reported paid")
	}
	if st := h.statement(c, "2026-11"); st.Status != StatementReady || st.TransferID != nil {
		t.Fatalf("after a failed transfer = %+v", st)
	}
	st, err := h.svc.Approve(ctx, c.ID, nov)
	if err != nil || st.Status != StatementPaid || st.TransferID == nil || *st.TransferID != "tr_creator-"+itoa(c.ID)+"-2026-11" {
		t.Fatalf("approve = %+v %v", st, err)
	}
	again, err := h.svc.Approve(ctx, c.ID, nov)
	if err != nil || h.gw.creates != 1 || *again.TransferID != *st.TransferID {
		t.Fatalf("second approve: creates = %d, %+v %v", h.gw.creates, again, err)
	}
	if p := h.gw.transfers[TransferGroup(c.ID, nov)]; p.AmountCents != BountyCents+1000 || p.Destination != c.StripeAccountID {
		t.Fatalf("transfer = %+v", p)
	}
	// The settled lines never come back.
	h.now = at("2027-01-02T00:00:00Z")
	if dec := h.statement(c, "2026-12"); dec.TotalCents != 0 || dec.Status != StatementNothing {
		t.Fatalf("december after november's payout = %+v", dec)
	}
	if _, err := h.svc.Approve(ctx, c.ID, mustMonth("2026-12")); !errors.Is(err, ErrNothingOwed) {
		t.Fatalf("approving a $0 month = %v", err)
	}
}

func itoa(n int64) string { return formatInt(n) }

func formatInt(n int64) string {
	if n == 0 {
		return "0"
	}
	var b []byte
	for n > 0 {
		b = append([]byte{byte('0' + n%10)}, b...)
		n /= 10
	}
	return string(b)
}

func TestCreatorPremiumGrantAndItsEnd(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	c := h.creator("grimbrew", "2026-11-08")
	uid := h.user(at("2026-11-01T00:00:00Z"))
	url, err := h.svc.Invite(ctx, c.ID)
	if err != nil || !strings.HasPrefix(url, "https://foreversixty.gg/creators/onboarding?invite=") {
		t.Fatalf("invite = %q %v", url, err)
	}
	token := strings.TrimPrefix(url, "https://foreversixty.gg/creators/onboarding?invite=")
	if _, err := h.svc.AcceptInvite(ctx, token, uid); err != nil {
		t.Fatal(err)
	}
	if _, err := h.svc.AcceptInvite(ctx, token, h.user(h.now)); !errors.Is(err, ErrBadInvite) {
		t.Fatalf("a used invite = %v", err)
	}
	source := func(id int64) (string, string) {
		var src, status string
		h.pool.QueryRow(ctx, `select source, status from entitlements where user_id = $1 and plan = 'premium'`, id).Scan(&src, &status)
		return src, status
	}
	if src, status := source(uid); src != "creator" || status != "active" {
		t.Fatalf("after accepting = %s %s", src, status)
	}
	if err := h.svc.Remove(ctx, c.ID); err != nil {
		t.Fatal(err)
	}
	if src, status := source(uid); src != "creator" || status != "canceled" {
		t.Fatalf("after removing = %s %s", src, status)
	}

	// A paid subscription is never replaced, and removing the creator leaves it alone.
	paid := h.user(at("2026-11-01T00:00:00Z"))
	h.exec(`insert into entitlements (user_id, plan, source, status, stripe_subscription_id, billing_user_id)
		values ($1, 'premium', 'stripe', 'active', 'sub_paid', $1)`, paid)
	d := h.creator("lunawick", "2026-11-08")
	if err := h.svc.bind(ctx, d.ID, paid); err != nil {
		t.Fatal(err)
	}
	if src, status := source(paid); src != "stripe" || status != "active" {
		t.Fatalf("a paid plan after the grant = %s %s", src, status)
	}
	if err := h.svc.Remove(ctx, d.ID); err != nil {
		t.Fatal(err)
	}
	if src, status := source(paid); src != "stripe" || status != "active" {
		t.Fatalf("a paid plan after removal = %s %s", src, status)
	}
}

func TestApplicationsOneOpenApproveAndDecline(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	uid := h.user(h.now.Add(-time.Hour))
	in := ApplicationInput{ChannelURL: "https://twitch.tv/emberfall", Platform: "twitch", About: "Hardcore runs", ContactEmail: "ember@example.com"}
	if err := h.svc.Apply(ctx, uid, in); err != nil {
		t.Fatal(err)
	}
	if err := h.svc.Apply(ctx, uid, in); !errors.Is(err, ErrApplicationOpen) {
		t.Fatalf("a second open application = %v", err)
	}
	mine, _ := h.svc.Store.Application(ctx, uid, h.now)
	if mine.Status == nil || *mine.Status != ApplicationPending || mine.Creator {
		t.Fatalf("pending = %+v", mine)
	}
	reqs, _ := h.svc.Store.Requests(ctx)
	if len(reqs) != 1 || reqs[0].Suggested.Slug != "emberfall" || reqs[0].ContactEmail == nil {
		t.Fatalf("requests = %+v", reqs)
	}
	c, err := h.svc.ApproveRequest(ctx, reqs[0].ID, Input{Name: "Emberfall", ChannelURL: in.ChannelURL, Slug: "emberfall", Code: "EMBERFALL", StartsOn: "2026-12-09"})
	if err != nil || c.UserID == nil || *c.UserID != uid {
		t.Fatalf("approve request = %+v %v", c, err)
	}
	mine, _ = h.svc.Store.Application(ctx, uid, h.now)
	if *mine.Status != ApplicationApproved || !mine.Creator {
		t.Fatalf("approved = %+v", mine)
	}
	if err := h.svc.Apply(ctx, uid, in); !errors.Is(err, ErrApplicationOpen) {
		t.Fatal("a creator applied again")
	}
	if _, err := h.svc.ApproveRequest(ctx, reqs[0].ID, Input{}); !errors.Is(err, ErrNotPending) {
		t.Fatalf("approving twice = %v", err)
	}

	other := h.user(h.now)
	h.svc.Apply(ctx, other, ApplicationInput{ChannelURL: "https://youtube.com/@salt", Platform: "youtube", About: "Guides"})
	reqs, _ = h.svc.Store.Requests(ctx)
	if err := h.svc.DeclineRequest(ctx, reqs[0].ID); err != nil {
		t.Fatal(err)
	}
	mine, _ = h.svc.Store.Application(ctx, other, h.now)
	if *mine.Status != ApplicationDeclined {
		t.Fatalf("declined = %+v", mine)
	}
	// Declined: not again for 30 days, then they may.
	again := ApplicationInput{ChannelURL: "https://youtube.com/@salt", Platform: "youtube", About: "Guides, again"}
	if err := h.svc.Apply(ctx, other, again); !errors.Is(err, ErrReapplyTooSoon) {
		t.Fatalf("reapplying at once = %v", err)
	}
	if mine.ReapplyOn == nil || *mine.ReapplyOn != "2027-01-08" {
		t.Fatalf("reapply_on = %v", mine.ReapplyOn)
	}
	h.now = h.now.Add(ReapplyAfter)
	if err := h.svc.Apply(ctx, other, again); err != nil {
		t.Fatal(err)
	}
}

func TestAdminRoutesAreTheOwnersAlone(t *testing.T) {
	h := newHarness(t)
	uid := h.user(h.now)
	routes := []struct{ method, path string }{
		{http.MethodGet, "/v1/admin/creators"},
		{http.MethodPost, "/v1/admin/creators"},
		{http.MethodGet, "/v1/admin/creators/statements"},
		{http.MethodPost, "/v1/admin/creators/1/statements/2026-11/approve"},
		{http.MethodGet, "/v1/admin/creators/requests"},
		{http.MethodPost, "/v1/admin/creators/requests/1/approve"},
		{http.MethodPost, "/v1/admin/creators/requests/1/decline"},
		{http.MethodPost, "/v1/admin/creators/1/invite"},
		{http.MethodDelete, "/v1/admin/creators/1"},
	}
	for _, r := range routes {
		if w := h.request(r.method, r.path, anonymous, map[string]any{}); w.Code != http.StatusUnauthorized {
			t.Errorf("%s %s anonymous = %d", r.method, r.path, w.Code)
		}
		if w := h.request(r.method, r.path, member(uid), map[string]any{}); w.Code != http.StatusForbidden {
			t.Errorf("%s %s member = %d", r.method, r.path, w.Code)
		}
	}
	for _, path := range []string{"/v1/admin/creators", "/v1/admin/creators/statements?month=2026-11", "/v1/admin/creators/requests"} {
		w := h.request(http.MethodGet, path, owner, nil)
		if w.Code != http.StatusOK || !strings.Contains(w.Header().Get("Cache-Control"), "no-store") {
			t.Errorf("owner %s = %d %s", path, w.Code, w.Body)
		}
	}
	if w := h.request(http.MethodGet, "/v1/creators/me", member(uid), nil); w.Code != http.StatusForbidden {
		t.Errorf("a non-creator's stats = %d, want 403", w.Code)
	}
	if w := h.request(http.MethodGet, "/v1/creators/me", anonymous, nil); w.Code != http.StatusUnauthorized {
		t.Errorf("anonymous stats = %d, want 401", w.Code)
	}
}

func TestHandlersValidateAndAnswerTheFlow(t *testing.T) {
	h := newHarness(t)
	w := h.request(http.MethodPost, "/v1/admin/creators", owner, map[string]any{"name": "", "channel_url": "http://x", "slug": "Bad Slug", "code": "lc", "starts_on": "soon"})
	if w.Code != http.StatusBadRequest {
		t.Fatalf("bad creator = %d", w.Code)
	}
	for _, f := range []string{"name", "channel_url", "slug", "code", "starts_on"} {
		if !strings.Contains(w.Body.String(), `"`+f+`"`) {
			t.Errorf("no %s error in %s", f, w.Body)
		}
	}
	w = h.request(http.MethodPost, "/v1/admin/creators", owner, map[string]any{"name": "Grimbrew", "channel_url": "https://twitch.tv/grimbrew", "slug": "grimbrew", "code": "grimbrew", "starts_on": "2026-11-08"})
	if w.Code != http.StatusCreated || !strings.Contains(w.Body.String(), `"code":"GRIMBREW"`) {
		t.Fatalf("add = %d %s", w.Code, w.Body)
	}
	w = h.request(http.MethodPost, "/v1/admin/creators", owner, map[string]any{"name": "Again", "channel_url": "https://twitch.tv/again", "slug": "grimbrew", "code": "AGAIN", "starts_on": "2026-11-08"})
	if w.Code != http.StatusBadRequest {
		t.Fatalf("a taken slug = %d", w.Code)
	}
	if w := h.request(http.MethodPost, "/v1/creators/application", member(h.user(h.now)), map[string]any{"channel_url": "https://evil.example/x", "platform": "twitch", "about": "x"}); w.Code != http.StatusBadRequest {
		t.Fatalf("a bad application = %d", w.Code)
	}
	if w := h.request(http.MethodGet, "/v1/admin/creators/statements?month=13", owner, nil); w.Code != http.StatusBadRequest {
		t.Fatalf("a bad month = %d", w.Code)
	}
}

// Nothing a person typed (an email, their words, a channel, an invite token) reaches a log line.
func TestNoPersonalDataInLogs(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	uid := h.user(h.now)
	secret := []string{"ember.private@example.com", "Hardcore runs on the quiet realm", "twitch.tv/emberfall"}
	if w := h.request(http.MethodPost, "/v1/creators/application", member(uid), map[string]any{
		"channel_url": "https://" + secret[2], "platform": "twitch", "about": secret[1], "contact_email": secret[0]}); w.Code != http.StatusCreated {
		t.Fatalf("apply = %d %s", w.Code, w.Body)
	}
	reqs, _ := h.svc.Store.Requests(ctx)
	c, err := h.svc.ApproveRequest(ctx, reqs[0].ID, Input{Name: "Emberfall", ChannelURL: "https://" + secret[2], Slug: "emberfall", Code: "EMBERFALL", StartsOn: "2026-11-08"})
	if err != nil {
		t.Fatal(err)
	}
	url, _ := h.svc.Invite(ctx, c.ID)
	token := url[strings.LastIndex(url, "=")+1:]
	h.request(http.MethodPost, "/v1/creators/invite", member(h.user(h.now)), map[string]string{"token": token})
	h.returning(c, at("2026-11-10T10:00:00Z"), "us-normal-logs")
	logs := h.logs.String()
	for _, s := range append(secret, token, "player#") {
		if strings.Contains(logs, s) {
			t.Errorf("the logs carry %q:\n%s", s, logs)
		}
	}
	if !strings.Contains(logs, "op=apply") || !strings.Contains(logs, "op=add") {
		t.Fatalf("the operations were not logged at all:\n%s", logs)
	}
}

func TestCreatorPagesAndStripeLinks(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	c := h.creator("grimbrew", "2026-11-08")
	uid := h.user(at("2026-11-01T00:00:00Z"))
	if err := h.svc.bind(ctx, c.ID, uid); err != nil {
		t.Fatal(err)
	}
	if _, err := h.svc.Dashboard(ctx, uid); !errors.Is(err, ErrNotPayable) {
		t.Fatalf("dashboard before onboarding = %v", err)
	}
	link, err := h.svc.Onboard(ctx, uid)
	if err != nil || !strings.Contains(link, "return=https://foreversixty.gg/creators/onboarding?done=1") ||
		!strings.Contains(link, "refresh=https://foreversixty.gg/creators/onboarding?refresh=1") {
		t.Fatalf("onboarding link = %q %v", link, err)
	}
	c, _ = h.svc.Store.Creator(ctx, c.ID)
	h.gw.accounts[c.StripeAccountID] = AccountFacts{ID: c.StripeAccountID, DetailsSubmitted: true, PayoutsEnabled: true, TransfersActive: true}
	me, err := h.svc.Me(ctx, uid, true)
	if err != nil || me.StripeStatus != StatusReady {
		t.Fatalf("me after ?done=1 = %+v %v", me, err)
	}
	if dash, err := h.svc.Dashboard(ctx, uid); err != nil || !strings.HasPrefix(dash, "https://connect.stripe.test/express/") {
		t.Fatalf("dashboard = %q %v", dash, err)
	}

	// Weeks: the running week shows sign-ups only.
	h.returning(c, at("2026-11-24T10:00:00Z"), "us-normal-w1")
	h.svc.Store.AttributeLink(ctx, h.user(at("2026-12-08T10:00:00Z")), c.Slug, at("2026-12-08T09:00:00Z"))
	me, _ = h.svc.Me(ctx, uid, false)
	if len(me.Weeks) != 2 || me.Weeks[0].WeekOf != "2026-12-07" || me.Weeks[0].Returning != nil ||
		me.Weeks[1].WeekOf != "2026-11-23" || *me.Weeks[1].Returning != 1 || me.Players != 2 || me.Month != "2026-12" {
		t.Fatalf("me = %+v", me)
	}
}
