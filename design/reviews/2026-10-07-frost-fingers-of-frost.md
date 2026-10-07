# Frost: Fingers of Frost, live Shatter, and what the searches adopted

2026-10-07, client 1.60.1.70009. Engine fork branch `frost-fof` (`sim/mage/fingers_of_frost.go`), site branch `frost-fof`.

## Why

The engine said Frostbite, Shatter and Ice Lance's Frozen bonus all key off a target being Frozen, which only Frost Nova sets, and Frost Nova cannot freeze a raid boss (`canFreeze`: level above the player cap). Against a boss Shatter was inert, Fingers of Frost was not implemented, and the guide build spent 3 points on Shatter for nothing.

## Client rows

| Row | Content |
|---|---|
| Talent 400647 "Fingers of Frost", max rank 2 | Tooltip (checked against the nether.wowhead `tooltip` field): "Gives your Chill effects a 15% chance to grant you the Fingers of Frost effect, which treats your next 1 [2] spell[s] cast as if the target were Frozen. Lasts 15 sec." SpellEffect: effect 1 aura 4 base points 15 (the chance), effect 0 aura 4 base points 2 (the charge count at max rank); SpellAuraOptions ProcChance 100, ProcTypeMask_0 65536 (harmful magic spell hit), no class mask |
| 401741 | The "Gain the Fingers of Frost ability" teaching spell (aura 332 trigger of 400647); generated rank 0, never registered |
| 400669 "Fingers of Frost" | The effect: aura 262, base points 2, EffectSpellClassMask_1 bit 20; duration index 8 = 15000 ms; aura text "Your next $s1 spells treat the target as if it were Frozen" |
| 400670 | Indicator aura: aura 4 base points 1, CumulativeAura 2 |
| Chill class bit | SpellClassMask_0 bit 20 (1048576): Frostbolt (all ranks), Cone of Cold, Frostfire Bolt, the "Chilled" aura spells (what Improved Blizzard applies), Frozen Orb, Spellfrost Bolt. Blizzard itself (524416), Frost Nova (524352) and Ice Lance (655360) do not carry it |
| Shatter 11170, max rank 3 | Effect aura 4, base points 1 for every rank; the per-rank figures are the talent tooltip's: "Increases the critical strike chance of all your spells against Frozen targets by 17% / 33% / 50%" |
| Frostbite 11071 | "Chill effects 5/10/15% to Freeze the target for 5 sec", trigger spell 12494. A root, so a boss is immune; no source in the engine |

## Model

- Chill effects that roll the 15%: a landed Frostbolt or Frostfire Bolt (direct hit only, not the Frostfire dot ticks) and every Improved Blizzard Chilled application (per tick, per target). Cone of Cold takes the same flag when it is registered. Frost Nova is not a Chill effect by the client's class bit, so it lost the engine's `SpellFlagChillSpell`.
- A proc sets the aura to the rank's charges (1 or 2) and refreshes the 15 s; it does not add to a held aura.
- "A spell cast" is a direct-damage cast of the mage's own (`SpellFlagMage|SpellFlagAPL`, `ProcMaskSpellDamage`, not channelled). Channelled spells (Arcane Missiles, Blizzard), non-damaging casts, procs and ticks neither spend a charge nor benefit.
- The cast that resolves while a charge is held is the one treated as Frozen, and it spends the charge as it resolves (TBC and WotLK behaviour). The chill of that same bolt lands at impact, after the charge is gone, so it can re-proc; tests pin both.
- While treated as Frozen: Shatter's crit bonus (17/33/50 crit points) applies to that cast, and Ice Lance's Frozen multiplier applies. Both also apply against a Frost Nova-frozen target, with no charge needed. The two compose without an interaction rule (base-damage multiplier versus crit chance).
- Assumptions the client does not settle: which spells consume (the proc mask has no class mask), the consumption timing, the Improved Blizzard per-tick roll, the x3 reading of Ice Lance's "300% increased" (a x4 reading only makes the Ice Lance line worth more), and Shatter's 17/33/50 coming from the tooltip rather than a spell row.

## Verification

