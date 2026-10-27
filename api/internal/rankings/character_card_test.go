package rankings

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/engine"
)

// TestCharacterCardMatchesTheRoute: the route's body is CharacterCard's answer, and
// CharacterCard refuses what the route 404s.
func TestCharacterCardMatchesTheRoute(t *testing.T) {
	h := newHarness(t)
	svc := &Service{Store: h.store, Log: slog.New(slog.NewTextHandler(io.Discard, nil)),
		Now: func() time.Time { return engine.FixtureBase }}
	mux := http.NewServeMux()
	Mount(mux, svc)
	h.server = httptest.NewServer(mux)
	t.Cleanup(h.server.Close)
	h.seedOwnedCharacter("carded", &h.owner)

	res := h.get("/v1/characters/US/hardcore/Carded")
	if res.StatusCode != http.StatusOK {
		t.Fatalf("route = %d", res.StatusCode)
	}
	var routeBody map[string]any
	h.data(res, &routeBody)

	card, found, err := svc.CharacterCard(context.Background(), "US", "hardcore", "Carded")
	if err != nil || !found {
		t.Fatalf("CharacterCard = %v, %v", found, err)
	}
	raw, err := json.Marshal(card)
	if err != nil {
		t.Fatal(err)
	}
	var cardBody map[string]any
	if err := json.Unmarshal(raw, &cardBody); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(routeBody, cardBody) {
		t.Fatalf("route body %v\ncard body %v", routeBody, cardBody)
	}

	for _, c := range [][3]string{{"xx", "hardcore", "carded"}, {"us", "nope", "carded"}, {"us", "hardcore", "nobody"}} {
		if _, found, err := svc.CharacterCard(context.Background(), c[0], c[1], c[2]); err != nil || found {
			t.Errorf("CharacterCard%v = %v, %v; want not found", c, found, err)
		}
	}
}
