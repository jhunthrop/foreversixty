# Leveling BiS: Survival

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-00000000000000000-500230100000000000)

Set DPS (verified): 57.1. Weights run: 1.9s. Verify run: 2.0s. 179 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.041 ± 0.008, strength=1.000 ± 0.001, crit=1.602 ± 0.050, hit=0.386 ± 0.016, melee_haste=not significant (0.818 ± 0.763)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.3 | yes | Defender's Leather Hood (252447, -0.05 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -0.34 DPS) [crafted]; Shadow Goggles (4373, -0.34 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.2 | yes | Erudite's Amulet (277204, -0.13 DPS, sim-verified) [quest]; Scholarly Pendant (277203, -0.25 DPS) [quest]; Tarnished Locket (279870, -0.25 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.2 | yes | Reinforced Woolen Shoulders (4315, -0.21 DPS) [crafted]; Forest Leather Mantle (4709, -0.21 DPS) [world_drop]; Double-Stitched Woolen Shoulders (4314, -0.33 DPS, sim-verified) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.2 | yes | Cape of the Brotherhood (5193, -0.04 DPS) [dungeon]; Dark Leather Cloak (2316, -0.05 DPS) [crafted]; Catacomb Cloak (279899, -0.16 DPS, sim-verified) [quest] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 12.3 | yes | Defender's Leather Armor (252434, -0.09 DPS) [crafted]; Prospector's Chestpiece (14562, -0.16 DPS) [world_drop]; Tunic of Westfall (2041, -0.19 DPS, sim-verified) [quest] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 6.2 | yes | Bristlebark Bindings (14569, -0.04 DPS) [world_drop]; Wolf Bracers (4794, -0.08 DPS) [vendor]; Forest Leather Bracers (3202, -0.20 DPS, sim-verified) [world_drop] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 22.4 | yes | Gloves of the Fang (10413, +0.00 DPS, sim-verified) [dungeon]; Bristlebark Gloves (14572, -0.58 DPS) [world_drop]; Brawler's Leather Gloves (252494, -0.58 DPS) [crafted] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 | yes | Brawler's Leather Belt (252428, -0.40 DPS) [crafted]; Ruffian Belt (5975, -0.48 DPS) [world]; Deviate Scale Belt (6468, -0.55 DPS, sim-verified) [crafted] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 15.4 | yes | Trapper's Leather Pants (252501, +0.00 DPS, sim-verified) [crafted]; Leggings of the Fang (10410, -0.04 DPS) [dungeon]; Defender's Leather Pants (252445, -0.09 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 11.3 | yes | Brawler's Leather Boots (252439, -0.08 DPS, sim-verified) [crafted]; Blackened Defias Boots (10402, -0.21 DPS) [dungeon]; Footpads of the Fang (10411, -0.21 DPS) [dungeon] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.2 | yes | Demon Band (12054, -0.17 DPS) [world_drop]; Loop of Sacrifice (281673, -0.21 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.25 DPS) [dungeon] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.2 | yes | Loop of Sacrifice (281673, -0.13 DPS) [quest]; Demon Band (12054, -0.16 DPS, sim-verified) [world_drop]; Lavishly Jeweled Ring (1156, -0.17 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 | yes | Duskbringer (2205, +0.00 DPS) [dungeon]; Living Root (6631, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 233.6 | yes | Cruel Barb (5191, +0.00 DPS, sim-verified) [dungeon]; Grayson's Torch (1172, -9.43 DPS) [quest]; Pulsating Hydra Heart (5183, -9.43 DPS) [world] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.2 | yes | Fine Longbow (11304, -0.03 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.08 DPS) [crafted]; Light Bow (4576, -0.08 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Bravo's Armbands; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Signet of the Zhevra; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 179, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 30 (dwarf, 0000000000000000-00000000000000000-500230131051000000)

Set DPS (verified): 88.2. Weights run: 2.0s. Verify run: 2.1s. 314 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.050 ± 0.013, strength=1.000 ± 0.001, crit=2.706 ± 0.100, hit=0.586 ± 0.019, melee_haste=not significant (0.559 ± 0.829)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 12.0 | yes | Brawler's Leather Helm (252512, -0.05 DPS, sim-verified) [crafted]; Cloudy Gustwoven Hood (277042, -0.14 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.14 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Kaleidoscope Chain (13084, -0.26 DPS) [world_drop]; Sentinel's Medallion (19541, -0.30 DPS, sim-verified) [rep]; Sentinel's Medallion (20444, -0.35 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 16.6 | yes | Barbaric Shoulders (5964, -0.29 DPS) [crafted]; Bristlebark Amice (14573, -0.33 DPS) [world_drop]; Mantle of Thieves (2264, -0.34 DPS, sim-verified) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 10.4 | yes | Wolfmaster Cape (6314, -0.05 DPS, sim-verified) [dungeon]; Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.10 DPS) [pvp] |
| chest | Raptorbane Armor (3566) | Ormer's Revenge [quest] | 16.0 | yes | Dusky Leather Armor (7374, -0.02 DPS, sim-verified) [crafted]; Brawler's Leather Tunic (252508, -0.07 DPS) [crafted]; Brawler's Leather Armor (252490, -0.17 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 10.3 | yes | Cultist's Armguards (270032, -0.05 DPS, sim-verified) [quest]; Jurassic Wristguards (6198, -0.09 DPS) [world]; Barbaric Bracers (18948, -0.10 DPS) [crafted] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 37.9 | yes | Heavy Earthen Gloves (7359, +0.00 DPS, sim-verified) [crafted]; Insignia Gloves (6408, -1.11 DPS) [world_drop]; Toughened Leather Gloves (4253, -1.16 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 | yes | Highlander's Leather Girdle (20117, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.27 DPS) [dungeon]; Skulker's Leather Belt (252520, -0.39 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Brawler's Leather Pants (252500, -0.48 DPS) [crafted]; Trapper's Leather Pants (252501, -0.48 DPS) [crafted]; Brawler's Leather Legguards (252516, -2.02 DPS, sim-verified) [crafted] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 | yes | Feet of the Lynx (1121, -0.01 DPS, sim-verified) [world_drop]; Brawler's Leather Boots (252439, -0.08 DPS) [crafted]; Insignia Boots (4055, -0.16 DPS) [world_drop] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 13.5 | yes | Thunderbrow Ring (13097, -0.10 DPS) [world_drop]; Insurgent's Band (272067, -0.20 DPS) [vendor]; Monkey Ring (6748, -0.28 DPS) [quest] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 12.3 | yes | Thunderbrow Ring (13097, -0.09 DPS, sim-verified) [world_drop]; Insurgent's Band (272067, -0.15 DPS) [vendor]; Protector's Band (20439, -0.19 DPS) [rep] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (66.2 DPS) | yes | Minor Recombobulator (4381, +0.00 DPS) [crafted]; Gnomish Universal Remote (7506, +0.00 DPS) [crafted]; Talisman of Arathor (21119, -1.96 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 340.4 | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Swinetusk Shank (6691, +0.00 DPS, sim-verified) [dungeon]; Morbid Dawn (7689, +0.00 DPS) [dungeon] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | sim-verified (88.2 DPS) | yes | Shoni's Disarming Tool (9608, -4.24 DPS) [quest]; Satyr's Rod (15962, -14.38 DPS) [world_drop]; Swinetusk Shank (6691, -22.04 DPS, sim-verified) [dungeon] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 | yes | Double-barreled Shotgun (2098, -0.15 DPS, sim-verified) [world_drop]; Silver Star (3463, -0.17 DPS) [quest]; BKP "Sparrow" Smallbore (3042, -0.22 DPS) [world_drop] |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Raptorbane Armor; wrist: Hawkeye's Bracers; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Ironspine's Eye; finger2: Protector's Band; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Pronged Reaver; off_hand: Bloody Brass Knuckles; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 314, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14145 Cursed Felblade

### Band 40 (dwarf, 0000000000000000-00000000000000000-500230131051120151)

Set DPS (verified): 123.7. Weights run: 1.9s. Verify run: 2.0s. 518 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.166 ± 0.016, strength=1.000 ± 0.001, crit=3.251 ± 0.109, hit=0.675 ± 0.026, melee_haste=not significant (3.438 ± 1.346)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 58.5 | yes | White Bandit Mask (10008, +0.00 DPS, sim-verified) [crafted]; Barbaric Iron Helm (7915, -2.03 DPS) [crafted]; Hawkeye's Helm (14591, -2.17 DPS) [world_drop] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19540, +0.00 DPS, sim-verified) [rep]; Sentinel's Medallion (19541, -0.24 DPS) [rep]; Kaleidoscope Chain (13084, -0.28 DPS) [world_drop] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.8 | yes | Barbaric Iron Shoulders (7913, -0.61 DPS) [crafted]; Nightscape Shoulders (8192, -0.62 DPS) [crafted]; Forest Tracker Epaulets (2278, -0.68 DPS, sim-verified) [world_drop] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 13.0 | yes | Wolfmaster Cape (6314, -0.16 DPS) [dungeon]; Yeti Fur Cloak (2805, -0.16 DPS) [quest]; Hawkeye's Cloak (14593, -0.64 DPS, sim-verified) [world_drop] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 22.2 | yes | Wolffear Harness (13110, +0.00 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -0.24 DPS) [crafted]; Tough Scorpid Breastplate (8203, -0.24 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Hawkeye's Bracers (14590, -0.47 DPS) [world_drop]; Cultist's Armguards (270032, -0.52 DPS) [quest]; Ravager's Armguards (14770, -0.73 DPS, sim-verified) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 65.5 | yes | Fletcher's Gloves (7348, -1.04 DPS) [crafted]; Shadowskin Gloves (18238, -1.04 DPS) [crafted]; Dragonscale Gauntlets (8347, -1.16 DPS, sim-verified) [crafted] |
| waist | Highlander's Chain Girdle (20089) | The League of Arathor [rep] | 53.5 | yes | Highlander's Leather Girdle (20116, +0.00 DPS, sim-verified) [rep]; Highlander's Chain Girdle (20090, -1.53 DPS) [rep]; Highlander's Leather Girdle (20117, -1.53 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Triprunner Dungarees (9624, -0.10 DPS) [quest]; Scarlet Leggings (10330, -0.26 DPS) [dungeon]; Basilisk Hide Pants (1718, -0.89 DPS, sim-verified) [world_drop] |
| feet | Blackforge Greaves (6423) | World drop [world_drop] | 20.7 | yes | Skulker's Leather Shoes (252531, +0.00 DPS, sim-verified) [crafted]; Prowler's Leather Shoes (252465, -0.08 DPS) [crafted]; Imperial Leather Boots (6431, -0.15 DPS) [world_drop] |
| finger1 | Assault Band (13095) | World drop [world_drop] | 20.0 | yes | Ironspine's Eye (7686, -0.29 DPS) [dungeon]; Ring of the Underwood (2951, -0.33 DPS) [world_drop]; Falcon's Hook (7552, -0.34 DPS) [world_drop] |
| finger2 | Protector's Band (19515) | Silverwing Sentinels [rep] | 17.3 | yes | Ring of the Underwood (2951, -0.19 DPS) [world_drop]; Falcon's Hook (7552, -0.20 DPS) [world_drop]; Ironspine's Eye (7686, -0.26 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | 442.0 | yes | Fiery War Axe (870, +0.00 DPS) [world_drop]; Staff of Jordan (873, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 440.1 | yes | Shoni's Disarming Tool (9608, -11.19 DPS) [quest]; Curve-bladed Ripper (2815, -16.30 DPS, sim-verified) [world_drop]; Stonecloth Branch (15963, -22.70 DPS) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | 14.0 | yes | Master Hunter's Rifle (17687, -0.28 DPS) [quest]; Swiftwind (13038, -0.30 DPS) [world_drop]; Booty Bay Bruiser's Buckshot (274748, -0.48 DPS, sim-verified) [vendor] |

**New at 40:** head: Raging Berserker's Helm; shoulder: Sunburn Spaulders; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Chain Girdle; feet: Blackforge Greaves; finger1: Assault Band; finger2: Protector's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Vanquisher's Sword; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 518, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 5000 Coral Band; 5008 Quicksilver Ring

### Band 50 (dwarf, 0000000000000000-32005000000000000-500230131051120151)

Set DPS (verified): 154.3. Weights run: 1.9s. Verify run: 2.0s. 669 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.196 ± 0.022, strength=1.000 ± 0.001, crit=4.313 ± 0.156, hit=0.907 ± 0.036, melee_haste=not significant (2.108 ± 0.907)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Chain Helmet (220822) | Captain Dirgehammer [vendor] | 79.5 | yes | Knight-Lieutenant's Mail Helmet (223075, -0.42 DPS) [vendor]; Eye of Theradras (17715, -0.98 DPS) [dungeon]; Raging Berserker's Helm (7719, -2.37 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 20.6 | yes | Ghostshard Talisman (7731, -0.33 DPS) [dungeon]; Sentinel's Medallion (19540, -0.38 DPS) [rep]; Sentinel's Medallion (19539, -0.83 DPS, sim-verified) [rep] |
| shoulder | Knight-Lieutenant's Chain Epaulets (220825) | Captain Dirgehammer [vendor] | 75.9 | yes | Knight-Lieutenant's Mail Epaulets (223073, +0.00 DPS, sim-verified) [vendor]; Sunburn Spaulders (274751, -2.59 DPS) [vendor]; Skulker's Leather Shoulder (252535, -2.75 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 17.9 | yes | Pridelord Cape (14673, -0.21 DPS) [world_drop]; Sergeant Major's Cape (16336, -0.24 DPS) [pvp]; Blackflame Cape (13109, -0.80 DPS, sim-verified) [world_drop] |
| chest | Knight's Chain Armor (220828) | Captain Dirgehammer [vendor] | 78.3 | yes | Knight's Mail Armor (223078, -0.60 DPS, sim-verified) [vendor]; Warbear Harness (15064, -2.34 DPS) [crafted]; Blazewind Breastplate (11193, -2.44 DPS) [quest] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 | yes | Branded Leather Bracers (19508, -0.41 DPS) [dungeon]; Skulker's Leather Bracers (252540, -0.52 DPS) [crafted]; Deepfury Bracers (13120, -1.52 DPS, sim-verified) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 80.4 | yes | Sergeant Major's Mail Gauntlets (223076, -0.28 DPS, sim-verified) [vendor]; Dragonscale Gauntlets (8347, -0.66 DPS) [crafted]; Fletcher's Gloves (7348, -1.02 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20088) (or Highlander's Leather Girdle (20115)) | The League of Arathor [rep] | 80.4 | yes | Highlander's Leather Girdle (20115, +0.00 DPS, sim-verified) [rep]; Highlander's Chain Girdle (20089, -0.61 DPS) [rep]; Highlander's Cloth Girdle (20097, -1.02 DPS) [rep] |
| legs | Knight's Chain Legplates (220832) | Captain Dirgehammer [vendor] | sim-verified (123.8 DPS) | yes | Knight's Mail Legplates (223074, -0.24 DPS) [vendor]; Gryphon Rider's Leggings (9652, -2.26 DPS) [quest]; Stormshroud Pants (15057, -2.59 DPS, sim-verified) [crafted] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 26.6 | yes | Prowler's Leather Boots (252468, -0.02 DPS) [crafted]; Sandstalker Ankleguards (12470, -0.05 DPS, sim-verified) [dungeon]; Albino Crocscale Boots (17728, -0.13 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 29.1 | yes | Assault Band (13095, -0.46 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.63 DPS) [quest]; Insurgent's Band (272065, -0.72 DPS) [vendor] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 20.8 | yes | Assault Band (13095, -0.09 DPS, sim-verified) [world_drop]; Protector's Band (19515, -0.16 DPS) [rep]; Masons Fraternity Ring (9533, -0.21 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (121.2 DPS) | yes | Thunderbrew's Boot Flask (744, -0.42 DPS) [quest]; Tidal Charm (1404, -0.42 DPS) [vendor]; Ankh of Life (1713, -0.42 DPS, sim-verified) [world_drop] |
| main_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | sim-verified (121.2 DPS) | yes | Hanzo Sword (8190, +0.00 DPS, sim-verified) [world_drop]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | sim-verified (151.1 DPS) | yes | Claw of Celebras (17738, -2.26 DPS) [dungeon]; Shoni's Disarming Tool (9608, -15.41 DPS) [quest]; Inventor's Focal Sword (17719, -29.93 DPS, sim-verified) [dungeon] |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (121.2 DPS) | yes | Precisely Calibrated Boomstick (2100, -0.01 DPS) [world_drop]; Houndmaster's Bow (11628, -0.10 DPS) [dungeon]; Dark Iron Rifle (16004, -1.56 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Chain Helmet; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Chain Epaulets; back: Dark Phantom Cape; chest: Knight's Chain Armor; wrist: Bracers of the Stone Princess; waist: Highlander's Chain Girdle; legs: Knight's Chain Legplates; feet: Skulker's Leather Boots; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Guardian Talisman; trinket2: Frozen Heart of the Mountain; main_hand: Dawn's Edge; off_hand: Thorium Cestus; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 669, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60 (dwarf, 0000000000000000-32005500005000000-500230131051120151)

Set DPS (verified): 260.1. Weights run: 2.0s. Verify run: 2.2s. 1223 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.205 ± 0.024, strength=1.000 ± 0.001, crit=6.402 ± 0.250, hit=1.190 ± 0.057, melee_haste=not significant (0.393 ± 1.319)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dawnstalker Headpiece (239540) | Leonid Barthalomew the Revered [vendor] | 238.3 | yes | Dawnstalker Visor (239532, -0.68 DPS) [vendor]; Lieutenant Commander's Chain Helm (23306, -1.88 DPS) [vendor]; Lieutenant Commander's Chain Greathelm (227086, -2.50 DPS, sim-verified) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (242.8 DPS) | yes | Blazefury Medallion (17111, -0.16 DPS, sim-verified) [world]; Beads of Ogre Might (22150, -3.92 DPS) [quest]; Amulet of the Darkmoon (19491, -4.07 DPS) [quest] |
| shoulder | Dawnstalker Pauldrons (239534) | Leonid Barthalomew the Revered [vendor] | sim-verified (247.0 DPS) | yes | Lieutenant Commander's Chain Pauldrons (227084, -2.67 DPS) [pvp]; Darkspear Pauldrons (272105, -3.04 DPS) [vendor]; Dawnstalker Spaulders (239542, -4.80 DPS, sim-verified) [vendor] |
| back | Chromatic Cloak (18509) (or Fel Cape (279269)) | Leatherworking [crafted] | 89.6 | yes | Fel Cape (279269, +0.00 DPS, sim-verified) [crafted]; Cloak of the Honor Guard (20073, -2.50 DPS) [rep]; Howler's Furs (272414, -2.51 DPS) [vendor] |
| chest | Dawnstalker Breastplate (239529) | Leonid Barthalomew the Revered [vendor] | sim-verified (246.4 DPS) | yes | Knight-Captain's Chain Armor (227089, -3.43 DPS) [vendor]; Dawnstalker Tunic (239543, -4.15 DPS, sim-verified) [vendor]; Knight-Captain's Chain Hauberk (23292, -4.52 DPS) [vendor] |
| wrist | Dawnstalker Vambraces (239536) | Leonid Barthalomew the Revered [vendor] | 137.4 | yes | Dawnstalker Wristguards (239544, -1.61 DPS, sim-verified) [vendor]; Windtalker's Wristguards (19582, -5.01 DPS) [rep]; Windtalker's Wristguards (19583, -5.21 DPS) [rep] |
| hands | Dawnstalker Gauntlets (239531) | Leonid Barthalomew the Revered [vendor] | sim-verified (245.7 DPS) | yes | Marshal's Chain Grips (231560, -2.16 DPS) [vendor]; Chromatic Gauntlets (19157, -2.16 DPS) [crafted]; Dawnstalker Handguards (239539, -3.48 DPS, sim-verified) [vendor] |
| waist | Dawnstalker Belt (239535) | Leonid Barthalomew the Revered [vendor] | 151.0 | yes | Highlander's Chain Girdle (20043, -1.38 DPS) [rep]; Highlander's Leather Girdle (20045, -1.38 DPS) [rep]; Dawnstalker Girdle (239538, -2.27 DPS, sim-verified) [vendor] |
| legs | Dawnstalker Legguards (239541) | Leonid Barthalomew the Revered [vendor] | 235.9 | yes | Sentinel's Chain Leggings (237819, -0.04 DPS, sim-verified) [vendor]; Dawnstalker Leggings (239533, -0.62 DPS) [vendor]; Knight-Captain's Chain Legplates (227085, -0.78 DPS) [vendor] |
| feet | Dawnstalker Greaves (239530) | Leonid Barthalomew the Revered [vendor] | sim-verified (248.2 DPS) | yes | Marshal's Chain Sabatons (231561, -5.77 DPS) [vendor]; Dawnstalker Boots (239537, -5.99 DPS, sim-verified) [vendor]; Scalegut Treaders (275618, -6.82 DPS) [crafted] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (230.6 DPS) | yes | Band of the Penitent (13217, -1.41 DPS) [quest]; Ring of Entropy (18543, -1.41 DPS) [world]; Wrath of Cenarius (21190, -6.49 DPS, sim-verified) [quest] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (230.6 DPS) | yes | Band of the Penitent (13217, -0.90 DPS) [quest]; Ring of Entropy (18543, -0.90 DPS) [world]; Wrath of Cenarius (21190, -5.26 DPS, sim-verified) [quest] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (230.5 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Ankh of Life (1713, -6.54 DPS, sim-verified) [world_drop] |
| trinket2 | Darkmoon Card: Heroism (19287) | Darkmoon Warlords Deck [quest] | sim-verified (230.6 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS, sim-verified) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Sword of Zeal (6622) | World drop [world_drop] | sim-verified (242.8 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [vendor]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Annihilator (12798, -3.34 DPS, sim-verified) [crafted] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 827.2 | yes | Grand Marshal's Left Hand Blade (18847, +0.00 DPS) [vendor]; Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [vendor]; Blackguard (19168, -5.02 DPS, sim-verified) [crafted] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (242.8 DPS) | yes | Dark Iron Rifle (16004, -2.09 DPS, sim-verified) [crafted]; Skull Splitting Crossbow (13039, -3.66 DPS) [world_drop]; Precisely Calibrated Boomstick (2100, -3.67 DPS) [world_drop] |

**New at 60:** head: Dawnstalker Headpiece; neck: Medallion of the Dawn; shoulder: Dawnstalker Pauldrons; back: Chromatic Cloak; chest: Dawnstalker Breastplate; wrist: Dawnstalker Vambraces; hands: Dawnstalker Gauntlets; waist: Dawnstalker Belt; legs: Dawnstalker Legguards; feet: Dawnstalker Greaves; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Darkmoon Card: Maelstrom; trinket2: Darkmoon Card: Heroism; main_hand: Sword of Zeal; off_hand: Ravencrest's Legacy; ranged: The Purifier

No-known-source sample (15 of 1223, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

## Horde

### Band 20 (troll, 0000000000000000-00000000000000000-500230100000000000)

Set DPS (verified): 56.9. Weights run: 1.9s. Verify run: 1.9s. 182 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.041 ± 0.008, strength=1.000 ± 0.001, crit=1.602 ± 0.050, hit=0.386 ± 0.016, melee_haste=not significant (0.818 ± 0.763)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.3 | yes | Defender's Leather Hood (252447, -0.05 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -0.34 DPS) [crafted]; Shadow Goggles (4373, -0.34 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.2 | yes | Erudite's Amulet (277204, -0.13 DPS, sim-verified) [quest]; Scholarly Pendant (277203, -0.25 DPS) [quest]; Tarnished Locket (279870, -0.25 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.2 | yes | Reinforced Woolen Shoulders (4315, -0.21 DPS) [crafted]; Forest Leather Mantle (4709, -0.21 DPS) [world_drop]; Double-Stitched Woolen Shoulders (4314, -0.32 DPS, sim-verified) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.2 | yes | Catacomb Cloak (279899, +0.00 DPS, sim-verified) [quest]; Cape of the Brotherhood (5193, -0.04 DPS) [dungeon]; Dark Leather Cloak (2316, -0.05 DPS) [crafted] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 12.3 | yes | Defender's Leather Armor (252434, -0.13 DPS, sim-verified) [crafted]; Prospector's Chestpiece (14562, -0.16 DPS) [world_drop]; Murloc Scale Breastplate (5781, -0.17 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.2 | yes | Bristlebark Bindings (14569, -0.01 DPS, sim-verified) [world_drop]; Wolf Bracers (4794, -0.04 DPS) [vendor]; Ratchet Wristwraps (274742, -0.08 DPS) [vendor] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 22.4 | yes | Gloves of the Fang (10413, +0.00 DPS, sim-verified) [dungeon]; Bristlebark Gloves (14572, -0.58 DPS) [world_drop]; Brawler's Leather Gloves (252494, -0.58 DPS) [crafted] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 | yes | Brawler's Leather Belt (252428, -0.40 DPS) [crafted]; Ruffian Belt (5975, -0.48 DPS) [world]; Deviate Scale Belt (6468, -0.55 DPS, sim-verified) [crafted] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 15.4 | yes | Trapper's Leather Pants (252501, +0.00 DPS, sim-verified) [crafted]; Leggings of the Fang (10410, -0.04 DPS) [dungeon]; Defender's Leather Pants (252445, -0.09 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 11.3 | yes | Brawler's Leather Boots (252439, -0.07 DPS, sim-verified) [crafted]; Blackened Defias Boots (10402, -0.21 DPS) [dungeon]; Footpads of the Fang (10411, -0.21 DPS) [dungeon] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.2 | yes | Demon Band (12054, -0.17 DPS) [world_drop]; Bounty Hunter's Ring (5351, -0.20 DPS) [quest]; Loop of Sacrifice (281673, -0.21 DPS) [quest] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.2 | yes | Bounty Hunter's Ring (5351, -0.13 DPS) [quest]; Loop of Sacrifice (281673, -0.13 DPS) [quest]; Demon Band (12054, -0.15 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 | yes | Duskbringer (2205, +0.00 DPS) [dungeon]; Living Root (6631, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 233.6 | yes | Cruel Barb (5191, +0.00 DPS, sim-verified) [dungeon]; Tork Wrench (11855, -9.35 DPS) [quest]; Grayson's Torch (1172, -9.43 DPS) [quest] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.2 | yes | Fine Longbow (11304, -0.02 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.08 DPS) [crafted]; Light Bow (4576, -0.08 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Signet of the Zhevra; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 182, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18856 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20437 Outrider's Bow; 20441 Scout's Blade; 209611 Insignia of the Alliance

### Band 30 (troll, 0000000000000000-00000000000000000-500230131051000000)

Set DPS (verified): 87.9. Weights run: 2.0s. Verify run: 2.0s. 318 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.050 ± 0.013, strength=1.000 ± 0.001, crit=2.706 ± 0.100, hit=0.586 ± 0.019, melee_haste=not significant (0.559 ± 0.829)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 12.0 | yes | Brawler's Leather Helm (252512, -0.03 DPS, sim-verified) [crafted]; Cloudy Gustwoven Hood (277042, -0.14 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.14 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Kaleidoscope Chain (13084, -0.26 DPS) [world_drop]; Scout's Medallion (19537, -0.28 DPS, sim-verified) [rep]; Scout's Medallion (20442, -0.35 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 16.6 | yes | Mantle of Thieves (2264, -0.25 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.29 DPS) [crafted]; Bristlebark Amice (14573, -0.33 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 10.4 | yes | Wildhunter Cloak (16658, -0.02 DPS) [quest]; Wolfmaster Cape (6314, -0.07 DPS, sim-verified) [dungeon]; Tigerstrike Mantle (13108, -0.09 DPS) [world_drop] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 14.7 | yes | Brawler's Leather Tunic (252508, +0.00 DPS, sim-verified) [crafted]; Brawler's Leather Armor (252490, -0.11 DPS) [crafted]; Defender's Leather Tunic (252450, -0.11 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 10.3 | yes | Cultist's Armguards (270032, -0.06 DPS, sim-verified) [quest]; Jurassic Wristguards (6198, -0.09 DPS) [world]; Barbaric Bracers (18948, -0.10 DPS) [crafted] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 37.9 | yes | Heavy Earthen Gloves (7359, +0.00 DPS, sim-verified) [crafted]; Insignia Gloves (6408, -1.11 DPS) [world_drop]; Toughened Leather Gloves (4253, -1.16 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 | yes | Defiler's Leather Girdle (20191, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.27 DPS) [dungeon]; Deftkin Belt (16659, -0.35 DPS) [quest] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Brawler's Leather Pants (252500, -0.48 DPS) [crafted]; Trapper's Leather Pants (252501, -0.48 DPS) [crafted]; Brawler's Leather Legguards (252516, -1.64 DPS, sim-verified) [crafted] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 | yes | Feet of the Lynx (1121, +0.00 DPS, sim-verified) [world_drop]; Brawler's Leather Boots (252439, -0.08 DPS) [crafted]; Insignia Boots (4055, -0.16 DPS) [world_drop] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 13.5 | yes | Thunderbrow Ring (13097, -0.10 DPS) [world_drop]; Insurgent's Band (272067, -0.20 DPS) [vendor]; Band of the Fist (17694, -0.24 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 12.3 | yes | Thunderbrow Ring (13097, -0.09 DPS, sim-verified) [world_drop]; Insurgent's Band (272067, -0.15 DPS) [vendor]; Band of the Fist (17694, -0.19 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (65.9 DPS) | yes | Minor Recombobulator (4381, +0.00 DPS) [crafted]; Gnomish Universal Remote (7506, +0.00 DPS) [crafted]; Defiler's Talisman (21120, -2.29 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 340.4 | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Swinetusk Shank (6691, +0.00 DPS, sim-verified) [dungeon]; Morbid Dawn (7689, +0.00 DPS) [dungeon] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | sim-verified (87.9 DPS) | yes | Tork Wrench (11855, -14.34 DPS) [quest]; Satyr's Rod (15962, -14.38 DPS) [world_drop]; Swinetusk Shank (6691, -21.84 DPS, sim-verified) [dungeon] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 | yes | Double-barreled Shotgun (2098, -0.15 DPS, sim-verified) [world_drop]; Silver Star (3463, -0.17 DPS) [quest]; BKP "Sparrow" Smallbore (3042, -0.22 DPS) [world_drop] |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Ironspine's Eye; finger2: Legionnaire's Band; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Pronged Reaver; off_hand: Bloody Brass Knuckles; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 318, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 40 (troll, 0000000000000000-00000000000000000-500230131051120151)

Set DPS (verified): 122.5. Weights run: 1.9s. Verify run: 2.0s. 515 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.166 ± 0.016, strength=1.000 ± 0.001, crit=3.251 ± 0.109, hit=0.675 ± 0.026, melee_haste=not significant (3.438 ± 1.346)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 58.5 | yes | White Bandit Mask (10008, +0.00 DPS, sim-verified) [crafted]; Barbaric Iron Helm (7915, -2.03 DPS) [crafted]; Hawkeye's Helm (14591, -2.17 DPS) [world_drop] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19536, +0.00 DPS, sim-verified) [rep]; Ethereal Talisman (4430, -0.23 DPS) [quest]; Scout's Medallion (19537, -0.24 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.8 | yes | Barbaric Iron Shoulders (7913, -0.61 DPS) [crafted]; Nightscape Shoulders (8192, -0.62 DPS) [crafted]; Forest Tracker Epaulets (2278, -0.67 DPS, sim-verified) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 11.2 | yes | Wildhunter Cloak (16658, -0.06 DPS) [quest]; Imperial Cloak (6432, -0.10 DPS) [world_drop]; Wolfmaster Cape (6314, -0.24 DPS, sim-verified) [dungeon] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 22.2 | yes | Wolffear Harness (13110, +0.00 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -0.24 DPS) [crafted]; Tough Scorpid Breastplate (8203, -0.24 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Hawkeye's Bracers (14590, -0.47 DPS) [world_drop]; Cultist's Armguards (270032, -0.52 DPS) [quest]; Ravager's Armguards (14770, -0.70 DPS, sim-verified) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 65.5 | yes | Fletcher's Gloves (7348, -1.04 DPS) [crafted]; Shadowskin Gloves (18238, -1.04 DPS) [crafted]; Dragonscale Gauntlets (8347, -1.10 DPS, sim-verified) [crafted] |
| waist | Defiler's Chain Girdle (20153) | The Defilers [rep] | 53.5 | yes | Defiler's Leather Girdle (20192, +0.00 DPS, sim-verified) [rep]; Defiler's Chain Girdle (20152, -1.53 DPS) [rep]; Defiler's Leather Girdle (20191, -1.53 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Triprunner Dungarees (9624, -0.10 DPS) [quest]; Scarlet Leggings (10330, -0.26 DPS) [dungeon]; Basilisk Hide Pants (1718, -1.56 DPS, sim-verified) [world_drop] |
| feet | Blackforge Greaves (6423) | World drop [world_drop] | 20.7 | yes | Skulker's Leather Shoes (252531, +0.00 DPS, sim-verified) [crafted]; Prowler's Leather Shoes (252465, -0.08 DPS) [crafted]; Imperial Leather Boots (6431, -0.15 DPS) [world_drop] |
| finger1 | Assault Band (13095) | World drop [world_drop] | 20.0 | yes | Ironspine's Eye (7686, -0.29 DPS) [dungeon]; Ring of the Underwood (2951, -0.33 DPS) [world_drop]; Falcon's Hook (7552, -0.34 DPS) [world_drop] |
| finger2 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 17.3 | yes | Ring of the Underwood (2951, -0.19 DPS) [world_drop]; Falcon's Hook (7552, -0.20 DPS) [world_drop]; Ironspine's Eye (7686, -0.25 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 442.0 | yes | Fiery War Axe (870, +0.00 DPS) [world_drop]; Staff of Jordan (873, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 440.1 | yes | Curve-bladed Ripper (2815, -15.93 DPS, sim-verified) [world_drop]; Stonecloth Branch (15963, -22.70 DPS) [world_drop]; Tork Wrench (11855, -22.76 DPS) [quest] |
| ranged | The Silencer (13138) | World drop [world_drop] | 14.0 | yes | Master Hunter's Rifle (17687, -0.28 DPS) [quest]; Swiftwind (13038, -0.30 DPS) [world_drop]; Booty Bay Bruiser's Buckshot (274748, -0.48 DPS, sim-verified) [vendor] |

**New at 40:** head: Raging Berserker's Helm; shoulder: Sunburn Spaulders; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Chain Girdle; feet: Blackforge Greaves; finger1: Assault Band; finger2: Legionnaire's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Vanquisher's Sword; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 515, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 50 (troll, 0000000000000000-32005000000000000-500230131051120151)

Set DPS (verified): 158.8. Weights run: 1.9s. Verify run: 1.9s. 667 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.196 ± 0.022, strength=1.000 ± 0.001, crit=4.313 ± 0.156, hit=0.907 ± 0.036, melee_haste=not significant (2.108 ± 0.907)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Chain Helmet (220821) | Lady Palanseer [vendor] | 79.5 | yes | Blood Guard's Mail Helmet (220820, -0.42 DPS) [vendor]; Eye of Theradras (17715, -0.98 DPS) [dungeon]; Raging Berserker's Helm (7719, -1.52 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 20.6 | yes | Scout's Medallion (19535, -0.32 DPS) [rep]; Ghostshard Talisman (7731, -0.33 DPS) [dungeon]; Woven Ivy Necklace (19159, -0.42 DPS, sim-verified) [quest] |
| shoulder | Blood Guard's Chain Epaulets (220824) | Lady Palanseer [vendor] | 75.9 | yes | Blood Guard's Mail Epaulets (220823, +0.00 DPS, sim-verified) [vendor]; Sunburn Spaulders (274751, -2.59 DPS) [vendor]; Skulker's Leather Shoulder (252535, -2.75 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 17.9 | yes | Pridelord Cape (14673, -0.21 DPS) [world_drop]; Serpentskin Cloak (8259, -0.24 DPS) [world_drop]; Blackflame Cape (13109, -0.26 DPS, sim-verified) [world_drop] |
| chest | Stone Guard's Chain Armor (220827) | Lady Palanseer [vendor] | 78.3 | yes | Stone Guard's Mail Armor (220826, -0.34 DPS, sim-verified) [vendor]; Warbear Harness (15064, -2.34 DPS) [crafted]; Blazewind Breastplate (11193, -2.44 DPS) [quest] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 | yes | Branded Leather Bracers (19508, -0.41 DPS) [dungeon]; Skulker's Leather Bracers (252540, -0.52 DPS) [crafted]; Deepfury Bracers (13120, -1.41 DPS, sim-verified) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 80.4 | yes | First Sergeant's Mail Gauntlets (220831, +0.00 DPS, sim-verified) [vendor]; Dragonscale Gauntlets (8347, -0.66 DPS) [crafted]; Fletcher's Gloves (7348, -1.02 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20151) (or Defiler's Leather Girdle (20193)) | The Defilers [rep] | 80.4 | yes | Defiler's Leather Girdle (20193, +0.00 DPS, sim-verified) [rep]; Defiler's Chain Girdle (20153, -0.61 DPS) [rep]; Highlander's Mail Girdle (20118, -1.02 DPS) [vendor] |
| legs | Stone Guard's Chain Legplates (220833) | Lady Palanseer [vendor] | sim-verified (126.4 DPS) | yes | Stone Guard's Mail Legplates (220834, -0.24 DPS) [vendor]; Stormshroud Pants (15057, -2.11 DPS, sim-verified) [crafted]; Serpentskin Leggings (8262, -2.30 DPS) [world_drop] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 26.6 | yes | Sandstalker Ankleguards (12470, +0.00 DPS, sim-verified) [dungeon]; Prowler's Leather Boots (252468, -0.02 DPS) [crafted]; Albino Crocscale Boots (17728, -0.13 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 29.1 | yes | Legionnaire's Band (19511, -0.42 DPS) [rep]; Assault Band (13095, -0.46 DPS) [world_drop]; Legionnaire's Band (19512, -0.59 DPS) [rep] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 | yes | Legionnaire's Band (19511, -0.19 DPS, sim-verified) [rep]; Assault Band (13095, -0.20 DPS) [world_drop]; Legionnaire's Band (19512, -0.33 DPS) [rep] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (125.7 DPS) | yes | Frozen Heart of the Mountain (249469, -2.05 DPS) [crafted]; Tidal Charm (1404, -2.47 DPS) [vendor]; Ankh of Life (1713, -3.48 DPS, sim-verified) [world_drop] |
| trinket2 | - | - |  |  |  |
| main_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | sim-verified (124.6 DPS) | yes | Hanzo Sword (8190, +0.00 DPS, sim-verified) [world_drop]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | sim-verified (156.2 DPS) | yes | Claw of Celebras (17738, -2.26 DPS) [dungeon]; White Bone Shredder (11863, -4.04 DPS) [quest]; Inventor's Focal Sword (17719, -31.90 DPS, sim-verified) [dungeon] |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (124.6 DPS) | yes | Precisely Calibrated Boomstick (2100, -0.01 DPS) [world_drop]; Houndmaster's Bow (11628, -0.10 DPS) [dungeon]; Dark Iron Rifle (16004, -1.54 DPS, sim-verified) [crafted] |

**New at 50:** head: Blood Guard's Chain Helmet; neck: Skibi's Pendant; shoulder: Blood Guard's Chain Epaulets; back: Dark Phantom Cape; chest: Stone Guard's Chain Armor; wrist: Bracers of the Stone Princess; waist: Defiler's Chain Girdle; legs: Stone Guard's Chain Legplates; feet: Skulker's Leather Boots; finger1: Blackstone Ring; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Guardian Talisman; main_hand: Dawn's Edge; off_hand: Thorium Cestus; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 667, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (troll, 0000000000000000-32005500005000000-500230131051120151)

Set DPS (verified): 265.5. Weights run: 2.0s. Verify run: 2.2s. 1221 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.205 ± 0.024, strength=1.000 ± 0.001, crit=6.402 ± 0.250, hit=1.190 ± 0.057, melee_haste=not significant (0.393 ± 1.319)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dawnstalker Headpiece (239540) | Leonid Barthalomew the Revered [vendor] | 238.3 | yes | Dawnstalker Visor (239532, -0.68 DPS) [vendor]; Champion's Chain Helm (23251, -1.88 DPS) [vendor]; Champion's Chain Greathelm (227080, -2.46 DPS, sim-verified) [vendor] |
| neck | Blazefury Medallion (17111) | Lord Kazzak [world] | sim-verified (248.9 DPS) | yes | Amulet of the Darkmoon (19491, +0.00 DPS) [quest]; Beads of Ogre Might (22150, +0.00 DPS) [quest]; Medallion of the Dawn (22659, +0.00 DPS, sim-verified) [quest] |
| shoulder | Dawnstalker Pauldrons (239534) | Leonid Barthalomew the Revered [vendor] | sim-verified (252.9 DPS) | yes | Champion's Chain Pauldrons (227078, -2.67 DPS) [pvp]; Darkspear Pauldrons (272105, -3.04 DPS) [vendor]; Dawnstalker Spaulders (239542, -4.75 DPS, sim-verified) [vendor] |
| back | Chromatic Cloak (18509) (or Fel Cape (279269)) | Leatherworking [crafted] | 89.6 | yes | Fel Cape (279269, +0.00 DPS, sim-verified) [crafted]; Deathguard's Cloak (20068, -2.50 DPS) [rep]; Howler's Furs (272414, -2.51 DPS) [vendor] |
| chest | Dawnstalker Breastplate (239529) | Leonid Barthalomew the Revered [vendor] | sim-verified (252.2 DPS) | yes | Legionnaire's Chain Armor (227083, -3.43 DPS) [vendor]; Dawnstalker Tunic (239543, -4.05 DPS, sim-verified) [vendor]; Legionnaire's Chain Hauberk (22874, -4.52 DPS) [vendor] |
| wrist | Dawnstalker Vambraces (239536) | Leonid Barthalomew the Revered [vendor] | 137.4 | yes | Dawnstalker Wristguards (239544, -1.56 DPS, sim-verified) [vendor]; Windtalker's Wristguards (19582, -5.01 DPS) [rep]; Windtalker's Wristguards (19583, -5.21 DPS) [rep] |
| hands | Dawnstalker Gauntlets (239531) | Leonid Barthalomew the Revered [vendor] | sim-verified (251.4 DPS) | yes | General's Chain Grips (231569, -2.16 DPS) [vendor]; Chromatic Gauntlets (19157, -2.16 DPS) [crafted]; Dawnstalker Handguards (239539, -3.20 DPS, sim-verified) [vendor] |
| waist | Dawnstalker Belt (239535) | Leonid Barthalomew the Revered [vendor] | 151.0 | yes | Defiler's Chain Girdle (20150, -1.38 DPS) [rep]; Defiler's Leather Girdle (20190, -1.38 DPS) [rep]; Dawnstalker Girdle (239538, -2.26 DPS, sim-verified) [vendor] |
| legs | Dawnstalker Legguards (239541) | Leonid Barthalomew the Revered [vendor] | 235.9 | yes | Sentinel's Chain Leggings (237819, +0.00 DPS, sim-verified) [vendor]; Dawnstalker Leggings (239533, -0.62 DPS) [vendor]; Legionnaire's Chain Legplates (227079, -0.78 DPS) [vendor] |
| feet | Dawnstalker Greaves (239530) | Leonid Barthalomew the Revered [vendor] | sim-verified (254.2 DPS) | yes | General's Chain Sabatons (231564, -5.77 DPS) [vendor]; Dawnstalker Boots (239537, -5.99 DPS, sim-verified) [vendor]; Scalegut Treaders (275618, -6.82 DPS) [crafted] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (236.2 DPS) | yes | Band of the Penitent (13217, -1.41 DPS) [quest]; Ring of Entropy (18543, -1.41 DPS) [world]; Wrath of Cenarius (21190, -6.28 DPS, sim-verified) [quest] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (236.2 DPS) | yes | Band of the Penitent (13217, -0.90 DPS) [quest]; Ring of Entropy (18543, -0.90 DPS) [world]; Wrath of Cenarius (21190, -5.04 DPS, sim-verified) [quest] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (232.1 DPS) | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (236.2 DPS) | yes | Tidal Charm (1404, -2.54 DPS) [vendor]; Guardian Talisman (1490, -2.54 DPS) [quest]; Frozen Heart of the Mountain (249469, -4.43 DPS, sim-verified) [crafted] |
| main_hand | Sword of Zeal (6622) | World drop [world_drop] | sim-verified (247.7 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [vendor]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Annihilator (12798, -0.49 DPS, sim-verified) [crafted] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 827.2 | yes | High Warlord's Left Claw (18848, +0.00 DPS) [vendor]; High Warlord's Left Claw (234558, +0.00 DPS) [vendor]; Blackguard (19168, -5.00 DPS, sim-verified) [crafted] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (248.9 DPS) | yes | Dark Iron Rifle (16004, -2.16 DPS, sim-verified) [crafted]; Skull Splitting Crossbow (13039, -3.66 DPS) [world_drop]; Precisely Calibrated Boomstick (2100, -3.67 DPS) [world_drop] |

**New at 60:** head: Dawnstalker Headpiece; neck: Blazefury Medallion; shoulder: Dawnstalker Pauldrons; back: Chromatic Cloak; chest: Dawnstalker Breastplate; wrist: Dawnstalker Vambraces; hands: Dawnstalker Gauntlets; waist: Dawnstalker Belt; legs: Dawnstalker Legguards; feet: Dawnstalker Greaves; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Sword of Zeal; off_hand: Ravencrest's Legacy; ranged: The Purifier

No-known-source sample (15 of 1221, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

