# Leveling BiS: Fury

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-353000000000000000-000000000000000000)

Set DPS (verified): 32.8. Weights run: 1.1s. Verify run: 1.3s. 261 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.036, strength=2.009 ± 0.045, agility=0.098 ± 0.015, crit=2.375 ± 0.071, hit=0.483 ± 0.020, melee_haste=1.740 ± 0.046

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.1 attack_power points (0.73 DPS) | yes | Defender's Leather Hood (252447, -0.38 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.69 DPS) [crafted]; Brawler's Leather Hood (252504, -0.70 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 0.6 attack_power points (0.02 DPS) | yes | Erudite's Amulet (277204, -0.08 DPS, sim-verified) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.22 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS, sim-verified) [crafted]; Serpent's Shoulders (5404, -0.20 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.29 DPS) | yes | Catacomb Cloak (279899, -0.07 DPS) [quest]; Grave Shroud (279865, -0.11 DPS, sim-verified) [quest]; Dark Leather Cloak (2316, -0.13 DPS) [crafted] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.1 attack_power points (0.73 DPS) | yes | Defender's Leather Armor (252434, -0.21 DPS) [crafted]; Totemic Leather Armor (252435, -0.22 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.34 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.36 DPS) | yes | Cryptwalker Bracers (280095, -0.14 DPS, sim-verified) [quest]; Bravo's Armbands (270015, -0.20 DPS) [quest]; Bristlebark Bindings (14569, -0.21 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | sim-verified (32.8 DPS) | yes | Gold-flecked Gloves (5195, -0.07 DPS) [dungeon]; Polar Gauntlets (7606, -0.15 DPS) [quest]; Fletcher's Gloves (7348, -0.61 DPS, sim-verified) [crafted] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (0.65 DPS) | yes | Ruffian Belt (5975, -0.22 DPS) [world]; Support Girdle (1215, -0.29 DPS) [world]; Cobrahn's Grasp (6460, -0.32 DPS, sim-verified) [dungeon] |
| legs | Chausses of Westfall (6087) | The Defias Brotherhood [quest] | 22.1 attack_power points (0.80 DPS) | yes | Defender's Leather Pants (252445, -0.13 DPS) [crafted]; Totemic Leather Pants (252446, -0.15 DPS) [crafted]; Veteran's Chain Leggings (250493, -0.23 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.5 attack_power points (0.38 DPS) | yes | Guard's Boots (250504, -0.02 DPS) [crafted]; Defender's Leather Boots (252441, -0.02 DPS) [crafted]; Veteran's Boots (250503, -0.04 DPS, sim-verified) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.4 attack_power points (0.31 DPS) | yes | Loop of Sacrifice (281673, -0.09 DPS) [quest]; The 1 Ring (8350, -0.23 DPS) [world]; Ring of the Moon (12052, -0.23 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.29 DPS) | yes | Loop of Sacrifice (281673, -0.14 DPS, sim-verified) [quest]; The 1 Ring (8350, -0.21 DPS) [world]; Ring of the Moon (12052, -0.22 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (9.02 DPS) | yes | Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon]; The Axe of Severing (23171, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 236.8 attack_power points (8.58 DPS) | yes | Diamond Hammer (2194, -0.55 DPS, sim-verified) [world_drop]; Redbeard Crest (12997, -8.14 DPS) [world_drop]; Bear Buckler (4821, -8.36 DPS) [vendor] |
| ranged | Dwarven Fishing Pole (3567) (or Cracked Blacksmith Hammer (285279)) | Murloc Poachers [quest] | 4.0 attack_power points (0.15 DPS) | yes | Cracked Blacksmith Hammer (285279, +0.00 DPS, sim-verified) [crafted]; Fine Longbow (11304, -0.00 DPS) [vendor]; Daryl's Hunting Rifle (2904, -0.07 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Dwarven Fishing Pole

No-known-source sample (15 of 261, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings

### Band 30 (human, 00000000000000000-353211005010000000-000000000000000000)

Set DPS (verified): 53.6. Weights run: 1.2s. Verify run: 1.3s. 444 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.218, strength=2.528 ± 0.266, agility=not significant (0.203 ± 0.078), crit=4.156 ± 0.307, hit=0.852 ± 0.038, melee_haste=2.879 ± 0.352

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 32.9 attack_power points (1.16 DPS) | yes | Defender's Leather Helm (252455, -0.09 DPS) [crafted]; Veteran's Chain Helm (250498, -0.12 DPS, sim-verified) [crafted]; Crusader's Chain Helm (250502, -0.18 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.49 DPS) | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; River Pride Choker (13087, -0.14 DPS) [world_drop]; Sentinel's Medallion (19541, -0.43 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 17.7 attack_power points (0.62 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS, sim-verified) [world_drop]; Barbaric Iron Shoulders (7913, -0.05 DPS) [crafted]; Elite Shoulders (4835, -0.09 DPS) [vendor] |
| back | Sergeant Major's Cape (16315) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 10.9 attack_power points (0.38 DPS) | yes | Slayer's Cape (14752, -0.03 DPS) [world_drop]; Wolfmaster Cape (6314, -0.03 DPS) [dungeon]; Lambent Scale Cloak (4706, -0.17 DPS, sim-verified) [world_drop] |
| chest | Shining Silver Breastplate (2870) | Blacksmithing [crafted] | sim-verified (52.8 DPS) | yes | Barbaric Iron Breastplate (7914, -0.18 DPS) [crafted]; Hard Gold Cuirass (250533, -0.27 DPS) [crafted]; Avenger's Armor (1488, -2.30 DPS, sim-verified) [dungeon] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 20.2 attack_power points (0.71 DPS) | yes | Yorgen Bracers (13012, +0.00 DPS, sim-verified) [world_drop]; Bands of Serra'kis (6902, -0.18 DPS) [dungeon]; Patterned Bronze Bracers (2868, -0.27 DPS) [crafted] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | sim-verified (51.5 DPS) | yes | Bonefist Gauntlets (4465, -0.03 DPS) [world]; Mail Combat Gauntlets (4075, -0.09 DPS) [world_drop]; Fletcher's Gloves (7348, -1.00 DPS, sim-verified) [crafted] |
| waist | Girdle of Golem Strength (9405) (or Highlander's Plate Girdle (20126)) | World drop [world_drop] | 30.3 attack_power points (1.07 DPS) | yes | Highlander's Plate Girdle (20126, +0.00 DPS, sim-verified) [rep]; Highlander's Lamellar Girdle (20108, -0.09 DPS) [rep]; Officer's Belt (250556, -0.13 DPS) [crafted] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 29.2 attack_power points (1.03 DPS) | yes | Chausses of Westfall (6087, -0.05 DPS) [quest]; Slayer's Pants (14757, -0.05 DPS) [world_drop]; Golden Scale Leggings (3843, -0.15 DPS, sim-verified) [crafted] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 19.1 attack_power points (0.67 DPS) | yes | Glimmering Mail Greaves (4073, -0.14 DPS) [world_drop]; Slayer's Slippers (14756, -0.14 DPS) [world_drop]; Hard Gold Boots (250534, -0.15 DPS, sim-verified) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.8 attack_power points (0.73 DPS) | yes | Tiger Band (6749, -0.20 DPS) [quest]; Silverlaine's Family Seal (6321, -0.29 DPS) [dungeon]; Ironspine's Eye (7686, -0.31 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 16.4 attack_power points (0.58 DPS) | yes | Silverlaine's Family Seal (6321, -0.13 DPS) [dungeon]; Ironspine's Eye (7686, -0.16 DPS) [dungeon]; Tiger Band (6749, -0.90 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 348.1 attack_power points (12.23 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 338.6 attack_power points (11.90 DPS) | yes | Shoni's Disarming Tool (9608, -4.00 DPS) [quest]; Shield of Thorsen (13079, -11.28 DPS) [world_drop]; Royal Diplomatic Scepter (9457, -13.13 DPS, sim-verified) [dungeon] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.32 DPS) | yes | Double-barreled Shotgun (2098, +0.00 DPS, sim-verified) [world_drop]; Long Battle Bow (15284, -0.05 DPS) [world_drop]; Dwarven Fishing Pole (3567, -0.14 DPS) [quest] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Sergeant Major's Cape; chest: Shining Silver Breastplate; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Girdle of Golem Strength; legs: Veteran's Silvered Chain Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Pronged Reaver; off_hand: Ironspine's Fist; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 444, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow

### Band 40 (human, 00000000000000000-353211005050010050-000000000000000000)

Set DPS (verified): 67.6. Weights run: 1.4s. Verify run: 1.4s. 628 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.333, strength=1.931 ± 0.393, agility=0.609 ± 0.145, crit=7.890 ± 0.549, hit=1.030 ± 0.052, melee_haste=3.559 ± 0.547

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 135.6 attack_power points (6.58 DPS) | yes | Chromite Barbute (8142, -0.69 DPS, sim-verified) [world_drop]; White Bandit Mask (10008, -5.22 DPS) [crafted]; Icemetal Barbute (10763, -5.27 DPS) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.97 DPS) | yes | Ghostshard Talisman (7731, +0.00 DPS, sim-verified) [dungeon]; Kaleidoscope Chain (13084, -0.48 DPS) [world_drop]; River Pride Choker (13087, -0.60 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 21.2 attack_power points (1.03 DPS) | yes | Shining Mithril Pauldrons (250541, -0.09 DPS) [crafted]; Sunburn Spaulders (274751, -0.12 DPS) [vendor]; Chromite Pauldrons (8144, -0.18 DPS, sim-verified) [world_drop] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 15.2 attack_power points (0.74 DPS) | yes | Sergeant Major's Cape (16315, -0.25 DPS) [pvp]; Hawkeye's Cloak (14593, -0.25 DPS) [world_drop]; Dark Hooded Cape (5257, -1.18 DPS, sim-verified) [world] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 31.3 attack_power points (1.52 DPS) | yes | Jouster's Chestplate (8157, -0.11 DPS) [world_drop]; Shining Silver Breastplate (2870, -0.21 DPS) [crafted]; Avenger's Armor (1488, -2.38 DPS, sim-verified) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.97 DPS) | yes | Ravager's Armguards (14770, -0.18 DPS, sim-verified) [world_drop]; Pugilist Bracers (4438, -0.22 DPS) [dungeon]; Yorgen Bracers (13012, -0.32 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 130.5 attack_power points (6.33 DPS) | yes | Fletcher's Gloves (7348, -0.97 DPS) [crafted]; Ornate Mithril Gloves (7927, -0.97 DPS) [crafted]; Dragonscale Gauntlets (8347, -1.35 DPS, sim-verified) [crafted] |
| waist | Highlander's Chain Girdle (20089) | The League of Arathor [rep] | 118.5 attack_power points (5.75 DPS) | yes | Highlander's Leather Girdle (20116, +0.00 DPS, sim-verified) [rep]; Boar Champion's Belt (10768, -4.34 DPS) [dungeon]; Highlander's Plate Girdle (20125, -4.34 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 40.5 attack_power points (1.97 DPS) | yes | Firemane Leggings (13129, -0.35 DPS, sim-verified) [world_drop]; Orcish War Leggings (7929, -0.37 DPS) [crafted]; Symbolic Legplates (14829, -0.38 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 29.4 attack_power points (1.42 DPS) | yes | Blackforge Greaves (6423, -0.29 DPS) [world_drop]; Obsidian Greaves (13068, -0.31 DPS) [world_drop]; Prowler's Leather Shoes (252465, -0.35 DPS, sim-verified) [crafted] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 20.3 attack_power points (0.99 DPS) | yes | Thunderbrow Ring (13097, -0.15 DPS) [world_drop]; Protector's Band (19517, -0.25 DPS) [rep]; Mark of Kern (2262, -1.76 DPS, sim-verified) [dungeon] |
| finger2 | Assault Band (13095) | World drop [world_drop] | sim-verified (60.3 DPS) | yes | Thunderbrow Ring (13097, -0.13 DPS) [world_drop]; Suspicious Spare Part (274754, -0.31 DPS) [vendor]; Mark of Kern (2262, -0.92 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (23.02 DPS) | yes | Ardent Custodian (868, +0.00 DPS, sim-verified) [world_drop]; Nightblade (1982, +0.00 DPS) [world_drop]; Illusionary Rod (7713, +0.00 DPS) [dungeon] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | sim-verified (65.9 DPS) | yes | Ardent Custodian (868, -6.49 DPS, sim-verified) [world_drop]; Shoni's Disarming Tool (9608, -10.54 DPS) [quest]; Savage Boar's Guard (10767, -20.41 DPS) [dungeon] |
| ranged | The Silencer (13138) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Monolithic Bow (9426, -0.03 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.19 DPS) [quest]; Bow of Searing Arrows (2825, -0.95 DPS, sim-verified) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Chain Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Protector's Band; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 628, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 50 (human, 35100000000000000-353211005050010051-000000000000000000)

Set DPS (verified): 91.9. Weights run: 1.5s. Verify run: 1.5s. 820 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.214, strength=1.584 ± 0.298, agility=not significant (0.351 ± 0.116), crit=4.668 ± 0.389, hit=0.999 ± 0.045, melee_haste=2.879 ± 0.401

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Plate Helm (220804) | Captain Dirgehammer [vendor] | 95.9 attack_power points (9.09 DPS) | yes | Raging Berserker's Helm (7719, -0.74 DPS, sim-verified) [dungeon]; Ornate Mithril Helm (7937, -1.40 DPS) [crafted]; Eye of Theradras (17715, -2.90 DPS) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | sim-verified (77.6 DPS) | yes | Skibi's Pendant (13089, -0.14 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.59 DPS) [world_drop]; Zealous Shadowshard Pendant (17772, -1.73 DPS, sim-verified) [quest] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) | Captain Dirgehammer [vendor] | 82.8 attack_power points (7.84 DPS) | yes | Officer's Pauldrons (250576, -0.14 DPS, sim-verified) [crafted]; Wyrmslayer Spaulders (13066, -5.77 DPS) [world_drop]; Earthslag Shoulders (11632, -5.89 DPS) [dungeon] |
| back | Bloodlust Cape (14801) | World drop [world_drop] | sim-verified (77.9 DPS) | yes | Sergeant Major's Cape (16336, -0.25 DPS) [pvp]; Wolfmaster Cape (6314, -0.40 DPS) [dungeon]; Blackveil Cape (11626, -2.00 DPS, sim-verified) [dungeon] |
| chest | Knight's Plate Hauberk (220794) | Captain Dirgehammer [vendor] | 89.1 attack_power points (8.44 DPS) | yes | Ornate Mithril Breastplate (7935, -2.42 DPS, sim-verified) [crafted]; Warforged Chestplate (11195, -4.84 DPS) [quest]; Grizzled Pelt (22274, -5.36 DPS) [quest] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 attack_power points (2.65 DPS) | yes | Branded Leather Bracers (19508, -0.76 DPS) [dungeon]; Officer's Wristguards (250581, -0.80 DPS) [crafted]; Runed Golem Shackles (12550, -2.41 DPS, sim-verified) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 85.3 attack_power points (8.09 DPS) | yes | Dragonscale Gauntlets (8347, -1.30 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -1.89 DPS) [crafted]; Ornate Mithril Gloves (7927, -1.89 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20088) (or Highlander's Leather Girdle (20115)) | The League of Arathor [rep] | 85.3 attack_power points (8.09 DPS) | yes | Highlander's Leather Girdle (20115, +0.00 DPS, sim-verified) [rep]; Highlander's Lamellar Girdle (20106, -0.24 DPS) [rep]; Highlander's Plate Girdle (20124, -0.39 DPS) [rep] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | sim-verified (77.9 DPS) | yes | Stormshroud Pants (15057, -2.00 DPS, sim-verified) [crafted]; Golem Shard Leggings (13074, -5.02 DPS) [world_drop]; Scarlet Leggings (10330, -5.17 DPS) [dungeon] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | sim-verified (79.0 DPS) | yes | Officer's Sabatons (250561, -0.07 DPS) [crafted]; Officer's Boots (250546, -0.13 DPS) [crafted]; Battlechaser's Greaves (12555, -3.08 DPS, sim-verified) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 30.0 attack_power points (2.84 DPS) | yes | Protector's Band (19516, -1.04 DPS) [rep]; Protector's Band (19515, -1.37 DPS) [rep]; Mark of Kern (2262, -2.90 DPS, sim-verified) [dungeon] |
| finger2 | Assault Band (13095) | World drop [world_drop] | sim-verified (78.0 DPS) | yes | Protector's Band (19516, -0.09 DPS) [rep]; Protector's Band (19515, -0.43 DPS) [rep]; Mark of Kern (2262, -2.13 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Thorium Greatmace (250613, +0.00 DPS) [crafted] |
| off_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-verified (83.5 DPS) | yes | Claw of Celebras (17738, -6.74 DPS) [dungeon]; Inventor's Focal Sword (17719, -7.64 DPS, sim-verified) [dungeon]; Shoni's Disarming Tool (9608, -31.13 DPS) [quest] |
| ranged | Stinging Bow (10624) | Sunken Temple: Atal'ai Warrior [dungeon] | 0.0 attack_power points (0.00 DPS) | yes | Bow of Searing Arrows (2825, +0.00 DPS, sim-verified) [world_drop]; Skull Splitting Crossbow (13039, +0.00 DPS) [world_drop]; Houndmaster's Bow (11628, -0.19 DPS) [dungeon] |

**New at 50:** head: Knight-Lieutenant's Plate Helm; neck: Ghostshard Talisman; shoulder: Knight-Lieutenant's Plate Pauldrons; back: Bloodlust Cape; chest: Knight's Plate Hauberk; wrist: Bracers of the Stone Princess; waist: Highlander's Chain Girdle; legs: Knight's Plate Leggings; feet: Prowler's Leather Boots; finger1: Blackstone Ring; trinket1: Frozen Heart of the Mountain; main_hand: Dawn's Edge; off_hand: Hammer of the Northern Wind; ranged: Stinging Bow

No-known-source sample (15 of 820, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 60 (human, 35311103002000000-353211005050010051-000000000000000000)

Set DPS (verified): 283.9. Weights run: 1.6s. Verify run: 1.5s. 1737 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.427, strength=not significant (2.344 ± 0.596), agility=not significant (0.450 ± 0.239), crit=13.060 ± 0.873, hit=2.273 ± 0.103, melee_haste=6.782 ± 0.884

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 453.3 attack_power points (26.17 DPS) | yes | Bloodvine Lens (19998, -5.06 DPS) [crafted]; Lightbreaker Greathelm (239517, -7.01 DPS) [vendor]; Ragefury Eyepatch (11735, -11.39 DPS, sim-verified) [dungeon] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (283.9 DPS) | yes | Rage of Mugamba (19577, -1.86 DPS, sim-verified) [quest]; Beads of Ogre Might (22150, -9.24 DPS) [quest]; Mark of Fordring (15411, -9.69 DPS) [quest] |
| shoulder | Lightbreaker Shoulders (239516) | Leonid Barthalomew the Revered [vendor] | 292.0 attack_power points (16.86 DPS) | yes | Lieutenant Commander's Plate Shoulders (23315, -4.00 DPS) [vendor]; Lieutenant Commander's Plate Shoulders (227045, -4.00 DPS) [vendor]; Darkspear Spaulders (272108, -14.79 DPS, sim-verified) [vendor] |
| back | Chromatic Cloak (18509) (or Fel Cape (279269)) | Leatherworking [crafted] | 182.8 attack_power points (10.56 DPS) | yes | Fel Cape (279269, +0.00 DPS, sim-verified) [crafted]; Howler's Furs (272414, -7.63 DPS) [vendor]; Puissant Cape (18541, -8.12 DPS) [world] |
| chest | Lightbreaker Cuirass (239519) | Leonid Barthalomew the Revered [vendor] | sim-verified (283.9 DPS) | yes | Bloodsoul Breastplate (19690, -6.71 DPS) [crafted]; Stormshroud Armor (15056, -6.94 DPS) [crafted]; Breastplate of Undead Slaying (23087, -25.25 DPS, sim-verified) [world] |
| wrist | Vambraces of the Sadist (13400) | Stratholme: Timmy the Cruel [dungeon] | sim-verified (283.9 DPS) | yes | Bracers of Undead Slaying (23090, -5.53 DPS, sim-verified) [world]; Lightbreaker Wrists (239512, -6.01 DPS) [vendor]; Deeprock Bracers (21184, -8.67 DPS) [quest] |
| hands | Lightbreaker Grips (239514) | Leonid Barthalomew the Revered [vendor] | sim-verified (283.9 DPS) | yes | Marshal's Plate Gauntlets (16484, -2.24 DPS) [vendor]; Marshal's Plate Gauntlets (231541, -2.24 DPS) [vendor]; Razor Gauntlets (18326, -15.51 DPS, sim-verified) [dungeon] |
| waist | Lightbreaker Belt (239513) | Leonid Barthalomew the Revered [vendor] | 271.7 attack_power points (15.69 DPS) | yes | Radiant Girdle of the Dawn (227814, -2.70 DPS, sim-verified) [vendor]; Highlander's Plate Girdle (20041, -2.83 DPS) [rep]; Highlander's Lamellar Girdle (20042, -3.10 DPS) [rep] |
| legs | Lightbreaker Tassets (239518) | Leonid Barthalomew the Revered [vendor] | sim-verified (283.9 DPS) | yes | Sentinel's Plate Legguards (237825, -2.06 DPS) [vendor]; Marshal's Plate Legguards (16479, -3.01 DPS) [vendor]; Cloudkeeper Legplates (14554, -25.61 DPS, sim-verified) [world_drop] |
| feet | Lightbreaker Greaves (239515) | Leonid Barthalomew the Revered [vendor] | 292.1 attack_power points (16.86 DPS) | yes | Boots of Heroism (21995, -10.39 DPS, sim-verified) [quest]; Marshal's Plate Boots (16483, -12.80 DPS) [vendor]; Marshal's Plate Boots (231539, -12.80 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (283.9 DPS) | yes | Band of the Penitent (13217, -2.24 DPS) [quest]; Ring of Entropy (18543, -2.24 DPS) [world]; Naglering (11669, -12.00 DPS, sim-verified) [dungeon] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (283.9 DPS) | yes | Band of the Penitent (13217, -2.12 DPS) [quest]; Ring of Entropy (18543, -2.12 DPS) [world]; Naglering (11669, -10.49 DPS, sim-verified) [dungeon] |
| trinket1 | Earthstrike (21180) | Champion's Battlegear [quest] | sim-verified (283.9 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (283.9 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -7.71 DPS, sim-verified) [crafted] |
| main_hand | Ebon Hand (19170) | Blacksmithing [crafted] | sim-verified (283.9 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [vendor]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Misplaced Servo Arm (23221, -81.66 DPS, sim-verified) [world_drop] |
| off_hand | Persuader (22384) | Blacksmithing [crafted] | sim-verified (283.9 DPS) | yes | Grand Marshal's Left Hand Blade (18847, +0.00 DPS) [vendor]; Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [vendor]; Eskhandar's Left Claw (18202, -76.05 DPS, sim-verified) [world] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (283.9 DPS) | yes | Bow of Searing Arrows (2825, -3.14 DPS, sim-verified) [world_drop]; Polished Ironwood Crossbow (20599, -9.17 DPS) [world]; Core Marksman Rifle (18282, -9.24 DPS) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Medallion of the Dawn; shoulder: Lightbreaker Shoulders; back: Chromatic Cloak; chest: Lightbreaker Cuirass; wrist: Vambraces of the Sadist; hands: Lightbreaker Grips; waist: Lightbreaker Belt; legs: Lightbreaker Tassets; feet: Lightbreaker Greaves; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Earthstrike; trinket2: Darkmoon Card: Maelstrom; main_hand: Ebon Hand; off_hand: Persuader; ranged: The Purifier

No-known-source sample (15 of 1737, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

## Horde

### Band 20 (troll, 00000000000000000-353000000000000000-000000000000000000)

Set DPS (verified): 29.7. Weights run: 1.1s. Verify run: 1.3s. 258 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.036, strength=2.009 ± 0.045, agility=0.098 ± 0.015, crit=2.375 ± 0.071, hit=0.483 ± 0.020, melee_haste=1.740 ± 0.046

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.1 attack_power points (0.73 DPS) | yes | Defender's Leather Hood (252447, -0.19 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.69 DPS) [crafted]; Brawler's Leather Hood (252504, -0.70 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 0.6 attack_power points (0.02 DPS) | yes | Erudite's Amulet (277204, -0.10 DPS, sim-verified) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.22 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS, sim-verified) [crafted]; Serpent's Shoulders (5404, -0.20 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.29 DPS) | yes | Grave Shroud (279865, -0.07 DPS, sim-verified) [quest]; Subterranean Cape (14149, -0.07 DPS) [dungeon]; Catacomb Cloak (279899, -0.07 DPS) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.1 attack_power points (0.73 DPS) | yes | Defender's Leather Armor (252434, -0.21 DPS) [crafted]; Totemic Leather Armor (252435, -0.22 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.25 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.36 DPS) | yes | Cryptwalker Bracers (280095, -0.14 DPS, sim-verified) [quest]; Raptorcrest Bracers (270010, -0.15 DPS) [quest]; Bristlebark Bindings (14569, -0.21 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | sim-verified (29.7 DPS) | yes | Gold-flecked Gloves (5195, -0.07 DPS) [dungeon]; Blackened Defias Gloves (10401, -0.15 DPS) [dungeon]; Fletcher's Gloves (7348, -0.34 DPS, sim-verified) [crafted] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (0.65 DPS) | yes | Cobrahn's Grasp (6460, -0.15 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.22 DPS) [world]; Support Girdle (1215, -0.29 DPS) [world] |
| legs | Veteran's Chain Leggings (250493) | Blacksmithing [crafted] | 18.6 attack_power points (0.67 DPS) | yes | Totemic Leather Pants (252446, -0.02 DPS) [crafted]; Defender's Leather Pants (252445, -0.06 DPS, sim-verified) [crafted]; Hulking Leggings (14748, -0.08 DPS) [world_drop] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.5 attack_power points (0.38 DPS) | yes | Guard's Boots (250504, -0.02 DPS) [crafted]; Defender's Leather Boots (252441, -0.02 DPS) [crafted]; Veteran's Boots (250503, -0.06 DPS, sim-verified) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.4 attack_power points (0.31 DPS) | yes | Loop of Sacrifice (281673, -0.09 DPS) [quest]; The 1 Ring (8350, -0.23 DPS) [world]; Ring of the Moon (12052, -0.23 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.29 DPS) | yes | Loop of Sacrifice (281673, -0.14 DPS, sim-verified) [quest]; The 1 Ring (8350, -0.21 DPS) [world]; Ring of the Moon (12052, -0.22 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (9.02 DPS) | yes | The Axe of Severing (23171, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest]; Hammerbone (270018, +0.00 DPS) [quest] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 236.8 attack_power points (8.58 DPS) | yes | Diamond Hammer (2194, -0.36 DPS, sim-verified) [world_drop]; Redbeard Crest (12997, -8.14 DPS) [world_drop]; Bear Buckler (4821, -8.36 DPS) [vendor] |
| ranged | Cracked Blacksmith Hammer (285279) | Blacksmithing [crafted] | 4.0 attack_power points (0.15 DPS) | yes | Fine Longbow (11304, +0.00 DPS, sim-verified) [vendor]; Heavy Shortbow (3036, -0.07 DPS) [world_drop]; Orcish Battle Bow (5346, -0.07 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Cracked Blacksmith Hammer

No-known-source sample (15 of 258, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 30 (troll, 00000000000000000-353211005010000000-000000000000000000)

Set DPS (verified): 53.5. Weights run: 1.2s. Verify run: 1.3s. 438 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.218, strength=2.528 ± 0.266, agility=not significant (0.203 ± 0.078), crit=4.156 ± 0.307, hit=0.852 ± 0.038, melee_haste=2.879 ± 0.352

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 32.9 attack_power points (1.16 DPS) | yes | Defender's Leather Helm (252455, -0.09 DPS) [crafted]; Veteran's Chain Helm (250498, -0.16 DPS, sim-verified) [crafted]; Crusader's Chain Helm (250502, -0.18 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.49 DPS) | yes | River Pride Choker (13087, -0.14 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.35 DPS, sim-verified) [world_drop]; Scout's Medallion (19537, -0.43 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 17.7 attack_power points (0.62 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS, sim-verified) [world_drop]; Barbaric Iron Shoulders (7913, -0.05 DPS) [crafted]; Elite Shoulders (4835, -0.09 DPS) [vendor] |
| back | Lambent Scale Cloak (4706) (or Slayer's Cape (14752)) | World drop [world_drop] | 10.1 attack_power points (0.36 DPS) | yes | Slayer's Cape (14752, +0.00 DPS, sim-verified) [world_drop]; Wolfmaster Cape (6314, -0.00 DPS) [dungeon]; Wildhunter Cloak (16658, -0.00 DPS) [quest] |
| chest | Shining Silver Breastplate (2870) | Blacksmithing [crafted] | sim-verified (53.0 DPS) | yes | Barbaric Iron Breastplate (7914, -0.18 DPS) [crafted]; Hard Gold Cuirass (250533, -0.27 DPS) [crafted]; Avenger's Armor (1488, -1.78 DPS, sim-verified) [dungeon] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 20.2 attack_power points (0.71 DPS) | yes | Bands of Serra'kis (6902, -0.18 DPS) [dungeon]; Yorgen Bracers (13012, -0.22 DPS, sim-verified) [world_drop]; Grimtoll Wristguards (15459, -0.26 DPS) [quest] |
| hands | Warsong Gauntlets (16978) | Warsong Supplies [quest] | sim-verified (51.9 DPS) | yes | Gauntlets of Ogre Strength (3341, -0.06 DPS) [world]; Bonefist Gauntlets (4465, -0.09 DPS) [world]; Fletcher's Gloves (7348, -0.69 DPS, sim-verified) [crafted] |
| waist | Girdle of Golem Strength (9405) (or Defiler's Plate Girdle (20207)) | World drop [world_drop] | 30.3 attack_power points (1.07 DPS) | yes | Defiler's Plate Girdle (20207, +0.00 DPS, sim-verified) [rep]; Officer's Belt (250556, -0.13 DPS) [crafted]; Defiler's Chain Girdle (20152, -0.22 DPS) [rep] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 29.2 attack_power points (1.03 DPS) | yes | Slayer's Pants (14757, -0.05 DPS) [world_drop]; Ferine Leggings (6690, -0.11 DPS) [dungeon]; Golden Scale Leggings (3843, -0.30 DPS, sim-verified) [crafted] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 19.1 attack_power points (0.67 DPS) | yes | Glimmering Mail Greaves (4073, -0.14 DPS) [world_drop]; Slayer's Slippers (14756, -0.14 DPS) [world_drop]; Hard Gold Boots (250534, -0.30 DPS, sim-verified) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.8 attack_power points (0.73 DPS) | yes | Tiger Band (6749, -0.20 DPS) [quest]; Silverlaine's Family Seal (6321, -0.29 DPS) [dungeon]; Ironspine's Eye (7686, -0.31 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 16.4 attack_power points (0.58 DPS) | yes | Silverlaine's Family Seal (6321, -0.13 DPS) [dungeon]; Ironspine's Eye (7686, -0.16 DPS) [dungeon]; Tiger Band (6749, -1.23 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 348.1 attack_power points (12.23 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 338.6 attack_power points (11.90 DPS) | yes | Shield of Thorsen (13079, -11.28 DPS) [world_drop]; Slayer's Shield (15892, -11.34 DPS) [world_drop]; Royal Diplomatic Scepter (9457, -13.64 DPS, sim-verified) [dungeon] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.32 DPS) | yes | Long Battle Bow (15284, -0.05 DPS) [world_drop]; Cracked Blacksmith Hammer (285279, -0.14 DPS) [crafted]; Double-barreled Shotgun (2098, -0.28 DPS, sim-verified) [world_drop] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; chest: Shining Silver Breastplate; wrist: Pugilist Bracers; hands: Warsong Gauntlets; waist: Girdle of Golem Strength; legs: Veteran's Silvered Chain Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Pronged Reaver; off_hand: Ironspine's Fist; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 438, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 40 (troll, 00000000000000000-353211005050010050-000000000000000000)

Set DPS (verified): 61.5. Weights run: 1.4s. Verify run: 1.5s. 620 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.333, strength=1.931 ± 0.393, agility=0.609 ± 0.145, crit=7.890 ± 0.549, hit=1.030 ± 0.052, melee_haste=3.559 ± 0.547

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 135.6 attack_power points (6.58 DPS) | yes | Chromite Barbute (8142, -0.59 DPS, sim-verified) [world_drop]; White Bandit Mask (10008, -5.22 DPS) [crafted]; Icemetal Barbute (10763, -5.27 DPS) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | sim-verified (59.4 DPS) | yes | Ethereal Talisman (4430, -0.09 DPS) [quest]; Kaleidoscope Chain (13084, -0.19 DPS) [world_drop]; Zealous Shadowshard Pendant (17772, -0.98 DPS, sim-verified) [quest] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 21.2 attack_power points (1.03 DPS) | yes | Shining Mithril Pauldrons (250541, -0.09 DPS) [crafted]; Sunburn Spaulders (274751, -0.12 DPS) [vendor]; Chromite Pauldrons (8144, -0.13 DPS, sim-verified) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-verified (59.3 DPS) | yes | Wolfmaster Cape (6314, -0.00 DPS) [dungeon]; Wildhunter Cloak (16658, -0.00 DPS) [quest]; Dark Hooded Cape (5257, -0.94 DPS, sim-verified) [world] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 31.3 attack_power points (1.52 DPS) | yes | Jouster's Chestplate (8157, -0.11 DPS) [world_drop]; Shining Silver Breastplate (2870, -0.21 DPS) [crafted]; Avenger's Armor (1488, -2.15 DPS, sim-verified) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.97 DPS) | yes | Ravager's Armguards (14770, -0.04 DPS, sim-verified) [world_drop]; Pugilist Bracers (4438, -0.22 DPS) [dungeon]; Darkspear Armsplints (4132, -0.31 DPS) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 130.5 attack_power points (6.33 DPS) | yes | Dragonscale Gauntlets (8347, -0.67 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -0.97 DPS) [crafted]; Ornate Mithril Gloves (7927, -0.97 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20153) | The Defilers [rep] | 118.5 attack_power points (5.75 DPS) | yes | Defiler's Leather Girdle (20192, -0.01 DPS, sim-verified) [rep]; Boar Champion's Belt (10768, -4.34 DPS) [dungeon]; Defiler's Plate Girdle (20206, -4.34 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 40.5 attack_power points (1.97 DPS) | yes | Firemane Leggings (13129, -0.03 DPS, sim-verified) [world_drop]; Orcish War Leggings (7929, -0.37 DPS) [crafted]; Symbolic Legplates (14829, -0.38 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 29.4 attack_power points (1.42 DPS) | yes | Prowler's Leather Shoes (252465, -0.03 DPS, sim-verified) [crafted]; Blackforge Greaves (6423, -0.29 DPS) [world_drop]; Obsidian Greaves (13068, -0.31 DPS) [world_drop] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 20.3 attack_power points (0.99 DPS) | yes | Thunderbrow Ring (13097, -0.15 DPS) [world_drop]; Legionnaire's Band (19513, -0.25 DPS) [rep]; Mark of Kern (2262, -1.37 DPS, sim-verified) [dungeon] |
| finger2 | Assault Band (13095) | World drop [world_drop] | sim-verified (59.6 DPS) | yes | Thunderbrow Ring (13097, -0.13 DPS) [world_drop]; Suspicious Spare Part (274754, -0.31 DPS) [vendor]; Mark of Kern (2262, -1.26 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (23.02 DPS) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Illusionary Rod (7713, +0.00 DPS) [dungeon]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor] |
| off_hand | Ardent Custodian (868) | World drop [world_drop] | 460.0 attack_power points (22.32 DPS) | yes | Vanquisher's Sword (10823, -1.20 DPS, sim-verified) [quest]; Savage Boar's Guard (10767, -21.29 DPS) [dungeon]; Skullance Shield (13081, -21.45 DPS) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Monolithic Bow (9426, -0.03 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.19 DPS) [quest]; Bow of Searing Arrows (2825, -0.67 DPS, sim-verified) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; shoulder: Hard Gold Pauldrons; back: Hawkeye's Cloak; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Chain Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Legionnaire's Band; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Ardent Custodian; ranged: The Silencer

No-known-source sample (15 of 620, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (troll, 35100000000000000-353211005050010051-000000000000000000)

Set DPS (verified): 97.8. Weights run: 1.5s. Verify run: 1.6s. 812 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.214, strength=1.584 ± 0.298, agility=not significant (0.351 ± 0.116), crit=4.668 ± 0.389, hit=0.999 ± 0.045, melee_haste=2.879 ± 0.401

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Plate Helm (220803) | Lady Palanseer [vendor] | 95.9 attack_power points (9.09 DPS) | yes | Raging Berserker's Helm (7719, -0.85 DPS, sim-verified) [dungeon]; Ornate Mithril Helm (7937, -1.40 DPS) [crafted]; Eye of Theradras (17715, -2.90 DPS) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | sim-verified (75.0 DPS) | yes | Woven Ivy Necklace (19159, -0.13 DPS) [quest]; Skibi's Pendant (13089, -0.14 DPS) [world_drop]; Zealous Shadowshard Pendant (17772, -1.77 DPS, sim-verified) [quest] |
| shoulder | Blood Guard's Plate Pauldrons (220796) | Lady Palanseer [vendor] | 82.8 attack_power points (7.84 DPS) | yes | Officer's Pauldrons (250576, +0.00 DPS, sim-verified) [crafted]; Wyrmslayer Spaulders (13066, -5.77 DPS) [world_drop]; Earthslag Shoulders (11632, -5.89 DPS) [dungeon] |
| back | Bloodlust Cape (14801) | World drop [world_drop] | sim-verified (75.3 DPS) | yes | Wolfmaster Cape (6314, -0.40 DPS) [dungeon]; Battlehard Cape (11858, -0.40 DPS) [quest]; Blackveil Cape (11626, -2.02 DPS, sim-verified) [dungeon] |
| chest | Stone Guard's Plate Armor (220801) | Lady Palanseer [vendor] | 89.1 attack_power points (8.44 DPS) | yes | Ornate Mithril Breastplate (7935, -2.88 DPS, sim-verified) [crafted]; Warforged Chestplate (11195, -4.84 DPS) [quest]; Grizzled Pelt (22274, -5.36 DPS) [quest] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 attack_power points (2.65 DPS) | yes | Branded Leather Bracers (19508, -0.76 DPS) [dungeon]; Officer's Wristguards (250581, -0.80 DPS) [crafted]; Runed Golem Shackles (12550, -2.48 DPS, sim-verified) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 85.3 attack_power points (8.09 DPS) | yes | Dragonscale Gauntlets (8347, -1.00 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -1.89 DPS) [crafted]; Ornate Mithril Gloves (7927, -1.89 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20151) (or Defiler's Leather Girdle (20193)) | The Defilers [rep] | 85.3 attack_power points (8.09 DPS) | yes | Defiler's Leather Girdle (20193, +0.00 DPS, sim-verified) [rep]; Defiler's Plate Girdle (20205, -0.39 DPS) [rep]; Defiler's Chain Girdle (20153, -1.14 DPS) [rep] |
| legs | Stone Guard's Plate Leggings (220798) | Lady Palanseer [vendor] | sim-verified (75.3 DPS) | yes | Stormshroud Pants (15057, -2.01 DPS, sim-verified) [crafted]; Golem Shard Leggings (13074, -5.02 DPS) [world_drop]; Scarlet Leggings (10330, -5.17 DPS) [dungeon] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | sim-verified (76.8 DPS) | yes | Officer's Sabatons (250561, -0.07 DPS) [crafted]; Officer's Boots (250546, -0.13 DPS) [crafted]; Battlechaser's Greaves (12555, -3.50 DPS, sim-verified) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 30.0 attack_power points (2.84 DPS) | yes | Mark of Kern (2262, -0.95 DPS) [dungeon]; Assault Band (13095, -0.95 DPS) [world_drop]; Legionnaire's Band (19511, -1.04 DPS) [rep] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (2.27 DPS) | yes | Assault Band (13095, -0.38 DPS) [world_drop]; Legionnaire's Band (19511, -0.47 DPS) [rep]; Mark of Kern (2262, -2.13 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+4.2 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 attack_power points (0.00 DPS) | yes | Smoking Heart of the Mountain (11811, -0.44 DPS, sim-verified) [crafted] |
| main_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Thorium Greatmace (250613, +0.00 DPS) [crafted] |
| off_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-verified (90.5 DPS) | yes | Claw of Celebras (17738, -6.74 DPS) [dungeon]; White Bone Shredder (11863, -10.37 DPS) [quest]; Inventor's Focal Sword (17719, -17.21 DPS, sim-verified) [dungeon] |
| ranged | Stinging Bow (10624) | Sunken Temple: Atal'ai Warrior [dungeon] | 0.0 attack_power points (0.00 DPS) | yes | Bow of Searing Arrows (2825, +0.00 DPS, sim-verified) [world_drop]; Skull Splitting Crossbow (13039, +0.00 DPS) [world_drop]; Houndmaster's Bow (11628, -0.19 DPS) [dungeon] |

**New at 50:** head: Blood Guard's Plate Helm; shoulder: Blood Guard's Plate Pauldrons; back: Bloodlust Cape; chest: Stone Guard's Plate Armor; wrist: Bracers of the Stone Princess; waist: Defiler's Chain Girdle; legs: Stone Guard's Plate Leggings; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Dawn's Edge; off_hand: Hammer of the Northern Wind; ranged: Stinging Bow

No-known-source sample (15 of 812, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (troll, 35311103002000000-353211005050010051-000000000000000000)

Set DPS (verified): 286.2. Weights run: 1.6s. Verify run: 1.5s. 1729 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.427, strength=not significant (2.344 ± 0.596), agility=not significant (0.450 ± 0.239), crit=13.060 ± 0.873, hit=2.273 ± 0.103, melee_haste=6.782 ± 0.884

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 453.3 attack_power points (26.17 DPS) | yes | Bloodvine Lens (19998, -5.06 DPS) [crafted]; Lightbreaker Greathelm (239517, -7.01 DPS) [vendor]; Ragefury Eyepatch (11735, -11.59 DPS, sim-verified) [dungeon] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (286.2 DPS) | yes | Blazefury Medallion (17111, -1.57 DPS, sim-verified) [world]; Beads of Ogre Might (22150, -9.24 DPS) [quest]; Mark of Fordring (15411, -9.69 DPS) [quest] |
| shoulder | Lightbreaker Shoulders (239516) | Leonid Barthalomew the Revered [vendor] | 292.0 attack_power points (16.86 DPS) | yes | Champion's Plate Shoulders (23243, -4.00 DPS) [vendor]; Champion's Plate Shoulders (227042, -4.00 DPS) [vendor]; Darkspear Spaulders (272108, -13.83 DPS, sim-verified) [vendor] |
| back | Chromatic Cloak (18509) (or Fel Cape (279269)) | Leatherworking [crafted] | 182.8 attack_power points (10.56 DPS) | yes | Fel Cape (279269, +0.00 DPS, sim-verified) [crafted]; Howler's Furs (272414, -7.63 DPS) [vendor]; Puissant Cape (18541, -8.12 DPS) [world] |
| chest | Lightbreaker Cuirass (239519) | Leonid Barthalomew the Revered [vendor] | sim-verified (286.2 DPS) | yes | Bloodsoul Breastplate (19690, -6.71 DPS) [crafted]; Stormshroud Armor (15056, -6.94 DPS) [crafted]; Breastplate of Undead Slaying (23087, -26.81 DPS, sim-verified) [world] |
| wrist | Vambraces of the Sadist (13400) | Stratholme: Timmy the Cruel [dungeon] | sim-verified (286.2 DPS) | yes | Lightbreaker Wrists (239512, -6.01 DPS) [vendor]; Bracers of Undead Slaying (23090, -6.49 DPS, sim-verified) [world]; Deeprock Bracers (21184, -8.67 DPS) [quest] |
| hands | Lightbreaker Grips (239514) | Leonid Barthalomew the Revered [vendor] | sim-verified (286.2 DPS) | yes | General's Plate Gauntlets (16548, -2.24 DPS) [vendor]; General's Plate Gauntlets (231532, -2.24 DPS) [vendor]; Razor Gauntlets (18326, -16.03 DPS, sim-verified) [dungeon] |
| waist | Lightbreaker Belt (239513) | Leonid Barthalomew the Revered [vendor] | 271.7 attack_power points (15.69 DPS) | yes | Defiler's Plate Girdle (20204, -2.83 DPS) [rep]; Defiler's Chain Girdle (20150, -3.17 DPS) [rep]; Radiant Girdle of the Dawn (227814, -3.73 DPS, sim-verified) [vendor] |
| legs | Lightbreaker Tassets (239518) | Leonid Barthalomew the Revered [vendor] | sim-verified (286.2 DPS) | yes | Sentinel's Plate Legguards (237825, -2.06 DPS) [vendor]; General's Plate Leggings (16543, -3.01 DPS) [vendor]; Cloudkeeper Legplates (14554, -26.75 DPS, sim-verified) [world_drop] |
| feet | Lightbreaker Greaves (239515) | Leonid Barthalomew the Revered [vendor] | 292.1 attack_power points (16.86 DPS) | yes | Boots of Heroism (21995, -9.75 DPS, sim-verified) [quest]; General's Plate Boots (16545, -12.80 DPS) [vendor]; General's Plate Boots (231531, -12.80 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (286.2 DPS) | yes | Band of the Penitent (13217, -2.24 DPS) [quest]; Ring of Entropy (18543, -2.24 DPS) [world]; Naglering (11669, -10.73 DPS, sim-verified) [dungeon] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (286.2 DPS) | yes | Band of the Penitent (13217, -2.12 DPS) [quest]; Ring of Entropy (18543, -2.12 DPS) [world]; Naglering (11669, -9.39 DPS, sim-verified) [dungeon] |
| trinket1 | Earthstrike (21180) | Champion's Battlegear [quest] | sim-verified (286.2 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Darkmoon Card: Maelstrom (19289, -3.39 DPS, sim-verified) [quest] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (286.2 DPS) | yes | Darkmoon Card: Maelstrom (19289, +0.00 DPS, sim-verified) [quest]; Blackhand's Breadth (13965, -1.84 DPS) [quest]; Counterattack Lodestone (18537, -2.07 DPS) [dungeon] |
| main_hand | Ebon Hand (19170) | Blacksmithing [crafted] | sim-verified (286.2 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [vendor]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Misplaced Servo Arm (23221, -81.05 DPS, sim-verified) [world_drop] |
| off_hand | Persuader (22384) | Blacksmithing [crafted] | sim-verified (286.2 DPS) | yes | High Warlord's Left Claw (18848, +0.00 DPS) [vendor]; High Warlord's Left Claw (234558, +0.00 DPS) [vendor]; Eskhandar's Left Claw (18202, -74.69 DPS, sim-verified) [world] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (286.2 DPS) | yes | Bow of Searing Arrows (2825, -3.53 DPS, sim-verified) [world_drop]; Polished Ironwood Crossbow (20599, -9.17 DPS) [world]; Core Marksman Rifle (18282, -9.24 DPS) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Medallion of the Dawn; shoulder: Lightbreaker Shoulders; back: Chromatic Cloak; chest: Lightbreaker Cuirass; wrist: Vambraces of the Sadist; hands: Lightbreaker Grips; waist: Lightbreaker Belt; legs: Lightbreaker Tassets; feet: Lightbreaker Greaves; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Earthstrike; trinket2: Rune of the Guard Captain; main_hand: Ebon Hand; off_hand: Persuader; ranged: The Purifier

No-known-source sample (15 of 1729, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

