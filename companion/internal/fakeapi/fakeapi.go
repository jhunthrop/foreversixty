// companion/internal/fakeapi/fakeapi.go
// Package fakeapi is the ingest contract implemented in memory, over
// httptest. Every companion test that talks to a server talks to this
// one, so the contract is written down once and a route that drifts
// breaks every test at the same time.
package fakeapi

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync"

	"github.com/klauspost/compress/zstd"

	"github.com/jhunthrop/foreversixty/companion/internal/client"
)

// MaxDecoded bounds decoding, per the contract: a chunk is 4 MiB of
// log, and anything claiming to be four times that is not one.
const MaxDecoded = 16 << 20

// decode unpacks a raw chunk the way the API does.
func decode(packed []byte) ([]byte, error) {
	d, err := zstd.NewReader(nil, zstd.WithDecoderMaxMemory(MaxDecoded))
	if err != nil {
		return nil, err
	}
	defer d.Close()
	return d.DecodeAll(packed, nil)
}

// Token is the device token the fake issues and the only one it
// accepts.
const Token = "fsd_aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

// PairCode is the pairing code the fake accepts.
const PairCode = "HORDE-42"

// Fight is one stored fight bundle, decoded.
type Fight struct {
	Index    int
	Summary  json.RawMessage
	Events   []byte
	Metrics  []client.MetricsRow
	RawRange client.RawRange
}

// Report is one stored report.
type Report struct {
	ID         string
	Visibility string
	Title      string
	Fights     map[int]Fight
	Live       map[int]client.Live
	Raw        map[int64][]byte
	Complete   *client.Complete
}

// Server is a fake ingest API.
type Server struct {
	*httptest.Server

	mu       sync.Mutex
	reports  map[string]*Report
	nextID   int
	order    []string
	offline  bool
	failNext map[string]int
	exports  []map[string]string
	inbox    json.RawMessage
	refuse   bool
}

// New starts a fake API. Close it with Close.
func New() *Server {
	s := &Server{
		reports:  map[string]*Report{},
		failNext: map[string]int{},
		inbox:    json.RawMessage(`{"builds":[]}`),
	}
	mux := http.NewServeMux()
	mux.HandleFunc("POST /v1/devices/claim", s.claim)
	mux.HandleFunc("POST /v1/reports", s.createReport)
	mux.HandleFunc("PUT /v1/reports/{id}/fights/{n}", s.putFight)
	mux.HandleFunc("PUT /v1/reports/{id}/fights/{n}/live", s.putLive)
	mux.HandleFunc("PUT /v1/reports/{id}/raw", s.putRaw)
	mux.HandleFunc("POST /v1/reports/{id}/complete", s.complete)
	mux.HandleFunc("POST /v1/addon/exports", s.addonExports)
	mux.HandleFunc("GET /v1/addon/inbox", s.addonInbox)
	s.Server = httptest.NewServer(s.wrap(mux))
	return s
}

// Offline makes every route answer 503, standing in for a dropped
// network without tearing the listener down.
func (s *Server) Offline(v bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.offline = v
}

// RefuseFights makes every fight PUT answer 409, standing in for the
// verification mismatch the contract describes. It is permanent, not
// retryable, which is the behaviour under test.
func (s *Server) RefuseFights(v bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.refuse = v
}

// FailNext makes the next n requests to a route pattern
// ("PUT /v1/reports/{id}/fights/{n}") answer 500.
func (s *Server) FailNext(pattern string, n int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.failNext[pattern] = n
}

// Order is every write the server accepted, in arrival order, as
// "fight 3", "raw 4194304", "complete" and so on. It is what an
// ordering assertion reads.
func (s *Server) Order() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]string(nil), s.order...)
}

// Reports returns a snapshot of everything stored.
func (s *Server) Reports() map[string]*Report {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := map[string]*Report{}
	for k, v := range s.reports {
		out[k] = v
	}
	return out
}

// Exports is every addon export the companion posted.
func (s *Server) Exports() []map[string]string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]map[string]string(nil), s.exports...)
}

// SetInbox replaces the body GET /v1/addon/inbox answers with.
func (s *Server) SetInbox(raw string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.inbox = json.RawMessage(raw)
}

