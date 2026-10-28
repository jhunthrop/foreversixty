package discordbot

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

const coachURL = "https://forever.test/coach/abc"

type fakeCoach struct {
	lines  []CoachLine
	err    error
	users  []int64
	budget time.Duration
}

func (f *fakeCoach) NightLinesFor(_ context.Context, userID int64, _ string, budget time.Duration) ([]CoachLine, error) {
	f.users = append(f.users, userID)
	f.budget = budget
	return f.lines, f.err
}

func coachWorld(t *testing.T, lines ...CoachLine) (*officerWorld, *fakeCoach) {
	t.Helper()
	w := newOfficerWorld(t)
	fc := &fakeCoach{lines: lines}
	w.svc.d.Coach = fc
	w.addChar(t, w.gid, w.uid, "us/normal/zulmara", "Zulmara", true)
	w.og.report = onyxiaReport(false, false)
	return w, fc
}

func line(name, stem string, on, of int) CoachLine {
	return CoachLine{Character: name, Stem: stem, On: on, Of: of, URL: coachURL}
}

func TestCoachLineTextVariants(t *testing.T) {
	open := ". [Open your coach's review](" + coachURL + ")"
	for name, tc := range map[string]struct {
		in   CoachLine
		want string
	}{
		"several": {line("A", "Bloodthirst casts lost", 5, 8), "Fix first: Bloodthirst casts lost — on 5 of 8 bosses" + open},
		"one":     {line("A", "Bloodthirst casts lost", 1, 1), "Fix first: Bloodthirst casts lost" + open},
		"all":     {line("A", "Bloodthirst casts lost", 8, 8), "Fix first: Bloodthirst casts lost — on all 8 bosses" + open},
		"pending": {CoachLine{Character: "A", Pending: true, URL: coachURL}, "[Your coach's review is on the site.](" + coachURL + ")"},
		"nocount": {line("A", "Stem", 0, 0), "Fix first: Stem" + open},
	} {
		if got := coachLineText(tc.in); got != tc.want {
			t.Errorf("%s = %q, want %q", name, got, tc.want)
		}
	}
}

func TestReport_PremiumOwnerGetsOneLinePerCharacter(t *testing.T) {
	w, fc := coachWorld(t, line("Zulmara", zulmaraStem, 5, 8), line("Vexlash", "Early pulls", 8, 8))
	rep := w.run(t, pressDiscord, "report")
	got := headFields(rep)[coachCopy.FieldName]
	want := "**Zulmara** · Fix first: " + zulmaraStem + " — on 5 of 8 bosses. [Open your coach's review](" + coachURL + ")\n" +
		"**Vexlash** · Fix first: Early pulls — on all 8 bosses. [Open your coach's review](" + coachURL + ")"
	if got != want {
		t.Fatalf("field = %q", got)
	}
	if len(fc.users) != 1 || fc.users[0] != w.uid || fc.budget <= 0 || fc.budget > coachLinesBudget {
		t.Fatalf("coach asked %v with %v", fc.users, fc.budget)
	}
}

func TestReport_CoachFieldSitsBeforeLoot(t *testing.T) {
	w, _ := coachWorld(t, line("Zulmara", "Stem", 3, 8))
	fields := w.run(t, pressDiscord, "report").Message.Embeds[0].Fields
	names := make([]string, len(fields))
	for i, f := range fields {
		names[i] = f.Name
	}
	coach := slices.Index(names, coachCopy.FieldName)
	loot := slices.Index(names, officerCopy.Loot)
	kills := slices.Index(names, officerCopy.KillOrder)
	if coach < 0 || loot < 0 || coach != loot-1 || kills < 0 || kills > coach {
		t.Fatalf("fields = %q: the coach field must follow Kill order and sit directly before Loot", names)
	}
}

