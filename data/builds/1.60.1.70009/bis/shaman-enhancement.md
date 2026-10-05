# Leveling BiS: Enhancement

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-253100000000000000-0000000000000000)

Set DPS (verified): 44.5. Weights run: 1.0s. Verify run: 0.9s. 225 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=0.235 ± 0.025, crit=0.309 ± 0.007 per rating point (14 rating = 1%, 4.333 per %), hit=0.439 ± 0.005 per rating point (10 rating = 1%, 4.391 per %), melee_haste=2.440 ± 0.609

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (0.55 DPS) | yes | Brawler's Leather Hood (252504, -0.43 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 1.4 attack_power points (0.05 DPS) | yes | Erudite's Amulet (277204, -0.02 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 1.2 attack_power points (0.04 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Grave Shroud (279865, -0.05 DPS) [quest]; Catacomb Cloak (279899, -0.07 DPS) [quest]; Dark Leather Cloak (2316, -0.11 DPS) [crafted] |
| chest | Totemic Leather Armor (252435) | Leatherworking [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Armor of the Fang (6473, -0.07 DPS) [dungeon]; Brawler's Leather Armor (252490, -0.08 DPS) [crafted]; Defender's Leather Armor (252434, -0.43 DPS, sim-verified) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 4.9 attack_power points (0.17 DPS) | yes | Bristlebark Bindings (14569, -0.01 DPS) [world_drop]; Light Leather Bracers (7281, -0.09 DPS) [crafted]; Cloudy Gustwoven Bracers (276999, -0.10 DPS) [crafted] |
| hands | Gold-flecked Gloves (5195) | The Deadmines: Sneed [dungeon] | 14.0 attack_power points (0.48 DPS) | yes | Blackened Defias Gloves (10401, -0.07 DPS) [dungeon]; Foreman's Gloves (2167, -0.14 DPS) [world]; Gloves of the Fang (10413, -0.16 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.62 DPS) | yes | Ruffian Belt (5975, -0.21 DPS, sim-verified) [world]; Support Girdle (1215, -0.28 DPS) [world]; Brawler's Leather Belt (252428, -0.31 DPS) [crafted] |
| legs | Totemic Leather Pants (252446) | Leatherworking [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Deepgrave Trousers (279900, -0.11 DPS) [quest]; Brawler's Leather Pants (252500, -0.13 DPS) [crafted]; Defender's Leather Pants (252445, -0.50 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 11.2 attack_power points (0.39 DPS) | yes | Defender's Leather Boots (252441, -0.04 DPS) [crafted]; Totemic Leather Boots (252442, -0.04 DPS) [crafted]; Forest Leather Boots (3057, -0.11 DPS) [world_drop] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.9 attack_power points (0.31 DPS) | yes | The 1 Ring (8350, -0.23 DPS) [world]; Ring of the Moon (12052, -0.24 DPS) [world_drop]; Signet of the Zhevra (285330, -0.26 DPS) [world] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | The 1 Ring (8350, -0.20 DPS, sim-verified) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop]; Signet of the Zhevra (285330, -0.23 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | sim-verified (+3.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Living Root (6631, -0.07 DPS) [dungeon]; Night Reaver (1318, -0.36 DPS) [dungeon]; The Axe of Severing (23171, -3.09 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Lambent Scale Cloak; chest: Totemic Leather Armor; wrist: Bravo's Armbands; hands: Gold-flecked Gloves; waist: Blackened Defias Belt; legs: Totemic Leather Pants; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: Smite's Mighty Hammer

No-known-source sample (15 of 225, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher

### Band 30 (dwarf, 0000000000000000-253130030004000000-0000000000000000)

Set DPS (verified): 98.4. Weights run: 1.0s. Verify run: 0.9s. 366 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.432 ± 0.100, crit=0.525 ± 0.024 per rating point (14 rating = 1%, 7.348 per %), hit=0.619 ± 0.009 per rating point (10 rating = 1%, 6.193 per %), melee_haste=2.835 ± 0.417

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (0.85 DPS) | yes | Azure Gustwoven Hood (277050, -0.21 DPS) [crafted]; Defender's Leather Hood (252447, -0.28 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.50 DPS, sim-verified) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.49 DPS) | yes | Kaleidoscope Chain (13084, -0.15 DPS) [world_drop]; River Pride Choker (13087, -0.21 DPS) [world_drop]; Sentinel's Medallion (19541, -0.37 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 14.7 attack_power points (0.52 DPS) | yes | Bristlebark Amice (14573, -0.22 DPS) [world_drop]; Watchman Pauldrons (7727, -0.31 DPS) [dungeon]; Barbaric Shoulders (5964, -0.55 DPS, sim-verified) [crafted] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 attack_power points (0.35 DPS) | yes | Sergeant Major's Cape (16315, -0.01 DPS) [pvp]; Hawkeye's Cloak (14593, -0.03 DPS) [world_drop]; Slayer's Cape (14752, -0.07 DPS) [world_drop] |
| chest | Thick Murloc Armor (5782) (or Nightwalker Armor (2234)) | Leatherworking [crafted] | 18.0 attack_power points (0.64 DPS) | yes | Nightwalker Armor (2234, +0.00 DPS) [world]; Totemic Leather Tunic (252451, -0.07 DPS) [crafted]; Raptorbane Armor (3566, -0.07 DPS) [quest] |
| wrist | Bands of Serra'kis (6902) | Blackfathom Deeps: Old Serra'kis [dungeon] | 12.0 attack_power points (0.42 DPS) | yes | Hawkeye's Bracers (14590, -0.05 DPS) [world_drop]; Technician's Bracers (270042, -0.07 DPS) [quest]; Cultist's Armguards (270032, -0.07 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 16.6 attack_power points (0.59 DPS) | yes | Brawler Gloves (720, -0.02 DPS) [world_drop]; Heavy Earthen Gloves (7359, -0.02 DPS) [crafted]; Toughened Leather Gloves (4253, -0.07 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (0.85 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Prowler's Leather Belt (252459, -0.12 DPS) [crafted]; Skirmisher's Leather Belt (252461, -0.21 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (0.92 DPS) | yes | Totemic Leather Pants (252446, -0.28 DPS) [crafted]; Totemic Leather Leggings (252458, -0.28 DPS) [crafted]; Defender's Leather Pants (252445, -1.66 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 12.2 attack_power points (0.43 DPS) | yes | Disjointed Shoes (277226, -0.01 DPS) [quest]; Draftsman Boots (6668, -0.08 DPS) [quest]; Defender's Leather Boots (252441, -0.08 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 17.3 attack_power points (0.61 DPS) | yes | Tiger Band (6749, -0.19 DPS) [quest]; Ironspine's Eye (7686, -0.19 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.26 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 14.6 attack_power points (0.52 DPS) | yes | Ironspine's Eye (7686, -0.10 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.16 DPS) [dungeon]; Tiger Band (6749, -0.55 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (98.4 DPS) | yes | Corpsemaker (6687, -0.17 DPS) [dungeon]; Darkspear Raider's Reaper (272082, -0.62 DPS) [vendor]; Viscous Hammer (13045, -22.31 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Thick Murloc Armor; wrist: Bands of Serra'kis; hands: Insignia Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 366, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler

### Band 40 (dwarf, 0000000000000000-253130030005102051-0000000000000000)

Set DPS (verified): 100.1. Weights run: 1.2s. Verify run: 1.0s. 592 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.386 ± 0.051, crit=0.551 ± 0.015 per rating point (14 rating = 1%, 7.710 per %), hit=0.521 ± 0.008 per rating point (10 rating = 1%, 5.206 per %), melee_haste=3.244 ± 0.122

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 33.7 attack_power points (1.25 DPS) | yes | White Bandit Mask (10008, -0.28 DPS) [crafted]; Tusken Helm (6686, -0.29 DPS) [dungeon]; Hard Gold Coif (250537, -1.82 DPS, sim-verified) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.74 DPS) | yes | Ghostshard Talisman (7731, -0.22 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.39 DPS) [world_drop]; River Pride Choker (13087, -0.44 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 attack_power points (0.81 DPS) | yes | Imperial Leather Spaulders (4737, -0.15 DPS) [dungeon]; Wrangling Spaulders (15698, -0.19 DPS) [quest]; Sunburn Spaulders (274751, -0.21 DPS) [vendor] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 14.3 attack_power points (0.53 DPS) | yes | Wolfmaster Cape (6314, -0.16 DPS) [dungeon]; Hawkeye's Cloak (14593, -0.21 DPS) [world_drop]; Dark Hooded Cape (5257, -1.56 DPS, sim-verified) [world] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 30.7 attack_power points (1.14 DPS) | yes | Shining Silver Breastplate (2870, -0.10 DPS) [crafted]; Golden Scale Cuirass (3845, -0.10 DPS) [crafted]; Avenger's Armor (1488, -1.04 DPS, sim-verified) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.74 DPS) | yes | Pugilist Bracers (4438, -0.15 DPS) [dungeon]; Ravager's Armguards (14770, -0.17 DPS) [world_drop]; Yorgen Bracers (13012, -0.25 DPS) [world_drop] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 attack_power points (1.18 DPS) | yes | Scarlet Gauntlets (10331, -0.15 DPS) [dungeon]; Gloves of Holy Might (867, -0.16 DPS) [world_drop]; Reticulated Bone Gauntlets (9435, -0.22 DPS) [world_drop] |
| waist | Boar Champion's Belt (10768) | Razorfen Downs: Ragglesnout [dungeon] | 30.0 attack_power points (1.11 DPS) | yes | Highlander's Leather Girdle (20116, -0.00 DPS) [rep]; Girdle of Golem Strength (9405, -0.22 DPS) [world_drop]; Scarlet Belt (10329, -0.22 DPS) [dungeon] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.55 DPS) | yes | Firemane Leggings (13129, -0.15 DPS) [world_drop]; Orcish War Leggings (7929, -0.30 DPS) [crafted]; Legguards of the Vault (9396, -0.52 DPS) [dungeon] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 24.7 attack_power points (0.91 DPS) | yes | Blackforge Greaves (6423, -0.11 DPS) [dungeon]; Ironheel Boots (4653, -0.17 DPS) [quest]; Skirmisher's Mail Boots (252564, -0.83 DPS, sim-verified) [crafted] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.74 DPS) | yes | Protector's Band (19515, -0.03 DPS) [rep]; Thunderbrow Ring (13097, -0.11 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.74 DPS) | yes | Protector's Band (19515, -0.03 DPS) [rep]; Thunderbrow Ring (13097, -0.11 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | The Jackhammer (9423) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-verified (100.1 DPS) | yes | Bonebiter (6830, +0.00 DPS) [quest]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor]; Pendulum of Doom (9425, -6.35 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Boar Champion's Belt; legs: Scarlet Leggings; feet: Prowler's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: The Jackhammer

No-known-source sample (15 of 592, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 50 (dwarf, 5500000000000000-253130030005102051-0000000000000000)

Set DPS (verified): 124.1. Weights run: 1.2s. Verify run: 1.1s. 762 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.512 ± 0.061, crit=0.710 ± 0.019 per rating point (14 rating = 1%, 9.934 per %), hit=0.802 ± 0.012 per rating point (10 rating = 1%, 8.019 per %), melee_haste=3.814 ± 0.090

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 48.0 attack_power points (1.74 DPS) | yes | Raging Berserker's Helm (7719, -0.44 DPS) [dungeon]; Knight-Lieutenant's Mail Helmet (223075, -0.58 DPS) [vendor]; Bloomsprout Headpiece (17767, -1.10 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.72 DPS) | yes | Skibi's Pendant (13089, -0.12 DPS) [world_drop]; Ghostshard Talisman (7731, -0.22 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.36 DPS) [world_drop] |
| shoulder | Knight-Lieutenant's Mail Epaulets (223073) | Captain Dirgehammer [vendor] | 27.9 attack_power points (1.01 DPS) | yes | Prowler's Leather Shoulder (252534, -0.05 DPS) [crafted]; Failed Flying Experiment (9647, -0.10 DPS) [quest]; Skulker's Leather Shoulder (252535, -0.16 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 19.2 attack_power points (0.69 DPS) | yes | Sergeant Major's Cape (16336, -0.15 DPS) [pvp]; Dark Hooded Cape (5257, -0.22 DPS) [world]; Bloodlust Cape (14801, -1.42 DPS, sim-verified) [world_drop] |
| chest | Mixologist's Tunic (12793) | Blackrock Depths: Plugger Spazzring [dungeon] | 41.6 attack_power points (1.51 DPS) | yes | Knight's Mail Armor (223078, -0.28 DPS) [vendor]; Kolkar Marauder Chain (6773, -0.36 DPS) [quest]; Warbear Harness (15064, -0.38 DPS) [crafted] |
| wrist | Arena Bands (18711) (or Bracers of the Stone Princess (17714)) | Arena Treasure Chest [world] | 28.0 attack_power points (1.01 DPS) | yes | Bracers of the Stone Princess (17714, +0.00 DPS) [dungeon]; Prowler's Leather Bracers (252539, -0.23 DPS) [crafted]; Branded Leather Bracers (19508, -0.29 DPS) [dungeon] |
| hands | Prowler's Leather Gauntlets (252547) | Leatherworking [crafted] | 32.6 attack_power points (1.18 DPS) | yes | Maddening Gauntlets (11867, -0.00 DPS) [quest]; Gauntlets of Divinity (7724, -0.02 DPS) [dungeon]; Raider Gloves (272100, -0.05 DPS) [vendor] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (1.66 DPS) | yes | Belt of the Gladiator (13134, -0.36 DPS) [world_drop]; Prowler's Leather Waistguard (252473, -0.43 DPS) [crafted]; Skulker's Leather Waistguard (252474, -0.54 DPS) [crafted] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.52 DPS) | yes | Firemane Leggings (13129, -0.14 DPS) [world_drop]; Gryphon Rider's Leggings (9652, -0.16 DPS) [quest]; Serpentskin Leggings (8262, -0.28 DPS) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 31.6 attack_power points (1.14 DPS) | yes | Skulker's Leather Boots (252469, -0.11 DPS) [crafted]; Skirmisher's Mail Sabatons (252578, -0.20 DPS) [crafted]; Shadefiend Boots (11675, -0.22 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 28.0 attack_power points (1.01 DPS) | yes | Mark of Kern (2262, -0.29 DPS) [dungeon]; Assault Band (13095, -0.29 DPS) [world_drop]; Thunderbrow Ring (13097, -0.38 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 24.6 attack_power points (0.89 DPS) | yes | Mark of Kern (2262, -0.17 DPS) [dungeon]; Assault Band (13095, -0.17 DPS) [world_drop]; Thunderbrow Ring (13097, -0.26 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (124.1 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (124.1 DPS) | yes | Smoking Heart of the Mountain (11811, +0.00 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (124.1 DPS) | yes | Ragehammer (10626, +0.00 DPS) [dungeon]; Thorium Greatmace (250613, -0.47 DPS) [crafted]; Darkspear Raider's Reaper (272080, -0.82 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Embrace of the Lycan; shoulder: Knight-Lieutenant's Mail Epaulets; back: Blackveil Cape; chest: Mixologist's Tunic; wrist: Arena Bands; hands: Prowler's Leather Gauntlets; waist: Girdle of Beastial Fury; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 762, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 60 (dwarf, 5533220000000000-253130030005102051-0000000000000000)

Set DPS (verified): 184.5. Weights run: 1.2s. Verify run: 1.1s. 1751 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.519 ± 0.068, crit=0.737 ± 0.023 per rating point (14 rating = 1%, 10.313 per %), hit=0.977 ± 0.017 per rating point (10 rating = 1%, 9.771 per %), melee_haste=3.972 ± 0.142

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 48.0 attack_power points (1.73 DPS) | yes | Warbear Helm (252485, -0.17 DPS) [crafted]; Face of The Five Thunders (227021, -0.21 DPS) [vendor]; Blue Suede Hat (252482, -0.21 DPS) [crafted] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 34.3 attack_power points (1.24 DPS) | yes | Beads of Ogre Might (22150, -0.02 DPS) [quest]; Imperial Jewel (11933, -0.08 DPS) [dungeon]; Will of the Martyr (17044, -0.16 DPS) [quest] |
| shoulder | Black Dragonscale Shoulders (15051) | Leatherworking [crafted] | 40.0 attack_power points (1.44 DPS) | yes | Highlander's Leather Shoulders (20059, -0.02 DPS) [rep]; Highlander's Lizardhide Shoulders (20060, -0.14 DPS) [rep]; Golden Mantle of the Dawn (19058, -0.30 DPS) [crafted] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 37.8 attack_power points (1.36 DPS) | yes | Shroud of Domination (22337, -0.14 DPS) [dungeon]; Cape of the Black Baron (13340, -0.36 DPS) [dungeon]; Cloak of the Honor Guard (20073, -1.63 DPS, sim-verified) [rep] |
| chest | Obsidian Mail Tunic (22191) | Blacksmithing [crafted] | sim-verified (184.5 DPS) | yes | Timbermaw Tunic (252484, -0.02 DPS) [crafted]; Cadaverous Armor (14637, -0.22 DPS) [dungeon]; Tunic of Undead Slaying (23089, -7.19 DPS, sim-verified) [world] |
| wrist | Windtalker's Wristguards (19582) | Silverwing Sentinels [rep] | sim-verified (184.5 DPS) | yes | Tranquil Wristguards (279256, -0.22 DPS) [crafted]; Forest Stalker's Bracers (19587, -0.22 DPS) [rep]; Wristwraps of Undead Slaying (23093, -7.93 DPS, sim-verified) [world] |
| hands | Timbermaw Brawlers (19049) | Leatherworking [crafted] | 49.3 attack_power points (1.78 DPS) | yes | Studded Timbermaw Brawlers (227809, -0.15 DPS) [vendor]; Cadaverous Gloves (14640, -0.19 DPS) [dungeon]; Skul's Fingerbone Claws (13395, -0.33 DPS) [dungeon] |
| waist | Dense Timbermaw Belt (227807) | Meilosh [vendor] | 64.0 attack_power points (2.31 DPS) | yes | Ferocity of the Timbermaw (227805, -0.11 DPS) [vendor]; Might of the Timbermaw (19044, -0.53 DPS) [crafted]; Windseeker's Belt (272410, -0.58 DPS) [vendor] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 65.4 attack_power points (2.36 DPS) | yes | Devilsaur Leggings (15062, -0.33 DPS) [crafted]; Black Dragonscale Leggings (15052, -0.41 DPS) [crafted]; Cadaverous Leggings (14638, -0.48 DPS) [dungeon] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 42.2 attack_power points (1.52 DPS) | yes | Pads of the Dread Wolf (13210, -0.08 DPS) [dungeon]; Drudge Boots (21532, -0.26 DPS) [quest]; Boots of Ferocity (22472, -0.34 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (184.5 DPS) | yes | Protector's Band (19514, -0.23 DPS) [rep]; Band of the Ogre King (18522, -0.29 DPS) [dungeon]; Naglering (11669, -5.95 DPS, sim-verified) [dungeon] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | sim-verified (184.5 DPS) | yes | Protector's Band (19514, -0.00 DPS) [rep]; Band of the Ogre King (18522, -0.06 DPS) [dungeon]; Naglering (11669, -4.22 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (184.5 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-verified (184.5 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Eye of the Beast (13968, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-verified (184.5 DPS) | yes | Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Seeping Willow (12969, -13.04 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** neck: Medallion of the Dawn; shoulder: Black Dragonscale Shoulders; back: Howler's Furs; chest: Obsidian Mail Tunic; wrist: Windtalker's Wristguards; hands: Timbermaw Brawlers; waist: Dense Timbermaw Belt; legs: Warbear Woolies; feet: Scalegut Treaders; finger1: Don Julio's Band; finger2: Blackstone Ring; trinket1: Darkmoon Card: Maelstrom; trinket2: Blackhand's Breadth; main_hand: The Unstoppable Force

No-known-source sample (15 of 1751, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

## Horde

### Band 20 (orc, 0000000000000000-253100000000000000-0000000000000000)

Set DPS (verified): 45.7. Weights run: 1.0s. Verify run: 0.8s. 205 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=0.235 ± 0.025, crit=0.309 ± 0.007 per rating point (14 rating = 1%, 4.333 per %), hit=0.439 ± 0.005 per rating point (10 rating = 1%, 4.391 per %), melee_haste=2.440 ± 0.609

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (0.55 DPS) | yes | Brawler's Leather Hood (252504, -0.50 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 1.4 attack_power points (0.05 DPS) | yes | Erudite's Amulet (277204, -0.02 DPS) [quest] |
| shoulder | Slime-encrusted Pads (6461) | Wailing Caverns: Mutanus the Devourer [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Serpent's Shoulders (5404, -0.49 DPS, sim-verified) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Grave Shroud (279865, -0.05 DPS) [quest]; Catacomb Cloak (279899, -0.07 DPS) [quest]; Subterranean Cape (14149, -0.07 DPS) [dungeon] |
| chest | Totemic Leather Armor (252435) | Leatherworking [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Armor of the Fang (6473, -0.07 DPS) [dungeon]; Brawler's Leather Armor (252490, -0.08 DPS) [crafted]; Defender's Leather Armor (252434, -0.75 DPS, sim-verified) [crafted] |
| wrist | Bristlebark Bindings (14569) | World drop [world_drop] | 4.7 attack_power points (0.16 DPS) | yes | Light Leather Bracers (7281, -0.09 DPS) [crafted]; Cloudy Gustwoven Bracers (276999, -0.09 DPS) [crafted]; Azure Gustwoven Bracers (277023, -0.09 DPS) [crafted] |
| hands | Gold-flecked Gloves (5195) | The Deadmines: Sneed [dungeon] | 14.0 attack_power points (0.48 DPS) | yes | Blackened Defias Gloves (10401, -0.07 DPS) [dungeon]; Foreman's Gloves (2167, -0.14 DPS) [world]; Gloves of the Fang (10413, -0.16 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.62 DPS) | yes | Ruffian Belt (5975, -0.22 DPS, sim-verified) [world]; Support Girdle (1215, -0.28 DPS) [world]; Brawler's Leather Belt (252428, -0.31 DPS) [crafted] |
| legs | Totemic Leather Pants (252446) | Leatherworking [crafted] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Deepgrave Trousers (279900, -0.11 DPS) [quest]; Brawler's Leather Pants (252500, -0.13 DPS) [crafted]; Defender's Leather Pants (252445, -0.89 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 11.2 attack_power points (0.39 DPS) | yes | Defender's Leather Boots (252441, -0.04 DPS) [crafted]; Totemic Leather Boots (252442, -0.04 DPS) [crafted]; Forest Leather Boots (3057, -0.11 DPS) [world_drop] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.9 attack_power points (0.31 DPS) | yes | Loop of Sacrifice (281673, -0.10 DPS) [quest]; The 1 Ring (8350, -0.23 DPS) [world]; Ring of the Moon (12052, -0.24 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Loop of Sacrifice (281673, +0.00 DPS, sim-verified) [quest]; The 1 Ring (8350, -0.20 DPS) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | The Axe of Severing (23171) | Shadowfang Keep: Sever [dungeon] | 322.0 attack_power points (11.11 DPS) | yes | Forsaken Greataxe (251533, -0.66 DPS) [quest]; Smite's Mighty Hammer (7230, -0.80 DPS) [dungeon]; Hammerbone (270018, -4.29 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Scout's Medallion; shoulder: Slime-encrusted Pads; back: Lambent Scale Cloak; chest: Totemic Leather Armor; wrist: Bristlebark Bindings; hands: Gold-flecked Gloves; waist: Blackened Defias Belt; legs: Totemic Leather Pants; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: The Axe of Severing

No-known-source sample (15 of 205, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield; 14389 Durability Shoulderpads; 14705 Brackwater Chain Shield

### Band 30 (orc, 0000000000000000-253130030004000000-0000000000000000)

Set DPS (verified): 91.6. Weights run: 1.0s. Verify run: 0.9s. 349 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.432 ± 0.100, crit=0.525 ± 0.024 per rating point (14 rating = 1%, 7.348 per %), hit=0.619 ± 0.009 per rating point (10 rating = 1%, 6.193 per %), melee_haste=2.835 ± 0.417

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (0.85 DPS) | yes | Cloudy Gustwoven Hood (277042, -0.21 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.21 DPS) [crafted]; Defender's Leather Hood (252447, -0.28 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.49 DPS) | yes | Kaleidoscope Chain (13084, -0.15 DPS) [world_drop]; River Pride Choker (13087, -0.21 DPS) [world_drop]; Scout's Medallion (19537, -0.37 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 14.7 attack_power points (0.52 DPS) | yes | Barbaric Shoulders (5964, -0.09 DPS) [crafted]; Bristlebark Amice (14573, -0.22 DPS) [world_drop]; Watchman Pauldrons (7727, -0.31 DPS) [dungeon] |
| back | Wildhunter Cloak (16658) (or Wolfmaster Cape (6314)) | The Hunt Completed [quest] | 10.0 attack_power points (0.35 DPS) | yes | Wolfmaster Cape (6314, +0.00 DPS) [dungeon]; Hawkeye's Cloak (14593, -0.03 DPS) [world_drop]; Slayer's Cape (14752, -0.07 DPS) [world_drop] |
| chest | Thick Murloc Armor (5782) (or Nightwalker Armor (2234)) | Leatherworking [crafted] | 18.0 attack_power points (0.64 DPS) | yes | Nightwalker Armor (2234, +0.00 DPS) [world]; Totemic Leather Tunic (252451, -0.07 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.09 DPS) [crafted] |
| wrist | Bands of Serra'kis (6902) | Blackfathom Deeps: Old Serra'kis [dungeon] | 12.0 attack_power points (0.42 DPS) | yes | Hawkeye's Bracers (14590, -0.05 DPS) [world_drop]; Technician's Bracers (270042, -0.07 DPS) [quest]; Cultist's Armguards (270032, -0.07 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 16.6 attack_power points (0.59 DPS) | yes | Brawler Gloves (720, -0.02 DPS) [world_drop]; Heavy Earthen Gloves (7359, -0.02 DPS) [crafted]; Toughened Leather Gloves (4253, -0.07 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (0.85 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Prowler's Leather Belt (252459, -0.12 DPS) [crafted]; Skirmisher's Leather Belt (252461, -0.21 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (0.92 DPS) | yes | Totemic Leather Pants (252446, -0.28 DPS) [crafted]; Totemic Leather Leggings (252458, -0.28 DPS) [crafted]; Defender's Leather Pants (252445, -1.67 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 12.2 attack_power points (0.43 DPS) | yes | Draftsman Boots (6668, -0.08 DPS) [quest]; Defender's Leather Boots (252441, -0.08 DPS) [crafted]; Totemic Leather Boots (252442, -0.08 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 17.3 attack_power points (0.61 DPS) | yes | Tiger Band (6749, -0.19 DPS) [quest]; Ironspine's Eye (7686, -0.19 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.26 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 14.6 attack_power points (0.52 DPS) | yes | Tiger Band (6749, -0.09 DPS) [quest]; Ironspine's Eye (7686, -0.10 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.16 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (91.6 DPS) | yes | Corpsemaker (6687, -0.17 DPS) [dungeon]; Darkspear Raider's Reaper (272082, -0.62 DPS) [vendor]; Viscous Hammer (13045, -21.47 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wildhunter Cloak; chest: Thick Murloc Armor; wrist: Bands of Serra'kis; hands: Insignia Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 349, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots

### Band 40 (orc, 0000000000000000-253130030005102051-0000000000000000)

Set DPS (verified): 102.9. Weights run: 1.2s. Verify run: 0.9s. 555 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.386 ± 0.051, crit=0.551 ± 0.015 per rating point (14 rating = 1%, 7.710 per %), hit=0.521 ± 0.008 per rating point (10 rating = 1%, 5.206 per %), melee_haste=3.244 ± 0.122

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 33.7 attack_power points (1.25 DPS) | yes | Hard Gold Coif (250537, -0.21 DPS) [crafted]; White Bandit Mask (10008, -0.28 DPS) [crafted]; Tusken Helm (6686, -0.29 DPS) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.74 DPS) | yes | Ghostshard Talisman (7731, -0.22 DPS) [dungeon]; Ethereal Talisman (4430, -0.31 DPS) [quest]; Kaleidoscope Chain (13084, -0.39 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 attack_power points (0.81 DPS) | yes | Imperial Leather Spaulders (4737, -0.15 DPS) [dungeon]; Wrangling Spaulders (15698, -0.19 DPS) [quest]; Sunburn Spaulders (274751, -0.21 DPS) [vendor] |
| back | First Sergeant's Cloak (16340) | PvP rank 9 · First Sergeant · Horde [pvp] | 14.3 attack_power points (0.53 DPS) | yes | Dark Hooded Cape (5257, -0.09 DPS) [world]; Wolfmaster Cape (6314, -0.16 DPS) [dungeon]; Wildhunter Cloak (16658, -0.16 DPS) [quest] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 30.7 attack_power points (1.14 DPS) | yes | Avenger's Armor (1488, -0.03 DPS) [dungeon]; Shining Silver Breastplate (2870, -0.10 DPS) [crafted]; Golden Scale Cuirass (3845, -0.10 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.74 DPS) | yes | Windtalker's Wristguards (19584, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Pugilist Bracers (4438, -0.15 DPS) [dungeon] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 attack_power points (1.18 DPS) | yes | Scarlet Gauntlets (10331, -0.15 DPS) [dungeon]; Gloves of Holy Might (867, -0.16 DPS) [world_drop]; Reticulated Bone Gauntlets (9435, -0.22 DPS) [world_drop] |
| waist | Boar Champion's Belt (10768) | Razorfen Downs: Ragglesnout [dungeon] | 30.0 attack_power points (1.11 DPS) | yes | Defiler's Leather Girdle (20192, -0.00 DPS) [rep]; Tharg's Shoelace (9705, -0.15 DPS) [quest]; Scarlet Belt (10329, -0.22 DPS) [dungeon] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.55 DPS) | yes | Firemane Leggings (13129, -0.15 DPS) [world_drop]; Orcish War Leggings (7929, -0.30 DPS) [crafted]; Legguards of the Vault (9396, -0.52 DPS) [dungeon] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 24.7 attack_power points (0.91 DPS) | yes | Skirmisher's Mail Boots (252564, -0.10 DPS) [crafted]; Blackforge Greaves (6423, -0.11 DPS) [dungeon]; Skulker's Leather Shoes (252531, -0.24 DPS) [crafted] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.74 DPS) | yes | Legionnaire's Band (19512, -0.03 DPS) [rep]; Thunderbrow Ring (13097, -0.11 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.74 DPS) | yes | Legionnaire's Band (19512, -0.03 DPS) [rep]; Thunderbrow Ring (13097, -0.11 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pendulum of Doom (9425) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-verified (102.9 DPS) | yes | Fiery War Axe (870, +0.00 DPS) [world_drop]; Staff of Jordan (873, +0.00 DPS) [world_drop]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: First Sergeant's Cloak; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Boar Champion's Belt; legs: Scarlet Leggings; feet: Prowler's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Pendulum of Doom

No-known-source sample (15 of 555, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (orc, 5500000000000000-253130030005102051-0000000000000000)

Set DPS (verified): 130.5. Weights run: 1.2s. Verify run: 1.1s. 704 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.512 ± 0.061, crit=0.710 ± 0.019 per rating point (14 rating = 1%, 9.934 per %), hit=0.802 ± 0.012 per rating point (10 rating = 1%, 8.019 per %), melee_haste=3.814 ± 0.090

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 48.0 attack_power points (1.74 DPS) | yes | Bloomsprout Headpiece (17767, -0.43 DPS) [dungeon]; Raging Berserker's Helm (7719, -0.44 DPS) [dungeon]; Blood Guard's Mail Helmet (220820, -0.58 DPS) [vendor] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.72 DPS) | yes | Skibi's Pendant (13089, -0.12 DPS) [world_drop]; Woven Ivy Necklace (19159, -0.12 DPS) [quest]; Ghostshard Talisman (7731, -0.22 DPS) [dungeon] |
| shoulder | Prowler's Leather Shoulder (252534) | Leatherworking [crafted] | 26.6 attack_power points (0.96 DPS) | yes | Blood Guard's Mail Epaulets (220823, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.05 DPS) [quest]; Skulker's Leather Shoulder (252535, -0.11 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 19.2 attack_power points (0.69 DPS) | yes | First Sergeant's Cloak (16340, -0.15 DPS) [pvp]; Dark Hooded Cape (5257, -0.22 DPS) [world]; Bloodlust Cape (14801, -1.49 DPS, sim-verified) [world_drop] |
| chest | Stone Guard's Mail Armor (220826) | Lady Palanseer [vendor] | sim-verified (130.5 DPS) | yes | Kolkar Marauder Chain (6773, -0.09 DPS) [quest]; Warbear Harness (15064, -0.10 DPS) [crafted]; Mixologist's Tunic (12793, -1.66 DPS, sim-verified) [dungeon] |
| wrist | Arena Bands (18711) (or Bracers of the Stone Princess (17714)) | Arena Treasure Chest [world] | 28.0 attack_power points (1.01 DPS) | yes | Bracers of the Stone Princess (17714, +0.00 DPS) [dungeon]; Windtalker's Wristguards (19583, +0.00 DPS) [pvp] |
| hands | Prowler's Leather Gauntlets (252547) | Leatherworking [crafted] | 32.6 attack_power points (1.18 DPS) | yes | Gauntlets of Divinity (7724, -0.02 DPS) [dungeon]; Raider Gloves (272100, -0.05 DPS) [vendor]; Gloves of Holy Might (867, -0.10 DPS) [world_drop] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (1.66 DPS) | yes | Belt of the Gladiator (13134, -0.36 DPS) [world_drop]; Prowler's Leather Waistguard (252473, -0.43 DPS) [crafted]; Skulker's Leather Waistguard (252474, -0.54 DPS) [crafted] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.52 DPS) | yes | Firemane Leggings (13129, -0.14 DPS) [world_drop]; Serpentskin Leggings (8262, -0.28 DPS) [world_drop]; Orcish War Leggings (7929, -0.29 DPS) [crafted] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 31.6 attack_power points (1.14 DPS) | yes | Skulker's Leather Boots (252469, -0.11 DPS) [crafted]; Skirmisher's Mail Sabatons (252578, -0.20 DPS) [crafted]; Shadefiend Boots (11675, -0.22 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 28.0 attack_power points (1.01 DPS) | yes | White Bone Band (11862, -0.15 DPS) [quest]; Mark of Kern (2262, -0.29 DPS) [dungeon]; Assault Band (13095, -0.29 DPS) [world_drop] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 24.6 attack_power points (0.89 DPS) | yes | White Bone Band (11862, -0.02 DPS) [quest]; Mark of Kern (2262, -0.17 DPS) [dungeon]; Assault Band (13095, -0.17 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+4.4 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -1.25 DPS, sim-verified) [crafted] |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thorium Greatmace (250613, +0.00 DPS) [crafted]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor]; Glowing Brightwood Staff (812, -1.28 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Embrace of the Lycan; shoulder: Prowler's Leather Shoulder; back: Blackveil Cape; chest: Stone Guard's Mail Armor; wrist: Arena Bands; hands: Prowler's Leather Gauntlets; waist: Girdle of Beastial Fury; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Fiery War Axe

No-known-source sample (15 of 704, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (orc, 5533220000000000-253130030005102051-0000000000000000)

Set DPS (verified): 175.1. Weights run: 1.2s. Verify run: 1.0s. 1672 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.519 ± 0.068, crit=0.737 ± 0.023 per rating point (14 rating = 1%, 10.313 per %), hit=0.977 ± 0.017 per rating point (10 rating = 1%, 9.771 per %), melee_haste=3.972 ± 0.142

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 48.0 attack_power points (1.73 DPS) | yes | Champion's Mail Headguard (227155, +0.00 DPS) [pvp]; Warbear Helm (252485, -0.17 DPS) [crafted]; Face of The Five Thunders (227021, -0.21 DPS) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 34.3 attack_power points (1.24 DPS) | yes | Beads of Ogre Might (22150, -0.02 DPS) [quest]; Imperial Jewel (11933, -0.08 DPS) [dungeon]; Will of the Martyr (17044, -0.16 DPS) [quest] |
| shoulder | Black Dragonscale Shoulders (15051) | Leatherworking [crafted] | 40.0 attack_power points (1.44 DPS) | yes | Champion's Mail Pauldrons (227154, +0.00 DPS) [pvp]; Warlord's Mail Pauldrons (231654, +0.00 DPS) [vendor]; Defiler's Leather Shoulders (20194, -0.02 DPS) [rep] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 37.8 attack_power points (1.36 DPS) | yes | Deathguard's Cloak (20068, -0.04 DPS) [rep]; Shroud of Domination (22337, -0.14 DPS) [dungeon]; Cape of the Black Baron (13340, -0.36 DPS) [dungeon] |
| chest | Obsidian Mail Tunic (22191) | Blacksmithing [crafted] | sim-verified (175.1 DPS) | yes | Timbermaw Tunic (252484, -0.02 DPS) [crafted]; Cadaverous Armor (14637, -0.22 DPS) [dungeon]; Tunic of Undead Slaying (23089, -8.83 DPS, sim-verified) [world] |
| wrist | Windtalker's Wristguards (19582) | Warsong Outriders [rep] | sim-verified (175.1 DPS) | yes | Tranquil Wristguards (279256, -0.22 DPS) [crafted]; Forest Stalker's Bracers (19587, -0.22 DPS) [rep]; Wristwraps of Undead Slaying (23093, -3.99 DPS, sim-verified) [world] |
| hands | Timbermaw Brawlers (19049) | Leatherworking [crafted] | 49.3 attack_power points (1.78 DPS) | yes | General's Mail Vices (231655, +0.00 DPS) [vendor]; Studded Timbermaw Brawlers (227809, -0.15 DPS) [vendor]; Cadaverous Gloves (14640, -0.19 DPS) [dungeon] |
| waist | Dense Timbermaw Belt (227807) | Meilosh [vendor] | 64.0 attack_power points (2.31 DPS) | yes | Ferocity of the Timbermaw (227805, -0.11 DPS) [vendor]; Might of the Timbermaw (19044, -0.53 DPS) [crafted]; Windseeker's Belt (272410, -0.58 DPS) [vendor] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 65.4 attack_power points (2.36 DPS) | yes | General's Mail Legguards (231658, +0.00 DPS) [vendor]; Legionnaire's Mail Legguards (227156, -0.05 DPS) [pvp]; Devilsaur Leggings (15062, -0.33 DPS) [crafted] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 42.2 attack_power points (1.52 DPS) | yes | General's Mail Greaves (231656, +0.00 DPS) [vendor]; Pads of the Dread Wolf (13210, -0.08 DPS) [dungeon]; Blood Guard's Mail Greaves (227158, -0.16 DPS) [pvp] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (175.1 DPS) | yes | Legionnaire's Band (19510, -0.23 DPS) [rep]; Band of the Ogre King (18522, -0.29 DPS) [dungeon]; Naglering (11669, -4.27 DPS, sim-verified) [dungeon] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | sim-verified (175.1 DPS) | yes | Legionnaire's Band (19510, -0.00 DPS) [rep]; Band of the Ogre King (18522, -0.06 DPS) [dungeon]; Naglering (11669, -2.34 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (175.1 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (175.1 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Draconic Infused Emblem (22268, -2.96 DPS, sim-verified) [dungeon] |
| main_hand | Gravestone War Axe (13983) | Scholomance: Kirtonos the Herald [dungeon] | sim-verified (175.1 DPS) | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; The Unstoppable Force (19323, -5.81 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** neck: Medallion of the Dawn; shoulder: Black Dragonscale Shoulders; back: Howler's Furs; chest: Obsidian Mail Tunic; wrist: Windtalker's Wristguards; hands: Timbermaw Brawlers; waist: Dense Timbermaw Belt; legs: Warbear Woolies; feet: Scalegut Treaders; finger1: Don Julio's Band; finger2: Blackstone Ring; trinket1: Darkmoon Card: Maelstrom; trinket2: Burst of Knowledge; main_hand: Gravestone War Axe

No-known-source sample (15 of 1672, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

