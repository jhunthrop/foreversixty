package notify

import (
	"context"
	"net/http"
	"net/url"
	"testing"
)

func TestWebhookConnectStillReturnsThroughTheSharedCallback(t *testing.T) {
	g := newGuildAPI(t, officerWithPlan())
	var linkCalls int
	g.svc.OnState(StateKindLink, func(http.ResponseWriter, *http.Request, ConnectState) { linkCalls++ })
	state := g.startGuildConnect()
	_, loc := g.callback("code="+authCode+"&state="+url.QueryEscape(state), nil)
	g.assertGuildOutcome(loc, outcomeConnected)
	if linkCalls != 0 {
		t.Fatal("a webhook state reached the link handler")
	}
}

func TestCallbackDispatchesAStateByKind(t *testing.T) {
	g := newGuildAPI(t, officerWithPlan())
	var got ConnectState
	g.svc.OnState(StateKindLink, func(w http.ResponseWriter, _ *http.Request, cs ConnectState) {
		got = cs
		w.WriteHeader(http.StatusNoContent)
	})
	state, err := g.svc.StartState(context.Background(), StateKindLink, g.uid, nil)
	if err != nil {
		t.Fatal(err)
	}
	res, _ := g.callback("code=x&state="+url.QueryEscape(state), nil)
	if res.StatusCode != http.StatusNoContent || got.Kind != StateKindLink || got.UserID != g.uid || got.GuildID != 0 {
		t.Fatalf("status %d, state %+v", res.StatusCode, got)
	}
	// Used once, and a kind with no handler fails to the account page.
	_, loc := g.callback("code=x&state="+url.QueryEscape(state), nil)
	g.assertOutcome(loc, outcomeFailed)
	other, err := g.svc.StartState(context.Background(), StateKindBotInstall, g.uid, &g.gid)
	if err != nil {
		t.Fatal(err)
	}
	_, loc = g.callback("code=x&state="+url.QueryEscape(other), nil)
	g.assertOutcome(loc, outcomeFailed)
}
