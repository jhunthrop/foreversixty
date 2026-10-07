package main

import (
	"testing"

	"github.com/jhunthrop/foreversixty/sim/request"
)

// TestDefaultConditionChooser covers the four shapes the task's own
// vocabulary calls for: dot, execute-only, buff/form/seal, and plain
// (cooldown or nuke).
func TestDefaultConditionChooser(t *testing.T) {
	tests := []struct {
		name   string
		a      request.LearnedAbility
		kind   conditionKind
		hasKey string // a key the resulting condition must carry, "" for none
	}{
		{
			name: "dot gets not(dotIsActive)",
			a:    request.LearnedAbility{Name: "Flame Shock", IDs: []int{29228}, Rank: 6, IsDOT: true, GCDMS: 1500},
			kind: condDOT, hasKey: "not",
		},
		{
			name: "execute-phase nuke gets isExecutePhase",
			a:    request.LearnedAbility{Name: "Execute", IDs: []int{5308}, IsDamage: true, GCDMS: 1500},
			kind: condExecute, hasKey: "isExecutePhase",
		},
		{
			name: "a persistent non-damage toggle gets not(auraIsActive)",
			a:    request.LearnedAbility{Name: "Moonkin Form", IDs: []int{24858}, DurationMS: -1, IsDamage: false, GCDMS: 1500},
			kind: condBuffOrForm, hasKey: "not",
		},
		{
			name: "a plain damage spell gets no condition",
			a:    request.LearnedAbility{Name: "Lightning Bolt", IDs: []int{15208}, Rank: 10, IsDamage: true, GCDMS: 1500},
			kind: condPlain, hasKey: "",
		},
		{
			name: "Ice Lance waits for a Fingers of Frost charge",
			a:    request.LearnedAbility{Name: "Ice Lance", IDs: []int{1240047}, Rank: 6, IsDamage: true, GCDMS: 1500},
			kind: condChargeGated, hasKey: "auraIsActive",
		},
		{
			name: "a cooldown with no dot/toggle/execute shape gets no condition",
			a:    request.LearnedAbility{Name: "Stormstrike", IDs: []int{17364}, IsDamage: true, GCDMS: 1500},
			kind: condPlain, hasKey: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cond, kind := defaultCondition(tt.a)
			if kind != tt.kind {
				t.Errorf("kind = %v, want %v", kind, tt.kind)
			}
			if tt.hasKey == "" {
				if cond != nil {
					t.Errorf("condition = %v, want nil", cond)
				}
				return
			}
			if cond == nil {
				t.Fatalf("condition = nil, want a %q node", tt.hasKey)
			}
			if _, ok := cond[tt.hasKey]; !ok {
				t.Errorf("condition = %v, want top-level key %q", cond, tt.hasKey)
			}
		})
	}
}

// TestDefaultConditionNeverFiresOnAnUngatedEntry is the DOT branch's
// own point: the built condition must actually be FALSE while the dot
// is up, so a maintenance line built from it does not recast on top
// of an active dot - i.e. it must be a "not(...)" wrapper, not the
// bare dotIsActive the probe's damage-gate lines use.
func TestDefaultConditionDOTIsANotWrapper(t *testing.T) {
	cond, _ := defaultCondition(request.LearnedAbility{Name: "Moonfire", IDs: []int{9835}, Rank: 10, IsDOT: true, GCDMS: 1500})
	not, ok := cond["not"].(map[string]any)
	if !ok {
		t.Fatalf("condition = %v, want a not(...) wrapper", cond)
	}
	val, ok := not["val"].(map[string]any)
	if !ok {
		t.Fatalf("not.val = %v, want a map", not["val"])
	}
	if _, ok := val["dotIsActive"]; !ok {
		t.Errorf("not.val = %v, want dotIsActive inside it", val)
	}
}

func TestIsProbeWorthy(t *testing.T) {
	tests := []struct {
		name string
		a    request.LearnedAbility
		want bool
	}{
		{"a damage spell with a real GCD is probe-worthy", request.LearnedAbility{Name: "Lava Burst", IsDamage: true, GCDMS: 2500}, true},
		{"a persistent non-damage toggle is probe-worthy", request.LearnedAbility{Name: "Moonkin Form", DurationMS: -1, GCDMS: 1500}, true},
		{"a plain non-damage, non-toggle ability is not probe-worthy", request.LearnedAbility{Name: "Rank 1 of something harmless", GCDMS: 1500}, false},
		{"a zero-GCD entry is never probe-worthy, even if it 'deals damage'", request.LearnedAbility{Name: "Attack", IsDamage: true, GCDMS: 0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isProbeWorthy(tt.a); got != tt.want {
				t.Errorf("isProbeWorthy(%+v) = %v, want %v", tt.a, got, tt.want)
			}
		})
	}
}

// TestTargetDebuffCurseIsProbedWithATargetAuraGate is a damage-taken debuff
// the caster lands on the target (Curse of the Elements): not damage and not a
// toggle, but a learned line worth probing, kept up while the target lacks it.
func TestTargetDebuffCurseIsProbedWithATargetAuraGate(t *testing.T) {
	curse := request.LearnedAbility{Name: "Curse of the Elements", IDs: []int{1311680}, Rank: 4, DurationMS: 300000, RaisesDamageTaken: true, GCDMS: 1500}
	if !isProbeWorthy(curse) {
		t.Fatal("a damage-taken debuff with a GCD is not probe-worthy")
	}
	cond, kind := defaultCondition(curse)
	if kind != condDebuff {
		t.Errorf("kind = %v, want %v", kind, condDebuff)
	}
	_, id, found := findNotActive(cond)
	if found {
		t.Errorf("the refresh mutators would rewrite this gate (found %+v): it names a target aura and they build self-aura conditions", id)
	}
	not, _ := cond["not"].(map[string]any)
	val, _ := not["val"].(map[string]any)
	gate, _ := val["auraIsActive"].(map[string]any)
	source, _ := gate["sourceUnit"].(map[string]any)
	if source["type"] != "CurrentTarget" {
		t.Errorf("gate = %v, want an auraIsActive on the CurrentTarget", gate)
	}
	if got, ok := readActionID(gate["auraId"]); !ok || got.SpellID != 1311680 {
		t.Errorf("gate aura = %v, want spell 1311680", gate["auraId"])
	}
}

// The Fingers of Frost gate must name the buff's own id (400669), not
// the Ice Lance rank's: the aura, not the spell, is what is held.
func TestIceLanceGateNamesTheFingersOfFrostAura(t *testing.T) {
	cond, _ := defaultCondition(request.LearnedAbility{Name: "Ice Lance", IDs: []int{1240047}, Rank: 6, IsDamage: true, GCDMS: 1500})
	inner, ok := cond["auraIsActive"].(map[string]any)
	if !ok {
		t.Fatalf("condition = %v, want an auraIsActive node", cond)
	}
	id, _ := inner["auraId"].(map[string]any)
	if id["spellId"] != float64(fingersOfFrostAuraID) {
		t.Errorf("gate aura id = %v, want %d", id["spellId"], fingersOfFrostAuraID)
	}
}
