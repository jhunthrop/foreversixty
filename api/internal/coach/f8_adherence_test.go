package coach

import (
	"testing"
)

func TestAdherence(t *testing.T) {
	cases := []struct {
		name       string
		in         Input
		wantSev    Severity
		wantSpells []string
		none       bool
	}{
		{name: "dps 70% vs median 90%", in: furyAdherence(t, 0.70, poolAdherence(0.90, 25)),
			wantSev: SeverityMinor, wantSpells: []string{"Hamstring", "Rend", "Overpower"}},
		{name: "dps 60% vs 90% is major", in: furyAdherence(t, 0.60, poolAdherence(0.90, 25)), wantSev: SeverityMajor},
		{name: "no pool: under 85% of the written rotation", in: furyAdherence(t, 0.80, noPool(3)), wantSev: SeverityMinor},
		{name: "no pool, 88%: nothing", in: furyAdherence(t, 0.88, noPool(3)), none: true},
		{name: "healer: off", in: holyAdherence(t, 0.40), none: true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			f := adherence(tc.in, CountedTime(tc.in)).Finding
			if tc.none != (f == nil) {
				t.Fatalf("finding %+v", f)
			}
			if tc.none {
				return
			}
			if f.Severity != tc.wantSev || len(f.Anchors) == 0 {
				t.Fatalf("got %+v", f)
			}
			assertSpells(t, f.Facts, tc.wantSpells...)
		})
	}
}

func TestAdherenceTextAndAnchors(t *testing.T) {
	in := furyAdherence(t, 0.70, poolAdherence(0.90, 25))
	o := adherence(in, CountedTime(in))
	f := o.Finding
	if f == nil {
		t.Fatal("no finding")
	}
	if f.Title != "70% of casts on the rotation" {
		t.Errorf("title %q", f.Title)
	}
	if f.Detail != "Most-cast spells outside the written rotation: Hamstring (15), Rend (9) and Overpower (6)." {
		t.Errorf("detail %q", f.Detail)
	}
	if f.Basis.Kind != BasisTopLogs || f.Basis.Label != "Top 25 Fury warriors on Onyxia (40): 90% on the rotation. You: 70%." {
		t.Errorf("basis %+v", f.Basis)
	}
	assertNumbers(t, f.Facts, "70", "90", "25")
	if len(o.Anchors) != 3 || len(f.Anchors) != 3 {
		t.Fatalf("anchors %+v / %v", o.Anchors, f.Anchors)
	}
	for i, a := range o.Anchors {
		if a.Kind != AnchorCast || a.SpellID != offSpells[i].id {
			t.Errorf("anchor %d = %+v", i, a)
		}
	}
}

func TestAdherenceRotationAloneBasis(t *testing.T) {
	in := furyAdherence(t, 0.70, noPool(3))
	f := adherence(in, CountedTime(in)).Finding
	if f == nil || f.Severity != SeverityMajor || f.Basis.Kind != BasisOwnLog {
		t.Fatalf("got %+v", f)
	}
}

func TestAdherenceTank(t *testing.T) {
	in := protAdherence(t, 0.60, poolAdherence(0.90, 25))
	f := adherence(in, CountedTime(in)).Finding
	if !in.Tables.RotationOK {
		t.Fatal("the Protection fixture must hold a shipped rotation")
	}
	if f == nil || f.Severity != SeverityMajor {
		t.Fatalf("got %+v", f)
	}
}

func TestAdherenceNoRotationClaimsNothing(t *testing.T) {
	in := furyAdherence(t, 0.40, poolAdherence(0.90, 25))
	in.Tables.RotationOK = false
	if o := adherence(in, CountedTime(in)); o.Finding != nil || len(o.Strengths) != 0 || textOf(o) != "" {
		t.Fatalf("got %+v", o)
	}
}

func TestAdherenceTooFewCasts(t *testing.T) {
	in := furyAdherence(t, 0.50, poolAdherence(0.90, 25))
	in.Fight = withoutCasts(in.Fight, in.Row.GUID)
	in.Fight = withCasts(in.Fight, in.Row.GUID, adherenceHamstring, "Hamstring", repeat(0, 9)...)
	if o := adherence(in, CountedTime(in)); o.Finding != nil || len(o.Strengths) != 0 {
		t.Fatalf("got %+v", o)
	}
}

