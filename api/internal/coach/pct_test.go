package coach

import (
	"testing"

	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
)

func TestPercentOfScalesAShareToPercent(t *testing.T) {
	if got := 0.25 * percentOf; got != 25 {
		t.Fatalf("a quarter is %v percent", got)
	}
	if percentOf-12.5 != 87.5 {
		t.Fatalf("the complement of 12.5 is %v", percentOf-12.5)
	}
}

func TestCountedBetween(t *testing.T) {
	ex := []ratingengine.Window{{FromMS: 10, ToMS: 20}, {FromMS: 50, ToMS: 60}}
	cases := []struct {
		name     string
		from, to int64
		want     int64
	}{
		{"no overlap", 20, 50, 30},
		{"whole range with both windows", 0, 100, 80},
		{"partial overlap", 15, 55, 30},
		{"inside a window", 12, 18, 0},
		{"empty or reversed", 30, 30, 0},
	}
	for _, c := range cases {
		if got := countedBetween(ex, c.from, c.to); got != c.want {
			t.Errorf("%s: got %d, want %d", c.name, got, c.want)
		}
	}
}
