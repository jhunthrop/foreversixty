// api/internal/guilds/raids_ping.go
//
// POST /v1/guilds/{id}/raids/{raid}/ping (spec 2026-11-05-web-layout §7.1 "Ping on Discord",
// D16): the bot posts one message in the raid's channel @mentioning the raiders who have not
// answered and have linked Discord. An optional body {"role": "tank"|"healer"|"dps"} mentions
// only the not-signed raiders of that role ("Mention the 2 healers"); no body mentions them all.
// No DMs. One ping per raid per six hours (pingEvery), whatever the role.
package guilds

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// RaidPing is the one message a ping posts: Content names MentionIDs (Discord user ids), and
// only they may be pinged by it.
type RaidPing struct {
	RaidID     int64
	Content    string
	MentionIDs []string
}

// RaidPinger posts a ping in the raid's channel (discordbot.RaidRouter): the channel the raid's
// post went to, else the bot's channel. posted is false, with no error, when the guild has no
// channel the bot can post in.
type RaidPinger interface {
	PingRaid(ctx context.Context, guildID int64, p RaidPing) (posted bool, err error)
}

// PingResult is the ping's answer.
type PingResult struct {
	Posted         bool      `json:"posted"`
	Mentioned      int       `json:"mentioned"`
	NoDiscord      int       `json:"no_discord"`
	NoDiscordNames []string  `json:"no_discord_names"`
	NextAllowedAt  *string   `json:"next_allowed_at"`
	mentionIDs     []string  // the Discord users the message names
	at             time.Time // the claim's time
}

// discordSnowflake is a Discord user id; anything else is never put in a message.
var discordSnowflake = regexp.MustCompile(`^[0-9]{15,21}$`)

// maxPingMentions is Discord's cap on the users one message's allowed_mentions may name.
const maxPingMentions = 100

// The ping's refusals.
const (
	pingNotConnected = "Connect the Forever Sixty bot to a channel to ping on Discord."
	pingRefused      = "Discord did not take the message. Try again in a minute."
	pingThrottled    = "You pinged this raid in the last 6 hours."
)

// splitPingTargets is the rows' linked Discord users (capped at maxPingMentions) and the names
// of the rows with none.
func splitPingTargets(rows []NotSignedRaider) PingResult {
	out := PingResult{NoDiscordNames: []string{}}
	for _, r := range rows {
		if r.DiscordLinked && discordSnowflake.MatchString(r.discordUserID) && len(out.mentionIDs) < maxPingMentions {
			out.mentionIDs = append(out.mentionIDs, r.discordUserID)
			continue
		}
		out.NoDiscordNames = append(out.NoDiscordNames, r.Name)
	}
	out.Mentioned, out.NoDiscord = len(out.mentionIDs), len(out.NoDiscordNames)
	return out
}

// pingRequest is the ping's optional body.
type pingRequest struct {
	Role string `json:"role"`
}

// pingRole reads the optional role ("" for every role); it answers 400 for a body that is not
// JSON and 422 for an unknown role, and ok is false then.
func pingRole(w http.ResponseWriter, r *http.Request) (string, bool) {
	var req pingRequest
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxJSONBody)).Decode(&req); err != nil && !errors.Is(err, io.EOF) {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON", nil)
		return "", false
	}
	switch req.Role {
	case "", RoleTank, RoleHealer, RoleDPS:
		return req.Role, true
	}
	writeInvalid(w, r, map[string]string{"role": "Pick tank, healer or dps."})
	return "", false
}

// rowsOfRole is the rows of role; every row for "".
func rowsOfRole(rows []NotSignedRaider, role string) []NotSignedRaider {
	if role == "" {
		return rows
	}
	out := []NotSignedRaider{}
	for _, r := range rows {
		if r.Role == role {
			out = append(out, r)
		}
	}
	return out
}

// pingContent is the message: the raid, its start in each reader's own time, the mentions and
// the link to sign up.
func pingContent(title string, startsAt time.Time, raidURL string, ids []string) string {
	mentions := make([]string, len(ids))
	for i, id := range ids {
		mentions[i] = "<@" + id + ">"
	}
	unix := strconv.FormatInt(startsAt.Unix(), 10)
	return fmt.Sprintf("Still need your answer for **%s** on <t:%s:F> (<t:%s:R>): %s\nSign up here: %s",
		title, unix, unix, strings.Join(mentions, " "), raidURL)
}

func (s *Service) pingRaid(w http.ResponseWriter, r *http.Request) {
	guildID, raid, ok := s.officerRaid(w, r, "raid_ping")
	if !ok {
		return
	}
	if raid.Status != RaidPlanned && raid.Status != RaidLocked {
		httpx.WriteError(w, r, http.StatusConflict, "conflict", "That raid is over, so there is no one to ping.", nil)
		return
	}
	if s.RaidPings == nil {
		httpx.WriteError(w, r, http.StatusConflict, "discord_not_connected", pingNotConnected, nil)
		return
	}
	role, ok := pingRole(w, r)
	if !ok {
		return
	}
	actor := auth.ActorFrom(r.Context()).UserID
	view, err := s.raidViewOf(r.Context(), guildID, actor, raid)
	if err != nil {
		s.fail(w, r, "raid_ping", err, raidFailMessage)
		return
	}
	res := splitPingTargets(rowsOfRole(*view.NotSignedRaiders, role))
	if res.Mentioned == 0 {
		httpx.WriteOK(w, r, http.StatusOK, res)
		return
	}
	s.sendPing(w, r, guildID, raid, view, res)
}

