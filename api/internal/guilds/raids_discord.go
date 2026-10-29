// api/internal/guilds/raids_discord.go
//
// The raid posts to the guild's Discord channel (design/specs/2026-10-26-guild-raids.md
// §6.2): raid_posted on a create, raid_changed on a time change or a cancel. A post needs
// the box ticked on that save, runs after the write commits, and never fails the request.
// Counts are those at posting; the message is never edited.
package guilds

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

// RaidEvents is the guild channel's poster (notify.GuildDeliverer); nil posts nothing.
type RaidEvents interface {
	RaidPosted(ctx context.Context, guildID int64, e notify.RaidPosted)
	RaidChanged(ctx context.Context, guildID int64, e notify.RaidChanged)
	// ReportReady posts the raid-night report; posted is false when it did not reach a channel.
	ReportReady(ctx context.Context, guildID int64, e notify.ReportReady) (posted bool, err error)
	// RaidReminder sends the day-before reminder; sent is false when it reached no channel.
	RaidReminder(ctx context.Context, guildID int64, e notify.RaidPosted) (ref notify.ReminderRef, sent bool, err error)
	// RaidReminderWithdrawn deletes a sent reminder again; a failure is logged, never returned.
	RaidReminderWithdrawn(ctx context.Context, guildID int64, ref notify.ReminderRef)
}

// raidPathFormat is the raid page under the guild page's path.
const raidPathFormat = "%s/raids/%d"

// cancelRequest is the optional body of POST …/raids/{raid}/cancel.
type cancelRequest struct {
	PostToDiscord *bool `json:"post_to_discord"`
}

// ticked reports whether the save carried the box, ticked.
func ticked(box *bool) bool { return box != nil && *box }

// decodeOptionalJSON is decodeJSON for a body that may be empty.
func decodeOptionalJSON(w http.ResponseWriter, r *http.Request, v any) error {
	err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxJSONBody)).Decode(v)
	if err == nil || errors.Is(err, io.EOF) {
		return nil
	}
	httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON", nil)
	return err
}

// rememberDiscordDefaults stores the two boxes as the guild's remembered defaults when the
// save carried either, on the webhook channel and on the bot's install, so whichever posts
// reads them back. A box the save left out keeps the guild's current default. A failure is
// logged: the raid itself was saved.
func (s *Service) rememberDiscordDefaults(ctx context.Context, guildID int64, post, remind *bool) {
	if post == nil && remind == nil {
		return
	}
	current, err := s.raidDiscord(ctx, guildID)
	if err != nil {
		s.logDiscordDefaults(guildID, "read", err)
		return
	}
	p, r := current.PostDefault, current.RemindDefault
	if post != nil {
		p = *post
	}
	if remind != nil {
		r = *remind
	}
	if s.discordOn() {
		s.logDiscordDefaults(guildID, "webhook", s.Discord.SetGuildDiscordDefaults(ctx, guildID, p, r))
	}
	if s.botOn() {
		s.logDiscordDefaults(guildID, "bot", s.Bot.SetDiscordDefaults(ctx, guildID, p, r))
	}
}

func (s *Service) logDiscordDefaults(guildID int64, destination string, err error) {
	if err != nil {
		s.logger().Error("guilds", "op", "raid_discord_defaults", "guild_id", guildID, "destination", destination, "err", err)
	}
}

// raidPost is what every raid embed needs for one raid view, in the guild's terms.
func (s *Service) raidPost(g Guild, v RaidView) notify.RaidPosted {
	in := notify.RaidPosted{
		GuildName: g.Name, Title: v.Title, StartsAt: v.StartsAt,
		RaidID:  v.ID,
		RaidURL: strings.TrimRight(s.PublicBaseURL, "/") + RaidPagePath(GuildPagePath(g), v.ID),
		Counts:  notify.RaidCounts(v.Counts),
	}
	if v.GuildTimezone != nil {
		in.Timezone = *v.GuildTimezone
	}
	if crest := CrestURL(s.Store.APIBaseURL, g.ID, g.CrestKey); crest != nil {
		in.CrestURL = *crest
	}
	if v.Size != nil {
		in.Size, in.SizeKnown = *v.Size, true
	}
	return in
}

// postGuild reads the guild the posts name; ok is false (and the failure logged) when it
// cannot, or when there is no poster.
func (s *Service) postGuild(ctx context.Context, guildID int64, op string) (Guild, bool) {
	if s.RaidEvents == nil {
		return Guild{}, false
	}
	g, err := s.Store.getGuild(ctx, guildID)
	if err != nil {
		s.logger().Error("guilds", "op", op, "guild_id", guildID, "err", err)
		return Guild{}, false
	}
	return g, true
}

// raidPostBudget bounds every post one request makes, together. The posts stay in the
// request (no goroutine outlives it on Cloud Run) but on a detached context, so a client
// that disconnects does not abort them.
const raidPostBudget = 8 * time.Second

// postContext is ctx detached from its cancel and bounded by the budget.
func (s *Service) postContext(ctx context.Context) (context.Context, context.CancelFunc) {
	budget := s.raidPostBudget
	if budget == 0 {
		budget = raidPostBudget
	}
	return context.WithTimeout(context.WithoutCancel(ctx), budget)
}

// postCreated posts one raid_posted per raid, soonest first, within the budget, and returns
// the ids it posted; once the budget is spent the rest are skipped with one log line (their
// schedule rows stay due, so the schedule run posts them).
func (s *Service) postCreated(ctx context.Context, guildID int64, views []RaidView) []int64 {
	ctx, cancel := s.postContext(ctx)
	defer cancel()
	g, ok := s.postGuild(ctx, guildID, "raid_posted")
	if !ok {
		return nil
	}
	sorted := slices.SortedFunc(slices.Values(views), func(a, b RaidView) int { return a.StartsAt.Compare(b.StartsAt) })
	posted := make([]int64, 0, len(sorted))
	for i, v := range sorted {
		if ctx.Err() != nil {
			s.logSkipped(guildID, sorted[i:])
			return posted
		}
		s.RaidEvents.RaidPosted(ctx, guildID, s.raidPost(g, v))
		posted = append(posted, v.ID)
	}
	return posted
}

// logSkipped names the raids whose posts the spent budget left out.
func (s *Service) logSkipped(guildID int64, views []RaidView) {
	ids := make([]int64, len(views))
	for i, v := range views {
		ids[i] = v.ID
	}
	s.logger().Warn("guilds", "op", "raid_post_skipped", "guild_id", guildID, "raid_ids", ids)
}

// postChanged posts raid_changed for v: moved from was, or cancelled when was is nil.
func (s *Service) postChanged(ctx context.Context, guildID int64, v RaidView, was *time.Time) {
	ctx, cancel := s.postContext(ctx)
	defer cancel()
	g, ok := s.postGuild(ctx, guildID, "raid_changed")
	if !ok {
		return
	}
	s.RaidEvents.RaidChanged(ctx, guildID, notify.RaidChanged{RaidPosted: s.raidPost(g, v), Was: was, Cancelled: was == nil})
}
