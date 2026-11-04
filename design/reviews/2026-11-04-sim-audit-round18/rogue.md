# Sim-accuracy round 18: rogue

Lane `rogue`. Site wt/r18-rogue (from int/r18 4600b7d82, int/r18 b291a4b55 merged in), fork lane/r18-rogue (from int/r18 b2c4a2d05, int/r18 8c24d6202 merged in). Client 1.60.1.70755 with the late hotfix re-layer. Specs: Assassination, Combat, Subtlety.

Method: spec-breakdown on the guide raid builds (`$R/rt`) and frozen BiS (`$R/bis-frozen`). Every reported number is 10000 iterations at seeds 1 and 50001; exploration used 3000 to 5000. The lane's starting engine reproduces `$R/before.tsv` for all twelve rogue rows. Client rows were read from wago 70755 CSVs (SpellEffect, SpellName, SkillLineAbility, TraitDefinition), the hotfix-applied curves in `c755/fetch/.../raw/CurvePoint.csv`, and `spelltips/rogue.json`.

Totals: **3 CONFIRMED-BUGs fixed in the engine, all `old`**: Hemorrhage's 15% never reached the bleeds, Deadly Poison kept its first stack's multiplier, and Thistle Tea was never drunk on a bar above 100 energy. There is also 1 card-text fix from the wow-player review (Subtlety); it changes no rotation action.

Commits: fork e10695cf2 (Deadly Poison), 05ccb0d1d (Hemorrhage), 1aa12229b (Thistle Tea, shared core), 0069e9b9c (APL notes sync). Site 185a31a9e (Subtlety card notes and pull prose), plus this report.

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| K1 | Logs or Forever sources for Death Mark storage and the Mutilate/Shank off-hand penalty | foreverlogs.gg API (`/api/reports/public`, `/api/reports/<id>`, browser headers), 2026-11-04 16:10 UTC: the newest public report is #2811 (2026-10-24, level 30 Stockade). #2812 to #2814 are empty "shield"/"Test" uploads. #2815 and later do not exist. `api.foreversixty.gg/v1/reports/recent` still lists only a 2026-10-20 report. Forever launches 23:00 UTC today, so no level 24+ (Shank) or 40+ (Mutilate, Death Mark) log can exist yet. A web search (Blizzard deep dive thread 2375390, Forever class guides) found no statement on either mechanic. Warcraft Logs is blocked (gcloud expired) and was not worked around | OPEN (no log). Mutilate off-hand stays pending with the owner | - | - |
| K2 | Regression: Backstab before Mutilate/Hemorrhage | Bare 20/30 horde: Assassination casts 31.3 and 33.6 Backstabs and no Sinister Strike. Combat at 20 casts 26.7 Backstabs. Subtlety bare 30 (Toxic Revenger) has Hemorrhage, as intended | Aligned | - | - |
| K3 | Regression: Dirty Deeds in Subtlety | Raid Garrote costs 30 (50 - 20). The opener costs 20 with Poised. 70755 curve 83003 gives -10/-20 | Aligned | - | - |
| K4 | Band-20 Subtlety set's mace (Diamond Hammer 2194): does anything besides Backstab assume a dagger? | Backstab is gated by `spellCanCast`, so with a mace the Sinister Strike fallback fires at 60 energy, not 45. Dropping that hold gains +0.3%/+0.5% on the mace set (41.3 to 41.4, 41.2 to 41.4), the cost round 17 noted. Nothing else at 20 needs a dagger: Garrote, Rupture, Slice and Dice and Premeditation are weapon-blind, and Ghostly Strike/Hemorrhage (not yet known at 20) use the plain weapon share without one. Opportunity 5 in the band-20 build buys only the opening Garrote with a mace. Both disappear when the refresh picks a dagger. No APL value reads the weapon type (round 17 K2) | OPEN (small, modelling limit; fixes itself at refresh) | - | - |

