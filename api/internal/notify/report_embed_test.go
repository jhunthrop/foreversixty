package notify

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func onyxiaReport() ReportReady {
	return ReportReady{
		RaidPosted: onyxiaThursday(),
		Kills:      []ReportKill{{"Onyxia", 102 * time.Second}},
		Wipes:      3, Came: 22, Signed: 23, LootItems: 6,
	}
}

func TestReportReadyEmbedMatchesTheCopyTable(t *testing.T) {
	want := `{"title":"Report: Onyxia's Lair · Thursday 10 December","description":"Kill times, loot and the biggest changes on Forever Sixty.",` +
		`"url":"https://foreversixty.gg/guild/us/normal/olympus-xxvii/raids/7#report",` +
		`"fields":[{"name":"Kills","value":"Onyxia 1:42","inline":true},{"name":"Wipes","value":"3","inline":true},` +
		`{"name":"Came","value":"22 of 23","inline":true},{"name":"Loot","value":"6 items","inline":true}],` +
		raidFoot + `"color":8377482}`
	if got := wireOf(t, ReportReadyEmbed(onyxiaReport(), fixedFinish)); got != wireHead+want+wireTail {
		t.Fatalf("\n got %s\nwant %s", got, wireHead+want+wireTail)
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
	if e.Fields[3].Value != "1 item" {
		t.Fatalf("loot = %q", e.Fields[3].Value)
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
