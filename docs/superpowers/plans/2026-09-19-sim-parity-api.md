# Simulator parity — the API lane — Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Teach the API the four new request kinds — gear, talents, drops, weights — so the premium lane sizes a request before it queues it, the job streams stage progress, the history list names every kind, and the two small routes Top Gear needs (`GET /v1/builds?mine=1`, `GET /v1/phases`) exist.

**Architecture:** Nothing about *which* sims run lives in the API. `sim/bulk` owns expansion, staging and ranking; the API reaches it exactly twice — once before queueing, through `forever-sim -plan`, which counts combinations without running any of them, and once inside the job, where the binary detects the kind and runs its own plan-rank loop. The API's own new work is four things: the `kind` column and the index the filtered history query needs, a pure `Headline` function that turns a stored result into one line of text, two submit-time refusals (`cap_exceeded`, `too_large`), and two list routes.

**Tech Stack:** Go 1.25 (workspace, `GOTOOLCHAIN=auto`), pgx v5, golang-migrate (embedded `api/internal/db/migrations/*.sql`), `net/http` `ServeMux`, the shared `httpx` envelope.

**Spec:**
- Design: `docs/superpowers/specs/2026-09-19-simulator-parity-design.md` (sections 10 and 11)
- Contract: `docs/superpowers/specs/2026-09-19-simulator-parity-interfaces.md` — section 8 is this lane's binding interface, **and section 10 overrides sections 1–9 wherever they disagree**. Section 10.6 is this lane's correction list; 10.1 (A1, A2, A3, A6, A7, A11) and 10.2 are what it consumes.

## Global Constraints

- **The contract is the authority, and section 10 wins inside it.** Section 10 is the amendment that earlier planning rounds asked for; there is no "amend the contract" task in this plan.
- **Go workspace toolchain:** `GOTOOLCHAIN=auto`. `go.work` uses `./api ./companion ./logs ./sim`; run every Go command from the repository root.
- **The api module must never import the engine.** `api/internal/sims`'s package doc says which halves of the `sim` module are importable: `sim/api`, `sim/enginever`, `sim/specs`, `sim/runner`, and — added by Task 6 — `sim/measure`. All five are engine-free. `sim/bulk` is not: it reads `sim/internal/simdb`, which imports `github.com/wowsims/classic/sim/core/proto`, and `sim/internal/...` is closed to other modules besides. Every combination count comes from the native binary as a subprocess (contract 10.2).
- **DB tests share one database.** `TEST_DATABASE_URL` points at a single Postgres; `newHarness` truncates on entry and the tests use literal twelve-character ids. Run the package's tests serially — `go test -p 1 ./api/internal/sims/...` — and if a run fails with a `sims_pkey` duplicate-key collision, re-run it once before investigating; a second collision is a real bug (two tests sharing a literal id).
- **Scoped tests per task:** `go test ./api/internal/sims/...` (or the package the task touches), `-run` for a single test while iterating. Never run the whole repository suite for a task — that is CI's job.
- **One commit per task**, message `feat(api): ...` or `test(api): ...`, ending with the trailer:

  ```
  Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
  ```
- **Never `git stash`.** The stash stack is shared with other worktrees and other sessions. Set work aside with a temporary WIP commit instead.
- **Migrations:** `0014_sim_kinds` (Task 1) and `0015_builds_owner` (Task 9). Everything one feature needs goes in its one file.
- **Numbers this lane must not re-derive.** The server cap is `simapi.Caps[simapi.LaneServer]`, which section 10.1's A2 sets to **5,000**; the engine's rate is `measure.NativeIterationsPerCPUSecond`, which `sim/measure` publishes from its own benchmark (1,218 today); the bulk budget is **840 seconds** against a job timeout that stays at 15 minutes. Write the names, never the literals.
- **Ordering.** Tasks 1–5 and 9–10 depend on nothing outside this lane. Tasks 6, 7, 8 and 11 name what they wait for from the `sim` and `data` lanes at the top of the task.

---

## File Structure

| File | Responsibility |
| --- | --- |
| `api/internal/db/migrations/0014_sim_kinds.up.sql` / `.down.sql` | New: `kind`, `headline`, `stage`, `combos_done`, `combos_total`, and the `(user_id, kind, created_at desc)` index. |
| `api/internal/db/migrations/0015_builds_owner.up.sql` / `.down.sql` | New: `builds.user_id` and its index. |
| `api/internal/sims/headline.go` | New: `Headline` and its formatting helpers. Pure functions, no I/O. |
| `api/internal/sims/headline_test.go` | New: the headline table test. No database. |
| `api/internal/sims/store.go` | `Save`/`Queue`/`Finish` write `kind` and `headline`; `Mine` filters by kind; `Tick`, `Advance` and `Progress` carry the stage fields. |
| `api/internal/sims/handler.go` | `Mount` requires a `Planner`; `mine` reads `kind=`; `trimTitle` generalised to `trimRunes`. |
| `api/internal/sims/run.go` | Forces the server lane's cap, validates per lane, then `cap_exceeded` / `too_large` before it queues. |
| `api/internal/sims/job.go` | Per-kind run timeout; the widened progress callback adapted to a `Tick`. |
| `api/internal/sims/simdep.go` | The `Planner` and `Engine` interfaces, and the estimate. |
| `api/internal/sims/specs.go` | `SpecFidelity.ReferenceStat`, read from `sim/specs`. |
| `api/internal/builds/store.go`, `handler.go` | An owner on a saved build, `Mine`, and one shared row scanner. |
| `api/internal/phase/phase.go` | JSON tags on `Boundary`, so the route can serve the table. |
| `api/internal/server/server.go` | `GET /v1/phases`. |
| `api/cmd/api/main.go` | `simEngine` returns a `sims.Engine`; the service gets a `Planner`; builds gets the accounts store. |
| `api/openapi.yaml`, `api/README.md` | The routes, schemas and operator documentation. |

---

### Task 1: Migration 0014 and the kind column

**Files:**
- Create: `api/internal/db/migrations/0014_sim_kinds.up.sql`
- Create: `api/internal/db/migrations/0014_sim_kinds.down.sql`
- Modify: `api/internal/sims/store.go` (`Save`, `Queue`)
- Test: `api/internal/sims/store_test.go`

**Interfaces:**
- Consumes: `simapi.SimRequest.Kind() string` and the `simapi.Kind*` constants (contract 1.1); `simapi.BulkSpec`, `simapi.WeightsSpec` (contract 1.3, 1.4).
- Produces: the `sims.kind`, `sims.headline`, `sims.stage`, `sims.combos_done` and `sims.combos_total` columns, and `sims_user_kind_idx`. Tasks 3, 4 and 7 read them.

- [ ] **Step 1: Write the failing test**

Add to `api/internal/sims/store_test.go`:

```go
func TestEveryRowRecordsWhichToolProducedIt(t *testing.T) {
	h := newHarness(t)
	// A browser save of a plain run.
	if err := h.store.Save(t.Context(), "aaaaaaaaaaaa", &h.owner, "",
		browserResult("warrior-fury", 1000)); err != nil {
		t.Fatal(err)
	}
	// A queued server run of a Top Gear request.
	gear := browserResult("warrior-fury", 0).Request
	gear.Bulk = &simapi.BulkSpec{
		Mode:      simapi.KindGear,
		Precision: simapi.PrecisionNormal,
		Cap:       simapi.Caps[simapi.LaneServer],
		Candidates: []simapi.Candidate{
			{Slot: "main_hand", ItemID: 19019, Origin: "bag"},
		},
	}
	if err := h.store.Queue(t.Context(), "bbbbbbbbbbbb", h.owner, gear); err != nil {
		t.Fatal(err)
	}
	// A queued weights run.
	weights := browserResult("warrior-fury", 0).Request
	weights.Weights = &simapi.WeightsSpec{
		Stats: []string{"strength", "melee_crit"}, Reference: "melee_crit",
	}
	if err := h.store.Queue(t.Context(), "cccccccccccc", h.owner, weights); err != nil {
		t.Fatal(err)
	}

	for _, c := range []struct{ id, want string }{
		{"aaaaaaaaaaaa", simapi.KindRun},
		{"bbbbbbbbbbbb", simapi.KindGear},
		{"cccccccccccc", simapi.KindWeights},
	} {
		var got string
		if err := h.store.Pool.QueryRow(t.Context(),
			`select kind from sims where id = $1`, c.id).Scan(&got); err != nil {
			t.Fatal(err)
		}
		if got != c.want {
			t.Errorf("%s: kind %q, want %q", c.id, got, c.want)
		}
	}
}
```

