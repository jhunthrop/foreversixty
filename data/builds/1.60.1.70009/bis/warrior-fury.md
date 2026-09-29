# Leveling BiS: Fury

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

## Alliance

### Band 20 (human, 00000000000000000-353000000000000000-000000000000000000)

Set DPS (verified): 32.3. Weights run: 0.9s. Verify run: 1.1s. 537 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.034, strength=2.084 ± 0.042, agility=0.085 ± 0.015, crit=2.317 ± 0.066, hit=1.130 ± 0.086, melee_haste=1.686 ± 0.042

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.8 | yes | Defender's Leather Hood (252447, -0.33 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.74 DPS) [crafted]; Brawler's Leather Hood (252504, -0.75 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 0.5 | yes | Tarnished Locket (279870, -0.08 DPS, sim-verified) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.3 | yes | Silvered Bronze Shoulders (3481, +0.00 DPS, sim-verified) [crafted]; Serpent's Shoulders (5404, -0.22 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.23 DPS) [crafted] |
| back | Lambent Scale Cloak (4706) | Gnomeregan: Caverndeep Burrower [dungeon] | 8.3 | yes | Catacomb Cloak (279899, -0.09 DPS) [quest]; Grave Shroud (279865, -0.10 DPS, sim-verified) [quest]; Dark Leather Cloak (2316, -0.15 DPS) [crafted] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.8 | yes | Defender's Leather Armor (252434, -0.22 DPS) [crafted]; Totemic Leather Armor (252435, -0.23 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.34 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.4 | yes | Cryptwalker Bracers (280095, -0.14 DPS, sim-verified) [quest]; Bravo's Armbands (270015, -0.22 DPS) [quest]; Runed Copper Bracers (2854, -0.23 DPS) [crafted] |
| hands | Gold-flecked Gloves (5195) | The Deadmines: Sneed [dungeon] | 14.6 | yes | Polar Gauntlets (7606, -0.08 DPS) [quest]; Blackened Defias Gloves (10401, -0.08 DPS) [dungeon]; Fletcher's Gloves (7348, -0.49 DPS, sim-verified) [crafted] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 | yes | Ruffian Belt (5975, -0.20 DPS) [world]; Cobrahn's Grasp (6460, -0.22 DPS, sim-verified) [dungeon]; Support Girdle (1215, -0.28 DPS) [world] |
| legs | Chausses of Westfall (6087) | The Defias Brotherhood [quest] | 22.9 | yes | Defender's Leather Pants (252445, -0.14 DPS) [crafted]; Veteran's Chain Leggings (250493, -0.15 DPS, sim-verified) [crafted]; Totemic Leather Pants (252446, -0.15 DPS) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.8 | yes | Veteran's Boots (250503, +0.00 DPS, sim-verified) [crafted]; Guard's Boots (250504, -0.02 DPS) [crafted]; Defender's Leather Boots (252441, -0.02 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.7 | yes | Signet of the Zhevra (285330, -0.30 DPS) [world]; Lavishly Jeweled Ring (1156, -0.32 DPS) [dungeon]; Minor Channeling Ring (1449, -0.32 DPS) [quest] |
| finger2 | The 1 Ring (8350) | Fishing [world] | 2.2 | yes | Signet of the Zhevra (285330, -0.00 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.07 DPS) [dungeon]; Minor Channeling Ring (1449, -0.08 DPS) [quest] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 | yes | Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest]; Hammerbone (270018, +0.00 DPS) [quest] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 237.1 | yes | Bear Buckler (4821, -8.57 DPS) [vendor]; Furen's Favor (6970, -8.57 DPS) [quest]; Veteran Shield (3651, -8.65 DPS) [dungeon] |
| ranged | Dwarven Fishing Pole (3567) (or Cracked Blacksmith Hammer (285279)) | Murloc Poachers [quest] | 4.2 | yes | Cracked Blacksmith Hammer (285279, +0.00 DPS, sim-verified) [crafted]; Fine Longbow (11304, -0.01 DPS) [vendor]; Daryl's Hunting Rifle (2904, -0.08 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Gold-flecked Gloves; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: The 1 Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Dwarven Fishing Pole

