# Leveling BiS: Fury

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

## Alliance

### Band 20 (human, 00000000000000000-353000000000000000-000000000000000000)

Set DPS (verified): 32.3. Weights run: 1.6s. Verify run: 1.6s. 537 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.034, strength=2.084 ± 0.042, agility=0.085 ± 0.015, crit=2.317 ± 0.066, hit=1.130 ± 0.086, melee_haste=1.686 ± 0.042

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.8 | yes | Defender's Leather Hood (252447, -4.2) [crafted]; Guard's Silvered Chain Helm (250529, -20.0) [crafted]; Brawler's Leather Hood (252504, -20.2) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 0.5 | yes | Tarnished Locket (279870, -0.5) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.3 | yes | Silvered Bronze Shoulders (3481, +0.0) [crafted]; Serpent's Shoulders (5404, -5.8) [dungeon]; Double-Stitched Woolen Shoulders (4314, -6.3) [crafted] |
| back | Lambent Scale Cloak (4706) | Gnomeregan: Caverndeep Burrower [dungeon] | 8.3 | yes | Grave Shroud (279865, -1.9) [quest]; Catacomb Cloak (279899, -2.3) [quest]; Dark Leather Cloak (2316, -3.9) [crafted] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.8 | yes | Veteran's Chain Shirt (250488, -5.9) [crafted]; Defender's Leather Armor (252434, -6.0) [crafted]; Totemic Leather Armor (252435, -6.3) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.4 | yes | Cryptwalker Bracers (280095, -2.1) [quest]; Bravo's Armbands (270015, -5.9) [quest]; Runed Copper Bracers (2854, -6.3) [crafted] |
| hands | Gold-flecked Gloves (5195) | The Deadmines: Sneed [dungeon] | 14.6 | yes | Fletcher's Gloves (7348, +17.8) [crafted]; Polar Gauntlets (7606, -2.1) [quest]; Blackened Defias Gloves (10401, -2.1) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 | yes | Cobrahn's Grasp (6460, -3.2) [dungeon]; Ruffian Belt (5975, -5.5) [world]; Support Girdle (1215, -7.6) [world] |
| legs | Chausses of Westfall (6087) | Quests [quest] | 22.9 | yes | Veteran's Chain Leggings (250493, -3.7) [crafted]; Defender's Leather Pants (252445, -3.8) [crafted]; Totemic Leather Pants (252446, -4.2) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.8 | yes | Veteran's Boots (250503, -0.1) [crafted]; Guard's Boots (250504, -0.4) [crafted]; Defender's Leather Boots (252441, -0.4) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.7 | yes | Signet of the Zhevra (285330, -8.2) [world]; Lavishly Jeweled Ring (1156, -8.5) [dungeon]; Minor Channeling Ring (1449, -8.7) [quest] |
| finger2 | The 1 Ring (8350) | Fishing [world] | 2.2 | yes | Signet of the Zhevra (285330, -1.7) [world]; Lavishly Jeweled Ring (1156, -2.0) [dungeon]; Minor Channeling Ring (1449, -2.2) [quest] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 | yes | Hammerbone (270018, +58.1) [quest]; Forsaken Greataxe (251533, +54.8) [quest]; Smite's Mighty Hammer (7230, +50.3) [dungeon] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 237.1 | yes | Bear Buckler (4821, -230.9) [vendor]; Furen's Favor (6970, -230.9) [quest]; Veteran Shield (3651, -232.9) [dungeon] |
| ranged | Dwarven Fishing Pole (3567) (or Cracked Blacksmith Hammer (285279)) | Quests [quest] | 4.2 | yes | Cracked Blacksmith Hammer (285279, +0.0) [crafted]; Fine Longbow (11304, -0.2) [vendor]; Daryl's Hunting Rifle (2904, -2.1) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Gold-flecked Gloves; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: The 1 Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Dwarven Fishing Pole

No-known-source sample (15 of 537, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 7956 Bronze Warhammer; 8178 Training Sword; 9602 Brushwood Blade; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse

### Band 30 (human, 00000000000000000-353211005010000000-000000000000000000)

Set DPS (verified): 53.1. Weights run: 1.6s. Verify run: 1.8s. 1006 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.160, strength=1.676 ± 0.208, agility=not significant (0.186 ± 0.061), crit=3.475 ± 0.238, hit=1.401 ± 0.221, melee_haste=1.775 ± 0.290

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 21.8 | yes | Veteran's Chain Helm (250498, -1.7) [crafted]; Defender's Leather Helm (252455, -1.7) [crafted]; Crusader's Chain Helm (250502, -3.4) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19541, -12.5) [rep]; Sentinel's Medallion (20444, -12.9) [rep]; Pendant of Myzrael (4614, -14.0) [dungeon] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 11.7 | yes | Mail Combat Spaulders (6404, +0.0) [dungeon]; Barbaric Iron Shoulders (7913, -0.6) [crafted]; Forest Tracker Epaulets (2278, -1.3) [dungeon] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Sergeant Major's Cape (16315, -2.6) [pvp]; Lambent Scale Cloak (4706, -3.3) [dungeon]; Catacomb Cloak (279899, -4.0) [quest] |
| chest | Shining Silver Breastplate (2870) | Blacksmithing [crafted] | 23.5 | yes | Barbaric Iron Breastplate (7914, -3.4) [crafted]; Hard Gold Cuirass (250533, -5.0) [crafted]; Veteran's Silvered Chain Shirt (250518, -5.6) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 13.4 | yes | Bands of Serra'kis (6902, -3.4) [dungeon]; Cultist's Armguards (270032, -3.4) [quest]; Patterned Bronze Bracers (2868, -5.0) [crafted] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 21.0 | yes | Fletcher's Gloves (7348, +27.6) [crafted]; Heavy Earthen Gloves (7359, -5.0) [crafted]; Bonefist Gauntlets (4465, -5.9) [world] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 | yes | Highlander's Leather Girdle (20117, +0.0) [rep]; Highlander's Plate Girdle (20126, -3.9) [rep]; Highlander's Lamellar Girdle (20108, -5.6) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Veteran's Silvered Chain Leggings (250523, -6.3) [crafted]; Golden Scale Leggings (3843, -7.6) [crafted]; Chausses of Westfall (6087, -7.6) [quest] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 13.0 | yes | Hard Gold Boots (250534, -1.3) [crafted]; Glimmering Mail Greaves (4073, -3.0) [dungeon]; Brawler's Leather Boots (252439, -3.7) [crafted] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 11.2 | yes | Insurgent's Band (272067, -2.2) [vendor]; Silverlaine's Family Seal (6321, -2.8) [dungeon]; Protector's Band (20439, -3.7) [rep] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 8.4 | yes | Insurgent's Band (272067, +0.6) [vendor]; Silverlaine's Family Seal (6321, -0.0) [dungeon]; Seal of Wrynn (2933, -2.8) [quest] |
| trinket1 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.0) [rep]; Rune of Duty (21568, +0.0) [rep]; Relentless Raider's Seal (272062, +0.0) [vendor] |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.0) [rep]; Rune of Duty (21568, +0.0) [rep]; Relentless Raider's Seal (272062, +0.0) [vendor] |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 343.8 | yes | Morbid Dawn (7689, +95.3) [dungeon]; Manual Crowd Pummeler (9449, +89.9) [dungeon]; Corpsemaker (6687, +86.6) [dungeon] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 332.6 | yes | Shoni's Disarming Tool (9608, -107.9) [quest]; Commander's Crest (6320, -324.2) [dungeon]; Lambent Scale Shield (3656, -325.9) [dungeon] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 | yes | Double-barreled Shotgun (2098, -3.4) [dungeon]; Moonsight Rifle (4383, -4.8) [crafted]; Precision Bow (217315, -4.8) [quest] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wolfmaster Cape; chest: Shining Silver Breastplate; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Protector's Band; finger2: Ironspine's Eye; trinket1: Talisman of Arathor; trinket2: Darkspear Voodoo Seal; main_hand: Pronged Reaver; off_hand: Ironspine's Fist; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 1006, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 7956 Bronze Warhammer

