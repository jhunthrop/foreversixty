package discordbot

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

const (
	maxChannelBody = 1 << 10
	// maxTestErrorRunes caps the Discord text a failed test shows.
	maxTestErrorRunes = 140

	removedInDiscordNote = "the bot was removed from the server"
	unavailableMessage   = "Discord is not answering for that server just now"
	noAnswerMessage      = "Discord did not answer"
	channelUnusableText  = "the bot cannot post in that channel"
)

var urlText = regexp.MustCompile(`https?://\S*`)

// serverChannels reads the server's postable channels. A missing-access answer marks the
// install removed in Discord. ok is false once it has answered.
func (s *Service) serverChannels(w http.ResponseWriter, r *http.Request, in Install) ([]Channel, bool) {
	layout, err := s.d.REST.ServerLayout(r.Context(), in.ServerID)
	if err != nil {
		var api *APIError
		if errors.As(err, &api) && api.MissingAccess() {
			if markErr := s.d.Store.MarkRemovedInDiscord(r.Context(), in.GuildID, removedInDiscordNote); markErr != nil {
				s.logFailure(r, "mark removed", markErr)
			}
		}
		s.log.Warn("discordbot", "id", httpx.RequestIDFrom(r.Context()), "op", "server layout", "err", err)
		httpx.WriteError(w, r, http.StatusBadGateway, "discord_unavailable", unavailableMessage, nil)
		return nil, false
	}
	return PostableChannels(in.ServerID, layout), true
}

// installOf is the guild's install, answering 404 when there is none.
func (s *Service) installOf(w http.ResponseWriter, r *http.Request, guildID int64) (Install, bool) {
	in, found, err := s.d.Store.InstallFor(r.Context(), guildID)
	if err != nil {
		s.internalError(w, r, "install read", err)
		return Install{}, false
	}
	if !found {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", noInstallMessage, nil)
		return Install{}, false
	}
	return in, true
}

func (s *Service) channelsList(w http.ResponseWriter, r *http.Request) {
	httpx.CachePrivateNoStore(w)
	guildID, _, ok := s.officer(w, r, true)
	if !ok {
		return
	}
	in, ok := s.installOf(w, r, guildID)
	if !ok {
		return
	}
	channels, ok := s.serverChannels(w, r, in)
	if !ok {
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, map[string][]Channel{"channels": channels})
}

type channelInput struct {
	ChannelID string `json:"channel_id"`
}

func decodeChannel(w http.ResponseWriter, r *http.Request) (string, bool) {
	dec := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxChannelBody))
	dec.DisallowUnknownFields()
	var in channelInput
	if err := dec.Decode(&in); err != nil || in.ChannelID == "" {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be {\"channel_id\": string}", nil)
		return "", false
	}
	return in.ChannelID, true
}

func (s *Service) channelSave(w http.ResponseWriter, r *http.Request) {
	httpx.CachePrivateNoStore(w)
	guildID, _, ok := s.officer(w, r, true)
	if !ok {
		return
	}
	channelID, ok := decodeChannel(w, r)
	if !ok {
		return
	}
	in, ok := s.installOf(w, r, guildID)
	if !ok {
		return
	}
	channels, ok := s.serverChannels(w, r, in)
	if !ok {
		return
	}
	picked, usable := findChannel(channels, channelID)
	if !usable {
		httpx.WriteError(w, r, http.StatusUnprocessableEntity, "channel_unusable", channelUnusableText, nil)
		return
	}
	if err := s.d.Store.SetChannel(r.Context(), guildID, picked.ID, picked.Name); err != nil {
		s.internalError(w, r, "channel save", err)
		return
	}
	view, err := s.d.Guilds.SettingsDiscordFor(r.Context(), guildID, auth.ActorFrom(r.Context()).UserID)
	if err != nil || view == nil {
		s.internalError(w, r, "channel settings", errors.Join(err, errors.New("no settings view")))
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, view)
}

