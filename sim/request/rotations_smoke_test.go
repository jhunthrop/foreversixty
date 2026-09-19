package request

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"sort"
	"sync"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
	"github.com/jhunthrop/foreversixty/sim/specs"
	engine "github.com/wowsims/classic/sim"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/mage"
	"github.com/wowsims/classic/sim/warrior"
)

// The rotation smoke test: every rotation the rotation lane has written
// is run through the engine this module is linked against, and asked
// three questions a priority list nobody ran cannot answer.
//
//  1. Does it deal damage? Zero DPS is a rotation that does nothing.
//  2. Did the character act at all? An empty cast set is a priority list
//     that never reached the engine.
//  3. Are the actions the engine could not resolve exactly the ones we
//     expect it not to resolve? The engine reports every unresolvable
//     reference in a priority list as a validation warning when the
//     rotation is parsed, and this test holds that set to an exact
//     expectation. Half of the expectation is the curated file's own
//     `inert` array - the spell ids the rotation lane's notes declare
//     this engine build cannot act on at all - and half is
//     smokeBuildWarnings below, the references this deliberately bare
//     character cannot resolve no matter what the rotation says.
//
// The run is bare on purpose: the class's allowed race, no gear, no
// consumables, no buffs, and no talents except where the class package
// ships a Forever reference build. What is measured is whether the
// priority list runs, not what it scores.
//
// A mismatch in (3) is never fixed by editing data/curated/apl here. It
// is a finding for the rotation lane: either the engine moved and a
// declaration went stale, or a line nobody meant to lose is being
// skipped.

const (
	// Enough iterations that a rarely-reached line still gets its turn,
	// few enough that twenty specs run inside a normal test timeout.
	// Not one of api.ValidIterations, which is why the request is built
	// with OpenIterations - see request.Options.
	smokeIterations = 200
	// One seed for every spec, so a failure reproduces.
	smokeSeed = 1
)

// repoRoot is the site repository from this package's directory. The
// curated rotations, the spec list and the build's own race and class
// tables all live there; this test reads the real files rather than a
// testdata copy, because a copy of the rotation lane's `inert`
// declaration would be one more thing that can drift from it.
const repoRoot = "../.."

// registerEngine registers every spec's agent factory exactly once for
// the whole test binary. The engine's own guard is a sync.Once too, but
// naming it here says why a test needs it at all: without the registry
// the engine has no agent to build and every spec fails identically.
var registerEngine = sync.Once{}

// curatedRotation is the half of data/curated/apl/<spec>.json this test
// reads. The rotation itself is not read here: sim/request embeds its
// own synced copy and `make apl-check` is what proves the two agree, so
// reading the curated rotation would test the copy rather than use it.
type curatedRotation struct {
	Spec  string `json:"spec"`
	State string `json:"state"`
	Inert []int  `json:"inert"`
}

// writtenRotation pairs a spec with what its curated file declares.
type writtenRotation struct {
	spec  string
	inert []int
}

// actionInWarning pulls the action out of an engine validation warning.
// The engine formats every unresolved reference through
// core.ActionID.String - see sim/core/agent.go - so "{SpellID: 400574}"
// and "{OtherID: 13}" are the engine's own rendering, not one we chose,
// and using it whole means a non-spell action is expressible too.
var actionInWarning = regexp.MustCompile(`\{[A-Za-z]+ID: \d+(?:, Tag: \d+)?\}`)

// spellAction renders a curated `inert` spell id the way the engine
// renders it in a warning, so the two sets are comparable.
func spellAction(spellID int) string { return fmt.Sprintf("{SpellID: %d}", spellID) }

