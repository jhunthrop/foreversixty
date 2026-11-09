package creators

import (
	"fmt"
	"time"
)

// ShareOf is the creator's share of a payment's base, rounded down to the cent.
func ShareOf(baseCents int64) int64 {
	if baseCents <= 0 {
		return 0
	}
	return baseCents * SharePercent / 100
}

// PaymentBase is what a payment earns a share on: the balance transaction's net (after Stripe's
// fees), less the invoice's tax unless the fees already took it (Stripe as merchant of record
// lists remitted tax among the fees).
func PaymentBase(netCents, taxCents int64, taxInFees bool) int64 {
	if taxInFees {
		return netCents
	}
	return netCents - taxCents
}

// RefundBase is the part of a refund that comes off the share: the refunded amount less its
// proportional tax. Stripe keeps its fee on a refund, so nothing else is returned.
func RefundBase(refundCents, amountCents, taxCents int64) int64 {
	if amountCents <= 0 || refundCents <= 0 {
		return 0
	}
	return refundCents - refundCents*taxCents/amountCents
}

// Settle is a month's total under the cap: earnings less refunds, never below zero, never above
// the cap. capped is what the cap took off.
func Settle(bountyCents, shareCents, refundCents, capCents int64) (total, capped int64) {
	sum := bountyCents + shareCents - refundCents
	if sum <= 0 {
		return 0, 0
	}
	if sum > capCents {
		return capCents, sum - capCents
	}
	return sum, 0
}

// Month is a UTC calendar month.
type Month struct{ Start time.Time }

// MonthOf is the month t falls in.
func MonthOf(t time.Time) Month {
	t = t.UTC()
	return Month{Start: time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)}
}

// ParseMonth reads "2026-11".
func ParseMonth(s string) (Month, error) {
	t, err := time.Parse("2006-01", s)
	if err != nil {
		return Month{}, fmt.Errorf("creators: month %q: want YYYY-MM", s)
	}
	return Month{Start: t.UTC()}, nil
}

// End is the first instant after the month.
func (m Month) End() time.Time { return m.Start.AddDate(0, 1, 0) }

// String is "2026-11".
func (m Month) String() string { return m.Start.Format("2006-01") }

// Closed reports whether the month has ended at now.
func (m Month) Closed(now time.Time) bool { return !now.Before(m.End()) }

// LinkCounts reports whether a link click at clickedAt attributes an account made at signedUpAt
// to a creator who started on startsOn: the account is new since the start, and the click came
// before the sign-up by at most LinkWindow.
func LinkCounts(startsOn, clickedAt, signedUpAt time.Time) bool {
	if signedUpAt.Before(startsOn) || clickedAt.After(signedUpAt) {
		return false
	}
	return signedUpAt.Sub(clickedAt) <= LinkWindow
}

// CodeCounts reports whether a code redeemed by an account made at signedUpAt counts for a
// creator who started on startsOn.
func CodeCounts(startsOn, signedUpAt time.Time) bool { return !signedUpAt.Before(startsOn) }

// SecondDay reports whether a signed-in request at seen is on a later UTC day than the sign-up
// and inside ReturnWindow.
func SecondDay(signedUpAt, seen time.Time) bool {
	if !seen.Before(signedUpAt.Add(ReturnWindow)) {
		return false
	}
	return utcDay(seen).After(utcDay(signedUpAt))
}

// WeekOf is the Monday (UTC) of t's week, the only date a creator sees of a player.
func WeekOf(t time.Time) time.Time {
	d := utcDay(t)
	offset := (int(d.Weekday()) + 6) % 7
	return d.AddDate(0, 0, -offset)
}

func utcDay(t time.Time) time.Time {
	t = t.UTC()
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}

// StripeStatusOf turns an account's facts into the page's status.
func StripeStatusOf(a AccountFacts) string {
	switch {
	case a.ID == "":
		return StatusNone
	case a.TransfersActive && a.PayoutsEnabled:
		return StatusReady
	case a.DetailsSubmitted || a.CurrentlyDue > 0:
		return StatusNeedsInfo
	default:
		return StatusOnboarding
	}
}
