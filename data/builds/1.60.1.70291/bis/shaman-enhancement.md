# Leveling BiS: Enhancement

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-254000000000000000-0000000000000000)

Set DPS (verified): 38.1. Weights run: 1.9s. Verify run: 3.2s. 226 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.578 ± 0.032, crit=1.073 ± 0.033 per rating point (14 rating = 1%, 15.018 per %), hit=1.760 ± 0.083 per rating point (10 rating = 1%, 17.598 per %), melee_haste=10.847 ± 0.323

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (0.57 DPS) | yes | Brawler's Leather Hood (252504, -0.31 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 9.5 attack_power points (0.34 DPS) | yes | Erudite's Amulet (277204, -0.22 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 7.9 attack_power points (0.28 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 9.5 attack_power points (0.34 DPS) | yes | Grave Shroud (279865, -0.01 DPS) [quest]; Dark Leather Cloak (2316, -0.03 DPS) [crafted]; Lambent Scale Cloak (4706, -0.05 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (38.1 DPS) | yes | Tunic of Westfall (2041, +0.00 DPS) [quest]; Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -0.80 DPS, sim-verified) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 10.3 attack_power points (0.37 DPS) | yes | Bristlebark Bindings (14569, -0.06 DPS) [world_drop]; Forest Leather Bracers (3202, -0.09 DPS) [world_drop]; Wolf Bracers (4794, -0.14 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (38.1 DPS) | yes | Fletcher's Gloves (7348, +0.00 DPS) [crafted]; Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Gloves of the Fang (10413, -1.07 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.64 DPS) | yes | Deviate Scale Belt (6468, -0.15 DPS) [crafted]; Ruffian Belt (5975, -0.21 DPS) [world]; Brawler's Leather Belt (252428, -1.37 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (38.1 DPS) | yes | Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -0.82 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 18.6 attack_power points (0.66 DPS) | yes | Brawler's Leather Boots (252439, -0.03 DPS) [crafted]; Defender's Leather Boots (252441, -0.31 DPS) [crafted]; Totemic Leather Boots (252442, -0.31 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 14.3 attack_power points (0.51 DPS) | yes | Demon Band (12054, -0.22 DPS) [world_drop]; The 1 Ring (8350, -0.38 DPS) [world]; Lavishly Jeweled Ring (1156, -0.40 DPS) [dungeon] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 9.5 attack_power points (0.34 DPS) | yes | Demon Band (12054, -0.05 DPS) [world_drop]; The 1 Ring (8350, -0.21 DPS) [world]; Lavishly Jeweled Ring (1156, -0.22 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | The Axe of Severing (23171) | Shadowfang Keep: Sever [dungeon] | 322.0 attack_power points (11.41 DPS) | yes | Smite's Mighty Hammer (7230, -0.87 DPS, sim-verified) [dungeon]; Living Root (6631, -0.89 DPS) [dungeon]; Night Reaver (1318, -1.19 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Bravo's Armbands; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Signet of the Zhevra; main_hand: The Axe of Severing

No-known-source sample (15 of 226, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher

### Band 30 (dwarf, 0000000000000000-255130030002000000-0000000000000000)

Set DPS (verified): 89.2. Weights run: 2.0s. Verify run: 1.9s. 395 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.006, agility=1.636 ± 0.037, crit=1.294 ± 0.039 per rating point (14 rating = 1%, 18.112 per %), hit=2.209 ± 0.100 per rating point (10 rating = 1%, 22.090 per %), melee_haste=10.672 ± 0.384

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (1.20 DPS) | yes | Cloudy Gustwoven Hood (277042, -0.30 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.30 DPS) [crafted]; Brawler's Leather Helm (252512, -0.38 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 14.5 attack_power points (0.73 DPS) | yes | Sentinel's Medallion (19541, -0.07 DPS) [rep]; Fallen Guard's Pendant (279837, -0.13 DPS) [quest]; Ghostshard Talisman (7731, -0.88 DPS, sim-verified) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 28.0 attack_power points (1.40 DPS) | yes | Mantle of Thieves (2264, -0.58 DPS) [dungeon]; Bristlebark Amice (14573, -0.61 DPS) [world_drop]; Barbaric Shoulders (5964, -0.68 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 17.5 attack_power points (0.88 DPS) | yes | Sergeant Major's Cape (16315, -0.15 DPS) [pvp]; Tigerstrike Mantle (13108, -0.22 DPS) [world_drop]; Wolfmaster Cape (6314, -0.37 DPS) [dungeon] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 25.1 attack_power points (1.26 DPS) | yes | Defender's Leather Tunic (252450, -0.16 DPS) [crafted]; Brawler's Leather Armor (252490, -0.18 DPS) [crafted]; Dusky Leather Armor (7374, -0.90 DPS, sim-verified) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 17.8 attack_power points (0.89 DPS) | yes | Barbaric Bracers (18948, -0.16 DPS) [crafted]; Jurassic Wristguards (6198, -0.20 DPS) [world]; Demonhide Bracers (270033, -0.26 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 23.8 attack_power points (1.19 DPS) | yes | Toughened Leather Gloves (4253, -0.10 DPS) [crafted]; Wolfclaw Gloves (1978, -0.20 DPS) [dungeon]; Fletcher's Gloves (7348, -0.29 DPS) [crafted] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 27.8 attack_power points (1.40 DPS) | yes | Highlander's Chain Girdle (20090, -0.19 DPS) [rep]; Highlander's Leather Girdle (20117, -0.19 DPS) [rep]; Skulker's Leather Belt (252520, -1.06 DPS, sim-verified) [crafted] |
| legs | Brawler's Leather Legguards (252516) | Leatherworking [crafted] | 28.7 attack_power points (1.44 DPS) | yes | Trapper's Leather Pants (252501, -0.10 DPS) [crafted]; Ferine Leggings (6690, -0.14 DPS) [dungeon]; Brawler's Leather Pants (252500, -0.86 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 19.1 attack_power points (0.96 DPS) | yes | Brawler's Leather Boots (252439, -0.05 DPS) [crafted]; Insignia Boots (4055, -0.30 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.30 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 22.7 attack_power points (1.14 DPS) | yes | Thunderbrow Ring (13097, -0.09 DPS) [world_drop]; Tiger Band (6749, -0.44 DPS) [quest]; Monkey Ring (6748, -0.57 DPS) [quest] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 21.8 attack_power points (1.09 DPS) | yes | Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Tiger Band (6749, -0.39 DPS) [quest]; Monkey Ring (6748, -0.52 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (89.2 DPS) | yes | Corpsemaker (6687, -0.55 DPS) [dungeon]; Darkspear Raider's Reaper (272082, -1.18 DPS) [vendor]; Viscous Hammer (13045, -21.74 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Brawler's Leather Legguards; finger1: Ironspine's Eye; finger2: Protector's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 395, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler

### Band 40 (dwarf, 0000000000000000-255130030005102031-0000000000000000)

Set DPS (verified): 112.7. Weights run: 2.4s. Verify run: 2.3s. 622 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.530 ± 0.044, crit=1.550 ± 0.056 per rating point (14 rating = 1%, 21.700 per %), hit=2.133 ± 0.105 per rating point (10 rating = 1%, 21.325 per %), melee_haste=7.904 ± 0.736

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 38.8 attack_power points (2.65 DPS) | yes | Barbaric Iron Helm (7915, -0.48 DPS) [crafted]; Hard Gold Coif (250537, -0.74 DPS) [crafted]; Tusken Helm (6686, -0.88 DPS) [dungeon] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Zealous Shadowshard Pendant (17772, +0.00 DPS) [quest]; Kaleidoscope Chain (13084, -0.19 DPS) [world_drop]; Ghostshard Talisman (7731, -0.19 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 28.8 attack_power points (1.97 DPS) | yes | Forest Tracker Epaulets (2278, -0.14 DPS) [world_drop]; Flintrock Shoulders (7755, -0.24 DPS) [dungeon]; Hard Gold Pauldrons (250539, -0.47 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 23.3 attack_power points (1.59 DPS) | yes | Sergeant Major's Cape (16336, -0.15 DPS) [pvp]; Hawkeye's Cloak (14593, -0.45 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.56 DPS) [quest] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | sim-verified (112.7 DPS) | yes | Quillward Harness (10583, +0.00 DPS) [dungeon]; Avenger's Armor (1488, -0.60 DPS) [dungeon]; Veteran's Silvered Chain Shirt (250518, -0.65 DPS) [crafted] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 20.1 attack_power points (1.38 DPS) | yes | Branded Leather Bracers (19508, -0.01 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.20 DPS) [world_drop]; Yorgen Bracers (13012, -0.24 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 41.7 attack_power points (2.85 DPS) | yes | Scarlet Gauntlets (10331, -0.16 DPS) [dungeon]; Prowler's Leather Gloves (252524, -0.54 DPS) [crafted]; Skulker's Leather Gloves (252525, -0.57 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 31.8 attack_power points (2.17 DPS) | yes | Highlander's Leather Girdle (20116, -0.12 DPS) [rep]; Boar Champion's Belt (10768, -0.12 DPS) [dungeon]; Highlander's Chain Girdle (20089, -0.14 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.87 DPS) | yes | Firemane Leggings (13129, -0.27 DPS) [world_drop]; Orcish War Leggings (7929, -0.55 DPS) [crafted]; Legguards of the Vault (9396, -0.57 DPS) [dungeon] |
| feet | Blackforge Greaves (6423) | Uldaman: Ancient Treasure [dungeon] | 33.3 attack_power points (2.28 DPS) | yes | Prowler's Leather Shoes (252465, -0.04 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.17 DPS) [crafted]; Excelsior Boots (4109, -0.38 DPS) [quest] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 28.2 attack_power points (1.93 DPS) | yes | Thunderbrow Ring (13097, -0.52 DPS) [world_drop]; Mark of Kern (2262, -0.56 DPS) [dungeon]; Assault Band (13095, -0.56 DPS) [world_drop] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 21.8 attack_power points (1.49 DPS) | yes | Thunderbrow Ring (13097, -0.08 DPS) [world_drop]; Mark of Kern (2262, -0.12 DPS) [dungeon]; Assault Band (13095, -0.12 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ravager (7717) | Scarlet Monastery: Herod [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonebiter (6830, +0.00 DPS) [quest]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor]; Pendulum of Doom (9425, -3.43 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Kolkar Marauder Chain; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Blackforge Greaves; finger1: Protector's Band; finger2: Ironspine's Eye; main_hand: Ravager

No-known-source sample (15 of 622, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 50 (dwarf, 3230000000000000-255130030005102051-0000000000000000)

Set DPS (verified): 156.2. Weights run: 2.4s. Verify run: 4.9s. 796 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.424 ± 0.049, crit=1.671 ± 0.066 per rating point (14 rating = 1%, 23.394 per %), hit=2.327 ± 0.117 per rating point (10 rating = 1%, 23.266 per %), melee_haste=9.099 ± 0.845

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Mail Helmet (223075) | Captain Dirgehammer [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | White Bandit Mask (10008, -0.55 DPS) [crafted]; Bloomsprout Headpiece (17767, -0.67 DPS) [dungeon]; Embrace of the Lycan (9479, -4.06 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 28.5 attack_power points (2.02 DPS) | yes | Sentinel's Medallion (19539, -0.81 DPS) [rep]; Zealous Shadowshard Pendant (17772, -1.33 DPS, sim-verified) [quest] |
| shoulder | Knight-Lieutenant's Mail Epaulets (223073) | Captain Dirgehammer [vendor] | 41.4 attack_power points (2.93 DPS) | yes | Failed Flying Experiment (9647, -0.51 DPS) [quest]; Skulker's Leather Shoulder (252535, -0.55 DPS) [crafted]; Prowler's Leather Shoulder (252534, -3.32 DPS, sim-verified) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, +0.00 DPS) [dungeon]; Blisterbane Wrap (12552, -0.06 DPS) [dungeon]; Dark Phantom Cape (13122, -0.06 DPS) [world_drop] |
| chest | Knight's Mail Armor (223078) | Captain Dirgehammer [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warbear Harness (15064, +0.00 DPS) [crafted]; Quillward Harness (10583, -0.31 DPS) [dungeon]; Mixologist's Tunic (12793, -5.43 DPS, sim-verified) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 29.4 attack_power points (2.08 DPS) | yes | Bracers of the Stone Princess (17714, -0.10 DPS) [dungeon]; Arena Bands (18711, -0.10 DPS) [world]; Prowler's Leather Bracers (252539, -0.10 DPS) [crafted] |
| hands | Sergeant Major's Mail Gauntlets (223076) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; Fists of The Five Thunders (227022, +0.00 DPS) [vendor]; Raider Gloves (272100, -2.81 DPS, sim-verified) [vendor] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | sim-verified (156.2 DPS) | yes | Girdle of Beastial Fury (11686, +0.00 DPS) [dungeon]; Skulker's Leather Waistguard (252474, -0.08 DPS) [crafted]; Highlander's Chain Girdle (20088, -0.12 DPS) [rep] |
| legs | Knight's Mail Legplates (223074) | Captain Dirgehammer [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpentskin Leggings (8262, +0.00 DPS) [world_drop]; Stormshroud Pants (15057, -0.04 DPS) [crafted]; Gryphon Rider's Leggings (9652, -5.64 DPS, sim-verified) [quest] |
| feet | Sergeant Major's Mail Sabatons (223077) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sandstalker Ankleguards (12470, +0.00 DPS) [dungeon]; Skulker's Leather Boots (252469, +0.00 DPS) [crafted]; Prowler's Leather Boots (252468, -2.63 DPS, sim-verified) [crafted] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 32.8 attack_power points (2.33 DPS) | yes | Thunderbrow Ring (13097, -0.89 DPS) [world_drop]; Mark of Kern (2262, -0.91 DPS) [dungeon]; Blackstone Ring (17713, -0.91 DPS) [dungeon] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 20.8 attack_power points (1.47 DPS) | yes | Thunderbrow Ring (13097, -0.04 DPS) [world_drop]; Mark of Kern (2262, -0.06 DPS) [dungeon]; Blackstone Ring (17713, -0.06 DPS) [dungeon] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thorium Greatmace (250613, -0.92 DPS) [crafted]; Darkspear Raider's Reaper (272080, -1.61 DPS) [vendor]; Ragehammer (10626, -8.66 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Knight-Lieutenant's Mail Helmet; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Mail Epaulets; chest: Knight's Mail Armor; wrist: Deepfury Bracers; hands: Sergeant Major's Mail Gauntlets; waist: Prowler's Leather Waistguard; legs: Knight's Mail Legplates; feet: Sergeant Major's Mail Sabatons; finger1: Protector's Band; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 796, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 60 (dwarf, 3230300000000000-255130030005102051-0520000000000000)

Set DPS (verified): 243.7. Weights run: 2.3s. Verify run: 9.7s. 1759 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.685 ± 0.065, crit=2.377 ± 0.093 per rating point (14 rating = 1%, 33.284 per %), hit=3.881 ± 0.176 per rating point (10 rating = 1%, 38.809 per %), melee_haste=11.635 ± 1.346

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Rend (12587, +0.00 DPS) [dungeon]; Crown of Tyranny (13359, -0.87 DPS) [dungeon]; Black Dragonscale Helm (252605, -1.21 DPS) [crafted] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-verified (243.7 DPS) | yes | Pendant of Celerity (22340, +0.00 DPS) [dungeon]; Mark of Fordring (15411, -0.26 DPS) [quest]; Medallion of the Dawn (22659, -0.40 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 101.6 attack_power points (7.37 DPS) | yes | Darkspear Epaulets (272106, -1.66 DPS) [vendor]; Wyrmhide Spaulders (12082, -1.74 DPS) [quest]; Darkspear Pauldrons (272105, -4.63 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 66.8 attack_power points (4.85 DPS) | yes | Cloak of the Honor Guard (20073, -1.77 DPS) [rep]; Windshear Cape (20691, -1.85 DPS) [world]; Cape of the Black Baron (13340, -3.46 DPS, sim-verified) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Timbermaw Tunic (252484, -0.35 DPS) [crafted]; Obsidian Mail Tunic (22191, -1.82 DPS) [crafted]; Tunic of Undead Slaying (23089, -13.29 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slashclaw Bracers (13211, -0.25 DPS) [dungeon]; Blackmist Armguards (12966, -0.38 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -6.39 DPS, sim-verified) [world] |
| hands | Bloodmail Gauntlets (14615) | Scholomance: Lady Illucia Barov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stormshroud Gloves (21278, +0.00 DPS) [crafted]; Raider Gloves (272099, +0.00 DPS) [vendor]; Voone's Vice Grips (13963, -8.82 DPS, sim-verified) [quest] |
| waist | Bloodmail Belt (14614) | Scholomance: Lady Illucia Barov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Ferocity of the Timbermaw (227805, +0.00 DPS) [vendor]; Belt of Preserved Heads (20216, -5.84 DPS, sim-verified) [quest] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 164.3 attack_power points (11.92 DPS) | yes | Sentinel's Leather Pants (237818, -3.79 DPS) [vendor]; Plaguehound Leggings (18736, -5.44 DPS) [dungeon]; Warbear Woolies (15065, -5.66 DPS) [crafted] |
| feet | Bloodmail Boots (14616) | Scholomance: Lady Illucia Barov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fine Dawn Treaders (227815, -1.06 DPS) [vendor]; Drudge Boots (21532, -1.26 DPS) [quest]; Windreaver Greaves (13967, -8.87 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.74 DPS) [dungeon]; Cutthroat's Signet (272408, -1.86 DPS) [vendor]; Naglering (11669, -11.22 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.65 DPS) [dungeon]; Cutthroat's Signet (272408, -1.77 DPS) [vendor]; Naglering (11669, -10.70 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of the Beast (13968, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Blackhand's Breadth (13965, -7.10 DPS, sim-verified) [quest] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Eye of the Beast (13968, +0.00 DPS) [quest]; Smolderweb's Eye (13213, -6.06 DPS, sim-verified) [dungeon] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Gravestone War Axe (13983, -15.96 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Beads of Ogre Might; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Bloodmail Gauntlets; waist: Bloodmail Belt; legs: Sentinel's Chain Leggings; feet: Bloodmail Boots; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: The Unstoppable Force

No-known-source sample (15 of 1759, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 60, raid preset (dwarf, 3230031000000000-255030031005102031-0530000000000000)

Set DPS (verified): 667.9. Weights run: 2.5s. Verify run: 10.2s. 1759 eligible items had no known source.

5 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.005, agility=2.391 ± 0.086, crit=3.038 ± 0.118 per rating point (14 rating = 1%, 42.531 per %), hit=5.440 ± 0.288 per rating point (10 rating = 1%, 54.403 per %), melee_haste=21.439 ± 2.132

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 114.0 attack_power points (16.07 DPS) | yes | Eye of Rend (12587, -0.05 DPS) [dungeon]; Backwood Helm (18421, -3.00 DPS) [quest]; Black Dragonscale Helm (252605, -3.00 DPS) [crafted] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pendant of Celerity (22340, +0.00 DPS) [dungeon]; Mark of Fordring (15411, -1.39 DPS) [quest]; Amulet of the Darkmoon (19491, -1.55 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wyrmhide Spaulders (12082, +0.00 DPS) [quest]; Truestrike Shoulders (12927, +0.00 DPS) [dungeon]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 82.4 attack_power points (11.62 DPS) | yes | Cape of the Black Baron (13340, -3.74 DPS) [dungeon]; Arcanoweave Cloak (272411, -3.95 DPS) [vendor]; Stalwart Cloak (272415, -3.95 DPS) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Timbermaw Tunic (252484, -3.84 DPS) [crafted]; Obsidian Mail Tunic (22191, -9.43 DPS) [crafted]; Tunic of Undead Slaying (23089, -21.08 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slashclaw Bracers (13211, +0.00 DPS) [dungeon]; Blackmist Armguards (12966, -0.60 DPS) [dungeon]; Bracers of Subterfuge (22668, -1.66 DPS) [quest] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; Gauntlets of Accuracy (18349, -0.94 DPS) [dungeon]; Raider Gloves (272099, -1.15 DPS) [vendor] |
| waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | sim-verified (667.9 DPS) | yes | Marksman's Girdle (22232, -2.32 DPS) [dungeon]; Ferocity of the Timbermaw (227805, -3.27 DPS) [vendor]; Might of the Timbermaw (19044, -5.83 DPS) [crafted] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 223.2 attack_power points (31.46 DPS) | yes | Sentinel's Leather Pants (237818, -10.37 DPS) [vendor]; Plaguehound Leggings (18736, -13.68 DPS) [dungeon] |
| feet | Bloodmail Boots (14616) | Scholomance: Lady Illucia Barov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fine Dawn Treaders (227815, -2.12 DPS) [vendor]; Drudge Boots (21532, -3.73 DPS) [quest]; Windreaver Greaves (13967, -10.32 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -3.20 DPS) [dungeon]; Cutthroat's Signet (272408, -3.53 DPS) [vendor]; Naglering (11669, -8.06 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -3.03 DPS) [dungeon]; Cutthroat's Signet (272408, -3.37 DPS) [vendor]; Naglering (11669, -7.85 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (+14.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Eye of the Beast (13968, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of the Beast (13968, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Blackhand's Breadth (13965, -8.30 DPS, sim-verified) [quest] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Seeping Willow (12969, -39.61 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Beads of Ogre Might; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Belt of Preserved Heads; legs: Sentinel's Chain Leggings; feet: Bloodmail Boots; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: The Unstoppable Force

No-known-source sample (15 of 1759, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

## Horde

### Band 20 (orc, 0000000000000000-254000000000000000-0000000000000000)

Set DPS (verified): 37.8. Weights run: 1.9s. Verify run: 3.0s. 206 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.578 ± 0.032, crit=1.073 ± 0.033 per rating point (14 rating = 1%, 15.018 per %), hit=1.760 ± 0.083 per rating point (10 rating = 1%, 17.598 per %), melee_haste=10.847 ± 0.323

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (0.57 DPS) | yes | Brawler's Leather Hood (252504, -0.30 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 9.5 attack_power points (0.34 DPS) | yes | Erudite's Amulet (277204, -0.22 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 7.9 attack_power points (0.28 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 9.5 attack_power points (0.34 DPS) | yes | Grave Shroud (279865, -0.01 DPS) [quest]; Dark Leather Cloak (2316, -0.03 DPS) [crafted]; Lambent Scale Cloak (4706, -0.05 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Murloc Scale Breastplate (5781, +0.00 DPS) [crafted]; Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -0.65 DPS, sim-verified) [crafted] |
| wrist | Bristlebark Bindings (14569) | World drop [world_drop] | 8.7 attack_power points (0.31 DPS) | yes | Forest Leather Bracers (3202, -0.03 DPS) [world_drop]; Wolf Bracers (4794, -0.09 DPS) [vendor]; Ratchet Wristwraps (274742, -0.14 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fletcher's Gloves (7348, +0.00 DPS) [crafted]; Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Gloves of the Fang (10413, -0.82 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.64 DPS) | yes | Deviate Scale Belt (6468, -0.15 DPS) [crafted]; Ruffian Belt (5975, -0.21 DPS) [world]; Brawler's Leather Belt (252428, -1.13 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -0.79 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 18.6 attack_power points (0.66 DPS) | yes | Brawler's Leather Boots (252439, -0.03 DPS) [crafted]; Defender's Leather Boots (252441, -0.31 DPS) [crafted]; Totemic Leather Boots (252442, -0.31 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 14.3 attack_power points (0.51 DPS) | yes | Demon Band (12054, -0.22 DPS) [world_drop]; Loop of Sacrifice (281673, -0.29 DPS) [quest]; Bounty Hunter's Ring (5351, -0.34 DPS) [quest] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 9.5 attack_power points (0.34 DPS) | yes | Demon Band (12054, -0.05 DPS) [world_drop]; Loop of Sacrifice (281673, -0.12 DPS) [quest]; Bounty Hunter's Ring (5351, -0.17 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hammerbone (270018) | Leaders of the Fang [quest] | sim-verified (37.8 DPS) | yes | The Axe of Severing (23171, +0.00 DPS) [dungeon]; Smite's Mighty Hammer (7230, -0.07 DPS) [dungeon]; Forsaken Greataxe (251533, -0.11 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Bristlebark Bindings; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Signet of the Zhevra; main_hand: Hammerbone

No-known-source sample (15 of 206, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield; 14389 Durability Shoulderpads; 14705 Brackwater Chain Shield

### Band 30 (orc, 0000000000000000-255130030002000000-0000000000000000)

Set DPS (verified): 89.9. Weights run: 2.0s. Verify run: 3.5s. 377 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.006, agility=1.636 ± 0.037, crit=1.294 ± 0.039 per rating point (14 rating = 1%, 18.112 per %), hit=2.209 ± 0.100 per rating point (10 rating = 1%, 22.090 per %), melee_haste=10.672 ± 0.384

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (1.20 DPS) | yes | Cloudy Gustwoven Hood (277042, -0.30 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.30 DPS) [crafted]; Brawler's Leather Helm (252512, -0.38 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 14.5 attack_power points (0.73 DPS) | yes | Ghostshard Talisman (7731, -0.03 DPS) [dungeon]; Scout's Medallion (19537, -0.07 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 28.0 attack_power points (1.40 DPS) | yes | Barbaric Shoulders (5964, -0.49 DPS) [crafted]; Mantle of Thieves (2264, -0.58 DPS) [dungeon]; Bristlebark Amice (14573, -0.61 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 17.5 attack_power points (0.88 DPS) | yes | Construct Cloak (279848, -0.17 DPS) [quest]; Tigerstrike Mantle (13108, -0.22 DPS) [world_drop]; Wildhunter Cloak (16658, -0.37 DPS) [quest] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 25.1 attack_power points (1.26 DPS) | yes | Defender's Leather Tunic (252450, -0.16 DPS) [crafted]; Brawler's Leather Armor (252490, -0.18 DPS) [crafted]; Dusky Leather Armor (7374, -0.50 DPS, sim-verified) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 17.8 attack_power points (0.89 DPS) | yes | Barbaric Bracers (18948, -0.16 DPS) [crafted]; Jurassic Wristguards (6198, -0.20 DPS) [world]; Bands of Serra'kis (6902, -0.29 DPS) [dungeon] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-verified (89.9 DPS) | yes | Wolfclaw Gloves (1978, +0.00 DPS) [dungeon]; Toughened Leather Gloves (4253, +0.00 DPS) [crafted]; Insignia Gloves (6408, +0.00 DPS) [world_drop] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 27.8 attack_power points (1.40 DPS) | yes | Defiler's Chain Girdle (20152, -0.19 DPS) [rep]; Defiler's Leather Girdle (20191, -0.19 DPS) [rep]; Skulker's Leather Belt (252520, -0.77 DPS, sim-verified) [crafted] |
| legs | Leggings of the Fang (10410) | Wailing Caverns: Lord Cobrahn [dungeon] | sim-verified (89.9 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Legguards (252516, -0.74 DPS, sim-verified) [crafted] |
| feet | Footpads of the Fang (10411) | Wailing Caverns: Lord Serpentis [dungeon] | sim-verified (89.9 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Vorrel's Boots (7751, +0.00 DPS) [quest]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 22.7 attack_power points (1.14 DPS) | yes | Thunderbrow Ring (13097, -0.09 DPS) [world_drop]; Band of the Fist (17694, -0.41 DPS) [quest]; Tiger Band (6749, -0.44 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 21.8 attack_power points (1.09 DPS) | yes | Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Band of the Fist (17694, -0.36 DPS) [quest]; Tiger Band (6749, -0.39 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (89.9 DPS) | yes | Corpsemaker (6687, -0.55 DPS) [dungeon]; Darkspear Raider's Reaper (272082, -1.18 DPS) [vendor]; Viscous Hammer (13045, -22.13 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Gloves of the Fang; waist: Prowler's Leather Belt; legs: Leggings of the Fang; feet: Footpads of the Fang; finger1: Ironspine's Eye; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 377, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots

### Band 40 (orc, 0000000000000000-255130030005102031-0000000000000000)

Set DPS (verified): 112.7. Weights run: 2.4s. Verify run: 2.5s. 584 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.530 ± 0.044, crit=1.550 ± 0.056 per rating point (14 rating = 1%, 21.700 per %), hit=2.133 ± 0.105 per rating point (10 rating = 1%, 21.325 per %), melee_haste=7.904 ± 0.736

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 38.8 attack_power points (2.65 DPS) | yes | Barbaric Iron Helm (7915, -0.48 DPS) [crafted]; Hard Gold Coif (250537, -0.74 DPS) [crafted]; Tusken Helm (6686, -0.88 DPS) [dungeon] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Zealous Shadowshard Pendant (17772, +0.00 DPS) [quest]; Ethereal Talisman (4430, -0.05 DPS) [quest]; Kaleidoscope Chain (13084, -0.19 DPS) [world_drop] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 28.8 attack_power points (1.97 DPS) | yes | Forest Tracker Epaulets (2278, -0.14 DPS) [world_drop]; Flintrock Shoulders (7755, -0.24 DPS) [dungeon]; Hard Gold Pauldrons (250539, -0.47 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | First Sergeant's Cloak (16340, -0.15 DPS) [pvp]; Hawkeye's Cloak (14593, -0.45 DPS) [world_drop]; Construct Cloak (279848, -0.64 DPS) [quest] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | sim-verified (112.7 DPS) | yes | Quillward Harness (10583, +0.00 DPS) [dungeon]; Avenger's Armor (1488, -0.60 DPS) [dungeon]; Veteran's Silvered Chain Shirt (250518, -0.65 DPS) [crafted] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 20.1 attack_power points (1.38 DPS) | yes | Windtalker's Wristguards (19584, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Branded Leather Bracers (19508, -0.01 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 41.7 attack_power points (2.85 DPS) | yes | Scarlet Gauntlets (10331, -0.16 DPS) [dungeon]; Prowler's Leather Gloves (252524, -0.54 DPS) [crafted]; Skulker's Leather Gloves (252525, -0.57 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 31.8 attack_power points (2.17 DPS) | yes | Defiler's Leather Girdle (20192, -0.12 DPS) [rep]; Boar Champion's Belt (10768, -0.12 DPS) [dungeon]; Defiler's Chain Girdle (20153, -0.14 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.87 DPS) | yes | Firemane Leggings (13129, -0.27 DPS) [world_drop]; Orcish War Leggings (7929, -0.55 DPS) [crafted]; Legguards of the Vault (9396, -0.57 DPS) [dungeon] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Skulker's Leather Shoes (252531, -0.13 DPS) [crafted]; Excelsior Boots (4109, -0.34 DPS) [quest]; Blackforge Greaves (6423, -1.54 DPS, sim-verified) [dungeon] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 28.2 attack_power points (1.93 DPS) | yes | Thunderbrow Ring (13097, -0.52 DPS) [world_drop]; Mark of Kern (2262, -0.56 DPS) [dungeon]; Assault Band (13095, -0.56 DPS) [world_drop] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 21.8 attack_power points (1.49 DPS) | yes | Thunderbrow Ring (13097, -0.08 DPS) [world_drop]; Mark of Kern (2262, -0.12 DPS) [dungeon]; Assault Band (13095, -0.12 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ravager (7717) | Scarlet Monastery: Herod [dungeon] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor]; Pendulum of Doom (9425, -1.54 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; neck: Scout's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Kolkar Marauder Chain; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Prowler's Leather Shoes; finger1: Legionnaire's Band; finger2: Ironspine's Eye; main_hand: Ravager

No-known-source sample (15 of 584, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (orc, 3230000000000000-255130030005102051-0000000000000000)

Set DPS (verified): 153.3. Weights run: 2.4s. Verify run: 4.8s. 737 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.424 ± 0.049, crit=1.671 ± 0.066 per rating point (14 rating = 1%, 23.394 per %), hit=2.327 ± 0.117 per rating point (10 rating = 1%, 23.266 per %), melee_haste=9.099 ± 0.845

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Embrace of the Lycan (9479, +0.00 DPS) [dungeon]; Blood Guard's Mail Helmet (220820, +0.00 DPS) [vendor]; Bloomsprout Headpiece (17767, -0.12 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 28.5 attack_power points (2.02 DPS) | yes | Woven Ivy Necklace (19159, -0.26 DPS) [quest]; Zealous Shadowshard Pendant (17772, -0.60 DPS) [quest]; Scout's Medallion (19535, -0.81 DPS) [rep] |
| shoulder | Prowler's Leather Shoulder (252534) | Leatherworking [crafted] | 34.8 attack_power points (2.47 DPS) | yes | Blood Guard's Mail Epaulets (220823, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.04 DPS) [quest]; Skulker's Leather Shoulder (252535, -0.08 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, +0.00 DPS) [dungeon]; Blisterbane Wrap (12552, -0.06 DPS) [dungeon]; Dark Phantom Cape (13122, -0.06 DPS) [world_drop] |
| chest | Stone Guard's Mail Armor (220826) | Lady Palanseer [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warbear Harness (15064, +0.00 DPS) [crafted]; Quillward Harness (10583, -0.31 DPS) [dungeon]; Mixologist's Tunic (12793, -3.46 DPS, sim-verified) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 29.4 attack_power points (2.08 DPS) | yes | Arena Bands (18711, +0.00 DPS) [world]; Forest Stalker's Bracers (19589, +0.00 DPS) [pvp] |
| hands | First Sergeant's Mail Gauntlets (220831) | PvP rank 9 · First Sergeant · Horde [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; Fists of The Five Thunders (227022, +0.00 DPS) [vendor]; Raider Gloves (272100, +0.00 DPS) [vendor] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | sim-verified (153.3 DPS) | yes | Girdle of Beastial Fury (11686, +0.00 DPS) [dungeon]; Skulker's Leather Waistguard (252474, -0.08 DPS) [crafted]; Defiler's Chain Girdle (20151, -0.12 DPS) [rep] |
| legs | Stone Guard's Mail Legplates (220834) | Lady Palanseer [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stormshroud Pants (15057, -0.04 DPS) [crafted]; Scarlet Leggings (10330, -0.38 DPS) [dungeon]; Serpentskin Leggings (8262, -3.48 DPS, sim-verified) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 41.7 attack_power points (2.95 DPS) | yes | Skulker's Leather Boots (252469, -0.08 DPS) [crafted]; Sandstalker Ankleguards (12470, -0.39 DPS) [dungeon]; Shadefiend Boots (11675, -0.43 DPS) [dungeon] |
| finger1 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 32.8 attack_power points (2.33 DPS) | yes | Ironspine's Eye (7686, -0.85 DPS) [dungeon]; Thunderbrow Ring (13097, -0.89 DPS) [world_drop]; Blackstone Ring (17713, -0.91 DPS) [dungeon] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.70 DPS) | yes | Ironspine's Eye (7686, -0.23 DPS) [dungeon]; Thunderbrow Ring (13097, -0.26 DPS) [world_drop]; Blackstone Ring (17713, -0.28 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -2.65 DPS) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thorium Greatmace (250613, -0.92 DPS) [crafted]; Wildstaff (20556, -1.52 DPS) [quest]; Ragehammer (10626, -8.61 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** neck: Skibi's Pendant; shoulder: Prowler's Leather Shoulder; chest: Stone Guard's Mail Armor; wrist: Deepfury Bracers; hands: First Sergeant's Mail Gauntlets; waist: Prowler's Leather Waistguard; legs: Stone Guard's Mail Legplates; feet: Prowler's Leather Boots; finger1: Legionnaire's Band; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Molten Heart of the Mountain; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 737, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (orc, 3230300000000000-255130030005102051-0520000000000000)

Set DPS (verified): 246.2. Weights run: 2.3s. Verify run: 9.2s. 1679 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.685 ± 0.065, crit=2.377 ± 0.093 per rating point (14 rating = 1%, 33.284 per %), hit=3.881 ± 0.176 per rating point (10 rating = 1%, 38.809 per %), melee_haste=11.635 ± 1.346

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Skyfury Helm (20134) | The Darkreaver Menace [quest] | 125.9 attack_power points (9.13 DPS) | yes | Eye of Rend (12587, -2.08 DPS, sim-verified) [dungeon]; Outlaw's Collar (279253, -2.94 DPS) [crafted]; Champion's Mail Headguard (227155, -3.38 DPS) [pvp] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pendant of Celerity (22340, +0.00 DPS) [dungeon]; Mark of Fordring (15411, -0.26 DPS) [quest]; Medallion of the Dawn (22659, -0.40 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 101.6 attack_power points (7.37 DPS) | yes | Warlord's Mail Pauldrons (231654, +0.00 DPS) [vendor]; Champion's Mail Pauldrons (227154, -0.26 DPS) [pvp]; Darkspear Pauldrons (272105, -1.66 DPS) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 66.8 attack_power points (4.85 DPS) | yes | Cape of the Black Baron (13340, -1.56 DPS) [dungeon]; Deathguard's Cloak (20068, -1.77 DPS) [rep]; Windshear Cape (20691, -1.85 DPS) [world] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Timbermaw Tunic (252484, -0.35 DPS) [crafted]; Legionnaire's Mail Hauberk (227157, -1.33 DPS) [pvp]; Tunic of Undead Slaying (23089, -6.51 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slashclaw Bracers (13211, -0.25 DPS) [dungeon]; Blackmist Armguards (12966, -0.38 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -3.56 DPS, sim-verified) [world] |
| hands | Bloodmail Gauntlets (14615) | Scholomance: Lady Illucia Barov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Savage Gladiator Grips (11730, +0.00 DPS) [dungeon]; Voone's Vice Grips (13963, +0.00 DPS) [quest]; General's Mail Vices (231655, +0.00 DPS) [vendor] |
| waist | Ferocity of the Timbermaw (227805) | Meilosh [vendor] | sim-verified (+3.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Marksman's Girdle (22232, -0.47 DPS) [dungeon]; Defiler's Chain Girdle (20150, -0.97 DPS) [rep]; Belt of Preserved Heads (20216, -3.23 DPS, sim-verified) [quest] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 164.3 attack_power points (11.92 DPS) | yes | Outrider's Chain Leggings (22673, -2.41 DPS) [rep]; General's Mail Legguards (231658, -2.63 DPS) [vendor]; Legionnaire's Mail Legguards (227156, -3.50 DPS) [pvp] |
| feet | Bloodmail Boots (14616) | Scholomance: Lady Illucia Barov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windreaver Greaves (13967, +0.00 DPS) [dungeon]; General's Mail Greaves (231656, +0.00 DPS) [vendor]; Savage Gladiator Greaves (11731, -3.26 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.74 DPS) [dungeon]; Cutthroat's Signet (272408, -1.86 DPS) [vendor]; Naglering (11669, -3.20 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.65 DPS) [dungeon]; Cutthroat's Signet (272408, -1.77 DPS) [vendor]; Naglering (11669, -5.35 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of the Beast (13968, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Blackhand's Breadth (13965, -11.85 DPS, sim-verified) [quest] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (246.2 DPS) | yes | Blackhand's Breadth (13965, -0.19 DPS) [quest]; Eye of the Beast (13968, -0.19 DPS) [quest]; Frozen Heart of the Mountain (249469, -2.48 DPS) [crafted] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; Gravestone War Axe (13983, -16.73 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Skyfury Helm; neck: Beads of Ogre Might; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Bloodmail Gauntlets; waist: Ferocity of the Timbermaw; legs: Sentinel's Chain Leggings; feet: Bloodmail Boots; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: The Unstoppable Force

No-known-source sample (15 of 1679, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60, raid preset (orc, 3230031000000000-255030031005102031-0530000000000000)

Set DPS (verified): 671.4. Weights run: 2.5s. Verify run: 9.9s. 1679 eligible items had no known source.

6 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.005, agility=2.391 ± 0.086, crit=3.038 ± 0.118 per rating point (14 rating = 1%, 42.531 per %), hit=5.440 ± 0.288 per rating point (10 rating = 1%, 54.403 per %), melee_haste=21.439 ± 2.132

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skyfury Helm (20134, +0.00 DPS) [quest]; Eye of Rend (12587, -0.05 DPS) [dungeon]; Champion's Mail Headguard (227155, -2.94 DPS) [pvp] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pendant of Celerity (22340, +0.00 DPS) [dungeon]; Mark of Fordring (15411, -1.39 DPS) [quest]; Amulet of the Darkmoon (19491, -1.55 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Truestrike Shoulders (12927, +0.00 DPS) [dungeon]; Champion's Mail Pauldrons (227154, +0.00 DPS) [pvp]; Warlord's Mail Pauldrons (231654, +0.00 DPS) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 82.4 attack_power points (11.62 DPS) | yes | Cape of the Black Baron (13340, -3.74 DPS) [dungeon]; Arcanoweave Cloak (272411, -3.95 DPS) [vendor]; Stalwart Cloak (272415, -3.95 DPS) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Timbermaw Tunic (252484, -3.84 DPS) [crafted]; Legionnaire's Mail Hauberk (227157, -5.65 DPS) [pvp]; Tunic of Undead Slaying (23089, -24.13 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slashclaw Bracers (13211, +0.00 DPS) [dungeon]; Blackmist Armguards (12966, -0.60 DPS) [dungeon]; Bracers of Subterfuge (22668, -1.66 DPS) [quest] |
| hands | Savage Gladiator Grips (11730) | Blackrock Depths: Eviscerator [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; Blood Guard's Mail Vices (227159, +0.00 DPS) [pvp]; General's Mail Vices (231655, +0.00 DPS) [vendor] |
| waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | 121.1 attack_power points (17.07 DPS) | yes | Ferocity of the Timbermaw (227805, -3.27 DPS) [vendor]; Might of the Timbermaw (19044, -5.83 DPS) [crafted]; Marksman's Girdle (22232, -6.19 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 223.2 attack_power points (31.46 DPS) | yes | Outrider's Chain Leggings (22673, -6.00 DPS) [rep]; General's Mail Legguards (231658, -9.11 DPS) [vendor]; Sentinel's Leather Pants (237818, -10.37 DPS) [vendor] |
| feet | Bloodmail Boots (14616) | Scholomance: Lady Illucia Barov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windreaver Greaves (13967, +0.00 DPS) [dungeon]; General's Mail Greaves (231656, -0.55 DPS) [vendor]; Savage Gladiator Greaves (11731, -9.93 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -3.20 DPS) [dungeon]; Cutthroat's Signet (272408, -3.53 DPS) [vendor]; Naglering (11669, -12.94 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -3.03 DPS) [dungeon]; Cutthroat's Signet (272408, -3.37 DPS) [vendor]; Naglering (11669, -11.45 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Eye of the Beast (13968, +0.00 DPS) [quest]; Heart of Wyrmthalak (22321, -18.41 DPS, sim-verified) [dungeon] |
| trinket2 | - | - |  |  |  |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; Seeping Willow (12969, -47.68 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Beads of Ogre Might; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Savage Gladiator Grips; waist: Belt of Preserved Heads; legs: Sentinel's Chain Leggings; feet: Bloodmail Boots; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; main_hand: The Unstoppable Force

No-known-source sample (15 of 1679, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