(The stat ids are the fork's `proto.Stat` enum names in snake case, per contract A7: `melee_crit`, never a bare `crit`.)

- [ ] **Step 2: Run it and watch it fail**

Run: `go test -p 1 ./api/internal/sims/... -run TestEveryRowRecordsWhichToolProducedIt -v`
Expected: FAIL — `column "kind" does not exist`. A build failure on `simapi.BulkSpec` instead means the `sim` lane's envelope (contract 1.3) is not merged yet; wait for it rather than stubbing the type.

- [ ] **Step 3: Write the migration**

`api/internal/db/migrations/0014_sim_kinds.up.sql`:

```sql
-- Every sim row says which tool produced it, carries the one line the
-- history list shows for it, and — while a bulk run is going — how far
-- through its stages it is (contract 10.6).
--
-- headline is stored rather than composed on read: the result blob is a
-- few hundred kilobytes and a history page is a hundred rows, so reading
-- it back to compose one line would detoast tens of megabytes per page.
-- A stored result never changes, so a stored headline never goes stale.

alter table sims add column if not exists kind         text not null default 'run';
alter table sims add column if not exists headline     text not null default '';
alter table sims add column if not exists stage        int  not null default 0;
alter table sims add column if not exists combos_done  int  not null default 0;
alter table sims add column if not exists combos_total int  not null default 0;

-- The history list filtered by kind. sims_user_idx still serves the
-- unfiltered list, whose ordering this index cannot satisfy.
create index if not exists sims_user_kind_idx on sims (user_id, kind, created_at desc);
```

`api/internal/db/migrations/0014_sim_kinds.down.sql`:

```sql
drop index if exists sims_user_kind_idx;
alter table sims drop column if exists combos_total;
alter table sims drop column if exists combos_done;
alter table sims drop column if exists stage;
alter table sims drop column if exists headline;
alter table sims drop column if exists kind;
```

- [ ] **Step 4: Write kind from both writers**

In `api/internal/sims/store.go`, in `Save`, replace the insert with:

```go
	_, err = s.Pool.Exec(ctx,
		`insert into sims (id, user_id, spec, kind, engine_version, lane, dps_mean, dps_error,
		   iterations, title, result, state)
		 values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		 on conflict (id) do nothing`,
		id, userID, res.Request.Spec, res.Request.Kind(), res.EngineVersion, res.Lane,
		res.DPS.Mean, res.DPS.Error, res.IterationsRun, t, body, StateDone)
```

and in `Queue`:

```go
	_, err = s.Pool.Exec(ctx,
		`insert into sims (id, user_id, spec, kind, engine_version, lane, dps_mean, dps_error,
		   iterations, result, state)
		 values ($1, $2, $3, $4, $5, $6, 0, 0, $7, $8, $9)`,
		id, userID, req.Spec, req.Kind(), req.EngineVersion, simapi.LaneServer,
		req.Iterations, body, StateQueued)
```

- [ ] **Step 5: Run the test**

Run: `go test -p 1 ./api/internal/sims/... -run TestEveryRowRecordsWhichToolProducedIt -v`
Expected: PASS

- [ ] **Step 6: Run the whole package**

Run: `go test -p 1 ./api/internal/sims/...`
Expected: PASS. Re-run once on a `sims_pkey` collision.

- [ ] **Step 7: Commit**

```bash
git add api/internal/db/migrations/0014_sim_kinds.up.sql \
        api/internal/db/migrations/0014_sim_kinds.down.sql \
        api/internal/sims/store.go api/internal/sims/store_test.go
git commit -m "$(cat <<'EOF'
feat(api): migration 0014, every sim row records its kind

sims gains kind, headline and the three bulk-progress columns, plus the
(user_id, kind, created_at desc) index the filtered history query needs.
Save and Queue write req.Kind().

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

### Task 2: The headline

One pure function, one table test, no database. Every number the history list shows is formatted here and nowhere else. The forms are contract section 8's, plus 10.6's two additions: several substitutions read "… and 2 more", and an empty result reads "no combinations", "no upgrades" or "no weights".

**Files:**
- Create: `api/internal/sims/headline.go`
- Create: `api/internal/sims/headline_test.go`
- Modify: `api/internal/sims/handler.go` (generalise `trimTitle`)

**Interfaces:**
- Consumes: `simapi.SimResult`, `simapi.Combo`, `simapi.StatWeight`, and `simapi.Substitution` with `Name` filled for items and `SourceName` copied from the candidate (contract A6); `simapi.SimRequest.Kind()`.
- Produces:
  - `func Headline(res simapi.SimResult) string`
  - `func withThousands(n int64) string`
  - `func trimRunes(s string, max int) string`
  - `const maxHeadline = 120`

- [ ] **Step 1: Write the failing test**

Create `api/internal/sims/headline_test.go`:

```go
package sims

import (
	"strings"
	"testing"
	"unicode/utf8"

	simapi "github.com/jhunthrop/foreversixty/sim/api"
)

// withKind returns a result of the given kind, so each case below says
// only what it is about.
func withKind(kind string, mutate func(*simapi.SimResult)) simapi.SimResult {
	res := browserResult("warrior-fury", 1204.4)
	switch kind {
	case simapi.KindGear, simapi.KindTalents, simapi.KindDrops:
		res.Request.Bulk = &simapi.BulkSpec{Mode: kind, Precision: simapi.PrecisionNormal}
	case simapi.KindWeights:
		res.Request.Weights = &simapi.WeightsSpec{Reference: "melee_crit"}
	}
	if mutate != nil {
		mutate(&res)
	}
	return res
}

func item(name string, delta float64) simapi.Combo {
	return simapi.Combo{
		Substitutions: []simapi.Substitution{
			{Kind: "item", Slot: "main_hand", ItemID: 17182, Name: name, Origin: "bag"},
		},
		Delta: simapi.Estimate{Mean: delta},
	}
}

func drop(name, source string, delta float64) simapi.Combo {
	return simapi.Combo{
		Substitutions: []simapi.Substitution{
			{Kind: "item", Slot: "main_hand", ItemID: 17182, Name: name,
				Origin: "drop:raid:molten-core:11502", SourceName: source},
		},
		Delta: simapi.Estimate{Mean: delta},
	}
}

func TestTheHeadlineSaysWhatEachKindFound(t *testing.T) {
	for _, c := range []struct {
		name string
		res  simapi.SimResult
		want string
	}{
		{
			"a run is its own dps, grouped",
			withKind(simapi.KindRun, nil),
			"1,204 DPS",
		},
		{
			"a three-figure run has no comma",
			withKind(simapi.KindRun, func(r *simapi.SimResult) { r.DPS.Mean = 999.4 }),
			"999 DPS",
		},
		{
			"a five-figure run groups once",
			withKind(simapi.KindRun, func(r *simapi.SimResult) { r.DPS.Mean = 12345.6 }),
			"12,346 DPS",
		},
		{
			"top gear names the best combination's item",
			withKind(simapi.KindGear, func(r *simapi.SimResult) {
				r.Combos = []simapi.Combo{item("Vis'kag the Bloodletter", 41.2), item("Brutality Blade", 8)}
			}),
			"+41 DPS from Vis'kag the Bloodletter",
		},
		{
			"top gear with three substitutions names the first and counts the rest",
			withKind(simapi.KindGear, func(r *simapi.SimResult) {
				r.Combos = []simapi.Combo{{
					Substitutions: []simapi.Substitution{
						{Kind: "item", ItemID: 17182, Name: "Vis'kag the Bloodletter", Origin: "bag"},
						{Kind: "item", ItemID: 16963, Name: "Onslaught Girdle", Origin: "bag"},
						{Kind: "item", ItemID: 18404, Name: "Blackhand's Breadth", Origin: "bank"},
					},
					Delta: simapi.Estimate{Mean: 63.5},
				}}
			}),
			"+64 DPS from Vis'kag the Bloodletter and 2 more",
		},
		{
			"an unnamed item falls back to its id",
			withKind(simapi.KindGear, func(r *simapi.SimResult) {
				r.Combos = []simapi.Combo{item("", 12)}
			}),
			"+12 DPS from item #17182",
		},
		{
			"a losing best combination is still signed",
			withKind(simapi.KindGear, func(r *simapi.SimResult) {
				r.Combos = []simapi.Combo{item("Brutality Blade", -13.7)}
			}),
			"-14 DPS from Brutality Blade",
		},
		{
			"a gear run with nothing in it says so",
			withKind(simapi.KindGear, nil),
			"no combinations",
		},
		{
			"talents quote the loadout",
			withKind(simapi.KindTalents, func(r *simapi.SimResult) {
				r.Combos = []simapi.Combo{{
					Substitutions: []simapi.Substitution{
						{Kind: "talents", Name: "Deep Fury", Talents: "-0550000505021051-05"},
					},
					Delta: simapi.Estimate{Mean: 18.4},
				}}
			}),
			"+18 DPS with 'Deep Fury'",
		},
		{
			"drops count the upgrades and name the source",
			withKind(simapi.KindDrops, func(r *simapi.SimResult) {
				r.Combos = []simapi.Combo{
					drop("Perdition's Blade", "Ragnaros", 55),
					drop("Spinal Reaper", "Ragnaros", 12),
					drop("Malistar's Defender", "Ragnaros", 1),
					drop("Band of Accuria", "Ragnaros", -4),
				}
			}),
			"3 upgrades on Ragnaros",
		},
		{
			"one upgrade is singular",
			withKind(simapi.KindDrops, func(r *simapi.SimResult) {
				r.Combos = []simapi.Combo{drop("Spinal Reaper", "Ragnaros", 12)}
			}),
			"1 upgrade on Ragnaros",
		},
		{
			"a two-word source reads whole",
			withKind(simapi.KindDrops, func(r *simapi.SimResult) {
				r.Combos = []simapi.Combo{drop("Maladath", "Broodlord Lashlayer", 30)}
			}),
			"1 upgrade on Broodlord Lashlayer",
		},
		{
			"nothing gained on a source is still an answer",
			withKind(simapi.KindDrops, func(r *simapi.SimResult) {
				r.Combos = []simapi.Combo{drop("Band of Accuria", "Ragnaros", -4)}
			}),
			"no upgrades on Ragnaros",
		},
		{
			"two sources in one request name neither",
			withKind(simapi.KindDrops, func(r *simapi.SimResult) {
				r.Combos = []simapi.Combo{
					drop("Perdition's Blade", "Ragnaros", 55),
					drop("Maladath", "Broodlord Lashlayer", 30),
				}
			}),
			"2 upgrades",
		},
		{
			"an unnamed source is counted, not named",
			withKind(simapi.KindDrops, func(r *simapi.SimResult) {
				r.Combos = []simapi.Combo{drop("Maladath", "", 30)}
			}),
			"1 upgrade",
		},
		{
			"weights are the top two, the reference first",
			withKind(simapi.KindWeights, func(r *simapi.SimResult) {
				r.Weights = []simapi.StatWeight{
					{Stat: "agility", Weight: 0.874},
					{Stat: "melee_crit", Weight: 1},
					{Stat: "attack_power", Weight: 0.5},
				}
			}),
			"Melee Crit 1.00 · Agility 0.87",
		},
		{
			"a multi-word stat id reads as words",
			withKind(simapi.KindWeights, func(r *simapi.SimResult) {
				r.Weights = []simapi.StatWeight{
					{Stat: "attack_power", Weight: 1},
					{Stat: "spell_haste", Weight: 0.4},
				}
			}),
			"Attack Power 1.00 · Spell Haste 0.40",
		},
		{
			"a weights run with nothing in it says so",
			withKind(simapi.KindWeights, nil),
			"no weights",
		},
	} {
		t.Run(c.name, func(t *testing.T) {
			if got := Headline(c.res); got != c.want {
				t.Errorf("Headline = %q, want %q", got, c.want)
			}
		})
	}
}

func TestALongHeadlineIsTrimmedRatherThanStoredWhole(t *testing.T) {
	res := withKind(simapi.KindGear, func(r *simapi.SimResult) {
		r.Combos = []simapi.Combo{item(strings.Repeat("ǝ", 400), 10)}
	})
	got := Headline(res)
	if n := len([]rune(got)); n != maxHeadline {
		t.Fatalf("%d runes, want %d", n, maxHeadline)
	}
	if !utf8.ValidString(got) {
		t.Error("the cut broke a rune in half")
	}
}

func TestThousandsAreGroupedFromTheRight(t *testing.T) {
	for _, c := range []struct {
		in   int64
		want string
	}{
		{0, "0"}, {7, "7"}, {999, "999"}, {1000, "1,000"}, {1204, "1,204"},
		{12345, "12,345"}, {1234567, "1,234,567"}, {-1500, "-1,500"},
	} {
		if got := withThousands(c.in); got != c.want {
			t.Errorf("withThousands(%d) = %q, want %q", c.in, got, c.want)
		}
	}
}
```

- [ ] **Step 2: Run it and watch it fail**

Run: `go test ./api/internal/sims/... -run 'TestTheHeadline|TestALongHeadline|TestThousands' -v`
Expected: FAIL to build — `undefined: Headline`, `undefined: withThousands`, `undefined: maxHeadline`.

- [ ] **Step 3: Write `headline.go`**

Create `api/internal/sims/headline.go`:

```go
package sims

import (
	"math"
	"slices"
	"strconv"
	"strings"
	"unicode"

	simapi "github.com/jhunthrop/foreversixty/sim/api"
)

// maxHeadline bounds the stored line, in runes. A set or loadout name is
// a member's own string, so the line it lands in needs a bound.
const maxHeadline = 120

// headlineWeights is how many stat weights the weights headline names:
// the reference and its nearest rival, which is what makes the line mean
// anything at a glance.
const headlineWeights = 2

// weightSeparator joins them. It is a middle dot with spaces, as the
// contract's example spells it.
const weightSeparator = " · "

// Headline is the one line a history row shows for a finished result.
// Contract section 8 fixes one form per kind and 10.6 adds the "… and N
// more" and empty-result forms; this function is the only place any of
// them is composed, and the only place a figure on that list is
// formatted.
//
// It reads the stored result and nothing else — no item table, no loot
// table — which is why sim/bulk fills Substitution.Name for items and
// copies SourceName onto the substitution (contract A6).
func Headline(res simapi.SimResult) string {
	var line string
	switch kind := res.Request.Kind(); kind {
	case simapi.KindWeights:
		line = weightsHeadline(res.Weights)
	case simapi.KindDrops:
		line = dropsHeadline(res.Combos)
	case simapi.KindGear, simapi.KindTalents:
		line = comboHeadline(kind, res.Combos)
	default:
		line = withThousands(roundDPS(res.DPS.Mean)) + " DPS"
	}
	return trimRunes(line, maxHeadline)
}

// comboHeadline is Top Gear's and talent compare's line: what the best
// combination gained, and what it swapped in to gain it.
func comboHeadline(kind string, combos []simapi.Combo) string {
	if len(combos) == 0 {
		return "no combinations"
	}
	best := combos[0]
	delta := signedDPS(best.Delta.Mean)
	name := substitutionName(best.Substitutions)
	switch {
	case name == "":
		return delta
	case kind == simapi.KindTalents:
		return delta + " with '" + name + "'"
	default:
		return delta + " from " + name
	}
}

// substitutionName names a combination: the first substitution, and how
// many others rode with it (contract 10.6). An item with no name falls
// back to its id rather than vanishing — an unnamed item is a data gap
// worth seeing on the page.
func substitutionName(subs []simapi.Substitution) string {
	if len(subs) == 0 {
		return ""
	}
	name := subs[0].Name
	if name == "" && subs[0].ItemID != 0 {
		name = "item #" + strconv.Itoa(subs[0].ItemID)
	}
	if name == "" {
		return ""
	}
	if rest := len(subs) - 1; rest > 0 {
		name += " and " + strconv.Itoa(rest) + " more"
	}
	return name
}

// dropsHeadline is Droptimizer's line: how many of a source's drops beat
// what is equipped, and which source they came from.
func dropsHeadline(combos []simapi.Combo) string {
	upgrades := 0
	for _, c := range combos {
		if c.Delta.Mean > 0 {
			upgrades++
		}
	}
	line := strconv.Itoa(upgrades) + " upgrades"
	switch upgrades {
	case 0:
		line = "no upgrades"
	case 1:
		line = "1 upgrade"
	}
	if source := sharedSource(combos); source != "" {
		return line + " on " + source
	}
	return line
}

// sharedSource is the one loot source every substitution came from, or
// "" when they came from more than one, or from one the page could not
// name. A Droptimizer run is normally one boss; a request that mixed two
// is counted without being named rather than named wrongly.
//
// The name is Substitution.SourceName (contract A6), which the page
// filled from loot.json when it built the request. The API has no loot
// table of its own and never invents one from an origin id.
func sharedSource(combos []simapi.Combo) string {
	source := ""
	for _, c := range combos {
		for _, s := range c.Substitutions {
			if s.SourceName == "" {
				return ""
			}
			if source == "" {
				source = s.SourceName
			} else if source != s.SourceName {
				return ""
			}
		}
	}
	return source
}

// weightsHeadline is the stat weights line: the two largest weights, each
// to two decimals. The reference stat is exactly 1 (contract 2), so it
// leads unless something beat it, which is itself worth seeing.
func weightsHeadline(weights []simapi.StatWeight) string {
	if len(weights) == 0 {
		return "no weights"
	}
	sorted := slices.Clone(weights)
	slices.SortFunc(sorted, func(a, b simapi.StatWeight) int {
		switch {
		case a.Weight > b.Weight:
			return -1
		case a.Weight < b.Weight:
			return 1
		default:
			// A stable order, so two equal weights do not swap between
			// reads and make the same result read differently twice.
			return strings.Compare(a.Stat, b.Stat)
		}
	})
	if len(sorted) > headlineWeights {
		sorted = sorted[:headlineWeights]
	}
	parts := make([]string, 0, len(sorted))
	for _, w := range sorted {
		parts = append(parts, statLabel(w.Stat)+" "+strconv.FormatFloat(w.Weight, 'f', 2, 64))
	}
	return strings.Join(parts, weightSeparator)
}

// statLabel turns a stat id into its display form: "attack_power" becomes
// "Attack Power", "melee_crit" becomes "Melee Crit". The vocabulary is
// the fork's proto.Stat enum names in snake case (contract A7), so there
// is no table to keep in step: a stat the engine gains is labelled the
// day it appears.
func statLabel(stat string) string { return titleWords(strings.ReplaceAll(stat, "_", " ")) }

// titleWords upper-cases the first rune of each word. It works by rune,
// so a name in any script keeps its characters whole.
func titleWords(s string) string {
	words := strings.Fields(s)
	for i, w := range words {
		r := []rune(w)
		words[i] = string(unicode.ToUpper(r[0])) + string(r[1:])
	}
	return strings.Join(words, " ")
}

// signedDPS is a delta: always signed, so "+0 DPS" reads as a comparison
// that found nothing rather than as an absolute figure.
func signedDPS(mean float64) string {
	n := roundDPS(mean)
	if n < 0 {
		return "-" + withThousands(-n) + " DPS"
	}
	return "+" + withThousands(n) + " DPS"
}

// roundDPS is how every DPS figure on the history list is rounded: to the
// nearest whole, halves away from zero, which is math.Round's own rule.
func roundDPS(mean float64) int64 { return int64(math.Round(mean)) }

// withThousands groups n into three-digit blocks with commas. There is no
// dependency for this: golang.org/x/text's printer is a locale package
// for one string, and the site's numbers are English-grouped everywhere
// else too.
func withThousands(n int64) string {
	sign := ""
	if n < 0 {
		sign, n = "-", -n
	}
	digits := strconv.FormatInt(n, 10)
	var b strings.Builder
	for i := range digits {
		if i > 0 && (len(digits)-i)%3 == 0 {
			b.WriteByte(',')
		}
		b.WriteByte(digits[i])
	}
	return sign + b.String()
}
```

- [ ] **Step 4: Generalise `trimTitle`**

In `api/internal/sims/handler.go`, replace the `trimTitle` function with:

```go
// trimTitle bounds the name a member gave a sim. A long title is cut
// rather than refused: it is a label, not data.
func trimTitle(s string) string { return trimRunes(s, maxTitle) }

// trimRunes cuts s to at most max runes. The cut is by rune, so a string
// in any script keeps its last character whole and the result is always
// valid UTF-8. Both the member's title and the composed headline are
// bounded this way.
func trimRunes(s string, max int) string {
	r := []rune(s)
	if len(r) > max {
		return string(r[:max])
	}
	return s
}
```

- [ ] **Step 5: Run the tests**

Run: `go test ./api/internal/sims/... -run 'TestTheHeadline|TestALongHeadline|TestThousands' -v`
Expected: PASS, all sub-tests.

- [ ] **Step 6: Run the whole package**

Run: `go test -p 1 ./api/internal/sims/...`
Expected: PASS.

- [ ] **Step 7: Commit**

```bash
git add api/internal/sims/headline.go api/internal/sims/headline_test.go \
        api/internal/sims/handler.go
git commit -m "$(cat <<'EOF'
feat(api): the history headline, one line per kind

Headline composes contract section 8's five forms, plus 10.6's "… and N
more" and empty-result forms, from a stored result and nothing else: no
item table, no loot table. Item names come from Substitution.Name and a
drop source from Substitution.SourceName, both filled by sim/bulk.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

### Task 3: The history row carries kind and headline

**Files:**
- Modify: `api/internal/sims/store.go` (`Row`, `Save`, `Finish`, `Mine`)
- Test: `api/internal/sims/store_test.go`

**Interfaces:**
- Consumes: `Headline(simapi.SimResult) string` (Task 2); the `kind` and `headline` columns (Task 1).
- Produces:
  - `Row` gains `Kind string \`json:"kind"\`` and `Headline string \`json:"headline"\``
  - `func (s *Store) Mine(ctx context.Context, userID int64, page int, kind string) (Page, error)` — `kind` `""` means every kind.

- [ ] **Step 1: Write the failing test**

Add to `api/internal/sims/store_test.go`:

```go
func TestMyHistoryCarriesTheKindAndHeadlineAndFiltersByKind(t *testing.T) {
	h := newHarness(t)
	plain := browserResult("warrior-fury", 1204.4)
	if err := h.store.Save(t.Context(), "aaaaaaaaaaaa", &h.owner, "", plain); err != nil {
		t.Fatal(err)
	}
	gear := browserResult("warrior-fury", 1204.4)
	gear.Request.Bulk = &simapi.BulkSpec{Mode: simapi.KindGear, Precision: simapi.PrecisionNormal}
	gear.Combos = []simapi.Combo{{
		Substitutions: []simapi.Substitution{
			{Kind: "item", ItemID: 17182, Name: "Vis'kag the Bloodletter", Origin: "bag"},
		},
		Delta: simapi.Estimate{Mean: 41.2},
	}}
	if err := h.store.Save(t.Context(), "bbbbbbbbbbbb", &h.owner, "", gear); err != nil {
		t.Fatal(err)
	}

	all, err := h.store.Mine(t.Context(), h.owner, 1, "")
	if err != nil {
		t.Fatal(err)
	}
	if all.Total != 2 {
		t.Fatalf("total %d, want both kinds", all.Total)
	}
	byID := map[string]Row{}
	for _, r := range all.Rows {
		byID[r.SimID] = r
	}
	if got := byID["aaaaaaaaaaaa"]; got.Kind != simapi.KindRun || got.Headline != "1,204 DPS" {
		t.Errorf("plain row: %+v", got)
	}
	if got := byID["bbbbbbbbbbbb"]; got.Kind != simapi.KindGear ||
		got.Headline != "+41 DPS from Vis'kag the Bloodletter" {
		t.Errorf("gear row: %+v", got)
	}

	only, err := h.store.Mine(t.Context(), h.owner, 1, simapi.KindGear)
	if err != nil {
		t.Fatal(err)
	}
	if only.Total != 1 || len(only.Rows) != 1 || only.Rows[0].SimID != "bbbbbbbbbbbb" {
		t.Fatalf("filtered: total %d rows %+v", only.Total, only.Rows)
	}

	none, err := h.store.Mine(t.Context(), h.owner, 1, simapi.KindDrops)
	if err != nil {
		t.Fatal(err)
	}
	if none.Total != 0 || len(none.Rows) != 0 {
		t.Fatalf("a kind with no rows: total %d rows %+v", none.Total, none.Rows)
	}
}

func TestAServerRunsHeadlineIsWrittenWhenItFinishes(t *testing.T) {
	h := newHarness(t)
	req := browserResult("warrior-fury", 0).Request
	req.Weights = &simapi.WeightsSpec{
		Stats: []string{"melee_crit", "agility"}, Reference: "melee_crit",
	}
	if err := h.store.Queue(t.Context(), "cccccccccccc", h.owner, req); err != nil {
		t.Fatal(err)
	}
	// A queued run has nothing to say yet.
	queued, err := h.store.Mine(t.Context(), h.owner, 1, simapi.KindWeights)
	if err != nil {
		t.Fatal(err)
	}
	if len(queued.Rows) != 1 || queued.Rows[0].Headline != "" {
		t.Fatalf("queued row: %+v", queued.Rows)
	}

	done := browserResult("warrior-fury", 1000)
	done.Lane, done.Request = simapi.LaneServer, req
	done.Weights = []simapi.StatWeight{
		{Stat: "melee_crit", Weight: 1}, {Stat: "agility", Weight: 0.874},
	}
	if err := h.store.Finish(t.Context(), "cccccccccccc", done); err != nil {
		t.Fatal(err)
	}
	finished, err := h.store.Mine(t.Context(), h.owner, 1, simapi.KindWeights)
	if err != nil {
		t.Fatal(err)
	}
	if len(finished.Rows) != 1 || finished.Rows[0].Headline != "Melee Crit 1.00 · Agility 0.87" {
		t.Fatalf("finished row: %+v", finished.Rows)
	}
}
```

Also change the existing `TestMyHistoryIsMineAndNewestFirst`'s call to `h.store.Mine(t.Context(), h.owner, 1, "")`.

- [ ] **Step 2: Run it and watch it fail**

Run: `go test -p 1 ./api/internal/sims/... -run 'TestMyHistoryCarries|TestAServerRunsHeadline' -v`
Expected: FAIL to build — `too many arguments in call to h.store.Mine`, `got.Kind undefined`.

- [ ] **Step 3: Widen `Row` and write the headline**

In `api/internal/sims/store.go`, replace `Row` with:

```go
// Row is one line of the caller's own sim history.
type Row struct {
	SimID string `json:"sim_id"`
	Spec  string `json:"spec"`
	// Kind is which tool produced this row: run, gear, talents, drops
	// or weights (contract 1.1).
	Kind string  `json:"kind"`
	DPS  float64 `json:"dps"`
	// Headline is the one line the list shows, composed by Headline at
	// write time and stored, because composing it on read would mean
	// detoasting the whole result blob for every row on the page.
	Headline      string    `json:"headline"`
	EngineVersion string    `json:"engine_version"`
	CreatedAt     time.Time `json:"created_at"`
	Title         string    `json:"title"`
}
```

In `Save`, add `headline` to the insert:

```go
	_, err = s.Pool.Exec(ctx,
		`insert into sims (id, user_id, spec, kind, headline, engine_version, lane,
		   dps_mean, dps_error, iterations, title, result, state)
		 values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		 on conflict (id) do nothing`,
		id, userID, res.Request.Spec, res.Request.Kind(), Headline(res), res.EngineVersion,
		res.Lane, res.DPS.Mean, res.DPS.Error, res.IterationsRun, t, body, StateDone)
```

In `Finish`, add it to the update:

```go
	_, err = s.Pool.Exec(ctx,
		`update sims set engine_version = $2, dps_mean = $3, dps_error = $4,
		   iterations = $5, result = $6, state = $7, headline = $8 where id = $1`,
		id, res.EngineVersion, res.DPS.Mean, res.DPS.Error, res.IterationsRun, body, state,
		Headline(res))
```

`Queue` is left alone: a queued run has produced nothing to headline, and the column's default is the empty string.

- [ ] **Step 4: Filter `Mine` by kind**

Replace `Mine` with:

```go
// Mine answers one page of a user's own sims, newest first. kind, when
// set, is one of simapi.Kinds and narrows the list to that tool; "" is
// every kind. The caller validates it — an unknown kind reaching here
// would simply return nothing, which reads as "you have none" rather
// than as the typo it is.
func (s *Store) Mine(ctx context.Context, userID int64, page int, kind string) (Page, error) {
	if page < 1 {
		page = 1
	}
	out := Page{Rows: []Row{}, Page: page, PerPage: PerPage}
	// One predicate, two queries: the count and the page must agree, and
	// a literal `$2 = '' or kind = $2` would make the planner ignore
	// sims_user_kind_idx for the filtered case, which is the case the
	// index exists for.
	where, args := "user_id = $1", []any{userID}
	if kind != "" {
		where += " and kind = $2"
		args = append(args, kind)
	}
	if err := s.Pool.QueryRow(ctx,
		`select count(*) from sims where `+where, args...).Scan(&out.Total); err != nil {
		return Page{}, fmt.Errorf("sims: count: %w", err)
	}
	rows, err := s.Pool.Query(ctx,
		`select id, spec, kind, dps_mean, headline, engine_version, created_at,
		        coalesce(title, '')
		 from sims where `+where+
			fmt.Sprintf(" order by created_at desc, id limit $%d offset $%d",
				len(args)+1, len(args)+2),
		append(args, PerPage, (page-1)*PerPage)...)
	if err != nil {
		return Page{}, fmt.Errorf("sims: list: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var r Row
		if err := rows.Scan(&r.SimID, &r.Spec, &r.Kind, &r.DPS, &r.Headline,
			&r.EngineVersion, &r.CreatedAt, &r.Title); err != nil {
			return Page{}, fmt.Errorf("sims: scan: %w", err)
		}
		out.Rows = append(out.Rows, r)
	}
	return out, rows.Err()
}
```

- [ ] **Step 5: Fix the one existing caller**

In `api/internal/sims/handler.go`'s `mine`, change the call to:

```go
	out, err := s.Store.Mine(r.Context(), auth.ActorFrom(r.Context()).UserID, page, "")
```

Task 4 replaces the `""` with the query parameter.

- [ ] **Step 6: Run the tests**

Run: `go test -p 1 ./api/internal/sims/...`
Expected: PASS. Re-run once on a `sims_pkey` collision.

- [ ] **Step 7: Commit**

```bash
git add api/internal/sims/store.go api/internal/sims/store_test.go api/internal/sims/handler.go
git commit -m "$(cat <<'EOF'
feat(api): history rows carry kind and headline, and filter by kind

Save and Finish compose the headline once and store it. Mine takes a
kind and builds one predicate for its count and its page, so the
filtered query uses sims_user_kind_idx.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

### Task 4: `GET /v1/sims?mine=1&kind=`

**Files:**
- Modify: `api/internal/sims/handler.go` (`mine`)
- Test: `api/internal/sims/handler_test.go`

**Interfaces:**
- Consumes: `Store.Mine(ctx, userID, page, kind)` (Task 3); `simapi.Kinds` (contract 1.1).
- Produces: the `kind=` query parameter on `GET /v1/sims`.

- [ ] **Step 1: Write the failing test**

Add to `api/internal/sims/handler_test.go`:

```go
func TestMyOwnSimsFilterByKind(t *testing.T) {
	h := newHarness(t)
	saveBrowserResult(h, "warrior-fury", 1204.4, "a plain run")

	gear := browserResult("mage-frost", 1100)
	gear.Request.Bulk = &simapi.BulkSpec{Mode: simapi.KindGear, Precision: simapi.PrecisionNormal}
	gear.Combos = []simapi.Combo{{
		Substitutions: []simapi.Substitution{
			{Kind: "item", ItemID: 19019, Name: "Thunderfury", Origin: "search"},
		},
		Delta: simapi.Estimate{Mean: 41.2},
	}}
	body, err := json.Marshal(struct {
		simapi.SimResult
		Title string `json:"title"`
	}{gear, "top gear"})
	if err != nil {
		t.Fatal(err)
	}
	h.data(h.json(http.MethodPost, "/v1/sims", string(body)), nil)

	var all Page
	h.data(h.do(http.MethodGet, "/v1/sims?mine=1", "", nil), &all)
	if all.Total != 2 {
		t.Fatalf("unfiltered total %d, want 2", all.Total)
	}

	var only Page
	h.data(h.do(http.MethodGet, "/v1/sims?mine=1&kind=gear", "", nil), &only)
	if only.Total != 1 || len(only.Rows) != 1 {
		t.Fatalf("filtered: %+v", only)
	}
	if only.Rows[0].Kind != simapi.KindGear ||
		only.Rows[0].Headline != "+41 DPS from Thunderfury" {
		t.Fatalf("row: %+v", only.Rows[0])
	}

	// An empty kind is "every kind", not "a kind called empty".
	var empty Page
	h.data(h.do(http.MethodGet, "/v1/sims?mine=1&kind=", "", nil), &empty)
	if empty.Total != 2 {
		t.Fatalf("kind= total %d, want 2", empty.Total)
	}
}

func TestAnUnknownKindIsRefusedRatherThanAnsweredEmpty(t *testing.T) {
	h := newHarness(t)
	res := h.do(http.MethodGet, "/v1/sims?mine=1&kind=topgear", "", nil)
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("status %d, want 400", res.StatusCode)
	}
	if code := h.errorCode(res); code != "invalid" {
		t.Fatalf("code %q, want invalid", code)
	}
}
```

- [ ] **Step 2: Run it and watch it fail**

Run: `go test -p 1 ./api/internal/sims/... -run 'TestMyOwnSimsFilterByKind|TestAnUnknownKind' -v`
Expected: FAIL — the filtered page reports 2, and the unknown kind is answered 200.

- [ ] **Step 3: Read and validate the parameter**

In `api/internal/sims/handler.go`'s `mine`, insert after the `page` block and before the `Store.Mine` call:

```go
	// An unknown kind is refused rather than answered with an empty
	// list: "you have no topgear sims" is a true sentence about a word
	// that does not exist, and a typo in a link would read as a real
	// (empty) answer forever.
	kind := r.URL.Query().Get("kind")
	if kind != "" && !slices.Contains(simapi.Kinds, kind) {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid",
			"kind must be one of "+strings.Join(simapi.Kinds, ", "),
			map[string]string{"kind": strings.Join(simapi.Kinds, "|")})
		return
	}
	out, err := s.Store.Mine(r.Context(), auth.ActorFrom(r.Context()).UserID, page, kind)
