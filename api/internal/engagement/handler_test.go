package engagement

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

type fakeReports struct {
	rep Report
	err error
	at  time.Time
}

func (f *fakeReports) Report(_ context.Context, now time.Time) (Report, error) {
	f.at = now
	return f.rep, f.err
}

func serve(f *fakeReports, actor auth.Actor) *httptest.ResponseRecorder {
	mux := http.NewServeMux()
	now := time.Date(2026, 11, 13, 9, 0, 0, 0, time.UTC)
	Mount(mux, &Service{Reports: f, Now: func() time.Time { return now }})
	r := httptest.NewRequest(http.MethodGet, AdminPath, nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, r.WithContext(auth.WithActor(r.Context(), actor)))
	return rec
}

func TestTheFiguresAreTheOwnersAlone(t *testing.T) {
	for _, tc := range []struct {
		actor auth.Actor
		want  int
	}{
		{auth.Actor{}, http.StatusUnauthorized},
		{browser, http.StatusForbidden},
		{auth.Actor{UserID: 78, Role: "admin", Method: "device", DeviceID: "d", DeviceKind: auth.DeviceKindApp}, http.StatusForbidden},
	} {
		if rec := serve(&fakeReports{}, tc.actor); rec.Code != tc.want {
			t.Errorf("%+v: status %d, want %d", tc.actor, rec.Code, tc.want)
		}
	}
}

func TestTheOwnerReadsTheReportPrivately(t *testing.T) {
	since := "2026-11-13"
	f := &fakeReports{rep: Report{Today: "2026-11-13", LiveSince: &since, InternalAccounts: 4,
		Days: []Period{{Start: "2026-11-13", Active: 3, New: 1, Returning: 2, Site: 3}}, Weeks: []Period{}, Cohorts: []Cohort{}}}
	rec := serve(f, auth.Actor{UserID: 78, Role: "admin", Method: "session"})
	if rec.Code != http.StatusOK {
		t.Fatalf("status %d: %s", rec.Code, rec.Body)
	}
	if cc := rec.Header().Get("Cache-Control"); cc == "" || cc == "public" {
		t.Fatalf("Cache-Control = %q", cc)
	}
	var env struct {
		Data Report `json:"data"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &env); err != nil {
		t.Fatal(err)
	}
	if env.Data.InternalAccounts != 4 || *env.Data.LiveSince != since || env.Data.Days[0].Returning != 2 {
		t.Fatalf("data = %+v", env.Data)
	}
	if !f.at.Equal(time.Date(2026, 11, 13, 9, 0, 0, 0, time.UTC)) {
		t.Fatalf("report asked at %v", f.at)
	}
}

func TestAFailedReadIsA500(t *testing.T) {
	rec := serve(&fakeReports{err: errors.New("db")}, auth.Actor{UserID: 78, Role: "admin", Method: "session"})
	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("status %d", rec.Code)
	}
}