No-known-source sample (15 of 537, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 7956 Bronze Warhammer; 8178 Training Sword; 9602 Brushwood Blade; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse

### Band 30 (human, 00000000000000000-353211005010000000-000000000000000000)

Set DPS (verified): 53.1. Weights run: 0.9s. Verify run: 1.2s. 1006 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.160, strength=1.676 ± 0.208, agility=not significant (0.186 ± 0.061), crit=3.475 ± 0.238, hit=1.401 ± 0.221, melee_haste=1.775 ± 0.290

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 21.8 | yes | Veteran's Chain Helm (250498, -0.07 DPS, sim-verified) [crafted]; Defender's Leather Helm (252455, -0.08 DPS) [crafted]; Crusader's Chain Helm (250502, -0.15 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (20444, -0.59 DPS) [rep]; Pendant of Myzrael (4614, -0.64 DPS) [dungeon]; Sentinel's Medallion (19541, -0.84 DPS, sim-verified) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 11.7 | yes | Mail Combat Spaulders (6404, +0.00 DPS, sim-verified) [dungeon]; Barbaric Iron Shoulders (7913, -0.03 DPS) [crafted]; Forest Tracker Epaulets (2278, -0.06 DPS) [dungeon] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Sergeant Major's Cape (16315, +0.21 DPS, sim-verified) [pvp]; Lambent Scale Cloak (4706, -0.15 DPS) [dungeon]; Catacomb Cloak (279899, -0.18 DPS) [quest] |
| chest | Shining Silver Breastplate (2870) | Blacksmithing [crafted] | 23.5 | yes | Barbaric Iron Breastplate (7914, -0.18 DPS, sim-verified) [crafted]; Hard Gold Cuirass (250533, -0.23 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.25 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 13.4 | yes | Cultist's Armguards (270032, -0.16 DPS) [quest]; Bands of Serra'kis (6902, -0.18 DPS, sim-verified) [dungeon]; Patterned Bronze Bracers (2868, -0.23 DPS) [crafted] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 21.0 | yes | Heavy Earthen Gloves (7359, -0.23 DPS) [crafted]; Bonefist Gauntlets (4465, -0.27 DPS) [world]; Fletcher's Gloves (7348, -0.97 DPS, sim-verified) [crafted] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 | yes | Highlander's Leather Girdle (20117, +0.00 DPS, sim-verified) [rep]; Highlander's Plate Girdle (20126, -0.18 DPS) [rep]; Highlander's Lamellar Girdle (20108, -0.25 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Veteran's Silvered Chain Leggings (250523, +0.21 DPS, sim-verified) [crafted]; Golden Scale Leggings (3843, -0.35 DPS) [crafted]; Chausses of Westfall (6087, -0.35 DPS) [quest] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 13.0 | yes | Glimmering Mail Greaves (4073, -0.14 DPS) [dungeon]; Brawler's Leather Boots (252439, -0.17 DPS) [crafted]; Hard Gold Boots (250534, -0.26 DPS, sim-verified) [crafted] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 11.2 | yes | Silverlaine's Family Seal (6321, -0.13 DPS) [dungeon]; Protector's Band (20439, -0.17 DPS) [rep]; Insurgent's Band (272067, -1.22 DPS, sim-verified) [vendor] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 8.4 | yes | Silverlaine's Family Seal (6321, -0.00 DPS) [dungeon]; Seal of Wrynn (2933, -0.13 DPS) [quest]; Insurgent's Band (272067, -0.56 DPS, sim-verified) [vendor] |
| trinket1 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 343.8 | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Morbid Dawn (7689, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 332.6 | yes | Shoni's Disarming Tool (9608, -4.92 DPS) [quest]; Commander's Crest (6320, -14.80 DPS) [dungeon]; Lambent Scale Shield (3656, -14.87 DPS) [dungeon] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 | yes | Double-barreled Shotgun (2098, +0.04 DPS, sim-verified) [dungeon]; Moonsight Rifle (4383, -0.22 DPS) [crafted]; Precision Bow (217315, -0.22 DPS) [quest] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wolfmaster Cape; chest: Shining Silver Breastplate; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Protector's Band; finger2: Ironspine's Eye; trinket1: Talisman of Arathor; trinket2: Darkspear Voodoo Seal; main_hand: Pronged Reaver; off_hand: Ironspine's Fist; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 1006, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 7956 Bronze Warhammer

### Band 40 (human, 00000000000000000-353211005050010050-000000000000000000)

Set DPS (verified): 88.1. Weights run: 0.9s. Verify run: 1.2s. 1451 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.203, strength=1.289 ± 0.264, agility=not significant (0.291 ± 0.084), crit=4.325 ± 0.374, hit=1.889 ± 0.312, melee_haste=2.652 ± 0.374

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 77.3 | yes | Icemetal Barbute (10763, -1.90 DPS, sim-verified) [dungeon]; Hard Gold Coif (250537, -4.40 DPS) [crafted]; White Bandit Mask (10008, -4.45 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19540, +0.02 DPS, sim-verified) [rep]; Sentinel's Medallion (19541, -0.87 DPS) [rep]; Sentinel's Medallion (20444, -0.91 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 15.2 | yes | Hard Gold Pauldrons (250539, +0.63 DPS, sim-verified) [crafted]; Shining Mithril Pauldrons (250541, -0.17 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.27 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | Rank 9 [pvp] | 9.5 | yes | Sergeant Major's Cape (16315, -0.23 DPS) [pvp]; Catacomb Cloak (279899, -0.26 DPS) [quest]; Wolfmaster Cape (6314, -1.18 DPS, sim-verified) [dungeon] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 20.1 | yes | Golden Scale Cuirass (3845, -0.15 DPS) [crafted]; Shining Mithril Breastplate (250540, -0.15 DPS) [crafted]; Shining Silver Breastplate (2870, -0.68 DPS, sim-verified) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Pugilist Bracers (4438, -0.29 DPS, sim-verified) [dungeon]; Ravager's Armguards (14770, -0.73 DPS) [world]; Cultist's Armguards (270032, -0.74 DPS) [quest] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 80.6 | yes | Dragonscale Gauntlets (8347, -0.86 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -1.49 DPS) [crafted]; Ornate Mithril Gloves (7927, -1.49 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20089) | The League of Arathor [rep] | 68.6 | yes | Highlander's Leather Girdle (20116, +0.38 DPS, sim-verified) [rep]; Highlander's Chain Girdle (20090, -3.31 DPS) [rep]; Highlander's Leather Girdle (20117, -3.31 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 27.1 | yes | Orcish War Leggings (7929, -0.38 DPS) [crafted]; Ferine Leggings (6690, -0.76 DPS, sim-verified) [dungeon]; Veteran's Silvered Chain Leggings (250523, -0.81 DPS) [crafted] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 18.8 | yes | Prowler's Leather Shoes (252465, -0.29 DPS, sim-verified) [crafted]; Blackforge Greaves (6423, -0.32 DPS) [dungeon]; Skirmisher's Mail Boots (252564, -0.34 DPS) [crafted] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 12.6 | yes | Protector's Band (19517, -0.23 DPS) [rep]; Suspicious Spare Part (274754, -0.27 DPS) [vendor]; Ironspine's Eye (7686, -0.36 DPS) [dungeon] |
| finger2 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 12.0 | yes | Suspicious Spare Part (274754, +0.78 DPS, sim-verified) [vendor]; Insurgent's Band (272067, -0.22 DPS) [vendor]; Ironspine's Eye (7686, -0.31 DPS) [dungeon] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [dungeon]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [dungeon]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Ardent Custodian (868) | Gnomeregan: Dark Iron Ambassador [dungeon] | 460.0 | yes | Staff of Jordan (873, +0.00 DPS) [dungeon]; Bonebiter (6830, +0.00 DPS) [quest]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 442.0 | yes | Shoni's Disarming Tool (9608, -16.14 DPS) [quest]; Salbac Shield (4652, -31.96 DPS) [quest]; Combat Shield (4065, -32.16 DPS) [dungeon] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 | yes | Explosive Shotgun (8188, +0.27 DPS, sim-verified) [world]; Mithril Blacksmith Hammer (285280, -0.19 DPS) [crafted]; Master Hunter's Rifle (17687, -0.20 DPS) [quest] |

**New at 40:** head: Raging Berserker's Helm; shoulder: Sunburn Spaulders; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Chain Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Protector's Band; finger2: Insurgent's Band; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Ardent Custodian; off_hand: Vanquisher's Sword

No-known-source sample (15 of 1451, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 50 (human, 35100000000000000-353211005050010051-000000000000000000)

Set DPS (verified): 92.5. Weights run: 1.0s. Verify run: 1.3s. 1949 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.217, strength=1.796 ± 0.296, agility=not significant (0.302 ± 0.101), crit=4.820 ± 0.391, hit=2.018 ± 0.331, melee_haste=2.822 ± 0.403

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Plate Helm (220803) (or Knight-Lieutenant's Plate Helm (220804)) | Lady Palanseer [vendor] | 111.0 | yes | Knight-Lieutenant's Plate Helm (220804, +0.00 DPS, sim-verified) [vendor]; Raging Berserker's Helm (7719, -1.94 DPS) [dungeon]; Ornate Mithril Helm (7937, -2.46 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19539, -0.72 DPS, sim-verified) [rep]; Talisman of the Naga Lord (5029, -1.00 DPS) [world]; Sentinel's Medallion (19540, -1.03 DPS) [rep] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) (or Blood Guard's Plate Pauldrons (220796)) | Captain Dirgehammer [vendor] | 87.2 | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS, sim-verified) [vendor]; Officer's Pauldrons (250576, -5.91 DPS) [crafted]; Wyrmslayer Spaulders (13066, -6.09 DPS) [world] |
| back | Bloodlust Cape (14801) | Maraudon: Princess Theradras [dungeon] | 16.2 | yes | Sergeant Major's Cape (16336, +0.14 DPS, sim-verified) [pvp]; Wolfmaster Cape (6314, -0.59 DPS) [dungeon]; Sergeant Major's Cape (16315, -0.75 DPS) [pvp] |
| chest | Knight's Plate Hauberk (220794) (or Stone Guard's Plate Armor (220801)) | Captain Dirgehammer [vendor] | 94.4 | yes | Stone Guard's Plate Armor (220801, +0.00 DPS, sim-verified) [vendor]; Ornate Mithril Breastplate (7935, -2.59 DPS) [crafted]; Warforged Chestplate (11195, -4.94 DPS) [quest] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 | yes | Officer's Wristguards (250581, +0.14 DPS, sim-verified) [crafted]; Branded Leather Bracers (19508, -0.77 DPS) [dungeon]; Prowler's Leather Bracers (252539, -0.94 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 87.5 | yes | Dragonscale Gauntlets (8347, -1.77 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -1.92 DPS) [crafted]; Ornate Mithril Gloves (7927, -1.92 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20088) (or Highlander's Leather Girdle (20115)) | The League of Arathor [rep] | 87.5 | yes | Highlander's Leather Girdle (20115, +0.00 DPS, sim-verified) [rep]; Highlander's Lamellar Girdle (20106, -0.02 DPS) [rep]; Highlander's Plate Girdle (20124, -0.20 DPS) [rep] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | 92.1 | yes | Stone Guard's Plate Leggings (220798, +0.00 DPS) [vendor]; Stormshroud Pants (15057, -1.96 DPS, sim-verified) [crafted]; Scarlet Leggings (10330, -5.23 DPS) [dungeon] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 26.7 | yes | Officer's Sabatons (250561, -0.06 DPS, sim-verified) [crafted]; Officer's Boots (250546, -0.12 DPS) [crafted]; Skulker's Leather Boots (252469, -0.29 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 40.2 | yes | Insurgent's Band (272065, -2.42 DPS) [vendor]; Suspicious Spare Part (274754, -2.66 DPS) [vendor]; Insurgent's Band (272066, -2.71 DPS) [vendor] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 20.7 | yes | Protector's Band (19515, -0.49 DPS, sim-verified) [rep]; Insurgent's Band (272065, -0.55 DPS) [vendor]; Protector's Band (19517, -0.78 DPS) [rep] |
| trinket1 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest] |
| trinket2 | Talisman of Arathor (21117) | The League of Arathor [rep] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest] |
| main_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | 577.5 | yes | Kindling Stave (11750, +0.00 DPS) [dungeon]; Bleakwood Hew (12769, +0.00 DPS) [crafted]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor] |
| off_hand | Hammer of the Northern Wind (810) | Blackrock Depths: Anvilrage Medic [dungeon] | 553.3 | yes | Claw of Celebras (17738, -6.84 DPS) [dungeon]; Shoni's Disarming Tool (9608, -31.61 DPS) [quest]; Shizzle's Drizzle Blocker (11915, -50.98 DPS) [quest] |
| ranged | Houndmaster's Bow (11628) | Blackrock Depths: Houndmaster Grebmar [dungeon] | 17.4 | yes | Arcanite Blacksmith Hammer (285281, -0.64 DPS) [crafted]; Booty Bay Bruiser's Buckshot (274748, -0.81 DPS) [vendor]; Explosive Shotgun (8188, -0.81 DPS) [world] |

**New at 50:** head: Blood Guard's Plate Helm; shoulder: Knight-Lieutenant's Plate Pauldrons; back: Bloodlust Cape; chest: Knight's Plate Hauberk; wrist: Bracers of the Stone Princess; waist: Highlander's Chain Girdle; legs: Knight's Plate Leggings; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Smoking Heart of the Mountain; trinket2: Talisman of Arathor; main_hand: Dawn's Edge; off_hand: Hammer of the Northern Wind; ranged: Houndmaster's Bow

No-known-source sample (15 of 1949, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (human, 35311103002000000-353211005050010051-000000000000000000)

Set DPS (verified): 269.7. Weights run: 1.0s. Verify run: 1.4s. 2882 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=not significant (1.000 ± 0.427), strength=not significant (2.344 ± 0.596), agility=not significant (0.450 ± 0.239), crit=13.060 ± 0.873, hit=not significant (0.000 ± 0.000), melee_haste=6.782 ± 0.884

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 407.9 | yes | Bloodvine Lens (19998, -2.44 DPS) [crafted]; Field Marshal's Plate Helm (16478, -9.20 DPS) [vendor]; Ragefury Eyepatch (11735, -9.65 DPS, sim-verified) [dungeon] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 206.8 | yes | Onyxia Tooth Pendant (18404, -1.20 DPS) [quest]; Choker of the Shifting Sands (21505, -9.52 DPS) [quest]; Pendant of the Shifting Sands (21506, -10.05 DPS) [quest] |
| shoulder | Champion's Plate Shoulders (23243) (or Lieutenant Commander's Plate Shoulders (23315), Champion's Plate Shoulders (227042), Lieutenant Commander's Plate Shoulders (227045)) | Lady Palanseer [vendor] | 222.7 | yes | Lieutenant Commander's Plate Shoulders (23315, +0.00 DPS, sim-verified) [vendor]; Champion's Plate Shoulders (227042, +0.00 DPS) [pvp]; Lieutenant Commander's Plate Shoulders (227045, +0.00 DPS) [pvp] |
| back | Drape of Unyielding Strength (21394) | Drape of Unyielding Strength [quest] | 39.2 | yes | Cloak of the Fallen God (21710, -0.10 DPS) [quest]; Cloak of the Honor Guard (20073, -0.17 DPS) [rep]; Chromatic Cloak (18509, -7.91 DPS, sim-verified) [crafted] |
| chest | Bloodsoul Breastplate (19690) | Blacksmithing [crafted] | 369.7 | yes | Stormshroud Armor (15056, -1.01 DPS, sim-verified) [crafted]; Savage Gladiator Chain (11726, -2.98 DPS) [dungeon]; Obsidian Mail Tunic (22191, -6.40 DPS) [crafted] |
| wrist | Deeprock Bracers (21184) | Stalwart's Battlegear [quest] | 49.0 | yes | Berserker Bracers (19578, -0.05 DPS) [rep]; Marshal's Plate Bracers (16481, -0.35 DPS) [pvp]; Vambraces of the Sadist (13400, -2.61 DPS, sim-verified) [dungeon] |
| hands | Marshal's Plate Gauntlets (16484) | Captain Dirgehammer [vendor] | 229.7 | yes | General's Plate Gauntlets (16548, +0.00 DPS, sim-verified) [vendor]; General's Plate Gauntlets (231532, +0.00 DPS) [pvp]; Marshal's Plate Gauntlets (231541, +0.00 DPS) [pvp] |
| waist | Zandalar Vindicator's Belt (19823) | Paragons of Power: The Vindicator's Belt [quest] | 241.4 | yes | Highlander's Plate Girdle (20041, -0.29 DPS, sim-verified) [rep]; Highlander's Lamellar Girdle (20042, -1.35 DPS) [rep]; Highlander's Chain Girdle (20043, -1.42 DPS) [rep] |
| legs | Marshal's Plate Legguards (16479) (or General's Plate Leggings (16543), General's Plate Leggings (231533), Marshal's Plate Legguards (231540)) | Captain Dirgehammer [vendor] | 412.6 | yes | General's Plate Leggings (16543, +0.00 DPS, sim-verified) [vendor]; General's Plate Leggings (231533, +0.00 DPS) [pvp]; Marshal's Plate Legguards (231540, +0.00 DPS) [pvp] |
| feet | Conqueror's Greaves (21333) | Conqueror's Greaves [quest] | 56.9 | yes | General's Plate Boots (231531, +2.54 DPS, sim-verified) [pvp]; Marshal's Plate Boots (16483, -0.54 DPS) [vendor]; General's Plate Boots (16545, -0.54 DPS) [vendor] |
| finger1 | Signet of Unyielding Strength (21393) | Signet of Unyielding Strength [quest] | 208.6 | yes | Band of Earthen Might (21182, -0.68 DPS) [quest]; Band of the Penitent (13217, -1.49 DPS) [quest]; Dragonslayer's Signet (18403, -1.49 DPS) [quest] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | 198.8 | yes | Band of Earthen Might (21182, -0.83 DPS, sim-verified) [quest]; Band of the Penitent (13217, -0.92 DPS) [quest]; Dragonslayer's Signet (18403, -0.92 DPS) [quest] |
| trinket1 | Onyxia Blood Talisman (18406) | For All To See [quest] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest] |
| trinket2 | Talisman of Arathor (20071) | The League of Arathor [rep] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest] |
| main_hand | High Warlord's Quickblade (234553) | Rank 18 [pvp] | 1078.0 | yes | High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Pulverizer (234545, +0.00 DPS) [pvp] |
| off_hand | Grand Marshal's Swiftblade (234579) | Rank 18 [pvp] | 1078.0 | yes | High Warlord's Left Claw (234558, -0.18 DPS) [pvp]; Grand Marshal's Left Hand Blade (234584, -0.18 DPS) [pvp]; Grand Marshal's Left Hand Blade (18847, -1.99 DPS) [vendor] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | 182.8 | yes | Bloodseeker (19107, -9.29 DPS) [quest]; Houndmaster's Bow (11628, -9.46 DPS) [dungeon]; Arcanite Blacksmith Hammer (285281, -9.74 DPS) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Medallion of the Dawn; shoulder: Champion's Plate Shoulders; back: Drape of Unyielding Strength; chest: Bloodsoul Breastplate; wrist: Deeprock Bracers; hands: Marshal's Plate Gauntlets; waist: Zandalar Vindicator's Belt; legs: Marshal's Plate Legguards; feet: Conqueror's Greaves; finger1: Signet of Unyielding Strength; finger2: Don Julio's Band; trinket1: Onyxia Blood Talisman; trinket2: Talisman of Arathor; main_hand: High Warlord's Quickblade; off_hand: Grand Marshal's Swiftblade; ranged: The Purifier

No-known-source sample (15 of 2882, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

## Horde

### Band 20 (troll, 00000000000000000-353000000000000000-000000000000000000)

Set DPS (verified): 29.0. Weights run: 0.9s. Verify run: 1.0s. 531 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.034, strength=2.084 ± 0.042, agility=0.085 ± 0.015, crit=2.317 ± 0.066, hit=1.130 ± 0.086, melee_haste=1.686 ± 0.042

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.8 | yes | Defender's Leather Hood (252447, -0.34 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.74 DPS) [crafted]; Brawler's Leather Hood (252504, -0.75 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 0.5 | yes | Tarnished Locket (279870, -0.15 DPS, sim-verified) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.3 | yes | Silvered Bronze Shoulders (3481, +0.00 DPS, sim-verified) [crafted]; Serpent's Shoulders (5404, -0.22 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.23 DPS) [crafted] |
| back | Lambent Scale Cloak (4706) | Gnomeregan: Caverndeep Burrower [dungeon] | 8.3 | yes | Subterranean Cape (14149, -0.08 DPS) [dungeon]; Catacomb Cloak (279899, -0.09 DPS) [quest]; Grave Shroud (279865, -0.16 DPS, sim-verified) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.8 | yes | Defender's Leather Armor (252434, -0.22 DPS) [crafted]; Totemic Leather Armor (252435, -0.23 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.45 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.4 | yes | Bravo's Armbands (270015, -0.22 DPS) [quest]; Runed Copper Bracers (2854, -0.23 DPS) [crafted]; Cryptwalker Bracers (280095, -0.25 DPS, sim-verified) [quest] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 32.4 | yes | Gold-flecked Gloves (5195, +0.08 DPS, sim-verified) [dungeon]; Blackened Defias Gloves (10401, -0.74 DPS) [dungeon]; Dagmire Gauntlets (6481, -0.81 DPS) [quest] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 | yes | Ruffian Belt (5975, -0.20 DPS) [world]; Cobrahn's Grasp (6460, -0.27 DPS, sim-verified) [dungeon]; Support Girdle (1215, -0.28 DPS) [world] |
| legs | Veteran's Chain Leggings (250493) | Blacksmithing [crafted] | 19.2 | yes | Totemic Leather Pants (252446, -0.02 DPS) [crafted]; Defender's Leather Pants (252445, -0.02 DPS, sim-verified) [crafted]; Deepgrave Trousers (279900, -0.16 DPS) [quest] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.8 | yes | Guard's Boots (250504, -0.02 DPS) [crafted]; Defender's Leather Boots (252441, -0.02 DPS) [crafted]; Veteran's Boots (250503, -0.02 DPS, sim-verified) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.7 | yes | Signet of the Zhevra (285330, -0.30 DPS) [world]; Bounty Hunter's Ring (5351, -0.31 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.32 DPS) [dungeon] |
| finger2 | The 1 Ring (8350) | Fishing [world] | 2.2 | yes | Signet of the Zhevra (285330, -0.03 DPS, sim-verified) [world]; Bounty Hunter's Ring (5351, -0.07 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.07 DPS) [dungeon] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 | yes | Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest]; Hammerbone (270018, +0.00 DPS) [quest] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 237.1 | yes | Bear Buckler (4821, -8.57 DPS) [vendor]; Ruga's Bulwark (7120, -8.57 DPS) [quest]; Faerleia's Shield (3450, -8.65 DPS) [quest] |
| ranged | Cracked Blacksmith Hammer (285279) | Blacksmithing [crafted] | 4.2 | yes | Fine Longbow (11304, +0.00 DPS, sim-verified) [vendor]; Heavy Shortbow (3036, -0.08 DPS) [dungeon]; Orcish Battle Bow (5346, -0.08 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: The 1 Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Cracked Blacksmith Hammer

No-known-source sample (15 of 531, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 7956 Bronze Warhammer; 8178 Training Sword; 9602 Brushwood Blade; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse

### Band 30 (troll, 00000000000000000-353211005010000000-000000000000000000)

Set DPS (verified): 53.1. Weights run: 0.9s. Verify run: 1.2s. 1000 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.160, strength=1.676 ± 0.208, agility=not significant (0.186 ± 0.061), crit=3.475 ± 0.238, hit=1.401 ± 0.221, melee_haste=1.775 ± 0.290

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 21.8 | yes | Veteran's Chain Helm (250498, +0.06 DPS, sim-verified) [crafted]; Defender's Leather Helm (252455, -0.08 DPS) [crafted]; Crusader's Chain Helm (250502, -0.15 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19537, -0.38 DPS, sim-verified) [rep]; Scout's Medallion (20442, -0.59 DPS) [rep]; Pendant of Myzrael (4614, -0.64 DPS) [dungeon] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 11.7 | yes | Mail Combat Spaulders (6404, +0.00 DPS, sim-verified) [dungeon]; Barbaric Iron Shoulders (7913, -0.03 DPS) [crafted]; Forest Tracker Epaulets (2278, -0.06 DPS) [dungeon] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Sergeant Major's Cape (16315, -0.12 DPS) [pvp]; Lambent Scale Cloak (4706, -0.15 DPS) [dungeon] |
| chest | Shining Silver Breastplate (2870) | Blacksmithing [crafted] | 23.5 | yes | Barbaric Iron Breastplate (7914, +0.07 DPS, sim-verified) [crafted]; Hard Gold Cuirass (250533, -0.23 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.25 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 13.4 | yes | Bands of Serra'kis (6902, +0.07 DPS, sim-verified) [dungeon]; Cultist's Armguards (270032, -0.16 DPS) [quest]; Grimtoll Wristguards (15459, -0.22 DPS) [quest] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 21.0 | yes | Warsong Gauntlets (16978, -0.19 DPS) [quest]; Heavy Earthen Gloves (7359, -0.23 DPS) [crafted]; Fletcher's Gloves (7348, -0.77 DPS, sim-verified) [crafted] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 | yes | Defiler's Leather Girdle (20191, +0.00 DPS, sim-verified) [rep]; Defiler's Plate Girdle (20207, -0.18 DPS) [rep]; Blackened Defias Belt (10403, -0.27 DPS) [dungeon] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 19.7 | yes | Golden Scale Leggings (3843, -0.06 DPS) [crafted]; Juggernaut Leggings (6671, -0.14 DPS) [quest]; Ferine Leggings (6690, -0.53 DPS, sim-verified) [dungeon] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 13.0 | yes | Glimmering Mail Greaves (4073, -0.14 DPS) [dungeon]; Brawler's Leather Boots (252439, -0.17 DPS) [crafted]; Hard Gold Boots (250534, -0.30 DPS, sim-verified) [crafted] |
| finger1 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 11.2 | yes | Ironspine's Eye (7686, -0.13 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.13 DPS) [dungeon]; Band of the Fist (17694, -0.17 DPS) [quest] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 | yes | Ironspine's Eye (7686, +0.31 DPS, sim-verified) [dungeon]; Silverlaine's Family Seal (6321, -0.03 DPS) [dungeon]; Band of the Fist (17694, -0.07 DPS) [quest] |
| trinket1 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 343.8 | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Morbid Dawn (7689, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 332.6 | yes | Commander's Crest (6320, -14.80 DPS) [dungeon]; Lambent Scale Shield (3656, -14.87 DPS) [dungeon]; Glimmering Shield (6400, -14.87 DPS) [dungeon] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 | yes | Double-barreled Shotgun (2098, +0.37 DPS, sim-verified) [dungeon]; Moonsight Rifle (4383, -0.22 DPS) [crafted]; Precision Bow (217315, -0.22 DPS) [quest] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wolfmaster Cape; chest: Shining Silver Breastplate; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Defiler's Chain Girdle; legs: Veteran's Silvered Chain Leggings; feet: Trouncing Boots; finger1: Legionnaire's Band; finger2: Insurgent's Band; trinket1: Defiler's Talisman; trinket2: Darkspear Voodoo Seal; main_hand: Pronged Reaver; off_hand: Ironspine's Fist; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 1000, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 7956 Bronze Warhammer

### Band 40 (troll, 00000000000000000-353211005050010050-000000000000000000)

Set DPS (verified): 83.5. Weights run: 0.9s. Verify run: 1.3s. 1445 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.203, strength=1.289 ± 0.264, agility=not significant (0.291 ± 0.084), crit=4.325 ± 0.374, hit=1.889 ± 0.312, melee_haste=2.652 ± 0.374

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 77.3 | yes | Icemetal Barbute (10763, -1.83 DPS, sim-verified) [dungeon]; Hard Gold Coif (250537, -4.40 DPS) [crafted]; White Bandit Mask (10008, -4.45 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Ethereal Talisman (4430, +0.16 DPS, sim-verified) [quest]; Scout's Medallion (19536, -0.80 DPS) [rep]; Scout's Medallion (19537, -0.87 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 15.2 | yes | Shining Mithril Pauldrons (250541, -0.17 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.27 DPS) [dungeon]; Hard Gold Pauldrons (250539, -0.90 DPS, sim-verified) [crafted] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Sergeant Major's Cape (16336, -0.04 DPS) [pvp]; Sergeant Major's Cape (16315, -0.27 DPS) [pvp] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 20.1 | yes | Golden Scale Cuirass (3845, -0.15 DPS) [crafted]; Shining Mithril Breastplate (250540, -0.15 DPS) [crafted]; Shining Silver Breastplate (2870, -0.67 DPS, sim-verified) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Pugilist Bracers (4438, -0.08 DPS, sim-verified) [dungeon]; Ravager's Armguards (14770, -0.73 DPS) [world]; Cultist's Armguards (270032, -0.74 DPS) [quest] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 80.6 | yes | Dragonscale Gauntlets (8347, -1.22 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -1.49 DPS) [crafted]; Ornate Mithril Gloves (7927, -1.49 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20153) | The Defilers [rep] | 68.6 | yes | Defiler's Leather Girdle (20192, -0.38 DPS, sim-verified) [rep]; Defiler's Chain Girdle (20152, -3.31 DPS) [rep]; Defiler's Leather Girdle (20191, -3.31 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 27.1 | yes | Orcish War Leggings (7929, -0.38 DPS) [crafted]; Veteran's Silvered Chain Leggings (250523, -0.81 DPS) [crafted]; Ferine Leggings (6690, -1.31 DPS, sim-verified) [dungeon] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 18.8 | yes | Prowler's Leather Shoes (252465, -0.08 DPS, sim-verified) [crafted]; Blackforge Greaves (6423, -0.32 DPS) [dungeon]; Skirmisher's Mail Boots (252564, -0.34 DPS) [crafted] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 12.6 | yes | Legionnaire's Band (19513, -0.23 DPS) [rep]; Suspicious Spare Part (274754, -0.27 DPS) [vendor]; Ironspine's Eye (7686, -0.36 DPS) [dungeon] |
| finger2 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 12.0 | yes | Insurgent's Band (272067, -0.22 DPS) [vendor]; Suspicious Spare Part (274754, -0.27 DPS, sim-verified) [vendor]; Ironspine's Eye (7686, -0.31 DPS) [dungeon] |
| trinket1 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [dungeon]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [dungeon]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Ardent Custodian (868) | Gnomeregan: Dark Iron Ambassador [dungeon] | 460.0 | yes | Fiery War Axe (870, +0.00 DPS) [dungeon]; Staff of Jordan (873, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 442.0 | yes | Pit Fighter's Shield (4507, -32.06 DPS) [quest]; Combat Shield (4065, -32.16 DPS) [dungeon]; Aegis of the Scarlet Commander (7726, -32.16 DPS) [dungeon] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 | yes | Mithril Blacksmith Hammer (285280, -0.19 DPS) [crafted]; Master Hunter's Rifle (17687, -0.20 DPS) [quest]; Explosive Shotgun (8188, -0.35 DPS, sim-verified) [world] |

**New at 40:** head: Raging Berserker's Helm; shoulder: Sunburn Spaulders; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Chain Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Legionnaire's Band; finger2: Insurgent's Band; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Ardent Custodian; off_hand: Vanquisher's Sword

No-known-source sample (15 of 1445, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 50 (troll, 35100000000000000-353211005050010051-000000000000000000)

Set DPS (verified): 98.3. Weights run: 1.0s. Verify run: 1.3s. 1943 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.217, strength=1.796 ± 0.296, agility=not significant (0.302 ± 0.101), crit=4.820 ± 0.391, hit=2.018 ± 0.331, melee_haste=2.822 ± 0.403

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Plate Helm (220803) (or Knight-Lieutenant's Plate Helm (220804)) | Lady Palanseer [vendor] | 111.0 | yes | Knight-Lieutenant's Plate Helm (220804, +0.00 DPS, sim-verified) [vendor]; Raging Berserker's Helm (7719, -1.94 DPS) [dungeon]; Ornate Mithril Helm (7937, -2.46 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Woven Ivy Necklace (19159, +0.24 DPS, sim-verified) [quest]; Ethereal Talisman (4430, -0.37 DPS) [quest]; Scout's Medallion (19535, -1.00 DPS) [rep] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) (or Blood Guard's Plate Pauldrons (220796)) | Captain Dirgehammer [vendor] | 87.2 | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS, sim-verified) [vendor]; Officer's Pauldrons (250576, -5.91 DPS) [crafted]; Wyrmslayer Spaulders (13066, -6.09 DPS) [world] |
| back | Bloodlust Cape (14801) | Maraudon: Princess Theradras [dungeon] | 16.2 | yes | Sergeant Major's Cape (16336, -0.51 DPS, sim-verified) [pvp]; Wolfmaster Cape (6314, -0.59 DPS) [dungeon]; Battlehard Cape (11858, -0.59 DPS) [quest] |
| chest | Knight's Plate Hauberk (220794) (or Stone Guard's Plate Armor (220801)) | Captain Dirgehammer [vendor] | 94.4 | yes | Stone Guard's Plate Armor (220801, +0.00 DPS, sim-verified) [vendor]; Ornate Mithril Breastplate (7935, -2.59 DPS) [crafted]; Warforged Chestplate (11195, -4.94 DPS) [quest] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 | yes | Officer's Wristguards (250581, -0.51 DPS, sim-verified) [crafted]; Branded Leather Bracers (19508, -0.77 DPS) [dungeon]; Prowler's Leather Bracers (252539, -0.94 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 87.5 | yes | Dragonscale Gauntlets (8347, -1.03 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -1.92 DPS) [crafted]; Ornate Mithril Gloves (7927, -1.92 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20151) (or Defiler's Leather Girdle (20193)) | The Defilers [rep] | 87.5 | yes | Defiler's Leather Girdle (20193, +0.00 DPS, sim-verified) [rep]; Defiler's Plate Girdle (20205, -0.20 DPS) [rep]; Defiler's Chain Girdle (20153, -1.15 DPS) [rep] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | 92.1 | yes | Stone Guard's Plate Leggings (220798, +0.00 DPS) [vendor]; Stormshroud Pants (15057, -1.28 DPS, sim-verified) [crafted]; Scarlet Leggings (10330, -5.23 DPS) [dungeon] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 26.7 | yes | Officer's Boots (250546, -0.12 DPS) [crafted]; Officer's Sabatons (250561, -0.18 DPS, sim-verified) [crafted]; Skulker's Leather Boots (252469, -0.29 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 40.2 | yes | Legionnaire's Band (19511, -1.88 DPS) [rep]; Legionnaire's Band (19512, -2.25 DPS) [rep]; Band of Allegiance (18585, -2.31 DPS) [quest] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 | yes | Legionnaire's Band (19511, -0.05 DPS, sim-verified) [rep]; Legionnaire's Band (19512, -0.69 DPS) [rep]; Band of Allegiance (18585, -0.75 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 56.1 | yes | Tidal Charm (1404, -5.40 DPS) [vendor]; Guardian Talisman (1490, -5.40 DPS) [quest]; Ankh of Life (1713, -5.40 DPS) [dungeon] |
| trinket2 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [dungeon] |
| main_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | 577.5 | yes | Kindling Stave (11750, +0.00 DPS) [dungeon]; Bleakwood Hew (12769, +0.00 DPS) [crafted]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor] |
| off_hand | Hammer of the Northern Wind (810) | Blackrock Depths: Anvilrage Medic [dungeon] | 553.3 | yes | Claw of Celebras (17738, -6.84 DPS) [dungeon]; White Bone Shredder (11863, -10.49 DPS) [quest]; Shizzle's Drizzle Blocker (11915, -50.98 DPS) [quest] |
| ranged | Houndmaster's Bow (11628) | Blackrock Depths: Houndmaster Grebmar [dungeon] | 17.4 | yes | Arcanite Blacksmith Hammer (285281, -0.64 DPS) [crafted]; Booty Bay Bruiser's Buckshot (274748, -0.81 DPS) [vendor]; Explosive Shotgun (8188, -0.81 DPS) [world] |

**New at 50:** head: Blood Guard's Plate Helm; shoulder: Knight-Lieutenant's Plate Pauldrons; back: Bloodlust Cape; chest: Knight's Plate Hauberk; wrist: Bracers of the Stone Princess; waist: Defiler's Chain Girdle; legs: Knight's Plate Leggings; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Smoking Heart of the Mountain; main_hand: Dawn's Edge; off_hand: Hammer of the Northern Wind; ranged: Houndmaster's Bow

No-known-source sample (15 of 1943, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (troll, 35311103002000000-353211005050010051-000000000000000000)

Set DPS (verified): 248.1. Weights run: 1.0s. Verify run: 1.4s. 2876 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=not significant (1.000 ± 0.427), strength=not significant (2.344 ± 0.596), agility=not significant (0.450 ± 0.239), crit=13.060 ± 0.873, hit=not significant (0.000 ± 0.000), melee_haste=6.782 ± 0.884

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 407.9 | yes | Bloodvine Lens (19998, -2.44 DPS) [crafted]; Ragefury Eyepatch (11735, -9.04 DPS, sim-verified) [dungeon]; Field Marshal's Plate Helm (16478, -9.20 DPS) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 206.8 | yes | Onyxia Tooth Pendant (18404, -1.20 DPS) [quest]; Choker of the Shifting Sands (21505, -9.52 DPS) [quest]; Pendant of the Shifting Sands (21506, -10.05 DPS) [quest] |
| shoulder | Champion's Plate Shoulders (23243) (or Lieutenant Commander's Plate Shoulders (23315), Champion's Plate Shoulders (227042), Lieutenant Commander's Plate Shoulders (227045)) | Lady Palanseer [vendor] | 222.7 | yes | Lieutenant Commander's Plate Shoulders (23315, +0.00 DPS, sim-verified) [vendor]; Champion's Plate Shoulders (227042, +0.00 DPS) [pvp]; Lieutenant Commander's Plate Shoulders (227045, +0.00 DPS) [pvp] |
| back | Drape of Unyielding Strength (21394) | Drape of Unyielding Strength [quest] | 39.2 | yes | Cloak of the Fallen God (21710, -0.10 DPS) [quest]; Deathguard's Cloak (20068, -0.17 DPS) [rep]; Chromatic Cloak (18509, -8.64 DPS, sim-verified) [crafted] |
| chest | Bloodsoul Breastplate (19690) | Blacksmithing [crafted] | 369.7 | yes | Stormshroud Armor (15056, -1.34 DPS, sim-verified) [crafted]; Savage Gladiator Chain (11726, -2.98 DPS) [dungeon]; Obsidian Mail Tunic (22191, -6.40 DPS) [crafted] |
| wrist | Vambraces of the Sadist (13400) | Stratholme: Timmy the Cruel [dungeon] | 199.2 | yes | Deeprock Bracers (21184, +2.13 DPS, sim-verified) [quest]; Berserker Bracers (19578, -8.72 DPS) [rep]; Marshal's Plate Bracers (16481, -9.02 DPS) [pvp] |
| hands | Marshal's Plate Gauntlets (16484) | Captain Dirgehammer [vendor] | 229.7 | yes | General's Plate Gauntlets (16548, +0.00 DPS, sim-verified) [vendor]; General's Plate Gauntlets (231532, -0.00 DPS) [pvp]; Marshal's Plate Gauntlets (231541, -0.00 DPS) [pvp] |
| waist | Zandalar Vindicator's Belt (19823) | Paragons of Power: The Vindicator's Belt [quest] | 241.4 | yes | Defiler's Plate Girdle (20204, -1.34 DPS, sim-verified) [rep]; Defiler's Chain Girdle (20150, -1.42 DPS) [rep]; Defiler's Leather Girdle (20190, -1.42 DPS) [rep] |
| legs | Marshal's Plate Legguards (16479) (or General's Plate Leggings (16543), General's Plate Leggings (231533), Marshal's Plate Legguards (231540)) | Captain Dirgehammer [vendor] | 412.6 | yes | General's Plate Leggings (16543, +0.00 DPS, sim-verified) [vendor]; General's Plate Leggings (231533, +0.00 DPS) [pvp]; Marshal's Plate Legguards (231540, +0.00 DPS) [pvp] |
| feet | General's Plate Boots (231531) | Rank 16 [pvp] | 47.6 | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Marshal's Plate Boots (16483, -0.00 DPS) [vendor]; Conqueror's Greaves (21333, -3.24 DPS, sim-verified) [quest] |
| finger1 | Signet of Unyielding Strength (21393) | Signet of Unyielding Strength [quest] | 208.6 | yes | Band of Earthen Might (21182, -0.68 DPS) [quest]; Band of the Penitent (13217, -1.49 DPS) [quest]; Dragonslayer's Signet (18403, -1.49 DPS) [quest] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 198.8 | yes | Band of Earthen Might (21182, -0.48 DPS, sim-verified) [quest]; Band of the Penitent (13217, -0.92 DPS) [quest]; Dragonslayer's Signet (18403, -0.92 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 42.0 | yes | Tidal Charm (1404, -2.42 DPS) [vendor]; Guardian Talisman (1490, -2.42 DPS) [quest]; Ankh of Life (1713, -2.42 DPS) [dungeon] |
| trinket2 | Onyxia Blood Talisman (18406) | For All To See [quest] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [dungeon] |
| main_hand | High Warlord's Quickblade (234553) | Rank 18 [pvp] | 1078.0 | yes | High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Pulverizer (234545, +0.00 DPS) [pvp] |
| off_hand | Grand Marshal's Swiftblade (234579) | Rank 18 [pvp] | 1078.0 | yes | High Warlord's Left Claw (234558, -0.18 DPS) [pvp]; Grand Marshal's Left Hand Blade (234584, -0.18 DPS) [pvp]; Grand Marshal's Left Hand Blade (18847, -1.99 DPS) [vendor] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | 182.8 | yes | Bloodseeker (19107, -9.29 DPS) [quest]; Houndmaster's Bow (11628, -9.46 DPS) [dungeon]; Arcanite Blacksmith Hammer (285281, -9.74 DPS) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Medallion of the Dawn; shoulder: Champion's Plate Shoulders; back: Drape of Unyielding Strength; chest: Bloodsoul Breastplate; wrist: Vambraces of the Sadist; hands: Marshal's Plate Gauntlets; waist: Zandalar Vindicator's Belt; legs: Marshal's Plate Legguards; feet: General's Plate Boots; finger1: Signet of Unyielding Strength; finger2: Don Julio's Band; trinket2: Onyxia Blood Talisman; main_hand: High Warlord's Quickblade; off_hand: Grand Marshal's Swiftblade; ranged: The Purifier

No-known-source sample (15 of 2876, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

