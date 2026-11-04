# Sim audit, round 14: physical group, fresh re-audit (2026-11-03)

This lane covers Arms, Fury, Assassination, Combat, Subtlety, Beast Mastery, Marksmanship and Survival (melee). Protection is lane tanks'. It re-audits them against client 1.60.1.70753 with the 2026-11-03 hotfix push layered in.

- **Baseline:** fork int/r14 5195a8889 with site int/r14 72471fcb3 (`r14/sb-before`, `r14/before.tsv`).
- **Branches:** fork `lane/r14-physical` (E), this report on `wt/r14-physical` (S). No site code changed.
- **Client rows:** the hotfix lane's layered 70753 CSVs (`scratchpad/r13/hotfix/root/1.60.1.70753/raw`), plus SpellShapeshift, SpellCategories and SpellEquippedItems fetched from wago.tools for 1.60.1.70753 on 2026-11-03 (`scratchpad/r14/physical/wago/`). Lane helpers are in `scratchpad/r14/physical/tools/` (sp.py, ranks.py, item.py, tal.py, apl.py). A tooltip and its client row count as one source.
- **Wowhead Forever tooltips** (read 2026-11-03 evening) still show the pre-push numbers for Aimed Shot (166), Arcane Shot (217), Serpent Sting (415 over 15 s), Trueshot Aura (50) and Aspect of the Hawk (55). Sniper Shot now reads 590, which matches the push, so that value has a second source.

**Result: 1 confirmed bug, fixed: Premeditation (old).** Everything else checked below is aligned or carried open.

