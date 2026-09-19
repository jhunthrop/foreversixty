# Simulator parity — engine fork lane — Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Teach the Forever engine fork the three encounter features the parity
design needs (scheduled movement, a target-count timeline, target-dummy mode)
and make it hand back the median-DPS iteration's cast log with resource
readings, so the site's fight styles and sample-iteration report have something
real underneath them.

**Architecture:** Three new `Encounter` proto fields (`movement`,
`targets_over_time`, `target_dummy`) parsed into the core `Encounter` struct in
`sim/core/target.go` and honoured by two new core files
(`sim/core/encounter_movement.go`, `sim/core/encounter_targets.go`) that hang
`PendingAction`s off `Simulation.reset`. The cast log is a recorder on
`Simulation` that `Spell.applyEffects` feeds, plus a two-pass replay: the normal
run records each iteration's DPS and its RNG seed, then one extra iteration is
re-run on a fresh `Simulation` at the median seed with the recorder on, so the
aggregated metrics are never polluted.

**Tech Stack:** Go 1.23.0, protobuf (protoc + protoc-gen-go v1.36.6), the
fork's own `go test --tags=with_db` suite.

**Repository:** `/Users/jh/code/wowsims-forever`, branch `forever`. Every task
in this plan is committed there, not in the site repository.

**Spec:** `/Users/jh/code/forever/docs/superpowers/specs/2026-09-19-simulator-parity-design.md`
(section 4.1 and section 5.1) and
`/Users/jh/code/forever/docs/superpowers/specs/2026-09-19-simulator-parity-interfaces.md`
(section 5 is this lane's binding interface; section 1.5 and section 2's
`SampleCast` are what this lane's output feeds).

## Global Constraints

- **Go 1.23.0 is fixed.** `go.mod` says `go 1.23.0` / `toolchain go1.23.4`.
  Never raise either, for any reason, including a tool telling you to.
- **Tests run with `--tags=with_db`.** `go test --tags=with_db ./sim/core/...`.
  Without the tag the item database is empty and database-dependent tests skip.
- **Never run `make update-tests`.** No task in this plan changes a checked-in
  `*.results` golden, and no task may regenerate one. If a `RunTestSuite`
  golden moves, that is a bug in the change, not a golden to refresh.
- **gofmt clean on touched files.** Run `gofmt -l <file>` on every file you
  create or modify and fix anything it names. CI does not gate gofmt (upstream's
  tree fails it) — you do.
- **Every proto change is regenerated in the same commit.** After editing any
  `proto/*.proto`, run `make proto` and commit `proto/*.proto`,
  `sim/core/proto/*.pb.go` and `ui/core/proto/*.ts` together.
  `go test --tags=with_db ./sim/core/proto/ -run TestGeneratedProtosMatchSources`
  is the gate and it must pass before the commit.
- **Commit messages** are `feat(core): ...` or `test(core): ...` (the fork's
  existing style), and each ends with the line:
  `Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>`
- **Never `git stash`.** The stash stack is shared across worktrees. Set work
  aside with a temporary WIP commit instead.
- **First task only:** run `make binary_dist/dist.go` once in a fresh checkout,
  or `./sim/...` will not build (`sim/web` embeds a generated, gitignored
  package).

---

### Task 1: The three `Encounter` proto fields

**Files:**
- Modify: `proto/common.proto:827-853` (the `Encounter` message)
- Regenerate: `sim/core/proto/common.pb.go`, `ui/core/proto/common.ts`
- Test: `sim/core/encounter_fields_test.go` (create)

**Interfaces:**
- Consumes: nothing.
- Produces: `proto.MovementPattern{IntervalSeconds, DurationSeconds float64; CastingOnly bool}`,
  `proto.TargetCountAt{AtSeconds float64; Count int32}`, and on `proto.Encounter`
  the fields `Movement *MovementPattern` (field 10), `TargetsOverTime []*TargetCountAt`
  (field 11), `TargetDummy bool` (field 12). Tasks 2, 4 and 6 read these.

Field numbers 1–7 and 20 are taken in `Encounter`; 10, 11 and 12 are free and
are the numbers the contract names, so use exactly those.

- [ ] **Step 1: Write the failing test**

Create `sim/core/encounter_fields_test.go`:

```go
package core

import (
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
)

// The three encounter fields the parity design adds are a wire contract
// with the site: the field numbers are pinned in
// docs/superpowers/specs/2026-09-19-simulator-parity-interfaces.md section 5,
// and a renumbering would silently misread every saved request. This test
// is cheap and it is the only thing that notices.
func TestEncounterCarriesTheParityFields(t *testing.T) {
	enc := &proto.Encounter{
		Duration: 180,
		Movement: &proto.MovementPattern{
			IntervalSeconds: 45,
			DurationSeconds: 5,
			CastingOnly:     true,
		},
		TargetsOverTime: []*proto.TargetCountAt{
			{AtSeconds: 0, Count: 1},
			{AtSeconds: 40, Count: 3},
		},
		TargetDummy: true,
	}

	if got := enc.Movement.GetIntervalSeconds(); got != 45 {
		t.Errorf("Movement.IntervalSeconds = %v, want 45", got)
	}
	if got := enc.Movement.GetDurationSeconds(); got != 5 {
		t.Errorf("Movement.DurationSeconds = %v, want 5", got)
	}
	if !enc.Movement.GetCastingOnly() {
		t.Error("Movement.CastingOnly = false, want true")
	}
	if len(enc.TargetsOverTime) != 2 || enc.TargetsOverTime[1].GetCount() != 3 {
		t.Errorf("TargetsOverTime = %v, want the two entries set above", enc.TargetsOverTime)
	}
	if !enc.GetTargetDummy() {
		t.Error("TargetDummy = false, want true")
	}
}

// The field numbers themselves, read off the descriptor rather than off
// the .proto text, so a renumbering fails here.
func TestEncounterParityFieldNumbers(t *testing.T) {
	fields := (&proto.Encounter{}).ProtoReflect().Descriptor().Fields()
	for name, want := range map[string]int32{
		"movement":          10,
		"targets_over_time": 11,
		"target_dummy":      12,
	} {
		field := fields.ByName(protoreflect.Name(name))
		if field == nil {
			t.Fatalf("Encounter has no field %q", name)
		}
		if got := int32(field.Number()); got != want {
			t.Errorf("Encounter.%s is field %d, want %d", name, got, want)
		}
	}
}
```

Add `"google.golang.org/protobuf/reflect/protoreflect"` to the import block.

- [ ] **Step 2: Run the test to verify it fails**

Run: `cd /Users/jh/code/wowsims-forever && go test --tags=with_db ./sim/core/ -run 'TestEncounterCarriesTheParityFields|TestEncounterParityFieldNumbers' -v`
Expected: FAIL — compile error, `enc.Movement undefined (type *proto.Encounter has no field or method Movement)`.

- [ ] **Step 3: Add the proto fields**

In `proto/common.proto`, inside `message Encounter`, after the `biome` field
(field 20) and before the closing brace, add:

```proto
	// A repeating window in which the player is away from the target.
	// interval_seconds is the gap between window starts; duration_seconds
	// is how long each window lasts. casting_only interrupts casting in
	// place instead of moving out of range, so melee continues.
	// Forever addition; see PORTING.md.
	MovementPattern movement = 10;

	// A timeline of how many of `targets` are active. Entries are
	// (at_seconds, count); the pool is sized to the largest count. When
	// set it overrides the plain target list's count.
	// Forever addition; see PORTING.md.
	repeated TargetCountAt targets_over_time = 11;

	// Target-dummy mode: the raid debuff panel is not applied, there is
	// no execute window, and nothing reduces the target's armor.
	// Forever addition; see PORTING.md.
	bool target_dummy = 12;
```

Directly above `message Encounter` (after the `Target` message ends at line
825), add the two new messages:

```proto
// MovementPattern is a repeating out-of-range window on an encounter.
message MovementPattern {
	double interval_seconds = 1;
	double duration_seconds = 2;
	bool casting_only = 3;
}

// TargetCountAt is one step of an encounter's target-count timeline.
message TargetCountAt {
	double at_seconds = 1;
	int32 count = 2;
}
```

- [ ] **Step 4: Regenerate the protobufs**

Run: `cd /Users/jh/code/wowsims-forever && make proto`
Then: `go test --tags=with_db ./sim/core/proto/ -run TestGeneratedProtosMatchSources -v`
Expected: PASS. If it skips with "protoc not installed", install protoc and
`go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.6` first — a
skip here is not a pass.

- [ ] **Step 5: Run the test to verify it passes**

Run: `go test --tags=with_db ./sim/core/ -run 'TestEncounterCarriesTheParityFields|TestEncounterParityFieldNumbers' -v`
Expected: PASS, both tests.

- [ ] **Step 6: gofmt and commit**

```bash
cd /Users/jh/code/wowsims-forever
gofmt -l sim/core/encounter_fields_test.go
git add proto/common.proto sim/core/proto ui/core/proto sim/core/encounter_fields_test.go
git commit -m "$(cat <<'MSG'
feat(core): Encounter carries movement, a target timeline and dummy mode

The three fields the parity contract's section 5 names, at the field
numbers it names, with the generated Go and TypeScript regenerated in the
same commit because the site consumes this module at a pinned sha.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
MSG
)"
```

---

### Task 2: Scheduled movement windows

**Files:**
- Modify: `sim/core/target.go:11-34` (the `Encounter` struct) and `:36-79` (`NewEncounter`)
- Modify: `sim/core/movement.go:70-90` (the `MovementHandler` struct and `initMovement`)
- Modify: `sim/core/spell.go:537` and `sim/core/cast.go:216` (the two cast gates)
- Modify: `sim/core/unit.go:504-530` (`Unit.reset`, to reset the handler)
- Modify: `sim/core/sim.go:420-422` (the tail of `Simulation.reset`)
- Create: `sim/core/encounter_movement.go`
- Test: `sim/core/encounter_movement_test.go` (create)

**Interfaces:**
- Consumes: `proto.Encounter.Movement` from Task 1.
- Produces:
  - `core.MovementPattern{Interval, Duration time.Duration; CastingOnly bool}`
  - `Encounter.Movement *MovementPattern`
  - `const EncounterMovementDistance = 12.0`
  - `func (unit *Unit) InterruptCast(sim *Simulation)`
  - `func (unit *Unit) IsCastingBlocked() bool`
  - `MovementHandler.CastingBlocked bool`
  - `func (sim *Simulation) initEncounterMovement()`
  Task 3 exercises these end to end.

- [ ] **Step 1: Write the failing test**

Create `sim/core/encounter_movement_test.go`:

```go
package core

import (
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core/proto"
)

// An encounter with no movement block is exactly today's encounter: the
// feature is additive and must not cost an unmoving fight anything.
func TestEncounterWithoutMovementHasNone(t *testing.T) {
	enc := NewEncounter(&proto.Encounter{
		Duration: 180,
		Targets:  []*proto.Target{DefaultTargetProtoLvl60},
	})
	if enc.Movement != nil {
		t.Errorf("Encounter.Movement = %+v, want nil for an encounter that sets none", enc.Movement)
	}
}

func TestEncounterParsesItsMovementPattern(t *testing.T) {
	enc := NewEncounter(&proto.Encounter{
		Duration: 180,
		Targets:  []*proto.Target{DefaultTargetProtoLvl60},
		Movement: &proto.MovementPattern{IntervalSeconds: 45, DurationSeconds: 5},
	})
	if enc.Movement == nil {
		t.Fatal("Encounter.Movement = nil, want a pattern")
	}
	if enc.Movement.Interval != 45*time.Second {
		t.Errorf("Interval = %s, want 45s", enc.Movement.Interval)
	}
	if enc.Movement.Duration != 5*time.Second {
		t.Errorf("Duration = %s, want 5s", enc.Movement.Duration)
	}
	if enc.Movement.CastingOnly {
		t.Error("CastingOnly = true, want false")
	}
}

// A zero interval or a zero duration is not a movement pattern, it is an
// unset one. Failing open here would schedule an infinite loop of
// zero-length windows.
func TestEncounterIgnoresADegenerateMovementPattern(t *testing.T) {
	for _, tc := range []struct {
		name    string
		pattern *proto.MovementPattern
	}{
		{"zero interval", &proto.MovementPattern{IntervalSeconds: 0, DurationSeconds: 5}},
		{"zero duration", &proto.MovementPattern{IntervalSeconds: 45, DurationSeconds: 0}},
		{"negative interval", &proto.MovementPattern{IntervalSeconds: -1, DurationSeconds: 5}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			enc := NewEncounter(&proto.Encounter{
				Duration: 180,
				Targets:  []*proto.Target{DefaultTargetProtoLvl60},
				Movement: tc.pattern,
			})
			if enc.Movement != nil {
				t.Errorf("Encounter.Movement = %+v, want nil", enc.Movement)
			}
		})
	}
}

// The away window puts the unit out of melee and stops it casting; when
// the window closes the unit is back at the distance it configured.
func TestAwayMovementWindowMovesOutAndBack(t *testing.T) {
	sim, unit := movementSim(t, &proto.MovementPattern{IntervalSeconds: 20, DurationSeconds: 5})
	startDistance := unit.DistanceFromTarget

	unit.startMovementWindow(sim, sim.Encounter.Movement)
	if !unit.IsMoving() {
		t.Error("unit.IsMoving() = false inside an away window, want true")
	}
	if !unit.IsCastingBlocked() {
		t.Error("unit.IsCastingBlocked() = false inside an away window, want true")
	}
	if unit.DistanceFromTarget != EncounterMovementDistance {
		t.Errorf("DistanceFromTarget = %v inside the window, want %v", unit.DistanceFromTarget, EncounterMovementDistance)
	}

	advanceSimTo(sim, 5*time.Second+time.Millisecond)
	if unit.IsMoving() {
		t.Error("unit.IsMoving() = true after the window closed, want false")
	}
	if unit.DistanceFromTarget != startDistance {
		t.Errorf("DistanceFromTarget = %v after the window, want the configured %v", unit.DistanceFromTarget, startDistance)
	}
}

// casting_only is the other half of the contract: casting stops, the unit
// never leaves melee, so autos keep swinging.
func TestCastingOnlyWindowBlocksCastingWithoutMoving(t *testing.T) {
	sim, unit := movementSim(t, &proto.MovementPattern{IntervalSeconds: 20, DurationSeconds: 5, CastingOnly: true})
	startDistance := unit.DistanceFromTarget

	unit.startMovementWindow(sim, sim.Encounter.Movement)
	if unit.IsMoving() {
		t.Error("unit.IsMoving() = true in a casting-only window, want false")
	}
	if !unit.IsCastingBlocked() {
		t.Error("unit.IsCastingBlocked() = false in a casting-only window, want true")
	}
	if unit.DistanceFromTarget != startDistance {
		t.Errorf("DistanceFromTarget = %v in a casting-only window, want the unchanged %v", unit.DistanceFromTarget, startDistance)
	}

	advanceSimTo(sim, 5*time.Second+time.Millisecond)
	if unit.IsCastingBlocked() {
		t.Error("unit.IsCastingBlocked() = true after the window closed, want false")
	}
}

// movementSim builds a one-player, one-target simulation with the given
// movement pattern, reset and ready to step.
func movementSim(t *testing.T, pattern *proto.MovementPattern) (*Simulation, *Unit) {
	t.Helper()
	sim := NewSim(&proto.RaidSimRequest{
		Raid: SinglePlayerRaidProto(&proto.Player{
			Name:               "Movement Test",
			Race:               proto.Race_RaceOrc,
			Class:              proto.Class_ClassShaman,
			Spec:               &proto.Player_ElementalShaman{ElementalShaman: &proto.ElementalShaman{}},
			Equipment:          &proto.EquipmentSpec{},
			DistanceFromTarget: 5,
		}, &proto.PartyBuffs{}, &proto.RaidBuffs{}, &proto.Debuffs{}),
		Encounter: &proto.Encounter{
			Duration: 180,
			Targets:  []*proto.Target{DefaultTargetProtoLvl60},
			Movement: pattern,
		},
		SimOptions: &proto.SimOptions{Iterations: 1, IsTest: true, RandomSeed: 1},
	}, simsignals.CreateSignals())
	sim.reset()
	return sim, &sim.Raid.Parties[0].Players[0].GetCharacter().Unit
}

// advanceSimTo steps the event loop until the clock passes `at`, so a
// delayed action scheduled inside a window actually fires.
func advanceSimTo(sim *Simulation, at time.Duration) {
	for sim.CurrentTime < at {
		if finished := sim.Step(); finished {
			return
		}
	}
}
```

Add `"github.com/wowsims/classic/sim/core/simsignals"` to the imports.

- [ ] **Step 2: Run the test to verify it fails**

Run: `cd /Users/jh/code/wowsims-forever && go test --tags=with_db ./sim/core/ -run 'Movement' -v`
Expected: FAIL — compile error, `enc.Movement undefined`, `unit.startMovementWindow undefined`, `EncounterMovementDistance undefined`.

- [ ] **Step 3: Parse the pattern into the core `Encounter`**

In `sim/core/target.go`, add to the `Encounter` struct, right after the `Biome`
field's block:

```go
	// Movement is a repeating window in which players are away from the
	// target (or, with CastingOnly, are interrupted in place). nil means
	// a stand-still fight, which is every encounter that predates the
	// parity work.
	Movement *MovementPattern
```

Directly under the `Encounter` struct, add:

```go
// MovementPattern is proto.MovementPattern in sim time units. Only a
// pattern with a positive interval and a positive duration is a pattern;
// anything else is an unset one, because a zero-length window repeated
// forever is not a fight, it is a hang.
type MovementPattern struct {
	Interval    time.Duration
	Duration    time.Duration
	CastingOnly bool
}

func newMovementPattern(options *proto.MovementPattern) *MovementPattern {
	if options == nil || options.IntervalSeconds <= 0 || options.DurationSeconds <= 0 {
		return nil
	}
	return &MovementPattern{
		Interval:    DurationFromSeconds(options.IntervalSeconds),
		Duration:    DurationFromSeconds(options.DurationSeconds),
		CastingOnly: options.CastingOnly,
	}
}
```

In `NewEncounter`, inside the `encounter := Encounter{...}` literal, add the
line `Movement: newMovementPattern(options.Movement),` after `Biome:`.

- [ ] **Step 4: Add the casting-blocked flag and the interrupt**

In `sim/core/movement.go`, add a field to `MovementHandler`:

```go
type MovementHandler struct {
	Moving    bool
	MoveSpeed float64

	// CastingBlocked is set while a casting-only movement window is open:
	// spells with a cast time are refused and melee continues. It is
	// separate from Moving because Moving also means "out of range", and
	// a casting-only window never leaves melee.
	CastingBlocked bool

	baseSpeed          float64
	moveAura           *Aura
	moveSpell          *Spell
	moveSpeedBonuses   *MoveHeap
	moveSpeedPenalties *MoveHeap
}
```

and, next to `IsMoving`, add:

```go
// IsCastingBlocked reports whether a cast with a cast time may start. It
// is what the cast gates ask; IsMoving stays the question "is this unit
// out of position", which item and talent code asks for other reasons.
func (unit *Unit) IsCastingBlocked() bool {
	return unit.MovementHandler.Moving || unit.MovementHandler.CastingBlocked
}

// InterruptCast cancels an in-progress hardcast or channel. The cost is
// already spent and is not refunded, which is what a real interrupt does
// and what a movement window models.
func (unit *Unit) InterruptCast(sim *Simulation) {
	if unit.IsChanneling(sim) {
		unit.ChanneledDot.Cancel(sim)
	}
	if !unit.IsCasting(sim) {
		return
	}
	if sim.Log != nil {
		unit.Log(sim, "Cast of %s interrupted", unit.Hardcast.ActionID)
	}
	unit.Hardcast = Hardcast{Expires: startingCDTime}
	if unit.hardcastAction != nil {
		unit.hardcastAction.Cancel(sim)
		unit.hardcastAction = nil
	}
}

// reset clears the per-iteration movement state. Unit.reset calls it, so
// a window left open by an aborted iteration cannot leak into the next.
func (move *MovementHandler) reset() {
	move.Moving = false
	move.CastingBlocked = false
}
```

- [ ] **Step 5: Point the two cast gates at the new question**

In `sim/core/spell.go:537`, change

```go
	if spell.DefaultCast.CastTime > 0 && spell.Unit.IsMoving() {
```

to

```go
	if spell.DefaultCast.CastTime > 0 && spell.Unit.IsCastingBlocked() {
```

In `sim/core/cast.go:216`, change

```go
		if (spell.CurCast.CastTime > 0) && spell.Unit.IsMoving() {
```

to

```go
		if (spell.CurCast.CastTime > 0) && spell.Unit.IsCastingBlocked() {
```

In `sim/core/unit.go`, inside `func (unit *Unit) reset(...)`, add
`unit.MovementHandler.reset()` immediately after the existing
`unit.DistanceFromTarget = unit.StartDistanceFromTarget` line.

- [ ] **Step 6: Write the scheduler**

Create `sim/core/encounter_movement.go`:

```go
package core

// EncounterMovementDistance is how far an away window puts a player from
// the target: past melee range (5) and at the minimum ranged distance
// (12), so the window costs melee and casting without also making ranged
// attacks illegal for a reason the fight style never asked for.
const EncounterMovementDistance = 12.0

// initEncounterMovement schedules the encounter's movement windows for
// every player, once per iteration. Simulation.reset calls it, beside
// initManaTickAction, because that is where per-iteration pending actions
// belong: reset has just emptied the queue and set sim.Duration.
func (sim *Simulation) initEncounterMovement() {
	pattern := sim.Encounter.Movement
	if pattern == nil {
		return
	}

	numTicks := int(sim.Duration / pattern.Interval)
	if numTicks <= 0 {
		return
	}

	for _, party := range sim.Raid.Parties {
		for _, player := range party.Players {
			unit := &player.GetCharacter().Unit
			sim.AddPendingAction(NewPeriodicAction(sim, PeriodicActionOptions{
				Period:          pattern.Interval,
				NumTicks:        numTicks,
				TickImmediately: false,
				Priority:        ActionPriorityAuto,
				OnAction: func(sim *Simulation) {
					unit.startMovementWindow(sim, pattern)
				},
			}))
		}
	}
}

// startMovementWindow opens one window on one unit and schedules its
// close. An away window activates the existing movement aura, which is
// what cancels auto attacks and channels; a casting-only window only
// interrupts the cast in progress and refuses new ones for the duration.
func (unit *Unit) startMovementWindow(sim *Simulation, pattern *MovementPattern) {
	unit.InterruptCast(sim)

	if pattern.CastingOnly {
		unit.MovementHandler.CastingBlocked = true
		if sim.Log != nil {
			unit.Log(sim, "Casting interrupted for %s", pattern.Duration)
		}
		StartDelayedAction(sim, DelayedActionOptions{
			DoAt:     sim.CurrentTime + pattern.Duration,
			Priority: ActionPriorityAuto,
			OnAction: func(sim *Simulation) {
				unit.MovementHandler.CastingBlocked = false
				unit.Rotation.DoNextAction(sim)
			},
		})
		return
	}

	unit.MovementHandler.moveAura.Activate(sim)
	unit.DistanceFromTarget = EncounterMovementDistance
	unit.MovementHandler.moveAura.SetStacks(sim, int32(EncounterMovementDistance))
	if sim.Log != nil {
		unit.Log(sim, "Moving out of range for %s", pattern.Duration)
	}
	StartDelayedAction(sim, DelayedActionOptions{
		DoAt:     sim.CurrentTime + pattern.Duration,
		Priority: ActionPriorityAuto,
		OnAction: func(sim *Simulation) {
			unit.DistanceFromTarget = unit.StartDistanceFromTarget
			unit.MovementHandler.moveAura.Deactivate(sim)
			unit.Rotation.DoNextAction(sim)
		},
	})
}
```

In `sim/core/sim.go`, at the end of `func (sim *Simulation) reset()`, change

```go
	sim.initManaTickAction()
```

to

```go
	sim.initManaTickAction()
	sim.initEncounterMovement()
```

- [ ] **Step 7: Run the test to verify it passes**

Run: `cd /Users/jh/code/wowsims-forever && go test --tags=with_db ./sim/core/ -run 'Movement' -v`
Expected: PASS, all five tests.

- [ ] **Step 8: Run the whole core suite so no golden moved**

Run: `go test --tags=with_db -count=1 ./sim/...`
Expected: PASS. An encounter with no `movement` block takes the `nil` path, so
every existing `*.results` golden must be unchanged. If one moved, the change
is wrong — do not regenerate it.

- [ ] **Step 9: gofmt and commit**

```bash
cd /Users/jh/code/wowsims-forever
gofmt -l sim/core/encounter_movement.go sim/core/encounter_movement_test.go sim/core/target.go sim/core/movement.go sim/core/spell.go sim/core/cast.go sim/core/unit.go sim/core/sim.go
git add sim/core/encounter_movement.go sim/core/encounter_movement_test.go sim/core/target.go sim/core/movement.go sim/core/spell.go sim/core/cast.go sim/core/unit.go sim/core/sim.go
git commit -m "$(cat <<'MSG'
feat(core): encounters schedule movement windows

An away window activates the movement aura, puts the player past melee
range and stops casting; a casting-only window interrupts the cast in
place and leaves melee swinging. The two cast gates now ask
IsCastingBlocked rather than IsMoving, so the casting-only case has a
question of its own and IsMoving keeps meaning "out of position".

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
MSG
)"
```

---

### Task 3: Movement, end to end, on a fury warrior and a frost mage

**Files:**
- Create: `sim/encounter_movement_e2e_test.go` (package `sim`)

**Interfaces:**
- Consumes: `proto.Encounter.Movement` (Task 1) and its behaviour (Task 2).
- Produces: `furyWarriorPlayer()`, `frostMagePlayer()`, `parityEncounter()` and
  `runParitySim()` — the shared end-to-end fixtures Tasks 5, 7 and 10 also use.
  They live in this file; later tasks import nothing, they are the same package.

These tests do **not** use `core.RunTestSuite`: that compares against a
checked-in `*.results` golden and this plan may never regenerate one. They build
a request, call `core.RunSim`, and assert a direction of travel.

- [ ] **Step 1: Write the failing test**

Create `sim/encounter_movement_e2e_test.go`:

```go
package sim

import (
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/mage"
	"github.com/wowsims/classic/sim/warrior"
)

// parityIterations is small on purpose: these are direction-of-travel
// checks over a three-minute fight, not a tuning fixture, and the whole
// file has to stay inside the fork's test budget.
const parityIterations = 30

// furyWarriorPlayer is the reference melee: rage, no cast times, so an
// away window costs it auto attacks and a casting-only window costs it
// nothing.
func furyWarriorPlayer() *proto.Player {
	return &proto.Player{
		Name:          "Fury Warrior",
		Race:          proto.Race_RaceOrc,
		Class:         proto.Class_ClassWarrior,
		Equipment:     core.GetGearSet("../ui/warrior/gear_sets", "phase_1").GearSet,
		TalentsString: warrior.ForeverFuryTalents,
		Rotation:      core.GetAplRotation("../ui/warrior/apls", "forever_fury").Rotation,
		Consumes:      &proto.Consumes{},
		Buffs:         &proto.IndividualBuffs{},
		Spec: &proto.Player_Warrior{Warrior: &proto.Warrior{
			Options: &proto.Warrior_Options{
				StartingRage: 50,
				Shout:        proto.WarriorShout_WarriorShoutBattle,
			},
		}},
		DistanceFromTarget: 5,
	}
}

// frostMagePlayer is the reference caster: mana and long cast times, so
// both window kinds cost it casts.
func frostMagePlayer() *proto.Player {
	return &proto.Player{
		Name:          "Frost Mage",
		Race:          proto.Race_RaceTroll,
		Class:         proto.Class_ClassMage,
		Equipment:     core.GetGearSet("../ui/mage/gear_sets", "p0.bis").GearSet,
		TalentsString: mage.ForeverFrostTalents,
		Rotation:      core.GetAplRotation("../ui/mage/apls", "forever_frost").Rotation,
		Consumes:      &proto.Consumes{},
		Buffs:         &proto.IndividualBuffs{},
		Spec: &proto.Player_Mage{Mage: &proto.Mage{
			Options: &proto.Mage_Options{Armor: proto.Mage_Options_MoltenArmor},
		}},
		DistanceFromTarget: 20,
	}
}

// parityEncounter is a plain three-minute Patchwerk with no duration
// variation, so two runs differ only by the feature under test.
func parityEncounter() *proto.Encounter {
	return &proto.Encounter{
		Duration: 180,
		Targets:  []*proto.Target{core.DefaultTargetProtoLvl60},
	}
}

// runParitySim runs one request and returns the raid's mean DPS.
func runParitySim(t *testing.T, player *proto.Player, encounter *proto.Encounter) float64 {
	t.Helper()
	result := core.RunSim(&proto.RaidSimRequest{
		Raid: core.SinglePlayerRaidProto(player, &proto.PartyBuffs{}, &proto.RaidBuffs{}, &proto.Debuffs{}),
		Encounter: encounter,
		SimOptions: &proto.SimOptions{
			Iterations: parityIterations,
			IsTest:     true,
			RandomSeed: 1,
		},
	}, nil, simsignals.CreateSignals())
	if result.Error != nil {
		t.Fatalf("sim failed: %s", result.Error.Message)
	}
	return result.RaidMetrics.Dps.Avg
}

// Heavy movement (5s away every 20s) is a quarter of the fight spent out
// of melee. A fury warrior's DPS must fall, and it must fall by less than
// the whole quarter, because rage carries across the window.
func TestHeavyMovementCostsTheFuryWarrior(t *testing.T) {
	still := runParitySim(t, furyWarriorPlayer(), parityEncounter())

	moving := parityEncounter()
	moving.Movement = &proto.MovementPattern{IntervalSeconds: 20, DurationSeconds: 5}
	movingDps := runParitySim(t, furyWarriorPlayer(), moving)

	if movingDps >= still {
		t.Errorf("heavy movement DPS %.1f is not below the standing DPS %.1f", movingDps, still)
	}
	if movingDps < still*0.5 {
		t.Errorf("heavy movement DPS %.1f is below half the standing DPS %.1f; a quarter of the fight away should not halve it", movingDps, still)
	}
}

// A casting-only window costs a melee nothing: it never leaves melee and
// it has no cast times to interrupt.
func TestCastingOnlyMovementIsFreeForTheFuryWarrior(t *testing.T) {
	still := runParitySim(t, furyWarriorPlayer(), parityEncounter())

	interrupted := parityEncounter()
	interrupted.Movement = &proto.MovementPattern{IntervalSeconds: 20, DurationSeconds: 5, CastingOnly: true}
	interruptedDps := runParitySim(t, furyWarriorPlayer(), interrupted)

	if interruptedDps != still {
		t.Errorf("casting-only DPS %.4f differs from the standing DPS %.4f; a melee with no cast times should be untouched", interruptedDps, still)
	}
}

// Both window kinds cost a frost mage, because every window interrupts a
// Frostbolt and refuses the next one for its duration.
func TestBothMovementKindsCostTheFrostMage(t *testing.T) {
	still := runParitySim(t, frostMagePlayer(), parityEncounter())

	away := parityEncounter()
	away.Movement = &proto.MovementPattern{IntervalSeconds: 20, DurationSeconds: 5}
	awayDps := runParitySim(t, frostMagePlayer(), away)

	castingOnly := parityEncounter()
	castingOnly.Movement = &proto.MovementPattern{IntervalSeconds: 20, DurationSeconds: 5, CastingOnly: true}
	castingOnlyDps := runParitySim(t, frostMagePlayer(), castingOnly)

	if awayDps >= still {
		t.Errorf("away-movement DPS %.1f is not below the standing DPS %.1f", awayDps, still)
	}
	if castingOnlyDps >= still {
		t.Errorf("casting-only DPS %.1f is not below the standing DPS %.1f", castingOnlyDps, still)
	}
}

// Light movement (5s every 45s) costs less than heavy movement (5s every
// 20s). This is the ordering the two fight styles promise the player.
func TestLightMovementCostsLessThanHeavy(t *testing.T) {
	light := parityEncounter()
	light.Movement = &proto.MovementPattern{IntervalSeconds: 45, DurationSeconds: 5}
	heavy := parityEncounter()
	heavy.Movement = &proto.MovementPattern{IntervalSeconds: 20, DurationSeconds: 5}

	lightDps := runParitySim(t, frostMagePlayer(), light)
	heavyDps := runParitySim(t, frostMagePlayer(), heavy)

	if lightDps <= heavyDps {
		t.Errorf("light-movement DPS %.1f is not above heavy-movement DPS %.1f", lightDps, heavyDps)
	}
}
```

- [ ] **Step 2: Prove the test can fail**

Task 2 already implements the behaviour, so a green run here proves nothing.
Comment out the `sim.initEncounterMovement()` line added to
`Simulation.reset` in Task 2 (edit in place — never `git stash`), then run:

`cd /Users/jh/code/wowsims-forever && go test --tags=with_db ./sim/ -run 'Movement' -v`

Expected: FAIL — `TestHeavyMovementCostsTheFuryWarrior` reports
"heavy movement DPS ... is not below the standing DPS ...".
Restore the line before Step 3.

- [ ] **Step 3: Run the test to verify it passes**

Run: `go test --tags=with_db ./sim/ -run 'Movement' -v`
Expected: PASS, all four tests.

- [ ] **Step 4: Commit**

```bash
cd /Users/jh/code/wowsims-forever
gofmt -l sim/encounter_movement_e2e_test.go
git add sim/encounter_movement_e2e_test.go
git commit -m "$(cat <<'MSG'
test(core): movement windows, end to end, on fury and frost

Two reference specs and four claims: away movement costs a melee, a
casting-only window does not; both cost a caster; and light movement
costs less than heavy, which is the ordering the fight styles promise.
No RunTestSuite and no golden, because this plan may not regenerate one.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
MSG
)"
```

---

### Task 4: The target-count timeline

**Files:**
- Modify: `sim/core/target.go:11-34` (the `Encounter` struct), `:36-79` (`NewEncounter`), `:84-87` (`updateAOECapMultiplier`)
- Modify: `sim/core/environment.go:78` (`env.AllUnits`), `:265-270` (`GetNumTargets`)
- Modify: `sim/core/sim.go` (the tail of `Simulation.reset`)
- Create: `sim/core/encounter_targets.go`
- Test: `sim/core/encounter_targets_test.go` (create)

**Interfaces:**
- Consumes: `proto.Encounter.TargetsOverTime` from Task 1.
- Produces:
  - `core.TargetCount{At time.Duration; Count int32}`
  - `Encounter.TargetsOverTime []TargetCount`
  - `Encounter.AllTargetUnits []*Unit` — the full pool; `Encounter.TargetUnits`
    becomes the **active prefix** of it
  - `func (encounter *Encounter) SetActiveTargetCount(sim *Simulation, count int32)`
  - `func (sim *Simulation) initEncounterTargets()`
  Task 5 exercises these end to end.

The pool is a prefix, which is what makes this cheap: targets `0..count-1` are
active and the rest are disabled, so every existing `range
Encounter.TargetUnits` loop (33 of them) becomes correct for free.

- [ ] **Step 1: Write the failing test**

Create `sim/core/encounter_targets_test.go`:

```go
package core

import (
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
)

// The pool is sized to the largest count the timeline asks for, by
// repeating the last configured target. The site sends one target and a
// timeline; it does not send five copies.
func TestTimelineSizesThePoolToItsLargestCount(t *testing.T) {
	enc := NewEncounter(&proto.Encounter{
		Duration: 180,
		Targets:  []*proto.Target{DefaultTargetProtoLvl60},
		TargetsOverTime: []*proto.TargetCountAt{
			{AtSeconds: 0, Count: 1},
			{AtSeconds: 40, Count: 5},
			{AtSeconds: 120, Count: 3},
		},
	})
	if got := len(enc.AllTargetUnits); got != 5 {
		t.Errorf("pool size = %d, want 5", got)
	}
	if got := len(enc.TargetUnits); got != 1 {
		t.Errorf("active count at construction = %d, want the timeline's t=0 count of 1", got)
	}
}

// Entries arrive in whatever order the request had them; the engine sorts.
func TestTimelineIsSortedByTime(t *testing.T) {
	enc := NewEncounter(&proto.Encounter{
		Duration: 180,
		Targets:  []*proto.Target{DefaultTargetProtoLvl60},
		TargetsOverTime: []*proto.TargetCountAt{
			{AtSeconds: 120, Count: 3},
			{AtSeconds: 0, Count: 1},
			{AtSeconds: 40, Count: 5},
		},
	})
	want := []TargetCount{
		{At: 0, Count: 1},
		{At: 40 * time.Second, Count: 5},
		{At: 120 * time.Second, Count: 3},
	}
	if len(enc.TargetsOverTime) != len(want) {
		t.Fatalf("TargetsOverTime = %v, want %v", enc.TargetsOverTime, want)
	}
	for i, tc := range want {
		if enc.TargetsOverTime[i] != tc {
			t.Errorf("TargetsOverTime[%d] = %+v, want %+v", i, enc.TargetsOverTime[i], tc)
		}
	}
}

// An encounter with no timeline is exactly today's encounter.
func TestNoTimelineLeavesEveryTargetActive(t *testing.T) {
	enc := NewEncounter(&proto.Encounter{
		Duration: 180,
		Targets:  []*proto.Target{DefaultTargetProtoLvl60, DefaultTargetProtoLvl60},
	})
	if len(enc.TargetsOverTime) != 0 {
		t.Errorf("TargetsOverTime = %v, want empty", enc.TargetsOverTime)
	}
	if got := len(enc.TargetUnits); got != 2 {
		t.Errorf("active count = %d, want both targets", got)
	}
}

// Activating and deactivating is what the timeline does at run time: the
// active prefix grows and shrinks, GetNumTargets follows it, and the AoE
// cap multiplier is recomputed against the active count, not the pool.
func TestSetActiveTargetCountMovesThePrefix(t *testing.T) {
	sim := timelineSim(t, []*proto.TargetCountAt{{AtSeconds: 0, Count: 1}, {AtSeconds: 40, Count: 3}})

	if got := sim.GetNumTargets(); got != 1 {
		t.Fatalf("GetNumTargets after reset = %d, want 1", got)
	}
	if sim.Encounter.Targets[2].IsEnabled() {
		t.Error("target 3 is enabled while the timeline says one target")
	}

	sim.Encounter.SetActiveTargetCount(sim, 3)
	if got := sim.GetNumTargets(); got != 3 {
		t.Errorf("GetNumTargets after activating = %d, want 3", got)
	}
	if !sim.Encounter.Targets[2].IsEnabled() {
		t.Error("target 3 is disabled after activating three")
	}
	if got := sim.Encounter.AOECapMultiplier(); got != 1 {
		t.Errorf("AOECapMultiplier with 3 active = %v, want 1", got)
	}

	sim.Encounter.SetActiveTargetCount(sim, 1)
	if got := sim.GetNumTargets(); got != 1 {
		t.Errorf("GetNumTargets after deactivating = %d, want 1", got)
	}
	if sim.Encounter.Targets[2].IsEnabled() {
		t.Error("target 3 is enabled after deactivating back to one")
	}
}

// A count below one or above the pool is clamped rather than panicking:
// the request is validated on the site, and the engine is not the place
// to crash a paying run over a bad number.
func TestSetActiveTargetCountClamps(t *testing.T) {
	sim := timelineSim(t, []*proto.TargetCountAt{{AtSeconds: 0, Count: 1}, {AtSeconds: 40, Count: 3}})

	sim.Encounter.SetActiveTargetCount(sim, 0)
	if got := sim.GetNumTargets(); got != 1 {
		t.Errorf("GetNumTargets after asking for 0 = %d, want the clamped 1", got)
	}
	sim.Encounter.SetActiveTargetCount(sim, 99)
	if got := sim.GetNumTargets(); got != 3 {
		t.Errorf("GetNumTargets after asking for 99 = %d, want the clamped pool size 3", got)
	}
}

// A player whose current target is deactivated is retargeted, or its
// rotation spends the rest of the fight attacking a corpse.
func TestDeactivatingRetargetsThePlayer(t *testing.T) {
	sim := timelineSim(t, []*proto.TargetCountAt{{AtSeconds: 0, Count: 3}, {AtSeconds: 40, Count: 1}})
	player := &sim.Raid.Parties[0].Players[0].GetCharacter().Unit
	player.CurrentTarget = sim.Encounter.AllTargetUnits[2]

	sim.Encounter.SetActiveTargetCount(sim, 1)

	if player.CurrentTarget != sim.Encounter.AllTargetUnits[0] {
		t.Errorf("CurrentTarget = %v after its target went away, want the first active target", player.CurrentTarget.Label)
	}
}

// timelineSim builds a one-player simulation with the given timeline,
// reset and ready to step.
func timelineSim(t *testing.T, timeline []*proto.TargetCountAt) *Simulation {
	t.Helper()
	sim := NewSim(&proto.RaidSimRequest{
		Raid: SinglePlayerRaidProto(&proto.Player{
			Name:      "Timeline Test",
			Race:      proto.Race_RaceOrc,
			Class:     proto.Class_ClassShaman,
			Spec:      &proto.Player_ElementalShaman{ElementalShaman: &proto.ElementalShaman{}},
			Equipment: &proto.EquipmentSpec{},
		}, &proto.PartyBuffs{}, &proto.RaidBuffs{}, &proto.Debuffs{}),
		Encounter: &proto.Encounter{
			Duration:        180,
			Targets:         []*proto.Target{DefaultTargetProtoLvl60},
			TargetsOverTime: timeline,
		},
		SimOptions: &proto.SimOptions{Iterations: 1, IsTest: true, RandomSeed: 1},
	}, simsignals.CreateSignals())
	sim.reset()
	return sim
}
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `cd /Users/jh/code/wowsims-forever && go test --tags=with_db ./sim/core/ -run 'Timeline|TargetCount|Deactivating' -v`
Expected: FAIL — compile error, `enc.AllTargetUnits undefined`, `TargetCount undefined`, `SetActiveTargetCount undefined`.

- [ ] **Step 3: Add the timeline to the core `Encounter`**

In `sim/core/target.go`, change the head of the `Encounter` struct to:

```go
type Encounter struct {
	Duration          time.Duration
	DurationVariation time.Duration
	Targets           []*Target

	// TargetUnits is the ACTIVE prefix of AllTargetUnits. Every AoE loop
	// and every GetNumTargets reads it, so shrinking it is how a target
	// stops existing for the rotation. AllTargetUnits is the whole pool
	// and is what the environment indexes and builds attack tables for.
	TargetUnits    []*Unit
	AllTargetUnits []*Unit

	// TargetsOverTime, when set, is the timeline of how many targets are
	// active, sorted by time. Empty means every pooled target is active
	// for the whole fight, which is every encounter that predates the
	// parity work.
	TargetsOverTime []TargetCount
```

(leave the rest of the struct — `Biome` onward — unchanged), and add under it:

```go
// TargetCount is one step of an encounter's target-count timeline.
type TargetCount struct {
	At    time.Duration
	Count int32
}
```

In `NewEncounter`, immediately after the `options.ExecuteProportion_35 = ...`
line at the top, add:

```go
	padTargetsForTimeline(options)
```

and in the `encounter := Encounter{...}` literal add
`TargetsOverTime: newTargetTimeline(options.TargetsOverTime),` after `Biome:`.

Replace the target-construction loop's tail so both slices are filled. Change:

```go
	for targetIndex, targetOptions := range options.Targets {
		target := NewTarget(targetOptions, int32(targetIndex))
		encounter.Targets = append(encounter.Targets, target)
		encounter.TargetUnits = append(encounter.TargetUnits, &target.Unit)
	}
	if len(encounter.Targets) == 0 {
		// Add a dummy target. The only case where targets aren't specified is when
		// computing character stats, and targets won't matter there.
		target := NewTarget(&proto.Target{}, 0)
		encounter.Targets = append(encounter.Targets, target)
		encounter.TargetUnits = append(encounter.TargetUnits, &target.Unit)
	}
```

to:

```go
	for targetIndex, targetOptions := range options.Targets {
		target := NewTarget(targetOptions, int32(targetIndex))
		encounter.Targets = append(encounter.Targets, target)
		encounter.AllTargetUnits = append(encounter.AllTargetUnits, &target.Unit)
	}
	if len(encounter.Targets) == 0 {
		// Add a dummy target. The only case where targets aren't specified is when
		// computing character stats, and targets won't matter there.
		target := NewTarget(&proto.Target{}, 0)
		encounter.Targets = append(encounter.Targets, target)
		encounter.AllTargetUnits = append(encounter.AllTargetUnits, &target.Unit)
	}
	encounter.TargetUnits = encounter.AllTargetUnits[:encounter.initialActiveCount()]
```

and change `updateAOECapMultiplier` to read the active count:

```go
func (encounter *Encounter) updateAOECapMultiplier() {
	encounter.aoeCapMultiplier = min(10/float64(len(encounter.TargetUnits)), 1)
}
```

- [ ] **Step 4: Write the pool and the schedule**

Create `sim/core/encounter_targets.go`:

```go
package core

import (
	"slices"

	"github.com/wowsims/classic/sim/core/proto"
)

// padTargetsForTimeline grows options.Targets to the largest count the
// timeline asks for by repeating the last one. The site sends one target
// and a timeline, not five copies of the target, and growing the proto
// here means the pooled targets are constructed, initialized and given
// attack tables by exactly the same code as a plain five-target fight.
// NewEncounter already mutates options (the execute proportions), so this
// is the file's established shape.
func padTargetsForTimeline(options *proto.Encounter) {
	if len(options.Targets) == 0 || len(options.TargetsOverTime) == 0 {
		return
	}
	maxCount := 0
	for _, entry := range options.TargetsOverTime {
		maxCount = max(maxCount, int(entry.Count))
	}
	for len(options.Targets) < maxCount {
		options.Targets = append(options.Targets, options.Targets[len(options.Targets)-1])
	}
}

// newTargetTimeline converts the proto entries to sim time and sorts them,
// so the request may list them in any order.
func newTargetTimeline(entries []*proto.TargetCountAt) []TargetCount {
	if len(entries) == 0 {
		return nil
	}
	timeline := make([]TargetCount, 0, len(entries))
	for _, entry := range entries {
		timeline = append(timeline, TargetCount{
			At:    DurationFromSeconds(entry.AtSeconds),
			Count: entry.Count,
		})
	}
	slices.SortStableFunc(timeline, func(a, b TargetCount) int {
		return int(a.At - b.At)
	})
	return timeline
}

// initialActiveCount is how many targets are up when the pull starts: the
// last timeline entry at or before zero, or the whole pool when there is
// no timeline.
func (encounter *Encounter) initialActiveCount() int32 {
	count := int32(len(encounter.AllTargetUnits))
	if len(encounter.TargetsOverTime) == 0 {
		return count
	}
	count = encounter.TargetsOverTime[0].Count
	for _, entry := range encounter.TargetsOverTime {
		if entry.At > 0 {
			break
		}
		count = entry.Count
	}
	return count
}

// SetActiveTargetCount makes the first `count` pooled targets active and
// every other one inactive. The count is clamped rather than panicking:
// the request is validated on the site, and a bad number must not crash a
// run that a player paid for.
func (encounter *Encounter) SetActiveTargetCount(sim *Simulation, count int32) {
	count = min(max(count, 1), int32(len(encounter.AllTargetUnits)))

	for i, target := range encounter.Targets {
		if int32(i) < count {
			target.activate(sim)
		} else {
			target.deactivate(sim)
		}
	}

	encounter.TargetUnits = encounter.AllTargetUnits[:count]
	encounter.updateAOECapMultiplier()
}

// activate brings a pooled target into the fight. It is a no-op on a
// target that is already in, which is what makes SetActiveTargetCount
// safe to call every step of the timeline.
func (target *Target) activate(sim *Simulation) {
	if target.IsEnabled() {
		return
	}
	target.enabled = true
	target.SetGCDTimer(sim, max(0, sim.CurrentTime))
	target.AutoAttacks.EnableAutoSwing(sim)
	if sim.Log != nil {
		target.Log(sim, "Target activated")
	}
}

// deactivate takes a pooled target out. Its auras expire, so the player's
// DoTs on it stop ticking, and anyone still pointed at it is retargeted
// at the first active target.
func (target *Target) deactivate(sim *Simulation) {
	if !target.IsEnabled() {
		return
	}
	target.enabled = false
	target.auraTracker.expireAll(sim)
	target.AutoAttacks.CancelAutoSwing(sim)
	target.CancelGCDTimer(sim)
	target.Hardcast = Hardcast{Expires: startingCDTime}

	first := target.Env.Encounter.AllTargetUnits[0]
	for _, unit := range target.Env.Raid.AllUnits {
		if unit.CurrentTarget == &target.Unit {
			unit.CurrentTarget = first
		}
	}

	if sim.Log != nil {
		target.Log(sim, "Target deactivated")
	}
}

// initEncounterTargets applies the timeline for one iteration.
// Simulation.reset calls it after Environment.reset has re-enabled every
// target, so the first thing it does is take the later ones back out.
func (sim *Simulation) initEncounterTargets() {
	encounter := &sim.Encounter
	encounter.SetActiveTargetCount(sim, encounter.initialActiveCount())

	for _, entry := range encounter.TargetsOverTime {
		if entry.At <= 0 {
			continue
		}
		count := entry.Count
		StartDelayedAction(sim, DelayedActionOptions{
			DoAt:     entry.At,
			Priority: ActionPriorityDOT,
			OnAction: func(sim *Simulation) {
				encounter.SetActiveTargetCount(sim, count)
			},
		})
	}
}
```

- [ ] **Step 5: Point the environment at the pool**

In `sim/core/environment.go`, change line 78 from

```go
	env.AllUnits = append(env.Encounter.TargetUnits, env.Raid.AllUnits...)
```

to

```go
	env.AllUnits = append(env.Encounter.AllTargetUnits, env.Raid.AllUnits...)
```

and change line 86 from

```go
		unit.CurrentTarget = env.Encounter.TargetUnits[0]
```

to

```go
		unit.CurrentTarget = env.Encounter.AllTargetUnits[0]
```

and, in the debuff block at line 90, change both references so every pooled
target gets the raid's debuffs:

```go
	if raidProto.Debuffs != nil && len(env.Encounter.AllTargetUnits) > 0 {
		for targetIdx, targetUnit := range env.Encounter.AllTargetUnits {
```

`GetNumTargets` already reads `len(env.Encounter.Targets)`; change it to the
active slice:

```go
func (env *Environment) GetNumTargets() int32 {
	return int32(len(env.Encounter.TargetUnits))
}
```

In `sim/core/sim.go`, extend the tail of `reset()` again:

```go
	sim.initManaTickAction()
	sim.initEncounterMovement()
	sim.initEncounterTargets()
```

- [ ] **Step 6: Run the test to verify it passes**

Run: `cd /Users/jh/code/wowsims-forever && go test --tags=with_db ./sim/core/ -run 'Timeline|TargetCount|Deactivating' -v`
Expected: PASS, all six tests.

- [ ] **Step 7: Run the whole suite so no golden moved**

Run: `go test --tags=with_db -count=1 ./sim/...`
Expected: PASS. With no timeline, `initialActiveCount` returns the whole pool
and `TargetUnits` is the whole slice, so every golden must be unchanged.

- [ ] **Step 8: gofmt and commit**

```bash
cd /Users/jh/code/wowsims-forever
gofmt -l sim/core/encounter_targets.go sim/core/encounter_targets_test.go sim/core/target.go sim/core/environment.go sim/core/sim.go
git add sim/core/encounter_targets.go sim/core/encounter_targets_test.go sim/core/target.go sim/core/environment.go sim/core/sim.go
git commit -m "$(cat <<'MSG'
feat(core): targets activate and deactivate on a timeline

The pool is sized to the timeline's largest count by repeating the last
configured target, and activity is the prefix of it, so all 33 existing
AoE loops over Encounter.TargetUnits become count-aware for free. A
deactivated target's auras expire and anyone aimed at it is retargeted.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
MSG
)"
```

---

### Task 5: The target timeline, end to end, on a fury warrior and a frost mage

**Files:**
- Create: `sim/encounter_targets_e2e_test.go` (package `sim`)

**Interfaces:**
- Consumes: `furyWarriorPlayer()`, `frostMagePlayer()`, `parityEncounter()`,
  `runParitySim()`, `parityIterations` from Task 3's file, same package.
- Produces: nothing later tasks use.

- [ ] **Step 1: Write the failing test**

Create `sim/encounter_targets_e2e_test.go`:

```go
package sim

import (
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
)

// The dungeon-pull timeline from the contract's section 1.6: one boss,
// then packs of three and five, then back down. Total damage must land
// between a pure single-target fight and a permanent five-target fight,
// because the adds are up for part of the time and not all of it.
func TestDungeonTimelineLandsBetweenOneAndFiveTargets(t *testing.T) {
	dungeon := []*proto.TargetCountAt{
		{AtSeconds: 0, Count: 1},
		{AtSeconds: 40, Count: 3},
		{AtSeconds: 80, Count: 5},
		{AtSeconds: 130, Count: 3},
		{AtSeconds: 160, Count: 1},
	}

	for _, tc := range []struct {
		name   string
		player func() *proto.Player
	}{
		{"fury warrior", furyWarriorPlayer},
		{"frost mage", frostMagePlayer},
	} {
		t.Run(tc.name, func(t *testing.T) {
			single := runParitySim(t, tc.player(), parityEncounter())

			five := parityEncounter()
			five.Targets = []*proto.Target{
				core.DefaultTargetProtoLvl60, core.DefaultTargetProtoLvl60, core.DefaultTargetProtoLvl60,
				core.DefaultTargetProtoLvl60, core.DefaultTargetProtoLvl60,
			}
			fiveDps := runParitySim(t, tc.player(), five)

			timeline := parityEncounter()
			timeline.TargetsOverTime = dungeon
			timelineDps := runParitySim(t, tc.player(), timeline)

			if timelineDps < single {
				t.Errorf("dungeon-timeline DPS %.1f is below the single-target DPS %.1f", timelineDps, single)
			}
			if timelineDps > fiveDps {
				t.Errorf("dungeon-timeline DPS %.1f is above the permanent five-target DPS %.1f", timelineDps, fiveDps)
			}
		})
	}
}

// The timeline's target count is what the rotation sees: an APL asking
// "how many targets" during the one-target opening must be told one, not
// the pool's five.
func TestNumTargetsFollowsTheTimeline(t *testing.T) {
	request := &proto.RaidSimRequest{
		Raid: core.SinglePlayerRaidProto(furyWarriorPlayer(), &proto.PartyBuffs{}, &proto.RaidBuffs{}, &proto.Debuffs{}),
		Encounter: &proto.Encounter{
			Duration: 180,
			Targets:  []*proto.Target{core.DefaultTargetProtoLvl60},
			TargetsOverTime: []*proto.TargetCountAt{
				{AtSeconds: 0, Count: 1},
				{AtSeconds: 90, Count: 5},
			},
		},
		SimOptions: &proto.SimOptions{Iterations: 1, IsTest: true, RandomSeed: 1},
	}

	sim := core.NewSim(request, simsignals.CreateSignals())
	sim.Reset()
	if got := sim.GetNumTargets(); got != 1 {
		t.Fatalf("GetNumTargets at the pull = %d, want 1", got)
	}

	result := core.RunSim(request, nil, simsignals.CreateSignals())
	if result.Error != nil {
		t.Fatalf("sim failed: %s", result.Error.Message)
	}
	if got := len(result.EncounterMetrics.Targets); got != 5 {
		t.Errorf("EncounterMetrics has %d targets, want the whole pool of 5 so the report can name the adds", got)
	}
}
```

- [ ] **Step 2: Run the test to verify it fails**

Temporarily comment out the `sim.initEncounterTargets()` line added in Task 4,
run: `cd /Users/jh/code/wowsims-forever && go test --tags=with_db ./sim/ -run 'Timeline|NumTargets' -v`
Expected: FAIL — `GetNumTargets at the pull = 5, want 1`. Restore the line.

- [ ] **Step 3: Run the test to verify it passes**

Run: `go test --tags=with_db ./sim/ -run 'Timeline|NumTargets' -v`
Expected: PASS, three subtests plus the second test.

- [ ] **Step 4: Commit**

```bash
cd /Users/jh/code/wowsims-forever
gofmt -l sim/encounter_targets_e2e_test.go
git add sim/encounter_targets_e2e_test.go
git commit -m "$(cat <<'MSG'
test(core): the dungeon timeline, end to end, on fury and frost

The contract's dungeon-pull timeline must land between a single-target
fight and a permanent five-target one for both a melee and a caster, and
the rotation's "how many targets" must follow the timeline rather than
the pool.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
MSG
)"
```

---

### Task 6: Target-dummy mode

**Files:**
- Modify: `sim/core/stats/stats.go:495-503` (the `PseudoStats` struct)
- Modify: `sim/core/unit.go:392-394` (`Unit.Armor`)
- Modify: `sim/core/target.go` (the `Encounter` struct and `NewEncounter`)
- Modify: `sim/core/environment.go:88-93` (the debuff block)
- Test: `sim/core/encounter_dummy_test.go` (create)

**Interfaces:**
- Consumes: `proto.Encounter.TargetDummy` from Task 1.
- Produces: `Encounter.Dummy bool`, `stats.PseudoStats.ArmorReductionDisabled bool`.
  Task 7 exercises them end to end.

The three clauses of the contract's sentence become three switches:

1. *no debuff application* — the raid debuff panel (which models other raiders)
   is not applied. A debuff the player's own rotation casts still lands, because
   a real dummy parse has the player's own Sunder and Curse on it.
2. *no execute window* — all three execute proportions are zeroed.
3. *no armor reduction* — the target's armor is read from its initial stats, so
   nothing lowers it.

- [ ] **Step 1: Write the failing test**

Create `sim/core/encounter_dummy_test.go`:

```go
package core

import (
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

// Dummy mode zeroes every execute proportion, whatever the request asked
// for: a dummy never drops below full health.
func TestDummyModeHasNoExecuteWindow(t *testing.T) {
	enc := NewEncounter(&proto.Encounter{
		Duration:              180,
		Targets:               []*proto.Target{DefaultTargetProtoLvl60},
		ExecuteProportion_20:  0.25,
		ExecuteProportion_25:  0.25,
		ExecuteProportion_35:  0.35,
		TargetDummy:           true,
	})
	if enc.Dummy != true {
		t.Error("Encounter.Dummy = false, want true")
	}
	if enc.ExecuteProportion_20 != 0 || enc.ExecuteProportion_25 != 0 || enc.ExecuteProportion_35 != 0 {
		t.Errorf("execute proportions = %v/%v/%v, want all zero on a dummy",
			enc.ExecuteProportion_20, enc.ExecuteProportion_25, enc.ExecuteProportion_35)
	}
}

// Without the flag nothing changes: the feature is additive.
func TestNonDummyKeepsItsExecuteWindow(t *testing.T) {
	enc := NewEncounter(&proto.Encounter{
		Duration:             180,
		Targets:              []*proto.Target{DefaultTargetProtoLvl60},
		ExecuteProportion_20: 0.25,
	})
	if enc.Dummy {
		t.Error("Encounter.Dummy = true, want false")
	}
	if enc.ExecuteProportion_20 != 0.25 {
		t.Errorf("ExecuteProportion_20 = %v, want 0.25", enc.ExecuteProportion_20)
	}
}

// Armor on a dummy is pinned to what the request configured: Sunder,
// Expose and Faerie Fire all lower stats.Armor dynamically, and on a
// dummy none of them may move the number the damage formula reads.
func TestDummyArmorIgnoresReduction(t *testing.T) {
	unit := &Unit{
		Type:         EnemyUnit,
		PseudoStats:  stats.NewPseudoStats(),
		initialStats: stats.Stats{stats.Armor: 3000},
	}
	unit.stats = unit.initialStats

	unit.stats[stats.Armor] = 1500
	if got := unit.Armor(); got != 1500 {
		t.Errorf("Armor() on a normal target = %v, want the reduced 1500", got)
	}

	unit.PseudoStats.ArmorReductionDisabled = true
	if got := unit.Armor(); got != 3000 {
		t.Errorf("Armor() on a dummy = %v, want the unreduced 3000", got)
	}
}

// The multiplier still applies on a dummy: it is gear and talents on the
// attacker's side of the table, not a debuff on the target.
func TestDummyArmorStillHonoursTheMultiplier(t *testing.T) {
	unit := &Unit{
		Type:         EnemyUnit,
		PseudoStats:  stats.NewPseudoStats(),
		initialStats: stats.Stats{stats.Armor: 3000},
	}
	unit.stats = unit.initialStats
	unit.PseudoStats.ArmorReductionDisabled = true
	unit.PseudoStats.ArmorMultiplier = 0.5

	if got := unit.Armor(); got != 1500 {
		t.Errorf("Armor() = %v, want 1500", got)
	}
}

// The raid debuff panel models other raiders, and a dummy has none, so
// none of it is applied and the target's armor is untouched by it.
func TestDummyModeSkipsTheRaidDebuffPanel(t *testing.T) {
	withDebuffs := dummyEnv(t, false, &proto.Debuffs{SunderArmor: true})
	onDummy := dummyEnv(t, true, &proto.Debuffs{SunderArmor: true})

	if !withDebuffs.Encounter.Targets[0].HasAura("Sunder Armor") {
		t.Fatal("a normal target has no Sunder Armor aura registered; the fixture is wrong, not the feature")
	}
	if onDummy.Encounter.Targets[0].HasAura("Sunder Armor") {
		t.Error("a dummy target has a Sunder Armor aura from the raid debuff panel, want none")
	}
}

func dummyEnv(t *testing.T, dummy bool, debuffs *proto.Debuffs) *Environment {
	t.Helper()
	env, _, _ := NewEnvironment(
		SinglePlayerRaidProto(&proto.Player{
			Name:      "Dummy Test",
			Race:      proto.Race_RaceOrc,
			Class:     proto.Class_ClassShaman,
			Spec:      &proto.Player_ElementalShaman{ElementalShaman: &proto.ElementalShaman{}},
			Equipment: &proto.EquipmentSpec{},
		}, &proto.PartyBuffs{}, &proto.RaidBuffs{}, debuffs),
		&proto.Encounter{
			Duration:    180,
			Targets:     []*proto.Target{DefaultTargetProtoLvl60},
			TargetDummy: dummy,
		},
		false,
	)
	return env
}
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `cd /Users/jh/code/wowsims-forever && go test --tags=with_db ./sim/core/ -run 'Dummy' -v`
Expected: FAIL — compile error, `enc.Dummy undefined`, `ArmorReductionDisabled undefined`.

- [ ] **Step 3: Add the pseudo-stat and the armor gate**

In `sim/core/stats/stats.go`, inside the `PseudoStats` struct, next to
`ArmorMultiplier`, add:

```go
	// ArmorReductionDisabled pins a unit's armor to what it started the
	// iteration with. Target-dummy encounters set it: a dummy takes no
	// Sunder, no Expose and no Faerie Fire, so nothing lowers the number
	// the damage formula reads.
	ArmorReductionDisabled bool
```

In `sim/core/unit.go`, change `Unit.Armor`:

```go
func (unit *Unit) Armor() float64 {
	armor := unit.stats[stats.Armor]
	if unit.PseudoStats.ArmorReductionDisabled {
		armor = unit.initialStats[stats.Armor]
	}
	return max(unit.PseudoStats.ArmorMultiplier*armor, 0.0)
}
```

- [ ] **Step 4: Add the flag to the encounter**

In `sim/core/target.go`, add to the `Encounter` struct, after the
`TargetsOverTime` block:

```go
	// Dummy is target-dummy mode: the raid debuff panel is not applied,
	// there is no execute window, and nothing reduces the target's armor.
	// The player's own rotation still debuffs the target, because a real
	// dummy parse has the player's own Sunder and Curse on it.
	Dummy bool
```

In `NewEncounter`, before the execute-proportion `max(...)` lines at the very
top of the function, add:

```go
	if options.TargetDummy {
		options.ExecuteProportion_20 = 0
		options.ExecuteProportion_25 = 0
		options.ExecuteProportion_35 = 0
	}
```

and add `Dummy: options.TargetDummy,` to the `encounter := Encounter{...}`
literal, after `Biome:`.

At the end of the target-construction loop in `NewEncounter`, after
`encounter.TargetUnits = encounter.AllTargetUnits[:encounter.initialActiveCount()]`,
add:

```go
	if encounter.Dummy {
		for _, target := range encounter.Targets {
			target.PseudoStats.ArmorReductionDisabled = true
		}
	}
```

- [ ] **Step 5: Skip the raid debuff panel on a dummy**

In `sim/core/environment.go`, change the debuff block's condition:

```go
	// Apply extra debuffs from raid. A target dummy gets none: the panel
	// models the other twenty-nine raiders, and a dummy stands alone.
	if raidProto.Debuffs != nil && !env.Encounter.Dummy && len(env.Encounter.AllTargetUnits) > 0 {
```

- [ ] **Step 6: Run the test to verify it passes**

Run: `cd /Users/jh/code/wowsims-forever && go test --tags=with_db ./sim/core/ -run 'Dummy' -v`
Expected: PASS, all five tests.

- [ ] **Step 7: Run the whole suite so no golden moved**

Run: `go test --tags=with_db -count=1 ./sim/...`
Expected: PASS.

- [ ] **Step 8: gofmt and commit**

```bash
cd /Users/jh/code/wowsims-forever
gofmt -l sim/core/stats/stats.go sim/core/unit.go sim/core/target.go sim/core/environment.go sim/core/encounter_dummy_test.go
git add sim/core/stats/stats.go sim/core/unit.go sim/core/target.go sim/core/environment.go sim/core/encounter_dummy_test.go
git commit -m "$(cat <<'MSG'
feat(core): target-dummy mode

Three switches for the contract's three clauses: the raid debuff panel is
not applied, every execute proportion is zeroed, and the target's armor is
read from its initial stats so nothing lowers it. The player's own
rotation still debuffs the target, which is what a real dummy parse shows.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
MSG
)"
```

---

### Task 7: Target-dummy mode, end to end, on a fury warrior and a frost mage

**Files:**
- Create: `sim/encounter_dummy_e2e_test.go` (package `sim`)

**Interfaces:**
- Consumes: `furyWarriorPlayer()`, `frostMagePlayer()`, `parityEncounter()`,
  `runParitySim()` from Task 3's file, same package.
- Produces: nothing later tasks use.

- [ ] **Step 1: Write the failing test**

Create `sim/encounter_dummy_e2e_test.go`:

```go
package sim

import (
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
)

// fullDebuffs is the raid debuff panel a normal fight carries and a dummy
// does not.
func fullDebuffs() *proto.Debuffs {
	return &proto.Debuffs{
		SunderArmor:     true,
		FaerieFire:      true,
		CurseOfElements: true,
	}
}

func runDummyParitySim(t *testing.T, player *proto.Player, encounter *proto.Encounter) float64 {
	t.Helper()
	result := core.RunSim(&proto.RaidSimRequest{
		Raid:      core.SinglePlayerRaidProto(player, &proto.PartyBuffs{}, &proto.RaidBuffs{}, fullDebuffs()),
		Encounter: encounter,
		SimOptions: &proto.SimOptions{
			Iterations: parityIterations,
			IsTest:     true,
			RandomSeed: 1,
		},
	}, nil, simsignals.CreateSignals())
	if result.Error != nil {
		t.Fatalf("sim failed: %s", result.Error.Message)
	}
	return result.RaidMetrics.Dps.Avg
}

// A dummy has none of the raid's debuffs and no execute window, so both
// reference specs do less damage on one than on a real boss.
func TestDummyModeIsWorseThanARealBossForBothSpecs(t *testing.T) {
	for _, tc := range []struct {
		name   string
		player func() *proto.Player
	}{
		{"fury warrior", furyWarriorPlayer},
		{"frost mage", frostMagePlayer},
	} {
		t.Run(tc.name, func(t *testing.T) {
			boss := parityEncounter()
			boss.ExecuteProportion_20 = 0.25
			bossDps := runDummyParitySim(t, tc.player(), boss)

			dummy := parityEncounter()
			dummy.ExecuteProportion_20 = 0.25
			dummy.TargetDummy = true
			dummyDps := runDummyParitySim(t, tc.player(), dummy)

			if dummyDps >= bossDps {
				t.Errorf("dummy DPS %.1f is not below the boss DPS %.1f", dummyDps, bossDps)
			}
		})
	}
}

// The execute half of the claim, isolated: a fury warrior's Execute is
// the largest single thing the window buys, so on a dummy the spell must
// never be cast at all.
func TestDummyModeNeverLetsTheWarriorExecute(t *testing.T) {
	dummy := parityEncounter()
	dummy.ExecuteProportion_20 = 1.0
	dummy.TargetDummy = true

	result := core.RunSim(&proto.RaidSimRequest{
		Raid:      core.SinglePlayerRaidProto(furyWarriorPlayer(), &proto.PartyBuffs{}, &proto.RaidBuffs{}, fullDebuffs()),
		Encounter: dummy,
		SimOptions: &proto.SimOptions{Iterations: parityIterations, IsTest: true, RandomSeed: 1},
	}, nil, simsignals.CreateSignals())
	if result.Error != nil {
		t.Fatalf("sim failed: %s", result.Error.Message)
	}

	for _, action := range result.RaidMetrics.Parties[0].Players[0].Actions {
		if action.Id.GetSpellId() == 0 {
			continue
		}
		casts := int32(0)
		for _, target := range action.Targets {
			casts += target.Casts
		}
		if casts > 0 && isWarriorExecute(action.Id.GetSpellId()) {
			t.Errorf("Execute (spell %d) was cast %d times on a target dummy", action.Id.GetSpellId(), casts)
		}
	}
}

// isWarriorExecute names the Execute ranks the client ships, so the test
// does not depend on which rank the reference gear affords.
func isWarriorExecute(spellID int32) bool {
	switch spellID {
	case 5308, 20658, 20660, 20661, 20662:
		return true
	}
	return false
}
```

- [ ] **Step 2: Run the test to verify it fails**

Temporarily change `NewEncounter`'s dummy block to not zero the execute
proportions, run: `cd /Users/jh/code/wowsims-forever && go test --tags=with_db ./sim/ -run 'Dummy' -v`
Expected: FAIL — `Execute (spell 20662) was cast N times on a target dummy`.
Restore the block.

- [ ] **Step 3: Run the test to verify it passes**

Run: `go test --tags=with_db ./sim/ -run 'Dummy' -v`
Expected: PASS, two subtests plus the Execute test.

- [ ] **Step 4: Commit**

```bash
cd /Users/jh/code/wowsims-forever
gofmt -l sim/encounter_dummy_e2e_test.go
git add sim/encounter_dummy_e2e_test.go
git commit -m "$(cat <<'MSG'
test(core): target-dummy mode, end to end, on fury and frost

A dummy is worse than a real boss for both reference specs, and a warrior
never casts Execute on one however wide the request's execute window was.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
MSG
)"
```

---

### Task 8: The `sample_iteration` proto and its opt-in flag

**Files:**
- Modify: `proto/api.proto:120-129` (`SimOptions`) and `:374-388` (`RaidSimResult`)
- Regenerate: `sim/core/proto/api.pb.go`, `ui/core/proto/api.ts`
- Test: `sim/core/sample_iteration_fields_test.go` (create)

**Interfaces:**
- Consumes: nothing.
- Produces:
  - `proto.SampleCast{AtMs int64; ActionId *ActionID; Target string; Resources map[string]int32}`
  - `proto.SampleIteration{Dps float64; DurationSeconds float64; Casts []*SampleCast}`
  - `proto.RaidSimResult.SampleIteration *SampleIteration` (field 8)
  - `proto.SimOptions.SampleIteration bool` (field 10)
  Task 9 fills these.

`SimOptions` uses 1, 2, 3, 5, 6, 7, 8, 9 — take 10. `RaidSimResult` uses 1–7 —
take 8. The cast carries the whole `ActionID`, not a bare spell id, because the
engine has no display names: `Spell` has no name field at all, and the site
already resolves `ActionID` to a name for the timeline. The sim module's
`SampleCast.Name` is filled there.

- [ ] **Step 1: Write the failing test**

Create `sim/core/sample_iteration_fields_test.go`:

```go
package core

import (
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
)

func TestSampleIterationFieldsExist(t *testing.T) {
	result := &proto.RaidSimResult{
		SampleIteration: &proto.SampleIteration{
			Dps:             1204.5,
			DurationSeconds: 180,
			Casts: []*proto.SampleCast{
				{
					AtMs:      -1500,
					ActionId:  &proto.ActionID{RawId: &proto.ActionID_SpellId{SpellId: 11565}},
					Target:    "Target 1",
					Resources: map[string]int32{"rage": 20},
				},
			},
		},
	}

	sample := result.GetSampleIteration()
	if sample.GetDps() != 1204.5 {
		t.Errorf("Dps = %v, want 1204.5", sample.GetDps())
	}
	if len(sample.GetCasts()) != 1 {
		t.Fatalf("Casts = %v, want one", sample.GetCasts())
	}
	cast := sample.GetCasts()[0]
	if cast.GetAtMs() != -1500 {
		t.Errorf("AtMs = %d, want -1500 so a pre-pull cast reads as negative", cast.GetAtMs())
	}
	if cast.GetActionId().GetSpellId() != 11565 {
		t.Errorf("ActionId.SpellId = %d, want 11565", cast.GetActionId().GetSpellId())
	}
	if cast.GetResources()["rage"] != 20 {
		t.Errorf("Resources[rage] = %d, want 20", cast.GetResources()["rage"])
	}
}

func TestSimOptionsHasTheSampleFlag(t *testing.T) {
	if !(&proto.SimOptions{SampleIteration: true}).GetSampleIteration() {
		t.Error("SimOptions.SampleIteration did not round-trip")
	}
}
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `cd /Users/jh/code/wowsims-forever && go test --tags=with_db ./sim/core/ -run 'Sample' -v`
Expected: FAIL — compile error, `unknown field SampleIteration`.

- [ ] **Step 3: Add the proto**

In `proto/api.proto`, add to `message SimOptions`, after `use_labeled_rands = 9`:

```proto
	// Produce RaidSimResult.sample_iteration: one extra iteration, re-run
	// at the median iteration's seed with a cast recorder attached.
	// Off by default because it costs an iteration and an environment.
	// Forever addition; see PORTING.md.
	bool sample_iteration = 10;
```

Add to `message RaidSimResult`, after `iterations_done = 7`:

```proto
	// The median-DPS iteration's cast log, when sim_options.sample_iteration
	// asked for it. Forever addition; see PORTING.md.
	SampleIteration sample_iteration = 8;
```

Directly above `message RaidSimResult`, add:

```proto
// SampleCast is one cast in the sampled iteration, with the caster's
// resources as they stood immediately after it.
message SampleCast {
	// Milliseconds from the pull; negative during the pre-pull.
	int64 at_ms = 1;
	ActionID action_id = 2;
	string target = 3;
	// Keyed by resource name: rage, energy, mana, combo_points. Only the
	// bars the caster actually has are present.
	map<string, int32> resources = 4;
}

// SampleIteration is one iteration's cast log. The engine picks the
// median-DPS iteration and replays it, so the log is representative
// rather than lucky.
message SampleIteration {
	double dps = 1;
	double duration_seconds = 2;
	repeated SampleCast casts = 3;
}
```

- [ ] **Step 4: Regenerate and verify**

Run:
```bash
cd /Users/jh/code/wowsims-forever
make proto
go test --tags=with_db ./sim/core/proto/ -run TestGeneratedProtosMatchSources -v
go test --tags=with_db ./sim/core/ -run 'Sample' -v
```
Expected: both PASS.

- [ ] **Step 5: gofmt and commit**

```bash
cd /Users/jh/code/wowsims-forever
gofmt -l sim/core/sample_iteration_fields_test.go
git add proto/api.proto sim/core/proto ui/core/proto sim/core/sample_iteration_fields_test.go
git commit -m "$(cat <<'MSG'
feat(core): RaidSimResult.sample_iteration and its opt-in flag

A cast carries the whole ActionID rather than a bare spell id: this
engine has no display names at all - Spell has no name field - and the
site already resolves ActionIDs for the timeline, so the name belongs
there and the id belongs here.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
MSG
)"
```

---

### Task 9: Recording and replaying the median iteration

**Files:**
- Modify: `sim/core/sim.go` (the `Simulation` struct, `reseedRands`, `run`, `runSim`)
- Modify: `sim/core/spell.go:593-598` (`applyEffects`)
- Modify: `sim/core/metrics_aggregator.go:15-55` (`DistributionMetrics`)
- Modify: `sim/core/sim_concurrent.go:330-352` (`CombineConcurrentSimResults`)
- Create: `sim/core/sample_iteration.go`
- Test: `sim/core/sample_iteration_test.go` (create)

**Interfaces:**
- Consumes: `proto.SimOptions.SampleIteration`, `proto.SampleIteration`,
  `proto.SampleCast` from Task 8.
- Produces:
  - `DistributionMetrics.LastIterationValue float64`
  - `func (sim *Simulation) iterationSeed(i int32) int64`
  - `func (sim *Simulation) reseedTo(rseed int64)`
  - `func (sim *Simulation) recordSampleCast(spell *Spell, target *Unit)`
  - `func runSampleIteration(rsr *proto.RaidSimRequest, mainSim *Simulation, signals simsignals.Signals) *proto.SampleIteration`
  Task 10 exercises them end to end.

The shape is a two-pass replay, and the reason matters: the median cannot be
known until every iteration has run, and buffering every iteration's casts
would cost megabytes. Because `reseedRands` derives each iteration's seed from
its index, one iteration can be reproduced exactly from its seed alone. So the
main run records `(dps, seed)` per iteration and nothing else; the replay is one
extra iteration on a **fresh** `Simulation`, which is what keeps the aggregated
metrics free of the extra sample.

- [ ] **Step 1: Write the failing test**

Create `sim/core/sample_iteration_test.go`:

```go
package core

import (
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
)

func sampleRequest(iterations int32, sample bool) *proto.RaidSimRequest {
	return &proto.RaidSimRequest{
		Raid: SinglePlayerRaidProto(&proto.Player{
			Name:      "Sample Test",
			Race:      proto.Race_RaceOrc,
			Class:     proto.Class_ClassShaman,
			Spec:      &proto.Player_ElementalShaman{ElementalShaman: &proto.ElementalShaman{}},
			Equipment: &proto.EquipmentSpec{},
		}, &proto.PartyBuffs{}, &proto.RaidBuffs{}, &proto.Debuffs{}),
		Encounter: &proto.Encounter{
			Duration: 60,
			Targets:  []*proto.Target{DefaultTargetProtoLvl60},
		},
		SimOptions: &proto.SimOptions{
			Iterations:      iterations,
			IsTest:          true,
			RandomSeed:      1,
			SampleIteration: sample,
		},
	}
}

// Off by default: a run that does not ask for a sample must not pay for
// one, and must not carry one.
func TestNoSampleWhenNotAsked(t *testing.T) {
	result := RunSim(sampleRequest(20, false), nil, simsignals.CreateSignals())
	if result.Error != nil {
		t.Fatalf("sim failed: %s", result.Error.Message)
	}
	if result.SampleIteration != nil {
		t.Errorf("SampleIteration = %+v, want nil", result.SampleIteration)
	}
}

// The sample exists, has casts, and its casts are in time order.
func TestSampleIterationIsAnOrderedCastLog(t *testing.T) {
	result := RunSim(sampleRequest(20, true), nil, simsignals.CreateSignals())
	if result.Error != nil {
		t.Fatalf("sim failed: %s", result.Error.Message)
	}
	sample := result.SampleIteration
	if sample == nil {
		t.Fatal("SampleIteration = nil, want a log")
	}
	if len(sample.Casts) == 0 {
		t.Fatal("SampleIteration has no casts")
	}
	for i := 1; i < len(sample.Casts); i++ {
		if sample.Casts[i].AtMs < sample.Casts[i-1].AtMs {
			t.Fatalf("cast %d at %dms precedes cast %d at %dms", i, sample.Casts[i].AtMs, i-1, sample.Casts[i-1].AtMs)
		}
	}
	if sample.DurationSeconds <= 0 {
		t.Errorf("DurationSeconds = %v, want the sampled iteration's length", sample.DurationSeconds)
	}
}

// The sampled iteration is the median one, not the first and not the
// luckiest: its DPS must sit inside the run's min-max band.
func TestSampledIterationIsNearTheMiddle(t *testing.T) {
	result := RunSim(sampleRequest(51, true), nil, simsignals.CreateSignals())
	if result.Error != nil {
		t.Fatalf("sim failed: %s", result.Error.Message)
	}
	dps := result.RaidMetrics.Dps
	got := result.SampleIteration.Dps
	if got < dps.Min || got > dps.Max {
		t.Errorf("sample DPS %.1f is outside the run's [%.1f, %.1f]", got, dps.Min, dps.Max)
	}
}

// The extra iteration must not be counted: a 20-iteration run reports 20.
func TestSampleDoesNotPolluteTheAggregates(t *testing.T) {
	with := RunSim(sampleRequest(20, true), nil, simsignals.CreateSignals())
	without := RunSim(sampleRequest(20, false), nil, simsignals.CreateSignals())
	if with.Error != nil || without.Error != nil {
		t.Fatal("sim failed")
	}
	if with.IterationsDone != 20 {
		t.Errorf("IterationsDone = %d with sampling on, want 20", with.IterationsDone)
	}
	if with.RaidMetrics.Dps.Avg != without.RaidMetrics.Dps.Avg {
		t.Errorf("mean DPS with sampling %.6f differs from without %.6f; the replay leaked into the aggregates",
			with.RaidMetrics.Dps.Avg, without.RaidMetrics.Dps.Avg)
	}
}

// An iteration is reproducible from its seed alone, which is the whole
// premise of the replay.
func TestIterationSeedIsDerivedFromTheIndex(t *testing.T) {
	sim := NewSim(sampleRequest(10, false), simsignals.CreateSignals())
	if got := sim.iterationSeed(3); got != sim.Options.RandomSeed+3 {
		t.Errorf("iterationSeed(3) = %d, want %d", got, sim.Options.RandomSeed+3)
	}
	if got := sim.iterationSeed(0); got != sim.rseed {
		t.Errorf("iterationSeed(0) = %d, want the constructed seed %d", got, sim.rseed)
	}
}
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `cd /Users/jh/code/wowsims-forever && go test --tags=with_db ./sim/core/ -run 'Sample|IterationSeed' -v`
Expected: FAIL — compile error, `sim.iterationSeed undefined`, and
`TestSampleIterationIsAnOrderedCastLog` fails with `SampleIteration = nil`.

- [ ] **Step 3: Record each iteration's DPS**

In `sim/core/metrics_aggregator.go`, add a field to `DistributionMetrics`:

```go
type DistributionMetrics struct {
	// Values for the current iteration. These are cleared after each iteration.
	Total float64

	// LastIterationValue is the value doneIteration last computed. The
	// sample-iteration replay reads it to learn which iteration was the
	// median one, without buffering anything per iteration.
	LastIterationValue float64
```

and set it as the first statement of `doneIteration`:

```go
func (distMetrics *DistributionMetrics) doneIteration(sim *Simulation) {
	dps := distMetrics.Total / sim.Duration.Seconds()
	distMetrics.LastIterationValue = dps
	distMetrics.add(dps)
```

- [ ] **Step 4: Make the seed addressable**

In `sim/core/sim.go`, replace `reseedRands` with a pair:

```go
func (sim *Simulation) reseedRands(i int64) {
	sim.reseedTo(sim.Options.RandomSeed + i)
}

// reseedTo puts the RNGs at an exact seed. reseedRands derives one from
// an iteration index; the sample replay uses the seed directly, which is
// what lets it reproduce a chosen iteration on a fresh Simulation.
func (sim *Simulation) reseedTo(rseed int64) {
	sim.rand.Seed(rseed)

	if sim.isTest {
		for label, rng := range sim.testRands {
			rng.Seed(makeTestRandSeed(rseed, label))
		}
	}
}

// iterationSeed is the seed iteration i ran at. Iteration 0 runs at the
// seed the Simulation was constructed with - which is time-based when the
// request left RandomSeed at zero - and every later iteration at
// RandomSeed+i, because that is what the run loop does.
func (sim *Simulation) iterationSeed(i int32) int64 {
	if i == 0 {
		return sim.rseed
	}
	return sim.Options.RandomSeed + int64(i)
}
```

- [ ] **Step 5: Write the recorder and the replay**

Create `sim/core/sample_iteration.go`:

```go
package core

import (
	"slices"
	"time"

	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
)

// sampleRecorder collects one unit's casts for one iteration.
type sampleRecorder struct {
	unit  *Unit
	casts []*proto.SampleCast
}

// iterationSample is one (dps, seed) pair from the main run. Sixteen
// bytes an iteration is the whole memory cost of knowing which iteration
// was the median one.
type iterationSample struct {
	dps  float64
	seed int64
}

// recordSampleCast appends one cast. Spell.applyEffects calls it after
// the effects have run, so the resources read are the ones the player saw
// after the cast - costs already spent, and any rage or energy the cast
// itself generated already added.
func (sim *Simulation) recordSampleCast(spell *Spell, target *Unit) {
	if sim.sample == nil || spell.Unit != sim.sample.unit {
		return
	}

	cast := &proto.SampleCast{
		AtMs:      sim.CurrentTime.Milliseconds(),
		ActionId:  spell.ActionID.ToProto(),
		Resources: unitResources(spell.Unit),
	}
	if target != nil {
		cast.Target = target.Label
	}
	sim.sample.casts = append(sim.sample.casts, cast)
}

// unitResources reads the bars the unit actually has. A mage has no rage
// bar and an empty key would read as "0 rage" on the report, which is a
// different claim from "this class has no rage".
func unitResources(unit *Unit) map[string]int32 {
	resources := make(map[string]int32, 4)
	if unit.HasManaBar() {
		resources["mana"] = int32(unit.CurrentMana())
	}
	if unit.HasRageBar() {
		resources["rage"] = int32(unit.CurrentRage())
	}
	if unit.HasEnergyBar() {
		resources["energy"] = int32(unit.CurrentEnergy())
		resources["combo_points"] = unit.ComboPoints()
	}
	return resources
}

// sampleUnit is the unit whose casts the log is of: the first player in
// the raid. Every sim the site sends is a single character, and a raid
// sim's sample log has no single subject to be about.
func (sim *Simulation) sampleUnit() *Unit {
	for _, party := range sim.Raid.Parties {
		for _, player := range party.Players {
			return &player.GetCharacter().Unit
		}
	}
	return nil
}

// runSampleIteration replays the median-DPS iteration of a finished run.
//
// It builds a FRESH Simulation from the same request rather than reusing
// the finished one, because runOnce ends in Cleanup, which feeds
// doneIteration, which would add a twenty-first sample to a
// twenty-iteration run's aggregates. A fresh environment costs one
// construction and buys exact aggregates.
func runSampleIteration(rsr *proto.RaidSimRequest, samples []iterationSample, baseDuration time.Duration, signals simsignals.Signals) *proto.SampleIteration {
	if len(samples) == 0 {
		return nil
	}

	ordered := slices.Clone(samples)
	slices.SortStableFunc(ordered, func(a, b iterationSample) int {
		switch {
		case a.dps < b.dps:
			return -1
		case a.dps > b.dps:
			return 1
		default:
			return 0
		}
	})
	median := ordered[len(ordered)/2]

	sim := NewSim(rsr, signals)
	unit := sim.sampleUnit()
	if unit == nil {
		return nil
	}

	// A health fight's duration is estimated from the presim; carry the
	// finished run's estimate over so the replay is the same fight.
	sim.BaseDuration = baseDuration
	sim.Encounter.DurationIsEstimate = false

	sim.sample = &sampleRecorder{unit: unit}
	sim.reseedTo(median.seed)
	sim.runOnce()

	return &proto.SampleIteration{
		Dps:             median.dps,
		DurationSeconds: sim.Duration.Seconds(),
		Casts:           sim.sample.casts,
	}
}
```

In `sim/core/sim.go`, add the recorder and the per-iteration buffer to the
`Simulation` struct, after the `Log` field:

```go
	// sample, when set, collects one unit's casts for the sample-iteration
	// log. It is only ever set on the throwaway Simulation the replay
	// builds, never on the one whose metrics are reported.
	sample *sampleRecorder

	// iterationSamples is (dps, seed) per iteration, filled only when the
	// request asked for a sample log.
	iterationSamples []iterationSample
```

In `sim/core/spell.go`, change `applyEffects`:

```go
func (spell *Spell) applyEffects(sim *Simulation, target *Unit) {
	spell.SpellMetrics[target.UnitIndex].Casts++
	spell.casts++

	spell.ApplyEffects(sim, target, spell)

	// After the effects, so the resources recorded are the ones the
	// player saw after the cast.
	sim.recordSampleCast(spell, target)
}
```

In `sim/core/sim.go`'s `run()`, record the pairs. Immediately before
`sim.runOnce()` (the first, un-looped one), add:

```go
	sampling := sim.Options.SampleIteration
	if sampling {
		sim.iterationSamples = make([]iterationSample, 0, sim.Options.Iterations)
	}
	sampleUnit := sim.sampleUnit()
	if sampleUnit == nil {
		sampling = false
	}
```

After that first `sim.runOnce()` and its duration bookkeeping, add:

```go
	if sampling {
		sim.iterationSamples = append(sim.iterationSamples, iterationSample{
			dps:  sampleUnit.Metrics.dps.LastIterationValue,
			seed: sim.iterationSeed(0),
		})
	}
```

and inside the `for i := int32(1); ...` loop, immediately after the loop's
`sim.runOnce()` and its duration bookkeeping, add:

```go
		if sampling {
			sim.iterationSamples = append(sim.iterationSamples, iterationSample{
				dps:  sampleUnit.Metrics.dps.LastIterationValue,
				seed: sim.iterationSeed(i),
			})
		}
```

In `runSim`, after `result = sim.run()` and before `return result`, add:

```go
	if rsr.SimOptions.SampleIteration && result.Error == nil {
		result.SampleIteration = runSampleIteration(rsr, sim.iterationSamples, sim.BaseDuration, signals)
	}
```

- [ ] **Step 6: Pick a sample when shards are combined**

In `sim/core/sim_concurrent.go`, at the end of `CombineConcurrentSimResults`,
before `return rsrc.Combined`, add:

```go
	rsrc.Combined.SampleIteration = pickSampleIteration(results, rsrc.Combined.RaidMetrics.Dps.Avg)
```

and add the function under it:

```go
// pickSampleIteration chooses which shard's sample log survives the
// combine: the one whose iteration is closest to the combined mean. Each
// shard picked its own median, and the median of one eighth of the run is
// not the median of the run.
func pickSampleIteration(results []*proto.RaidSimResult, mean float64) *proto.SampleIteration {
	var best *proto.SampleIteration
	bestDistance := math.Inf(1)
	for _, result := range results {
		sample := result.SampleIteration
		if sample == nil {
			continue
		}
		if distance := math.Abs(sample.Dps - mean); distance < bestDistance {
			best, bestDistance = sample, distance
		}
	}
	return best
}
```

Add `"math"` to that file's imports.

- [ ] **Step 7: Run the test to verify it passes**

Run: `cd /Users/jh/code/wowsims-forever && go test --tags=with_db ./sim/core/ -run 'Sample|IterationSeed' -v`
Expected: PASS, all six tests.

- [ ] **Step 8: Run the whole suite so no golden moved**

Run: `go test --tags=with_db -count=1 ./sim/...`
Expected: PASS. `SampleIteration` defaults false, so the recorder is nil, the
extra branch in `applyEffects` is one nil check, and no golden may move.

- [ ] **Step 9: gofmt and commit**

```bash
cd /Users/jh/code/wowsims-forever
gofmt -l sim/core/sample_iteration.go sim/core/sample_iteration_test.go sim/core/sim.go sim/core/spell.go sim/core/metrics_aggregator.go sim/core/sim_concurrent.go
git add sim/core/sample_iteration.go sim/core/sample_iteration_test.go sim/core/sim.go sim/core/spell.go sim/core/metrics_aggregator.go sim/core/sim_concurrent.go
git commit -m "$(cat <<'MSG'
feat(core): the median iteration's cast log, with resources

Two passes, not a buffer: the run records sixteen bytes an iteration -
its DPS and its seed - and the median iteration is then replayed once on
a fresh Simulation with a recorder attached. Fresh, because runOnce ends
in Cleanup and reusing the finished sim would add a twenty-first sample
to a twenty-iteration run. A concurrent combine keeps the shard sample
closest to the combined mean.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
MSG
)"
```

---

### Task 10: The sample log, end to end, on a fury warrior and a frost mage

**Files:**
- Create: `sim/sample_iteration_e2e_test.go` (package `sim`)

**Interfaces:**
- Consumes: `furyWarriorPlayer()`, `frostMagePlayer()`, `parityEncounter()`,
  `parityIterations` from Task 3's file, same package.
- Produces: nothing later tasks use.

This is the task that proves the contract's `SampleCast` can actually be filled:
a rage reading for the warrior, a mana reading for the mage, and a pre-pull cast
at a negative time for both.

- [ ] **Step 1: Write the failing test**

Create `sim/sample_iteration_e2e_test.go`:

```go
package sim

import (
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
)

func runSampleSim(t *testing.T, player *proto.Player) *proto.SampleIteration {
	t.Helper()
	result := core.RunSim(&proto.RaidSimRequest{
		Raid:      core.SinglePlayerRaidProto(player, &proto.PartyBuffs{}, &proto.RaidBuffs{}, &proto.Debuffs{}),
		Encounter: parityEncounter(),
		SimOptions: &proto.SimOptions{
			Iterations:      parityIterations,
			IsTest:          true,
			RandomSeed:      1,
			SampleIteration: true,
		},
	}, nil, simsignals.CreateSignals())
	if result.Error != nil {
		t.Fatalf("sim failed: %s", result.Error.Message)
	}
	if result.SampleIteration == nil {
		t.Fatal("SampleIteration = nil, want a log")
	}
	return result.SampleIteration
}

// A fury warrior's log reads rage after every cast, and never mana.
func TestFuryWarriorSampleLogReadsRage(t *testing.T) {
	sample := runSampleSim(t, furyWarriorPlayer())

	sawRage := false
	for _, cast := range sample.Casts {
		if _, ok := cast.Resources["rage"]; ok {
			sawRage = true
		}
		if _, ok := cast.Resources["mana"]; ok {
			t.Fatalf("a warrior's cast at %dms reported mana", cast.AtMs)
		}
		if cast.Resources["rage"] < 0 || cast.Resources["rage"] > 100 {
			t.Fatalf("rage %d at %dms is outside the bar", cast.Resources["rage"], cast.AtMs)
		}
	}
	if !sawRage {
		t.Error("no cast in the warrior's sample log reported rage")
	}
}

// A frost mage's log reads mana after every cast, and never rage.
func TestFrostMageSampleLogReadsMana(t *testing.T) {
	sample := runSampleSim(t, frostMagePlayer())

	sawMana := false
	for _, cast := range sample.Casts {
		if _, ok := cast.Resources["mana"]; ok {
			sawMana = true
		}
		if _, ok := cast.Resources["rage"]; ok {
			t.Fatalf("a mage's cast at %dms reported rage", cast.AtMs)
		}
	}
	if !sawMana {
		t.Error("no cast in the mage's sample log reported mana")
	}
}

// Pre-pull casts are what the report separates out, and they are
// identified by a negative time. Both reference rotations open with one.
func TestSampleLogSeparatesThePrePull(t *testing.T) {
	for _, tc := range []struct {
		name   string
		player func() *proto.Player
	}{
		{"fury warrior", furyWarriorPlayer},
		{"frost mage", frostMagePlayer},
	} {
		t.Run(tc.name, func(t *testing.T) {
			sample := runSampleSim(t, tc.player())

			prePull := 0
			for _, cast := range sample.Casts {
				if cast.AtMs < 0 {
					prePull++
				}
			}
			if prePull == 0 {
				t.Errorf("no pre-pull cast in the log; the reference rotation opens with one, so a negative time is never produced")
			}
			if sample.Casts[0].AtMs >= 0 && prePull > 0 {
				t.Error("a pre-pull cast is not first in the log; the log is not in time order")
			}
		})
	}
}

// Every cast names a target and a spell, so the report can render a row.
func TestSampleLogCastsAreRenderable(t *testing.T) {
	sample := runSampleSim(t, furyWarriorPlayer())
	for _, cast := range sample.Casts {
		if cast.ActionId == nil {
			t.Fatalf("a cast at %dms has no ActionId", cast.AtMs)
		}
		if cast.ActionId.GetSpellId() == 0 && cast.ActionId.GetItemId() == 0 && cast.ActionId.GetOtherId() == 0 {
			t.Fatalf("a cast at %dms has an empty ActionId", cast.AtMs)
		}
	}
}
```

- [ ] **Step 2: Run the test to verify it fails**

Temporarily change `runSim` in `sim/core/sim.go` to not call
`runSampleIteration`, run:
`cd /Users/jh/code/wowsims-forever && go test --tags=with_db ./sim/ -run 'Sample' -v`
Expected: FAIL — `SampleIteration = nil, want a log`. Restore the call.

- [ ] **Step 3: Run the test to verify it passes**

Run: `go test --tags=with_db ./sim/ -run 'Sample' -v`
Expected: PASS, all four tests (five with the subtests).

If `TestSampleLogSeparatesThePrePull` fails because a reference rotation has no
pre-pull action, do not weaken the test: add a pre-pull to the request instead
by setting `player.Cooldowns` — but check first with
`grep -c prepull ui/warrior/apls/forever_fury.apl.json`, because both reference
rotations do have one.

- [ ] **Step 4: Commit**

```bash
cd /Users/jh/code/wowsims-forever
gofmt -l sim/sample_iteration_e2e_test.go
git add sim/sample_iteration_e2e_test.go
git commit -m "$(cat <<'MSG'
test(core): the sample log, end to end, on fury and frost

The four things the contract's SampleCast needs, proved on a real
rotation each: rage on the warrior and never mana, mana on the mage and
never rage, a pre-pull cast at a negative time first in the log, and an
ActionId on every row so the report can render it.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
MSG
)"
```

---

### Task 11: The divergence note and the sha handoff

**Files:**
- Modify: `PORTING.md` (add a section before "No Forever artifacts are built here")

**Interfaces:**
- Consumes: everything above.
- Produces: the short sha the site's sim module lane needs.

`PORTING.md` is where this fork records what it added on top of
`wowsims/classic`, so the next upstream merge knows what is ours. Four proto
messages and five fields went in; they belong on that list.

- [ ] **Step 1: Add the section**

Insert into `PORTING.md`, immediately before the `## No Forever artifacts are
built here` heading:

```markdown
## Encounter and result fields upstream does not have

The Forever Sixty site's fight styles and sample-iteration report need
five fields upstream has no equivalent for. They are additive: an
encounter that sets none behaves exactly as it did before, which is why
no `*.results` golden moved when they landed.

| Field | Message | Number | What it does |
| --- | --- | --- | --- |
| `movement` | `Encounter` | 10 | A repeating window out of melee, or (with `casting_only`) an interrupt in place |
| `targets_over_time` | `Encounter` | 11 | A target-count timeline; the pool is sized to the largest count |
| `target_dummy` | `Encounter` | 12 | No raid debuff panel, no execute window, no armor reduction |
| `sample_iteration` | `RaidSimResult` | 8 | The median-DPS iteration's cast log, with resources after each cast |
| `sample_iteration` | `SimOptions` | 10 | Opt in to the above; it costs one extra iteration and one environment |

Supporting messages: `MovementPattern` and `TargetCountAt` in
`common.proto`, `SampleCast` and `SampleIteration` in `api.proto`.

A `SampleCast` carries the whole `ActionID`, not a display name: this
engine has none — `core.Spell` has no name field at all — and the site
already resolves action ids to names for its cast timeline.

The behaviour lives in `sim/core/encounter_movement.go`,
`sim/core/encounter_targets.go` and `sim/core/sample_iteration.go`, all
three of which are Forever files with no upstream counterpart, so an
upstream merge cannot conflict with them.
```

- [ ] **Step 2: Commit**

```bash
cd /Users/jh/code/wowsims-forever
git add PORTING.md
git commit -m "$(cat <<'MSG'
docs(sim): record the five parity fields as deliberate divergences

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
MSG
)"
```

- [ ] **Step 3: Run the fork's whole gate one last time**

```bash
cd /Users/jh/code/wowsims-forever
go test --tags=with_db ./sim/core/proto/ -run TestGeneratedProtosMatchSources -v
make binary_dist/dist.go
go test --tags=with_db -count=1 ./sim/...
go test --tags=with_db -count=1 ./tools/...
GOOS=js GOARCH=wasm go build ./sim/...
git status --porcelain
```
Expected: every command passes and `git status --porcelain` is empty. A dirty
tree cannot be pinned — the site's `make engine-pin` refuses one.

- [ ] **Step 4: Print the sha to hand over**

```bash
cd /Users/jh/code/wowsims-forever
git rev-parse --short=9 HEAD
```

Hand that nine-character sha to the **sim module lane**, which owns the pin
bump. The engine lane never edits the site repository. What the sim module lane
does with it:

1. `cd /Users/jh/code/forever && make engine-pin` — rewrites
   `sim/enginever/version.go` from this checkout's `HEAD`. It refuses a dirty
   engine tree, which Step 3 has already ruled out. Verify the constant in
   `sim/enginever/version.go` now equals the sha printed above (it replaces
   `edc0c8e9a`).
2. CI's `.github/actions/pin-engine` then rewrites `sim/go.mod`'s development
   replace to `github.com/jhunthrop/wowsims-forever@<that sha>`; nothing to do
   by hand. The site's `sim/` module consumes the generated Go protobuf package
   straight out of the fork at that pseudo-version, which is why Task 1 and
   Task 8 had to commit `sim/core/proto/*.pb.go` alongside the `.proto` edit.

There is a **second** pin, and it is not the sim module lane's: the data lane
vendors the engine's `.proto` files into `data/proto/` with the sha recorded in
`data/proto/ENGINE_SHA` (currently `464d1a14a`). Tasks 1 and 8 changed
`common.proto` and `api.proto`, so that copy is now stale. Tell the **data
lane** to run `python -m pipeline genproto` against this checkout, which
re-copies the three vendored `.proto` files, regenerates
`data/pipeline/simproto/` and rewrites `data/proto/ENGINE_SHA` to the same sha.
The parity design's section 12 lane list does not mention this pin; it is real
and it must be bumped in the same round or the data pipeline's bindings will not
know the `Encounter` fields exist.

---

## Self-review

**Spec coverage** (contract section 5, the binding one for this lane):

| Requirement | Task |
| --- | --- |
| `MovementPattern` message, `movement = 10` | 1 |
| `TargetCountAt` message, `targets_over_time = 11` | 1 |
| `target_dummy = 12` | 1 |
| Movement schedules moves out of range every interval for a duration | 2, 3 |
| `casting_only` interrupts casting without moving | 2, 3 |
| Targets activate and deactivate on a timeline | 4, 5 |
| Pool sized to the maximum count | 4 |
| `target_dummy` disables debuff application | 6, 7 |
| `target_dummy` disables execute windows | 6, 7 |
| `target_dummy` disables the target's armor reduction | 6 |
| `RaidSimResult.sample_iteration`, median-DPS iteration | 8, 9 |
| Resource readings after each cast (rage, energy, mana, combo points) | 9, 10 |
| Pre-pull casts at negative times | 10 |
| Regenerating the Go protobufs the way this repo does | 1, 8 (`make proto` + `TestGeneratedProtosMatchSources`) |
| The site's vendored copy / pin handoff | 11 |
| Tests per task in the fork's own style | every task |
| Fury warrior and frost mage end to end per encounter feature | 3, 5, 7, 10 |

**Types used consistently across tasks:** `MovementPattern` (core) in Tasks 2
and 3; `TargetCount` / `AllTargetUnits` / `SetActiveTargetCount` in Tasks 4 and
5; `Encounter.Dummy` / `ArmorReductionDisabled` in Tasks 6 and 7;
`iterationSample` / `sampleRecorder` / `iterationSeed` / `reseedTo` /
`LastIterationValue` in Task 9 only; `furyWarriorPlayer` / `frostMagePlayer` /
`parityEncounter` / `runParitySim` / `parityIterations` defined once in Task 3
and reused by Tasks 5, 7 and 10 in the same package.
