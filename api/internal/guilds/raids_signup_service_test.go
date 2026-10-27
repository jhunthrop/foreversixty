package guilds

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"sync"
	"testing"
	"time"
)

// fakeMarker records the raids the service marks dirty.
type fakeMarker struct {
	mu  sync.Mutex
	ids []int64
}

func (m *fakeMarker) MarkRaidDirty(_ context.Context, raidID int64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.ids = append(m.ids, raidID)
}

func (m *fakeMarker) count() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return len(m.ids)
}

func (m *fakeMarker) reset() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.ids = nil
}

func (w raidWorld) svcSignUp(cmd SignupCommand) (SignupResult, error) {
	return w.h.svc.SignUp(context.Background(), cmd)
}

func (w raidWorld) cmd(raidID, actor int64, key, state string) SignupCommand {
	return SignupCommand{GuildID: w.gid, RaidID: raidID, ActorID: actor, CharacterKey: key, State: state}
}

func TestSignUpIsThePutRoute(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	res, err := w.svcSignUp(w.cmd(raid.ID, w.member, memberKeyA, SignupIn))
	if err != nil {
		t.Fatal(err)
	}
	if res.Raid.Counts.In != 1 || res.Moved != nil || res.Character.Name != "Kraggor" || !res.Character.Verified {
		t.Fatalf("result = %+v", res)
	}
	res, err = w.svcSignUp(w.cmd(raid.ID, w.member, memberKeyB, SignupIn))
	if err != nil {
		t.Fatal(err)
	}
	if res.Moved == nil || *res.Moved != "Kraggor" || res.Raid.Counts.In != 1 {
		t.Fatalf("moved = %v counts = %+v", res.Moved, res.Raid.Counts)
	}
}

func (w raidWorld) signupUpdatedAt(t *testing.T, raidID int64, key string) time.Time {
	t.Helper()
	var at time.Time
	if err := w.h.pool.QueryRow(t.Context(),
		`select updated_at from guild_raid_signups where raid_id = $1 and character_key = $2`, raidID, key).Scan(&at); err != nil {
		t.Fatal(err)
	}
	return at
}

func TestSignUpUnchangedOnTheSameState(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	marker := &fakeMarker{}
	w.h.svc.RaidPosts = marker
	if _, err := w.svcSignUp(w.cmd(raid.ID, w.member, memberKeyA, SignupIn)); err != nil {
		t.Fatal(err)
	}
	before := w.signupUpdatedAt(t, raid.ID, memberKeyA)
	marker.reset()
	res, err := w.svcSignUp(w.cmd(raid.ID, w.member, memberKeyA, SignupIn))
	if err != nil {
		t.Fatal(err)
	}
	if !res.Unchanged || !w.signupUpdatedAt(t, raid.ID, memberKeyA).Equal(before) || marker.count() != 0 {
		t.Fatalf("unchanged = %v, marks = %d", res.Unchanged, marker.count())
	}
}

func TestSignUpKeepsTheNoteWhenNil(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	w.as(w.member)
	w.signup(t, raid.ID, memberKeyA, `{"state":"late","note":"late 15m"}`)
	if _, err := w.svcSignUp(w.cmd(raid.ID, w.member, memberKeyA, SignupIn)); err != nil {
		t.Fatal(err)
	}
	got := w.getRaid(t, raid.ID)
	if len(got.Signups) != 1 || got.Signups[0].State != SignupIn || got.Signups[0].Note != "late 15m" {
		t.Fatalf("signups = %+v", got.Signups)
	}
	empty := ""
	cmd := w.cmd(raid.ID, w.member, memberKeyA, SignupIn)
	cmd.Note = &empty
	if res, err := w.svcSignUp(cmd); err != nil || res.Unchanged {
		t.Fatalf("clearing the note: %+v %v", res, err)
	}
	if got := w.getRaid(t, raid.ID); got.Signups[0].Note != "" {
		t.Fatalf("note = %q, want cleared", got.Signups[0].Note)
	}
}

func TestSignUpRefusals(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	outsider := seedUser(t, w.h.pool, "outsider@example.com")
	if _, err := w.svcSignUp(w.cmd(raid.ID, w.member, memberKeyC, SignupIn)); !errors.Is(err, ErrCharacterNotVerified) {
		t.Fatalf("pending = %v", err)
	}
	if _, err := w.svcSignUp(w.cmd(raid.ID, outsider, memberKeyA, SignupIn)); !errors.Is(err, ErrNotMember) {
		t.Fatalf("non-member = %v", err)
	}
	if _, err := w.svcSignUp(w.cmd(raid.ID, w.member2, memberKeyA, SignupIn)); !errors.Is(err, ErrNotYourCharacter) {
		t.Fatalf("another's character = %v", err)
	}
	if _, err := w.svcSignUp(w.cmd(raid.ID, w.member, "us/hardcore/nobody", SignupIn)); !errors.Is(err, ErrCharacterNotFound) {
		t.Fatalf("unknown character = %v", err)
	}
	if _, err := w.svcSignUp(w.cmd(raid.ID+9999, w.member, memberKeyA, SignupIn)); !errors.Is(err, ErrNotFound) {
		t.Fatalf("unknown raid = %v", err)
	}
	for _, c := range []struct {
		status string
		from   []string
		want   error
	}{{RaidLocked, []string{RaidPlanned}, ErrRaidLocked}, {RaidCancelled, []string{RaidLocked}, ErrRaidCancelled}} {
		if _, err := w.h.store.SetRaidStatus(t.Context(), w.gid, raid.ID, c.from, c.status); err != nil {
			t.Fatal(err)
		}
		if _, err := w.svcSignUp(w.cmd(raid.ID, w.member, memberKeyA, SignupIn)); !errors.Is(err, c.want) {
			t.Fatalf("%s raid = %v, want %v", c.status, err, c.want)
		}
	}
}

