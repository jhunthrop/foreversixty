# Leveling BiS: Retribution

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

## Alliance

### Band 20 (human, 00000000000000000-0000000000000000-55100000000000000)

Set DPS (verified): 34.2. Weights run: 1.0s. Verify run: 0.7s. 487 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.133 ± 0.018, crit=2.410 ± 0.076, hit=1.703 ± 0.139, melee_haste=-2.419 ± 0.186

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 | yes | Defender's Leather Hood (252447, -0.16 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.65 DPS) [crafted]; Brawler's Leather Hood (252504, -0.66 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 0.8 | yes | Tarnished Locket (279870, -0.07 DPS, sim-verified) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 | yes | Silvered Bronze Shoulders (3481, +0.02 DPS, sim-verified) [crafted]; Serpent's Shoulders (5404, -0.18 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.21 DPS) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 | yes | Grave Shroud (279865, -0.06 DPS, sim-verified) [quest]; Catacomb Cloak (279899, -0.07 DPS) [quest]; Dark Leather Cloak (2316, -0.12 DPS) [crafted] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 | yes | Veteran's Chain Shirt (250488, -0.18 DPS, sim-verified) [crafted]; Defender's Leather Armor (252434, -0.19 DPS) [crafted]; Totemic Leather Armor (252435, -0.21 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 | yes | Cryptwalker Bracers (280095, -0.08 DPS, sim-verified) [quest]; Bravo's Armbands (270015, -0.19 DPS) [quest]; Runed Copper Bracers (2854, -0.21 DPS) [crafted] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 33.7 | yes | Gold-flecked Gloves (5195, +0.21 DPS, sim-verified) [dungeon]; Polar Gauntlets (7606, -0.75 DPS) [quest]; Blackened Defias Gloves (10401, -0.75 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 | yes | Cobrahn's Grasp (6460, -0.11 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.21 DPS) [world]; Support Girdle (1215, -0.28 DPS) [world] |
| legs | Chausses of Westfall (6087) | The Defias Brotherhood [quest] | 22.0 | yes | Veteran's Chain Leggings (250493, -0.07 DPS, sim-verified) [crafted]; Defender's Leather Pants (252445, -0.12 DPS) [crafted]; Totemic Leather Pants (252446, -0.14 DPS) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.7 | yes | Veteran's Boots (250503, -0.01 DPS, sim-verified) [crafted]; Guard's Boots (250504, -0.02 DPS) [crafted]; Defender's Leather Boots (252441, -0.02 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.5 | yes | Signet of the Zhevra (285330, -0.27 DPS) [world]; Lavishly Jeweled Ring (1156, -0.29 DPS) [dungeon]; Minor Channeling Ring (1449, -0.30 DPS) [quest] |
| finger2 | The 1 Ring (8350) | Fishing [world] | 2.1 | yes | Signet of the Zhevra (285330, -0.02 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.06 DPS) [dungeon]; Minor Channeling Ring (1449, -0.07 DPS) [quest] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | Verigan's Fist (6953) | The Test of Righteousness [quest] | 358.7 | yes | Smite's Mighty Hammer (7230, -2.09 DPS) [dungeon]; Duskbringer (2205, -2.23 DPS) [dungeon]; Forsaken Greataxe (251533, -3.70 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: The 1 Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Verigan's Fist

No-known-source sample (15 of 487, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 7956 Bronze Warhammer; 8178 Training Sword; 9602 Brushwood Blade; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9753 Nomad Buckler

### Band 30 (human, 00000000000000000-0000000000000000-55223310000000000)

Set DPS (verified): 71.1. Weights run: 1.0s. Verify run: 0.9s. 935 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.118 ± 0.014, crit=3.235 ± 0.092, hit=1.979 ± 0.168, melee_haste=1.778 ± 0.112

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 | yes | Defender's Leather Helm (252455, -0.08 DPS) [crafted]; Veteran's Chain Helm (250498, -0.11 DPS, sim-verified) [crafted]; Crusader's Chain Helm (250502, -0.17 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (20444, -0.55 DPS) [rep]; Sentinel's Medallion (19541, -0.57 DPS, sim-verified) [rep]; Pendant of Myzrael (4614, -0.58 DPS) [dungeon] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 14.0 | yes | Mail Combat Spaulders (6404, +0.30 DPS, sim-verified) [world_drop]; Barbaric Iron Shoulders (7913, -0.05 DPS) [crafted]; Elite Shoulders (4835, -0.08 DPS) [vendor] |
| back | Sergeant Major's Cape (16315) | Rank 9 [pvp] | 8.5 | yes | Lambent Scale Cloak (4706, -0.02 DPS) [world_drop]; Grave Shroud (279865, -0.09 DPS) [quest]; Wolfmaster Cape (6314, -0.62 DPS, sim-verified) [dungeon] |
| chest | Shining Silver Breastplate (2870) | Blacksmithing [crafted] | 28.0 | yes | Barbaric Iron Breastplate (7914, -0.22 DPS, sim-verified) [crafted]; Hard Gold Cuirass (250533, -0.25 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.30 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 | yes | Bands of Serra'kis (6902, -0.22 DPS, sim-verified) [dungeon]; Patterned Bronze Bracers (2868, -0.25 DPS) [crafted]; Technician's Bracers (270042, -0.25 DPS) [quest] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 | yes | Bonefist Gauntlets (4465, -0.17 DPS) [world]; Mail Combat Gauntlets (4075, -0.23 DPS) [world_drop]; Fletcher's Gloves (7348, -0.65 DPS, sim-verified) [crafted] |
| waist | Highlander's Plate Girdle (20126) | The League of Arathor [rep] | 24.0 | yes | Highlander's Chain Girdle (20090, +0.00 DPS, sim-verified) [rep]; Highlander's Leather Girdle (20117, -0.00 DPS) [rep]; Highlander's Lamellar Girdle (20108, -0.08 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Golden Scale Leggings (3843, -0.17 DPS) [crafted]; Chausses of Westfall (6087, -0.17 DPS) [quest]; Veteran's Silvered Chain Leggings (250523, -1.92 DPS, sim-verified) [crafted] |
| feet | Hard Gold Boots (250534) | Blacksmithing [crafted] | 14.0 | yes | Glimmering Mail Greaves (4073, -0.08 DPS) [world_drop]; Brawler's Leather Boots (252439, -0.14 DPS) [crafted]; Trouncing Boots (4464, -0.70 DPS, sim-verified) [world] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 12.7 | yes | Ironspine's Eye (7686, -0.15 DPS) [dungeon]; Insurgent's Band (272067, -0.15 DPS) [vendor]; Protector's Band (20439, -0.18 DPS) [rep] |
| finger2 | Silverlaine's Family Seal (6321) | Shadowfang Keep: Baron Silverlaine [dungeon] | 10.0 | yes | Ironspine's Eye (7686, +0.11 DPS, sim-verified) [dungeon]; Insurgent's Band (272067, -0.04 DPS) [vendor]; Seal of Wrynn (2933, -0.15 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 438.6 | yes | Corpsemaker (6687, -0.14 DPS) [dungeon]; Darkspear Raider's Reaper (272082, -0.67 DPS) [vendor]; Morbid Dawn (7689, -12.34 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Sergeant Major's Cape; chest: Shining Silver Breastplate; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Highlander's Plate Girdle; legs: Ferine Leggings; feet: Hard Gold Boots; finger1: Protector's Band; finger2: Silverlaine's Family Seal; trinket1: Darkspear Voodoo Seal; trinket2: Talisman of Arathor; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 935, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 7956 Bronze Warhammer; 7957 Bronze Greatsword

### Band 40 (human, 00000000000000000-0000000000000000-55223331211000210)

Set DPS (verified): 72.4. Weights run: 1.0s. Verify run: 0.8s. 1349 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.228 ± 0.031, crit=4.287 ± 0.122, hit=2.725 ± 0.232, melee_haste=2.170 ± 0.198

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 86.0 | yes | Icemetal Barbute (10763, -0.42 DPS, sim-verified) [dungeon]; Hard Gold Coif (250537, -2.46 DPS) [crafted]; Tusken Helm (6686, -2.55 DPS) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19540, -0.36 DPS, sim-verified) [rep]; Sentinel's Medallion (19541, -0.52 DPS) [rep]; Sentinel's Medallion (20444, -0.54 DPS) [rep] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 | yes | Shining Mithril Pauldrons (250541, +0.52 DPS, sim-verified) [crafted]; Imperial Leather Spaulders (4737, -0.17 DPS) [world_drop]; Wrangling Spaulders (15698, -0.24 DPS) [quest] |
| back | Sergeant Major's Cape (16336) | Rank 9 [pvp] | 13.4 | yes | Sergeant Major's Cape (16315, -0.19 DPS) [pvp]; Lambent Scale Cloak (4706, -0.23 DPS) [world_drop]; Wolfmaster Cape (6314, -1.29 DPS, sim-verified) [dungeon] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 29.6 | yes | Golden Scale Cuirass (3845, -0.07 DPS) [crafted]; Shining Mithril Breastplate (250540, -0.07 DPS) [crafted]; Shining Silver Breastplate (2870, -0.21 DPS, sim-verified) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Pugilist Bracers (4438, -0.21 DPS, sim-verified) [dungeon]; Ravager's Armguards (14770, -0.22 DPS) [world]; Golden Scale Bracers (6040, -0.34 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 80.0 | yes | Fletcher's Gloves (7348, -0.85 DPS) [crafted]; Ornate Mithril Gloves (7927, -0.85 DPS) [crafted]; Dragonscale Gauntlets (8347, -0.87 DPS, sim-verified) [crafted] |
| waist | Highlander's Chain Girdle (20089) | The League of Arathor [rep] | 68.0 | yes | Highlander's Leather Girdle (20116, +0.48 DPS, sim-verified) [rep]; Highlander's Plate Girdle (20125, -1.61 DPS) [rep]; Highlander's Chain Girdle (20090, -1.87 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 | yes | Orcish War Leggings (7929, -0.43 DPS, sim-verified) [crafted]; Ferine Leggings (6690, -0.68 DPS) [dungeon]; Ornate Mithril Pants (7926, -0.76 DPS) [crafted] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 27.6 | yes | Prowler's Leather Shoes (252465, +0.27 DPS, sim-verified) [crafted]; Skirmisher's Mail Boots (252564, -0.24 DPS) [crafted]; Blackforge Greaves (6423, -0.31 DPS) [world_drop] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 17.8 | yes | Protector's Band (19517, -0.19 DPS) [rep]; Insurgent's Band (272066, -0.25 DPS) [vendor]; Ironspine's Eye (7686, -0.33 DPS) [dungeon] |
| finger2 | Suspicious Spare Part (274754) | Rettrick [vendor] | 14.0 | yes | Insurgent's Band (272066, -0.11 DPS, sim-verified) [vendor]; Ironspine's Eye (7686, -0.17 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.17 DPS) [dungeon] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Bonebiter (6830) | In the Name of the Light [quest] | 583.5 | yes | Darkspear Raider's Reaper (272081, -0.28 DPS) [vendor]; Primitive Fishing Pole (276203, -0.57 DPS) [vendor]; Fiery War Axe (870, -0.79 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Chain Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Protector's Band; finger2: Suspicious Spare Part; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Bonebiter

No-known-source sample (15 of 1349, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat

### Band 50 (human, 00000000000000000-5500000000000000-55223331211000210)

Set DPS (verified): 100.4. Weights run: 1.0s. Verify run: 0.9s. 1766 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.231 ± 0.026, crit=5.537 ± 0.167, hit=3.474 ± 0.300, melee_haste=2.628 ± 0.325

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Plate Helm (220803) (or Knight-Lieutenant's Plate Helm (220804)) | Lady Palanseer [vendor] | 138.3 | yes | Knight-Lieutenant's Plate Helm (220804, +0.00 DPS, sim-verified) [vendor]; Raging Berserker's Helm (7719, -1.49 DPS) [dungeon]; Ornate Mithril Helm (7937, -1.75 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19539, -0.48 DPS) [rep]; Sentinel's Medallion (19540, -0.49 DPS) [rep]; Talisman of the Naga Lord (5029, -0.60 DPS, sim-verified) [world] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) (or Blood Guard's Plate Pauldrons (220796)) | Captain Dirgehammer [vendor] | 99.5 | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS, sim-verified) [vendor]; Officer's Pauldrons (250576, -3.08 DPS) [crafted]; Earthslag Shoulders (11632, -3.16 DPS) [dungeon] |
| back | Bloodlust Cape (14801) | Maraudon: Princess Theradras [dungeon] | 18.0 | yes | Sergeant Major's Cape (16336, +0.49 DPS, sim-verified) [pvp]; Wolfmaster Cape (6314, -0.34 DPS) [dungeon]; Sergeant Major's Cape (16315, -0.39 DPS) [pvp] |
| chest | Knight's Plate Hauberk (220794) (or Stone Guard's Plate Armor (220801)) | Captain Dirgehammer [vendor] | 107.5 | yes | Stone Guard's Plate Armor (220801, +0.00 DPS, sim-verified) [vendor]; Ornate Mithril Breastplate (7935, -1.29 DPS) [crafted]; Warforged Chestplate (11195, -2.56 DPS) [quest] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 | yes | Branded Leather Bracers (19508, -0.34 DPS) [dungeon]; Prowler's Leather Bracers (252539, -0.36 DPS) [crafted]; Officer's Wristguards (250581, -2.11 DPS, sim-verified) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 97.5 | yes | Sergeant Major's Lamellar Gauntlets (220817, +0.44 DPS, sim-verified) [vendor]; Dragonscale Gauntlets (8347, -0.80 DPS) [crafted]; Fletcher's Gloves (7348, -0.86 DPS) [crafted] |
| waist | Highlander's Lamellar Girdle (20106) | The League of Arathor [rep] | 99.5 | yes | Highlander's Leather Girdle (20115, -0.09 DPS) [rep]; Highlander's Plate Girdle (20124, -0.09 DPS) [rep]; Highlander's Chain Girdle (20088, -2.50 DPS, sim-verified) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | 155.0 | yes | Knight's Plate Leggings (220797, +0.89 DPS, sim-verified) [vendor]; Stone Guard's Plate Leggings (220798, -2.20 DPS) [vendor]; Scarlet Leggings (10330, -4.86 DPS) [dungeon] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 34.7 | yes | Officer's Sabatons (250561, -0.29 DPS) [crafted]; Officer's Boots (250546, -0.31 DPS) [crafted]; Prowler's Leather Boots (252468, -2.75 DPS, sim-verified) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 54.7 | yes | Insurgent's Band (272065, -1.71 DPS) [vendor]; Suspicious Spare Part (274754, -1.75 DPS) [vendor]; Insurgent's Band (272066, -1.84 DPS) [vendor] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 22.1 | yes | Protector's Band (19515, -0.29 DPS, sim-verified) [rep]; Insurgent's Band (272065, -0.30 DPS) [vendor]; Suspicious Spare Part (274754, -0.35 DPS) [vendor] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 31.3 | yes | Thunderbrew's Boot Flask (744, -0.83 DPS, sim-verified) [quest]; Tidal Charm (1404, -1.34 DPS) [vendor]; Guardian Talisman (1490, -1.34 DPS) [quest] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | 705.0 | yes | Darkspear Raider's Reaper (272080, -0.42 DPS) [vendor]; Bleakwood Hew (12769, -1.39 DPS) [crafted]; Dark Iron Pulverizer (11608, -1.76 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Blood Guard's Plate Helm; shoulder: Knight-Lieutenant's Plate Pauldrons; back: Bloodlust Cape; chest: Knight's Plate Hauberk; wrist: Bracers of the Stone Princess; waist: Highlander's Lamellar Girdle; legs: Stormshroud Pants; feet: Greaves of Withering Despair; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Ankh of Life; trinket2: Frozen Heart of the Mountain; main_hand: Thorium Greatmace

No-known-source sample (15 of 1766, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 60 (human, 00000000000000000-5532500000000000-55223331211000210)

Set DPS (verified): 166.9. Weights run: 1.0s. Verify run: 0.9s. 2604 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.339 ± 0.051, crit=7.362 ± 0.238, hit=9.891 ± 0.551, melee_haste=3.638 ± 0.629

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 440.0 | yes | Mask of the Unforgiven (13404, -4.10 DPS, sim-verified) [dungeon]; Bloodvine Goggles (19999, -6.08 DPS) [crafted]; Blood Guard's Plate Helm (220803, -9.27 DPS) [vendor] |
| neck | Onyxia Tooth Pendant (18404) | Celebrating Good Times [quest] | 204.3 | yes | Medallion of the Dawn (22659, -3.38 DPS) [quest]; Beads of Ogre Might (22150, -3.56 DPS) [quest]; Choker of the Shifting Sands (21505, -7.10 DPS) [quest] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) (or Blood Guard's Plate Pauldrons (220796)) | Captain Dirgehammer [vendor] | 125.1 | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS, sim-verified) [vendor]; Lightforge Spaulders (16729, -0.71 DPS) [dungeon]; Stormshroud Shoulders (15058, -0.96 DPS) [crafted] |
| back | Earthweave Cloak (21187) | Volunteer's Battlegear [quest] | 104.0 | yes | Chromatic Cloak (18509, -0.02 DPS, sim-verified) [crafted]; Cloak of the Honor Guard (20073, -2.99 DPS) [rep]; Cloak of the Fallen God (21710, -3.20 DPS) [quest] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 214.5 | yes | Stormshroud Armor (15056, -0.37 DPS) [crafted]; Bloodvine Vest (19682, -0.73 DPS) [crafted]; Bloodsoul Breastplate (19690, -0.85 DPS, sim-verified) [crafted] |
| wrist | Vambraces of the Sadist (13400) | Stratholme: Timmy the Cruel [dungeon] | 117.1 | yes | Rockfury Bracers (21186, -0.79 DPS) [quest]; Primal Batskin Bracers (19687, -1.40 DPS, sim-verified) [crafted]; Deeprock Bracers (21184, -3.31 DPS) [quest] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | 202.0 | yes | Primal Batskin Gloves (19686, -0.68 DPS, sim-verified) [crafted]; Chromatic Gauntlets (19157, -2.40 DPS) [crafted]; Marshal's Lamellar Gloves (16471, -2.75 DPS) [vendor] |
| waist | Highlander's Plate Girdle (20041) | The League of Arathor [rep] | 137.1 | yes | Highlander's Chain Girdle (20043, +0.00 DPS, sim-verified) [rep]; Highlander's Leather Girdle (20045, -0.00 DPS) [rep]; Light Obsidian Belt (22195, -0.09 DPS) [crafted] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | 360.9 | yes | Sentinel's Plate Legguards (22672, -0.56 DPS, sim-verified) [rep]; Sentinel's Lamellar Legguards (22753, -5.11 DPS) [rep]; Knight-Captain's Lamellar Leggings (16435, -6.15 DPS) [pvp] |
| feet | Redemption Boots (22430) | Redemption Boots [quest] | 103.1 | yes | Bloodvine Boots (19684, -0.18 DPS) [crafted]; Greaves of Withering Despair (22240, -0.18 DPS) [dungeon]; Knight-Lieutenant's Lamellar Sabatons (227146, -3.24 DPS, sim-verified) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 218.0 | yes | Master Dragonslayer's Ring (19384, -3.11 DPS) [quest]; Blackstone Ring (17713, -4.33 DPS) [dungeon]; Signet Ring of the Bronze Dragonflight (21205, -4.85 DPS) [quest] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | 214.0 | yes | Master Dragonslayer's Ring (19384, +1.15 DPS, sim-verified) [quest]; Blackstone Ring (17713, -4.16 DPS) [dungeon]; Signet Ring of the Bronze Dragonflight (21205, -4.67 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 89.0 | yes | Thunderbrew's Boot Flask (744, -3.89 DPS) [quest]; Tidal Charm (1404, -3.89 DPS) [vendor]; Guardian Talisman (1490, -3.89 DPS) [quest] |
| trinket2 | Onyxia Blood Talisman (18406) | Celebrating Good Times [quest] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest] |
| main_hand | High Warlord's Greatsword (234542) | Rank 18 [pvp] | 1280.7 | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Lionheart Helm; neck: Onyxia Tooth Pendant; back: Earthweave Cloak; chest: Dawn Armor; wrist: Vambraces of the Sadist; hands: Stormshroud Gloves; waist: Highlander's Plate Girdle; legs: Titanic Leggings; feet: Redemption Boots; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Frozen Heart of the Mountain; trinket2: Onyxia Blood Talisman; main_hand: High Warlord's Greatsword

No-known-source sample (15 of 2604, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

## Horde

### Band 20 (undead, 00000000000000000-0000000000000000-55100000000000000)

Set DPS (verified): 31.6. Weights run: 1.0s. Verify run: 0.7s. 484 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.133 ± 0.018, crit=2.410 ± 0.076, hit=1.703 ± 0.139, melee_haste=-2.419 ± 0.186

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 | yes | Defender's Leather Hood (252447, -0.16 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.65 DPS) [crafted]; Brawler's Leather Hood (252504, -0.66 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 0.8 | yes | Tarnished Locket (279870, +0.03 DPS, sim-verified) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 | yes | Silvered Bronze Shoulders (3481, +0.14 DPS, sim-verified) [crafted]; Serpent's Shoulders (5404, -0.18 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.21 DPS) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 | yes | Grave Shroud (279865, -0.03 DPS, sim-verified) [quest]; Catacomb Cloak (279899, -0.07 DPS) [quest]; Subterranean Cape (14149, -0.07 DPS) [dungeon] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 | yes | Veteran's Chain Shirt (250488, -0.17 DPS, sim-verified) [crafted]; Defender's Leather Armor (252434, -0.19 DPS) [crafted]; Totemic Leather Armor (252435, -0.21 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 | yes | Cryptwalker Bracers (280095, -0.08 DPS, sim-verified) [quest]; Runed Copper Bracers (2854, -0.21 DPS) [crafted]; Burnished Bracers (3211, -0.21 DPS) [world_drop] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 33.7 | yes | Gold-flecked Gloves (5195, +0.22 DPS, sim-verified) [dungeon]; Blackened Defias Gloves (10401, -0.75 DPS) [dungeon]; Dagmire Gauntlets (6481, -0.81 DPS) [quest] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 | yes | Cobrahn's Grasp (6460, -0.11 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.21 DPS) [world]; Support Girdle (1215, -0.28 DPS) [world] |
| legs | Veteran's Chain Leggings (250493) | Blacksmithing [crafted] | 18.7 | yes | Defender's Leather Pants (252445, -0.01 DPS, sim-verified) [crafted]; Totemic Leather Pants (252446, -0.02 DPS) [crafted]; Deepgrave Trousers (279900, -0.15 DPS) [quest] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.7 | yes | Veteran's Boots (250503, -0.01 DPS, sim-verified) [crafted]; Guard's Boots (250504, -0.02 DPS) [crafted]; Defender's Leather Boots (252441, -0.02 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.5 | yes | Signet of the Zhevra (285330, -0.27 DPS) [world]; Bounty Hunter's Ring (5351, -0.28 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.29 DPS) [dungeon] |
| finger2 | The 1 Ring (8350) | Fishing [world] | 2.1 | yes | Signet of the Zhevra (285330, +0.01 DPS, sim-verified) [world]; Bounty Hunter's Ring (5351, -0.06 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.06 DPS) [dungeon] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | Hammerbone (270018) | Leaders of the Fang [quest] | 306.2 | yes | Smite's Mighty Hammer (7230, -0.27 DPS) [dungeon]; Duskbringer (2205, -0.41 DPS) [dungeon]; Forsaken Greataxe (251533, -1.28 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: The 1 Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Hammerbone

No-known-source sample (15 of 484, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2879 Antipodean Rod; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 7956 Bronze Warhammer; 8178 Training Sword; 9602 Brushwood Blade; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse

### Band 30 (undead, 00000000000000000-0000000000000000-55223310000000000)

Set DPS (verified): 68.8. Weights run: 1.0s. Verify run: 0.9s. 934 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.118 ± 0.014, crit=3.235 ± 0.092, hit=1.979 ± 0.168, melee_haste=1.778 ± 0.112

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 | yes | Defender's Leather Helm (252455, -0.08 DPS) [crafted]; Veteran's Chain Helm (250498, -0.10 DPS, sim-verified) [crafted]; Crusader's Chain Helm (250502, -0.17 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19537, -0.47 DPS, sim-verified) [rep]; Scout's Medallion (20442, -0.55 DPS) [rep]; Pendant of Myzrael (4614, -0.58 DPS) [dungeon] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 14.0 | yes | Mail Combat Spaulders (6404, +0.42 DPS, sim-verified) [world_drop]; Barbaric Iron Shoulders (7913, -0.05 DPS) [crafted]; Elite Shoulders (4835, -0.08 DPS) [vendor] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Sergeant Major's Cape (16315, -0.06 DPS) [pvp]; Lambent Scale Cloak (4706, -0.08 DPS) [world_drop] |
| chest | Shining Silver Breastplate (2870) | Blacksmithing [crafted] | 28.0 | yes | Barbaric Iron Breastplate (7914, -0.20 DPS, sim-verified) [crafted]; Hard Gold Cuirass (250533, -0.25 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.30 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 | yes | Bands of Serra'kis (6902, -0.20 DPS, sim-verified) [dungeon]; Grimtoll Wristguards (15459, -0.24 DPS) [quest]; Patterned Bronze Bracers (2868, -0.25 DPS) [crafted] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 45.3 | yes | Gauntlets of Ogre Strength (3341, +0.51 DPS, sim-verified) [world]; Warsong Gauntlets (16978, -1.05 DPS) [quest]; Bonefist Gauntlets (4465, -1.14 DPS) [world] |
| waist | Defiler's Plate Girdle (20207) | The Defilers [rep] | 24.0 | yes | Defiler's Chain Girdle (20152, +0.00 DPS, sim-verified) [rep]; Defiler's Leather Girdle (20191, -0.00 DPS) [rep]; Officer's Belt (250556, -0.14 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Golden Scale Leggings (3843, -0.17 DPS) [crafted]; Juggernaut Leggings (6671, -0.25 DPS) [quest]; Veteran's Silvered Chain Leggings (250523, -1.67 DPS, sim-verified) [crafted] |
| feet | Hard Gold Boots (250534) | Blacksmithing [crafted] | 14.0 | yes | Glimmering Mail Greaves (4073, -0.08 DPS) [world_drop]; Brawler's Leather Boots (252439, -0.14 DPS) [crafted]; Trouncing Boots (4464, -0.65 DPS, sim-verified) [world] |
| finger1 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 12.7 | yes | Ironspine's Eye (7686, -0.15 DPS) [dungeon]; Insurgent's Band (272067, -0.15 DPS) [vendor]; Band of the Fist (17694, -0.18 DPS) [quest] |
| finger2 | Silverlaine's Family Seal (6321) | Shadowfang Keep: Baron Silverlaine [dungeon] | 10.0 | yes | Ironspine's Eye (7686, +0.14 DPS, sim-verified) [dungeon]; Insurgent's Band (272067, -0.04 DPS) [vendor]; Band of the Fist (17694, -0.06 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 438.6 | yes | Corpsemaker (6687, -0.14 DPS) [dungeon]; Darkspear Raider's Reaper (272082, -0.67 DPS) [vendor]; Morbid Dawn (7689, -16.62 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wolfmaster Cape; chest: Shining Silver Breastplate; wrist: Pugilist Bracers; waist: Defiler's Plate Girdle; legs: Ferine Leggings; feet: Hard Gold Boots; finger1: Legionnaire's Band; finger2: Silverlaine's Family Seal; trinket1: Darkspear Voodoo Seal; trinket2: Defiler's Talisman; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 934, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 7956 Bronze Warhammer

### Band 40 (undead, 00000000000000000-0000000000000000-55223331211000210)

Set DPS (verified): 71.1. Weights run: 1.0s. Verify run: 0.8s. 1347 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.228 ± 0.031, crit=4.287 ± 0.122, hit=2.725 ± 0.232, melee_haste=2.170 ± 0.198

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 86.0 | yes | Icemetal Barbute (10763, -0.43 DPS, sim-verified) [dungeon]; Hard Gold Coif (250537, -2.46 DPS) [crafted]; Tusken Helm (6686, -2.55 DPS) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Ethereal Talisman (4430, -0.07 DPS, sim-verified) [quest]; Scout's Medallion (19536, -0.49 DPS) [rep]; Scout's Medallion (19537, -0.52 DPS) [rep] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 | yes | Shining Mithril Pauldrons (250541, +0.53 DPS, sim-verified) [crafted]; Imperial Leather Spaulders (4737, -0.17 DPS) [world_drop]; Wrangling Spaulders (15698, -0.24 DPS) [quest] |
| back | Sergeant Major's Cape (16336) | Rank 9 [pvp] | 13.4 | yes | Wildhunter Cloak (16658, -0.14 DPS) [quest]; Sergeant Major's Cape (16315, -0.19 DPS) [pvp]; Wolfmaster Cape (6314, -0.92 DPS, sim-verified) [dungeon] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 29.6 | yes | Golden Scale Cuirass (3845, -0.07 DPS) [crafted]; Shining Mithril Breastplate (250540, -0.07 DPS) [crafted]; Shining Silver Breastplate (2870, -0.27 DPS, sim-verified) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Pugilist Bracers (4438, -0.21 DPS, sim-verified) [dungeon]; Ravager's Armguards (14770, -0.22 DPS) [world]; Darkspear Armsplints (4132, -0.25 DPS) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 80.0 | yes | Fletcher's Gloves (7348, -0.85 DPS) [crafted]; Ornate Mithril Gloves (7927, -0.85 DPS) [crafted]; Dragonscale Gauntlets (8347, -0.85 DPS, sim-verified) [crafted] |
| waist | Defiler's Chain Girdle (20153) | The Defilers [rep] | 68.0 | yes | Defiler's Leather Girdle (20192, +0.41 DPS, sim-verified) [rep]; Defiler's Plate Girdle (20206, -1.61 DPS) [rep]; Tharg's Shoelace (9705, -1.78 DPS) [quest] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 | yes | Orcish War Leggings (7929, -0.43 DPS, sim-verified) [crafted]; Ferine Leggings (6690, -0.68 DPS) [dungeon]; Ornate Mithril Pants (7926, -0.76 DPS) [crafted] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 27.6 | yes | Prowler's Leather Shoes (252465, +0.56 DPS, sim-verified) [crafted]; Skirmisher's Mail Boots (252564, -0.24 DPS) [crafted]; Blackforge Greaves (6423, -0.31 DPS) [world_drop] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 17.8 | yes | Legionnaire's Band (19513, -0.19 DPS) [rep]; Insurgent's Band (272066, -0.25 DPS) [vendor]; Ironspine's Eye (7686, -0.33 DPS) [dungeon] |
| finger2 | Suspicious Spare Part (274754) | Rettrick [vendor] | 14.0 | yes | Insurgent's Band (272066, -0.11 DPS, sim-verified) [vendor]; Ironspine's Eye (7686, -0.17 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.17 DPS) [dungeon] |
| trinket1 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Darkspear Raider's Reaper (272081) | Creeg Bothunk [vendor] | 576.8 | yes | Primitive Fishing Pole (276203, -0.29 DPS) [vendor]; Fiery War Axe (870, -0.51 DPS) [world_drop]; Mograine's Might (7723, -0.67 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Chain Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Legionnaire's Band; finger2: Suspicious Spare Part; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Darkspear Raider's Reaper

No-known-source sample (15 of 1347, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 50 (undead, 00000000000000000-5500000000000000-55223331211000210)

Set DPS (verified): 100.1. Weights run: 1.0s. Verify run: 0.8s. 1764 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.231 ± 0.026, crit=5.537 ± 0.167, hit=3.474 ± 0.300, melee_haste=2.628 ± 0.325

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Plate Helm (220803) | Lady Palanseer [vendor] | 138.3 | yes | Knight-Lieutenant's Plate Helm (220804, +0.00 DPS, sim-verified) [vendor]; Raging Berserker's Helm (7719, -1.49 DPS) [dungeon]; Ornate Mithril Helm (7937, -1.75 DPS) [crafted] |
| neck | Woven Ivy Necklace (19159) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 14.1 | yes | Ethereal Talisman (4430, -0.14 DPS) [quest]; Ghostshard Talisman (7731, -0.37 DPS, sim-verified) [dungeon]; Talisman of the Naga Lord (5029, -0.43 DPS) [world] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) (or Blood Guard's Plate Pauldrons (220796)) | Captain Dirgehammer [vendor] | 99.5 | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS, sim-verified) [vendor]; Officer's Pauldrons (250576, -3.08 DPS) [crafted]; Earthslag Shoulders (11632, -3.16 DPS) [dungeon] |
| back | Bloodlust Cape (14801) | Maraudon: Princess Theradras [dungeon] | 18.0 | yes | Sergeant Major's Cape (16336, +0.81 DPS, sim-verified) [pvp]; Wolfmaster Cape (6314, -0.34 DPS) [dungeon]; Battlehard Cape (11858, -0.34 DPS) [quest] |
| chest | Knight's Plate Hauberk (220794) (or Stone Guard's Plate Armor (220801)) | Captain Dirgehammer [vendor] | 107.5 | yes | Stone Guard's Plate Armor (220801, +0.00 DPS, sim-verified) [vendor]; Ornate Mithril Breastplate (7935, -1.29 DPS) [crafted]; Warforged Chestplate (11195, -2.56 DPS) [quest] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 | yes | Branded Leather Bracers (19508, -0.34 DPS) [dungeon]; Prowler's Leather Bracers (252539, -0.36 DPS) [crafted]; Officer's Wristguards (250581, -0.67 DPS, sim-verified) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 97.5 | yes | Sergeant Major's Lamellar Gauntlets (220817, +0.39 DPS, sim-verified) [vendor]; Dragonscale Gauntlets (8347, -0.80 DPS) [crafted]; Fletcher's Gloves (7348, -0.86 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20151) (or Defiler's Leather Girdle (20193)) | The Defilers [rep] | 97.5 | yes | Defiler's Leather Girdle (20193, +0.00 DPS, sim-verified) [rep]; Defiler's Plate Girdle (20205, -0.00 DPS) [rep]; Defiler's Chain Girdle (20153, -0.52 DPS) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | 155.0 | yes | Knight's Plate Leggings (220797, +0.90 DPS, sim-verified) [vendor]; Stone Guard's Plate Leggings (220798, -2.20 DPS) [vendor]; Scarlet Leggings (10330, -4.86 DPS) [dungeon] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 34.7 | yes | Officer's Sabatons (250561, -0.29 DPS) [crafted]; Officer's Boots (250546, -0.31 DPS) [crafted]; Prowler's Leather Boots (252468, -1.69 DPS, sim-verified) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 54.7 | yes | Legionnaire's Band (19511, -1.40 DPS) [rep]; Band of Allegiance (18585, -1.58 DPS) [quest]; Legionnaire's Band (19512, -1.59 DPS) [rep] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 | yes | Legionnaire's Band (19511, +0.17 DPS, sim-verified) [rep]; Band of Allegiance (18585, -0.26 DPS) [quest]; Legionnaire's Band (19512, -0.26 DPS) [rep] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 66.3 | yes | Frozen Heart of the Mountain (249469, -1.51 DPS) [crafted]; Tidal Charm (1404, -2.85 DPS) [vendor]; Guardian Talisman (1490, -2.85 DPS) [quest] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Frozen Heart of the Mountain (249469, +0.61 DPS, sim-verified) [crafted]; Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | 705.0 | yes | Darkspear Raider's Reaper (272080, -0.42 DPS) [vendor]; Bleakwood Hew (12769, -1.39 DPS) [crafted]; Dark Iron Pulverizer (11608, -1.76 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Blood Guard's Plate Helm; neck: Woven Ivy Necklace; shoulder: Knight-Lieutenant's Plate Pauldrons; back: Bloodlust Cape; chest: Knight's Plate Hauberk; wrist: Bracers of the Stone Princess; waist: Defiler's Chain Girdle; legs: Stormshroud Pants; feet: Greaves of Withering Despair; finger1: Blackstone Ring; finger2: White Bone Band; trinket1: Rune of the Guard Captain; main_hand: Thorium Greatmace

No-known-source sample (15 of 1764, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 60 (undead, 00000000000000000-5532500000000000-55223331211000210)

Set DPS (verified): 159.2. Weights run: 1.0s. Verify run: 0.9s. 2608 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.339 ± 0.051, crit=7.362 ± 0.238, hit=9.891 ± 0.551, melee_haste=3.638 ± 0.629

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 440.0 | yes | Mask of the Unforgiven (13404, -3.89 DPS, sim-verified) [dungeon]; Bloodvine Goggles (19999, -6.08 DPS) [crafted]; Knight-Lieutenant's Plate Helm (220804, -9.27 DPS) [vendor] |
| neck | Onyxia Tooth Pendant (18404) | For All To See [quest] | 204.3 | yes | Medallion of the Dawn (22659, -3.38 DPS) [quest]; Beads of Ogre Might (22150, -3.56 DPS) [quest]; Choker of the Shifting Sands (21505, -7.10 DPS) [quest] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) (or Blood Guard's Plate Pauldrons (220796)) | Captain Dirgehammer [vendor] | 125.1 | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS, sim-verified) [vendor]; Lightforge Spaulders (16729, -0.71 DPS) [dungeon]; Stormshroud Shoulders (15058, -0.96 DPS) [crafted] |
| back | Earthweave Cloak (21187) | Volunteer's Battlegear [quest] | 104.0 | yes | Chromatic Cloak (18509, -0.06 DPS, sim-verified) [crafted]; Deathguard's Cloak (20068, -2.99 DPS) [rep]; Cloak of the Fallen God (21710, -3.20 DPS) [quest] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 214.5 | yes | Stormshroud Armor (15056, -0.37 DPS) [crafted]; Bloodvine Vest (19682, -0.73 DPS) [crafted]; Bloodsoul Breastplate (19690, -0.90 DPS, sim-verified) [crafted] |
| wrist | Vambraces of the Sadist (13400) | Stratholme: Timmy the Cruel [dungeon] | 117.1 | yes | Rockfury Bracers (21186, -0.79 DPS) [quest]; Primal Batskin Bracers (19687, -1.27 DPS, sim-verified) [crafted]; Deeprock Bracers (21184, -3.31 DPS) [quest] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | 202.0 | yes | Primal Batskin Gloves (19686, -0.59 DPS, sim-verified) [crafted]; Chromatic Gauntlets (19157, -2.40 DPS) [crafted]; Marshal's Lamellar Gloves (16471, -2.75 DPS) [vendor] |
| waist | Defiler's Plate Girdle (20204) | The Defilers [rep] | 137.1 | yes | Defiler's Chain Girdle (20150, +0.00 DPS, sim-verified) [rep]; Defiler's Leather Girdle (20190, -0.00 DPS) [rep]; Light Obsidian Belt (22195, -0.09 DPS) [crafted] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | 360.9 | yes | Outrider's Plate Legguards (22651, -0.56 DPS, sim-verified) [rep]; Knight-Captain's Lamellar Leggings (16435, -6.15 DPS) [pvp]; Outrider's Chain Leggings (22673, -6.43 DPS) [rep] |
| feet | Knight-Lieutenant's Lamellar Sabatons (227146) | Captain Dirgehammer [vendor] | 127.1 | yes | Bloodvine Boots (19684, +1.43 DPS, sim-verified) [crafted]; Greaves of Withering Despair (22240, -1.23 DPS) [dungeon]; Scalegut Treaders (275618, -3.78 DPS) [crafted] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 218.0 | yes | Master Dragonslayer's Ring (19384, -3.11 DPS) [quest]; Blackstone Ring (17713, -4.33 DPS) [dungeon]; Signet Ring of the Bronze Dragonflight (21205, -4.85 DPS) [quest] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | 214.0 | yes | Master Dragonslayer's Ring (19384, +1.05 DPS, sim-verified) [quest]; Blackstone Ring (17713, -4.16 DPS) [dungeon]; Signet Ring of the Bronze Dragonflight (21205, -4.67 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 111.2 | yes | Frozen Heart of the Mountain (249469, -0.97 DPS) [crafted]; Tidal Charm (1404, -4.86 DPS) [vendor]; Guardian Talisman (1490, -4.86 DPS) [quest] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -0.21 DPS, sim-verified) [crafted] |
| main_hand | High Warlord's Greatsword (234542) | Rank 18 [pvp] | 1280.7 | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Lionheart Helm; neck: Onyxia Tooth Pendant; back: Earthweave Cloak; chest: Dawn Armor; wrist: Vambraces of the Sadist; hands: Stormshroud Gloves; waist: Defiler's Plate Girdle; legs: Titanic Leggings; feet: Knight-Lieutenant's Lamellar Sabatons; finger1: Don Julio's Band; finger2: Band of Earthen Might; main_hand: High Warlord's Greatsword

No-known-source sample (15 of 2608, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

