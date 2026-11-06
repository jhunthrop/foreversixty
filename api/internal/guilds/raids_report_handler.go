// api/internal/guilds/raids_report_handler.go
//
// Lane G2: POST …/raids/{raid}/report/post, the officer's one "Post to Discord" for a done
// raid's report (R-A6). The post is claimed first by a conditional update, so two presses
// post once; a post that did not reach the channel releases the claim. A body of
// {"repost": true} posts an already posted report again (the officer asked for it); a repost
// that does not reach the channel puts the earlier stamp back.
package guilds

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
	"github.com/jhunthrop/foreversixty/api/internal/rating/presence"
)

func mountRaidReportRoutes(mux *http.ServeMux, s *Service, officerLimit func(http.Handler) http.Handler) {
	mux.Handle("POST /v1/guilds/{id}/raids/{raid}/report/post", officerLimit(raidRouteSession(s.postReport)))
	mountAttendanceOverrideRoutes(mux, s, officerLimit)
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

// ClaimReportRepost restamps report_posted_at on a done raid that was already posted, and
// returns the stamp it replaced; ok is false when the raid was never posted (or is not done).
func (s *Store) ClaimReportRepost(ctx context.Context, guildID, raidID int64) (time.Time, time.Time, bool, error) {
	var at, prev time.Time
	err := s.Pool.QueryRow(ctx, `
		with old as (
		  select report_posted_at from guild_raids
		  where id = $1 and guild_id = $2 and status = 'done' and report_posted_at is not null for update)
		update guild_raids g set report_posted_at = now() from old
		where g.id = $1 returning g.report_posted_at, old.report_posted_at`, raidID, guildID).Scan(&at, &prev)
	if errors.Is(err, pgx.ErrNoRows) {
		return time.Time{}, time.Time{}, false, nil
	}
	if err != nil {
		return time.Time{}, time.Time{}, false, fmt.Errorf("guilds: claim report repost: %w", err)
	}
	return at.UTC(), prev.UTC(), true, nil
}

// RestoreReportPost puts back prev over a repost claim that did not happen (only the claim it
// was given: this guild's raid, still stamped with claimedAt).
func (s *Store) RestoreReportPost(ctx context.Context, guildID, raidID int64, claimedAt, prev time.Time) error {
	if _, err := s.Pool.Exec(ctx,
		`update guild_raids set report_posted_at = $4 where id = $1 and guild_id = $2 and report_posted_at = $3`,
		raidID, guildID, claimedAt, prev); err != nil {
		return fmt.Errorf("guilds: restore report post: %w", err)
	}
	return nil
}

// matchedLogIsEvidence reports whether raidID holds a log that is evidence (not private, not
// flagged); false for a raid with no log.
func (s *Store) matchedLogIsEvidence(ctx context.Context, raidID int64) (bool, error) {
	var ok bool
	err := s.Pool.QueryRow(ctx, `
		select exists (select 1 from guild_raids g join reports r on r.id = g.report_id
		               where g.id = $1 and `+presence.EvidenceReport("r")+`)`, raidID).Scan(&ok)
	if err != nil {
		return false, fmt.Errorf("guilds: matched log evidence: %w", err)
	}
	return ok, nil
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
	guildID, raid, ok := s.matchOfficerRaid(w, r, op)
	if !ok {
		return
	}
	if raid.Status != RaidDone {
		httpx.WriteError(w, r, http.StatusConflict, "conflict", "only a raid with a matched log has a report", nil)
		return
	}
	public, err := s.Store.matchedLogIsEvidence(r.Context(), raid.ID)
	if err != nil {
		s.fail(w, r, op, err, raidFailMessage)
		return
	}
	if !public {
		// The presence rules leave such a raid with no report, so there is nothing to post.
		httpx.WriteError(w, r, http.StatusConflict, "conflict", "this raid's log is not public", nil)
		return
	}
	repost, ok := repostAsked(w, r)
	if !ok {
		return
	}
	claim, err := s.claimPost(r.Context(), guildID, raid.ID, repost)
	if err != nil {
		s.fail(w, r, op, err, raidFailMessage)
		return
	}
	if !claim.ok {
		code, msg := "already_posted", "this report was already posted"
		if repost {
			code, msg = "not_posted", "this report has not been posted yet"
		}
		httpx.WriteError(w, r, http.StatusConflict, code, msg, nil)
		return
	}
	if err := s.sendReport(r, guildID, raid); err != nil {
		s.releaseReportClaim(r.Context(), guildID, raid.ID, claim)
		if errors.Is(err, errReportNotPosted) {
			httpx.WriteError(w, r, http.StatusBadGateway, "post_failed", "could not post to Discord just now", nil)
			return
		}
		s.fail(w, r, op, err, "could not post that report just now")
		return
	}
	s.logger().Info("guilds", "op", op, "guild_id", guildID, "raid_id", raid.ID, "repost", repost)
	httpx.WriteOK(w, r, http.StatusOK, map[string]time.Time{"posted_at": claim.at})
}

// errReportNotPosted is a post that reached no channel.
var errReportNotPosted = errors.New("guilds: the report did not reach the channel")

// postClaim is a claimed post: its stamp, and for a repost the stamp it replaced.
type postClaim struct {
	at, prev time.Time
	repost   bool
	ok       bool
}

func (s *Service) claimPost(ctx context.Context, guildID, raidID int64, repost bool) (postClaim, error) {
	if repost {
		at, prev, ok, err := s.Store.ClaimReportRepost(ctx, guildID, raidID)
		return postClaim{at: at, prev: prev, repost: true, ok: ok}, err
	}
	at, ok, err := s.Store.ClaimReportPost(ctx, guildID, raidID)
	return postClaim{at: at, ok: ok}, err
}

// repostAsked reads the optional {"repost": bool} body; an empty body is a first post. ok is
// false once it has answered a malformed body.
func repostAsked(w http.ResponseWriter, r *http.Request) (bool, bool) {
	body, err := io.ReadAll(http.MaxBytesReader(w, r.Body, maxJSONBody))
	if err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON", nil)
		return false, false
	}
	if len(bytes.TrimSpace(body)) == 0 {
		return false, true
	}
	var req struct {
		Repost bool `json:"repost"`
	}
	if err := json.Unmarshal(body, &req); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON", nil)
		return false, false
	}
	return req.Repost, true
}

