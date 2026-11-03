package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/golinks"
)

type oneLink struct{}

func (oneLink) Target(context.Context, string) (string, error) { return "/", nil }
func (oneLink) CountClick(context.Context, string, bool) (bool, error) {
	return true, nil
}
func (oneLink) RecordArrival(context.Context, int64, string) (bool, error) { return false, nil }
func (oneLink) Create(_ context.Context, l golinks.Link) (golinks.Link, error) {
	return l, nil
}
func (oneLink) Funnel(context.Context) ([]golinks.FunnelRow, error) { return nil, nil }

// Every /go/ click reaches the API from the Worker's few egress addresses, so the resolve and
// the count skip the per-IP budget; the owner's routes keep it.
func TestTheWorkersLinkRoutesSkipTheRouterLimiter(t *testing.T) {
	h := NewRouter(Deps{Version: "test", Auth: &auth.Authenticator{}, GoLinks: &golinks.Service{Links: oneLink{}}})
	for i := 0; i <= routerRequestsPerMinute; i++ {
		for _, req := range []*http.Request{
			httptest.NewRequest(http.MethodGet, "/v1/golinks/reddit", nil),
			httptest.NewRequest(http.MethodPost, "/v1/golinks/reddit/clicks", nil),
		} {
			rec := httptest.NewRecorder()
			h.ServeHTTP(rec, req)
			if rec.Code == http.StatusTooManyRequests {
				t.Fatalf("%s %s was limited at request %d", req.Method, req.URL.Path, i+1)
			}
		}
	}
	last := 0
	for i := 0; i <= routerRequestsPerMinute; i++ {
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/v1/admin/golinks", nil))
		last = rec.Code
	}
	if last != http.StatusTooManyRequests {
		t.Fatalf("the admin list past the budget = %d, want 429", last)
	}
}
