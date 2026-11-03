# Leveling BiS: Enhancement

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-254000000000000000-0000000000000000)

Set DPS (verified): 37.4. Weights run: 2.1s. Verify run: 2.1s. 258 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.435 ± 0.032, crit=0.971 ± 0.031 per rating point (14 rating = 1%, 13.596 per %), hit=1.573 ± 0.083 per rating point (10 rating = 1%, 15.729 per %), melee_haste=9.558 ± 0.324, intellect=1.166 ± 0.027, spell_power=0.984 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (0.57 DPS) | yes | Trapper's Leather Hood (252505, -0.24 DPS) [crafted]; Totemic Leather Hood (252448, -0.25 DPS) [crafted]; Brawler's Leather Hood (252504, -0.29 DPS, sim-verified) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 2.9 attack_power points (0.10 DPS) | yes | Scholarly Pendant (277203, +0.00 DPS) [quest]; Tarnished Locket (279870, -0.32 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 15.4 attack_power points (0.55 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.21 DPS) [crafted]; Serpent's Shoulders (5404, -0.29 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.41 DPS) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 8.6 attack_power points (0.31 DPS) | yes | Dark Leather Cloak (2316, -0.01 DPS) [crafted]; Lambent Scale Cloak (4706, -0.02 DPS) [world_drop]; Cape of the Brotherhood (5193, -0.05 DPS) [dungeon] |
| chest | Totemic Leather Armor (252435) | Leatherworking [crafted] | 23.4 attack_power points (0.83 DPS) | yes | Defender's Leather Armor (252434, -0.18 DPS) [crafted]; Brawler's Leather Armor (252490, -0.26 DPS, sim-verified) [crafted]; Tunic of Westfall (2041, -0.27 DPS) [quest] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 10.3 attack_power points (0.37 DPS) | yes | Bristlebark Bindings (14569, -0.07 DPS) [world_drop]; Forest Leather Bracers (3202, -0.11 DPS) [world_drop]; Owl Bracers (4796, -0.16 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 16.6 attack_power points (0.59 DPS) | yes | Serpent Gloves (5970, -0.04 DPS) [dungeon]; Gold-flecked Gloves (5195, -0.09 DPS) [dungeon]; Bristlebark Gloves (14572, -0.10 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.64 DPS) | yes | Deviate Scale Belt (6468, -0.17 DPS) [crafted]; Totemic Leather Belt (252431, -0.22 DPS) [crafted]; Brawler's Leather Belt (252428, -0.23 DPS, sim-verified) [crafted] |
| legs | Trapper's Leather Pants (252501) | Leatherworking [crafted] | 29.6 attack_power points (1.05 DPS) | yes | Totemic Leather Pants (252446, -0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -0.17 DPS) [crafted]; Defender's Leather Pants (252445, -0.21 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 17.5 attack_power points (0.62 DPS) | yes | Brawler's Leather Boots (252439, -0.01 DPS) [crafted]; Totemic Leather Boots (252442, -0.06 DPS) [crafted]; Trapper's Leather Boots (252440, -0.20 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 9.9 attack_power points (0.35 DPS) | yes | Minor Channeling Ring (1449, -0.09 DPS) [quest]; The 1 Ring (8350, -0.19 DPS) [world]; Volcanic Rock Ring (12053, -0.23 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.29 DPS) | yes | Minor Channeling Ring (1449, -0.03 DPS) [quest]; The 1 Ring (8350, -0.12 DPS) [world]; Volcanic Rock Ring (12053, -0.16 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 303.7 attack_power points (10.83 DPS) | yes | Twisted Chanter's Staff (890, -0.21 DPS) [world_drop]; Rockslicer (872, -0.50 DPS) [dungeon]; Night Reaver (1318, -1.00 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Erudite's Amulet; shoulder: Magician's Mantle; back: Glowing Lizardscale Cloak; chest: Totemic Leather Armor; wrist: Bravo's Armbands; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Trapper's Leather Pants; feet: Feet of the Lynx; finger1: Lavishly Jeweled Ring; finger2: Demon Band; main_hand: Smite's Mighty Hammer

No-known-source sample (15 of 258, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18845 Insignia of the Horde

### Band 30 (dwarf, 0000000000000000-255130030002000000-0000000000000000)

Set DPS (verified): 93.8. Weights run: 2.1s. Verify run: 2.2s. 458 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.006, agility=1.337 ± 0.033, crit=1.070 ± 0.035 per rating point (14 rating = 1%, 14.975 per %), hit=1.950 ± 0.092 per rating point (10 rating = 1%, 19.500 per %), melee_haste=9.662 ± 0.411, intellect=1.966 ± 0.028, spell_power=0.583 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 31.5 attack_power points (1.57 DPS) | yes | Dokebi Hat (14584, -0.09 DPS) [world_drop]; Resilient Cap (14401, -0.36 DPS) [world_drop]; Defender's Leather Helm (252455, -0.37 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 23.8 attack_power points (1.18 DPS) | yes | Kaleidoscope Chain (13084, -0.13 DPS) [world_drop]; River Pride Choker (13087, -0.79 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.79 DPS) [world_drop] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 25.1 attack_power points (1.25 DPS) | yes | Bloodmage Mantle (7684, -0.11 DPS) [dungeon]; Feline Mantle (3748, -0.22 DPS) [dungeon]; Forest Tracker Epaulets (2278, -0.54 DPS, sim-verified) [world_drop] |
| back | Sergeant Major's Cape (16315) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 21.2 attack_power points (1.06 DPS) | yes | Resilient Cape (14400, -0.47 DPS) [world_drop]; Tigerstrike Mantle (13108, -0.52 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.87 DPS, sim-verified) [world_drop] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 41.5 attack_power points (2.07 DPS) | yes | Spirewind Fetter (9406, -0.20 DPS) [dungeon]; Green Silk Armor (7065, -0.53 DPS) [crafted]; Totemic Leather Tunic (252451, -0.58 DPS) [crafted] |
| wrist | Glowing Magical Bracelets (13106) | World drop [world_drop] | 21.7 attack_power points (1.08 DPS) | yes | Technician's Bracers (270042, -0.29 DPS) [quest]; Barbaric Bracers (18948, -0.42 DPS) [crafted]; Hawkeye's Bracers (14590, -0.82 DPS, sim-verified) [world_drop] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 22.0 attack_power points (1.10 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.05 DPS) [dungeon]; Toughened Leather Gloves (4253, -0.10 DPS) [crafted]; Wolfclaw Gloves (1978, -0.20 DPS) [dungeon] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 33.9 attack_power points (1.69 DPS) | yes | Skirmisher's Leather Belt (252461, -0.20 DPS) [crafted]; Skulker's Leather Belt (252520, -0.49 DPS) [crafted]; Stalker's Leather Belt (252521, -0.50 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 49.6 attack_power points (2.47 DPS) | yes | Trapper's Leather Pants (252501, -0.88 DPS) [crafted]; Totemic Leather Pants (252446, -0.98 DPS) [crafted]; Totemic Leather Leggings (252458, -1.16 DPS, sim-verified) [crafted] |
| feet | Acidic Walkers (9454) | Gnomeregan: Viscous Fallout [dungeon] | 18.6 attack_power points (0.93 DPS) | yes | Gilded Slippers (254001, -0.04 DPS) [crafted]; Feet of the Lynx (1121, -0.10 DPS) [world_drop]; Brawler's Leather Boots (252439, -0.10 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.0 attack_power points (1.00 DPS) | yes | Lavishly Jeweled Ring (1156, -0.28 DPS) [dungeon]; Snake Hoop (6750, -0.33 DPS) [quest]; Tiger Band (6749, -0.50 DPS) [quest] |
| finger2 | Seal of Wrynn (2933) | An Audience with the King [quest] | 17.9 attack_power points (0.89 DPS) | yes | Lavishly Jeweled Ring (1156, -0.17 DPS) [dungeon]; Snake Hoop (6750, -0.23 DPS) [quest]; Tiger Band (6749, -0.39 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (93.8 DPS) | yes | Corpsemaker (6687, -0.47 DPS) [dungeon]; Cobalt Crusher (7730, -1.99 DPS) [dungeon]; Viscous Hammer (13045, -22.50 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Archmagister's Faceted Pendant; shoulder: Death Speaker Mantle; back: Sergeant Major's Cape; chest: Garb of Florid Feathers; wrist: Glowing Magical Bracelets; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Ferine Leggings; feet: Acidic Walkers; finger1: Thunderbrow Ring; finger2: Seal of Wrynn; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 458, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 40 (dwarf, 0000000000000000-255130030005102031-0000000000000000)

Set DPS (verified): 130.8. Weights run: 2.6s. Verify run: 2.4s. 706 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.292 ± 0.044, crit=1.308 ± 0.055 per rating point (14 rating = 1%, 18.309 per %), hit=1.888 ± 0.106 per rating point (10 rating = 1%, 18.879 per %), melee_haste=7.503 ± 0.814, intellect=3.498 ± 0.058, spell_power=0.474 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 63.3 attack_power points (4.44 DPS) | yes | Thinking Cap (2624, -0.27 DPS) [world]; Enduring Cap (3020, -0.51 DPS) [world_drop]; Big Voodoo Mask (8201, -0.70 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 33.0 attack_power points (2.31 DPS) | yes | Glowing Eye of Mordresh (10769, -0.12 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.41 DPS) [world_drop]; Necklace of Calisea (1714, -0.60 DPS) [dungeon] |
| shoulder | Sheepshear Mantle (13115) | World drop [world_drop] | 56.0 attack_power points (3.92 DPS) | yes | Windchaser Amice (14432, -0.46 DPS) [world_drop]; Inquisitor's Shawl (19507, -0.50 DPS) [dungeon]; Green Silken Shoulders (7057, -0.96 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 40.7 attack_power points (2.86 DPS) | yes | Glyphed Cloak (4732, -1.14 DPS) [world_drop]; Resilient Cape (14400, -1.38 DPS) [world_drop]; Blackforge Cape (6424, -1.48 DPS, sim-verified) [dungeon] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 64.1 attack_power points (4.49 DPS) | yes | Spirewind Fetter (9406, -0.57 DPS) [dungeon]; Silksand Tunic (14417, -0.57 DPS) [world_drop]; Red Mageweave Vest (10007, -1.15 DPS, sim-verified) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 34.1 attack_power points (2.39 DPS) | yes | Glowing Magical Bracelets (13106, -0.01 DPS) [world_drop]; Warden's Wristbands (14600, -0.15 DPS) [world_drop]; Tough Scorpid Bracers (8205, -0.28 DPS) [crafted] |
| hands | Tough Scorpid Gloves (8204) | Leatherworking [crafted] | 44.4 attack_power points (3.11 DPS) | yes | Engineering Gloves (1659, -0.03 DPS) [world]; Stormcloth Gloves (10011, -0.04 DPS) [crafted]; Gnomish Mechanic's Gloves (6732, -0.27 DPS) [quest] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 55.8 attack_power points (3.91 DPS) | yes | Saltscale Girdle (273023, -0.44 DPS) [dungeon]; Razzeric's Customized Seatbelt (6726, -0.97 DPS) [quest]; Teacher's Sash (10747, -0.97 DPS) [quest] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 68.0 attack_power points (4.77 DPS) | yes | Red Mageweave Pants (10009, -1.36 DPS) [crafted]; Misplaced Pantaloons (276201, -1.58 DPS) [vendor]; Crimson Silk Pantaloons (7062, -1.59 DPS, sim-verified) [crafted] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 48.5 attack_power points (3.40 DPS) | yes | Skirmisher's Mail Boots (252564, -0.33 DPS) [crafted]; Stalker's Mail Boots (252562, -0.69 DPS) [crafted]; Briar Tredders (10582, -0.74 DPS) [dungeon] |
| finger1 | Voodoo Band (1996) (or Ogremind Ring (1993), Mindbender Loop (5009)) | Bloodscalp Witch Doctor [world] | 24.5 attack_power points (1.72 DPS) | yes | Mindbender Loop (5009, +0.00 DPS) [world_drop]; Seal of Wrynn (2933, -0.04 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.06 DPS) [dungeon] |
| finger2 | Ogremind Ring (1993) (or Mindbender Loop (5009)) | World drop [world_drop] | 24.5 attack_power points (1.72 DPS) | yes | Mindbender Loop (5009, +0.00 DPS) [world_drop]; Seal of Wrynn (2933, -0.04 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.06 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-verified (130.8 DPS) | yes | Mograine's Might (7723, +0.00 DPS) [dungeon]; Bonebiter (6830, -0.86 DPS) [quest]; Pendulum of Doom (9425, -3.85 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; shoulder: Sheepshear Mantle; back: Sergeant Major's Cape; wrist: Windchaser Cuffs; hands: Tough Scorpid Gloves; waist: Deathmage Sash; feet: Prowler's Leather Shoes; finger1: Voodoo Band; finger2: Ogremind Ring; main_hand: Fiery War Axe

No-known-source sample (15 of 706, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 50 (dwarf, 3200300000000000-255130030005102051-0000000000000000)

Set DPS (verified): 172.2. Weights run: 2.6s. Verify run: 3.0s. 909 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.397 ± 0.051, crit=1.653 ± 0.067 per rating point (14 rating = 1%, 23.145 per %), hit=3.088 ± 0.140 per rating point (10 rating = 1%, 30.879 per %), melee_haste=10.925 ± 1.128, intellect=4.307 ± 0.066, spell_power=0.780 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | 105.0 attack_power points (6.98 DPS) | yes | Winged Helm (13112, -0.39 DPS) [world_drop]; Bad Mojo Mask (9470, -0.97 DPS) [dungeon]; Red Mageweave Headband (10033, -1.49 DPS, sim-verified) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 60.3 attack_power points (4.01 DPS) | yes | Glowing Eye of Mordresh (10769, -1.36 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.49 DPS) [world]; Kaleidoscope Chain (13084, -1.96 DPS) [world_drop] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 97.8 attack_power points (6.51 DPS) | yes | Rotgrip Mantle (17732, -0.68 DPS) [dungeon]; Rockshard Pauldrons (9411, -1.76 DPS) [dungeon]; Imperial Red Mantle (8250, -1.84 DPS) [world_drop] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 47.4 attack_power points (3.15 DPS) | yes | Sergeant Major's Cape (16336, -0.08 DPS) [pvp]; Keeper's Cloak (14665, -0.29 DPS) [world_drop]; Big Voodoo Cloak (8216, -0.31 DPS) [crafted] |
| chest | Knight's Mail Armor (223078) | Captain Dirgehammer [vendor] | sim-verified (172.2 DPS) | yes | Relentless Chain (17777, -0.70 DPS) [quest]; Hibernal Robe (8113, -0.84 DPS) [world_drop]; Acumen Robes (17775, -1.86 DPS, sim-verified) [quest] |
| wrist | Forgotten Wraps (9433) | World drop [world_drop] | 55.9 attack_power points (3.72 DPS) | yes | Shizzle's Nozzle Wiper (11917, -0.28 DPS) [quest]; Runic Leather Bracers (15092, -0.33 DPS) [crafted]; Bracers of the Stone Princess (17714, -0.42 DPS) [dungeon] |
| hands | Stonerender Gauntlets (17007) | World drop [world_drop] | 77.2 attack_power points (5.13 DPS) | yes | Prowler's Leather Gauntlets (252547, -0.71 DPS) [crafted]; Skirmisher's Mail Gauntlets (252586, -0.98 DPS) [crafted]; Sergeant Major's Mail Gauntlets (223076, -1.10 DPS) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 86.5 attack_power points (5.75 DPS) | yes | Serenity Belt (13144, -0.88 DPS) [world_drop]; Sagebrush Girdle (17778, -0.92 DPS) [quest]; Imperial Red Sash (8253, -0.92 DPS) [world_drop] |
| legs | Knight's Mail Legplates (223074) | Captain Dirgehammer [vendor] | 98.8 attack_power points (6.57 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -1.26 DPS) [dungeon]; Ferine Leggings (6690, -1.41 DPS) [dungeon]; Imperial Red Pants (8251, -1.57 DPS, sim-verified) [world_drop] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 82.5 attack_power points (5.49 DPS) | yes | Skycaller's Leather Boots (252471, -1.56 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -1.56 DPS) [crafted]; Rancor Boots (11865, -1.84 DPS, sim-verified) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 50.9 attack_power points (3.38 DPS) | yes | Ogremind Ring (1993, -1.38 DPS) [world_drop]; Mindbender Loop (5009, -1.38 DPS) [world_drop]; Philanthropist's Ring (281635, -1.43 DPS) [quest] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993), Mindbender Loop (5009)) | Bloodscalp Witch Doctor [world] | 30.2 attack_power points (2.00 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Mindbender Loop (5009, +0.00 DPS) [world_drop]; Philanthropist's Ring (281635, -0.05 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+3.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -1.54 DPS) [world_drop] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bleakwood Hew (12769, -3.34 DPS) [crafted]; Dark Iron Pulverizer (11608, -3.92 DPS) [crafted]; Soulkeeper (1607, -4.85 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Eye of Theradras; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Imperial Red Cloak; chest: Knight's Mail Armor; wrist: Forgotten Wraps; hands: Stonerender Gauntlets; waist: Dawnspire Cord; legs: Knight's Mail Legplates; feet: Greaves of Withering Despair; finger1: Blackstone Ring; finger2: Voodoo Band; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Thorium Greatmace

No-known-source sample (15 of 909, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60 (dwarf, 3200330000000000-255130030005102051-0520000000000000)

Set DPS (verified): 228.5. Weights run: 2.4s. Verify run: 5.1s. 1856 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.958 ± 0.076, crit=2.739 ± 0.106 per rating point (14 rating = 1%, 38.342 per %), hit=3.690 ± 0.207 per rating point (10 rating = 1%, 36.903 per %), melee_haste=15.809 ± 1.595, intellect=4.796 ± 0.095, spell_power=0.623 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 261.7 attack_power points (20.54 DPS) | yes | Magister's Crown (16686, -6.38 DPS, sim-verified) [dungeon]; Blue Dragonscale Helm (252604, -9.65 DPS) [crafted]; Horns of Eranikus (10833, -10.18 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 67.1 attack_power points (5.27 DPS) | yes | Opaline Medallion (16623, -0.38 DPS) [quest]; Archlight Talisman (15856, -0.92 DPS) [quest]; Glowing Eye of Mordresh (10769, -1.92 DPS) [dungeon] |
| shoulder | Shroud of the Nathrezim (18720) | Stratholme: Balnazzar [dungeon] | 119.4 attack_power points (9.38 DPS) | yes | Mantle of the Timbermaw (19050, -0.64 DPS) [crafted]; Magister's Mantle (16689, -0.70 DPS) [dungeon]; Ironfeather Shoulders (15067, -1.11 DPS) [crafted] |
| back | Sergeant Major's Cape (16337) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 70.0 attack_power points (5.50 DPS) | yes | Cape of the Black Baron (13340, -1.62 DPS) [dungeon]; Keeper's Cloak (14665, -1.73 DPS) [world_drop]; Imperial Red Cloak (8248, -1.81 DPS, sim-verified) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 147.7 attack_power points (11.59 DPS) | yes | Magister's Robes (16688, -0.46 DPS) [dungeon]; Timbermaw Tunic (252484, -0.64 DPS) [crafted]; Blue Dragonscale Breastplate (15048, -0.66 DPS) [crafted] |
| wrist | Magiskull Cuffs (13107) | World drop [world_drop] | 81.5 attack_power points (6.40 DPS) | yes | Tranquil Wristguards (279256, +0.00 DPS, sim-verified) [crafted]; Spitfire Bracers (20481, -1.24 DPS) [crafted]; Hero's Bracers (8302, -1.27 DPS) [dungeon] |
| hands | Mourning Handguards (279257) | Leatherworking [crafted] | 108.3 attack_power points (8.50 DPS) | yes | Spitfire Gauntlets (20480, -1.60 DPS) [crafted]; Hero's Gauntlets (8305, -1.61 DPS) [dungeon]; Savage Gladiator Grips (11730, -1.81 DPS) [dungeon] |
| waist | Elunarian Belt (14465) | World drop [world_drop] | 107.9 attack_power points (8.47 DPS) | yes | Magister's Belt (16685, -0.22 DPS) [dungeon]; Cord of Elements (16673, -0.66 DPS) [dungeon]; Devout Belt (16696, -0.75 DPS) [dungeon] |
| legs | Red Dragonscale Leggings (252603) | Leatherworking [crafted] | 165.9 attack_power points (13.03 DPS) | yes | Ironfeather Leggings (252486, +0.00 DPS, sim-verified) [crafted]; Knight's Mail Legplates (223074, -3.61 DPS) [vendor]; Beaststalker's Pants (16678, -4.32 DPS) [dungeon] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 91.7 attack_power points (7.20 DPS) | yes | Timmy's Galoshes (13402, -0.82 DPS) [dungeon]; Rancor Boots (11865, -1.21 DPS) [quest]; Mooncloth Boots (15802, -1.54 DPS) [crafted] |
| finger1 | Glowing Crystal Ring (18402) | Uncovering Past Secrets [quest] | 57.6 attack_power points (4.52 DPS) | yes | Band of the Hierophant (13096, -0.75 DPS) [world_drop]; Frigid Ring (18679, -0.75 DPS) [world]; Maiden's Circle (13001, -1.00 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 56.9 attack_power points (4.47 DPS) | yes | Band of the Hierophant (13096, -0.70 DPS) [world_drop]; Frigid Ring (18679, -0.70 DPS) [world]; Maiden's Circle (13001, -0.95 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (228.5 DPS) | yes | Uther's Strength (11302, -2.31 DPS) [world_drop] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (228.5 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Forest Defender's Axe (279260) | Blacksmithing [crafted] | sim-verified (228.5 DPS) | yes | Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Staff of Hale Magefire (13000, -9.93 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Black Dragonscale Helm; shoulder: Shroud of the Nathrezim; back: Sergeant Major's Cape; chest: Dawn Armor; wrist: Magiskull Cuffs; hands: Mourning Handguards; waist: Elunarian Belt; legs: Red Dragonscale Leggings; finger1: Glowing Crystal Ring; finger2: Blackstone Ring; main_hand: Forest Defender's Axe

No-known-source sample (15 of 1856, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60, raid preset (dwarf, 0533001000000000-253030031005102051-0530000000000000)

Set DPS (verified): 631.8. Weights run: 2.5s. Verify run: 5.1s. 1856 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.005, agility=2.339 ± 0.100, crit=2.957 ± 0.133 per rating point (14 rating = 1%, 41.402 per %), hit=5.516 ± 0.321 per rating point (10 rating = 1%, 55.157 per %), melee_haste=20.868 ± 2.640, intellect=3.392 ± 0.123, spell_power=0.878 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 219.4 attack_power points (33.21 DPS) | yes | Blue Dragonscale Helm (252604, -12.30 DPS) [crafted]; Magister's Crown (16686, -16.34 DPS) [dungeon]; Mask of the Unforgiven (13404, -21.02 DPS, sim-verified) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 47.5 attack_power points (7.19 DPS) | yes | Archlight Talisman (15856, -0.46 DPS) [quest]; Opaline Medallion (16623, -0.51 DPS) [quest]; Skibi's Pendant (13089, -0.92 DPS) [world_drop] |
| shoulder | Shroud of the Nathrezim (18720) | Stratholme: Balnazzar [dungeon] | 101.8 attack_power points (15.41 DPS) | yes | Mantle of the Timbermaw (19050, -0.21 DPS) [crafted]; Magister's Mantle (16689, -3.05 DPS) [dungeon]; Ironfeather Shoulders (15067, -3.15 DPS) [crafted] |
| back | Sergeant Major's Cape (16337) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 63.4 attack_power points (9.60 DPS) | yes | Cape of the Black Baron (13340, -1.27 DPS) [dungeon]; Fel Cape (279269, -3.34 DPS) [crafted]; Cloak of Flames (3475, -6.20 DPS, sim-verified) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 183.1 attack_power points (27.71 DPS) | yes | Timbermaw Tunic (252484, +0.00 DPS, sim-verified) [crafted]; Savage Gladiator Chain (11726, -4.90 DPS) [dungeon]; Icemail Jerkin (1981, -10.87 DPS) [world_drop] |
| wrist | Tranquil Wristguards (279256) | Leatherworking [crafted] | 69.1 attack_power points (10.46 DPS) | yes | Spitfire Bracers (20481, -1.59 DPS) [crafted]; Magiskull Cuffs (13107, -1.73 DPS) [world_drop]; Lordly Armguards (13135, -2.06 DPS) [world_drop] |
| hands | Mourning Handguards (279257) | Leatherworking [crafted] | 96.6 attack_power points (14.62 DPS) | yes | Dreadmist Wraps (16705, -1.65 DPS) [dungeon]; Timbermaw Brawlers (19049, -2.67 DPS) [crafted]; Gauntlets of Elements (16672, -7.21 DPS, sim-verified) [dungeon] |
| waist | Elunarian Belt (14465) | World drop [world_drop] | 81.0 attack_power points (12.26 DPS) | yes | Might of the Timbermaw (19044, -0.32 DPS) [crafted]; Cord of Elements (16673, -0.54 DPS) [dungeon]; Magister's Belt (16685, -0.55 DPS) [dungeon] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 135.8 attack_power points (20.56 DPS) | yes | Red Dragonscale Leggings (252603, -1.47 DPS) [crafted]; Pristine Scorpid Leggings (252606, -3.91 DPS) [crafted]; Knight's Mail Legplates (223074, -4.14 DPS) [vendor] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 98.7 attack_power points (14.94 DPS) | yes | Timmy's Galoshes (13402, -3.79 DPS) [dungeon]; Shadowcraft Boots (16711, -5.01 DPS) [dungeon]; Rancor Boots (11865, -5.05 DPS) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 75.2 attack_power points (11.37 DPS) | yes | Glowing Crystal Ring (18402, -5.21 DPS) [quest]; Maiden's Circle (13001, -5.39 DPS) [world_drop]; Frigid Ring (18679, -6.24 DPS) [world] |
| finger2 | Band of the Penitent (13217) | Houses of the Holy [quest] | 41.4 attack_power points (6.27 DPS) | yes | Glowing Crystal Ring (18402, -0.11 DPS) [quest]; Maiden's Circle (13001, -0.28 DPS) [world_drop]; Frigid Ring (18679, -1.13 DPS) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (631.8 DPS) | yes | Uther's Strength (11302, -6.72 DPS) [world_drop] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (631.8 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Forest Defender's Axe (279260) | Blacksmithing [crafted] | sim-verified (631.8 DPS) | yes | Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Staff of Hale Magefire (13000, -22.39 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Black Dragonscale Helm; shoulder: Shroud of the Nathrezim; back: Sergeant Major's Cape; chest: Dawn Armor; wrist: Tranquil Wristguards; hands: Mourning Handguards; waist: Elunarian Belt; legs: Ironfeather Leggings; finger2: Band of the Penitent; main_hand: Forest Defender's Axe

No-known-source sample (15 of 1856, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

## Horde

### Band 20 (orc, 0000000000000000-254000000000000000-0000000000000000)

Set DPS (verified): 37.7. Weights run: 2.1s. Verify run: 1.8s. 247 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.435 ± 0.032, crit=0.971 ± 0.031 per rating point (14 rating = 1%, 13.596 per %), hit=1.573 ± 0.083 per rating point (10 rating = 1%, 15.729 per %), melee_haste=9.558 ± 0.324, intellect=1.166 ± 0.027, spell_power=0.984 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (0.57 DPS) | yes | Trapper's Leather Hood (252505, -0.24 DPS) [crafted]; Totemic Leather Hood (252448, -0.25 DPS) [crafted]; Brawler's Leather Hood (252504, -0.29 DPS, sim-verified) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 2.9 attack_power points (0.10 DPS) | yes | Roadwatcher's Confidence (281265, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 15.4 attack_power points (0.55 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.21 DPS) [crafted]; Serpent's Shoulders (5404, -0.29 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.41 DPS) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 8.6 attack_power points (0.31 DPS) | yes | Dark Leather Cloak (2316, -0.01 DPS) [crafted]; Lambent Scale Cloak (4706, -0.02 DPS) [world_drop]; Cape of the Brotherhood (5193, -0.05 DPS) [dungeon] |
| chest | Totemic Leather Armor (252435) | Leatherworking [crafted] | 23.4 attack_power points (0.83 DPS) | yes | Brawler's Leather Armor (252490, -0.12 DPS) [crafted]; Defender's Leather Armor (252434, -0.18 DPS) [crafted]; Murloc Scale Breastplate (5781, -0.32 DPS) [crafted] |
| wrist | Bristlebark Bindings (14569) | World drop [world_drop] | 8.3 attack_power points (0.30 DPS) | yes | Forest Leather Bracers (3202, -0.04 DPS) [world_drop]; Mindthrust Bracers (1974, -0.09 DPS) [dungeon]; Owl Bracers (4796, -0.09 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 16.6 attack_power points (0.59 DPS) | yes | Serpent Gloves (5970, -0.04 DPS) [dungeon]; Gold-flecked Gloves (5195, -0.09 DPS) [dungeon]; Bristlebark Gloves (14572, -0.10 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.64 DPS) | yes | Brawler's Leather Belt (252428, -0.15 DPS) [crafted]; Deviate Scale Belt (6468, -0.17 DPS) [crafted]; Totemic Leather Belt (252431, -0.22 DPS) [crafted] |
| legs | Trapper's Leather Pants (252501) | Leatherworking [crafted] | 29.6 attack_power points (1.05 DPS) | yes | Totemic Leather Pants (252446, +0.00 DPS, sim-verified) [crafted]; Brawler's Leather Pants (252500, -0.17 DPS) [crafted]; Defender's Leather Pants (252445, -0.21 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 17.5 attack_power points (0.62 DPS) | yes | Brawler's Leather Boots (252439, -0.01 DPS) [crafted]; Totemic Leather Boots (252442, -0.06 DPS) [crafted]; Trapper's Leather Boots (252440, -0.20 DPS) [crafted] |
| finger1 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 11.8 attack_power points (0.42 DPS) | yes | Demon Band (12054, -0.14 DPS) [world_drop]; The 1 Ring (8350, -0.26 DPS) [world]; Bounty Hunter's Ring (5351, -0.27 DPS) [quest] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 9.9 attack_power points (0.35 DPS) | yes | Demon Band (12054, -0.07 DPS) [world_drop]; The 1 Ring (8350, -0.19 DPS) [world]; Bounty Hunter's Ring (5351, -0.20 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 303.7 attack_power points (10.83 DPS) | yes | Forsaken Greataxe (251533, -0.03 DPS) [quest]; Hammerbone (270018, -0.13 DPS) [quest]; Twisted Chanter's Staff (890, -0.21 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Erudite's Amulet; shoulder: Magician's Mantle; back: Glowing Lizardscale Cloak; chest: Totemic Leather Armor; wrist: Bristlebark Bindings; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Trapper's Leather Pants; feet: Feet of the Lynx; finger1: Loop of Sacrifice; finger2: Lavishly Jeweled Ring; main_hand: Smite's Mighty Hammer

No-known-source sample (15 of 247, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion; 21566 Rune of Perfection

### Band 30 (orc, 0000000000000000-255130030002000000-0000000000000000)

Set DPS (verified): 94.1. Weights run: 2.1s. Verify run: 2.1s. 451 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.006, agility=1.337 ± 0.033, crit=1.070 ± 0.035 per rating point (14 rating = 1%, 14.975 per %), hit=1.950 ± 0.092 per rating point (10 rating = 1%, 19.500 per %), melee_haste=9.662 ± 0.411, intellect=1.966 ± 0.028, spell_power=0.583 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 31.5 attack_power points (1.57 DPS) | yes | Resilient Cap (14401, -0.36 DPS) [world_drop]; Defender's Leather Helm (252455, -0.37 DPS) [crafted]; Dokebi Hat (14584, -0.53 DPS, sim-verified) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 23.8 attack_power points (1.18 DPS) | yes | Kaleidoscope Chain (13084, -0.13 DPS) [world_drop]; River Pride Choker (13087, -0.79 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.79 DPS) [world_drop] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 25.1 attack_power points (1.25 DPS) | yes | Bloodmage Mantle (7684, -0.11 DPS) [dungeon]; Feline Mantle (3748, -0.22 DPS) [dungeon]; Forest Tracker Epaulets (2278, -0.70 DPS, sim-verified) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 15.4 attack_power points (0.76 DPS) | yes | Resilient Cape (14400, -0.18 DPS) [world_drop]; Tigerstrike Mantle (13108, -0.23 DPS) [world_drop]; Construct Cloak (279848, -0.27 DPS) [quest] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 41.5 attack_power points (2.07 DPS) | yes | Spirewind Fetter (9406, -0.20 DPS) [dungeon]; Green Silk Armor (7065, -0.53 DPS) [crafted]; Totemic Leather Tunic (252451, -0.58 DPS) [crafted] |
| wrist | Glowing Magical Bracelets (13106) | World drop [world_drop] | 21.7 attack_power points (1.08 DPS) | yes | Technician's Bracers (270042, -0.29 DPS) [quest]; Barbaric Bracers (18948, -0.42 DPS) [crafted]; Hawkeye's Bracers (14590, -0.75 DPS, sim-verified) [world_drop] |
| hands | Gloves of Kapelan (6744) | Alliance Relations [quest] | 25.8 attack_power points (1.28 DPS) | yes | Insignia Gloves (6408, -0.19 DPS) [world_drop]; Hotshot Pilot's Gloves (9491, -0.23 DPS) [dungeon]; Toughened Leather Gloves (4253, -0.29 DPS) [crafted] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 33.9 attack_power points (1.69 DPS) | yes | Skirmisher's Leather Belt (252461, -0.20 DPS) [crafted]; Skulker's Leather Belt (252520, -0.49 DPS) [crafted]; Stalker's Leather Belt (252521, -0.50 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 49.6 attack_power points (2.47 DPS) | yes | Trapper's Leather Pants (252501, -0.88 DPS) [crafted]; Totemic Leather Pants (252446, -0.98 DPS) [crafted]; Totemic Leather Leggings (252458, -1.16 DPS, sim-verified) [crafted] |
| feet | Acidic Walkers (9454) | Gnomeregan: Viscous Fallout [dungeon] | 18.6 attack_power points (0.93 DPS) | yes | Gilded Slippers (254001, -0.04 DPS) [crafted]; Feet of the Lynx (1121, -0.10 DPS) [world_drop]; Brawler's Leather Boots (252439, -0.10 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.0 attack_power points (1.00 DPS) | yes | Lavishly Jeweled Ring (1156, -0.28 DPS) [dungeon]; Band of the Fist (17694, -0.33 DPS) [quest]; Snake Hoop (6750, -0.33 DPS) [quest] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 15.8 attack_power points (0.79 DPS) | yes | Band of the Fist (17694, -0.12 DPS) [quest]; Snake Hoop (6750, -0.12 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.54 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (94.1 DPS) | yes | Corpsemaker (6687, -0.47 DPS) [dungeon]; Cobalt Crusher (7730, -1.99 DPS) [dungeon]; Viscous Hammer (13045, -23.10 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Archmagister's Faceted Pendant; shoulder: Death Speaker Mantle; back: Hawkeye's Cloak; chest: Garb of Florid Feathers; wrist: Glowing Magical Bracelets; hands: Gloves of Kapelan; waist: Prowler's Leather Belt; legs: Ferine Leggings; feet: Acidic Walkers; finger1: Thunderbrow Ring; finger2: Loop of Sacrifice; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 451, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape

### Band 40 (orc, 0000000000000000-255130030005102031-0000000000000000)

Set DPS (verified): 129.3. Weights run: 2.6s. Verify run: 2.3s. 688 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.292 ± 0.044, crit=1.308 ± 0.055 per rating point (14 rating = 1%, 18.309 per %), hit=1.888 ± 0.106 per rating point (10 rating = 1%, 18.879 per %), melee_haste=7.503 ± 0.814, intellect=3.498 ± 0.058, spell_power=0.474 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 63.3 attack_power points (4.44 DPS) | yes | Thinking Cap (2624, -0.27 DPS) [world]; Enduring Cap (3020, -0.51 DPS) [world_drop]; Big Voodoo Mask (8201, -0.70 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 33.0 attack_power points (2.31 DPS) | yes | Glowing Eye of Mordresh (10769, -0.12 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.41 DPS) [world_drop]; Necklace of Calisea (1714, -0.60 DPS) [dungeon] |
| shoulder | Sheepshear Mantle (13115) | World drop [world_drop] | 56.0 attack_power points (3.92 DPS) | yes | Windchaser Amice (14432, -0.46 DPS) [world_drop]; Inquisitor's Shawl (19507, -0.50 DPS) [dungeon]; Green Silken Shoulders (7057, -0.96 DPS) [crafted] |
| back | Blackforge Cape (6424) | Uldaman: Ancient Treasure [dungeon] | 28.0 attack_power points (1.96 DPS) | yes | Glyphed Cloak (4732, -0.25 DPS) [world_drop]; Resilient Cape (14400, -0.49 DPS) [world_drop]; Windchaser Cloak (14430, -0.49 DPS) [world_drop] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 64.1 attack_power points (4.49 DPS) | yes | Red Mageweave Vest (10007, -0.08 DPS) [crafted]; Spirewind Fetter (9406, -0.57 DPS) [dungeon]; Silksand Tunic (14417, -0.57 DPS) [world_drop] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 34.1 attack_power points (2.39 DPS) | yes | Glowing Magical Bracelets (13106, -0.01 DPS) [world_drop]; Warden's Wristbands (14600, -0.15 DPS) [world_drop]; Tough Scorpid Bracers (8205, -0.28 DPS) [crafted] |
| hands | Rustler Gloves (9704) | Identifying the Brood [quest] | 49.5 attack_power points (3.47 DPS) | yes | Tough Scorpid Gloves (8204, -0.36 DPS) [crafted]; Engineering Gloves (1659, -0.39 DPS) [world]; Stormcloth Gloves (10011, -0.39 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 55.8 attack_power points (3.91 DPS) | yes | Saltscale Girdle (273023, -0.44 DPS) [dungeon]; Razzeric's Customized Seatbelt (6726, -0.97 DPS) [quest]; Prowler's Leather Belt (252459, -1.13 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 68.0 attack_power points (4.77 DPS) | yes | Red Mageweave Pants (10009, -1.36 DPS) [crafted]; Misplaced Pantaloons (276201, -1.58 DPS) [vendor]; Crimson Silk Pantaloons (7062, -1.84 DPS, sim-verified) [crafted] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 48.5 attack_power points (3.40 DPS) | yes | Skirmisher's Mail Boots (252564, -0.33 DPS) [crafted]; Stalker's Mail Boots (252562, -0.69 DPS) [crafted]; Briar Tredders (10582, -0.74 DPS) [dungeon] |
| finger1 | Voodoo Band (1996) (or Ogremind Ring (1993), Mindbender Loop (5009)) | Bloodscalp Witch Doctor [world] | 24.5 attack_power points (1.72 DPS) | yes | Mindbender Loop (5009, +0.00 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.06 DPS) [dungeon]; Loop of Sacrifice (281673, -0.07 DPS) [quest] |
| finger2 | Ogremind Ring (1993) (or Mindbender Loop (5009)) | World drop [world_drop] | 24.5 attack_power points (1.72 DPS) | yes | Mindbender Loop (5009, +0.00 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.06 DPS) [dungeon]; Loop of Sacrifice (281673, -0.07 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-verified (129.3 DPS) | yes | Mograine's Might (7723, +0.00 DPS) [dungeon]; Thornstone Sledgehammer (1722, -2.47 DPS) [world_drop]; Pendulum of Doom (9425, -2.72 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; shoulder: Sheepshear Mantle; back: Blackforge Cape; wrist: Windchaser Cuffs; hands: Rustler Gloves; waist: Deathmage Sash; feet: Prowler's Leather Shoes; finger1: Voodoo Band; finger2: Ogremind Ring; main_hand: Fiery War Axe

No-known-source sample (15 of 688, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7748 Forcestone Buckler

### Band 50 (orc, 3200300000000000-255130030005102051-0000000000000000)

Set DPS (verified): 175.3. Weights run: 2.6s. Verify run: 2.6s. 879 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.397 ± 0.051, crit=1.653 ± 0.067 per rating point (14 rating = 1%, 23.145 per %), hit=3.088 ± 0.140 per rating point (10 rating = 1%, 30.879 per %), melee_haste=10.925 ± 1.128, intellect=4.307 ± 0.066, spell_power=0.780 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | 105.0 attack_power points (6.98 DPS) | yes | Red Mageweave Headband (10033, -0.27 DPS) [crafted]; Winged Helm (13112, -0.39 DPS) [world_drop]; Bad Mojo Mask (9470, -0.97 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 60.3 attack_power points (4.01 DPS) | yes | Glowing Eye of Mordresh (10769, -1.36 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.49 DPS) [world]; Kaleidoscope Chain (13084, -1.96 DPS) [world_drop] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 97.8 attack_power points (6.51 DPS) | yes | Rotgrip Mantle (17732, -0.68 DPS) [dungeon]; Rockshard Pauldrons (9411, -1.76 DPS) [dungeon]; Imperial Red Mantle (8250, -1.84 DPS) [world_drop] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 48.1 attack_power points (3.20 DPS) | yes | Imperial Red Cloak (8248, -0.05 DPS) [world_drop]; Keeper's Cloak (14665, -0.34 DPS) [world_drop]; Big Voodoo Cloak (8216, -0.36 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 101.0 attack_power points (6.71 DPS) | yes | Stone Guard's Mail Armor (220826, -0.14 DPS) [vendor]; Stone Guard's Pulsing Breastplate (220844, -0.29 DPS) [vendor]; Relentless Chain (17777, -0.84 DPS) [quest] |
| wrist | Forgotten Wraps (9433) | World drop [world_drop] | 55.9 attack_power points (3.72 DPS) | yes | Shizzle's Nozzle Wiper (11917, -0.28 DPS) [quest]; Runic Leather Bracers (15092, -0.33 DPS) [crafted]; Bracers of the Stone Princess (17714, -0.42 DPS) [dungeon] |
| hands | Stonerender Gauntlets (17007) | World drop [world_drop] | 77.2 attack_power points (5.13 DPS) | yes | Prowler's Leather Gauntlets (252547, -0.71 DPS) [crafted]; Skirmisher's Mail Gauntlets (252586, -0.98 DPS) [crafted]; First Sergeant's Mail Gauntlets (220831, -1.10 DPS) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 86.5 attack_power points (5.75 DPS) | yes | Serenity Belt (13144, -0.88 DPS) [world_drop]; Sagebrush Girdle (17778, -0.92 DPS) [quest]; Imperial Red Sash (8253, -0.92 DPS) [world_drop] |
| legs | Stone Guard's Mail Legplates (220834) | PvP rank 10 · Stone Guard · Horde [vendor] | 98.8 attack_power points (6.57 DPS) | yes | Imperial Red Pants (8251, -0.88 DPS) [world_drop]; Kilt of the Atal'ai Prophet (10807, -1.26 DPS) [dungeon]; Stone Guard's Pulsing Legplates (220847, -1.98 DPS, sim-verified) [vendor] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 82.5 attack_power points (5.49 DPS) | yes | Skycaller's Leather Boots (252471, -1.56 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -1.56 DPS) [crafted]; Rancor Boots (11865, -3.09 DPS, sim-verified) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 50.9 attack_power points (3.38 DPS) | yes | Ogremind Ring (1993, -1.38 DPS) [world_drop]; Mindbender Loop (5009, -1.38 DPS) [world_drop]; Philanthropist's Ring (281635, -1.43 DPS) [quest] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993), Mindbender Loop (5009)) | Bloodscalp Witch Doctor [world] | 30.2 attack_power points (2.00 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Mindbender Loop (5009, +0.00 DPS) [world_drop]; Philanthropist's Ring (281635, -0.05 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (175.3 DPS) | yes | Uther's Strength (11302, -3.92 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (175.3 DPS) | yes | Uther's Strength (11302, -1.54 DPS) [world_drop]; Molten Heart of the Mountain (249470, -3.33 DPS, sim-verified) [crafted] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-verified (175.3 DPS) | yes | Bleakwood Hew (12769, -3.34 DPS) [crafted]; Dark Iron Pulverizer (11608, -3.92 DPS) [crafted]; Soulkeeper (1607, -5.17 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Eye of Theradras; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Forgotten Wraps; hands: Stonerender Gauntlets; waist: Dawnspire Cord; legs: Stone Guard's Mail Legplates; feet: Greaves of Withering Despair; finger1: Blackstone Ring; finger2: Voodoo Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Thorium Greatmace

No-known-source sample (15 of 879, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders

### Band 60 (orc, 3200330000000000-255130030005102051-0520000000000000)

Set DPS (verified): 230.9. Weights run: 2.4s. Verify run: 5.0s. 1812 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.958 ± 0.076, crit=2.739 ± 0.106 per rating point (14 rating = 1%, 38.342 per %), hit=3.690 ± 0.207 per rating point (10 rating = 1%, 36.903 per %), melee_haste=15.809 ± 1.595, intellect=4.796 ± 0.095, spell_power=0.623 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 261.7 attack_power points (20.54 DPS) | yes | Magister's Crown (16686, -4.76 DPS, sim-verified) [dungeon]; Warlord's Mail Headguard (231657, -5.55 DPS) [vendor]; Warlord's Mail Helm (231663, -8.88 DPS) [pvp] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 67.1 attack_power points (5.27 DPS) | yes | Archlight Talisman (15856, -0.92 DPS) [quest]; Glowing Eye of Mordresh (10769, -1.92 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -2.07 DPS) [world] |
| shoulder | Shroud of the Nathrezim (18720) | Stratholme: Balnazzar [dungeon] | 119.4 attack_power points (9.38 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Warlord's Mail Pauldrons (231654, +0.00 DPS) [vendor]; Warlord's Mail Spaulders (231659, +0.00 DPS) [pvp] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 52.8 attack_power points (4.14 DPS) | yes | Deep Woodlands Cloak (19121, -0.17 DPS) [quest]; Cape of the Black Baron (13340, -0.27 DPS) [dungeon]; Keeper's Cloak (14665, -0.38 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 147.7 attack_power points (11.59 DPS) | yes | Warlord's Mail Armor (16577, +0.00 DPS) [vendor]; Warlord's Mail Hauberk (231653, +0.00 DPS) [vendor]; Magister's Robes (16688, -5.14 DPS, sim-verified) [dungeon] |
| wrist | Magiskull Cuffs (13107) | World drop [world_drop] | 81.5 attack_power points (6.40 DPS) | yes | Tranquil Wristguards (279256, +0.00 DPS, sim-verified) [crafted]; Spitfire Bracers (20481, -1.24 DPS) [crafted]; Hero's Bracers (8302, -1.27 DPS) [dungeon] |
| hands | Mourning Handguards (279257) | Leatherworking [crafted] | 108.3 attack_power points (8.50 DPS) | yes | Spitfire Gauntlets (20480, +0.00 DPS) [crafted]; General's Mail Vices (231655, +0.00 DPS) [vendor]; General's Mail Gauntlets (231660, +0.00 DPS) [pvp] |
| waist | Elunarian Belt (14465) | World drop [world_drop] | 107.9 attack_power points (8.47 DPS) | yes | General's Mail Waistband (16575, +0.00 DPS) [pvp]; Magister's Belt (16685, -0.22 DPS) [dungeon]; Cord of Elements (16673, -0.66 DPS) [dungeon] |
| legs | Red Dragonscale Leggings (252603) | Leatherworking [crafted] | 165.9 attack_power points (13.03 DPS) | yes | General's Mail Legguards (231658, +0.00 DPS) [vendor]; General's Mail Leggings (231664, +0.00 DPS) [pvp]; Ironfeather Leggings (252486, +0.00 DPS) [crafted] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 91.7 attack_power points (7.20 DPS) | yes | General's Mail Greaves (231656, +0.00 DPS) [vendor]; General's Mail Sabatons (231661, +0.00 DPS) [vendor]; Timmy's Galoshes (13402, -3.03 DPS, sim-verified) [dungeon] |
| finger1 | Glowing Crystal Ring (18402) | Uncovering Past Secrets [quest] | 57.6 attack_power points (4.52 DPS) | yes | Band of the Hierophant (13096, -0.75 DPS) [world_drop]; Chemist's Ring (15702, -0.75 DPS) [quest]; Frigid Ring (18679, -0.75 DPS) [world] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 56.9 attack_power points (4.47 DPS) | yes | Band of the Hierophant (13096, -0.70 DPS) [world_drop]; Chemist's Ring (15702, -0.70 DPS) [quest]; Frigid Ring (18679, -4.61 DPS, sim-verified) [world] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (230.9 DPS) | yes | Frozen Heart of the Mountain (249469, -2.63 DPS, sim-verified) [crafted]; Uther's Strength (11302, -5.03 DPS) [world_drop] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (230.9 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Forest Defender's Axe (279260) | Blacksmithing [crafted] | sim-verified (230.9 DPS) | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Staff of Hale Magefire (13000, -8.81 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Black Dragonscale Helm; shoulder: Shroud of the Nathrezim; back: Imperial Red Cloak; chest: Dawn Armor; wrist: Magiskull Cuffs; hands: Mourning Handguards; waist: Elunarian Belt; legs: Red Dragonscale Leggings; finger1: Glowing Crystal Ring; finger2: Blackstone Ring; trinket2: Molten Heart of the Mountain; main_hand: Forest Defender's Axe

No-known-source sample (15 of 1812, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders

### Band 60, raid preset (orc, 0533001000000000-253030031005102051-0530000000000000)

Set DPS (verified): 639.0. Weights run: 2.5s. Verify run: 5.1s. 1812 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.005, agility=2.339 ± 0.100, crit=2.957 ± 0.133 per rating point (14 rating = 1%, 41.402 per %), hit=5.516 ± 0.321 per rating point (10 rating = 1%, 55.157 per %), melee_haste=20.868 ± 2.640, intellect=3.392 ± 0.123, spell_power=0.878 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 219.4 attack_power points (33.21 DPS) | yes | Warlord's Mail Headguard (231657, -8.93 DPS) [vendor]; Blue Dragonscale Helm (252604, -12.30 DPS) [crafted]; Mask of the Unforgiven (13404, -22.24 DPS, sim-verified) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 47.5 attack_power points (7.19 DPS) | yes | Archlight Talisman (15856, -0.46 DPS) [quest]; Skibi's Pendant (13089, -0.92 DPS) [world_drop]; Woven Ivy Necklace (19159, -2.00 DPS) [quest] |
| shoulder | Shroud of the Nathrezim (18720) | Stratholme: Balnazzar [dungeon] | 101.8 attack_power points (15.41 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Champion's Mail Pauldrons (227154, +0.00 DPS) [pvp]; Warlord's Mail Pauldrons (231654, +0.00 DPS) [vendor] |
| back | Cloak of Flames (3475) | World drop [world_drop] | 55.2 attack_power points (8.35 DPS) | yes | Cape of the Black Baron (13340, -0.01 DPS) [dungeon]; Fel Cape (279269, -2.08 DPS) [crafted]; Deep Woodlands Cloak (19121, -2.13 DPS) [quest] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 183.1 attack_power points (27.71 DPS) | yes | Warlord's Mail Hauberk (231653, +0.00 DPS) [vendor]; Timbermaw Tunic (252484, -4.09 DPS) [crafted]; Savage Gladiator Chain (11726, -4.90 DPS) [dungeon] |
| wrist | Tranquil Wristguards (279256) | Leatherworking [crafted] | 69.1 attack_power points (10.46 DPS) | yes | Magiskull Cuffs (13107, -1.73 DPS) [world_drop]; Lordly Armguards (13135, -2.06 DPS) [world_drop]; Spitfire Bracers (20481, -5.96 DPS, sim-verified) [crafted] |
| hands | Mourning Handguards (279257) | Leatherworking [crafted] | 96.6 attack_power points (14.62 DPS) | yes | Blood Guard's Mail Vices (227159, +0.00 DPS) [pvp]; General's Mail Vices (231655, +0.00 DPS) [vendor]; Gauntlets of Elements (16672, -8.16 DPS, sim-verified) [dungeon] |
| waist | Elunarian Belt (14465) | World drop [world_drop] | 81.0 attack_power points (12.26 DPS) | yes | General's Mail Waistband (16575, +0.00 DPS) [pvp]; Might of the Timbermaw (19044, -0.32 DPS) [crafted]; Cord of Elements (16673, -0.54 DPS) [dungeon] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 135.8 attack_power points (20.56 DPS) | yes | General's Mail Legguards (231658, +0.00 DPS) [vendor]; General's Mail Leggings (231664, +0.00 DPS) [pvp]; Red Dragonscale Leggings (252603, +0.00 DPS) [crafted] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 98.7 attack_power points (14.94 DPS) | yes | Timmy's Galoshes (13402, +0.00 DPS) [dungeon]; General's Mail Greaves (231656, +0.00 DPS) [vendor]; Blood Guard's Mail Greaves (227158, -1.93 DPS) [pvp] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 75.2 attack_power points (11.37 DPS) | yes | Glowing Crystal Ring (18402, -5.21 DPS) [quest]; Maiden's Circle (13001, -5.39 DPS) [world_drop]; Frigid Ring (18679, -6.24 DPS) [world] |
| finger2 | Band of the Penitent (13217) | Houses of the Holy [quest] | 41.4 attack_power points (6.27 DPS) | yes | Glowing Crystal Ring (18402, -0.11 DPS) [quest]; Maiden's Circle (13001, -0.28 DPS) [world_drop]; Frigid Ring (18679, -1.13 DPS) [world] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (639.0 DPS) | yes | Frozen Heart of the Mountain (249469, -8.83 DPS, sim-verified) [crafted]; Uther's Strength (11302, -11.40 DPS) [world_drop] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (639.0 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Forest Defender's Axe (279260) | Blacksmithing [crafted] | sim-verified (639.0 DPS) | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Staff of Hale Magefire (13000, -26.00 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Black Dragonscale Helm; shoulder: Shroud of the Nathrezim; back: Cloak of Flames; chest: Dawn Armor; wrist: Tranquil Wristguards; hands: Mourning Handguards; waist: Elunarian Belt; legs: Ironfeather Leggings; finger2: Band of the Penitent; trinket2: Molten Heart of the Mountain; main_hand: Forest Defender's Axe

No-known-source sample (15 of 1812, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders

