# Heating Up: model, Pyroblast lines, Fire guide (2026-10-07)

Lane `heating-up`. Engine branch `heating-up` (sim/mage), site branch `heating-up`. Client build 1.60.1.70009.

## The live text

Talent "Heating Up" (node 105786, one rank, tier 3 column 2, requires Pyroblast; `data/builds/1.60.1.70009/talents/mage.json`):

> Non-periodic critical strikes with Fireball, Frostfire Bolt, Fire Blast, and Scorch reduce the cast time of your next Pyroblast cast within 20 sec by 25%, stacking up to 3 times.

Client rows: 400624 is the talent (a dummy aura named "Hot Streak" in the data); 400625 is the buff: `duration_ms` 20000, effect 0 aura 108 (cast time modifier) misc 10 amount -25, effect 1 a dummy aura of -3 (the stack ceiling). Pyroblast has eight player ranks (11366 ... 18809), all 6000 ms cast, 12 s duration, 3 s ticks.

## The model (`sim/mage/heating_up.go`)

- A permanent trigger aura listens for direct hits (the periodic path is separate, which is the "non-periodic" of the text). A landed crit from Fireball, Frostfire Bolt, Fire Blast or Scorch (by class spell mask) activates the buff, which refreshes its 20 s, and adds a stack, capped at 3.
- Each stack takes 25% off the cast-time multiplier of every registered Pyroblast rank, so 3 stacks cast 6 s as 1.5 s. Applied and removed in `OnStacksChange`.
- Any Pyroblast cast that completes spends all stacks ("your next Pyroblast cast").
- Pyroblast itself does not feed the buff. Frostfire Bolt is in the feeder mask but has no ability in the package, so it is inert until one exists.
- Pyroblast's cast time now reads `PyroblastCastTime` from the generated client table instead of a literal 6 s. Its direct roll already read the client through `clientRoll`; mana cost already read `PyroblastManaCost`; the per-tick dot amount is the existing hand table under `spellconst_damage_test.go`.
- Tests written first (`sim/mage/heating_up_test.go`): cast time from the client table, stack per feeder crit with the 75/50/25% cast times, the cap, non-crits and Pyroblast's own crit giving nothing, the cast spending the buff, the 20 s expiry and refresh. The existing mage package tests and `.results` goldens did not move (no preset takes the talent).

## Candidates measured

Tool: `rotation-search -spec mage-fire -probe-only`, seed 7, 800 iterations, alliance BiS band. Hand-written line: Pyroblast (rank 8, 18809) after the Scorch line, condition `auraNumStacks(400625) >= N`. The guide build does not carry Pyroblast or Heating Up, so a Pyroblast line on that build is inert (the probe's own insertion row says "no effect"); the build therefore has to change with the line. Two points were taken from Master of Elements (2/3 to 0/3); the other sources are in the second table.

| Arm | Build | Rotation | DPS | ± |
|---|---|---|---|---|
| A | guide build (before) | curated | 357.2 | 2.5 |
| B | MoE 2 to 0, Pyroblast + Heating Up | curated, no Pyroblast line | 349.9 | 2.5 |
| C3 | same as B | Pyroblast at >= 3 stacks | 370.0 | 2.8 |
| C2 | same as B | Pyroblast at >= 2 stacks | 358.4 | 2.6 |
| C1 | same as B | Pyroblast at >= 1 stack | 337.9 | 2.2 |

C3 against A: +12.8 DPS (+3.6%), combined error 3.7, so it clears 1% plus the error. C2 is level with A (+1.2, inside error); C1 loses 19. The points alone cost 2.0% (B against A), so the line more than pays for them. Removal probes inside C3: Pyroblast +20.1 (helps), Combustion +8.1, Fire Blast +25.9, all "helps". Seeds 11 and 23 reproduce it (A 357.5 / 357.6, C3 369.7 / 369.6).

Where the two points come from (all with the >= 3 line):

| Points taken from | DPS | ± |
|---|---|---|
| Master of Elements 2 to 0 | 370.0 | 2.8 |
| Wake of Fire 2 to 0 | 370.4 | 2.7 |
| Incineration 3 to 1 | 368.7 | 2.7 |
| Master of Elements 2 to 1 and Wake of Fire 2 to 1 | 363.3 | 2.7 |
| Critical Mass 3 to 1 | 356.6 | 2.5 |
| Improved Fireball 5 to 4, Master of Elements 2 to 1 | 358.4 | 2.5 |
| Ignite 5 to 3 | 356.1 | 2.6 |
| Improved Fireball 5 to 3 | 347.2 | 2.4 |

The first three are equal within error. Master of Elements was adopted: it is fully modeled, while Wake of Fire's kill bonus and Incineration's Scorch crit are only partly or no value here, and the talent adoption rule keeps what the sim cannot value.

## Adopted

- Rotation: one new line, Pyroblast (rank 8) when Heating Up has three stacks, guarded by `auraIsKnown(400625)`. Without the guard a build that has Pyroblast but not Heating Up (ladder level 30) reads the unresolved stack check as always true and spams Pyroblast; the Scorch line uses the same idiom.
- Build: guide FS1 Fire digits `23552100030023051` to `23552100130103051`.
- Bookkeeping: 400625 added to the curated `inert` list; `expected_idle` entry for 18809 (the ladder's level-30 build has Pyroblast but not yet Heating Up); smoke warning for 18809 (a talent spell, absent from the bare smoke build); mage-fire ladder golden regenerated; `make apl-sync` run against the worktree fork (`ENGINE_DIR`), `make apl-check` clean.
- Full search on the adopted list (`rotation-search`, all mutations): no mutation beat the incumbent beyond error. Best variant +0.0 (±3.9). No wrong lines and no missing lines. The earlier held Fire Blast drop is closed for now: with the current engine, removing Fire Blast costs DPS (+14.3 helps on the old build, +25.9 on the new).
- Guide: Talents and Rotation sections now describe Heating Up and the three-stack Pyroblast, the Combustion line reads three crits (it said four), and the stale +8.7% figure is gone.

## Open

- Consumption timing: the buff is spent when a Pyroblast cast completes. A crit that lands while a Pyroblast is mid-cast (a Fireball already in flight) therefore adds a stack that the same completion spends. If the client spends at cast start, that stack survives. The text does not say; revisit when the beta shows it.
- Whether a Pyroblast taken at fewer than three stacks should also spend them is assumed yes ("your next Pyroblast cast").
- Frostfire Bolt feeds the buff in the text but has no ability file, so it feeds nothing here.
- Absolute DPS is far below the 2026-09 search file (416 then, 357 now for the old build): the engine moved since (Combustion at three charges, Ignite no longer double-counting, client-rolled damage). Only paired comparisons within this run are meaningful.
- An accidental `make apl-sync` without `ENGINE_DIR` wrote into the main fork checkout (mage fire and a stale feral APL) before I caught it; both files were restored with `git checkout`. If the feral lane had uncommitted edits to `ui/feral_druid/apls/forever_feral.apl.json` in the main checkout, they were overwritten and need to be redone.
- The site worktree's `sim/go.mod` replace points at the fork worktree and is not committed.
