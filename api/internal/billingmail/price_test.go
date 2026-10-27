package billingmail

import (
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
		{1000, "", "$10"},
		{1250, "", "$12.50"},
		{500, "monthly", "$5 a month"},
		{25000, "yearly", "$250 a year"},
		{4000, "yearly", "$40 a year"},
	}
	for _, c := range cases {
		if got := FormatPrice(c.cents, c.interval); got != c.want {
			t.Errorf("FormatPrice(%d, %q) = %q, want %q", c.cents, c.interval, got, c.want)
		}
	}
}

func TestEmailDates(t *testing.T) {
	d := time.Date(2027, time.January, 6, 23, 30, 0, 0, time.FixedZone("x", -8*3600))
	if got := EmailDate(d); got != "7 January 2027" {
		t.Errorf("EmailDate = %q (UTC expected)", got)
	}
	if got := EmailDayMonth(d); got != "7 January" {
		t.Errorf("EmailDayMonth = %q", got)
	}
}
