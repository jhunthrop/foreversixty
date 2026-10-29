package discordbot

import (
	"context"
	"fmt"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
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
	// Read fresh on purpose, not through a cache: by ruling a press always answers from the
	// member's characters as they stand now.
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
	movedFrom, err := s.d.Store.MovedFrom(ctx, next.ID)
	if err != nil {
		return Reply{}, fmt.Errorf("discordbot: raid moved: %w", err)
	}
	msg := PostMessage(postFor(facts, s.postStateOf(facts, movedFrom), movedFrom), facts.RaidURL, facts.NightURL, now)
	msg.Embeds[0] = s.raidAnswerEmbed(ctx, facts, movedFrom, next.Own, cands)
	msg.Ephemeral = true
	return Reply{Message: msg}, nil
}

// ownStateSentence is the viewer's state on the raid, named by the character that signed.
// A character missing from the candidates is not named (its key is not a name).
func ownStateSentence(own *guilds.NextRaidOwn, cands []guilds.SignupCandidate) string {
	if own == nil {
		return commandCopy.RaidUnsigned
	}
	name, named := "", false
	for _, c := range cands {
		if c.CharacterKey == own.CharacterKey {
			name, named = memberText(c.Name), true
		}
	}
	templates := map[string][2]string{
		guilds.SignupIn:    {commandCopy.RaidInNameless, commandCopy.RaidIn},
		guilds.SignupLate:  {commandCopy.RaidLateNameless, commandCopy.RaidLate},
		guilds.SignupBench: {commandCopy.RaidBenchNameless, commandCopy.RaidBench},
		guilds.SignupOut:   {commandCopy.RaidOutNameless, commandCopy.RaidOut},
	}
	pair, ok := templates[own.State]
	if !ok {
		return commandCopy.RaidUnsigned
	}
	template := pair[0]
	if named {
		template = pair[1]
	}
	return fill(template, "{Name}", name, "{Role}", roleWord(own.Role))
}

// postStateOf is the state the raid's post shows now.
func (s *Service) postStateOf(facts guilds.RaidPostFacts, movedFrom *time.Time) notify.RaidPostState {
	return PostState(facts.Status, facts.Post.StartsAt, facts.EndsAt, movedFrom, s.now())
}

// raidAnswerEmbed is the raid's post embed with its description replaced by the viewer's own
// state: the /raid answer, and the embed of the /signup card.
func (s *Service) raidAnswerEmbed(_ context.Context, facts guilds.RaidPostFacts, movedFrom *time.Time, own *guilds.NextRaidOwn, cands []guilds.SignupCandidate) notify.Embed {
	embed := notify.RaidPostEmbed(postFor(facts, s.postStateOf(facts, movedFrom), movedFrom), s.now())
	embed.Description = ownStateSentence(own, cands)
	return embed
}
