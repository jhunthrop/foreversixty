package sims

import (
	"strings"
	"testing"
	"unicode/utf8"

	simapi "github.com/jhunthrop/foreversixty/sim/api"
)

// withKind returns a result of the given kind, so each case below says
// only what it is about.
func withKind(kind string, mutate func(*simapi.SimResult)) simapi.SimResult {
	res := browserResult("warrior-fury", 1204.4)
	switch kind {
	case simapi.KindGear, simapi.KindTalents, simapi.KindDrops:
		res.Request.Bulk = &simapi.BulkSpec{Mode: kind, Precision: simapi.PrecisionNormal}
	case simapi.KindWeights:
		res.Request.Weights = &simapi.WeightsSpec{Reference: "crit"}
	}
	if mutate != nil {
		mutate(&res)
	}
	return res
}

func item(name string, delta float64) simapi.Combo {
	return simapi.Combo{
		Substitutions: []simapi.Substitution{
			{Kind: "item", Slot: "main_hand", ItemID: 17182, Name: name, Origin: "bag"},
		},
		Delta: simapi.Estimate{Mean: delta},
	}
}

func drop(name, source string, delta float64) simapi.Combo {
	return simapi.Combo{
		Substitutions: []simapi.Substitution{
			{Kind: "item", Slot: "main_hand", ItemID: 17182, Name: name,
				Origin: "drop:raid:molten-core:11502", SourceName: source},
		},
		Delta: simapi.Estimate{Mean: delta},
	}
}

