package discordbot

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

// Autocomplete bounds: Discord takes at most 25 choices, each name at most 100.
const (
	maxChoices     = 25
	maxChoiceRunes = 100
	// raidChoiceDate is "Thu 10 Dec", in the guild's raid zone.
	raidChoiceDate = "Mon 2 Jan"
	choiceJoin     = " · "
)

// autocompleteSignup suggests /signup's raid (the guild's planned raids, soonest first) and
// character (the account's own verified characters in the guild). Anyone who is not a
// verified member through a verified character gets nothing (R1).
func (s *Service) autocompleteSignup(ctx context.Context, in Interaction) ([]Choice, error) {
	focused, typed := in.Focused()
	if focused != optRaid && focused != optCharacter {
		return []Choice{}, nil
	}
	userID, linked, err := s.viewer(ctx, in)
	if err != nil || !linked {
		return []Choice{}, err
	}
	guildID, bound, err := s.serverGuild(ctx, in)
	if err != nil || !bound {
		return []Choice{}, err
	}
	cands, err := s.cachedCandidates(ctx, guildID, userID)
	if err != nil {
		return nil, err
	}
	verified := verifiedOf(cands)
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
	loc := raidLocation(zone)
	out := []Choice{}
	for _, r := range raids {
		if len(out) == maxChoices {
			break
		}
		date := choiceJoin + r.StartsAt.In(loc).Format(raidChoiceDate)
		title := truncateRunes(r.Title, maxChoiceRunes-len([]rune(date)))
		if matches(title, typed) {
			out = append(out, Choice{Name: title + date, Value: strconv.FormatInt(r.ID, 10)})
		}
	}
	return out, nil
}

// raidLocation is the guild's raid zone, UTC when unset or unknown.
func raidLocation(zone string) *time.Location {
	if zone == "" {
		return time.UTC
	}
	loc, err := time.LoadLocation(zone)
	if err != nil {
		return time.UTC
	}
	return loc
}

func matches(name, typed string) bool {
	return strings.Contains(strings.ToLower(name), strings.ToLower(strings.TrimSpace(typed)))
}
