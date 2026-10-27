// Package billingmail builds the plain-text billing emails.
package billingmail

import (
	"errors"
	"fmt"
	"time"
	_ "time/tzdata" // the zone must load on minimal images
)

const (
	intervalMonthly = "monthly"
	intervalYearly  = "yearly"
	centsPerDollar  = 100

	// emailZone is the founder window's zone; dates read in it so a Pacific
	// evening shows that evening's date.
	emailZone = "America/Los_Angeles"

	dateLayout     = "2 January 2006"
	dayMonthLayout = "2 January"
)

// ErrNegativeAmount rejects a negative price.
var ErrNegativeAmount = errors.New("billingmail: negative amount")

var pacific = mustLoadZone(emailZone)

func mustLoadZone(name string) *time.Location {
	loc, err := time.LoadLocation(name)
	if err != nil {
		panic(fmt.Sprintf("billingmail: load zone %s: %v", name, err))
	}
	return loc
}

// FormatPrice renders cents as "$5 a month", "$40 a year", or a bare amount
// for an empty interval. Cents show only when the amount is not whole.
func FormatPrice(cents int64, interval string) (string, error) {
	if cents < 0 {
		return "", fmt.Errorf("%w: %d cents", ErrNegativeAmount, cents)
	}
	amount := fmt.Sprintf("$%d", cents/centsPerDollar)
	if cents%centsPerDollar != 0 {
		amount = fmt.Sprintf("$%d.%02d", cents/centsPerDollar, cents%centsPerDollar)
	}
	switch interval {
	case intervalMonthly:
		return amount + " a month", nil
	case intervalYearly:
		return amount + " a year", nil
	default:
		return amount, nil
	}
}

// EmailDate renders a date in the Pacific zone as "6 January 2027".
func EmailDate(t time.Time) string { return t.In(pacific).Format(dateLayout) }

// EmailDayMonth renders a date in the Pacific zone as "6 January".
func EmailDayMonth(t time.Time) string { return t.In(pacific).Format(dayMonthLayout) }
