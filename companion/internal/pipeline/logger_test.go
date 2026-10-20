package pipeline_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/companion/internal/character"
	"github.com/jhunthrop/foreversixty/companion/internal/client"
	"github.com/jhunthrop/foreversixty/companion/internal/fixture"
	"github.com/jhunthrop/foreversixty/companion/internal/logger"
	"github.com/jhunthrop/foreversixty/companion/internal/state"
)

// morrowlyn is the addon's record of the fixture's logging character.
var morrowlyn = logger.Record{Name: "Morrowlyn", Realm: "Hardcore", Region: "us", Class: "paladin"}

var morrowlynChar = character.Character{Region: "us", Ruleset: "hardcore", Name: "Morrowlyn"}

// play writes the log's opening and n fights, ticking after each.
func (r *rig) play(n int) {
	r.t.Helper()
	r.write(fixture.Header + fixture.Zone)
	if err := r.tick(); err != nil {
		r.t.Fatal(err)
	}
	for i := range n {
		r.write(fixture.Encounter(i) + fixture.Heartbeat(i))
		if err := r.tick(); err != nil {
			r.t.Fatal(err)
		}
	}
}

func (r *rig) report() (created, now *character.Character, patches []json.RawMessage) {
	r.t.Helper()
	reps := r.srv.Reports()
	if len(reps) != 1 {
		r.t.Fatalf("the server has %d reports", len(reps))
	}
	for _, rep := range reps {
		return rep.Created, rep.LoggingCharacter, rep.Patches
	}
	return
}

func TestWithNothingChosenTheLogNamesTheCharacterAtCreation(t *testing.T) {
	r := newRig(t, 1<<20)
	r.records = []logger.Record{
		{Name: "Brannic", Realm: "Hardcore", Region: "eu"},
		morrowlyn,
	}
	r.play(1)

	created, now, patches := r.report()
	if created == nil || *created != morrowlynChar {
		t.Fatalf("the report was created for %+v, want %+v", created, morrowlynChar)
	}
	if now == nil || *now != morrowlynChar || len(patches) != 0 {
		t.Errorf("credited to %+v after %d patches", now, len(patches))
	}
	if got := r.pipe.Status().Detected; got != "Morrowlyn" {
		t.Errorf("Status.Detected = %q", got)
	}
}

func TestAChosenCharacterIsNeverOverriddenByTheLog(t *testing.T) {
	r := newRig(t, 1<<20)
	chosen := character.Character{Region: "eu", Ruleset: "pvp", Name: "Thalgrit"}
	r.chosen = &chosen
	r.open()
	r.records = []logger.Record{morrowlyn}
	r.play(2)

	created, now, patches := r.report()
	if created == nil || *created != chosen || *now != chosen || len(patches) != 0 {
		t.Errorf("created %+v, credited %+v, %d patches; want %+v throughout", created, now, len(patches), chosen)
	}
}

func TestACharacterTheAddonHasNotRecordedIsPatchedOnceWhenItArrives(t *testing.T) {
	r := newRig(t, 1<<20)
	r.play(1)

	created, now, patches := r.report()
	if created != nil || now != nil || len(patches) != 0 {
		t.Fatalf("credited to %+v before the addon named anyone", now)
	}
	// The log is Morrowlyn's, and the INFO line says what to do.
	for _, want := range []string{"the log names Morrowlyn, but the addon has not exported that character",
		"log out once with the addon loaded on it"} {
		if !strings.Contains(r.logs.String(), want) {
			t.Errorf("the log has no %q: %s", want, r.logs.String())
		}
	}
	if got := r.pipe.Status().Detected; got != "Morrowlyn" {
		t.Errorf("Status.Detected = %q before the record arrives", got)
	}
	if n := strings.Count(r.logs.String(), "the log names Morrowlyn"); n != 1 {
		t.Errorf("the notice was logged %d times", n)
	}

	// The player logs out; the record arrives; the next drain patches.
	r.records = []logger.Record{morrowlyn}
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}
	created, now, patches = r.report()
	if created != nil {
		t.Errorf("created with %+v", created)
	}
	if now == nil || *now != morrowlynChar {
		t.Fatalf("credited to %+v after the record arrived", now)
	}
	if len(patches) != 1 {
		t.Fatalf("%d patches, want one", len(patches))
	}
	var body client.UpdateReport
	if err := json.Unmarshal(patches[0], &body); err != nil || body.LoggingCharacter == nil || *body.LoggingCharacter != morrowlynChar {
		t.Errorf("patch body = %s (%v)", patches[0], err)
	}

	// Once is once: more drains, more fights, no second patch.
	r.write(fixture.Encounter(1) + fixture.Heartbeat(1))
	r.settle()
	if _, _, patches = r.report(); len(patches) != 1 {
		t.Errorf("%d patches after more fights", len(patches))
	}
}

