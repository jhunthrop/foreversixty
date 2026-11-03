# Sim audit, round 12: physical group (2026-11-03)

This lane covers Arms, Fury, Protection, Assassination, Combat, Subtlety, Beast Mastery, Marksmanship and Survival (melee), on client 1.60.1.70753.

- **Baseline:** fork pin 2606bc05e with site c6bdc855c (the shared `r12/sb-before`, `r12/before.tsv`).
- **Branches:** fork fixes on `lane/r12-physical` (E), this report on `wt/r12-physical` (S). No site code changed.
- **Measurements:** `sim/cmd/spec-breakdown`, 10,000 iterations, seeds 1 and 50001, on the published sets.
- **Client rows:** wago.tools 1.60.1.70753 CSVs (round 11's copy, `scratchpad/r11/c70753/builds/1.60.1.70753/raw`). The 70753 hotfix push is empty: every table reads added 0, changed 0 (`data/builds/1.60.1.70753/hotfix-push.json`). Wowhead Forever tooltips read 2026-11-03; a tooltip and its client row count as one source.

Origin tags follow the brief. `new-data` is 70753.

## Assigned items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | **Poised was not simulated.** 70753 teaches every rogue Poised at 18, not only Subtlety. SkillLineAbility 59888 puts passive 1328825 on skill line 39 with class mask 8 (rogue) and acquire method 0 (trainer). Skill line 39 is the rogue's Subtlety school, the same line that teaches Stealth, so the trainables file's "Subtlety" label is the school, not the spec. The tooltip says: "While at maximum Energy, and out of combat, reduce the Energy cost of your next ability by 10." The passive is a dummy. It applies buff 1328826: no duration (index 21), one charge (SpellAuraOptions 249076), and one effect, SpellEffect 1368039: aura 107, SpellMod 14 (cost), -10, on family mask 255727518 / 2099456 / 67108864 / 0. Against SpellClassOptions set 8, that mask covers every rogue ability that costs Energy: Sinister Strike, Backstab, Gouge, Kick, Garrote, Ambush, Cheap Shot, Eviscerate, Slice and Dice, Expose Armor, Rupture, Hemorrhage, Ghostly Strike, Feint, Riposte, Blade Flurry, Mutilate and Venom (also Sap, Distract, Blind and Kidney Shot, which the engine does not model). A fight starts out of combat at full Energy, so the opener gets the discount: Garrote costs 40 (20 with Dirty Deeds) instead of 50. The mechanism is `core.RegisterPoised`, shared with the cat. The rogue's Energy abilities now have engine class masks, read through one client-family table (`rogueClassMasks`); the Tier 1 5-piece reads that table too. Tests: `TestPoisedTakesTenEnergyOffTheFirstAbility` and `TestPoisedIsNotKnownBelowEighteen` failed first. `TestPoisedReachesEveryEnergyAbility` checks that every Energy ability is in the mask and that nothing else is. | SkillLineAbility 59888; SpellEffect 1368038, 1368039; SpellMisc 874124; SpellAuraOptions 249076; SpellClassOptions set 8; https://nether.wowhead.com/forever/tooltip/spell/1328825 and /1328826 | CONFIRMED-BUG (client rows) | new-data | E 3322675af |
| A2 | **Thunder Clap's slow: the trained 20% is right.** SkillLineAbility 6073 to 6078 teach 6343, 8198, 8204, 8205, 11580 and 11581. Every rank has aura 319 at -20, a 6 s category cooldown, and durations of 10, 14, 18, 22, 26 and 30 s. The 461810 family (-10, 4 s) has no SkillLineAbility row, so no one trains it. The tooltip for 11581 reads "increasing the time between their attacks by 20% for 30 sec". The warrior's own Thunder Clap already used 20% and the per-rank durations. The slow multiplies the boss's attack speed by 1/1.2, so a Protection warrior in the tank harness takes fewer swings. That cuts damage taken and the rage it gives. The Protection APL keeps the slow up at every band: 10.2, 8.9 and 7.4 casts a fight at bands 40, 50 and 60, which matches each rank's duration. | SpellEffect 688690 (11581) and 1170519 (461810); SpellMisc; SpellCooldowns; https://nether.wowhead.com/forever/tooltip/spell/11581 | aligned | - | - |
| A2b | **The Thunder Clap raid debuff (`thunder_clap` in the request vocabulary) still had Era's 10% (15% improved) on 8205.** It is now 11581 at 20%, and 30% improved (Conqueror's 5-piece, +50%). Both values come from core constants, which the warrior's spell now reads too. No published preset carries this debuff, so no published number moves. Fork goldens: TestP1TankWarrior was regenerated. TestProtection (paladin) and TestFeralTank move too, because FullDebuffs carries the improved debuff; see Files to regenerate. Test `TestThunderClapRaidDebuffSlowsTwentyPercent` failed first. | `sim/core/debuffs.go:124` (upstream 3dfbe03c6, 2024-12-31) | CONFIRMED-BUG (client rows) | old (upstream; no round looked at it) | E 9291833a7 |
| A3 | **The fork's dps_rogue tests failed with "No item with id: 12783".** The fork loads its item database only under `-tags=with_db`, which CI passes and a bare `go test` does not. The Mutilate tests were built to avoid the database (synthetic weapons). The dagger check (56f97823e, 2026-10-31) then equipped Heartseeker and Scarlet Kris by item id, and the rotation test wore the whole prebis set. Without the tag the package panicked before any test ran. The four weapons are now a fixture `SimDatabase` on the player (`core.Character` adds a player's Database to the item table). The rotation test now wears the fixture mace and sword; it checks cast counts, which do not change (4.24 Ruptures either way). The package passes with and without the tag. The `rogue` package's own tests still need the tag for its gear sets, as CI runs them. | `sim/rogue/dps_rogue/mutilate_test.go`, `finisher_rotation_test.go`, `mutilate_forever_test.go` | CONFIRMED (test fixture, not a sim bug) | old | E 8497fe603 |
| A4 | **Aimed Shot is baseline, consistently.** SkillLineAbility 10956 and 11870 to 11875 teach 19434 and 20900 to 20904 on skill line 163 at 20/28/36/44/52/60. Forever's trait tree (talents/hunter.json) has no Aimed Shot node; the legacy Talent row 1345 is not in the tree. The engine registers it by level alone (`sim/hunter/aimed_shot.go:76`), and the BM and MM APLs cast rank 6. The cast time is 2.0 s (SpellCastTimes for 19434 and 20904) and the engine uses 2.0 s. One mismatch, for the trainables lane: `trainables/hunter.json` gives Aimed Shot `cast_time_ms: 3000`. | SkillLineAbility; SpellMisc / SpellCastTimes; talents/hunter.json | aligned (trainables field flagged) | - | - |
| A5 | **The orc Command rows: no new hunter consequence.** The hunter row 20576 (set 9) is unchanged from 70707: SpellEffect 697944, aura 107, op 8 (all effects), +5 on family 33554432. That bit is Tamed Pet Passive 8875 (aura 79, all schools), which makes it +5% to all pet damage, and the engine applies exactly that (`sim/core/racials.go:245`). The new row is the warlock's: 20575 gains SpellEffect 697943, aura 108, op 0 (damage), +5 on family 12288 of set 5, which is Firebolt (4096) and Lash of Pain (8192). Reported to the auras and warlock lanes. | SpellEffect 697942 to 697944, 698201; SpellClassOptions 46429, 46430, 46483 | aligned for hunters | - | - |

