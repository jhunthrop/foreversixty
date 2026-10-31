package discordbot

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

const (
	testRaidURL  = "https://foreversixty.gg/guild/us/normal/olympus/raids/42"
	testNightURL = testRaidURL + "#report"
)

var postStart = time.Date(2026, 12, 11, 2, 0, 0, 0, time.UTC)

func testPost(state notify.RaidPostState) notify.RaidPost {
	return notify.RaidPost{
		State: state,
		RaidPosted: notify.RaidPosted{
			RaidID: 42, GuildName: "Olympus", Title: "Onyxia", StartsAt: postStart,
			RaidURL: testRaidURL, Counts: notify.RaidCounts{In: 3}, Size: 40, SizeKnown: true,
		},
	}
}

func TestPostStatePrecedence(t *testing.T) {
	end := postStart.Add(3 * time.Hour)
	before, during, after := postStart.Add(-time.Hour), postStart.Add(time.Hour), end.Add(time.Minute)
	moved := postStart.Add(-24 * time.Hour)
	cases := []struct {
		name   string
		status string
		moved  *time.Time
		now    time.Time
		want   notify.RaidPostState
	}{
		{"open", "planned", nil, before, notify.PostOpen},
		{"locked", "locked", nil, before, notify.PostLocked},
		{"locked past end is ended", "locked", nil, after, notify.PostEnded},
		{"ended exactly at the end", "planned", nil, end, notify.PostEnded},
		{"cancelled beats done", "cancelled", nil, after, notify.PostCancelled},
		{"done beats ended", "done", nil, after, notify.PostDone},
		{"moved before start", "planned", &moved, before, notify.PostMoved},
		{"moved after start is open", "planned", &moved, during, notify.PostOpen},
		{"locked beats moved", "locked", &moved, before, notify.PostLocked},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := PostState(c.status, postStart, end, c.moved, c.now); got != c.want {
				t.Fatalf("got %s want %s", got, c.want)
			}
		})
	}
}

func TestPostMessageRows(t *testing.T) {
	now := postStart
	for _, state := range []notify.RaidPostState{notify.PostOpen, notify.PostMoved} {
		row := PostMessage(testPost(state), testRaidURL, testNightURL, now).Components[0]
		wantStyles := []ButtonStyle{StyleSecondary, StyleSecondary, StyleSecondary, StyleSecondary, StyleLink}
		wantIDs := []string{"raid:42:in", "raid:42:late", "raid:42:bench", "raid:42:out", ""}
		if len(row) != 5 {
			t.Fatalf("%s: %d components", state, len(row))
		}
		for i, c := range row {
			if c.Style != wantStyles[i] || c.CustomID != wantIDs[i] || c.Disabled {
				t.Fatalf("%s button %d = %+v", state, i, c)
			}
		}
	}
	for _, state := range []notify.RaidPostState{notify.PostLocked, notify.PostEnded} {
		row := PostMessage(testPost(state), testRaidURL, testNightURL, now).Components[0]
		for i, c := range row[:4] {
			if !c.Disabled {
				t.Fatalf("%s button %d is live", state, i)
			}
		}
		if row[4].Disabled || row[4].URL != testRaidURL {
			t.Fatalf("%s link = %+v", state, row[4])
		}
	}
	cancelled := PostMessage(testPost(notify.PostCancelled), testRaidURL, testNightURL, now).Components[0]
	if len(cancelled) != 1 || cancelled[0].Label != "Raid page" || cancelled[0].URL != testRaidURL {
		t.Fatalf("cancelled row = %+v", cancelled)
	}
	done := PostMessage(testPost(notify.PostDone), testRaidURL, testNightURL, now).Components[0]
	if len(done) != 1 || done[0].Label != "Open the night" || done[0].URL != testNightURL {
		t.Fatalf("done row = %+v", done)
	}
}

func TestRaidCustomIDRoundTrip(t *testing.T) {
	id := raidCustomID(42, "late")
	if raidID, state, ok := parseRaidCustomID(id); !ok || raidID != 42 || state != "late" {
		t.Fatalf("parse %q = %d %q %v", id, raidID, state, ok)
	}
	for _, bad := range []string{"", "raid:x:in", "raid:42", "pick:42:in", "raid:0:in", "raid:42:", "raid:42:in:x", "raid:42:maybe"} {
		if _, _, ok := parseRaidCustomID(bad); ok {
			t.Fatalf("%q parsed", bad)
		}
	}
	if len(raidCustomID(1<<62, "bench")) > 100 || pickIDPrefix == raidIDPrefix {
		t.Fatal("custom id over 100 or prefixes collide")
	}
}

func TestPostSurvivesHostileTitles(t *testing.T) {
	states := []notify.RaidPostState{notify.PostOpen, notify.PostMoved, notify.PostCancelled, notify.PostDone}
	for _, state := range states {
		t.Run(string(state), func(t *testing.T) {
			p := testPost(state)
			p.Title = strings.Repeat("*", 80) + "@everyone <@123>"
			p.Kills = hostileKills()
			checkHostile(t, PostMessage(p, testRaidURL, testNightURL, postStart))
		})
	}
}

func hostileKills() []notify.KillLine {
	kills := make([]notify.KillLine, 8)
	for i := range kills {
		kills[i] = notify.KillLine{Boss: strings.Repeat("_", 80), Duration: time.Minute}
	}
	return kills
}

func TestEmptyLinksAreLeftOut(t *testing.T) {
	open := PostMessage(testPost(notify.PostOpen), "", "", postStart)
	if len(open.Components[0]) != 4 {
		t.Fatalf("open row = %+v", open.Components[0])
	}
	for _, state := range []notify.RaidPostState{notify.PostCancelled, notify.PostDone} {
		m := PostMessage(testPost(state), "", "", postStart)
		if len(m.Components) != 0 {
			t.Fatalf("%s rows = %+v", state, m.Components)
		}
		if _, err := json.Marshal(m.edit()); err != nil {
			t.Fatalf("%s edit: %v", state, err)
		}
	}
}

func checkHostile(t *testing.T, m Message) {
	t.Helper()
	e := m.Embeds[0]
	if n := len([]rune(e.Title)); n > 256 {
		t.Fatalf("title %d runes", n)
	}
	if strings.Contains(e.Title, "**") && !strings.Contains(e.Title, `\*`) {
		t.Fatalf("title not escaped: %q", e.Title)
	}
	total := len([]rune(e.Title + e.Description + e.Footer))
	for _, f := range e.Fields {
		total += len([]rune(f.Name + f.Value))
	}
	if total > 6000 {
		t.Fatalf("embed %d chars", total)
	}
	raw, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	var wire struct {
		AllowedMentions struct {
			Parse []string `json:"parse"`
		} `json:"allowed_mentions"`
	}
	if err := json.Unmarshal(raw, &wire); err != nil || wire.AllowedMentions.Parse == nil || len(wire.AllowedMentions.Parse) != 0 {
		t.Fatalf("allowed_mentions in %s", raw)
	}
}

func TestClearedRowsStillEdit(t *testing.T) {
	body, err := json.Marshal(PostMessage(testPost(notify.PostCancelled), testRaidURL, testNightURL, postStart).edit())
	if err != nil || !strings.Contains(string(body), `"components":[{`) {
		t.Fatalf("edit body %s %v", body, err)
	}
}
