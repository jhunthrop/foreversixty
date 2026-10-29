package discordbot

import (
	"context"
	"strings"
)

// cardVersion is the grammar version in every guided-card custom id.
const cardVersion = "1"

// Card custom-id prefixes, one per command; the router keys on them (customPrefix).
const (
	prefixBis        = "gbis"
	prefixChar       = "gchar"
	prefixSignup     = "gsgn"
	prefixAttendance = "gatt"
	prefixReadiness  = "grdy"
	prefixReport     = "grep"
)

// cardFooter is the footer of every card embed.
const cardFooter = "Forever Sixty"

// cardMinFields is prefix, version and part: the fewest fields a card id has.
const cardMinFields = 3

// lockedOptionValue is the one value of a disabled select, which Discord needs to have an option.
const lockedOptionValue = "none"

// cardCustomID is `{prefix}:1:{part}:{state joined by ":"}`, keeping empty fields. State
// never names another account and never carries a credential.
func cardCustomID(prefix, part string, state ...string) string {
	return strings.Join([]string{prefix, cardVersion, part, strings.Join(state, ":")}, ":")
}

// cardPress is a parsed card custom id.
type cardPress struct {
	Prefix, Version, Part string
	State                 []string
}

// parseCardID splits a card custom id; ok is false when it has fewer than three fields.
// The state is whatever follows the part, split on ":" with empty fields kept.
func parseCardID(id string) (cardPress, bool) {
	fields := strings.SplitN(id, ":", cardMinFields+1)
	if len(fields) < cardMinFields {
		return cardPress{}, false
	}
	p := cardPress{Prefix: fields[0], Version: fields[1], Part: fields[2]}
	if len(fields) > cardMinFields {
		p.State = strings.Split(fields[cardMinFields], ":")
	}
	return p, true
}

// staleCard replaces a card of an older shape in place with the way out. command must be a
// constant: it is inserted unescaped.
func staleCard(command string) Reply {
	content := strings.Replace(guideCopy.CardStale, "{command}", command, 1)
	return Reply{Update: true, Message: Message{Content: content, Ephemeral: true}}
}

// lockedSelect is a disabled select showing only its placeholder.
func lockedSelect(customID, placeholder string) Component {
	return Component{
		Kind: KindStringSelect, CustomID: customID, Placeholder: placeholder, Disabled: true,
		Options: []SelectOption{{Label: placeholder, Value: lockedOptionValue}},
	}
}

// cardPart answers one component of a card.
type cardPart func(ctx context.Context, in Interaction, p cardPress) (Reply, error)

// cardHandler routes a press on a command's card to its part. An id of another version
// answers the stale card; an unknown part or unreadable id answers the failure sentence.
func (s *Service) cardHandler(command string, parts map[string]cardPart) Handler {
	return func(ctx context.Context, in Interaction) (Reply, error) {
		p, ok := parseCardID(in.CustomID())
		if ok && p.Version != cardVersion {
			return staleCard(command), nil
		}
		part := parts[p.Part]
		if !ok || part == nil {
			rep := failureReply()
			rep.Update = true
			return rep, nil
		}
		return part(ctx, in, p)
	}
}
