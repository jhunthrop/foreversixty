package discordbot

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Interaction callback response types.
const (
	respPong               = 1
	respMessage            = 4
	respDeferredMessage    = 5
	respDeferredUpdate     = 6
	respUpdateMessage      = 7
	respAutocompleteResult = 8
)

type callbackBody struct {
	Type int `json:"type"`
	Data any `json:"data,omitempty"`
}

func writeCallback(w http.ResponseWriter, body callbackBody) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(body)
}

func writeUnauthorized(w http.ResponseWriter) {
	http.Error(w, "invalid request signature", http.StatusUnauthorized)
}

// serveInteraction verifies the signature on the raw body, then dispatches.
func (s *Service) serveInteraction(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(io.LimitReader(r.Body, MaxInteractionBytes+1))
	if err != nil || len(body) > MaxInteractionBytes {
		writeUnauthorized(w)
		return
	}
	received := s.now()
	err = Verify(s.d.PublicKey, r.Header.Get("X-Signature-Ed25519"), r.Header.Get("X-Signature-Timestamp"), body, received)
	if err != nil {
		writeUnauthorized(w)
		return
	}
	in, err := ParseInteraction(body, received)
	if err != nil {
		http.Error(w, "unreadable interaction", http.StatusBadRequest)
		return
	}
	switch in.Type {
	case typePing:
		writeCallback(w, callbackBody{Type: respPong})
	case typeCommand, typeComponent:
		s.dispatch(r.Context(), w, in)
	case typeAutocomplete:
		s.autocomplete(r.Context(), w, in)
	default:
		s.writeReply(w, failureReply(), respMessage)
	}
}

func failureReply() Reply {
	return Reply{Message: Message{Content: commonCopy.Failure, Ephemeral: true}}
}

func unboundReply() Reply {
	return Reply{Message: Message{Content: commonCopy.Unbound, Ephemeral: true}}
}

func (s *Service) writeReply(w http.ResponseWriter, rep Reply, kind int) {
	if rep.Update && kind == respMessage {
		// An update replaces the message: what it leaves out (the select) must go, so it
		// sends the edit body, which writes every part explicitly.
		writeCallback(w, callbackBody{Type: respUpdateMessage, Data: rep.Message.edit()})
		return
	}
	writeCallback(w, callbackBody{Type: kind, Data: rep.Message})
}

// route finds the handler and whether its press updates the message it sits on.
func (s *Service) route(in Interaction) (h Handler, update, ok bool) {
	if in.Type == typeCommand {
		h, ok = s.commands[in.CommandName()]
		return h, false, ok
	}
	prefix := customPrefix(in.CustomID())
	h, ok = s.components[prefix]
	return h, s.updating[prefix], ok
}

// guildFreeCommands are the commands that need no installed guild: they answer the same in
// an unbound server or a DM.
func guildFreeCommands() map[string]bool { return map[string]bool{"link": true, "help": true} }

// dispatch resolves the route and the server, then answers or defers.
func (s *Service) dispatch(ctx context.Context, w http.ResponseWriter, in Interaction) {
	h, update, ok := s.route(in)
	if !ok {
		s.writeReply(w, failureReply(), respMessage)
		return
	}
	if in.Type == typeCommand && guildFreeCommands()[in.CommandName()] {
		s.answerOrDefer(ctx, w, in, h, update)
		return
	}
	_, bound, err := s.serverGuild(ctx, in)
	switch {
	case err != nil:
		s.log.Error("discordbot", "op", "discord_answer", "outcome", "server_lookup_failed", "err", err)
		s.writeReply(w, failureReply(), respMessage)
	case !bound:
		s.writeReply(w, unboundReply(), respMessage)
	default:
		s.answerOrDefer(ctx, w, in, h, update)
	}
}

type handled struct {
	reply Reply
	err   error
}

var errShuttingDown = errors.New("discordbot: shutting down")

