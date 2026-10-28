package coach

import (
	"strings"
	"testing"
)

func TestRanks(t *testing.T) {
	cases := []struct {
		name                  string
		in                    Input
		wantTitle, wantDetail string
		wantSev               Severity
		wantCannot, none      bool
	}{
		{name: "prot Sunder Rank 4 at level 60", in: protSunderRank(t, 60, 4, 22),
			wantTitle: "Sunder Armor: 22 casts of Rank 4", wantSev: SeverityMinor},
		{name: "deliberate downrank is never flagged", in: withDownrank(protSunderRank(t, 60, 4, 22), "Sunder Armor"), none: true},
		{name: "no level: couldn't check", in: protSunderRank(t, 0, 4, 22), none: true, wantCannot: true},
		{name: "one stray cast of a once-cast family is quiet", in: protSunderRank(t, 60, 4, 1), none: true},
		{name: "2 of 2 casts below the top rank fires on the percent branch", in: protSunderRank(t, 60, 4, 2),
			wantTitle: "Sunder Armor: 2 casts of Rank 4", wantSev: SeverityMinor},
		{name: "top rank trainable is fine", in: protSunderRank(t, 60, 5, 22), none: true},
		{name: "holy Greater Heal R4 44% vs 22% is minor (22 points, §5.3)", in: holyOverheal(t, 10965, 23, 44, poolOverheal(10965, 22, 21)),
			wantTitle: "Greater Heal (Rank 4): 44% overheal", wantSev: SeverityMinor, wantDetail: "Lower ranks cost less mana."},
		{name: "holy Greater Heal R4 55% vs 22% is major (33 points, §6.2 F5)", in: holyOverheal(t, 10965, 23, 55, poolOverheal(10965, 22, 21)),
			wantTitle: "Greater Heal (Rank 4): 55% overheal", wantSev: SeverityMajor, wantDetail: "Lower ranks cost less mana."},
		{name: "holy under 15 points over is quiet", in: holyOverheal(t, 10965, 23, 36, poolOverheal(10965, 22, 21)), none: true},
		{name: "holy rank with 9 casts is ignored", in: holyOverheal(t, 10965, 9, 60, poolOverheal(10965, 22, 21)), none: true},
		{name: "holy with no pool: top logs only", in: holyOverheal(t, 10965, 23, 44, noPool(5)), none: true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			o := ranks(tc.in, CountedTime(tc.in))
			if tc.none != (o.Finding == nil) {
				t.Fatalf("finding %+v", o.Finding)
			}
			if tc.wantCannot && len(o.CouldNotCheck) == 0 {
				t.Error("no level must say so")
			}
			if tc.none {
				return
			}
			f := o.Finding
			if f.Title != tc.wantTitle || f.Severity != tc.wantSev || len(f.Anchors) == 0 {
				t.Fatalf("got %+v", f)
			}
			if tc.wantDetail != "" && !strings.HasSuffix(f.Detail, tc.wantDetail) {
				t.Errorf("the healer detail ends at %q: %q", tc.wantDetail, f.Detail)
			}
			if strings.Contains(strings.ToLower(f.Title+f.Detail), "heal more") {
				t.Error("never heal more")
			}
		})
	}
}

func TestRanksFacts(t *testing.T) {
	o := ranks(protSunderRank(t, 60, 4, 22), Counted{})
	assertNumbers(t, o.Finding.Facts, "22", "4", "5", "60")
	assertSpells(t, o.Finding.Facts, "Sunder Armor")
	if o.Finding.Detail != "Rank 5 is trainable at your level. Rank 4 was cast 22 times." {
		t.Errorf("detail %q", o.Finding.Detail)
	}
	if len(o.Anchors) != 1 || o.Anchors[0].Kind != AnchorCast || o.Anchors[0].SpellID != sunderRank4ID {
		t.Errorf("anchors %+v", o.Anchors)
	}

	h := ranks(holyOverheal(t, 10965, 23, 55, poolOverheal(10965, 22, 21)), Counted{})
	assertNumbers(t, h.Finding.Facts, "55", "22", "23", "4")
	if len(h.Anchors) != 1 || h.Anchors[0].Kind != AnchorHealing || h.Finding.Basis.Kind != BasisTopLogs {
		t.Errorf("anchors %+v basis %+v", h.Anchors, h.Finding.Basis)
	}
}

// A heal rank at or under the top logs' overheal for the same rank is a healer strength
// (A12, StrengthOverheal, §7); one over it, or with no pool, is not.
func TestRanks_OverhealAtOrUnderTheMedianIsAStrength(t *testing.T) {
	o := ranks(holyOverheal(t, 10965, 23, 18, poolOverheal(10965, 22, 21)), Counted{})
	if o.Finding != nil || len(o.Strengths) != 1 {
		t.Fatalf("finding %+v strengths %+v", o.Finding, o.Strengths)
	}
	s := o.Strengths[0]
	if s.Sentence != "Greater Heal (Rank 4): 18% overheal. Top 21: 22%." {
		t.Errorf("sentence %q", s.Sentence)
	}
	assertNumbers(t, s.Facts, "18", "22", "21", "4")
	if len(s.Anchors) != 1 || s.Anchors[0] != "#0" || len(o.Anchors) != 1 || o.Anchors[0].Kind != AnchorHealing {
		t.Errorf("anchors %v %+v", s.Anchors, o.Anchors)
	}
	for _, in := range []Input{
		holyOverheal(t, 10965, 23, 30, poolOverheal(10965, 22, 21)),
		holyOverheal(t, 10965, 23, 18, noPool(5)),
	} {
		if o := ranks(in, Counted{}); len(o.Strengths) != 0 {
			t.Errorf("strength %+v", o.Strengths)
		}
	}
}
