# Leveling BiS: Fury

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-353000000000000000-000000000000000000)

Set DPS (verified): 32.8. Weights run: 0.9s. Verify run: 0.9s. 247 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.036, strength=2.009 ± 0.045, agility=0.098 ± 0.015, crit=2.375 ± 0.071, hit=not significant (0.000 ± 0.000), melee_haste=1.740 ± 0.046

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.1 | yes | Defender's Leather Hood (252447, -0.38 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.69 DPS) [crafted]; Brawler's Leather Hood (252504, -0.70 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 0.6 | yes | Scholarly Pendant (277203, -0.02 DPS) [quest]; Tarnished Locket (279870, -0.02 DPS) [quest]; Erudite's Amulet (277204, -0.08 DPS, sim-verified) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 | yes | Silvered Bronze Shoulders (3481, +0.00 DPS, sim-verified) [crafted]; Serpent's Shoulders (5404, -0.20 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.22 DPS) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 | yes | Catacomb Cloak (279899, -0.07 DPS) [quest]; Grave Shroud (279865, -0.11 DPS, sim-verified) [quest]; Dark Leather Cloak (2316, -0.13 DPS) [crafted] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.1 | yes | Defender's Leather Armor (252434, -0.21 DPS) [crafted]; Totemic Leather Armor (252435, -0.22 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.34 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 | yes | Cryptwalker Bracers (280095, -0.14 DPS, sim-verified) [quest]; Bravo's Armbands (270015, -0.20 DPS) [quest]; Bristlebark Bindings (14569, -0.21 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | sim-verified (32.8 DPS) | yes | Gold-flecked Gloves (5195, -0.07 DPS) [dungeon]; Polar Gauntlets (7606, -0.15 DPS) [quest]; Fletcher's Gloves (7348, -0.61 DPS, sim-verified) [crafted] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 | yes | Ruffian Belt (5975, -0.22 DPS) [world]; Support Girdle (1215, -0.29 DPS) [world]; Cobrahn's Grasp (6460, -0.32 DPS, sim-verified) [dungeon] |
| legs | Chausses of Westfall (6087) | The Defias Brotherhood [quest] | 22.1 | yes | Defender's Leather Pants (252445, -0.13 DPS) [crafted]; Totemic Leather Pants (252446, -0.15 DPS) [crafted]; Veteran's Chain Leggings (250493, -0.23 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.5 | yes | Guard's Boots (250504, -0.02 DPS) [crafted]; Defender's Leather Boots (252441, -0.02 DPS) [crafted]; Veteran's Boots (250503, -0.04 DPS, sim-verified) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.4 | yes | Loop of Sacrifice (281673, -0.09 DPS) [quest]; The 1 Ring (8350, -0.23 DPS) [world]; Ring of the Moon (12052, -0.23 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 | yes | Loop of Sacrifice (281673, -0.14 DPS, sim-verified) [quest]; The 1 Ring (8350, -0.21 DPS) [world]; Ring of the Moon (12052, -0.22 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 | yes | Living Root (6631, +0.00 DPS) [dungeon]; Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 236.8 | yes | Diamond Hammer (2194, -0.55 DPS, sim-verified) [world_drop]; Redbeard Crest (12997, -8.14 DPS) [world_drop]; Bear Buckler (4821, -8.36 DPS) [vendor] |
| ranged | Dwarven Fishing Pole (3567) (or Cracked Blacksmith Hammer (285279)) | Murloc Poachers [quest] | 4.0 | yes | Cracked Blacksmith Hammer (285279, +0.00 DPS, sim-verified) [crafted]; Fine Longbow (11304, -0.00 DPS) [vendor]; Daryl's Hunting Rifle (2904, -0.07 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Dwarven Fishing Pole

No-known-source sample (15 of 247, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6478 Rat Stompers; 7955 Copper Claymore; 7956 Bronze Warhammer; 9602 Brushwood Blade

### Band 30 (human, 00000000000000000-353211005010000000-000000000000000000)

Set DPS (verified): 53.6. Weights run: 1.0s. Verify run: 1.0s. 422 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.218, strength=2.528 ± 0.266, agility=not significant (0.203 ± 0.078), crit=4.156 ± 0.307, hit=not significant (0.000 ± 0.000), melee_haste=2.879 ± 0.352

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 32.9 | yes | Veteran's Chain Helm (250498, -0.09 DPS, sim-verified) [crafted]; Defender's Leather Helm (252455, -0.09 DPS) [crafted]; Crusader's Chain Helm (250502, -0.18 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | River Pride Choker (13087, -0.14 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.30 DPS, sim-verified) [world_drop]; Sentinel's Medallion (19541, -0.43 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 17.7 | yes | Mail Combat Spaulders (6404, +0.00 DPS, sim-verified) [world_drop]; Barbaric Iron Shoulders (7913, -0.05 DPS) [crafted]; Elite Shoulders (4835, -0.09 DPS) [vendor] |
| back | Sergeant Major's Cape (16315) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 10.9 | yes | Slayer's Cape (14752, -0.03 DPS) [world_drop]; Wolfmaster Cape (6314, -0.03 DPS) [dungeon]; Lambent Scale Cloak (4706, -0.23 DPS, sim-verified) [world_drop] |
| chest | Shining Silver Breastplate (2870) | Blacksmithing [crafted] | 35.4 | yes | Barbaric Iron Breastplate (7914, -0.22 DPS, sim-verified) [crafted]; Hard Gold Cuirass (250533, -0.27 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.31 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 20.2 | yes | Yorgen Bracers (13012, -0.07 DPS, sim-verified) [world_drop]; Bands of Serra'kis (6902, -0.18 DPS) [dungeon]; Patterned Bronze Bracers (2868, -0.27 DPS) [crafted] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | sim-verified (53.6 DPS) | yes | Bonefist Gauntlets (4465, -0.03 DPS) [world]; Mail Combat Gauntlets (4075, -0.09 DPS) [world_drop]; Fletcher's Gloves (7348, -0.82 DPS, sim-verified) [crafted] |
| waist | Girdle of Golem Strength (9405) (or Highlander's Plate Girdle (20126)) | World drop [world_drop] | 30.3 | yes | Highlander's Plate Girdle (20126, +0.00 DPS, sim-verified) [rep]; Highlander's Lamellar Girdle (20108, -0.09 DPS) [rep]; Officer's Belt (250556, -0.13 DPS) [crafted] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 29.2 | yes | Chausses of Westfall (6087, -0.05 DPS) [quest]; Slayer's Pants (14757, -0.05 DPS) [world_drop]; Golden Scale Leggings (3843, -0.48 DPS, sim-verified) [crafted] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 19.1 | yes | Glimmering Mail Greaves (4073, -0.14 DPS) [world_drop]; Slayer's Slippers (14756, -0.14 DPS) [world_drop]; Hard Gold Boots (250534, -0.48 DPS, sim-verified) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.8 | yes | Silverlaine's Family Seal (6321, -0.29 DPS) [dungeon]; Ironspine's Eye (7686, -0.31 DPS) [dungeon]; Demon Band (12054, -0.38 DPS) [world_drop] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 16.4 | yes | Ironspine's Eye (7686, -0.16 DPS) [dungeon]; Protector's Band (20439, -0.19 DPS) [rep]; Silverlaine's Family Seal (6321, -0.32 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 348.1 | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 338.6 | yes | Shoni's Disarming Tool (9608, -4.00 DPS) [quest]; Shield of Thorsen (13079, -11.28 DPS) [world_drop]; Swinetusk Shank (6691, -21.36 DPS, sim-verified) [dungeon] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 | yes | Double-barreled Shotgun (2098, +0.00 DPS, sim-verified) [world_drop]; Long Battle Bow (15284, -0.05 DPS) [world_drop]; Dwarven Fishing Pole (3567, -0.14 DPS) [quest] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Sergeant Major's Cape; chest: Shining Silver Breastplate; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Girdle of Golem Strength; legs: Veteran's Silvered Chain Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Pronged Reaver; off_hand: Ironspine's Fist; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 422, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 40 (human, 00000000000000000-353211005050010050-000000000000000000)

Set DPS (verified): 89.1. Weights run: 1.1s. Verify run: 1.2s. 596 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=not significant (1.000 ± 0.333), strength=1.931 ± 0.393, agility=0.609 ± 0.145, crit=7.890 ± 0.549, hit=not significant (0.000 ± 0.000), melee_haste=3.559 ± 0.547

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 135.6 | yes | Chromite Barbute (8142, -0.50 DPS, sim-verified) [world_drop]; White Bandit Mask (10008, -5.22 DPS) [crafted]; Icemetal Barbute (10763, -5.27 DPS) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; River Pride Choker (13087, -0.30 DPS) [world_drop]; Gazlowe's Charm (13088, -0.30 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 21.2 | yes | Chromite Pauldrons (8144, +0.00 DPS, sim-verified) [world_drop]; Shining Mithril Pauldrons (250541, -0.09 DPS) [crafted]; Sunburn Spaulders (274751, -0.12 DPS) [vendor] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 15.2 | yes | Sergeant Major's Cape (16315, -0.15 DPS, sim-verified) [pvp]; Hawkeye's Cloak (14593, -0.25 DPS) [world_drop]; Wolfmaster Cape (6314, -0.25 DPS) [dungeon] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 31.3 | yes | Shining Silver Breastplate (2870, -0.21 DPS) [crafted]; Golden Scale Cuirass (3845, -0.21 DPS) [crafted]; Jouster's Chestplate (8157, -0.35 DPS, sim-verified) [world_drop] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Ravager's Armguards (14770, +0.00 DPS, sim-verified) [world_drop]; Pugilist Bracers (4438, -0.22 DPS) [dungeon]; Yorgen Bracers (13012, -0.32 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 130.5 | yes | Dragonscale Gauntlets (8347, -0.95 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -0.97 DPS) [crafted]; Ornate Mithril Gloves (7927, -0.97 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20089) | The League of Arathor [rep] | 118.5 | yes | Highlander's Leather Girdle (20116, -0.12 DPS, sim-verified) [rep]; Highlander's Plate Girdle (20125, -4.34 DPS) [rep]; Highlander's Chain Girdle (20090, -4.58 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 40.5 | yes | Firemane Leggings (13129, +0.00 DPS, sim-verified) [world_drop]; Orcish War Leggings (7929, -0.37 DPS) [crafted]; Symbolic Legplates (14829, -0.38 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 29.4 | yes | Prowler's Leather Shoes (252465, +0.00 DPS, sim-verified) [crafted]; Blackforge Greaves (6423, -0.29 DPS) [world_drop]; Obsidian Greaves (13068, -0.31 DPS) [world_drop] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 20.3 | yes | Thunderbrow Ring (13097, -0.15 DPS) [world_drop]; Protector's Band (19517, -0.25 DPS) [rep]; Suspicious Spare Part (274754, -0.33 DPS) [vendor] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 | yes | Thunderbrow Ring (13097, +0.00 DPS, sim-verified) [world_drop]; Suspicious Spare Part (274754, -0.31 DPS) [vendor]; Ironspine's Eye (7686, -0.33 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ardent Custodian (868) | World drop [world_drop] | 460.0 | yes | Bonebiter (6830, +0.00 DPS) [quest]; Illusionary Rod (7713, +0.00 DPS) [dungeon]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | 442.0 | yes | Jhordy's Misplaced Screwdriver (274753, -5.06 DPS, sim-verified) [vendor]; Shoni's Disarming Tool (9608, -10.54 DPS) [quest]; Skullance Shield (13081, -20.58 DPS) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | 14.0 | yes | Master Hunter's Rifle (17687, +0.00 DPS, sim-verified) [quest]; Explosive Shotgun (8188, -0.21 DPS) [world_drop]; Mithril Blacksmith Hammer (285280, -0.21 DPS) [crafted] |

**New at 40:** head: Raging Berserker's Helm; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Chain Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Protector's Band; finger2: Assault Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Ardent Custodian; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 596, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band

### Band 50 (human, 35100000000000000-353211005050010051-000000000000000000)

Set DPS (verified): 93.1. Weights run: 1.2s. Verify run: 1.1s. 775 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.214, strength=1.584 ± 0.298, agility=not significant (0.351 ± 0.116), crit=4.668 ± 0.389, hit=not significant (0.000 ± 0.000), melee_haste=2.879 ± 0.401

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Plate Helm (220804) | Captain Dirgehammer [vendor] | sim-verified (86.2 DPS) | yes | Ornate Mithril Helm (7937, -0.45 DPS) [crafted]; Raging Berserker's Helm (7719, -1.10 DPS, sim-verified) [dungeon]; Eye of Theradras (17715, -1.95 DPS) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Skibi's Pendant (13089, -0.27 DPS, sim-verified) [world_drop]; Kaleidoscope Chain (13084, -0.59 DPS) [world_drop]; River Pride Choker (13087, -0.73 DPS) [world_drop] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) | Captain Dirgehammer [vendor] | 82.8 | yes | Officer's Pauldrons (250576, -0.99 DPS, sim-verified) [crafted]; Wyrmslayer Spaulders (13066, -5.77 DPS) [world_drop]; Earthslag Shoulders (11632, -5.89 DPS) [dungeon] |
| back | Bloodlust Cape (14801) | World drop [world_drop] | 14.3 | yes | Wolfmaster Cape (6314, -0.40 DPS) [dungeon]; Pridelord Cape (14673, -0.60 DPS) [world_drop]; Sergeant Major's Cape (16336, -0.65 DPS, sim-verified) [pvp] |
| chest | Knight's Plate Hauberk (220794) | Captain Dirgehammer [vendor] | 89.1 | yes | Ornate Mithril Breastplate (7935, -3.65 DPS, sim-verified) [crafted]; Warforged Chestplate (11195, -4.84 DPS) [quest]; Valorous Chestguard (8274, -5.44 DPS) [world_drop] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 | yes | Officer's Wristguards (250581, -0.80 DPS) [crafted]; Branded Leather Bracers (19508, -0.96 DPS, sim-verified) [dungeon]; Giantslayer Bracers (13076, -0.99 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 85.3 | yes | Fletcher's Gloves (7348, -1.89 DPS) [crafted]; Ornate Mithril Gloves (7927, -1.89 DPS) [crafted]; Dragonscale Gauntlets (8347, -2.10 DPS, sim-verified) [crafted] |
| waist | Highlander's Chain Girdle (20088) (or Highlander's Leather Girdle (20115)) | The League of Arathor [rep] | 85.3 | yes | Highlander's Leather Girdle (20115, +0.00 DPS, sim-verified) [rep]; Highlander's Lamellar Girdle (20106, -0.24 DPS) [rep]; Highlander's Plate Girdle (20124, -0.39 DPS) [rep] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | sim-verified (86.6 DPS) | yes | Stormshroud Pants (15057, -1.54 DPS, sim-verified) [crafted]; Golem Shard Leggings (13074, -5.02 DPS) [world_drop]; Scarlet Leggings (10330, -5.17 DPS) [dungeon] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 24.5 | yes | Officer's Boots (250546, -0.13 DPS) [crafted]; Officer's Sabatons (250561, -0.19 DPS, sim-verified) [crafted]; Skulker's Leather Boots (252469, -0.23 DPS) [crafted] |
| finger1 | Assault Band (13095) (or Blackstone Ring (17713)) | World drop [world_drop] | 20.0 | yes | Protector's Band (19516, -0.09 DPS) [rep]; Protector's Band (19515, -0.43 DPS) [rep]; Insurgent's Band (272065, -0.47 DPS) [vendor] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 20.0 | yes | Protector's Band (19516, -0.33 DPS, sim-verified) [rep]; Protector's Band (19515, -0.43 DPS) [rep]; Insurgent's Band (272065, -0.47 DPS) [vendor] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (84.8 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Smoking Heart of the Mountain (11811, -1.22 DPS, sim-verified) [crafted] |
| trinket2 | - | - |  |  |  |
| main_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | sim-verified (84.8 DPS) | yes | Kindling Stave (11750, +0.00 DPS) [dungeon]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor] |
| off_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-verified (90.8 DPS) | yes | Inventor's Focal Sword (17719, -5.70 DPS, sim-verified) [dungeon]; Claw of Celebras (17738, -6.74 DPS) [dungeon]; Shoni's Disarming Tool (9608, -31.13 DPS) [quest] |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (84.8 DPS) | yes | Houndmaster's Bow (11628, -0.19 DPS) [dungeon]; The Silencer (13138, -0.45 DPS) [world_drop]; Dark Iron Rifle (16004, -2.55 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Plate Helm; shoulder: Knight-Lieutenant's Plate Pauldrons; back: Bloodlust Cape; chest: Knight's Plate Hauberk; wrist: Bracers of the Stone Princess; waist: Highlander's Chain Girdle; legs: Knight's Plate Leggings; feet: Prowler's Leather Boots; finger1: Assault Band; finger2: Blackstone Ring; trinket1: Frozen Heart of the Mountain; trinket2: Guardian Talisman; main_hand: Dawn's Edge; off_hand: Hammer of the Northern Wind; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 775, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60 (human, 35311103002000000-353211005050010051-000000000000000000)

Set DPS (verified): 318.4. Weights run: 1.3s. Verify run: 1.4s. 1430 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=not significant (1.000 ± 0.427), strength=not significant (2.344 ± 0.596), agility=not significant (0.450 ± 0.239), crit=13.060 ± 0.873, hit=not significant (0.000 ± 0.000), melee_haste=6.782 ± 0.884

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 407.9 | yes | Bloodvine Lens (19998, -2.44 DPS) [crafted]; Lightbreaker Greathelm (239517, -5.70 DPS) [vendor]; Ragefury Eyepatch (11735, -11.51 DPS, sim-verified) [dungeon] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (308.6 DPS) | yes | Blazefury Medallion (17111, -0.99 DPS, sim-verified) [world]; Amulet of the Darkmoon (19491, -10.09 DPS) [quest]; Dragonheart Necklace (20622, -10.56 DPS) [world] |
| shoulder | Lightbreaker Shoulders (239516) | Leonid Barthalomew the Revered [vendor] | 269.3 | yes | Lieutenant Commander's Plate Shoulders (23315, -2.69 DPS) [vendor]; Lieutenant Commander's Plate Shoulders (227045, -2.69 DPS) [vendor]; Darkspear Spaulders (272108, -13.32 DPS, sim-verified) [vendor] |
| back | Cloak of the Honor Guard (20073) | The League of Arathor [rep] | sim-verified (310.6 DPS) | yes | Howler's Furs (272414, -0.48 DPS) [vendor]; Cape of the Black Baron (13340, -0.55 DPS) [dungeon]; Chromatic Cloak (18509, -3.88 DPS, sim-verified) [crafted] |
| chest | Lightbreaker Cuirass (239519) | Leonid Barthalomew the Revered [vendor] | 485.9 | yes | Stormshroud Armor (15056, -6.94 DPS) [crafted]; Savage Gladiator Chain (11726, -9.69 DPS) [dungeon]; Bloodsoul Breastplate (19690, -16.94 DPS, sim-verified) [crafted] |
| wrist | Lightbreaker Wrists (239512) | Leonid Barthalomew the Revered [vendor] | sim-verified (315.8 DPS) | yes | Deeprock Bracers (21184, -1.35 DPS) [quest]; Berserker Bracers (19578, -1.41 DPS) [rep]; Vambraces of the Sadist (13400, -9.06 DPS, sim-verified) [dungeon] |
| hands | Lightbreaker Grips (239514) | Leonid Barthalomew the Revered [vendor] | 268.5 | yes | Marshal's Plate Gauntlets (16484, -2.24 DPS) [vendor]; Marshal's Plate Gauntlets (231541, -2.24 DPS) [vendor]; Chromatic Gauntlets (19157, -5.25 DPS, sim-verified) [crafted] |
| waist | Lightbreaker Belt (239513) | Leonid Barthalomew the Revered [vendor] | 271.7 | yes | Highlander's Plate Girdle (20041, -2.83 DPS) [rep]; Highlander's Lamellar Girdle (20042, -3.10 DPS) [rep]; Radiant Girdle of the Dawn (227814, -4.78 DPS, sim-verified) [vendor] |
| legs | Lightbreaker Tassets (239518) | Leonid Barthalomew the Revered [vendor] | 487.4 | yes | Marshal's Plate Legguards (16479, -4.32 DPS) [vendor]; Marshal's Plate Legguards (231540, -4.32 DPS) [vendor]; Sentinel's Plate Legguards (237825, -5.50 DPS, sim-verified) [vendor] |
| feet | Lightbreaker Greaves (239515) | Leonid Barthalomew the Revered [vendor] | 269.4 | yes | Boots of Heroism (21995, -12.18 DPS, sim-verified) [quest]; Marshal's Plate Boots (16483, -12.80 DPS) [vendor]; Marshal's Plate Boots (231539, -12.80 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (308.6 DPS) | yes | Band of the Penitent (13217, -0.92 DPS) [quest]; Ring of Entropy (18543, -0.92 DPS) [world]; Wrath of Cenarius (21190, -12.25 DPS, sim-verified) [quest] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (308.6 DPS) | yes | Band of the Penitent (13217, -0.81 DPS) [quest]; Ring of Entropy (18543, -0.81 DPS) [world]; Wrath of Cenarius (21190, -11.74 DPS, sim-verified) [quest] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (308.6 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Darkmoon Card: Heroism (19287, -10.65 DPS, sim-verified) [quest] |
| trinket2 | - | - |  |  |  |
| main_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | sim-verified (308.6 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [vendor]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Ebon Hand (19170, -29.79 DPS, sim-verified) [crafted] |
| off_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (308.6 DPS) | yes | Grand Marshal's Left Hand Blade (18847, +0.00 DPS) [vendor]; Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [vendor]; Force Reactive Disk (18168, -79.53 DPS, sim-verified) [crafted] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (308.6 DPS) | yes | Dark Iron Rifle (16004, -4.05 DPS, sim-verified) [crafted]; Bloodseeker (19107, -9.29 DPS) [quest]; Skull Splitting Crossbow (13039, -9.34 DPS) [world_drop] |

**New at 60:** head: Lionheart Helm; neck: Medallion of the Dawn; shoulder: Lightbreaker Shoulders; back: Cloak of the Honor Guard; chest: Lightbreaker Cuirass; wrist: Lightbreaker Wrists; hands: Lightbreaker Grips; waist: Lightbreaker Belt; legs: Lightbreaker Tassets; feet: Lightbreaker Greaves; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Darkmoon Card: Maelstrom; trinket2: Darkmoon Card: Blue Dragon; main_hand: Ravencrest's Legacy; off_hand: Shadowsong's Sorrow; ranged: The Purifier

No-known-source sample (15 of 1430, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band

## Horde

### Band 20 (troll, 00000000000000000-353000000000000000-000000000000000000)

Set DPS (verified): 29.7. Weights run: 0.9s. Verify run: 1.0s. 244 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.036, strength=2.009 ± 0.045, agility=0.098 ± 0.015, crit=2.375 ± 0.071, hit=not significant (0.000 ± 0.000), melee_haste=1.740 ± 0.046

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.1 | yes | Defender's Leather Hood (252447, -0.19 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.69 DPS) [crafted]; Brawler's Leather Hood (252504, -0.70 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 0.6 | yes | Scholarly Pendant (277203, -0.02 DPS) [quest]; Tarnished Locket (279870, -0.02 DPS) [quest]; Erudite's Amulet (277204, -0.10 DPS, sim-verified) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 | yes | Silvered Bronze Shoulders (3481, +0.00 DPS, sim-verified) [crafted]; Serpent's Shoulders (5404, -0.20 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.22 DPS) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 | yes | Grave Shroud (279865, -0.07 DPS, sim-verified) [quest]; Subterranean Cape (14149, -0.07 DPS) [dungeon]; Catacomb Cloak (279899, -0.07 DPS) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.1 | yes | Defender's Leather Armor (252434, -0.21 DPS) [crafted]; Totemic Leather Armor (252435, -0.22 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.25 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 | yes | Cryptwalker Bracers (280095, -0.14 DPS, sim-verified) [quest]; Raptorcrest Bracers (270010, -0.15 DPS) [quest]; Bristlebark Bindings (14569, -0.21 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | sim-verified (29.7 DPS) | yes | Gold-flecked Gloves (5195, -0.07 DPS) [dungeon]; Blackened Defias Gloves (10401, -0.15 DPS) [dungeon]; Fletcher's Gloves (7348, -0.34 DPS, sim-verified) [crafted] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 | yes | Cobrahn's Grasp (6460, -0.15 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.22 DPS) [world]; Support Girdle (1215, -0.29 DPS) [world] |
| legs | Veteran's Chain Leggings (250493) | Blacksmithing [crafted] | 18.6 | yes | Totemic Leather Pants (252446, -0.02 DPS) [crafted]; Defender's Leather Pants (252445, -0.06 DPS, sim-verified) [crafted]; Hulking Leggings (14748, -0.08 DPS) [world_drop] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.5 | yes | Guard's Boots (250504, -0.02 DPS) [crafted]; Defender's Leather Boots (252441, -0.02 DPS) [crafted]; Veteran's Boots (250503, -0.06 DPS, sim-verified) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.4 | yes | Loop of Sacrifice (281673, -0.09 DPS) [quest]; The 1 Ring (8350, -0.23 DPS) [world]; Ring of the Moon (12052, -0.23 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 | yes | Loop of Sacrifice (281673, -0.14 DPS, sim-verified) [quest]; The 1 Ring (8350, -0.21 DPS) [world]; Ring of the Moon (12052, -0.22 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 | yes | Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest]; Hammerbone (270018, +0.00 DPS) [quest] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 236.8 | yes | Diamond Hammer (2194, -0.36 DPS, sim-verified) [world_drop]; Redbeard Crest (12997, -8.14 DPS) [world_drop]; Bear Buckler (4821, -8.36 DPS) [vendor] |
| ranged | Cracked Blacksmith Hammer (285279) | Blacksmithing [crafted] | 4.0 | yes | Fine Longbow (11304, +0.00 DPS, sim-verified) [vendor]; Heavy Shortbow (3036, -0.07 DPS) [world_drop]; Orcish Battle Bow (5346, -0.07 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Demon Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Cracked Blacksmith Hammer

No-known-source sample (15 of 244, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7955 Copper Claymore; 7956 Bronze Warhammer; 9602 Brushwood Blade; 10047 Simple Kilt; 10421 Rough Copper Vest; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 30 (troll, 00000000000000000-353211005010000000-000000000000000000)

Set DPS (verified): 53.5. Weights run: 1.0s. Verify run: 1.0s. 418 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.218, strength=2.528 ± 0.266, agility=not significant (0.203 ± 0.078), crit=4.156 ± 0.307, hit=not significant (0.000 ± 0.000), melee_haste=2.879 ± 0.352

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 32.9 | yes | Defender's Leather Helm (252455, -0.09 DPS) [crafted]; Veteran's Chain Helm (250498, -0.11 DPS, sim-verified) [crafted]; Crusader's Chain Helm (250502, -0.18 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; River Pride Choker (13087, -0.14 DPS) [world_drop]; Scout's Medallion (19537, -0.43 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 17.7 | yes | Mail Combat Spaulders (6404, +0.00 DPS, sim-verified) [world_drop]; Barbaric Iron Shoulders (7913, -0.05 DPS) [crafted]; Elite Shoulders (4835, -0.09 DPS) [vendor] |
| back | Lambent Scale Cloak (4706) (or Slayer's Cape (14752)) | World drop [world_drop] | 10.1 | yes | Slayer's Cape (14752, +0.00 DPS, sim-verified) [world_drop]; Wolfmaster Cape (6314, -0.00 DPS) [dungeon]; Wildhunter Cloak (16658, -0.00 DPS) [quest] |
| chest | Shining Silver Breastplate (2870) | Blacksmithing [crafted] | 35.4 | yes | Barbaric Iron Breastplate (7914, +0.00 DPS, sim-verified) [crafted]; Hard Gold Cuirass (250533, -0.27 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.31 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 20.2 | yes | Yorgen Bracers (13012, +0.00 DPS, sim-verified) [world_drop]; Bands of Serra'kis (6902, -0.18 DPS) [dungeon]; Grimtoll Wristguards (15459, -0.26 DPS) [quest] |
| hands | Warsong Gauntlets (16978) | Warsong Supplies [quest] | sim-verified (53.5 DPS) | yes | Gauntlets of Ogre Strength (3341, -0.06 DPS) [world]; Bonefist Gauntlets (4465, -0.09 DPS) [world]; Fletcher's Gloves (7348, -0.57 DPS, sim-verified) [crafted] |
| waist | Girdle of Golem Strength (9405) (or Defiler's Plate Girdle (20207)) | World drop [world_drop] | 30.3 | yes | Defiler's Plate Girdle (20207, +0.00 DPS, sim-verified) [rep]; Officer's Belt (250556, -0.13 DPS) [crafted]; Defiler's Chain Girdle (20152, -0.22 DPS) [rep] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 29.2 | yes | Slayer's Pants (14757, -0.05 DPS) [world_drop]; Ferine Leggings (6690, -0.11 DPS) [dungeon]; Golden Scale Leggings (3843, -0.28 DPS, sim-verified) [crafted] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 19.1 | yes | Glimmering Mail Greaves (4073, -0.14 DPS) [world_drop]; Slayer's Slippers (14756, -0.14 DPS) [world_drop]; Hard Gold Boots (250534, -0.28 DPS, sim-verified) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.8 | yes | Silverlaine's Family Seal (6321, -0.29 DPS) [dungeon]; Ironspine's Eye (7686, -0.31 DPS) [dungeon]; Band of the Fist (17694, -0.35 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 16.4 | yes | Ironspine's Eye (7686, -0.16 DPS) [dungeon]; Band of the Fist (17694, -0.19 DPS) [quest]; Silverlaine's Family Seal (6321, -0.43 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 348.1 | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 338.6 | yes | Shield of Thorsen (13079, -11.28 DPS) [world_drop]; Slayer's Shield (15892, -11.34 DPS) [world_drop]; Swinetusk Shank (6691, -21.50 DPS, sim-verified) [dungeon] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 | yes | Double-barreled Shotgun (2098, +0.00 DPS, sim-verified) [world_drop]; Long Battle Bow (15284, -0.05 DPS) [world_drop]; Cracked Blacksmith Hammer (285279, -0.14 DPS) [crafted] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; chest: Shining Silver Breastplate; wrist: Pugilist Bracers; hands: Warsong Gauntlets; waist: Girdle of Golem Strength; legs: Veteran's Silvered Chain Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Pronged Reaver; off_hand: Ironspine's Fist; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 418, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7955 Copper Claymore; 7956 Bronze Warhammer; 7957 Bronze Greatsword

### Band 40 (troll, 00000000000000000-353211005050010050-000000000000000000)

Set DPS (verified): 84.6. Weights run: 1.1s. Verify run: 1.2s. 592 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=not significant (1.000 ± 0.333), strength=1.931 ± 0.393, agility=0.609 ± 0.145, crit=7.890 ± 0.549, hit=not significant (0.000 ± 0.000), melee_haste=3.559 ± 0.547

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 135.6 | yes | Chromite Barbute (8142, -1.17 DPS, sim-verified) [world_drop]; White Bandit Mask (10008, -5.22 DPS) [crafted]; Icemetal Barbute (10763, -5.27 DPS) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Kaleidoscope Chain (13084, -0.19 DPS) [world_drop]; River Pride Choker (13087, -0.30 DPS) [world_drop]; Ethereal Talisman (4430, -0.34 DPS, sim-verified) [quest] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 21.2 | yes | Shining Mithril Pauldrons (250541, -0.09 DPS) [crafted]; Sunburn Spaulders (274751, -0.12 DPS) [vendor]; Chromite Pauldrons (8144, -0.62 DPS, sim-verified) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 10.1 | yes | Wildhunter Cloak (16658, -0.00 DPS) [quest]; Lambent Scale Cloak (4706, -0.11 DPS) [world_drop]; Wolfmaster Cape (6314, -0.43 DPS, sim-verified) [dungeon] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 31.3 | yes | Shining Silver Breastplate (2870, -0.21 DPS) [crafted]; Golden Scale Cuirass (3845, -0.21 DPS) [crafted]; Jouster's Chestplate (8157, -0.77 DPS, sim-verified) [world_drop] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Pugilist Bracers (4438, -0.22 DPS) [dungeon]; Darkspear Armsplints (4132, -0.31 DPS) [quest]; Ravager's Armguards (14770, -0.39 DPS, sim-verified) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 130.5 | yes | Dragonscale Gauntlets (8347, -0.88 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -0.97 DPS) [crafted]; Ornate Mithril Gloves (7927, -0.97 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20153) | The Defilers [rep] | 118.5 | yes | Defiler's Leather Girdle (20192, -0.68 DPS, sim-verified) [rep]; Defiler's Plate Girdle (20206, -4.34 DPS) [rep]; Tharg's Shoelace (9705, -4.53 DPS) [quest] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 40.5 | yes | Orcish War Leggings (7929, -0.37 DPS) [crafted]; Symbolic Legplates (14829, -0.38 DPS) [world_drop]; Firemane Leggings (13129, -0.78 DPS, sim-verified) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 29.4 | yes | Blackforge Greaves (6423, -0.29 DPS) [world_drop]; Obsidian Greaves (13068, -0.31 DPS) [world_drop]; Prowler's Leather Shoes (252465, -0.78 DPS, sim-verified) [crafted] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 20.3 | yes | Thunderbrow Ring (13097, -0.15 DPS) [world_drop]; Legionnaire's Band (19513, -0.25 DPS) [rep]; Suspicious Spare Part (274754, -0.33 DPS) [vendor] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 | yes | Suspicious Spare Part (274754, -0.31 DPS) [vendor]; Ironspine's Eye (7686, -0.33 DPS) [dungeon]; Thunderbrow Ring (13097, -0.50 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ardent Custodian (868) | World drop [world_drop] | 460.0 | yes | Illusionary Rod (7713, +0.00 DPS) [dungeon]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 442.0 | yes | Jhordy's Misplaced Screwdriver (274753, +0.00 DPS, sim-verified) [vendor]; Skullance Shield (13081, -20.58 DPS) [world_drop]; Pit Fighter's Shield (4507, -20.70 DPS) [quest] |
| ranged | The Silencer (13138) | World drop [world_drop] | 14.0 | yes | Explosive Shotgun (8188, -0.21 DPS) [world_drop]; Mithril Blacksmith Hammer (285280, -0.21 DPS) [crafted]; Master Hunter's Rifle (17687, -0.39 DPS, sim-verified) [quest] |

**New at 40:** head: Raging Berserker's Helm; shoulder: Hard Gold Pauldrons; back: Hawkeye's Cloak; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Chain Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Legionnaire's Band; finger2: Assault Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Ardent Custodian; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 592, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 50 (troll, 35100000000000000-353211005050010051-000000000000000000)

Set DPS (verified): 99.3. Weights run: 1.2s. Verify run: 1.2s. 771 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.214, strength=1.584 ± 0.298, agility=not significant (0.351 ± 0.116), crit=4.668 ± 0.389, hit=not significant (0.000 ± 0.000), melee_haste=2.879 ± 0.401

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Plate Helm (220803) | Lady Palanseer [vendor] | sim-verified (79.8 DPS) | yes | Ornate Mithril Helm (7937, -0.45 DPS) [crafted]; Raging Berserker's Helm (7719, -1.47 DPS, sim-verified) [dungeon]; Eye of Theradras (17715, -1.95 DPS) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Woven Ivy Necklace (19159, +0.00 DPS, sim-verified) [quest]; Skibi's Pendant (13089, -0.14 DPS) [world_drop]; Ethereal Talisman (4430, -0.44 DPS) [quest] |
| shoulder | Blood Guard's Plate Pauldrons (220796) | Lady Palanseer [vendor] | 82.8 | yes | Officer's Pauldrons (250576, -0.28 DPS, sim-verified) [crafted]; Wyrmslayer Spaulders (13066, -5.77 DPS) [world_drop]; Earthslag Shoulders (11632, -5.89 DPS) [dungeon] |
| back | Bloodlust Cape (14801) | World drop [world_drop] | 14.3 | yes | Battlehard Cape (11858, -0.40 DPS) [quest]; Wildhunter Cloak (16658, -0.40 DPS) [quest]; Wolfmaster Cape (6314, -0.70 DPS, sim-verified) [dungeon] |
| chest | Stone Guard's Plate Armor (220801) | Lady Palanseer [vendor] | 89.1 | yes | Ornate Mithril Breastplate (7935, -2.60 DPS, sim-verified) [crafted]; Warforged Chestplate (11195, -4.84 DPS) [quest]; Valorous Chestguard (8274, -5.44 DPS) [world_drop] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 | yes | Branded Leather Bracers (19508, -0.70 DPS, sim-verified) [dungeon]; Officer's Wristguards (250581, -0.80 DPS) [crafted]; Giantslayer Bracers (13076, -0.99 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 85.3 | yes | Dragonscale Gauntlets (8347, -0.71 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -1.89 DPS) [crafted]; Ornate Mithril Gloves (7927, -1.89 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20151) (or Defiler's Leather Girdle (20193)) | The Defilers [rep] | 85.3 | yes | Defiler's Leather Girdle (20193, +0.00 DPS, sim-verified) [rep]; Defiler's Plate Girdle (20205, -0.39 DPS) [rep]; Defiler's Chain Girdle (20153, -1.14 DPS) [rep] |
| legs | Stone Guard's Plate Leggings (220798) | Lady Palanseer [vendor] | sim-verified (80.0 DPS) | yes | Stormshroud Pants (15057, -1.63 DPS, sim-verified) [crafted]; Golem Shard Leggings (13074, -5.02 DPS) [world_drop]; Scarlet Leggings (10330, -5.17 DPS) [dungeon] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 24.5 | yes | Officer's Sabatons (250561, -0.06 DPS, sim-verified) [crafted]; Officer's Boots (250546, -0.13 DPS) [crafted]; Skulker's Leather Boots (252469, -0.23 DPS) [crafted] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 | yes | Legionnaire's Band (19511, -0.47 DPS) [rep]; Legionnaire's Band (19512, -0.81 DPS) [rep]; Assault Band (13095, -1.56 DPS, sim-verified) [world_drop] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | sim-verified (79.8 DPS) | yes | Legionnaire's Band (19511, -0.09 DPS) [rep]; Legionnaire's Band (19512, -0.43 DPS) [rep]; Assault Band (13095, -1.47 DPS, sim-verified) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (77.6 DPS) | yes | Tidal Charm (1404, -3.98 DPS) [vendor]; Guardian Talisman (1490, -3.98 DPS) [quest]; Smoking Heart of the Mountain (11811, -4.53 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (78.6 DPS) | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest]; Smoking Heart of the Mountain (11811, -0.67 DPS, sim-verified) [crafted] |
| main_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | sim-verified (78.6 DPS) | yes | Kindling Stave (11750, +0.00 DPS) [dungeon]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor] |
| off_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-verified (95.7 DPS) | yes | Claw of Celebras (17738, -6.74 DPS) [dungeon]; White Bone Shredder (11863, -10.37 DPS) [quest]; Inventor's Focal Sword (17719, -17.39 DPS, sim-verified) [dungeon] |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (78.6 DPS) | yes | Houndmaster's Bow (11628, -0.19 DPS) [dungeon]; The Silencer (13138, -0.45 DPS) [world_drop]; Dark Iron Rifle (16004, -1.08 DPS, sim-verified) [crafted] |

**New at 50:** head: Blood Guard's Plate Helm; shoulder: Blood Guard's Plate Pauldrons; back: Bloodlust Cape; chest: Stone Guard's Plate Armor; wrist: Bracers of the Stone Princess; waist: Defiler's Chain Girdle; legs: Stone Guard's Plate Leggings; feet: Prowler's Leather Boots; finger1: White Bone Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Dawn's Edge; off_hand: Hammer of the Northern Wind; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 771, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (troll, 35311103002000000-353211005050010051-000000000000000000)

Set DPS (verified): 309.2. Weights run: 1.3s. Verify run: 1.4s. 1427 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=not significant (1.000 ± 0.427), strength=not significant (2.344 ± 0.596), agility=not significant (0.450 ± 0.239), crit=13.060 ± 0.873, hit=not significant (0.000 ± 0.000), melee_haste=6.782 ± 0.884

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 407.9 | yes | Bloodvine Lens (19998, -2.44 DPS) [crafted]; Lightbreaker Greathelm (239517, -5.70 DPS) [vendor]; Ragefury Eyepatch (11735, -11.34 DPS, sim-verified) [dungeon] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (297.8 DPS) | yes | Blazefury Medallion (17111, -3.21 DPS, sim-verified) [world]; Amulet of the Darkmoon (19491, -10.09 DPS) [quest]; Dragonheart Necklace (20622, -10.56 DPS) [world] |
| shoulder | Lightbreaker Shoulders (239516) | Leonid Barthalomew the Revered [vendor] | 269.3 | yes | Champion's Plate Shoulders (23243, -2.69 DPS) [vendor]; Champion's Plate Shoulders (227042, -2.69 DPS) [vendor]; Darkspear Spaulders (272108, -13.84 DPS, sim-verified) [vendor] |
| back | Deathguard's Cloak (20068) | The Defilers [rep] | sim-verified (298.5 DPS) | yes | Howler's Furs (272414, -0.48 DPS) [vendor]; Cape of the Black Baron (13340, -0.55 DPS) [dungeon]; Chromatic Cloak (18509, -3.91 DPS, sim-verified) [crafted] |
| chest | Lightbreaker Cuirass (239519) | Leonid Barthalomew the Revered [vendor] | 485.9 | yes | Stormshroud Armor (15056, -6.94 DPS) [crafted]; Savage Gladiator Chain (11726, -9.69 DPS) [dungeon]; Bloodsoul Breastplate (19690, -16.88 DPS, sim-verified) [crafted] |
| wrist | Lightbreaker Wrists (239512) | Leonid Barthalomew the Revered [vendor] | sim-verified (305.8 DPS) | yes | Deeprock Bracers (21184, -1.35 DPS) [quest]; Berserker Bracers (19578, -1.41 DPS) [rep]; Vambraces of the Sadist (13400, -11.27 DPS, sim-verified) [dungeon] |
| hands | Lightbreaker Grips (239514) | Leonid Barthalomew the Revered [vendor] | 268.5 | yes | General's Plate Gauntlets (16548, -2.24 DPS) [vendor]; General's Plate Gauntlets (231532, -2.24 DPS) [vendor]; Chromatic Gauntlets (19157, -6.23 DPS, sim-verified) [crafted] |
| waist | Lightbreaker Belt (239513) | Leonid Barthalomew the Revered [vendor] | 271.7 | yes | Defiler's Plate Girdle (20204, -2.83 DPS) [rep]; Defiler's Chain Girdle (20150, -3.17 DPS) [rep]; Radiant Girdle of the Dawn (227814, -4.43 DPS, sim-verified) [vendor] |
| legs | Lightbreaker Tassets (239518) | Leonid Barthalomew the Revered [vendor] | 487.4 | yes | General's Plate Leggings (16543, -4.32 DPS) [vendor]; General's Plate Leggings (231533, -4.32 DPS) [vendor]; Sentinel's Plate Legguards (237825, -6.59 DPS, sim-verified) [vendor] |
| feet | Lightbreaker Greaves (239515) | Leonid Barthalomew the Revered [vendor] | 269.4 | yes | Boots of Heroism (21995, -10.42 DPS, sim-verified) [quest]; General's Plate Boots (16545, -12.80 DPS) [vendor]; General's Plate Boots (231531, -12.80 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (297.8 DPS) | yes | Band of the Penitent (13217, -0.92 DPS) [quest]; Ring of Entropy (18543, -0.92 DPS) [world]; Wrath of Cenarius (21190, -10.26 DPS, sim-verified) [quest] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (297.8 DPS) | yes | Band of the Penitent (13217, -0.81 DPS) [quest]; Ring of Entropy (18543, -0.81 DPS) [world]; Wrath of Cenarius (21190, -9.10 DPS, sim-verified) [quest] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (288.0 DPS) | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (297.8 DPS) | yes | Tidal Charm (1404, -2.42 DPS) [vendor]; Guardian Talisman (1490, -2.42 DPS) [quest]; Darkmoon Card: Heroism (19287, -8.85 DPS, sim-verified) [quest] |
| main_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | sim-verified (297.8 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [vendor]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Ebon Hand (19170, -8.36 DPS, sim-verified) [crafted] |
| off_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (299.4 DPS) | yes | High Warlord's Left Claw (18848, +0.00 DPS) [vendor]; High Warlord's Left Claw (234558, +0.00 DPS) [vendor]; Force Reactive Disk (18168, -83.05 DPS, sim-verified) [crafted] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (299.4 DPS) | yes | Dark Iron Rifle (16004, -3.05 DPS, sim-verified) [crafted]; Bloodseeker (19107, -9.29 DPS) [quest]; Skull Splitting Crossbow (13039, -9.34 DPS) [world_drop] |

**New at 60:** head: Lionheart Helm; neck: Medallion of the Dawn; shoulder: Lightbreaker Shoulders; back: Deathguard's Cloak; chest: Lightbreaker Cuirass; wrist: Lightbreaker Wrists; hands: Lightbreaker Grips; waist: Lightbreaker Belt; legs: Lightbreaker Tassets; feet: Lightbreaker Greaves; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Ravencrest's Legacy; off_hand: Shadowsong's Sorrow; ranged: The Purifier

No-known-source sample (15 of 1427, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

