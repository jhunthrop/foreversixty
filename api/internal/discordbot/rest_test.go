package discordbot

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

type seen struct {
	method, path, auth string
	body               map[string]any
}

func serve(t *testing.T, status int, answer string) (*HTTPREST, *seen) {
	t.Helper()
	got := &seen{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		raw, _ := io.ReadAll(r.Body)
		got.method, got.path, got.auth = r.Method, r.URL.EscapedPath(), r.Header.Get("Authorization")
		got.body = nil
		_ = json.Unmarshal(raw, &got.body)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = w.Write([]byte(answer))
	}))
	t.Cleanup(srv.Close)
	rest := NewHTTPREST("app1", "t0k")
	rest.BaseURL = srv.URL
	return rest, got
}

func TestCreateMessageSendsBotAuthAndPingsNoOne(t *testing.T) {
	rest, got := serve(t, 200, `{"id":"m1","channel_id":"c1"}`)
	id, err := rest.CreateMessage(context.Background(), "c1", Message{Content: "hi"})
	if err != nil || id != "m1" {
		t.Fatalf("id=%q err=%v", id, err)
	}
	if got.method != "POST" || got.path != "/channels/c1/messages" || got.auth != "Bot t0k" {
		t.Fatalf("saw %+v", got)
	}
	parse, ok := got.body["allowed_mentions"].(map[string]any)["parse"].([]any)
	if !ok || len(parse) != 0 {
		t.Fatalf("body = %v", got.body)
	}
}

func TestEditMessageAndLeaveServer(t *testing.T) {
	rest, got := serve(t, 204, ``)
	if err := rest.EditMessage(context.Background(), "c1", "m9", Message{Content: "x"}); err != nil {
		t.Fatal(err)
	}
	if got.method != "PATCH" || got.path != "/channels/c1/messages/m9" {
		t.Fatalf("saw %+v", got)
	}
	if err := rest.LeaveServer(context.Background(), "srv"); err != nil {
		t.Fatal(err)
	}
	if got.method != "DELETE" || got.path != "/users/@me/guilds/srv" {
		t.Fatalf("saw %+v", got)
	}
}

func TestEditOriginalUsesTheInteractionWebhookWithoutBotAuth(t *testing.T) {
	rest, got := serve(t, 200, `{}`)
	if err := rest.EditOriginal(context.Background(), "itok", Message{Content: "x"}); err != nil {
		t.Fatal(err)
	}
	if got.method != "PATCH" || got.path != "/webhooks/app1/itok/messages/@original" || got.auth != "" {
		t.Fatalf("saw %+v", got)
	}
	if err := rest.FollowUp(context.Background(), "itok", Message{Content: "x"}); err != nil {
		t.Fatal(err)
	}
	if got.method != "POST" || got.path != "/webhooks/app1/itok" || got.auth != "" {
		t.Fatalf("saw %+v", got)
	}
}

func TestAPIErrorClassifies(t *testing.T) {
	cases := []struct {
		status int
		answer string
		check  func(*APIError) bool
		retry  time.Duration
	}{
		{429, `{"message":"You are being rate limited.","retry_after":1.5,"global":false}`, (*APIError).RateLimited, 1500 * time.Millisecond},
		{403, `{"message":"Missing Access","code":50001}`, (*APIError).MissingAccess, 0},
		{404, `{"message":"Unknown Message","code":10008}`, (*APIError).Gone, 0},
	}
	for _, c := range cases {
		rest, _ := serve(t, c.status, c.answer)
		_, err := rest.CreateMessage(context.Background(), "c1", Message{})
		apiErr, ok := err.(*APIError)
		if !ok || !c.check(apiErr) || apiErr.RetryAfter != c.retry || apiErr.Status != c.status {
			t.Fatalf("status %d: err = %#v", c.status, err)
		}
		if strings.Contains(err.Error(), "t0k") {
			t.Fatalf("error leaks the token: %v", err)
		}
	}
	if !(&APIError{Status: 404, Code: 10015}).Gone() || !(&APIError{Status: 401, Code: 50027}).Gone() || !(&APIError{Status: 404, Code: 10003}).Gone() || !(&APIError{Status: 400, Code: 50013}).MissingAccess() ||
		!(&APIError{Status: 404, Code: 10004}).MissingAccess() || (&APIError{Status: 500}).Gone() {
		t.Fatal("classification table wrong")
	}
}

