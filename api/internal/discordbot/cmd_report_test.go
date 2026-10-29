package discordbot

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// officerGuilds is the real guilds service with the officer reads scripted: what each read
// answers is the guilds package's to test (discord_reads_test.go); here the bot's rendering is.
type officerGuilds struct {
	*pressGuilds
	report     *guilds.ReportAnswer
	reportSeen []int64 // the viewer ids ReportFor was asked for
	records    []guilds.AttendanceRecord
	recordsErr error
	byName     map[string]guilds.AttendanceRecord // AttendanceFor answers these for a typed name
	lowest     []guilds.AttendanceRecord
	lowestErr  error
	readiness  guilds.ReadinessAnswer
	readyErr   error
	// planned and done script the raid lists; raidReady and reports answer by raid id.
	planned   []guilds.Raid
	done      []guilds.Raid
	raidReady map[int64]guilds.ReadinessAnswer
	reports   map[int64]*guilds.ReportAnswer
}

func (o *officerGuilds) PlannedRaids(ctx context.Context, guildID int64, now time.Time) ([]guilds.Raid, error) {
	if o.planned != nil {
		return o.planned, nil
	}
	return o.pressGuilds.PlannedRaids(ctx, guildID, now)
}

func (o *officerGuilds) DoneRaids(ctx context.Context, guildID int64, publicOnly bool, limit int) ([]guilds.Raid, error) {
	if o.done != nil {
		return o.done, nil
	}
	return o.pressGuilds.DoneRaids(ctx, guildID, publicOnly, limit)
}

func (o *officerGuilds) IsOfficer(context.Context, int64, int64) (bool, error) {
	return o.lowestErr == nil, nil
}

func (o *officerGuilds) RaidReadiness(_ context.Context, _, _, raidID int64) (guilds.ReadinessAnswer, error) {
	if o.readyErr != nil {
		return guilds.ReadinessAnswer{}, o.readyErr
	}
	return o.raidReady[raidID], nil
}

// ReportFor scripts the report but takes membership from the real guild (GuildRank, as
// guilds.ReportFor does): anyone who is not a verified member sees a public log only, as no member.
func (o *officerGuilds) ReportFor(ctx context.Context, guildID, viewerID int64, raidID *int64) (guilds.ReportAnswer, bool, error) {
	o.reportSeen = append(o.reportSeen, viewerID)
	_, member, err := o.Accounts.GuildRank(ctx, guildID, viewerID)
	if err != nil {
		return guilds.ReportAnswer{}, false, err
	}
	script := o.report
	if raidID != nil && o.reports[*raidID] != nil {
		script = o.reports[*raidID]
	}
	if script == nil || (!member && !script.Public) {
		return guilds.ReportAnswer{}, false, nil
	}
	ans := *script
	ans.Member = member
	ans.Officer = ans.Officer && member
	return ans, true, nil
}

func (o *officerGuilds) AttendanceFor(_ context.Context, _, _ int64, name string) ([]guilds.AttendanceRecord, error) {
	if rec, ok := o.byName[name]; ok && o.recordsErr == nil {
		return []guilds.AttendanceRecord{rec}, nil
	}
	if name != "" && o.byName != nil && o.recordsErr == nil {
		return nil, guilds.ErrNoSuchMember
	}
	return o.records, o.recordsErr
}

func (o *officerGuilds) LowestAttendance(context.Context, int64, int64, int) ([]guilds.AttendanceRecord, error) {
	return o.lowest, o.lowestErr
}

func (o *officerGuilds) NextRaidReadiness(context.Context, int64, int64, time.Time) (guilds.ReadinessAnswer, error) {
	return o.readiness, o.readyErr
}

func (o *officerGuilds) ActiveBuild() (string, bool) { return lookupBuild, true }
func (o *officerGuilds) BuildItem(string, string, int) (trees.Item, bool) {
	return trees.Item{ItemLevel: 74}, true
}

type officerWorld struct {
	*pressWorld
	og *officerGuilds
}

func newOfficerWorld(t *testing.T) *officerWorld {
	t.Helper()
	pw := newPressWorld(t)
	og := &officerGuilds{pressGuilds: pw.port, lowestErr: guilds.ErrOfficerOnly}
	pw.svc = New(Deps{Store: pw.st, REST: pw.rest, Guilds: og, PublicBaseURL: pressPublic, APIBaseURL: pressAPI, Log: slog_(pw.log)})
	return &officerWorld{pressWorld: pw, og: og}
}

