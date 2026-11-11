package billingmail

import (
	"errors"
	"testing"
	"time"
)

func TestFormatPrice(t *testing.T) {
	cases := []struct {
		cents    int64
		interval string
		want     string
	}{
		{333, "", "$3.33"},
		{5, "", "$0.05"},
		{0, "", "$0"},
		{1000, "", "$10"},
		{1250, "", "$12.50"},
		{500, "monthly", "$5 a month"},
		{25000, "yearly", "$250 a year"},
		{4000, "yearly", "$40 a year"},
	}
	for _, c := range cases {
		got, err := FormatPrice(c.cents, c.interval)
		if err != nil || got != c.want {
			t.Errorf("FormatPrice(%d, %q) = %q, %v; want %q", c.cents, c.interval, got, err, c.want)
		}
	}
}

func TestFormatPriceRejectsNegative(t *testing.T) {
	if _, err := FormatPrice(-150, "yearly"); !errors.Is(err, ErrNegativeAmount) {
		t.Errorf("err = %v, want ErrNegativeAmount", err)
	}
}

func TestEmailDatesUsePacificTime(t *testing.T) {
	cases := []struct {
		name string
		at   time.Time
		want string
	}{
		// 05:00 UTC on 19 November is 21:00 PST on the 18th.
		{"evening", time.Date(2027, time.November, 19, 5, 0, 0, 0, time.UTC), "18 Nov 2027"},
		// 08:00 UTC on the 19th is 00:00 PST: the 19th begins.
		{"midnight", time.Date(2027, time.November, 19, 8, 0, 0, 0, time.UTC), "19 Nov 2027"},
		// Daylight time: 06:30 UTC on 15 July is 23:30 PDT on the 14th.
		{"summer", time.Date(2027, time.July, 15, 6, 30, 0, 0, time.UTC), "14 Jul 2027"},
	}
	for _, c := range cases {
		if got := EmailDate(c.at); got != c.want {
			t.Errorf("%s: EmailDate = %q, want %q", c.name, got, c.want)
		}
	}
}
