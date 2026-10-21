package auth

import (
	"context"
	"encoding/json"
	"errors"
	"net"
	"net/http"

	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// devLogin signs the caller in as a seeded account, by battletag, with no password: the
// local walkthrough stack (tools/dev-stack) has no Battle.net to sign in through, and the
// guild workflows need a browser session for the guild master, an officer and a member.
//
// This route must never ship enabled. It is mounted only when Service.DevLogin is set,
// which main sets only from FS_DEV_LOGIN=1, and config.Load refuses that flag unless the
// listener is bound to a loopback address and PUBLIC_BASE_URL is not production. This
// handler repeats the loopback test on the connecting peer so that a listener bound
// wider than intended, or a proxy in front of it, still cannot reach the route from off
// the machine. It signs in an account that already exists and never creates one.
func (s *Service) devLogin(w http.ResponseWriter, r *http.Request) {
	httpx.CachePrivate(w)
	if !isLoopbackPeer(r.RemoteAddr) {
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden", "the dev sign-in answers loopback callers only", nil)
		return
	}
	var in struct {
		Battletag string `json:"battletag"`
	}
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxAuthBody)).Decode(&in); err != nil || in.Battletag == "" {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON with a battletag", nil)
		return
	}
	u, err := s.Store.UserByBattletag(r.Context(), in.Battletag)
	if errors.Is(err, ErrNotFound) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no seeded account has that battletag", nil)
		return
	}
	if err != nil {
		s.fail(w, r, "devlogin", err, "could not sign you in just now")
		return
	}
	if _, err := s.Auth.StartSession(r.Context(), w, u.ID, "dev"); err != nil {
		s.fail(w, r, "devlogin", err, "could not sign you in just now")
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, map[string]any{"user_id": u.ID, "battletag": u.Battletag})
}

// isLoopbackPeer reports whether a request's RemoteAddr is a loopback IP.
func isLoopbackPeer(remoteAddr string) bool {
	host, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		return false
	}
	ip := net.ParseIP(host)
	return ip != nil && ip.IsLoopback()
}

// UserByBattletag reads the oldest account holding a battletag, for the dev sign-in.
func (s *Store) UserByBattletag(ctx context.Context, battletag string) (User, error) {
	return scanUser(s.Pool.QueryRow(ctx,
		`select `+userColumns+` from users where battletag = $1 order by id limit 1`, battletag))
}
