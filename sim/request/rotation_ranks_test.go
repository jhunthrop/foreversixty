package request

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
)

// hunter-beast-mastery: one of the two hunter rotations left that still
// reference Serpent Sting's max rank (25295) and Multi-Shot (2643) --
// hunter-survival used to be the third until the rotation-accuracy
// program's melee rewrite (2026-09-28) dropped every ranged shot and
// Serpent Sting from it (they require a range this now-melee spec never
// occupies; see data/curated/apl/hunter-survival.json's own notes), so
// this file points at beast mastery instead (marksmanship dropped Serpent
// Sting on 2026-10-08 on the raid rotation search). Multi-Shot's own
// spellranks.json row is rank 0 (the reference table never lists this
// build's higher client ranks 14288/14289/14290/25294 at all), but it is
// NOT unranked in the engine's own sense: sim/hunter/multi_shot.go gates
// rank 1 (2643) on RequiredLevel 18. sim/internal/spellranks'
// singleTierLevelOverrides (rotation-accuracy program, 2026-09-28)
// carries this id specifically so LearnedSpellID agrees with the
// engine below level 18, rather than treating 2643 as always-learned the
// way a genuinely unranked ability (Bloodrage, Judgement) is.
const hunterSpecForRankTests = "hunter-beast-mastery"

func readHunterAPL(t *testing.T) []byte {
	t.Helper()
	b, err := aplFS.ReadFile("apl/" + hunterSpecForRankTests + ".apl.json")
	if err != nil {
		t.Fatal(err)
	}
	return b
}

// unmarshalGeneric is a JSON blob as a generic tree, for structural
// comparison independent of key order or byte-for-byte formatting.
func unmarshalGeneric(t *testing.T, b []byte) any {
	t.Helper()
	var v any
	if err := json.Unmarshal(b, &v); err != nil {
		t.Fatal(err)
	}
	return v
}

// A level-38 hunter has learned Serpent Sting through rank 5 (id
// 13552, learned at 34; rank 6 needs 42) and Multi-Shot (2643, learned
// at 18 - well below 38).
func TestRewriteRotationRanksResolvesAHunterAtLevel38(t *testing.T) {
	raw := readHunterAPL(t)
	got, err := rewriteRotationRanks(raw, "hunter", 38, nil)
	if err != nil {
		t.Fatal(err)
	}
	s := string(got)
	if strings.Contains(s, `"spellId":25295`) {
		t.Error("Serpent Sting's level-60 rank (25295) should have been rewritten away for a level-38 character")
	}
	if !strings.Contains(s, `"spellId":13552`) {
		t.Error(`Serpent Sting should resolve to its level-38 rank, 13552 (rank 5, learned at 34)`)
	}
	if !strings.Contains(s, `"spellId":2643`) {
		t.Error("Multi-Shot (2643, learned at 18) should still be present at level 38")
	}
}

// A level-18 hunter has Serpent Sting (rank 3, 13550) and Multi-Shot
// but not Aimed Shot (learned at 20). Every hunter rotation conditions
// Serpent Sting and Rapid Fire on Aimed Shot's cooldown, so dropping an
// action for ANY unlearned spell it named left an 18 with auto shots and
// Multi-Shot alone (a live level-18 export, 2026-09-28). Only an
// unlearned CAST drops an action; a condition naming one is left for
// the engine to nil out.
func TestRewriteRotationRanksKeepsAnActionWhoseConditionNamesAnUnlearnedSpell(t *testing.T) {
	raw := readHunterAPL(t)
	got, err := rewriteRotationRanks(raw, "hunter", 18, nil)
	if err != nil {
		t.Fatal(err)
	}
	s := string(got)
	if !strings.Contains(s, `"spellId":13550`) {
		t.Error("Serpent Sting should be kept at its level-18 rank (13550) although its condition names Aimed Shot")
	}
	if !strings.Contains(s, `"spellId":2643`) {
		t.Error("Multi-Shot (2643, learned at 18) should be kept")
	}
	if strings.Contains(s, `"castSpell":{"spellId":{"spellId":20904`) {
		t.Error("the Aimed Shot cast itself (20904, learned at 20) should have been dropped")
	}
}

