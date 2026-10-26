package guilds

import (
	"net/http"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
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
