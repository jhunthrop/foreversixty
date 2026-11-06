package coach

import (
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"testing"
)

// The spells the talent-cooldown tests cast (1.60.1.70790/spellconst).
const (
	whirlwindID    = int64(1680)  // warrior.json: 10 s category cooldown, 25 rage
	shadowburnR6ID = int64(18871) // warlock.json: 15 s category cooldown, one Soul Shard
)

// talentCoreFight is a 60 second fight in which the player's rule file lists one core
// spell (family, a base cooldown of cooldownMS, cost checked as none) cast at castsMS.
func talentCoreFight(t *testing.T, family string, id, cooldownMS int64, level int64, castsMS ...int64) Input {
	t.Helper()
	in := protInput(t)
	in.Fight = withLevel(withDuration(withoutCasts(in.Fight, furyGUID), 60_000), furyGUID, level)
	in.Fight = withCasts(in.Fight, furyGUID, id, family, castsMS...)
	in.Row.Level = level
	in.Row.Role = RoleDPS
	in.Tables.Rules = SpecRules{Spec: "test", Build: fixtureVerified,
		Core: []RuleEntry{{Family: family, SpellIDs: []int64{id}, Verified: fixtureVerified}}}
	in.Tables.Spells = map[int64]SpellConst{id: {Name: family, CooldownMS: cooldownMS, GCDMS: 1500, CostKnown: true, CostType: PowerRage}}
	in.Tables.Trainable = nil
	return in
}

// whirlwindFight is a Fury warrior casting Whirlwind at 0:00 and 0:30 of a minute.
func whirlwindFight(t *testing.T, talents PlayerTalents) Input {
	t.Helper()
	in := talentCoreFight(t, "Whirlwind", whirlwindID, 10_000, 60, 0, 30_000)
	in.Tables.Talents = talents
	return in
}

// With Improved Whirlwind 2/2 on file, Whirlwind is an 8 second cooldown: two 22 second
// gaps hold two lost casts each, and the basis names 8 seconds.
func TestCoreCasts_ImprovedWhirlwindShortensTheCooldown(t *testing.T) {
	in := whirlwindFight(t, PlayerTalents{"Improved Whirlwind": 2})
	f := coreCasts(in, CountedTime(in)).Finding
	if f == nil || f.Title != "Whirlwind: 4 casts lost" {
		t.Fatalf("got %+v", f)
	}
	if want := "Measured against Whirlwind's 8-second cooldown and your own rage bar."; f.Basis.Label != want {
		t.Errorf("basis %q, want %q", f.Basis.Label, want)
	}
}

// Talents on file without the talent: the base 10 seconds, stated as the player's own.
func TestCoreCasts_KnownTalentsWithoutTheTalentKeepTheBase(t *testing.T) {
	in := whirlwindFight(t, PlayerTalents{"Flurry": 5})
	f := coreCasts(in, CountedTime(in)).Finding
	if f == nil || f.Title != "Whirlwind: 2 casts lost" {
		t.Fatalf("got %+v", f)
	}
	if want := "Measured against Whirlwind's 10-second cooldown and your own rage bar."; f.Basis.Label != want {
		t.Errorf("basis %q, want %q", f.Basis.Label, want)
	}
}

// Talents not on file: the base cooldown, and the basis says it is the base one and that
// the count is the fewest lost.
func TestCoreCasts_UnknownTalentsSayTheBaseCooldown(t *testing.T) {
	in := whirlwindFight(t, nil)
	f := coreCasts(in, CountedTime(in)).Finding
	if f == nil || f.Title != "Whirlwind: 2 casts lost" {
		t.Fatalf("got %+v", f)
	}
	if f.Basis.Label != Fill(BasisCooldownBase, map[string]string{"Spell": "Whirlwind", "n": "10", "resource": "rage"}) {
		t.Errorf("basis %q", f.Basis.Label)
	}
}

