package coach

import (
	"strings"
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
	if f.Title != "70% of casts from the rotation" {
		t.Errorf("title %q", f.Title)
	}
	if !strings.Contains(f.Detail, "Hamstring, Rend, Overpower") {
		t.Errorf("detail %q", f.Detail)
	}
	if f.Basis.Kind != BasisTopLogs || !strings.Contains(f.Basis.Label, "90% of casts from the rotation. You: 70%.") {
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
		t.Skip("the Protection fixture has no shipped rotation")
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
	if s.Sentence != "95% of casts from the rotation, at or above the comparison." {
		t.Errorf("sentence %q", s.Sentence)
	}
	assertNumbers(t, s.Facts, "95", "90")
}

func TestAdherenceBetweenThresholdsIsQuiet(t *testing.T) {
	in := furyAdherence(t, 0.85, poolAdherence(0.90, 25))
	if o := adherence(in, CountedTime(in)); o.Finding != nil || len(o.Strengths) != 0 {
		t.Fatalf("got %+v", o)
	}
}
