// Package billingmail builds the plain-text billing emails.
package billingmail

import (
	"fmt"
	"time"
)

const (
	intervalMonthly = "monthly"
	intervalYearly  = "yearly"
	centsPerDollar  = 100

	dateLayout     = "2 January 2006"
	dayMonthLayout = "2 January"
)

// FormatPrice renders cents as "$5 a month", "$40 a year", or a bare amount
// for an empty interval. Cents show only when the amount is not whole.
func FormatPrice(cents int64, interval string) string {
	amount := fmt.Sprintf("$%d", cents/centsPerDollar)
	if cents%centsPerDollar != 0 {
		amount = fmt.Sprintf("$%d.%02d", cents/centsPerDollar, cents%centsPerDollar)
	}
	switch interval {
	case intervalMonthly:
		return amount + " a month"
	case intervalYearly:
		return amount + " a year"
	default:
		return amount
	}
}

// EmailDate renders a UTC date as "6 January 2027".
func EmailDate(t time.Time) string { return t.UTC().Format(dateLayout) }

// EmailDayMonth renders a UTC date as "6 January".
func EmailDayMonth(t time.Time) string { return t.UTC().Format(dayMonthLayout) }
