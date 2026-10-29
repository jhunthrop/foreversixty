package guilds

import (
	"context"
	"errors"
	"net/http"
	"testing"
)

func (w raidWorld) saveSpec(raidID, actor int64, key, spec string) (RaidSpecResult, error) {
	return w.h.svc.SaveRaidSpec(context.Background(), RaidSpecCommand{
		GuildID: w.gid, ActorID: actor, RaidID: raidID, CharacterKey: key, Spec: spec,
	})
}

// The raid spec wins over the logs: the candidate carries it, the sign-up takes its name and role.
func TestRaidSpecIsTheSignupsSpecAndRole(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	cands, err := w.h.svc.SignupCandidates(context.Background(), w.gid, w.member)
	if err != nil || len(cands) == 0 || cands[0].RaidSpec != "" {
		t.Fatalf("before = %+v err %v", cands, err)
	}
	res, err := w.saveSpec(0, w.member, memberKeyA, "warrior-protection")
	if err != nil || res.Character.Spec != "Protection" || res.Character.Role != RoleTank || res.Character.RaidSpec != "warrior-protection" {
		t.Fatalf("save = %+v err %v", res, err)
	}
	sg, err := w.svcSignUp(w.cmd(raid.ID, w.member, memberKeyA, SignupIn))
	if err != nil || sg.Character.Role != RoleTank || sg.Character.Spec != "Protection" {
		t.Fatalf("signup = %+v err %v", sg.Character, err)
	}
	signups, err := w.h.svc.Store.RaidSignups(context.Background(), raid.ID)
	if err != nil || len(signups) != 1 || signups[0].Spec != "Protection" || signups[0].Role != RoleTank {
		t.Fatalf("rows = %+v err %v", signups, err)
	}
}

// Change spec moves the row's spec and role, never its state; an Out row keeps its snapshot.
func TestRaidSpecChangeKeepsTheState(t *testing.T) {
	w := newRaidWorld(t, true)
	marker := &fakeMarker{}
	w.h.svc.RaidPosts = marker
	raid := w.createRaid(t)
	if _, err := w.svcSignUp(w.cmd(raid.ID, w.member, memberKeyA, SignupBench)); err != nil {
		t.Fatal(err)
	}
	marker.reset()
	res, err := w.saveSpec(raid.ID, w.member, memberKeyA, "warrior-fury")
	if err != nil || res.State != SignupBench || marker.count() != 1 {
		t.Fatalf("change = %+v err %v marks %d", res, err, marker.count())
	}
	signups, _ := w.h.svc.Store.RaidSignups(context.Background(), raid.ID)
	if signups[0].State != SignupBench || signups[0].Spec != "Fury" || signups[0].Role != RoleDPS {
		t.Fatalf("row = %+v", signups[0])
	}
	if _, err := w.svcSignUp(w.cmd(raid.ID, w.member, memberKeyA, SignupOut)); err != nil {
		t.Fatal(err)
	}
	res, err = w.saveSpec(raid.ID, w.member, memberKeyA, "warrior-protection")
	signups, _ = w.h.svc.Store.RaidSignups(context.Background(), raid.ID)
	if err != nil || res.State != SignupOut || signups[0].State != SignupOut || signups[0].Spec != "Fury" {
		t.Fatalf("out = %+v row %+v err %v", res, signups[0], err)
	}
}

func TestRaidSpecRefusesAnotherAccountAnotherClassAndPending(t *testing.T) {
	w := newRaidWorld(t, true)
	cases := []struct {
		actor     int64
		key, spec string
		want      error
	}{
		{w.member2, memberKeyA, "warrior-fury", ErrNotYourCharacter},
		{w.member, memberKeyA, "mage-frost", ErrSpecNotOfClass},
		{w.member, memberKeyA, "warrior-nope", ErrSpecNotOfClass},
		{w.member, memberKeyC, "rogue-combat", ErrCharacterNotVerified},
		{w.member, "us/hardcore/nobody", "rogue-combat", ErrCharacterNotFound},
	}
	for _, c := range cases {
		if _, err := w.saveSpec(0, c.actor, c.key, c.spec); !errors.Is(err, c.want) {
			t.Errorf("%s %s: err %v, want %v", c.key, c.spec, err, c.want)
		}
	}
}

// The raid page's PUT takes a spec, saves it and signs up with its role; the view carries it.
func TestPutSignupSavesTheSpec(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	w.as(w.member)
	res := w.putSignup(raid.ID, memberKeyA, `{"state":"in","role":"dps","spec":"warrior-protection"}`)
	if res.StatusCode != http.StatusOK {
		t.Fatalf("status %d: %s", res.StatusCode, readBody(res))
	}
	_ = res.Body.Close()
	view := w.getRaid(t, raid.ID)
	for _, c := range view.Viewer.Characters {
		if c.CharacterKey == memberKeyA && (c.RaidSpec == nil || *c.RaidSpec != "warrior-protection" || c.Role != RoleTank) {
			t.Fatalf("viewer character = %+v", c)
		}
	}
	bad := w.putSignup(raid.ID, memberKeyA, `{"state":"in","spec":"mage-frost"}`)
	expectError(t, bad, http.StatusUnprocessableEntity, "invalid")
	unknown := w.putSignup(raid.ID, memberKeyA, `{"state":"in","spec":"nope"}`)
	expectError(t, unknown, http.StatusUnprocessableEntity, "invalid")
}

// A locked raid's PUT with a spec changes nothing: no raid spec saved, the row's role unmoved.
func TestPutSignupWithASpecOnALockedRaidChangesNothing(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	if _, err := w.svcSignUp(w.cmd(raid.ID, w.member, memberKeyA, SignupIn)); err != nil {
		t.Fatal(err)
	}
	if _, err := w.h.pool.Exec(context.Background(), `update guild_raids set status = 'locked' where id = $1`, raid.ID); err != nil {
		t.Fatal(err)
	}
	w.as(w.member)
	res := w.putSignup(raid.ID, memberKeyA, `{"state":"in","spec":"warrior-protection"}`)
	expectError(t, res, http.StatusConflict, "raid_locked")
	var spec *string
	if err := w.h.pool.QueryRow(context.Background(), `select raid_spec from guild_characters where character_key = $1`, memberKeyA).
		Scan(&spec); err != nil || spec != nil {
		t.Fatalf("raid_spec = %v err %v", spec, err)
	}
	signups, err := w.h.svc.Store.RaidSignups(context.Background(), raid.ID)
	if err != nil || len(signups) != 1 || signups[0].Role != RoleDPS || signups[0].State != SignupIn {
		t.Fatalf("rows = %+v err %v", signups, err)
	}
}
