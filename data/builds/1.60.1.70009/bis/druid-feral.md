# Leveling BiS: Feral

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-5420000000000000000-0000000000000000)

Set DPS (verified): 62.4. Weights run: 1.4s. Verify run: 1.3s. 168 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.002, agility=1.436 ± 0.049, crit=8.959 ± 0.224, hit=1.745 ± 0.067, melee_haste=5.281 ± 0.416

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 18.6 attack_power points (1.03 DPS) | yes | Brawler's Leather Hood (252504, -0.33 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 8.6 attack_power points (0.48 DPS) | yes | Erudite's Amulet (277204, -0.18 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 7.2 attack_power points (0.40 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.43 DPS, sim-verified) [crafted] |
| back | Grave Shroud (279865) | Abominable Creatures [quest] | 9.8 attack_power points (0.55 DPS) | yes | Lambent Scale Cloak (4706, -0.05 DPS, sim-verified) [world_drop]; Dark Leather Cloak (2316, -0.05 DPS) [crafted]; Glowing Lizardscale Cloak (6449, -0.07 DPS) [dungeon] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 21.7 attack_power points (1.20 DPS) | yes | Defender's Leather Armor (252434, -0.08 DPS, sim-verified) [crafted]; Totemic Leather Armor (252435, -0.30 DPS) [crafted]; Murloc Scale Breastplate (5781, -0.32 DPS) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 10.4 attack_power points (0.58 DPS) | yes | Bristlebark Bindings (14569, -0.09 DPS, sim-verified) [world_drop]; Forest Leather Bracers (3202, -0.18 DPS) [world_drop]; Wolf Bracers (4794, -0.26 DPS) [vendor] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 125.4 attack_power points (6.95 DPS) | yes | Gloves of the Fang (10413, +0.00 DPS, sim-verified) [dungeon]; Gold-flecked Gloves (5195, -6.05 DPS) [dungeon]; Bristlebark Gloves (14572, -6.12 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (1.00 DPS) | yes | Brawler's Leather Belt (252428, -0.14 DPS, sim-verified) [crafted]; Deviate Scale Belt (6468, -0.21 DPS) [crafted]; Ruffian Belt (5975, -0.23 DPS) [world] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 26.8 attack_power points (1.49 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS, sim-verified) [crafted]; Defender's Leather Pants (252445, -0.01 DPS) [crafted]; Leggings of the Fang (10410, -0.13 DPS) [dungeon] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 18.8 attack_power points (1.04 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS, sim-verified) [world_drop]; Defender's Leather Boots (252441, -0.40 DPS) [crafted]; Totemic Leather Boots (252442, -0.40 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 15.0 attack_power points (0.83 DPS) | yes | Signet of the Zhevra (285330, -0.36 DPS) [world]; Loop of Sacrifice (281673, -0.45 DPS) [quest]; The 1 Ring (8350, -0.62 DPS) [world] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 9.3 attack_power points (0.51 DPS) | yes | Signet of the Zhevra (285330, +0.00 DPS, sim-verified) [world]; Loop of Sacrifice (281673, -0.13 DPS) [quest]; The 1 Ring (8350, -0.31 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | Westfall: Mr. Smite [dungeon] | 307.2 attack_power points (17.03 DPS) | yes | Gargoyle's Bite (12989, -1.02 DPS) [world_drop]; Staff of Westfall (2042, -1.12 DPS) [quest]; Living Root (6631, -1.80 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Grave Shroud; chest: Brawler's Leather Armor; wrist: Bravo's Armbands; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Smite's Mighty Hammer; ranged: Idol of the Huntress

No-known-source sample (15 of 168, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14148 Crystalline Cuffs

### Band 30 (night-elf, 0000000000000000-5423222100000000000-0000000000000000)

Set DPS (verified): 101.7. Weights run: 1.6s. Verify run: 1.6s. 296 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.002, agility=1.494 ± 0.064, crit=10.211 ± 0.292, hit=2.054 ± 0.079, melee_haste=6.340 ± 0.693

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 27.8 attack_power points (1.45 DPS) | yes | Azure Gustwoven Hood (277050, -0.36 DPS) [crafted]; Defender's Leather Hood (252447, -0.48 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.50 DPS, sim-verified) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.3 attack_power points (0.80 DPS) | yes | Sentinel's Medallion (19541, -0.17 DPS) [rep]; River Pride Choker (13087, -0.31 DPS) [world_drop]; Ghostshard Talisman (7731, -0.32 DPS, sim-verified) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 28.0 attack_power points (1.46 DPS) | yes | Bristlebark Amice (14573, -0.63 DPS) [world_drop]; Mantle of Thieves (2264, -0.68 DPS) [dungeon]; Barbaric Shoulders (5964, -0.73 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 17.4 attack_power points (0.91 DPS) | yes | Sergeant Major's Cape (16315, -0.03 DPS, sim-verified) [pvp]; Tigerstrike Mantle (13108, -0.29 DPS) [world_drop]; Wolfmaster Cape (6314, -0.39 DPS) [dungeon] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 25.9 attack_power points (1.35 DPS) | yes | Brawler's Leather Armor (252490, -0.20 DPS) [crafted]; Defender's Leather Tunic (252450, -0.23 DPS, sim-verified) [crafted]; Dusky Leather Armor (7374, -0.26 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 18.2 attack_power points (0.95 DPS) | yes | Bands of Serra'kis (6902, -0.23 DPS) [dungeon]; Barbaric Bracers (18948, -0.23 DPS, sim-verified) [crafted]; Jurassic Wristguards (6198, -0.24 DPS) [world] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 143.0 attack_power points (7.47 DPS) | yes | Insignia Gloves (6408, +0.00 DPS, sim-verified) [world_drop]; Toughened Leather Gloves (4253, -6.27 DPS) [crafted]; Wolfclaw Gloves (1978, -6.40 DPS) [dungeon] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 29.8 attack_power points (1.56 DPS) | yes | Skulker's Leather Belt (252520, -0.30 DPS, sim-verified) [crafted]; Highlander's Chain Girdle (20090, -0.31 DPS) [rep]; Highlander's Leather Girdle (20117, -0.31 DPS) [rep] |
| legs | Brawler's Leather Legguards (252516) | Leatherworking [crafted] | 29.7 attack_power points (1.55 DPS) | yes | Trapper's Leather Pants (252501, -0.12 DPS) [crafted]; Defender's Leather Pants (252445, -0.15 DPS) [crafted]; Brawler's Leather Pants (252500, -0.17 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 19.1 attack_power points (1.00 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS, sim-verified) [world_drop]; Disjointed Shoes (277226, -0.37 DPS) [quest]; Insignia Boots (4055, -0.37 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 23.0 attack_power points (1.20 DPS) | yes | Ironspine's Eye (7686, -0.02 DPS) [dungeon]; Tiger Band (6749, -0.48 DPS) [quest]; Silverlaine's Family Seal (6321, -0.60 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 22.9 attack_power points (1.20 DPS) | yes | Ironspine's Eye (7686, +0.00 DPS, sim-verified) [dungeon]; Protector's Band (20439, -0.40 DPS) [rep]; Tiger Band (6749, -0.47 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (101.7 DPS) | yes | Cobalt Crusher (7730, -2.40 DPS) [dungeon]; Wind Spirit Staff (6689, -4.03 DPS) [dungeon]; Viscous Hammer (13045, -21.98 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; waist: Prowler's Leather Belt; legs: Brawler's Leather Legguards; finger1: Thunderbrow Ring; finger2: Protector's Band; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 296, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (night-elf, 0000000000000000-5423222121032010001-0000000000000000)

Set DPS (verified): 131.6. Weights run: 1.7s. Verify run: 1.8s. 416 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.003, agility=1.534 ± 0.071, crit=11.632 ± 0.333, hit=2.271 ± 0.096, melee_haste=7.650 ± 0.975

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 42.4 attack_power points (2.35 DPS) | yes | Hawkeye's Helm (14591, -0.90 DPS) [world_drop]; Cloudy Gustwoven Hood (277042, -1.19 DPS) [crafted]; Defender's Leather Helm (252455, -1.30 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 16.9 attack_power points (0.93 DPS) | yes | Kaleidoscope Chain (13084, -0.09 DPS, sim-verified) [world_drop]; Ghostshard Talisman (7731, -0.16 DPS) [dungeon]; Sentinel's Medallion (19541, -0.25 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 28.9 attack_power points (1.60 DPS) | yes | Forest Tracker Epaulets (2278, -0.03 DPS, sim-verified) [world_drop]; Flintrock Shoulders (7755, -0.11 DPS) [dungeon]; Imperial Leather Spaulders (4737, -0.44 DPS) [world_drop] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | sim-verified (129.3 DPS) | yes | Hawkeye's Cloak (14593, -0.30 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.39 DPS) [quest]; Dark Hooded Cape (5257, -2.07 DPS, sim-verified) [world] |
| chest | Barbaric Harness (5739) | Leatherworking [crafted] | sim-verified (129.6 DPS) | yes | Brawler's Leather Tunic (252508, -0.13 DPS) [crafted]; Wolffear Harness (13110, -0.13 DPS) [world_drop]; Quillward Harness (10583, -2.32 DPS, sim-verified) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.11 DPS) | yes | Hawkeye's Bracers (14590, -0.03 DPS, sim-verified) [world_drop]; Barbaric Bracers (18948, -0.25 DPS) [crafted]; Bands of Serra'kis (6902, -0.34 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 182.9 attack_power points (10.13 DPS) | yes | Shadowskin Gloves (18238, -1.11 DPS) [crafted]; Fletcher's Gloves (7348, -1.54 DPS, sim-verified) [crafted]; Prowler's Leather Gloves (252524, -8.08 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 34.7 attack_power points (1.92 DPS) | yes | Highlander's Leather Girdle (20116, -0.26 DPS) [rep]; Prowler's Leather Belt (252459, -0.27 DPS, sim-verified) [crafted]; Skulker's Leather Belt (252520, -0.39 DPS) [crafted] |
| legs | Triprunner Dungarees (9624) | The Grand Betrayal [quest] | 34.6 attack_power points (1.91 DPS) | yes | Basilisk Hide Pants (1718, -0.14 DPS, sim-verified) [world_drop]; Brawler's Leather Legguards (252516, -0.25 DPS) [crafted]; Brawler's Leather Pants (252500, -0.38 DPS) [crafted] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 36.3 attack_power points (2.01 DPS) | yes | Excelsior Boots (4109, -0.34 DPS) [quest]; Skulker's Leather Shoes (252531, -0.36 DPS, sim-verified) [crafted]; Imperial Leather Boots (6431, -0.43 DPS) [world_drop] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 30.8 attack_power points (1.71 DPS) | yes | Protector's Band (19517, -0.43 DPS) [rep]; Ironspine's Eye (7686, -0.43 DPS) [dungeon]; Falcon's Hook (7552, -0.56 DPS) [world_drop] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 23.2 attack_power points (1.28 DPS) | yes | Ironspine's Eye (7686, +0.00 DPS, sim-verified) [dungeon]; Falcon's Hook (7552, -0.13 DPS) [world_drop]; Assault Band (13095, -0.18 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (127.2 DPS) | yes | Thornstone Sledgehammer (1722, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Illusionary Rod (7713, -27.86 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Sergeant Major's Cape; chest: Barbaric Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Triprunner Dungarees; feet: Prowler's Leather Shoes; finger1: Protector's Band; finger2: Thunderbrow Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty

No-known-source sample (15 of 416, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 50 (night-elf, 0000000000000000-5423222121032010001-5500000000000000)

Set DPS (verified): 128.1. Weights run: 1.7s. Verify run: 1.6s. 558 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.003, agility=1.597 ± 0.077, crit=12.908 ± 0.369, hit=2.557 ± 0.109, melee_haste=7.662 ± 1.230

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) | Captain Dirgehammer [vendor] | 222.3 attack_power points (12.32 DPS) | yes | Eye of Theradras (17715, -1.58 DPS, sim-verified) [dungeon]; White Bandit Mask (10008, -9.93 DPS) [crafted]; Scorpashi Skullcap (14658, -10.69 DPS) [world_drop] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 32.4 attack_power points (1.79 DPS) | yes | Sentinel's Medallion (19539, -0.82 DPS, sim-verified) [rep]; Sentinel's Medallion (19540, -0.82 DPS) [rep]; Kaleidoscope Chain (13084, -0.93 DPS) [world_drop] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) | Captain Dirgehammer [vendor] | 214.3 attack_power points (11.87 DPS) | yes | Prowler's Leather Shoulder (252534, +0.00 DPS, sim-verified) [crafted]; Failed Flying Experiment (9647, -9.70 DPS) [quest]; Skulker's Leather Shoulder (252535, -9.74 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 36.3 attack_power points (2.01 DPS) | yes | Dark Hooded Cape (5257, +0.00 DPS, sim-verified) [world]; Blisterbane Wrap (12552, -0.68 DPS) [dungeon]; Dark Phantom Cape (13122, -0.68 DPS) [world_drop] |
| chest | Knight's Leather Armor (220854) | Captain Dirgehammer [vendor] | 224.3 attack_power points (12.43 DPS) | yes | Knight's Crackling Leather Tunic (220868, -1.84 DPS, sim-verified) [vendor]; Grizzled Pelt (22274, -8.83 DPS) [quest]; Mixologist's Tunic (12793, -9.14 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 33.2 attack_power points (1.84 DPS) | yes | Prowler's Leather Bracers (252539, -0.01 DPS, sim-verified) [crafted]; Skulker's Leather Bracers (252540, -0.15 DPS) [crafted]; Pridelord Bands (14672, -0.31 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 200.7 attack_power points (11.12 DPS) | yes | Sergeant Major's Leather Gauntlets (220856, -0.38 DPS, sim-verified) [vendor]; Fletcher's Gloves (7348, -1.11 DPS) [crafted]; Shadowskin Gloves (18238, -1.11 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20115) | The League of Arathor [rep] | 200.7 attack_power points (11.12 DPS) | yes | Highlander's Lizardhide Girdle (20103, -1.11 DPS) [rep]; Highlander's Cloth Girdle (20097, -1.16 DPS, sim-verified) [rep]; Prowler's Leather Waistguard (252473, -8.26 DPS) [crafted] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | sim-verified (128.1 DPS) | yes | Knight's Crackling Leather Leggings (220864, -1.00 DPS) [vendor]; Stormshroud Pants (15057, -1.28 DPS, sim-verified) [crafted]; Knight's Restored Leather Leggings (220882, -2.41 DPS) [vendor] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 47.7 attack_power points (2.64 DPS) | yes | Skulker's Leather Boots (252469, -0.10 DPS, sim-verified) [crafted]; Sandstalker Ankleguards (12470, -0.37 DPS) [dungeon]; Shadefiend Boots (11675, -0.39 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 45.6 attack_power points (2.53 DPS) | yes | Ironspine's Eye (7686, -1.21 DPS) [dungeon]; Thunderbrow Ring (13097, -1.23 DPS) [world_drop]; Masons Fraternity Ring (9533, -1.29 DPS) [quest] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 37.6 attack_power points (2.08 DPS) | yes | Protector's Band (19515, -0.38 DPS, sim-verified) [rep]; Ironspine's Eye (7686, -0.77 DPS) [dungeon]; Protector's Band (19517, -0.78 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (126.8 DPS) | yes | Smoking Heart of the Mountain (11811, -0.81 DPS, sim-verified) [crafted] |
| main_hand | Blight (7959) | Blacksmithing [crafted] | sim-verified (126.8 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Kindling Stave (11750, -3.57 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Knight-Lieutenant's Leather Headband; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Leather Shoulders; back: Blackveil Cape; chest: Knight's Leather Armor; wrist: Deepfury Bracers; waist: Highlander's Leather Girdle; legs: Knight's Leather Pants; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Guardian Talisman; trinket2: Frozen Heart of the Mountain; main_hand: Blight

No-known-source sample (15 of 558, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60 (night-elf, 0000000000000000-5423222121032010001-5553200000000000)

Set DPS (verified): 178.5. Weights run: 1.7s. Verify run: 1.7s. 1298 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.003, agility=1.760 ± 0.098, crit=13.720 ± 0.409, hit=3.054 ± 0.134, melee_haste=not significant (7.156 ± 1.916)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Waywatcher Cowl (240064) | Leonid Barthalomew the Revered [vendor] | sim-verified (176.7 DPS) | yes | Ragefury Eyepatch (11735, -6.82 DPS) [dungeon]; Bloodvine Lens (19998, -7.57 DPS) [crafted]; Waywatcher Hood (240072, -8.63 DPS, sim-verified) [vendor] |
| neck | Blazefury Medallion (17111) | Lord Kazzak [world] | sim-verified (168.1 DPS) | yes | Amulet of the Darkmoon (19491, +0.00 DPS) [quest]; Beads of Ogre Might (22150, +0.00 DPS) [quest]; Medallion of the Dawn (22659, -2.00 DPS, sim-verified) [quest] |
| shoulder | Waywatcher Mantle (240070) | Leonid Barthalomew the Revered [vendor] | 253.2 attack_power points (13.73 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS, sim-verified) [vendor]; Darkspear Pauldrons (272105, -0.74 DPS) [vendor]; Knight-Lieutenant's Leather Shoulders (220852, -1.22 DPS) [vendor] |
| back | Chromatic Cloak (18509) (or Fel Cape (279269)) | Leatherworking [crafted] | 192.1 attack_power points (10.42 DPS) | yes | Fel Cape (279269, +0.00 DPS, sim-verified) [crafted]; Howler's Furs (272414, -7.24 DPS) [vendor]; Earthweave Cloak (21187, -7.33 DPS) [quest] |
| chest | Waywatcher Leathers (240075) | Leonid Barthalomew the Revered [vendor] | sim-verified (168.1 DPS) | yes | Tunic of Undead Slaying (23089, -2.76 DPS, sim-verified) [world]; Waywatcher Vest (240067, -10.37 DPS) [vendor]; Stormshroud Armor (15056, -10.42 DPS) [crafted] |
| wrist | Waywatcher Wristguards (240084) | Leonid Barthalomew the Revered [vendor] | sim-verified (168.1 DPS) | yes | Wristwraps of Undead Slaying (23093, -1.07 DPS, sim-verified) [world]; Waywatcher Wraps (240060, -6.15 DPS) [vendor]; Waywatcher Bracers (240076, -6.74 DPS) [vendor] |
| hands | Waywatcher Mitts (240073) | Leonid Barthalomew the Revered [vendor] | 414.7 attack_power points (22.50 DPS) | yes | Marshal's Dragonhide Grips (231694, +0.00 DPS, sim-verified) [vendor]; Stormshroud Gloves (21278, -10.42 DPS) [crafted]; Devilsaur Gauntlets (15063, -10.56 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20045) | The League of Arathor [rep] | 226.1 attack_power points (12.26 DPS) | yes | Highlander's Leather Girdle (20115, -0.97 DPS, sim-verified) [rep]; Belt of the Archmage (18405, -1.84 DPS) [crafted]; Highlander's Lizardhide Girdle (20046, -1.84 DPS) [rep] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | sim-verified (169.9 DPS) | yes | Waywatcher Legguards (240087, -1.82 DPS, sim-verified) [vendor]; Stormshroud Pants (15057, -2.58 DPS) [crafted]; Sentinel's Silk Leggings (237815, -2.58 DPS) [vendor] |
| feet | Waywatcher Sandals (240074) | Leonid Barthalomew the Revered [vendor] | 414.7 attack_power points (22.50 DPS) | yes | Marshal's Dragonhide Treads (231692, +0.00 DPS, sim-verified) [vendor]; Knight-Lieutenant's Dragonhide Treads (227182, -9.83 DPS) [vendor]; Waywatcher Stompers (240066, -16.40 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (168.1 DPS) | yes | Band of the Penitent (13217, -2.52 DPS) [quest]; Ring of Entropy (18543, -2.52 DPS) [world]; Naglering (11669, -2.90 DPS, sim-verified) [dungeon] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (168.1 DPS) | yes | Band of the Penitent (13217, -2.41 DPS) [quest]; Ring of Entropy (18543, -2.41 DPS) [world]; Naglering (11669, -2.76 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (168.1 DPS) | yes | Frozen Heart of the Mountain (249469, -6.59 DPS, sim-verified) [crafted] |
| trinket2 | Darkmoon Card: Heroism (19287) | Darkmoon Warlords Deck [quest] | sim-verified (168.1 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS, sim-verified) [crafted] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-verified (168.1 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [vendor]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Sulfuron Hammer (17193, -1.58 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Moon (23197) (or Howling Idol (272427), Enraged Idol (272428), Idol of Synthesis (272429), Swarming Idol (272430), Idol of Swiftness (279250), Idol of the Ursine Twins (279251), Idol of the Dream (220606), Idol of the Huntress (227444), Idol of the Raging Shambler (220915), Talons of Wrath (249441), Idol of the Heckler (213594), Idol of the Wild (210534), Mystic Mushroom (249396), Ferocious Idol (208689), Idol of Ursine Rage (206954), Unbalanced Idol (210195), Lunar Idol (208414)) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Howling Idol (272427, +0.00 DPS, sim-verified) [vendor] |

**New at 60:** head: Waywatcher Cowl; neck: Blazefury Medallion; shoulder: Waywatcher Mantle; back: Chromatic Cloak; chest: Waywatcher Leathers; wrist: Waywatcher Wristguards; hands: Waywatcher Mitts; waist: Highlander's Leather Girdle; legs: Sentinel's Leather Pants; feet: Waywatcher Sandals; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Darkmoon Card: Maelstrom; trinket2: Darkmoon Card: Heroism; main_hand: The Unstoppable Force; ranged: Idol of the Moon

No-known-source sample (15 of 1298, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

## Horde

### Band 20 (tauren, 0000000000000000-5420000000000000000-0000000000000000)

Set DPS (verified): 61.1. Weights run: 1.4s. Verify run: 1.3s. 170 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.002, agility=1.436 ± 0.049, crit=8.959 ± 0.224, hit=1.745 ± 0.067, melee_haste=5.281 ± 0.416

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 18.6 attack_power points (1.03 DPS) | yes | Brawler's Leather Hood (252504, -0.34 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 8.6 attack_power points (0.48 DPS) | yes | Erudite's Amulet (277204, -0.17 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 7.2 attack_power points (0.40 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.43 DPS, sim-verified) [crafted] |
| back | Grave Shroud (279865) | Abominable Creatures [quest] | 9.8 attack_power points (0.55 DPS) | yes | Lambent Scale Cloak (4706, -0.04 DPS, sim-verified) [world_drop]; Dark Leather Cloak (2316, -0.05 DPS) [crafted]; Glowing Lizardscale Cloak (6449, -0.07 DPS) [dungeon] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 21.7 attack_power points (1.20 DPS) | yes | Defender's Leather Armor (252434, -0.09 DPS, sim-verified) [crafted]; Totemic Leather Armor (252435, -0.30 DPS) [crafted]; Murloc Scale Breastplate (5781, -0.32 DPS) [crafted] |
| wrist | Bristlebark Bindings (14569) | World drop [world_drop] | 8.9 attack_power points (0.50 DPS) | yes | Forest Leather Bracers (3202, -0.09 DPS, sim-verified) [world_drop]; Wolf Bracers (4794, -0.18 DPS) [vendor]; Ratchet Wristwraps (274742, -0.26 DPS) [vendor] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 125.4 attack_power points (6.95 DPS) | yes | Gloves of the Fang (10413, +0.00 DPS, sim-verified) [dungeon]; Gold-flecked Gloves (5195, -6.05 DPS) [dungeon]; Bristlebark Gloves (14572, -6.12 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (1.00 DPS) | yes | Brawler's Leather Belt (252428, -0.13 DPS, sim-verified) [crafted]; Deviate Scale Belt (6468, -0.21 DPS) [crafted]; Ruffian Belt (5975, -0.23 DPS) [world] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 26.8 attack_power points (1.49 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS, sim-verified) [crafted]; Defender's Leather Pants (252445, -0.01 DPS) [crafted]; Leggings of the Fang (10410, -0.13 DPS) [dungeon] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 18.8 attack_power points (1.04 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS, sim-verified) [world_drop]; Defender's Leather Boots (252441, -0.40 DPS) [crafted]; Totemic Leather Boots (252442, -0.40 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 15.0 attack_power points (0.83 DPS) | yes | Signet of the Zhevra (285330, -0.36 DPS) [world]; Loop of Sacrifice (281673, -0.45 DPS) [quest]; Bounty Hunter's Ring (5351, -0.59 DPS) [quest] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 9.3 attack_power points (0.51 DPS) | yes | Signet of the Zhevra (285330, +0.00 DPS, sim-verified) [world]; Loop of Sacrifice (281673, -0.13 DPS) [quest]; Bounty Hunter's Ring (5351, -0.28 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hammerbone (270018) | Leaders of the Fang [quest] | 309.4 attack_power points (17.15 DPS) | yes | Smite's Mighty Hammer (7230, +0.00 DPS, sim-verified) [dungeon]; Living Root (6631, -0.69 DPS) [dungeon]; Crescent Staff (6505, -0.81 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Grave Shroud; chest: Brawler's Leather Armor; wrist: Bristlebark Bindings; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Demon Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Hammerbone; ranged: Idol of the Huntress

No-known-source sample (15 of 170, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 20425 Advisor's Gnarled Staff

### Band 30 (tauren, 0000000000000000-5423222100000000000-0000000000000000)

Set DPS (verified): 100.4. Weights run: 1.6s. Verify run: 1.6s. 301 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.002, agility=1.494 ± 0.064, crit=10.211 ± 0.292, hit=2.054 ± 0.079, melee_haste=6.340 ± 0.693

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 27.8 attack_power points (1.45 DPS) | yes | Azure Gustwoven Hood (277050, -0.36 DPS) [crafted]; Defender's Leather Hood (252447, -0.48 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.49 DPS, sim-verified) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.3 attack_power points (0.80 DPS) | yes | Scout's Medallion (19537, -0.17 DPS) [rep]; River Pride Choker (13087, -0.31 DPS) [world_drop]; Ghostshard Talisman (7731, -0.32 DPS, sim-verified) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 28.0 attack_power points (1.46 DPS) | yes | Bristlebark Amice (14573, -0.63 DPS) [world_drop]; Mantle of Thieves (2264, -0.68 DPS) [dungeon]; Barbaric Shoulders (5964, -0.74 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 17.4 attack_power points (0.91 DPS) | yes | Tigerstrike Mantle (13108, -0.37 DPS, sim-verified) [world_drop]; Wolfmaster Cape (6314, -0.39 DPS) [dungeon]; Wildhunter Cloak (16658, -0.39 DPS) [quest] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 25.9 attack_power points (1.35 DPS) | yes | Brawler's Leather Armor (252490, -0.20 DPS) [crafted]; Defender's Leather Tunic (252450, -0.25 DPS, sim-verified) [crafted]; Dusky Leather Armor (7374, -0.26 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 18.2 attack_power points (0.95 DPS) | yes | Bands of Serra'kis (6902, -0.23 DPS) [dungeon]; Jurassic Wristguards (6198, -0.24 DPS) [world]; Barbaric Bracers (18948, -0.25 DPS, sim-verified) [crafted] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 143.0 attack_power points (7.47 DPS) | yes | Insignia Gloves (6408, +0.00 DPS, sim-verified) [world_drop]; Toughened Leather Gloves (4253, -6.27 DPS) [crafted]; Wolfclaw Gloves (1978, -6.40 DPS) [dungeon] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 29.8 attack_power points (1.56 DPS) | yes | Skulker's Leather Belt (252520, -0.29 DPS, sim-verified) [crafted]; Defiler's Chain Girdle (20152, -0.31 DPS) [rep]; Defiler's Leather Girdle (20191, -0.31 DPS) [rep] |
| legs | Brawler's Leather Legguards (252516) | Leatherworking [crafted] | 29.7 attack_power points (1.55 DPS) | yes | Trapper's Leather Pants (252501, -0.12 DPS) [crafted]; Defender's Leather Pants (252445, -0.15 DPS) [crafted]; Brawler's Leather Pants (252500, -0.16 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 19.1 attack_power points (1.00 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS, sim-verified) [world_drop]; Stomping Boots (3741, -0.20 DPS) [quest]; Disjointed Shoes (277226, -0.37 DPS) [quest] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 23.0 attack_power points (1.20 DPS) | yes | Ironspine's Eye (7686, -0.02 DPS) [dungeon]; Band of the Fist (17694, -0.41 DPS) [quest]; Tiger Band (6749, -0.48 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 22.9 attack_power points (1.20 DPS) | yes | Ironspine's Eye (7686, +0.00 DPS, sim-verified) [dungeon]; Band of the Fist (17694, -0.40 DPS) [quest]; Legionnaire's Band (20429, -0.40 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (100.4 DPS) | yes | Cobalt Crusher (7730, -2.40 DPS) [dungeon]; Wind Spirit Staff (6689, -4.03 DPS) [dungeon]; Viscous Hammer (13045, -21.55 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; waist: Prowler's Leather Belt; legs: Brawler's Leather Legguards; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 301, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5821 Darkstalker Boots; 5968 Rugged Boots

### Band 40 (tauren, 0000000000000000-5423222121032010001-0000000000000000)

Set DPS (verified): 129.5. Weights run: 1.7s. Verify run: 1.8s. 421 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.003, agility=1.534 ± 0.071, crit=11.632 ± 0.333, hit=2.271 ± 0.096, melee_haste=7.650 ± 0.975

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 42.4 attack_power points (2.35 DPS) | yes | Hawkeye's Helm (14591, -0.90 DPS) [world_drop]; Cloudy Gustwoven Hood (277042, -1.19 DPS) [crafted]; Defender's Leather Helm (252455, -1.28 DPS, sim-verified) [crafted] |
| neck | Ethereal Talisman (4430) | The Crown of Will [quest] | 17.7 attack_power points (0.98 DPS) | yes | Scout's Medallion (19536, +0.00 DPS, sim-verified) [rep]; Kaleidoscope Chain (13084, -0.13 DPS) [world_drop]; Ghostshard Talisman (7731, -0.21 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 28.9 attack_power points (1.60 DPS) | yes | Forest Tracker Epaulets (2278, -0.03 DPS, sim-verified) [world_drop]; Flintrock Shoulders (7755, -0.11 DPS) [dungeon]; Imperial Leather Spaulders (4737, -0.44 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-verified (127.2 DPS) | yes | Scorpashi Cape (14656, -0.30 DPS) [world_drop]; Imperial Cloak (6432, -0.30 DPS) [world_drop]; Dark Hooded Cape (5257, -1.42 DPS, sim-verified) [world] |
| chest | Barbaric Harness (5739) | Leatherworking [crafted] | sim-verified (128.0 DPS) | yes | Brawler's Leather Tunic (252508, -0.13 DPS) [crafted]; Wolffear Harness (13110, -0.13 DPS) [world_drop]; Quillward Harness (10583, -2.25 DPS, sim-verified) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.11 DPS) | yes | Hawkeye's Bracers (14590, -0.05 DPS, sim-verified) [world_drop]; Barbaric Bracers (18948, -0.25 DPS) [crafted]; Bands of Serra'kis (6902, -0.34 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 182.9 attack_power points (10.13 DPS) | yes | Shadowskin Gloves (18238, -1.11 DPS) [crafted]; Fletcher's Gloves (7348, -1.51 DPS, sim-verified) [crafted]; Prowler's Leather Gloves (252524, -8.08 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 34.7 attack_power points (1.92 DPS) | yes | Prowler's Leather Belt (252459, -0.25 DPS) [crafted]; Defiler's Leather Girdle (20192, -0.26 DPS) [rep]; Tharg's Shoelace (9705, -0.47 DPS, sim-verified) [quest] |
| legs | Triprunner Dungarees (9624) | Rig Wars [quest] | 34.6 attack_power points (1.91 DPS) | yes | Basilisk Hide Pants (1718, -0.15 DPS, sim-verified) [world_drop]; Brawler's Leather Legguards (252516, -0.25 DPS) [crafted]; Brawler's Leather Pants (252500, -0.38 DPS) [crafted] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 36.3 attack_power points (2.01 DPS) | yes | Excelsior Boots (4109, -0.34 DPS) [quest]; Skulker's Leather Shoes (252531, -0.35 DPS, sim-verified) [crafted]; Imperial Leather Boots (6431, -0.43 DPS) [world_drop] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 30.8 attack_power points (1.71 DPS) | yes | Legionnaire's Band (19513, -0.43 DPS) [rep]; Ironspine's Eye (7686, -0.43 DPS) [dungeon]; Falcon's Hook (7552, -0.56 DPS) [world_drop] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 23.2 attack_power points (1.28 DPS) | yes | Ironspine's Eye (7686, +0.00 DPS, sim-verified) [dungeon]; Falcon's Hook (7552, -0.13 DPS) [world_drop]; Assault Band (13095, -0.18 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (125.8 DPS) | yes | Thornstone Sledgehammer (1722, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Illusionary Rod (7713, -27.41 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; neck: Ethereal Talisman; shoulder: Sunburn Spaulders; chest: Barbaric Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Triprunner Dungarees; feet: Prowler's Leather Shoes; finger1: Legionnaire's Band; finger2: Thunderbrow Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty

No-known-source sample (15 of 421, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (tauren, 0000000000000000-5423222121032010001-5500000000000000)

Set DPS (verified): 130.7. Weights run: 1.7s. Verify run: 1.7s. 563 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.003, agility=1.597 ± 0.077, crit=12.908 ± 0.369, hit=2.557 ± 0.109, melee_haste=7.662 ± 1.230

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Leather Headband (220851) | Lady Palanseer [vendor] | 222.3 attack_power points (12.32 DPS) | yes | Eye of Theradras (17715, -1.55 DPS, sim-verified) [dungeon]; White Bandit Mask (10008, -9.93 DPS) [crafted]; Undercity Reservist's Cap (20643, -10.06 DPS) [quest] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 32.4 attack_power points (1.79 DPS) | yes | Woven Ivy Necklace (19159, -0.28 DPS, sim-verified) [quest]; Scout's Medallion (19535, -0.73 DPS) [rep]; Ethereal Talisman (4430, -0.80 DPS) [quest] |
| shoulder | Blood Guard's Leather Shoulders (220853) | Lady Palanseer [vendor] | 214.3 attack_power points (11.87 DPS) | yes | Prowler's Leather Shoulder (252534, +0.00 DPS, sim-verified) [crafted]; Failed Flying Experiment (9647, -9.70 DPS) [quest]; Skulker's Leather Shoulder (252535, -9.74 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 36.3 attack_power points (2.01 DPS) | yes | Dark Hooded Cape (5257, +0.00 DPS, sim-verified) [world]; Blisterbane Wrap (12552, -0.68 DPS) [dungeon]; Dark Phantom Cape (13122, -0.68 DPS) [world_drop] |
| chest | Stone Guard's Leather Armor (220855) | Lady Palanseer [vendor] | 224.3 attack_power points (12.43 DPS) | yes | Stone Guard's Crackling Leather Tunic (220869, -1.82 DPS, sim-verified) [vendor]; Grizzled Pelt (22274, -8.83 DPS) [quest]; Mixologist's Tunic (12793, -9.14 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 33.2 attack_power points (1.84 DPS) | yes | Prowler's Leather Bracers (252539, -0.02 DPS, sim-verified) [crafted]; Skulker's Leather Bracers (252540, -0.15 DPS) [crafted]; Pridelord Bands (14672, -0.31 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 200.7 attack_power points (11.12 DPS) | yes | First Sergeant's Leather Gauntlets (220857, -0.37 DPS, sim-verified) [vendor]; Fletcher's Gloves (7348, -1.11 DPS) [crafted]; Shadowskin Gloves (18238, -1.11 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20193) | The Defilers [rep] | 200.7 attack_power points (11.12 DPS) | yes | Defiler's Lizardhide Girdle (20174, -1.11 DPS) [rep]; Defiler's Cloth Girdle (20165, -1.14 DPS, sim-verified) [rep]; Prowler's Leather Waistguard (252473, -8.26 DPS) [crafted] |
| legs | Stone Guard's Leather Pants (220859) | Lady Palanseer [vendor] | sim-verified (130.7 DPS) | yes | Stone Guard's Crackling Leather Leggings (220865, -1.00 DPS) [vendor]; Stormshroud Pants (15057, -1.62 DPS, sim-verified) [crafted]; Stone Guard's Restored Leather Leggings (220883, -2.41 DPS) [vendor] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 47.7 attack_power points (2.64 DPS) | yes | Skulker's Leather Boots (252469, -0.08 DPS, sim-verified) [crafted]; Sandstalker Ankleguards (12470, -0.37 DPS) [dungeon]; Shadefiend Boots (11675, -0.39 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 45.6 attack_power points (2.53 DPS) | yes | White Bone Band (11862, -1.20 DPS) [quest]; Ironspine's Eye (7686, -1.21 DPS) [dungeon]; Thunderbrow Ring (13097, -1.23 DPS) [world_drop] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 37.6 attack_power points (2.08 DPS) | yes | Legionnaire's Band (19512, -0.39 DPS, sim-verified) [rep]; White Bone Band (11862, -0.75 DPS) [quest]; Ironspine's Eye (7686, -0.77 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (129.1 DPS) | yes | Frozen Heart of the Mountain (249469, -2.04 DPS) [crafted]; Smoking Heart of the Mountain (11811, -3.32 DPS, sim-verified) [crafted] |
| trinket2 | - | - |  |  |  |
| main_hand | Blight (7959) | Blacksmithing [crafted] | sim-verified (129.1 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Hammer of the Northern Wind (810, -3.58 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Blood Guard's Leather Headband; neck: Skibi's Pendant; shoulder: Blood Guard's Leather Shoulders; back: Blackveil Cape; chest: Stone Guard's Leather Armor; wrist: Deepfury Bracers; waist: Defiler's Leather Girdle; legs: Stone Guard's Leather Pants; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Guardian Talisman; main_hand: Blight

No-known-source sample (15 of 563, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (tauren, 0000000000000000-5423222121032010001-5553200000000000)

Set DPS (verified): 182.0. Weights run: 1.7s. Verify run: 1.7s. 1302 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.003, agility=1.760 ± 0.098, crit=13.720 ± 0.409, hit=3.054 ± 0.134, melee_haste=not significant (7.156 ± 1.916)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Waywatcher Cowl (240064) | Leonid Barthalomew the Revered [vendor] | sim-verified (180.2 DPS) | yes | Ragefury Eyepatch (11735, -6.82 DPS) [dungeon]; Bloodvine Lens (19998, -7.57 DPS) [crafted]; Waywatcher Hood (240072, -8.60 DPS, sim-verified) [vendor] |
| neck | Blazefury Medallion (17111) | Lord Kazzak [world] | sim-verified (171.6 DPS) | yes | Amulet of the Darkmoon (19491, +0.00 DPS) [quest]; Beads of Ogre Might (22150, +0.00 DPS) [quest]; Medallion of the Dawn (22659, -1.96 DPS, sim-verified) [quest] |
| shoulder | Waywatcher Mantle (240070) | Leonid Barthalomew the Revered [vendor] | 253.2 attack_power points (13.73 DPS) | yes | Warlord's Dragonhide Shoulders (231684, +0.00 DPS, sim-verified) [vendor]; Darkspear Pauldrons (272105, -0.74 DPS) [vendor]; Blood Guard's Leather Shoulders (220853, -1.22 DPS) [vendor] |
| back | Chromatic Cloak (18509) (or Fel Cape (279269)) | Leatherworking [crafted] | 192.1 attack_power points (10.42 DPS) | yes | Fel Cape (279269, +0.00 DPS, sim-verified) [crafted]; Howler's Furs (272414, -7.24 DPS) [vendor]; Earthweave Cloak (21187, -7.33 DPS) [quest] |
| chest | Waywatcher Leathers (240075) | Leonid Barthalomew the Revered [vendor] | sim-verified (171.6 DPS) | yes | Tunic of Undead Slaying (23089, -3.00 DPS, sim-verified) [world]; Waywatcher Vest (240067, -10.37 DPS) [vendor]; Stormshroud Armor (15056, -10.42 DPS) [crafted] |
| wrist | Waywatcher Wristguards (240084) | Leonid Barthalomew the Revered [vendor] | sim-verified (171.6 DPS) | yes | Wristwraps of Undead Slaying (23093, -1.15 DPS, sim-verified) [world]; Waywatcher Wraps (240060, -6.15 DPS) [vendor]; Waywatcher Bracers (240076, -6.74 DPS) [vendor] |
| hands | Waywatcher Mitts (240073) | Leonid Barthalomew the Revered [vendor] | 414.7 attack_power points (22.50 DPS) | yes | General's Dragonhide Grips (231688, +0.00 DPS, sim-verified) [vendor]; Stormshroud Gloves (21278, -10.42 DPS) [crafted]; Devilsaur Gauntlets (15063, -10.56 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20190) | The Defilers [rep] | 226.1 attack_power points (12.26 DPS) | yes | Defiler's Leather Girdle (20193, -0.97 DPS, sim-verified) [rep]; Belt of the Archmage (18405, -1.84 DPS) [crafted]; Defiler's Cloth Girdle (20163, -1.84 DPS) [rep] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | sim-verified (173.4 DPS) | yes | Waywatcher Legguards (240087, -1.83 DPS, sim-verified) [vendor]; Stormshroud Pants (15057, -2.58 DPS) [crafted]; Sentinel's Silk Leggings (237815, -2.58 DPS) [vendor] |
| feet | Waywatcher Sandals (240074) | Leonid Barthalomew the Revered [vendor] | 414.7 attack_power points (22.50 DPS) | yes | General's Dragonhide Treads (231683, +0.00 DPS, sim-verified) [vendor]; Blood Guard's Dragonhide Treads (227181, -9.83 DPS) [vendor]; Waywatcher Stompers (240066, -16.40 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (171.6 DPS) | yes | Band of the Penitent (13217, -2.52 DPS) [quest]; Ring of Entropy (18543, -2.52 DPS) [world]; Naglering (11669, -2.76 DPS, sim-verified) [dungeon] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (171.6 DPS) | yes | Band of the Penitent (13217, -2.41 DPS) [quest]; Ring of Entropy (18543, -2.41 DPS) [world]; Naglering (11669, -2.62 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (171.6 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (171.6 DPS) | yes | Frozen Heart of the Mountain (249469, -2.84 DPS, sim-verified) [crafted] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (171.6 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [vendor]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Sulfuron Hammer (17193, -1.15 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Moon (23197) (or Howling Idol (272427), Enraged Idol (272428), Idol of Synthesis (272429), Swarming Idol (272430), Idol of Swiftness (279250), Idol of the Ursine Twins (279251), Idol of the Dream (220606), Idol of the Huntress (227444), Idol of the Raging Shambler (220915), Talons of Wrath (249441), Idol of the Heckler (213594), Idol of the Wild (210534), Mystic Mushroom (249396), Ferocious Idol (208689), Idol of Ursine Rage (206954), Unbalanced Idol (210195), Lunar Idol (208414)) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Howling Idol (272427, +0.00 DPS, sim-verified) [vendor] |

**New at 60:** head: Waywatcher Cowl; neck: Blazefury Medallion; shoulder: Waywatcher Mantle; back: Chromatic Cloak; chest: Waywatcher Leathers; wrist: Waywatcher Wristguards; hands: Waywatcher Mitts; waist: Defiler's Leather Girdle; legs: Sentinel's Leather Pants; feet: Waywatcher Sandals; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: The Unstoppable Force; ranged: Idol of the Moon

No-known-source sample (15 of 1302, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

