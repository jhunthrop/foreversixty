package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/apptesters"
	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

type noTesters struct{}

func (noTesters) Add(context.Context, string, string) (bool, error) { return true, nil }
func (noTesters) List(context.Context) ([]apptesters.Tester, error) { return nil, nil }

// The beta signup is mounted, open to a visitor with no session, and the owner's list is not.
func TestTheAppTesterRoutesAreMounted(t *testing.T) {
	h := NewRouter(Deps{Version: "test", Auth: &auth.Authenticator{}, AppTesters: &apptesters.Service{Testers: noTesters{}}})

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, apptesters.SignupPath, strings.NewReader(`{"email":"a@example.com","platform":"ios"}`))
	h.ServeHTTP(rec, req)
	if rec.Code != http.StatusCreated {
		t.Fatalf("signup: %d %s", rec.Code, rec.Body)
	}

	rec = httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, apptesters.AdminPath, nil))
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("list without a session: %d", rec.Code)
	}
}
