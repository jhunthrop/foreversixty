package discordbot

import (
	"encoding/json"
	"os"
	"testing"
)

// siteLinksPath is the site's community links, the source CommunityDiscordInvite copies.
const siteLinksPath = "../../../web/src/data/links.json"

func TestCommunityDiscordInviteIsTheSites(t *testing.T) {
	raw, err := os.ReadFile(siteLinksPath)
	if err != nil {
		t.Fatal(err)
	}
	var links struct {
		DiscordInvite string `json:"discordInvite"`
	}
	if err := json.Unmarshal(raw, &links); err != nil {
		t.Fatal(err)
	}
	if links.DiscordInvite != CommunityDiscordInvite {
		t.Fatalf("CommunityDiscordInvite = %q, the site's discordInvite = %q", CommunityDiscordInvite, links.DiscordInvite)
	}
}