### Band 40 (human, 00000000000000000-353211005050010050-000000000000000000)

Set DPS (verified): 88.1. Weights run: 1.9s. Verify run: 2.0s. 1451 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.203, strength=1.289 ± 0.264, agility=not significant (0.291 ± 0.084), crit=4.325 ± 0.374, hit=1.889 ± 0.312, melee_haste=2.652 ± 0.374

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 77.3 | yes | Icemetal Barbute (10763, -59.3) [dungeon]; Hard Gold Coif (250537, -59.3) [crafted]; White Bandit Mask (10008, -59.9) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19540, -10.8) [rep]; Sentinel's Medallion (19541, -11.7) [rep]; Sentinel's Medallion (20444, -12.3) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 15.2 | yes | Hard Gold Pauldrons (250539, -1.0) [crafted]; Shining Mithril Pauldrons (250541, -2.3) [crafted]; Imperial Leather Spaulders (4737, -3.6) [dungeon] |
| back | Sergeant Major's Cape (16336) | Rank 9 [pvp] | 9.5 | yes | Wolfmaster Cape (6314, +0.5) [dungeon]; Sergeant Major's Cape (16315, -3.2) [pvp]; Catacomb Cloak (279899, -3.5) [quest] |
| chest | Kolkar Marauder Chain (6773) | Quests [quest] | 20.1 | yes | Shining Silver Breastplate (2870, -2.0) [crafted]; Golden Scale Cuirass (3845, -2.0) [crafted]; Shining Mithril Breastplate (250540, -2.0) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Pugilist Bracers (4438, -9.7) [dungeon]; Ravager's Armguards (14770, -9.8) [world]; Cultist's Armguards (270032, -10.0) [quest] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 80.6 | yes | Dragonscale Gauntlets (8347, -18.3) [crafted]; Fletcher's Gloves (7348, -20.0) [crafted]; Ornate Mithril Gloves (7927, -20.0) [crafted] |
| waist | Highlander's Chain Girdle (20089) | The League of Arathor [rep] | 68.6 | yes | Highlander's Leather Girdle (20116, -38.6) [rep]; Highlander's Chain Girdle (20090, -44.6) [rep]; Highlander's Leather Girdle (20117, -44.6) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 27.1 | yes | Ferine Leggings (6690, -1.1) [dungeon]; Orcish War Leggings (7929, -5.2) [crafted]; Veteran's Silvered Chain Leggings (250523, -10.9) [crafted] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 18.8 | yes | Prowler's Leather Shoes (252465, -2.6) [crafted]; Blackforge Greaves (6423, -4.3) [dungeon]; Skirmisher's Mail Boots (252564, -4.6) [crafted] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 12.6 | yes | Protector's Band (19517, -3.2) [rep]; Suspicious Spare Part (274754, -3.6) [vendor]; Ironspine's Eye (7686, -4.9) [dungeon] |
| finger2 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 12.0 | yes | Suspicious Spare Part (274754, -3.0) [vendor]; Insurgent's Band (272067, -3.0) [vendor]; Ironspine's Eye (7686, -4.2) [dungeon] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Tidal Charm (1404, +0.0) [vendor]; Blazing Emblem (2802, +0.0) [dungeon]; Cold Basilisk Eye (5079, +0.0) [world] |
| trinket2 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | 0.0 | yes | Tidal Charm (1404, +0.0) [vendor]; Blazing Emblem (2802, +0.0) [dungeon]; Cold Basilisk Eye (5079, +0.0) [world] |
| main_hand | Ardent Custodian (868) | Gnomeregan: Dark Iron Ambassador [dungeon] | 460.0 | yes | Bonebiter (6830, +109.2) [quest]; Staff of Jordan (873, +105.7) [dungeon]; Primitive Fishing Pole (276203, +105.6) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Quests [quest] | 442.0 | yes | Shoni's Disarming Tool (9608, -217.3) [quest]; Salbac Shield (4652, -430.4) [quest]; Combat Shield (4065, -433.0) [dungeon] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 | yes | Explosive Shotgun (8188, -2.6) [world]; Mithril Blacksmith Hammer (285280, -2.6) [crafted]; Master Hunter's Rifle (17687, -2.7) [quest] |

