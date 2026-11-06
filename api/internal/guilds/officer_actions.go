// api/internal/guilds/officer_actions.go
//
// GET /v1/guilds/{id}/officer/actions (spec 2026-11-05-web-layout §7.1 "Action needed"): the
// officer Overview's counts in one read. approvals are the characters waiting for an officer;
// not_signed is the next raid's accounts that have not answered (raids_not_signed.go);
// drops_to_award is the newest loot council night with drops (the loot council nights: what the
// council awarded, voted on or flagged) and how many of them are still open.
// applications is the pending applications (recruitment, applications.go), oldest first.
package guilds

import (
	"context"
	"fmt"
	"net/http"
	"time"

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

// ActionDrops is the newest council night with drops: Count is the drops still to award (Tracked
// is always true now; Count stays a pointer for the contract); Awarded is the night's awards.
type ActionDrops struct {
	Tracked bool   `json:"tracked"`
	Count   *int   `json:"count"`
	RaidID  *int64 `json:"raid_id"`
	Label   string `json:"label"`
	Awarded int    `json:"awarded"`
}

// ActionApplications is the applications waiting for a decision, by the applicant's character.
type ActionApplications struct {
	Count int      `json:"count"`
	Names []string `json:"names"`
}

// OfficerActions is the body of GET …/officer/actions.
type OfficerActions struct {
	Approvals    ActionApprovals    `json:"approvals"`
	Applications ActionApplications `json:"applications"`
	NotSigned    ActionNotSigned    `json:"not_signed"`
	DropsToAward *ActionDrops       `json:"drops_to_award"`
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
	if out.Applications, err = s.Store.pendingApplications(ctx, guildID); err != nil {
		return out, err
	}
	if out.NotSigned, err = s.nextNotSigned(ctx, guildID, viewerID); err != nil {
		return out, err
	}
	out.DropsToAward, err = s.Store.lastNightDrops(ctx, guildID, time.Now())
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

// pendingApplications is the applications waiting for a decision: the count, and the applicants
// by their character's name, oldest application first (the order the officer list reads them).
func (s *Store) pendingApplications(ctx context.Context, guildID int64) (ActionApplications, error) {
	count, err := s.PendingApplicationsCount(ctx, guildID)
	if err != nil || count == 0 {
		return ActionApplications{Count: count, Names: []string{}}, err
	}
	rows, err := s.Pool.Query(ctx, `
		select coalesce(nullif(c.name, ''), split_part(ga.character_key, '/', 3))
		from guild_applications ga
		left join characters c on c.key = ga.character_key
		where ga.guild_id = $1 and ga.status = 'pending'
		order by ga.created_at, ga.id`, guildID)
	if err != nil {
		return ActionApplications{}, fmt.Errorf("guilds: pending applicants: %w", err)
	}
	defer rows.Close()
	out := ActionApplications{Count: count, Names: make([]string, 0, count)}
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return ActionApplications{}, fmt.Errorf("guilds: pending applicants: %w", err)
		}
		out.Names = append(out.Names, name)
	}
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

// actionNightsScan is how many council nights Action needed looks back through for drops.
const actionNightsScan = councilNightsDefault

// lastNightDrops is the newest council night that has drops (awarded, voted on or flagged) with
// its open ones to award; nil when no night has any. A drop is known only once the council
// touched it, so a night nobody voted or flagged on has nothing to count.
func (s *Store) lastNightDrops(ctx context.Context, guildID int64, now time.Time) (*ActionDrops, error) {
	view, err := s.CouncilNights(ctx, guildID, actionNightsScan, now)
	if err != nil {
		return nil, err
	}
	for _, n := range view.Nights {
		if len(n.Drops) == 0 {
			continue
		}
		return s.nightDrops(ctx, guildID, n)
	}
	return nil, nil
}

// nightDrops is night n as Action needed counts it: the open drops, the awards, and its label.
func (s *Store) nightDrops(ctx context.Context, guildID int64, n LootNight) (*ActionDrops, error) {
	open := 0
	for _, d := range n.Drops {
		if d.State == dropOpen {
			open++
		}
	}
	tz, err := s.GuildRaidTimezone(ctx, guildID)
	if err != nil {
		return nil, fmt.Errorf("guilds: last night drops: %w", err)
	}
	at, err := time.Parse(time.RFC3339, n.At)
	if err != nil {
		return nil, fmt.Errorf("guilds: last night drops: %w", err)
	}
	label := n.Zone + ", " + raidDay(at, tz)
	return &ActionDrops{Tracked: true, Count: &open, RaidID: n.RaidID, Label: label, Awarded: len(n.Decisions)}, nil
}
