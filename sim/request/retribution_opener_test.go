package request

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

// Sim-accuracy round 20 (guide-accuracy round 17's "Sims requested" 1): on a
// fresh boss the opener is Seal of the Crusader before the pull, its
// Judgement on the first swing, then the damage Seal. With Unsealed Wrath
// that Seal is Seal of Command straight away (three swings store three
// Commands and the twist starts), as the pre-pull already does on a judged
// boss; without it, Seal of Righteousness. Raid 60, Human, guide raid build,
// frozen BiS, 10,000 iterations: 782.7/782.4 against 779.9/779.3 for the old
// card's Righteousness-first line (+0.36/+0.40% at seeds 1 and 50001).
func TestRetributionOpenerSplitsTheDamageSealByUnsealedWrath(t *testing.T) {
	raw, err := aplFS.ReadFile("apl/paladin-retribution.apl.json")
	if err != nil {
		t.Fatal(err)
	}
	var apl struct {
		PriorityList []json.RawMessage `json:"priorityList"`
	}
	if err := json.Unmarshal(raw, &apl); err != nil {
		t.Fatal(err)
	}
	type sealLine struct {
		known bool // Unsealed Wrath known
		cast  int32
	}
	var got []sealLine
	for _, item := range apl.PriorityList {
		var line struct {
			Action struct {
				Condition struct {
					And struct {
						Vals []map[string]json.RawMessage `json:"vals"`
					} `json:"and"`
				} `json:"condition"`
				CastSpell struct {
					SpellID struct {
						SpellID int32 `json:"spellId"`
					} `json:"spellId"`
				} `json:"castSpell"`
			} `json:"action"`
		}
		if err := json.Unmarshal(item, &line); err != nil || len(line.Action.Condition.And.Vals) != 2 {
			continue
		}
		first, second := line.Action.Condition.And.Vals[0], line.Action.Condition.And.Vals[1]
		if !strings.Contains(openerCompact(t, second["auraIsActive"]), `"spellId":20308`) {
			continue
		}
		cast := line.Action.CastSpell.SpellID.SpellID
		switch {
		case strings.Contains(openerCompact(t, first["auraIsKnown"]), `"spellId":1310735`):
			got = append(got, sealLine{true, cast})
		case strings.Contains(openerCompact(t, first["not"]), `"auraIsKnown":{"auraId":{"spellId":1310735}}`):
			got = append(got, sealLine{false, cast})
		}
	}
	want := []sealLine{{true, 20920}, {false, 20293}}
	if len(got) != len(want) || got[0] != want[0] || got[1] != want[1] {
		t.Errorf("the lines after the Crusader Judgement cast %v, want Seal of Command with Unsealed Wrath then Seal of Righteousness without it (%v)", got, want)
	}
}

// openerCompact is raw JSON without whitespace, or "" for nothing.
func openerCompact(t *testing.T, raw json.RawMessage) string {
	t.Helper()
	if raw == nil {
		return ""
	}
	var buf bytes.Buffer
	if err := json.Compact(&buf, raw); err != nil {
		t.Fatal(err)
	}
	return buf.String()
}