// At MaxLevel the rewrite is a no-op: every APL is authored against the
// highest rank of everything it casts, so resolving "the highest rank
// learned by level 60" always returns the id already there.
func TestRewriteRotationRanksAtMaxLevelIsUnchanged(t *testing.T) {
	raw := readHunterAPL(t)
	got, err := rewriteRotationRanks(raw, "hunter", api.MaxLevel, nil)
	if err != nil {
		t.Fatal(err)
	}
	want := unmarshalGeneric(t, raw)
	have := unmarshalGeneric(t, got)
	if !reflect.DeepEqual(want, have) {
		t.Errorf("rewriteRotationRanks at MaxLevel changed the rotation:\n got %s\nwant %s", got, raw)
	}
}

// rotation() itself skips the rewrite entirely at MaxLevel (an
// optimization and a stronger guarantee than "rewrites to the same
// JSON"): the parsed proto is built from the untouched embed, so a
// level-MaxLevel rotation is exactly what it always was.
func TestRotationAtMaxLevelParsesTheEmbedDirectly(t *testing.T) {
	raw := readHunterAPL(t)
	viaRewrite, err := rewriteRotationRanks(raw, "hunter", api.MaxLevel, nil)
	if err != nil {
		t.Fatal(err)
	}
	got, err := rotation(hunterSpecForRankTests, "hunter", RaidPreset, HordeFaction, api.MaxLevel)
	if err != nil {
		t.Fatal(err)
	}
	// The rewrite is a JSON no-op (proved above); parsing either the raw
	// embed or its rewritten twin must agree with what rotation() built.
	direct := unmarshalGeneric(t, raw)
	rewritten := unmarshalGeneric(t, viaRewrite)
	if !reflect.DeepEqual(direct, rewritten) {
		t.Fatal("test setup: the rewrite was not a no-op at MaxLevel")
	}
	if got == nil {
		t.Fatal("rotation() returned a nil rotation at MaxLevel")
	}
}

// A level-3 hunter has not learned Serpent Sting at all (its first rank
// needs level 4): every action that casts it or conditions on it must
// be dropped, not left pointing at an id nothing can cast. Multi-Shot
// (learned at 18) is dropped for the same reason - rotation-accuracy
// program (2026-09-28): this used to assert the opposite ("Multi-Shot
// is unranked and always available regardless of level"), which was the
// bug sim/internal/spellranks' singleTierLevelOverrides fixes; see that
// package's own comment and this file's hunterSpecForRankTests comment.
func TestRewriteRotationRanksDropsAnUnlearnedRankedSpell(t *testing.T) {
	raw := readHunterAPL(t)
	got, err := rewriteRotationRanks(raw, "hunter", 3, nil)
	if err != nil {
		t.Fatal(err)
	}
	s := string(got)
	for _, id := range []string{
		"1978", "425728", "13549", "425729", "13550", "425730",
		"13551", "425732", "13552", "425733", "13553", "425734",
		"13554", "425735", "13555", "425736", "25295", "425737",
		"2643",
	} {
		if strings.Contains(s, `"spellId":`+id) {
			t.Errorf("id %s survived the rewrite at level 3, where nothing has learned it", id)
		}
	}
}

// The rewrite is JSON that protojson can still parse: dropping an
// action never leaves the document malformed, and Build() itself
// succeeds for a low-level character even though some of its
// rotation's actions were dropped.
func TestBuildSucceedsForALowLevelHunterWithDroppedActions(t *testing.T) {
	req := api.SimRequest{
		EngineVersion: enginever.Version,
		Spec:          hunterSpecForRankTests,
		Source:        api.CharacterSource{Kind: api.SourceManual},
		Character: api.CharacterSpec{
			Name:  "Lowbie",
			Race:  "orc",
			Class: "hunter",
			Level: 3,
		},
		Encounter:  api.DefaultEncounter(),
		Iterations: 3000,
		RandomSeed: 1,
	}
	if _, err := Build(req); err != nil {
		t.Fatalf("Build rejected a low-level hunter: %v", err)
	}
}

