package push

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"regexp"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// maxRegisterBody bounds a register body; an FCM token is at most a few hundred bytes.
const maxRegisterBody = 8 << 10

// The token shapes each platform issues: APNs device tokens are hex, FCM registration
// tokens are URL-safe text with a colon.
var tokenShapes = map[Platform]*regexp.Regexp{
	IOS:     regexp.MustCompile(`^[0-9a-fA-F]{64,200}$`),
	Android: regexp.MustCompile(`^[A-Za-z0-9_:\-]+$`),
}

// The token length bounds migration 0061 holds (a regexp repeat stops at 1000).
const (
	minTokenLen = 16
	maxTokenLen = 4096
)

// appVersionShape is a release like 1.0.0 or 1.0 (42).
var appVersionShape = regexp.MustCompile(`^[0-9A-Za-z .()+_\-]{1,32}$`)

// TokenStore is what the routes and the sink use (Store).
type TokenStore interface {
	Register(ctx context.Context, t Token) error
	Unregister(ctx context.Context, deviceID, token string) (bool, error)
	TokensOf(ctx context.Context, userID int64) ([]Token, error)
	Delete(ctx context.Context, token string) error
}

// Handler serves POST /v1/me/push-tokens and DELETE /v1/me/push-tokens/{token}.
type Handler struct {
	Store TokenStore
	Log   *slog.Logger
}

// Mount adds the token routes. Both need a device token: a token belongs to the app install
// that holds it.
func Mount(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("POST /v1/me/push-tokens", auth.RequireDevice(h.register))
	mux.HandleFunc("DELETE /v1/me/push-tokens/{token}", auth.RequireDevice(h.unregister))
}

func (h *Handler) logger() *slog.Logger {
	if h.Log != nil {
		return h.Log
	}
	return slog.Default()
}

type registerRequest struct {
	Platform   Platform `json:"platform"`
	Token      string   `json:"token"`
	AppVersion string   `json:"app_version"`
}

// validate names each field that is wrong.
func (req registerRequest) validate() map[string]string {
	fields := map[string]string{}
	shape, ok := tokenShapes[req.Platform]
	if !ok {
		fields["platform"] = "ios or android"
	} else if len(req.Token) < minTokenLen || len(req.Token) > maxTokenLen || !shape.MatchString(req.Token) {
		fields["token"] = "not a push token for that platform"
	}
	if req.AppVersion != "" && !appVersionShape.MatchString(req.AppVersion) {
		fields["app_version"] = "a release, at most 32 characters"
	}
	if len(fields) == 0 {
		return nil
	}
	return fields
}

func (h *Handler) register(w http.ResponseWriter, r *http.Request) {
	var req registerRequest
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxRegisterBody)).Decode(&req); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON", nil)
		return
	}
	if fields := req.validate(); fields != nil {
		httpx.WriteError(w, r, http.StatusUnprocessableEntity, "invalid", "that push token is not valid", fields)
		return
	}
	actor := auth.ActorFrom(r.Context())
	err := h.Store.Register(r.Context(), Token{Token: req.Token, Platform: req.Platform, UserID: actor.UserID,
		DeviceID: actor.DeviceID, AppVersion: req.AppVersion})
	if err != nil {
		h.logger().Error("push", "op", "register", "device", actor.DeviceID, "err", err)
		httpx.WriteError(w, r, http.StatusInternalServerError, "internal", "could not save that push token", nil)
		return
	}
	h.logger().Info("push", "op", "register", "user", actor.UserID, "device", actor.DeviceID, "platform", string(req.Platform))
	httpx.WriteOK(w, r, http.StatusCreated, map[string]string{"platform": string(req.Platform), "app_version": req.AppVersion})
}

// unregister deletes the calling device's token. It answers 204 whether or not the token
// was there, and never touches another device's token.
func (h *Handler) unregister(w http.ResponseWriter, r *http.Request) {
	actor := auth.ActorFrom(r.Context())
	if _, err := h.Store.Unregister(r.Context(), actor.DeviceID, r.PathValue("token")); err != nil {
		h.logger().Error("push", "op", "unregister", "device", actor.DeviceID, "err", err)
		httpx.WriteError(w, r, http.StatusInternalServerError, "internal", "could not remove that push token", nil)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
