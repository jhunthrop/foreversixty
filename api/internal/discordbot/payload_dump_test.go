//go:build payloaddump

package discordbot

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
	"github.com/jhunthrop/foreversixty/api/internal/rankings"
)

// payloadDirEnv names the folder TestDumpGuidedPayloads writes into. Unset, the test skips, so
// a normal run never writes files. The committed copies live in
// design/mocks/2026-10-28-discord-guided-cards/payloads (see README there).
const payloadDirEnv = "GUIDED_PAYLOADS_DIR"

// gearPayloadDirEnv names the folder TestDumpGearPanelPayloads writes the gear panel's real
// frames into: design/mocks/2026-10-29-discord-gear-panel/payloads-real, drawn by
// gen_gear_panel.mjs beside the mock's own.
const gearPayloadDirEnv = "GEAR_PAYLOADS_DIR"

// signedUpOfForty is the roster size the reminder boards show: 33 of 40.
const signedUpOfForty = 33

// dumpFrame is one frame file: the board id it answers to, the message the handler produced
// (as Discord is sent it), a message above it when the frame is a reply, and a modal when the
// answer opens one.
type dumpFrame struct {
	ID      string          `json:"id"`
	Title   string          `json:"title"`
	Clock   time.Time       `json:"clock"`
	Message json.RawMessage `json:"message,omitempty"`
	Above   json.RawMessage `json:"above,omitempty"`
	Modal   json.RawMessage `json:"modal,omitempty"`
}

type frameDumper struct {
	t   *testing.T
	dir string
}

func (d frameDumper) write(f dumpFrame) {
	d.t.Helper()
	f.Clock = boardClock
	raw, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		d.t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(d.dir, f.ID+".json"), append(raw, '\n'), 0o600); err != nil {
		d.t.Fatal(err)
	}
}

// stamped pins every embed timestamp of a drained post to the board clock, which the drainer
// takes from the machine's clock.
var stampPattern = regexp.MustCompile(`"timestamp":"[^"]+"`)

func stamped(raw json.RawMessage) json.RawMessage {
	return stampPattern.ReplaceAll(raw, []byte(`"timestamp":"`+boardClock.Format(time.RFC3339)+`"`))
}

func (d frameDumper) message(id, title string, m Message) {
	d.t.Helper()
	d.write(dumpFrame{ID: id, Title: title, Message: marshalled(d.t, m)})
}

func (d frameDumper) reply(id, title string, rep Reply) {
	d.t.Helper()
	if rep.Modal != nil {
		d.write(dumpFrame{ID: id, Title: title, Modal: marshalled(d.t, rep.Modal)})
		return
	}
	d.message(id, title, rep.Message)
}