## Assassination, Combat, Subtlety

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | **Premeditation does not need Stealth in Forever, and the engine required it.** In Classic Era, Premeditation only works from Stealth. In Forever's client it doesn't: its SpellShapeshift row has no Stealth form, while Garrote, Ambush, Cheap Shot and Sap all do, and it lacks the stealth-only attribute that Garrote has. Wowhead's Forever tooltip has no "Requires Stealth" line. Wowhead's Classic tooltip for the same spell has one, and so does Forever's own Garrote tooltip. The engine only allowed it in Stealth, so Subtlety used it once, before the pull. Now it comes back in combat every 2 minutes. The automatic use waits until the rogue has 3 combo points or fewer, so neither point is lost over the cap of 5. Test `TestPremeditationDoesNotNeedStealth` failed first ("not castable out of stealth"). Round 12's regmelee L7 marked "stealth is required" as aligned. | SpellShapeshift 39913 (14183: ShapeshiftMask 0) against 39860/39865/39457/39625 (Garrote, Ambush, Cheap Shot, Sap: 536870912); SpellMisc 14183 Attributes_0 0x50010 against Garrote 0x270010 (ONLY_STEALTHED 0x20000); https://nether.wowhead.com/forever/tooltip/spell/14183 and https://nether.wowhead.com/classic/tooltip/spell/14183 (2026-11-03); `sim/rogue/premeditation.go` | CONFIRMED-BUG | old (upstream stealth gate; round 12 regmelee L7 looked and missed it) | E e01d5596a |
| R2 | **Finisher amounts and shapes.** Eviscerate rank 8 is 96 plus 151 a combo point. It costs 35 Energy on a 1 s GCD, and Flawless Execution makes it 25 (448 Energy over 17.9 casts). Rupture rank 6 is 35 plus 4.73 a combo point a tick, every 2 s, for 3 + combo points ticks. Serrated Blades adds +30% (aura 108 misc 22 on the Rupture mask) and Hemorrhage +15%. Slice and Dice gives +30% attack speed (aura 319) for 9/12/15/18/21 s, and Combat's uptime matches its 1 and 2 point casts (8.19×9 + 4.17×12 = 124 s of 180). Garrote is 92 a tick for 6 ticks and 50 Energy, less Poised's 10 and Dirty Deeds' 20. | SpellEffect 689090, 687822, 684799, 689794, 690916; `rogue/rupture.go`, `eviscerate.go`; `bd/rogue-*.txt` | aligned | - | - |
| R3 | **Builders.** Backstab: 150% weapon damage of normalized damage +140, 60 Energy. Sinister Strike: normalized +68, 45 Energy. Mutilate rank 4: 75% of (normalized +67) per hand, 60 Energy, +20% against a poisoned target. Hemorrhage: 100% normalized (145% with a dagger), 35 Energy. Ghostly Strike: 125%/180% of unnormalized weapon damage (effect 31 only, no 121), 40 Energy, 20 s. The engine's normalized/unnormalized choice follows effect 121 in every case. | SpellEffect rows for 11281, 11294, 1241584/86/90, 16511, 14278; `rogue/backstab.go:71`, `ghostly_strike.go:64`, `hemorrhage.go:97`, `mutilate.go:157` | aligned | - | - |
| R4 | **Cooldowns.** Blade Flurry: +20% attack speed (aura 319) for 15 s, 2 min, 25 Energy; 2 casts and 16.7% uptime. Adrenaline Rush: +100% Energy regeneration for 15 s, 5 min. Cold Blood: +100% crit on the next Sinister Strike, Backstab, Ambush, Eviscerate or Mutilate, 3 min. Energy regenerates at 10 a second with no tick (Combat: 1977.7 over 180 s, less Adrenaline Rush's 150). | SpellEffect 691767, 691038, 692068; SpellCooldowns | aligned | - | - |
| R5 | **Poisons.** Instant Poison VI: 88 ±27.7% Nature, 20% on a hit. Deadly Poison IV: 18 a tick every 3 s for 12 s, stacking to 5, 30% on a hit. Both use the spell hit table and crit at 1.5× (SpellCategories DefenseType 1 on 11337 and 11354). Stormshroud's Lightning procs on 10% of hits (26.4 of about 279 landed for Assassination). Revitalize procs on 4% for 30 Energy. | SpellEffect 688955, 689744; SpellAuraOptions 18979, 23863; SpellCategories 73648, 73906 | aligned | - | - |
| R6 | **Talents not re-read in round 13.** Murder is +4% against Humanoid and Giant only. The engine gives nothing against the type-neutral target, and raid-keeps.json keeps it on purpose. Vile Poisons, Improved Poisons, Relentless Strikes (20% a combo point, 25 Energy) and Ruthlessness (60%) match their texts. | talents/rogue.json; `rogue/talents.go:142`; design/reviews/talent-search/raid-keeps.json | aligned | - | - |
| R7 | **Openers.** Stealth at -3 s, Premeditation at -2 s (Subtlety), Garrote at 0 s. Poised takes 10 Energy off the Garrote (Assassination and Combat 40, Subtlety 20). Major Frenzy Potion (+80 attack power for 30 s, 2 min) twice. Thistle Tea is SpellCategory 1153, not the potion category 4, so it doesn't share the potion cooldown. | SpellEffect 1368039; ItemEffect 98288 (Thistle Tea category 1153), 214812 | aligned | - | - |
| R8 | **Cold Blood in Subtlety (asked by lane guides): the build takes it and the APL has no line for it. That is not a bug.** The engine models it (+100% crit on the next Sinister Strike, Backstab, Ambush, Eviscerate or Mutilate, the client mask; Hemorrhage and Ghostly Strike do not use it up). The engine registers it as a major cooldown, so the rotation's first line casts it on cooldown: 1.48 casts a fight, waiting for an Eviscerate (Eviscerate at 3 points crits 61%). An explicit "Cold Blood before Eviscerate" line (3+ points, Rupture up, 35 Energy) gives the same DPS: raid 60 546.9 / 546.6 and bare 60 223.4 / 223.3 → 223.2, seeds 1 / 50001, 10,000 iterations. Without the talent the raid build does 544.5 / 544.2 (-0.44%). The point stays under the adoption rule (under 1%), and the APL needs no line. | `sim/rogue/talents.go:175` (registerColdBloodCD, AddMajorCooldown); SpellEffect 692068; variant `r14/physical/rot/sub-cb.json` | aligned (no change) | - | - |

## Arms, Fury

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| W1 | **Rage cost of every cast, from the breakdown's rage spent divided by casts.** Mortal Strike 30, Bloodthirst 30, Whirlwind 25, Slam 15, Heroic Strike 15 (12 with Improved Heroic Strike), Overpower 5, Rend 10, Hamstring 10, Execute 15 less Improved Execute's 5, plus all remaining rage at 15 damage a point. Death Wish 10. All match SpellPower. | SpellPower rows; `bd/warrior-*.txt` | aligned | - | - |
| W2 | **Stance requirements against SpellShapeshift.** Overpower, Sweeping Strikes, Charge and Retaliation: Battle Stance. Whirlwind, Recklessness, Berserker Rage, Intercept and Pummel: Berserker Stance. Execute and Hamstring: Battle or Berserker. Rend and Thunder Clap: Battle or Defensive. Fury lists Overpower in Berserker Stance, where it can never cast. That is inert, not a loss. | SpellShapeshift masks 65536 / 262144 / 327680 / 196608 | aligned | - | - |
| W3 | **Whirlwind hits with both weapons in Forever.** The off-hand part (462891) is a separate normalized strike with the off-hand attribute. Off-hand hits average 0.64 of main-hand hits, which matches 50% off-hand damage plus Dual Wield Specialization's 25%. 10 s cooldown. | SpellEffect 680030, 1173118; SpellMisc 462891 Attributes_3 0x1000200 | aligned | - | - |
| W4 | **Recklessness:** +100% crit (aura 290) and +20% damage taken for 15 s, 30 min, on the GCD. **Death Wish:** 30 s, 3 min, on the GCD. **Sweeping Strikes:** 5 charges, 20 s, 30 rage (not cast single target). **Anger Management:** 1 rage every 3 s (59.5 a fight). **Bloodrage:** 1 min, off the GCD, +50% from Improved Bloodrage. Every engine GCD matches SpellCooldowns StartRecoveryTime. | SpellEffect 687178/9, 690203/5, 689601, 689409; SpellCooldowns | aligned | - | - |
| W5 | **Normalized rage.** Fury main hand (2.4 s one-hander): 8.3 rage a landed swing, ×2 on a crit, extra attacks included. Expected about 768 a fight; the engine gives 823.6. Off hand (2.3 s): ×0.5, ×1.5 with Dual Wield Specialization, expected about 692 against 720. The off-hand factor and the crit multiplier are still the memory note's unverified constants. | `core/rage.go:100`; ItemSparse ItemDelay 279259 / 279261; forever-rage-normalization-gap.md | aligned (constants still unverified, carried) | - | - |
| W6 | **Arms Hamstring weaving.** Without it, Arms loses 2.8% (730.1 → 710.2, seed 1, 3000 iterations). A lower Heroic Strike threshold (50 rage) loses 2.0%. Both APL choices hold. | `r14/physical/rot/arms-*.json` | aligned | - | - |
| W7 | **Deep Wounds.** Forever's bleed is 412609 "Deep Wound": a 12 s periodic dummy every 3 s, so 4 ticks at 60% of average weapon damage, as the engine does. Whether the attacker's damage multiplier applies twice is still open (round 13 W6). Warcraft Logs is unreachable (gcloud login expired). | SpellMisc/SpellEffect 412609; `warrior/deep_wounds.go` | OPEN (carried) | - | - |
| W8 | **Proc gear.** Arcanite Champion's Strength of the Champion is +120 Strength for 30 s. Molten Heart of the Mountain is +55 attack power for 15 s on a 5 min cooldown (doubled only in volcanic areas, which the type-neutral fight isn't in). Rune of the Guard Captain is +14 attack power, with its forest/grassland rider left out by `equip.py`'s area rule. Devilsaur's 2-piece is +2% hit. All match. | SpellEffect 694634, 1252393, 1351263; `data/pipeline/simdb/equip.py:382` | aligned | - | - |

