package guilds

import (
	"net/http"
	"slices"
	"strconv"
	"testing"
)

func readyKeys(items []ReadyItem) []string {
	out := make([]string, len(items))
	for i, it := range items {
		out[i] = it.Key + ":" + it.State
	}
	return out
}

// The checklist's items read as ready lines: attunement, the flask and food, one line per
// unenchanted slot; a check with no reading is left out.
func TestReadyItemsWordTheChecklist(t *testing.T) {
	ony := zoneOf(raidZoneOny)
	level := 60
	checks := []ChecklistItem{
		{Key: CheckLevel, State: CheckOK, Level: &level, RequiredLevel: &level},
		{Key: CheckAttunement, State: CheckOK, KeyItem: &KeyItem{Name: "Drakefire Amulet"}},
		{Key: CheckGearGap, State: CheckFail},
		{Key: CheckEnchants, State: CheckFail, MissingSlots: []string{"feet"}},
		{Key: CheckFlaskFood, State: CheckFail, Missing: []string{missingFlaskOrElixir},
			Supplies: []SupplyStack{{Name: "Grilled Squid", Kind: supplyFood, Count: 12}}},
	}
	got := readyItems(checks, &ony, nil)
	want := []string{"attune.onyxias-lair:ok", "enchant.feet:warn", "flask:miss", "food:ok"}
	if !slices.Equal(readyKeys(got), want) {
		t.Fatalf("items %v, want %v", readyKeys(got), want)
	}
	if got[0].Title != "Drakefire Amulet" || got[1].Title != "Boots not enchanted" || got[1].Line != "" || got[1].Where != nil || got[2].Where == nil ||
		got[2].Where.Label != "Auction house" || got[3].Title != "Grilled Squid" || got[3].Line != "12 in your bags" {
		t.Fatalf("words %+v", got)
	}
	if got[0].ItemID != nil || got[1].ItemID != nil || got[2].ItemID != nil {
		t.Fatalf("no ids on file, none sent: %+v", got)
	}
	withIDs := slices.Clone(checks)
	withIDs[1].KeyItem = &KeyItem{ID: 16309, Name: "Drakefire Amulet"}
	withIDs[4].Supplies = []SupplyStack{{ID: 13931, Name: "Nightfin Soup", Kind: supplyFood, Count: 4}}
	ids := readyItems(withIDs, &ony, map[string]int{"feet": 16912})
	for i, want := range map[int]int{0: 16309, 1: 16912, 3: 13931} {
		if ids[i].ItemID == nil || *ids[i].ItemID != want {
			t.Errorf("%s item_id %v, want %d", ids[i].Key, ids[i].ItemID, want)
		}
	}
	if ids[2].ItemID != nil {
		t.Errorf("a missing flask names no item, got %d", *ids[2].ItemID)
	}
	unseen := []ChecklistItem{{Key: CheckAttunement, State: CheckNotTracked, Reason: reasonKeysNotInExport},
		{Key: CheckFlaskFood, State: CheckNotTracked, Reason: reasonBagsNotShared}}
	if got := readyItems(unseen, &ony, nil); len(got) != 1 || got[0].State != ReadyWarn {
		t.Fatalf("unseen %+v", got)
	}
}

// Only the owner reads it; a raid of a guild the character is not in reads 404; with the
// export the items come back with its time.
func TestReadyCheckRoute(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	if _, err := w.h.pool.Exec(t.Context(), `
		insert into addon_exports (character_key, user_id, region, ruleset, name, export, captured_at, updated_at)
		values ($1, $2, 'us', 'hardcore', 'Mirelle', 'FS1:1.60.1.70009:priest:human:0/0/0:chest=11726|level=60', now(), now())`,
		member2Key, w.member2); err != nil {
		t.Fatal(err)
	}
	path := "/v1/characters/us/hardcore/mirelle/ready"
	w.as(w.member)
	if res := w.h.do(http.MethodGet, path, ""); res.StatusCode != http.StatusNotFound {
		t.Fatalf("another's = %d", res.StatusCode)
	}
	w.as(w.member2)
	other := seedGuild(t, w.h.pool, "Elsewhere")
	if _, err := w.h.pool.Exec(t.Context(), `insert into guild_raids (guild_id, zone, title, starts_at, duration_minutes, created_by)
		values ($1, $2, 'Ony', now() + interval '2 days', 180, $3)`, other, raidZoneOny, w.member2); err != nil {
		t.Fatal(err)
	}
	var otherRaid int64
	if err := w.h.pool.QueryRow(t.Context(), `select id from guild_raids where guild_id = $1`, other).Scan(&otherRaid); err != nil {
		t.Fatal(err)
	}
	if res := w.h.do(http.MethodGet, path+"?raid="+strconv.FormatInt(otherRaid, 10), ""); res.StatusCode != http.StatusNotFound {
		t.Fatalf("another guild's raid = %d", res.StatusCode)
	}
	res := w.h.do(http.MethodGet, path+"?raid="+strconv.FormatInt(raid.ID, 10), "")
	if res.StatusCode != http.StatusOK {
		t.Fatalf("own = %d: %s", res.StatusCode, readBody(res))
	}
	var check ReadyCheck
	w.h.data(res, &check)
	if check.ExportedAt == nil || len(check.Items) == 0 {
		t.Fatalf("check %+v", check)
	}
	if !slices.Contains(readyKeys(check.Items), "enchant.chest:warn") {
		t.Fatalf("items %v", readyKeys(check.Items))
	}
}