// sendPing claims the raid's throttle, posts the message and answers; a post that did not land
// gives the claim back.
func (s *Service) sendPing(w http.ResponseWriter, r *http.Request, guildID int64, raid Raid, view RaidView, res PingResult) {
	ctx := r.Context()
	// timestamptz keeps microseconds: the release matches the stored value exactly.
	res.at = time.Now().UTC().Truncate(time.Microsecond)
	claimed, prev, err := s.Store.claimPing(ctx, raid.ID, auth.ActorFrom(ctx).UserID, res.Mentioned, res.at)
	if err != nil {
		s.fail(w, r, "raid_ping", err, raidFailMessage)
		return
	}
	if !claimed {
		writePingThrottled(w, r, prev)
		return
	}
	g, err := s.Store.getGuild(ctx, guildID)
	if err != nil {
		s.releasePing(ctx, raid.ID, res.at, prev)
		s.fail(w, r, "raid_ping", err, raidFailMessage)
		return
	}
	url := strings.TrimRight(s.PublicBaseURL, "/") + RaidPagePath(GuildPagePath(g), raid.ID)
	postCtx, cancel := s.postContext(ctx)
	defer cancel()
	posted, err := s.RaidPings.PingRaid(postCtx, guildID, RaidPing{
		RaidID: raid.ID, Content: pingContent(view.Title, view.StartsAt, url, res.mentionIDs), MentionIDs: res.mentionIDs,
	})
	if err != nil || !posted {
		s.releasePing(ctx, raid.ID, res.at, prev)
		s.answerPingFailure(w, r, guildID, err)
		return
	}
	next := res.at.Add(pingEvery).Format(time.RFC3339)
	res.Posted, res.NextAllowedAt = true, &next
	s.logger().Info("guilds", "op", "raid_ping", "guild_id", guildID, "raid_id", raid.ID, "mentioned", res.Mentioned)
	httpx.WriteOK(w, r, http.StatusOK, res)
}

// answerPingFailure answers a post that did not land: 502 when Discord refused it, 409 when the
// guild has no channel the bot can post in.
func (s *Service) answerPingFailure(w http.ResponseWriter, r *http.Request, guildID int64, err error) {
	if err != nil {
		s.logger().Warn("guilds", "op", "raid_ping", "guild_id", guildID, "err", err)
		httpx.WriteError(w, r, http.StatusBadGateway, "discord_refused", pingRefused, nil)
		return
	}
	httpx.WriteError(w, r, http.StatusConflict, "discord_not_connected", pingNotConnected, nil)
}

// writePingThrottled answers 429 with when the next ping may go, in next_allowed_at and
// Retry-After.
func writePingThrottled(w http.ResponseWriter, r *http.Request, prev *time.Time) {
	fields := map[string]string{}
	if prev != nil {
		next := prev.Add(pingEvery).UTC()
		fields["next_allowed_at"] = next.Format(time.RFC3339)
		if wait := int(time.Until(next).Seconds()) + 1; wait > 0 {
			w.Header().Set("Retry-After", strconv.Itoa(wait))
		}
	}
	httpx.WriteError(w, r, http.StatusTooManyRequests, "ping_throttled", pingThrottled, fields)
}

// claimPing stamps the raid's ping at at when its last one is pingEvery old or there is none;
// claimed is false otherwise. prev is the last ping before this call (nil: never).
func (s *Store) claimPing(ctx context.Context, raidID, by int64, mentioned int, at time.Time) (claimed bool, prev *time.Time, err error) {
	err = s.Pool.QueryRow(ctx, `
		with prev as (select last_pinged_at from guild_raid_pings where raid_id = $1),
		up as (
		  insert into guild_raid_pings (raid_id, last_pinged_at, pinged_by, mentioned) values ($1, $2, $3, $4)
		  on conflict (raid_id) do update set last_pinged_at = excluded.last_pinged_at,
		    pinged_by = excluded.pinged_by, mentioned = excluded.mentioned
		  where guild_raid_pings.last_pinged_at <= $2::timestamptz - make_interval(secs => $5)
		  returning 1)
		select exists (select 1 from up), (select last_pinged_at from prev)`,
		raidID, at, by, mentioned, pingEvery.Seconds()).Scan(&claimed, &prev)
	if err != nil {
		return false, nil, fmt.Errorf("guilds: claim ping: %w", err)
	}
	return claimed, prev, nil
}

// releasePing gives back a claim made at at: the row returns to prev, or goes when there was
// none. A failure is logged: the throttle then holds for six hours, never the other way.
func (s *Service) releasePing(ctx context.Context, raidID int64, at time.Time, prev *time.Time) {
	var err error
	if prev == nil {
		_, err = s.Store.Pool.Exec(ctx, `delete from guild_raid_pings where raid_id = $1 and last_pinged_at = $2`, raidID, at)
	} else {
		_, err = s.Store.Pool.Exec(ctx, `update guild_raid_pings set last_pinged_at = $3
			where raid_id = $1 and last_pinged_at = $2`, raidID, at, *prev)
	}
	if err != nil {
		s.logger().Error("guilds", "op", "raid_ping_release", "raid_id", raidID, "err", err)
	}
}