## Beast Mastery, Marksmanship, Survival

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | **Shot ranks against the pushed client, rank by rank.** Aimed Shot 40/70/115/185/260/350 at 20-60, costs 75-310. Arcane Shot 25 to 280, costs 25-190. Serpent Sting 3 to 100 a tick, rank 9 (135) is AQ. All are 6 s category cooldowns except Serpent Sting. Multi-Shot is one rank in Forever: normalized weapon damage with no flat bonus, 13.9% of base mana. Aimed Shot has a 2 s cast scaled by ranged haste. All match the engine's tables. | SpellEffect/SpellPower/SpellLevels for 19434-20904, 3044-14287, 1978-25295, 2643; `hunter/aimed_shot.go`, `arcane_shot.go`, `serpent_sting.go`, `multi_shot.go` | aligned | - | - |
| H2 | **Shared cooldown and defense types.** Multi-Shot and Aimed Shot share category 2. Arcane Shot is category 1173, and Summon Hawk shares Arcane Shot's timer. Serpent Sting is DefenseType 3 (ranged), so its ticks crit at the ranged 2× plus Mortal Shots, whose class mask includes the sting. | SpellCategories 71097, 76230, 74283, 74172; SpellEffect 696507 mask | aligned | - | - |
| H3 | **Talents read again.** Lethal Attacks +5% crit, Careful Aim 100% of Intellect to both attack powers, Mortal Shots +30% crit bonus, Barrage 3/7/10%, Ranged Weapon Specialization +5%, Improved Stings 6/13/20%, Focused Fire +2%, Lone Wolf +20% with no pet, Rapid Killing -2 min, Efficiency -15%, Deadly Aspects 10% for +30% for 12 s, Frenzy 100% for +30% for 8 s, Bestial Wrath +50% for 18 s every 2 min at 12% of base mana, Ferocity +10% and Unleashed Fury +15%. | SpellEffect 697561, 1212378/1323702, 696507, 696995, 1340331, 1085262, 697222/3; `hunter/talents.go` | aligned | - | - |
| H4 | **Ranged Weapon Specialization's reach.** The client aura is 79 on every school (misc 127) and requires a ranged weapon (SpellEquippedItems 179969: bow, gun, crossbow). Under the usual server rule, that applies to every spell while a ranged weapon is equipped, Serpent Sting included. The engine applies it to Arcane Shot but excludes Serpent Sting, an upstream choice (8b2bd3d32). This is worth at most 0.3% for Marksmanship. | SpellEffect 696995; SpellEquippedItems 179969; `hunter/talents.go:143` | OPEN (a log of Serpent Sting ticks with and without the talent would settle it) | - | - |
| H5 | **Rapid Fire:** +40% ranged (aura 140) and +40% melee (aura 319) attack speed for 15 s, 5 min, 100 mana. The engine speeds up both, so Survival's melee benefits. **Aspect of the Hawk:** rank 6 is 110 for 110 mana. **Hunter's Mark:** 143 ranged attack power at 58, ranged attacks only. **Trueshot Aura:** 100, ranged only. | SpellEffect 679813, 1339717, 692138, 690991; `core/debuffs.go:693` | aligned | - | - |
| H6 | **Survival's melee kit.** Raptor Strike rank 8: +70 unnormalized, 100 mana, 6 s, next swing. Mongoose Bite rank 4: +57 normalized, 65 mana, 5 s, opened by your own dodge or by Expose Prey (10%, 10 s) against a Hunter's Mark target. Wing Clip: 50 Physical school damage, 80 mana. Strider Kick: 100% normalized weapon damage, 5.81% of base mana, 8 s. Lacerate: 58 a tick, 7 ticks, 95 mana. Lacerating Strikes: 40% of the bite over 7 ticks. Aspect of the Beast: +110 melee attack power. Immolation Trap: 138 a tick for 5 ticks, 30 s cooldown (Forever, not Era's 15), 18 s with Survivalist's Discipline 2/2 (9.7 casts). Resourcefulness takes 60% off melee and trap costs. Every rank and level matches. | SpellEffect/SpellPower/SpellLevels for 2973-14266, 1495-14271, 2974-14268, 1317257, 1299332, 1310533/6, 1299447, 14301/14305, 1340058; `hunter/*.go` | aligned | - | - |
| H7 | **Lacerating Strikes on refresh.** A new bite replaces the bleed's snapshot, so ticks still to come are lost. The bleed comes to 22.7% of bite damage, not 40%. Whether Forever rolls the rest of the bleed into the new one isn't stated in the client. | `hunter/lacerating_strikes.go:96`; `bd/hunter-survival.txt` (52.6 ticks × 63 against 18.1 bites × 805) | OPEN (a Survival log with overlapping bites would settle it) | - | - |
| H8 | **Improved Tracking against the type-neutral boss.** The engine grants it, on purpose: every Phase 1 boss is a trackable type, unlike Murder's two types. This is a stated modelling choice, unchanged. | `hunter/talents.go:661` | aligned (choice) | - | - |
| H9 | **Auto Shot cadence.** BM fires 98.0 shots in 180 s from a 2.5 s rifle. Quiver 15%, Deadly Aspects 30% at 45% uptime, Rapid Fire and Berserking predict 98.2. | `bd/hunter-beast-mastery.txt` | aligned | - | - |
| H10 | **Survival's dual-wield white misses** (main hand 18.4%, off hand 15.5%). A queued Raptor Strike lifts the off hand's dual-wield penalty while it waits: the vanilla on-next-swing behaviour, which isn't confirmed for Forever. | `hunter/raptor_strike.go` (DisableDWMissPenalty) | OPEN (carried from round 13 H3) | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **Requirement sweep after R1.** Every stealth, stance and "after a dodge" ability in the eight rotations was checked against SpellShapeshift and Attributes_0. The cannot-be-dodged flag (0x200000) is only on Overpower, Garrote, Ambush and Lacerating Strikes, and the engine matches. Nothing else differs from the engine. | SpellShapeshift; SpellMisc Attributes_0 | aligned | - | - |
| X2 | **Racials of the published races.** Troll Berserking: +10% speed for 15 s, 3 min, flat (not health-scaled). The 1.48 casts a fight is a second use when the fight runs past 180 s. Orc Blood Fury: +10% attack power for 15 s, 2 min. | SpellEffect 697808, 698685; `core/racials.go:519` | aligned | - | - |
| X3 | **The hotfix push's SpellEquippedItems row.** The landing record lists 1 added row. Against the 11-02 copy, wago now has three new rows: Holy Shield 1239570 (needs a shield; lane tanks), Night Watchman's Torch and Enchant Boots - Speed. The push's row is one of these. None touches this group. | `scratchpad/r14/physical/wago/SpellEquippedItems.csv` against `r11/items/wago2/` | aligned (tanks: Holy Shield's shield row, FYI) | - | - |

