package coach

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/contract"
)

// The coach routes the apps read (W3-C4): the night line, a fight's review and the account's
// recent line, each pinned to its contract fixture. COACH_CONTRACT_DUMP=1 prints the bodies.
func (f *t31) raw(t *testing.T, actor auth.Actor, path string, want int) []byte {
	t.Helper()
	req := httptest.NewRequest(http.MethodGet, path, nil)
	req = req.WithContext(auth.WithActor(req.Context(), actor))
	rec := httptest.NewRecorder()
	f.mux.ServeHTTP(rec, req)
	if rec.Code != want {
		t.Fatalf("%s = %d, want %d: %s", path, rec.Code, want, rec.Body.String())
	}
	if os.Getenv("COACH_CONTRACT_DUMP") != "" {
		t.Logf("%s\n%s", path, rec.Body.String())
	}
	return rec.Body.Bytes()
}

var coachFightOptional = []string{"data.cap"}

func TestAppContractCoachNight(t *testing.T) {
	f := newT31(t)
	at := capsNow.Add(-24 * time.Hour)
	f.rate(t, f.furyKey, 1, at)
	contract.AssertFixture(t, "coach.night", f.raw(t, user(f.premium), f.nightPath(), http.StatusOK), contract.Options{})
	contract.AssertFixture(t, "coach.night.error-premium-required",
		f.raw(t, user(f.freeOwner), f.nightPath(), http.StatusForbidden), contract.Options{})
}

func TestAppContractCoachFight(t *testing.T) {
	f := newT31(t)
	f.store(t, SourceModel, false)
	path := "/v1/reports/" + t30Report + "/fights/1/coach?player=" + f.furyKey
	contract.AssertFixture(t, "coach.fight", f.raw(t, user(f.premium), path, http.StatusOK),
		contract.Options{Optional: coachFightOptional})
}

func TestAppContractCoachRecent(t *testing.T) {
	f := newT31(t)
	f.linkCharacter(t, f.premium, true)
	f.rate(t, f.furyKey, 1, capsNow.Add(-24*time.Hour))
	f.seedReviews(t, 1)
	contract.AssertFixture(t, "coach.recent", f.raw(t, user(f.premium), "/v1/me/coach/recent", http.StatusOK),
		contract.Options{Optional: []string{"data.last_raid.first"}})
}
