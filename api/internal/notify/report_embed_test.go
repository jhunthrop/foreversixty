package notify

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/spectable"
)

const coachSpecPath = "../../../design/specs/2026-10-28-log-coach.md"

func onyxiaReport() ReportReady {
	return ReportReady{
		RaidPosted: onyxiaThursday(),
		Kills:      []ReportKill{{"Onyxia", 102 * time.Second}},
		Wipes:      3, Came: 22, Signed: 23, LootItems: 6,
	}
}

func TestReportReadyEmbedMatchesTheCopyTable(t *testing.T) {
	want := `{"title":"Report: Onyxia's Lair · Thu 10 Dec",` +
		`"url":"https://foreversixty.gg/guild/us/normal/olympus-xxvii/raids/7#report",` +
		`"fields":[{"name":"Kills","value":"Onyxia 1:42","inline":true},{"name":"Wipes","value":"3","inline":true},` +
		`{"name":"Came","value":"22 of 23","inline":true},{"name":"Deaths","value":"0","inline":true},{"name":"Loot","value":"6 items","inline":true}],` +
		raidFoot + `"color":8377482}`
	if got := wireOf(t, ReportReadyEmbed(onyxiaReport(), fixedFinish)); got != wireHead+want+wireTail {
		t.Fatalf("\n got %s\nwant %s", got, wireHead+want+wireTail)
	}
}

// The night's summary: its time, its deaths as a count, the best parses and the loot by name.
// Nobody is named for dying.
func TestReportReadyEmbedSummarisesTheNight(t *testing.T) {
	in := onyxiaReport()
	in.Time, in.Deaths = 2*time.Hour+3*time.Minute+4*time.Second, 7
	in.TopParses = []ReportParse{{Name: "Frostquill", Metric: "dps", Boss: "Onyxia", Value: 812},
		{Name: "Doom_whisper", Metric: "dps", Boss: "Onyxia", Value: 744}}
	in.Loot = []ReportLoot{{Item: "Ring of Binding", Winner: "Doomwhisper"}}
	e := ReportReadyEmbed(in, fixedFinish)
	got := map[string]EmbedField{}
	for _, f := range e.Fields {
		got[f.Name] = f
	}
	if got["Time"].Value != "2 h 3 min" || got["Deaths"].Value != "7" {
		t.Fatalf("time %q deaths %q", got["Time"].Value, got["Deaths"].Value)
	}
	if v := got["Top parses"].Value; v != "Frostquill 812 DPS on Onyxia\nDoom\\_whisper 744 DPS on Onyxia" || got["Top parses"].Inline {
		t.Fatalf("top parses = %q", v)
	}
	if l := got["Loot"]; l.Value != "Ring of Binding to Doomwhisper" || l.Inline {
		t.Fatalf("loot = %+v", l)
	}
	in.Loot = nil
	for i := 0; i < 7; i++ {
		in.Loot = append(in.Loot, ReportLoot{Item: fmt.Sprintf("Item %d", i), Winner: "Kraggor"})
	}
	last := ReportReadyEmbed(in, fixedFinish).Fields
	if lines := strings.Split(last[len(last)-1].Value, "\n"); len(lines) != 6 || lines[5] != "+2 more" {
		t.Fatalf("loot lines = %q", lines)
	}
}

func TestReportReadyEmbedCapsKillsAndEscapes(t *testing.T) {
	in := onyxiaReport()
	in.Kills = nil
	for i := 0; i < 8; i++ {
		in.Kills = append(in.Kills, ReportKill{fmt.Sprintf("*Boss* %d", i), time.Hour + 5*time.Second})
	}
	in.LootItems = 1
	e := ReportReadyEmbed(in, fixedFinish)
	lines := strings.Split(e.Fields[0].Value, "\n")
	if len(lines) != 7 || lines[0] != `\*Boss\* 0 1:00:05` || lines[6] != "+2 more" {
		t.Fatalf("kills = %q", lines)
	}
	if loot := e.Fields[len(e.Fields)-1]; loot.Value != "1 item" {
		t.Fatalf("loot = %q", loot.Value)
	}
}

func TestReportReadyEmbedEdges(t *testing.T) {
	in := onyxiaReport()
	in.Kills, in.RaidURL = nil, "/guild/us/normal/olympus-xxvii/raids/7"
	e := ReportReadyEmbed(in, fixedFinish)
	if e.Fields[0].Value != "No kills" || e.URL != "" {
		t.Fatalf("no kills / relative link = %q %q", e.Fields[0].Value, e.URL)
	}
}

func TestReportReadyEmbed_CoachSentenceNamesNobody(t *testing.T) {
	in := onyxiaReport()
	in.CoachReady = true
	want := spectable.Collapse(spectable.Cell(t, coachSpecPath, "### 8.3 Discord", "raidNight")[0])
	if got := ReportReadyEmbed(in, fixedFinish).Description; got != want {
		t.Fatalf("description = %q, want %q", got, want)
	}
}

func TestReportReadyEmbed_NoSentenceWithoutCoachReady(t *testing.T) {
	if got := ReportReadyEmbed(onyxiaReport(), fixedFinish).Description; got != "" {
		t.Fatal(got)
	}
}
