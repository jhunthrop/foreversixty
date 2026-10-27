package guilds

import (
	"context"
	"net/http"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/refusal"
)

// officerNotFrozen answers 409 claim_contested and returns false while the guild's claim is
// contested (freezeCheck); true when officer powers may run. A failed check is a 500 with
// failMessage.
func (s *Service) officerNotFrozen(w http.ResponseWriter, r *http.Request, guildID int64, op, failMessage string) bool {
	frozen, err := s.freezeCheck(r.Context(), guildID, auth.ActorFrom(r.Context()).IsModerator())
	if err != nil {
		s.fail(w, r, op, err, failMessage)
		return false
	}
	if frozen {
		httpx.WriteError(w, r, http.StatusConflict, "claim_contested", refusal.ClaimContested, nil)
		return false
	}
	return true
}

// raidGuildTools reports whether the Raid guild tools (loot council, readiness board) are
// unlocked for guildID. Like officerTools, a failed ask never takes the home down: it reads
// false and logs at warn.
func (s *Service) raidGuildTools(ctx context.Context, guildID, userID int64) bool {
	if s.Entitlements == nil {
		return false
	}
	ok, _, err := s.Entitlements.CanForGuild(ctx, guildID, userID, entitlements.FeatureRaidGuildTools)
	if err != nil {
		s.logger().Warn("guilds", "id", httpx.RequestIDFrom(ctx), "op", "raid_guild_tools", "guild_id", guildID, "err", err)
		return false
	}
	return ok
}
