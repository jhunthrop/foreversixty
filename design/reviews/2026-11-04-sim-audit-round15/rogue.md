# Sim-accuracy round 15: rogue

Lane `rogue`. Site wt/r15-rogue (from int/r15 fd320b60e), fork lane/r15-rogue (from int/r15 d1d9f2775). Client 1.60.1.70755. Specs: Assassination, Combat, Subtlety.

Fork commits (engine): c9759342f Preparation, a3146478d Blood Splatter aura, plus APL sync commits aae6399c1, bd2dfe400, de21aa7f8 (hand sync of the curated rogue rotations).
Site commits: 96f9fd6a9, 493a3d038, 22e8c28bb, 6e178531b, 4eba2d9d1, and this report.

Measurements: spec-breakdown, raid 60, 10000 iterations, seeds 1 and 50001. Before = $R/before.tsv (fd320b60e / d1d9f2775); after = this lane's engine and site. The three rogues run on their guide raid builds (`$R/rt`) on both sides, as in the brief.

## Assassination

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | The rotation never cast Rupture. In Forever bleeds always crit, and the raid build's Blood Splatter adds 6 s to Rupture. A Rupture at four or more points, with Slice and Dice up and Rupture not ticking, beats an Eviscerate | Raid 60: +0.97/+0.83% alliance, +0.97/+1.06% horde (seeds 1/50001). Without Blood Splatter the same line loses 1.3% at bare 50 (193.3 to 190.8), so it is gated on the talent: `auraIsKnown(1329256)`, a passive aura the engine now gives Blood Splatter. Lead SIM-4(b) | CONFIRMED-BUG | 70755 d009d90a3 (Blood Splatter is new in 70755, and the 70755 rotation pass left Rupture out); guides-handoff | S 22e8c28bb, S 4eba2d9d1, E a3146478d |
| A2 | test_apl: the rotation labelled Mutilate 1241584 "rank 4" and Shank 1329549 "rank 2" | 70755 spellconst: 1241584 is rank 3 (70755 dropped 70753's 1310707 rank at 30), and Shank has no rank subtext (rank 0). The engine resolves casts by id, so these are label fixes. The rotation was wrong, not the test | CONFIRMED-BUG (label only, no DPS change) | 70755 d009d90a3 | S 96f9fd6a9 |
| A3 | TestMutilateRankOneIsLearnedAtThirty was red | In 70755 Mutilate is the 31-point talent (node 105709, tier 6, prerequisite Venom; talents/rogue.json), with rank 1 399956 at 40. 1310707 is gone from trainables and talents. The test was wrong; it is now TestMutilateRankOneIsLearnedAtForty (399956 at 40, 1241582 at 50) | test was wrong | 70755 (data) | S 96f9fd6a9 |
| A4 | Mutilate's 70755 values | 399960/399961 +44 inside 90% (e0 121, e1 31), 1241585/8 +63, 1241586/90 +88; three ranks at 40/50/60 (spellconst, trainables). The engine matches (`mutilate.go`, `MutilateDamage`) | Aligned | - | - |
| A5 | Poisons "up 40-50%" (round-14 build R3) | The changed amounts are 11340/11343/11356/11358/25347/25351/13227/13230 e1/e2 (175 to 250, 160/180 to 240), the enchant item spells' charge counts. The damage spells are unchanged (Instant VI 11337 is 88 ±27.7%, Deadly IV 11354 is 18 a tick, both 70753 = 70755). `client_damage.go` matches | Aligned (charges only; the sim does not model charges) | - | - |
| A6 | Venom (1329328 to 1310703, 6 s on every Shank): e0 aura 108 +30% damage, e1 aura 108 misc 22 +30% periodic, e2 aura 107 misc 18 +20% chance | The engine has +30% for Instant and Deadly and +20% proc chance (`venom.go`, `poisons.go`) | Aligned | - | - |
| A7 | Spells and talents checked: Cold Blood (14177, 3 min, next Eviscerate/Mutilate crit), Slice and Dice / Improved Slice and Dice 15% a rank, Eviscerate with Coup de Grace 7/13/20% (curve 83054), Relentless Strikes 4% a point a rank (curve 128113) on every finisher, Ruthlessness 20% a rank, Seal Fate 33/66/100%, Lethality 10% a rank, Vile Poisons, Improved Poisons, Opportunity 4% a rank on Mutilate, Puncturing Wounds 5% a rank on Mutilate, Malice 1% a rank, Vigor 5 a rank, Garrote opener with Poised | spellconst 70755, the 70753-to-70755 rogue spellconst diff (below), engine `talents.go`/`client_values.go` | Aligned | - | - |

## Combat

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | The raid build spent 2 points on Efficient Strikes, which discounts only Sinister Strike and Hemorrhage (13732; the 70755 rank text). The raid build is a Backstab build: Sinister Strike casts 0 times a fight at raid 60. Composure 2/2 (Poised, which opens every pull, costs 10 more energy less) fills the same row 0-2 gate to Flawless Execution, Weapon Expertise and Aggression | raidBuild `005203002/3243030000012321451/005` to `005203002/3043030002012321451/005`. Within error at seeds 1 and 50001 (756.4/755.7 to 756.2/755.9). The two points did nothing; their replacement is legal, does something and costs nothing measurable (adoption rule). Lead SIM-10 | CONFIRMED-BUG | 70755 3d4712981; guides-handoff | S 6e178531b |
| C2 | Other raid-build points | Lightning Reflexes 4 is gate filler: rows 0-2 hold nothing else that does damage on a Backstab build (Deflection, Improved Sprint, Endurance, Riposte). Aggression 3 works on Eviscerate only (mask 131074), and the guide already says so. Tested Hack and Slash 5 / Flash of Steel 4: -1.2%. Ruthlessness 3 / Improved Slice and Dice 2: within error. Kept | Aligned | - | - |
| C3 | Rupture on a boss (SIM-4a) | A five-point Rupture line ahead of Eviscerate: -2.1% at raid 60. The prose "Rupture is not worth a global on a boss" stands | Aligned | - | - |
| C4 | Level 60 / leveling build: Improved Gouge 3, Remorseless Attacks 2, Master of Deception 3, Efficient Strikes 1 | Efficient Strikes works with the leveling Sinister Strike. The other three are leveling and utility talents that the guide argues for in prose, not points that provably do nothing for a leveling build | Aligned (taste, not a bug) | - | - |
| C5 | Shank rank 1 never cast from 24 to 43 (all specs; see X1) | Combat bare 30 +7.2%, bare 40 +11.6% | CONFIRMED-BUG (X1) | X1 | S 96f9fd6a9 |
| C6 | Shank (1329283/1329549: 5 energy, 15/25 back, 10 s, 1 s GCD, 50% off-hand, ALWAYS_HIT, procs suppressed), Flash of Steel +3 energy and -0.4 s a rank (curves 83037/128124), Adrenaline Rush, Blade Flurry (13877: 25 energy, 2 min), Backstab with Puncturing Wounds and Opportunity, Hack and Slash (dagger crit is melee crit only), Weapon Expertise, Dual Wield Specialization 25% a rank, Flawless Execution -15 (1310711) | spellconst 70755 and the engine (`shank.go`, `talents.go`, `client_values.go`) | Aligned | - | - |

## Subtlety

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | Eviscerate (three points, Rupture ticking) sat above nothing that waited for Death Mark, so a five-point Death Mark came only when a 2-point jump happened to land on five: 6.3 marks a fight at raid 60 out of about 17 possible, and 0 at level 40 on the ladder. Death Mark stores up to 60% (72% with Coup de Grace 3) of everything for 25 energy. Eviscerate is now the finisher only before Death Mark is learned | Rotation variants at raid 60, horde, seed 1: Eviscerate gated on Death Mark more than 5 s from ready 653.4; more than 8 s 663.9; never once Death Mark is known 663.9 to 664.2 (base 635.8). Leads SIM-2, ROGUE-9 | CONFIRMED-BUG | 70755 d009d90a3; guides-handoff | S 493a3d038 |
| S2 | Shank with Boot Knife 5/5 always adds 2 points (1329286 trigger 1329285, 100%), and Ghostly Strike adds 1. They were pressed with no combo-point check: Ghostly Strike's points 8.7 gained, 4.5 kept; Boot Knife's 36.3 gained, 33.9 kept | Shank now at three points or fewer and Ghostly Strike at four or fewer. Shank at ≤3 or ≤4 is a tie (679.1 vs 679.5, ±0.2), so it follows the guide-lane prose (≤3). Web search for Forever sources found none (see Sources). Leads SIM-8, ROGUE-4 | CONFIRMED-BUG | old (Ghostly Strike) / 70755 d009d90a3 (Boot Knife Shank); guides-handoff | S 493a3d038 |
| S3 | Ghostly Strike's 1329326 ("next Ambush or Garrote does not require Stealth", 7 s) was modelled in the engine (`ghostly_strike.go` canOpen) but never used by the rotation | A Garrote line right after Ghostly Strike (Garrote not ticking, four points or fewer): +0.8% on top of S1+S2 (663.9 to 669.4). Lead SIM-5 | CONFIRMED-BUG | 70755 d009d90a3; guides-handoff | S 493a3d038 |
| S4 | Preparation was never used, and reset too little. `preparation.go` reset only Cold Blood, Shadowstep and Vanish, and its automatic use waited for Vanish to be on cooldown, which nothing casts (0 casts a fight). 14185: "immediately finishes the cooldown on your other Rogue abilities" (70755 rank text; its only effect is a server dummy, effect 3) | It now resets every rogue cooldown (Premeditation, Shank, Ghostly Strike, Death Mark, Cold Blood, Adrenaline Rush, Blade Flurry, Evasion, Kick, Gouge, Shadowstep, Vanish). It goes once a long DPS cooldown has more than 30 s left, Shank and Ghostly Strike are spent, no Death Mark is storing, the rogue is out of Stealth and Premeditation's points would fit. A naive "on cooldown" use cost 1% (a GCD at the pull), hence the gates. +0.5% on the old rotation, +1.5% on the new one (669.4 to 679.5). Tests: TestPreparationFinishesEveryOtherRogueCooldown, TestPreparationIsUsedOnceAResetIsWorthIt. Lead SIM-7 | CONFIRMED-BUG | old (upstream 1608caf21, 2023; no sim round looked); guides-handoff | E c9759342f |
| S5 | Slice and Dice is shadowed by Eviscerate (SIM-3) | Moving it ahead of Eviscerate changes nothing at bands 30, 40, 50 or raid 60 (identical DPS). With S1, Eviscerate no longer runs from 40, so the line is live there. Not changed | Aligned (no effect) | - | - |
| S6 | Premeditation in the fight (SIM-6) | Already fixed in round 14 (E e01d5596a, at the pin). The major cooldown fires at ≤3 points, so no APL line is needed (autocastOtherCooldowns) | Aligned | - | - |
| S7 | Shank rank 1 never cast from 24 to 43 (X1) | Subtlety bare 30 +6.1% (S1-S4 do not apply there) | CONFIRMED-BUG (X1) | X1 | S 96f9fd6a9 |
| S8 | Death Mark (1316709/1325756/7/8 at 40/48/56/60: 25 energy, 10 s, 6 s; storage 12% a point, bleeds ×2; cap from SpellDescriptionVariables 1021), Hemorrhage (35 energy, 100/145%, debuff +15% Rupture and Garrote, mask 1048832), Ghostly Strike (40 energy, 20 s, 125/180%), Blood Splatter +6/+9 s (curves 109811/128090), Serrated Blades 3% armor a rank and +30% Rupture, Initiative 33/67/100%, Dirty Deeds -10 a rank, Efficient Strikes -3/-5 on Hemorrhage, Opportunity on Garrote's periodic (misc 22), Ruthlessness and Relentless Strikes on Death Mark (OnComboPointsSpent) | spellconst 70755 and the engine | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | No rogue cast Shank from 24 to 43. Shank's two trainer rows (1329283 at 24, 1329549 at 44; SkillLineAbility 59899/59900) carry no "Rank N" subtext, so spellranks.json lists both at rank 0. The rotation rewrite then treated 1329549 as untracked and left it unresolved below 44 (ladder `unresolved_id {SpellID: 1329549}` at 30, 38 and 40 for every spec). `sim/internal/spellranks` now chains the two rows by trainer level (`unrankedRankOverrides`). Test: TestShankRanksChainByLevel | Bare 30: Assassination +9.0%, Combat +7.2%, Subtlety +6.1%; bare 40: +6.7%, +11.6% (Subtlety +21.4% with S1-S4). Lead SIM-1 | CONFIRMED-BUG | 70755 d009d90a3; guides-handoff | S 96f9fd6a9 |
| X2 | Client diff 70753 to 70755, rogue spellconst | Changed: poison charge counts (A5), Mutilate (A4), Venom (A6), Flawless Execution -15, Coup de Grace (gains e1), Efficient Strikes (renamed), Ghostly Strike (gains 1329326), Quick Getaway. New: Blood Splatter, Composure, Shank, Flash of Steel, Boot Knife, Venom 1329328, Death Mark. Gone: 1310705/6/7, 1310709. Every changed or new spell above is followed by the engine (fork a57b0221b) | Aligned | - | - |
| X3 | Mutilate's off-hand skips the 0.5 dual-wield penalty | Pending owner (brief). Recorded only | pending owner | - | - |

## Before/after

Raid 60, guide raid builds (relative %, seeds 1 / 50001; absolute at seed 1):

| spec | alliance | horde | seed 1 absolute |
|---|---|---|---|
| rogue-assassination | +0.97 / +0.83 | +0.97 / +1.06 | A 745.7 to 752.9, H 738.7 to 745.9 |
| rogue-combat | -0.03 / +0.00 | -0.03 / +0.03 | A 756.0 to 755.8, H 756.4 to 756.2 |
| rogue-subtlety | +6.68 / +6.78 | +6.81 / +6.76 | A 640.3 to 683.1, H 635.8 to 679.1 |

Bare bands, horde, seed 1, 3000 iterations, ladder-truncated guide `build`:

| spec | 30 | 40 | 50 | 60 |
|---|---|---|---|---|
| Assassination | 57.9 to 63.1 (+9.0%) | 92.5 to 98.7 (+6.7%) | 193.3 to 193.3 (after the Blood Splatter gate) | 302.5 to 302.5 |
| Combat | 64.2 to 68.8 (+7.2%) | 106.4 to 118.7 (+11.6%) | 187.0 to 187.0 | 274.2 to 274.2 |
| Subtlety | 58.7 to 62.3 (+6.1%) | 94.3 to 114.5 (+21.4%) | 162.5 to 176.9 (+8.9%) | 270.6 to 299.2 (+10.6%) |

No commit moves DPS for any spec outside the rogue class. The spellranks change touches only Shank's two ids. The Preparation and Blood Splatter changes are in `sim/rogue`.

## Files to regenerate

- Ladder goldens `sim/request/testdata/ladder/rogue-assassination|combat|subtlety.golden.md` (Shank resolves 24-43, the new Subtlety and Assassination lines, Eviscerate idle for Subtlety from 40, Preparation casts).
- `data/builds/1.60.1.70755/addon-builds.json` (Combat raidBuild changed; TestCommittedAddonBuildsFollowTheGuides fails until then), then Data.lua.
- `dps-per-point.json` for the three rogues (after the pin), then Data.lua.
- apl-sync / apl-check: the rogue copies are already hand-synced on both sides (site `sim/request/apl/rogue-*.apl.json`, fork `ui/rogue/apls/forever_*.apl.json`); `make apl-check` should pass.
- Fork pin: the site's Assassination Rupture line needs fork a3146478d (Blood Splatter's aura). On the old pin, `auraIsKnown(1329256)` is unresolved and the line simply never fires, so the merge order is safe either way. The smoke test entry added in 4eba2d9d1 expects the aura to be unresolved on a build without the talent.
- Fork goldens: none changed. `./sim/rogue/...` (with_db) and `TestConformance/rogue` are green on lane/r15-rogue.

