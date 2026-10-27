package discordbot

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

// /signup's option names.
const (
	optStatus    = "status"
	optRaid      = "raid"
	optCharacter = "character"
)

// cmdSignup is /signup status:<…> [raid] [character]: the press's flow, the raid defaulting to
// the guild's next planned one and the character option skipping the select.
func (s *Service) cmdSignup(ctx context.Context, in Interaction) (Reply, error) {
	status, _ := in.Option(optStatus)
	state := strings.ToLower(strings.TrimSpace(status))
	if !knownSignUpState(state) {
		return pressReply(pressCopy.Failure, false), nil
	}
	a, refused, ok := s.resolveAsker(ctx, in, false)
	if !ok {
		return refused, nil
	}
	a.State = state
	a.CharacterKey, _ = in.Option(optCharacter)
	raidRef, _ := in.Option(optRaid)
	if raidRef != "" {
		id, err := strconv.ParseInt(raidRef, 10, 64)
		if err != nil || id <= 0 {
			return a.reply(pressCopy.Failure), nil
		}
		a.RaidID = id
		// A named raid's state is told only to a member, so its ids cannot be probed.
		if rep, member, err := s.memberGate(ctx, a); err != nil || !member {
			return s.gateAnswer(in, rep, err), nil
		}
		return s.runFlow(ctx, in, a), nil
	}
	raidID, rep, err := s.nextRaidFor(ctx, a)
	if err != nil {
		return s.saveFailed(in, false, err), nil
	}
	if raidID == 0 {
		return rep, nil
	}
	a.RaidID = raidID
	return s.runFlow(ctx, in, a), nil
}

// memberGate answers an account with no verified character in the guild with the member
// sentences (R1: nothing of the calendar before that); member is true otherwise.
func (s *Service) memberGate(ctx context.Context, a signupAsk) (Reply, bool, error) {
	cands, err := s.freshCandidates(ctx, a.GuildID, a.UserID)
	if err != nil {
		return Reply{}, false, err
	}
	if len(verifiedOf(cands)) > 0 {
		return Reply{}, true, nil
	}
	rep, err := s.notSignable(ctx, a, cands)
	return rep, false, err
}

func (s *Service) gateAnswer(in Interaction, rep Reply, err error) Reply {
	if err != nil {
		return s.saveFailed(in, false, err)
	}
	return rep
}

// nextRaidFor is the guild's next signable raid, asked only for a member. raidID is 0 with the
// answer when there is none.
func (s *Service) nextRaidFor(ctx context.Context, a signupAsk) (int64, Reply, error) {
	if rep, member, err := s.memberGate(ctx, a); err != nil || !member {
		return 0, rep, err
	}
	raids, err := s.signableRaids(ctx, a.GuildID)
	if err != nil {
		return 0, Reply{}, err
	}
	if len(raids) == 0 {
		return 0, a.reply(pressCopy.NoRaid), nil
	}
	return raids[0].ID, Reply{}, nil
}

// signableRaids are the guild's Signable raids, soonest first.
func (s *Service) signableRaids(ctx context.Context, guildID int64) ([]guilds.Raid, error) {
	now := s.now()
	raids, err := s.d.Guilds.PlannedRaids(ctx, guildID, now)
	if err != nil {
		return nil, fmt.Errorf("discordbot: planned raids: %w", err)
	}
	out := make([]guilds.Raid, 0, len(raids))
	for _, r := range raids {
		if r.Signable(now) {
			out = append(out, r)
		}
	}
	return out, nil
}
