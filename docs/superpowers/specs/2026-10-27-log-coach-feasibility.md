# Log coach feasibility (2026-10-27)

**Bottom line:** most of what a log coach needs is already parsed and stored. Two pieces are missing: what each player brought into the pull (two lookup tables are never passed to the parser), and talents and race in the log. Both have workarounds. One related bug should be fixed first.

## 1. What the parser stores per player per fight

The per-fight table is `summary.Summary` in `logs/engine/summary/summary.go`. It is stored as the companion's posted JSON in the bucket, and read back with `sims.fightSummary` (`api/internal/sims/summaries.go`). The rater gets the same posted summary (`api/internal/reports/ingest.go:404`, `rate`).

| Fact | Exists? | Where |
|---|---|---|
| Casts with timestamps | **Yes.** `CastRow` per caster and spell: started, succeeded, failed, fail reasons, total cast time, and `Sequence` (the ms offset of every successful cast). Pets are mapped to their owner. | `logs/engine/summary/deaths.go:116`, filled at `:440` |
| Buffs and debuffs with uptime | **Yes.** `AuraTrack` per (target, spell): uptime, applications, max stacks, segments with start, end and source, and appliers. Auras already up at the pull are seeded from COMBATANT_INFO (`seedAuras`). The track is keyed by target and spell, **not by caster**, so two warlocks' Corruption on one boss merge into one track. | `deaths.go:86`, `summary.go:396` |
| Resources | **Yes.** `ResourceTrack` per power type: per-second series, gained, spent, `ZeroMS` (time at zero), `Max`, `AtMaxMS` (time capped) and `Wasted` (gained past the cap). | `deaths.go:163` |
| Deaths | **Yes.** `Death`: time, killing blow, last 10 hits, last 40 heals, auras held and lost, release time. | `deaths.go:57` |
| Activity / dead time | **Yes.** `RosterRow.ActiveMS` and `ActivityPct` (1.5 s gap rule). `rating.ActivityIdleMS` subtracts the curated downtime windows. | `roster.go:27`, `logs/engine/rating/activity.go:47` |
| Damage by ability | **Yes.** `Ability`: hits, crits, ticks, misses, min, max, plus a per-second series per actor. | `damage.go:13` |
| Interrupts and dispels | **Yes.** `ExchangeRow` | `deaths.go:143` |
| Mechanics | **Yes, where curated:** avoidable hits, interrupt duty, debuff uptime, phases. | `summary/mechanics.go`, `logs/engine/mechanics` |
| Gear snapshot | **Yes, boss pulls only.** `CombatantRow.Gear` (item id, item level, enchants, gems), `ItemLevel` and the stat block. Forever writes log format v22 (`layout/retail.go:113`). It is not on trash, and not in the Parquet rebuild (`parquet/schema.go:215`), so the posted JSON is the only copy. | `roster.go:13` |
| Consumables and raid buffs at the pull | **Columns exist but are always empty in production.** `CombatantRow.Consumables`, `RaidBuffs` and `MissingBuffs` are filled only when `Options.ConsumableSpells` and `RaidBuffSpells` are passed in. Neither the API (`api/internal/engine/engine.go:89`) nor the companion (`companion/internal/pipeline/pipeline.go:296`) passes them. The same facts can be recovered from the seeded `Auras` tracks. | `roster.go:112-141` |
| Talents | **No.** v22 writes talent triples and the decoder deliberately leaves `Talents` empty (`layout/retail.go:250`). | — |
| Race | **No.** No log format records it (`api/internal/sims/score.go:71`). | — |
| Level | Only in the advanced-logging fields (`AdvLevel`). | `parquet/schema.go:95` |

**Bug to fix first (HIGH).** These read `CombatantRow.Consumables`, which nothing fills:
- `rating.consumablesMissed` (`api/internal/rating/last_raid.go:179`)
- `fightDiscipline` / `carries` (`api/internal/rating/score_v2.go:567-590`)
- the engine's `scorePreparation` (`logs/engine/rating/preparation.go:26-39`)

So any player with a gear snapshot is reported as missing flask and food, Preparation gets almost no flask, food or world-buff credit, and Discipline is false. I couldn't confirm this against production data from here. The fix is either to pass the catalogue into `engine.SummaryOptions` and the companion, or to read the seeded `Auras` tracks (segment `StartMS == 0` on the player). The second needs no companion release.

## 2. What the sim engine can answer

