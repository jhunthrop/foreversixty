package creators

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

const maxBody = 4 << 10

// applicationsPerHour is POST /v1/creators/application's budget, per address and per account.
const applicationsPerHour = 5

// Mount registers the program's routes: the owner's under /v1/admin/creators, the signed-in
// account's under /v1/creators.
func Mount(mux *http.ServeMux, s *Service, trustedProxyHops int) {
	admin := func(pattern string, h http.HandlerFunc) { mux.HandleFunc(pattern, auth.RequireAdmin(h)) }
	admin("GET /v1/admin/creators", s.list)
	admin("POST /v1/admin/creators", s.add)
	admin("PATCH /v1/admin/creators/{id}", s.edit)
	admin("DELETE /v1/admin/creators/{id}", s.remove)
	admin("POST /v1/admin/creators/{id}/invite", s.invite)
	admin("POST /v1/admin/creators/{id}/stripe-status", s.refresh)
	admin("GET /v1/admin/creators/statements", s.statements)
	admin("POST /v1/admin/creators/{id}/statements/{month}/approve", s.approve)
	admin("GET /v1/admin/creators/balance", s.balance)
	admin("GET /v1/admin/creators/requests", s.requests)
	admin("POST /v1/admin/creators/requests/{id}/approve", s.approveRequest)
	admin("POST /v1/admin/creators/requests/{id}/decline", s.declineRequest)

	session := func(pattern string, h http.HandlerFunc) { mux.HandleFunc(pattern, auth.RequireSession(h)) }
	session("POST /v1/creators/invite/preview", s.invitePreview)
	session("POST /v1/creators/invite", s.acceptInvite)
	session("GET /v1/creators/me", s.me)
	session("POST /v1/creators/me/onboarding", s.onboarding)
	session("POST /v1/creators/me/dashboard", s.dashboard)
	session("GET /v1/creators/application", s.application)
	// Two budgets: per address before the session check, and per account behind it, so a shared
	// address (a proxy hop, a campus) and a rotating one are both held to five an hour.
	perAddress := httpx.RateLimitPer(applicationsPerHour, time.Hour, trustedProxyHops)
	perAccount := httpx.RateLimitByKey(applicationsPerHour, time.Hour, func(r *http.Request) (string, bool) {
		id := auth.ActorFrom(r.Context()).UserID
		return strconv.FormatInt(id, 10), id != 0
	})
	mux.Handle("POST /v1/creators/application", perAddress(auth.RequireSession(perAccount(http.HandlerFunc(s.apply)).ServeHTTP)))
}

// fail answers an error: the program's own errors as what they mean, anything else as a 500 with
// the error logged (ids only; nothing a person typed reaches a log).
func (s *Service) fail(w http.ResponseWriter, r *http.Request, op string, err error) {
	var fe *FieldError
	switch {
	case errors.As(err, &fe):
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "check the fields", fe.Fields)
	case errors.Is(err, ErrNotFound):
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such creator", nil)
	case errors.Is(err, ErrNotCreator):
		httpx.WriteError(w, r, http.StatusForbidden, "not_creator", "this account is not in the creator program", nil)
	case errors.Is(err, ErrBadInvite):
		httpx.WriteError(w, r, http.StatusGone, "invite_used", "that invite has expired or was already used", nil)
	case errors.Is(err, ErrAlreadyCreator):
		httpx.WriteError(w, r, http.StatusConflict, "already_creator", "that account is already in the creator program", nil)
	case errors.Is(err, ErrTaken):
		httpx.WriteError(w, r, http.StatusConflict, "taken", "that account or link is already a creator's", nil)
	case errors.Is(err, ErrReapplyTooSoon):
		httpx.WriteError(w, r, http.StatusConflict, "reapply_too_soon", "you can apply again 30 days after a decline", nil)
	case errors.Is(err, ErrApplicationOpen):
		httpx.WriteError(w, r, http.StatusConflict, "application_open", "an application is already open", nil)
	case errors.Is(err, ErrNotPending):
		httpx.WriteError(w, r, http.StatusConflict, "decided", "that request was already decided", nil)
	case errors.Is(err, ErrMonthOpen):
		httpx.WriteError(w, r, http.StatusConflict, "month_open", "the month is still open", nil)
	case errors.Is(err, ErrBelowMinimum):
		httpx.WriteError(w, r, http.StatusConflict, "below_minimum", "under the $10 minimum; it carries over to next month", nil)
	case errors.Is(err, ErrNothingOwed):
		httpx.WriteError(w, r, http.StatusConflict, "nothing_owed", "nothing is owed this month", nil)
	case errors.Is(err, ErrNotPayable):
		httpx.WriteError(w, r, http.StatusConflict, "not_payable", "stripe cannot pay this creator yet", nil)
	case errors.Is(err, ErrInsufficientBalance):
		httpx.WriteError(w, r, http.StatusConflict, "balance_insufficient",
			"Not enough in your Stripe balance for this payout. Nothing was paid; approve it again once the balance covers it.", nil)
	case errors.Is(err, ErrConnectDenied):
		httpx.WriteError(w, r, http.StatusServiceUnavailable, "connect_denied", "the stripe payouts key is missing or lacks connect permissions", nil)
	case errors.Is(err, ErrPayoutsOff):
		httpx.WriteError(w, r, http.StatusServiceUnavailable, "payouts_off", "payouts are switched off on the server", nil)
	case errors.Is(err, ErrPayoutsKeyMissing):
		httpx.WriteError(w, r, http.StatusServiceUnavailable, "payouts_key_missing", "the stripe payouts key is not configured", nil)
	case errors.Is(err, ErrStripeOff):
		httpx.WriteError(w, r, http.StatusServiceUnavailable, "stripe_off", "stripe is not configured", nil)
	default:
		s.logger().Error("creators", "id", httpx.RequestIDFrom(r.Context()), "op", op, "err", err)
		httpx.WriteError(w, r, http.StatusInternalServerError, "internal", "that didn't work just now", nil)
	}
}

