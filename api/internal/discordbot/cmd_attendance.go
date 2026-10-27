package discordbot

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

const (
	// optAttendanceName is /attendance's officers-only name option.
	optAttendanceName = "name"
	// attendanceLineDate is "Thu 10 Dec", in the guild's raid zone.
	attendanceLineDate = "Mon 2 Jan"
	// rosterTab is the guild page's Roster tab.
	rosterTab = "/roster"
)

// cmdAttendance answers /attendance: one embed per own character (main first, at most five),
// or an officer's look-up of a named one. Ephemeral, never shared.
func (s *Service) cmdAttendance(ctx context.Context, in Interaction) (Reply, error) {
	guildID, viewer, bound, err := s.lookupCaller(ctx, in)
	if err != nil || !bound {
		return pressReply(commonCopy.Unbound, false), err
	}
	if viewer == nil {
		return pressReply(officerCopy.AttendanceUnlinked, false, s.linkButton()), nil
	}
	name, _ := in.Option(optAttendanceName)
	g, err := s.d.Guilds.BotGuildFacts(ctx, guildID)
	if err != nil {
		return Reply{}, fmt.Errorf("discordbot: attendance guild: %w", err)
	}
	guildName := memberText(g.Name)
	recs, err := s.d.Guilds.AttendanceFor(ctx, guildID, *viewer, name)
	switch {
	case errors.Is(err, guilds.ErrNotMember):
		return pressReply(fill(officerCopy.AttendanceMembers, "{Guild}", guildName), false), nil
	case errors.Is(err, guilds.ErrOfficerOnly):
		return pressReply(officerCopy.OfficersOnly, false), nil
	case errors.Is(err, guilds.ErrNoSuchMember):
		return pressReply(fill(officerCopy.NoMember, "{Guild}", guildName), false), nil
	case err != nil:
		return Reply{}, fmt.Errorf("discordbot: attendance: %w", err)
	}
	if noRaids(recs) {
		return pressReply(officerCopy.NoRaids, false), nil
	}
	return s.attendanceReply(ctx, guildID, g.Name, recs)
}

// noRaids is true when no record has a night or a raid line.
func noRaids(recs []guilds.AttendanceRecord) bool {
	for _, r := range recs {
		if r.Chip != "" || len(r.Last) > 0 {
			return false
		}
	}
	return true
}

func (s *Service) attendanceReply(ctx context.Context, guildID int64, guildName string, recs []guilds.AttendanceRecord) (Reply, error) {
	path, err := s.d.Guilds.GuildPagePathOf(ctx, guildID)
	if err != nil {
		return Reply{}, fmt.Errorf("discordbot: attendance path: %w", err)
	}
	zone, err := s.d.Guilds.GuildRaidTimezone(ctx, guildID)
	if err != nil {
		return Reply{}, fmt.Errorf("discordbot: attendance zone: %w", err)
	}
	roster := s.d.PublicBaseURL + path + rosterTab
	embeds := make([]notify.Embed, 0, len(recs))
	for _, r := range recs {
		embeds = append(embeds, attendanceEmbed(r, roster, zone))
	}
	closeEmbeds(embeds, fill(guildFooter, "{Guild}", notify.Collapse(guildName)), s.now())
	link := Component{Kind: KindButton, Style: StyleLink, Label: officerCopy.Roster, URL: roster}
	return Reply{Message: Message{Embeds: embeds, Components: []ActionRow{{link}}, Ephemeral: true}}, nil
}

// attendanceEmbed is one record: the chip, then its last raids in G2's outcome words.
func attendanceEmbed(r guilds.AttendanceRecord, roster, zone string) notify.Embed {
	e := notify.Embed{
		Title: fill(officerCopy.AttendanceTitle, "{Name}", memberText(r.Name)), URL: roster,
		Color: classColors[r.Class], Description: r.Chip,
	}
	if e.Description == "" {
		e.Description = officerCopy.NoRaids
	}
	if len(r.Last) == 0 {
		return e
	}
	loc := raidLocation(zone)
	lines := make([]string, len(r.Last))
	for i, l := range r.Last {
		lines[i] = l.StartsAt.In(loc).Format(attendanceLineDate) + " " + outcomeLine(l)
	}
	e.Fields = []notify.EmbedField{{Name: officerCopy.LastRaids, Value: strings.Join(lines, "\n")}}
	return e
}

// outcomeLine is "Present", "Late, excused", "Not in the log" or "Came as {Name}".
func outcomeLine(l guilds.AttendanceLine) string {
	word := outcomeWords[l.Outcome]
	if l.CameAs != "" {
		word = fill(cameAsWords, "{Name}", memberText(l.CameAs))
	}
	if l.Excused {
		word += excusedSuffix
	}
	return word
}

// autocompleteMember suggests /attendance's name from the roster, for officers only.
func (s *Service) autocompleteMember(ctx context.Context, in Interaction) ([]Choice, error) {
	focused, typed := in.Focused()
	if focused != optAttendanceName {
		return []Choice{}, nil
	}
	guildID, viewer, bound, err := s.lookupCaller(ctx, in)
	if err != nil || !bound || viewer == nil {
		return []Choice{}, err
	}
	names, err := s.d.Guilds.AttendanceNames(ctx, guildID, *viewer, typed, maxChoices)
	if err != nil {
		return nil, fmt.Errorf("discordbot: attendance autocomplete: %w", err)
	}
	out := make([]Choice, 0, len(names))
	for _, n := range names {
		if len([]rune(n)) <= maxChoiceRunes {
			out = append(out, Choice{Name: n, Value: n})
		}
	}
	return out, nil
}
