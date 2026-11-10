// api/internal/guilds/raid_badge.go
//
// GET /v1/guilds/{id}/raid-badge: the apps' Raid tab badge in one small read (spec
// 2026-11-08-app-flows F3). A member's count is 1 while their own next raid is unanswered; an
// officer's is the open count, the next raid's unanswered accounts plus the pending applicants,
// the same number as the Raid tab's Needs you. Unanswered follows the officer's not-signed list
// (raids_not_signed.go): one per account, accounts whose characters are all under the zone's
// level left out. The next raid is the officer Overview's (the next raid, else the next run).
package guilds

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// RaidBadge is the body of GET …/raid-badge. RaidID is the raid it counts, null with none
// planned. Unanswered and Applicants are an officer's (0 for a member). SimsRank and
// LastRaidLoot are the Raid tab's Guild sims and Loot history rows (raid_badge_rows.go), null
// when there is none.
type RaidBadge struct {
	Count         int           `json:"count"`
	RaidID        *int64        `json:"raid_id"`
	YouUnanswered bool          `json:"you_unanswered"`
	Officer       bool          `json:"officer"`
	Unanswered    int           `json:"unanswered"`
	Applicants    int           `json:"applicants"`
	SimsRank      *SimsRank     `json:"sims_rank"`
	LastRaidLoot  *LastRaidLoot `json:"last_raid_loot"`
}

func (s *Service) raidBadge(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	viewerID := auth.ActorFrom(r.Context()).UserID
	rank, verified, err := s.Accounts.GuildRank(r.Context(), guildID, viewerID)
	if err != nil {
		s.fail(w, r, "raid_badge", err, "could not load the raid badge just now")
		return
	}
	if !verified {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	out, err := s.RaidBadgeOf(r.Context(), guildID, viewerID, isOfficerRank(rank))
	if err != nil {
		s.fail(w, r, "raid_badge", err, "could not load the raid badge just now")
		return
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, out)
}

// isOfficerRank reports whether a verified member's rank reads every raid's answers.
func isOfficerRank(rank string) bool { return rank == "officer" || rank == "leader" }

// RaidBadgeOf is viewerID's badge in guildID; the caller has checked the viewer is a verified
// member, and officer is their standing.
func (s *Service) RaidBadgeOf(ctx context.Context, guildID, viewerID int64, officer bool) (RaidBadge, error) {
	out := RaidBadge{Officer: officer}
	rows, raidID, err := s.nextUnanswered(ctx, guildID, viewerID)
	if err != nil {
		return out, err
	}
	out.RaidID = raidID
	if err := s.fillBadgeRows(ctx, guildID, viewerID, &out); err != nil {
		return out, err
	}
	you := accountKeyFor(viewerID)
	out.YouUnanswered = slices.ContainsFunc(rows, func(r NotSignedRaider) bool { return r.AccountKey == you })
	if !officer {
		if out.YouUnanswered {
			out.Count = 1
		}
		return out, nil
	}
	out.Unanswered = len(rows)
	if out.Applicants, err = s.Store.PendingApplicationsCount(ctx, guildID); err != nil {
		return out, err
	}
	out.Count = out.Unanswered + out.Applicants
	return out, nil
}

// nextUnanswered is the not-signed raiders of the guild's next raid (else its next run) and
// that raid's id; nil rows and id with none planned.
func (s *Service) nextUnanswered(ctx context.Context, guildID, viewerID int64) ([]NotSignedRaider, *int64, error) {
	now := time.Now()
	raid, found, err := s.Store.NextRaid(ctx, guildID, now)
	if err == nil && !found {
		raid, found, err = s.Store.NextPlanned(ctx, guildID, now)
	}
	if err != nil || !found {
		return nil, nil, err
	}
	rows, err := s.unansweredOn(ctx, guildID, viewerID, raid)
	if err != nil {
		return nil, nil, err
	}
	id := raid.ID
	return rows, &id, nil
}

// unansweredOn is the officer's not-signed raiders of raid: the badge's and raid.short's count.
func (s *Service) unansweredOn(ctx context.Context, guildID, viewerID int64, raid Raid) ([]NotSignedRaider, error) {
	views, err := s.Store.RaidViews(ctx, guildID, viewerID, []Raid{raid})
	if err != nil {
		return nil, err
	}
	facts, err := s.Store.officerRosterFacts(ctx, guildID)
	if err != nil {
		return nil, err
	}
	return notSignedRaiders(views[0], facts), nil
}
