// api/internal/guilds/raid_assign_post.go
//
// POST /v1/guilds/{id}/raids/{raid}/assignments/post: the bot posts the raid's assignments as
// one message in the raid's channel, a reply under the raid's post when it is up (the ping's
// path, RaidPinger). It names characters in plain text and pings no one; never a DM. A leader
// re-posts after an edit, so the throttle only stops a double press (assignmentPostEvery).
package guilds

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// assignmentPostEvery is the re-post throttle: one post per raid per minute.
const assignmentPostEvery = time.Minute

// maxDiscordContent is Discord's cap on one message's content.
const maxDiscordContent = 2000

// The post's refusals.
const (
	assignmentsNothing   = "Pick at least one name before posting."
	assignmentsThrottled = "You posted these a moment ago. Try again in a minute."
	assignmentsOpenSlot  = "open"
)

// AssignmentPostResult is the post's answer.
type AssignmentPostResult struct {
	Posted        bool    `json:"posted"`
	PostedAt      *string `json:"posted_at"`
	NextAllowedAt *string `json:"next_allowed_at"`
}

func (s *Service) postAssignments(w http.ResponseWriter, r *http.Request) {
	guildID, raid, ok := s.openAssignmentRaid(w, r, "raid_assignments_post")
	if !ok {
		return
	}
	if s.RaidPings == nil {
		httpx.WriteError(w, r, http.StatusConflict, "discord_not_connected", pingNotConnected, nil)
		return
	}
	view, err := s.raidViewOf(r.Context(), guildID, auth.ActorFrom(r.Context()).UserID, raid)
	if err != nil {
		s.fail(w, r, "raid_assignments_post", err, raidFailMessage)
		return
	}
	if !anyAssigned(view.Assignments) {
		httpx.WriteError(w, r, http.StatusConflict, "nothing_assigned", assignmentsNothing, nil)
		return
	}
	s.sendAssignments(w, r, guildID, view)
}

// anyAssigned reports whether any slot of the block holds a name.
func anyAssigned(a *RaidAssignments) bool {
	if a == nil {
		return false
	}
	for _, e := range a.Encounters {
		for _, slot := range e.Slots {
			if len(slot.Names) > 0 {
				return true
			}
		}
	}
	return false
}

// sendAssignments claims the throttle, posts and answers; a post that did not land gives the
// claim back.
func (s *Service) sendAssignments(w http.ResponseWriter, r *http.Request, guildID int64, view RaidView) {
	ctx := r.Context()
	at := time.Now().UTC().Truncate(time.Microsecond)
	claimed, prev, err := s.Store.claimAssignmentPost(ctx, view.ID, auth.ActorFrom(ctx).UserID, at, assignmentPostEvery)
	if err != nil {
		s.fail(w, r, "raid_assignments_post", err, raidFailMessage)
		return
	}
	if !claimed {
		writeAssignmentsThrottled(w, r, prev)
		return
	}
	postCtx, cancel := s.postContext(ctx)
	defer cancel()
	posted, err := s.RaidPings.PingRaid(postCtx, guildID, RaidPing{RaidID: view.ID, Content: assignmentsContent(view)})
	if err != nil || !posted {
		if rerr := s.Store.releaseAssignmentPost(ctx, view.ID, at, prev); rerr != nil {
			s.logger().Error("guilds", "op", "raid_assignments_release", "raid_id", view.ID, "err", rerr)
		}
		s.answerPingFailure(w, r, guildID, err)
		return
	}
	stamp, next := at.Format(time.RFC3339), at.Add(assignmentPostEvery).Format(time.RFC3339)
	s.logger().Info("guilds", "op", "raid_assignments_post", "guild_id", guildID, "raid_id", view.ID)
	httpx.WriteOK(w, r, http.StatusOK, AssignmentPostResult{Posted: true, PostedAt: &stamp, NextAllowedAt: &next})
}

// writeAssignmentsThrottled answers 429 with when the next post may go.
func writeAssignmentsThrottled(w http.ResponseWriter, r *http.Request, prev *time.Time) {
	fields := map[string]string{}
	if prev != nil {
		next := prev.Add(assignmentPostEvery).UTC()
		fields["next_allowed_at"] = next.Format(time.RFC3339)
		if wait := int(time.Until(next).Seconds()) + 1; wait > 0 {
			w.Header().Set("Retry-After", strconv.Itoa(wait))
		}
	}
	httpx.WriteError(w, r, http.StatusTooManyRequests, "assignments_throttled", assignmentsThrottled, fields)
}

// assignmentsContent is the message: the raid and its start in each reader's own time, then
// per encounter one line per slot, "open" for a slot no one holds. Cut to Discord's cap on a
// whole line.
func assignmentsContent(view RaidView) string {
	unix := strconv.FormatInt(view.StartsAt.Unix(), 10)
	lines := []string{fmt.Sprintf("Assignments for **%s** on <t:%s:F>", view.Title, unix)}
	for _, e := range view.Assignments.Encounters {
		lines = append(lines, "", "**"+e.Name+"**")
		for _, slot := range e.Slots {
			names := assignedNames(slot.Names)
			if names == "" {
				names = assignmentsOpenSlot
			}
			lines = append(lines, fmt.Sprintf("%s: %s", slot.Label, names))
		}
	}
	return joinWithin(lines, maxDiscordContent)
}

// joinWithin joins lines with newlines, dropping the lines that would pass max bytes.
func joinWithin(lines []string, max int) string {
	var b strings.Builder
	for i, line := range lines {
		extra := len(line)
		if i > 0 {
			extra++
		}
		if b.Len()+extra > max {
			break
		}
		if i > 0 {
			b.WriteByte('\n')
		}
		b.WriteString(line)
	}
	return b.String()
}