func (s *Server) wrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.mu.Lock()
		off := s.offline
		s.mu.Unlock()
		if off {
			fail(w, http.StatusServiceUnavailable, "the network is down")
			return
		}
		if r.URL.Path != "/v1/devices/claim" {
			if r.Header.Get("Authorization") != "Bearer "+Token {
				fail(w, http.StatusUnauthorized, "unknown device token")
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

// shouldFail consumes one injected failure for this pattern.
func (s *Server) shouldFail(pattern string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.failNext[pattern] > 0 {
		s.failNext[pattern]--
		return true
	}
	return false
}

func ok(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]any{
		"ok": true, "data": data, "error": nil, "request_id": "fake",
	})
}

func fail(w http.ResponseWriter, status int, msg string, fields ...map[string]string) {
	body := map[string]any{"message": msg}
	if len(fields) > 0 {
		body["fields"] = fields[0]
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]any{
		"ok": false, "data": nil, "error": body, "request_id": "fake",
	})
}

func (s *Server) claim(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Code     string `json:"code"`
		Name     string `json:"name"`
		Platform string `json:"platform"`
	}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		fail(w, http.StatusBadRequest, "malformed body")
		return
	}
	if in.Code != PairCode {
		fail(w, http.StatusBadRequest, "that pairing code is not valid",
			map[string]string{"code": "expired or unknown"})
		return
	}
	ok(w, http.StatusCreated, map[string]string{
		"device_id": "dev00000000aaaa", "token": Token,
	})
}

func (s *Server) createReport(w http.ResponseWriter, r *http.Request) {
	if s.shouldFail("POST /v1/reports") {
		fail(w, http.StatusInternalServerError, "injected failure")
		return
	}
	var in client.CreateReport
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		fail(w, http.StatusBadRequest, "malformed body")
		return
	}
	s.mu.Lock()
	s.nextID++
	id := fmt.Sprintf("rpt%09d", s.nextID)
	s.reports[id] = &Report{ID: id, Visibility: in.Visibility, Title: in.Title,
		Fights: map[int]Fight{}, Live: map[int]client.Live{}, Raw: map[int64][]byte{}}
	s.order = append(s.order, "report "+id)
	s.mu.Unlock()
	ok(w, http.StatusCreated, map[string]any{"id": id, "created_at": "2026-12-09T20:00:00Z"})
}

func (s *Server) report(w http.ResponseWriter, r *http.Request) (*Report, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	rep, found := s.reports[r.PathValue("id")]
	if !found {
		fail(w, http.StatusNotFound, "no such report")
		return nil, false
	}
	return rep, true
}

func (s *Server) putFight(w http.ResponseWriter, r *http.Request) {
	if s.shouldFail("PUT /v1/reports/{id}/fights/{n}") {
		fail(w, http.StatusInternalServerError, "injected failure")
		return
	}
	s.mu.Lock()
	refuse := s.refuse
	s.mu.Unlock()
	if refuse {
		fail(w, http.StatusConflict, "the metrics do not match the events",
			map[string]string{"metrics": "dps differs by more than 0.5%"})
		return
	}
	rep, found := s.report(w, r)
	if !found {
		return
	}
	n, err := strconv.Atoi(r.PathValue("n"))
	if err != nil {
		fail(w, http.StatusBadRequest, "fight index is not a number")
		return
	}
	_, params, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		fail(w, http.StatusBadRequest, "content type: "+err.Error())
		return
	}
	f := Fight{Index: n}
	mr := multipart.NewReader(r.Body, params["boundary"])
	seen := map[string]bool{}
	for {
		part, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			fail(w, http.StatusBadRequest, "multipart: "+err.Error())
			return
		}
		b, err := io.ReadAll(part)
		if err != nil {
			fail(w, http.StatusBadRequest, "multipart: "+err.Error())
			return
		}
		seen[part.FormName()] = true
		switch part.FormName() {
		case "summary":
			f.Summary = json.RawMessage(b)
		case "events":
			f.Events = b
		case "metrics":
			if err := json.Unmarshal(b, &f.Metrics); err != nil {
				fail(w, http.StatusBadRequest, "metrics: "+err.Error())
				return
			}
		case "raw_range":
			if err := json.Unmarshal(b, &f.RawRange); err != nil {
				fail(w, http.StatusBadRequest, "raw_range: "+err.Error())
				return
			}
		}
	}
	for _, want := range []string{"summary", "events", "metrics", "raw_range"} {
		if !seen[want] {
			fail(w, http.StatusBadRequest, "the bundle is missing the "+want+" part")
			return
		}
	}
	s.mu.Lock()
	prior, already := rep.Fights[n]
	if already && prior.RawRange.SHA256 == f.RawRange.SHA256 {
		s.mu.Unlock()
		ok(w, http.StatusOK, map[string]any{"fight_index": n, "verified": true})
		return
	}
	rep.Fights[n] = f
	s.order = append(s.order, "fight "+strconv.Itoa(n))
	s.mu.Unlock()
	ok(w, http.StatusCreated, map[string]any{"fight_index": n, "verified": true})
}

