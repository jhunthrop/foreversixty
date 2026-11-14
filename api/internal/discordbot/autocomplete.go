package discordbot

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/textfmt"
)

// Autocomplete bounds: Discord takes at most 25 choices, each name at most 100.
const (
	maxChoices     = 25
	maxChoiceRunes = 100
	choiceJoin     = " · "
)

// autocompleteSignup suggests /signup's raid (the guild's planned raids, soonest first) and
// character (the account's own verified characters in the guild). A Discord user with no site
// account, or a Discord-only one, is also offered the guild's directory roster as character
// picks (pickChoices). Otherwise anyone who is not a verified member through a verified
// character gets nothing (R1), and the raid choices stay members only.
func (s *Service) autocompleteSignup(ctx context.Context, in Interaction) ([]Choice, error) {
	focused, typed := in.Focused()
	if focused != optRaid && focused != optCharacter {
		return []Choice{}, nil
	}
	userID, linked, err := s.viewer(ctx, in)
	if err != nil {
		return []Choice{}, err
	}
	guildID, bound, err := s.serverGuild(ctx, in)
	if err != nil || !bound {
		return []Choice{}, err
	}
	var verified []guilds.SignupCandidate
	if linked {
		cands, err := s.cachedCandidates(ctx, guildID, userID)
		if err != nil {
			return nil, err
		}
		verified = verifiedOf(cands)
	}
	if focused == optCharacter {
		if may, err := s.mayPick(ctx, userID, linked); err != nil || may {
			if err != nil {
				return nil, err
			}
			return s.pickChoices(ctx, guildID, verified, typed)
		}
	}
	if len(verified) == 0 {
		return []Choice{}, nil
	}
	if focused == optCharacter {
		return characterChoices(verified, typed), nil
	}
	return s.raidChoices(ctx, guildID, typed)
}

func characterChoices(verified []guilds.SignupCandidate, typed string) []Choice {
	out := []Choice{}
	for _, c := range verified {
		if len(out) == maxChoices {
			break
		}
		if len([]rune(c.CharacterKey)) <= maxChoiceRunes && matches(c.Name, typed) {
			out = append(out, Choice{Name: truncateRunes(c.Name, maxChoiceRunes), Value: c.CharacterKey})
		}
	}
	return out
}

func (s *Service) raidChoices(ctx context.Context, guildID int64, typed string) ([]Choice, error) {
	raids, err := s.signableRaids(ctx, guildID)
	if err != nil {
		return nil, err
	}
	zone, err := s.d.Guilds.GuildRaidTimezone(ctx, guildID)
	if err != nil {
		return nil, fmt.Errorf("discordbot: signup autocomplete zone: %w", err)
	}
	loc := textfmt.Location(zone)
	out := []Choice{}
	for _, r := range raids {
		if len(out) == maxChoices {
			break
		}
		date := choiceJoin + textfmt.RaidDay(r.StartsAt, loc)
		title := truncateRunes(r.Title, maxChoiceRunes-len([]rune(date)))
		if matches(title, typed) {
			out = append(out, Choice{Name: title + date, Value: strconv.FormatInt(r.ID, 10)})
		}
	}
	return out, nil
}

func matches(name, typed string) bool {
	return strings.Contains(strings.ToLower(name), strings.ToLower(strings.TrimSpace(typed)))
}
