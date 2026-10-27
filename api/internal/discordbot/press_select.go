package discordbot

import (
	"context"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

// Discord's string select bounds (experience §3.3.0).
const (
	maxSelectOptions = 25
	maxOptionRunes   = 100
)

// characterSelect asks which verified character to sign up: the main first, then the
// service's order, at most 25, the text saying how many were left out. With no option Discord
// would take, it answers the failure sentence with the raid page, where they can be signed up.
func (s *Service) characterSelect(ctx context.Context, a signupAsk, raid guilds.Raid, verified []guilds.SignupCandidate) Reply {
	options := make([]SelectOption, 0, maxSelectOptions)
	selectable := 0
	for _, c := range verified {
		if utf8.RuneCountInString(c.CharacterKey) > maxOptionRunes {
			continue // a value Discord would refuse; the raid page still has it
		}
		selectable++
		if len(options) < maxSelectOptions {
			options = append(options, characterOption(c))
		}
	}
	if len(options) == 0 {
		return a.reply(pressCopy.Failure, linkButton(postCopy.RaidPage, s.raidPageURL(ctx, raid.ID)))
	}
	text := fill(pressCopy.SelectPrompt, "{State}", "**"+stateWord(a.State)+"**",
		"{Title}", memberText(raid.Title), "{unix}", strconv.FormatInt(raid.StartsAt.Unix(), 10))
	if selectable > maxSelectOptions {
		text += " " + fill(pressCopy.Showing, "{n}", strconv.Itoa(selectable))
	}
	sel := Component{
		Kind: KindStringSelect, CustomID: pickCustomID(raid.ID, a.State),
		Placeholder: pressCopy.SelectPlaceholder, Options: options,
	}
	return Reply{Message: Message{Content: text, Ephemeral: true, Components: []ActionRow{{sel}}}, Update: a.Update}
}

// characterOption is one select entry: the name, and "{Spec} {Class} · {Role}" with
// " · your main" on the main, each inside Discord's 100.
func characterOption(c guilds.SignupCandidate) SelectOption {
	return SelectOption{
		Label: truncateRunes(c.Name, maxOptionRunes), Value: c.CharacterKey, Description: optionDescription(c),
	}
}

// optionDescription cuts the spec, never the role or the main mark, to stay inside 100.
func optionDescription(c guilds.SignupCandidate) string {
	build := func(spec string) string {
		text := strings.TrimSpace(fill(pressCopy.OptionDescription,
			"{Spec}", spec, "{Class}", classWord(c.Class), "{Role}", roleWord(c.Role)))
		if c.Main {
			text += pressCopy.OptionMain
		}
		return text
	}
	text := build(c.Spec)
	if over := utf8.RuneCountInString(text) - maxOptionRunes; over > 0 {
		spec := []rune(c.Spec)
		text = build(string(spec[:max(0, len(spec)-over)]))
	}
	return truncateRunes(text, maxOptionRunes)
}

// classWord is a class as the game names it: "hunter" reads "Hunter".
func classWord(class string) string {
	r, size := utf8.DecodeRuneInString(class)
	if size == 0 {
		return ""
	}
	return string(unicode.ToUpper(r)) + class[size:]
}

func truncateRunes(s string, n int) string {
	runes := []rune(s)
	if len(runes) <= n {
		return s
	}
	return string(runes[:n])
}

// pickCustomID is the select's custom id: pick:{raid}:{state}.
func pickCustomID(raidID int64, state string) string {
	return pickIDPrefix + strings.TrimPrefix(raidCustomID(raidID, state), raidIDPrefix)
}

// parsePickCustomID reads a select's custom id; ok is false for any other shape.
func parsePickCustomID(id string) (raidID int64, state string, ok bool) {
	rest, found := strings.CutPrefix(id, pickIDPrefix+idSeparator)
	if !found {
		return 0, "", false
	}
	return parseRaidCustomID(raidIDPrefix + idSeparator + rest)
}

// pick answers a choice in the character select: the flow again, with the character chosen
// (re-checked against the account's verified characters), replacing the select's message.
func (s *Service) pick(ctx context.Context, in Interaction) (Reply, error) {
	raidID, state, ok := parsePickCustomID(in.CustomID())
	values := in.Values()
	if !ok || len(values) != 1 || values[0] == "" {
		return pressReply(pressCopy.Failure, true), nil
	}
	a, refused, ok := s.resolveAsker(ctx, in, true)
	if !ok {
		return refused, nil
	}
	a.RaidID, a.State, a.CharacterKey = raidID, state, values[0]
	return s.runFlow(ctx, in, a), nil
}
