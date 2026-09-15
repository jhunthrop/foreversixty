# Forever Sixty — combat simulator design

Date: 2026-09-14. Status: approved in conversation, section by section. Research behind it:
`research/07-simulator.md` (WoWSims and SimulationCraft internals, measured throughput,
Raidbots' product surface).

The simulator is the fourth pillar of the one-stop site after the planner, the addon, and the
logs. It answers three questions in order: what is my DPS and why (Quick Sim), which of the
gear I own is best (Top Gear), and which boss to run for my next upgrade (Droptimizer). The bar
the user set: a world-class UX on a proven engine, free to run in the browser, with server
compute as the site's first paid feature.

## Decisions

| Question | Decision |
|---|---|
| Engine | WoWSims, used upstream rather than forked adversarially: a Forever repository in the WoWSims organisation if they accept it, otherwise hosted by us under the same MIT terms with the link back they ask for. Consumed as a Go module; our value is the UX, the data, the integration, and the validation loop. |
| SimulationCraft | Rejected: no vanilla class code, server-only, GPL-3 makes a browser build "conveying", C++ beside a Go and TypeScript codebase. |
| Own engine | Rejected: the user does not want to build an engine, and the research sized it at 25 to 37 engineer-weeks against 11 to 15 for the WoWSims route. |
| Scope at launch (Nov 4) | Quick Sim for DPS specs. Tanks and healers are research problems and stay out of the first cut. Top Gear and Droptimizer follow for the Dec 9 raids. |
| Class models | Written on top of the vanilla implementations, since Forever keeps vanilla combat, and corrected against beta logs. Two specs validated for launch (Fury Warrior, Frost Mage); the rest added in rankings-population order through December. |
| Rotations | One default action priority list (APL) per spec, written by us and checked against what top parses actually cast. APLs are stored as data from day one so sharing and editing are later features, not a redesign. The APL builder UI is deferred. |
| UI | Our own, in the site's design system, built from the planner, gear picker, and report components. None of the WoWSims TypeScript UI ships. |
| Where it runs | Browser for free, unlimited, no queue: the engine as WebAssembly in a worker pool. Server for premium: the same engine natively as a Cloud Run job, gated on a signed-in account with a premium flag. One request and one result format for both lanes. |
| Character sources | Armory by Battle.net sign-in, addon export by companion or paste, a planner build by link, a logged fight from a report or ranking. One character model behind all four, shared with the planner. |
| Results | The sim emits the logs engine's event stream, so the summary code that renders a real fight renders a sim. Compare mode puts a real fight beside the sim per ability and per buff. |
| Integration | The sim is mostly ambient: an execution score on every ranked fight, live DPS in the planner, a one-button landing state for members, and later tooltip deltas, the biggest-lever hint, guild execution, and results pushed to the addon (section 4.6). |
| Validation | A nightly job sims the top parses per spec from rankings and publishes each spec's fidelity on a public support page. "Validated" means a median gap under 5% on the top 50 parses. |
| Iterations | 3,000 by default with a live-updating estimate; a precision toggle for 10,000. Enough for ±0.25% and ±0.14% respectively; direct comparisons use paired seeds. |
| Data | The existing pipeline generates the engine's database (items, enchants, sets, talents, consumables) from the Forever client tables; nothing scraped at runtime. |

## 1. Architecture

```
wowsims/forever (Go, MIT) ── protobuf API ──┬─▶ sim.wasm ─▶ browser worker pool (free lane)
                                            └─▶ native binary ─▶ Cloud Run job (premium lane)
data/ pipeline ─▶ engine database (items, enchants, sets, talents, consumables) ─▶ both lanes
web/ sim pages (Svelte islands) ◀── character model ◀── Armory | addon export | planner build | logged fight
                                ◀── result (protobuf → logs engine events → summary JSON) ─▶ report components
api/ ── characters, saved sims, premium lane dispatch, validation job ─▶ Postgres, R2
```

Budgets, measured in CI:

| Path | Budget |
|---|---|
| Quick Sim, 3,000 iterations, 3-minute fight, 4-worker laptop | first estimate within 1 s, final within 8 s |
| Quick Sim, premium server lane, 10,000 iterations | final within 3 s of dispatch |
| Sim page first paint | LCP within the site's 1.6 s mobile budget; the engine loads after paint, 4 MB gzipped, cached immutably per engine version |
| Compare mode | no extra sim; the fight's summary is the one the report page already serves |
| Validation job | 50 parses × 9 specs × 3,000 iterations nightly, under 20 minutes on 4 CPUs |

Measured baseline, single core, 300-second fight, Fury Warrior: native Go 1,218 iterations per
second, WebAssembly 155. The browser number is the design constraint: a laptop with four workers
runs 3,000 iterations in about 5 seconds and 10,000 in about 16. The estimate updates as
iterations accumulate, so the page is never waiting on a spinner.

## 2. Engine

### 2.1 Repository and relationship to WoWSims

A repository `forever` in the WoWSims organisation, created by them at our request, or by us at
`github.com/jhunthrop/wowsims-forever` if they decline, in which case it stays mergeable and
carries the link back the MIT README asks for. The user's stance: not adversarial; contribute
back. The site imports the repository as a Go module pinned by version and builds `sim.wasm`
from it in CI; the same version builds the native binary in the premium lane's image.

Starting point: `wowsims/classic` at its current head. It is dormant on class code, which is
fine: the engine core, APL system, metrics, bulk sim, stat weights, and the WebAssembly and
worker plumbing are what we keep, and the class code is rewritten for Forever regardless.

### 2.2 Core changes for Forever's rules

Confirmed rules from the Deep Dive panel and their engine cost, from reading the code:

| Rule | Where | Cost |
|---|---|---|
| Spell, melee, and ranged hit are one stat; likewise crit | `sim/core/stats` enum and `proto/common.proto`; 159 mechanical call sites in `sim/`, none in the UI | days |
| Expertise reduces parry and dodge | already modelled in the attack table from the later-expansion lineage; dormant in vanilla, live for Forever | none |
| Bonus healing carries one third as bonus damage | one entry in the existing stat-dependency table | hours |
| Caster weapons grant spell damage | item data; the stat exists | none |
| Weapon skill kept, less per item | fully modelled; item data | none |
| Talent trees: seven rows, milestones at 11, 16, 21, 31; buff talents baseline | per-class talent proto and tree JSON regenerated from the client; per-class `talents.go` partially rewritten, roughly half survives | 2 to 3 days per class |
| Reworked racials, two active and two passive per race | one file rewritten | days |
| New baseline abilities per class | one small file each in the existing one-ability-per-file pattern | hours each |
| Re-itemised world; biome- and creature-type trinkets | item effects and set files as content; one new engine concept, an encounter environment (biome, target creature type) that conditional effects read | a week for the concept, content ongoing |
| `SpellScaling` absent from the Classic-lineage tables | spell coefficients stay the vanilla convention with per-spell overrides, as they are today | none |
| Periodic damage can critically strike (vanilla's could not) | the engine already carries the machinery: `Dot.OutcomeTickPhysicalCrit`, the `CritTicks` and `ResistedCritTicks` metrics, and `crit_ticks`/`resisted_crit_ticks` in `TargetedActionMetrics`. What is missing is the magic-school equivalent of that one function, plus switching the 64 call sites that apply the plain `OutcomeTick` to the critting variant where Forever says a DoT crits, and the crit multiplier the periodic case uses | one new outcome function; a call-site sweep per spec |

The `spell_mod.go` declarative modifier system from the Season of Discovery repository is
cherry-picked, since Forever's reworked talents are mostly "these spells cost or crit or hit
differently", which it expresses as data rather than closures.

### 2.3 Class models

Each spec is a hand-written Go package in the existing pattern: one ability per file, constants
in code, talents applied in one function. Forever's numbers come from the client tables through
the pipeline, so a regenerated constants file per class carries base damage, costs, cooldowns,
and durations, and the ability file references it rather than a literal. Behaviour that no table
holds, which is procs, internal cooldowns, coefficients, and interactions, is written from the
tooltip and the beta and marked `unconfirmed` until the validation job clears it.

Order: Fury Warrior and Frost Mage first, as the two most-played DPS specs in every vanilla
launch and the two with the most different mechanics (melee attack table and rage; caster with
resistances and mana). Then, in rankings-population order as the beta reveals it, the remaining
DPS specs: Rogue, Hunter, Warlock, Shadow Priest, Balance and Feral Druid, Elemental and
Enhancement Shaman, Retribution Paladin, Fire and Arcane Mage, Arms Warrior.

### 2.4 Rotations

Each spec ships one default APL, a priority list with conditions in the engine's existing APL
schema, stored as a JSON document in `data/curated/apl/<spec>.json` with sources, like every
other curated fact. The validation job's cast-frequency comparison against top parses is the
evidence for changing one. Player-authored and shared APLs are a later feature over the same
schema.

### 2.5 Results as logs events

The engine's result is a protobuf: DPS distribution, per-action metrics, aura uptimes, resource
timelines, and one iteration's cast log. A Go adapter in our repository (`sim/` package under
`logs/`) turns the cast log and metrics into the logs engine's `event.Event` stream and the same
per-fight summary the report page renders. Sized at one to two weeks. It is what makes compare
mode a layout rather than a second renderer, and it means every improvement to the report views
reaches the sim for free.

## 3. Data

The pipeline gains a `simdb` output: the engine's item, enchant, gem-free, set, consumable, and
talent database in its protobuf schema, generated from the same client tables the planner uses,
with our icons and names. Wowhead's Forever data environment exists and matches the shape the
engine's generator already parses; it and wago.tools' DB2 export both carry the beta client's
tables within a day of Sept 17. The pipeline's fetch step already reads DB2 from wago.tools.

A per-class constants file (section 2.3) is generated alongside, keyed by spell id, so a beta
patch that changes a number is a pipeline run and a constants regeneration, not a code edit.

## 4. The product

### 4.1 Sim page, `/sim`

- **Character strip**: name, race, class, spec, level, gear as the planner's slot grid with
  tooltips, and a source pill: "Armory, 2 hours ago", "Addon export, just now", "Build from
  planner", "From this fight". The source switcher offers the four sources; Armory needs
  sign-in, the others do not.
- **Settings bar**: fight length (default 3 minutes, 1 to 8), targets (default 1, up to 10),
  execute phase on, buffs and consumables as a named preset ("raid-buffed", "solo", "custom"
  with the full list), the rotation shown as "Default for Fury" with a link to what it does.
- **Run**: one button. The DPS figure appears within a second and refines as iterations land,
  with its error range and a progress line. Precision toggle for 10,000 iterations. Premium
  accounts see "Run on our servers" beside it.
- **Results**: a one-sentence summary ("Bloodthirst and white hits are 61% of your damage;
  Flurry uptime 78%"), then the report page's components: damage breakdown per ability, buff
  and debuff uptimes, resource timeline, a sample cast timeline from one iteration, and the
  DPS distribution. Everything is the same component the logs report uses.
- **Spec not supported**: instead of a number, the spec's card from the support page: its
  state, what is being validated, and when it is expected.
- **Phone**: the strip collapses to a row, settings to a sheet, results stack.

### 4.2 Compare mode

Reached from a fight in a report ("Sim this fight"), a rankings row, or the sim page's picker
over the signed-in player's own fights. The character is loaded from the fight's recorded gear,
talents, and buffs, the sim runs, and the results show two columns, simulated and actual, per
ability and per buff, with the difference highlighted and explained in words: "Bloodthirst cast
41 times, the sim expects 52", "Flurry uptime 61% against 78%". The fight's summary is the one
the report already serves; no second parse.

### 4.3 Saved sims, `/sim/<id>`

A public, shareable result with an unfurl card, stored like a build: the request, the engine
version, and the summary. Signed-in players keep a history and can name entries. Ids are the
same 12-character base32 as reports.

### 4.4 Spec support page, `/sim/specs`

One card per spec: state (validated, in progress, not yet), the fidelity figure and the parse
count behind it, the abilities with the largest gaps, the engine version, and the date. The
answer to "can I trust this" lives on the site.

### 4.6 Integration surfaces

The sim is mostly not a page the member visits. Ordered by launch phase.

At launch (S1 and S2):

- **Execution score.** The nightly job (section 6) also sims every ranked fight of a signed-in
  member with the gear, talents, and buffs recorded for that fight. Each report row, character
  page, and rankings row shows "92% of what your gear can do" beside the parse percentile, with
  compare mode one click away. The score is stored on the fight-metrics row and served by the
  same rankings and character reads; unsigned members see it for fights whose spec is
  validated, computed on demand in the browser.
- **Live DPS in the planner.** A talent or slot change in the planner runs a 500-iteration sim
  in the browser pool and updates a DPS estimate with its error in under a second; a "Sim this
  build" control opens the full results. The planner is where gear decisions are made, the sim
  page where they are explained.
- **Signed-in landing state.** `/sim` for a member shows their characters with last-logout gear
  and one button per character. No form is shown unless the member opens one.

With Top Gear (S3):

- **Gear tooltip deltas.** Any item tooltip on the site, a dungeon drop, a ranking row's
  trinket, a planner candidate, shows "+42 DPS for you" from a paired single-swap sim against
  the current gear, run in the browser on hover and cached per item for the session.
- **Biggest lever.** After the main sim, the browser runs a handful of paired variants (a missing
  consumable, an alternative in each talent milestone, the weapon enchant) and the results name
  the one that helps most.
- **Guild execution.** The guild page shows execution scores per raider per fight, from the
  same stored column.
- **Result to the addon.** Expected DPS and cast counts pushed through the addon inbox, so the
  in-game meter carries an expected column beside the actual one.

### 4.7 Cross-pillar features

What each pillar gains from the sim, and what the sim gains from each. Ordered by phase.

At launch (S1 and S2):

- **Execution leaderboard.** A second sort order on every rankings page: by execution score
  rather than DPS, so a player in dungeon gear competes with a full raider on how well they
  played. One column, one sort; the fairest ranking on offer anywhere.
- **Execution score at fight close.** For members, the score is computed when the fight closes
  (three seconds of native compute per fight on the ingest path's job runner), not nightly, so a
  report carries it the moment the fight ends. The nightly job remains the fidelity source.
- **Build unfurls with DPS.** A shared planner link's card carries the build's simmed DPS and the
  engine version it was simmed on.

With Top Gear (S3):

- **Fight-shaped encounters.** Per boss, an encounter profile generated from the logs corpus:
  median kill length, target counts over time, the execute window, and the downtime players
  actually had. A sim of "this boss as raids fight it", offered beside the plain single-target
  fight, refreshed weekly from the corpus. Raidbots' hand-made fight styles have no equivalent.
- **Computed BiS lists.** Top Gear over the whole item universe per spec per phase on the server
  lane, published on the class guides and the item pages with per-slot alternatives and the
  date, regenerated on every data change. A content feature the incumbents hand-write.
- **Personalised loot on boss pages.** Every dungeon and raid boss page shows a signed-in
  member "best drop here for you: +31 DPS" from a Droptimizer run cached per character per data
  version.
- **Coaching lines.** Compare mode adds cooldown-delay and downtime analysis in words from the
  fight's cast sequence: "Death Wish was ready for 14 seconds before you used it", "9 seconds
  with no cast during the add phase".
- **Sim-derived stat weights and addon deltas.** The Phase 2 curated stat weights are replaced
  per spec and phase by weights the engine produces, and Top Gear's per-item deltas are pushed
  through the addon inbox so the in-game tooltip reads "+42 DPS, simmed".

### 4.5 Later, on the same engine

- **Top Gear**: tick bag items, enchants, and consumables from the addon export; the premium
  lane sims combinations in stages, low precision first, pruning, then paired high-precision
  runs on the finalists; results grouped with sidegrades within twice the error.
- **Droptimizer**: pick a raid; every drop simmed one at a time against current gear; per boss
  best drop, expected value, and chance of an upgrade, from the same loot tables the dungeon
  and raid pages use.
- **Stat weights**: available from the engine, shown with the same warning Raidbots gives.
- **APL editing and sharing**, and the raid sim, when there is demand.

## 5. Compute lanes

### 5.1 Browser

`sim.wasm` built from the pinned engine version, served from `/_sim/<version>/` with immutable
caching, loaded after first paint on the sim page. A worker pool sized to
`navigator.hardwareConcurrency` (capped at 8), each worker holding one engine instance; the
request is split by iterations, partial results combined on the main thread as they arrive.
Abort on navigation. Nothing leaves the device unless the player saves the result.

### 5.2 Server, premium

`POST /v1/sims` with the same protobuf request; the API checks the premium flag, writes a job
row, and starts a Cloud Run job running the native engine with the same pinned version, which
streams progress to the API and writes the result to R2 like a report. The page follows it over
the same polling path live reports use. Cost per run is on the order of one CPU-second per
thousand iterations. Premium is a flag on the user row, set by hand for testers until payments
are designed. Top Gear and Droptimizer run only on this lane.

### 5.3 One contract

Request and result are the engine's protobufs plus a small envelope (character source, engine
version, lane). Saved sims, compare mode, and the validation job read the envelope and never
the lane.

## 6. Validation

A nightly Cloud Run job, per spec: take the top 50 parses from rankings for the current phase,
each with the recorded gear, talents, buffs, and encounter length; build a sim request with the
spec's default APL; run 3,000 iterations natively; store per-ability and per-aura differences
against the fight's summary. Publish per spec: median DPS gap, the five abilities with the
largest cast-count gaps, the five auras with the largest uptime gaps. "Validated" is a median
gap under 5% on at least 50 parses; the page shows the figure either way.

Beta parses arrive through the same logs pipeline from Sept 17, so validation starts the day the
first Forever fight is logged, on beta fights first and launch fights after.

## 7. Repository layout

```
sim/                         Go: result adapter to logs events, request envelope, validation job (imports the engine module)
web/src/pages/sim/           /sim, /sim/[id], /sim/specs
web/src/lib/sim/             engine loader, worker pool, character model, sources, result → summary
web/src/components/sim/      character strip, settings bar, run control, compare layout
api/internal/sims/           saved sims, premium dispatch, validation results
data/pipeline/simdb/         engine database and per-class constants generation
data/curated/apl/            default APLs with sources
```

The engine lives outside this repository; its version is pinned in `sim/go.mod` and in
`web/package.json` (for the wasm artifact), and a CI job builds both artifacts from that pin.

## 8. Phasing and lanes

| Phase | Dates | Engine lane | Data lane | Web lane | API lane |
|---|---|---|---|---|---|
| S0 | Sept 15 to 19 | Forever repository; core stat changes; spell_mod cherry-pick; encounter environment concept; browser throughput spike | simdb generator against Era tables; switch to Forever tables on beta day | engine loader and worker pool spike; character model | result adapter to logs events |
| S1 | Sept 22 to Oct 17 | Fury Warrior and Frost Mage models with constants files; racials; default APLs | per-class constants; Forever items and sets | sim page with the signed-in landing state, four sources, results, compare mode, saved sims, spec page, phone; live DPS in the planner; execution leaderboard sort; build unfurl DPS | saved sims, validation job, spec fidelity, execution score at fight close |
| S2 | Oct 20 to Nov 4 | fixes from validation; two specs validated | launch build | polish, budgets, Lighthouse | premium flag and server lane |
| S3 | Nov 5 to Dec 9 | remaining DPS specs in rankings order | raid loot tables | Top Gear, Droptimizer, gear tooltip deltas, biggest lever, guild execution, computed BiS on guides, personalised loot on boss pages, coaching lines | staged premium sims, addon inbox results and deltas, encounter profiles from the corpus, sim-derived stat weights |

Each lane is one plan with parallel sub-lanes for independent tasks, per the execution rules in
memory. The engine lane's work happens in the engine repository and is the one place a
different review standard applies: their contribution guidelines, when the repository is theirs.

## 9. Risks

- **Forever's numbers move weekly through October.** Mitigated by constants files regenerated
  from the client and by the nightly validation figure, which shows drift the morning after a
  patch.
- **Unified hit against the vanilla weapon-skill miss table is unspecified.** Load-bearing for
  every melee spec. The first beta logs settle it; until then the engine keeps the vanilla table
  and the support page says so.
- **WoWSims may create its own Forever repository.** Then it is upstream and we contribute; the
  outreach message settles this early.
- **WebAssembly is eight times slower than native.** Designed around: live-updating estimates,
  3,000-iteration default, big jobs on the server lane. A SIMD or TinyGo build is a later
  optimisation, not a dependency.
- **Two specs at launch will disappoint players of the other seven.** The support page and the
  in-page state are the honest answer; the order follows who is actually playing.

## 10. User-owned steps

- Post the WoWSims Discord message (drafted in the scratchpad) asking about a Forever
  repository in their organisation.
- Decide the premium price and payment provider before the server lane opens to the public;
  testers use a hand-set flag.
