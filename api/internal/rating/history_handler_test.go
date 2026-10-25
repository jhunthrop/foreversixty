package rating

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
)

type stubEntitlements struct{ allowed bool }

func (s stubEntitlements) Can(context.Context, int64, entitlements.Feature) (bool, entitlements.Reason, error) {
	return s.allowed, "", nil
}

const historyPath = "/v1/characters/us/normal/Histowner/history"

// historyFixture links us/normal/histowner to a fresh account and returns the service, the
// owner's id and a second account's id.
func historyFixture(t *testing.T, entitled bool) (*Service, int64, int64) {
	t.Helper()
	svc, rs := newTestService(t)
	pool := rs.Pool
	ctx := context.Background()
	ids := make([]int64, 2)
	for i, email := range []string{"history-owner@example.com", "history-other@example.com"} {
		if err := pool.QueryRow(ctx,
			`insert into users (email) values ($1)
			 on conflict (email) do update set email = excluded.email returning id`, email).Scan(&ids[i]); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := pool.Exec(ctx,
		`insert into characters (key, region, ruleset, name, user_id) values ('us/normal/histowner','us','normal','Histowner',$1)
		 on conflict (key) do update set user_id = $1`, ids[0]); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		pool.Exec(ctx, `delete from characters where key = 'us/normal/histowner'`)
		pool.Exec(ctx, `delete from users where email in ('history-owner@example.com','history-other@example.com')`)
	})
	svc.Owners = &auth.Store{Pool: pool}
	svc.Entitlements = stubEntitlements{allowed: entitled}
	return svc, ids[0], ids[1]
}

func getHistory(svc *Service, userID int64) *httptest.ResponseRecorder {
	mux := http.NewServeMux()
	Mount(mux, svc)
	req := httptest.NewRequest(http.MethodGet, historyPath, nil)
	if userID != 0 {
		req = req.WithContext(auth.WithActor(req.Context(), auth.Actor{UserID: userID, Role: "user", Method: "session"}))
	}
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	return rec
}

func TestCharacterHistoryServesTheOwnerEveryArray(t *testing.T) {
	svc, owner, _ := historyFixture(t, true)
	rec := getHistory(svc, owner)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d: %s", rec.Code, rec.Body)
	}
	if cc := rec.Header().Get("Cache-Control"); cc != "private, no-store" {
		t.Fatalf("Cache-Control = %q", cc)
	}
	var body struct {
		Data map[string]json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	for _, k := range []string{"score", "encounters", "gaps", "fights"} {
		if string(body.Data[k]) != "[]" {
			t.Fatalf("%s = %s, want []", k, body.Data[k])
		}
	}
	if len(body.Data["as_of"]) == 0 {
		t.Fatal("as_of missing")
	}
}

func TestCharacterHistoryIs402ForAnUnentitledOwner(t *testing.T) {
	svc, owner, _ := historyFixture(t, false)
	if rec := getHistory(svc, owner); rec.Code != http.StatusPaymentRequired {
		t.Fatalf("status = %d, want 402", rec.Code)
	}
}

func TestCharacterHistoryIs404ForAnotherAccount(t *testing.T) {
	svc, _, other := historyFixture(t, true)
	if rec := getHistory(svc, other); rec.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want 404", rec.Code)
	}
}

func TestCharacterHistoryIs401WhenAnonymous(t *testing.T) {
	svc, _, _ := historyFixture(t, true)
	if rec := getHistory(svc, 0); rec.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, want 401", rec.Code)
	}
}

func TestCharacterHistoryIs404ForAnAnonymisedOwner(t *testing.T) {
	svc, owner, _ := historyFixture(t, true)
	if _, err := svc.Store.Pool.Exec(context.Background(),
		`update users set anonymize = true where id = $1`, owner); err != nil {
		t.Fatal(err)
	}
	if rec := getHistory(svc, owner); rec.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want 404", rec.Code)
	}
}

func TestCharacterHistoryIsNotMountedWithoutOwnersAndEntitlements(t *testing.T) {
	svc, owner, _ := historyFixture(t, true)
	svc.Owners = nil
	if rec := getHistory(svc, owner); rec.Code == http.StatusOK {
		t.Fatalf("status = %d, want the route unmounted", rec.Code)
	}
}
