// api/internal/guilds/roster_account.go
package guilds

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// accountKeyPrefix is accountKeyFor's prefix; parseAccountKey is its inverse.
const accountKeyPrefix = "u:"

// parseAccountKey reads the user id back out of an account key ("u:42"), the handle the roster
// gives an invited member, who has no character key of the guild's own to remove them by.
func parseAccountKey(key string) (int64, bool) {
	raw, ok := strings.CutPrefix(key, accountKeyPrefix)
	if !ok {
		return 0, false
	}
	id, err := strconv.ParseInt(raw, 10, 64)
	if err != nil || id <= 0 {
		return 0, false
	}
	return id, true
}

// rankOrderSQL orders guild_characters rows of alias by rank, leader first.
func rankOrderSQL(alias string) string {
	return "case " + alias + ".rank when 'leader' then 0 when 'officer' then 1 else 2 end"
}

// AccountRank is the highest rank among userID's rows in guildID (leader, then officer, then
// member), the rank the rank-protects-rank rule reads when an officer removes a whole account.
// ErrNotFound when the account has no row in the guild.
func (s *Store) AccountRank(ctx context.Context, guildID, userID int64) (string, error) {
	var rank string
	err := s.Pool.QueryRow(ctx,
		`select gc.rank from guild_characters gc where gc.guild_id = $1 and gc.user_id = $2
		 order by `+rankOrderSQL("gc")+` limit 1`,
		guildID, userID).Scan(&rank)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", ErrNotFound
	}
	if err != nil {
		return "", fmt.Errorf("guilds: account rank: %w", err)
	}
	return rank, nil
}

// removeMember answers DELETE /v1/guilds/{id}/members/{account}: an officer removes a whole
// account from the guild, the way Leave removes the caller's own. It exists for the member who
// joined by invite link: their row's key is the synthetic account: one, which the per-character
// remove route cannot name. Standing and the contested-claim freeze are exactly
// removeCharacter's, read against the account's highest rank.
func (s *Service) removeMember(w http.ResponseWriter, r *http.Request) {
	const op, failMessage = "remove_member", "could not remove that member just now"
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	target, ok := parseAccountKey(r.PathValue("account"))
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such member", nil)
		return
	}
	actor := auth.ActorFrom(r.Context())
	targetRank, err := s.Store.AccountRank(r.Context(), guildID, target)
	if errors.Is(err, ErrNotFound) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such member of this guild", nil)
		return
	}
	if err != nil {
		s.fail(w, r, op, err, failMessage)
		return
	}
	selfOrModerator := target == actor.UserID || actor.IsModerator()
	var verifiedOfficer bool
	if !selfOrModerator {
		if verifiedOfficer, err = s.verifiedOfficerOrLeader(r, guildID); err != nil {
			s.fail(w, r, op, err, failMessage)
			return
		}
	}
	allowed, err := s.Store.mayRemoveCharacter(r.Context(), guildID, actor.UserID, target, targetRank, actor.IsModerator(), verifiedOfficer)
	if err != nil {
		s.fail(w, r, op, err, failMessage)
		return
	}
	if !allowed {
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden", "you do not have standing to remove this member", nil)
		return
	}
	if !selfOrModerator && !s.officerNotFrozen(w, r, guildID, op, failMessage) {
		return
	}
	if err := s.Store.Leave(r.Context(), guildID, target); err != nil {
		if errors.Is(err, ErrNotFound) {
			httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such member of this guild", nil)
			return
		}
		s.fail(w, r, op, err, failMessage)
		return
	}
	s.logger().Info("guilds", "op", op, "guild_id", guildID, "target_user_id", target, "user_id", actor.UserID)
	httpx.WriteOK(w, r, http.StatusOK, map[string]string{"status": "removed"})
}
