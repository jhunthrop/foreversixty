// api/internal/bnetimport/realm_answer_handler.go
package bnetimport

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// RealmAnswerPath is the unknown-realm question's answer: POST {key, ruleset}.
const RealmAnswerPath = "/v1/me/realm-answers"

// realmAnswersPerHour bounds each account's answers: a few characters and a
// changed mind or two, never a loop.
const realmAnswersPerHour = 30

// maxRealmAnswerBody bounds the body: a key and a ruleset.
const maxRealmAnswerBody = 1 << 10

const (
	msgAnswerInvalid  = "send a character key and one of normal, pvp, rp or hardcore"
	msgAnswerNotYours = "that character is not one of yours waiting for its realm type"
	msgAnswerTaken    = "another account already has a character of that name on that realm type"
	msgAnswerFailed   = "could not save the realm type just now"
)

type realmAnswerBody struct {
	Key     string `json:"key"`
	Ruleset string `json:"ruleset"`
}

// MountRealmAnswers registers POST /v1/me/realm-answers behind a session and
// the per-account budget.
func MountRealmAnswers(mux *http.ServeMux, s *Service) {
	limit := httpx.RateLimitByKey(realmAnswersPerHour, time.Hour, auth.AccountRateLimitKey)
	mux.Handle("POST "+RealmAnswerPath, limit(auth.RequireSession(s.answerRealm)))
}

func (s *Service) answerRealm(w http.ResponseWriter, r *http.Request) {
	var body realmAnswerBody
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxRealmAnswerBody)).Decode(&body); err != nil ||
		strings.TrimSpace(body.Key) == "" {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", msgAnswerInvalid, nil)
		return
	}
	userID := auth.ActorFrom(r.Context()).UserID
	key, err := s.AnswerRealm(r.Context(), userID, body.Key, body.Ruleset)
	switch {
	case err == nil:
		httpx.WriteOK(w, r, http.StatusOK, map[string]string{"key": key, "ruleset": body.Ruleset})
	case errors.Is(err, ErrInvalidAnswer):
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", msgAnswerInvalid, map[string]string{"ruleset": "normal, pvp, rp or hardcore"})
	case errors.Is(err, ErrNotAwaitingAnswer):
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", msgAnswerNotYours, nil)
	case errors.Is(err, ErrAnswerKeyTaken):
		httpx.WriteError(w, r, http.StatusConflict, "taken", msgAnswerTaken, nil)
	default:
		s.logger().Error("bnetimport", "op", "realm_answer", "user_id", userID, "err", err)
		httpx.WriteError(w, r, http.StatusInternalServerError, "internal", msgAnswerFailed, nil)
	}
}