```

and delete the old `out, err := s.Store.Mine(...)` line. Add `"slices"` and `"strings"` to the import block.

If `simapi.Kinds` does not exist yet, the `sim` lane has not merged contract 1.1's closed set; wait for it rather than declaring a second copy here.

- [ ] **Step 4: Run the tests**

Run: `go test -p 1 ./api/internal/sims/...`
Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git add api/internal/sims/handler.go api/internal/sims/handler_test.go
git commit -m "$(cat <<'EOF'
test(api): the history list filters by kind and refuses an unknown one

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

### Task 5: Counting a request before it is queued — `cap_exceeded`

**Waits on:** the `sim` lane's `simapi.PlanSummary` and `SimRequest.ValidateLane` (contract 1.8 as amended, and A1). The `Planner` interface and its test double are declared here, so nothing else in this task needs a binary.

**Files:**
- Modify: `api/internal/sims/simdep.go` (the `Planner` and `Engine` interfaces)
- Modify: `api/internal/sims/handler.go` (`Service.Planner`, `Mount`)
- Modify: `api/internal/sims/run.go`
- Test: `api/internal/sims/run_test.go`, `api/internal/sims/harness_test.go`, `api/internal/sims/handler_test.go`, `api/internal/server/sims_test.go`

**Interfaces:**
- Consumes: `simapi.PlanSummary{Kind, Combinations, Cap, IterationsTotal}`; `simapi.Caps` (A2: server 5,000); `simapi.SimRequest.ValidateLane(lane string) error` (A1).
- Produces:
  - `type Planner interface { Plan(ctx context.Context, req simapi.SimRequest) (simapi.PlanSummary, error) }`
  - `type Engine interface { runner.Runner; Planner }`
  - `Service.Planner Planner`
  - `func (s *Service) checkSize(w http.ResponseWriter, r *http.Request, req simapi.SimRequest) bool`
  - `const planTimeout = 30 * time.Second`

- [ ] **Step 1: Write the failing test**

Add to `api/internal/sims/harness_test.go`, next to `fakePremium`:

```go
// fakePlanner answers the plan-only count without a binary. The real one
// invokes `forever-sim -plan` (contract 10.2); nothing about the
// handler's decision needs a subprocess to be tested.
type fakePlanner struct {
	summary simapi.PlanSummary
	err     error
	asked   []simapi.SimRequest
}

