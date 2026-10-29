package discordbot

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

// optGuided marks a /char answer reached through the card. answerChar ignores it; Share
// stores it so the presser's own copy keeps its controls.
const (
	optGuided = "guided"
	guidedOn  = "1"
)

// Bounds of a character name the modal accepts.
const (
	charNameMin = 2
	charNameMax = 12
)

// ownChar is one of the viewer's verified characters with the level the card shows.
type ownChar struct {
	guilds.SignupCandidate
	Level *int
}

// charOptions are the resolved options of a /char answer.
func charOptions(name string, guided bool) map[string]string {
	opts := map[string]string{optName: name}
	if guided {
		opts[optGuided] = guidedOn
	}
	return opts
}

// charModal is the name field behind Type a name and Someone else.
func charModal() *Modal {
	return &Modal{
		CustomID: cardCustomID(prefixChar, "name"), Title: guideCopy.ModalCharTitle,
		InputLabel: guideCopy.ModalLabel, Placeholder: guideCopy.ModalPlaceholder,
		MinLength: charNameMin, MaxLength: charNameMax,
	}
}

// ownChars are the viewer's verified characters in the guild, the main first, at most the
// select's 25, each with its level. Empty for a visitor.
func (s *Service) ownChars(ctx context.Context, guildID int64, viewer *int64) ([]ownChar, error) {
	if viewer == nil {
		return nil, nil
	}
	cands, err := s.d.Guilds.SignupCandidates(ctx, guildID, *viewer)
	if err != nil {
		return nil, fmt.Errorf("discordbot: char own: %w", err)
	}
	verified := verifiedOf(cands)
	out := make([]ownChar, 0, min(len(verified), maxSelectOptions))
	for _, c := range verified[:min(len(verified), maxSelectOptions)] {
		level, err := s.levelOf(ctx, c.CharacterKey)
		if err != nil {
			return nil, err
		}
		out = append(out, ownChar{c, level})
	}
	return out, nil
}

// levelOf is a character's level from its public card; nil when unknown.
func (s *Service) levelOf(ctx context.Context, key string) (*int, error) {
	parts := strings.SplitN(key, "/", 3)
	if len(parts) != 3 {
		return nil, nil
	}
	card, found, err := s.d.Characters.CharacterCard(ctx, parts[0], parts[1], parts[2])
	if err != nil || !found {
		return nil, err
	}
	return card.Character.Level, nil
}

// ownSelect is the own-characters select; shown is the name on the card (its option is Default).
func ownSelect(own []ownChar, shown string) Component {
	options := make([]SelectOption, 0, len(own))
	for _, c := range own {
		options = append(options, SelectOption{
			Label: truncateRunes(c.Name, maxOptionRunes), Value: truncateRunes(c.Name, maxOptionRunes),
			Description: truncateRunes(ownDescription(c), maxOptionRunes), Default: strings.EqualFold(c.Name, shown),
		})
	}
	return Component{Kind: KindStringSelect, CustomID: cardCustomID(prefixChar, "pick"),
		Placeholder: guideCopy.CharSelect, Options: options}
}

// ownDescription is "{Spec} {Class} · Level {n}", without the level part when it is unknown.
func ownDescription(c ownChar) string {
	template := guideCopy.CharSelectDesc
	level := ""
	if c.Level != nil {
		level = strconv.Itoa(*c.Level)
	} else {
		template, _, _ = strings.Cut(template, lineJoin)
	}
	return strings.TrimSpace(fill(template, "{Spec}", c.Spec, "{Class}", classWord(c.Class), "{n}", level))
}

func askButton(label string, style ButtonStyle) Component {
	return Component{Kind: KindButton, Style: style, Label: label, CustomID: cardCustomID(prefixChar, "ask")}
}

// charAskCard asks for a character: no embed, Type a name (with Link Discord when not
// linked) and the own-characters select when the viewer has characters.
func (s *Service) charAskCard(content string, linked bool, own []ownChar, update bool) Reply {
	row := ActionRow{askButton(guideCopy.CharTypeName, StyleSuccess)}
	if !linked {
		row = append(row, s.linkButton())
	}
	rows := []ActionRow{row}
	if len(own) > 0 {
		rows = append(rows, ActionRow{ownSelect(own, "")})
	}
	return Reply{Message: Message{Content: content, Ephemeral: true, Components: rows}, Update: update}
}

