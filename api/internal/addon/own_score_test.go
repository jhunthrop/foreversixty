package addon

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"
)

// fakeOwnScores answers a fixed record per character key.
type fakeOwnScores struct {
	records map[string]OwnScore
	err     error
}

func (f fakeOwnScores) OwnScore(_ context.Context, key string) (OwnScore, bool, error) {
	r, ok := f.records[key]
	return r, ok, f.err
}

func intp(n int) *int { return &n }

// The own_score fixture is a flat message that MessageEntry decodes strictly and writes back.
func TestOwnScoreFixtureGolden(t *testing.T) {
	raw := contractFixture(t, "inbox-own-score.json")
	var m MessageEntry
	contractStrict(t, raw, &m)
	contractRoundTrip(t, raw, m)
	if m.Type != messageOwnScore || m.Score == nil || *m.Score != 642 || m.Tier != "epic" || m.ScoreAsOf != "2026-11-07" {
		t.Fatalf("message = %+v", m)
	}
}

// One own_score per character of the account with a record; a part the record lacks is left out,
// and a character with none sends nothing.
func TestOwnScoreMessages(t *testing.T) {
	asOf := time.Date(2026, 11, 7, 21, 30, 0, 0, time.UTC)
	src := fakeOwnScores{records: map[string]OwnScore{
		contractCharacter: {Score: 642, Tier: "epic", AsOf: asOf, Fights: 18,
			Gear: intp(310), Experience: intp(92), Parses: intp(140), Attendance: intp(60), Discipline: intp(40)},
		"us/normal/alt": {Score: 120, Tier: "common", AsOf: asOf, Fights: 2, Gear: intp(120)},
	}}
	s := &Service{Log: quietLog(), OwnScores: src}
	got := s.ownScoreMessages(context.Background(), []string{"us/normal/alt", contractCharacter, "us/normal/nolog"})
	if len(got) != 2 {
		t.Fatalf("messages = %+v", got)
	}
	contractRoundTrip(t, contractFixture(t, "inbox-own-score.json"), got[1])
	raw, _ := json.Marshal(got[0])
	var fields map[string]any
	_ = json.Unmarshal(raw, &fields)
	for _, absent := range []string{"experience", "parses", "attendance", "discipline"} {
		if _, ok := fields[absent]; ok {
			t.Errorf("a part the record lacks went out: %s", raw)
		}
	}
	if fields["gear"] != float64(120) || fields["score"] != float64(120) {
		t.Errorf("alt = %s", raw)
	}
	if failed := (&Service{Log: quietLog(), OwnScores: fakeOwnScores{err: errors.New("down")}}).
		ownScoreMessages(context.Background(), []string{contractCharacter}); len(failed) != 0 {
		t.Errorf("a failed read sent %+v", failed)
	}
	if none := (&Service{Log: quietLog()}).ownScoreMessages(context.Background(), []string{contractCharacter}); none != nil {
		t.Errorf("no source sent %+v", none)
	}
}
