# Simulator parity — the API lane — Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Teach `api/internal/sims` the four new request kinds — gear, talents, drops, weights — so the premium lane sizes them before it queues them, the job streams stage progress, and the history list names each one.

**Architecture:** Nothing about *which* sims run lives in the API. `sim/bulk` owns expansion, staging and ranking; the API reaches it exactly twice — once before queueing, through a plan-only invocation of the native binary that counts combinations without simulating anything, and once inside the job, where the binary detects the kind and runs its own plan-rank loop. The API's own new work is three things: a `kind` column and the index the filtered history query needs, a pure `Headline` function that turns a stored result into one line of text, and two submit-time refusals (`cap_exceeded`, `too_large`).

**Tech Stack:** Go 1.25 (workspace, `GOTOOLCHAIN=auto`), pgx v5, golang-migrate (embedded `api/internal/db/migrations/*.sql`), `net/http` `ServeMux`, the shared `httpx` envelope.

**Spec:**
- Design: `docs/superpowers/specs/2026-09-19-simulator-parity-design.md` (sections 10 and 11)
- Contract: `docs/superpowers/specs/2026-09-19-simulator-parity-interfaces.md` (section 8 is this lane's binding interface; sections 1, 2 and 4 are what it consumes)

## Global Constraints

- **The contract is the authority.** Where this plan and the contract disagree, the contract wins — and a name this lane needs that the contract does not have is added to the contract first, in its own commit (Task 1), before any code uses it.
- **Go workspace toolchain:** `GOTOOLCHAIN=auto`. The repo's `go.work` uses `./api ./companion ./logs ./sim`; run every Go command from the repository root.
- **The api module must never import the engine.** `api/internal/sims`'s package doc says it: only `sim/api`, `sim/enginever`, `sim/specs` and `sim/runner` are importable. `sim/bulk` reads `sim/internal/simdb`, which imports `github.com/wowsims/classic/sim/core/proto` — so the API can never import `sim/bulk`, and additionally `sim/internal/...` is unreachable from another module. Every count this lane needs comes from the native binary as a subprocess.
- **DB tests share one database.** `TEST_DATABASE_URL` points at a single Postgres; `newHarness` truncates `sims` on entry and the tests use literal twelve-character ids. Run the package's tests serially — `go test -p 1 ./api/internal/sims/...` — and if a run fails with a `sims_pkey` duplicate-key collision, re-run it once before investigating; a second collision is a real bug (two tests using the same literal id).
- **Scoped tests per task:** `go test ./api/internal/sims/...` (add `-run` for a single test while iterating). Never run the whole repository suite for a task — that is CI's job.
- **One commit per task**, message `feat(api): ...` or `test(api): ...` (Task 1 is the exception: it is a spec change, `docs(sim): ...`). Every commit ends with the trailer:

  ```
  Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
  ```
- **Never `git stash`.** The stash stack is shared with other worktrees and other sessions. Set work aside with a temporary WIP commit instead.
- **Migration numbering:** `0014_sim_kinds` is the only migration this lane adds. Every column it needs goes in that one file.
- **Blocked tasks are marked.** Tasks 8, 9 and 10 consume names the `sim` and `data` lanes implement. Each says exactly what it waits for. Tasks 2–7 depend on nothing outside this lane.

---

## File Structure

| File | Responsibility |
| --- | --- |
| `docs/superpowers/specs/2026-09-19-simulator-parity-interfaces.md` | Contract. Task 1 amends sections 1, 2, 4 and 8. |
| `api/internal/db/migrations/0014_sim_kinds.up.sql` / `.down.sql` | New: `kind`, `headline`, `stage`, `combos_done`, `combos_total`, and the `(user_id, kind, created_at desc)` index. |
| `api/internal/sims/headline.go` | New: `Headline` and its formatting helpers. Pure functions, no I/O, no database. |
| `api/internal/sims/headline_test.go` | New: the headline table test. No database. |
| `api/internal/sims/store.go` | `Save`/`Queue`/`Finish` write `kind` and `headline`; `Mine` filters by kind; `Advance`/`Progress` carry the stage fields. |
| `api/internal/sims/handler.go` | `Mount` requires a `Planner`; `mine` reads `kind=`; `trimTitle` generalised to `trimRunes`. |
| `api/internal/sims/run.go` | Forces the server lane's cap onto the request, then `cap_exceeded` / `too_large` before it queues. |
| `api/internal/sims/job.go` | Per-kind run timeout; the widened progress callback. |
| `api/internal/sims/simdep.go` | The `Planner` and `Engine` interfaces, and the native-rate constants. |
| `api/internal/sims/specs.go` | `SpecFidelity.ReferenceStat`, filled from `sim/specs`. |
| `api/cmd/api/main.go` | `simEngine` returns a `sims.Engine`; the service gets a `Planner`. |
| `api/openapi.yaml`, `api/README.md` | The route, schema and operator documentation. |

---

### Task 1: The contract amendment

This lane needs six names the contract does not have. The contract's own rule — "A lane that needs a name not written here adds it here first, in its own commit, and the other lanes pick it up" — makes this the first task, and it is a documentation commit only.

**Files:**
- Modify: `docs/superpowers/specs/2026-09-19-simulator-parity-interfaces.md`

**Interfaces:**
- Consumes: nothing.
- Produces: `api.Kinds`, `api.PlanSummary`, `forever-sim -plan`, `runner.Planner`, `runner.Progress func(api.Progress)`, `Substitution.Name` for items, `specs.Spec.ReferenceStat`. Tasks 2–10 code against these names; the `sim` and `data` lanes implement the ones outside `api/`.

- [ ] **Step 1: Add the kinds slice to section 1.1**

Directly under the `Kind()` declaration in section 1.1, add:

````markdown
```go
// Kinds is the closed set. The API validates ?kind= against it and
// the history filter has no other vocabulary.
var Kinds = []string{KindRun, KindGear, KindTalents, KindDrops, KindWeights}
```
````

- [ ] **Step 2: Add `PlanSummary` as a new section 1.8, at the end of section 1**

````markdown
### 1.8 PlanSummary

Counting a bulk request is not running one. The API cannot import
`sim/bulk` — it reads `sim/internal/simdb`, which imports the engine,
and the api module's image build (`GOWORK=off`) has no replace for it —
so the count crosses the boundary as JSON from the native binary.

```go
type PlanSummary struct {
    Kind            string `json:"kind"`
    Combinations    int    `json:"combinations"`     // what Expand found; reported even when it is over Cap
    Cap             int    `json:"cap"`              // the lane's cap, 0 for a kind that has none (weights)
    IterationsTotal int    `json:"iterations_total"` // every stage of the precision's ladder summed, the equipped set included
}
```

`IterationsTotal` is `sim/bulk`'s arithmetic, not its caller's: the
stage ladder in 1.3 is stated in one place and summed in one place.
````

- [ ] **Step 3: Add the native plan mode to section 4**

After the paragraph in section 4 that begins "`simRun` is unchanged", add:

````markdown
`forever-sim -plan` is the same expansion without a simulation: it reads
a SimRequest on stdin and writes a `PlanSummary` (1.8) on stdout, then
exits 0. `ErrCapExceeded` is not a failure here — the count is what the
caller asked for — so the summary carries `Combinations > Cap` and the
caller decides. A request whose kind is `run` gets
`{"kind":"run","combinations":0,"cap":0,"iterations_total":<iterations>}`.

`sim/runner` gains the Go side of both:

```go
// Planner counts a request without running it.
type Planner interface {
    Plan(ctx context.Context, req api.SimRequest) (api.PlanSummary, error)
}

// Progress widens: the tick on stderr now carries the stage fields,
// so the callback carries the whole api.Progress rather than two
// numbers. Native and Fixture both implement Planner.
type Progress func(p api.Progress)
```
````

- [ ] **Step 4: Name items in `Substitution.Name`, in section 2**

Replace the `Name` line of the `Substitution` struct with:

```go
    Name    string `json:"name,omitempty"`    // item: the item's name from simdb; talents/set: the loadout or set name
```

Add under the struct: "`Name` is filled for every kind. The API composes
history headlines from stored results and has no item table of its own;
`sim/bulk` has simdb open already."

- [ ] **Step 5: Correct and complete section 8**

Replace the first bullet of section 8 with:

````markdown
- `POST /v1/sims/run` (the premium submit; `POST /v1/sims` is the
  browser-result save and is unchanged): unchanged path and premium
  gate. The server overwrites `bulk.cap` with the server lane's cap
  before validating — an echoed cap records what bounded a saved
  request, it never raises the bound. A cap breach answers `400
  cap_exceeded`, a request the planner estimates past the job's timeout
  answers `400 too_large`. Both carry their numbers in the envelope's
  `error.fields`, which is `map[string]string`, so they are decimal
  strings: `{"cap":"20000","combinations":"31200"}` and
  `{"estimate_sec":"1406","budget_sec":"840"}`.
````

Replace the `GET /v1/sims?mine=1` bullet's parenthesis with: "(the API
composes and stores it at write time: reading it back out of the result
blob would detoast a few hundred kilobytes of JSON per row, a hundred
rows to a page. `sims.headline` is the column.)"

