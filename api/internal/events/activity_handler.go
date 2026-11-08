package events

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// Activity page sizes.
const (
	DefaultActivityLimit = 30
	MaxActivityLimit     = 50
	// maxReadBody bounds a mark-read body.
	maxReadBody = 1 << 10
)

// ActivityStore is what the Activity routes read and write (Store).
type ActivityStore interface {
	Activity(ctx context.Context, userID int64, q ActivityQuery) ([]ActivityRow, *int64, error)
	UnreadActivity(ctx context.Context, userID int64) (int, error)
	MarkActivityRead(ctx context.Context, userID, through int64) (int, error)
}

// ActivityPage is the GET /v1/me/activity body. Unread is the account's whole unread count.
type ActivityPage struct {
	Rows       []ActivityRow `json:"rows"`
	Unread     int           `json:"unread"`
	NextBefore *int64        `json:"next_before"`
}

// ActivityUnread is the mark-read answer.
type ActivityUnread struct {
	Unread int `json:"unread"`
}

func mountActivity(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("GET /v1/me/activity", auth.Require(h.activity))
	mux.HandleFunc("POST /v1/me/activity/read", auth.Require(h.markRead))
}

func (h *Handler) activity(w http.ResponseWriter, r *http.Request) {
	q, fields := activityQuery(r)
	if fields != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "some parameters are not valid", fields)
		return
	}
	ctx, uid := r.Context(), auth.ActorFrom(r.Context()).UserID
	rows, next, err := h.Activity.Activity(ctx, uid, q)
	if err != nil {
		h.internal(w, r, "activity_get", err)
		return
	}
	unread, err := h.Activity.UnreadActivity(ctx, uid)
	if err != nil {
		h.internal(w, r, "activity_unread", err)
		return
	}
	if rows == nil {
		rows = []ActivityRow{}
	}
	httpx.CachePrivateNoStore(w)
	httpx.WriteOK(w, r, http.StatusOK, ActivityPage{Rows: rows, Unread: unread, NextBefore: next})
}

// activityQuery reads group, before and limit; fields names what is wrong.
func activityQuery(r *http.Request) (ActivityQuery, map[string]string) {
	v := r.URL.Query()
	q := ActivityQuery{Group: Group(v.Get("group")), Limit: DefaultActivityLimit}
	fields := map[string]string{}
	if q.Group != "" && !KnownGroup(q.Group) {
		fields["group"] = "raids, gear, logs, ladder, tiers, officer or account"
	}
	if s := v.Get("before"); s != "" {
		n, err := strconv.ParseInt(s, 10, 64)
		if err != nil || n < 1 {
			fields["before"] = "a row id"
		}
		q.Before = n
	}
	if s := v.Get("limit"); s != "" {
		n, err := strconv.Atoi(s)
		if err != nil || n < 1 || n > MaxActivityLimit {
			fields["limit"] = "1 to 50"
		}
		q.Limit = n
	}
	if len(fields) > 0 {
		return ActivityQuery{}, fields
	}
	return q, nil
}

func (h *Handler) markRead(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Through *int64 `json:"through"`
	}
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxReadBody)).Decode(&req); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON", nil)
		return
	}
	if req.Through == nil || *req.Through < 1 {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "through is a row id",
			map[string]string{"through": "a row id"})
		return
	}
	unread, err := h.Activity.MarkActivityRead(r.Context(), auth.ActorFrom(r.Context()).UserID, *req.Through)
	if err != nil {
		h.internal(w, r, "activity_read", err)
		return
	}
	httpx.CachePrivateNoStore(w)
	httpx.WriteOK(w, r, http.StatusOK, ActivityUnread{Unread: unread})
}

func (h *Handler) internal(w http.ResponseWriter, r *http.Request, op string, err error) {
	h.logger().Error("events", "op", op, "err", err)
	httpx.WriteError(w, r, http.StatusInternalServerError, "internal", "could not read your activity", nil)
}
