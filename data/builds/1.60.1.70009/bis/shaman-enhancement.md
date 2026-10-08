# Leveling BiS: Enhancement

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-254000000000000000-0000000000000000)

Set DPS (verified): 37.0. Weights run: 1.4s. Verify run: 2.3s. 225 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.778 ± 0.024, crit=1.089 ± 0.033 per rating point (14 rating = 1%, 15.241 per %), hit=1.682 ± 0.086 per rating point (10 rating = 1%, 16.824 per %), melee_haste=10.899 ± 0.322

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (0.56 DPS) | yes | Brawler's Leather Hood (252504, -0.43 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 4.7 attack_power points (0.16 DPS) | yes | Erudite's Amulet (277204, -0.05 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 3.9 attack_power points (0.14 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS, sim-verified) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Dark Leather Cloak (2316, -0.06 DPS) [crafted]; Catacomb Cloak (279899, -0.07 DPS) [quest]; Grave Shroud (279865, -0.20 DPS, sim-verified) [quest] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (37.0 DPS) | yes | Totemic Leather Armor (252435, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, +0.00 DPS) [crafted]; Defender's Leather Armor (252434, -0.90 DPS, sim-verified) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 7.1 attack_power points (0.25 DPS) | yes | Forest Leather Bracers (3202, -0.11 DPS) [world_drop]; Wolf Bracers (4794, -0.14 DPS) [vendor]; Bristlebark Bindings (14569, -0.19 DPS, sim-verified) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (37.0 DPS) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Fletcher's Gloves (7348, -1.20 DPS, sim-verified) [crafted] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.63 DPS) | yes | Brawler's Leather Belt (252428, -0.24 DPS) [crafted]; Support Girdle (1215, -0.28 DPS) [world]; Ruffian Belt (5975, -1.32 DPS, sim-verified) [world] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (37.0 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Defender's Leather Pants (252445, -0.95 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 13.9 attack_power points (0.48 DPS) | yes | Feet of the Lynx (1121, -0.06 DPS) [world_drop]; Defender's Leather Boots (252441, -0.14 DPS) [crafted]; Totemic Leather Boots (252442, -0.14 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 11.1 attack_power points (0.39 DPS) | yes | Signet of the Zhevra (285330, -0.22 DPS) [world]; The 1 Ring (8350, -0.29 DPS) [world]; Ring of the Moon (12052, -0.32 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Signet of the Zhevra (285330, -0.12 DPS) [world]; The 1 Ring (8350, -0.18 DPS) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | The Axe of Severing (23171) | Shadowfang Keep: Sever [dungeon] | 322.0 attack_power points (11.22 DPS) | yes | Living Root (6631, -0.87 DPS) [dungeon]; Smite's Mighty Hammer (7230, -0.97 DPS, sim-verified) [dungeon]; Night Reaver (1318, -1.17 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Lambent Scale Cloak; chest: Blackened Defias Armor; wrist: Bravo's Armbands; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: The Axe of Severing

No-known-source sample (15 of 225, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher

### Band 30 (dwarf, 0000000000000000-255130030002000000-0000000000000000)

Set DPS (verified): 86.1. Weights run: 1.5s. Verify run: 1.5s. 366 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.006, agility=0.989 ± 0.029, crit=1.385 ± 0.042 per rating point (14 rating = 1%, 19.395 per %), hit=2.098 ± 0.101 per rating point (10 rating = 1%, 20.981 per %), melee_haste=10.824 ± 0.391

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (1.18 DPS) | yes | Cloudy Gustwoven Hood (277042, -0.30 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.30 DPS) [crafted]; Defender's Leather Hood (252447, -0.39 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.69 DPS) | yes | Kaleidoscope Chain (13084, -0.10 DPS) [world_drop]; River Pride Choker (13087, -0.30 DPS) [world_drop]; Sentinel's Medallion (19541, -0.30 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 20.9 attack_power points (1.03 DPS) | yes | Bristlebark Amice (14573, -0.44 DPS) [world_drop]; Mantle of Thieves (2264, -0.54 DPS) [dungeon]; Barbaric Shoulders (5964, -0.58 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 12.9 attack_power points (0.64 DPS) | yes | Sergeant Major's Cape (16315, -0.05 DPS) [pvp]; Wolfmaster Cape (6314, -0.14 DPS) [dungeon]; Slayer's Cape (14752, -0.24 DPS) [world_drop] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 19.9 attack_power points (0.98 DPS) | yes | Nightwalker Armor (2234, -0.09 DPS) [world]; Defender's Leather Tunic (252450, -0.10 DPS) [crafted]; Thick Murloc Armor (5782, -0.79 DPS, sim-verified) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 13.9 attack_power points (0.69 DPS) | yes | Bands of Serra'kis (6902, -0.10 DPS) [dungeon]; Barbaric Bracers (18948, -0.10 DPS) [crafted]; Cultist's Armguards (270032, -0.19 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 19.9 attack_power points (0.98 DPS) | yes | Fletcher's Gloves (7348, -0.03 DPS) [crafted]; Toughened Leather Gloves (4253, -0.10 DPS) [crafted]; Heavy Earthen Gloves (7359, -0.19 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.18 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Prowler's Leather Belt (252459, -0.00 DPS) [crafted]; Skulker's Leather Belt (252520, -0.15 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.28 DPS) | yes | Defender's Leather Pants (252445, -0.20 DPS) [crafted]; Barbaric Leggings (5963, -0.25 DPS) [crafted]; Brawler's Leather Legguards (252516, -2.60 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 14.9 attack_power points (0.74 DPS) | yes | Feet of the Lynx (1121, -0.05 DPS) [world_drop]; Disjointed Shoes (277226, -0.15 DPS) [quest]; Draftsman Boots (6668, -0.24 DPS) [quest] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.0 attack_power points (0.93 DPS) | yes | Ironspine's Eye (7686, -0.10 DPS) [dungeon]; Tiger Band (6749, -0.34 DPS) [quest]; Silverlaine's Family Seal (6321, -0.44 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 17.9 attack_power points (0.88 DPS) | yes | Ironspine's Eye (7686, -0.05 DPS) [dungeon]; Tiger Band (6749, -0.29 DPS) [quest]; Silverlaine's Family Seal (6321, -0.39 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (86.1 DPS) | yes | Corpsemaker (6687, -0.38 DPS) [dungeon]; Darkspear Raider's Reaper (272082, -1.00 DPS) [vendor]; Viscous Hammer (13045, -21.46 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 366, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler

### Band 40 (dwarf, 0000000000000000-255130030005102031-0000000000000000)

Set DPS (verified): 110.8. Weights run: 1.8s. Verify run: 1.6s. 592 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.107 ± 0.038, crit=1.565 ± 0.055 per rating point (14 rating = 1%, 21.904 per %), hit=2.095 ± 0.105 per rating point (10 rating = 1%, 20.946 per %), melee_haste=9.272 ± 0.704

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 47.9 attack_power points (3.15 DPS) | yes | White Bandit Mask (10008, -0.90 DPS) [crafted]; Hard Gold Coif (250537, -1.31 DPS) [crafted]; Barbaric Iron Helm (7915, -1.31 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.32 DPS) | yes | Ghostshard Talisman (7731, -0.39 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.50 DPS) [world_drop]; Sentinel's Medallion (19540, -0.51 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.2 attack_power points (1.59 DPS) | yes | Forest Tracker Epaulets (2278, -0.13 DPS) [world_drop]; Hard Gold Pauldrons (250539, -0.14 DPS) [crafted]; Flintrock Shoulders (7755, -0.20 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | sim-verified (110.8 DPS) | yes | Hawkeye's Cloak (14593, -0.32 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.39 DPS) [quest]; Dark Hooded Cape (5257, -2.28 DPS, sim-verified) [world] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 37.0 attack_power points (2.44 DPS) | yes | Kolkar Marauder Chain (6773, -0.08 DPS) [quest]; Avenger's Armor (1488, -0.46 DPS) [dungeon]; Golden Scale Cuirass (3845, -0.59 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.32 DPS) | yes | Ravager's Armguards (14770, -0.10 DPS) [world_drop]; Pugilist Bracers (4438, -0.26 DPS) [dungeon]; Yorgen Bracers (13012, -0.31 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 41.9 attack_power points (2.76 DPS) | yes | Scarlet Gauntlets (10331, -0.45 DPS) [dungeon]; Gauntlets of Divinity (7724, -0.65 DPS) [dungeon]; Prowler's Leather Gloves (252524, -0.79 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (1.97 DPS) | yes | Boar Champion's Belt (10768, -0.00 DPS) [dungeon]; Highlander's Chain Girdle (20089, -0.01 DPS) [rep]; Ogron's Sash (13117, -0.13 DPS) [world_drop] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.76 DPS) | yes | Firemane Leggings (13129, -0.26 DPS) [world_drop]; Orcish War Leggings (7929, -0.53 DPS) [crafted]; Legguards of the Vault (9396, -0.69 DPS) [dungeon] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 29.7 attack_power points (1.96 DPS) | yes | Blackforge Greaves (6423, -0.04 DPS) [dungeon]; Skulker's Leather Shoes (252531, -0.23 DPS) [crafted]; Excelsior Boots (4109, -0.38 DPS) [quest] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 24.9 attack_power points (1.63 DPS) | yes | Assault Band (13095, -0.32 DPS) [world_drop]; Thunderbrow Ring (13097, -0.36 DPS) [world_drop]; Ironspine's Eye (7686, -0.45 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.32 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.04 DPS) [world_drop]; Ironspine's Eye (7686, -0.13 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ravager (7717) | Scarlet Monastery: Herod [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonebiter (6830, +0.00 DPS) [quest]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor]; Pendulum of Doom (9425, -2.38 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Sergeant Major's Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; legs: Scarlet Leggings; feet: Prowler's Leather Shoes; finger1: Protector's Band; finger2: Mark of Kern; main_hand: Ravager

No-known-source sample (15 of 592, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 50 (dwarf, 3230000000000000-255130030005102051-0000000000000000)

Set DPS (verified): 154.1. Weights run: 1.9s. Verify run: 3.5s. 762 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.227 ± 0.044, crit=1.718 ± 0.063 per rating point (14 rating = 1%, 24.055 per %), hit=2.126 ± 0.116 per rating point (10 rating = 1%, 21.255 per %), melee_haste=9.940 ± 0.813

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Mail Helmet (223075) | Captain Dirgehammer [vendor] | sim-verified (154.1 DPS) | yes | Embrace of the Lycan (9479, +0.00 DPS) [dungeon]; Bloomsprout Headpiece (17767, -0.70 DPS) [dungeon]; Raging Berserker's Helm (7719, -2.70 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 26.0 attack_power points (1.80 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.41 DPS) [quest]; Sentinel's Medallion (19539, -0.78 DPS) [rep]; Ghostshard Talisman (7731, -0.83 DPS) [dungeon] |
| shoulder | Knight-Lieutenant's Mail Epaulets (223073) | Captain Dirgehammer [vendor] | 42.1 attack_power points (2.92 DPS) | yes | Failed Flying Experiment (9647, -0.68 DPS) [quest]; Skulker's Leather Shoulder (252535, -0.73 DPS) [crafted]; Prowler's Leather Shoulder (252534, -2.93 DPS, sim-verified) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 29.2 attack_power points (2.03 DPS) | yes | Dark Hooded Cape (5257, -0.62 DPS) [world]; Sergeant Major's Cape (16336, -0.68 DPS) [pvp]; Dark Phantom Cape (13122, -0.75 DPS) [world_drop] |
| chest | Knight's Mail Armor (223078) | Captain Dirgehammer [vendor] | sim-verified (154.1 DPS) | yes | Warbear Harness (15064, -0.28 DPS) [crafted]; Quillward Harness (10583, -0.61 DPS) [dungeon]; Mixologist's Tunic (12793, -4.55 DPS, sim-verified) [dungeon] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | sim-verified (154.1 DPS) | yes | Arena Bands (18711, +0.00 DPS) [world]; Prowler's Leather Bracers (252539, -0.10 DPS) [crafted]; Deepfury Bracers (13120, -0.11 DPS) [world_drop] |
| hands | Sergeant Major's Mail Gauntlets (223076) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-verified (154.1 DPS) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; Fists of The Five Thunders (227022, +0.00 DPS) [vendor]; Raider Gloves (272100, -1.70 DPS, sim-verified) [vendor] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (3.20 DPS) | yes | Highlander's Chain Girdle (20088, -0.14 DPS) [rep]; Highlander's Leather Girdle (20115, -0.14 DPS) [rep]; Prowler's Leather Waistguard (252473, -0.23 DPS) [crafted] |
| legs | Knight's Mail Legplates (223074) | Captain Dirgehammer [vendor] | sim-verified (154.1 DPS) | yes | Stormshroud Pants (15057, +0.00 DPS) [crafted]; Serpentskin Leggings (8262, -0.17 DPS) [world_drop]; Gryphon Rider's Leggings (9652, -4.68 DPS, sim-verified) [quest] |
| feet | Sergeant Major's Mail Sabatons (223077) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-verified (154.1 DPS) | yes | Shadefiend Boots (11675, +0.00 DPS) [dungeon]; Prowler's Leather Boots (252468, +0.00 DPS) [crafted]; Skulker's Leather Boots (252469, +0.00 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 41.3 attack_power points (2.87 DPS) | yes | Mark of Kern (2262, -1.48 DPS) [dungeon]; Assault Band (13095, -1.48 DPS) [world_drop]; Thunderbrow Ring (13097, -1.50 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 31.0 attack_power points (2.16 DPS) | yes | Mark of Kern (2262, -0.77 DPS) [dungeon]; Assault Band (13095, -0.77 DPS) [world_drop]; Thunderbrow Ring (13097, -0.79 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (154.1 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (154.1 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (154.1 DPS) | yes | Thorium Greatmace (250613, -0.90 DPS) [crafted]; Darkspear Raider's Reaper (272080, -1.58 DPS) [vendor]; Ragehammer (10626, -5.93 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Knight-Lieutenant's Mail Helmet; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Mail Epaulets; back: Blackveil Cape; chest: Knight's Mail Armor; wrist: Bracers of the Stone Princess; hands: Sergeant Major's Mail Gauntlets; waist: Girdle of Beastial Fury; legs: Knight's Mail Legplates; feet: Sergeant Major's Mail Sabatons; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 762, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 60 (dwarf, 3230300000000000-255130030005102051-0520000000000000)

Set DPS (verified): 236.2. Weights run: 1.8s. Verify run: 6.9s. 1751 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.685 ± 0.065, crit=2.377 ± 0.093 per rating point (14 rating = 1%, 33.284 per %), hit=3.881 ± 0.176 per rating point (10 rating = 1%, 38.809 per %), melee_haste=11.635 ± 1.346

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Face of The Five Thunders (227021) | Mokvar [vendor] | sim-verified (236.2 DPS) | yes | Mask of the Unforgiven (13404, +0.00 DPS) [dungeon]; Outlaw's Collar (279253, +0.00 DPS) [crafted]; Eye of Rend (12587, -6.73 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 64.1 attack_power points (4.65 DPS) | yes | Beads of Ogre Might (22150, -0.09 DPS) [quest]; Mark of Fordring (15411, -0.35 DPS) [quest]; Medallion of the Dawn (22659, -0.49 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 101.6 attack_power points (7.37 DPS) | yes | Darkspear Epaulets (272106, -1.66 DPS) [vendor]; Wyrmhide Spaulders (12082, -1.74 DPS) [quest]; Darkspear Pauldrons (272105, -2.60 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 66.8 attack_power points (4.85 DPS) | yes | Cloak of the Honor Guard (20073, -1.77 DPS) [rep]; Windshear Cape (20691, -1.85 DPS) [world]; Cape of the Black Baron (13340, -2.95 DPS, sim-verified) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (236.2 DPS) | yes | Timbermaw Tunic (252484, -0.35 DPS) [crafted]; Obsidian Mail Tunic (22191, -1.82 DPS) [crafted]; Tunic of Undead Slaying (23089, -11.23 DPS, sim-verified) [world] |
| wrist | Bands of The Five Thunders (227017) | Mokvar [vendor] | sim-verified (236.2 DPS) | yes | Blackmist Armguards (12966, +0.00 DPS) [dungeon]; Slashclaw Bracers (13211, +0.00 DPS) [dungeon]; Forest Stalker's Bracers (19587, -4.45 DPS, sim-verified) [rep] |
| hands | Fists of The Five Thunders (227022) | Mokvar [vendor] | sim-verified (236.2 DPS) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; Bloodmail Gauntlets (14615, +0.00 DPS) [dungeon]; Stormshroud Gloves (21278, +0.00 DPS) [crafted] |
| waist | Girdle of The Five Thunders (227018) | Mokvar [vendor] | sim-verified (236.2 DPS) | yes | Bloodmail Belt (14614, +0.00 DPS) [dungeon]; Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Ferocity of the Timbermaw (227805, +0.00 DPS) [vendor] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 164.3 attack_power points (11.92 DPS) | yes | Sentinel's Leather Pants (237818, -3.79 DPS) [vendor]; Plaguehound Leggings (18736, -5.44 DPS) [dungeon]; Warbear Woolies (15065, -5.66 DPS) [crafted] |
| feet | Bloodmail Boots (14616) | Scholomance: Lady Illucia Barov [dungeon] | sim-verified (236.2 DPS) | yes | Fine Dawn Treaders (227815, -1.06 DPS) [vendor]; Drudge Boots (21532, -1.26 DPS) [quest]; Windreaver Greaves (13967, -6.65 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (236.2 DPS) | yes | Tarnished Elven Ring (18500, -1.74 DPS) [dungeon]; Cutthroat's Signet (272408, -1.86 DPS) [vendor]; Naglering (11669, -5.68 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (236.2 DPS) | yes | Tarnished Elven Ring (18500, -0.37 DPS) [dungeon]; Cutthroat's Signet (272408, -0.49 DPS) [vendor]; Naglering (11669, -3.49 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (236.2 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Eye of the Beast (13968, +0.00 DPS) [quest]; Hand of Justice (11815, -6.15 DPS, sim-verified) [dungeon] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (236.2 DPS) | yes | Eye of the Beast (13968, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Blackhand's Breadth (13965, -3.09 DPS, sim-verified) [quest] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-verified (236.2 DPS) | yes | Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Seeping Willow (12969, -18.04 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Face of The Five Thunders; neck: Pendant of Celerity; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Bands of The Five Thunders; hands: Fists of The Five Thunders; waist: Girdle of The Five Thunders; legs: Sentinel's Chain Leggings; feet: Bloodmail Boots; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Second Wind; main_hand: The Unstoppable Force

No-known-source sample (15 of 1751, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 60, raid preset (dwarf, 3230031000000000-255030031005102031-0530000000000000)

Set DPS (verified): 611.4. Weights run: 1.8s. Verify run: 7.4s. 1751 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.902 ± 0.068, crit=2.661 ± 0.096 per rating point (14 rating = 1%, 37.260 per %), hit=4.491 ± 0.260 per rating point (10 rating = 1%, 44.911 per %), melee_haste=21.838 ± 1.891

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Face of The Five Thunders (227021) | Mokvar [vendor] | sim-verified (611.4 DPS) | yes | Mask of the Unforgiven (13404, +0.00 DPS) [dungeon]; Outlaw's Collar (279253, +0.00 DPS) [crafted]; Eye of Rend (12587, -8.38 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 73.4 attack_power points (10.73 DPS) | yes | Beads of Ogre Might (22150, -0.66 DPS) [quest]; Mark of Fordring (15411, -1.49 DPS) [quest]; Medallion of the Dawn (22659, -1.78 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 113.8 attack_power points (16.63 DPS) | yes | Darkspear Pauldrons (272105, -3.68 DPS) [vendor]; Darkspear Epaulets (272106, -3.68 DPS) [vendor]; Wyrmhide Spaulders (12082, -5.04 DPS, sim-verified) [quest] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-verified (611.4 DPS) | yes | Howler's Furs (272414, +0.00 DPS) [vendor]; Arcanoweave Cloak (272411, -0.53 DPS) [vendor]; Stalwart Cloak (272415, -0.53 DPS) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (611.4 DPS) | yes | Timbermaw Tunic (252484, -2.20 DPS) [crafted]; Savage Gladiator Chain (11726, -5.56 DPS) [dungeon]; Tunic of Undead Slaying (23089, -29.53 DPS, sim-verified) [world] |
| wrist | Bands of The Five Thunders (227017) | Mokvar [vendor] | sim-verified (611.4 DPS) | yes | Blackmist Armguards (12966, +0.00 DPS) [dungeon]; Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Slashclaw Bracers (13211, -11.54 DPS, sim-verified) [dungeon] |
| hands | Fists of The Five Thunders (227022) | Mokvar [vendor] | sim-verified (611.4 DPS) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; Bloodmail Gauntlets (14615, +0.00 DPS) [dungeon]; Stormshroud Gloves (21278, +0.00 DPS) [crafted] |
| waist | Girdle of The Five Thunders (227018) | Mokvar [vendor] | sim-verified (611.4 DPS) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Ferocity of the Timbermaw (227805, +0.00 DPS) [vendor] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 186.0 attack_power points (27.19 DPS) | yes | Sentinel's Leather Pants (237818, -8.79 DPS) [vendor]; Plaguehound Leggings (18736, -12.28 DPS) [dungeon]; Blademaster Leggings (12963, -13.79 DPS) [dungeon] |
| feet | Bloodmail Boots (14616) | Scholomance: Lady Illucia Barov [dungeon] | sim-verified (611.4 DPS) | yes | Fine Dawn Treaders (227815, -2.07 DPS) [vendor]; Drudge Boots (21532, -3.17 DPS) [quest]; Windreaver Greaves (13967, -4.21 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (611.4 DPS) | yes | Tarnished Elven Ring (18500, -3.61 DPS) [dungeon]; Cutthroat's Signet (272408, -3.89 DPS) [vendor]; Naglering (11669, -19.99 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (611.4 DPS) | yes | Tarnished Elven Ring (18500, -0.83 DPS) [dungeon]; Cutthroat's Signet (272408, -1.11 DPS) [vendor]; Naglering (11669, -16.36 DPS, sim-verified) [dungeon] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (611.4 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Eye of the Beast (13968, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (611.4 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Eye of the Beast (13968, +0.00 DPS) [quest]; Counterattack Lodestone (18537, -16.23 DPS, sim-verified) [dungeon] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-verified (611.4 DPS) | yes | Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Gravestone War Axe (13983, -43.53 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Face of The Five Thunders; neck: Pendant of Celerity; shoulder: Truestrike Shoulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Bands of The Five Thunders; hands: Fists of The Five Thunders; waist: Girdle of The Five Thunders; legs: Sentinel's Chain Leggings; feet: Bloodmail Boots; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Hand of Justice; trinket2: Darkmoon Card: Maelstrom; main_hand: The Unstoppable Force

No-known-source sample (15 of 1751, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

## Horde

### Band 20 (orc, 0000000000000000-254000000000000000-0000000000000000)

Set DPS (verified): 36.9. Weights run: 1.4s. Verify run: 2.2s. 205 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.778 ± 0.024, crit=1.089 ± 0.033 per rating point (14 rating = 1%, 15.241 per %), hit=1.682 ± 0.086 per rating point (10 rating = 1%, 16.824 per %), melee_haste=10.899 ± 0.322

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (0.56 DPS) | yes | Brawler's Leather Hood (252504, -0.42 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 4.7 attack_power points (0.16 DPS) | yes | Erudite's Amulet (277204, -0.05 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 3.9 attack_power points (0.14 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS, sim-verified) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Grave Shroud (279865, -0.02 DPS) [quest]; Dark Leather Cloak (2316, -0.06 DPS) [crafted]; Catacomb Cloak (279899, -0.07 DPS) [quest] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (36.9 DPS) | yes | Totemic Leather Armor (252435, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, +0.00 DPS) [crafted]; Defender's Leather Armor (252434, -0.88 DPS, sim-verified) [crafted] |
| wrist | Bristlebark Bindings (14569) | World drop [world_drop] | 6.3 attack_power points (0.22 DPS) | yes | Forest Leather Bracers (3202, -0.09 DPS) [world_drop]; Wolf Bracers (4794, -0.11 DPS) [vendor]; Light Leather Bracers (7281, -0.12 DPS) [crafted] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (36.9 DPS) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Fletcher's Gloves (7348, -1.16 DPS, sim-verified) [crafted] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.63 DPS) | yes | Brawler's Leather Belt (252428, -0.24 DPS) [crafted]; Support Girdle (1215, -0.28 DPS) [world]; Ruffian Belt (5975, -1.31 DPS, sim-verified) [world] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (36.9 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Defender's Leather Pants (252445, -0.92 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 13.9 attack_power points (0.48 DPS) | yes | Feet of the Lynx (1121, -0.06 DPS) [world_drop]; Defender's Leather Boots (252441, -0.14 DPS) [crafted]; Totemic Leather Boots (252442, -0.14 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 11.1 attack_power points (0.39 DPS) | yes | Loop of Sacrifice (281673, -0.18 DPS) [quest]; Signet of the Zhevra (285330, -0.22 DPS) [world]; The 1 Ring (8350, -0.29 DPS) [world] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Loop of Sacrifice (281673, -0.07 DPS) [quest]; Signet of the Zhevra (285330, -0.12 DPS) [world]; The 1 Ring (8350, -0.18 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | The Axe of Severing (23171) | Shadowfang Keep: Sever [dungeon] | 322.0 attack_power points (11.22 DPS) | yes | Hammerbone (270018, -0.50 DPS, sim-verified) [quest]; Forsaken Greataxe (251533, -0.66 DPS) [quest]; Smite's Mighty Hammer (7230, -0.73 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Lambent Scale Cloak; chest: Blackened Defias Armor; wrist: Bristlebark Bindings; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: The Axe of Severing

No-known-source sample (15 of 205, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield; 14389 Durability Shoulderpads; 14705 Brackwater Chain Shield

### Band 30 (orc, 0000000000000000-255130030002000000-0000000000000000)

Set DPS (verified): 86.2. Weights run: 1.5s. Verify run: 1.4s. 349 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.006, agility=0.989 ± 0.029, crit=1.385 ± 0.042 per rating point (14 rating = 1%, 19.395 per %), hit=2.098 ± 0.101 per rating point (10 rating = 1%, 20.981 per %), melee_haste=10.824 ± 0.391

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (1.18 DPS) | yes | Cloudy Gustwoven Hood (277042, -0.30 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.30 DPS) [crafted]; Defender's Leather Hood (252447, -0.39 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.69 DPS) | yes | Kaleidoscope Chain (13084, -0.10 DPS) [world_drop]; River Pride Choker (13087, -0.30 DPS) [world_drop]; Scout's Medallion (19537, -0.30 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 20.9 attack_power points (1.03 DPS) | yes | Bristlebark Amice (14573, -0.44 DPS) [world_drop]; Mantle of Thieves (2264, -0.54 DPS) [dungeon]; Barbaric Shoulders (5964, -1.07 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 12.9 attack_power points (0.64 DPS) | yes | Wolfmaster Cape (6314, -0.14 DPS) [dungeon]; Slayer's Cape (14752, -0.24 DPS) [world_drop]; Wildhunter Cloak (16658, -0.85 DPS, sim-verified) [quest] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 19.9 attack_power points (0.98 DPS) | yes | Nightwalker Armor (2234, -0.09 DPS) [world]; Thick Murloc Armor (5782, -0.09 DPS) [crafted]; Defender's Leather Tunic (252450, -0.10 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 13.9 attack_power points (0.69 DPS) | yes | Barbaric Bracers (18948, -0.10 DPS) [crafted]; Cultist's Armguards (270032, -0.19 DPS) [quest]; Bands of Serra'kis (6902, -0.77 DPS, sim-verified) [dungeon] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 19.9 attack_power points (0.98 DPS) | yes | Fletcher's Gloves (7348, -0.03 DPS) [crafted]; Toughened Leather Gloves (4253, -0.10 DPS) [crafted]; Heavy Earthen Gloves (7359, -0.19 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.18 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Prowler's Leather Belt (252459, -0.00 DPS) [crafted]; Skulker's Leather Belt (252520, -0.15 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.28 DPS) | yes | Defender's Leather Pants (252445, -0.20 DPS) [crafted]; Barbaric Leggings (5963, -0.25 DPS) [crafted]; Brawler's Leather Legguards (252516, -1.86 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 14.9 attack_power points (0.74 DPS) | yes | Feet of the Lynx (1121, -0.05 DPS) [world_drop]; Stomping Boots (3741, -0.15 DPS) [quest]; Draftsman Boots (6668, -0.24 DPS) [quest] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.0 attack_power points (0.93 DPS) | yes | Ironspine's Eye (7686, -0.10 DPS) [dungeon]; Tiger Band (6749, -0.34 DPS) [quest]; Band of the Fist (17694, -0.35 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 17.9 attack_power points (0.88 DPS) | yes | Ironspine's Eye (7686, -0.05 DPS) [dungeon]; Tiger Band (6749, -0.29 DPS) [quest]; Band of the Fist (17694, -0.29 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (86.2 DPS) | yes | Corpsemaker (6687, -0.38 DPS) [dungeon]; Darkspear Raider's Reaper (272082, -1.00 DPS) [vendor]; Viscous Hammer (13045, -21.85 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 349, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots

### Band 40 (orc, 0000000000000000-255130030005102031-0000000000000000)

Set DPS (verified): 107.0. Weights run: 1.8s. Verify run: 1.6s. 555 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.107 ± 0.038, crit=1.565 ± 0.055 per rating point (14 rating = 1%, 21.904 per %), hit=2.095 ± 0.105 per rating point (10 rating = 1%, 20.946 per %), melee_haste=9.272 ± 0.704

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 47.9 attack_power points (3.15 DPS) | yes | White Bandit Mask (10008, -1.19 DPS, sim-verified) [crafted]; Hard Gold Coif (250537, -1.31 DPS) [crafted]; Barbaric Iron Helm (7915, -1.31 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.32 DPS) | yes | Ethereal Talisman (4430, -0.37 DPS) [quest]; Ghostshard Talisman (7731, -0.39 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.50 DPS) [world_drop] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.2 attack_power points (1.59 DPS) | yes | Forest Tracker Epaulets (2278, -0.13 DPS) [world_drop]; Hard Gold Pauldrons (250539, -0.14 DPS) [crafted]; Flintrock Shoulders (7755, -0.20 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 19.1 attack_power points (1.25 DPS) | yes | First Sergeant's Cloak (16340, -0.03 DPS) [pvp]; Hawkeye's Cloak (14593, -0.35 DPS) [world_drop]; Wildhunter Cloak (16658, -0.60 DPS) [quest] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 37.0 attack_power points (2.44 DPS) | yes | Kolkar Marauder Chain (6773, -0.08 DPS) [quest]; Avenger's Armor (1488, -0.46 DPS) [dungeon]; Golden Scale Cuirass (3845, -0.59 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.32 DPS) | yes | Windtalker's Wristguards (19584, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Ravager's Armguards (14770, -0.10 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 41.9 attack_power points (2.76 DPS) | yes | Scarlet Gauntlets (10331, -0.45 DPS) [dungeon]; Gauntlets of Divinity (7724, -0.65 DPS) [dungeon]; Prowler's Leather Gloves (252524, -0.79 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (1.97 DPS) | yes | Boar Champion's Belt (10768, -0.00 DPS) [dungeon]; Defiler's Chain Girdle (20153, -0.01 DPS) [rep]; Ogron's Sash (13117, -0.13 DPS) [world_drop] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.76 DPS) | yes | Firemane Leggings (13129, -0.26 DPS) [world_drop]; Orcish War Leggings (7929, -0.53 DPS) [crafted]; Legguards of the Vault (9396, -0.69 DPS) [dungeon] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 29.7 attack_power points (1.96 DPS) | yes | Skulker's Leather Shoes (252531, -0.23 DPS) [crafted]; Excelsior Boots (4109, -0.38 DPS) [quest]; Blackforge Greaves (6423, -1.82 DPS, sim-verified) [dungeon] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 24.9 attack_power points (1.63 DPS) | yes | Assault Band (13095, -0.32 DPS) [world_drop]; Thunderbrow Ring (13097, -0.36 DPS) [world_drop]; Ironspine's Eye (7686, -0.45 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.32 DPS) | yes | Thunderbrow Ring (13097, -0.04 DPS) [world_drop]; Ironspine's Eye (7686, -0.13 DPS) [dungeon]; Assault Band (13095, -1.19 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pendulum of Doom (9425) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-verified (107.0 DPS) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Ravager (7717, +0.00 DPS) [dungeon]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; legs: Scarlet Leggings; feet: Prowler's Leather Shoes; finger1: Legionnaire's Band; finger2: Mark of Kern; main_hand: Pendulum of Doom

No-known-source sample (15 of 555, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (orc, 3230000000000000-255130030005102051-0000000000000000)

Set DPS (verified): 153.3. Weights run: 1.9s. Verify run: 2.2s. 704 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.227 ± 0.044, crit=1.718 ± 0.063 per rating point (14 rating = 1%, 24.055 per %), hit=2.126 ± 0.116 per rating point (10 rating = 1%, 21.255 per %), melee_haste=9.940 ± 0.813

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 50.1 attack_power points (3.48 DPS) | yes | Embrace of the Lycan (9479, -0.14 DPS) [dungeon]; Blood Guard's Mail Helmet (220820, -0.28 DPS) [vendor]; Bloomsprout Headpiece (17767, -0.98 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 26.0 attack_power points (1.80 DPS) | yes | Woven Ivy Necklace (19159, -0.20 DPS) [quest]; Zealous Shadowshard Pendant (17772, -0.41 DPS) [quest]; Ethereal Talisman (4430, -0.77 DPS) [quest] |
| shoulder | Prowler's Leather Shoulder (252534) | Leatherworking [crafted] | 33.0 attack_power points (2.30 DPS) | yes | Blood Guard's Mail Epaulets (220823, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.05 DPS) [quest]; Skulker's Leather Shoulder (252535, -0.11 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 29.2 attack_power points (2.03 DPS) | yes | Dark Hooded Cape (5257, -0.62 DPS) [world]; First Sergeant's Cloak (16340, -0.68 DPS) [pvp]; Dark Phantom Cape (13122, -0.75 DPS) [world_drop] |
| chest | Stone Guard's Mail Armor (220826) | Lady Palanseer [vendor] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Warbear Harness (15064, -0.28 DPS) [crafted]; Quillward Harness (10583, -0.61 DPS) [dungeon]; Mixologist's Tunic (12793, -2.17 DPS, sim-verified) [dungeon] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Windtalker's Wristguards (19583, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Arena Bands (18711, -2.07 DPS, sim-verified) [world] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 47.0 attack_power points (3.27 DPS) | yes | Gloves of Holy Might (867, -0.20 DPS) [world_drop]; Fists of The Five Thunders (227022, -0.34 DPS) [vendor]; First Sergeant's Mail Gauntlets (220831, -0.48 DPS) [vendor] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (3.20 DPS) | yes | Defiler's Chain Girdle (20151, -0.14 DPS) [rep]; Defiler's Leather Girdle (20193, -0.14 DPS) [rep]; Prowler's Leather Waistguard (252473, -0.23 DPS) [crafted] |
| legs | Stone Guard's Mail Legplates (220834) | Lady Palanseer [vendor] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Serpentskin Leggings (8262, -0.17 DPS) [world_drop]; Scarlet Leggings (10330, -0.42 DPS) [dungeon]; Stormshroud Pants (15057, -2.66 DPS, sim-verified) [crafted] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 39.5 attack_power points (2.74 DPS) | yes | Skulker's Leather Boots (252469, -0.11 DPS) [crafted]; Shadefiend Boots (11675, -0.42 DPS) [dungeon]; Sandstalker Ankleguards (12470, -0.46 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 41.3 attack_power points (2.87 DPS) | yes | White Bone Band (11862, -1.20 DPS) [quest]; Mark of Kern (2262, -1.48 DPS) [dungeon]; Assault Band (13095, -1.48 DPS) [world_drop] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 31.0 attack_power points (2.16 DPS) | yes | White Bone Band (11862, -0.49 DPS) [quest]; Mark of Kern (2262, -0.77 DPS) [dungeon]; Assault Band (13095, -0.77 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+4.8 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -1.70 DPS, sim-verified) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thorium Greatmace (250613, -0.90 DPS) [crafted]; Darkspear Raider's Reaper (272080, -1.58 DPS) [vendor]; Ragehammer (10626, -3.92 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** neck: Skibi's Pendant; shoulder: Prowler's Leather Shoulder; back: Blackveil Cape; chest: Stone Guard's Mail Armor; wrist: Bracers of the Stone Princess; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Stone Guard's Mail Legplates; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 704, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (orc, 3230300000000000-255130030005102051-0520000000000000)

Set DPS (verified): 239.1. Weights run: 1.8s. Verify run: 6.7s. 1672 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.685 ± 0.065, crit=2.377 ± 0.093 per rating point (14 rating = 1%, 33.284 per %), hit=3.881 ± 0.176 per rating point (10 rating = 1%, 38.809 per %), melee_haste=11.635 ± 1.346

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Skyfury Helm (20134) | The Darkreaver Menace [quest] | 125.9 attack_power points (9.13 DPS) | yes | Eye of Rend (12587, -2.41 DPS) [dungeon]; Outlaw's Collar (279253, -2.94 DPS) [crafted]; Mask of the Unforgiven (13404, -6.79 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 64.1 attack_power points (4.65 DPS) | yes | Beads of Ogre Might (22150, -0.09 DPS) [quest]; Mark of Fordring (15411, -0.35 DPS) [quest]; Medallion of the Dawn (22659, -0.49 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | sim-verified (239.1 DPS) | yes | Champion's Mail Pauldrons (227154, +0.00 DPS) [pvp]; Warlord's Mail Pauldrons (231654, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -2.02 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 66.8 attack_power points (4.85 DPS) | yes | Deathguard's Cloak (20068, -1.77 DPS) [rep]; Windshear Cape (20691, -1.85 DPS) [world]; Cape of the Black Baron (13340, -2.44 DPS, sim-verified) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (239.1 DPS) | yes | Timbermaw Tunic (252484, -0.35 DPS) [crafted]; Warlord's Mail Hauberk (231653, -1.33 DPS) [vendor]; Tunic of Undead Slaying (23089, -12.56 DPS, sim-verified) [world] |
| wrist | Bands of The Five Thunders (227017) | Mokvar [vendor] | sim-verified (239.1 DPS) | yes | Blackmist Armguards (12966, +0.00 DPS) [dungeon]; Slashclaw Bracers (13211, +0.00 DPS) [dungeon]; Forest Stalker's Bracers (19587, -2.00 DPS, sim-verified) [rep] |
| hands | Fists of The Five Thunders (227022) | Mokvar [vendor] | sim-verified (239.1 DPS) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; Bloodmail Gauntlets (14615, +0.00 DPS) [dungeon]; General's Mail Vices (231655, +0.00 DPS) [vendor] |
| waist | Girdle of The Five Thunders (227018) | Mokvar [vendor] | sim-verified (239.1 DPS) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Ferocity of the Timbermaw (227805, +0.00 DPS) [vendor]; Bloodmail Belt (14614, -2.15 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 164.3 attack_power points (11.92 DPS) | yes | General's Mail Legguards (231658, -2.63 DPS) [vendor]; Outrider's Chain Leggings (22673, -3.34 DPS, sim-verified) [rep]; Legionnaire's Mail Legguards (227156, -3.50 DPS) [pvp] |
| feet | Bloodmail Boots (14616) | Scholomance: Lady Illucia Barov [dungeon] | sim-verified (239.1 DPS) | yes | General's Mail Greaves (231656, +0.00 DPS) [vendor]; Blood Guard's Mail Greaves (227158, -0.37 DPS) [pvp]; Windreaver Greaves (13967, -2.01 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (239.1 DPS) | yes | Tarnished Elven Ring (18500, -1.74 DPS) [dungeon]; Cutthroat's Signet (272408, -1.86 DPS) [vendor]; Naglering (11669, -8.87 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (239.1 DPS) | yes | Tarnished Elven Ring (18500, -0.37 DPS) [dungeon]; Cutthroat's Signet (272408, -0.49 DPS) [vendor]; Naglering (11669, -7.05 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (239.1 DPS) | yes | Eye of the Beast (13968, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Blackhand's Breadth (13965, -6.97 DPS, sim-verified) [quest] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (239.1 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, -0.19 DPS) [quest]; Eye of the Beast (13968, -0.19 DPS) [quest] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (239.1 DPS) | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; Gravestone War Axe (13983, -17.76 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Skyfury Helm; neck: Pendant of Celerity; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Bands of The Five Thunders; hands: Fists of The Five Thunders; waist: Girdle of The Five Thunders; legs: Sentinel's Chain Leggings; feet: Bloodmail Boots; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: The Unstoppable Force

No-known-source sample (15 of 1672, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60, raid preset (orc, 3230031000000000-255030031005102031-0530000000000000)

Set DPS (verified): 610.3. Weights run: 1.8s. Verify run: 7.0s. 1672 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.902 ± 0.068, crit=2.661 ± 0.096 per rating point (14 rating = 1%, 37.260 per %), hit=4.491 ± 0.260 per rating point (10 rating = 1%, 44.911 per %), melee_haste=21.838 ± 1.891

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Skyfury Helm (20134) | The Darkreaver Menace [quest] | 137.8 attack_power points (20.14 DPS) | yes | Eye of Rend (12587, -5.45 DPS) [dungeon]; Outlaw's Collar (279253, -6.13 DPS) [crafted]; Mask of the Unforgiven (13404, -8.46 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 73.4 attack_power points (10.73 DPS) | yes | Beads of Ogre Might (22150, -0.66 DPS) [quest]; Mark of Fordring (15411, -1.49 DPS) [quest]; Medallion of the Dawn (22659, -1.78 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 113.8 attack_power points (16.63 DPS) | yes | Warlord's Mail Pauldrons (231654, +0.00 DPS) [vendor]; Champion's Mail Pauldrons (227154, -0.83 DPS) [pvp]; Wyrmhide Spaulders (12082, -5.13 DPS, sim-verified) [quest] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 72.9 attack_power points (10.66 DPS) | yes | Cape of the Black Baron (13340, -3.56 DPS) [dungeon]; Arcanoweave Cloak (272411, -4.09 DPS) [vendor]; Stalwart Cloak (272415, -4.09 DPS) [vendor] |
| chest | Savage Gladiator Chain (11726) | Blackrock Depths: Gorosh the Dervish [dungeon] | sim-verified (610.3 DPS) | yes | Legionnaire's Mail Hauberk (227157, +0.00 DPS) [pvp]; Dawn Armor (252483, +0.00 DPS) [crafted]; Timbermaw Tunic (252484, +0.00 DPS) [crafted] |
| wrist | Slashclaw Bracers (13211) | Blackrock Spire: Halycon [dungeon] | sim-verified (610.3 DPS) | yes | Forest Stalker's Bracers (19587, -0.01 DPS) [rep]; Blackmist Armguards (12966, -0.48 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -5.44 DPS, sim-verified) [world] |
| hands | Bloodmail Gauntlets (14615) | Scholomance: Lady Illucia Barov [dungeon] | sim-verified (610.3 DPS) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; General's Mail Vices (231655, +0.00 DPS) [vendor]; Savage Gladiator Grips (11730, -7.14 DPS, sim-verified) [dungeon] |
| waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | 101.4 attack_power points (14.83 DPS) | yes | Ferocity of the Timbermaw (227805, -2.50 DPS) [vendor]; Marksman's Girdle (22232, -4.29 DPS, sim-verified) [dungeon]; Defiler's Chain Girdle (20150, -4.41 DPS) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 186.0 attack_power points (27.19 DPS) | yes | Outrider's Chain Leggings (22673, -4.61 DPS, sim-verified) [rep]; General's Mail Legguards (231658, -6.99 DPS) [vendor]; Legionnaire's Mail Legguards (227156, -8.75 DPS) [pvp] |
| feet | Bloodmail Boots (14616) | Scholomance: Lady Illucia Barov [dungeon] | sim-verified (610.3 DPS) | yes | Windreaver Greaves (13967, +0.00 DPS) [dungeon]; General's Mail Greaves (231656, -0.16 DPS) [vendor]; Savage Gladiator Greaves (11731, -5.75 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (610.3 DPS) | yes | Tarnished Elven Ring (18500, -3.61 DPS) [dungeon]; Cutthroat's Signet (272408, -3.89 DPS) [vendor]; Naglering (11669, -11.00 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (610.3 DPS) | yes | Tarnished Elven Ring (18500, -0.83 DPS) [dungeon]; Cutthroat's Signet (272408, -1.11 DPS) [vendor]; Naglering (11669, -7.34 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (610.3 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Eye of the Beast (13968, +0.00 DPS) [quest]; Heart of Wyrmthalak (22321, -8.46 DPS, sim-verified) [dungeon] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (610.3 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Eye of the Beast (13968, +0.00 DPS) [quest]; Darkmoon Card: Heroism (19287, -12.61 DPS, sim-verified) [quest] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (610.3 DPS) | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; Gravestone War Axe (13983, -38.59 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Skyfury Helm; neck: Pendant of Celerity; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Savage Gladiator Chain; wrist: Slashclaw Bracers; hands: Bloodmail Gauntlets; waist: Belt of Preserved Heads; legs: Sentinel's Chain Leggings; feet: Bloodmail Boots; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: The Unstoppable Force

No-known-source sample (15 of 1672, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

