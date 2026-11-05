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
	budget := func(perSecond float64) string {
		return fmt.Sprintf(`{"cmp":{"op":"OpGe",
			"lhs":{"math":{"op":"OpAdd","lhs":{"currentMana":{}},
				"rhs":{"math":{"op":"OpMul","lhs":{"manaIncomePerSecond":{"windowSeconds":10}},"rhs":{"remainingTime":{}}}}}},
			"rhs":{"math":{"op":"OpAdd","lhs":{"currentMana":{}},
				"rhs":{"math":{"op":"OpMul","lhs":{"const":{"val":"%g"}},"rhs":{"remainingTime":{}}}}}}}}`, perSecond)
	}
	if got := shadowWordPainCastsGatedOn(t, budget(0.1)); got == 0 {
		t.Error("a 0.1 mana per second budget never cast: the value read no regen")
	}
	if got := shadowWordPainCastsGatedOn(t, budget(100000)); got != 0 {
		t.Errorf("a 100,000 mana per second budget cast %d times", got)
	}
}

// Round 24 (platform): Protection paladin's Consecration gate reads the rate the
// rotation spends mana (manaSpentPerSecond) beside what comes in. Smite spends
// every cast, so a spend floor it clears lets Shadow Word: Pain through and one it
// never reaches does not; "income keeps up with spend" parses and reads the real
// bar (Smite outspends a priest's regen, so it holds only some of the time).
func TestALineGatedOnManaSpentReadsTheCharactersSpend(t *testing.T) {
	spendAbove := func(perSecond float64) string {
		return fmt.Sprintf(`{"cmp":{"op":"OpGt","lhs":{"manaSpentPerSecond":{"windowSeconds":10}},"rhs":{"const":{"val":"%g"}}}}`, perSecond)
	}
	if got := shadowWordPainCastsGatedOn(t, spendAbove(1)); got == 0 {
		t.Error("a 1 mana per second spend floor never cast: the value read no spend")
	}
	if got := shadowWordPainCastsGatedOn(t, spendAbove(100000)); got != 0 {
		t.Errorf("a 100,000 mana per second spend floor cast %d times", got)
	}
	keepsUp := `{"cmp":{"op":"OpGe","lhs":{"manaIncomePerSecond":{"windowSeconds":10}},"rhs":{"manaSpentPerSecond":{"windowSeconds":10}}}}`
	shadowWordPainCastsGatedOn(t, keepsUp)
}

// shadowWordPainCastsGatedOn runs a raid-60 Shadow priest whose list is Shadow Word:
// Pain on the given condition, then Smite while above 1000 mana, through the same
// JSON rank rewrite and protojson parse an embedded APL takes, and counts the Pains.
func shadowWordPainCastsGatedOn(t *testing.T, condition string) int64 {
	t.Helper()
	const shadowWordPainID = 10894
	registerEngine.Do(engine.RegisterAll)
	raw := []byte(`{"priorityList":[{"action":{"condition":` + condition + `,
		"castSpell":{"spellId":{"spellId":10894}}}},
		{"action":{"condition":{"cmp":{"op":"OpGt","lhs":{"currentMana":{}},"rhs":{"const":{"val":"1000"}}}},
			"castSpell":{"spellId":{"spellId":10934}}}}]}`)
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
