package fakeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

func get(t *testing.T, url, token string) (int, string) {
	t.Helper()
	req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, url, nil)
	if err != nil {
		t.Fatal(err)
	}
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	return resp.StatusCode, string(b)
}

func TestEveryRouteNeedsTheDeviceTokenExceptTheClaim(t *testing.T) {
	s := New()
	defer s.Close()
	if code, _ := get(t, s.URL+"/v1/addon/inbox", ""); code != http.StatusUnauthorized {
		t.Errorf("inbox without a token = %d", code)
	}
	if code, body := get(t, s.URL+"/v1/addon/inbox", Token); code != http.StatusOK ||
		!strings.Contains(body, `"builds"`) {
		t.Errorf("inbox = %d %s", code, body)
	}
}

func TestOfflineFailsEveryRoute(t *testing.T) {
	s := New()
	defer s.Close()
	s.Offline(true)
	if code, _ := get(t, s.URL+"/v1/addon/inbox", Token); code != http.StatusServiceUnavailable {
		t.Errorf("offline inbox = %d", code)
	}
	s.Offline(false)
	if code, _ := get(t, s.URL+"/v1/addon/inbox", Token); code != http.StatusOK {
		t.Errorf("back online = %d", code)
	}
}

func TestFailNextIsConsumedOnce(t *testing.T) {
	s := New()
	defer s.Close()
	s.FailNext("POST /v1/reports", 1)
	if !s.shouldFail("POST /v1/reports") {
		t.Fatal("the first call did not fail")
	}
	if s.shouldFail("POST /v1/reports") {
		t.Fatal("the failure was not consumed")
	}
}

func TestSetInboxReplacesTheBody(t *testing.T) {
	s := New()
	defer s.Close()
	s.SetInbox(`{"builds":[{"id":"b1","name":"Holy","code":"FSB1:x"}]}`)
	_, body := get(t, s.URL+"/v1/addon/inbox", Token)
	if !strings.Contains(body, "FSB1:x") {
		t.Fatalf("inbox = %s", body)
	}
	if len(s.Order()) != 0 || len(s.Reports()) != 0 || len(s.Exports()) != 0 {
		t.Error("a read changed the recorded state")
	}
}

func post(t *testing.T, url, body string) (int, string) {
	t.Helper()
	req, err := http.NewRequestWithContext(t.Context(), http.MethodPost, url, strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+Token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	return resp.StatusCode, string(b)
}

const exportWithOutbox = `{"characters":[{"name":"Obnoxious","ruleset":"normal","region":"us","export":"FS1:x"}],
	"outbox":{"format":1,"export":{"character":"Obnoxious-Yell","code":"FS1:x","saved_at":"2026-11-04 14:02"},
	"sim_request":{"character":"Obnoxious-Yell","spec":"warrior-arms","requested_at":"2026-11-04 14:10","gear_code":"FS1:x"},
	"roster_snapshot":{"guild":"OLYMPUS XXVII","members":[{"name":"Obnoxious-Yell","level":20},{"level":30}]}}}`

func TestAnExportWithAnOutboxIsAnsweredPerPartAndAcked(t *testing.T) {
	s := New()
	defer s.Close()
	at := time.Date(2026, 11, 4, 14, 3, 0, 0, time.UTC)
	s.SetClock(func() time.Time { return at })

	// Before any outbox, the inbox is v2 and carries no ack.
	_, body := get(t, s.URL+"/v1/addon/inbox", Token)
	if !strings.Contains(body, `"format":2`) || strings.Contains(body, `"ack"`) {
		t.Fatalf("inbox before an outbox = %s", body)
	}

	code, body := post(t, s.URL+"/v1/addon/exports", exportWithOutbox)
	if code != http.StatusOK {
		t.Fatalf("exports = %d %s", code, body)
	}
	var env struct {
		Data struct {
			Stored int `json:"stored"`
			Outbox struct {
				Status         string `json:"status"`
				Export         string `json:"export"`
				SimRequest     string `json:"sim_request"`
				RosterSnapshot string `json:"roster_snapshot"`
				DroppedMembers int    `json:"dropped_members"`
			} `json:"outbox"`
		} `json:"data"`
	}
	if err := json.Unmarshal([]byte(body), &env); err != nil {
		t.Fatal(err)
	}
	if o := env.Data.Outbox; env.Data.Stored != 1 || o.Status != "read" || o.Export != "stored" ||
		o.SimRequest != "queued" || o.RosterSnapshot != "stored" || o.DroppedMembers != 1 {
		t.Errorf("answer = %s", body)
	}
	posts := s.ExportPosts()
	if len(posts) != 1 || len(posts[0].Characters) != 1 || !strings.Contains(string(posts[0].Outbox), "warrior-arms") {
		t.Fatalf("recorded = %+v", posts)
	}

	_, body = get(t, s.URL+"/v1/addon/inbox", Token)
	if !strings.Contains(body, `{"type":"ack","character":"us/normal/obnoxious","export_seen_at":"2026-11-04T14:03:00Z","sim_request_state":"queued"}`) {
		t.Errorf("inbox after the outbox = %s", body)
	}

	// A premium refusal is the ack's needs_premium; a refused outbox is a 400 that stores nothing.
	s.SetSimOutcome("premium_required")
	post(t, s.URL+"/v1/addon/exports", exportWithOutbox)
	if _, body = get(t, s.URL+"/v1/addon/inbox", Token); !strings.Contains(body, `"sim_request_state":"needs_premium"`) {
		t.Errorf("inbox after a premium refusal = %s", body)
	}
	s.RefuseOutbox(true)
	if code, _ := post(t, s.URL+"/v1/addon/exports", exportWithOutbox); code != http.StatusBadRequest {
		t.Errorf("a refused outbox = %d", code)
	}
	if len(s.ExportPosts()) != 2 {
		t.Errorf("a refused post was recorded: %d", len(s.ExportPosts()))
	}
}
