// api/internal/guilds/raids_signup_service.go
//
// The raid sign-up as a function of a command rather than of an HTTP request: the site's PUT
// route and the Discord bot's buttons both run it, so both answer to the same standing rules
// (G1 R1, R3) and the same per-account budget.
package guilds

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
)

// The refusals SignUp answers with, besides the store's ErrRaidLocked, ErrRaidCancelled and
// ErrNotFound (an unknown raid).
var (
	ErrNotMember            = errors.New("guilds: actor is not a verified member of the guild")
	ErrCharacterNotFound    = errors.New("guilds: no such character in this guild")
	ErrNotYourCharacter     = errors.New("guilds: that character belongs to another account")
	ErrCharacterNotVerified = errors.New("guilds: character is not verified in this guild")
	ErrSignupRateLimited    = errors.New("guilds: too many sign-ups just now")
	ErrPlanRequired         = errors.New("guilds: the guild plan is required")
	ErrClaimContested       = errors.New("guilds: the guild's claim is contested")
)

// errOfficerTargetUnverified is ErrCharacterNotVerified for an officer answering for another
// account's character; the PUT words it differently from a member's own pending character.
var errOfficerTargetUnverified = fmt.Errorf("%w: officer target", ErrCharacterNotVerified)

// RaidPostMarker is told when a raid's channel post is out of date. nil on the Service means
// nothing is told.
type RaidPostMarker interface {
	MarkRaidDirty(ctx context.Context, raidID int64)
	// MarkRaidMoved is MarkRaidDirty for a start that moved from was, announced or not: the
	// post shows the Moved state either way.
	MarkRaidMoved(ctx context.Context, raidID int64, was time.Time)
}

// SignupCommand is one sign-up answer. Role "" is the character's default; a nil Note keeps
// the stored note ("" on a first answer).
type SignupCommand struct {
	GuildID, RaidID, ActorID int64
	CharacterKey, State      string
	Role                     string
	Note                     *string
}

// SignupCandidate is one of an account's characters in a guild.
type SignupCandidate struct {
	CharacterKey string
	Name         string
	Class        string
	Spec         string
	Role         string
	// RaidSpec is the picked raid spec (a specs.json key), "" when never picked: the bot asks
	// "Which spec tonight?" before an In, Late or Bench while it is empty (spec §P.10.2).
	RaidSpec string
	Verified bool
	Main     bool
}

// SignupResult is the raid as it now stands, the character answered for, the character R3
// moved to bench (nil when none), and whether the stored row already said this (no write).
type SignupResult struct {
	Raid      RaidView
	Character SignupCandidate
	Moved     *string
	Unchanged bool
}

// SignUp answers for a character: it spends the account's sign-up budget, then runs signUp.
func (s *Service) SignUp(ctx context.Context, cmd SignupCommand) (SignupResult, error) {
	if !s.signupLimiter().Allow(signupLimitKey(cmd.ActorID)) {
		return SignupResult{}, ErrSignupRateLimited
	}
	return s.signUp(ctx, cmd, true)
}

// signUp is SignUp past the budget (the PUT route spends it in its middleware). With
// skipUnchanged a row that already holds the asked state, role and note is left alone; the
// PUT route always writes, as it always has, so a re-PUT refreshes the snapshot.
func (s *Service) signUp(ctx context.Context, cmd SignupCommand, skipUnchanged bool) (SignupResult, error) {
	raid, char, err := s.signupTarget(ctx, cmd)
	if err != nil {
		return SignupResult{}, err
	}
	asOfficer, err := s.signupStanding(ctx, cmd.ActorID, raid, char)
	if err != nil {
		return SignupResult{}, err
	}
	role := cmd.Role
	if role == "" {
		role = char.defaultRole()
	}
	result := SignupResult{Character: char.candidate(role, false)}
	// SignedBy is the character's account, not the actor: R3 groups by it, so an officer's
	// write for a member benches the member's other character, never the officer's.
	sg := Signup{
		RaidID: raid.ID, CharacterKey: char.CharacterKey, Name: char.Name, Class: char.Class, Spec: char.Spec,
		Role: role, State: cmd.State, SignedBy: char.userID, ActedBy: cmd.ActorID, AsOfficer: asOfficer,
		KeepNote: cmd.Note == nil,
	}
	if cmd.Note != nil {
		sg.Note = *cmd.Note
	}
	was, err := s.rosterPlaceBefore(ctx, raid, cmd, char, asOfficer)
	if err != nil {
		return SignupResult{}, err
	}
	wrote := true
	if skipUnchanged {
		result.Moved, wrote, err = s.Store.UpsertSignupIfChanged(ctx, sg)
	} else {
		result.Moved, err = s.Store.UpsertSignup(ctx, sg)
	}
	if err != nil {
		return SignupResult{}, err
	}
	if !wrote {
		result.Unchanged = true
		result.Raid, err = s.signupRaidView(ctx, cmd)
		return result, err
	}
	s.logger().Info("guilds", "op", "raid_signup", "guild_id", cmd.GuildID, "raid_id", raid.ID,
		"user_id", char.userID, "actor", cmd.ActorID, "character_key", char.CharacterKey, "as_officer", asOfficer)
	s.markRaid(ctx, raid.ID)
	if cmd.State == SignupOut {
		s.publishLateChange(ctx, raid, char.CharacterKey, char.Name, time.Now())
	}
	if was != nil && *was != cmd.State {
		s.publishRosterMoved(ctx, raid, char.userID, char.Name, *was, cmd.State)
	}
	result.Raid, err = s.signupRaidView(ctx, cmd)
	return result, err
}

// signupTarget runs the first three checks: the actor is a verified member, the raid is the
// guild's, the character is on the guild's roster.
func (s *Service) signupTarget(ctx context.Context, cmd SignupCommand) (Raid, raidRosterChar, error) {
	_, verified, err := s.Accounts.GuildRank(ctx, cmd.GuildID, cmd.ActorID)
	if err != nil {
		return Raid{}, raidRosterChar{}, err
	}
	if !verified {
		return Raid{}, raidRosterChar{}, ErrNotMember
	}
	raid, found, err := s.Store.GetRaid(ctx, cmd.GuildID, cmd.RaidID)
	if err != nil {
		return Raid{}, raidRosterChar{}, err
	}
	if !found {
		return Raid{}, raidRosterChar{}, ErrNotFound
	}
	chars, err := s.Store.raidRoster(ctx, cmd.GuildID, cmd.CharacterKey)
	if err != nil {
		return Raid{}, raidRosterChar{}, err
	}
	if len(chars) == 0 {
		return Raid{}, raidRosterChar{}, ErrCharacterNotFound
	}
	return raid, chars[0], nil
}

// signupStanding decides whether the actor may answer for char, and whether as an officer
// (which a locked raid needs).
func (s *Service) signupStanding(ctx context.Context, actorID int64, raid Raid, char raidRosterChar) (bool, error) {
	if char.userID != actorID {
		// Someone else's character: the full officer gate (rank, plan, freeze).
		if err := s.officerGateFor(ctx, raid.GuildID, actorID); err != nil {
			return false, err
		}
		if !char.verified {
			return false, errOfficerTargetUnverified
		}
		return true, nil
	}
	if !char.verified {
		return false, ErrCharacterNotVerified
	}
	if raid.Status != RaidLocked && raid.Status != RaidDone {
		return false, nil
	}
	officer, err := s.officerOrLeaderOf(ctx, raid.GuildID, actorID)
	if err != nil {
		return false, err
	}
	if !officer || !s.officerTools(ctx, raid.GuildID, actorID) {
		return false, nil
	}
	// Writing to a locked raid is an officer power, frozen like the others.
	if err := s.notFrozen(ctx, raid.GuildID, actorID); err != nil {
		return false, err
	}
	return true, nil
}

