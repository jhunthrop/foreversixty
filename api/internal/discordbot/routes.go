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
	// Modal, when set, opens a dialog instead of sending a message (response type 9). Its
	// handler must do no I/O, so it answers inside the window; a modal is never deferred.
	Modal *Modal
	// After runs once the answer has gone to Discord, off the request: a follow-up message
	// needs the interaction acknowledged first. Nil runs nothing. When the answer could not be
	// delivered, After still runs, under a context marked dropped (answerDropped): it must then
	// only release what the handler held (a claimed token), never send.
	After func(ctx context.Context)
}

// Handler answers a command or a component press. An error it returns is
// logged, so it must carry no member text or credential.
type Handler func(ctx context.Context, in Interaction) (Reply, error)

// Autocomplete answers an option's live suggestions.
type Autocomplete func(ctx context.Context, in Interaction) ([]Choice, error)

// commandRoutes maps a slash command's name to its handler. Each later task
// adds one entry.
func (s *Service) commandRoutes() map[string]Handler {
	return map[string]Handler{
		"link": s.cmdLink, "signup": s.cmdSignup, "raid": s.cmdRaid, "help": s.cmdHelp,
		"char": s.cmdChar, "bis": s.cmdBis,
		"report": s.cmdReport, "attendance": s.cmdAttendance, "readiness": s.cmdReadiness,
	}
}

// componentRoutes maps a custom id's prefix (the part before the first colon)
// to its handler.
func (s *Service) componentRoutes() map[string]Handler {
	return map[string]Handler{raidIDPrefix: s.press, pickIDPrefix: s.pick, shareIDPrefix: s.share, prefixBis: s.bisPress(), prefixChar: s.charPress(), prefixSignup: s.signupPress(), prefixAttendance: s.attendancePress(),
		prefixReadiness: s.readinessPress(), prefixReport: s.reportPress(), prefixSpec: s.specPress()}
}

// autocompleteRoutes maps a command's name to its suggester.
func (s *Service) autocompleteRoutes() map[string]Autocomplete {
	return map[string]Autocomplete{
		"signup": s.autocompleteSignup, "char": s.autocompleteChar, "bis": s.autocompleteSpec,
		"report": s.autocompleteReport, "attendance": s.autocompleteMember, "readiness": s.autocompleteReadiness,
	}
}

// updateComponents names the custom-id prefixes whose press edits the message
// it sits on, so a slow one is deferred as type 6, not type 5.
func updateComponents() map[string]bool {
	return map[string]bool{pickIDPrefix: true, shareIDPrefix: true, prefixBis: true, prefixChar: true, prefixSignup: true, prefixAttendance: true,
		prefixReadiness: true, prefixReport: true, prefixSpec: true}
}

// customPrefix is the route key of a custom id.
func customPrefix(customID string) string {
	prefix, _, _ := strings.Cut(customID, ":")
	return prefix
}