## Measure at refresh

- All three rogues' BiS picks and weights: Subtlety +6.8% raid and +6 to +21% on the bands, with crit and energy weights likely to move (Death Mark's cap scales with crit, Preparation, Garrote); Assassination about +1% (Rupture, more weight on crit and AP for bleeds); Combat with Composure (no DPS change).
- The published BiS talent strings for the rogues are illegal on 70755 (brief); the refresh should take the guide strings above.

## For other lanes

- Warrior lane: `data/tests/test_apl.py::test_every_named_spell_matches_its_own_label_by_name` is red on warrior-fury (its sidecar of expected names lacks 1329555). Not rogue; untouched.
- Ladder mirror (owner of `sim/request/ladder.go`): `idLearnLevel` skips rank-0 rows, so it does not know Shank's two learn levels from `unrankedRankOverrides`. This does not cause violations (the rewrite drops Shank below 24 and resolves 1329283 from 24), but the "learned but unused" report will not list Shank as a chain. If wanted, mirror `spellranks.unrankedRankOverrides` there the way `idLearnLevelOverrides` mirrors `singleTierLevelOverrides`.

## For the guide lane

Line numbers are from origin/land/guides-r1.

1. `combat.md:52`, sentence "The level 60 build is the one you level and run dungeons with: 14/34/3: ... and the fifth points of Hack and Slash and Lightning Reflexes in place of the fifth Flash of Steel and the second Efficient Strikes." Replace with: "The level 60 build is the one you level and run dungeons with: 14/34/3: Improved Gouge and Remorseless Attacks in place of Lethality and two points of Improved Slice and Dice, Master of Deception in place of Opportunity, a third point of Ruthlessness, and a point of Efficient Strikes and the fifth points of Hack and Slash and Lightning Reflexes in place of Composure and the fifth Flash of Steel." (12 out, 12 in.)
2. `combat.md:63`, sentences "Efficient Strikes takes energy off Sinister Strike, 3 at one point and 5 at two: your builder's discount in the level 60 build. In the raid build its two points and four of Lightning Reflexes are the toll to reach Flawless Execution, Weapon Expertise and Aggression." Replace with: "Efficient Strikes takes 3 energy off Sinister Strike at one point: your builder's discount in the level 60 build. The raid build never presses Sinister Strike, so it takes Composure instead, 10 more energy off your Poised opener, and Composure and four points of Lightning Reflexes are the toll to reach Flawless Execution, Weapon Expertise and Aggression."
3. `assassination.md:80`, sentence "Rupture and Expose Armor do not make the list." Replace with: "With Blood Splatter, put Rupture up at four points whenever Slice and Dice is running and Rupture is not: bleeds always crit, and 6 more seconds makes it worth the global. Expose Armor does not make the list." Add to the raid-build talent prose, if it explains Blood Splatter: "Blood Splatter is what puts Rupture in the rotation."
4. `subtlety.md:76`, sentence "While Rupture ticks and Death Mark is on cooldown, spend three or more points on Eviscerate, or on Slice and Dice when that is down or about to fall off." Replace with: "Once you have Death Mark, every point Rupture does not need goes into the next mark, never into Eviscerate, even with Death Mark on cooldown. While Rupture ticks, spend three or more on Slice and Dice when that is down or about to fall off." Add after the Ghostly Strike sentence: "When a Ghostly Strike lands and Garrote has dropped off, put Garrote straight back on: the strike lets it in without Stealth."
5. `subtlety.md:117` (leveling): "Eviscerate at five" is still right before 40.
6. `rogue-subtlety.json` notes, which I did not edit. Shank: "Shank on cooldown at three combo points or fewer: an off-hand stab that cannot miss, pays back more energy than it costs and, with Boot Knife, adds two combo points." Ghostly Strike: "Ghostly Strike on cooldown at four combo points or fewer: it hits harder than Hemorrhage for 10 more energy and gives a combo point." Eviscerate: "Eviscerate at three or more combo points while Rupture is ticking, only before you have Death Mark." The new Garrote entry carries a note I wrote ("Garrote right after Ghostly Strike whenever it has dropped off: ..."); please review it for voice. The top-level `notes` still says "Eviscerate and Slice and Dice spend the combo points that build while it runs"; replace with "Death Mark (Eviscerate before you have it) and Slice and Dice spend the combo points that build while it runs".
7. `rogue-assassination.json`: the new Rupture entry has a note I wrote; please review it. The top-level `notes` says "combo points are spent from three: Cold Blood and Eviscerate, with Slice and Dice ahead of it whenever it has lapsed". Add: "With Blood Splatter, Rupture goes up at four whenever Slice and Dice is running and Rupture is not."
8. `subtlety.md:86` (Preparation) now matches the engine; no change.

