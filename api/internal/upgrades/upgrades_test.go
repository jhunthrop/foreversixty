package upgrades

import (
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/bis"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

func ptr[T any](v T) *T { return &v }

func TestSlotsAreTheFrameOrder(t *testing.T) {
	var keys []string
	for _, s := range Slots {
		keys = append(keys, s.Key)
	}
	want := "head,neck,shoulder,back,chest,wrist,hands,waist,legs,feet,finger1,finger2,trinket1,trinket2,main_hand,off_hand,ranged"
	if strings.Join(keys, ",") != want {
		t.Fatalf("slots = %v", keys)
	}
	if SlotLabel("finger1") != "Finger 1" || SlotLabel("main_hand") != "Main hand" || SlotLabel("nope") != "nope" {
		t.Fatalf("labels %q %q", SlotLabel("finger1"), SlotLabel("main_hand"))
	}
}

func TestLevelBandAndPreset(t *testing.T) {
	for _, c := range []struct {
		level *int
		band  int
		ok    bool
	}{{nil, 0, false}, {ptr(19), 0, false}, {ptr(20), 20, true}, {ptr(39), 30, true}, {ptr(60), 60, true}, {ptr(63), 60, true}} {
		if b, ok := LevelBand(c.level); b != c.band || ok != c.ok {
			t.Errorf("level %v = %d %v", c.level, b, ok)
		}
	}
	if PresetFor(60) != bis.RaidPreset || PresetFor(50) != bis.BarePreset {
		t.Fatal("the Phase 1 raid is the preset at 60 only")
	}
}

// The gain rule (web/src/lib/home/upgrades.ts): a tie, a listed alternative's measured gap, a
// weapon with no listed alternative, then the stat-weight difference.
func TestJudge(t *testing.T) {
	per := 0.5
	band := bis.Band{ReferenceDpsPerPoint: &per}
	head := bis.Slot{Slot: "head", ItemID: 1, Score: ptr(100.0),
		Alternatives: []bis.Alternative{{ItemID: 2, DpsDelta: -4.2}, {ItemID: 3, DpsDelta: 0}}}
	ranged := bis.Slot{Slot: "ranged", ItemID: 10, Alternatives: []bis.Alternative{{ItemID: 11, DpsDelta: -9}}}
	score := func(id int) (float64, bool) { return map[int]float64{4: 60}[id], id == 4 }
	cases := []struct {
		name  string
		pick  bis.Slot
		worn  int
		score ItemScore
		want  Verdict
	}{
		{"the pick", head, 1, score, Verdict{Tie: true}},
		{"a zero-delta alternative", head, 3, score, Verdict{Tie: true}},
		{"a listed alternative", head, 2, nil, Verdict{Gain: ptr(4.2)}},
		{"a weapon with no listed alternative", ranged, 12, score, Verdict{NotSimChecked: true}},
		{"a weapon's listed alternative", ranged, 11, nil, Verdict{Gain: ptr(9.0)}},
		{"an empty weapon slot", ranged, 0, score, Verdict{NotSimChecked: true}},
		{"the stat-weight difference", head, 4, score, Verdict{Gain: ptr(20.0)}},
		{"an item the data lacks", head, 5, score, Verdict{}},
		{"no scorer", head, 4, nil, Verdict{}},
		{"an empty slot", head, 0, nil, Verdict{Gain: ptr(50.0)}},
		{"a pick with no score", bis.Slot{Slot: "trinket1", ItemID: 7}, 0, score, Verdict{}},
	}
	for _, c := range cases {
		got := Judge(band, c.pick, c.worn, c.score)
		if got.Tie != c.want.Tie || got.NotSimChecked != c.want.NotSimChecked || !sameGain(got.Gain, c.want.Gain) {
			t.Errorf("%s = %+v, want %+v", c.name, got, c.want)
		}
	}
	if got := Judge(bis.Band{}, head, 0, nil); got.Gain != nil {
		t.Errorf("a band with no DPS per point measures nothing: %+v", got)
	}
}

func sameGain(a, b *float64) bool {
	if a == nil || b == nil {
		return a == b
	}
	d := *a - *b
	return d < 1e-9 && d > -1e-9
}

func TestScoreItemMapsThePlannerStatNames(t *testing.T) {
	item := trees.Item{Armor: 100, Stats: map[string]int{"agility": 10, "fire_res": 5, "healing": 20, "stamina": 3}}
	weights := map[string]float64{"agility": 2, "fire_resistance": 1, "healing_power": 0.5, "armor": 0.01}
	if got := ScoreItem(item, weights); got != 20+5+10+1 {
		t.Fatalf("score = %v", got)
	}
}

func TestOrderIsTheBotsOrder(t *testing.T) {
	type row struct {
		slot  string
		group Group
		value float64
	}
	rows := []row{{"head", GroupUnmodelled, 11}, {"chest", GroupMeasured, 11}, {"wrist", GroupMeasured, 4.7},
		{"feet", GroupUnmodelled, 7.6}, {"trinket2", GroupEmpty, 0}, {"ranged", GroupMeasured, 23.1}}
	out := Order(rows, func(r row) (Group, float64) { return r.group, r.value })
	var got []string
	for _, r := range out {
		got = append(got, r.slot)
	}
	if strings.Join(got, ",") != "trinket2,ranged,chest,wrist,head,feet" {
		t.Fatalf("order = %v", got)
	}
	if rows[0].slot != "head" {
		t.Fatal("Order sorted its input in place")
	}
}

func TestUnmatchedPointsIsTheBandCompare(t *testing.T) {
	for _, c := range []struct {
		worn []string
		band string
		want int
	}{
		{[]string{"5", "0", ""}, "5-0-0", 0},
		{[]string{"53", "", ""}, "35-0-0", 2},
		{[]string{"", "", "55"}, "55-0-0", 10},
		{[]string{"5"}, "5-3", 3},
		{[]string{"5000001"}, "5", 0},
	} {
		if got := UnmatchedPoints(c.worn, c.band); got != c.want {
			t.Errorf("%v vs %q = %d, want %d", c.worn, c.band, got, c.want)
		}
	}
}

func TestWords(t *testing.T) {
	if SignedFigure(9.24, "DPS") != "+9.2 DPS" || SignedFigure(-15, "DPS") != "−15.0 DPS" {
		t.Fatalf("%q %q", SignedFigure(9.24, "DPS"), SignedFigure(-15, "DPS"))
	}
	if SourceLine("Leatherworking", "crafted") != "Leatherworking · crafted" || SourceLine("Onyxia", "raid") != "Onyxia" {
		t.Fatal("source line")
	}
	for role, want := range map[string]Unit{"tank": UnitScore, "healer": UnitHPS, "": UnitDPS} {
		if u := UnitOf(bis.Band{Role: role}); u != want {
			t.Errorf("role %q = %q", role, u)
		}
	}
	if UnitScore.Word() != "score" || UnitHPS.Word() != "HPS" || UnitDPS.Word() != "DPS" {
		t.Fatal("unit words")
	}
}