**New at 40:** head: Raging Berserker's Helm; shoulder: Sunburn Spaulders; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Chain Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Protector's Band; finger2: Insurgent's Band; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Ardent Custodian; off_hand: Vanquisher's Sword

No-known-source sample (15 of 1451, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 50 (human, 35100000000000000-353211005050010051-000000000000000000)

Set DPS (verified): 92.5. Weights run: 2.0s. Verify run: 1.9s. 1949 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.217, strength=1.796 ± 0.296, agility=not significant (0.302 ± 0.101), crit=4.820 ± 0.391, hit=2.018 ± 0.331, melee_haste=2.822 ± 0.403

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Plate Helm (220803) | Lady Palanseer [vendor] | 111.0 | yes | Knight-Lieutenant's Plate Helm (220804, -0.0) [vendor]; Raging Berserker's Helm (7719, -20.2) [dungeon]; Ornate Mithril Helm (7937, -25.6) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19539, -10.4) [rep]; Talisman of the Naga Lord (5029, -10.4) [world]; Sentinel's Medallion (19540, -10.7) [rep] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) (or Blood Guard's Plate Pauldrons (220796)) | Captain Dirgehammer [vendor] | 87.2 | yes | Blood Guard's Plate Pauldrons (220796, +0.0) [vendor]; Officer's Pauldrons (250576, -61.5) [crafted]; Wyrmslayer Spaulders (13066, -63.3) [world] |
| back | Bloodlust Cape (14801) | Maraudon: Princess Theradras [dungeon] | 16.2 | yes | Sergeant Major's Cape (16336, -3.6) [pvp]; Wolfmaster Cape (6314, -6.2) [dungeon]; Sergeant Major's Cape (16315, -7.8) [pvp] |
| chest | Knight's Plate Hauberk (220794) (or Stone Guard's Plate Armor (220801)) | Captain Dirgehammer [vendor] | 94.4 | yes | Stone Guard's Plate Armor (220801, +0.0) [vendor]; Ornate Mithril Breastplate (7935, -26.9) [crafted]; Warforged Chestplate (11195, -51.3) [quest] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 | yes | Officer's Wristguards (250581, -6.4) [crafted]; Branded Leather Bracers (19508, -8.0) [dungeon]; Prowler's Leather Bracers (252539, -9.7) [crafted] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 87.5 | yes | Dragonscale Gauntlets (8347, -18.2) [crafted]; Fletcher's Gloves (7348, -20.0) [crafted]; Ornate Mithril Gloves (7927, -20.0) [crafted] |
| waist | Highlander's Chain Girdle (20088) (or Highlander's Leather Girdle (20115)) | The League of Arathor [rep] | 87.5 | yes | Highlander's Leather Girdle (20115, +0.0) [rep]; Highlander's Lamellar Girdle (20106, -0.2) [rep]; Highlander's Plate Girdle (20124, -2.0) [rep] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | 92.1 | yes | Stormshroud Pants (15057, +42.9) [crafted]; Stone Guard's Plate Leggings (220798, +0.0) [vendor]; Scarlet Leggings (10330, -54.3) [dungeon] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 26.7 | yes | Officer's Sabatons (250561, -0.6) [crafted]; Officer's Boots (250546, -1.2) [crafted]; Skulker's Leather Boots (252469, -3.0) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 40.2 | yes | Insurgent's Band (272065, -25.2) [vendor]; Suspicious Spare Part (274754, -27.6) [vendor]; Insurgent's Band (272066, -28.2) [vendor] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 20.7 | yes | Protector's Band (19515, -3.9) [rep]; Insurgent's Band (272065, -5.7) [vendor]; Protector's Band (19517, -8.1) [rep] |
| trinket1 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.0) [quest]; Tidal Charm (1404, +0.0) [vendor]; Guardian Talisman (1490, +0.0) [quest] |
| trinket2 | Talisman of Arathor (21117) | The League of Arathor [rep] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.0) [quest]; Tidal Charm (1404, +0.0) [vendor]; Guardian Talisman (1490, +0.0) [quest] |
| main_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | 577.5 | yes | Kindling Stave (11750, +132.0) [dungeon]; Darkspear Raider's Reaper (272080, +112.3) [vendor]; Bleakwood Hew (12769, +93.0) [crafted] |
| off_hand | Hammer of the Northern Wind (810) | Blackrock Depths: Anvilrage Medic [dungeon] | 553.3 | yes | Claw of Celebras (17738, -71.1) [dungeon]; Shoni's Disarming Tool (9608, -328.6) [quest]; Shizzle's Drizzle Blocker (11915, -529.9) [quest] |
| ranged | Houndmaster's Bow (11628) | Blackrock Depths: Houndmaster Grebmar [dungeon] | 17.4 | yes | Arcanite Blacksmith Hammer (285281, -6.6) [crafted]; Booty Bay Bruiser's Buckshot (274748, -8.4) [vendor]; Explosive Shotgun (8188, -8.4) [world] |

**New at 50:** head: Blood Guard's Plate Helm; shoulder: Knight-Lieutenant's Plate Pauldrons; back: Bloodlust Cape; chest: Knight's Plate Hauberk; wrist: Bracers of the Stone Princess; waist: Highlander's Chain Girdle; legs: Knight's Plate Leggings; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Smoking Heart of the Mountain; trinket2: Talisman of Arathor; main_hand: Dawn's Edge; off_hand: Hammer of the Northern Wind; ranged: Houndmaster's Bow

