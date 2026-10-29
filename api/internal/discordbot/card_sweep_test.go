package discordbot

import (
	"context"
	"fmt"
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
		w.attendanceModal(t, "Vexlash").Message,
	}
	// Each world truncates the database, so every world's frames are drawn before the next.
	frames = append(frames, signupFrames(t)...)
	frames = append(frames, specFrames(t)...)
	frames = append(frames, newOfficerWorld(t).run(t, pressDiscord, "char").Message)
	return append(frames, lookupFrames(t)...)
}

// signupFrames are the bare /signup card and the card a raid pick redraws.
func signupFrames(t *testing.T) []Message {
	t.Helper()
	w := newOfficerWorld(t)
	w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", true)
	bare := w.run(t, pressDiscord, "signup").Message
	later := w.addRaid(t, w.gid, w.raidStart.Add(24*time.Hour))
	raidPick := w.signupCardPress(t, fmt.Sprintf("gsgn:1:raid:%d:%s", w.raidID, keyZulmara), fmt.Sprint(later))
	return []Message{bare, raidPick.Message}
}

// specFrames are the spec card an In asks, the outcome its pick answers and the Change spec card.
func specFrames(t *testing.T) []Message {
	t.Helper()
	w := newPressWorld(t)
	paladin(t, w)
	ask := w.press(t, "in").Message
	done := w.mustSpecPress(t, specPickID(w, "in", keyZulmara), "paladin-holy").Message
	change := w.mustSpecPress(t, fmt.Sprintf("gspc:1:chg:%d:%s", w.raidID, keyZulmara)).Message
	return []Message{ask, done, change}
}

// lookupFrames are the /char card a typed name redraws and a /bis range frame.
func lookupFrames(t *testing.T) []Message {
	t.Helper()
	lookup := newLookupWorld(t)
	linkTwoChars(lookup)
	charName := lookup.modal(t, pressDiscord, charNameID, "Zulmara").Message
	bisRange := newBisWorld(t).pressBis(t, pressDiscord, "gbis:1:r:warrior:fury:p1:", "50").Message
	return []Message{charName, bisRange}
}

// cardModals are the modals each card's ask part opens, keyed by prefix: derived by pressing
// ask on every routed card, so a new modal is checked without being listed here. ask reads
// nothing, so an empty Service answers it.
func cardModals(t *testing.T) map[string]*Modal {
	t.Helper()
	s := &Service{}
	routes := s.componentRoutes()
	out := map[string]*Modal{}
	for _, prefix := range cardPrefixes(s) {
		in := interactionOf(typeComponent, pressDiscord, map[string]any{"custom_id": cardCustomID(prefix, askPart)})
		rep, err := routes[prefix](context.Background(), in)
		if err != nil {
			t.Fatalf("%s ask: %v", prefix, err)
		}
		if rep.Modal != nil {
			out[prefix] = rep.Modal
		}
	}
	return out
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

	routed := map[string]bool{}
	for _, prefix := range prefixes {
		routed[prefix] = true
	}
	for prefix := range cardFailureRoutes {
		if !routed[prefix] {
			t.Errorf("%s is failure-routed but has no component route", prefix)
		}
	}
	for prefix := range updates {
		if !routed[prefix] && !legacyComponentPrefixes[prefix] {
			t.Errorf("%s is an update route but has no component route", prefix)
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
	modals := cardModals(t)
	// The two cards that ask for a typed name today; a card losing its modal fails here.
	for _, prefix := range []string{prefixChar, prefixAttendance} {
		if modals[prefix] == nil {
			t.Errorf("%s ask opened no modal", prefix)
		}
	}
	routes := (&Service{}).componentRoutes()
	for prefix, modal := range modals {
		if !cardFailureRoutes[prefix] || modal.CustomID != cardCustomID(prefix, "name") ||
			customPrefix(modal.CustomID) != prefix || routes[prefix] == nil ||
			utf8.RuneCountInString(modal.CustomID) > maxCustomIDLen || utf8.RuneCountInString(modal.Title) > 45 ||
			utf8.RuneCountInString(modal.InputLabel) > 45 || utf8.RuneCountInString(modal.Placeholder) > 100 {
			t.Errorf("%s modal = %+v", prefix, modal)
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
