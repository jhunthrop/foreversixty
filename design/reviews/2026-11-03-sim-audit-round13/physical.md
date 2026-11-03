# Sim audit, round 13: physical group, fresh re-audit (2026-11-03)

This lane covers Arms, Fury, Protection, Assassination, Combat, Subtlety, Beast Mastery, Marksmanship and Survival (melee), on client 1.60.1.70753. It re-audits them from scratch. Round 12's own changes (Poised, the Arms opener, Thunder Clap, missile speeds, Arcane Shot's spell power, Juju Might, the Fury build) belong to the regmelee lane.

- **Baseline:** fork int/r12 076a5a517 with site int/r12 93c5d2510 (`r13/sb-before`, `r13/before.tsv`).
- **Branches:** fork `lane/r13-physical` (E), report on `wt/r13-physical` (S). No site code changed.
- **Client rows:** the wago.tools 1.60.1.70753 CSVs (`scratchpad/r11/c70753/builds/1.60.1.70753/raw`), read with a lane helper (`r13/physical/tools/sp.py`). A tooltip and its client row count as one source.
- **Hotfix push.** A hunter hotfix push went live at 14:14 UTC on 2026-11-03: Aimed Shot, Arcane Shot, Serpent Sting, Trueshot Aura, the low ranks of Aspect of the Hawk, Sniper Shot, and Lacerating Strikes' attributes. The `hotfix` lane owns it. This lane changed no hunter spell amount and no hunter APL.

**Result: 1 confirmed bug, fixed (old).** Every other spell and talent checked below is aligned or carried open.

