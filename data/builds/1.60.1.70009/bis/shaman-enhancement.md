# Leveling BiS: Enhancement

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-254000000000000000-0000000000000000)

Set DPS (verified): 32.1. Weights run: 1.7s. Verify run: 1.1s. 225 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.751 ± 0.023, crit=1.051 ± 0.033 per rating point (14 rating = 1%, 14.720 per %), hit=1.560 ± 0.086 per rating point (10 rating = 1%, 15.601 per %), melee_haste=10.899 ± 0.322

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (0.56 DPS) | yes | Brawler's Leather Hood (252504, -0.43 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 4.5 attack_power points (0.16 DPS) | yes | Erudite's Amulet (277204, -0.05 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 3.8 attack_power points (0.13 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Grave Shroud (279865, -0.02 DPS) [quest]; Dark Leather Cloak (2316, -0.06 DPS) [crafted]; Catacomb Cloak (279899, -0.07 DPS) [quest] |
| chest | Defender's Leather Armor (252434) | Leatherworking [crafted] | 16.3 attack_power points (0.57 DPS) | yes | Brawler's Leather Armor (252490, -0.03 DPS) [crafted]; Totemic Leather Armor (252435, -0.08 DPS) [crafted]; Murloc Scale Breastplate (5781, -0.14 DPS) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 7.0 attack_power points (0.24 DPS) | yes | Bristlebark Bindings (14569, -0.03 DPS) [world_drop]; Forest Leather Bracers (3202, -0.11 DPS) [world_drop]; Wolf Bracers (4794, -0.14 DPS) [vendor] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 14.7 attack_power points (0.51 DPS) | yes | Gold-flecked Gloves (5195, -0.03 DPS) [dungeon]; Gloves of the Fang (10413, -0.08 DPS) [dungeon]; Blackened Defias Gloves (10401, -0.09 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.63 DPS) | yes | Ruffian Belt (5975, -0.21 DPS, sim-verified) [world]; Brawler's Leather Belt (252428, -0.24 DPS) [crafted]; Support Girdle (1215, -0.28 DPS) [world] |
| legs | Defender's Leather Pants (252445) | Leatherworking [crafted] | 21.0 attack_power points (0.73 DPS) | yes | Brawler's Leather Pants (252500, -0.08 DPS) [crafted]; Trapper's Leather Pants (252501, -0.08 DPS) [crafted]; Totemic Leather Pants (252446, -0.10 DPS) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 13.8 attack_power points (0.48 DPS) | yes | Feet of the Lynx (1121, -0.06 DPS) [world_drop]; Defender's Leather Boots (252441, -0.13 DPS) [crafted]; Totemic Leather Boots (252442, -0.13 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 11.0 attack_power points (0.38 DPS) | yes | Signet of the Zhevra (285330, -0.23 DPS) [world]; The 1 Ring (8350, -0.29 DPS) [world]; Ring of the Moon (12052, -0.31 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Signet of the Zhevra (285330, -0.18 DPS, sim-verified) [world]; The 1 Ring (8350, -0.18 DPS) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | The Axe of Severing (23171) | Shadowfang Keep: Sever [dungeon] | 322.0 attack_power points (11.22 DPS) | yes | Smite's Mighty Hammer (7230, -0.83 DPS, sim-verified) [dungeon]; Living Root (6631, -0.87 DPS) [dungeon]; Night Reaver (1318, -1.17 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Lambent Scale Cloak; chest: Defender's Leather Armor; wrist: Bravo's Armbands; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Defender's Leather Pants; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: The Axe of Severing

No-known-source sample (15 of 225, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher

### Band 30 (dwarf, 0000000000000000-255130030002000000-0000000000000000)

Set DPS (verified): 80.5. Weights run: 1.7s. Verify run: 1.3s. 366 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.006, agility=0.967 ± 0.029, crit=1.355 ± 0.042 per rating point (14 rating = 1%, 18.967 per %), hit=1.975 ± 0.100 per rating point (10 rating = 1%, 19.749 per %), melee_haste=10.824 ± 0.391

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (1.18 DPS) | yes | Cloudy Gustwoven Hood (277042, -0.30 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.30 DPS) [crafted]; Defender's Leather Hood (252447, -0.39 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.69 DPS) | yes | Kaleidoscope Chain (13084, -0.11 DPS) [world_drop]; River Pride Choker (13087, -0.30 DPS) [world_drop]; Sentinel's Medallion (19541, -0.31 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 20.6 attack_power points (1.02 DPS) | yes | Barbaric Shoulders (5964, -0.29 DPS) [crafted]; Bristlebark Amice (14573, -0.44 DPS) [world_drop]; Mantle of Thieves (2264, -0.54 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 12.8 attack_power points (0.63 DPS) | yes | Sergeant Major's Cape (16315, -0.04 DPS) [pvp]; Wolfmaster Cape (6314, -0.14 DPS) [dungeon]; Slayer's Cape (14752, -0.24 DPS) [world_drop] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 19.7 attack_power points (0.97 DPS) | yes | Nightwalker Armor (2234, -0.09 DPS) [world]; Thick Murloc Armor (5782, -0.09 DPS) [crafted]; Defender's Leather Tunic (252450, -0.10 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 13.8 attack_power points (0.68 DPS) | yes | Bands of Serra'kis (6902, -0.09 DPS) [dungeon]; Barbaric Bracers (18948, -0.10 DPS) [crafted]; Technician's Bracers (270042, -0.19 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 19.8 attack_power points (0.98 DPS) | yes | Toughened Leather Gloves (4253, -0.10 DPS) [crafted]; Brawler Gloves (720, -0.19 DPS) [world_drop]; Fletcher's Gloves (7348, -0.70 DPS, sim-verified) [crafted] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.18 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Prowler's Leather Belt (252459, -0.01 DPS) [crafted]; Skulker's Leather Belt (252520, -0.16 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.28 DPS) | yes | Defender's Leather Pants (252445, -0.20 DPS) [crafted]; Barbaric Leggings (5963, -0.26 DPS) [crafted]; Brawler's Leather Legguards (252516, -1.78 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 14.8 attack_power points (0.73 DPS) | yes | Feet of the Lynx (1121, -0.05 DPS) [world_drop]; Disjointed Shoes (277226, -0.14 DPS) [quest]; Draftsman Boots (6668, -0.24 DPS) [quest] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 18.9 attack_power points (0.93 DPS) | yes | Ironspine's Eye (7686, -0.11 DPS) [dungeon]; Tiger Band (6749, -0.34 DPS) [quest]; Silverlaine's Family Seal (6321, -0.44 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 17.8 attack_power points (0.88 DPS) | yes | Ironspine's Eye (7686, -0.05 DPS) [dungeon]; Tiger Band (6749, -0.29 DPS) [quest]; Silverlaine's Family Seal (6321, -0.38 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (80.5 DPS) | yes | Corpsemaker (6687, -0.37 DPS) [dungeon]; Darkspear Raider's Reaper (272082, -1.00 DPS) [vendor]; Viscous Hammer (13045, -21.46 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 366, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler

### Band 40 (dwarf, 0000000000000000-255130030005102031-0000000000000000)

Set DPS (verified): 111.5. Weights run: 2.0s. Verify run: 1.2s. 592 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.131 ± 0.039, crit=1.594 ± 0.055 per rating point (14 rating = 1%, 22.313 per %), hit=1.929 ± 0.100 per rating point (10 rating = 1%, 19.295 per %), melee_haste=9.901 ± 0.692

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 48.3 attack_power points (3.30 DPS) | yes | White Bandit Mask (10008, -0.95 DPS) [crafted]; Barbaric Iron Helm (7915, -1.38 DPS) [crafted]; Hard Gold Coif (250537, -1.39 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.37 DPS) | yes | Ghostshard Talisman (7731, -0.41 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.51 DPS) [world_drop]; Sentinel's Medallion (19540, -0.52 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.4 attack_power points (1.67 DPS) | yes | Forest Tracker Epaulets (2278, -0.14 DPS) [world_drop]; Hard Gold Pauldrons (250539, -0.17 DPS) [crafted]; Flintrock Shoulders (7755, -0.21 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | sim-verified (111.5 DPS) | yes | Hawkeye's Cloak (14593, -0.33 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.41 DPS) [quest]; Dark Hooded Cape (5257, -2.29 DPS, sim-verified) [world] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 37.5 attack_power points (2.56 DPS) | yes | Kolkar Marauder Chain (6773, -0.11 DPS) [quest]; Avenger's Armor (1488, -0.51 DPS) [dungeon]; Golden Scale Cuirass (3845, -0.65 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.37 DPS) | yes | Ravager's Armguards (14770, -0.10 DPS) [world_drop]; Pugilist Bracers (4438, -0.27 DPS) [dungeon]; Yorgen Bracers (13012, -0.31 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 42.3 attack_power points (2.89 DPS) | yes | Scarlet Gauntlets (10331, -0.48 DPS) [dungeon]; Gauntlets of Divinity (7724, -0.71 DPS) [dungeon]; Prowler's Leather Gloves (252524, -0.83 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20089) | The League of Arathor [rep] | 30.3 attack_power points (2.07 DPS) | yes | Boar Champion's Belt (10768, -0.02 DPS) [dungeon]; Highlander's Leather Girdle (20116, -0.02 DPS) [rep]; Ogron's Sash (13117, -0.15 DPS) [world_drop] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.87 DPS) | yes | Firemane Leggings (13129, -0.27 DPS) [world_drop]; Orcish War Leggings (7929, -0.55 DPS) [crafted]; Legguards of the Vault (9396, -0.71 DPS) [dungeon] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 29.9 attack_power points (2.05 DPS) | yes | Skulker's Leather Shoes (252531, -0.24 DPS) [crafted]; Excelsior Boots (4109, -0.39 DPS) [quest]; Blackforge Greaves (6423, -1.85 DPS, sim-verified) [dungeon] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 25.0 attack_power points (1.71 DPS) | yes | Assault Band (13095, -0.35 DPS) [world_drop]; Thunderbrow Ring (13097, -0.39 DPS) [world_drop]; Ironspine's Eye (7686, -0.47 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.37 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.04 DPS) [world_drop]; Ironspine's Eye (7686, -0.12 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ravager (7717) | Scarlet Monastery: Herod [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonebiter (6830, +0.00 DPS) [quest]; Pendulum of Doom (9425, +0.00 DPS) [dungeon]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Sergeant Major's Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Chain Girdle; legs: Scarlet Leggings; feet: Prowler's Leather Shoes; finger1: Protector's Band; finger2: Mark of Kern; main_hand: Ravager

No-known-source sample (15 of 592, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 50 (dwarf, 3230000000000000-255130030005102051-0000000000000000)

Set DPS (verified): 152.2. Weights run: 2.0s. Verify run: 1.5s. 762 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.199 ± 0.044, crit=1.676 ± 0.063 per rating point (14 rating = 1%, 23.463 per %), hit=2.204 ± 0.112 per rating point (10 rating = 1%, 22.043 per %), melee_haste=10.681 ± 0.826

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 49.5 attack_power points (3.52 DPS) | yes | Embrace of the Lycan (9479, +0.00 DPS, sim-verified) [dungeon]; Knight-Lieutenant's Mail Helmet (223075, -0.28 DPS) [vendor]; Bloomsprout Headpiece (17767, -0.96 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.6 attack_power points (1.82 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.40 DPS) [quest]; Sentinel's Medallion (19539, -0.80 DPS) [rep]; Ghostshard Talisman (7731, -0.82 DPS) [dungeon] |
| shoulder | Knight-Lieutenant's Mail Epaulets (223073) | Captain Dirgehammer [vendor] | 41.5 attack_power points (2.95 DPS) | yes | Prowler's Leather Shoulder (252534, -0.62 DPS) [crafted]; Failed Flying Experiment (9647, -0.67 DPS) [quest]; Skulker's Leather Shoulder (252535, -0.73 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 28.8 attack_power points (2.05 DPS) | yes | Dark Hooded Cape (5257, -0.63 DPS) [world]; Sergeant Major's Cape (16336, -0.68 DPS) [pvp]; Bloodlust Cape (14801, -0.77 DPS) [world_drop] |
| chest | Knight's Mail Armor (223078) | Captain Dirgehammer [vendor] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Warbear Harness (15064, -0.28 DPS) [crafted]; Quillward Harness (10583, -0.62 DPS) [dungeon]; Mixologist's Tunic (12793, -1.97 DPS, sim-verified) [dungeon] |
| wrist | Arena Bands (18711) (or Bracers of the Stone Princess (17714)) | Arena Treasure Chest [world] | 28.0 attack_power points (1.99 DPS) | yes | Bracers of the Stone Princess (17714, +0.00 DPS) [dungeon]; Prowler's Leather Bracers (252539, -0.11 DPS) [crafted]; Deepfury Bracers (13120, -0.14 DPS) [world_drop] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 46.4 attack_power points (3.30 DPS) | yes | Gloves of Holy Might (867, -0.21 DPS) [world_drop]; Fists of The Five Thunders (227022, -0.35 DPS) [vendor]; Sergeant Major's Mail Gauntlets (223076, -0.49 DPS) [vendor] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (3.27 DPS) | yes | Highlander's Chain Girdle (20088, -0.18 DPS) [rep]; Highlander's Leather Girdle (20115, -0.18 DPS) [rep]; Prowler's Leather Waistguard (252473, -0.26 DPS) [crafted] |
| legs | Knight's Mail Legplates (223074) | Captain Dirgehammer [vendor] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Stormshroud Pants (15057, -0.04 DPS) [crafted]; Serpentskin Leggings (8262, -0.16 DPS) [world_drop]; Gryphon Rider's Leggings (9652, -1.96 DPS, sim-verified) [quest] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 39.2 attack_power points (2.79 DPS) | yes | Skulker's Leather Boots (252469, -0.11 DPS) [crafted]; Shadefiend Boots (11675, -0.43 DPS) [dungeon]; Sandstalker Ankleguards (12470, -0.48 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 42.0 attack_power points (2.99 DPS) | yes | Mark of Kern (2262, -1.57 DPS) [dungeon]; Assault Band (13095, -1.57 DPS) [world_drop]; Thunderbrow Ring (13097, -1.60 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 30.8 attack_power points (2.19 DPS) | yes | Mark of Kern (2262, -0.77 DPS) [dungeon]; Assault Band (13095, -0.77 DPS) [world_drop]; Thunderbrow Ring (13097, -0.80 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thorium Greatmace (250613, -0.93 DPS) [crafted]; Darkspear Raider's Reaper (272080, -1.62 DPS) [vendor]; Ragehammer (10626, -7.96 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Mail Epaulets; back: Blackveil Cape; chest: Knight's Mail Armor; wrist: Arena Bands; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Knight's Mail Legplates; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 762, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 60 (dwarf, 3230300000000000-255130030005102051-0520000000000000)

Set DPS (verified): 226.0. Weights run: 2.0s. Verify run: 1.4s. 1751 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.827 ± 0.065, crit=2.563 ± 0.092 per rating point (14 rating = 1%, 35.877 per %), hit=3.902 ± 0.161 per rating point (10 rating = 1%, 39.018 per %), melee_haste=15.373 ± 1.415

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Rend (12587) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (+6.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Outlaw's Collar (279253, -0.40 DPS) [crafted]; Ragefury Eyepatch (11735, -1.02 DPS) [dungeon]; Mask of the Unforgiven (13404, -6.29 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 66.4 attack_power points (4.83 DPS) | yes | Beads of Ogre Might (22150, -0.25 DPS) [quest]; Mark of Fordring (15411, -0.33 DPS) [quest]; Medallion of the Dawn (22659, -0.48 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 102.0 attack_power points (7.41 DPS) | yes | Darkspear Pauldrons (272105, +0.00 DPS, sim-verified) [vendor]; Darkspear Epaulets (272106, -1.22 DPS) [vendor]; Wyrmhide Spaulders (12082, -1.74 DPS) [quest] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 67.0 attack_power points (4.87 DPS) | yes | Cape of the Black Baron (13340, -1.43 DPS) [dungeon]; Windshear Cape (20691, -1.72 DPS) [world]; Cloak of the Honor Guard (20073, -1.74 DPS) [rep] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Timbermaw Tunic (252484, -0.56 DPS) [crafted]; Obsidian Mail Tunic (22191, -2.22 DPS) [crafted]; Tunic of Undead Slaying (23089, -6.92 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slashclaw Bracers (13211, -0.36 DPS) [dungeon]; Blackmist Armguards (12966, -0.56 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -4.45 DPS, sim-verified) [world] |
| hands | Voone's Vice Grips (13963) | Maxwell's Mission [quest] | 94.5 attack_power points (6.86 DPS) | yes | Stormshroud Gloves (21278, +0.00 DPS, sim-verified) [crafted]; Raider Gloves (272099, -1.68 DPS) [vendor]; Timbermaw Brawlers (19049, -1.95 DPS) [crafted] |
| waist | Ferocity of the Timbermaw (227805) | Meilosh [vendor] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Marksman's Girdle (22232, -0.41 DPS) [dungeon]; Highlander's Chain Girdle (20043, -0.96 DPS) [rep]; Belt of Preserved Heads (20216, -2.66 DPS, sim-verified) [quest] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 174.7 attack_power points (12.70 DPS) | yes | Sentinel's Leather Pants (237818, -3.90 DPS) [vendor]; Plaguehound Leggings (18736, -5.88 DPS) [dungeon]; Warbear Woolies (15065, -6.24 DPS) [crafted] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 75.6 attack_power points (5.49 DPS) | yes | Bloodmail Boots (14616, +0.00 DPS, sim-verified) [dungeon]; Fine Dawn Treaders (227815, -1.19 DPS) [vendor]; Drudge Boots (21532, -1.34 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.78 DPS) [dungeon]; Cutthroat's Signet (272408, -1.91 DPS) [vendor]; Naglering (11669, -4.96 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -0.40 DPS) [dungeon]; Cutthroat's Signet (272408, -0.53 DPS) [vendor]; Naglering (11669, -2.87 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of the Beast (13968, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Blackhand's Breadth (13965, -6.08 DPS, sim-verified) [quest] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Eye of the Beast (13968, +0.00 DPS) [quest]; Counterattack Lodestone (18537, -3.93 DPS, sim-verified) [dungeon] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Gravestone War Axe (13983, -14.17 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Eye of Rend; neck: Pendant of Celerity; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Voone's Vice Grips; waist: Ferocity of the Timbermaw; legs: Sentinel's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: The Unstoppable Force

No-known-source sample (15 of 1751, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 60, raid preset (dwarf, 3230300000000000-255130030005102051-0520000000000000)

Set DPS (verified): 502.2. Weights run: 2.2s. Verify run: 1.6s. 1751 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.757 ± 0.073, crit=2.485 ± 0.103 per rating point (14 rating = 1%, 34.795 per %), hit=4.179 ± 0.216 per rating point (10 rating = 1%, 41.788 per %), melee_haste=19.127 ± 1.849

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Rend (12587) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (502.2 DPS) | yes | Outlaw's Collar (279253, -0.92 DPS) [crafted]; Ragefury Eyepatch (11735, -1.99 DPS) [dungeon]; Mask of the Unforgiven (13404, -8.47 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 68.1 attack_power points (9.68 DPS) | yes | Beads of Ogre Might (22150, -0.33 DPS) [quest]; Mark of Fordring (15411, -1.04 DPS) [quest]; Medallion of the Dawn (22659, -1.33 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 107.6 attack_power points (15.29 DPS) | yes | Darkspear Pauldrons (272105, -3.60 DPS) [vendor]; Darkspear Epaulets (272106, -3.60 DPS) [vendor]; Wyrmhide Spaulders (12082, -4.84 DPS, sim-verified) [quest] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 69.8 attack_power points (9.92 DPS) | yes | Cape of the Black Baron (13340, -3.33 DPS) [dungeon]; Cloak of the Honor Guard (20073, -3.84 DPS) [rep]; Windshear Cape (20691, -3.90 DPS) [world] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Timbermaw Tunic (252484, -1.30 DPS) [crafted]; Obsidian Mail Tunic (22191, -4.38 DPS) [crafted]; Tunic of Undead Slaying (23089, -12.61 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slashclaw Bracers (13211, -0.18 DPS) [dungeon]; Blackmist Armguards (12966, -0.51 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -7.51 DPS, sim-verified) [world] |
| hands | Voone's Vice Grips (13963) | Maxwell's Mission [quest] | 99.4 attack_power points (14.12 DPS) | yes | Stormshroud Gloves (21278, -3.24 DPS) [crafted]; Raider Gloves (272099, -4.26 DPS) [vendor]; Gauntlets of Accuracy (18349, -4.44 DPS) [dungeon] |
| waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | 96.1 attack_power points (13.66 DPS) | yes | Ferocity of the Timbermaw (227805, +0.00 DPS, sim-verified) [vendor]; Marksman's Girdle (22232, -2.48 DPS) [dungeon]; Highlander's Chain Girdle (20043, -3.89 DPS) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 172.9 attack_power points (24.57 DPS) | yes | Sentinel's Leather Pants (237818, -7.94 DPS) [vendor]; Plaguehound Leggings (18736, -11.14 DPS) [dungeon]; Warbear Woolies (15065, -12.11 DPS) [crafted] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 76.9 attack_power points (10.93 DPS) | yes | Bloodmail Boots (14616, +0.00 DPS, sim-verified) [dungeon]; Fine Dawn Treaders (227815, -2.25 DPS) [vendor]; Drudge Boots (21532, -2.99 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -3.47 DPS) [dungeon]; Cutthroat's Signet (272408, -3.72 DPS) [vendor]; Naglering (11669, -8.33 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -0.75 DPS) [dungeon]; Cutthroat's Signet (272408, -1.00 DPS) [vendor]; Naglering (11669, -4.28 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (+12.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Eye of the Beast (13968, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of the Beast (13968, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Blackhand's Breadth (13965, -4.51 DPS, sim-verified) [quest] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Seeping Willow (12969, -32.60 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Eye of Rend; neck: Pendant of Celerity; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Voone's Vice Grips; waist: Belt of Preserved Heads; legs: Sentinel's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: The Unstoppable Force

No-known-source sample (15 of 1751, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

## Horde

### Band 20 (orc, 0000000000000000-254000000000000000-0000000000000000)

Set DPS (verified): 32.4. Weights run: 1.7s. Verify run: 1.0s. 205 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.751 ± 0.023, crit=1.051 ± 0.033 per rating point (14 rating = 1%, 14.720 per %), hit=1.560 ± 0.086 per rating point (10 rating = 1%, 15.601 per %), melee_haste=10.899 ± 0.322

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (0.56 DPS) | yes | Brawler's Leather Hood (252504, -0.43 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 4.5 attack_power points (0.16 DPS) | yes | Erudite's Amulet (277204, -0.05 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 3.8 attack_power points (0.13 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Grave Shroud (279865, -0.02 DPS) [quest]; Dark Leather Cloak (2316, -0.06 DPS) [crafted]; Catacomb Cloak (279899, -0.07 DPS) [quest] |
| chest | Defender's Leather Armor (252434) | Leatherworking [crafted] | 16.3 attack_power points (0.57 DPS) | yes | Brawler's Leather Armor (252490, -0.03 DPS) [crafted]; Totemic Leather Armor (252435, -0.08 DPS) [crafted]; Murloc Scale Breastplate (5781, -0.14 DPS) [crafted] |
| wrist | Bristlebark Bindings (14569) | World drop [world_drop] | 6.3 attack_power points (0.22 DPS) | yes | Forest Leather Bracers (3202, -0.09 DPS) [world_drop]; Wolf Bracers (4794, -0.11 DPS) [vendor]; Light Leather Bracers (7281, -0.12 DPS) [crafted] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 14.7 attack_power points (0.51 DPS) | yes | Gold-flecked Gloves (5195, -0.03 DPS) [dungeon]; Gloves of the Fang (10413, -0.08 DPS) [dungeon]; Blackened Defias Gloves (10401, -0.09 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.63 DPS) | yes | Ruffian Belt (5975, -0.22 DPS, sim-verified) [world]; Brawler's Leather Belt (252428, -0.24 DPS) [crafted]; Support Girdle (1215, -0.28 DPS) [world] |
| legs | Defender's Leather Pants (252445) | Leatherworking [crafted] | 21.0 attack_power points (0.73 DPS) | yes | Brawler's Leather Pants (252500, -0.08 DPS) [crafted]; Trapper's Leather Pants (252501, -0.08 DPS) [crafted]; Totemic Leather Pants (252446, -0.10 DPS) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 13.8 attack_power points (0.48 DPS) | yes | Feet of the Lynx (1121, -0.06 DPS) [world_drop]; Defender's Leather Boots (252441, -0.13 DPS) [crafted]; Totemic Leather Boots (252442, -0.13 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 11.0 attack_power points (0.38 DPS) | yes | Loop of Sacrifice (281673, -0.17 DPS) [quest]; Signet of the Zhevra (285330, -0.23 DPS) [world]; The 1 Ring (8350, -0.29 DPS) [world] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Loop of Sacrifice (281673, -0.07 DPS) [quest]; Signet of the Zhevra (285330, -0.12 DPS) [world]; The 1 Ring (8350, -0.18 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | The Axe of Severing (23171) | Shadowfang Keep: Sever [dungeon] | 322.0 attack_power points (11.22 DPS) | yes | Forsaken Greataxe (251533, -0.66 DPS) [quest]; Smite's Mighty Hammer (7230, -0.73 DPS) [dungeon]; Hammerbone (270018, -0.81 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Lambent Scale Cloak; chest: Defender's Leather Armor; wrist: Bristlebark Bindings; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Defender's Leather Pants; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: The Axe of Severing

No-known-source sample (15 of 205, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield; 14389 Durability Shoulderpads; 14705 Brackwater Chain Shield

### Band 30 (orc, 0000000000000000-255130030002000000-0000000000000000)

Set DPS (verified): 80.6. Weights run: 1.7s. Verify run: 1.2s. 349 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.006, agility=0.967 ± 0.029, crit=1.355 ± 0.042 per rating point (14 rating = 1%, 18.967 per %), hit=1.975 ± 0.100 per rating point (10 rating = 1%, 19.749 per %), melee_haste=10.824 ± 0.391

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (1.18 DPS) | yes | Azure Gustwoven Hood (277050, -0.30 DPS) [crafted]; Defender's Leather Hood (252447, -0.39 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.45 DPS, sim-verified) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.69 DPS) | yes | Kaleidoscope Chain (13084, -0.11 DPS) [world_drop]; River Pride Choker (13087, -0.30 DPS) [world_drop]; Scout's Medallion (19537, -0.31 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 20.6 attack_power points (1.02 DPS) | yes | Barbaric Shoulders (5964, -0.29 DPS) [crafted]; Bristlebark Amice (14573, -0.44 DPS) [world_drop]; Mantle of Thieves (2264, -0.54 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 12.8 attack_power points (0.63 DPS) | yes | Wolfmaster Cape (6314, -0.14 DPS) [dungeon]; Wildhunter Cloak (16658, -0.14 DPS) [quest]; Slayer's Cape (14752, -0.24 DPS) [world_drop] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 19.7 attack_power points (0.97 DPS) | yes | Nightwalker Armor (2234, -0.09 DPS) [world]; Thick Murloc Armor (5782, -0.09 DPS) [crafted]; Defender's Leather Tunic (252450, -0.10 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 13.8 attack_power points (0.68 DPS) | yes | Bands of Serra'kis (6902, -0.09 DPS) [dungeon]; Barbaric Bracers (18948, -0.10 DPS) [crafted]; Technician's Bracers (270042, -0.19 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 19.8 attack_power points (0.98 DPS) | yes | Fletcher's Gloves (7348, -0.04 DPS) [crafted]; Toughened Leather Gloves (4253, -0.10 DPS) [crafted]; Brawler Gloves (720, -0.19 DPS) [world_drop] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.18 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Prowler's Leather Belt (252459, -0.01 DPS) [crafted]; Skulker's Leather Belt (252520, -0.16 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.28 DPS) | yes | Defender's Leather Pants (252445, -0.20 DPS) [crafted]; Barbaric Leggings (5963, -0.26 DPS) [crafted]; Brawler's Leather Legguards (252516, -1.80 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 14.8 attack_power points (0.73 DPS) | yes | Feet of the Lynx (1121, -0.05 DPS) [world_drop]; Stomping Boots (3741, -0.15 DPS) [quest]; Draftsman Boots (6668, -0.24 DPS) [quest] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 18.9 attack_power points (0.93 DPS) | yes | Ironspine's Eye (7686, -0.11 DPS) [dungeon]; Tiger Band (6749, -0.34 DPS) [quest]; Band of the Fist (17694, -0.35 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 17.8 attack_power points (0.88 DPS) | yes | Ironspine's Eye (7686, -0.05 DPS) [dungeon]; Tiger Band (6749, -0.29 DPS) [quest]; Band of the Fist (17694, -0.29 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (80.6 DPS) | yes | Corpsemaker (6687, -0.37 DPS) [dungeon]; Darkspear Raider's Reaper (272082, -1.00 DPS) [vendor]; Viscous Hammer (13045, -21.17 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 349, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots

### Band 40 (orc, 0000000000000000-255130030005102031-0000000000000000)

Set DPS (verified): 109.1. Weights run: 2.0s. Verify run: 1.1s. 555 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.131 ± 0.039, crit=1.594 ± 0.055 per rating point (14 rating = 1%, 22.313 per %), hit=1.929 ± 0.100 per rating point (10 rating = 1%, 19.295 per %), melee_haste=9.901 ± 0.692

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 48.3 attack_power points (3.30 DPS) | yes | White Bandit Mask (10008, -0.95 DPS) [crafted]; Barbaric Iron Helm (7915, -1.38 DPS) [crafted]; Hard Gold Coif (250537, -1.39 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.37 DPS) | yes | Ethereal Talisman (4430, -0.37 DPS) [quest]; Ghostshard Talisman (7731, -0.41 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.51 DPS) [world_drop] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.4 attack_power points (1.67 DPS) | yes | Forest Tracker Epaulets (2278, -0.14 DPS) [world_drop]; Hard Gold Pauldrons (250539, -0.17 DPS) [crafted]; Flintrock Shoulders (7755, -0.21 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 19.3 attack_power points (1.32 DPS) | yes | First Sergeant's Cloak (16340, -0.04 DPS) [pvp]; Hawkeye's Cloak (14593, -0.37 DPS) [world_drop]; Wildhunter Cloak (16658, -0.64 DPS) [quest] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 37.5 attack_power points (2.56 DPS) | yes | Kolkar Marauder Chain (6773, -0.11 DPS) [quest]; Avenger's Armor (1488, -0.51 DPS) [dungeon]; Golden Scale Cuirass (3845, -0.65 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.37 DPS) | yes | Windtalker's Wristguards (19584, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Ravager's Armguards (14770, -0.10 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 42.3 attack_power points (2.89 DPS) | yes | Scarlet Gauntlets (10331, -0.48 DPS) [dungeon]; Gauntlets of Divinity (7724, -0.71 DPS) [dungeon]; Prowler's Leather Gloves (252524, -0.83 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20153) | The Defilers [rep] | 30.3 attack_power points (2.07 DPS) | yes | Boar Champion's Belt (10768, -0.02 DPS) [dungeon]; Defiler's Leather Girdle (20192, -0.02 DPS) [rep]; Ogron's Sash (13117, -0.15 DPS) [world_drop] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.87 DPS) | yes | Firemane Leggings (13129, -0.27 DPS) [world_drop]; Orcish War Leggings (7929, -0.55 DPS) [crafted]; Legguards of the Vault (9396, -0.71 DPS) [dungeon] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 29.9 attack_power points (2.05 DPS) | yes | Skulker's Leather Shoes (252531, -0.24 DPS) [crafted]; Excelsior Boots (4109, -0.39 DPS) [quest]; Blackforge Greaves (6423, -2.57 DPS, sim-verified) [dungeon] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 25.0 attack_power points (1.71 DPS) | yes | Assault Band (13095, -0.35 DPS) [world_drop]; Thunderbrow Ring (13097, -0.39 DPS) [world_drop]; Ironspine's Eye (7686, -0.47 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.37 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.04 DPS) [world_drop]; Ironspine's Eye (7686, -0.12 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ravager (7717) | Scarlet Monastery: Herod [dungeon] | sim-verified (109.1 DPS) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Pendulum of Doom (9425, +0.00 DPS) [dungeon]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Chain Girdle; legs: Scarlet Leggings; feet: Prowler's Leather Shoes; finger1: Legionnaire's Band; finger2: Mark of Kern; main_hand: Ravager

No-known-source sample (15 of 555, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (orc, 3230000000000000-255130030005102051-0000000000000000)

Set DPS (verified): 156.5. Weights run: 2.0s. Verify run: 1.3s. 704 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.199 ± 0.044, crit=1.676 ± 0.063 per rating point (14 rating = 1%, 23.463 per %), hit=2.204 ± 0.112 per rating point (10 rating = 1%, 22.043 per %), melee_haste=10.681 ± 0.826

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 49.5 attack_power points (3.52 DPS) | yes | Embrace of the Lycan (9479, -0.10 DPS) [dungeon]; Blood Guard's Mail Helmet (220820, -0.28 DPS) [vendor]; Bloomsprout Headpiece (17767, -0.96 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.6 attack_power points (1.82 DPS) | yes | Woven Ivy Necklace (19159, -0.20 DPS) [quest]; Zealous Shadowshard Pendant (17772, -0.40 DPS) [quest]; Ethereal Talisman (4430, -0.77 DPS) [quest] |
| shoulder | Prowler's Leather Shoulder (252534) | Leatherworking [crafted] | 32.8 attack_power points (2.33 DPS) | yes | Blood Guard's Mail Epaulets (220823, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.06 DPS) [quest]; Skulker's Leather Shoulder (252535, -0.11 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 28.8 attack_power points (2.05 DPS) | yes | Dark Hooded Cape (5257, -0.63 DPS) [world]; First Sergeant's Cloak (16340, -0.68 DPS) [pvp]; Bloodlust Cape (14801, -0.77 DPS) [world_drop] |
| chest | Mixologist's Tunic (12793) | Blackrock Depths: Plugger Spazzring [dungeon] | 49.2 attack_power points (3.50 DPS) | yes | Stone Guard's Mail Armor (220826, +0.00 DPS, sim-verified) [vendor]; Warbear Harness (15064, -0.40 DPS) [crafted]; Quillward Harness (10583, -0.74 DPS) [dungeon] |
| wrist | Arena Bands (18711) (or Bracers of the Stone Princess (17714)) | Arena Treasure Chest [world] | 28.0 attack_power points (1.99 DPS) | yes | Bracers of the Stone Princess (17714, +0.00 DPS) [dungeon]; Windtalker's Wristguards (19583, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19589, +0.00 DPS) [pvp] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 46.4 attack_power points (3.30 DPS) | yes | Gloves of Holy Might (867, -0.21 DPS) [world_drop]; Fists of The Five Thunders (227022, -0.35 DPS) [vendor]; First Sergeant's Mail Gauntlets (220831, -0.49 DPS) [vendor] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (3.27 DPS) | yes | Defiler's Chain Girdle (20151, -0.18 DPS) [rep]; Defiler's Leather Girdle (20193, -0.18 DPS) [rep]; Prowler's Leather Waistguard (252473, -0.26 DPS) [crafted] |
| legs | Stone Guard's Mail Legplates (220834) | Lady Palanseer [vendor] | 47.5 attack_power points (3.38 DPS) | yes | Serpentskin Leggings (8262, -0.16 DPS) [world_drop]; Scarlet Leggings (10330, -0.39 DPS) [dungeon]; Stormshroud Pants (15057, -2.40 DPS, sim-verified) [crafted] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 39.2 attack_power points (2.79 DPS) | yes | Skulker's Leather Boots (252469, -0.11 DPS) [crafted]; Shadefiend Boots (11675, -0.43 DPS) [dungeon]; Sandstalker Ankleguards (12470, -0.48 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 42.0 attack_power points (2.99 DPS) | yes | White Bone Band (11862, -1.28 DPS) [quest]; Mark of Kern (2262, -1.57 DPS) [dungeon]; Assault Band (13095, -1.57 DPS) [world_drop] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 30.8 attack_power points (2.19 DPS) | yes | White Bone Band (11862, -0.48 DPS) [quest]; Mark of Kern (2262, -0.77 DPS) [dungeon]; Assault Band (13095, -0.77 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (156.5 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (156.5 DPS) | yes | Molten Heart of the Mountain (249470, -1.70 DPS, sim-verified) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (156.5 DPS) | yes | Thorium Greatmace (250613, -0.93 DPS) [crafted]; Wildstaff (20556, -1.60 DPS) [quest]; Ragehammer (10626, -8.08 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** neck: Skibi's Pendant; shoulder: Prowler's Leather Shoulder; back: Blackveil Cape; chest: Mixologist's Tunic; wrist: Arena Bands; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Stone Guard's Mail Legplates; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 704, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (orc, 3230300000000000-255130030005102051-0520000000000000)

Set DPS (verified): 223.5. Weights run: 2.0s. Verify run: 1.3s. 1672 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.827 ± 0.065, crit=2.563 ± 0.092 per rating point (14 rating = 1%, 35.877 per %), hit=3.902 ± 0.161 per rating point (10 rating = 1%, 39.018 per %), melee_haste=15.373 ± 1.415

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Skyfury Helm (20134) | The Darkreaver Menace [quest] | 133.6 attack_power points (9.71 DPS) | yes | Eye of Rend (12587, -2.61 DPS) [dungeon]; Outlaw's Collar (279253, -3.01 DPS) [crafted]; Mask of the Unforgiven (13404, -8.47 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 66.4 attack_power points (4.83 DPS) | yes | Beads of Ogre Might (22150, -0.25 DPS) [quest]; Mark of Fordring (15411, -0.33 DPS) [quest]; Medallion of the Dawn (22659, -0.48 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 102.0 attack_power points (7.41 DPS) | yes | Warlord's Mail Pauldrons (231654, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, +0.00 DPS, sim-verified) [vendor]; Champion's Mail Pauldrons (227154, -0.08 DPS) [pvp] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 67.0 attack_power points (4.87 DPS) | yes | Cape of the Black Baron (13340, -1.43 DPS) [dungeon]; Windshear Cape (20691, -1.72 DPS) [world]; Deathguard's Cloak (20068, -1.74 DPS) [rep] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Timbermaw Tunic (252484, -0.56 DPS) [crafted]; Legionnaire's Mail Hauberk (227157, -1.71 DPS) [pvp]; Tunic of Undead Slaying (23089, -7.31 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slashclaw Bracers (13211, -0.36 DPS) [dungeon]; Blackmist Armguards (12966, -0.56 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -6.07 DPS, sim-verified) [world] |
| hands | Voone's Vice Grips (13963) | Warlord's Command [quest] | 94.5 attack_power points (6.86 DPS) | yes | Blood Guard's Mail Vices (227159, +0.00 DPS) [pvp]; General's Mail Vices (231655, +0.00 DPS) [vendor]; Stormshroud Gloves (21278, -1.42 DPS) [crafted] |
| waist | Ferocity of the Timbermaw (227805) | Meilosh [vendor] | sim-verified (223.5 DPS) | yes | Marksman's Girdle (22232, -0.41 DPS) [dungeon]; Defiler's Chain Girdle (20150, -0.96 DPS) [rep]; Belt of Preserved Heads (20216, -2.27 DPS, sim-verified) [quest] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 174.7 attack_power points (12.70 DPS) | yes | General's Mail Legguards (231658, -3.18 DPS) [vendor]; Outrider's Chain Leggings (22673, -3.46 DPS, sim-verified) [rep]; Sentinel's Leather Pants (237818, -3.90 DPS) [vendor] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 75.6 attack_power points (5.49 DPS) | yes | Bloodmail Boots (14616, +0.00 DPS, sim-verified) [dungeon]; General's Mail Greaves (231656, -0.18 DPS) [vendor]; Blood Guard's Mail Greaves (227158, -0.62 DPS) [pvp] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.78 DPS) [dungeon]; Cutthroat's Signet (272408, -1.91 DPS) [vendor]; Naglering (11669, -5.35 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -0.40 DPS) [dungeon]; Cutthroat's Signet (272408, -0.53 DPS) [vendor]; Naglering (11669, -3.66 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+5.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Eye of the Beast (13968, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -2.48 DPS) [crafted]; Counterattack Lodestone (18537, -3.44 DPS) [dungeon] |
| trinket2 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of the Beast (13968, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -2.66 DPS) [crafted]; Counterattack Lodestone (18537, -3.62 DPS) [dungeon] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; Gravestone War Axe (13983, -17.81 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Skyfury Helm; neck: Pendant of Celerity; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Voone's Vice Grips; waist: Ferocity of the Timbermaw; legs: Sentinel's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket2: Blackhand's Breadth; main_hand: The Unstoppable Force

No-known-source sample (15 of 1672, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60, raid preset (orc, 3230300000000000-255130030005102051-0520000000000000)

Set DPS (verified): 519.1. Weights run: 2.2s. Verify run: 1.5s. 1672 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.757 ± 0.073, crit=2.485 ± 0.103 per rating point (14 rating = 1%, 34.795 per %), hit=4.179 ± 0.216 per rating point (10 rating = 1%, 41.788 per %), melee_haste=19.127 ± 1.849

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Skyfury Helm (20134) | The Darkreaver Menace [quest] | 130.4 attack_power points (18.53 DPS) | yes | Eye of Rend (12587, -4.94 DPS) [dungeon]; Outlaw's Collar (279253, -5.87 DPS) [crafted]; Mask of the Unforgiven (13404, -13.43 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 68.1 attack_power points (9.68 DPS) | yes | Beads of Ogre Might (22150, -0.33 DPS) [quest]; Mark of Fordring (15411, -1.04 DPS) [quest]; Medallion of the Dawn (22659, -1.33 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 107.6 attack_power points (15.29 DPS) | yes | Warlord's Mail Pauldrons (231654, +0.00 DPS) [vendor]; Champion's Mail Pauldrons (227154, -0.71 DPS) [pvp]; Wyrmhide Spaulders (12082, -5.03 DPS, sim-verified) [quest] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 69.8 attack_power points (9.92 DPS) | yes | Cape of the Black Baron (13340, -3.33 DPS) [dungeon]; Deathguard's Cloak (20068, -3.84 DPS) [rep]; Windshear Cape (20691, -3.90 DPS) [world] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Timbermaw Tunic (252484, -1.30 DPS) [crafted]; Legionnaire's Mail Hauberk (227157, -2.99 DPS) [pvp]; Tunic of Undead Slaying (23089, -13.47 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slashclaw Bracers (13211, -0.18 DPS) [dungeon]; Blackmist Armguards (12966, -0.51 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -8.90 DPS, sim-verified) [world] |
| hands | Voone's Vice Grips (13963) | Warlord's Command [quest] | 99.4 attack_power points (14.12 DPS) | yes | General's Mail Vices (231655, +0.00 DPS) [vendor]; Blood Guard's Mail Vices (227159, -0.40 DPS) [pvp]; Stormshroud Gloves (21278, -3.24 DPS) [crafted] |
| waist | Ferocity of the Timbermaw (227805) | Meilosh [vendor] | sim-verified (519.1 DPS) | yes | Marksman's Girdle (22232, -0.45 DPS) [dungeon]; Defiler's Chain Girdle (20150, -1.86 DPS) [rep]; Belt of Preserved Heads (20216, -6.46 DPS, sim-verified) [quest] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 172.9 attack_power points (24.57 DPS) | yes | Outrider's Chain Leggings (22673, -4.33 DPS, sim-verified) [rep]; General's Mail Legguards (231658, -5.73 DPS) [vendor]; Legionnaire's Mail Legguards (227156, -7.43 DPS) [pvp] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 76.9 attack_power points (10.93 DPS) | yes | Bloodmail Boots (14616, +0.00 DPS, sim-verified) [dungeon]; General's Mail Greaves (231656, -0.16 DPS) [vendor]; Blood Guard's Mail Greaves (227158, -1.01 DPS) [pvp] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -3.47 DPS) [dungeon]; Cutthroat's Signet (272408, -3.72 DPS) [vendor]; Naglering (11669, -6.35 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Naglering (11669, +0.00 DPS) [dungeon]; Tarnished Elven Ring (18500, -0.75 DPS) [dungeon]; Cutthroat's Signet (272408, -1.00 DPS) [vendor] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (+12.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Eye of the Beast (13968, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Eye of the Beast (13968, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; Gravestone War Axe (13983, -39.92 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Skyfury Helm; neck: Pendant of Celerity; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Voone's Vice Grips; waist: Ferocity of the Timbermaw; legs: Sentinel's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Hand of Justice; trinket2: Darkmoon Card: Maelstrom; main_hand: The Unstoppable Force

No-known-source sample (15 of 1672, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