func TestTheHeadlineSaysWhatEachKindFound(t *testing.T) {
	for _, c := range []struct {
		name string
		res  simapi.SimResult
		want string
	}{
		{
			"a run is its own dps, grouped",
			withKind(simapi.KindRun, nil),
			"1,204 DPS",
		},
		{
			"a three-figure run has no comma",
			withKind(simapi.KindRun, func(r *simapi.SimResult) { r.DPS.Mean = 999.4 }),
			"999 DPS",
		},
		{
			"a five-figure run groups once",
			withKind(simapi.KindRun, func(r *simapi.SimResult) { r.DPS.Mean = 12345.6 }),
			"12,346 DPS",
		},
		{
			"top gear names the best combination's item",
			withKind(simapi.KindGear, func(r *simapi.SimResult) {
				r.Combos = []simapi.Combo{item("Vis'kag the Bloodletter", 41.2), item("Brutality Blade", 8)}
			}),
			"+41 DPS from Vis'kag the Bloodletter",
		},
		{
			"top gear with three substitutions names the first and counts the rest",
			withKind(simapi.KindGear, func(r *simapi.SimResult) {
				r.Combos = []simapi.Combo{{
					Substitutions: []simapi.Substitution{
						{Kind: "item", ItemID: 17182, Name: "Vis'kag the Bloodletter", Origin: "bag"},
						{Kind: "item", ItemID: 16963, Name: "Onslaught Girdle", Origin: "bag"},
						{Kind: "item", ItemID: 18404, Name: "Blackhand's Breadth", Origin: "bank"},
					},
					Delta: simapi.Estimate{Mean: 63.5},
				}}
			}),
			"+64 DPS from Vis'kag the Bloodletter and 2 more",
		},
		{
			"an unnamed item falls back to its id",
			withKind(simapi.KindGear, func(r *simapi.SimResult) {
				r.Combos = []simapi.Combo{item("", 12)}
			}),
			"+12 DPS from item #17182",
		},
		{
			"a losing best combination is still signed",
			withKind(simapi.KindGear, func(r *simapi.SimResult) {
				r.Combos = []simapi.Combo{item("Brutality Blade", -13.7)}
			}),
			"-14 DPS from Brutality Blade",
		},
		{
			"a gear run with nothing in it says so",
			withKind(simapi.KindGear, nil),
			"no combinations",
		},
		{
			"talents quote the loadout",
			withKind(simapi.KindTalents, func(r *simapi.SimResult) {
				r.Combos = []simapi.Combo{{
					Substitutions: []simapi.Substitution{
						{Kind: "talents", Name: "Deep Fury", Talents: "-0550000505021051-05"},
					},
					Delta: simapi.Estimate{Mean: 18.4},
				}}
			}),
			"+18 DPS with 'Deep Fury'",
		},
		{
			"drops count the upgrades and name the source",
			withKind(simapi.KindDrops, func(r *simapi.SimResult) {
				r.Combos = []simapi.Combo{
					drop("Perdition's Blade", "Ragnaros", 55),
					drop("Spinal Reaper", "Ragnaros", 12),
					drop("Malistar's Defender", "Ragnaros", 1),
					drop("Band of Accuria", "Ragnaros", -4),
				}
			}),
			"3 upgrades on Ragnaros",
		},
		{
			"one upgrade is singular",
			withKind(simapi.KindDrops, func(r *simapi.SimResult) {
				r.Combos = []simapi.Combo{drop("Spinal Reaper", "Ragnaros", 12)}
			}),
			"1 upgrade on Ragnaros",
		},
		{
			"a two-word source reads whole",
			withKind(simapi.KindDrops, func(r *simapi.SimResult) {
				r.Combos = []simapi.Combo{drop("Maladath", "Broodlord Lashlayer", 30)}
			}),
			"1 upgrade on Broodlord Lashlayer",
		},
		{
			"nothing gained on a source is still an answer",
			withKind(simapi.KindDrops, func(r *simapi.SimResult) {
				r.Combos = []simapi.Combo{drop("Band of Accuria", "Ragnaros", -4)}
			}),
			"no upgrades on Ragnaros",
		},
		{
			"two sources in one request name neither",
			withKind(simapi.KindDrops, func(r *simapi.SimResult) {
				r.Combos = []simapi.Combo{
					drop("Perdition's Blade", "Ragnaros", 55),
					drop("Maladath", "Broodlord Lashlayer", 30),
				}
			}),
			"2 upgrades",
		},
		{
			"an unnamed source is counted, not named",
			withKind(simapi.KindDrops, func(r *simapi.SimResult) {
				r.Combos = []simapi.Combo{drop("Maladath", "", 30)}
			}),
			"1 upgrade",
		},
		{
			"weights are the top two, the reference first",
			withKind(simapi.KindWeights, func(r *simapi.SimResult) {
				r.Weights = []simapi.StatWeight{
					{Stat: "agility", Weight: 0.874},
					{Stat: "crit", Weight: 1},
					{Stat: "attack_power", Weight: 0.5},
				}
			}),
			"Crit 1.00 · Agility 0.87",
		},
		{
			"a multi-word stat id reads as words",
			withKind(simapi.KindWeights, func(r *simapi.SimResult) {
				r.Weights = []simapi.StatWeight{
					{Stat: "attack_power", Weight: 1},
					{Stat: "spell_haste", Weight: 0.4},
				}
			}),
			"Attack Power 1.00 · Spell Haste 0.40",
		},
		{
			"a weights run with nothing in it says so",
			withKind(simapi.KindWeights, nil),
			"no weights",
		},
	} {
		t.Run(c.name, func(t *testing.T) {
			if got := Headline(c.res); got != c.want {
				t.Errorf("Headline = %q, want %q", got, c.want)
			}
		})
	}
}

func TestALongHeadlineIsTrimmedRatherThanStoredWhole(t *testing.T) {
	res := withKind(simapi.KindGear, func(r *simapi.SimResult) {
		r.Combos = []simapi.Combo{item(strings.Repeat("ǝ", 400), 10)}
	})
	got := Headline(res)
	if n := len([]rune(got)); n != maxHeadline {
		t.Fatalf("%d runes, want %d", n, maxHeadline)
	}
	if !utf8.ValidString(got) {
		t.Error("the cut broke a rune in half")
	}
}

func TestThousandsAreGroupedFromTheRight(t *testing.T) {
	for _, c := range []struct {
		in   int64
		want string
	}{
		{0, "0"}, {7, "7"}, {999, "999"}, {1000, "1,000"}, {1204, "1,204"},
		{12345, "12,345"}, {1234567, "1,234,567"}, {-1500, "-1,500"},
	} {
		if got := withThousands(c.in); got != c.want {
			t.Errorf("withThousands(%d) = %q, want %q", c.in, got, c.want)
		}
	}
}
