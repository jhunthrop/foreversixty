// api/internal/guilds/raids_discord_facts.go
//
// What the bot's one post per raid (spec 2026-10-27-discord-bot §3.3.1) is drawn from: the
// raid's counts, the guild's name, crest and zone, and for a done raid the G2 night's Came
// figures and the matched log's kills. It is read for nobody in particular: no viewer.
package guilds

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

// RaidPostFacts is one raid as its channel post shows it. Post carries no State or Was: the
// bot derives those from Status, the clock and its own moved-from record.
type RaidPostFacts struct {
	GuildID int64
	Status  string
	Post    notify.RaidPost
	EndsAt  time.Time
	// RaidURL is the raid page; NightURL the night (the done raid's page), empty until done.
	RaidURL, NightURL string
}

// noViewer reads a raid view as nobody: no own characters, no viewer-scoped rows.
const noViewer int64 = 0

// RaidPostFacts reads raidID's post facts; ok is false when the raid no longer exists.
func (s *Service) RaidPostFacts(ctx context.Context, raidID int64) (RaidPostFacts, bool, error) {
	raid, found, err := s.Store.raidByID(ctx, raidID)
	if err != nil || !found {
		return RaidPostFacts{}, false, err
	}
	g, err := s.Store.getGuild(ctx, raid.GuildID)
	if err != nil {
		return RaidPostFacts{}, false, fmt.Errorf("guilds: raid post guild: %w", err)
	}
	view, err := s.raidViewOf(ctx, raid.GuildID, noViewer, raid)
	if err != nil {
		return RaidPostFacts{}, false, err
	}
	post := notify.RaidPost{RaidPosted: s.raidPost(g, view)}
	if view.Night != nil {
		post.Came, post.CameOf = view.Night.Came, view.Night.Signed
	}
	post.Kills = killLines(view.Report)
	facts := RaidPostFacts{
		GuildID: raid.GuildID, Status: raid.Status, Post: post, EndsAt: raid.EndsAt(), RaidURL: post.RaidURL,
	}
	if raid.Status == RaidDone {
		facts.NightURL = post.RaidURL
	}
	return facts, true, nil
}

// killLines is the report's kills in kill order; nil without a report.
func killLines(r *ReportView) []notify.KillLine {
	if r == nil {
		return nil
	}
	var out []notify.KillLine
	for _, k := range r.KillOrder {
		if k.Killed && k.KillMS != nil {
			out = append(out, notify.KillLine{Boss: k.Name, Duration: time.Duration(*k.KillMS) * time.Millisecond})
		}
	}
	return out
}

// raidByID reads one raid by its id alone.
func (s *Store) raidByID(ctx context.Context, id int64) (Raid, bool, error) {
	raid, err := scanRaid(s.Pool.QueryRow(ctx, `select `+raidColumns+` from guild_raids where id = $1`, id))
	if errors.Is(err, pgx.ErrNoRows) {
		return Raid{}, false, nil
	}
	if err != nil {
		return Raid{}, false, fmt.Errorf("guilds: raid by id: %w", err)
	}
	return raid, true, nil
}
