package guilds

import (
	"reflect"
	"sort"
	"testing"
)

// consumablesByRoster is the consumables block as it was first built, from the whole home
// roster: the reference the narrow query must equal.
func consumablesByRoster(t *testing.T, w raidWorld, vc raidViewContext, reportID string) []ConsumableRow {
	t.Helper()
	inLog, err := w.h.store.reportPlayerKeys(t.Context(), reportID)
	if err != nil {
		t.Fatal(err)
	}
	g, err := w.h.store.getGuild(t.Context(), w.gid)
	if err != nil {
		t.Fatal(err)
	}
	roster, err := w.h.store.HomeRoster(t.Context(), w.gid, g, 0, false, false)
	if err != nil {
		t.Fatal(err)
	}
	out := []ConsumableRow{}
	for _, row := range roster {
		if row.Verified && inLog[row.CharacterKey] {
			out = append(out, ConsumableRow{CharacterKey: row.CharacterKey, Name: row.Name, Class: row.className,
				Before: consumablesState(row.Consent, row.bags, row.hasBagsSection), own: vc.ownsKey(row.CharacterKey)})
		}
	}
	sort.SliceStable(out, func(i, j int) bool {
		if consumableOrder[out[i].Before] != consumableOrder[out[j].Before] {
			return consumableOrder[out[i].Before] < consumableOrder[out[j].Before]
		}
		return out[i].Name < out[j].Name
	})
	return out
}

func TestTheConsumablesQueryEqualsTheRosterPath(t *testing.T) {
	w, _, _ := reportNight(t)
	seedExport(t, w.h.pool, w.member2, member2Key, "us", "hardcore", "Mirelle")
	w.exec(t, `update guild_members set consent = 'gear_bags' where user_id = (select id from users where email = 'officer@example.com')`)
	w.exec(t, `update guild_members set consent = 'gear' where user_id = (select id from users where email = 'member2@example.com')`)
	for _, viewer := range []int64{w.officer, w.member} {
		vc, err := w.h.store.loadRaidOwnership(t.Context(), w.gid, viewer)
		if err != nil {
			t.Fatal(err)
		}
		want := consumablesByRoster(t, w, vc, "night")
		got, err := w.h.store.reportConsumables(t.Context(), vc, "night")
		if err != nil {
			t.Fatal(err)
		}
		if len(want) < 3 || !reflect.DeepEqual(got.Rows, want) {
			t.Fatalf("viewer %d: rows %+v, want %+v", viewer, got.Rows, want)
		}
	}
}
