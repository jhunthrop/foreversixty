package creators

import (
	"strings"
	"testing"
	"time"
)

func at(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err)
	}
	return t
}

func TestShareAndBases(t *testing.T) {
	if ShareOf(407) != 203 || ShareOf(0) != 0 || ShareOf(-5) != 0 {
		t.Fatalf("ShareOf: %d %d %d", ShareOf(407), ShareOf(0), ShareOf(-5))
	}
	// The sandbox's own numbers (spec section 2): $4.50 paid, $0.43 fee, net $4.07, no tax.
	if PaymentBase(407, 0, false) != 407 {
		t.Fatal("no tax: the base is the net")
	}
	if PaymentBase(1000, 80, false) != 920 || PaymentBase(1000, 80, true) != 1000 {
		t.Fatal("tax comes off once: not again when Stripe's fees already took it")
	}
	if RefundBase(200, 450, 0) != 200 || RefundBase(540, 540, 40) != 500 || RefundBase(100, 0, 0) != 0 {
		t.Fatalf("RefundBase: %d %d", RefundBase(200, 450, 0), RefundBase(540, 540, 40))
	}
}

func TestSettleUnderTheCap(t *testing.T) {
	cases := []struct {
		bounty, share, refund, cap, total, capped int64
	}{
		{12800, 4310, 250, 50000, 16860, 0},
		{6200, 1240, 0, 5000, 5000, 2440},
		{0, 100, 300, 50000, 0, 0},
		{0, 0, 0, 50000, 0, 0},
		{50000, 0, 0, 50000, 50000, 0},
	}
	for _, c := range cases {
		total, capped := Settle(c.bounty, c.share, c.refund, c.cap)
		if total != c.total || capped != c.capped {
			t.Errorf("Settle(%d,%d,%d,%d) = %d,%d; want %d,%d", c.bounty, c.share, c.refund, c.cap, total, capped, c.total, c.capped)
		}
	}
}

func TestLinkWindowAndStartDate(t *testing.T) {
	start := at("2026-11-08T00:00:00Z")
	signup := at("2026-12-10T12:00:00Z")
	cases := []struct {
		name    string
		start   time.Time
		click   time.Time
		signup  time.Time
		counted bool
	}{
		{"click a day before", start, signup.Add(-24 * time.Hour), signup, true},
		{"exactly 30 days", start, signup.Add(-LinkWindow), signup, true},
		{"30 days and a second", start, signup.Add(-LinkWindow - time.Second), signup, false},
		{"click after sign-up", start, signup.Add(time.Minute), signup, false},
		{"account before the start", at("2026-12-11T00:00:00Z"), signup.Add(-time.Hour), signup, false},
		{"account on the start day", at("2026-12-10T00:00:00Z"), signup.Add(-time.Hour), signup, true},
	}
	for _, c := range cases {
		if got := LinkCounts(c.start, c.click, c.signup); got != c.counted {
			t.Errorf("%s: LinkCounts = %v", c.name, got)
		}
	}
	if CodeCounts(start, start.Add(-time.Second)) || !CodeCounts(start, start) {
		t.Fatal("a code counts from the start date on")
	}
}

func TestSecondDayIsALaterUTCDayInsideFourteenDays(t *testing.T) {
	signup := at("2026-12-01T23:30:00Z")
	cases := []struct {
		seen time.Time
		want bool
	}{
		{at("2026-12-01T23:59:00Z"), false},            // same UTC day
		{at("2026-12-02T00:01:00Z"), true},             // 31 minutes later, the next UTC day
		{signup.Add(ReturnWindow - time.Minute), true}, // inside day 14
		{signup.Add(ReturnWindow), false},              // at day 14 exactly
		{signup.Add(ReturnWindow + time.Hour), false},  // past it
		{at("2026-12-01T22:00:00Z"), false},            // before sign-up
	}
	for _, c := range cases {
		if got := SecondDay(signup, c.seen); got != c.want {
			t.Errorf("SecondDay(%s) = %v, want %v", c.seen, got, c.want)
		}
	}
}

func TestWeekOfIsTheMonday(t *testing.T) {
	for in, want := range map[string]string{
		"2026-11-23T00:00:00Z": "2026-11-23", "2026-11-29T23:59:00Z": "2026-11-23", "2026-11-30T00:00:00Z": "2026-11-30",
	} {
		if got := WeekOf(at(in)).Format(time.DateOnly); got != want {
			t.Errorf("WeekOf(%s) = %s, want %s", in, got, want)
		}
	}
}

func TestMonths(t *testing.T) {
	m, err := ParseMonth("2026-11")
	if err != nil || m.String() != "2026-11" || !m.End().Equal(at("2026-12-01T00:00:00Z")) {
		t.Fatalf("ParseMonth = %v %v", m, err)
	}
	if m.Closed(at("2026-11-30T23:59:59Z")) || !m.Closed(at("2026-12-01T00:00:00Z")) {
		t.Fatal("a month closes at its end")
	}
	if _, err := ParseMonth("2026-13"); err == nil {
		t.Fatal("a bad month parsed")
	}
	if MonthOf(at("2026-12-31T23:00:00-05:00")).String() != "2027-01" {
		t.Fatal("MonthOf reads UTC")
	}
	if TransferGroup(7, m) != "creator-7-2026-11" {
		t.Fatalf("group = %s", TransferGroup(7, m))
	}
}

