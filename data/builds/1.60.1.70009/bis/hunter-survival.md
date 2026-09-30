# Leveling BiS: Survival

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

## Alliance

### Band 20 (dwarf, 0000000000000000-00000000000000000-500230100000000000)

Set DPS (verified): 57.0. Weights run: 1.4s. Verify run: 1.3s. 385 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.037 ± 0.007, strength=1.000 ± 0.001, crit=1.607 ± 0.051, hit=0.966 ± 0.082, melee_haste=not significant (0.795 ± 0.764)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.3 | yes | Defender's Leather Hood (252447, -0.05 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -0.34 DPS) [crafted]; Shadow Goggles (4373, -0.34 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.2 | yes | Tarnished Locket (279870, -0.32 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.2 | yes | Reinforced Woolen Shoulders (4315, -0.21 DPS) [crafted]; Forest Leather Mantle (4709, -0.21 DPS) [world_drop]; Double-Stitched Woolen Shoulders (4314, -0.32 DPS, sim-verified) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.2 | yes | Cape of the Brotherhood (5193, -0.04 DPS) [dungeon]; Dark Leather Cloak (2316, -0.04 DPS) [crafted]; Catacomb Cloak (279899, -0.16 DPS, sim-verified) [quest] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 12.3 | yes | Defender's Leather Armor (252434, -0.09 DPS) [crafted]; Murloc Scale Breastplate (5781, -0.17 DPS) [crafted]; Tunic of Westfall (2041, -0.19 DPS, sim-verified) [quest] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 6.1 | yes | Wolf Bracers (4794, -0.08 DPS) [vendor]; Ratchet Wristwraps (274742, -0.12 DPS) [vendor]; Forest Leather Bracers (3202, -0.20 DPS, sim-verified) [world_drop] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 22.5 | yes | Gloves of the Fang (10413, +0.28 DPS, sim-verified) [dungeon]; Brawler's Leather Gloves (252494, -0.58 DPS) [crafted]; Gold-flecked Gloves (5195, -0.63 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 | yes | Brawler's Leather Belt (252428, -0.40 DPS) [crafted]; Ruffian Belt (5975, -0.49 DPS) [world]; Deviate Scale Belt (6468, -0.55 DPS, sim-verified) [crafted] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 15.3 | yes | Trapper's Leather Pants (252501, +0.43 DPS, sim-verified) [crafted]; Leggings of the Fang (10410, -0.04 DPS) [dungeon]; Defender's Leather Pants (252445, -0.09 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 11.3 | yes | Brawler's Leather Boots (252439, -0.08 DPS, sim-verified) [crafted]; Blackened Defias Boots (10402, -0.21 DPS) [dungeon]; Footpads of the Fang (10411, -0.21 DPS) [dungeon] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.1 | yes | Lavishly Jeweled Ring (1156, -0.25 DPS) [dungeon]; The 1 Ring (8350, -0.25 DPS) [world]; Minor Channeling Ring (1449, -0.33 DPS) [quest] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.2 | yes | Lavishly Jeweled Ring (1156, +0.25 DPS, sim-verified) [dungeon]; The 1 Ring (8350, -0.17 DPS) [world]; Minor Channeling Ring (1449, -0.25 DPS) [quest] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 | yes | Duskbringer (2205, +0.00 DPS) [dungeon]; Living Root (6631, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 233.6 | yes | Grayson's Torch (1172, -9.44 DPS) [quest]; Pulsating Hydra Heart (5183, -9.44 DPS) [world]; Tear of Grief (5611, -9.44 DPS) [quest] |
| ranged | Fine Longbow (11304) | Naela Trance [vendor] | 4.0 | yes | Light Bow (4576, -0.08 DPS) [world_drop]; Owlsight Rifle (15205, -0.08 DPS) [quest]; Deadly Blunderbuss (4369, -0.10 DPS, sim-verified) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Bravo's Armbands; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Signet of the Zhevra; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Fine Longbow

No-known-source sample (15 of 385, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 8178 Training Sword; 9602 Brushwood Blade; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers; 9757 Nomad Tunic

### Band 30 (dwarf, 0000000000000000-00000000000000000-500230131051000000)

Set DPS (verified): 87.4. Weights run: 1.4s. Verify run: 1.5s. 734 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.052 ± 0.014, strength=1.000 ± 0.001, crit=2.688 ± 0.095, hit=3.081 ± 0.212, melee_haste=not significant (0.562 ± 0.827)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 12.0 | yes | Brawler's Leather Helm (252512, -0.03 DPS, sim-verified) [crafted]; Cloudy Gustwoven Hood (277042, -0.14 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.14 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19541, -0.27 DPS, sim-verified) [rep]; Sentinel's Medallion (20444, -0.35 DPS) [rep]; Pendant of Myzrael (4614, -0.64 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 16.6 | yes | Barbaric Shoulders (5964, -0.29 DPS) [crafted]; Mantle of Thieves (2264, -0.32 DPS, sim-verified) [dungeon]; Dark Leather Shoulders (4252, -0.42 DPS) [crafted] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Sergeant Major's Cape (16315, +0.36 DPS, sim-verified) [pvp]; Cloak of Night (4447, -0.17 DPS) [world]; Fenrus' Hide (6340, -0.17 DPS) [dungeon] |
| chest | Raptorbane Armor (3566) | Ormer's Revenge [quest] | 16.0 | yes | Dusky Leather Armor (7374, -0.00 DPS, sim-verified) [crafted]; Brawler's Leather Tunic (252508, -0.07 DPS) [crafted]; Brawler's Leather Armor (252490, -0.17 DPS) [crafted] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.0 | yes | Jurassic Wristguards (6198, -0.06 DPS, sim-verified) [world]; Barbaric Bracers (18948, -0.08 DPS) [crafted]; Insignia Bracers (6410, -0.17 DPS) [world_drop] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 37.6 | yes | Heavy Earthen Gloves (7359, +0.54 DPS, sim-verified) [crafted]; Insignia Gloves (6408, -1.12 DPS) [world_drop]; Toughened Leather Gloves (4253, -1.16 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 | yes | Highlander's Leather Girdle (20117, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.28 DPS) [dungeon]; Skulker's Leather Belt (252520, -0.39 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Brawler's Leather Pants (252500, -0.48 DPS) [crafted]; Trapper's Leather Pants (252501, -0.48 DPS) [crafted]; Brawler's Leather Legguards (252516, -1.93 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 11.4 | yes | Brawler's Leather Boots (252439, -0.08 DPS, sim-verified) [crafted]; Insignia Boots (4055, -0.14 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.14 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 13.5 | yes | Insurgent's Band (272067, -0.21 DPS) [vendor]; Monkey Ring (6748, -0.28 DPS) [quest]; Ring of Precision (1491, -0.33 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 12.3 | yes | Protector's Band (20439, -0.19 DPS) [rep]; Insurgent's Band (272067, -0.21 DPS, sim-verified) [vendor]; Monkey Ring (6748, -0.23 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 340.4 | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Morbid Dawn (7689, +0.00 DPS) [dungeon]; Darkspear Raider's Reaper (272082, +0.00 DPS) [vendor] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 318.3 | yes | Shoni's Disarming Tool (9608, -4.30 DPS) [quest]; Grayson's Torch (1172, -14.64 DPS) [quest]; Rod of Molten Fire (2565, -14.64 DPS) [world_drop] |
| ranged | Moonsight Rifle (4383) (or Precision Bow (217315)) | Engineering [crafted] | 9.2 | yes | Precision Bow (217315, +0.00 DPS, sim-verified) [quest]; Booty Bay Bruiser's Buckshot (274748, -0.01 DPS) [vendor]; Double-barreled Shotgun (2098, -0.14 DPS) [world_drop] |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Raptorbane Armor; wrist: Cultist's Armguards; waist: Highlander's Chain Girdle; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Protector's Band; trinket1: Darkspear Voodoo Seal; trinket2: Talisman of Arathor; main_hand: Pronged Reaver; off_hand: Bloody Brass Knuckles; ranged: Moonsight Rifle

No-known-source sample (15 of 734, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 7957 Bronze Greatsword; 7958 Bronze Battle Axe

### Band 40 (dwarf, 0000000000000000-00000000000000000-500230131051120151)

Set DPS (verified): 123.1. Weights run: 1.5s. Verify run: 1.5s. 1260 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.169 ± 0.018, strength=1.000 ± 0.001, crit=3.231 ± 0.106, hit=3.490 ± 0.282, melee_haste=not significant (3.178 ± 1.330)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 58.2 | yes | White Bandit Mask (10008, +0.13 DPS, sim-verified) [crafted]; Barbaric Iron Helm (7915, -2.04 DPS) [crafted]; Hawkeye's Helm (14591, -2.18 DPS) [world] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19540, +0.07 DPS, sim-verified) [rep]; Sentinel's Medallion (19541, -0.24 DPS) [rep]; Sentinel's Medallion (20444, -0.37 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.9 | yes | Barbaric Iron Shoulders (7913, -0.62 DPS) [crafted]; Nightscape Shoulders (8192, -0.63 DPS) [crafted]; Forest Tracker Epaulets (2278, -0.69 DPS, sim-verified) [world_drop] |
| back | Sergeant Major's Cape (16336) | Rank 9 [pvp] | 13.0 | yes | Wolfmaster Cape (6314, -0.16 DPS) [dungeon]; Imperial Cloak (6432, -0.19 DPS) [world_drop]; Yeti Fur Cloak (2805, -1.56 DPS, sim-verified) [quest] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 22.2 | yes | Tough Scorpid Breastplate (8203, -0.24 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.27 DPS) [crafted]; Nightscape Tunic (8175, -0.32 DPS, sim-verified) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Cultist's Armguards (270032, -0.53 DPS) [quest]; Imperial Leather Bracers (4061, -0.56 DPS) [world_drop]; Ravager's Armguards (14770, -0.77 DPS, sim-verified) [world] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 65.2 | yes | Fletcher's Gloves (7348, -1.05 DPS) [crafted]; Shadowskin Gloves (18238, -1.05 DPS) [crafted]; Dragonscale Gauntlets (8347, -1.17 DPS, sim-verified) [crafted] |
| waist | Highlander's Chain Girdle (20089) | The League of Arathor [rep] | 53.2 | yes | Highlander's Leather Girdle (20116, +0.98 DPS, sim-verified) [rep]; Highlander's Chain Girdle (20090, -1.54 DPS) [rep]; Highlander's Leather Girdle (20117, -1.54 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Triprunner Dungarees (9624, -0.10 DPS) [quest]; Scarlet Leggings (10330, -0.26 DPS) [dungeon]; Basilisk Hide Pants (1718, -1.57 DPS, sim-verified) [world_drop] |
| feet | Blackforge Greaves (6423) | World drop [world_drop] | 20.7 | yes | Skulker's Leather Shoes (252531, -0.02 DPS, sim-verified) [crafted]; Prowler's Leather Shoes (252465, -0.08 DPS) [crafted]; Imperial Leather Boots (6431, -0.15 DPS) [world_drop] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 17.4 | yes | Ring of the Underwood (2951, -0.19 DPS) [world_drop]; Protector's Band (19517, -0.23 DPS) [rep]; Insurgent's Band (272066, -0.28 DPS) [vendor] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 14.5 | yes | Ring of the Underwood (2951, -0.06 DPS, sim-verified) [world_drop]; Insurgent's Band (272066, -0.13 DPS) [vendor]; Disengagement Ring (276202, -0.27 DPS) [vendor] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | 442.0 | yes | Fiery War Axe (870, +0.00 DPS) [world_drop]; Staff of Jordan (873, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 440.2 | yes | Shoni's Disarming Tool (9608, -11.35 DPS) [quest]; Stonecloth Branch (15963, -23.02 DPS) [world]; Grayson's Torch (1172, -23.18 DPS) [quest] |
| ranged | Moonsight Rifle (4383) (or Precision Bow (217315)) | Engineering [crafted] | 10.5 | yes | Precision Bow (217315, +0.00 DPS, sim-verified) [quest]; Booty Bay Bruiser's Buckshot (274748, -0.08 DPS) [vendor]; Master Hunter's Rifle (17687, -0.09 DPS) [quest] |

**New at 40:** head: Raging Berserker's Helm; shoulder: Sunburn Spaulders; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Chain Girdle; feet: Blackforge Greaves; finger1: Protector's Band; finger2: Ironspine's Eye; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Vanquisher's Sword; off_hand: Jhordy's Misplaced Screwdriver

No-known-source sample (15 of 1260, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 50 (dwarf, 0000000000000000-32005000000000000-500230131051120151)

Set DPS (verified): 154.6. Weights run: 1.5s. Verify run: 1.5s. 1602 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.177 ± 0.018, strength=1.000 ± 0.001, crit=4.300 ± 0.154, hit=5.121 ± 0.388, melee_haste=not significant (1.910 ± 0.883)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Chain Helmet (220821) (or Knight-Lieutenant's Chain Helmet (220822)) | Lady Palanseer [vendor] | 79.0 | yes | Knight-Lieutenant's Chain Helmet (220822, +0.00 DPS, sim-verified) [vendor]; Raging Berserker's Helm (7719, -0.30 DPS) [dungeon]; Blood Guard's Mail Helmet (220820, -0.41 DPS) [vendor] |
| neck | Sentinel's Medallion (19539) | Silverwing Sentinels [rep] | 14.1 | yes | Sentinel's Medallion (19540, -0.06 DPS) [rep]; Ghostshard Talisman (7731, -0.15 DPS, sim-verified) [dungeon]; Sentinel's Medallion (19541, -0.24 DPS) [rep] |
| shoulder | Blood Guard's Chain Epaulets (220824) (or Knight-Lieutenant's Chain Epaulets (220825)) | Lady Palanseer [vendor] | 75.5 | yes | Knight-Lieutenant's Chain Epaulets (220825, +0.00 DPS, sim-verified) [vendor]; Blood Guard's Mail Epaulets (220823, -0.33 DPS) [vendor]; Knight-Lieutenant's Mail Epaulets (223073, -0.33 DPS) [vendor] |
| back | Sergeant Major's Cape (16336) | Rank 9 [pvp] | 13.1 | yes | Serpentskin Cloak (8259, -0.01 DPS) [dungeon]; Nightscape Cloak (8195, -0.07 DPS) [crafted]; Pridelord Cape (14673, -1.38 DPS, sim-verified) [dungeon] |
| chest | Stone Guard's Chain Armor (220827) (or Knight's Chain Armor (220828)) | Lady Palanseer [vendor] | 77.8 | yes | Knight's Chain Armor (220828, +0.00 DPS, sim-verified) [vendor]; Stone Guard's Mail Armor (220826, -0.29 DPS) [vendor]; Knight's Mail Armor (223078, -0.29 DPS) [vendor] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 | yes | Branded Leather Bracers (19508, -0.41 DPS) [dungeon]; Skulker's Leather Bracers (252540, -0.54 DPS) [crafted]; Deepfury Bracers (13120, -1.22 DPS, sim-verified) [world] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 80.2 | yes | First Sergeant's Mail Gauntlets (220831, -0.25 DPS, sim-verified) [vendor]; Sergeant Major's Mail Gauntlets (223076, -0.62 DPS) [vendor]; Dragonscale Gauntlets (8347, -0.67 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20088) (or Highlander's Leather Girdle (20115)) | The League of Arathor [rep] | 80.2 | yes | Highlander's Leather Girdle (20115, +0.00 DPS, sim-verified) [rep]; Highlander's Chain Girdle (20089, -0.62 DPS) [rep]; Highlander's Cloth Girdle (20097, -1.04 DPS) [rep] |
| legs | Knight's Chain Legplates (220832) | Captain Dirgehammer [vendor] | 76.7 | yes | Stone Guard's Chain Legplates (220833, +0.00 DPS) [vendor]; Stone Guard's Mail Legplates (220834, -0.23 DPS) [vendor]; Stormshroud Pants (15057, -2.21 DPS, sim-verified) [crafted] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 51.2 | yes | Skulker's Leather Boots (252469, -0.25 DPS, sim-verified) [crafted]; Sandstalker Ankleguards (12470, -1.31 DPS) [dungeon]; Prowler's Leather Boots (252468, -1.31 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 71.2 | yes | Masons Fraternity Ring (9533, -2.83 DPS) [quest]; Insurgent's Band (272065, -2.91 DPS) [vendor]; Ironspine's Eye (7686, -2.93 DPS) [dungeon] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 20.6 | yes | Masons Fraternity Ring (9533, -0.21 DPS) [quest]; Insurgent's Band (272065, -0.29 DPS) [vendor]; Protector's Band (19515, -0.31 DPS, sim-verified) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 46.1 | yes | Thunderbrew's Boot Flask (744, -2.39 DPS) [quest]; Tidal Charm (1404, -2.39 DPS) [vendor]; Guardian Talisman (1490, -2.39 DPS) [quest] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.01 DPS, sim-verified) [quest] |
| main_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | 570.2 | yes | Kindling Stave (11750, +0.00 DPS) [dungeon]; Bleakwood Hew (12769, +0.00 DPS) [crafted]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | 526.4 | yes | Claw of Celebras (17738, -2.29 DPS) [dungeon]; Shoni's Disarming Tool (9608, -15.62 DPS) [quest]; Thermotastic Egg Timer (9644, -27.07 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | 16.5 | yes | Moonsight Rifle (4383, -0.06 DPS) [crafted]; Precision Bow (217315, -0.06 DPS) [quest]; Houndmaster's Bow (11628, -0.08 DPS) [dungeon] |

**New at 50:** head: Blood Guard's Chain Helmet; neck: Sentinel's Medallion; shoulder: Blood Guard's Chain Epaulets; chest: Stone Guard's Chain Armor; wrist: Bracers of the Stone Princess; waist: Highlander's Chain Girdle; legs: Knight's Chain Legplates; feet: Greaves of Withering Despair; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Frozen Heart of the Mountain; main_hand: Dawn's Edge; off_hand: Thorium Cestus; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1602, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (dwarf, 0000000000000000-32005500005000000-500230131051120151)

Set DPS (verified): 241.9. Weights run: 1.5s. Verify run: 1.7s. 2288 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.205 ± 0.024, strength=1.000 ± 0.001, crit=6.402 ± 0.250, hit=5.883 ± 0.716, melee_haste=not significant (0.393 ± 1.319)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Champion's Chain Greathelm (227080) (or Lieutenant Commander's Chain Greathelm (227086)) | Lady Palanseer [vendor] | 227.5 | yes | Lieutenant Commander's Chain Greathelm (227086, +0.00 DPS, sim-verified) [vendor]; Cryptstalker Headpiece (22438, -0.55 DPS) [quest]; Mask of the Unforgiven (13404, -1.02 DPS) [dungeon] |
| neck | Blazefury Medallion (17111) | Lord Kazzak [world] | 15.7 | yes | Onyxia Tooth Pendant (18404, +0.26 DPS, sim-verified) [quest]; Beads of Ogre Might (22150, +0.00 DPS) [quest]; Medallion of the Dawn (22659, +0.00 DPS) [quest] |
| shoulder | Champion's Chain Pauldrons (227078) (or Lieutenant Commander's Chain Pauldrons (227084)) | Rank 14 [pvp] | 129.5 | yes | Lieutenant Commander's Chain Pauldrons (227084, +0.00 DPS, sim-verified) [pvp]; Cryptstalker Spaulders (22439, -0.25 DPS) [quest]; Champion's Chain Shoulders (23252, -0.92 DPS) [vendor] |
| back | Cloak of the Unseen Path (21403) | Cloak of the Unseen Path [quest] | 79.3 | yes | Earthweave Cloak (21187, -0.12 DPS) [quest]; Cloak of the Fallen God (21710, -1.86 DPS) [quest]; Chromatic Cloak (18509, -2.37 DPS, sim-verified) [crafted] |
| chest | Legionnaire's Chain Armor (227083) (or Knight-Captain's Chain Armor (227089)) | Lady Palanseer [vendor] | 220.3 | yes | Knight-Captain's Chain Armor (227089, +0.00 DPS, sim-verified) [vendor]; Cryptstalker Tunic (22436, -0.89 DPS) [quest]; Legionnaire's Chain Hauberk (22874, -1.10 DPS) [vendor] |
| wrist | Cryptstalker Wristguards (22443) | Cryptstalker Wristguards [quest] | 90.1 | yes | Rockfury Bracers (21186, -1.58 DPS) [quest]; Primal Batskin Bracers (19687, -2.14 DPS, sim-verified) [crafted]; Windtalker's Wristguards (19582, -2.63 DPS) [rep] |
| hands | Marshal's Chain Grips (16463) | Captain Dirgehammer [vendor] | 114.9 | yes | Chromatic Gauntlets (19157, +0.00 DPS) [crafted]; Stormshroud Gloves (21278, +0.00 DPS) [crafted]; Marshal's Chain Grips (231560, +0.00 DPS) [pvp] |
| waist | Cryptstalker Girdle (22442) | Cryptstalker Girdle [quest] | 176.2 | yes | Highlander's Leather Girdle (20045, -2.65 DPS) [rep]; Light Obsidian Belt (22195, -2.75 DPS) [crafted]; Highlander's Chain Girdle (20043, -6.36 DPS, sim-verified) [rep] |
| legs | Legionnaire's Chain Legplates (227079) (or Knight-Captain's Chain Legplates (227085)) | Lady Palanseer [vendor] | 220.3 | yes | Knight-Captain's Chain Legplates (227085, +0.00 DPS, sim-verified) [vendor]; Legionnaire's Chain Legguards (22875, -1.10 DPS) [vendor]; Knight-Captain's Chain Legguards (23293, -1.10 DPS) [vendor] |
| feet | General's Chain Sabatons (231564) | Rank 16 [pvp] | 115.5 | yes | Marshal's Chain Boots (16462, -1.28 DPS) [vendor]; General's Chain Sabatons (16569, -1.28 DPS) [vendor]; Cryptstalker Boots (22440, -1.30 DPS, sim-verified) [quest] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 164.5 | yes | Band of the Penitent (13217, -3.77 DPS) [quest]; Dragonslayer's Signet (18403, -3.77 DPS) [quest]; Band of Earthen Might (21182, -6.29 DPS, sim-verified) [quest] |
| finger2 | Master Dragonslayer's Ring (19384) | The Lord of Blackrock [quest] | 106.8 | yes | Band of the Penitent (13217, -0.87 DPS) [quest]; Dragonslayer's Signet (18403, -0.87 DPS) [quest]; Band of Earthen Might (21182, -2.94 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 52.9 | yes | Ankh of Life (1713, -0.78 DPS, sim-verified) [world_drop]; Tidal Charm (1404, -2.67 DPS) [vendor]; Guardian Talisman (1490, -2.67 DPS) [quest] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | 0.0 | yes | Ankh of Life (1713, +0.57 DPS, sim-verified) [world_drop]; Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest] |
| main_hand | High Warlord's Quickblade (234553) | Rank 18 [pvp] | 984.8 | yes | Atiesh, Greatstaff of the Guardian (22632, +0.00 DPS) [quest]; High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Battle Axe (234543, +0.00 DPS) [pvp] |
| off_hand | Grand Marshal's Swiftblade (234579) | Rank 18 [pvp] | 984.8 | yes | High Warlord's Left Claw (234558, -0.16 DPS) [pvp]; Grand Marshal's Left Hand Blade (234584, -0.16 DPS) [pvp]; Grand Marshal's Left Hand Blade (18847, -1.74 DPS) [vendor] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | 89.6 | yes | Fahrad's Reloading Repeater (22347, -1.31 DPS) [quest]; Core Marksman Rifle (18282, -1.55 DPS) [crafted]; Blackcrow (12651, -1.67 DPS) [dungeon] |

**New at 60:** head: Champion's Chain Greathelm; neck: Blazefury Medallion; shoulder: Champion's Chain Pauldrons; back: Cloak of the Unseen Path; chest: Legionnaire's Chain Armor; wrist: Cryptstalker Wristguards; hands: Marshal's Chain Grips; waist: Cryptstalker Girdle; legs: Legionnaire's Chain Legplates; feet: General's Chain Sabatons; finger1: Don Julio's Band; finger2: Master Dragonslayer's Ring; trinket2: Thunderbrew's Boot Flask; main_hand: High Warlord's Quickblade; off_hand: Grand Marshal's Swiftblade; ranged: The Purifier

No-known-source sample (15 of 2288, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

## Horde

### Band 20 (troll, 0000000000000000-00000000000000000-500230100000000000)

Set DPS (verified): 56.9. Weights run: 1.4s. Verify run: 1.2s. 383 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.037 ± 0.007, strength=1.000 ± 0.001, crit=1.607 ± 0.051, hit=0.966 ± 0.082, melee_haste=not significant (0.795 ± 0.764)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.3 | yes | Defender's Leather Hood (252447, -0.06 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -0.34 DPS) [crafted]; Shadow Goggles (4373, -0.34 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.2 | yes | Tarnished Locket (279870, -0.27 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.2 | yes | Reinforced Woolen Shoulders (4315, -0.21 DPS) [crafted]; Forest Leather Mantle (4709, -0.21 DPS) [world_drop]; Double-Stitched Woolen Shoulders (4314, -0.33 DPS, sim-verified) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.2 | yes | Catacomb Cloak (279899, +0.06 DPS, sim-verified) [quest]; Cape of the Brotherhood (5193, -0.04 DPS) [dungeon]; Dark Leather Cloak (2316, -0.04 DPS) [crafted] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 12.3 | yes | Defender's Leather Armor (252434, -0.14 DPS, sim-verified) [crafted]; Murloc Scale Breastplate (5781, -0.17 DPS) [crafted]; Trapper's Leather Armor (252491, -0.20 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.2 | yes | Wolf Bracers (4794, -0.06 DPS, sim-verified) [vendor]; Ratchet Wristwraps (274742, -0.08 DPS) [vendor]; Spare Part Bindings (279875, -0.08 DPS) [quest] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 22.5 | yes | Gloves of the Fang (10413, +0.27 DPS, sim-verified) [dungeon]; Brawler's Leather Gloves (252494, -0.58 DPS) [crafted]; Gold-flecked Gloves (5195, -0.63 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 | yes | Brawler's Leather Belt (252428, -0.40 DPS) [crafted]; Ruffian Belt (5975, -0.49 DPS) [world]; Deviate Scale Belt (6468, -0.55 DPS, sim-verified) [crafted] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 15.3 | yes | Trapper's Leather Pants (252501, +0.47 DPS, sim-verified) [crafted]; Leggings of the Fang (10410, -0.04 DPS) [dungeon]; Defender's Leather Pants (252445, -0.09 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 11.3 | yes | Brawler's Leather Boots (252439, -0.08 DPS, sim-verified) [crafted]; Blackened Defias Boots (10402, -0.21 DPS) [dungeon]; Footpads of the Fang (10411, -0.21 DPS) [dungeon] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.1 | yes | Bounty Hunter's Ring (5351, -0.20 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.25 DPS) [dungeon]; The 1 Ring (8350, -0.25 DPS) [world] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.2 | yes | Bounty Hunter's Ring (5351, -0.10 DPS, sim-verified) [quest]; Lavishly Jeweled Ring (1156, -0.17 DPS) [dungeon]; The 1 Ring (8350, -0.17 DPS) [world] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 | yes | Duskbringer (2205, +0.00 DPS) [dungeon]; Living Root (6631, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 233.6 | yes | Tork Wrench (11855, -9.36 DPS) [quest]; Grayson's Torch (1172, -9.44 DPS) [quest]; Nightglow Concoction (3451, -9.44 DPS) [quest] |
| ranged | Fine Longbow (11304) | Naela Trance [vendor] | 4.0 | yes | Light Bow (4576, -0.08 DPS) [world_drop]; Privateer Musket (5309, -0.08 DPS) [quest]; Deadly Blunderbuss (4369, -0.11 DPS, sim-verified) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Signet of the Zhevra; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Fine Longbow

No-known-source sample (15 of 383, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2879 Antipodean Rod; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 8178 Training Sword; 9602 Brushwood Blade; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers

### Band 30 (troll, 0000000000000000-00000000000000000-500230131051000000)

Set DPS (verified): 87.3. Weights run: 1.4s. Verify run: 1.4s. 734 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.052 ± 0.014, strength=1.000 ± 0.001, crit=2.688 ± 0.095, hit=3.081 ± 0.212, melee_haste=not significant (0.562 ± 0.827)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 12.0 | yes | Brawler's Leather Helm (252512, -0.04 DPS, sim-verified) [crafted]; Cloudy Gustwoven Hood (277042, -0.14 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.14 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19537, -0.29 DPS, sim-verified) [rep]; Scout's Medallion (20442, -0.35 DPS) [rep]; Pendant of Myzrael (4614, -0.64 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 16.6 | yes | Mantle of Thieves (2264, -0.26 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.29 DPS) [crafted]; Dark Leather Shoulders (4252, -0.42 DPS) [crafted] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Sergeant Major's Cape (16315, -0.08 DPS) [pvp]; Cloak of Night (4447, -0.17 DPS) [world] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 14.7 | yes | Brawler's Leather Tunic (252508, +0.23 DPS, sim-verified) [crafted]; Brawler's Leather Armor (252490, -0.11 DPS) [crafted]; Defender's Leather Tunic (252450, -0.11 DPS) [crafted] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.0 | yes | Jurassic Wristguards (6198, -0.07 DPS, sim-verified) [world]; Barbaric Bracers (18948, -0.08 DPS) [crafted]; Insignia Bracers (6410, -0.17 DPS) [world_drop] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 37.6 | yes | Heavy Earthen Gloves (7359, +0.54 DPS, sim-verified) [crafted]; Insignia Gloves (6408, -1.12 DPS) [world_drop]; Toughened Leather Gloves (4253, -1.16 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 | yes | Defiler's Leather Girdle (20191, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.28 DPS) [dungeon]; Deftkin Belt (16659, -0.36 DPS) [quest] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Brawler's Leather Pants (252500, -0.48 DPS) [crafted]; Trapper's Leather Pants (252501, -0.48 DPS) [crafted]; Brawler's Leather Legguards (252516, -1.64 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 11.4 | yes | Brawler's Leather Boots (252439, -0.09 DPS, sim-verified) [crafted]; Insignia Boots (4055, -0.14 DPS) [world_drop]; Warsong Boots (16977, -0.14 DPS) [quest] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 13.5 | yes | Insurgent's Band (272067, -0.21 DPS) [vendor]; Band of the Fist (17694, -0.24 DPS) [quest]; Monkey Ring (6748, -0.28 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 12.3 | yes | Band of the Fist (17694, -0.19 DPS) [quest]; Legionnaire's Band (20429, -0.19 DPS) [rep]; Insurgent's Band (272067, -0.22 DPS, sim-verified) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 340.4 | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Morbid Dawn (7689, +0.00 DPS) [dungeon]; Darkspear Raider's Reaper (272082, +0.00 DPS) [vendor] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 318.3 | yes | Tork Wrench (11855, -14.54 DPS) [quest]; Grayson's Torch (1172, -14.64 DPS) [quest]; Rod of Molten Fire (2565, -14.64 DPS) [world_drop] |
| ranged | Moonsight Rifle (4383) (or Precision Bow (217315)) | Engineering [crafted] | 9.2 | yes | Precision Bow (217315, +0.00 DPS, sim-verified) [quest]; Booty Bay Bruiser's Buckshot (274748, -0.01 DPS) [vendor]; Double-barreled Shotgun (2098, -0.14 DPS) [world_drop] |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Dusky Leather Armor; wrist: Cultist's Armguards; waist: Defiler's Chain Girdle; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Legionnaire's Band; trinket1: Darkspear Voodoo Seal; trinket2: Defiler's Talisman; main_hand: Pronged Reaver; off_hand: Bloody Brass Knuckles; ranged: Moonsight Rifle

No-known-source sample (15 of 734, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 7957 Bronze Greatsword

### Band 40 (troll, 0000000000000000-00000000000000000-500230131051120151)

Set DPS (verified): 123.1. Weights run: 1.5s. Verify run: 1.5s. 1258 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.169 ± 0.018, strength=1.000 ± 0.001, crit=3.231 ± 0.106, hit=3.490 ± 0.282, melee_haste=not significant (3.178 ± 1.330)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 58.2 | yes | White Bandit Mask (10008, +0.12 DPS, sim-verified) [crafted]; Barbaric Iron Helm (7915, -2.04 DPS) [crafted]; Hawkeye's Helm (14591, -2.18 DPS) [world] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19536, +0.09 DPS, sim-verified) [rep]; Ethereal Talisman (4430, -0.23 DPS) [quest]; Scout's Medallion (19537, -0.24 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.9 | yes | Barbaric Iron Shoulders (7913, -0.62 DPS) [crafted]; Nightscape Shoulders (8192, -0.63 DPS) [crafted]; Forest Tracker Epaulets (2278, -0.68 DPS, sim-verified) [world_drop] |
| back | Sergeant Major's Cape (16336) | Rank 9 [pvp] | 13.0 | yes | Wildhunter Cloak (16658, -0.16 DPS) [quest]; Imperial Cloak (6432, -0.19 DPS) [world_drop]; Wolfmaster Cape (6314, -1.63 DPS, sim-verified) [dungeon] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 22.2 | yes | Tough Scorpid Breastplate (8203, -0.24 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.27 DPS) [crafted]; Nightscape Tunic (8175, -0.31 DPS, sim-verified) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Cultist's Armguards (270032, -0.53 DPS) [quest]; Imperial Leather Bracers (4061, -0.56 DPS) [world_drop]; Ravager's Armguards (14770, -0.73 DPS, sim-verified) [world] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 65.2 | yes | Fletcher's Gloves (7348, -1.05 DPS) [crafted]; Shadowskin Gloves (18238, -1.05 DPS) [crafted]; Dragonscale Gauntlets (8347, -1.15 DPS, sim-verified) [crafted] |
| waist | Defiler's Chain Girdle (20153) | The Defilers [rep] | 53.2 | yes | Defiler's Leather Girdle (20192, +0.98 DPS, sim-verified) [rep]; Defiler's Chain Girdle (20152, -1.54 DPS) [rep]; Defiler's Leather Girdle (20191, -1.54 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Triprunner Dungarees (9624, -0.10 DPS) [quest]; Scarlet Leggings (10330, -0.26 DPS) [dungeon]; Basilisk Hide Pants (1718, -1.16 DPS, sim-verified) [world_drop] |
| feet | Blackforge Greaves (6423) | World drop [world_drop] | 20.7 | yes | Prowler's Leather Shoes (252465, -0.08 DPS) [crafted]; Imperial Leather Boots (6431, -0.15 DPS) [world_drop]; Skulker's Leather Shoes (252531, -0.38 DPS, sim-verified) [crafted] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 17.4 | yes | Ring of the Underwood (2951, -0.19 DPS) [world_drop]; Legionnaire's Band (19513, -0.23 DPS) [rep]; Insurgent's Band (272066, -0.28 DPS) [vendor] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 14.5 | yes | Ring of the Underwood (2951, -0.07 DPS, sim-verified) [world_drop]; Insurgent's Band (272066, -0.13 DPS) [vendor]; Disengagement Ring (276202, -0.27 DPS) [vendor] |
| trinket1 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 442.0 | yes | Fiery War Axe (870, +0.00 DPS) [world_drop]; Staff of Jordan (873, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 440.2 | yes | Stonecloth Branch (15963, -23.02 DPS) [world]; Tork Wrench (11855, -23.08 DPS) [quest]; Grayson's Torch (1172, -23.18 DPS) [quest] |
| ranged | Moonsight Rifle (4383) (or Precision Bow (217315)) | Engineering [crafted] | 10.5 | yes | Precision Bow (217315, +0.00 DPS, sim-verified) [quest]; Booty Bay Bruiser's Buckshot (274748, -0.08 DPS) [vendor]; Master Hunter's Rifle (17687, -0.09 DPS) [quest] |

**New at 40:** head: Raging Berserker's Helm; shoulder: Sunburn Spaulders; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Chain Girdle; feet: Blackforge Greaves; finger1: Legionnaire's Band; finger2: Ironspine's Eye; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Vanquisher's Sword; off_hand: Jhordy's Misplaced Screwdriver

No-known-source sample (15 of 1258, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 50 (troll, 0000000000000000-32005000000000000-500230131051120151)

Set DPS (verified): 160.4. Weights run: 1.5s. Verify run: 1.5s. 1601 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.177 ± 0.018, strength=1.000 ± 0.001, crit=4.300 ± 0.154, hit=5.121 ± 0.388, melee_haste=not significant (1.910 ± 0.883)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Chain Helmet (220821) (or Knight-Lieutenant's Chain Helmet (220822)) | Lady Palanseer [vendor] | 79.0 | yes | Knight-Lieutenant's Chain Helmet (220822, +0.00 DPS, sim-verified) [vendor]; Raging Berserker's Helm (7719, -0.30 DPS) [dungeon]; Blood Guard's Mail Helmet (220820, -0.41 DPS) [vendor] |
| neck | Woven Ivy Necklace (19159) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 16.6 | yes | Ghostshard Talisman (7731, -0.13 DPS) [dungeon]; Scout's Medallion (19535, -0.19 DPS, sim-verified) [rep]; Scout's Medallion (19536, -0.19 DPS) [rep] |
| shoulder | Blood Guard's Chain Epaulets (220824) (or Knight-Lieutenant's Chain Epaulets (220825)) | Lady Palanseer [vendor] | 75.5 | yes | Knight-Lieutenant's Chain Epaulets (220825, +0.00 DPS, sim-verified) [vendor]; Blood Guard's Mail Epaulets (220823, -0.33 DPS) [vendor]; Knight-Lieutenant's Mail Epaulets (223073, -0.33 DPS) [vendor] |
| back | Sergeant Major's Cape (16336) | Rank 9 [pvp] | 13.1 | yes | Serpentskin Cloak (8259, -0.01 DPS) [dungeon]; Nightscape Cloak (8195, -0.07 DPS) [crafted]; Pridelord Cape (14673, -1.56 DPS, sim-verified) [dungeon] |
| chest | Stone Guard's Chain Armor (220827) (or Knight's Chain Armor (220828)) | Lady Palanseer [vendor] | 77.8 | yes | Knight's Chain Armor (220828, +0.00 DPS, sim-verified) [vendor]; Stone Guard's Mail Armor (220826, -0.29 DPS) [vendor]; Knight's Mail Armor (223078, -0.29 DPS) [vendor] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 | yes | Branded Leather Bracers (19508, -0.41 DPS) [dungeon]; Skulker's Leather Bracers (252540, -0.54 DPS) [crafted]; Deepfury Bracers (13120, -0.57 DPS, sim-verified) [world] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 80.2 | yes | First Sergeant's Mail Gauntlets (220831, -0.02 DPS, sim-verified) [vendor]; Sergeant Major's Mail Gauntlets (223076, -0.62 DPS) [vendor]; Dragonscale Gauntlets (8347, -0.67 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20151) (or Defiler's Leather Girdle (20193)) | The Defilers [rep] | 80.2 | yes | Defiler's Leather Girdle (20193, +0.00 DPS, sim-verified) [rep]; Defiler's Chain Girdle (20153, -0.62 DPS) [rep]; Highlander's Mail Girdle (20118, -1.04 DPS) [vendor] |
| legs | Knight's Chain Legplates (220832) | Captain Dirgehammer [vendor] | 76.7 | yes | Stone Guard's Chain Legplates (220833, +0.00 DPS) [vendor]; Stone Guard's Mail Legplates (220834, -0.23 DPS) [vendor]; Stormshroud Pants (15057, -3.20 DPS, sim-verified) [crafted] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 51.2 | yes | Skulker's Leather Boots (252469, +1.18 DPS, sim-verified) [crafted]; Sandstalker Ankleguards (12470, -1.31 DPS) [dungeon]; Prowler's Leather Boots (252468, -1.31 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 71.2 | yes | Legionnaire's Band (19511, -2.62 DPS) [rep]; Legionnaire's Band (19512, -2.79 DPS) [rep]; Masons Fraternity Ring (9533, -2.83 DPS) [quest] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 | yes | Legionnaire's Band (19511, -0.19 DPS, sim-verified) [rep]; Legionnaire's Band (19512, -0.34 DPS) [rep]; Masons Fraternity Ring (9533, -0.39 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 77.8 | yes | Ankh of Life (1713, -3.66 DPS, sim-verified) [world_drop]; Tidal Charm (1404, -4.03 DPS) [vendor]; Guardian Talisman (1490, -4.03 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 46.1 | yes | Ankh of Life (1713, +0.12 DPS, sim-verified) [world_drop]; Tidal Charm (1404, -2.39 DPS) [vendor]; Guardian Talisman (1490, -2.39 DPS) [quest] |
| main_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | 570.2 | yes | Kindling Stave (11750, +0.00 DPS) [dungeon]; Bleakwood Hew (12769, +0.00 DPS) [crafted]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | 526.4 | yes | Claw of Celebras (17738, -2.29 DPS) [dungeon]; White Bone Shredder (11863, -4.10 DPS) [quest]; Thermotastic Egg Timer (9644, -27.07 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | 16.5 | yes | Moonsight Rifle (4383, -0.06 DPS) [crafted]; Precision Bow (217315, -0.06 DPS) [quest]; Houndmaster's Bow (11628, -0.08 DPS) [dungeon] |

**New at 50:** head: Blood Guard's Chain Helmet; neck: Woven Ivy Necklace; shoulder: Blood Guard's Chain Epaulets; chest: Stone Guard's Chain Armor; wrist: Bracers of the Stone Princess; waist: Defiler's Chain Girdle; legs: Knight's Chain Legplates; feet: Greaves of Withering Despair; finger1: Blackstone Ring; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Dawn's Edge; off_hand: Thorium Cestus; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1601, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60 (troll, 0000000000000000-32005500005000000-500230131051120151)

Set DPS (verified): 244.4. Weights run: 1.5s. Verify run: 1.7s. 2285 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.205 ± 0.024, strength=1.000 ± 0.001, crit=6.402 ± 0.250, hit=5.883 ± 0.716, melee_haste=not significant (0.393 ± 1.319)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Champion's Chain Greathelm (227080) (or Lieutenant Commander's Chain Greathelm (227086)) | Lady Palanseer [vendor] | 227.5 | yes | Lieutenant Commander's Chain Greathelm (227086, +0.00 DPS, sim-verified) [vendor]; Cryptstalker Headpiece (22438, -0.55 DPS) [quest]; Mask of the Unforgiven (13404, -1.02 DPS) [dungeon] |
| neck | Blazefury Medallion (17111) | Lord Kazzak [world] | 15.7 | yes | Beads of Ogre Might (22150, +0.00 DPS) [quest]; Medallion of the Dawn (22659, +0.00 DPS) [quest]; Onyxia Tooth Pendant (18404, -0.14 DPS, sim-verified) [quest] |
| shoulder | Champion's Chain Pauldrons (227078) (or Lieutenant Commander's Chain Pauldrons (227084)) | Rank 14 [pvp] | 129.5 | yes | Lieutenant Commander's Chain Pauldrons (227084, +0.00 DPS, sim-verified) [pvp]; Cryptstalker Spaulders (22439, -0.25 DPS) [quest]; Champion's Chain Shoulders (23252, -0.92 DPS) [vendor] |
| back | Chromatic Cloak (18509) | Leatherworking [crafted] | 89.6 | yes | Cloak of the Unseen Path (21403, +1.99 DPS, sim-verified) [quest]; Earthweave Cloak (21187, -0.64 DPS) [quest]; Cloak of the Fallen God (21710, -2.38 DPS) [quest] |
| chest | Legionnaire's Chain Armor (227083) (or Knight-Captain's Chain Armor (227089)) | Lady Palanseer [vendor] | 220.3 | yes | Knight-Captain's Chain Armor (227089, +0.00 DPS, sim-verified) [vendor]; Cryptstalker Tunic (22436, -0.89 DPS) [quest]; Legionnaire's Chain Hauberk (22874, -1.10 DPS) [vendor] |
| wrist | Cryptstalker Wristguards (22443) | Cryptstalker Wristguards [quest] | 90.1 | yes | Rockfury Bracers (21186, -1.58 DPS) [quest]; Primal Batskin Bracers (19687, -2.12 DPS, sim-verified) [crafted]; Windtalker's Wristguards (19582, -2.63 DPS) [rep] |
| hands | Marshal's Chain Grips (16463) | Captain Dirgehammer [vendor] | 114.9 | yes | Chromatic Gauntlets (19157, +0.00 DPS) [crafted]; Stormshroud Gloves (21278, +0.00 DPS) [crafted]; Marshal's Chain Grips (231560, +0.00 DPS) [pvp] |
| waist | Cryptstalker Girdle (22442) | Cryptstalker Girdle [quest] | 176.2 | yes | Defiler's Leather Girdle (20190, -2.65 DPS) [rep]; Light Obsidian Belt (22195, -2.75 DPS) [crafted]; Defiler's Chain Girdle (20150, -5.89 DPS, sim-verified) [rep] |
| legs | Legionnaire's Chain Legplates (227079) (or Knight-Captain's Chain Legplates (227085)) | Lady Palanseer [vendor] | 220.3 | yes | Knight-Captain's Chain Legplates (227085, +0.00 DPS, sim-verified) [vendor]; Legionnaire's Chain Legguards (22875, -1.10 DPS) [vendor]; Knight-Captain's Chain Legguards (23293, -1.10 DPS) [vendor] |
| feet | General's Chain Sabatons (231564) | Rank 16 [pvp] | 115.5 | yes | Marshal's Chain Boots (16462, -1.28 DPS) [vendor]; General's Chain Sabatons (16569, -1.28 DPS) [vendor]; Cryptstalker Boots (22440, -1.52 DPS, sim-verified) [quest] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 164.5 | yes | Band of the Penitent (13217, -3.77 DPS) [quest]; Dragonslayer's Signet (18403, -3.77 DPS) [quest]; Band of Earthen Might (21182, -6.07 DPS, sim-verified) [quest] |
| finger2 | Master Dragonslayer's Ring (19384) | The Lord of Blackrock [quest] | 106.8 | yes | Band of the Penitent (13217, -0.87 DPS) [quest]; Dragonslayer's Signet (18403, -0.87 DPS) [quest]; Band of Earthen Might (21182, -2.89 DPS, sim-verified) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 83.2 | yes | Frozen Heart of the Mountain (249469, -1.52 DPS) [crafted]; Tidal Charm (1404, -4.19 DPS) [vendor]; Guardian Talisman (1490, -4.19 DPS) [quest] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Frozen Heart of the Mountain (249469, +0.69 DPS, sim-verified) [crafted]; Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest] |
| main_hand | High Warlord's Quickblade (234553) | Rank 18 [pvp] | 984.8 | yes | Atiesh, Greatstaff of the Guardian (22632, +0.00 DPS) [quest]; High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Battle Axe (234543, +0.00 DPS) [pvp] |
| off_hand | Grand Marshal's Swiftblade (234579) | Rank 18 [pvp] | 984.8 | yes | High Warlord's Left Claw (234558, -0.16 DPS) [pvp]; Grand Marshal's Left Hand Blade (234584, -0.16 DPS) [pvp]; Grand Marshal's Left Hand Blade (18847, -1.74 DPS) [vendor] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | 89.6 | yes | Fahrad's Reloading Repeater (22347, -1.31 DPS) [quest]; Core Marksman Rifle (18282, -1.55 DPS) [crafted]; Blackcrow (12651, -1.67 DPS) [dungeon] |

**New at 60:** head: Champion's Chain Greathelm; neck: Blazefury Medallion; shoulder: Champion's Chain Pauldrons; back: Chromatic Cloak; chest: Legionnaire's Chain Armor; wrist: Cryptstalker Wristguards; hands: Marshal's Chain Grips; waist: Cryptstalker Girdle; legs: Legionnaire's Chain Legplates; feet: General's Chain Sabatons; finger1: Don Julio's Band; finger2: Master Dragonslayer's Ring; trinket2: Ankh of Life; main_hand: High Warlord's Quickblade; off_hand: Grand Marshal's Swiftblade; ranged: The Purifier

No-known-source sample (15 of 2285, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

