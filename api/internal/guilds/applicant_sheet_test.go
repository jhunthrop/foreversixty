package guilds

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"
)

// fakeScores answers one score for every character.
type fakeScores struct{ score *int }

func (f fakeScores) ScoreOf(context.Context, string) (*int, error) { return f.score, nil }

// The officer's list carries what the applicant sheet shows: the score, the latest sim, the
// enchants from the export, the guilds the character logged raids with, and the export's time.
func TestOfficerApplicationsCarryTheSheet(t *testing.T) {
	h := newHTTPHarness(t)
	gid, officer := recruitGuild(t, h, "Sheets")
	openRecruiting(t, h, gid, officer)
	uid, key := applicant(t, h, "sheet-app@example.com", "sheetchar")
	ctx := context.Background()
	exec := func(sql string, args ...any) {
		t.Helper()
		if _, err := h.pool.Exec(ctx, sql, args...); err != nil {
			t.Fatal(err)
		}
	}
	// A two-hander and a chest, the chest enchanted: 1 of 2 enchantable slots.
	exec(`insert into addon_exports (character_key, user_id, region, ruleset, name, export, captured_at, updated_at)
		values ($1, $2, 'us', 'hardcore', 'sheetchar', 'FS1:1.60.1.70009:paladin:human:0/0/0:main_hand=12640,chest=11726:41|level=60', now(), now())`,
		key, uid)
	exec(`insert into sims (id, user_id, spec, engine_version, lane, dps_mean, dps_error, iterations, result, state, finished_at)
		values ('sheet-sim', $1, 'paladin-holy', 'e', 'server', 612.4, 1, 3000, '{}', 'done', now())`, uid)
	exec(`insert into character_sims (sim_id, character_key, user_id, kind, spec, role, band, preset)
		values ('sheet-sim', $1, $2, 'import', 'paladin-holy', 'healer', 60, 'raid')`, key, uid)
	old := seedGuild(t, h.pool, "Dawnbreak")
	at := time.Now().Add(-20 * 24 * time.Hour).UTC().Truncate(time.Hour)
	seedNight(t, h.pool, old, "sheet-r1", "Onyxia's Lair", at, []seedFight{{encounterID: 1084, kill: true, duration: time.Minute,
		players: []seedPlayer{{key: key, name: "sheetchar", class: "paladin", spec: "Holy", role: "healer"}}}})
	exec(`update reports set visibility = 'public' where id = 'sheet-r1'`)

	h.as(uid)
	if res := h.do(http.MethodPost, fmt.Sprintf("/v1/guilds/%d/applications", gid), applyBody(key, false)); res.StatusCode != http.StatusCreated {
		t.Fatalf("apply = %d: %s", res.StatusCode, readBody(res))
	}
	score := 74
	h.svc.Scores = fakeScores{score: &score}
	h.as(officer)
	var page struct{ Rows []OfficerApplication }
	h.data(h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/applications", gid), ""), &page)
	if len(page.Rows) != 1 || page.Rows[0].Sheet == nil {
		t.Fatalf("rows %+v", page.Rows)
	}
	sh := page.Rows[0].Sheet
	if sh.Score == nil || *sh.Score != 74 {
		t.Fatalf("score %+v", sh.Score)
	}
	if sh.Sim == nil || sh.Sim.Value != 612.4 || sh.Sim.Unit != gainUnitHPS {
		t.Fatalf("sim %+v", sh.Sim)
	}
	if sh.Enchants == nil || sh.Enchants.Slots != 2 || sh.Enchants.Enchanted != 1 || len(sh.Enchants.Missing) != 1 || sh.Enchants.Missing[0] != "main_hand" {
		t.Fatalf("enchants %+v", sh.Enchants)
	}
	if len(sh.History) != 1 || sh.History[0].GuildName != "Dawnbreak" || sh.History[0].Raids != 1 {
		t.Fatalf("history %+v", sh.History)
	}
	if sh.ExportedAt == nil {
		t.Fatal("no export time")
	}
}
