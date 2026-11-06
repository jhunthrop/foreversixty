package guilds

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

const dropGUID = "Creature-0-4469-249-1234-10184-00004F3A2B"

// dropsExport is an export whose drops= section holds corpses.
func dropsExport(corpses ...string) string {
	out := "FS1:1.60.1.70790:warrior:human:0/0/0:head=16900|drops="
	for i, c := range corpses {
		if i > 0 {
			out += ";"
		}
		out += c
	}
	return out
}

// dropCorpse is one corpse of a drops= section, looted at `at`.
func dropCorpse(at time.Time, encounter int64, source, items string) string {
	return fmt.Sprintf("%d:249:%d:master:%s:%s", at.Unix(), encounter, source, items)
}

// record stores key's export the way an upload does.
func (w raidWorld) record(t *testing.T, uid int64, key, export string) {
	t.Helper()
	tx, err := w.h.pool.Begin(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = tx.Rollback(t.Context()) }()
	if err := RecordExportDrops(t.Context(), tx, uid, key, export, time.Now()); err != nil {
		t.Fatal(err)
	}
	if err := tx.Commit(t.Context()); err != nil {
		t.Fatal(err)
	}
}

type storedDrop struct {
	raidID   *int64
	reportID *string
	attached bool
	count    int
	winners  []string
	method   string
}

func (w raidWorld) storedDrop(t *testing.T, item int) (storedDrop, int) {
	t.Helper()
	rows, err := w.h.pool.Query(t.Context(), `
		select raid_id, report_id, attached_by_match, count, winners, loot_method from raid_drops
		where guild_id = $1 and item_id = $2`, w.gid, item)
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
	var d storedDrop
	n := 0
	for rows.Next() {
		n++
		if err := rows.Scan(&d.raidID, &d.reportID, &d.attached, &d.count, &d.winners, &d.method); err != nil {
			t.Fatal(err)
		}
	}
	return d, n
}

func TestRecordedDropsKeepACorpseOnceAcrossRaiders(t *testing.T) {
	w := newRaidWorld(t, false)
	now := time.Now().UTC()
	w.record(t, w.member, memberKeyA, dropsExport(dropCorpse(now, 1084, dropGUID, fmt.Sprintf("%dx1", councilItem))))
	w.record(t, w.member2, member2Key, dropsExport(
		dropCorpse(now.Add(time.Second), 1084, dropGUID, fmt.Sprintf("%dx1@Kraggor,99999x1", councilItem)),
		dropCorpse(now, 0, "Creature-0-1-249-1-555-1", "99998x1"),
		dropCorpse(now.Add(-40*24*time.Hour), 1084, "Creature-0-1-249-1-10184-9", "18205x1")))
	w.record(t, w.member, memberKeyC, dropsExport(dropCorpse(now, 1084, "Creature-0-1-249-1-10184-8", "18813x1")))

	d, n := w.storedDrop(t, councilItem)
	if n != 1 || d.count != 1 || len(d.winners) != 1 || d.winners[0] != "Kraggor" || d.method != "master" {
		t.Fatalf("drape = %+v (%d rows)", d, n)
	}
	for item, want := range map[int]int{99999: 1, 99998: 0, 18205: 0, 18813: 0} {
		if _, n := w.storedDrop(t, item); n != want {
			t.Errorf("item %d rows = %d, want %d (unknown boss, too old, unverified)", item, n, want)
		}
	}
	// A reupload of the same export changes nothing; a GUID-less corpse is the kill of the day.
	w.record(t, w.member, memberKeyA, dropsExport(dropCorpse(now, 1084, "", "18205x1"), dropCorpse(now, 1084, "", "18205x2")))
	if d, n := w.storedDrop(t, 18205); n != 1 || d.count != 2 {
		t.Fatalf("kill-keyed corpse = %+v (%d rows)", d, n)
	}
}

func TestARecordedDropFilesUnderItsNight(t *testing.T) {
	w := newRaidWorld(t, false)
	now := time.Now().UTC().Truncate(time.Minute)
	raid := w.raidAt(t, raidZoneOny, now.Add(-time.Hour))
	w.record(t, w.member, memberKeyA, dropsExport(dropCorpse(now, 1084, dropGUID, fmt.Sprintf("%dx1", councilItem))))
	if d, _ := w.storedDrop(t, councilItem); !onRaid(d.raidID, raid.ID) || d.attached {
		t.Fatalf("drop in the raid's window = %+v, want raid %d", d, raid.ID)
	}
	// A logged night with no planned raid takes a drop looted after its kill.
	logged := now.Add(-5 * 24 * time.Hour)
	w.oneFightNight(t, "old", onyLogZone, logged)
	w.record(t, w.member, memberKeyA, dropsExport(dropCorpse(logged.Add(20*time.Minute), 1084, "Creature-0-1-249-1-10184-7", "18205x1")))
	if d, _ := w.storedDrop(t, 18205); d.raidID != nil || d.reportID == nil || *d.reportID != "old" {
		t.Fatalf("drop after the logged kill = %+v, want report old", d)
	}
}