func (s *Service) releaseReportClaim(ctx context.Context, guildID, raidID int64, claim postClaim) {
	ctx, cancel := s.postContext(ctx)
	defer cancel()
	var err error
	if claim.repost {
		err = s.Store.RestoreReportPost(ctx, guildID, raidID, claim.at, claim.prev)
	} else {
		err = s.Store.ReleaseReportPost(ctx, guildID, raidID, claim.at)
	}
	if err != nil {
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
	posted, err := s.RaidEvents.ReportReady(ctx, guildID, s.reportReady(g, view, s.coachReady(ctx, view)))
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

// coachReady is whether the post gains the coach sentence; a failed read logs and leaves it out.
func (s *Service) coachReady(ctx context.Context, v RaidView) bool {
	if s.CoachHolders == nil || v.Match == nil {
		return false
	}
	holds, err := s.CoachHolders.AnyCoachHolder(ctx, v.Match.ReportID)
	if err != nil {
		s.logger().Warn("guilds", "op", "report_coach_holders", "raid_id", v.ID, "err", err)
		return false
	}
	return holds
}

// reportTopParses is how many best parses the post names.
const reportTopParses = 3

// reportReady is the embed input of a done raid's view (an officer's view: every row). It names
// the best parses and the loot winners, never who died (notify.ReportReady).
func (s *Service) reportReady(g Guild, v RaidView, coachReady bool) notify.ReportReady {
	in := notify.ReportReady{RaidPosted: s.raidPost(g, v), Wipes: v.Night.Wipes, Came: v.Night.Came, Signed: v.Night.Signed,
		CoachReady: coachReady, Time: time.Duration(v.Night.DurationMS) * time.Millisecond, Deaths: v.Report.Headline.Deaths}
	in.RaidURL = strings.TrimRight(in.RaidURL, "/")
	for _, k := range v.Report.KillOrder {
		if k.Killed && k.KillMS != nil {
			in.Kills = append(in.Kills, notify.ReportKill{Boss: k.Name, Duration: time.Duration(*k.KillMS) * time.Millisecond})
		}
	}
	for _, b := range v.Report.Loot {
		in.LootItems += len(b.Awards)
		for _, a := range b.Awards {
			in.Loot = append(in.Loot, notify.ReportLoot{Item: a.ItemName, Winner: a.Name})
		}
	}
	for _, p := range topParses(v.Report.Parses, reportTopParses) {
		in.TopParses = append(in.TopParses, notify.ReportParse{Name: p.Name, Metric: p.Metric, Boss: p.Encounter,
			Value: int(math.Round(p.Value))})
	}
	return in
}