## Assassination

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | **Deadly Poison kept the multiplier of its first stack for the whole fight.** A five-stack Deadly Poison is re-applied by every new proc, and each re-application takes the caster's modifiers of that moment. The engine snapshotted the multiplier only when stack 1 landed (`poisons.go` OnSnapshot, "only the first stack snapshots"). Venom (1310703, 70755: Shank gives +30% poison damage for 6 s, about 60% uptime) therefore stuck to every tick of the fight when the first stack landed under it, and never applied when it did not. Now every application re-takes the multiplier; only a new stack adds base damage | Test `TestDeadlyPoisonReapplicationTakesVenomOfTheMoment` (sim/rogue): before the fix, a re-application under Venom keeps x1.32, want x1.65; it passes after. Raid 60 +0.30% both factions, both seeds (DP alone) | CONFIRMED-BUG | old (upstream 3b0fb2a68 snapshot; mattered once Venom became a 6 s Shank buff in 70755 a57b0221b; no round looked) | E e10695cf2 |
| A2 | Ambush opener instead of Garrote (daggers) | Raid 60, 5000 its: Ambush 723.3 vs Garrote 729.3 | Aligned (Garrote stays) | - | - |
| A3 | Re-audit: Mutilate (90% + 79.2 per hand at r3, 20% vs poisoned, dagger in each hand), Seal Fate 33/66/100 (27.6 CP a fight = P(either hand crits) x casts), Relentless Strikes 4%/rank per point (569.5 energy = 113.7 points x 0.2 x 25), Ruthlessness, Improved Slice and Dice 15%/rank, Vigor 5/10, Lethality 10/20, Vile Poisons 4%/rank, Improved Poisons 2%/rank, Venom, Cold Blood on both Mutilate hands, Blood Splatter 6/9 s, Puncturing Wounds 10/15/5 (Backstab crit, Backstab point, Mutilate crit), Precision e0/e1 (melee and spell hit, which Forever merges), Malice aura 290. Costs: Eviscerate 35, Rupture 25, Slice and Dice 25, Mutilate 60, Shank 5 (net +20 shown). Poisons: Instant 20%, Deadly 30% (spelltips 70755) | 70755 SpellEffect, hotfix-applied curves, spelltips; raid breakdown | Aligned | - | - |
| A4 | Deadly Poison IV at 60, not V | Deadly Poison V (25347 craft) is "Recipe Hidden Until Known" on Wowhead Forever with no trainer, the Era AQ book (Handbook of Deadly Poison V); the engine keeps it behind IncludeAQ=false, the same as the pending level-60 book ranks | Aligned (pending book-rank ruling covers it) | - | - |