## Before/after

Before: `r14/sb-before` (from `r14/site-before/sim`) and `r14/before.tsv`. After: `lane/r14-physical` e01d5596a with `wt/r14-physical`. 10,000 iterations; each cell is seed 1 / seed 50001.

| Spec, band | horde % | alliance % | seed 1 absolute (horde / alliance) | cause |
|---|---|---|---|---|
| Subtlety raid 60 | +0.53 / +0.51 | +0.46 / +0.53 | 544.0 → 546.9 / 544.7 → 547.2 | R1 |
| Subtlety bare 60 | +0.95 / +1.00 | - | 221.3 → 223.4 | R1 |
| Subtlety bare 50 | +0.65 / +0.65 | - | 138.4 → 139.3 | R1 |
| Subtlety bare 40 | +0.41 / +0.51 | - | 97.5 → 97.9 | R1 |
| Subtlety bare 20, 30 | 0 | - | 37.5, 60.4 unchanged (no Premeditation yet) | - |
| Every other spec in the group | 0 | 0 | no other engine change | - |

Outside the group, nothing moves: only a rogue with the Premeditation talent reads the changed file.

## Files to regenerate

- Site `sim/request` TestRotationLadder goldens: rogue-subtlety at bands 40 to 60 (R1). Not regenerated.
- dps-per-point.json for rogue-subtlety, then Data.lua (R1).
- No fork golden moved: `go test -tags=with_db ./sim/rogue/...` passes.