- [ ] **Step 6: Name the reference stat's Go field, in section 1.4**

Under `WeightsSpec`, add: "`Reference`'s default per spec is
`data/curated/specs.json`'s new `reference_stat` field, which the data
lane's generator emits as `specs.Spec.ReferenceStat string
\`json:\"reference_stat\"\``. `GET /v1/specs` reads it from there."

- [ ] **Step 7: Commit**

```bash
git add docs/superpowers/specs/2026-09-19-simulator-parity-interfaces.md
git commit -m "$(cat <<'EOF'
docs(sim): the API's plan-only count, the kinds slice and the headline column

The API cannot import sim/bulk — it reads sim/internal/simdb, which
imports the engine — so the combination count crosses as JSON from
forever-sim -plan. Names PlanSummary, runner.Planner, the widened
runner.Progress, api.Kinds, Substitution.Name for items, and
specs.Spec.ReferenceStat. Corrects section 8's route name and says
error.fields carries decimal strings.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

### Task 2: Migration 0014 and the kind column

**Files:**
- Create: `api/internal/db/migrations/0014_sim_kinds.up.sql`
- Create: `api/internal/db/migrations/0014_sim_kinds.down.sql`
- Modify: `api/internal/sims/store.go` (`Save`, `Queue`)
- Test: `api/internal/sims/store_test.go`

**Interfaces:**
- Consumes: `simapi.SimRequest.Kind() string` and the `simapi.Kind*` constants (contract 1.1).
- Produces: the `sims.kind`, `sims.headline`, `sims.stage`, `sims.combos_done` and `sims.combos_total` columns, and `sims_user_kind_idx`. Tasks 4, 5 and 8 read them.

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
		Stats: []string{"strength", "crit"}, Reference: "crit",
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

- [ ] **Step 2: Run it and watch it fail**

Run: `go test -p 1 ./api/internal/sims/... -run TestEveryRowRecordsWhichToolProducedIt -v`
Expected: FAIL — `column "kind" does not exist` (or a compile error on `simapi.BulkSpec` if the `sim` lane's envelope has not landed; in that case this task waits on contract 1.3).

- [ ] **Step 3: Write the migration**

`api/internal/db/migrations/0014_sim_kinds.up.sql`:

```sql
-- Every sim row says which tool produced it, carries the one line the
-- history list shows for it, and — while a bulk run is going — how far
-- through its stages it is.
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
Expected: PASS. On a `sims_pkey` duplicate-key failure, re-run once.

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

### Task 3: The headline

One pure function, one table test, no database. Every number the history list shows is formatted here and nowhere else.

**Files:**
- Create: `api/internal/sims/headline.go`
- Create: `api/internal/sims/headline_test.go`
- Modify: `api/internal/sims/handler.go` (generalise `trimTitle`)

**Interfaces:**
- Consumes: `simapi.SimResult`, `simapi.Combo`, `simapi.Substitution`, `simapi.StatWeight` (contract 2); `simapi.SimRequest.Kind()` (contract 1.1).
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
		mode := kind
		res.Request.Bulk = &simapi.BulkSpec{Mode: mode, Precision: simapi.PrecisionNormal}
	case simapi.KindWeights:
		res.Request.Weights = &simapi.WeightsSpec{Reference: "crit"}
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