// charUnknownCard says no public character has the typed name and asks again.
func (s *Service) charUnknownCard(ctx context.Context, guildID int64, viewer *int64, typed string, update bool) (Reply, error) {
	own, err := s.ownChars(ctx, guildID, viewer)
	if err != nil {
		return Reply{}, err
	}
	content := fill(guideCopy.CharUnknown, "{typed}", memberText(truncateRunes(typed, charNameMax)))
	return s.charAskCard(content, viewer != nil, own, update), nil
}

// charStart is a bare /char: the main's card, or the ask card.
func (s *Service) charStart(ctx context.Context, in Interaction, guildID int64, viewer *int64) (Reply, error) {
	own, err := s.ownChars(ctx, guildID, viewer)
	if err != nil {
		return Reply{}, err
	}
	switch {
	case viewer == nil:
		return s.charAskCard(guideCopy.CharLeadUnlinked, false, nil, false), nil
	case len(own) == 0:
		return s.charAskCard(guideCopy.CharLeadNone, true, nil, false), nil
	}
	return s.guidedChar(ctx, in, guildID, viewer, own[0].Name, false)
}

// guidedChar is the /char answer for a name with Share, Someone else and the own-characters
// select; a name that does not resolve draws the ask card instead.
func (s *Service) guidedChar(ctx context.Context, in Interaction, guildID int64, viewer *int64, name string, update bool) (Reply, error) {
	opts := charOptions(name, true)
	msg, found, err := s.answerChar(ctx, guildID, viewer, opts)
	if err != nil {
		return Reply{}, err
	}
	if !found {
		return s.charUnknownCard(ctx, guildID, viewer, name, update)
	}
	button, err := s.offerShare(ctx, in, guildID, "char", opts)
	if err != nil {
		return Reply{}, err
	}
	msg, err = s.withGuidedControls(ctx, "char", opts, guildID, viewer, msg)
	if err != nil {
		return Reply{}, err
	}
	msg = withShare(msg, button)
	msg.Ephemeral = true
	return Reply{Message: msg, Update: update}, nil
}

// withGuidedControls adds Someone else and the select to the presser's own copy of a guided
// /char answer; any other answer is returned as it is. The public follow-up never gets them.
func (s *Service) withGuidedControls(ctx context.Context, command string, opts map[string]string, guildID int64, viewer *int64, msg Message) (Message, error) {
	if command != "char" || opts[optGuided] != guidedOn || len(msg.Components) == 0 || len(msg.Components[0]) == 0 {
		return msg, nil
	}
	own, err := s.ownChars(ctx, guildID, viewer)
	if err != nil {
		return Message{}, err
	}
	rows := append([]ActionRow(nil), msg.Components...)
	row := rows[0]
	last := len(row) - 1
	rows[0] = append(append(append(ActionRow(nil), row[:last]...), askButton(guideCopy.CharSomeoneElse, StyleSecondary)), row[last:]...)
	if len(own) >= 2 {
		rows = append(rows, ActionRow{ownSelect(own, opts[optName])})
	}
	msg.Components = rows
	return msg, nil
}

// charPress answers the presses on a /char card. ask only opens the modal: it does no I/O
// and needs no guild, because it reads nothing.
func (s *Service) charPress() Handler {
	return s.cardHandler("/char", map[string]cardPart{
		"ask":  func(context.Context, Interaction, cardPress) (Reply, error) { return Reply{Modal: charModal()}, nil },
		"pick": s.charLookup(picked),
		"name": s.charLookup(func(in Interaction) string { return in.ModalValue(modalInputID) }),
	})
}

// charLookup answers in place for the name a select or the modal gave.
func (s *Service) charLookup(nameOf func(Interaction) string) cardPart {
	return func(ctx context.Context, in Interaction, _ cardPress) (Reply, error) {
		guildID, viewer, bound, err := s.lookupCaller(ctx, in)
		if err != nil {
			return Reply{}, err
		}
		if !bound {
			return pressReply(commonCopy.Unbound, true), nil
		}
		return s.guidedChar(ctx, in, guildID, viewer, nameOf(in), true)
	}
}
