package rating

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/rankings"
)

func TestHistoryDTOWireShape(t *testing.T) {
	at := time.Date(2026, 12, 9, 21, 30, 0, 0, time.FixedZone("x", 3600))
	parse := 87.5
	h := History{
		AsOf: "2026-12-09",
		Score: []HistoryPoint{
			{At: at, Score: 812, Tier: "gold", ReportID: "r1", FightIndex: 3, Encounter: "Boss", EncounterID: 667,
				Difficulty: 3, Kill: true, Visibility: "public"},
			{At: at, Score: 820, Tier: "gold", ReportID: "r2", FightIndex: 1, Encounter: "Boss", EncounterID: 667,
				Difficulty: 3, ParsePercentile: &parse, Kill: false, Visibility: "private"},
		},
		Encounters: []HistoryEncounter{{Encounter: "Boss", EncounterID: 667, Difficulty: 3,
			Points: []HistoryParsePoint{{At: at, ParsePercentile: parse, ReportID: "r2", FightIndex: 1}}}},
		Gaps: []HistoryGap{{From: "2026-11-01", To: "2026-12-09", Reason: GapNoRatedFights}},
		Fights: []rankings.OwnerFight{{CharacterFight: rankings.CharacterFight{
			Encounter: "Boss", EncounterID: 667, Difficulty: 3, ReportID: "r2", FightIndex: 1, FoughtAt: at},
			Visibility: "private"}},
	}
	raw, err := json.Marshal(toHistoryDTO(h))
	if err != nil {
		t.Fatal(err)
	}
	var got map[string]any
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatal(err)
	}
	if got["as_of"] != "2026-12-09" {
		t.Fatalf("as_of = %v", got["as_of"])
	}
	score := got["score"].([]any)
	first, second := score[0].(map[string]any), score[1].(map[string]any)
	wantFirst := map[string]any{"at": "2026-12-09T20:30:00Z", "score": 812.0, "tier": "gold", "report_id": "r1",
		"fight_index": 3.0, "encounter": "Boss", "encounter_id": 667.0, "difficulty": 3.0, "kill": true,
		"visibility": "public"}
	if !reflect.DeepEqual(first, wantFirst) {
		t.Fatalf("score[0] = %v, want %v (no parse_percentile)", first, wantFirst)
	}
	if second["parse_percentile"] != 87.5 || second["visibility"] != "private" || second["kill"] != false {
		t.Fatalf("score[1] = %v", second)
	}
	wantEnc := map[string]any{"encounter": "Boss", "encounter_id": 667.0, "difficulty": 3.0,
		"points": []any{map[string]any{"at": "2026-12-09T20:30:00Z", "parse_percentile": 87.5,
			"report_id": "r2", "fight_index": 1.0}}}
	if enc := got["encounters"].([]any)[0]; !reflect.DeepEqual(enc, wantEnc) {
		t.Fatalf("encounters[0] = %v, want %v", enc, wantEnc)
	}
	wantGap := map[string]any{"from": "2026-11-01", "to": "2026-12-09", "reason": GapNoRatedFights}
	if !reflect.DeepEqual(got["gaps"].([]any)[0], wantGap) {
		t.Fatalf("gaps[0] = %v", got["gaps"])
	}
	fight := got["fights"].([]any)[0].(map[string]any)
	for _, k := range []string{"encounter", "encounter_id", "difficulty", "kill", "metric", "value", "duration_ms",
		"fought_at", "report_id", "fight_index", "execution_score", "visibility"} {
		if _, ok := fight[k]; !ok {
			t.Fatalf("fights[0] lacks %q: %v", k, fight)
		}
	}
	if fight["visibility"] != "private" {
		t.Fatalf("fights[0].visibility = %v", fight["visibility"])
	}
}
