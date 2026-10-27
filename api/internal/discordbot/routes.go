package discordbot

import (
	"context"
	"strings"
)

// Reply is a handler's answer. Update replaces the message a component sits
// on (response type 7, deferred type 6) instead of sending a new one.
type Reply struct {
	Message Message
	Update  bool
}

// Handler answers a command or a component press. An error it returns is
// logged, so it must carry no member text or credential.
type Handler func(ctx context.Context, in Interaction) (Reply, error)

// Autocomplete answers an option's live suggestions.
type Autocomplete func(ctx context.Context, in Interaction) ([]Choice, error)

// commandRoutes maps a slash command's name to its handler. Each later task
// adds one entry.
func (s *Service) commandRoutes() map[string]Handler {
	return map[string]Handler{"link": s.cmdLink, "signup": s.cmdSignup, "raid": s.cmdRaid, "help": s.cmdHelp}
}

// componentRoutes maps a custom id's prefix (the part before the first colon)
// to its handler.
func (s *Service) componentRoutes() map[string]Handler {
	return map[string]Handler{raidIDPrefix: s.press, pickIDPrefix: s.pick}
}

// autocompleteRoutes maps a command's name to its suggester.
func (s *Service) autocompleteRoutes() map[string]Autocomplete {
	return map[string]Autocomplete{"signup": s.autocompleteSignup}
}

// updateComponents names the custom-id prefixes whose press edits the message
// it sits on, so a slow one is deferred as type 6, not type 5.
func updateComponents() map[string]bool { return map[string]bool{pickIDPrefix: true} }

// customPrefix is the route key of a custom id.
func customPrefix(customID string) string {
	prefix, _, _ := strings.Cut(customID, ":")
	return prefix
}
