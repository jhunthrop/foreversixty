# Leveling BiS: Retribution

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-0000000000000000-55100000000000000)

Set DPS (verified): 34.6. Weights run: 1.3s. Verify run: 1.2s. 218 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.133 ± 0.018, crit=0.172 ± 0.005 per rating point (14 rating = 1%, 2.410 per %), hit=0.187 ± 0.003 per rating point (10 rating = 1%, 1.869 per %), melee_haste=-2.419 ± 0.186

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.69 DPS) | yes | Defender's Leather Hood (252447, -0.16 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.65 DPS) [crafted]; Brawler's Leather Hood (252504, -0.66 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 0.8 attack_power points (0.03 DPS) | yes | Erudite's Amulet (277204, -0.04 DPS, sim-verified) [quest] |
| shoulder | Silvered Bronze Shoulders (3481) (or Rough Bronze Shoulders (3480)) | Blacksmithing [crafted] | 6.0 attack_power points (0.21 DPS) | yes | Rough Bronze Shoulders (3480, +0.00 DPS, sim-verified) [crafted]; Serpent's Shoulders (5404, -0.18 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Grave Shroud (279865, -0.01 DPS, sim-verified) [quest]; Catacomb Cloak (279899, -0.07 DPS) [quest]; Dark Leather Cloak (2316, -0.12 DPS) [crafted] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.69 DPS) | yes | Veteran's Chain Shirt (250488, -0.17 DPS, sim-verified) [crafted]; Defender's Leather Armor (252434, -0.19 DPS) [crafted]; Totemic Leather Armor (252435, -0.21 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Cryptwalker Bracers (280095, -0.08 DPS, sim-verified) [quest]; Bravo's Armbands (270015, -0.19 DPS) [quest]; Bristlebark Bindings (14569, -0.19 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.0 attack_power points (0.55 DPS) | yes | Gold-flecked Gloves (5195, -0.08 DPS, sim-verified) [dungeon]; Polar Gauntlets (7606, -0.14 DPS) [quest]; Blackened Defias Gloves (10401, -0.14 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (0.62 DPS) | yes | Cobrahn's Grasp (6460, -0.11 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.21 DPS) [world]; Warchief's Girdle (5750, -0.28 DPS) [world] |
| legs | Chausses of Westfall (6087) | The Defias Brotherhood [quest] | 22.0 attack_power points (0.76 DPS) | yes | Veteran's Chain Leggings (250493, -0.07 DPS, sim-verified) [crafted]; Defender's Leather Pants (252445, -0.12 DPS) [crafted]; Totemic Leather Pants (252446, -0.14 DPS) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.7 attack_power points (0.37 DPS) | yes | Veteran's Boots (250503, -0.02 DPS, sim-verified) [crafted]; Guard's Boots (250504, -0.02 DPS) [crafted]; Defender's Leather Boots (252441, -0.02 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.5 attack_power points (0.30 DPS) | yes | Loop of Sacrifice (281673, -0.09 DPS) [quest]; The 1 Ring (8350, -0.22 DPS) [world]; Ring of the Moon (12052, -0.23 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Loop of Sacrifice (281673, +0.00 DPS, sim-verified) [quest]; The 1 Ring (8350, -0.20 DPS) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Verigan's Fist (6953) | The Test of Righteousness [quest] | 358.7 attack_power points (12.43 DPS) | yes | Forsaken Greataxe (251533, -1.93 DPS) [quest]; Smite's Mighty Hammer (7230, -2.09 DPS) [dungeon]; The Axe of Severing (23171, -19.88 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Silvered Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: Verigan's Fist

No-known-source sample (15 of 218, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 30 (human, 00000000000000000-0000000000000000-55223310000000000)

Set DPS (verified): 71.6. Weights run: 1.4s. Verify run: 1.4s. 382 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.118 ± 0.014, crit=0.231 ± 0.007 per rating point (14 rating = 1%, 3.235 per %), hit=0.226 ± 0.004 per rating point (10 rating = 1%, 2.259 per %), melee_haste=1.778 ± 0.112

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 attack_power points (1.08 DPS) | yes | Defender's Leather Helm (252455, -0.08 DPS) [crafted]; Veteran's Chain Helm (250498, -0.14 DPS, sim-verified) [crafted]; Crusader's Chain Helm (250502, -0.17 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.58 DPS) | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; River Pride Choker (13087, -0.25 DPS) [world_drop]; Sentinel's Medallion (19541, -0.54 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 14.0 attack_power points (0.58 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS, sim-verified) [world_drop]; Barbaric Iron Shoulders (7913, -0.05 DPS) [crafted]; Elite Shoulders (4835, -0.08 DPS) [vendor] |
| back | Sergeant Major's Cape (16315) | PvP rank 9 · Sergeant Major · Alliance [pvp] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Lambent Scale Cloak (4706, -0.02 DPS) [world_drop]; Slayer's Cape (14752, -0.02 DPS) [world_drop]; Wolfmaster Cape (6314, -0.71 DPS, sim-verified) [dungeon] |
| chest | Shining Silver Breastplate (2870) | Blacksmithing [crafted] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Barbaric Iron Breastplate (7914, -0.17 DPS) [crafted]; Hard Gold Cuirass (250533, -0.25 DPS) [crafted]; Avenger's Armor (1488, -1.93 DPS, sim-verified) [dungeon] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.67 DPS) | yes | Yorgen Bracers (13012, -0.16 DPS, sim-verified) [world_drop]; Bands of Serra'kis (6902, -0.17 DPS) [dungeon]; Technician's Bracers (270042, -0.25 DPS) [quest] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (0.92 DPS) | yes | Bonefist Gauntlets (4465, -0.17 DPS) [world]; Mail Combat Gauntlets (4075, -0.23 DPS) [world_drop]; The Frozen Clutch (23170, -1.51 DPS, sim-verified) [dungeon] |
| waist | Girdle of Golem Strength (9405) (or Highlander's Plate Girdle (20126)) | World drop [world_drop] | 24.0 attack_power points (1.00 DPS) | yes | Highlander's Plate Girdle (20126, +0.00 DPS, sim-verified) [rep]; Highlander's Chain Girdle (20090, -0.00 DPS) [rep]; Highlander's Leather Girdle (20117, -0.00 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.08 DPS) | yes | Golden Scale Leggings (3843, -0.17 DPS) [crafted]; Slayer's Pants (14757, -0.17 DPS) [world_drop]; Veteran's Silvered Chain Leggings (250523, -2.29 DPS, sim-verified) [crafted] |
| feet | Hard Gold Boots (250534) | Blacksmithing [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Glimmering Mail Greaves (4073, -0.08 DPS) [world_drop]; Slayer's Slippers (14756, -0.08 DPS) [world_drop]; Trouncing Boots (4464, -1.02 DPS, sim-verified) [world] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.4 attack_power points (0.68 DPS) | yes | Tiger Band (6749, -0.18 DPS) [quest]; Silverlaine's Family Seal (6321, -0.26 DPS) [dungeon]; Ironspine's Eye (7686, -0.30 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 12.7 attack_power points (0.53 DPS) | yes | Silverlaine's Family Seal (6321, -0.11 DPS) [dungeon]; Ironspine's Eye (7686, -0.15 DPS) [dungeon]; Tiger Band (6749, -1.04 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Talisman of Arathor (21119, -2.93 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 0.0 attack_power points (0.00 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Corpsemaker (6687, -0.14 DPS) [dungeon]; Viscous Hammer (13045, -15.75 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Sergeant Major's Cape; chest: Shining Silver Breastplate; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Girdle of Golem Strength; legs: Ferine Leggings; feet: Hard Gold Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; trinket1: Darkspear Voodoo Seal; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 382, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants

### Band 40 (human, 00000000000000000-0000000000000000-55223331211000210)

Set DPS (verified): 85.8. Weights run: 1.4s. Verify run: 1.4s. 538 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.228 ± 0.031, crit=0.306 ± 0.009 per rating point (14 rating = 1%, 4.287 per %), hit=0.312 ± 0.005 per rating point (10 rating = 1%, 3.119 per %), melee_haste=2.170 ± 0.198

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 30.3 attack_power points (1.29 DPS) | yes | Hard Gold Coif (250537, -0.10 DPS) [crafted]; Chromite Barbute (8142, -0.15 DPS) [world_drop]; Icemetal Barbute (10763, -0.44 DPS, sim-verified) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Kaleidoscope Chain (13084, -0.22 DPS) [world_drop]; Gazlowe's Charm (13088, -0.25 DPS) [world_drop]; Zealous Shadowshard Pendant (17772, -1.01 DPS, sim-verified) [quest] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 attack_power points (0.93 DPS) | yes | Shining Mithril Pauldrons (250541, -0.08 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.17 DPS) [world_drop]; Chromite Pauldrons (8144, -0.91 DPS, sim-verified) [world_drop] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 13.4 attack_power points (0.57 DPS) | yes | Wolfmaster Cape (6314, -0.14 DPS) [dungeon]; Sergeant Major's Cape (16315, -0.19 DPS) [pvp]; Dark Hooded Cape (5257, -1.60 DPS, sim-verified) [world] |
| chest | Jouster's Chestplate (8157) (or Avenger's Armor (1488)) | World drop [world_drop] | 30.0 attack_power points (1.27 DPS) | yes | Kolkar Marauder Chain (6773, -0.02 DPS) [quest]; Shining Mithril Breastplate (250540, -0.08 DPS) [crafted]; Avenger's Armor (1488, -2.16 DPS, sim-verified) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.85 DPS) | yes | Ravager's Armguards (14770, -0.22 DPS) [world_drop]; Pugilist Bracers (4438, -0.29 DPS, sim-verified) [dungeon]; Yorgen Bracers (13012, -0.31 DPS) [world_drop] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 attack_power points (1.36 DPS) | yes | Truesilver Gauntlets (7938, +0.00 DPS, sim-verified) [crafted]; Scarlet Gauntlets (10331, -0.24 DPS) [dungeon]; Reticulated Bone Gauntlets (9435, -0.25 DPS) [world_drop] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (1.27 DPS) | yes | Highlander's Plate Girdle (20125, +0.00 DPS, sim-verified) [rep]; Boar Champion's Belt (10768, -0.00 DPS) [dungeon]; Highlander's Chain Girdle (20090, -0.25 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.78 DPS) | yes | Firemane Leggings (13129, -0.29 DPS, sim-verified) [world_drop]; Orcish War Leggings (7929, -0.34 DPS) [crafted]; Symbolic Legplates (14829, -0.45 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 27.6 attack_power points (1.17 DPS) | yes | Prowler's Leather Shoes (252465, +0.00 DPS, sim-verified) [crafted]; Skirmisher's Mail Boots (252564, -0.24 DPS) [crafted]; Obsidian Greaves (13068, -0.26 DPS) [world_drop] |
| finger1 | Assault Band (13095) (or Mark of Kern (2262)) | World drop [world_drop] | 20.0 attack_power points (0.85 DPS) | yes | Thunderbrow Ring (13097, -0.14 DPS) [world_drop]; Suspicious Spare Part (274754, -0.25 DPS) [vendor]; Mark of Kern (2262, -1.65 DPS, sim-verified) [dungeon] |
| finger2 | Protector's Band (19515) | Silverwing Sentinels [rep] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Suspicious Spare Part (274754, -0.16 DPS) [vendor]; Mark of Kern (2262, -1.49 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 0.0 attack_power points (0.00 DPS) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Bonebiter (6830, +0.00 DPS) [quest]; Fiery War Axe (870, -10.04 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Jouster's Chestplate; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Highlander's Leather Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Assault Band; finger2: Protector's Band

No-known-source sample (15 of 538, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 50 (human, 00000000000000000-5500000000000000-55223331211000210)

Set DPS (verified): 90.9. Weights run: 1.5s. Verify run: 1.3s. 707 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.231 ± 0.026, crit=0.395 ± 0.012 per rating point (14 rating = 1%, 5.537 per %), hit=0.389 ± 0.006 per rating point (10 rating = 1%, 3.895 per %), melee_haste=2.628 ± 0.325

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 48.0 attack_power points (2.06 DPS) | yes | Fury Visor (20521, +0.00 DPS, sim-verified) [quest]; Ebon Mask (19984, -0.50 DPS) [quest]; Bloomsprout Headpiece (17767, -0.52 DPS) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.86 DPS) | yes | Ghostshard Talisman (7731, +0.00 DPS, sim-verified) [dungeon]; Skibi's Pendant (13089, -0.30 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.48 DPS) [world_drop] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 27.8 attack_power points (1.20 DPS) | yes | Knight-Lieutenant's Plate Pauldrons (220795, +0.00 DPS, sim-verified) [vendor]; Earthslag Shoulders (11632, -0.08 DPS) [dungeon]; Wyrmslayer Spaulders (13066, -0.09 DPS) [world_drop] |
| back | Bloodlust Cape (14801) | World drop [world_drop] | 18.0 attack_power points (0.77 DPS) | yes | Sergeant Major's Cape (16336, -0.20 DPS) [pvp]; Dark Hooded Cape (5257, -0.33 DPS) [world]; Blackveil Cape (11626, -0.97 DPS, sim-verified) [dungeon] |
| chest | Warforged Chestplate (11195) | Tremors of the Earth [quest] | 48.0 attack_power points (2.06 DPS) | yes | Mixologist's Tunic (12793, -0.41 DPS) [dungeon]; Coldmetal Guard (274758, -0.43 DPS) [vendor]; Valorous Chestguard (8274, -0.44 DPS, sim-verified) [world_drop] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 attack_power points (1.20 DPS) | yes | Officer's Wristguards (250581, -0.20 DPS) [crafted]; Giantslayer Bracers (13076, -0.29 DPS) [world_drop]; Runed Golem Shackles (12550, -2.75 DPS, sim-verified) [dungeon] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 46.3 attack_power points (1.99 DPS) | yes | Gauntlets of Divinity (7724, -0.61 DPS) [dungeon]; Truesilver Gauntlets (7938, -0.61 DPS) [crafted]; Officer's Gloves (250551, -0.69 DPS, sim-verified) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (1.98 DPS) | yes | Atal'alarion's Tusk Ring (10798, +0.00 DPS, sim-verified) [dungeon]; Belt of the Gladiator (13134, -0.43 DPS) [world_drop]; Prowler's Leather Waistguard (252473, -0.65 DPS) [crafted] |
| legs | Golem Shard Leggings (13074) | World drop [world_drop] | 44.0 attack_power points (1.89 DPS) | yes | Scarlet Leggings (10330, -0.11 DPS, sim-verified) [dungeon]; Silvershell Leggings (10633, -0.17 DPS) [dungeon]; Elemental Rockridge Leggings (17711, -0.17 DPS) [dungeon] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | sim-verified (90.9 DPS) | yes | Officer's Sabatons (250561, -0.02 DPS) [crafted]; Officer's Boots (250546, -0.04 DPS) [crafted]; Battlechaser's Greaves (12555, -1.91 DPS, sim-verified) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 23.9 attack_power points (1.03 DPS) | yes | Mark of Kern (2262, -0.17 DPS) [dungeon]; Assault Band (13095, -0.17 DPS) [world_drop]; Thunderbrow Ring (13097, -0.31 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 22.1 attack_power points (0.95 DPS) | yes | Mark of Kern (2262, -0.09 DPS) [dungeon]; Protector's Band (19515, -0.18 DPS) [rep]; Assault Band (13095, -0.51 DPS, sim-verified) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | 0.0 attack_power points (0.00 DPS) | yes | Darkspear Raider's Reaper (272080, -0.42 DPS) [vendor]; Taran Icebreaker (2915, -0.48 DPS) [world_drop]; Blight (7959, -1.30 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Embrace of the Lycan; neck: Zealous Shadowshard Pendant; shoulder: Officer's Pauldrons; back: Bloodlust Cape; chest: Warforged Chestplate; wrist: Bracers of the Stone Princess; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Golem Shard Leggings; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Frozen Heart of the Mountain; main_hand: Thorium Greatmace

No-known-source sample (15 of 707, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60 (human, 00000000000000000-5532500000000000-55223331211000210)

Set DPS (verified): 179.3. Weights run: 1.4s. Verify run: 1.4s. 1554 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.339 ± 0.051, crit=0.526 ± 0.017 per rating point (14 rating = 1%, 7.362 per %), hit=0.466 ± 0.008 per rating point (10 rating = 1%, 4.660 per %), melee_haste=3.638 ± 0.629

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Inquisition Helmet (240027) | Leonid Barthalomew the Revered [vendor] | 116.7 attack_power points (5.10 DPS) | yes | Lionheart Helm (12640, -2.48 DPS) [crafted]; Field Marshal's Lamellar Headguard (231648, -2.86 DPS) [vendor]; Inquisition Crown (240035, -5.27 DPS, sim-verified) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (179.3 DPS) | yes | Will of the Martyr (17044, -0.06 DPS) [quest]; Beads of Ogre Might (22150, -0.12 DPS) [quest]; Imperial Jewel (11933, -3.08 DPS, sim-verified) [dungeon] |
| shoulder | Inquisition Shoulderplates (240025) | Leonid Barthalomew the Revered [vendor] | 84.0 attack_power points (3.67 DPS) | yes | Field Marshal's Lamellar Shoulders (231651, -1.84 DPS) [vendor]; Highlander's Plate Spaulders (20057, -1.85 DPS) [rep]; Inquisition Pauldrons (240033, -5.57 DPS, sim-verified) [vendor] |
| back | Cloak of the Honor Guard (20073) | The League of Arathor [rep] | 35.7 attack_power points (1.56 DPS) | yes | Howler's Furs (272414, -0.13 DPS) [vendor]; Shadewood Cloak (18328, -0.42 DPS) [dungeon]; Shroud of Domination (22337, -2.59 DPS, sim-verified) [dungeon] |
| chest | Inquisition Breastplate (240030) | Leonid Barthalomew the Revered [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Obsidian Mail Tunic (22191, -1.46 DPS) [crafted]; Timbermaw Tunic (252484, -1.63 DPS) [crafted]; Breastplate of Undead Slaying (23087, -9.75 DPS, sim-verified) [world] |
| wrist | Inquisition Vambraces (240023) | Leonid Barthalomew the Revered [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Berserker Bracers (19578, -1.11 DPS) [rep]; Windtalker's Wristguards (19582, -1.23 DPS) [rep]; Bracers of Undead Slaying (23090, -6.99 DPS, sim-verified) [world] |
| hands | Inquisition Gloves (240028) | Leonid Barthalomew the Revered [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Raider Gauntlets (272095, -1.12 DPS) [vendor]; Chromatic Gauntlets (19157, -1.40 DPS) [crafted]; Razor Gauntlets (18326, -6.49 DPS, sim-verified) [dungeon] |
| waist | Inquisition Belt (240024) | Leonid Barthalomew the Revered [vendor] | 86.7 attack_power points (3.79 DPS) | yes | Radiant Girdle of the Dawn (227814, -1.11 DPS) [vendor]; Ferocity of the Timbermaw (227805, -1.27 DPS) [vendor]; Dense Timbermaw Belt (227807, -3.56 DPS, sim-verified) [vendor] |
| legs | Inquisition Leggings (240026) | Leonid Barthalomew the Revered [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Titanic Leggings (22385, -1.87 DPS) [crafted]; Sentinel's Plate Legguards (237825, -2.01 DPS) [vendor]; Cloudkeeper Legplates (14554, -12.02 DPS, sim-verified) [world_drop] |
| feet | Inquisition Greaves (240029) | Leonid Barthalomew the Revered [vendor] | 86.0 attack_power points (3.76 DPS) | yes | Pads of the Dread Wolf (13210, -2.01 DPS) [dungeon]; Clutchlord's Stompers (275627, -2.10 DPS) [crafted]; Scalegut Treaders (275618, -5.66 DPS, sim-verified) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Signet Ring of the Bronze Dragonflight (234030, -0.17 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234026, -0.26 DPS) [vendor]; Naglering (11669, -5.62 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | 0.0 attack_power points (0.00 DPS) | yes | Band of the Ogre King (18522, -0.00 DPS) [dungeon]; Protector's Band (19514, -0.01 DPS) [rep]; Naglering (11669, -4.91 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | 0.0 attack_power points (0.00 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Shard of the Fallen Star (21891, -5.75 DPS, sim-verified) [world_drop] |
| main_hand | Sulfuron Hammer (17193) | Blacksmithing [crafted] | 0.0 attack_power points (0.00 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Nightfall (19169, -0.71 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Inquisition Helmet; neck: Medallion of the Dawn; shoulder: Inquisition Shoulderplates; back: Cloak of the Honor Guard; chest: Inquisition Breastplate; wrist: Inquisition Vambraces; hands: Inquisition Gloves; waist: Inquisition Belt; legs: Inquisition Leggings; feet: Inquisition Greaves; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Talisman of Ascendance; trinket2: Darkmoon Card: Maelstrom; main_hand: Sulfuron Hammer

No-known-source sample (15 of 1554, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

## Horde

### Band 20 (undead, 00000000000000000-0000000000000000-55100000000000000)

Set DPS (verified): 32.3. Weights run: 1.3s. Verify run: 1.4s. 217 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.133 ± 0.018, crit=0.172 ± 0.005 per rating point (14 rating = 1%, 2.410 per %), hit=0.187 ± 0.003 per rating point (10 rating = 1%, 1.869 per %), melee_haste=-2.419 ± 0.186

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.69 DPS) | yes | Defender's Leather Hood (252447, -0.14 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.65 DPS) [crafted]; Brawler's Leather Hood (252504, -0.66 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 0.8 attack_power points (0.03 DPS) | yes | Erudite's Amulet (277204, -0.01 DPS, sim-verified) [quest] |
| shoulder | Silvered Bronze Shoulders (3481) (or Rough Bronze Shoulders (3480)) | Blacksmithing [crafted] | 6.0 attack_power points (0.21 DPS) | yes | Rough Bronze Shoulders (3480, -0.10 DPS, sim-verified) [crafted]; Serpent's Shoulders (5404, -0.18 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Grave Shroud (279865, -0.04 DPS, sim-verified) [quest]; Catacomb Cloak (279899, -0.07 DPS) [quest]; Subterranean Cape (14149, -0.07 DPS) [dungeon] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.69 DPS) | yes | Veteran's Chain Shirt (250488, -0.18 DPS, sim-verified) [crafted]; Defender's Leather Armor (252434, -0.19 DPS) [crafted]; Totemic Leather Armor (252435, -0.21 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Cryptwalker Bracers (280095, -0.07 DPS, sim-verified) [quest]; Raptorcrest Bracers (270010, -0.14 DPS) [quest]; Bristlebark Bindings (14569, -0.19 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.0 attack_power points (0.55 DPS) | yes | Gold-flecked Gloves (5195, -0.07 DPS, sim-verified) [dungeon]; Blackened Defias Gloves (10401, -0.14 DPS) [dungeon]; Dagmire Gauntlets (6481, -0.19 DPS) [quest] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (0.62 DPS) | yes | Cobrahn's Grasp (6460, -0.12 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.21 DPS) [world]; Warchief's Girdle (5750, -0.28 DPS) [world] |
| legs | Veteran's Chain Leggings (250493) | Blacksmithing [crafted] | 18.7 attack_power points (0.65 DPS) | yes | Defender's Leather Pants (252445, -0.01 DPS, sim-verified) [crafted]; Totemic Leather Pants (252446, -0.02 DPS) [crafted]; Hulking Leggings (14748, -0.08 DPS) [world_drop] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.7 attack_power points (0.37 DPS) | yes | Veteran's Boots (250503, -0.01 DPS, sim-verified) [crafted]; Guard's Boots (250504, -0.02 DPS) [crafted]; Defender's Leather Boots (252441, -0.02 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.5 attack_power points (0.30 DPS) | yes | Loop of Sacrifice (281673, -0.09 DPS) [quest]; The 1 Ring (8350, -0.22 DPS) [world]; Ring of the Moon (12052, -0.23 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Loop of Sacrifice (281673, -0.04 DPS, sim-verified) [quest]; The 1 Ring (8350, -0.20 DPS) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hammerbone (270018) | Leaders of the Fang [quest] | sim-verified (32.3 DPS) | yes | Forsaken Greataxe (251533, -0.11 DPS) [quest]; Smite's Mighty Hammer (7230, -0.27 DPS) [dungeon]; The Axe of Severing (23171, -17.69 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Silvered Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: Hammerbone

No-known-source sample (15 of 217, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5968 Rugged Boots; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield

### Band 30 (undead, 00000000000000000-0000000000000000-55223310000000000)

Set DPS (verified): 70.1. Weights run: 1.4s. Verify run: 1.4s. 381 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.118 ± 0.014, crit=0.231 ± 0.007 per rating point (14 rating = 1%, 3.235 per %), hit=0.226 ± 0.004 per rating point (10 rating = 1%, 2.259 per %), melee_haste=1.778 ± 0.112

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 attack_power points (1.08 DPS) | yes | Defender's Leather Helm (252455, -0.08 DPS) [crafted]; Veteran's Chain Helm (250498, -0.14 DPS, sim-verified) [crafted]; Crusader's Chain Helm (250502, -0.17 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.58 DPS) | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; River Pride Choker (13087, -0.25 DPS) [world_drop]; Scout's Medallion (19537, -0.54 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 14.0 attack_power points (0.58 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS, sim-verified) [world_drop]; Barbaric Iron Shoulders (7913, -0.05 DPS) [crafted]; Elite Shoulders (4835, -0.08 DPS) [vendor] |
| back | Wildhunter Cloak (16658) (or Wolfmaster Cape (6314)) | The Hunt Completed [quest] | 10.0 attack_power points (0.42 DPS) | yes | Wolfmaster Cape (6314, +0.00 DPS, sim-verified) [dungeon]; Lambent Scale Cloak (4706, -0.08 DPS) [world_drop]; Slayer's Cape (14752, -0.08 DPS) [world_drop] |
| chest | Shining Silver Breastplate (2870) | Blacksmithing [crafted] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Barbaric Iron Breastplate (7914, -0.17 DPS) [crafted]; Hard Gold Cuirass (250533, -0.25 DPS) [crafted]; Avenger's Armor (1488, -1.92 DPS, sim-verified) [dungeon] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.67 DPS) | yes | Yorgen Bracers (13012, -0.16 DPS, sim-verified) [world_drop]; Bands of Serra'kis (6902, -0.17 DPS) [dungeon]; Grimtoll Wristguards (15459, -0.24 DPS) [quest] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (0.92 DPS) | yes | The Frozen Clutch (23170, -0.08 DPS) [dungeon]; Warsong Gauntlets (16978, -0.14 DPS, sim-verified) [quest]; Bonefist Gauntlets (4465, -0.17 DPS) [world] |
| waist | Girdle of Golem Strength (9405) (or Defiler's Plate Girdle (20207)) | World drop [world_drop] | 24.0 attack_power points (1.00 DPS) | yes | Defiler's Plate Girdle (20207, +0.00 DPS, sim-verified) [rep]; Defiler's Chain Girdle (20152, -0.00 DPS) [rep]; Defiler's Leather Girdle (20191, -0.00 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.08 DPS) | yes | Golden Scale Leggings (3843, -0.17 DPS) [crafted]; Slayer's Pants (14757, -0.17 DPS) [world_drop]; Veteran's Silvered Chain Leggings (250523, -1.93 DPS, sim-verified) [crafted] |
| feet | Hard Gold Boots (250534) | Blacksmithing [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Glimmering Mail Greaves (4073, -0.08 DPS) [world_drop]; Slayer's Slippers (14756, -0.08 DPS) [world_drop]; Trouncing Boots (4464, -1.02 DPS, sim-verified) [world] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.4 attack_power points (0.68 DPS) | yes | Tiger Band (6749, -0.18 DPS) [quest]; Silverlaine's Family Seal (6321, -0.26 DPS) [dungeon]; Ironspine's Eye (7686, -0.30 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 12.7 attack_power points (0.53 DPS) | yes | Silverlaine's Family Seal (6321, -0.11 DPS) [dungeon]; Ironspine's Eye (7686, -0.15 DPS) [dungeon]; Tiger Band (6749, -1.01 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Defiler's Talisman (21120, -2.38 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 0.0 attack_power points (0.00 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Corpsemaker (6687, -0.14 DPS) [dungeon]; Viscous Hammer (13045, -15.45 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wildhunter Cloak; chest: Shining Silver Breastplate; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Girdle of Golem Strength; legs: Ferine Leggings; feet: Hard Gold Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; trinket1: Darkspear Voodoo Seal; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 381, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 40 (undead, 00000000000000000-0000000000000000-55223331211000210)

Set DPS (verified): 83.3. Weights run: 1.4s. Verify run: 1.4s. 535 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.228 ± 0.031, crit=0.306 ± 0.009 per rating point (14 rating = 1%, 4.287 per %), hit=0.312 ± 0.005 per rating point (10 rating = 1%, 3.119 per %), melee_haste=2.170 ± 0.198

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 30.3 attack_power points (1.29 DPS) | yes | Hard Gold Coif (250537, -0.10 DPS) [crafted]; Chromite Barbute (8142, -0.15 DPS) [world_drop]; Icemetal Barbute (10763, -0.65 DPS, sim-verified) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Ethereal Talisman (4430, -0.13 DPS) [quest]; Kaleidoscope Chain (13084, -0.22 DPS) [world_drop]; Zealous Shadowshard Pendant (17772, -1.00 DPS, sim-verified) [quest] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 attack_power points (0.93 DPS) | yes | Shining Mithril Pauldrons (250541, -0.08 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.17 DPS) [world_drop]; Chromite Pauldrons (8144, -0.87 DPS, sim-verified) [world_drop] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 10.3 attack_power points (0.44 DPS) | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Wolfmaster Cape (6314, -0.01 DPS) [dungeon]; Khan's Cloak (14781, -0.10 DPS) [world_drop] |
| chest | Jouster's Chestplate (8157) (or Avenger's Armor (1488)) | World drop [world_drop] | 30.0 attack_power points (1.27 DPS) | yes | Kolkar Marauder Chain (6773, -0.02 DPS) [quest]; Shining Mithril Breastplate (250540, -0.08 DPS) [crafted]; Avenger's Armor (1488, -2.14 DPS, sim-verified) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.85 DPS) | yes | Ravager's Armguards (14770, -0.22 DPS) [world_drop]; Darkspear Armsplints (4132, -0.25 DPS) [quest]; Pugilist Bracers (4438, -0.28 DPS, sim-verified) [dungeon] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 attack_power points (1.36 DPS) | yes | Truesilver Gauntlets (7938, +0.00 DPS, sim-verified) [crafted]; Scarlet Gauntlets (10331, -0.24 DPS) [dungeon]; Reticulated Bone Gauntlets (9435, -0.25 DPS) [world_drop] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (1.27 DPS) | yes | Defiler's Plate Girdle (20206, +0.00 DPS, sim-verified) [rep]; Boar Champion's Belt (10768, -0.00 DPS) [dungeon]; Tharg's Shoelace (9705, -0.17 DPS) [quest] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.78 DPS) | yes | Firemane Leggings (13129, -0.28 DPS, sim-verified) [world_drop]; Orcish War Leggings (7929, -0.34 DPS) [crafted]; Symbolic Legplates (14829, -0.45 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 27.6 attack_power points (1.17 DPS) | yes | Prowler's Leather Shoes (252465, +0.00 DPS, sim-verified) [crafted]; Skirmisher's Mail Boots (252564, -0.24 DPS) [crafted]; Obsidian Greaves (13068, -0.26 DPS) [world_drop] |
| finger1 | Assault Band (13095) (or Mark of Kern (2262)) | World drop [world_drop] | 20.0 attack_power points (0.85 DPS) | yes | Thunderbrow Ring (13097, -0.14 DPS) [world_drop]; Suspicious Spare Part (274754, -0.25 DPS) [vendor]; Mark of Kern (2262, -1.69 DPS, sim-verified) [dungeon] |
| finger2 | Legionnaire's Band (19512) | Warsong Outriders [rep] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Suspicious Spare Part (274754, -0.16 DPS) [vendor]; Mark of Kern (2262, -1.51 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 0.0 attack_power points (0.00 DPS) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor]; Fiery War Axe (870, -9.55 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; shoulder: Hard Gold Pauldrons; back: Dark Hooded Cape; chest: Jouster's Chestplate; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Defiler's Leather Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Assault Band; finger2: Legionnaire's Band

No-known-source sample (15 of 535, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band

### Band 50 (undead, 00000000000000000-5500000000000000-55223331211000210)

Set DPS (verified): 94.1. Weights run: 1.5s. Verify run: 1.3s. 718 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.231 ± 0.026, crit=0.395 ± 0.012 per rating point (14 rating = 1%, 5.537 per %), hit=0.389 ± 0.006 per rating point (10 rating = 1%, 3.895 per %), melee_haste=2.628 ± 0.325

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 48.0 attack_power points (2.06 DPS) | yes | Fury Visor (20521, +0.00 DPS, sim-verified) [quest]; Ebon Mask (19984, -0.50 DPS) [quest]; Bloomsprout Headpiece (17767, -0.52 DPS) [dungeon] |
| neck | Woven Ivy Necklace (19159) | Wanted: Vile Priestess Hexx and Her Minions [quest] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Ghostshard Talisman (7731, -0.00 DPS) [dungeon]; Skibi's Pendant (13089, -0.05 DPS) [world_drop]; Zealous Shadowshard Pendant (17772, -1.03 DPS, sim-verified) [quest] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 27.8 attack_power points (1.20 DPS) | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS, sim-verified) [vendor]; Earthslag Shoulders (11632, -0.08 DPS) [dungeon]; Wyrmslayer Spaulders (13066, -0.09 DPS) [world_drop] |
| back | Bloodlust Cape (14801) | World drop [world_drop] | 18.0 attack_power points (0.77 DPS) | yes | Dark Hooded Cape (5257, -0.33 DPS) [world]; Battlehard Cape (11858, -0.34 DPS) [quest]; Blackveil Cape (11626, -0.97 DPS, sim-verified) [dungeon] |
| chest | Warforged Chestplate (11195) | Broken Alliances [quest] | 48.0 attack_power points (2.06 DPS) | yes | Valorous Chestguard (8274, -0.10 DPS, sim-verified) [world_drop]; Mixologist's Tunic (12793, -0.41 DPS) [dungeon]; Coldmetal Guard (274758, -0.43 DPS) [vendor] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 attack_power points (1.20 DPS) | yes | Officer's Wristguards (250581, -0.20 DPS) [crafted]; Giantslayer Bracers (13076, -0.29 DPS) [world_drop]; Runed Golem Shackles (12550, -2.42 DPS, sim-verified) [dungeon] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 46.3 attack_power points (1.99 DPS) | yes | Gauntlets of Divinity (7724, -0.61 DPS) [dungeon]; Truesilver Gauntlets (7938, -0.61 DPS) [crafted]; Officer's Gloves (250551, -0.70 DPS, sim-verified) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (1.98 DPS) | yes | Atal'alarion's Tusk Ring (10798, +0.00 DPS, sim-verified) [dungeon]; Belt of the Gladiator (13134, -0.43 DPS) [world_drop]; Prowler's Leather Waistguard (252473, -0.65 DPS) [crafted] |
| legs | Golem Shard Leggings (13074) | World drop [world_drop] | 44.0 attack_power points (1.89 DPS) | yes | Scarlet Leggings (10330, -0.11 DPS, sim-verified) [dungeon]; Silvershell Leggings (10633, -0.17 DPS) [dungeon]; Elemental Rockridge Leggings (17711, -0.17 DPS) [dungeon] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Officer's Sabatons (250561, -0.02 DPS) [crafted]; Officer's Boots (250546, -0.04 DPS) [crafted]; Battlechaser's Greaves (12555, -1.87 DPS, sim-verified) [dungeon] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.03 DPS) | yes | Legionnaire's Band (19511, -0.08 DPS) [rep]; Mark of Kern (2262, -0.17 DPS) [dungeon]; Assault Band (13095, -0.17 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 23.9 attack_power points (1.03 DPS) | yes | Mark of Kern (2262, -0.17 DPS) [dungeon]; Assault Band (13095, -0.17 DPS) [world_drop]; Legionnaire's Band (19511, -0.35 DPS, sim-verified) [rep] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+2.9 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 attack_power points (0.00 DPS) | yes | Smoking Heart of the Mountain (11811, -0.76 DPS, sim-verified) [crafted] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | 0.0 attack_power points (0.00 DPS) | yes | Darkspear Raider's Reaper (272080, -0.42 DPS) [vendor]; Taran Icebreaker (2915, -0.48 DPS) [world_drop]; Blight (7959, -0.99 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Embrace of the Lycan; neck: Woven Ivy Necklace; shoulder: Officer's Pauldrons; back: Bloodlust Cape; chest: Warforged Chestplate; wrist: Bracers of the Stone Princess; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Golem Shard Leggings; feet: Prowler's Leather Boots; finger1: White Bone Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Thorium Greatmace

No-known-source sample (15 of 718, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band

### Band 60 (undead, 00000000000000000-5532500000000000-55223331211000210)

Set DPS (verified): 178.2. Weights run: 1.4s. Verify run: 1.4s. 1602 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.339 ± 0.051, crit=0.526 ± 0.017 per rating point (14 rating = 1%, 7.362 per %), hit=0.466 ± 0.008 per rating point (10 rating = 1%, 4.660 per %), melee_haste=3.638 ± 0.629

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Inquisition Helmet (240027) | Leonid Barthalomew the Revered [vendor] | 116.7 attack_power points (5.10 DPS) | yes | Lionheart Helm (12640, -2.48 DPS) [crafted]; Embrace of the Lycan (9479, -3.00 DPS) [dungeon]; Inquisition Crown (240035, -5.35 DPS, sim-verified) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (178.2 DPS) | yes | Will of the Martyr (17044, -0.06 DPS) [quest]; Beads of Ogre Might (22150, -0.12 DPS) [quest]; Imperial Jewel (11933, -3.03 DPS, sim-verified) [dungeon] |
| shoulder | Inquisition Shoulderplates (240025) | Leonid Barthalomew the Revered [vendor] | 84.0 attack_power points (3.67 DPS) | yes | Defiler's Plate Spaulders (20212, -1.85 DPS) [rep]; Black Dragonscale Shoulders (15051, -1.92 DPS) [crafted]; Inquisition Pauldrons (240033, -5.78 DPS, sim-verified) [vendor] |
| back | Deathguard's Cloak (20068) | The Defilers [rep] | 35.7 attack_power points (1.56 DPS) | yes | Howler's Furs (272414, -0.13 DPS) [vendor]; Shadewood Cloak (18328, -0.42 DPS) [dungeon]; Shroud of Domination (22337, -2.63 DPS, sim-verified) [dungeon] |
| chest | Inquisition Breastplate (240030) | Leonid Barthalomew the Revered [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Obsidian Mail Tunic (22191, -1.46 DPS) [crafted]; Timbermaw Tunic (252484, -1.63 DPS) [crafted]; Breastplate of Undead Slaying (23087, -9.79 DPS, sim-verified) [world] |
| wrist | Inquisition Vambraces (240023) | Leonid Barthalomew the Revered [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Berserker Bracers (19578, -1.11 DPS) [rep]; Windtalker's Wristguards (19582, -1.23 DPS) [rep]; Bracers of Undead Slaying (23090, -7.08 DPS, sim-verified) [world] |
| hands | Inquisition Gloves (240028) | Leonid Barthalomew the Revered [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Raider Gauntlets (272095, -1.12 DPS) [vendor]; Chromatic Gauntlets (19157, -1.40 DPS) [crafted]; Razor Gauntlets (18326, -6.53 DPS, sim-verified) [dungeon] |
| waist | Inquisition Belt (240024) | Leonid Barthalomew the Revered [vendor] | 86.7 attack_power points (3.79 DPS) | yes | Radiant Girdle of the Dawn (227814, -1.11 DPS) [vendor]; Ferocity of the Timbermaw (227805, -1.27 DPS) [vendor]; Dense Timbermaw Belt (227807, -3.66 DPS, sim-verified) [vendor] |
| legs | Inquisition Leggings (240026) | Leonid Barthalomew the Revered [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Titanic Leggings (22385, -1.87 DPS) [crafted]; Sentinel's Plate Legguards (237825, -2.01 DPS) [vendor]; Cloudkeeper Legplates (14554, -12.08 DPS, sim-verified) [world_drop] |
| feet | Inquisition Greaves (240029) | Leonid Barthalomew the Revered [vendor] | 86.0 attack_power points (3.76 DPS) | yes | Pads of the Dread Wolf (13210, -2.01 DPS) [dungeon]; Clutchlord's Stompers (275627, -2.10 DPS) [crafted]; Scalegut Treaders (275618, -5.78 DPS, sim-verified) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Signet Ring of the Bronze Dragonflight (234030, -0.17 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234026, -0.26 DPS) [vendor]; Naglering (11669, -5.66 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 0.0 attack_power points (0.00 DPS) | yes | Band of the Ogre King (18522, -0.00 DPS) [dungeon]; Legionnaire's Band (19510, -0.01 DPS) [rep]; Naglering (11669, -5.13 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (+4.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | 0.0 attack_power points (0.00 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS, sim-verified) [quest] |
| main_hand | Sulfuron Hammer (17193) | Blacksmithing [crafted] | 0.0 attack_power points (0.00 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Nightfall (19169, -0.55 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Inquisition Helmet; neck: Medallion of the Dawn; shoulder: Inquisition Shoulderplates; back: Deathguard's Cloak; chest: Inquisition Breastplate; wrist: Inquisition Vambraces; hands: Inquisition Gloves; waist: Inquisition Belt; legs: Inquisition Leggings; feet: Inquisition Greaves; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Talisman of Ascendance; main_hand: Sulfuron Hammer

No-known-source sample (15 of 1602, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band

