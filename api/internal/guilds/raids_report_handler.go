// api/internal/guilds/raids_report_handler.go
//
// Lane G2: POST …/raids/{raid}/report/post, the officer's one "Post to Discord" for a done
// raid's report (R-A6). The post is claimed first by a conditional update, so two presses
// post once; a post that did not reach the channel releases the claim.
package guilds

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

func mountRaidReportRoutes(mux *http.ServeMux, s *Service, officerLimit func(http.Handler) http.Handler) {
	mux.Handle("POST /v1/guilds/{id}/raids/{raid}/report/post", officerLimit(auth.RequireSession(s.postReport)))
}

// ClaimReportPost stamps report_posted_at on a done raid that has none; ok is false when it
// is already stamped (or the raid is not done).
func (s *Store) ClaimReportPost(ctx context.Context, guildID, raidID int64) (time.Time, bool, error) {
	var at time.Time
	err := s.Pool.QueryRow(ctx, `
		update guild_raids set report_posted_at = now()
		where id = $1 and guild_id = $2 and status = 'done' and report_posted_at is null
		returning report_posted_at`, raidID, guildID).Scan(&at)
	if errors.Is(err, pgx.ErrNoRows) {
		return time.Time{}, false, nil
	}
	if err != nil {
		return time.Time{}, false, fmt.Errorf("guilds: claim report post: %w", err)
	}
	return at.UTC(), true, nil
}

// ReleaseReportPost clears the claim of a post that did not happen.
// It only clears the claim it was given: this guild's raid, still stamped with claimedAt.
func (s *Store) ReleaseReportPost(ctx context.Context, guildID, raidID int64, claimedAt time.Time) error {
	if _, err := s.Pool.Exec(ctx,
		`update guild_raids set report_posted_at = null where id = $1 and guild_id = $2 and report_posted_at = $3`,
		raidID, guildID, claimedAt); err != nil {
		return fmt.Errorf("guilds: release report post: %w", err)
	}
	return nil
}

func (s *Service) postReport(w http.ResponseWriter, r *http.Request) {
	const op = "raid_report_post"
	guildID, raid, ok := s.officerRaid(w, r, op)
	if !ok {
		return
	}
	if raid.Status != RaidDone {
		httpx.WriteError(w, r, http.StatusConflict, "conflict", "only a raid with a matched log has a report", nil)
		return
	}
	at, claimed, err := s.Store.ClaimReportPost(r.Context(), guildID, raid.ID)
	if err != nil {
		s.fail(w, r, op, err, raidFailMessage)
		return
	}
	if !claimed {
		httpx.WriteError(w, r, http.StatusConflict, "already_posted", "this report was already posted", nil)
		return
	}
	if err := s.sendReport(r, guildID, raid); err != nil {
		s.releaseReportClaim(r.Context(), guildID, raid.ID, at)
		if errors.Is(err, errReportNotPosted) {
			httpx.WriteError(w, r, http.StatusBadGateway, "post_failed", "could not post to Discord just now", nil)
			return
		}
		s.fail(w, r, op, err, "could not post that report just now")
		return
	}
	s.logger().Info("guilds", "op", op, "guild_id", guildID, "raid_id", raid.ID)
	httpx.WriteOK(w, r, http.StatusOK, map[string]time.Time{"posted_at": at})
}

// errReportNotPosted is a post that reached no channel.
var errReportNotPosted = errors.New("guilds: the report did not reach the channel")

func (s *Service) releaseReportClaim(ctx context.Context, guildID, raidID int64, claimedAt time.Time) {
	ctx, cancel := s.postContext(ctx)
	defer cancel()
	if err := s.Store.ReleaseReportPost(ctx, guildID, raidID, claimedAt); err != nil {
		s.logger().Error("guilds", "op", "raid_report_release", "guild_id", guildID, "raid_id", raidID, "err", err)
	}
}

// sendReport posts the report_ready embed on the post budget.
func (s *Service) sendReport(r *http.Request, guildID int64, raid Raid) error {
	view, err := s.raidViewOf(r.Context(), guildID, auth.ActorFrom(r.Context()).UserID, raid)
	if err != nil {
		return err
	}
	if s.RaidEvents == nil || view.Report == nil || view.Night == nil {
		return errReportNotPosted
	}
	ctx, cancel := s.postContext(r.Context())
	defer cancel()
	g, err := s.Store.getGuild(ctx, guildID)
	if err != nil {
		return err
	}
	posted, err := s.RaidEvents.ReportReady(ctx, guildID, s.reportReady(g, view))
	if err != nil {
		return fmt.Errorf("guilds: post report: %w", err)
	}
	if !posted {
		if ctx.Err() != nil {
			// The budget ended; Discord may have taken the post, so a retry can duplicate it.
			s.logger().Warn("guilds", "op", "report_post_timeout", "guild_id", guildID, "raid_id", raid.ID)
		}
		return errReportNotPosted
	}
	return nil
}

// reportReady is the embed input of a done raid's view: counts and kill times only.
func (s *Service) reportReady(g Guild, v RaidView) notify.ReportReady {
	in := notify.ReportReady{RaidPosted: s.raidPost(g, v), Wipes: v.Night.Wipes, Came: v.Night.Came, Signed: v.Night.Signed}
	in.RaidURL = strings.TrimRight(in.RaidURL, "/")
	for _, k := range v.Report.KillOrder {
		if k.Killed && k.KillMS != nil {
			in.Kills = append(in.Kills, notify.ReportKill{Boss: k.Name, Duration: time.Duration(*k.KillMS) * time.Millisecond})
		}
	}
	for _, b := range v.Report.Loot {
		in.LootItems += len(b.Awards)
	}
	return in
}
