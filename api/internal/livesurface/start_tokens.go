package livesurface

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// StartTokenRequest is POST /v1/me/live-activities/start-token.
type StartTokenRequest struct {
	Platform string `json:"platform"`
	Kind     string `json:"kind"`
	Token    string `json:"token"`
}

// StartTokenStored is its answer.
type StartTokenStored struct {
	Kind     string `json:"kind"`
	Platform string `json:"platform"`
}

// StartToken is one install's push-to-start token, as the senders read it.
type StartToken struct {
	DeviceID string
	Platform string
	Token    string
}

// storeStartToken answers POST /v1/me/live-activities/start-token: the app's device keeps one
// push-to-start token per kind, replaced on each call.
func (s *Service) storeStartToken(w http.ResponseWriter, r *http.Request) {
	actor := auth.ActorFrom(r.Context())
	if !s.RegisterLimiter.Allow("live-start:" + strconv.FormatInt(actor.UserID, 10)) {
		httpx.WriteError(w, r, http.StatusTooManyRequests, "rate_limited", "too many requests", nil)
		return
	}
	if actor.DeviceID == "" {
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden", "only the app's device stores a start token", nil)
		return
	}
	var req StartTokenRequest
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxBody)).Decode(&req); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON", nil)
		return
	}
	if fields := validateStartToken(req); fields != nil {
		httpx.WriteError(w, r, http.StatusUnprocessableEntity, "invalid", "that start token is not valid", fields)
		return
	}
	if err := s.Tokens.SaveStartToken(r.Context(), actor.UserID, actor.DeviceID, req); err != nil {
		s.fail(w, r, "live_start_token", err, "could not store the start token just now")
		return
	}
	httpx.CachePrivateNoStore(w)
	httpx.WriteOK(w, r, http.StatusOK, StartTokenStored{Kind: req.Kind, Platform: req.Platform})
}

// validateStartToken names each bad field; nil when the request is good.
func validateStartToken(req StartTokenRequest) map[string]string {
	fields := map[string]string{}
	if req.Kind != KindRaid {
		fields["kind"] = "raid"
	}
	rule := tokenRule(req.Platform)
	if rule == nil {
		fields["platform"] = "ios or android"
	} else if !validToken(req.Platform, rule, req.Token) {
		fields["token"] = "not a push token"
	}
	if len(fields) > 0 {
		return fields
	}
	return nil
}

// SaveStartToken keeps the device's token for the kind, replacing the one before.
func (s *Store) SaveStartToken(ctx context.Context, userID int64, deviceID string, req StartTokenRequest) error {
	if _, err := s.Pool.Exec(ctx, `insert into live_activity_start_tokens (user_id, device_id, kind, platform, token)
		values ($1, $2, $3, $4, $5)
		on conflict (user_id, device_id, kind) do update set platform = excluded.platform, token = excluded.token,
		  updated_at = now()`, userID, deviceID, req.Kind, req.Platform, req.Token); err != nil {
		return fmt.Errorf("livesurface: save start token: %w", err)
	}
	return nil
}

// StartTokens is the account's push-to-start tokens for kind.
func (s *Store) StartTokens(ctx context.Context, userID int64, kind string) ([]StartToken, error) {
	rows, err := s.Pool.Query(ctx, `select device_id, platform, token from live_activity_start_tokens
		where user_id = $1 and kind = $2 order by device_id`, userID, kind)
	if err != nil {
		return nil, fmt.Errorf("livesurface: start tokens: %w", err)
	}
	defer rows.Close()
	var out []StartToken
	for rows.Next() {
		var t StartToken
		if err := rows.Scan(&t.DeviceID, &t.Platform, &t.Token); err != nil {
			return nil, fmt.Errorf("livesurface: start tokens: %w", err)
		}
		out = append(out, t)
	}
	return out, rows.Err()
}

// DropStartToken removes a start token the push service refused.
func (s *Store) DropStartToken(ctx context.Context, userID int64, deviceID, kind string) error {
	if _, err := s.Pool.Exec(ctx, `delete from live_activity_start_tokens where user_id = $1 and device_id = $2 and kind = $3`,
		userID, deviceID, kind); err != nil {
		return fmt.Errorf("livesurface: drop start token: %w", err)
	}
	return nil
}

// Running reports whether the account already keeps a live surface of kind for objectID.
func (s *Store) Running(ctx context.Context, userID int64, kind, objectID string) (bool, error) {
	var running bool
	if err := s.Pool.QueryRow(ctx, `select exists (select 1 from live_activity_tokens
		where user_id = $1 and kind = $2 and object_id = $3 and expires_at > now())`, userID, kind, objectID).Scan(&running); err != nil {
		return false, fmt.Errorf("livesurface: running: %w", err)
	}
	return running, nil
}