func drop(name, origin string, delta float64) simapi.Combo {
	return simapi.Combo{
		Substitutions: []simapi.Substitution{
			{Kind: "item", Slot: "main_hand", ItemID: 17182, Name: name, Origin: origin},
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
			"a four-figure run keeps one comma",
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
			"top gear with two substitutions names the first and counts the rest",
			withKind(simapi.KindGear, func(r *simapi.SimResult) {
				r.Combos = []simapi.Combo{{
					Substitutions: []simapi.Substitution{
						{Kind: "item", ItemID: 17182, Name: "Vis'kag the Bloodletter", Origin: "bag"},
						{Kind: "item", ItemID: 16963, Name: "Onslaught Girdle", Origin: "bag"},
					},
					Delta: simapi.Estimate{Mean: 63.5},
				}}
			}),
			"+64 DPS from Vis'kag the Bloodletter and 1 more",
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
					drop("Perdition's Blade", "drop:raid:mc:ragnaros", 55),
					drop("Spinal Reaper", "drop:raid:mc:ragnaros", 12),
					drop("Malistar's Defender", "drop:raid:mc:ragnaros", 1),
					drop("Band of Accuria", "drop:raid:mc:ragnaros", -4),
				}
			}),
			"3 upgrades on Ragnaros",
		},
		{
			"one upgrade is singular",
			withKind(simapi.KindDrops, func(r *simapi.SimResult) {
				r.Combos = []simapi.Combo{drop("Spinal Reaper", "drop:raid:mc:ragnaros", 12)}
			}),
			"1 upgrade on Ragnaros",
		},
		{
			"a hyphenated source id becomes words",
			withKind(simapi.KindDrops, func(r *simapi.SimResult) {
				r.Combos = []simapi.Combo{drop("Maladath", "drop:raid:bwl:broodlord-lashlayer", 30)}
			}),
			"1 upgrade on Broodlord Lashlayer",
		},
		{
			"nothing gained on a source is still an answer",
			withKind(simapi.KindDrops, func(r *simapi.SimResult) {
				r.Combos = []simapi.Combo{drop("Band of Accuria", "drop:raid:mc:ragnaros", -4)}
			}),
			"no upgrades on Ragnaros",
		},
		{
			"two sources in one request name neither",
			withKind(simapi.KindDrops, func(r *simapi.SimResult) {
				r.Combos = []simapi.Combo{
					drop("Perdition's Blade", "drop:raid:mc:ragnaros", 55),
					drop("Maladath", "drop:raid:bwl:broodlord-lashlayer", 30),
				}
			}),
			"2 upgrades",
		},
		{
			"weights are the top two, the reference first",
			withKind(simapi.KindWeights, func(r *simapi.SimResult) {
				r.Weights = []simapi.StatWeight{
					{Stat: "agility", Weight: 0.874},
					{Stat: "crit", Weight: 1},
					{Stat: "attack_power", Weight: 0.5},
				}
			}),
			"Crit 1.00 · Agility 0.87",
		},
		{
			"a multi-word stat id reads as words",
			withKind(simapi.KindWeights, func(r *simapi.SimResult) {
				r.Weights = []simapi.StatWeight{
					{Stat: "attack_power", Weight: 1},
					{Stat: "spell_power", Weight: 0.4},
				}
			}),
			"Attack Power 1.00 · Spell Power 0.40",
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

The import block is `"strings"`, `"testing"`, `"unicode/utf8"` and `simapi "github.com/jhunthrop/foreversixty/sim/api"`.

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

// maxHeadline bounds the stored line, in runes. A set or loadout name
// is a member's own string, so the line it lands in needs a bound.
const maxHeadline = 120

// headlineWeights is how many stat weights the weights headline names:
// the reference and its nearest rival, which is what makes the line
// mean anything at a glance.
const headlineWeights = 2

// weightSeparator joins them. It is a middle dot with spaces, as the
// contract's example spells it.
const weightSeparator = " · "

// dropOrigin is the prefix every Droptimizer candidate's origin
// carries: "drop:<source-id>" (contract 1.3).
const dropOrigin = "drop:"

// Headline is the one line a history row shows for a finished result.
// The contract (section 8) fixes one form per kind; this function is
// the only place any of them is composed, and the only place a figure
// on that list is formatted.
//
// It reads the stored result and nothing else — no item table, no loot
// table — which is why sim/bulk fills Substitution.Name for items
// (contract 2) and why a drop source is read off the origin id.
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
// many others rode with it. An item with no name falls back to its id
// rather than vanishing — an unnamed item is a data gap worth seeing.
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
// "" when they came from more than one (or from somewhere that is not a
// drop at all). A Droptimizer run is normally one boss; a request that
// mixed two is counted without being named rather than named wrongly.
func sharedSource(combos []simapi.Combo) string {
	origin := ""
	for _, c := range combos {
		for _, s := range c.Substitutions {
			if !strings.HasPrefix(s.Origin, dropOrigin) {
				return ""
			}
			if origin == "" {
				origin = s.Origin
			} else if origin != s.Origin {
				return ""
			}
		}
	}
	return sourceLabel(origin)
}

// sourceLabel turns "drop:raid:mc:broodlord-lashlayer" into "Broodlord
// Lashlayer". The API has no loot table — it is the data lane's file and
// the web reads it — so the label comes off the id, whose last segment
// is a slug of the source's name by construction (contract 6.1).
func sourceLabel(origin string) string {
	id := strings.TrimPrefix(origin, dropOrigin)
	if id == "" {
		return ""
	}
	segments := strings.Split(id, ":")
	return titleWords(strings.ReplaceAll(segments[len(segments)-1], "-", " "))
}

// weightsHeadline is the stat weights line: the two largest weights,
// each to two decimals. The reference stat is exactly 1 (contract 2), so
// it leads unless something beat it, which is itself worth seeing.
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

// statLabel turns an IDS.md stat id into its display form:
// "attack_power" becomes "Attack Power". There is no table to keep in
// step, which is the point: a stat the engine gains is labelled the day
// it appears.
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

// signedDPS is a delta: always signed, so "+0 DPS" reads as a
// comparison that found nothing rather than as an absolute figure.
func signedDPS(mean float64) string {
	n := roundDPS(mean)
	if n < 0 {
		return "-" + withThousands(-n) + " DPS"
	}
	return "+" + withThousands(n) + " DPS"
}

// roundDPS is how every DPS figure on the history list is rounded: to
// the nearest whole, halves away from zero, which is math.Round's own
// rule.
func roundDPS(mean float64) int64 { return int64(math.Round(mean)) }

// withThousands groups n into three-digit blocks with commas. There is
// no dependency for this: golang.org/x/text's printer is a locale
// package for one string, and the site's numbers are English-grouped
// everywhere else too.
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

// trimRunes cuts s to at most max runes. The cut is by rune, so a
// string in any script keeps its last character whole and the result is
// always valid UTF-8. Both the member's title and the composed headline
// are bounded this way.
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

Headline composes the contract's five forms from a stored result and
nothing else: no item table, no loot table. Item names come from
Substitution.Name, a drop source from the origin id's last segment.
Every DPS figure on the history list is rounded and grouped here.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

### Task 4: The history row carries kind and headline

**Files:**
- Modify: `api/internal/sims/store.go` (`Row`, `Save`, `Finish`, `Mine`)
- Test: `api/internal/sims/store_test.go`

**Interfaces:**
- Consumes: `Headline(simapi.SimResult) string` (Task 3); the `kind` and `headline` columns (Task 2).
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
	req.Weights = &simapi.WeightsSpec{Stats: []string{"crit", "agility"}, Reference: "crit"}
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
		{Stat: "crit", Weight: 1}, {Stat: "agility", Weight: 0.874},
	}
	if err := h.store.Finish(t.Context(), "cccccccccccc", done); err != nil {
		t.Fatal(err)
	}
	finished, err := h.store.Mine(t.Context(), h.owner, 1, simapi.KindWeights)
	if err != nil {
		t.Fatal(err)
	}
	if len(finished.Rows) != 1 || finished.Rows[0].Headline != "Crit 1.00 · Agility 0.87" {
		t.Fatalf("finished row: %+v", finished.Rows)
	}
}
```

Also update the existing `TestMyHistoryIsMineAndNewestFirst` call to `h.store.Mine(t.Context(), h.owner, 1, "")`.

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

Task 5 replaces the `""` with the query parameter.

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

### Task 5: `GET /v1/sims?mine=1&kind=`

**Files:**
- Modify: `api/internal/sims/handler.go` (`mine`)
- Test: `api/internal/sims/handler_test.go`

**Interfaces:**
- Consumes: `Store.Mine(ctx, userID, page, kind)` (Task 4); `simapi.Kinds` (contract 1.1, Task 1).
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

### Task 6: Counting a bulk request before it is queued — `cap_exceeded`

**Files:**
- Modify: `api/internal/sims/simdep.go` (the `Planner` and `Engine` interfaces)
- Modify: `api/internal/sims/handler.go` (`Service.Planner`, `Mount`)
- Modify: `api/internal/sims/run.go`
- Test: `api/internal/sims/run_test.go`, `api/internal/sims/harness_test.go`

**Interfaces:**
- Consumes: `simapi.PlanSummary`, `simapi.Caps`, `simapi.BulkSpec` (contract 1.3, 1.8).
- Produces:
  - `type Planner interface { Plan(ctx context.Context, req simapi.SimRequest) (simapi.PlanSummary, error) }`
  - `type Engine interface { runner.Runner; Planner }`
  - `Service.Planner Planner`
  - `func (s *Service) checkSize(w http.ResponseWriter, r *http.Request, req simapi.SimRequest) bool`
  - `const planTimeout = 30 * time.Second`

- [ ] **Step 1: Write the failing test**

Add to `api/internal/sims/harness_test.go`, next to `fakePremium`:

```go
// fakePlanner answers the plan-only count without a binary. The real
// one shells out to forever-sim -plan; nothing about the handler's
// decision needs a subprocess to be tested.
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