func TestTransportErrorNeverLeaksTheToken(t *testing.T) {
	rest := NewHTTPREST("app1", "t0k")
	rest.BaseURL = "http://127.0.0.1:1"
	_, err := rest.CreateMessage(context.Background(), "c1", Message{})
	if err == nil || strings.Contains(err.Error(), "t0k") {
		t.Fatalf("err = %v", err)
	}
}

func TestPutCommandsGlobalAndServerScoped(t *testing.T) {
	rest, got := serve(t, 200, `[]`)
	cmds := []Command{{Name: "bis", Description: "d", Released: true}}
	if err := rest.PutCommands(context.Background(), "", cmds); err != nil {
		t.Fatal(err)
	}
	if got.method != "PUT" || got.path != "/applications/app1/commands" || got.auth != "Bot t0k" {
		t.Fatalf("saw %+v", got)
	}
	if err := rest.PutCommands(context.Background(), "srv", cmds); err != nil {
		t.Fatal(err)
	}
	if got.path != "/applications/app1/guilds/srv/commands" {
		t.Fatalf("saw %+v", got)
	}
}

func TestServerLayoutReadsChannelsRolesAndTheBotMember(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/guilds/srv/channels":
			_, _ = w.Write([]byte(`[{"id":"1","name":"general","type":0,"position":2,"parent_id":null,"permission_overwrites":[{"id":"srv","type":0,"allow":"0","deny":"2048"}]}]`))
		case "/guilds/srv/roles":
			_, _ = w.Write([]byte(`[{"id":"srv","name":"@everyone","permissions":"1024","position":0}]`))
		case "/guilds/srv/members/app1":
			_, _ = w.Write([]byte(`{"roles":["r1","r2"]}`))
		default:
			w.WriteHeader(404)
		}
	}))
	defer srv.Close()
	rest := NewHTTPREST("app1", "t0k")
	rest.BaseURL = srv.URL
	l, err := rest.ServerLayout(context.Background(), "srv")
	if err != nil {
		t.Fatal(err)
	}
	if len(l.Channels) != 1 || l.Channels[0].Overwrites[0].Deny != PermSendMessages || l.Roles[0].Permissions != PermViewChannel ||
		strings.Join(l.BotRoleIDs, ",") != "r1,r2" || l.BotUserID != "app1" {
		t.Fatalf("layout = %+v", l)
	}
}

func TestMessageJSON(t *testing.T) {
	m := Message{
		Content: "c", Ephemeral: true, Embeds: []notify.Embed{{Title: "t", Author: "a"}},
		Components: []ActionRow{{
			{Kind: KindButton, Style: StyleSuccess, Label: "Go", CustomID: "go"},
			{Kind: KindButton, Style: StyleLink, Label: "Site", URL: "https://x.test"},
		}, {
			{Kind: KindStringSelect, CustomID: "pick", Placeholder: "p", Options: []SelectOption{{Label: "l", Value: "v", Description: "d"}}},
		}},
	}
	raw, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	var out map[string]any
	_ = json.Unmarshal(raw, &out)
	if out["flags"] != float64(64) || out["allowed_mentions"] == nil {
		t.Fatalf("json = %s", raw)
	}
	row := out["components"].([]any)[0].(map[string]any)
	if row["type"] != float64(1) {
		t.Fatalf("row = %v", row)
	}
	parts := row["components"].([]any)
	selRow := out["components"].([]any)[1].(map[string]any)["components"].([]any)
	btn := parts[0].(map[string]any)
	if btn["type"] != float64(2) || btn["style"] != float64(3) || btn["custom_id"] != "go" {
		t.Fatalf("button = %v", btn)
	}
	link := parts[1].(map[string]any)
	if link["url"] != "https://x.test" || link["custom_id"] != nil {
		t.Fatalf("link = %v", link)
	}
	sel := selRow[0].(map[string]any)
	if sel["type"] != float64(3) || sel["options"] == nil || sel["style"] != nil {
		t.Fatalf("select = %v", sel)
	}
	plain, _ := json.Marshal(Message{Content: "c"})
	if strings.Contains(string(plain), "flags") || strings.Contains(string(plain), "components") {
		t.Fatalf("plain = %s", plain)
	}
}