No-known-source sample (15 of 1949, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (human, 35311103002000000-353211005050010051-000000000000000000)

Set DPS (verified): 380.4. Weights run: 2.1s. Verify run: 2.4s. 2761 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=not significant (1.000 ± 0.427), strength=not significant (2.344 ± 0.596), agility=not significant (0.450 ± 0.239), crit=13.060 ± 0.873, hit=not significant (0.000 ± 0.000), melee_haste=6.782 ± 0.884

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 407.9 | yes | Ragefury Eyepatch (11735, -28.1) [dungeon]; Bloodvine Lens (19998, -42.2) [crafted]; Field Marshal's Plate Helm (16478, -159.4) [vendor] |
| neck | Stormrage's Talisman of Seething (23053) | Naxxramas [raid] | 391.7 | yes | Gem of Trapped Innocents (23057, -26.0) [raid]; Barbed Choker (21664, -164.8) [raid]; Medallion of the Dawn (22659, -184.8) [quest] |
| shoulder | Champion's Plate Shoulders (23243) (or Lieutenant Commander's Plate Shoulders (23315), Champion's Plate Shoulders (227042), Lieutenant Commander's Plate Shoulders (227045)) | Lady Palanseer [vendor] | 222.7 | yes | Lieutenant Commander's Plate Shoulders (23315, +0.0) [vendor]; Champion's Plate Shoulders (227042, +0.0) [pvp]; Lieutenant Commander's Plate Shoulders (227045, +0.0) [pvp] |
| back | Drape of Unyielding Strength (21394) | Quests [quest] | 39.2 | yes | Chromatic Cloak (18509, +143.6) [crafted]; Cloak of the Fallen God (21710, -1.7) [quest]; Cloak of the Honor Guard (20073, -3.0) [rep] |
| chest | Bloodsoul Breastplate (19690) | Blacksmithing [crafted] | 369.7 | yes | Stormshroud Armor (15056, -4.1) [crafted]; Savage Gladiator Chain (11726, -51.7) [dungeon]; Obsidian Mail Tunic (22191, -110.9) [crafted] |
| wrist | Hive Defiler Wristguards (21618) | Ahn'Qiraj [raid] | 62.0 | yes | Vambraces of the Sadist (13400, +137.2) [dungeon]; Deeprock Bracers (21184, -13.0) [quest]; Berserker Bracers (19578, -13.9) [rep] |
| hands | Gauntlets of Annihilation (21581) | Ahn'Qiraj [raid] | 264.9 | yes | Marshal's Plate Gauntlets (16484, -35.2) [vendor]; General's Plate Gauntlets (16548, -35.2) [vendor]; General's Plate Gauntlets (231532, -35.2) [pvp] |
| waist | Belt of Never-ending Agony (21586) | Ahn'Qiraj [raid] | 246.8 | yes | Zandalar Vindicator's Belt (19823, -5.4) [quest]; Highlander's Plate Girdle (20041, -24.2) [rep]; Highlander's Lamellar Girdle (20042, -28.8) [rep] |
| legs | Marshal's Plate Legguards (16479) (or General's Plate Leggings (16543), General's Plate Leggings (231533), Marshal's Plate Legguards (231540)) | Captain Dirgehammer [vendor] | 412.6 | yes | General's Plate Leggings (16543, +0.0) [vendor]; General's Plate Leggings (231533, +0.0) [pvp]; Marshal's Plate Legguards (231540, +0.0) [pvp] |
| feet | Conqueror's Greaves (21333) | Quests [quest] | 56.9 | yes | Marshal's Plate Boots (16483, -9.3) [vendor]; General's Plate Boots (16545, -9.3) [vendor]; General's Plate Boots (231531, -9.3) [pvp] |
| finger1 | Band of Unnatural Forces (23038) | Naxxramas [raid] | 234.8 | yes | Ring of the Qiraji Fury (21677, -12.0) [raid]; Signet of Unyielding Strength (21393, -26.2) [quest]; Don Julio's Band (19325, -36.0) [rep] |
| finger2 | Quick Strike Ring (18821) | Molten Core: Magmadar [raid] | 224.6 | yes | Ring of the Qiraji Fury (21677, -1.7) [raid]; Signet of Unyielding Strength (21393, -15.9) [quest]; Don Julio's Band (19325, -25.7) [rep] |
| trinket1 | Slayer's Crest (23041) | Naxxramas [raid] | 64.0 | yes | Eye of Diminution (23001, +301.7) [raid]; Drake Fang Talisman (19406, -8.0) [raid]; Thunderbrew's Boot Flask (744, -64.0) [quest] |
| trinket2 | Kiss of the Spider (22954) | Naxxramas [raid] | 182.8 | yes | Eye of Diminution (23001, +182.8) [raid]; Drake Fang Talisman (19406, -126.8) [raid]; Thunderbrew's Boot Flask (744, -182.8) [quest] |
| main_hand | Thunderfury, Blessed Blade of the Windseeker (19019) | Quests [quest] | 757.6 | yes | High Warlord's Greatsword (234542, +612.2) [pvp]; High Warlord's Battle Axe (234543, +612.2) [pvp]; High Warlord's Pulverizer (234545, +612.2) [pvp] |
| off_hand | Grand Marshal's Swiftblade (234579) | Rank 18 [pvp] | 1078.0 | yes | High Warlord's Left Claw (234558, -3.1) [pvp]; Grand Marshal's Left Hand Blade (234584, -3.1) [pvp]; Grand Marshal's Left Hand Blade (18847, -34.4) [vendor] |
| ranged | Larvae of the Great Worm (23557) | Ahn'Qiraj [raid] | 200.8 | yes | The Purifier (22656, -18.0) [quest]; Bloodseeker (19107, -178.9) [quest]; Houndmaster's Bow (11628, -181.8) [dungeon] |

**New at 60:** head: Lionheart Helm; neck: Stormrage's Talisman of Seething; shoulder: Champion's Plate Shoulders; back: Drape of Unyielding Strength; chest: Bloodsoul Breastplate; wrist: Hive Defiler Wristguards; hands: Gauntlets of Annihilation; waist: Belt of Never-ending Agony; legs: Marshal's Plate Legguards; feet: Conqueror's Greaves; finger1: Band of Unnatural Forces; finger2: Quick Strike Ring; trinket1: Slayer's Crest; trinket2: Kiss of the Spider; main_hand: Thunderfury, Blessed Blade of the Windseeker; off_hand: Grand Marshal's Swiftblade; ranged: Larvae of the Great Worm

