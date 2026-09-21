# Performance rating engine and ratings on the site — design

**Date:** 2026-09-21
**Status:** DRAFT for review. Implements `docs/superpowers/specs/2026-09-21-guild-and-premium-proposal.md`
section 3.4 ("The performance analyzer") and honours 3.4's fairness rules and the
2026-09-21 decisions recorded there: one number first then the six parts then the log
moments; ratings are public the way parses are, honouring report visibility and
`anonymize`; single lookups free; no public leaderboard of ratings; the player's own
report card free.
**Scope:** the scoring engine, its storage, and the two FREE site surfaces — the per-fight
report card inside a report, and the rating on the character page. Officer views, roster
checks, payments, assignments UI and in-game display are **out of scope** (recorded in
§9), but the data model carries a defined hook so they can be built without a schema
change (§3.1).
**House style:** follows `docs/superpowers/specs/2026-09-21-one-product-design.md` (lane
ownership, honest copy, no layout shift, signed-out-first, design system discipline) and
`docs/superpowers/specs/2026-09-19-simulator-parity-interfaces.md` (a binding interfaces
document the lanes code against; amendments are numbered and additive, never silent).

---

## 0. What already exists, in one paragraph, cited

A fight's `summary.Summary` (`logs/engine/summary/summary.go:95-122`) already carries
everything the six components read: `DamageDone`/`DamageTaken`/`Healing`/`HealingTaken`
tables with per-second `Series` and `ActiveMS` (`logs/engine/summary/damage.go:46-58`),
`Deaths` with the last ten damage events, healing events, and auras held/lost around each
death (`logs/engine/summary/deaths.go:57-70`), `Auras` with per-applier uptime segments
(`logs/engine/summary/deaths.go:86-99`), `Casts`/`Interrupts`/`Dispels` with source and
target GUIDs (`logs/engine/summary/deaths.go:116-160`), `Resources`, `Threat`/`ThreatByTarget`
(`logs/engine/summary/threat.go:77-98`, model incomplete — see §2.6), `Taunts`
(`logs/engine/summary/taunts.go:9-22`), `Combatants` (gear/talents/consumables/raid buffs
at pull, `logs/engine/summary/roster.go:13-24`), `Roster` (role, DPS/HPS/DTPS,
`ActivityPct`, `logs/engine/summary/roster.go:27-45`), a curated `Mechanics` block
(`logs/engine/summary/mechanics.go:16-45`) and `Phases`
(`logs/engine/summary/phases.go:13-17`). The curated tables live in
`logs/engine/mechanics/tables/*.json`, embedded and validated at build time
(`logs/engine/mechanics/mechanics.go:102-135`), one avoidable/unavoidable/interrupt/dispel
classification per spell id, with an optional `role` on an avoidable entry
(`logs/engine/mechanics/mechanics.go:44-47`) and an optional phase table
(`logs/engine/mechanics/mechanics.go:73-90`). The execution score already exists:
`api/internal/sims/score.go:227-300` runs the fight's own recorded gear through the
simulator and writes `row.DPS / res.DPS.Mean` into `fight_metrics.execution_score`
(migration `api/internal/db/migrations/0013_sims.up.sql:9-11`) — DPS role only
(`score.go:243-247`, "tanks and healers are research problems"), spec-validated only
(`score.go:254-258`), and only when the character is a signed-in, claimed member
(`api/internal/reports/ingest.go:395-421`, `MemberKeys` at `api/internal/auth/store.go:369`).
The percentile machinery already exists as a merging t-digest
(`api/internal/digest/digest.go:1-53`) folded per `(encounter_id, difficulty, spec, phase,
metric)` bracket at fight close (`api/internal/rankings/store.go:129-224`) and read back by
`Store.Percentile` (`api/internal/rankings/store.go:329-352`). Report visibility is a single
function, `Service.mayView` (`api/internal/reports/handler.go:521-538`): public and unlisted
readable by anyone with the link, private by owner/moderator, guild by
`Accounts.GuildRank`. **`anonymize` today governs exactly one thing** —
`auth.User.PublicName()` (`api/internal/auth/store.go:67-79`), which masks the *report
owner's* battletag — and nothing filters `fight_metrics`, rankings or a character page by
it (§5.4 records this gap and the ruling that follows from it).

---

## 1. The scoring model

### 1.1 Components, inputs, and what "no data" means

Every component is scored 0–100. A component the engine cannot score for a given fight
(no table, no digest, spec unmodeled, threat model incomplete, etc.) is **excluded**, not
zeroed: its weight is redistributed proportionally across the remaining scored components,
and the card states which parts are missing and why (§7.4 has the exact copy). This is the
single rule every formula below obeys; it is not restated per component.

| Component | Primary input (struct/field) | Fallback input | "No data" trigger |
|---|---|---|---|
| **Output** | `fight_metrics.execution_score` (`sims/score.go:298-299`) | `RosterRow.DPS`/`HPS` (`roster.go:42-43`) | Spec unvalidated (`sims.Store.Validated`), character unclaimed (`MemberKeys`), or bracket has no digest and no absolute standard |
| **Survival** | `Death.AtMS`, `Death.KillingBlow.SpellID` (`deaths.go:57-70`), `MechanicsBlock.Rows[kind=avoidable].Players` (`mechanics.go:22-45`) | — | Fight's encounter has no mechanics table (`MechanicsBlock.TableFound == false`) |
| **Mechanics** | `Interrupts`/`Dispels` (`ExchangeRow`, `deaths.go:143-160`), `MechanicsBlock.Rows[kind=interrupt\|dispel]` | — | No mechanics table, or table has no interrupt/dispel rows |
| **Utility** | `Auras` (`AuraTrack.UptimeMS`/`Appliers`, `deaths.go:86-99`) against new per-spec utility table (§3.2) | — | No utility table entry for this spec |
| **Preparation** | `CombatantRow.Consumables`/`RaidBuffs` (`roster.go:13-24`), `Casts` filtered to consumable ids | — | No consumable catalogue for this role (§3.3) |
| **Activity** | `RosterRow.ActiveMS`/`ActivityPct` (`roster.go:36-37`), `Death.AtMS`, new per-boss downtime windows (§3.4) | Whole-fight `ActivityPct` with downtime unexcluded | No downtime windows curated for this encounter (falls back to whole-fight `ActivityPct`, flagged approximate, never excluded — activity is always computable) |

### 1.2 Percentile-within-bracket vs. absolute standard

**The bracket.** `(encounter_id, difficulty, spec, role, kill_time_band, component)`.
`encounter_id`/`difficulty`/`spec` mirror `percentile_digests`
(`api/internal/db/migrations/0005_logs.up.sql`'s `percentile_digests` table); `role` is
added because a component's meaning differs by role (Survival's avoidable-damage share for
a tank is not comparable to a healer's); `kill_time_band` is new (§5.2).

`kill_time_band` is one of `fast` / `typical` / `slow`, computed against the encounter's
own trailing distribution rather than a curated per-boss duration (curating expected kill
times for every boss is exactly the maintenance burden the existing phase digest avoids by
keying on data, not on a hand-written table):

```
band(duration_ms, median_ms) =
  fast    if duration_ms <  0.85 * median_ms
  typical if 0.85 * median_ms <= duration_ms <= 1.15 * median_ms
  slow    if duration_ms >  1.15 * median_ms
```

`median_ms` is the bracket's own kill-duration digest (a seventh, tiny t-digest per
`(encounter_id, difficulty)` keyed on `duration_ms` of kills only, reusing
`api/internal/digest`), read at write time. **RULING: three fixed bands, not a continuous
percentile-of-duration.** A continuous band would need re-bucketing every existing rating
whenever the median shifts; three bands are stable once assigned to a fight (the median can
drift without moving already-scored fights out of their band) and match how raid leaders
already talk about a pull ("that was a fast kill"). Cost if wrong: a boss whose kill times
cluster tightly (most Molten Core bosses, once execute-checked) will put almost everything
in `typical`, which is harmless — the band degenerates to a no-op, not a wrong answer.

**Percentile vs. absolute, the exact rule:**

```
if bracket.n >= MIN_SAMPLE (20 kills) and an execution-score-shaped input exists:
    use percentile-within-bracket (the component's raw value placed in its own digest,
    via digest.Digest.Placement, api/internal/digest/digest.go:212)
elif an absolute standard exists for this component (§1.3 per component):
    use the absolute standard
else:
    component excluded; weight renormalised; card says why
```