func (s *Server) putLive(w http.ResponseWriter, r *http.Request) {
	rep, found := s.report(w, r)
	if !found {
		return
	}
	n, err := strconv.Atoi(r.PathValue("n"))
	if err != nil {
		fail(w, http.StatusBadRequest, "fight index is not a number")
		return
	}
	var in client.Live
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		fail(w, http.StatusBadRequest, "malformed body")
		return
	}
	s.mu.Lock()
	rep.Live[n] = in
	s.order = append(s.order, "live "+strconv.Itoa(n))
	s.mu.Unlock()
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) putRaw(w http.ResponseWriter, r *http.Request) {
	if s.shouldFail("PUT /v1/reports/{id}/raw") {
		fail(w, http.StatusInternalServerError, "injected failure")
		return
	}
	rep, found := s.report(w, r)
	if !found {
		return
	}
	offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 64)
	if err != nil {
		fail(w, http.StatusBadRequest, "offset is not a number")
		return
	}
	body, err := io.ReadAll(io.LimitReader(r.Body, 8<<20+1))
	if err != nil {
		fail(w, http.StatusBadRequest, err.Error())
		return
	}
	if len(body) > 8<<20 {
		fail(w, http.StatusRequestEntityTooLarge, "a raw chunk may not exceed 8 MiB")
		return
	}
	// The header is the hash of the DECODED chunk, so the server
	// decodes first -- bounded, because the bytes are a stranger's --
	// and the decoded length is what gives the range its end.
	decoded, derr := decode(body)
	if derr != nil {
		fail(w, http.StatusBadRequest, "the chunk did not decode: "+derr.Error())
		return
	}
	sum := sha256.Sum256(decoded)
	if got, want := r.Header.Get("X-Raw-SHA256"), hex.EncodeToString(sum[:]); got != want {
		fail(w, http.StatusBadRequest, "X-Raw-SHA256 does not match the decoded chunk")
		return
	}
	s.mu.Lock()
	if prior, already := rep.Raw[offset]; already {
		same := string(prior) == string(body)
		s.mu.Unlock()
		if !same {
			fail(w, http.StatusConflict, "that offset holds different bytes")
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}
	rep.Raw[offset] = body
	s.order = append(s.order, "raw "+strconv.FormatInt(offset, 10))
	s.mu.Unlock()
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) complete(w http.ResponseWriter, r *http.Request) {
	rep, found := s.report(w, r)
	if !found {
		return
	}
	var in client.Complete
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		fail(w, http.StatusBadRequest, "malformed body")
		return
	}
	s.mu.Lock()
	rep.Complete = &in
	s.order = append(s.order, "complete "+rep.ID)
	s.mu.Unlock()
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) addonExports(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Characters []map[string]string `json:"characters"`
	}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		fail(w, http.StatusBadRequest, "malformed body")
		return
	}
	s.mu.Lock()
	s.exports = append(s.exports, in.Characters...)
	s.mu.Unlock()
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) addonInbox(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	body := s.inbox
	s.mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"ok":true,"data":%s,"error":null,"request_id":"fake"}`, strings.TrimSpace(string(body)))
}
