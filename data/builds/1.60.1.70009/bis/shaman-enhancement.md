# Leveling BiS: Enhancement

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-254000000000000000-0000000000000000)

Set DPS (verified): 37.0. Weights run: 1.7s. Verify run: 2.8s. 225 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.777 ± 0.024, crit=1.088 ± 0.033 per rating point (14 rating = 1%, 15.232 per %), hit=1.705 ± 0.086 per rating point (10 rating = 1%, 17.052 per %), melee_haste=10.899 ± 0.322

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (0.56 DPS) | yes | Brawler's Leather Hood (252504, -0.43 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 4.7 attack_power points (0.16 DPS) | yes | Erudite's Amulet (277204, -0.05 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 3.9 attack_power points (0.14 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS, sim-verified) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Grave Shroud (279865, -0.02 DPS) [quest]; Dark Leather Cloak (2316, -0.06 DPS) [crafted]; Catacomb Cloak (279899, -0.07 DPS) [quest] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (37.0 DPS) | yes | Totemic Leather Armor (252435, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, +0.00 DPS) [crafted]; Defender's Leather Armor (252434, -0.84 DPS, sim-verified) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 7.1 attack_power points (0.25 DPS) | yes | Bristlebark Bindings (14569, -0.03 DPS) [world_drop]; Forest Leather Bracers (3202, -0.11 DPS) [world_drop]; Wolf Bracers (4794, -0.14 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (37.0 DPS) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Fletcher's Gloves (7348, -1.14 DPS, sim-verified) [crafted] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.63 DPS) | yes | Brawler's Leather Belt (252428, -0.24 DPS) [crafted]; Support Girdle (1215, -0.28 DPS) [world]; Ruffian Belt (5975, -1.26 DPS, sim-verified) [world] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (37.0 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Defender's Leather Pants (252445, -0.88 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 13.9 attack_power points (0.48 DPS) | yes | Feet of the Lynx (1121, -0.06 DPS) [world_drop]; Defender's Leather Boots (252441, -0.14 DPS) [crafted]; Totemic Leather Boots (252442, -0.14 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 11.1 attack_power points (0.39 DPS) | yes | Signet of the Zhevra (285330, -0.22 DPS) [world]; The 1 Ring (8350, -0.29 DPS) [world]; Ring of the Moon (12052, -0.32 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Signet of the Zhevra (285330, -0.12 DPS) [world]; The 1 Ring (8350, -0.18 DPS) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | The Axe of Severing (23171) | Shadowfang Keep: Sever [dungeon] | 322.0 attack_power points (11.22 DPS) | yes | Living Root (6631, -0.87 DPS) [dungeon]; Smite's Mighty Hammer (7230, -0.90 DPS, sim-verified) [dungeon]; Night Reaver (1318, -1.17 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Lambent Scale Cloak; chest: Blackened Defias Armor; wrist: Bravo's Armbands; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: The Axe of Severing

No-known-source sample (15 of 225, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher

### Band 30 (dwarf, 0000000000000000-255130030002000000-0000000000000000)

Set DPS (verified): 85.7. Weights run: 1.9s. Verify run: 1.8s. 366 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.006, agility=0.988 ± 0.029, crit=1.385 ± 0.042 per rating point (14 rating = 1%, 19.386 per %), hit=2.111 ± 0.101 per rating point (10 rating = 1%, 21.112 per %), melee_haste=10.824 ± 0.391

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (1.18 DPS) | yes | Cloudy Gustwoven Hood (277042, -0.30 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.30 DPS) [crafted]; Defender's Leather Hood (252447, -0.39 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.69 DPS) | yes | Kaleidoscope Chain (13084, -0.10 DPS) [world_drop]; River Pride Choker (13087, -0.30 DPS) [world_drop]; Sentinel's Medallion (19541, -0.30 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 20.9 attack_power points (1.03 DPS) | yes | Bristlebark Amice (14573, -0.44 DPS) [world_drop]; Mantle of Thieves (2264, -0.54 DPS) [dungeon]; Barbaric Shoulders (5964, -0.58 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 12.9 attack_power points (0.64 DPS) | yes | Sergeant Major's Cape (16315, -0.05 DPS) [pvp]; Wolfmaster Cape (6314, -0.14 DPS) [dungeon]; Slayer's Cape (14752, -0.24 DPS) [world_drop] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 19.9 attack_power points (0.98 DPS) | yes | Nightwalker Armor (2234, -0.09 DPS) [world]; Thick Murloc Armor (5782, -0.09 DPS) [crafted]; Defender's Leather Tunic (252450, -0.10 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 13.9 attack_power points (0.69 DPS) | yes | Bands of Serra'kis (6902, -0.10 DPS) [dungeon]; Barbaric Bracers (18948, -0.10 DPS) [crafted]; Technician's Bracers (270042, -0.19 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 19.9 attack_power points (0.98 DPS) | yes | Fletcher's Gloves (7348, -0.03 DPS) [crafted]; Toughened Leather Gloves (4253, -0.10 DPS) [crafted]; Brawler Gloves (720, -0.19 DPS) [world_drop] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.18 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Prowler's Leather Belt (252459, -0.00 DPS) [crafted]; Skulker's Leather Belt (252520, -0.15 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.28 DPS) | yes | Defender's Leather Pants (252445, -0.20 DPS) [crafted]; Barbaric Leggings (5963, -0.25 DPS) [crafted]; Brawler's Leather Legguards (252516, -1.73 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 14.9 attack_power points (0.74 DPS) | yes | Feet of the Lynx (1121, -0.05 DPS) [world_drop]; Disjointed Shoes (277226, -0.14 DPS) [quest]; Draftsman Boots (6668, -0.24 DPS) [quest] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.0 attack_power points (0.93 DPS) | yes | Ironspine's Eye (7686, -0.10 DPS) [dungeon]; Tiger Band (6749, -0.34 DPS) [quest]; Silverlaine's Family Seal (6321, -0.44 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 17.9 attack_power points (0.88 DPS) | yes | Ironspine's Eye (7686, -0.05 DPS) [dungeon]; Tiger Band (6749, -0.29 DPS) [quest]; Silverlaine's Family Seal (6321, -0.39 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (85.7 DPS) | yes | Corpsemaker (6687, -0.38 DPS) [dungeon]; Darkspear Raider's Reaper (272082, -1.00 DPS) [vendor]; Viscous Hammer (13045, -21.03 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 366, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler

### Band 40 (dwarf, 0000000000000000-255130030005102031-0000000000000000)

Set DPS (verified): 113.7. Weights run: 2.3s. Verify run: 2.0s. 592 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.123 ± 0.039, crit=1.585 ± 0.055 per rating point (14 rating = 1%, 22.195 per %), hit=2.153 ± 0.103 per rating point (10 rating = 1%, 21.532 per %), melee_haste=8.960 ± 0.729

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 48.2 attack_power points (3.25 DPS) | yes | White Bandit Mask (10008, -0.93 DPS) [crafted]; Barbaric Iron Helm (7915, -1.35 DPS) [crafted]; Hard Gold Coif (250537, -1.36 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.35 DPS) | yes | Ghostshard Talisman (7731, -0.40 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.51 DPS) [world_drop]; Sentinel's Medallion (19540, -0.52 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.4 attack_power points (1.64 DPS) | yes | Forest Tracker Epaulets (2278, -0.13 DPS) [world_drop]; Hard Gold Pauldrons (250539, -0.16 DPS) [crafted]; Flintrock Shoulders (7755, -0.21 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | sim-verified (113.7 DPS) | yes | Hawkeye's Cloak (14593, -0.33 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.40 DPS) [quest]; Dark Hooded Cape (5257, -2.89 DPS, sim-verified) [world] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 37.3 attack_power points (2.52 DPS) | yes | Kolkar Marauder Chain (6773, -0.10 DPS) [quest]; Avenger's Armor (1488, -0.50 DPS) [dungeon]; Golden Scale Cuirass (3845, -0.63 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.35 DPS) | yes | Ravager's Armguards (14770, -0.10 DPS) [world_drop]; Pugilist Bracers (4438, -0.27 DPS) [dungeon]; Yorgen Bracers (13012, -0.31 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 42.2 attack_power points (2.84 DPS) | yes | Scarlet Gauntlets (10331, -0.47 DPS) [dungeon]; Gauntlets of Divinity (7724, -0.69 DPS) [dungeon]; Prowler's Leather Gloves (252524, -0.81 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20089) | The League of Arathor [rep] | 30.2 attack_power points (2.04 DPS) | yes | Boar Champion's Belt (10768, -0.01 DPS) [dungeon]; Highlander's Leather Girdle (20116, -0.01 DPS) [rep]; Ogron's Sash (13117, -0.14 DPS) [world_drop] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.83 DPS) | yes | Firemane Leggings (13129, -0.27 DPS) [world_drop]; Orcish War Leggings (7929, -0.54 DPS) [crafted]; Legguards of the Vault (9396, -0.70 DPS) [dungeon] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 29.9 attack_power points (2.01 DPS) | yes | Blackforge Greaves (6423, -0.04 DPS) [dungeon]; Skulker's Leather Shoes (252531, -0.24 DPS) [crafted]; Excelsior Boots (4109, -0.39 DPS) [quest] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 25.0 attack_power points (1.68 DPS) | yes | Assault Band (13095, -0.34 DPS) [world_drop]; Thunderbrow Ring (13097, -0.38 DPS) [world_drop]; Ironspine's Eye (7686, -0.46 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.35 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.04 DPS) [world_drop]; Ironspine's Eye (7686, -0.13 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ravager (7717) | Scarlet Monastery: Herod [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonebiter (6830, +0.00 DPS) [quest]; Pendulum of Doom (9425, +0.00 DPS) [dungeon]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Sergeant Major's Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Chain Girdle; legs: Scarlet Leggings; feet: Prowler's Leather Shoes; finger1: Protector's Band; finger2: Mark of Kern; main_hand: Ravager

No-known-source sample (15 of 592, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 50 (dwarf, 3230000000000000-255130030005102051-0000000000000000)

Set DPS (verified): 157.3. Weights run: 2.3s. Verify run: 4.3s. 762 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.226 ± 0.045, crit=1.714 ± 0.064 per rating point (14 rating = 1%, 24.000 per %), hit=2.194 ± 0.120 per rating point (10 rating = 1%, 21.938 per %), melee_haste=9.570 ± 0.845

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Mail Helmet (223075) | Captain Dirgehammer [vendor] | sim-verified (157.3 DPS) | yes | Embrace of the Lycan (9479, +0.00 DPS) [dungeon]; Bloomsprout Headpiece (17767, -0.70 DPS) [dungeon]; Raging Berserker's Helm (7719, -4.08 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.9 attack_power points (1.83 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.42 DPS) [quest]; Sentinel's Medallion (19539, -0.79 DPS) [rep]; Ghostshard Talisman (7731, -0.84 DPS) [dungeon] |
| shoulder | Knight-Lieutenant's Mail Epaulets (223073) | Captain Dirgehammer [vendor] | 42.0 attack_power points (2.96 DPS) | yes | Failed Flying Experiment (9647, -0.69 DPS) [quest]; Skulker's Leather Shoulder (252535, -0.74 DPS) [crafted]; Prowler's Leather Shoulder (252534, -3.03 DPS, sim-verified) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 29.2 attack_power points (2.05 DPS) | yes | Sergeant Major's Cape (16336, -0.69 DPS) [pvp]; Dark Phantom Cape (13122, -0.76 DPS) [world_drop]; Dark Hooded Cape (5257, -1.31 DPS, sim-verified) [world] |
| chest | Knight's Mail Armor (223078) | Captain Dirgehammer [vendor] | sim-verified (157.3 DPS) | yes | Warbear Harness (15064, -0.28 DPS) [crafted]; Quillward Harness (10583, -0.61 DPS) [dungeon]; Mixologist's Tunic (12793, -5.18 DPS, sim-verified) [dungeon] |
| wrist | Arena Bands (18711) (or Bracers of the Stone Princess (17714)) | Arena Treasure Chest [world] | 28.0 attack_power points (1.97 DPS) | yes | Bracers of the Stone Princess (17714, +0.00 DPS) [dungeon]; Prowler's Leather Bracers (252539, -0.10 DPS) [crafted]; Deepfury Bracers (13120, -0.11 DPS) [world_drop] |
| hands | Sergeant Major's Mail Gauntlets (223076) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-verified (157.3 DPS) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; Fists of The Five Thunders (227022, +0.00 DPS) [vendor]; Raider Gloves (272100, -2.84 DPS, sim-verified) [vendor] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (3.24 DPS) | yes | Highlander's Chain Girdle (20088, -0.14 DPS) [rep]; Highlander's Leather Girdle (20115, -0.14 DPS) [rep]; Prowler's Leather Waistguard (252473, -0.23 DPS) [crafted] |
| legs | Knight's Mail Legplates (223074) | Captain Dirgehammer [vendor] | sim-verified (157.3 DPS) | yes | Stormshroud Pants (15057, +0.00 DPS) [crafted]; Serpentskin Leggings (8262, -0.17 DPS) [world_drop]; Gryphon Rider's Leggings (9652, -5.12 DPS, sim-verified) [quest] |
| feet | Sergeant Major's Mail Sabatons (223077) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-verified (157.3 DPS) | yes | Shadefiend Boots (11675, +0.00 DPS) [dungeon]; Prowler's Leather Boots (252468, +0.00 DPS) [crafted]; Skulker's Leather Boots (252469, +0.00 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 41.9 attack_power points (2.95 DPS) | yes | Mark of Kern (2262, -1.54 DPS) [dungeon]; Assault Band (13095, -1.54 DPS) [world_drop]; Thunderbrow Ring (13097, -1.57 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 31.0 attack_power points (2.18 DPS) | yes | Mark of Kern (2262, -0.78 DPS) [dungeon]; Assault Band (13095, -0.78 DPS) [world_drop]; Thunderbrow Ring (13097, -0.80 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (157.3 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (157.3 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (157.3 DPS) | yes | Thorium Greatmace (250613, -0.92 DPS) [crafted]; Darkspear Raider's Reaper (272080, -1.60 DPS) [vendor]; Ragehammer (10626, -7.96 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Knight-Lieutenant's Mail Helmet; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Mail Epaulets; back: Blackveil Cape; chest: Knight's Mail Armor; wrist: Arena Bands; hands: Sergeant Major's Mail Gauntlets; waist: Girdle of Beastial Fury; legs: Knight's Mail Legplates; feet: Sergeant Major's Mail Sabatons; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 762, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 60 (dwarf, 3230300000000000-255130030005102051-0520000000000000)

Set DPS (verified): 235.0. Weights run: 2.2s. Verify run: 6.5s. 1751 eligible items had no known source.

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

Set DPS (verified): 664.5. Weights run: 2.3s. Verify run: 9.1s. 1751 eligible items had no known source.

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

Set DPS (verified): 37.1. Weights run: 1.7s. Verify run: 2.7s. 205 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.777 ± 0.024, crit=1.088 ± 0.033 per rating point (14 rating = 1%, 15.232 per %), hit=1.705 ± 0.086 per rating point (10 rating = 1%, 17.052 per %), melee_haste=10.899 ± 0.322

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (0.56 DPS) | yes | Brawler's Leather Hood (252504, -0.41 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 4.7 attack_power points (0.16 DPS) | yes | Erudite's Amulet (277204, -0.05 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 3.9 attack_power points (0.14 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS, sim-verified) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Grave Shroud (279865, -0.02 DPS) [quest]; Dark Leather Cloak (2316, -0.06 DPS) [crafted]; Subterranean Cape (14149, -0.07 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (37.1 DPS) | yes | Totemic Leather Armor (252435, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, +0.00 DPS) [crafted]; Defender's Leather Armor (252434, -0.71 DPS, sim-verified) [crafted] |
| wrist | Bristlebark Bindings (14569) | World drop [world_drop] | 6.3 attack_power points (0.22 DPS) | yes | Forest Leather Bracers (3202, -0.09 DPS) [world_drop]; Wolf Bracers (4794, -0.11 DPS) [vendor]; Light Leather Bracers (7281, -0.12 DPS) [crafted] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (37.1 DPS) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Fletcher's Gloves (7348, -1.00 DPS, sim-verified) [crafted] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.63 DPS) | yes | Brawler's Leather Belt (252428, -0.24 DPS) [crafted]; Support Girdle (1215, -0.28 DPS) [world]; Ruffian Belt (5975, -1.15 DPS, sim-verified) [world] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (37.1 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Defender's Leather Pants (252445, -0.76 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 13.9 attack_power points (0.48 DPS) | yes | Feet of the Lynx (1121, -0.06 DPS) [world_drop]; Defender's Leather Boots (252441, -0.14 DPS) [crafted]; Totemic Leather Boots (252442, -0.14 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 11.1 attack_power points (0.39 DPS) | yes | Loop of Sacrifice (281673, -0.18 DPS) [quest]; Signet of the Zhevra (285330, -0.22 DPS) [world]; The 1 Ring (8350, -0.29 DPS) [world] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Loop of Sacrifice (281673, -0.07 DPS) [quest]; Signet of the Zhevra (285330, -0.12 DPS) [world]; The 1 Ring (8350, -0.18 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | The Axe of Severing (23171) | Shadowfang Keep: Sever [dungeon] | 322.0 attack_power points (11.22 DPS) | yes | Hammerbone (270018, -0.45 DPS, sim-verified) [quest]; Forsaken Greataxe (251533, -0.66 DPS) [quest]; Smite's Mighty Hammer (7230, -0.73 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Lambent Scale Cloak; chest: Blackened Defias Armor; wrist: Bristlebark Bindings; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: The Axe of Severing

No-known-source sample (15 of 205, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield; 14389 Durability Shoulderpads; 14705 Brackwater Chain Shield

### Band 30 (orc, 0000000000000000-255130030002000000-0000000000000000)

Set DPS (verified): 86.2. Weights run: 1.9s. Verify run: 1.7s. 349 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.006, agility=0.988 ± 0.029, crit=1.385 ± 0.042 per rating point (14 rating = 1%, 19.386 per %), hit=2.111 ± 0.101 per rating point (10 rating = 1%, 21.112 per %), melee_haste=10.824 ± 0.391

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (1.18 DPS) | yes | Cloudy Gustwoven Hood (277042, -0.30 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.30 DPS) [crafted]; Defender's Leather Hood (252447, -0.39 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.69 DPS) | yes | Kaleidoscope Chain (13084, -0.10 DPS) [world_drop]; River Pride Choker (13087, -0.30 DPS) [world_drop]; Scout's Medallion (19537, -0.30 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 20.9 attack_power points (1.03 DPS) | yes | Bristlebark Amice (14573, -0.44 DPS) [world_drop]; Mantle of Thieves (2264, -0.54 DPS) [dungeon]; Barbaric Shoulders (5964, -0.89 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 12.9 attack_power points (0.64 DPS) | yes | Wolfmaster Cape (6314, -0.14 DPS) [dungeon]; Slayer's Cape (14752, -0.24 DPS) [world_drop]; Wildhunter Cloak (16658, -0.71 DPS, sim-verified) [quest] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 19.9 attack_power points (0.98 DPS) | yes | Nightwalker Armor (2234, -0.09 DPS) [world]; Thick Murloc Armor (5782, -0.09 DPS) [crafted]; Defender's Leather Tunic (252450, -0.10 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 13.9 attack_power points (0.69 DPS) | yes | Barbaric Bracers (18948, -0.10 DPS) [crafted]; Technician's Bracers (270042, -0.19 DPS) [quest]; Bands of Serra'kis (6902, -0.59 DPS, sim-verified) [dungeon] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 19.9 attack_power points (0.98 DPS) | yes | Fletcher's Gloves (7348, -0.03 DPS) [crafted]; Toughened Leather Gloves (4253, -0.10 DPS) [crafted]; Brawler Gloves (720, -0.19 DPS) [world_drop] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.18 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Prowler's Leather Belt (252459, -0.00 DPS) [crafted]; Skulker's Leather Belt (252520, -0.15 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.28 DPS) | yes | Defender's Leather Pants (252445, -0.20 DPS) [crafted]; Barbaric Leggings (5963, -0.25 DPS) [crafted]; Brawler's Leather Legguards (252516, -1.75 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 14.9 attack_power points (0.74 DPS) | yes | Feet of the Lynx (1121, -0.05 DPS) [world_drop]; Stomping Boots (3741, -0.15 DPS) [quest]; Draftsman Boots (6668, -0.24 DPS) [quest] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.0 attack_power points (0.93 DPS) | yes | Ironspine's Eye (7686, -0.10 DPS) [dungeon]; Tiger Band (6749, -0.34 DPS) [quest]; Band of the Fist (17694, -0.35 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 17.9 attack_power points (0.88 DPS) | yes | Ironspine's Eye (7686, -0.05 DPS) [dungeon]; Tiger Band (6749, -0.29 DPS) [quest]; Band of the Fist (17694, -0.29 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (86.2 DPS) | yes | Corpsemaker (6687, -0.38 DPS) [dungeon]; Darkspear Raider's Reaper (272082, -1.00 DPS) [vendor]; Viscous Hammer (13045, -21.56 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 349, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots

### Band 40 (orc, 0000000000000000-255130030005102031-0000000000000000)

Set DPS (verified): 111.6. Weights run: 2.3s. Verify run: 2.0s. 555 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.123 ± 0.039, crit=1.585 ± 0.055 per rating point (14 rating = 1%, 22.195 per %), hit=2.153 ± 0.103 per rating point (10 rating = 1%, 21.532 per %), melee_haste=8.960 ± 0.729

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 48.2 attack_power points (3.25 DPS) | yes | White Bandit Mask (10008, -0.93 DPS) [crafted]; Barbaric Iron Helm (7915, -1.35 DPS) [crafted]; Hard Gold Coif (250537, -1.36 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.35 DPS) | yes | Ghostshard Talisman (7731, -0.40 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.51 DPS) [world_drop]; Ethereal Talisman (4430, -0.98 DPS, sim-verified) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.4 attack_power points (1.64 DPS) | yes | Forest Tracker Epaulets (2278, -0.13 DPS) [world_drop]; Hard Gold Pauldrons (250539, -0.16 DPS) [crafted]; Flintrock Shoulders (7755, -0.21 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 19.2 attack_power points (1.30 DPS) | yes | First Sergeant's Cloak (16340, +0.00 DPS, sim-verified) [pvp]; Hawkeye's Cloak (14593, -0.36 DPS) [world_drop]; Wildhunter Cloak (16658, -0.62 DPS) [quest] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 37.3 attack_power points (2.52 DPS) | yes | Kolkar Marauder Chain (6773, -0.10 DPS) [quest]; Avenger's Armor (1488, -0.50 DPS) [dungeon]; Golden Scale Cuirass (3845, -0.63 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.35 DPS) | yes | Windtalker's Wristguards (19584, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Ravager's Armguards (14770, -0.10 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 42.2 attack_power points (2.84 DPS) | yes | Scarlet Gauntlets (10331, -0.47 DPS) [dungeon]; Gauntlets of Divinity (7724, -0.69 DPS) [dungeon]; Prowler's Leather Gloves (252524, -0.81 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20153) | The Defilers [rep] | 30.2 attack_power points (2.04 DPS) | yes | Boar Champion's Belt (10768, -0.01 DPS) [dungeon]; Defiler's Leather Girdle (20192, -0.01 DPS) [rep]; Ogron's Sash (13117, -0.14 DPS) [world_drop] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.83 DPS) | yes | Firemane Leggings (13129, -0.27 DPS) [world_drop]; Orcish War Leggings (7929, -0.54 DPS) [crafted]; Legguards of the Vault (9396, -0.70 DPS) [dungeon] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 29.9 attack_power points (2.01 DPS) | yes | Skulker's Leather Shoes (252531, -0.24 DPS) [crafted]; Excelsior Boots (4109, -0.39 DPS) [quest]; Blackforge Greaves (6423, -1.80 DPS, sim-verified) [dungeon] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 25.0 attack_power points (1.68 DPS) | yes | Assault Band (13095, -0.34 DPS) [world_drop]; Thunderbrow Ring (13097, -0.38 DPS) [world_drop]; Ironspine's Eye (7686, -0.46 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.35 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.04 DPS) [world_drop]; Ironspine's Eye (7686, -0.13 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ravager (7717) | Scarlet Monastery: Herod [dungeon] | sim-verified (111.6 DPS) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor]; Pendulum of Doom (9425, -1.22 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Chain Girdle; legs: Scarlet Leggings; feet: Prowler's Leather Shoes; finger1: Legionnaire's Band; finger2: Mark of Kern; main_hand: Ravager

No-known-source sample (15 of 555, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (orc, 3230000000000000-255130030005102051-0000000000000000)

Set DPS (verified): 155.7. Weights run: 2.3s. Verify run: 2.5s. 704 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.226 ± 0.045, crit=1.714 ± 0.064 per rating point (14 rating = 1%, 24.000 per %), hit=2.194 ± 0.120 per rating point (10 rating = 1%, 21.938 per %), melee_haste=9.570 ± 0.845

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 50.0 attack_power points (3.52 DPS) | yes | Embrace of the Lycan (9479, -0.14 DPS) [dungeon]; Blood Guard's Mail Helmet (220820, -0.28 DPS) [vendor]; Bloomsprout Headpiece (17767, -0.99 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.9 attack_power points (1.83 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.42 DPS) [quest]; Ethereal Talisman (4430, -0.78 DPS) [quest]; Woven Ivy Necklace (19159, -1.77 DPS, sim-verified) [quest] |
| shoulder | Prowler's Leather Shoulder (252534) | Leatherworking [crafted] | 33.0 attack_power points (2.33 DPS) | yes | Blood Guard's Mail Epaulets (220823, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.05 DPS) [quest]; Skulker's Leather Shoulder (252535, -0.11 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 29.2 attack_power points (2.05 DPS) | yes | First Sergeant's Cloak (16340, -0.69 DPS) [pvp]; Dark Phantom Cape (13122, -0.76 DPS) [world_drop]; Dark Hooded Cape (5257, -1.36 DPS, sim-verified) [world] |
| chest | Stone Guard's Mail Armor (220826) | Lady Palanseer [vendor] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Warbear Harness (15064, -0.28 DPS) [crafted]; Quillward Harness (10583, -0.61 DPS) [dungeon]; Mixologist's Tunic (12793, -2.22 DPS, sim-verified) [dungeon] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Windtalker's Wristguards (19583, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Arena Bands (18711, -1.98 DPS, sim-verified) [world] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 47.0 attack_power points (3.31 DPS) | yes | Gloves of Holy Might (867, -0.21 DPS) [world_drop]; Fists of The Five Thunders (227022, -0.35 DPS) [vendor]; First Sergeant's Mail Gauntlets (220831, -0.49 DPS) [vendor] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (3.24 DPS) | yes | Defiler's Chain Girdle (20151, -0.14 DPS) [rep]; Defiler's Leather Girdle (20193, -0.14 DPS) [rep]; Prowler's Leather Waistguard (252473, -0.23 DPS) [crafted] |
| legs | Stone Guard's Mail Legplates (220834) | Lady Palanseer [vendor] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Serpentskin Leggings (8262, -0.17 DPS) [world_drop]; Scarlet Leggings (10330, -0.42 DPS) [dungeon]; Stormshroud Pants (15057, -2.15 DPS, sim-verified) [crafted] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 39.5 attack_power points (2.78 DPS) | yes | Skulker's Leather Boots (252469, -0.11 DPS) [crafted]; Shadefiend Boots (11675, -0.42 DPS) [dungeon]; Sandstalker Ankleguards (12470, -0.47 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 41.9 attack_power points (2.95 DPS) | yes | White Bone Band (11862, -1.26 DPS) [quest]; Mark of Kern (2262, -1.54 DPS) [dungeon]; Assault Band (13095, -1.54 DPS) [world_drop] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 31.0 attack_power points (2.18 DPS) | yes | White Bone Band (11862, -0.50 DPS) [quest]; Mark of Kern (2262, -0.78 DPS) [dungeon]; Assault Band (13095, -0.78 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+5.0 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -2.08 DPS, sim-verified) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thorium Greatmace (250613, -0.92 DPS) [crafted]; Wildstaff (20556, -1.56 DPS) [quest]; Ragehammer (10626, -3.97 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** neck: Skibi's Pendant; shoulder: Prowler's Leather Shoulder; back: Blackveil Cape; chest: Stone Guard's Mail Armor; wrist: Bracers of the Stone Princess; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Stone Guard's Mail Legplates; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 704, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (orc, 3230300000000000-255130030005102051-0520000000000000)

Set DPS (verified): 241.0. Weights run: 2.2s. Verify run: 8.3s. 1672 eligible items had no known source.

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

Set DPS (verified): 670.2. Weights run: 2.3s. Verify run: 8.4s. 1672 eligible items had no known source.

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