No-known-source sample (15 of 2761, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

## Horde

### Band 20 (troll, 00000000000000000-353000000000000000-000000000000000000)

Set DPS (verified): 29.0. Weights run: 1.6s. Verify run: 1.5s. 531 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.034, strength=2.084 ± 0.042, agility=0.085 ± 0.015, crit=2.317 ± 0.066, hit=1.130 ± 0.086, melee_haste=1.686 ± 0.042

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.8 | yes | Defender's Leather Hood (252447, -4.2) [crafted]; Guard's Silvered Chain Helm (250529, -20.0) [crafted]; Brawler's Leather Hood (252504, -20.2) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 0.5 | yes | Tarnished Locket (279870, -0.5) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.3 | yes | Silvered Bronze Shoulders (3481, +0.0) [crafted]; Serpent's Shoulders (5404, -5.8) [dungeon]; Double-Stitched Woolen Shoulders (4314, -6.3) [crafted] |
| back | Lambent Scale Cloak (4706) | Gnomeregan: Caverndeep Burrower [dungeon] | 8.3 | yes | Grave Shroud (279865, -1.9) [quest]; Subterranean Cape (14149, -2.1) [dungeon]; Catacomb Cloak (279899, -2.3) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.8 | yes | Veteran's Chain Shirt (250488, -5.9) [crafted]; Defender's Leather Armor (252434, -6.0) [crafted]; Totemic Leather Armor (252435, -6.3) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.4 | yes | Cryptwalker Bracers (280095, -2.1) [quest]; Bravo's Armbands (270015, -5.9) [quest]; Runed Copper Bracers (2854, -6.3) [crafted] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 32.4 | yes | Gold-flecked Gloves (5195, -17.8) [dungeon]; Blackened Defias Gloves (10401, -19.9) [dungeon]; Dagmire Gauntlets (6481, -21.8) [quest] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 | yes | Cobrahn's Grasp (6460, -3.2) [dungeon]; Ruffian Belt (5975, -5.5) [world]; Support Girdle (1215, -7.6) [world] |
| legs | Veteran's Chain Leggings (250493) | Blacksmithing [crafted] | 19.2 | yes | Defender's Leather Pants (252445, -0.1) [crafted]; Totemic Leather Pants (252446, -0.4) [crafted]; Deepgrave Trousers (279900, -4.3) [quest] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.8 | yes | Veteran's Boots (250503, -0.1) [crafted]; Guard's Boots (250504, -0.4) [crafted]; Defender's Leather Boots (252441, -0.4) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.7 | yes | Signet of the Zhevra (285330, -8.2) [world]; Bounty Hunter's Ring (5351, -8.4) [quest]; Lavishly Jeweled Ring (1156, -8.5) [dungeon] |
| finger2 | The 1 Ring (8350) | Fishing [world] | 2.2 | yes | Signet of the Zhevra (285330, -1.7) [world]; Bounty Hunter's Ring (5351, -1.9) [quest]; Lavishly Jeweled Ring (1156, -2.0) [dungeon] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 | yes | Hammerbone (270018, +58.1) [quest]; Forsaken Greataxe (251533, +54.8) [quest]; Smite's Mighty Hammer (7230, +50.3) [dungeon] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 237.1 | yes | Bear Buckler (4821, -230.9) [vendor]; Ruga's Bulwark (7120, -230.9) [quest]; Faerleia's Shield (3450, -232.9) [quest] |
| ranged | Cracked Blacksmith Hammer (285279) | Blacksmithing [crafted] | 4.2 | yes | Fine Longbow (11304, -0.2) [vendor]; Heavy Shortbow (3036, -2.1) [dungeon]; Orcish Battle Bow (5346, -2.1) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: The 1 Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Cracked Blacksmith Hammer

No-known-source sample (15 of 531, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 7956 Bronze Warhammer; 8178 Training Sword; 9602 Brushwood Blade; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse

### Band 30 (troll, 00000000000000000-353211005010000000-000000000000000000)

Set DPS (verified): 53.1. Weights run: 1.6s. Verify run: 1.8s. 1000 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.160, strength=1.676 ± 0.208, agility=not significant (0.186 ± 0.061), crit=3.475 ± 0.238, hit=1.401 ± 0.221, melee_haste=1.775 ± 0.290

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 21.8 | yes | Veteran's Chain Helm (250498, -1.7) [crafted]; Defender's Leather Helm (252455, -1.7) [crafted]; Crusader's Chain Helm (250502, -3.4) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19537, -12.5) [rep]; Scout's Medallion (20442, -12.9) [rep]; Pendant of Myzrael (4614, -14.0) [dungeon] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 11.7 | yes | Mail Combat Spaulders (6404, +0.0) [dungeon]; Barbaric Iron Shoulders (7913, -0.6) [crafted]; Forest Tracker Epaulets (2278, -1.3) [dungeon] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.0) [quest]; Sergeant Major's Cape (16315, -2.6) [pvp]; Lambent Scale Cloak (4706, -3.3) [dungeon] |
| chest | Shining Silver Breastplate (2870) | Blacksmithing [crafted] | 23.5 | yes | Barbaric Iron Breastplate (7914, -3.4) [crafted]; Hard Gold Cuirass (250533, -5.0) [crafted]; Veteran's Silvered Chain Shirt (250518, -5.6) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 13.4 | yes | Bands of Serra'kis (6902, -3.4) [dungeon]; Cultist's Armguards (270032, -3.4) [quest]; Grimtoll Wristguards (15459, -4.8) [quest] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 21.0 | yes | Fletcher's Gloves (7348, +27.6) [crafted]; Warsong Gauntlets (16978, -4.3) [quest]; Heavy Earthen Gloves (7359, -5.0) [crafted] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 | yes | Defiler's Leather Girdle (20191, +0.0) [rep]; Defiler's Plate Girdle (20207, -3.9) [rep]; Blackened Defias Belt (10403, -6.0) [dungeon] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 19.7 | yes | Ferine Leggings (6690, +6.3) [dungeon]; Golden Scale Leggings (3843, -1.3) [crafted]; Juggernaut Leggings (6671, -3.0) [quest] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 13.0 | yes | Hard Gold Boots (250534, -1.3) [crafted]; Glimmering Mail Greaves (4073, -3.0) [dungeon]; Brawler's Leather Boots (252439, -3.7) [crafted] |
| finger1 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 11.2 | yes | Ironspine's Eye (7686, -2.8) [dungeon]; Silverlaine's Family Seal (6321, -2.8) [dungeon]; Band of the Fist (17694, -3.7) [quest] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 | yes | Ironspine's Eye (7686, -0.6) [dungeon]; Silverlaine's Family Seal (6321, -0.6) [dungeon]; Band of the Fist (17694, -1.6) [quest] |
| trinket1 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.0) [rep]; Rune of Duty (21568, +0.0) [rep]; Relentless Raider's Seal (272062, +0.0) [vendor] |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.0) [rep]; Rune of Duty (21568, +0.0) [rep]; Relentless Raider's Seal (272062, +0.0) [vendor] |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 343.8 | yes | Morbid Dawn (7689, +95.3) [dungeon]; Manual Crowd Pummeler (9449, +89.9) [dungeon]; Corpsemaker (6687, +86.6) [dungeon] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 332.6 | yes | Commander's Crest (6320, -324.2) [dungeon]; Lambent Scale Shield (3656, -325.9) [dungeon]; Glimmering Shield (6400, -325.9) [dungeon] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 | yes | Double-barreled Shotgun (2098, -3.4) [dungeon]; Moonsight Rifle (4383, -4.8) [crafted]; Precision Bow (217315, -4.8) [quest] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wolfmaster Cape; chest: Shining Silver Breastplate; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Defiler's Chain Girdle; legs: Veteran's Silvered Chain Leggings; feet: Trouncing Boots; finger1: Legionnaire's Band; finger2: Insurgent's Band; trinket1: Defiler's Talisman; trinket2: Darkspear Voodoo Seal; main_hand: Pronged Reaver; off_hand: Ironspine's Fist; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 1000, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 7956 Bronze Warhammer

