package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/parquet"
	"github.com/jhunthrop/foreversixty/logs/engine/store"
)

// writeSample lays out a sample cut by engine 0.5: one fight whose line carries the
// double-counted npc_kills, its events, and a meta.json with a key recount must keep.
func writeSample(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	at := time.Date(2026, 9, 17, 20, 0, 0, 0, time.UTC)
	logger, mob := "Player-4184-000000B1", "Creature-0-1-2-3-4-0000AA0001"
	events := []event.Event{
		{Time: at, Kind: event.Damage, Name: "SWING_DAMAGE",
			Source: event.Unit{GUID: logger, Flags: 0x511}, Dest: event.Unit{GUID: mob, Flags: 0xa48}},
		{Time: at, Kind: event.PartyKill, Name: "PARTY_KILL",
			Source: event.Unit{GUID: logger, Flags: 0x511}, Dest: event.Unit{GUID: mob, Flags: 0xa48}},
		{Time: at, Kind: event.Death, Name: "UNIT_DIED", Dest: event.Unit{GUID: mob, Flags: 0xa48}},
	}
	pq, err := parquet.Marshal(events)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(dir, "fights", "1"), 0o755); err != nil {
		t.Fatal(err)
	}
	mustWrite(t, filepath.Join(dir, "fights", "1", "events.parquet"), pq)
	line := `{"index":1,"kind":"trash","name":"Trash","kill":false,"in_progress":false,` +
		`"start":"2026-09-17T20:00:00Z","end":"2026-09-17T20:00:09Z","duration_ms":9000,` +
		`"players":["` + logger + `"],"deaths":0,"npc_kills":2}`
	mustWrite(t, filepath.Join(dir, "report.json"),
		[]byte(`{"report_id":"sample","engine_version":"0.5.3","health":{},"fights":[`+line+`],"units":[]}`))
	mustWrite(t, filepath.Join(dir, "meta.json"),
		[]byte(`{"engine_version":"0.5.3","fights":[`+line+`],"phase":"launch","title":"Beta night"}`))
	return dir
}

func mustWrite(t *testing.T, path string, b []byte) {
	t.Helper()
	if err := os.WriteFile(path, b, 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestRecountRewritesTheCountsFromEachFightsEvents(t *testing.T) {
	dir := writeSample(t)
	if err := recount(dir); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"report.json", "meta.json"} {
		b, err := os.ReadFile(filepath.Join(dir, name))
		if err != nil {
			t.Fatal(err)
		}
		if strings.Contains(string(b), "npc_kills") {
			t.Errorf("%s still carries npc_kills", name)
		}
		var doc struct {
			Fights []store.FightEntry `json:"fights"`
		}
		if err := json.Unmarshal(b, &doc); err != nil {
			t.Fatal(err)
		}
		f := doc.Fights[0]
		if f.MobDeaths == nil || *f.MobDeaths != 1 || f.GroupKills == nil || *f.GroupKills != 1 {
			t.Errorf("%s: counts = %v/%v, want one mob death, the group's", name, f.MobDeaths, f.GroupKills)
		}
		if f.GroupDeaths == nil || len(f.Group) != 1 {
			t.Errorf("%s: group deaths %v, group %v", name, f.GroupDeaths, f.Group)
		}
	}
	meta, err := os.ReadFile(filepath.Join(dir, "meta.json"))
	if err != nil {
		t.Fatal(err)
	}
	for _, keep := range []string{`"phase": "launch"`, `"title": "Beta night"`, `"engine_version": "0.5.3"`} {
		if !strings.Contains(string(meta), keep) {
			t.Errorf("meta.json lost %s:\n%s", keep, meta)
		}
	}
}
