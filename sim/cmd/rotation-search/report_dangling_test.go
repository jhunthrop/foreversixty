package main

import "testing"

// TestDanglingGatesFlagsAGateNothingSatisfies covers the exact shape
// the greedy search can produce: a line gated on dotIsActive(X) while
// no other line in the same rotation ever casts X.
func TestDanglingGatesFlagsAGateNothingSatisfies(t *testing.T) {
	r := rotation{PriorityList: []entry{
		{Action: action{"castSpell": map[string]any{"spellId": buildActionID(100, 3)}, "condition": dotIsActiveCondition(200, 6)}},
		{Action: action{"castSpell": map[string]any{"spellId": buildActionID(300, 0)}}},
	}}
	got := danglingGates(r, map[int]string{100: "Lava Burst", 200: "Flame Shock"})
	if len(got) != 1 {
		t.Fatalf("got %v, want exactly one dangling gate", got)
	}
}

// TestDanglingGatesIgnoresAMaintenanceLinesOwnNotCondition proves the
// detector never flags a line's own not(dotIsActive(...)) maintenance
// condition as a dangling presence gate - that shape means "cast
// because it is NOT up", not "requires it to be up".
func TestDanglingGatesIgnoresAMaintenanceLinesOwnNotCondition(t *testing.T) {
	r := rotation{PriorityList: []entry{
		{Action: action{"castSpell": map[string]any{"spellId": buildActionID(200, 6)}, "condition": notCondition(dotIsActiveCondition(200, 6))}},
	}}
	if got := danglingGates(r, nil); len(got) != 0 {
		t.Fatalf("got %v, want none", got)
	}
}

// TestDanglingGatesIsSatisfiedWhenSomethingElseCastsTheGate covers the
// healthy case the search normally produces: the gated line survives
// alongside whatever applies its dot, so nothing is dangling.
func TestDanglingGatesIsSatisfiedWhenSomethingElseCastsTheGate(t *testing.T) {
	r := rotation{PriorityList: []entry{
		{Action: action{"castSpell": map[string]any{"spellId": buildActionID(200, 6)}, "condition": notCondition(dotIsActiveCondition(200, 6))}},
		{Action: action{"castSpell": map[string]any{"spellId": buildActionID(100, 3)}, "condition": dotIsActiveCondition(200, 6)}},
	}}
	if got := danglingGates(r, nil); len(got) != 0 {
		t.Fatalf("got %v, want none (Flame Shock is still cast above Lava Burst)", got)
	}
}

// TestDanglingGatesIgnoresAPassiveProcGate: the Fingers of Frost buff is
// granted by the talent, never cast, so a line gated on it is live.
func TestDanglingGatesIgnoresAPassiveProcGate(t *testing.T) {
	r := rotation{PriorityList: []entry{
		{Action: action{"castSpell": map[string]any{"spellId": buildActionID(1240047, 6)}, "condition": auraIsActiveCondition(fingersOfFrostAuraID, 0)}},
	}}
	if got := danglingGates(r, nil); len(got) != 0 {
		t.Fatalf("got %v, want no dangling gate", got)
	}
}