// runHandler runs h on its own goroutine, tracked by Background, bounded by
// the token's work budget and detached from the request's cancellation (its
// values, the request id, stay). A panic is an error carrying only its type,
// never its value. Handler errors must be free of member data: they are logged.
func (s *Service) runHandler(reqCtx context.Context, in Interaction, h Handler) <-chan handled {
	done := make(chan handled, 1)
	ctx, cancel := context.WithDeadline(context.WithoutCancel(reqCtx), in.ReceivedAt.Add(workBudget))
	started := s.bg.Start(func() {
		defer cancel()
		defer func() {
			if p := recover(); p != nil {
				done <- handled{err: fmt.Errorf("handler panicked: %T", p)}
			}
		}()
		rep, err := h(ctx, in)
		done <- handled{rep, err}
	})
	if !started {
		cancel()
		done <- handled{err: errShuttingDown}
	}
	return done
}

// remaining is what is left of the answer window, counted from when Discord
// sent the interaction, so the server lookup before it is not extra.
func (s *Service) remaining(in Interaction) time.Duration {
	return max(0, s.answerWithin-s.now().Sub(in.ReceivedAt))
}

func (s *Service) answerOrDefer(ctx context.Context, w http.ResponseWriter, in Interaction, h Handler, update bool) {
	done := s.runHandler(ctx, in, h)
	timer := time.NewTimer(s.remaining(in))
	defer timer.Stop()
	select {
	case res := <-done:
		s.writeReply(w, s.replyOf(in, res), respMessage)
	case <-timer.C:
		s.writeDeferral(w, update)
		if !s.bg.Start(func() { s.deliver(in, <-done) }) {
			s.log.Warn("discordbot", "op", "discord_answer", "outcome", "dropped_shutdown", "interaction", in.ID)
		}
	}
}

func (s *Service) writeDeferral(w http.ResponseWriter, update bool) {
	if update {
		writeCallback(w, callbackBody{Type: respDeferredUpdate})
		return
	}
	writeCallback(w, callbackBody{Type: respDeferredMessage, Data: map[string]int{"flags": ephemeralFlag}})
}

// replyOf turns a handler's outcome into the reply to send; an error is the
// ephemeral failure sentence, logged without the interaction token.
func (s *Service) replyOf(in Interaction, res handled) Reply {
	if res.err != nil {
		s.log.Error("discordbot", "op", "discord_answer", "outcome", "handler_failed", "interaction", in.ID, "err", res.err)
		return failureReply()
	}
	return res.reply
}

// deliver edits the deferred answer in, unless the token has expired.
func (s *Service) deliver(in Interaction, res handled) {
	if s.now().Sub(in.ReceivedAt) > interactionTokenLife {
		s.log.Warn("discordbot", "op", "discord_answer", "outcome", "expired", "interaction", in.ID)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.d.REST.EditOriginal(ctx, in.Token, s.replyOf(in, res).Message); err != nil {
		s.log.Warn("discordbot", "op", "discord_answer", "outcome", "edit_failed", "interaction", in.ID, "err", err)
	}
}

// autocomplete never defers: it answers inside answerWithin or with no choices.
func (s *Service) autocomplete(ctx context.Context, w http.ResponseWriter, in Interaction) {
	choices := []Choice{}
	if h, ok := s.autocompletes[in.CommandName()]; ok {
		choices = s.suggest(ctx, in, h)
	}
	writeCallback(w, callbackBody{Type: respAutocompleteResult, Data: map[string][]Choice{"choices": choices}})
}

type suggestions struct {
	choices []Choice
	err     error
}

// suggest runs h through Background, bounded by the answer window. A timeout,
// an error, a panic or a shutdown is no choices.
func (s *Service) suggest(ctx context.Context, in Interaction, h Autocomplete) []Choice {
	ctx, cancel := context.WithTimeout(context.WithoutCancel(ctx), s.remaining(in))
	defer cancel()
	done := make(chan suggestions, 1)
	started := s.bg.Start(func() {
		defer func() {
			if p := recover(); p != nil {
				done <- suggestions{err: fmt.Errorf("autocomplete panicked: %T", p)}
			}
		}()
		c, err := h(ctx, in)
		done <- suggestions{c, err}
	})
	if !started {
		return []Choice{}
	}
	select {
	case res := <-done:
		if res.err != nil {
			s.log.Error("discordbot", "op", "discord_autocomplete", "interaction", in.ID, "err", res.err)
		} else if res.choices != nil {
			return res.choices
		}
	case <-ctx.Done():
	}
	return []Choice{}
}
