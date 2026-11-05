package request

import (
	"fmt"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	engine "github.com/wowsims/classic/sim"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

// Round 23 (platform): the paladin pacing rules read mana income through
// manaIncomePerSecond, "currentMana + income x remainingTime covers what is
// left". The pacing shape goes through the same JSON rank rewrite and protojson
// parse an embedded APL does, and reads the character's real regen: a budget
// the regen can meet casts, one it cannot never does. Smite spends below it, so
// the bar has room for what comes in (income is net of overflow), and stops
// with enough left for Shadow Word: Pain once the regen ticks resume.
func TestALineGatedOnManaIncomeReadsTheCharactersRegen(t *testing.T) {
	const shadowWordPainID = 10894
	registerEngine.Do(engine.RegisterAll)
	casts := func(perSecond float64) int64 {
		t.Helper()
		raw := []byte(fmt.Sprintf(`{"priorityList":[{"action":{
			"condition":{"cmp":{"op":"OpGe",
				"lhs":{"math":{"op":"OpAdd","lhs":{"currentMana":{}},
					"rhs":{"math":{"op":"OpMul","lhs":{"manaIncomePerSecond":{"windowSeconds":10}},"rhs":{"remainingTime":{}}}}}},
				"rhs":{"math":{"op":"OpAdd","lhs":{"currentMana":{}},
					"rhs":{"math":{"op":"OpMul","lhs":{"const":{"val":"%g"}},"rhs":{"remainingTime":{}}}}}}}},
			"castSpell":{"spellId":{"spellId":10894}}}},
			{"action":{"condition":{"cmp":{"op":"OpGt","lhs":{"currentMana":{}},"rhs":{"const":{"val":"1000"}}}},
				"castSpell":{"spellId":{"spellId":10934}}}}]}`, perSecond)) // Shadow Word: Pain on the budget; Smite spends down to 1000
		rewritten, err := rewriteRotationRanks(raw, "priest", 60, nil)
		if err != nil {
			t.Fatal(err)
		}
		rot := &proto.APLRotation{}
		if err := protojson.Unmarshal(rewritten, rot); err != nil {
			t.Fatalf("parse: %v", err)
		}
		req := smokeRequest(t, "priest-shadow")
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
		return CastSpellCounts(player)[shadowWordPainID]
	}
	if got := casts(0.1); got == 0 {
		t.Error("a 0.1 mana per second budget never cast: the value read no regen")
	}
	if got := casts(100000); got != 0 {
		t.Errorf("a 100,000 mana per second budget cast %d times", got)
	}
}