### Band 40 (troll, 00000000000000000-353211005050010050-000000000000000000)

Set DPS (verified): 83.5. Weights run: 1.9s. Verify run: 2.0s. 1445 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.203, strength=1.289 ± 0.264, agility=not significant (0.291 ± 0.084), crit=4.325 ± 0.374, hit=1.889 ± 0.312, melee_haste=2.652 ± 0.374

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 77.3 | yes | Icemetal Barbute (10763, -59.3) [dungeon]; Hard Gold Coif (250537, -59.3) [crafted]; White Bandit Mask (10008, -59.9) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Ethereal Talisman (4430, -6.4) [quest]; Scout's Medallion (19536, -10.8) [rep]; Scout's Medallion (19537, -11.7) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 15.2 | yes | Hard Gold Pauldrons (250539, -1.0) [crafted]; Shining Mithril Pauldrons (250541, -2.3) [crafted]; Imperial Leather Spaulders (4737, -3.6) [dungeon] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.0) [quest]; Sergeant Major's Cape (16336, -0.5) [pvp]; Sergeant Major's Cape (16315, -3.7) [pvp] |
| chest | Kolkar Marauder Chain (6773) | Quests [quest] | 20.1 | yes | Shining Silver Breastplate (2870, -2.0) [crafted]; Golden Scale Cuirass (3845, -2.0) [crafted]; Shining Mithril Breastplate (250540, -2.0) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Pugilist Bracers (4438, -9.7) [dungeon]; Ravager's Armguards (14770, -9.8) [world]; Cultist's Armguards (270032, -10.0) [quest] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 80.6 | yes | Dragonscale Gauntlets (8347, -18.3) [crafted]; Fletcher's Gloves (7348, -20.0) [crafted]; Ornate Mithril Gloves (7927, -20.0) [crafted] |
| waist | Defiler's Chain Girdle (20153) | The Defilers [rep] | 68.6 | yes | Defiler's Leather Girdle (20192, -38.6) [rep]; Defiler's Chain Girdle (20152, -44.6) [rep]; Defiler's Leather Girdle (20191, -44.6) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 27.1 | yes | Ferine Leggings (6690, -1.1) [dungeon]; Orcish War Leggings (7929, -5.2) [crafted]; Veteran's Silvered Chain Leggings (250523, -10.9) [crafted] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 18.8 | yes | Prowler's Leather Shoes (252465, -2.6) [crafted]; Blackforge Greaves (6423, -4.3) [dungeon]; Skirmisher's Mail Boots (252564, -4.6) [crafted] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 12.6 | yes | Legionnaire's Band (19513, -3.2) [rep]; Suspicious Spare Part (274754, -3.6) [vendor]; Ironspine's Eye (7686, -4.9) [dungeon] |
| finger2 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 12.0 | yes | Suspicious Spare Part (274754, -3.0) [vendor]; Insurgent's Band (272067, -3.0) [vendor]; Ironspine's Eye (7686, -4.2) [dungeon] |
| trinket1 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Tidal Charm (1404, +0.0) [vendor]; Blazing Emblem (2802, +0.0) [dungeon]; Cold Basilisk Eye (5079, +0.0) [world] |
| trinket2 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | 0.0 | yes | Tidal Charm (1404, +0.0) [vendor]; Blazing Emblem (2802, +0.0) [dungeon]; Cold Basilisk Eye (5079, +0.0) [world] |
| main_hand | Ardent Custodian (868) | Gnomeregan: Dark Iron Ambassador [dungeon] | 460.0 | yes | Staff of Jordan (873, +105.7) [dungeon]; Primitive Fishing Pole (276203, +105.6) [vendor]; Fiery War Axe (870, +104.7) [dungeon] |
| off_hand | Vanquisher's Sword (10823) | Quests [quest] | 442.0 | yes | Pit Fighter's Shield (4507, -431.7) [quest]; Combat Shield (4065, -433.0) [dungeon]; Aegis of the Scarlet Commander (7726, -433.0) [dungeon] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 | yes | Explosive Shotgun (8188, -2.6) [world]; Mithril Blacksmith Hammer (285280, -2.6) [crafted]; Master Hunter's Rifle (17687, -2.7) [quest] |

