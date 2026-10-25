package notify

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"time"
	"unicode"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

const (
	testsPerMinute  = 5
	putsPerMinute   = 10
	maxPutBodyBytes = 4 << 10
	// maxTestErrorRunes is the spec's cap on the Discord text a failed test shows.
	maxTestErrorRunes = 140
	// maxWebhookNameRunes is the spec's cap on the stored webhook name.
	maxWebhookNameRunes = 80

	invalidURLMessage  = "That is not a Discord webhook URL. It starts with https://discord.com/api/webhooks/."
	noAnswerMessage    = "Discord did not answer"
	invalidBodyMessage = "the body must be {webhook_url?, events: {sim_finished, upload_parsed, upload_private}}"
)

var (
	errTrailingData  = errors.New("notify: data after the request object")
	errMissingEvents = errors.New("notify: all three event keys are required")
)

// Entitlementer answers whether an account may use a feature. It has the
// shape sims.Entitlementer has (entitlements.Store satisfies both); it is
// declared here because sims will import this package for its own posts.
type Entitlementer interface {
	Can(ctx context.Context, userID int64, feature entitlements.Feature) (bool, entitlements.Reason, error)
}

// Service serves the notification routes and answers /v1/me's block.
type Service struct {
	Store        *Store
	Sealer       *Sealer
	Entitlements Entitlementer
	Poster       Poster
	// PublicBaseURL is the site origin the test message links to.
	PublicBaseURL string
	Log           *slog.Logger
}

// Mount registers the Discord routes. The caller mounts only when the
// deployment holds an encryption key (a Sealer exists).
func Mount(mux *http.ServeMux, s *Service) {
	limit := httpx.RateLimitByKey(testsPerMinute, time.Minute, auth.AccountRateLimitKey)
	putLimit := httpx.RateLimitByKey(putsPerMinute, time.Minute, auth.AccountRateLimitKey)
	mux.Handle("PUT /v1/notifications/discord", putLimit(auth.RequireSession(s.put)))
	mux.HandleFunc("DELETE /v1/notifications/discord", auth.RequireSession(s.remove))
	mux.Handle("POST /v1/notifications/discord/test", limit(auth.RequireSession(s.test)))
}

// DiscordBlock implements auth.NotificationsView.
func (s *Service) DiscordBlock(ctx context.Context, userID int64, entitled bool) (*auth.DiscordBlock, error) {
	w, ok, err := s.Store.Get(ctx, userID, KindDiscord)
	if err != nil || !ok {
		return nil, err
	}
	return blockOf(w, entitled), nil
}

type putRequest struct {
	WebhookURL *string `json:"webhook_url"`
	Events     *struct {
		SimFinished   *bool `json:"sim_finished"`
		UploadParsed  *bool `json:"upload_parsed"`
		UploadPrivate *bool `json:"upload_private"`
	} `json:"events"`
}

// decodePut reads the body strictly: both event keys present, no unknown
// fields, nothing after the object. The error is never shown to the caller.
func decodePut(w http.ResponseWriter, r *http.Request) (putRequest, map[EventName]bool, error) {
	var req putRequest
	dec := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxPutBodyBytes))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&req); err != nil {
		return req, nil, err
	}
	if _, err := dec.Token(); !errors.Is(err, io.EOF) {
		return req, nil, errTrailingData
	}
	if req.Events == nil || req.Events.SimFinished == nil || req.Events.UploadParsed == nil || req.Events.UploadPrivate == nil {
		return req, nil, errMissingEvents
	}
	return req, map[EventName]bool{
		EventSimFinished:   *req.Events.SimFinished,
		EventUploadParsed:  *req.Events.UploadParsed,
		EventUploadPrivate: *req.Events.UploadPrivate,
	}, nil
}

func (s *Service) put(w http.ResponseWriter, r *http.Request) {
	actor := auth.ActorFrom(r.Context())
	httpx.CachePrivateNoStore(w)
	if !s.entitled(w, r, actor) {
		return
	}
	req, events, err := decodePut(w, r)
	if err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", invalidBodyMessage, nil)
		return
	}
	if req.WebhookURL == nil {
		s.putEvents(w, r, actor, events)
		return
	}
	s.putWebhook(w, r, actor, *req.WebhookURL, events)
}

func (s *Service) putEvents(w http.ResponseWriter, r *http.Request, actor auth.Actor, events map[EventName]bool) {
	err := s.Store.SaveEvents(r.Context(), actor.UserID, KindDiscord, events)
	if errors.Is(err, ErrNotFound) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no webhook is saved", nil)
		return
	}
	if err != nil {
		s.fail(w, r, "save events", err)
		return
	}
	s.answerBlock(w, r, actor.UserID)
}

func (s *Service) putWebhook(w http.ResponseWriter, r *http.Request, actor auth.Actor, raw string, events map[EventName]bool) {
	parsed, err := ParseWebhookURL(strings.TrimSpace(raw))
	if err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", invalidURLMessage, nil)
		return
	}
	found, err := s.Poster.Lookup(r.Context(), parsed.URL())
	if err != nil {
		// Any failure to reach the webhook reads the same to the member: it does not work.
		httpx.WriteError(w, r, http.StatusUnprocessableEntity, "invalid", invalidURLMessage, nil)
		return
	}
	box, err := s.Sealer.Seal(parsed.URL())
	if err != nil {
		s.fail(w, r, "seal", err)
		return
	}
	if err := s.Store.Save(r.Context(), actor.UserID, KindDiscord, found.ChannelID, cleanWebhookName(found.Name), box, events); err != nil {
		s.fail(w, r, "save", err)
		return
	}
	s.answerBlock(w, r, actor.UserID)
}

