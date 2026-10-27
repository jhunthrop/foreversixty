package guilds

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

func TestSignUpChangedNoteWrites(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	if _, err := w.svcSignUp(w.cmd(raid.ID, w.member, memberKeyA, SignupIn)); err != nil {
		t.Fatal(err)
	}
	note := "bringing flasks"
	cmd := w.cmd(raid.ID, w.member, memberKeyA, SignupIn)
	cmd.Note = &note
	res, err := w.svcSignUp(cmd)
	if err != nil || res.Unchanged {
		t.Fatalf("changed note: %+v %v", res, err)
	}
	if got := w.getRaid(t, raid.ID).Signups[0].Note; got != note {
		t.Fatalf("note = %q", got)
	}
}

func TestSignUpUnchangedStillRefusesALockedRaid(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	if _, err := w.svcSignUp(w.cmd(raid.ID, w.member, memberKeyA, SignupIn)); err != nil {
		t.Fatal(err)
	}
	if _, err := w.h.store.SetRaidStatus(t.Context(), w.gid, raid.ID, []string{RaidPlanned}, RaidLocked); err != nil {
		t.Fatal(err)
	}
	if _, err := w.svcSignUp(w.cmd(raid.ID, w.member, memberKeyA, SignupIn)); !errors.Is(err, ErrRaidLocked) {
		t.Fatalf("same row on a locked raid = %v, want ErrRaidLocked", err)
	}
}

// A second character's In benches the first; the first's repeat In must then write, never
// report Unchanged on the stale row. The concurrent run ends with exactly one attending row.
func TestSignUpUnchangedSeesABench(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	for _, key := range []string{memberKeyA, memberKeyB, memberKeyA} {
		res, err := w.svcSignUp(w.cmd(raid.ID, w.member, key, SignupIn))
		if err != nil || res.Unchanged || res.Raid.Counts.In != 1 {
			t.Fatalf("%s: %+v %v", key, res, err)
		}
	}
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		for _, key := range []string{memberKeyA, memberKeyB} {
			wg.Add(1)
			go func() {
				defer wg.Done()
				if _, err := w.svcSignUp(w.cmd(raid.ID, w.member, key, SignupIn)); err != nil {
					t.Error(err)
				}
			}()
		}
	}
	wg.Wait()
	if got := w.getRaid(t, raid.ID).Counts.In; got != 1 {
		t.Fatalf("attending after the race = %d, want 1", got)
	}
}

func TestSignUpOfficerForAMemberAndTheFreeze(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	res, err := w.svcSignUp(w.cmd(raid.ID, w.officer, memberKeyA, SignupIn))
	if err != nil || res.Raid.Counts.In != 1 {
		t.Fatalf("officer for a member: %+v %v", res, err)
	}
	w.setContested(t, true)
	if _, err := w.svcSignUp(w.cmd(raid.ID, w.officer, memberKeyA, SignupOut)); !errors.Is(err, ErrClaimContested) {
		t.Fatalf("frozen = %v", err)
	}
	// A moderator in the context is not the answering account: no bypass.
	ctx := auth.WithActor(t.Context(), auth.Actor{UserID: w.member2, Role: "moderator", Method: "session"})
	if _, err := w.h.svc.SignUp(ctx, w.cmd(raid.ID, w.officer, memberKeyA, SignupOut)); !errors.Is(err, ErrClaimContested) {
		t.Fatalf("foreign moderator ctx = %v", err)
	}
	w.setContested(t, false)
	w.dropPlan(t)
	if _, err := w.svcSignUp(w.cmd(raid.ID, w.officer, memberKeyA, SignupOut)); !errors.Is(err, ErrPlanRequired) {
		t.Fatalf("no plan = %v", err)
	}
}

func TestClearingAMatchMarksThePost(t *testing.T) {
	w, raid := matchedNight(t)
	marker := &fakeMarker{}
	w.h.svc.RaidPosts = marker
	w.as(w.officer)
	if res := w.h.do(http.MethodDelete, w.path("/raids/%d/match", raid.ID), ""); res.StatusCode != http.StatusOK {
		t.Fatalf("clear = %d: %s", res.StatusCode, readBody(res))
	}
	if marker.count() != 1 || marker.ids[0] != raid.ID {
		t.Fatalf("clear marks = %v", marker.ids)
	}
}

func TestGuildMoveClearMarksThePost(t *testing.T) {
	w, raid := matchedNight(t)
	marker := &fakeMarker{}
	w.h.svc.RaidPosts = marker
	other := seedGuild(t, w.h.pool, "Elsewhere")
	w.exec(t, fmt.Sprintf(`update reports set guild_id = %d where id = 'night'`, other))
	w.h.svc.ReportGuildMoved(t.Context(), "night", w.officer)
	if marker.count() != 1 || marker.ids[0] != raid.ID {
		t.Fatalf("guild-move marks = %v", marker.ids)
	}
}
