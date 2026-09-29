// api/internal/auth/ratelimit.go
package auth

import (
	"net/http"
	"strconv"
)

// AccountRateLimitKey keys an httpx.RateLimitByKey limiter by the
// request's signed-in account id, for routes that must budget per member
// rather than per IP address (night-api-security finding, LOW:
// POST /v1/sims/run and POST /v1/uploads dispatch real Cloud Run compute
// and R2 storage and had only the router-wide per-IP budget). ok is
// false for an unsigned request, which is then exempt from that
// limiter's budget entirely - every route this key is used on also
// wraps RequireSession, which answers an unsigned request with 401
// before the budget would ever matter, and an unsigned caller must not
// be able to exhaust a shared "no account" bucket for everyone else.
func AccountRateLimitKey(r *http.Request) (key string, ok bool) {
	uid := ActorFrom(r.Context()).UserID
	if uid == 0 {
		return "", false
	}
	return strconv.FormatInt(uid, 10), true
}
