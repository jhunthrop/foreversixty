package guilds

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"
)

type raidFixture struct {
	store   *Store
	guildID int64
	userID  int64
}

func newRaidFixture(t *testing.T) raidFixture {
	t.Helper()
	pool := testPool(t)
	return raidFixture{
		store:   &Store{Pool: pool},
		guildID: seedGuild(t, pool, "Raiders"),
		userID:  seedUser(t, pool, "raider@example.com"),
	}
}

func (f raidFixture) input(startsAt time.Time) RaidInput {
	return RaidInput{
		GuildID: f.guildID, Zone: "raid:onyxias-lair", Title: "Onyxia", StartsAt: startsAt,
		DurationMinutes: 180, CreatedBy: f.userID,
	}
}

func (f raidFixture) createRaid(t *testing.T, startsAt time.Time) Raid {
	t.Helper()
	r, err := f.store.CreateRaid(context.Background(), f.input(startsAt))
	if err != nil {
		t.Fatal(err)
	}
	return r
}

func (f raidFixture) signup(raidID int64, key, state string) Signup {
	return Signup{RaidID: raidID, CharacterKey: key, Name: key, Class: "warrior", Spec: "arms",
		Role: "dps", State: state, SignedBy: f.userID}
}

func TestCreateRaidRefusesTheThirteenth(t *testing.T) {
	f := newRaidFixture(t)
	ctx := context.Background()
	start := time.Now().Add(24 * time.Hour)
	var first Raid
	for i := 0; i < RaidCap; i++ {
		r := f.createRaid(t, start.Add(time.Duration(i)*time.Hour))
		if i == 0 {
			first = r
		}
	}
	if _, err := f.store.CreateRaid(ctx, f.input(start)); !errors.Is(err, ErrRaidCap) {
		t.Fatalf("13th raid: err = %v, want ErrRaidCap", err)
	}
	if _, err := f.store.SetRaidStatus(ctx, first.ID, RaidCancelled); err != nil {
		t.Fatal(err)
	}
	if _, err := f.store.CreateRaid(ctx, f.input(start)); err != nil {
		t.Fatalf("after a cancel: %v", err)
	}
}

func TestPlannedRaidsWindow(t *testing.T) {
	f := newRaidFixture(t)
	ctx := context.Background()
	now := time.Now()
	f.createRaid(t, now.Add(-10*time.Hour)) // ended 7 h ago (3 h long): gone
	live := f.createRaid(t, now.Add(-4*time.Hour))
	cancelledFuture := f.createRaid(t, now.Add(2*time.Hour))
	cancelledPast := f.createRaid(t, now.Add(-1*time.Hour))
	for _, id := range []int64{cancelledFuture.ID, cancelledPast.ID} {
		if _, err := f.store.SetRaidStatus(ctx, id, RaidCancelled); err != nil {
			t.Fatal(err)
		}
	}
	got, err := f.store.PlannedRaids(ctx, f.guildID, now)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || got[0].ID != live.ID || got[1].ID != cancelledFuture.ID {
		t.Fatalf("PlannedRaids = %+v, want [live, cancelledFuture]", got)
	}
}

func TestNextRaidSkipsEnded(t *testing.T) {
	f := newRaidFixture(t)
	ctx := context.Background()
	now := time.Now()
	if _, ok, err := f.store.NextRaid(ctx, f.guildID, now); err != nil || ok {
		t.Fatalf("empty guild: ok=%v err=%v", ok, err)
	}
	f.createRaid(t, now.Add(-5*time.Hour)) // ended
	running := f.createRaid(t, now.Add(-1*time.Hour))
	f.createRaid(t, now.Add(30*time.Hour))
	got, ok, err := f.store.NextRaid(ctx, f.guildID, now)
	if err != nil || !ok || got.ID != running.ID {
		t.Fatalf("NextRaid = %+v ok=%v err=%v, want the running raid", got, ok, err)
	}
}

func TestUpsertSignupMovesTheOtherOwnCharacterToBench(t *testing.T) {
	f := newRaidFixture(t)
	ctx := context.Background()
	raid := f.createRaid(t, time.Now().Add(24*time.Hour))
	if moved, err := f.store.UpsertSignup(ctx, f.signup(raid.ID, "first", "in")); err != nil || moved != nil {
		t.Fatalf("first in: moved=%v err=%v", moved, err)
	}
	moved, err := f.store.UpsertSignup(ctx, f.signup(raid.ID, "second", "late"))
	if err != nil || moved == nil || *moved != "first" {
		t.Fatalf("second late: moved=%v err=%v, want first", moved, err)
	}
	if moved, err := f.store.UpsertSignup(ctx, f.signup(raid.ID, "first", "out")); err != nil || moved != nil {
		t.Fatalf("out moves nothing: moved=%v err=%v", moved, err)
	}
	got, err := f.store.RaidSignups(ctx, raid.ID)
	if err != nil {
		t.Fatal(err)
	}
	states := map[string]string{}
	for _, s := range got {
		states[s.CharacterKey] = s.State
	}
	if states["first"] != "out" || states["second"] != "late" {
		t.Fatalf("states = %v", states)
	}
}

