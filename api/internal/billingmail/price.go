// Package billingmail builds the billing emails, drawn in the mail package's branded shell.
package billingmail

import (
	"errors"
	"fmt"
	"time"
	_ "time/tzdata" // the zone must load on minimal images

	"github.com/jhunthrop/foreversixty/api/internal/copygen"
	"github.com/jhunthrop/foreversixty/api/internal/textfmt"
)

const (
	intervalMonthly = "monthly"
	intervalYearly  = "yearly"
	centsPerDollar  = 100

	// emailZone is the founder window's zone; dates read in it so a Pacific
	// evening shows that evening's date.
	emailZone = "America/Los_Angeles"
)

// ErrNegativeAmount rejects a negative price.
var ErrNegativeAmount = errors.New("billingmail: negative amount")

// ErrNoMemberLimit rejects a size notice built without the tier's member limit.
var ErrNoMemberLimit = errors.New("billingmail: size notice needs the member limit")

var pacific = mustLoadZone(emailZone)

func mustLoadZone(name string) *time.Location {
	loc, err := time.LoadLocation(name)
	if err != nil {
		panic(fmt.Sprintf("billingmail: load zone %s: %v", name, err))
	}
	return loc
}

// FormatPrice renders cents as "$5 a month", "$40 a year" (the app's plan.price.* words), or a
// bare amount for an empty interval. Cents show only when the amount is not whole.
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
		return copygen.Fill(copygen.PlanPriceMonthly, "{amount}", amount), nil
	case intervalYearly:
		return copygen.Fill(copygen.PlanPriceYearly, "{amount}", amount), nil
	default:
		return amount, nil
	}
}

// EmailDate renders a date in the Pacific zone as the account page writes a plan's date:
// "6 Jan 2027" (formats.json day_month_year).
func EmailDate(t time.Time) string { return textfmt.DayMonthYear(t, pacific) }
