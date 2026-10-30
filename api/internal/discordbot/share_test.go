package discordbot

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

func (w *lookupWorld) pressShare(t *testing.T, customID string) Reply {
	t.Helper()
	in := interactionOf(typeComponent, pressDiscord, map[string]any{"custom_id": customID})
	in.Token = "tok-share"
	rep, err := w.svc.share(context.Background(), in)
	if err != nil {
		t.Fatal(err)
	}
	return rep
}

func (w *lookupWorld) followUps() []RESTCall {
	var out []RESTCall
	for _, c := range w.rest.Calls() {
		if c.Method == "FollowUp" {
			out = append(out, c)
		}
	}
	return out
}

func TestShareRecomputesAsAVisitorOnce(t *testing.T) {
	w := newLookupWorld(t)
	w.port.candidates = []guilds.SignupCandidate{{CharacterKey: keyZulmaraLookup, Name: "Zulmara", Verified: true, Main: true}}
	w.lg.readiness[keyZulmaraLookup] = guilds.ReadinessFacts{Word: "Stocked", SyncedAt: ref(time.Now()), OnRoster: true}
	answer := w.command(t, pressDiscord, "char", map[string]string{"name": "Zulmara"}).Message
	shareID := answer.Components[0][0].CustomID
	if !strings.HasPrefix(shareID, "share:") || len(shareID) > 100 {
		t.Fatalf("share id = %q", shareID)
	}

	rep := w.pressShare(t, shareID)
	if !rep.Update || rep.After == nil {
		t.Fatalf("first press = %+v", rep)
	}
	if b := rep.Message.Components[0][0]; b.Label != "Shared" || !b.Disabled || b.CustomID != shareID {
		t.Fatalf("button after = %+v", b)
	}
	if len(w.followUps()) != 0 {
		t.Fatal("the follow-up went before the press was answered")
	}
	rep.After(context.Background())
	ups := w.followUps()
	if len(ups) != 1 || ups[0].Token != "tok-share" {
		t.Fatalf("follow-ups = %+v", ups)
	}
	public := ups[0].Message
	raw, _ := json.Marshal(public)
	if public.Ephemeral || wireOf(t, public)["flags"] != float64(componentsV2Flag) {
		t.Fatalf("follow-up is not public: %s", raw)
	}
	if got := headOf(t, public).Description; strings.Contains(got, "Readiness") || strings.Contains(got, "Synced") {
		t.Fatalf("the visitor's card shows member clauses: %s", got)
	}
	for _, row := range public.Components {
		for _, c := range row {
			if c.Style != StyleLink {
				t.Fatalf("the public answer carries a button that acts: %+v", c)
			}
		}
	}

	again := w.pressShare(t, shareID)
	if !again.Update || again.After != nil || again.Message.Components[0][0].Label != "Shared" {
		t.Fatalf("second press = %+v", again)
	}
	if len(w.followUps()) != 1 {
		t.Fatal("a second press shared again")
	}

	fresh := w.command(t, pressDiscord, "char", map[string]string{"name": "Zulmara"}).Message
	delete(w.cards.cards, "zulmara") // the character turned private since
	gone := w.pressShare(t, fresh.Components[0][0].CustomID)
	if gone.Update || contentOf(gone.Message) != "That is no longer public, so it was not shared." || !gone.Message.Ephemeral {
		t.Fatalf("private since = %+v", gone)
	}
	if len(w.followUps()) != 1 {
		t.Fatal("a private character was shared")
	}
}

func TestShareRefusesAnotherGuildsOrAnExpiredToken(t *testing.T) {
	w := newLookupWorld(t)
	answer := w.command(t, pressDiscord, "bis", map[string]string{"spec": "hunter-marksmanship", "phase": "60"}).Message
	id := answer.Components[0][0].CustomID
	w.exec(t, `update discord_shares set expires_at = now() - interval '1 minute'`)
	if rep := w.pressShare(t, id); rep.Message.Content != commonCopy.Failure {
		t.Fatalf("expired = %+v", rep)
	}
	other := seedGuild(t, w.st.Pool, "Other")
	w.exec(t, `update discord_shares set expires_at = now() + interval '1 hour', guild_id = $1`, other)
	if rep := w.pressShare(t, id); rep.Message.Content != commonCopy.Failure {
		t.Fatalf("another guild's token = %+v", rep)
	}
}

// TestReplyAfterRunsOnceTheAnswerIsOut: inline, After runs after the answer is written;
// deferred, after the deferred answer's edit.
func TestReplyAfterRunsOnceTheAnswerIsOut(t *testing.T) {
	for _, c := range []struct {
		name  string
		delay time.Duration
		first string
	}{{"inline", time.Millisecond, "FollowUp"}, {"deferred", 200 * time.Millisecond, "EditOriginal"}} {
		t.Run(c.name, func(t *testing.T) {
			var h *harness
			h = newHarness(t, func(ctx context.Context, in Interaction) (Reply, error) {
				time.Sleep(c.delay)
				return Reply{Message: Message{Content: "done", Ephemeral: true}, After: func(ctx context.Context) {
					_ = h.rest.FollowUp(ctx, in.Token, Message{Content: "public"})
				}}, nil
			})
			h.svc.answerWithin = 50 * time.Millisecond
			h.post(commandBody("spy"), time.Now(), nil)
			if err := h.svc.Drain(context.Background()); err != nil {
				t.Fatal(err)
			}
			calls := h.rest.Calls()
			if len(calls) == 0 || calls[0].Method != c.first || calls[len(calls)-1].Method != "FollowUp" {
				t.Fatalf("calls %+v", calls)
			}
		})
	}
}

