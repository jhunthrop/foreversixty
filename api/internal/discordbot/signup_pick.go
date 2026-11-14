package discordbot

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/directory"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

// DiscordClaims is what the sign-up by character pick needs of the directory (directory.Claims).
type DiscordClaims interface {
	EnsureDiscordAccount(ctx context.Context, discordUserID, displayName string) (int64, bool, error)
	IsDiscordOnly(ctx context.Context, userID int64) (bool, error)
	GuildPicks(ctx context.Context, guildID int64, typed string, limit int) ([]directory.Pick, error)
	ClaimForDiscord(ctx context.Context, guildID, userID, bnetID int64) (directory.Claim, error)
	MergeDiscordOnly(ctx context.Context, from, into int64) error
}

// mayPick reports whether the account picks its characters from the directory: no account
// yet (linked false), or a Discord-only one. An account with a Battle.net sign-in never does.
func (s *Service) mayPick(ctx context.Context, userID int64, linked bool) (bool, error) {
	if s.d.Claims == nil {
		return false, nil
	}
	if !linked {
		return true, nil
	}
	return s.d.Claims.IsDiscordOnly(ctx, userID)
}

// pickChoices is /signup's character autocomplete for an account that picks: its own picked
// characters first, then the guild's directory roster no account holds, at most 25 in all.
func (s *Service) pickChoices(ctx context.Context, guildID int64, verified []guilds.SignupCandidate, typed string) ([]Choice, error) {
	out := characterChoices(verified, typed)
	if len(out) >= maxChoices {
		return out, nil
	}
	picks, err := s.d.Claims.GuildPicks(ctx, guildID, typed, maxChoices-len(out))
	if err != nil {
		return nil, fmt.Errorf("discordbot: guild picks: %w", err)
	}
	for _, p := range picks {
		out = append(out, Choice{Name: truncateRunes(pickName(p), maxChoiceRunes), Value: directory.PickValue(p.BnetID)})
	}
	return out, nil
}

// pickName is a pick's autocomplete line, "Vela Thornbrook · 41 Mage"; the level and class
// drop out when the directory has neither.
func pickName(p directory.Pick) string {
	if p.Level == 0 && p.Class == "" {
		return p.Name
	}
	class := p.Class
	if class != "" {
		class = strings.ToUpper(class[:1]) + class[1:]
	}
	line := fill(pickCopy.Choice, "{Name}", p.Name, "{Level}", strconv.Itoa(p.Level), "{Class}", class)
	return strings.Join(strings.Fields(line), " ")
}

// signupByPick is /signup with a picked directory character: the Discord-only account (made on
// first use), the claim, then the sign-up as /signup runs it, In unless a status was given.
// A new claim's answer leads with the line that officers see it marked.
func (s *Service) signupByPick(ctx context.Context, in Interaction, state string, bnetID int64) Reply {
	guildID, bound, err := s.serverGuild(ctx, in)
	if err != nil {
		return s.saveFailed(in, false, fmt.Errorf("discordbot: pick server: %w", err))
	}
	if !bound {
		return pressReply(commonCopy.Unbound, false)
	}
	userID, refusal, ok := s.pickAccount(ctx, in)
	if !ok {
		return refusal
	}
	claim, err := s.d.Claims.ClaimForDiscord(ctx, guildID, userID, bnetID)
	switch {
	case errors.Is(err, directory.ErrTaken):
		return pressReply(fill(pickCopy.Taken, "{Name}", memberText(claim.Name)), false)
	case errors.Is(err, directory.ErrTooManyPicks):
		return pressReply(fill(pickCopy.TooMany, "{n}", strconv.Itoa(directory.MaxDiscordPicks)), false)
	case errors.Is(err, directory.ErrNotInGuild), errors.Is(err, directory.ErrNotDiscordOnly):
		return pressReply(pressCopy.Failure, false)
	case err != nil:
		return s.saveFailed(in, false, err)
	}
	if state == "" {
		state = guilds.SignupIn
	}
	a := signupAsk{GuildID: guildID, UserID: userID, State: state, CharacterKey: claim.Key}
	rep := s.signupNamedRaid(ctx, in, a)
	if claim.Created {
		return leadWith(rep, fill(pickCopy.Claimed, "{Name}", memberText(claim.Name)))
	}
	return rep
}

// leadWith puts line above an answer: its first Text Display in a Components V2 message (the
// spec card), else the first line of its content.
func leadWith(rep Reply, line string) Reply {
	m := rep.Message
	switch {
	case m.ComponentsV2:
		m.Layout = append([]Node{textNode(line)}, m.Layout...)
	case m.Content == "":
		m.Content = line
	default:
		m.Content = line + "\n" + m.Content
	}
	rep.Message = m
	return rep
}

// pickAccount is the account a pick signs up with: the Discord-only account, made now when the
// Discord user has none. ok is false with the answer for an account with a Battle.net sign-in.
func (s *Service) pickAccount(ctx context.Context, in Interaction) (int64, Reply, bool) {
	userID, linked, err := s.viewer(ctx, in)
	if err != nil {
		return 0, s.saveFailed(in, false, fmt.Errorf("discordbot: pick viewer: %w", err)), false
	}
	if may, err := s.mayPick(ctx, userID, linked); err != nil || !may {
		if err != nil {
			return 0, s.saveFailed(in, false, err), false
		}
		return 0, pressReply(pickCopy.OwnAccount, false), false
	}
	if linked {
		return userID, Reply{}, true
	}
	userID, _, err = s.d.Claims.EnsureDiscordAccount(ctx, in.DiscordUserID, in.DisplayName)
	if err != nil {
		return 0, s.saveFailed(in, false, fmt.Errorf("discordbot: discord account: %w", err)), false
	}
	s.log.Info("discordbot", "op", "discord_only_account", "user_id", userID)
	return userID, Reply{}, true
}
