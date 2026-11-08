package addon

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

// fakeRaidAnswers answers the raid_signup rows from a fixed list and records each answer,
// answering every one with the next outcome in line (accepted when the line is empty).
type fakeRaidAnswers struct {
	rows     []guilds.AddonRaidSignup
	err      error
	got      []guilds.AddonRaidAnswer
	outcomes []guilds.AddonAnswerOutcome
	fail     error
}

func (f *fakeRaidAnswers) AddonRaidSignups(context.Context, []string, time.Time) ([]guilds.AddonRaidSignup, error) {
	return f.rows, f.err
}

func (f *fakeRaidAnswers) AnswerFromAddon(_ context.Context, a guilds.AddonRaidAnswer) (guilds.AddonAnswerOutcome, error) {
	f.got = append(f.got, a)
	if f.fail != nil {
		return guilds.AddonAnswerOutcome{}, f.fail
	}
	if len(f.outcomes) == 0 {
		return guilds.AddonAnswerOutcome{Status: guilds.AddonAnswerAccepted}, nil
	}
	o := f.outcomes[0]
	f.outcomes = f.outcomes[1:]
	return o, nil
}

// The raid_signup fixture is a flat message that MessageEntry decodes strictly and writes back.
func TestRaidSignupFixtureGolden(t *testing.T) {
	raw := contractFixture(t, "inbox-raid-signup.json")
	var m MessageEntry
	contractStrict(t, raw, &m)
	contractRoundTrip(t, raw, m)
	if m.Type != messageRaidSignup || m.Answer != "late" || m.LateMinutes != 15 || m.AnsweredAt == "" {
		t.Fatalf("message = %+v", m)
	}
}

func TestRaidSignupMessages(t *testing.T) {
	at := time.Date(2026, 12, 10, 2, 0, 0, 0, time.UTC)
	answered := time.Date(2026, 12, 8, 19, 4, 0, 0, time.UTC)
	m := int16(15)
	src := &fakeRaidAnswers{rows: []guilds.AddonRaidSignup{
		{CharacterKey: contractCharacter, RaidID: 42, ZoneName: "Onyxia's Lair", Title: "Onyxia", StartsAt: at,
			Answer: "late", Role: "dps", LateMinutes: &m, AnsweredAt: &answered},
		{CharacterKey: contractCharacter, RaidID: 43, ZoneName: "Onyxia's Lair", Title: "Onyxia", StartsAt: at.Add(24 * time.Hour),
			Answer: guilds.AddonAnswerNone},
	}}
	s := &Service{Log: quietLog(), RaidAnswers: src}
	got := s.raidSignupMessages(context.Background(), []string{contractCharacter}, at)
	if len(got) != 2 {
		t.Fatalf("messages = %+v", got)
	}
	contractRoundTrip(t, contractFixture(t, "inbox-raid-signup.json"), got[0])
	raw, _ := json.Marshal(got[1])
	var fields map[string]any
	_ = json.Unmarshal(raw, &fields)
	for _, absent := range []string{"role", "late_minutes", "answered_at"} {
		if _, ok := fields[absent]; ok {
			t.Errorf("an unanswered raid carries %s: %s", absent, raw)
		}
	}
	if fields["answer"] != "none" {
		t.Errorf("answer %v", fields["answer"])
	}
	src.err = errors.New("down")
	if failed := s.raidSignupMessages(context.Background(), []string{contractCharacter}, at); failed != nil {
		t.Errorf("a failed read sent %+v", failed)
	}
}

// The outbox fixture's raid_answer entry is read, applied for the account's own character, and
// answered per entry.
func TestOutboxRaidAnswers(t *testing.T) {
	var wire struct {
		RaidAnswer json.RawMessage `json:"raid_answer"`
	}
	if err := json.Unmarshal(contractFixture(t, "outbox-raid-answer.json"), &wire); err != nil {
		t.Fatal(err)
	}
	entries := readRaidAnswers(wire.RaidAnswer)
	if len(entries) != 1 || entries[0].Character != "Obnoxious-Yell" || entries[0].RaidID != 42 ||
		entries[0].Answer != "in" || entries[0].AnsweredAt != 1796753040 {
		t.Fatalf("entries %+v", entries)
	}
	exports := []Export{{Region: "us", Ruleset: "normal", Name: "Obnoxious"}}
	src := &fakeRaidAnswers{outcomes: []guilds.AddonAnswerOutcome{{Status: guilds.AddonAnswerRefused, Reason: guilds.AddonAnswerStale}}}
	s := &Service{Log: quietLog(), RaidAnswers: src}
	late := 30
	in := []RaidAnswerEntry{
		{Character: "Obnoxious-Yell", RaidID: 42, Answer: "in", AnsweredAt: 1796753040},
		{Character: "Nobody-Yell", RaidID: 42, Answer: "in", AnsweredAt: 1796753040},
		{Character: "Obnoxious-Yell", RaidID: 43, Answer: "late", LateMinutes: &late, AnsweredAt: 1796753041},
		{Character: "Obnoxious-Yell", RaidID: 0, Answer: "in", AnsweredAt: 1796753040},
	}
	res, err := s.applyRaidAnswers(context.Background(), 7, exports, in)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"refused:stale", "refused:unknown_character", "accepted:", "refused:invalid"}
	for i, r := range res {
		if got := r.Status + ":" + r.Reason; got != want[i] {
			t.Errorf("entry %d: %s, want %s", i, got, want[i])
		}
		if r.Character != in[i].Character || r.RaidID != in[i].RaidID || r.AnsweredAt != in[i].AnsweredAt {
			t.Errorf("entry %d does not echo its entry: %+v", i, r)
		}
	}
	if len(src.got) != 2 || src.got[0].CharacterKey != "us/normal/obnoxious" || src.got[0].UserID != 7 ||
		!src.got[0].AnsweredAt.Equal(time.Unix(1796753040, 0)) || src.got[1].LateMinutes == nil || *src.got[1].LateMinutes != 30 {
		t.Fatalf("answers %+v", src.got)
	}
	// Without a source every entry waits: the companion keeps it and sends it again.
	none, err := (&Service{Log: quietLog()}).applyRaidAnswers(context.Background(), 7, exports, in[:1])
	if err != nil || len(none) != 1 || none[0].Status != RaidAnswerDeferred {
		t.Fatalf("no source: %+v %v", none, err)
	}
	src.fail = errors.New("db down")
	if _, err := s.applyRaidAnswers(context.Background(), 7, exports, in[:1]); err == nil {
		t.Fatal("a failed write was answered")
	}
}

// Past the cap, entries wait for the next upload.
func TestOutboxRaidAnswersCap(t *testing.T) {
	exports := []Export{{Region: "us", Ruleset: "normal", Name: "Obnoxious"}}
	in := make([]RaidAnswerEntry, maxRaidAnswers+2)
	for i := range in {
		in[i] = RaidAnswerEntry{Character: "Obnoxious-Yell", RaidID: int64(i + 1), Answer: "out", AnsweredAt: 1796753040}
	}
	res, err := (&Service{Log: quietLog(), RaidAnswers: &fakeRaidAnswers{}}).applyRaidAnswers(context.Background(), 7, exports, in)
	if err != nil || len(res) != len(in) || res[maxRaidAnswers].Status != RaidAnswerDeferred || res[0].Status != guilds.AddonAnswerAccepted {
		t.Fatalf("res %+v %v", res, err)
	}
}
