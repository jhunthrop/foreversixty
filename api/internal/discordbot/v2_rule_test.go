package discordbot

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"os"
	"testing"
)

// legacyBisCard is the /bis card as the build before Components V2 drew it (the guided-card
// spec's bis-card-main payload, kept in testdata): a press arriving from it after deploy.
func legacyBisCard(t *testing.T) sourceMessage {
	t.Helper()
	raw, err := os.ReadFile("testdata/legacy-bis-card-main.json")
	if err != nil {
		t.Fatal(err)
	}
	var frame struct {
		Message sourceMessage `json:"message"`
	}
	if err := json.Unmarshal(raw, &frame); err != nil {
		t.Fatal(err)
	}
	if frame.Message.componentsV2() || len(frame.Message.Embeds) == 0 {
		t.Fatal("the legacy fixture is not a legacy card")
	}
	return frame.Message
}

// legacySource is a legacy message with an embed and Share in channel, as an answer drawn
// before Components V2.
func legacySource(t *testing.T, share string) sourceMessage {
	t.Helper()
	rows, err := json.Marshal([]ActionRow{{shareButton(share, false)}})
	if err != nil {
		t.Fatal(err)
	}
	return sourceMessage{Flags: ephemeralFlag, Embeds: json.RawMessage(`[{"title":"old"}]`), Components: rows}
}

