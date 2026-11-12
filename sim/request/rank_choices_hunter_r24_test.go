package request

import "testing"

// hunterRankCase is one level of a rank row: the spell id the rotation must cast
// and a neighbouring rank it must not.
type hunterRankCase struct {
	name, spec, preset, race string
	level                    int
	want, not                int32
}

func checkHunterRankCases(t *testing.T, cases []hunterRankCase) {
	t.Helper()
	for _, c := range cases {
		ids := rotationIDs(t, c.spec, "hunter", c.level, specRankChoices(c.spec, c.preset, c.race))
		if !ids[c.want] || ids[c.not] {
			t.Errorf("%s: want %d cast and %d not", c.name, c.want, c.not)
		}
	}
}

const (
	hMM = "hunter-marksmanship"
	hBM = "hunter-beast-mastery"
	hSV = "hunter-survival"
)