func TestAdherenceStrengthAtOrAboveMedian(t *testing.T) {
	in := furyAdherence(t, 0.95, poolAdherence(0.90, 25))
	o := adherence(in, CountedTime(in))
	if o.Finding != nil || len(o.Strengths) != 1 {
		t.Fatalf("got %+v", o)
	}
	s := o.Strengths[0]
	if s.Sentence != "95% of your casts were on the written rotation. Top 25: 90%." {
		t.Errorf("sentence %q", s.Sentence)
	}
	assertNumbers(t, s.Facts, "95", "90")
	if len(s.Anchors) == 0 || len(o.Anchors) != len(s.Anchors) || o.Anchors[0].Kind != AnchorCast || o.Anchors[0].SpellID != adherenceRotationID {
		t.Errorf("strength anchors %v / %+v", s.Anchors, o.Anchors)
	}
}

func TestAdherenceStrengthExactlyAtMedian(t *testing.T) {
	in := furyAdherence(t, 0.90, poolAdherence(0.90, 25))
	o := adherence(in, CountedTime(in))
	if o.Finding != nil || len(o.Strengths) != 1 {
		t.Fatalf("got %+v", o)
	}
}

func TestAdherenceRotationAloneHasNoStrength(t *testing.T) {
	in := furyAdherence(t, 0.99, noPool(3))
	if o := adherence(in, CountedTime(in)); o.Finding != nil || len(o.Strengths) != 0 {
		t.Fatalf("got %+v", o)
	}
}

// A pool whose median has no adherence falls back to the written rotation alone.
func TestAdherencePoolMedianWithoutAdherence(t *testing.T) {
	refs := poolOf(25, Profile{})
	low := furyAdherence(t, 0.80, refs)
	f := adherence(low, CountedTime(low)).Finding
	if f == nil || f.Severity != SeverityMinor || f.Basis.Kind != BasisOwnLog {
		t.Fatalf("got %+v", f)
	}
	high := furyAdherence(t, 0.99, refs)
	if o := adherence(high, CountedTime(high)); o.Finding != nil || len(o.Strengths) != 0 {
		t.Fatalf("got %+v", o)
	}
}

// Whole-point boundaries land on the spec's bands: fires 10 under, major 25 under, and
// without a pool fires under 85 and is major 25 under the whole rotation.
func TestAdherenceBandEdges(t *testing.T) {
	cases := []struct {
		name         string
		median, you  float64 // whole percent; median 0 = no pool
		fires, major bool
	}{
		{"100/90 is 10 under", 100, 90, true, false},
		{"100/91 is 9 under", 100, 91, false, false},
		{"90/80 is 10 under", 90, 80, true, false},
		{"90/81 is 9 under", 90, 81, false, false},
		{"90/65 is 25 under", 90, 65, true, true},
		{"90/66 is 24 under", 90, 66, true, false},
		{"100/75 is 25 under", 100, 75, true, true},
		{"no pool 85", 0, 85, false, false},
		{"no pool 84", 0, 84, true, false},
		{"no pool 76", 0, 76, true, false},
		{"no pool 75", 0, 75, true, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			refs := noPool(3)
			if tc.median > 0 {
				refs = poolAdherence(tc.median/100, 25)
			}
			in := furyAdherence(t, tc.you/100, refs)
			f := adherence(in, CountedTime(in)).Finding
			if tc.fires != (f != nil) {
				t.Fatalf("finding %+v", f)
			}
			if f != nil && (f.Severity == SeverityMajor) != tc.major {
				t.Fatalf("severity %s", f.Severity)
			}
		})
	}
}

func TestAdherenceBetweenThresholdsIsQuiet(t *testing.T) {
	in := furyAdherence(t, 0.85, poolAdherence(0.90, 25))
	if o := adherence(in, CountedTime(in)); o.Finding != nil || len(o.Strengths) != 0 {
		t.Fatalf("got %+v", o)
	}
}