New tests in `sim/mage/fingers_of_frost_test.go`: charges follow rank (1, 2) and last 15 s; absent without the talent; charge spent by the next damaging cast; consumer set (Frostbolt, Ice Lance, Frost Nova yes; Blizzard, Ice Barrier no); Ice Lance on a charge versus none against a boss; Shatter crit gap on a charge and on a Frost Nova-frozen target; the bonus does not leak past the cast; Frostbolt and Frostfire Bolt proc; Frost Nova does not; 15% roll rate; the consuming bolt can re-proc. The conformance report (`mage.golden.md`) is spell-table only and did not move: Fingers of Frost has no spell row to declare, so every mage row stays `declared, matches`. `TestP1Mage.results` moved (reference Frost build has Shatter 3 and Fingers of Frost 2, both now live) and was adopted for that reason.

## Searches (seed 7, 800 iterations, paired, level 60 alliance gnome band)

Guide build before: `2030050001/113023/253511130000030105`, Frostbolt-only rotation.

| Build | Rotation | DPS | vs old guide |
|---|---|---|---|
| Old guide (Shatter 3, no Ice Lance or Fingers of Frost) | Frostbolt | 282.2 ± 1.6 | baseline |
| Search's best clean candidate on the old guide (Shatter -> Frost Channeling) | Frostbolt | 303.4 ± 1.4 | +21.2 ± 2.1 |
| New guide: -Cold Snap, -Improved Frost Nova, -Arcane Blast, +Ice Lance, +Fingers of Frost 2 | Frostbolt | 310.3 ± 2.0 | +28.1 ± 2.6 (+10.0%) |
| Same, plus Ice Lance on a Fingers of Frost charge | Frostbolt + Ice Lance | 319.0 ± 1.9 | +36.8 ± 2.5 (+13.0%) |

Hand variants of the Fingers of Frost spend (same rotation): dropping Shatter for Frost Channeling 3 gave 304.7 ± 1.3, Shatter 2 plus Frost Channeling 2 gave 307.7 ± 1.6, dropping one Shatter for Elemental Precision gave 305.9 ± 1.9; none beat the adopted spend. Fingers of Frost 1/2 in the same spend is not a legal build.

The talent search's generator cannot reach this build: single swaps need a destination with no prerequisite chain, and Ice Lance (a prerequisite with no credit under a Frostbolt-only rotation) is only picked up by the structural builds that drop unmodeled guide talents. The build was composed by hand from the free points (Cold Snap, Improved Frost Nova and Arcane Blast are modeled with no damage and gate nothing the build keeps; Wand Specialization and the Fire tier-0 points hold tier gates for kept talents) and verified legal.

Rotation search on the new guide: Ice Lance inserted with the aura gate (`auraIsActive` 400669) is a missing line beyond error: +9.0 ± 2.7 probe, +8.7 ± 2.7 (+2.8%) as the final rotation, one accepted mutation (Ice Lance at position 2). The search tool gained a charge gate for Ice Lance (`chargeGatedSpells` in `sim/cmd/rotation-search/spells.go`), and its dangling-gate check no longer flags a gate on a talent-granted aura.

Talent search re-run on the new guide (rotation now with Ice Lance): guide 319.0 ± 1.9; the best variant that keeps every unmodeled guide talent is Improved Fireball -> Incineration at +1.1 ± 2.6, within error, so nothing further is adopted. Winners that beat it by more (a deep Frost build at +49.1) drop Frostbite, Permafrost, Frost Warding, Impact and Flame Throwing, which the engine does not model, and are not adopted.

## Adopted

- Guide build `FS1:1.60.1.70009:mage:gnome:203005/113023/253510130100030025:` (passes the rule: keeps all unmodeled guide talents, holds the Frost tree, +10.0% beyond ±2.6 error, drops nothing the rotation casts).
- `data/curated/apl/mage-frost.json`: Ice Lance (rank 6) while the Fingers of Frost aura is active, ahead of Frostbolt; mirrored to the fork's `forever_frost.apl.json` by `make apl-sync`.
- Guide prose: Fingers of Frost and Shatter described, build paragraph and rotation rewritten with relative gains only.
- Ladder golden regenerated: the build change moves the level 30 to 60 rungs. Rungs below the Fingers of Frost row list 400669 as unresolved, expected until the build reaches the talent.

## Open

- Frostbite's Freeze has no source on a boss; on a freezable target it would add a third Frozen source.
- Cone of Cold is not registered; it should carry `SpellFlagChillSpell` when it is.
- Beta data on the proc rules (consuming cast, per-tick Blizzard rolls) would settle the assumptions above.
