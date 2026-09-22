// api/internal/dataaddon/aggregate_test.go
package dataaddon

import "testing"

func score(v float64) *float64 { return &v }

func TestAggregateCharacterAveragesOverallAndEachNonExcludedComponent(t *testing.T) {
	fights := []fightScore{
		{Overall: 80, Components: []componentScore{
			{Name: "output", Score: score(90)}, {Name: "survival", Score: score(70)},
			{Name: "mechanics", Score: score(85)}, {Name: "utility", Score: score(60)},
			{Name: "preparation", Score: score(95)}, {Name: "activity", Score: nil}, // excluded this fight
		}},
		{Overall: 82, Components: []componentScore{
			{Name: "output", Score: score(92)}, {Name: "survival", Score: score(74)},
			{Name: "mechanics", Score: score(83)}, {Name: "utility", Score: score(64)},
			{Name: "preparation", Score: score(93)}, {Name: "activity", Score: score(91)},
		}},
	}
	row, ok := aggregateCharacter(fights)
	if !ok {
		t.Fatal("aggregateCharacter reported no data for two fights")
	}
	if row.Rating != 81 {
		t.Errorf("rating = %d, want 81", row.Rating)
	}
	if row.Fights != 2 {
		t.Errorf("fights = %d, want 2", row.Fights)
	}
	want := map[string]int{"output": 91, "survival": 72, "mechanics": 84, "utility": 62, "preparation": 94, "activity": 91}
	for name, v := range want {
		if row.Components[name] != v {
			t.Errorf("component %s = %d, want %d", name, row.Components[name], v)
		}
	}
}

func TestAggregateCharacterOmitsAComponentWithNoNonExcludedFights(t *testing.T) {
	fights := []fightScore{
		{Overall: 88, Components: []componentScore{
			{Name: "output", Score: score(70)}, {Name: "survival", Score: nil},
			{Name: "mechanics", Score: score(70)}, {Name: "utility", Score: nil},
			{Name: "preparation", Score: nil}, {Name: "activity", Score: nil},
		}},
	}
	row, ok := aggregateCharacter(fights)
	if !ok {
		t.Fatal("aggregateCharacter reported no data for one fight")
	}
	if len(row.Components) != 2 {
		t.Fatalf("components = %v, want exactly output and mechanics", row.Components)
	}
	if row.Components["output"] != 70 || row.Components["mechanics"] != 70 {
		t.Errorf("components = %v", row.Components)
	}
	if _, present := row.Components["survival"]; present {
		t.Error("survival should be omitted, not written as any number")
	}
}

func TestAggregateCharacterReportsNoDataForZeroFights(t *testing.T) {
	_, ok := aggregateCharacter(nil)
	if ok {
		t.Error("aggregateCharacter should report ok=false for zero fights")
	}
}

func TestDecodeComponentsReadsExcludedAsNilScore(t *testing.T) {
	raw := []byte(`[{"name":"output","score":70,"excluded":false},{"name":"survival","score":null,"excluded":true,"reason":"wipe"}]`)
	got := decodeComponents(raw)
	if len(got) != 2 {
		t.Fatalf("decoded %d components, want 2", len(got))
	}
	if got[0].Name != "output" || got[0].Score == nil || *got[0].Score != 70 {
		t.Errorf("output = %+v", got[0])
	}
	if got[1].Name != "survival" || got[1].Score != nil {
		t.Errorf("survival = %+v, want a nil score", got[1])
	}
}

func TestDecodeComponentsReturnsNilForUnreadableJSON(t *testing.T) {
	if got := decodeComponents([]byte(`not json`)); got != nil {
		t.Errorf("decodeComponents(garbage) = %v, want nil", got)
	}
}