// decode reads one JSON object and nothing else: unknown fields and trailing data are refused.
func decode(w http.ResponseWriter, r *http.Request, v any) bool {
	dec := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxBody))
	dec.DisallowUnknownFields()
	if err := dec.Decode(v); err != nil || dec.More() {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "the body must be JSON", nil)
		return false
	}
	return true
}

func pathID(w http.ResponseWriter, r *http.Request) (int64, bool) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id <= 0 {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such creator", nil)
		return 0, false
	}
	return id, true
}

func ok(w http.ResponseWriter, r *http.Request, status int, data any) {
	httpx.CachePrivateNoStore(w)
	httpx.WriteOK(w, r, status, data)
}

func (s *Service) list(w http.ResponseWriter, r *http.Request) {
	list, err := s.List(r.Context())
	if err != nil {
		s.fail(w, r, "list", err)
		return
	}
	ok(w, r, http.StatusOK, map[string]any{"creators": list, "stripe_enabled": s.Gateway != nil, "payouts_enabled": s.PayoutsEnabled,
		"connect_denied": s.ConnectDenied()})
}

func (s *Service) add(w http.ResponseWriter, r *http.Request) {
	var in Input
	if !decode(w, r, &in) {
		return
	}
	c, err := s.Add(r.Context(), AddInput{Input: in})
	if err != nil {
		s.fail(w, r, "add", err)
		return
	}
	ok(w, r, http.StatusCreated, c)
}

func (s *Service) edit(w http.ResponseWriter, r *http.Request) {
	id, found := pathID(w, r)
	var in Input
	if !found || !decode(w, r, &in) {
		return
	}
	c, err := s.Edit(r.Context(), id, in)
	if err != nil {
		s.fail(w, r, "edit", err)
		return
	}
	ok(w, r, http.StatusOK, c)
}

func (s *Service) remove(w http.ResponseWriter, r *http.Request) {
	id, found := pathID(w, r)
	if !found {
		return
	}
	if err := s.Remove(r.Context(), id); err != nil {
		s.fail(w, r, "remove", err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s *Service) invite(w http.ResponseWriter, r *http.Request) {
	id, found := pathID(w, r)
	if !found {
		return
	}
	url, err := s.Invite(r.Context(), id)
	if err != nil {
		s.fail(w, r, "invite", err)
		return
	}
	ok(w, r, http.StatusOK, map[string]string{"url": url})
}

func (s *Service) refresh(w http.ResponseWriter, r *http.Request) {
	id, found := pathID(w, r)
	if !found {
		return
	}
	c, err := s.Store.Creator(r.Context(), id)
	if err != nil {
		s.fail(w, r, "refresh", err)
		return
	}
	status, err := s.RefreshStatus(r.Context(), c)
	if err != nil {
		s.fail(w, r, "refresh", err)
		return
	}
	ok(w, r, http.StatusOK, map[string]string{"stripe_status": status})
}

// lastClosedMonth is the month the statements page opens on.
func lastClosedMonth(now time.Time) Month { return Month{Start: MonthOf(now).Start.AddDate(0, -1, 0)} }

func (s *Service) statements(w http.ResponseWriter, r *http.Request) {
	m := lastClosedMonth(s.now())
	if q := r.URL.Query().Get("month"); q != "" {
		var err error
		if m, err = ParseMonth(q); err != nil {
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "month must be YYYY-MM", map[string]string{"month": "YYYY-MM"})
			return
		}
	}
	page, err := s.Statements(r.Context(), m)
	if err != nil {
		s.fail(w, r, "statements", err)
		return
	}
	ok(w, r, http.StatusOK, page)
}

func (s *Service) approve(w http.ResponseWriter, r *http.Request) {
	id, found := pathID(w, r)
	if !found {
		return
	}
	m, err := ParseMonth(r.PathValue("month"))
	if err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "month must be YYYY-MM", map[string]string{"month": "YYYY-MM"})
		return
	}
	st, err := s.Approve(r.Context(), id, m)
	if err != nil {
		s.fail(w, r, "approve", err)
		return
	}
	ok(w, r, http.StatusOK, st)
}

