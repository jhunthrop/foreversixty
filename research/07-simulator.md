# Combat Simulator Research — WoW: Forever / foreversixty.gg

Research date: 2026-09-14. Forever beta opens 2026-09-17; launch 2026-11-04; first raids 2026-12-09.

All WoWSims facts below come from shallow clones taken 2026-09-14 at
`/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad/{classic,sod}`
(`wowsims/classic` @ HEAD, `wowsims/sod` @ HEAD). Line counts are `wc -l` on those trees.
Benchmarks were run locally on an Apple M4 Pro (14 cores), Go 1.25.4. SimulationCraft was built from
source and benchmarked on the same machine.

### Summary

1. **WoWSims (MIT, Go) is 80% of the backend**, and more of it survives a Forever port than expected:
   **expertise is already implemented** in the attack table, weapon skill is fully modelled, and
   `stats.SpellDamage`/`HealingPower` already exist. Merging hit and crit is 159 mechanical call sites.
2. **SimulationCraft is not a candidate.** No Classic support, ever (the "EPIC: Classic Support" issue
   has 1 of 15 boxes ticked since 2021; zero of 787 forks default to a Classic branch), and GPL-3.0
   forbids serving a WASM build to a browser. Take its ideas — the text APL, the self-expiring hotfix
   pattern — not its code.
3. **Raidbots will not compete here** (retail-only, because SimC is). Its product surface is the
   reference design; its Smart Sim precision ladder is the cost model to copy.
4. **The data pipeline is nearly free.** Wowhead's Forever data environment is **live today**
   (`nether.wowhead.com/forever/data/gear-planner`, internal name `classicplus`) and wago.tools' DB2 CSV
   API is *already* what `wowsims/classic`'s generator calls. Both are one-line changes.
5. **Measured: WASM is 7.9× slower than native Go** (155 vs 1,218 iters/sec/core). 10k iterations =
   ~16 s in a 4-worker browser vs ~1 s server-side. Plan for server compute.
6. **The differentiator is not the engine.** No sim anywhere has a log-validation loop; our `logs/`
   engine can close it automatically from beta day one. That, plus UX, is the product.
7. **The field is three days old and wide open** — four Forever repos, all created 2026-09-13/14, none
   with a licence or a real engine. `wowsims/forever` does not exist yet but almost certainly will.

---

## 1. WoWSims

### 1.1 The four repos, at a glance

