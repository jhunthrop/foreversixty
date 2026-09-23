// api/internal/guilds/settings.go
package guilds

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// ErrInvalidSettings is returned for a bad default_visibility or
// officer_max_rank_index in a settings PATCH.
var ErrInvalidSettings = errors.New("guilds: invalid settings")

// minOfficerRank and maxOfficerRank bound officer_max_rank_index: a WoW
// guild has ranks 0-9, and 0 is always the guild master (never
// configurable), so an officer threshold sits somewhere in 1-9.
const (
	minOfficerRank = 1
	maxOfficerRank = 9
)

// validDefaultVisibility is the contract's list for a guild's own
// default: never "private" for a guild default.
var validDefaultVisibility = map[string]bool{"public": true, "unlisted": true, "guild": true}

type SettingsPatch struct {
	DefaultVisibility   *string
	OfficerMaxRankIndex *int
}

type MemberRef struct {
	Battletag string `json:"battletag"`
}

type ClaimPendingView struct {
	By        MemberRef `json:"by"`
	ExpiresAt time.Time `json:"expires_at"`
}

type InviteView struct {
	RotatedAt *time.Time `json:"rotated_at"`
}

type SettingsView struct {
	DefaultVisibility   string            `json:"default_visibility"`
	OfficerMaxRankIndex int               `json:"officer_max_rank_index"`
	ClaimedBy           *MemberRef        `json:"claimed_by"`
	ClaimPending        *ClaimPendingView `json:"claim_pending"`
	Claim               ClaimStateView    `json:"claim"`
	Invite              InviteView        `json:"invite"`
}

func (s *Store) Settings(ctx context.Context, guildID int64) (SettingsView, error) {
	g, err := s.getGuild(ctx, guildID)
	if err != nil {
		return SettingsView{}, err
	}
	return SettingsView{
		DefaultVisibility: g.DefaultVisibility, OfficerMaxRankIndex: g.OfficerMaxRankIndex,
		Invite: InviteView{RotatedAt: g.InviteTokenRotatedAt},
		Claim:  claimState(g, time.Now()),
	}, nil
}

func (s *Store) UpdateSettings(ctx context.Context, guildID int64, in SettingsPatch) (SettingsView, error) {
	if in.DefaultVisibility != nil && !validDefaultVisibility[*in.DefaultVisibility] {
		return SettingsView{}, ErrInvalidSettings
	}
	if in.OfficerMaxRankIndex != nil &&
		(*in.OfficerMaxRankIndex < minOfficerRank || *in.OfficerMaxRankIndex > maxOfficerRank) {
		return SettingsView{}, ErrInvalidSettings
	}
	if _, err := s.getGuild(ctx, guildID); err != nil {
		return SettingsView{}, err
	}

	if in.DefaultVisibility != nil {
		if _, err := s.Pool.Exec(ctx, `update guilds set default_visibility = $2 where id = $1`,
			guildID, *in.DefaultVisibility); err != nil {
			return SettingsView{}, fmt.Errorf("guilds: update settings: %w", err)
		}
	}
	if in.OfficerMaxRankIndex != nil {
		tx, err := s.Pool.Begin(ctx)
		if err != nil {
			return SettingsView{}, fmt.Errorf("guilds: update settings: begin: %w", err)
		}
		defer func() { _ = tx.Rollback(ctx) }()
		if _, err := tx.Exec(ctx, `update guilds set officer_max_rank_index = $2 where id = $1`,
			guildID, *in.OfficerMaxRankIndex); err != nil {
			return SettingsView{}, fmt.Errorf("guilds: update settings: %w", err)
		}
		if _, err := tx.Exec(ctx, `
			update guild_characters set rank = case
			    when rank_index = 0 then 'leader'
			    when rank_index <= $2 then 'officer'
			    else 'member'
			  end
			where guild_id = $1 and rank_index is not null`, guildID, *in.OfficerMaxRankIndex); err != nil {
			return SettingsView{}, fmt.Errorf("guilds: re-derive ranks: %w", err)
		}
		if err := RecomputeMembership(ctx, tx, guildID, nil); err != nil {
			return SettingsView{}, err
		}
		if err := tx.Commit(ctx); err != nil {
			return SettingsView{}, fmt.Errorf("guilds: update settings: commit: %w", err)
		}
	}
	return s.Settings(ctx, guildID)
}

