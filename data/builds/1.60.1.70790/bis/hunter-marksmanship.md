# Leveling BiS: Marksmanship

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-0051500000000000-000000000000000000)

Set DPS (verified): 67.5. Weights run: 7.2s. Verify run: 7.9s. 258 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.163 ± 0.012, crit=0.671 ± 0.013 per rating point (14 rating = 1%, 9.397 per %), hit=1.167 ± 0.039 per rating point (10 rating = 1%, 11.667 per %), melee_haste=2.959 ± 0.588, intellect=1.439 ± 0.032, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 25.3 ranged_attack_power points (1.73 DPS) | yes | Trapper's Leather Hood (252505, -1.13 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.24 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 19.0 ranged_attack_power points (1.29 DPS) | yes | Erudite's Amulet (277204, -0.84 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 15.8 ranged_attack_power points (1.08 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.69 DPS) [crafted]; Magician's Mantle (12998, -0.84 DPS, sim-verified) [world_drop] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Brotherhood (5193, -0.21 DPS, sim-verified) [dungeon]; Sentry Cloak (2059, -0.43 DPS) [world_drop]; Bristlebark Cape (14571, -0.43 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 34.8 ranged_attack_power points (2.37 DPS) | yes | Trapper's Leather Armor (252491, -0.47 DPS, sim-verified) [crafted]; Brawler's Leather Armor (252490, -0.86 DPS) [crafted]; Mystic's Wrap (14369, -1.04 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 15.8 ranged_attack_power points (1.08 DPS) | yes | Wolf Bracers (4794, -0.21 DPS, sim-verified) [vendor]; Bristlebark Bindings (14569, -0.43 DPS) [world_drop]; Bravo's Armbands (270015, -0.43 DPS) [quest] |
| hands | Gloves of the Fang (10413) (or Serpent Gloves (5970)) | Wailing Caverns: Druid of the Fang [dungeon] | 19.0 ranged_attack_power points (1.29 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Trapper's Leather Gloves (252495, -0.14 DPS) [crafted]; Riding Gloves (1304, -0.35 DPS) [quest] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.23 DPS) | yes | Deviate Scale Belt (6468, -0.15 DPS) [crafted]; Trapper's Leather Belt (252429, -0.22 DPS, sim-verified) [crafted]; Dark Leather Belt (4249, -0.37 DPS) [crafted] |
| legs | Leggings of the Fang (10410) | Wailing Caverns: Lord Cobrahn [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.22 DPS) [world]; Trapper's Leather Pants (252501, -0.46 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 25.3 ranged_attack_power points (1.73 DPS) | yes | Trapper's Leather Boots (252440, -0.29 DPS, sim-verified) [crafted]; Blackened Defias Boots (10402, -0.43 DPS) [dungeon]; Footpads of the Fang (10411, -0.43 DPS) [dungeon] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 15.0 ranged_attack_power points (1.02 DPS) | yes | The 1 Ring (8350, -0.71 DPS) [world]; Volcanic Rock Ring (12053, -0.73 DPS) [world_drop]; Minor Channeling Ring (1449, -0.82 DPS) [quest] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 12.7 ranged_attack_power points (0.86 DPS) | yes | The 1 Ring (8350, -0.54 DPS, sim-verified) [world]; Volcanic Rock Ring (12053, -0.57 DPS) [world_drop]; Minor Channeling Ring (1449, -0.67 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Impaling Harpoon (5200) | The Deadmines: Captain Greenskin [dungeon] | 35.7 ranged_attack_power points (2.43 DPS) | yes | Bronze Dory (250603, -0.33 DPS, sim-verified) [crafted]; Lupine Axe (1220, -1.14 DPS) [world]; Piercing Axe (6094, -1.14 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.5 ranged_attack_power points (12.25 DPS) | yes | Lil Timmy's Peashooter (13136, -0.91 DPS) [world_drop]; Venomstrike (6469, -1.60 DPS, sim-verified) [dungeon]; Lovingly Crafted Boomstick (4372, -2.96 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Lavishly Jeweled Ring; finger2: Protector's Band; main_hand: Impaling Harpoon; ranged: Ranger Bow

No-known-source sample (15 of 258, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (dwarf, 0000000000000000-0051550001400000-000000000000000000)

Set DPS (verified): 108.8. Weights run: 7.5s. Verify run: 9.5s. 447 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.331 ± 0.018, crit=1.289 ± 0.025 per rating point (14 rating = 1%, 18.052 per %), hit=1.803 ± 0.062 per rating point (10 rating = 1%, 18.030 per %), melee_haste=not significant (-1.650 ± 0.831), intellect=2.764 ± 0.043, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 44.2 ranged_attack_power points (3.07 DPS) | yes | Resilient Cap (14401, -0.70 DPS, sim-verified) [world_drop]; Brawler's Leather Helm (252512, -0.76 DPS) [crafted]; Nightsky Cowl (4039, -0.77 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 28.6 ranged_attack_power points (1.98 DPS) | yes | Kaleidoscope Chain (13084, -0.29 DPS) [world_drop]; Sentinel's Medallion (19541, -0.33 DPS, sim-verified) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 36.6 ranged_attack_power points (2.54 DPS) | yes | Mantle of Thieves (2264, -0.38 DPS, sim-verified) [dungeon]; Death Speaker Mantle (6685, -0.43 DPS) [dungeon]; Feline Mantle (3748, -0.55 DPS) [dungeon] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sergeant Major's Cape (16315, -0.14 DPS, sim-verified) [pvp]; Hawkeye's Cloak (14593, -0.23 DPS) [world_drop]; Expeditionary's Cape (271923, -0.31 DPS) [rep] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 71.4 ranged_attack_power points (4.95 DPS) | yes | Green Leather Armor (4255, -1.42 DPS, sim-verified) [crafted]; Dusky Leather Armor (7374, -1.72 DPS) [crafted]; Resilient Tunic (14398, -1.92 DPS) [world_drop] |
| wrist | Glowing Magical Bracelets (13106) | World drop [world_drop] | 22.1 ranged_attack_power points (1.53 DPS) | yes | Insignia Bracers (6410, -0.15 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.15 DPS) [world_drop]; Forest Leather Bracers (3202, -0.38 DPS) [world_drop] |
| hands | Hotshot Pilot's Gloves (9491) | Gnomeregan [dungeon] | 35.4 ranged_attack_power points (2.46 DPS) | yes | Cartilage Shapers (273054, -0.38 DPS) [world]; Pilferer's Gloves (7358, -0.61 DPS) [crafted]; Resilient Handgrips (14403, -0.61 DPS) [world_drop] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 46.6 ranged_attack_power points (3.23 DPS) | yes | Highlander's Lizardhide Girdle (20105, -0.93 DPS) [rep]; Highlander's Mail Girdle (20120, -0.93 DPS) [vendor]; Green Leather Belt (4257, -0.93 DPS, sim-verified) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 59.2 ranged_attack_power points (4.10 DPS) | yes | Trapper's Leather Legguards (252517, -0.68 DPS, sim-verified) [crafted]; Petrolspill Leggings (9509, -0.87 DPS) [dungeon]; Troll's Bane Leggings (13114, -0.87 DPS) [world_drop] |
| feet | Trapper's Leather Boots (252440) | Leatherworking [crafted] | 27.7 ranged_attack_power points (1.92 DPS) | yes | Feet of the Lynx (1121, -0.07 DPS) [world_drop]; Insignia Boots (4055, -0.07 DPS) [world_drop]; Acidic Walkers (9454, -0.39 DPS) [dungeon] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 23.2 ranged_attack_power points (1.61 DPS) | yes | Ring of Precision (1491, -0.23 DPS) [dungeon]; Protector's Band (19517, -0.23 DPS) [rep]; Monkey Ring (6748, -0.46 DPS) [quest] |
| finger2 | Seal of Wrynn (2933) | An Audience with the King [quest] | 21.0 ranged_attack_power points (1.46 DPS) | yes | Ring of Precision (1491, -0.07 DPS) [dungeon]; Protector's Band (19517, -0.07 DPS) [rep]; Monkey Ring (6748, -0.30 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Felbough Arm (273037) | Fel Ancient [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Impaling Harpoon (5200, -0.52 DPS) [dungeon]; Kam's Walking Stick (2280, -0.68 DPS) [dungeon]; Claw of the Shadowmancer (2912, -1.67 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Crystalpine Stinger (13037) | World drop [world_drop] | sim-verified (108.8 DPS) | yes | Ironweaver (13137, -0.32 DPS) [world_drop]; Harpyclaw Short Bow (13019, -0.71 DPS) [world_drop]; Golemsight Long Gun (273029, -2.07 DPS, sim-verified) [dungeon] |

**New at 30:** head: Enduring Cap; neck: Archmagister's Faceted Pendant; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Garb of Florid Feathers; wrist: Glowing Magical Bracelets; hands: Hotshot Pilot's Gloves; waist: Stalker's Leather Belt; legs: Ferine Leggings; feet: Trapper's Leather Boots; finger2: Seal of Wrynn; main_hand: Felbough Arm; ranged: Crystalpine Stinger

No-known-source sample (15 of 447, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade

### Band 40 (dwarf, 1000000000000000-0051550001503050-000000000000000000)

Set DPS (verified): 170.8. Weights run: 7.8s. Verify run: 20.8s. 681 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.368 ± 0.023, crit=1.870 ± 0.038 per rating point (14 rating = 1%, 26.176 per %), hit=2.314 ± 0.095 per rating point (10 rating = 1%, 23.139 per %), melee_haste=not significant (5.772 ± 1.473), intellect=2.916 ± 0.063, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 52.5 ranged_attack_power points (4.14 DPS) | yes | Papal Fez (9431, -0.23 DPS) [dungeon]; Enduring Cap (3020, -0.46 DPS) [world_drop]; Thinking Cap (2624, -0.55 DPS, sim-verified) [world] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 37.0 ranged_attack_power points (2.92 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.90 DPS, sim-verified) [world]; Kaleidoscope Chain (13084, -0.94 DPS) [world_drop]; Glowing Eye of Mordresh (10769, -1.08 DPS) [dungeon] |
| shoulder | Windchaser Amice (14432) | World drop [world_drop] | 48.0 ranged_attack_power points (3.78 DPS) | yes | Sheepshear Mantle (13115, -0.11 DPS) [world_drop]; Inquisitor's Shawl (19507, -0.80 DPS) [dungeon]; Nightscape Shoulders (8192, -0.86 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-verified (170.8 DPS) | yes | Imperial Cloak (6432, -0.85 DPS) [dungeon]; Parachute Cloak (10518, -0.85 DPS) [crafted]; Expeditionary's Cape (271922, -1.09 DPS, sim-verified) [rep] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 104.0 ranged_attack_power points (8.19 DPS) | yes | Tough Scorpid Breastplate (8203, -2.60 DPS) [crafted]; Garb of Florid Feathers (273026, -3.19 DPS, sim-verified) [dungeon]; Wolffear Harness (13110, -3.68 DPS) [world_drop] |
| wrist | Tough Scorpid Bracers (8205) | Leatherworking [crafted] | 41.1 ranged_attack_power points (3.24 DPS) | yes | Green Leather Bracers (4259, -0.59 DPS, sim-verified) [crafted]; Warden's Wristbands (14600, -0.60 DPS) [world_drop]; Windchaser Cuffs (14429, -0.64 DPS) [world_drop] |
| hands | Tough Scorpid Gloves (8204) | Leatherworking [crafted] | sim-verified (170.8 DPS) | yes | Stalker's Leather Gloves (252526, -0.59 DPS, sim-verified) [crafted]; Aurora Gloves (4042, -0.76 DPS) [world_drop]; Dragonscale Gauntlets (8347, -1.07 DPS) [crafted] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 47.8 ranged_attack_power points (3.77 DPS) | yes | Deathmage Sash (10771, -0.32 DPS) [dungeon]; Highlander's Mail Girdle (20119, -0.32 DPS) [vendor]; Highlander's Lizardhide Girdle (20104, -1.27 DPS, sim-verified) [rep] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 70.7 ranged_attack_power points (5.57 DPS) | yes | Triprunner Dungarees (9624, -0.80 DPS) [quest]; Ferine Leggings (6690, -1.45 DPS, sim-verified) [dungeon]; Trapper's Leather Legguards (252517, -1.58 DPS) [crafted] |
| feet | Stalker's Mail Boots (252562) | Leatherworking [crafted] | 57.5 ranged_attack_power points (4.53 DPS) | yes | Swampwalker Boots (2276, -1.08 DPS) [world_drop]; Briar Tredders (10582, -1.13 DPS) [dungeon]; Worn Running Boots (9398, -1.62 DPS, sim-verified) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 33.7 ranged_attack_power points (2.65 DPS) | yes | Protector's Band (19515, -0.53 DPS) [rep]; Disengagement Ring (276202, -0.53 DPS) [vendor]; Lavishly Jeweled Ring (1156, -0.74 DPS) [dungeon] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 30.3 ranged_attack_power points (2.39 DPS) | yes | Protector's Band (19515, -0.27 DPS) [rep]; Disengagement Ring (276202, -0.27 DPS) [vendor]; Lavishly Jeweled Ring (1156, -0.48 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (170.8 DPS) | yes | Felbough Arm (273037, -2.85 DPS) [world]; Steel Spear (250605, -3.09 DPS) [crafted]; Claw of the Shadowmancer (2912, -5.82 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | sim-verified (170.8 DPS) | yes | Skystriker Bow (13020, -2.25 DPS) [world_drop]; Swiftwind (13038, -2.28 DPS) [world_drop]; The Silencer (13138, -3.78 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Sentinel's Medallion; shoulder: Windchaser Amice; back: Sergeant Major's Cape; chest: Icemail Jerkin; wrist: Tough Scorpid Bracers; hands: Tough Scorpid Gloves; legs: Basilisk Hide Pants; feet: Stalker's Mail Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Bow of Searing Arrows

No-known-source sample (15 of 681, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 5000 Coral Band

### Band 50 (dwarf, 5320000100000000-0051550001503050-000000000000000000)

Set DPS (verified): 236.9. Weights run: 8.1s. Verify run: 30.2s. 854 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.234 ± 0.018, crit=2.042 ± 0.032 per rating point (14 rating = 1%, 28.594 per %), hit=2.804 ± 0.083 per rating point (10 rating = 1%, 28.045 per %), melee_haste=14.385 ± 1.086, intellect=2.776 ± 0.046, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Chain Helmet (220822, +0.00 DPS) [vendor]; Tough Scorpid Helm (8208, -0.60 DPS, sim-verified) [crafted]; Khan's Helmet (14785, -1.31 DPS) [world_drop] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 42.0 ranged_attack_power points (3.86 DPS) | yes | Horizon Choker (13085, -0.29 DPS) [world_drop]; Sentinel's Medallion (19539, -0.30 DPS) [rep] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Chain Epaulets (220825, +0.00 DPS) [vendor]; Stalker's Mail Shoulders (252566, -0.05 DPS) [crafted]; Tough Scorpid Shoulders (8207, -0.53 DPS, sim-verified) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackflame Cape (13109, -0.89 DPS) [world_drop]; Sergeant Major's Cape (16336, -1.14 DPS) [pvp]; Expeditionary's Cape (271921, -1.16 DPS, sim-verified) [rep] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 105.7 ranged_attack_power points (9.70 DPS) | yes | Knight's Chain Armor (220828, -0.07 DPS) [vendor]; Blazewind Breastplate (11193, -2.87 DPS) [quest]; Garb of Florid Feathers (273026, -3.21 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 48.5 ranged_attack_power points (4.45 DPS) | yes | Stalker's Mail Bracers (252570, -0.25 DPS) [crafted]; Forgotten Wraps (9433, -0.50 DPS) [world_drop]; Bracers of the Stone Princess (17714, -0.61 DPS) [dungeon] |
| hands | Stalker's Mail Gauntlets (252584) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tough Scorpid Gloves (8204, -1.19 DPS) [crafted]; Stalker's Leather Gloves (252526, -1.70 DPS) [crafted]; Hibernal Gloves (8110, -2.09 DPS, sim-verified) [world_drop] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 81.8 ranged_attack_power points (7.51 DPS) | yes | Stalker's Mail Belt (252588, -1.29 DPS, sim-verified) [crafted]; Highlander's Lizardhide Girdle (20103, -2.34 DPS) [rep]; Highlander's Mail Girdle (20118, -2.34 DPS) [vendor] |
| legs | Knight's Chain Legplates (220832) | Captain Dirgehammer [vendor] | 101.6 ranged_attack_power points (9.33 DPS) | yes | Imperial Red Pants (8251, -1.38 DPS, sim-verified) [world_drop]; Tough Scorpid Leggings (8206, -1.73 DPS) [crafted]; Searingscale Leggings (11749, -2.16 DPS) [dungeon] |
| feet | Fleetfoot Greaves (11627) | Blackrock Depths: Houndmaster Grebmar [dungeon] | 78.1 ranged_attack_power points (7.17 DPS) | yes | Rancor Boots (11865, -0.91 DPS, sim-verified) [quest]; Stalker's Mail Sabatons (252576, -1.02 DPS) [crafted]; Albino Crocscale Boots (17728, -1.23 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 48.0 ranged_attack_power points (4.41 DPS) | yes | Ring of the Underwood (2951, -1.44 DPS) [world_drop]; Falcon's Hook (7552, -1.74 DPS) [dungeon]; Protector's Band (19516, -1.74 DPS) [rep] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 45.3 ranged_attack_power points (4.16 DPS) | yes | Ring of the Underwood (2951, -1.21 DPS, sim-verified) [world_drop]; Falcon's Hook (7552, -1.48 DPS) [dungeon]; Protector's Band (19516, -1.48 DPS) [rep] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -0.83 DPS, sim-verified) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warmonger (13052, -1.24 DPS) [world_drop]; Glowing Brightwood Staff (812, -1.57 DPS) [world_drop]; Claw of the Shadowmancer (2912, -7.65 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gryphonwing Long Bow (13022, -2.84 DPS, sim-verified) [dungeon]; Megashot Rifle (17717, -3.18 DPS) [dungeon]; Dark Iron Rifle (16004, -3.73 DPS) [crafted] |

**New at 50:** head: Eye of Theradras; neck: Skibi's Pendant; shoulder: Ironfeather Shoulders; back: Dark Phantom Cape; wrist: Deepfury Bracers; hands: Stalker's Mail Gauntlets; waist: Sagebrush Girdle; legs: Knight's Chain Legplates; feet: Fleetfoot Greaves; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Thunderbrew's Boot Flask; trinket2: Molten Heart of the Mountain; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 854, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band

### Band 60 (dwarf, 5320000501000000-0051550001503050-500000000000000000)

Set DPS (verified): 352.7. Weights run: 7.8s. Verify run: 34.5s. 1594 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.208 ± 0.022, crit=2.541 ± 0.046 per rating point (14 rating = 1%, 35.568 per %), hit=3.378 ± 0.147 per rating point (10 rating = 1%, 33.782 per %), melee_haste=12.249 ± 1.971, intellect=3.059 ± 0.084, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 219.2 ranged_attack_power points (22.64 DPS) | yes | Field Marshal's Chain Helm (231580, -3.91 DPS) [pvp]; Field Marshal's Chain Greathelm (231562, -5.82 DPS) [vendor]; Outlaw's Collar (279253, -12.39 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (19538) | Silverwing Sentinels [rep] | 48.1 ranged_attack_power points (4.97 DPS) | yes | Skibi's Pendant (13089, -0.66 DPS) [world_drop]; Opaline Medallion (16623, -0.86 DPS) [quest]; Horizon Choker (13085, -1.88 DPS, sim-verified) [world_drop] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 122.2 ranged_attack_power points (12.62 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor]; Theramore Epaulets (271956, +0.00 DPS) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-verified (352.7 DPS) | yes | Cloak of the Honor Guard (20073, -1.87 DPS) [rep]; Expeditionary's Cape (271907, -1.98 DPS) [rep]; Sergeant Major's Cape (16337, -2.88 DPS, sim-verified) [pvp] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 188.0 ranged_attack_power points (19.42 DPS) | yes | Field Marshal's Chain Breastplate (16466, -0.69 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -0.69 DPS) [vendor]; Savage Gladiator Chain (11726, -9.11 DPS, sim-verified) [dungeon] |
| wrist | Windtalker's Wristguards (19582) | Silverwing Sentinels [rep] | 62.5 ranged_attack_power points (6.45 DPS) | yes | Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Spitfire Bracers (20481, -0.63 DPS) [crafted]; Forest Stalker's Bracers (19587, -1.86 DPS, sim-verified) [rep] |
| hands | Mourning Handguards (279257) | Leatherworking [crafted] | sim-verified (352.7 DPS) | yes | Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Marshal's Chain Grips (231560, -0.86 DPS) [pvp]; Theramore Gloves (271948, -3.66 DPS, sim-verified) [rep] |
| waist | Highlander's Mail Girdle (20044) (or Highlander's Lizardhide Girdle (20046)) | Samuel Hawke [vendor] | 87.6 ranged_attack_power points (9.05 DPS) | yes | Marshal's Chain Girdle (16464, +0.00 DPS) [pvp]; Highlander's Lizardhide Girdle (20046, +0.00 DPS) [rep]; Sagebrush Girdle (17778, -0.28 DPS) [quest] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 217.2 ranged_attack_power points (22.44 DPS) | yes | Marshal's Chain Legguards (231577, -3.70 DPS) [pvp]; Marshal's Chain Legplates (231558, -5.62 DPS) [vendor]; Sentinel's Leather Pants (237818, -6.14 DPS) [vendor] |
| feet | Timmy's Galoshes (13402) | Stratholme: Timmy the Cruel [dungeon] | 85.1 ranged_attack_power points (8.79 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor]; Fleetfoot Greaves (11627, -1.94 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 83.4 ranged_attack_power points (8.61 DPS) | yes | Masons Fraternity Ring (9533, -3.97 DPS) [quest]; Glowing Crystal Ring (18402, -4.82 DPS) [quest]; Band of the Penitent (13217, -4.94 DPS) [quest] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 53.8 ranged_attack_power points (5.56 DPS) | yes | Masons Fraternity Ring (9533, -0.92 DPS) [quest]; Glowing Crystal Ring (18402, -1.76 DPS) [quest]; Band of the Penitent (13217, -1.88 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (352.7 DPS) | yes | - |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (352.7 DPS) | yes | Molten Heart of the Mountain (249470, -1.71 DPS, sim-verified) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (352.7 DPS) | yes | Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Electrified Dagger (19100, -7.76 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (352.7 DPS) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -12.60 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; neck: Sentinel's Medallion; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Windtalker's Wristguards; hands: Mourning Handguards; waist: Highlander's Mail Girdle; legs: Sentinel's Chain Leggings; feet: Timmy's Galoshes; finger1: Don Julio's Band; finger2: Blackstone Ring; trinket1: Frozen Heart of the Mountain; trinket2: Thunderbrew's Boot Flask; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1594, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band

### Band 60, raid preset (dwarf, 5000000000000000-0053550011503050-500050030000000000)

Set DPS (verified): 786.4. Weights run: 4.8s. Verify run: 51.3s. 1594 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.264 ± 0.025, crit=2.668 ± 0.052 per rating point (14 rating = 1%, 37.352 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=20.259 ± 2.014, intellect=1.834 ± 0.036, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 175.6 ranged_attack_power points (36.86 DPS) | yes | Field Marshal's Chain Helm (231580, -1.10 DPS) [pvp]; Field Marshal's Chain Greathelm (231562, -5.19 DPS) [vendor]; Outlaw's Collar (279253, -12.23 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (19538) | Silverwing Sentinels [rep] | 49.0 ranged_attack_power points (10.28 DPS) | yes | Skibi's Pendant (13089, -1.37 DPS) [world_drop]; Sunprism Pendant (20649, -2.06 DPS) [quest]; Horizon Choker (13085, -4.89 DPS) [world_drop] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 125.5 ranged_attack_power points (26.34 DPS) | yes | Theramore Epaulets (271956, +0.00 DPS) [rep]; Field Marshal's Chain Spaulders (16468, -4.68 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, -4.68 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-verified (786.4 DPS) | yes | Cloak of the Honor Guard (20073, -3.58 DPS, sim-verified) [rep]; Dark Phantom Cape (13122, -4.20 DPS) [world_drop]; Sergeant Major's Cape (16337, -5.92 DPS) [pvp] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 158.1 ranged_attack_power points (33.19 DPS) | yes | Field Marshal's Chain Breastplate (16466, +0.00 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, +0.00 DPS) [vendor]; Savage Gladiator Chain (11726, -18.87 DPS, sim-verified) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 62.0 ranged_attack_power points (13.02 DPS) | yes | Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Knight-Captain's Chain Armguards (16402, -1.37 DPS) [pvp]; Windtalker's Wristguards (19582, -1.96 DPS) [rep] |
| hands | Mourning Handguards (279257) | Leatherworking [crafted] | sim-verified (786.4 DPS) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Theramore Gloves (271948, -1.28 DPS) [rep] |
| waist | Highlander's Chain Girdle (20043) (or Highlander's Leather Girdle (20045)) | The League of Arathor [rep] | 71.4 ranged_attack_power points (14.98 DPS) | yes | Marshal's Chain Girdle (16464, +0.00 DPS) [pvp]; Highlander's Leather Girdle (20045, +0.00 DPS) [rep]; Sagebrush Girdle (17778, -0.08 DPS) [quest] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 188.9 ranged_attack_power points (39.66 DPS) | yes | Marshal's Chain Legguards (231577, -3.91 DPS) [pvp]; Marshal's Chain Legplates (231558, -8.00 DPS) [vendor]; Sentinel's Leather Pants (237818, -12.09 DPS, sim-verified) [vendor] |
| feet | Beaststalker's Boots (16675) | Stratholme: Nerub'enkan [dungeon] | 74.0 ranged_attack_power points (15.54 DPS) | yes | Timmy's Galoshes (13402, +0.00 DPS) [dungeon]; Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 51.4 ranged_attack_power points (10.78 DPS) | yes | Band of the Penitent (13217, -2.94 DPS) [quest]; Don Rodrigo's Band (21563, -2.94 DPS) [rep]; Protector's Band (19514, -3.24 DPS) [rep] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 45.7 ranged_attack_power points (9.59 DPS) | yes | Band of the Penitent (13217, -1.75 DPS) [quest]; Don Rodrigo's Band (21563, -1.75 DPS) [rep]; Protector's Band (19514, -2.06 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (786.4 DPS) | yes | Molten Heart of the Mountain (249470, -3.44 DPS, sim-verified) [crafted] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (786.4 DPS) | yes | Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Electrified Dagger (19100, -12.88 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (786.4 DPS) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Eaglehorn Long Bow (13023, -5.64 DPS, sim-verified) [world_drop] |

**New at 60:** head: Black Dragonscale Helm; neck: Sentinel's Medallion; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Mourning Handguards; waist: Highlander's Chain Girdle; legs: Sentinel's Chain Leggings; feet: Beaststalker's Boots; finger1: Don Julio's Band; trinket1: Frozen Heart of the Mountain; trinket2: Thunderbrew's Boot Flask; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1594, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band

## Horde

### Band 20 (troll, 0000000000000000-0051500000000000-000000000000000000)

Set DPS (verified): 68.4. Weights run: 7.2s. Verify run: 7.8s. 254 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.163 ± 0.012, crit=0.671 ± 0.013 per rating point (14 rating = 1%, 9.397 per %), hit=1.167 ± 0.039 per rating point (10 rating = 1%, 11.667 per %), melee_haste=2.959 ± 0.588, intellect=1.439 ± 0.032, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 25.3 ranged_attack_power points (1.73 DPS) | yes | Trapper's Leather Hood (252505, -0.95 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.24 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 19.0 ranged_attack_power points (1.29 DPS) | yes | Erudite's Amulet (277204, -0.92 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 15.8 ranged_attack_power points (1.08 DPS) | yes | Magician's Mantle (12998, -0.63 DPS, sim-verified) [world_drop]; Reinforced Woolen Shoulders (4315, -0.69 DPS) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Brotherhood (5193, -0.22 DPS, sim-verified) [dungeon]; Sentry Cloak (2059, -0.43 DPS) [world_drop]; Bristlebark Cape (14571, -0.43 DPS) [world_drop] |
| chest | Trapper's Leather Armor (252491) | Leatherworking [crafted] | 26.5 ranged_attack_power points (1.81 DPS) | yes | Brawler's Leather Armor (252490, -0.32 DPS, sim-verified) [crafted]; Mystic's Wrap (14369, -0.47 DPS) [world_drop]; Mystic's Robe (14371, -0.47 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 15.8 ranged_attack_power points (1.08 DPS) | yes | Wolf Bracers (4794, -0.22 DPS, sim-verified) [vendor]; Bristlebark Bindings (14569, -0.43 DPS) [world_drop]; Spare Part Bindings (279875, -0.43 DPS) [quest] |
| hands | Gloves of the Fang (10413) (or Serpent Gloves (5970)) | Wailing Caverns: Druid of the Fang [dungeon] | 19.0 ranged_attack_power points (1.29 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Trapper's Leather Gloves (252495, -0.14 DPS) [crafted]; Forest Leather Gloves (3058, -0.43 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.23 DPS) | yes | Deviate Scale Belt (6468, -0.15 DPS) [crafted]; Trapper's Leather Belt (252429, -0.33 DPS, sim-verified) [crafted]; Dark Leather Belt (4249, -0.37 DPS) [crafted] |
| legs | Leggings of the Fang (10410) | Wailing Caverns: Lord Cobrahn [dungeon] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.22 DPS) [world]; Trapper's Leather Pants (252501, -0.36 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 25.3 ranged_attack_power points (1.73 DPS) | yes | Trapper's Leather Boots (252440, -0.36 DPS, sim-verified) [crafted]; Blackened Defias Boots (10402, -0.43 DPS) [dungeon]; Footpads of the Fang (10411, -0.43 DPS) [dungeon] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 15.0 ranged_attack_power points (1.02 DPS) | yes | Bounty Hunter's Ring (5351, -0.37 DPS) [quest]; Loop of Sacrifice (281673, -0.53 DPS) [quest]; The 1 Ring (8350, -0.71 DPS) [world] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 12.7 ranged_attack_power points (0.86 DPS) | yes | Bounty Hunter's Ring (5351, -0.22 DPS, sim-verified) [quest]; Loop of Sacrifice (281673, -0.37 DPS) [quest]; The 1 Ring (8350, -0.55 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Impaling Harpoon (5200) | The Deadmines: Captain Greenskin [dungeon] | 35.7 ranged_attack_power points (2.43 DPS) | yes | Crescent Staff (6505, -0.26 DPS, sim-verified) [quest]; Bronze Dory (250603, -0.30 DPS) [crafted]; Staff of Orgrimmar (15444, -0.55 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.5 ranged_attack_power points (12.25 DPS) | yes | Lil Timmy's Peashooter (13136, -0.91 DPS) [world_drop]; Venomstrike (6469, -1.64 DPS, sim-verified) [dungeon]; Lovingly Crafted Boomstick (4372, -2.96 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Trapper's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Lavishly Jeweled Ring; finger2: Legionnaire's Band; main_hand: Impaling Harpoon; ranged: Ranger Bow

No-known-source sample (15 of 254, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18856 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20437 Outrider's Bow; 20441 Scout's Blade

### Band 30 (troll, 0000000000000000-0051550001400000-000000000000000000)

Set DPS (verified): 110.4. Weights run: 7.5s. Verify run: 8.8s. 446 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.331 ± 0.018, crit=1.289 ± 0.025 per rating point (14 rating = 1%, 18.052 per %), hit=1.803 ± 0.062 per rating point (10 rating = 1%, 18.030 per %), melee_haste=not significant (-1.650 ± 0.831), intellect=2.764 ± 0.043, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 44.2 ranged_attack_power points (3.07 DPS) | yes | Resilient Cap (14401, -0.50 DPS) [world_drop]; Brawler's Leather Helm (252512, -0.76 DPS) [crafted]; Nightsky Cowl (4039, -0.77 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 28.6 ranged_attack_power points (1.98 DPS) | yes | Scout's Medallion (19537, -0.13 DPS) [rep]; Kaleidoscope Chain (13084, -0.29 DPS) [world_drop] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (110.4 DPS) | yes | Mantle of Thieves (2264, +0.00 DPS) [dungeon]; Forest Tracker Epaulets (2278, +0.00 DPS) [world_drop]; Feline Mantle (3748, -0.11 DPS) [dungeon] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hawkeye's Cloak (14593, -0.23 DPS) [world_drop]; Darkspear Raider's Cloak (272078, -0.31 DPS) [vendor]; Swiftrunner Cape (6745, -0.46 DPS) [quest] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 71.4 ranged_attack_power points (4.95 DPS) | yes | Green Leather Armor (4255, -1.57 DPS) [crafted]; Dusky Leather Armor (7374, -1.72 DPS) [crafted]; Resilient Tunic (14398, -1.92 DPS) [world_drop] |
| wrist | Glowing Magical Bracelets (13106) | World drop [world_drop] | 22.1 ranged_attack_power points (1.53 DPS) | yes | Insignia Bracers (6410, -0.15 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.15 DPS) [world_drop]; Loamflake Bracers (15462, -0.38 DPS) [quest] |
| hands | Hotshot Pilot's Gloves (9491) | Gnomeregan [dungeon] | 35.4 ranged_attack_power points (2.46 DPS) | yes | Cartilage Shapers (273054, -0.38 DPS) [world]; Pilferer's Gloves (7358, -0.61 DPS) [crafted]; Resilient Handgrips (14403, -0.61 DPS) [world_drop] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 46.6 ranged_attack_power points (3.23 DPS) | yes | Green Leather Belt (4257, -0.69 DPS) [crafted]; Defiler's Lizardhide Girdle (20172, -0.93 DPS) [rep]; Defiler's Mail Girdle (20197, -0.93 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 59.2 ranged_attack_power points (4.10 DPS) | yes | Trapper's Leather Legguards (252517, -0.68 DPS) [crafted]; Petrolspill Leggings (9509, -0.87 DPS) [dungeon]; Troll's Bane Leggings (13114, -0.87 DPS) [world_drop] |
| feet | Trapper's Leather Boots (252440) | Leatherworking [crafted] | 27.7 ranged_attack_power points (1.92 DPS) | yes | Feet of the Lynx (1121, -0.07 DPS) [world_drop]; Insignia Boots (4055, -0.07 DPS) [world_drop]; Warsong Boots (16977, -0.07 DPS) [quest] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 23.2 ranged_attack_power points (1.61 DPS) | yes | Legionnaire's Band (19513, -0.23 DPS) [rep]; Monkey Ring (6748, -0.46 DPS) [quest]; Azora's Will (4999, -0.65 DPS) [world_drop] |
| finger2 | Ring of Precision (1491) (or Legionnaire's Band (19513)) | Blackfathom Deeps [dungeon] | 20.0 ranged_attack_power points (1.39 DPS) | yes | Legionnaire's Band (19513, +0.00 DPS) [rep]; Monkey Ring (6748, -0.23 DPS) [quest]; Azora's Will (4999, -0.43 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Crescent Staff (6505) | Leaders of the Fang [quest] | sim-verified (110.4 DPS) | yes | Impaling Harpoon (5200, +0.00 DPS) [dungeon]; Felbough Arm (273037, +0.00 DPS) [world]; Kam's Walking Stick (2280, -0.07 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Crystalpine Stinger (13037) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Golemsight Long Gun (273029, +0.00 DPS) [dungeon]; Ironweaver (13137, -0.32 DPS) [world_drop]; Harpyclaw Short Bow (13019, -0.71 DPS) [world_drop] |

**New at 30:** head: Enduring Cap; neck: Archmagister's Faceted Pendant; shoulder: Death Speaker Mantle; back: Tigerstrike Mantle; chest: Garb of Florid Feathers; wrist: Glowing Magical Bracelets; hands: Hotshot Pilot's Gloves; waist: Stalker's Leather Belt; legs: Ferine Leggings; feet: Trapper's Leather Boots; finger2: Ring of Precision; main_hand: Crescent Staff; ranged: Crystalpine Stinger

No-known-source sample (15 of 446, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 40 (troll, 1000000000000000-0051550001503050-000000000000000000)

Set DPS (verified): 169.4. Weights run: 7.8s. Verify run: 25.0s. 671 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.368 ± 0.023, crit=1.870 ± 0.038 per rating point (14 rating = 1%, 26.176 per %), hit=2.314 ± 0.095 per rating point (10 rating = 1%, 23.139 per %), melee_haste=not significant (5.772 ± 1.473), intellect=2.916 ± 0.063, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 52.5 ranged_attack_power points (4.14 DPS) | yes | Thinking Cap (2624, -0.23 DPS) [world]; Papal Fez (9431, -0.23 DPS) [dungeon]; Enduring Cap (3020, -0.46 DPS) [world_drop] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 37.0 ranged_attack_power points (2.92 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.59 DPS) [world]; Kaleidoscope Chain (13084, -0.94 DPS) [world_drop]; Glowing Eye of Mordresh (10769, -1.08 DPS) [dungeon] |
| shoulder | Sheepshear Mantle (13115) | World drop [world_drop] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Windchaser Amice (14432, -0.35 DPS, sim-verified) [world_drop]; Inquisitor's Shawl (19507, -0.69 DPS) [dungeon]; Nightscape Shoulders (8192, -0.76 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Cloak (6432, -0.40 DPS) [dungeon]; Parachute Cloak (10518, -0.40 DPS) [crafted]; Tigerstrike Mantle (13108, -0.40 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 104.0 ranged_attack_power points (8.19 DPS) | yes | Garb of Florid Feathers (273026, -2.27 DPS, sim-verified) [dungeon]; Tough Scorpid Breastplate (8203, -2.60 DPS) [crafted]; Wolffear Harness (13110, -3.68 DPS) [world_drop] |
| wrist | Tough Scorpid Bracers (8205) | Leatherworking [crafted] | 41.1 ranged_attack_power points (3.24 DPS) | yes | Green Leather Bracers (4259, -0.39 DPS, sim-verified) [crafted]; Warden's Wristbands (14600, -0.60 DPS) [world_drop]; Windchaser Cuffs (14429, -0.64 DPS) [world_drop] |
| hands | Tough Scorpid Gloves (8204) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stalker's Leather Gloves (252526, -0.46 DPS) [crafted]; Aurora Gloves (4042, -0.76 DPS) [world_drop]; Dragonscale Gauntlets (8347, -1.07 DPS) [crafted] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 47.8 ranged_attack_power points (3.77 DPS) | yes | Deathmage Sash (10771, -0.32 DPS) [dungeon]; Defiler's Lizardhide Girdle (20173, -0.32 DPS) [rep]; Defiler's Mail Girdle (20198, -0.32 DPS) [rep] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 70.7 ranged_attack_power points (5.57 DPS) | yes | Ferine Leggings (6690, -0.56 DPS, sim-verified) [dungeon]; Triprunner Dungarees (9624, -0.80 DPS) [quest]; Trapper's Leather Legguards (252517, -1.58 DPS) [crafted] |
| feet | Stalker's Mail Boots (252562) | Leatherworking [crafted] | 57.5 ranged_attack_power points (4.53 DPS) | yes | Swampwalker Boots (2276, -1.08 DPS) [world_drop]; Briar Tredders (10582, -1.13 DPS) [dungeon]; Worn Running Boots (9398, -1.16 DPS, sim-verified) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 33.7 ranged_attack_power points (2.65 DPS) | yes | Legionnaire's Band (19512, -0.53 DPS) [rep]; Disengagement Ring (276202, -0.53 DPS) [vendor]; Lavishly Jeweled Ring (1156, -0.74 DPS) [dungeon] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 30.3 ranged_attack_power points (2.39 DPS) | yes | Legionnaire's Band (19512, -0.27 DPS) [rep]; Disengagement Ring (276202, -0.27 DPS) [vendor]; Lavishly Jeweled Ring (1156, -0.48 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Felbough Arm (273037, -2.85 DPS) [world]; Steel Spear (250605, -3.09 DPS) [crafted]; Claw of the Shadowmancer (2912, -4.22 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Silencer (13138, -2.06 DPS, sim-verified) [world_drop]; Skystriker Bow (13020, -2.25 DPS) [world_drop]; Swiftwind (13038, -2.28 DPS) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Scout's Medallion; shoulder: Sheepshear Mantle; back: Darkspear Raider's Cloak; chest: Icemail Jerkin; wrist: Tough Scorpid Bracers; hands: Tough Scorpid Gloves; legs: Basilisk Hide Pants; feet: Stalker's Mail Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Bow of Searing Arrows

No-known-source sample (15 of 671, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 50 (troll, 5320000100000000-0051550001503050-000000000000000000)

Set DPS (verified): 238.3. Weights run: 8.1s. Verify run: 27.7s. 845 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.234 ± 0.018, crit=2.042 ± 0.032 per rating point (14 rating = 1%, 28.594 per %), hit=2.804 ± 0.083 per rating point (10 rating = 1%, 28.045 per %), melee_haste=14.385 ± 1.086, intellect=2.776 ± 0.046, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Tough Scorpid Helm (8208, -0.59 DPS, sim-verified) [crafted]; Khan's Helmet (14785, -1.31 DPS) [world_drop] |
| neck | Horizon Choker (13085) | World drop [world_drop] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Scout's Medallion (19535, -0.01 DPS) [rep]; Skibi's Pendant (13089, -0.64 DPS, sim-verified) [world_drop] |
| shoulder | Tough Scorpid Shoulders (8207) | Leatherworking [crafted] | 60.1 ranged_attack_power points (5.52 DPS) | yes | Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Ironfeather Shoulders (15067, -0.42 DPS) [crafted]; Stalker's Mail Shoulders (252566, -0.47 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Raider's Cloak (272076, -0.89 DPS) [vendor]; Blackflame Cape (13109, -0.89 DPS) [world_drop]; Serpentskin Cloak (8259, -1.19 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 105.7 ranged_attack_power points (9.70 DPS) | yes | Stone Guard's Chain Armor (220827, -0.69 DPS, sim-verified) [vendor]; Blazewind Breastplate (11193, -2.87 DPS) [quest]; Garb of Florid Feathers (273026, -3.21 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 48.5 ranged_attack_power points (4.45 DPS) | yes | Stalker's Mail Bracers (252570, -0.25 DPS) [crafted]; Forgotten Wraps (9433, -0.50 DPS) [world_drop]; Bracers of the Stone Princess (17714, -0.61 DPS) [dungeon] |
| hands | Stalker's Mail Gauntlets (252584) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tough Scorpid Gloves (8204, -1.19 DPS) [crafted]; Hibernal Gloves (8110, -1.51 DPS, sim-verified) [world_drop]; Stalker's Leather Gloves (252526, -1.70 DPS) [crafted] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 81.8 ranged_attack_power points (7.51 DPS) | yes | Stalker's Mail Belt (252588, -1.31 DPS, sim-verified) [crafted]; Defiler's Lizardhide Girdle (20174, -2.34 DPS) [rep]; Defiler's Mail Girdle (20196, -2.34 DPS) [rep] |
| legs | Stone Guard's Chain Legplates (220833) | Lady Palanseer [vendor] | 101.6 ranged_attack_power points (9.33 DPS) | yes | Imperial Red Pants (8251, -1.42 DPS, sim-verified) [world_drop]; Tough Scorpid Leggings (8206, -1.73 DPS) [crafted]; Searingscale Leggings (11749, -2.16 DPS) [dungeon] |
| feet | Fleetfoot Greaves (11627) | Blackrock Depths: Houndmaster Grebmar [dungeon] | 78.1 ranged_attack_power points (7.17 DPS) | yes | Stalker's Mail Sabatons (252576, -1.02 DPS) [crafted]; Rancor Boots (11865, -1.20 DPS, sim-verified) [quest]; Albino Crocscale Boots (17728, -1.23 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 48.0 ranged_attack_power points (4.41 DPS) | yes | Ring of the Underwood (2951, -1.44 DPS) [world_drop]; Falcon's Hook (7552, -1.74 DPS) [dungeon]; Legionnaire's Band (19511, -1.74 DPS) [rep] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 45.3 ranged_attack_power points (4.16 DPS) | yes | Ring of the Underwood (2951, -1.22 DPS, sim-verified) [world_drop]; Falcon's Hook (7552, -1.48 DPS) [dungeon]; Legionnaire's Band (19511, -1.48 DPS) [rep] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+4.0 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -1.41 DPS, sim-verified) [crafted] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warmonger (13052, -1.24 DPS) [world_drop]; Glowing Brightwood Staff (812, -1.57 DPS) [world_drop]; Claw of the Shadowmancer (2912, -6.55 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gryphonwing Long Bow (13022, -1.85 DPS, sim-verified) [dungeon]; Megashot Rifle (17717, -3.18 DPS) [dungeon]; Dark Iron Rifle (16004, -3.73 DPS) [crafted] |

**New at 50:** head: Eye of Theradras; neck: Horizon Choker; shoulder: Tough Scorpid Shoulders; back: Dark Phantom Cape; wrist: Deepfury Bracers; hands: Stalker's Mail Gauntlets; waist: Sagebrush Girdle; legs: Stone Guard's Chain Legplates; feet: Fleetfoot Greaves; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 845, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (troll, 5320000501000000-0051550001503050-500000000000000000)

Set DPS (verified): 348.3. Weights run: 7.8s. Verify run: 34.1s. 1595 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.208 ± 0.022, crit=2.541 ± 0.046 per rating point (14 rating = 1%, 35.568 per %), hit=3.378 ± 0.147 per rating point (10 rating = 1%, 33.782 per %), melee_haste=12.249 ± 1.971, intellect=3.059 ± 0.084, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 219.2 ranged_attack_power points (22.64 DPS) | yes | Warlord's Chain Helmet (16566, -3.91 DPS) [vendor]; Warlord's Chain Helm (231571, -3.91 DPS) [vendor]; Outlaw's Collar (279253, -8.31 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (19534) | Warsong Outriders [rep] | 48.1 ranged_attack_power points (4.97 DPS) | yes | Horizon Choker (13085, -0.55 DPS) [world_drop]; Skibi's Pendant (13089, -0.66 DPS) [world_drop]; Sunprism Pendant (20649, -0.99 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 122.2 ranged_attack_power points (12.62 DPS) | yes | Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-verified (348.3 DPS) | yes | Deathguard's Cloak (20068, -1.84 DPS, sim-verified) [rep]; Darkspear Raider's Cloak (272063, -1.98 DPS) [vendor]; Dark Phantom Cape (13122, -2.07 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 188.0 ranged_attack_power points (19.42 DPS) | yes | Warlord's Chain Chestpiece (16565, -0.69 DPS) [vendor]; Warlord's Chain Hauberk (231573, -0.69 DPS) [vendor]; Savage Gladiator Chain (11726, -8.65 DPS, sim-verified) [dungeon] |
| wrist | Windtalker's Wristguards (19582) | Warsong Outriders [rep] | 62.5 ranged_attack_power points (6.45 DPS) | yes | General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19587, -0.16 DPS) [rep]; Spitfire Bracers (20481, -0.63 DPS) [crafted] |
| hands | Mourning Handguards (279257) | Leatherworking [crafted] | sim-verified (348.3 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; Raider Gloves (272099, -1.86 DPS, sim-verified) [vendor] |
| waist | Defiler's Lizardhide Girdle (20171) (or Defiler's Mail Girdle (20195)) | The Defilers [rep] | 87.6 ranged_attack_power points (9.05 DPS) | yes | General's Chain Girdle (16572, +0.00 DPS) [pvp]; Defiler's Mail Girdle (20195, +0.00 DPS) [rep]; Sagebrush Girdle (17778, -0.28 DPS) [quest] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | 181.6 ranged_attack_power points (18.76 DPS) | yes | General's Chain Legguards (231574, -0.03 DPS) [pvp]; General's Chain Legplates (231567, -1.94 DPS) [vendor]; Outrider's Leather Pants (22740, -5.17 DPS, sim-verified) [rep] |
| feet | Timmy's Galoshes (13402) | Stratholme: Timmy the Cruel [dungeon] | 85.1 ranged_attack_power points (8.79 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Fleetfoot Greaves (11627, -0.60 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 83.4 ranged_attack_power points (8.61 DPS) | yes | Masons Fraternity Ring (9533, -3.97 DPS) [quest]; Glowing Crystal Ring (18402, -4.82 DPS) [quest]; Band of the Penitent (13217, -4.94 DPS) [quest] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 53.8 ranged_attack_power points (5.56 DPS) | yes | Masons Fraternity Ring (9533, -0.92 DPS) [quest]; Glowing Crystal Ring (18402, -1.76 DPS) [quest]; Band of the Penitent (13217, -1.88 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (348.3 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (348.3 DPS) | yes | Ankh of Life (1713, -2.16 DPS, sim-verified) [world_drop] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (348.3 DPS) | yes | High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Glacial Blade (19099, -5.97 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (348.3 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -9.93 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; neck: Scout's Medallion; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Windtalker's Wristguards; hands: Mourning Handguards; waist: Defiler's Lizardhide Girdle; legs: Outrider's Chain Leggings; feet: Timmy's Galoshes; finger1: Don Julio's Band; finger2: Blackstone Ring; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1595, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60, raid preset (troll, 5000000000000000-0053550011503050-500050030000000000)

Set DPS (verified): 785.2. Weights run: 4.8s. Verify run: 49.4s. 1595 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.237 ± 0.025, crit=2.519 ± 0.050 per rating point (14 rating = 1%, 35.268 per %), hit=3.610 ± 0.158 per rating point (10 rating = 1%, 36.098 per %), melee_haste=14.124 ± 1.730, intellect=2.344 ± 0.051, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 192.3 ranged_attack_power points (34.47 DPS) | yes | Warlord's Chain Helmet (16566, -3.38 DPS) [vendor]; Warlord's Chain Helm (231571, -3.38 DPS) [vendor]; Outlaw's Collar (279253, -11.89 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (19534) | Warsong Outriders [rep] | 48.6 ranged_attack_power points (8.70 DPS) | yes | Skibi's Pendant (13089, -1.16 DPS) [world_drop]; Sunprism Pendant (20649, -1.74 DPS) [quest]; Horizon Choker (13085, -2.82 DPS) [world_drop] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 122.7 ranged_attack_power points (21.98 DPS) | yes | Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-verified (785.2 DPS) | yes | Dark Phantom Cape (13122, -3.58 DPS) [world_drop]; Deathguard's Cloak (20068, -3.80 DPS, sim-verified) [rep]; Blackflame Cape (13109, -5.32 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 191.2 ranged_attack_power points (34.25 DPS) | yes | Warlord's Chain Chestpiece (16565, -3.17 DPS) [vendor]; Warlord's Chain Hauberk (231573, -3.17 DPS) [vendor]; Savage Gladiator Chain (11726, -18.63 DPS, sim-verified) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 61.5 ranged_attack_power points (11.02 DPS) | yes | General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Windtalker's Wristguards (19582, -0.85 DPS) [rep]; Legionnaire's Chain Bracers (16517, -1.16 DPS) [pvp] |
| hands | Mourning Handguards (279257) | Leatherworking [crafted] | sim-verified (785.2 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; Raider Gloves (272099, +0.00 DPS) [vendor] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 76.7 ranged_attack_power points (13.74 DPS) | yes | General's Chain Girdle (16572, +0.00 DPS) [pvp]; Defiler's Lizardhide Girdle (20171, -0.28 DPS) [rep]; Defiler's Mail Girdle (20195, -0.28 DPS) [rep] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | 184.7 ranged_attack_power points (33.09 DPS) | yes | General's Chain Legguards (231574, -2.01 DPS) [pvp]; General's Chain Legplates (231567, -5.42 DPS) [vendor]; Outrider's Leather Pants (22740, -11.59 DPS, sim-verified) [rep] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 78.8 ranged_attack_power points (14.12 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Timmy's Galoshes (13402, -0.06 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 85.4 ranged_attack_power points (15.30 DPS) | yes | Masons Fraternity Ring (9533, -7.17 DPS) [quest]; Legionnaire's Band (19510, -8.92 DPS) [rep]; Band of the Penitent (13217, -8.98 DPS) [quest] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 56.1 ranged_attack_power points (10.05 DPS) | yes | Masons Fraternity Ring (9533, -1.93 DPS) [quest]; Legionnaire's Band (19510, -3.67 DPS) [rep]; Band of the Penitent (13217, -3.73 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (785.2 DPS) | yes | Frozen Heart of the Mountain (249469, -2.81 DPS, sim-verified) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (785.2 DPS) | yes | Frozen Heart of the Mountain (249469, -3.21 DPS, sim-verified) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (785.2 DPS) | yes | High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; High Warlord's War Staff (234549, -3.13 DPS) [pvp]; Glacial Blade (19099, -10.88 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (785.2 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Eaglehorn Long Bow (13023, -6.30 DPS, sim-verified) [world_drop] |

**New at 60:** head: Black Dragonscale Helm; neck: Scout's Medallion; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Mourning Handguards; legs: Outrider's Chain Leggings; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Blackstone Ring; trinket2: Molten Heart of the Mountain; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1595, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

