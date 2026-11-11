package discordbot

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/db"
)

// seedLoggedNight is a guild-visible Onyxia log no planned raid holds: two pulls, a wipe then
// a kill, with Zulmara and Pyrewisp in both.
func seedLoggedNight(t *testing.T, w *officerWorld, at time.Time) {
	t.Helper()
	ctx := context.Background()
	if err := db.EnsureMetricsPartition(ctx, w.st.Pool, at); err != nil {
		t.Fatal(err)
	}
	if _, err := w.st.Pool.Exec(ctx, `insert into reports (id, guild_id, visibility, status, zone, log_zone, created_at)
		values ('logged-night-1', $1, 'guild', 'complete', 'Onyxia''s Lair', 'Onyxia''s Lair', $2)`, w.gid, at); err != nil {
		t.Fatal(err)
	}
	for i, kill := range []bool{false, true} {
		start := at.Add(time.Duration(i) * 10 * time.Minute)
		if _, err := w.st.Pool.Exec(ctx, `insert into fights (report_id, fight_index, name, encounter_id, kill, start_ms, duration_ms, deaths, players)
			values ('logged-night-1', $1, 'Onyxia', 1084, $2, $3, 102000, $4, '{}')`, i, kill, start.UnixMilli(), 2-2*i); err != nil {
			t.Fatal(err)
		}
		for _, p := range [][2]string{{"us/normal/zulmara", "Zulmara"}, {"us/normal/pyrewisp", "Pyrewisp"}} {
			if _, err := w.st.Pool.Exec(ctx, `insert into fight_metrics (report_id, fight_index, player_key, player_name, class, spec, role,
				metric_dps, deaths, encounter_id, duration_ms, kill, fought_at)
				values ('logged-night-1', $1, $2, $3, 'mage', 'fire', 'dps', 100, 0, 1084, 102000, $4, $5)`,
				i, p[0], p[1], kill, start); err != nil {
				t.Fatal(err)
			}
		}
	}
}

// A guild whose night matches no planned raid still answers a bare /report: the newest night
// on its Raids tab, linking the report page, with no Share and no loot.
func TestReportFallsBackToTheLatestLoggedNight(t *testing.T) {
	w := newOfficerWorld(t)
	w.addChar(t, w.gid, w.uid, "us/normal/zulmara", "Zulmara", true)
	seedLoggedNight(t, w, time.Date(2026, 12, 10, 20, 0, 0, 0, time.UTC))

	rep := w.run(t, pressDiscord, "report")
	if len(rep.Message.Embeds) != 1 {
		t.Fatalf("embeds = %+v, want the head alone", rep.Message.Embeds)
	}
	head := rep.Message.Embeds[0]
	if head.Title != officerCopy.ReportPrefix+"Onyxia's Lair · Thu 10 Dec" {
		t.Fatalf("title = %q", head.Title)
	}
	want := "2 present · 1 kill · 1 wipe · 2 deaths · 12 min\n" + officerCopy.LoggedNote
	if head.Description != want {
		t.Fatalf("description = %q, want %q", head.Description, want)
	}
	if got := headFields(rep)[officerCopy.KillOrder]; got != "1. Onyxia 1:42 after 1 wipe" {
		t.Fatalf("kill order = %q", got)
	}
	if b := rep.Message.Components[0][0]; b.URL != pressPublic+"/reports/logged-night-1" || b.Label != officerCopy.OpenReport {
		t.Fatalf("button = %+v", b)
	}
	if hasShare(rep) || !rep.Message.Ephemeral {
		t.Fatal("the fallback offers Share, or is public")
	}
}

// A planned Onyxia raid still waiting for a log that night: the note says to match the log on
// the raid page, linking it, instead of "no planned raid" (guild round 3).
func TestReportNamesTheUnmatchedRaidPlannedThatNight(t *testing.T) {
	w := newOfficerWorld(t)
	w.addChar(t, w.gid, w.uid, "us/normal/zulmara", "Zulmara", true)
	at := time.Date(2026, 12, 10, 20, 0, 0, 0, time.UTC)
	w.addRaid(t, w.gid, at.Add(-15*time.Minute))
	seedLoggedNight(t, w, at)

	head := w.run(t, pressDiscord, "report").Message.Embeds[0]
	_, note, _ := strings.Cut(head.Description, "\n")
	const lead = "Onyxia's Lair was planned that night: match the log on [the raid page]("
	if !strings.HasPrefix(note, lead) || !strings.HasSuffix(note, ").") {
		t.Fatalf("note = %q, want %q…", note, lead)
	}
	if got := plannedNote("Onyxia's Lair", ""); got != "Onyxia's Lair was planned that night: match the log on the raid page." {
		t.Fatalf("unlinked note = %q", got)
	}
}

// A matched night with no awards says so in the site's words.
func TestLootLineWithNoAwards(t *testing.T) {
	if got := lootLine(0); got != "No loot awarded." {
		t.Fatalf("lootLine(0) = %q", got)
	}
}

// With no log at all, /report still says so.
func TestReportNoneWithNoLoggedNight(t *testing.T) {
	w := newOfficerWorld(t)
	w.addChar(t, w.gid, w.uid, "us/normal/zulmara", "Zulmara", true)
	rep := w.run(t, pressDiscord, "report")
	if rep.Message.Content != officerCopy.ReportNone {
		t.Fatalf("content = %q", rep.Message.Content)
	}
}