func TestCoachFieldIndex_FallsBackToAfterKillOrder(t *testing.T) {
	kills := []notify.EmbedField{{Name: "Other"}, {Name: officerCopy.KillOrder}, {Name: "Tail"}}
	if got := coachFieldIndex(kills); got != 2 {
		t.Fatalf("index without Loot = %d, want 2", got)
	}
	if got := coachFieldIndex([]notify.EmbedField{{Name: "Other"}}); got != 1 {
		t.Fatalf("index without Loot or Kill order = %d, want 1", got)
	}
}

func TestReport_OneCharacterIsNotNamed(t *testing.T) {
	w, _ := coachWorld(t, line("Zulmara", "Stem", 3, 8))
	got := headFields(w.run(t, pressDiscord, "report"))[coachCopy.FieldName]
	if strings.Contains(got, "Zulmara") || !strings.HasPrefix(got, "Fix first: Stem") {
		t.Fatalf("field = %q", got)
	}
}

func TestReport_AtMostThreeLines(t *testing.T) {
	var lines []CoachLine
	for _, n := range []string{"A", "B", "C", "D", "E"} {
		lines = append(lines, line(n, "Stem", 2, 8))
	}
	w, _ := coachWorld(t, lines...)
	got := headFields(w.run(t, pressDiscord, "report"))[coachCopy.FieldName]
	if n := strings.Count(got, "\n") + 1; n != 3 {
		t.Fatalf("%d lines: %q", n, got)
	}
}

func TestReport_PendingLineReadsTheSiteLine(t *testing.T) {
	w, _ := coachWorld(t, CoachLine{Character: "Zulmara", URL: coachURL, Pending: true})
	got := headFields(w.run(t, pressDiscord, "report"))[coachCopy.FieldName]
	if got != "[Your coach's review is on the site.]("+coachURL+")" {
		t.Fatalf("field = %q", got)
	}
}

func TestReport_NoCoachDepsNoLine(t *testing.T) {
	w, _ := coachWorld(t, line("Zulmara", "Stem", 2, 8))
	with := w.run(t, pressDiscord, "report")
	w.svc.d.Coach = nil
	without := w.run(t, pressDiscord, "report")
	if _, ok := headFields(without)[coachCopy.FieldName]; ok {
		t.Fatal("a coach field without Deps.Coach")
	}
	if len(headFields(with)) != len(headFields(without))+1 {
		t.Fatalf("only the coach field may differ: %v vs %v", headFields(with), headFields(without))
	}
}

func TestReport_NoLinesOrAFailedReadLeaveTheAnswerAlone(t *testing.T) {
	w, fc := coachWorld(t)
	if _, ok := headFields(w.run(t, pressDiscord, "report"))[coachCopy.FieldName]; ok {
		t.Fatal("a coach field with no lines")
	}
	fc.err = errors.New("pack store down")
	rep := w.run(t, pressDiscord, "report")
	if _, ok := headFields(rep)[coachCopy.FieldName]; ok || len(rep.Message.Embeds) == 0 {
		t.Fatal("a failed coach read changed the answer")
	}
}

func TestReport_ShareInChannelRendersWithoutTheLine(t *testing.T) {
	w, _ := coachWorld(t, line("Zulmara", "Stem", 2, 8))
	w.og.report = onyxiaReport(true, true)
	answer := w.run(t, pressDiscord, "report").Message
	if _, ok := headFields(Reply{Message: answer})[coachCopy.FieldName]; !ok {
		t.Fatal("the ephemeral answer has no coach field")
	}
	in := interactionOf(typeComponent, pressDiscord, map[string]any{"custom_id": answer.Components[0][0].CustomID})
	in.Token = "tok-coach"
	rep, err := w.svc.share(context.Background(), in)
	if err != nil || rep.After == nil {
		t.Fatalf("press = %+v %v", rep, err)
	}
	rep.After(context.Background())
	for _, c := range w.rest.Calls() {
		if c.Method != "FollowUp" {
			continue
		}
		for _, e := range c.Message.Embeds {
			for _, f := range e.Fields {
				if f.Name == coachCopy.FieldName {
					t.Fatal("the shared post carries the coach field")
				}
			}
		}
	}
}