func findChannel(channels []Channel, id string) (Channel, bool) {
	for _, c := range channels {
		if c.ID == id {
			return c, true
		}
	}
	return Channel{}, false
}

func (s *Service) botTest(w http.ResponseWriter, r *http.Request) {
	httpx.CachePrivateNoStore(w)
	guildID, standing, ok := s.officer(w, r, true)
	if !ok {
		return
	}
	in, ok := s.installOf(w, r, guildID)
	if !ok {
		return
	}
	if in.ChannelID == nil {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", noChannelMessage, nil)
		return
	}
	embed := notify.GuildTestEmbed(s.d.PublicBaseURL, standing.Name, standing.Path, s.now())
	if _, err := s.d.REST.CreateMessage(r.Context(), *in.ChannelID, Message{Embeds: []notify.Embed{embed}}); err != nil {
		// A manual test never counts toward the failures, but it does learn the bot is gone.
		s.markIfBotGone(r, in, err)
		s.testFailed(w, r, err)
		return
	}
	if err := s.d.Store.RecordPostSuccess(r.Context(), guildID); err != nil {
		s.internalError(w, r, "test record", err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// markIfBotGone records removed_in_discord when Discord says the bot cannot see the server
// (missing access or unknown server); a missing permission in one channel is not that.
func (s *Service) markIfBotGone(r *http.Request, in Install, err error) {
	var api *APIError
	if !errors.As(err, &api) || (api.Code != codeMissingAccess && api.Code != codeUnknownGuild) {
		return
	}
	if markErr := s.d.Store.MarkRemovedInDiscord(r.Context(), in.GuildID, removedInDiscordNote); markErr != nil {
		s.logFailure(r, "mark removed", markErr)
	}
}

const noChannelMessage = "no channel is chosen"

// testFailed answers a refused test with Discord's own sanitised words.
func (s *Service) testFailed(w http.ResponseWriter, r *http.Request, err error) {
	message, status := noAnswerMessage, http.StatusGatewayTimeout
	var api *APIError
	if errors.As(err, &api) {
		s.log.Warn("discordbot", "id", httpx.RequestIDFrom(r.Context()), "op", "test", "status", api.Status)
		status = http.StatusBadGateway
		if text := sanitizeDiscordText(api.Message); text != "" {
			message = text
		}
	} else {
		s.log.Warn("discordbot", "id", httpx.RequestIDFrom(r.Context()), "op", "test", "err", "no answer")
	}
	httpx.WriteError(w, r, status, "upstream", message, nil)
}

// sanitizeDiscordText drops URLs, collapses space and caps the length.
func sanitizeDiscordText(text string) string {
	text = strings.Join(strings.Fields(urlText.ReplaceAllString(text, "")), " ")
	if runes := []rune(text); len(runes) > maxTestErrorRunes {
		return string(runes[:maxTestErrorRunes])
	}
	return text
}

func (s *Service) botRemove(w http.ResponseWriter, r *http.Request) {
	httpx.CachePrivateNoStore(w)
	guildID, _, ok := s.officer(w, r, false)
	if !ok {
		return
	}
	in, found, err := s.d.Store.InstallFor(r.Context(), guildID)
	if err != nil {
		s.internalError(w, r, "remove read", err)
		return
	}
	if found {
		s.leaveServer(r.Context(), in.ServerID)
	}
	if err := s.d.Store.RemoveBot(r.Context(), guildID); err != nil {
		s.internalError(w, r, "remove", err)
		return
	}
	s.log.Info("discordbot", "op", "remove", "guild_id", guildID)
	w.WriteHeader(http.StatusNoContent)
}

// leaveServer asks the bot to leave; a failure is logged and never fails the remove.
func (s *Service) leaveServer(ctx context.Context, serverID string) {
	leaveCtx, cancel := context.WithTimeout(ctx, requestTimeout+time.Second)
	defer cancel()
	if err := s.d.REST.LeaveServer(leaveCtx, serverID); err != nil {
		s.log.Warn("discordbot", "op", "leave server", "err", err)
	}
}
