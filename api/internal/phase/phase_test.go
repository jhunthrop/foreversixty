package phase

import (
	"testing"
	"time"
)

func TestAtNamesThePhaseFromTheSitesDates(t *testing.T) {
	for _, tc := range []struct {
		when time.Time
		want string
	}{
		{time.Date(2026, 9, 1, 0, 0, 0, 0, time.UTC), PreBeta},
		{time.Date(2026, 9, 17, 0, 0, 0, 0, time.UTC), Beta},
		{time.Date(2026, 11, 4, 22, 59, 0, 0, time.UTC), Beta},
		{time.Date(2026, 11, 4, 23, 0, 0, 0, time.UTC), Launch},
		{time.Date(2026, 12, 8, 23, 59, 0, 0, time.UTC), Launch},
		{time.Date(2026, 12, 9, 0, 0, 0, 0, time.UTC), Raids1},
		{time.Date(2027, 3, 1, 0, 0, 0, 0, time.UTC), Raids1},
	} {
		if got := At(tc.when); got != tc.want {
			t.Errorf("At(%s) = %q, want %q", tc.when, got, tc.want)
		}
	}
}

func TestNamesAndValid(t *testing.T) {
	if len(Names()) != 4 {
		t.Fatalf("names = %v", Names())
	}
	if !Valid(Raids1) || Valid("season-of-mastery") {
		t.Fatal("Valid must accept exactly the four phases")
	}
}