// verifiedOfficerOrModerator is the standing check GET/PATCH settings and
// the invite-rotate route share: a verified officer/leader of guildID, or
// a moderator.
func (s *Service) verifiedOfficerOrModerator(r *http.Request, guildID int64) (bool, error) {
	actor := auth.ActorFrom(r.Context())
	if actor.IsModerator() {
		return true, nil
	}
	rank, ok, err := s.Accounts.GuildRank(r.Context(), guildID, actor.UserID)
	if err != nil || !ok {
		return false, err
	}
	return rank == "officer" || rank == "leader", nil
}

func (s *Service) getSettings(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	allowed, err := s.verifiedOfficerOrModerator(r, guildID)
	if err != nil {
		s.fail(w, r, "settings", err, "could not read those settings just now")
		return
	}
	if !allowed {
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden",
			"you must be a verified officer of this guild to see its settings", nil)
		return
	}
	view, err := s.Store.Settings(r.Context(), guildID)
	if errors.Is(err, ErrNotFound) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	if err != nil {
		s.fail(w, r, "settings", err, "could not read those settings just now")
		return
	}
	g, err := s.Store.getGuild(r.Context(), guildID)
	if err != nil {
		s.fail(w, r, "settings", err, "could not read those settings just now")
		return
	}
	if g.ClaimedBy != nil {
		if u, err := s.Accounts.User(r.Context(), *g.ClaimedBy); err == nil {
			view.ClaimedBy = &MemberRef{Battletag: u.PublicName()}
		}
	}
	if g.pendingActive(time.Now()) {
		if u, err := s.Accounts.User(r.Context(), *g.ClaimPendingBy); err == nil {
			view.ClaimPending = &ClaimPendingView{
				By: MemberRef{Battletag: u.PublicName()}, ExpiresAt: g.ClaimRequestedAt.Add(ClaimPendingTTL),
			}
		}
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, view)
}

type settingsPatchInput struct {
	DefaultVisibility   *string `json:"default_visibility"`
	OfficerMaxRankIndex *int    `json:"officer_max_rank_index"`
}

func (s *Service) patchSettings(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	actor := auth.ActorFrom(r.Context())
	allowed, err := s.verifiedOfficerOrModerator(r, guildID)
	if err != nil {
		s.fail(w, r, "patch_settings", err, "could not change those settings just now")
		return
	}
	if !allowed {
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden",
			"you must be a verified officer of this guild to change its settings", nil)
		return
	}
	if frozen, err := s.freezeCheck(r.Context(), guildID, actor.IsModerator()); err != nil {
		s.fail(w, r, "patch_settings", err, "could not change those settings just now")
		return
	} else if frozen {
		httpx.WriteError(w, r, http.StatusConflict, "claim_contested",
			"this guild's claim is contested; officer actions are frozen until a moderator resolves it", nil)
		return
	}
	var in settingsPatchInput
	if err := decodeJSON(w, r, &in); err != nil {
		return
	}
	view, err := s.Store.UpdateSettings(r.Context(), guildID, SettingsPatch{
		DefaultVisibility: in.DefaultVisibility, OfficerMaxRankIndex: in.OfficerMaxRankIndex,
	})
	switch {
	case errors.Is(err, ErrInvalidSettings):
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid",
			"that is not a valid setting", map[string]string{
				"default_visibility":     "one of public, unlisted, guild",
				"officer_max_rank_index": fmt.Sprintf("%d to %d", minOfficerRank, maxOfficerRank),
			})
	case errors.Is(err, ErrNotFound):
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
	case err != nil:
		s.fail(w, r, "patch_settings", err, "could not change those settings just now")
	default:
		s.logger().Info("guilds", "op", "settings_update", "guild_id", guildID, "user_id", actor.UserID)
		httpx.WriteOK(w, r, http.StatusOK, view)
	}
}
