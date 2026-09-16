# Mechanics Mode Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** A Mechanics mode on the report page that lists, per fight and per player, the avoidable damage taken, the interrupts that went through and the debuffs that ran their course, from a curated per-encounter table, with a whole-night roll-up.

**Architecture:** Curated tables live inside the engine module as embedded JSON (one file per encounter id), so the parse job that runs the engine ships them. The engine attributes each fight's events to the table's mechanics at parse time and writes a `mechanics` block into `summary.json`; the web draws the mode from that block and the whole-night fold sums it. A CLI subcommand drafts a table for an encounter from a log, with the evidence a curator needs.

**Tech Stack:** Go 1.22+ (`logs/engine`, `go:embed`), Svelte 5 islands + Astro 7 (`web/`), vitest, Playwright.

**Spec:** `docs/superpowers/specs/2026-09-16-mechanics-mode-design.md`

## Global Constraints

- Work on a git worktree branch (`mechanics-mode`), never on `main`: the persona review loop runs a harness built from `main` and must not see half-built work. Merge to `main` only when the whole plan is done.
- Table location differs from the spec's `data/curated/mechanics/`: tables live at `logs/engine/mechanics/tables/<encounter_id>.json` so `go:embed` ships them with the engine. The format is the spec's, verbatim.
- `kind` is exactly one of `avoidable`, `unavoidable`, `interrupt`, `dispel`.
- Every summary change bumps `logs/engine/session/session.go` `Version` (currently `0.2.4` → `0.3.0`), regenerates goldens with `FOREVER_UPDATE_GOLDEN=1 go test ./...` in `logs/`, and the web fixture with `node scripts/make-report-fixture.mjs` in `web/`.
- Web copy rules: no "Mechanics later" note once the mode is enabled; every number has a unit or a word beside it; every new term goes in `web/src/components/report/Glossary.svelte`.
- The persona review harness serves the `main` build on port 4321. In the worktree, never run `scratchpad/local-report.sh`, and run Playwright against your own preview on port 4322: `E2E_PORT=4322 npx astro preview --port 4322` with `E2E_PORT=4322 npx playwright test ...`. The first task of whichever plan runs first makes `web/playwright.config.ts` read `process.env.E2E_PORT ?? '4321'` for both `use.baseURL` and `webServer.port`/`command` (one small commit: `chore(web): e2e port from E2E_PORT`).
- Run only the tests for files you touch (`go test ./engine/summary/`, `npx vitest run src/lib/report/<file>.test.ts`, `npx playwright test tests/e2e/report-mechanics.spec.ts`); CI runs the rest.
- After `npm run build` in `web/`, check `ls dist/report-island.js` exists: the island build runs as `postbuild` and a Svelte compile error there fails silently otherwise.
- Type check with `NO_COLOR=1 npx astro check 2>&1 | grep -E "^- [0-9]+ errors?"` and read the count; colour codes defeat a grep for the word "error".

---

### Task 1: The mechanics table package with embedded tables

