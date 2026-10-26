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
	"fmt"
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

// rememberPostDefault stores the box as the guild's remembered default when the save
// carried it. A failure is logged: the raid itself was saved.
func (s *Service) rememberPostDefault(ctx context.Context, guildID int64, box *bool) {
	if box == nil || !s.discordOn() {
		return
	}
	if err := s.Discord.SetGuildPostDefault(ctx, guildID, *box); err != nil {
		s.logger().Error("guilds", "op", "raid_post_default", "guild_id", guildID, "err", err)
	}
}

// raidPost is what every raid embed needs for one raid view, in the guild's terms.
func (s *Service) raidPost(g Guild, v RaidView) notify.RaidPosted {
	in := notify.RaidPosted{
		GuildName: g.Name, Title: v.Title, StartsAt: v.StartsAt,
		RaidURL:  strings.TrimRight(s.PublicBaseURL, "/") + fmt.Sprintf(raidPathFormat, GuildPagePath(g), v.ID),
		SignedUp: v.Counts.In + v.Counts.Late,
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

// postCreated posts one raid_posted per created raid, soonest first.
func (s *Service) postCreated(ctx context.Context, guildID int64, views []RaidView) {
	g, ok := s.postGuild(ctx, guildID, "raid_posted")
	if !ok {
		return
	}
	sorted := slices.SortedFunc(slices.Values(views), func(a, b RaidView) int { return a.StartsAt.Compare(b.StartsAt) })
	for _, v := range sorted {
		s.RaidEvents.RaidPosted(ctx, guildID, s.raidPost(g, v))
	}
}

// postChanged posts raid_changed for v: moved from was, or cancelled when was is nil.
func (s *Service) postChanged(ctx context.Context, guildID int64, v RaidView, was *time.Time) {
	g, ok := s.postGuild(ctx, guildID, "raid_changed")
	if !ok {
		return
	}
	s.RaidEvents.RaidChanged(ctx, guildID, notify.RaidChanged{RaidPosted: s.raidPost(g, v), Was: was, Cancelled: was == nil})
}
