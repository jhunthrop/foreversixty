package discordbot

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

const (
	// maxCoachLines is how many characters' lines the field carries.
	maxCoachLines = 3
	// coachLinesBudget bounds the coach read inside one /report answer, clipped by the context's own deadline.
	coachLinesBudget = 2 * time.Second
)

// CoachLines is the coach service as the bot sees it. The adapter in api/cmd/api/coach.go copies
// the coach's night lines field for field and words nothing: the copy lives here (§18).
// NightLinesFor returns nothing for a viewer without log_coach.
type CoachLines interface {
	NightLinesFor(ctx context.Context, userID int64, reportID string, budget time.Duration) ([]CoachLine, error)
}

// CoachLine is one character's what-to-fix-first, as data. Stem is the finding's stem; On of Of
// bosses had it as their first fix. Pending: the packs were not cached inside the budget, so
// only a link to the site is offered.
type CoachLine struct {
	Character, Stem string
	On, Of          int
	URL             string
	Pending         bool
}

// withCoachLines adds the Your coach field to msg's head embed: one line per character, three at
// most. A failed or empty read leaves msg as it was; /report never fails on the coach.
func (s *Service) withCoachLines(ctx context.Context, msg Message, viewer int64, reportID string) Message {
	if s.d.Coach == nil || len(msg.Embeds) == 0 {
		return msg
	}
	lines, err := s.d.Coach.NightLinesFor(ctx, viewer, reportID, coachBudget(ctx))
	if err != nil {
		s.log.Warn("discordbot", "op", "coach lines", "err", err)
		return msg
	}
	value := coachFieldValue(lines)
	if value == "" {
		return msg
	}
	embeds := append([]notify.Embed(nil), msg.Embeds...)
	head := embeds[0]
	head.Fields = append(append([]notify.EmbedField(nil), head.Fields...), notify.EmbedField{Name: coachCopy.FieldName, Value: value})
	embeds[0] = head
	msg.Embeds = embeds
	return msg
}

// coachBudget is the time left for the coach read: coachLinesBudget, or less when ctx ends sooner.
func coachBudget(ctx context.Context) time.Duration {
	if deadline, ok := ctx.Deadline(); ok {
		return max(0, min(coachLinesBudget, time.Until(deadline)))
	}
	return coachLinesBudget
}

// coachFieldValue is the field's lines, led by the character when there are two or more; "" for none.
func coachFieldValue(lines []CoachLine) string {
	if len(lines) > maxCoachLines {
		lines = lines[:maxCoachLines]
	}
	words := make([]string, 0, len(lines))
	for _, l := range lines {
		text := coachLineText(l)
		if len(lines) > 1 {
			text = fill(coachCopy.Named, "{Character}", memberText(l.Character), "{line}", text)
		}
		words = append(words, text)
	}
	return boundedLines(words)
}

// coachLineText is one line's words, from the typed value alone.
func coachLineText(l CoachLine) string {
	if l.Pending {
		return fill(coachCopy.Pending, "{url}", l.URL)
	}
	pairs := []string{"{stem}", memberText(l.Stem), "{url}", l.URL, "{n}", strconv.Itoa(l.On), "{m}", strconv.Itoa(l.Of)}
	return strings.TrimSpace(fill(coachLineTemplate(l), pairs...))
}

// coachLineTemplate is the variant for the night's boss count: one boss (or an unusable count)
// names no count, every boss says "all", otherwise "{n} of {m}".
func coachLineTemplate(l CoachLine) string {
	switch {
	case l.Of <= 1 || l.On < 1 || l.On > l.Of:
		return coachCopy.LineOne
	case l.On == l.Of:
		return coachCopy.LineAll
	default:
		return coachCopy.Line
	}
}
