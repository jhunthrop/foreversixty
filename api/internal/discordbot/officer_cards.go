package discordbot

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

// lowestAttendanceLimit is how many members the officer select lists (Discord's cap).
const lowestAttendanceLimit = maxSelectOptions

// attendanceModal is the name field behind Type a name.
func attendanceModal() *Modal { return nameModal(prefixAttendance, guideCopy.ModalAttendanceTitle) }

// attendanceOfficerRows are the rows an officer's /attendance answer gains: the lowest-attendance
// select (names as labels, the roster chip verbatim as descriptions, shown as the default) and
// the Roster row with Type a name before it. The select is left out when nobody has a night.
func (s *Service) attendanceOfficerRows(ctx context.Context, guildID int64, recs []guilds.AttendanceRecord, shown string) ([]ActionRow, error) {
	roster, err := s.rosterURL(ctx, guildID)
	if err != nil {
		return nil, err
	}
	var rows []ActionRow
	if options := lowestOptions(recs, shown); len(options) > 0 {
		rows = append(rows, ActionRow{Component{
			Kind: KindStringSelect, CustomID: cardCustomID(prefixAttendance, "pick"),
			Placeholder: guideCopy.AttendancePlaceholder, Options: options,
		}})
	}
	ask := Component{Kind: KindButton, Style: StyleSecondary, Label: guideCopy.CharTypeName,
		CustomID: cardCustomID(prefixAttendance, "ask")}
	return append(rows, ActionRow{ask, rosterLink(roster)}), nil
}

// lowestOptions are the select's options; a name listed twice (two realms) is offered once,
// because a select value must be unique.
func lowestOptions(recs []guilds.AttendanceRecord, shown string) []SelectOption {
	seen := map[string]bool{}
	options := make([]SelectOption, 0, min(len(recs), maxSelectOptions))
	for _, r := range recs {
		folded := strings.ToLower(r.Name)
		if r.Name == "" || seen[folded] || len(options) == maxSelectOptions {
			continue
		}
		seen[folded] = true
		name := truncateRunes(r.Name, maxOptionRunes)
		options = append(options, SelectOption{Label: name, Value: name,
			Description: truncateRunes(r.Chip, maxOptionRunes), Default: strings.EqualFold(r.Name, shown)})
	}
	return options
}

// withOfficerRows swaps the answer's Roster row for the officer's rows. A viewer who is not an
// officer (any more) keeps the plain answer.
func (s *Service) withOfficerRows(ctx context.Context, guildID, viewer int64, rep Reply, shown string) (Reply, error) {
	lowest, err := s.d.Guilds.LowestAttendance(ctx, guildID, viewer, lowestAttendanceLimit)
	if errors.Is(err, guilds.ErrOfficerOnly) {
		return rep, nil
	}
	if err != nil {
		return Reply{}, fmt.Errorf("discordbot: attendance lowest: %w", err)
	}
	rows, err := s.attendanceOfficerRows(ctx, guildID, lowest, shown)
	if err != nil {
		return Reply{}, err
	}
	rep.Message.Components = rows
	return rep, nil
}

// attendanceUnknown is the officer's own answer under the unknown-name line.
func (s *Service) attendanceUnknown(ctx context.Context, guildID, viewer int64, ask attendanceAsk, guildName string) (Reply, error) {
	rep, err := s.attendanceAnswer(ctx, guildID, viewer, attendanceAsk{officerRows: true, update: ask.update})
	if err != nil {
		return Reply{}, err
	}
	rep.Message.Content = fill(guideCopy.AttendanceUnknown, "{Guild}", guildName,
		"{typed}", memberText(truncateRunes(ask.name, charNameMax)))
	return rep, nil
}

// attendancePress answers the presses on an /attendance card. ask only opens the modal: it does
// no I/O and needs no guild. pick and name re-check rank through AttendanceFor every time.
func (s *Service) attendancePress() Handler {
	return s.cardHandler("/attendance", map[string]cardPart{
		"ask": func(context.Context, Interaction, cardPress) (Reply, error) {
			return Reply{Modal: attendanceModal()}, nil
		},
		"pick": s.attendanceLookup(picked, false),
		"name": s.attendanceLookup(func(in Interaction) string { return in.ModalValue(modalInputID) }, true),
	})
}

// attendanceLookup answers in place for the name a select or the modal gave.
func (s *Service) attendanceLookup(nameOf func(Interaction) string, typed bool) cardPart {
	return func(ctx context.Context, in Interaction, _ cardPress) (Reply, error) {
		guildID, viewer, bound, err := s.lookupCaller(ctx, in)
		if err != nil {
			return Reply{}, err
		}
		if !bound {
			return pressReply(commonCopy.Unbound, true), nil
		}
		if viewer == nil {
			return pressReply(officerCopy.AttendanceUnlinked, true, s.linkButton()), nil
		}
		name := strings.TrimSpace(nameOf(in))
		return s.attendanceAnswer(ctx, guildID, *viewer, attendanceAsk{
			name: name, officerRows: true, update: true, typedUnknown: typed && name != "",
		})
	}
}
