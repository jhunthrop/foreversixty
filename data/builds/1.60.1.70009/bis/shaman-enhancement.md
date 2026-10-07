# Leveling BiS: Enhancement

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-253100000000000000-0000000000000000)

Set DPS (verified): 32.3. Weights run: 2.3s. Verify run: 1.1s. 225 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.180 ± 0.003, crit=0.253 ± 0.004 per rating point (14 rating = 1%, 3.543 per %), hit=0.206 ± 0.002 per rating point (10 rating = 1%, 2.064 per %), melee_haste=2.162 ± 0.027

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (0.55 DPS) | yes | Brawler's Leather Hood (252504, -0.43 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 1.1 attack_power points (0.04 DPS) | yes | Erudite's Amulet (277204, -0.01 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 0.9 attack_power points (0.03 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Grave Shroud (279865, -0.06 DPS) [quest]; Catacomb Cloak (279899, -0.07 DPS) [quest]; Dark Leather Cloak (2316, -0.12 DPS) [crafted] |
| chest | Defender's Leather Armor (252434) | Leatherworking [crafted] | 14.5 attack_power points (0.50 DPS) | yes | Totemic Leather Armor (252435, -0.02 DPS) [crafted]; Armor of the Fang (6473, -0.09 DPS) [dungeon]; Brawler's Leather Armor (252490, -0.11 DPS) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 4.7 attack_power points (0.16 DPS) | yes | Bristlebark Bindings (14569, -0.01 DPS) [world_drop]; Light Leather Bracers (7281, -0.09 DPS) [crafted]; Cloudy Gustwoven Bracers (276999, -0.09 DPS) [crafted] |
| hands | Gold-flecked Gloves (5195) | The Deadmines: Sneed [dungeon] | 14.0 attack_power points (0.48 DPS) | yes | Blackened Defias Gloves (10401, -0.07 DPS) [dungeon]; Foreman's Gloves (2167, -0.14 DPS) [world]; Gloves of the Fang (10413, -0.17 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.62 DPS) | yes | Ruffian Belt (5975, -0.21 DPS, sim-verified) [world]; Support Girdle (1215, -0.28 DPS) [world]; Brawler's Leather Belt (252428, -0.32 DPS) [crafted] |
| legs | Defender's Leather Pants (252445) | Leatherworking [crafted] | 18.7 attack_power points (0.65 DPS) | yes | Totemic Leather Pants (252446, -0.02 DPS) [crafted]; Deepgrave Trousers (279900, -0.14 DPS) [quest]; Brawler's Leather Pants (252500, -0.18 DPS) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.9 attack_power points (0.38 DPS) | yes | Defender's Leather Boots (252441, -0.03 DPS) [crafted]; Totemic Leather Boots (252442, -0.03 DPS) [crafted]; Forest Leather Boots (3057, -0.10 DPS) [world_drop] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.7 attack_power points (0.30 DPS) | yes | The 1 Ring (8350, -0.23 DPS) [world]; Ring of the Moon (12052, -0.23 DPS) [world_drop]; Signet of the Zhevra (285330, -0.26 DPS) [world] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | The 1 Ring (8350, -0.18 DPS, sim-verified) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop]; Signet of the Zhevra (285330, -0.24 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | The Axe of Severing (23171) | Shadowfang Keep: Sever [dungeon] | 322.0 attack_power points (11.11 DPS) | yes | Smite's Mighty Hammer (7230, -0.84 DPS, sim-verified) [dungeon]; Living Root (6631, -0.86 DPS) [dungeon]; Night Reaver (1318, -1.15 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Lambent Scale Cloak; chest: Defender's Leather Armor; wrist: Bravo's Armbands; hands: Gold-flecked Gloves; waist: Blackened Defias Belt; legs: Defender's Leather Pants; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: The Axe of Severing

No-known-source sample (15 of 225, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher

### Band 30 (dwarf, 0000000000000000-253130030004000000-0000000000000000)

Set DPS (verified): 81.1. Weights run: 2.4s. Verify run: 1.3s. 366 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.337 ± 0.006, crit=0.472 ± 0.008 per rating point (14 rating = 1%, 6.610 per %), hit=0.294 ± 0.003 per rating point (10 rating = 1%, 2.941 per %), melee_haste=2.980 ± 0.040

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (0.85 DPS) | yes | Cloudy Gustwoven Hood (277042, -0.21 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.21 DPS) [crafted]; Defender's Leather Hood (252447, -0.28 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.49 DPS) | yes | Kaleidoscope Chain (13084, -0.16 DPS) [world_drop]; River Pride Choker (13087, -0.21 DPS) [world_drop]; Sentinel's Medallion (19541, -0.40 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 13.7 attack_power points (0.48 DPS) | yes | Barbaric Shoulders (5964, -0.07 DPS) [crafted]; Bristlebark Amice (14573, -0.20 DPS) [world_drop]; Watchman Pauldrons (7727, -0.27 DPS) [dungeon] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 attack_power points (0.35 DPS) | yes | Sergeant Major's Cape (16315, +0.00 DPS, sim-verified) [pvp]; Hawkeye's Cloak (14593, -0.06 DPS) [world_drop]; Slayer's Cape (14752, -0.07 DPS) [world_drop] |
| chest | Thick Murloc Armor (5782) (or Nightwalker Armor (2234)) | Leatherworking [crafted] | 18.0 attack_power points (0.64 DPS) | yes | Nightwalker Armor (2234, +0.00 DPS) [world]; Raptorbane Armor (3566, -0.07 DPS) [quest]; Totemic Leather Tunic (252451, -0.07 DPS) [crafted] |
| wrist | Bands of Serra'kis (6902) | Blackfathom Deeps: Old Serra'kis [dungeon] | 12.0 attack_power points (0.42 DPS) | yes | Hawkeye's Bracers (14590, -0.07 DPS) [world_drop]; Cultist's Armguards (270032, -0.07 DPS) [quest]; Technician's Bracers (270042, -0.07 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 16.0 attack_power points (0.57 DPS) | yes | Heavy Earthen Gloves (7359, -0.00 DPS) [crafted]; Brawler Gloves (720, -0.00 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.07 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (0.85 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Prowler's Leather Belt (252459, -0.14 DPS) [crafted]; Blackened Defias Belt (10403, -0.21 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (0.92 DPS) | yes | Totemic Leather Pants (252446, -0.28 DPS) [crafted]; Totemic Leather Leggings (252458, -0.28 DPS) [crafted]; Defender's Leather Pants (252445, -1.74 DPS, sim-verified) [crafted] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 attack_power points (0.42 DPS) | yes | Brawler's Leather Boots (252439, -0.01 DPS) [crafted]; Draftsman Boots (6668, -0.07 DPS) [quest]; Defender's Leather Boots (252441, -0.07 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 17.0 attack_power points (0.60 DPS) | yes | Tiger Band (6749, -0.18 DPS) [quest]; Ironspine's Eye (7686, -0.21 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.25 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 14.0 attack_power points (0.50 DPS) | yes | Tiger Band (6749, -0.07 DPS) [quest]; Ironspine's Eye (7686, -0.11 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.14 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (81.1 DPS) | yes | Corpsemaker (6687, -0.15 DPS) [dungeon]; Darkspear Raider's Reaper (272082, -0.60 DPS) [vendor]; Viscous Hammer (13045, -21.04 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Thick Murloc Armor; wrist: Bands of Serra'kis; hands: Insignia Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 366, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler

### Band 40 (dwarf, 0000000000000000-253130030005102051-0000000000000000)

Set DPS (verified): 111.1. Weights run: 2.9s. Verify run: 1.3s. 592 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.393 ± 0.007, crit=0.553 ± 0.011 per rating point (14 rating = 1%, 7.743 per %), hit=0.317 ± 0.003 per rating point (10 rating = 1%, 3.172 per %), melee_haste=3.187 ± 0.113

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 33.7 attack_power points (1.39 DPS) | yes | White Bandit Mask (10008, -0.31 DPS) [crafted]; Tusken Helm (6686, -0.32 DPS) [dungeon]; Hard Gold Coif (250537, -1.68 DPS, sim-verified) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.82 DPS) | yes | Ghostshard Talisman (7731, -0.25 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.43 DPS) [world_drop]; River Pride Choker (13087, -0.49 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 attack_power points (0.91 DPS) | yes | Imperial Leather Spaulders (4737, -0.16 DPS) [dungeon]; Wrangling Spaulders (15698, -0.21 DPS) [quest]; Sunburn Spaulders (274751, -0.23 DPS) [vendor] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 14.4 attack_power points (0.59 DPS) | yes | Wolfmaster Cape (6314, -0.18 DPS) [dungeon]; Hawkeye's Cloak (14593, -0.23 DPS) [world_drop]; Dark Hooded Cape (5257, -2.79 DPS, sim-verified) [world] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 30.8 attack_power points (1.27 DPS) | yes | Avenger's Armor (1488, -0.03 DPS) [dungeon]; Shining Silver Breastplate (2870, -0.11 DPS) [crafted]; Golden Scale Cuirass (3845, -0.11 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.82 DPS) | yes | Pugilist Bracers (4438, -0.16 DPS) [dungeon]; Ravager's Armguards (14770, -0.18 DPS) [world_drop]; Yorgen Bracers (13012, -0.28 DPS) [world_drop] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 attack_power points (1.32 DPS) | yes | Scarlet Gauntlets (10331, -0.17 DPS) [dungeon]; Gloves of Holy Might (867, -0.18 DPS) [world_drop]; Reticulated Bone Gauntlets (9435, -0.25 DPS) [world_drop] |
| waist | Boar Champion's Belt (10768) | Razorfen Downs: Ragglesnout [dungeon] | 30.0 attack_power points (1.24 DPS) | yes | Highlander's Leather Girdle (20116, -0.00 DPS) [rep]; Girdle of Golem Strength (9405, -0.25 DPS) [world_drop]; Scarlet Belt (10329, -0.25 DPS) [dungeon] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.73 DPS) | yes | Firemane Leggings (13129, -0.16 DPS) [world_drop]; Orcish War Leggings (7929, -0.33 DPS) [crafted]; Legguards of the Vault (9396, -0.58 DPS) [dungeon] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 24.8 attack_power points (1.02 DPS) | yes | Skirmisher's Mail Boots (252564, -0.11 DPS) [crafted]; Blackforge Greaves (6423, -0.12 DPS) [dungeon]; Ironheel Boots (4653, -0.20 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.82 DPS) | yes | Protector's Band (19515, -0.04 DPS) [rep]; Thunderbrow Ring (13097, -0.12 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.82 DPS) | yes | Protector's Band (19515, -0.04 DPS) [rep]; Thunderbrow Ring (13097, -0.12 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pendulum of Doom (9425) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-verified (111.1 DPS) | yes | Bonebiter (6830, +0.00 DPS) [quest]; Ravager (7717, +0.00 DPS) [dungeon]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Boar Champion's Belt; legs: Scarlet Leggings; feet: Prowler's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Pendulum of Doom

No-known-source sample (15 of 592, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 50 (dwarf, 5500000000000000-253130030005102051-0000000000000000)

Set DPS (verified): 151.1. Weights run: 2.9s. Verify run: 1.5s. 762 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.462 ± 0.009, crit=0.649 ± 0.013 per rating point (14 rating = 1%, 9.091 per %), hit=0.392 ± 0.004 per rating point (10 rating = 1%, 3.920 per %), melee_haste=3.749 ± 0.150

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 48.0 attack_power points (2.02 DPS) | yes | Raging Berserker's Helm (7719, -0.54 DPS) [dungeon]; Knight-Lieutenant's Mail Helmet (223075, -0.71 DPS) [vendor]; Bloomsprout Headpiece (17767, -1.29 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.84 DPS) | yes | Skibi's Pendant (13089, -0.17 DPS) [world_drop]; Ghostshard Talisman (7731, -0.25 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.43 DPS) [world_drop] |
| shoulder | Knight-Lieutenant's Mail Epaulets (223073) | Captain Dirgehammer [vendor] | 27.1 attack_power points (1.14 DPS) | yes | Prowler's Leather Shoulder (252534, -0.04 DPS) [crafted]; Failed Flying Experiment (9647, -0.10 DPS) [quest]; Skulker's Leather Shoulder (252535, -0.17 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 18.5 attack_power points (0.78 DPS) | yes | Sergeant Major's Cape (16336, -0.16 DPS) [pvp]; Dark Hooded Cape (5257, -0.25 DPS) [world]; Bloodlust Cape (14801, -1.34 DPS, sim-verified) [world_drop] |
| chest | Knight's Mail Armor (223078) | Captain Dirgehammer [vendor] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Kolkar Marauder Chain (6773, -0.08 DPS) [quest]; Warbear Harness (15064, -0.12 DPS) [crafted]; Mixologist's Tunic (12793, -1.89 DPS, sim-verified) [dungeon] |
| wrist | Arena Bands (18711) (or Bracers of the Stone Princess (17714)) | Arena Treasure Chest [world] | 28.0 attack_power points (1.18 DPS) | yes | Bracers of the Stone Princess (17714, +0.00 DPS) [dungeon]; Prowler's Leather Bracers (252539, -0.29 DPS) [crafted]; Branded Leather Bracers (19508, -0.34 DPS) [dungeon] |
| hands | Prowler's Leather Gauntlets (252547) | Leatherworking [crafted] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Gauntlets of Divinity (7724, -0.01 DPS) [dungeon]; Raider Gloves (272100, -0.08 DPS) [vendor]; Maddening Gauntlets (11867, -1.65 DPS, sim-verified) [quest] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (1.94 DPS) | yes | Belt of the Gladiator (13134, -0.42 DPS) [world_drop]; Prowler's Leather Waistguard (252473, -0.52 DPS) [crafted]; Skulker's Leather Waistguard (252474, -0.65 DPS) [crafted] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.77 DPS) | yes | Firemane Leggings (13129, -0.17 DPS) [world_drop]; Gryphon Rider's Leggings (9652, -0.21 DPS) [quest]; Orcish War Leggings (7929, -0.34 DPS) [crafted] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 31.1 attack_power points (1.31 DPS) | yes | Skulker's Leather Boots (252469, -0.13 DPS) [crafted]; Skirmisher's Mail Sabatons (252578, -0.21 DPS) [crafted]; Prowler's Leather Shoes (252465, -0.25 DPS) [crafted] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 24.2 attack_power points (1.02 DPS) | yes | Mark of Kern (2262, -0.18 DPS) [dungeon]; Assault Band (13095, -0.18 DPS) [world_drop]; Thunderbrow Ring (13097, -0.29 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 23.9 attack_power points (1.01 DPS) | yes | Assault Band (13095, -0.17 DPS) [world_drop]; Thunderbrow Ring (13097, -0.28 DPS) [world_drop]; Mark of Kern (2262, -2.93 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thorium Greatmace (250613, -0.55 DPS) [crafted]; Darkspear Raider's Reaper (272080, -0.96 DPS) [vendor]; Ragehammer (10626, -10.72 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Embrace of the Lycan; shoulder: Knight-Lieutenant's Mail Epaulets; back: Blackveil Cape; chest: Knight's Mail Armor; wrist: Arena Bands; hands: Prowler's Leather Gauntlets; waist: Girdle of Beastial Fury; feet: Prowler's Leather Boots; finger1: Protector's Band; finger2: Blackstone Ring; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 762, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 60 (dwarf, 5533220000000000-253130030005102051-0000000000000000)

Set DPS (verified): 203.3. Weights run: 2.7s. Verify run: 1.4s. 1751 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.547 ± 0.011, crit=0.771 ± 0.015 per rating point (14 rating = 1%, 10.797 per %), hit=0.532 ± 0.004 per rating point (10 rating = 1%, 5.325 per %), melee_haste=4.412 ± 0.175

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 48.0 attack_power points (1.98 DPS) | yes | Warbear Helm (252485, -0.18 DPS) [crafted]; Face of The Five Thunders (227021, -0.22 DPS) [vendor]; Blue Suede Hat (252482, -0.22 DPS) [crafted] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 34.8 attack_power points (1.44 DPS) | yes | Imperial Jewel (11933, -0.12 DPS) [dungeon]; Amulet of the Darkmoon (19491, -0.18 DPS) [quest]; Will of the Martyr (17044, -0.20 DPS) [quest] |
| shoulder | Black Dragonscale Shoulders (15051) | Leatherworking [crafted] | 40.0 attack_power points (1.65 DPS) | yes | Highlander's Leather Shoulders (20059, +0.00 DPS, sim-verified) [rep]; Highlander's Lizardhide Shoulders (20060, -0.14 DPS) [rep]; Golden Mantle of the Dawn (19058, -0.33 DPS) [crafted] |
| back | Cloak of the Honor Guard (20073) | The League of Arathor [rep] | 36.7 attack_power points (1.52 DPS) | yes | Shroud of Domination (22337, -0.11 DPS) [dungeon]; Howler's Furs (272414, -0.14 DPS) [vendor]; Cape of the Black Baron (13340, -0.35 DPS) [dungeon] |
| chest | Obsidian Mail Tunic (22191) | Blacksmithing [crafted] | sim-verified (203.3 DPS) | yes | Timbermaw Tunic (252484, -0.01 DPS) [crafted]; Cadaverous Armor (14637, -0.27 DPS) [dungeon]; Tunic of Undead Slaying (23089, -8.76 DPS, sim-verified) [world] |
| wrist | Windtalker's Wristguards (19582) | Silverwing Sentinels [rep] | sim-verified (203.3 DPS) | yes | Forest Stalker's Bracers (19587, -0.23 DPS) [rep]; Tranquil Wristguards (279256, -0.25 DPS) [crafted]; Wristwraps of Undead Slaying (23093, -5.97 DPS, sim-verified) [world] |
| hands | Timbermaw Brawlers (19049) | Leatherworking [crafted] | 49.7 attack_power points (2.05 DPS) | yes | Studded Timbermaw Brawlers (227809, -0.17 DPS) [vendor]; Cadaverous Gloves (14640, -0.23 DPS) [dungeon]; Skul's Fingerbone Claws (13395, -0.40 DPS) [dungeon] |
| waist | Dense Timbermaw Belt (227807) | Meilosh [vendor] | 64.0 attack_power points (2.65 DPS) | yes | Ferocity of the Timbermaw (227805, +0.00 DPS, sim-verified) [vendor]; Might of the Timbermaw (19044, -0.59 DPS) [crafted]; Windseeker's Belt (272410, -0.66 DPS) [vendor] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 65.9 attack_power points (2.72 DPS) | yes | Devilsaur Leggings (15062, -0.37 DPS) [crafted]; Black Dragonscale Leggings (15052, -0.49 DPS) [crafted]; Cadaverous Leggings (14638, -0.57 DPS) [dungeon] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 42.4 attack_power points (1.75 DPS) | yes | Pads of the Dread Wolf (13210, -0.10 DPS) [dungeon]; Drudge Boots (21532, -0.29 DPS) [quest]; Boots of Ferocity (22472, -0.38 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (203.3 DPS) | yes | Band of the Ogre King (18522, -0.17 DPS) [dungeon]; Blackstone Ring (17713, -0.28 DPS) [dungeon]; Naglering (11669, -3.21 DPS, sim-verified) [dungeon] |
| finger2 | Protector's Band (19514) | Silverwing Sentinels [rep] | sim-verified (203.3 DPS) | yes | Band of the Ogre King (18522, -0.08 DPS) [dungeon]; Blackstone Ring (17713, -0.19 DPS) [dungeon]; Naglering (11669, -3.07 DPS, sim-verified) [dungeon] |
| trinket1 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-verified (203.3 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Eye of the Beast (13968) | General Drakkisath's Demise [quest] | sim-verified (203.3 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Second Wind (11819, -1.88 DPS, sim-verified) [dungeon] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-verified (203.3 DPS) | yes | Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Gravestone War Axe (13983, -14.75 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** neck: Medallion of the Dawn; shoulder: Black Dragonscale Shoulders; back: Cloak of the Honor Guard; chest: Obsidian Mail Tunic; wrist: Windtalker's Wristguards; hands: Timbermaw Brawlers; waist: Dense Timbermaw Belt; legs: Warbear Woolies; feet: Scalegut Treaders; finger1: Don Julio's Band; finger2: Protector's Band; trinket1: Blackhand's Breadth; trinket2: Eye of the Beast; main_hand: The Unstoppable Force

No-known-source sample (15 of 1751, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

## Horde

### Band 20 (orc, 0000000000000000-253100000000000000-0000000000000000)

Set DPS (verified): 32.5. Weights run: 2.3s. Verify run: 1.0s. 205 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.180 ± 0.003, crit=0.253 ± 0.004 per rating point (14 rating = 1%, 3.543 per %), hit=0.206 ± 0.002 per rating point (10 rating = 1%, 2.064 per %), melee_haste=2.162 ± 0.027

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (0.55 DPS) | yes | Brawler's Leather Hood (252504, -0.40 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 1.1 attack_power points (0.04 DPS) | yes | Erudite's Amulet (277204, -0.01 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 0.9 attack_power points (0.03 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Grave Shroud (279865, -0.06 DPS) [quest]; Catacomb Cloak (279899, -0.07 DPS) [quest]; Subterranean Cape (14149, -0.07 DPS) [dungeon] |
| chest | Defender's Leather Armor (252434) | Leatherworking [crafted] | 14.5 attack_power points (0.50 DPS) | yes | Totemic Leather Armor (252435, -0.02 DPS) [crafted]; Armor of the Fang (6473, -0.09 DPS) [dungeon]; Brawler's Leather Armor (252490, -0.11 DPS) [crafted] |
| wrist | Bristlebark Bindings (14569) | World drop [world_drop] | 4.5 attack_power points (0.16 DPS) | yes | Light Leather Bracers (7281, -0.08 DPS) [crafted]; Cloudy Gustwoven Bracers (276999, -0.09 DPS) [crafted]; Azure Gustwoven Bracers (277023, -0.09 DPS) [crafted] |
| hands | Gold-flecked Gloves (5195) | The Deadmines: Sneed [dungeon] | 14.0 attack_power points (0.48 DPS) | yes | Blackened Defias Gloves (10401, -0.07 DPS) [dungeon]; Foreman's Gloves (2167, -0.14 DPS) [world]; Gloves of the Fang (10413, -0.17 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.62 DPS) | yes | Ruffian Belt (5975, -0.21 DPS, sim-verified) [world]; Support Girdle (1215, -0.28 DPS) [world]; Brawler's Leather Belt (252428, -0.32 DPS) [crafted] |
| legs | Defender's Leather Pants (252445) | Leatherworking [crafted] | 18.7 attack_power points (0.65 DPS) | yes | Totemic Leather Pants (252446, -0.02 DPS) [crafted]; Deepgrave Trousers (279900, -0.14 DPS) [quest]; Brawler's Leather Pants (252500, -0.18 DPS) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.9 attack_power points (0.38 DPS) | yes | Defender's Leather Boots (252441, -0.03 DPS) [crafted]; Totemic Leather Boots (252442, -0.03 DPS) [crafted]; Forest Leather Boots (3057, -0.10 DPS) [world_drop] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.7 attack_power points (0.30 DPS) | yes | Loop of Sacrifice (281673, -0.09 DPS) [quest]; The 1 Ring (8350, -0.23 DPS) [world]; Ring of the Moon (12052, -0.23 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Loop of Sacrifice (281673, -0.07 DPS) [quest]; The 1 Ring (8350, -0.20 DPS) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | The Axe of Severing (23171) | Shadowfang Keep: Sever [dungeon] | 322.0 attack_power points (11.11 DPS) | yes | Forsaken Greataxe (251533, -0.66 DPS) [quest]; Hammerbone (270018, -0.74 DPS, sim-verified) [quest]; Smite's Mighty Hammer (7230, -0.81 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Lambent Scale Cloak; chest: Defender's Leather Armor; wrist: Bristlebark Bindings; hands: Gold-flecked Gloves; waist: Blackened Defias Belt; legs: Defender's Leather Pants; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: The Axe of Severing

No-known-source sample (15 of 205, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield; 14389 Durability Shoulderpads; 14705 Brackwater Chain Shield

### Band 30 (orc, 0000000000000000-253130030004000000-0000000000000000)

Set DPS (verified): 81.9. Weights run: 2.4s. Verify run: 1.2s. 349 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.337 ± 0.006, crit=0.472 ± 0.008 per rating point (14 rating = 1%, 6.610 per %), hit=0.294 ± 0.003 per rating point (10 rating = 1%, 2.941 per %), melee_haste=2.980 ± 0.040

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (0.85 DPS) | yes | Cloudy Gustwoven Hood (277042, -0.21 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.21 DPS) [crafted]; Defender's Leather Hood (252447, -0.28 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.49 DPS) | yes | Kaleidoscope Chain (13084, -0.16 DPS) [world_drop]; River Pride Choker (13087, -0.21 DPS) [world_drop]; Scout's Medallion (19537, -0.40 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 13.7 attack_power points (0.48 DPS) | yes | Barbaric Shoulders (5964, -0.07 DPS) [crafted]; Bristlebark Amice (14573, -0.20 DPS) [world_drop]; Watchman Pauldrons (7727, -0.27 DPS) [dungeon] |
| back | Wildhunter Cloak (16658) (or Wolfmaster Cape (6314)) | The Hunt Completed [quest] | 10.0 attack_power points (0.35 DPS) | yes | Wolfmaster Cape (6314, +0.00 DPS) [dungeon]; Hawkeye's Cloak (14593, -0.06 DPS) [world_drop]; Slayer's Cape (14752, -0.07 DPS) [world_drop] |
| chest | Thick Murloc Armor (5782) (or Nightwalker Armor (2234)) | Leatherworking [crafted] | 18.0 attack_power points (0.64 DPS) | yes | Nightwalker Armor (2234, +0.00 DPS) [world]; Totemic Leather Tunic (252451, -0.07 DPS) [crafted]; Defender's Leather Armor (252434, -0.11 DPS) [crafted] |
| wrist | Bands of Serra'kis (6902) | Blackfathom Deeps: Old Serra'kis [dungeon] | 12.0 attack_power points (0.42 DPS) | yes | Hawkeye's Bracers (14590, -0.07 DPS) [world_drop]; Cultist's Armguards (270032, -0.07 DPS) [quest]; Technician's Bracers (270042, -0.07 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 16.0 attack_power points (0.57 DPS) | yes | Heavy Earthen Gloves (7359, -0.00 DPS) [crafted]; Brawler Gloves (720, -0.00 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.07 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (0.85 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Prowler's Leather Belt (252459, -0.14 DPS) [crafted]; Blackened Defias Belt (10403, -0.21 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (0.92 DPS) | yes | Totemic Leather Pants (252446, -0.28 DPS) [crafted]; Totemic Leather Leggings (252458, -0.28 DPS) [crafted]; Defender's Leather Pants (252445, -1.68 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 11.7 attack_power points (0.41 DPS) | yes | Draftsman Boots (6668, -0.06 DPS) [quest]; Defender's Leather Boots (252441, -0.06 DPS) [crafted]; Totemic Leather Boots (252442, -0.06 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 17.0 attack_power points (0.60 DPS) | yes | Tiger Band (6749, -0.18 DPS) [quest]; Ironspine's Eye (7686, -0.21 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.25 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 14.0 attack_power points (0.50 DPS) | yes | Tiger Band (6749, -0.07 DPS) [quest]; Ironspine's Eye (7686, -0.11 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.14 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (81.9 DPS) | yes | Corpsemaker (6687, -0.15 DPS) [dungeon]; Darkspear Raider's Reaper (272082, -0.60 DPS) [vendor]; Viscous Hammer (13045, -21.05 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wildhunter Cloak; chest: Thick Murloc Armor; wrist: Bands of Serra'kis; hands: Insignia Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 349, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots

### Band 40 (orc, 0000000000000000-253130030005102051-0000000000000000)

Set DPS (verified): 112.2. Weights run: 2.9s. Verify run: 1.1s. 555 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.393 ± 0.007, crit=0.553 ± 0.011 per rating point (14 rating = 1%, 7.743 per %), hit=0.317 ± 0.003 per rating point (10 rating = 1%, 3.172 per %), melee_haste=3.187 ± 0.113

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 33.7 attack_power points (1.39 DPS) | yes | White Bandit Mask (10008, -0.31 DPS) [crafted]; Tusken Helm (6686, -0.32 DPS) [dungeon]; Hard Gold Coif (250537, -2.34 DPS, sim-verified) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.82 DPS) | yes | Ghostshard Talisman (7731, -0.25 DPS) [dungeon]; Ethereal Talisman (4430, -0.35 DPS) [quest]; Kaleidoscope Chain (13084, -0.43 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 attack_power points (0.91 DPS) | yes | Imperial Leather Spaulders (4737, -0.16 DPS) [dungeon]; Wrangling Spaulders (15698, -0.21 DPS) [quest]; Sunburn Spaulders (274751, -0.23 DPS) [vendor] |
| back | First Sergeant's Cloak (16340) | PvP rank 9 · First Sergeant · Horde [pvp] | 14.4 attack_power points (0.59 DPS) | yes | Wolfmaster Cape (6314, -0.18 DPS) [dungeon]; Wildhunter Cloak (16658, -0.18 DPS) [quest]; Dark Hooded Cape (5257, -2.19 DPS, sim-verified) [world] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 30.8 attack_power points (1.27 DPS) | yes | Avenger's Armor (1488, -0.03 DPS) [dungeon]; Shining Silver Breastplate (2870, -0.11 DPS) [crafted]; Golden Scale Cuirass (3845, -0.11 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.82 DPS) | yes | Windtalker's Wristguards (19584, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Pugilist Bracers (4438, -0.16 DPS) [dungeon] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 attack_power points (1.32 DPS) | yes | Scarlet Gauntlets (10331, -0.17 DPS) [dungeon]; Gloves of Holy Might (867, -0.18 DPS) [world_drop]; Reticulated Bone Gauntlets (9435, -0.25 DPS) [world_drop] |
| waist | Boar Champion's Belt (10768) | Razorfen Downs: Ragglesnout [dungeon] | 30.0 attack_power points (1.24 DPS) | yes | Defiler's Leather Girdle (20192, -0.00 DPS) [rep]; Tharg's Shoelace (9705, -0.16 DPS) [quest]; Scarlet Belt (10329, -0.25 DPS) [dungeon] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.73 DPS) | yes | Firemane Leggings (13129, -0.16 DPS) [world_drop]; Orcish War Leggings (7929, -0.33 DPS) [crafted]; Legguards of the Vault (9396, -0.58 DPS) [dungeon] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 24.8 attack_power points (1.02 DPS) | yes | Skirmisher's Mail Boots (252564, -0.11 DPS) [crafted]; Blackforge Greaves (6423, -0.12 DPS) [dungeon]; Skulker's Leather Shoes (252531, -0.26 DPS) [crafted] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.82 DPS) | yes | Legionnaire's Band (19512, -0.04 DPS) [rep]; Thunderbrow Ring (13097, -0.12 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.82 DPS) | yes | Legionnaire's Band (19512, -0.04 DPS) [rep]; Thunderbrow Ring (13097, -0.12 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ravager (7717) | Scarlet Monastery: Herod [dungeon] | sim-verified (112.2 DPS) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor]; Pendulum of Doom (9425, -2.86 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: First Sergeant's Cloak; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Boar Champion's Belt; legs: Scarlet Leggings; feet: Prowler's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Ravager

No-known-source sample (15 of 555, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (orc, 5500000000000000-253130030005102051-0000000000000000)

Set DPS (verified): 151.1. Weights run: 2.9s. Verify run: 1.4s. 704 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.462 ± 0.009, crit=0.649 ± 0.013 per rating point (14 rating = 1%, 9.091 per %), hit=0.392 ± 0.004 per rating point (10 rating = 1%, 3.920 per %), melee_haste=3.749 ± 0.150

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 48.0 attack_power points (2.02 DPS) | yes | Raging Berserker's Helm (7719, -0.54 DPS) [dungeon]; Blood Guard's Mail Helmet (220820, -0.71 DPS) [vendor]; Bloomsprout Headpiece (17767, -1.36 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.84 DPS) | yes | Woven Ivy Necklace (19159, -0.16 DPS) [quest]; Skibi's Pendant (13089, -0.17 DPS) [world_drop]; Ghostshard Talisman (7731, -0.25 DPS) [dungeon] |
| shoulder | Prowler's Leather Shoulder (252534) | Leatherworking [crafted] | 26.2 attack_power points (1.10 DPS) | yes | Blood Guard's Mail Epaulets (220823, +0.00 DPS) [vendor]; Skulker's Leather Shoulder (252535, -0.13 DPS) [crafted]; Failed Flying Experiment (9647, -1.60 DPS, sim-verified) [quest] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 18.5 attack_power points (0.78 DPS) | yes | First Sergeant's Cloak (16340, -0.16 DPS) [pvp]; Dark Hooded Cape (5257, -0.25 DPS) [world]; Bloodlust Cape (14801, -2.13 DPS, sim-verified) [world_drop] |
| chest | Mixologist's Tunic (12793) | Blackrock Depths: Plugger Spazzring [dungeon] | 41.1 attack_power points (1.73 DPS) | yes | Stone Guard's Mail Armor (220826, -0.34 DPS) [vendor]; Kolkar Marauder Chain (6773, -0.42 DPS) [quest]; Warbear Harness (15064, -0.45 DPS) [crafted] |
| wrist | Arena Bands (18711) (or Bracers of the Stone Princess (17714)) | Arena Treasure Chest [world] | 28.0 attack_power points (1.18 DPS) | yes | Bracers of the Stone Princess (17714, +0.00 DPS) [dungeon]; Windtalker's Wristguards (19583, +0.00 DPS) [pvp] |
| hands | Prowler's Leather Gauntlets (252547) | Leatherworking [crafted] | 32.2 attack_power points (1.36 DPS) | yes | Raider Gloves (272100, -0.08 DPS) [vendor]; Gloves of Holy Might (867, -0.13 DPS) [world_drop]; Gauntlets of Divinity (7724, -2.89 DPS, sim-verified) [dungeon] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (1.94 DPS) | yes | Belt of the Gladiator (13134, -0.42 DPS) [world_drop]; Prowler's Leather Waistguard (252473, -0.52 DPS) [crafted]; Skulker's Leather Waistguard (252474, -0.65 DPS) [crafted] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.77 DPS) | yes | Firemane Leggings (13129, -0.17 DPS) [world_drop]; Orcish War Leggings (7929, -0.34 DPS) [crafted]; Serpentskin Leggings (8262, -0.36 DPS) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 31.1 attack_power points (1.31 DPS) | yes | Skulker's Leather Boots (252469, -0.13 DPS) [crafted]; Skirmisher's Mail Sabatons (252578, -0.21 DPS) [crafted]; Prowler's Leather Shoes (252465, -0.25 DPS) [crafted] |
| finger1 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 24.2 attack_power points (1.02 DPS) | yes | Blackstone Ring (17713, -0.01 DPS) [dungeon]; Mark of Kern (2262, -0.18 DPS) [dungeon]; Assault Band (13095, -0.18 DPS) [world_drop] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.01 DPS) | yes | Mark of Kern (2262, -0.17 DPS) [dungeon]; Assault Band (13095, -0.17 DPS) [world_drop]; Blackstone Ring (17713, -1.30 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (151.1 DPS) | yes | Frozen Heart of the Mountain (249469, -1.74 DPS) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (151.1 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (151.1 DPS) | yes | Thorium Greatmace (250613, -0.55 DPS) [crafted]; Darkspear Raider's Reaper (272080, -0.96 DPS) [vendor]; Ragehammer (10626, -9.65 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Embrace of the Lycan; shoulder: Prowler's Leather Shoulder; back: Blackveil Cape; chest: Mixologist's Tunic; wrist: Arena Bands; hands: Prowler's Leather Gauntlets; waist: Girdle of Beastial Fury; feet: Prowler's Leather Boots; finger1: Legionnaire's Band; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Molten Heart of the Mountain; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 704, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (orc, 5533220000000000-253130030005102051-0000000000000000)

Set DPS (verified): 207.4. Weights run: 2.7s. Verify run: 1.3s. 1672 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.547 ± 0.011, crit=0.771 ± 0.015 per rating point (14 rating = 1%, 10.797 per %), hit=0.532 ± 0.004 per rating point (10 rating = 1%, 5.325 per %), melee_haste=4.412 ± 0.175

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 48.0 attack_power points (1.98 DPS) | yes | Champion's Mail Headguard (227155, +0.00 DPS) [pvp]; Warbear Helm (252485, -0.18 DPS) [crafted]; Face of The Five Thunders (227021, -0.22 DPS) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 34.8 attack_power points (1.44 DPS) | yes | Imperial Jewel (11933, -0.12 DPS) [dungeon]; Amulet of the Darkmoon (19491, -0.18 DPS) [quest]; Will of the Martyr (17044, -0.20 DPS) [quest] |
| shoulder | Black Dragonscale Shoulders (15051) | Leatherworking [crafted] | 40.0 attack_power points (1.65 DPS) | yes | Champion's Mail Pauldrons (227154, +0.00 DPS) [pvp]; Warlord's Mail Pauldrons (231654, +0.00 DPS) [vendor]; Defiler's Leather Shoulders (20194, -0.01 DPS) [rep] |
| back | Deathguard's Cloak (20068) | The Defilers [rep] | 36.7 attack_power points (1.52 DPS) | yes | Shroud of Domination (22337, -0.11 DPS) [dungeon]; Howler's Furs (272414, -0.14 DPS) [vendor]; Cape of the Black Baron (13340, -0.35 DPS) [dungeon] |
| chest | Obsidian Mail Tunic (22191) | Blacksmithing [crafted] | sim-verified (207.4 DPS) | yes | Timbermaw Tunic (252484, -0.01 DPS) [crafted]; Cadaverous Armor (14637, -0.27 DPS) [dungeon]; Tunic of Undead Slaying (23089, -8.83 DPS, sim-verified) [world] |
| wrist | Windtalker's Wristguards (19582) | Warsong Outriders [rep] | sim-verified (207.4 DPS) | yes | Forest Stalker's Bracers (19587, -0.23 DPS) [rep]; Tranquil Wristguards (279256, -0.25 DPS) [crafted]; Wristwraps of Undead Slaying (23093, -6.09 DPS, sim-verified) [world] |
| hands | Timbermaw Brawlers (19049) | Leatherworking [crafted] | 49.7 attack_power points (2.05 DPS) | yes | General's Mail Vices (231655, -0.06 DPS) [vendor]; Studded Timbermaw Brawlers (227809, -0.17 DPS) [vendor]; Cadaverous Gloves (14640, -0.23 DPS) [dungeon] |
| waist | Dense Timbermaw Belt (227807) | Meilosh [vendor] | 64.0 attack_power points (2.65 DPS) | yes | Ferocity of the Timbermaw (227805, -0.11 DPS) [vendor]; Might of the Timbermaw (19044, -0.59 DPS) [crafted]; Windseeker's Belt (272410, -0.66 DPS) [vendor] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 65.9 attack_power points (2.72 DPS) | yes | General's Mail Legguards (231658, +0.00 DPS) [vendor]; Legionnaire's Mail Legguards (227156, -0.24 DPS) [pvp]; Devilsaur Leggings (15062, -0.37 DPS) [crafted] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 42.4 attack_power points (1.75 DPS) | yes | Pads of the Dread Wolf (13210, -0.10 DPS) [dungeon]; General's Mail Greaves (231656, -0.13 DPS) [vendor]; Drudge Boots (21532, -0.29 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (207.4 DPS) | yes | Band of the Ogre King (18522, -0.17 DPS) [dungeon]; Blackstone Ring (17713, -0.28 DPS) [dungeon]; Naglering (11669, -3.15 DPS, sim-verified) [dungeon] |
| finger2 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-verified (207.4 DPS) | yes | Band of the Ogre King (18522, -0.08 DPS) [dungeon]; Blackstone Ring (17713, -0.19 DPS) [dungeon]; Naglering (11669, -2.69 DPS, sim-verified) [dungeon] |
| trinket1 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-verified (207.4 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| trinket2 | Eye of the Beast (13968) | For The Horde! [quest] | sim-verified (207.4 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Darkmoon Card: Maelstrom (19289, +0.00 DPS) [quest] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (207.4 DPS) | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; Gravestone War Axe (13983, -15.75 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** neck: Medallion of the Dawn; shoulder: Black Dragonscale Shoulders; back: Deathguard's Cloak; chest: Obsidian Mail Tunic; wrist: Windtalker's Wristguards; hands: Timbermaw Brawlers; waist: Dense Timbermaw Belt; legs: Warbear Woolies; feet: Scalegut Treaders; finger1: Don Julio's Band; finger2: Legionnaire's Band; trinket1: Blackhand's Breadth; trinket2: Eye of the Beast; main_hand: The Unstoppable Force

No-known-source sample (15 of 1672, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

