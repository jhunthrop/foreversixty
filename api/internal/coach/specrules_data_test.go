package coach

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"testing"
)

// The shipped rule files and the client tables they are stamped against, relative to this
// package.
const (
	ruleFilesDir    = "../../../data/curated/coach"
	ruleBuildsDir   = "../../../data/builds"
	activeBuildPath = "../../../web/src/data/active-build.json"
)

// ruleFileTodos is every list a rule file leaves out on purpose because nothing in the
// repo grounds it yet, by spec and list name. A left-out list sends its findings to
// "Couldn't check" (§6.1); the test reports each one and fails on any omission not named
// here, or on a name here whose list the file does carry.
var ruleFileTodos = map[string]map[string]string{
	"mage-fire": {
		"auras": "Improved Scorch's debuff (12873, the aura the APL stacks) is not in this build's mage spellconst, so its id cannot be stamped",
	},
}

// newestBuild is the active client build: web/src/data/active-build.json's build, never a
// sorted listing of data/builds (forever-prebeta sorts last).
func newestBuild(t *testing.T) string {
	t.Helper()
	data, err := os.ReadFile(activeBuildPath)
	if err != nil {
		t.Fatalf("read active build: %v", err)
	}
	var f struct {
		Build string `json:"build"`
	}
	if err := json.Unmarshal(data, &f); err != nil || f.Build == "" {
		t.Fatalf("active build %q: %v", f.Build, err)
	}
	return f.Build
}

// classTables is one class's spellconst, trainable family names and talent names at a build.
type classTables struct {
	spells    map[int64]SpellConst
	trainable map[string][]TrainableRank
	talents   map[string]bool
}

func loadClassTables(t *testing.T, build, class string) classTables {
	t.Helper()
	read := func(dir string) []byte {
		data, err := os.ReadFile(filepath.Join(ruleBuildsDir, build, dir, class+".json"))
		if err != nil {
			t.Fatalf("%s %s: %v", dir, class, err)
		}
		return data
	}
	spells, err := ParseSpellConst(read("spellconst"))
	if err != nil {
		t.Fatal(err)
	}
	trainable, err := ParseTrainables(read("trainables"))
	if err != nil {
		t.Fatal(err)
	}
	var tal struct {
		Trees []struct {
			Talents []struct {
				Name string `json:"name"`
			} `json:"talents"`
		} `json:"trees"`
	}
	if err := json.Unmarshal(read("talents"), &tal); err != nil {
		t.Fatalf("talents %s: %v", class, err)
	}
	talents := map[string]bool{}
	for _, tree := range tal.Trees {
		for _, x := range tree.Talents {
			talents[x.Name] = true
		}
	}
	return classTables{spells: spells, trainable: trainable, talents: talents}
}

// shippedRules is one spec's rule file with its specs.json row and class tables.
type shippedRules struct {
	info  specInfo
	rules SpecRules
	class classTables
}

// loadShippedRules loads the rule file of every spec in data/curated/specs.json.
func loadShippedRules(t *testing.T) (string, []shippedRules) {
	t.Helper()
	build := newestBuild(t)
	classes := map[string]classTables{}
	var out []shippedRules
	for _, s := range loadSpecs(t) {
		r, ok, err := LoadSpecRules(ruleFilesDir, s.Spec)
		if err != nil || !ok {
			t.Fatalf("%s: loaded %v, %v", s.Spec, ok, err)
		}
		if _, seen := classes[s.Class]; !seen {
			classes[s.Class] = loadClassTables(t, build, s.Class)
		}
		out = append(out, shippedRules{info: s, rules: r, class: classes[s.Class]})
	}
	return build, out
}

// allRuleSpellIDs is every spell id a rule file names.
func allRuleSpellIDs(r SpecRules) []int64 {
	var ids []int64
	for _, list := range [][]RuleEntry{r.Core, r.Downranks, r.Defensives, r.Efficiency} {
		for _, e := range list {
			ids = append(ids, e.SpellIDs...)
		}
	}
	for _, a := range r.Auras {
		ids = append(ids, a.IDs()...)
	}
	return ids
}

// missingRanks is every ranked (rank ≥ 1) spellconst spell named name that ids leaves out.
// A rank left out would read as the ability never cast or the aura never applied (§20.4).
// Unranked spells of the same name (NPC copies, triggered effects) are not the player's.
func missingRanks(spells map[int64]SpellConst, name string, ids []int64) []int64 {
	var out []int64
	for id, sc := range spells {
		if sc.Name == name && sc.Rank >= 1 && !slices.Contains(ids, id) {
			out = append(out, id)
		}
	}
	slices.Sort(out)
	return out
}

// checkRankList: every id is a spell named name, and no rank of name is missing.
func checkRankList(t *testing.T, s shippedRules, what, name string, ids []int64) {
	t.Helper()
	for _, id := range ids {
		if sc, ok := s.class.spells[id]; ok && sc.Name != name {
			t.Errorf("%s %s: %s spell %d is %q in spellconst", s.info.Spec, what, name, id, sc.Name)
		}
	}
	if miss := missingRanks(s.class.spells, name, ids); len(miss) > 0 {
		t.Errorf("%s %s: %s leaves out ranks %v", s.info.Spec, what, name, miss)
	}
}

// stampID checks a verified stamp is <build>/spellconst/<class>.json#<id> and returns the id.
func stampID(stamp, build, class string) (int64, error) {
	prefix := fmt.Sprintf("%s/spellconst/%s.json#", build, class)
	rest, ok := strings.CutPrefix(stamp, prefix)
	if !ok {
		return 0, fmt.Errorf("stamp %q is not against %s", stamp, prefix)
	}
	return strconv.ParseInt(rest, 10, 64)
}

