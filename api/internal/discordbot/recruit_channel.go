package discordbot

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/copygen"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
) // The recruiting channel (guild onboarding spec 2026-11-09 §2.3): an officer picks one of the
// server's channels, and the bot posts each new application there. Free for any claimed guild.

var _ guilds.ApplicationPoster = (*Service)(nil)

// recruitCopy is the recruiting post's words.
var recruitCopy = struct {
	Title, Level, ItemLevel, Availability, Footer, BestParse, Parse string
}{
	Title:        copygen.DiscordRecruitTitle,
	Level:        copygen.PlanLevelSheet,
	ItemLevel:    copygen.RaidApplicantItemLevel,
	Availability: copygen.DiscordRecruitAvailability,
	Footer:       copygen.DiscordCmdReportGuildFooter,
	BestParse:    copygen.DiscordRecruitBestParse,
	Parse:        copygen.DiscordRecruitParse,
}

// maxApplicationRunes caps the message the post quotes; the applications page has the rest.
const maxApplicationRunes = 600

// recruitChannelSave is PUT …/discord/bot/recruit-channel {"channel_id"}: the server's channel
// the bot can post in becomes the recruiting channel; an empty id clears it.
func (s *Service) recruitChannelSave(w http.ResponseWriter, r *http.Request) {
	httpx.CachePrivateNoStore(w)
	guildID, _, ok := s.officer(w, r)
	if !ok {
		return
	}
	channelID, ok := decodeRecruitChannel(w, r)
	if !ok {
		return
	}
	in, ok := s.installOf(w, r, guildID)
	if !ok {
		return
	}
	if !s.saveRecruitChannel(w, r, in, channelID) {
		return
	}
	view, err := s.d.Guilds.SettingsDiscordFor(r.Context(), guildID, auth.ActorFrom(r.Context()).UserID)
	if err != nil || view == nil {
		s.internalError(w, r, "recruit channel settings", errors.Join(err, errors.New("no settings view")))
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, view)
}

// saveRecruitChannel stores the pick (or clears it); false once it has answered.
func (s *Service) saveRecruitChannel(w http.ResponseWriter, r *http.Request, in Install, channelID string) bool {
	if channelID == "" {
		if err := s.d.Store.SetRecruitChannel(r.Context(), in.GuildID, nil, nil); err != nil {
			s.internalError(w, r, "recruit channel clear", err)
			return false
		}
		return true
	}
	channels, ok := s.serverChannels(w, r, in)
	if !ok {
		return false
	}
	picked, usable := findChannel(channels, channelID)
	if !usable {
		httpx.WriteError(w, r, http.StatusUnprocessableEntity, "channel_unusable", channelUnusableText, nil)
		return false
	}
	if err := s.d.Store.SetRecruitChannel(r.Context(), in.GuildID, &picked.ID, &picked.Name); err != nil {
		s.internalError(w, r, "recruit channel save", err)
		return false
	}
	return true
}

// decodeRecruitChannel reads {"channel_id": string}; the id may be empty (clear).
func decodeRecruitChannel(w http.ResponseWriter, r *http.Request) (string, bool) {
	dec := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxChannelBody))
	dec.DisallowUnknownFields()
	var in channelInput
	if err := dec.Decode(&in); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be {\"channel_id\": string}", nil)
		return "", false
	}
	return in.ChannelID, true
}

// SetRecruitChannel saves (or, with nil, clears) the guild's recruiting channel.
func (s *Store) SetRecruitChannel(ctx context.Context, guildID int64, channelID, channelName *string) error {
	return s.exec(ctx, `update guild_discord_installs set recruit_channel_id = $2, recruit_channel_name = $3,
		updated_at = now() where guild_id = $1`, guildID, channelID, channelName)
}

// PostApplication implements guilds.ApplicationPoster: the application's embed in the guild's
// recruiting channel. No install, the bot removed, or no recruiting channel is OutcomeNoBot.
func (s *Service) PostApplication(ctx context.Context, guildID int64, post guilds.ApplicationPost) (string, error) {
	in, found, err := s.d.Store.InstallFor(ctx, guildID)
	if err != nil {
		return "", err
	}
	if !found || in.RemovedInDiscord || in.RecruitChannelID == nil {
		return guilds.OutcomeNoBot, nil
	}
	msg := Message{Embeds: []notify.Embed{applicationEmbed(post)},
		Components: []ActionRow{applicationButtons(post.ID, post.ReviewURL, s.characterURL(post.CharacterKey))}}
	if _, err := s.d.REST.CreateMessage(ctx, *in.RecruitChannelID, msg); err != nil {
		return refusedOr(err)
	}
	return guilds.OutcomeSent, nil
}

// applicationEmbed is the recruiting post: the character and spec, the message, then the facts.
func applicationEmbed(p guilds.ApplicationPost) notify.Embed {
	e := notify.Embed{
		Title: fill(recruitCopy.Title, "{Name}", notify.Collapse(p.Character)), URL: p.ReviewURL,
		Author: p.SpecClass, Description: clipRunes(p.Message, maxApplicationRunes), Color: accentColor,
		Footer: fill(recruitCopy.Footer, "{Guild}", notify.Collapse(p.GuildName)),
	}
	if p.Level != nil {
		e.Fields = append(e.Fields, notify.EmbedField{Name: recruitCopy.Level, Value: strconv.Itoa(*p.Level), Inline: true})
	}
	if p.ItemLevel != nil {
		e.Fields = append(e.Fields, notify.EmbedField{Name: recruitCopy.ItemLevel, Value: strconv.Itoa(*p.ItemLevel), Inline: true})
	}
	if b := p.BestParse; b != nil && b.Percentile != nil && b.Encounter != "" {
		pct := strconv.Itoa(int(*b.Percentile))
		e.Fields = append(e.Fields, notify.EmbedField{Name: recruitCopy.BestParse, Inline: true,
			Value: fill(recruitCopy.Parse, "{pct}", pct, "{Boss}", notify.Collapse(b.Encounter))})
	}
	if p.Availability != "" {
		e.Fields = append(e.Fields, notify.EmbedField{Name: recruitCopy.Availability, Value: notify.Collapse(p.Availability)})
	}
	return e
}

// clipRunes cuts text to at most n runes, ending on an ellipsis when cut.
func clipRunes(text string, n int) string {
	runes := []rune(text)
	if len(runes) <= n {
		return text
	}
	return fmt.Sprintf("%s…", string(runes[:n-1]))
}