func (s *Service) balance(w http.ResponseWriter, r *http.Request) {
	cents, err := s.AvailableUSD(r.Context())
	if err != nil {
		s.fail(w, r, "balance", err)
		return
	}
	ok(w, r, http.StatusOK, map[string]int64{"available_cents": cents})
}

func (s *Service) requests(w http.ResponseWriter, r *http.Request) {
	list, err := s.Store.Requests(r.Context())
	if err != nil {
		s.fail(w, r, "requests", err)
		return
	}
	ok(w, r, http.StatusOK, map[string]any{"requests": list})
}

func (s *Service) approveRequest(w http.ResponseWriter, r *http.Request) {
	id, found := pathID(w, r)
	var in Input
	if !found || !decode(w, r, &in) {
		return
	}
	c, err := s.ApproveRequest(r.Context(), id, in)
	if err != nil {
		s.fail(w, r, "approve_request", err)
		return
	}
	ok(w, r, http.StatusCreated, c)
}

func (s *Service) declineRequest(w http.ResponseWriter, r *http.Request) {
	id, found := pathID(w, r)
	if !found {
		return
	}
	if err := s.DeclineRequest(r.Context(), id); err != nil {
		s.fail(w, r, "decline_request", err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

type tokenBody struct {
	Token string `json:"token"`
}

func (s *Service) invitePreview(w http.ResponseWriter, r *http.Request) {
	var in tokenBody
	if !decode(w, r, &in) {
		return
	}
	name, err := s.InviteName(r.Context(), in.Token)
	if err != nil {
		s.fail(w, r, "invite_preview", err)
		return
	}
	ok(w, r, http.StatusOK, map[string]string{"name": name})
}

func (s *Service) acceptInvite(w http.ResponseWriter, r *http.Request) {
	var in tokenBody
	if !decode(w, r, &in) {
		return
	}
	c, err := s.AcceptInvite(r.Context(), in.Token, auth.ActorFrom(r.Context()).UserID)
	if err != nil {
		s.fail(w, r, "accept_invite", err)
		return
	}
	ok(w, r, http.StatusOK, map[string]string{"name": c.Name})
}

func (s *Service) me(w http.ResponseWriter, r *http.Request) {
	me, err := s.Me(r.Context(), auth.ActorFrom(r.Context()).UserID, r.URL.Query().Get("sync") == "1")
	if err != nil {
		s.fail(w, r, "me", err)
		return
	}
	ok(w, r, http.StatusOK, me)
}

func (s *Service) onboarding(w http.ResponseWriter, r *http.Request) {
	url, err := s.Onboard(r.Context(), auth.ActorFrom(r.Context()).UserID)
	if err != nil {
		s.fail(w, r, "onboarding", err)
		return
	}
	ok(w, r, http.StatusOK, map[string]string{"url": url})
}

func (s *Service) dashboard(w http.ResponseWriter, r *http.Request) {
	url, err := s.Dashboard(r.Context(), auth.ActorFrom(r.Context()).UserID)
	if err != nil {
		s.fail(w, r, "dashboard", err)
		return
	}
	ok(w, r, http.StatusOK, map[string]string{"url": url})
}

func (s *Service) application(w http.ResponseWriter, r *http.Request) {
	a, err := s.Store.Application(r.Context(), auth.ActorFrom(r.Context()).UserID, s.now())
	if err != nil {
		s.fail(w, r, "application", err)
		return
	}
	ok(w, r, http.StatusOK, a)
}

func (s *Service) apply(w http.ResponseWriter, r *http.Request) {
	var in ApplicationInput
	if !decode(w, r, &in) {
		return
	}
	if err := s.Apply(r.Context(), auth.ActorFrom(r.Context()).UserID, in); err != nil {
		s.fail(w, r, "apply", err)
		return
	}
	ok(w, r, http.StatusCreated, map[string]string{"status": ApplicationPending})
}
