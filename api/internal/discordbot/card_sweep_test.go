package discordbot

import (
	"sort"
	"strings"
	"testing"
	"time"
	"unicode/utf8"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

// preAmendmentCommandsHash is commandsHash of the nine commands before the guided cards
// (commit 61b459472); the amendment must change what Discord is sent.
const preAmendmentCommandsHash = "320cbf017bb1bd518beb613361a51cc979aaa23cdaf71b4757df3407fca10d5e"

// legacyComponentPrefixes are the routes that are not guided cards.
var legacyComponentPrefixes = map[string]bool{raidIDPrefix: true, pickIDPrefix: true, shareIDPrefix: true}

// cardPrefixes is every card prefix in the routing table.
func cardPrefixes(s *Service) []string {
	var out []string
	for prefix := range s.componentRoutes() {
		if !legacyComponentPrefixes[prefix] {
			out = append(out, prefix)
		}
	}
	sort.Strings(out)
	return out
}

// sweepFrames are the cards each command draws and the cards a press redraws.
func sweepFrames(t *testing.T) []Message {
	t.Helper()
	w := newOfficerWorld(t)
	w.addChar(t, w.gid, w.uid, "us/normal/zulmara", "Zulmara", true)
	officerLowest(w)
	raids := plannedRaids()
	w.og.planned = raids
	w.og.readiness = readinessFor(raids[0], 20)
	w.og.raidReady = map[int64]guilds.ReadinessAnswer{4: readinessFailing(raids[1])}
	last := onyxiaReport(true, true)
	w.og.report = last
	w.og.reports = map[int64]*guilds.ReportAnswer{7: last}
	w.og.done = []guilds.Raid{
		{ID: 7, Title: last.Raid.Title, StartsAt: last.Raid.StartsAt},
		{ID: 6, Title: "Molten Core", StartsAt: last.Raid.StartsAt.Add(-7 * 24 * time.Hour)},
	}
	older := onyxiaReport(false, false)
	older.Raid.ID = 6
	w.og.reports[6] = older

	frames := []Message{
		w.run(t, pressDiscord, "bis").Message,
		w.run(t, pressDiscord, "bis", opt("class", "warrior")).Message,
		w.run(t, pressDiscord, "attendance").Message,
		w.run(t, pressDiscord, "readiness").Message,
		w.run(t, pressDiscord, "report").Message,
		w.attendanceCard(t, attPickID, "Thalwyn").Message,
		w.attendanceCard(t, rdyPickID, "4").Message,
		w.attendanceCard(t, repPickID, "6").Message,
	}
	signup := newOfficerWorld(t)
	signup.addChar(t, signup.gid, signup.uid, "us/normal/zulmara", "Zulmara", true)
	return append(frames, signup.run(t, pressDiscord, "signup").Message, newOfficerWorld(t).run(t, pressDiscord, "char").Message)
}

func TestEveryCardHoldsDiscordsLimits(t *testing.T) {
	s := &Service{}
	prefixes := cardPrefixes(s)
	if len(prefixes) != len(cardFailureRoutes) {
		t.Fatalf("routed card prefixes %v, failure-routed %d", prefixes, len(cardFailureRoutes))
	}
	updates := updateComponents()
	for _, prefix := range prefixes {
		if !cardFailureRoutes[prefix] || !updates[prefix] {
			t.Errorf("%s: failure route %v, update route %v", prefix, cardFailureRoutes[prefix], updates[prefix])
		}
	}

	seen := map[string]bool{}
	for i, m := range sweepFrames(t) {
		assertCardLimits(t, m)
		if len(m.Components) == 0 && len(m.Embeds) == 0 {
			t.Errorf("frame %d is empty", i)
		}
		for _, row := range m.Components {
			for _, c := range row {
				if c.CustomID != "" && !strings.HasPrefix(c.CustomID, shareIDPrefix+":") {
					seen[customPrefix(c.CustomID)] = true
				}
			}
		}
	}
	for _, prefix := range prefixes {
		if !seen[prefix] {
			t.Errorf("no swept frame carries a %s component", prefix)
		}
	}
}

func TestCardModalsHoldDiscordsLimits(t *testing.T) {
	for prefix, modal := range map[string]*Modal{prefixChar: charModal(), prefixAttendance: attendanceModal()} {
		if !cardFailureRoutes[prefix] || modal.CustomID != cardCustomID(prefix, "name") ||
			utf8.RuneCountInString(modal.CustomID) > maxCustomIDLen || utf8.RuneCountInString(modal.Title) > 45 ||
			utf8.RuneCountInString(modal.InputLabel) > 45 || utf8.RuneCountInString(modal.Placeholder) > 100 {
			t.Errorf("%s modal = %+v", prefix, modal)
		}
	}
	routes := (&Service{}).componentRoutes()
	for _, prefix := range []string{prefixChar, prefixAttendance} {
		if routes[prefix] == nil {
			t.Errorf("%s does not route a modal submit", prefix)
		}
	}
}

func TestGuidedCommandSetChangedAndKeptItsSize(t *testing.T) {
	cmds := Commands()
	if len(cmds) != 9 {
		t.Fatalf("%d commands", len(cmds))
	}
	hash, err := commandsHash(cmds)
	if err != nil {
		t.Fatal(err)
	}
	if hash == preAmendmentCommandsHash {
		t.Fatal("the command set Discord holds is unchanged by the amendment")
	}
}
