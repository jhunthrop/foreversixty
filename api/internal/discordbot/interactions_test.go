package discordbot

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

// harness is a service on a real store with one bound server, signed requests
// and a spy command route.
type harness struct {
	svc   *Service
	rest  *FakeREST
	priv  ed25519.PrivateKey
	mux   *http.ServeMux
	calls atomic.Int32
	log   *bytes.Buffer
	t     *testing.T
}

const (
	testServerID = "9001"
	testToken    = "tok-secret-123"
)

func newHarness(t *testing.T, handler Handler) *harness {
	t.Helper()
	st := testStore(t)
	u := seedUser(t, st.Pool)
	g := seedGuild(t, st.Pool, "Olympus")
	if _, err := st.BindServer(context.Background(), g, testServerID, "Srv", u); err != nil {
		t.Fatal(err)
	}
	pub, priv := newKeys(t)
	h := &harness{rest: &FakeREST{}, priv: priv, mux: http.NewServeMux(), log: &bytes.Buffer{}, t: t}
	h.svc = New(Deps{
		Store: st, REST: h.rest, PublicKey: pub, ApplicationID: "app",
		Log: slog_(h.log),
	})
	spy := func(ctx context.Context, in Interaction) (Reply, error) {
		h.calls.Add(1)
		if handler != nil {
			return handler(ctx, in)
		}
		return Reply{Message: Message{Content: "ok", Ephemeral: true}}, nil
	}
	h.svc.commands = map[string]Handler{"spy": spy}
	h.svc.components = map[string]Handler{"spy": spy}
	h.svc.updating = map[string]bool{"upd": true}
	h.svc.components["upd"] = spy
	Mount(h.mux, h.svc)
	return h
}

func (h *harness) post(body string, at time.Time, tamper func(*http.Request)) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodPost, InteractionsPath, strings.NewReader(body))
	sig, ts := signedAt(h.t, h.priv, at, []byte(body))
	req.Header.Set("X-Signature-Ed25519", sig)
	req.Header.Set("X-Signature-Timestamp", ts)
	if tamper != nil {
		tamper(req)
	}
	rec := httptest.NewRecorder()
	h.mux.ServeHTTP(rec, req)
	return rec
}

func commandBody(name string) string {
	return `{"id":"i1","application_id":"app","type":2,"token":"` + testToken + `","guild_id":"` + testServerID +
		`","member":{"user":{"id":"d1"}},"data":{"name":"` + name + `"}}`
}

func decode(t *testing.T, rec *httptest.ResponseRecorder) map[string]any {
	t.Helper()
	var out map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &out); err != nil {
		t.Fatalf("body %q: %v", rec.Body.String(), err)
	}
	return out
}

func TestInteractionsAnswersPing(t *testing.T) {
	h := newHarness(t, nil)
	rec := h.post(`{"type":1}`, time.Now(), nil)
	if rec.Code != 200 || decode(t, rec)["type"] != float64(1) {
		t.Fatalf("ping = %d %s", rec.Code, rec.Body)
	}
}

func TestInteractionsRejectsReplayAndTamper(t *testing.T) {
	h := newHarness(t, nil)
	body := commandBody("spy")
	otherSig, otherTS := signedAt(t, h.priv, time.Now(), []byte(commandBody("other")))
	cases := map[string]*httptest.ResponseRecorder{
		"replay": h.post(body, time.Now().Add(-10*time.Minute), nil),
		"different body": h.post(body, time.Now(), func(r *http.Request) {
			r.Header.Set("X-Signature-Ed25519", otherSig)
			r.Header.Set("X-Signature-Timestamp", otherTS)
		}),
		"no headers": h.post(body, time.Now(), func(r *http.Request) {
			r.Header.Del("X-Signature-Ed25519")
			r.Header.Del("X-Signature-Timestamp")
		}),
		"oversized": h.post(strings.Repeat("a", MaxInteractionBytes+1), time.Now(), nil),
	}
	for name, rec := range cases {
		if rec.Code != http.StatusUnauthorized {
			t.Errorf("%s: status %d", name, rec.Code)
		}
	}
	if h.calls.Load() != 0 {
		t.Fatal("a handler ran for a rejected request")
	}
}

func TestUnknownCommandAnswersTheFailureSentence(t *testing.T) {
	h := newHarness(t, nil)
	rec := h.post(commandBody("nope"), time.Now(), nil)
	data, _ := decode(t, rec)["data"].(map[string]any)
	if data["content"] != commonCopy.Failure {
		t.Fatalf("body %s", rec.Body)
	}
}

func TestUnboundServerAnswersTheSpecSentence(t *testing.T) {
	h := newHarness(t, nil)
	body := strings.Replace(commandBody("spy"), testServerID, "424242", 1)
	rec := h.post(body, time.Now(), nil)
	data, _ := decode(t, rec)["data"].(map[string]any)
	want := specString(t, "### 6.2 In Discord", "Unbound server", 0, nil)
	if data["content"] != want || h.calls.Load() != 0 {
		t.Fatalf("body %s calls %d", rec.Body, h.calls.Load())
	}
}

func TestAutocompleteNeverDefers(t *testing.T) {
	h := newHarness(t, nil)
	h.svc.answerWithin = 20 * time.Millisecond
	h.svc.autocompletes = map[string]Autocomplete{"spy": func(ctx context.Context, _ Interaction) ([]Choice, error) {
		<-ctx.Done()
		return []Choice{{Name: "late", Value: "late"}}, nil
	}}
	body := `{"id":"i","type":4,"token":"t","guild_id":"` + testServerID + `","data":{"name":"spy","options":[{"name":"q","value":"x","focused":true}]}}`
	rec := h.post(body, time.Now(), nil)
	got := decode(t, rec)
	data, _ := got["data"].(map[string]any)
	choices, ok := data["choices"].([]any)
	if got["type"] != float64(8) || !ok || len(choices) != 0 {
		t.Fatalf("body %s", rec.Body)
	}
}

func TestInteractionAccessors(t *testing.T) {
	in, err := ParseInteraction([]byte(`{"id":"i","type":2,"token":"t","guild_id":"g","user":{"id":"u"},
		"message":{"id":"m"},"data":{"name":"bis","custom_id":"raid:1:in","values":["a"],
		"options":[{"name":"slot","value":"head"},{"name":"n","value":3,"focused":true}]}}`), time.Now())
	if err != nil {
		t.Fatal(err)
	}
	slot, _ := in.Option("slot")
	name, value := in.Focused()
	if in.DiscordUserID != "u" || in.MessageID != "m" || in.CommandName() != "bis" || in.CustomID() != "raid:1:in" ||
		in.Values()[0] != "a" || slot != "head" || name != "n" || value != "3" {
		t.Fatalf("parsed %+v", in)
	}
}