func TestReport_EveryLimitHolds(t *testing.T) {
	stem := strings.Repeat("s", 60)
	w, _ := coachWorld(t, line("Abcdefghijkl", stem, 2, 8), line("Mnopqrstuvwx", stem, 3, 8), line("Yzabcdefghij", stem, 8, 8))
	rep := w.run(t, pressDiscord, "report")
	assertMessageLimits(t, rep.Message)
	if v := headFields(rep)[coachCopy.FieldName]; len(v) > fieldValueMax || strings.Count(v, "\n") != 2 {
		t.Fatalf("field = %d bytes: %q", len(v), v)
	}
	raw, err := rep.Message.MarshalJSON()
	if err != nil || !strings.Contains(string(raw), `"allowed_mentions":{"parse":[]}`) {
		t.Fatalf("mentions: %s %v", raw, err)
	}
}

func TestCoachBudgetIsClippedByTheContext(t *testing.T) {
	if got := coachBudget(context.Background()); got != coachLinesBudget {
		t.Fatalf("budget = %v", got)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	if got := coachBudget(ctx); got <= 0 || got > 300*time.Millisecond {
		t.Fatalf("clipped budget = %v", got)
	}
}

// ---- The A11 capture payloads ----

const (
	dumpEnv  = "COACH_DUMP_PAYLOADS"
	dumpPath = "../../../web/tests/captures/log-coach/payloads.json"
	// dumpCoachURL is the Coach tab of the example raid, as the real adapter links it.
	dumpCoachURL = "https://foreversixty.gg/guild/us/normal/olympus-xxvii/raids/7#coach"
)

type dumpedFrame struct {
	ID      string  `json:"id"`
	Row     string  `json:"row"`
	Title   string  `json:"title"`
	Surface string  `json:"surface"`
	Clock   string  `json:"clock"`
	Message Message `json:"message"`
}

var dumpClock = time.Date(2026, 12, 11, 4, 0, 0, 0, time.UTC)

// moltenCoreNight is the §5.3 example night: eight bosses, so "5 of 8 bosses" is real.
// moltenCoreKills is a night that stopped after eight of Molten Core's ten bosses, in the order
// a raid meets them: no Majordomo Executus, so no Ragnaros.
var moltenCoreKills = []string{"Lucifron", "Magmadar", "Gehennas", "Garr", "Baron Geddon", "Shazzrah",
	"Sulfuron Harbinger", "Golemagg the Incinerator"}

func moltenCoreNight(officer, public bool) *guilds.ReportAnswer {
	ans := onyxiaReport(officer, public)
	ans.Raid.Title = "Molten Core"
	order := make([]guilds.KillOrderView, 0, len(moltenCoreKills))
	for i, name := range moltenCoreKills {
		ms := int64(150_000 + i*20_000)
		order = append(order, guilds.KillOrderView{Name: name, Killed: true, KillMS: &ms})
	}
	ans.Report.KillOrder = order
	ans.Raid.Night.Kills, ans.Raid.Night.Wipes = len(moltenCoreKills), 0
	return ans
}

// zulmaraStem is a Marksmanship hunter's Fix first (the account and guild mocks' Zulmara): a
// core cast of the hunter-marksmanship rule file.
const zulmaraStem = "Multi-Shot casts lost"

// TestDumpCoachPayloads writes the messages the A11 captures render, produced by the bot's own
// handlers and the notify embed: COACH_DUMP_PAYLOADS=1 go test -run DumpCoachPayloads.
func TestDumpCoachPayloads(t *testing.T) {
	if os.Getenv(dumpEnv) == "" {
		t.Skip(dumpEnv + " is not set")
	}
	w, fc := coachWorld(t)
	w.svc.now = func() time.Time { return dumpClock }
	w.og.report = moltenCoreNight(false, false)
	frames := []dumpedFrame{
		w.dumpReport(t, fc, "report-coach-one", "one character", line("Zulmara", zulmaraStem, 5, 8)),
		w.dumpReport(t, fc, "report-coach-two", "two characters, every boss",
			line("Zulmara", zulmaraStem, 5, 8), line("Duskstrider", "Aimed Shot casts lost", 8, 8)),
		w.dumpReport(t, fc, "report-coach-pending", "pending", CoachLine{Character: "Zulmara", Pending: true}),
		w.dumpShared(t, fc, "report-shared", "after Share in channel", line("Zulmara", zulmaraStem, 5, 8)),
		dumpRaidNight(true),
		dumpRaidNight(false),
	}
	raw, err := json.MarshalIndent(frames, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(dumpPath, append(raw, '\n'), 0o600); err != nil {
		t.Fatal(err)
	}
}

func (w *officerWorld) dumpReport(t *testing.T, fc *fakeCoach, id, title string, lines ...CoachLine) dumpedFrame {
	t.Helper()
	for i := range lines {
		lines[i].URL = dumpCoachURL
	}
	fc.lines = lines
	rep := w.run(t, pressDiscord, "report")
	return dumpedFrame{ID: id, Row: "A11", Title: "/report, Premium: " + title, Surface: "ephemeral", Clock: dumpClock.Format(time.RFC3339), Message: rep.Message}
}

// dumpShared is the channel post Share in channel sends: the same report, re-rendered without
// the coach field.
func (w *officerWorld) dumpShared(t *testing.T, fc *fakeCoach, id, title string, lines ...CoachLine) dumpedFrame {
	t.Helper()
	for i := range lines {
		lines[i].URL = dumpCoachURL
	}
	fc.lines = lines
	prev := w.og.report
	w.og.report = moltenCoreNight(false, true) // only a public report can be shared
	defer func() { w.og.report = prev }()
	answer := w.run(t, pressDiscord, "report").Message
	in := interactionOf(typeComponent, pressDiscord, map[string]any{"custom_id": answer.Components[0][0].CustomID})
	in.Token = "tok-dump"
	rep, err := w.svc.share(context.Background(), in)
	if err != nil || rep.After == nil {
		t.Fatalf("share = %+v %v", rep, err)
	}
	rep.After(context.Background())
	calls := w.rest.Calls()
	for i := len(calls) - 1; i >= 0; i-- {
		if calls[i].Method == "FollowUp" {
			return dumpedFrame{ID: id, Row: "A11", Title: "/report, " + title, Surface: "channel",
				Clock: dumpClock.Format(time.RFC3339), Message: calls[i].Message}
		}
	}
	t.Fatal("share sent no follow-up")
	return dumpedFrame{}
}

func dumpRaidNight(coachReady bool) dumpedFrame {
	post := notify.ReportReady{
		RaidPosted: notify.RaidPosted{
			RaidID: 7, GuildName: "Olympus XXVII", Title: "Molten Core",
			StartsAt: time.Date(2026, 12, 11, 2, 0, 0, 0, time.UTC), Timezone: "America/Chicago",
			RaidURL: "https://foreversixty.gg/guild/us/normal/olympus-xxvii/raids/7",
		},
		Came: 22, Signed: 23, LootItems: 6, CoachReady: coachReady,
	}
	for i, boss := range moltenCoreKills {
		post.Kills = append(post.Kills, notify.ReportKill{Boss: boss, Duration: time.Duration(150+20*i) * time.Second})
	}
	id, title := "report-ready-coach", "report_ready with the coach sentence"
	if !coachReady {
		id, title = "report-ready-plain", "report_ready without the coach sentence"
	}
	return dumpedFrame{ID: id, Row: "A11", Title: title, Surface: "channel",
		Clock: dumpClock.Format(time.RFC3339), Message: Message{Embeds: []notify.Embed{notify.ReportReadyEmbed(post, dumpClock)}}}
}
