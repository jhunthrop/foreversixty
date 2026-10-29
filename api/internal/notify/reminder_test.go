package notify

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/spectable"
)

const (
	guidedSpecPath    = "../../../design/specs/2026-10-28-discord-guided-cards.md"
	guidedCopyHeading = "## P.7 Copy, verbatim"
	reminderUnix      = "1796954400"
	reminderToken     = "reminder-secret-token"
)

func guidedCell(t *testing.T, key string) []string {
	t.Helper()
	cells := spectable.Cell(t, guidedSpecPath, guidedCopyHeading, "`"+key+"`")
	if len(cells) == 0 {
		t.Fatalf("spec has no string for %s", key)
	}
	return cells
}

// boldLead puts the clause before the first comma in bold, as §P.4 says the line is drawn.
func boldLead(line string) string {
	lead, rest, _ := strings.Cut(line, ",")
	return "**" + lead + "**," + rest
}

func TestReminderCopyMatchesTheSpec(t *testing.T) {
	title := guidedCell(t, "post.tomorrowTitle")[0]
	wantTitle := strings.NewReplacer("{Title}", "Onyxia's Lair", "{Weekday}", "Thursday", "{d}", "10", "{Month}", "December").Replace(title)
	e := RaidPostEmbed(RaidPost{RaidPosted: onyxiaThursday(), State: PostTomorrow}, fixedFinish)
	if e.Title != wantTitle {
		t.Fatalf("title = %q, want %q", e.Title, wantTitle)
	}
	if want := guidedCell(t, "post.tomorrowDescription")[0]; e.Description != want {
		t.Fatalf("description = %q, want %q", e.Description, want)
	}

	line := guidedCell(t, "reminder.line")
	fill := strings.NewReplacer("{Title}", "Onyxia's Lair", "{unix}", reminderUnix, "{signed}", "33", "{size}", "40", "{n}", "12")
	in := onyxiaThursday()
	in.Counts = RaidCounts{In: 30, Late: 3}
	if got, want := ReminderLine(in), boldLead(fill.Replace(line[0])); got != want {
		t.Fatalf("known size = %q, want %q", got, want)
	}
	in.Counts, in.SizeKnown = RaidCounts{In: 12}, false
	head := strings.TrimSuffix(fill.Replace(line[0]), " 33 of 40 signed up so far.")
	if got, want := ReminderLine(in), boldLead(head+" "+fill.Replace(line[1])); got != want {
		t.Fatalf("unknown size = %q, want %q", got, want)
	}
	in.Counts = RaidCounts{}
	if got, want := ReminderLine(in), boldLead(head+" "+line[2]); got != want {
		t.Fatalf("nobody = %q, want %q", got, want)
	}

	in = onyxiaThursday()
	link := strings.ReplaceAll(guidedCell(t, "reminder.webhookLink")[0], "{raid url}", in.RaidURL)
	if got, want := WebhookReminderContent(in), ReminderLine(in)+" "+link; got != want {
		t.Fatalf("webhook content = %q, want %q", got, want)
	}
}

func TestTomorrowEmbed(t *testing.T) {
	e := RaidPostEmbed(RaidPost{RaidPosted: onyxiaThursday(), State: PostTomorrow}, fixedFinish)
	open := RaidPostEmbed(RaidPost{RaidPosted: onyxiaThursday(), State: PostOpen}, fixedFinish)
	if e.Title != "Tomorrow: Onyxia's Lair · Thursday 10 December" ||
		e.Description != "The raid is tomorrow. Sign up with the buttons below." {
		t.Fatalf("embed = %q / %q", e.Title, e.Description)
	}
	if e.Color != open.Color || len(e.Fields) != len(open.Fields) || e.Fields[0] != open.Fields[0] {
		t.Fatalf("colour and fields are not open's: %+v", e)
	}
}

func TestReminderLine(t *testing.T) {
	in := onyxiaThursday()
	in.Counts = RaidCounts{In: 30, Late: 3}
	const lead = "**Onyxia's Lair is tomorrow**, <t:1796954400:F> (<t:1796954400:R>)."
	if got, want := ReminderLine(in), lead+" 33 of 40 signed up so far."; got != want {
		t.Fatalf("got %q want %q", got, want)
	}
	in.SizeKnown, in.Counts = false, RaidCounts{In: 12}
	if got, want := ReminderLine(in), lead+" 12 signed up so far."; got != want {
		t.Fatalf("got %q want %q", got, want)
	}
	in.Counts = RaidCounts{Bench: 2}
	if got, want := ReminderLine(in), lead+" Nobody has signed up yet."; got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}

func TestReminderLineSurvivesHostileTitles(t *testing.T) {
	in := onyxiaThursday()
	in.Title = "@everyone **x**"
	line := ReminderLine(in)
	if !strings.Contains(line, `**@everyone \*\*x\*\* is tomorrow**`) {
		t.Fatalf("title is not escaped: %q", line)
	}
	in.Title = strings.Repeat("*", 80)
	if n := len([]rune(WebhookReminderContent(in))); n > 2000 {
		t.Fatalf("content is %d runes", n)
	}
}

