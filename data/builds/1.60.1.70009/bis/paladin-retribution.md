# Leveling BiS: Retribution

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-0000000000000000-55100000000000000)

Set DPS (verified): 34.4. Weights run: 1.0s. Verify run: 0.9s. 220 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.133 ± 0.018, crit=2.410 ± 0.076, hit=1.869 ± 0.027, melee_haste=-2.419 ± 0.186

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.69 DPS) | yes | Defender's Leather Hood (252447, -0.16 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.65 DPS) [crafted]; Brawler's Leather Hood (252504, -0.66 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 0.8 attack_power points (0.03 DPS) | yes | Erudite's Amulet (277204, -0.04 DPS, sim-verified) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.21 DPS) | yes | Silvered Bronze Shoulders (3481, -0.03 DPS, sim-verified) [crafted]; Serpent's Shoulders (5404, -0.18 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Catacomb Cloak (279899, -0.07 DPS) [quest]; Grave Shroud (279865, -0.10 DPS, sim-verified) [quest]; Dark Leather Cloak (2316, -0.12 DPS) [crafted] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.69 DPS) | yes | Veteran's Chain Shirt (250488, -0.18 DPS, sim-verified) [crafted]; Defender's Leather Armor (252434, -0.19 DPS) [crafted]; Totemic Leather Armor (252435, -0.21 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Cryptwalker Bracers (280095, -0.08 DPS, sim-verified) [quest]; Bravo's Armbands (270015, -0.19 DPS) [quest]; Bristlebark Bindings (14569, -0.19 DPS) [world_drop] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 33.7 attack_power points (1.17 DPS) | yes | Thorbia's Gauntlets (12994, +0.00 DPS, sim-verified) [world_drop]; Gold-flecked Gloves (5195, -0.68 DPS) [dungeon]; Polar Gauntlets (7606, -0.75 DPS) [quest] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (0.62 DPS) | yes | Cobrahn's Grasp (6460, -0.12 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.21 DPS) [world]; Support Girdle (1215, -0.28 DPS) [world] |
| legs | Chausses of Westfall (6087) | The Defias Brotherhood [quest] | 22.0 attack_power points (0.76 DPS) | yes | Veteran's Chain Leggings (250493, -0.08 DPS, sim-verified) [crafted]; Defender's Leather Pants (252445, -0.12 DPS) [crafted]; Totemic Leather Pants (252446, -0.14 DPS) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.7 attack_power points (0.37 DPS) | yes | Veteran's Boots (250503, -0.02 DPS, sim-verified) [crafted]; Guard's Boots (250504, -0.02 DPS) [crafted]; Defender's Leather Boots (252441, -0.02 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.5 attack_power points (0.30 DPS) | yes | Loop of Sacrifice (281673, -0.09 DPS) [quest]; The 1 Ring (8350, -0.22 DPS) [world]; Ring of the Moon (12052, -0.23 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Loop of Sacrifice (281673, +0.00 DPS, sim-verified) [quest]; The 1 Ring (8350, -0.20 DPS) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Verigan's Fist (6953) | The Test of Righteousness [quest] | 358.7 attack_power points (12.43 DPS) | yes | Forsaken Greataxe (251533, -1.93 DPS) [quest]; Smite's Mighty Hammer (7230, -2.09 DPS) [dungeon]; The Axe of Severing (23171, -20.13 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Verigan's Fist; ranged: Libram of Banishment

No-known-source sample (15 of 220, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5031 ZZZZZZZZ; 5036 ZZZZZ; 5255 Quilboar Tomahawk

### Band 30 (human, 00000000000000000-0000000000000000-55223310000000000)

Set DPS (verified): 71.6. Weights run: 1.1s. Verify run: 1.1s. 384 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.118 ± 0.014, crit=3.235 ± 0.092, hit=2.259 ± 0.035, melee_haste=1.778 ± 0.112

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 attack_power points (1.08 DPS) | yes | Defender's Leather Helm (252455, -0.08 DPS) [crafted]; Veteran's Chain Helm (250498, -0.14 DPS, sim-verified) [crafted]; Crusader's Chain Helm (250502, -0.17 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.58 DPS) | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; River Pride Choker (13087, -0.25 DPS) [world_drop]; Sentinel's Medallion (19541, -0.54 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 14.0 attack_power points (0.58 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS, sim-verified) [world_drop]; Barbaric Iron Shoulders (7913, -0.05 DPS) [crafted]; Elite Shoulders (4835, -0.08 DPS) [vendor] |
| back | Sergeant Major's Cape (16315) | PvP rank 9 · Sergeant Major · Alliance [pvp] | sim-verified (67.9 DPS) | yes | Lambent Scale Cloak (4706, -0.02 DPS) [world_drop]; Slayer's Cape (14752, -0.02 DPS) [world_drop]; Wolfmaster Cape (6314, -0.70 DPS, sim-verified) [dungeon] |
| chest | Shining Silver Breastplate (2870) | Blacksmithing [crafted] | sim-verified (69.1 DPS) | yes | Barbaric Iron Breastplate (7914, -0.17 DPS) [crafted]; Hard Gold Cuirass (250533, -0.25 DPS) [crafted]; Avenger's Armor (1488, -1.95 DPS, sim-verified) [dungeon] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.67 DPS) | yes | Yorgen Bracers (13012, -0.17 DPS, sim-verified) [world_drop]; Bands of Serra'kis (6902, -0.17 DPS) [dungeon]; Patterned Bronze Bracers (2868, -0.25 DPS) [crafted] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | sim-verified (67.9 DPS) | yes | Bonefist Gauntlets (4465, -0.17 DPS) [world]; Mail Combat Gauntlets (4075, -0.23 DPS) [world_drop]; Fletcher's Gloves (7348, -0.75 DPS, sim-verified) [crafted] |
| waist | Girdle of Golem Strength (9405) (or Highlander's Plate Girdle (20126)) | World drop [world_drop] | 24.0 attack_power points (1.00 DPS) | yes | Highlander's Plate Girdle (20126, +0.00 DPS, sim-verified) [rep]; Highlander's Chain Girdle (20090, -0.00 DPS) [rep]; Highlander's Leather Girdle (20117, -0.00 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.08 DPS) | yes | Golden Scale Leggings (3843, -0.17 DPS) [crafted]; Chausses of Westfall (6087, -0.17 DPS) [quest]; Veteran's Silvered Chain Leggings (250523, -2.31 DPS, sim-verified) [crafted] |
| feet | Hard Gold Boots (250534) | Blacksmithing [crafted] | sim-verified (68.2 DPS) | yes | Glimmering Mail Greaves (4073, -0.08 DPS) [world_drop]; Slayer's Slippers (14756, -0.08 DPS) [world_drop]; Trouncing Boots (4464, -1.01 DPS, sim-verified) [world] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.4 attack_power points (0.68 DPS) | yes | Tiger Band (6749, -0.18 DPS) [quest]; Silverlaine's Family Seal (6321, -0.26 DPS) [dungeon]; Ironspine's Eye (7686, -0.30 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 12.7 attack_power points (0.53 DPS) | yes | Silverlaine's Family Seal (6321, -0.11 DPS) [dungeon]; Ironspine's Eye (7686, -0.15 DPS) [dungeon]; Tiger Band (6749, -1.08 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (67.1 DPS) | yes | Talisman of Arathor (21119, -2.95 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (67.1 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Corpsemaker (6687, -0.14 DPS) [dungeon]; Viscous Hammer (13045, -15.54 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Sergeant Major's Cape; chest: Shining Silver Breastplate; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Girdle of Golem Strength; legs: Ferine Leggings; feet: Hard Gold Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 384, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants

### Band 40 (human, 00000000000000000-0000000000000000-55223331211000210)

Set DPS (verified): 74.8. Weights run: 1.1s. Verify run: 1.1s. 546 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.228 ± 0.031, crit=4.287 ± 0.122, hit=3.119 ± 0.049, melee_haste=2.170 ± 0.198

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 86.0 attack_power points (3.65 DPS) | yes | Icemetal Barbute (10763, -0.21 DPS, sim-verified) [dungeon]; Hard Gold Coif (250537, -2.46 DPS) [crafted]; Chromite Barbute (8142, -2.52 DPS) [world_drop] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.59 DPS) | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; River Pride Choker (13087, -0.25 DPS) [world_drop]; Gazlowe's Charm (13088, -0.25 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 attack_power points (0.93 DPS) | yes | Shining Mithril Pauldrons (250541, -0.08 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.17 DPS) [world_drop]; Chromite Pauldrons (8144, -0.65 DPS, sim-verified) [world_drop] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 13.4 attack_power points (0.57 DPS) | yes | Wolfmaster Cape (6314, -0.14 DPS) [dungeon]; Sergeant Major's Cape (16315, -0.19 DPS) [pvp]; Dark Hooded Cape (5257, -1.31 DPS, sim-verified) [world] |
| chest | Jouster's Chestplate (8157) | World drop [world_drop] | sim-verified (74.8 DPS) | yes | Kolkar Marauder Chain (6773, -0.02 DPS) [quest]; Shining Silver Breastplate (2870, -0.08 DPS) [crafted]; Avenger's Armor (1488, -1.60 DPS, sim-verified) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.85 DPS) | yes | Pugilist Bracers (4438, -0.21 DPS, sim-verified) [dungeon]; Ravager's Armguards (14770, -0.22 DPS) [world_drop]; Yorgen Bracers (13012, -0.31 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 80.0 attack_power points (3.40 DPS) | yes | Fletcher's Gloves (7348, -0.85 DPS) [crafted]; Ornate Mithril Gloves (7927, -0.85 DPS) [crafted]; Dragonscale Gauntlets (8347, -0.86 DPS, sim-verified) [crafted] |
| waist | Highlander's Chain Girdle (20089) | The League of Arathor [rep] | 68.0 attack_power points (2.89 DPS) | yes | Highlander's Leather Girdle (20116, +0.00 DPS, sim-verified) [rep]; Boar Champion's Belt (10768, -1.61 DPS) [dungeon]; Highlander's Plate Girdle (20125, -1.61 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.78 DPS) | yes | Firemane Leggings (13129, -0.21 DPS, sim-verified) [world_drop]; Orcish War Leggings (7929, -0.34 DPS) [crafted]; Symbolic Legplates (14829, -0.45 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 27.6 attack_power points (1.17 DPS) | yes | Prowler's Leather Shoes (252465, +0.00 DPS, sim-verified) [crafted]; Skirmisher's Mail Boots (252564, -0.24 DPS) [crafted]; Obsidian Greaves (13068, -0.26 DPS) [world_drop] |
| finger1 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.85 DPS) | yes | Thunderbrow Ring (13097, -0.14 DPS) [world_drop]; Suspicious Spare Part (274754, -0.25 DPS) [vendor]; Insurgent's Band (272066, -0.34 DPS) [vendor] |
| finger2 | Protector's Band (19515) | Silverwing Sentinels [rep] | 17.8 attack_power points (0.76 DPS) | yes | Suspicious Spare Part (274754, -0.16 DPS) [vendor]; Thunderbrow Ring (13097, -0.18 DPS, sim-verified) [world_drop]; Protector's Band (19517, -0.19 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-verified (73.2 DPS) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Bonebiter (6830, +0.00 DPS) [quest]; Frost Tiger Blade (3854, -4.28 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Jouster's Chestplate; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Chain Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Assault Band; finger2: Protector's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Fiery War Axe; ranged: Libram of Invocation

No-known-source sample (15 of 546, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 50 (human, 00000000000000000-5500000000000000-55223331211000210)

Set DPS (verified): 102.4. Weights run: 1.1s. Verify run: 1.0s. 714 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.231 ± 0.026, crit=5.537 ± 0.167, hit=3.895 ± 0.064, melee_haste=2.628 ± 0.325

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Plate Helm (220804) | Captain Dirgehammer [vendor] | 142.5 attack_power points (6.12 DPS) | yes | Raging Berserker's Helm (7719, -0.95 DPS, sim-verified) [dungeon]; Ornate Mithril Helm (7937, -1.93 DPS) [crafted]; Eye of Theradras (17715, -2.79 DPS) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.60 DPS) | yes | Skibi's Pendant (13089, +0.00 DPS, sim-verified) [world_drop]; Kaleidoscope Chain (13084, -0.22 DPS) [world_drop]; River Pride Choker (13087, -0.26 DPS) [world_drop] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) | Captain Dirgehammer [vendor] | 99.5 attack_power points (4.28 DPS) | yes | Officer's Pauldrons (250576, -0.27 DPS, sim-verified) [crafted]; Earthslag Shoulders (11632, -3.16 DPS) [dungeon]; Wyrmslayer Spaulders (13066, -3.17 DPS) [world_drop] |
| back | Bloodlust Cape (14801) | World drop [world_drop] | 18.0 attack_power points (0.77 DPS) | yes | Sergeant Major's Cape (16336, -0.20 DPS) [pvp]; Dark Hooded Cape (5257, -0.33 DPS) [world]; Blackveil Cape (11626, -1.20 DPS, sim-verified) [dungeon] |
| chest | Knight's Plate Hauberk (220794) | Captain Dirgehammer [vendor] | 107.5 attack_power points (4.62 DPS) | yes | Ornate Mithril Breastplate (7935, -2.01 DPS, sim-verified) [crafted]; Warforged Chestplate (11195, -2.56 DPS) [quest]; Valorous Chestguard (8274, -2.90 DPS) [world_drop] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 attack_power points (1.20 DPS) | yes | Officer's Wristguards (250581, -0.20 DPS) [crafted]; Giantslayer Bracers (13076, -0.29 DPS) [world_drop]; Runed Golem Shackles (12550, -3.49 DPS, sim-verified) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 97.5 attack_power points (4.19 DPS) | yes | Sergeant Major's Lamellar Gauntlets (220817, +0.00 DPS, sim-verified) [vendor]; Dragonscale Gauntlets (8347, -0.80 DPS) [crafted]; Fletcher's Gloves (7348, -0.86 DPS) [crafted] |
| waist | Highlander's Lamellar Girdle (20106) | The League of Arathor [rep] | 99.5 attack_power points (4.28 DPS) | yes | Highlander's Leather Girdle (20115, -0.09 DPS) [rep]; Highlander's Plate Girdle (20124, -0.09 DPS) [rep]; Highlander's Chain Girdle (20088, -2.11 DPS, sim-verified) [rep] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | sim-verified (102.4 DPS) | yes | Stormshroud Pants (15057, -1.17 DPS, sim-verified) [crafted]; Golem Shard Leggings (13074, -2.57 DPS) [world_drop]; Scarlet Leggings (10330, -2.66 DPS) [dungeon] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 38.9 attack_power points (1.67 DPS) | yes | Prowler's Leather Boots (252468, -0.45 DPS) [crafted]; Officer's Sabatons (250561, -0.47 DPS) [crafted]; Battlechaser's Greaves (12555, -4.57 DPS, sim-verified) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 58.9 attack_power points (2.53 DPS) | yes | Assault Band (13095, -1.67 DPS) [world_drop]; Thunderbrow Ring (13097, -1.82 DPS) [world_drop]; Insurgent's Band (272065, -1.89 DPS) [vendor] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 22.1 attack_power points (0.95 DPS) | yes | Protector's Band (19515, -0.18 DPS) [rep]; Thunderbrow Ring (13097, -0.23 DPS) [world_drop]; Assault Band (13095, -0.38 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (101.3 DPS) | yes | Frozen Heart of the Mountain (249469, -0.58 DPS, sim-verified) [crafted] |
| main_hand | Warmonger (13052) | World drop [world_drop] | sim-verified (101.3 DPS) | yes | Thorium Greatmace (250613, -2.05 DPS) [crafted]; Darkspear Raider's Reaper (272080, -2.47 DPS) [vendor]; Blight (7959, -3.89 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Knight-Lieutenant's Plate Helm; shoulder: Knight-Lieutenant's Plate Pauldrons; back: Bloodlust Cape; chest: Knight's Plate Hauberk; wrist: Bracers of the Stone Princess; waist: Highlander's Lamellar Girdle; legs: Knight's Plate Leggings; feet: Greaves of Withering Despair; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Guardian Talisman; trinket2: Ankh of Life; main_hand: Warmonger

No-known-source sample (15 of 714, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60 (human, 00000000000000000-5532500000000000-55223331211000210)

Set DPS (verified): 182.7. Weights run: 1.1s. Verify run: 1.0s. 1504 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.339 ± 0.051, crit=7.362 ± 0.238, hit=4.660 ± 0.076, melee_haste=3.638 ± 0.629

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Inquisition Helmet (240027) | Leonid Barthalomew the Revered [vendor] | sim-verified (180.9 DPS) | yes | Lionheart Helm (12640, -3.56 DPS, sim-verified) [crafted]; Ragefury Eyepatch (11735, -3.93 DPS) [dungeon]; Bloodvine Lens (19998, -4.46 DPS) [crafted] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (177.4 DPS) | yes | Blazefury Medallion (17111, -2.00 DPS, sim-verified) [world]; Beads of Ogre Might (22150, -2.47 DPS) [quest]; Amulet of the Darkmoon (19491, -4.40 DPS) [quest] |
| shoulder | Inquisition Shoulderplates (240025) | Leonid Barthalomew the Revered [vendor] | 221.7 attack_power points (9.69 DPS) | yes | Inquisition Spaulders (240021, +0.00 DPS, sim-verified) [vendor]; Inquisition Epaulets (246061, -3.15 DPS) [vendor]; Darkspear Spaulders (272108, -4.14 DPS) [vendor] |
| back | Chromatic Cloak (18509) (or Fel Cape (279269)) | Leatherworking [crafted] | 103.1 attack_power points (4.51 DPS) | yes | Fel Cape (279269, +0.00 DPS, sim-verified) [crafted]; Howler's Furs (272414, -1.24 DPS) [vendor]; Earthweave Cloak (21187, -2.25 DPS) [quest] |
| chest | Inquisition Cuirass (246060) | Leonid Barthalomew the Revered [vendor] | sim-verified (177.4 DPS) | yes | Inquisition Breastplate (240030, -1.65 DPS) [vendor]; Bloodsoul Breastplate (19690, -5.98 DPS) [crafted]; Breastplate of Undead Slaying (23087, -13.84 DPS, sim-verified) [world] |
| wrist | Inquisition Armbraces (246056) | Leonid Barthalomew the Revered [vendor] | sim-verified (177.4 DPS) | yes | Vambraces of the Sadist (13400, -0.99 DPS) [dungeon]; Inquisition Vambraces (240023, -1.39 DPS) [vendor]; Bracers of Undead Slaying (23090, -7.21 DPS, sim-verified) [world] |
| hands | Inquisition Gloves (240028) | Leonid Barthalomew the Revered [vendor] | sim-verified (177.4 DPS) | yes | Stormshroud Gloves (21278, -1.29 DPS) [crafted]; Chromatic Gauntlets (19157, -1.40 DPS) [crafted]; Razor Gauntlets (18326, -6.91 DPS, sim-verified) [dungeon] |
| waist | Inquisition Belt (240024) | Leonid Barthalomew the Revered [vendor] | 278.1 attack_power points (12.16 DPS) | yes | Inquisition Cord (246059, +0.00 DPS, sim-verified) [vendor]; Radiant Girdle of the Dawn (227814, -5.29 DPS) [vendor]; Highlander's Plate Girdle (20041, -6.17 DPS) [rep] |
| legs | Inquisition Leggings (240026) | Leonid Barthalomew the Revered [vendor] | sim-verified (177.4 DPS) | yes | Inquisition Legguards (240020, -1.90 DPS) [vendor]; Sentinel's Plate Legguards (237825, -2.01 DPS) [vendor]; Cloudkeeper Legplates (14554, -10.76 DPS, sim-verified) [world_drop] |
| feet | Inquisition Greaves (240029) | Leonid Barthalomew the Revered [vendor] | 223.7 attack_power points (9.78 DPS) | yes | Inquisition Stompers (246057, +0.00 DPS, sim-verified) [vendor]; Knight-Lieutenant's Lamellar Sabatons (227146, -4.22 DPS) [vendor]; Inquisition Boots (240022, -5.27 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (177.4 DPS) | yes | Wrath of Cenarius (21190, -1.73 DPS, sim-verified) [quest]; Band of the Penitent (13217, -2.74 DPS) [quest]; Ring of Entropy (18543, -2.74 DPS) [world] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (177.4 DPS) | yes | Wrath of Cenarius (21190, -1.44 DPS, sim-verified) [quest]; Band of the Penitent (13217, -2.56 DPS) [quest]; Ring of Entropy (18543, -2.56 DPS) [world] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (177.4 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Darkmoon Card: Blue Dragon (19288, -6.31 DPS, sim-verified) [quest] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (179.2 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Darkmoon Card: Heroism (19287, -1.80 DPS, sim-verified) [quest] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-verified (177.4 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [vendor]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Sulfuron Hammer (17193, -0.30 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Inquisition Helmet; neck: Medallion of the Dawn; shoulder: Inquisition Shoulderplates; back: Chromatic Cloak; chest: Inquisition Cuirass; wrist: Inquisition Armbraces; hands: Inquisition Gloves; waist: Inquisition Belt; legs: Inquisition Leggings; feet: Inquisition Greaves; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Darkmoon Card: Maelstrom; trinket2: Thunderbrew's Boot Flask; main_hand: The Unstoppable Force; ranged: Libram of Fervor

No-known-source sample (15 of 1504, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

## Horde

### Band 20 (undead, 00000000000000000-0000000000000000-55100000000000000)

Set DPS (verified): 32.2. Weights run: 1.0s. Verify run: 1.1s. 219 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.133 ± 0.018, crit=2.410 ± 0.076, hit=1.869 ± 0.027, melee_haste=-2.419 ± 0.186

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.69 DPS) | yes | Defender's Leather Hood (252447, -0.14 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.65 DPS) [crafted]; Brawler's Leather Hood (252504, -0.66 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 0.8 attack_power points (0.03 DPS) | yes | Erudite's Amulet (277204, -0.01 DPS, sim-verified) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.21 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS, sim-verified) [crafted]; Serpent's Shoulders (5404, -0.18 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Catacomb Cloak (279899, -0.07 DPS) [quest]; Subterranean Cape (14149, -0.07 DPS) [dungeon]; Grave Shroud (279865, -0.09 DPS, sim-verified) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.69 DPS) | yes | Veteran's Chain Shirt (250488, -0.18 DPS, sim-verified) [crafted]; Defender's Leather Armor (252434, -0.19 DPS) [crafted]; Totemic Leather Armor (252435, -0.21 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Cryptwalker Bracers (280095, -0.07 DPS, sim-verified) [quest]; Raptorcrest Bracers (270010, -0.14 DPS) [quest]; Bristlebark Bindings (14569, -0.19 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | sim-verified (14.5 DPS) | yes | Gold-flecked Gloves (5195, -0.07 DPS) [dungeon]; Blackened Defias Gloves (10401, -0.14 DPS) [dungeon]; Fletcher's Gloves (7348, -0.41 DPS, sim-verified) [crafted] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (0.62 DPS) | yes | Cobrahn's Grasp (6460, -0.12 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.21 DPS) [world]; Support Girdle (1215, -0.28 DPS) [world] |
| legs | Veteran's Chain Leggings (250493) | Blacksmithing [crafted] | 18.7 attack_power points (0.65 DPS) | yes | Defender's Leather Pants (252445, -0.01 DPS, sim-verified) [crafted]; Totemic Leather Pants (252446, -0.02 DPS) [crafted]; Hulking Leggings (14748, -0.08 DPS) [world_drop] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.7 attack_power points (0.37 DPS) | yes | Veteran's Boots (250503, -0.01 DPS, sim-verified) [crafted]; Guard's Boots (250504, -0.02 DPS) [crafted]; Defender's Leather Boots (252441, -0.02 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.5 attack_power points (0.30 DPS) | yes | Loop of Sacrifice (281673, -0.09 DPS) [quest]; The 1 Ring (8350, -0.22 DPS) [world]; Ring of the Moon (12052, -0.23 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Loop of Sacrifice (281673, +0.00 DPS, sim-verified) [quest]; The 1 Ring (8350, -0.20 DPS) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hammerbone (270018) | Leaders of the Fang [quest] | sim-verified (31.9 DPS) | yes | Forsaken Greataxe (251533, -0.11 DPS) [quest]; Smite's Mighty Hammer (7230, -0.27 DPS) [dungeon]; The Axe of Severing (23171, -17.82 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Demon Band; trinket1: Rune of Duty; trinket2: Rune of Perfection; main_hand: Hammerbone; ranged: Libram of Banishment

No-known-source sample (15 of 219, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 5031 ZZZZZZZZ; 5036 ZZZZZ; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5968 Rugged Boots; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 30 (undead, 00000000000000000-0000000000000000-55223310000000000)

Set DPS (verified): 70.8. Weights run: 1.1s. Verify run: 1.1s. 383 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.118 ± 0.014, crit=3.235 ± 0.092, hit=2.259 ± 0.035, melee_haste=1.778 ± 0.112

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 attack_power points (1.08 DPS) | yes | Defender's Leather Helm (252455, -0.08 DPS) [crafted]; Veteran's Chain Helm (250498, -0.14 DPS, sim-verified) [crafted]; Crusader's Chain Helm (250502, -0.17 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.58 DPS) | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; River Pride Choker (13087, -0.25 DPS) [world_drop]; Scout's Medallion (19537, -0.54 DPS) [rep] |
| shoulder | Mail Combat Spaulders (6404) | World drop [world_drop] | sim-verified (67.1 DPS) | yes | Barbaric Iron Shoulders (7913, -0.05 DPS) [crafted]; Elite Shoulders (4835, -0.08 DPS) [vendor]; Golden Scale Shoulders (3841, -0.69 DPS, sim-verified) [crafted] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 attack_power points (0.42 DPS) | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Lambent Scale Cloak (4706, -0.08 DPS) [world_drop]; Slayer's Cape (14752, -0.08 DPS) [world_drop] |
| chest | Shining Silver Breastplate (2870) | Blacksmithing [crafted] | sim-verified (68.4 DPS) | yes | Barbaric Iron Breastplate (7914, -0.17 DPS) [crafted]; Hard Gold Cuirass (250533, -0.25 DPS) [crafted]; Avenger's Armor (1488, -1.94 DPS, sim-verified) [dungeon] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.67 DPS) | yes | Yorgen Bracers (13012, -0.16 DPS, sim-verified) [world_drop]; Bands of Serra'kis (6902, -0.17 DPS) [dungeon]; Grimtoll Wristguards (15459, -0.24 DPS) [quest] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | sim-verified (67.2 DPS) | yes | Warsong Gauntlets (16978, -0.08 DPS) [quest]; Bonefist Gauntlets (4465, -0.17 DPS) [world]; Fletcher's Gloves (7348, -0.76 DPS, sim-verified) [crafted] |
| waist | Girdle of Golem Strength (9405) (or Defiler's Plate Girdle (20207)) | World drop [world_drop] | 24.0 attack_power points (1.00 DPS) | yes | Defiler's Plate Girdle (20207, +0.00 DPS, sim-verified) [rep]; Defiler's Chain Girdle (20152, -0.00 DPS) [rep]; Defiler's Leather Girdle (20191, -0.00 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.08 DPS) | yes | Golden Scale Leggings (3843, -0.17 DPS) [crafted]; Slayer's Pants (14757, -0.17 DPS) [world_drop]; Veteran's Silvered Chain Leggings (250523, -1.95 DPS, sim-verified) [crafted] |
| feet | Hard Gold Boots (250534) | Blacksmithing [crafted] | sim-verified (67.4 DPS) | yes | Glimmering Mail Greaves (4073, -0.08 DPS) [world_drop]; Slayer's Slippers (14756, -0.08 DPS) [world_drop]; Trouncing Boots (4464, -1.01 DPS, sim-verified) [world] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.4 attack_power points (0.68 DPS) | yes | Tiger Band (6749, -0.18 DPS) [quest]; Silverlaine's Family Seal (6321, -0.26 DPS) [dungeon]; Ironspine's Eye (7686, -0.30 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 12.7 attack_power points (0.53 DPS) | yes | Silverlaine's Family Seal (6321, -0.11 DPS) [dungeon]; Ironspine's Eye (7686, -0.15 DPS) [dungeon]; Tiger Band (6749, -1.03 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (66.4 DPS) | yes | Defiler's Talisman (21120, -2.38 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (66.4 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Corpsemaker (6687, -0.14 DPS) [dungeon]; Viscous Hammer (13045, -15.22 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Mail Combat Spaulders; back: Wolfmaster Cape; chest: Shining Silver Breastplate; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Girdle of Golem Strength; legs: Ferine Leggings; feet: Hard Gold Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 383, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 40 (undead, 00000000000000000-0000000000000000-55223331211000210)

Set DPS (verified): 72.8. Weights run: 1.1s. Verify run: 1.0s. 543 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.228 ± 0.031, crit=4.287 ± 0.122, hit=3.119 ± 0.049, melee_haste=2.170 ± 0.198

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 86.0 attack_power points (3.65 DPS) | yes | Icemetal Barbute (10763, -0.35 DPS, sim-verified) [dungeon]; Hard Gold Coif (250537, -2.46 DPS) [crafted]; Chromite Barbute (8142, -2.52 DPS) [world_drop] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.59 DPS) | yes | Ethereal Talisman (4430, -0.05 DPS, sim-verified) [quest]; Kaleidoscope Chain (13084, -0.22 DPS) [world_drop]; River Pride Choker (13087, -0.25 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 attack_power points (0.93 DPS) | yes | Shining Mithril Pauldrons (250541, -0.08 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.17 DPS) [world_drop]; Chromite Pauldrons (8144, -0.63 DPS, sim-verified) [world_drop] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 10.3 attack_power points (0.44 DPS) | yes | Wolfmaster Cape (6314, +0.00 DPS, sim-verified) [dungeon]; Wildhunter Cloak (16658, -0.01 DPS) [quest]; Lambent Scale Cloak (4706, -0.10 DPS) [world_drop] |
| chest | Jouster's Chestplate (8157) | World drop [world_drop] | sim-verified (72.8 DPS) | yes | Kolkar Marauder Chain (6773, -0.02 DPS) [quest]; Shining Silver Breastplate (2870, -0.08 DPS) [crafted]; Avenger's Armor (1488, -1.59 DPS, sim-verified) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.85 DPS) | yes | Pugilist Bracers (4438, -0.21 DPS, sim-verified) [dungeon]; Ravager's Armguards (14770, -0.22 DPS) [world_drop]; Darkspear Armsplints (4132, -0.25 DPS) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 80.0 attack_power points (3.40 DPS) | yes | Dragonscale Gauntlets (8347, -0.82 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -0.85 DPS) [crafted]; Ornate Mithril Gloves (7927, -0.85 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20153) | The Defilers [rep] | 68.0 attack_power points (2.89 DPS) | yes | Defiler's Leather Girdle (20192, +0.00 DPS, sim-verified) [rep]; Boar Champion's Belt (10768, -1.61 DPS) [dungeon]; Defiler's Plate Girdle (20206, -1.61 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.78 DPS) | yes | Firemane Leggings (13129, -0.21 DPS, sim-verified) [world_drop]; Orcish War Leggings (7929, -0.34 DPS) [crafted]; Symbolic Legplates (14829, -0.45 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 27.6 attack_power points (1.17 DPS) | yes | Prowler's Leather Shoes (252465, -0.12 DPS, sim-verified) [crafted]; Skirmisher's Mail Boots (252564, -0.24 DPS) [crafted]; Obsidian Greaves (13068, -0.26 DPS) [world_drop] |
| finger1 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.85 DPS) | yes | Thunderbrow Ring (13097, -0.14 DPS) [world_drop]; Suspicious Spare Part (274754, -0.25 DPS) [vendor]; Insurgent's Band (272066, -0.34 DPS) [vendor] |
| finger2 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 17.8 attack_power points (0.76 DPS) | yes | Suspicious Spare Part (274754, -0.16 DPS) [vendor]; Legionnaire's Band (19513, -0.19 DPS) [rep]; Thunderbrow Ring (13097, -0.19 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-verified (71.2 DPS) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor]; Frost Tiger Blade (3854, -9.23 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; shoulder: Hard Gold Pauldrons; back: Dark Hooded Cape; chest: Jouster's Chestplate; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Chain Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Assault Band; finger2: Legionnaire's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Fiery War Axe; ranged: Libram of Invocation

No-known-source sample (15 of 543, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band

### Band 50 (undead, 00000000000000000-5500000000000000-55223331211000210)

Set DPS (verified): 99.3. Weights run: 1.1s. Verify run: 1.0s. 725 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.231 ± 0.026, crit=5.537 ± 0.167, hit=3.895 ± 0.064, melee_haste=2.628 ± 0.325

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Plate Helm (220803) | Lady Palanseer [vendor] | 142.5 attack_power points (6.12 DPS) | yes | Raging Berserker's Helm (7719, -0.91 DPS, sim-verified) [dungeon]; Ornate Mithril Helm (7937, -1.93 DPS) [crafted]; Eye of Theradras (17715, -2.79 DPS) [dungeon] |
| neck | Woven Ivy Necklace (19159) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 14.1 attack_power points (0.61 DPS) | yes | Skibi's Pendant (13089, -0.05 DPS) [world_drop]; Ethereal Talisman (4430, -0.14 DPS) [quest]; Ghostshard Talisman (7731, -0.29 DPS, sim-verified) [dungeon] |
| shoulder | Blood Guard's Plate Pauldrons (220796) | Lady Palanseer [vendor] | 99.5 attack_power points (4.28 DPS) | yes | Officer's Pauldrons (250576, -0.29 DPS, sim-verified) [crafted]; Earthslag Shoulders (11632, -3.16 DPS) [dungeon]; Wyrmslayer Spaulders (13066, -3.17 DPS) [world_drop] |
| back | Bloodlust Cape (14801) | World drop [world_drop] | 18.0 attack_power points (0.77 DPS) | yes | Dark Hooded Cape (5257, -0.33 DPS) [world]; Wolfmaster Cape (6314, -0.34 DPS) [dungeon]; Blackveil Cape (11626, -1.08 DPS, sim-verified) [dungeon] |
| chest | Stone Guard's Plate Armor (220801) | Lady Palanseer [vendor] | 107.5 attack_power points (4.62 DPS) | yes | Ornate Mithril Breastplate (7935, -1.79 DPS, sim-verified) [crafted]; Warforged Chestplate (11195, -2.56 DPS) [quest]; Valorous Chestguard (8274, -2.90 DPS) [world_drop] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 attack_power points (1.20 DPS) | yes | Officer's Wristguards (250581, -0.20 DPS) [crafted]; Giantslayer Bracers (13076, -0.29 DPS) [world_drop]; Runed Golem Shackles (12550, -2.69 DPS, sim-verified) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 97.5 attack_power points (4.19 DPS) | yes | Fletcher's Gloves (7348, -0.86 DPS) [crafted]; Ornate Mithril Gloves (7927, -0.86 DPS) [crafted]; Dragonscale Gauntlets (8347, -0.92 DPS, sim-verified) [crafted] |
| waist | Defiler's Chain Girdle (20151) (or Defiler's Leather Girdle (20193)) | The Defilers [rep] | 97.5 attack_power points (4.19 DPS) | yes | Defiler's Leather Girdle (20193, +0.00 DPS, sim-verified) [rep]; Defiler's Plate Girdle (20205, -0.00 DPS) [rep]; Defiler's Chain Girdle (20153, -0.52 DPS) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | 155.0 attack_power points (6.66 DPS) | yes | Stone Guard's Plate Leggings (220798, +0.00 DPS, sim-verified) [vendor]; Golem Shard Leggings (13074, -4.77 DPS) [world_drop]; Scarlet Leggings (10330, -4.86 DPS) [dungeon] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 38.9 attack_power points (1.67 DPS) | yes | Prowler's Leather Boots (252468, -0.45 DPS) [crafted]; Officer's Sabatons (250561, -0.47 DPS) [crafted]; Battlechaser's Greaves (12555, -4.39 DPS, sim-verified) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 58.9 attack_power points (2.53 DPS) | yes | Legionnaire's Band (19511, -1.58 DPS) [rep]; Assault Band (13095, -1.67 DPS) [world_drop]; Band of Allegiance (18585, -1.76 DPS) [quest] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.03 DPS) | yes | Legionnaire's Band (19511, +0.00 DPS, sim-verified) [rep]; Assault Band (13095, -0.17 DPS) [world_drop]; Band of Allegiance (18585, -0.26 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (99.3 DPS) | yes | Ankh of Life (1713, -3.13 DPS, sim-verified) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (99.3 DPS) | yes | Ankh of Life (1713, -0.66 DPS, sim-verified) [world_drop] |
| main_hand | Blight (7959) | Blacksmithing [crafted] | sim-verified (99.3 DPS) | yes | Thorium Greatmace (250613, +0.00 DPS) [crafted]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor]; Warmonger (13052, -3.09 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Blood Guard's Plate Helm; neck: Woven Ivy Necklace; shoulder: Blood Guard's Plate Pauldrons; back: Bloodlust Cape; chest: Stone Guard's Plate Armor; wrist: Bracers of the Stone Princess; waist: Defiler's Chain Girdle; legs: Stormshroud Pants; feet: Greaves of Withering Despair; finger1: Blackstone Ring; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Blight

No-known-source sample (15 of 725, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band

### Band 60 (undead, 00000000000000000-5532500000000000-55223331211000210)

Set DPS (verified): 186.2. Weights run: 1.1s. Verify run: 1.0s. 1552 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.339 ± 0.051, crit=7.362 ± 0.238, hit=4.660 ± 0.076, melee_haste=3.638 ± 0.629

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Inquisition Helmet (240027) | Leonid Barthalomew the Revered [vendor] | sim-verified (184.0 DPS) | yes | Lionheart Helm (12640, -3.66 DPS, sim-verified) [crafted]; Ragefury Eyepatch (11735, -3.93 DPS) [dungeon]; Bloodvine Lens (19998, -4.46 DPS) [crafted] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (180.3 DPS) | yes | Blazefury Medallion (17111, -2.21 DPS, sim-verified) [world]; Beads of Ogre Might (22150, -2.47 DPS) [quest]; Amulet of the Darkmoon (19491, -4.40 DPS) [quest] |
| shoulder | Inquisition Spaulders (240021) | Leonid Barthalomew the Revered [vendor] | sim-verified (182.6 DPS) | yes | Inquisition Shoulderplates (240025, -2.27 DPS, sim-verified) [vendor]; Inquisition Epaulets (246061, -2.47 DPS) [vendor]; Darkspear Spaulders (272108, -3.46 DPS) [vendor] |
| back | Chromatic Cloak (18509) (or Fel Cape (279269)) | Leatherworking [crafted] | 103.1 attack_power points (4.51 DPS) | yes | Fel Cape (279269, +0.00 DPS, sim-verified) [crafted]; Howler's Furs (272414, -1.24 DPS) [vendor]; Earthweave Cloak (21187, -2.25 DPS) [quest] |
| chest | Inquisition Cuirass (246060) | Leonid Barthalomew the Revered [vendor] | sim-verified (180.3 DPS) | yes | Inquisition Breastplate (240030, -1.65 DPS) [vendor]; Bloodsoul Breastplate (19690, -5.98 DPS) [crafted]; Breastplate of Undead Slaying (23087, -13.75 DPS, sim-verified) [world] |
| wrist | Inquisition Armbraces (246056) | Leonid Barthalomew the Revered [vendor] | sim-verified (180.3 DPS) | yes | Vambraces of the Sadist (13400, -0.99 DPS) [dungeon]; Inquisition Vambraces (240023, -1.39 DPS) [vendor]; Bracers of Undead Slaying (23090, -5.88 DPS, sim-verified) [world] |
| hands | Inquisition Gloves (240028) | Leonid Barthalomew the Revered [vendor] | sim-verified (180.3 DPS) | yes | Stormshroud Gloves (21278, -1.29 DPS) [crafted]; Chromatic Gauntlets (19157, -1.40 DPS) [crafted]; Razor Gauntlets (18326, -6.77 DPS, sim-verified) [dungeon] |
| waist | Inquisition Belt (240024) | Leonid Barthalomew the Revered [vendor] | 278.1 attack_power points (12.16 DPS) | yes | Inquisition Cord (246059, +0.00 DPS, sim-verified) [vendor]; Radiant Girdle of the Dawn (227814, -5.29 DPS) [vendor]; Defiler's Plate Girdle (20204, -6.17 DPS) [rep] |
| legs | Inquisition Leggings (240026) | Leonid Barthalomew the Revered [vendor] | sim-verified (180.3 DPS) | yes | Inquisition Legguards (240020, -1.90 DPS) [vendor]; Sentinel's Plate Legguards (237825, -2.01 DPS) [vendor]; Cloudkeeper Legplates (14554, -10.14 DPS, sim-verified) [world_drop] |
| feet | Inquisition Greaves (240029) | Leonid Barthalomew the Revered [vendor] | 223.7 attack_power points (9.78 DPS) | yes | Inquisition Stompers (246057, +0.00 DPS, sim-verified) [vendor]; Inquisition Boots (240022, -5.27 DPS) [vendor]; Fine Dawn Treaders (227815, -7.58 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (180.3 DPS) | yes | Wrath of Cenarius (21190, -2.01 DPS, sim-verified) [quest]; Band of the Penitent (13217, -2.74 DPS) [quest]; Ring of Entropy (18543, -2.74 DPS) [world] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (180.3 DPS) | yes | Wrath of Cenarius (21190, -1.73 DPS, sim-verified) [quest]; Band of the Penitent (13217, -2.56 DPS) [quest]; Ring of Entropy (18543, -2.56 DPS) [world] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (180.3 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (180.3 DPS) | yes | Frozen Heart of the Mountain (249469, -1.43 DPS) [crafted]; Darkmoon Card: Heroism (19287, -3.95 DPS, sim-verified) [quest] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (180.3 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [vendor]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Sulfuron Hammer (17193, -1.47 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Inquisition Helmet; neck: Medallion of the Dawn; shoulder: Inquisition Spaulders; back: Chromatic Cloak; chest: Inquisition Cuirass; wrist: Inquisition Armbraces; hands: Inquisition Gloves; waist: Inquisition Belt; legs: Inquisition Leggings; feet: Inquisition Greaves; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: The Unstoppable Force; ranged: Libram of Fervor

No-known-source sample (15 of 1552, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band

