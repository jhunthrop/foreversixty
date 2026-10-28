package discordbot

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

const (
	// readinessTab is the guild page's Readiness tab.
	readinessTab = "/readiness"
	// goldColor is #e5b955, a neutral answer.
	goldColor = 15055189
	// readinessLineMax bounds one failing row; the guilds read keeps the rows inside a field.
	readinessLineMax = 300
	// readinessMoreReserve is room kept in the field for the "+n more on the board" line.
	readinessMoreReserve = 40
)

// cmdReadiness answers /readiness: the next raid's readiness board for an officer, its failing
// rows first. Rank, plan and freeze are checked on every call; never shared.
func (s *Service) cmdReadiness(ctx context.Context, in Interaction) (Reply, error) {
	guildID, viewer, bound, err := s.lookupCaller(ctx, in)
	if err != nil || !bound {
		return pressReply(commonCopy.Unbound, false), err
	}
	g, err := s.d.Guilds.BotGuildFacts(ctx, guildID)
	if err != nil {
		return Reply{}, fmt.Errorf("discordbot: readiness guild: %w", err)
	}
	guildName := memberText(g.Name)
	if viewer == nil {
		return pressReply(fill(officerCopy.ReadinessUnlinked, "{Guild}", guildName), false, s.linkButton()), nil
	}
	path, err := s.d.Guilds.GuildPagePathOf(ctx, guildID)
	if err != nil {
		return Reply{}, fmt.Errorf("discordbot: readiness path: %w", err)
	}
	ans, err := s.d.Guilds.NextRaidReadiness(ctx, guildID, *viewer, s.now())
	if refusal, ok := s.readinessRefusal(err, guildName, checkoutGuildURL(s.d.PublicBaseURL, guildID)); ok {
		return refusal, nil
	}
	if err != nil {
		return Reply{}, fmt.Errorf("discordbot: readiness: %w", err)
	}
	if ans.Raid == nil {
		return pressReply(commandCopy.RaidNone, false), nil
	}
	zone, err := s.d.Guilds.GuildRaidTimezone(ctx, guildID)
	if err != nil {
		return Reply{}, fmt.Errorf("discordbot: readiness zone: %w", err)
	}
	board := s.d.PublicBaseURL + path + readinessTab
	embeds := []notify.Embed{readinessEmbed(ans, board, zone)}
	closeEmbeds(embeds, fill(guildFooter, "{Guild}", notify.Collapse(g.Name)), s.now())
	link := Component{Kind: KindButton, Style: StyleLink, Label: officerCopy.Board, URL: board}
	return Reply{Message: Message{Embeds: embeds, Components: []ActionRow{{link}}, Ephemeral: true}}, nil
}

// checkoutGuildURL opens the checkout's guild step with the Raid guild plan chosen for guildID,
// where an officer subscribes or moves the guild up.
func checkoutGuildURL(base string, guildID int64) string {
	return base + catalog.CheckoutPath + "?plan=raid_guild&interval=yearly&guild_id=" + strconv.FormatInt(guildID, 10)
}

// readinessRefusal is the sentence for a gate's refusal; ok is false for any other error.
func (s *Service) readinessRefusal(err error, guildName, billing string) (Reply, bool) {
	planLink := func(label string) Component {
		return Component{Kind: KindButton, Style: StyleLink, Label: label, URL: billing}
	}
	switch {
	case errors.Is(err, guilds.ErrNotOfficer):
		return pressReply(fill(officerCopy.NotOfficer, "{Guild}", guildName), false), true
	case errors.Is(err, guilds.ErrPlanRequired):
		return pressReply(officerCopy.PlanLapsed, false, planLink(officerCopy.Subscribe)), true
	case errors.Is(err, guilds.ErrRaidGuildRequired):
		return pressReply(officerCopy.RaidGuild, false, planLink(officerCopy.MoveToRaidGuild)), true
	case errors.Is(err, guilds.ErrClaimContested):
		return pressReply(officerCopy.Frozen, false), true
	}
	return Reply{}, false
}

// readinessEmbed is the board summary: who signed, who has something to fix, and the first ten
// failing rows in the board's words.
func readinessEmbed(ans guilds.ReadinessAnswer, board, zone string) notify.Embed {
	z := zone
	e := notify.Embed{
		Title: officerCopy.ReadinessTitle + raidTitle(ans.Raid.Title, ans.Raid.StartsAt, &z), URL: board,
		Color: goldColor, Description: readinessLead(ans),
	}
	if ans.Failing == 0 {
		return e
	}
	lines, more := fittedRows(ans.Rows, ans.More)
	if more > 0 {
		lines = append(lines, fill(officerCopy.ReadinessMore, "{n}", strconv.Itoa(more)))
	}
	e.Fields = []notify.EmbedField{{Name: officerCopy.ToFix, Value: strings.Join(lines, "\n")}}
	return e
}

// fittedRows are the rows escaped, as many as fit the field with room for the "+n more" line,
// measured after escaping; more counts the rows that did not fit with the ones never listed.
func fittedRows(rows []string, more int) ([]string, int) {
	out := make([]string, 0, len(rows)+1)
	used := 0
	for i, r := range rows {
		line := notify.EscapeWithin(notify.Collapse(r), readinessLineMax, readinessLineMax)
		size := len([]rune(line)) + 1 // the newline before it, or the "+n more" line after it
		if used+size > fieldValueMax-readinessMoreReserve {
			return out, more + len(rows) - i
		}
		out = append(out, line)
		used += size
	}
	return out, more
}

func readinessLead(ans guilds.ReadinessAnswer) string {
	n := strconv.Itoa(ans.Signed)
	switch {
	case ans.Signed == 0:
		return officerCopy.ReadinessNobodySigned
	case ans.Failing == 0:
		return officerCopy.ReadinessAllReady
	case ans.Failing == 1:
		return fill(officerCopy.ReadinessLeadOne, "{n}", n)
	default:
		return fill(officerCopy.ReadinessLead, "{n}", n, "{k}", strconv.Itoa(ans.Failing))
	}
}