func TestSpecRuleFiles_AllStampedAndReal(t *testing.T) {
	build, all := loadShippedRules(t)
	if len(all) != 28 {
		t.Fatalf("%d specs in specs.json, want 28", len(all))
	}
	for _, s := range all {
		r := s.rules
		if r.Spec != s.info.Spec || r.Build != build {
			t.Errorf("%s: header spec %q build %q, want build %s", s.info.Spec, r.Spec, r.Build, build)
		}
		for _, id := range allRuleSpellIDs(r) {
			if _, ok := s.class.spells[id]; !ok {
				t.Errorf("%s: spell %d is not in %s's %s spellconst", r.Spec, id, build, s.info.Class)
			}
		}
		lists := map[string][]RuleEntry{"core": r.Core, "downranks": r.Downranks, "defensives": r.Defensives, "efficiency": r.Efficiency}
		for list, entries := range lists {
			for _, e := range entries {
				checkShippedEntry(t, s, build, list, e)
			}
		}
		for _, a := range r.Auras {
			checkShippedAura(t, s, build, a)
		}
	}
}

// checkShippedEntry: every id is a rank of the family by name, the stamp names one of them
// at this build, the talent flag matches the class's talent trees, and a baseline
// defensive is trainable (F7 counts it from the trainable level).
func checkShippedEntry(t *testing.T, s shippedRules, build, list string, e RuleEntry) {
	t.Helper()
	spec := s.info.Spec
	checkRankList(t, s, list, e.Family, e.SpellIDs)
	id, err := stampID(e.Verified, build, s.info.Class)
	if err != nil || !slices.Contains(e.SpellIDs, id) {
		t.Errorf("%s %s: %s stamp %q (%v) names none of its ids", spec, list, e.Family, e.Verified, err)
	}
	if e.Talent != s.class.talents[e.Family] {
		t.Errorf("%s %s: %s talent=%v, the talent trees say %v", spec, list, e.Family, e.Talent, s.class.talents[e.Family])
	}
	if list == "defensives" && !e.Talent && len(s.class.trainable[e.Family]) == 0 {
		t.Errorf("%s: baseline defensive %s is not trainable", spec, e.Family)
	}
}

func checkShippedAura(t *testing.T, s shippedRules, build string, a AuraRule) {
	t.Helper()
	spec := s.info.Spec
	if len(a.SpellIDs) == 0 {
		t.Errorf("%s auras: %s lists no ranks; a lower rank would read as never applied", spec, a.Name)
	}
	checkRankList(t, s, "auras", a.Name, a.IDs())
	if id, err := stampID(a.Verified, build, s.info.Class); err != nil || id != a.SpellID {
		t.Errorf("%s auras: %s stamp %q (%v) is not its spell %d", spec, a.Name, a.Verified, err, a.SpellID)
	}
	if a.Talent != s.class.talents[a.Name] {
		t.Errorf("%s auras: %s talent=%v, the talent trees say %v", spec, a.Name, a.Talent, s.class.talents[a.Name])
	}
	if !slices.Contains(uptimeKindsByRole[s.info.Role], a.Kind) {
		t.Errorf("%s auras: %s is a %s aura, which F2 never judges for a %s", spec, a.Name, a.Kind, s.info.Role)
	}
}

func TestSpecRuleFiles_HealersHaveOnlyEfficiencyCore(t *testing.T) {
	_, all := loadShippedRules(t)
	healers := 0
	for _, s := range all {
		if s.info.Role == RoleHealer {
			healers++
			if len(s.rules.Core) > 0 {
				t.Errorf("%s: a healer's F1 reads efficiency cooldowns only, but core lists %d", s.info.Spec, len(s.rules.Core))
			}
			continue
		}
		if len(s.rules.Efficiency) > 0 {
			t.Errorf("%s: efficiency is for healers, but a %s lists %d", s.info.Spec, s.info.Role, len(s.rules.Efficiency))
		}
	}
	if healers != 5 {
		t.Errorf("%d healer specs, want 5", healers)
	}
}

func TestSpecRuleFiles_CoreCooldownsInRange(t *testing.T) {
	build, all := loadShippedRules(t)
	for _, s := range all {
		for _, e := range slices.Concat(s.rules.Core, s.rules.Efficiency) {
			id, err := stampID(e.Verified, build, s.info.Class)
			if err != nil {
				continue // reported by AllStampedAndReal
			}
			cd := s.class.spells[id].CooldownMS
			if cd < CoreCooldownMinMS || cd > CoreCooldownMaxMS {
				t.Errorf("%s: %s (%d) has a %d ms cooldown, outside F1's %d to %d", s.info.Spec, e.Family, id, cd, CoreCooldownMinMS, CoreCooldownMaxMS)
			}
		}
	}
}

// TestSpecRuleFiles_OmittedListsAreTodos reports every left-out list and fails on one the
// todo table does not name.
func TestSpecRuleFiles_OmittedListsAreTodos(t *testing.T) {
	_, all := loadShippedRules(t)
	todos := 0
	for _, s := range all {
		r := s.rules
		present := map[string]bool{
			"core": r.Core != nil, "auras": r.Auras != nil, "downranks": r.Downranks != nil,
			"defensives": r.Defensives != nil, "efficiency": r.Efficiency != nil,
		}
		for list, has := range present {
			note, isTodo := ruleFileTodos[r.Spec][list]
			switch {
			case !has && isTodo:
				todos++
				t.Logf("todo %s %s: %s", r.Spec, list, note)
			case !has:
				t.Errorf("%s: %s is left out with no todo note", r.Spec, list)
			case isTodo:
				t.Errorf("%s: %s has a todo note but the file carries the list", r.Spec, list)
			}
		}
	}
	t.Logf("%d rule files, %d todo notes", len(all), todos)
}
