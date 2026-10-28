package coach

import (
	"strings"
	"testing"
)

func refsValidate(t *testing.T, d ReviewDoc) []Problem {
	t.Helper()
	return Validator{Names: buildNames(t)}.Validate(d, examplePack(), fixturePersona(t))
}

func hasProblemContaining(probs []Problem, kind, contains string) bool {
	for _, p := range probs {
		if p.Kind == kind && strings.Contains(p.Value, contains) {
			return true
		}
	}
	return false
}

// f1's fix-first item cannot cite f6's anchor: the "show me" link would open another
// finding's moment under f1's text.
func TestValidate_FixFirstCannotCiteAnotherFindingsAnchor(t *testing.T) {
	d := exampleReview()
	d.FixFirst.Anchors = []string{"a1", "a4"}
	if probs := refsValidate(t, d); !hasProblemContaining(probs, ProblemShape, `anchor "a4"`) {
		t.Fatalf("a4 belongs to f6: %+v", probs)
	}
}

func TestValidate_ThenFixCannotCiteAnotherFindingsAnchor(t *testing.T) {
	d := exampleReview()
	d.ThenFix[0].Anchors = []string{"a1"} // f6's item citing f1's anchor
	if probs := refsValidate(t, d); !hasProblemContaining(probs, ProblemShape, `anchor "a1"`) {
		t.Fatalf("a1 belongs to f1: %+v", probs)
	}
}

// §5.3: f1's item cites a5 (f4's rage cap) through l1, the link that joins f1 and f4.
func TestValidate_AnAnchorOfALinkedFindingIsOwnedThroughTheLink(t *testing.T) {
	if probs := refsValidate(t, exampleReview()); len(probs) != 0 {
		t.Fatalf("the worked example validates: %+v", probs)
	}
	d := exampleReview()
	d.FixFirst.Links = []string{}
	d.FixFirst.Why = "Bloodthirst is the button the Fury rotation is built around, and 9 of a possible 40 went unused while you had the rage."
	if probs := refsValidate(t, d); !hasProblemContaining(probs, ProblemShape, `anchor "a5"`) {
		t.Fatalf("without l1, a5 is not f1's: %+v", probs)
	}
}

// A finding item may cite only links that join its own finding.
func TestValidate_AFindingItemsLinkMustJoinItsFinding(t *testing.T) {
	d := exampleReview()
	d.ThenFix[0].Links = []string{"l1"} // f6's item; l1 joins f1 and f4
	if probs := refsValidate(t, d); !hasProblemContaining(probs, ProblemLink, `link "l1"`) {
		t.Fatalf("l1 does not join f6: %+v", probs)
	}
}

func TestValidate_DoingWellCitesOnlyItsStrengthsAnchors(t *testing.T) {
	d := exampleReview()
	d.DoingWell[0].Anchors = []string{"a1"}
	if probs := refsValidate(t, d); !hasProblemContaining(probs, ProblemShape, `anchor "a1"`) {
		t.Fatalf("s1 has no anchors: %+v", probs)
	}
}