func TestAFailedPatchIsRetriedBecauseTheNetworkMayReturn(t *testing.T) {
	r := newRig(t, 1<<20)
	r.play(1)
	r.records = []logger.Record{morrowlyn}
	r.srv.FailNext("PATCH /v1/reports/{id}", 1)
	if err := r.tick(); err == nil {
		t.Fatal("the drain hid a failed patch")
	}
	if _, now, _ := r.report(); now != nil {
		t.Fatalf("credited to %+v by a failed patch", now)
	}
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}
	if _, now, patches := r.report(); now == nil || len(patches) != 1 {
		t.Errorf("credited to %+v after %d patches", now, len(patches))
	}
}

func TestTheDetectionSurvivesARestart(t *testing.T) {
	r := newRig(t, 1<<20)
	r.play(1)
	if got := r.pipe.Status().Detected; got != "Morrowlyn" {
		t.Fatalf("Status.Detected = %q", got)
	}
	r.open() // the companion restarts mid-report
	if got := r.pipe.Status().Detected; got != "Morrowlyn" {
		t.Errorf("Status.Detected = %q after the restart", got)
	}
	reps, err := state.List(r.dirs.state)
	if err != nil || len(reps) != 1 {
		t.Fatalf("state: %v, %v", reps, err)
	}
	if got := reps[0]; !got.Auto || got.DetectedName != "Morrowlyn" || got.DetectedRealm != "Nightslayer" {
		t.Errorf("the state file holds auto=%v name=%q realm=%q", got.Auto, got.DetectedName, got.DetectedRealm)
	}
	// And it still resolves after the restart.
	r.records = []logger.Record{morrowlyn}
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}
	if _, now, _ := r.report(); now == nil || *now != morrowlynChar {
		t.Errorf("credited to %+v after the restart", now)
	}
}

func TestTheLastDetectionIsKeptWhenNothingIsOpen(t *testing.T) {
	r := newRig(t, 1<<20)
	r.records = []logger.Record{morrowlyn}
	r.play(1)
	r.settle()
	st := r.pipe.Status()
	if st.Logging {
		t.Fatal("the report did not close")
	}
	if st.Detected != "Morrowlyn" {
		t.Errorf("Status.Detected = %q once the report closed", st.Detected)
	}
	r.open()
	if got := r.pipe.Status().Detected; got != "Morrowlyn" {
		t.Errorf("Status.Detected = %q after a restart with nothing open", got)
	}
}

// An automatic report is not created until the log has named its
// logger, so the character goes up with the creation rather than by a
// patch behind it.
func TestAnAutomaticReportWaitsForTheLoggerBeforeItIsCreated(t *testing.T) {
	r := newRig(t, 1<<20)
	r.records = []logger.Record{morrowlyn}
	r.write(fixture.Header + fixture.Zone)
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}
	if n := len(r.srv.Reports()); n != 0 {
		t.Fatalf("%d reports on the server before the log named anyone", n)
	}
	r.write(fixture.Encounter(0))
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}
	created, _, patches := r.report()
	if created == nil || *created != morrowlynChar || len(patches) != 0 {
		t.Errorf("created for %+v after %d patches", created, len(patches))
	}
}

