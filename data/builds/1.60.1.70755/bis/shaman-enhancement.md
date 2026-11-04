# Leveling BiS: Enhancement

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-055010000000000000-0000000000000000)

Set DPS (verified): 35.0. Weights run: 5.4s. Verify run: 11.9s. 250 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.881 ± 0.021, crit=0.686 ± 0.021 per rating point (14 rating = 1%, 9.600 per %), hit=1.197 ± 0.055 per rating point (10 rating = 1%, 11.969 per %), melee_haste=6.382 ± 0.207, intellect=1.196 ± 0.023, spell_power=0.867 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (0.57 DPS) | yes | Trapper's Leather Hood (252505, -0.23 DPS) [crafted]; Totemic Leather Hood (252448, -0.29 DPS) [crafted]; Brawler's Leather Hood (252504, -0.32 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 5.3 attack_power points (0.19 DPS) | yes | Erudite's Amulet (277204, -0.13 DPS) [quest]; Tarnished Locket (279870, -0.40 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 15.1 attack_power points (0.54 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.21 DPS) [crafted]; Serpent's Shoulders (5404, -0.38 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.42 DPS) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | sim-verified (35.0 DPS) | yes | Dark Leather Cloak (2316, -0.05 DPS) [crafted]; Grave Shroud (279865, -0.07 DPS) [quest]; Pearl-clasped Cloak (5542, -0.10 DPS) [crafted] |
| chest | Totemic Leather Armor (252435) | Leatherworking [crafted] | 22.8 attack_power points (0.82 DPS) | yes | Defender's Leather Armor (252434, -0.22 DPS) [crafted]; Brawler's Leather Armor (252490, -0.24 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.35 DPS) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 8.6 attack_power points (0.31 DPS) | yes | Bristlebark Bindings (14569, -0.07 DPS) [world_drop]; Mindthrust Bracers (1974, -0.10 DPS) [dungeon]; Owl Bracers (4796, -0.10 DPS) [vendor] |
| hands | Gold-flecked Gloves (5195) | The Deadmines: Sneed [dungeon] | 14.0 attack_power points (0.50 DPS) | yes | Gloves of the Fang (10413, -0.03 DPS) [dungeon]; Totemic Leather Gloves (252497, -0.06 DPS) [crafted]; Blackened Defias Gloves (10401, -0.07 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.64 DPS) | yes | Brawler's Leather Belt (252428, -0.22 DPS, sim-verified) [crafted]; Totemic Leather Belt (252431, -0.23 DPS) [crafted]; Deviate Scale Belt (6468, -0.27 DPS) [crafted] |
| legs | Totemic Leather Pants (252446) | Leatherworking [crafted] | 28.9 attack_power points (1.03 DPS) | yes | Trapper's Leather Pants (252501, -0.15 DPS) [crafted]; Defender's Leather Pants (252445, -0.26 DPS) [crafted]; Brawler's Leather Pants (252500, -0.32 DPS) [crafted] |
| feet | Totemic Leather Boots (252442) | Leatherworking [crafted] | 15.2 attack_power points (0.54 DPS) | yes | Brawler's Leather Boots (252439, -0.03 DPS) [crafted]; Feet of the Lynx (1121, -0.08 DPS) [world_drop]; Stormrider's Leather Boots (252443, -0.14 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 11.5 attack_power points (0.41 DPS) | yes | Demon Band (12054, -0.13 DPS) [world_drop]; Minor Channeling Ring (1449, -0.17 DPS) [quest]; Lorekeeper's Ring (20431, -0.26 DPS) [rep] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 8.9 attack_power points (0.32 DPS) | yes | Demon Band (12054, -0.03 DPS) [world_drop]; Minor Channeling Ring (1449, -0.08 DPS) [quest]; Lorekeeper's Ring (20431, -0.16 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 301.5 attack_power points (10.80 DPS) | yes | Twisted Chanter's Staff (890, -0.17 DPS, sim-verified) [world_drop]; Rockslicer (872, -0.42 DPS) [dungeon]; Night Reaver (1318, -0.92 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Sentinel's Medallion; shoulder: Magician's Mantle; back: Lambent Scale Cloak; chest: Totemic Leather Armor; wrist: Bravo's Armbands; hands: Gold-flecked Gloves; waist: Blackened Defias Belt; legs: Totemic Leather Pants; feet: Totemic Leather Boots; finger1: Protector's Band; finger2: Lavishly Jeweled Ring; main_hand: Smite's Mighty Hammer

No-known-source sample (15 of 250, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18845 Insignia of the Horde

### Band 30 (dwarf, 0000000000000000-055030310040000000-0000000000000000)

Set DPS (verified): 77.9. Weights run: 5.5s. Verify run: 9.2s. 434 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.006, agility=1.447 ± 0.038, crit=1.327 ± 0.044 per rating point (14 rating = 1%, 18.580 per %), hit=2.038 ± 0.097 per rating point (10 rating = 1%, 20.378 per %), melee_haste=9.035 ± 0.357, intellect=2.551 ± 0.033, spell_power=0.812 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 40.8 attack_power points (1.99 DPS) | yes | Dokebi Hat (14584, -0.47 DPS, sim-verified) [world_drop]; Resilient Cap (14401, -0.48 DPS) [world_drop]; Nightsky Cowl (4039, -0.50 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 27.3 attack_power points (1.33 DPS) | yes | Kaleidoscope Chain (13084, -0.16 DPS) [world_drop]; Seal of the Expedition (271920, -0.71 DPS) [rep]; Sentinel's Medallion (19541, -0.77 DPS) [rep] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 32.9 attack_power points (1.60 DPS) | yes | Bloodmage Mantle (7684, -0.13 DPS) [dungeon]; Magician's Mantle (12998, -0.29 DPS) [world_drop]; Feline Mantle (3748, -0.31 DPS) [dungeon] |
| back | Sergeant Major's Cape (16315) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-verified (77.9 DPS) | yes | Expeditionary's Cape (271923, -0.17 DPS) [rep]; Hawkeye's Cloak (14593, -0.38 DPS) [world_drop]; Resilient Cape (14400, -0.42 DPS) [world_drop] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 51.3 attack_power points (2.50 DPS) | yes | Green Silk Armor (7065, -0.53 DPS) [crafted]; Spirewind Fetter (9406, -0.53 DPS, sim-verified) [dungeon]; Beguiler Robes (7728, -0.65 DPS) [dungeon] |
| wrist | Glowing Magical Bracelets (13106) | World drop [world_drop] | 26.4 attack_power points (1.29 DPS) | yes | Hawkeye's Bracers (14590, -0.47 DPS) [world_drop]; Technician's Bracers (270042, -0.54 DPS, sim-verified) [quest]; Nightsky Wristbands (6407, -0.54 DPS) [world_drop] |
| hands | Hotshot Pilot's Gloves (9491) | Gnomeregan [dungeon] | 26.2 attack_power points (1.28 DPS) | yes | Toughened Leather Gloves (4253, -0.27 DPS) [crafted]; Town Clerk's Mittens (270029, -0.29 DPS) [quest]; Insignia Gloves (6408, -0.46 DPS, sim-verified) [world_drop] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 36.9 attack_power points (1.80 DPS) | yes | Highlander's Lizardhide Girdle (20105, -0.31 DPS) [rep]; Highlander's Mail Girdle (20120, -0.31 DPS) [vendor]; Skirmisher's Leather Belt (252461, -0.62 DPS, sim-verified) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 56.6 attack_power points (2.76 DPS) | yes | Totemic Leather Leggings (252458, -1.00 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, -1.04 DPS) [crafted]; Totemic Leather Pants (252446, -1.11 DPS) [crafted] |
| feet | Acidic Walkers (9454) | Gnomeregan: Viscous Fallout [dungeon] | 24.5 attack_power points (1.19 DPS) | yes | Gilded Slippers (254001, -0.05 DPS) [crafted]; Frothing Slippers (254003, -0.32 DPS) [crafted]; Fiery Slippers (254005, -0.32 DPS) [crafted] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 20.7 attack_power points (1.01 DPS) | yes | Thunderbrow Ring (13097, -0.02 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.12 DPS) [dungeon]; Snake Hoop (6750, -0.15 DPS) [quest] |
| finger2 | Seal of Wrynn (2933) | An Audience with the King [quest] | 20.5 attack_power points (1.00 DPS) | yes | Thunderbrow Ring (13097, -0.01 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.11 DPS) [dungeon]; Snake Hoop (6750, -0.14 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (77.9 DPS) | yes | Corpsemaker (6687, -0.49 DPS) [dungeon]; Cobalt Crusher (7730, -1.98 DPS) [dungeon]; Viscous Hammer (13045, -4.19 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Archmagister's Faceted Pendant; shoulder: Death Speaker Mantle; back: Sergeant Major's Cape; chest: Garb of Florid Feathers; wrist: Glowing Magical Bracelets; hands: Hotshot Pilot's Gloves; waist: Prowler's Leather Belt; legs: Ferine Leggings; feet: Acidic Walkers; finger1: Protector's Band; finger2: Seal of Wrynn; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 434, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 40 (dwarf, 0000000000000000-055030310051202031-0000000000000000)

Set DPS (verified): 154.3. Weights run: 6.0s. Verify run: 16.9s. 661 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.504 ± 0.055, crit=1.675 ± 0.066 per rating point (14 rating = 1%, 23.454 per %), hit=2.568 ± 0.122 per rating point (10 rating = 1%, 25.677 per %), melee_haste=6.948 ± 0.943, intellect=3.230 ± 0.040, spell_power=0.645 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 60.1 attack_power points (5.25 DPS) | yes | Enduring Cap (3020, -0.73 DPS) [world_drop]; Big Voodoo Mask (8201, -0.79 DPS) [crafted]; Thinking Cap (2624, -1.18 DPS, sim-verified) [world] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 31.4 attack_power points (2.74 DPS) | yes | Glowing Eye of Mordresh (10769, -0.09 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.39 DPS) [world_drop]; Seal of the Expedition (271919, -0.77 DPS) [rep] |
| shoulder | Sheepshear Mantle (13115) | World drop [world_drop] | 51.7 attack_power points (4.52 DPS) | yes | Windchaser Amice (14432, -0.45 DPS) [world_drop]; Green Silken Shoulders (7057, -0.96 DPS) [crafted]; Inquisitor's Shawl (19507, -1.59 DPS, sim-verified) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Expeditionary's Cape (271922, -0.43 DPS) [rep]; Blackforge Cape (6424, -1.27 DPS) [dungeon] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 62.0 attack_power points (5.42 DPS) | yes | Red Mageweave Vest (10007, -0.34 DPS) [crafted]; Spirewind Fetter (9406, -0.81 DPS) [dungeon]; Icemail Jerkin (1981, -1.22 DPS, sim-verified) [world_drop] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 32.1 attack_power points (2.80 DPS) | yes | Glowing Magical Bracelets (13106, -0.02 DPS) [world_drop]; Warden's Wristbands (14600, -0.15 DPS) [world_drop]; Tough Scorpid Bracers (8205, -0.19 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Engineering Gloves (1659, +0.00 DPS) [world]; Stormcloth Gloves (10011, +0.00 DPS) [crafted]; Tough Scorpid Gloves (8204, -1.74 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 53.0 attack_power points (4.63 DPS) | yes | Highlander's Lizardhide Girdle (20104, -0.39 DPS) [rep]; Highlander's Mail Girdle (20119, -0.39 DPS) [vendor]; Saltscale Girdle (273023, -0.52 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 64.8 attack_power points (5.66 DPS) | yes | Crimson Silk Pantaloons (7062, -1.54 DPS) [crafted]; Red Mageweave Pants (10009, -1.64 DPS, sim-verified) [crafted]; Totemic Leather Leggings (252458, -1.94 DPS) [crafted] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 48.7 attack_power points (4.25 DPS) | yes | Skirmisher's Mail Boots (252564, -0.41 DPS) [crafted]; Stalker's Mail Boots (252562, -0.83 DPS) [crafted]; Briar Tredders (10582, -1.06 DPS) [dungeon] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 28.0 attack_power points (2.45 DPS) | yes | Ogremind Ring (1993, -0.47 DPS) [world_drop]; Voodoo Band (1996, -0.47 DPS) [world]; Mindbender Loop (5009, -0.47 DPS) [world_drop] |
| finger2 | Seal of Wrynn (2933) | An Audience with the King [quest] | 23.4 attack_power points (2.05 DPS) | yes | Ogremind Ring (1993, -0.07 DPS) [world_drop]; Voodoo Band (1996, -0.07 DPS) [world]; Mindbender Loop (5009, -0.07 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mograine's Might (7723, +0.00 DPS) [dungeon]; Bonebiter (6830, -1.08 DPS) [quest]; Pendulum of Doom (9425, -6.82 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; shoulder: Sheepshear Mantle; back: Sergeant Major's Cape; wrist: Windchaser Cuffs; hands: Gloves of Holy Might; waist: Deathmage Sash; feet: Prowler's Leather Shoes; finger1: Protector's Band; trinket1: Ankh of Life; main_hand: Fiery War Axe

No-known-source sample (15 of 661, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 50 (dwarf, 3200300000000000-055030310051202051-0000000000000000)

Set DPS (verified): 217.6. Weights run: 6.3s. Verify run: 17.1s. 848 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.415 ± 0.064, crit=1.838 ± 0.085 per rating point (14 rating = 1%, 25.729 per %), hit=3.301 ± 0.169 per rating point (10 rating = 1%, 33.006 per %), melee_haste=9.842 ± 1.399, intellect=3.315 ± 0.059, spell_power=0.587 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | 88.7 attack_power points (8.99 DPS) | yes | Red Mageweave Headband (10033, -1.14 DPS) [crafted]; Winged Helm (13112, -1.26 DPS) [world_drop]; Knight-Lieutenant's Mail Helmet (223075, -1.80 DPS) [vendor] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 46.4 attack_power points (4.71 DPS) | yes | Glowing Eye of Mordresh (10769, -1.60 DPS) [dungeon]; Seal of the Expedition (271918, -1.68 DPS) [rep]; Archmagister's Faceted Pendant (273049, -3.66 DPS, sim-verified) [world] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 75.1 attack_power points (7.61 DPS) | yes | Rotgrip Mantle (17732, -0.79 DPS) [dungeon]; Knight-Lieutenant's Mail Epaulets (223073, -1.50 DPS) [vendor]; Rockshard Pauldrons (9411, -1.92 DPS) [dungeon] |
| back | Expeditionary's Cape (271921) | Theramore Expeditionary Force [rep] | sim-verified (217.6 DPS) | yes | Imperial Red Cloak (8248, -1.01 DPS) [world_drop]; Keeper's Cloak (14665, -1.34 DPS) [world_drop]; Sergeant Major's Cape (16336, -2.64 DPS, sim-verified) [pvp] |
| chest | Knight's Mail Armor (223078) | Captain Dirgehammer [vendor] | 89.5 attack_power points (9.07 DPS) | yes | Acumen Robes (17775, -1.22 DPS) [quest]; Relentless Chain (17777, -1.53 DPS) [quest]; Icemail Jerkin (1981, -2.29 DPS) [world_drop] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 44.6 attack_power points (4.52 DPS) | yes | Prowler's Leather Bracers (252539, -0.01 DPS) [crafted]; Forgotten Wraps (9433, -0.06 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -0.49 DPS) [quest] |
| hands | Stonerender Gauntlets (17007) | World drop [world_drop] | sim-verified (217.6 DPS) | yes | Prowler's Leather Gauntlets (252547, -0.80 DPS) [crafted]; Sergeant Major's Mail Gauntlets (223076, -1.03 DPS) [vendor]; Skirmisher's Mail Gauntlets (252586, -1.43 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 66.5 attack_power points (6.74 DPS) | yes | Highlander's Lizardhide Girdle (20103, -0.77 DPS) [rep]; Highlander's Mail Girdle (20118, -0.77 DPS) [vendor]; Sagebrush Girdle (17778, -2.86 DPS, sim-verified) [quest] |
| legs | Knight's Mail Legplates (223074) | Captain Dirgehammer [vendor] | 89.5 attack_power points (9.07 DPS) | yes | Imperial Red Pants (8251, -1.61 DPS, sim-verified) [world_drop]; Ferine Leggings (6690, -2.41 DPS) [dungeon]; Searingscale Leggings (11749, -2.84 DPS) [dungeon] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 72.6 attack_power points (7.36 DPS) | yes | Prowler's Leather Shoes (252465, -2.45 DPS) [crafted]; Stalker's Mail Sabatons (252576, -2.47 DPS) [crafted]; Rancor Boots (11865, -3.39 DPS, sim-verified) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 53.0 attack_power points (5.37 DPS) | yes | Seal of Wrynn (2933, -2.99 DPS) [quest]; Ogremind Ring (1993, -3.02 DPS) [world_drop]; Voodoo Band (1996, -3.02 DPS) [world] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 32.7 attack_power points (3.32 DPS) | yes | Seal of Wrynn (2933, -0.94 DPS) [quest]; Ogremind Ring (1993, -0.97 DPS) [world_drop]; Voodoo Band (1996, -0.97 DPS) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (217.6 DPS) | yes | Uther's Strength (11302, -2.65 DPS) [world_drop] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (217.6 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Molten Heart of the Mountain (249470, -2.36 DPS, sim-verified) [crafted] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-verified (217.6 DPS) | yes | Soulkeeper (1607, -3.04 DPS) [world_drop]; Bleakwood Hew (12769, -4.64 DPS) [crafted]; Executioner's Cleaver (13018, -5.45 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Eye of Theradras; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Expeditionary's Cape; chest: Knight's Mail Armor; wrist: Bracers of the Stone Princess; hands: Stonerender Gauntlets; waist: Dawnspire Cord; legs: Knight's Mail Legplates; feet: Greaves of Withering Despair; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Frozen Heart of the Mountain; trinket2: Thunderbrew's Boot Flask; main_hand: Thorium Greatmace

No-known-source sample (15 of 848, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60 (dwarf, 3200330000000000-055030310051202051-0520000000000000)

Set DPS (verified): 313.6. Weights run: 5.8s. Verify run: 55.3s. 1668 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.226 ± 0.058, crit=1.963 ± 0.083 per rating point (14 rating = 1%, 27.476 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=12.029 ± 1.368, intellect=2.797 ± 0.056, spell_power=0.647 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 159.5 attack_power points (18.99 DPS) | yes | Coif of Elements (16667, -8.74 DPS) [dungeon]; Eye of Theradras (17715, -9.39 DPS) [dungeon]; Magister's Crown (16686, -10.49 DPS, sim-verified) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 39.2 attack_power points (4.66 DPS) | yes | Opaline Medallion (16623, -0.33 DPS) [quest]; Archlight Talisman (15856, -0.41 DPS) [quest]; Archmagister's Faceted Pendant (273049, -1.24 DPS) [world] |
| shoulder | Theramore Shoulderguards (272957) | Theramore Expeditionary Force [rep] | 85.9 attack_power points (10.22 DPS) | yes | Theramore Shoulderpads (271953, -0.31 DPS) [rep]; Theramore Shoulders (271954, -0.31 DPS) [rep]; Highlander's Mail Pauldrons (20056, -0.48 DPS) [vendor] |
| back | Sergeant Major's Cape (16337) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-verified (313.6 DPS) | yes | Expeditionary's Cape (271907, -0.41 DPS) [rep]; Cloak of the Honor Guard (20073, -0.96 DPS) [rep] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 115.5 attack_power points (13.76 DPS) | yes | Magister's Robes (16688, -3.51 DPS) [dungeon]; Knight's Mail Armor (223078, -3.63 DPS) [vendor]; Savage Gladiator Chain (11726, -3.98 DPS, sim-verified) [dungeon] |
| wrist | Windtalker's Wristguards (19582) | Silverwing Sentinels [rep] | 60.4 attack_power points (7.19 DPS) | yes | Tranquil Wristguards (279256, -0.05 DPS) [crafted]; Magiskull Cuffs (13107, -1.53 DPS) [world_drop]; Forest Stalker's Bracers (19587, -1.80 DPS) [rep] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (313.6 DPS) | yes | Theramore Handwraps (271946, +0.00 DPS) [rep]; Theramore Handguards (271950, +0.00 DPS) [rep] |
| waist | Highlander's Mail Girdle (20044) (or Highlander's Lizardhide Girdle (20046)) | Samuel Hawke [vendor] | 75.0 attack_power points (8.93 DPS) | yes | Highlander's Lizardhide Girdle (20046, +0.00 DPS) [rep]; Elunarian Belt (14465, -0.84 DPS) [world_drop]; Cord of Elements (16673, -1.13 DPS) [dungeon] |
| legs | Sentinel's Lizardhide Pants (22750) | Silverwing Sentinels [rep] | 124.9 attack_power points (14.87 DPS) | yes | Red Dragonscale Leggings (252603, -2.67 DPS) [crafted]; Ironfeather Leggings (252486, -2.68 DPS) [crafted]; Sentinel's Silk Leggings (237815, -4.72 DPS, sim-verified) [vendor] |
| feet | Timmy's Galoshes (13402) | Stratholme: Timmy the Cruel [dungeon] | 48.8 attack_power points (5.81 DPS) | yes | Highlander's Lizardhide Boots (20053, -0.07 DPS) [rep]; Scalegut Treaders (275618, -0.12 DPS) [crafted]; Omnicast Boots (11822, -0.28 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 41.5 attack_power points (4.94 DPS) | yes | Deep Rooted Ring (19109, -0.68 DPS) [world]; Glowing Crystal Ring (18402, -0.94 DPS) [quest]; Maiden's Circle (13001, -1.22 DPS) [world_drop] |
| finger2 | Protector's Band (19514) | Silverwing Sentinels [rep] | 37.5 attack_power points (4.46 DPS) | yes | Glowing Crystal Ring (18402, -0.47 DPS) [quest]; Maiden's Circle (13001, -0.75 DPS) [world_drop]; Deep Rooted Ring (19109, -1.90 DPS, sim-verified) [world] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (313.6 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (313.6 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Molten Heart of the Mountain (249470, -3.93 DPS, sim-verified) [crafted] |
| main_hand | Forest Defender's Axe (279260) | Blacksmithing [crafted] | sim-verified (313.6 DPS) | yes | The Unstoppable Force (19323, +0.00 DPS) [rep]; Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Black Dragonscale Helm; shoulder: Theramore Shoulderguards; back: Sergeant Major's Cape; chest: Timbermaw Tunic; wrist: Windtalker's Wristguards; hands: Gloves of Holy Might; waist: Highlander's Mail Girdle; legs: Sentinel's Lizardhide Pants; feet: Timmy's Galoshes; finger1: Don Julio's Band; finger2: Protector's Band; trinket1: Thunderbrew's Boot Flask; trinket2: Frozen Heart of the Mountain; main_hand: Forest Defender's Axe

No-known-source sample (15 of 1668, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60, raid preset (dwarf, 0533001000000000-053030310051202051-0530000000000000)

Set DPS (verified): 822.8. Weights run: 5.4s. Verify run: 60.0s. 1668 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.332 ± 0.052, crit=2.158 ± 0.076 per rating point (14 rating = 1%, 30.210 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=14.786 ± 1.752, intellect=2.195 ± 0.019, spell_power=0.694 ± 0.001

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 141.6 attack_power points (32.30 DPS) | yes | Coif of Elements (16667, -14.44 DPS, sim-verified) [dungeon]; Ragefury Eyepatch (11735, -15.50 DPS) [dungeon]; Magister's Crown (16686, -15.53 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 30.7 attack_power points (7.01 DPS) | yes | Archlight Talisman (15856, -0.10 DPS) [quest]; Opaline Medallion (16623, -0.50 DPS) [quest]; Skibi's Pendant (13089, -0.55 DPS) [world_drop] |
| shoulder | Highlander's Mail Pauldrons (20056) | Samuel Hawke [vendor] | 74.8 attack_power points (17.07 DPS) | yes | Highlander's Lizardhide Shoulders (20060, -0.57 DPS) [rep]; Mantle of the Timbermaw (19050, -0.98 DPS) [crafted]; Theramore Shoulderguards (272957, -1.02 DPS) [rep] |
| back | Sergeant Major's Cape (16337) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of the Honor Guard (20073, -1.18 DPS) [rep]; Cape of the Black Baron (13340, -1.33 DPS) [dungeon]; Expeditionary's Cape (271907, -2.44 DPS) [rep] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 126.8 attack_power points (28.92 DPS) | yes | Savage Gladiator Chain (11726, -2.85 DPS) [dungeon]; Knight's Mail Armor (223078, -10.00 DPS) [vendor]; Dawn Armor (252483, -10.79 DPS) [crafted] |
| wrist | Tranquil Wristguards (279256) | Leatherworking [crafted] | 57.2 attack_power points (13.04 DPS) | yes | Windtalker's Wristguards (19582, -0.36 DPS) [rep]; Forest Stalker's Bracers (19587, -1.75 DPS) [rep] |
| hands | Timbermaw Brawlers (19049) | Leatherworking [crafted] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Theramore Handwraps (271945, -0.01 DPS) [rep]; Theramore Gloves (271948, -1.07 DPS) [rep]; Stonerender Gauntlets (17007, -1.85 DPS, sim-verified) [world_drop] |
| waist | Highlander's Mail Girdle (20044) (or Highlander's Lizardhide Girdle (20046)) | Samuel Hawke [vendor] | 67.5 attack_power points (15.40 DPS) | yes | Highlander's Lizardhide Girdle (20046, +0.00 DPS) [rep]; Might of the Timbermaw (19044, -0.61 DPS) [crafted]; Highlander's Chain Girdle (20043, -0.76 DPS) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | sim-verified (+8.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Sentinel's Silk Leggings (237815, -0.56 DPS) [vendor]; Sentinel's Lizardhide Pants (237817, -0.56 DPS) [vendor] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 48.7 attack_power points (11.10 DPS) | yes | Timmy's Galoshes (13402, -0.93 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.95 DPS) [crafted]; Highlander's Lizardhide Boots (20053, -1.01 DPS) [rep] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 44.2 attack_power points (10.08 DPS) | yes | Deep Rooted Ring (19109, -3.18 DPS) [world]; Band of the Penitent (13217, -3.19 DPS) [quest]; Don Rodrigo's Band (21563, -3.19 DPS) [rep] |
| finger2 | Protector's Band (19514) | Silverwing Sentinels [rep] | 41.0 attack_power points (9.36 DPS) | yes | Deep Rooted Ring (19109, -2.46 DPS) [world]; Band of the Penitent (13217, -2.47 DPS) [quest]; Don Rodrigo's Band (21563, -2.47 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+7.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Forest Defender's Axe (279260) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Unstoppable Force (19323, +0.00 DPS) [rep]; Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Black Dragonscale Helm; shoulder: Highlander's Mail Pauldrons; back: Sergeant Major's Cape; chest: Timbermaw Tunic; wrist: Tranquil Wristguards; hands: Timbermaw Brawlers; waist: Highlander's Mail Girdle; legs: Sentinel's Chain Leggings; feet: Scalegut Treaders; finger1: Don Julio's Band; finger2: Protector's Band; trinket2: Molten Heart of the Mountain; main_hand: Forest Defender's Axe

No-known-source sample (15 of 1668, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

## Horde

### Band 20 (orc, 0000000000000000-055010000000000000-0000000000000000)

Set DPS (verified): 35.6. Weights run: 5.4s. Verify run: 13.4s. 239 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.214 ± 0.030, crit=0.937 ± 0.029 per rating point (14 rating = 1%, 13.120 per %), hit=1.608 ± 0.077 per rating point (10 rating = 1%, 16.081 per %), melee_haste=8.539 ± 0.281, intellect=1.186 ± 0.023, spell_power=0.837 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (0.55 DPS) | yes | Trapper's Leather Hood (252505, -0.22 DPS) [crafted]; Totemic Leather Hood (252448, -0.29 DPS) [crafted]; Brawler's Leather Hood (252504, -0.35 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 7.3 attack_power points (0.25 DPS) | yes | Erudite's Amulet (277204, -0.17 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 14.9 attack_power points (0.51 DPS) | yes | Serpent's Shoulders (5404, -0.30 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.40 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.45 DPS, sim-verified) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | sim-verified (35.6 DPS) | yes | Glowing Lizardscale Cloak (6449, -0.02 DPS) [dungeon]; Cape of the Brotherhood (5193, -0.07 DPS) [dungeon]; Dark Leather Cloak (2316, -0.35 DPS, sim-verified) [crafted] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (35.6 DPS) | yes | Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.72 DPS, sim-verified) [crafted] |
| wrist | Bristlebark Bindings (14569) | World drop [world_drop] | 7.6 attack_power points (0.26 DPS) | yes | Forest Leather Bracers (3202, -0.05 DPS) [world_drop]; Mindthrust Bracers (1974, -0.06 DPS) [dungeon]; Owl Bracers (4796, -0.06 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (35.6 DPS) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Serpent Gloves (5970, +0.00 DPS) [dungeon]; Gloves of the Fang (10413, -1.14 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.62 DPS) | yes | Deviate Scale Belt (6468, -0.20 DPS) [crafted]; Totemic Leather Belt (252431, -0.23 DPS) [crafted]; Brawler's Leather Belt (252428, -1.42 DPS, sim-verified) [crafted] |
| legs | Totemic Leather Pants (252446) | Leatherworking [crafted] | 28.6 attack_power points (0.98 DPS) | yes | Brawler's Leather Pants (252500, -0.20 DPS) [crafted]; Defender's Leather Pants (252445, -0.20 DPS) [crafted]; Trapper's Leather Pants (252501, -0.24 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (35.6 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Totemic Leather Boots (252442, +0.00 DPS) [crafted]; Brawler's Leather Boots (252439, -0.88 DPS, sim-verified) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 12.9 attack_power points (0.44 DPS) | yes | Lavishly Jeweled Ring (1156, -0.11 DPS) [dungeon]; Demon Band (12054, -0.17 DPS) [world_drop]; The 1 Ring (8350, -0.29 DPS) [world] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 11.9 attack_power points (0.41 DPS) | yes | Demon Band (12054, -0.14 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.18 DPS, sim-verified) [dungeon]; The 1 Ring (8350, -0.26 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Forsaken Greataxe (251533) | The Wrath of Rath'mael [quest] | 303.0 attack_power points (10.43 DPS) | yes | Hammerbone (270018, -0.10 DPS) [quest]; Twisted Chanter's Staff (890, -0.17 DPS) [world_drop]; Smite's Mighty Hammer (7230, -0.44 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Scout's Medallion; shoulder: Magician's Mantle; back: Lambent Scale Cloak; chest: Blackened Defias Armor; wrist: Bristlebark Bindings; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Totemic Leather Pants; feet: Blackened Defias Boots; finger1: Legionnaire's Band; finger2: Loop of Sacrifice; main_hand: Forsaken Greataxe

No-known-source sample (15 of 239, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 20425 Advisor's Gnarled Staff; 20441 Scout's Blade; 23173 Abomination Skin Leggings; 206376 Bulwark Icon; 206383 Lithic Icon; 206384 Astral Icon; 206385 Clastic Icon

### Band 30 (orc, 0000000000000000-055030310040000000-0000000000000000)

Set DPS (verified): 77.3. Weights run: 5.5s. Verify run: 10.5s. 428 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.005, agility=1.471 ± 0.034, crit=1.376 ± 0.038 per rating point (14 rating = 1%, 19.270 per %), hit=2.334 ± 0.095 per rating point (10 rating = 1%, 23.343 per %), melee_haste=9.824 ± 0.322, intellect=2.439 ± 0.027, spell_power=1.031 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 39.0 attack_power points (2.00 DPS) | yes | Dokebi Hat (14584, -0.22 DPS) [world_drop]; Enchanter's Cowl (4322, -0.43 DPS) [crafted]; Resilient Cap (14401, -0.47 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 26.6 attack_power points (1.37 DPS) | yes | Kaleidoscope Chain (13084, -0.15 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.74 DPS) [vendor]; Scout's Medallion (19537, -0.76 DPS) [rep] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 33.0 attack_power points (1.69 DPS) | yes | Bloodmage Mantle (7684, -0.09 DPS) [dungeon]; Magician's Mantle (12998, -0.30 DPS) [world_drop]; Feline Mantle (3748, -0.33 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-verified (77.3 DPS) | yes | Darkspear Raider's Cloak (272078, +0.00 DPS) [vendor]; Resilient Cape (14400, -0.09 DPS) [world_drop]; Sporid Cape (6629, -1.19 DPS, sim-verified) [dungeon] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 49.8 attack_power points (2.56 DPS) | yes | Spirewind Fetter (9406, -0.34 DPS) [dungeon]; Green Silk Armor (7065, -0.45 DPS) [crafted]; Beguiler Robes (7728, -0.58 DPS) [dungeon] |
| wrist | Glowing Magical Bracelets (13106) | World drop [world_drop] | 25.5 attack_power points (1.31 DPS) | yes | Technician's Bracers (270042, -0.40 DPS) [quest]; Hawkeye's Bracers (14590, -0.45 DPS) [world_drop]; Nightsky Wristbands (6407, -0.56 DPS) [world_drop] |
| hands | Gloves of Kapelan (6744) | Alliance Relations [quest] | 28.6 attack_power points (1.47 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.17 DPS) [dungeon]; Insignia Gloves (6408, -0.30 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.40 DPS) [crafted] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 36.6 attack_power points (1.88 DPS) | yes | Skirmisher's Leather Belt (252461, -0.08 DPS) [crafted]; Defiler's Lizardhide Girdle (20172, -0.38 DPS) [rep]; Defiler's Mail Girdle (20197, -0.38 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 55.3 attack_power points (2.84 DPS) | yes | Totemic Leather Leggings (252458, -0.81 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, -1.04 DPS) [crafted]; Totemic Leather Pants (252446, -1.04 DPS) [crafted] |
| feet | Acidic Walkers (9454) | Gnomeregan: Viscous Fallout [dungeon] | 24.7 attack_power points (1.27 DPS) | yes | Gilded Slippers (254001, -0.02 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.32 DPS) [crafted]; Feet of the Lynx (1121, -0.35 DPS) [world_drop] |
| finger1 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 20.8 attack_power points (1.07 DPS) | yes | Snake Hoop (6750, -0.13 DPS) [quest]; Loop of Sacrifice (281673, -0.13 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.17 DPS) [dungeon] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.4 attack_power points (1.05 DPS) | yes | Snake Hoop (6750, -0.10 DPS) [quest]; Loop of Sacrifice (281673, -0.11 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.15 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (77.3 DPS) | yes | Corpsemaker (6687, -0.52 DPS) [dungeon]; Cobalt Crusher (7730, -2.09 DPS) [dungeon]; Viscous Hammer (13045, -3.91 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Archmagister's Faceted Pendant; shoulder: Death Speaker Mantle; back: Hawkeye's Cloak; chest: Garb of Florid Feathers; wrist: Glowing Magical Bracelets; hands: Gloves of Kapelan; waist: Prowler's Leather Belt; legs: Ferine Leggings; feet: Acidic Walkers; finger1: Legionnaire's Band; finger2: Thunderbrow Ring; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 428, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape

### Band 40 (orc, 0000000000000000-055030310051202031-0000000000000000)

Set DPS (verified): 155.1. Weights run: 6.0s. Verify run: 15.5s. 644 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.504 ± 0.055, crit=1.675 ± 0.066 per rating point (14 rating = 1%, 23.454 per %), hit=2.568 ± 0.122 per rating point (10 rating = 1%, 25.677 per %), melee_haste=6.948 ± 0.943, intellect=3.230 ± 0.040, spell_power=0.645 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 60.1 attack_power points (5.25 DPS) | yes | Enduring Cap (3020, -0.73 DPS) [world_drop]; Big Voodoo Mask (8201, -0.79 DPS) [crafted]; Thinking Cap (2624, -1.48 DPS, sim-verified) [world] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 31.4 attack_power points (2.74 DPS) | yes | Glowing Eye of Mordresh (10769, -0.09 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.39 DPS) [world_drop]; Darkspear Warding Pendant (272074, -0.77 DPS) [vendor] |
| shoulder | Sheepshear Mantle (13115) | World drop [world_drop] | 51.7 attack_power points (4.52 DPS) | yes | Windchaser Amice (14432, -0.45 DPS) [world_drop]; Green Silken Shoulders (7057, -0.96 DPS) [crafted]; Inquisitor's Shawl (19507, -1.57 DPS, sim-verified) [dungeon] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | sim-verified (155.1 DPS) | yes | Glyphed Cloak (4732, -1.13 DPS) [world_drop]; Long Silken Cloak (4326, -1.36 DPS) [crafted]; Blackforge Cape (6424, -1.50 DPS, sim-verified) [dungeon] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 62.0 attack_power points (5.42 DPS) | yes | Red Mageweave Vest (10007, -0.34 DPS) [crafted]; Spirewind Fetter (9406, -0.81 DPS) [dungeon]; Icemail Jerkin (1981, -4.54 DPS, sim-verified) [world_drop] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 32.1 attack_power points (2.80 DPS) | yes | Glowing Magical Bracelets (13106, -0.02 DPS) [world_drop]; Warden's Wristbands (14600, -0.15 DPS) [world_drop]; Tough Scorpid Bracers (8205, -0.19 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (155.1 DPS) | yes | Engineering Gloves (1659, +0.00 DPS) [world]; Rustler Gloves (9704, +0.00 DPS) [quest]; Tough Scorpid Gloves (8204, -2.88 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 53.0 attack_power points (4.63 DPS) | yes | Defiler's Lizardhide Girdle (20173, -0.39 DPS) [rep]; Defiler's Mail Girdle (20198, -0.39 DPS) [rep]; Saltscale Girdle (273023, -0.52 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 64.8 attack_power points (5.66 DPS) | yes | Crimson Silk Pantaloons (7062, -1.54 DPS) [crafted]; Red Mageweave Pants (10009, -1.65 DPS, sim-verified) [crafted]; Totemic Leather Leggings (252458, -1.94 DPS) [crafted] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 48.7 attack_power points (4.25 DPS) | yes | Stalker's Mail Boots (252562, -0.83 DPS) [crafted]; Briar Tredders (10582, -1.06 DPS) [dungeon]; Skirmisher's Mail Boots (252564, -1.31 DPS, sim-verified) [crafted] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 28.0 attack_power points (2.45 DPS) | yes | Ogremind Ring (1993, -0.47 DPS) [world_drop]; Mindbender Loop (5009, -0.47 DPS) [world_drop]; Philanthropist's Ring (281635, -0.47 DPS) [quest] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993), Mindbender Loop (5009)) | Bloodscalp Witch Doctor [world] | 22.6 attack_power points (1.98 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Mindbender Loop (5009, +0.00 DPS) [world_drop]; Philanthropist's Ring (281635, -0.00 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-verified (155.1 DPS) | yes | Thornstone Sledgehammer (1722, -3.08 DPS) [world_drop]; Staff of Jordan (873, -3.14 DPS) [world_drop]; Mograine's Might (7723, -7.35 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; shoulder: Sheepshear Mantle; back: Darkspear Raider's Cloak; wrist: Windchaser Cuffs; hands: Gloves of Holy Might; waist: Deathmage Sash; feet: Prowler's Leather Shoes; finger1: Legionnaire's Band; finger2: Voodoo Band; trinket1: Ankh of Life; main_hand: Fiery War Axe

No-known-source sample (15 of 644, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7748 Forcestone Buckler

### Band 50 (orc, 3200300000000000-055030310051202051-0000000000000000)

Set DPS (verified): 216.5. Weights run: 6.3s. Verify run: 15.7s. 819 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.136 ± 0.045, crit=1.480 ± 0.061 per rating point (14 rating = 1%, 20.720 per %), hit=2.727 ± 0.124 per rating point (10 rating = 1%, 27.266 per %), melee_haste=10.040 ± 0.889, intellect=3.730 ± 0.057, spell_power=0.646 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | 91.6 attack_power points (7.54 DPS) | yes | Winged Helm (13112, -0.48 DPS) [world_drop]; Bad Mojo Mask (9470, -1.09 DPS) [dungeon]; Red Mageweave Headband (10033, -2.06 DPS, sim-verified) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 52.2 attack_power points (4.30 DPS) | yes | Glowing Eye of Mordresh (10769, -1.47 DPS) [dungeon]; Darkspear Warding Pendant (272073, -1.54 DPS) [vendor]; Archmagister's Faceted Pendant (273049, -1.88 DPS, sim-verified) [world] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 84.3 attack_power points (6.94 DPS) | yes | Rotgrip Mantle (17732, -0.72 DPS) [dungeon]; Rockshard Pauldrons (9411, -1.82 DPS) [dungeon]; Imperial Red Mantle (8250, -1.96 DPS) [world_drop] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Red Cloak (8248, -0.92 DPS) [world_drop]; Keeper's Cloak (14665, -1.23 DPS) [world_drop]; Deep Woodlands Cloak (19121, -1.64 DPS, sim-verified) [quest] |
| chest | Stone Guard's Mail Armor (220826) | Lady Palanseer [vendor] | 89.5 attack_power points (7.37 DPS) | yes | Stone Guard's Pulsing Breastplate (220844, -0.48 DPS) [vendor]; Relentless Chain (17777, -0.76 DPS) [quest]; Acumen Robes (17775, -1.94 DPS, sim-verified) [quest] |
| wrist | Forgotten Wraps (9433) | World drop [world_drop] | 48.2 attack_power points (3.97 DPS) | yes | Bracers of the Stone Princess (17714, -0.12 DPS) [dungeon]; Shizzle's Nozzle Wiper (11917, -0.28 DPS) [quest]; Prowler's Leather Bracers (252539, -0.29 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Skirmisher's Mail Gauntlets (252586, +0.00 DPS) [crafted]; Stonerender Gauntlets (17007, -0.84 DPS, sim-verified) [world_drop] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 74.7 attack_power points (6.15 DPS) | yes | Serenity Belt (13144, -0.93 DPS) [world_drop]; Sagebrush Girdle (17778, -1.07 DPS) [quest]; Imperial Red Sash (8253, -1.35 DPS, sim-verified) [world_drop] |
| legs | Stone Guard's Mail Legplates (220834) | Lady Palanseer [vendor] | 89.5 attack_power points (7.37 DPS) | yes | Stone Guard's Pulsing Legplates (220847, -0.48 DPS) [vendor]; Imperial Red Pants (8251, -1.36 DPS) [world_drop]; Ferine Leggings (6690, -1.54 DPS) [dungeon] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 71.7 attack_power points (5.90 DPS) | yes | Skycaller's Leather Boots (252471, -1.73 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -1.73 DPS) [crafted]; Rancor Boots (11865, -2.10 DPS, sim-verified) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 47.3 attack_power points (3.89 DPS) | yes | Ogremind Ring (1993, -1.74 DPS) [world_drop]; Voodoo Band (1996, -1.74 DPS) [world]; Mindbender Loop (5009, -1.74 DPS) [world_drop] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 30.2 attack_power points (2.49 DPS) | yes | Ogremind Ring (1993, -0.34 DPS) [world_drop]; Mindbender Loop (5009, -0.34 DPS) [world_drop]; Voodoo Band (1996, -2.00 DPS, sim-verified) [world] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -2.41 DPS) [world_drop]; Molten Heart of the Mountain (249470, -2.80 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.70 DPS) [world_drop]; Molten Heart of the Mountain (249470, -2.09 DPS, sim-verified) [crafted] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Soulkeeper (1607, -2.21 DPS) [world_drop]; Glowing Brightwood Staff (812, -3.83 DPS) [world_drop]; Bleakwood Hew (12769, -3.88 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Eye of Theradras; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Darkspear Raider's Cloak; chest: Stone Guard's Mail Armor; wrist: Forgotten Wraps; waist: Dawnspire Cord; legs: Stone Guard's Mail Legplates; feet: Greaves of Withering Despair; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Thorium Greatmace

No-known-source sample (15 of 819, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders

### Band 60 (orc, 3200330000000000-055030310051202051-0520000000000000)

Set DPS (verified): 315.1. Weights run: 5.8s. Verify run: 54.9s. 1628 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.226 ± 0.058, crit=1.963 ± 0.083 per rating point (14 rating = 1%, 27.476 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=12.029 ± 1.368, intellect=2.797 ± 0.056, spell_power=0.647 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 159.5 attack_power points (18.99 DPS) | yes | Warlord's Mail Headguard (231657, -3.73 DPS) [vendor]; Warlord's Mail Helm (231663, -7.29 DPS) [pvp]; Magister's Crown (16686, -15.13 DPS, sim-verified) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 39.2 attack_power points (4.66 DPS) | yes | Archmagister's Faceted Pendant (273049, -1.24 DPS) [world]; Darkspear Warding Pendant (272072, -1.33 DPS) [vendor]; Archlight Talisman (15856, -7.50 DPS, sim-verified) [quest] |
| shoulder | Darkspear Shoulderguards (272958) | Creeg Bothunk [vendor] | 85.9 attack_power points (10.22 DPS) | yes | Warlord's Mail Pauldrons (231654, +0.00 DPS) [vendor]; Warlord's Mail Spaulders (231659, -0.01 DPS) [pvp]; Darkspear Shoulderpads (272103, -0.31 DPS) [vendor] |
| back | Darkspear Raider's Cloak (272063) | Creeg Bothunk [vendor] | sim-verified (315.1 DPS) | yes | Deathguard's Cloak (20068, -0.55 DPS) [rep]; Cape of the Black Baron (13340, -0.76 DPS) [dungeon]; Cloak of Flames (3475, -6.95 DPS, sim-verified) [world_drop] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 115.5 attack_power points (13.76 DPS) | yes | Warlord's Mail Hauberk (231653, -0.92 DPS) [vendor]; Warlord's Mail Armor (16577, -0.99 DPS) [vendor]; Savage Gladiator Chain (11726, -3.38 DPS, sim-verified) [dungeon] |
| wrist | Windtalker's Wristguards (19582) | Warsong Outriders [rep] | 60.4 attack_power points (7.19 DPS) | yes | Tranquil Wristguards (279256, -0.05 DPS) [crafted]; Magiskull Cuffs (13107, -1.53 DPS) [world_drop]; Forest Stalker's Bracers (19587, -1.80 DPS) [rep] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-verified (315.1 DPS) | yes | General's Mail Gauntlets (231660, +0.00 DPS) [pvp]; General's Mail Gloves (231666, -0.26 DPS) [vendor]; Gloves of Holy Might (867, -4.42 DPS, sim-verified) [world_drop] |
| waist | Defiler's Lizardhide Girdle (20171) (or Defiler's Mail Girdle (20195)) | The Defilers [rep] | 75.0 attack_power points (8.93 DPS) | yes | General's Mail Waistband (16575, +0.00 DPS) [pvp]; Defiler's Mail Girdle (20195, +0.00 DPS) [rep]; Elunarian Belt (14465, -0.84 DPS) [world_drop] |
| legs | Outrider's Lizardhide Pants (22741) | Warsong Outriders [rep] | 124.9 attack_power points (14.87 DPS) | yes | General's Mail Leggings (231664, -1.45 DPS) [pvp]; General's Mail Legguards (231658, -1.60 DPS) [vendor]; Outrider's Mail Leggings (22676, -3.61 DPS, sim-verified) [rep] |
| feet | Timmy's Galoshes (13402) | Stratholme: Timmy the Cruel [dungeon] | 48.8 attack_power points (5.81 DPS) | yes | General's Mail Greaves (231656, +0.00 DPS) [vendor]; General's Mail Sabatons (231661, +0.00 DPS) [vendor]; Defiler's Lizardhide Boots (20167, -6.27 DPS, sim-verified) [rep] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 41.5 attack_power points (4.94 DPS) | yes | Deep Rooted Ring (19109, -0.68 DPS) [world]; Glowing Crystal Ring (18402, -0.94 DPS) [quest]; Maiden's Circle (13001, -1.22 DPS) [world_drop] |
| finger2 | Legionnaire's Band (19510) | Warsong Outriders [rep] | 37.5 attack_power points (4.46 DPS) | yes | Glowing Crystal Ring (18402, -0.47 DPS) [quest]; Maiden's Circle (13001, -0.75 DPS) [world_drop]; Deep Rooted Ring (19109, -3.57 DPS, sim-verified) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (315.1 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (315.1 DPS) | yes | Uther's Strength (11302, -1.20 DPS) [world_drop]; Molten Heart of the Mountain (249470, -7.87 DPS, sim-verified) [crafted] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (315.1 DPS) | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Forest Defender's Axe (279260, -3.31 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Black Dragonscale Helm; shoulder: Darkspear Shoulderguards; back: Darkspear Raider's Cloak; chest: Timbermaw Tunic; wrist: Windtalker's Wristguards; hands: Raider Handwraps; waist: Defiler's Lizardhide Girdle; legs: Outrider's Lizardhide Pants; feet: Timmy's Galoshes; finger1: Don Julio's Band; finger2: Legionnaire's Band; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: The Unstoppable Force

No-known-source sample (15 of 1628, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders

### Band 60, raid preset (orc, 0533001000000000-053030310051202051-0530000000000000)

Set DPS (verified): 830.5. Weights run: 5.4s. Verify run: 58.5s. 1628 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.332 ± 0.052, crit=2.158 ± 0.076 per rating point (14 rating = 1%, 30.210 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=14.786 ± 1.752, intellect=2.195 ± 0.019, spell_power=0.694 ± 0.001

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 141.6 attack_power points (32.30 DPS) | yes | Warlord's Mail Headguard (231657, -5.36 DPS) [vendor]; Champion's Mail Headguard (227155, -10.86 DPS) [pvp]; Coif of Elements (16667, -20.14 DPS, sim-verified) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 30.7 attack_power points (7.01 DPS) | yes | Archlight Talisman (15856, -0.10 DPS) [quest]; Skibi's Pendant (13089, -0.55 DPS) [world_drop]; Woven Ivy Necklace (19159, -1.27 DPS) [quest] |
| shoulder | Defiler's Mail Pauldrons (20203) | The Defilers [rep] | 74.8 attack_power points (17.07 DPS) | yes | Defiler's Lizardhide Shoulders (20175, +0.00 DPS) [rep]; Warlord's Mail Pauldrons (231654, +0.00 DPS) [vendor]; Warlord's Mail Spaulders (231659, +0.00 DPS) [pvp] |
| back | Deathguard's Cloak (20068) | The Defilers [rep] | sim-verified (830.5 DPS) | yes | Cape of the Black Baron (13340, -0.16 DPS) [dungeon]; Darkspear Raider's Cloak (272063, -1.26 DPS) [vendor]; Cloak of Flames (3475, -13.80 DPS, sim-verified) [world_drop] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 126.8 attack_power points (28.92 DPS) | yes | Warlord's Mail Hauberk (231653, -4.48 DPS) [vendor]; Warlord's Mail Armor (16577, -6.07 DPS) [vendor]; Savage Gladiator Chain (11726, -7.79 DPS, sim-verified) [dungeon] |
| wrist | Tranquil Wristguards (279256) | Leatherworking [crafted] | 57.2 attack_power points (13.04 DPS) | yes | Windtalker's Wristguards (19582, -0.36 DPS) [rep]; Forest Stalker's Bracers (19587, -1.75 DPS) [rep] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-verified (830.5 DPS) | yes | General's Mail Vices (231655, +0.00 DPS) [vendor]; General's Mail Gauntlets (231660, +0.00 DPS) [pvp]; Stonerender Gauntlets (17007, -17.95 DPS, sim-verified) [world_drop] |
| waist | Defiler's Lizardhide Girdle (20171) (or Defiler's Mail Girdle (20195)) | The Defilers [rep] | 67.5 attack_power points (15.40 DPS) | yes | Defiler's Mail Girdle (20195, +0.00 DPS) [rep]; General's Mail Waistband (16575, -0.38 DPS) [pvp]; Might of the Timbermaw (19044, -0.61 DPS) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (830.5 DPS) | yes | Outrider's Mail Leggings (22676, +0.00 DPS) [rep]; General's Mail Legguards (231658, +0.00 DPS) [vendor]; Outrider's Lizardhide Pants (22741, -11.11 DPS, sim-verified) [rep] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 48.7 attack_power points (11.10 DPS) | yes | Timmy's Galoshes (13402, +0.00 DPS) [dungeon]; General's Mail Greaves (231656, +0.00 DPS) [vendor]; General's Mail Sabatons (231661, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 44.2 attack_power points (10.08 DPS) | yes | Deep Rooted Ring (19109, -3.18 DPS) [world]; Band of the Penitent (13217, -3.19 DPS) [quest]; Don Rodrigo's Band (21563, -3.19 DPS) [rep] |
| finger2 | Legionnaire's Band (19510) | Warsong Outriders [rep] | 41.0 attack_power points (9.36 DPS) | yes | Deep Rooted Ring (19109, -2.46 DPS) [world]; Band of the Penitent (13217, -2.47 DPS) [quest]; Don Rodrigo's Band (21563, -2.47 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (830.5 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (830.5 DPS) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted]; Uther's Strength (11302, -2.24 DPS) [world_drop] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (830.5 DPS) | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Forest Defender's Axe (279260, -9.08 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Black Dragonscale Helm; shoulder: Defiler's Mail Pauldrons; back: Deathguard's Cloak; chest: Timbermaw Tunic; wrist: Tranquil Wristguards; hands: Devilsaur Gauntlets; waist: Defiler's Lizardhide Girdle; legs: Devilsaur Leggings; feet: Scalegut Treaders; finger1: Don Julio's Band; finger2: Legionnaire's Band; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: The Unstoppable Force

No-known-source sample (15 of 1628, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders

