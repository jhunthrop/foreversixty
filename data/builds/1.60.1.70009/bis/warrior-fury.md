# Leveling BiS: Fury

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power (Onyxia Tooth Pendant, Earthweave Cloak) still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-353000000000000000-000000000000000000)

Set DPS (verified): 32.3. Weights run: 1.5s. Verify run: 1.5s. 438 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.036, strength=2.009 ± 0.045, agility=0.098 ± 0.015, crit=2.375 ± 0.071, hit=not significant (0.000 ± 0.000), melee_haste=1.740 ± 0.046

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.1 | yes | Defender's Leather Hood (252447, -0.33 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.69 DPS) [crafted]; Brawler's Leather Hood (252504, -0.70 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 0.6 | yes | Scholarly Pendant (277203, -0.02 DPS) [quest]; Tarnished Locket (279870, -0.02 DPS) [quest]; Erudite's Amulet (277204, -0.04 DPS, sim-verified) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 | yes | Silvered Bronze Shoulders (3481, +0.00 DPS, sim-verified) [crafted]; Serpent's Shoulders (5404, -0.20 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.22 DPS) [crafted] |
| back | Lambent Scale Cloak (4706) | Gnomeregan: Caverndeep Burrower [dungeon] | 8.0 | yes | Catacomb Cloak (279899, -0.07 DPS) [quest]; Grave Shroud (279865, -0.10 DPS, sim-verified) [quest]; Dark Leather Cloak (2316, -0.13 DPS) [crafted] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.1 | yes | Defender's Leather Armor (252434, -0.21 DPS) [crafted]; Totemic Leather Armor (252435, -0.22 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.34 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 | yes | Cryptwalker Bracers (280095, -0.14 DPS, sim-verified) [quest]; Bravo's Armbands (270015, -0.20 DPS) [quest]; Runed Copper Bracers (2854, -0.22 DPS) [crafted] |
| hands | Gold-flecked Gloves (5195) | The Deadmines: Sneed [dungeon] | sim-verified (32.3 DPS) | yes | Polar Gauntlets (7606, -0.07 DPS) [quest]; Blackened Defias Gloves (10401, -0.07 DPS) [dungeon]; Fletcher's Gloves (7348, -0.49 DPS, sim-verified) [crafted] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 | yes | Ruffian Belt (5975, -0.22 DPS) [world]; Cobrahn's Grasp (6460, -0.22 DPS, sim-verified) [dungeon]; Support Girdle (1215, -0.29 DPS) [world] |
| legs | Chausses of Westfall (6087) | The Defias Brotherhood [quest] | 22.1 | yes | Defender's Leather Pants (252445, -0.13 DPS) [crafted]; Totemic Leather Pants (252446, -0.15 DPS) [crafted]; Veteran's Chain Leggings (250493, -0.15 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.5 | yes | Veteran's Boots (250503, +0.00 DPS, sim-verified) [crafted]; Guard's Boots (250504, -0.02 DPS) [crafted]; Defender's Leather Boots (252441, -0.02 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.4 | yes | The 1 Ring (8350, -0.23 DPS) [world]; Signet of the Zhevra (285330, -0.28 DPS) [world]; Lavishly Jeweled Ring (1156, -0.30 DPS) [dungeon] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 6.0 | yes | Signet of the Zhevra (285330, -0.20 DPS) [world]; Lavishly Jeweled Ring (1156, -0.21 DPS) [dungeon]; The 1 Ring (8350, -0.31 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 | yes | Living Root (6631, +0.00 DPS) [dungeon]; Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 236.8 | yes | Diamond Hammer (2194, -0.40 DPS, sim-verified) [dungeon]; Bear Buckler (4821, -8.36 DPS) [vendor]; Furen's Favor (6970, -8.36 DPS) [quest] |
| ranged | Ranger Bow (3021) | Blackfathom Deeps: Murkshallow Snapclaw [dungeon] | 347.5 | yes | Cracked Blacksmith Hammer (285279, +0.00 DPS, sim-verified) [crafted]; Lovingly Crafted Boomstick (4372, -2.87 DPS) [crafted]; Venomstrike (6469, -3.01 DPS) [dungeon] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Gold-flecked Gloves; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Loop of Sacrifice; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Ranger Bow

No-known-source sample (15 of 438, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 7956 Bronze Warhammer; 8178 Training Sword; 9602 Brushwood Blade; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse

### Band 30 (human, 00000000000000000-353211005010000000-000000000000000000)

Set DPS (verified): 53.1. Weights run: 1.7s. Verify run: 1.7s. 891 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.218, strength=2.528 ± 0.266, agility=not significant (0.203 ± 0.078), crit=4.156 ± 0.307, hit=not significant (0.000 ± 0.000), melee_haste=2.879 ± 0.352

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 32.9 | yes | Defender's Leather Helm (252455, -0.09 DPS) [crafted]; Crusader's Chain Helm (250502, -0.18 DPS) [crafted]; Veteran's Chain Helm (250498, -0.19 DPS, sim-verified) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (20444, -0.45 DPS) [rep]; Erudite's Amulet (277204, -0.46 DPS) [quest]; Sentinel's Medallion (19541, -0.75 DPS, sim-verified) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 17.7 | yes | Mail Combat Spaulders (6404, +0.00 DPS, sim-verified) [dungeon]; Barbaric Iron Shoulders (7913, -0.05 DPS) [crafted]; Elite Shoulders (4835, -0.09 DPS) [vendor] |
| back | Sergeant Major's Cape (16315) | Rank 9 (Alliance) [pvp] | 10.9 | yes | Wolfmaster Cape (6314, -0.03 DPS) [dungeon]; Grave Shroud (279865, -0.10 DPS) [quest]; Lambent Scale Cloak (4706, -0.19 DPS, sim-verified) [dungeon] |
| chest | Shining Silver Breastplate (2870) | Blacksmithing [crafted] | 35.4 | yes | Hard Gold Cuirass (250533, -0.27 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.31 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.36 DPS, sim-verified) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 20.2 | yes | Patterned Bronze Bracers (2868, -0.27 DPS) [crafted]; Technician's Bracers (270042, -0.27 DPS) [quest]; Bands of Serra'kis (6902, -0.36 DPS, sim-verified) [dungeon] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | sim-verified (53.1 DPS) | yes | Bonefist Gauntlets (4465, -0.03 DPS) [world]; Mail Combat Gauntlets (4075, -0.09 DPS) [dungeon]; Fletcher's Gloves (7348, -0.92 DPS, sim-verified) [crafted] |
| waist | Highlander's Plate Girdle (20126) | The League of Arathor [rep] | 30.3 | yes | Officer's Belt (250556, -0.13 DPS) [crafted]; Highlander's Lamellar Girdle (20108, -0.19 DPS, sim-verified) [rep]; Highlander's Chain Girdle (20090, -0.22 DPS) [rep] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 29.2 | yes | Chausses of Westfall (6087, -0.05 DPS) [quest]; Ferine Leggings (6690, -0.11 DPS) [dungeon]; Golden Scale Leggings (3843, -0.30 DPS, sim-verified) [crafted] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 19.1 | yes | Glimmering Mail Greaves (4073, -0.14 DPS) [dungeon]; Brawler's Leather Boots (252439, -0.19 DPS) [crafted]; Hard Gold Boots (250534, -0.30 DPS, sim-verified) [crafted] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 16.4 | yes | Ironspine's Eye (7686, -0.16 DPS) [dungeon]; Protector's Band (20439, -0.19 DPS) [rep]; Insurgent's Band (272067, -0.26 DPS) [vendor] |
| finger2 | Silverlaine's Family Seal (6321) | Shadowfang Keep: Baron Silverlaine [dungeon] | 12.6 | yes | Ironspine's Eye (7686, +0.00 DPS, sim-verified) [dungeon]; Insurgent's Band (272067, -0.13 DPS) [vendor]; Seal of Wrynn (2933, -0.16 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 348.1 | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Morbid Dawn (7689, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 338.6 | yes | Shoni's Disarming Tool (9608, -4.00 DPS) [quest]; Commander's Crest (6320, -11.45 DPS) [dungeon]; Swinetusk Shank (6691, -21.07 DPS, sim-verified) [dungeon] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 522.4 | yes | Silver Star (3463, -0.59 DPS, sim-verified) [quest]; Nightstalker Bow (6696, -2.15 DPS) [dungeon]; Satchel of Bronze Bombs (285276, -3.11 DPS) [crafted] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Sergeant Major's Cape; chest: Shining Silver Breastplate; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Highlander's Plate Girdle; legs: Veteran's Silvered Chain Leggings; feet: Trouncing Boots; finger1: Protector's Band; finger2: Silverlaine's Family Seal; trinket1: Talisman of Arathor; trinket2: Darkspear Voodoo Seal; main_hand: Pronged Reaver; off_hand: Ironspine's Fist; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 891, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 7956 Bronze Warhammer

### Band 40 (human, 00000000000000000-353211005050010050-000000000000000000)

Set DPS (verified): 87.3. Weights run: 1.9s. Verify run: 1.9s. 1330 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=not significant (1.000 ± 0.333), strength=1.931 ± 0.393, agility=0.609 ± 0.145, crit=7.890 ± 0.549, hit=not significant (0.000 ± 0.000), melee_haste=3.559 ± 0.547

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 135.6 | yes | White Bandit Mask (10008, -2.01 DPS, sim-verified) [crafted]; Icemetal Barbute (10763, -5.27 DPS) [dungeon]; Hard Gold Coif (250537, -5.27 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19541, -0.44 DPS) [rep]; Sentinel's Medallion (20444, -0.50 DPS) [rep]; Sentinel's Medallion (19540, -0.75 DPS, sim-verified) [rep] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 21.2 | yes | Sunburn Spaulders (274751, -0.12 DPS) [vendor]; Imperial Leather Spaulders (4737, -0.19 DPS) [dungeon]; Shining Mithril Pauldrons (250541, -0.47 DPS, sim-verified) [crafted] |
| back | Sergeant Major's Cape (16336) | Rank 9 (Alliance) [pvp] | 15.2 | yes | Wolfmaster Cape (6314, -0.25 DPS) [dungeon]; Yeti Fur Cloak (2805, -0.28 DPS) [quest]; Sergeant Major's Cape (16315, -0.99 DPS, sim-verified) [pvp] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 31.3 | yes | Golden Scale Cuirass (3845, -0.21 DPS) [crafted]; Shining Mithril Breastplate (250540, -0.21 DPS) [crafted]; Shining Silver Breastplate (2870, -0.98 DPS, sim-verified) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Pugilist Bracers (4438, -0.22 DPS) [dungeon]; Golden Scale Bracers (6040, -0.41 DPS) [crafted]; Ravager's Armguards (14770, -0.69 DPS, sim-verified) [world] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 130.5 | yes | Fletcher's Gloves (7348, -0.97 DPS) [crafted]; Ornate Mithril Gloves (7927, -0.97 DPS) [crafted]; Dragonscale Gauntlets (8347, -1.34 DPS, sim-verified) [crafted] |
| waist | Highlander's Chain Girdle (20089) | The League of Arathor [rep] | 118.5 | yes | Highlander's Leather Girdle (20116, -0.25 DPS, sim-verified) [rep]; Highlander's Plate Girdle (20125, -4.34 DPS) [rep]; Highlander's Chain Girdle (20090, -4.58 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 40.5 | yes | Ferine Leggings (6690, -0.71 DPS) [dungeon]; Veteran's Silvered Chain Leggings (250523, -0.73 DPS) [crafted]; Orcish War Leggings (7929, -1.08 DPS, sim-verified) [crafted] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 29.4 | yes | Blackforge Greaves (6423, -0.29 DPS) [dungeon]; Skirmisher's Mail Boots (252564, -0.39 DPS) [crafted]; Prowler's Leather Shoes (252465, -0.70 DPS, sim-verified) [crafted] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 20.3 | yes | Ironspine's Eye (7686, -0.35 DPS) [dungeon]; Insurgent's Band (272066, -0.40 DPS) [vendor]; Protector's Band (19517, -0.99 DPS, sim-verified) [rep] |
| finger2 | Suspicious Spare Part (274754) | Rettrick [vendor] | 13.5 | yes | Insurgent's Band (272066, -0.07 DPS) [vendor]; Ironspine's Eye (7686, -0.17 DPS, sim-verified) [dungeon]; Ring of the Underwood (2951, -0.17 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ardent Custodian (868) | Gnomeregan: Dark Iron Ambassador [dungeon] | 460.0 | yes | Bonebiter (6830, +0.00 DPS) [quest]; Illusionary Rod (7713, +0.00 DPS) [dungeon]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | 442.0 | yes | Jhordy's Misplaced Screwdriver (274753, -4.79 DPS, sim-verified) [vendor]; Shoni's Disarming Tool (9608, -10.54 DPS) [quest]; Salbac Shield (4652, -20.60 DPS) [quest] |
| ranged | Sniper Rifle (3430) (or Mithril Heavy-bore Rifle (10510)) | Gnomeregan: Dark Iron Agent [dungeon] | 415.7 | yes | Mithril Heavy-bore Rifle (10510, +0.00 DPS, sim-verified) [crafted]; Master Hunter's Rifle (17687, -0.26 DPS) [quest]; Mithril Blacksmith Hammer (285280, -0.41 DPS) [crafted] |

**New at 40:** head: Raging Berserker's Helm; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Chain Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Protector's Band; finger2: Suspicious Spare Part; trinket1: Ankh of Life; trinket2: Rune of Perfection; main_hand: Ardent Custodian; off_hand: Vanquisher's Sword; ranged: Sniper Rifle

No-known-source sample (15 of 1330, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 50 (human, 35100000000000000-353211005050010051-000000000000000000)

Set DPS (verified): 93.3. Weights run: 2.1s. Verify run: 2.0s. 1762 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.214, strength=1.584 ± 0.298, agility=not significant (0.351 ± 0.116), crit=4.668 ± 0.389, hit=not significant (0.000 ± 0.000), melee_haste=2.879 ± 0.401

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Plate Helm (220803) | Lady Palanseer [vendor] | sim-verified (85.7 DPS) | yes | Knight-Lieutenant's Plate Helm (220804, +0.00 DPS) [vendor]; Ornate Mithril Helm (7937, -0.45 DPS) [crafted]; Raging Berserker's Helm (7719, -1.05 DPS, sim-verified) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19539, -0.54 DPS, sim-verified) [rep]; Sentinel's Medallion (19540, -0.96 DPS) [rep]; Talisman of the Naga Lord (5029, -1.03 DPS) [world] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) (or Blood Guard's Plate Pauldrons (220796)) | Captain Dirgehammer [vendor] | 82.8 | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS, sim-verified) [vendor]; Officer's Pauldrons (250576, -5.62 DPS) [crafted]; Wyrmslayer Spaulders (13066, -5.77 DPS) [world_drop] |
| back | Bloodlust Cape (14801) | Maraudon: Princess Theradras [dungeon] | 14.3 | yes | Sergeant Major's Cape (16336, -0.20 DPS, sim-verified) [pvp]; Wolfmaster Cape (6314, -0.40 DPS) [dungeon]; Pridelord Cape (14673, -0.60 DPS) [dungeon] |
| chest | Knight's Plate Hauberk (220794) (or Stone Guard's Plate Armor (220801)) | Captain Dirgehammer [vendor] | 89.1 | yes | Stone Guard's Plate Armor (220801, +0.00 DPS, sim-verified) [vendor]; Ornate Mithril Breastplate (7935, -2.25 DPS) [crafted]; Warforged Chestplate (11195, -4.84 DPS) [quest] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 | yes | Officer's Wristguards (250581, -0.80 DPS) [crafted]; Prowler's Leather Bracers (252539, -1.07 DPS) [crafted]; Branded Leather Bracers (19508, -1.13 DPS, sim-verified) [dungeon] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 85.3 | yes | Dragonscale Gauntlets (8347, -1.84 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -1.89 DPS) [crafted]; Ornate Mithril Gloves (7927, -1.89 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20088) (or Highlander's Leather Girdle (20115)) | The League of Arathor [rep] | 85.3 | yes | Highlander's Leather Girdle (20115, +0.00 DPS, sim-verified) [rep]; Highlander's Lamellar Girdle (20106, -0.24 DPS) [rep]; Highlander's Plate Girdle (20124, -0.39 DPS) [rep] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | sim-verified (86.1 DPS) | yes | Stone Guard's Plate Leggings (220798, +0.00 DPS) [vendor]; Stormshroud Pants (15057, -1.51 DPS, sim-verified) [crafted]; Scarlet Leggings (10330, -5.17 DPS) [dungeon] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 24.5 | yes | Officer's Boots (250546, -0.13 DPS) [crafted]; Officer's Sabatons (250561, -0.19 DPS, sim-verified) [crafted]; Skulker's Leather Boots (252469, -0.23 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 20.0 | yes | Insurgent's Band (272065, -0.47 DPS) [vendor]; Insurgent's Band (272066, -0.76 DPS) [vendor]; Suspicious Spare Part (274754, -0.84 DPS) [vendor] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 19.0 | yes | Insurgent's Band (272065, -0.38 DPS) [vendor]; Protector's Band (19515, -0.41 DPS, sim-verified) [rep]; Insurgent's Band (272066, -0.66 DPS) [vendor] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (84.8 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [dungeon] |
| trinket2 | - | - |  |  |  |
| main_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | sim-verified (84.8 DPS) | yes | Kindling Stave (11750, +0.00 DPS) [dungeon]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor] |
| off_hand | Hammer of the Northern Wind (810) | Blackrock Depths: Anvilrage Medic [dungeon] | sim-verified (90.5 DPS) | yes | Inventor's Focal Sword (17719, -5.89 DPS, sim-verified) [dungeon]; Claw of Celebras (17738, -6.74 DPS) [dungeon]; Shoni's Disarming Tool (9608, -31.13 DPS) [quest] |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | sim-verified (84.8 DPS) | yes | Dark Iron Rifle (16004, -1.39 DPS, sim-verified) [crafted]; Houndmaster's Bow (11628, -4.14 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -4.81 DPS) [dungeon] |

**New at 50:** head: Blood Guard's Plate Helm; shoulder: Knight-Lieutenant's Plate Pauldrons; back: Bloodlust Cape; chest: Knight's Plate Hauberk; wrist: Bracers of the Stone Princess; waist: Highlander's Chain Girdle; legs: Knight's Plate Leggings; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Frozen Heart of the Mountain; trinket2: Smoking Heart of the Mountain; main_hand: Dawn's Edge; off_hand: Hammer of the Northern Wind; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 1762, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60 (human, 35311103002000000-353211005050010051-000000000000000000)

Set DPS (verified): 300.8. Weights run: 2.1s. Verify run: 2.3s. 2500 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=not significant (1.000 ± 0.427), strength=not significant (2.344 ± 0.596), agility=not significant (0.450 ± 0.239), crit=13.060 ± 0.873, hit=not significant (0.000 ± 0.000), melee_haste=6.782 ± 0.884

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 407.9 | yes | Bloodvine Lens (19998, -2.44 DPS) [crafted]; Lightbreaker Greathelm (239517, -5.70 DPS) [vendor]; Ragefury Eyepatch (11735, -9.78 DPS, sim-verified) [dungeon] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (287.7 DPS) | yes | Blazefury Medallion (17111, +0.00 DPS, sim-verified) [world]; Onyxia Tooth Pendant (18404, -1.20 DPS) [quest]; Choker of the Shifting Sands (21505, -9.52 DPS) [quest] |
| shoulder | Lightbreaker Shoulders (239516) | Leonid Barthalomew the Revered [vendor] | 269.3 | yes | Champion's Plate Shoulders (23243, -2.69 DPS) [vendor]; Lieutenant Commander's Plate Shoulders (23315, -2.69 DPS) [vendor]; Darkspear Spaulders (272108, -11.35 DPS, sim-verified) [vendor] |
| back | Drape of Unyielding Strength (21394) | Drape of Unyielding Strength [quest] | sim-verified (294.5 DPS) | yes | Cloak of the Fallen God (21710, -0.10 DPS) [quest]; Cloak of the Honor Guard (20073, -0.17 DPS) [rep]; Chromatic Cloak (18509, -9.26 DPS, sim-verified) [crafted] |
| chest | Lightbreaker Cuirass (239519) | Leonid Barthalomew the Revered [vendor] | 485.9 | yes | Stormshroud Armor (15056, -6.94 DPS) [crafted]; Savage Gladiator Chain (11726, -9.69 DPS) [dungeon]; Bloodsoul Breastplate (19690, -15.70 DPS, sim-verified) [crafted] |
| wrist | Lightbreaker Wrists (239512) | Leonid Barthalomew the Revered [vendor] | sim-verified (297.8 DPS) | yes | Deeprock Bracers (21184, -1.35 DPS) [quest]; Berserker Bracers (19578, -1.41 DPS) [rep]; Vambraces of the Sadist (13400, -12.58 DPS, sim-verified) [dungeon] |
| hands | Lightbreaker Grips (239514) | Leonid Barthalomew the Revered [vendor] | 268.5 | yes | Marshal's Plate Gauntlets (16484, -2.24 DPS) [vendor]; General's Plate Gauntlets (16548, -2.24 DPS) [vendor]; Chromatic Gauntlets (19157, -4.51 DPS, sim-verified) [crafted] |
| waist | Lightbreaker Belt (239513) | Leonid Barthalomew the Revered [vendor] | 271.7 | yes | Zandalar Vindicator's Belt (19823, -1.75 DPS) [quest]; Highlander's Plate Girdle (20041, -2.83 DPS) [rep]; Radiant Girdle of the Dawn (227814, -3.12 DPS, sim-verified) [vendor] |
| legs | Lightbreaker Tassets (239518) | Leonid Barthalomew the Revered [vendor] | 487.4 | yes | Sentinel's Plate Legguards (237825, -4.15 DPS, sim-verified) [vendor]; Marshal's Plate Legguards (16479, -4.32 DPS) [vendor]; General's Plate Leggings (16543, -4.32 DPS) [vendor] |
| feet | Lightbreaker Greaves (239515) | Leonid Barthalomew the Revered [vendor] | 269.4 | yes | Conqueror's Greaves (21333, -9.85 DPS, sim-verified) [quest]; Marshal's Plate Boots (16483, -12.80 DPS) [vendor]; General's Plate Boots (16545, -12.80 DPS) [vendor] |
| finger1 | Signet of Unyielding Strength (21393) | Signet of Unyielding Strength [quest] | 208.6 | yes | Band of Earthen Might (21182, -0.68 DPS) [quest]; Band of the Penitent (13217, -1.49 DPS) [quest]; Dragonslayer's Signet (18403, -1.49 DPS) [quest] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | 198.8 | yes | Band of Earthen Might (21182, +0.00 DPS, sim-verified) [quest]; Band of the Penitent (13217, -0.92 DPS) [quest]; Dragonslayer's Signet (18403, -0.92 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | High Warlord's Hacker (235476) | Sergeant Thunderhorn [vendor] | sim-verified (287.7 DPS) | yes | High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Ebon Hand (19170, -12.99 DPS, sim-verified) [crafted] |
| off_hand | High Warlord's Bonecracker (235477) | Sergeant Thunderhorn [vendor] | sim-verified (287.7 DPS) | yes | High Warlord's Left Claw (234558, -0.18 DPS) [vendor]; Grand Marshal's Left Hand Blade (234584, -0.18 DPS) [vendor]; Force Reactive Disk (18168, -81.14 DPS, sim-verified) [crafted] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (287.7 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [vendor]; High Warlord's Crossbow (234560, +0.00 DPS) [vendor]; Dark Iron Rifle (16004, -2.71 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Medallion of the Dawn; shoulder: Lightbreaker Shoulders; back: Drape of Unyielding Strength; chest: Lightbreaker Cuirass; wrist: Lightbreaker Wrists; hands: Lightbreaker Grips; waist: Lightbreaker Belt; legs: Lightbreaker Tassets; feet: Lightbreaker Greaves; finger1: Signet of Unyielding Strength; finger2: Don Julio's Band; trinket1: Onyxia Blood Talisman; trinket2: Weakness Analyzer; main_hand: High Warlord's Hacker; off_hand: High Warlord's Bonecracker; ranged: The Purifier

No-known-source sample (15 of 2500, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

## Horde

### Band 20 (troll, 00000000000000000-353000000000000000-000000000000000000)

Set DPS (verified): 29.0. Weights run: 1.5s. Verify run: 1.4s. 440 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.036, strength=2.009 ± 0.045, agility=0.098 ± 0.015, crit=2.375 ± 0.071, hit=not significant (0.000 ± 0.000), melee_haste=1.740 ± 0.046

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.1 | yes | Defender's Leather Hood (252447, -0.34 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.69 DPS) [crafted]; Brawler's Leather Hood (252504, -0.70 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 0.6 | yes | Scholarly Pendant (277203, -0.02 DPS) [quest]; Tarnished Locket (279870, -0.02 DPS) [quest]; Erudite's Amulet (277204, -0.05 DPS, sim-verified) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 | yes | Silvered Bronze Shoulders (3481, +0.00 DPS, sim-verified) [crafted]; Serpent's Shoulders (5404, -0.20 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.22 DPS) [crafted] |
| back | Lambent Scale Cloak (4706) | Gnomeregan: Caverndeep Burrower [dungeon] | 8.0 | yes | Subterranean Cape (14149, -0.07 DPS) [dungeon]; Catacomb Cloak (279899, -0.07 DPS) [quest]; Grave Shroud (279865, -0.16 DPS, sim-verified) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.1 | yes | Defender's Leather Armor (252434, -0.21 DPS) [crafted]; Totemic Leather Armor (252435, -0.22 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.45 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 | yes | Raptorcrest Bracers (270010, -0.15 DPS) [quest]; Runed Copper Bracers (2854, -0.22 DPS) [crafted]; Cryptwalker Bracers (280095, -0.25 DPS, sim-verified) [quest] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 33.3 | yes | Gold-flecked Gloves (5195, +0.00 DPS, sim-verified) [dungeon]; Blackened Defias Gloves (10401, -0.77 DPS) [dungeon]; Dagmire Gauntlets (6481, -0.83 DPS) [quest] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 | yes | Ruffian Belt (5975, -0.22 DPS) [world]; Cobrahn's Grasp (6460, -0.27 DPS, sim-verified) [dungeon]; Support Girdle (1215, -0.29 DPS) [world] |
| legs | Veteran's Chain Leggings (250493) | Blacksmithing [crafted] | 18.6 | yes | Totemic Leather Pants (252446, -0.02 DPS) [crafted]; Defender's Leather Pants (252445, -0.02 DPS, sim-verified) [crafted]; Deepgrave Trousers (279900, -0.15 DPS) [quest] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.5 | yes | Guard's Boots (250504, -0.02 DPS) [crafted]; Defender's Leather Boots (252441, -0.02 DPS) [crafted]; Veteran's Boots (250503, -0.02 DPS, sim-verified) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.4 | yes | The 1 Ring (8350, -0.23 DPS) [world]; Signet of the Zhevra (285330, -0.28 DPS) [world]; Bounty Hunter's Ring (5351, -0.29 DPS) [quest] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 6.0 | yes | Signet of the Zhevra (285330, -0.20 DPS) [world]; Bounty Hunter's Ring (5351, -0.21 DPS) [quest]; The 1 Ring (8350, -0.33 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 | yes | Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest]; Hammerbone (270018, +0.00 DPS) [quest] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 236.8 | yes | Diamond Hammer (2194, -0.49 DPS, sim-verified) [dungeon]; Bear Buckler (4821, -8.36 DPS) [vendor]; Ruga's Bulwark (7120, -8.36 DPS) [quest] |
| ranged | Ranger Bow (3021) | Blackfathom Deeps: Murkshallow Snapclaw [dungeon] | 347.5 | yes | Cracked Blacksmith Hammer (285279, +0.00 DPS, sim-verified) [crafted]; Lovingly Crafted Boomstick (4372, -2.87 DPS) [crafted]; Venomstrike (6469, -3.01 DPS) [dungeon] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Loop of Sacrifice; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Ranger Bow

No-known-source sample (15 of 440, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 7956 Bronze Warhammer; 8178 Training Sword; 9602 Brushwood Blade; 9747 Simple Britches; 9748 Simple Robe

### Band 30 (troll, 00000000000000000-353211005010000000-000000000000000000)

Set DPS (verified): 52.9. Weights run: 1.7s. Verify run: 1.7s. 896 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.218, strength=2.528 ± 0.266, agility=not significant (0.203 ± 0.078), crit=4.156 ± 0.307, hit=not significant (0.000 ± 0.000), melee_haste=2.879 ± 0.352

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 32.9 | yes | Defender's Leather Helm (252455, -0.09 DPS) [crafted]; Crusader's Chain Helm (250502, -0.18 DPS) [crafted]; Veteran's Chain Helm (250498, -0.34 DPS, sim-verified) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (20442, -0.45 DPS) [rep]; Erudite's Amulet (277204, -0.46 DPS) [quest]; Scout's Medallion (19537, -0.82 DPS, sim-verified) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 17.7 | yes | Mail Combat Spaulders (6404, +0.00 DPS, sim-verified) [dungeon]; Barbaric Iron Shoulders (7913, -0.05 DPS) [crafted]; Elite Shoulders (4835, -0.09 DPS) [vendor] |
| back | Lambent Scale Cloak (4706) | Gnomeregan: Caverndeep Burrower [dungeon] | 10.1 | yes | Wolfmaster Cape (6314, +0.00 DPS, sim-verified) [dungeon]; Wildhunter Cloak (16658, -0.00 DPS) [quest]; Grave Shroud (279865, -0.07 DPS) [quest] |
| chest | Shining Silver Breastplate (2870) | Blacksmithing [crafted] | 35.4 | yes | Hard Gold Cuirass (250533, -0.27 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.31 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.66 DPS, sim-verified) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 20.2 | yes | Grimtoll Wristguards (15459, -0.26 DPS) [quest]; Patterned Bronze Bracers (2868, -0.27 DPS) [crafted]; Bands of Serra'kis (6902, -0.66 DPS, sim-verified) [dungeon] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 58.2 | yes | Warsong Gauntlets (16978, +0.00 DPS, sim-verified) [quest]; Gauntlets of Ogre Strength (3341, -1.22 DPS) [world]; Bonefist Gauntlets (4465, -1.24 DPS) [world] |
| waist | Defiler's Plate Girdle (20207) | The Defilers [rep] | 30.3 | yes | Defiler's Chain Girdle (20152, -0.22 DPS) [rep]; Defiler's Leather Girdle (20191, -0.22 DPS) [rep]; Officer's Belt (250556, -0.54 DPS, sim-verified) [crafted] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 29.2 | yes | Ferine Leggings (6690, -0.11 DPS) [dungeon]; Juggernaut Leggings (6671, -0.14 DPS) [quest]; Golden Scale Leggings (3843, -0.29 DPS, sim-verified) [crafted] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 19.1 | yes | Glimmering Mail Greaves (4073, -0.14 DPS) [dungeon]; Brawler's Leather Boots (252439, -0.19 DPS) [crafted]; Hard Gold Boots (250534, -0.29 DPS, sim-verified) [crafted] |
| finger1 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 16.4 | yes | Ironspine's Eye (7686, -0.16 DPS) [dungeon]; Band of the Fist (17694, -0.19 DPS) [quest]; Legionnaire's Band (20429, -0.19 DPS) [rep] |
| finger2 | Silverlaine's Family Seal (6321) | Shadowfang Keep: Baron Silverlaine [dungeon] | 12.6 | yes | Ironspine's Eye (7686, +0.00 DPS, sim-verified) [dungeon]; Band of the Fist (17694, -0.06 DPS) [quest]; Insurgent's Band (272067, -0.13 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 348.1 | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Morbid Dawn (7689, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 338.6 | yes | Commander's Crest (6320, -11.45 DPS) [dungeon]; Lambent Scale Shield (3656, -11.54 DPS) [dungeon]; Swinetusk Shank (6691, -21.84 DPS, sim-verified) [dungeon] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 522.4 | yes | Silver Star (3463, -0.57 DPS, sim-verified) [quest]; Nightstalker Bow (6696, -2.15 DPS) [dungeon]; Satchel of Bronze Bombs (285276, -3.11 DPS) [crafted] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; chest: Shining Silver Breastplate; wrist: Pugilist Bracers; waist: Defiler's Plate Girdle; legs: Veteran's Silvered Chain Leggings; feet: Trouncing Boots; finger1: Legionnaire's Band; finger2: Silverlaine's Family Seal; trinket1: Defiler's Talisman; trinket2: Darkspear Voodoo Seal; main_hand: Pronged Reaver; off_hand: Ironspine's Fist; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 896, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore

### Band 40 (troll, 00000000000000000-353211005050010050-000000000000000000)

Set DPS (verified): 82.4. Weights run: 1.9s. Verify run: 2.0s. 1335 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=not significant (1.000 ± 0.333), strength=1.931 ± 0.393, agility=0.609 ± 0.145, crit=7.890 ± 0.549, hit=not significant (0.000 ± 0.000), melee_haste=3.559 ± 0.547

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 135.6 | yes | White Bandit Mask (10008, -0.98 DPS, sim-verified) [crafted]; Icemetal Barbute (10763, -5.27 DPS) [dungeon]; Hard Gold Coif (250537, -5.27 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Ethereal Talisman (4430, +0.00 DPS, sim-verified) [quest]; Scout's Medallion (19536, -0.35 DPS) [rep]; Scout's Medallion (19537, -0.44 DPS) [rep] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 21.2 | yes | Sunburn Spaulders (274751, -0.12 DPS) [vendor]; Imperial Leather Spaulders (4737, -0.19 DPS) [dungeon]; Shining Mithril Pauldrons (250541, -0.24 DPS, sim-verified) [crafted] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Lambent Scale Cloak (4706, -0.11 DPS) [dungeon]; Warden's Cloak (14602, -0.11 DPS) [world] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 31.3 | yes | Golden Scale Cuirass (3845, -0.21 DPS) [crafted]; Shining Mithril Breastplate (250540, -0.21 DPS) [crafted]; Shining Silver Breastplate (2870, -0.48 DPS, sim-verified) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Ravager's Armguards (14770, +0.00 DPS, sim-verified) [world]; Pugilist Bracers (4438, -0.22 DPS) [dungeon]; Darkspear Armsplints (4132, -0.31 DPS) [quest] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 130.5 | yes | Dragonscale Gauntlets (8347, -0.83 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -0.97 DPS) [crafted]; Ornate Mithril Gloves (7927, -0.97 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20153) | The Defilers [rep] | 118.5 | yes | Defiler's Leather Girdle (20192, -0.01 DPS, sim-verified) [rep]; Defiler's Plate Girdle (20206, -4.34 DPS) [rep]; Tharg's Shoelace (9705, -4.53 DPS) [quest] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 40.5 | yes | Orcish War Leggings (7929, -0.33 DPS, sim-verified) [crafted]; Ferine Leggings (6690, -0.71 DPS) [dungeon]; Veteran's Silvered Chain Leggings (250523, -0.73 DPS) [crafted] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 29.4 | yes | Prowler's Leather Shoes (252465, -0.04 DPS, sim-verified) [crafted]; Blackforge Greaves (6423, -0.29 DPS) [dungeon]; Skirmisher's Mail Boots (252564, -0.39 DPS) [crafted] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 20.3 | yes | Ironspine's Eye (7686, -0.35 DPS) [dungeon]; Insurgent's Band (272066, -0.40 DPS) [vendor]; Legionnaire's Band (19513, -0.41 DPS, sim-verified) [rep] |
| finger2 | Suspicious Spare Part (274754) | Rettrick [vendor] | 13.5 | yes | Ironspine's Eye (7686, +0.00 DPS, sim-verified) [dungeon]; Insurgent's Band (272066, -0.07 DPS) [vendor]; Band of the Fist (17694, -0.16 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ardent Custodian (868) | Gnomeregan: Dark Iron Ambassador [dungeon] | 460.0 | yes | Illusionary Rod (7713, +0.00 DPS) [dungeon]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 442.0 | yes | Jhordy's Misplaced Screwdriver (274753, +0.00 DPS, sim-verified) [vendor]; Pit Fighter's Shield (4507, -20.70 DPS) [quest]; Combat Shield (4065, -20.79 DPS) [dungeon] |
| ranged | Sniper Rifle (3430) (or Mithril Heavy-bore Rifle (10510)) | Gnomeregan: Dark Iron Agent [dungeon] | 415.7 | yes | Mithril Heavy-bore Rifle (10510, +0.00 DPS, sim-verified) [crafted]; Master Hunter's Rifle (17687, -0.26 DPS) [quest]; Mithril Blacksmith Hammer (285280, -0.41 DPS) [crafted] |

**New at 40:** head: Raging Berserker's Helm; shoulder: Hard Gold Pauldrons; back: Wolfmaster Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Chain Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Legionnaire's Band; finger2: Suspicious Spare Part; trinket1: Ankh of Life; trinket2: Rune of Perfection; main_hand: Ardent Custodian; off_hand: Vanquisher's Sword; ranged: Sniper Rifle

No-known-source sample (15 of 1335, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 50 (troll, 35100000000000000-353211005050010051-000000000000000000)

Set DPS (verified): 97.3. Weights run: 2.1s. Verify run: 2.0s. 1767 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.214, strength=1.584 ± 0.298, agility=not significant (0.351 ± 0.116), crit=4.668 ± 0.389, hit=not significant (0.000 ± 0.000), melee_haste=2.879 ± 0.401

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Plate Helm (220803) | Lady Palanseer [vendor] | sim-verified (79.1 DPS) | yes | Knight-Lieutenant's Plate Helm (220804, +0.00 DPS) [vendor]; Ornate Mithril Helm (7937, -0.45 DPS) [crafted]; Raging Berserker's Helm (7719, -1.04 DPS, sim-verified) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Woven Ivy Necklace (19159, +0.00 DPS, sim-verified) [quest]; Ethereal Talisman (4430, -0.44 DPS) [quest]; Scout's Medallion (19535, -0.93 DPS) [rep] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) (or Blood Guard's Plate Pauldrons (220796)) | Captain Dirgehammer [vendor] | 82.8 | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS, sim-verified) [vendor]; Officer's Pauldrons (250576, -5.62 DPS) [crafted]; Wyrmslayer Spaulders (13066, -5.77 DPS) [world_drop] |
| back | Bloodlust Cape (14801) | Maraudon: Princess Theradras [dungeon] | 14.3 | yes | Battlehard Cape (11858, -0.40 DPS) [quest]; Wildhunter Cloak (16658, -0.40 DPS) [quest]; Wolfmaster Cape (6314, -0.48 DPS, sim-verified) [dungeon] |
| chest | Knight's Plate Hauberk (220794) (or Stone Guard's Plate Armor (220801)) | Captain Dirgehammer [vendor] | 89.1 | yes | Stone Guard's Plate Armor (220801, +0.00 DPS, sim-verified) [vendor]; Ornate Mithril Breastplate (7935, -2.25 DPS) [crafted]; Warforged Chestplate (11195, -4.84 DPS) [quest] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 | yes | Branded Leather Bracers (19508, -0.48 DPS, sim-verified) [dungeon]; Officer's Wristguards (250581, -0.80 DPS) [crafted]; Prowler's Leather Bracers (252539, -1.07 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 85.3 | yes | Dragonscale Gauntlets (8347, -1.72 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -1.89 DPS) [crafted]; Ornate Mithril Gloves (7927, -1.89 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20151) (or Defiler's Leather Girdle (20193)) | The Defilers [rep] | 85.3 | yes | Defiler's Leather Girdle (20193, +0.00 DPS, sim-verified) [rep]; Defiler's Plate Girdle (20205, -0.39 DPS) [rep]; Defiler's Chain Girdle (20153, -1.14 DPS) [rep] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | sim-verified (79.3 DPS) | yes | Stone Guard's Plate Leggings (220798, +0.00 DPS) [vendor]; Stormshroud Pants (15057, -1.26 DPS, sim-verified) [crafted]; Scarlet Leggings (10330, -5.17 DPS) [dungeon] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 24.5 | yes | Officer's Boots (250546, -0.13 DPS) [crafted]; Officer's Sabatons (250561, -0.14 DPS, sim-verified) [crafted]; Skulker's Leather Boots (252469, -0.23 DPS) [crafted] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 | yes | Legionnaire's Band (19511, -0.47 DPS) [rep]; Legionnaire's Band (19512, -0.81 DPS) [rep]; Insurgent's Band (272065, -0.85 DPS) [vendor] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 20.0 | yes | Legionnaire's Band (19511, -0.28 DPS, sim-verified) [rep]; Legionnaire's Band (19512, -0.43 DPS) [rep]; Insurgent's Band (272065, -0.47 DPS) [vendor] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (78.4 DPS) | yes | Frozen Heart of the Mountain (249469, -3.53 DPS, sim-verified) [crafted]; Guardian Talisman (1490, -3.98 DPS) [quest]; Ankh of Life (1713, -3.98 DPS) [dungeon] |
| trinket2 | - | - |  |  |  |
| main_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | sim-verified (78.4 DPS) | yes | Kindling Stave (11750, +0.00 DPS) [dungeon]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor] |
| off_hand | Hammer of the Northern Wind (810) | Blackrock Depths: Anvilrage Medic [dungeon] | sim-verified (95.2 DPS) | yes | Claw of Celebras (17738, -6.74 DPS) [dungeon]; White Bone Shredder (11863, -10.37 DPS) [quest]; Inventor's Focal Sword (17719, -17.14 DPS, sim-verified) [dungeon] |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | sim-verified (78.4 DPS) | yes | Dark Iron Rifle (16004, -0.50 DPS, sim-verified) [crafted]; Houndmaster's Bow (11628, -4.14 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -4.81 DPS) [dungeon] |

**New at 50:** head: Blood Guard's Plate Helm; shoulder: Knight-Lieutenant's Plate Pauldrons; back: Bloodlust Cape; chest: Knight's Plate Hauberk; wrist: Bracers of the Stone Princess; waist: Defiler's Chain Girdle; legs: Knight's Plate Leggings; feet: Prowler's Leather Boots; finger1: White Bone Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Smoking Heart of the Mountain; main_hand: Dawn's Edge; off_hand: Hammer of the Northern Wind; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 1767, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60 (troll, 35311103002000000-353211005050010051-000000000000000000)

Set DPS (verified): 315.9. Weights run: 2.1s. Verify run: 2.4s. 2506 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=not significant (1.000 ± 0.427), strength=not significant (2.344 ± 0.596), agility=not significant (0.450 ± 0.239), crit=13.060 ± 0.873, hit=not significant (0.000 ± 0.000), melee_haste=6.782 ± 0.884

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 407.9 | yes | Bloodvine Lens (19998, -2.44 DPS) [crafted]; Lightbreaker Greathelm (239517, -5.70 DPS) [vendor]; Ragefury Eyepatch (11735, -11.61 DPS, sim-verified) [dungeon] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (300.8 DPS) | yes | Blazefury Medallion (17111, +0.00 DPS, sim-verified) [world]; Onyxia Tooth Pendant (18404, -1.20 DPS) [quest]; Choker of the Shifting Sands (21505, -9.52 DPS) [quest] |
| shoulder | Lightbreaker Shoulders (239516) | Leonid Barthalomew the Revered [vendor] | 269.3 | yes | Champion's Plate Shoulders (23243, -2.69 DPS) [vendor]; Lieutenant Commander's Plate Shoulders (23315, -2.69 DPS) [vendor]; Darkspear Spaulders (272108, -14.30 DPS, sim-verified) [vendor] |
| back | Drape of Unyielding Strength (21394) | Drape of Unyielding Strength [quest] | sim-verified (307.1 DPS) | yes | Cloak of the Fallen God (21710, -0.10 DPS) [quest]; Deathguard's Cloak (20068, -0.17 DPS) [rep]; Chromatic Cloak (18509, -8.32 DPS, sim-verified) [crafted] |
| chest | Lightbreaker Cuirass (239519) | Leonid Barthalomew the Revered [vendor] | 485.9 | yes | Stormshroud Armor (15056, -6.94 DPS) [crafted]; Savage Gladiator Chain (11726, -9.69 DPS) [dungeon]; Bloodsoul Breastplate (19690, -16.64 DPS, sim-verified) [crafted] |
| wrist | Lightbreaker Wrists (239512) | Leonid Barthalomew the Revered [vendor] | sim-verified (308.8 DPS) | yes | Deeprock Bracers (21184, -1.35 DPS) [quest]; Berserker Bracers (19578, -1.41 DPS) [rep]; Vambraces of the Sadist (13400, -10.10 DPS, sim-verified) [dungeon] |
| hands | Lightbreaker Grips (239514) | Leonid Barthalomew the Revered [vendor] | 268.5 | yes | Marshal's Plate Gauntlets (16484, -2.24 DPS) [vendor]; General's Plate Gauntlets (16548, -2.24 DPS) [vendor]; Chromatic Gauntlets (19157, -4.28 DPS, sim-verified) [crafted] |
| waist | Lightbreaker Belt (239513) | Leonid Barthalomew the Revered [vendor] | 271.7 | yes | Zandalar Vindicator's Belt (19823, -1.75 DPS) [quest]; Defiler's Plate Girdle (20204, -2.83 DPS) [rep]; Radiant Girdle of the Dawn (227814, -3.49 DPS, sim-verified) [vendor] |
| legs | Lightbreaker Tassets (239518) | Leonid Barthalomew the Revered [vendor] | 487.4 | yes | Marshal's Plate Legguards (16479, -4.32 DPS) [vendor]; General's Plate Leggings (16543, -4.32 DPS) [vendor]; Sentinel's Plate Legguards (237825, -6.08 DPS, sim-verified) [vendor] |
| feet | Lightbreaker Greaves (239515) | Leonid Barthalomew the Revered [vendor] | 269.4 | yes | Conqueror's Greaves (21333, -11.19 DPS, sim-verified) [quest]; Marshal's Plate Boots (16483, -12.80 DPS) [vendor]; General's Plate Boots (16545, -12.80 DPS) [vendor] |
| finger1 | Signet of Unyielding Strength (21393) | Signet of Unyielding Strength [quest] | 208.6 | yes | Band of Earthen Might (21182, -0.68 DPS) [quest]; Band of the Penitent (13217, -1.49 DPS) [quest]; Dragonslayer's Signet (18403, -1.49 DPS) [quest] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 198.8 | yes | Band of Earthen Might (21182, +0.00 DPS, sim-verified) [quest]; Band of the Penitent (13217, -0.92 DPS) [quest]; Dragonslayer's Signet (18403, -0.92 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (300.8 DPS) | yes | Guardian Talisman (1490, -2.42 DPS) [quest]; Ankh of Life (1713, -2.42 DPS) [dungeon]; Blazing Emblem (2802, -2.42 DPS) [dungeon] |
| trinket2 | - | - |  |  |  |
| main_hand | High Warlord's Hacker (235476) | Sergeant Thunderhorn [vendor] | sim-verified (300.8 DPS) | yes | High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Ebon Hand (19170, -14.29 DPS, sim-verified) [crafted] |
| off_hand | High Warlord's Bonecracker (235477) | Sergeant Thunderhorn [vendor] | sim-verified (300.8 DPS) | yes | High Warlord's Left Claw (234558, -0.18 DPS) [vendor]; Grand Marshal's Left Hand Blade (234584, -0.18 DPS) [vendor]; Force Reactive Disk (18168, -85.26 DPS, sim-verified) [crafted] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (300.8 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [vendor]; High Warlord's Crossbow (234560, +0.00 DPS) [vendor]; Dark Iron Rifle (16004, -3.84 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Medallion of the Dawn; shoulder: Lightbreaker Shoulders; back: Drape of Unyielding Strength; chest: Lightbreaker Cuirass; wrist: Lightbreaker Wrists; hands: Lightbreaker Grips; waist: Lightbreaker Belt; legs: Lightbreaker Tassets; feet: Lightbreaker Greaves; finger1: Signet of Unyielding Strength; finger2: Don Julio's Band; trinket2: Onyxia Blood Talisman; main_hand: High Warlord's Hacker; off_hand: High Warlord's Bonecracker; ranged: The Purifier

No-known-source sample (15 of 2506, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