func (w *officerWorld) run(t *testing.T, discordUser, name string, options ...map[string]any) Reply {
	t.Helper()
	opts := append([]map[string]any{}, options...)
	rep, err := w.svc.commands[name](context.Background(), interactionOf(typeCommand, discordUser, map[string]any{"name": name, "options": opts}))
	if err != nil {
		t.Fatal(err)
	}
	return rep
}

// strangerDiscord is a Discord user with no link; outsiderDiscord one linked to an account with
// no character in the guild.
const (
	strangerDiscord = "999000999"
	outsiderDiscord = "888000888"
)

func onyxiaReport(officer, public bool) *guilds.ReportAnswer {
	tz := "America/Chicago"
	kill := int64(102_000)
	var loot []guilds.ReportLootAward
	for i := 0; i < 10; i++ {
		loot = append(loot, guilds.ReportLootAward{ItemID: 1000 + i, ItemName: "Item", Quality: 2 + i%3, Icon: "inv_x", Name: "Pyrewisp"})
	}
	report := &guilds.ReportView{
		KillOrder: []guilds.KillOrderView{{Name: "Onyxia", Killed: true, KillMS: &kill, WipesBefore: 3}},
		Deaths: []guilds.ReportDeathView{{Name: "Windtalon", OnWipes: 3}, {Name: "Pyrewisp", OnWipes: 3, OnKills: 1},
			{Name: "Vexlash", OnKills: 2}, {Name: "Quiet", OnKills: 1}},
		Consumables: guilds.ReportConsumablesView{Rows: []guilds.ConsumableRow{{Before: "short"}, {Before: "short"}, {Before: "short"}, {Before: "stocked"}}},
		Loot:        []guilds.ReportLootBoss{{Boss: "Onyxia", Awards: loot}},
	}
	return &guilds.ReportAnswer{
		Raid: guilds.RaidView{ID: 7, Title: "Onyxia's Lair", Status: guilds.RaidDone, GuildTimezone: &tz,
			StartsAt: time.Date(2026, 12, 11, 2, 0, 0, 0, time.UTC), Match: &guilds.MatchView{ReportID: "abc"},
			Night: &guilds.NightView{Signed: 23, Came: 22, Kills: 1, Wipes: 3, Deaths: 14, DurationMS: (2*60 + 41) * 60_000}},
		Report: report, Officer: officer, Public: public,
	}
}

func headFields(rep Reply) map[string]string {
	out := map[string]string{}
	for _, f := range rep.Message.Embeds[0].Fields {
		out[f.Name] = f.Value
	}
	return out
}

func hasShare(rep Reply) bool {
	for _, row := range rep.Message.Components {
		for _, c := range row {
			if strings.HasPrefix(c.CustomID, shareIDPrefix+":") {
				return true
			}
		}
	}
	return false
}

