# Sim audit round 12: auras lane

Scope: raid buffs, debuffs, the consumable buffs in the presets, and the racials (`sim/core/racials.go`).

Sources:
- The 1.60.1.70753 client tables (wago CSVs fetched 2026-11-02, `scratchpad/r11/c70753/builds/1.60.1.70753/raw`), compared with 70707 where it matters. They were read with `$R/auras/tools/sp.py` (a 70753 copy of round 10's tool, with every attribute word added), `cons.py` (consumable to item to spell) and round 10's `ranks.py`.
- Wowhead Forever server tooltips (`nether.wowhead.com/forever/tooltip/spell/<id>`), read 2026-11-03. A tooltip and its client row count as one source.
- forever-bugs #396, read 2026-11-03.

Engine branch `lane/r12-auras` (from 2606bc05e). Site branch `wt/r12-auras` (from c6bdc855c).

**Found: 5 CONFIRMED-BUG rows, all fixed. By origin: old 4, new-data 1.** No R11, landing, R9 or R10 bugs. There are also 3 note fixes: the Berserking comment, the Endurance comment, and four preset reasons.

No published number moves. Every published race was re-simmed and is identical: Enhancement, Elemental, Fire, Arms, Beast Mastery and Demonology, horde, seed 1. The fixes reach orc healers, orc hunters and orc warlocks, and no published entry is one of those. They do move /sim for a player who picks one.

## Assigned items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | **Windfury Totem's bonus outlived the totem.**<br>- **What was wrong:** the shaman's own Windfury Totem kept a tracker aura under vanilla's weapon-enchant ids (8514, 10607, 10611). The aura lasted 10 s and the totem refreshed it every 5 s. When the totem ended it ran on, the vanilla "totem twist" linger.<br>- **What the client says:** the totem (8512, 10613, 10614) summons an NPC whose aura is 8515, 10609 or 10612. That aura is an area party aura (SpellEffect effect 65, aura 42, the 20% proc of 8516 / 10608 / 10610) with no duration. 8514, 10607 and 10611 are not in the client at all. So the bonus ends the moment the totem does, and Forever has no twisting.<br>- **What changed:** the tracker now uses the client's ids and lasts exactly as long as the totem, like every other buff totem.<br>- The raid buff (`core/buffs.go` ApplyWindfury, the raid preset's Windfury Totem) is a permanent party aura and was right already. No published rotation casts Windfury Totem: Enhancement's air slot is Grace of Air (round 9). Horde melee and druids in forms get the raid buff, so they do not move.<br>- The upstream fork-only test rotation (`ui/enhancement_shaman/apls/default.apl.json`) twisted Windfury against 10611's remaining time. With no linger there is no twist, so its Windfury lines are gone. | SpellEffect, SpellMisc and SpellName for 8512 to 10614; Wowhead Forever 8512 and 10614 ("enhances the melee attacks of all party or raid members within 30 yards ... Lasts 5 min"), 8515 and 10612 ("Proc chance: 20%, 100ms cooldown"). The test `TestWindfuryTotemBonusEndsWithTheTotem` (shaman/enhancement) failed first: no aura under 10612. TestEnhancement golden regenerated for the test rotation. | CONFIRMED-BUG (latent: no published number) | old (upstream wowsims; round 9 hybrids X3 saw the Era ids and called them cosmetic) | E 651660849 |
| A2 | **Blood Fury's spell power did not reach heals.**<br>- **What was wrong:** the engine multiplied spell power and spell damage by 1.10 but not healing power, so an orc healer's heals gained nothing.<br>- **What the client says:** Blood Fury 20572's third effect is aura 317 (spell power percent) at 10 with misc value 0. It has no school mask and no damage-only selector. The tooltip says "Increases Attack Power and Spell Power by 10% for 15 sec". The same client's other aura-317 rows (Totemic Wrath 77747, Templar 1240574) are spell power buffs with no damage qualifier. In this engine, healing is its own stat (`core/spell_result.go` HealingPower).<br>- **What changed:** healing power is now in the multiplied set. Nothing feeds spell damage from healing power (`stats/deps.go`), so no bonus counts twice.<br>- Orc healers that Forever allows: Restoration shaman (orc is a horde shaman race). The published Restoration shaman is tauren. An orc Restoration shaman gains +0.4 to +0.5% effective HPS. Damage specs do not move. | SpellEffect 20572; Wowhead Forever 20572. The test `TestBloodFuryIsTenPercentOfTotalPowers` failed first: healing power 600, want 660. | CONFIRMED-BUG (one source: the client row and its tooltip; the server's handling of 317 is not visible) | old (d1af3529f, the client racials) | E f085a6c29 |
| A3 | **The Berserking note said 10 s.** The code and the 70753 client say 15 s (20554, SpellMisc 15000 ms, 3 min cooldown, no global cooldown). The note and the constant's comment now cite the shipped row. | SpellMisc and SpellCooldowns 20554 | note fix | old | E 0dea180b6 |
| A4a | **Orc Command gave an orc hunter's or warlock's item guardians 5%.**<br>- **What the client says:** 20576 (hunter) and 20575 (warlock) are aura 107 (flat modifier), misc 8 (every effect), +5 on the pet's "Tamed Pet Passive (DND)". For the hunter that is 8875, on every pet family's skill line. For the warlock it is 18727 to 18730, on the Voidwalker (and Felguard), Imp, Succubus and Felhunter lines. That passive is aura 79 (damage done, all schools) at 0, so the class pet deals 5% more on everything: melee and every ability. An item guardian (Emerald Dragon Whelp, the Core Hound set) has no such passive.<br>- **What the engine did:** it gave the 5% to every pet in `character.Pets`, guardians included.<br>- **What changed:** guardians are skipped. | SpellEffect and SkillLineAbility 20575, 20576, 8875, 18727 to 18730; SkillLine names 188, 189, 204, 205, 2887. The test `TestCommandReachesClassPetsAndTheTwoDemonBolts` failed first: guardian at 1.05. | CONFIRMED-BUG (latent) | old (round 9 K1 narrowed it to hunters and warlocks and missed the guardians) | E 8e4de235c |
| A4b | **70753 gives Command a second warlock effect that the engine did not have.**<br>- 20575 effect 1 (new in 70753, absent in 70707) is aura 108 (percent modifier), misc 0 (damage), +5, on warlock class mask 0x3000. That mask covers every Firebolt rank (0x1000: 3110, 7799 to 7802, 11762, 11763, 20801) and every Lash of Pain rank (0x2000: 7814 to 7816, 11778 to 11780).<br>- The owner's spell modifiers reach its pet's spells, so the Imp's Firebolt and the Succubus' Lash of Pain take 5% more again. Torment and Shadow Bite do not.<br>- The warlock lane found the same (their O1). | SpellEffect 20575 (70707 vs 70753); SpellClassOptions set 5. Same failing test as A4a: Firebolt and Lash at 1, want 1.05. | CONFIRMED-BUG | new-data | E 8e4de235c |
| A4c | **Command and the pet talents multiplied; in the client they add.**<br>- **What the client says:** Unholy Power 18769 is aura 107, misc 8, +10 at 5/5, on the same demon passive (mask bit 26). Unleashed Fury 19616 is aura 107 on the hunter pet passive 8875. Points on one aura add, so an orc warlock with Unholy Power 5 has 1 + 0.10 + 0.05 = 1.15 on the demon, not 1.10 × 1.05 = 1.155. Firebolt and Lash of Pain are 1.15 × 1.05 = 1.2075.<br>- **What changed:** `Pet.AddPassiveDamagePercent` (`core/pet.go`) holds the summed points. Command, Unholy Power (`warlock/pet.go`) and Unleashed Fury (`hunter/talents.go`) call it.<br>- Only orc hunters and warlocks move: orc hunters -0.1 to -0.15%, orc warlocks +0.02 to +0.06%. A non-orc pet's multiplier is unchanged. | SpellEffect 18769, 19616, 20575, 20576. The test `TestCommandAddsToTheTalentPointsOnThePetPassive` fails on the old stacking (1.155, want 1.15). `TestUnholyPowerRaisesAllDemonDamageOnce` now expects 1.15, with Lash 1.05. | CONFIRMED-BUG | old (the multiplicative talents predate round 9) | E 46aa4f61a |
| A5 | **Blessing of Kings and Divine Spirit are baseline trainer spells in Forever.** The trainables lane makes the ruling; the client reads:<br>- Divine Spirit: SkillLineAbility Discipline 613, acquire 0, ranks 14752 / 14818 / 14819 / 27841 at 30 / 40 / 50 / 60, plus Prayer of Spirit 27681 at 60. It is in `trainables/priest.json` and not in `talents/priest.json`.<br>- Blessing of Kings: Protection 267, acquire 0, from 20. Greater Blessing of Kings 25898 at 60.<br>- **Do the presets agree?** The engine's `DivineSpiritRanks` starts at 30 with the client's amounts (17, 23, 33, 40), so every band at 30 and above gets it from any priest. Kings stays out of the raid preset for the separate reason that it may not stack with Heart of the Lion (B5, open). Paladins are on both factions, so either faction's raid can bring both.<br>- **Text fixed:** the preset's Divine Spirit reason ("a Discipline talent a Phase 1 raid priest takes") and the kit comment in `sim/leveling/kit.go` now say trainer spell. | SkillLineAbility; trainables and talents json; `core/buff_ranks.go:226` | Aligned (note fix) | - | S f8b327ab8 |
| A6 | **Blessings last 1 hour in Forever.** Every Blessing of Might rank (19740 to 19838, 25291), Wisdom (19742 to 19854, 25290), Kings 20217, and the Greater Blessings 25782, 25916, 25894, 25918 and 25898 are 3,600,000 ms. The tooltips agree, for example 19838 "increasing melee attack power by 112 for 1 hour" and 25916 "...by 133 for 1 hour". The guides' claim is right. The engine's `paladin/blessing_ranks.go` blessingDuration is already `time.Hour` with a client test, and the raid blessings are permanent. No rebuff assumption depends on it, since a fight is 180 to 300 s. | SpellMisc DurationIndex; Wowhead Forever 19838, 25916, 20217, 25898 | Aligned | - | - |
| A7 | **Heart of the Lion is 280 attack power at 60, plus 10% of all stats, and the hunter gets another 10%.**<br>- 409580, the hunter's own spell (learned at 1, Survival skill line 51, permanent): aura 137 (all stats, misc -1) +10%. It also has a periodic dummy (aura 226, every 14 s) that sends 409583.<br>- 409583 (15 s, refreshed): aura 137 +10% all stats, plus aura 99 and aura 124 at 40 + 4 per level, MaxLevel 60, BaseLevel 0. That is 40 + 4 × 60 = 280 at 60, and the tooltip reads 280.<br>- Round 11's 276 came from the old (TrinityCore 3.3.5) rule that subtracts SpellLevel. The modern rule this client follows subtracts BaseLevel, which is 0.<br>- The engine has 280 (`HeartOfTheLionRanks`) and 1.10 × 1.10 for the hunter. The preset note's "276" is corrected. | SpellEffect, SpellLevels and SpellMisc 409580, 409583; Wowhead Forever 409580 ("Attack Power by 280 for all nearby allies ... additional 10%") and 409583 | Aligned (closes round 11 B4; note fix) | - | S f8b327ab8 |

## Raid buffs (fresh re-audit)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| B1 | Every `BuffRanks` table was re-run on 70753: spell, learn level and amount for Arcane Intellect, Mark of the Wild (armor, stats, resistances), Power Word: Fortitude, Divine Spirit, Blessing of Might and Wisdom, Battle Shout, Trueshot Aura, Devotion and Retribution Aura, Thorns, Stoneskin, Blood Pact and Mana Tide. Only the round-11 by-design mismatches flag. The totem tables were re-checked by hand on their aura spells:<br>- Strength of Earth 8076 / 8162 / 8163 / 10441 = 7 / 14 / 25 / 42 (25362's 53 is an Ahn'Qiraj book).<br>- Grace of Air 8836 / 10626 = 49 / 77.<br>- Mana Spring 5677 to 10494 = 4 / 6 / 8 / 10 every 2 s, so 25 mp5 at 60 and 31.25 with Restorative Totems. | `ranks.py` over `engine/sim/core` (`$R/auras/ranks.txt`); SpellEffect | Aligned | - | - |
| B2 | **Trueshot Aura is ranged attack power only** in Forever (aura 124; Era had aura 99 as well). It is 15 / 22 / 30 / 40 / 50 at 25 / 32 / 40 / 50 / 60 (1299346, 1299348, 19506, 20905, 20906), 30 minutes. The engine grants ranged attack power only, by rank. The preset reason said "Attack power" and now says ranged. The guides' "50 ranged attack power at 60" is right. | SpellEffect 1299346 to 20906; `core/buffs.go:1408` | Aligned (note fix) | - | S f8b327ab8 |
| B3 | **Leader of the Pack (24932) and Moonkin Aura (24907) are both aura 290 at 3**: crit on every table, not Era's melee-only and spell-only crit. The engine adds 3% unified crit for each, with no shared exclusivity, as the client has none. Reach levels 30 and 40 follow the talent tiers (Feral tier 4, Balance tier 6, `10 + 5 × tier`). | SpellEffect 24932, 24907, 24858; talents/druid.json | Aligned | - | - |
| B4 | Windfury Totem raid buff: 95 / 179 / 246 attack power (8516, 10608, 10610), 1 s, 2 charges, 20% proc with a 100 ms recovery (8515, 10612). It is learned at 32 / 42 / 52. Each of these matches the engine. | SpellEffect and SpellAuraOptions | Aligned | - | - |
| B5 | Blessing of Kings with Heart of the Lion: still no source. Both are aura 137, misc -1, at 10, with no shared exclusivity row. forever-bugs has no issue on it (searched 2026-11-03). | client rows; gh search | OPEN (round 7) | - | - |

## Raid debuffs

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| D1 | Each of these was re-read on 70753 and matches: Sunder Armor, Faerie Fire, Curse of Recklessness, Curse of the Elements (damage and resistance), Hunter's Mark and Judgement of the Crusader (161). Judgement of Wisdom and Light are stated on their triggered spells and Expose Armor is scripted, as round 11 found. | `ranks.py` | Aligned | - | - |

## Consumables in the presets

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | Every id in `sim/request/raid-consumes.json` was traced item → ItemEffect → spell → SpellEffect on 70753 and compared with `core/consumes.go`. All match:<br>- Elixirs: Sages 25 Spirit and 2% crit on every table (aura 290); Greater Mageblood 20 mp5; Greater Cleric's 40 healing; Mongoose 25 Agility and 2%; Greater Arcane 35; Fire Power 10; Shadow Power 40; Frost Power 15; Fortitude 200 health; Greater Defense 450 armor.<br>- Flasks: Distilled Wisdom 2000 mana, Supreme Power 150, Titans 1200.<br>- Juju, Blasted Lands and alcohol: Juju Might 40 / 40, Juju Power 30, R.O.I.D.S., Scorpok, Cortex and Gizzard Gum 25, Rumsey Rum 15.<br>- Foods: Baked Salmon 28, Nightfin 22, Poached Sunscale 40 attack power, Sage's Tea 44 healing, Smoked Desert Dumplings 20 Strength (20452, level 45).<br>- Potions and conjured: Frenzy 28, Major Frenzy 80, Major Spellblasting 47; Major Mana 1800 at 0.5 variance (1350 to 2250); Mighty Rage 45 to 75 rage and 60 Strength; Demonic Rune 900 to 1500; Thistle Tea 100 energy. | `$R/auras/tools/cons.py`; `core/consumes.go` | Aligned | - | - |

## Racials of the published races

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | Every racial below was re-read on 70753 and matches the engine:<br>- **Human:** Ruthless 1% (aura 290); The Human Spirit 10% Spirit.<br>- **Dwarf:** Artisan's Precision, 1% melee and spell hit (aura 54 + 55, the engine's single Hit stat); Big Game Hunter 5% against beasts.<br>- **Night Elf:** Elune's Light 10% crit on every table for 15 s, 3 min cooldown, no global cooldown; Quickness 1% dodge.<br>- **Undead:** Touch of the Grave 20% / 30%, 15 s proc recovery.<br>- **Tauren:** Endurance 5% health.<br>- **Gnome:** Expansive Mind 5% mana, rage and energy; Eureka! 2 min.<br>- **Troll:** Berserking 10% for 15 s, 3 min, no global cooldown; Beast Slaying 5%.<br>- **Orc:** Blood Fury 2 min, no global cooldown; see A2 and A4. | SpellEffect, SpellMisc, SpellAuraOptions, SpellCooldowns; Wowhead Forever 20550, 1259719, 20582 | Aligned | - | - |
| R2 | **Tauren Endurance's 2% crit damage is gone from the client.** forever-bugs #396 ("Endurance 2% crit damage not working") was closed "fixed!" on 2026-10-27. But 20550 is a single effect (aura 133, 5% health) in both 70707 and 70753, and the tooltip reads "Total Health increased by 5%". The engine follows the client. Its stale comment, which called the 2% open, is corrected. | SpellEffect 20550 (70707, 70753); Wowhead Forever 20550; forever-bugs #396 | Aligned (note fix) | - | E f1b0d5a66 |

## Before/after

"Before" is `$R/sb-before` from `$R/site-before/sim`, or `$R/before.tsv`. "After" is lane engine f1b0d5a66 with site f8b327ab8. Raid preset, band 60, horde, 10,000 iterations.

**Published entries (identical):**

| spec (published race) | horde s1 before → after |
|---|---|
| shaman-enhancement (orc) | 586.0 → 586.0 |
| shaman-elemental (orc) | 479.6 → 479.6 |
| mage-fire (orc) | 672.9 → 672.9 |
| warrior-arms (orc) | 731.7 → 731.7 |
| hunter-beast-mastery (troll) | 626.9 → 626.9 |
| warlock-demonology (troll) | 641.4 → 641.4 |

**The same entries with the race set to orc** (BiS copies in `$R/auras/bis-orc`; /sim reach only):

| spec | s1 | s50001 | s1 absolute |
|---|---|---|---|
| shaman-restoration (effective HPS) | +0.44 | +0.51 | 842.6 → 846.3 |
| hunter-beast-mastery | -0.14 | -0.16 | 633.1 → 632.2 |
| hunter-marksmanship | -0.10 | -0.11 | 618.6 → 618.0 |
| hunter-survival | 0.00 | 0.00 | 746.9 → 746.9 |
| warlock-affliction | +0.02 | 0.00 | 578.7 → 578.8 |
| warlock-demonology | +0.03 | +0.03 | 646.0 → 646.2 |
| warlock-destruction | +0.05 | +0.06 | 645.4 → 645.7 |

What moved:
- The orc Restoration shaman gains from Blood Fury on heals (A2).
- Orc hunters lose about 0.1%, because Command now adds to Unleashed Fury instead of multiplying it (A4c).
- Orc warlocks gain a little: Lash of Pain's +5% (A4b) outweighs the additive passive (A4c). This matches the warlock lane's estimate.

## Files to regenerate

- None of the nightly or ladder files: no published race moves.
- Fork golden regenerated in its own commit: `sim/shaman/enhancement/TestEnhancement.results` (the test rotation lost its Windfury twist, E 651660849).

## Measure at refresh

- None from this lane. Race picks are within 1% (A2's +0.5% for an orc Restoration shaman does not beat the tauren pick's reason to change).

## For other lanes

- **Warlock lane:** O1 is fixed here. E 46aa4f61a changes one line in `sim/warlock/pet.go` (Unholy Power now calls `AddPassiveDamagePercent`) and updates `warlock/pet_spells_test.go` (the demon multiplier is 1.15, Lash 1.05). Merge on top of any warlock-lane edit of those files.
- **Physical lane (hunters):** E 46aa4f61a changes one line in `sim/hunter/talents.go`: Unleashed Fury now adds into the pet passive. Please also check that Unleashed Fury's per-rank amount matches 19616 (aura 107, misc 3, +20 on 8875 in the shown row; the engine has 3% a rank). Beast Mastery 409368 (+15 on 8875, misc 8) also raises the same passive. If either is modelled, it should call the same helper.
- **Guides lane (wow-player questions):**
  - Blessings and Greater Blessings last 1 hour (A6).
  - Heart of the Lion gives everyone near the hunter 10% to every stat plus 280 melee and ranged attack power at 60 (40 + 4 a level), and the hunter a further 10% stats (A7).
  - Trueshot Aura is 50 ranged attack power at 60, ranged only (B2).
  - Divine Spirit is a priest trainer spell from 30, not a talent (A5).
  - Windfury Totem's bonus ends the moment the totem ends: no 10-second linger, so no totem twisting (A1).
- **Platform:** `data/curated/engine-spell-names.json` names the vanilla Windfury ids 8514, 10607 and 10611. The engine no longer uses them; its tracker is now 8515, 10609 and 10612, which the client names. The three entries can go. No published result carries them.
- **Trainables lane:** A5's client reading (Divine Spirit and Kings are trainer spells) is for your ruling; the presets already agree with it.

## Still open

1. B5: Blessing of Kings with Heart of the Lion (round 7). A beta character sheet with both settles it.
2. A2 rests on the client row and its tooltip only. An orc healer's log with Blood Fury up, at a gear level where 10% of bonus healing is visible, would confirm it.
3. Round 11's C1 (the new elixirs stacking with Mongoose) and C7 (Greater and Superior potions below 55) are unchanged.