// smokeBuildWarnings is what THIS character cannot resolve, as opposed
// to what the rotation names that no character could.
//
// A bare smoke build has no talents, no consumables and no raid
// debuffs, and the engine registers a great many abilities only when
// one of those is present: Mortal Strike is a talent, the potion action
// needs a potion, the Improved Scorch debuff is a raid setting. Those
// warnings say nothing about the rotation, so they are expected here
// rather than mistaken for inert lines - and expected EXACTLY, so a
// warning nobody has accounted for still fails the test.
//
// The reason is carried with each entry because that is the whole value
// of the table: an id with no explanation is indistinguishable from a
// rotation quietly losing a line. Entries marked FINDING are not
// artefacts of this build at all - they are real defects reported to
// the rotation and engine lanes, recorded here so the suite is green
// while they are fixed elsewhere, and deleted when they are.
//
// Pinned to the engine at sim/enginever.Version.
var smokeBuildWarnings = map[string]map[string]string{
	"hunter-beast-mastery": {
		"{SpellID: 19574}": "Bestial Wrath is a talent (sim/hunter/talents.go); this build takes none.",
	},
	"mage-arcane": {
		"{SpellID: 12042}": "Arcane Power is a talent (sim/mage/talents.go); this build takes none.",
		"{SpellID: 12043}": "Presence of Mind is a talent (sim/mage/talents.go); this build takes none.",
	},
	"mage-fire": {
		"{SpellID: 11129}": "Combustion is a talent (sim/mage/talents.go); this build takes none.",
		"{SpellID: 12873}": "Improved Scorch's debuff aura only exists once sim/mage/talents.go's " +
			"applyImprovedScorch wires it up (Talents.ImprovedScorch != 0) or the raid debuff toggle " +
			"is on; this build takes neither, so core.GetAPLAura finds no such aura on the current " +
			"target and warns. Every auraIsActive/auraNumStacks/auraRemainingTime reference to it " +
			"resolves to a nil APLValue, which the Scorch line's `or` condition (newValueOr filters " +
			"nil operands the same way newValueAnd does) collapses to nil when every operand is nil, " +
			"and APLAction.IsReady reads a nil condition as always true -- so Scorch recasts every " +
			"global under this one build, the same shape as the hunters' zero-SwingSpeed auto-shot " +
			"loop: a real artifact of this bare build, not a defect in the gate.",
	},
	"priest-shadow": {
		"{SpellID: 15473}": "Shadowform is a talent (sim/priest/talents.go); this build takes none.",
		"{SpellID: 14751}": "Inner Focus is a talent (sim/priest/talents.go); this build takes none.",
		"{SpellID: 18807}": "Mind Flay is gated on its talent (sim/priest/mind_flay.go); this build takes none.",
	},
	"rogue-assassination": {
		"{SpellID: 14177}": "Cold Blood is a talent (sim/rogue/talents.go); this build takes none.",
	},
	"rogue-combat": {
		"{SpellID: 13750}": "Adrenaline Rush is a talent (sim/rogue/talents.go); this build takes none.",
	},
	"rogue-subtlety": {
		"{SpellID: 14183}": "Premeditation is a talent (sim/rogue/premeditation.go); this build takes none.",
		"{SpellID: 14278}": "Ghostly Strike is a talent (sim/rogue/ghostly_strike.go); this build takes none.",
		"{SpellID: 16511}": "Hemorrhage is a talent (sim/rogue/talents.go); this build takes none.",
	},
	"shaman-enhancement": {
		"{SpellID: 17364}": "Stormstrike is a talent (sim/shaman/stormstrike.go); this build takes none.",
	},
	"warlock-affliction": {
		"{OtherID: 13}":    noPotionWarning,
		"{SpellID: 18288}": "Amplify Curse is a talent (sim/warlock/curses.go); this build takes none.",
	},
	"warlock-demonology": {"{OtherID: 13}": noPotionWarning},
	"warlock-destruction": {
		"{OtherID: 13}":    noPotionWarning,
		"{SpellID: 18871}": "Shadowburn is gated on its talent (sim/warlock/shadowburn.go); this build takes none.",
		"{SpellID: 18932}": "Conflagrate is gated on its talent (sim/warlock/conflagrate.go); this build takes none.",
	},
	"warrior-arms": {
		"{SpellID: 21553}": "Mortal Strike is gated on its talent (sim/warrior/mortal_strike.go); this build takes none.",
	},
}

const noPotionWarning = "OtherActionPotion: this build carries no consumables, so there is " +
	"no potion for the engine to resolve the action against."

func TestEveryWrittenRotationRunsInTheEngine(t *testing.T) {
	registerEngine.Do(engine.RegisterAll)

	written := writtenRotations(t)
	if len(written) == 0 {
		t.Fatal("no curated rotation is marked written; the whole check is vacuous")
	}
	for _, rot := range written {
		t.Run(rot.spec, func(t *testing.T) {
			req := smokeRequest(t, rot.spec)

			// Two builds of the same request: core.ComputeStats builds
			// an environment of its own out of the raid it is handed,
			// and handing the same protobuf to the sim afterwards would
			// make the run depend on what the stats pass did to it.
			stats, err := BuildWith(req, Options{OpenIterations: true})
			if err != nil {
				t.Fatalf("building the request: %v", err)
			}
			warned := warnedActions(t, stats)
			want := expectedWarnings(rot)
			if !slices.Equal(warned, want) {
				t.Errorf("the engine could not resolve %v; expected %v.\n"+
					"An id here that the curated file does not declare inert and "+
					"smokeBuildWarnings does not explain is a finding for the rotation lane: "+
					"report it, do not edit the priority list.", warned, want)
			}

			engineReq, err := BuildWith(req, Options{OpenIterations: true})
			if err != nil {
				t.Fatalf("building the request: %v", err)
			}
			res := core.RunRaidSim(engineReq)
			if err := adapter.ResultError(res); err != nil {
				t.Fatalf("the sim failed: %v", err)
			}
			player, err := adapter.PlayerMetrics(res)
			if err != nil {
				t.Fatalf("reading the player's metrics: %v", err)
			}
			dps := player.Dps.GetAvg()
			if dps <= 0 {
				t.Errorf("DPS = %v; a rotation that deals no damage is not a rotation", dps)
			}
			cast := castSet(player)
			if len(cast) == 0 {
				t.Error("the player cast nothing; the priority list never fired a single action")
			}
			t.Logf("SMOKE\t%s\tdps=%.1f\tspells=%d\ttop=%s\tinert=%v\tunresolved=%v",
				rot.spec, dps, spellCasts(cast), top(cast, 3), rot.inert, warned)
		})
	}
}

// Every written rotation must also be a spec this module can build a
// request for. The two lists are maintained in different lanes - the
// rotation lane writes the curated file, this module carries the spec's
// options - so a rotation with no options here is a rotation nothing can
// run, and it would otherwise only show up as a missing subtest.
func TestEveryWrittenRotationHasSpecOptions(t *testing.T) {
	for _, rot := range writtenRotations(t) {
		if _, ok := specOptions[rot.spec]; !ok {
			t.Errorf("%s has a written rotation but no entry in specOptions; the engine cannot build an agent for it", rot.spec)
		}
	}
}

// smokeBuildWarnings is pinned to one engine build, so an entry for a
// spec nobody smokes is an entry nothing proves.
func TestSmokeBuildWarningsNameWrittenSpecs(t *testing.T) {
	written := map[string]bool{}
	for _, rot := range writtenRotations(t) {
		written[rot.spec] = true
	}
	for spec, warnings := range smokeBuildWarnings {
		if !written[spec] {
			t.Errorf("smokeBuildWarnings carries %q, which has no written rotation", spec)
		}
		for action, reason := range warnings {
			if reason == "" {
				t.Errorf("%s %s is expected with no reason given; an unexplained warning is indistinguishable from a lost line", spec, action)
			}
		}
	}
}

// expectedWarnings is everything this run is allowed not to resolve:
// what the curated file declares inert, plus what the bare build
// explains. Sorted, so the comparison is of sets.
func expectedWarnings(rot writtenRotation) []string {
	seen := map[string]bool{}
	for _, id := range rot.inert {
		seen[spellAction(id)] = true
	}
	for action := range smokeBuildWarnings[rot.spec] {
		seen[action] = true
	}
	return sortedKeys(seen)
}

// smokeRequest is the bare request described in this file's header.
func smokeRequest(t *testing.T, spec string) api.SimRequest {
	t.Helper()
	class := specs.ByKey[spec].ClassSlug
	return api.SimRequest{
		EngineVersion: enginever.Version,
		Spec:          spec,
		Source:        api.CharacterSource{Kind: api.Sources[0]},
		Character: api.CharacterSpec{
			Name:    spec,
			Race:    smokeRace(t, class),
			Class:   class,
			Level:   api.SimLevel,
			Talents: referenceTalents[spec],
		},
		Encounter:  api.DefaultEncounter(),
		Iterations: smokeIterations,
		RandomSeed: smokeSeed,
	}
}

// referenceTalents is the talent string a spec is smoked with.
//
// Empty is the default and the honest one: a talent string is a claim
// about a build, and this test makes none. The exceptions are the two
// class packages that ship a Forever reference build of their own - the
// same constants their own regression suites run - because for those
// two specs an untalented run would exercise a different rotation than
// the fork itself measures. They are read from the engine rather than
// copied, so a retuned build cannot go stale here.
var referenceTalents = map[string]string{
	"warrior-fury": warrior.ForeverFuryTalents,
	"mage-frost":   mage.ForeverFrostTalents,
}

// warnedActions is the set of actions the engine could not resolve
// while building this rotation, sorted and rendered the engine's own way.
//
// core.ComputeStats is what reports them: the warnings are attached to
// the player's stats, not to a sim result, because they are produced
// once when the priority list is parsed rather than per iteration.
//
// A warning that names no action at all is a failure rather than a
// skipped row. Every shape this test knows about - an unknown spell, an
// unregistered aura, a DoT that does not exist - names one, so a warning
// without one is the engine telling us something new.
func warnedActions(t *testing.T, req *proto.RaidSimRequest) []string {
	t.Helper()
	res := core.ComputeStats(&proto.ComputeStatsRequest{Raid: req.Raid, Encounter: req.Encounter})
	rotation := res.GetRaidStats().GetParties()[0].GetPlayers()[0].GetRotationStats()
	seen := map[string]bool{}
	for _, action := range slices.Concat(rotation.GetPrepullActions(), rotation.GetPriorityList()) {
		for _, warning := range action.GetWarnings() {
			match := actionInWarning.FindString(warning)
			if match == "" {
				t.Errorf("the engine warned without naming an action: %q", warning)
				continue
			}
			seen[match] = true
		}
	}
	return sortedKeys(seen)
}

func sortedKeys(set map[string]bool) []string {
	out := make([]string, 0, len(set))
	for key := range set {
		out = append(out, key)
	}
	slices.Sort(out)
	return out
}

// castCount is one action and how many times the player performed it
// per iteration.
type castCount struct {
	name  string
	casts float64
}

// castSet is every action the player actually performed, most first.
// Pets are left out: a hunter's cat attacking is not this rotation
// casting anything.
func castSet(player *proto.UnitMetrics) []castCount {
	var out []castCount
	for _, action := range player.GetActions() {
		var casts int32
		for _, target := range action.GetTargets() {
			casts += target.GetCasts()
		}
		if casts == 0 {
			continue
		}
		_, name := adapter.ActionName(action.GetId())
		out = append(out, castCount{name: name, casts: float64(casts) / smokeIterations})
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].casts != out[j].casts {
			return out[i].casts > out[j].casts
		}
		return out[i].name < out[j].name
	})
	return out
}

// spellCasts counts how many DISTINCT spells the priority list cast, as
// opposed to the autoattacks and resource gains the engine records
// alongside them. It is reported rather than asserted: a rotation whose
// every line is talent-gated casts no spell under this bare build, and
// that is a fact about the build, not about the rotation.
func spellCasts(cast []castCount) int {
	n := 0
	for _, c := range cast {
		if len(c.name) > 6 && c.name[:6] == "spell:" {
			n++
		}
	}
	return n
}