func TestCommandJSON(t *testing.T) {
	raw, _ := json.Marshal(Command{Name: "bis", Description: "d", Released: true,
		Options: []CommandOption{{Name: "class", Description: "c", Type: 3, Required: true, Autocomplete: true, Choices: []Choice{{Name: "A", Value: "a"}}}}})
	var out map[string]any
	_ = json.Unmarshal(raw, &out)
	if out["released"] != nil || out["Released"] != nil || out["contexts"].([]any)[0] != float64(0) || out["integration_types"].([]any)[0] != float64(0) {
		t.Fatalf("json = %s", raw)
	}
	opt := out["options"].([]any)[0].(map[string]any)
	if opt["required"] != true || opt["autocomplete"] != true || opt["choices"] == nil {
		t.Fatalf("option = %v", opt)
	}
}

func TestEditsClearWhatTheyLeaveEmptyAndNeverSendFlags(t *testing.T) {
	rest, got := serve(t, 200, `{}`)
	m := Message{Ephemeral: true}
	if err := rest.EditMessage(context.Background(), "c1", "m1", m); err != nil {
		t.Fatal(err)
	}
	assertCleared := func() {
		t.Helper()
		for _, k := range []string{"embeds", "components"} {
			if v, ok := got.body[k].([]any); !ok || len(v) != 0 {
				t.Fatalf("%s = %v in %v", k, got.body[k], got.body)
			}
		}
		if c, ok := got.body["content"]; !ok || c != "" || got.body["flags"] != nil {
			t.Fatalf("body = %v", got.body)
		}
	}
	assertCleared()
	if err := rest.EditOriginal(context.Background(), "tok", m); err != nil {
		t.Fatal(err)
	}
	assertCleared()
}

func TestRetryAfterFallsBackToTheHeaderOnANonJSON429(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Retry-After", "2")
		w.WriteHeader(429)
		_, _ = w.Write([]byte("<html>blocked</html>"))
	}))
	defer srv.Close()
	rest := NewHTTPREST("app1", "t0k")
	rest.BaseURL = srv.URL
	_, err := rest.CreateMessage(context.Background(), "c1", Message{})
	apiErr, ok := err.(*APIError)
	if !ok || !apiErr.RateLimited() || apiErr.RetryAfter != 2*time.Second {
		t.Fatalf("err = %#v", err)
	}
}

func TestInvalidButtonsFailFast(t *testing.T) {
	bad := []Component{
		{Kind: KindButton, Style: StyleSuccess, Label: "no id"},
		{Kind: KindButton, Style: StyleLink, Label: "no url"},
		{Kind: KindButton, Style: StyleLink, Label: "both", URL: "https://x.test", CustomID: "x"},
		{Kind: KindButton, Style: StyleSuccess, Label: "url", CustomID: "x", URL: "https://x.test"},
	}
	for _, c := range bad {
		if _, err := json.Marshal(Message{Components: []ActionRow{{c}}}); err == nil {
			t.Fatalf("%q marshalled", c.Label)
		}
	}
}
