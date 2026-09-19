// Package phase is the content phase a fight belongs to, which is the
// bracket rankings are grouped by.
//
// The boundaries are data/curated/phases.json, copied here as a table
// rather than read at runtime: the API does not ship the site's source,
// the dates are four fixed instants, and a ranking bracket that could
// change under a running deployment would silently re-bucket stored
// rows. `python -m pipeline phases` emits the web's copy from the same
// source.
//
// When a date moves or a phase is added, change that file and this
// table together; boundaries_test.go fails if the two disagree.
package phase

import "time"

// The phase names. They are stored in fight_metrics.phase and appear in
// the rankings URL, so they are lowercase and hyphenated.
const (
	PreBeta = "pre-beta"
	Beta    = "beta"
	Launch  = "launch"
	Raids1  = "raids-1"
)

// Boundary is one phase and the instant it opens. The tags are what
// GET /v1/phases serves; pre-beta's zero Start marshals as
// "0001-01-01T00:00:00Z", which is the truth about a phase that has no
// opening instant.
type Boundary struct {
	Name  string    `json:"name"`
	Start time.Time `json:"start"`
}

// Boundaries are the phases in order. Launch is 15:00 PST on Nov 4,
// which is 23:00 UTC; the other two open at midnight UTC on their day.
var Boundaries = []Boundary{
	{PreBeta, time.Time{}},
	{Beta, time.Date(2026, 9, 17, 0, 0, 0, 0, time.UTC)},
	{Launch, time.Date(2026, 11, 4, 23, 0, 0, 0, time.UTC)},
	{Raids1, time.Date(2026, 12, 9, 0, 0, 0, 0, time.UTC)},
}

// At names the phase a moment falls in.
func At(t time.Time) string {
	name := PreBeta
	for _, b := range Boundaries {
		if !t.UTC().Before(b.Start) {
			name = b.Name
		}
	}
	return name
}

// Names lists every phase, for the rankings filter.
func Names() []string {
	out := make([]string, 0, len(Boundaries))
	for _, b := range Boundaries {
		out = append(out, b.Name)
	}
	return out
}

// Valid reports whether name is a phase.
func Valid(name string) bool {
	for _, b := range Boundaries {
		if b.Name == name {
			return true
		}
	}
	return false
}
