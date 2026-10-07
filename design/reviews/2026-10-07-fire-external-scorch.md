# Fire mage and the external Improved Scorch debuff, 2026-10-07

Lane `fire-scorch`. Site branch `fire-scorch` (from `main`), fork branch `fire-scorch` (from `forever`).

## Question

The raid preset's `improved_scorch` debuff lowered a level-60 mage-fire by about 4% in the lane's ablation, where a Fire vulnerability debuff should raise it. Engine defect, or something else?

## Reproduction

Level-60 band character (the ranker's alliance gnome: band-60 gear from `data/builds/1.60.1.70009/bis/mage-fire.json`, kit buffs and consumables, the raid preset layered on top), paired seeds 1 to 3, 5000 iterations, default encounter with variation 0. The only difference between the two arms is the `improved_scorch` entry in the buff list. Seeds agree to within 0.02 DPS, so one seed is shown.

Before the fix (Fireball filler, Scorch only to keep the debuff up):

| Band | Fight | Without external | With external | Change |
|---|---|---|---|---|
| 60 | 60 s | 674.05 | 704.87 | +4.6% |
| 60 | 180 s | 630.07 | 624.13 | -0.9% |
| 60 | 300 s | 601.08 | 576.02 | -4.2% |
| 50 | 180 s | 327.31 | 323.14 | -1.3% |
| 50 | 300 s | 290.84 | 287.04 | -1.3% |
| 40 | 180 s | 237.13 | 237.28 | +0.1% |
| 40 | 300 s | 217.75 | 218.55 | +0.3% |

The 300 s row is the lane's -4%. A bare, ungeared level-60 mage with the same talents shows it more strongly (111.2 to 104.3 at 180 s, -6.2%). The absolute DPS differs from the ranker's `set_dps` (324.7 at band 60), which is a different fight setup; the ratio is what matters here.

## Cause

Not the engine. Every candidate in the brief was checked and ruled out:

- **One aura, not two.** `mage.ImprovedScorchAuras.Get(target)` and `core.ImprovedScorchAura(target)` are the same `*Aura` (both go through `GetOrRegisterAura` on label "Improved Scorch"). The APL gate reads that instance.
- **The external debuff never lapses.** `SchedulePeriodicDebuffApplication` sets the shared aura's `Duration` to `NeverExpires` on reset, so five ticks give five stacks and they stay to the end of the fight (log: the aura fades only at the final second).
- **The multiplier is right.** Five stacks raise Fire damage taken by exactly 1.15 (Fire Blast: 500.8 to 575.9 in the debug log). With the Scorch line removed from the rotation, adding the external debuff raised DPS from 516.8 to 624.8.
- **Activate does not reset anything.** A Scorch cast on a full debuff renews the 30 s and adds no stack.
- **Combustion timing** is the same in both arms (popped at the first five-stack moment, 7.5 s without, 3.0 s with).

The cause is the rotation. With Fireball as the filler, the mage's only Scorch casts are the ones that rebuild the debuff each time it lapses. The refresh gate fires at 1.5 s remaining, a 3 s Fireball steps past that window, so the debuff lapses and the mage recasts five Scorches about every 37 s: 26 Scorch casts in 180 s (log: `10207=26.0` against `10151=30.7`). Scorch is the mage's most mana-efficient damage (150 mana for a 1.5 s cast, 207 damage bare, against Fireball's 395 mana for 3 s, 667 bare) and a Fire mage is mana-bound. The external debuff removes those casts and the mage falls back to Fireball and the wand (wand casts 8.8 to 18.2 per fight), which costs more than the 15% it adds.

A fixed-cast-set check confirms it: the same casts under the external debuff are worth +15%.

## Fix

No engine change; the fork gains tests only (commit on branch `fire-scorch`), `sim/mage/improved_scorch_test.go`:

- own and raid Improved Scorch are one aura instance;
- the raid debuff reaches five stacks within 10 s and holds past 100 s;
- five stacks raise Fire damage taken by exactly 15%;
- a sixth Scorch on a full debuff renews the 30 s and adds no stack.

On the site, `data/curated/apl/mage-fire.json` makes Scorch the filler (rank 7, rewritten to the learned rank at lower levels) and keeps Fireball as the line after it for the levels before Scorch is trained (22). The refresh gate stays and still runs ahead of Pyroblast and Fire Blast. Fireball joins `expected_idle` (it is reached only below level 22). The guide's Rotation section is rewritten to match.

Alternatives measured at band 60, 180 s, without the external debuff (the external arm tells the same story):

| Filler | DPS |
|---|---|
| Fireball (before) | 630.07 |
| Fireball above 70% mana, Scorch below | 672.55 |
| Fireball above 40% mana, Scorch below | 657.66 |
| Scorch (shipped) | 696.61 |
| Scorch, refresh gate removed | 695.57 |
| Refresh gate at 3 s, 4.5 s, 6 s (Fireball filler) | 629.6, 629.3, 630.0 |

Widening the gate alone does not help: the external arm stays at 624, because the debuff is not the point, the Scorch casts are.

## After

Same character, seeds and fight lengths:

| Band | Fight | Without external | With external | Change | Without, vs before |
|---|---|---|---|---|---|
| 60 | 60 s | 709.17 | 727.93 | +2.6% | +5.2% |
| 60 | 180 s | 696.61 | 703.50 | +1.0% | +10.6% |
| 60 | 300 s | 693.66 | 698.28 | +0.7% | +15.4% |
| 50 | 180 s | 370.40 | 370.75 | +0.1% | +13.2% |
| 50 | 300 s | 330.67 | 330.65 | 0.0% | +13.7% |
| 40 | 180 s | 254.20 | 254.57 | +0.1% | +7.2% |
| 40 | 300 s | 232.16 | 232.45 | +0.1% | +6.6% |

The external debuff is now neutral to positive everywhere, and the mage's own DPS is 5% to 15% higher than before.

The external debuff gains little at 180 s and beyond because the rotation already holds near-full uptime with its own filler Scorches, which is the right behaviour: a raid mage's debuff is worth only the casts it saves, and a Scorch cast is as good as a Fireball cast.

## Ladder golden

`sim/request/testdata/ladder/mage-fire.golden.md` moved for this reason only (the filler changed). The bare, ungeared ladder character is wand-dominated: level 30 25.5 to 25.2, 38 41.3 to 41.0, 40 42.6 to 41.8 (about 1% to 2% lower), level 50 73.8 to 75.4, level 60 212.6 to 251.1. The geared band characters above are the better measure of the rotation; the small bare-ladder dip at 30 to 40 is the wand covering more of the fight there. No other golden moved.

## Verification

- Fork: `go test --tags=with_db ./sim/mage/... ./sim/core/` passes; `TestP1Mage.results` did not move.
- Site: `go test ./request/ ./cmd/leveling-bis/ ./leveling/` from `sim/` passes; `make apl-check ENGINE_DIR=<fork worktree>` passes.
- `data/tests/test_apl.py::test_every_named_spell_matches_its_own_label_by_name` fails on mage-frost's sidecar, which this change does not touch.

## Follow-ups for the owner

- The nightly ranker will re-pick mage-fire gear and weights against the new rotation; the published `bis/mage-fire.json`, `addon-data.json` and `Data.lua` are intentionally not regenerated here.
- Scorch outscoring Fireball is a model result at these gear levels and mana pools. If real-client Fire guidance disagrees, the Scorch coefficient and Fireball's 3 s cast time in the client tables are the first things to verify.
- Other casters' rotations with a similar "refresh gate also doubles as the best filler" shape (not looked for here).
