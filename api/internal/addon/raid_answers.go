// api/internal/addon/raid_answers.go
//
// The addon's raid answer loop (docs/contracts/addon-v2.md; spec 2026-11-08-addon-ux §4.3).
// Contract 3's raid_signup message is one per open raid of a character's guild in the next
// week, answered or not, so the addon panel can show Tonight with the answer the site holds.
// Contract 4's raid_answer outbox entries are the answers the player gave in game; each is
// applied (guilds.Service.AnswerFromAddon) and answered accepted or refused, which the
// companion removes from the game's file, or deferred, which it keeps and sends again.
package addon

import (
	"context"
	"encoding/json"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

const messageRaidSignup = "raid_signup"

// maxRaidAnswers bounds the answers one upload applies; the rest are deferred to the next.
const maxRaidAnswers = 20

// RaidAnswerDeferred is an answer this upload did not apply: past maxRaidAnswers, or with no
// sign-up source wired. The companion keeps it and sends it again.
const RaidAnswerDeferred = "deferred"

// RaidAnswerSource is guilds.Store's raid_signup read and guilds.Service's answer, together.
type RaidAnswerSource interface {
	AddonRaidSignups(ctx context.Context, keys []string, now time.Time) ([]guilds.AddonRaidSignup, error)
	AnswerFromAddon(ctx context.Context, a guilds.AddonRaidAnswer) (guilds.AddonAnswerOutcome, error)
}

// raidSignupMessages is one raid_signup message per open raid of each of keys' guilds in the
// next week; none on a failed read, which is logged (the sync never fails over it).
func (s *Service) raidSignupMessages(ctx context.Context, keys []string, now time.Time) []MessageEntry {
	if s.RaidAnswers == nil || len(keys) == 0 {
		return nil
	}
	rows, err := s.RaidAnswers.AddonRaidSignups(ctx, keys, now)
	if err != nil {
		s.logger().Warn("addon", "op", "messages_raid_signup", "err", err)
		return nil
	}
	out := make([]MessageEntry, 0, len(rows))
	for _, r := range rows {
		m := MessageEntry{Type: messageRaidSignup, Character: r.CharacterKey, RaidID: r.RaidID, Raid: r.ZoneName,
			Title: r.Title, StartsAt: r.StartsAt.UTC().Format(time.RFC3339), Answer: r.Answer, Role: r.Role}
		if r.LateMinutes != nil {
			m.LateMinutes = int(*r.LateMinutes)
		}
		if r.AnsweredAt != nil {
			m.AnsweredAt = r.AnsweredAt.UTC().Format(time.RFC3339)
		}
		out = append(out, m)
	}
	return out
}

// RaidAnswerEntry is one outbox raid_answer entry: Character is the export's Name-Realm form,
// AnsweredAt the client's time(), Unix seconds UTC.
type RaidAnswerEntry struct {
	Character   string `json:"character"`
	RaidID      int64  `json:"raid_id"`
	Answer      string `json:"answer"`
	LateMinutes *int   `json:"late_minutes,omitempty"`
	AnsweredAt  int64  `json:"answered_at"`
}

// RaidAnswerResult is what became of one entry, echoing the entry's character, raid and time so
// the companion can find it in the game's file.
type RaidAnswerResult struct {
	Character  string `json:"character"`
	RaidID     int64  `json:"raid_id"`
	AnsweredAt int64  `json:"answered_at"`
	Status     string `json:"status"`
	Reason     string `json:"reason,omitempty"`
}

// readRaidAnswers decodes the outbox's raid_answer list entry by entry; an entry that is not
// one decodes as the zero entry, which is answered invalid. Anything but a list reads as none.
func readRaidAnswers(raw json.RawMessage) []RaidAnswerEntry {
	var list []json.RawMessage
	if err := json.Unmarshal(raw, &list); err != nil {
		return nil
	}
	out := make([]RaidAnswerEntry, len(list))
	for i, item := range list {
		var e RaidAnswerEntry
		if json.Unmarshal(item, &e) == nil {
			out[i] = e
		}
	}
	return out
}

// applyRaidAnswers applies each entry for the account's own character it names and answers it.
// Only a failed read or write is an error: the upload fails and the companion sends it again.
func (s *Service) applyRaidAnswers(ctx context.Context, userID int64, exports []Export, entries []RaidAnswerEntry) ([]RaidAnswerResult, error) {
	out := make([]RaidAnswerResult, len(entries))
	for i, e := range entries {
		out[i] = RaidAnswerResult{Character: e.Character, RaidID: e.RaidID, AnsweredAt: e.AnsweredAt}
		if s.RaidAnswers == nil || i >= maxRaidAnswers {
			out[i].Status = RaidAnswerDeferred
			continue
		}
		outcome, err := s.applyRaidAnswer(ctx, userID, exports, e)
		if err != nil {
			return nil, err
		}
		out[i].Status, out[i].Reason = outcome.Status, outcome.Reason
	}
	return out, nil
}

// applyRaidAnswer answers one entry: invalid when it is not a usable entry, unknown_character
// when it names none of the account's synced characters, else what the sign-up made of it.
func (s *Service) applyRaidAnswer(ctx context.Context, userID int64, exports []Export, e RaidAnswerEntry) (guilds.AddonAnswerOutcome, error) {
	refused := func(reason string) (guilds.AddonAnswerOutcome, error) {
		return guilds.AddonAnswerOutcome{Status: guilds.AddonAnswerRefused, Reason: reason}, nil
	}
	if e.RaidID <= 0 || e.AnsweredAt <= 0 || !validUnit(e.Character) {
		return refused(guilds.AddonAnswerInvalid)
	}
	own, ok := resolveOwn(exports, e.Character)
	if !ok {
		return refused(OutboxUnknownCharacter)
	}
	var late *int16
	if e.LateMinutes != nil {
		m := int16(*e.LateMinutes)
		late = &m
	}
	return s.RaidAnswers.AnswerFromAddon(ctx, guilds.AddonRaidAnswer{
		UserID: userID, CharacterKey: character.Key(own.Region, character.RulesetFromRealm(own.Ruleset, ""), own.Name),
		RaidID: e.RaidID, State: e.Answer, LateMinutes: late, AnsweredAt: time.Unix(e.AnsweredAt, 0).UTC(),
	})
}