func TestRecordedDropsFollowTheMatchAndTheClear(t *testing.T) {
	w := newRaidWorld(t, false)
	start := pastStart()
	w.record(t, w.member, memberKeyA, dropsExport(dropCorpse(start.Add(40*time.Minute), 1084, dropGUID, fmt.Sprintf("%dx1", councilItem))))
	raid := w.raidAt(t, raidZoneOny, start)
	w.oneFightNight(t, "log1", onyLogZone, start.Add(30*time.Minute))
	if _, err := w.h.store.SetMatch(t.Context(), w.gid, raid.ID, "log1", MatchedHand); err != nil {
		t.Fatal(err)
	}
	if d, _ := w.storedDrop(t, councilItem); !onRaid(d.raidID, raid.ID) || !d.attached {
		t.Fatalf("after match = %+v, want attached to raid %d", d, raid.ID)
	}
	if _, err := w.h.store.ClearMatch(t.Context(), w.gid, raid.ID, w.officer); err != nil {
		t.Fatal(err)
	}
	if d, _ := w.storedDrop(t, councilItem); d.raidID != nil || d.attached {
		t.Fatalf("after clear = %+v, want no raid", d)
	}
}

// nightDrop is the night's drop of item, failing when it has none.
func nightDrop(t *testing.T, n LootNight, item int) LootNightDrop {
	t.Helper()
	for _, d := range n.Drops {
		if d.ItemID == item {
			return d
		}
	}
	t.Fatalf("night %s has no drop %d: %+v", n.Night, item, n.Drops)
	return LootNightDrop{}
}

func TestTheNightsListRecordedDropsOpenUntilAwarded(t *testing.T) {
	w := newRaidWorld(t, true)
	now := time.Now().UTC().Truncate(time.Minute)
	raid := w.raidAt(t, raidZoneOny, now.Add(-time.Hour))
	key := nightKeyRaid + fmt.Sprint(raid.ID)
	w.record(t, w.member, memberKeyA, dropsExport(dropCorpse(now, 1084, dropGUID,
		fmt.Sprintf("%dx1@Kraggor,18205x1", councilItem))))

	w.as(w.officer)
	var view LootNightsView
	w.h.data(w.h.do(http.MethodGet, w.path("/loot/nights"), ""), &view)
	if len(view.Nights) != 1 || view.Nights[0].Night != key || len(view.Nights[0].Drops) != 2 {
		t.Fatalf("nights = %+v", view.Nights)
	}
	drape := nightDrop(t, view.Nights[0], councilItem)
	if !drape.Recorded || drape.Count != 1 || drape.State != dropOpen || len(drape.LootedBy) != 1 || drape.LootedBy[0] != "Kraggor" {
		t.Fatalf("drape = %+v", drape)
	}
	var actions OfficerActions
	w.h.data(w.h.do(http.MethodGet, w.path("/officer/actions"), ""), &actions)
	if a := actions.DropsToAward; a == nil || a.Count == nil || *a.Count != 2 || a.RaidID == nil || *a.RaidID != raid.ID {
		t.Fatalf("drops to award = %+v", a)
	}

	if res := w.h.do(http.MethodPost, w.path("/loot/awards"), fmt.Sprintf(
		`{"encounter_id":1084,"item_id":%d,"character_key":%q,"night":%q}`, councilItem, memberKeyA, key)); res.StatusCode != http.StatusOK {
		t.Fatalf("award = %d: %s", res.StatusCode, readBody(res))
	}
	w.h.data(w.h.do(http.MethodGet, w.path("/officer/actions"), ""), &actions)
	if a := actions.DropsToAward; a == nil || *a.Count != 1 || a.Awarded != 1 {
		t.Fatalf("drops to award after the award = %+v", a)
	}

	w.as(w.member)
	w.h.data(w.h.do(http.MethodGet, w.path("/loot/nights"), ""), &view)
	drape = nightDrop(t, view.Nights[0], councilItem)
	if !drape.Recorded || drape.State != dropAwarded || drape.Award == nil || drape.Award.VotesOfficers != nil {
		t.Fatalf("member's drape = %+v award %+v", drape, drape.Award)
	}
	if other := nightDrop(t, view.Nights[0], 18205); !other.Recorded || other.State != dropOpen || other.LootedBy == nil {
		t.Fatalf("member's open drop = %+v", other)
	}
}
