package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequireAdminLetsOnlyAnAdminBrowserSessionThrough(t *testing.T) {
	cases := []struct {
		name  string
		actor Actor
		want  int
	}{
		{"anonymous", Actor{}, http.StatusUnauthorized},
		{"a user", Actor{UserID: 1, Role: "user", Method: "session"}, http.StatusForbidden},
		{"a moderator", Actor{UserID: 1, Role: "moderator", Method: "session"}, http.StatusForbidden},
		{"an admin's companion", Actor{UserID: 1, Role: "admin", Method: "device", DeviceID: "d", DeviceKind: DeviceKindCompanion}, http.StatusUnauthorized},
		{"an admin's app", Actor{UserID: 1, Role: "admin", Method: "device", DeviceID: "d", DeviceKind: DeviceKindApp}, http.StatusForbidden},
		{"an admin", Actor{UserID: 1, Role: "admin", Method: "session"}, http.StatusNoContent},
	}
	h := RequireAdmin(func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusNoContent) })
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/v1/admin/golinks", nil)
			h(rec, r.WithContext(WithActor(r.Context(), c.actor)))
			if rec.Code != c.want {
				t.Fatalf("status = %d, want %d", rec.Code, c.want)
			}
		})
	}
}

func TestIsAdmin(t *testing.T) {
	if !(Actor{UserID: 1, Role: "admin", Method: "session"}).IsAdmin() {
		t.Fatal("an admin session is not an admin")
	}
	if (Actor{UserID: 1, Role: "admin", DeviceID: "d", DeviceKind: DeviceKindApp}).IsAdmin() {
		t.Fatal("an app token carried admin powers")
	}
	if (Actor{UserID: 1, Role: "moderator"}).IsAdmin() {
		t.Fatal("a moderator is an admin")
	}
}
