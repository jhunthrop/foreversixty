package addon

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// An upload of a verified member's export stores its drops= section for the guild, once however
// often it is uploaded, and the export itself is stored as before.
func TestAnUploadStoresTheExportsDrops(t *testing.T) {
	h := newHarness(t)
	guildID := joinByInviteAndSync(t, h, "2")
	ctx := context.Background()
	export := fmt.Sprintf("FS1:1.60.1.69893:warrior:tauren:0/0/0:|drops=%d:249:1084:master:Creature-0-1-249-1-10184-5:17078x1@Baelgrim|guild=Forever:2",
		time.Now().Unix())
	for range 2 {
		if err := h.store.PutExports(ctx, h.owner, []Export{{Name: "Baelgrim", Region: "us", Ruleset: "hardcore", Export: export}}); err != nil {
			t.Fatal(err)
		}
	}
	var n int
	var by string
	if err := h.pool.QueryRow(ctx, `select count(*), max(reported_character) from raid_drops where guild_id = $1 and item_id = 17078`,
		guildID).Scan(&n, &by); err != nil {
		t.Fatal(err)
	}
	if n != 1 || by != "us/hardcore/baelgrim" {
		t.Fatalf("drops = %d by %q, want one by the character", n, by)
	}
}
