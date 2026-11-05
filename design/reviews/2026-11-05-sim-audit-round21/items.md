# Sim audit round 21: items lane (shared item effects, sim/common)

**SHARED CODE.** The one engine commit (E 1bd3ad862, fork `lane/r21-items`) is in `sim/common` and moves every spec that wears Hand of Justice, Warblade of Caer Darrow, Storm Gauntlets or Blade of Hanna. No published BiS set wears HoJ, Warblade or Hanna; Storm Gauntlets sit on Elemental's band-60 sets and do nothing there (no melee). Every published number is unchanged. Two bare-60 weapon picks should flip at the next re-rank (Measure at refresh).

Leads came from the MythicSim tier-list comparison (`competitor/lanes/physical.md`, `hybrids.md`). Every claim was re-checked against the 70790 client rows (wago `ItemXItemEffect`, `ItemEffect`, `SpellAuraOptions`, `SpellEffect`, `SpellMisc`, `SpellCategories`, `SpellDuration`, `Spell` text, fetched for build 1.60.1.70790 on 2026-11-05) and Wowhead Forever's tooltips (`nether.wowhead.com/forever/tooltip/spell/<id>`, read 2026-11-05 ~04:40 UTC). `hotfix-push.json` carries no row for any of these spells or item effects. A tooltip plus the client table count as one source; where the client's own text formula settles it, that is said.

## Items

| # | finding | evidence and source | confirmed? | origin | fix |
|---|---|---|---|---|---|
| I1 | **Hand of Justice procs 2% a melee hit; Forever's is 1% against anything that is not a Dwarf.** The engine hard-coded `sim.Proc(0.02, "HandOfJustice")` (`sim/common/item_effects.go:2819`). The client's own description formula divides the proc chance by 3: the 3% in the aura row is the rate against Dwarves. A raid boss is never a Dwarf, so every hit rolls 1%. The 2 s cooldown, the melee-hit trigger (auto attacks and specials) and the one extra attack were already right. | 11815 → ItemEffect 99238 → 15600. `Spell.Description`: "${$h/3}% chance on Melee hit to gain $s1 extra attack. Attacks against Dwarves are $s2 times as likely to activate this effect." SpellAuraOptions 128409: ProcChance 3, ProcCategoryRecovery 2000, ProcTypeMask 20. SpellEffect 1325832 (index 1, aura 4): $s2 = 3. Wowhead Forever 15600: "1% chance on Melee hit to gain 1 extra attack... (Proc chance: 3%, 2s cooldown)". Test `TestHandOfJusticeProcsAtTheClientsNonDwarfRate` (failed first: no constant). | CONFIRMED-BUG (client text formula + aura row; single source family, but the formula is the client's own) | old (Era value from upstream 77a9302c1, 2024; HoJ is on no published set, so no round looked) | E 1bd3ad862 |
| I2 | **Warblade of Caer Darrow's Forever equip had no handler.** Forever gives it 22 to 33 Frost on every landed melee hit with the weapon. Added: 27.72 with variance 0.4, on every landed auto attack or melee special from the hand holding it. The spell is melee-defense with ALWAYS_HIT, so it never misses and crits on the melee table (double damage). It also suppresses procs on both sides, so it triggers nothing. | 13982 → ItemXItemEffect 166865 → ItemEffect 234705 (trigger 1, equip) → 1298500. SpellAuraOptions 246336: ProcChance 100, ProcTypeMask 20 (auto attacks + melee specials). Trigger 1298499: SpellEffect 1322194, Effect 2, base 27.72, Variance 0.4 (22.2 to 33.3); SpellMisc school 16 (Frost), Attributes_3 0x40070000 (suppress caster procs, suppress target procs, always hit); SpellCategories DefenseType 2 (melee). Wowhead Forever 1298500: "Melee attacks with this weapon deal 22 to 33 Frost damage." In a full Arms sim it lands on 167 hits per fight at 35% crit (the warrior's melee crit). Tests: `TestForeverMeleeEquipsRollTheClientsDamage`, `...ProcOnEveryLandedHit`, `TestWarbladeFrostNeverMissesAndCritsOnTheMeleeTable` (all failed first: spell not registered). | CONFIRMED-BUG | new-data (a Forever-only equip; the engine never had it) | E 1bd3ad862 |
| I3 | **Storm Gauntlets' Lightning Strike had no handler.** Added: 3 Nature on every landed melee hit. The spell is magic-defense with no always-hit attribute, so it rolls the spell hit and crit table. | 12632 → ItemEffect 99469 → 16615 "Add Lightning Dam - Weap 03". SpellAuraOptions 128107: ProcChance 100, ProcTypeMask 20. Trigger 16614 "Lightning Strike": SpellEffect 692412, Effect 2, base 3, Variance 0; school 8; DefenseType 1. Wowhead Forever 16615 and item 12632: "Adds 3 Lightning damage to your melee attacks." Same tests as I2. | CONFIRMED-BUG (trivial size) | old (Era item; no round looked) | E 1bd3ad862 |
| I4 | **Blade of Hanna's Forever equip had no handler.** Forever gives it Improving Form: a stack per landed auto attack; at 30 stacks, Perfected Form hits the target for 490 Physical and gives 20 to all stats for 20 sec, and the stacks clear. Added as the client states it: auto attacks only (the aura row's trigger mask is auto attacks, although the text says "each melee hit"), the 490 strike on the melee table and through armour, the stat buff on the caster. | 2801 → ItemEffect 238936 → 1325535 Improving Form. SpellAuraOptions 248750: ProcChance 101, ProcTypeMask 4 (auto attacks only). Stack aura 1325509: CumulativeAura 30, duration index 21 (until cancelled), effect 1362880 aura 542 misc 30 → 1325508. 1325508 Perfected Form: Effect 2 base 490 (Physical, DefenseType 2), aura 29 misc -1 base 20 (all stats), duration index 18 = 20 s, Effect 164 removing 1325509. Wowhead Forever 1325535 and item 2801: "Each melee hit grants you a stack of Improving Form. At 30 stacks perfect your form, inflicting 490 Physical damage to your current target and increasing all your stats by 20 for 20 sec." Tests `TestBladeOfHannaPerfectsItsFormEveryThirtyHits` (failed first: nothing registered), `TestBladeOfHannaStacksOnlyOnAutoAttacks`. | CONFIRMED-BUG | new-data (the round-5 hybrids and ranker lanes compared Blade of Hanna as a weapon without its equip; its 70707 hotfix moved it to 2.0 s and dropped its stats) | E 1bd3ad862 |
| I5 | Ironfoe (11684), for the record: the engine's 0.8 PPM is about 3.2% a swing on its 2.4 s speed. Forever's 1301046 carries ProcChance 6 with "Attacks against Orcs are 2 times as likely" ($s2 = 2); by HoJ's pattern the 6% is the Orc rate and a boss rolls 3%. The client text gives no `$h/2` formula, so this is inferred. Not on any published set. | SpellAuraOptions 246591 (ProcChance 6, 100 ms recovery, mask 20); SpellEffect 1325972 (aura 4, base 2); Wowhead Forever 1301046 "(Proc chance: 6%, 100ms cooldown)". | Aligned within 0.2 points (no change) | - | - |

