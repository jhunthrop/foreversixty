# Sim-accuracy round 16: rogue

Lane `rogue`. Site wt/r16-rogue (from int/r16 c62d77ab8), fork lane/r16-rogue (from int/r16 9f7f1f6e4). Client 1.60.1.70755 with the late hotfix re-layer. Specs: Assassination, Combat, Subtlety.

Method: spec-breakdown, raid 60, guide raid builds (`$R/rt`), frozen BiS (`$R/bis-frozen`), 10000 iterations at seeds 1 and 50001 for every reported number; 3000 iterations to explore. The lane's starting point reproduces `$R/before.tsv` exactly for all twelve rogue rows.

Fork commits: bf99ed477, 1470d4ba9, d69ca607a (APL copy syncs). Site commits: f89d5bc98 (rotation and tests), 8160ad885 (guide text and this report), and a sha fill-in.

Totals: 3 CONFIRMED-BUGs fixed (2 rotation, 1 guide text), all Subtlety or text. Origins: 1 `R15 493a3d038`, 1 `70755 d009d90a3`, 1 `R15` text hand-off (round 15's guide items never landed). No engine code changed; the fork commits are APL copy syncs.

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| K1 | Does Death Mark store white swings and poisons? | No source says either way. The 70755 rank text and the Wowhead Forever tooltip (https://nether.wowhead.com/forever/tooltip/spell/1325758, read 2026-11-04) say "storing a portion of the damage dealt by your attacks", Garrote and Rupture at twice the rate. Searched 2026-11-04: Blizzard deep dive (https://us.forums.blizzard.com/en/wow/t/2375390), wow.gg Subtlety overview (https://wow.gg/guides/rogue-subtlety-forever-overview, 2026-09-19: no Death Mark text), wowtbc.gg Forever Subtlety (https://wowtbc.gg/warcraftforever/class-guides/subtlety-rogue/, undated: "Cast Death Mark at 5 Combo Points", nothing on what it stores), mobalytics (403). Logs: foreverlogs.gg has no report after #2813 (2026-10-29) and every beta log is level 13 to 33; Death Mark is learned at 40. 97 rogue encounter timelines from 18 characters were pulled (the highest-geared beta rogues): no Death Mark, no Mutilate, no Shank. The engine keeps storing every hit except proc-suppressed ones (Shank). A level 40+ log of a mark's payout against its cap settles it | OPEN | - | - |
| K2 | Mutilate and Shank off-hand penalty on live logs | No log carries either spell (K1's pull: Shank arrived with 70753/70755, Mutilate is learned at 40). The engine model is the server's: the off-hand damage stat is halved for every OFF_ATTACK strike (TrinityCore's UNIT_MOD_DAMAGE_OFFHAND TOTAL_PCT 0.5), and both spells carry REQUIRES_OFF_HAND_WEAPON (round 15 W1). So Shank is 50% of the halved off-hand roll (`core.OHWeaponDamage`). One thing to know when a log appears: Shank at bare 30 averages 14 damage a hit; if a log shows about twice that, the off-hand halving does not apply to Shank | OPEN (no log exists) | - | - |
| K3 | Regression: Shank ranks below 44 | Bare 30, all three specs cast Shank 1329283 on cooldown (18.2 casts a fight, 17.7 for Subtlety with its gate) | Aligned | - | - |
| K4 | Regression: Subtlety Eviscerate/Death Mark gate, Shank ≤3 and Ghostly Strike ≤4, Preparation | Eviscerate is idle once Death Mark is known (raid 60: 0 casts); Shank 17.8 and Ghostly Strike 9.6 casts a fight, combo-point gated; Preparation 1 cast, Premeditation 3 | Aligned | - | - |
| K5 | Regression: Subtlety's Garrote after Ghostly Strike | **The line rarely fired: see S1** | CONFIRMED-BUG | R15 493a3d038 | S f89d5bc98 |
| K6 | Regression: Assassination's Rupture with Blood Splatter; the Mutilate off-hand penalty | Rupture 7.2 casts a fight at 4 and 5 points; variants: Rupture at five only -0.6%, refreshing under 2 s -0.1%, Shank below Mutilate -0.1%. Main-hand Mutilate 620 average, off hand 402 (the off hand halved, flat whole) | Aligned | - | - |
| K7 | Regression: Combat's Composure | Garrote opener costs 30 (50, Poised -10, Composure -10), as 1329265's curve (-5/-10) says | Aligned | - | - |
| K8 | Regression: Cold Blood order | Shank, Slice and Dice, Rupture, Cold Blood, Eviscerate, Mutilate; 1.43 Cold Bloods a fight | Aligned | - | - |

## Subtlety

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Garrote after Ghostly Strike was starved by Hemorrhage.** Ghostly Strike's 1329326 lets the next Garrote in without Stealth for 7 seconds, and round 15 put a Garrote line right under Ghostly Strike. But a priority list moves on when a line cannot pay: Hemorrhage (30 energy) fired every time Garrote's 40 was not there yet, so energy never reached 40 inside the window. Garrote went out 1.9 times a fight (the opener and one more) against 9.6 Ghostly Strikes, with 27% uptime. Now Hemorrhage holds while 1329326 is up, and Garrote goes on after every Ghostly Strike at four points or fewer, refreshing a Garrote that still ticks | Raid 60 horde, 3000 iterations: 717.4 base; Garrote with no combo gate 718.7; Hemorrhage held, Garrote only when dropped 738.6 (4.7 casts); refresh under 8 s 743.2, under 10 s 744.5, under 14 s 744.6, always 744.5 (8.7 casts). The plateau is taken in its simplest form (always). Garrote at five points as well loses (727.8). Test `TestSubtletyGarrotesAfterEveryGhostlyStrike` (sim/request/rogue_rotation_test.go): 1.07 Garrotes against 8.74 Ghostly Strikes before, passes after. Client: 14278 e5 triggers 1329326 (aura 275 on mask 768, Garrote 256 + Ambush 512, 7 s) | CONFIRMED-BUG | R15 493a3d038 | S f89d5bc98, E bf99ed477 |
| S2 | **Death Mark waited for a fifth point.** It comes back every 10 seconds, so holding it for a fifth combo point delays every mark after it. Marking at four or more | On top of S1, horde: raid 60 744.6/744.4 to 746.0/746.0 (+0.19%, seeds 1/50001, ±0.3); bare 40 133.3 to 137.2 (+2.9%), 50 201.8 to 203.4 (+0.8%), 60 345.2 to 345.5 (+0.1%). Rupture at five instead of four: -0.7%. Test `TestSubtletyDeathMarksAtFourComboPoints` (red on the five-point line, green after) | CONFIRMED-BUG (rotation) | 70755 d009d90a3 (the line came in at five with Death Mark; round 15 kept it) | S f89d5bc98, E 1470d4ba9 |
| S3 | Hemorrhage's 15% debuff reaches Garrote as well as Rupture? | The shipped client row 693309 has EffectSpellClassMask_0 1048576 (Rupture only), but the server push 113483 (wago.tools/hotfixes, 2026-11-03, build 70755) sets it to 1048832, Rupture 1048576 + Garrote 256. The engine's debuff covers both, and the tooltip says "Rupture and Garrote". Note for every lane: wago's CSV export does not carry the push | Aligned | - | - |
| S4 | Hemorrhage's 145% "if a Dagger is equipped" (SIM-R5-1) | 16511 e4 is a server dummy (145) with no hand named; Ghostly Strike's text says "in your Main Hand", Hemorrhage's does not. The published Subtlety set is two swords (Teebu's Blazing Longsword, Persuader), so the published number is untouched either way | OPEN (in-game check of a sword main hand with a dagger off hand) | - | - |
| S5 | Re-audit, spell by spell against spellconst 70755 and the Wowhead Forever tooltips: Rupture 6 (35 + 4.73 a point a tick, 3 + points ticks, Blood Splatter +3 ticks, AP 4/10/18/21/24% over the ticks), Garrote 6 (92 every 3 s, 18 s, +9 s, 552 on the tooltip), Death Mark ranks and caps (12% a point, bleeds ×2, cap 156 + 245 a point + 4% AP a point, × crit, × Coup de Grace), Ghostly Strike (40 energy, 20 s, 125/180% main-hand dagger, Lethality), Hemorrhage (35 - 5, 15 s), Shank and Boot Knife (2 points, round 15 S2), Premeditation (2 points, 20 s, 2 min, no Stealth), Preparation (10 min), Slice and Dice 2 (30% haste, 9 + 3 s a point), Serrated Blades, Initiative, Dirty Deeds (-10 Garrote), Opportunity, Camouflage, Elusiveness (no damage) | client rows; engine `sim/rogue/*.go` | Aligned | - | - |

## Assassination

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | Rotation re-checked (K6, K8) and the raid build point by point: Malice (+5% crit, every table, poisons included), Ruthlessness, Improved Slice and Dice 3, Vigor 1, Blood Splatter, Lethality, Vile Poisons, Cold Blood, Improved Poisons (the charge half is not modelled; charges are not), Relentless Strikes, Venom, Seal Fate, Mutilate, Coup de Grace (Eviscerate), Lightning Reflexes 2 (gate), Puncturing Wounds (Mutilate +15% crit), Precision, Opportunity. Every point does something or is a gate | talents/rogue.json 70755 rank texts; engine | Aligned | - | - |
| A2 | Gnome Assassination's Eureka! (SIM-R3-1) | The rogue row 1259812 covers Mutilate's cast (1241584, cost -10%), Eviscerate, Backstab, Sinister Strike, Ghostly Strike and Hemorrhage, not Mutilate's hand strikes 1241586/1241590 (their class flags are outside the masks). Expansive Mind re-enables the energy bar after Vigor, so it multiplies 105 to 110.25. A Gnome below the other races is the race's base stats, not a lost racial | Aligned | - | - |
| A3 | Prose: "Rupture and Expose Armor do not make the list" contradicted the rotation since round 15 | See X1 | CONFIRMED (text) | R15 | S 8160ad885 |

## Combat

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | Openers (SIM-R2-2): Garrote vs Ambush vs Backstab from Stealth | Raid 60 horde, seeds 1/50001: Garrote 776.6/776.6, Ambush 774.7/774.7 (-0.25%), Backstab 773.8/773.3 (-0.4%). The guide's Garrote opener stands, and combat.md:71 makes no ranking claim between the other two | Aligned | - | - |
| C2 | Crit's low weight (SIM-R6-1) | White swings: 19.0% miss, 4.4% dodge, 40.0% glancing, so the white crit cap is 36.6%, and the white crit chance is 36.5%: the swings sit at the cap. Backstab crits 63.6% (Puncturing Wounds, Hack and Slash dagger crit, Malice). The weight is the cap, not a talent gap | Aligned | - | - |
| C3 | Eviscerate at four instead of five, and Slice and Dice at one to three | Raid 60 horde: 776.6 to 781.7/781.9 (+0.66/+0.68%, seeds 1/50001). But on the bare sets (the BiS band talents) it loses: 30 -1.4%, 40 -1.8%, 50 -1.1%, 60 -0.3%; Eviscerate at four alone: raid +0.5%, bare 30 -0.4%, 40 -0.2%, 50 +0.1%, 60 +0.7%. The raid gain comes from Backstab's second combo point (Puncturing Wounds) on daggers; the bands that lose build with Sinister Strike. The priority list has no main-hand weapon test to split them | OPEN (needs a dagger condition in the rotation language, or the owner's call raid vs leveling) | - | - |
| C4 | Adrenaline Rush moved above Backstab | -0.8%: kept where it is | Aligned | - | - |
| C5 | Thistle Tea inside Adrenaline Rush (SIM-R3-2) | The engine drinks at 10 energy or less, inside Adrenaline Rush or not. combat.md:83 says only "drink it when you are near empty", which makes no claim about Adrenaline Rush | Aligned (no claim to test) | - | - |
| C6 | Combat build points, Blade Flurry (13877, 25 energy, 2 min, +20% speed, mask cleave only with a second target), Adrenaline Rush (5 min, 15 s, +100% regen), Composure, Flawless Execution, Weapon Expertise, Aggression (Eviscerate), Dual Wield Specialization | client rows; engine `talents.go`, `poised.go` | Aligned | - | - |
| C7 | Prose: combat.md:52 and :63 still described Efficient Strikes 2/2 in the raid build (round 15 C1 moved those points to Composure) | See X1 | CONFIRMED (text) | R15 6e178531b | S 8160ad885 |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **Round 15's rogue guide text never landed.** Its "For the guide lane" items 1 to 7 were still open on int/r16: combat.md (Efficient Strikes in the raid build), assassination.md (no Rupture), subtlety.md (Eviscerate while Death Mark is on cooldown, nothing on Garrote after Ghostly Strike), and the Subtlety Shank, Ghostly Strike and Eviscerate notes without their gates. Written in this lane with S1/S2's prose, in guide voice; "bleeds always crit" became "bleeds can crit in Forever" in the prose and the Assassination Rupture note (the owner's rule means they are able to crit, not that every tick crits) | round 15 rogue report, "For the guide lane"; integration round 15 "Subtlety Shank note" | CONFIRMED (text) | R15 (hand-off not landed); guides-handoff | S f89d5bc98 (notes), S 8160ad885 |
| X2 | Hit is one stat for melee and poisons (Precision and gear hit lower a poison's 17% miss) | Shared core decision (`client_sets_test.go` "hit is one stat"), not rogue-owned | For other lanes (recorded only) | - | - |

## Before/after

Raid 60, guide raid builds, frozen BiS (relative %, seeds 1 / 50001; absolute at seed 1):

| spec | alliance | horde | seed 1 absolute |
|---|---|---|---|
| rogue-assassination | 0 / 0 | 0 / 0 | A 722.9, H 726.2 (note text only) |
| rogue-combat | 0 / 0 | 0 / 0 | A 774.0, H 776.6 |
| rogue-subtlety | +4.01 / +3.93 | +4.07 / +4.02 | A 722.3 to 751.3, H 716.8 to 746.0 |

Of Subtlety's +4.1% (horde), S1 is +3.9% and S2 +0.2%.

Bare sets, horde, seed 1, 3000 iterations (the BiS band talents), Subtlety: 30 72.2 to 81.9 (+13.4%), 40 120.9 to 137.2 (+13.5%), 50 182.2 to 203.4 (+11.6%), 60 314.7 to 345.5 (+9.8%).

No commit moves DPS for any spec outside the rogue class.

## Files to regenerate

- Ladder golden `sim/request/testdata/ladder/rogue-subtlety.golden.md` (Garrote after every Ghostly Strike from 26, Hemorrhage hold, Death Mark at four from 40; aura 1329326 is unresolved on rungs below Ghostly Strike, as the smoke test now expects).
- `addon-data.json` (the Subtlety rotation lines and notes: the card and the addon), then Data.lua.
- `dps-per-point.json` for Subtlety, then Data.lua.
- apl-sync: the Subtlety and Assassination copies are already hand-synced on both sides (site `sim/request/apl/`, fork `ui/rogue/apls/forever_*.apl.json`).
- Fork goldens: none.

## Measure at refresh

- Subtlety BiS and weights: Garrote's share rises from 2.7% to about 8.5% and Death Mark marks more often, so attack power, crit (bleed ticks crit, the mark's cap) and energy move.

## For other lanes

- Data/pipeline: wago's DB2 CSV export at build=1.60.1.70755 lacks server-push rows (Hemorrhage 693309's mask is 1048576 there, 1048832 in push 113483). Any lane reading masks from a plain wago CSV should read the layered copy.
- Core: hit is one stat (X2), recorded only.

## For the guide lane

Nothing left over: this lane wrote the rogue text for S1, S2 and round 15's items (see "Rendered review"). The Combat C3 question, if the owner takes it, would change combat.md's "Five points always go to Eviscerate".

## Rendered review

wow-player on the rendered Assassination, Combat and Subtlety guide pages (astro dev on this worktree, real data, 1440 px full-page captures). First pass: Subtlety NOT SHIP (the leveling section still said Eviscerate at five past 40; the pull and the priority disagreed on Death Mark's bleeds), Assassination NOT SHIP (Rupture given unconditionally on a page whose header loads the level 60 build without Blood Splatter), Combat SHIP. All MUSTs and most SHOULDs applied (Garrote's Initiative reason, Garrote before Rupture at exactly four, "Rupture before Death Mark", "every Garrote 9", "two bleeds", the priority split in three, "With daggers" on Combat's Composure sentence). Second pass: **Subtlety SHIP, Assassination SHIP, Combat SHIP.** Open SHOULDs passed on, none in this lane: the "Expertise cap" label (should read dodge cap), Teebu's Blazing Longsword shown as "level 62", blank item and Vanish icons in the capture (confirm on the live page), and a Dirty Deeds "from the front" clause in the level 60 talent paragraph. The rendered rotation card still shows the old Subtlety rows until addon-data is regenerated (Files to regenerate).

## Still open

- K1 Death Mark's storage scope and K2 the off-hand penalty on Shank and Mutilate: no Forever log at 40+ exists yet (foreverlogs.gg stops at report 2813, 2026-10-29). Warcraft Logs stays blocked (gcloud login expired); not worked around.
- S4 Hemorrhage's dagger hand.
- C3 Combat's Eviscerate threshold depends on the weapon.
