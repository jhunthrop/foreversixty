// logs/engine/rating/insufficient.go
package rating

import "strings"

// MinCoverage is spec §1.5's dated (2026-09-21) coverage ruling: below
// this fraction of a role's total weight actually scored, a card has no
// overall at all. Found in production on a wipe with no curated mechanics
// table and no validated simulator spec — three of six components
// excluded, the surviving three renormalised to carry the full weight —
// where the resulting number (7.81, 13.97) read as a judgement built on
// almost nothing. 0.5 means at least half the role's weight must be
// backed by an actually-measured component before Overall is published;
// below it, the card still reports every component's own score and
// reason (so the page can show what WAS measured) but Overall,
// OverallUncapped and OverallCapped are zero-valued and Insufficient is
// true instead.
const MinCoverage = 0.5

// buildInsufficientReason names, in plain words built from the excluded
// components' own Reason codes, what was missing — one clause per
// distinct cause, in component order, joined "; ". Called only when
// Coverage < MinCoverage.
func buildInsufficientReason(components [6]Component) string {
	seen := map[string]bool{}
	var clauses []string
	for _, c := range components {
		if !c.Excluded {
			continue
		}
		phrase := reasonPhrase(c.Name, c.Reason)
		if seen[phrase] {
			continue
		}
		seen[phrase] = true
		clauses = append(clauses, phrase)
	}
	return strings.Join(clauses, "; ")
}

// reasonPhrase turns one excluded component's machine Reason code into a
// plain-words clause. This is the one place the engine itself writes
// reader-facing prose rather than a machine code for the web layer to
// translate (spec §5.2/§6.4's usual convention): spec §1.5's coverage
// rule can zero the whole Overall, and a blank number with no explanation
// is a worse failure than a wrong one.
func reasonPhrase(component, reason string) string {
	switch reason {
	case ReasonWipe:
		return "no output on a wipe"
	case ReasonNoMechanicsTable:
		return "no mechanics table for this encounter"
	case ReasonNoUtilityTable:
		return "no utility table for this spec"
	case ReasonNoConsumableCatalogue:
		return "no consumable catalogue for this role"
	case ReasonNotEnoughSamples:
		return "not enough logs to score " + component
	default:
		return component + " not scored"
	}
}
