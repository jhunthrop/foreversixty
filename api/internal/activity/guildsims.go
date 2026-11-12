package activity

import (
	"context"
	"fmt"
	"log/slog"
	"math"
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/events"
	"github.com/jhunthrop/foreversixty/api/internal/textfmt"
)

// postChange is the smallest move in a member's figure that posts a new row (ruling P7).
const postChange = 0.02

// GuildSimWriter writes a guild member's sim to their guild's feed when it finishes (ruling
// P7). The sim job calls it; only a sim linked to a character (character_sims) of a verified,
// non-anonymised guild member writes anything.
type GuildSimWriter struct {
	Store *Store
	// Events, when set, receives guildsim.passed for each member the sim moved above
	// (guildsim_passed.go); PublicBaseURL roots its link.
	Events        events.Publisher
	PublicBaseURL string
	Log           *slog.Logger
	Now           func() time.Time
}

type simFacts struct {
	key, name, class, spec, preset, role, guild string
	guildID                                     int64
	band, level                                 int
	value                                       float64
	last                                        *float64
	at                                          time.Time
}

const simFactsSQL = `select cs.character_key, c.name, lower(coalesce(c.class, '')), cs.spec, cs.preset, cs.role, cs.band,
	coalesce(c.level, 0), gc.guild_id, g.name, sm.dps_mean, coalesce(sm.finished_at, now()),
	(select p.posted_value from character_sims p where p.character_key = cs.character_key and p.sim_id <> cs.sim_id
	   and p.role = cs.role and p.band = cs.band and p.preset = cs.preset and p.posted_value is not null
	 order by p.created_at desc limit 1)
	from character_sims cs
	join sims sm on sm.id = cs.sim_id and sm.state = 'done'
	join characters c on c.key = cs.character_key
	join guild_characters gc on gc.character_key = cs.character_key and gc.verified_at is not null
	join guilds g on g.id = gc.guild_id
	left join users u on u.id = c.user_id
	where cs.sim_id = $1 and not coalesce(u.anonymize, false)`

// SimFinished writes the row when the sim earns one. A sim whose headline figure is 0 (a healer
// with nothing to heal) measured nothing and earns no row, no ladder move and no posted value.
// Failures are logged: the sim is done.
func (w *GuildSimWriter) SimFinished(ctx context.Context, simID string) {
	if w == nil || w.Store == nil {
		return
	}
	if err := w.write(ctx, simID); err != nil {
		w.logger().Error("activity", "op", "guild sim", "sim", simID, "err", err)
	}
}

func (w *GuildSimWriter) write(ctx context.Context, simID string) error {
	var f simFacts
	err := w.Store.Pool.QueryRow(ctx, simFactsSQL, simID).Scan(&f.key, &f.name, &f.class, &f.spec, &f.preset, &f.role,
		&f.band, &f.level, &f.guildID, &f.guild, &f.value, &f.at, &f.last)
	if isNoRows(err) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("sim facts: %w", err)
	}
	if f.value <= 0 {
		return nil
	}
	ladder, err := w.Store.GuildLadder(ctx, LadderGroup{f.guildID, f.role, f.band, f.preset}, w.now())
	if err != nil {
		return err
	}
	if err := w.publishPassed(ctx, f, simID, ladder); err != nil {
		return err
	}
	if f.last != nil && *f.last > 0 && math.Abs(f.value-*f.last)/(*f.last) < postChange {
		return nil
	}
	if err := w.Store.Upsert(ctx, guildSimEvent(f, simID, ladder), false); err != nil {
		return err
	}
	_, err = w.Store.Pool.Exec(ctx, `update character_sims set posted_value = $2 where sim_id = $1`, simID, f.value)
	return err
}

func guildSimEvent(f simFacts, simID string, ladder []LadderRow) Event {
	members, unit := RoleWords(f.role)
	figure := textfmt.Figure(f.value, f.value)
	place := ""
	for _, r := range ladder {
		if r.Key == f.key && len(ladder) >= ladderRankWords {
			place = fmt.Sprintf(", #%d of %d %s", r.Rank, len(ladder), members)
		}
	}
	words := LevelWords(f.level, f.band)
	if f.preset != barePreset {
		words += " " + f.preset
	}
	return Event{
		Kind: KindGuildSim, OccurredAt: f.at, Spec: f.spec, Band: f.band, GuildID: f.guildID, CharacterKey: f.key,
		Crest: f.class, Href: "/sim/" + simID,
		Segments: []Segment{Name(f.name, f.class), Text(fmt.Sprintf(" simmed %s %s%s in ", figure, unit, place)),
			Strong(f.guild), Text(fmt.Sprintf(" (%s, %s)", SpecName(f.spec), words))},
		DedupeKey: "guild-sim:" + strconv.FormatInt(f.guildID, 10) + ":" + f.key + ":" + f.at.UTC().Format("20060102"),
	}
}

func (w *GuildSimWriter) now() time.Time {
	if w.Now != nil {
		return w.Now().UTC()
	}
	return time.Now().UTC()
}

func (w *GuildSimWriter) logger() *slog.Logger {
	if w.Log != nil {
		return w.Log
	}
	return slog.Default()
}