**New at 40:** head: Raging Berserker's Helm; shoulder: Sunburn Spaulders; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Chain Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Legionnaire's Band; finger2: Insurgent's Band; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Ardent Custodian; off_hand: Vanquisher's Sword

No-known-source sample (15 of 1445, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 50 (troll, 35100000000000000-353211005050010051-000000000000000000)

Set DPS (verified): 93.7. Weights run: 2.0s. Verify run: 2.0s. 1943 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.217, strength=1.796 ± 0.296, agility=not significant (0.302 ± 0.101), crit=4.820 ± 0.391, hit=2.018 ± 0.331, melee_haste=2.822 ± 0.403

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Plate Helm (220803) (or Knight-Lieutenant's Plate Helm (220804)) | Lady Palanseer [vendor] | 111.0 | yes | Knight-Lieutenant's Plate Helm (220804, +0.0) [vendor]; Raging Berserker's Helm (7719, -20.2) [dungeon]; Ornate Mithril Helm (7937, -25.6) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Woven Ivy Necklace (19159, -0.5) [quest]; Ethereal Talisman (4430, -3.8) [quest]; Scout's Medallion (19535, -10.4) [rep] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) (or Blood Guard's Plate Pauldrons (220796)) | Captain Dirgehammer [vendor] | 87.2 | yes | Blood Guard's Plate Pauldrons (220796, +0.0) [vendor]; Officer's Pauldrons (250576, -61.5) [crafted]; Wyrmslayer Spaulders (13066, -63.3) [world] |
| back | Bloodlust Cape (14801) | Maraudon: Princess Theradras [dungeon] | 16.2 | yes | Sergeant Major's Cape (16336, -3.6) [pvp]; Wolfmaster Cape (6314, -6.2) [dungeon]; Battlehard Cape (11858, -6.2) [quest] |
| chest | Knight's Plate Hauberk (220794) (or Stone Guard's Plate Armor (220801)) | Captain Dirgehammer [vendor] | 94.4 | yes | Stone Guard's Plate Armor (220801, +0.0) [vendor]; Ornate Mithril Breastplate (7935, -26.9) [crafted]; Warforged Chestplate (11195, -51.3) [quest] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 | yes | Officer's Wristguards (250581, -6.4) [crafted]; Branded Leather Bracers (19508, -8.0) [dungeon]; Prowler's Leather Bracers (252539, -9.7) [crafted] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 87.5 | yes | Dragonscale Gauntlets (8347, -18.2) [crafted]; Fletcher's Gloves (7348, -20.0) [crafted]; Ornate Mithril Gloves (7927, -20.0) [crafted] |
| waist | Defiler's Chain Girdle (20151) (or Defiler's Leather Girdle (20193)) | The Defilers [rep] | 87.5 | yes | Defiler's Leather Girdle (20193, +0.0) [rep]; Defiler's Plate Girdle (20205, -2.0) [rep]; Defiler's Chain Girdle (20153, -12.0) [rep] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | 92.1 | yes | Stormshroud Pants (15057, +42.9) [crafted]; Stone Guard's Plate Leggings (220798, +0.0) [vendor]; Scarlet Leggings (10330, -54.3) [dungeon] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 26.7 | yes | Officer's Sabatons (250561, -0.6) [crafted]; Officer's Boots (250546, -1.2) [crafted]; Skulker's Leather Boots (252469, -3.0) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 40.2 | yes | Legionnaire's Band (19511, -19.5) [rep]; Legionnaire's Band (19512, -23.4) [rep]; Band of Allegiance (18585, -24.0) [quest] |
| finger2 | White Bone Band (11862) | Quests [quest] | 24.0 | yes | Legionnaire's Band (19511, -3.3) [rep]; Legionnaire's Band (19512, -7.2) [rep]; Band of Allegiance (18585, -7.8) [quest] |
| trinket1 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | 0.0 | yes | Rune of the Guard Captain (19120, +56.1) [quest]; Tidal Charm (1404, +0.0) [vendor]; Guardian Talisman (1490, +0.0) [quest] |
| trinket2 | Defiler's Talisman (21115) | The Defilers [rep] | 0.0 | yes | Rune of the Guard Captain (19120, +56.1) [quest]; Tidal Charm (1404, +0.0) [vendor]; Guardian Talisman (1490, +0.0) [quest] |
| main_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | 577.5 | yes | Kindling Stave (11750, +132.0) [dungeon]; Darkspear Raider's Reaper (272080, +112.3) [vendor]; Bleakwood Hew (12769, +93.0) [crafted] |
| off_hand | Hammer of the Northern Wind (810) | Blackrock Depths: Anvilrage Medic [dungeon] | 553.3 | yes | Claw of Celebras (17738, -71.1) [dungeon]; White Bone Shredder (11863, -109.0) [quest]; Shizzle's Drizzle Blocker (11915, -529.9) [quest] |
| ranged | Houndmaster's Bow (11628) | Blackrock Depths: Houndmaster Grebmar [dungeon] | 17.4 | yes | Arcanite Blacksmith Hammer (285281, -6.6) [crafted]; Booty Bay Bruiser's Buckshot (274748, -8.4) [vendor]; Explosive Shotgun (8188, -8.4) [world] |

**New at 50:** head: Blood Guard's Plate Helm; shoulder: Knight-Lieutenant's Plate Pauldrons; back: Bloodlust Cape; chest: Knight's Plate Hauberk; wrist: Bracers of the Stone Princess; waist: Defiler's Chain Girdle; legs: Knight's Plate Leggings; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: White Bone Band; trinket1: Smoking Heart of the Mountain; trinket2: Defiler's Talisman; main_hand: Dawn's Edge; off_hand: Hammer of the Northern Wind; ranged: Houndmaster's Bow