func TestShareIsThePressersAlone(t *testing.T) {
	w := newLookupWorld(t)
	id := w.command(t, pressDiscord, "char", map[string]string{"name": "Zulmara"}).Message.Components[0][0].CustomID
	in := interactionOf(typeComponent, "d-someone-else", map[string]any{"custom_id": id})
	rep, err := w.svc.share(context.Background(), in)
	if err != nil || rep.Message.Content != commonCopy.Failure || rep.After != nil {
		t.Fatalf("another presser = %+v, %v", rep, err)
	}
	if again := w.pressShare(t, id); again.After == nil {
		t.Fatal("the answer's own presser could not share after a refused press")
	}
}

// TestADroppedShareAnswerFreesTheToken: when the Shared answer never reaches Discord (the
// deferred edit fails, the token expired, or shutdown refused the work), After runs marked
// dropped, posts nothing and frees the token.
func TestADroppedShareAnswerFreesTheToken(t *testing.T) {
	w := newLookupWorld(t)
	id := w.command(t, pressDiscord, "char", map[string]string{"name": "Zulmara"}).Message.Components[0][0].CustomID
	rep := w.pressShare(t, id)
	w.rest.ErrFor = map[string]error{"EditOriginal": &APIError{Status: 500, Message: "boom"}}
	in := interactionOf(typeComponent, pressDiscord, map[string]any{"custom_id": id})
	in.ReceivedAt = time.Now()
	w.svc.deliver(in, handled{reply: rep}, true)
	if len(w.followUps()) != 0 {
		t.Fatal("a dropped answer still shared")
	}
	w.rest.ErrFor = nil
	if again := w.pressShare(t, id); again.After == nil || again.Message.Components[0][0].Label != "Shared" {
		t.Fatalf("the token was not freed: %+v", again)
	}
	w.exec(t, `update discord_shares set used_at = null`)
	rep = w.pressShare(t, id)
	in.ReceivedAt = time.Now().Add(-interactionTokenLife - time.Minute)
	w.svc.deliver(in, handled{reply: rep}, true)
	var used bool
	if err := w.st.Pool.QueryRow(context.Background(), `select used_at is not null from discord_shares`).Scan(&used); err != nil || used {
		t.Fatalf("an expired answer kept its claim: %v %v", used, err)
	}
}

func TestAShareRefusedAtShutdownFreesTheToken(t *testing.T) {
	w := newLookupWorld(t)
	id := w.command(t, pressDiscord, "char", map[string]string{"name": "Zulmara"}).Message.Components[0][0].CustomID
	rep := w.pressShare(t, id)
	if err := w.svc.Drain(context.Background()); err != nil {
		t.Fatal(err)
	}
	w.svc.runAfter(httptest.NewRecorder(), Interaction{ID: "i1"}, rep)
	var used bool
	if err := w.st.Pool.QueryRow(context.Background(), `select used_at is not null from discord_shares`).Scan(&used); err != nil || used {
		t.Fatalf("a share refused at shutdown kept its claim: %v %v", used, err)
	}
	if len(w.followUps()) != 0 {
		t.Fatal("a share refused at shutdown still posted")
	}
}

func TestSweepSharesDeletesADayPastExpiry(t *testing.T) {
	w := newLookupWorld(t)
	w.command(t, pressDiscord, "char", map[string]string{"name": "Zulmara"})
	w.command(t, pressDiscord, "char", map[string]string{"name": "Zulmara"})
	w.exec(t, `update discord_shares set expires_at = now() - interval '25 hours'
		where token = (select token from discord_shares limit 1)`)
	if err := w.st.SweepShares(context.Background(), time.Now()); err != nil {
		t.Fatal(err)
	}
	var left int
	if err := w.st.Pool.QueryRow(context.Background(), `select count(*) from discord_shares`).Scan(&left); err != nil || left != 1 {
		t.Fatalf("left = %d, %v", left, err)
	}
}

func TestShareFollowUpFailureFreesTheToken(t *testing.T) {
	w := newLookupWorld(t)
	answer := w.command(t, pressDiscord, "char", map[string]string{"name": "Zulmara"}).Message
	id := answer.Components[0][0].CustomID
	w.rest.ErrFor = map[string]error{"FollowUp": &APIError{Status: 500, Message: "boom"}}
	rep := w.pressShare(t, id)
	rep.After(context.Background())
	var edits []RESTCall
	for _, c := range w.rest.Calls() {
		if c.Method == "EditOriginal" {
			edits = append(edits, c)
		}
	}
	if len(edits) != 1 || edits[0].Message.Components[0][0].Label != lookupCopy.Share || edits[0].Message.Components[0][0].Disabled ||
		contentOf(edits[0].Message) != commonCopy.Failure {
		t.Fatalf("restore edit = %+v", edits)
	}
	// The restore edits the V2 Shared answer: the flag, never content or embeds (spec §A.7).
	body := wireOf(t, edits[0].Message.edit())
	if _, has := body["content"]; has || body["flags"] != float64(componentsV2Flag) {
		t.Fatalf("restore body = %v", body)
	}
	if _, has := body["embeds"]; has {
		t.Fatalf("restore body = %v", body)
	}
	w.rest.ErrFor = nil
	if again := w.pressShare(t, id); again.After == nil {
		t.Fatal("the freed token did not share on the next press")
	}
}