func TestACharacterTheAddonHasRecordedOnAnotherRealmIsToldApartFromOneItHasNot(t *testing.T) {
	r := newRig(t, 1<<20)
	// Two Morrowlyns, neither on the log's realm: the name is known, the realm is what is missing.
	r.records = []logger.Record{
		{Name: "Morrowlyn", Realm: "Hardcore", Region: "us", Class: "paladin"},
		{Name: "Morrowlyn", Realm: "Hardcore", Region: "eu", Class: "paladin"},
	}
	r.play(1)

	if _, now, _ := r.report(); now != nil {
		t.Fatalf("credited to %+v with two candidates", now)
	}
	want := "the log names Morrowlyn, but no addon record carries that realm"
	if !strings.Contains(r.logs.String(), want) || strings.Contains(r.logs.String(), "has not exported") {
		t.Errorf("the log has no %q: %s", want, r.logs.String())
	}
}

// storeDetected rewrites the open report's stored detection, the way
// an older companion would have left it, and restarts the pipeline.
func (r *rig) storeDetected(name, realm string) {
	r.t.Helper()
	reps, err := state.List(r.dirs.state)
	if err != nil || len(reps) != 1 {
		r.t.Fatalf("state: %v, %v", reps, err)
	}
	rep := reps[0]
	rep.DetectedName, rep.DetectedRealm, rep.DetectedClass = name, realm, ""
	if err := state.Save(r.dirs.state, rep); err != nil {
		r.t.Fatal(err)
	}
	r.open()
}

// A report opened by companion 0.2.5 holds "Bow-Jackzon"; the new
// rule reads the same log unit as "Bow Jackzon", and the report must
// be credited to that character after the update rather than ask for
// an export the addon already made.
func TestAReportStoredUnderTheOldTwoWordNameIsCreditedAfterTheUpdate(t *testing.T) {
	r := newRig(t, 1<<20)
	r.play(1)
	r.storeDetected("Bow-Jackzon", "ClassicBetaPvP")
	if got := r.pipe.Status().Detected; got != "Bow Jackzon" {
		t.Errorf("Status.Detected = %q after the update", got)
	}
	r.records = []logger.Record{{Name: "Bow Jackzon", Realm: "ClassicBetaPvP", Region: "us"}}
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}
	want := character.Character{Region: "us", Ruleset: "pvp", Name: "Bow Jackzon"}
	if _, now, _ := r.report(); now == nil || *now != want {
		t.Errorf("credited to %+v, want %+v", now, want)
	}
}

func TestAStoredLoggerNameIsReReadOnceAndOnlyWhenTheOldSplitMadeIt(t *testing.T) {
	for _, tc := range []struct {
		name, realm         string
		wantName, wantRealm string
	}{
		{"Bow-Jackzon", "ClassicBetaPvP", "Bow Jackzon", "ClassicBetaPvP"},
		{"Bow Jackzon", "ClassicBetaPvP", "Bow Jackzon", "ClassicBetaPvP"},
		{"Xuen", "ClassicBetaPvP", "Xuen", "ClassicBetaPvP"},
		{"Bow-Jackzon", "", "Bow-Jackzon", ""},
	} {
		t.Run(tc.name+"/"+tc.realm, func(t *testing.T) {
			r := newRig(t, 1<<20)
			r.play(1)
			r.storeDetected(tc.name, tc.realm)
			// A second restart finds nothing more to change.
			r.open()
			reps, err := state.List(r.dirs.state)
			if err != nil || len(reps) != 1 {
				t.Fatalf("state: %v, %v", reps, err)
			}
			if got := reps[0]; got.DetectedName != tc.wantName || got.DetectedRealm != tc.wantRealm {
				t.Errorf("stored name %q realm %q, want %q and %q",
					got.DetectedName, got.DetectedRealm, tc.wantName, tc.wantRealm)
			}
		})
	}
}