// TestDumpGuidedPayloads writes the real handler answer of every guided-card and raid-post
// board scenario to $GUIDED_PAYLOADS_DIR, one <frame>.json each, for the mock generator to draw
// beside the designer's boards.
func TestDumpGuidedPayloads(t *testing.T) {
	dir := os.Getenv(payloadDirEnv)
	if dir == "" {
		t.Skip(payloadDirEnv + " is not set")
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	d := frameDumper{t: t, dir: dir}
	dumpBisFrames(d)
	dumpCharFrames(d)
	dumpSignupFrames(d)
	dumpOfficerFrames(d)
	dumpPostFrames(d)
	dumpSpecFrames(d)
}

// hunterBeastMasteryBands is a one-slot band so a Show BiS answer exists for the board's spec.
func hunterBeastMasteryBands() []map[string]any {
	return []map[string]any{{"band": 60, "faction": "alliance", "slots": []map[string]any{
		{"slot": "head", "item_id": 100, "item_name": "Pick head", "alternatives": []map[string]any{
			{"item_id": 200, "item_name": "Worn head", "dps_delta": -9.2}}},
	}}}
}

func dumpBisFrames(d frameDumper) {
	t := d.t
	w := newBisWorld(t)
	w.svc.now = func() time.Time { return boardClock }
	writeBis(t, w.lg.dir, "hunter-beast-mastery", hunterBeastMasteryBands())
	stranger := "d-stranger"
	press := func(user, id string, values ...string) Message { return w.pressBis(t, user, id, values...).Message }

	d.message("bis-card-1-class", "not linked: the bare card", w.bisCardOf(t, stranger, nil))
	d.message("bis-card-2-spec", "after the class pick", press(stranger, "gbis:1:c:::p1:", "hunter"))
	ready := press(stranger, "gbis:1:s:hunter::p1:", "beast-mastery")
	d.message("bis-card-3-range", "after the spec pick", ready)
	d.message("bis-card-4-ready", "every pick made", ready)
	d.message("bis-card-nudge", "Show BiS before a spec", press(stranger, "gbis:1:go:hunter::p1:"))

	linkMain(w)
	d.message("bis-card-main", "linked: picked from the main", w.bisCardOf(t, pressDiscord, nil))
	// The board's changed card is the main's Beast Mastery card with Marksmanship at level 50.
	press(pressDiscord, "gbis:1:s:hunter:beast-mastery:p1:", "marksmanship")
	d.message("bis-card-main-changed", "linked, a pick changed",
		press(pressDiscord, "gbis:1:r:hunter:marksmanship:p1:", "50"))
	d.reply("bis-card-stale", "a press on a card of an older shape", w.card(t, stranger, "gbis:2:c:::p1:", "hunter"))
	d.message("bis-result", "Show BiS: the card becomes the typed answer",
		press(pressDiscord, "gbis:1:go:hunter:beast-mastery:p1:"))
}

func dumpCharFrames(d frameDumper) {
	t := d.t
	w := newLookupWorld(t)
	w.svc.now = func() time.Time { return boardClock }
	d.reply("char-card-unlinked", "/char, not linked: asks", w.command(t, "d-stranger", "char", nil))
	linkTwoChars(w)
	d.reply("char-card-own", "/char, linked: the main and the own-characters select", w.command(t, pressDiscord, "char", nil))
	d.reply("char-card-modal", "Someone else: the name field", w.card(t, pressDiscord, charAskID))
	d.reply("char-card-unknown", "a name that is nobody", w.modal(t, pressDiscord, charNameID, "Nobody"))
}

func dumpSignupFrames(d frameDumper) {
	t := d.t
	w := newPressWorld(t)
	w.svc.now = func() time.Time { return boardClock }
	w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", true)
	w.addChar(t, w.gid, w.uid, "us/normal/duskstrider", "Duskstrider", true)
	second := w.addRaid(t, w.gid, w.raidStart.Add(72*time.Hour))
	w.exec(t, `update guild_raids set title = 'Molten Core', starts_at = $2 where id = $1`, second, boardClock.Add(120*time.Hour))
	w.exec(t, `update guild_raids set starts_at = $2 where id = $1`, w.raidID, boardClock.Add(48*time.Hour))
	card := w.bareSignup(t)
	d.reply("signup-card", "/signup, two raids and two characters", card)
	d.reply("signup-card-in", "a status press: In", w.signupCardPress(t, buttonsOf(card.Message)[0].CustomID))
}

func dumpOfficerFrames(d frameDumper) {
	t := d.t
	w := newOfficerWorld(t)
	w.svc.now = func() time.Time { return boardClock }
	w.addChar(t, w.gid, w.uid, "us/normal/zulmara", "Zulmara", true)
	officerLowest(w)
	raids := plannedRaids()
	w.og.planned = raids
	w.og.readiness = readinessFor(raids[0], 20)
	w.og.raidReady = map[int64]guilds.ReadinessAnswer{4: readinessFailing(raids[1])}
	last := onyxiaReport(true, true)
	older := onyxiaReport(false, false)
	older.Raid.ID, older.Raid.Title = 6, "Molten Core"
	w.og.report = last
	w.og.reports = map[int64]*guilds.ReportAnswer{7: last, 6: older}
	w.og.done = []guilds.Raid{
		{ID: 7, Title: last.Raid.Title, StartsAt: last.Raid.StartsAt},
		{ID: 6, Title: "Molten Core", StartsAt: last.Raid.StartsAt.Add(-7 * 24 * time.Hour)},
	}
	d.reply("attendance-select", "/attendance, an officer: the lowest-attendance select", w.run(t, pressDiscord, "attendance"))
	d.reply("attendance-modal", "Type a name: the name field", w.attendanceCard(t, attAskID))
	d.reply("attendance-pick", "an officer's pick", w.attendanceCard(t, attPickID, "Thalwyn"))
	d.reply("readiness-select", "/readiness: the Another raid select", w.run(t, pressDiscord, "readiness"))
	d.reply("readiness-pick", "a raid picked", w.attendanceCard(t, rdyPickID, "4"))
	d.reply("report-select", "/report: the Another night select", w.run(t, pressDiscord, "report"))
	d.reply("report-pick", "a night picked", w.attendanceCard(t, repPickID, "6"))
}

// boardStart is the boards' raid: Thursday 10 December 2026, 8:00 pm in Chicago.
var boardStart = time.Date(2026, time.December, 11, 2, 0, 0, 0, time.UTC)

// boardRaid pins the raid the post worlds hold to the boards' night.
func boardRaid(w *postWorld, counts notify.RaidCounts) {
	w.facts.edit(w.raidID, func(f *guilds.RaidPostFacts) {
		f.Post.StartsAt, f.Post.Timezone, f.Post.Counts = boardStart, "America/Chicago", counts
		f.EndsAt = boardStart.Add(3 * time.Hour)
	})
}

// TestDumpGearPanelPayloads writes the real handler frames of every gear-panel board to
// $GEAR_PAYLOADS_DIR, one <board>.json each, in the mock's own frame shape (id, title, budget,
// message), for gen_gear_panel.mjs to draw as <board>-real beside the board.
func TestDumpGearPanelPayloads(t *testing.T) {
	dir := os.Getenv(gearPayloadDirEnv)
	if dir == "" {
		t.Skip(gearPayloadDirEnv + " is not set")
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	for id, m := range gearBoardFrames(t) {
		frame := map[string]any{"id": id, "title": "real payload: " + id,
			"budget": map[string]int{"components": m.componentCount(), "text": textLength(m.Layout)}, "message": json.RawMessage(marshalled(t, m))}
		raw, err := json.MarshalIndent(frame, "", "  ")
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(dir, id+".json"), append(raw, '\n'), 0o600); err != nil {
			t.Fatal(err)
		}
	}
}

// gearBoardFrames are the handlers' answers for each of the mock's nine boards.
func gearBoardFrames(t *testing.T) map[string]Message {
	t.Helper()
	w := newGearWorld(t)
	p1 := w.bisPanel(t, gearBisOpts)
	in := interactionOf(typeComponent, pressDiscord, map[string]any{"custom_id": p1.Components[0][1].CustomID})
	in.Source = sourceOf(t, p1)
	out := map[string]Message{
		"bis-panel-p1":   p1,
		"bis-panel-p2":   w.route(t, in).Message,
		"bis-panel-slot": w.bisPanel(t, map[string]string{"spec": "hunter-beast-mastery", "phase": "p1", "slot": "legs"}),
		"char-panel":     w.command(t, pressDiscord, "char", nil).Message,
	}
	share := w.pressShareFrom(t, p1)
	out["bis-panel-p1-sharefail"] = withLiveShare(share.Message, p1.Components[0][0].CustomID[len(shareIDPrefix)+1:])
	share.After(context.Background())
	out["share-bis"] = w.followUps()[0].Message
	w.sheets.err = errors.New("no sheet")
	again := w.pressShareFrom(t, w.bisPanel(t, gearBisOpts))
	again.After(context.Background())
	out["share-bis-nosheet"] = w.followUps()[1].Message
	w.sheets.err = nil
	charShare := w.pressShareFrom(t, out["char-panel"])
	charShare.After(context.Background())
	out["share-char"] = w.followUps()[2].Message
	bare := gearZulmara()
	bare.Gear = []rankings.GearSlot{}
	w.cards.cards["zulmara"] = bare
	w.lg.readiness[keyZulmara] = guilds.ReadinessFacts{Word: "Not synced", OnRoster: true}
	out["char-panel-nogear"] = w.command(t, pressDiscord, "char", nil).Message
	return out
}
