package request

import (
	"testing"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	engine "github.com/wowsims/classic/sim"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

// Round 19 (priest lane): a line gated on the character's level fires only in
// its range. Discipline's step-down to Heal rank 4 applies from 40 to 59 and
// waits on an owner ruling at 60; no ranked id can gate that, since the rank
// rewrite resolves a ranked id inside spellIsKnown to the highest learned rank.
// The line goes through the same JSON rank rewrite and protojson parse an
// embedded APL does.
func TestALineGatedOnCurrentLevelFiresOnlyInItsRange(t *testing.T) {
	registerEngine.Do(engine.RegisterAll)
	raw := []byte(`{"priorityList":[{"action":{
		"condition":{"and":{"vals":[
			{"cmp":{"op":"OpGe","lhs":{"currentLevel":{}},"rhs":{"const":{"val":"40"}}}},
			{"cmp":{"op":"OpLt","lhs":{"currentLevel":{}},"rhs":{"const":{"val":"60"}}}}
		]}},
		"castSpell":{"spellId":{"spellId":10894}}}}]}`) // Shadow Word: Pain; the rewrite casts the learned rank
	casts := func(level int) int64 {
		t.Helper()
		rewritten, err := rewriteRotationRanks(raw, "priest", level, nil)
		if err != nil {
			t.Fatal(err)
		}
		rot := &proto.APLRotation{}
		if err := protojson.Unmarshal(rewritten, rot); err != nil {
			t.Fatalf("level %d: %v", level, err)
		}
		req := smokeRequest(t, "priest-shadow")
		req.Character.Level = level
		req.Encounter.TargetLevel = 0
		built, err := BuildWith(req, Options{OpenIterations: true})
		if err != nil {
			t.Fatal(err)
		}
		built.Raid.Parties[0].Players[0].Rotation = rot
		res := core.RunRaidSim(built)
		if err := adapter.ResultError(res); err != nil {
			t.Fatal(err)
		}
		player, err := adapter.PlayerMetrics(res)
		if err != nil {
			t.Fatal(err)
		}
		// The rotation's one line is the only thing that casts a spell.
		var n int64
		for _, count := range CastSpellCounts(player) {
			n += count
		}
		return n
	}
	if got := casts(45); got == 0 {
		t.Error("the 40-59 line cast nothing at level 45")
	}
	if got := casts(60); got != 0 {
		t.Errorf("the 40-59 line cast %d times at level 60", got)
	}
}