// pressCallback runs a press through the answer path and returns the callback Discord is sent.
func (w *gearWorld) pressCallback(t *testing.T, in Interaction) (int, map[string]any) {
	t.Helper()
	in.ReceivedAt = boardClock
	h, update, ok := w.svc.route(in)
	if !ok {
		t.Fatalf("no route for %q", in.CustomID())
	}
	rec := httptest.NewRecorder()
	w.svc.answerOrDefer(context.Background(), rec, in, h, update)
	var cb struct {
		Type int            `json:"type"`
		Data map[string]any `json:"data"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &cb); err != nil {
		t.Fatal(err)
	}
	return cb.Type, cb.Data
}

// assertUpdateBody holds spec 2026-10-29 §A.7 rule 2 on one update body: the V2 flag always;
// over a legacy source content null and embeds empty; over a V2 source neither key.
func assertUpdateBody(t *testing.T, name string, body map[string]any, fromLegacy bool) {
	t.Helper()
	content, hasContent := body["content"]
	embeds, hasEmbeds := body["embeds"]
	if body["flags"] != float64(componentsV2Flag) {
		t.Errorf("%s: flags = %v", name, body["flags"])
	}
	if fromLegacy {
		if !hasContent || content != nil || !hasEmbeds || len(embeds.([]any)) != 0 {
			t.Errorf("%s over a legacy message: content %v (%v), embeds %v (%v)", name, content, hasContent, embeds, hasEmbeds)
		}
		return
	}
	if hasContent || hasEmbeds {
		t.Errorf("%s over a V2 message sends content %v or embeds %v", name, content, embeds)
	}
}

func TestEveryUpdateDecidesByItsSource(t *testing.T) {
	w := newGearWorld(t)
	p1 := w.bisPanel(t, gearBisOpts)
	share := p1.Components[0][0].CustomID[len(shareIDPrefix)+1:]
	v2Card := w.command(t, pressDiscord, "bis", nil).Message
	presses := []struct {
		name, id string
		values   []string
		legacy   sourceMessage
		v2       sourceMessage
	}{
		{"the /bis card's Show BiS", "gbis:1:go:hunter:beast-mastery:p1:", nil, legacyBisCard(t), sourceOf(t, v2Card)},
		{"a /char select pick", charPickID, []string{"Zulmara"}, legacySource(t, share), sourceOf(t, w.command(t, pressDiscord, "char", nil).Message)},
		{"a page press", "gbis:1:pg:hunter:beast-mastery:p1::2", nil, legacySource(t, share), sourceOf(t, p1)},
		{"Change picks", "gbis:1:chg:hunter:beast-mastery:p1:", nil, legacySource(t, share), sourceOf(t, p1)},
	}
	for _, p := range presses {
		for _, fromLegacy := range []bool{true, false} {
			data := map[string]any{"custom_id": p.id}
			if p.values != nil {
				data["values"] = p.values
			}
			in := interactionOf(typeComponent, pressDiscord, data)
			in.ReceivedAt, in.Source = boardClock, p.v2
			if fromLegacy {
				in.Source = p.legacy
			}
			kind, body := w.pressCallback(t, in)
			if kind != respUpdateMessage {
				t.Fatalf("%s: callback type %d", p.name, kind)
			}
			assertUpdateBody(t, p.name, body, fromLegacy)

			// The deferred path edits through the press's token with the same body.
			w.svc.deliver(in, handled{reply: w.route(t, in)}, true)
			calls := w.editsOriginal()
			assertUpdateBody(t, p.name+" (deferred)", wireOf(t, calls[len(calls)-1].Message.edit()), fromLegacy)
		}
	}
}

func TestEveryBisAndCharFrameIsComponentsV2(t *testing.T) {
	w := newGearWorld(t)
	frames := map[string]Message{
		"/bis card":          w.command(t, pressDiscord, "bis", nil).Message,
		"/bis page 1":        w.bisPanel(t, gearBisOpts),
		"/char":              w.command(t, pressDiscord, "char", nil).Message,
		"/char unknown":      w.command(t, pressDiscord, "char", map[string]string{"name": "Nobody"}).Message,
		"/char ask":          w.command(t, "d-stranger", "char", nil).Message,
		"a stale /bis card":  w.card(t, pressDiscord, "gbis:2:go:hunter:beast-mastery:p1:").Message,
		"a stale /char card": w.card(t, pressDiscord, "gchar:2:pick:", "Zulmara").Message,
		"/bis with no band":  w.card(t, pressDiscord, "gbis:1:go:hunter:marksmanship:p1:").Message,
	}
	for name, m := range frames {
		if !m.componentsV2() {
			t.Errorf("%s is not V2: %+v", name, m)
			continue
		}
		if f := wireOf(t, m)["flags"]; f != float64(componentsV2Flag|ephemeralFlag) {
			t.Errorf("%s: flags %v", name, f)
		}
		if n := m.componentCount(); n > maxV2Components {
			t.Errorf("%s: %d components", name, n)
		}
	}
	// The guided card keeps its layout: one container, four rows (spec §A.7 counts at most 14).
	if n := frames["/bis card"].componentCount(); n > 14 || len(frames["/bis card"].Components) != 4 {
		t.Errorf("the /bis card counts %d components", n)
	}
}

// A failed press on a V2 answer puts the answer back with the failure line on top, in V2.
func TestACardFailureOverAV2AnswerStaysV2(t *testing.T) {
	w := newGearWorld(t)
	p1 := w.bisPanel(t, gearBisOpts)
	in := interactionOf(typeComponent, pressDiscord, map[string]any{"custom_id": "gbis:1:pg:hunter:beast-mastery:p1::2"})
	in.Source = sourceOf(t, p1)
	body := wireOf(t, cardFailureReply(in).Message.edit())
	assertUpdateBody(t, "card failure", body, false)
	list := body["components"].([]any)
	if first := list[0].(map[string]any); first["type"] != float64(KindTextDisplay) || first["content"] != guideCopy.CardFailure ||
		len(list) != len(wireOf(t, p1)["components"].([]any))+1 {
		t.Fatalf("components = %v", list)
	}
	// A second failure replaces the line, never stacks it.
	raw, _ := json.Marshal(list)
	in.Source.Components = raw
	again := wireOf(t, cardFailureReply(in).Message.edit())["components"].([]any)
	if len(again) != len(list) {
		t.Fatalf("the failure line stacked: %d then %d", len(list), len(again))
	}
}