// top renders the n most-performed actions for the test log.
func top(cast []castCount, n int) string {
	out := ""
	for i, c := range cast[:min(n, len(cast))] {
		if i > 0 {
			out += ", "
		}
		out += fmt.Sprintf("%s=%.1f", c.name, c.casts)
	}
	return out
}

// writtenRotations reads every curated rotation the rotation lane has
// written, in sim/specs' canonical order.
func writtenRotations(t *testing.T) []writtenRotation {
	t.Helper()
	var out []writtenRotation
	for _, spec := range specs.All {
		path := filepath.Join(repoRoot, "data", "curated", "apl", spec.Spec+".json")
		b, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("reading %s: %v", path, err)
		}
		var doc curatedRotation
		if err := json.Unmarshal(b, &doc); err != nil {
			t.Fatalf("parsing %s: %v", path, err)
		}
		if doc.Spec != spec.Spec {
			t.Fatalf("%s declares spec %q", path, doc.Spec)
		}
		if doc.State != "written" {
			continue
		}
		out = append(out, writtenRotation{spec: spec.Spec, inert: doc.Inert})
	}
	return out
}

// smokeRaceOverrides names a race for a class the engine gates an
// ability on. sim/priest/priest.go registers Devouring Plague only for
// an undead priest, and the shadow rotation casts it, so smoking a
// human priest would report that line unresolvable and say nothing
// about the rotation. A class not named here takes the default below.
var smokeRaceOverrides = map[string]string{"priest": "undead"}

// smokeRace is a race the class is allowed to be.
//
// Read from the active build's own tables rather than from a table
// typed in here: Forever adds race and class pairs (the Deep Dive's six
// new combos), and a hard-coded pairing would be a second, staler copy
// of data/builds/<build>/combos.json. The lowest race id the class can
// be is picked, which is stable and, for every class, a vanilla race.
func smokeRace(t *testing.T, class string) string {
	t.Helper()
	build := activeBuild(t)
	type row struct {
		ID   int    `json:"id"`
		Slug string `json:"slug"`
	}
	var races, classes []row
	readBuildJSON(t, build, "races.json", &races)
	readBuildJSON(t, build, "classes.json", &classes)
	var combos []struct {
		RaceID  int `json:"race_id"`
		ClassID int `json:"class_id"`
	}
	readBuildJSON(t, build, "combos.json", &combos)

	classID := 0
	for _, c := range classes {
		if c.Slug == class {
			classID = c.ID
		}
	}
	if classID == 0 {
		t.Fatalf("build %s has no class %q", build, class)
	}
	want := smokeRaceOverrides[class]
	slug, lowest := "", 0
	for _, combo := range combos {
		if combo.ClassID != classID {
			continue
		}
		for _, r := range races {
			if r.ID != combo.RaceID {
				continue
			}
			// An override still has to be a pairing the build allows,
			// so it is chosen from this list rather than returned
			// before it.
			if r.Slug == want {
				return r.Slug
			}
			if lowest == 0 || r.ID < lowest {
				slug, lowest = r.Slug, r.ID
			}
		}
	}
	if want != "" {
		t.Fatalf("build %s does not pair %q with class %q", build, want, class)
	}
	if slug == "" {
		t.Fatalf("build %s pairs no race with class %q", build, class)
	}
	return slug
}

// activeBuild is the client build the site is serving, which is the one
// whose tables `make simdb` embeds.
func activeBuild(t *testing.T) string {
	t.Helper()
	path := filepath.Join(repoRoot, "web", "src", "data", "active-build.json")
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading %s: %v", path, err)
	}
	var active struct {
		Build string `json:"build"`
	}
	if err := json.Unmarshal(b, &active); err != nil {
		t.Fatalf("parsing %s: %v", path, err)
	}
	if active.Build == "" {
		t.Fatalf("%s names no build", path)
	}
	return active.Build
}

func readBuildJSON(t *testing.T, build, name string, into any) {
	t.Helper()
	path := filepath.Join(repoRoot, "data", "builds", build, name)
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading %s: %v", path, err)
	}
	if err := json.Unmarshal(b, into); err != nil {
		t.Fatalf("parsing %s: %v", path, err)
	}
}
