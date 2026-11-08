package activity

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/zoneicons"
)

// GuildRaidWriter keeps one Your guild row per planned raid (ruling P3): written when the
// raid is posted and rewritten with its counts as members sign up. guilds.Service calls it.
type GuildRaidWriter struct {
	Store *Store
	Log   *slog.Logger
}

type raidFacts struct {
	id, guildID         int64
	guild, title, zone  string
	region, ruleset     string
	tz                  string
	startsAt, createdAt time.Time
	status              string
	in, late            int
}

// RaidChanged writes or refreshes the raid's row. A raid that is not planned or locked
// writes nothing. Failures are logged: the raid write already succeeded.
func (w *GuildRaidWriter) RaidChanged(ctx context.Context, raidID int64) {
	if w == nil || w.Store == nil {
		return
	}
	f, err := w.facts(ctx, raidID)
	if err != nil {
		w.logger().Error("activity", "op", "guild raid", "raid", raidID, "err", err)
		return
	}
	if f.status != "planned" && f.status != "locked" {
		return
	}
	if err := w.Store.Upsert(ctx, guildRaidEvent(f), true); err != nil {
		w.logger().Error("activity", "op", "guild raid", "raid", raidID, "err", err)
	}
}

func (w *GuildRaidWriter) facts(ctx context.Context, raidID int64) (raidFacts, error) {
	f := raidFacts{id: raidID}
	err := w.Store.Pool.QueryRow(ctx,
		`select g.id, g.name, g.region, g.ruleset, r.title, r.zone, coalesce(g.raid_timezone, ''), r.starts_at, r.created_at, r.status,
		        (select count(*) from guild_raid_signups s where s.raid_id = r.id and s.state = 'in'),
		        (select count(*) from guild_raid_signups s where s.raid_id = r.id and s.state = 'late')
		 from guild_raids r join guilds g on g.id = r.guild_id where r.id = $1`, raidID).
		Scan(&f.guildID, &f.guild, &f.region, &f.ruleset, &f.title, &f.zone, &f.tz, &f.startsAt, &f.createdAt, &f.status, &f.in, &f.late)
	if err != nil {
		return f, fmt.Errorf("activity: raid facts %d: %w", raidID, err)
	}
	return f, nil
}

func guildRaidEvent(f raidFacts) Event {
	loc, err := time.LoadLocation(f.tz)
	if err != nil || f.tz == "" {
		loc = time.UTC
	}
	day := f.startsAt.In(loc).Format("Mon 2 Jan")
	icon := zoneicons.For(f.zone)
	return Event{
		Kind: KindGuildRaid, OccurredAt: f.createdAt, GuildID: f.guildID, Icon: icon,
		Href: guildPath(f.region, f.ruleset, f.guild) + "/raids/" + strconv.FormatInt(f.id, 10),
		Segments: []Segment{Strong(f.guild), Text(" posted "), Strong(f.title),
			Text(fmt.Sprintf(", %s: %d in, %d late", day, f.in, f.late))},
		DedupeKey: "guild-raid:" + strconv.FormatInt(f.id, 10),
	}
}

func (w *GuildRaidWriter) logger() *slog.Logger {
	if w.Log != nil {
		return w.Log
	}
	return slog.Default()
}

// guildPath is guilds.GuildPagePath for a guild known by its parts.
func guildPath(region, ruleset, name string) string {
	return "/guild/" + strings.ToLower(region) + "/" + strings.ToLower(ruleset) + "/" + url.PathEscape(character.Slug(name))
}
