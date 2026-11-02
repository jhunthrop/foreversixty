# Sim audit round 10: crit lane

The owner's ask: "pay extra attention to anything the round-9 crit split could have broken: every effect that reads crit, the Succubus, procs, and talents that add crit to one school." This lane owns the engine-wide sweep and the crit plumbing.

Engine branch `lane/r10-crit` (fork, from 4d3129cae). Site branch `wt/r10-crit` (this report only). Client 1.60.1.70707 with the hotfix push layered in (raw tables copied from the round-9 physical lane's layered `raw/`; `SpellEquippedItems` fetched from wago.tools 2026-11-02).

## How the sweep was done

1. **A dynamic trace of every crit roll.** A temporary hook (never committed) in the three crit functions (`PhysicalCritChance`, `SpellCritChance`, `HealingCritChance`, `sim/core/spell_result.go`) logged, once per spell: the unit (player or pet), spell id, which table the roll read, school, declared DefenseType, ProcMask and the crit chance with its parts (every-table `Crit`, `MeleeCrit`, `SpellCrit`, the spell's `BonusCritRating`, the school's crit). It ran every published spec at band 60 raid and bands 20 to 50 bare (horde, 300 iterations): 434 distinct (spell, table) rows. Each row was joined to the client's `SpellCategories.DefenseType` (1 magic rolls spell crit; 2 melee and 3 ranged roll melee and ranged crit). Script: `scratchpad/r10/crit/py/join.py`, output `join.txt`.
2. **A static scan** of every engine `SpellConfig` with a literal spell id (`py/staticscan.py`), the weapon-proc table (`sim/common/weapon_proc_rolls.go`, 30 rows) and every `itemhelpers.CreateWeapon*ProcDamage` call (10), each against the client DefenseType.
3. **Every crit source by grep**: every write of `stats.Crit`, `stats.MeleeCrit`, `stats.SpellCrit`, `SchoolBonusCritChance`, `SchoolCritTakenChance`, `BonusCritRating` sources, and every "merged from MeleeCrit + SpellCrit" comment the pre-split code left (22 sites).
4. **Every talent rank** of all nine classes whose client effects touch crit (auras 52, 57, 71, 290, 163, and 107/108 on misc 7 crit chance and 15 crit damage), with triggered spells followed (`py/talentcrit.py`, 41 talents).
5. **Every set bonus** the published BiS wears (every spec, every band) against its client spells (`py/sets.py`). Only Stormshroud 2/3, Devilsaur 2 and Ironfeather 2 are active; none touches crit except Stormshroud's Lightning proc (18980, DefenseType 1, spell table: aligned).
6. **The beta logs** (foreverlogs.gg public reports, levels 12 to 35, the round-9 lane's cached downloads plus roster snapshots with Agility, Intellect and talent points): per-class and per-talent-tree crit on each table, pets, totems.

## The fixes

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| C1 | **Conviction reached the Holy damage spells.** The engine gave its 6% (3 ranks) to every table, so Consecration, Exorcism, Holy Shock and Holy Wrath got it. In Forever it reaches weapon attacks, not those spells. Heals keep it (see C1b). | Client: 20117 is aura 290 but carries a SpellEquippedItems row (180908: item class 2, subclass mask 173555, the melee weapons), Cruelty's restriction. Logs: 105 paladins with snapshots, levels 19 to 35. Retribution paladins (5+ points; Conviction is tier 1) crit Consecration 6.48% of 84,212 ticks; the rest 8.05% of 53,536 (SE 0.1). Their melee crit 10.1% against 6.2%. If Conviction reached Consecration, Retribution's would sit about 3 to 6 points above the rest; it sits 1.6 below. The sim at band 30 bare gave Retribution's Consecration 11.3% (6 of it Conviction); the band-30 logs show 5 to 7.5%. Test `TestConvictionIsWeaponCritAndHealCritOnly` (first `TestConvictionIsWeaponCritOnly`) failed first (6.00 spell crit). | CONFIRMED-BUG | E fc8c92a4e, golden E 4eb6a73fb, E 03ec2ecd8 |
| C1b | **Conviction on heals stays.** The logs hold one Retribution paladin with heals (154 Flash of Light, 7.8%), too few to judge, and Thundering Strikes, the same aura 290 with the same weapon row, does reach a shaman's heals there (Healing Wave 8.8% of 1,433 for Enhancement against 5.0% of 966). So heals keep the 6% as the spell's own crit bonus. Without this the first commit cut Holy paladin's published HPS by 5.6% on no heal evidence. | `py/a290.py`, shaman per-spell tally | OPEN (heals; kept as before) | E 03ec2ecd8 |
| C2 | **Gloves of Holy Might's smite rolled the spell table.** Its client DefenseType is 3 (ranged), like Hammer of Wrath's. It now rolls the ranged table: melee and ranged crit and the physical (double) crit damage. Survival wears the gloves at band 40. | `SpellCategories` 1325229 DefenseType 3, school holy. `TestGlovesOfHolyMightRollsTheRangedTable` failed first (largest crit 184.6, a spell crit; want about 220). | CONFIRMED-BUG | E 495779597 |
| C3 | **A school's spell crit did not reach that school's heals.** `HealingCritChance` read the spell table and the spell's own bonus but not `SchoolBonusCritChance`. Vestments of Prophecy 5-piece ("Holy spells", 21092) and The Ten Storms 5-piece ("Nature spells", 23570), both aura 57, are healer sets, and their crit missed every heal. Neither set is in a published band. | `sim/core/spell_result.go` `healingCritRating`; `sim/priest/item_sets_pve.go:22`, `sim/shaman/item_sets_pve.go:53`. `TestSchoolSpellCritReachesThatSchoolsHeals` failed first (healing +0, want +2). | CONFIRMED-BUG (no published number) | E b4546cdb6 |
| C4 | **Emerald Dragon Whelp (Dragon's Call) merged its crit.** The pre-split code wrote melee 4.5% and spell 13% as one 13% on every table, so its melee crit 13%. Each is on its own table again. Not in a published band. | `sim/common/guardians/emerald_dragon_whelp.go` "merged from MeleeCrit 4.5 + SpellCrit 13". `TestEmeraldDragonWhelpKeepsItsCritTablesApart` fails on the old sheet. | CONFIRMED-BUG (no published number) | E 49db95da4 |

## Every crit read in the engine (the site table)

Which table each roll reads after the split. "Trace" means the row was seen in the dynamic trace for a published spec; "client" is the client's DefenseType.

| site | spells that use it | table read | client agrees? |
|---|---|---|---|
| `PhysicalCritChance` (`spell_result.go:139`): every melee and ranged outcome (`OutcomeMelee*`, `OutcomeRanged*`, `applyAttackTableCrit*`), `OutcomeTickPhysicalCrit`, the druid's Rake/Shred snapshots | white swings, every weapon ability of every class, Raptor Strike, Mongoose Bite, Wing Clip, Holy Strike, seals and judgements (DT 2), Hammer of Wrath, Serpent Sting, Arcane/Aimed/Multi-Shot (DT 3), Windfury and other extra attacks, Deep Wounds, bleeds, hunter pet Claw/Bite, demon melee | Crit + MeleeCrit, unless the spell is DefenseType magic (then Crit + SpellCrit) | yes, every traced row |
| `SpellCritChance` (`spell_result.go:254`): `OutcomeMagic*`, magic dot ticks, snapshot crits, `warlock/apl_values.go:137` | every damage spell, poisons (8680 and ranks, DT 1), weapon-imbue procs (Flametongue, Frostbrand), item procs (Fiery Weapon 13897, Lifestealing 20004, Stormshroud 18980, 39 weapon procs), totems (Searing, Magma, Fire Nova), Ignite, traps, Holy Shield, Lash of Pain, Firebolt | Crit + SpellCrit + school crit, except a spell with DefenseType melee or ranged, or physical with none, which reads Crit + MeleeCrit | yes, except C2 (fixed) |
| `HealingCritChance` (`spell_result.go:279`): `OutcomeHealing*`, heal-over-time ticks | every heal, Healing Stream ticks, Prayer of Mending, Penance, Lightwell | Crit + SpellCrit + spell bonus, now + school crit (C3) | yes |
| `critMultiplierAs` (`spell_outcome.go:1080`) | crit damage | physical 2x for melee and ranged defense types, magic 1.5x | yes |
| `periodicCritDefense` (`spell_outcome.go:91`) | dot ticks | the spell's DT; none: physical melee, else magic | yes (owner ruling on dot crit unchanged) |
| Hunter pets (`hunter/pet.go:215`) | pet melee and abilities | inherits the hunter's Crit and MeleeCrit (415429 channel 52) | logs: pet melee 8.1% of 40,189 against the hunters' Auto Shot 7.4%: aligned |
| Warlock demons (`warlock/pet.go`) | Lash of Pain, Firebolt, demon melee | the warlock lane's E de87d2393 (flat measured rates) | see "Pets" below |
| Totems (`shaman/fire_totems.go`, `water_totems.go`) | Searing, Magma, Healing Stream | the shaman's own spell table (the spells are the shaman's) | logs: Searing "Attack" 11.2% and Magma 10.7% against the same shamans' Flame Shock 10.3%: aligned |
| Guardians (`common/guardians`) | Emerald Dragon Whelp, Core Hound | own sheet | C4 fixed the whelp; the hound has no crit |
| Damage shields (Thorns, Retribution Aura, Lightning Shield, Cloak of Flames, Aegis) | `registerStruckDamage`, aura procs | do not roll crit | logs: 0 crits in 16,000 hits: aligned |

## Crit sources by table (engine after this lane)

| source | client row | table | engine |
|---|---|---|---|
| Agility / Intellect | PlayerExpectedStat CritPerAgility / SpellCritPerIntellect | melee / spell | MeleeCrit / SpellCrit |
| item crit rating | ItemSparse stat 32 | all | Crit |
| Elixir of the Mongoose, Grilled Squid, Elixir of Sages | 17538, 1249523, 1250941: aura 290 | all | Crit |
| Elemental Sharpening Stone | 22755: aura 52, SpellEquippedItems mask 42483 (melee weapons only) | melee, not ranged | MeleeCrit, minus 2 on ranged (hybrids lane's Hammer of Wrath question: aligned, see below) |
| Brilliant Wizard Oil, Firestone | 25113, 23480-23483: aura 57 | spell | SpellCrit |
| Leader of the Pack, Moonkin Aura | 24932, 24907: aura 290 | all | Crit |
| Atiesh, Cenarion 5, Bloodvine 3 | aura 57 | spell | SpellCrit |
| Prophecy 5, Ten Storms 5 | 21092, 23570: aura 57 | spell, one school | SchoolBonusCritChance, now on heals too (C3) |
| Rallying Cry, Songflower, Slip'kik's | Forever tooltips (nether.wowhead.com/forever/tooltip/spell/22888, 15366, 22820, read 2026-11-02): "spells and attacks" | all | Crit (world buffs gated off) |
| Malice, Thundering Strikes, Nature's Majesty, Sharpened Claws, Arcane Instability, Lethal Attacks, Killer Instinct, Hack and Slash, Weaponmaster, Berserker Stance, Recklessness | aura 290 | all | Crit. Logs: Malice reaches poisons (Assassination 4.70% of 6,879 against 0.43% of 469), Thundering Strikes reaches every shaman spell and heal (Enhancement 10.5% against 5.7%), Balance's 290 talents reach Wrath and Moonfire |
| Conviction | 20117: aura 290, melee weapons only | melee, and heals | MeleeCrit plus a heal-only bonus (C1, C1b) |
| Cruelty, Elemental Devastation | aura 52 | melee | MeleeCrit |
| School and ability crit talents (Improved Moonfire, Improved Regrowth, Savage Strikes, Arcane Impact, Incineration, Improved Flamestrike, Critical Mass, Combustion, Divine Favor, Inner Focus, Holy Specialization, Cold Blood, Puncturing Wounds, Improved Ambush, Tidal Mastery, Malevolence, Agonizing Flames, Fire and Brimstone, Improved Overpower) | aura 107 misc 7 with class masks | the spell's own | BonusCritRating on the masked spells; the trace shows each on its spells only (e.g. Malevolence on the Succubus's Lash of Pain for Affliction 60, Improved Overpower +30 bonus on Overpower). Class lanes own the per-mask checks |

## Pets

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | **The Succubus's Lash of Pain (2.7% simmed against 7.1% logged in round 9).** The 7.1% (23 of 324) was a pool: 209 of the hits came from Affliction warlocks with 19+ points, who have Malevolence, whose class mask reaches Lash of Pain (the engine models it: 8.34% at 60 Affliction). Without Malevolence the logs give 7 of 115. The cause of the low sim number was the merged sheet plus the warrior Agility conversion on the demon (melee 7.6% before suppression against 2.0% of 1,470 Succubus swings in the logs) and Firebolt (Intellect at the mage rate: 8.4% against 2.2% of 2,383). The warlock lane fixed all three with measured rates. | `py/demons.py`. Warlock lane E de87d2393. | Fixed by the warlock lane; this lane agrees | (E de87d2393) |
| P2 | Hunter pets inherit the hunter's Crit and MeleeCrit. | Logs: pet Auto Attack 8.1% of 40,189, Claw 7.7%, Bite 8.7%, hunters' Auto Shot 7.4%. | Aligned | - |
| P3 | Hunter pet magic abilities read Crit + the pet's own SpellCrit (none), below its melee crit. The logs show Lightning Breath 12.8% of 406, above the pets' melee 8.1%. No published hunter uses a non-cat pet. | `py/petlogs.py` | OPEN (hunter lane: which channel feeds pet spells) | - |
| P4 | Totems roll the shaman's spell table. | Searing 11.2%, Magma 10.7% against Flame Shock 10.3% for the same shamans. | Aligned | - |

## Cross-cutting rows checked and found aligned

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| X1 | Every traced (spell, table) pair of every published spec and band matches the client DefenseType. The flagged rows are id reuse: Windfury's extra attack logs under the imbue id 8232 (a swing, melee); Arcane Missiles' channel 5143 (DT 0) is not the missile; poisons log under the item spell, the damage spell 8680 is DT 1; Immolation Trap's 13795 is the trap, its damage 13797 is DT 1; Penance and Prayer of Mending heal ids carry DT 0 and heal on the spell table. | `join.txt` | Aligned | - |
| X2 | Weapon procs: 30 table rows and 10 helper calls match the client DefenseType, except Thornspike 1293183 (client 3, engine melee: both read the melee table; only the dodge/parry rules differ). | `py/staticscan.py` | Aligned for crit (items lane note) | - |
| X3 | Crusader (20007, a Strength buff), Fiery Weapon (13897, DT 1), Lifestealing (20004, DT 1) | client rows | Aligned | - |
| X4 | Sharpening stone on Hammer of Wrath (hybrids lane's question). 22755 is aura 52 with SpellEquippedItems mask 42483: axes, maces, swords, polearms, staves, fists and daggers, no bow, gun, crossbow, thrown or wand. A weapon-dependent crit reaches that weapon's attacks; Hammer of Wrath is DefenseType 3, the ranged attack type, so the engine's minus 2 on ranged spells stands. | client rows | Aligned (no change) | - |
| X5 | Crit-triggered procs (Flurry, Enrage, Blood Craze, Deep Wounds, Ignite, Improved Shadow Bolt, Nature's Grace, Illumination, Infusion of Light, Elemental Devastation, Unbridled Wrath, Seal Fate, Reckoning, Water Shield): each keys on its ProcMask and the outcome its own spell rolled; the split changed only probabilities, not which spells can trigger them. | 21 sites, `grep OutcomeCrit\|DidCrit` | Aligned | - |
| X6 | simdb equip crit: aura 57 misc 0 equip spells are on three items only (Eye of the Beast 13968, Atiesh 22589, 236401); none is picked. | `py/sets.py`, ItemEffect TriggerType 1 | Aligned | - |
| X7 | Remaining pre-split merges: Paladin 1+1 (`item_sets_pve.go:26`) and crafted 1+1 (`crafted.go:64`) are equal on both tables, so the merge is exact. Rallying Cry and Songflower are "spells and attacks" in Forever's tooltips. | grep "merged from" | Aligned | - |

## Base crit columns (item 4)

The 70707 client still lacks the tables. `wago.tools/api/casc/<id>?version=1.60.1.70707` answers 200 with an empty body for chancetomeleecrit 3999262, chancetomeleecritbase 3999263, chancetospellcritbase 3999264 and chancetospellcrit 3999265 (basemp 1391664 returns 8,129 bytes), checked 2026-11-02; there is no DB2 in their place, and PlayerExpectedStat carries only the per-point rates. A log fit of base spell crit for players whose talent trees cannot add spell crit (`py/spellbase.py`):

| class | players, hits | observed | Int crit | implied base | Era base |
|---|---|---|---|---|---|
| druid (no Balance) | 26, 2,589 | 5.91% | 3.34 | 2.6 ± 0.5 | 1.8 |
| priest (no Holy) | 29, 2,107 | 6.55% | 4.22 | 2.3 ± 0.5 | 0.8 |
| shaman (Restoration only) | 8, 1,470 | 4.97% | 2.83 | 2.1 ± 0.6 | 2.3 |
| paladin (no Retribution, Consecration) | 32, 53,536 | 8.05% | 2.09 | 6.0 ± 0.1 | 3.5 |

The druid, priest and shaman rows sit within about 1.5 points of Era's column, with talent confounds (Inner Focus, Tidal Mastery) that lift them. The paladin row is 2.5 points above Era's 3.5. The Retribution group, after C1, implies 4.5. Era's columns stay; "Still open" has what would settle them.

## Before/after

Engine lane/r10-crit (head 03ec2ecd8) with site origin/main, against `$R/sb-before`. Headline metric: DPS, threat per second for tanks, effective HPS for healers. 10000 iterations.

| spec | alliance s1 | alliance s50001 | horde s1 | horde s50001 | horde s1 absolute |
|---|---|---|---|---|---|
| paladin-retribution 60 raid | -0.08% | -0.10% | -0.08% | -0.08% | 621.3 -> 620.8 |
| paladin-protection 60 raid (threat/s) | -0.80% | -0.80% | -0.78% | -0.78% | 993.8 -> 986.0 |
| paladin-holy 60 raid (eff. HPS) | 0.00% | 0.00% | 0.00% | 0.00% | 641.4 -> 641.4 |

Bands 20 to 50, published bare preset, horde, seed 1 / seed 50001:

| spec | 20 | 30 | 40 | 50 |
|---|---|---|---|---|
| paladin-retribution | -0.26% / 0.00% | -0.14% / -0.14% | -0.20% / -0.10% | -0.24% / -0.24% |
| paladin-protection | 0 / 0 | 0 / 0 | 0 / 0 | 0 / 0 |
| paladin-holy | 0 / 0 | 0 / 0 | 0 / 0 | 0 / 0 |
| hunter-survival (C2) | 0 / 0 | 0 / 0 | +0.20% / +0.20% (152.8 -> 153.1) | 0 / 0 |

Retribution moves little because Consecration is a small share of its damage; Protection loses 0.8% threat through Consecration and Holy Shield. Every other spec is unchanged (C3 and C4 touch nothing published; the dynamic trace confirms no other published spell reads a moved stat).

## For other lanes

- **Hybrids (paladin code):** C1 changes `sim/paladin/talents.go` applyConviction. Exorcism never crits in the beta logs (0 of about 200 hits across 13 paladins, Holy and Retribution alike; the sim gives it the spell table). Worth a look at the Forever rows for Exorcism (879 to 10314, 415068) and a tooltip or log check before launch.
- **Healers:** Healing Stream Totem in the logs is bimodal: 0 crits in 5,488, 4,582, 8,072 and 2,664 ticks for four shamans, about 5% for the rest (`py/hst.py`). Probably a hotfix boundary on the PERIODIC_CAN_CRIT flag; the engine (crits) matches the later logs.
- **Hunter:** P3 (pet spells' crit channel). Gloves of Holy Might (C2) moves Survival at band 40.
- **Items:** Blazing Rapier's Blaze (16898) is 20 a tick in the 70707 row (engine 10) and rolls a melee hit for a DT 1 spell; Brain Damage (24388, DT 2) uses a magic hit roll; Thornspike (1293183) is DT 3 not melee. None changes a crit table.
- **Auras/world buffs:** Forever's Rallying Cry of the Dragonslayer adds "spell damage and healing by up to 80" and both it and Slip'kik's say "No effect while in a raid instance"; the engine has neither (world buffs are gated off).
- **Warlock:** agree with E de87d2393. The Blizzard reply on #331 (demons take the warlock's plain crit stat) stays unwired here too: single source.
- **Integrator:** C1 moves Retribution and Protection (table above); Holy is unchanged after C1b. C2 moves hunter-survival at band 40. C3 and C4 move nothing published.

## Still open

- **Why Thundering Strikes reaches spells and Conviction does not.** Both are aura 290 with the same SpellEquippedItems row (class 2, mask 173555). The logs say Thundering Strikes reaches every shaman spell and heal (+4.7) and Conviction does not reach Consecration (−1.6 against non-Retribution). C1 follows the logs for each. Hack and Slash (13706) and Weaponmaster (12700) carry weapon rows too and stay every-table; nothing in the logs separates them. A Forever character sheet with the talent toggled would settle all four.
- **Base crit columns.** The client has no table. A level-60 character sheet with no gear and no talents per class (spell crit and melee crit) would settle the nine Era values; the paladin spell column looks low (above).
- **Blood Siphon** (Spire of Hakkar 1299943) has no SpellCategories row (DefenseType 0) and is physical; the engine declares it magic, so its tick crits on the spell table. No published spec at band 60; the band-50 casters carry it. Under the engine's own rule (physical with no defense type rolls melee) it would read melee crit; the server may not let a DefenseType-0 leech crit at all.
