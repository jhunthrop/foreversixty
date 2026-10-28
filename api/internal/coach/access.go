package coach

import (
	"context"
	"fmt"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/rating"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
)

// Access is what a viewer may see of one row.
type Access string

const (
	AccessNone    Access = ""
	AccessTeaser  Access = "teaser"
	AccessFull    Access = "full"
	AccessOfficer Access = "officer"
)

// verifiedMemberSQL: the character's owner is a verified member of the guild.
const verifiedMemberSQL = `
select exists (select 1 from characters c join guild_members gm on gm.user_id = c.user_id
  where c.key = $1 and gm.guild_id = $2 and gm.verified_at is not null)`

// AccessFor decides what actor may see of playerKey's row on rep (§3, §19), and the
// account a generation counts against: the report's viewers only; the owner (full with
// log_coach, else the teaser); an officer with officer_views for the report's guild, on a
// verified member's row; a moderator as an officer, except on an anonymised character.
// Anyone else gets AccessNone.
func (s *Service) AccessFor(ctx context.Context, actor auth.Actor, rep reports.Report, playerKey string) (Access, int64, error) {
	if !actor.Signed() || !reports.MayView(ctx, actor, rep, s.Accounts) {
		return AccessNone, 0, nil
	}
	owns, err := s.Owners.OwnsCharacter(ctx, actor.UserID, playerKey)
	if err != nil {
		return AccessNone, 0, fmt.Errorf("coach: owner check: %w", err)
	}
	if owns {
		return s.ownerAccess(ctx, actor.UserID)
	}
	officer, err := s.officerOf(ctx, actor, rep, playerKey)
	if err != nil || officer {
		return officerOr(officer), actor.UserID, err
	}
	if !actor.IsModerator() {
		return AccessNone, 0, nil
	}
	anon, err := s.anonymized(ctx, playerKey)
	if err != nil || anon {
		return AccessNone, 0, err
	}
	return AccessOfficer, actor.UserID, nil
}

func officerOr(officer bool) Access {
	if officer {
		return AccessOfficer
	}
	return AccessNone
}

func (s *Service) ownerAccess(ctx context.Context, userID int64) (Access, int64, error) {
	can, _, err := s.Entitlements.Can(ctx, userID, entitlements.FeatureLogCoach)
	if err != nil {
		return AccessNone, 0, fmt.Errorf("coach: entitlement: %w", err)
	}
	if can {
		return AccessFull, userID, nil
	}
	return AccessTeaser, userID, nil
}

// officerOf: the report has a guild, actor holds officer_views for it, and the row's
// owner is a verified member of it.
func (s *Service) officerOf(ctx context.Context, actor auth.Actor, rep reports.Report, playerKey string) (bool, error) {
	if rep.GuildID == nil {
		return false, nil
	}
	can, _, err := s.Entitlements.CanForGuild(ctx, *rep.GuildID, actor.UserID, entitlements.FeatureOfficerViews)
	if err != nil {
		return false, fmt.Errorf("coach: officer entitlement: %w", err)
	}
	if !can {
		return false, nil
	}
	var member bool
	if err := s.Store.DB.QueryRow(ctx, verifiedMemberSQL, playerKey, *rep.GuildID).Scan(&member); err != nil {
		return false, fmt.Errorf("coach: verified member: %w", err)
	}
	return member, nil
}

func (s *Service) anonymized(ctx context.Context, playerKey string) (bool, error) {
	anon, err := rating.CharacterAnonymized(ctx, s.Store.DB, playerKey)
	if err != nil {
		return false, fmt.Errorf("coach: %w", err)
	}
	return anon, nil
}
