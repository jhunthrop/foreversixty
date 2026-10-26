package guilds

import (
	"testing"
	"time"
)

// seedAward stores an Onyxia award of onyxiaLoot[item] to key at awardedAt, under raidID (nil
// for none), and returns its id.
func (w raidWorld) seedAward(t *testing.T, item int, key string, awardedAt time.Time, raidID *int64) int64 {
	t.Helper()
	var id int64
	if err := w.h.pool.QueryRow(t.Context(), `
		insert into loot_awards (guild_id, encounter_id, item_id, character_key, awarded_by, awarded_at, raid_id)
		values ($1, $2, $3, $4, $5, $6, $7) returning id`,
		w.gid, onyxiaEncounterID, onyxiaLoot[item].ItemID, key, w.officer, awardedAt, raidID).Scan(&id); err != nil {
		t.Fatal(err)
	}
	return id
}

// awardRaid reads an award's raid_id and attached_by_match.
func (w raidWorld) awardRaid(t *testing.T, id int64) (*int64, bool) {
	t.Helper()
	var raidID *int64
	var attached bool
	if err := w.h.pool.QueryRow(t.Context(),
		`select raid_id, attached_by_match from loot_awards where id = $1`, id).Scan(&raidID, &attached); err != nil {
		t.Fatal(err)
	}
	return raidID, attached
}

func onRaid(got *int64, raidID int64) bool { return got != nil && *got == raidID }

func TestAnAwardMadeMidRaidLandsOnTheRaidWhenTheLogMatches(t *testing.T) {
	w := newRaidWorld(t, false)
	start := pastStart()
	raid := w.raidAt(t, raidZoneOny, start)
	mid := w.seedAward(t, 0, officerKey, start.Add(20*time.Minute), nil)
	early := w.seedAward(t, 1, officerKey, start.Add(-3*time.Hour), nil)
	w.oneFightNight(t, "log1", onyLogZone, start.Add(30*time.Minute))
	if _, ok, err := w.h.store.AutoMatch(t.Context(), "log1"); err != nil || !ok {
		t.Fatalf("auto match = %v %v", ok, err)
	}
	if got, attached := w.awardRaid(t, mid); !onRaid(got, raid.ID) || !attached {
		t.Fatalf("mid-raid award raid %v attached %v, want raid %d by the match", got, attached, raid.ID)
	}
	if got, _ := w.awardRaid(t, early); got != nil {
		t.Fatalf("an award before the window landed on raid %d", *got)
	}
}

func TestClearingTheMatchDetachesOnlyWhatTheMatchAttached(t *testing.T) {
	w := newRaidWorld(t, false)
	start := pastStart()
	raid := w.raidAt(t, raidZoneOny, start)
	w.oneFightNight(t, "log1", onyLogZone, start.Add(30*time.Minute))
	mid := w.seedAward(t, 0, officerKey, start.Add(20*time.Minute), nil)
	if _, err := w.h.store.SetMatch(t.Context(), w.gid, raid.ID, "log1", MatchedHand); err != nil {
		t.Fatal(err)
	}
	explicit := w.seedAward(t, 1, officerKey, start.Add(40*time.Minute), &raid.ID)
	if got, _ := w.awardRaid(t, mid); !onRaid(got, raid.ID) {
		t.Fatalf("hand match left the mid-raid award on %v", got)
	}
	if _, err := w.h.store.ClearMatch(t.Context(), w.gid, raid.ID, w.officer); err != nil {
		t.Fatal(err)
	}
	if got, attached := w.awardRaid(t, mid); got != nil || attached {
		t.Fatalf("after clear the attached award is on %v (attached %v), want no raid", got, attached)
	}
	if got, _ := w.awardRaid(t, explicit); !onRaid(got, raid.ID) {
		t.Fatalf("after clear the explicit award is on %v, want raid %d", got, raid.ID)
	}
}

func TestAnAttachCollisionKeepsTheNewerAwardOnTheRaid(t *testing.T) {
	w := newRaidWorld(t, false)
	start := pastStart()
	raid := w.raidAt(t, raidZoneOny, start)
	w.oneFightNight(t, "log1", onyLogZone, start.Add(30*time.Minute))
	// The raid already holds an award of item 0; a newer raid-less award of the same drop,
	// and an older raid-less award of item 1 against a newer one on the raid.
	olderOnRaid := w.seedAward(t, 0, officerKey, start.Add(10*time.Minute), &raid.ID)
	newerLoose := w.seedAward(t, 0, memberKeyA, start.Add(50*time.Minute), nil)
	newerOnRaid := w.seedAward(t, 1, officerKey, start.Add(70*time.Minute), &raid.ID)
	olderLoose := w.seedAward(t, 1, memberKeyA, start.Add(5*time.Minute), nil)
	if _, ok, err := w.h.store.AutoMatch(t.Context(), "log1"); err != nil || !ok {
		t.Fatalf("auto match = %v %v", ok, err)
	}
	for _, c := range []struct {
		name   string
		id     int64
		onRaid bool
	}{
		{"newer raid-less of item 0", newerLoose, true},
		{"older on-raid of item 0", olderOnRaid, false},
		{"newer on-raid of item 1", newerOnRaid, true},
		{"older raid-less of item 1", olderLoose, false},
	} {
		got, _ := w.awardRaid(t, c.id)
		if onRaid(got, raid.ID) != c.onRaid {
			t.Errorf("%s: raid %v, want on the raid %v", c.name, got, c.onRaid)
		}
	}
}
