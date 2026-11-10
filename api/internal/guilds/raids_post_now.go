// api/internal/guilds/raids_post_now.go
//
// POST /v1/guilds/{id}/raids/{raid}/discord/post (the free calendar gate): the officer's "Post now" for an open raid
// that has not posted (the box was off, or a send failed), and "Repost" for one that has (the
// message was buried or deleted). A first post goes through the schedule row, so the run and a
// second press cannot post it twice; a repost asks the router for a fresh message, and the
// bot deletes the old one.
package guilds

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

// RaidReposter is the router's repost: the raid's post goes out again as a new message. sent
// is false when it reached no channel.
type RaidReposter interface {
	RepostRaid(ctx context.Context, guildID int64, e notify.RaidPosted) (sent bool, err error)
}

const (
	raidPostNowOp       = "raid_post_now"
	raidPostNotOpenText = "only a planned or locked raid that has not ended can post"
	raidPostNoChannel   = "connect a Discord channel first"
	raidPostBusyText    = "this raid is posting right now"
	raidPostFailedText  = "could not post to Discord just now"
)

var (
	errRaidPostNoChannel = errors.New("guilds: the raid post has no channel")
	errRaidPostBusy      = errors.New("guilds: the raid post is leased")
)

func mountRaidPostNowRoute(mux *http.ServeMux, s *Service, officerLimit func(http.Handler) http.Handler) {
	mux.Handle("POST /v1/guilds/{id}/raids/{raid}/discord/post", officerLimit(raidRouteSession(s.postRaidNow)))
}

func (s *Service) postRaidNow(w http.ResponseWriter, r *http.Request) {
	guildID, raid, ok := s.calendarRaid(w, r, raidPostNowOp)
	if !ok {
		return
	}
	now := time.Now()
	if (raid.Status != RaidPlanned && raid.Status != RaidLocked) || raid.Ended(now) {
		httpx.WriteError(w, r, http.StatusConflict, "conflict", raidPostNotOpenText, nil)
		return
	}
	reposted, err := s.postOrRepost(r.Context(), guildID, raid, now)
	if !s.answeredPostNow(w, r, err) {
		return
	}
	s.logger().Info("guilds", "op", raidPostNowOp, "guild_id", guildID, "raid_id", raid.ID, "repost", reposted)
	s.writeRaidView(w, r, raidPostNowOp, guildID, raid, nil)
}

// answeredPostNow writes the refusal for err; true when err is nil.
func (s *Service) answeredPostNow(w http.ResponseWriter, r *http.Request, err error) bool {
	switch {
	case err == nil:
		return true
	case errors.Is(err, errRaidPostNoChannel):
		httpx.WriteError(w, r, http.StatusConflict, "no_channel", raidPostNoChannel, nil)
	case errors.Is(err, errRaidPostBusy):
		httpx.WriteError(w, r, http.StatusConflict, "conflict", raidPostBusyText, nil)
	case errors.Is(err, errPostReachedNoChannel):
		httpx.WriteError(w, r, http.StatusBadGateway, "post_failed", raidPostFailedText, nil)
	default:
		s.fail(w, r, raidPostNowOp, err, raidPostFailedText)
	}
	return false
}

// postOrRepost posts an unposted raid now, or reposts a posted one; reposted reports which.
func (s *Service) postOrRepost(ctx context.Context, guildID int64, raid Raid, now time.Time) (reposted bool, err error) {
	dest, err := s.raidDiscord(ctx, guildID)
	if err != nil {
		return false, err
	}
	if s.RaidEvents == nil || !dest.Connected {
		return false, errRaidPostNoChannel
	}
	rows, err := s.Store.SchedulesOf(ctx, []int64{raid.ID})
	if err != nil {
		return false, err
	}
	row, has := rows[raid.ID]
	post, err := s.raidPostFor(ctx, guildID, raid)
	if err != nil {
		return false, err
	}
	if has && row.PostedAt != nil {
		return true, s.repost(ctx, guildID, post)
	}
	remind := row.Remind
	if !has {
		remind = s.remindChoice(ctx, guildID, nil)
	}
	return false, s.postFirst(ctx, guildID, raid, post, remind, now)
}

// raidPostFor is the raid's post as the channel shows it.
func (s *Service) raidPostFor(ctx context.Context, guildID int64, raid Raid) (notify.RaidPosted, error) {
	g, err := s.Store.getGuild(ctx, guildID)
	if err != nil {
		return notify.RaidPosted{}, err
	}
	view, err := s.raidViewOf(ctx, guildID, noViewer, raid)
	if err != nil {
		return notify.RaidPosted{}, err
	}
	return s.raidPost(g, view), nil
}

// postFirst records the row due now, leases it so the run cannot send it too, and sends it.
func (s *Service) postFirst(ctx context.Context, guildID int64, raid Raid, post notify.RaidPosted, remind bool, now time.Time) error {
	ctx = context.WithoutCancel(ctx)
	if err := s.Store.UpsertSchedule(ctx, raid.ID, guildID, now, remind); err != nil {
		return err
	}
	lease, ok, err := s.Store.leaseUnposted(ctx, raid.ID, now)
	if err != nil {
		return err
	}
	if !ok {
		return errRaidPostBusy
	}
	if err := s.sendNow(ctx, guildID, post); err != nil {
		s.releaseDue(ctx, raid.ID, lease)
		return err
	}
	s.markPosted(ctx, raid.ID, lease, raid.StartsAt, now)
	return nil
}

// sendNow sends a first post through whichever destination is live.
func (s *Service) sendNow(ctx context.Context, guildID int64, post notify.RaidPosted) error {
	ctx, cancel := s.postContext(ctx)
	defer cancel()
	sender, ok := s.RaidEvents.(RaidPostSender)
	if !ok {
		s.RaidEvents.RaidPosted(ctx, guildID, post)
		return nil
	}
	sent, err := sender.SendRaidPost(ctx, guildID, post)
	if err == nil && !sent {
		err = errPostReachedNoChannel
	}
	return err
}

// repost asks the router for a fresh message; a poster that cannot repost sends it as a first
// post (the webhook's posts are never edited, so a repost is a new one there too).
func (s *Service) repost(ctx context.Context, guildID int64, post notify.RaidPosted) error {
	reposter, ok := s.RaidEvents.(RaidReposter)
	if !ok {
		return s.sendNow(ctx, guildID, post)
	}
	ctx, cancel := s.postContext(ctx)
	defer cancel()
	sent, err := reposter.RepostRaid(ctx, guildID, post)
	if err == nil && !sent {
		err = errPostReachedNoChannel
	}
	return err
}
