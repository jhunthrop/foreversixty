package guilds

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/contract"
)

// The Raid tab's new reads keep their app fixtures.
func TestAppContractRaidTabReads(t *testing.T) {
	w := newRaidWorld(t, true)
	w.recruit(t, w.gid, true, "alliance", "Wed and Sat, 8:00 pm CST", `[]`)
	w.as(0)
	res := w.h.do(http.MethodGet, "/v1/guilds/recruiting?spec=warrior-fury&faction=alliance&region=us&ruleset=hardcore", "")
	contract.AssertFixture(t, "guild.recruiting", guildBody(t, res, http.StatusOK), contract.Options{})

	raid := w.createRaid(t)
	if _, err := w.h.pool.Exec(t.Context(), `
		insert into addon_exports (character_key, user_id, region, ruleset, name, export, captured_at, updated_at)
		values ($1, $2, 'us', 'hardcore', 'Mirelle', 'FS1:1.60.1.70009:priest:human:0/0/0:chest=11726|level=60', now(), now())`,
		member2Key, w.member2); err != nil {
		t.Fatal(err)
	}
	w.as(w.member2)
	res = w.h.do(http.MethodGet, "/v1/characters/us/hardcore/mirelle/ready?raid="+strconv.FormatInt(raid.ID, 10), "")
	// A line's icon is the item's icon when the build's data has it (openapi ReadyItem.icon, omitempty);
	// this world loads no item data, so its chest line carries the item_id alone.
	contract.AssertFixture(t, "character.ready", guildBody(t, res, http.StatusOK), contract.Options{
		Optional: []string{"data.items[].icon"},
	})
}