func (f *fakePlanner) Plan(_ context.Context, req simapi.SimRequest) (simapi.PlanSummary, error) {
	f.asked = append(f.asked, req)
	if f.err != nil {
		return simapi.PlanSummary{}, f.err
	}
	return f.summary, nil
}
```

Add a `planner *fakePlanner` field to `harness`, and in `newHarness` replace the two lines that build the service with:

```go
	h.jobs, h.premium = &jobs.Fake{}, &fakePremium{}
	// A plan that fits: one combination, well inside both bounds. A test
	// that cares sets its own.
	h.planner = &fakePlanner{summary: simapi.PlanSummary{
		Kind: simapi.KindGear, Combinations: 1,
		Cap: simapi.Caps[simapi.LaneServer], IterationsTotal: 4000,
	}}
	h.service = &Service{
		Store: h.store, Accounts: h.premium, Jobs: h.jobs, Planner: h.planner,
		EngineVersion: testEngine, Log: quiet,
	}
```

Add to `api/internal/sims/run_test.go`:

```go
// bulkBody is a well-formed Top Gear submit: the same envelope with a
// bulk block on it. Iterations is the precision's final-stage count,
// which is what Validate expects of a bulk request (contract A3).
func bulkBody(t *testing.T) string {
	t.Helper()
	b, err := json.Marshal(simapi.SimRequest{
		EngineVersion: "an old one the page was holding", Spec: "warrior-fury",
		Iterations: defaultIterations,
		Source:     simapi.CharacterSource{Kind: simapi.SourceAddon, Ref: "us/normal/baelgrim"},
		Character:  aCharacter("warrior", "orc"),
		Bulk: &simapi.BulkSpec{
			Mode: simapi.KindGear, Precision: simapi.PrecisionNormal,
			// A client-chosen cap the server must overwrite.
			Cap: 1_000_000,
			Candidates: []simapi.Candidate{
				{Slot: "main_hand", ItemID: 19019, Origin: "bag"},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

// refusal reads a failing envelope's code and fields together.
func refusal(t *testing.T, res *http.Response) (string, map[string]string) {
	t.Helper()
	defer res.Body.Close()
	var env struct {
		Error struct {
			Code   string            `json:"code"`
			Fields map[string]string `json:"fields"`
		} `json:"error"`
	}
	if err := json.NewDecoder(res.Body).Decode(&env); err != nil {
		t.Fatal(err)
	}
	return env.Error.Code, env.Error.Fields
}

func TestABulkRunPastTheLanesCapIsRefusedWithBothNumbers(t *testing.T) {
	h := newHarness(t)
	h.premium.premium = true
	h.planner.summary = simapi.PlanSummary{
		Kind: simapi.KindGear, Combinations: 31200,
		Cap: simapi.Caps[simapi.LaneServer], IterationsTotal: 100,
	}

	res := h.json(http.MethodPost, "/v1/sims/run", bulkBody(t))
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("status %d, want 400", res.StatusCode)
	}
	code, fields := refusal(t, res)
	if code != "cap_exceeded" {
		t.Fatalf("code %q, want cap_exceeded", code)
	}
	// Decimal strings: httpx.ErrorBody.Fields is map[string]string
	// (contract 10.6).
	if fields["combinations"] != "31200" ||
		fields["cap"] != strconv.Itoa(simapi.Caps[simapi.LaneServer]) {
		t.Fatalf("fields %+v", fields)
	}
	if ran := h.jobs.Ran(); len(ran) != 0 {
		t.Fatalf("a refused run was dispatched anyway: %v", ran)
	}
}

func TestTheServerSetsTheCapItself(t *testing.T) {
	h := newHarness(t)
	h.premium.premium = true
	if res := h.json(http.MethodPost, "/v1/sims/run", bulkBody(t)); res.StatusCode != http.StatusAccepted {
		t.Fatalf("status %d, want 202", res.StatusCode)
	}
	if len(h.planner.asked) != 1 {
		t.Fatalf("%d plans", len(h.planner.asked))
	}
	if got := h.planner.asked[0].Bulk.Cap; got != simapi.Caps[simapi.LaneServer] {
		t.Errorf("cap %d, want the server lane's %d; a client may not raise its own bound",
			got, simapi.Caps[simapi.LaneServer])
	}
}

func TestAPlainRunIsNeverPlanned(t *testing.T) {
	h := newHarness(t)
	h.premium.premium = true
	if res := h.json(http.MethodPost, "/v1/sims/run", runBody(t)); res.StatusCode != http.StatusAccepted {
		t.Fatalf("status %d, want 202", res.StatusCode)
	}
	if len(h.planner.asked) != 0 {
		t.Errorf("a plain run has nothing to expand, but the planner was asked: %+v", h.planner.asked)
	}
}
```

Add `"strconv"` to `run_test.go`'s imports.

- [ ] **Step 2: Run it and watch it fail**

Run: `go test -p 1 ./api/internal/sims/... -run 'TestABulkRunPastTheLanesCap|TestTheServerSetsTheCap|TestAPlainRunIsNeverPlanned' -v`
Expected: FAIL to build — `unknown field Planner in struct literal of type Service`.

- [ ] **Step 3: Declare the interfaces**

Add to `api/internal/sims/simdep.go`, after the `roleDPS` constant:

```go
// Planner counts a bulk or weights request without running it. The real
// implementation is sim/runner's, which invokes `forever-sim -plan`
// (contract 10.2): sim/bulk cannot be imported here — it reads
// sim/internal/simdb, which imports the engine, and sim/internal is
// closed to other modules besides — so the count crosses the boundary as
// a subprocess's JSON, exactly the way a run does.
//
// The interface is declared here, by the consumer, so this package's
// tests can exercise the decision without a binary.
type Planner interface {
	Plan(ctx context.Context, req simapi.SimRequest) (simapi.PlanSummary, error)
}

// Engine is a Runner that can also plan: what cmd/api hands the service
// and the jobs alike, so one binary-presence check serves both.
type Engine interface {
	runner.Runner
	Planner
}
```

Add `"context"` and `"github.com/jhunthrop/foreversixty/sim/runner"` to that file's imports, and add `sim/runner` to the list of importable `sim` packages in the package doc at the top of `simdep.go`.

- [ ] **Step 4: Carry a planner on the service and require it to mount the route**

In `api/internal/sims/handler.go`, add to `Service`, under `Jobs`:

```go
	// Planner counts a bulk or weights request before it is queued. Nil
	// means the premium lane cannot size one, and the run route is not
	// mounted: accepting a request we cannot bound would be worse than
	// not offering the route.
	Planner Planner
```

and change `Mount`'s condition:

```go
	if s.Jobs != nil && s.Accounts != nil && s.Planner != nil {
		mux.HandleFunc("POST /v1/sims/run", auth.RequireSession(s.run))
	}
```

Update `Mount`'s doc comment's last sentence to: "without the job runner, the accounts store or the planner there is nothing behind it, and a 404 is a truer answer than a 500."

- [ ] **Step 5: Force the cap, validate per lane, and check the count**

In `api/internal/sims/run.go`, replace the block from `req.Encounter = withEncounterDefaults(req.Encounter)` through the `Validate` error branch with:

```go
	req.Encounter = withEncounterDefaults(req.Encounter)
	if req.Bulk != nil {
		// The lane's cap is the server's to set. A request echoes the
		// cap that bounded it (contract 1.3) so a saved request says
		// what it ran under; it is not a control the client holds.
		req.Bulk.Cap = simapi.Caps[simapi.LaneServer]
	}
	// ValidateLane, not Validate: the envelope's plain Validate checks
	// the largest lane's cap, and this is the server lane (contract A1).
	if err := req.ValidateLane(simapi.LaneServer); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", err.Error(), nil)
		return
	}
	if req.Kind() != simapi.KindRun && s.checkSize(w, r, req) {
		return
	}
```

Then add at the bottom of `run.go`:

```go
// planTimeout bounds the plan-only subprocess. Expansion loads the item
// database and walks the candidates; it runs no iterations, so thirty
// seconds bounds something pathological rather than budgeting the work.
const planTimeout = 30 * time.Second

// checkSize counts what the request would expand to, without running any
// of it, and refuses the two ways it can be too big. It reports whether
// it has already written a response.
func (s *Service) checkSize(w http.ResponseWriter, r *http.Request, req simapi.SimRequest) bool {
	ctx, cancel := context.WithTimeout(r.Context(), planTimeout)
	defer cancel()
	plan, err := s.Planner.Plan(ctx, req)
	if err != nil {
		s.fail(w, r, "plan", err, "could not size that run just now")
		return true
	}
	if plan.Cap > 0 && plan.Combinations > plan.Cap {
		// Both numbers, because the page says how far over it is and by
		// how much to trim. error.Fields is map[string]string, so they
		// go over as decimal strings (contract 10.6).
		httpx.WriteError(w, r, http.StatusBadRequest, "cap_exceeded",
			fmt.Sprintf("that is %s combinations; a run on our servers is at most %s",
				withThousands(int64(plan.Combinations)), withThousands(int64(plan.Cap))),
			map[string]string{
				"cap":          strconv.Itoa(plan.Cap),
				"combinations": strconv.Itoa(plan.Combinations),
			})
		return true
	}
	return false
}
```

Add `"fmt"` and `"strconv"` to `run.go`'s imports (`context`, `time` and `httpx` are already there).

- [ ] **Step 6: Fix the two mount tests**

Add to `handler_test.go`:

```go
func TestTheRunRouteNeedsAPlannerToo(t *testing.T) {
	mux := http.NewServeMux()
	Mount(mux, &Service{Jobs: &jobs.Fake{}, Accounts: &fakePremium{}})
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, httptest.NewRequest(http.MethodPost, "/v1/sims/run", strings.NewReader("{}")))
	if w.Code != http.StatusNotFound {
		t.Errorf("status %d, want 404: a lane that cannot size a bulk request does not offer one", w.Code)
	}
}
```

with `"github.com/jhunthrop/foreversixty/api/internal/jobs"` imported. Then in `api/internal/server/sims_test.go`, the `withRun` service literal needs a planner; add above it:

```go
// planNothing satisfies sims.Planner for the mount check, which never
// plans anything.
type planNothing struct{}

func (planNothing) Plan(context.Context, simapi.SimRequest) (simapi.PlanSummary, error) {
	return simapi.PlanSummary{}, nil
}
```

and set `Planner: planNothing{}` on that literal, importing `"context"` and `simapi "github.com/jhunthrop/foreversixty/sim/api"`.

- [ ] **Step 7: Run the tests**

Run: `go test -p 1 ./api/internal/sims/... ./api/internal/server/...`
Expected: PASS.

- [ ] **Step 8: Commit**

```bash
git add api/internal/sims/simdep.go api/internal/sims/handler.go api/internal/sims/run.go \
        api/internal/sims/run_test.go api/internal/sims/harness_test.go \
        api/internal/sims/handler_test.go api/internal/server/sims_test.go
git commit -m "$(cat <<'EOF'
feat(api): the premium lane sizes a bulk request before it queues one

The server overwrites bulk.cap with its own lane's, validates with
ValidateLane(server), then counts the expansion through a plan-only call
— no iterations — and answers 400 cap_exceeded with the cap and the
count as decimal strings. sim/bulk cannot be imported here, so the count
crosses as JSON from the binary; the interface is declared by the
consumer so the handler is testable without one.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

### Task 6: The `too_large` estimate

**Waits on:** the `sim` lane publishing `measure.NativeIterationsPerCPUSecond` in `sim/measure` (contract A2). `sim/measure` is engine-free, so the api module may import it.

**Files:**
- Modify: `api/internal/sims/job.go` (`BulkBudget`, `timeoutFor`)
- Modify: `api/internal/sims/simdep.go` (the package doc, `simJobCPUs`, `estimateSec`)
- Modify: `api/internal/sims/run.go` (`checkSize`)
- Test: `api/internal/sims/run_test.go`, `api/internal/sims/simdep_test.go`

**Interfaces:**
- Consumes: `simapi.PlanSummary.IterationsTotal`; `measure.NativeIterationsPerCPUSecond`.
- Produces:
  - `const BulkBudget = 840 * time.Second`
  - `func timeoutFor(req simapi.SimRequest) time.Duration`
  - `func estimateSec(iterations int) int`
  - `const nativeRate = measure.NativeIterationsPerCPUSecond * simJobCPUs`

- [ ] **Step 1: Write the failing test**

Add to `api/internal/sims/simdep_test.go`:

```go
func TestTheEstimateIsIterationsOverTheJobsMeasuredRate(t *testing.T) {
	// Stated in multiples of the rate, not in literals: the rate is
	// sim/measure's benchmark figure and moves when the benchmark does.
	for _, c := range []struct {
		iterations int
		want       int
	}{
		{0, 0},
		{1, 1}, // rounded up: a run is never estimated at no time at all
		{nativeRate, 1},
		{nativeRate * 60, 60},
		{nativeRate*4000 + 1, 4001},
	} {
		if got := estimateSec(c.iterations); got != c.want {
			t.Errorf("estimateSec(%d) = %d, want %d", c.iterations, got, c.want)
		}
	}
}

func TestTheBulkBudgetFitsInsideTheCloudRunTaskTimeout(t *testing.T) {
	// api/README.md creates sim-run with --task-timeout 15m, and
	// contract A2 fixes the budget at 840 seconds. The in-process bound
	// has to leave room for start-up and the two writes at the end, or
	// the platform kills the job mid-write and the row never leaves
	// "running".
	const taskTimeout = 15 * time.Minute
	if BulkBudget != 840*time.Second {
		t.Fatalf("BulkBudget %s, want the contract's 840s", BulkBudget)
	}
	if BulkBudget >= taskTimeout {
		t.Fatalf("BulkBudget %s, want less than the job's %s", BulkBudget, taskTimeout)
	}
	if RunTimeout > BulkBudget {
		t.Fatalf("a plain run (%s) may not outlast a bulk one (%s)", RunTimeout, BulkBudget)
	}
}

func TestTheServerCapAndTheBudgetAgree(t *testing.T) {
	// Contract A2 lowered the server cap to 5,000 so that a full-cap run
	// can actually finish: the largest expansion the lane accepts must
	// not be one the budget always refuses. A fast ladder over the cap is
	// cap×100 + cap/4×1,000 + 11×3,000 iterations.
	cap := simapi.Caps[simapi.LaneServer]
	fastLadder := cap*100 + (cap/4)*1000 + 11*3000
	if est := estimateSec(fastLadder); est > int(BulkBudget.Seconds()) {
		t.Fatalf("a full-cap fast run estimates %ds against a %ds budget: "+
			"the cap and the budget disagree", est, int(BulkBudget.Seconds()))
	}
}
```

Add `"time"` and `simapi "github.com/jhunthrop/foreversixty/sim/api"` to `simdep_test.go`'s imports if they are not already there.

Add to `api/internal/sims/run_test.go`:

```go
func TestABulkRunPastTheBudgetIsRefusedWithItsEstimate(t *testing.T) {
	h := newHarness(t)
	h.premium.premium = true
	// Exactly four thousand seconds of engine time, whatever the
	// benchmark's current figure is.
	h.planner.summary = simapi.PlanSummary{
		Kind: simapi.KindGear, Combinations: 4000,
		Cap: simapi.Caps[simapi.LaneServer], IterationsTotal: nativeRate * 4000,
	}

	res := h.json(http.MethodPost, "/v1/sims/run", bulkBody(t))
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("status %d, want 400", res.StatusCode)
	}
	code, fields := refusal(t, res)
	if code != "too_large" {
		t.Fatalf("code %q, want too_large", code)
	}
	if fields["estimate_sec"] != "4000" ||
		fields["budget_sec"] != strconv.Itoa(int(BulkBudget.Seconds())) {
		t.Fatalf("fields %+v", fields)
	}
	if ran := h.jobs.Ran(); len(ran) != 0 {
		t.Fatalf("a refused run was dispatched anyway: %v", ran)
	}
}

func TestAWeightsRunHasNoCapButStillHasABudget(t *testing.T) {
	h := newHarness(t)
	h.premium.premium = true
	// Cap 0: weights expand to no combinations, so the cap test must not
	// fire on a zero and refuse every one of them.
	h.planner.summary = simapi.PlanSummary{
		Kind: simapi.KindWeights, Combinations: 0, Cap: 0, IterationsTotal: 60_000,
	}
	b, err := json.Marshal(simapi.SimRequest{
		EngineVersion: testEngine, Spec: "warrior-fury", Iterations: defaultIterations,
		Source:    simapi.CharacterSource{Kind: simapi.SourceAddon, Ref: "us/normal/baelgrim"},
		Character: aCharacter("warrior", "orc"),
		Weights: &simapi.WeightsSpec{
			Stats: []string{"strength", "melee_crit"}, Reference: "melee_crit",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if res := h.json(http.MethodPost, "/v1/sims/run", string(b)); res.StatusCode != http.StatusAccepted {
		t.Fatalf("status %d, want 202", res.StatusCode)
	}
	if len(h.planner.asked) != 1 {
		t.Fatalf("a weights run must still be sized: %d plans", len(h.planner.asked))
	}
}
```

- [ ] **Step 2: Run it and watch it fail**

Run: `go test -p 1 ./api/internal/sims/... -run 'TestTheEstimate|TestTheBulkBudget|TestTheServerCapAndTheBudget|TestABulkRunPastTheBudget|TestAWeightsRun' -v`
Expected: FAIL to build — `undefined: estimateSec`, `undefined: nativeRate`, `undefined: BulkBudget`.

- [ ] **Step 3: Write the rate and the estimate**

Add to `api/internal/sims/simdep.go`, after the iteration constants:

```go
// The shape of the job the estimate is made against. The engine's own
// rate is measure.NativeIterationsPerCPUSecond, which sim/measure
// publishes from its benchmark (contract A2) — it is not restated here,
// because a second copy of a benchmark figure is a copy that goes stale.
// simJobCPUs is what api/README.md creates the sim-run job with
// (`--cpu 4`); the two have to change together, and the README says so.
const (
	simJobCPUs = 4
	nativeRate = measure.NativeIterationsPerCPUSecond * simJobCPUs
)

// estimateSec is how long the sim-run job needs to complete iterations
// iterations, rounded up: a run that would take a fraction of a second
// still takes some. Multiplied out, nativeRate × BulkBudget is the
// iteration ceiling contract A2 names.
func estimateSec(iterations int) int {
	if iterations <= 0 {
		return 0
	}
	return (iterations + nativeRate - 1) / nativeRate
}
```

Add `"github.com/jhunthrop/foreversixty/sim/measure"` to `simdep.go`'s imports, and add `sim/measure` to the list of importable `sim` packages in the package doc at the top of the file — it is engine-free, which is why it may be imported here at all.

- [ ] **Step 4: Give a bulk run its own bound**

In `api/internal/sims/job.go`, add under `RunTimeout`:

```go
// BulkBudget bounds one bulk or weights run: every stage of the ladder,
// not one sim. Contract A2 fixes it at 840 seconds. api/README.md creates
// the sim-run job with --task-timeout 15m, and a job killed by the
// platform dies between the bucket write and the row write, leaving the
// page polling "running" forever — so the in-process bound stops a minute
// short of it and fails the row on its way out. It is also the budget the
// submit-time estimate is refused against, so a run that is accepted is
// a run that can finish.
const BulkBudget = 840 * time.Second

// timeoutFor is the bound one request's run gets. A plain run keeps the
// tighter one: ten minutes for a single sim is already an outlier worth
// failing.
func timeoutFor(req simapi.SimRequest) time.Duration {
	if req.Kind() == simapi.KindRun {
		return RunTimeout
	}
	return BulkBudget
}
```

and in `Run`, replace `runCtx, cancel := context.WithTimeout(ctx, RunTimeout)` with:

```go
	runCtx, cancel := context.WithTimeout(ctx, timeoutFor(req))
```

- [ ] **Step 5: Refuse the oversized request**

In `api/internal/sims/run.go`'s `checkSize`, after the cap block and before `return false`:

```go
	budget := int(BulkBudget.Seconds())
	if est := estimateSec(plan.IterationsTotal); est > budget {
		// Refused at submit rather than started and killed: a job the
		// platform stops leaves a row that never reaches a terminal
		// state, and the member has waited a quarter of an hour to find
		// out. The estimate is what they trim against.
		httpx.WriteError(w, r, http.StatusBadRequest, "too_large",
			fmt.Sprintf("that run is about %d seconds of engine time; a run on our servers stops at %d",
				est, budget),
			map[string]string{
				"estimate_sec": strconv.Itoa(est),
				"budget_sec":   strconv.Itoa(budget),
			})
		return true
	}
```

- [ ] **Step 6: Run the tests**

Run: `go test -p 1 ./api/internal/sims/...`
Expected: PASS. If `TestTheServerCapAndTheBudgetAgree` fails, stop: the cap, the budget and the benchmark no longer agree and that is a contract question, not a code fix.

- [ ] **Step 7: Commit**

```bash
git add api/internal/sims/simdep.go api/internal/sims/simdep_test.go \
        api/internal/sims/job.go api/internal/sims/run.go api/internal/sims/run_test.go
git commit -m "$(cat <<'EOF'
feat(api): a run past the job's budget is refused at submit, with the estimate

The estimate is the plan's total iterations over sim/measure's published
native rate times the job's four CPUs, against contract A2's 840-second
budget — inside Cloud Run's 15-minute task timeout. A test asserts the
5,000-combination server cap and that budget still agree, so the lane
never accepts a full-cap run it cannot finish.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

### Task 7: Stage progress through the job

**Waits on:** the `sim` lane widening `runner.Progress` with `Stage`, `CombosDone` and `CombosTotal` (contract A11). A11 makes that widening additive, so nothing here replaces a type — but **read `sim/runner/runner.go` before Step 4** and write the callback against whatever shape landed. This package's own storage vocabulary is `Tick`, declared here, so the store and its tests do not depend on the runner's callback shape at all.

**Files:**
- Modify: `api/internal/sims/store.go` (`Tick`, `Progress`, `Advance`)
- Modify: `api/internal/sims/job.go` (the callback)
- Test: `api/internal/sims/job_test.go`, `api/internal/sims/store_test.go`, `api/internal/sims/validate_test.go`

**Interfaces:**
- Consumes: `runner.Progress` with the three added fields.
- Produces:
  - `type Tick struct { IterationsDone int; Mean float64; Stage int; CombosDone int; CombosTotal int }`
  - `func (s *Store) Advance(ctx context.Context, id string, t Tick) error`
  - `Progress` gains `Stage int \`json:"stage"\``, `CombosDone int \`json:"combos_done"\``, `CombosTotal int \`json:"combos_total"\``

- [ ] **Step 1: Write the failing test**

Add to `api/internal/sims/job_test.go`:

```go
// bulkRunner answers the way forever-sim does for a bulk request: it
// streams a tick per stage, then returns a ranked result. The binary
// detects the kind and runs sim/bulk's plan-rank loop itself, so the job
// hands it the request whole and reads what comes back — which is exactly
// what this stands in for.
//
// The onProgress literal below is written against runner.Progress as the
// sim lane widened it (contract A11); if it does not compile, read
// sim/runner/runner.go and match the shape that landed — the three
// numbers it must carry are the same either way.
type bulkRunner struct{ got simapi.SimRequest }

func (b *bulkRunner) Run(_ context.Context, req simapi.SimRequest,
	onProgress runner.Progress) (simapi.SimResult, error) {
	b.got = req
	if onProgress != nil {
		onProgress(runner.Progress{
			IterationsRun: 1000, DPS: simapi.Estimate{Mean: 1000},
			Stage: 1, CombosDone: 40, CombosTotal: 96,
		})
		onProgress(runner.Progress{
			IterationsRun: 4000, DPS: simapi.Estimate{Mean: 1040},
			Stage: 2, CombosDone: 10, CombosTotal: 10,
		})
	}
	return simapi.SimResult{
		EngineVersion: req.EngineVersion, Request: req, Lane: simapi.LaneServer,
		DPS:           simapi.Estimate{Mean: 1042.5, Error: 1.6},
		IterationsRun: 4000, DurationMS: 90_000,
		Equipped:      &simapi.Estimate{Mean: 1001.3},
		Stages:        []simapi.Stage{{Iterations: 1000, Combos: 96}, {Iterations: 3000, Combos: 10}},
		Combos: []simapi.Combo{{
			Substitutions: []simapi.Substitution{
				{Kind: "item", Slot: "main_hand", ItemID: 17182,
					Name: "Vis'kag the Bloodletter", Origin: "bag"},
			},
			DPS:   simapi.Estimate{Mean: 1042.5},
			Delta: simapi.Estimate{Mean: 41.2, Error: 2.2},
		}},
	}, nil
}

// queuedBulk puts a Top Gear run in the queued state.
func (h *harness) queuedBulk(t *testing.T, id string) simapi.SimRequest {
	t.Helper()
	req := browserResult("warrior-fury", 0).Request
	req.Bulk = &simapi.BulkSpec{
		Mode: simapi.KindGear, Precision: simapi.PrecisionNormal,
		Cap:        simapi.Caps[simapi.LaneServer],
		Candidates: []simapi.Candidate{{Slot: "main_hand", ItemID: 17182, Origin: "bag"}},
	}
	if err := h.store.Queue(t.Context(), id, h.owner, req); err != nil {
		t.Fatal(err)
	}
	return req
}

func TestABulkJobStreamsItsStagesAndStoresTheRanking(t *testing.T) {
	h := newHarness(t)
	want := h.queuedBulk(t, "ffffffffffff")
	engine := &bulkRunner{}

	if err := Run(t.Context(), h.jobDeps(engine), "ffffffffffff"); err != nil {
		t.Fatal(err)
	}

	// The whole request reached the binary, bulk block and all: the
	// binary detects the kind and runs the plan-rank loop itself.
	if engine.got.Bulk == nil || engine.got.Bulk.Mode != want.Bulk.Mode ||
		len(engine.got.Bulk.Candidates) != 1 {
		t.Fatalf("the bulk block did not reach the engine: %+v", engine.got.Bulk)
	}

	p, err := h.store.Progress(t.Context(), "ffffffffffff")
	if err != nil {
		t.Fatal(err)
	}
	if p.State != StateDone {
		t.Fatalf("state %q", p.State)
	}
	// The last tick is what the row remembers.
	if p.Stage != 2 || p.CombosDone != 10 || p.CombosTotal != 10 {
		t.Errorf("progress %+v, want the final stage's counts", p)
	}

	stored, err := h.store.Get(t.Context(), "ffffffffffff")
	if err != nil {
		t.Fatal(err)
	}
	if len(stored.Combos) != 1 || stored.Equipped == nil || len(stored.Stages) != 2 {
		t.Fatalf("the ranking did not survive storage: %+v", stored)
	}

	// And the history row says what it found.
	page, err := h.store.Mine(t.Context(), h.owner, 1, simapi.KindGear)
	if err != nil {
		t.Fatal(err)
	}
	if len(page.Rows) != 1 || page.Rows[0].Headline != "+41 DPS from Vis'kag the Bloodletter" {
		t.Fatalf("history row: %+v", page.Rows)
	}
}

func TestAPlainRunReportsNoStages(t *testing.T) {
	h := newHarness(t)
	h.queued(t, "gggggggggggg", "warrior-fury")
	if err := Run(t.Context(), h.jobDeps(&runner.Fixture{}), "gggggggggggg"); err != nil {
		t.Fatal(err)
	}
	p, err := h.store.Progress(t.Context(), "gggggggggggg")
	if err != nil {
		t.Fatal(err)
	}
	if p.Stage != 0 || p.CombosDone != 0 || p.CombosTotal != 0 {
		t.Errorf("a plain run reported stages: %+v", p)
	}
}
```

Add to `api/internal/sims/validate_test.go`:

```go
func TestTheValidationJobStillRunsPlainSims(t *testing.T) {
	// The nightly pass measures rotation fidelity, not the optimiser
	// (design section 11). Every request it builds must stay a plain
	// run: a bulk or weights request there would sim fifty parses dozens
	// of times each and blow the job's budget.
	h := newHarness(t)
	engine := &runner.Fixture{}
	d := h.validateDeps(engine)
	_ = Validate(t.Context(), d, []string{"warrior-fury"}, "raids-1", testEngine)
	asked := engine.Asked()
	if len(asked) == 0 {
		t.Fatal("the validation job ran nothing; the assertion below would be vacuous")
	}
	for _, req := range asked {
		if req.Kind() != simapi.KindRun {
			t.Fatalf("the validation job asked for a %s run: %+v", req.Kind(), req)
		}
		if req.Bulk != nil || req.Weights != nil {
			t.Fatalf("the validation job built a bulk request: %+v", req)
		}
	}
}
```

Read `validate_test.go` first: it already builds a `ValidateDeps` with its own fakes for `Top`, `Build` and `Scores`. Extract that construction into a `func (h *harness) validateDeps(engine runner.Runner) ValidateDeps` helper in that file and have both the existing tests and this one call it, rather than adding a second set of fakes.

- [ ] **Step 2: Run it and watch it fail**

Run: `go test -p 1 ./api/internal/sims/... -run 'TestABulkJobStreams|TestAPlainRunReportsNoStages|TestTheValidationJobStill' -v`
Expected: FAIL — `p.Stage undefined`.

- [ ] **Step 3: Widen the stored progress**

In `api/internal/sims/store.go`, replace `Progress` with:

```go
// Progress is what a page following a server run is shown. The three
// stage fields are zero for a plain run, which has no stages (contract
// 2); a bulk run's carry the last tick the job wrote.
type Progress struct {
	State          string   `json:"state"`
	IterationsDone int      `json:"iterations_done"`
	DPS            *float64 `json:"dps,omitempty"`
	Stage          int      `json:"stage"`
	CombosDone     int      `json:"combos_done"`
	CombosTotal    int      `json:"combos_total"`
}

// Tick is one progress report on its way into the row. It is this
// package's own vocabulary rather than the runner's callback type, so
// the store and its tests do not move when the engine's progress stream
// gains a field.
type Tick struct {
	IterationsDone int
	Mean           float64
	Stage          int
	CombosDone     int
	CombosTotal    int
}
```

and its reader:

```go
func (s *Store) Progress(ctx context.Context, id string) (Progress, error) {
	var (
		p    Progress
		mean float64
		iter int
	)
	err := s.Pool.QueryRow(ctx,
		`select state, dps_mean, iterations, stage, combos_done, combos_total
		 from sims where id = $1`, id).Scan(&p.State, &mean, &iter,
		&p.Stage, &p.CombosDone, &p.CombosTotal)
	if isNoRows(err) {
		return Progress{}, ErrNotFound
	}
	if err != nil {
		return Progress{}, fmt.Errorf("sims: progress %s: %w", id, err)
	}
	if p.State == StateRunning || p.State == StateDone {
		p.IterationsDone, p.DPS = iter, &mean
	}
	return p, nil
}
```

Replace `Advance` with:

```go
// Advance records a running job's partial estimate, so the page's DPS
// figure refines and its stage line moves while the job is still going.
// A row already in a terminal state is left alone: a progress tick that
// arrives after the result, or after a failure, would otherwise walk the
// run backwards or revive an errored one.
func (s *Store) Advance(ctx context.Context, id string, t Tick) error {
	_, err := s.Pool.Exec(ctx,
		`update sims set state = $2, iterations = $3, dps_mean = $4,
		   stage = $5, combos_done = $6, combos_total = $7
		 where id = $1 and state in ($8, $9)`,
		id, StateRunning, t.IterationsDone, t.Mean,
		t.Stage, t.CombosDone, t.CombosTotal, StateQueued, StateRunning)
	if err != nil {
		return fmt.Errorf("sims: advance %s: %w", id, err)
	}
	return nil
}
```

- [ ] **Step 4: Adapt the callback in the job**

Read `sim/runner/runner.go` first, then in `api/internal/sims/job.go` replace the `Engine.Run` callback with the adaptation — one `Tick` built from whatever the runner hands over:

```go
	res, err := d.Engine.Run(runCtx, req, func(p runner.Progress) {
		// A progress write that fails is logged and the run carries on:
		// the figure on the page is a courtesy, the result is not.
		if err := d.Store.Advance(ctx, simID, Tick{
			IterationsDone: p.IterationsRun, Mean: p.DPS.Mean,
			Stage: p.Stage, CombosDone: p.CombosDone, CombosTotal: p.CombosTotal,
		}); err != nil {
			d.logger().Error("sims", "op", "progress", "sim", simID, "err", err)
		}
	})
```

If the `sim` lane kept a two-argument func and added the stage numbers as further parameters, the literal's signature changes and its body does not: build the same `Tick` from the parameters it gives you.

- [ ] **Step 5: Fix the three existing `Advance` callers in tests**

`TestAServerRunWalksQueuedThenRunningThenDone`, `TestAdvanceCannotReviveAFailedRun` and `TestAdvanceCannotReopenAFinishedRun` in `store_test.go` each call `h.store.Advance(ctx, id, 1500, 1000)`. Rewrite each call as, keeping that test's own id and expectations:

```go
	if err := h.store.Advance(t.Context(), "dddddddddddd", Tick{
		IterationsDone: 1500, Mean: 1000,
	}); err != nil {
```

- [ ] **Step 6: Run the tests**

Run: `go test -p 1 ./api/internal/sims/...`
Expected: PASS.

- [ ] **Step 7: Commit**

```bash
git add api/internal/sims/store.go api/internal/sims/store_test.go \
        api/internal/sims/job.go api/internal/sims/job_test.go \
        api/internal/sims/validate_test.go
git commit -m "$(cat <<'EOF'
feat(api): the sim-run job streams stage, combos_done and combos_total

The binary detects the kind and runs sim/bulk's plan-rank loop itself,
so the job hands it the request whole and writes what each tick carries.
Store.Advance takes this package's own Tick rather than the runner's
callback type, so the row's vocabulary does not move when the engine's
progress stream does. Asserts the nightly validation job still builds
plain runs only.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

### Task 8: `GET /v1/specs` carries the reference stat

**Waits on:** the data lane adding `reference_stat` to `data/curated/specs.json` and regenerating `sim/specs/specs.go` so `specs.Spec` has `ReferenceStat string \`json:"reference_stat"\`` (contract A7). The vocabulary is the fork's `proto.Stat` enum names in snake case.

**Files:**
- Modify: `api/internal/sims/specs.go`
- Test: `api/internal/sims/specs_test.go`

**Interfaces:**
- Consumes: `specs.All` and `specs.ByKey` with `ReferenceStat`.
- Produces: `SpecFidelity` gains `ReferenceStat string \`json:"reference_stat"\``.

- [ ] **Step 1: Write the failing test**

Add to `api/internal/sims/specs_test.go`:

```go
func TestEverySpecCardNamesItsReferenceStat(t *testing.T) {
	h := newHarness(t)
	// A measured row and an unmeasured one both carry it: the weights
	// page reads the reference off the card before anything has been
	// simmed.
	gap := 0.02
	if err := h.store.PutSpec(t.Context(), SpecFidelity{
		Spec: "warrior-fury", Parses: 50, MedianGap: &gap, EngineVersion: testEngine,
	}); err != nil {
		t.Fatal(err)
	}
	cards, err := h.store.Specs(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	if len(cards) == 0 {
		t.Fatal("no cards at all")
	}
	for _, c := range cards {
		want := specs.ByKey[c.Spec].ReferenceStat
		if want == "" {
			t.Fatalf("%s has no reference_stat in the generated spec list", c.Spec)
		}
		if c.ReferenceStat != want {
			t.Errorf("%s: reference_stat %q, want %q", c.Spec, c.ReferenceStat, want)
		}
	}
}
```

Add `"github.com/jhunthrop/foreversixty/sim/specs"` to `specs_test.go`'s imports, and reuse that file's existing pointer helper instead of the local `gap` variable if it has one.

- [ ] **Step 2: Run it and watch it fail**

Run: `go test -p 1 ./api/internal/sims/... -run TestEverySpecCardNamesItsReferenceStat -v`
Expected: FAIL — `c.ReferenceStat undefined`.

- [ ] **Step 3: Carry it on the card**

In `api/internal/sims/specs.go`, add to `SpecFidelity`, under `Spec`:

```go
	// ReferenceStat is the stat the weights tool normalises to 1.0 for
	// this spec (contract A7), in the fork's proto.Stat vocabulary. It
	// is read from the generated spec list rather than stored: it is a
	// property of the spec, not of a measurement, and a row measured
	// before the data lane changed it must not report the old one.
	ReferenceStat string `json:"reference_stat"`
```

In `Store.Specs`, replace the seeding loop with one that has the whole spec, not just its key:

```go
	for _, s := range specs.All {
		if s.Role != roleDPS {
			continue
		}
		// No MedianGap and no UpdatedAt: nothing has measured this spec,
		// and both fields marshal as null to say so.
		byspec[s.Spec] = SpecFidelity{
			Spec: s.Spec, ReferenceStat: s.ReferenceStat,
			State: SpecUnsupported, WorstActions: []WorstAction{},
		}
	}
```

and replace the `byspec[f.Spec] = f` line and the comment above it with:

```go
		// A measured row for a spec that is no longer in the list is
		// still shown: a rename should be visible, not silent. Its
		// reference stat is simply blank, because the list no longer has
		// one for it.
		f.ReferenceStat = specs.ByKey[f.Spec].ReferenceStat
		byspec[f.Spec] = f
```

Add `"github.com/jhunthrop/foreversixty/sim/specs"` to `specs.go`'s imports. `DPSSpecs()` in `simdep.go` is unchanged and still has its own callers.

- [ ] **Step 4: Run the tests**

Run: `go test -p 1 ./api/internal/sims/...`
Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git add api/internal/sims/specs.go api/internal/sims/specs_test.go
git commit -m "$(cat <<'EOF'
feat(api): spec cards carry the weights tool's reference stat

Read from the generated sim/specs struct, not stored: it is a property
of the spec, so a card measured before the data lane changed it must not
report the old one.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

### Task 9: A build has an owner, and `GET /v1/builds?mine=1`

Top Gear's talent list reads the signed-in player's own builds (contract 10.6).

**Files:**
- Create: `api/internal/db/migrations/0015_builds_owner.up.sql`
- Create: `api/internal/db/migrations/0015_builds_owner.down.sql`
- Modify: `api/internal/builds/store.go` (`Save`, `Get`, `GetMany`, `Mine`, `scanBuild`)
- Modify: `api/internal/builds/handler.go` (`Storer`, `Mount`, `save`, `mine`)
- Test: `api/internal/builds/store_test.go`, `api/internal/builds/handler_test.go`

**Interfaces:**
- Consumes: `auth.ActorFrom(ctx)`, `auth.RequireSession` (both already used by `api/internal/sims`).
- Produces:
  - `func (s *Store) Save(ctx context.Context, b Build, userID *int64) (Build, bool, error)`
  - `func (s *Store) Mine(ctx context.Context, userID int64, page int) (Page, error)`
  - `type Page struct { Rows []Build; Total, Page, PerPage int }`
  - `const PerPage = 100`
  - `GET /v1/builds?mine=1`

- [ ] **Step 1: Write the failing test**

Add to `api/internal/builds/store_test.go` (follow that file's own harness for a pool and a user id; it already has one for `Save`):

```go
func TestASavedBuildRemembersWhoSavedIt(t *testing.T) {
	s, owner := storeWithUser(t)
	b := aBuild("Fury")
	if _, _, err := s.Save(t.Context(), b, &owner); err != nil {
		t.Fatal(err)
	}
	page, err := s.Mine(t.Context(), owner, 1)
	if err != nil {
		t.Fatal(err)
	}
	if page.Total != 1 || len(page.Rows) != 1 || page.Rows[0].ID != b.ID {
		t.Fatalf("mine: %+v", page)
	}
	if page.Rows[0].Title != "Fury" || page.Rows[0].TreeVersion != b.TreeVersion {
		t.Errorf("the row is not the whole build: %+v", page.Rows[0])
	}
}

func TestAnAnonymousBuildBelongsToNobodyUntilSomeoneSavesIt(t *testing.T) {
	s, owner := storeWithUser(t)
	b := aBuild("Fury")
	// Saved by nobody first.
	if _, created, err := s.Save(t.Context(), b, nil); err != nil || !created {
		t.Fatalf("created %v err %v", created, err)
	}
	page, err := s.Mine(t.Context(), owner, 1)
	if err != nil {
		t.Fatal(err)
	}
	if page.Total != 0 {
		t.Fatalf("an unowned build showed in a history: %+v", page)
	}
	// The same build saved again by a signed-in player claims the row:
	// ids are content hashes, so this is the same row, and leaving it
	// ownerless would mean a player could never list a build they saved.
	if _, created, err := s.Save(t.Context(), b, &owner); err != nil || created {
		t.Fatalf("created %v err %v; a content-hash id is saved once", created, err)
	}
	page, err = s.Mine(t.Context(), owner, 1)
	if err != nil {
		t.Fatal(err)
	}
	if page.Total != 1 {
		t.Fatalf("the second saver did not claim the unowned row: %+v", page)
	}
}

func TestAnOwnedBuildIsNotReassignedByASecondSaver(t *testing.T) {
	s, first := storeWithUser(t)
	second := anotherUser(t, s)
	b := aBuild("Fury")
	if _, _, err := s.Save(t.Context(), b, &first); err != nil {
		t.Fatal(err)
	}
	if _, _, err := s.Save(t.Context(), b, &second); err != nil {
		t.Fatal(err)
	}
	mine, err := s.Mine(t.Context(), first, 1)
	if err != nil {
		t.Fatal(err)
	}
	theirs, err := s.Mine(t.Context(), second, 1)
	if err != nil {
		t.Fatal(err)
	}
	if mine.Total != 1 || theirs.Total != 0 {
		t.Fatalf("the first saver keeps it: mine %d theirs %d", mine.Total, theirs.Total)
	}
}
```

`storeWithUser`, `anotherUser` and `aBuild` are this file's own helpers: read `store_test.go` and add only the ones it does not already have, following how `api/internal/sims/harness_test.go` makes a user (`auth.Store{Pool: pool}.UpsertEmailUser`).

Add to `api/internal/builds/handler_test.go`:

```go
func TestMyOwnBuildsNeedASessionAndMineEqualsOne(t *testing.T) {
	// The handler tests run against fakeStore, so this one checks the
	// route's shape: the parameter is required and the session is.
	store := &fakeStore{}
	h := testRouter(t, store)

	w := httptest.NewRecorder()
	h.ServeHTTP(w, signedIn(httptest.NewRequest(http.MethodGet, "/v1/builds", nil)))
	if w.Code != http.StatusBadRequest {
		t.Errorf("status %d, want 400 without mine=1", w.Code)
	}

	w = httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/v1/builds?mine=1", nil))
	if w.Code != http.StatusUnauthorized {
		t.Errorf("status %d, want 401 without a session", w.Code)
	}
}
```

`signedIn` wraps a request with an actor the way `api/internal/sims/harness_test.go`'s server does (`auth.WithActor`); add it to `handler_test.go` if it is not there, and extend `testRouter` so its mux carries the actor.

- [ ] **Step 2: Run it and watch it fail**

Run: `go test -p 1 ./api/internal/builds/...`
Expected: FAIL to build — `too many arguments in call to s.Save`, `s.Mine undefined`.

- [ ] **Step 3: Write the migration**

`api/internal/db/migrations/0015_builds_owner.up.sql`:

```sql
-- A saved build remembers who saved it, so a signed-in player can list
-- their own (contract 10.6). Nullable: an anonymous save is still saved
-- and still shareable, it simply has no owner.
alter table builds add column if not exists user_id bigint references users(id) on delete set null;
create index if not exists builds_user_idx on builds (user_id, created_at desc);
```

`api/internal/db/migrations/0015_builds_owner.down.sql`:

```sql
drop index if exists builds_user_idx;
alter table builds drop column if exists user_id;
```

- [ ] **Step 4: One row scanner, then the owner and the list**

In `api/internal/builds/store.go`, add above `Get`:

```go
// buildRow is what both a single-row query and a multi-row one satisfy,
// so one function reads a build in one column order.
type buildRow interface{ Scan(dest ...any) error }

// buildColumns is that column order. Every query below selects exactly
// these, in this order, and scanBuild reads them.
const buildColumns = `id, class_id, race_id, tree_version, point_order, gear, title,
	created_at, views`

// scanBuild reads one row. It exists because Get, GetMany and Mine had
// three copies of the same conversions between them, and a fourth was
// one too many.
func scanBuild(row buildRow) (Build, error) {
	var (
		b               Build
		classID, raceID int16
		order           []int32
		title           *string
	)
	if err := row.Scan(&b.ID, &classID, &raceID, &b.TreeVersion, &order, &b.Gear, &title,
		&b.CreatedAt, &b.Views); err != nil {
		return Build{}, err
	}
	b.ClassID, b.RaceID = int(classID), int(raceID)
	b.PointOrder = make([]int, len(order))
	for i, v := range order {
		b.PointOrder[i] = int(v)
	}
	if b.Gear == nil {
		b.Gear = map[string]int{}
	}
	if title != nil {
		b.Title = *title
	}
	return b, nil
}
```

Rewrite `Get` and `GetMany` to select `buildColumns` and call `scanBuild`, keeping each one's own error wrapping and its `ErrNotFound` branch.

Change `Save`'s signature and its insert:

```go
// Save inserts b and returns the stored record with created true. When
// the id already exists nothing is written and the existing record is
// returned with created false, provided it really is the same build: an
// id is a content hash, so the title that was saved first wins. An
// existing row whose content differs is an id collision and returns
// ErrIDCollision.
//
// userID may be nil: an anonymous save is saved and shareable, it simply
// has no owner and never appears in anyone's list. A signed-in save of a
// build that already exists claims the row when nobody owns it yet —
// otherwise a player could save a build somebody had already shared and
// never find it in their own list — and leaves an owned row alone, the
// way the first title wins.
func (s *Store) Save(ctx context.Context, b Build, userID *int64) (Build, bool, error) {
```

with the insert gaining the column:

```go
		`insert into builds (id, class_id, race_id, tree_version, point_order, gear, title, user_id)
		 values ($1, $2, $3, $4, $5, $6, $7, $8)
		 on conflict (id) do nothing
		 returning created_at, views`,
		b.ID, classID, raceID, b.TreeVersion, order, gear, title, userID).
```

and, in the existing-row branch — after `sameContent` has confirmed it is the same build and before the function returns it — the claim:

```go
	if userID != nil {
		if _, err := s.Pool.Exec(ctx,
			`update builds set user_id = $2 where id = $1 and user_id is null`,
			b.ID, *userID); err != nil {
			return Build{}, false, fmt.Errorf("builds: claim %s: %w", b.ID, err)
		}
	}
```

Add the list, at the bottom of the file:

```go
// PerPage is the page size of a player's own build list, the same
// hundred the sim history and the rankings use.
const PerPage = 100

// Page is one page of a player's own builds.
type Page struct {
	Rows    []Build `json:"rows"`
	Total   int     `json:"total"`
	Page    int     `json:"page"`
	PerPage int     `json:"per_page"`
}

// Mine answers one page of a player's own builds, newest first.
func (s *Store) Mine(ctx context.Context, userID int64, page int) (Page, error) {
	if page < 1 {
		page = 1
	}
	out := Page{Rows: []Build{}, Page: page, PerPage: PerPage}
	if err := s.Pool.QueryRow(ctx,
		`select count(*) from builds where user_id = $1`, userID).Scan(&out.Total); err != nil {
		return Page{}, fmt.Errorf("builds: count: %w", err)
	}
	rows, err := s.Pool.Query(ctx,
		`select `+buildColumns+` from builds where user_id = $1
		 order by created_at desc, id limit $2 offset $3`,
		userID, PerPage, (page-1)*PerPage)
	if err != nil {
		return Page{}, fmt.Errorf("builds: list: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		b, err := scanBuild(rows)
		if err != nil {
			return Page{}, fmt.Errorf("builds: scan: %w", err)
		}
		out.Rows = append(out.Rows, b)
	}
	return out, rows.Err()
}
```

- [ ] **Step 5: Mount the route and pass the owner through**

In `api/internal/builds/handler.go`, widen `Storer`:

```go
// Storer is the part of Store the handlers use, so they can be tested
// without Postgres.
type Storer interface {
	Save(ctx context.Context, b Build, userID *int64) (Build, bool, error)
	Get(ctx context.Context, id string) (Build, error)
	Mine(ctx context.Context, userID int64, page int) (Page, error)
}
```

add the route:

```go
	mux.HandleFunc("GET /v1/builds", auth.RequireSession(s.mine))
```

pass the owner in `save`, replacing the `Store.Save` call:

```go
	var owner *int64
	if a := auth.ActorFrom(r.Context()); a.Signed() {
		id := a.UserID
		owner = &id
	}
	stored, created, err := s.Store.Save(r.Context(), b, owner)
```

and add the handler:

```go
// mine is the signed-in player's own builds. Like the sim history, the
// parameter is required so the route's meaning is on the URL: there is
// no "everyone's builds" list and inventing one by omission would be a
// surprise.
func (s *Service) mine(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("mine") != "1" {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "this list is mine=1 only",
			map[string]string{"mine": "1"})
		return
	}
	page := 1
	if v := r.URL.Query().Get("page"); v != "" {
		n, err := strconv.Atoi(v)
		if err != nil || n < 1 {
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "page must be 1 or more",
				map[string]string{"page": "a page number from 1"})
			return
		}
		page = n
	}
	out, err := s.Store.Mine(r.Context(), auth.ActorFrom(r.Context()).UserID, page)
	if err != nil {
		s.logger().Error("builds", "id", httpx.RequestIDFrom(r.Context()), "op", "mine", "err", err)
		httpx.WriteError(w, r, http.StatusInternalServerError, "internal",
			"could not read your builds just now", nil)
		return
	}
	// Per-account: never cached at a shared edge.
	w.Header().Set("Cache-Control", "private, no-store")
	httpx.WriteOK(w, r, http.StatusOK, out)
}
```

Add `"strconv"` and `"github.com/jhunthrop/foreversixty/api/internal/auth"` to the imports, and update `Mount`'s doc comment to name the third route. Update `fakeStore` in `handler_test.go` to the new `Save` signature and give it a `Mine`.

- [ ] **Step 6: Run the tests**

Run: `go test -p 1 ./api/internal/builds/...`
Expected: PASS.

- [ ] **Step 7: Commit**

```bash
git add api/internal/db/migrations/0015_builds_owner.up.sql \
        api/internal/db/migrations/0015_builds_owner.down.sql \
        api/internal/builds/store.go api/internal/builds/store_test.go \
        api/internal/builds/handler.go api/internal/builds/handler_test.go
git commit -m "$(cat <<'EOF'
feat(api): a saved build remembers its owner, and GET /v1/builds?mine=1

Top Gear's talent list reads a player's own builds. Ids are content
hashes, so a signed-in save claims a row nobody owns yet and leaves an
owned one alone. Get, GetMany and the new list now share one row scanner
instead of three copies of the same conversions.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

### Task 10: `GET /v1/phases`, and the boundaries the site and the API share

**Waits on:** the data lane creating `data/curated/phases.json` as `[ { "name", "start" } ]` (contract 10.4). The route serves the API's compiled table; the test is what keeps the two honest.

**Files:**
- Modify: `api/internal/phase/phase.go` (JSON tags)
- Create: `api/internal/phase/boundaries_test.go`
- Modify: `api/internal/server/server.go` (the route)
- Test: `api/internal/server/server_test.go`

**Interfaces:**
- Consumes: `data/curated/phases.json`.
- Produces: `GET /v1/phases` → `{"phases":[{"name","start"}]}`.

- [ ] **Step 1: Write the failing test**

Create `api/internal/phase/boundaries_test.go`:

```go
package phase

import (
	"encoding/json"
	"os"
	"testing"
	"time"
)

// TestBoundariesMatchTheCuratedFile is the reason the table in phase.go
// may be a compiled constant. The dates live in data/curated/phases.json,
// which the pipeline also emits to the site; this table is a copy, and a
// copy nothing compares is a copy that drifts. A phase that moved in one
// place and not the other re-buckets stored rankings silently, which is
// exactly what this catches.
func TestBoundariesMatchTheCuratedFile(t *testing.T) {
	b, err := os.ReadFile("../../../data/curated/phases.json")
	if err != nil {
		t.Fatal(err)
	}
	var file []struct {
		Name  string `json:"name"`
		Start string `json:"start"`
	}
	if err := json.Unmarshal(b, &file); err != nil {
		t.Fatalf("phases.json is not [{name,start}]: %v", err)
	}
	if len(file) != len(Boundaries) {
		t.Fatalf("%d phases in the file, %d in the table", len(file), len(Boundaries))
	}
	for i, want := range file {
		got := Boundaries[i]
		if got.Name != want.Name {
			t.Errorf("phase %d: name %q, want %q", i, got.Name, want.Name)
		}
		// An empty or absent start is the zero time: pre-beta has no
		// opening instant, it is simply everything before beta.
		var start time.Time
		if want.Start != "" {
			start, err = time.Parse(time.RFC3339, want.Start)
			if err != nil {
				t.Fatalf("phase %s: start %q is not RFC3339: %v", want.Name, want.Start, err)
			}
		}
		if !got.Start.Equal(start) {
			t.Errorf("phase %s: start %s, want %s", want.Name, got.Start, start)
		}
	}
}
```

Add to `api/internal/server/server_test.go`:

```go
func TestThePhasesRouteServesTheBoundaries(t *testing.T) {
	h := NewRouter(Deps{Version: "test"})
	w := httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/v1/phases", nil))
	if w.Code != http.StatusOK {
		t.Fatalf("status %d, want 200", w.Code)
	}
	var env struct {
		OK   bool `json:"ok"`
		Data struct {
			Phases []struct {
				Name  string    `json:"name"`
				Start time.Time `json:"start"`
			} `json:"phases"`
		} `json:"data"`
	}
	if err := json.NewDecoder(w.Body).Decode(&env); err != nil {
		t.Fatal(err)
	}
	if !env.OK || len(env.Data.Phases) != len(phase.Boundaries) {
		t.Fatalf("envelope: %+v", env)
	}
	if env.Data.Phases[0].Name != phase.Boundaries[0].Name {
		t.Errorf("first phase %q, want %q", env.Data.Phases[0].Name, phase.Boundaries[0].Name)
	}
	if cc := w.Header().Get("Cache-Control"); !strings.Contains(cc, "public") {
		t.Errorf("Cache-Control %q: four fixed instants are cacheable", cc)
	}
}
```

with `"time"`, `"encoding/json"` and `"github.com/jhunthrop/foreversixty/api/internal/phase"` imported.

- [ ] **Step 2: Run it and watch it fail**

Run: `go test ./api/internal/phase/... ./api/internal/server/...`
Expected: FAIL — no such file `data/curated/phases.json` (wait for the data lane), and `/v1/phases` answers 404.

- [ ] **Step 3: Tag the boundary for JSON**

In `api/internal/phase/phase.go`, replace `Boundary` with:

```go
// Boundary is one phase and the instant it opens. The tags are what
// GET /v1/phases serves; pre-beta's zero Start marshals as
// "0001-01-01T00:00:00Z", which is the truth about a phase that has no
// opening instant.
type Boundary struct {
	Name  string    `json:"name"`
	Start time.Time `json:"start"`
}
```

and extend the package doc's last paragraph to: "When a date moves or a phase is added, change this table, `data/curated/phases.json` and the site's `dates.json` together; `boundaries_test.go` fails if the first two disagree."

- [ ] **Step 4: Serve the route**

In `api/internal/server/server.go`, after the `GET /version` handler:

```go
	// The phase boundaries, for any client that has to bucket a date the
	// same way the rankings do. Four fixed instants that change only with
	// a deploy, so an hour at the edge is safe.
	mux.HandleFunc("GET /v1/phases", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=3600")
		httpx.WriteOK(w, r, http.StatusOK, map[string]any{"phases": phase.Boundaries})
	})
```

Add `"github.com/jhunthrop/foreversixty/api/internal/phase"` to the imports.

- [ ] **Step 5: Run the tests**

Run: `go test ./api/internal/phase/... ./api/internal/server/...`
Expected: PASS.

- [ ] **Step 6: Commit**

```bash
git add api/internal/phase/phase.go api/internal/phase/boundaries_test.go \
        api/internal/server/server.go api/internal/server/server_test.go
git commit -m "$(cat <<'EOF'
feat(api): GET /v1/phases, and a test that the table matches phases.json

The API's phase table is a compiled copy of data/curated/phases.json;
a copy nothing compares is a copy that drifts, and a phase that moved in
one place silently re-buckets stored rankings. Now the test compares it.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

### Task 11: Wiring, OpenAPI and the operator documentation

**Waits on:** the `sim` lane's `Plan` method on `runner.Native` and `runner.Fixture` (contract 10.2). Until both exist, `simEngine` cannot return a `sims.Engine`.

**Files:**
- Modify: `api/cmd/api/main.go`
- Modify: `api/openapi.yaml`
- Modify: `api/internal/server/openapi_test.go`
- Modify: `api/README.md`

**Interfaces:**
- Consumes: `sims.Engine` (Task 5); `runner.Native.Plan`, `runner.Fixture.Plan`.
- Produces: a deployment that mounts `POST /v1/sims/run`, `GET /v1/builds` and `GET /v1/phases`.

- [ ] **Step 1: Give the service a planner**

In `api/cmd/api/main.go`, change `simEngine`'s signature and comment:

```go
// simEngine is what every simulator job and the submit handler use: the
// real binary when the image carries one, and the checked-in fixture
// when it does not, so a deployment without the artifact still answers
// instead of failing. It runs sims and it counts bulk requests without
// running them (`forever-sim -plan`), which is why one value serves both
// the jobs and the service.
func simEngine(log *slog.Logger) sims.Engine {
	if _, err := os.Stat(runner.DefaultBinary); err == nil {
		return &runner.Native{}
	}
	log.Warn("sims", "state", "no engine binary at "+runner.DefaultBinary,
		"effect", "sims answer from the checked-in fixture result")
	return &runner.Fixture{}
}
```

and the service construction:

```go
	deps.Sims = &sims.Service{
		Store: simStore, Accounts: authStore, Planner: simEngine(log),
		EngineVersion: enginever.Version, Log: log,
	}
```

- [ ] **Step 2: Build it**

Run: `go build ./api/...`
Expected: success. A failure naming `Plan` means the `sim` lane's half is not merged; stop and wait rather than adding a shim.

- [ ] **Step 3: Add the new paths to the document's own test**

In `api/internal/server/openapi_test.go`, add `"/v1/phases"` to `requiredPaths` (`"/v1/builds"` is already there).

- [ ] **Step 4: Update the OpenAPI document**

In `api/openapi.yaml`:

Add `kind` and `headline` to `SimRow`:

```yaml
    SimRow:
      type: object
      properties:
        sim_id: { type: string }
        spec: { type: string }
        kind: { type: string, enum: [run, gear, talents, drops, weights] }
        dps: { type: number }
        headline:
          type: string
          maxLength: 120
          description: >-
            The one line the history list shows, composed when the result
            was written: "1,204 DPS", "+41 DPS from Vis'kag the
            Bloodletter", "3 upgrades on Ragnaros", "+18 DPS with 'Deep
            Fury'", "Melee Crit 1.00 · Agility 0.87".
        engine_version: { type: string }
        created_at: { type: string, format: date-time }
        title: { type: string }
```

Add to `SpecFidelity`, under `spec`:

```yaml
        reference_stat:
          type: string
          description: >-
            The stat the weights tool normalises to 1.0 for this spec, in
            the engine's own vocabulary (melee_crit, attack_power, …).
```

Add the `kind` parameter to `listMySims`:

```yaml
        - name: kind
          in: query
          schema: { type: string, enum: [run, gear, talents, drops, weights] }
          description: Narrows the list to one tool. Omitted or empty is every kind; an unknown value is 400.
```

Replace `POST /v1/sims/run`'s `'400'` line:

```yaml
        '400':
          description: >-
            The request cannot be run. error.code is "invalid" for a
            malformed envelope, "cap_exceeded" when the expansion is past
            the server lane's cap (error.fields carries cap and
            combinations as decimal strings), or "too_large" when the
            planner estimates the run past the job's budget (error.fields
            carries estimate_sec and budget_sec).
```

Add to `GET /v1/sims/{id}/progress`'s data properties, under `dps`:

```yaml
                          stage: { type: integer, description: Which stage of a bulk run is going; 0 for a plain run. }
                          combos_done: { type: integer }
                          combos_total: { type: integer }
```

Add `get` to the existing `/v1/builds` path item, beside its `post`:

```yaml
    get:
      summary: The caller's own saved builds
      operationId: listMyBuilds
      parameters:
        - { name: mine, in: query, required: true, schema: { type: string, enum: ['1'] } }
        - { name: page, in: query, schema: { type: integer, minimum: 1 } }
      responses:
        '200':
          description: One page of the caller's builds, newest first
          content:
            application/json:
              schema:
                allOf:
                  - $ref: '#/components/schemas/Envelope'
                  - type: object
                    properties:
                      data:
                        type: object
                        properties:
                          rows:
                            type: array
                            items: { $ref: '#/components/schemas/Build' }
                          total: { type: integer }
                          page: { type: integer }
                          per_page: { type: integer, example: 100 }
        '400': { description: This list is mine=1 only }
        '401': { description: Sign in first }
```

and a new path item:

```yaml
  /v1/phases:
    get:
      summary: The content phase boundaries
      operationId: getPhases
      description: >-
        The same four instants the rankings bucket by, from
        data/curated/phases.json. Cached an hour: they change only with a
        deploy.
      responses:
        '200':
          description: The phases, in order
          content:
            application/json:
              schema:
                allOf:
                  - $ref: '#/components/schemas/Envelope'
                  - type: object
                    properties:
                      data:
                        type: object
                        properties:
                          phases:
                            type: array
                            items:
                              type: object
                              properties:
                                name: { type: string, example: raids-1 }
                                start: { type: string, format: date-time }
```

- [ ] **Step 5: Run the document's own test**

Run: `go test ./api/internal/server/...`
Expected: PASS.

- [ ] **Step 6: Update the README**

In `api/README.md`, under "The simulator's Cloud Run jobs", after the paragraph beginning "`sim-run` is executed by the API", add:

```markdown
The `--cpu 4 --task-timeout 15m` on `sim-run` is load-bearing, not just a
ceiling. A Top Gear, Droptimizer, talent-compare or stat-weights submit is
sized before it is queued: the API asks the binary to expand the request
without running it (`forever-sim -plan`), divides the precision ladder's
total iterations by the engine's measured rate times those four CPUs, and
refuses anything past `sims.BulkBudget` (840 seconds, one minute inside
the task timeout) with `400 too_large` and the estimate. Changing the
job's CPU count means changing `simJobCPUs` in
`api/internal/sims/simdep.go` in the same commit, or every estimate is
wrong. The rate itself is `measure.NativeIterationsPerCPUSecond`, which
`sim/measure` publishes from its own benchmark — never restate it here.
Jobs are created by hand, so nothing enforces this but this paragraph.
```

In the route table, add three lines:

```markdown
| `GET /v1/sims?mine=1&kind=` | The caller's own sims, newest first, optionally narrowed to one tool (`run`, `gear`, `talents`, `drops`, `weights`). Each row carries its kind and a composed one-line headline. An unknown kind is 400. |
| `GET /v1/builds?mine=1` | The signed-in player's own saved builds, newest first. A build saved anonymously has no owner; a later signed-in save of the same build (ids are content hashes) claims it if nobody owns it yet. |
| `GET /v1/phases` | The content phase boundaries, cached an hour. The table is compiled in; `api/internal/phase/boundaries_test.go` holds it to `data/curated/phases.json`. |
```

No new environment variable: the CPU count is a constant because it describes the job definition, not a deployment.

- [ ] **Step 7: Run the API's tests once more**

Run: `go test -p 1 ./api/...`
Expected: PASS. Re-run once on a `sims_pkey` collision.

- [ ] **Step 8: Commit**

```bash
git add api/cmd/api/main.go api/openapi.yaml api/internal/server/openapi_test.go api/README.md
git commit -m "$(cat <<'EOF'
feat(api): wire the planner, and document the kinds, the lists and the budget

simEngine now returns a sims.Engine — one value that both runs sims and
counts bulk requests without running them — so the service can size a
submit. OpenAPI gains kind, headline, reference_stat, the two refusals,
the stage progress fields, GET /v1/builds and GET /v1/phases; the README
says why the job's --cpu 4 is load-bearing.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

## Self-review against the contract

**Section 8 and 10.6 coverage.** Premium submit on `POST /v1/sims/run` with both refusals: Tasks 5 and 6. Decimal-string error fields: Tasks 5 and 6, asserted in both refusal tests. Migration `0014_sim_kinds` with all five columns: Task 1. `headline` composed at save by section 8's rules plus 10.6's "… and N more" and empty cases: Tasks 2 and 3. `GET /v1/sims?mine=1&kind=` with `kind` and `headline` on the row: Tasks 3 and 4. Progress rows carrying `stage`, `combos_done`, `combos_total`: Task 7. `GET /v1/specs` `reference_stat`: Task 8. `builds.user_id` and `GET /v1/builds?mine=1`: Task 9. `GET /v1/phases` and the `phase.Boundaries` test: Task 10. `GET /v1/sims/<id>` unchanged — nothing to do; Task 7 asserts `Combos`, `Equipped` and `Stages` survive the round trip.

**Section 10.1 coverage.** A1 `ValidateLane(LaneServer)`: Task 5 step 5. A2 cap 5,000, rate from `sim/measure`, 840-second budget: Task 6, with `TestTheServerCapAndTheBudgetAgree` pinning the ruling's own reasoning. A3 a bulk request's iterations are the precision's final-stage count: Task 5's `bulkBody`. A6 `Substitution.Name` for items and `SourceName`: Task 2 — no name is ever derived from an origin id. A7 the `proto.Stat` vocabulary and `specs.Spec.ReferenceStat`: Tasks 2 and 8. A11 additive progress widening: Task 7, which is no longer blocking and keeps the store on its own `Tick`.

**Section 10.2 coverage.** `forever-sim -plan` is how the API counts: Task 5's `Planner`, wired in Task 11. `simCount`, `simNeedsMore` and `simValidate` are browser exports and belong to the web lane; nothing here calls them.

**Design sections 10.2, 10.3 and 11 coverage.** "The job row records the kind": Task 1. "Runs the native planner loop and streams stage progress": Task 7. "A request the planner estimates past that is refused at submit with the estimate": Task 6 and Task 11's README paragraph. "Saved sims of every kind are public at `/sim/<id>`": already true. "The API test suite covers the kind column, submit-time cap refusal and job progress for a bulk request": Tasks 1, 5 and 7. "The nightly validation job is unchanged": Task 7's `TestTheValidationJobStillRunsPlainSims`.

**Placeholder scan.** Every code step carries the code. Four steps direct the engineer to read an existing test file before adding to it — `validate_test.go`'s deps helper (Task 7), `specs_test.go`'s pointer helper (Task 8), `builds/store_test.go`'s fixtures and `builds/handler_test.go`'s router (Task 9) — because reusing that file's fakes is the point of the instruction; each names exactly which helper and what it must do.

**Type consistency.** `Headline`, `withThousands`, `trimRunes`, `maxHeadline` (Task 2) are used unchanged in Tasks 3, 5 and 6. `Store.Mine(ctx, userID, page, kind)` (Task 3) is called with four arguments in Tasks 3, 4 and 7. `Planner`/`Engine` (Task 5) are the types Task 11 wires. `Tick` (Task 7) is the only shape `Advance` takes after that task. `BulkBudget` is one name throughout Tasks 6 and 11 — there is no `BulkRunTimeout`. `builds.Save(ctx, b, userID)` (Task 9) has one signature everywhere it appears.

---

## Rulings from section 10 that could not be applied as written

1. **A11's "additive" is not achievable for `runner.Progress` as it stands today** — it is `func(done int, mean float64)`, and a Go func type cannot gain fields — so Task 7 tells the engineer to read `sim/runner/runner.go` and match whichever shape the `sim` lane landed, and keeps the API's storage on its own `Tick` so only one callback literal ever has to move.
