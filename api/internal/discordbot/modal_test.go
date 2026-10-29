package discordbot

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"testing"
	"time"
)

const (
	testCardEmbeds     = `[{"description":"Ready. Show BiS, or change a pick.","color":5793266}]`
	testCardComponents = `[{"type":1,"components":[{"type":3,"custom_id":"gbis:1:class:","options":[{"label":"Hunter","value":"hunter","default":true}]}]},{"type":1,"components":[{"type":2,"style":3,"label":"Show BiS","custom_id":"gbis:1:show:hunter"}]}]`
)

func componentBodyFor(customID, extra string) string {
	return `{"id":"i1","application_id":"app","type":3,"token":"` + testToken + `","guild_id":"` + testServerID +
		`","member":{"user":{"id":"d1"}},"data":{"custom_id":"` + customID + `"},"message":{"id":"m1",` + extra + `}}`
}

func cardMessageFields() string {
	return `"embeds":` + testCardEmbeds + `,"components":` + testCardComponents
}

func failing(context.Context, Interaction) (Reply, error) { return Reply{}, errors.New("boom") }

func cardHarness(t *testing.T, prefix string, handler Handler) *harness {
	t.Helper()
	h := newHarness(t, handler)
	h.svc.components[prefix] = h.svc.components["spy"]
	h.svc.updating[prefix] = true
	return h
}

func compact(t *testing.T, v any) string {
	t.Helper()
	if s, ok := v.(string); ok {
		var parsed any
		if err := json.Unmarshal([]byte(s), &parsed); err != nil {
			t.Fatal(err)
		}
		v = parsed
	}
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func TestModalOpensAtOnce(t *testing.T) {
	h := newHarness(t, func(context.Context, Interaction) (Reply, error) {
		return Reply{Modal: &Modal{
			CustomID: "gchar:1:name:", Title: "Type a name", InputLabel: "Character name",
			Placeholder: "Zulmara", MinLength: 2, MaxLength: 12,
		}}, nil
	})
	body := strings.Replace(commandBody("spy"), `"type":2`, `"type":3`, 1)
	body = strings.Replace(body, `"name":"spy"`, `"custom_id":"spy:1"`, 1)
	rec := h.post(body, time.Now(), nil)
	want := `{"type":9,"data":{"custom_id":"gchar:1:name:","title":"Type a name","components":[{"type":1,"components":[` +
		`{"type":4,"custom_id":"name","style":1,"label":"Character name","placeholder":"Zulmara","min_length":2,"max_length":12,"required":true}]}]}}`
	if strings.TrimSpace(rec.Body.String()) != want {
		t.Fatalf("modal %s", rec.Body)
	}
}

func TestModalSubmitRoutesByCustomID(t *testing.T) {
	var got string
	h := newHarness(t, func(_ context.Context, in Interaction) (Reply, error) {
		got = in.ModalValue("name") + "|" + in.ModalValue("other")
		return Reply{Message: Message{Content: "ok", Ephemeral: true}}, nil
	})
	body := `{"id":"i1","application_id":"app","type":5,"token":"` + testToken + `","guild_id":"` + testServerID +
		`","member":{"user":{"id":"d1"}},"data":{"custom_id":"spy:1:name:","components":[{"type":1,"components":[{"type":4,"custom_id":"name","value":"Zulmara"}]}]}}`
	data, _ := decode(t, h.post(body, time.Now(), nil))["data"].(map[string]any)
	if data["content"] != "ok" || got != "Zulmara|" {
		t.Fatalf("content %v value %q", data["content"], got)
	}
}

func TestCardFailureKeepsTheCard(t *testing.T) {
	h := cardHarness(t, "gbis", failing)
	got := decode(t, h.post(componentBodyFor("gbis:1:show:hunter", cardMessageFields()), time.Now(), nil))
	data, _ := got["data"].(map[string]any)
	embeds, _ := data["embeds"].([]any)
	first, _ := embeds[0].(map[string]any)
	if got["type"] != float64(7) || first["description"] != guideCopy.CardFailure || first["color"] != float64(5793266) || data["content"] != "" {
		t.Fatalf("answer %v", got)
	}
	if compact(t, data["components"]) != compact(t, testCardComponents) {
		t.Fatalf("components changed: %v", data["components"])
	}

	h = cardHarness(t, "gsgn", failing)
	data, _ = decode(t, h.post(componentBodyFor("gsgn:1:raid:", cardMessageFields()), time.Now(), nil))["data"].(map[string]any)
	embeds, _ = data["embeds"].([]any)
	first, _ = embeds[0].(map[string]any)
	if data["content"] != guideCopy.CardFailure || first["description"] != "Ready. Show BiS, or change a pick." {
		t.Fatalf("answer %v", data)
	}
	if compact(t, data["components"]) != compact(t, testCardComponents) {
		t.Fatalf("components changed: %v", data["components"])
	}
}

func TestCardFailureDeferredEditKeepsTheCard(t *testing.T) {
	h := cardHarness(t, "gsgn", func(context.Context, Interaction) (Reply, error) {
		time.Sleep(150 * time.Millisecond)
		return Reply{}, errors.New("boom")
	})
	h.svc.answerWithin = 20 * time.Millisecond
	h.post(componentBodyFor("gsgn:1:raid:", cardMessageFields()), time.Now(), nil)
	if err := h.svc.Drain(context.Background()); err != nil {
		t.Fatal(err)
	}
	calls := h.rest.Calls()
	if len(calls) != 1 || calls[0].Method != "EditOriginal" {
		t.Fatalf("calls %+v", calls)
	}
	edit := compact(t, calls[0].Message.edit())
	for _, part := range []string{`"content":"` + guideCopy.CardFailure + `"`, "Show BiS", `"allowed_mentions":{"parse":[]}`} {
		if !strings.Contains(edit, part) {
			t.Fatalf("edit lacks %s: %s", part, edit)
		}
	}
}

func TestNonCardFailureIsUnchanged(t *testing.T) {
	h := newHarness(t, failing)
	got := decode(t, h.post(componentBodyFor("upd:1", `"embeds":`+testCardEmbeds), time.Now(), nil))
	data, _ := got["data"].(map[string]any)
	if got["type"] != float64(7) || data["content"] != commonCopy.Failure || len(data["embeds"].([]any)) != 0 {
		t.Fatalf("answer %v", got)
	}
}