**RULING: `MIN_SAMPLE = 20`.** `DefaultCompression = 100`
(`api/internal/digest/digest.go:19-22`, "a digest is a few hundred centroids ... the
median is accurate to well under a percent") is accurate at any *n*, but a percentile
computed from fewer than ~20 points is a percentile of noise dressed as a number, and the
existing rankings page has the same problem today for a brand-new encounter — this spec
does not fix that, it just refuses to let ratings inherit it silently. Twenty is a round
number chosen for defensibility, not derived from the digest's own error bound; if December
raid nights produce ten pulls of a boss on the first night, every rating on it runs on the
absolute standard until the eleventh kill crosses the threshold within its band, which is
the honest state of the world. Cost if wrong: too high and ratings sit on absolute standards
longer than necessary (more conservative, not wrong); too low and early percentiles are
noisy (the confidence band in §1.5 covers this).

The card states which rule was used, in these exact words: *"Compared with other `<spec>`
`<role>`s on this fight (`<n>` logs)"* for percentile, *"Measured against the encounter's own
numbers — not enough logs yet to compare players"* for absolute, and the excluded-component
copy from §7.4 otherwise.

### 1.3 Formulas, per component

**Output.**

*Value*: `execution_score` when the fight has one (DPS role, spec validated, character
claimed — `sims/score.go:227-300`); else raw `metric_dps`/`metric_hps`
(`fight_metrics.metric_dps`/`metric_hps`, written at `rankings/store.go:172-181`) as the
fallback value, exactly as the proposal states ("falls back to the spec percentile where
the simulator does not model the spec yet", `2026-09-21-guild-and-premium-proposal.md:102`).

*Absolute standard* (execution_score path only, since it is already normalised 0→≈1.3):
`score = min(100, execution_score * 100)`. **RULING: capped at 100, not scaled past it.**
A pull with favourable crit variance can post `execution_score > 1.0`; scoring it above 100
would make "how good was this player" indistinguishable from "how lucky was this pull," and
the percentile path (which *does* reward beating peers on the same lucky night) is exactly
where that variance belongs. Cost if wrong: a genuinely exceptional played fight (perfect
rotation, no wasted GCDs) on top of good variance reads identically to a merely-good fight
with great variance, until the bracket is large enough to run the percentile path instead.

*Raw-DPS fallback has no absolute standard* (DPS is unbounded and gear-dependent): with
`bracket.n < 20`, Output is excluded rather than scored off an ungrounded number.

*Tanks*: value is the tank's own `metric_dps` (tanks still deal damage); weight is small
(§1.4). No execution-score path exists for tanks today (`score.go:243-247`); this is a
known gap, not modeled around — see §9's testing note for when the simulator's tank
coverage lands.

*Healers*: value is `metric_hps`, which is **already effective healing, not raw** —
`RosterRow.HealingDone = hd[guid].Effective` (`roster.go:173`), and `Actor.Effective`
excludes overheal by construction (`damage.go:46-58`, `Overheal` is tracked as a separate
field). No engine change needed here; the proposal's "healers scored on effective healing"
requirement (§2 of this doc) is already true of the underlying number.

**Survival.** Two parts, combined `0.65 * DeathScore + 0.35 * AvoidableHitScore`.
**RULING: this 65/35 split, not a single blended number.** Deaths are discrete and
catastrophic; sub-lethal avoidable damage is continuous and forgiving; collapsing them into
one running total would let ten small avoidable hits look as bad as a death, which they are
not — a death removes the player from the fight, ten hits do not. Cost if wrong: the split
is a judgment call with no data to validate it before real raid logs exist; it is exposed on
the weights page (§7.3) and adjustable the same way every other weight is, so a wrong split
is a config change, not a redeploy.

*DeathScore*: `100 - Σ penalty_i` (floored at 0) over the fight's deaths for this player
(`Deaths` filtered by `GUID`), where

```
penalty_i = 70 * causeMultiplier_i * timeRemaining_i
timeRemaining_i = 1 - (Death.AtMS / DurationMS)      // fraction of the fight left when they died
causeMultiplier_i =
    1.00  if KillingBlow.SpellID classifies "avoidable" in this encounter's mechanics.Table
    0.15  if it classifies "unavoidable"
    0.50  if unclassified (no table, or the spell id is not listed) — "an unclassified
          killing blow is neutral," per the brief
```

**RULING: "when" reads as time-remaining, not boss-health-remaining.** The proposal's own
words — "a death at 95% costs more than one at 3%" — read most directly as *95% of the
fight still to go* costing more than *3% left*, and `Death.AtMS`/`Summary.DurationMS` are
always present (`deaths.go:61`, `summary.go:98`) where a boss's own health reading depends
on advanced combat logging being on and the table naming the right unit
(`phases.go:64-70` shows how fragile that lookup already is). Boss-health timing was
considered and rejected: it is unavailable on any fight without advanced logging, and
"time left" already captures the intuition (a death that removes you from most of the
remaining fight costs more) without a second data dependency. Cost if wrong: on a fight with
a long, low-intensity DPS-race final phase, a death at 90% time-remaining is scored as
costly even if the boss was nearly dead at that instant (Golemagg-style burn fights);
inspect against a real raid log before Dec 9 and move to a boss-health reading (already
partially available per `phases.go`) if this reads wrong in practice.

*Killing blow classification, exact lookup*: `mechanics.Table.Lookup(KillingBlow.SpellID)`
(`mechanics.go:202-209`); no table (`MechanicsBlock.TableFound == false`) makes Survival's
death half unclassifiable and Survival is excluded entirely for that fight (not just the
cause term), since a death-timing-only signal with no cause information is not the
"avoidable damage taken" component the brief asks for.

*AvoidableHitScore*: `100 - percentile(avoidableDamageTakenPerSecond)` within the same
bracket, where `avoidableDamageTakenPerSecond = Σ MechanicHit.Damage / (DurationMS/1000)`
summed over `MechanicsBlock.Rows[kind == avoidable].Players[GUID == this player]`
(`mechanics.go:22-45`, `MechanicHit` at `mechanics.go:49-61`), with a `Role`-tagged row (e.g.
the tank in `tables/1084.json`'s `role: "tank"` rows) going through the assignment check
below rather than being excluded unconditionally. A *non*-tank hit by a tank-role mechanic
still counts fully against them regardless, exactly matching `Mechanic.Role`'s existing doc
comment: "a hit on anyone else is then that role's problem as much as the victim's"
(`mechanics.go:44-47`).

**CORRECTION (whole-branch review):** this section originally read "excluding any row whose
`Role` equals this player's own role" unconditionally — meaning a `role: "tank"` cleave that
hit the *off* tank standing in it scored identically to the *main* tank taking the exact same
mechanic as their actual job, because the curated table alone cannot say which player of a
role a given hit was "for." **RULING:** a `Role`-tagged row excuses a same-role player only
when (a) the fight carries no `Assignment` (§2) overlapping that hit's own window for any
player of that role, in which case every player of that role is excused — the table cannot
tell the main tank doing their job from the off tank standing in it, and wrongly punishing
the main tank for every cleave is the worse of the two errors — or (b) an `Assignment`
overlapping that hit's window names this specific player. When one or more assignments for
that role overlap the hit's window and none of them names this player, the hit counts in
full: it was demonstrably someone else's to take. This is the literal mechanism behind
"tanks' ... Survival judged on avoidable damage only, since taking damage is the job": with
no assignment data (today's default, since officer tooling does not exist yet, §2), every
`role: "tank"` hit is still excused for every tank exactly as before; the assignment path only
ever makes scoring *stricter*, once a guild actually marks who had which add or cleave.

**Mechanics.** Three parts, averaged with equal weight (⅓ each) among the parts that have
data (each is independently excludable):

1. **Interrupts made, value-weighted.** For each `ExchangeRow{Kind: "interrupt",
   SourceGUID: this player}` (`deaths.go:143-160`), credit
   `MechanicRow.Damage / max(1, MechanicRow.Casts)` (the average cost of one uninterrupted
   cast of that spell, from `mechanics.go:24-44`) as "damage prevented." Sum, then
   percentile-within-bracket (bracket value: prevented-damage per second of the player's
   time alive).
2. **Dispels made, value-weighted.** Same shape, over `ExchangeRow{Kind: "dispel"}` and
   `MechanicRow.Healed` (a dispel's value is the enemy healing it denied, per
   `mechanics.go:39-44`'s existing "healing it gave the enemies" field) plus
   `MechanicRow.Damage` for a dispel that was itself doing damage (e.g. a DoT debuff).
3. **Own dispellable-debuff uptime**, lower-is-better. For every `AuraTrack` where
   `Type == "DEBUFF"`, `TargetGUID == this player`, and `SpellID` classifies `dispel` in the
   encounter's table, sum `UptimeMS`; percentile-within-bracket, inverted
   (`100 - percentile`). This is the "debuffs carried that should have been removed" half of
   Mechanics, and it scores the *target*, not the healer who should have dispelled them —
   see §2's fairness rule on assigned jobs for why a raid-wide "who should have dispelled"
   judgment is explicitly not attempted here.

**Utility.** Percentile-within-bracket of mean uptime-share across every buff/debuff/
cooldown the new per-spec utility table (§3.2) lists as "owned" by this spec, each uptime
measured as `AuraTrack.UptimeMS / DurationMS` for the aura the player is the source of
(`AuraTrack.Appliers` contains the player's GUID, `deaths.go:98`). **RULING: mean across
owned utilities, not minimum.** A single dropped buff (say, a Warrior who let Demoralizing
Shout lapse once) under a minimum rule would floor the whole component even if every other
owned utility ran at 100%; a mean rewards broad coverage the way a raid leader actually
reads it — "mostly on top of their jobs" rather than "failed the worst one." Cost if wrong:
a player who reliably drops one specific buff every fight (a genuine, correctable habit)
is under-penalized relative to a "worst-of" rule; the per-part breakdown (§7.1) still shows
the specific dropped buff, so the coaching value survives even if the number is generous.
Threat ("kept under the tank") is **not scored** until `BaseThreat.Complete()` is true
(`threat.go:41-54` — the modifier table is empty pending Forever's own numbers); the card
lists "Threat" among Utility's parts with a "not modeled yet" note rather than silently
omitting it, so a guild does not read its absence as "this player never pulls."

**Preparation.** Percentile-within-bracket of a completeness score built from the new
consumable catalogue (§3.3): `Σ weight_j` over every catalogue entry present, where entries
are matched against `CombatantRow.Consumables` (pull-time flask/elixir/food/weapon
oil-or-stone/world-buff snapshot, `roster.go:13-24`, `Options.ConsumableSpells`/
`RaidBuffSpells`, `summary.go:39-44`) for persistent buffs, and `Casts` filtered to
potion-classed consumable ids (`deaths.go:116-133`) for in-combat potion use, capped at the
catalogue's declared max-uses-per-fight (shared cooldown). `weight_j` sums to 100 per the
catalogue's own per-role weighting (§3.3's schema).

**Activity.** `100 * percentile(activeShare)`, where

```
activeShare = ActiveMS_excluding_downtime / (timeAliveMS - downtimeOverlapMS)
timeAliveMS = min(DurationMS, firstDeathAtMS or DurationMS if no death)
downtimeOverlapMS = Σ overlap(downtime window, [0, timeAliveMS])   // curated windows, §3.4
```

`ActiveMS_excluding_downtime` requires a small engine change: today `Actor.ActiveMS`
(`damage.go:79-105`, `markActive`) is a whole-fight figure with no concept of a downtime
window to subtract; the rating package (not the summary accumulator — see §5.1's "why a new
package, not a change to `summary`") re-derives a downtime-adjusted active time from the
same `Casts`/damage timestamps `markActive` already folds, filtered to exclude any interval
inside a curated downtime window. A fight with no curated downtime windows falls back to
the existing whole-fight `RosterRow.ActivityPct` (`roster.go:37`), flagged `approximate:
true` on the card rather than excluded — Activity is the one component that is always
computable, just sometimes less precisely.

### 1.4 Weights, per role, default

| Component | DPS | Healer | Tank |
|---|---:|---:|---:|
| Output | 35 | 30 | 10 |
| Survival | 15 | 10 | 35 |
| Mechanics | 20 | 20 | 20 |
| Utility | 15 | 20 | 20 |
| Preparation | 10 | 10 | 10 |
| Activity | 5 | 10 | 5 |
| **Total** | 100 | 100 | 100 |

These are the site-wide defaults; a guild's officer tools may change them per §3.4 of the
proposal ("no hidden weights ... a guild can change them") — that control is officer-tooling
and out of scope here, but the weights table itself lives in the curated-data layer (§3.2's
storage note) so the officer feature, when built, edits data, not code.

### 1.5 Combining into the overall score

```
overall = Σ (weight_c / Σ weight_scored) * score_c    over every scored (non-excluded) component
```

a plain weighted mean over the renormalised weights, **except**:

**RULING: a hard cap when Survival is catastrophic.** If `DeathScore == 0` for this fight
(the player's deaths alone already floor the death half of Survival — i.e. two or more
avoidable deaths with meaningful time remaining, or one very early avoidable death), the
overall score is capped at `min(overall, 40)` regardless of how the other five components
scored. **Reason:** a player who died avoidably early and then, playing a corpse, cannot
generate poor Output/Activity numbers that the weighted mean would otherwise read as
"balanced" — the weighted mean of a component that is structurally impossible to fail
(you cannot mis-time a heal while dead) against ones that are catastrophically failed
produces a misleadingly middling score. The cap is the mechanism the brief's "public the
way parses are, honest and coaching" tone needs: a 55 that hides "died in the first minute"
is a worse UI failure than a 40 that says why. **Cost if wrong:** a player who dies early
to genuine bad luck (a mechanic nobody could have told them about, on an unclassified
mechanic — but note `causeMultiplier` already discounts unclassified deaths to 0.50, so
`DeathScore == 0` mainly triggers on classified-avoidable deaths) is capped alongside a
player who died to obvious standing-in-fire; the per-part breakdown makes the distinction
visible one click away, which is the same mitigation §1.3 already leans on elsewhere.

**Coordinator ruling (accepted): the cap stays, and is a guild-adjustable setting,
published in plain words.** Alongside the weights table (§1.4, itself already guild-
adjustable officer tooling, out of scope), a guild may turn the 40-point cap off entirely.
Because that is a *display* choice over an already-computed fight, not a re-judgment of
what happened in it, `rating.Score` (§4.1) always computes and stores **both** figures —
`OverallUncapped` (the plain weighted mean, this section's formula with no exception) and
`Overall` (`OverallUncapped` capped at 40 when `OverallCapped` fired) — so that a guild
turning the cap off, and any future recompute of the cap's own threshold, read the
uncapped figure straight back with **no backfill**: the API serves whichever of the two a
viewer's settings call for (§5.2's schema carries both). The explanation page (§6.3)
publishes the rule in exactly these words: *"An avoidable death early in a fight caps the
overall score at 40, because nothing else in the fight makes up for it."*

`overall` and every `score_c` round to the nearest integer for display; the stored value
keeps two decimal places (§5.2's schema) so a recompute is stable to the same input.

### 1.6 Worked examples

All three use a `typical` kill-time band, `bracket.n >= 20` (percentile path), Molten Core's
Shazzrah (`encounter_id: 667`, `logs/engine/mechanics/tables/667.json`), a 3-minute
(180,000 ms) kill.

**DPS — Warrior Fury, one avoidable death at 140,000 ms (78% through), classified
avoidable (Arcane Explosion, `spell_id: 19712`):**

| Component | Raw / percentile | Score |
|---|---|---:|
| Output | `execution_score = 0.97` → 71st percentile of Fury on this bracket | 71 |
| Survival | 1 death: `penalty = 70 * 1.00 * (1 - 140000/180000) = 70 * 0.222 = 15.6` → DeathScore 84.4; AvoidableHitScore 60th pct → 60 | `0.65*84.4 + 0.35*60 = 75.9` |
| Mechanics | 2 interrupts made (raid needed 3), no dispels owned by Warrior on this fight, no dispellable debuff carried | 55th pct → 55 |
| Utility | Demo Shout 88% uptime, Sunder stacks 92% uptime on primary target, mean 90% → | 62nd pct → 62 |
| Preparation | Flask + food + weapon oil present, one Rage/combat potion used | 95 |
| Activity | 91% active share, no curated downtime for this encounter yet (approximate) | 80th pct → 80 |

`overall (unrounded) = .35*71 + .15*75.9 + .20*55 + .15*62 + .10*95 + .05*80`
`= 24.85 + 11.39 + 11.0 + 9.3 + 9.5 + 4.0 = 70.0` → **70**, no cap triggered
(`DeathScore = 84.4 ≠ 0`).

**Healer — Priest Holy, no deaths, one dropped dispellable debuff:**

| Component | Raw / percentile | Score |
|---|---|---:|
| Output | `metric_hps` (effective) 41st percentile — validated healer sim does not exist yet, so this *is* the percentile path directly on `metric_hps`, no execution-score step | 41 |
| Survival | No deaths; AvoidableHitScore 70th pct (a Holy priest standing in Rain of Fire once) → DeathScore 100 | `0.65*100 + 0.35*70 = 89.5` |
| Mechanics | 1 dispel made (Gehennas' Curse) value-weighted, high; owns no interrupt assignment on this table | 78 |
| Utility | Power Word: Fortitude 96% uptime (raid-wide, but only this priest's applications count toward their own share) | 85 |
| Preparation | Flask present, food present, no weapon oil (not applicable to casters) | 100 |
| Activity | 78% active share (healers idle between damage windows is expected; bracket already compares healer-to-healer) | 66th pct → 66 |

`overall = .30*41 + .10*89.5 + .20*78 + .20*85 + .10*100 + .10*66`
`= 12.3 + 8.95 + 15.6 + 17.0 + 10.0 + 6.6 = 70.45` → **70**

**Tank — Warrior Protection, one unavoidable death at 175,000 ms (from raid-wipe tank
death after a healer CC, classified unavoidable per this encounter — Vicious Headbutt-style
tank damage) plus two avoidable hits earlier (stood in a cleave meant for the other tank):**

An officer has marked `Assignment{PlayerKey: "<the other tank>", Job: "cleave duty", FromMS:
50000, ToMS: 70000}` (§2) — the raid's actual plan for this window. The cleave (`role:
"tank"`, spec §3.1's `Mechanic.Role`) lands on *this* tank at 60,000–65,000 ms, inside that
window, but the assignment names someone else for it: per §1.3's corrected rule, an
assignment overlapping the hit that does not name this player means the hit counts in full,
not the unconditional same-role excuse the uncorrected text described. Had no assignment
existed for this window at all, the same two hits would have been excused instead (§1.3's
case (a)) and AvoidableHitScore would read a percentile of `0`, not of `4000/180 ≈ 22.2`.

| Component | Raw / percentile | Score |
|---|---|---:|
| Output | tank's own `metric_dps`, 55th pct (small weight regardless) | 55 |
| Survival | 1 death: `penalty = 70*0.15*(1-175000/180000) = 70*0.15*0.028 = 0.29` → DeathScore 99.7; AvoidableHitScore (two hits, 4000 damage over the 180s fight ≈ 22.2/s, counted in full per the assignment above) 35th pct → 35 | `0.65*99.7 + 0.35*35 = 77.06` |
| Mechanics | Taunted off the add duty correctly (logged, not scored — see §1.3 Utility taunt note) | excluded from the ⅓ split with the remaining two parts reweighted |
| Utility | Sunder Armor 97% uptime, Demoralizing Shout n/a for Protection kit here | 91 |
| Preparation | Flask + food + weapon stone present | 100 |
| Activity | 85% active share | 70th pct → 70 |

`overall = .10*55 + .35*77.06 + .20*(Mechanics excluded → reweighted onto Output/Survival/
Utility/Preparation/Activity's shares proportionally, per §1.1) ...`

showing the renormalisation explicitly: Mechanics' 20 points redistribute across the other
five in proportion to their own §1.4 tank weights (`10, 35, 20, 10, 5` sum to 80; each gets
`+20 * (own/80)`): Output `+2.5→12.5`, Survival `+8.75→43.75`, Utility `+5→25`,
Preparation `+2.5→12.5`, Activity `+1.25→6.25` (all divided by 100 as usual, new total 100):

`overall = .125*55 + .4375*77.06 + .25*91 + .125*100 + .0625*70`
`= 6.875 + 33.71 + 22.75 + 12.5 + 4.375 = 80.21` → **80**

**CORRECTION (implementation pass, 2026-09-21):** this worked example originally read
Survival's combined score as `76.06` (an arithmetic slip — `0.65*99.7 + 0.35*35` is `77.06`,
not `76.06`) and the tank weight table's own Utility figure as `15` in the redistribution
step (transcribed wrong — §1.4's tank row gives Utility `20`, and the five non-Mechanics
weights `10+35+20+10+5` sum to `80`, not `75`). Both are fixed above; the corrected final
score is **80**, not 75. Verified against `logs/engine/rating`'s own test suite
(`TestCombineRenormalisesWeightsAcrossExcludedComponents` in `combine_test.go` and
`TestWorkedExampleTankWarriorProtection` in `score_test.go`), which reproduce this exact
arithmetic.

**CORRECTION (whole-branch review):** this example originally described the two avoidable
hits as excused unconditionally because the cleave was `role: "tank"` and this player is a
tank — §1.3's own text said the same, unconditionally, and both were wrong for the same
reason (the review's CRITICAL finding): a curated table's `role: "tank"` tag cannot say
*which* tank a hit was for, so the uncorrected rule credited the off tank standing in a
cleave exactly like the main tank doing their actual job. The example now includes the
`Assignment` that makes the hit demonstrably someone else's to take, so it counts in full;
the final numbers were already computed on the fully-counted value (`22.2`/s, 35th
percentile) and do not change.

---

## 2. Fairness details

**Deaths weighted by when and by cause.** The curve and the classification are §1.3's
`penalty_i` formula, restated for reference: linear in time-remaining, three discrete
cause multipliers (avoidable 1.00, unclassified 0.50, unavoidable 0.15). No table means no
Survival score at all (§1.1), not a silent "unavoidable" default — the brief's "unclassified
is neutral" applies to a *listed-but-uncertain* spell id inside a table that exists, not to
an encounter nobody has curated yet.

**Wipes score Survival and Mechanics but not Output.** `fight_metrics`/rankings already
exclude non-kills from the percentile digests today ("only kills feed the digests," §comment
at `api/internal/rankings/store.go:196-203`). Ratings extend the same rule one step
further: a wipe's Survival and Mechanics components are computed and stored (a wipe is
exactly where avoidable deaths and missed interrupts matter most, per the proposal's own
"Wipes count, differently" rule) but **do not feed the Output/Utility/Preparation/Activity
percentile digests**, and a wipe's Output component itself is excluded outright (not merely
unranked) since a wipe's damage window is truncated by definition and comparing partial
output to a full kill's is not a fair standard, absolute or percentile. A wipe's Mechanics
and Survival *do* feed their own digests, on the same `kill: false` rows — the digest bracket
key needs a `kill` boolean dimension so a wipe's Survival/Mechanics percentile is compared
against other wipes' Survival/Mechanics, not folded in with kills (§5.2 schema).

**Forced downtime windows.** New field on the mechanics table schema (§3.4), consumed
by Activity (§1.3) only — a downtime window is deliberately *not* subtracted from Survival's
or Mechanics' denominators, since "the boss is untargetable" does not mean "avoidable damage
stopped mattering" (add-phase cleaves, floor mechanics during a transition are exactly where
positioning still matters).

**Healers scored on effective healing.** Already true of `metric_hps` by construction
(§1.3's Output section); no new engine work.

**Tanks' Output weight small, Survival judged on avoidable damage only.** §1.4's weight
table (Output 10 for tanks) and §1.3's `AvoidableHitScore` role-exclusion rule
(`mechanics.go:44-47`'s existing `Role` semantics) are the exact mechanisms.

**The assignment hook.** An officer marking "this player is on decurse/interrupt-rotation/
add-duty/kiting for this fight" is officer tooling and out of scope (§9), but the coordinator
asked for the data shape to be settled now so that tooling can be built later without a
model change. The shape is deliberately small: an optional, per-fight, per-character list of
named time windows, supplied later by an officer:

```go
// Assignment is an officer-marked job for one player on one fight, supplied later by
// officer tooling that does not exist yet (§9). Score's assignments parameter is always
// an empty slice today — no caller populates it — but the signature accepts it now so
// officer tooling can be built against a stable contract without a model change once it
// exists.
type Assignment struct {
    PlayerKey string
    // Job is the officer's own free-text label for what this window covers —
    // "decurse", "interrupt rotation", "add duty", "kiting" — shown back on the card
    // wherever the exemption applies, never interpreted by the engine.
    Job    string
    FromMS int64
    ToMS   int64 // exclusive
}
```

During `[FromMS, ToMS)` for `PlayerKey`, **Output and Activity are measured only on the
player's free time**: the window is excluded from Output's percentile/absolute input the
same way a wipe's truncated window already is (§1.1), and from Activity's denominator the
same way a curated downtime window already is (§1.3's `downtimeOverlapMS`, which becomes
`Σ overlap(downtime windows ∪ this player's assignment windows, [0, timeAliveMS])` once
assignments exist — encounter-level downtime is shared by the whole raid, an assignment
window is one player's own).

**"The job is credited under Mechanics or Utility" needs no new formula.** Interrupts made,
dispels made, and buff/debuff uptime (§1.3's Mechanics and Utility components) are already
read straight off raid-wide event data (`ExchangeRow`, `AuraTrack`) with no reference to
*why* the player was doing it — a dispel made during an assignment window already counts
under Mechanics today, with or without the assignment existing. The assignment's only actual
effect on scoring is the Output/Activity exemption above; "credited under Mechanics or
Utility" describes behaviour the existing formulas already have, not a formula this spec
still owes.

`Assignment` is threaded through `Score(..., assignments []Assignment)` (§4.1's engine
signature) as an optional, possibly-empty slice; every formula above already reads correctly
with it empty, which is its state on every fight until officer tooling exists.
**RULING: the hook is a parameter, not a database table this spec creates.** Creating
`guild_assignments` now, unused, would be dead schema per the user's own coding-style rule
against speculative structure; the Go type above is the contract the officer-tooling spec
(whenever it is written) implements storage for, and the scoring package's public function
signature is what makes that a config change rather than a scoring-engine rewrite.

---

## 3. New curated data and its schema

### 3.1 Storage convention

Every new curated file lives beside the existing mechanics tables, in the same package,
following the same embed-and-validate-at-build pattern
(`logs/engine/mechanics/mechanics.go:102-135`): a panic on a malformed file at process
start, never a runtime surprise. Files:

- `logs/engine/mechanics/utility/tables/<spec-slug>.json` — one file per spec (27 files,
  matching `data/curated/specs.json`'s 27 entries).
- `logs/engine/mechanics/consumables/catalogue.json` — one file, all roles.
- `logs/engine/mechanics/tables/<encounter_id>.json` — the **existing** mechanics table
  file gains an optional `downtime` array (§3.4); no new file, an additive field, so every
  existing table stays valid with zero edits (`downtime` omitted means none, same pattern
  `phases` already uses — `mechanics.go:99`, `phases,omitempty`).
- `logs/engine/mechanics/weights/roles.json` — the default weight table (§1.4), as data,
  so the (out-of-scope) officer per-guild override can read the same shape a guild's
  override would write.

### 3.2 Per-spec utility table

```json
{
  "spec": "warrior-protection",
  "owned": [
    { "spell_id": 1160, "name": "Demoralizing Shout", "kind": "debuff", "target": "enemy",
      "verified": "1.60.1.69893/spells.json#1160", "note": "Party-wide melee attack power debuff." },
    { "spell_id": 7386, "name": "Sunder Armor", "kind": "debuff", "target": "enemy",
      "verified": "1.60.1.69893/spells.json#7386" },
    { "spell_id": 20647, "name": "Improved Sunder Armor", "kind": "talent-modifier",
      "verified": "1.60.1.69893/spells.json#20647", "note": "Not independently trackable as an aura; informational only, not scored." }
  ]
}
```

`spell_id` is verified against `data/builds/1.60.1.69893/spells.json`, a flat
`[{id, name}]` list of 31,754 rows (checked directly: `python3 -c "..." → <class 'list'>
31754`, row 1 `{'id': 1, 'name': 'Word of Recall (OLD)'}`) — the same source
`logs/engine/mechanics/encounters.json`'s own provenance note points at for spell ids
("Every id below was read from the Forever beta client itself," `encounters.json`'s
top-level `note` field). **Verification procedure:** the curator looks up the id in
`spells.json`, confirms the name matches (a mismatch means the wrong client/family — see
`encounters.json`'s own worked example of exactly this problem, "Flame Spear and Shazzrah's
Curse are the two abilities... whose ids the first real log should be checked against"),
and records the source as `verified: "<build>/spells.json#<id>"` in the entry. An id that
does not resolve is not committed; the drafting tool (§3.5) flags any spell id it proposes
that `spells.json` does not contain.

`kind` is one of `buff` (self/raid buff the spec applies), `debuff` (enemy debuff the spec
applies), `cooldown` (a burst/utility cooldown whose *use*, not uptime, is what should be
tracked — e.g. Innervate, Power Infusion; scored by cast count against fight length rather
than by `AuraTrack.UptimeMS`), or `talent-modifier` (informational, never independently
scored — most specs will have a handful of these the curator records for completeness and
marks `"note"` explaining why it is not scored, so a reviewer does not wonder why it is
missing from the card).

### 3.3 Consumable catalogue by role

```json
{
  "roles": {
    "dps": {
      "weights": { "flask": 30, "food": 15, "weapon_enchant": 15, "world_buffs": 20, "combat_potion": 20 },
      "flask": [ { "spell_id": 17627, "name": "Flask of the Titans", "verified": "1.60.1.69893/spells.json#17627" }, "..." ],
      "food": [ "..." ],
      "weapon_enchant": [ "..." ],
      "world_buffs": [ { "spell_id": 22888, "name": "Rallying Cry of the Dragonslayer", "verified": "..." }, "..." ],
      "combat_potion": { "max_uses": 2, "entries": [ "..." ] }
    },
    "healer": { "weights": { "flask": 30, "food": 15, "weapon_enchant": 0, "world_buffs": 20, "combat_potion": 35 }, "...": "..." },
    "tank": { "weights": { "flask": 30, "food": 15, "weapon_enchant": 15, "world_buffs": 20, "combat_potion": 20 }, "...": "..." }
  }
}
```

Each role's `weights` sum to 100 (validated at parse, same pattern as
`mechanics.Parse`'s field checks, `mechanics.go:138-185`). `weapon_enchant` is `0` for
casters (no melee weapon to oil) — the catalogue, not the scoring code, is where "this
category does not apply to this role" lives, so a future role split (e.g. a caster DPS vs.
a melee DPS getting different weights) is a data change.

### 3.4 Downtime windows — extends the existing mechanics table

```json
{
  "encounter_id": 666,
  "name": "Garr",
  "mechanics": [ "...unchanged..." ],
  "downtime": [
    { "trigger": { "spell_id": 19497, "on": "aura_applied" }, "duration_ms": 3000,
      "note": "A Firesworn's death detonation window: melee are expected to step out, not to keep swinging." }
  ]
}
```

`trigger` reuses `mechanics.PhaseStart`'s exact shape (`mechanics.go:73-90`: either a
`spell_id`+`on` pair or a `health_pct`) — **RULING: reuse `PhaseStart`, do not invent a
second trigger vocabulary.** The engine already has one validated, tested way to say "this
fires when," and a downtime window is structurally identical to a phase boundary (an
instant, named by a cast/aura/health event) except it also carries a `duration_ms` and does
not rename the fight's phase. Cost if wrong: none foreseen — the two concepts share
literally the same detection code path in `summary.phaseFires` (`phases.go:47-71`), which
the engine package's downtime detector calls unmodified.

### 3.5 Who curates, how the drafting tool helps, what happens with no table

Curation is a person's judgment call over the drafting tool's evidence, exactly as the
existing mechanics tables already work (`mechanics.go:6-7`'s own package doc: "a curator
corrects it before it lands here"). `logs/cmd/forever-logs`'s `mechanics-draft` subcommand
(`logs/cmd/forever-logs/mechanics_draft.go:27-60`) already classifies avoidable vs.
unavoidable from evidence (non-tank hit counts, `mechanics_draft.go:151-170`) and proposes
phase candidates from enemy casts/auras and boss health readings
(`mechanics_draft.go:280-320`+). This spec extends it, not replaces it, with two more
evidence passes run over the same drafted fights:

- **Utility evidence**: for every player's spec, which auras they were the source of that
  are *not already* in the fight's own buff/debuff vocabulary the drafter already reads —
  cross-referenced against `data/builds/1.60.1.69893/spells.json` by name, proposed as
  candidate `owned` entries with a `"draft:"`-prefixed note (matching the existing
  convention at `mechanics_draft.go:216`'s `evidence` string) naming how many fights showed
  the aura and its mean uptime.
- **Downtime evidence**: any enemy cast/aura the drafter already collects as a phase
  candidate (`mechanics_draft.go:280-320`) that repeats *within* a single phase rather than
  opening a new one, with a consistent duration before raid activity resumes, is proposed as
  a downtime window candidate instead of a phase candidate — the same underlying evidence
  (enemy casts + damage-taken drop-off), read through the second lens.

**An encounter with no table yet still gets ratings.** Per §1.1's exclusion rule, Output
(execution-score or raw-DPS path), Preparation, and Activity (whole-fight fallback) need no
mechanics table at all; only Survival, Mechanics, and Utility (which needs a *spec* table,
not an *encounter* table) are gated on encounter curation. A brand-new boss on patch day
therefore ships with three of six components live from the first kill, and the card says so
in the same words §7.4 defines for any excluded component ("Mechanics: not scored — this
fight's encounter has no curated table yet").

---

## 4. Where it runs and where it is stored

### 4.1 A new engine package, pure functions over a summary plus curated tables

`logs/engine/rating/` (new package, sibling to `summary` and `mechanics`, not a change to
either). **RULING: a new package, not a method on `summary.Accumulator`.** The rating
engine needs a *finished* `summary.Summary` (deaths, mechanics block, auras — all
post-`Snapshot`) plus curated data `summary` has no reason to import (utility tables,
consumable catalogue, weights); folding it into the accumulator would make `summary` depend
on rating-specific curated data it does not otherwise need, and would make every summary
computed during a live-tailed fight (`Accumulator.Snapshot` is called "every few seconds
during a live fight," `summary.go:398-401`) redo rating math it cannot yet finish (deaths
and mechanics are still accumulating). A rating is a **post-fight** computation over an
already-closed fight's summary, same lifecycle stage as the execution scorer
(`sims/score.go` runs after fight close too).

```go
package rating

// Card is one player's rating for one fight: the product the brief calls "never a bare
// number." Six Components plus Overall; the six always list six entries even when a
// component is excluded (Excluded: true, Reason set), so the web layer never has to
// special-case "missing from the array" vs. "excluded."
type Card struct {
    // Overall is the site-default figure: OverallUncapped, with §1.5's cap applied if
    // OverallCapped fired. OverallUncapped is always stored too, so a guild-adjustable
    // setting that turns the cap off — or a future recompute of the cap's own threshold —
    // reads it straight back with no backfill (§1.5, §4.2).
    Overall         float64
    OverallUncapped float64
    OverallCapped   bool     // §1.5's Survival-catastrophe cap condition fired this fight
    Components      [6]Component
    Basis           string   // "percentile" | "absolute" | "mixed" (per-component; see Component.Basis)
    ModelVersion    string
    KillTimeBand    string   // "fast" | "typical" | "slow" | "" (excluded/wipe)
}

type Component struct {
    Name       string  // "output" | "survival" | "mechanics" | "utility" | "preparation" | "activity"
    Score      float64 // 0-100; meaningless if Excluded
    Weight     float64 // the *renormalised* weight actually applied
    Basis      string  // "percentile" | "absolute" | ""
    Percentile *float64
    BracketN   int64
    Excluded   bool
    Reason     string  // set iff Excluded; one of the fixed strings in §7.4
    Moments    []Moment // the "opens into the specific moments" data; see §7.1
}

// Score computes one player's Card. Every input is already-loaded data; Score makes no
// network or database call, which is what makes it unit-testable with golden fixtures
// (§8) the same way summary.Snapshot is.
func Score(
    fight summary.Summary,
    player string, // GUID
    tables CuratedTables, // mechanics.Table (existing) + the three new files (§3), preloaded
    percentiles PercentileSource, // an interface over api/internal/digest's Placement, injected so the engine package never imports api/internal
    assignments []Assignment, // §2's hook; always an empty slice today — no caller populates it
    now ModelInfo, // ModelVersion string, the weights table, and the cap's enabled flag and threshold (§1.5) — every guild-adjustable setting lives here
) Card
```

`PercentileSource` is an interface (`Placement(bracket Bracket, value float64) (pct float64,
n int64, ok bool)`) rather than a direct `api/internal/digest` dependency, so `logs/engine`
— which today has zero dependency on `api/internal` (verified: `logs/engine/mechanics`
imports only `embed`/`encoding/json`/`fmt`/`strconv`, `mechanics.go:10-15`) — stays that
way. The API layer's job wires a concrete `PercentileSource` backed by
`api/internal/digest.Unmarshal`/`Placement` (`digest.go:212-236`, `253-...`) over rows read
from the new `rating_percentile_digests` table (§4.2).

### 4.2 Storage schema

Migration `0019_ratings.up.sql`. `0017_reports_recent_idx` is the latest migration landed in
this checkout (confirmed via `ls api/internal/db/migrations/*.up.sql | sort | tail`); two
parallel specs claim the numbers immediately after it — `0018` is the guild-membership spec
(`2026-09-21-guild-and-premium-proposal.md` §3.1) and `0020` is the payments spec (§4.2 of
the same proposal) — so this one lands as `0019`. §8 notes that the coordinator assigns the
final number at merge if the three specs land in a different order than this.

```sql
-- Player ratings: one row per (report, fight, player), partitioned by fought_at exactly
-- as fight_metrics is (api/internal/db/migrations/0005_logs.up.sql), so the same monthly
-- partition job and the same EnsureMetricsPartition-shaped helper serve both. The engine's
-- Card (logs/engine/rating) is stored whole as components; overall/overall_uncapped are
-- denormalised for the sortable/filterable columns a query needs without unpacking jsonb.
-- Both the capped and uncapped figures are kept (coordinator ruling, §1.5): the cap is a
-- guild-adjustable display setting over an already-computed fight, not a re-judgment of
-- it, so toggling it — or changing its threshold later — reads back from this row with no
-- recompute and no backfill.

create table if not exists rating_scores (
  report_id      text not null,
  fight_index    int not null,
  player_key     text not null,
  player_name    text not null default '',
  class          text,
  spec           text,
  role           text,
  encounter_id   int,
  difficulty     int,
  size           int,
  duration_ms    int,
  kill           boolean not null,
  kill_time_band text not null default '',
  overall           numeric,  -- site-default figure: overall_uncapped, capped at 40 if overall_capped
  overall_uncapped  numeric,  -- the plain weighted mean, §1.5, with no exception applied
  overall_capped    boolean not null default false, -- whether §1.5's cap CONDITION fired this fight
  components     jsonb not null,   -- [6]rating.Component, exactly as the engine returns it
  model_version  text not null,
  fought_at      timestamptz not null,
  computed_at    timestamptz not null default now(),
  primary key (report_id, fight_index, player_key, fought_at)
) partition by range (fought_at);

create index if not exists rating_scores_player_idx on rating_scores (player_key, fought_at desc);
create index if not exists rating_scores_report_idx on rating_scores (report_id, fight_index);
create index if not exists rating_scores_stale_idx on rating_scores (model_version) where model_version <> '';

-- One digest per (encounter, difficulty, spec, role, kill_time_band, kill, component),
-- the same shape as percentile_digests (0005_logs.up.sql) with role/kill_time_band/kill
-- added per §1.2 and §2's wipe rule, keyed by component name rather than "metric" since a
-- rating percentile is always of a component's own 0-100-or-raw value, never dps/hps/taken.
create table if not exists rating_percentile_digests (
  encounter_id   int not null,
  difficulty     int not null,
  spec           text not null,
  role           text not null,
  kill_time_band text not null,
  kill           boolean not null,
  component      text not null,
  digest         bytea not null,
  n              bigint not null default 0,
  updated_at     timestamptz not null default now(),
  primary key (encounter_id, difficulty, spec, role, kill_time_band, kill, component)
);

-- The tiny digest kill_time_band buckets against: duration_ms of kills only, per
-- (encounter_id, difficulty). Read at write time to classify a fight's band (§1.2).
create table if not exists kill_duration_digests (
  encounter_id int not null,
  difficulty   int not null,
  digest       bytea not null,
  n            bigint not null default 0,
  updated_at   timestamptz not null default now(),
  primary key (encounter_id, difficulty)
);
```

`EnsureMetricsPartition` (`api/internal/db/partitions.go:28-38`) hard-codes `partition of
fight_metrics` in its `fmt.Sprintf`; this spec generalises it to
`EnsurePartition(ctx, pool, table string, at time.Time)` with `EnsureMetricsPartition`
becoming a one-line wrapper (`EnsurePartition(ctx, pool, "fight_metrics", at)`), so
`rating_scores` reuses the exact same monthly-partition logic and the exact same startup
call site rather than a forked copy.

### 4.3 Invocation: where in the parse job

`api/internal/reports/ingest.go`'s per-fight verify handler already runs, in order,
`i.rank(...)` (writes `fight_metrics`, `ingest.go:365-379`) then `i.score(...)` (schedules
execution scores for signed-in members only, `ingest.go:396-421`). This spec adds a third,
parallel call, `i.rate(...)`, in the same best-effort, never-blocks-storage spirit as
`i.score` (`ingest.go`'s own comment: "It is deliberately best-effort and returns nothing"):

```go
// after i.score(r.Context(), rep, n, f.EncounterID, derived):
i.rate(r.Context(), rep, n, f.EncounterID, f.Start, derived, rebuilt.Combatants)
```

**Unlike `i.score`, `i.rate` runs for every roster player, not signed-in members only.**
`i.score`'s member-only gate exists because the execution scorer needs to *rebuild a
simulator character* (`CombatantBuilder.FightCharacter`, `sims/score.go:91-121`), which
today always fails anyway (no race data — `score.go:119`, "the fight records no race").
Rating's Output component reads `execution_score` if `fight_metrics` already has one
(computed independently by the existing scorer, for whichever subset of players qualify)
and falls back to raw `metric_dps`/`metric_hps` otherwise (§1.3) — nothing about computing a
*rating* requires the player to be a claimed, signed-in character, so gating it the same way
would silently withhold ratings from every non-member in a pug, which contradicts the
proposal's "public the way parses are" rule directly. **RULING: rate every roster row;
`i.score`'s member gate is specific to needing a simulator character, not a rating
precondition.** Cost if wrong: computing a rating for a player nobody has claimed is
"wasted" work if they never look it up — but the alternative (member-gated ratings) makes
the free single-lookup promise a lie for exactly the pug-leader-checks-an-applicant case the
proposal calls out as the main reason people use a logs site
(`2026-09-21-guild-and-premium-proposal.md:125`).

`i.rate` is skipped under the same conditions `i.rank`/`i.score` already are — no ranker
wired, no `EncounterID` (trash), or `!Ranked(rep.Visibility)` (private reports never
compute ratings at all, matching `Ranked` at `reports/report.go:44-45`, and matching §2's
wipe rule for whether a fight's rows *feed digests* — a private report's ratings are simply
never computed, full stop, rather than computed-but-hidden, since there is no card to
show anyone but the owner and the owner already sees the full report).

### 4.4 Recomputation: `model_version` and a backfill job

`ModelInfo.ModelVersion` (§4.1's `rating.Score` signature) is a short string
(`"rating-2026-09-21"`, date-stamped like the mechanics tables' own provenance convention)
bumped whenever a formula, a weight, or **any** curated table changes in a way that could
change a stored score. **RULING: one global version, not per-encounter or per-component.**
A curated-table edit to one boss technically only invalidates that boss's stored ratings,
but tracking per-table versions multiplies the bookkeeping (which version applied to which
table, cross-referenced against which stored row) for a problem a single global version
solves adequately: the backfill job below is cheap enough that over-invalidating (recomputing
ratings a table edit did not actually touch) costs compute, not correctness, and compute is
bounded (§4.5). Cost if wrong: unnecessary backfill runs after a small table tweak; the job
is idempotent and safe to run as often as needed, so this is a cost-not-correctness tradeoff
by design.

**Backfill job.** A Cloud Run job (the same shape as the existing `sim-validate` job the
proposal already references for the nightly snapshot,
`2026-09-21-guild-and-premium-proposal.md:166`), triggered manually after a
model/weights/table change and nightly as a sweep: selects `rating_scores` where
`model_version <> current`, re-reads each fight's stored summary via the **same** `Getter`/
`fightSummary` path the execution scorer already uses
(`api/internal/sims/summaries.go:19-42`), recomputes via `rating.Score`, and upserts. Digest
tables (§4.2) are **not** rebuilt from scratch on a backfill — same tradeoff `RemoveReport`
already accepts for `percentile_digests` ("a t-digest cannot have a value taken back out...
the alternative... would cost far more than one tampered report distorts,"
`rankings/store.go:303-320`); a formula change shifts the digest's shape slightly until
enough new fights are folded in, which the confidence band (§1.2's bracket-size copy)
already communicates honestly.

### 4.5 Cost

Per fight: one JSON summary read (already paid by `i.score` when it runs; `i.rate` pays it
independently since the two calls do not currently share state — a follow-up could thread
the already-decoded summary through both, but `i.rank`/`i.score`/`i.rate` are three
sequential handler-local calls each starting from `derived`/`rebuilt.Combatants`, already in
memory from the verify handler's own earlier work at `ingest.go:280-300`, so no *extra*
summary read is needed at all — only `i.score`'s simulator run, which `i.rate` does not
call, is the expensive part it avoids). Per-player cost is arithmetic over already-loaded
structs plus up to six digest reads (`Store.Percentile`-shaped queries, each a single-row
`select` by primary key, `rankings/store.go:329-352`'s existing pattern) and up to six
digest writes (`updateDigest`'s existing lock-insert-update pattern,
`rankings/store.go:255-302`, reused verbatim for the new tables). A 40-player raid fight is
therefore ~240 small index-scoped reads and ~240 locked upserts per fight-close, the same
order of magnitude `WriteFight` already performs today for `fight_metrics`
(`rankings/store.go:129-224`'s per-row insert plus per-metric digest fold). No new
infrastructure; runs on the existing API Cloud Run service, in the same request that already
writes `fight_metrics`.

---

## 5. API

### 5.1 Endpoints

```
GET /v1/reports/{id}/fights/{n}/ratings
GET /v1/characters/{region}/{ruleset}/{name}/rating
```

**`GET /v1/reports/{id}/fights/{n}/ratings`** — the per-fight report card, every player on
the roster. Visibility: **reuses `Service.mayView(r, rep)` unmodified**
(`reports/handler.go:521-538`) — this is the whole rule, stated as the brief asks,
testably: *public and unlisted reports' ratings are visible to anyone with the report link;
private reports' ratings are visible to the owner and moderators; guild reports' ratings are
visible to the owner, moderators, and members of that guild (`Accounts.GuildRank`)*. No new
access-control code; the rating endpoint sits in the `reports` service and calls the same
function the report body already calls.

**`GET /v1/characters/{region}/{ruleset}/{name}/rating`** — the aggregate on a character
page: overall trend and best/worst component over the character's recent rated fights.
**RULING: this endpoint includes only fights whose report was `public` at time of
query — never unlisted, guild, or private, even though `mayView` would let some of those
through for the *right* viewer.** The character page has no report-scoped access context
(it is not "this viewer, looking at this specific report" — it is "anyone, looking at this
character"), so it cannot ask "does this viewer have this specific unlisted link." This
matches the proposal's own sentence precisely: "a rating computed from a public report shows
on the player's character page... Ratings from private, unlisted or guild-only reports are
visible only to whoever can already see those reports" — the character page is not "whoever
can already see" an unlisted report (an unlisted link is not the character page), so it is
correctly excluded. **Cost if wrong:** if this ruling is too conservative, a guild's own
character-page view of their own raiders shows fewer rated fights than a roster check would
(a premium, in-scope-later feature); that gap is exactly what "premium is scale and depth,
not the lookup itself" (`2026-09-21-guild-and-premium-proposal.md:132`) already describes as
the paid feature's value, so under-including here is the correct default, not a bug to widen
casually.

`anonymize` — **§0 found that nothing today filters rankings/character pages by it.** This
spec does not invent a new mechanism to retrofit onto existing rankings (out of scope: that
is a `fight_metrics`/`rankings` gap, not a ratings one), but ratings must not make the gap
worse by being the first surface that actually honours the flag inconsistently with
everything beside it on the same page. **RULING: both rating endpoints join
`characters.user_id → users.anonymize` (the same join `auth.User.PublicName` would need if
it were applied to a character rather than a report owner) and, when `anonymize = true`,
answer `404 not_found` for the character-rating endpoint and omit that player's row entirely
from the per-fight report card's roster** (rather than a name-masked placeholder — a report
card with an unnamed row a viewer can still map back to a real player by class/spec/gear is
not actually anonymous). **This is stricter than the current, unenforced rankings behaviour**
and is flagged here explicitly as a known inconsistency for a follow-up (not this spec's
scope) to bring rankings/character pages up to the same standard, rather than an argument for
ratings to match rankings' current gap. Cost if wrong: an anonymize-set player's rating
becomes literally unlookupable even by someone who could otherwise see the report (e.g. an
officer who already knows who's who) — the alternative (name-masked but still shown) risks
deanonymizing by class+spec+gear combination in a small raid, which is the worse failure for
a flag whose entire purpose is "don't show my performance publicly."

### 5.2 Response shapes, in the existing envelope

Both endpoints answer inside `httpx.Envelope` (`api/internal/httpx/envelope.go:28-33`):
`{ ok, data, error, request_id }`.

```json
// GET /v1/reports/{id}/fights/{n}/ratings → data:
{
  "fight_index": 3,
  "kill": true,
  "kill_time_band": "typical",
  "model_version": "rating-2026-09-21",
  "players": [
    {
      "player_key": "us/normal/simfury",
      "player_name": "Simfury",
      "class": "Warrior", "spec": "Fury", "role": "dps",
      "overall": 70, "overall_uncapped": 70, "overall_capped": false,
      "basis": "percentile",
      "components": [
        { "name": "output", "score": 71, "weight": 35, "basis": "percentile",
          "percentile": 71.2, "bracket_n": 142, "excluded": false, "moments": [] },
        { "name": "survival", "score": 76, "weight": 15, "basis": "percentile",
          "percentile": 60.0, "bracket_n": 89, "excluded": false,
          "moments": [ { "kind": "death", "at_ms": 140000, "spell_id": 19712,
                          "spell_name": "Arcane Explosion", "avoidable": true,
                          "anchor": "death-<guid>-140000" } ] }
        // ... mechanics, utility, preparation, activity
      ]
    }
    // ... every roster player mayView/anonymize allows
  ]
}
```

```json
// GET /v1/characters/{region}/{ruleset}/{name}/rating → data:
{
  "player_key": "us/normal/simfury",
  "sample_size": 34,
  "trend": [ { "fought_at": "2026-11-10T02:14:00Z", "overall": 62, "report_id": "...", "fight_index": 2 }, "..." ],
  "best_component": "preparation",
  "worst_component": "activity",
  "latest": { "...": "same per-fight shape as above, one player" }
}
```

`components` always has six entries (§4.1's `Card.Components [6]Component`); an excluded
entry carries `"excluded": true, "score": null, "reason": "no_mechanics_table"` — the fixed
reason strings are §7.4's.

`overall` and `overall_uncapped` are equal whenever `overall_capped` is `false`; the web
layer (§6.1) shows `overall` by default and, when `overall_capped` is `true`, the §6.4 copy
naming what capped it. Officer tooling with the cap turned off for its guild (§1.5) reads
`overall_uncapped` instead — a read-time substitution both endpoints already support with no
API change beyond serving the field that is already on the row.

### 5.3 Caching

`Cache-Control: public, max-age=30` on both, matching the existing rankings routes'
`cache(w)` helper (`rankings/handler.go:58-63`, `cacheSeconds = 30`) — ratings are exactly
as volatile as a rankings page (a new fight can update a bracket's percentile at any time)
and reuse the same edge-cache policy rather than inventing a different number. The per-fight
endpoint additionally sets `Cache-Control: private` when the report's visibility is anything
but `public` (an unlisted/guild/private card must not be cached at a shared edge for a
different viewer to receive) — the same distinction the report body's own visibility-gated
route already needs to make, if it does not already (flagged for the implementer to confirm
against the existing report-body caching, which this spec did not audit).

---

## 6. Web

### 6.1 The report card inside a report

New tab, `'rating'`, added to `web/src/lib/report/url.ts`'s `Tab` union
(`url.ts:14-25`), so the card is reachable at `?fight=<n>&tab=rating` and — combined with
the existing `source=<guid>` scoping param the URL scheme already defines — at
`?fight=<n>&tab=rating&source=<guid>` for one player's card. A `RatingPanel.svelte` mirrors
`SummaryPanels.svelte`'s dashboard pattern (`SummaryPanels.svelte:1-4`, a short bar list per
category linking into its own tab) on `SummaryTab`: one line per roster player, overall
score plus a compact six-segment bar, linking to `tab=rating&source=<guid>` for the full
card. The full `RatingTab.svelte` (new component) follows `ActorRow.svelte`'s existing
`<details>`/`aria-expanded` expansion pattern (`ActorRow.svelte:339`) for each of the six
parts: **one number first** (the `overall`), **then the six parts** as a row of expandable
segments, **then the moments** inside each expanded segment — the exact three-level
structure the proposal's rule 7 requires ("one number first, then the parts... opens into
the moments in the log behind each part").

**Jumping to a moment.** Survival's death moments link directly to the Deaths tab's
existing per-death DOM anchor, `id={death-${guid}-${at_ms}}`
(`DeathsTab.svelte:294-295`) — already built, reused as-is via
`?tab=deaths#death-<guid>-<at_ms>`. Mechanics' interrupt/dispel moments and Utility's
dropped-buff moments have **no existing per-event anchor** (`ActorRow`/`CastTable`/
`ExchangeTable` render lists with no per-row `id`) — this is a real gap, not assumed away:
this spec asks the web lane to add a `data-testid`/`id` anchor convention to
`ExchangeTable.svelte` and `AuraTable.svelte` rows (`id={exchange-<sourceGuid>-<atMs>}`,
mirroring the Deaths tab's own naming) as a small, additive change alongside the rating tab,
not a rating-specific feature — every other tab benefits from the same deep-link
capability going forward.

### 6.2 The character page rating, with trend

`web/src/components/Character.svelte` already fetches `CharacterPage` via
`fetchCharacter` (`Character.svelte:8`); this spec adds a `fetchCharacterRating` call
(new function in `web/src/lib/rankings/api.ts`, same module `fetchCharacter` lives in)
rendered as a new panel beside the existing progression/roster-bests panels: overall score
big, six-segment bar, a small trend sparkline over `trend[]` (reusing whatever sparkline
primitive the codebase already has for a per-second `Series` — `TimeChart.svelte` is the
existing time-series renderer and is the natural component to point the trend chart at,
confirmed present at `web/src/components/report/TimeChart.svelte`). Empty/low-sample states
per §7.4.

### 6.3 The explanation page

`/ratings` (new route, `web/src/pages/ratings.astro` or equivalent, mirroring how
`/premium` is scoped as a single reference page in the guild proposal): publishes the exact
weight table (§1.4), the exact formulas in plain language (not the Go/SQL of this spec, the
reader-facing translation of §1.3), the percentile-vs-absolute rule in one paragraph, the
cap in the coordinator's own words — *"An avoidable death early in a fight caps the overall
score at 40, because nothing else in the fight makes up for it"* — named alongside the
weights as one of the settings a guild can change (§1.5), and a link from every rating
surface ("How is this calculated?") — **no hidden weights**, per the proposal's own rule 5.

### 6.4 Copy for every state

- **Scored, percentile basis:** *"71 — better than 71% of Fury Warriors on this fight
  (142 logs)."*
- **Scored, absolute basis:** *"70 — measured against the encounter's numbers. Not enough
  logs yet to compare players."*
- **Excluded, no table:** *"Mechanics — not scored. This fight's encounter has no curated
  mechanics table yet."*
- **Excluded, spec not modeled:** *"Output — not scored. The simulator does not model this
  spec yet; scores return once it does."*
- **Excluded, threat not modeled (Utility sub-part):** *"Threat — not modeled yet. This
  does not count for or against Utility."*
- **Capped by Survival:** *"70, capped from a higher weighted average — a costly avoidable
  death outweighs the rest of the fight. See Survival."* (coaching tone, not shaming — see
  §8).
- **Low sample / new player:** *"Not enough rated fights yet to show a trend (2 of 5
  needed)."*
- **Anonymized:** the row is omitted; no placeholder text is shown (§5.1's ruling — nothing
  on the page implies "someone is hidden here").

### 6.5 Phone layout and the design system

No emoji, no marketing buttons, restrained motion, 44px hit targets
(`2026-09-21-one-product-design.md:16-17`'s rules for every lane, which this spec's web work
is bound by exactly as every other lane is). The six-segment bar collapses to a stacked list
below `md` width rather than a horizontal bar (matching the existing nav's own
phone-collapse pattern, `2026-09-21-one-product-design.md:131-135`). Colors: score bars use
`--gold`/`--gold-hover` for the overall figure (matching every other primary numeric
headline on the site, `design/DESIGN-SYSTEM.md`'s accent section) and **not** a red/green
traffic-light scheme for the six parts — a low score reads as muted text plus the coaching
copy above, not a red bar, per the tone rule in §8.

---

## 7. Abuse and tone

**No public leaderboard of ratings.** No endpoint in §5 returns "top" or "bottom" N players
sorted by overall rating across the site; `GET /v1/reports/{id}/fights/{n}/ratings` returns
one fight's roster (already bounded by who was in that raid), and the character endpoint
returns one player. Nothing in this spec's data model prevents an officer-tooling feature
from building a *guild's own* raid-night sheet (the proposal's explicit, paid officer
feature) — the guard is specifically **no cross-guild, cross-raid ranked list**, which no
endpoint here provides even as raw material (there is no "give me every rating above N"
query).

**Coaching copy, not shaming.** §6.4's exact strings model the tone: a low Survival score
reads as *"70, capped from a higher weighted average — a costly avoidable death outweighs
the rest of the fight. See Survival."* — names the mechanism, points at the fix, never says
"bad," "failed," or ranks the player against a judgment word. The per-part moments (§6.1)
show the specific hit/death/dropped-buff with its timestamp, which is the coaching content;
the number is the headline, the moment is the lesson.

**Gaming resistance.**

- **Padding** (chasing raw damage instead of playing well): Output's primary path is
  `execution_score` (a ratio against the player's *own* simulated potential, not a raw
  number to inflate by any means available) wherever the simulator models the spec; the raw-
  DPS fallback only applies where it does not, and even there it is percentile-ranked
  against the same-spec/same-bracket peers, not an absolute number a player can pad against
  no ceiling.
- **Dying on purpose to skip a mechanic** (e.g. a healer letting themselves die rather than
  soak an unavoidable mechanic): the death's `causeMultiplier` reads the killing blow's own
  classification (§1.3) — a death to an *unavoidable* mechanic already costs only 0.15× the
  penalty of an avoidable one, so "dying to the unavoidable thing on purpose" is nearly free
  to Survival already, which sounds like a hole, but is closed by Activity and Output: a
  player dead for the rest of the fight posts near-zero Activity and Output for that
  stretch (both computed over `timeAliveMS`, §1.3), which the weighted mean (§1.5) still
  punishes across the components that *aren't* Survival. **RULING: this is treated as
  already handled by the existing formula shape, not given a dedicated anti-gaming rule.**
  Adding an explicit "did they die suspiciously" detector is speculative complexity against
  a pattern the component interactions already discourage; if real raid logs show this
  gamed in practice, the fix is very likely just Activity's weight, not new logic.
- **Sitting out wipes** (only logging/attending kills to keep an average up): ratings are
  computed and stored for wipes too (§2), including Survival and Mechanics — a player who
  never appears on wipe logs at all is a roster/attendance question the (out-of-scope)
  officer attendance tool answers, not something the rating number itself can detect from
  inside one fight's data. Flagged as a known limitation, not solved here.

---

## 8. Testing — lane split

Three lanes, file ownership, order constraints, and what December 9 actually needs.

| Lane | Owns | Depends on |
|---|---|---|
| **Engine + curated data** | `logs/engine/rating/` (new), `logs/engine/mechanics/utility/`, `logs/engine/mechanics/consumables/`, `logs/engine/mechanics/weights/`, `downtime` additions to `logs/engine/mechanics/tables/*.json`, the extended `mechanics-draft` evidence passes (`logs/cmd/forever-logs/mechanics_draft.go`) | Nothing new (existing `summary`/`mechanics` packages) — can start first |
| **API + storage + job** | migration `0019_ratings...sql`, `api/internal/rating/` (new: handler, store, the `PercentileSource` adapter over `api/internal/digest`), `i.rate` in `api/internal/reports/ingest.go`, the backfill job, `EnsureMetricsPartition` generalisation in `api/internal/db/partitions.go` | Engine lane's `rating.Card`/`rating.Score` signature (§4.1) — can start against the signature before the engine lane lands, per the simulator-parity spec's own precedent ("if Lane B has not landed, Lane A builds against the signatures," `2026-09-21-one-product-design.md:88-90`) |

**Migration numbering.** `0019` assumes this spec's migration lands after the guild-
membership spec's `0018` and before the payments spec's `0020` (§4.2's note). If the three
land in a different order, the coordinator assigns the final number at merge time; nothing
in this spec's schema depends on the specific number beyond the filename itself.
| **Web** | `web/src/lib/report/url.ts`'s `Tab` addition, `RatingPanel.svelte`, `RatingTab.svelte`, `web/src/lib/rankings/api.ts`'s rating fetchers, `Character.svelte`'s new panel, `/ratings` explanation page, the `ExchangeTable`/`AuraTable` anchor additions (§6.1) | API lane's response shapes (§5.2) — can build against the shapes before the API lane lands |

**Order constraints:** engine lane's `rating.Card`/`Component` types and `rating.Score`
signature (§4.1) must land — even as a stub returning excluded-everywhere cards — before the
API lane can compile its store layer against real types rather than a hand-rolled mock;
after that, all three lanes proceed in parallel. The curated-data sub-lane (utility tables,
consumable catalogue, downtime windows) is explicitly **not** a blocker for the engine
package's own unit tests (which run against golden fixtures with hand-written minimal
tables, same as `mechanics_test.go` already does today) but **is** a blocker for real
ratings being non-excluded on real encounters — see the Dec 9 analysis below.

**Golden fixtures.** `logs/engine/summary/testdata/{v16,v22}.summary.json.golden` already
exist and are exactly the shape `rating.Score` consumes (`summary.Summary`); the rating
package's own tests add `logs/engine/rating/testdata/*.golden` following the identical
pattern already established at `logs/engine/summary/golden_test.go:19-34` (a `goldenDialect`
table, a `FOREVER_UPDATE_GOLDEN` regen env var, JSON as the portable comparison format) —
this spec's engine lane reuses that harness rather than inventing a second one.

**Encounter mechanics-table coverage, the long pole, checked against
`logs/engine/mechanics/encounters.json` and `logs/engine/mechanics/tables/`:**

| Instance | Type | Encounters | Tables today |
|---|---|---|---|
| Molten Core (map 409) | raid, open at launch | 663–672 (Lucifron → Ragnaros), 10 bosses | **10 / 10** |
| Onyxia's Lair (map 249) | raid, opens 2026-12-09 | 1084 | **1 / 1** |
| Barrow Deeps | raid, opens 2026-12-09 per the roadmap language `encounters.json` itself records | — | **no `DungeonEncounter` ids exist yet** (`encounters.json`'s `no_client_rows_yet.raids`) |
| Hyjal Summit | raid, opens 2026-12-09 | — | **no ids yet**, same note |
| The Hall of Thanes (dungeon, map 3065) | dungeon, opens 2026-11-04 | 3493–3496, 4 bosses | 1 / 4 |
| Ruins of Lordaeron (dungeon, map 2999) | dungeon, opens 2026-11-04 | 7 bosses | 0 / 7 |
| City of Dalaran (dungeon, map 2959) | dungeon, opens 2026-11-04 | 9 bosses | 1 / 9 |
| Excavation Site: Wetlands (dungeon, map 2998) | dungeon, opens 2026-11-04 | 4 bosses | 0 / 4 |

**The genuine December 9 risk is not Molten Core or Onyxia** — both are fully tabled
already, so every raider on the raids most likely to draw the first serious rating traffic
gets Survival/Mechanics/Utility scored from day one. **It is Barrow Deeps and Hyjal
Summit**, which the beta client does not even assign encounter ids to yet (`encounters.json`:
"The beta's level cap is 30, so the higher dungeons and both new raids are not in it yet.
These get tables when a later build or a real log gives their ids"). **This is the same
constraint the guild proposal itself already names** ("the mechanics tables cover a handful
of encounters today and need one per raid boss before December 9: that curation is the long
pole," `2026-09-21-guild-and-premium-proposal.md:154-157`) — this spec adds the precise,
checked accounting of exactly which bosses that means, and turns it into a checklist rather
than restating the general warning.

**Before December 9: a checklist, with who does what.**

| Step | Owner | Status / dependency |
|---|---|---|
| Ship a client build with `DungeonEncounter` rows for Barrow Deeps and Hyjal Summit | the game/client pipeline that produces `data/builds/<build>/` — not a lane of this spec | blocked on a build newer than `1.60.1.69893`, whose level cap (30) does not reach either raid; nothing in this spec can move this |
| Curate Survival/Mechanics/Utility tables for those two raids' bosses, the way `tables/663.json`–`672.json` and `1084.json` already exist for Molten Core and Onyxia | a curator, helped by the extended `mechanics-draft` tool (§3.5) | blocked on the ids above **and** on real raid logs from those encounters — a table cannot be curated from a boss nobody has pulled yet, so this categorically cannot land before players are in the raid, whatever the ids' timing |
| Rate a fight on an encounter with no table using only the table-free components (Output, Preparation, Activity), and say so on the card | the engine lane (already specified, §1.1/§3.5 — nothing new to build here) | **not blocked** — this is the fallback that makes the two rows above non-blocking for ratings *existing* on these raids at all |
| Everything else Molten Core and Onyxia need (engine package + golden tests, API storage/job wired into `ingest.go`, the utility table and consumable catalogue for all 27 specs, downtime windows for Molten Core's ten bosses at minimum, the report-card and character-page web surfaces, the explanation page) | the three lanes in the table above | ordinary lane work, no external dependency |

So the honest state on December 9 is: every raid boss that anyone can have logged before
that date (Molten Core, open since launch, and Onyxia, opening the same day with an id and
table that already exist) is fully rated across all six components from its first kill;
Barrow Deeps and Hyjal Summit run on Output/Preparation/Activity only — excluded-with-reason
on the other three — until a curator can table them from logs that, by definition, cannot
exist before players are standing in those raids. The officer-tooling assignment hook's
actual UI (§2's `Assignment` type ships unused until that spec is written) is out of scope
regardless of date.

---

## 9. Out of scope, recorded

Officer views (raid-night sheet, per-player trend-over-weeks premium depth, wipe analysis,
side-by-side applicant comparison), roster checks, payments/entitlements, the assignments
UI (the data hook exists, §2), and in-game display (the Raider.IO-style snapshot, proposal
§3.5) are not designed here. The rating engine's public function signature (§4.1) and
storage schema (§4.2) are built so each of those can be added as a new reader over existing
data — an officer raid-night sheet is `select ... from rating_scores where report_id = $1`
plus sort, a roster check is the same query parameterised over a guild's `characters`, an
applicant comparison is two character-rating fetches side by side — rather than requiring a
second write path or a schema change.
