// Package creators is the creator program (design/specs/2026-11-08-creator-program.md): creators
// with a tracked link and a code of their own, the accounts each one brought, the $2 bounty per
// returning player and the 50% revenue share for a player's first 12 months, and the monthly
// statements the owner approves before any Stripe transfer is made. Every money rule lives here;
// the pages only format what this package answers.
//
// Nothing about a player reaches a log line or a URL: logs carry account and creator ids only.
package creators

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

// The program's terms (spec section 1). Changing one changes what creators are paid.
const (
	// BountyCents is paid once per returning player.
	BountyCents = 200
	// SharePercent of a payment's net goes to the player's creator.
	SharePercent = 50
	// MinPayoutCents is the least a month pays out; a smaller total carries over to the next month.
	MinPayoutCents = 1000
	// DefaultCapCents is a new creator's monthly cap.
	DefaultCapCents = 50000
	// MaxCapCents bounds what the owner may type (the creators migration's check).
	MaxCapCents = 1000000
	// CodeTrialDays is the free month a creator's code gives at checkout, on Premium or a guild plan
	// (billing reads it through CodeTrialDays()).
	CodeTrialDays = 30
)

// The windows the rules read.
const (
	// LinkWindow is how long before sign-up a creator's link click still counts.
	LinkWindow = 30 * 24 * time.Hour
	// ReturnWindow is how long after sign-up the second day must come.
	ReturnWindow = 14 * 24 * time.Hour
	// ShareMonths is how long after a player's first payment their payments earn a share.
	ShareMonths = 12
	// ReapplyAfter is how long after a decline an account may apply again.
	ReapplyAfter = 30 * 24 * time.Hour
	// InviteTTL is how long an owner's invite works.
	InviteTTL = 14 * 24 * time.Hour
	// PrivacyAge is how old a player must be before the creator sees whether they returned or pay.
	PrivacyAge = 7 * 24 * time.Hour
	// RecentWeeks is the most week rows the creator's page lists.
	RecentWeeks = 26
)

// Limits the creators migration's checks hold; the handler refuses past them first.
const (
	MaxNameLen    = 60
	MaxChannelLen = 200
)

// Stripe onboarding states, as the pages show them.
const (
	StatusNone       = "none"
	StatusOnboarding = "onboarding"
	StatusNeedsInfo  = "needs_info"
	StatusReady      = "ready"
)

// Statement states.
const (
	StatementOpen    = "open"
	StatementReady   = "ready"
	StatementNothing = "nothing"
	StatementBlocked = "blocked"
	// StatementCarry: owes money, under MinPayoutCents; its lines stay for the next month.
	StatementCarry = "carry"
	StatementPaid  = "paid"
)

// Line kinds.
const (
	LineBounty = "bounty"
	LineShare  = "share"
	LineRefund = "refund"
)

var (
	// ErrNotFound: no creator (or invite, or statement) answers that id.
	ErrNotFound = errors.New("creators: not found")
	// ErrTaken: the slug or the code is already used.
	ErrTaken = errors.New("creators: that link or code is taken")
	// ErrAlreadyCreator: the account to bind is a creator already (a second application).
	ErrAlreadyCreator = errors.New("creators: that account is already a creator")
	// ErrNotCreator: the signed-in account is not bound to a creator.
	ErrNotCreator = errors.New("creators: this account is not a creator")
	// ErrBadInvite: the invite is unknown, used or expired.
	ErrBadInvite = errors.New("creators: invite expired or used")
	// ErrMonthOpen: the month has not ended, so it cannot be approved.
	ErrMonthOpen = errors.New("creators: the month is still open")
	// ErrNotPayable: the creator's Stripe account cannot receive a transfer yet.
	ErrNotPayable = errors.New("creators: stripe cannot pay this creator yet")
	// ErrStripeOff: Stripe is not configured on this deployment.
	ErrStripeOff = errors.New("creators: stripe is not configured")
)

var (
	codePattern = regexp.MustCompile(`^[A-Z0-9]{4,20}$`)
	slugPattern = regexp.MustCompile(`^[a-z0-9][a-z0-9-]{0,39}$`)
)

// ValidCode reports whether s is a code a creator can say on stream.
func ValidCode(s string) bool { return codePattern.MatchString(s) }

