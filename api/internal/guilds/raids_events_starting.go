// api/internal/guilds/raids_events_starting.go
//
// What the scheduled officer and player raid events read beyond the raid's post facts:
// raid.short's not-answered count (the officer badge's list, raid_badge.go) and raid.starting's
// per-player line, the first thing the player's ready check (ready_check.go) misses on the
// character they signed. raid.starting goes to each In or Late account as its own event so its
// body can name that miss; without a recipient reader (s.Nudges) it goes to the raid audience
// with no miss.
package guilds

import (
	"context"
	"fmt"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/events"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

// readyMissCopy is how raid.starting names a ready check miss, by ready item key (an
// attunement key starts "attune.").
var readyMissCopy = struct {
	Flask, Food, Attune, Level string
}{
	Flask:  "No flask in your bags.",
	Food:   "No buff food in your bags.",
	Attune: "Not attuned.",
	Level:  "%s.",
}

// readyMissLine is the first miss of items as raid.starting says it; "" with none.
func readyMissLine(items []ReadyItem) string {
	for _, it := range items {
		if it.State != ReadyMiss {
			continue
		}
		switch {
		case it.Key == "flask":
			return readyMissCopy.Flask
		case it.Key == "food":
			return readyMissCopy.Food
		case strings.HasPrefix(it.Key, "attune."):
			return readyMissCopy.Attune
		case it.Key == "level" && it.Line != "":
			return fmt.Sprintf(readyMissCopy.Level, it.Line)
		}
	}
	return ""
}

// shortUnanswered is the raid's not-answered count for raid.short; unansweredUnknown (logged)
// when it cannot be read.
func (s *Service) shortUnanswered(ctx context.Context, guildID, raidID int64) int {
	raid, found, err := s.Store.raidByID(ctx, raidID)
	if err == nil && found {
		var rows []NotSignedRaider
		if rows, err = s.unansweredOn(ctx, guildID, noViewer, raid); err == nil {
			return len(rows)
		}
	}
	if err != nil {
		s.logger().Error("guilds", "op", "event_short_unanswered", "raid_id", raidID, "err", err)
	}
	return unansweredUnknown
}

// startingEvents is raid.starting, one per In or Late account with its ready check miss. A
// failed read falls back to the one raid-audience event with no miss.
func (s *Service) startingEvents(ctx context.Context, guildID int64, p notify.RaidPosted) []events.Event {
	all := events.Audience{Kind: events.ToRaidInLate, GuildID: guildID, RaidID: p.RaidID}
	plain := []events.Event{raidStartingEvent(guildID, p, startingLead, all, "")}
	if s.Nudges == nil {
		return plain
	}
	users, err := s.Nudges.Recipients(ctx, all)
	if err == nil {
		var chars map[int64]string
		if chars, err = s.Store.inLateCharacters(ctx, p.RaidID); err == nil {
			zone := zoneOf(p.Zone)
			out := make([]events.Event, 0, len(users))
			for _, uid := range users {
				out = append(out, raidStartingEvent(guildID, p, startingLead, events.Users(uid), s.readyMiss(ctx, chars[uid], &zone)))
			}
			return out
		}
	}
	s.logger().Error("guilds", "op", "event_starting", "raid_id", p.RaidID, "err", err)
	return plain
}

// readyMiss is the first miss of key's ready check for zone; "" without a character, a miss or
// a readable check (logged).
func (s *Service) readyMiss(ctx context.Context, key string, zone *Zone) string {
	if key == "" {
		return ""
	}
	check, err := s.Store.ReadyCheckOf(ctx, key, zone)
	if err != nil {
		s.logger().Error("guilds", "op", "event_starting_ready", "character_key", key, "err", err)
		return ""
	}
	return readyMissLine(check.Items)
}

// inLateCharacters is each account's character signed In (else Late) on the raid.
func (s *Store) inLateCharacters(ctx context.Context, raidID int64) (map[int64]string, error) {
	rows, err := s.Pool.Query(ctx, `
		select distinct on (signed_by) signed_by, character_key from guild_raid_signups
		where raid_id = $1 and state in ('in', 'late')
		order by signed_by, case state when 'in' then 0 else 1 end, character_key`, raidID)
	if err != nil {
		return nil, fmt.Errorf("guilds: in and late characters: %w", err)
	}
	defer rows.Close()
	out := map[int64]string{}
	for rows.Next() {
		var uid int64
		var key string
		if err := rows.Scan(&uid, &key); err != nil {
			return nil, fmt.Errorf("guilds: in and late characters: %w", err)
		}
		out[uid] = key
	}
	return out, rows.Err()
}
