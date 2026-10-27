package discordbot

import (
	"context"
	"errors"
	"fmt"
)

// The permissions a raid channel needs, by name and in the order settings lists them. A post
// is a create (Send Messages, Embed Links for the embed) and edits of the bot's own message,
// which Discord gates on nothing more; View Channel underlies both. Read Message History is
// not needed: the bot never replies to or reads a message.
// postPermissions is channelNeeds as one set: the channel picker's gate.
var postPermissions = PermViewChannel | PermSendMessages | PermEmbedLinks

var channelNeeds = []struct {
	Perm Perm
	Name string
}{
	{PermViewChannel, "View Channel"},
	{PermSendMessages, "Send Messages"},
	{PermEmbedLinks, "Embed Links"},
}

// channelVerdict is what a channel-level refusal turned out to mean.
type channelVerdict int

const (
	// verdictRetry: the probe could not tell, so the call is retried after a backoff.
	verdictRetry channelVerdict = iota
	// verdictRemoved: the bot is not in the server (state 3, removed_in_discord).
	verdictRemoved
	// verdictUnusable: the bot is in the server but cannot post in the channel.
	verdictUnusable
)

// channelRefused reports an answer about the channel or the server rather than one message:
// unknown channel, or any missing access or permission.
func channelRefused(api *APIError) bool {
	return api.Code == codeUnknownChannel || api.MissingAccess()
}

// judgeChannel works out a channel-level refusal with one read of the server. A server the
// bot cannot read (403, 50001, 10004) is removed_in_discord. A server it can read decides the
// channel: absent is deleted, lacking a needed permission is permissions. Either way the
// install records it. Anything else (no answer, a 429, a channel that looks fine) is a retry.
func judgeChannel(ctx context.Context, store *Store, rest REST, guildID int64, channelID string, refusal error) (channelVerdict, error) {
	in, found, err := store.InstallFor(ctx, guildID)
	if err != nil || !found {
		return verdictRetry, err
	}
	layout, err := rest.ServerLayout(ctx, in.ServerID)
	var probe *APIError
	if errors.As(err, &probe) && probe.MissingAccess() {
		return verdictRemoved, store.MarkRemovedInDiscord(ctx, guildID, removedInDiscordNote)
	}
	if err != nil {
		return verdictRetry, nil
	}
	u, unusable := unusableIn(in.ServerID, layout, channelID)
	if !unusable {
		return verdictRetry, nil
	}
	if err := store.MarkChannelUnusable(ctx, guildID, u, describe(refusal)); err != nil {
		return verdictRetry, fmt.Errorf("discordbot: mark channel unusable: %w", err)
	}
	return verdictUnusable, nil
}

// unusableIn is the channel's unusable state in layout; false when it exists and the bot
// holds every permission the post needs.
func unusableIn(serverID string, layout ServerLayout, channelID string) (ChannelUnusable, bool) {
	for _, c := range layout.Channels {
		if c.ID != channelID {
			continue
		}
		perms := channelPermissions(basePermissions(serverID, layout), serverID, layout, c)
		missing := []string{}
		for _, need := range channelNeeds {
			if perms&need.Perm == 0 {
				missing = append(missing, need.Name)
			}
		}
		if len(missing) == 0 {
			return ChannelUnusable{}, false
		}
		return ChannelUnusable{Reason: UnusablePermissions, Missing: missing}, true
	}
	// Absent means deleted: GET /guilds/{id}/channels answers a bot every channel of the
	// server, those hidden from it by overwrites included, so a hidden channel is listed (and
	// judged on its permissions above), never mistaken for a deleted one.
	return ChannelUnusable{Reason: UnusableDeleted, Missing: []string{}}, true
}
