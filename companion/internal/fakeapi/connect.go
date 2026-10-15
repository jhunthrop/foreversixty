// companion/internal/fakeapi/connect.go
// The five device-authorization routes (companion-connect-contract.md),
// enough of them for the companion's connect flow to be tested end to
// end: a started connect stays pending until a test approves, denies
// or expires it, an approved row answers its token exactly once, and a
// second poll is a 410. The session-only routes (lookup, approve,
// deny) are served without a session here; the real API's session
// check is its own tests' business.
package fakeapi

import (
	"encoding/json"
	"net/http"
	"strings"
)

// ConnectID and ConnectCode are the first connect the fake issues;
// later ones count up from the same stem.
const (
	ConnectID   = "connect-secret-0001"
	ConnectCode = "KQ4-7TM"
)

type connectRow struct {
	code     string
	name     string
	platform string
	status   string // pending|approved|denied|expired
	taken    bool   // the approved answer was collected
	polls    int
}

// Connects is how many connects have been started.
func (s *Server) Connects() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.connectSeq
}

// ConnectPolls is how many times the connect with this id was polled.
func (s *Server) ConnectPolls(id string) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	if r := s.connects[id]; r != nil {
		return r.polls
	}
	return 0
}

// ConnectName is the device name the connect was started with.
func (s *Server) ConnectName(id string) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	if r := s.connects[id]; r != nil {
		return r.name
	}
	return ""
}

// ConnectDecide sets the newest started connect's status, as the site
// does when the player presses Connect ("approved"), Cancel
// ("denied") or lets ten minutes pass ("expired").
func (s *Server) ConnectDecide(status string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if r := s.connects[connectIDOf(s.connectSeq)]; r != nil {
		r.status = status
	}
}

func connectIDOf(n int) string {
	if n <= 1 {
		return ConnectID
	}
	return ConnectID[:len(ConnectID)-1] + string(rune('0'+n%10))
}

func (s *Server) connectStart(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Name     string `json:"name"`
		Platform string `json:"platform"`
	}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil || strings.TrimSpace(in.Name) == "" {
		fail(w, http.StatusBadRequest, "a device needs a name", map[string]string{"name": "required"})
		return
	}
	s.mu.Lock()
	s.connectSeq++
	id := connectIDOf(s.connectSeq)
	s.connects[id] = &connectRow{code: ConnectCode, name: in.Name, platform: in.Platform, status: "pending"}
	s.mu.Unlock()
	ok(w, http.StatusCreated, map[string]any{
		"connect_id": id, "user_code": ConnectCode,
		"url":        s.URL + "/connect?code=" + ConnectCode,
		"expires_in": 600, "interval": 2,
	})
}

func (s *Server) connectPoll(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	row := s.connects[r.PathValue("id")]
	if row == nil {
		s.mu.Unlock()
		fail(w, http.StatusNotFound, "unknown connect")
		return
	}
	row.polls++
	status, taken := row.status, row.taken
	if status == "approved" {
		row.taken = true
	}
	s.mu.Unlock()
	switch {
	case status == "approved" && taken:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusGone)
		_ = json.NewEncoder(w).Encode(map[string]any{"ok": false, "data": map[string]string{"status": "used"},
			"error": map[string]string{"message": "already collected"}, "request_id": "fake"})
	case status == "approved":
		s.mu.Lock()
		s.revoked = false
		s.mu.Unlock()
		ok(w, http.StatusOK, map[string]string{"status": "approved",
			"device_id": "dev00000000bbbb", "token": Token, "account_label": AccountLabel})
	default:
		ok(w, http.StatusOK, map[string]string{"status": status})
	}
}

func (s *Server) connectLookup(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, row := range s.connects {
		if row.code == code && row.status == "pending" {
			ok(w, http.StatusOK, map[string]any{"name": row.name, "platform": row.platform, "expires_in": 600})
			return
		}
	}
	fail(w, http.StatusNotFound, "unknown code")
}

func (s *Server) decideByCode(w http.ResponseWriter, r *http.Request, status string) {
	var in struct {
		Code string `json:"code"`
	}
	_ = json.NewDecoder(r.Body).Decode(&in)
	s.mu.Lock()
	defer s.mu.Unlock()
	// The fake hands every connect the same code, so a decided row
	// from an earlier connect must not shadow the pending one.
	decided := false
	for _, row := range s.connects {
		if row.code != in.Code {
			continue
		}
		if row.status != "pending" {
			decided = true
			continue
		}
		row.status = status
		if status == "approved" {
			ok(w, http.StatusOK, map[string]string{"name": row.name, "platform": row.platform})
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
		return
	}
	if decided {
		fail(w, http.StatusConflict, "already decided")
		return
	}
	fail(w, http.StatusNotFound, "unknown code")
}

func (s *Server) connectApprove(w http.ResponseWriter, r *http.Request) {
	s.decideByCode(w, r, "approved")
}

func (s *Server) connectDeny(w http.ResponseWriter, r *http.Request) {
	s.decideByCode(w, r, "denied")
}