## Combat

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | Re-audit: Backstab 60 energy (40.7 casts, crit 63.6% incl. dodges), Eviscerate 20 (Flawless Execution), Shank 8 s cooldown and 40 energy (Flash of Steel 5), Puncturing Wounds +45% point (17.5 a fight = 0.45 x 38.9 hits), Ruthlessness 2 (40%), Composure 2 (opener Garrote 30), Hack and Slash dagger crit 1%/rank, Weapon Expertise 1/2, Blade Flurry 20% speed 15 s 2 min, Adrenaline Rush, Aggression 5%/rank with Coup de Grace 7/13/20 additive | 70755 curves and rows | Aligned | - | - |
| C2 | Ambush opener instead of Garrote | Raid 60, 5000 its: 774.9 vs 777.0 | Aligned | - | - |
| C3 | Gnome (the guide's Alliance pick) gets the Thistle Tea fix X2 | Gnome raid 60: tea 0.67 to 1.00 a fight, +0.56% (773.0 to 777.3, 5000 its) | see X2 | - | - |

## Subtlety

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Hemorrhage's 15% never reached Rupture or Garrote.** Hemorrhage (16511, hotfixed text "take 15% increased Rupture and Garrote damage from the Rogue") is a target debuff read at each tick. `dealBleedTick` scaled the spell's live `DamageMultiplier` around the tick, but a snapshot tick (`Dot.CalcSnapshotDamage`, core spell_result.go:410) reads only `SnapshotAttackerMultiplier`, which was taken at the cast. So every bleed tick ignored the debuff, and so did Death Mark's storage of those ticks. The tick now scales the snapshot multiplier for that tick | Test `TestHemorrhageRaisesRuptureAndGarroteTicks` (sim/rogue): x1.0000, want x1.15 before; passes after. Raid 60 +2.97 to +3.00%, bare 30 to 60 +4.5 to +5.6% | CONFIRMED-BUG | old (c043f79cc 2026-10-07 for Rupture, extended to Garrote by 70755 a57b0221b; rounds 13 and 16 audited Hemorrhage and missed it) | E 05ccb0d1d |
| S2 | Wow-player NO-SHIP items on the Subtlety cards (coordinator's message) | (1) Points held for the next Death Mark: an Eviscerate fallback at five (or four) while the mark is more than 2 s from ready measures 780.3/780.8 (779.0/779.5 at four) against 782.7/782.8, so holding stays. The Death Mark card now says that points Rupture and Slice and Dice do not need wait for the next mark. (2) Hemorrhage's clause is no longer circular. (3) Garrote goes "ahead of Rupture at exactly four". (4) The opener card lists pre-pull casts for every spec (shared layout); the Rupture row and the pull prose ("Then Ghostly Strike for a fifth point, Rupture right away") carry the follow-up. (5) Vanish's icon is the client's real `ability_vanish` (grey smoke with eyes; spellicons maps 1856/1857 to it and the build ships the file), not a placeholder: nothing to regenerate | copy (not counted as a sim bug) | - | S 185a31a9e, E 0069e9b9c |
| S3 | Rupture at three instead of four, after S1 | Raid 60 +0.25% both seeds (784.6/784.9 vs 782.7/782.8), but bare 20 -0.8%, 30 +0.9%, 40 -0.2%, 50 +0.7%, 60 +0.25%. Mixed by band; four stays | Aligned (held) | - | - |
| S4 | Slice and Dice before Rupture at low levels | Bare 20 39.6 vs 41.3, 30 78.2 vs 81.9, 40 127.5 vs 137.3: Rupture first wins everywhere | Aligned | - | - |
| S5 | Re-audit: Ghostly Strike 125/180% (non-normalized, effect 31 alone) and its 7 s Garrote window, Hemorrhage 100/145% normalized (effect 121), Garrote 92 a tick and 30 energy, Initiative 33/67/100, Opportunity 4%/rank on Garrote's periodic, Serrated Blades 3/6/9% armor and 10/20/30% Rupture, Boot Knife 20%/rank for 2 points, Death Mark (25 energy, 10 s, 12% a point, bleeds x2, Coup de Grace on rate and cap), Premeditation without Stealth (3 casts with Preparation), Efficient Strikes -3/-5 (Hemorrhage 30), Camouflage, Elusiveness and Master of Deception as tier gates | 70755 curves and rows; raid breakdown | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Occult Poison I/II learnability (round 17 X1) | Correction to round 17: the CRAFT spells 458822 and 1214168 do have SkillLineAbility rows (51327, 51712; skill line 40, AcquireMethod 0). But so do Season of Discovery's Numbing, Atrophic and Sebacious poisons, so the row does not prove a Forever trainer sells it. Wowhead Forever's spell page could not be read (bot wall) | OPEN (low; in-game check of the poisons trainer at 54/60) | - | - |
| X2 | **Thistle Tea was not drunk on a bar above 100 energy.** The rule was "10 energy or less". A rogue whose bar is above 100 (Vigor, a gnome's Expansive Mind) seldom bottoms out that far: a gnome Assassination rogue drank no tea in a fight, and a gnome Combat rogue drank it in two fights of three. The tea is now drunk when its 100 energy fits under the bar's own cap with at most 10 lost, which is the same rule on a 100-energy bar | Raid 60 gnome Assassination: tea 0 to 1.00 a fight, 711.7 to 727.0 (+2.1%), which puts gnome level with Night Elf (726.8). Test `TestThistleTeaFitsUnderTheBarsOwnCap` (core) failed to build before the helper existed and passes after. **Shared core code** (`core/consumes.go`): only characters with maximum energy above 100 change. Night Elf Assassination (Vigor, 105) gains +0.1 to +0.2% from earlier tea. A 100-energy cat is unchanged | CONFIRMED-BUG | old (d5c560d31, round 10; made worse by round 14's Expansive Mind) | E 1aa12229b |
| X3 | Race re-run (guide-accuracy rounds 13/14 request), raid 60 seed 1, after S1/X2 | Subtlety: NE 788.4, Undead 788.4, Human 786.5, Gnome ~786.7, Orc 785.4, Dwarf 784.5, Troll 782.7. Assassination: Undead 732.6, Orc 729.1, Troll 728.4, NE ~726.8, Gnome ~727.0, Dwarf 725.0, Human 724.6. Combat: Undead 780.0, Orc 777.8, Gnome ~777.3, Troll 776.6, NE 774.0, Human 771.7, Dwarf 770.5. Every guide race line stays inside 1% of the best on its faction | Aligned | - | - |
| X4 | Berserking for Troll rogues (ActionID 26297) | Behaves as the 70755 row 20554 (10%, 15 s, 3 min, no cost); 26297 is only the label | Aligned | - | - |

## Before/after

Raid 60, guide raid builds, frozen BiS, 10000 iterations (relative %, seeds 1 / 50001; absolute at seed 1):

| spec | alliance | horde | seed 1 absolute |
|---|---|---|---|
| rogue-assassination | +0.51 / +0.46 | +0.40 / +0.43 | A 722.9 to 726.6, H 726.2 to 729.1 |
| rogue-combat | 0 / 0 | 0 / 0 | A 774.0, H 776.6 |
| rogue-subtlety | +3.00 / +2.99 | +2.97 / +2.99 | A 765.4 to 788.4, H 760.1 to 782.7 |

Bare bands, horde, published (frozen) sets, seeds 1 / 50001:
- Assassination: 20 0/0, 30 +0.14/+0.14, 40 +0.33/+0.33, 50 +0.57/+0.57, 60 +0.64/+0.67.
- Subtlety: 20 0/0 (no Hemorrhage yet), 30 +4.52/+4.51, 40 +4.59/+4.67, 50 +4.63/+4.63, 60 +5.64/+5.64.
- Combat is unchanged on the published races (Night Elf, Troll; no Vigor). Gnome Combat gains about +0.6%.

DPS outside the rogue class: X2 (fork 1aa12229b) changes only characters with maximum energy above 100. No published druid has that, so no DPS moved outside the rogues in these runs. A gnome druid, if that combination exists, would see it.

## Files to regenerate

- `addon-data.json` (Subtlety notes on Garrote, Rupture, Death Mark and Hemorrhage), then Data.lua. The rendered cards show the old clauses until then. The review below was on a temporary hand-patched copy (same `_one_line` rule), reverted before commit.
- Ladder goldens `sim/request/testdata/ladder/rogue-subtlety.golden.md` and `rogue-assassination.golden.md` (DPS moves: S1, A1).
- `dps-per-point.json` for the three rogues (bleed and poison weights move), then Data.lua.
- Fork goldens: none changed (`sim/rogue/...` and `sim/rogue/dps_rogue` pass unchanged). `sim/druid/tank` TestFeralTank fails on int/r18 with and without this lane's change, and the failure output is identical (pre-existing, see For other lanes).

## Measure at refresh

- Subtlety, every band and both factions: weights and picks after S1 (Garrote, Rupture and Death Mark are worth more, so armor-ignore and AP shares move).
- Assassination weights after A1 (poison share).
- Subtlety band 20: a dagger main hand replaces Diamond Hammer (K4; round 17 S2 already listed it).
- Gnome rogues after X2.

## For other lanes

- Platform/druid: `sim/druid/tank` TestFeralTank fails on fork int/r18 before any rogue change (same diff with and without this lane).
- Anyone with an energy bar: X2 is in `core/consumes.go`; flagged here because it is shared.

## For the guide lane

- Nothing left over. This lane wrote the Subtlety text (card notes, the pull sentence, and "Between marks, every point Rupture and Slice and Dice do not need waits for the next mark"). `_prose.test.ts` and `src/lib/guides` pass (18 and 543 tests).

## Rendered review

wow-player on the rendered Subtlety rotation section (astro dev on this worktree after `npm run sync`, 1440 px, with the hand-patched addon-data): **SHIP**, no MUSTs. SHOULDs: one was applied (the prose now names Slice and Dice and says "Between marks"). Two were left: splitting the Death Mark clause would drop the hold from the card, which shows only the first sentence; and adding Hemorrhage's holds to its card clause would push it past the card length. The prose carries both holds. Vanish reads as a dull grey tile at 16 px but is the real icon.

## Still open

- K1: Death Mark storage and the Mutilate/Shank off-hand penalty, until a level 40+ Forever log exists (Mutilate off-hand stays pending with the owner).
- K4: Subtlety band 20's 60-energy Sinister Strike hold on a mace (-0.3 to -0.5%), until the refresh moves the set to a dagger.
- X1: Occult Poison learnability (in-game check).
- Carried: Combat Eviscerate at four (owner: keep five).
