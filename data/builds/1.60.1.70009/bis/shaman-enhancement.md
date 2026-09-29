# Leveling BiS: Enhancement

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

## Alliance

### Band 20 (dwarf, 0000000000000000-253100000000000000-0000000000000000)

Set DPS (verified): 42.7. Weights run: 1.0s. Verify run: 1.0s. 480 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=0.235 ± 0.025, crit=4.333 ± 0.103, hit=4.241 ± 0.303, melee_haste=2.440 ± 0.609

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 | yes | Defender's Leather Hood (252447, -0.17 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.61 DPS) [crafted]; Brawler's Leather Hood (252504, -0.63 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 1.4 | yes | Tarnished Locket (279870, -0.21 DPS, sim-verified) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 | yes | Silvered Bronze Shoulders (3481, +0.06 DPS, sim-verified) [crafted]; Serpent's Shoulders (5404, -0.17 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.21 DPS) [crafted] |
| back | Lambent Scale Cloak (4706) | Gnomeregan: Caverndeep Burrower [dungeon] | 8.0 | yes | Catacomb Cloak (279899, -0.07 DPS) [quest]; Dark Leather Cloak (2316, -0.11 DPS) [crafted]; Grave Shroud (279865, -0.26 DPS, sim-verified) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 | yes | Veteran's Chain Shirt (250488, -0.16 DPS, sim-verified) [crafted]; Defender's Leather Armor (252434, -0.18 DPS) [crafted]; Totemic Leather Armor (252435, -0.21 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 | yes | Cryptwalker Bracers (280095, -0.09 DPS, sim-verified) [quest]; Bravo's Armbands (270015, -0.17 DPS) [quest]; Runed Copper Bracers (2854, -0.21 DPS) [crafted] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 60.7 | yes | Gold-flecked Gloves (5195, +0.10 DPS, sim-verified) [dungeon]; Polar Gauntlets (7606, -1.68 DPS) [quest]; Blackened Defias Gloves (10401, -1.68 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 | yes | Cobrahn's Grasp (6460, -0.10 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.21 DPS) [world]; Support Girdle (1215, -0.28 DPS) [world] |
| legs | Chausses of Westfall (6087) | The Defias Brotherhood [quest] | 22.0 | yes | Veteran's Chain Leggings (250493, -0.04 DPS, sim-verified) [crafted]; Defender's Leather Pants (252445, -0.11 DPS) [crafted]; Totemic Leather Pants (252446, -0.14 DPS) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 11.2 | yes | Veteran's Boots (250503, -0.03 DPS, sim-verified) [crafted]; Guard's Boots (250504, -0.04 DPS) [crafted]; Defender's Leather Boots (252441, -0.04 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.9 | yes | Signet of the Zhevra (285330, -0.26 DPS) [world]; Lavishly Jeweled Ring (1156, -0.29 DPS) [dungeon]; Minor Channeling Ring (1449, -0.31 DPS) [quest] |
| finger2 | The 1 Ring (8350) | Fishing [world] | 2.2 | yes | Lavishly Jeweled Ring (1156, -0.06 DPS) [dungeon]; Minor Channeling Ring (1449, -0.08 DPS) [quest]; Signet of the Zhevra (285330, -0.08 DPS, sim-verified) [world] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 237.0 | yes | Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest]; Hammerbone (270018, +0.00 DPS) [quest] |
| off_hand | Diamond Hammer (2194) | Gnomeregan: Caverndeep Ambusher [dungeon] | 232.8 | yes | Bear Buckler (4821, -7.82 DPS) [vendor]; Veteran Shield (3651, -7.89 DPS) [dungeon]; Burnished Shield (3655, -7.89 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: The 1 Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Butcher's Cleaver; off_hand: Diamond Hammer

No-known-source sample (15 of 480, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 3738 Brewing Rod; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7956 Bronze Warhammer; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9753 Nomad Buckler; 9756 Nomad Trousers; 9757 Nomad Tunic; 9763 Cadet Leggings

### Band 30 (dwarf, 0000000000000000-253130030004000000-0000000000000000)

Set DPS (verified): 70.9. Weights run: 0.9s. Verify run: 1.2s. 918 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.432 ± 0.100, crit=7.348 ± 0.337, hit=5.948 ± 0.471, melee_haste=2.835 ± 0.417

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 | yes | Defender's Leather Helm (252455, -0.07 DPS) [crafted]; Veteran's Chain Helm (250498, -0.12 DPS, sim-verified) [crafted]; Crusader's Chain Helm (250502, -0.14 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19541, -0.06 DPS, sim-verified) [rep]; Sentinel's Medallion (20444, -0.40 DPS) [rep]; Pendant of Myzrael (4614, -0.49 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | Gnomeregan: Mechanized Sentry [dungeon] | 14.7 | yes | Golden Scale Shoulders (3841, -0.03 DPS) [crafted]; Mail Combat Spaulders (6404, -0.03 DPS) [dungeon]; Barbaric Iron Shoulders (7913, -0.08 DPS, sim-verified) [crafted] |
| back | Sergeant Major's Cape (16315) | Rank 9 [pvp] | 9.7 | yes | Lambent Scale Cloak (4706, -0.06 DPS) [dungeon]; Grave Shroud (279865, -0.10 DPS) [quest]; Wolfmaster Cape (6314, -1.00 DPS, sim-verified) [dungeon] |
| chest | Shining Silver Breastplate (2870) | Blacksmithing [crafted] | 28.0 | yes | Veteran's Silvered Chain Shirt (250518, -0.19 DPS) [crafted]; Hard Gold Cuirass (250533, -0.21 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.23 DPS, sim-verified) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 | yes | Patterned Bronze Bracers (2868, -0.21 DPS) [crafted]; Technician's Bracers (270042, -0.21 DPS) [quest]; Bands of Serra'kis (6902, -0.23 DPS, sim-verified) [dungeon] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 102.9 | yes | Gauntlets of Ogre Strength (3341, +0.08 DPS, sim-verified) [world]; Bonefist Gauntlets (4465, -3.00 DPS) [world]; Mail Combat Gauntlets (4075, -3.01 DPS) [dungeon] |
| waist | Highlander's Plate Girdle (20126) | The League of Arathor [rep] | 24.0 | yes | Highlander's Chain Girdle (20090, +0.00 DPS, sim-verified) [rep]; Highlander's Leather Girdle (20117, -0.00 DPS) [rep]; Officer's Belt (250556, -0.05 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Golden Scale Leggings (3843, -0.14 DPS) [crafted]; Chausses of Westfall (6087, -0.14 DPS) [quest]; Veteran's Silvered Chain Leggings (250523, -0.81 DPS, sim-verified) [crafted] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 17.0 | yes | Hard Gold Boots (250534, +0.01 DPS, sim-verified) [crafted]; Brawler's Leather Boots (252439, -0.17 DPS) [crafted]; Glimmering Mail Greaves (4073, -0.18 DPS) [dungeon] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 14.6 | yes | Silverlaine's Family Seal (6321, -0.16 DPS) [dungeon]; Protector's Band (20439, -0.17 DPS) [rep]; Insurgent's Band (272067, -0.20 DPS) [vendor] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 11.9 | yes | Insurgent's Band (272067, -0.10 DPS) [vendor]; Seal of Wrynn (2933, -0.16 DPS) [quest]; Silverlaine's Family Seal (6321, -0.48 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 345.4 | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Darkspear Raider's Reaper (272082, +0.00 DPS) [vendor] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 334.9 | yes | Shoni's Disarming Tool (9608, -3.89 DPS) [quest]; Commander's Crest (6320, -11.48 DPS) [dungeon]; Lambent Scale Shield (3656, -11.55 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Sergeant Major's Cape; chest: Shining Silver Breastplate; wrist: Pugilist Bracers; waist: Highlander's Plate Girdle; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Protector's Band; finger2: Ironspine's Eye; trinket1: Darkspear Voodoo Seal; trinket2: Talisman of Arathor; main_hand: Pronged Reaver; off_hand: Ironspine's Fist

No-known-source sample (15 of 918, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7956 Bronze Warhammer; 7958 Bronze Battle Axe; 9362 Brilliant Gold Ring

### Band 40 (dwarf, 0000000000000000-253130030005102051-0000000000000000)

Set DPS (verified): 93.3. Weights run: 0.9s. Verify run: 1.2s. 1255 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.386 ± 0.051, crit=7.710 ± 0.213, hit=5.187 ± 0.448, melee_haste=3.244 ± 0.122

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 133.9 | yes | Hard Gold Coif (250537, -0.71 DPS, sim-verified) [crafted]; White Bandit Mask (10008, -3.99 DPS) [crafted]; Tusken Helm (6686, -4.00 DPS) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19541, -0.40 DPS) [rep]; Sentinel's Medallion (20444, -0.43 DPS) [rep]; Sentinel's Medallion (19540, -0.58 DPS, sim-verified) [rep] |
| shoulder | Imperial Leather Spaulders (4737) | Gnomeregan: Dark Iron Agent [dungeon] | 18.0 | yes | Wrangling Spaulders (15698, -0.05 DPS) [quest]; Sunburn Spaulders (274751, -0.07 DPS) [vendor]; Hard Gold Pauldrons (250539, -0.95 DPS, sim-verified) [crafted] |
| back | Sergeant Major's Cape (16336) | Rank 9 [pvp] | 14.3 | yes | Sergeant Major's Cape (16315, -0.18 DPS) [pvp]; Yeti Fur Cloak (2805, -0.22 DPS) [quest]; Wolfmaster Cape (6314, -1.60 DPS, sim-verified) [dungeon] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 30.7 | yes | Golden Scale Cuirass (3845, -0.10 DPS) [crafted]; Mail Combat Armor (4074, -0.17 DPS) [dungeon]; Shining Silver Breastplate (2870, -0.53 DPS, sim-verified) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Ravager's Armguards (14770, -0.17 DPS) [world]; Pugilist Bracers (4438, -0.23 DPS, sim-verified) [dungeon]; Golden Scale Bracers (6040, -0.30 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 127.9 | yes | Fletcher's Gloves (7348, -0.74 DPS) [crafted]; Shadowskin Gloves (18238, -0.74 DPS) [crafted]; Dragonscale Gauntlets (8347, -1.15 DPS, sim-verified) [crafted] |
| waist | Highlander's Chain Girdle (20089) | The League of Arathor [rep] | 115.9 | yes | Highlander's Leather Girdle (20116, +0.14 DPS, sim-verified) [rep]; Scarlet Belt (10329, -3.40 DPS) [dungeon]; Highlander's Plate Girdle (20126, -3.40 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 | yes | Orcish War Leggings (7929, -0.46 DPS, sim-verified) [crafted]; Ferine Leggings (6690, -0.59 DPS) [dungeon]; Veteran's Silvered Chain Leggings (250523, -0.64 DPS) [crafted] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 24.7 | yes | Blackforge Greaves (6423, -0.11 DPS) [dungeon]; Skirmisher's Mail Boots (252564, -0.17 DPS, sim-verified) [crafted]; Ironheel Boots (4653, -0.17 DPS) [quest] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 19.1 | yes | Protector's Band (19517, -0.22 DPS, sim-verified) [rep]; Insurgent's Band (272066, -0.26 DPS) [vendor]; Ironspine's Eye (7686, -0.28 DPS) [dungeon] |
| finger2 | Suspicious Spare Part (274754) | Rettrick [vendor] | 14.0 | yes | Ironspine's Eye (7686, -0.09 DPS) [dungeon]; Insurgent's Band (272066, -0.11 DPS, sim-verified) [vendor]; Silverlaine's Family Seal (6321, -0.15 DPS) [dungeon] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [dungeon]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [dungeon]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Ardent Custodian (868) | Gnomeregan: Dark Iron Ambassador [dungeon] | 460.0 | yes | Bonebiter (6830, +0.00 DPS) [quest]; Illusionary Rod (7713, +0.00 DPS) [dungeon]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 439.1 | yes | Shoni's Disarming Tool (9608, -7.94 DPS) [quest]; Salbac Shield (4652, -15.59 DPS) [quest]; Combat Shield (4065, -15.74 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; shoulder: Imperial Leather Spaulders; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Chain Girdle; legs: Scarlet Leggings; feet: Prowler's Leather Shoes; finger1: Protector's Band; finger2: Suspicious Spare Part; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Ardent Custodian; off_hand: Jhordy's Misplaced Screwdriver

No-known-source sample (15 of 1255, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

### Band 50 (dwarf, 5500000000000000-253130030005102051-0000000000000000)

Set DPS (verified): 120.6. Weights run: 1.0s. Verify run: 1.2s. 1673 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.512 ± 0.061, crit=9.934 ± 0.268, hit=7.977 ± 0.662, melee_haste=3.814 ± 0.090

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 165.1 | yes | Eye of Theradras (17715, +0.67 DPS, sim-verified) [dungeon]; Hard Gold Coif (250537, -4.96 DPS) [crafted]; White Bandit Mask (10008, -4.97 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19540, -0.30 DPS) [rep]; Sentinel's Medallion (19539, -0.35 DPS, sim-verified) [rep]; Sentinel's Medallion (19541, -0.36 DPS) [rep] |
| shoulder | Prowler's Leather Shoulder (252534) | Leatherworking [crafted] | 26.6 | yes | Skulker's Leather Shoulder (252535, -0.11 DPS) [crafted]; Kaylari Shoulders (10745, -0.17 DPS) [quest]; Failed Flying Experiment (9647, -0.52 DPS, sim-verified) [quest] |
| back | Bloodlust Cape (14801) | Maraudon: Princess Theradras [dungeon] | 18.0 | yes | Sergeant Major's Cape (16336, -0.22 DPS, sim-verified) [pvp]; Pridelord Cape (14673, -0.27 DPS) [dungeon]; Sergeant Major's Cape (16315, -0.29 DPS) [pvp] |
| chest | Wildthorn Mail (12624) | Blacksmithing [crafted] | 79.8 | yes | Kolkar Marauder Chain (6773, -0.51 DPS, sim-verified) [quest]; Warbear Harness (15064, -1.76 DPS) [crafted]; Relentless Chain (17777, -1.87 DPS) [quest] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 | yes | Branded Leather Bracers (19508, -0.29 DPS) [dungeon]; Skulker's Leather Bracers (252540, -0.34 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.34 DPS, sim-verified) [crafted] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 159.1 | yes | Fletcher's Gloves (7348, -0.72 DPS) [crafted]; Shadowskin Gloves (18238, -0.72 DPS) [crafted]; Dragonscale Gauntlets (8347, -1.16 DPS, sim-verified) [crafted] |
| waist | Highlander's Chain Girdle (20088) (or Highlander's Leather Girdle (20115)) | The League of Arathor [rep] | 159.1 | yes | Highlander's Leather Girdle (20115, +0.00 DPS, sim-verified) [rep]; Highlander's Chain Girdle (20089, -0.43 DPS) [rep]; Highlander's Cloth Girdle (20097, -0.72 DPS) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | 278.1 | yes | Scarlet Leggings (10330, -0.63 DPS, sim-verified) [dungeon]; Gryphon Rider's Leggings (9652, -8.70 DPS) [quest]; Serpentskin Leggings (8262, -8.83 DPS) [dungeon] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 79.8 | yes | Prowler's Leather Boots (252468, -1.71 DPS, sim-verified) [crafted]; Skulker's Leather Boots (252469, -1.85 DPS) [crafted]; Skirmisher's Mail Sabatons (252578, -1.95 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 99.8 | yes | Insurgent's Band (272065, -3.07 DPS) [vendor]; Suspicious Spare Part (274754, -3.10 DPS) [vendor]; Ironspine's Eye (7686, -3.15 DPS) [dungeon] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 24.6 | yes | Protector's Band (19517, -0.35 DPS) [rep]; Insurgent's Band (272065, -0.35 DPS) [vendor]; Protector's Band (19515, -0.51 DPS, sim-verified) [rep] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, -4.37 DPS, sim-verified) [dungeon] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, -0.24 DPS, sim-verified) [dungeon] |
| main_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | 649.1 | yes | Kindling Stave (11750, +0.00 DPS) [dungeon]; Bleakwood Hew (12769, +0.00 DPS) [crafted]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor] |
| off_hand | Hammer of the Northern Wind (810) | Blackrock Depths: Anvilrage Medic [dungeon] | 553.3 | yes | Claw of Celebras (17738, -2.57 DPS) [dungeon]; Shoni's Disarming Tool (9608, -11.89 DPS) [quest]; Shizzle's Drizzle Blocker (11915, -19.08 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 50:** shoulder: Prowler's Leather Shoulder; back: Bloodlust Cape; chest: Wildthorn Mail; wrist: Bracers of the Stone Princess; waist: Highlander's Chain Girdle; legs: Stormshroud Pants; feet: Greaves of Withering Despair; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Darkspear Voodoo Seal; trinket2: Thunderbrew's Boot Flask; main_hand: Dawn's Edge; off_hand: Hammer of the Northern Wind

No-known-source sample (15 of 1673, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60 (dwarf, 5533220000000000-253130030005102051-0000000000000000)

Set DPS (verified): 193.5. Weights run: 1.1s. Verify run: 1.3s. 2554 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.519 ± 0.068, crit=10.313 ± 0.324, hit=11.159 ± 1.127, melee_haste=3.972 ± 0.142

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bloodvine Goggles (19999) | Engineering [crafted] | 367.6 | yes | Ragefury Eyepatch (11735, -2.41 DPS) [dungeon]; Bloodvine Lens (19998, -2.84 DPS) [crafted]; Mask of the Unforgiven (13404, -3.07 DPS, sim-verified) [dungeon] |
| neck | Onyxia Tooth Pendant (18404) | For All To See [quest] | 259.6 | yes | Medallion of the Dawn (22659, -3.29 DPS) [quest]; Beads of Ogre Might (22150, -4.47 DPS) [quest]; Choker of the Shifting Sands (21505, -7.85 DPS) [quest] |
| shoulder | Champion's Mail Pauldrons (227154) | Rank 14 [pvp] | 282.0 | yes | Champion's Mail Pauldrons (23260, -0.60 DPS, sim-verified) [vendor]; Stormshroud Shoulders (15058, -4.96 DPS) [crafted]; Shroud of the Nathrezim (18720, -4.96 DPS) [dungeon] |
| back | Chromatic Cloak (18509) | Leatherworking [crafted] | 144.4 | yes | Earthweave Cloak (21187, -0.30 DPS, sim-verified) [quest]; Cloak of the Honor Guard (20073, -3.89 DPS) [rep]; Cloak of the Fallen God (21710, -3.93 DPS) [quest] |
| chest | Legionnaire's Mail Hauberk (227157) | Rank 12 [pvp] | 300.0 | yes | Stormshroud Armor (15056, -0.40 DPS) [crafted]; Savage Gladiator Chain (11726, -1.59 DPS) [dungeon]; Bloodsoul Breastplate (19690, -2.97 DPS, sim-verified) [crafted] |
| wrist | Primal Batskin Bracers (19687) | Leatherworking [crafted] | 118.9 | yes | Rockfury Bracers (21186, -2.08 DPS, sim-verified) [quest]; Windtalker's Wristguards (19582, -2.92 DPS) [rep]; Windtalker's Wristguards (19583, -3.06 DPS) [rep] |
| hands | Blood Guard's Mail Vices (227159) | Rank 11 [pvp] | 276.0 | yes | Primal Batskin Gloves (19686, -1.72 DPS) [crafted]; Stormshroud Gloves (21278, -1.81 DPS, sim-verified) [crafted]; Chromatic Gauntlets (19157, -3.16 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20043) (or Highlander's Leather Girdle (20045)) | The League of Arathor [rep] | 178.4 | yes | Highlander's Leather Girdle (20045, +0.00 DPS, sim-verified) [rep]; Light Obsidian Belt (22195, -0.07 DPS) [crafted]; Highlander's Chain Girdle (20088, -0.50 DPS) [rep] |
| legs | Legionnaire's Mail Legguards (227156) | Rank 12 [pvp] | 300.0 | yes | Sentinel's Chain Leggings (22748, -0.93 DPS) [rep]; Devilsaur Leggings (15062, -3.95 DPS) [crafted]; Stormshroud Pants (15057, -4.00 DPS, sim-verified) [crafted] |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 111.6 | yes | Greaves of Withering Despair (22240, +0.00 DPS) [dungeon]; Blood Guard's Mail Greaves (227158, -2.22 DPS, sim-verified) [pvp]; Shadowcraft Boots (16711, -2.42 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 272.0 | yes | Master Dragonslayer's Ring (19384, -4.05 DPS) [quest]; Band of the Penitent (13217, -4.60 DPS) [quest]; Dragonslayer's Signet (18403, -4.60 DPS) [quest] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | 268.0 | yes | Master Dragonslayer's Ring (19384, -0.65 DPS, sim-verified) [quest]; Band of the Penitent (13217, -4.46 DPS) [quest]; Dragonslayer's Signet (18403, -4.46 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest] |
| trinket2 | Onyxia Blood Talisman (18406) | For All To See [quest] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest] |
| main_hand | High Warlord's Cleaver (234554) | Rank 18 [pvp] | 1036.5 | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp] |
| off_hand | High Warlord's Bludgeon (234555) | Rank 18 [pvp] | 1036.5 | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Grand Marshal's Left Hand Blade (18847, -1.13 DPS) [vendor] |
| ranged | - | - |  |  |  |

**New at 60:** head: Bloodvine Goggles; neck: Onyxia Tooth Pendant; shoulder: Champion's Mail Pauldrons; back: Chromatic Cloak; chest: Legionnaire's Mail Hauberk; wrist: Primal Batskin Bracers; hands: Blood Guard's Mail Vices; waist: Highlander's Chain Girdle; legs: Legionnaire's Mail Legguards; feet: Bloodvine Boots; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Ankh of Life; trinket2: Onyxia Blood Talisman; main_hand: High Warlord's Cleaver; off_hand: High Warlord's Bludgeon

No-known-source sample (15 of 2554, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6478 Rat Stompers

## Horde

### Band 20 (orc, 0000000000000000-253100000000000000-0000000000000000)

Set DPS (verified): 45.3. Weights run: 1.0s. Verify run: 1.0s. 475 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=0.235 ± 0.025, crit=4.333 ± 0.103, hit=4.241 ± 0.303, melee_haste=2.440 ± 0.609

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 | yes | Defender's Leather Hood (252447, -0.19 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.61 DPS) [crafted]; Brawler's Leather Hood (252504, -0.63 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 1.4 | yes | Tarnished Locket (279870, -0.09 DPS, sim-verified) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 | yes | Silvered Bronze Shoulders (3481, -0.02 DPS, sim-verified) [crafted]; Serpent's Shoulders (5404, -0.17 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.21 DPS) [crafted] |
| back | Lambent Scale Cloak (4706) | Gnomeregan: Caverndeep Burrower [dungeon] | 8.0 | yes | Catacomb Cloak (279899, -0.07 DPS) [quest]; Subterranean Cape (14149, -0.07 DPS) [dungeon]; Grave Shroud (279865, -0.25 DPS, sim-verified) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 | yes | Veteran's Chain Shirt (250488, -0.15 DPS, sim-verified) [crafted]; Defender's Leather Armor (252434, -0.18 DPS) [crafted]; Totemic Leather Armor (252435, -0.21 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 | yes | Cryptwalker Bracers (280095, -0.09 DPS, sim-verified) [quest]; Bravo's Armbands (270015, -0.17 DPS) [quest]; Runed Copper Bracers (2854, -0.21 DPS) [crafted] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 60.7 | yes | Gold-flecked Gloves (5195, +0.10 DPS, sim-verified) [dungeon]; Blackened Defias Gloves (10401, -1.68 DPS) [dungeon]; Dagmire Gauntlets (6481, -1.72 DPS) [quest] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 | yes | Cobrahn's Grasp (6460, -0.05 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.21 DPS) [world]; Support Girdle (1215, -0.28 DPS) [world] |
| legs | Veteran's Chain Leggings (250493) | Blacksmithing [crafted] | 19.2 | yes | Defender's Leather Pants (252445, -0.02 DPS, sim-verified) [crafted]; Totemic Leather Pants (252446, -0.04 DPS) [crafted]; Deepgrave Trousers (279900, -0.15 DPS) [quest] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 11.2 | yes | Veteran's Boots (250503, -0.02 DPS, sim-verified) [crafted]; Guard's Boots (250504, -0.04 DPS) [crafted]; Defender's Leather Boots (252441, -0.04 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.9 | yes | Signet of the Zhevra (285330, -0.26 DPS) [world]; Bounty Hunter's Ring (5351, -0.28 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.29 DPS) [dungeon] |
| finger2 | The 1 Ring (8350) | Fishing [world] | 2.2 | yes | Bounty Hunter's Ring (5351, -0.05 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.06 DPS) [dungeon]; Signet of the Zhevra (285330, -0.17 DPS, sim-verified) [world] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 237.0 | yes | Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest]; Hammerbone (270018, +0.00 DPS) [quest] |
| off_hand | Diamond Hammer (2194) | Gnomeregan: Caverndeep Ambusher [dungeon] | 232.8 | yes | Bear Buckler (4821, -7.82 DPS) [vendor]; Faerleia's Shield (3450, -7.89 DPS) [quest]; Veteran Shield (3651, -7.89 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: The 1 Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Butcher's Cleaver; off_hand: Diamond Hammer

No-known-source sample (15 of 475, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 3738 Brewing Rod; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7956 Bronze Warhammer; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9753 Nomad Buckler; 9756 Nomad Trousers; 9757 Nomad Tunic; 9763 Cadet Leggings

### Band 30 (orc, 0000000000000000-253130030004000000-0000000000000000)

Set DPS (verified): 75.1. Weights run: 0.9s. Verify run: 1.1s. 913 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.432 ± 0.100, crit=7.348 ± 0.337, hit=5.948 ± 0.471, melee_haste=2.835 ± 0.417

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 | yes | Defender's Leather Helm (252455, -0.07 DPS) [crafted]; Veteran's Chain Helm (250498, -0.13 DPS, sim-verified) [crafted]; Crusader's Chain Helm (250502, -0.14 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (20442, -0.40 DPS) [rep]; Scout's Medallion (19537, -0.44 DPS, sim-verified) [rep]; Pendant of Myzrael (4614, -0.49 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | Gnomeregan: Mechanized Sentry [dungeon] | 14.7 | yes | Barbaric Iron Shoulders (7913, +0.14 DPS, sim-verified) [crafted]; Golden Scale Shoulders (3841, -0.03 DPS) [crafted]; Mail Combat Spaulders (6404, -0.03 DPS) [dungeon] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Sergeant Major's Cape (16315, -0.01 DPS) [pvp]; Lambent Scale Cloak (4706, -0.07 DPS) [dungeon] |
| chest | Shining Silver Breastplate (2870) | Blacksmithing [crafted] | 28.0 | yes | Veteran's Silvered Chain Shirt (250518, -0.19 DPS) [crafted]; Hard Gold Cuirass (250533, -0.21 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.26 DPS, sim-verified) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 | yes | Grimtoll Wristguards (15459, -0.20 DPS) [quest]; Patterned Bronze Bracers (2868, -0.21 DPS) [crafted]; Bands of Serra'kis (6902, -0.26 DPS, sim-verified) [dungeon] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 | yes | Warsong Gauntlets (16978, -0.07 DPS) [quest]; Bonefist Gauntlets (4465, -0.14 DPS) [world]; Fletcher's Gloves (7348, -0.77 DPS, sim-verified) [crafted] |
| waist | Defiler's Plate Girdle (20207) | The Defilers [rep] | 24.0 | yes | Defiler's Leather Girdle (20191, -0.00 DPS) [rep]; Officer's Belt (250556, -0.05 DPS) [crafted]; Defiler's Chain Girdle (20152, -0.06 DPS, sim-verified) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Golden Scale Leggings (3843, -0.14 DPS) [crafted]; Veteran's Chain Leggings (250493, -0.21 DPS) [crafted]; Veteran's Silvered Chain Leggings (250523, -1.14 DPS, sim-verified) [crafted] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 17.0 | yes | Hard Gold Boots (250534, +0.36 DPS, sim-verified) [crafted]; Brawler's Leather Boots (252439, -0.17 DPS) [crafted]; Glimmering Mail Greaves (4073, -0.18 DPS) [dungeon] |
| finger1 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 14.6 | yes | Silverlaine's Family Seal (6321, -0.16 DPS) [dungeon]; Band of the Fist (17694, -0.17 DPS) [quest]; Legionnaire's Band (20429, -0.17 DPS) [rep] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 11.9 | yes | Silverlaine's Family Seal (6321, -0.05 DPS, sim-verified) [dungeon]; Band of the Fist (17694, -0.08 DPS) [quest]; Insurgent's Band (272067, -0.10 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 345.4 | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Darkspear Raider's Reaper (272082, +0.00 DPS) [vendor] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 334.9 | yes | Commander's Crest (6320, -11.48 DPS) [dungeon]; Lambent Scale Shield (3656, -11.55 DPS) [dungeon]; Glimmering Shield (6400, -11.55 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Shining Silver Breastplate; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Defiler's Plate Girdle; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Legionnaire's Band; finger2: Ironspine's Eye; trinket1: Darkspear Voodoo Seal; trinket2: Defiler's Talisman; main_hand: Pronged Reaver; off_hand: Ironspine's Fist

No-known-source sample (15 of 913, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7956 Bronze Warhammer; 7958 Bronze Battle Axe; 9362 Brilliant Gold Ring

### Band 40 (orc, 0000000000000000-253130030005102051-0000000000000000)

Set DPS (verified): 87.2. Weights run: 0.9s. Verify run: 1.1s. 1251 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.386 ± 0.051, crit=7.710 ± 0.213, hit=5.187 ± 0.448, melee_haste=3.244 ± 0.122

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 133.9 | yes | Hard Gold Coif (250537, -0.51 DPS, sim-verified) [crafted]; White Bandit Mask (10008, -3.99 DPS) [crafted]; Tusken Helm (6686, -4.00 DPS) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Ethereal Talisman (4430, +0.32 DPS, sim-verified) [quest]; Scout's Medallion (19536, -0.36 DPS) [rep]; Scout's Medallion (19537, -0.40 DPS) [rep] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 | yes | Wrangling Spaulders (15698, -0.19 DPS) [quest]; Imperial Leather Spaulders (4737, -0.21 DPS, sim-verified) [dungeon]; Sunburn Spaulders (274751, -0.21 DPS) [vendor] |
| back | Sergeant Major's Cape (16336) | Rank 9 [pvp] | 14.3 | yes | Wildhunter Cloak (16658, -0.16 DPS) [quest]; Sergeant Major's Cape (16315, -0.18 DPS) [pvp]; Wolfmaster Cape (6314, -2.41 DPS, sim-verified) [dungeon] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 30.7 | yes | Golden Scale Cuirass (3845, -0.10 DPS) [crafted]; Mail Combat Armor (4074, -0.17 DPS) [dungeon]; Shining Silver Breastplate (2870, -0.46 DPS, sim-verified) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Ravager's Armguards (14770, -0.17 DPS) [world]; Pugilist Bracers (4438, -0.21 DPS, sim-verified) [dungeon]; Darkspear Armsplints (4132, -0.22 DPS) [quest] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 127.9 | yes | Fletcher's Gloves (7348, -0.74 DPS) [crafted]; Shadowskin Gloves (18238, -0.74 DPS) [crafted]; Dragonscale Gauntlets (8347, -0.87 DPS, sim-verified) [crafted] |
| waist | Defiler's Chain Girdle (20153) | The Defilers [rep] | 115.9 | yes | Defiler's Leather Girdle (20192, +0.45 DPS, sim-verified) [rep]; Tharg's Shoelace (9705, -3.33 DPS) [quest]; Scarlet Belt (10329, -3.40 DPS) [dungeon] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 | yes | Orcish War Leggings (7929, -0.53 DPS, sim-verified) [crafted]; Ferine Leggings (6690, -0.59 DPS) [dungeon]; Veteran's Silvered Chain Leggings (250523, -0.64 DPS) [crafted] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 24.7 | yes | Skirmisher's Mail Boots (252564, -0.09 DPS, sim-verified) [crafted]; Blackforge Greaves (6423, -0.11 DPS) [dungeon]; Skulker's Leather Shoes (252531, -0.24 DPS) [crafted] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 19.1 | yes | Insurgent's Band (272066, -0.26 DPS) [vendor]; Ironspine's Eye (7686, -0.28 DPS) [dungeon]; Legionnaire's Band (19513, -0.38 DPS, sim-verified) [rep] |
| finger2 | Suspicious Spare Part (274754) | Rettrick [vendor] | 14.0 | yes | Ironspine's Eye (7686, -0.09 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.15 DPS) [dungeon]; Insurgent's Band (272066, -0.16 DPS, sim-verified) [vendor] |
| trinket1 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [dungeon]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [dungeon]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Ardent Custodian (868) | Gnomeregan: Dark Iron Ambassador [dungeon] | 460.0 | yes | Staff of Jordan (873, +0.00 DPS) [dungeon]; Illusionary Rod (7713, +0.00 DPS) [dungeon]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor] |
| off_hand | Curve-bladed Ripper (2815) | Maraudon: Theradrim Shardling [dungeon] | 439.6 | yes | Pit Fighter's Shield (4507, -15.68 DPS) [quest]; Combat Shield (4065, -15.76 DPS) [dungeon]; Aegis of the Scarlet Commander (7726, -15.76 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Chain Girdle; legs: Scarlet Leggings; feet: Prowler's Leather Shoes; finger1: Legionnaire's Band; finger2: Suspicious Spare Part; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Ardent Custodian; off_hand: Curve-bladed Ripper

No-known-source sample (15 of 1251, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

### Band 50 (orc, 5500000000000000-253130030005102051-0000000000000000)

Set DPS (verified): 117.5. Weights run: 1.0s. Verify run: 1.2s. 1670 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.512 ± 0.061, crit=9.934 ± 0.268, hit=7.977 ± 0.662, melee_haste=3.814 ± 0.090

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 165.1 | yes | Eye of Theradras (17715, -0.57 DPS, sim-verified) [dungeon]; Hard Gold Coif (250537, -4.96 DPS) [crafted]; White Bandit Mask (10008, -4.97 DPS) [crafted] |
| neck | Woven Ivy Necklace (19159) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 16.6 | yes | Ethereal Talisman (4430, -0.16 DPS) [quest]; Scout's Medallion (19535, -0.38 DPS) [rep]; Ghostshard Talisman (7731, -0.85 DPS, sim-verified) [dungeon] |
| shoulder | Prowler's Leather Shoulder (252534) | Leatherworking [crafted] | 26.6 | yes | Skulker's Leather Shoulder (252535, -0.11 DPS) [crafted]; Hard Gold Pauldrons (250539, -0.17 DPS) [crafted]; Failed Flying Experiment (9647, -1.52 DPS, sim-verified) [quest] |
| back | Sergeant Major's Cape (16336) | Rank 9 [pvp] | 15.1 | yes | Pridelord Cape (14673, -0.16 DPS) [dungeon]; Sergeant Major's Cape (16315, -0.18 DPS) [pvp]; Bloodlust Cape (14801, -1.32 DPS, sim-verified) [dungeon] |
| chest | Wildthorn Mail (12624) | Blacksmithing [crafted] | 79.8 | yes | Kolkar Marauder Chain (6773, -0.42 DPS, sim-verified) [quest]; Warbear Harness (15064, -1.76 DPS) [crafted]; Relentless Chain (17777, -1.87 DPS) [quest] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 | yes | Prowler's Leather Bracers (252539, +0.16 DPS, sim-verified) [crafted]; Branded Leather Bracers (19508, -0.29 DPS) [dungeon]; Skulker's Leather Bracers (252540, -0.34 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 159.1 | yes | Dragonscale Gauntlets (8347, -0.66 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -0.72 DPS) [crafted]; Shadowskin Gloves (18238, -0.72 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20151) (or Defiler's Leather Girdle (20193)) | The Defilers [rep] | 159.1 | yes | Defiler's Leather Girdle (20193, +0.00 DPS, sim-verified) [rep]; Defiler's Chain Girdle (20153, -0.43 DPS) [rep]; Highlander's Mail Girdle (20118, -0.72 DPS) [vendor] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | 278.1 | yes | Scarlet Leggings (10330, -1.08 DPS, sim-verified) [dungeon]; Serpentskin Leggings (8262, -8.83 DPS) [dungeon]; Orcish War Leggings (7929, -8.83 DPS) [crafted] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 79.8 | yes | Prowler's Leather Boots (252468, -1.55 DPS, sim-verified) [crafted]; Skulker's Leather Boots (252469, -1.85 DPS) [crafted]; Skirmisher's Mail Sabatons (252578, -1.95 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 99.8 | yes | White Bone Band (11862, -2.74 DPS) [quest]; Band of Allegiance (18585, -2.96 DPS) [quest]; Insurgent's Band (272065, -3.07 DPS) [vendor] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 24.6 | yes | Legionnaire's Band (19512, -0.16 DPS) [rep]; Band of Allegiance (18585, -0.24 DPS) [quest]; White Bone Band (11862, -0.75 DPS, sim-verified) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 97.8 | yes | Tidal Charm (1404, -3.54 DPS) [vendor]; Guardian Talisman (1490, -3.54 DPS) [quest]; Ankh of Life (1713, -3.54 DPS) [dungeon] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, -4.11 DPS, sim-verified) [dungeon] |
| main_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | 649.1 | yes | Kindling Stave (11750, +0.00 DPS) [dungeon]; Bleakwood Hew (12769, +0.00 DPS) [crafted]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor] |
| off_hand | Hammer of the Northern Wind (810) | Blackrock Depths: Anvilrage Medic [dungeon] | 553.3 | yes | Claw of Celebras (17738, -2.57 DPS) [dungeon]; White Bone Shredder (11863, -3.86 DPS) [quest]; Shizzle's Drizzle Blocker (11915, -19.08 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 50:** neck: Woven Ivy Necklace; shoulder: Prowler's Leather Shoulder; chest: Wildthorn Mail; wrist: Bracers of the Stone Princess; waist: Defiler's Chain Girdle; legs: Stormshroud Pants; feet: Greaves of Withering Despair; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Darkspear Voodoo Seal; main_hand: Dawn's Edge; off_hand: Hammer of the Northern Wind

No-known-source sample (15 of 1670, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60 (orc, 5533220000000000-253130030005102051-0000000000000000)

Set DPS (verified): 209.8. Weights run: 1.1s. Verify run: 1.2s. 2549 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.519 ± 0.068, crit=10.313 ± 0.324, hit=11.159 ± 1.127, melee_haste=3.972 ± 0.142

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bloodvine Goggles (19999) | Engineering [crafted] | 367.6 | yes | Ragefury Eyepatch (11735, -2.41 DPS) [dungeon]; Bloodvine Lens (19998, -2.84 DPS) [crafted]; Mask of the Unforgiven (13404, -3.18 DPS, sim-verified) [dungeon] |
| neck | Onyxia Tooth Pendant (18404) | For All To See [quest] | 259.6 | yes | Medallion of the Dawn (22659, -3.29 DPS) [quest]; Beads of Ogre Might (22150, -4.47 DPS) [quest]; Choker of the Shifting Sands (21505, -7.85 DPS) [quest] |
| shoulder | Champion's Mail Pauldrons (227154) | Rank 14 [pvp] | 282.0 | yes | Champion's Mail Pauldrons (23260, -0.39 DPS, sim-verified) [vendor]; Stormshroud Shoulders (15058, -4.96 DPS) [crafted]; Shroud of the Nathrezim (18720, -4.96 DPS) [dungeon] |
| back | Earthweave Cloak (21187) | Volunteer's Battlegear [quest] | 119.4 | yes | Chromatic Cloak (18509, -2.10 DPS, sim-verified) [crafted]; Deathguard's Cloak (20068, -2.98 DPS) [rep]; Cloak of the Fallen God (21710, -3.02 DPS) [quest] |
| chest | Legionnaire's Mail Hauberk (227157) | Rank 12 [pvp] | 300.0 | yes | Stormshroud Armor (15056, -0.40 DPS) [crafted]; Bloodsoul Breastplate (19690, -1.22 DPS, sim-verified) [crafted]; Savage Gladiator Chain (11726, -1.59 DPS) [dungeon] |
| wrist | Primal Batskin Bracers (19687) | Leatherworking [crafted] | 118.9 | yes | Rockfury Bracers (21186, -0.89 DPS, sim-verified) [quest]; Windtalker's Wristguards (19582, -2.92 DPS) [rep]; Windtalker's Wristguards (19583, -3.06 DPS) [rep] |
| hands | Blood Guard's Mail Vices (227159) | Rank 11 [pvp] | 276.0 | yes | Primal Batskin Gloves (19686, -1.72 DPS) [crafted]; Stormshroud Gloves (21278, -1.96 DPS, sim-verified) [crafted]; Chromatic Gauntlets (19157, -3.16 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20150) (or Defiler's Leather Girdle (20190)) | The Defilers [rep] | 178.4 | yes | Defiler's Leather Girdle (20190, +0.00 DPS, sim-verified) [rep]; Light Obsidian Belt (22195, -0.07 DPS) [crafted]; Defiler's Chain Girdle (20151, -0.50 DPS) [rep] |
| legs | Legionnaire's Mail Legguards (227156) | Rank 12 [pvp] | 300.0 | yes | Outrider's Chain Leggings (22673, -0.93 DPS) [rep]; Stormshroud Pants (15057, -2.60 DPS, sim-verified) [crafted]; Devilsaur Leggings (15062, -3.95 DPS) [crafted] |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 111.6 | yes | Greaves of Withering Despair (22240, +0.00 DPS) [dungeon]; Shadowcraft Boots (16711, -2.42 DPS) [dungeon]; Blood Guard's Mail Greaves (227158, -3.79 DPS, sim-verified) [pvp] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 272.0 | yes | Master Dragonslayer's Ring (19384, -4.05 DPS) [quest]; Band of the Penitent (13217, -4.60 DPS) [quest]; Dragonslayer's Signet (18403, -4.60 DPS) [quest] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | 268.0 | yes | Master Dragonslayer's Ring (19384, +0.34 DPS, sim-verified) [quest]; Band of the Penitent (13217, -4.46 DPS) [quest]; Dragonslayer's Signet (18403, -4.46 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 120.1 | yes | Tidal Charm (1404, -4.33 DPS) [vendor]; Guardian Talisman (1490, -4.33 DPS) [quest]; Blazing Emblem (2802, -4.33 DPS) [dungeon] |
| trinket2 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest]; Blazing Emblem (2802, +0.00 DPS) [dungeon] |
| main_hand | High Warlord's Cleaver (234554) | Rank 18 [pvp] | 1036.5 | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp] |
| off_hand | High Warlord's Bludgeon (234555) | Rank 18 [pvp] | 1036.5 | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Grand Marshal's Left Hand Blade (18847, -1.13 DPS) [vendor] |
| ranged | - | - |  |  |  |

**New at 60:** head: Bloodvine Goggles; neck: Onyxia Tooth Pendant; shoulder: Champion's Mail Pauldrons; back: Earthweave Cloak; chest: Legionnaire's Mail Hauberk; wrist: Primal Batskin Bracers; hands: Blood Guard's Mail Vices; waist: Defiler's Chain Girdle; legs: Legionnaire's Mail Legguards; feet: Bloodvine Boots; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket2: Ankh of Life; main_hand: High Warlord's Cleaver; off_hand: High Warlord's Bludgeon

No-known-source sample (15 of 2549, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6478 Rat Stompers