## Measure at refresh

- Subtlety picks and weights at bands 40 to 60. Premeditation now gives 2 more combo points every 2 minutes in the fight.

## For other lanes

- **guides:** `web/src/content/guides/rogue/subtlety.md` (lines 58 and 69) treats Premeditation as an opener only. In Forever it doesn't need Stealth: press it again when it comes off cooldown in the fight, with 3 or fewer combo points. The in-fight use comes from the engine: Premeditation is a major cooldown (`sim/rogue/premeditation.go`, used when the rogue holds 3 points or fewer), which the APL's first line casts, so no APL line is needed. Suggested player note (lane guides' wording): "Premeditation whenever it is ready and you have three combo points or fewer: it needs no Stealth in Forever, so it is two free points every 2 minutes. Spend them inside 20 seconds." The `notes` of `data/curated/apl/rogue-subtlety.json` say "the stealth opener (Stealth, Premeditation, Ambush)". The opener is unchanged, and the in-fight reuse is automatic (a cooldown, no APL line).
- **tanks:** the push's new SpellEquippedItems row is Holy Shield 1239570 requiring a shield (X3).

## Still open

- H4: Ranged Weapon Specialization on Serpent Sting (misc 127, ranged-weapon requirement); the engine excludes the sting.
- H7: Lacerating Strikes' refresh (snapshot replaced, the rest of the bleed lost).
- H10 and W5: Survival's queue lifting the off-hand miss penalty; normalized rage's off-hand factor and crit multiplier.
- W7: Deep Wounds' double multiplier. Blocked on logs (Warcraft Logs needs `gcloud auth login`).
- Wowhead's Forever tooltips still lag the push, except Sniper Shot's 590.

## Origin totals

1 confirmed bug (R1): old 1. R13 0, R12 0, landing 0, new-data 0.

Sources (read 2026-11-03): the wago.tools 1.60.1.70753 tables (SpellEffect, SpellMisc, SpellPower, SpellCooldowns, SpellAuraOptions, SpellLevels, SpellShapeshift, SpellCategories, SpellEquippedItems, ItemEffect, ItemSparse, ItemSet, ItemSetSpell) with the hotfix push layered; data/builds/1.60.1.70753/talents/{warrior,rogue,hunter}.json; the published BiS files; nether.wowhead.com Forever and Classic tooltips.
