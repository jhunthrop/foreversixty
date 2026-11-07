package leveling

import (
	"path/filepath"
	"testing"
)

// Sim-accuracy round 29 (walks-heal): each healer's leveling order spent points on talents
// that heal nothing in the band's fight (or a toll it did not need yet) while a live healing
// talent waited, and the order that takes the live ones first heals more at every level it
// changes, seeds 1 and 50001, both factions, on the band's own heal fight
// (design/reviews/2026-11-07-sim-audit-round29/walksheal.md). Each want is a level the lane
// measured.
func TestHealerWalksTakeTheirLiveTalentsFirst(t *testing.T) {
	repoRoot := filepath.Join("..", "..")
	active, err := ReadActiveBuild(repoRoot)
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		spec  string
		level int
		want  string
	}{
		// Gift of Nature, Improved Rejuvenation, Swiftmend, Living Spirit and Genesis before
		// the rest of Nature's Focus, Naturalist and Tranquil Spirit; Nature's Splendor at 44
		// through its Balance rows; Nature's Swiftness (behind Naturalist 5) last.
		{"druid-restoration", 18, "0000000000000000-00000000000000000000-1503000000000000"},
		{"druid-restoration", 25, "0100000000000000-00000000000000000000-3503003100000000"},
		{"druid-restoration", 30, "0100000000000000-00000000000000000000-3503004103100000"},
		{"druid-restoration", 40, "0500200000000000-00000000000000000000-3503005103103000"},
		{"druid-restoration", 44, "0530200100000000-00000000000000000000-3503005103103000"},
		{"druid-restoration", 50, "0530200100000000-00000000000000000000-5503005113103210"},
		// Reverence before Healing Light, Illumination before Divine Favor, Divine Intellect
		// before the dead tolls, Benediction and Conviction before Holy Shock, and the
		// Protection points that do nothing for a healer last.
		{"paladin-holy", 18, "50220000000000000-00000000000000000-00000000000000000"},
		{"paladin-holy", 25, "50320005000100000-00000000000000000-00000000000000000"},
		{"paladin-holy", 30, "50320015000100000-40000000000000000-00000000000000000"},
		{"paladin-holy", 37, "50320015000100000-50000000000000000-05010000000000000"},
		{"paladin-holy", 40, "50320115000100000-50000000000000000-05030000000000000"},
		{"paladin-holy", 45, "50320215012101000-50000000000000000-05030000000000000"},
		{"paladin-holy", 50, "50320215012101230-50000000000000000-05030000000000000"},
		// Improved Renew before Twilight Focus, Holy Binding at 25, Spiritual Guidance before
		// the last Divine Fury points, and Meditation right behind its Discipline rows.
		{"priest-holy", 10, "000000000000000000-01000000000000000-00000000000000000"},
		{"priest-holy", 20, "000000000000000000-03502001000000000-00000000000000000"},
		{"priest-holy", 25, "000000000000000000-23502003000100000-00000000000000000"},
		{"priest-holy", 30, "000000000000000000-33502003030110000-00000000000000000"},
		{"priest-holy", 40, "000000000000000000-33502003030121431-00000000000000000"},
		{"priest-holy", 51, "005005000010000000-33502003030121431-00000000000000000"},
		// Meditation 3 before Mental Agility; Improved Renew and Holy Specialization before
		// Renewed Hope and Divine Aegis; Improved Healing before the last Divine Fury points.
		{"priest-discipline", 22, "005005000030000000-00000000000000000-00000000000000000"},
		{"priest-discipline", 30, "005005031330001000-00000000000000000-00000000000000000"},
		{"priest-discipline", 35, "005005031330011000-13000000000000000-00000000000000000"},
		{"priest-discipline", 40, "005005031330011000-33300000000000000-00000000000000000"},
		{"priest-discipline", 50, "005005031330011430-33501000000000000-00000000000000000"},
		// Tidal Focus before Improved Healing Wave, Restorative Totems and Healing Way before
		// Healing Focus and Tidal Mastery, Riptide at 40, Healing Focus last.
		{"shaman-restoration", 10, "0000000000000000-000000000000000000-0100000000000000"},
		{"shaman-restoration", 14, "0000000000000000-000000000000000000-2300000000000000"},
		{"shaman-restoration", 20, "0000000000000000-000000000000000000-2530000010000000"},
		{"shaman-restoration", 25, "0000000000000000-000000000000000000-5530001010010000"},
		{"shaman-restoration", 30, "0000000000000000-000000000000000000-5530001010411000"},
		{"shaman-restoration", 40, "0000000000000000-000000000000000000-5530001011513051"},
		{"shaman-restoration", 50, "0000000000000000-055000000000000000-5530001011513051"},
	}
	for _, c := range cases {
		info := ladderFirstSpecs[c.spec]
		trees, err := LoadTalentTrees(repoRoot, active, info.class)
		if err != nil {
			t.Fatal(err)
		}
		_, digits, err := GuideBuildTalents(repoRoot, info.class, info.slug)
		if err != nil {
			t.Fatal(err)
		}
		targets := GuideTalentTargets(trees, digits)
		if got := LadderTalentString(trees, targets, info.tree, c.level, ForSpec(c.spec)); got != c.want {
			t.Errorf("%s level %d: %s, want %s", c.spec, c.level, got, c.want)
		}
	}
}

// A reservation may name another tree's deeper talent: it is taken in list order once the
// ranks the list reserved before it fill that tree's rows above it and its prerequisite
// (Nature's Splendor needs ten Balance points and Nature's Majesty 2). Listed ahead of
// them, it is never held illegally: the walk spends the point in the spec's own tree.
func TestReservationHoldsADeeperOffTreeTalentOnceItsRowsAre(t *testing.T) {
	repoRoot := filepath.Join("..", "..")
	active, err := ReadActiveBuild(repoRoot)
	if err != nil {
		t.Fatal(err)
	}
	trees, err := LoadTalentTrees(repoRoot, active, "druid")
	if err != nil {
		t.Fatal(err)
	}
	_, digits, err := GuideBuildTalents(repoRoot, "druid", "restoration")
	if err != nil {
		t.Fatal(err)
	}
	targets := GuideTalentTargets(trees, digits)
	rows := []LadderFirstRank{{Talent: "Genesis", Rank: 5}, {Talent: "Nature's Majesty", Rank: 2}, {Talent: "Moonglow", Rank: 3}}
	splendor := LadderFirstRank{Talent: "Nature's Splendor", Rank: 1}
	after := withFirst(append(append([]LadderFirstRank{}, rows...), splendor))
	ahead := withFirst(append([]LadderFirstRank{splendor}, rows...))
	cases := []struct {
		name  string
		opt   LadderOption
		level int
		want  string
	}{
		{"listed after its rows, level 14", after, 14, "0500000000000000-00000000000000000000-0000000000000000"},
		{"listed after its rows, level 19", after, 19, "0530200000000000-00000000000000000000-0000000000000000"},
		{"listed after its rows, level 20", after, 20, "0530200100000000-00000000000000000000-0000000000000000"},
		{"listed ahead of its rows, level 20", ahead, 20, "0530200000000000-00000000000000000000-1000000000000000"},
	}
	for _, c := range cases {
		if got := LadderTalentString(trees, targets, 2, c.level, c.opt); got != c.want {
			t.Errorf("%s: %s, want %s", c.name, got, c.want)
		}
	}
}