// reminderHook is a guild whose sealed webhook points at a recording httptest Discord.
type reminderHook struct {
	r       guildRig
	logs    *bytes.Buffer
	methods []string
	paths   []string
	queries []string
	bodies  []map[string]any
}

func newReminderHook(t *testing.T, deleteStatus int) *reminderHook {
	t.Helper()
	h := &reminderHook{r: newGuildRig(t), logs: &bytes.Buffer{}}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		raw, _ := io.ReadAll(req.Body)
		var body map[string]any
		_ = json.Unmarshal(raw, &body)
		h.methods, h.paths = append(h.methods, req.Method), append(h.paths, req.URL.Path)
		h.queries, h.bodies = append(h.queries, req.URL.RawQuery), append(h.bodies, body)
		if req.Method == http.MethodDelete {
			w.WriteHeader(deleteStatus)
			return
		}
		_, _ = w.Write([]byte(`{"id":"1234"}`))
	}))
	t.Cleanup(srv.Close)
	box, err := h.r.d.Sealer.Seal(srv.URL + "/api/webhooks/1/" + reminderToken)
	if err != nil {
		t.Fatal(err)
	}
	saved := h.r.hook(t)
	if err := h.r.store.Save(context.Background(), GuildSaveInput{
		GuildID: h.r.gid, ChannelID: saved.ChannelID, TokenBox: box, ConnectedBy: *saved.ConnectedBy,
	}); err != nil {
		t.Fatal(err)
	}
	h.r.d.Poster = NewHTTPPoster("")
	h.r.d.Log = slog.New(slog.NewTextHandler(h.logs, nil))
	return h
}

func assertNoToken(t *testing.T, logs *bytes.Buffer) {
	t.Helper()
	if strings.Contains(logs.String(), reminderToken) {
		t.Fatal("the log holds the webhook token")
	}
}

func TestWebhookReminderPostsAndDeletes(t *testing.T) {
	h := newReminderHook(t, http.StatusNotFound)
	ctx := context.Background()
	ref, sent, err := h.r.d.RaidReminder(ctx, h.r.gid, onyxiaThursday())
	if err != nil || !sent || ref != (ReminderRef{Via: ReminderViaWebhook, MessageID: "1234"}) {
		t.Fatalf("reminder = %+v, %v, %v", ref, sent, err)
	}
	if h.queries[0] != "wait=true" || h.methods[0] != http.MethodPost {
		t.Fatalf("post = %s ?%s", h.methods[0], h.queries[0])
	}
	body := h.bodies[0]
	// §P.9.7: the webhook post in the Tomorrow state rides under the line, with no components.
	embeds, _ := body["embeds"].([]any)
	if len(embeds) != 1 || embeds[0].(map[string]any)["title"] != "Tomorrow: Onyxia's Lair · Thursday 10 December" {
		t.Fatalf("a reminder carries one Tomorrow embed: %v", body["embeds"])
	}
	if _, has := body["components"]; has {
		t.Fatalf("a webhook reminder carries components: %v", body)
	}
	mentions, _ := json.Marshal(body["allowed_mentions"])
	if body["content"] != WebhookReminderContent(onyxiaThursday()) || string(mentions) != `{"parse":[]}` {
		t.Fatalf("body = %v", body)
	}

	h.r.d.RaidReminderWithdrawn(ctx, h.r.gid, ref) // the server answers 404: success
	if h.methods[1] != http.MethodDelete || h.paths[1] != "/api/webhooks/1/"+reminderToken+"/messages/1234" {
		t.Fatalf("delete = %s %s", h.methods[1], h.paths[1])
	}
	if strings.Contains(h.logs.String(), "withdraw_failed") {
		t.Fatalf("a 404 was logged as a failure: %s", h.logs)
	}
	assertNoToken(t, h.logs)
}

func TestWebhookReminderFailuresNeverLeakTheToken(t *testing.T) {
	h := newReminderHook(t, http.StatusInternalServerError)
	ctx := context.Background()
	h.r.d.RaidReminderWithdrawn(ctx, h.r.gid, ReminderRef{Via: ReminderViaWebhook, MessageID: "1234"})
	h.r.d.RaidReminderWithdrawn(ctx, h.r.gid, ReminderRef{Via: ReminderViaBot, MessageID: "1234"})
	if len(h.methods) != 1 || !strings.Contains(h.logs.String(), "withdraw_failed") {
		t.Fatalf("calls %v, logs %s", h.methods, h.logs)
	}
	assertNoToken(t, h.logs)
}

func TestRaidReminderSkipsWithoutAChannel(t *testing.T) {
	d := newGuildRig(t).d
	if _, sent, err := d.RaidReminder(context.Background(), 987654321, onyxiaThursday()); sent || err != nil {
		t.Fatalf("sent = %v, err = %v", sent, err)
	}
	var none *GuildDeliverer
	if _, sent, err := none.RaidReminder(context.Background(), 1, onyxiaThursday()); sent || err != nil {
		t.Fatalf("nil deliverer: sent = %v, err = %v", sent, err)
	}
}