// rewriteRankedSpellIDs is the walker BuildRotationRanks depends on: it
// must rewrite an ActionID it finds directly, recurse through whatever
// field happens to hold one, and signal a drop without panicking on
// the surrounding structure.
func TestRewriteRankedSpellIDsWalksGenerically(t *testing.T) {
	resolve := func(id int32) (int32, bool) {
		switch id {
		case 100:
			return 100, true // untouched: not a ranked spell
		case 200:
			return 150, true // downgraded to a lower learned rank
		case 300:
			return 0, false // no rank learned: drop
		default:
			t.Fatalf("resolve called with unexpected id %d", id)
			return 0, false
		}
	}

	t.Run("a bare ActionID rewrites in place", func(t *testing.T) {
		node := map[string]any{"spellId": float64(200), "rank": float64(3)}
		drop := false
		rewriteRankedSpellIDs(node, resolve, &drop, true)
		if drop {
			t.Fatal("unexpected drop")
		}
		if node["spellId"] != float64(150) {
			t.Errorf("spellId = %v, want 150", node["spellId"])
		}
	})

	t.Run("an ActionID nested under an arbitrary key is found", func(t *testing.T) {
		node := map[string]any{
			"castSpell": map[string]any{
				"spellId": map[string]any{"spellId": float64(300)},
			},
		}
		drop := false
		rewriteRankedSpellIDs(node, resolve, &drop, true)
		if !drop {
			t.Fatal("expected drop for an unlearned ranked spell")
		}
	})

	t.Run("a list of actions is walked element by element", func(t *testing.T) {
		node := []any{
			map[string]any{"spellId": float64(100)},
			map[string]any{"spellId": float64(200)},
		}
		drop := false
		rewriteRankedSpellIDs(node, resolve, &drop, true)
		if drop {
			t.Fatal("unexpected drop")
		}
		if node[0].(map[string]any)["spellId"] != float64(100) {
			t.Error("an untouched id should stay unchanged")
		}
		if node[1].(map[string]any)["spellId"] != float64(150) {
			t.Error("a downgraded id should be rewritten")
		}
	})

	t.Run("an unrelated spellId-shaped scalar is not mistaken for one", func(t *testing.T) {
		// "val" is a string operand (proto.APLValueConst), never a
		// number, so it must never reach resolve.
		node := map[string]any{"const": map[string]any{"val": "1.5"}}
		drop := false
		rewriteRankedSpellIDs(node, resolve, &drop, true)
		if drop {
			t.Fatal("unexpected drop from a node with no spellId at all")
		}
	})
}

func TestFilterRankedActionsDropsWholeTopLevelEntries(t *testing.T) {
	resolve := func(id int32) (int32, bool) {
		if id == 300 {
			return 0, false
		}
		return id, true
	}
	entries := []any{
		map[string]any{"action": map[string]any{
			"castSpell": map[string]any{"spellId": map[string]any{"spellId": float64(100)}},
		}},
		map[string]any{"action": map[string]any{
			"castSpell": map[string]any{"spellId": map[string]any{"spellId": float64(300)}},
		}},
		// A sequence with one droppable step: the WHOLE entry goes, not
		// just the one step.
		map[string]any{"action": map[string]any{
			"strictSequence": map[string]any{
				"actions": []any{
					map[string]any{"castSpell": map[string]any{"spellId": map[string]any{"spellId": float64(100)}}},
					map[string]any{"castSpell": map[string]any{"spellId": map[string]any{"spellId": float64(300)}}},
				},
			},
		}},
	}
	got := filterRankedActions(entries, resolve)
	arr, ok := got.([]any)
	if !ok {
		t.Fatalf("filterRankedActions did not return a slice: %T", got)
	}
	if len(arr) != 1 {
		t.Fatalf("filterRankedActions kept %d entries, want 1 (only the first)", len(arr))
	}

	// A non-array value passes through unchanged.
	if v := filterRankedActions("not an array", resolve); v != "not an array" {
		t.Errorf("filterRankedActions altered a non-array value: %v", v)
	}
}

// rotationIDs is every spellId number in a rotation rewritten with choices.
func rotationIDs(t *testing.T, spec, class string, level int, choices rankChoices) map[int32]bool {
	t.Helper()
	raw, err := aplFS.ReadFile("apl/" + spec + ".apl.json")
	if err != nil {
		t.Fatal(err)
	}
	got, err := rewriteRotationRanks(raw, class, level, choices)
	if err != nil {
		t.Fatal(err)
	}
	ids := map[int32]bool{}
	var walk func(any)
	walk = func(node any) {
		switch v := node.(type) {
		case map[string]any:
			if n, ok := v["spellId"].(float64); ok {
				ids[int32(n)] = true
			}
			for _, val := range v {
				walk(val)
			}
		case []any:
			for _, val := range v {
				walk(val)
			}
		}
	}
	walk(unmarshalGeneric(t, got))
	return ids
}

