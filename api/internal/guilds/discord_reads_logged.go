// api/internal/guilds/discord_reads_logged.go
//
// /report's fallback (bot spec §3.3.3, amended 2026-11-05): a guild whose nights are logged but
// matched to no planned raid still has a last night. The answer is the Raids tab's newest row as
// the viewer sees it there, so the bot and the tab name the same night.
package guilds

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

// LoggedNightAnswer is the guild's newest logged night the viewer may see, read from its log
// alone: no plan, so no sign-ups, loot or raid page.
type LoggedNightAnswer struct {
	ReportID string
	// Zone is the log's zone ("Onyxia's Lair"); Title the report's own title.
	Zone, Title string
	// At is when the report was made, the date the Raids tab shows.
	At time.Time
	// Present is how many characters the log shows (the tab's "N of N present").
	Present int
	// Kills and Wipes are boss pulls only (a named encounter); TrashPacks the trash packs
	// cleared, the pulls with no named encounter, as the Raids tab and the Overview say them.
	Kills, Wipes, Deaths int
	TrashPacks           int
	DurationMS           int64
	KillOrder            []KillOrderView
	Unnamed              *UnnamedView
	Public               bool
	GuildTimezone        *string
	// PlannedRaid is the guild's planned raid of the log's zone, still waiting for a log, whose
	// night holds this one (the auto-matcher's window, plannedRaidForLog): the log was that
	// raid's and wants matching on its page. Nil when none.
	PlannedRaid *RaidRef
}

// LatestLoggedNight is the newest report on the guild's Raids tab for viewerID (0: a signed-out
// visitor), with its kill order. found is false when the tab lists none for the viewer.
func (s *Service) LatestLoggedNight(ctx context.Context, guildID, viewerID int64) (LoggedNightAnswer, bool, error) {
	member := false
	if viewerID != noViewer {
		_, verified, err := s.Accounts.GuildRank(ctx, guildID, viewerID)
		if err != nil {
			return LoggedNightAnswer{}, false, fmt.Errorf("guilds: logged night viewer: %w", err)
		}
		member = verified
	}
	page, err := s.Store.Raids(ctx, guildID, viewerID, member, 1, nil)
	if err != nil || len(page.Rows) == 0 {
		return LoggedNightAnswer{}, false, err
	}
	row := page.Rows[0]
	fights, err := s.Store.reportFights(ctx, row.ID)
	if err != nil {
		return LoggedNightAnswer{}, false, err
	}
	kills, unnamed := killOrderOf(fights)
	public, err := s.Store.reportIsPublic(ctx, row.ID)
	if err != nil {
		return LoggedNightAnswer{}, false, err
	}
	zone, err := s.GuildRaidTimezone(ctx, guildID)
	if err != nil {
		return LoggedNightAnswer{}, false, err
	}
	ans := LoggedNightAnswer{
		ReportID: row.ID, Zone: row.Zone, Title: row.Title, At: row.CreatedAt, Present: row.Raiders,
		Deaths: row.Deaths, DurationMS: row.DurationMS, KillOrder: kills, Unnamed: unnamed, Public: public,
	}
	ans.Kills, ans.Wipes, ans.TrashPacks = bossTally(fights)
	if zone != "" {
		ans.GuildTimezone = &zone
	}
	if ans.PlannedRaid, err = s.Store.plannedRaidForLog(ctx, guildID, row.ID); err != nil {
		return LoggedNightAnswer{}, false, err
	}
	return ans, true, nil
}

// plannedRaidForLog is guildID's planned or locked raid with no log, of a zone whose log name is
// the report's, whose match window (AutoMatch's: an hour before the start to six after the end)
// holds the report's anchor and that no officer cleared this log from; the start nearest the
// anchor wins. Nil when none, or when the report is not matchable evidence.
func (s *Store) plannedRaidForLog(ctx context.Context, guildID int64, reportID string) (*RaidRef, error) {
	a, ok, err := s.reportAnchor(ctx, reportID)
	if err != nil || !ok {
		return nil, err
	}
	keys := zoneKeysForLog(a.logZone)
	if len(keys) == 0 {
		return nil, nil
	}
	var ref RaidRef
	err = s.Pool.QueryRow(ctx, `
		select c.id, c.title, c.starts_at from guild_raids c
		where `+openRaidForLog("c", "$1", "$2", "$3", "$4", "$5", "$6")+`
		order by abs(extract(epoch from (c.starts_at - $3))), c.id limit 1`,
		guildID, keys, a.at, reportID, matchLeadHours, matchTailHours).Scan(&ref.ID, &ref.Title, &ref.StartsAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("guilds: planned raid for log: %w", err)
	}
	return &ref, nil
}

// bossTally counts a night's boss kills and wipes (pulls of a named encounter) and the trash packs
// it cleared (kills with none). A wiped trash pack counts nowhere: it is not a wipe on the night.
func bossTally(fights []reportFight) (kills, wipes, trash int) {
	for _, f := range fights {
		switch {
		case f.encounterID == nil || *f.encounterID <= 0:
			if f.kill {
				trash++
			}
		case f.kill:
			kills++
		default:
			wipes++
		}
	}
	return kills, wipes, trash
}
