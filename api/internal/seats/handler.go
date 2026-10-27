package seats

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/refusal"
)

const (
	// seatRateLimitPerHour is every seat route's per-IP budget.
	seatRateLimitPerHour = 60
	maxSeatBody          = 1 << 10
)

// Mount registers the seat routes.
func Mount(mux *http.ServeMux, s *Service, trustedProxyHops int) {
	limit := httpx.RateLimitPer(seatRateLimitPerHour, time.Hour, trustedProxyHops)
	route := func(pattern string, h http.HandlerFunc) {
		mux.Handle(pattern, limit(auth.RequireSession(h)))
	}
	const base = "/v1/guilds/{id}/seats"
	route("GET "+base+"/candidates", s.candidates)
	route("POST "+base+"/{slot}/give", s.recipientAction(s.Give))
	route("POST "+base+"/{slot}/move", s.recipientAction(s.Move))
	route("POST "+base+"/{slot}/withdraw", s.withdraw)
	route("POST "+base+"/offer/accept", s.offerAction(s.Accept))
	route("POST "+base+"/offer/decline", s.offerAction(s.Decline))
}

func pathInt(r *http.Request, name string) (int64, bool) {
	v, err := strconv.ParseInt(r.PathValue(name), 10, 64)
	return v, err == nil && v > 0
}

func (s *Service) candidates(w http.ResponseWriter, r *http.Request) {
	guildID, ok := pathInt(r, "id")
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	actor := auth.ActorFrom(r.Context()).UserID
	if _, err := s.authorize(r.Context(), guildID, actor); err != nil {
		s.refuse(w, r, "candidates", err)
		return
	}
	list, err := s.Store.Candidates(r.Context(), guildID)
	if err != nil {
		s.refuse(w, r, "candidates", err)
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, map[string]any{"candidates": list})
}

func (s *Service) recipientAction(act func(context.Context, int64, int, int64, int64) (Seat, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		guildID, slot, ok := guildAndSlot(r)
		if !ok {
			httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such seat", nil)
			return
		}
		var in struct {
			UserID int64 `json:"user_id"`
		}
		if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxSeatBody)).Decode(&in); err != nil || in.UserID <= 0 {
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON with a user_id", nil)
			return
		}
		seat, err := act(r.Context(), guildID, slot, auth.ActorFrom(r.Context()).UserID, in.UserID)
		if err != nil {
			s.refuse(w, r, "seat action", err)
			return
		}
		httpx.WriteOK(w, r, http.StatusOK, seat)
	}
}

func (s *Service) withdraw(w http.ResponseWriter, r *http.Request) {
	guildID, slot, ok := guildAndSlot(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such seat", nil)
		return
	}
	seat, err := s.Withdraw(r.Context(), guildID, slot, auth.ActorFrom(r.Context()).UserID)
	if err != nil {
		s.refuse(w, r, "withdraw", err)
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, seat)
}

func (s *Service) offerAction(act func(context.Context, int64, int64) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		guildID, ok := pathInt(r, "id")
		if !ok {
			httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
			return
		}
		if err := act(r.Context(), guildID, auth.ActorFrom(r.Context()).UserID); err != nil {
			s.refuse(w, r, "offer", err)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func guildAndSlot(r *http.Request) (guildID int64, slot int, ok bool) {
	guildID, ok = pathInt(r, "id")
	if !ok {
		return 0, 0, false
	}
	slot64, ok := pathInt(r, "slot")
	return guildID, int(slot64), ok
}

// refuse maps a service error to its HTTP answer.
func (s *Service) refuse(w http.ResponseWriter, r *http.Request, op string, err error) {
	var locked ErrLocked
	switch {
	case errors.Is(err, ErrNotPayer):
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden", "only the guild's payer can manage seats", nil)
	case errors.Is(err, ErrNoPlan):
		httpx.WriteError(w, r, http.StatusConflict, "no_plan", "this guild's plan has no seats", nil)
	case errors.Is(err, ErrNoSuchSlot):
		httpx.WriteError(w, r, http.StatusConflict, "seat_count", "that seat is outside the plan's seat count", nil)
	case errors.Is(err, ErrOfferPending):
		httpx.WriteError(w, r, http.StatusConflict, "offer_pending", "an offer is waiting on that seat; withdraw it first", nil)
	case errors.Is(err, ErrNotMember):
		httpx.WriteError(w, r, http.StatusConflict, "not_member", "that account is not a verified member of this guild", nil)
	case errors.As(err, &locked):
		httpx.WriteError(w, r, http.StatusConflict, "seat_locked", "this seat moved recently",
			map[string]string{"movable_at": locked.MovableAt.UTC().Format(time.RFC3339)})
	case errors.Is(err, ErrTaken), errors.Is(err, ErrNotHeld):
		httpx.WriteError(w, r, http.StatusConflict, "conflict", "that seat or account is already taken", nil)
	case errors.Is(err, ErrNoOffer):
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "there is no offer to answer", nil)
	case errors.Is(err, ErrClaimContested):
		httpx.WriteError(w, r, http.StatusConflict, "claim_contested", refusal.ClaimContested, nil)
	case errors.Is(err, ErrBillingUnavailable):
		httpx.WriteError(w, r, http.StatusServiceUnavailable, "billing_unavailable",
			"billing is not configured on this deployment", nil)
	default:
		s.logger().Error("seats", "id", httpx.RequestIDFrom(r.Context()), "op", op, "err", err)
		httpx.WriteError(w, r, http.StatusInternalServerError, "internal", "something went wrong", nil)
	}
}
