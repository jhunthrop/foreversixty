// api/internal/guilds/officer_actions.go
//
// GET /v1/guilds/{id}/officer/actions (spec 2026-11-05-web-layout §7.1 "Action needed"): the
// officer Overview's counts in one read. approvals are the characters waiting for an officer;
// not_signed is the next raid's accounts that have not answered (raids_not_signed.go);
// drops_to_award is the last logged night with its awards so far. The logs carry no loot, so
// what dropped is not known: tracked is false and count null until a drop record exists.
// Applications arrive with their own slice.
package guilds

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// ActionApprovals is the characters waiting for approval.
type ActionApprovals struct {
	Count int      `json:"count"`
	Names []string `json:"names"`
}

// ActionNotSigned is the next raid's accounts that have not answered; RaidID null with none planned.
type ActionNotSigned struct {
	RaidID *int64   `json:"raid_id"`
	Count  int      `json:"count"`
	Names  []string `json:"names"`
}

// ActionDrops is the last logged night's loot: Count is the drops still to award, null while no
// drop record exists (Tracked false); Awarded is the awards filed under that night so far.
type ActionDrops struct {
	Tracked bool   `json:"tracked"`
	Count   *int   `json:"count"`
	RaidID  *int64 `json:"raid_id"`
	Label   string `json:"label"`
	Awarded int    `json:"awarded"`
}

// OfficerActions is the body of GET …/officer/actions.
type OfficerActions struct {
	Approvals    ActionApprovals `json:"approvals"`
	NotSigned    ActionNotSigned `json:"not_signed"`
	DropsToAward *ActionDrops    `json:"drops_to_award"`
}

func (s *Service) officerActions(w http.ResponseWriter, r *http.Request) {
	guildID, ok := s.officerRead(w, r, "officer_actions")
	if !ok {
		return
	}
	out, err := s.officerActionsOf(r.Context(), guildID, auth.ActorFrom(r.Context()).UserID)
	if err != nil {
		s.fail(w, r, "officer_actions", err, "could not load what needs doing just now")
		return
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, out)
}

func (s *Service) officerActionsOf(ctx context.Context, guildID, viewerID int64) (OfficerActions, error) {
	var out OfficerActions
	var err error
	if out.Approvals, err = s.Store.waitingApprovals(ctx, guildID); err != nil {
		return out, err
	}
	if out.NotSigned, err = s.nextNotSigned(ctx, guildID, viewerID); err != nil {
		return out, err
	}
	out.DropsToAward, err = s.Store.lastNightDrops(ctx, guildID)
	return out, err
}

// waitingApprovals is every character of the guild no officer has approved yet, by name.
func (s *Store) waitingApprovals(ctx context.Context, guildID int64) (ActionApprovals, error) {
	rows, err := s.Pool.Query(ctx, `
		select coalesce(nullif(c.name, ''), nullif(ae.name, ''), split_part(gc.character_key, '/', 3)) as name
		from guild_characters gc
		left join characters c on c.key = gc.character_key
		left join addon_exports ae on ae.character_key = gc.character_key
		where gc.guild_id = $1 and gc.verified_at is null and gc.character_key not like 'account:%'
		order by name`, guildID)
	if err != nil {
		return ActionApprovals{}, fmt.Errorf("guilds: waiting approvals: %w", err)
	}
	defer rows.Close()
	out := ActionApprovals{Names: []string{}}
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return ActionApprovals{}, fmt.Errorf("guilds: waiting approvals: %w", err)
		}
		out.Names = append(out.Names, name)
	}
	out.Count = len(out.Names)
	return out, rows.Err()
}

// nextNotSigned is the not-signed raiders of the next raid (else the next run) as the officer's
// raid view lists them.
func (s *Service) nextNotSigned(ctx context.Context, guildID, viewerID int64) (ActionNotSigned, error) {
	out := ActionNotSigned{Names: []string{}}
	now := time.Now()
	raid, found, err := s.Store.NextRaid(ctx, guildID, now)
	if err == nil && !found {
		raid, found, err = s.Store.NextPlanned(ctx, guildID, now)
	}
	if err != nil || !found {
		return out, err
	}
	view, err := s.raidViewOf(ctx, guildID, viewerID, raid)
	if err != nil {
		return out, err
	}
	id := raid.ID
	out.RaidID = &id
	if view.NotSignedRaiders != nil {
		out.Names = raiderNames(*view.NotSignedRaiders)
	}
	out.Count = len(out.Names)
	return out, nil
}

// lastNightDrops is the guild's last logged raid with its awards so far; nil when it has none.
func (s *Store) lastNightDrops(ctx context.Context, guildID int64) (*ActionDrops, error) {
	var id int64
	var startsAt time.Time
	var zone, tz string
	var awarded int
	err := s.Pool.QueryRow(ctx, `
		select g.id, g.starts_at, g.zone, coalesce(gu.raid_timezone, ''),
		       (select count(*) from loot_awards la where la.raid_id = g.id)
		from guild_raids g join guilds gu on gu.id = g.guild_id
		where g.id in (`+loggedRaids+`)`, guildID, 1).Scan(&id, &startsAt, &zone, &tz, &awarded)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("guilds: last night drops: %w", err)
	}
	label := zoneOf(zone).Name + ", " + raidDay(startsAt, tz)
	return &ActionDrops{RaidID: &id, Label: label, Awarded: awarded}, nil
}
