package discordbot

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"strings"
	"testing"
	"time"
)

func slog_(w io.Writer) *slog.Logger { return slog.New(slog.NewTextHandler(w, nil)) }

func slowHandler(d time.Duration) Handler {
	return func(context.Context, Interaction) (Reply, error) {
		time.Sleep(d)
		return Reply{Message: Message{Content: "done", Ephemeral: true}}, nil
	}
}

func TestRespondAnswersInline(t *testing.T) {
	h := newHarness(t, slowHandler(time.Millisecond))
	h.svc.answerWithin = 500 * time.Millisecond
	got := decode(t, h.post(commandBody("spy"), time.Now(), nil))
	data, _ := got["data"].(map[string]any)
	if got["type"] != float64(4) || data["flags"] != float64(64) || data["content"] != "done" {
		t.Fatalf("answer %v", got)
	}
}

func TestRespondDefersThenEditsOriginal(t *testing.T) {
	h := newHarness(t, slowHandler(200*time.Millisecond))
	h.svc.answerWithin = 50 * time.Millisecond
	rec := h.post(commandBody("spy"), time.Now(), nil)
	if strings.TrimSpace(rec.Body.String()) != `{"type":5,"data":{"flags":64}}` {
		t.Fatalf("deferral %s", rec.Body)
	}
	if err := h.svc.Drain(context.Background()); err != nil {
		t.Fatal(err)
	}
	calls := h.rest.Calls()
	if len(calls) != 1 || calls[0].Method != "EditOriginal" || calls[0].Token != testToken || calls[0].Message.Content != "done" {
		t.Fatalf("calls %+v", calls)
	}
}

func TestRespondDefersComponentUpdatesAsType6(t *testing.T) {
	h := newHarness(t, slowHandler(200*time.Millisecond))
	h.svc.answerWithin = 50 * time.Millisecond
	body := strings.Replace(commandBody("x"), `"type":2`, `"type":3`, 1)
	body = strings.Replace(body, `"name":"x"`, `"custom_id":"upd:1"`, 1)
	rec := h.post(body, time.Now(), nil)
	if strings.TrimSpace(rec.Body.String()) != `{"type":6}` {
		t.Fatalf("deferral %s", rec.Body)
	}
	_ = h.svc.Drain(context.Background())
}

func TestRespondHandlerErrorAnswersFailure(t *testing.T) {
	h := newHarness(t, func(context.Context, Interaction) (Reply, error) { return Reply{}, errors.New("boom") })
	got := decode(t, h.post(commandBody("spy"), time.Now(), nil))
	data, _ := got["data"].(map[string]any)
	if data["content"] != commonCopy.Failure || data["flags"] != float64(64) {
		t.Fatalf("answer %v", got)
	}
}

func TestRespondDropsAnswersPastTokenLife(t *testing.T) {
	h := newHarness(t, nil)
	in := Interaction{ID: "i9", Token: testToken, ReceivedAt: time.Now().Add(-interactionTokenLife - time.Second)}
	h.svc.deliver(in, handled{reply: Reply{Message: Message{Content: "late"}}})
	if len(h.rest.Calls()) != 0 {
		t.Fatal("an expired answer was sent")
	}
	logged := h.log.String()
	if !strings.Contains(logged, "outcome=expired") || strings.Contains(logged, testToken) {
		t.Fatalf("log %q", logged)
	}
}
