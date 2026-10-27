package discordbot

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"strings"
	"sync"
	"sync/atomic"
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
	if err := h.svc.Drain(context.Background()); err != nil {
		t.Fatal(err)
	}
	calls := h.rest.Calls()
	if len(calls) != 1 || calls[0].Method != "EditOriginal" || calls[0].Message.Content != "done" {
		t.Fatalf("calls %+v", calls)
	}
}

func TestRespondHandlerPanicAnswersFailureAndLogsOnlyTheType(t *testing.T) {
	h := newHarness(t, func(context.Context, Interaction) (Reply, error) { panic("secret member text") })
	got := decode(t, h.post(commandBody("spy"), time.Now(), nil))
	data, _ := got["data"].(map[string]any)
	if data["content"] != commonCopy.Failure {
		t.Fatalf("answer %v", got)
	}
	if logged := h.log.String(); !strings.Contains(logged, "panicked: string") || strings.Contains(logged, "secret member text") {
		t.Fatalf("log %q", logged)
	}
}

func TestAutocompletePanicAnswersNoChoices(t *testing.T) {
	h := newHarness(t, nil)
	h.svc.autocompletes = map[string]Autocomplete{"spy": func(context.Context, Interaction) ([]Choice, error) { panic("boom") }}
	body := `{"id":"i","type":4,"token":"t","guild_id":"` + testServerID + `","data":{"name":"spy"}}`
	rec := h.post(body, time.Now(), nil)
	data, _ := decode(t, rec)["data"].(map[string]any)
	if choices, _ := data["choices"].([]any); rec.Code != 200 || len(choices) != 0 {
		t.Fatalf("body %s", rec.Body)
	}
}

func TestBackgroundDrainTimesOutOnSlowWork(t *testing.T) {
	bg := &inProcess{}
	release := make(chan struct{})
	bg.Start(func() { <-release })
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	if err := bg.Drain(ctx); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("drain = %v", err)
	}
	close(release)
}

func TestBackgroundRefusesWorkOnceDraining(t *testing.T) {
	bg := &inProcess{}
	if err := bg.Drain(context.Background()); err != nil {
		t.Fatal(err)
	}
	ran := false
	if bg.Start(func() { ran = true }) || ran {
		t.Fatal("work started after Drain")
	}
}

func TestRespondAfterDrainAnswersFailure(t *testing.T) {
	h := newHarness(t, nil)
	if err := h.svc.Drain(context.Background()); err != nil {
		t.Fatal(err)
	}
	data, _ := decode(t, h.post(commandBody("spy"), time.Now(), nil))["data"].(map[string]any)
	if data["content"] != commonCopy.Failure || h.calls.Load() != 0 {
		t.Fatalf("data %v calls %d", data, h.calls.Load())
	}
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
	h.svc.deliver(in, handled{reply: Reply{Message: Message{Content: "late"}}}, false)
	if len(h.rest.Calls()) != 0 {
		t.Fatal("an expired answer was sent")
	}
	logged := h.log.String()
	if !strings.Contains(logged, "outcome=expired") || strings.Contains(logged, testToken) {
		t.Fatalf("log %q", logged)
	}
}

// firstOnly is a Background that runs its first work and refuses the rest, as a shutdown
// that began while a handler ran.
type firstOnly struct {
	mu      sync.Mutex
	started bool
}

func (b *firstOnly) Start(work func()) bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.started {
		return false
	}
	b.started = true
	go work()
	return true
}

func (b *firstOnly) Drain(context.Context) error { return nil }

// A deferral whose delivery shutdown refused still undoes what the handler holds: After runs
// once the handler finishes, marked dropped, and nothing reaches Discord.
func TestRespondRefusedDeferralRunsAfterDropped(t *testing.T) {
	var dropped atomic.Int32
	h := newHarness(t, func(context.Context, Interaction) (Reply, error) {
		time.Sleep(100 * time.Millisecond)
		return Reply{Message: Message{Content: "shared", Ephemeral: true}, After: func(ctx context.Context) {
			if answerDropped(ctx) {
				dropped.Add(1)
			}
		}}, nil
	})
	h.svc.bg = &firstOnly{}
	h.svc.answerWithin = 20 * time.Millisecond
	rec := h.post(commandBody("spy"), time.Now(), nil)
	if strings.TrimSpace(rec.Body.String()) != `{"type":5,"data":{"flags":64}}` {
		t.Fatalf("deferral %s", rec.Body)
	}
	if dropped.Load() != 1 || len(h.rest.Calls()) != 0 {
		t.Fatalf("dropped After runs = %d, calls %+v", dropped.Load(), h.rest.Calls())
	}
}
