package main

import (
	"testing"

	"github.com/jhunthrop/foreversixty/sim/internal/spellranks"
	"github.com/jhunthrop/foreversixty/sim/request"
)

// The search sims each candidate the way the published number is simmed: under
// the bare preset at 60 the rotation casts the bare-60 rank choices
// (data/curated/rank-choices.json: Shadow's Shadow Word: Death rank 2), and the
// request names the preset. The candidate itself keeps its authored ids, so the
// pasteable rotation in the report is never rewritten.
func TestCandidatesCastThePresetsRankChoices(t *testing.T) {
	const authored, spec = int32(1309636), "priest-shadow"
	rank2, ok := spellranks.SpellIDOfRank("priest", authored, 2, 60)
	if !ok || rank2 == authored {
		t.Fatalf("Shadow Word: Death rank 2 at 60 = %d, %v", rank2, ok)
	}
	r := rotation{Type: "TypeAPL", PriorityList: []entry{{Action: action{"castSpell": map[string]any{"spellId": map[string]any{"spellId": float64(authored)}}}}}}
	bare := simSetup{spec: specInfo{Spec: spec, ClassSlug: "priest"}, level: 60, seed: 1}

	apl, err := bare.engineRotation(r)
	if err != nil {
		t.Fatal(err)
	}
	if got := apl.GetPriorityList()[0].GetAction().GetCastSpell().GetSpellId().GetSpellId(); got != rank2 {
		t.Errorf("bare-60 candidate casts %d, want rank 2 (%d)", got, rank2)
	}
	if got := bare.request(10).Preset; got != request.BarePreset {
		t.Errorf("bare request preset = %q, want %q", got, request.BarePreset)
	}
	if got := r.PriorityList[0].Action["castSpell"].(map[string]any)["spellId"].(map[string]any)["spellId"]; got != float64(authored) {
		t.Errorf("the candidate was rewritten in place: %v", got)
	}

	raid := bare
	raid.preset = &request.ResolvedPreset{}
	apl, err = raid.engineRotation(r)
	if err != nil {
		t.Fatal(err)
	}
	if got := apl.GetPriorityList()[0].GetAction().GetCastSpell().GetSpellId().GetSpellId(); got != authored {
		t.Errorf("raid-60 candidate casts %d, want the authored %d", got, authored)
	}
	if got := raid.request(10).Preset; got != request.RaidPreset {
		t.Errorf("raid request preset = %q, want %q", got, request.RaidPreset)
	}
}