## Arms, Fury, Protection

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| W1 | Cost per cast against SpellPower, from the raid breakdowns. Arms: Mortal Strike 30, Slam 15, Rend 10, Overpower 5, Hamstring 10, Battle Shout 10, Heroic Strike 15 (no Improved Heroic Strike in the Arms build). Fury: Bloodthirst 30, Whirlwind 25, Death Wish 10, Heroic Strike 12 (Improved Heroic Strike -3). Execute: 15 less Improved Execute, plus all extra rage (28.3 and 23.6 a cast). Cooldowns: Mortal Strike and Bloodthirst 6 s, Whirlwind 10 s, Slam 18 s, Death Wish 3 min. All match SpellPower and SpellCooldowns. | `bd/warrior-*.txt`; SpellPower; SpellCooldowns | aligned | - | - |
| W2 | Thunder Clap (A2). Protection's Defensive Stance, Last Stand and Shield Wall lines and its openers (Bloodrage, Battle Shout before the pull) were read, not changed. The landing's Arms Charge at -0.5 s belongs to the regmelee lane. | `data/curated/apl/warrior-*.json` | aligned | - | - |

## Assassination, Combat, Subtlety

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | Poised (A1). | | CONFIRMED-BUG | new-data | E 3322675af |
| R2 | Cost per cast: Mutilate 60, Backstab 60, Ghostly Strike 40, Hemorrhage 35 (31.3 on average with Thousand Cuts), Rupture and Slice and Dice 25, Eviscerate 35 (25 with Flawless Execution in the Assassination and Combat builds), Blade Flurry 25, and Garrote 50 less Dirty Deeds 20 less Poised 10. Relentless Strikes ("20% chance per Combo Point to restore 25 Energy") returns 333.5 Energy a fight on 66.5 combo points spent, against 332.5 expected. Ghostly Strike's dodge aura has 33.0% uptime on 8.65 casts, which is the client's 7 s; its 20 s cooldown fits 8.65 casts in 180 s. Eviscerate is 48 to 144 plus 151 a point (SpellEffect 689090: 96, variance 1, 151 per resource), and Improved Eviscerate and Aggression add together, as aura 108 percent mods do. Thousand Cuts is -3 a stack, five stacks. Frozen Heart of the Mountain (249469) is Forever's own trinket: equip +0.9% hit; its use, Chthonic Power, is Frost and Shadow spell damage and does nothing for a rogue. The Major Frenzy Potion is up 33% of the time (two 30 s potions). | `bd/rogue-*.txt`; SpellPower; SpellEffect 689090, 1252390; talents/rogue.json | aligned | - | - |
| R3 | Openers: Stealth at -3 s, Premeditation at -2 s for Subtlety (it costs no Energy and is not in Poised's mask, so Poised waits for the Garrote), Garrote at 0 s. Each spends Poised as the client says. | `data/curated/apl/rogue-*.json` | aligned | - | - |

## Beast Mastery, Marksmanship, Survival

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | **Every engine missile flew at 24 yards a second.** SpellMisc Speed says: Auto Shot (75, also Shoot Bow, Shoot Gun and Shoot Crossbow) 40; Aimed Shot, Arcane Shot and Serpent Sting 40; Multi-Shot 30; Sniper Shot 60; Summon Hawk 35; the three traps 0; wand Shoot (5019) 20, which reaches every wand user. Each now reads its own row (`sim/hunter/missile_speeds.go`, core's ranged auto and Shoot). The warlock lane found it. A shot lands sooner or later; the DPS moves by noise only (-0.2 to +0.2%, table below). Test `TestHunterShotsTravelAtTheClientsSpeed` and the Shoot check in `TestRegisterShootSpellUsesTheEquippedWand` failed first. | SpellMisc rows for 75, 2480, 7918, 7919, 2764, 19434, 20904, 3044, 14287, 1978, 25295, 2643, 1310687, 1310786, 1293241, 1293527, 1499, 13795, 13813, 5019 | CONFIRMED-BUG (client rows) | old (upstream's default; the wand from 3f5ea92c1, 2026-09-28) | E 8a54e923e |
| H2 | Mana per cast: Efficiency is "Shots, Stings, and melee abilities by 15%". BM, with 4 points, pays 88% on Arcane Shot, Serpent Sting, Aimed Shot and Multi-Shot; MM, with 5 points, pays 85%. Survival takes none and pays Resourcefulness's 40% on Raptor Strike (40), Mongoose Bite (26), Lacerate (38), Wing Clip (32) and Immolation Trap (98). Strider Kick is 5.81% of base mana × 0.4 = 40. Multi-Shot is 13.9% base mana and Bestial Wrath 12%. Cooldowns: Arcane, Aimed and Multi 6 s, Mongoose Bite 5 s, Strider Kick 8 s, traps 30 s. | `bd/hunter-*.txt`; SpellPower; SpellCooldowns; talents/hunter.json | aligned | - | - |
| H3 | Rapid Fire for melee Survival: 3045 carries aura 140 (ranged) and aura 319 (melee) at 40, and the tooltip reads "ranged and melee attack speed by 40%". The engine hastes both. Expose Prey ("enemies afflicted by Hunter's Mark"): any Hunter's Mark counts, so the raid's mark works. Survival casts no mark of its own. With one added before the pull, bare 30 and 60 do not gain (80.3 against 80.4; 293.8 against 295.0), so the rotation stays. | SpellEffect 679813, 1339717; tooltips 3045, 1310532; `sim/hunter/talents.go:776` | aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | The 70753 hotfix push is empty, so the shipped tables are the truth this round. | `hotfix-push.json` search_totals | aligned | - | - |

## Before/after

**Before:** `r12/sb-before` (2606bc05e, site c6bdc855c) and `r12/before.tsv`. **After:** `lane/r12-physical` 8a54e923e with `wt/r12-physical`. Each cell is seed 1 / seed 50001.

| Spec, band | alliance % | horde % | seed 1 alliance / horde | cause |
|---|---|---|---|---|
| Arms, Fury, Protection 60 raid | 0.00 / 0.00 | 0.00 / 0.00 | 719.5 / 731.7 (Arms), 754.9 / 757.2 (Fury), 1080.0 / 1069.2 (Protection, threat/s) unchanged | - |
| Assassination 60 raid | +0.21 / +0.16 | +0.18 / +0.08 | 625.1 → 626.4 / 621.0 → 622.1 | A1 |
| Combat 60 raid | +0.24 / +0.22 | +0.22 / +0.24 | 625.5 → 627.0 / 626.8 → 628.2 | A1 |
| Subtlety 60 raid | +0.22 / +0.18 | +0.22 / +0.18 | 549.6 → 550.8 / 549.0 → 550.2 | A1 |
| Beast Mastery 60 raid | -0.06 / +0.05 | -0.05 / -0.03 | 623.0 → 622.6 / 626.9 → 626.6 | H1 (noise) |
| Marksmanship 60 raid | +0.05 / -0.03 | +0.08 / -0.03 | 610.7 → 611.0 / 615.7 → 616.2 | H1 (noise) |
| Survival 60 raid | 0.00 / 0.00 | 0.00 / 0.00 | 741.1 / 744.7 unchanged | - |
| Frost mage, Affliction, Shadow priest, cat 60 raid (outside the group) | 0.00 / 0.00 | 0.00 / 0.00 | unchanged | - |

Bare bands, horde, seed 1 / seed 50001:

| spec | 20 | 30 | 40 | 50 | 60 | cause |
|---|---|---|---|---|---|---|
| Assassination | 0.0 / +0.3 | 0.0 / +0.2 | +0.3 / +0.2 | +0.3 / +0.3 | +0.3 / +0.3 | A1 |
| Combat | +0.3 / +0.3 | +0.2 / +0.2 | +0.1 / +0.2 | +0.2 / +0.3 | +0.2 / +0.2 | A1 |
| Subtlety | 0.0 / 0.0 | +0.2 / 0.0 | +0.2 / +0.3 | +0.3 / +0.3 | +0.3 / +0.3 | A1 |
| Beast Mastery | -0.2 / -0.2 | - | -0.1 / -0.2 | - | 0.0 / -0.1 | H1 (noise) |
| Marksmanship | +0.2 / +0.2 | - | -0.1 / -0.1 | - | -0.1 / -0.1 | H1 (noise) |
| Survival | 0.0 / 0.0 | - | 0.0 / 0.0 | - | 0.0 / 0.0 | H1 |
| Shadow priest (wand) | -0.3 / 0.0 | - | 0.0 / 0.0 | - | 0.0 / 0.0 | H1 (wand) |
| Affliction, Frost mage (wand) | 0.0 / 0.0 | - | 0.0 / 0.0 | - | 0.0 / 0.0 | H1 (wand) |

## Files to regenerate

- Fork goldens moved by E 9291833a7: `sim/paladin/protection/TestProtection.results` and `sim/druid/tank/TestFeralTank.results`. FullDebuffs carries the improved Thunder Clap debuff, which is now 30% (it was 15%), so DTPS and TPS fall. I regenerated only `sim/warrior/tank_warrior/TestP1TankWarrior.results`.
- Fork golden moved by E 8a54e923e: `sim/mage/TestP1Mage.results`, by a hundredth of a DPS, from wand Shoot's flight time.
- Site TestRotationLadder rows for the three rogue specs (Poised) and the hunters (missile speeds) move at the pin. Not regenerated.

## Measure at refresh

- Rogue gear and talent picks under Poised (+0.2 to +0.3%, uniform; no pick is expected to move).

## For other lanes

- **Hybrids (cat Poised):** use `core.RegisterPoised` (`sim/core/poised.go`, E 3322675af) and do not add a second mechanism. Notes are in `scratchpad/r12/physical/POISED-NOTES.txt`. The cat's passive is 1328829 at **level 32** (SkillLineAbility 59889, skill line 134, class mask 1024). `sim/druid/constants_auto_gen.go`'s `PoisedSpellId` / `PoisedLevel` are the buff row (1328827, level 1), not the passive, so do not gate on them. The buff's mask (8622080, 536870912, 0, 0; SpellEffect 1368040) resolves in set 7 to Claw, Rake, Shred, Ferocious Bite, Rip, Pounce, Ravage and Cower. Tiger's Fury is not in it.
- **Auras / warlock:** 70753's new Command effect is the warlock's (20575, SpellEffect 697943: +5% damage, aura 108 op 0, on Firebolt and Lash of Pain), on top of the +5% all pet damage that both classes already get through Tamed Pet Passive. Hunters see nothing new (A5). The engine's Command (`racials.go:245`) gives 5% to all pet damage. Whether the warlock's imp and succubus spells should take a second 5% belongs to those lanes.
- **Auras:** E 9291833a7 changes the core raid debuff `ThunderClapAura` (`sim/core/debuffs.go`). If you also touched debuffs.go, merge around the Thunder Clap block.
- **Healprofile:** the sub-60 heal profile builds the boss's melee from creature_template with no Thunder Clap slow. A warrior-tanked group would take about 17% less melee. Note only; groups vary.
- **Trainables:** `trainables/hunter.json` gives Aimed Shot `cast_time_ms: 3000`; the client's SpellCastTimes for 19434 to 20904 is 2000.
- **Warlock (relay):** your missile-speed finding is confirmed and fixed for the hunter shots, Auto Shot and wand Shoot (E 8a54e923e). The Shadow Bolt and Incinerate travel are yours.
- **Integrator:** E 8a54e923e moves wand users outside this group (priests, mages, warlocks) by noise only (table above). E 9291833a7 moves no published number. E 3322675af moves only the three rogue specs.

## Still open

- Raptor Strike's queue and the off-hand penalty, normalized rage's off-hand factor and crit, Lacerating Strikes' refresh, and the one-a-minute proc rate (round 11 H2, A2 to A4). These are carried. No new logs exist.
- Marksmanship's held +0.9% build (round 11 A1).
- Carried: Eviscerate's and Rupture's attack power, Venom, Combat's Slice and Dice gate, the Hawk proc mask, Serpent Sting's and Arcane Shot's spell power, threat values, and weapon-proc rates.

Sources (read 2026-11-03): the wago.tools 1.60.1.70753 CSVs above; https://nether.wowhead.com/forever/tooltip/spell/{1328825,1328826,1328827,1328829,11581,3045,1310532,11300,1249110} and /item/249469.
