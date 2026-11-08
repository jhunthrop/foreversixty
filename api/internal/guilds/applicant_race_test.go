package guilds

import (
	"fmt"
	"net/http"
	"testing"
)

// The applicant's character carries race and faction: Battle.net's, else the export's race.
func TestApplicantCharacterCarriesRaceAndFaction(t *testing.T) {
	h := newHTTPHarness(t)
	gid, officer := recruitGuild(t, h, "Races")
	openRecruiting(t, h, gid, officer)
	uid, key := applicant(t, h, "race-app@example.com", "racechar")
	if _, err := h.pool.Exec(t.Context(), `update characters set race = 'Dwarf', faction = 'alliance' where key = $1`, key); err != nil {
		t.Fatal(err)
	}
	h.as(uid)
	var a Application
	h.data(h.do(http.MethodPost, fmt.Sprintf("/v1/guilds/%d/applications", gid), applyBody(key, false)), &a)
	if a.Character.Race == nil || *a.Character.Race != "Dwarf" || a.Character.Faction == nil || *a.Character.Faction != "alliance" {
		t.Fatalf("character %+v", a.Character)
	}
	if raceWords("night-elf") != "Night Elf" {
		t.Fatalf("race words %q", raceWords("night-elf"))
	}
}