| Repo | Created | License | Go LOC | Commits | Contributors | Forks | Stars | Last push |
|---|---|---|---|---|---|---|---|---|
| [wowsims/classic](https://github.com/wowsims/classic) | 2024-11-19 | MIT | 72,932 | 14,169 | 97 | 36 | 43 | 2026-07-24 |
| [wowsims/sod](https://github.com/wowsims/sod) | 2024-01-28 | MIT | 119,574 | 15,281 | 99 | 59 | 23 | 2026-07-24 |
| [wowsims/cata](https://github.com/wowsims/cata) | 2024-03-09 | MIT | — | 14,602 | 98 | 123 | 43 | 2026-09-09 |
| [wowsims/mop](https://github.com/wowsims/mop) | 2025-04-25 | MIT | — | 20,771 | 114 | 42 | 22 | 2026-09-13 |

Source: `gh api repos/wowsims/<r>`; contributor/commit counts from the `Link: rel="last"` page count on
`/contributors?per_page=1` and `/commits?per_page=1`.

None of the four is a GitHub fork of another (`"fork": false`, `parent: null` on all) — they are
independent repos seeded by copying the tree, then diverging. Practically they are the *same codebase*:
`diff <(ls classic/sim/core) <(ls sod/sim/core)` differs by exactly four files (SoD adds
`apl_values_rune.go`, `spell_mod.go`, `wowhead.go`; classic adds `apl_values_stats.go`). This matters a
lot for the fork decision — see §6.

License: MIT (`classic/LICENSE`, "Copyright (c) 2024 wowsims team"). The README adds a non-binding
request: *"We request that anyone using this software in their own project to make sure there is a user
visible link back to the original project."* That is a request, not a licence term. MIT permits a
closed-source commercial hosted fork; keep the copyright notice and add the courtesy backlink.

### 1.2 Architecture

Five layers, one repo:

```
proto/*.proto          14 files, ~86 KB     the wire contract (hand-written)
sim/core/              29,861 LOC (classic)  engine: event loop, auras, spells, attack table, APL, metrics
sim/<class>/           2,125–4,740 LOC each  hand-written spell/talent/item implementations
sim/wasm/main.go       396 LOC               syscall/js bridge → browser
sim/web/main.go        558 LOC               HTTP server → local binary / hosted
ui/                    46,548 TS LOC         TypeScript SPA, Vite, one sub-app per spec
tools/database/        3,614 LOC             the data generator
```

The protobufs are the only interface between Go and TypeScript. `makefile:197` runs
`protoc -I=./proto --go_out=./sim/core ./proto/*.proto`; `makefile:82` runs
`npx protoc --ts_out ui/core/proto --proto_path proto proto/api.proto`. Every request
(`RaidSimRequest`, `StatWeightsRequest`, `BulkSimRequest`, `ComputeStatsRequest`) and every result
(`RaidSimResult`, `StatWeightsResult`, `BulkSimResult`) is a protobuf message, serialised as bytes
across the WASM boundary or as HTTP POST bodies. Nothing else crosses.

Core engine files, largest first (`classic/sim/core/`):

```
1893 buffs.go        989 spell_outcome.go   688 spell.go         659 bulksim.go
1252 consumes.go     936 aura.go            675 apl_values_operators.go
1040 attack.go       754 character.go       662 metrics_aggregator.go
 982 debuffs.go      702 spell_result.go    653 sim.go   536 sim_concurrent.go
```

The event loop is `Simulation.run → runOnce → runPendingActions → Step`
(`sim/core/sim.go:299,368,477,491`) over a `pendingActions` heap (`sim/core/pending_action.go`) plus a
separate `advanceWeaponAttacks` path (`sim/core/sim.go:531`). Iterations are embarrassingly parallel and
are split two ways:

- **Native**: `sim/core/sim_concurrent.go` — `SplitSimRequestForConcurrency(req, splitCount)` clones the
  request N times, divides iterations, and offsets `RandomSeed` by the iteration count of each earlier
  split so the RNG stream matches a serial run. Results are recombined by `raidSimResultCombiner`, which
  merges `DistributionMetrics` (histograms, min/max with their seeds) rather than averaging averages.
- **Browser**: the same split runs across N Web Workers, each holding its own WASM instance.
  `ui/core/sim.ts:119-130` sizes the pool as `min(wasmConcurrency, navigator.hardwareConcurrency)`, with
  the default `min(4, floor(hardwareConcurrency/2))`, user-overridable up to `hardwareConcurrency` in
  `ui/core/components/settings_menu.tsx:155`.

`sim/wasm/main.go:29-43` exports exactly 13 functions to JS: `computeStats`, `raidSim`, `raidSimAsync`,
`raidSimRequestSplit`, `raidSimResultCombination`, `statWeights`, `statWeightsAsync`,
`statWeightRequests`, `statWeightCompute`, `bulkSimAsync`, `abortById`, and JSON variants. `ui/worker/`
has three interchangeable worker shims — `sim_worker.ts` (WASM in-browser), `local_worker.ts` (POST to
`http://localhost:3333`), `net_worker.ts` (POST to the serving origin). **The same Go code runs in the
browser, in a local desktop binary, and on a server, with no branching in the sim** other than
`core.SetRunningInWasm()`. This is the single most valuable architectural property of WoWSims for us.

### 1.3 How one spec is implemented — DPS Warrior (classic)

`sim/warrior/` is 3,515 lines across 40 files. The shape is one file per ability:

```
768 item_sets_pve.go   161 heroic_strike_cleave.go   64 thunder_clap.go
498 talents.go         158 items.go                  64 sunder_armor.go
290 warrior.go         150 stances.go                63 shield_slam.go
                       111 sweeping_strikes.go       60 execute.go
                        91 deep_wounds.go            56 mortal_strike.go
                        85 revenge.go                51 whirlwind.go / bloodthirst.go
```

A whole ability is one function. `sim/warrior/mortal_strike.go` in full is 56 lines:

```go
func (warrior *Warrior) registerMortalStrikeSpell(cdTimer *core.Timer) {
	if !warrior.Talents.MortalStrike { return }
	bonusDamage := 160.0
	spellID := int32(21553)
	warrior.MortalStrike = warrior.RegisterSpell(AnyStance, core.SpellConfig{
		SpellCode: SpellCode_WarriorMortalStrike,
		ActionID:  core.ActionID{SpellID: spellID},
		SpellSchool: core.SpellSchoolPhysical,
		DefenseType: core.DefenseTypeMelee,
		ProcMask:    core.ProcMaskMeleeMHSpecial,
		Flags: core.SpellFlagMeleeMetrics | core.SpellFlagAPL | SpellFlagOffensive,
		RageCost: core.RageCostOptions{Cost: 30, Refund: 0.8},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{GCD: core.GCDDefault},
			IgnoreHaste: true,
			CD: core.Cooldown{Timer: cdTimer, Duration: time.Second * 6},
		},
		CritDamageBonus: warrior.impale(),
		DamageMultiplier: 1, ThreatMultiplier: 1, BonusCoefficient: 1,
		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := bonusDamage + spell.Unit.MHNormalizedWeaponDamage(sim, spell.MeleeAttackPower(target))
			result := spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMeleeWeaponSpecialHitAndCrit)
			if !result.Landed() { spell.IssueRefund(sim) }
		},
	})
}
```

**Every number is hardcoded.** `160.0` base damage, `30` rage, `6s` cooldown, `0.8` refund, spell ID
`21553`. Nothing is read from DB2 at runtime. This is the central fact for §5 and §6.

Talents follow the same pattern. `sim/warrior/talents.go:16-35` is a single `ApplyTalents()` that adds
flat stats and then dispatches to per-talent functions:

```go
warrior.AddStat(stats.MeleeCrit, core.CritRatingPerCritChance*1*float64(warrior.Talents.Cruelty))
warrior.AddStat(stats.Defense, 2*float64(warrior.Talents.Anticipation))
warrior.applyAngerManagement(); warrior.applyDeepWounds(); warrior.applyFlurry(); ...
```

Each talent function early-returns if the talent is unspent, then either adds a stat, registers a
`RegisterResetEffect` periodic, or hooks `warrior.OnSpellRegistered(func(spell *core.Spell){...})` to
patch multipliers onto matching spells (e.g. `applyTwoHandedWeaponSpecialization`, talents.go:57-67).

Talent *shape* is generated, talent *behaviour* is not:
- `proto/warrior.proto` — a flat `WarriorTalents` message, one field per talent (`int32` for ranked,
  `bool` for 1-point), 49+ fields, generated by `tools/scrape_talents_proto.py` which Seleniums
  `https://wowhead.com/classic/talent-calc/<class>`.
- `ui/core/talents/trees/warrior.json` — grid layout for the UI (`rowIdx`, `colIdx`, `spellIds[]`,
  `maxPoints`), generated by `tools/scrape_talents_config.py`.
- `sim/warrior/talents.go` — hand-written, 498 lines.

### 1.4 The APL system

WoWSims' APL is **structured data, not a text DSL**. It is defined entirely in `proto/apl.proto`
(13,936 bytes in classic, 14,662 in sod), executed by `sim/core/apl*.go` (~13 files), and edited by a
drag-and-drop tree builder in the UI. The proto comment says outright:
*"Rotation options are based heavily on APL. See https://github.com/simulationcraft/simc/wiki/ActionLists."*

Top level:

```proto
message APLRotation {
  enum Type { TypeUnknown=0; TypeAuto=1; TypeSimple=2; TypeAPL=3; TypeLegacy=4; }
  Type type = 3;
  SimpleRotation simple = 4;
  repeated APLPrepullAction prepull_actions = 1;  // each has an APLValue do_at (negative time)
  repeated APLListItem priority_list = 2;         // hide, notes, action
}
```

`APLAction` is a `oneof` over ~20 action kinds, grouped: **casting** (`cast_spell`, `channel_spell`,
`multidot`, `multishield`, `autocast_other_cooldowns`), **timing** (`wait`, `wait_until`, `schedule`),
**sequences** (`sequence`, `reset_sequence`, `strict_sequence`), **misc** (`change_target`,
`activate_aura`, `activate_aura_with_stacks`, `cancel_aura`, `trigger_icd`, `item_swap`, `move`,
`add_combo_points`), plus escape hatches for spec-specific logic (`cat_optimal_rotation_action`,
`cast_paladin_primary_seal`, `custom_rotation`). Every action carries an optional `APLValue condition`.

`APLValue` is a `oneof` over ~78 value kinds: operators (`const`, `and`, `or`, `not`, `cmp`, `math`,
`max`, `min`), encounter (`current_time`, `remaining_time_percent`, `is_execute_phase`,
`number_targets`), resources (`current_rage`, `current_energy`, `current_combo_points`,
`time_to_energy_tick`, `energy_threshold`), GCD, autoattack timing (`auto_time_to_next`,
`auto_swing_time`), spell/dot/aura state, and stats (`current_attack_power`). SoD adds
`apl_values_rune.go` on top.

Because it is a proto, the "language" is free: JSON Schema autocompletion falls out of
`protoc-gen-jsonschema` (documented in `tools/APL.md`), and the UI builder is generated from the same
oneof. A prepull entry looks like:

```json
{"action": {"castSpell": {"spellId": {"spellId": 1}}}, "doAt": "-1s"}
```

and a condition like "Flame Shock DoT remaining > Lava Burst cast time":

```json
{"condition": {"cmp": {"op": "OpGt",
  "lhs": {"dotRemainingTime": {"spellId": {"spellId": 49233}}},
  "rhs": {"spellCastTime":   {"spellId": {"spellId": 60043}}}}},
 "castSpell": {"spellId": {"spellId": 60043}}}
```

Default rotations ship as checked-in files: `ui/<spec>/apls/*.apl.json` — **27 files in classic, 166 in
sod**. `core.GetAplRotation("../../../ui/warrior/apls", "dps_reck")` loads one directly in tests, so the
shipped default rotations are the ones the regression suite runs.

Validation feeds back to the UI: `ComputeStatsResult` carries `APLStats{ prepull_actions,
priority_list }` of `APLActionStats{ repeated string warnings }`, so a rotation referencing a spell the
character cannot cast lights up in the editor.

### 1.5 Where the data comes from

`tools/database/gen_db/main.go` (789 lines) is a multi-stage scraper, driven by `-gen=<stage>`, whose
usage block is literally in the file header:

```
go run ./tools/database/gen_db -outDir=assets -gen=atlasloot
go run ./tools/database/gen_db -outDir=assets -gen=wowhead-items
go run ./tools/database/gen_db -outDir=assets -gen=wowhead-spells -maxid=31000
go run ./tools/database/gen_db -outDir=assets -gen=wowhead-gearplannerdb
go run ./tools/database/gen_db -outDir=assets -gen=wago-db2-items
go run ./tools/database/gen_db -outDir=assets -gen=db     # offline merge → db.bin + db.json
```

The five sources:

| Stage | Source | File |
|---|---|---|
| `atlasloot` | the AtlasLoot addon's Lua loot tables | `tools/database/atlasloot.go` (434) |
| `wowhead-items` | per-item tooltip HTML scrape | `tools/database/wowhead_tooltips.go` (768) |
| `wowhead-spells` | per-spell tooltip HTML scrape (ids 1..31000) | same |
| `wowhead-gearplannerdb` | `https://nether.wowhead.com/classic/data/gear-planner?dv=100` | `tools/database/wowhead_db.go` (320) |
| `wago-db2-items` | `https://wago.tools/db2/ItemSparse/csv?build=1.15.3.55646` | `tools/database/wago_db.go` (92) |

Note the DB2 use is **narrow**: one table, `ItemSparse`, pinned to one build string. There is no
`dbc_extract` equivalent, no `Spell`/`SpellEffect`/`SpellMisc` ingestion, no client CASC reader. Spell
behaviour never comes from data.

Hand-curation is first-class: `tools/database/overrides.go` (439 lines) and
`enchant_overrides.go` (232) patch the scraped output, and `gen_db/main.go` carries an inline commentary
on de-duplicating SoD-reworked items that share a name with their Era originals.

The merge stage emits `assets/database/db.bin` (protobuf `UIDatabase`) plus `db.json`, and a
`leftover_db.*` pair. `assets/database/loader.go` `//go:embed db.bin` and
`sim/core/database_load.go` (behind `//go:build with_db`) inflate it into `proto.SimDatabase` at
`init()`. Without the `with_db` tag the DB is empty — **`go test ./sim/...` silently fails on
`No item with id: N` unless you pass `--tags=with_db`** (this is what `makefile:221` does).

Item *effects* (procs, on-use, set bonuses) are, again, hand-written Go:
`sim/common/item_effects.go` is 3,701 lines, `sim/common/enchant_effects.go` 365,
`sim/common/item_sets/` 1,319, with `sim/common/itemhelpers/` (207 lines) providing
`stat_bonus_procs.go` / `weaponprocs.go` factories for the common shapes.

Consumables are `sim/core/consumes.go`, 1,252 lines, all hardcoded. Buffs/debuffs are
`buffs.go` (1,893) + `debuffs.go` (982).

### 1.6 Features: raid sim, bulk sim, stat weights

- **Single vs raid** is not a code fork: `RaidSimRequest` always carries a `Raid{ parties[] }`;
  a single-character sim is one player in one party (`core.SinglePlayerRaidProto`). `sim/raid_test.go`
  and `sim/raid_bench_test.go` exercise multi-party compositions.
- **Bulk sim / "sim all gear"** is `sim/core/bulksim.go` (659 lines, + 377 of tests).
  `BulkSettings` takes `repeated ItemSpec items`, optional `repeated TalentLoadout talents_to_sim`, and
  an iteration budget (`defaultIterationsPerCombo = 1000`). It builds an `equipmentSubstitution` per
  combination, keeps a `raidSimRequestChangeLog`, and returns `BulkComboResult{ items_added[], ... }`.
  It refuses to run for multi-player raids ("Bulk simming is only supported for the single-player use").
- **Stat weights** are `sim/core/statweight.go` (355 lines). The request names
  `repeated Stat stats_to_weigh` + `repeated PseudoStat pseudo_stats_to_weigh` and an
  `EPReferenceStat`; the result is six `StatWeightValues` blocks (`Dps, Hps, Tps, Dtps, Tmi, PDeath`),
  each with `Weights`, `WeightsStdev`, `EpValues`, `EpValuesStdev` over a `UnitStats{ Stats[],
  PseudoStats[] }`. The three-call split (`statWeightRequests` → run → `statWeightCompute`) exists
  precisely so the N perturbation sims can be farmed out to the worker pool.

### 1.7 The results protobuf

`proto/api.proto`. `RaidSimResult` → `RaidMetrics{ dps, hps, parties[] }` → `PartyMetrics` →
`UnitMetrics`, plus `EncounterMetrics{ targets[] }`.

`UnitMetrics` carries seven `DistributionMetrics` — `dps, dpasp, threat, dtps, tmi, hps, tto` — and
`seconds_oom_avg`, `chance_of_death`, `actions[]`, `auras[]`, `resources[]`, and recursive `pets[]`.

```proto
message DistributionMetrics {
  double avg = 1; double stdev = 2;
  double max = 3; int64 max_seed = 5;   // seed replay for the best/worst iteration
  double min = 6; int64 min_seed = 7;
  map<int32,int32> hist = 4;            // the DPS histogram
  repeated double all_values = 8;
  AggregatorData aggregator_data = 9;
}
```

`max_seed` / `min_seed` are the nice trick: the UI can re-run the single best or worst iteration with
full logging to show a timeline, instead of storing every timeline.

`ActionMetrics{ id, is_melee, spell_school, is_passive, targets[] }` fans out to
`TargetedActionMetrics` which counts `casts, hits, resisted_hits, crits, resisted_crits, ticks, ...`
(field 37 and climbing, with two reserved slots from a past crit-block removal).

There is **no per-event timeline in the result protobuf.** The detailed-results UI
(`ui/detailed_results/`) works off aggregated metrics plus an optional debug log from a single
`DebugFirstIteration` run. If Forever's sim needs a shareable, inspectable timeline (and our logs engine
suggests we want one), that is net-new work in any option.

### 1.8 Regression / correctness approach

16 golden-output files (`sim/**/Test*.results`), text-format protobufs holding exact expected
`final_stats` arrays and DPS values per configuration. `make update-tests` regenerates them. Adding or
changing anything shows up as a reviewable diff of numbers. Combined with the checked-in
`ui/<spec>/gear_sets/*.gear.json` and `ui/<spec>/apls/*.apl.json`, the whole test corpus is data.

### 1.9 How WoWSims handled a fresh, rules-divergent Classic launch — Season of Discovery

This is the closest historical analogue to Forever and it is worth reading carefully.

SoD dropped new runes (effectively new abilities and passives) on a vanilla base, in phases, with no
prior theorycrafting. WoWSims' response was **not** to build a data-driven ability system. It was:

1. **A new hand-written layer per class.** `sim/<class>/runes.go` exists for all nine classes, totalling
   **5,252 lines**: warlock 768, shaman 586, warrior 525, hunter 489, paladin 483, mage 474, rogue 432,
   druid 364, priest 225. `ApplyRunes()` reads the rune id off the equipped slot
   (`warrior.Equipment.Shoulders().Rune`) and switches into hand-written effects, exactly like
   `ApplyTalents()`.
2. **One generic core primitive to absorb the new modifier space.** SoD's `sim/core/spell_mod.go`
   (790 lines, absent from classic) is a declarative spell-modifier system —
   `SpellModConfig{ Kind, ClassMask, ClassSpellsOnly, School, SpellFlags, SpellFlagsExclude,
   DefenseType, ProcMask, CastType, IntValue, TimeValue, FloatValue, ApplyCustom, RemoveCustom }` —
   so "this rune makes all Fire spells cost 10% less and crit 5% more" is a config literal rather than a
   bespoke `OnSpellRegistered` closure. Cata/MoP use the same system.
3. **Sheer headcount and iteration.** 99 contributors, 15,281 commits, per-spec ownership. The result is
   that SoD spec code is roughly **2× the size of the equivalent classic spec** (SoD warrior 5,999 LOC
   vs classic 3,515; SoD warlock 8,418 vs 4,029; SoD paladin 6,660 vs 2,381).
4. **Living with uncertainty in comments.** `grep -c "TODO\|FIXME\|not sure\|unconfirmed\|needs testing"`
   over `sod/sim` returns **271** hits. Mechanics were shipped as best guesses and corrected against logs
   afterwards.

The honest reading: a rules-divergent Classic launch cost WoWSims roughly **+580 LOC per class of new
ability code, +790 LOC of core modifier infrastructure, and a doubling of per-spec code over ~2 years
and ~100 contributors.** Forever is a bigger divergence than SoD (reworked talents across all 27 trees,
new stats, reworked racials, reworked itemization), not a smaller one.

### 1.10 Measured performance

I generated the protobufs and ran a real single-target sim (details in §6.5):

```
PERF iters=1000   fight=300s  elapsed=0.833s  1,201 iters/sec/core   360,182 sim-seconds/sec
PERF iters=10000  fight=300s  elapsed=8.208s  1,218 iters/sec/core   365,496 sim-seconds/sec
```

Phase-1 BiS Fury Warrior, full buffs/consumes, shipped `dps_reck` APL, 300 s single target,
single-threaded, Apple M4 Pro. **~0.82 ms per 300-second simulated fight.** A 10,000-iteration Quick Sim
is 8.2 s on one core, ~1.2 s across 7 cores natively.

I then built the same code to WASM and ran the identical serialised `RaidSimRequest` through the
exported `raidSim` entrypoint under Node 22 (V8, the same engine as Chrome):

```
lib.wasm (--tags=with_db)  23 MB   3.95 MB gzipped
lib.wasm (as shipped)      18 MB   3.53 MB gzipped   + assets/database/db.bin 4.9 MB (420 KB gzipped)

WASM(node) iters=10000 fight=300s elapsed=64.62s -> 155 iters/sec/core
```

**WASM is 7.9× slower than native Go here** (155 vs 1,218 iters/sec/core), not the 1.5–3× usually quoted
for Go/WASM. That is the single most important number in this document for the browser-vs-server
decision:

| | 1 core | 4 workers | 8 cores |
|---|---|---|---|
| Native Go (server) | 8.2 s | — | ~1.2 s |
| WASM (browser) | 64.6 s | **~16 s** | ~8 s |

A 10,000-iteration Quick Sim costs ~16 s on a typical 4-worker laptop and ~1 s on a modest server core
budget. For a site whose stated #1 priority is UX, that gap is not a rounding error. Caveats: measured
under Node, not a real browser (same V8, but no browser scheduling overhead modelled); Safari and
Firefox will differ; a WASM-SIMD or TinyGo build was not attempted.

### 1.11 What a Forever fork would actually touch

Reading the code against the confirmed Forever rules (§5.1):

**Changes (core):**
- `sim/core/stats/stats.go` — the `Stat` enum (589 lines total; enum is lines 20-63). Merging
  `MeleeHit`+`SpellHit` → `Hit` and `MeleeCrit`+`SpellCrit` → `Crit` touches
  **159 call sites in `sim/`** (`stats.MeleeHit` 14, `stats.SpellHit` 13, `stats.MeleeCrit` 76,
  `stats.SpellCrit` 56) and zero in `ui/` (the UI reads the proto enum, 0 direct hits). The enum must
  stay index-synced with `proto.Stat` — the file says so explicitly. Mechanical, high-volume, low-risk.
- `proto/common.proto` `Stat` enum + `ui/core/proto_utils/` stat display names.
- `sim/core/base_stats_auto_gen.go` + `tools/base_stats_parser.py` — regenerate
  `CritRatingPerCritChance`, `ExpertisePerQuarterPercentReduction`,
  `ExpertiseRatingPerExpertiseChance`, per-class-per-level base stats, from Forever values.
- `sim/core/stats/deps.go` (271 lines) — already a stat-dependency engine; "bonus healing grants ⅓ as
  spell damage" is one new dependency entry. Days, not weeks.

**Already done, no work:**
- **Expertise is already modelled.** `stats.Expertise` is in the enum, and
  `sim/core/spell_outcome.go:711-735` already reduces dodge and parry:
  ```go
  expertiseDodgeReduction := attackTable.Attacker.stats[stats.Expertise] / 100
  *chance += max(0, attackTable.BaseDodgeChance - ...DodgeReduction - expertiseDodgeReduction)
  ```
  This is inherited from the wotlk/cata lineage and is dead weight in vanilla — it becomes live for
  Forever for free. This surprised me and it is the strongest single argument for forking.
- **Weapon skill is fully modelled** — a 15-value `WeaponSkill` enum in `stats.go:67-90` and 52 call
  sites. Forever keeps weapon skill (with smaller per-item values), so this survives intact.
- **Spell damage as a weapon stat**: `stats.SpellDamage`, `stats.HealingPower`, and the seven per-school
  power stats (`ArcanePower`…`ShadowPower`) already exist. Caster weapons granting spell damage is an
  item-data change, not an engine change.
- The whole attack table — glancing blows, the vanilla dual-wield miss, the one-roll special table,
  `spell_resistances.go`, `spell_school.go` multischool handling — is vanilla-correct and Forever keeps
  "combat feels the same."

**Rewrites (per class):**
- `proto/<class>.proto` talent messages — regenerate (the scraper targets
  `wowhead.com/<env>/talent-calc/<class>`; a Forever env exists, see §5.2).
- `ui/core/talents/trees/<class>.json` — regenerate.
- `sim/<class>/talents.go` — **rewrite**. 427–498 lines per class in classic. Forever keeps seven rows
  and the 11/21/31 gold-medal talents, adds a 16-point one per tree, deletes the buff-only talents
  (Improved Battle Shout, Divine Spirit, Blessing of Kings, Improved Mark of the Wild) and makes those
  baseline, keeps many talents "unchanged from 2006", and changes others. So this is a *partial*
  rewrite: perhaps 40–60% of each `talents.go` survives, which is much better than it sounds.
- New baseline abilities per class (Paladin's `Holy Strike` at 6, baseline `Consecration` at 20,
  `Seal of Fury`, etc.) — one new 50–100 line file each, in the `mortal_strike.go` mould.
- `sim/core/racials.go` — reworked racials (two active + two passive per race, new `Touch of the Grave`,
  changed Stoneform, Mace Specialization now affecting spell crit). Full rewrite of one ~300-line file.
- `sim/common/item_effects.go` (3,701 lines) + `item_sets/` (1,319) — Forever re-examined **every**
  dungeon drop, made unique boss items blue, improved set bonuses, added hundreds of new drops,
  specialised trinkets, and added biome- and creature-type-conditional effects. Biome/creature
  conditions are a *new mechanic class* with no precedent in the engine — they need an encounter-level
  "environment" and "target creature type" concept. `proto.Target` would need a creature type and the
  encounter a biome. Small engine change, large content surface.

**Survives untouched:** the event loop, aura system, APL (all ~13 `apl_*.go` files + `apl.proto`),
metrics aggregation, bulk sim, stat weights, WASM/worker plumbing, the whole `ui/` shell, the build
system, and the test harness.


---

## 2. SimulationCraft

All facts from `simulationcraft/simc` at commit `f8352ef`, branch **`midnight`** (the current default —
`master` is not), fetched 2026-09-14. GPL-3.0, 1,591 stars, 784 forks, ~5 GB repo. The binary was
**built from source and benchmarked** for §2.6, so those numbers are measured, not quoted.
Self-report: `SimulationCraft 1210-01 for World of Warcraft 12.1.0.69814 Live (hotfix 2026-09-12/69814)`.

### 2.1 Architecture and size

| Directory | Files | LOC | Note |
|---|---:|---:|---|
| `engine/` total (excl. generated) | — | **387,061** | |
| `engine/class_modules/` | 89 | **164,199** | 42% of the engine |
| `engine/player/` | 71 | 87,346 | `player.cpp` alone is 16,650 |
| `engine/sim/` | 42 | 19,565 | |
| `engine/action/` | 30 | 16,138 | |
| `engine/report/` | 32 | 14,639 | |
| `engine/dbc/` hand-written | ~81 | 21,640 | |
| **`engine/dbc/generated/*.inc`** | 68 | **1,241,121** | **157.6 MB** of machine-generated C arrays |

The generated DBC is 158 MB of a 173 MB `engine/` tree: `sc_spell_data_ptr.inc` 29.6 MB,
`sc_spell_data.inc` 29.5 MB, `item_data_ptr.inc` 26.8 MB, `item_data.inc` 26.4 MB.

Core objects:
- **`sim_t`** — [`engine/sim/sim.hpp:65`](https://github.com/simulationcraft/simc/blob/midnight/engine/sim/sim.hpp),
  `struct sim_t : private sc_thread_t`. **It *is* the thread**; worker threads are child `sim_t`
  instances. `sim.cpp` is 4,912 lines, of which `sim.cpp:3700–3950` is one long `add_option(...)` block.
- **`player_t`** — `engine/player/player.hpp:139`. 1,544-line header, 16,650-line `.cpp`. Owns the
  `create_action()` dispatch and `create_expression()`.
- **`action_t`** — `engine/action/action.hpp:60`. Hierarchy: `action_t` → `spell_base_t` →
  `spell_t` / `heal_t` / `absorb_t`, and `action_t` → `attack_t` → `melee_attack_t` /
  `ranged_attack_t`. `engine/action/parse_effects.cpp` (2,078 lines) is the modern "read the DBC effect
  and auto-wire the multiplier" layer class modules build on.
- **`buff_t`** — 4,401 lines. Its `stack_uptime` sample data lands verbatim in the JSON report.
- **`dbc_t`** — `engine/dbc/dbc.hpp:311`. A thin stateful accessor over the static generated arrays; a
  `ptr` flag selects `__spell_data` vs `__ptr_spell_data`.
- **`spell_data_t`** — `engine/dbc/spell_data.hpp:439`, with fields annotated by their source table:
  ```cpp
  const char* _name;        // Spell name from Spell.dbc stringblock (enGB)
  unsigned    _cooldown;    // SpellCooldown.dbc, milliseconds
  unsigned    _max_stack;   // SpellAuraOptions.dbc
  double      _rppm;        // Base real procs per minute
  unsigned    _class_flags[NUM_CLASS_FAMILY_FLAGS];  // SpellClassOptions.dbc
  ```
  Live sizes from the generated headers: **31,636 spells, 53,234 effects, 68,238 item stat rows.**

**The event queue is a hashed timing wheel, not a binary heap.** This is commonly mis-stated and it is
the most interesting engineering detail in the project. `engine/sim/event_manager.hpp` holds a
`std::vector<event_t*> timing_wheel`; `event_manager_t::init()` (`event_manager.cpp:331`) comments:

> *"Timing wheel depth defaults to about 17 minutes with a granularity of 32 buckets per second. This
> makes wheel_size = 32K and it's fully used."*

`add_event` (line 110) computes `slice = (time_ms >> wheel_shift) & wheel_mask` (`wheel_shift = 5`, so
32 ms buckets) and does an ordered linked-list insert within the bucket, ties broken by a monotonic
`id`. Events beyond the wheel are parked at `wheel_time - 1s` with a `reschedule_time`. **O(1)
amortized, not O(log n).** Events are never `delete`d mid-sim — `allocate_event()` pulls from a
`recycled_event_list`. `event_t` subclasses are size-capped so everything fits one fixed pool.
Tunable via `wheel_granularity=` / `wheel_seconds=` / `wheel_shift=`.

`sim_t::combat()` (`sim.cpp:1790`) is three lines: `combat_begin(); event_mgr.execute(); combat_end();`.

**Threading is across iterations, not within one.** `sim_t::partition()` (`sim.cpp:3372`) creates
`threads - 1` child `sim_t` clones, each with its own event manager, actors, and RNG. By default all
children **share one `work_queue_t`** (a mutex around work vectors) so load balances naturally; with
`deterministic=1` or `strict_work_queue=1` each thread gets a fixed slice. `sim_t::merge()`
(`sim.cpp:3264`) folds every sample-data container back into the parent under a `merge_mutex`.
`threads=` accepts negatives as relative: *"`threads=-2` on a 8 core machine will result in 6 threads."*

`sim_t::analyze_error()` (`sim.cpp:2131`) runs on thread 0 every `analyze_error_interval` iterations,
computes relative standard error, and calls `interrupt()` once it drops below `target_error` — otherwise
re-projects `n * (current_error² / target_error²)` into the queue. Defaults: `iterations=1000000`,
`target_error=0.200`.

### 2.2 The APL text DSL

In-repo docs are thin (`doc/` is a Doxyfile and a 9-line mainpage). The real reference is the wiki,
which is itself a cloneable git repo (`github.com/simulationcraft/simc.wiki.git`, ~63 pages):
[ActionLists](https://github.com/simulationcraft/simc/wiki/ActionLists) (644 lines) and
[Action-List-Conditional-Expressions](https://github.com/simulationcraft/simc/wiki/Action-List-Conditional-Expressions)
(573).

```
actions=first_spell
actions+=/second_spell,if=<expr>
actions.<listname>=first
actions.<listname>+=/second,if=<expr>
actions.precombat=snapshot_stats
```
`/` separates entries, `,` separates options within an entry, `+=` appends. Semantics, from the wiki:

> *"Actions lists are priorities lists: periodically, Simulationcraft scans your character's actions
> list, starting with the first action (the highest priority) and continuing until an available action
> is found."*

**Sub-lists**, both in `engine/action/action.hpp:1165–1193`: `call_action_list_t`, `swap_action_list_t`,
and `run_action_list_t` (which *derives from* `swap_action_list_t`). The wiki's distinction:
*"`call_action_list` effectively inserts the action list into the default APL where it was called"*,
whereas *"if `run_action_list` is called and there are no usable actions in the sub-APL, it will return
to the top of the default action list with a delay."* `action_priority_list_t` carries an
`internal_id_mask` explicitly commented as the guard against *"potential infinite loops in the APL"*.

**Variables** — `engine/action/variable.cpp` (410 lines). Full `op=` set (`variable.hpp:73–87`):
`set print reset add sub mul div pow mod min max floor ceil setif report`. `setif` takes `value=`,
`value_else=`, `condition=`. There is also a `cycling_variable`.

**The expression grammar** lives in `engine/sim/expressions.hpp` (291) and `expressions.cpp` (**1,686**):
a hand-written lexer (`lexer_t::next()`, line 60), shunting-yard to RPN (~line 1310), tree build, and a
large constant-folding/short-circuit optimizer (lines 600–1000) driven by `optimize_expressions=`.

Note the operator quirk — **`%` is division**, because `/` is the option separator (`expressions.cpp:85`):
```cpp
case '%': { if ( match('%') ) return yield_token( TOK_MOD ); return yield_token( TOK_DIV ); }
case '@': return yield_token( TOK_ABS );
case '<': { if (match('=')) return TOK_LTEQ; if (match('?')) return TOK_MAX; return TOK_LT; }
```
`~` prefixes float-tolerant comparisons (`~=`, `~<`, … tolerance 1e-9); bare `~` / `!~` are substring
`in` / `not in` (SpellQuery only). Precedence high→low (`expressions.cpp:1081`): `floor`/`ceil` 9 ·
unary `! + - @` 8 · `* % %%` 7 · `+ -` 6 · `<? >?` 5 · comparisons 4 · `&` 3 · `^` 2 · `|` 1.
**Everything evaluates to a double**; 0 is false.

`action_t::create_expression()` (`action.cpp:3179`, ~1,090 lines) resolves action properties and the
dotted prefixes `dot` `debuff` `target` `action` `prev` `prev_gcd` `spell_targets` `in_flight` `self`
`sim`, then falls through to `player_t::create_expression` (`player.cpp:11769` — `buff` `cooldown`
`talent` `spec` `hero_tree` `stat` `variable` `trinket` `set_bonus` `equipped` `race` `pet` `movement`
`active_dot` …) and `sim_t::create_expression` (`sim.cpp:3563` — `time` `fight_remains` `active_enemies`
`desired_targets` `target` …). `raid_event_t::evaluate_raid_event_expression` (`raid_event.cpp:2585`)
handles `raid_event.<type>.<exists|up|in|remains|duration|cooldown|distance|amount|count>`, folding to a
constant when no matching event exists.

Real APL, verbatim from `ActionPriorityLists/default/paladin_retribution.simc`:

```
# Executed every time the actor is available.
actions=auto_attack
actions+=/rebuke
actions+=/call_action_list,name=cooldowns
actions+=/call_action_list,name=generators

actions.cooldowns=use_item,name=algethar_puzzle_box,if=(cooldown.avenging_wrath.remains=0&!talent.radiant_glory|(!talent.execution_sentence&cooldown.wake_of_ashes.remains=0|cooldown.execution_sentence.remains=0)&talent.radiant_glory)
actions.cooldowns+=/potion,if=buff.avenging_wrath.up|fight_remains<30|talent.radiant_glory&cooldown.wake_of_ashes.remains=0&(!talent.holy_flames|dot.expurgation.ticking)
actions.cooldowns+=/invoke_external_buff,name=power_infusion,if=buff.avenging_wrath.up|...
actions.cooldowns+=/lights_judgment,if=!raid_event.adds.exists|raid_event.adds.in>75|raid_event.adds.up
actions.cooldowns+=/avenging_wrath,if=(!raid_event.adds.up|target.time_to_die>10)&(!talent.holy_flames|dot.expurgation.ticking)

actions.finishers=variable,name=ds_castable,value=(active_enemies>=2|buff.divine_arbiter_verdict.up)&!buff.empyrean_legacy.up
actions.finishers+=/divine_storm,if=variable.ds_castable
actions.finishers+=/templars_verdict

actions.generators=call_action_list,name=finishers,if=holy_power=5&cooldown.wake_of_ashes.remains|buff.hammer_of_light_free.remains<gcd*2
actions.generators+=/wake_of_ashes,if=(cooldown.avenging_wrath.remains>6|talent.radiant_glory)&(!talent.execution_sentence|cooldown.execution_sentence.remains>4|target.time_to_die<10)&(!raid_event.adds.exists|raid_event.adds.in>10|raid_event.adds.up)
actions.generators+=/blade_of_justice,if=(buff.art_of_war.stack=2|buff.righteous_cause.stack=2)&(!talent.walk_into_light|!buff.avenging_wrath.up)
actions.generators+=/call_action_list,name=finishers
actions.generators+=/judgment
```

`setif` with modulus, from the precombat block of the same file:
```
actions.precombat+=/variable,name=trinket_1_sync,op=setif,value=1,value_else=0.5,condition=variable.trinket_1_buffs&(trinket.1.cooldown.duration%%cooldown.avenging_wrath.duration=0|cooldown.avenging_wrath.duration%%trinket.1.cooldown.duration=0)
```

**The comparison to WoWSims matters for us.** SimC's APL is *text*: postable in a forum, diffable in
git, greppable, hand-writable, and the thing the theorycrafting audience already knows. WoWSims' is a
protobuf tree edited by a GUI (§1.4) — better for a builder UI, worse for sharing and review. Raidbots
exposes raw SimC APL text as an "Expert Mode" escape hatch on every tool (§3.1) precisely because the
text form is what power users want.

Fight styles (`engine/sc_enums.hpp:1395`): `Patchwerk` `CastingPatchwerk` `HecticAddCleave`
`DungeonSlice` `DungeonRoute` `CleaveAdd` `LightMovement` `HeavyMovement` `Beastlord` `HelterSkelter`
`Ultraxion`. Raid event types (`raid_event.cpp:2388`): `adds` `pull` `move_enemy` `casting`
`distraction` `invulnerable` `interrupt` `movement` `damage` `heal` `stun` `vulnerable` `absorb`
`position_switch` `flying` `damage_taken_debuff` `damage_done_buff` `buff`.

### 2.3 Class module structure

Single-file modules (`engine/class_modules/*.cpp`, 112,262 lines):

| File | LOC | | File | LOC |
|---|---:|---|---|---:|
| `sc_death_knight.cpp` | **17,730** | | `sc_evoker.cpp` | 11,554 |
| `sc_shaman.cpp` | 14,599 | | `sc_warrior.cpp` | 9,816 |
| `sc_druid.cpp` | 14,448 | | `sc_hunter.cpp` | 9,130 |
| `sc_demon_hunter.cpp` | 13,170 | | `sc_mage.cpp` | 7,569 |
| `sc_rogue.cpp` | 11,933 | | `sc_enemy.cpp` | 2,221 |

Split (directory) modules: `warlock/` 14,358 · `priest/` 12,128 · `paladin/` 10,362 · `monk/` 10,255 ·
`apl/` 5,023. **`engine/class_modules/` is 164,199 LOC across 89 files — the largest subsystem in the
project by a wide margin.**

Taking **paladin** as the worked example (`sc_paladin.cpp` 5,367 · `sc_paladin.hpp` 1,893 ·
`_retribution` 1,263 · `_protection` 1,124 · `_holy` 715):

```cpp
template <class Base> struct paladin_action_t : public parse_action_effects_t<Base>  // .hpp:1122
struct paladin_spell_t        : public paladin_spell_base_t<spell_t>                 // :1358
struct paladin_heal_t         : public paladin_spell_base_t<heal_t>                  // :1365
struct paladin_melee_attack_t : public paladin_action_t<melee_attack_t>              // :1417
template <class Base> struct holy_power_consumer_t : public Base                     // :1431
```

`paladin_action_t`'s constructor is where DBC meets gameplay — it asks the spell data whether it is
affected by each aura's effect and builds an `affected_by` whitelist:
```cpp
this->affected_by.highlords_judgment = this->data().affected_by( p->mastery.highlords_judgment->effectN( 1 ) );
this->affected_by.crusade            = this->data().affected_by( p->spells.crusade->effectN( 1 ) );
if ( this->data().ok() ) { p->apply_action_effects( this ); }
```
Concrete spells take a `spell_data_t*`:
```cpp
: paladin_spell_t( "consecration", p, p->find_spell( 26573 ) )
: paladin_spell_t( "blessing_of_protection", p, p->find_talent_spell( talent_tree::CLASS, "Blessing of Protection" ) )
```

**Talents.** `engine/player/talent.hpp` (55 lines) — `player_talent_t` wraps a `trait_data_t*`, a
`spell_data_t*`, and a rank, with `enabled() { return m_rank > 0 && m_spell->ok(); }` and an implicit
`operator const spell_data_t*()` so a talent passes anywhere a spell does. Five `find_talent_spell`
overloads (`player.hpp:1017–1023`). `paladin_t::init_spells()` (`sc_paladin.cpp:4155`) is a ~500-line
wall of `talents.fist_of_justice = find_talent_spell( talent_tree::CLASS, "Fist of Justice" );` with a
`register_passive_effect_override` escape hatch for server-side script effects the DBC doesn't model.

**Default APLs are generated C++.** `engine/class_modules/apl/` (26 files, 5,023 LOC):
`apl_hunter.cpp` 585 · `apl_demon_hunter.cpp` 576 · `apl_monk.cpp` 566 · `apl_warrior.cpp` 423 ·
`apl_death_knight.cpp` 409 · `warlock.cpp` 380 · `apl_priest.cpp` 344 · `mage.cpp` 343 ·
`apl_evoker.cpp` 316 · `apl_rogue.cpp` 272 · `apl_shaman.cpp` 223 · `apl_paladin.cpp` 120. (No
`apl_druid.cpp` — druid's APL lives inside `sc_druid.cpp`; naming is inconsistent.)

`ConvertAPL.py` (158 lines) reads plain APL text and writes C++ between sentinel comments:
```cpp
//retribution_apl_start
void retribution( player_t* p ) {
  action_priority_list_t* default_   = p->get_action_priority_list( "default" );
  action_priority_list_t* cooldowns  = p->get_action_priority_list( "cooldowns" );
  precombat->add_action( "snapshot_stats", "Snapshot raid buffed stats before combat begins..." );
  default_->add_action( "call_action_list,name=cooldowns" );
```
`apl_paladin.cpp` and `retribution_apl.inc` are line-for-line identical modulo syntax. Its README:
*"These files only contain **generated text** and should not be manually modified at any time"* — and
usefully, *"The contents of these files can be placed directly within any SimC input… or custom APL
box/Advanced Sim in raidbots.com."* CI regenerates two classes automatically
(`generate_apl_modules_ci.sh`: `demon_hunter shaman`).

`profiles/` is 1.63 MB / 157 `.simc` files — `MID1/`, `MID2/` (48 current-tier profiles including
hero-talent variants), `PreRaids/`, `tests/`, plus `CI.simc`.

### 2.4 Spell data pipeline

**Everything is in the main repo — there is no separate extractor project.**

`dbc_extract3/` is **11,742 lines of Python**: `dbc/generator.py` **5,327** · `dbc/wdc1.py` 1,381 ·
`dbc/filter.py` 929 · `dbc/data.py` 726 · `dbc/wdc2.py` 494 · `dbc_extract.py` 454, plus WDC1–WDC5
readers and a per-build JSON schema directory `formats/`.

**The source of truth is Blizzard's own CDN/NGDP, not wago.tools and not a local install.**
`casc_extract/casc.py:657` talks to `http://us.patch.battle.net:1119/...` (the Ribbit patch protocol),
implements CASC/BLTE decoding and Salsa20 for encrypted files. Workflow:
`python casc_extract.py --cdn -m batch -o <path>` (`--ptr`, `--beta`, `--product=wowxptr`).
**CDN download → DB2 → Python extraction → `.inc` generation.**

Tables linked explicitly (`generator.py:2787–2799`):
```python
self._data_store.link('SpellEffect',    'id_parent', 'SpellName', 'add_effect')
self._data_store.link('SpellMisc',      'id_parent', 'SpellName', 'misc')
self._data_store.link('SpellLevels',    'id_parent', 'SpellName', 'level')
self._data_store.link('SpellCooldowns', 'id_parent', 'SpellName', 'cooldown')
self._data_store.link('SpellScaling',   'id_spell',  'SpellName', 'scaling')
```
plus `SpellPower`, `SpellCategories`, `SpellAuraOptions`, `SpellClassOptions`, `SpellLabel`,
`SpellRange`, `SpellDuration`, `SpellRadius`, `SpellInterrupts`, `SpellTargetRestrictions`,
`SpellShapeshift`, `SpellMechanic`, `SpellEquippedItems`, `ItemSparse`, `ItemEffect`, `ItemBonus`,
`ItemSet`/`ItemSetSpell`, `SpellItemEnchantment`, `GemProperties`, `RandPropPoints`, `ContentTuning`,
`ExpectedStat`, the `TraitNode*` family, `AssistedCombat*`, and the `GameTables/` scaling curves.

`live.conf` is the batch manifest mapping generators to the 28 output files
(`[sc_spell_data.inc] generators = SpellDataGenerator`, etc.); `ptr.conf` mirrors it. Every `.inc`
ships twice — `foo.inc` and `foo_ptr.inc` — selected by `dbc_t::ptr`.

**Two separate hotfix mechanisms:**
1. **Blizzard's**, baked in at extraction. Blizzard pushes DB2 overrides into the client's
   `Cache/ADB/<locale>/DBCache.bin`; `dbc/db.py:146 __apply_hotfixes()` merges them field-by-field and
   stamps a `_flags` bitmask. Result versioned in `client_data_version.inc`:
   ```c
   #define CLIENT_DATA_WOW_VERSION "12.1.0.69814"
   #define CLIENT_DATA_HOTFIX_DATE "2026-09-12"
   #define CLIENT_DATA_HOTFIX_HASH "dca34b30…"
   ```
   So the base data is CDN, but the hotfix cache does come from a real client install.
2. **SimC's own runtime overrides** — `namespace hotfix` (`dbc.hpp:230–290`), a fluent builder applied
   at startup by each module's `register_hotfixes()`:
   ```cpp
   hotfix::register_effect( "Death Knight", "2026-08-22", "Frost aura (direct) buffed 6%", 179689, hotfix::HOTFIX_FLAG_LIVE )
       .field( "base_value" ).operation( hotfix::HOTFIX_SET ).modifier( 2 ).verification_value( -4 );
   ```
   `verification_value` makes the override self-disable once Blizzard's real data catches up. **This is
   a good pattern to copy for a game with unknown numbers: a dated, self-expiring override with the
   value you expected to see.**

`SpellDataDump/` (135.6 MB) holds human-readable per-class text dumps, regenerated on every data update.

### 2.5 Report output

```cpp
add_option( opt_func( "json",  parse_json_reports ) );   // sim.cpp:3865
add_option( opt_func( "json2", replace_json2 ) );        // rewrites to json=X,version=2
```
Syntax: `json=<file>[,version=<semver>][,full_states=0|1][,pretty_print=0|1][,decimal_places=N]`.
Supported versions (`report_configuration.cpp:34`): `"3.0.0-alpha1"` (default) and `"2.0.0"`, selected
by real semver-range matching. Other outputs: `html=`, `output=` (text), `xml=`, `report_details=0|1`.

Files: `report_html_player.cpp` **5,039** · `charts.cpp` 1,834 · `json/report_json.cpp` **1,479** ·
`report_text.cpp` 1,377 · `report_html_sim.cpp` 1,351.

Measured on a default Ret Paladin sim (converged at 1,019 iterations, target_error 0.2%):

```
rb_out.json    506,687 bytes   (report_details=1, the default)
rb_min.json     16,573 bytes   (report_details=0)
rb_out.html  1,521,343 bytes
```
**`report_details=0` shrinks the JSON ~30×** — worth knowing if you store reports.

Structure: `sim.players[i]` carries `collected_data` with `fight_length waiting_time dmg dps dpse
target_metric buffed_stats resource_lost timeline_dmg resource_timelines action_sequence_precombat
action_sequence`, plus `dtps heal hps absorb deaths timeline_dmg_taken` for tank/healer roles.
`extended_sample_data_t` serialises as:
```json
"dps": {"sum":236479246.7,"count":1019,"mean":232069.918,"min":210467.888,"max":259756.198,
        "median":231925.695,"variance":52482374.09,"std_dev":7244.472,"mean_std_dev":226.944}
```
Per-action entries carry `id spell_name school type num_executes compound_amount total_intervals`;
per-buff entries carry `start_count refresh_count interval trigger duration uptime benefit
overflow_stacks stack_uptime`.

**Gap worth knowing: the DPS distribution histogram is NOT in the JSON.**
`extended_sample_data_t::create_histogram(50)` exists (`engine/util/sample_data.hpp:483`) and populates
a `distribution` vector, but `report_json.cpp` never serialises it — it is HTML-chart-only. From JSON
you get min/max/mean/median/variance/std_dev plus full per-second `timeline_dmg` and
`resource_timelines` arrays, and you would rebuild a histogram from those. (Contrast WoWSims, which
puts `map<int32,int32> hist` right in `DistributionMetrics` — §1.7.)

Text report throughput block (`report_text.cpp:755`), where `SpeedUp` =
`iterations * mean_fight_length / elapsed_cpu`, i.e. **simulated game-seconds per CPU-second**:
```
Iterations = 10000   TotalEvents = 58464245   MaxEventQueue = 89
SimSeconds = 2999729.737   CpuSeconds = 30.21763   WallSeconds = 31.024921792
SpeedUp = 99281
```

### 2.6 Performance — measured

**There is no published benchmark page.** All ~63 wiki pages were enumerated; none covers performance.
The only relevant issue, [#3989](https://github.com/simulationcraft/simc/issues/3989) (2018), is a
qualitative complaint about merge cost at 20+ threads, since fixed (measured `MergeSeconds` at 14
threads: 0.0017 s). So it was built and measured: `make -j14 SC_NO_NETWORKING=1`, Apple Silicon 14-core,
stock current-tier `MID2` profiles, `iterations=10000 threads=1 max_time=300 fight_style=Patchwerk
target_error=0`.

| Profile | Events | CPU-s | **iters/s/core** | events/CPU-s |
|---|---:|---:|---:|---:|
| MID2 Paladin Retribution | 58.5 M | 30.22 | **331** | 1.94 M |
| MID1 Death Knight Frost | 89.0 M | 31.98 | **313** | 2.78 M |
| MID2 Priest Shadow | 98.0 M | 40.96 | **244** | 2.39 M |
| MID2 Mage Fire | 132.7 M | 46.04 | **217** | 2.88 M |
| MID2 Death Knight Frost | 114.9 M | 46.88 | **213** | 2.45 M |
| MID2 Warlock Demonology | 139.2 M | 54.93 | **182** | 2.53 M |

**~180–330 iterations/sec/core**, or more stably **~2–3 million events/CPU-second**. A default
10,000-iteration sim costs 30–55 CPU-seconds. Tier inflation is real: the same spec costs **31.98 CPU-s
on MID1 gear vs 46.88 on MID2**, +47% for one tier.

Thread scaling (MID2 Ret, 10,000 iterations): 1 thread 31.0 s wall · 4 threads 7.97 s (**3.89×**, 97%
efficiency) · 14 threads 3.05 s (**10.18×**, 73%). Total CPU work grows 26% from 1→14 threads. The
14-thread drop is partly Amdahl and partly this SoC's P-core/E-core split — **do not generalise 73% to
a homogeneous server**.

Realistic service workloads:
- **Default converged sim** (no iteration override; `target_error=0.2`): Ret Paladin, 14 threads →
  converged at **1,019 iterations, 3.82 CPU-seconds, 0.40 s wall.** This matches Raidbots' own
  published telemetry almost exactly (p50 Quick Sim ≈4.6 s CPU at ≈17.8k iterations — §3.6).
- **Droptimizer-shaped batch** — baseline + 50 profilesets at 0.2%, 14 threads: 16.09 s wall,
  **196.93 CPU-s**, 299.5 M events → **~3.9 CPU-seconds per gear combination.** A 200-item Droptimizer
  at flat 0.2% is ~780 CPU-s ≈ 25 s on a 32-core box before any culling.

Monte Carlo error goes as 1/√n, so iterations go as 1/error². 0.2% → 0.05% is **16×** the work;
culling at 1% costs **1/25th** of a 0.2% run. That arithmetic is the entire economics of Raidbots'
Smart Sim staging (§3.6), and it applies identically to whatever we build.

**Caveat on cross-engine comparison:** SimC's 180–330 iters/s/core is for *retail* fights with 58–139 M
events. My WoWSims measurement of 1,218 iters/s/core (§1.10) is for a *vanilla* fight with a fraction of
the event count. **These numbers are not an engine-quality comparison** — they say that vanilla combat
is cheap to simulate, which is good news for us either way. Both were measured on the same machine.
All figures are order-of-magnitude, ±2×; re-benchmark on target hardware before sizing anything.

### 2.7 Licence: GPL-3.0, and why that does not block a hosted service

Confirmed GPL-3.0; `LICENSE` and `COPYING` are both the verbatim 35,147-byte GPLv3 text. Individual
source files carry no per-file notice. Bundled third-party libs are all GPLv3-compatible: RapidJSON
(MIT), RapidXML (MIT), {fmt} (MIT), cpp-semver (MIT), UTF-8 CPP (Boost), utf8.h (Unlicense),
MSInttypes (BSD-3), and **Qt (LGPL-3.0 — GUI only)**.

**GPLv3 is not AGPL, and the difference is the whole question.** GPLv3 §0:

> *"To 'convey' a work means any kind of propagation that enables other parties to make or receive
> copies. **Mere interaction with a user through a computer network, with no transfer of a copy, is not
> conveying.**"*

AGPLv3 §13 adds the network clause GPLv3 deliberately omits. So **running simc server-side, taking
input over HTTP and returning output, is not conveying** — no obligation to release your frontend, API,
scheduler, or orchestration. This is the posture Raidbots has operated under publicly since 2017.

You *would* trigger obligations by:
1. **Shipping the binary** — a desktop app, a published Docker image, or **a WASM build served to the
   browser. Compiling simc to WASM and sending it to the client is conveying** and pulls the combined
   work under GPLv3. (This alone rules SimC out of WoWSims' browser-compute model.)
2. **Linking it as a library into your own binary.** SimC is C++ with no stable C ABI; the safe and
   standard pattern — Raidbots' pattern — is to **exec the `simc` process and read its JSON**.
3. Distributing a modified binary without publishing the source changes.
4. Bundled Qt is LGPL-3.0 — irrelevant if you build CLI-only (`BUILD_GUI=OFF`).

*(This is a reading of the licence text, not legal advice.)*

### 2.8 Classic support: confirmed absent, briefly attempted, abandoned

**There is no `classic` branch.** All 15 branches track retail (`midnight`, `thewarwithin`,
`dragonflight`, `shadowlands`, `legion-dev`, `bfa-dev`, `wod`, `mop` [= retail Mists, 2014]). Two
Classic-adjacent branches exist and both are dead:

- **[`tbc`](https://github.com/simulationcraft/simc/tree/tbc)** — last commit **2021-03-02**, exactly
  **2 commits ahead** of its base and ~20,000 behind `midnight`. Those two commits are
  `"start the purge"` and `"remove spell data from priest spells"`. **That is the entire effort.**
- **[`old-tbc`](https://github.com/simulationcraft/simc/tree/old-tbc)** — 2021-03-01, a resurrected copy
  of the 2008-era TBC simc code as a reference point. Its README: *"Old TBC simc code ripped from
  [this commit]."*

The README contains no "classic", "vanilla", "era", "TBC" or "WotLK". Of 11,822 issues and PRs, three
mention Classic in the title and zero mention vanilla:
- **[#5821 "EPIC: Classic Support"](https://github.com/simulationcraft/simc/issues/5821)** — opened
  2021-03-01, **still open**, last activity 2024-09-06, **1 of 15 checklist boxes ticked** ("Branch
  created"). Maintainer `seanpeters86`: *"No there has been no work on spell data"* and *"Need more
  contributors to help out on the project."*
- **[#8549](https://github.com/simulationcraft/simc/issues/8549)** — `vituscze`: *"WotLK classic was
  never supported."*
- **[#7574](https://github.com/simulationcraft/simc/issues/7574)** — `vituscze`: *"I don't believe it
  was ever supported. In any case, **no support for classic is currently planned**."*

The addon ([simulationcraft/simc-addon](https://github.com/simulationcraft/simc-addon)) declares
`## Interface: 120005, 120007, 120100` — retail TOCs only, so `/simc` gear export does not work on a
Classic client.

**All 787 forks were enumerated.** Default-branch distribution: `legion-dev` 168, `shadowlands` 149,
`bfa-dev` 144, `dragonflight` 111, `thewarwithin` 92, `midnight` 89, `master` 25. **Zero forks default
to a Classic branch, and no `simc-classic`-style fork exists on GitHub.**

The reason is structural, not political: simc is built end-to-end around modern DB2 spell data, the
Dragonflight+ trait/talent system, and an item/bonus-ID model that Classic clients simply do not have.

**What the community uses instead is wowsims — and it is a from-scratch Go engine, not a simc fork.**
Verified: every `wowsims/*` repo reports `"fork": false`; `go.mod` declares its own module with no simc
linkage; MIT not GPL; first repo created 2021-09-05, **six months after the simc Classic EPIC stalled.**
Activity in the last 90 days: `tbc-new` **423 commits**, `mop` **545**, `cata` 2, `tbc` 3,
**`classic` 1**, `sod` 2, `wotlk` **0** (despite being the highest-starred at 183). Unlike simc,
**healing and tank specs are first-class** across the wowsims repos.

### 2.9 What Raidbots runs on

Confirmed from the operator's own writing (Dave Hendler, GitHub `seriallos`):

> *"Raidbots is a web application with two goals: Provide an easy-to-use web UI to generate scripts for
> SimulationCraft… Provide access to powerful cloud hardware to run simulations (SimC is incredibly
> hungry for CPU power)"* — [Technical Architecture](https://medium.com/raidbots/raidbots-technical-architecture-303349d82784)

> *"**What's the difference between Raidbots and SimulationCraft?** Raidbots runs SimulationCraft on
> cloud hardware."* — [FAQ](https://medium.com/raidbots/frequently-asked-questions-2933b01a2d6e)

And the process-exec boundary from §2.7, in their own words: *"Grab sim jobs from the queue. **Spawn and
monitor the SimC process and hand in the SimC input.** On success, save the HTML, JSON and stdout/stderr
output to permanent storage."*

**Stock builds or patched?** Evidence favours stock upstream source on a custom build *cadence*; source
modification is **unverified in either direction**. The FAQ says *"Raidbots uses a custom 'weekly' build
that is compiled from the latest changes on Monday nights"* — "custom" is the schedule. Their frontend
config pins an upstream branch name directly:
```js
live: { simcBranch: "midnight", simcDataDomain: "live", simcChannel: "latest" }
simcBranch: { help: "Branch the data-update PR targets and the env's SimC build comes from." }
```
`seriallos/simc` is a fork but has been **dead since 2016-12-08**; it is not the build source. Smart Sim
and SwiftSim are orchestration *over profilesets*, outside the engine.

**They contribute back, and they own critical infrastructure.** `seriallos` has **1,085 commits and
1,078 PRs** on `simulationcraft/simc`. Of the 100 most recent PRs, **99 are automated
`[live] Game data update (Build NNNNN)`** — i.e. **Raidbots operates the DBC extraction pipeline
(§2.4) for upstream SimC**, running it on every WoW build and publishing the result as PRs.

The dependency runs both ways. Upstream simc has Raidbots baked in:
- `engine/report/report_html_player.cpp` renders talent images from
  `https://{www|mimiron}.raidbots.com/simbot/render/talents/{...}`.
- `engine/item/item.cpp` carries `std::string DUMMY_CONTEXT; // not used by simc but used by 3rd parties (raidbots)`.
- `qt/WebPage.cpp` special-cases `raidbots.com`.
- Raidbots grants SimC devs a free internal tier (`simcdev: { level: 80, iterationsLimit: 16e6, apiAccess: true }`).

Infrastructure claims over time (see §3.6 for the fuller picture): Apr 2017 *"16-CPU machines for
workers… roughly 2–4× as fast as decently specced gaming PCs"*; May 2017 *"~500 CPUs"* at peak;
Jul 2023 *"all sims have been running on VMs using 32 CPU cores"*, with SwiftSim using two such servers.
**One conflict to note:** the Flightmaster post describes workers as ephemeral/preemptible and says that
is why sims were iteration-capped before chunking, while the architecture post never names spot VMs.
Both readings are defensible from the public text; current fleet composition is **unverified**.

---

## 3. Raidbots' product surface

Retail-only, SimulationCraft-in-the-cloud. Seven tools. Sources are Raidbots' own help centre, their
Medium archive, and their public (unauthenticated) frontend config bundle
`https://www.raidbots.com/frontend/simbot.4435b7a9e650185158cb.js` — where a number comes from that
bundle rather than prose documentation, it is marked *(config)*.

### 3.1 The tools

| Tool | Route | Input | Output |
|---|---|---|---|
| **Quick Sim** | `/simbot/quick` | one character (Armory or `/simc` paste) + sim options | DPS/HPS + **Damage Breakdown**, **Buff Uptime**, **Sample Ability Log** (one iteration), **Simulation Details**, link to full SimC HTML report |
| **Top Gear** | `/simbot/topgear` | `/simc` paste (needs bags) + per-slot item ticks, Item Search at arbitrary ilvl, gem/enchant sets, socket adds, tier-set requirement, talent loadouts, consumable combos, Great Vault items | ranked DPS table of gear × talent × gem combinations, "Your Top Gear" set, sidegrade grouping, Shopping List, Great Vault Summary |
| **Droptimizer** | `/simbot/droptimizer` | character + a loot source (raid / M+ dungeon / world boss / profession / delve / PvP / bonus-roll pool) + difficulty or key level | per-item **%DPS if you win and equip it**, **Expected Value** (avg gain over the group; non-upgrades count 0, not negative), **Best Drop**, **Gamble/Buy**, **Priority** (bosses grouped within 0.2% EV, ranked by chance-of-upgrade, tie-broken by Best Drop) |
| **Gear Compare** | `/simbot/gear` | two or more explicit gear sets | DPS per set. Demoted to "Legacy" 2024-04-23: *"Top Gear item search is a better tool in most cases."* |
| **Stat Weights** | `/simbot/stats` | one character | per-stat weights + a **Pawn string** |
| **Talent Compare** | `/simbot/talents` | talent loadout strings | DPS-ranked list |
| **Advanced** | `/simbot/advanced` | a raw SimC script (syntax-highlighted editor since 2026-04-09) | whatever the script asks for |

Plus **SimC Expert Mode** on every non-Advanced tool: raw SimC injected at 11 named positions (Script
Header, Base Actor, Consumables, Expansion Options, Custom APL, Pre-Actor, Actors, Post-Actor, Custom
Enemies, Simulation Options, Script Footer). *"Raidbots does not perform any validation on your input."*
([article 15](https://support.raidbots.com/article/15-simc-expert-mode))

### 3.2 The limits that define the product

- **Droptimizer changes one piece at a time.** *"it only changes one piece at a time, not simulating all
  possible combinations… It does not account for multiple gear changes or look further out than 1 drop."*
  Trinkets and rings are tried in both slots; dual-wielders try weapons in both.
  ([article 59](https://support.raidbots.com/article/59-droptimizer-how-does-it-work))
- **Top Gear combinatorics are hard-capped** *(config)*: `maxRawCombinations 200,000`,
  `maxCombinations 10,000`, `maxTotalCombinations 10,000`, `smartThreshold 2,000,000`, plus per-slot
  enhancement caps. Top Gear and Advanced bill every gear set as 5,000 iterations, so the tier iteration
  cap ÷ 5,000 is roughly your combination budget
  ([article 55](https://support.raidbots.com/article/55-smart-sim)).
- **Sidegrade** has an exact definition: *"a DPS result within `2 * error` of the top result"* — 0.1% at
  default precision ([article 61](https://support.raidbots.com/article/61-top-gear-sidegrades)).
- **Full talent trees are impossible and they say so**: *"The average size for just a full spec tree sim
  … is about 60 million combinations which is about 300 billion iterations. This is 18,750 times larger
  than the largest sim Raidbots can currently handle."*
  ([article 42](https://support.raidbots.com/article/42-talent-tree-combination-size))
- **Stat weights are actively discouraged by their own docs**: *"Friends don't let friends abuse stat
  weights… Direct sims (Top Gear, Droptimizer) are almost always better."* They evaluate only raw
  primary/secondary stats — no procs, set bonuses, enchants, or weapons
  ([article 66](https://support.raidbots.com/article/66-beware-of-stat-weights)). Method: run the
  character, then run it +238 of each stat, then `dps_increase / stat_increase`. 238 comes from "3.5%
  haste".
- **Healing and tanking are not really supported**: *"Party/raid healing is not simulated/quantified at
  all and survival/tanking metrics are generally unreliable."*
  ([article 69](https://support.raidbots.com/article/69-why-isnt-my-spec-supported))

### 3.3 Character import and the addon string

Raidbots prefers the addon and says why: *"The Blizzard Armory API is often out-of-date… if you have a
choice, use the addon!"* ([article 54](https://support.raidbots.com/article/54-installing-and-using-the-simulationcraft-addon))
The addon is **SimulationCraft's own** ([CurseForge](https://www.curseforge.com/wow/addons/simulationcraft),
[source](https://github.com/simulationcraft/simc-addon)), not a Raidbots addon. Commands: `/simc`,
`/simc nobags`, `/simc [Item Link]`.

The export string (emission order per `GetSimcProfile` in the addon's `core.lua`):

```
# <Name> - <Spec> - <YYYY-MM-DD HH:MM> - <region>/<realm>
# SimC Addon <version>
# WoW <version>.<build>, TOC <toc>

mage="Nickolaki"
level=70 / race=troll / region=eu / server=twisting_nether / role=spell
professions=alchemy=175/engineering=1
spec=arcane
talents=B4DArSxcnei16P8xFL3rzzOyRSCQLRCRKgkkQEJRSkAJBJtkAAAAAAAAAAAAikkkEJJJBCA
# Saved Loadout: Cleave
# talents=...                          <- inactive loadouts, commented

# Underlight Conjurer's Arcanocowl (441)
head=,id=202551,bonus_id=6652/9414/7977/...
finger1=,id=192999,enchant_id=6556,gem_id=192948,bonus_id=...
main_hand=,id=190511,enchant_id=6643,bonus_id=...

### Gear from Bags
# head=,id=202551,bonus_id=...          <- every bag item, COMMENTED OUT
### Weekly Reward Choices / ### Merchant items / ### Linked gear
### Additional Character Info
# catalyst_currencies= / upgrade_currencies= / slot_high_watermarks= / bonus_roll_items=
# Checksum: <adler32 hex>
```

Two design points worth stealing: **bag items are emitted commented out** so a plain SimC run ignores
them while Raidbots parses the comments to fill its item pickers; and the trailing **adler32 checksum**
line detects truncated or tampered pastes. Raidbots validates the paste with four regexes on the header
lines and enforces a minimum addon version *(config: SimC addon 12.1.0, WoW 12.1.0)*. Anything the
client has in memory leaks in — opening the vault, bank, or a vendor before `/simc` includes those items.

### 3.4 Report pages

`https://www.raidbots.com/simbot/report/<22-char base62 id>`, e.g.
[`.../report/imvT4F1wPbyFgceMdFfVHH`](https://www.raidbots.com/simbot/report/imvT4F1wPbyFgceMdFfVHH).
Sub-routes: `/simc` = SimC's native HTML report, `/banner` = character banner image, `/data.json` =
302 to `https://storage.googleapis.com/simbot-reports/reports/<id>/data.json`.

Sidebar controls: Share Report URL, Share Setup, Relative DPS toggle, **Run Sim Again** (restores the
original setup), Raw Files, original addon input, user custom APL, SimulationCraft Log, Full HTML
Report, SimC Error Info, Simple/Detailed toggle. Full HTML is suppressed above `htmlComboLimit: 20`
*(config)*.

**Retention conflicts between sources.** The [privacy policy](https://www.raidbots.com/privacy) says
*"sim reports are stored for 28 days"*; the tier config says 30 days free, **120 days** for all paid
tiers, 365 for internal tiers. Unreconciled — treat 120/365 as config-only. Shared setup links expire
after 30 days. Import history: 30 days Premium, 7 days free.

### 3.5 Premium: price and exactly what it gates

Prices (Stripe plan objects in the public bundle; USD and EUR numerically identical, unchanged since the
2018 launch):

| Tier | 1 mo | 3 mo | 6 mo |
|---|---|---|---|
| Rare | $3.00 | $8.25 | $15.00 |
| Heirloom | $5.00 | $13.50 | $24.00 |
| Epic | $10.00 | $27.00 | $48.00 |

The real limit table *(config `userLevels`)*:

| Tier | iteration cap | report retention | submits/hr | CPU budget/hr | sim time limit | advanced actors | concurrency | queue priority |
|---|---|---|---|---|---|---|---|---|
| anonymous | 300,000 | 30 d | 20 | 600 | 15 min | 60 | 1 | low |
| registered free | 300,000 | 30 d | 24 | 900 | 15 min | 60 | 1 | low |
| **Rare $3** | 2,000,000 | 120 d | 40 | 2,400 | 30 min | 120 | 1 | normal |
| **Heirloom $5** | 5,500,000 | 120 d | 40 | 4,800 | 30 min | 200 | 1 | normal |
| **Epic $10** | 16,000,000 | 120 d | 40 | 9,600 | 60 min | 200 | **2** | normal |
| Theorycrafter / SimC Dev (internal) | 16,000,000 | 365 d | 150 | unlimited | 120–480 min | 200 | 2 | normal + **API** |

`noAds: true` on every tier — the site is ad-free for everyone. API access is internal-only. The unit of
`simCpuBudgetPerHour` is **unverified**.

Premium buys exactly five things
([article 28](https://support.raidbots.com/article/28-raidbots-premium-rewards)):
1. **Skip the Line** (all tiers) — *"your sim jumps ahead of all free users… sim processing that starts
   almost immediately."* This is the headline value.
2. **More Top Gear iterations.**
3. **Premium Droptimizer features** — currently just *"Run all M+ dungeons in a single Droptimizer sim"*,
   plus the combined Bonus Rolls source and Top Gear's "Evaluate Bonus Rolls" (2026-08-11 changelog).
4. **SwiftSim** (Epic only) — *"your sim is run on double the usual amount of CPU cores (64 instead of
   32)"*; eligible above 4 combinations/profilesets; *"a sim that would normally take 10 minutes will
   finish in closer to 5 or 6"* ([article 63](https://support.raidbots.com/article/63-swift-sim)).
5. **Guild Discord bot skip-the-line** (Epic only).

**Free-tier rate limits arrived August 2026** — *"To prevent abuse and manage server costs, Raidbots now
imposes some rate limits on free sims as of August 2026… designed to allow a handful of bigger sims
(Droptimizer, Top Gear) per hour which should be enough for about 99% of free users."* Anonymous users
are limited **by IP**, so VPN and shared-ISP users collide.
([article 74](https://support.raidbots.com/article/74-sim-usage-limits))

### 3.6 Infrastructure (they have published a lot of it)

Components, from Seriallos' architecture posts
([2017 overview](https://medium.com/raidbots/raidbots-technical-architecture-303349d82784),
[Flightmaster 2018](https://medium.com/raidbots/raidbots-tech-flightmaster-440692101355),
[the main queue 2020](https://medium.com/raidbots/raidbots-tech-the-main-queue-bd56ad8393fa)):

- **Frontend** React SPA · **Web server** Node/Express — *"The critical architectural piece of the web
  server is what it doesn't do — run SimC. All it does is create a job."*
- **Worker** — pulls a job, spawns SimC, writes HTML/JSON/stdout to Cloud Storage, kills overruns.
- **Warchief** — singleton that scales the worker fleet (loop every 5 s), does queue maintenance, ships
  metrics to Discord and Datadog.
- **Flightmaster** — ~900 LoC Node service that splits profileset sims into chunks, runs them, merges
  the JSON, and drives Smart Sim's multistage. Progress is checkpointed in Datastore so a preempted
  Flightmaster resumes. This exists *because* workers run on **preemptible/spot instances** and a SimC
  process is atomic.
- **Queue**: Kue → **Bull** on Redis, single Redis instance. *"a first-in-first-out priority queue where
  paying Premium members skip ahead of free users."* Kue's per-job pubsub broadcast "has quadratic
  properties"; Bull "scales linearly with site load."
- **Cloud**: GCP (*"I initially chose GCP because of their low CPU pricing"*), Compute Engine + Datastore
  + Cloud Storage, Cloudflare in front. Confirmed live 2026-09-14: responses carry
  `server: cloudflare` + `via: 1.1 google` + `x-cloud-trace-context`, and report data 302s to
  `storage.googleapis.com/simbot-reports/`. Workers are 32-core AMD GCP instances *(config
  `worker.numCores: 32`; the SimC Stats page maps `t2d`/`n2d`→Zen 3, `c3d`→Zen 4, `c4d`/`n4d`→Zen 5)*.

**Smart Sim** is the cost-control algorithm and the thing most worth copying
([article 55](https://support.raidbots.com/article/55-smart-sim)): *"inspired by AutoSimc,"* three
stages at **1%, 0.2%, 0.05%** target error, chunk sizes `[8, 32, 256]`. *"Low precision sims can be run
stupidly fast. Instead of requiring 10,000 or more iterations, often only 100–200 iterations are
needed."* Cull the candidate set at 1% error, re-rank the survivors at 0.2%, finish the top few at 0.05%.

**Published scale figures:** 2017 peak ~500 worker CPUs, ~25k sims/day. Antorus launch Dec 2017: ~190k
sims in a day, ~5 sims/second, peak ~120 machines / 2,880 cores, "23 billion virtual boss fights." 2020:
*"over a million jobs a day with peaks of over a thousand sims running at the same time."*

**Live 2026 telemetry** (public feed,
`https://www.raidbots.com/static/analysis/simc-stats/overview.json`): 28-day error rate 0.51%; fight-style
mix **Patchwerk 85.3%**, DungeonSlice 9.6%, TargetDummy 1.4%, HecticAddCleave 1.3%, DungeonRoute 0.63%;
**77% of sims are single-target**; reference Quick Sim p50 ≈ 4.6 s CPU at ≈17.8k iterations
(**≈0.15 ms/iteration**), p50 memory ≈9.5 GB; all-sims p50 elapsed **36 s**, p90 **198 s**. Since
2026-07-28 they build SimC with PGO.

**Cost:** the only hard published figure is *"site currently costs ~$400/month"*
([Seriallos, Reddit, Jan 2017](https://www.reddit.com/r/wow/comments/5r2zms/raidbots_now_with_relic_compare_and_advanced_simc/)).
No current cost-per-sim or cloud bill is published — **unverified**.

### 3.7 Game versions

**Retail only, no Classic of any kind.** The app's `wowEnvs` config lists exactly four environments —
`live` (product `wow`, labelled "Retail"), `ptr` (`wowt`), `xptr` (`wowxptr`), `beta` (`wow_beta`). No
Classic product key appears anywhere in the bundle, and the class/spec data is current-retail (Evoker,
hero-talent `subTreeId`, per-spec `talentsTraitTreeId`). The reason is upstream: SimC does not support
Classic (§2).

**This is the gap foreversixty.gg is walking into.** Raidbots is the reference product for what a sim
site's surface should be, and it will not be competing in Forever.

---

## 4. Classic-era simulators — what exists, and what is reusable

Licences below were read from the actual `LICENSE` file via `gh api repos/<r>/license`, not from
GitHub's badge.

### 4.1 General-purpose engines

| Tool | Lang | License | Status (last commit) | Stars / forks / contributors |
|---|---|---|---|---|
| [wowsims/classic](https://github.com/wowsims/classic) | Go + TS | **MIT** | **Dormant** — 2026-07-24, **1 commit in 90 days** | 23 / 36 / 97 |
| [wowsims/sod](https://github.com/wowsims/sod) | Go + TS | MIT | Dormant — 2 commits in 90 days | 23 / 59 / 99 |
| [wowsims/tbc-new](https://github.com/wowsims/tbc-new) | Go + TS | MIT | **Very active** — 423 commits in 90 days | 19 / 39 |
| [wowsims/mop](https://github.com/wowsims/mop) | Go + TS | MIT | **Very active** — 545 commits in 90 days | 22 / 42 / 114 |
| [timhul/ClassicSim](https://github.com/timhul/ClassicSim) | C++ / Qt5 QML | **Custom LGPLv3 + field-of-use restrictions** (GitHub: NOASSERTION) | **Abandoned** — last *code* commit 2021-03-21 | 114 / 65 / effectively 1 |
| [Zwyk/ClassicCraft](https://github.com/Zwyk/ClassicCraft) | C# | MIT | Dormant (2024-08-18) | 18 |
| [sleep2death/vanilla](https://github.com/sleep2death/vanilla) | **Go** | GPL-3.0 | Dead (2023-05) | 0 |

### 4.2 wowsims/classic is no longer the SoD codebase

Worth stating plainly because the repo *description* still says "Season of Discovery": the commit
history through 2026 H1 is a deliberate vanilla migration ("found a few more sod leftovers like divine
storm and crusader strike that I removed", "started work to finish sod paladin migration", "atiesh
actually stacks in vanilla"). Structurally confirmed in §1.1 — `sod/sim/core` has `apl_values_rune.go`
and `spell_mod.go`; `classic/sim/core` has neither. **`wowsims/classic` is the Classic Era / Anniversary
sim.**

Its own spec status, read from `ui/core/launched_sims.ts` (the authoritative file, verified in my clone):

| Launched | Beta | Alpha | Unlaunched |
|---|---|---|---|
| Elemental Shaman (P2), Enhancement Shaman (P1), Warden Shaman (P1) | Balance Druid (P5) | Feral Druid (P2), Mage, Rogue, Hunter, Shadow Priest, Warlock, DPS Warrior, Tank Warrior (P1), Prot Pal, Ret Pal (P6) | Feral Tank Druid, Resto Druid, Resto Shaman, Holy Paladin, Healing Priest, **Raid sim** |

**Three specs launched out of 21 after ~2 years, and the repo is now dormant** — 1 commit in the last
90 days, measured `gh api repos/wowsims/classic/commits?since=2026-06-14`, against 545 for `mop` and
423 for `tbc-new`. That is the realistic yardstick for §6's estimates, and it is also the market opening — the one-off class sims below exist precisely because wowsims Classic
never finished.

### 4.3 timhul/ClassicSim — the deepest vanilla codebase, and it is dead

Event-driven C++ DES with a Qt5/QML **desktop** GUI. ~1,055 source files, ~86,000 LOC, 80 XML data
files. Layering: `Engine/` → `Queue/` → `CombatRoll/AttackTables/` (four tables: `MeleeWhiteHitTable`,
`MeleeSpecialTable`, `RangedWhiteHitTable`, `MagicAttackTable`) → `Mechanics/Mechanics.cpp` → per-class
`Class/<Class>/{Spells,Buffs,Procs,TalentTrees}/`. Integer rolls on 0–9999, xoroshiro128+.

Eight classes complete; **Priest is a stub** (no `TalentTrees/`, `Buffs/`, or `Procs/`). Tanking and
healing were roadmapped for v0.6/v0.7 and never shipped. Effectively a solo project — timhul 1,701
commits, next contributor 69.

**There was never a live web build**, and the domain is now hostile. The last commit on master,
[`558f84a` (2024-03-03)](https://github.com/timhul/ClassicSim/commit/558f84a), is titled "Remove expired
website link from README" and replaces `classicsim.org` with a GitHub Releases link. Today
`classicsim.org` 302-redirects to affiliate spam (`medpatceo.com/match-...`); `classicsim.io` is
NXDOMAIN. **Do not link either anywhere on foreversixty.gg.** Distribution was always desktop binaries;
last release v0.5-alpha-1, 2021-03-23.

**The licence is a trap.** `LICENSE` prepends a custom preamble to GPLv3:

> The distribution of this software … shall always: 1. Be free of charge 2. Only be provided as a
> digital product 3. Be restricted to the ®World of Warcraft IP **4. Be free of advertisements**
> 5. Include licensing information

"Free of charge" and "free of advertisements" are field-of-use restrictions; this is not OSI-free and
GitHub classifies it NOASSERTION. If foreversixty.gg ever carries ads or a paid tier, vendoring
ClassicSim code **or its data files** is a violation.

**What is worth taking anyway (as knowledge, not as files):**
- **`Mechanics/Mechanics.cpp`** — ~250 lines of pure functions, the densest vanilla formula bank
  anywhere: two-regime miss, dual-wield miss (`white_miss*0.8 + 0.2`), glancing chance and damage
  clamps, dodge/parry, **boss base armor 3750**, armour DR `armor/(armor + 400 + 85*clvl)`, melee crit
  suppression, the spell-miss level-diff table (…0.06 / **0.17** / 0.28 / 0.39), and piecewise-linear
  partial-resist tables. Re-derive it in Go **from the wiki prose**, not from the file (see below).
- **`Equipment/EquipmentDb/**/*.xml`** — **1,284 curated vanilla items with explicit proc rates.** The
  best curated proc dataset in existence, with `rate=`, `internal_cd=`, `mutex` exclusion groups, and
  `phase=`. 25 proc archetypes. Use it as a **validation oracle only** — the numbers are facts about the
  game, but the curation and structure are the author's licensed work.
- **[Mechanics Details wiki](https://github.com/timhul/ClassicSim/wiki/Mechanics-Details)** — a
  plain-English spec of every modelling decision, including the complete chance-on-hit proc ruleset
  (procs cannot recurse; extra attacks and Seal of Command *can* trigger others; extra attacks roll the
  full hit table, generate rage, renew seals, and reset the MH swing timer) and the two-roll spell table.
  **This page is prose describing game facts. Port from it freely.**
- **[Questions and Investigations wiki](https://github.com/timhul/ClassicSim/wiki/Questions-and-Investigations)**
  — a list of what they never resolved. Read it as a pre-made list of known-unknowns.
- Its `Rotation/Rotations/<Class>/*.xml` DSL is more human-writable than wowsims' JSON APL:
  ```xml
  <cast_if name="Heroic Strike">
      variable "time_remaining_execute" greater 3
      and resource "Rage" greater 50
  </cast_if>
  ```

`jokal2/CSIMDevTools` is **nothing** — three Python Qt-packaging scripts, 0 stars, dead 2019-11-04.

### 4.4 Sixty Upgrades — confirmed a gear planner, not a sim

[sixtyupgrades.com](https://sixtyupgrades.com), live, bundle last modified 2026-08-22. **No combat
engine, no rotation model, no Monte Carlo.** It computes Equivalency Points (a linear dot-product of
per-stat weights against item stats), a full character sheet with a *defensive* attack table (a live set
page shows `Boss Miss 4.40% / Boss Crit 5.60% / Boss Crush 15.00%` — level-63 values), a talent
calculator, and a greedy per-slot autofill optimizer. The author is upfront: *"these have their
shortcomings and do not remain constant as your balance of stats changes."*
([launch post](https://www.reddit.com/r/classicwow/comments/c1ncso/))

Versions: `/era`, `/sod`, `/tbc`, `/wotlk`, `/cata`. Not open source — GitHub search for
`sixtyupgrades` returns **0 repositories**. Its API (`https://api.sixtyupgrades.com/classic/graphql`,
AWS Cognito pool `us-west-2_LciI88NwD`) returns `403 MissingAuthenticationTokenException`
unauthenticated; no community client exists.

**But there is a clean, documented interop surface, and it is the cheapest ecosystem win available.**
Both halves are already implemented in the wowsims source:

- **Import (60U → sim):** [`individual_60u_importer.tsx`](https://github.com/wowsims/classic/blob/master/ui/core/components/individual_sim_ui/importers/individual_60u_importer.tsx)
  parses `{ character: {gameClass, race}, talents: [{spellId}], items: [{id, name, enchant:{id}, suffixId, reforge:{id}}] }`.
  Its own UI copy: *"This feature imports gear, race, and (optionally) talents. It does NOT import
  buffs, debuffs, consumes, rotation, or custom stats."* A live code comment notes 60U exports wrong
  random suffixes, so wowsims strips them.
- **Export (sim → 60U stat weights):** a plain unauthenticated GET, from
  [`individual_60u_ep_exporter.tsx`](https://github.com/wowsims/classic/blob/master/ui/core/components/individual_sim_ui/exporters/individual_60u_ep_exporter.tsx):
  `https://sixtyupgrades.com/sod/ep/import?name=<x>&strength=1.234&agility=…&spellPower=…`
  (note the live bug: the *Classic* repo hardcodes `/sod/`, not `/era/`).

Competitors in the planner space: [Stonetavern](https://stonetavern.app/planner) (positioned explicitly
against 60U: "No ads, no tracking"), [Wowhead Classic gear planner](https://www.wowhead.com/classic/gear-planner),
[Warcraft Tavern](https://www.warcrafttavern.com/wow-classic/tools/gear-planner/).

### 4.5 Class-specific sims and spreadsheets

| Tool | Class | Lang | License | Last commit | Stars |
|---|---|---|---|---|---|
| [GuybrushGit/WarriorSim](https://guybrushgit.github.io/WarriorSim/classic.html) | Warrior | JS | **NONE** (`/license` → 404) | 2024-12-16 | 57 / **92 forks** |
| [tzcnt/WarriorSim](https://fleetcode.com/WarriorSim/) | Warrior | C++→Emscripten | **NONE** | **2026-09-14** | 0 |
| [wow-aurana/bigdick](https://github.com/wow-aurana/bigdick) | Fury Warrior | JS | LGPL-3.0 | **2026-09-14** | 6 |
| [Cheesehyvel/magesim-vanilla](https://cheesehyvel.github.io/magesim-vanilla/) | Mage | **Rust**→WASM + Vue | GPL-3.0 (crate) | 2025-05-04 | 2 |
| [Cheesehyvel/magesim-tbc2](https://github.com/Cheesehyvel/magesim-tbc2) | Mage | C++→WASM + Vue | **MIT** | 2026-03-12 (active) | 10 |
| [ronkuby-mage/fire-mage-simulation](https://github.com/ronkuby-mage/fire-mage-simulation) | Mage (team ignite) | Python | **Apache-2.0** | 2025-08-17 | 12 |
| [watchyoursixx/HunterSim](https://watchyoursixx.github.io/HunterSim/) | Hunter | JS | **GPL-3.0** | 2022-07-16 | 5 |
| [Maarslet/WarlockSim](https://maarslet.github.io/WarlockSim/) | Warlock | JS | **MIT** | 2025-02-06 | 0 |
| [Kristoferhh/WarlockSimulatorTBC](https://kristoferhh.github.io/WarlockSimulatorTBC/) | Warlock | TS + C++/WASM | GPL-3.0 | 2022-07-16 | 10 |
| [TheSorm/RetSim](https://github.com/TheSorm/RetSim) | Ret Paladin (TBC) | C# desktop | **MIT** | 2022-05-30 | 12 |
| [NerdEgghead/classic_cat_sim](https://github.com/NerdEgghead/classic_cat_sim) (+ TBC / WotLK) | Feral Druid | Python | **NONE** | 2021-12 / 2023-08 | 1 / 3 / 12 |
| [Ceridwyn/tbcc-moonkin-dps-simulator](https://github.com/Ceridwyn/tbcc-moonkin-dps-simulator) | Balance (TBC) | Jupyter | Apache-2.0 | 2021-11 | 2 |
| Rogue sims ([system787](https://github.com/system787/ClassicRogueSim), [imwilson](https://github.com/imwilson/RogueSim), [pedrobefi](https://github.com/pedrobefi/Nordie-sims)) | Rogue | Java / C++ / — | — | 2020–2025 | 0 each |

Notes that matter:

- **GuybrushGit/WarriorSim has no licence file at all** (confirmed: `gh api repos/GuybrushGit/WarriorSim/license`
  → 404). It is the most-forked tool in the ecosystem (92 forks) and legally unusable. Its `gear/`
  directory is quietly the most valuable thing in it: **raw 1.12 DBC exports as CSV** —
  `itemsparse.csv` (7.8 MB), `spelleffect.csv` (3.1 MB), `spellmisc.csv`, `spell.csv`,
  `spellcategories.csv`, `itemeffect.csv`, `itemset.csv`, `shieldblockvalue.csv` — same licence problem.
- **`shadowcraft-classic` does not exist.** ShadowCraft is retail-only and itself dead
  ([Aldriana/ShadowCraft-Engine](https://github.com/Aldriana/ShadowCraft-Engine), 37★, LGPL-3.0, last
  push **2012**). Classic rogues use spreadsheets
  ([Classic Rogue Craft index](https://classicroguecraft.com/formulas-spreadsheets/)).
- **No Classic healing sim exists for any class.** No Classic Wrathcalcs analogue. Feral, rogue, and all
  healing specs are spreadsheet-only territory.
- Spreadsheets worth knowing (all live, `/edit`-accessible so formulas are extractable):
  [Konst's Warrior sheet](https://docs.google.com/spreadsheets/d/1DfYmb-ycl6hW8Xx_wfcx6tVscJ1kNyELxFeR5waLoZw/edit)
  (the canonical one; data from Magey's beta testing),
  [Infra's 1.13 Mage sheet](https://docs.google.com/spreadsheets/d/1uCXQni8Ndf0RuKfgNUlQ_jzfKIYmsCVuYXS_P3OaKUY/edit),
  [WatchYourSixx's Hunter sheet](https://docs.google.com/spreadsheets/d/1BIlB2P1kyV_QdD4ULQzvZvS6hK6BDouUQkyHQzCvBGI/edit).
  The names "Baranor", "Sarumon", "Aeriwen", "Bumples", "kmmagesim", "elu's mage sim" and
  "Warlock Optimizer" produced **no evidence** and appear to be misremembered.

### 4.6 Private-server and notable wowsims forks

All 32 forks of `classic` and all 58 of `sod` were enumerated; most are stale mirrors. The ones that
actually diverged:

| Fork | Target | Licence | Last push | What changed |
|---|---|---|---|---|
| [isfir/wowsims-turtle](https://isfir.github.io/wowsims-turtle/) | **Turtle WoW** | MIT | 2026-04-17 | The most substantive private-server fork. "Add Vampirism stat", "Implement Resilience and fix item effects parsing", Turtle item procs, patch data imports. |
| [Vellasta/classic](https://vellasta.github.io/classic/hunter/) | Turtle Hunter | MIT | 2026-04-13 | Hunter-focused retitle |
| [Zephryl87/classic](https://github.com/Zephryl87/classic) | Classic Era | MIT | 2025-05-05 | Era retitle, predates upstream's own migration |
| [denari94/EnhanceSim](https://github.com/denari94/EnhanceSim) | Enh Shaman | MIT | 2025-01-01 | Single-spec strip-down |

**Project Epoch and Ascension have no sim of any kind** — `gh search repos "project epoch wow"` returns
17 repos, all addons; Ascension returns zero. No published combat-formula documentation for Turtle WoW
or Project Epoch either; both are Discord-only.

### 4.7 The Forever field, as of today

**Four repos, all created in the last 48 hours, none with a licence, none with an engine worth the name:**

| Repo | Created | Lang | Notes |
|---|---|---|---|
| [tzcnt/WarriorSim](https://github.com/tzcnt/WarriorSim) | 2026-09-13 | C++→WASM | "Blazing fast Warrior DPS simulator for WoW Forever and Classic Era". Detached fork of GuybrushGit's. Reads talents from `talentsforever.com/warrior`; ships `data/forever/RACIALS.md` with tiered provisional values. Has a **distributed "Share Compute" browser worker pool** — genuinely novel, nobody else in this space has it. |
| [turlockmike/wow-forever-boomkin-sim](https://github.com/turlockmike/wow-forever-boomkin-sim) | 2026-09-14 | JS (Node) | Cast-by-cast Monte Carlo + talent optimizer, Balance Druid only |
| [andycarlson13/foreversims](https://foreversims.vercel.app) | 2026-09-14 | TS | Warlock, "evidence-tiered data. No guessing beyond tooltips." |
| [ClassicWoWCommunity/forever-bugs](https://github.com/ClassicWoWCommunity/forever-bugs) | 2026-09-14 | — | Community bug/research tracker — worth watching as shared research substrate |

**There is no `wowsims/forever` yet.** I checked the full `wowsims` org repo list on 2026-09-14: cata,
classic, sod, mop, tbc, tbc-new, wotlk, exporter, dev-docs, pages-deploy, ptr-deploy, wowsimsapp, YALPS,
DB2ToSqliteTool, AtlasLootClassic_SoD. Given they stood up `sod`, `classic`, `cata`, `mop`, and
`tbc-new` each within weeks of the corresponding launch, **assume `wowsims/forever` appears within
2–6 weeks of the 2026-09-17 beta.** That is a strategic fact, not a technical one — see §6.7.

### 4.8 Formula references worth citing

These are what to cite in a public methodology page; the community recognises them and showing your work
is how a new sim earns trust.

**Tier 1 — empirically tested, Classic-specific**
- **[magey/classic-warrior wiki](https://github.com/magey/classic-warrior/wiki)** — 116★, no licence,
  frozen 2019-05-28, class-agnostic despite the name. The
  [Attack-table page](https://github.com/magey/classic-warrior/wiki/Attack-table) is *the* reference:
  Blizzard-confirmed values, then their own tests **with sample sizes** (n=7,044 equal-level; n=31,779
  vs +3; n=28,063 vs +3 with +5 skill). Miss `5% + (TargetLevel*5 − AttackerSkill)*0.2%` for skill gap
  ≥11 with hit suppression `(gap−10)*0.2%`, or `*0.1%` with no suppression for gap ≤10; dual wield
  = normal + 19%; dodge `5% + gap*0.1%`; parry 14% vs +3; glancing `10% + gap*2%` → 40% vs +3; glancing
  damage low `1.3 − 0.05*gap`, high `1.2 − 0.03*gap`; crit suppression 1%/level **plus a flat ~1.8%
  applied only to crit from auras/talents/buffs**. Also
  [Crit-aura-suppression](https://github.com/magey/classic-warrior/wiki/Crit-aura-suppression) (16
  datasets, 2,376–8,075 attacks each, 95% CIs),
  [Parry-haste](https://github.com/magey/classic-warrior/wiki/Parry-haste),
  [Spell-batching](https://github.com/magey/classic-warrior/wiki/Spell-batching),
  [Threat-Mechanics](https://github.com/magey/classic-warrior/wiki/Threat-Mechanics),
  [Windfury-Totem](https://github.com/magey/classic-warrior/wiki/Windfury-Totem), and
  [issue #8](https://github.com/magey/classic-warrior/issues/8) documenting that 100% resists roll down
  to 75%, so the practical average-mitigation cap is **~69–70%, not 75%**.
- [ClassicSim Mechanics-Details wiki](https://github.com/timhul/ClassicSim/wiki/Mechanics-Details) —
  cross-check against magey; the glancing clamps differ slightly and are worth reconciling.
- [Marrow's Compendium of Dragonslaying](https://bookdown.org/marrowwar/marrow_compendium/mechanics.html)
  — numbered equations incl. the crit cap derivation; debunks "33% crit = 100% Flurry uptime".

**Tier 2 — Blizzard primary sources**
- [Not a Bug: Combat Table Values in Classic WoW](https://www.wowhead.com/classic/news/not-a-bug-combat-table-values-in-classic-wow-291996)
  (2019-05-29) — Blizzard ran a Reference-client diff. 14% parry vs +3, 5% dodge +0.5%/level, crit
  reduced 1%/level, mobs have `5 × level` weapon skill/defense, 0.04% per point of defense difference.
- **[Hit Cap in Classic WoW Clarifications](https://www.wowhead.com/classic/news/hit-cap-in-classic-wow-clarifications-292085)**
  (2019-06-03) — the single most important blue post for a DPS model: *"code in 1.12 that explicitly
  adds a modifier that causes the first 1% of +hit gained from talents or gear to be ignored against
  monsters with more than 10 Defense Skill above the attacking player's Weapon Skill… With a Weapon
  Skill of 305 … this hit modifier is no longer in place."* **This is exactly the interaction that
  Forever's unified Hit stat and retuned weapon skill will perturb, and nobody knows how yet.**

**Tier 3 — vanilla wikis** (Fandom 403s to curl and 402s to WebFetch; reach via `/api.php` or a browser)
- [vanilla-wow-archive/Hit](https://vanilla-wow-archive.fandom.com/wiki/Hit) — best single page for the
  miss formula; base 5% 2H / **24% dual-wield**; explains the 294→295 skill discontinuity; worked
  crit-cap example.
- [/Attack_table](https://vanilla-wow-archive.fandom.com/wiki/Attack_table) — one-roll precedence
  Miss → Dodge → Parry → Glancing → Block → Crit → Crushing → Hit.
- [/Attack_power](https://vanilla-wow-archive.fandom.com/wiki/Attack_power) · [/Resistance](https://vanilla-wow-archive.fandom.com/wiki/Resistance)
  (`(Res/(CasterLevel*5)) * 0.75`, cap 75%, 315 cap at L63) · [/Glancing_blow](https://vanilla-wow-archive.fandom.com/wiki/Glancing_blow).
- [/Armor](https://vanilla-wow-archive.fandom.com/wiki/Armor) — **use with care**: its DR section is
  tagged `{{source needed}}` and is the **2.0.8/TBC** formula. For vanilla use only
  `Armor / (Armor + 400 + 85 × AttackerLevel)` (L63 → `Armor/(Armor + 5755)`, cap 75%).
- [wowwiki-archive/Formulas:Instant_Melee_Attacks](https://wowwiki-archive.fandom.com/wiki/Formulas:Instant_Melee_Attacks)
  — normalized AP coefficients: two-handers **3.3**, daggers **1.7**, other one-handers **2.4**, cat paw
  1.0, bear paw 2.5.
- [wowwiki-archive/Spell_power](https://wowwiki-archive.fandom.com/wiki/Spell_power) — the fullest
  coefficient writeup. (Its `Spell_power_coefficient` page is **WotLK-patched** — do not use.)

**Spell coefficients**, corroborated across four sources: direct `C = CastTime/3.5` clamped [1.5, 7.0];
DoT/HoT `C = Duration/15` clamped at 15 s; hybrid `x = Dur/15, y = Cast/3.5 → C_DoT = x²/(x+y),
C_DD = y²/(x+y)`; AoE = half single-target; sub-20 penalty `1 − ((20 − SpellLevel) × 0.0375)`. In
original vanilla there was **no** downranking penalty beyond the sub-20 rule.
([Ozgar's Downranking Guide](https://www.warcrafttavern.com/wow-classic/guides/ozgars-downranking-guide-tool/)
is the cleanest Classic-targeted writeup.)

**Threat** — [Kenco's research on threat](https://wowwiki-archive.fandom.com/wiki/Kenco%27s_research_on_threat)
(Jan 2006, the primary source: 1 damage = 1 threat, healing 0.5/point excluding overheal, **rage
5.0/point**, warrior stance multipliers 0.8 / 1.3 / 1.45 with Defiance, measured raw ability threat);
the [2008 WoWWiki Threat snapshot](https://web.archive.org/web/20080101022824/www.wowwiki.com/Threat)
for per-class threat auras; [Wowhead's Classic Threat Overview](https://www.wowhead.com/classic/guide/threat-overview-classic-wow)
for the rank-by-rank table; and **[LibThreatClassic2](https://github.com/dfherr/LibThreatClassic2)**,
which is the machine-readable version (spell-ID-keyed threat tables) that ThreatClassic2 and Details
actually ship.

**Elitist Jerks** survives only as a Wayback snapshot of the
[Theorycrafting Think Tank](https://web.archive.org/web/20080927163119/http://elitistjerks.com/f47/t21302-theorycrafting_think_tank/)
(35 articles, all verified reachable). **Caveat: it is a 2.x/TBC snapshot.** EJ peaked in TBC and its
vanilla coverage was thin and is largely lost. The one genuinely valuable vanilla EJ thread still
reachable is [Resisting Magical Damage & Its Relation to Resistance Levels](https://web.archive.org/web/20110808083353/http://elitistjerks.com/f15/t10712-resisting_magical_damage_its_relation_resistance_levels/)
(2007) — the dataset behind the 69–70% effective resist cap.

**Raw game data**

| Source | Licence | Notes |
|---|---|---|
| [nexus-devs/wow-classic-items](https://github.com/nexus-devs/wow-classic-items) | **MIT** | 58★. All Vanilla/TBC/WotLK items, professions, zones as JSON. Frozen 2023-01. |
| [gtker/wow_vanilla_dbc](https://github.com/gtker/wow_vanilla_dbc) | **Apache-2.0** | Rust library for reading/writing **1.12 DBC** files |
| [vmangos/core](https://github.com/vmangos/core) | GPL-2.0 | 924★, active. Progressive vanilla emulator — a large second opinion on every formula. **Viral licence: read, don't copy.** |
| [cmangos/mangos-classic](https://github.com/cmangos/mangos-classic) | GPL-2.0 | 1,079★, active. Same caveat. [cmangos/issues#1193](https://github.com/cmangos/issues/issues/1193) reverse-engineers the resistance tables. |
| [wowsims/exporter](https://github.com/wowsims/exporter) | **MIT** | The in-game Lua addon — the de-facto standard character export for Classic. Pushed 2026-08-05. |

---

## 5. Data availability for Forever

### 5.1 What Blizzard has confirmed about the ruleset (verified, 2026-09-12/13)

From the [Deep Dive Panel Recap](https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap) (Blizzard, 2026-09-13) — direct quote:

> "World of Warcraft: Forever keeps the familiar stats from Classic, but reworks several of them and
> adds new ones… **Spell, melee, and ranged hit chance are combined, as are spell, melee, and ranged
> critical chance.** … **Weapon skill still works as it always has, but items with weapon skill offer
> less of it per item** so they are not always the obvious best choice. **Bonus healing now also includes
> one third as much bonus damage** … Some items can **reduce the chance for attacks to be parried or
> dodged**, while **spellcaster-oriented weapons now grant spell damage and healing**."

[Wowhead's liveblog](https://www.wowhead.com/forever/news/world-of-warcraft-forever-deep-dive-liveblog-382862) adds the word explicitly:
> "Weapon skill was a little too easy to get and a little too strong in Classic. It's a little harder to
> get now. **Expertise has also been added.**"

[Massively OP](https://massivelyop.com/2026/09/13/blizzcon-2026-world-of-warcraft-forevers-deep-dive-panel-talks-group-play-progression-and-item-updates/) corroborates: "stats like expertise have been added in."

Talents, from [Output Lag's transcription](https://outputlag.com/news/world-of-warcraft-forever-adds-a-fourth-one-point-talent-at-16-points-to-every-tree/) of Kris Zierhut:
> "All classes still have three talent trees, they still have seven rows, they still have those one
> point talents at 11, 21, 31, the gold medal talents, we called them. We've even added a fourth gold
> medal talent at the 16 mark at every single one of the talent trees." … "Improved Battle Shout, Divine
> Spirit, Blessing of Kings, Improved Mark of the Wild. They're all gone. All of those bonuses are now
> baseline to the class."

27 talent trees, all intended raid-viable. The number of talent points at 60 was **not stated**, and the
16-point talents were **not named for any class** — unverified until beta.

Other confirmed-relevant facts: level cap 60; new race (Skyborne Elves); six new race/class combos
(Gnome Priest, Human Hunter, Dwarf Shaman, Orc Mage, Troll Warlock, Undead Paladin); nine new dungeons,
two new raids at launch, more on 2026-12-09; reworked racials (two active + two passive each);
trinkets specialised with biome- and creature-type-conditional effects; every dungeon drop re-tuned.

### 5.2 What data will actually be available, and when

**Wowhead already has a live Forever data environment.** Verified by direct fetch on 2026-09-14:

```
GET https://nether.wowhead.com/forever/data/gear-planner?dv=100    → 200, 5,058,446 bytes
     payload keys: WH.setPageData("wow.gearPlanner.classicplus.item", ...)
GET https://nether.wowhead.com/forever/tooltip/item/2825?dataEnv=100&locale=0 → 200, JSON tooltip
```

The internal data-env name is **`classicplus`**, and the payload exposes exactly the sections WoWSims'
`wowhead_db.go` parses: `item`, `gem`, `itemSet`, `enchant`, `randomEnchant`, `randPropPoints`,
`baseStats`, `statToRating`, `baseCritPhysical`, `critPhysical`, `baseCritSpell`, `critSpell`,
`weaponProficiencies`, `armorProficiencies`, `talent`, `runeSpells`, `buffs`, `buffAuras`.

**Caveat, and it matters:** as of 2026-09-14 that endpoint still serves *Classic-Era-shaped* data.
Items carry `versionNum: 11300` (patch 1.13.0) and the old split stat keys — `mlehitpct`, `rgdhitpct`,
`mlecritstrkpct`, `splcritstrkpct` — not a unified hit/crit or an expertise key. The environment is a
scaffold; the real Forever payload lands when the beta client ships on 2026-09-17. So: **the pipeline
endpoint is confirmed to exist and to be the same shape WoWSims already parses; the Forever content in
it is not yet there.**

**wago.tools** serves DB2 as CSV at `https://wago.tools/db2/<Table>/csv?build=<x.y.z.build>` — which is
already precisely what `wowsims/classic` calls (`gen_db/main.go:70`:
`https://wago.tools/db2/ItemSparse/csv?build=1.15.3.55646`). Verified working against the live Classic
Era build `1.15.9.69547` on 2026-09-14; every table a sim would want returns 200:

| Table | HTTP | bytes |
|---|---|---|
| SpellName | 200 | 813,507 |
| SpellEffect | 200 | 3,658,058 |
| SpellMisc | 200 | 2,625,060 |
| SpellPower | 200 | 123,103 |
| SpellCooldowns | 200 | 97,569 |
| SpellCategories | 200 | 295,698 |
| SpellClassOptions | 200 | 164,131 |
| SpellLevels / SpellDuration / SpellCastTimes / SpellRange / SpellAuraOptions / SpellTargetRestrictions | 200 | — |
| ItemSparse | 200 | 8,561,018 |
| ItemEffect | 200 | 645,103 |
| ItemDamageOneHand | 200 | 10,979 |
| Talent | 200 | 31,494 |
| **SpellScaling** | **404** | `{"errors":"Table not found."}` |

`SpellScaling` does not exist in the Classic-lineage schema — expect the same for Forever. Spell
coefficients will have to come from the classic formula tradition (cast-time/3.5 for direct, duration/15
for DoTs, with per-spell overrides), not from a table.

wago.tools' [builds list](https://wago.tools/builds) shows it auto-discovers Blizzard CDN branches
within hours (`wow_classic_era`, `wow_anniversary`, `wow_classic`, `wow_classic_titan`, `wowt`, …). A
`wow_classic_titan` branch at patch 3.80.2 exists (build 69496, detected 2026-08-25), but I checked its
tables and it is **not** Forever: max spell id 1,316,372, Death Knight `TalentTab` rows, Mastery columns
in `ChrClasses`, no "Seal of Fury", no Skyborne. **Which product code Forever ships under is unverified**
— but wago will have it within a day of the beta going live, and `gen_db` needs a one-line URL change to
consume it.

### 5.3 What the tables give you, and what they do not

**DB2 gives you, reliably:** spell names and icons, school, mechanic, dispel type, GCD category
(`SpellCategories.StartRecoveryCategory`), cooldown and category cooldown (`SpellCooldowns`), mana cost
(`SpellPower`), cast time bucket (`SpellCastTimes`), range, duration index, the spell-family bitmask
(`SpellClassOptions.SpellClassMask_0..3` — invaluable for "which spells does this talent modify"),
effect type / aura type / base points / `EffectBonusCoefficient` / radius / target (`SpellEffect`),
attribute flags (`SpellMisc.Attributes_0..N`, which encode things like "not affected by haste", "cannot
crit", "ignores armor"), talent grid position and rank spell ids (`Talent.TierID, ColumnIndex, TabID,
SpellRank_0..`), and the whole item universe (`ItemSparse`, `ItemEffect`, weapon damage tables).

Rebuilding the *gear planner* from DB2 + Wowhead is a largely-solved, mechanical job. Call it a data
engineering problem.

**DB2 does not give you:**

1. **Proc mechanics.** `ItemEffect.TriggerType` tells you "on equip / on use / on proc" and
   `SpellID`, but not the proc chance, the PPM, or whether the trigger is on-hit / on-crit /
   on-spell-cast / on-being-hit. Forever's biome- and creature-type-conditional trinkets are worse:
   the condition lives in `PlayerConditionID` and/or in spell script, not in a readable column.
2. **Hidden internal cooldowns.** ICDs are frequently script-side. `SpellCooldowns` carries the visible
   cooldown; a trinket's 45-second hidden ICD is nowhere.
3. **AP/SP coefficients.** `SpellEffect.EffectBonusCoefficient` exists in the modern schema but for
   Classic-lineage spells is routinely 0 or wrong; vanilla coefficients are a *convention*
   (`cast_time/3.5`, `duration/15`, halved for hybrid classes, with dozens of hand-known exceptions)
   rather than data. Forever's new abilities will have new coefficients that nobody knows on day one.
4. **Actual talent behaviour.** `Talent.Description_lang` is prose. "Twist of Light" and "Seal of Fury"
   will have a tooltip and a spell id; whether Seal of Fury's absorb scales with AP or with spell power,
   whether it stacks, whether its Judgment-taunt shares a cooldown — none of that is in a table.
5. **Interaction rules.** Whether a new proc can trigger itself, whether two new absorbs stack, whether
   Holy Strike's threat is a flat bonus or a multiplier, how the unified Hit stat interacts with the
   vanilla weapon-skill miss table (this one is genuinely load-bearing for every melee spec and is
   **completely unspecified** by anything public).

### 5.4 How the two engines bridge that gap — and what it means for us

**SimulationCraft** extracts DB2 into `engine/dbc/generated/*.inc` and then writes the *behaviour* in
C++ class modules on top. The data supplies constants; the module supplies semantics. See §2.

**WoWSims** does not extract spell data at all. It hardcodes the constants **inside** the behaviour
(`bonusDamage := 160.0`, `Duration: time.Second * 6`) and uses DB2/Wowhead only for the item database
and the UI's icons and tooltips. §1.3.

Both therefore converge on the same truth: **the sim is the hand-written spell implementations, and
those are validated against combat logs.** WoWSims' validation loop is the golden-`.results` regression
corpus (§1.8) plus the community per-spec ownership model; SoD's 271 `TODO/unconfirmed/needs testing`
comments are the visible residue of shipping guesses and fixing them against logs later.

For Forever specifically:

- There is **no historical data to validate against**, by construction. Every Forever coefficient,
  proc rate, and ICD will be unknown on 2026-09-17 and only partially known on 2026-11-04.
- The vanilla *base* is 20 years theorycrafted and Blizzard has said combat "feels the same" — so the
  attack table, armor mitigation, resistance, glancing blows, dual-wield penalty, weapon skill, rage
  generation, energy ticks, and the spell coefficient conventions all carry over. That is the large
  majority of the engine, and it is exactly what a wowsims fork gives you for free.
- The *divergence* is concentrated in: the unified Hit/Crit stat and its interaction with weapon skill,
  expertise values and caps, per-class talent rewrites, new baseline abilities, reworked racials, and a
  fully re-itemized world with conditional trinkets.
- **Our own logs engine is the decisive asset here.** WoWSims has no log-ingestion path; validation is
  manual and social. If foreversixty.gg can, on day one of beta, automatically extract observed
  coefficients, proc rates, hit tables, and ICDs from real combat logs and diff them against the sim's
  predictions, we close the loop that WoWSims closes by hand over months. That capability — not the
  engine — is the defensible part of the product, and it argues for an architecture where the sim's
  event model is the *same* event model the logs engine already speaks (§6, option C's one real merit).


---

## 6. Recommendation inputs

### 6.0 How these estimates are built

Units are **engineer-weeks** for one competent Go engineer working with heavy AI assistance, which is
how this codebase is being built. They are not calendar weeks unless one person is on it. Every figure
is anchored to a measured line count or a measured historical analogue, cited inline. Where I am
guessing, I say so.

**SimulationCraft is not one of the options, and §2 says why in two sentences.** It has never
supported Classic (no `classic` branch; the "EPIC: Classic Support" issue has 1 of 15 boxes ticked
since 2021; zero of 787 forks default to a Classic branch), and its GPL-3.0 licence means compiling it
to WASM and serving it to a browser **is conveying**, which would pull our whole client under GPLv3.
What we take from SimC is ideas, not code: its text APL, its self-expiring hotfix pattern, and
Raidbots' precision-laddering economics.

Two calibration anchors, both measured in §1 and §4:
- **SoD's rules divergence cost wowsims ~5,252 LOC of new per-class ability code (~580/class) plus 790
  LOC of core `spell_mod.go` infrastructure**, and doubled every spec's size, over ~2 years with ~100
  contributors.
- **wowsims/classic has 3 specs "Launched" out of 21 after ~2 years** (`ui/core/launched_sims.ts`),
  and is now **dormant — 1 commit in 90 days** against 545 for `mop` and 423 for `tbc-new`. Any
  estimate that implies we finish 21 specs quickly is wrong. (The dormancy cuts both ways for a fork:
  almost no upstream churn to merge, but also almost no upstream fixes to inherit.)

The dominant cost in every option is **not writing code — it is determining unknown numbers.**
`mortal_strike.go` is 56 lines; the hard part is knowing that the bonus damage is 160 and the rage cost
is 30. For Forever, nobody knows any of those numbers yet.

### 6.1 Option A — fork `wowsims/classic`, port to Forever's rules

**Engine adaptation: 5–7 weeks**

| Task | Weeks | Justification |
|---|---|---|
| Merge `MeleeHit`+`SpellHit` → `Hit`, `MeleeCrit`+`SpellCrit` → `Crit` | 1.0 | 159 call sites in `sim/` (measured), 0 in `ui/`; must stay index-synced with `proto.Stat` |
| Regenerate `base_stats_auto_gen.go` (crit-per-agi, per-class-per-level bases, expertise constants) | 0.5 | `tools/base_stats_parser.py` exists; blocked on beta data |
| Bonus healing → ⅓ spell damage | 0.2 | `sim/core/stats/deps.go` is already a stat-dependency engine; one entry |
| **Expertise** | **0** | **already implemented** — `stats.Expertise` + `spell_outcome.go:711-735` reduce dodge and parry |
| **Weapon skill** | **0** | already fully modelled — 15-value enum, 52 call sites; Forever keeps it |
| **Spell damage / healing on caster weapons** | **0** | `stats.SpellDamage` / `stats.HealingPower` already exist; item-data change only |
| Backport `spell_mod.go` from SoD | 0.5 | 790 LOC, drops in cleanly — the cores differ by 4 files |
| New item stat keys (dodge/parry reduction) through the scrape → proto → sim path | 0.5 | `wowhead_db.go` + `wowhead_tooltips.go` stat-key maps |
| Biome / creature-type conditional item effects | 1.5 | **net-new mechanic class.** Needs a zone/biome on the encounter and a creature type on `proto.Target`, plus UI. No precedent in the engine. |
| Rework `sim/core/racials.go` (2 active + 2 passive per race, new Touch of the Grave, Stoneform now % phys reduction, Mace Spec now affects spell crit) | 1.0 | one ~300-line file, full rewrite |
| Slack for the unified-Hit × weapon-skill interaction (the [1% hit suppression rule](https://www.wowhead.com/classic/news/hit-cap-in-classic-wow-clarifications-292085) meets a merged stat — genuinely unknown) | 1.0 | see §4.8; this is the single most load-bearing unknown for every melee spec |

**Data pipeline: 3–4 weeks initial, then continuous**

| Task | Weeks | Justification |
|---|---|---|
| Repoint `gen_db` at Forever sources | 0.2 | **Verified live today**: `https://nether.wowhead.com/forever/data/gear-planner?dv=100` returns 5 MB keyed `wow.gearPlanner.classicplus.*`; wago.tools CSV API already used at `gen_db/main.go:70`. Both are one-line changes. |
| Absorb Wowhead `classicplus` schema drift (unified hit/crit keys, expertise, dodge/parry reduction) | 1.0 | today's payload still has `mlehitpct`/`splcritstrkpct`; the real shape lands 2026-09-17 |
| Replace the AtlasLoot loot-source stage | 1.5 | `tools/database/atlasloot.go` (434 LOC) depends on the AtlasLoot addon, which will have no Forever data. Every dungeon drop re-tuned, 9 new dungeons, 2 new raids. Needs a Wowhead zone-page scraper instead. |
| Rebuild `overrides.go` (439) + `enchant_overrides.go` (232) | 1.0 + ongoing | both are vanilla-specific hand curation and mostly wrong for Forever |

**First spec: 2–3 weeks.** Warrior as worked example: `talents.go` is 498 LOC of which I estimate
40–60% survives (Forever keeps seven rows, the 11/21/31 talents, and "many talents exactly the same as
they were, unchanged from 2006"); 2–3 new baseline-ability files at ~60 LOC each; one APL; one gear
set; one golden `.results` file. **This produces "credible", not "trusted".** Trust needs log validation
(§6.6).

**Every spec: 30–40 engineer-weeks, and not achievable solo.** 21 sim entries across 9 classes. Post-
first-spec learning amortises to ~1.5–2 weeks/spec. Realistic sequencing: DPS-only first (~12 specs,
**18–25 weeks**), tanks next, healers last or never — Raidbots itself says *"Party/raid healing is not
simulated/quantified at all"* (§3.2), so healing is not table stakes.

**UI: 3 weeks (seam) or 12–20 weeks (native)**
- Per-spec UI is tiny: measured 180–544 TS LOC per spec directory. `ui/core/` is 36,459 LOC and
  survives wholesale.
- **Cheap path (~3 weeks):** ship the wowsims Vite SPA at `foreversixty.gg/sim/*`, rebranded. Works,
  but it is a Bootstrap SPA bolted onto an Astro + Svelte-islands site, and UX is stated priority #1.
- **Native path (12–20 weeks):** rewrite the UI in Svelte against the same protobufs. Tractable
  *because* the protos are the entire contract and the TS client is generated (`makefile:82`) — you are
  not reverse-engineering anything. This is where the product differentiation actually lives.

**Total to a shippable single-spec product: ~11–15 engineer-weeks** (engine 6 + data 3.5 + spec 2.5 +
seam UI 3, with some overlap).

### 6.2 Option B — fork `wowsims/sod`

Identical to A except for the delta, and the delta is **net negative**:

**What B buys you**
- `spell_mod.go` (790 LOC) — but this is a 0.5-week cherry-pick into A, since the cores differ by four
  files.
- 166 checked-in `.apl.json` presets vs classic's 27 — but they are SoD rotations for SoD abilities.
  Useful as worked examples, not as content.
- A demonstrated pattern for layering new abilities onto a vanilla base (`ApplyRunes()` alongside
  `ApplyTalents()`), which is exactly the shape Forever's reworked talents want. This is real, but it is
  a pattern you read in an afternoon, not code you keep.

**What B costs you**
- **~5,252 LOC of rune code to delete**, and it is entangled, not isolated: `sim/warrior/runes.go`
  carries comments like "Furious Thunder implemented in thunder_clap.go" and "Gladiator implemented on
  stances.go". The rune concept also reaches into `proto/<class>.proto` (`WarriorRune` enums), the
  equipment model (`warrior.Equipment.Shoulders().Rune`), `apl_values_rune.go`, and the UI.
- SoD specs are **~2× the size of classic's** (warrior 5,999 vs 3,515; warlock 8,418 vs 4,029; paladin
  6,660 vs 2,381), so every file you touch is bigger and noisier.
- Phase-structured content dead weight: `item_sets_pve_phase_4.go` … `_phase_8.go`.
- SoD is dormant (2 commits in 90 days; upstream effort has moved to `mop` at 545 and `tbc-new` at
  423), so you inherit a branch nobody is maintaining.

**Delta vs A: +2–3 weeks of deletion, buys ~0.5 week.** B is strictly worse.
**The correct move is A, cherry-picking `spell_mod.go` and the rune-layer pattern from B.**

*(For completeness: `wowsims/mop` has the most modern core and the most activity — 20,771 commits, 114
contributors, pushed yesterday — but it is MoP. No weapon skill, no vanilla attack table, mastery and
rating conversions throughout, a completely different talent system. For Forever the **vanilla combat
model is the asset**, and only `classic` and `sod` have it. `tbc-new` is the second-best base: vanilla-
adjacent, MIT, and the most actively maintained Classic-lineage repo — but TBC already merged hit ratings
and dropped some vanilla quirks, so it is further from Forever than `classic`, not closer.)*

### 6.3 Option C — own engine in Go, sharing the logs engine's event model, with a text APL

| Task | Weeks | Justification |
|---|---|---|
| Engine core (event heap, auras, attack table, resources, metrics) | 10–16 | wowsims' `sim/core` is 29,861 LOC, of which ~4,127 (`buffs.go` 1,893 + `consumes.go` 1,252 + `debuffs.go` 982) is content, not engine. The irreducible engine is ~12–15k LOC. But LOC is the wrong measure: the vanilla attack table (two-roll vs one-roll, glancing chance *and* damage clamps, dual-wield miss, weapon skill vs defense, the 1% hit suppression, spell resistance partials, armour DR) is a set of **known-solved, subtle, empirically-derived rules** you would re-derive and re-bug. `spell_outcome.go` is 989 lines and `spell_school_test.go` alone is 494. |
| Text APL: lexer, parser, expression evaluator, action/value vocabulary | 4–6 | SimC's expression parser is the reference (§2). A text DSL is genuinely **better** than wowsims' proto tree — shareable in a forum post, diffable in git, greppable, and it is what the SimC-literate audience already knows. This is the one place Option C's output is superior. |
| Data pipeline | 3–4 | identical to A; no saving |
| First spec | 2–3 | identical to A; the cost is the unknown numbers |
| Every spec | 30–40 | identical to A |
| UI (Svelte-native from day one) | 12–20 | same as A's native path, minus the fork seam |

**Total to a shippable single-spec product: ~25–37 engineer-weeks** — roughly **+14–22 weeks over A.**

**The one stated merit needs examining honestly.** "Sharing our logs engine's event model" does not
survive contact with the code. `logs/engine/event/` (999 LOC) defines a **decoded Blizzard combat-log
line**: `Source/Dest Unit`, `Spell`, `Amount`, `Overkill`, `Absorbed`, `MissType`,
`Critical/Glancing/Crushing/OffHand`, `Combatant` with gear and talents. A simulator's internal events
are `PendingAction`s on a heap — aura ticks, swing timers, GCD expiry. **These are different models and
unifying them saves nothing.**

What *is* genuinely valuable, and is the real insight buried in option C:

> Have the sim **emit** `logs/engine/event.Event` values, so `logs/engine/summary` (1,395 LOC —
> damage, healing, deaths, casts, auras, resources) renders a simulated fight and a real logged fight
> through one code path, and one detailed-results UI serves both.

That is an **adapter, 1–2 weeks** — and it is available in options A and B just as much as in C. It is
also a real product differentiator (§3.7: WoWSims' result protobuf has no timeline at all, only
aggregated metrics plus a `max_seed`/`min_seed` replay trick). **Take this idea; do not take the
engine rewrite it was attached to.**

### 6.3b Three cheap ideas to steal, in any option

These are small, they are independent of A/B/C, and each one is disproportionately valuable.

1. **A text APL front-end over the proto APL — 2–3 weeks.** SimC's APL is text
   (`actions+=/spell,if=buff.x.up&cooldown.y.remains>3`); WoWSims' is a protobuf tree edited by a GUI
   (§1.4, §2.2). Text is postable in a Discord message, diffable in git, greppable, and is the form the
   theorycrafting audience already knows — which is exactly why Raidbots ships raw-SimC "Expert Mode"
   on every tool (§3.1). WoWSims' `APLAction`/`APLValue` oneofs are a perfectly good *AST*; writing a
   lexer + parser that targets them, and a pretty-printer that goes back, gives us both the builder UI
   and shareable text. Nobody in the Classic space has this.
2. **SimC's self-expiring hotfix pattern — days.** `hotfix::register_effect(...).field("base_value")
   .operation(HOTFIX_SET).modifier(2).verification_value(-4)` (§2.4): a dated override that carries the
   value you *expected* the data to have, and disables itself once the real data matches. For a game
   where every number is a guess on 2026-09-17, a first-class "this is our provisional value, here is
   what the tooltip said, here is when we guessed it" record is the right primitive — and it feeds the
   evidence-tier methodology page in §6.6.
3. **Smart-Sim precision laddering — 1 week, and it is the whole cost model.** Measured in §2.6: Monte
   Carlo error goes as 1/√n, so iterations go as 1/error². Going 0.2% → 0.05% is **16×** the work;
   culling candidates at 1% first costs **1/25th** of a 0.2% run. Raidbots' three-stage
   1% → 0.2% → 0.05% ladder with chunk sizes 8/32/256 (§3.6) is the single highest-leverage thing to
   copy, and it applies to bulk gear sims, Droptimizer-style runs, and stat weights alike.

### 6.4 The calendar, which probably decides it

Today is 2026-09-14. Beta **9/17** (3 days). Launch **11/4** (7 weeks). First raids **12/9** (12 weeks).

| Milestone | Option A | Option C |
|---|---|---|
| Beta week (0–1 wk) | Data pipeline pointed at Forever, gear planner sketch | Nothing shippable |
| Launch, 11/4 (7 wks) | Gear planner + 1–2 DPS specs, credible | Engine still in progress |
| First raids, 12/9 (12 wks) | 4–6 DPS specs, stat weights, bulk sim | First spec, maybe |

**No option ships all specs by launch.** But A can have something real in front of users on 11/4 and C
cannot. For a community site whose value compounds with early adoption during a launch window, that is
close to decisive.

### 6.5 What I measured, and how (so the estimates are auditable)

- Cloned `wowsims/classic` and `wowsims/sod` at HEAD (shallow, blob-filtered); all LOC figures are
  `wc -l` on those trees.
- Installed `protobuf` + `protoc-gen-go`, ran `protoc -I=./proto --go_out=./sim/core ./proto/*.proto`
  (`makefile:197`), and built `./sim/...`. Everything compiles except `sim/web` (needs a generated
  `binary_dist`).
- **Discovered that `go test ./sim/...` fails with `No item with id: N` unless you pass
  `--tags=with_db`** (`sim/core/database_load.go` is behind that tag; `makefile:221` passes it). Worth
  knowing before anyone wastes an afternoon.
- Wrote a throughput test in `sim/warrior/dps_warrior/` using the repo's own `GetGearSet` /
  `GetAplRotation` / `FullBuffs` fixtures, ran P1 BiS Fury Warrior, 300 s single target, full
  buffs/consumes, shipped `dps_reck` APL. Native: **1,218 iters/sec/core**. Built the same code with
  `GOOS=js GOARCH=wasm` and ran the identical serialised request through `raidSim` under Node 22:
  **155 iters/sec/core, a 7.9× penalty** (§1.10). Artefact sizes: 18 MB wasm / 3.53 MB gzipped, plus
  4.9 MB `db.bin` / 420 KB gzipped.
- Probed `nether.wowhead.com/forever/...` and `wago.tools/db2/<Table>/csv?build=...` directly (§5.2),
  including checking whether the `wow_classic_titan` branch is Forever (**it is not**).

### 6.6 The thing that actually differentiates us

Neither wowsims nor SimC has a log-validation loop. WoWSims validates by hand, socially, per-spec, over
months — SoD's 271 `TODO/unconfirmed/needs testing` comments are the visible residue (§1.9).

We already have `logs/` (10,495 Go LOC: lexer, layout inference for classic *and* retail formats, fight
segmentation, event decode, summary aggregation, parquet, store). On beta day, that engine can start
extracting **observed** coefficients, proc rates, ICDs, and hit/crit/dodge tables from real Forever
combat logs and diffing them against whatever the sim currently predicts. That closes, automatically and
continuously, the loop that every other sim closes by hand.

Concretely, three things worth building regardless of which option wins:
1. **Sim → `event.Event` adapter** (1–2 weeks) so one detailed-results UI serves logs and sims.
2. **Log → parameter extractor** — observed proc rate, observed ICD, observed coefficient per ability,
   with sample sizes and confidence intervals in the style of
   [magey/classic-warrior](https://github.com/magey/classic-warrior/wiki/Crit-aura-suppression). This
   is publishable research that earns community trust, which is the currency in this space.
3. **A public methodology page** citing sources per mechanic, with an explicit
   confidence tier per number (`tzcnt/WarriorSim`'s `data/forever/RACIALS.md` and
   `andycarlson13/foreversims`' "evidence-tiered data. No guessing beyond tooltips." show the
   community is already converging on this norm — §4.7).

### 6.7 Two strategic facts

1. **`wowsims/forever` does not exist yet, but it will.** The org has stood up `sod`, `classic`, `cata`,
   `mop`, and `tbc-new` each within weeks of the corresponding launch (§4.7). Assume it appears 2–6
   weeks after 2026-09-17, with ~100 contributors behind it. Forking means we can merge their fixes and
   contribute back rather than race them. Writing our own engine means competing head-on with a hundred
   volunteers on their home turf, and losing.
2. **Raidbots will not be in this market.** Retail-only, because SimC is retail-only (§2, §3.7). The
   best-in-class sim *product* surface — Droptimizer, Top Gear, Smart Sim's three-stage precision
   ladder, the addon import flow — has no Classic incumbent. The product design is the opportunity; the
   engine is a commodity we should acquire, not build.

### 6.8 Final comparison

| | **A — fork wowsims/classic** | **B — fork wowsims/sod** | **C — own Go engine** |
|---|---|---|---|
| **Licence** | MIT. Commercial hosted fork is fine; keep the copyright notice, add the requested user-visible backlink. | MIT, identical. | Ours. But the reference material we'd learn from (ClassicSim) is field-of-use restricted, and the emulators (vmangos, cmangos) are GPL-2.0 — read, don't copy. |
| **Engine adaptation** | **5–7 wk.** Expertise, weapon skill, spell-damage stats, the whole vanilla attack table already correct. Merging hit/crit is 159 mechanical call sites. | 7–10 wk. Same work plus ~5,252 LOC of entangled rune code to delete. | **10–16 wk** to re-derive a combat model that already exists and is right. |
| **Data pipeline** | **3–4 wk.** `gen_db` already calls the exact endpoints (`nether.wowhead.com/<env>/data/gear-planner`, `wago.tools/db2/…/csv`); the Forever env is live today. AtlasLoot stage needs replacing. | Same 3–4 wk. | Same 3–4 wk. No saving. |
| **Effort to first spec** | **2–3 wk** (11–15 wk end-to-end shippable) | 2–3 wk (13–18 wk end-to-end) | 2–3 wk (**25–37 wk end-to-end**) |
| **Effort to all specs** | 30–40 engineer-wk (DPS-only ≈ 18–25). Not solo-achievable; wowsims has 3/21 launched after 2 years with ~100 contributors. | Same 30–40. | Same 30–40. The spec work is identical in all three — it is dominated by unknown numbers, not by engine choice. |
| **UI** | 3 wk seam (ship their Vite SPA on a path) or 12–20 wk Svelte-native against the same protos. Per-spec UI is only 180–544 TS LOC; `ui/core` (36,459 LOC) is free. | Identical. | 12–20 wk, no seam, but nothing free. |
| **Browser vs server compute** | Both, free. Same Go builds to native, WASM, and an HTTP server with no branching (`sim/wasm/main.go`, `sim/web/main.go`, three interchangeable worker shims). **But measured WASM is 7.9× slower than native** — 10k iterations is ~16 s on a 4-worker laptop vs ~1 s on a few server cores (§1.10). Given UX is priority #1, plan for server compute with browser as an offline/fallback path, and copy Raidbots' Smart Sim ladder. Note MIT is what makes the browser path legal at all: **SimC's GPL-3.0 forbids it**, since serving a WASM build to a client is conveying (§2.7). | Identical. | You would have to build both paths yourself (~2–3 wk of the engine estimate), but you would own the licence, so both stay open. |
| **Risk** | **Low-moderate.** Fork divergence from upstream; the biome/creature-type mechanic is net-new; the unified-Hit × weapon-skill interaction is a genuine unknown that could force an attack-table rework. Fork divergence is *less* of a risk than it looks: upstream `classic` is dormant (1 commit in 90 days), so there is almost nothing to merge — but equally almost nothing to inherit. | **Moderate.** All of A's risks plus deletion risk — rune logic is entangled across `talents.go`, ability files, protos, and UI, so removal is where bugs hide. Also inheriting a dormant branch (2 commits in 90 days). | **High.** Re-deriving 20 years of community-verified vanilla mechanics, from scratch, against a 7-week launch deadline, while ~100 volunteers ship a working alternative. The stated benefit (shared event model) turns out to be a 1–2 week adapter available in every option. |

**Reading of the evidence:** Option A, cherry-picking `spell_mod.go` from B, plus the
`event.Event` emission adapter and the log-validation loop from C's good idea. Server-side compute with
Smart-Sim-style precision laddering, browser WASM as an offline/fallback path. Ship the wowsims UI on a
path for launch, replace it with Svelte natively once specs exist.
# Raidbots product surface (walked in the browser, Sept 14, 2026)

Source pages: raidbots.com/simbot, /simbot/topgear, /simbot/quick (with a prefilled Armory character), /developers, support.raidbots.com articles 28, 42, 43, 54, 59, 61, 63, 64, 65, 66, 69; Medium posts on queue, limits, premium.

## What it is
A web front end that builds a SimulationCraft input, queues it to 32-core worker VMs, stores SimC's output files in Google Cloud Storage, and renders a report page. No engine of its own. "Results are only as good as the SimulationCraft model for your spec" is on every form. Healers are explicitly unsupported; tanks get damage only.

## Tools (input → output)
| Tool | Input | Output |
|---|---|---|
| Quick Sim | SimC addon paste, or Armory (region/realm/name; no bag items), or History | DPS with error bar, damage breakdown per ability (WCL-style, expandable, uptimes, school colours), buff uptimes (perfect-uptime buffs separated), sample ability log for one iteration (time, ability, target, active buffs), talent tree render, SimC HTML report |
| Top Gear | Same, plus tick boxes for every bag item, enchant, gem, consumable and talent alternative; combination counter at the bottom against the tier's iteration budget | Ranked list of gear combinations with DPS and delta; "Your Top Gear" block; sidegrades grouped (within 2×error, ~0.1%); "show gear differences from" selector; copy /simc of any row; Smart Sim runs all combos at low precision, prunes, re-sims at higher precision, repeats (multistage) |
| Droptimizer | Character + a source (raid difficulty, all M+ dungeons, PvP, professions) + upgrade level | Per boss: Best Drop, Expected Value (mean gain, losses count as 0), upgrade probability, Priority rank (groups bosses within 0.2%, then by upgrade chance, then Best Drop); per item: DPS gain. One item at a time, both trinket/ring/weapon slots tried |
| Advanced | Raw SimC script (hardware/IO options blocked) | SimC output as is |
| Stat Weights (legacy) | Character | Scale factors for primary/secondary/weapon dps; site itself says "Beware of Stat Weights" and steers users to direct sims |
| Gear Compare (legacy) | Character + gear sets | Side-by-side DPS |
| Talent Compare | Not offered: "60 million combinations ≈ 300 billion iterations"; they plan targeted sub-tree tools instead |
| Expert Mode | Inject SimC text at 11 points (header, base actor, consumables, expansion options, custom APL, pre/post actor, actors, enemies, sim options, footer); no validation, no support | |

## Simulation options (Quick Sim form)
Fight style: Patchwerk, Dungeon Slice, Target Dummy, Execute Patchwerk, Hectic Add Cleave, Light Movement, Heavy Movement, Casting Patchwerk, Cleave Add (Dungeon Slice marked unreliable for many specs). Bosses 1–10. Fight length 20 s to 10 min, varied ±20% per iteration. SimC version weekly/nightly/latest. Trinket-specific toggles per expansion. High Precision checkbox (2× more precise, 4× slower). Default iterations=100000, max_time=300, raid buffs overridden on, scale_only list for stat weights.

## Report and data
- Report URL is public, shareable, expires (old sims removed). Files at `<report>/data.json` (top 200 actors), `input.txt`, `preview.png` (Discord embed), `data.full.json`, `data.csv`, `index.html` (SimC HTML), `output.txt`.
- Static data JSON generated from the client with SimC's casc/dbc tools: equippable-items, item-names, bonuses, talents, instances, enchantments, crafting, item curves, item sets, item-limit-categories; per environment (live/ptr/beta) and per build.
- Daily "top 100 actors per spec" summary.csv/details.json for SimC developers.
- Talent tree embed (server-rendered iframe with a Blizzard talent string).
- Armory prefill query params (`?region=&realm=&name=`) on any tool.

## Business and infrastructure
- Free tier: end of the queue, iteration cap (historically 600k Advanced), time limit. Premium tiers Rare $3 / Heirloom $5 / Epic $10 per month (2018 prices): skip the line in tier order, larger Top Gear budgets (1.5M / 4M / 12M iterations), all Droptimizer sources at once, Epic gets SwiftSim (two 32-core VMs, ~2× on big sims) and guild Discord-bot skip-the-line.
- Queue goal: public wait under 5 minutes; workers scale with queue depth.
- Discord bot runs sims from chat.
- SimC addon export is the primary input (character, gear with bonus ids, bag items, talents); Armory second.

## Takeaways for Forever Sixty
1. The product is a compute queue plus a report renderer over someone else's engine; the moat is the engine's per-spec fidelity, which they do not own and openly disclaim.
2. The features people actually use: Top Gear (bag items → best combination), Droptimizer (which boss to run), Quick Sim (what is my DPS and why). Stat weights are deprecated by the site itself.
3. Our addon export already exists in Phase 2 (FS1) and can carry bag items; our data pipeline already produces the static data they publish; our logs product gives us the validation signal SimC lacks for a new ruleset.
4. What they cannot do that we could: healer and tank models, browser-side compute (no queue, no premium gating), talent-tree search on our own trees, validation of the model against live rankings from our logs.
