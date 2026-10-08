# Leveling BiS: Enhancement

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-254000000000000000-0000000000000000)

Set DPS (verified): 38.1. Weights run: 1.2s. Verify run: 2.1s. 225 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.578 ± 0.032, crit=1.073 ± 0.033 per rating point (14 rating = 1%, 15.018 per %), hit=1.760 ± 0.083 per rating point (10 rating = 1%, 17.598 per %), melee_haste=10.847 ± 0.323

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (0.57 DPS) | yes | Brawler's Leather Hood (252504, -0.31 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 9.5 attack_power points (0.34 DPS) | yes | Erudite's Amulet (277204, -0.11 DPS) [quest] |
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

No-known-source sample (15 of 225, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher

### Band 30 (dwarf, 0000000000000000-255130030002000000-0000000000000000)

Set DPS (verified): 89.2. Weights run: 1.3s. Verify run: 1.3s. 366 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.006, agility=1.636 ± 0.037, crit=1.294 ± 0.039 per rating point (14 rating = 1%, 18.112 per %), hit=2.209 ± 0.100 per rating point (10 rating = 1%, 22.090 per %), melee_haste=10.672 ± 0.384

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (1.20 DPS) | yes | Cloudy Gustwoven Hood (277042, -0.30 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.30 DPS) [crafted]; Brawler's Leather Helm (252512, -0.38 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 14.5 attack_power points (0.73 DPS) | yes | Sentinel's Medallion (19541, -0.07 DPS) [rep]; Ghostshard Talisman (7731, -0.88 DPS, sim-verified) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 28.0 attack_power points (1.40 DPS) | yes | Mantle of Thieves (2264, -0.58 DPS) [dungeon]; Bristlebark Amice (14573, -0.61 DPS) [world_drop]; Barbaric Shoulders (5964, -0.68 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 17.5 attack_power points (0.88 DPS) | yes | Sergeant Major's Cape (16315, -0.15 DPS) [pvp]; Tigerstrike Mantle (13108, -0.22 DPS) [world_drop]; Wolfmaster Cape (6314, -0.37 DPS) [dungeon] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 25.1 attack_power points (1.26 DPS) | yes | Defender's Leather Tunic (252450, -0.16 DPS) [crafted]; Brawler's Leather Armor (252490, -0.18 DPS) [crafted]; Dusky Leather Armor (7374, -0.90 DPS, sim-verified) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 17.8 attack_power points (0.89 DPS) | yes | Barbaric Bracers (18948, -0.16 DPS) [crafted]; Jurassic Wristguards (6198, -0.20 DPS) [world]; Demonhide Bracers (270033, -0.26 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 23.8 attack_power points (1.19 DPS) | yes | Toughened Leather Gloves (4253, -0.10 DPS) [crafted]; Wolfclaw Gloves (1978, -0.20 DPS) [dungeon]; Fletcher's Gloves (7348, -0.29 DPS) [crafted] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 27.8 attack_power points (1.40 DPS) | yes | Highlander's Chain Girdle (20090, -0.19 DPS) [rep]; Highlander's Leather Girdle (20117, -0.19 DPS) [rep]; Skulker's Leather Belt (252520, -1.06 DPS, sim-verified) [crafted] |
| legs | Brawler's Leather Legguards (252516) | Leatherworking [crafted] | 28.7 attack_power points (1.44 DPS) | yes | Trapper's Leather Pants (252501, -0.10 DPS) [crafted]; Ferine Leggings (6690, -0.14 DPS) [dungeon]; Brawler's Leather Pants (252500, -0.86 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 19.1 attack_power points (0.96 DPS) | yes | Brawler's Leather Boots (252439, -0.05 DPS) [crafted]; Insignia Boots (4055, -0.30 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.30 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 22.7 attack_power points (1.14 DPS) | yes | Thunderbrow Ring (13097, -0.09 DPS) [world_drop]; Tiger Band (6749, -0.54 DPS) [quest]; Monkey Ring (6748, -0.57 DPS) [quest] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 21.8 attack_power points (1.09 DPS) | yes | Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Tiger Band (6749, -0.49 DPS) [quest]; Monkey Ring (6748, -0.52 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (89.2 DPS) | yes | Corpsemaker (6687, -0.55 DPS) [dungeon]; Darkspear Raider's Reaper (272082, -1.18 DPS) [vendor]; Viscous Hammer (13045, -21.74 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Brawler's Leather Legguards; finger1: Ironspine's Eye; finger2: Protector's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 366, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler

### Band 40 (dwarf, 0000000000000000-255130030005102031-0000000000000000)

Set DPS (verified): 114.4. Weights run: 1.5s. Verify run: 1.4s. 592 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.530 ± 0.044, crit=1.550 ± 0.056 per rating point (14 rating = 1%, 21.700 per %), hit=2.133 ± 0.105 per rating point (10 rating = 1%, 21.325 per %), melee_haste=7.904 ± 0.736

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 47.7 attack_power points (3.26 DPS) | yes | White Bandit Mask (10008, -0.61 DPS) [crafted]; Barbaric Iron Helm (7915, -1.09 DPS) [crafted]; Hard Gold Coif (250537, -1.35 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.37 DPS) | yes | Kaleidoscope Chain (13084, -0.40 DPS) [world_drop]; Ghostshard Talisman (7731, -0.41 DPS) [dungeon]; Sentinel's Medallion (19540, -1.76 DPS, sim-verified) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 28.8 attack_power points (1.97 DPS) | yes | Forest Tracker Epaulets (2278, -0.14 DPS) [world_drop]; Flintrock Shoulders (7755, -0.24 DPS) [dungeon]; Hard Gold Pauldrons (250539, -0.47 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 23.3 attack_power points (1.59 DPS) | yes | Sergeant Major's Cape (16336, -0.15 DPS) [pvp]; Hawkeye's Cloak (14593, -0.45 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.56 DPS) [quest] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 45.1 attack_power points (3.08 DPS) | yes | Kolkar Marauder Chain (6773, -0.44 DPS) [quest]; Avenger's Armor (1488, -1.03 DPS) [dungeon]; Veteran's Silvered Chain Shirt (250518, -1.09 DPS) [crafted] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 20.1 attack_power points (1.38 DPS) | yes | Branded Leather Bracers (19508, -0.01 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.20 DPS) [world_drop]; Yorgen Bracers (13012, -0.24 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 41.7 attack_power points (2.85 DPS) | yes | Scarlet Gauntlets (10331, -0.16 DPS) [dungeon]; Prowler's Leather Gloves (252524, -0.54 DPS) [crafted]; Skulker's Leather Gloves (252525, -0.57 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 31.8 attack_power points (2.17 DPS) | yes | Highlander's Leather Girdle (20116, -0.12 DPS) [rep]; Boar Champion's Belt (10768, -0.12 DPS) [dungeon]; Highlander's Chain Girdle (20089, -0.14 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.87 DPS) | yes | Firemane Leggings (13129, -0.27 DPS) [world_drop]; Orcish War Leggings (7929, -0.55 DPS) [crafted]; Legguards of the Vault (9396, -0.57 DPS) [dungeon] |
| feet | Blackforge Greaves (6423) | Uldaman: Ancient Treasure [dungeon] | 33.3 attack_power points (2.28 DPS) | yes | Prowler's Leather Shoes (252465, -0.04 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.17 DPS) [crafted]; Excelsior Boots (4109, -0.38 DPS) [quest] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 28.2 attack_power points (1.93 DPS) | yes | Thunderbrow Ring (13097, -0.52 DPS) [world_drop]; Mark of Kern (2262, -0.56 DPS) [dungeon]; Assault Band (13095, -0.56 DPS) [world_drop] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 21.8 attack_power points (1.49 DPS) | yes | Thunderbrow Ring (13097, -0.08 DPS) [world_drop]; Mark of Kern (2262, -0.12 DPS) [dungeon]; Assault Band (13095, -0.12 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ravager (7717) | Scarlet Monastery: Herod [dungeon] | sim-verified (114.4 DPS) | yes | Bonebiter (6830, +0.00 DPS) [quest]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor]; Pendulum of Doom (9425, -3.55 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Blackforge Greaves; finger1: Protector's Band; finger2: Ironspine's Eye; main_hand: Ravager

No-known-source sample (15 of 592, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 50 (dwarf, 3230000000000000-255130030005102051-0000000000000000)

Set DPS (verified): 159.8. Weights run: 1.5s. Verify run: 3.2s. 762 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.424 ± 0.049, crit=1.671 ± 0.066 per rating point (14 rating = 1%, 23.394 per %), hit=2.327 ± 0.117 per rating point (10 rating = 1%, 23.266 per %), melee_haste=9.099 ± 0.845

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Mail Helmet (223075) | Captain Dirgehammer [vendor] | sim-verified (159.8 DPS) | yes | Embrace of the Lycan (9479, +0.00 DPS) [dungeon]; White Bandit Mask (10008, -0.55 DPS) [crafted]; Raging Berserker's Helm (7719, -3.77 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 28.5 attack_power points (2.02 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.60 DPS) [quest]; Sentinel's Medallion (19539, -0.81 DPS) [rep] |
| shoulder | Knight-Lieutenant's Mail Epaulets (223073) | Captain Dirgehammer [vendor] | 41.4 attack_power points (2.93 DPS) | yes | Failed Flying Experiment (9647, -0.51 DPS) [quest]; Skulker's Leather Shoulder (252535, -0.55 DPS) [crafted]; Prowler's Leather Shoulder (252534, -2.66 DPS, sim-verified) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 31.9 attack_power points (2.26 DPS) | yes | Dark Hooded Cape (5257, -0.69 DPS) [world]; Blisterbane Wrap (12552, -0.75 DPS) [dungeon]; Dark Phantom Cape (13122, -0.75 DPS) [world_drop] |
| chest | Knight's Mail Armor (223078) | Captain Dirgehammer [vendor] | sim-verified (159.8 DPS) | yes | Warbear Harness (15064, +0.00 DPS) [crafted]; Quillward Harness (10583, -0.31 DPS) [dungeon]; Mixologist's Tunic (12793, -4.86 DPS, sim-verified) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 29.4 attack_power points (2.08 DPS) | yes | Bracers of the Stone Princess (17714, -0.10 DPS) [dungeon]; Arena Bands (18711, -0.10 DPS) [world]; Prowler's Leather Bracers (252539, -0.10 DPS) [crafted] |
| hands | Sergeant Major's Mail Gauntlets (223076) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-verified (159.8 DPS) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; Fists of The Five Thunders (227022, +0.00 DPS) [vendor]; Raider Gloves (272100, -2.50 DPS, sim-verified) [vendor] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (3.26 DPS) | yes | Prowler's Leather Waistguard (252473, -0.06 DPS) [crafted]; Skulker's Leather Waistguard (252474, -0.15 DPS) [crafted]; Highlander's Chain Girdle (20088, -0.18 DPS) [rep] |
| legs | Knight's Mail Legplates (223074) | Captain Dirgehammer [vendor] | sim-verified (159.8 DPS) | yes | Serpentskin Leggings (8262, +0.00 DPS) [world_drop]; Stormshroud Pants (15057, -0.04 DPS) [crafted]; Gryphon Rider's Leggings (9652, -4.65 DPS, sim-verified) [quest] |
| feet | Sergeant Major's Mail Sabatons (223077) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-verified (159.8 DPS) | yes | Sandstalker Ankleguards (12470, +0.00 DPS) [dungeon]; Skulker's Leather Boots (252469, +0.00 DPS) [crafted]; Prowler's Leather Boots (252468, -2.29 DPS, sim-verified) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 43.3 attack_power points (3.07 DPS) | yes | Ironspine's Eye (7686, -1.59 DPS) [dungeon]; Thunderbrow Ring (13097, -1.63 DPS) [world_drop]; Mark of Kern (2262, -1.65 DPS) [dungeon] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 32.8 attack_power points (2.33 DPS) | yes | Ironspine's Eye (7686, -0.85 DPS) [dungeon]; Thunderbrow Ring (13097, -0.89 DPS) [world_drop]; Mark of Kern (2262, -0.91 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (159.8 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (159.8 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (159.8 DPS) | yes | Thorium Greatmace (250613, -0.92 DPS) [crafted]; Darkspear Raider's Reaper (272080, -1.61 DPS) [vendor]; Ragehammer (10626, -9.48 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Knight-Lieutenant's Mail Helmet; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Mail Epaulets; back: Blackveil Cape; chest: Knight's Mail Armor; wrist: Deepfury Bracers; hands: Sergeant Major's Mail Gauntlets; waist: Girdle of Beastial Fury; legs: Knight's Mail Legplates; feet: Sergeant Major's Mail Sabatons; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 762, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 60 (dwarf, 3230300000000000-255130030005102051-0520000000000000)

Set DPS (verified): 235.0. Weights run: 1.5s. Verify run: 5.0s. 1751 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.685 ± 0.065, crit=2.377 ± 0.093 per rating point (14 rating = 1%, 33.284 per %), hit=3.881 ± 0.176 per rating point (10 rating = 1%, 38.809 per %), melee_haste=11.635 ± 1.346

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Rend (12587) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (+5.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Outlaw's Collar (279253, -0.53 DPS) [crafted]; Ragefury Eyepatch (11735, -1.02 DPS) [dungeon]; Mask of the Unforgiven (13404, -5.60 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 64.1 attack_power points (4.65 DPS) | yes | Beads of Ogre Might (22150, -0.09 DPS) [quest]; Mark of Fordring (15411, -0.35 DPS) [quest]; Medallion of the Dawn (22659, -0.49 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | sim-verified (+3.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Wyrmhide Spaulders (12082, -0.08 DPS) [quest]; Truestrike Shoulders (12927, -3.88 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 66.8 attack_power points (4.85 DPS) | yes | Cape of the Black Baron (13340, -1.56 DPS) [dungeon]; Cloak of the Honor Guard (20073, -1.77 DPS) [rep]; Windshear Cape (20691, -1.85 DPS) [world] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Timbermaw Tunic (252484, -0.35 DPS) [crafted]; Obsidian Mail Tunic (22191, -1.82 DPS) [crafted]; Tunic of Undead Slaying (23089, -6.33 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slashclaw Bracers (13211, -0.25 DPS) [dungeon]; Blackmist Armguards (12966, -0.38 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -4.86 DPS, sim-verified) [world] |
| hands | Voone's Vice Grips (13963) | Maxwell's Mission [quest] | 92.8 attack_power points (6.73 DPS) | yes | Stormshroud Gloves (21278, +0.00 DPS, sim-verified) [crafted]; Raider Gloves (272099, -1.83 DPS) [vendor]; Timbermaw Brawlers (19049, -1.97 DPS) [crafted] |
| waist | Ferocity of the Timbermaw (227805) | Meilosh [vendor] | sim-verified (+3.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Marksman's Girdle (22232, -0.47 DPS) [dungeon]; Highlander's Chain Girdle (20043, -0.97 DPS) [rep]; Belt of Preserved Heads (20216, -3.23 DPS, sim-verified) [quest] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 164.3 attack_power points (11.92 DPS) | yes | Sentinel's Leather Pants (237818, -3.79 DPS) [vendor]; Plaguehound Leggings (18736, -5.44 DPS) [dungeon]; Warbear Woolies (15065, -5.66 DPS) [crafted] |
| feet | Bloodmail Boots (14616) | Scholomance: Lady Illucia Barov [dungeon] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Fine Dawn Treaders (227815, -1.06 DPS) [vendor]; Drudge Boots (21532, -1.26 DPS) [quest]; Windreaver Greaves (13967, -2.41 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.74 DPS) [dungeon]; Cutthroat's Signet (272408, -1.86 DPS) [vendor]; Naglering (11669, -3.89 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -0.37 DPS) [dungeon]; Cutthroat's Signet (272408, -0.49 DPS) [vendor]; Naglering (11669, -1.85 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (+10.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Eye of the Beast (13968, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Eye of the Beast (13968, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Seeping Willow (12969, -16.25 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Eye of Rend; neck: Pendant of Celerity; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Voone's Vice Grips; waist: Ferocity of the Timbermaw; legs: Sentinel's Chain Leggings; feet: Bloodmail Boots; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: The Unstoppable Force

No-known-source sample (15 of 1751, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 60, raid preset (dwarf, 3230031000000000-255030031005102031-0530000000000000)

Set DPS (verified): 664.5. Weights run: 1.6s. Verify run: 6.8s. 1751 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.005, agility=2.391 ± 0.086, crit=3.038 ± 0.118 per rating point (14 rating = 1%, 42.531 per %), hit=5.440 ± 0.288 per rating point (10 rating = 1%, 54.403 per %), melee_haste=21.439 ± 2.132

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | sim-verified (664.5 DPS) | yes | Eye of Rend (12587, -0.05 DPS) [dungeon]; Ragefury Eyepatch (11735, -2.22 DPS) [dungeon]; Mask of the Unforgiven (13404, -7.79 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 90.3 attack_power points (12.72 DPS) | yes | Beads of Ogre Might (22150, -1.67 DPS) [quest]; Mark of Fordring (15411, -3.06 DPS) [quest]; Amulet of the Darkmoon (19491, -3.22 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 132.8 attack_power points (18.72 DPS) | yes | Darkspear Pauldrons (272105, -3.63 DPS) [vendor]; Darkspear Epaulets (272106, -3.63 DPS) [vendor]; Wyrmhide Spaulders (12082, -4.87 DPS, sim-verified) [quest] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 82.4 attack_power points (11.62 DPS) | yes | Cape of the Black Baron (13340, -3.74 DPS) [dungeon]; Arcanoweave Cloak (272411, -3.95 DPS) [vendor]; Stalwart Cloak (272415, -3.95 DPS) [vendor] |
| chest | Savage Gladiator Chain (11726) | Blackrock Depths: Gorosh the Dervish [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dawn Armor (252483, +0.00 DPS) [crafted]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Obsidian Mail Tunic (22191, -1.97 DPS) [crafted] |
| wrist | Slashclaw Bracers (13211) | Blackrock Spire: Halycon [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Stalker's Bracers (19587, -0.21 DPS) [rep]; Blackmist Armguards (12966, -0.81 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -4.83 DPS, sim-verified) [world] |
| hands | Bloodmail Gauntlets (14615) | Scholomance: Lady Illucia Barov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; Stormshroud Gloves (21278, +0.00 DPS) [crafted]; Savage Gladiator Grips (11730, -5.54 DPS, sim-verified) [dungeon] |
| waist | Bloodmail Belt (14614) | Scholomance: Lady Illucia Barov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Ferocity of the Timbermaw (227805, +0.00 DPS) [vendor] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 223.2 attack_power points (31.46 DPS) | yes | Sentinel's Leather Pants (237818, -10.37 DPS) [vendor]; Plaguehound Leggings (18736, -13.68 DPS) [dungeon] |
| feet | Bloodmail Boots (14616) | Scholomance: Lady Illucia Barov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windreaver Greaves (13967, +0.00 DPS) [dungeon]; Fine Dawn Treaders (227815, -2.12 DPS) [vendor]; Savage Gladiator Greaves (11731, -6.05 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -3.20 DPS) [dungeon]; Cutthroat's Signet (272408, -3.53 DPS) [vendor]; Naglering (11669, -12.31 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.01 DPS) [dungeon]; Cutthroat's Signet (272408, -1.35 DPS) [vendor]; Naglering (11669, -9.11 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Eye of the Beast (13968, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of the Beast (13968, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -5.09 DPS) [crafted]; Counterattack Lodestone (18537, -8.89 DPS) [dungeon] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Gravestone War Axe (13983, -41.59 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Pendant of Celerity; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Savage Gladiator Chain; wrist: Slashclaw Bracers; hands: Bloodmail Gauntlets; waist: Bloodmail Belt; legs: Sentinel's Chain Leggings; feet: Bloodmail Boots; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Blackhand's Breadth; main_hand: The Unstoppable Force

No-known-source sample (15 of 1751, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

## Horde

### Band 20 (orc, 0000000000000000-254000000000000000-0000000000000000)

Set DPS (verified): 38.1. Weights run: 1.2s. Verify run: 2.1s. 205 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.578 ± 0.032, crit=1.073 ± 0.033 per rating point (14 rating = 1%, 15.018 per %), hit=1.760 ± 0.083 per rating point (10 rating = 1%, 17.598 per %), melee_haste=10.847 ± 0.323

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (0.57 DPS) | yes | Brawler's Leather Hood (252504, -0.30 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 9.5 attack_power points (0.34 DPS) | yes | Erudite's Amulet (277204, -0.11 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 7.9 attack_power points (0.28 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 9.5 attack_power points (0.34 DPS) | yes | Grave Shroud (279865, -0.01 DPS) [quest]; Dark Leather Cloak (2316, -0.03 DPS) [crafted]; Lambent Scale Cloak (4706, -0.05 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (38.1 DPS) | yes | Murloc Scale Breastplate (5781, +0.00 DPS) [crafted]; Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -0.65 DPS, sim-verified) [crafted] |
| wrist | Bristlebark Bindings (14569) | World drop [world_drop] | 8.7 attack_power points (0.31 DPS) | yes | Forest Leather Bracers (3202, -0.03 DPS) [world_drop]; Wolf Bracers (4794, -0.09 DPS) [vendor]; Ratchet Wristwraps (274742, -0.14 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (38.1 DPS) | yes | Fletcher's Gloves (7348, +0.00 DPS) [crafted]; Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Gloves of the Fang (10413, -0.82 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.64 DPS) | yes | Deviate Scale Belt (6468, -0.15 DPS) [crafted]; Ruffian Belt (5975, -0.21 DPS) [world]; Brawler's Leather Belt (252428, -1.13 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (38.1 DPS) | yes | Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -0.79 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 18.6 attack_power points (0.66 DPS) | yes | Brawler's Leather Boots (252439, -0.03 DPS) [crafted]; Defender's Leather Boots (252441, -0.31 DPS) [crafted]; Totemic Leather Boots (252442, -0.31 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 14.3 attack_power points (0.51 DPS) | yes | Demon Band (12054, -0.22 DPS) [world_drop]; Loop of Sacrifice (281673, -0.29 DPS) [quest]; Bounty Hunter's Ring (5351, -0.34 DPS) [quest] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 9.5 attack_power points (0.34 DPS) | yes | Demon Band (12054, -0.05 DPS) [world_drop]; Loop of Sacrifice (281673, -0.12 DPS) [quest]; Bounty Hunter's Ring (5351, -0.17 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | The Axe of Severing (23171) | Shadowfang Keep: Sever [dungeon] | 322.0 attack_power points (11.41 DPS) | yes | Hammerbone (270018, -0.37 DPS, sim-verified) [quest]; Smite's Mighty Hammer (7230, -0.63 DPS) [dungeon]; Forsaken Greataxe (251533, -0.67 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Bristlebark Bindings; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Signet of the Zhevra; main_hand: The Axe of Severing

No-known-source sample (15 of 205, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield; 14389 Durability Shoulderpads; 14705 Brackwater Chain Shield

### Band 30 (orc, 0000000000000000-255130030002000000-0000000000000000)

Set DPS (verified): 89.9. Weights run: 1.3s. Verify run: 2.5s. 349 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.006, agility=1.636 ± 0.037, crit=1.294 ± 0.039 per rating point (14 rating = 1%, 18.112 per %), hit=2.209 ± 0.100 per rating point (10 rating = 1%, 22.090 per %), melee_haste=10.672 ± 0.384

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (1.20 DPS) | yes | Cloudy Gustwoven Hood (277042, -0.30 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.30 DPS) [crafted]; Brawler's Leather Helm (252512, -0.38 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 14.5 attack_power points (0.73 DPS) | yes | Ghostshard Talisman (7731, -0.03 DPS) [dungeon]; Scout's Medallion (19537, -0.07 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 28.0 attack_power points (1.40 DPS) | yes | Barbaric Shoulders (5964, -0.49 DPS) [crafted]; Mantle of Thieves (2264, -0.58 DPS) [dungeon]; Bristlebark Amice (14573, -0.61 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 17.5 attack_power points (0.88 DPS) | yes | Tigerstrike Mantle (13108, -0.22 DPS) [world_drop]; Wolfmaster Cape (6314, -0.37 DPS) [dungeon]; Wildhunter Cloak (16658, -0.37 DPS) [quest] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 25.1 attack_power points (1.26 DPS) | yes | Defender's Leather Tunic (252450, -0.16 DPS) [crafted]; Brawler's Leather Armor (252490, -0.18 DPS) [crafted]; Dusky Leather Armor (7374, -0.50 DPS, sim-verified) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 17.8 attack_power points (0.89 DPS) | yes | Barbaric Bracers (18948, -0.16 DPS) [crafted]; Jurassic Wristguards (6198, -0.20 DPS) [world]; Bands of Serra'kis (6902, -0.29 DPS) [dungeon] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-verified (89.9 DPS) | yes | Wolfclaw Gloves (1978, +0.00 DPS) [dungeon]; Toughened Leather Gloves (4253, +0.00 DPS) [crafted]; Insignia Gloves (6408, +0.00 DPS) [world_drop] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 27.8 attack_power points (1.40 DPS) | yes | Defiler's Chain Girdle (20152, -0.19 DPS) [rep]; Defiler's Leather Girdle (20191, -0.19 DPS) [rep]; Skulker's Leather Belt (252520, -0.77 DPS, sim-verified) [crafted] |
| legs | Leggings of the Fang (10410) | Wailing Caverns: Lord Cobrahn [dungeon] | sim-verified (89.9 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Legguards (252516, -0.74 DPS, sim-verified) [crafted] |
| feet | Footpads of the Fang (10411) | Wailing Caverns: Lord Serpentis [dungeon] | sim-verified (89.9 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Stomping Boots (3741, +0.00 DPS) [quest]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 22.7 attack_power points (1.14 DPS) | yes | Thunderbrow Ring (13097, -0.09 DPS) [world_drop]; Band of the Fist (17694, -0.41 DPS) [quest]; Tiger Band (6749, -0.54 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 21.8 attack_power points (1.09 DPS) | yes | Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Band of the Fist (17694, -0.36 DPS) [quest]; Tiger Band (6749, -0.49 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (89.9 DPS) | yes | Corpsemaker (6687, -0.55 DPS) [dungeon]; Darkspear Raider's Reaper (272082, -1.18 DPS) [vendor]; Viscous Hammer (13045, -22.13 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Gloves of the Fang; waist: Prowler's Leather Belt; legs: Leggings of the Fang; feet: Footpads of the Fang; finger1: Ironspine's Eye; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 349, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots

### Band 40 (orc, 0000000000000000-255130030005102031-0000000000000000)

Set DPS (verified): 114.7. Weights run: 1.5s. Verify run: 1.6s. 555 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.530 ± 0.044, crit=1.550 ± 0.056 per rating point (14 rating = 1%, 21.700 per %), hit=2.133 ± 0.105 per rating point (10 rating = 1%, 21.325 per %), melee_haste=7.904 ± 0.736

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 47.7 attack_power points (3.26 DPS) | yes | White Bandit Mask (10008, -1.03 DPS, sim-verified) [crafted]; Barbaric Iron Helm (7915, -1.09 DPS) [crafted]; Hard Gold Coif (250537, -1.35 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.37 DPS) | yes | Scout's Medallion (19536, -0.22 DPS) [rep]; Ethereal Talisman (4430, -0.27 DPS) [quest]; Kaleidoscope Chain (13084, -0.40 DPS) [world_drop] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 28.8 attack_power points (1.97 DPS) | yes | Forest Tracker Epaulets (2278, -0.14 DPS) [world_drop]; Flintrock Shoulders (7755, -0.24 DPS) [dungeon]; Hard Gold Pauldrons (250539, -0.47 DPS) [crafted] |
| back | First Sergeant's Cloak (16340) | PvP rank 9 · First Sergeant · Horde [pvp] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Hawkeye's Cloak (14593, -0.31 DPS) [world_drop]; Parachute Cloak (10518, -0.61 DPS) [crafted]; Dark Hooded Cape (5257, -2.27 DPS, sim-verified) [world] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 45.1 attack_power points (3.08 DPS) | yes | Kolkar Marauder Chain (6773, -0.44 DPS) [quest]; Avenger's Armor (1488, -1.03 DPS) [dungeon]; Veteran's Silvered Chain Shirt (250518, -1.09 DPS) [crafted] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 20.1 attack_power points (1.38 DPS) | yes | Windtalker's Wristguards (19584, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Branded Leather Bracers (19508, -0.01 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 41.7 attack_power points (2.85 DPS) | yes | Scarlet Gauntlets (10331, -0.16 DPS) [dungeon]; Prowler's Leather Gloves (252524, -0.54 DPS) [crafted]; Skulker's Leather Gloves (252525, -0.57 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 31.8 attack_power points (2.17 DPS) | yes | Defiler's Leather Girdle (20192, -0.12 DPS) [rep]; Boar Champion's Belt (10768, -0.12 DPS) [dungeon]; Defiler's Chain Girdle (20153, -0.14 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.87 DPS) | yes | Firemane Leggings (13129, -0.27 DPS) [world_drop]; Orcish War Leggings (7929, -0.55 DPS) [crafted]; Legguards of the Vault (9396, -0.57 DPS) [dungeon] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Skulker's Leather Shoes (252531, -0.13 DPS) [crafted]; Excelsior Boots (4109, -0.34 DPS) [quest]; Blackforge Greaves (6423, -1.49 DPS, sim-verified) [dungeon] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 28.2 attack_power points (1.93 DPS) | yes | Thunderbrow Ring (13097, -0.52 DPS) [world_drop]; Mark of Kern (2262, -0.56 DPS) [dungeon]; Assault Band (13095, -0.56 DPS) [world_drop] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 21.8 attack_power points (1.49 DPS) | yes | Thunderbrow Ring (13097, -0.08 DPS) [world_drop]; Mark of Kern (2262, -0.12 DPS) [dungeon]; Assault Band (13095, -0.12 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pendulum of Doom (9425) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Ravager (7717, +0.00 DPS) [dungeon]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: First Sergeant's Cloak; chest: Quillward Harness; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Prowler's Leather Shoes; finger1: Legionnaire's Band; finger2: Ironspine's Eye; main_hand: Pendulum of Doom

No-known-source sample (15 of 555, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (orc, 3230000000000000-255130030005102051-0000000000000000)

Set DPS (verified): 157.4. Weights run: 1.5s. Verify run: 3.0s. 704 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.424 ± 0.049, crit=1.671 ± 0.066 per rating point (14 rating = 1%, 23.394 per %), hit=2.327 ± 0.117 per rating point (10 rating = 1%, 23.266 per %), melee_haste=9.099 ± 0.845

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 49.4 attack_power points (3.50 DPS) | yes | Embrace of the Lycan (9479, -0.10 DPS) [dungeon]; Blood Guard's Mail Helmet (220820, -0.28 DPS) [vendor]; White Bandit Mask (10008, -0.83 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 28.5 attack_power points (2.02 DPS) | yes | Woven Ivy Necklace (19159, -0.26 DPS) [quest]; Zealous Shadowshard Pendant (17772, -0.60 DPS) [quest]; Scout's Medallion (19535, -0.81 DPS) [rep] |
| shoulder | Prowler's Leather Shoulder (252534) | Leatherworking [crafted] | 34.8 attack_power points (2.47 DPS) | yes | Blood Guard's Mail Epaulets (220823, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.04 DPS) [quest]; Skulker's Leather Shoulder (252535, -0.08 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 31.9 attack_power points (2.26 DPS) | yes | Dark Hooded Cape (5257, -0.69 DPS) [world]; Blisterbane Wrap (12552, -0.75 DPS) [dungeon]; Dark Phantom Cape (13122, -0.75 DPS) [world_drop] |
| chest | Stone Guard's Mail Armor (220826) | Lady Palanseer [vendor] | sim-verified (157.4 DPS) | yes | Warbear Harness (15064, +0.00 DPS) [crafted]; Quillward Harness (10583, -0.31 DPS) [dungeon]; Mixologist's Tunic (12793, -1.80 DPS, sim-verified) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 29.4 attack_power points (2.08 DPS) | yes | Arena Bands (18711, +0.00 DPS) [world]; Forest Stalker's Bracers (19589, +0.00 DPS) [pvp] |
| hands | First Sergeant's Mail Gauntlets (220831) | PvP rank 9 · First Sergeant · Horde [vendor] | sim-verified (157.4 DPS) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; Fists of The Five Thunders (227022, +0.00 DPS) [vendor]; Raider Gloves (272100, +0.00 DPS) [vendor] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (3.26 DPS) | yes | Prowler's Leather Waistguard (252473, -0.06 DPS) [crafted]; Skulker's Leather Waistguard (252474, -0.15 DPS) [crafted]; Defiler's Chain Girdle (20151, -0.18 DPS) [rep] |
| legs | Stone Guard's Mail Legplates (220834) | Lady Palanseer [vendor] | sim-verified (157.4 DPS) | yes | Stormshroud Pants (15057, -0.04 DPS) [crafted]; Scarlet Leggings (10330, -0.38 DPS) [dungeon]; Serpentskin Leggings (8262, -2.00 DPS, sim-verified) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 41.7 attack_power points (2.95 DPS) | yes | Skulker's Leather Boots (252469, -0.08 DPS) [crafted]; Sandstalker Ankleguards (12470, -0.39 DPS) [dungeon]; Shadefiend Boots (11675, -0.43 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 43.3 attack_power points (3.07 DPS) | yes | White Bone Band (11862, -1.37 DPS) [quest]; Ironspine's Eye (7686, -1.59 DPS) [dungeon]; Thunderbrow Ring (13097, -1.63 DPS) [world_drop] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 32.8 attack_power points (2.33 DPS) | yes | White Bone Band (11862, -0.62 DPS) [quest]; Ironspine's Eye (7686, -0.85 DPS) [dungeon]; Thunderbrow Ring (13097, -0.89 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (157.4 DPS) | yes | Molten Heart of the Mountain (249470, -4.91 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (157.4 DPS) | yes | Molten Heart of the Mountain (249470, -1.61 DPS, sim-verified) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (157.4 DPS) | yes | Thorium Greatmace (250613, -0.92 DPS) [crafted]; Wildstaff (20556, -1.52 DPS) [quest]; Ragehammer (10626, -9.65 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** neck: Skibi's Pendant; shoulder: Prowler's Leather Shoulder; back: Blackveil Cape; chest: Stone Guard's Mail Armor; wrist: Deepfury Bracers; hands: First Sergeant's Mail Gauntlets; waist: Girdle of Beastial Fury; legs: Stone Guard's Mail Legplates; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 704, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (orc, 3230300000000000-255130030005102051-0520000000000000)

Set DPS (verified): 241.0. Weights run: 1.5s. Verify run: 6.2s. 1672 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.685 ± 0.065, crit=2.377 ± 0.093 per rating point (14 rating = 1%, 33.284 per %), hit=3.881 ± 0.176 per rating point (10 rating = 1%, 38.809 per %), melee_haste=11.635 ± 1.346

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Skyfury Helm (20134) | The Darkreaver Menace [quest] | 125.9 attack_power points (9.13 DPS) | yes | Eye of Rend (12587, -2.41 DPS) [dungeon]; Outlaw's Collar (279253, -2.94 DPS) [crafted]; Mask of the Unforgiven (13404, -6.83 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 64.1 attack_power points (4.65 DPS) | yes | Beads of Ogre Might (22150, -0.09 DPS) [quest]; Mark of Fordring (15411, -0.35 DPS) [quest]; Medallion of the Dawn (22659, -0.49 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | sim-verified (241.0 DPS) | yes | Champion's Mail Pauldrons (227154, +0.00 DPS) [pvp]; Warlord's Mail Pauldrons (231654, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -2.36 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 66.8 attack_power points (4.85 DPS) | yes | Deathguard's Cloak (20068, -1.77 DPS) [rep]; Windshear Cape (20691, -1.85 DPS) [world]; Cape of the Black Baron (13340, -2.15 DPS, sim-verified) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (241.0 DPS) | yes | Timbermaw Tunic (252484, -0.35 DPS) [crafted]; Legionnaire's Mail Hauberk (227157, -1.33 DPS) [pvp]; Tunic of Undead Slaying (23089, -11.91 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-verified (241.0 DPS) | yes | Slashclaw Bracers (13211, -0.25 DPS) [dungeon]; Blackmist Armguards (12966, -0.38 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -4.32 DPS, sim-verified) [world] |
| hands | Bloodmail Gauntlets (14615) | Scholomance: Lady Illucia Barov [dungeon] | sim-verified (241.0 DPS) | yes | Blood Guard's Mail Vices (227159, +0.00 DPS) [pvp]; General's Mail Vices (231655, +0.00 DPS) [vendor]; Voone's Vice Grips (13963, -6.06 DPS, sim-verified) [quest] |
| waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | 92.1 attack_power points (6.68 DPS) | yes | Ferocity of the Timbermaw (227805, -0.83 DPS) [vendor]; Marksman's Girdle (22232, -1.30 DPS) [dungeon]; Defiler's Chain Girdle (20150, -1.80 DPS) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 164.3 attack_power points (11.92 DPS) | yes | Outrider's Chain Leggings (22673, -1.96 DPS, sim-verified) [rep]; General's Mail Legguards (231658, -2.63 DPS) [vendor]; Legionnaire's Mail Legguards (227156, -3.50 DPS) [pvp] |
| feet | Bloodmail Boots (14616) | Scholomance: Lady Illucia Barov [dungeon] | sim-verified (241.0 DPS) | yes | General's Mail Greaves (231656, +0.00 DPS) [vendor]; Blood Guard's Mail Greaves (227158, -0.37 DPS) [pvp]; Windreaver Greaves (13967, -5.78 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (241.0 DPS) | yes | Tarnished Elven Ring (18500, -1.74 DPS) [dungeon]; Cutthroat's Signet (272408, -1.86 DPS) [vendor]; Naglering (11669, -7.68 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (241.0 DPS) | yes | Tarnished Elven Ring (18500, -0.37 DPS) [dungeon]; Cutthroat's Signet (272408, -0.49 DPS) [vendor]; Naglering (11669, -5.31 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (241.0 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Eye of the Beast (13968, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (241.0 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, -0.19 DPS) [quest]; Eye of the Beast (13968, -0.19 DPS) [quest] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (241.0 DPS) | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; Gravestone War Axe (13983, -16.77 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Skyfury Helm; neck: Pendant of Celerity; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Bloodmail Gauntlets; waist: Belt of Preserved Heads; legs: Sentinel's Chain Leggings; feet: Bloodmail Boots; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: The Unstoppable Force

No-known-source sample (15 of 1672, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60, raid preset (orc, 3230031000000000-255030031005102031-0530000000000000)

Set DPS (verified): 670.2. Weights run: 1.6s. Verify run: 6.3s. 1672 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.005, agility=2.391 ± 0.086, crit=3.038 ± 0.118 per rating point (14 rating = 1%, 42.531 per %), hit=5.440 ± 0.288 per rating point (10 rating = 1%, 54.403 per %), melee_haste=21.439 ± 2.132

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Skyfury Helm (20134) | The Darkreaver Menace [quest] | 156.2 attack_power points (22.02 DPS) | yes | Outlaw's Collar (279253, -5.95 DPS) [crafted]; Eye of Rend (12587, -6.00 DPS) [dungeon]; Mask of the Unforgiven (13404, -19.38 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 90.3 attack_power points (12.72 DPS) | yes | Mark of Fordring (15411, -3.06 DPS) [quest]; Amulet of the Darkmoon (19491, -3.22 DPS) [quest]; Beads of Ogre Might (22150, -4.36 DPS, sim-verified) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 132.8 attack_power points (18.72 DPS) | yes | Warlord's Mail Pauldrons (231654, +0.00 DPS) [vendor]; Champion's Mail Pauldrons (227154, -1.03 DPS) [pvp]; Wyrmhide Spaulders (12082, -5.08 DPS, sim-verified) [quest] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 82.4 attack_power points (11.62 DPS) | yes | Cape of the Black Baron (13340, -3.74 DPS) [dungeon]; Arcanoweave Cloak (272411, -3.95 DPS) [vendor]; Stalwart Cloak (272415, -3.95 DPS) [vendor] |
| chest | Savage Gladiator Chain (11726) | Blackrock Depths: Gorosh the Dervish [dungeon] | sim-verified (670.2 DPS) | yes | Legionnaire's Mail Hauberk (227157, +0.00 DPS) [pvp]; Dawn Armor (252483, +0.00 DPS) [crafted]; Timbermaw Tunic (252484, +0.00 DPS) [crafted] |
| wrist | Slashclaw Bracers (13211) | Blackrock Spire: Halycon [dungeon] | sim-verified (670.2 DPS) | yes | Forest Stalker's Bracers (19587, -0.21 DPS) [rep]; Blackmist Armguards (12966, -0.81 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -8.56 DPS, sim-verified) [world] |
| hands | Voone's Vice Grips (13963) | Warlord's Command [quest] | 130.3 attack_power points (18.37 DPS) | yes | General's Mail Vices (231655, +0.00 DPS) [vendor]; Blood Guard's Mail Vices (227159, -1.61 DPS) [pvp]; Stormshroud Gloves (21278, -4.71 DPS) [crafted] |
| waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | 121.1 attack_power points (17.07 DPS) | yes | Ferocity of the Timbermaw (227805, -3.27 DPS) [vendor]; Marksman's Girdle (22232, -4.34 DPS, sim-verified) [dungeon]; Might of the Timbermaw (19044, -5.83 DPS) [crafted] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 223.2 attack_power points (31.46 DPS) | yes | General's Mail Legguards (231658, -9.11 DPS) [vendor]; Sentinel's Leather Pants (237818, -10.37 DPS) [vendor]; Outrider's Chain Leggings (22673, -10.53 DPS, sim-verified) [rep] |
| feet | Savage Gladiator Greaves (11731) | Blackrock Depths: Anub'shiah [dungeon] | sim-verified (670.2 DPS) | yes | Windreaver Greaves (13967, +0.00 DPS) [dungeon]; Bloodmail Boots (14616, +0.00 DPS) [dungeon]; General's Mail Greaves (231656, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (670.2 DPS) | yes | Tarnished Elven Ring (18500, -3.20 DPS) [dungeon]; Cutthroat's Signet (272408, -3.53 DPS) [vendor]; Naglering (11669, -17.55 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (670.2 DPS) | yes | Tarnished Elven Ring (18500, -1.01 DPS) [dungeon]; Cutthroat's Signet (272408, -1.35 DPS) [vendor]; Naglering (11669, -14.25 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (670.2 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Eye of the Beast (13968, +0.00 DPS) [quest]; Heart of Wyrmthalak (22321, -17.85 DPS, sim-verified) [dungeon] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (670.2 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Eye of the Beast (13968, +0.00 DPS) [quest]; Counterattack Lodestone (18537, -15.05 DPS, sim-verified) [dungeon] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (670.2 DPS) | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; Seeping Willow (12969, -51.50 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Skyfury Helm; neck: Pendant of Celerity; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Savage Gladiator Chain; wrist: Slashclaw Bracers; hands: Voone's Vice Grips; waist: Belt of Preserved Heads; legs: Sentinel's Chain Leggings; feet: Savage Gladiator Greaves; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: The Unstoppable Force

No-known-source sample (15 of 1672, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