// officerGateFor is officerRaidGate as errors: a verified officer or leader (else
// ErrNotYourCharacter, the only reason the sign-up asks), the guild plan, the claim freeze.
func (s *Service) officerGateFor(ctx context.Context, guildID, actorID int64) error {
	officer, err := s.officerOrLeaderOf(ctx, guildID, actorID)
	if err != nil {
		return err
	}
	if !officer {
		return ErrNotYourCharacter
	}
	if s.Entitlements == nil {
		return ErrPlanRequired
	}
	ok, _, err := s.Entitlements.CanForGuild(ctx, guildID, actorID, entitlements.FeatureOfficerViews)
	if err != nil {
		return err
	}
	if !ok {
		return ErrPlanRequired
	}
	return s.notFrozen(ctx, guildID, actorID)
}

// notFrozen refuses while the guild's claim is contested. A moderator's bypass counts only
// when the context's actor is the very account answering.
func (s *Service) notFrozen(ctx context.Context, guildID, actorID int64) error {
	ctxActor := auth.ActorFrom(ctx)
	frozen, err := s.freezeCheck(ctx, guildID, ctxActor.UserID == actorID && ctxActor.IsModerator())
	if err != nil {
		return err
	}
	if frozen {
		return ErrClaimContested
	}
	return nil
}

// signupRaidView re-reads the raid (its status may have moved meanwhile) as the actor sees it.
func (s *Service) signupRaidView(ctx context.Context, cmd SignupCommand) (RaidView, error) {
	raid, found, err := s.Store.GetRaid(ctx, cmd.GuildID, cmd.RaidID)
	if err != nil {
		return RaidView{}, err
	}
	if !found {
		return RaidView{}, ErrNotFound
	}
	return s.raidViewOf(ctx, cmd.GuildID, cmd.ActorID, raid)
}

// SignupCandidates lists the account's characters in the guild: its main first, then by
// name. A pending character is included with Verified false.
func (s *Service) SignupCandidates(ctx context.Context, guildID, userID int64) ([]SignupCandidate, error) {
	chars, err := s.Store.raidRoster(ctx, guildID, "")
	if err != nil {
		return nil, err
	}
	mainKey, err := s.Store.mainCharacterKey(ctx, userID)
	if err != nil {
		return nil, err
	}
	out := []SignupCandidate{}
	for _, c := range chars {
		if c.userID == userID {
			cand := c.candidate(c.defaultRole(), c.CharacterKey == mainKey)
			cand.Verified = c.verified
			out = append(out, cand)
		}
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Main != out[j].Main {
			return out[i].Main
		}
		return strings.ToLower(out[i].Name) < strings.ToLower(out[j].Name)
	})
	return out, nil
}

func (c raidRosterChar) candidate(role string, main bool) SignupCandidate {
	return SignupCandidate{
		CharacterKey: c.CharacterKey, Name: c.Name, Class: c.Class, Spec: c.Spec,
		Role: role, RaidSpec: c.raidSpec, Verified: c.verified, Main: main,
	}
}

// mainCharacterKey is the account's chosen main, "" when it has none.
func (s *Store) mainCharacterKey(ctx context.Context, userID int64) (string, error) {
	var key *string
	if err := s.Pool.QueryRow(ctx, `select main_character_key from users where id = $1`, userID).Scan(&key); err != nil {
		return "", fmt.Errorf("guilds: main character: %w", err)
	}
	if key == nil {
		return "", nil
	}
	return *key, nil
}

func signupLimitKey(userID int64) string { return fmt.Sprintf("signup:%d", userID) }

// markRaid tells the post marker the raid changed; a no-op without one.
func (s *Service) markRaid(ctx context.Context, raidID int64) {
	if s.RaidPosts != nil {
		s.RaidPosts.MarkRaidDirty(ctx, raidID)
	}
}

// markRaidMoved tells the post the raid's start moved from was.
func (s *Service) markRaidMoved(ctx context.Context, raidID int64, was time.Time) {
	if s.RaidPosts != nil {
		s.RaidPosts.MarkRaidMoved(ctx, raidID, was)
	}
}