**Files:**
- Create: `logs/engine/mechanics/mechanics.go`
- Create: `logs/engine/mechanics/mechanics_test.go`
- Create: `logs/engine/mechanics/tables/2363.json` (General Kaal, Sanguine Depths)
- Create: `logs/engine/mechanics/tables/2360.json` (Kryxis the Voracious)
- Create: `logs/engine/mechanics/tables/9001.json` (Warden Kelthas, the web fixture's boss)

**Interfaces:**
- Produces: `package mechanics`; `type Kind string` with constants `Avoidable`, `Unavoidable`, `Interrupt`, `Dispel`; `type Mechanic struct { SpellID int64; Name string; Kind Kind; Note string }`; `type Table struct { EncounterID int64; Name string; Mechanics []Mechanic }`; `func Parse(data []byte) (Table, error)`; `func Load(encounterID int64) (Table, bool)`; `func (t Table) Lookup(spellID int64) (Mechanic, bool)`.

- [ ] **Step 1: Write the failing tests**

```go
// logs/engine/mechanics/mechanics_test.go
package mechanics

import "testing"

func TestParseAcceptsTheSpecFormatAndRejectsAnUnknownKind(t *testing.T) {
	good := []byte(`{"encounter_id": 2363, "name": "General Kaal", "mechanics": [
		{"spell_id": 331415, "name": "Wicked Gash", "kind": "avoidable", "note": "Frontal."}]}`)
	table, err := Parse(good)
	if err != nil {
		t.Fatal(err)
	}
	if table.EncounterID != 2363 || len(table.Mechanics) != 1 || table.Mechanics[0].Kind != Avoidable {
		t.Fatalf("table = %+v", table)
	}
	if _, ok := table.Lookup(331415); !ok {
		t.Fatal("Lookup by spell id must find the mechanic")
	}
	bad := []byte(`{"encounter_id": 1, "name": "X", "mechanics": [{"spell_id": 5, "name": "Y", "kind": "soak"}]}`)
	if _, err := Parse(bad); err == nil {
		t.Fatal("an unknown kind must be refused")
	}
	zero := []byte(`{"encounter_id": 1, "name": "X", "mechanics": [{"spell_id": 0, "name": "Y", "kind": "avoidable"}]}`)
	if _, err := Parse(zero); err == nil {
		t.Fatal("a spell id of zero must be refused")
	}
}

func TestLoadFindsAnEmbeddedTableAndSaysWhenThereIsNone(t *testing.T) {
	if table, ok := Load(2363); !ok || table.Name != "General Kaal" {
		t.Fatalf("Load(2363) = %+v, %v", table, ok)
	}
	if _, ok := Load(424242); ok {
		t.Fatal("an encounter without a table must report false")
	}
}
```

- [ ] **Step 2: Run the tests to verify they fail**

Run: `cd logs && go test ./engine/mechanics/`
Expected: FAIL, package does not exist.

- [ ] **Step 3: Write the package**

```go
// logs/engine/mechanics/mechanics.go
// Package mechanics holds the curated per-encounter tables that say which of a
// boss's abilities are avoidable, which casts should be interrupted and which
// debuffs should be dispelled. The log only says an ability hit someone; a
// person decided that standing in it was a mistake. The tables ship inside the
// engine so the parse job has them; the drafting tool in cmd/forever-logs
// proposes one from a log, and a curator corrects it before it lands here.
package mechanics

import (
	"embed"
	"encoding/json"
	"fmt"
	"strconv"
)

// Kind is what failing a mechanic means.
type Kind string

const (
	// Avoidable damage should not have been taken.
	Avoidable Kind = "avoidable"
	// Unavoidable damage is the fight's, listed so it is not mistaken for a gap.
	Unavoidable Kind = "unavoidable"
	// Interrupt casts should have been stopped.
	Interrupt Kind = "interrupt"
	// Dispel debuffs should have been removed.
	Dispel Kind = "dispel"
)

var kinds = map[Kind]bool{Avoidable: true, Unavoidable: true, Interrupt: true, Dispel: true}

// Mechanic is one ability the table classifies.
type Mechanic struct {
	SpellID int64  `json:"spell_id"`
	Name    string `json:"name"`
	Kind    Kind   `json:"kind"`
	Note    string `json:"note,omitempty"`
}

// Table is one encounter's mechanics.
type Table struct {
	EncounterID int64      `json:"encounter_id"`
	Name        string     `json:"name"`
	Mechanics   []Mechanic `json:"mechanics"`
}

//go:embed tables/*.json
var tables embed.FS

// Parse reads a table and refuses one that is not the format.
func Parse(data []byte) (Table, error) {
	var t Table
	if err := json.Unmarshal(data, &t); err != nil {
		return Table{}, fmt.Errorf("mechanics: %w", err)
	}
	if t.EncounterID <= 0 {
		return Table{}, fmt.Errorf("mechanics: encounter_id must be positive")
	}
	for i, m := range t.Mechanics {
		if m.SpellID <= 0 {
			return Table{}, fmt.Errorf("mechanics[%d]: spell_id must be positive", i)
		}
		if !kinds[m.Kind] {
			return Table{}, fmt.Errorf("mechanics[%d]: kind %q is not avoidable, unavoidable, interrupt or dispel", i, m.Kind)
		}
	}
	return t, nil
}

// Load returns the embedded table for an encounter, and false when there is none.
func Load(encounterID int64) (Table, bool) {
	data, err := tables.ReadFile("tables/" + strconv.FormatInt(encounterID, 10) + ".json")
	if err != nil {
		return Table{}, false
	}
	t, err := Parse(data)
	if err != nil {
		// An embedded table that does not parse is a build defect, not a runtime case.
		panic(err)
	}
	return t, true
}

// Lookup finds a mechanic by spell id.
func (t Table) Lookup(spellID int64) (Mechanic, bool) {
	for _, m := range t.Mechanics {
		if m.SpellID == spellID {
			return m, true
		}
	}
	return Mechanic{}, false
}
```

- [ ] **Step 4: Write the three tables**

`logs/engine/mechanics/tables/2363.json` (from the sample log's General Kaal pulls; spell ids read from `fights/20/summary.json`'s damage taken):

```json
{
  "encounter_id": 2363,
  "name": "General Kaal",
  "mechanics": [
    { "spell_id": 331415, "name": "Wicked Gash", "kind": "avoidable", "note": "Frontal cleave. Only the tank should be in it." },
    { "spell_id": 322903, "name": "Gloom Squall", "kind": "avoidable", "note": "Room-wide unless behind a pillar." },
    { "spell_id": 323810, "name": "Piercing Blur", "kind": "avoidable", "note": "Charge along a line; step out of it." },
    { "spell_id": 0, "name": "Melee", "kind": "unavoidable", "note": "Tank damage." }
  ]
}
```

Before writing it, replace the `"spell_id": 0` Melee line: Melee is spell id 0 and `Parse` refuses zero, so leave Melee out of the table entirely (the mode lists unclassified damage separately). Confirm the three ids in `scratchpad/report-real/fights/20/summary.json` (`damage_taken[*].abilities[*].spell_id`).

`logs/engine/mechanics/tables/2360.json`:

```json
{
  "encounter_id": 2360,
  "name": "Kryxis the Voracious",
  "mechanics": [
    { "spell_id": 319713, "name": "Hungering Drain", "kind": "interrupt", "note": "Kick it; each tick heals the boss." },
    { "spell_id": 319654, "name": "Vicious Headbutt", "kind": "unavoidable", "note": "Tank damage." },
    { "spell_id": 319650, "name": "Severing Smash", "kind": "avoidable", "note": "Ground effect; move." }
  ]
}
```

Confirm each id against `fights/5/summary.json` (casts for the boss, damage taken for the raid) and correct any that differ; the names are what matter to a reader, the ids are what the engine matches.

`logs/engine/mechanics/tables/9001.json` (the web fixture's boss; read `web/src/fixtures/report/fixture.log` for the ids of the abilities Warden Kelthas casts; `Anima Lash` 334660 kills the tank in the fixture):

```json
{
  "encounter_id": 9001,
  "name": "Warden Kelthas",
  "mechanics": [
    { "spell_id": 334660, "name": "Anima Lash", "kind": "avoidable", "note": "Fixture: the hit that kills the tank." }
  ]
}
```

- [ ] **Step 5: Run the tests to verify they pass**

Run: `cd logs && go test ./engine/mechanics/`
Expected: PASS.

- [ ] **Step 6: Commit**

```bash
git add logs/engine/mechanics
git commit -m "feat(logs): curated mechanics tables, embedded in the engine"
```

---

### Task 2: The engine attributes each fight's events to its table

**Files:**
- Modify: `logs/engine/summary/summary.go` (Options, Accumulator fields, Snapshot)
- Create: `logs/engine/summary/mechanics.go`
- Modify: `logs/engine/summary/damage.go:150-160` (one call in the Damage case)
- Modify: `logs/engine/session/session.go:300` (set the table when a fight opens) and `Version`
- Test: `logs/engine/summary/mechanics_test.go`

**Interfaces:**
- Consumes: `mechanics.Load`, `mechanics.Table`, `mechanics.Kind` from Task 1; the existing `Death.KillingBlow`, `castRow` (started/succeeded per caster and spell), `exchangeRows("interrupt")`, `exchangeRows("dispel")`, `auraRows()`.
- Produces on `Summary`: `Mechanics MechanicsBlock \`json:"mechanics"\`` where

```go
type MechanicsBlock struct {
	// TableFound is false when the fight's encounter has no table; the block is then empty.
	TableFound bool           `json:"table_found"`
	Rows       []MechanicRow  `json:"rows"`
}
type MechanicRow struct {
	SpellID int64          `json:"spell_id"`
	Name    string         `json:"name"`
	Kind    mechanics.Kind `json:"kind"`
	Note    string         `json:"note,omitempty"`
	// Avoidable and unavoidable: who it hit.
	Players []MechanicHit `json:"players,omitempty"`
	// Interrupt: casts the enemies started and how many were stopped.
	Casts   int64 `json:"casts,omitempty"`
	Stopped int64 `json:"stopped,omitempty"`
	// Dispel: applications on players and how many were dispelled.
	Applied   int64 `json:"applied,omitempty"`
	Dispelled int64 `json:"dispelled,omitempty"`
}
type MechanicHit struct {
	GUID    string `json:"guid"`
	Name    string `json:"name"`
	Hits    int64  `json:"hits"`
	Damage  int64  `json:"damage"`
	FirstMS int64  `json:"first_ms"`
	LastMS  int64  `json:"last_ms"`
	// Killed is true when this player's killing blow was this mechanic.
	Killed bool `json:"killed"`
}
```

- `summary.Options` gains `Mechanics *mechanics.Table` (nil = no table).

- [ ] **Step 1: Write the failing test**

```go
// logs/engine/summary/mechanics_test.go
package summary

import (
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
)

// The build() fixture: the boss (guid `boss`) hits the tank with Anima Lash 334660
// twice, the second killing them; the mage casts Frostbolt; see summary_test.go.
func TestMechanicsAttributeAvoidableHitsAndTheKill(t *testing.T) {
	o, reg := opts(t)
	o.Mechanics = &mechanics.Table{EncounterID: 9001, Name: "Warden Kelthas", Mechanics: []mechanics.Mechanic{
		{SpellID: 334660, Name: "Anima Lash", Kind: mechanics.Avoidable},
	}}
	a := New(o)
	a.Start(at(0))
	for _, e := range fixtureEvents() { // the same events build() feeds; extract them if build() does not expose them
		reg.Observe(e)
		a.Add(e)
	}
	s := a.Snapshot(fixtureFight(), "test")
	if !s.Mechanics.TableFound || len(s.Mechanics.Rows) != 1 {
		t.Fatalf("mechanics = %+v", s.Mechanics)
	}
	row := s.Mechanics.Rows[0]
	if row.SpellID != 334660 || row.Kind != mechanics.Avoidable || len(row.Players) != 1 {
		t.Fatalf("row = %+v", row)
	}
	hit := row.Players[0]
	if hit.GUID != tank || hit.Hits != 2 || !hit.Killed {
		t.Fatalf("tank's hit = %+v, want two hits and the kill", hit)
	}
}

func TestMechanicsWithoutATableSaySo(t *testing.T) {
	_, _, s := build(t)
	if s.Mechanics.TableFound || len(s.Mechanics.Rows) != 0 {
		t.Fatalf("mechanics without a table = %+v", s.Mechanics)
	}
}
```

Read `summary_test.go`'s `build()` (line 126) first: if it feeds events inline, lift them into `fixtureEvents()` and the fight into `fixtureFight()` so both tests share them. Adjust the expected hit count to what the fixture actually does with Anima Lash (`dmg(...)` calls with spell 334660).

- [ ] **Step 2: Run the test to verify it fails**

Run: `cd logs && go test ./engine/summary/ -run TestMechanics`
Expected: FAIL, `o.Mechanics` undefined.

- [ ] **Step 3: Implement the accounting**

In `summary.go`: add to `Options`:

```go
	// Mechanics is the curated table for the fight's encounter; nil when there is none.
	Mechanics *mechanics.Table
```

Add to `Accumulator`: `mechanicHits map[int64]map[string]*MechanicHit` (spell id → player guid → hit), initialised in `New`. Add `Mechanics MechanicsBlock \`json:"mechanics"\`` to `Summary` and set `Mechanics: a.mechanicsBlock(deaths)` in `Snapshot` after `Deaths` is built (pass the death rows in, so `Killed` can be read off them).

Create `mechanics.go`:

```go
// logs/engine/summary/mechanics.go
package summary

import (
	"sort"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
)

// noteMechanicHit records a listed ability landing on a player. Called from the
// Damage case of addDamageAndHealing, after the tables are folded.
func (a *Accumulator) noteMechanicHit(e event.Event) {
	if a.opt.Mechanics == nil {
		return
	}
	m, ok := a.opt.Mechanics.Lookup(e.Spell.ID)
	if !ok || (m.Kind != mechanics.Avoidable && m.Kind != mechanics.Unavoidable) {
		return
	}
	u, known := a.opt.Registry.Get(e.Dest.GUID)
	if !known || !u.IsPlayer() {
		return
	}
	bySpell := a.mechanicHits[e.Spell.ID]
	if bySpell == nil {
		bySpell = map[string]*MechanicHit{}
		a.mechanicHits[e.Spell.ID] = bySpell
	}
	hit := bySpell[e.Dest.GUID]
	if hit == nil {
		hit = &MechanicHit{GUID: e.Dest.GUID, Name: a.name(e.Dest.GUID), FirstMS: a.ms(e.Time)}
		bySpell[e.Dest.GUID] = hit
	}
	hit.Hits++
	hit.Damage += e.Effective()
	hit.LastMS = a.ms(e.Time)
}

// mechanicsBlock renders the table's rows for the fight. Interrupt and dispel rows
// read the casts, exchanges and aura tracks the accumulator already keeps.
func (a *Accumulator) mechanicsBlock(deaths []Death) MechanicsBlock {
	if a.opt.Mechanics == nil {
		return MechanicsBlock{Rows: []MechanicRow{}}
	}
	killedBy := map[int64]map[string]bool{}
	for _, d := range deaths {
		if d.KillingBlow == nil {
			continue
		}
		if killedBy[d.KillingBlow.SpellID] == nil {
			killedBy[d.KillingBlow.SpellID] = map[string]bool{}
		}
		killedBy[d.KillingBlow.SpellID][d.GUID] = true
	}
	stopped := map[int64]int64{}
	for _, x := range a.exchangeRows("interrupt") {
		stopped[x.ExtraSpellID] += x.Count
	}
	dispelled := map[int64]int64{}
	for _, x := range a.exchangeRows("dispel") {
		dispelled[x.ExtraSpellID] += x.Count
	}
	rows := make([]MechanicRow, 0, len(a.opt.Mechanics.Mechanics))
	for _, m := range a.opt.Mechanics.Mechanics {
		row := MechanicRow{SpellID: m.SpellID, Name: m.Name, Kind: m.Kind, Note: m.Note}
		switch m.Kind {
		case mechanics.Avoidable, mechanics.Unavoidable:
			for _, hit := range a.mechanicHits[m.SpellID] {
				h := *hit
				h.Killed = killedBy[m.SpellID][h.GUID]
				row.Players = append(row.Players, h)
			}
			sort.Slice(row.Players, func(i, j int) bool { return row.Players[i].Damage > row.Players[j].Damage })
		case mechanics.Interrupt:
			for _, c := range a.castRows() {
				if c.SpellID == m.SpellID && !a.isPlayer(c.GUID) {
					row.Casts += max(c.Started, c.Succeeded)
				}
			}
			row.Stopped = stopped[m.SpellID]
			if row.Casts < row.Stopped {
				row.Casts = row.Stopped
			}
		case mechanics.Dispel:
			for _, tr := range a.auraRows() {
				if tr.SpellID == m.SpellID && tr.Type == "DEBUFF" && a.isPlayer(tr.TargetGUID) {
					row.Applied += tr.Applications
				}
			}
			row.Dispelled = dispelled[m.SpellID]
		}
		rows = append(rows, row)
	}
	return MechanicsBlock{TableFound: true, Rows: rows}
}

// isPlayer reports whether the registry knows the guid as a player.
func (a *Accumulator) isPlayer(guid string) bool {
	u, ok := a.opt.Registry.Get(guid)
	return ok && u.IsPlayer()
}
```

Check the exact names: `castRows()` and `exchangeRows(kind)` exist in `summary.go` (used by `Snapshot`); `ExchangeRow.ExtraSpellID` and `.Count` match `damage.go`/`deaths.go`'s `ExchangeRow` struct; `castRow` fields `Started`/`Succeeded`/`SpellID`/`GUID` match `CastRow`; `Registry.Get(guid) (*units.Unit, bool)` and `Unit.IsPlayer()` exist (`units/units.go`). Adjust names to what the code has; do not invent parallel structures.

In `damage.go`, in `case event.Damage:` after `a.threat[src] += ...` add `a.noteMechanicHit(e)`.

In `session.go` where a fight opens and `summary.New(o)` is called (line ~300): before `New`, set `o.Mechanics = nil` and, if the fight is an encounter, `if table, ok := mechanics.Load(f.EncounterID); ok { o.Mechanics = &table }`. Read the surrounding code to find the fight value in scope. Bump `Version` to `"0.3.0"`.

- [ ] **Step 4: Run the tests and regenerate goldens**

Run: `cd logs && go test ./engine/summary/ -run TestMechanics && FOREVER_UPDATE_GOLDEN=1 go test ./... && go test ./...`
Expected: PASS; the golden gains a `mechanics` block (`table_found: false, rows: []` for the retail sample's fights whose encounters have no table, and rows for 2360/2363).

- [ ] **Step 5: Commit**

```bash
git add logs/engine/summary logs/engine/session
git commit -m "feat(logs): the summary carries each fight's mechanics, from its encounter's table"
```

---

### Task 3: The drafting tool

**Files:**
- Modify: `logs/cmd/forever-logs/main.go` (new subcommand `mechanics-draft`)
- Create: `logs/cmd/forever-logs/mechanics_draft.go`
- Test: `logs/cmd/forever-logs/mechanics_draft_test.go`

**Interfaces:**
- Consumes: the session/parse path `parse` already uses (read `runParse` in `main.go` for how a log becomes fights and summaries; reuse it, do not re-implement parsing).
- Produces: `forever-logs mechanics-draft [-encounter <id>] <log>` printing one JSON table per encounter in the log (or the one asked for) to stdout, in the spec's format with a drafted `kind` and a `note` carrying the evidence: `"note": "draft: hit 3 non-tanks in one cast 4 times, same player twice 2 times, 12% of damage taken, killed 1"`.

- [ ] **Step 1: Write the failing test**

```go
// logs/cmd/forever-logs/mechanics_draft_test.go
package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
)

func TestMechanicsDraftProposesATableFromTheFixtureLog(t *testing.T) {
	var out, errOut bytes.Buffer
	err := run([]string{"mechanics-draft", "-encounter", "9001", "../../../web/src/fixtures/report/fixture.log"}, &out, &errOut)
	if err != nil {
		t.Fatal(err, errOut.String())
	}
	table, err := mechanics.Parse(out.Bytes())
	if err != nil {
		t.Fatalf("the draft must be a valid table: %v\n%s", err, out.String())
	}
	if table.EncounterID != 9001 {
		t.Fatalf("encounter = %d", table.EncounterID)
	}
	var lash *mechanics.Mechanic
	for i := range table.Mechanics {
		if table.Mechanics[i].SpellID == 334660 {
			lash = &table.Mechanics[i]
		}
	}
	if lash == nil || !strings.HasPrefix(lash.Note, "draft:") {
		t.Fatalf("Anima Lash must be drafted with evidence: %+v", table.Mechanics)
	}
	_ = json.Valid
}
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `cd logs && go test ./cmd/forever-logs/ -run TestMechanicsDraft`
Expected: FAIL, unknown subcommand.

- [ ] **Step 3: Implement the subcommand**

In `main.go`'s `run` switch add `case "mechanics-draft": return runMechanicsDraft(args[1:], out, errOut)`. In `mechanics_draft.go`, parse the log the way `runParse` does (read that function: it builds a session, feeds lines, collects fights and summaries), then per encounter fight compute, from the fight's `summary.DamageTaken` rows (players only, via the registry) and their abilities:

- per enemy spell id: total effective damage to players, share of all damage taken by players, number of distinct players hit, hits on the tank (roster role `tank`) vs others, and whether a death's killing blow was that spell;
- draft `kind`: `avoidable` when it hit two or more non-tank players, or hit a non-tank on more than one occasion (hits ≥ 2 on a non-tank); `unavoidable` when every hit landed on a tank; otherwise `avoidable` with the note saying "unclassified: check"; Melee (spell id 0) is skipped;
- `note`: `"draft: hit N players (M non-tanks), P% of damage taken, killed K"`.

Print `json.MarshalIndent` of a `mechanics.Table` per encounter, one after another, to `out`. With `-encounter`, only that one.

- [ ] **Step 4: Run the test to verify it passes**

Run: `cd logs && go test ./cmd/forever-logs/ -run TestMechanicsDraft`
Expected: PASS.

- [ ] **Step 5: Draft the sample tables and reconcile Task 1's files**

Run: `cd logs && go run ./cmd/forever-logs mechanics-draft /private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad/logs/wowp.txt > /tmp/claude-501/drafts.json`
Read the drafts for 2360 and 2363. Where a drafted spell id is missing from Task 1's tables and the evidence says avoidable, add it; where Task 1 guessed an id that never appears, remove it. Keep the notes hand-written.

- [ ] **Step 6: Commit**

```bash
git add logs/cmd/forever-logs logs/engine/mechanics/tables
git commit -m "feat(logs): forever-logs mechanics-draft proposes a table from a log"
```

---

### Task 4: Web types and the whole-night fold

**Files:**
- Modify: `web/src/lib/report/types.ts` (Summary gains `mechanics`)
- Modify: `web/src/lib/report/night.ts` (fold the block across pulls)
- Test: `web/src/lib/report/night.test.ts`

**Interfaces:**
- Consumes: the JSON shape from Task 2.
- Produces:

```ts
export type MechanicKind = 'avoidable' | 'unavoidable' | 'interrupt' | 'dispel';
export interface MechanicHit { guid: string; name: string; hits: number; damage: number; first_ms: number; last_ms: number; killed: boolean }
export interface MechanicRow { spell_id: number; name: string; kind: MechanicKind; note?: string; players?: MechanicHit[]; casts?: number; stopped?: number; applied?: number; dispelled?: number }
export interface MechanicsBlock { table_found: boolean; rows: MechanicRow[] }
// on Summary:
mechanics?: MechanicsBlock; // absent from summaries written before engine 0.3.0
```

and, on the night's folded `Summary`, `mechanics` folded per spell id with per-player hits summed, plus a new field on `MechanicRow` the fold sets: `pulls_hit?: number` (pulls on which the mechanic hit anyone) and on `MechanicHit`: `pulls?: number` (pulls on which this player was hit).

- [ ] **Step 1: Write the failing test** (in `night.test.ts`, alongside the existing fixtures `fights`/`summaries`; give two summaries a `mechanics` block with the same spell hitting the same player)

```ts
it('folds mechanics across pulls, counting the pulls each one hit anyone on', () => {
  const withMechanics = new Map(summaries);
  const block = (damage: number) => ({
    table_found: true,
    rows: [{ spell_id: 331415, name: 'Wicked Gash', kind: 'avoidable' as const,
      players: [{ guid: 'Player-1', name: 'Hobolol', hits: 1, damage, first_ms: 1000, last_ms: 1000, killed: false }] }],
  });
  withMechanics.set(3, { ...summaries.get(3)!, mechanics: block(100) });
  withMechanics.set(4, { ...summaries.get(4)!, mechanics: block(250) });
  const night = nightSummary(fights, withMechanics);
  const row = night.mechanics?.rows.find((entry) => entry.spell_id === 331415);
  expect(row?.pulls_hit).toBe(2);
  expect(row?.players?.[0]).toMatchObject({ guid: 'Player-1', hits: 2, damage: 350, pulls: 2 });
  expect(night.mechanics?.table_found).toBe(true);
});
```

Use whichever fight indexes the existing test fixtures use for the two pulls of one boss (read the file's `fights` constant).

- [ ] **Step 2: Run the test to verify it fails**

Run: `cd web && npx vitest run src/lib/report/night.test.ts`
Expected: FAIL, `mechanics` undefined on the night summary.

- [ ] **Step 3: Implement the fold** in `nightSummary`: keep `mechanicRows = new Map<number, MechanicRow>()`; for each pull with `summary.mechanics?.table_found`, for each row: merge by `spell_id` (sum `casts/stopped/applied/dispelled`, merge `players` by guid summing `hits`/`damage`, min `first_ms`+offset, max `last_ms`+offset, `killed` OR, `pulls` +1 per pull the player appears in), `pulls_hit` +1 when the row has any player, cast or application on that pull. Set `mechanics: { table_found: anyTable, rows: [...mechanicRows.values()] }` on the folded summary. Add the `pulls_hit`/`pulls` optional fields to the types.

- [ ] **Step 4: Run the test to verify it passes**

Run: `cd web && npx vitest run src/lib/report/night.test.ts`
Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git add web/src/lib/report/types.ts web/src/lib/report/night.ts web/src/lib/report/night.test.ts
git commit -m "feat(web): mechanics on the summary type, folded over the night"
```

---

### Task 5: The Mechanics mode

**Files:**
- Modify: `web/src/lib/report/url.ts:12,40` (`Mode` gains `'mechanics'`; the MODES entry becomes `enabled: true` with no note)
- Modify: `web/src/lib/report/url.test.ts:16` (the ids assertion still lists five; the enabled flag changes, so read the test and update whatever asserts on `enabled`/`note`)
- Create: `web/src/components/report/MechanicsMode.svelte`
- Modify: `web/src/components/report/ReportView.svelte:1296-1316` (mount beside Compare and Rankings; Mechanics mounts in night mode too, on the folded summary)
- Modify: `web/src/components/report/Glossary.svelte` (terms: "Avoidable damage", "Ran their course" if absent, "Went through" if absent)
- Test: `web/tests/e2e/report-mechanics.spec.ts`

**Interfaces:**
- Consumes: `scoped.mechanics` (Task 4's shape), `patch` from ReportView, `classOf`, `playerSet`, `nightMode`, `fight`.
- Produces: `MechanicsMode` props: `{ summary: Summary; classOf: Map<string, string>; nightMode: boolean; onPatch: (patch: Partial<ReportState>) => void; onDraft?: () => void }`.

- [ ] **Step 1: Write the failing e2e test**

```ts
// web/tests/e2e/report-mechanics.spec.ts
import { expect, test } from '@playwright/test';

const FIGHT = '/reports/fixture2abcd?fight=3';

test('mechanics lists the avoidable hit and the death it caused, and links to the detail', async ({ page }) => {
  await page.goto(`${FIGHT}&mode=mechanics`);
  await expect(page.getByTestId('mechanics-mode')).toBeVisible();
  const problems = page.getByTestId('mechanics-problems');
  await expect(problems).toContainText('Anima Lash');
  await expect(problems).toContainText('died to it');
  await problems.getByRole('button', { name: /Damage Taken/ }).first().click();
  await expect(page).toHaveURL(/tab=damage-taken/);
  await expect(page).toHaveURL(/ability=334660/);
});

test('a player card says what to tell them', async ({ page }) => {
  await page.goto(`${FIGHT}&mode=mechanics`);
  const card = page.getByTestId('mechanics-player-Player-4184-000000A1');
  await expect(card).toContainText('Anima Lash');
  await expect(card).toContainText(/avoidable/i);
});

test('a boss without a table says so and offers the draft', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=4&mode=mechanics');
  await expect(page.getByTestId('mechanics-no-table')).toContainText('No mechanics table');
});

test('the night rolls mechanics up per pull', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=all&mode=mechanics');
  await expect(page.getByTestId('mechanics-night')).toContainText('Anima Lash');
  await expect(page.getByTestId('mechanics-night')).toContainText(/pull/);
});
```

Check the fixture's guids and encounter ids in `web/src/fixtures/report/report.json`: fight 3 is encounter 9001 (Warden Kelthas), fight 4 is 9002 (no table), `Player-4184-000000A1` is the tank. Regenerate the fixture first (Task 2 changed the engine): `cd web && node scripts/make-report-fixture.mjs`.

- [ ] **Step 2: Run the test to verify it fails**

Run: `cd web && FOREVER_DATA=fixture npm run build && ls dist/report-island.js && (npx astro preview --port 4321 &) && sleep 4 && npx playwright test tests/e2e/report-mechanics.spec.ts`
Expected: FAIL, `mechanics-mode` not found (the mode is disabled).

- [ ] **Step 3: Enable the mode and write the component**

`url.ts`: `export type Mode = 'analyze' | 'compare' | 'rankings' | 'mechanics';` and `{ id: 'mechanics', label: 'Mechanics', enabled: true }`. `ModeOption.id` becomes `Mode | 'replay'`. Fix `url.test.ts` accordingly.

`MechanicsMode.svelte`:

```svelte
<!-- web/src/components/report/MechanicsMode.svelte -->
<!-- Who stood in what, and what it cost: the problems list a raid leader reads after a
     wipe, from the encounter's curated mechanics table, with a card per player and a
     roll-up over the night. Whole-fight figures: the Analyze window does not apply. -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import { classColorVar, formatAmount, formatDuration } from '../../lib/report/format';
  import type { ReportState } from '../../lib/report/url';
  import type { MechanicHit, MechanicRow, Summary } from '../../lib/report/types';

  let {
    summary,
    classOf,
    nightMode,
    onPatch,
  }: {
    summary: Summary;
    classOf: Map<string, string>;
    nightMode: boolean;
    onPatch: (patch: Partial<ReportState>) => void;
  } = $props();

  const block = $derived(summary.mechanics ?? { table_found: false, rows: [] });

  interface Problem {
    key: string;
    row: MechanicRow;
    hit?: MechanicHit;
    cost: number;
    text: string;
  }

  /** One line per failure, most costly first: a death outranks any amount of damage. */
  const problems = $derived.by<Problem[]>(() => {
    const out: Problem[] = [];
    for (const row of block.rows) {
      if (row.kind === 'avoidable') {
        for (const hit of row.players ?? []) {
          out.push({
            key: `${row.spell_id}-${hit.guid}`,
            row,
            hit,
            cost: hit.damage + (hit.killed ? 1e12 : 0),
            text: `${splitUnitName(hit.name).name} took ${row.name} ${hit.hits === 1 ? 'once' : `${hit.hits} times`} for ${formatAmount(hit.damage)}${hit.killed ? ' and died to it' : ''}`,
          });
        }
      } else if (row.kind === 'interrupt' && (row.casts ?? 0) > (row.stopped ?? 0)) {
        out.push({
          key: `${row.spell_id}-through`,
          row,
          cost: (row.casts ?? 0) - (row.stopped ?? 0),
          text: `${row.name} went through ${(row.casts ?? 0) - (row.stopped ?? 0)} of ${row.casts} times`,
        });
      } else if (row.kind === 'dispel' && (row.applied ?? 0) > (row.dispelled ?? 0)) {
        out.push({
          key: `${row.spell_id}-uncured`,
          row,
          cost: (row.applied ?? 0) - (row.dispelled ?? 0),
          text: `${row.name} ran its course ${(row.applied ?? 0) - (row.dispelled ?? 0)} of ${row.applied} times`,
        });
      }
    }
    return out.sort((a, b) => b.cost - a.cost);
  });

  /** Per player: their avoidable hits, most damage first, and the mechanics that never touched them. */
  const players = $derived.by(() => {
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const byGuid = new Map<string, { guid: string; name: string; hits: { row: MechanicRow; hit: MechanicHit }[]; damage: number }>();
    for (const row of block.rows) {
      if (row.kind !== 'avoidable') continue;
      for (const hit of row.players ?? []) {
        const found = byGuid.get(hit.guid) ?? { guid: hit.guid, name: hit.name, hits: [], damage: 0 };
        found.hits.push({ row, hit });
        found.damage += hit.damage;
        byGuid.set(hit.guid, found);
      }
    }
    for (const row of summary.roster) {
      if (!byGuid.has(row.guid)) byGuid.set(row.guid, { guid: row.guid, name: row.name, hits: [], damage: 0 });
    }
    return [...byGuid.values()].sort((a, b) => b.damage - a.damage);
  });

  const avoidableRows = $derived(block.rows.filter((row) => row.kind === 'avoidable'));

  function takenOf(guid: string): number {
    return summary.damage_taken.find((actor) => actor.guid === guid)?.effective ?? 0;
  }
  function openDamageTaken(spellId: number, guid?: string): void {
    onPatch({ mode: 'analyze', view: 'tables', tab: 'damage-taken', ability: spellId, ...(guid ? { source: guid } : {}) });
  }
</script>

<div class="flex flex-col gap-4" data-testid="mechanics-mode">
  {#if !block.table_found}
    <p class="text-muted text-[14px]" data-testid="mechanics-no-table">
      No mechanics table for this boss yet. A table says which of its abilities are avoidable, which casts to
      interrupt and which debuffs to dispel; without one this page has nothing to judge. Run
      <code class="font-mono text-[13px]">forever-logs mechanics-draft</code> on the log for a draft to review.
    </p>
  {:else}
    <p class="text-muted text-[12px]">
      Whole {nightMode ? 'night' : 'fight'}, from the encounter’s mechanics table. The time window above does not
      apply here.
    </p>

    <section class="flex flex-col gap-1" data-testid="mechanics-problems">
      <h2 class="label text-muted">Problems, most costly first</h2>
      {#if problems.length === 0}
        <p class="text-[14px]">Nothing the table lists went wrong.</p>
      {:else}
        <ol class="flex flex-col">
          {#each problems as problem (problem.key)}
            <li class="border-line-soft flex flex-wrap items-center gap-x-3 gap-y-1 border-b px-2 py-2 text-[14px]">
              <span class:text-death={problem.hit?.killed}>{problem.text}</span>
              {#if problem.row.note}<span class="text-muted text-[12px]">{problem.row.note}</span>{/if}
              {#if problem.row.kind === 'avoidable'}
                <button type="button" class="text-gold min-h-11 text-[12px] underline-offset-2 hover:underline md:min-h-0"
                  onclick={() => openDamageTaken(problem.row.spell_id, problem.hit?.guid)}>Damage Taken</button>
                {#if problem.hit?.killed}
                  <button type="button" class="text-gold min-h-11 text-[12px] underline-offset-2 hover:underline md:min-h-0"
                    onclick={() => onPatch({ mode: 'analyze', view: 'tables', tab: 'deaths' })}>Deaths</button>
                {/if}
              {:else if problem.row.kind === 'interrupt'}
                <button type="button" class="text-gold min-h-11 text-[12px] underline-offset-2 hover:underline md:min-h-0"
                  onclick={() => onPatch({ mode: 'analyze', view: 'tables', tab: 'interrupts' })}>Interrupts</button>
              {:else}
                <button type="button" class="text-gold min-h-11 text-[12px] underline-offset-2 hover:underline md:min-h-0"
                  onclick={() => onPatch({ mode: 'analyze', view: 'tables', tab: 'dispels' })}>Dispels</button>
              {/if}
            </li>
          {/each}
        </ol>
      {/if}
    </section>

    {#if nightMode}
      <section class="flex flex-col gap-1" data-testid="mechanics-night">
        <h2 class="label text-muted">Over the night</h2>
        <ul class="flex flex-col">
          {#each avoidableRows as row (row.spell_id)}
            <li class="border-line-soft border-b px-2 py-2 text-[14px]">
              <span class="font-semibold">{row.name}</span>
              hit someone on <span class="tabular font-mono">{row.pulls_hit ?? 0}</span> {row.pulls_hit === 1 ? 'pull' : 'pulls'}
              {#if (row.players ?? []).length > 0}
                · most often {splitUnitName([...(row.players ?? [])].sort((a, b) => (b.pulls ?? 0) - (a.pulls ?? 0))[0].name).name}
              {/if}
            </li>
          {/each}
        </ul>
      </section>
    {/if}

    <section class="grid grid-cols-1 gap-3 md:grid-cols-2" data-testid="mechanics-players">
      {#each players as player (player.guid)}
        <article class="border-line rounded-panel bg-raised flex flex-col gap-2 border p-3" data-testid={`mechanics-player-${player.guid}`}>
          <h3 class="font-semibold" style={`color: ${classColorVar(classOf.get(player.guid))}`}>{splitUnitName(player.name).name}</h3>
          <p class="text-[13px]">
            <span class="tabular font-mono">{formatAmount(player.damage)}</span> avoidable damage
            {#if takenOf(player.guid) > 0}
              · <span class="tabular font-mono">{Math.round((player.damage / takenOf(player.guid)) * 100)}%</span> of what they took
            {/if}
          </p>
          {#if player.hits.length === 0}
            <p class="text-muted text-[13px]">Clean: nothing avoidable landed on them.</p>
          {:else}
            <ul class="flex flex-col gap-1 text-[13px]">
              {#each player.hits as entry (entry.row.spell_id)}
                <li class:text-death={entry.hit.killed}>
                  {entry.row.name} · <span class="tabular font-mono">{entry.hit.hits}</span> {entry.hit.hits === 1 ? 'hit' : 'hits'} ·
                  <span class="tabular font-mono">{formatAmount(entry.hit.damage)}</span>
                  {#if !nightMode}· first at <span class="tabular font-mono">{formatDuration(entry.hit.first_ms)}</span>{/if}
                  {#if entry.hit.killed}· died to it{/if}
                  {#if entry.hit.pulls}· on {entry.hit.pulls} {entry.hit.pulls === 1 ? 'pull' : 'pulls'}{/if}
                </li>
              {/each}
            </ul>
          {/if}
        </article>
      {/each}
    </section>
  {/if}
</div>
```

Mount in `ReportView.svelte` after the Rankings block:

```svelte
      {#if state.mode === 'mechanics' && scoped !== null}
        <MechanicsMode summary={scoped} {classOf} {nightMode} onPatch={patch} />
      {/if}
```

`ModeBar.svelte` hides the mode tablist in night mode (`hidden={nightMode}`); Mechanics must be reachable over the night, so change that to hide only the modes that do not apply over the night: keep the tablist visible in night mode and disable Compare and Rankings there (`disabled={!option.enabled || (nightMode && option.id !== 'analyze' && option.id !== 'mechanics')}`). Check `patch()` in ReportView: it pushes history on `mode` changes already.

Glossary entries to add (`Glossary.svelte` TERMS): `Avoidable damage` — "Damage from an ability the boss's mechanics table says a player should not have been standing in. The table is curated by a person; the log only says what hit whom."; `Mechanics table` — "The per-boss list behind Mechanics mode: which abilities are avoidable, which casts to interrupt, which debuffs to dispel. A boss without one has no Mechanics page yet."

- [ ] **Step 4: Run the tests to verify they pass**

Run: `cd web && npx vitest run src/lib/report/url.test.ts && FOREVER_DATA=fixture npm run build && ls dist/report-island.js && npx playwright test tests/e2e/report-mechanics.spec.ts tests/e2e/report-tabs.spec.ts`
Expected: PASS. `report-tabs.spec.ts` covers the mode bar; if it asserts "Mechanics later", update it to the enabled mode.

- [ ] **Step 5: Lint and type check**

Run: `cd web && npx eslint src/components/report src/lib/report && NO_COLOR=1 npx astro check 2>&1 | grep -E "^- [0-9]+ errors?"`
Expected: no errors; `- 0 errors`.

- [ ] **Step 6: Commit**

```bash
git add web/src web/tests/e2e/report-mechanics.spec.ts
git commit -m "feat(web): mechanics mode: problems, player cards, the night"
```

---

### Task 6: Merge, re-parse, and hand to the personas

**Files:** none new.

- [ ] **Step 1: Merge the branch to main** (fast-forward or merge commit; run the full web e2e and `go test ./...` in `logs/` and `api/` first).
- [ ] **Step 2: Wait for API CI, then re-parse the production sample** with `gcloud run jobs execute parse-report --region us-east1 --args=parse-report,5lop7n5kwysf --wait`, so its summaries carry `mechanics`.
- [ ] **Step 3: Restage the review harness**: re-parse `scratchpad/logs/wowp.txt` with `forever-logs parse`, copy into `scratchpad/report-real` (keep `meta.json`), run `scratchpad/local-report.sh`.
- [ ] **Step 4: Dispatch the raid leader and tank personas** with the standing brief plus: "Mechanics mode is live for General Kaal and Kryxis; other bosses say they have no table yet." Fold their findings in.
