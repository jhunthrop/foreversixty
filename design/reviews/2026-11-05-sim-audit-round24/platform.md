# Sim-accuracy round 24: platform lane

Site branch `wt/r24-platform` (from int/r24 f05c281fe). Fork branch `lane/r24-platform` (from int/r24 3b4d6b054). The lane owns sim/core, sim/leveling, the web and the addon data. Client 1.60.1.70790. BUILD MODE: touched tests only, no regens, no wow-player gate.

Method: the four seed items in the brief's order, then a fresh re-audit of round 23's shared changes (manaIncomePerSecond, number x duration, the empty-tree parser, talent-gated card rows). Every sim went through `$R/sim.sh`. Every CONFIRMED-BUG has a test that was red before its fix.

**Confirmed and fixed: 5** (L2 `old`, C1 `R22 6c639ea7f`, C2 `R22 6c639ea7f`, A1 `R23 7bb3c301a`, A2 `old`). Plus one feature (M1, the spend value). **Cross-class DPS movers: none.** No raid or band checkpoint number moves. L2 moves Fire's in-between leveling levels 41 to 44 only (mage, measured below). C1 and C2 move card row start levels for Enhancement and the three mage specs (card data, no DPS).

## Live logs

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| L1 | **Still no live Forever logs we can read.** api.foreversixty.gg `/v1/reports/recent` lists only `z2opt2x2gs5u` (20 Oct). foreverlogs.gg `/api/reports/public` answers 403. The Warcraft Logs key is with the owner (401 last round). | last looked 14:29 UTC 5 Nov | OPEN (no data) | - | - |

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| M1 | **The mana-spent APL value is built.** `APLValue.mana_spent_per_second = 80 { double window_seconds = 1 }` (0 means 10 s): every cost paid through `SpendMana` in the trailing window, over the window, or over the time elapsed early in the fight. The same windowing and reset as `manaIncomePerSecond`: both now share one flow log (`sim/core/mana_flow.go`, renamed from `mana_income.go`), which records nothing until a rotation reads it. Site: `apl-gates` knows the node; `apl-condition` words `manaIncomePerSecond >= manaSpentPerSecond` as "what you regain keeps up with what you spend". Usage note for the paladin lane: `$R/platform/for-paladin.md`. | fork `proto/apl.proto`, `sim/core/proto/apl.pb.go` (protoc 36.2, protoc-gen-go v1.36.6), `mana_flow.go`, `mana.go`, `apl_value.go`, `apl_values_resources.go`; test `TestValueManaSpentPerSecondReadsTheTrailingWindow` (red: undefined); site `sim/request/mana_income_test.go` `TestALineGatedOnManaSpentReadsTheCharactersSpend` (red on the pinned engine: unknown field "manaSpentPerSecond"; green on the lane: a 1 mana/s floor casts, a 100,000 mana/s floor never does, income >= spend parses, on a real priest); vitest `reads income against spend` (red, then green) | Feature (not a bug) | - | E 568d4a7fe, S afe672388 |
| L2 | **The leveling order could not put an off-tree point between two own-tree points.** `reserveOffTree` spent every off-tree reserved rank before the spec's own tree. Shadow's Wand Specialization is meant to come first, so it was right. Fire's could only come first or after all 34 Fire points, so it landed at 44 and 45. Now a reservation is spent in its own list order across trees: an off-tree rank waits until the budget has paid for the own-tree ranks listed before it (`ownRanksAhead`). Fire's new ladder-order row takes Wand Specialization 2 right after Combustion, at 41 and 42. Level 45 holds the same build either way, so no checkpoint moves (`TestLadderOrderMovesNoCheckpoint`). The web's port of the walk (`leveling-builds.ts`) got the same rule. | `sim/leveling/talents.go`, `data/curated/ladder-order.json` (+ embedded copy), `web/src/lib/guides/leveling-builds.ts`; tests `TestReservationSpendsAnOffTreeRankInListOrder` (red: Wand Specialization 1 at 14 ahead of Incineration 5), `TestMageLadderTakesItsCastTalentsEarly` 41/42/43 (red), vitest `takes Fire's Wand Specialization at 41 and 42` (red); all of `go test ./sim/leveling` and `leveling-builds.test.ts` pass. **Addon order dry run** (addon-builds run, diffed, reverted): only `mage-fire leveling` moves, points 41 to 45; every other spec's leveling and raid order is unchanged, Shadow's included. | CONFIRMED-BUG (leveling order, Fire) | old (the reservation code, round 10/15; round 23 mage measured the cost) | S 502efcb28, 0b547981f |
| C1 | **Enhancement's Fire Nova row showed from 12.** The line is gated on Clearcasting (16246), the aura Elemental Focus (16164) gives, not on the talent's own spell, so round 22's talent-gated row start never found it. `data/curated/talent-owned-auras.json` is now the one aura -> talent table: `talent_levels.load_spec_point_levels` gives an owned aura its talent's level, and `sim/request`'s ladder harness reads the same file in place of its hard-coded mage map. Fire Nova starts at 52, where the raid order takes Elemental Focus (the leveling build never takes it). | `sim/shaman/talents.go:116-125` (aura registered only with the talent); `data/pipeline/talent_levels.py`, `addonrotation.py`, `sim/request/ladder.go`; tests `test_a_row_gated_on_a_talent_starts_where_the_published_build_takes_it[...Fire Nova...-52]` (red: 12), `TestAddTalentOwnedAurasExemptsClearcasting` | CONFIRMED-BUG (card data) | R22 6c639ea7f (the row-start rule never read aura gates) | S 33fb1b5f0 |
| C2 | **The same gap on the mage cards.** A scan of every curated APL for `auraIsKnown` gates on an id that is not a talent spell found three more talent auras: the Missile Barrage buff (400589, talent 400588), Heating Up (400625, talent 400624) and the Improved Scorch debuff (12873, talent 11095). Rows moved: Fire's Heating Up Pyroblast 20 -> 30, Fire's Scorch filler 22 -> 30, Improved Scorch stacking 22 -> 27, Missile Barrage Missiles 10 -> 55 (Fire), 10 -> 51 (Frost), 10 -> 22 (Arcane), Arcane's Heating Up Pyroblast 54 -> 57 and its raid-build Fire Blast, Fireball and Frostbolt rows 10 -> 57. Every new start is where the published order takes the talent. Other non-talent aura gates in the scan are `auraIsActive` only (debuffs, buffs you cast, Windfury Totem) and do not floor a row. | `sim/mage/missile_barrage.go:46`, `heating_up.go:37`, `talents.go:382`; five new parametrized cases (red without the table rows: 5 failed) | CONFIRMED-BUG (card data) | R22 6c639ea7f | S 0f29d428e |
| X1 | **Cost reductions: the engine combines them the way the server does.** Per-spell percent mods (aura 108 misc 14: Arcane Power, Missile Barrage, Clearcasting, Devouring Contagion, Mental Agility) are summed into one multiplier; the school multiplier (aura 72: Shadowform) multiplies after them; flat mods come first. vmangos `Spell::CalculatePowerCost` does exactly this: flat school mod, then `ApplySpellMod(SPELLMOD_COST)` (which sums `totalpct` and `totalflat`), then `powerCost * (1 + UNIT_FIELD_POWER_COST_MULTIPLIER[school])`. Round 20 mage S2 already showed from the client rows that aura 108 adds (only the mage's Clearcasting is -1000). Priest S6's 985 x (1 - 0.5 - 0.1) x 0.5 = 197 is the server's number. One small difference: the server truncates to whole mana, the engine keeps fractions (under 1 mana a cast). | `sim/core/spell.go:720-724`, `spell_mod.go:577-599`, `spell_school.go:165-170`; https://github.com/vmangos/core/blob/development/src/game/Spells/Spell.cpp (`CalculatePowerCost`), https://github.com/vmangos/core/blob/development/src/game/Objects/Player.cpp (`ApplySpellMod`), read 5 Nov | Aligned (mage and priest) | - | - |

## Fresh re-audit of round 23's shared changes

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | **A duration constant in number x duration math read 0.** Round 23's M2 coerces both operands to float. A const like "30s" is parsed as a duration and kept no float value, and coercing a const copies its fields, so `level x "30s"` gave 0. A duration const now keeps its seconds as its float value. Duration ranks above float in comparisons, so no comparison ever coerced a duration const to float; only M2's math and the energy-threshold scan (which never sees a duration) read it. No curated APL multiplies by a duration const. | `sim/core/apl_values_operators.go:40-46`; test `TestMathOfANumberAndADurationConstantWorksInSeconds` (red: 0, want 60) | CONFIRMED-BUG (engine, latent) | R23 7bb3c301a | E 568d4a7fe |
| A2 | **The web's build-code reader still needed a digit in every tree.** Round 23 fixed sim/leveling's FS1 regexes (E1) but not `buildCodeDigits` in `web/src/lib/guides/leveling-builds.ts`, which the Leveling strip, the talents by level tree and `/data/leveling-plans.json` use. It reads only `build:` lines, and none ends in an empty tree today, so nothing renders differently; a future one would show no leveling tree while the sim and the addon read it. The re-encoder (`web/scripts/reencode-guide-builds.mjs`) and the planner's FS1 parser already accept an empty tree. | test `reads a build code by tree` (red: undefined for Assassination's raid line) | CONFIRMED-BUG (web, latent) | old | S 73b075b10 |
| A3 | **manaIncomePerSecond (E 61dc24ce5).** Reset each iteration, pruned to the longest window, net of overflow. Early in the fight it divides by the elapsed time, so the first regen tick reads high for a second or two, and any gain before the pull (negative time) is counted against the elapsed time after it. Retribution's budget only uses it on a bar that is not full, where this lag is what round 23 measured. The spend value inherits the same shape; the paladin note says so. | `sim/core/mana_flow.go` | Aligned (documented behavior) | - | - |
| A4 | **Talent-gated card rows (S 6c639ea7f, round 23 C-rows).** Bare `spellIsKnown`/`auraIsKnown` and top-level `and` terms floor a row; `or` and `not` do not, by design. The gap was aura gates (C1, C2). | `data/pipeline/addonrotation.py:275-301` | Aligned after C1/C2 | - | - |
| A5 | **Empty-tree parser (S d8c453fdf).** sim/leveling and the re-encoder agree; see A2 for the web. | - | Aligned after A2 | - | - |

## Before/after

No published number moves. Raid 60, horde, seed 1, 3,000 iterations, frozen BiS, guide raid build, `$R/sb-before` against `$R/platform/sb`: paladin-retribution 779.8 / 779.8, mage-fire 711.0 / 711.0, priest-shadow 718.2 / 718.2, shaman-enhancement 838.3 / 838.3.

Fire, bare, band 40, each level's ladder build passed with `-talents` (Wand Specialization last, then right after Combustion), 10,000 iterations (`$R/platform/fire-wand.tsv`):

| level | horde s1 | alliance s1 | horde s50001 | alliance s50001 |
|---|---|---|---|---|
| 41 | 160.6 -> 161.1 (+0.31%) | 164.7 -> 165.1 (+0.24%) | +0.31% | +0.24% |
| 42 | 162.8 -> 163.8 (+0.61%) | 166.5 -> 167.4 (+0.54%) | +0.67% | +0.54% |
| 43 | 164.1 -> 165.1 (+0.61%) | 167.7 -> 168.6 (+0.54%) | +0.55% | +0.48% |
| 44 | 166.0 -> 166.4 (+0.24%) | 169.3 -> 169.7 (+0.24%) | +0.24% | +0.24% |

Error is +/- 0.2 DPS. 45 and every checkpoint are the same build.

## Files to regenerate (integrator, at the pin)

- `python -m pipeline simproto --engine <fork at the pin>` (data/proto/apl.proto, `pipeline/simproto/apl_pb2.py`) for `manaSpentPerSecond`. Until then `data/tests/test_apl.py` rejects a curated APL that uses it.
- `go run ./sim/cmd/addon-builds -repo-root .` (mage-fire leveling, points 41 to 45), then the addon data step (`addon-data.json`, `Data.lua`: also the C1/C2 row starts) and `make simdb`. `TestCommittedAddonBuildsFollowTheGuides` fails until addon-builds is regenerated.
- Fork `ui/core/proto` TS is not regenerated (as with the last two values).

## Measure at refresh

- Nothing new. The BiS re-rank and dps-per-point read checkpoint builds, which do not move.

## For other lanes

- **Paladin:** `$R/platform/for-paladin.md`. Until the integrator merges, cherry-pick E 568d4a7fe onto your engine lane (or point a go.work at `$R/platform/engine`). The card words only `manaIncomePerSecond >= manaSpentPerSecond` (OpGe, that order); tell the integrator if you adopt another shape so a reader is added. OR the gate with a full-bar clause: income is 0 at a full bar.
- **Mage:** Fire's Wand Specialization now comes at 41 and 42 (ladder-order row); the round-23 "Still open" item is done. Card rows for Missile Barrage, Heating Up and Improved Scorch now start where the build takes the talent (C2): check your card notes still read right from those levels. Cost stacking (X1) is aligned.
- **Priest:** S6 cost stacking is the server's (X1).
- **Shaman:** Fire Nova's row starts at 52 (C1), where the raid order takes Elemental Focus.
- **Any lane:** a gate on an aura a talent gives needs a row in `data/curated/talent-owned-auras.json` (aura, talent spell, the engine line that ties them) for its card row to start at the talent.

## For the guide lane

- None.

## Still open

- L1 (logs).
- Whole-mana truncation of costs (X1): the engine keeps fractions of a mana point. Under 1 mana a cast; not worth moving every caster's numbers for.
