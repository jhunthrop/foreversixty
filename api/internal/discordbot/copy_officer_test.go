package discordbot

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

// The fakeGuilds stubs the install tests' port needs for the officer reads.
func (f *fakeGuilds) ReportFor(context.Context, int64, int64, *int64) (guilds.ReportAnswer, bool, error) {
	return guilds.ReportAnswer{}, false, nil
}
func (f *fakeGuilds) DoneRaids(context.Context, int64, bool, int) ([]guilds.Raid, error) {
	return nil, nil
}
func (f *fakeGuilds) AttendanceFor(context.Context, int64, int64, string) ([]guilds.AttendanceRecord, error) {
	return nil, nil
}
func (f *fakeGuilds) LowestAttendance(context.Context, int64, int64, int) ([]guilds.AttendanceRecord, error) {
	return nil, nil
}
func (f *fakeGuilds) AttendanceNames(context.Context, int64, int64, string, int) ([]string, error) {
	return nil, nil
}
func (f *fakeGuilds) NextRaidReadiness(context.Context, int64, int64, time.Time) (guilds.ReadinessAnswer, error) {
	return guilds.ReadinessAnswer{}, nil
}
func (f *fakeGuilds) GuildPagePathOf(context.Context, int64) (string, error) { return "", nil }

func TestOfficerCopyMatchesTheSpec(t *testing.T) {
	cell := func(label string) []string { return specCell(t, pressHeading, label) }
	c := officerCopy
	for _, row := range []struct {
		name      string
		got, want []string
	}{
		{"report", cell("`/report`"), []string{c.ReportPrefix + "{Title} · {Weekday} {d} {Month}", c.KillOrder, c.MostDeaths,
			c.Consumables, c.Loot, c.LootAll, c.LootCount, c.LootFirst, c.LootTo, c.OpenReport, c.ReportNone, c.ReportUnlinked}},
		{"attendance", cell("`/attendance`"), []string{c.AttendanceTitle, c.LastRaids, c.Roster, c.OfficersOnly, c.NoMember,
			c.NoRaids, c.AttendanceUnlinked, c.AttendanceMembers}},
		{"readiness", cell("`/readiness`"), []string{c.ReadinessTitle + "{Title} · {Weekday} {d} {Month}", c.ReadinessLead,
			"1 has", c.ReadinessAllReady, c.ToFix, c.ReadinessMore, c.Board, c.NotOfficer, c.ReadinessUnlinked, c.PlanLapsed,
			c.Subscribe, c.Frozen, c.RaidGuild, c.MoveToRaidGuild, c.ReadinessNobodySigned}},
	} {
		if strings.Join(row.got, "|") != strings.Join(row.want, "|") {
			t.Errorf("%s: spec %q\ncopy %q", row.name, row.got, row.want)
		}
	}
	if want := strings.Replace(c.ReadinessLead, "{k} have", "1 has", 1); c.ReadinessLeadOne != want {
		t.Errorf("one = %q, want %q", c.ReadinessLeadOne, want)
	}
	if outcomeWords["present"] != "Present" || outcomeWords["no_show"] != "Not in the log" || excusedSuffix != ", excused" {
		t.Errorf("outcome words %v", outcomeWords)
	}
}