// isInvisibleFormatting is the line/paragraph separators and the bidi
// controls: they are not unicode.IsControl but can disguise member text.
func isInvisibleFormatting(r rune) bool {
	return r == 0x2028 || r == 0x2029 || (r >= 0x202A && r <= 0x202E) || (r >= 0x2066 && r <= 0x2069)
}

// cleanWebhookName is the member-controlled name as stored: trimmed, control
// characters removed, at most maxWebhookNameRunes runes; nil when nothing is left.
func cleanWebhookName(raw string) *string {
	kept := strings.Map(func(r rune) rune {
		if unicode.IsControl(r) || isInvisibleFormatting(r) {
			return -1
		}
		return r
	}, strings.TrimSpace(raw))
	cleaned := strings.TrimSpace(truncateRunes(kept, maxWebhookNameRunes))
	if cleaned == "" {
		return nil
	}
	return &cleaned
}

func (s *Service) remove(w http.ResponseWriter, r *http.Request) {
	httpx.CachePrivateNoStore(w)
	if err := s.Store.Delete(r.Context(), auth.ActorFrom(r.Context()).UserID, KindDiscord); err != nil {
		s.fail(w, r, "delete", err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s *Service) test(w http.ResponseWriter, r *http.Request) {
	actor := auth.ActorFrom(r.Context())
	httpx.CachePrivateNoStore(w)
	if !s.entitled(w, r, actor) {
		return
	}
	hook, url, err := s.openSaved(r.Context(), actor.UserID)
	if errors.Is(err, ErrNotFound) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no webhook is saved", nil)
		return
	}
	if err != nil {
		s.fail(w, r, "open", err)
		return
	}
	if err := s.Poster.Post(r.Context(), url, TestEmbed(s.PublicBaseURL, hook.Events, time.Now())); err != nil {
		s.testFailed(w, r, err)
		return
	}
	if err := s.Store.RecordSuccess(r.Context(), actor.UserID, KindDiscord); err != nil {
		s.fail(w, r, "record success", err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// openSaved loads the member's webhook and unseals its URL; ErrNotFound when none is saved.
func (s *Service) openSaved(ctx context.Context, userID int64) (Webhook, string, error) {
	hook, ok, err := s.Store.Get(ctx, userID, KindDiscord)
	if err != nil {
		return Webhook{}, "", err
	}
	if !ok {
		return Webhook{}, "", ErrNotFound
	}
	box, err := s.Store.TokenBox(ctx, userID, KindDiscord)
	if err != nil {
		return Webhook{}, "", err
	}
	url, err := s.Sealer.Open(box)
	return hook, url, err
}

// testFailed answers a failed test post. It changes no counter: only the
// jobs' own posts count toward switching a webhook off.
func (s *Service) testFailed(w http.ResponseWriter, r *http.Request, err error) {
	message, status := noAnswerMessage, http.StatusGatewayTimeout
	var discord *DiscordError
	if errors.As(err, &discord) {
		s.logger().Warn("notify", "id", httpx.RequestIDFrom(r.Context()), "op", "test", "status", discord.Status)
		status = http.StatusBadGateway
		if body := truncateRunes(discord.Body, maxTestErrorRunes); body != "" {
			message = body
		}
	} else {
		s.logger().Warn("notify", "id", httpx.RequestIDFrom(r.Context()), "op", "test", "err", "no answer")
	}
	httpx.WriteError(w, r, status, "upstream", message, nil)
}

func truncateRunes(text string, limit int) string {
	if runes := []rune(text); len(runes) > limit {
		return string(runes[:limit])
	}
	return text
}

// entitled reports whether the actor holds the notifications feature, and
// has already written the refusal when not.
func (s *Service) entitled(w http.ResponseWriter, r *http.Request, actor auth.Actor) bool {
	allowed, _, err := s.Entitlements.Can(r.Context(), actor.UserID, entitlements.FeatureNotifications)
	if err != nil {
		s.fail(w, r, "premium", err)
		return false
	}
	if !allowed {
		httpx.WriteError(w, r, http.StatusPaymentRequired, "premium_required",
			"Discord messages are part of Premium", nil)
	}
	return allowed
}

// answerBlock writes the member's current block as the PUT answer.
func (s *Service) answerBlock(w http.ResponseWriter, r *http.Request, userID int64) {
	block, err := s.DiscordBlock(r.Context(), userID, true)
	if err != nil || block == nil {
		s.fail(w, r, "read block", err)
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, block)
}

func (s *Service) logger() *slog.Logger {
	if s.Log != nil {
		return s.Log
	}
	return slog.Default()
}

func (s *Service) fail(w http.ResponseWriter, r *http.Request, op string, err error) {
	s.logger().Error("notify", "id", httpx.RequestIDFrom(r.Context()), "op", op, "err", err)
	httpx.WriteError(w, r, http.StatusInternalServerError, "internal", "could not save that just now", nil)
}