## Survey: weapon and trinket equips with no engine handler, by published BiS use

Method: every chosen main hand, off hand, ranged and trinket in `$R/bis-frozen` (the frozen published BiS, all bands), asked of the linked engine itself (`core.HasItemEffect` / `HasWeaponEffect`, as `sim/cmd/leveling-bis/engine_effects.go` does), then the item's 70790 `ItemXItemEffect` rows with pure stat auras set aside (simdb carries those). Items with no client effect rows were checked on Wowhead Forever: the top 26 carry only stat equips (hit, crit, spell power, attack power).

| bands | item | unhandled effect | matters for DPS/heals/threat? |
|---|---|---|---|
| 40 | Uther's Strength 11302 | 8397, a holy-shield absorb when struck | No (defensive, when struck) |
| 14 | The Unstoppable Force 19323 | 23454 stun on hit | No (bosses are stun immune) |
| 6 | Cookie's Stirring Rod 5198 | cooking speed | No |
| 5+3+3+3+2+2 | Talisman of Arathor / Defiler's Talisman (20071, 20072, 21115-21121) | use: physical absorb | No for threat; a tank survival use only |
| 5 / 4 | Frostwolf / Stormpike Insignia 6 | use: recall | No |
| 5 | Orb of Soran'ruk 6898 | use: heal over time | No |
| 4 | Arena Grand Master 19024 | use: absorb; Battle Ring damage taken | No |
| 4 | Lifeforce Dirk 10750 | on hit: self heal | No |
| 4 / 3 | Staff of Westfall 2042, Forsaken Greataxe 251533 | zone run speed | No |
| 2 | Julie's Dagger 6660 | on hit: self heal | No |
| 1 | Satchel of Copper Bombs 285275 | thrown attacks explode | No (no thrown attack in the Protection rotation) |