func TestSignupLimitIsSharedBySiteAndService(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	w.as(w.member)
	for i := 0; i < signupsPerMinute; i++ {
		if res := w.putSignup(raid.ID, memberKeyA, `{"state":"in"}`); res.StatusCode != http.StatusOK {
			t.Fatalf("PUT #%d = %d: %s", i+1, res.StatusCode, readBody(res))
		}
	}
	if _, err := w.svcSignUp(w.cmd(raid.ID, w.member, memberKeyA, SignupIn)); !errors.Is(err, ErrSignupRateLimited) {
		t.Fatalf("SignUp after 30 PUTs = %v", err)
	}
	if res := w.putSignup(raid.ID, memberKeyA, `{"state":"in"}`); res.StatusCode != http.StatusTooManyRequests {
		t.Fatalf("31st PUT = %d", res.StatusCode)
	}
}

func TestSignupCandidatesMainFirst(t *testing.T) {
	w := newRaidWorld(t, true)
	w.exec(t, fmt.Sprintf(`update users set main_character_key = '%s' where id = %d`, memberKeyB, w.member))
	got, err := w.h.svc.SignupCandidates(t.Context(), w.gid, w.member)
	if err != nil {
		t.Fatal(err)
	}
	var names []string
	for _, c := range got {
		names = append(names, c.Name)
	}
	if !slices.Equal(names, []string{"Sunderfel", "Kraggor", "Unverifiedalt"}) {
		t.Fatalf("order = %v", names)
	}
	if !got[0].Main || got[1].Main || !got[0].Verified || got[2].Verified {
		t.Fatalf("flags = %+v", got)
	}
	if got[0].Role == "" || got[0].Class != "mage" {
		t.Fatalf("first = %+v", got[0])
	}
}

func TestEveryRaidChangeMarksThePost(t *testing.T) {
	w := newRaidWorld(t, true)
	marker := &fakeMarker{}
	w.h.svc.RaidPosts = marker
	raid := w.createRaid(t)
	if marker.count() != 0 {
		t.Fatalf("create marked %d times, want 0", marker.count())
	}
	expectMark := func(name string, do func() *http.Response) {
		t.Helper()
		marker.reset()
		if res := do(); res.StatusCode != http.StatusOK {
			t.Fatalf("%s = %d: %s", name, res.StatusCode, readBody(res))
		}
		if marker.count() != 1 || marker.ids[0] != raid.ID {
			t.Fatalf("%s marks = %v, want [%d]", name, marker.ids, raid.ID)
		}
	}
	w.as(w.member)
	expectMark("signup", func() *http.Response { return w.putSignup(raid.ID, memberKeyA, `{"state":"in"}`) })
	w.as(w.officer)
	expectMark("patch", func() *http.Response {
		return w.h.do(http.MethodPatch, w.path("/raids/%d", raid.ID), raidBody(futureStart(), 0))
	})
	for _, action := range []string{"lock", "unlock", "cancel"} {
		expectMark(action, func() *http.Response { return w.h.do(http.MethodPost, w.path("/raids/%d/%s", raid.ID, action), "") })
	}
}

func TestMatchMarksThePost(t *testing.T) {
	w, raid := attendanceNight(t)
	w.grantPlan(t)
	marker := &fakeMarker{}
	w.h.svc.RaidPosts = marker
	w.as(w.officer)
	res := w.h.do(http.MethodPut, w.path("/raids/%d/match", raid.ID), `{"report_id":"night"}`)
	if res.StatusCode != http.StatusOK {
		t.Fatalf("put match = %d: %s", res.StatusCode, readBody(res))
	}
	if marker.count() != 1 || marker.ids[0] != raid.ID {
		t.Fatalf("match marks = %v", marker.ids)
	}
}

func TestRaidEndedIsPastItsEndAndStillOpen(t *testing.T) {
	start := time.Date(2026, 11, 1, 20, 0, 0, 0, time.UTC)
	raid := Raid{StartsAt: start, DurationMinutes: 180, Status: RaidPlanned}
	if !raid.EndsAt().Equal(start.Add(3 * time.Hour)) {
		t.Fatalf("EndsAt = %v", raid.EndsAt())
	}
	after := raid.EndsAt().Add(time.Minute)
	if raid.Ended(raid.EndsAt().Add(-time.Minute)) || !raid.Ended(after) {
		t.Fatal("Ended must turn true only after the end")
	}
	for _, status := range []string{RaidDone, RaidCancelled} {
		raid.Status = status
		if raid.Ended(after) {
			t.Fatalf("a %s raid is not Ended", status)
		}
	}
}