## Still open

- Death Mark storage scope: the engine stores every hit except proc-suppressed ones (Shank), including white swings and poisons. "the damage dealt by your attacks" is the only source, and Titan Reforged notes were not found. An in-game log of a mark's payout against its cap would settle it. OPEN.
- Mutilate off-hand 0.5 penalty: pending owner (X3).
- Warcraft Logs is blocked (gcloud login expired). Not worked around.
- Web research (owner rule, round 15). Death Mark vs Eviscerate is not a tie (+4.5%), so it is decided on the sim. For the Shank ≤3 vs ≤4 tie I found no Forever source. Searched 2026-11-03: Warcraft Tavern Forever Subtlety guide (https://www.warcrafttavern.com/forever/guides/subtlety-rogue/, undated, pre-70755, no Death Mark or Shank); Icy Veins Forever Subtlety (https://www.icy-veins.com/wow-forever/subtlety-rogue-melee-dps-pve-guide, 403); Blizzard Forever class deep dive (https://us.forums.blizzard.com/en/wow/t/2375390) for Death Mark's origin. The tie followed the guide lane's existing prose (≤3).

## Wave 2: Mutilate's off-hand penalty

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| W1 | Mutilate's off-hand strike read the off-hand dagger at full weight. It should take the 50% off-hand penalty on its weapon damage; the flat per-hand bonus stays whole, and Dual Wield Specialization raises the off-hand weapon damage | Client: SpellMisc (wago 1.60.1.70755). The off-hand hand spells 399961, 1241588 and 1241590 have Attributes_3 16777728, which includes REQUIRES_OFF_HAND_WEAPON (0x1000000). The main-hand 399960, 1241585 and 1241586 have 1536, without it. That flag makes the strike an off-hand attack, the same flag Shank carries, so it uses the off-hand weapon damage the server halves for every off-hand swing. Second source: guide-accuracy round 1 owner item 1 (origin/land/guides-r1 `design/reviews/guide-accuracy/round1/rogue.md:119-122`). wowsims SoD and Wrath `mutilate.go` use `OHNormalizedWeaponDamage` (0.5 times the roll) plus the full flat. Forever notes are silent, so the Classic rule holds. Test: TestMutilateOffHandTakesTheOffHandPenalty (179.7 vs 89.9 before the fix). The brief listed this as pending owner; the coordinator reopened it for wave 2 on this evidence | CONFIRMED-BUG | old (fork 56f97823e's "does not apply to a special ability that explicitly swings the off-hand" comment); guides-handoff | E db91a347d |