- **The rotation per spec: yes.** Every spec has a written APL with ranks and a plain note per line: `data/curated/apl/<spec>.json`, copied to `sim/request/apl/*.apl.json`. The addon's level-banded rotation is loaded by `sims.RotationFor` (`api/internal/sims/rotation.go:69`), and `Rotation.Family` maps any rank to its ability.
- **Cooldowns:** `CooldownSpec` in `sim/api/envelope.go:212` and `sim/request/cooldowns.go`. Each spell's cooldown, category cooldown, GCD and cast time are in `data/builds/<build>/spellconst/<class>.json`. Max rank by level is in `trainables/` and `spellranks.json`.
- **Expected casts: yes, not just DPS.** `SimResult.Summary` is the same `summary.Summary` type, built by `sim/adapter`. Cast counts and aura uptimes are averaged per iteration (`adapter.go:172`, `:546`). `Sequence` is empty there; one iteration's timeline is in `SimResult.Sample` (`SampleCast`: time, action, resources). So "expected casts per minute and uptime for this spec" is directly comparable with a log summary.
- **"The sim with this player's gear and talents": built but always skipped.** `sims.Score` and `CombatantBuilder.FightCharacter` (`api/internal/sims/score.go:91`) always fail with "no race" (and on Forever also "no talents"), so `execution_score` never fills in. The addon export does carry race and talent trees (`api/internal/fs1/fs1.go:43-53`, in `addon_exports.export`). A builder that takes gear from the log and race and talents from the export would unblock it for any player whose addon export is synced. That is a natural fit for Premium. Cost is about 3 s of native compute per fight at `ScoreIterations`; run it on demand, not at ingest. DPS specs only (`score.go:243`).

## 3. Reuse candidates

- **Rating card**, `logs/engine/rating`: six components (output, survival, mechanics, utility, preparation, activity), each with `Moments` (death, avoidable hit, interrupt, dropped debuff, with anchors into the report page), stored in `rating_scores` (migration 0022).
- **Within-bracket percentiles:** `PercentileSource` and `rating_percentile_digests` keyed (encounter, difficulty, spec, role, kill-time band, kill, component). `MinSample = 20`, falling back to an absolute standard below that.
- **Last-raid facts** (`api/internal/rating/last_raid.go`): biggest loss as a plain phrase (downtime, dead early, died to X, standing in fire, missing consumables), loss seconds, consumables missed. This is already coach-shaped copy.
- **Rotation adherence** (`api/internal/rating/adherence.go`, `adherence_exclusions.go`): the share of casts that are on the shipped rotation.
- **Readiness** (`api/internal/guilds/readiness_core.go`): enchant gaps on weapon, chest, cloak and boots, unspent talent points (from the addon export), consumables in bags, gear gap against BiS. Pure function. Reads the addon export, not the log.
- **Discipline enchant check** (`score_v2.go:594`), which reads the log gear.
- **Parse percentile / rankings:** `fight_metrics` and `percentile_digests` (migration 0005), and `api/internal/rankings/query.go`.
- **Premium gating:** `entitlements.Store.Can(ctx, userID, Feature)`. Add `FeatureLogCoach Feature = "log_coach"`; it counts as a personal Premium feature (`Standard()`) automatically.

## 4. The comparison pool

A "top logs for this spec" comparison needs per-ability numbers for the top players. Today those live only in per-fight `summary.json` objects in the bucket; `fight_metrics` holds only headline DPS and HPS. Reading N summaries per request is too slow.

**Proposed table:** `coach_profiles`, partitioned by `fought_at` like `fight_metrics`.
- Columns: `report_id, fight_index, player_key, spec, role, encounter_id, difficulty, size, kill, duration_ms, metric, fought_at, model_version, profile jsonb`.
- `profile` holds casts per minute per ability family, uptime % per key aura, active %, % of time capped, OOM ms, and pull consumables.
- Written by the rater beside `rating_scores`, public and guild reports only (the same rule as `reports.Ranked`).
- Index: `(encounter_id, difficulty, spec, kill, metric desc)` plus `(player_key, fought_at desc)`.
- Comparison set: the top 10–25 kills in the bracket, compared by median.

**Fallback when the pool is thin** (below 20 kills, the existing `MinSample` rule), first match wins:
1. The same spec across every encounter of the same size, normalised per minute.
2. **A sim reference profile:** the spec's APL at the player's level band, preset gear for that band, cached per (build, spec, band, rough duration). It always exists, so the coach is never empty on launch day.
3. Absolute ceilings from spell data, e.g. `floor(duration / cooldown) + 1` casts possible.

Every finding states its basis ("vs the top 20 Ragnaros kills" / "vs the simulated rotation").

## 5. Raid sizes (10 / 20 / 40)

- **Parser:** nothing assumes 40. `fight.Fight.Size` comes from ENCOUNTER_START and is stored on `fights.size`, `fight_metrics.size` and `rating_scores.size`. Difficulty ids are kept as logged; the web labels 3 = 10, 148 = 20, 9/186 = 40 (`web/src/lib/difficulty.ts`).
- **Rating brackets** key on (encounter, difficulty) without size. That is fine as long as each size has its own difficulty id (true for these ids). The coach should still put `size` in its own key explicitly.
- **Readiness** is level-based, not size-based. Only `api/cmd/seedguild` hard-codes 186.
- **Risk:** `MissingBuffs` (`roster.go:125`) flags every listed raid buff that is absent, whether or not the raid had a class to provide it. In a 10-man without a priest that is wrong. The coach must only flag a buff when its provider class was in the roster.
- World-buff and consumable expectations should be relaxed for 10-player content. The catalogue weights are per role, not per size.

## 6. A v1 that can ship fast

**New package:** `api/internal/coach`.
- A pure `coach.Review(fight summary.Summary, row summary.RosterRow, ref Reference, tables Tables) []Finding`. Tables are the APL and rotation, spellconst, spellranks, trainables and the consumables catalogue.
- `Finding { Code, Severity, Title, Detail, Basis, Evidence []Moment, GainEstimate *float64 }`.
- Copy is plain WoW words from a single copy table in the package.
- Reference resolver: `coach_profiles`, then the sim reference.

**Findings per spec family, all computable from today's data:**

1. **Core ability under-used / clipped.** For each cooldown-gated APL ability (Bloodthirst, Mortal Strike, Shadow Burn, Judgement, Starsurge-type, Aimed Shot, Death Wish, Recklessness, Arcane Power, Innervate, …), compare casts against the possible count from `spellconst` cooldown and fight length, and list the gaps longer than cooldown + 1.5 s from `CastRow.Sequence`. Example: "Bloodthirst: 31 of a possible 38 — 7 casts lost, longest gap 0:42–0:55."
2. **Key buff/debuff uptime.** The player's own DoTs and debuffs (Corruption, Sunder, Serpent Sting, Faerie Fire, Judgement debuffs) on the boss track, attributed through `Appliers`, plus self-buffs (Battle Shout, Slice and Dice, Inner Fire). Shown against the reference uptime. Shared debuffs are flagged as "raid uptime".
3. **Dead time.** `ActivityIdleMS` minus curated downtime, plus gaps between casts longer than GCD + 1 s while alive, with the three longest windows timestamped.
4. **Resource waste.** `AtMaxMS` / `Wasted` for rage and energy ("energy full for 26 s"), `ZeroMS` / OOM time for mana users, and when OOM first happened relative to the fight's end.
5. **Wrong-rank spells.** Casts whose spell id is a lower rank (via `Rotation.Family` and `spellranks`) than the highest trainable rank at the player's level, excluding deliberate downranks the APL lists.
6. **Pull preparation.** Flask, food, weapon oil or stone, and world buffs at the pull from the seeded `Auras` tracks; combat potion casts from `Casts` against the catalogue's `MaxUses`; missing permanent enchants from log gear (`score_v2.enchanted`); raid buffs only when the provider was present.
7. **Death review.** First death: time, killing blow, whether it was avoidable (mechanics table), and the defensive cooldowns or potions available and unused (spellconst and casts). Reuses the rating's `Moments`.
8. **Rotation adherence + off-rotation casts.** `fightAdherence` plus the top three off-rotation spells by cast count.

Healer and tank variants of 1–3 use the same machinery, with overheal % from `Actor.Overheal` and defensive uptime.

**Out of v1:**
- Talent points unspent: comes only from the addon export, via the readiness route.
- "Sim of your own character": v2, needs the export-based builder.

**API:**
- `GET /v1/reports/{id}/fights/{n}/coach?player=<key>` returns `{ player, spec, size, basis: { kind: "top_logs"|"sim"|"absolute", n }, findings: [...] }`.
- `GET /v1/me/coach/recent` returns the newest fights of the signed-in user's characters.
- Computed on request from the bucket summary and cached per (report, fight, player, model_version).

**Premium gate:**
- **Free:** the finding count and the top finding's title. The existing last-raid "biggest loss" already gives a free teaser.
- **Premium** (`FeatureLogCoach`): every finding with details, timestamps, the top-logs comparison and coach history.
- Viewing is limited to players allowed to see the report (the existing visibility rules). Coaching someone else's character stays guild-officer-only (`FeatureOfficerViews`).

## 7. Risks and what isn't possible yet

- **Consumables tables never passed to the parser:** the rating bug from §1. Fix it before the coach repeats it.
- **No talents or race in Forever logs:** no talent findings and no personal sim without a synced addon export.
- **Gear only on boss pulls, and only in the posted JSON:** trash fights get no gear or consumable findings. Parquet re-parses can't recover it (`engine.go:150` follow-up).
- **Aura tracks merge across casters:** uptime of a debuff shared between two players of the same class can't be split per player.
- **Positioning, facing and movement** fields exist in advanced logging but nothing in the summary uses them. "Stood in fire" depends on curated mechanics tables, which cover few encounters so far.
- **Weapon oils and stones:** unverified how they show up in COMBATANT_INFO (`score_v2.go:546`).
- **Classic layout without advanced logging:** no character sheet and fewer facts. The coach must say what it couldn't check.
- **Pool size in production is unknown.** Early on most players will be compared with the sim reference, not top logs. Label the basis honestly.
- **Sim reference fidelity:** APLs are curated and some talents are unmodelled. The reference is "the written rotation", not a ceiling.
- **Pet and guardian casts** are mapped to the owner but kept as separate rows. Hunter and warlock findings need to fold them.

---
