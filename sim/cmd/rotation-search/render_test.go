package main

import "testing"

func TestRenderConditionTable(t *testing.T) {
	names := map[int]string{100: "Flame Shock", 200: "Lava Burst"}
	dot := map[string]any{"dotIsActive": map[string]any{"spellId": buildActionID(100, 0)}}
	cases := []struct {
		name string
		node any
		want string
	}{
		{"scalar falls through to %v", 42, "42"},
		{"and joins parenthesised clauses", map[string]any{"and": map[string]any{"vals": []any{dot, dot}}},
			"(Flame Shock dot is active) and (Flame Shock dot is active)"},
		{"or uses or", map[string]any{"or": map[string]any{"vals": []any{dot, map[string]any{"isExecutePhase": map[string]any{}}}}},
			"(Flame Shock dot is active) or (execute phase)"},
		{"not wraps", map[string]any{"not": map[string]any{"val": dot}}, "not (Flame Shock dot is active)"},
		{"cmp renders operator words", map[string]any(cmpCondition("OpGe", action{"currentManaPercent": map[string]any{}}, "30%")),
			"currentManaPercent >= 30%"},
		{"const", map[string]any{"const": map[string]any{"val": "3s"}}, "3s"},
		{"dot remaining", map[string]any{"dotRemainingTime": map[string]any{"spellId": buildActionID(100, 2)}}, "Flame Shock (rank 2) dot remaining time"},
		{"aura active", map[string]any(auraIsActiveCondition(200, 0)), "Lava Burst aura is active"},
		{"aura remaining", map[string]any{"auraRemainingTime": map[string]any{"auraId": buildActionID(999, 0)}}, "spell 999 aura remaining time"},
		{"unreadable id", map[string]any{"dotIsActive": map[string]any{"spellId": "x"}}, "? dot is active"},
		{"execute", map[string]any{"isExecutePhase": map[string]any{}}, "execute phase"},
		{"targets", map[string]any{"numberTargets": map[string]any{}}, "number of targets"},
		{"rage key", map[string]any{"currentRagePercent": map[string]any{}}, "currentRagePercent"},
		{"unknown shape renders its key", map[string]any{"weird": 1}, "weird"},
		{"empty map", map[string]any{}, "(empty)"},
		{"empty and", map[string]any{"and": map[string]any{"vals": []any{}}}, "(empty)"},
	}
	for _, c := range cases {
		if got := renderCondition(c.node, names); got != c.want {
			t.Errorf("%s: got %q, want %q", c.name, got, c.want)
		}
	}
}

func TestOpWords(t *testing.T) {
	for op, want := range map[string]string{"OpEq": "==", "OpNe": "!=", "OpLt": "<", "OpLe": "<=", "OpGt": ">", "OpGe": ">=", "OpXx": "OpXx"} {
		if got := opWords(op); got != want {
			t.Errorf("opWords(%s) = %q, want %q", op, got, want)
		}
	}
}

func TestRenderEntry(t *testing.T) {
	names := map[int]string{100: "Flame Shock"}
	cases := []struct {
		name string
		e    entry
		want string
	}{
		{"cast without condition", buildCastEntry("", 100, 0, nil), "Flame Shock -- no condition"},
		{"cast with condition", buildCastEntry("", 100, 3, notCondition(dotIsActiveCondition(100, 0))),
			"Flame Shock (rank 3) -- condition: not (Flame Shock dot is active)"},
		{"non-cast action is named by key", entry{Action: action{"autocastOtherCooldowns": map[string]any{}}}, "(autocastOtherCooldowns) -- no condition"},
		{"empty action", entry{Action: action{}}, "(empty action) -- no condition"},
		{"unknown spell", buildCastEntry("", 7, 0, nil), "spell 7 -- no condition"},
	}
	for _, c := range cases {
		if got := renderEntry(c.e, names); got != c.want {
			t.Errorf("%s: got %q, want %q", c.name, got, c.want)
		}
	}
}
