# Leveling BiS: Enhancement

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-253100000000000000-0000000000000000)

Set DPS (verified): 44.5. Weights run: 2.7s. Verify run: 1.4s. 225 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.205 ± 0.016, crit=0.317 ± 0.005 per rating point (14 rating = 1%, 4.432 per %), hit=0.433 ± 0.004 per rating point (10 rating = 1%, 4.331 per %), melee_haste=2.472 ± 0.427

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (0.55 DPS) | yes | Brawler's Leather Hood (252504, -0.43 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 1.2 attack_power points (0.04 DPS) | yes | Erudite's Amulet (277204, -0.01 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 1.0 attack_power points (0.04 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Grave Shroud (279865, -0.05 DPS) [quest]; Catacomb Cloak (279899, -0.07 DPS) [quest]; Dark Leather Cloak (2316, -0.12 DPS) [crafted] |
| chest | Totemic Leather Armor (252435) | Leatherworking [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Armor of the Fang (6473, -0.07 DPS) [dungeon]; Brawler's Leather Armor (252490, -0.09 DPS) [crafted]; Defender's Leather Armor (252434, -0.43 DPS, sim-verified) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 4.8 attack_power points (0.17 DPS) | yes | Bristlebark Bindings (14569, -0.01 DPS) [world_drop]; Light Leather Bracers (7281, -0.09 DPS) [crafted]; Cloudy Gustwoven Bracers (276999, -0.10 DPS) [crafted] |
| hands | Gold-flecked Gloves (5195) | The Deadmines: Sneed [dungeon] | 14.0 attack_power points (0.48 DPS) | yes | Blackened Defias Gloves (10401, -0.07 DPS) [dungeon]; Foreman's Gloves (2167, -0.14 DPS) [world]; Gloves of the Fang (10413, -0.16 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.62 DPS) | yes | Ruffian Belt (5975, -0.21 DPS, sim-verified) [world]; Support Girdle (1215, -0.28 DPS) [world]; Brawler's Leather Belt (252428, -0.32 DPS) [crafted] |
| legs | Totemic Leather Pants (252446) | Leatherworking [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Deepgrave Trousers (279900, -0.12 DPS) [quest]; Brawler's Leather Pants (252500, -0.14 DPS) [crafted]; Defender's Leather Pants (252445, -0.50 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 11.0 attack_power points (0.38 DPS) | yes | Defender's Leather Boots (252441, -0.04 DPS) [crafted]; Totemic Leather Boots (252442, -0.04 DPS) [crafted]; Forest Leather Boots (3057, -0.10 DPS) [world_drop] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.8 attack_power points (0.30 DPS) | yes | The 1 Ring (8350, -0.23 DPS) [world]; Ring of the Moon (12052, -0.24 DPS) [world_drop]; Signet of the Zhevra (285330, -0.26 DPS) [world] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | The 1 Ring (8350, -0.20 DPS, sim-verified) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop]; Signet of the Zhevra (285330, -0.23 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | sim-verified (+3.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Living Root (6631, -0.06 DPS) [dungeon]; Night Reaver (1318, -0.35 DPS) [dungeon]; The Axe of Severing (23171, -3.09 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Lambent Scale Cloak; chest: Totemic Leather Armor; wrist: Bravo's Armbands; hands: Gold-flecked Gloves; waist: Blackened Defias Belt; legs: Totemic Leather Pants; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: Smite's Mighty Hammer

No-known-source sample (15 of 225, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher

### Band 30 (dwarf, 0000000000000000-253130030004000000-0000000000000000)

Set DPS (verified): 103.4. Weights run: 2.8s. Verify run: 1.5s. 366 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.350 ± 0.073, crit=0.562 ± 0.017 per rating point (14 rating = 1%, 7.874 per %), hit=0.629 ± 0.006 per rating point (10 rating = 1%, 6.285 per %), melee_haste=3.222 ± 0.285

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (0.85 DPS) | yes | Azure Gustwoven Hood (277050, -0.21 DPS) [crafted]; Defender's Leather Hood (252447, -0.28 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.50 DPS, sim-verified) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.49 DPS) | yes | Kaleidoscope Chain (13084, -0.16 DPS) [world_drop]; River Pride Choker (13087, -0.21 DPS) [world_drop]; Sentinel's Medallion (19541, -0.40 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 13.9 attack_power points (0.49 DPS) | yes | Bristlebark Amice (14573, -0.20 DPS) [world_drop]; Watchman Pauldrons (7727, -0.28 DPS) [dungeon]; Barbaric Shoulders (5964, -0.56 DPS, sim-verified) [crafted] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 attack_power points (0.35 DPS) | yes | Sergeant Major's Cape (16315, +0.00 DPS, sim-verified) [pvp]; Hawkeye's Cloak (14593, -0.05 DPS) [world_drop]; Slayer's Cape (14752, -0.07 DPS) [world_drop] |
| chest | Thick Murloc Armor (5782) (or Nightwalker Armor (2234)) | Leatherworking [crafted] | 18.0 attack_power points (0.64 DPS) | yes | Nightwalker Armor (2234, +0.00 DPS) [world]; Raptorbane Armor (3566, -0.07 DPS) [quest]; Totemic Leather Tunic (252451, -0.07 DPS) [crafted] |
| wrist | Bands of Serra'kis (6902) | Blackfathom Deeps: Old Serra'kis [dungeon] | 12.0 attack_power points (0.42 DPS) | yes | Hawkeye's Bracers (14590, -0.07 DPS) [world_drop]; Cultist's Armguards (270032, -0.07 DPS) [quest]; Technician's Bracers (270042, -0.07 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 16.1 attack_power points (0.57 DPS) | yes | Heavy Earthen Gloves (7359, -0.00 DPS) [crafted]; Brawler Gloves (720, -0.00 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.07 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (0.85 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Prowler's Leather Belt (252459, -0.14 DPS) [crafted]; Blackened Defias Belt (10403, -0.21 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (0.92 DPS) | yes | Totemic Leather Pants (252446, -0.28 DPS) [crafted]; Totemic Leather Leggings (252458, -0.28 DPS) [crafted]; Defender's Leather Pants (252445, -2.92 DPS, sim-verified) [crafted] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 attack_power points (0.42 DPS) | yes | Brawler's Leather Boots (252439, -0.01 DPS) [crafted]; Draftsman Boots (6668, -0.07 DPS) [quest]; Defender's Leather Boots (252441, -0.07 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 17.1 attack_power points (0.60 DPS) | yes | Tiger Band (6749, -0.18 DPS) [quest]; Ironspine's Eye (7686, -0.21 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.25 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 14.1 attack_power points (0.50 DPS) | yes | Ironspine's Eye (7686, -0.10 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.14 DPS) [dungeon]; Tiger Band (6749, -0.56 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (103.4 DPS) | yes | Corpsemaker (6687, -0.16 DPS) [dungeon]; Darkspear Raider's Reaper (272082, -0.61 DPS) [vendor]; Viscous Hammer (13045, -23.79 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Thick Murloc Armor; wrist: Bands of Serra'kis; hands: Insignia Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 366, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler

### Band 40 (dwarf, 0000000000000000-253130030005102051-0000000000000000)

Set DPS (verified): 109.9. Weights run: 3.4s. Verify run: 1.6s. 592 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.418 ± 0.037, crit=0.639 ± 0.012 per rating point (14 rating = 1%, 8.943 per %), hit=0.529 ± 0.006 per rating point (10 rating = 1%, 5.295 per %), melee_haste=3.765 ± 0.092

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 34.9 attack_power points (1.36 DPS) | yes | White Bandit Mask (10008, -0.32 DPS) [crafted]; Tusken Helm (6686, -0.35 DPS) [dungeon]; Hard Gold Coif (250537, -1.35 DPS, sim-verified) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.78 DPS) | yes | Ghostshard Talisman (7731, -0.23 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.40 DPS) [world_drop]; River Pride Choker (13087, -0.47 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 attack_power points (0.86 DPS) | yes | Imperial Leather Spaulders (4737, -0.16 DPS) [dungeon]; Wrangling Spaulders (15698, -0.20 DPS) [quest]; Sunburn Spaulders (274751, -0.21 DPS) [vendor] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 14.5 attack_power points (0.56 DPS) | yes | Wolfmaster Cape (6314, -0.18 DPS) [dungeon]; Hawkeye's Cloak (14593, -0.22 DPS) [world_drop]; Dark Hooded Cape (5257, -2.29 DPS, sim-verified) [world] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 30.9 attack_power points (1.20 DPS) | yes | Avenger's Armor (1488, -0.04 DPS) [dungeon]; Shining Silver Breastplate (2870, -0.11 DPS) [crafted]; Golden Scale Cuirass (3845, -0.11 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.78 DPS) | yes | Pugilist Bracers (4438, -0.16 DPS) [dungeon]; Ravager's Armguards (14770, -0.17 DPS) [world_drop]; Yorgen Bracers (13012, -0.26 DPS) [world_drop] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 attack_power points (1.24 DPS) | yes | Gloves of Holy Might (867, -0.12 DPS) [world_drop]; Scarlet Gauntlets (10331, -0.15 DPS) [dungeon]; Reticulated Bone Gauntlets (9435, -0.23 DPS) [world_drop] |
| waist | Boar Champion's Belt (10768) | Razorfen Downs: Ragglesnout [dungeon] | 30.0 attack_power points (1.17 DPS) | yes | Highlander's Leather Girdle (20116, -0.00 DPS) [rep]; Girdle of Golem Strength (9405, -0.23 DPS) [world_drop]; Scarlet Belt (10329, -0.23 DPS) [dungeon] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.63 DPS) | yes | Firemane Leggings (13129, -0.16 DPS) [world_drop]; Orcish War Leggings (7929, -0.31 DPS) [crafted]; Legguards of the Vault (9396, -0.54 DPS) [dungeon] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 24.9 attack_power points (0.97 DPS) | yes | Skirmisher's Mail Boots (252564, -0.11 DPS) [crafted]; Ironheel Boots (4653, -0.19 DPS) [quest]; Blackforge Greaves (6423, -2.14 DPS, sim-verified) [dungeon] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.78 DPS) | yes | Protector's Band (19515, -0.03 DPS) [rep]; Thunderbrow Ring (13097, -0.11 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.78 DPS) | yes | Protector's Band (19515, -0.03 DPS) [rep]; Thunderbrow Ring (13097, -0.11 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | The Jackhammer (9423) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-verified (109.9 DPS) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor]; Bonebiter (6830, -8.92 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Boar Champion's Belt; legs: Scarlet Leggings; feet: Prowler's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: The Jackhammer

No-known-source sample (15 of 592, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 50 (dwarf, 5500000000000000-253130030005102051-0000000000000000)

Set DPS (verified): 140.9. Weights run: 3.4s. Verify run: 1.9s. 762 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.606 ± 0.051, crit=0.787 ± 0.015 per rating point (14 rating = 1%, 11.024 per %), hit=0.804 ± 0.009 per rating point (10 rating = 1%, 8.035 per %), melee_haste=4.279 ± 0.068

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 48.0 attack_power points (1.83 DPS) | yes | Raging Berserker's Helm (7719, -0.42 DPS) [dungeon]; Bloomsprout Headpiece (17767, -0.46 DPS) [dungeon]; Knight-Lieutenant's Mail Helmet (223075, -0.57 DPS) [vendor] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.76 DPS) | yes | Skibi's Pendant (13089, -0.08 DPS) [world_drop]; Ghostshard Talisman (7731, -0.23 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.36 DPS) [world_drop] |
| shoulder | Knight-Lieutenant's Mail Epaulets (223073) | Captain Dirgehammer [vendor] | 29.0 attack_power points (1.11 DPS) | yes | Failed Flying Experiment (9647, -0.11 DPS) [quest]; Skulker's Leather Shoulder (252535, -0.17 DPS) [crafted]; Prowler's Leather Shoulder (252534, -1.35 DPS, sim-verified) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 20.5 attack_power points (0.78 DPS) | yes | Sergeant Major's Cape (16336, -0.18 DPS) [pvp]; Dark Hooded Cape (5257, -0.24 DPS) [world]; Bloodlust Cape (14801, -1.71 DPS, sim-verified) [world_drop] |
| chest | Knight's Mail Armor (223078) | Captain Dirgehammer [vendor] | sim-verified (140.9 DPS) | yes | Warbear Harness (15064, -0.08 DPS) [crafted]; Kolkar Marauder Chain (6773, -0.11 DPS) [quest]; Mixologist's Tunic (12793, -1.67 DPS, sim-verified) [dungeon] |
| wrist | Arena Bands (18711) (or Bracers of the Stone Princess (17714)) | Arena Treasure Chest [world] | 28.0 attack_power points (1.07 DPS) | yes | Bracers of the Stone Princess (17714, +0.00 DPS) [dungeon]; Prowler's Leather Bracers (252539, -0.22 DPS) [crafted]; Branded Leather Bracers (19508, -0.30 DPS) [dungeon] |
| hands | Prowler's Leather Gauntlets (252547) | Leatherworking [crafted] | 33.4 attack_power points (1.27 DPS) | yes | Raider Gloves (272100, -0.00 DPS) [vendor]; Maddening Gauntlets (11867, -0.02 DPS) [quest]; Gauntlets of Divinity (7724, -0.06 DPS) [dungeon] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (1.75 DPS) | yes | Belt of the Gladiator (13134, -0.38 DPS) [world_drop]; Prowler's Leather Waistguard (252473, -0.41 DPS) [crafted]; Skulker's Leather Waistguard (252474, -0.52 DPS) [crafted] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.60 DPS) | yes | Gryphon Rider's Leggings (9652, -0.11 DPS) [quest]; Firemane Leggings (13129, -0.15 DPS) [world_drop]; Serpentskin Leggings (8262, -0.24 DPS) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 32.7 attack_power points (1.24 DPS) | yes | Skulker's Leather Boots (252469, -0.11 DPS) [crafted]; Shadefiend Boots (11675, -0.23 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.24 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 28.0 attack_power points (1.07 DPS) | yes | Mark of Kern (2262, -0.31 DPS) [dungeon]; Assault Band (13095, -0.31 DPS) [world_drop]; Thunderbrow Ring (13097, -0.39 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 25.4 attack_power points (0.97 DPS) | yes | Mark of Kern (2262, -0.21 DPS) [dungeon]; Assault Band (13095, -0.21 DPS) [world_drop]; Thunderbrow Ring (13097, -0.29 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smoking Heart of the Mountain (11811, +0.00 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ragehammer (10626, +0.00 DPS) [dungeon]; Thorium Greatmace (250613, -0.50 DPS) [crafted]; Darkspear Raider's Reaper (272080, -0.87 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Embrace of the Lycan; shoulder: Knight-Lieutenant's Mail Epaulets; back: Blackveil Cape; chest: Knight's Mail Armor; wrist: Arena Bands; hands: Prowler's Leather Gauntlets; waist: Girdle of Beastial Fury; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 762, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 60 (dwarf, 5533220000000000-253130030005102051-0000000000000000)

Set DPS (verified): 202.7. Weights run: 3.6s. Verify run: 1.8s. 1751 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.644 ± 0.056, crit=0.869 ± 0.018 per rating point (14 rating = 1%, 12.172 per %), hit=0.974 ± 0.012 per rating point (10 rating = 1%, 9.739 per %), melee_haste=4.623 ± 0.114

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 48.0 attack_power points (1.82 DPS) | yes | Warbear Helm (252485, -0.11 DPS) [crafted]; Blue Suede Hat (252482, -0.14 DPS) [crafted]; Face of The Five Thunders (227021, -0.15 DPS) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 36.2 attack_power points (1.37 DPS) | yes | Beads of Ogre Might (22150, -0.09 DPS) [quest]; Amulet of the Darkmoon (19491, -0.15 DPS) [quest]; Imperial Jewel (11933, -0.16 DPS) [dungeon] |
| shoulder | Highlander's Leather Shoulders (20059) | The League of Arathor [rep] | 41.6 attack_power points (1.58 DPS) | yes | Highlander's Lizardhide Shoulders (20060, -0.15 DPS) [rep]; Golden Mantle of the Dawn (19058, -0.32 DPS) [crafted]; Black Dragonscale Shoulders (15051, -1.60 DPS, sim-verified) [crafted] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 37.7 attack_power points (1.43 DPS) | yes | Cloak of the Honor Guard (20073, -0.02 DPS) [rep]; Shroud of Domination (22337, -0.14 DPS) [dungeon]; Cape of the Black Baron (13340, -0.31 DPS) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (202.7 DPS) | yes | Obsidian Mail Tunic (22191, -0.06 DPS) [crafted]; Cadaverous Armor (14637, -0.33 DPS) [dungeon]; Tunic of Undead Slaying (23089, -11.04 DPS, sim-verified) [world] |
| wrist | Windtalker's Wristguards (19582) | Silverwing Sentinels [rep] | sim-verified (202.7 DPS) | yes | Forest Stalker's Bracers (19587, -0.14 DPS) [rep]; Tranquil Wristguards (279256, -0.23 DPS) [crafted]; Wristwraps of Undead Slaying (23093, -4.41 DPS, sim-verified) [world] |
| hands | Timbermaw Brawlers (19049) | Leatherworking [crafted] | 51.0 attack_power points (1.93 DPS) | yes | Studded Timbermaw Brawlers (227809, -0.17 DPS) [vendor]; Cadaverous Gloves (14640, -0.27 DPS) [dungeon]; Devilsaur Gauntlets (15063, -0.41 DPS) [crafted] |
| waist | Dense Timbermaw Belt (227807) | Meilosh [vendor] | 64.0 attack_power points (2.43 DPS) | yes | Ferocity of the Timbermaw (227805, -0.04 DPS) [vendor]; Might of the Timbermaw (19044, -0.49 DPS) [crafted]; Windseeker's Belt (272410, -0.61 DPS) [vendor] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 67.6 attack_power points (2.56 DPS) | yes | Devilsaur Leggings (15062, -0.36 DPS) [crafted]; Sentinel's Chain Leggings (237819, -0.42 DPS) [vendor]; Black Dragonscale Leggings (15052, -0.52 DPS) [crafted] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 43.1 attack_power points (1.64 DPS) | yes | Pads of the Dread Wolf (13210, -0.12 DPS) [dungeon]; Drudge Boots (21532, -0.24 DPS) [quest]; Boots of Ferocity (22472, -0.33 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (202.7 DPS) | yes | Blackstone Ring (17713, -0.31 DPS) [dungeon]; Band of the Ogre King (18522, -0.38 DPS) [dungeon]; Naglering (11669, -5.19 DPS, sim-verified) [dungeon] |
| finger2 | Protector's Band (19514) | Silverwing Sentinels [rep] | sim-verified (202.7 DPS) | yes | Blackstone Ring (17713, -0.05 DPS) [dungeon]; Band of the Ogre King (18522, -0.12 DPS) [dungeon]; Naglering (11669, -3.55 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (202.7 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (202.7 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-verified (202.7 DPS) | yes | Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Seeping Willow (12969, -15.26 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** neck: Medallion of the Dawn; shoulder: Highlander's Leather Shoulders; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Windtalker's Wristguards; hands: Timbermaw Brawlers; waist: Dense Timbermaw Belt; legs: Warbear Woolies; feet: Scalegut Treaders; finger1: Don Julio's Band; finger2: Protector's Band; trinket1: Burst of Knowledge; trinket2: Darkmoon Card: Maelstrom; main_hand: The Unstoppable Force

No-known-source sample (15 of 1751, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

## Horde

### Band 20 (orc, 0000000000000000-253100000000000000-0000000000000000)

Set DPS (verified): 45.7. Weights run: 2.7s. Verify run: 1.2s. 205 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.205 ± 0.016, crit=0.317 ± 0.005 per rating point (14 rating = 1%, 4.432 per %), hit=0.433 ± 0.004 per rating point (10 rating = 1%, 4.331 per %), melee_haste=2.472 ± 0.427

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (0.55 DPS) | yes | Brawler's Leather Hood (252504, -0.50 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 1.2 attack_power points (0.04 DPS) | yes | Erudite's Amulet (277204, -0.01 DPS) [quest] |
| shoulder | Slime-encrusted Pads (6461) | Wailing Caverns: Mutanus the Devourer [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Serpent's Shoulders (5404, -0.49 DPS, sim-verified) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Grave Shroud (279865, -0.05 DPS) [quest]; Catacomb Cloak (279899, -0.07 DPS) [quest]; Subterranean Cape (14149, -0.07 DPS) [dungeon] |
| chest | Totemic Leather Armor (252435) | Leatherworking [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Armor of the Fang (6473, -0.07 DPS) [dungeon]; Brawler's Leather Armor (252490, -0.09 DPS) [crafted]; Defender's Leather Armor (252434, -0.75 DPS, sim-verified) [crafted] |
| wrist | Bristlebark Bindings (14569) | World drop [world_drop] | 4.6 attack_power points (0.16 DPS) | yes | Light Leather Bracers (7281, -0.08 DPS) [crafted]; Cloudy Gustwoven Bracers (276999, -0.09 DPS) [crafted]; Azure Gustwoven Bracers (277023, -0.09 DPS) [crafted] |
| hands | Gold-flecked Gloves (5195) | The Deadmines: Sneed [dungeon] | 14.0 attack_power points (0.48 DPS) | yes | Blackened Defias Gloves (10401, -0.07 DPS) [dungeon]; Foreman's Gloves (2167, -0.14 DPS) [world]; Gloves of the Fang (10413, -0.16 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.62 DPS) | yes | Ruffian Belt (5975, -0.22 DPS, sim-verified) [world]; Support Girdle (1215, -0.28 DPS) [world]; Brawler's Leather Belt (252428, -0.32 DPS) [crafted] |
| legs | Totemic Leather Pants (252446) | Leatherworking [crafted] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Deepgrave Trousers (279900, -0.12 DPS) [quest]; Brawler's Leather Pants (252500, -0.14 DPS) [crafted]; Defender's Leather Pants (252445, -0.89 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 11.0 attack_power points (0.38 DPS) | yes | Defender's Leather Boots (252441, -0.04 DPS) [crafted]; Totemic Leather Boots (252442, -0.04 DPS) [crafted]; Forest Leather Boots (3057, -0.10 DPS) [world_drop] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.8 attack_power points (0.30 DPS) | yes | Loop of Sacrifice (281673, -0.10 DPS) [quest]; The 1 Ring (8350, -0.23 DPS) [world]; Ring of the Moon (12052, -0.24 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Loop of Sacrifice (281673, +0.00 DPS, sim-verified) [quest]; The 1 Ring (8350, -0.20 DPS) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | The Axe of Severing (23171) | Shadowfang Keep: Sever [dungeon] | 322.0 attack_power points (11.11 DPS) | yes | Forsaken Greataxe (251533, -0.66 DPS) [quest]; Smite's Mighty Hammer (7230, -0.80 DPS) [dungeon]; Hammerbone (270018, -4.29 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Scout's Medallion; shoulder: Slime-encrusted Pads; back: Lambent Scale Cloak; chest: Totemic Leather Armor; wrist: Bristlebark Bindings; hands: Gold-flecked Gloves; waist: Blackened Defias Belt; legs: Totemic Leather Pants; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: The Axe of Severing

No-known-source sample (15 of 205, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield; 14389 Durability Shoulderpads; 14705 Brackwater Chain Shield

### Band 30 (orc, 0000000000000000-253130030004000000-0000000000000000)

Set DPS (verified): 95.8. Weights run: 2.8s. Verify run: 1.4s. 349 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.350 ± 0.073, crit=0.562 ± 0.017 per rating point (14 rating = 1%, 7.874 per %), hit=0.629 ± 0.006 per rating point (10 rating = 1%, 6.285 per %), melee_haste=3.222 ± 0.285

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (0.85 DPS) | yes | Cloudy Gustwoven Hood (277042, -0.21 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.21 DPS) [crafted]; Defender's Leather Hood (252447, -0.28 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.49 DPS) | yes | Kaleidoscope Chain (13084, -0.16 DPS) [world_drop]; River Pride Choker (13087, -0.21 DPS) [world_drop]; Scout's Medallion (19537, -0.40 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 13.9 attack_power points (0.49 DPS) | yes | Barbaric Shoulders (5964, -0.07 DPS) [crafted]; Bristlebark Amice (14573, -0.20 DPS) [world_drop]; Watchman Pauldrons (7727, -0.28 DPS) [dungeon] |
| back | Wildhunter Cloak (16658) (or Wolfmaster Cape (6314)) | The Hunt Completed [quest] | 10.0 attack_power points (0.35 DPS) | yes | Wolfmaster Cape (6314, +0.00 DPS) [dungeon]; Hawkeye's Cloak (14593, -0.05 DPS) [world_drop]; Slayer's Cape (14752, -0.07 DPS) [world_drop] |
| chest | Thick Murloc Armor (5782) (or Nightwalker Armor (2234)) | Leatherworking [crafted] | 18.0 attack_power points (0.64 DPS) | yes | Nightwalker Armor (2234, +0.00 DPS) [world]; Totemic Leather Tunic (252451, -0.07 DPS) [crafted]; Defender's Leather Armor (252434, -0.10 DPS) [crafted] |
| wrist | Bands of Serra'kis (6902) | Blackfathom Deeps: Old Serra'kis [dungeon] | 12.0 attack_power points (0.42 DPS) | yes | Hawkeye's Bracers (14590, -0.07 DPS) [world_drop]; Cultist's Armguards (270032, -0.07 DPS) [quest]; Technician's Bracers (270042, -0.07 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 16.1 attack_power points (0.57 DPS) | yes | Heavy Earthen Gloves (7359, -0.00 DPS) [crafted]; Brawler Gloves (720, -0.00 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.07 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (0.85 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Prowler's Leather Belt (252459, -0.14 DPS) [crafted]; Blackened Defias Belt (10403, -0.21 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (0.92 DPS) | yes | Totemic Leather Pants (252446, -0.28 DPS) [crafted]; Totemic Leather Leggings (252458, -0.28 DPS) [crafted]; Defender's Leather Pants (252445, -2.58 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 11.8 attack_power points (0.42 DPS) | yes | Draftsman Boots (6668, -0.06 DPS) [quest]; Defender's Leather Boots (252441, -0.06 DPS) [crafted]; Totemic Leather Boots (252442, -0.06 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 17.1 attack_power points (0.60 DPS) | yes | Tiger Band (6749, -0.18 DPS) [quest]; Ironspine's Eye (7686, -0.21 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.25 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 14.1 attack_power points (0.50 DPS) | yes | Tiger Band (6749, -0.07 DPS) [quest]; Ironspine's Eye (7686, -0.10 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.14 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (95.8 DPS) | yes | Corpsemaker (6687, -0.16 DPS) [dungeon]; Darkspear Raider's Reaper (272082, -0.61 DPS) [vendor]; Viscous Hammer (13045, -22.79 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wildhunter Cloak; chest: Thick Murloc Armor; wrist: Bands of Serra'kis; hands: Insignia Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 349, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots

### Band 40 (orc, 0000000000000000-253130030005102051-0000000000000000)

Set DPS (verified): 112.0. Weights run: 3.4s. Verify run: 1.5s. 555 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.418 ± 0.037, crit=0.639 ± 0.012 per rating point (14 rating = 1%, 8.943 per %), hit=0.529 ± 0.006 per rating point (10 rating = 1%, 5.295 per %), melee_haste=3.765 ± 0.092

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 34.9 attack_power points (1.36 DPS) | yes | White Bandit Mask (10008, -0.32 DPS) [crafted]; Tusken Helm (6686, -0.35 DPS) [dungeon]; Hard Gold Coif (250537, -1.88 DPS, sim-verified) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.78 DPS) | yes | Ghostshard Talisman (7731, -0.23 DPS) [dungeon]; Ethereal Talisman (4430, -0.32 DPS) [quest]; Kaleidoscope Chain (13084, -0.40 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 attack_power points (0.86 DPS) | yes | Imperial Leather Spaulders (4737, -0.16 DPS) [dungeon]; Wrangling Spaulders (15698, -0.20 DPS) [quest]; Sunburn Spaulders (274751, -0.21 DPS) [vendor] |
| back | First Sergeant's Cloak (16340) | PvP rank 9 · First Sergeant · Horde [pvp] | 14.5 attack_power points (0.56 DPS) | yes | Wolfmaster Cape (6314, -0.18 DPS) [dungeon]; Wildhunter Cloak (16658, -0.18 DPS) [quest]; Dark Hooded Cape (5257, -2.13 DPS, sim-verified) [world] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 30.9 attack_power points (1.20 DPS) | yes | Avenger's Armor (1488, -0.04 DPS) [dungeon]; Shining Silver Breastplate (2870, -0.11 DPS) [crafted]; Golden Scale Cuirass (3845, -0.11 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.78 DPS) | yes | Windtalker's Wristguards (19584, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Pugilist Bracers (4438, -0.16 DPS) [dungeon] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 attack_power points (1.24 DPS) | yes | Gloves of Holy Might (867, +0.00 DPS, sim-verified) [world_drop]; Scarlet Gauntlets (10331, -0.15 DPS) [dungeon]; Reticulated Bone Gauntlets (9435, -0.23 DPS) [world_drop] |
| waist | Boar Champion's Belt (10768) | Razorfen Downs: Ragglesnout [dungeon] | 30.0 attack_power points (1.17 DPS) | yes | Defiler's Leather Girdle (20192, -0.00 DPS) [rep]; Tharg's Shoelace (9705, -0.16 DPS) [quest]; Scarlet Belt (10329, -0.23 DPS) [dungeon] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.63 DPS) | yes | Firemane Leggings (13129, -0.16 DPS) [world_drop]; Orcish War Leggings (7929, -0.31 DPS) [crafted]; Legguards of the Vault (9396, -0.54 DPS) [dungeon] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 24.9 attack_power points (0.97 DPS) | yes | Blackforge Greaves (6423, -0.11 DPS) [dungeon]; Skirmisher's Mail Boots (252564, -0.11 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.25 DPS) [crafted] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.78 DPS) | yes | Legionnaire's Band (19512, -0.03 DPS) [rep]; Thunderbrow Ring (13097, -0.11 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.78 DPS) | yes | Legionnaire's Band (19512, -0.03 DPS) [rep]; Thunderbrow Ring (13097, -0.11 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-verified (112.0 DPS) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Mograine's Might (7723, -0.15 DPS) [dungeon]; Darkspear Raider's Reaper (272081, -1.82 DPS, sim-verified) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: First Sergeant's Cloak; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Boar Champion's Belt; legs: Scarlet Leggings; feet: Prowler's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Fiery War Axe

No-known-source sample (15 of 555, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (orc, 5500000000000000-253130030005102051-0000000000000000)

Set DPS (verified): 147.5. Weights run: 3.4s. Verify run: 1.8s. 704 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.606 ± 0.051, crit=0.787 ± 0.015 per rating point (14 rating = 1%, 11.024 per %), hit=0.804 ± 0.009 per rating point (10 rating = 1%, 8.035 per %), melee_haste=4.279 ± 0.068

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 48.0 attack_power points (1.83 DPS) | yes | Raging Berserker's Helm (7719, -0.42 DPS) [dungeon]; Bloomsprout Headpiece (17767, -0.46 DPS) [dungeon]; Blood Guard's Mail Helmet (220820, -0.57 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Woven Ivy Necklace (19159, -0.02 DPS) [quest]; Ghostshard Talisman (7731, -0.15 DPS) [dungeon]; Zealous Shadowshard Pendant (17772, -1.56 DPS, sim-verified) [quest] |
| shoulder | Prowler's Leather Shoulder (252534) | Leatherworking [crafted] | 27.4 attack_power points (1.05 DPS) | yes | Blood Guard's Mail Epaulets (220823, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.05 DPS) [quest]; Skulker's Leather Shoulder (252535, -0.11 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 20.5 attack_power points (0.78 DPS) | yes | Bloodlust Cape (14801, -0.09 DPS) [world_drop]; First Sergeant's Cloak (16340, -0.18 DPS) [pvp]; Dark Hooded Cape (5257, -0.24 DPS) [world] |
| chest | Stone Guard's Mail Armor (220826) | Lady Palanseer [vendor] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Warbear Harness (15064, -0.08 DPS) [crafted]; Kolkar Marauder Chain (6773, -0.11 DPS) [quest]; Mixologist's Tunic (12793, -2.66 DPS, sim-verified) [dungeon] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Windtalker's Wristguards (19583, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Arena Bands (18711, -2.22 DPS, sim-verified) [world] |
| hands | Prowler's Leather Gauntlets (252547) | Leatherworking [crafted] | 33.4 attack_power points (1.27 DPS) | yes | Raider Gloves (272100, -0.00 DPS) [vendor]; Gauntlets of Divinity (7724, -0.06 DPS) [dungeon]; Gloves of Holy Might (867, -0.09 DPS) [world_drop] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (1.75 DPS) | yes | Belt of the Gladiator (13134, -0.38 DPS) [world_drop]; Prowler's Leather Waistguard (252473, -0.41 DPS) [crafted]; Skulker's Leather Waistguard (252474, -0.52 DPS) [crafted] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.60 DPS) | yes | Firemane Leggings (13129, -0.15 DPS) [world_drop]; Serpentskin Leggings (8262, -0.24 DPS) [world_drop]; Stone Guard's Mail Legplates (220834, -0.27 DPS) [vendor] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 32.7 attack_power points (1.24 DPS) | yes | Skulker's Leather Boots (252469, -0.11 DPS) [crafted]; Shadefiend Boots (11675, -0.23 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.24 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 28.0 attack_power points (1.07 DPS) | yes | White Bone Band (11862, -0.15 DPS) [quest]; Mark of Kern (2262, -0.31 DPS) [dungeon]; Assault Band (13095, -0.31 DPS) [world_drop] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 25.4 attack_power points (0.97 DPS) | yes | White Bone Band (11862, -0.06 DPS) [quest]; Mark of Kern (2262, -0.21 DPS) [dungeon]; Assault Band (13095, -0.21 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+4.4 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -2.08 DPS, sim-verified) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fiery War Axe (870, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, -0.50 DPS) [crafted]; Darkspear Raider's Reaper (272080, -0.87 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Embrace of the Lycan; neck: Skibi's Pendant; shoulder: Prowler's Leather Shoulder; back: Blackveil Cape; chest: Stone Guard's Mail Armor; wrist: Bracers of the Stone Princess; hands: Prowler's Leather Gauntlets; waist: Girdle of Beastial Fury; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 704, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (orc, 5533220000000000-253130030005102051-0000000000000000)

Set DPS (verified): 191.1. Weights run: 3.6s. Verify run: 1.8s. 1672 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.644 ± 0.056, crit=0.869 ± 0.018 per rating point (14 rating = 1%, 12.172 per %), hit=0.974 ± 0.012 per rating point (10 rating = 1%, 9.739 per %), melee_haste=4.623 ± 0.114

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 48.0 attack_power points (1.82 DPS) | yes | Champion's Mail Headguard (227155, +0.00 DPS) [pvp]; Warbear Helm (252485, -0.11 DPS) [crafted]; Blue Suede Hat (252482, -0.14 DPS) [crafted] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 36.2 attack_power points (1.37 DPS) | yes | Beads of Ogre Might (22150, -0.09 DPS) [quest]; Amulet of the Darkmoon (19491, -0.15 DPS) [quest]; Imperial Jewel (11933, -0.16 DPS) [dungeon] |
| shoulder | Defiler's Leather Shoulders (20194) | The Defilers [rep] | 41.6 attack_power points (1.58 DPS) | yes | Champion's Mail Pauldrons (227154, +0.00 DPS) [pvp]; Warlord's Mail Pauldrons (231654, +0.00 DPS) [vendor]; Black Dragonscale Shoulders (15051, -0.06 DPS) [crafted] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 37.7 attack_power points (1.43 DPS) | yes | Deathguard's Cloak (20068, -0.02 DPS) [rep]; Shroud of Domination (22337, -0.14 DPS) [dungeon]; Cape of the Black Baron (13340, -0.31 DPS) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Obsidian Mail Tunic (22191, -0.06 DPS) [crafted]; Cadaverous Armor (14637, -0.33 DPS) [dungeon]; Tunic of Undead Slaying (23089, -9.96 DPS, sim-verified) [world] |
| wrist | Windtalker's Wristguards (19582) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Stalker's Bracers (19587, -0.14 DPS) [rep]; Tranquil Wristguards (279256, -0.23 DPS) [crafted]; Wristwraps of Undead Slaying (23093, -3.61 DPS, sim-verified) [world] |
| hands | Timbermaw Brawlers (19049) | Leatherworking [crafted] | 51.0 attack_power points (1.93 DPS) | yes | General's Mail Vices (231655, +0.00 DPS) [vendor]; Studded Timbermaw Brawlers (227809, -0.17 DPS) [vendor]; Cadaverous Gloves (14640, -0.27 DPS) [dungeon] |
| waist | Dense Timbermaw Belt (227807) | Meilosh [vendor] | 64.0 attack_power points (2.43 DPS) | yes | Ferocity of the Timbermaw (227805, -0.04 DPS) [vendor]; Might of the Timbermaw (19044, -0.49 DPS) [crafted]; Windseeker's Belt (272410, -0.61 DPS) [vendor] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 67.6 attack_power points (2.56 DPS) | yes | General's Mail Legguards (231658, +0.00 DPS) [vendor]; Legionnaire's Mail Legguards (227156, -0.06 DPS) [pvp]; Devilsaur Leggings (15062, -0.36 DPS) [crafted] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 43.1 attack_power points (1.64 DPS) | yes | General's Mail Greaves (231656, +0.00 DPS) [vendor]; Pads of the Dread Wolf (13210, -0.12 DPS) [dungeon]; Blood Guard's Mail Greaves (227158, -0.21 DPS) [pvp] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -0.31 DPS) [dungeon]; Band of the Ogre King (18522, -0.38 DPS) [dungeon]; Naglering (11669, -4.30 DPS, sim-verified) [dungeon] |
| finger2 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -0.05 DPS) [dungeon]; Band of the Ogre King (18522, -0.12 DPS) [dungeon]; Naglering (11669, -3.45 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkmoon Card: Maelstrom (19289, +0.00 DPS) [quest]; Counterattack Lodestone (18537, -1.02 DPS) [dungeon]; Hand of Justice (11815, -1.09 DPS) [dungeon] |
| trinket2 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Eye of the Beast (13968, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon] |
| main_hand | Gravestone War Axe (13983) | Scholomance: Kirtonos the Herald [dungeon] | sim-verified (191.1 DPS) | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; The Unstoppable Force (19323, -4.26 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** neck: Medallion of the Dawn; shoulder: Defiler's Leather Shoulders; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Windtalker's Wristguards; hands: Timbermaw Brawlers; waist: Dense Timbermaw Belt; legs: Warbear Woolies; feet: Scalegut Treaders; finger1: Don Julio's Band; finger2: Legionnaire's Band; trinket2: Blackhand's Breadth; main_hand: Gravestone War Axe

No-known-source sample (15 of 1672, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