// Owner ruling 3 (round 9): a deliberate band-60 downrank is respected below
// 60, never lifted to the top learned rank. Holy and Discipline cast Renew
// rank 3 (6075, learned at 20) and Discipline casts Power Word: Shield rank 4
// (3747, learned at 24); the rewrite used to lift them to Renew rank 8
// (10928) and Shield rank 8 (10899) at band 50. Below the level the authored
// rank is learned at, the highest rank under it stands in.
func TestRewriteKeepsTheAuthoredDownranksBelowSixty(t *testing.T) {
	cases := []struct {
		spec, class string
		level       int
		want        []int32
		lifted      []int32
	}{
		{"priest-holy", "priest", 50, []int32{6075}, []int32{10928}},
		{"priest-holy", "priest", 30, []int32{6075}, []int32{6076}},
		{"priest-holy", "priest", 16, []int32{6074}, []int32{6075}},
		{"priest-discipline", "priest", 50, []int32{6075, 3747}, []int32{10928, 10899}},
		{"priest-discipline", "priest", 40, []int32{3747}, []int32{6066}},
		{"priest-discipline", "priest", 20, []int32{600}, []int32{3747}},
		// The party's Flash of Light rank 4 and the tank's rank 1 both survive.
		{"paladin-holy", "paladin", 50, []int32{19941, 19750}, []int32{19942}},
	}
	for _, c := range cases {
		ids := rotationIDs(t, c.spec, c.class, c.level, nil)
		for _, id := range c.want {
			if !ids[id] {
				t.Errorf("%s at %d: the authored downrank should cast %d", c.spec, c.level, id)
			}
		}
		for _, id := range c.lifted {
			if ids[id] {
				t.Errorf("%s at %d: %d is a lift above the authored rank", c.spec, c.level, id)
			}
		}
	}
}

// A per-band rank choice is honoured over its levels, and only there:
// outside every row the authored rank's rule stands.
func TestRewriteHonoursAPerBandRankChoice(t *testing.T) {
	// One authored line, so the check reads the choice alone: the real
	// Elemental list also carries a rank 10 line paced to the mana bar,
	// which no choice row names.
	raw := []byte(`{"type":"TypeAPL","priorityList":[{"action":{"castSpell":{"spellId":{"spellId":15207,"rank":9}}}}]}`)
	const topBolt = 15207 // Lightning Bolt rank 9, as the line is authored
	choices := rankChoices{topBolt: {{FromLevel: 40, ToLevel: 49, Rank: 6}, {FromLevel: 50, ToLevel: 55, Rank: 7}}}
	for _, c := range []struct {
		level      int
		want, lost int32
	}{
		{38, 10391, 0},     // no row yet: rank 7, the top learned at 38
		{40, 6041, 10391},  // rank 6 from 40
		{49, 6041, 10392},  // still rank 6 at 49
		{50, 10391, 15207}, // rank 7 from 50
		{55, 10391, 10392}, // through 55
		{56, 15207, 10391}, // past the row: the authored rank 9 again
	} {
		got, err := rewriteRotationRanks(raw, "shaman", c.level, choices)
		if err != nil {
			t.Fatal(err)
		}
		s := string(got)
		if !strings.Contains(s, `"spellId":`+strconv.Itoa(int(c.want))) {
			t.Errorf("level %d: want Lightning Bolt %d", c.level, c.want)
		}
		if c.lost != 0 && strings.Contains(s, `"spellId":`+strconv.Itoa(int(c.lost))) {
			t.Errorf("level %d: %d should have been replaced by the band's choice", c.level, c.lost)
		}
	}
}

// The curated table reaches the published rotation: Elemental's band-50 row
// casts Lightning Bolt rank 4 (915) where the default rule casts rank 9. The
// list's other Lightning Bolt line, the top rank while mana is ahead of the
// fight (15208), has its own row since round 12 and casts rank 7 (10391).
func TestThePublishedRotationTakesTheCuratedRankChoices(t *testing.T) {
	ids := rotationIDs(t, "shaman-elemental", "shaman", 50, specRankChoices("shaman-elemental", BarePreset, HordeFaction))
	if !ids[915] {
		t.Errorf("Elemental at 50 should cast Lightning Bolt rank 4 (915) on its filler line")
	}
	if !ids[10391] || ids[15207] {
		t.Errorf("Elemental at 50 should cast rank 7 (10391) on its paced line, not rank 9 (15207)")
	}
}
