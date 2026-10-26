package notify

import (
	"strings"
	"testing"
	"time"
)

// onyxiaThursday is a guild in America/Chicago raiding at 20:00 on Thursday 10 December
// 2026 (02:00 UTC on the 11th), 23 of 40 signed up.
func onyxiaThursday() RaidPosted {
	return RaidPosted{
		GuildName: "Olympus XXVII", Title: "Onyxia's Lair",
		StartsAt: time.Date(2026, 12, 11, 2, 0, 0, 0, time.UTC), Timezone: "America/Chicago",
		RaidURL:  "https://foreversixty.gg/guild/us/normal/olympus-xxvii/raids/7",
		CrestURL: "https://api.foreversixty.gg/v1/guilds/7/crest.webp?v=abc",
		SignedUp: 23, Size: 40, SizeKnown: true,
	}
}

const (
	wireHead   = `{"username":"Forever Sixty","avatar_url":"https://foreversixty.gg/seal-400.png","embeds":[`
	wireTail   = `],"allowed_mentions":{"parse":[]}}`
	raidLink   = `"url":"https://foreversixty.gg/guild/us/normal/olympus-xxvii/raids/7",`
	raidFoot   = `"footer":{"text":"Olympus XXVII · Forever Sixty"},"thumbnail":{"url":"https://api.foreversixty.gg/v1/guilds/7/crest.webp?v=abc"},"timestamp":"2026-10-25T18:30:00Z",`
	raidWhen   = `{"name":"When","value":"\u003ct:1796954400:F\u003e (\u003ct:1796954400:R\u003e)","inline":true}`
	raidSigned = `{"name":"Signed up","value":"23 of 40","inline":true}`
)

func TestRaidEmbedsMatchTheCopyTable(t *testing.T) {
	moved := RaidChanged{RaidPosted: onyxiaThursday()}
	was := moved.StartsAt
	moved.StartsAt, moved.Was = was.Add(24*time.Hour), &was
	cancelled := RaidChanged{RaidPosted: onyxiaThursday(), Cancelled: true}
	cases := map[string]struct {
		embed Embed
		want  string
	}{
		"posted": {
			RaidPostedEmbed(onyxiaThursday(), fixedFinish),
			`{"title":"Onyxia's Lair · Thursday 10 December","description":"Sign up on Forever Sixty.",` + raidLink +
				`"fields":[` + raidWhen + `,{"name":"Guild time","value":"8:00 pm CST","inline":true},` + raidSigned + `],` +
				raidFoot + `"color":15055189}`,
		},
		"moved": {
			RaidChangedEmbed(moved, fixedFinish),
			`{"title":"Moved: Onyxia's Lair · Friday 11 December","description":"Sign-ups carry over. Check you can still make it.",` +
				raidLink + `"fields":[{"name":"When","value":"\u003ct:1797040800:F\u003e (\u003ct:1797040800:R\u003e)","inline":true},` +
				`{"name":"Was","value":"\u003ct:1796954400:F\u003e","inline":true},` + raidSigned + `],` +
				raidFoot + `"color":15055189}`,
		},
		"cancelled": {
			RaidChangedEmbed(cancelled, fixedFinish),
			`{"title":"Cancelled: Onyxia's Lair · Thursday 10 December",` + raidLink +
				`"fields":[{"name":"When","value":"~~\u003ct:1796954400:F\u003e (\u003ct:1796954400:R\u003e)~~","inline":true},` +
				raidSigned + `],` + raidFoot + `"color":15758187}`,
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if got, want := wireOf(t, c.embed), wireHead+c.want+wireTail; got != want {
				t.Fatalf("\n got %s\nwant %s", got, want)
			}
		})
	}
}

func TestRaidPostedEdges(t *testing.T) {
	in := onyxiaThursday()
	in.Title, in.GuildName = "*Ony* @everyone", "Olympus_XXVII"
	in.SignedUp, in.Timezone, in.CrestURL = 0, "", "/v1/guilds/7/crest.webp?v=abc"
	e := RaidPostedEmbed(in, fixedFinish)
	if e.Title != `\*Ony\* @everyone · Friday 11 December` {
		t.Fatalf("escaped title in UTC = %q", e.Title)
	}
	if e.Fields[1].Value != "2:00 am UTC" || e.Fields[2].Value != "None yet" {
		t.Fatalf("fields = %+v", e.Fields)
	}
	if e.Thumbnail != "" {
		t.Fatalf("relative crest posted as thumbnail %q", e.Thumbnail)
	}
	if e.Footer != "Olympus_XXVII · Forever Sixty" {
		t.Fatalf("footer = %q", e.Footer)
	}
	in.SignedUp, in.SizeKnown = 12, false
	if got := RaidPostedEmbed(in, fixedFinish).Fields[2].Value; got != "12 signed up" {
		t.Fatalf("unknown size = %q", got)
	}
	in.Title = strings.Repeat("_", 80)
	if got := []rune(RaidPostedEmbed(in, fixedFinish).Title); len(got) > maxTitleRunes {
		t.Fatalf("title %d runes", len(got))
	}
}
