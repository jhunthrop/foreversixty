# Leveling BiS: Retribution

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-0000000000000000-55100000000000000)

Set DPS (verified): 34.4. Weights run: 0.8s. Verify run: 0.7s. 208 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.133 ± 0.018, crit=2.410 ± 0.076, hit=1.703 ± 0.139, melee_haste=-2.419 ± 0.186

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 | yes | Defender's Leather Hood (252447, -0.16 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.65 DPS) [crafted]; Brawler's Leather Hood (252504, -0.66 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 0.8 | yes | Scholarly Pendant (277203, -0.03 DPS) [quest]; Tarnished Locket (279870, -0.03 DPS) [quest]; Erudite's Amulet (277204, -0.04 DPS, sim-verified) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 | yes | Silvered Bronze Shoulders (3481, -0.03 DPS, sim-verified) [crafted]; Serpent's Shoulders (5404, -0.18 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.21 DPS) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 | yes | Catacomb Cloak (279899, -0.07 DPS) [quest]; Grave Shroud (279865, -0.10 DPS, sim-verified) [quest]; Dark Leather Cloak (2316, -0.12 DPS) [crafted] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 | yes | Veteran's Chain Shirt (250488, -0.18 DPS, sim-verified) [crafted]; Defender's Leather Armor (252434, -0.19 DPS) [crafted]; Totemic Leather Armor (252435, -0.21 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 | yes | Cryptwalker Bracers (280095, -0.08 DPS, sim-verified) [quest]; Bravo's Armbands (270015, -0.19 DPS) [quest]; Bristlebark Bindings (14569, -0.19 DPS) [world_drop] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 33.7 | yes | Thorbia's Gauntlets (12994, +0.00 DPS, sim-verified) [world_drop]; Gold-flecked Gloves (5195, -0.68 DPS) [dungeon]; Polar Gauntlets (7606, -0.75 DPS) [quest] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 | yes | Cobrahn's Grasp (6460, -0.12 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.21 DPS) [world]; Support Girdle (1215, -0.28 DPS) [world] |
| legs | Chausses of Westfall (6087) | The Defias Brotherhood [quest] | 22.0 | yes | Veteran's Chain Leggings (250493, -0.08 DPS, sim-verified) [crafted]; Defender's Leather Pants (252445, -0.12 DPS) [crafted]; Totemic Leather Pants (252446, -0.14 DPS) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.7 | yes | Veteran's Boots (250503, -0.02 DPS, sim-verified) [crafted]; Guard's Boots (250504, -0.02 DPS) [crafted]; Defender's Leather Boots (252441, -0.02 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.5 | yes | Loop of Sacrifice (281673, -0.09 DPS) [quest]; The 1 Ring (8350, -0.22 DPS) [world]; Ring of the Moon (12052, -0.23 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 | yes | Loop of Sacrifice (281673, +0.00 DPS, sim-verified) [quest]; The 1 Ring (8350, -0.20 DPS) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Verigan's Fist (6953) | The Test of Righteousness [quest] | 358.7 | yes | Smite's Mighty Hammer (7230, -2.09 DPS) [dungeon]; Duskbringer (2205, -2.23 DPS) [dungeon]; Forsaken Greataxe (251533, -3.62 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Verigan's Fist; ranged: Libram of Banishment

No-known-source sample (15 of 208, see the JSON for more): 1189 Overseer's Ring; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6478 Rat Stompers; 7955 Copper Claymore; 7956 Bronze Warhammer; 9602 Brushwood Blade; 10047 Simple Kilt; 10421 Rough Copper Vest

### Band 30 (human, 00000000000000000-0000000000000000-55223310000000000)

Set DPS (verified): 71.6. Weights run: 0.8s. Verify run: 0.9s. 364 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.118 ± 0.014, crit=3.235 ± 0.092, hit=1.979 ± 0.168, melee_haste=1.778 ± 0.112

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 | yes | Defender's Leather Helm (252455, -0.08 DPS) [crafted]; Veteran's Chain Helm (250498, -0.14 DPS, sim-verified) [crafted]; Crusader's Chain Helm (250502, -0.17 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; River Pride Choker (13087, -0.25 DPS) [world_drop]; Sentinel's Medallion (19541, -0.54 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 14.0 | yes | Mail Combat Spaulders (6404, +0.00 DPS, sim-verified) [world_drop]; Barbaric Iron Shoulders (7913, -0.05 DPS) [crafted]; Elite Shoulders (4835, -0.08 DPS) [vendor] |
| back | Sergeant Major's Cape (16315) | PvP rank 9 · Sergeant Major · Alliance [pvp] | sim-verified (69.8 DPS) | yes | Lambent Scale Cloak (4706, -0.02 DPS) [world_drop]; Slayer's Cape (14752, -0.02 DPS) [world_drop]; Wolfmaster Cape (6314, -0.71 DPS, sim-verified) [dungeon] |
| chest | Shining Silver Breastplate (2870) | Blacksmithing [crafted] | 28.0 | yes | Hard Gold Cuirass (250533, -0.25 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.28 DPS, sim-verified) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.30 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 | yes | Yorgen Bracers (13012, -0.16 DPS, sim-verified) [world_drop]; Bands of Serra'kis (6902, -0.17 DPS) [dungeon]; Patterned Bronze Bracers (2868, -0.25 DPS) [crafted] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | sim-verified (69.8 DPS) | yes | Bonefist Gauntlets (4465, -0.17 DPS) [world]; Mail Combat Gauntlets (4075, -0.23 DPS) [world_drop]; Fletcher's Gloves (7348, -0.72 DPS, sim-verified) [crafted] |
| waist | Girdle of Golem Strength (9405) (or Highlander's Plate Girdle (20126)) | World drop [world_drop] | 24.0 | yes | Highlander's Plate Girdle (20126, +0.00 DPS, sim-verified) [rep]; Highlander's Chain Girdle (20090, -0.00 DPS) [rep]; Highlander's Leather Girdle (20117, -0.00 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Golden Scale Leggings (3843, -0.17 DPS) [crafted]; Chausses of Westfall (6087, -0.17 DPS) [quest]; Veteran's Silvered Chain Leggings (250523, -2.31 DPS, sim-verified) [crafted] |
| feet | Hard Gold Boots (250534) | Blacksmithing [crafted] | sim-verified (70.1 DPS) | yes | Glimmering Mail Greaves (4073, -0.08 DPS) [world_drop]; Slayer's Slippers (14756, -0.08 DPS) [world_drop]; Trouncing Boots (4464, -1.00 DPS, sim-verified) [world] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.4 | yes | Silverlaine's Family Seal (6321, -0.26 DPS) [dungeon]; Ironspine's Eye (7686, -0.30 DPS) [dungeon]; Insurgent's Band (272067, -0.31 DPS) [vendor] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 12.7 | yes | Ironspine's Eye (7686, -0.15 DPS) [dungeon]; Insurgent's Band (272067, -0.15 DPS) [vendor]; Silverlaine's Family Seal (6321, -0.39 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (52.9 DPS) | yes | Minor Recombobulator (4381, +0.00 DPS) [crafted]; Gnomish Universal Remote (7506, +0.00 DPS) [crafted]; Talisman of Arathor (21119, -2.96 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (68.8 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Corpsemaker (6687, -0.14 DPS) [dungeon]; Viscous Hammer (13045, -16.10 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Sergeant Major's Cape; chest: Shining Silver Breastplate; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Girdle of Golem Strength; legs: Ferine Leggings; feet: Hard Gold Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 364, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6478 Rat Stompers; 7955 Copper Claymore

### Band 40 (human, 00000000000000000-0000000000000000-55223331211000210)

Set DPS (verified): 85.1. Weights run: 0.8s. Verify run: 0.8s. 511 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.228 ± 0.031, crit=4.287 ± 0.122, hit=2.725 ± 0.232, melee_haste=2.170 ± 0.198

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 86.0 | yes | Icemetal Barbute (10763, -0.34 DPS, sim-verified) [dungeon]; Hard Gold Coif (250537, -2.46 DPS) [crafted]; Chromite Barbute (8142, -2.52 DPS) [world_drop] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; River Pride Choker (13087, -0.25 DPS) [world_drop]; Gazlowe's Charm (13088, -0.25 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 | yes | Shining Mithril Pauldrons (250541, -0.08 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.17 DPS) [world_drop]; Chromite Pauldrons (8144, -0.93 DPS, sim-verified) [world_drop] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 13.4 | yes | Sergeant Major's Cape (16315, -0.19 DPS) [pvp]; Lambent Scale Cloak (4706, -0.23 DPS) [world_drop]; Wolfmaster Cape (6314, -0.90 DPS, sim-verified) [dungeon] |
| chest | Jouster's Chestplate (8157) | World drop [world_drop] | 30.0 | yes | Kolkar Marauder Chain (6773, +0.00 DPS, sim-verified) [quest]; Shining Silver Breastplate (2870, -0.08 DPS) [crafted]; Golden Scale Cuirass (3845, -0.08 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Ravager's Armguards (14770, -0.22 DPS) [world_drop]; Pugilist Bracers (4438, -0.30 DPS, sim-verified) [dungeon]; Yorgen Bracers (13012, -0.31 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 80.0 | yes | Fletcher's Gloves (7348, -0.85 DPS) [crafted]; Ornate Mithril Gloves (7927, -0.85 DPS) [crafted]; Dragonscale Gauntlets (8347, -1.24 DPS, sim-verified) [crafted] |
| waist | Highlander's Chain Girdle (20089) | The League of Arathor [rep] | 68.0 | yes | Highlander's Leather Girdle (20116, +0.00 DPS, sim-verified) [rep]; Highlander's Plate Girdle (20125, -1.61 DPS) [rep]; Highlander's Chain Girdle (20090, -1.87 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 | yes | Firemane Leggings (13129, -0.30 DPS, sim-verified) [world_drop]; Orcish War Leggings (7929, -0.34 DPS) [crafted]; Symbolic Legplates (14829, -0.45 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 27.6 | yes | Prowler's Leather Shoes (252465, +0.00 DPS, sim-verified) [crafted]; Skirmisher's Mail Boots (252564, -0.24 DPS) [crafted]; Obsidian Greaves (13068, -0.26 DPS) [world_drop] |
| finger1 | Assault Band (13095) | World drop [world_drop] | 20.0 | yes | Thunderbrow Ring (13097, -0.14 DPS) [world_drop]; Suspicious Spare Part (274754, -0.25 DPS) [vendor]; Insurgent's Band (272066, -0.34 DPS) [vendor] |
| finger2 | Protector's Band (19515) | Silverwing Sentinels [rep] | 17.8 | yes | Suspicious Spare Part (274754, -0.16 DPS) [vendor]; Protector's Band (19517, -0.19 DPS) [rep]; Thunderbrow Ring (13097, -0.23 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (84.9 DPS) | yes | Bonebiter (6830, +0.00 DPS) [quest]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor]; Fiery War Axe (870, -10.31 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Jouster's Chestplate; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Chain Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Assault Band; finger2: Protector's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; ranged: Libram of Invocation

No-known-source sample (15 of 511, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 50 (human, 00000000000000000-5500000000000000-55223331211000210)

Set DPS (verified): 102.4. Weights run: 0.9s. Verify run: 0.8s. 669 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.231 ± 0.026, crit=5.537 ± 0.167, hit=3.474 ± 0.300, melee_haste=2.628 ± 0.325

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Plate Helm (220804) | Captain Dirgehammer [vendor] | 138.3 | yes | Raging Berserker's Helm (7719, -0.95 DPS, sim-verified) [dungeon]; Ornate Mithril Helm (7937, -1.75 DPS) [crafted]; Eye of Theradras (17715, -2.61 DPS) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Skibi's Pendant (13089, +0.00 DPS, sim-verified) [world_drop]; Kaleidoscope Chain (13084, -0.22 DPS) [world_drop]; River Pride Choker (13087, -0.26 DPS) [world_drop] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) | Captain Dirgehammer [vendor] | 99.5 | yes | Officer's Pauldrons (250576, -0.27 DPS, sim-verified) [crafted]; Earthslag Shoulders (11632, -3.16 DPS) [dungeon]; Wyrmslayer Spaulders (13066, -3.17 DPS) [world_drop] |
| back | Bloodlust Cape (14801) | World drop [world_drop] | 18.0 | yes | Sergeant Major's Cape (16336, +0.00 DPS, sim-verified) [pvp]; Wolfmaster Cape (6314, -0.34 DPS) [dungeon]; Sergeant Major's Cape (16315, -0.39 DPS) [pvp] |
| chest | Knight's Plate Hauberk (220794) | Captain Dirgehammer [vendor] | 107.5 | yes | Ornate Mithril Breastplate (7935, -2.01 DPS, sim-verified) [crafted]; Warforged Chestplate (11195, -2.56 DPS) [quest]; Valorous Chestguard (8274, -2.90 DPS) [world_drop] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 | yes | Giantslayer Bracers (13076, -0.29 DPS) [world_drop]; Branded Leather Bracers (19508, -0.34 DPS) [dungeon]; Officer's Wristguards (250581, -1.72 DPS, sim-verified) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 97.5 | yes | Sergeant Major's Lamellar Gauntlets (220817, +0.00 DPS, sim-verified) [vendor]; Dragonscale Gauntlets (8347, -0.80 DPS) [crafted]; Fletcher's Gloves (7348, -0.86 DPS) [crafted] |
| waist | Highlander's Lamellar Girdle (20106) | The League of Arathor [rep] | 99.5 | yes | Highlander's Leather Girdle (20115, -0.09 DPS) [rep]; Highlander's Plate Girdle (20124, -0.09 DPS) [rep]; Highlander's Chain Girdle (20088, -2.11 DPS, sim-verified) [rep] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | sim-verified (102.4 DPS) | yes | Stormshroud Pants (15057, -1.17 DPS, sim-verified) [crafted]; Golem Shard Leggings (13074, -2.57 DPS) [world_drop]; Scarlet Leggings (10330, -2.66 DPS) [dungeon] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 34.7 | yes | Officer's Sabatons (250561, -0.29 DPS) [crafted]; Officer's Boots (250546, -0.31 DPS) [crafted]; Prowler's Leather Boots (252468, -2.31 DPS, sim-verified) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 54.7 | yes | Assault Band (13095, -1.49 DPS) [world_drop]; Thunderbrow Ring (13097, -1.63 DPS) [world_drop]; Insurgent's Band (272065, -1.71 DPS) [vendor] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 22.1 | yes | Protector's Band (19515, -0.18 DPS) [rep]; Thunderbrow Ring (13097, -0.23 DPS) [world_drop]; Assault Band (13095, -0.38 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (100.9 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -0.58 DPS, sim-verified) [crafted] |
| main_hand | Warmonger (13052) | World drop [world_drop] | sim-verified (100.9 DPS) | yes | Thorium Greatmace (250613, -1.51 DPS) [crafted]; Darkspear Raider's Reaper (272080, -1.93 DPS) [vendor]; Blight (7959, -3.89 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Knight-Lieutenant's Plate Helm; shoulder: Knight-Lieutenant's Plate Pauldrons; back: Bloodlust Cape; chest: Knight's Plate Hauberk; wrist: Bracers of the Stone Princess; waist: Highlander's Lamellar Girdle; legs: Knight's Plate Leggings; feet: Greaves of Withering Despair; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Guardian Talisman; trinket2: Ankh of Life; main_hand: Warmonger

No-known-source sample (15 of 669, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (human, 00000000000000000-5532500000000000-55223331211000210)

Set DPS (verified): 182.9. Weights run: 0.9s. Verify run: 0.9s. 1232 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.339 ± 0.051, crit=7.362 ± 0.238, hit=9.891 ± 0.551, melee_haste=3.638 ± 0.629

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Inquisition Helmet (240027) | Leonid Barthalomew the Revered [vendor] | sim-verified (180.8 DPS) | yes | Mask of the Unforgiven (13404, -0.32 DPS) [dungeon]; Bloodvine Goggles (19999, -0.32 DPS) [crafted]; Lionheart Helm (12640, -4.86 DPS, sim-verified) [crafted] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (176.4 DPS) | yes | Beads of Ogre Might (22150, -0.18 DPS) [quest]; Blazefury Medallion (17111, -1.08 DPS, sim-verified) [world]; Amulet of the Darkmoon (19491, -4.40 DPS) [quest] |
| shoulder | Inquisition Spaulders (240021) | Leonid Barthalomew the Revered [vendor] | sim-verified (178.4 DPS) | yes | Inquisition Epaulets (246061, -0.18 DPS) [vendor]; Inquisition Shoulderplates (240025, -2.43 DPS, sim-verified) [vendor]; Darkspear Spaulders (272108, -3.46 DPS) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 126.9 | yes | Earthweave Cloak (21187, -0.97 DPS, sim-verified) [quest]; Chromatic Cloak (18509, -1.04 DPS) [crafted]; Arcanoweave Cloak (272411, -1.22 DPS) [vendor] |
| chest | Inquisition Cuirass (246060) | Leonid Barthalomew the Revered [vendor] | 502.9 | yes | Inquisition Breastplate (240030, -4.42 DPS, sim-verified) [vendor]; Dawn Armor (252483, -12.61 DPS) [crafted]; Bloodsoul Breastplate (19690, -12.84 DPS) [crafted] |
| wrist | Inquisition Armbraces (246056) | Leonid Barthalomew the Revered [vendor] | 296.7 | yes | Inquisition Vambraces (240023, -1.25 DPS, sim-verified) [vendor]; Inquisition Bracers (240031, -7.42 DPS) [vendor]; Vambraces of the Sadist (13400, -7.85 DPS) [dungeon] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | 202.0 | yes | Primal Batskin Gloves (19686, -0.72 DPS, sim-verified) [crafted]; Inquisition Gloves (240028, -1.00 DPS) [vendor]; Chromatic Gauntlets (19157, -2.40 DPS) [crafted] |
| waist | Inquisition Belt (240024) | Leonid Barthalomew the Revered [vendor] | 278.1 | yes | Inquisition Cord (246059, +0.00 DPS, sim-verified) [vendor]; Radiant Girdle of the Dawn (227814, -5.29 DPS) [vendor]; Highlander's Plate Girdle (20041, -6.17 DPS) [rep] |
| legs | Inquisition Leggings (240026) | Leonid Barthalomew the Revered [vendor] | 405.0 | yes | Sentinel's Plate Legguards (237825, -2.01 DPS) [vendor]; Sentinel's Chain Leggings (237819, -3.85 DPS) [vendor]; Titanic Leggings (22385, -4.41 DPS, sim-verified) [crafted] |
| feet | Inquisition Greaves (240029) | Leonid Barthalomew the Revered [vendor] | 276.0 | yes | Inquisition Stompers (246057, +0.00 DPS, sim-verified) [vendor]; Knight-Lieutenant's Lamellar Sabatons (227146, -6.51 DPS) [vendor]; Inquisition Boots (240022, -7.56 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (173.5 DPS) | yes | Wrath of Cenarius (21190, -1.47 DPS, sim-verified) [quest]; Signet Ring of the Bronze Dragonflight (234034, -3.11 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234030, -3.28 DPS) [vendor] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (173.5 DPS) | yes | Wrath of Cenarius (21190, -1.18 DPS, sim-verified) [quest]; Signet Ring of the Bronze Dragonflight (234034, -2.93 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234030, -3.11 DPS) [vendor] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (174.2 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Darkmoon Card: Blue Dragon (19288, -6.14 DPS, sim-verified) [quest] |
| trinket2 | Darkmoon Card: Heroism (19287) | Darkmoon Warlords Deck [quest] | sim-verified (173.5 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS, sim-verified) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Arcanite Champion (12790) | Blacksmithing [crafted] | sim-verified (176.4 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [vendor]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; The Unstoppable Force (19323, -3.51 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Inquisition Helmet; neck: Medallion of the Dawn; shoulder: Inquisition Spaulders; back: Howler's Furs; chest: Inquisition Cuirass; wrist: Inquisition Armbraces; hands: Stormshroud Gloves; waist: Inquisition Belt; legs: Inquisition Leggings; feet: Inquisition Greaves; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Darkmoon Card: Maelstrom; trinket2: Darkmoon Card: Heroism; main_hand: Arcanite Champion; ranged: Libram of Fervor

No-known-source sample (15 of 1232, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

## Horde

### Band 20 (undead, 00000000000000000-0000000000000000-55100000000000000)

Set DPS (verified): 31.9. Weights run: 0.8s. Verify run: 0.7s. 206 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.133 ± 0.018, crit=2.410 ± 0.076, hit=1.703 ± 0.139, melee_haste=-2.419 ± 0.186

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 | yes | Defender's Leather Hood (252447, -0.16 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.65 DPS) [crafted]; Brawler's Leather Hood (252504, -0.66 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 0.8 | yes | Scholarly Pendant (277203, -0.03 DPS) [quest]; Tarnished Locket (279870, -0.03 DPS) [quest]; Erudite's Amulet (277204, -0.05 DPS, sim-verified) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 | yes | Silvered Bronze Shoulders (3481, +0.00 DPS, sim-verified) [crafted]; Serpent's Shoulders (5404, -0.18 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.21 DPS) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 | yes | Catacomb Cloak (279899, -0.07 DPS) [quest]; Subterranean Cape (14149, -0.07 DPS) [dungeon]; Grave Shroud (279865, -0.15 DPS, sim-verified) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 | yes | Veteran's Chain Shirt (250488, -0.18 DPS, sim-verified) [crafted]; Defender's Leather Armor (252434, -0.19 DPS) [crafted]; Totemic Leather Armor (252435, -0.21 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 | yes | Cryptwalker Bracers (280095, -0.08 DPS, sim-verified) [quest]; Raptorcrest Bracers (270010, -0.14 DPS) [quest]; Bristlebark Bindings (14569, -0.19 DPS) [world_drop] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 33.7 | yes | Thorbia's Gauntlets (12994, +0.00 DPS, sim-verified) [world_drop]; Gold-flecked Gloves (5195, -0.68 DPS) [dungeon]; Blackened Defias Gloves (10401, -0.75 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 | yes | Cobrahn's Grasp (6460, -0.12 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.21 DPS) [world]; Support Girdle (1215, -0.28 DPS) [world] |
| legs | Veteran's Chain Leggings (250493) | Blacksmithing [crafted] | 18.7 | yes | Defender's Leather Pants (252445, -0.02 DPS, sim-verified) [crafted]; Totemic Leather Pants (252446, -0.02 DPS) [crafted]; Hulking Leggings (14748, -0.08 DPS) [world_drop] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.7 | yes | Veteran's Boots (250503, -0.02 DPS, sim-verified) [crafted]; Guard's Boots (250504, -0.02 DPS) [crafted]; Defender's Leather Boots (252441, -0.02 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.5 | yes | Loop of Sacrifice (281673, -0.09 DPS) [quest]; The 1 Ring (8350, -0.22 DPS) [world]; Ring of the Moon (12052, -0.23 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 | yes | Loop of Sacrifice (281673, +0.00 DPS, sim-verified) [quest]; The 1 Ring (8350, -0.20 DPS) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hammerbone (270018) | Leaders of the Fang [quest] | 306.2 | yes | Smite's Mighty Hammer (7230, -0.27 DPS) [dungeon]; Duskbringer (2205, -0.41 DPS) [dungeon]; Forsaken Greataxe (251533, -1.56 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Demon Band; trinket1: Rune of Duty; trinket2: Rune of Perfection; main_hand: Hammerbone; ranged: Libram of Banishment

No-known-source sample (15 of 206, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7955 Copper Claymore; 7956 Bronze Warhammer; 9602 Brushwood Blade; 10047 Simple Kilt; 10421 Rough Copper Vest; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 15402 Noosegrip Gauntlets; 18864 Insignia of the Alliance

### Band 30 (undead, 00000000000000000-0000000000000000-55223310000000000)

Set DPS (verified): 70.8. Weights run: 0.8s. Verify run: 0.9s. 364 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.118 ± 0.014, crit=3.235 ± 0.092, hit=1.979 ± 0.168, melee_haste=1.778 ± 0.112

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 | yes | Defender's Leather Helm (252455, -0.08 DPS) [crafted]; Veteran's Chain Helm (250498, -0.14 DPS, sim-verified) [crafted]; Crusader's Chain Helm (250502, -0.17 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; River Pride Choker (13087, -0.25 DPS) [world_drop]; Scout's Medallion (19537, -0.54 DPS) [rep] |
| shoulder | Mail Combat Spaulders (6404) | World drop [world_drop] | sim-verified (69.0 DPS) | yes | Barbaric Iron Shoulders (7913, -0.05 DPS) [crafted]; Elite Shoulders (4835, -0.08 DPS) [vendor]; Golden Scale Shoulders (3841, -0.69 DPS, sim-verified) [crafted] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Lambent Scale Cloak (4706, -0.08 DPS) [world_drop]; Slayer's Cape (14752, -0.08 DPS) [world_drop] |
| chest | Shining Silver Breastplate (2870) | Blacksmithing [crafted] | 28.0 | yes | Hard Gold Cuirass (250533, -0.25 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.28 DPS, sim-verified) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.30 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 | yes | Yorgen Bracers (13012, -0.16 DPS, sim-verified) [world_drop]; Bands of Serra'kis (6902, -0.17 DPS) [dungeon]; Grimtoll Wristguards (15459, -0.24 DPS) [quest] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | sim-verified (69.1 DPS) | yes | Warsong Gauntlets (16978, -0.08 DPS) [quest]; Bonefist Gauntlets (4465, -0.17 DPS) [world]; Fletcher's Gloves (7348, -0.74 DPS, sim-verified) [crafted] |
| waist | Girdle of Golem Strength (9405) (or Defiler's Plate Girdle (20207)) | World drop [world_drop] | 24.0 | yes | Defiler's Plate Girdle (20207, +0.00 DPS, sim-verified) [rep]; Defiler's Chain Girdle (20152, -0.00 DPS) [rep]; Defiler's Leather Girdle (20191, -0.00 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Golden Scale Leggings (3843, -0.17 DPS) [crafted]; Slayer's Pants (14757, -0.17 DPS) [world_drop]; Veteran's Silvered Chain Leggings (250523, -1.95 DPS, sim-verified) [crafted] |
| feet | Hard Gold Boots (250534) | Blacksmithing [crafted] | sim-verified (69.4 DPS) | yes | Glimmering Mail Greaves (4073, -0.08 DPS) [world_drop]; Slayer's Slippers (14756, -0.08 DPS) [world_drop]; Trouncing Boots (4464, -1.00 DPS, sim-verified) [world] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.4 | yes | Silverlaine's Family Seal (6321, -0.26 DPS) [dungeon]; Ironspine's Eye (7686, -0.30 DPS) [dungeon]; Insurgent's Band (272067, -0.31 DPS) [vendor] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 12.7 | yes | Ironspine's Eye (7686, -0.15 DPS) [dungeon]; Insurgent's Band (272067, -0.15 DPS) [vendor]; Silverlaine's Family Seal (6321, -0.35 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (52.3 DPS) | yes | Minor Recombobulator (4381, +0.00 DPS) [crafted]; Gnomish Universal Remote (7506, +0.00 DPS) [crafted]; Defiler's Talisman (21120, -2.38 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (68.2 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Corpsemaker (6687, -0.14 DPS) [dungeon]; Viscous Hammer (13045, -15.77 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Mail Combat Spaulders; back: Wolfmaster Cape; chest: Shining Silver Breastplate; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Girdle of Golem Strength; legs: Ferine Leggings; feet: Hard Gold Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 364, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7955 Copper Claymore; 7956 Bronze Warhammer; 7957 Bronze Greatsword; 7958 Bronze Battle Axe; 9362 Brilliant Gold Ring

### Band 40 (undead, 00000000000000000-0000000000000000-55223331211000210)

Set DPS (verified): 83.5. Weights run: 0.8s. Verify run: 0.8s. 511 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.228 ± 0.031, crit=4.287 ± 0.122, hit=2.725 ± 0.232, melee_haste=2.170 ± 0.198

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 86.0 | yes | Icemetal Barbute (10763, -0.70 DPS, sim-verified) [dungeon]; Hard Gold Coif (250537, -2.46 DPS) [crafted]; Chromite Barbute (8142, -2.52 DPS) [world_drop] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Ethereal Talisman (4430, -0.13 DPS, sim-verified) [quest]; Kaleidoscope Chain (13084, -0.22 DPS) [world_drop]; River Pride Choker (13087, -0.25 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 | yes | Shining Mithril Pauldrons (250541, -0.08 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.17 DPS) [world_drop]; Chromite Pauldrons (8144, -0.89 DPS, sim-verified) [world_drop] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Lambent Scale Cloak (4706, -0.08 DPS) [world_drop]; Warden's Cloak (14602, -0.08 DPS) [world_drop] |
| chest | Jouster's Chestplate (8157) | World drop [world_drop] | 30.0 | yes | Kolkar Marauder Chain (6773, +0.00 DPS, sim-verified) [quest]; Shining Silver Breastplate (2870, -0.08 DPS) [crafted]; Golden Scale Cuirass (3845, -0.08 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Ravager's Armguards (14770, -0.22 DPS) [world_drop]; Darkspear Armsplints (4132, -0.25 DPS) [quest]; Pugilist Bracers (4438, -0.29 DPS, sim-verified) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 80.0 | yes | Fletcher's Gloves (7348, -0.85 DPS) [crafted]; Ornate Mithril Gloves (7927, -0.85 DPS) [crafted]; Dragonscale Gauntlets (8347, -1.20 DPS, sim-verified) [crafted] |
| waist | Defiler's Chain Girdle (20153) | The Defilers [rep] | 68.0 | yes | Defiler's Leather Girdle (20192, +0.00 DPS, sim-verified) [rep]; Defiler's Plate Girdle (20206, -1.61 DPS) [rep]; Tharg's Shoelace (9705, -1.78 DPS) [quest] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 | yes | Firemane Leggings (13129, -0.29 DPS, sim-verified) [world_drop]; Orcish War Leggings (7929, -0.34 DPS) [crafted]; Symbolic Legplates (14829, -0.45 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 27.6 | yes | Prowler's Leather Shoes (252465, -0.07 DPS, sim-verified) [crafted]; Skirmisher's Mail Boots (252564, -0.24 DPS) [crafted]; Obsidian Greaves (13068, -0.26 DPS) [world_drop] |
| finger1 | Assault Band (13095) | World drop [world_drop] | 20.0 | yes | Thunderbrow Ring (13097, -0.14 DPS) [world_drop]; Suspicious Spare Part (274754, -0.25 DPS) [vendor]; Insurgent's Band (272066, -0.34 DPS) [vendor] |
| finger2 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 17.8 | yes | Suspicious Spare Part (274754, -0.16 DPS) [vendor]; Legionnaire's Band (19513, -0.19 DPS) [rep]; Thunderbrow Ring (13097, -0.25 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (83.6 DPS) | yes | Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Fiery War Axe (870, -10.12 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; shoulder: Hard Gold Pauldrons; chest: Jouster's Chestplate; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Chain Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Assault Band; finger2: Legionnaire's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; ranged: Libram of Invocation

No-known-source sample (15 of 511, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7748 Forcestone Buckler; 7922 Steel Plate Helm

### Band 50 (undead, 00000000000000000-5500000000000000-55223331211000210)

Set DPS (verified): 99.3. Weights run: 0.9s. Verify run: 0.8s. 682 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.231 ± 0.026, crit=5.537 ± 0.167, hit=3.474 ± 0.300, melee_haste=2.628 ± 0.325

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Plate Helm (220803) | Lady Palanseer [vendor] | 138.3 | yes | Raging Berserker's Helm (7719, -0.91 DPS, sim-verified) [dungeon]; Ornate Mithril Helm (7937, -1.75 DPS) [crafted]; Eye of Theradras (17715, -2.61 DPS) [dungeon] |
| neck | Woven Ivy Necklace (19159) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 14.1 | yes | Skibi's Pendant (13089, -0.05 DPS) [world_drop]; Ethereal Talisman (4430, -0.14 DPS) [quest]; Ghostshard Talisman (7731, -0.29 DPS, sim-verified) [dungeon] |
| shoulder | Blood Guard's Plate Pauldrons (220796) | Lady Palanseer [vendor] | 99.5 | yes | Officer's Pauldrons (250576, -0.29 DPS, sim-verified) [crafted]; Earthslag Shoulders (11632, -3.16 DPS) [dungeon]; Wyrmslayer Spaulders (13066, -3.17 DPS) [world_drop] |
| back | Bloodlust Cape (14801) | World drop [world_drop] | 18.0 | yes | Battlehard Cape (11858, -0.34 DPS) [quest]; Wildhunter Cloak (16658, -0.34 DPS) [quest]; Wolfmaster Cape (6314, -0.48 DPS, sim-verified) [dungeon] |
| chest | Stone Guard's Plate Armor (220801) | Lady Palanseer [vendor] | 107.5 | yes | Ornate Mithril Breastplate (7935, -1.79 DPS, sim-verified) [crafted]; Warforged Chestplate (11195, -2.56 DPS) [quest]; Valorous Chestguard (8274, -2.90 DPS) [world_drop] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 | yes | Giantslayer Bracers (13076, -0.29 DPS) [world_drop]; Branded Leather Bracers (19508, -0.34 DPS) [dungeon]; Officer's Wristguards (250581, -1.05 DPS, sim-verified) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 97.5 | yes | Fletcher's Gloves (7348, -0.86 DPS) [crafted]; Ornate Mithril Gloves (7927, -0.86 DPS) [crafted]; Dragonscale Gauntlets (8347, -0.92 DPS, sim-verified) [crafted] |
| waist | Defiler's Chain Girdle (20151) (or Defiler's Leather Girdle (20193)) | The Defilers [rep] | 97.5 | yes | Defiler's Leather Girdle (20193, +0.00 DPS, sim-verified) [rep]; Defiler's Plate Girdle (20205, -0.00 DPS) [rep]; Defiler's Chain Girdle (20153, -0.52 DPS) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | 155.0 | yes | Stone Guard's Plate Leggings (220798, +0.00 DPS, sim-verified) [vendor]; Golem Shard Leggings (13074, -4.77 DPS) [world_drop]; Scarlet Leggings (10330, -4.86 DPS) [dungeon] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 34.7 | yes | Officer's Sabatons (250561, -0.29 DPS) [crafted]; Officer's Boots (250546, -0.31 DPS) [crafted]; Prowler's Leather Boots (252468, -2.41 DPS, sim-verified) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 54.7 | yes | Legionnaire's Band (19511, -1.40 DPS) [rep]; Assault Band (13095, -1.49 DPS) [world_drop]; Band of Allegiance (18585, -1.58 DPS) [quest] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 | yes | Legionnaire's Band (19511, +0.00 DPS, sim-verified) [rep]; Assault Band (13095, -0.17 DPS) [world_drop]; Band of Allegiance (18585, -0.26 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (95.9 DPS) | yes | Tidal Charm (1404, -2.85 DPS) [vendor]; Guardian Talisman (1490, -2.85 DPS) [quest]; Ankh of Life (1713, -3.13 DPS, sim-verified) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (95.9 DPS) | yes | Ankh of Life (1713, -0.66 DPS, sim-verified) [world_drop]; Tidal Charm (1404, -1.34 DPS) [vendor]; Guardian Talisman (1490, -1.34 DPS) [quest] |
| main_hand | Blight (7959) | Blacksmithing [crafted] | sim-verified (98.9 DPS) | yes | Thorium Greatmace (250613, +0.00 DPS) [crafted]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor]; Warmonger (13052, -3.09 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Blood Guard's Plate Helm; neck: Woven Ivy Necklace; shoulder: Blood Guard's Plate Pauldrons; back: Bloodlust Cape; chest: Stone Guard's Plate Armor; wrist: Bracers of the Stone Princess; waist: Defiler's Chain Girdle; legs: Stormshroud Pants; feet: Greaves of Withering Despair; finger1: Blackstone Ring; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Blight

No-known-source sample (15 of 682, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60 (undead, 00000000000000000-5532500000000000-55223331211000210)

Set DPS (verified): 182.3. Weights run: 0.9s. Verify run: 0.9s. 1283 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.339 ± 0.051, crit=7.362 ± 0.238, hit=9.891 ± 0.551, melee_haste=3.638 ± 0.629

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Inquisition Helmet (240027) | Leonid Barthalomew the Revered [vendor] | sim-verified (180.1 DPS) | yes | Mask of the Unforgiven (13404, -0.32 DPS) [dungeon]; Bloodvine Goggles (19999, -0.32 DPS) [crafted]; Lionheart Helm (12640, -4.64 DPS, sim-verified) [crafted] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (174.7 DPS) | yes | Beads of Ogre Might (22150, -0.18 DPS) [quest]; Blazefury Medallion (17111, -2.24 DPS, sim-verified) [world]; Amulet of the Darkmoon (19491, -4.40 DPS) [quest] |
| shoulder | Inquisition Spaulders (240021) | Leonid Barthalomew the Revered [vendor] | sim-verified (177.8 DPS) | yes | Inquisition Epaulets (246061, -0.18 DPS) [vendor]; Inquisition Shoulderplates (240025, -2.32 DPS, sim-verified) [vendor]; Darkspear Spaulders (272108, -3.46 DPS) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 126.9 | yes | Earthweave Cloak (21187, -0.85 DPS, sim-verified) [quest]; Chromatic Cloak (18509, -1.04 DPS) [crafted]; Arcanoweave Cloak (272411, -1.22 DPS) [vendor] |
| chest | Inquisition Cuirass (246060) | Leonid Barthalomew the Revered [vendor] | 502.9 | yes | Inquisition Breastplate (240030, -2.78 DPS, sim-verified) [vendor]; Dawn Armor (252483, -12.61 DPS) [crafted]; Bloodsoul Breastplate (19690, -12.84 DPS) [crafted] |
| wrist | Inquisition Armbraces (246056) | Leonid Barthalomew the Revered [vendor] | 296.7 | yes | Inquisition Vambraces (240023, -0.36 DPS, sim-verified) [vendor]; Inquisition Bracers (240031, -7.42 DPS) [vendor]; Vambraces of the Sadist (13400, -7.85 DPS) [dungeon] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | 202.0 | yes | Primal Batskin Gloves (19686, -0.65 DPS, sim-verified) [crafted]; Inquisition Gloves (240028, -1.00 DPS) [vendor]; Chromatic Gauntlets (19157, -2.40 DPS) [crafted] |
| waist | Inquisition Belt (240024) | Leonid Barthalomew the Revered [vendor] | 278.1 | yes | Inquisition Cord (246059, +0.00 DPS, sim-verified) [vendor]; Radiant Girdle of the Dawn (227814, -5.29 DPS) [vendor]; Defiler's Plate Girdle (20204, -6.17 DPS) [rep] |
| legs | Inquisition Leggings (240026) | Leonid Barthalomew the Revered [vendor] | 405.0 | yes | Sentinel's Plate Legguards (237825, -2.01 DPS) [vendor]; Sentinel's Chain Leggings (237819, -3.85 DPS) [vendor]; Titanic Leggings (22385, -4.20 DPS, sim-verified) [crafted] |
| feet | Inquisition Greaves (240029) | Leonid Barthalomew the Revered [vendor] | 276.0 | yes | Inquisition Stompers (246057, +0.00 DPS, sim-verified) [vendor]; Inquisition Boots (240022, -7.56 DPS) [vendor]; Fine Dawn Treaders (227815, -7.58 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (174.7 DPS) | yes | Wrath of Cenarius (21190, -1.83 DPS, sim-verified) [quest]; Signet Ring of the Bronze Dragonflight (234034, -3.11 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234030, -3.28 DPS) [vendor] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (174.7 DPS) | yes | Wrath of Cenarius (21190, -1.55 DPS, sim-verified) [quest]; Signet Ring of the Bronze Dragonflight (234034, -2.93 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234030, -3.11 DPS) [vendor] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (171.7 DPS) | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (174.7 DPS) | yes | Frozen Heart of the Mountain (249469, -0.97 DPS) [crafted]; Darkmoon Card: Heroism (19287, -4.10 DPS, sim-verified) [quest]; Tidal Charm (1404, -4.86 DPS) [vendor] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (174.7 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [vendor]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Sulfuron Hammer (17193, -1.42 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Inquisition Helmet; neck: Medallion of the Dawn; shoulder: Inquisition Spaulders; back: Howler's Furs; chest: Inquisition Cuirass; wrist: Inquisition Armbraces; hands: Stormshroud Gloves; waist: Inquisition Belt; legs: Inquisition Leggings; feet: Inquisition Greaves; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: The Unstoppable Force; ranged: Libram of Fervor

No-known-source sample (15 of 1283, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

