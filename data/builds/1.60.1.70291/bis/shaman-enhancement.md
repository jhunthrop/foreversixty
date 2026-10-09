# Leveling BiS: Enhancement

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-254000000000000000-0000000000000000)

Set DPS (verified): 38.3. Weights run: 1.9s. Verify run: 3.1s. 226 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.569 ± 0.033, crit=1.061 ± 0.032 per rating point (14 rating = 1%, 14.850 per %), hit=1.757 ± 0.083 per rating point (10 rating = 1%, 17.572 per %), melee_haste=10.826 ± 0.322

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (0.57 DPS) | yes | Brawler's Leather Hood (252504, -0.34 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 9.4 attack_power points (0.33 DPS) | yes | Erudite's Amulet (277204, -0.22 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 7.8 attack_power points (0.28 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 9.4 attack_power points (0.33 DPS) | yes | Grave Shroud (279865, -0.01 DPS) [quest]; Dark Leather Cloak (2316, -0.03 DPS) [crafted]; Lambent Scale Cloak (4706, -0.05 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (38.3 DPS) | yes | Tunic of Westfall (2041, +0.00 DPS) [quest]; Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -0.96 DPS, sim-verified) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 10.3 attack_power points (0.36 DPS) | yes | Bristlebark Bindings (14569, -0.06 DPS) [world_drop]; Forest Leather Bracers (3202, -0.09 DPS) [world_drop]; Wolf Bracers (4794, -0.14 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (38.3 DPS) | yes | Fletcher's Gloves (7348, +0.00 DPS) [crafted]; Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Gloves of the Fang (10413, -1.28 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.64 DPS) | yes | Deviate Scale Belt (6468, -0.15 DPS) [crafted]; Ruffian Belt (5975, -0.21 DPS) [world]; Brawler's Leather Belt (252428, -1.59 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (38.3 DPS) | yes | Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -1.26 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 18.5 attack_power points (0.66 DPS) | yes | Brawler's Leather Boots (252439, -0.03 DPS) [crafted]; Defender's Leather Boots (252441, -0.30 DPS) [crafted]; Totemic Leather Boots (252442, -0.30 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 14.3 attack_power points (0.51 DPS) | yes | Demon Band (12054, -0.22 DPS) [world_drop]; The 1 Ring (8350, -0.38 DPS) [world]; Lavishly Jeweled Ring (1156, -0.40 DPS) [dungeon] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 9.4 attack_power points (0.33 DPS) | yes | Demon Band (12054, -0.05 DPS) [world_drop]; The 1 Ring (8350, -0.21 DPS) [world]; Lavishly Jeweled Ring (1156, -0.22 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | The Axe of Severing (23171) | Shadowfang Keep: Sever [dungeon] | 322.0 attack_power points (11.44 DPS) | yes | Living Root (6631, -0.89 DPS) [dungeon]; Smite's Mighty Hammer (7230, -1.13 DPS, sim-verified) [dungeon]; Night Reaver (1318, -1.19 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Bravo's Armbands; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Signet of the Zhevra; main_hand: The Axe of Severing

No-known-source sample (15 of 226, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher

### Band 30 (dwarf, 0000000000000000-255130030002000000-0000000000000000)

Set DPS (verified): 90.5. Weights run: 1.9s. Verify run: 2.0s. 395 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.005, agility=1.513 ± 0.034, crit=1.229 ± 0.038 per rating point (14 rating = 1%, 17.200 per %), hit=2.263 ± 0.094 per rating point (10 rating = 1%, 22.628 per %), melee_haste=11.034 ± 0.408

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (1.20 DPS) | yes | Cloudy Gustwoven Hood (277042, -0.30 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.30 DPS) [crafted]; Defender's Leather Hood (252447, -0.40 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 14.1 attack_power points (0.70 DPS) | yes | Sentinel's Medallion (19541, -0.10 DPS) [rep]; Fallen Guard's Pendant (279837, -0.10 DPS) [quest]; Ghostshard Talisman (7731, -0.55 DPS, sim-verified) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 26.6 attack_power points (1.33 DPS) | yes | Barbaric Shoulders (5964, -0.45 DPS) [crafted]; Mantle of Thieves (2264, -0.57 DPS) [dungeon]; Bristlebark Amice (14573, -0.58 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 16.6 attack_power points (0.83 DPS) | yes | Sergeant Major's Cape (16315, +0.00 DPS, sim-verified) [pvp]; Tigerstrike Mantle (13108, -0.22 DPS) [world_drop]; Wolfmaster Cape (6314, -0.33 DPS) [dungeon] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 24.1 attack_power points (1.20 DPS) | yes | Defender's Leather Tunic (252450, -0.15 DPS) [crafted]; Brawler's Leather Armor (252490, -0.17 DPS) [crafted]; Dusky Leather Armor (7374, -0.96 DPS, sim-verified) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 17.1 attack_power points (0.85 DPS) | yes | Barbaric Bracers (18948, -0.15 DPS) [crafted]; Jurassic Wristguards (6198, -0.20 DPS) [world]; Demonhide Bracers (270033, -0.25 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 23.1 attack_power points (1.15 DPS) | yes | Toughened Leather Gloves (4253, -0.10 DPS) [crafted]; Wolfclaw Gloves (1978, -0.20 DPS) [dungeon]; Fletcher's Gloves (7348, -0.29 DPS) [crafted] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 27.1 attack_power points (1.35 DPS) | yes | Highlander's Chain Girdle (20090, -0.15 DPS) [rep]; Highlander's Leather Girdle (20117, -0.15 DPS) [rep]; Skulker's Leather Belt (252520, -0.73 DPS, sim-verified) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-verified (90.5 DPS) | yes | Brawler's Leather Pants (252500, -0.02 DPS) [crafted]; Trapper's Leather Pants (252501, -0.02 DPS) [crafted]; Brawler's Leather Legguards (252516, -2.15 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 18.1 attack_power points (0.90 DPS) | yes | Brawler's Leather Boots (252439, -0.03 DPS) [crafted]; Insignia Boots (4055, -0.30 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.30 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 21.6 attack_power points (1.08 DPS) | yes | Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Tiger Band (6749, -0.38 DPS) [quest]; Monkey Ring (6748, -0.55 DPS) [quest] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 21.1 attack_power points (1.05 DPS) | yes | Thunderbrow Ring (13097, -0.03 DPS) [world_drop]; Tiger Band (6749, -0.35 DPS) [quest]; Monkey Ring (6748, -0.52 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Corpsemaker (6687, -0.51 DPS) [dungeon]; Darkspear Raider's Reaper (272082, -1.14 DPS) [vendor]; Viscous Hammer (13045, -20.94 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Protector's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 395, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler

### Band 40 (dwarf, 0000000000000000-255130030005102031-0000000000000000)

Set DPS (verified): 112.0. Weights run: 2.4s. Verify run: 2.3s. 622 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.532 ± 0.045, crit=1.499 ± 0.055 per rating point (14 rating = 1%, 20.988 per %), hit=2.219 ± 0.102 per rating point (10 rating = 1%, 22.194 per %), melee_haste=8.528 ± 0.688

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 38.9 attack_power points (2.65 DPS) | yes | Barbaric Iron Helm (7915, -0.48 DPS) [crafted]; Hard Gold Coif (250537, -0.74 DPS) [crafted]; Tusken Helm (6686, -0.88 DPS) [dungeon] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Zealous Shadowshard Pendant (17772, +0.00 DPS) [quest]; Kaleidoscope Chain (13084, -0.19 DPS) [world_drop]; Ghostshard Talisman (7731, -0.19 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 28.9 attack_power points (1.97 DPS) | yes | Forest Tracker Epaulets (2278, -0.14 DPS) [world_drop]; Flintrock Shoulders (7755, -0.24 DPS) [dungeon]; Hard Gold Pauldrons (250539, -0.47 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Hawkeye's Cloak (14593, -0.31 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.41 DPS) [quest]; Dark Hooded Cape (5257, -1.38 DPS, sim-verified) [world] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | sim-verified (112.0 DPS) | yes | Quillward Harness (10583, +0.00 DPS) [dungeon]; Avenger's Armor (1488, -0.60 DPS) [dungeon]; Veteran's Silvered Chain Shirt (250518, -0.65 DPS) [crafted] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 20.1 attack_power points (1.37 DPS) | yes | Branded Leather Bracers (19508, -0.01 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.20 DPS) [world_drop]; Yorgen Bracers (13012, -0.24 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 41.0 attack_power points (2.80 DPS) | yes | Scarlet Gauntlets (10331, -0.11 DPS) [dungeon]; Prowler's Leather Gloves (252524, -0.49 DPS) [crafted]; Skulker's Leather Gloves (252525, -0.52 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 31.8 attack_power points (2.17 DPS) | yes | Boar Champion's Belt (10768, -0.12 DPS) [dungeon]; Highlander's Leather Girdle (20116, -0.12 DPS) [rep]; Officer's Belt (250556, -0.18 DPS) [crafted] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.87 DPS) | yes | Firemane Leggings (13129, -0.27 DPS) [world_drop]; Orcish War Leggings (7929, -0.55 DPS) [crafted]; Legguards of the Vault (9396, -0.57 DPS) [dungeon] |
| feet | Blackforge Greaves (6423) | Uldaman: Ancient Treasure [dungeon] | 33.3 attack_power points (2.28 DPS) | yes | Prowler's Leather Shoes (252465, -0.04 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.17 DPS) [crafted]; Excelsior Boots (4109, -0.38 DPS) [quest] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 28.3 attack_power points (1.93 DPS) | yes | Thunderbrow Ring (13097, -0.52 DPS) [world_drop]; Mark of Kern (2262, -0.56 DPS) [dungeon]; Assault Band (13095, -0.56 DPS) [world_drop] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 21.8 attack_power points (1.49 DPS) | yes | Thunderbrow Ring (13097, -0.08 DPS) [world_drop]; Mark of Kern (2262, -0.12 DPS) [dungeon]; Assault Band (13095, -0.12 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ravager (7717) | Scarlet Monastery: Herod [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonebiter (6830, +0.00 DPS) [quest]; Pendulum of Doom (9425, +0.00 DPS) [dungeon]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Blackforge Greaves; finger1: Protector's Band; finger2: Ironspine's Eye; main_hand: Ravager

No-known-source sample (15 of 622, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 50 (dwarf, 3230000000000000-255130030005102051-0000000000000000)

Set DPS (verified): 154.3. Weights run: 2.3s. Verify run: 4.7s. 796 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.383 ± 0.045, crit=1.584 ± 0.058 per rating point (14 rating = 1%, 22.175 per %), hit=2.411 ± 0.111 per rating point (10 rating = 1%, 24.107 per %), melee_haste=9.952 ± 0.859

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Mail Helmet (223075) | Captain Dirgehammer [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Embrace of the Lycan (9479, +0.00 DPS) [dungeon]; White Bandit Mask (10008, -0.49 DPS) [crafted]; Bloomsprout Headpiece (17767, -0.58 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 28.0 attack_power points (1.98 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.56 DPS) [quest]; Sentinel's Medallion (19539, -0.80 DPS) [rep] |
| shoulder | Knight-Lieutenant's Mail Epaulets (223073) | Captain Dirgehammer [vendor] | 40.2 attack_power points (2.84 DPS) | yes | Failed Flying Experiment (9647, -0.45 DPS) [quest]; Skulker's Leather Shoulder (252535, -0.49 DPS) [crafted]; Prowler's Leather Shoulder (252534, -2.20 DPS, sim-verified) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, +0.00 DPS) [dungeon]; Blisterbane Wrap (12552, -0.08 DPS) [dungeon]; Dark Phantom Cape (13122, -0.08 DPS) [world_drop] |
| chest | Knight's Mail Armor (223078) | Captain Dirgehammer [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warbear Harness (15064, +0.00 DPS) [crafted]; Quillward Harness (10583, -0.28 DPS) [dungeon]; Mixologist's Tunic (12793, -2.41 DPS, sim-verified) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 28.7 attack_power points (2.03 DPS) | yes | Bracers of the Stone Princess (17714, -0.05 DPS) [dungeon]; Arena Bands (18711, -0.05 DPS) [world]; Prowler's Leather Bracers (252539, -0.08 DPS) [crafted] |
| hands | Sergeant Major's Mail Gauntlets (223076) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Raider Gloves (272100, +0.00 DPS) [vendor] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | sim-verified (154.3 DPS) | yes | Girdle of Beastial Fury (11686, +0.00 DPS) [dungeon]; Skulker's Leather Waistguard (252474, -0.09 DPS) [crafted]; Highlander's Chain Girdle (20088, -0.17 DPS) [rep] |
| legs | Knight's Mail Legplates (223074) | Captain Dirgehammer [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpentskin Leggings (8262, +0.00 DPS) [world_drop]; Stormshroud Pants (15057, -0.13 DPS) [crafted]; Gryphon Rider's Leggings (9652, -3.08 DPS, sim-verified) [quest] |
| feet | Sergeant Major's Mail Sabatons (223077) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sandstalker Ankleguards (12470, +0.00 DPS) [dungeon]; Prowler's Leather Boots (252468, +0.00 DPS) [crafted]; Skulker's Leather Boots (252469, +0.00 DPS) [crafted] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 32.4 attack_power points (2.29 DPS) | yes | Thunderbrow Ring (13097, -0.87 DPS) [world_drop]; Mark of Kern (2262, -0.88 DPS) [dungeon]; Blackstone Ring (17713, -0.88 DPS) [dungeon] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 20.4 attack_power points (1.45 DPS) | yes | Thunderbrow Ring (13097, -0.02 DPS) [world_drop]; Mark of Kern (2262, -0.03 DPS) [dungeon]; Blackstone Ring (17713, -0.03 DPS) [dungeon] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smoking Heart of the Mountain (11811, +0.00 DPS) [crafted]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thorium Greatmace (250613, -0.92 DPS) [crafted]; Darkspear Raider's Reaper (272080, -1.61 DPS) [vendor]; Ragehammer (10626, -7.33 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Knight-Lieutenant's Mail Helmet; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Mail Epaulets; back: Dark Hooded Cape; chest: Knight's Mail Armor; wrist: Deepfury Bracers; hands: Sergeant Major's Mail Gauntlets; waist: Prowler's Leather Waistguard; legs: Knight's Mail Legplates; feet: Sergeant Major's Mail Sabatons; finger1: Protector's Band; trinket1: Molten Heart of the Mountain; trinket2: Mark of the Chosen; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 796, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 60 (dwarf, 3230300000000000-255130030005102051-0520000000000000)

Set DPS (verified): 242.7. Weights run: 2.3s. Verify run: 9.4s. 1759 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.768 ± 0.063, crit=2.465 ± 0.090 per rating point (14 rating = 1%, 34.510 per %), hit=3.737 ± 0.164 per rating point (10 rating = 1%, 37.368 per %), melee_haste=16.117 ± 1.387

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Rend (12587, +0.00 DPS) [dungeon]; Crown of Tyranny (13359, -1.04 DPS) [dungeon]; Black Dragonscale Helm (252605, -1.25 DPS) [crafted] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-verified (242.7 DPS) | yes | Pendant of Celerity (22340, +0.00 DPS) [dungeon]; Mark of Fordring (15411, -0.06 DPS) [quest]; Medallion of the Dawn (22659, -0.21 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 98.7 attack_power points (7.13 DPS) | yes | Darkspear Epaulets (272106, -1.19 DPS) [vendor]; Wyrmhide Spaulders (12082, -1.73 DPS) [quest]; Darkspear Pauldrons (272105, -4.94 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 65.4 attack_power points (4.72 DPS) | yes | Cloak of the Honor Guard (20073, -1.63 DPS) [rep]; Windshear Cape (20691, -1.65 DPS) [world]; Cape of the Black Baron (13340, -4.57 DPS, sim-verified) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Timbermaw Tunic (252484, -0.36 DPS) [crafted]; Obsidian Mail Tunic (22191, -1.93 DPS) [crafted]; Tunic of Undead Slaying (23089, -13.52 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slashclaw Bracers (13211, -0.42 DPS) [dungeon]; Blackmist Armguards (12966, -0.59 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -5.87 DPS, sim-verified) [world] |
| hands | Bloodmail Gauntlets (14615) | Scholomance: Lady Illucia Barov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stormshroud Gloves (21278, +0.00 DPS) [crafted]; Raider Gloves (272099, +0.00 DPS) [vendor]; Voone's Vice Grips (13963, -11.80 DPS, sim-verified) [quest] |
| waist | Bloodmail Belt (14614) | Scholomance: Lady Illucia Barov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Ferocity of the Timbermaw (227805, -4.52 DPS, sim-verified) [vendor] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 168.3 attack_power points (12.14 DPS) | yes | Sentinel's Leather Pants (237818, -3.72 DPS) [vendor]; Plaguehound Leggings (18736, -5.62 DPS) [dungeon]; Warbear Woolies (15065, -5.81 DPS) [crafted] |
| feet | Bloodmail Boots (14616) | Scholomance: Lady Illucia Barov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fine Dawn Treaders (227815, -1.04 DPS) [vendor]; Drudge Boots (21532, -1.10 DPS) [quest]; Windreaver Greaves (13967, -10.13 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.73 DPS) [dungeon]; Cutthroat's Signet (272408, -1.86 DPS) [vendor]; Naglering (11669, -9.70 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.55 DPS) [dungeon]; Cutthroat's Signet (272408, -1.68 DPS) [vendor]; Naglering (11669, -10.48 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (+15.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Eye of the Beast (13968, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Eye of the Beast (13968, +0.00 DPS) [quest]; Burst of Knowledge (11832, -7.04 DPS, sim-verified) [dungeon] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Seeping Willow (12969, -20.51 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Beads of Ogre Might; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Bloodmail Gauntlets; waist: Bloodmail Belt; legs: Sentinel's Chain Leggings; feet: Bloodmail Boots; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: The Unstoppable Force

No-known-source sample (15 of 1759, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 60, raid preset (dwarf, 3230031000000000-255030031005102031-0530000000000000)

Set DPS (verified): 647.0. Weights run: 2.4s. Verify run: 9.4s. 1759 eligible items had no known source.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.005, agility=2.279 ± 0.082, crit=2.936 ± 0.110 per rating point (14 rating = 1%, 41.107 per %), hit=5.472 ± 0.293 per rating point (10 rating = 1%, 54.721 per %), melee_haste=24.131 ± 2.115

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Rend (12587, +0.00 DPS) [dungeon]; Backwood Helm (18421, -2.85 DPS) [quest]; Black Dragonscale Helm (252605, -2.85 DPS) [crafted] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pendant of Celerity (22340, +0.00 DPS) [dungeon]; Mark of Fordring (15411, -1.61 DPS) [quest]; Amulet of the Darkmoon (19491, -1.86 DPS) [quest] |
| shoulder | Highlander's Chain Pauldrons (20055) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wyrmhide Spaulders (12082, +0.00 DPS) [quest]; Darkspear Pauldrons (272105, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -13.32 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 82.7 attack_power points (11.49 DPS) | yes | Shroud of Arcane Mastery (22330, -3.89 DPS) [dungeon]; Arcanoweave Cloak (272411, -3.89 DPS) [vendor]; Stalwart Cloak (272415, -5.79 DPS, sim-verified) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Timbermaw Tunic (252484, -3.53 DPS) [crafted]; Obsidian Mail Tunic (22191, -8.76 DPS) [crafted]; Tunic of Undead Slaying (23089, -24.30 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slashclaw Bracers (13211, +0.00 DPS) [dungeon]; Blackmist Armguards (12966, -0.25 DPS) [dungeon]; Bracers of Subterfuge (22668, -1.57 DPS) [quest] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | sim-verified (647.0 DPS) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; Gauntlets of Accuracy (18349, -0.96 DPS) [dungeon]; Raider Gloves (272099, -1.40 DPS) [vendor] |
| waist | Highlander's Chain Girdle (20043) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Ferocity of the Timbermaw (227805, +0.00 DPS) [vendor]; Belt of Preserved Heads (20216, -5.32 DPS, sim-verified) [quest] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 216.7 attack_power points (30.10 DPS) | yes | Sentinel's Leather Pants (237818, -10.13 DPS) [vendor]; Plaguehound Leggings (18736, -13.00 DPS) [dungeon]; Blademaster Leggings (12963, -15.20 DPS) [dungeon] |
| feet | Highlander's Chain Greaves (20050) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodmail Boots (14616, +0.00 DPS) [dungeon]; Fine Dawn Treaders (227815, +0.00 DPS) [vendor]; Windreaver Greaves (13967, -6.37 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -3.18 DPS) [dungeon]; Cutthroat's Signet (272408, -3.50 DPS) [vendor]; Naglering (11669, -18.44 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -2.85 DPS) [dungeon]; Cutthroat's Signet (272408, -3.17 DPS) [vendor]; Naglering (11669, -17.93 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (+8.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Eye of the Beast (13968, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of the Beast (13968, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Blackhand's Breadth (13965, -4.63 DPS, sim-verified) [quest] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Seeping Willow (12969, -39.36 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Beads of Ogre Might; shoulder: Highlander's Chain Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Highlander's Chain Girdle; legs: Sentinel's Chain Leggings; feet: Highlander's Chain Greaves; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: The Unstoppable Force

No-known-source sample (15 of 1759, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

## Horde

### Band 20 (orc, 0000000000000000-254000000000000000-0000000000000000)

Set DPS (verified): 37.8. Weights run: 1.9s. Verify run: 2.9s. 206 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.569 ± 0.033, crit=1.061 ± 0.032 per rating point (14 rating = 1%, 14.850 per %), hit=1.757 ± 0.083 per rating point (10 rating = 1%, 17.572 per %), melee_haste=10.826 ± 0.322

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (0.57 DPS) | yes | Brawler's Leather Hood (252504, -0.31 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 9.4 attack_power points (0.33 DPS) | yes | Erudite's Amulet (277204, -0.22 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 7.8 attack_power points (0.28 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 9.4 attack_power points (0.33 DPS) | yes | Grave Shroud (279865, -0.01 DPS) [quest]; Dark Leather Cloak (2316, -0.03 DPS) [crafted]; Lambent Scale Cloak (4706, -0.05 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Murloc Scale Breastplate (5781, +0.00 DPS) [crafted]; Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -0.69 DPS, sim-verified) [crafted] |
| wrist | Bristlebark Bindings (14569) | World drop [world_drop] | 8.7 attack_power points (0.31 DPS) | yes | Forest Leather Bracers (3202, -0.03 DPS) [world_drop]; Wolf Bracers (4794, -0.09 DPS) [vendor]; Ratchet Wristwraps (274742, -0.14 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fletcher's Gloves (7348, +0.00 DPS) [crafted]; Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Gloves of the Fang (10413, -0.85 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.64 DPS) | yes | Deviate Scale Belt (6468, -0.15 DPS) [crafted]; Ruffian Belt (5975, -0.21 DPS) [world]; Brawler's Leather Belt (252428, -1.17 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -0.81 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 18.5 attack_power points (0.66 DPS) | yes | Brawler's Leather Boots (252439, -0.03 DPS) [crafted]; Defender's Leather Boots (252441, -0.30 DPS) [crafted]; Totemic Leather Boots (252442, -0.30 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 14.3 attack_power points (0.51 DPS) | yes | Demon Band (12054, -0.22 DPS) [world_drop]; Loop of Sacrifice (281673, -0.29 DPS) [quest]; Bounty Hunter's Ring (5351, -0.34 DPS) [quest] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 9.4 attack_power points (0.33 DPS) | yes | Demon Band (12054, -0.05 DPS) [world_drop]; Loop of Sacrifice (281673, -0.12 DPS) [quest]; Bounty Hunter's Ring (5351, -0.17 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hammerbone (270018) | Leaders of the Fang [quest] | sim-verified (37.8 DPS) | yes | The Axe of Severing (23171, +0.00 DPS) [dungeon]; Smite's Mighty Hammer (7230, -0.07 DPS) [dungeon]; Forsaken Greataxe (251533, -0.11 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Bristlebark Bindings; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Signet of the Zhevra; main_hand: Hammerbone

No-known-source sample (15 of 206, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield; 14389 Durability Shoulderpads; 14705 Brackwater Chain Shield

### Band 30 (orc, 0000000000000000-255130030002000000-0000000000000000)

Set DPS (verified): 90.2. Weights run: 1.9s. Verify run: 1.8s. 377 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.005, agility=1.513 ± 0.034, crit=1.229 ± 0.038 per rating point (14 rating = 1%, 17.200 per %), hit=2.263 ± 0.094 per rating point (10 rating = 1%, 22.628 per %), melee_haste=11.034 ± 0.408

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (1.20 DPS) | yes | Cloudy Gustwoven Hood (277042, -0.30 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.30 DPS) [crafted]; Defender's Leather Hood (252447, -0.40 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 14.1 attack_power points (0.70 DPS) | yes | Ghostshard Talisman (7731, -0.00 DPS) [dungeon]; Scout's Medallion (19537, -0.10 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 26.6 attack_power points (1.33 DPS) | yes | Mantle of Thieves (2264, -0.57 DPS) [dungeon]; Bristlebark Amice (14573, -0.58 DPS) [world_drop]; Barbaric Shoulders (5964, -0.62 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 16.6 attack_power points (0.83 DPS) | yes | Construct Cloak (279848, -0.13 DPS) [quest]; Tigerstrike Mantle (13108, -0.22 DPS) [world_drop]; Wildhunter Cloak (16658, -0.33 DPS) [quest] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 24.1 attack_power points (1.20 DPS) | yes | Dusky Leather Armor (7374, -0.15 DPS) [crafted]; Defender's Leather Tunic (252450, -0.15 DPS) [crafted]; Brawler's Leather Armor (252490, -0.17 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 17.1 attack_power points (0.85 DPS) | yes | Barbaric Bracers (18948, -0.15 DPS) [crafted]; Jurassic Wristguards (6198, -0.20 DPS) [world]; Bands of Serra'kis (6902, -0.25 DPS) [dungeon] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 23.1 attack_power points (1.15 DPS) | yes | Toughened Leather Gloves (4253, -0.10 DPS) [crafted]; Wolfclaw Gloves (1978, -0.20 DPS) [dungeon]; Fletcher's Gloves (7348, -0.29 DPS) [crafted] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 27.1 attack_power points (1.35 DPS) | yes | Defiler's Chain Girdle (20152, -0.15 DPS) [rep]; Defiler's Leather Girdle (20191, -0.15 DPS) [rep]; Skulker's Leather Belt (252520, -1.00 DPS, sim-verified) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-verified (90.2 DPS) | yes | Brawler's Leather Pants (252500, -0.02 DPS) [crafted]; Trapper's Leather Pants (252501, -0.02 DPS) [crafted]; Brawler's Leather Legguards (252516, -1.50 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 18.1 attack_power points (0.90 DPS) | yes | Brawler's Leather Boots (252439, -0.03 DPS) [crafted]; Vorrel's Boots (7751, -0.15 DPS) [quest]; Stomping Boots (3741, -0.20 DPS) [quest] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 21.6 attack_power points (1.08 DPS) | yes | Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Band of the Fist (17694, -0.38 DPS) [quest]; Tiger Band (6749, -0.38 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 21.1 attack_power points (1.05 DPS) | yes | Thunderbrow Ring (13097, -0.03 DPS) [world_drop]; Band of the Fist (17694, -0.35 DPS) [quest]; Tiger Band (6749, -0.35 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Corpsemaker (6687, -0.51 DPS) [dungeon]; Darkspear Raider's Reaper (272082, -1.14 DPS) [vendor]; Viscous Hammer (13045, -21.12 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 377, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots

### Band 40 (orc, 0000000000000000-255130030005102031-0000000000000000)

Set DPS (verified): 110.4. Weights run: 2.4s. Verify run: 2.4s. 584 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.532 ± 0.045, crit=1.499 ± 0.055 per rating point (14 rating = 1%, 20.988 per %), hit=2.219 ± 0.102 per rating point (10 rating = 1%, 22.194 per %), melee_haste=8.528 ± 0.688

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 38.9 attack_power points (2.65 DPS) | yes | Barbaric Iron Helm (7915, -0.48 DPS) [crafted]; Hard Gold Coif (250537, -0.74 DPS) [crafted]; Tusken Helm (6686, -0.88 DPS) [dungeon] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Zealous Shadowshard Pendant (17772, +0.00 DPS) [quest]; Ethereal Talisman (4430, -0.05 DPS) [quest]; Kaleidoscope Chain (13084, -0.19 DPS) [world_drop] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 28.9 attack_power points (1.97 DPS) | yes | Forest Tracker Epaulets (2278, -0.14 DPS) [world_drop]; Flintrock Shoulders (7755, -0.24 DPS) [dungeon]; Hard Gold Pauldrons (250539, -0.47 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 23.3 attack_power points (1.59 DPS) | yes | First Sergeant's Cloak (16340, -0.15 DPS) [pvp]; Hawkeye's Cloak (14593, -0.45 DPS) [world_drop]; Construct Cloak (279848, -0.64 DPS) [quest] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | sim-verified (110.4 DPS) | yes | Quillward Harness (10583, +0.00 DPS) [dungeon]; Avenger's Armor (1488, -0.60 DPS) [dungeon]; Veteran's Silvered Chain Shirt (250518, -0.65 DPS) [crafted] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 20.1 attack_power points (1.37 DPS) | yes | Windtalker's Wristguards (19584, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Branded Leather Bracers (19508, -0.01 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 41.0 attack_power points (2.80 DPS) | yes | Scarlet Gauntlets (10331, -0.11 DPS) [dungeon]; Prowler's Leather Gloves (252524, -0.49 DPS) [crafted]; Skulker's Leather Gloves (252525, -0.52 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 31.8 attack_power points (2.17 DPS) | yes | Boar Champion's Belt (10768, -0.12 DPS) [dungeon]; Defiler's Leather Girdle (20192, -0.12 DPS) [rep]; Officer's Belt (250556, -0.18 DPS) [crafted] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.87 DPS) | yes | Firemane Leggings (13129, -0.27 DPS) [world_drop]; Orcish War Leggings (7929, -0.55 DPS) [crafted]; Legguards of the Vault (9396, -0.57 DPS) [dungeon] |
| feet | Blackforge Greaves (6423) | Uldaman: Ancient Treasure [dungeon] | 33.3 attack_power points (2.28 DPS) | yes | Prowler's Leather Shoes (252465, -0.04 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.17 DPS) [crafted]; Excelsior Boots (4109, -0.38 DPS) [quest] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 28.3 attack_power points (1.93 DPS) | yes | Thunderbrow Ring (13097, -0.52 DPS) [world_drop]; Mark of Kern (2262, -0.56 DPS) [dungeon]; Assault Band (13095, -0.56 DPS) [world_drop] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 21.8 attack_power points (1.49 DPS) | yes | Thunderbrow Ring (13097, -0.08 DPS) [world_drop]; Mark of Kern (2262, -0.12 DPS) [dungeon]; Assault Band (13095, -0.12 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ravager (7717) | Scarlet Monastery: Herod [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Pendulum of Doom (9425, +0.00 DPS) [dungeon]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; neck: Scout's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Kolkar Marauder Chain; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Blackforge Greaves; finger1: Legionnaire's Band; finger2: Ironspine's Eye; main_hand: Ravager

No-known-source sample (15 of 584, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (orc, 3230000000000000-255130030005102051-0000000000000000)

Set DPS (verified): 154.4. Weights run: 2.3s. Verify run: 4.6s. 737 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.383 ± 0.045, crit=1.584 ± 0.058 per rating point (14 rating = 1%, 22.175 per %), hit=2.411 ± 0.111 per rating point (10 rating = 1%, 24.107 per %), melee_haste=9.952 ± 0.859

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Embrace of the Lycan (9479, +0.00 DPS) [dungeon]; Blood Guard's Mail Helmet (220820, +0.00 DPS) [vendor]; Bloomsprout Headpiece (17767, -0.09 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 28.0 attack_power points (1.98 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.56 DPS) [quest]; Scout's Medallion (19535, -0.80 DPS) [rep]; Woven Ivy Necklace (19159, -1.27 DPS, sim-verified) [quest] |
| shoulder | Prowler's Leather Shoulder (252534) | Leatherworking [crafted] | 34.4 attack_power points (2.43 DPS) | yes | Blood Guard's Mail Epaulets (220823, +0.00 DPS) [vendor]; Skulker's Leather Shoulder (252535, -0.09 DPS) [crafted]; Failed Flying Experiment (9647, -1.95 DPS, sim-verified) [quest] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, +0.00 DPS) [dungeon]; Blisterbane Wrap (12552, -0.08 DPS) [dungeon]; Dark Phantom Cape (13122, -0.08 DPS) [world_drop] |
| chest | Stone Guard's Mail Armor (220826) | Lady Palanseer [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warbear Harness (15064, +0.00 DPS) [crafted]; Quillward Harness (10583, -0.28 DPS) [dungeon]; Mixologist's Tunic (12793, -2.03 DPS, sim-verified) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 28.7 attack_power points (2.03 DPS) | yes | Arena Bands (18711, +0.00 DPS) [world]; Forest Stalker's Bracers (19589, +0.00 DPS) [pvp] |
| hands | First Sergeant's Mail Gauntlets (220831) | PvP rank 9 · First Sergeant · Horde [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Raider Gloves (272100, +0.00 DPS) [vendor] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | sim-verified (154.4 DPS) | yes | Girdle of Beastial Fury (11686, +0.00 DPS) [dungeon]; Skulker's Leather Waistguard (252474, -0.09 DPS) [crafted]; Defiler's Chain Girdle (20151, -0.17 DPS) [rep] |
| legs | Stone Guard's Mail Legplates (220834) | Lady Palanseer [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stormshroud Pants (15057, -0.13 DPS) [crafted]; Scarlet Leggings (10330, -0.30 DPS) [dungeon]; Serpentskin Leggings (8262, -2.58 DPS, sim-verified) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 41.2 attack_power points (2.91 DPS) | yes | Skulker's Leather Boots (252469, -0.09 DPS) [crafted]; Sandstalker Ankleguards (12470, -0.40 DPS) [dungeon]; Shadefiend Boots (11675, -0.42 DPS) [dungeon] |
| finger1 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 32.4 attack_power points (2.29 DPS) | yes | Ironspine's Eye (7686, -0.85 DPS) [dungeon]; Thunderbrow Ring (13097, -0.87 DPS) [world_drop]; Blackstone Ring (17713, -0.88 DPS) [dungeon] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.70 DPS) | yes | Ironspine's Eye (7686, -0.25 DPS) [dungeon]; Thunderbrow Ring (13097, -0.27 DPS) [world_drop]; Blackstone Ring (17713, -0.28 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -6.18 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -3.08 DPS, sim-verified) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thorium Greatmace (250613, -0.92 DPS) [crafted]; Wildstaff (20556, -1.54 DPS) [quest]; Ragehammer (10626, -11.86 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** neck: Skibi's Pendant; shoulder: Prowler's Leather Shoulder; chest: Stone Guard's Mail Armor; wrist: Deepfury Bracers; hands: First Sergeant's Mail Gauntlets; waist: Prowler's Leather Waistguard; legs: Stone Guard's Mail Legplates; feet: Prowler's Leather Boots; finger1: Legionnaire's Band; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 737, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (orc, 3230300000000000-255130030005102051-0520000000000000)

Set DPS (verified): 239.9. Weights run: 2.3s. Verify run: 9.1s. 1679 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.768 ± 0.063, crit=2.465 ± 0.090 per rating point (14 rating = 1%, 34.510 per %), hit=3.737 ± 0.164 per rating point (10 rating = 1%, 37.368 per %), melee_haste=16.117 ± 1.387

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Skyfury Helm (20134) | The Darkreaver Menace [quest] | 129.5 attack_power points (9.35 DPS) | yes | Outlaw's Collar (279253, -2.93 DPS) [crafted]; Champion's Mail Headguard (227155, -3.54 DPS) [pvp]; Eye of Rend (12587, -4.03 DPS, sim-verified) [dungeon] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pendant of Celerity (22340, +0.00 DPS) [dungeon]; Mark of Fordring (15411, -0.06 DPS) [quest]; Medallion of the Dawn (22659, -0.21 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Truestrike Shoulders (12927, +0.00 DPS) [dungeon]; Champion's Mail Pauldrons (227154, +0.00 DPS) [pvp]; Warlord's Mail Pauldrons (231654, +0.00 DPS) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 65.4 attack_power points (4.72 DPS) | yes | Deathguard's Cloak (20068, -1.63 DPS) [rep]; Windshear Cape (20691, -1.65 DPS) [world]; Cape of the Black Baron (13340, -2.37 DPS, sim-verified) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Timbermaw Tunic (252484, -0.36 DPS) [crafted]; Legionnaire's Mail Hauberk (227157, -1.54 DPS) [pvp]; Tunic of Undead Slaying (23089, -10.88 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slashclaw Bracers (13211, -0.42 DPS) [dungeon]; Blackmist Armguards (12966, -0.59 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -5.68 DPS, sim-verified) [world] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | sim-verified (239.9 DPS) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; Blood Guard's Mail Vices (227159, +0.00 DPS) [pvp]; General's Mail Vices (231655, +0.00 DPS) [vendor] |
| waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | 91.9 attack_power points (6.63 DPS) | yes | Ferocity of the Timbermaw (227805, -0.71 DPS) [vendor]; Marksman's Girdle (22232, -1.26 DPS) [dungeon]; Defiler's Chain Girdle (20150, -1.69 DPS) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 168.3 attack_power points (12.14 DPS) | yes | General's Mail Legguards (231658, -2.91 DPS) [vendor]; Outrider's Chain Leggings (22673, -3.34 DPS, sim-verified) [rep]; Sentinel's Leather Pants (237818, -3.72 DPS) [vendor] |
| feet | Bloodmail Boots (14616) | Scholomance: Lady Illucia Barov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Mail Greaves (231656, +0.00 DPS) [vendor]; Blood Guard's Mail Greaves (227158, -0.43 DPS) [pvp]; Windreaver Greaves (13967, -5.75 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.73 DPS) [dungeon]; Cutthroat's Signet (272408, -1.86 DPS) [vendor]; Naglering (11669, -8.03 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.55 DPS) [dungeon]; Cutthroat's Signet (272408, -1.68 DPS) [vendor]; Naglering (11669, -8.70 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Eye of the Beast (13968, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, -4.41 DPS, sim-verified) [quest] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Eye of the Beast (13968, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; Gravestone War Axe (13983, -15.74 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Skyfury Helm; neck: Beads of Ogre Might; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Belt of Preserved Heads; legs: Sentinel's Chain Leggings; feet: Bloodmail Boots; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: The Unstoppable Force

No-known-source sample (15 of 1679, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60, raid preset (orc, 3230031000000000-255030031005102031-0530000000000000)

Set DPS (verified): 665.3. Weights run: 2.4s. Verify run: 9.6s. 1679 eligible items had no known source.

6 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.005, agility=2.279 ± 0.082, crit=2.936 ± 0.110 per rating point (14 rating = 1%, 41.107 per %), hit=5.472 ± 0.293 per rating point (10 rating = 1%, 54.721 per %), melee_haste=24.131 ± 2.115

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Skyfury Helm (20134) | The Darkreaver Menace [quest] | 151.9 attack_power points (21.10 DPS) | yes | Outlaw's Collar (279253, -5.89 DPS) [crafted]; Eye of Rend (12587, -6.49 DPS, sim-verified) [dungeon]; Champion's Mail Headguard (227155, -8.36 DPS) [pvp] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pendant of Celerity (22340, +0.00 DPS) [dungeon]; Mark of Fordring (15411, -1.61 DPS) [quest]; Amulet of the Darkmoon (19491, -1.86 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Truestrike Shoulders (12927, +0.00 DPS) [dungeon]; Champion's Mail Pauldrons (227154, +0.00 DPS) [pvp]; Warlord's Mail Pauldrons (231654, +0.00 DPS) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 82.7 attack_power points (11.49 DPS) | yes | Shroud of Arcane Mastery (22330, -3.89 DPS) [dungeon]; Arcanoweave Cloak (272411, -3.89 DPS) [vendor]; Stalwart Cloak (272415, -5.84 DPS, sim-verified) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Timbermaw Tunic (252484, -3.53 DPS) [crafted]; Legionnaire's Mail Hauberk (227157, -4.99 DPS) [pvp]; Tunic of Undead Slaying (23089, -20.20 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slashclaw Bracers (13211, +0.00 DPS) [dungeon]; Blackmist Armguards (12966, -0.25 DPS) [dungeon]; Bracers of Subterfuge (22668, -1.57 DPS) [quest] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; Blood Guard's Mail Vices (227159, +0.00 DPS) [pvp]; General's Mail Vices (231655, +0.00 DPS) [vendor] |
| waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marksman's Girdle (22232, -2.38 DPS) [dungeon]; Ferocity of the Timbermaw (227805, -3.30 DPS) [vendor]; Might of the Timbermaw (19044, -5.78 DPS) [crafted] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 216.7 attack_power points (30.10 DPS) | yes | Outrider's Chain Leggings (22673, -6.49 DPS, sim-verified) [rep]; General's Mail Legguards (231658, -8.23 DPS) [vendor]; Legionnaire's Mail Legguards (227156, -10.06 DPS) [pvp] |
| feet | Bloodmail Boots (14616) | Scholomance: Lady Illucia Barov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Mail Greaves (231656, -0.40 DPS) [vendor]; Blood Guard's Mail Greaves (227158, -1.32 DPS) [pvp]; Windreaver Greaves (13967, -10.43 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -3.18 DPS) [dungeon]; Cutthroat's Signet (272408, -3.50 DPS) [vendor]; Naglering (11669, -10.89 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -2.85 DPS) [dungeon]; Cutthroat's Signet (272408, -3.17 DPS) [vendor]; Naglering (11669, -9.70 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Eye of the Beast (13968, +0.00 DPS) [quest]; Hand of Justice (11815, -18.23 DPS, sim-verified) [dungeon] |
| trinket2 | - | - |  |  |  |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; Gravestone War Axe (13983, -44.82 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Skyfury Helm; neck: Beads of Ogre Might; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Belt of Preserved Heads; legs: Sentinel's Chain Leggings; feet: Bloodmail Boots; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; main_hand: The Unstoppable Force

No-known-source sample (15 of 1679, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