func TestReportMemberAndOfficer(t *testing.T) {
	w := newOfficerWorld(t)
	segment := func(i int) string { return specString(t, pressHeading, "`/report`", i, nil) }
	w.addChar(t, w.gid, w.uid, "us/normal/zulmara", "Zulmara", true)

	w.og.report = onyxiaReport(false, false)
	member := w.run(t, pressDiscord, "report")
	if b := member.Message.Components[0][0]; b.URL != pressPublic+"/guild/us/normal/olympus-xxvii/raids/7#report" {
		t.Fatalf("member link = %q", b.URL)
	}
	assertMessageLimits(t, member.Message)
	head := member.Message.Embeds[0]
	wantTitle := specString(t, pressHeading, "`/report`", 0, map[string]string{"Title": "Onyxia's Lair", "Weekday": "Thursday", "d": "10", "Month": "December"})
	if head.Title != wantTitle || head.Color != doneGreen || head.Description != "22 of 23 came · 1 kill · 3 wipes · 14 deaths · 2 h 41 min" {
		t.Fatalf("head = %q %d %q", head.Title, head.Color, head.Description)
	}
	fields := headFields(member)
	if fields[segment(1)] != "1. Onyxia 1:42 after 3 wipes" {
		t.Fatalf("kill order = %q", fields[segment(1)])
	}
	if _, ok := fields[segment(2)]; ok {
		t.Fatal("a member sees Most deaths")
	}
	if _, ok := fields[segment(3)]; ok {
		t.Fatal("a member sees Consumables")
	}
	if fields[segment(4)] != "10 items. The first 8:" {
		t.Fatalf("loot = %q", fields[segment(4)])
	}
	if len(member.Message.Embeds) != 9 || member.Message.Embeds[1].Color != qualityColors[4] ||
		member.Message.Embeds[8].Color == qualityColors[4] || !strings.Contains(member.Message.Embeds[1].Description, "to Pyrewisp · Onyxia") {
		t.Fatalf("loot embeds = %+v", member.Message.Embeds[1:])
	}
	if hasShare(member) || !member.Message.Ephemeral {
		t.Fatal("a private log offers Share, or the answer is public")
	}

	w.og.report = onyxiaReport(true, true)
	officer := w.run(t, pressDiscord, "report")
	assertMessageLimits(t, officer.Message)
	fields = headFields(officer)
	if fields[segment(2)] != "Pyrewisp 4 · Windtalon 3 · Vexlash 2" || fields[segment(3)] != "3 short before the raid" {
		t.Fatalf("officer fields = %v", fields)
	}
	assertDistinctEmbedURLs(t, officer.Message.Embeds)
	if !hasShare(officer) || len(officer.Message.Embeds) > 9 {
		t.Fatal("a public log offers no Share")
	}
	visitor, found, err := w.svc.answerReport(context.Background(), w.gid, nil, map[string]string{optReportRaid: "7"})
	if err != nil || !found || len(visitor.Embeds) != 1 || len(visitor.Embeds[0].Fields) != 1 ||
		visitor.Embeds[0].URL != pressPublic+"/reports/abc" || visitor.Components[0][0].URL != pressPublic+"/reports/abc" {
		t.Fatalf("shared answer = %+v %v %v", visitor, found, err)
	}

	// A linked account with no character in the guild is no member: a private log is not
	// there for it, and a public one is the visitor's answer (no loot, no raid page).
	outsider := seedUser(t, w.st.Pool)
	if err := w.st.SaveLink(context.Background(), outsider, outsiderDiscord, "outsider"); err != nil {
		t.Fatal(err)
	}
	w.og.report = onyxiaReport(false, false)
	if rep := w.run(t, outsiderDiscord, "report"); rep.Message.Content != segment(10) {
		t.Fatalf("non-member on a private log = %+v", rep.Message)
	}
	w.og.report = onyxiaReport(false, true)
	out := w.run(t, outsiderDiscord, "report")
	assertMessageLimits(t, out.Message)
	if len(out.Message.Embeds) != 1 || len(out.Message.Embeds[0].Fields) != 1 ||
		out.Message.Embeds[0].URL != pressPublic+"/reports/abc" || !hasShare(out) {
		t.Fatalf("non-member on a public log = %+v", out.Message)
	}
	if _, ok := headFields(out)[segment(4)]; ok {
		t.Fatal("a non-member sees the Loot field")
	}
	for _, row := range out.Message.Components {
		for _, c := range row {
			if strings.Contains(c.URL, "/raids/") {
				t.Fatalf("a non-member gets the raid page link %q", c.URL)
			}
		}
	}

	w.og.report = nil
	if rep := w.run(t, pressDiscord, "report"); rep.Message.Content != segment(10) {
		t.Fatalf("none = %q", rep.Message.Content)
	}
	rep := w.run(t, strangerDiscord, "report")
	if rep.Message.Content != segment(11) || rep.Message.Components[0][0].Label != linkCopy.LinkButton {
		t.Fatalf("unlinked = %+v", rep.Message)
	}
}

// TestReportShareIsTheNightAndKillOrder: a public report's Share posts the visitor's answer,
// the night line and kill order linking the public report, with no loot, deaths or consumables.
func TestReportShareIsTheNightAndKillOrder(t *testing.T) {
	w := newOfficerWorld(t)
	w.og.report = onyxiaReport(true, true)
	answer := w.run(t, pressDiscord, "report").Message
	in := interactionOf(typeComponent, pressDiscord, map[string]any{"custom_id": answer.Components[0][0].CustomID})
	in.Token = "tok-report"
	w.og.reportSeen = nil
	rep, err := w.svc.share(context.Background(), in)
	if err != nil || rep.After == nil {
		t.Fatalf("press = %+v %v", rep, err)
	}
	rep.After(context.Background())
	var public *Message
	for _, c := range w.rest.Calls() {
		if c.Method == "FollowUp" {
			m := c.Message
			public = &m
		}
	}
	if public == nil || public.Ephemeral || len(public.Embeds) != 1 {
		t.Fatalf("public = %+v", public)
	}
	e := public.Embeds[0]
	if len(e.Fields) != 1 || e.Fields[0].Name != officerCopy.KillOrder || e.URL != pressPublic+"/reports/abc" ||
		!strings.HasPrefix(e.Description, "22 of 23 came") {
		t.Fatalf("shared embed = %+v", e)
	}
	if len(w.og.reportSeen) == 0 || w.og.reportSeen[0] != 0 {
		t.Fatalf("the share was not recomputed as a visitor first: %v", w.og.reportSeen)
	}
}