// A share-of-cooldown talent: Endurance 2/2 takes 60% off Evasion's five minutes.
func TestCooldownFor_PercentTalent(t *testing.T) {
	in := Input{Tables: Tables{Talents: PlayerTalents{"Endurance": 2}}}
	if ms, short := cooldownFor(in, "Evasion", SpellConst{CooldownMS: 300_000}); ms != 120_000 || short {
		t.Errorf("Evasion with Endurance 2: %d ms (shortenable %v), want 120000", ms, short)
	}
	if ms, short := cooldownFor(Input{}, "Bloodthirst", SpellConst{CooldownMS: 6_000}); ms != 6_000 || short {
		t.Errorf("a family no talent shortens: %d ms (shortenable %v)", ms, short)
	}
}

// shadowburnFight is a Destruction warlock at level who casts Shadowburn once, at 0:00, in a
// minute: three 15 second cooldowns sat unused.
func shadowburnFight(t *testing.T, level int64, talents PlayerTalents) Input {
	t.Helper()
	in := talentCoreFight(t, "Shadowburn", shadowburnR6ID, 15_000, level, 0)
	in.Tables.Spells[shadowburnR6ID] = SpellConst{Name: "Shadowburn", Rank: 6, CooldownMS: 15_000, GCDMS: 1500, CostKnown: true}
	in.Tables.Talents = talents
	return in
}

// Shadowburn costs a Soul Shard and the log shows no shards: it is judged only where every
// cast refunds its shard (Shadow and Flame 5/5), or, talents unknown, from level 44.
func TestCoreCasts_ShadowburnNeedsTheShardRefund(t *testing.T) {
	for _, tc := range []struct {
		name    string
		level   int64
		talents PlayerTalents
		judged  bool
	}{
		{"Shadow and Flame 5/5", 60, PlayerTalents{"Shadow and Flame": 5}, true},
		{"Shadow and Flame 3/5", 40, PlayerTalents{"Shadow and Flame": 3}, false},
		{"no Shadow and Flame", 60, PlayerTalents{"Ruin": 1}, false},
		{"talents unknown at 44", 44, nil, true},
		{"talents unknown at 43", 43, nil, false},
		{"talents and level unknown", 0, nil, false},
	} {
		in := shadowburnFight(t, tc.level, tc.talents)
		f := coreCasts(in, CountedTime(in)).Finding
		if judged := f != nil; judged != tc.judged {
			t.Errorf("%s: judged %v, want %v (%+v)", tc.name, judged, tc.judged, f)
		}
	}
}

// classTalentFile is one class's talents/<class>.json at the active build, the fields the
// guards read.
type classTalentFile struct {
	Trees []struct {
		Position int `json:"position"`
		Talents  []struct {
			Name    string `json:"name"`
			MaxRank int    `json:"max_rank"`
			Tier    int    `json:"tier"`
			Column  int    `json:"column"`
			Ranks   []struct {
				Description string `json:"description"`
			} `json:"ranks"`
		} `json:"talents"`
	} `json:"trees"`
}

func readClassTalents(t *testing.T, class string) classTalentFile {
	t.Helper()
	var f classTalentFile
	if err := json.Unmarshal(buildFile(t, "talents", class+".json"), &f); err != nil {
		t.Fatal(err)
	}
	return f
}

var coachClasses = []string{"druid", "hunter", "mage", "paladin", "priest", "rogue", "shaman", "warlock", "warrior"}

// cutValue matches a number and its unit in a talent's text ("by 2 sec", "by 1 min", "by 40%").
var cutValue = regexp.MustCompile(`([0-9]+(?:\.[0-9]+)?)\s*(secs?|min|%)`)

