package discordbot

import (
	"context"
	"encoding/json"
	"strings"
	"testing"
	"time"
)

func TestModalAfterDeferralIsAFailedPress(t *testing.T) {
	h := cardHarness(t, "gsgn", func(context.Context, Interaction) (Reply, error) {
		time.Sleep(150 * time.Millisecond)
		return Reply{Modal: &Modal{CustomID: "gsgn:1:name:", Title: "t"}}, nil
	})
	h.svc.answerWithin = 20 * time.Millisecond
	h.post(componentBodyFor("gsgn:1:raid:", cardMessageFields()), time.Now(), nil)
	if err := h.svc.Drain(context.Background()); err != nil {
		t.Fatal(err)
	}
	calls := h.rest.Calls()
	if len(calls) != 1 || calls[0].Method != "EditOriginal" || calls[0].Message.Content != guideCopy.CardFailure {
		t.Fatalf("calls %+v", calls)
	}
}

func TestAskPressSkipsTheServerLookup(t *testing.T) {
	h := cardHarness(t, "gchar", func(context.Context, Interaction) (Reply, error) {
		return Reply{Modal: &Modal{CustomID: "gchar:1:name:", Title: "t"}}, nil
	})
	body := strings.Replace(componentBodyFor("gchar:1:ask:", `"embeds":[]`), testServerID, "unbound-server", 1)
	if got := decode(t, h.post(body, time.Now(), nil)); got["type"] != float64(9) {
		t.Fatalf("answer %v", got)
	}
}

func submitWith(data string) Interaction {
	return Interaction{Type: typeModalSubmit, Data: json.RawMessage(data)}
}

func TestModalValueIsTrimmedAndBounded(t *testing.T) {
	in := submitWith(`{"components":[{"components":[{"custom_id":"name","value":"  Zul  "}]}]}`)
	if got := in.ModalValue("name"); got != "Zul" {
		t.Fatalf("trim %q", got)
	}
	long := strings.Repeat("é", maxModalValue+20)
	in = submitWith(`{"components":[{"components":[{"custom_id":"name","value":"` + long + `"}]}]}`)
	if got := []rune(in.ModalValue("name")); len(got) != maxModalValue {
		t.Fatalf("bound %d", len(got))
	}
}

func TestModalValueOfMalformedSubmitsIsEmpty(t *testing.T) {
	for name, data := range map[string]string{
		"missing components": `{}`,
		"wrong-typed":        `{"components":"x"}`,
		"wrong-typed rows":   `{"components":[{"components":5}]}`,
		"wrong-typed value":  `{"components":[{"components":[{"custom_id":"name","value":7}]}]}`,
		"empty value":        `{"components":[{"components":[{"custom_id":"name","value":""}]}]}`,
		"blank value":        `{"components":[{"components":[{"custom_id":"name","value":"   "}]}]}`,
		"non-json":           `not json`,
		"empty data":         ``,
	} {
		if got := submitWith(data).ModalValue("name"); got != "" {
			t.Errorf("%s: %q", name, got)
		}
	}
}
