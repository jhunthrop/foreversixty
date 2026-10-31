package livesurface

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// RegisterRequest is POST /v1/me/live-activities (W3-C13).
type RegisterRequest struct {
	Kind             string `json:"kind"`
	Platform         string `json:"platform"`
	PushToken        string `json:"push_token"`
	PushToStartToken string `json:"push_to_start_token"`
	ObjectID         string `json:"object_id"`
}

// Registered is the answer: the id DELETE takes.
type Registered struct {
	ID int64 `json:"id"`
}

const (
	maxBody = 8 << 10
	// raidTail keeps a raid countdown alive past the raid's end, as spec §5.6 ends it at the
	// pull plus 15 minutes; a raid's window covers that, the tail is the cancelled note's 5
	// minutes and a little slack.
	raidTail = 15 * time.Minute
	// simLife is how long a sim registration lives: a run that has not finished by then is
	// failed by the runner, and a finished one stays an hour.
	simLife = 6 * time.Hour
)

var (
	// An APNs token (ActivityKit's update and push-to-start tokens too) is hex; an FCM
	// registration token is base64url with colons. The bounds are the push-token route's.
	iosToken     = regexp.MustCompile(`^[0-9a-fA-F]{64,200}$`)
	androidToken = regexp.MustCompile(`^[A-Za-z0-9_:-]+$`)
	raidObject   = regexp.MustCompile(`^[1-9][0-9]{0,18}$`)
	simObject    = regexp.MustCompile(`^[A-Za-z0-9_-]{1,64}$`)
	errNotFound  = errors.New("livesurface: no such object for this account")
)

// register answers POST /v1/me/live-activities.
func (s *Service) register(w http.ResponseWriter, r *http.Request) {
	actor := auth.ActorFrom(r.Context())
	if !s.RegisterLimiter.Allow("live:" + strconv.FormatInt(actor.UserID, 10)) {
		httpx.WriteError(w, r, http.StatusTooManyRequests, "rate_limited", "too many requests", nil)
		return
	}
	var req RegisterRequest
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxBody)).Decode(&req); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "send a JSON body", nil)
		return
	}
	if fields := validate(req); len(fields) > 0 {
		httpx.WriteError(w, r, http.StatusUnprocessableEntity, "invalid", "check the fields", fields)
		return
	}
	if req.Kind == KindSim && !s.premium(w, r, actor.UserID) {
		return
	}
	expires, err := s.objectExpiry(r.Context(), actor.UserID, req)
	if errors.Is(err, errNotFound) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such "+req.Kind, nil)
		return
	}
	if err != nil {
		s.fail(w, r, "live_object", err, "could not start that just now")
		return
	}
	id, err := s.Tokens.Register(r.Context(), Registration{
		UserID: actor.UserID, DeviceID: actor.DeviceID, Kind: req.Kind, Platform: req.Platform,
		ObjectID: req.ObjectID, PushToken: req.PushToken, PushToStartToken: req.PushToStartToken,
		ExpiresAt: expires,
	})
	if err != nil {
		s.fail(w, r, "live_register", err, "could not start that just now")
		return
	}
	// The tokens are never logged; the row id and object are enough to follow a report.
	s.logger().Info("livesurface", "op", "live_register", "user_id", actor.UserID, "kind", req.Kind,
		"platform", req.Platform, "object_id", req.ObjectID, "registration", id)
	httpx.WriteOK(w, r, http.StatusCreated, Registered{ID: id})
}

// unregister answers DELETE /v1/me/live-activities/{id}: 204 whether or not it was there.
func (s *Service) unregister(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id <= 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if err := s.Tokens.Unregister(r.Context(), auth.ActorFrom(r.Context()).UserID, id); err != nil {
		s.fail(w, r, "live_unregister", err, "could not end that just now")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// premium reports whether the caller holds server_sims, having written the 402 when not.
func (s *Service) premium(w http.ResponseWriter, r *http.Request, userID int64) bool {
	ok, _, err := s.Entitlements.Can(r.Context(), userID, entitlements.FeatureServerSims)
	if err != nil {
		s.fail(w, r, "live_premium", err, "could not check your account just now")
		return false
	}
	if !ok {
		httpx.WriteError(w, r, http.StatusPaymentRequired, "premium_required",
			"server sims are part of Premium", nil)
	}
	return ok
}

// objectExpiry checks the object is the caller's to follow and answers when its registration
// ends; errNotFound when it is not.
func (s *Service) objectExpiry(ctx context.Context, userID int64, req RegisterRequest) (time.Time, error) {
	if req.Kind == KindSim {
		owns, err := s.Sims.OwnsServerRun(ctx, userID, req.ObjectID)
		if err != nil {
			return time.Time{}, err
		}
		if !owns {
			return time.Time{}, errNotFound
		}
		return s.now().Add(simLife), nil
	}
	raidID, err := strconv.ParseInt(req.ObjectID, 10, 64)
	if err != nil {
		return time.Time{}, fmt.Errorf("livesurface: raid id: %w", err)
	}
	end, found, err := s.Raids.LiveRaidEnd(ctx, userID, raidID)
	if err != nil {
		return time.Time{}, err
	}
	if !found || !end.After(s.now()) {
		return time.Time{}, errNotFound
	}
	return end.Add(raidTail), nil
}

// validate names each bad field of req.
func validate(req RegisterRequest) map[string]string {
	fields := map[string]string{}
	switch req.Kind {
	case KindRaid:
		if !raidObject.MatchString(req.ObjectID) {
			fields["object_id"] = "a raid id"
		}
	case KindSim:
		if !simObject.MatchString(req.ObjectID) {
			fields["object_id"] = "a sim id"
		}
	default:
		fields["kind"] = "raid or sim"
	}
	token := tokenRule(req.Platform)
	if token == nil {
		fields["platform"] = "ios or android"
		return fields
	}
	if !validToken(req.Platform, token, req.PushToken) {
		fields["push_token"] = "the push token the system gave the app"
	}
	if req.PushToStartToken != "" && (req.Platform != PlatformIOS || !validToken(req.Platform, token, req.PushToStartToken)) {
		fields["push_to_start_token"] = "an iOS push-to-start token, or nothing"
	}
	return fields
}

func tokenRule(platform string) *regexp.Regexp {
	switch platform {
	case PlatformIOS:
		return iosToken
	case PlatformAndroid:
		return androidToken
	default:
		return nil
	}
}

// validToken applies the platform's pattern and, for an FCM token, the push-token route's 16 to
// 4096 characters (Go's regexp caps a repeat at 1000).
func validToken(platform string, rule *regexp.Regexp, token string) bool {
	if platform == PlatformAndroid && (len(token) < 16 || len(token) > 4096) {
		return false
	}
	return rule.MatchString(token)
}