## Arms, Fury, Protection

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| W1 | **Arms abilities against the client.** Overpower 11585: weapon damage +35, normalized (effect 121), 5 rage, 5 s, 50% more crit from Improved Overpower. In the breakdown it crits 77.4%. Mortal Strike 21553: +160 normalized, 30 rage, 6 s. Slam 11605: +87, 15 rage, a 1.5 s cast and an 18 s cooldown. Improved Slam 12862 takes 0.5 s off the cast and the GCD (misc 10 and 21) and 3 s off the cooldown (misc 11), and keeps the swing; the engine matches (`slam.go`). Rend 11574: 21 a tick for 7 ticks (147), 10 rage. Hamstring 7373: 45 flat, 10 rage; it is woven only with Windfury Totem up and stops at 25%. Execute 20662: 600, plus 15 per point of extra rage. | SpellEffect 688815, 697710, 688660, 689691, 685506, 696172; SpellPower; SpellCooldowns; `bd/warrior-arms-raid60.txt` | aligned | - | - |
| W2 | **Bloodthrill 5/5.** The tree text is 4% a point; the engine uses 20%, on landed main-hand attacks against your own Rend. The window is 5 s (1282733 DurationIndex 28). It procs 29.4 times a fight. The estimate is 168 landed main-hand attacks × 84% Rend uptime × 20% = 28.2. | talents/warrior.json node 110524; SpellMisc 1282733; `bloodthrill.go` | aligned | - | - |
| W3 | **Weaponmaster 5/5.** Axe or polearm: +5% crit. Mace or staff: 15% of the target's armor ignored, as a share of the armor (`ArmorIgnorePercent`). Sword: a 5% extra attack. These are 1290261's three effects (5, 15, 5); the engine applies 1%, 3% and 1% a point. Arms raid wields Arcanite Champion (a sword). | SpellEffect 1309067 to 1309069; `warrior/talents.go:405` | aligned | - | - |
| W4 | **Fury talents.** Bloodthirst 23894: 48 plus 45% of attack power, 30 rage, 6 s (`bloodthirst.go:30`). Death Wish 12328: 10 rage, 3 minutes, +20% Physical (aura 79), +5% damage taken (aura 87). Flurry: 25% for 3 swings. Unbridled Wrath 12322: proc type 4 (melee auto only), 60%, 1 rage; the engine procs on white hits only. Dual Wield Specialization 23584: +25% off-hand damage, +50% off-hand rage. Furious Precision: +10% off-hand hit. Improved Execute: -5. Cruelty: +5% melee crit (aura 52). Enrage: 2% a point, 30% on hits taken, so a raid DPS warrior gets none. | SpellEffect 701080, 1339740, 690203 to 690205, 689427, 689439, 699937; SpellAuraOptions 12322 | aligned | - | - |
| W5 | **Protection talents.** Shield Specialization: 5% block, and 5 rage on every block. Master of Defense: 5 rage on a dodge or parry with a shield. Defiance: +15% threat. Focused Rage: -3 on offensive abilities (`SpellFlagOffensive`). Bastion: +10% Physical with a shield. Improved Revenge: 20% a point. Anticipation: 4 Defense a point. Improved Thunder Clap: -2 rage a point. Shield Slam: 421 to 439 plus block value. Thunder Clap was round 12's (regmelee). | talents/warrior.json; `warrior/talents.go`, `talents_protection.go` | aligned | - | - |
| W6 | **Deep Wounds.** The engine adds the average weapon damage times the attacker multiplier, and then applies the attacker multiplier again on every tick (`deep_wounds.go`, "Double dips on attackers mods", upstream's Season of Discovery note). Rounds 2 to 11 accepted it. Nothing for Forever or Era settles whether the bleed's snapshot double-dips. | `sim/warrior/deep_wounds.go` | OPEN (a log of Deep Wounds ticks with Death Wish on and off would settle it) | - | - |

## Assassination, Combat, Subtlety

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | **Serrated Blades ignored a flat 126 armor, not 9% of the target's armor.** The client row 14171 is aura 280, "ignore 9% of your target's Armor" (SpellEffect 690915, 9 at rank 3). That is a share of the target's current armor. The engine added 9 × 13.99 to the flat ArmorPenetration stat; 13.99 is a TBC rating-per-percent constant. The armor formula then subtracts that as 126 flat armor. Against the raid preset's debuffed boss (Sunder Armor, Faerie Fire and Curse of Recklessness leave about 470 armor), 126 is far more than 9%. Against an undebuffed level-63 target, 126 is far less than 9%. Round 4 checked the percent per rank and missed the shape. The fix writes `PseudoStats.ArmorIgnorePercent`, as the warrior's Weaponmaster mace clause already does. The rewritten `TestSerratedBladesIgnoresThreePercentOfArmorPerRank` failed first (0 against 0.09, and 125.91 flat). | SpellEffect 690915 (aura 280); `sim/rogue/talents.go:100`; `core/spell_resistances.go:84` | CONFIRMED-BUG (client row) | old E e4b87c1e1 (2026-10-07, which corrected the per-rank percent but kept the rating conversion) | E b08a1c83c |
| R2 | **Hack and Slash's mace clause had the same shape.** "Mace: Your attacks ignore 15% of your target's armor" (13960 effect 1309137, 15) gave 210 flat armor. It now ignores 15% of the target's armor. No published rogue band wields a mace (all daggers), so no published number moves. Test `TestHackAndSlashMaceIgnoresArmorPercent` failed first (209.85 flat). Counted with R1: same commit, same cause. | SpellEffect 1309137; bis/rogue-*.json weapons | CONFIRMED-BUG (part of R1) | old (as R1) | E b08a1c83c |
| R3 | **Hack and Slash's other clauses.** Dagger or fist: +5% melee crit on the dagger hand only. Axe or sword: a 5% extra attack with a 200 ms cooldown (SpellAuraOptions: ProcCategoryRecovery 200, proc type 20). Both match. | SpellAuraOptions 13960; `rogue/talents.go:635` | aligned | - | - |
| R4 | **Combo points and Energy.** Puncturing Wounds 1224716: Backstab +30% crit, Mutilate +15%, and a 45% extra combo point on Backstab (15.6 points against 34.6 Backstab hits × 0.45 = 15.6). Seal Fate 5/5: at most one extra point per Mutilate cast (24.5 against 36.5 casts × P(either hand crits) 0.678 = 24.7). Ruthlessness: 60% (12.5 points on 21.3 finishers). Relentless Strikes: 295.6 Energy against 305 expected for Combat's finishers by combo point. Revitalize (23863, Rune of the Guard Captain): 4% on melee for 30 Energy. Weapon Expertise: dodge 4.5% against 6.4% without it. Precision: +3% hit. | SpellEffect 1213532, 1340439, 1348793, 701497, 700855; `bd/rogue-*-raid60.txt` | aligned | - | - |
| R5 | **Poisons.** Instant Poison VI procs on 20% of off-hand hits for Combat (23.1 of 115.6) and on 30% with Improved Poisons 5/5 for Assassination. Its 8 to 11% miss is the spell table. Deadly Poison IV, not V, at 60: V (25347) needs the Ahn'Qiraj book (round 6 R1; `2026-10-07-aq-ranks-off.md`). | SkillLineAbility 13240 (25347, poisons line 40); SpellItemEnchantment | aligned | - | - |
| R6 | **Talent texts against the engine.** Opportunity +10% (Backstab, Garrote, Ambush, Mutilate), Lethality +20% crit damage, Improved Eviscerate +20%, Aggression +6%, Flawless Execution -10 Energy, Thousand Cuts -3 a stack, Quietus +10% below 35%, Ghostly Strike 180% with a dagger, Hemorrhage 145% with a dagger and a 15% Rupture amplifier. Round 12 checked the openers and Poised. | talents/rogue.json; `rogue/*.go` | aligned | - | - |

## Beast Mastery, Marksmanship, Survival

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | **Trueshot Aura in the raid preset is ranged attack power only.** Every rank is aura 124, ranged only. The engine gives no melee attack power (`core/buffs.go:1405`). The push's new amount (100) belongs to the hotfix lane. | SpellEffect 1323534, 697683 | aligned (the amount is the hotfix lane's) | - | - |
| H2 | **The pet is a cat with a 1.2 s swing** (`sim/request/request.go:688`). It has 1.25 happiness and family scalars. Frenzy, Bestial Wrath, Unleashed Fury, Ferocity and Focused Fire are in the build. Classic pets keep the swing speed of the beast that was tamed. The 1.2 s is a choice of tame (common for cats), not a client value. | `hunter/pet.go:42` | aligned (a modelling choice; noted) | - | - |
| H3 | **Survival dual-wield table.** Specials (Raptor Strike, Mongoose Bite, Strider Kick, Wing Clip) take no dual-wield miss, which is right for yellow attacks. Main-hand white misses 18.4% and off-hand white 15.6%. The gap comes from Raptor Strike replacing main-hand swings in the queue, which is round 11's open "Raptor Strike's queue" item. | `bd/hunter-survival-raid60.txt` | OPEN (carried) | - | - |
| H4 | **Armor debuffs the hunters and melee meet in the raid preset.** In Forever every armor debuff is aura 674 (misc 1). Sunder Armor 11597: 450 a stack. Faerie Fire 9907: 505. Curse of Recklessness 11717: 505, with no attack power for the target. All three match `core/raid_debuff_ranks.go`. | SpellEffect 688831, 687337, 690026 | aligned | - | - |
| H5 | Hunter shots, stings and Aimed Shot's amounts were not re-audited. They are the hotfix lane's this round. | hotfix lane message, 2026-11-03 | not audited (handed off) | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Before the fix, flat `stats.ArmorPenetration` had one percent-based writer, the rogue (R1, R2). The only other writer is an item stack effect (`common/item_effects.go:3630`, 200 a stack), which is flat by its text. Percent armor ignore now has one path, `PseudoStats.ArmorIgnorePercent`, for warrior and rogue. | grep of `stats.ArmorPenetration` | aligned after R1 | - | - |

