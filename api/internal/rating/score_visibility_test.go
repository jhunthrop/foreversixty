package rating

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/rankings"
)

// hideScore sets characters.score_visible for key.
func hideScore(t *testing.T, fx scoreFixture, key string, visible bool) {
	t.Helper()
	if _, err := fx.pool.Exec(context.Background(), `update characters set score_visible = $2 where key = $1`, key, visible); err != nil {
		t.Fatal(err)
	}
}

// characterPageAs reads the character page as viewer (0: no one signed in).
func characterPageAs(t *testing.T, fx scoreFixture, viewer int64, name string) (map[string]json.RawMessage, http.Header) {
	t.Helper()
	svc := &Service{Store: fx.store}
	mux := http.NewServeMux()
	rankings.Mount(mux, &rankings.Service{Store: &rankings.Store{Pool: fx.pool}, ScoreGear: svc})
	req := httptest.NewRequest(http.MethodGet, "/v1/characters/us/normal/"+name, nil)
	if viewer != 0 {
		req = req.WithContext(auth.WithActor(req.Context(), auth.Actor{UserID: viewer, Role: "user", Method: "session"}))
	}
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("status %d: %s", rec.Code, rec.Body)
	}
	var env struct {
		Data map[string]json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &env); err != nil {
		t.Fatal(err)
	}
	return env.Data, rec.Header()
}

// A hidden Forever Score is off the character page for everyone but its owner: no score, no
// history, no shared cache; the owner reads it whole. Visible again, it is back for everyone.
func TestAHiddenScoreIsOffTheCharacterPageButForItsOwner(t *testing.T) {
	fx := seedLoggedRaider(t)
	hideScore(t, fx, fx.main, false)
	name := strings.TrimPrefix(fx.main, "us/normal/")

	for _, viewer := range []int64{0, fx.userID + 1000} {
		page, header := characterPageAs(t, fx, viewer, name)
		for _, field := range []string{"score", "tier", "components", "as_of", "experience"} {
			if string(page[field]) != "null" {
				t.Errorf("viewer %d: %s = %s, want null", viewer, field, page[field])
			}
		}
		if string(page["score_history"]) != "[]" || string(page["score_hidden"]) != "true" {
			t.Errorf("viewer %d: score_history %s, score_hidden %s", viewer, page["score_history"], page["score_hidden"])
		}
		if _, ok := page["gear"]; !ok {
			t.Errorf("viewer %d: the gear panel went with the score", viewer)
		}
		if cc := header.Get("Cache-Control"); !strings.Contains(cc, "private") {
			t.Errorf("viewer %d: Cache-Control %q, want private", viewer, cc)
		}
	}

	own, header := characterPageAs(t, fx, fx.userID, name)
	if string(own["score"]) == "null" || own["score_hidden"] != nil {
		t.Fatalf("the owner's own page: score %s, score_hidden %s", own["score"], own["score_hidden"])
	}
	if cc := header.Get("Cache-Control"); !strings.Contains(cc, "private") {
		t.Errorf("owner: Cache-Control %q, want private", cc)
	}

	hideScore(t, fx, fx.main, true)
	back, header := characterPageAs(t, fx, 0, name)
	if string(back["score"]) == "null" || back["score_hidden"] != nil || strings.Contains(header.Get("Cache-Control"), "private") {
		t.Fatalf("visible again: score %s, score_hidden %s, cache %q", back["score"], back["score_hidden"], header.Get("Cache-Control"))
	}
}

// The scorer every other surface reads (the applicant sheet, the bot's card) answers no score
// to anyone but the owner.
func TestForeverScoreHidesAHiddenScoreFromOthers(t *testing.T) {
	fx := seedLoggedRaider(t)
	hideScore(t, fx, fx.main, false)
	svc := &Service{Store: fx.store}
	officer := auth.WithActor(context.Background(), auth.Actor{UserID: fx.userID + 1000, Role: "user", Method: "session"})
	if sc, err := svc.ForeverScore(officer, fx.main); err != nil || sc.Score != nil {
		t.Fatalf("another account: %+v %v", sc, err)
	}
	if sc, err := svc.ForeverScore(context.Background(), fx.main); err != nil || sc.Score != nil {
		t.Fatalf("no one signed in: %+v %v", sc, err)
	}
	owner := auth.WithActor(context.Background(), auth.Actor{UserID: fx.userID, Role: "user", Method: "session"})
	if sc, err := svc.ForeverScore(owner, fx.main); err != nil || sc.Score == nil {
		t.Fatalf("the owner: %+v %v", sc, err)
	}
	if hidden, err := fx.store.ScoreHidden(context.Background(), fx.main); err != nil || !hidden {
		t.Fatalf("ScoreHidden %v %v", hidden, err)
	}
}

// The Discord bot's /char reads the same card (rankings.Service.CharacterCard) as the bot, no
// account at all: a hidden score is off it.
func TestTheBotsCharacterCardHasNoHiddenScore(t *testing.T) {
	fx := seedLoggedRaider(t)
	hideScore(t, fx, fx.main, false)
	svc := &rankings.Service{Store: &rankings.Store{Pool: fx.pool}, ScoreGear: &Service{Store: fx.store}}
	card, found, err := svc.CharacterCard(context.Background(), "us", "normal", strings.TrimPrefix(fx.main, "us/normal/"))
	if err != nil || !found {
		t.Fatalf("card %v %v", found, err)
	}
	if card.Score != nil || card.Tier != nil || len(card.ScoreHistory) != 0 || !card.ScoreHidden {
		t.Fatalf("card score %v tier %v history %v hidden %v", card.Score, card.Tier, card.ScoreHistory, card.ScoreHidden)
	}
}
