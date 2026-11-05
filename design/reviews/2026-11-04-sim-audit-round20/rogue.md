# Sim-accuracy round 20: rogue

Lane: fork `lane/r20-rogue` (from int/r20 7704edf50), site `wt/r20-rogue` (from int/r20 1c6357521). One confirmed bug: Forever's poisons scale with attack power and the engine's did not. It moves every rogue spec at every band from 20 up, and it moves the Assassination poison hands at 60.

Measurements: spec-breakdown, 10000 iterations, seeds 1 and 50001, frozen BiS (`$R/bis-frozen`), guide raid builds (`$R/rt`). Before = `$R/sb-before` on `$R/site-before`.

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix |
|---|---|---|---|---|---|
| K1 | Death Mark storage and Mutilate/Shank off-hand on a new log | Launch was 23:00 UTC 4 Nov. Last looked 01:50 UTC 5 Nov: `api.foreversixty.gg/v1/reports/recent` still lists only the 2026-10-20 report; foreverlogs.gg reports 2815, 2830, 2860, 2900 and 3000 answer "Report not found" (nothing past beta #2814). Death Mark and Mutilate are learned at 40, so no launch-night log can carry them anyway | OPEN (needs a level 40+ Forever log with Death Mark / Mutilate) | - | - |
| K4 | Subtlety band 20 Sinister Strike hold on the published mace | Unchanged; the BiS re-rank decides the weapon | Measure at refresh | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix |
|---|---|---|---|---|---|
| X1 | **Instant Poison and Deadly Poison did not scale with attack power.** Blizzard's Forever beta development notes, post of 3 November 2026 (Kaivax, Classes > Rogue): "Deadly Poison Attack Power coefficient increased to 0.009 (Was 0.0045)" and "Instant Poison Attack Power coefficient increased to 0.01 (Was 0.005)". The client cannot show this: 70755 SpellEffect carries 0 in BonusCoefficientFromAP on every Instant (8680...11337) and Deadly (2818...25349) row, as it does for Garrote's server-side 0.03. The engine (poisons.go) dealt the client amount only, so a poison hit at 60 with ~1750 attack power was ~17 short (Instant, 76-100 roll) and a Deadly stack tick ~16 short (client 18). Now each Instant hit is its client roll + 0.01 AP; each Deadly stack ticks for its client amount + 0.009 AP, read at the latest application (the whole dot re-applies, as round 18's Venom fix already treats the multiplier) | Two sources. (1) Blizzard: https://us.forums.blizzard.com/en/wow/t/wow-forever-beta-development-notes-%E2%80%93-updated-november-2/2360696 (read 2026-11-05; copied by Icy Veins, theforeverera.com, patchbot). (2) Beta logs (foreverlogs, 42 rogue entries in reports 2701-2811, levels 29-35, fit in `$R/rogue/fl/poisons.py`): Deadly Poison I single-stack ticks land at 7-8 against the client 6, stacks at 14-16 / 21-24 / 30-32 (not 12/18/24); Instant Poison II non-crits average 24-25 with a max of 28-29 against the client 19-27 roll (centre 23): both above the client by about the old 0.0045 / 0.005 x attack power of a level 30-35 rogue. Round 19 C0 read the 0 on both sides and called it Aligned. Tests `TestDeadlyPoisonTicksScaleWithAttackPower`, `TestInstantPoisonHitsScaleWithAttackPower` (sim/rogue, `-tags with_db`): with the coefficients at 0, five stacks tick for 90.00 (want 168.80) and an Instant hit of 77.44 falls outside 93.33-117.70; both pass with the fix. Full `sim/rogue` and `sim/rogue/dps_rogue` packages pass | CONFIRMED-BUG | old (the 0.0045/0.005 coefficients were live in beta, logs above; round 19 C0 looked at the client row and missed it) + new-data (the 3 Nov doubling) | E e1d7ae036 |
| X2 | **Assassination at 60 carries Instant Poison on the main hand, Deadly on the off hand** (follows X1) | With X1, raid 60 on the guide build: Instant MH / Deadly OH 777.3 / 776.8 (alliance) and 780.7 / 781.4 (horde) against Deadly MH 765.5 / 764.8 and 768.4 / 768.0: +1.5 to +1.7%, both seeds, both factions. Bare 60: 318.6 against 313.4 (+1.7%). Bare 40 and 50 tie (129.1 vs 128.8, 205.4 vs 205.6, 3000 iterations); bare 30 loses (73.2 vs 75.4), where Backstab builds. Combat at raid 60 keeps Deadly MH (798.2, against 792.2 mirrored and 767.1 two Instants); Subtlety's Deadly MH 809.8 ties two Deadlys (810.6) and beats the mirror (793.1), so both keep the existing rule. The off hand alone holds five Deadly stacks, so the main hand's extra procs are worth more as Instant hits once Instant scales | CONFIRMED (consumable kit; any gain beyond error at both seeds) | follows X1 | S 8a9c0565b |
| X3 | Dead talent points in the rogue guide builds | Added the rogue to `TestGuideBuildsSpendNothingOnDeadTalents` (sim/leveling/dead_points_test.go): Improved Expose Armor (Sunder replaces Expose on a boss), Improved Kidney Shot (bosses are stun-immune), Riposte (the boss never attacks the rogue), Setup (no dodges/resists on a boss; 1 Oct notes: only on the dodged/resisted target) and Improved Ambush (every pull opens with Garrote). No rogue build at any band spends a point on them; the test guards it | Aligned (guard added) | - | S fdb0ffc14 |
| X4 | The 3 Nov notes' other rogue lines | Poised (level 18, -10 energy on the next ability at full energy out of combat): `poised.go`. Shank (24, 5 energy, 10 s cooldown, 50% off-hand weapon damage, 15 / 25 energy at 24 / 44, cannot miss): `shank.go`:11-38 against SpellPower, SpellCooldowns and SpellEffect. The talent-tree moves (Murder removed, Lethality 2 points, Venom, Mutilate, Blood Splatter, Coup de Grace, Efficient Strikes, Composure, Flash of Steel, Aggression, Dual Wield Specialization, Setup, Camouflage, Hemorrhage, Quick Getaway, Boot Knife, Death Mark) are all in the 70755 client trees the engine reads; rounds 15-19 audited each | Aligned | - | - |
| X5 | Live logs | See K1 | OPEN | - | - |

## Assassination

| # | finding | evidence and source | confirmed? | origin | fix |
|---|---|---|---|---|---|
| A1 | Poisons (X1) and the poison hands at 60 (X2) | Above. Guide: Poisons, Best in slot, Enchants and consumables and Levels 20 to 60 now say Instant MH / Deadly OH at 60; Stat priority adds that attack power feeds the poisons | CONFIRMED-BUG (X1) | see X1 | E e1d7ae036, S 8a9c0565b |
| A2 | Spell-by-spell fresh pass | Mutilate, Shank, Eviscerate, Rupture, Slice and Dice, Cold Blood, Garrote opener, Seal Fate, Vile/Improved Poisons, Venom, Blood Splatter, Lethality, Malice: round 19 A1 values re-read against the breakdown (unchanged by this round except the poison lines); no new gap found | Aligned | - | - |

## Combat

| # | finding | evidence and source | confirmed? | origin | fix |
|---|---|---|---|---|---|
| C1 | Poisons (X1) | Deadly MH kept (X2). Guide: Stat priority and Enchants and consumables add the attack-power line | CONFIRMED-BUG (X1) | see X1 | E e1d7ae036, S 8a9c0565b |
| C2 | Spell-by-spell fresh pass | Backstab, Sinister Strike fallback, Eviscerate (five, pending owner), Slice and Dice, Blade Flurry, Adrenaline Rush, Shank, Flawless Execution, Efficient Strikes, Coup de Grace, Hack and Slash, Aggression, Dual Wield Specialization: unchanged from round 19 C1/C2; nothing new | Aligned | - | - |

## Subtlety

| # | finding | evidence and source | confirmed? | origin | fix |
|---|---|---|---|---|---|
| S1 | Poisons (X1) | Deadly MH kept (X2). Guide: Stat priority and Enchants and consumables add the attack-power line | CONFIRMED-BUG (X1) | see X1 | E e1d7ae036, S 8a9c0565b |
| S2 | Spell-by-spell fresh pass | Hemorrhage, Ghostly Strike, Garrote, Rupture, Death Mark (ranks by level), Preparation, Shank, Slice and Dice, Eviscerate, Boot Knife, Quick Getaway: unchanged from round 19 S1/S2; Death Mark storage stays K1 | Aligned (storage OPEN) | - | - |

## Before/after

Raid 60, guide raid builds, frozen BiS, 10000 iterations (before -> after, seeds 1 / 50001):

| spec | alliance | horde | change |
|---|---|---|---|
| rogue-assassination | 726.6 / 726.0 -> 777.3 / 776.8 | 729.1 / 728.8 -> 780.7 / 781.4 | +7.0 to +7.2% (poisons +5.3%, Instant MH +1.5%) |
| rogue-combat | 774.0 / 774.2 -> 798.1 / 798.5 | 776.6 / 776.6 -> 800.9 / 800.9 | +3.1% |
| rogue-subtlety | 788.4 / 788.6 -> 809.9 / 810.2 | 782.7 / 782.8 -> 804.3 / 804.3 | +2.7 to +2.8% |

Bare bands, horde, seeds 1 / 50001 agree within 0.1 DPS (seed 1 shown): Assassination 30 70.5 -> 75.2 (+6.7%), 40 119.9 -> 129.0 (+7.6%), 50 193.5 -> 205.3 (+6.1%), 60 300.6 -> 318.6 (+6.0%, with the swap). Combat 30 +4.3%, 40 +3.5%, 50 +3.2%, 60 +2.9%. Subtlety 30 +3.9%, 40 +3.4%, 50 +3.0%, 60 +2.5%. Outputs: `$R/rogue/before-raid.tsv`, `after-raid-final.tsv`, `bands.tsv`, `bare60.tsv`, `hands.tsv`, `hands2.tsv`.

No spec outside rogue moves: the edits to the shared `sim/leveling/kit.go`, `sim/request/raid-consumes.json` and `sim/request/published-presets.json` add rows for `rogue-assassination` only (from level 60). The two JSON files are the committed form of the Go sets; `TestRaidConsumesJSONIsCommitted` and `TestPublishedPresetsJSONIsCommitted` pass.

## Files to regenerate

- Ladder goldens (sim/request `TestRotationLadder`) for the three rogue specs.
- Stat weights, dps-per-point and planner vectors for the rogues (attack power, Strength and Agility now feed the poisons; the guide stat tables render from them).
- addon-data / addon-builds / Data.lua for the rogues (poison hands at 60 for Assassination, stat weights).
- `web` raid-preset (`npm run sync:raid-preset`) picks up the Assassination consumables row.

## Measure at refresh

- Rogue BiS (all three specs, 60 and leveling): attack power weighs more; the re-rank decides.
- K4: Subtlety band 20 weapon.
- X2 at the new BiS: re-check Assassination Instant MH against Deadly MH (expected to hold: it gains with attack power).

## For other lanes

- Every lane: Blizzard's 3 November post states server-side attack-power coefficients the client rows show as 0. Any class whose spells read 0 in BonusCoefficientFromAP should check that post's lines for its class.
- Fork package tests need `-tags with_db` (round 19 X1).

## For the guide lane

- Nothing left over; this lane wrote the poison text on all three pages. `_prose.test.ts` passes (18 tests).

## Rendered review

wow-player on the rendered pages (astro dev on this worktree after `npm run sync`, 1440 px, captures in `$R/rogue/shots2`, `crops2`, `crops3`): Combat SHIP, Subtlety SHIP. Assassination first HOLD (the "while you level ... off Backstab" line was wrong for 40-59, and Best in slot / Enchants named Instant on the main hand without "at 60"); after the rewrite, **SHIP**.

## Sources

- Blizzard, "WoW Forever Beta Development Notes", post of 3 November 2026 (Kaivax): https://us.forums.blizzard.com/en/wow/t/wow-forever-beta-development-notes-%E2%80%93-updated-november-2/2360696 (read 2026-11-05)
- Icy Veins, Development Notes November 2nd: https://www.icy-veins.com/wow-forever/news/class-power-and-enemy-hp-increased-in-massive-class-tuning-wow-forever-development-notes-november-2nd/ ; The Forever Era copy: https://theforeverera.com/en/news/us-30732723/ (search results, 2026-11-05)
- foreverlogs.gg beta reports 2701-2811 (rogue damage timelines cached in `$R/rogue/fl/tl`)

## Still open

- K1: Death Mark storage and the Mutilate/Shank off-hand penalty, until a level 40+ Forever log exists (Mutilate off-hand stays pending with the owner).
- X1 on live logs: a level 20+ rogue with Instant Poison and a known attack power (roster snapshot) would confirm the doubled 0.01 / 0.009 (the beta logs predate the 3 Nov change).
- Carried: Combat Eviscerate at five (owner: keep five).