Nothing the sim should model is missing on the published BiS. Shard of the Splithooves (10659, "+21 Attack Power against Demons") sits on ten caster bands as a filler trinket with no score; it gives a caster nothing either way.

## Before/after

10000 iterations, frozen published BiS (`$R/bis-frozen`), guide raid builds (`$R/rt`). Before = this site worktree on fork int/r21 8aa8e55fa; after = the same plus E 1bd3ad862.

| case | faction | seed 1 before → after | seed 50001 before → after | change |
|---|---|---|---|---|
| Elemental raid 60 (published, Storm Gauntlets) | alliance | 669.4 → 669.4 | 669.5 → 669.5 | 0.0% |
| Elemental raid 60 (published, Storm Gauntlets) | horde | 666.1 → 666.1 | 666.6 → 666.6 | 0.0% |
| Elemental bare 60 (published, Storm Gauntlets) | alliance | 240.0 → 240.0 | 240.2 → 240.2 | 0.0% |
| Arms raid 60, Warblade in the main hand | horde | 804.7 → 841.0 | 805.5 → 841.5 | +4.5% |
| Retribution raid 60, Blade of Hanna | horde | 696.6 → 713.3 | 695.9 → 712.4 | +2.4% |
| Fury raid 60, Hand of Justice in trinket 2 | horde | 918.4 → 912.4 | 918.9 → 912.5 | -0.7% |
| Enhancement raid 60, Storm Gauntlets | horde | 849.3 → 851.7 | 850.3 → 852.8 | +0.3% |

Neither raid pick flips: published Arms (Arcanite Champion) is 855.5 / 854.4 against Warblade's 841; published Retribution 797.0 / 796.6 against Hanna's 713.

Bare 60 (`-preset bare`, 10000 iterations, seeds 1 / 50001), published weapon against the new equips:

| spec | faction | published | candidate before | candidate after |
|---|---|---|---|---|
| Arms (The Unstoppable Force) vs Warblade | alliance | 318.8 / 318.7 | 298.4 / 298.5 | **320.9 / 320.9 (+0.7%)** |
| Arms vs Warblade | horde | 320.5 / 320.5 | 299.8 / 299.7 | **322.4 / 322.7 (+0.6%)** |
| Retribution (Forest Defender's Axe) vs Blade of Hanna | alliance | 292.9 / 292.8 | 288.4 / 288.2 | **295.6 / 295.5 (+0.9%)** |
| Retribution vs Blade of Hanna | horde | 295.4 / 295.2 | 293.5 / 293.8 | **301.0 / 301.1 (+1.9%)** |

## Files to regenerate

- Fork goldens outside `sim/common` (their test gear sets wear Hand of Justice): `sim/warrior/dps_warrior/TestP1DPSWarrior.results`, `sim/shaman/enhancement/TestEnhancement.results`. Both fail on the lane head, as expected; `sim/hunter`, `sim/rogue`, `sim/rogue/dps_rogue`, `sim/shaman` and `sim/` (movement e2e) pass.
- Site: nothing. No published set wears an item whose number moved.

## Measure at refresh

- Arms bare 60 should take Warblade of Caer Darrow over The Unstoppable Force (+0.6-0.7%), both factions.
- Retribution bare 60 should take Blade of Hanna over Forest Defender's Axe (+0.9% alliance, +1.9% horde).
- Hand of Justice's ranker value halves; it was on no published set, so no pick should move.

## For other lanes

- Warrior and paladin lanes: the two bare-60 weapon flips above land with the BiS refresh, not by hand.
- Coordinator: the MythicSim "+6% pre-raid items" figure counted HoJ at 2%; at 1% it is smaller.

## Still open

- Hand of Justice and Ironfoe proc rates on a level-60 log (any extra-attack count over known landed hits). Live logs: site API (one 20 Oct report) and foreverlogs.gg's public list (empty) last looked 05:01 UTC 5 Nov; WCL had no Forever zone at the platform lane's 03:02 UTC check.
- Whether Perfected Form's stat buff applies when its 490 strike misses or is dodged. The engine applies it either way (the aura is on the caster); a log of a dodged Perfected Form would settle it.
