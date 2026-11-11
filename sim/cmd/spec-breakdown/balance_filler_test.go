package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/internal/inproc"
)

// Below 20 a Balance druid has no Starfire, so the card's mana-pacing Wrath
// (cast only while mana is behind the fight) left nothing to cast whenever
// mana was ahead and Moonfire was ticking: a level 14 moonkin with a full
// pool stood idle. Before Starfire is trained Wrath is the filler, as a real
// druid plays it. Geared for mana (mp5 pieces well above its level, so the
// pool never falls behind the fight's pace), a level 14 moonkin must spend a
// 60 s fight casting Wrath between Moonfires.
func TestBalanceBelowTwentyCastsWrathWhileManaIsAhead(t *testing.T) {
	dir := bandTenBalanceEntries(t)
	o := options{
		repoRoot: "../../..", spec: "druid-balance", preset: "bare", band: 10, level: 14,
		faction: "horde", iterations: 50, seed: 1, duration: 60, bisDir: dir,
		gear:    "main_hand:22632,head:20327,chest:20331,legs:20329,finger1:23037,trinket1:17064,feet:19391,hands:21619",
		talents: "5000000000000000-00000000000000000000-0000000000000000",
	}
	req, err := buildRequest(o)
	if err != nil {
		t.Fatal(err)
	}
	est, player, err := inproc.PlainRunWithRotation(req, nil)
	if err != nil {
		t.Fatal(err)
	}
	r := newReport(req, fightSecondsOf(req), est, player, nil)
	wraths := 0.0
	for _, a := range r.Actions {
		for _, id := range wrathRankIDs {
			if strings.HasSuffix(a.Name, fmt.Sprintf("(%d)", id)) {
				wraths += a.CastsPerIter
			}
		}
	}
	// Wrath is a 1.5 s cast with Improved Wrath 5/5: about 35 fit in 60 s
	// beside the Moonfires. The idle card cast three.
	if wraths < 30 {
		t.Errorf("level 14, 60 s with mana ahead: %.1f Wraths cast, want at least 30 (the card idles while mana is ahead)", wraths)
	}
}

// wrathRankIDs are Wrath ranks 1 to 8.
var wrathRankIDs = []int{5176, 5177, 5178, 5179, 5180, 6780, 8905, 9912}

// bandTenBalanceEntries writes a bis directory holding the published band 20
// bare entries relabelled band 10 with no gear, so spec-breakdown can sim a
// level 10 to 19 moonkin (the ranker publishes no band 10).
func bandTenBalanceEntries(t *testing.T) string {
	t.Helper()
	build, err := buildDirectory("../../..")
	if err != nil {
		t.Fatal(err)
	}
	raw, err := os.ReadFile(filepath.Join(build, "bis", "druid-balance.json"))
	if err != nil {
		t.Fatal(err)
	}
	var file map[string]json.RawMessage
	if err := json.Unmarshal(raw, &file); err != nil {
		t.Fatal(err)
	}
	var bands []map[string]any
	if err := json.Unmarshal(file["bands"], &bands); err != nil {
		t.Fatal(err)
	}
	var tens []map[string]any
	for _, b := range bands {
		if b["band"] == float64(20) && b["preset"] == "bare" {
			ten := map[string]any{}
			for k, v := range b {
				ten[k] = v
			}
			ten["band"] = 10
			ten["slots"] = []any{}
			tens = append(tens, ten)
		}
	}
	if len(tens) == 0 {
		t.Fatal("no band 20 bare Balance entry to relabel")
	}
	out, err := json.Marshal(map[string]any{"spec": "druid-balance", "bands": tens})
	if err != nil {
		t.Fatal(err)
	}
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "druid-balance.json"), out, 0o644); err != nil {
		t.Fatal(err)
	}
	return dir
}
