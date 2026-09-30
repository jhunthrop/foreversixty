# Leveling BiS: Feral

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-5420000000000000000-0000000000000000)

Set DPS (verified): 62.4. Weights run: 1.1s. Verify run: 1.0s. 165 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.002, agility=1.436 ± 0.049, crit=8.959 ± 0.224, hit=not significant (0.000 ± 0.000), melee_haste=5.281 ± 0.416

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 18.6 | yes | Brawler's Leather Hood (252504, -0.33 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.03 DPS) [crafted]; Shadow Goggles (4373, -1.03 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 8.6 | yes | Erudite's Amulet (277204, -0.18 DPS, sim-verified) [quest]; Scholarly Pendant (277203, -0.48 DPS) [quest]; Tarnished Locket (279870, -0.48 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 7.2 | yes | Reinforced Woolen Shoulders (4315, -0.40 DPS) [crafted]; Forest Leather Mantle (4709, -0.40 DPS) [world_drop]; Double-Stitched Woolen Shoulders (4314, -0.43 DPS, sim-verified) [crafted] |
| back | Grave Shroud (279865) | Abominable Creatures [quest] | 9.8 | yes | Lambent Scale Cloak (4706, -0.05 DPS, sim-verified) [world_drop]; Dark Leather Cloak (2316, -0.05 DPS) [crafted]; Glowing Lizardscale Cloak (6449, -0.07 DPS) [dungeon] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 21.7 | yes | Defender's Leather Armor (252434, -0.08 DPS, sim-verified) [crafted]; Totemic Leather Armor (252435, -0.30 DPS) [crafted]; Murloc Scale Breastplate (5781, -0.32 DPS) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 10.4 | yes | Bristlebark Bindings (14569, -0.09 DPS, sim-verified) [world_drop]; Forest Leather Bracers (3202, -0.18 DPS) [world_drop]; Wolf Bracers (4794, -0.26 DPS) [vendor] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 125.4 | yes | Gloves of the Fang (10413, +0.00 DPS, sim-verified) [dungeon]; Gold-flecked Gloves (5195, -6.05 DPS) [dungeon]; Bristlebark Gloves (14572, -6.12 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 | yes | Brawler's Leather Belt (252428, -0.14 DPS, sim-verified) [crafted]; Deviate Scale Belt (6468, -0.21 DPS) [crafted]; Ruffian Belt (5975, -0.23 DPS) [world] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 26.8 | yes | Trapper's Leather Pants (252501, +0.00 DPS, sim-verified) [crafted]; Defender's Leather Pants (252445, -0.01 DPS) [crafted]; Leggings of the Fang (10410, -0.13 DPS) [dungeon] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 18.8 | yes | Feet of the Lynx (1121, +0.00 DPS, sim-verified) [world_drop]; Defender's Leather Boots (252441, -0.40 DPS) [crafted]; Totemic Leather Boots (252442, -0.40 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 15.0 | yes | Signet of the Zhevra (285330, -0.36 DPS) [world]; Loop of Sacrifice (281673, -0.45 DPS) [quest]; The 1 Ring (8350, -0.62 DPS) [world] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 9.3 | yes | Signet of the Zhevra (285330, +0.00 DPS, sim-verified) [world]; Loop of Sacrifice (281673, -0.13 DPS) [quest]; The 1 Ring (8350, -0.31 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | Westfall: Mr. Smite [dungeon] | 307.2 | yes | Gargoyle's Bite (12989, -1.02 DPS) [world_drop]; Staff of Westfall (2042, -1.12 DPS) [quest]; Living Root (6631, -1.80 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Grave Shroud; chest: Brawler's Leather Armor; wrist: Bravo's Armbands; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Smite's Mighty Hammer; ranged: Idol of the Huntress

No-known-source sample (15 of 165, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 7956 Bronze Warhammer; 10047 Simple Kilt; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 30 (night-elf, 0000000000000000-5423222100000000000-0000000000000000)

Set DPS (verified): 101.7. Weights run: 1.2s. Verify run: 1.2s. 288 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.002, agility=1.494 ± 0.064, crit=10.211 ± 0.292, hit=not significant (0.000 ± 0.000), melee_haste=6.340 ± 0.693

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 27.8 | yes | Azure Gustwoven Hood (277050, -0.36 DPS) [crafted]; Defender's Leather Hood (252447, -0.48 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.50 DPS, sim-verified) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.3 | yes | Sentinel's Medallion (19541, -0.17 DPS) [rep]; River Pride Choker (13087, -0.31 DPS) [world_drop]; Ghostshard Talisman (7731, -0.32 DPS, sim-verified) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 28.0 | yes | Bristlebark Amice (14573, -0.63 DPS) [world_drop]; Mantle of Thieves (2264, -0.68 DPS) [dungeon]; Barbaric Shoulders (5964, -0.73 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 17.4 | yes | Sergeant Major's Cape (16315, -0.03 DPS, sim-verified) [pvp]; Tigerstrike Mantle (13108, -0.29 DPS) [world_drop]; Wolfmaster Cape (6314, -0.39 DPS) [dungeon] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 25.9 | yes | Brawler's Leather Armor (252490, -0.20 DPS) [crafted]; Defender's Leather Tunic (252450, -0.23 DPS, sim-verified) [crafted]; Dusky Leather Armor (7374, -0.26 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 18.2 | yes | Bands of Serra'kis (6902, -0.23 DPS) [dungeon]; Barbaric Bracers (18948, -0.23 DPS, sim-verified) [crafted]; Jurassic Wristguards (6198, -0.24 DPS) [world] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 143.0 | yes | Insignia Gloves (6408, +0.00 DPS, sim-verified) [world_drop]; Toughened Leather Gloves (4253, -6.27 DPS) [crafted]; Wolfclaw Gloves (1978, -6.40 DPS) [dungeon] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 29.8 | yes | Skulker's Leather Belt (252520, -0.30 DPS, sim-verified) [crafted]; Highlander's Chain Girdle (20090, -0.31 DPS) [rep]; Highlander's Leather Girdle (20117, -0.31 DPS) [rep] |
| legs | Brawler's Leather Legguards (252516) | Leatherworking [crafted] | 29.7 | yes | Trapper's Leather Pants (252501, -0.12 DPS) [crafted]; Defender's Leather Pants (252445, -0.15 DPS) [crafted]; Brawler's Leather Pants (252500, -0.17 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 19.1 | yes | Feet of the Lynx (1121, +0.00 DPS, sim-verified) [world_drop]; Disjointed Shoes (277226, -0.37 DPS) [quest]; Insignia Boots (4055, -0.37 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 23.0 | yes | Ironspine's Eye (7686, -0.02 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.60 DPS) [dungeon]; Seal of Wrynn (2933, -0.61 DPS) [quest] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 22.9 | yes | Ironspine's Eye (7686, +0.00 DPS, sim-verified) [dungeon]; Protector's Band (20439, -0.40 DPS) [rep]; Silverlaine's Family Seal (6321, -0.59 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (102.0 DPS) | yes | Cobalt Crusher (7730, -2.40 DPS) [dungeon]; Gnarled Ash Staff (791, -4.80 DPS) [world_drop]; Viscous Hammer (13045, -21.98 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; waist: Prowler's Leather Belt; legs: Brawler's Leather Legguards; finger1: Thunderbrow Ring; finger2: Protector's Band; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 288, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7956 Bronze Warhammer; 9362 Brilliant Gold Ring; 10047 Simple Kilt

### Band 40 (night-elf, 0000000000000000-5423222121032010001-0000000000000000)

Set DPS (verified): 131.6. Weights run: 1.3s. Verify run: 1.3s. 402 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.003, agility=1.534 ± 0.071, crit=11.632 ± 0.333, hit=not significant (0.000 ± 0.000), melee_haste=7.650 ± 0.975

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 42.4 | yes | Hawkeye's Helm (14591, -0.90 DPS) [world_drop]; Cloudy Gustwoven Hood (277042, -1.19 DPS) [crafted]; Defender's Leather Helm (252455, -1.26 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 16.9 | yes | Kaleidoscope Chain (13084, -0.07 DPS, sim-verified) [world_drop]; Ghostshard Talisman (7731, -0.16 DPS) [dungeon]; Sentinel's Medallion (19541, -0.25 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 28.9 | yes | Forest Tracker Epaulets (2278, -0.03 DPS, sim-verified) [world_drop]; Imperial Leather Spaulders (4737, -0.44 DPS) [world_drop]; Fleshhide Shoulders (10774, -0.45 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 23.1 | yes | Yeti Fur Cloak (2805, -0.39 DPS) [quest]; Sergeant Major's Cape (16315, -0.43 DPS) [pvp]; Hawkeye's Cloak (14593, -0.61 DPS, sim-verified) [world_drop] |
| chest | Barbaric Harness (5739) | Leatherworking [crafted] | 28.5 | yes | Wolffear Harness (13110, -0.13 DPS) [world_drop]; Brawler's Leather Tunic (252508, -0.18 DPS, sim-verified) [crafted]; Defender's Leather Tunic (252450, -0.30 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Hawkeye's Bracers (14590, -0.04 DPS, sim-verified) [world_drop]; Barbaric Bracers (18948, -0.25 DPS) [crafted]; Bands of Serra'kis (6902, -0.34 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 182.9 | yes | Shadowskin Gloves (18238, -1.11 DPS) [crafted]; Fletcher's Gloves (7348, -1.56 DPS, sim-verified) [crafted]; Prowler's Leather Gloves (252524, -8.08 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 34.7 | yes | Highlander's Leather Girdle (20116, -0.26 DPS) [rep]; Prowler's Leather Belt (252459, -0.28 DPS, sim-verified) [crafted]; Skulker's Leather Belt (252520, -0.39 DPS) [crafted] |
| legs | Triprunner Dungarees (9624) | The Grand Betrayal [quest] | 34.6 | yes | Basilisk Hide Pants (1718, -0.12 DPS, sim-verified) [world_drop]; Brawler's Leather Legguards (252516, -0.25 DPS) [crafted]; Brawler's Leather Pants (252500, -0.38 DPS) [crafted] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 36.3 | yes | Skulker's Leather Shoes (252531, -0.34 DPS, sim-verified) [crafted]; Excelsior Boots (4109, -0.34 DPS) [quest]; Imperial Leather Boots (6431, -0.43 DPS) [world_drop] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 30.8 | yes | Protector's Band (19517, -0.43 DPS) [rep]; Ironspine's Eye (7686, -0.43 DPS) [dungeon]; Falcon's Hook (7552, -0.56 DPS) [world_drop] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 23.2 | yes | Ironspine's Eye (7686, +0.00 DPS, sim-verified) [dungeon]; Falcon's Hook (7552, -0.13 DPS) [world_drop]; Assault Band (13095, -0.18 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (131.4 DPS) | yes | Thornstone Sledgehammer (1722, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Illusionary Rod (7713, -28.99 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Sergeant Major's Cape; chest: Barbaric Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Triprunner Dungarees; feet: Prowler's Leather Shoes; finger1: Protector's Band; finger2: Thunderbrow Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty

No-known-source sample (15 of 402, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 50 (night-elf, 0000000000000000-5423222121032010001-5500000000000000)

Set DPS (verified): 155.8. Weights run: 1.3s. Verify run: 1.3s. 535 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.003, agility=1.597 ± 0.077, crit=12.908 ± 0.369, hit=not significant (0.000 ± 0.000), melee_haste=7.662 ± 1.230

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) | Captain Dirgehammer [vendor] | 196.7 | yes | Eye of Theradras (17715, -1.90 DPS, sim-verified) [dungeon]; White Bandit Mask (10008, -8.51 DPS) [crafted]; Scorpashi Skullcap (14658, -9.28 DPS) [world_drop] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 32.4 | yes | Sentinel's Medallion (19540, -0.82 DPS) [rep]; Kaleidoscope Chain (13084, -0.93 DPS) [world_drop]; Sentinel's Medallion (19539, -1.10 DPS, sim-verified) [rep] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) | Captain Dirgehammer [vendor] | 188.7 | yes | Prowler's Leather Shoulder (252534, +0.00 DPS, sim-verified) [crafted]; Failed Flying Experiment (9647, -8.29 DPS) [quest]; Skulker's Leather Shoulder (252535, -8.33 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 24.0 | yes | Sergeant Major's Cape (16336, +0.00 DPS, sim-verified) [pvp]; Pridelord Cape (14673, -0.15 DPS) [world_drop]; Bloodlust Cape (14801, -0.17 DPS) [world_drop] |
| chest | Knight's Leather Armor (220854) | Captain Dirgehammer [vendor] | 198.7 | yes | Knight's Crackling Leather Tunic (220868, -2.26 DPS, sim-verified) [vendor]; Warbear Harness (15064, -8.00 DPS) [crafted]; Blazewind Breastplate (11193, -8.59 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 33.2 | yes | Prowler's Leather Bracers (252539, -0.06 DPS, sim-verified) [crafted]; Skulker's Leather Bracers (252540, -0.15 DPS) [crafted]; Pridelord Bands (14672, -0.31 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 200.7 | yes | Sergeant Major's Leather Gauntlets (220856, -0.50 DPS, sim-verified) [vendor]; Fletcher's Gloves (7348, -1.11 DPS) [crafted]; Shadowskin Gloves (18238, -1.11 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20115) | The League of Arathor [rep] | 200.7 | yes | Highlander's Lizardhide Girdle (20103, -1.11 DPS) [rep]; Highlander's Cloth Girdle (20097, -1.51 DPS, sim-verified) [rep]; Prowler's Leather Waistguard (252473, -8.26 DPS) [crafted] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | 361.4 | yes | Knight's Leather Pants (220858, +0.00 DPS, sim-verified) [vendor]; Knight's Crackling Leather Leggings (220864, -10.01 DPS) [vendor]; Knight's Restored Leather Leggings (220882, -10.01 DPS) [vendor] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 47.7 | yes | Skulker's Leather Boots (252469, -0.08 DPS, sim-verified) [crafted]; Sandstalker Ankleguards (12470, -0.37 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.61 DPS) [crafted] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 37.6 | yes | Protector's Band (19515, -0.52 DPS, sim-verified) [rep]; Protector's Band (19517, -0.78 DPS) [rep]; Thunderbrow Ring (13097, -0.79 DPS) [world_drop] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 23.7 | yes | Masons Fraternity Ring (9533, -0.07 DPS) [quest]; Thunderbrow Ring (13097, -0.11 DPS, sim-verified) [world_drop]; Falcon's Hook (7552, -0.13 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (124.4 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Smoking Heart of the Mountain (11811, -0.97 DPS, sim-verified) [crafted] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (156.0 DPS) | yes | Kindling Stave (11750, +0.00 DPS) [dungeon]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Blight (7959, -28.10 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Knight-Lieutenant's Leather Headband; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Leather Shoulders; back: Dark Phantom Cape; chest: Knight's Leather Armor; wrist: Deepfury Bracers; waist: Highlander's Leather Girdle; legs: Stormshroud Pants; feet: Prowler's Leather Boots; finger1: Protector's Band; finger2: Ironspine's Eye; trinket1: Guardian Talisman; trinket2: Frozen Heart of the Mountain

No-known-source sample (15 of 535, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60 (night-elf, 0000000000000000-5423222121032010001-5553200000000000)

Set DPS (verified): 186.8. Weights run: 1.3s. Verify run: 1.3s. 1111 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.003, agility=1.760 ± 0.098, crit=13.720 ± 0.409, hit=not significant (0.000 ± 0.000), melee_haste=not significant (7.156 ± 1.916)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Waywatcher Cowl (240064) | Leonid Barthalomew the Revered [vendor] | sim-verified (178.5 DPS) | yes | Ragefury Eyepatch (11735, -6.82 DPS) [dungeon]; Bloodvine Lens (19998, -7.57 DPS) [crafted]; Waywatcher Hood (240072, -8.54 DPS, sim-verified) [vendor] |
| neck | Blazefury Medallion (17111) | Lord Kazzak [world] | sim-verified (170.5 DPS) | yes | Skibi's Pendant (13089, +0.00 DPS) [world_drop]; Amulet of the Darkmoon (19491, +0.00 DPS) [quest]; Medallion of the Dawn (22659, -2.01 DPS, sim-verified) [quest] |
| shoulder | Field Marshal's Dragonhide Shoulders (231693) | Captain Dirgehammer [vendor] | 245.7 | yes | Darkspear Pauldrons (272105, -0.72 DPS, sim-verified) [vendor]; Lieutenant Commander's Dragonhide Shoulders (227172, -0.82 DPS) [vendor]; Knight-Lieutenant's Leather Shoulders (220852, -2.47 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-verified (172.1 DPS) | yes | Cloak of the Honor Guard (20073, -0.20 DPS) [rep]; Sergeant Major's Cape (16337, -0.75 DPS) [pvp]; Chromatic Cloak (18509, -2.15 DPS, sim-verified) [crafted] |
| chest | Waywatcher Leathers (240075) | Leonid Barthalomew the Revered [vendor] | 576.3 | yes | Stormshroud Armor (15056, -1.08 DPS, sim-verified) [crafted]; Waywatcher Tunic (240091, -10.42 DPS) [vendor]; Waywatcher Vest (240067, -12.03 DPS) [vendor] |
| wrist | Waywatcher Wraps (240060) | Leonid Barthalomew the Revered [vendor] | sim-verified (174.3 DPS) | yes | Forest Stalker's Bracers (19587, -1.07 DPS) [rep]; Forest Stalker's Bracers (19589, -1.51 DPS) [rep]; Waywatcher Wristguards (240084, -4.28 DPS, sim-verified) [vendor] |
| hands | Waywatcher Mitts (240073) | Leonid Barthalomew the Revered [vendor] | 384.2 | yes | Marshal's Dragonhide Grips (231694, +0.00 DPS, sim-verified) [vendor]; Devilsaur Gauntlets (15063, -8.90 DPS) [crafted]; Gloves of Holy Might (867, -9.34 DPS) [world_drop] |
| waist | Highlander's Leather Girdle (20045) | The League of Arathor [rep] | 226.1 | yes | Highlander's Leather Girdle (20115, -0.96 DPS, sim-verified) [rep]; Belt of the Archmage (18405, -1.84 DPS) [crafted]; Highlander's Lizardhide Girdle (20046, -1.84 DPS) [rep] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | sim-verified (171.8 DPS) | yes | Waywatcher Legguards (240087, -1.80 DPS, sim-verified) [vendor]; Stormshroud Pants (15057, -2.58 DPS) [crafted]; Sentinel's Silk Leggings (237815, -2.58 DPS) [vendor] |
| feet | Waywatcher Sandals (240074) | Leonid Barthalomew the Revered [vendor] | 384.2 | yes | Marshal's Dragonhide Treads (231692, +0.00 DPS, sim-verified) [vendor]; Knight-Lieutenant's Dragonhide Treads (227182, -8.18 DPS) [vendor]; Waywatcher Stompers (240066, -14.74 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (168.1 DPS) | yes | Band of the Penitent (13217, -0.87 DPS) [quest]; Ring of Entropy (18543, -0.87 DPS) [world]; Wrath of Cenarius (21190, -3.05 DPS, sim-verified) [quest] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (168.1 DPS) | yes | Band of the Penitent (13217, -0.76 DPS) [quest]; Ring of Entropy (18543, -0.76 DPS) [world]; Wrath of Cenarius (21190, -2.91 DPS, sim-verified) [quest] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (168.1 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Darkmoon Card: Heroism (19287, -7.16 DPS, sim-verified) [quest] |
| trinket2 | - | - |  |  |  |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-verified (168.1 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [vendor]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Sulfuron Hammer (17193, -1.89 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Moon (23197) (or Howling Idol (272427), Enraged Idol (272428), Idol of Synthesis (272429), Swarming Idol (272430), Idol of Swiftness (279250), Idol of the Ursine Twins (279251), Idol of the Dream (220606), Idol of the Huntress (227444), Idol of the Raging Shambler (220915), Talons of Wrath (249441), Idol of the Heckler (213594), Idol of the Wild (210534), Mystic Mushroom (249396), Ferocious Idol (208689), Idol of Ursine Rage (206954), Unbalanced Idol (210195), Lunar Idol (208414)) | World drop [world_drop] | 0.0 | yes | Howling Idol (272427, +0.00 DPS, sim-verified) [vendor]; Enraged Idol (272428, +0.00 DPS) [vendor]; Idol of Synthesis (272429, +0.00 DPS) [vendor] |

**New at 60:** head: Waywatcher Cowl; neck: Blazefury Medallion; shoulder: Field Marshal's Dragonhide Shoulders; back: Cape of the Black Baron; chest: Waywatcher Leathers; wrist: Waywatcher Wraps; hands: Waywatcher Mitts; waist: Highlander's Leather Girdle; legs: Sentinel's Leather Pants; feet: Waywatcher Sandals; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Darkmoon Card: Maelstrom; trinket2: Darkmoon Card: Blue Dragon; main_hand: The Unstoppable Force; ranged: Idol of the Moon

No-known-source sample (15 of 1111, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

## Horde

### Band 20 (tauren, 0000000000000000-5420000000000000000-0000000000000000)

Set DPS (verified): 61.1. Weights run: 1.1s. Verify run: 1.0s. 167 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.002, agility=1.436 ± 0.049, crit=8.959 ± 0.224, hit=not significant (0.000 ± 0.000), melee_haste=5.281 ± 0.416

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 18.6 | yes | Brawler's Leather Hood (252504, -0.34 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.03 DPS) [crafted]; Shadow Goggles (4373, -1.03 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 8.6 | yes | Erudite's Amulet (277204, -0.17 DPS, sim-verified) [quest]; Scholarly Pendant (277203, -0.48 DPS) [quest]; Tarnished Locket (279870, -0.48 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 7.2 | yes | Reinforced Woolen Shoulders (4315, -0.40 DPS) [crafted]; Forest Leather Mantle (4709, -0.40 DPS) [world_drop]; Double-Stitched Woolen Shoulders (4314, -0.43 DPS, sim-verified) [crafted] |
| back | Grave Shroud (279865) | Abominable Creatures [quest] | 9.8 | yes | Lambent Scale Cloak (4706, -0.04 DPS, sim-verified) [world_drop]; Dark Leather Cloak (2316, -0.05 DPS) [crafted]; Glowing Lizardscale Cloak (6449, -0.07 DPS) [dungeon] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 21.7 | yes | Defender's Leather Armor (252434, -0.09 DPS, sim-verified) [crafted]; Totemic Leather Armor (252435, -0.30 DPS) [crafted]; Murloc Scale Breastplate (5781, -0.32 DPS) [crafted] |
| wrist | Bristlebark Bindings (14569) | World drop [world_drop] | 8.9 | yes | Forest Leather Bracers (3202, -0.09 DPS, sim-verified) [world_drop]; Wolf Bracers (4794, -0.18 DPS) [vendor]; Ratchet Wristwraps (274742, -0.26 DPS) [vendor] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 125.4 | yes | Gloves of the Fang (10413, +0.00 DPS, sim-verified) [dungeon]; Gold-flecked Gloves (5195, -6.05 DPS) [dungeon]; Bristlebark Gloves (14572, -6.12 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 | yes | Brawler's Leather Belt (252428, -0.13 DPS, sim-verified) [crafted]; Deviate Scale Belt (6468, -0.21 DPS) [crafted]; Ruffian Belt (5975, -0.23 DPS) [world] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 26.8 | yes | Trapper's Leather Pants (252501, +0.00 DPS, sim-verified) [crafted]; Defender's Leather Pants (252445, -0.01 DPS) [crafted]; Leggings of the Fang (10410, -0.13 DPS) [dungeon] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 18.8 | yes | Feet of the Lynx (1121, +0.00 DPS, sim-verified) [world_drop]; Defender's Leather Boots (252441, -0.40 DPS) [crafted]; Totemic Leather Boots (252442, -0.40 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 15.0 | yes | Signet of the Zhevra (285330, -0.36 DPS) [world]; Loop of Sacrifice (281673, -0.45 DPS) [quest]; Bounty Hunter's Ring (5351, -0.59 DPS) [quest] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 9.3 | yes | Signet of the Zhevra (285330, +0.00 DPS, sim-verified) [world]; Loop of Sacrifice (281673, -0.13 DPS) [quest]; Bounty Hunter's Ring (5351, -0.28 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hammerbone (270018) | Leaders of the Fang [quest] | 309.4 | yes | Smite's Mighty Hammer (7230, +0.00 DPS, sim-verified) [dungeon]; Living Root (6631, -0.69 DPS) [dungeon]; Crescent Staff (6505, -0.81 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Grave Shroud; chest: Brawler's Leather Armor; wrist: Bristlebark Bindings; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Demon Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Hammerbone; ranged: Idol of the Huntress

No-known-source sample (15 of 167, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 7956 Bronze Warhammer; 10047 Simple Kilt; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20441 Scout's Blade; 209617 Insignia of the Alliance; 209624 Insignia of the Horde; 241089 Scarlet Dagger

### Band 30 (tauren, 0000000000000000-5423222100000000000-0000000000000000)

Set DPS (verified): 100.4. Weights run: 1.2s. Verify run: 1.2s. 294 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.002, agility=1.494 ± 0.064, crit=10.211 ± 0.292, hit=not significant (0.000 ± 0.000), melee_haste=6.340 ± 0.693

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 27.8 | yes | Azure Gustwoven Hood (277050, -0.36 DPS) [crafted]; Defender's Leather Hood (252447, -0.48 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.49 DPS, sim-verified) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.3 | yes | Scout's Medallion (19537, -0.17 DPS) [rep]; River Pride Choker (13087, -0.31 DPS) [world_drop]; Ghostshard Talisman (7731, -0.32 DPS, sim-verified) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 28.0 | yes | Bristlebark Amice (14573, -0.63 DPS) [world_drop]; Mantle of Thieves (2264, -0.68 DPS) [dungeon]; Barbaric Shoulders (5964, -0.74 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 17.4 | yes | Tigerstrike Mantle (13108, -0.37 DPS, sim-verified) [world_drop]; Wolfmaster Cape (6314, -0.39 DPS) [dungeon]; Wildhunter Cloak (16658, -0.39 DPS) [quest] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 25.9 | yes | Brawler's Leather Armor (252490, -0.20 DPS) [crafted]; Defender's Leather Tunic (252450, -0.25 DPS, sim-verified) [crafted]; Dusky Leather Armor (7374, -0.26 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 18.2 | yes | Bands of Serra'kis (6902, -0.23 DPS) [dungeon]; Jurassic Wristguards (6198, -0.24 DPS) [world]; Barbaric Bracers (18948, -0.25 DPS, sim-verified) [crafted] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 143.0 | yes | Insignia Gloves (6408, +0.00 DPS, sim-verified) [world_drop]; Toughened Leather Gloves (4253, -6.27 DPS) [crafted]; Wolfclaw Gloves (1978, -6.40 DPS) [dungeon] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 29.8 | yes | Skulker's Leather Belt (252520, -0.29 DPS, sim-verified) [crafted]; Defiler's Chain Girdle (20152, -0.31 DPS) [rep]; Defiler's Leather Girdle (20191, -0.31 DPS) [rep] |
| legs | Brawler's Leather Legguards (252516) | Leatherworking [crafted] | 29.7 | yes | Trapper's Leather Pants (252501, -0.12 DPS) [crafted]; Defender's Leather Pants (252445, -0.15 DPS) [crafted]; Brawler's Leather Pants (252500, -0.16 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 19.1 | yes | Feet of the Lynx (1121, +0.00 DPS, sim-verified) [world_drop]; Stomping Boots (3741, -0.20 DPS) [quest]; Disjointed Shoes (277226, -0.37 DPS) [quest] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 23.0 | yes | Ironspine's Eye (7686, -0.02 DPS) [dungeon]; Band of the Fist (17694, -0.41 DPS) [quest]; Silverlaine's Family Seal (6321, -0.60 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 22.9 | yes | Ironspine's Eye (7686, +0.00 DPS, sim-verified) [dungeon]; Band of the Fist (17694, -0.40 DPS) [quest]; Legionnaire's Band (20429, -0.40 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (100.7 DPS) | yes | Cobalt Crusher (7730, -2.40 DPS) [dungeon]; Gnarled Ash Staff (791, -4.80 DPS) [world_drop]; Viscous Hammer (13045, -21.55 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; waist: Prowler's Leather Belt; legs: Brawler's Leather Legguards; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 294, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7956 Bronze Warhammer; 9362 Brilliant Gold Ring; 10047 Simple Kilt; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 40 (tauren, 0000000000000000-5423222121032010001-0000000000000000)

Set DPS (verified): 129.5. Weights run: 1.3s. Verify run: 1.3s. 408 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.003, agility=1.534 ± 0.071, crit=11.632 ± 0.333, hit=not significant (0.000 ± 0.000), melee_haste=7.650 ± 0.975

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 42.4 | yes | Hawkeye's Helm (14591, -0.90 DPS) [world_drop]; Cloudy Gustwoven Hood (277042, -1.19 DPS) [crafted]; Defender's Leather Helm (252455, -1.32 DPS, sim-verified) [crafted] |
| neck | Ethereal Talisman (4430) | The Crown of Will [quest] | 17.7 | yes | Scout's Medallion (19536, +0.00 DPS, sim-verified) [rep]; Kaleidoscope Chain (13084, -0.13 DPS) [world_drop]; Ghostshard Talisman (7731, -0.21 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 28.9 | yes | Forest Tracker Epaulets (2278, -0.03 DPS, sim-verified) [world_drop]; Imperial Leather Spaulders (4737, -0.44 DPS) [world_drop]; Fleshhide Shoulders (10774, -0.45 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 17.7 | yes | Imperial Cloak (6432, -0.30 DPS) [world_drop]; Parachute Cloak (10518, -0.30 DPS) [crafted]; Scorpashi Cape (14656, -0.45 DPS, sim-verified) [world_drop] |
| chest | Barbaric Harness (5739) | Leatherworking [crafted] | 28.5 | yes | Wolffear Harness (13110, -0.13 DPS) [world_drop]; Brawler's Leather Tunic (252508, -0.18 DPS, sim-verified) [crafted]; Defender's Leather Tunic (252450, -0.30 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Hawkeye's Bracers (14590, -0.00 DPS, sim-verified) [world_drop]; Barbaric Bracers (18948, -0.25 DPS) [crafted]; Bands of Serra'kis (6902, -0.34 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 182.9 | yes | Shadowskin Gloves (18238, -1.11 DPS) [crafted]; Fletcher's Gloves (7348, -1.52 DPS, sim-verified) [crafted]; Prowler's Leather Gloves (252524, -8.08 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 34.7 | yes | Prowler's Leather Belt (252459, -0.25 DPS) [crafted]; Defiler's Leather Girdle (20192, -0.26 DPS) [rep]; Tharg's Shoelace (9705, -0.53 DPS, sim-verified) [quest] |
| legs | Triprunner Dungarees (9624) | Rig Wars [quest] | 34.6 | yes | Basilisk Hide Pants (1718, -0.13 DPS, sim-verified) [world_drop]; Brawler's Leather Legguards (252516, -0.25 DPS) [crafted]; Brawler's Leather Pants (252500, -0.38 DPS) [crafted] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 36.3 | yes | Skulker's Leather Shoes (252531, -0.34 DPS, sim-verified) [crafted]; Excelsior Boots (4109, -0.34 DPS) [quest]; Imperial Leather Boots (6431, -0.43 DPS) [world_drop] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 30.8 | yes | Legionnaire's Band (19513, -0.43 DPS) [rep]; Ironspine's Eye (7686, -0.43 DPS) [dungeon]; Falcon's Hook (7552, -0.56 DPS) [world_drop] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 23.2 | yes | Ironspine's Eye (7686, +0.00 DPS, sim-verified) [dungeon]; Falcon's Hook (7552, -0.13 DPS) [world_drop]; Assault Band (13095, -0.18 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (129.3 DPS) | yes | Thornstone Sledgehammer (1722, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Illusionary Rod (7713, -28.31 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; neck: Ethereal Talisman; shoulder: Sunburn Spaulders; chest: Barbaric Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Triprunner Dungarees; feet: Prowler's Leather Shoes; finger1: Legionnaire's Band; finger2: Thunderbrow Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty

No-known-source sample (15 of 408, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7956 Bronze Warhammer; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 10047 Simple Kilt

### Band 50 (tauren, 0000000000000000-5423222121032010001-5500000000000000)

Set DPS (verified): 158.7. Weights run: 1.3s. Verify run: 1.3s. 541 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.003, agility=1.597 ± 0.077, crit=12.908 ± 0.369, hit=not significant (0.000 ± 0.000), melee_haste=7.662 ± 1.230

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Leather Headband (220851) | Lady Palanseer [vendor] | 196.7 | yes | Eye of Theradras (17715, -1.88 DPS, sim-verified) [dungeon]; White Bandit Mask (10008, -8.51 DPS) [crafted]; Undercity Reservist's Cap (20643, -8.64 DPS) [quest] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 32.4 | yes | Woven Ivy Necklace (19159, -0.37 DPS, sim-verified) [quest]; Scout's Medallion (19535, -0.73 DPS) [rep]; Ethereal Talisman (4430, -0.80 DPS) [quest] |
| shoulder | Blood Guard's Leather Shoulders (220853) | Lady Palanseer [vendor] | 188.7 | yes | Prowler's Leather Shoulder (252534, +0.00 DPS, sim-verified) [crafted]; Failed Flying Experiment (9647, -8.29 DPS) [quest]; Skulker's Leather Shoulder (252535, -8.33 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 24.0 | yes | Bloodlust Cape (14801, -0.17 DPS) [world_drop]; Blackflame Cape (13109, -0.27 DPS) [world_drop]; Pridelord Cape (14673, -0.27 DPS, sim-verified) [world_drop] |
| chest | Stone Guard's Leather Armor (220855) | Lady Palanseer [vendor] | 198.7 | yes | Stone Guard's Crackling Leather Tunic (220869, -2.22 DPS, sim-verified) [vendor]; Warbear Harness (15064, -8.00 DPS) [crafted]; Blazewind Breastplate (11193, -8.59 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 33.2 | yes | Prowler's Leather Bracers (252539, -0.03 DPS, sim-verified) [crafted]; Skulker's Leather Bracers (252540, -0.15 DPS) [crafted]; Pridelord Bands (14672, -0.31 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 200.7 | yes | First Sergeant's Leather Gauntlets (220857, -0.50 DPS, sim-verified) [vendor]; Fletcher's Gloves (7348, -1.11 DPS) [crafted]; Shadowskin Gloves (18238, -1.11 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20193) | The Defilers [rep] | 200.7 | yes | Defiler's Lizardhide Girdle (20174, -1.11 DPS) [rep]; Defiler's Cloth Girdle (20165, -1.50 DPS, sim-verified) [rep]; Prowler's Leather Waistguard (252473, -8.26 DPS) [crafted] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | 361.4 | yes | Stone Guard's Leather Pants (220859, +0.00 DPS, sim-verified) [vendor]; Stone Guard's Crackling Leather Leggings (220865, -10.01 DPS) [vendor]; Stone Guard's Restored Leather Leggings (220883, -10.01 DPS) [vendor] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 47.7 | yes | Skulker's Leather Boots (252469, -0.08 DPS, sim-verified) [crafted]; Sandstalker Ankleguards (12470, -0.37 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.61 DPS) [crafted] |
| finger1 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 37.6 | yes | Legionnaire's Band (19512, -0.51 DPS, sim-verified) [rep]; Ironspine's Eye (7686, -0.77 DPS) [dungeon]; Legionnaire's Band (19513, -0.78 DPS) [rep] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 | yes | Ironspine's Eye (7686, +0.00 DPS, sim-verified) [dungeon]; Thunderbrow Ring (13097, -0.04 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.09 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (126.7 DPS) | yes | Tidal Charm (1404, -2.33 DPS) [vendor]; Ankh of Life (1713, -2.33 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -3.25 DPS, sim-verified) [crafted] |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (158.8 DPS) | yes | Kindling Stave (11750, +0.00 DPS) [dungeon]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Blight (7959, -28.70 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Blood Guard's Leather Headband; neck: Skibi's Pendant; shoulder: Blood Guard's Leather Shoulders; back: Dark Phantom Cape; chest: Stone Guard's Leather Armor; wrist: Deepfury Bracers; waist: Defiler's Leather Girdle; legs: Stormshroud Pants; feet: Prowler's Leather Boots; finger1: Legionnaire's Band; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Guardian Talisman

No-known-source sample (15 of 541, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7956 Bronze Warhammer; 8708 Hammer of Expertise

### Band 60 (tauren, 0000000000000000-5423222121032010001-5553200000000000)

Set DPS (verified): 191.5. Weights run: 1.3s. Verify run: 1.3s. 1117 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.003, agility=1.760 ± 0.098, crit=13.720 ± 0.409, hit=not significant (0.000 ± 0.000), melee_haste=not significant (7.156 ± 1.916)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Waywatcher Cowl (240064) | Leonid Barthalomew the Revered [vendor] | sim-verified (181.6 DPS) | yes | Ragefury Eyepatch (11735, -6.82 DPS) [dungeon]; Bloodvine Lens (19998, -7.57 DPS) [crafted]; Waywatcher Hood (240072, -8.53 DPS, sim-verified) [vendor] |
| neck | Blazefury Medallion (17111) | Lord Kazzak [world] | sim-verified (173.3 DPS) | yes | Skibi's Pendant (13089, +0.00 DPS) [world_drop]; Amulet of the Darkmoon (19491, +0.00 DPS) [quest]; Medallion of the Dawn (22659, -1.93 DPS, sim-verified) [quest] |
| shoulder | Warlord's Dragonhide Shoulders (231684) | Lady Palanseer [vendor] | 245.7 | yes | Darkspear Pauldrons (272105, -0.69 DPS, sim-verified) [vendor]; Champion's Dragonhide Shoulders (227175, -0.82 DPS) [vendor]; Blood Guard's Leather Shoulders (220853, -2.47 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-verified (175.3 DPS) | yes | Deathguard's Cloak (20068, -0.20 DPS) [rep]; Stoneskin Gargoyle Cape (13397, -0.87 DPS) [dungeon]; Chromatic Cloak (18509, -2.15 DPS, sim-verified) [crafted] |
| chest | Waywatcher Leathers (240075) | Leonid Barthalomew the Revered [vendor] | 576.3 | yes | Stormshroud Armor (15056, -1.24 DPS, sim-verified) [crafted]; Waywatcher Tunic (240091, -10.42 DPS) [vendor]; Waywatcher Vest (240067, -12.03 DPS) [vendor] |
| wrist | Waywatcher Wraps (240060) | Leonid Barthalomew the Revered [vendor] | sim-verified (177.4 DPS) | yes | Forest Stalker's Bracers (19587, -1.07 DPS) [rep]; Forest Stalker's Bracers (19589, -1.51 DPS) [rep]; Waywatcher Wristguards (240084, -4.29 DPS, sim-verified) [vendor] |
| hands | Waywatcher Mitts (240073) | Leonid Barthalomew the Revered [vendor] | 384.2 | yes | General's Dragonhide Grips (231688, +0.00 DPS, sim-verified) [vendor]; Devilsaur Gauntlets (15063, -8.90 DPS) [crafted]; Gloves of Holy Might (867, -9.34 DPS) [world_drop] |
| waist | Defiler's Leather Girdle (20190) | The Defilers [rep] | 226.1 | yes | Defiler's Leather Girdle (20193, -0.96 DPS, sim-verified) [rep]; Belt of the Archmage (18405, -1.84 DPS) [crafted]; Defiler's Cloth Girdle (20163, -1.84 DPS) [rep] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | sim-verified (174.9 DPS) | yes | Waywatcher Legguards (240087, -1.80 DPS, sim-verified) [vendor]; Stormshroud Pants (15057, -2.58 DPS) [crafted]; Sentinel's Silk Leggings (237815, -2.58 DPS) [vendor] |
| feet | General's Dragonhide Treads (231683) | Lady Palanseer [vendor] | sim-verified (174.9 DPS) | yes | Blood Guard's Dragonhide Treads (227181, -0.47 DPS) [vendor]; Waywatcher Sandals (240074, -1.77 DPS, sim-verified) [vendor]; Waywatcher Stompers (240066, -7.04 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (170.9 DPS) | yes | Band of the Penitent (13217, -0.87 DPS) [quest]; Ring of Entropy (18543, -0.87 DPS) [world]; Wrath of Cenarius (21190, -2.59 DPS, sim-verified) [quest] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (170.9 DPS) | yes | Band of the Penitent (13217, -0.76 DPS) [quest]; Ring of Entropy (18543, -0.76 DPS) [world]; Wrath of Cenarius (21190, -2.45 DPS, sim-verified) [quest] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (167.5 DPS) | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (170.9 DPS) | yes | Tidal Charm (1404, -2.28 DPS) [vendor]; Guardian Talisman (1490, -2.28 DPS) [quest]; Shard of the Fallen Star (21891, -3.38 DPS, sim-verified) [world_drop] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (170.9 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [vendor]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Sulfuron Hammer (17193, -1.69 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Moon (23197) (or Howling Idol (272427), Enraged Idol (272428), Idol of Synthesis (272429), Swarming Idol (272430), Idol of Swiftness (279250), Idol of the Ursine Twins (279251), Idol of the Dream (220606), Idol of the Huntress (227444), Idol of the Raging Shambler (220915), Talons of Wrath (249441), Idol of the Heckler (213594), Idol of the Wild (210534), Mystic Mushroom (249396), Ferocious Idol (208689), Idol of Ursine Rage (206954), Unbalanced Idol (210195), Lunar Idol (208414)) | World drop [world_drop] | 0.0 | yes | Howling Idol (272427, +0.00 DPS, sim-verified) [vendor]; Enraged Idol (272428, +0.00 DPS) [vendor]; Idol of Synthesis (272429, +0.00 DPS) [vendor] |

**New at 60:** head: Waywatcher Cowl; neck: Blazefury Medallion; shoulder: Warlord's Dragonhide Shoulders; back: Cape of the Black Baron; chest: Waywatcher Leathers; wrist: Waywatcher Wraps; hands: Waywatcher Mitts; waist: Defiler's Leather Girdle; legs: Sentinel's Leather Pants; feet: General's Dragonhide Treads; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: The Unstoppable Force; ranged: Idol of the Moon

No-known-source sample (15 of 1117, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7956 Bronze Warhammer; 8708 Hammer of Expertise