Measured on Assassination, wave-1 engine to wave-2 engine, 10000 iterations:

| case | seed 1 | seed 50001 |
|---|---|---|
| raid 60 alliance | 752.9 to 707.8 (-5.99%) | 752.6 to 707.6 (-5.98%) |
| raid 60 horde | 745.9 to 701.2 (-5.99%) | 745.8 to 701.1 (-5.99%) |
| bare 40 horde | 97.8 to 97.8 (the truncated level 40 build has no Mutilate yet) | 97.8 to 97.8 |
| bare 50 horde | 193.3 to 185.4 (-4.09%) | 193.4 to 185.5 (-4.08%) |
| bare 60 horde | 302.3 to 290.6 (-3.87%) | 302.2 to 290.6 (-3.84%) |

Net against the round baseline ($R/before.tsv), raid 60: alliance 745.7 to 707.8 (-5.1%), horde 738.7 to 701.2 (-5.1%). No other spec casts Mutilate.

The Rupture line (A1) still gains after the fix: 695.2/694.7 without it, 701.2/701.1 with it (+0.86%/+0.92%, seeds 1/50001, horde). Guide-accuracy owner item 4 cites Forever guides that skip Rupture, but the sim decides beyond error with Blood Splatter and the always-crit bleed rule, which those guides do not price in. So A1 stays.

Wave 2 additions:
- Files to regenerate: the Assassination ladder golden, Assassination dps-per-point and Data.lua (already listed).
- Measure at refresh: Assassination BiS and weights. The main-hand dagger's weight rises against the off hand.
- For the guide lane: `assassination.md:103` on land/guides-r1 already says the off-hand Mutilate strike gets half its dagger's weapon damage with the flat bonus whole. It now matches the engine; no change.