and give the harness one, by adding a field to `harness`:

```go
	planner *fakePlanner
```

and setting it in `newHarness`, replacing the two lines that build the service:

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
// bulk block on it.
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
	if env.Error.Code != "cap_exceeded" {
		t.Fatalf("code %q, want cap_exceeded", env.Error.Code)
	}
	if env.Error.Fields["combinations"] != "31200" ||
		env.Error.Fields["cap"] != strconv.Itoa(simapi.Caps[simapi.LaneServer]) {
		t.Fatalf("fields %+v", env.Error.Fields)
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
// (contract 4): sim/bulk cannot be imported here — it reads
// sim/internal/simdb, which imports the engine, and the api module's
// image build has no replace for it — so the count crosses the boundary
// as a subprocess's JSON, exactly the way a run does.
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

Add `"context"` and `"github.com/jhunthrop/foreversixty/sim/runner"` to that file's imports.

- [ ] **Step 4: Carry a planner on the service and require it to mount the route**

In `api/internal/sims/handler.go`, add to `Service`, under `Jobs`:

```go
	// Planner counts a bulk or weights request before it is queued.
	// Nil means the premium lane cannot size one, and the run route is
	// not mounted: accepting a request we cannot bound would be worse
	// than not offering the route.
	Planner Planner
```

and change `Mount`'s condition:

```go
	if s.Jobs != nil && s.Accounts != nil && s.Planner != nil {
		mux.HandleFunc("POST /v1/sims/run", auth.RequireSession(s.run))
	}
```

Update `Mount`'s doc comment's last sentence to: "without the job runner, the accounts store or the planner there is nothing behind it, and a 404 is a truer answer than a 500."

- [ ] **Step 5: Force the cap and check it**

In `api/internal/sims/run.go`, insert between `req.Encounter = withEncounterDefaults(...)` and `if err := req.Validate()`:

```go
	if req.Bulk != nil {
		// The lane's cap is the server's to set. A request echoes the
		// cap that bounded it (contract 1.3) so a saved request says
		// what it ran under; it is not a control the client holds.
		req.Bulk.Cap = simapi.Caps[simapi.LaneServer]
	}
```

and, after the `Validate` block and before the `id := auth.Base32ID(...)` line:

```go
	if req.Kind() != simapi.KindRun && s.checkSize(w, r, req) {
		return
	}
```

Then add at the bottom of `run.go`:

```go
// planTimeout bounds the plan-only subprocess. Expansion loads the item
// database and walks the candidates; it runs no iterations, so thirty
// seconds is a bound on something pathological rather than a budget.
const planTimeout = 30 * time.Second

// checkSize counts what the request would expand to, without running
// any of it, and refuses the two ways it can be too big. It reports
// whether it has already written a response.
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
		// go over as decimal strings (contract 8).
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

Add `"fmt"`, `"strconv"` and `"github.com/jhunthrop/foreversixty/api/internal/httpx"` to `run.go`'s imports (`httpx` is already there; `context` and `time` are too).

- [ ] **Step 6: Fix the mount tests**

`TestTheRunRouteIsNotMountedWithoutJobsAndAccounts` in `handler_test.go` and `api/internal/server/sims_test.go` both build a `Service` without a planner. In `handler_test.go`, add a case that a service with jobs and accounts but no planner also 404s:

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

with `"github.com/jhunthrop/foreversixty/api/internal/jobs"` imported. Then in `api/internal/server/sims_test.go`, find the service literal at line 79's `withRun` and add `Planner: ...` to it — read the file first and give it a stub that satisfies `sims.Planner`:

```go
// planNothing satisfies sims.Planner for the mount check, which never
// plans anything.
type planNothing struct{}

func (planNothing) Plan(context.Context, simapi.SimRequest) (simapi.PlanSummary, error) {
	return simapi.PlanSummary{}, nil
}
```

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

The server overwrites bulk.cap with its own lane's, then counts the
expansion through a plan-only call — no iterations — and answers 400
cap_exceeded with the cap and the count. sim/bulk cannot be imported
here (it reads sim/internal/simdb, which imports the engine), so the
count crosses as JSON from the binary; the interface is declared by the
consumer so the handler is testable without one.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

### Task 7: The `too_large` estimate

**Files:**
- Modify: `api/internal/sims/job.go` (`BulkRunTimeout`, `timeoutFor`)
- Modify: `api/internal/sims/simdep.go` (the rate constants, `estimateSec`)
- Modify: `api/internal/sims/run.go` (`checkSize`)
- Test: `api/internal/sims/run_test.go`, `api/internal/sims/simdep_test.go`

**Interfaces:**
- Consumes: `simapi.PlanSummary.IterationsTotal` (contract 1.8).
- Produces:
  - `const BulkRunTimeout = 14 * time.Minute`
  - `func timeoutFor(req simapi.SimRequest) time.Duration`
  - `func estimateSec(iterations int) int`

- [ ] **Step 1: Write the failing test**

Add to `api/internal/sims/simdep_test.go`:

```go
func TestTheEstimateIsIterationsOverTheJobsMeasuredRate(t *testing.T) {
	for _, c := range []struct {
		iterations int
		want       int
	}{
		{0, 0},
		{1, 1},                 // rounded up: a run is never estimated at no time at all
		{nativeRate, 1},        //
		{nativeRate * 60, 60},  // a minute of the job's four CPUs
		{20_033_000, 4007},     // 20,000 combinations at normal precision
	} {
		if got := estimateSec(c.iterations); got != c.want {
			t.Errorf("estimateSec(%d) = %d, want %d", c.iterations, got, c.want)
		}
	}
}

func TestTheBulkTimeoutFitsInsideTheCloudRunTaskTimeout(t *testing.T) {
	// api/README.md creates sim-run with --task-timeout 15m. The
	// in-process bound has to leave room for start-up and the two
	// writes at the end, or the platform kills the job mid-write and
	// the row never leaves "running".
	const taskTimeout = 15 * time.Minute
	if BulkRunTimeout >= taskTimeout {
		t.Fatalf("BulkRunTimeout %s, want less than the job's %s", BulkRunTimeout, taskTimeout)
	}
	if RunTimeout > BulkRunTimeout {
		t.Fatalf("a plain run (%s) may not outlast a bulk one (%s)", RunTimeout, BulkRunTimeout)
	}
}
```

Add to `api/internal/sims/run_test.go`:

```go
func TestABulkRunPastTheJobsTimeoutIsRefusedWithItsEstimate(t *testing.T) {
	h := newHarness(t)
	h.premium.premium = true
	h.planner.summary = simapi.PlanSummary{
		Kind: simapi.KindGear, Combinations: 19_000,
		Cap: simapi.Caps[simapi.LaneServer], IterationsTotal: 20_033_000,
	}

	res := h.json(http.MethodPost, "/v1/sims/run", bulkBody(t))
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("status %d, want 400", res.StatusCode)
	}
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
	if env.Error.Code != "too_large" {
		t.Fatalf("code %q, want too_large", env.Error.Code)
	}
	if env.Error.Fields["estimate_sec"] != "4007" ||
		env.Error.Fields["budget_sec"] != strconv.Itoa(int(BulkRunTimeout.Seconds())) {
		t.Fatalf("fields %+v", env.Error.Fields)
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
		Weights:   &simapi.WeightsSpec{Stats: []string{"strength", "crit"}, Reference: "crit"},
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

Run: `go test -p 1 ./api/internal/sims/... -run 'TestTheEstimate|TestTheBulkTimeout|TestABulkRunPastTheJobsTimeout|TestAWeightsRun' -v`
Expected: FAIL to build — `undefined: estimateSec`, `undefined: nativeRate`, `undefined: BulkRunTimeout`.

- [ ] **Step 3: Write the rate and the estimate**

Add to `api/internal/sims/simdep.go`, after the iteration constants:

```go
// The native engine's measured throughput, and the shape of the job it
// runs in. These two numbers are the whole of the submit-time estimate.
//
// nativeIterationsPerCPUSecond is the figure job.go's RunTimeout was
// sized from and is stated with: ten thousand iterations of a
// three-minute fight is about eight CPU-seconds natively, so 1,250 an
// iteration-second per core. simJobCPUs is what api/README.md creates
// the sim-run job with (`--cpu 4`); the two have to be changed
// together, and TestTheEstimateIsIterationsOverTheJobsMeasuredRate
// pins the product.
//
// They are constants rather than configuration because they are a
// property of the engine and the job definition, not of a deployment:
// an operator who changes the job's CPU count is editing the README
// line and this one in the same commit.
const (
	nativeIterationsPerCPUSecond = 1250
	simJobCPUs                   = 4
	nativeRate                   = nativeIterationsPerCPUSecond * simJobCPUs
)

// estimateSec is how long the sim-run job needs to complete iterations
// iterations, rounded up: a run that would take a fraction of a second
// still takes some.
func estimateSec(iterations int) int {
	if iterations <= 0 {
		return 0
	}
	return (iterations + nativeRate - 1) / nativeRate
}
```

- [ ] **Step 4: Give a bulk run its own bound**

In `api/internal/sims/job.go`, add under `RunTimeout`:

```go
// BulkRunTimeout bounds one bulk or weights run: every stage of the
// ladder, not one sim. api/README.md creates the sim-run job with
// --task-timeout 15m, and a job killed by the platform dies between the
// bucket write and the row write, leaving the page polling "running"
// forever — so the in-process bound stops a minute short of it and
// fails the row on its way out.
const BulkRunTimeout = 14 * time.Minute

// timeoutFor is the bound one request's run gets. A plain run keeps the
// tighter one: ten minutes for a single sim is already an outlier worth
// failing.
func timeoutFor(req simapi.SimRequest) time.Duration {
	if req.Kind() == simapi.KindRun {
		return RunTimeout
	}
	return BulkRunTimeout
}
```

and in `Run`, replace `runCtx, cancel := context.WithTimeout(ctx, RunTimeout)` with:

```go
	runCtx, cancel := context.WithTimeout(ctx, timeoutFor(req))
```

- [ ] **Step 5: Refuse the oversized request**

In `api/internal/sims/run.go`'s `checkSize`, after the cap block and before `return false`:

```go
	budget := int(BulkRunTimeout.Seconds())
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
Expected: PASS.

- [ ] **Step 7: Commit**

```bash
git add api/internal/sims/simdep.go api/internal/sims/simdep_test.go \
        api/internal/sims/job.go api/internal/sims/run.go api/internal/sims/run_test.go
git commit -m "$(cat <<'EOF'
feat(api): a run past the job's timeout is refused at submit, with the estimate

The estimate is the plan's total iterations over the job's measured rate
— 1,250 an iteration-second per core, four cores — against a 14-minute
in-process bound that stops short of Cloud Run's 15-minute task timeout.
A weights request has no cap but still has a budget.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

### Task 8: Stage progress through the job

**Blocked on:** the `sim` lane's widened `runner.Progress` (`func(p api.Progress)`) and the three new fields on `api.Progress` (contract 2, Task 1 step 3). Until both land, this task will not compile. Everything it tests on this side uses a local stub runner, not `runner.Fixture`, so nothing else waits on the `sim` lane's fixture work.

**Files:**
- Modify: `api/internal/sims/store.go` (`Progress`, `Advance`)
- Modify: `api/internal/sims/job.go` (the callback)
- Test: `api/internal/sims/job_test.go`, `api/internal/sims/store_test.go`, `api/internal/sims/validate_test.go`

**Interfaces:**
- Consumes: `simapi.Progress{IterationsRun, DPS, Stage, CombosDone, CombosTotal}`; `runner.Progress func(api.Progress)`.
- Produces:
  - `Progress` gains `Stage int \`json:"stage"\``, `CombosDone int \`json:"combos_done"\``, `CombosTotal int \`json:"combos_total"\``
  - `func (s *Store) Advance(ctx context.Context, id string, p simapi.Progress) error`

- [ ] **Step 1: Write the failing test**

Add to `api/internal/sims/job_test.go`:

```go
// bulkRunner answers the way forever-sim does for a bulk request: it
// streams a tick per stage, then returns a ranked result. The binary
// detects the kind and runs sim/bulk's plan-rank loop itself, so the
// job hands it the request whole and reads what comes back — which is
// exactly what this stands in for.
type bulkRunner struct{ got simapi.SimRequest }

func (b *bulkRunner) Run(_ context.Context, req simapi.SimRequest,
	onProgress runner.Progress) (simapi.SimResult, error) {
	b.got = req
	if onProgress != nil {
		onProgress(simapi.Progress{
			IterationsRun: 1000, DPS: simapi.Estimate{Mean: 1000},
			Stage: 1, CombosDone: 40, CombosTotal: 96,
		})
		onProgress(simapi.Progress{
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
	// run: a bulk or weights request there would sim fifty parses
	// dozens of times each and blow the job's budget.
	h := newHarness(t)
	engine := &runner.Fixture{}
	d := ValidateDeps{
		Store: h.store, Top: fixedParses(t), Engine: engine, Build: CombatantBuilder{},
		Scores: noScores{}, Log: slog.New(slog.NewTextHandler(io.Discard, nil)),
	}
	_ = Validate(t.Context(), d, []string{"warrior-fury"}, "raids-1", testEngine)
	for _, req := range engine.Asked() {
		if req.Kind() != simapi.KindRun {
			t.Fatalf("the validation job asked for a %s run: %+v", req.Kind(), req)
		}
		if req.Bulk != nil || req.Weights != nil {
			t.Fatalf("the validation job built a bulk request: %+v", req)
		}
	}
}
```

Read `validate_test.go` first and reuse whatever it already has for `Top`, `Build` and `Scores` — the names `fixedParses` and `noScores` above are placeholders for that file's own existing fakes; substitute the real ones rather than adding new ones.

- [ ] **Step 2: Run it and watch it fail**

Run: `go test -p 1 ./api/internal/sims/... -run 'TestABulkJobStreams|TestAPlainRunReportsNoStages|TestTheValidationJobStill' -v`
Expected: FAIL — `p.Stage undefined`, and a type error on the `bulkRunner`'s `onProgress` signature until the `sim` lane's widened `runner.Progress` is in.

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
func (s *Store) Advance(ctx context.Context, id string, p simapi.Progress) error {
	_, err := s.Pool.Exec(ctx,
		`update sims set state = $2, iterations = $3, dps_mean = $4,
		   stage = $5, combos_done = $6, combos_total = $7
		 where id = $1 and state in ($8, $9)`,
		id, StateRunning, p.IterationsRun, p.DPS.Mean,
		p.Stage, p.CombosDone, p.CombosTotal, StateQueued, StateRunning)
	if err != nil {
		return fmt.Errorf("sims: advance %s: %w", id, err)
	}
	return nil
}
```

- [ ] **Step 4: Pass the whole tick through the job**

In `api/internal/sims/job.go`, replace the `Engine.Run` callback:

```go
	res, err := d.Engine.Run(runCtx, req, func(p simapi.Progress) {
		// A progress write that fails is logged and the run carries
		// on: the figure on the page is a courtesy, the result is not.
		if err := d.Store.Advance(ctx, simID, p); err != nil {
			d.logger().Error("sims", "op", "progress", "sim", simID, "err", err)
		}
	})
```

- [ ] **Step 5: Fix the two existing `Advance` callers in tests**

`TestAServerRunWalksQueuedThenRunningThenDone`, `TestAdvanceCannotReviveAFailedRun` and `TestAdvanceCannotReopenAFinishedRun` in `store_test.go` each call `h.store.Advance(ctx, id, 1500, 1000)`. Rewrite each as:

```go
	if err := h.store.Advance(t.Context(), "dddddddddddd", simapi.Progress{
		IterationsRun: 1500, DPS: simapi.Estimate{Mean: 1000},
	}); err != nil {
```

keeping each test's own id and expectations.

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
Progress rows gain the three stage fields; a plain run reports zeros.
Asserts the nightly validation job still builds plain runs only.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

### Task 9: `GET /v1/specs` carries the reference stat

**Blocked on:** the data lane adding `reference_stat` to `data/curated/specs.json` and regenerating `sim/specs/specs.go` so `specs.Spec` has `ReferenceStat string \`json:"reference_stat"\`` (contract 1.4, Task 1 step 6).

**Files:**
- Modify: `api/internal/sims/specs.go`
- Test: `api/internal/sims/specs_test.go`

**Interfaces:**
- Consumes: `specs.ByKey map[string]specs.Spec` with `ReferenceStat`.
- Produces: `SpecFidelity` gains `ReferenceStat string \`json:"reference_stat"\``.

- [ ] **Step 1: Write the failing test**

Add to `api/internal/sims/specs_test.go`:

```go
func TestEverySpecCardNamesItsReferenceStat(t *testing.T) {
	h := newHarness(t)
	// A measured row and an unmeasured one both carry it: the weights
	// page reads the reference off the card before anything has been
	// simmed.
	if err := h.store.PutSpec(t.Context(), SpecFidelity{
		Spec: "warrior-fury", Parses: 50, MedianGap: ptr(0.02), EngineVersion: testEngine,
	}); err != nil {
		t.Fatal(err)
	}
	cards, err := h.store.Specs(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	seen := 0
	for _, c := range cards {
		want := specs.ByKey[c.Spec].ReferenceStat
		if want == "" {
			t.Fatalf("%s has no reference_stat in the generated spec list", c.Spec)
		}
		if c.ReferenceStat != want {
			t.Errorf("%s: reference_stat %q, want %q", c.Spec, c.ReferenceStat, want)
		}
		seen++
	}
	if seen == 0 {
		t.Fatal("no cards at all")
	}
}
```

Reuse the file's existing pointer helper if it has one instead of `ptr`; read `specs_test.go` first. Add `"github.com/jhunthrop/foreversixty/sim/specs"` to its imports.

- [ ] **Step 2: Run it and watch it fail**

Run: `go test -p 1 ./api/internal/sims/... -run TestEverySpecCardNamesItsReferenceStat -v`
Expected: FAIL — `c.ReferenceStat undefined`.

- [ ] **Step 3: Carry it on the card**

In `api/internal/sims/specs.go`, add to `SpecFidelity`, under `Spec`:

```go
	// ReferenceStat is the stat the weights tool normalises to 1.0 for
	// this spec (contract 1.4). It is the data lane's figure, read from
	// the generated spec list rather than stored: it is a property of
	// the spec, not of a measurement, and a row measured before the
	// data lane changed it must not report the old one.
	ReferenceStat string `json:"reference_stat"`
```

In `Store.Specs`, fill it on the seeded card:

```go
	for _, spec := range specs.All {
		if spec.Role != roleDPS {
			continue
		}
		byspec[spec.Spec] = SpecFidelity{
			Spec: spec.Spec, ReferenceStat: spec.ReferenceStat,
			State: SpecUnsupported, WorstActions: []WorstAction{},
		}
	}
```

(this replaces the `for _, spec := range DPSSpecs()` loop, which had only the key), and on every measured row, immediately after the `byspec[f.Spec] = f` line is reached — replace that line with:

```go
		// A measured row for a spec that is no longer in the list is
		// still shown: a rename should be visible, not silent. Its
		// reference stat is simply blank, because the list no longer
		// has one for it.
		f.ReferenceStat = specs.ByKey[f.Spec].ReferenceStat
		byspec[f.Spec] = f
```

and delete the old comment above it. Add `"github.com/jhunthrop/foreversixty/sim/specs"` to `specs.go`'s imports.

- [ ] **Step 4: Run the tests**

Run: `go test -p 1 ./api/internal/sims/...`
Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git add api/internal/sims/specs.go api/internal/sims/specs_test.go
git commit -m "$(cat <<'EOF'
feat(api): spec cards carry the weights tool's reference stat

Read from the generated spec list, not stored: it is a property of the
spec, so a card measured before the data lane changed it must not report
the old one.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

### Task 10: Wiring, OpenAPI and the operator documentation

**Blocked on:** the `sim` lane's `Plan` method on `runner.Native` and `runner.Fixture` (contract 4, Task 1 step 3). Until both exist, `simEngine` cannot return a `sims.Engine`.

**Files:**
- Modify: `api/cmd/api/main.go`
- Modify: `api/openapi.yaml`
- Modify: `api/README.md`
- Test: `api/internal/server/openapi_test.go` (no change expected — confirm it still passes)

**Interfaces:**
- Consumes: `sims.Engine` (Task 6); `runner.Native.Plan`, `runner.Fixture.Plan`.
- Produces: a deployment that mounts `POST /v1/sims/run`.

- [ ] **Step 1: Give the service a planner**

In `api/cmd/api/main.go`, change `simEngine`'s signature and comment:

```go
// simEngine is what every simulator job and the submit handler use: the
// real binary when the image carries one, and the checked-in fixture
// when it does not, so a deployment without the artifact still answers
// instead of failing. It runs sims and it counts bulk requests without
// running them (forever-sim -plan), which is why one value serves both
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
Expected: success. A failure naming `Plan` means the `sim` lane's half is not merged yet; stop and wait rather than adding a shim.

- [ ] **Step 3: Update the OpenAPI document**

In `api/openapi.yaml`:

Add `kind` and `headline` to `SimRow` (around line 237):

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
            Fury'", "Crit 1.00 · Agility 0.87".
        engine_version: { type: string }
        created_at: { type: string, format: date-time }
        title: { type: string }
```

Add `reference_stat` to `SpecFidelity`, under `spec`:

```yaml
        reference_stat:
          type: string
          description: The stat the weights tool normalises to 1.0 for this spec.
```

Add the `kind` parameter to `GET /v1/sims` (the `listMySims` parameter list):

```yaml
        - name: kind
          in: query
          schema: { type: string, enum: [run, gear, talents, drops, weights] }
          description: Narrows the list to one tool. Omitted or empty is every kind; an unknown value is 400.
```

Add the two refusals to `POST /v1/sims/run`'s responses, replacing its `'400'` line:

```yaml
        '400':
          description: >-
            The request cannot be run. error.code is "invalid" for a
            malformed envelope, "cap_exceeded" when the expansion is past
            the server lane's cap (error.fields carries cap and
            combinations as decimal strings), or "too_large" when the
            planner estimates the run past the job's timeout
            (error.fields carries estimate_sec and budget_sec).
```

Add the stage fields to `GET /v1/sims/{id}/progress`'s data properties, under `dps`:

```yaml
                          stage: { type: integer, description: Which stage of a bulk run is going; 0 for a plain run. }
                          combos_done: { type: integer }
                          combos_total: { type: integer }
```

- [ ] **Step 4: Run the document's own test**

Run: `go test ./api/internal/server/...`
Expected: PASS (`TestOpenAPIListsEveryRoute` checks paths and schemas; no new path was added).

- [ ] **Step 5: Update the README**

In `api/README.md`, under "The simulator's Cloud Run jobs", after the paragraph beginning "`sim-run` is executed by the API", add:

```markdown
The `--cpu 4 --task-timeout 15m` on `sim-run` is load-bearing and is not
just a ceiling. A Top Gear, Droptimizer, talent-compare or stat-weights
submit is sized before it is queued: the API asks the binary to expand
the request without running it, multiplies the precision ladder's total
iterations by the engine's measured rate — about 1,250 iterations a
CPU-second, times those four CPUs — and refuses anything past
`sims.BulkRunTimeout` (14 minutes, one minute inside the task timeout)
with `400 too_large` and the estimate. Changing the job's CPU count means
changing `simJobCPUs` in `api/internal/sims/simdep.go` in the same
commit, or every estimate is wrong. Jobs are created by hand, so nothing
enforces this but this paragraph.
```

In the route table further down, add a line beside the other sim routes:

```markdown
| `GET /v1/sims?mine=1&kind=` | The caller's own sims, newest first, optionally narrowed to one tool (`run`, `gear`, `talents`, `drops`, `weights`). Each row carries its kind and a composed one-line headline. An unknown kind is 400. |
```

No new environment variable: the rate and the CPU count are constants, not configuration, because they describe the engine and the job definition rather than a deployment.

- [ ] **Step 6: Run the API's tests once more**

Run: `go test -p 1 ./api/...`
Expected: PASS. Re-run once on a `sims_pkey` collision.

- [ ] **Step 7: Commit**

```bash
git add api/cmd/api/main.go api/openapi.yaml api/README.md
git commit -m "$(cat <<'EOF'
feat(api): wire the planner, and document the kinds and the job's budget

simEngine now returns a sims.Engine — one value that both runs sims and
counts bulk requests without running them — so the service can size a
submit. OpenAPI gains kind, headline, reference_stat, the two refusals
and the stage progress fields; the README says why the job's --cpu 4 is
load-bearing.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
EOF
)"
```

---

## Self-review against the contract

**Section 8 coverage.** `POST /v1/sims/run` cap and estimate: Tasks 6 and 7. `sims.kind` and migration `0014_sim_kinds`, `Store.Queue` from `req.Kind()`: Task 2. `GET /v1/sims/<id>` unchanged: nothing to do — the result blob is stored whole by `Save`/`Finish` already, and Task 8's job test asserts `Combos`, `Equipped` and `Stages` survive the round trip. `GET /v1/sims?mine=1&kind=` with `kind` and `headline`: Tasks 3, 4, 5. Progress rows: Task 8. `GET /v1/specs` `reference_stat`: Task 9.

**Design section 10.2 and 10.3 coverage.** "The job row records the kind": Task 2. "The `sim-run` Cloud Run job runs the native planner loop and streams stage progress": Task 8. "Bulk requests get the 4-CPU job at a 15-minute timeout; a request the planner estimates past that is refused at submit with the estimate": Task 7 and Task 10's README paragraph. "Saved sims of every kind are public at `/sim/<id>`": already true — `get` has no kind in it.

**Design section 11 coverage.** "The API test suite covers the kind column, submit-time cap refusal and job progress for a bulk request with the fixture engine": Tasks 2, 6 and 8. "The nightly validation job is unchanged": Task 8's `TestTheValidationJobStillRunsPlainSims`.

---

## Where the contract was ambiguous or wrong for this lane

These are recorded here as well as in the Task 1 amendment, so an executor reading only the plan knows which sentences were inferred.

1. **Section 8 names `POST /v1/sims` for the premium submit; that route is the browser-result save.** The premium submit is `POST /v1/sims/run`. Task 1 step 5 corrects it.
2. **`bulk.Expand` is unreachable from the api module.** It reads `sim/internal/simdb`, which imports the engine, and `sim/internal/...` is closed to other modules besides. The count crosses as JSON from `forever-sim -plan` (Task 1 steps 2 and 3).
3. **The envelope's `error.fields` is `map[string]string`.** `{"cap": 20000}` cannot be a number without widening `httpx.ErrorBody` and every handler's tests with it; the numbers go over as decimal strings.
4. **The `too_large` budget has no number in the contract.** It is `BulkRunTimeout`, 14 minutes, derived from the README's `--task-timeout 15m`; the rate is job.go's own measured "10,000 iterations ≈ 8 CPU-seconds" times the job's four CPUs.
5. **The server cap and the job timeout disagree.** 20,000 combinations at `fast` precision is about 7 million iterations, roughly 23 minutes at the measured rate — so `too_large`, not `cap_exceeded`, is the binding refusal for almost every large request. Both checks are kept, in that order, and the cap is left at the contract's number.
6. **`Substitution` has no name for an item,** so "+41 DPS from Vis'kag" is uncomposable from a stored result. Task 1 step 4 extends `Name` to items.
7. **The drop source's display name has no home in the envelope.** `sourceLabel` derives it from the origin id's last segment (`drop:raid:mc:ragnaros` → "Ragnaros"), which holds because those ids are slugs of the source names by `loot.json`'s construction.
8. **The headline examples only show one substitution.** A Top Gear winner can swap several slots; this plan names the first and counts the rest ("… and 2 more"), and names the empty cases ("no combinations", "no upgrades", "no weights") which the contract does not.
9. **Progress has nowhere to be stored.** Migration 0014 adds `stage`, `combos_done` and `combos_total` alongside `kind`; the contract's section 8 mentions only `kind`.
10. **`headline` is a stored column, not a computed one.** Composing it on read would detoast a few hundred kilobytes of result JSON per row, a hundred rows to a page.
