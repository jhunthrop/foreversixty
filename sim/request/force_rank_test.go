package request

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/internal/spellranks"
)

// Round 18 priest X4: a rank-swap probe that rewrote only the cast left the
// Shadow Word: Pain line's own dotIsActive watching the authored rank, so the
// swapped line recast on every global (rank 4 showed -29%). ForceRotationRank
// rewrites every reference to the authored id, conditions included, on every
// line.
func TestForceRotationRankRewritesEveryReferenceToTheAuthoredID(t *testing.T) {
	const swp = 10894 // Shadow Word: Pain rank 8
	rank4, ok := spellranks.SpellIDOfRank("priest", swp, 4, 60)
	if !ok || rank4 == swp {
		t.Fatalf("no Shadow Word: Pain rank 4 in the table: %d %v", rank4, ok)
	}
	raw := []byte(`{"priorityList":[
		{"action":{"condition":{"not":{"val":{"dotIsActive":{"spellId":{"spellId":10894}}}}},"castSpell":{"spellId":{"spellId":10894,"rank":8}}}},
		{"action":{"condition":{"cmp":{"op":"OpLt","lhs":{"dotRemainingTime":{"spellId":{"spellId":10894}}},"rhs":{"const":{"val":"2s"}}}},"castSpell":{"spellId":{"spellId":10947}}}}
	],"prepullActions":[{"action":{"castSpell":{"spellId":{"spellId":10894}}},"doAtValue":{"const":{"val":"-1s"}}}]}`)
	out, err := ForceRotationRank(raw, "priest", swp, 4, 60)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(out), "10894") {
		t.Errorf("a reference to the authored rank survived: %s", out)
	}
	var doc any
	if err := json.Unmarshal(out, &doc); err != nil {
		t.Fatal(err)
	}
	if got := strings.Count(string(out), `"spellId":`+itoa(rank4)); got != 4 {
		t.Errorf("want the cast, both conditions and the pre-pull cast on rank 4 (%d), got %d: %s", rank4, got, out)
	}
	if _, err := ForceRotationRank(raw, "priest", swp, 4, 10); err == nil {
		t.Error("a rank not learned at the level must fail loudly")
	}
}

func itoa(v int32) string {
	b, _ := json.Marshal(v)
	return string(b)
}