// The table's values are the client's: every rank's text names the table's cut at that rank.
func TestCooldownTalentsMatchTheClientText(t *testing.T) {
	for _, ct := range cooldownTalents {
		var descs []string
		for _, class := range coachClasses {
			for _, tree := range readClassTalents(t, class).Trees {
				for _, tal := range tree.Talents {
					if tal.Name == ct.Talent {
						for _, r := range tal.Ranks {
							descs = append(descs, r.Description)
						}
					}
				}
			}
		}
		if len(descs) == 0 {
			t.Errorf("%s: no such talent in the client data", ct.Talent)
			continue
		}
		for i, d := range descs {
			rank := float64(i + 1)
			if !textNamesCut(d, ct, rank) {
				t.Errorf("%s rank %d: the text does not name the table's cut: %q", ct.Talent, i+1, d)
			}
		}
	}
}

func textNamesCut(desc string, ct cooldownTalent, rank float64) bool {
	for _, m := range cutValue.FindAllStringSubmatch(desc, -1) {
		v, err := strconv.ParseFloat(m[1], 64)
		if err != nil {
			continue
		}
		switch {
		case m[2] == "%" && ct.CutPctRank > 0 && v == ct.CutPctRank*rank:
			return true
		case m[2] == "min" && ct.CutMSRank > 0 && v*60_000 == float64(ct.CutMSRank)*rank:
			return true
		case (m[2] == "sec" || m[2] == "secs") && ct.CutMSRank > 0 && v*1_000 == float64(ct.CutMSRank)*rank:
			return true
		}
	}
	return false
}

// Every family the table shortens is one a rule file judges, so no entry is dead.
func TestCooldownTalentFamiliesAreRuleFileFamilies(t *testing.T) {
	families := map[string]bool{}
	paths, err := filepath.Glob(filepath.Join(ruleFilesDir, "*.json"))
	if err != nil || len(paths) == 0 {
		t.Fatalf("no rule files: %v", err)
	}
	for _, p := range paths {
		body, err := os.ReadFile(p)
		if err != nil {
			t.Fatal(err)
		}
		r, err := ParseSpecRules(body)
		if err != nil {
			continue // not a rule file (a persona's JSON)
		}
		for _, list := range [][]RuleEntry{r.Core, r.Efficiency, r.Defensives} {
			for _, e := range list {
				families[e.Family] = true
			}
		}
	}
	for _, ct := range cooldownTalents {
		for _, fam := range ct.Families {
			if !families[fam] {
				t.Errorf("%s shortens %s, which no rule file judges", ct.Talent, fam)
			}
		}
	}
	for fam := range shardGates {
		if !families[fam] {
			t.Errorf("the shard gate %s is not a rule-file family", fam)
		}
	}
}

// The shard gate's level is the last level at which a shipped Destruction order fills
// Shadow and Flame (each order's point i is spent at level 10 + i).
func TestShardGateLevelIsWhereTheShippedOrdersFillTheTalent(t *testing.T) {
	g := shardGates["Shadowburn"]
	var at [3]int // the talent's 1-based tab, tier and column, as the orders write them
	found := false
	for _, tree := range readClassTalents(t, "warlock").Trees {
		for _, tal := range tree.Talents {
			if tal.Name == g.Talent {
				at, found = [3]int{tree.Position + 1, tal.Tier + 1, tal.Column + 1}, true
				if tal.MaxRank != g.Rank {
					t.Errorf("%s max rank %d, the gate wants %d", g.Talent, tal.MaxRank, g.Rank)
				}
			}
		}
	}
	if !found {
		t.Fatalf("no %s in the warlock talents", g.Talent)
	}
	var builds map[string]map[string]struct {
		Order [][3]int `json:"order"`
	}
	if err := json.Unmarshal(buildFile(t, "addon-builds.json"), &builds); err != nil {
		t.Fatal(err)
	}
	latest := int64(0)
	for kind, b := range builds["warlock-destruction"] {
		points, level := 0, int64(0)
		for i, p := range b.Order {
			if p == at {
				if points++; points == g.Rank {
					level = int64(10 + i)
				}
			}
		}
		if level == 0 {
			t.Errorf("the %s order never fills %s", kind, g.Talent)
		}
		latest = max(latest, level)
	}
	if latest != g.Level {
		t.Errorf("the shipped orders fill %s by %d, the gate says %d", g.Talent, latest, g.Level)
	}
}