// NormalizeCode is a code as typed (any case, spaces around it) as it is stored.
func NormalizeCode(s string) string { return strings.ToUpper(strings.TrimSpace(s)) }

// ValidSlug is the go link's own rule (golinks.ValidSlug, migration 0066).
func ValidSlug(s string) bool { return slugPattern.MatchString(s) }

// ValidName reports whether s is a creator's display name.
func ValidName(s string) bool {
	t := strings.TrimSpace(s)
	return t != "" && len(t) <= MaxNameLen
}

// ValidChannel reports whether s is an https address with no whitespace.
func ValidChannel(s string) bool {
	if len(s) > MaxChannelLen || !strings.HasPrefix(s, "https://") || len(s) == len("https://") {
		return false
	}
	return !strings.ContainsAny(s, " \t\r\n")
}

// Creator is one creator as the owner's page shows it.
type Creator struct {
	ID           int64   `json:"id"`
	Name         string  `json:"name"`
	ChannelURL   string  `json:"channel_url"`
	Slug         string  `json:"slug"`
	Code         string  `json:"code"`
	StartsOn     string  `json:"starts_on"`
	CapCents     int64   `json:"cap_cents"`
	StripeStatus string  `json:"stripe_status"`
	Figures      Figures `json:"figures"`

	UserID          *int64 `json:"-"`
	StripeAccountID string `json:"-"`
	// Ended: removed from the program (no new players, no invites).
	Ended bool `json:"-"`
}

// Figures are a creator's totals since the start and what is owed this month so far.
type Figures struct {
	Clicks       int64 `json:"clicks"`
	SignUps      int64 `json:"sign_ups"`
	Returning    int64 `json:"returning"`
	RevenueCents int64 `json:"revenue_cents"`
	// Month is the open month OwedCents covers; the four counts above are all time.
	Month     string `json:"month"`
	OwedCents int64  `json:"owed_cents"`
	Capped    bool   `json:"capped"`
}

// Input is what the owner types to add or edit a creator. Slug and Code are fixed after adding.
type Input struct {
	Name       string `json:"name"`
	ChannelURL string `json:"channel_url"`
	Slug       string `json:"slug"`
	Code       string `json:"code"`
	StartsOn   string `json:"starts_on"`
	CapCents   *int64 `json:"cap_cents"`
}

// Line is one statement line: a bounty, a share of a payment, or a refund.
type Line struct {
	Kind        string `json:"kind"`
	Player      int64  `json:"player"`
	On          string `json:"on"`
	AmountCents int64  `json:"amount_cents"`
}

// Statement is one creator's month.
type Statement struct {
	CreatorID   int64      `json:"creator_id"`
	Name        string     `json:"name"`
	Month       string     `json:"month"`
	Status      string     `json:"status"`
	BountyCount int64      `json:"bounty_count"`
	BountyCents int64      `json:"bounty_cents"`
	ShareCents  int64      `json:"share_cents"`
	RefundCents int64      `json:"refund_cents"`
	CappedCents int64      `json:"capped_cents"`
	TotalCents  int64      `json:"total_cents"`
	PaidAt      *time.Time `json:"paid_at"`
	TransferID  *string    `json:"transfer_id"`
	Lines       []Line     `json:"lines"`
}

// Week is one week (a Monday, UTC) of a creator's sign-ups as the creator may see it: returning
// and paying are nil while the week is still running (spec section 1, Privacy).
type Week struct {
	WeekOf    string `json:"week_of"`
	SignedUp  int64  `json:"signed_up"`
	Returning *int64 `json:"returning"`
	Paying    *int64 `json:"paying"`
}

// Payout is one approved month on the creator's own page.
type Payout struct {
	Month       string    `json:"month"`
	AmountCents int64     `json:"amount_cents"`
	PaidAt      time.Time `json:"paid_at"`
}

// Me is the creator's own page.
type Me struct {
	Name         string   `json:"name"`
	Slug         string   `json:"slug"`
	Code         string   `json:"code"`
	CapCents     int64    `json:"cap_cents"`
	Month        string   `json:"month"`
	MonthCents   int64    `json:"month_cents"`
	Players      int64    `json:"players"`
	Returning    int64    `json:"returning"`
	Paying       int64    `json:"paying"`
	Weeks        []Week   `json:"weeks"`
	Payouts      []Payout `json:"payouts"`
	StripeStatus string   `json:"stripe_status"`
}
