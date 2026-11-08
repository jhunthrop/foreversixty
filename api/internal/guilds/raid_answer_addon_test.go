package guilds

import (
	"context"
	"net/http"
	"testing"
	"time"
)

// answerFromGame answers for memberKeyA as the member from the addon panel at answeredAt.
func (w raidWorld) answerFromGame(t *testing.T, raidID int64, state string, answeredAt time.Time) AddonAnswerOutcome {
	t.Helper()
	out, err := w.h.svc.AnswerFromAddon(context.Background(), AddonRaidAnswer{
		UserID: w.member, CharacterKey: memberKeyA, RaidID: raidID, State: state, AnsweredAt: answeredAt,
	})
	if err != nil {
		t.Fatal(err)
	}
	return out
}

// An answer from the game applies when it is newer than the stored row, keeps the stored role,
// and is refused when the app, site or Discord answered later (addon-v2 contract 4).
func TestAnswerFromAddonAppliesOnlyANewerAnswer(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	w.as(w.member)
	if res := w.putSignup(raid.ID, memberKeyA, `{"state":"bench","role":"tank"}`); res.StatusCode != http.StatusOK {
		t.Fatalf("site answer = %d: %s", res.StatusCode, readBody(res))
	}
	if got := w.answerFromGame(t, raid.ID, SignupIn, time.Now().Add(-time.Hour)); got.Status != AddonAnswerRefused || got.Reason != AddonAnswerStale {
		t.Fatalf("an older answer: %+v", got)
	}
	if got := w.answerFromGame(t, raid.ID, SignupIn, time.Now().Add(time.Second)); got.Status != AddonAnswerAccepted {
		t.Fatalf("a newer answer: %+v", got)
	}
	signups, err := w.h.svc.Store.SignupsByRaid(context.Background(), []int64{raid.ID})
	if err != nil {
		t.Fatal(err)
	}
	var found bool
	for _, sg := range signups[raid.ID] {
		if sg.CharacterKey == memberKeyA {
			found = true
			if sg.State != SignupIn || sg.Role != RoleTank {
				t.Fatalf("stored %+v", sg)
			}
		}
	}
	if !found {
		t.Fatal("no row")
	}
}

// With no stored row the answer takes the character's own role; a late answer keeps its minutes.
func TestAnswerFromAddonFirstAnswer(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	m := int16(30)
	out, err := w.h.svc.AnswerFromAddon(context.Background(), AddonRaidAnswer{
		UserID: w.member, CharacterKey: memberKeyA, RaidID: raid.ID, State: SignupLate, LateMinutes: &m, AnsweredAt: time.Now(),
	})
	if err != nil || out.Status != AddonAnswerAccepted {
		t.Fatalf("first answer %+v %v", out, err)
	}
	rows, err := w.h.svc.Store.AddonRaidSignups(context.Background(), []string{memberKeyA}, time.Now())
	if err != nil || len(rows) != 1 {
		t.Fatalf("rows %+v %v", rows, err)
	}
	r := rows[0]
	if r.Answer != SignupLate || r.Role == "" || r.LateMinutes == nil || *r.LateMinutes != 30 || r.AnsweredAt == nil {
		t.Fatalf("row %+v", r)
	}
}

// Refusals: a cancelled raid, an unknown raid, a character of another account, a bad answer.
func TestAnswerFromAddonRefusals(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	cases := []struct {
		name   string
		in     AddonRaidAnswer
		reason string
	}{
		{"unknown raid", AddonRaidAnswer{UserID: w.member, CharacterKey: memberKeyA, RaidID: raid.ID + 999, State: SignupIn}, AddonAnswerNotFound},
		{"another's character", AddonRaidAnswer{UserID: w.member, CharacterKey: member2Key, RaidID: raid.ID, State: SignupIn}, AddonAnswerNotYours},
		{"bad answer", AddonRaidAnswer{UserID: w.member, CharacterKey: memberKeyA, RaidID: raid.ID, State: "maybe"}, AddonAnswerInvalid},
	}
	for _, c := range cases {
		c.in.AnsweredAt = time.Now()
		out, err := w.h.svc.AnswerFromAddon(context.Background(), c.in)
		if err != nil || out.Status != AddonAnswerRefused || out.Reason != c.reason {
			t.Errorf("%s: %+v %v", c.name, out, err)
		}
	}
	w.as(w.officer)
	if res := w.h.do(http.MethodPost, w.path("/raids/%d/cancel", raid.ID), ""); res.StatusCode != http.StatusOK {
		t.Fatalf("cancel = %d", res.StatusCode)
	}
	if got := w.answerFromGame(t, raid.ID, SignupIn, time.Now()); got.Status != AddonAnswerRefused || got.Reason != AddonAnswerClosed {
		t.Fatalf("cancelled: %+v", got)
	}
}

// The inbox's raid_signup rows: every open raid of the character's guild in the next 7 days,
// answered or not; none past the window, cancelled, or for a character with no guild.
func TestAddonRaidSignupsWindow(t *testing.T) {
	w := newRaidWorld(t, true)
	soon := w.createRaids(t, raidBody(time.Now().Add(2*24*time.Hour).UTC().Truncate(time.Minute).Format(time.RFC3339), 0))[0]
	w.createRaids(t, raidBody(time.Now().Add(9*24*time.Hour).UTC().Truncate(time.Minute).Format(time.RFC3339), 0))
	cancelled := w.createRaids(t, raidBody(time.Now().Add(3*24*time.Hour).UTC().Truncate(time.Minute).Format(time.RFC3339), 0))[0]
	if res := w.h.do(http.MethodPost, w.path("/raids/%d/cancel", cancelled.ID), ""); res.StatusCode != http.StatusOK {
		t.Fatalf("cancel = %d", res.StatusCode)
	}
	rows, err := w.h.svc.Store.AddonRaidSignups(context.Background(), []string{memberKeyA, "us/hardcore/nobody"}, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 1 || rows[0].RaidID != soon.ID || rows[0].Answer != AddonAnswerNone || rows[0].Role != "" ||
		rows[0].AnsweredAt != nil || rows[0].ZoneName != "Onyxia's Lair" || rows[0].CharacterKey != memberKeyA {
		t.Fatalf("rows %+v", rows)
	}
}
