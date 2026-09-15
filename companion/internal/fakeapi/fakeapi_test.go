package fakeapi

import (
	"io"
	"net/http"
	"strings"
	"testing"
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
