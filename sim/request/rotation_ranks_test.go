package request

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
)

// The three hunter rotations that reference Serpent Sting's max rank
// (25295) and Multi-Shot (2643). Multi-Shot's own spellranks.json row is
// rank 0 (the reference table never lists this build's higher client
// ranks 14288/14289/14290/25294 at all), but it is NOT unranked in the
// engine's own sense: sim/hunter/multi_shot.go gates rank 1 (2643) on
// RequiredLevel 18. sim/internal/spellranks' singleTierLevelOverrides
// (rotation-accuracy program, 2026-09-28) carries this id specifically
// so HighestLearnedSpellID agrees with the engine below level 18,
// rather than treating 2643 as always-learned the way a genuinely
// unranked ability (Bloodrage, Judgement) is.
const hunterSpecForRankTests = "hunter-survival"

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
	got, err := rewriteRotationRanks(raw, "hunter", 38)
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
	got, err := rewriteRotationRanks(raw, "hunter", 18)
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
	got, err := rewriteRotationRanks(raw, "hunter", api.MaxLevel)
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
	viaRewrite, err := rewriteRotationRanks(raw, "hunter", api.MaxLevel)
	if err != nil {
		t.Fatal(err)
	}
	got, err := rotation(hunterSpecForRankTests, "hunter", api.MaxLevel)
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
	got, err := rewriteRotationRanks(raw, "hunter", 3)
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
