# Leveling BiS: Retribution

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-0000000000000000-55100000000000000)

Set DPS (verified): 34.6. Weights run: 2.4s. Verify run: 1.1s. 239 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.142 ± 0.013, crit=0.177 ± 0.004 per rating point (14 rating = 1%, 2.476 per %), hit=0.184 ± 0.002 per rating point (10 rating = 1%, 1.843 per %), melee_haste=-2.567 ± 0.137

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.69 DPS) | yes | Defender's Leather Hood (252447, -0.16 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.64 DPS) [crafted]; Brawler's Leather Hood (252504, -0.65 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 0.9 attack_power points (0.03 DPS) | yes | Erudite's Amulet (277204, -0.01 DPS) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.21 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.18 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Grave Shroud (279865, -0.06 DPS) [quest]; Catacomb Cloak (279899, -0.07 DPS) [quest]; Dark Leather Cloak (2316, -0.12 DPS) [crafted] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.69 DPS) | yes | Veteran's Chain Shirt (250488, -0.17 DPS, sim-verified) [crafted]; Defender's Leather Armor (252434, -0.19 DPS) [crafted]; Totemic Leather Armor (252435, -0.21 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Cryptwalker Bracers (280095, -0.07 DPS) [quest]; Bravo's Armbands (270015, -0.19 DPS) [quest]; Bristlebark Bindings (14569, -0.19 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.0 attack_power points (0.55 DPS) | yes | Gold-flecked Gloves (5195, -0.07 DPS) [dungeon]; Polar Gauntlets (7606, -0.14 DPS) [quest]; Blackened Defias Gloves (10401, -0.14 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.62 DPS) | yes | Cobrahn's Grasp (6460, -0.12 DPS) [dungeon]; Ruffian Belt (5975, -0.21 DPS) [world]; Hulking Belt (14746, -0.28 DPS) [world_drop] |
| legs | Chausses of Westfall (6087) | The Defias Brotherhood [quest] | 22.0 attack_power points (0.76 DPS) | yes | Veteran's Chain Leggings (250493, -0.11 DPS) [crafted]; Defender's Leather Pants (252445, -0.12 DPS) [crafted]; Totemic Leather Pants (252446, -0.14 DPS) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.7 attack_power points (0.37 DPS) | yes | Veteran's Boots (250503, -0.00 DPS) [crafted]; Guard's Boots (250504, -0.02 DPS) [crafted]; Defender's Leather Boots (252441, -0.02 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.6 attack_power points (0.30 DPS) | yes | The 1 Ring (8350, -0.22 DPS) [world]; Ring of the Moon (12052, -0.23 DPS) [world_drop]; Signet of the Zhevra (285330, -0.27 DPS) [world] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | The 1 Ring (8350, -0.20 DPS, sim-verified) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop]; Signet of the Zhevra (285330, -0.25 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Verigan's Fist (6953) | The Test of Righteousness [quest] | 358.7 attack_power points (12.43 DPS) | yes | Duskbringer (2205, -2.23 DPS) [dungeon]; Monstrous Cleaver (279864, -2.53 DPS) [quest]; Smite's Mighty Hammer (7230, -3.35 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: Verigan's Fist

No-known-source sample (15 of 239, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler

### Band 30 (human, 00000000000000000-0000000000000000-55223310000000000)

Set DPS (verified): 68.5. Weights run: 2.4s. Verify run: 1.4s. 404 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.123 ± 0.011, crit=0.229 ± 0.005 per rating point (14 rating = 1%, 3.210 per %), hit=0.224 ± 0.002 per rating point (10 rating = 1%, 2.238 per %), melee_haste=1.733 ± 0.079

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 attack_power points (1.08 DPS) | yes | Veteran's Chain Helm (250498, -0.08 DPS) [crafted]; Defender's Leather Helm (252455, -0.08 DPS) [crafted]; Crusader's Chain Helm (250502, -0.17 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.58 DPS) | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; River Pride Choker (13087, -0.25 DPS) [world_drop]; Sentinel's Medallion (19541, -0.54 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 14.0 attack_power points (0.58 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS) [world_drop]; Barbaric Iron Shoulders (7913, -0.05 DPS) [crafted]; Glimmering Mail Pauldrons (6388, -0.08 DPS) [world_drop] |
| back | Sergeant Major's Cape (16315) | PvP rank 9 · Sergeant Major · Alliance [pvp] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Lambent Scale Cloak (4706, -0.02 DPS) [world_drop]; Slayer's Cape (14752, -0.02 DPS) [world_drop]; Wolfmaster Cape (6314, -0.68 DPS, sim-verified) [dungeon] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.0 attack_power points (1.25 DPS) | yes | Shining Silver Breastplate (2870, -0.08 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.25 DPS) [crafted]; Hard Gold Cuirass (250533, -0.33 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.67 DPS) | yes | Yorgen Bracers (13012, -0.15 DPS) [world_drop]; Bands of Serra'kis (6902, -0.17 DPS) [dungeon]; Patterned Bronze Bracers (2868, -0.25 DPS) [crafted] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (0.91 DPS) | yes | The Frozen Clutch (23170, -0.08 DPS) [dungeon]; Bonefist Gauntlets (4465, -0.17 DPS) [world]; Mail Combat Gauntlets (4075, -0.23 DPS) [world_drop] |
| waist | Girdle of Golem Strength (9405) (or Highlander's Plate Girdle (20126)) | World drop [world_drop] | 24.0 attack_power points (1.00 DPS) | yes | Highlander's Plate Girdle (20126, +0.00 DPS) [rep]; Highlander's Chain Girdle (20090, -0.00 DPS) [rep]; Highlander's Leather Girdle (20117, -0.00 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.08 DPS) | yes | Golden Scale Leggings (3843, -0.17 DPS) [crafted]; Chausses of Westfall (6087, -0.17 DPS) [quest]; Veteran's Silvered Chain Leggings (250523, -1.65 DPS, sim-verified) [crafted] |
| feet | Hard Gold Boots (250534) | Blacksmithing [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Glimmering Mail Greaves (4073, -0.08 DPS) [world_drop]; Slayer's Slippers (14756, -0.08 DPS) [world_drop]; Trouncing Boots (4464, -0.81 DPS, sim-verified) [world] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.4 attack_power points (0.68 DPS) | yes | Tiger Band (6749, -0.18 DPS) [quest]; Silverlaine's Family Seal (6321, -0.26 DPS) [dungeon]; Ironspine's Eye (7686, -0.30 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 12.7 attack_power points (0.53 DPS) | yes | Tiger Band (6749, -0.03 DPS) [quest]; Silverlaine's Family Seal (6321, -0.11 DPS) [dungeon]; Ironspine's Eye (7686, -0.15 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Cobalt Crusher (7730, -1.41 DPS) [dungeon]; Viscous Hammer (13045, -15.80 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Sergeant Major's Cape; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Girdle of Golem Strength; legs: Ferine Leggings; feet: Hard Gold Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 404, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak

### Band 40 (human, 00000000000000000-0000000000000000-55223331211000210)

Set DPS (verified): 86.1. Weights run: 2.4s. Verify run: 1.3s. 562 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.229 ± 0.022, crit=0.310 ± 0.006 per rating point (14 rating = 1%, 4.341 per %), hit=0.308 ± 0.003 per rating point (10 rating = 1%, 3.084 per %), melee_haste=2.205 ± 0.140

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 30.3 attack_power points (1.29 DPS) | yes | Icemetal Barbute (10763, -0.10 DPS) [dungeon]; Hard Gold Coif (250537, -0.10 DPS) [crafted]; Chromite Barbute (8142, -0.15 DPS) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.85 DPS) | yes | Ghostshard Talisman (7731, -0.43 DPS, sim-verified) [dungeon]; Kaleidoscope Chain (13084, -0.47 DPS) [world_drop]; River Pride Choker (13087, -0.51 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 attack_power points (0.93 DPS) | yes | Shining Mithril Pauldrons (250541, -0.08 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.17 DPS) [dungeon]; Chromite Pauldrons (8144, -0.91 DPS, sim-verified) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 13.4 attack_power points (0.57 DPS) | yes | Wolfmaster Cape (6314, -0.14 DPS) [dungeon]; Slayer's Cape (14752, -0.23 DPS) [world_drop]; Dark Hooded Cape (5257, -0.58 DPS, sim-verified) [world] |
| chest | Jouster's Chestplate (8157) (or Avenger's Armor (1488)) | Uldaman: Ancient Treasure [dungeon] | 30.0 attack_power points (1.27 DPS) | yes | Avenger's Armor (1488, +0.00 DPS) [dungeon]; Kolkar Marauder Chain (6773, -0.02 DPS) [quest]; Shining Mithril Breastplate (250540, -0.08 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.85 DPS) | yes | Pugilist Bracers (4438, -0.17 DPS) [dungeon]; Ravager's Armguards (14770, -0.22 DPS) [world_drop]; Yorgen Bracers (13012, -0.31 DPS) [world_drop] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 attack_power points (1.36 DPS) | yes | Truesilver Gauntlets (7938, -0.00 DPS) [crafted]; Scarlet Gauntlets (10331, -0.24 DPS) [dungeon]; Reticulated Bone Gauntlets (9435, -0.25 DPS) [world_drop] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (1.27 DPS) | yes | Boar Champion's Belt (10768, -0.00 DPS) [dungeon]; Highlander's Plate Girdle (20125, -0.00 DPS) [rep]; Highlander's Chain Girdle (20090, -0.25 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.78 DPS) | yes | Firemane Leggings (13129, -0.17 DPS) [world_drop]; Orcish War Leggings (7929, -0.34 DPS) [crafted]; Symbolic Legplates (14829, -0.45 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 27.6 attack_power points (1.17 DPS) | yes | Prowler's Leather Shoes (252465, -0.17 DPS) [crafted]; Skirmisher's Mail Boots (252564, -0.24 DPS) [crafted]; Obsidian Greaves (13068, -0.26 DPS) [world_drop] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.85 DPS) | yes | Protector's Band (19515, -0.09 DPS) [rep]; Thunderbrow Ring (13097, -0.14 DPS) [world_drop]; Suspicious Spare Part (274754, -0.25 DPS) [vendor] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.85 DPS) | yes | Protector's Band (19515, -0.09 DPS) [rep]; Thunderbrow Ring (13097, -0.14 DPS) [world_drop]; Suspicious Spare Part (274754, -0.25 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (86.1 DPS) | yes | X'caliboar (10758, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Nightblade (1982, -2.65 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Jouster's Chestplate; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Highlander's Leather Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Mark of Kern; finger2: Assault Band

No-known-source sample (15 of 562, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

### Band 50 (human, 00000000000000000-5500000000000000-55223331211000210)

Set DPS (verified): 100.4. Weights run: 2.6s. Verify run: 1.4s. 722 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.228 ± 0.021, crit=0.404 ± 0.009 per rating point (14 rating = 1%, 5.652 per %), hit=0.386 ± 0.004 per rating point (10 rating = 1%, 3.864 per %), melee_haste=2.792 ± 0.231

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 48.0 attack_power points (2.06 DPS) | yes | Knight-Lieutenant's Plate Helm (220804, -0.54 DPS) [vendor]; Sunscale Helmet (14849, -0.63 DPS) [world_drop]; Bloomsprout Headpiece (17767, -1.00 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.86 DPS) | yes | Ghostshard Talisman (7731, -0.26 DPS) [dungeon]; Skibi's Pendant (13089, -0.30 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.48 DPS) [world_drop] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 27.8 attack_power points (1.19 DPS) | yes | Knight-Lieutenant's Plate Pauldrons (220795, -0.01 DPS) [vendor]; Earthslag Shoulders (11632, -0.08 DPS) [dungeon]; Wyrmslayer Spaulders (13066, -0.09 DPS) [world_drop] |
| back | Bloodlust Cape (14801) | World drop [world_drop] | 18.0 attack_power points (0.77 DPS) | yes | Blackveil Cape (11626, -0.12 DPS) [dungeon]; Sergeant Major's Cape (16336, -0.20 DPS) [pvp]; Dark Hooded Cape (5257, -0.33 DPS) [world] |
| chest | Warforged Chestplate (11195) | Tremors of the Earth [quest] | 48.0 attack_power points (2.06 DPS) | yes | Mixologist's Tunic (12793, -0.41 DPS) [dungeon]; Coldmetal Guard (274758, -0.43 DPS) [vendor]; Valorous Chestguard (8274, -0.63 DPS, sim-verified) [world_drop] |
| wrist | Arena Bands (18711) (or Bracers of the Stone Princess (17714)) | Arena Treasure Chest [world] | 28.0 attack_power points (1.20 DPS) | yes | Bracers of the Stone Princess (17714, +0.00 DPS, sim-verified) [dungeon]; Runed Golem Shackles (12550, -0.00 DPS) [dungeon]; Officer's Wristguards (250581, -0.20 DPS) [crafted] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 46.3 attack_power points (1.99 DPS) | yes | Gauntlets of Divinity (7724, -0.61 DPS) [dungeon]; Truesilver Gauntlets (7938, -0.61 DPS) [crafted]; Officer's Gloves (250551, -0.75 DPS, sim-verified) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (1.97 DPS) | yes | Belt of the Gladiator (13134, -0.43 DPS) [world_drop]; Atal'alarion's Tusk Ring (10798, -0.60 DPS, sim-verified) [dungeon]; Prowler's Leather Waistguard (252473, -0.66 DPS) [crafted] |
| legs | Golem Shard Leggings (13074) | World drop [world_drop] | 44.0 attack_power points (1.89 DPS) | yes | Scarlet Leggings (10330, -0.09 DPS) [dungeon]; Silvershell Leggings (10633, -0.17 DPS) [dungeon]; Elemental Rockridge Leggings (17711, -0.17 DPS) [dungeon] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 31.0 attack_power points (1.33 DPS) | yes | Prowler's Leather Boots (252468, -0.11 DPS) [crafted]; Officer's Sabatons (250561, -0.13 DPS) [crafted]; Officer's Boots (250546, -0.14 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 23.9 attack_power points (1.02 DPS) | yes | Mark of Kern (2262, -0.17 DPS) [dungeon]; Assault Band (13095, -0.17 DPS) [world_drop]; Thunderbrow Ring (13097, -0.31 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 22.1 attack_power points (0.95 DPS) | yes | Mark of Kern (2262, -0.09 DPS) [dungeon]; Assault Band (13095, -0.09 DPS) [world_drop]; Thunderbrow Ring (13097, -0.23 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (100.4 DPS) | yes | - |
| trinket2 | Sanctified Orb (20512) | Forging the Mightstone [quest] | sim-verified (100.4 DPS) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Nightblade (1982) | World drop [world_drop] | sim-verified (100.4 DPS) | yes | Taran Icebreaker (2915, +0.00 DPS) [world_drop]; Drakefang Butcher (12463, +0.00 DPS) [dungeon]; Thorium Greatmace (250613, -3.50 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Embrace of the Lycan; shoulder: Officer's Pauldrons; back: Bloodlust Cape; chest: Warforged Chestplate; wrist: Arena Bands; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Golem Shard Leggings; feet: Battlechaser's Greaves; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Frozen Heart of the Mountain; trinket2: Sanctified Orb; main_hand: Nightblade

No-known-source sample (15 of 722, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

### Band 60 (human, 00000000000000000-5532500000000000-55223331211000210)

Set DPS (verified): 163.3. Weights run: 2.6s. Verify run: 1.4s. 1674 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.336 ± 0.038, crit=0.511 ± 0.012 per rating point (14 rating = 1%, 7.150 per %), hit=0.460 ± 0.005 per rating point (10 rating = 1%, 4.603 per %), melee_haste=3.548 ± 0.431

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 59.5 attack_power points (2.60 DPS) | yes | Field Marshal's Lamellar Headguard (231648, -0.37 DPS) [vendor]; Field Marshal's Lamellar Faceguard (16474, -0.45 DPS) [vendor]; Embrace of the Lycan (9479, -2.60 DPS, sim-verified) [dungeon] |
| neck | Imperial Jewel (11933) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | 32.0 attack_power points (1.40 DPS) | yes | Medallion of the Dawn (22659, +0.00 DPS, sim-verified) [quest]; Will of the Martyr (17044, -0.09 DPS) [quest]; Beads of Ogre Might (22150, -0.15 DPS) [quest] |
| shoulder | Highlander's Lamellar Spaulders (20058) | The League of Arathor [rep] | sim-verified (163.3 DPS) | yes | Field Marshal's Lamellar Shoulders (231651, +0.00 DPS) [vendor]; Black Dragonscale Shoulders (15051, -0.07 DPS) [crafted]; Highlander's Plate Spaulders (20057, -2.31 DPS, sim-verified) [rep] |
| back | Cloak of the Honor Guard (20073) | The League of Arathor [rep] | 35.7 attack_power points (1.56 DPS) | yes | Shroud of Domination (22337, -0.07 DPS) [dungeon]; Howler's Furs (272414, -0.13 DPS) [vendor]; Shadewood Cloak (18328, -0.42 DPS) [dungeon] |
| chest | Obsidian Mail Tunic (22191) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Timbermaw Tunic (252484, -0.17 DPS) [crafted]; Cadaverous Armor (14637, -0.20 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -6.53 DPS, sim-verified) [world] |
| wrist | Berserker Bracers (19578) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windtalker's Wristguards (19582, -0.12 DPS) [rep]; Gordok Bracers of Power (18533, -0.29 DPS) [dungeon]; Bracers of Undead Slaying (23090, -3.11 DPS, sim-verified) [world] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Radiant Gloves of the Dawn (227817, -0.36 DPS) [vendor]; Timbermaw Brawlers (19049, -0.48 DPS) [crafted]; Razor Gauntlets (18326, -3.36 DPS, sim-verified) [dungeon] |
| waist | Dense Timbermaw Belt (227807) | Meilosh [vendor] | 64.0 attack_power points (2.80 DPS) | yes | Radiant Girdle of the Dawn (227814, +0.00 DPS, sim-verified) [vendor]; Ferocity of the Timbermaw (227805, -0.27 DPS) [vendor]; Heavy Obsidian Belt (22197, -0.61 DPS) [crafted] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentinel's Plate Legguards (237825, -0.15 DPS) [vendor]; Cloudkeeper Legplates (14554, -2.21 DPS, sim-verified) [world_drop] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 40.7 attack_power points (1.78 DPS) | yes | Pads of the Dread Wolf (13210, -0.03 DPS) [dungeon]; Clutchlord's Stompers (275627, -0.12 DPS) [crafted]; Goregasher Stompers (275624, -0.12 DPS) [crafted] |
| finger1 | Band of the Ogre King (18522) | Dire Maul: King Gordok [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Protector's Band (19514, -0.01 DPS) [rep]; Blackstone Ring (17713, -0.15 DPS) [dungeon]; Naglering (11669, -1.93 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Protector's Band (19514, -0.00 DPS) [rep]; Blackstone Ring (17713, -0.14 DPS) [dungeon]; Naglering (11669, -2.74 DPS, sim-verified) [dungeon] |
| trinket1 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Burst of Knowledge (11832, -0.75 DPS, sim-verified) [dungeon] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Darkmoon Card: Maelstrom (19289, +0.00 DPS) [quest] |
| main_hand | Blackblade of Shahram (12592) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Runeblade of Baron Rivendare (13505, -3.12 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Lionheart Helm; neck: Imperial Jewel; shoulder: Highlander's Lamellar Spaulders; back: Cloak of the Honor Guard; chest: Obsidian Mail Tunic; wrist: Berserker Bracers; hands: Raider Gauntlets; waist: Dense Timbermaw Belt; legs: Titanic Leggings; feet: Scalegut Treaders; finger1: Band of the Ogre King; finger2: Don Julio's Band; trinket1: Second Wind; trinket2: Draconic Infused Emblem; main_hand: Blackblade of Shahram

No-known-source sample (15 of 1674, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

## Horde

### Band 20 (undead, 00000000000000000-0000000000000000-55100000000000000)

Set DPS (verified): 32.2. Weights run: 2.4s. Verify run: 1.0s. 219 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.142 ± 0.013, crit=0.177 ± 0.004 per rating point (14 rating = 1%, 2.476 per %), hit=0.184 ± 0.002 per rating point (10 rating = 1%, 1.843 per %), melee_haste=-2.567 ± 0.137

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.69 DPS) | yes | Defender's Leather Hood (252447, -0.16 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.64 DPS) [crafted]; Brawler's Leather Hood (252504, -0.65 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 0.9 attack_power points (0.03 DPS) | yes | Erudite's Amulet (277204, -0.01 DPS) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.21 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.18 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Catacomb Cloak (279899, -0.07 DPS) [quest]; Subterranean Cape (14149, -0.07 DPS) [dungeon]; Grave Shroud (279865, -0.15 DPS, sim-verified) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.69 DPS) | yes | Veteran's Chain Shirt (250488, -0.17 DPS, sim-verified) [crafted]; Defender's Leather Armor (252434, -0.19 DPS) [crafted]; Totemic Leather Armor (252435, -0.21 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Raptorcrest Bracers (270010, -0.16 DPS, sim-verified) [quest]; Bristlebark Bindings (14569, -0.19 DPS) [world_drop]; Runed Copper Bracers (2854, -0.21 DPS) [crafted] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.0 attack_power points (0.55 DPS) | yes | Gold-flecked Gloves (5195, -0.07 DPS) [dungeon]; Blackened Defias Gloves (10401, -0.14 DPS) [dungeon]; Dagmire Gauntlets (6481, -0.19 DPS) [quest] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.62 DPS) | yes | Cobrahn's Grasp (6460, -0.12 DPS) [dungeon]; Ruffian Belt (5975, -0.21 DPS) [world]; Hulking Belt (14746, -0.28 DPS) [world_drop] |
| legs | Veteran's Chain Leggings (250493) | Blacksmithing [crafted] | 18.7 attack_power points (0.65 DPS) | yes | Defender's Leather Pants (252445, -0.00 DPS) [crafted]; Totemic Leather Pants (252446, -0.02 DPS) [crafted]; Hulking Leggings (14748, -0.08 DPS) [world_drop] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.7 attack_power points (0.37 DPS) | yes | Veteran's Boots (250503, -0.00 DPS) [crafted]; Guard's Boots (250504, -0.02 DPS) [crafted]; Defender's Leather Boots (252441, -0.02 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.6 attack_power points (0.30 DPS) | yes | Loop of Sacrifice (281673, -0.09 DPS) [quest]; The 1 Ring (8350, -0.22 DPS) [world]; Ring of the Moon (12052, -0.23 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Loop of Sacrifice (281673, -0.07 DPS) [quest]; The 1 Ring (8350, -0.20 DPS) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hammerbone (270018) | Leaders of the Fang [quest] | 306.2 attack_power points (10.61 DPS) | yes | Duskbringer (2205, -0.41 DPS) [dungeon]; Monstrous Cleaver (279864, -0.71 DPS) [quest]; Smite's Mighty Hammer (7230, -1.23 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: Hammerbone

No-known-source sample (15 of 219, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5968 Rugged Boots; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield; 9602 Brushwood Blade

### Band 30 (undead, 00000000000000000-0000000000000000-55223310000000000)

Set DPS (verified): 67.7. Weights run: 2.4s. Verify run: 1.4s. 381 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.123 ± 0.011, crit=0.229 ± 0.005 per rating point (14 rating = 1%, 3.210 per %), hit=0.224 ± 0.002 per rating point (10 rating = 1%, 2.238 per %), melee_haste=1.733 ± 0.079

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 attack_power points (1.08 DPS) | yes | Veteran's Chain Helm (250498, -0.08 DPS) [crafted]; Defender's Leather Helm (252455, -0.08 DPS) [crafted]; Crusader's Chain Helm (250502, -0.17 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.58 DPS) | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; River Pride Choker (13087, -0.25 DPS) [world_drop]; Scout's Medallion (19537, -0.54 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 14.0 attack_power points (0.58 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS) [world_drop]; Barbaric Iron Shoulders (7913, -0.05 DPS) [crafted]; Elite Shoulders (4835, -0.08 DPS) [vendor] |
| back | Wildhunter Cloak (16658) (or Wolfmaster Cape (6314)) | The Hunt Completed [quest] | 10.0 attack_power points (0.42 DPS) | yes | Wolfmaster Cape (6314, +0.00 DPS) [dungeon]; Lambent Scale Cloak (4706, -0.08 DPS) [world_drop]; Slayer's Cape (14752, -0.08 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.0 attack_power points (1.25 DPS) | yes | Shining Silver Breastplate (2870, -0.08 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.25 DPS) [crafted]; Hard Gold Cuirass (250533, -0.33 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.67 DPS) | yes | Yorgen Bracers (13012, -0.15 DPS) [world_drop]; Bands of Serra'kis (6902, -0.17 DPS) [dungeon]; Grimtoll Wristguards (15459, -0.24 DPS) [quest] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (0.91 DPS) | yes | Warsong Gauntlets (16978, -0.08 DPS) [quest]; The Frozen Clutch (23170, -0.08 DPS) [dungeon]; Bonefist Gauntlets (4465, -0.17 DPS) [world] |
| waist | Girdle of Golem Strength (9405) (or Defiler's Plate Girdle (20207)) | World drop [world_drop] | 24.0 attack_power points (1.00 DPS) | yes | Defiler's Plate Girdle (20207, +0.00 DPS) [rep]; Defiler's Chain Girdle (20152, -0.00 DPS) [rep]; Defiler's Leather Girdle (20191, -0.00 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.08 DPS) | yes | Golden Scale Leggings (3843, -0.17 DPS) [crafted]; Slayer's Pants (14757, -0.17 DPS) [world_drop]; Veteran's Silvered Chain Leggings (250523, -1.66 DPS, sim-verified) [crafted] |
| feet | Hard Gold Boots (250534) | Blacksmithing [crafted] | sim-verified (67.7 DPS) | yes | Glimmering Mail Greaves (4073, -0.08 DPS) [world_drop]; Slayer's Slippers (14756, -0.08 DPS) [world_drop]; Trouncing Boots (4464, -0.82 DPS, sim-verified) [world] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.4 attack_power points (0.68 DPS) | yes | Tiger Band (6749, -0.18 DPS) [quest]; Silverlaine's Family Seal (6321, -0.26 DPS) [dungeon]; Ironspine's Eye (7686, -0.30 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 12.7 attack_power points (0.53 DPS) | yes | Tiger Band (6749, -0.03 DPS) [quest]; Silverlaine's Family Seal (6321, -0.11 DPS) [dungeon]; Ironspine's Eye (7686, -0.15 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Cobalt Crusher (7730, -1.41 DPS) [dungeon]; Viscous Hammer (13045, -15.95 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wildhunter Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Girdle of Golem Strength; legs: Ferine Leggings; feet: Hard Gold Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 381, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 40 (undead, 00000000000000000-0000000000000000-55223331211000210)

Set DPS (verified): 85.5. Weights run: 2.4s. Verify run: 1.3s. 531 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.229 ± 0.022, crit=0.310 ± 0.006 per rating point (14 rating = 1%, 4.341 per %), hit=0.308 ± 0.003 per rating point (10 rating = 1%, 3.084 per %), melee_haste=2.205 ± 0.140

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 30.3 attack_power points (1.29 DPS) | yes | Icemetal Barbute (10763, -0.10 DPS) [dungeon]; Hard Gold Coif (250537, -0.10 DPS) [crafted]; Chromite Barbute (8142, -0.15 DPS) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.85 DPS) | yes | Ethereal Talisman (4430, -0.39 DPS) [quest]; Ghostshard Talisman (7731, -0.43 DPS, sim-verified) [dungeon]; Kaleidoscope Chain (13084, -0.47 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 attack_power points (0.93 DPS) | yes | Shining Mithril Pauldrons (250541, -0.08 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.17 DPS) [dungeon]; Chromite Pauldrons (8144, -0.89 DPS, sim-verified) [dungeon] |
| back | First Sergeant's Cloak (16340) | PvP rank 9 · First Sergeant · Horde [pvp] | 13.4 attack_power points (0.57 DPS) | yes | Wolfmaster Cape (6314, -0.14 DPS) [dungeon]; Wildhunter Cloak (16658, -0.14 DPS) [quest]; Dark Hooded Cape (5257, -0.80 DPS, sim-verified) [world] |
| chest | Jouster's Chestplate (8157) (or Avenger's Armor (1488)) | Uldaman: Ancient Treasure [dungeon] | 30.0 attack_power points (1.27 DPS) | yes | Avenger's Armor (1488, +0.00 DPS) [dungeon]; Kolkar Marauder Chain (6773, -0.02 DPS) [quest]; Shining Mithril Breastplate (250540, -0.08 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.85 DPS) | yes | Pugilist Bracers (4438, +0.00 DPS) [dungeon]; Berserker Bracers (19581, +0.00 DPS) [pvp]; Windtalker's Wristguards (19584, +0.00 DPS) [pvp] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 attack_power points (1.36 DPS) | yes | Truesilver Gauntlets (7938, -0.00 DPS) [crafted]; Scarlet Gauntlets (10331, -0.24 DPS) [dungeon]; Reticulated Bone Gauntlets (9435, -0.25 DPS) [world_drop] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (1.27 DPS) | yes | Boar Champion's Belt (10768, -0.00 DPS) [dungeon]; Defiler's Plate Girdle (20206, -0.00 DPS) [rep]; Tharg's Shoelace (9705, -0.17 DPS) [quest] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.78 DPS) | yes | Firemane Leggings (13129, -0.17 DPS) [world_drop]; Orcish War Leggings (7929, -0.34 DPS) [crafted]; Symbolic Legplates (14829, -0.45 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 27.6 attack_power points (1.17 DPS) | yes | Prowler's Leather Shoes (252465, +0.00 DPS, sim-verified) [crafted]; Skirmisher's Mail Boots (252564, -0.24 DPS) [crafted]; Obsidian Greaves (13068, -0.26 DPS) [world_drop] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.85 DPS) | yes | Legionnaire's Band (19512, -0.09 DPS) [rep]; Thunderbrow Ring (13097, -0.14 DPS) [world_drop]; Suspicious Spare Part (274754, -0.25 DPS) [vendor] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.85 DPS) | yes | Legionnaire's Band (19512, -0.09 DPS) [rep]; Thunderbrow Ring (13097, -0.14 DPS) [world_drop]; Suspicious Spare Part (274754, -0.25 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (85.5 DPS) | yes | X'caliboar (10758, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Nightblade (1982, -8.61 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: First Sergeant's Cloak; chest: Jouster's Chestplate; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Defiler's Leather Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Mark of Kern; finger2: Assault Band

No-known-source sample (15 of 531, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band

### Band 50 (undead, 00000000000000000-5500000000000000-55223331211000210)

Set DPS (verified): 100.3. Weights run: 2.6s. Verify run: 1.4s. 702 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.228 ± 0.021, crit=0.404 ± 0.009 per rating point (14 rating = 1%, 5.652 per %), hit=0.386 ± 0.004 per rating point (10 rating = 1%, 3.864 per %), melee_haste=2.792 ± 0.231

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 48.0 attack_power points (2.06 DPS) | yes | Blood Guard's Plate Helm (220803, -0.54 DPS) [vendor]; Bloomsprout Headpiece (17767, -0.61 DPS, sim-verified) [dungeon]; Sunscale Helmet (14849, -0.63 DPS) [world_drop] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.86 DPS) | yes | Woven Ivy Necklace (19159, -0.26 DPS) [quest]; Ghostshard Talisman (7731, -0.26 DPS) [dungeon]; Skibi's Pendant (13089, -0.30 DPS) [world_drop] |
| shoulder | Earthslag Shoulders (11632) | Blackrock Depths: Lord Roccor [dungeon] | sim-verified (100.3 DPS) | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.01 DPS) [world_drop]; Officer's Pauldrons (250576, -1.86 DPS, sim-verified) [crafted] |
| back | Bloodlust Cape (14801) | World drop [world_drop] | 18.0 attack_power points (0.77 DPS) | yes | Blackveil Cape (11626, +0.00 DPS, sim-verified) [dungeon]; First Sergeant's Cloak (16340, -0.20 DPS) [pvp]; Dark Hooded Cape (5257, -0.33 DPS) [world] |
| chest | Warforged Chestplate (11195) | Broken Alliances [quest] | 48.0 attack_power points (2.06 DPS) | yes | Valorous Chestguard (8274, -0.34 DPS) [world_drop]; Mixologist's Tunic (12793, -0.41 DPS) [dungeon]; Coldmetal Guard (274758, -0.43 DPS) [vendor] |
| wrist | Arena Bands (18711) (or Bracers of the Stone Princess (17714)) | Arena Treasure Chest [world] | 28.0 attack_power points (1.20 DPS) | yes | Bracers of the Stone Princess (17714, +0.00 DPS, sim-verified) [dungeon]; Berserker Bracers (19580, +0.00 DPS) [pvp]; Windtalker's Wristguards (19583, +0.00 DPS) [pvp] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 46.3 attack_power points (1.99 DPS) | yes | Gauntlets of Divinity (7724, -0.61 DPS) [dungeon]; Truesilver Gauntlets (7938, -0.61 DPS) [crafted]; Officer's Gloves (250551, -0.70 DPS, sim-verified) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (1.97 DPS) | yes | Belt of the Gladiator (13134, -0.43 DPS) [world_drop]; Atal'alarion's Tusk Ring (10798, -0.54 DPS, sim-verified) [dungeon]; Prowler's Leather Waistguard (252473, -0.66 DPS) [crafted] |
| legs | Golem Shard Leggings (13074) | World drop [world_drop] | 44.0 attack_power points (1.89 DPS) | yes | Scarlet Leggings (10330, -0.09 DPS) [dungeon]; Silvershell Leggings (10633, -0.17 DPS) [dungeon]; Elemental Rockridge Leggings (17711, -0.17 DPS) [dungeon] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 31.0 attack_power points (1.33 DPS) | yes | Prowler's Leather Boots (252468, -0.11 DPS) [crafted]; Officer's Sabatons (250561, -0.13 DPS) [crafted]; Officer's Boots (250546, -0.14 DPS) [crafted] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.03 DPS) | yes | Legionnaire's Band (19511, -0.08 DPS) [rep]; Mark of Kern (2262, -0.17 DPS) [dungeon]; Assault Band (13095, -0.17 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 23.9 attack_power points (1.02 DPS) | yes | Mark of Kern (2262, -0.17 DPS) [dungeon]; Assault Band (13095, -0.17 DPS) [world_drop]; Legionnaire's Band (19511, -0.74 DPS, sim-verified) [rep] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+3.2 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -0.84 DPS, sim-verified) [crafted] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Taran Icebreaker (2915, -0.47 DPS) [world_drop]; Drakefang Butcher (12463, -1.37 DPS) [dungeon]; Blight (7959, -1.42 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Embrace of the Lycan; shoulder: Earthslag Shoulders; back: Bloodlust Cape; chest: Warforged Chestplate; wrist: Arena Bands; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Golem Shard Leggings; feet: Battlechaser's Greaves; finger1: White Bone Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Thorium Greatmace

No-known-source sample (15 of 702, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band

### Band 60 (undead, 00000000000000000-5532500000000000-55223331211000210)

Set DPS (verified): 155.8. Weights run: 2.6s. Verify run: 1.3s. 1699 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.336 ± 0.038, crit=0.511 ± 0.012 per rating point (14 rating = 1%, 7.150 per %), hit=0.460 ± 0.005 per rating point (10 rating = 1%, 4.603 per %), melee_haste=3.548 ± 0.431

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 59.5 attack_power points (2.60 DPS) | yes | Warbear Helm (252485, -0.82 DPS) [crafted]; Crown of Tyranny (13359, -0.83 DPS) [dungeon]; Embrace of the Lycan (9479, -5.06 DPS, sim-verified) [dungeon] |
| neck | Imperial Jewel (11933) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | 32.0 attack_power points (1.40 DPS) | yes | Medallion of the Dawn (22659, +0.00 DPS, sim-verified) [quest]; Will of the Martyr (17044, -0.09 DPS) [quest]; Beads of Ogre Might (22150, -0.15 DPS) [quest] |
| shoulder | Defiler's Plate Spaulders (20212) | The Defilers [rep] | 41.7 attack_power points (1.82 DPS) | yes | Defiler's Leather Shoulders (20194, -0.25 DPS) [rep]; Defiler's Lizardhide Shoulders (20175, -0.34 DPS) [rep]; Black Dragonscale Shoulders (15051, -1.00 DPS, sim-verified) [crafted] |
| back | Deathguard's Cloak (20068) | The Defilers [rep] | 35.7 attack_power points (1.56 DPS) | yes | Shroud of Domination (22337, -0.07 DPS) [dungeon]; Howler's Furs (272414, -0.13 DPS) [vendor]; Shadewood Cloak (18328, -0.42 DPS) [dungeon] |
| chest | Obsidian Mail Tunic (22191) | Blacksmithing [crafted] | sim-verified (155.8 DPS) | yes | Timbermaw Tunic (252484, -0.17 DPS) [crafted]; Cadaverous Armor (14637, -0.20 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -6.29 DPS, sim-verified) [world] |
| wrist | Berserker Bracers (19578) | Warsong Outriders [rep] | sim-verified (155.8 DPS) | yes | Windtalker's Wristguards (19582, -0.12 DPS) [rep]; Gordok Bracers of Power (18533, -0.29 DPS) [dungeon]; Bracers of Undead Slaying (23090, -3.05 DPS, sim-verified) [world] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | sim-verified (155.8 DPS) | yes | Radiant Gloves of the Dawn (227817, -0.36 DPS) [vendor]; Timbermaw Brawlers (19049, -0.48 DPS) [crafted]; Razor Gauntlets (18326, -3.27 DPS, sim-verified) [dungeon] |
| waist | Dense Timbermaw Belt (227807) | Meilosh [vendor] | 64.0 attack_power points (2.80 DPS) | yes | Radiant Girdle of the Dawn (227814, +0.00 DPS, sim-verified) [vendor]; Ferocity of the Timbermaw (227805, -0.27 DPS) [vendor]; Heavy Obsidian Belt (22197, -0.61 DPS) [crafted] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | sim-verified (155.8 DPS) | yes | Sentinel's Plate Legguards (237825, -0.15 DPS) [vendor]; Outrider's Plate Legguards (22651, -0.38 DPS) [rep]; Cloudkeeper Legplates (14554, -5.37 DPS, sim-verified) [world_drop] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 40.7 attack_power points (1.78 DPS) | yes | Pads of the Dread Wolf (13210, -0.03 DPS) [dungeon]; Clutchlord's Stompers (275627, -0.12 DPS) [crafted]; Goregasher Stompers (275624, -0.12 DPS) [crafted] |
| finger1 | Band of the Ogre King (18522) | Dire Maul: King Gordok [dungeon] | sim-verified (155.8 DPS) | yes | Legionnaire's Band (19510, -0.01 DPS) [rep]; Blackstone Ring (17713, -0.15 DPS) [dungeon]; Naglering (11669, -1.81 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (155.8 DPS) | yes | Legionnaire's Band (19510, -0.00 DPS) [rep]; Blackstone Ring (17713, -0.14 DPS) [dungeon]; Naglering (11669, -3.86 DPS, sim-verified) [dungeon] |
| trinket1 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (155.8 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Burst of Knowledge (11832, -1.85 DPS, sim-verified) [dungeon] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (155.8 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Darkmoon Card: Maelstrom (19289, -0.85 DPS, sim-verified) [quest] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (155.8 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Blackblade of Shahram (12592, -4.09 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Lionheart Helm; neck: Imperial Jewel; shoulder: Defiler's Plate Spaulders; back: Deathguard's Cloak; chest: Obsidian Mail Tunic; wrist: Berserker Bracers; hands: Raider Gauntlets; waist: Dense Timbermaw Belt; legs: Titanic Leggings; feet: Scalegut Treaders; finger1: Band of the Ogre King; finger2: Don Julio's Band; trinket1: Second Wind; trinket2: Draconic Infused Emblem; main_hand: The Unstoppable Force

No-known-source sample (15 of 1699, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band