No-known-source sample (15 of 1943, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (troll, 35311103002000000-353211005050010051-000000000000000000)

Set DPS (verified): 341.6. Weights run: 2.1s. Verify run: 2.5s. 2755 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=not significant (1.000 ± 0.427), strength=not significant (2.344 ± 0.596), agility=not significant (0.450 ± 0.239), crit=13.060 ± 0.873, hit=not significant (0.000 ± 0.000), melee_haste=6.782 ± 0.884

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 407.9 | yes | Ragefury Eyepatch (11735, -28.1) [dungeon]; Bloodvine Lens (19998, -42.2) [crafted]; Field Marshal's Plate Helm (16478, -159.4) [vendor] |
| neck | Stormrage's Talisman of Seething (23053) | Naxxramas [raid] | 391.7 | yes | Gem of Trapped Innocents (23057, -26.0) [raid]; Barbed Choker (21664, -164.8) [raid]; Medallion of the Dawn (22659, -184.8) [quest] |
| shoulder | Champion's Plate Shoulders (23243) (or Lieutenant Commander's Plate Shoulders (23315), Champion's Plate Shoulders (227042), Lieutenant Commander's Plate Shoulders (227045)) | Lady Palanseer [vendor] | 222.7 | yes | Lieutenant Commander's Plate Shoulders (23315, +0.0) [vendor]; Champion's Plate Shoulders (227042, +0.0) [pvp]; Lieutenant Commander's Plate Shoulders (227045, +0.0) [pvp] |
| back | Drape of Unyielding Strength (21394) | Quests [quest] | 39.2 | yes | Chromatic Cloak (18509, +143.6) [crafted]; Cloak of the Fallen God (21710, -1.7) [quest]; Deathguard's Cloak (20068, -3.0) [rep] |
| chest | Bloodsoul Breastplate (19690) | Blacksmithing [crafted] | 369.7 | yes | Stormshroud Armor (15056, -4.1) [crafted]; Savage Gladiator Chain (11726, -51.7) [dungeon]; Obsidian Mail Tunic (22191, -110.9) [crafted] |
| wrist | Vambraces of the Sadist (13400) | Stratholme: Timmy the Cruel [dungeon] | 199.2 | yes | Hive Defiler Wristguards (21618, -137.2) [raid]; Deeprock Bracers (21184, -150.2) [quest]; Berserker Bracers (19578, -151.1) [rep] |
| hands | Gauntlets of Annihilation (21581) | Ahn'Qiraj [raid] | 264.9 | yes | Marshal's Plate Gauntlets (16484, -35.2) [vendor]; General's Plate Gauntlets (16548, -35.2) [vendor]; General's Plate Gauntlets (231532, -35.2) [pvp] |
| waist | Belt of Never-ending Agony (21586) | Ahn'Qiraj [raid] | 246.8 | yes | Zandalar Vindicator's Belt (19823, -5.4) [quest]; Defiler's Plate Girdle (20204, -24.2) [rep]; Defiler's Chain Girdle (20150, -30.0) [rep] |
| legs | Marshal's Plate Legguards (16479) (or General's Plate Leggings (16543), General's Plate Leggings (231533), Marshal's Plate Legguards (231540)) | Captain Dirgehammer [vendor] | 412.6 | yes | General's Plate Leggings (16543, +0.0) [vendor]; General's Plate Leggings (231533, +0.0) [pvp]; Marshal's Plate Legguards (231540, +0.0) [pvp] |
| feet | Conqueror's Greaves (21333) | Quests [quest] | 56.9 | yes | Marshal's Plate Boots (16483, -9.3) [vendor]; General's Plate Boots (16545, -9.3) [vendor]; General's Plate Boots (231531, -9.3) [pvp] |
| finger1 | Band of Unnatural Forces (23038) | Naxxramas [raid] | 234.8 | yes | Ring of the Qiraji Fury (21677, -12.0) [raid]; Signet of Unyielding Strength (21393, -26.2) [quest]; Don Julio's Band (19325, -36.0) [rep] |
| finger2 | Quick Strike Ring (18821) | Molten Core: Magmadar [raid] | 224.6 | yes | Ring of the Qiraji Fury (21677, -1.7) [raid]; Signet of Unyielding Strength (21393, -15.9) [quest]; Don Julio's Band (19325, -25.7) [rep] |
| trinket1 | Slayer's Crest (23041) | Naxxramas [raid] | 64.0 | yes | Eye of Diminution (23001, +301.7) [raid]; Drake Fang Talisman (19406, -8.0) [raid]; Rune of the Guard Captain (19120, -22.0) [quest] |
| trinket2 | Kiss of the Spider (22954) | Naxxramas [raid] | 182.8 | yes | Eye of Diminution (23001, +182.8) [raid]; Drake Fang Talisman (19406, -126.8) [raid]; Rune of the Guard Captain (19120, -140.8) [quest] |
| main_hand | Thunderfury, Blessed Blade of the Windseeker (19019) | Quests [quest] | 757.6 | yes | High Warlord's Greatsword (234542, +612.2) [pvp]; High Warlord's Battle Axe (234543, +612.2) [pvp]; High Warlord's Pulverizer (234545, +612.2) [pvp] |
| off_hand | Grand Marshal's Swiftblade (234579) | Rank 18 [pvp] | 1078.0 | yes | High Warlord's Left Claw (234558, -3.1) [pvp]; Grand Marshal's Left Hand Blade (234584, -3.1) [pvp]; Grand Marshal's Left Hand Blade (18847, -34.4) [vendor] |
| ranged | Larvae of the Great Worm (23557) | Ahn'Qiraj [raid] | 200.8 | yes | The Purifier (22656, -18.0) [quest]; Bloodseeker (19107, -178.9) [quest]; Houndmaster's Bow (11628, -181.8) [dungeon] |

**New at 60:** head: Lionheart Helm; neck: Stormrage's Talisman of Seething; shoulder: Champion's Plate Shoulders; back: Drape of Unyielding Strength; chest: Bloodsoul Breastplate; wrist: Vambraces of the Sadist; hands: Gauntlets of Annihilation; waist: Belt of Never-ending Agony; legs: Marshal's Plate Legguards; feet: Conqueror's Greaves; finger1: Band of Unnatural Forces; finger2: Quick Strike Ring; trinket1: Slayer's Crest; trinket2: Kiss of the Spider; main_hand: Thunderfury, Blessed Blade of the Windseeker; off_hand: Grand Marshal's Swiftblade; ranged: Larvae of the Great Worm

No-known-source sample (15 of 2755, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