func TestUpsertSignupBenchesOnlyOwnCharacters(t *testing.T) {
	f := newRaidFixture(t)
	ctx := context.Background()
	raid := f.createRaid(t, time.Now().Add(24*time.Hour))
	other := seedUser(t, f.store.Pool, "other@example.com")
	theirs := f.signup(raid.ID, "theirs", "in")
	theirs.SignedBy = other
	if _, err := f.store.UpsertSignup(ctx, theirs); err != nil {
		t.Fatal(err)
	}
	if moved, err := f.store.UpsertSignup(ctx, f.signup(raid.ID, "mine", "in")); err != nil || moved != nil {
		t.Fatalf("moved=%v err=%v", moved, err)
	}
}

func TestUpsertSignupIsAtomicUnderRace(t *testing.T) {
	f := newRaidFixture(t)
	ctx := context.Background()
	raid := f.createRaid(t, time.Now().Add(24*time.Hour))
	var wg sync.WaitGroup
	for _, key := range []string{"alpha", "beta"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, err := f.store.UpsertSignup(ctx, f.signup(raid.ID, key, "in")); err != nil {
				t.Error(err)
			}
		}()
	}
	wg.Wait()
	got, err := f.store.RaidSignups(ctx, raid.ID)
	if err != nil {
		t.Fatal(err)
	}
	in := 0
	for _, s := range got {
		if s.State == "in" {
			in++
		}
	}
	if in != 1 {
		t.Fatalf("%d characters in, want exactly 1: %+v", in, got)
	}
}

func TestUpsertSignupRefusesLockedAndCancelled(t *testing.T) {
	f := newRaidFixture(t)
	ctx := context.Background()
	raid := f.createRaid(t, time.Now().Add(24*time.Hour))
	if _, err := f.store.SetRaidStatus(ctx, raid.ID, RaidLocked); err != nil {
		t.Fatal(err)
	}
	if _, err := f.store.UpsertSignup(ctx, f.signup(raid.ID, "a", "in")); !errors.Is(err, ErrRaidLocked) {
		t.Fatalf("locked: err = %v", err)
	}
	officer := f.signup(raid.ID, "a", "in")
	officer.AsOfficer = true
	if _, err := f.store.UpsertSignup(ctx, officer); err != nil {
		t.Fatalf("officer on locked: %v", err)
	}
	if _, err := f.store.SetRaidStatus(ctx, raid.ID, RaidCancelled); err != nil {
		t.Fatal(err)
	}
	if _, err := f.store.UpsertSignup(ctx, officer); !errors.Is(err, ErrRaidCancelled) {
		t.Fatalf("officer on cancelled: err = %v", err)
	}
	if _, err := f.store.UpsertSignup(ctx, f.signup(raid.ID, "a", "in")); !errors.Is(err, ErrRaidCancelled) {
		t.Fatalf("cancelled: err = %v", err)
	}
}

func TestRaidCrudAndTimezone(t *testing.T) {
	f := newRaidFixture(t)
	ctx := context.Background()
	raid := f.createRaid(t, time.Now().Add(24*time.Hour))
	in := f.input(raid.StartsAt)
	in.Title, in.Notes = "Renamed", "bring flasks"
	updated, err := f.store.UpdateRaid(ctx, raid.ID, in)
	if err != nil || updated.Title != "Renamed" || updated.Notes != "bring flasks" {
		t.Fatalf("UpdateRaid = %+v err=%v", updated, err)
	}
	got, ok, err := f.store.GetRaid(ctx, f.guildID, raid.ID)
	if err != nil || !ok || got.Title != "Renamed" {
		t.Fatalf("GetRaid = %+v ok=%v err=%v", got, ok, err)
	}
	if _, ok, _ := f.store.GetRaid(ctx, f.guildID+999, raid.ID); ok {
		t.Fatal("GetRaid crossed guilds")
	}
	if tz, err := f.store.GuildRaidTimezone(ctx, f.guildID); err != nil || tz != "" {
		t.Fatalf("default tz = %q err=%v", tz, err)
	}
	if err := f.store.SetGuildRaidTimezone(ctx, f.guildID, "America/Chicago"); err != nil {
		t.Fatal(err)
	}
	if tz, err := f.store.GuildRaidTimezone(ctx, f.guildID); err != nil || tz != "America/Chicago" {
		t.Fatalf("tz = %q err=%v", tz, err)
	}
}
