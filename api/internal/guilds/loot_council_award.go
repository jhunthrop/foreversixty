// api/internal/guilds/loot_council_award.go
//
// An award's decision record: main or off spec, a short reason, and the votes at award time
// (yes, no and how many officers could vote), snapshotted onto the award row so the record
// stands when the votes change later.
package guilds

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// awardReasonMax is the longest reason an award keeps, in characters.
const awardReasonMax = 200

// invalidNightText is the refusal for a night key that names no loot night of the guild.
const invalidNightText = "that is not a raid night of this guild"

// AwardDecision is what an officer adds to an award: the council night it is for (nil for the
// raid_id or current-night rule), main or off spec ("" to take the candidate's flag, else the
// derivation rule) and a reason ("" for none).
type AwardDecision struct {
	Night  *awardNight
	MsOs   string
	Reason string
}

// decisionSnapshot is the record the award keeps: the votes on characterKey for the drop now, the
// officer count, the MS/OS (given, else declared, else derived) and the reason.
func (s *Store) decisionSnapshot(ctx context.Context, guildID int64, night awardNight, itemID int, characterKey string, d AwardDecision, now time.Time) (awardDecisionRecord, error) {
	votes, err := s.candidateVotes(ctx, guildID, night, itemID, characterKey, 0, now)
	if err != nil {
		return awardDecisionRecord{}, err
	}
	rec := awardDecisionRecord{votesYes: &votes.Yes, votesNo: &votes.No, officers: &votes.Voters}
	if d.Reason != "" {
		rec.reason = &d.Reason
	}
	msOs, err := s.awardMsOs(ctx, guildID, night, itemID, characterKey, d.MsOs, now)
	if err != nil {
		return awardDecisionRecord{}, err
	}
	rec.msOs = msOs
	return rec, nil
}

// awardMsOs is the given flag, else the candidate's declared one, else the derivation rule
// (candidateMsOs); nil for a character the drop has no candidate row for.
func (s *Store) awardMsOs(ctx context.Context, guildID int64, night awardNight, itemID int, characterKey, given string, now time.Time) (*string, error) {
	if given != "" {
		return &given, nil
	}
	flags, err := s.dropInterest(ctx, guildID, night, itemID, now)
	if err != nil {
		return nil, err
	}
	if declared := flags[characterKey]; declared != "" {
		return &declared, nil
	}
	group, ok := lootGroupOf(itemID)
	if !ok {
		return nil, ErrInvalidAward
	}
	lr, err := s.lootRoster(ctx, guildID)
	if err != nil {
		return nil, err
	}
	iv := s.lootItemView(lr.roster, group, lr.guild.Faction)
	for _, c := range iv.Candidates {
		if c.CharacterKey == characterKey {
			msOs, _ := candidateMsOs(c, iv.Kind, "")
			return msOs, nil
		}
	}
	return nil, nil
}

// validMsOs reports whether v is "ms" or "os".
func validMsOs(v string) bool { return v == msOsMS || v == msOsOS }

// awardDecisionFrom reads and checks the award body's decision fields, writing the refusal and
// returning false when they are bad.
func (s *Service) awardDecisionFrom(w http.ResponseWriter, r *http.Request, guildID int64, in lootAwardInput) (AwardDecision, bool) {
	var d AwardDecision
	if in.MsOs != nil {
		if !validMsOs(*in.MsOs) {
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", `ms_os must be "ms" or "os"`, nil)
			return d, false
		}
		d.MsOs = *in.MsOs
	}
	if in.Reason != nil {
		d.Reason = strings.TrimSpace(*in.Reason)
		if utf8.RuneCountInString(d.Reason) > awardReasonMax {
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "a reason is at most 200 characters", nil)
			return d, false
		}
	}
	if in.Night == "" {
		return d, true
	}
	night, ok := s.councilNightFrom(w, r, guildID, in.Night, "loot_award", "could not record that award just now")
	if !ok {
		return d, false
	}
	d.Night = &night
	return d, true
}

// councilNightFrom resolves a night key (councilNight), writing the refusal and returning false
// when it names no night of the guild.
func (s *Service) councilNightFrom(w http.ResponseWriter, r *http.Request, guildID int64, key, op, failMessage string) (awardNight, bool) {
	night, err := s.Store.councilNight(r.Context(), guildID, key, time.Now())
	switch {
	case err == nil:
		return night, true
	case errors.Is(err, ErrInvalidNight):
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", invalidNightText, nil)
	default:
		s.fail(w, r, op, err, failMessage)
	}
	return awardNight{}, false
}
