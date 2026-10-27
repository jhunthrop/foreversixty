package discordbot

import (
	"context"
	"fmt"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

// cmdRaid answers /raid: the guild's next raid as its post shows it, the description replaced
// by the viewer's own state, with the post's row. Ephemeral, no Share.
func (s *Service) cmdRaid(ctx context.Context, in Interaction) (Reply, error) {
	userID, linked, err := s.viewer(ctx, in)
	if err != nil {
		return Reply{}, fmt.Errorf("discordbot: raid viewer: %w", err)
	}
	if !linked {
		return pressReply(commandCopy.RaidUnlinked, false, s.linkButton()), nil
	}
	guildID, bound, err := s.serverGuild(ctx, in)
	if err != nil {
		return Reply{}, fmt.Errorf("discordbot: raid server: %w", err)
	}
	if !bound {
		return pressReply(commonCopy.Unbound, false), nil
	}
	ask := signupAsk{GuildID: guildID, UserID: userID}
	cands, err := s.d.Guilds.SignupCandidates(ctx, guildID, userID)
	if err != nil {
		return Reply{}, fmt.Errorf("discordbot: raid candidates: %w", err)
	}
	if len(cands) == 0 { // a non-member learns nothing of the calendar (R1)
		return s.notMember(ctx, ask)
	}
	now := s.now()
	next, err := s.d.Guilds.NextRaidFor(ctx, guildID, userID, nil, now)
	if err != nil {
		return Reply{}, fmt.Errorf("discordbot: raid next: %w", err)
	}
	if next == nil {
		return pressReply(commandCopy.RaidNone, false), nil
	}
	facts, found, err := s.d.Guilds.RaidPostFacts(ctx, next.ID)
	if err != nil {
		return Reply{}, fmt.Errorf("discordbot: raid facts: %w", err)
	}
	if !found {
		return pressReply(commandCopy.RaidNone, false), nil
	}
	state := PostState(facts.Status, facts.Post.StartsAt, facts.EndsAt, nil, now)
	msg := PostMessage(postFor(facts, state, nil), facts.RaidURL, facts.NightURL, now)
	msg.Embeds[0].Description = ownStateSentence(next.Own, cands)
	msg.Ephemeral = true
	return Reply{Message: msg}, nil
}

// ownStateSentence is the viewer's state on the raid, named by the character that signed.
func ownStateSentence(own *guilds.NextRaidOwn, cands []guilds.SignupCandidate) string {
	if own == nil {
		return commandCopy.RaidUnsigned
	}
	name := own.CharacterKey
	for _, c := range cands {
		if c.CharacterKey == own.CharacterKey {
			name = c.Name
		}
	}
	template := map[string]string{
		guilds.SignupIn: commandCopy.RaidIn, guilds.SignupLate: commandCopy.RaidLate,
		guilds.SignupBench: commandCopy.RaidBench, guilds.SignupOut: commandCopy.RaidOut,
	}[own.State]
	if template == "" {
		return commandCopy.RaidUnsigned
	}
	return fill(template, "{Name}", memberText(name), "{Role}", roleWord(own.Role))
}
