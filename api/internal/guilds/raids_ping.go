// api/internal/guilds/raids_ping.go
//
// POST /v1/guilds/{id}/raids/{raid}/ping (spec 2026-11-05-web-layout §7.1 "Ping on Discord",
// D16): the bot posts one message in the raid's channel @mentioning the raiders who have not
// answered and have linked Discord. An optional body {"role": "tank"|"healer"|"dps"} mentions
// only the not-signed raiders of that role ("Mention the 2 healers"); no body mentions them all.
// No DMs. The throttle is per member: a ping mentions only the raiders not mentioned for this
// raid in the last six hours (pingEvery, raids_ping_claim.go), and is 429 when that leaves nobody.
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
	Posted         bool     `json:"posted"`
	Mentioned      int      `json:"mentioned"`
	NoDiscord      int      `json:"no_discord"`
	NoDiscordNames []string `json:"no_discord_names"`
	NextAllowedAt  *string  `json:"next_allowed_at"`
	// RecentlyMentioned is the linked raiders left out because a ping mentioned them for this
	// raid in the last six hours.
	RecentlyMentioned int          `json:"recently_mentioned"`
	targets           []pingTarget // the raiders the message names
	at                time.Time    // the claim's time
}

// pingTarget is one raider a ping may mention: the account and its Discord user.
type pingTarget struct {
	userID    int64
	discordID string
}

// mentionIDs is the targets' Discord users, in order.
func (p PingResult) mentionIDs() []string {
	out := make([]string, len(p.targets))
	for i, t := range p.targets {
		out[i] = t.discordID
	}
	return out
}

// discordSnowflake is a Discord user id; anything else is never put in a message.
var discordSnowflake = regexp.MustCompile(`^[0-9]{15,21}$`)

// maxPingMentions is Discord's cap on the users one message's allowed_mentions may name.
const maxPingMentions = 100

// The ping's refusals.
const (
	pingNotConnected = "Connect the Forever Sixty bot to a channel to ping on Discord."
	pingRefused      = "Discord did not take the message. Try again in a minute."
	pingThrottled    = "Everyone left to ping was mentioned in the last 6 hours."
)

// splitPingTargets is the rows' linked Discord users (capped at maxPingMentions) and the names
// of the rows with none.
func splitPingTargets(rows []NotSignedRaider) PingResult {
	out := PingResult{NoDiscordNames: []string{}}
	for _, r := range rows {
		if r.DiscordLinked && discordSnowflake.MatchString(r.discordUserID) && len(out.targets) < maxPingMentions {
			out.targets = append(out.targets, pingTarget{userID: r.userID, discordID: r.discordUserID})
			continue
		}
		out.NoDiscordNames = append(out.NoDiscordNames, r.Name)
	}
	out.Mentioned, out.NoDiscord = len(out.targets), len(out.NoDiscordNames)
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

// sendPing claims the targets not mentioned in the last six hours, posts the message and
// answers; a post that did not land gives the claim back.
func (s *Service) sendPing(w http.ResponseWriter, r *http.Request, guildID int64, raid Raid, view RaidView, res PingResult) {
	ctx := r.Context()
	// timestamptz keeps microseconds: the release matches the stored value exactly.
	res.at = time.Now().UTC().Truncate(time.Microsecond)
	claim, err := s.Store.claimMentions(ctx, raid.ID, auth.ActorFrom(ctx).UserID, res.targets, res.at)
	if err != nil {
		s.fail(w, r, "raid_ping", err, raidFailMessage)
		return
	}
	if len(claim.kept) == 0 {
		writePingThrottled(w, r, claim.earliest)
		return
	}
	res.targets, res.Mentioned, res.RecentlyMentioned = claim.kept, len(claim.kept), len(claim.skipped)
	g, err := s.Store.getGuild(ctx, guildID)
	if err != nil {
		s.releaseMentions(ctx, raid.ID, res.at, claim)
		s.fail(w, r, "raid_ping", err, raidFailMessage)
		return
	}
	url := strings.TrimRight(s.PublicBaseURL, "/") + RaidPagePath(GuildPagePath(g), raid.ID)
	postCtx, cancel := s.postContext(ctx)
	defer cancel()
	posted, err := s.RaidPings.PingRaid(postCtx, guildID, RaidPing{
		RaidID: raid.ID, Content: pingContent(view.Title, view.StartsAt, url, res.mentionIDs()), MentionIDs: res.mentionIDs(),
	})
	if err != nil || !posted {
		s.releaseMentions(ctx, raid.ID, res.at, claim)
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