func TestStripeStatusOf(t *testing.T) {
	cases := map[string]AccountFacts{
		StatusNone:       {},
		StatusOnboarding: {ID: "acct_1"},
		StatusNeedsInfo:  {ID: "acct_1", DetailsSubmitted: true},
		StatusReady:      {ID: "acct_1", DetailsSubmitted: true, PayoutsEnabled: true, TransfersActive: true},
	}
	for want, facts := range cases {
		if got := StripeStatusOf(facts); got != want {
			t.Errorf("StripeStatusOf(%+v) = %s, want %s", facts, got, want)
		}
	}
	if StripeStatusOf(AccountFacts{ID: "acct_1", CurrentlyDue: 2}) != StatusNeedsInfo {
		t.Fatal("requirements due is needs_info")
	}
}

func TestFieldRules(t *testing.T) {
	for _, ok := range []string{"GRIMBREW", "ABCD", "A1B2C3"} {
		if !ValidCode(ok) {
			t.Errorf("ValidCode(%q) = false", ok)
		}
	}
	for _, bad := range []string{"abc", "ABC", "GRIM BREW", strings.Repeat("A", 21), "GRIM-BREW"} {
		if ValidCode(bad) {
			t.Errorf("ValidCode(%q) = true", bad)
		}
	}
	if NormalizeCode(" grimbrew ") != "GRIMBREW" {
		t.Fatal("NormalizeCode")
	}
	if !ValidChannel("https://twitch.tv/grimbrew") || ValidChannel("http://twitch.tv/x") || ValidChannel("https://") || ValidChannel("https://a b") {
		t.Fatal("ValidChannel")
	}
}

func TestApplicationValidation(t *testing.T) {
	n := int64(1200)
	good := ApplicationInput{ChannelURL: " https://twitch.tv/emberfall ", Platform: "twitch", Audience: &n, About: " Hardcore runs ", ContactEmail: ""}
	if bad := good.Validate(); len(bad) != 0 || good.About != "Hardcore runs" {
		t.Fatalf("good input refused: %v", bad)
	}
	neg := int64(-1)
	cases := []struct {
		in    ApplicationInput
		field string
	}{
		{ApplicationInput{ChannelURL: "https://evil.example/x", Platform: "twitch", About: "x"}, "channel_url"},
		{ApplicationInput{ChannelURL: "https://youtube.com/@x", Platform: "twitch", About: "x"}, "channel_url"},
		{ApplicationInput{ChannelURL: "https://twitch.tv/", Platform: "twitch", About: "x"}, "channel_url"},
		{ApplicationInput{ChannelURL: "https://twitch.tv/x", Platform: "kick", About: "x"}, "platform"},
		{ApplicationInput{ChannelURL: "https://twitch.tv/x", Platform: "twitch", About: "  "}, "about"},
		{ApplicationInput{ChannelURL: "https://twitch.tv/x", Platform: "twitch", About: strings.Repeat("é", MaxAbout+1)}, "about"},
		{ApplicationInput{ChannelURL: "https://twitch.tv/x", Platform: "twitch", About: "x", Audience: &neg}, "audience"},
		{ApplicationInput{ChannelURL: "https://twitch.tv/x", Platform: "twitch", About: "x", ContactEmail: "not an email"}, "contact_email"},
		{ApplicationInput{ChannelURL: "https://twitch.tv/x", Platform: "twitch", About: "x\x00y"}, "about"},
		{ApplicationInput{ChannelURL: "https://twitch.tv/x", Platform: "twitch", About: "x\x1b[31my"}, "about"},
	}
	lines := ApplicationInput{ChannelURL: "https://twitch.tv/x", Platform: "twitch", About: "line one\nline two\ttabbed"}
	if bad := lines.Validate(); len(bad) != 0 {
		t.Errorf("line breaks and tabs refused: %v", bad)
	}
	for _, c := range cases {
		if bad := c.in.Validate(); bad[c.field] == "" {
			t.Errorf("%+v: no %s error (%v)", c.in, c.field, bad)
		}
	}
	exact := ApplicationInput{ChannelURL: "https://www.youtube.com/@salt", Platform: "youtube", About: strings.Repeat("é", MaxAbout)}
	if bad := exact.Validate(); len(bad) != 0 {
		t.Fatalf("500 characters refused: %v", bad)
	}
}

func TestSuggest(t *testing.T) {
	cases := map[string]Suggested{
		"https://twitch.tv/emberfall":            {Name: "emberfall", Slug: "emberfall", Code: "EMBERFALL"},
		"https://youtube.com/@Salt_Marsh.Guides": {Name: "Salt_Marsh.Guides", Slug: "salt-marsh-guides", Code: "SALTMARSHGUIDES"},
	}
	for in, want := range cases {
		if got := Suggest(in); got != want {
			t.Errorf("Suggest(%s) = %+v, want %+v", in, got, want)
		}
	}
}