## Before/after

Before: `r13/sb-before` (int/r12) and `r13/before.tsv`. After: `lane/r13-physical` b08a1c83c with `wt/r13-physical`. 10,000 iterations; each cell is seed 1 / seed 50001.

| Spec, band | alliance % | horde % | seed 1 alliance / horde | cause |
|---|---|---|---|---|
| Subtlety raid 60 | -1.11 / -1.13 | -1.13 / -1.11 | 550.8 → 544.7 / 550.2 → 544.0 | R1 |
| Subtlety bare 60 | - | +1.47 / +1.47 | 218.1 → 221.3 (horde) | R1 |
| Subtlety bare 50 | - | +1.32 / +1.24 | 136.6 → 138.4 | R1 |
| Subtlety bare 40 | - | +1.04 / +1.04 | 96.5 → 97.5 | R1 |
| Subtlety bare 30 | - | +0.50 / +0.50 | 60.1 → 60.4 | R1 |
| Assassination, Combat (every band) | 0 | 0 | unchanged: no Serrated Blades in their builds, no mace | - |
| Warriors, hunters | 0 | 0 | no change in their engine files | - |

Below 30, the Subtlety build has no Serrated Blades. The raid loss is real: with Sunder, Faerie Fire and Curse of Recklessness up, a boss has about 470 armor left, so 9% is about 42 armor, not 126.

## Files to regenerate

- Site `sim/request` TestRotationLadder goldens: rogue-subtlety at bands 30 to 60 (R1). Not regenerated.
- No fork golden moved: `go test -tags=with_db ./sim/rogue/...` passes.

## Measure at refresh

- Subtlety picks and weights at every band. Armor penetration is now worth less at raid and more at bare bands, so any armor-ignore item or the Serrated Blades points may move.

## For other lanes

- **hotfix:** this lane changed no hunter spell amount and no hunter APL. Trueshot Aura's 100 is yours (H1): `core.TrueshotAura` reads `TrueshotAuraRanks`, ranged attack power only. Shots and stings were left to you (H5).
- **Integrator:** E b08a1c83c moves only rogue-subtlety (the mace branch moves no published set).

## Still open

- W6: Deep Wounds' double-dipped attacker multiplier (upstream's Season of Discovery behaviour); a log would settle it.
- H3: Raptor Strike's queue and the white miss split (carried from round 11), plus the off-hand penalty, normalized rage's off-hand factor and crit, Lacerating Strikes' refresh and the one-a-minute proc rate (carried).
- Marksmanship's held +0.9% build: not re-run this round. The hotfix lane re-tunes hunter builds on the new values, so the round-11 measurement is stale.

## Origin totals

1 confirmed bug (R1 and R2, one cause): old 1. R12 0, landing 0, R9 to R11 0, new-data 0.

Sources (read 2026-11-03): the wago.tools 1.60.1.70753 CSVs (SpellEffect, SpellMisc, SpellPower, SpellCooldowns, SpellAuraOptions, SpellLevels, SkillLineAbility); data/builds/1.60.1.70753/talents/{warrior,rogue,hunter}.json; the published BiS files.
