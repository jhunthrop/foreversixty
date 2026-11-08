package dataaddon

import (
	"context"
	"testing"
	"time"
)

// A character whose owner hid its Forever Score has no record in the nightly file and is in
// no guild's member list; nothing else changes (contract 4 export.score_visible, contract 5).
func TestAHiddenScoreIsLeftOutOfTheNightlyFile(t *testing.T) {
	pool := testPool(t)
	data := useRepoData(t)
	ctx := context.Background()
	now := time.Now().UTC()
	seedGolden(t, pool, now.Truncate(24*time.Hour))
	if _, err := pool.Exec(ctx, `update characters set score_visible = false where key = 'us/normal/dav2-obnoxious-yell'`); err != nil {
		t.Fatal(err)
	}
	d, _, err := collect(ctx, Deps{Store: &Store{Pool: pool}, Trees: data, Build: "1.60.1.70291", Now: now})
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := d.Characters["Obnoxious Yell-ClassicBetaPvE2"]; ok {
		t.Fatal("the hidden character has a record")
	}
	if _, ok := d.Characters["Mörk-ClassicBetaPvP"]; !ok {
		t.Fatal("another character lost its record")
	}
	for _, m := range d.Guilds["Dav2 Olympus"].Members {
		if m == "Obnoxious Yell-ClassicBetaPvE2" {
			t.Fatal("the guild still lists the hidden character")
		}
	}
}
