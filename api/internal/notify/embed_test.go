package notify

import (
	"encoding/json"
	"strings"
	"testing"
	"time"
)

var fixedFinish = time.Date(2026, 10, 25, 18, 30, 0, 0, time.UTC)

const (
	testBase   = "https://foreversixty.gg"
	testAvatar = "https://foreversixty.gg/seal-400.png"
)

func wireOf(t *testing.T, e Embed) string {
	t.Helper()
	b, err := json.Marshal(newMessage(testAvatar, e))
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

const common = `"footer":{"text":"Forever Sixty"},"timestamp":"2026-10-25T18:30:00Z","color":15055189`
const footer = `"footer":{"text":"Forever Sixty"},`
const stamp = `"timestamp":"2026-10-25T18:30:00Z","color":15055189`

func TestEmbedsMatchTheCopyTable(t *testing.T) {
	head := `{"username":"Forever Sixty","avatar_url":"https://foreversixty.gg/seal-400.png","embeds":[`
	tail := `],"allowed_mentions":{"parse":[]}}`
	cases := map[string]struct {
		embed Embed
		want  string
	}{
		"test": {
			TestEmbed(testBase, bothEvents, fixedFinish),
			`{"title":"Test message from Forever Sixty","description":"Your webhook works. You will get a message here when a server sim finishes or a combat log you upload finishes parsing.","url":"https://foreversixty.gg/account#discord",` + common + `}`,
		},
		"sim": {
			SimFinishedEmbed(testBase, SimFinished{
				SimID: "abc123", Spec: "mage-frost", DPS: 1204.4, Title: "Phase 1 BiS check",
				CreatedAt: fixedFinish.Add(-192 * time.Second), FinishedAt: fixedFinish,
			}),
			`{"title":"Your Frost Mage sim finished","url":"https://foreversixty.gg/sim/abc123","fields":[{"name":"Result","value":"1,204 DPS","inline":true},{"name":"Run","value":"Phase 1 BiS check","inline":true},{"name":"Took","value":"3 min 12 s","inline":true}],` + footer + `"thumbnail":{"url":"https://foreversixty.gg/icons/hd/crests/mage.webp"},` + stamp + `}`,
		},
		"sim headline, no run": {
			SimFinishedEmbed(testBase, SimFinished{
				SimID: "abc123", Spec: "warrior-fury", DPS: 900, Headline: "+41 DPS from Vis'kag",
				CreatedAt: fixedFinish.Add(-48 * time.Second), FinishedAt: fixedFinish,
			}),
			`{"title":"Your Fury Warrior sim finished","url":"https://foreversixty.gg/sim/abc123","fields":[{"name":"Result","value":"+41 DPS from Vis'kag","inline":true},{"name":"Took","value":"48 s","inline":true}],` + footer + `"thumbnail":{"url":"https://foreversixty.gg/icons/hd/crests/warrior.webp"},` + stamp + `}`,
		},
		"log": {
			UploadParsedEmbed(testBase, UploadParsed{
				ReportID: "rep1", Title: "Tuesday Molten Core", Zone: "Molten Core",
				Visibility: "unlisted", Fights: 9, Kills: 7, FinishedAt: fixedFinish,
			}),
			`{"title":"Tuesday Molten Core finished parsing","url":"https://foreversixty.gg/reports/rep1","fields":[{"name":"Zone","value":"Molten Core","inline":true},{"name":"Fights","value":"9 fights · 7 kills","inline":true},{"name":"Visible to","value":"Anyone with the link","inline":true}],` + common + `}`,
		},
		"log, bare": {
			UploadParsedEmbed(testBase, UploadParsed{
				ReportID: "rep2", Visibility: "private", Fights: 1, Kills: 1, FinishedAt: fixedFinish,
			}),
			`{"title":"Your log finished parsing","url":"https://foreversixty.gg/reports/rep2","fields":[{"name":"Fights","value":"1 fight · 1 kill","inline":true},{"name":"Visible to","value":"Only you","inline":true}],` + common + `}`,
		},
	}
	for name, c := range cases {
		if got, want := wireOf(t, c.embed), head+c.want+tail; got != want {
			t.Errorf("%s:\n got %s\nwant %s", name, got, want)
		}
	}
}

func TestFightsAndVisibilityWording(t *testing.T) {
	for _, c := range []struct {
		fights, kills int
		want          string
	}{{9, 7, "9 fights · 7 kills"}, {1, 1, "1 fight · 1 kill"}, {3, 0, "3 fights · no kills"}} {
		if got := fightsLine(c.fights, c.kills); got != c.want {
			t.Errorf("%d/%d = %q, want %q", c.fights, c.kills, got, c.want)
		}
	}
	for v, want := range map[string]string{"public": "Everyone", "guild": "Your guild"} {
		e := UploadParsedEmbed(testBase, UploadParsed{Visibility: v})
		if e.Fields[len(e.Fields)-1].Value != want {
			t.Errorf("%s: %+v", v, e.Fields)
		}
	}
}

func TestRelativeBaseOmitsLinksAndThumbnail(t *testing.T) {
	e := SimFinishedEmbed("/relative", SimFinished{SimID: "x", Spec: "mage-frost", FinishedAt: fixedFinish})
	if e.URL != "" || e.Thumbnail != "" || AvatarURL("/relative") != "" {
		t.Fatalf("%+v", e)
	}
}

func TestLongMemberTextStaysInsideDiscordLimits(t *testing.T) {
	long := strings.Repeat("*", 500)
	e := UploadParsedEmbed(testBase, UploadParsed{Title: long, Zone: long, Visibility: "public"})
	if n := len([]rune(e.Title)); n > maxTitleRunes {
		t.Errorf("title %d runes", n)
	}
	s := SimFinishedEmbed(testBase, SimFinished{Spec: "mage-frost", Title: long, Headline: long})
	for _, f := range append(e.Fields, s.Fields...) {
		if len([]rune(f.Value)) > maxValueRunes {
			t.Errorf("%s: %d runes", f.Name, len([]rune(f.Value)))
		}
	}
	if run := s.Fields[1].Value; len([]rune(run)) != 200 {
		t.Errorf("run field = %d runes, want 100 escaped stars", len([]rune(run)))
	}
}

func TestEscapeMarkdown(t *testing.T) {
	for in, want := range map[string]string{
		"plain":           "plain",
		`a\b`:             `a\\b`,
		"*bold* _it_":     `\*bold\* \_it\_`,
		"~~s~~ `c` |s|":   "\\~\\~s\\~\\~ \\`c\\` \\|s\\|",
		"> quote # head":  `\> quote \# head`,
		"- item":          `\- item`,
		"a - b":           "a - b",
		"x\n- y":          "x\n\\- y",
		"[link](u)":       `\[link\](u)`,
		"@everyone":       "@everyone",
		"":                "",
		"Phase 1 BiS (a)": "Phase 1 BiS (a)",
	} {
		if got := escapeMarkdown(in); got != want {
			t.Errorf("escape(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestDurationAndThousands(t *testing.T) {
	for d, want := range map[time.Duration]string{
		48 * time.Second: "48 s", 192 * time.Second: "3 min 12 s", 180 * time.Second: "3 min",
		65 * time.Minute: "1 h 5 min", time.Hour: "1 h", -time.Second: "0 s",
	} {
		if got := duration(d); got != want {
			t.Errorf("%v = %q, want %q", d, got, want)
		}
	}
	for v, want := range map[float64]string{0: "0", 999.4: "999", 1204.5: "1,205", 1234567: "1,234,567"} {
		if got := thousands(v); got != want {
			t.Errorf("%v = %q, want %q", v, got, want)
		}
	}
}
