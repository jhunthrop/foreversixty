// api/internal/guilds/raid_answer_addon.go
//
// The addon's raid answer loop (docs/contracts/addon-v2.md contracts 3 and 4; spec
// 2026-11-08-addon-ux §4.3): the inbox's raid_signup rows, one per open raid of a character's
// guild in the next week, and the outbox's raid_answer entries, the answers a player gave in
// game. An answer from the game applies only while the raid is open to the character and only
// when it is newer than the stored row, so a later answer from the app, the site or Discord
// wins; it keeps the stored role, else the character's, and rule R3 holds as for any sign-up.
package guilds

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"time"
)

// AddonSignupWindow is how far ahead the inbox lists a guild's raids.
const AddonSignupWindow = 7 * 24 * time.Hour

// AddonAnswerNone is a raid_signup's answer for a character with no sign-up row.
const AddonAnswerNone = "none"

// What became of one addon answer: accepted (stored, or already what was stored) or refused,
// with the reason. The companion removes both from the game's file; anything else (a failed
// read) fails the upload so the companion sends it again.
const (
	AddonAnswerAccepted = "accepted"
	AddonAnswerRefused  = "refused"

	AddonAnswerStale    = "stale"
	AddonAnswerClosed   = "closed"
	AddonAnswerNotFound = "not_found"
	AddonAnswerNotYours = "not_yours"
	AddonAnswerInvalid  = "invalid"
)

// AddonRaidAnswer is one answer the player gave in game for one of their own characters.
type AddonRaidAnswer struct {
	UserID       int64
	CharacterKey string
	RaidID       int64
	State        string
	LateMinutes  *int16
	AnsweredAt   time.Time
}

// AddonAnswerOutcome is what became of an AddonRaidAnswer; Reason is set on a refusal.
type AddonAnswerOutcome struct {
	Status string `json:"status"`
	Reason string `json:"reason,omitempty"`
}

func refusedFor(reason string) AddonAnswerOutcome {
	return AddonAnswerOutcome{Status: AddonAnswerRefused, Reason: reason}
}

// AnswerFromAddon applies one answer from the game. Refusals are outcomes; only a failed read
// or write is an error.
func (s *Service) AnswerFromAddon(ctx context.Context, a AddonRaidAnswer) (AddonAnswerOutcome, error) {
	if !slices.Contains([]string{SignupIn, SignupLate, SignupBench, SignupOut}, a.State) ||
		!validLateMinutes(a.State, a.LateMinutes) || a.AnsweredAt.IsZero() {
		return refusedFor(AddonAnswerInvalid), nil
	}
	raid, found, err := s.Store.raidByID(ctx, a.RaidID)
	if err != nil || !found {
		return refusedFor(AddonAnswerNotFound), err
	}
	// A logged or past raid is an officer's correction on the site, never a game answer.
	if raid.Status == RaidDone || time.Now().After(raid.EndsAt()) {
		return refusedFor(AddonAnswerClosed), nil
	}
	at := a.AnsweredAt
	_, err = s.signUp(ctx, SignupCommand{GuildID: raid.GuildID, RaidID: raid.ID, ActorID: a.UserID,
		CharacterKey: a.CharacterKey, State: a.State, LateMinutes: a.LateMinutes, AnsweredAt: &at}, true)
	return addonOutcome(err)
}

// addonOutcome words a sign-up's error as an addon answer's outcome; an error it cannot word
// stays an error.
func addonOutcome(err error) (AddonAnswerOutcome, error) {
	switch {
	case err == nil:
		return AddonAnswerOutcome{Status: AddonAnswerAccepted}, nil
	case errors.Is(err, ErrSignupStale):
		return refusedFor(AddonAnswerStale), nil
	case errors.Is(err, ErrRaidLocked), errors.Is(err, ErrRaidCancelled), errors.Is(err, ErrRaidEnded),
		errors.Is(err, ErrClaimContested):
		return refusedFor(AddonAnswerClosed), nil
	case errors.Is(err, ErrNotFound):
		return refusedFor(AddonAnswerNotFound), nil
	case errors.Is(err, ErrNotMember), errors.Is(err, ErrCharacterNotFound), errors.Is(err, ErrNotYourCharacter),
		errors.Is(err, ErrCharacterNotVerified):
		return refusedFor(AddonAnswerNotYours), nil
	}
	var role *RoleNotOfClassError
	if errors.As(err, &role) {
		return refusedFor(AddonAnswerInvalid), nil
	}
	return AddonAnswerOutcome{}, err
}

// AddonRaidSignup is one raid_signup row: a raid of the character's guild and the character's
// answer for it. Role, LateMinutes and AnsweredAt are nil with no sign-up row (Answer "none").
type AddonRaidSignup struct {
	CharacterKey string
	RaidID       int64
	ZoneName     string
	Title        string
	StartsAt     time.Time
	Answer       string
	Role         string
	LateMinutes  *int16
	AnsweredAt   *time.Time
}

// AddonRaidSignups is the Store's, so one Service answers the addon both ways.
func (s *Service) AddonRaidSignups(ctx context.Context, keys []string, now time.Time) ([]AddonRaidSignup, error) {
	return s.Store.AddonRaidSignups(ctx, keys, now)
}

// AddonRaidSignups is, for each of keys verified in a guild, every planned or locked raid of
// that guild not over at now and starting within AddonSignupWindow, by character then start.
func (s *Store) AddonRaidSignups(ctx context.Context, keys []string, now time.Time) ([]AddonRaidSignup, error) {
	if len(keys) == 0 {
		return nil, nil
	}
	rows, err := s.Pool.Query(ctx, `
		select gc.character_key, r.id, r.zone, r.title, r.starts_at,
		       coalesce(sg.state, $4), coalesce(sg.role, ''), sg.late_minutes, sg.updated_at
		from guild_characters gc
		join guild_raids r on r.guild_id = gc.guild_id
		left join guild_raid_signups sg on sg.raid_id = r.id and sg.character_key = gc.character_key
		where gc.character_key = any($1) and gc.verified_at is not null and r.status in ('planned', 'locked')
		  and r.starts_at + make_interval(mins => r.duration_minutes) > $2 and r.starts_at < $3
		order by gc.character_key, r.starts_at, r.id`, keys, now, now.Add(AddonSignupWindow), AddonAnswerNone)
	if err != nil {
		return nil, fmt.Errorf("guilds: addon raid signups: %w", err)
	}
	defer rows.Close()
	var out []AddonRaidSignup
	for rows.Next() {
		var r AddonRaidSignup
		var zone string
		if err := rows.Scan(&r.CharacterKey, &r.RaidID, &zone, &r.Title, &r.StartsAt, &r.Answer, &r.Role,
			&r.LateMinutes, &r.AnsweredAt); err != nil {
			return nil, fmt.Errorf("guilds: addon raid signups: %w", err)
		}
		r.ZoneName = zoneOf(zone).Name
		out = append(out, r)
	}
	return out, rows.Err()
}
