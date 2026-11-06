package guilds

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strings"
	"testing"
)

// seedDecidedAward files an award with no night and a full decision record.
func (f exportFixture) seedDecidedAward(day string, itemID int) {
	f.h.t.Helper()
	if _, err := f.h.pool.Exec(f.h.t.Context(), `
		insert into loot_awards (guild_id, encounter_id, item_id, character_key, awarded_by, awarded_at,
			ms_os, reason, votes_yes, votes_no, votes_officers)
		values ($1, 1084, $4, 'us/hardcore/zed', $2, ($3 || ' 21:00+00')::timestamptz, 'os', '=Tank first', 3, 1, 4)`,
		f.gid, f.officer, day, itemID); err != nil {
		f.h.t.Fatal(err)
	}
}

func TestExportLoot_CarriesTheDecisionRecord(t *testing.T) {
	f := newExportFixture(t, "Export Decided")
	f.seedDecidedAward("2026-10-20", 16963)
	if _, err := f.h.pool.Exec(t.Context(),
		`insert into characters (key, region, ruleset, name, class, user_id) values ('us/hardcore/zed', 'us', 'hardcore', 'Zed', 'warrior', $1)`,
		f.officer); err != nil {
		t.Fatal(err)
	}
	rows := f.rows("loot")
	if len(rows) != 2 {
		t.Fatalf("rows = %v", rows)
	}
	got := rows[1][6:11]
	want := []string{"warrior", "os", "3", "1", "'=Tank first"}
	if fmt.Sprint(got) != fmt.Sprint(want) {
		t.Fatalf("decision cells = %v, want %v", got, want)
	}
}

func TestExportLootJSON_KeepsTheRange(t *testing.T) {
	f := newExportFixture(t, "Export Range")
	f.seedDecidedAward("2026-10-01", 16963)
	f.seedDecidedAward("2026-10-20", councilItem)
	var view LootHistoryView
	f.h.data(f.h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/export/loot.json?from=2026-10-10&to=2026-10-31", f.gid), ""), &view)
	if len(view.Rows) != 1 || view.Rows[0].Date != "2026-10-20" || view.Rows[0].VotesYes == nil || *view.Rows[0].VotesYes != 3 ||
		view.From == nil || *view.From != "2026-10-10" {
		t.Fatalf("view = %+v", view)
	}
	for _, q := range []string{"from=10/01/2026", "from=2026-10-31&to=2026-10-01"} {
		res := f.h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/export/loot.json?%s", f.gid, q), "")
		expectError(t, res, http.StatusBadRequest, "invalid")
	}
	res := f.h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/export/loot.csv?to=2026-10-05", f.gid), "")
	body := readBody(res)
	rows, err := csv.NewReader(strings.NewReader(body)).ReadAll()
	if err != nil || res.StatusCode != http.StatusOK || len(rows) != 2 || rows[1][0] != "2026-10-01" {
		t.Fatalf("csv to= = %d %q", res.StatusCode, body)
	}
}
