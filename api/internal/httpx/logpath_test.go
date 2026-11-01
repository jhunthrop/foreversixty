package httpx

import (
	"bytes"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLogPathRedactsTheCalendarToken(t *testing.T) {
	cases := map[string]string{
		"/v1/calendar/abcDEF_-123.ics": "/v1/calendar/:redacted",
		"/v1/calendar/":                "/v1/calendar/",
		"/v1/me/calendar-feed":         "/v1/me/calendar-feed",
		"/v1/guilds/2/home":            "/v1/guilds/2/home",
		"//v1/calendar/tok.ics":        "/v1/calendar/:redacted",
		"/v1/./calendar/tok.ics":       "/v1/calendar/:redacted",
		"/a/../v1/calendar/tok.ics":    "/v1/calendar/:redacted",
		"/v1/calendar//tok.ics":        "/v1/calendar/:redacted",
	}
	for in, want := range cases {
		if got := LogPath(in); got != want {
			t.Errorf("LogPath(%q) = %q, want %q", in, got, want)
		}
	}
}

// The request log and the panic log never carry a calendar token.
func TestLoggerNeverWritesTheCalendarToken(t *testing.T) {
	const secret = "s3cretTokenValueThatMustNotAppearInTheLogs0"
	var buf bytes.Buffer
	log := slog.New(slog.NewTextHandler(&buf, nil))
	h := Chain(http.HandlerFunc(func(http.ResponseWriter, *http.Request) { panic("boom") }), Recover(log), Logger(log))
	h.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/v1/calendar/"+secret+".ics", nil))
	if strings.Contains(buf.String(), secret) {
		t.Fatalf("the log carries the token: %s", buf.String())
	}
	if !strings.Contains(buf.String(), "/v1/calendar/:redacted") {
		t.Fatalf("the log names the route: %s", buf.String())
	}
}
