# Leveling BiS: Enhancement

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-253100000000000000-0000000000000000)

Set DPS (verified): 42.8. Weights run: 1.1s. Verify run: 1.2s. 229 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=0.235 ± 0.025, crit=4.333 ± 0.103, hit=4.391 ± 0.054, melee_haste=2.440 ± 0.609

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.69 DPS) | yes | Defender's Leather Hood (252447, -0.17 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.61 DPS) [crafted]; Brawler's Leather Hood (252504, -0.63 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 1.4 attack_power points (0.05 DPS) | yes | Erudite's Amulet (277204, -0.05 DPS, sim-verified) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.21 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS, sim-verified) [crafted]; Serpent's Shoulders (5404, -0.17 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Catacomb Cloak (279899, -0.07 DPS) [quest]; Dark Leather Cloak (2316, -0.11 DPS) [crafted]; Grave Shroud (279865, -0.13 DPS, sim-verified) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.69 DPS) | yes | Veteran's Chain Shirt (250488, -0.16 DPS, sim-verified) [crafted]; Defender's Leather Armor (252434, -0.18 DPS) [crafted]; Totemic Leather Armor (252435, -0.21 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Cryptwalker Bracers (280095, -0.08 DPS, sim-verified) [quest]; Bravo's Armbands (270015, -0.17 DPS) [quest]; Bristlebark Bindings (14569, -0.18 DPS) [world_drop] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 60.7 attack_power points (2.09 DPS) | yes | Thorbia's Gauntlets (12994, +0.00 DPS, sim-verified) [world_drop]; Gold-flecked Gloves (5195, -1.61 DPS) [dungeon]; Polar Gauntlets (7606, -1.68 DPS) [quest] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (0.62 DPS) | yes | Cobrahn's Grasp (6460, -0.09 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.21 DPS) [world]; Support Girdle (1215, -0.28 DPS) [world] |
| legs | Chausses of Westfall (6087) | The Defias Brotherhood [quest] | 22.0 attack_power points (0.76 DPS) | yes | Veteran's Chain Leggings (250493, -0.04 DPS, sim-verified) [crafted]; Defender's Leather Pants (252445, -0.11 DPS) [crafted]; Totemic Leather Pants (252446, -0.14 DPS) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 11.2 attack_power points (0.39 DPS) | yes | Veteran's Boots (250503, -0.02 DPS, sim-verified) [crafted]; Guard's Boots (250504, -0.04 DPS) [crafted]; Defender's Leather Boots (252441, -0.04 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.9 attack_power points (0.31 DPS) | yes | Loop of Sacrifice (281673, -0.10 DPS) [quest]; The 1 Ring (8350, -0.23 DPS) [world]; Ring of the Moon (12052, -0.24 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Loop of Sacrifice (281673, +0.00 DPS, sim-verified) [quest]; The 1 Ring (8350, -0.20 DPS) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 237.0 attack_power points (8.18 DPS) | yes | Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon]; The Axe of Severing (23171, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest] |
| off_hand | Diamond Hammer (2194) | World drop [world_drop] | 232.8 attack_power points (8.03 DPS) | yes | Blackfang (2236, -1.95 DPS, sim-verified) [world_drop]; Redbeard Crest (12997, -7.62 DPS) [world_drop]; Bear Buckler (4821, -7.82 DPS) [vendor] |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: Butcher's Cleaver; off_hand: Diamond Hammer

No-known-source sample (15 of 229, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher

### Band 30 (dwarf, 0000000000000000-253130030004000000-0000000000000000)

Set DPS (verified): 71.5. Weights run: 1.2s. Verify run: 1.3s. 393 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.432 ± 0.100, crit=7.348 ± 0.337, hit=6.193 ± 0.087, melee_haste=2.835 ± 0.417

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 attack_power points (0.92 DPS) | yes | Defender's Leather Helm (252455, -0.07 DPS) [crafted]; Veteran's Chain Helm (250498, -0.12 DPS, sim-verified) [crafted]; Crusader's Chain Helm (250502, -0.14 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | sim-verified (69.5 DPS) | yes | River Pride Choker (13087, -0.06 DPS) [world_drop]; Sentinel's Medallion (19541, -0.22 DPS) [rep]; Ghostshard Talisman (7731, -1.08 DPS, sim-verified) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 14.7 attack_power points (0.52 DPS) | yes | Golden Scale Shoulders (3841, -0.03 DPS) [crafted]; Mail Combat Spaulders (6404, -0.03 DPS) [world_drop]; Barbaric Iron Shoulders (7913, -0.09 DPS, sim-verified) [crafted] |
| back | Sergeant Major's Cape (16315) | PvP rank 9 · Sergeant Major · Alliance [pvp] | sim-verified (69.7 DPS) | yes | Hawkeye's Cloak (14593, -0.02 DPS) [world_drop]; Lambent Scale Cloak (4706, -0.06 DPS) [world_drop]; Wolfmaster Cape (6314, -1.32 DPS, sim-verified) [dungeon] |
| chest | Shining Silver Breastplate (2870) | Blacksmithing [crafted] | sim-verified (70.0 DPS) | yes | Barbaric Iron Breastplate (7914, -0.14 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.19 DPS) [crafted]; Avenger's Armor (1488, -1.63 DPS, sim-verified) [dungeon] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.57 DPS) | yes | Yorgen Bracers (13012, -0.08 DPS, sim-verified) [world_drop]; Bands of Serra'kis (6902, -0.14 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.19 DPS) [world_drop] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 102.9 attack_power points (3.64 DPS) | yes | Gauntlets of Ogre Strength (3341, +0.00 DPS, sim-verified) [world]; The Frozen Clutch (23170, -2.93 DPS) [dungeon]; Bonefist Gauntlets (4465, -3.00 DPS) [world] |
| waist | Girdle of Golem Strength (9405) (or Highlander's Plate Girdle (20126)) | World drop [world_drop] | 24.0 attack_power points (0.85 DPS) | yes | Highlander's Plate Girdle (20126, +0.00 DPS, sim-verified) [rep]; Highlander's Chain Girdle (20090, -0.00 DPS) [rep]; Highlander's Leather Girdle (20117, -0.00 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (0.92 DPS) | yes | Golden Scale Leggings (3843, -0.14 DPS) [crafted]; Chausses of Westfall (6087, -0.14 DPS) [quest]; Veteran's Silvered Chain Leggings (250523, -0.81 DPS, sim-verified) [crafted] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 17.0 attack_power points (0.60 DPS) | yes | Hard Gold Boots (250534, +0.00 DPS, sim-verified) [crafted]; Brawler's Leather Boots (252439, -0.17 DPS) [crafted]; Alacritous Treads (277234, -0.17 DPS) [quest] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 17.3 attack_power points (0.61 DPS) | yes | Tiger Band (6749, -0.19 DPS) [quest]; Ironspine's Eye (7686, -0.19 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.26 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 14.6 attack_power points (0.52 DPS) | yes | Ironspine's Eye (7686, -0.10 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.16 DPS) [dungeon]; Tiger Band (6749, -0.97 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Talisman of Arathor (21119, -1.30 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 345.4 attack_power points (12.21 DPS) | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 334.9 attack_power points (11.83 DPS) | yes | Shoni's Disarming Tool (9608, -3.89 DPS) [quest]; Royal Diplomatic Scepter (9457, -4.78 DPS, sim-verified) [dungeon]; Shield of Thorsen (13079, -11.34 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 30:** head: Tusken Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Sergeant Major's Cape; chest: Shining Silver Breastplate; wrist: Pugilist Bracers; waist: Girdle of Golem Strength; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; trinket1: Darkspear Voodoo Seal; main_hand: Pronged Reaver; off_hand: Ironspine's Fist

No-known-source sample (15 of 393, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 40 (dwarf, 0000000000000000-253130030005102051-0000000000000000)

Set DPS (verified): 87.0. Weights run: 1.4s. Verify run: 1.4s. 543 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.386 ± 0.051, crit=7.710 ± 0.213, hit=5.206 ± 0.082, melee_haste=3.244 ± 0.122

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 133.9 attack_power points (4.96 DPS) | yes | Hard Gold Coif (250537, -0.92 DPS, sim-verified) [crafted]; White Bandit Mask (10008, -3.99 DPS) [crafted]; Tusken Helm (6686, -4.00 DPS) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | sim-verified (87.0 DPS) | yes | Kaleidoscope Chain (13084, -0.17 DPS) [world_drop]; River Pride Choker (13087, -0.22 DPS) [world_drop]; Zealous Shadowshard Pendant (17772, -0.91 DPS, sim-verified) [quest] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 attack_power points (0.81 DPS) | yes | Imperial Leather Spaulders (4737, +0.00 DPS, sim-verified) [world_drop]; Wrangling Spaulders (15698, -0.19 DPS) [quest]; Sunburn Spaulders (274751, -0.21 DPS) [vendor] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 14.3 attack_power points (0.53 DPS) | yes | Wolfmaster Cape (6314, -0.16 DPS) [dungeon]; Sergeant Major's Cape (16315, -0.18 DPS) [pvp]; Dark Hooded Cape (5257, -2.32 DPS, sim-verified) [world] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 30.7 attack_power points (1.14 DPS) | yes | Shining Silver Breastplate (2870, -0.10 DPS) [crafted]; Golden Scale Cuirass (3845, -0.10 DPS) [crafted]; Avenger's Armor (1488, -2.30 DPS, sim-verified) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.74 DPS) | yes | Ravager's Armguards (14770, -0.17 DPS) [world_drop]; Yorgen Bracers (13012, -0.25 DPS) [world_drop]; Pugilist Bracers (4438, -0.26 DPS, sim-verified) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 127.9 attack_power points (4.74 DPS) | yes | Fletcher's Gloves (7348, -0.74 DPS) [crafted]; Shadowskin Gloves (18238, -0.74 DPS) [crafted]; Dragonscale Gauntlets (8347, -0.77 DPS, sim-verified) [crafted] |
| waist | Highlander's Chain Girdle (20089) | The League of Arathor [rep] | 115.9 attack_power points (4.29 DPS) | yes | Boar Champion's Belt (10768, -1.83 DPS, sim-verified) [dungeon]; Highlander's Leather Girdle (20116, -3.18 DPS) [rep]; Girdle of Golem Strength (9405, -3.40 DPS) [world_drop] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.55 DPS) | yes | Firemane Leggings (13129, -0.26 DPS, sim-verified) [world_drop]; Orcish War Leggings (7929, -0.30 DPS) [crafted]; Legguards of the Vault (9396, -0.52 DPS) [dungeon] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 24.7 attack_power points (0.91 DPS) | yes | Blackforge Greaves (6423, -0.11 DPS) [world_drop]; Skirmisher's Mail Boots (252564, -0.15 DPS, sim-verified) [crafted]; Ironheel Boots (4653, -0.17 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.74 DPS) | yes | Protector's Band (19515, -0.03 DPS) [rep]; Thunderbrow Ring (13097, -0.11 DPS) [world_drop]; Protector's Band (19517, -0.21 DPS) [rep] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.74 DPS) | yes | Protector's Band (19515, -0.05 DPS, sim-verified) [rep]; Thunderbrow Ring (13097, -0.11 DPS) [world_drop]; Protector's Band (19517, -0.21 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (17.56 DPS) | yes | Bonebiter (6830, +0.00 DPS) [quest]; Illusionary Rod (7713, +0.00 DPS) [dungeon]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor] |
| off_hand | Ardent Custodian (868) | World drop [world_drop] | 460.0 attack_power points (17.03 DPS) | yes | Shoni's Disarming Tool (9608, -8.71 DPS) [quest]; Curve-bladed Ripper (2815, -9.27 DPS, sim-verified) [world_drop]; Savage Boar's Guard (10767, -16.22 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; neck: Ghostshard Talisman; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Chain Girdle; legs: Scarlet Leggings; feet: Prowler's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Ardent Custodian

No-known-source sample (15 of 543, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 50 (dwarf, 5500000000000000-253130030005102051-0000000000000000)

Set DPS (verified): 123.4. Weights run: 1.4s. Verify run: 1.5s. 716 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.512 ± 0.061, crit=9.934 ± 0.268, hit=8.019 ± 0.121, melee_haste=3.814 ± 0.090

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 165.1 attack_power points (5.97 DPS) | yes | Knight-Lieutenant's Mail Helmet (223075, -0.27 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -0.94 DPS) [dungeon]; Embrace of the Lycan (9479, -4.24 DPS) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.72 DPS) | yes | Skibi's Pendant (13089, +0.00 DPS, sim-verified) [world_drop]; Ghostshard Talisman (7731, -0.22 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.36 DPS) [world_drop] |
| shoulder | Knight-Lieutenant's Mail Epaulets (223073) | Captain Dirgehammer [vendor] | 157.1 attack_power points (5.68 DPS) | yes | Prowler's Leather Shoulder (252534, -1.15 DPS, sim-verified) [crafted]; Failed Flying Experiment (9647, -4.77 DPS) [quest]; Skulker's Leather Shoulder (252535, -4.83 DPS) [crafted] |
| back | Bloodlust Cape (14801) | World drop [world_drop] | sim-verified (120.4 DPS) | yes | Sergeant Major's Cape (16336, -0.11 DPS) [pvp]; Dark Hooded Cape (5257, -0.18 DPS) [world]; Blackveil Cape (11626, -1.25 DPS, sim-verified) [dungeon] |
| chest | Knight's Mail Armor (223078) | Captain Dirgehammer [vendor] | 163.1 attack_power points (5.90 DPS) | yes | Wildthorn Mail (12624, -2.70 DPS, sim-verified) [crafted]; Grizzled Pelt (22274, -4.37 DPS) [quest]; Mixologist's Tunic (12793, -4.39 DPS) [dungeon] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 attack_power points (1.01 DPS) | yes | Prowler's Leather Bracers (252539, -0.25 DPS, sim-verified) [crafted]; Branded Leather Bracers (19508, -0.29 DPS) [dungeon]; Skulker's Leather Bracers (252540, -0.34 DPS) [crafted] |
| hands | Fists of The Five Thunders (227022) | Mokvar [vendor] | sim-verified (120.5 DPS) | yes | Sergeant Major's Mail Gauntlets (223076, -0.07 DPS) [vendor]; Dragonscale Gauntlets (8347, -0.54 DPS) [crafted]; Gloves of Holy Might (867, -1.39 DPS, sim-verified) [world_drop] |
| waist | Highlander's Chain Girdle (20088) (or Highlander's Leather Girdle (20115)) | The League of Arathor [rep] | 159.1 attack_power points (5.76 DPS) | yes | Highlander's Leather Girdle (20115, +0.00 DPS, sim-verified) [rep]; Highlander's Chain Girdle (20089, -0.43 DPS) [rep]; Highlander's Cloth Girdle (20097, -0.72 DPS) [rep] |
| legs | Knight's Mail Legplates (223074) | Captain Dirgehammer [vendor] | sim-verified (120.7 DPS) | yes | Stormshroud Pants (15057, -1.60 DPS, sim-verified) [crafted]; Scarlet Leggings (10330, -4.38 DPS) [dungeon]; Firemane Leggings (13129, -4.53 DPS) [world_drop] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 80.2 attack_power points (2.90 DPS) | yes | Prowler's Leather Boots (252468, -1.07 DPS, sim-verified) [crafted]; Skulker's Leather Boots (252469, -1.86 DPS) [crafted]; Skirmisher's Mail Sabatons (252578, -1.96 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 100.2 attack_power points (3.63 DPS) | yes | Mark of Kern (2262, -2.90 DPS) [dungeon]; Assault Band (13095, -2.90 DPS) [world_drop]; Thunderbrow Ring (13097, -2.99 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 24.6 attack_power points (0.89 DPS) | yes | Mark of Kern (2262, -0.17 DPS) [dungeon]; Assault Band (13095, -0.17 DPS) [world_drop]; Protector's Band (19515, -0.44 DPS, sim-verified) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Thunderbrew's Boot Flask (744, -0.41 DPS, sim-verified) [quest] |
| main_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | 0.0 attack_power points (0.00 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Shadowblade (2163, +0.00 DPS, sim-verified) [world_drop]; Kindling Stave (11750, +0.00 DPS) [dungeon] |
| off_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | 553.3 attack_power points (20.02 DPS) | yes | Claw of Celebras (17738, -2.57 DPS) [dungeon]; Shoni's Disarming Tool (9608, -11.89 DPS) [quest]; Shadowblade (2163, -12.25 DPS, sim-verified) [world_drop] |
| ranged | - | - |  |  |  |

**New at 50:** neck: Zealous Shadowshard Pendant; shoulder: Knight-Lieutenant's Mail Epaulets; back: Bloodlust Cape; chest: Knight's Mail Armor; wrist: Bracers of the Stone Princess; hands: Fists of The Five Thunders; waist: Highlander's Chain Girdle; legs: Knight's Mail Legplates; feet: Greaves of Withering Despair; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Frozen Heart of the Mountain; trinket2: Ankh of Life; main_hand: Dawn's Edge; off_hand: Hammer of the Northern Wind

No-known-source sample (15 of 716, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 60 (dwarf, 5533220000000000-253130030005102051-0000000000000000)

Set DPS (verified): 235.3. Weights run: 1.4s. Verify run: 1.7s. 1544 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.519 ± 0.068, crit=10.313 ± 0.324, hit=9.771 ± 0.172, melee_haste=3.972 ± 0.142

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soulcrusher Headpiece (240096) | Leonid Barthalomew the Revered [vendor] | 433.2 attack_power points (15.62 DPS) | yes | Bloodvine Goggles (19999, -3.37 DPS) [crafted]; Mask of the Unforgiven (13404, -3.45 DPS, sim-verified) [dungeon]; Ragefury Eyepatch (11735, -4.77 DPS) [dungeon] |
| neck | Blazefury Medallion (17111) | Lord Kazzak [world] | 0.0 attack_power points (0.00 DPS) | yes | Mark of Fordring (15411, +0.00 DPS) [quest]; Beads of Ogre Might (22150, +0.00 DPS) [quest]; Medallion of the Dawn (22659, -1.07 DPS, sim-verified) [quest] |
| shoulder | Soulcrusher Mantle (240125) | Leonid Barthalomew the Revered [vendor] | 242.1 attack_power points (8.73 DPS) | yes | Soulcrusher Epaulets (240135, -0.79 DPS, sim-verified) [vendor]; Knight-Lieutenant's Mail Epaulets (223073, -2.87 DPS) [vendor]; Darkspear Pauldrons (272105, -3.02 DPS) [vendor] |
| back | Chromatic Cloak (18509) (or Fel Cape (279269)) | Leatherworking [crafted] | 144.4 attack_power points (5.21 DPS) | yes | Fel Cape (279269, +0.00 DPS, sim-verified) [crafted]; Howler's Furs (272414, -0.67 DPS) [vendor]; Earthweave Cloak (21187, -1.40 DPS) [quest] |
| chest | Soulcrusher Armor (240128) | Leonid Barthalomew the Revered [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Bloodsoul Breastplate (19690, -10.41 DPS) [crafted]; Stormshroud Armor (15056, -10.58 DPS) [crafted]; Tunic of Undead Slaying (23089, -23.06 DPS, sim-verified) [world] |
| wrist | Soulcrusher Vambraces (240137) | Leonid Barthalomew the Revered [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Soulcrusher Bracers (240108, -1.15 DPS) [vendor]; Primal Batskin Bracers (19687, -1.61 DPS) [crafted]; Wristwraps of Undead Slaying (23093, -8.37 DPS, sim-verified) [world] |
| hands | Soulcrusher Grips (240130) | Leonid Barthalomew the Revered [vendor] | 352.8 attack_power points (12.72 DPS) | yes | Soulcrusher Mitts (240122, -2.31 DPS) [vendor]; Stormshroud Gloves (21278, -3.99 DPS) [crafted]; Soulcrusher Handguards (240095, -5.34 DPS, sim-verified) [vendor] |
| waist | Soulcrusher Girdle (240099) | Leonid Barthalomew the Revered [vendor] | 288.8 attack_power points (10.41 DPS) | yes | Soulcrusher Waistguard (240107, -0.53 DPS, sim-verified) [vendor]; Soulcrusher Belt (240136, -2.61 DPS) [vendor]; Highlander's Chain Girdle (20043, -3.98 DPS) [rep] |
| legs | Soulcrusher Leggings (240134) | Leonid Barthalomew the Revered [vendor] | sim-verified (235.3 DPS) | yes | Sentinel's Leather Pants (237818, -3.10 DPS) [vendor]; Stormshroud Pants (15057, -3.61 DPS) [crafted]; Sentinel's Chain Leggings (237819, -4.73 DPS, sim-verified) [vendor] |
| feet | Soulcrusher Sabatons (240102) | Leonid Barthalomew the Revered [vendor] | 121.7 attack_power points (4.39 DPS) | yes | Bloodvine Boots (19684, -0.87 DPS) [crafted]; Greaves of Withering Despair (22240, -0.87 DPS) [dungeon]; Fine Dawn Treaders (227815, -7.08 DPS, sim-verified) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 0.0 attack_power points (0.00 DPS) | yes | Signet Ring of the Bronze Dragonflight (234034, -4.05 DPS) [vendor]; Band of the Penitent (13217, -4.10 DPS) [quest]; Wrath of Cenarius (21190, -7.59 DPS, sim-verified) [quest] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | 0.0 attack_power points (0.00 DPS) | yes | Signet Ring of the Bronze Dragonflight (234034, -3.91 DPS) [vendor]; Band of the Penitent (13217, -3.96 DPS) [quest]; Wrath of Cenarius (21190, -7.18 DPS, sim-verified) [quest] |
| trinket1 | Earthstrike (21180) | Champion's Battlegear [quest] | sim-verified (+6.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | 0.0 attack_power points (0.00 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -10.84 DPS, sim-verified) [crafted] |
| main_hand | Ebon Hand (19170) | Blacksmithing [crafted] | 0.0 attack_power points (0.00 DPS) | yes | Grand Marshal's Sunderer (234566, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor]; The Lobotomizer (19324, -25.43 DPS, sim-verified) [rep] |
| off_hand | Persuader (22384) | Blacksmithing [crafted] | 0.0 attack_power points (0.00 DPS) | yes | Grand Marshal's Left Hand Blade (18847, +0.00 DPS) [vendor]; Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [vendor]; Eskhandar's Left Claw (18202, -23.00 DPS, sim-verified) [world] |
| ranged | Totem of the Storm (23199) (or Totem of Thunder (228176), Tidal Totem (272431), Totem of the Storm (272432), Burning Totem (272433), Totem of Urgency (279249), Totem of Ancestral Protection (249443), Kajaric Icon (206387), Polished Driftwood Icon (249398), Tempest Icon (206382), Dyadic Icon (206381), Galvanic Icon (206386), Sulfurous Icon (206388), Voltaic Icon (225838)) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Totem of Thunder (228176, +0.00 DPS, sim-verified) [vendor] |

**New at 60:** head: Soulcrusher Headpiece; neck: Blazefury Medallion; shoulder: Soulcrusher Mantle; back: Chromatic Cloak; chest: Soulcrusher Armor; wrist: Soulcrusher Vambraces; hands: Soulcrusher Grips; waist: Soulcrusher Girdle; legs: Soulcrusher Leggings; feet: Soulcrusher Sabatons; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Earthstrike; trinket2: Darkmoon Card: Maelstrom; main_hand: Ebon Hand; off_hand: Persuader; ranged: Totem of the Storm

No-known-source sample (15 of 1544, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

## Horde

### Band 20 (orc, 0000000000000000-253100000000000000-0000000000000000)

Set DPS (verified): 45.8. Weights run: 1.1s. Verify run: 1.2s. 224 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=0.235 ± 0.025, crit=4.333 ± 0.103, hit=4.391 ± 0.054, melee_haste=2.440 ± 0.609

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.69 DPS) | yes | Defender's Leather Hood (252447, -0.19 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.61 DPS) [crafted]; Brawler's Leather Hood (252504, -0.63 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 1.4 attack_power points (0.05 DPS) | yes | Erudite's Amulet (277204, -0.05 DPS, sim-verified) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.21 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS, sim-verified) [crafted]; Serpent's Shoulders (5404, -0.17 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Catacomb Cloak (279899, -0.07 DPS) [quest]; Subterranean Cape (14149, -0.07 DPS) [dungeon]; Grave Shroud (279865, -0.18 DPS, sim-verified) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.69 DPS) | yes | Veteran's Chain Shirt (250488, -0.17 DPS, sim-verified) [crafted]; Defender's Leather Armor (252434, -0.18 DPS) [crafted]; Totemic Leather Armor (252435, -0.21 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Cryptwalker Bracers (280095, -0.09 DPS, sim-verified) [quest]; Raptorcrest Bracers (270010, -0.14 DPS) [quest]; Bristlebark Bindings (14569, -0.18 DPS) [world_drop] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 60.7 attack_power points (2.09 DPS) | yes | Thorbia's Gauntlets (12994, +0.00 DPS, sim-verified) [world_drop]; Gold-flecked Gloves (5195, -1.61 DPS) [dungeon]; Blackened Defias Gloves (10401, -1.68 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (0.62 DPS) | yes | Cobrahn's Grasp (6460, -0.06 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.21 DPS) [world]; Support Girdle (1215, -0.28 DPS) [world] |
| legs | Veteran's Chain Leggings (250493) | Blacksmithing [crafted] | 19.2 attack_power points (0.66 DPS) | yes | Defender's Leather Pants (252445, -0.03 DPS, sim-verified) [crafted]; Totemic Leather Pants (252446, -0.04 DPS) [crafted]; Hulking Leggings (14748, -0.09 DPS) [world_drop] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 11.2 attack_power points (0.39 DPS) | yes | Veteran's Boots (250503, -0.03 DPS, sim-verified) [crafted]; Guard's Boots (250504, -0.04 DPS) [crafted]; Defender's Leather Boots (252441, -0.04 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.9 attack_power points (0.31 DPS) | yes | The 1 Ring (8350, -0.23 DPS) [world]; Ring of the Moon (12052, -0.24 DPS) [world_drop]; Demon Band (12054, -0.48 DPS, sim-verified) [world_drop] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | sim-verified (45.8 DPS) | yes | The 1 Ring (8350, -0.13 DPS) [world]; Ring of the Moon (12052, -0.14 DPS) [world_drop]; Demon Band (12054, -0.47 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 237.0 attack_power points (8.18 DPS) | yes | The Axe of Severing (23171, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest]; Hammerbone (270018, +0.00 DPS) [quest] |
| off_hand | Diamond Hammer (2194) | World drop [world_drop] | 232.8 attack_power points (8.03 DPS) | yes | Blackfang (2236, -0.23 DPS, sim-verified) [world_drop]; Redbeard Crest (12997, -7.62 DPS) [world_drop]; Bear Buckler (4821, -7.82 DPS) [vendor] |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Loop of Sacrifice; main_hand: Butcher's Cleaver; off_hand: Diamond Hammer

No-known-source sample (15 of 224, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5968 Rugged Boots; 6189 Durable Chain Shoulders

### Band 30 (orc, 0000000000000000-253130030004000000-0000000000000000)

Set DPS (verified): 74.9. Weights run: 1.2s. Verify run: 1.2s. 388 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.432 ± 0.100, crit=7.348 ± 0.337, hit=6.193 ± 0.087, melee_haste=2.835 ± 0.417

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 attack_power points (0.92 DPS) | yes | Defender's Leather Helm (252455, -0.07 DPS) [crafted]; Veteran's Chain Helm (250498, -0.13 DPS, sim-verified) [crafted]; Crusader's Chain Helm (250502, -0.14 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.49 DPS) | yes | Kaleidoscope Chain (13084, -0.10 DPS, sim-verified) [world_drop]; River Pride Choker (13087, -0.21 DPS) [world_drop]; Scout's Medallion (19537, -0.37 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 14.7 attack_power points (0.52 DPS) | yes | Golden Scale Shoulders (3841, -0.03 DPS) [crafted]; Mail Combat Spaulders (6404, -0.03 DPS) [world_drop]; Barbaric Iron Shoulders (7913, -0.13 DPS, sim-verified) [crafted] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 attack_power points (0.35 DPS) | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Hawkeye's Cloak (14593, -0.03 DPS) [world_drop]; Lambent Scale Cloak (4706, -0.07 DPS) [world_drop] |
| chest | Shining Silver Breastplate (2870) | Blacksmithing [crafted] | sim-verified (74.9 DPS) | yes | Barbaric Iron Breastplate (7914, -0.14 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.19 DPS) [crafted]; Avenger's Armor (1488, -1.82 DPS, sim-verified) [dungeon] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.57 DPS) | yes | Bands of Serra'kis (6902, -0.14 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.19 DPS) [world_drop]; Yorgen Bracers (13012, -0.46 DPS, sim-verified) [world_drop] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 102.9 attack_power points (3.64 DPS) | yes | Gauntlets of Ogre Strength (3341, +0.00 DPS, sim-verified) [world]; Warsong Gauntlets (16978, -2.93 DPS) [quest]; The Frozen Clutch (23170, -2.93 DPS) [dungeon] |
| waist | Girdle of Golem Strength (9405) (or Defiler's Plate Girdle (20207)) | World drop [world_drop] | 24.0 attack_power points (0.85 DPS) | yes | Defiler's Plate Girdle (20207, +0.00 DPS, sim-verified) [rep]; Defiler's Chain Girdle (20152, -0.00 DPS) [rep]; Defiler's Leather Girdle (20191, -0.00 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (0.92 DPS) | yes | Golden Scale Leggings (3843, -0.14 DPS) [crafted]; Slayer's Pants (14757, -0.14 DPS) [world_drop]; Veteran's Silvered Chain Leggings (250523, -1.37 DPS, sim-verified) [crafted] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 17.0 attack_power points (0.60 DPS) | yes | Hard Gold Boots (250534, +0.00 DPS, sim-verified) [crafted]; Brawler's Leather Boots (252439, -0.17 DPS) [crafted]; Alacritous Treads (277234, -0.17 DPS) [quest] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 17.3 attack_power points (0.61 DPS) | yes | Tiger Band (6749, -0.19 DPS) [quest]; Ironspine's Eye (7686, -0.19 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.26 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 14.6 attack_power points (0.52 DPS) | yes | Ironspine's Eye (7686, -0.10 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.16 DPS) [dungeon]; Tiger Band (6749, -1.00 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Defiler's Talisman (21120, -1.50 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 345.4 attack_power points (12.21 DPS) | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 334.9 attack_power points (11.83 DPS) | yes | Royal Diplomatic Scepter (9457, -2.98 DPS, sim-verified) [dungeon]; Shield of Thorsen (13079, -11.34 DPS) [world_drop]; Slayer's Shield (15892, -11.36 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Shining Silver Breastplate; wrist: Pugilist Bracers; waist: Girdle of Golem Strength; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; trinket1: Darkspear Voodoo Seal; main_hand: Pronged Reaver; off_hand: Ironspine's Fist

No-known-source sample (15 of 388, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band

### Band 40 (orc, 0000000000000000-253130030005102051-0000000000000000)

Set DPS (verified): 82.6. Weights run: 1.4s. Verify run: 1.4s. 537 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.386 ± 0.051, crit=7.710 ± 0.213, hit=5.206 ± 0.082, melee_haste=3.244 ± 0.122

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 133.9 attack_power points (4.96 DPS) | yes | Hard Gold Coif (250537, -0.74 DPS, sim-verified) [crafted]; White Bandit Mask (10008, -3.99 DPS) [crafted]; Tusken Helm (6686, -4.00 DPS) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | sim-verified (79.6 DPS) | yes | Ethereal Talisman (4430, -0.09 DPS) [quest]; Kaleidoscope Chain (13084, -0.17 DPS) [world_drop]; Zealous Shadowshard Pendant (17772, -0.82 DPS, sim-verified) [quest] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 attack_power points (0.81 DPS) | yes | Wrangling Spaulders (15698, -0.19 DPS) [quest]; Sunburn Spaulders (274751, -0.21 DPS) [vendor]; Imperial Leather Spaulders (4737, -0.21 DPS, sim-verified) [world_drop] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 11.9 attack_power points (0.44 DPS) | yes | Wolfmaster Cape (6314, +0.00 DPS, sim-verified) [dungeon]; Wildhunter Cloak (16658, -0.07 DPS) [quest]; Hawkeye's Cloak (14593, -0.12 DPS) [world_drop] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 30.7 attack_power points (1.14 DPS) | yes | Shining Silver Breastplate (2870, -0.10 DPS) [crafted]; Golden Scale Cuirass (3845, -0.10 DPS) [crafted]; Avenger's Armor (1488, -2.12 DPS, sim-verified) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.74 DPS) | yes | Ravager's Armguards (14770, -0.17 DPS) [world_drop]; Pugilist Bracers (4438, -0.20 DPS, sim-verified) [dungeon]; Darkspear Armsplints (4132, -0.22 DPS) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 127.9 attack_power points (4.74 DPS) | yes | Fletcher's Gloves (7348, -0.74 DPS) [crafted]; Shadowskin Gloves (18238, -0.74 DPS) [crafted]; Dragonscale Gauntlets (8347, -0.77 DPS, sim-verified) [crafted] |
| waist | Defiler's Chain Girdle (20153) | The Defilers [rep] | 115.9 attack_power points (4.29 DPS) | yes | Boar Champion's Belt (10768, -1.56 DPS, sim-verified) [dungeon]; Defiler's Leather Girdle (20192, -3.18 DPS) [rep]; Tharg's Shoelace (9705, -3.33 DPS) [quest] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.55 DPS) | yes | Firemane Leggings (13129, -0.24 DPS, sim-verified) [world_drop]; Orcish War Leggings (7929, -0.30 DPS) [crafted]; Legguards of the Vault (9396, -0.52 DPS) [dungeon] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 24.7 attack_power points (0.91 DPS) | yes | Blackforge Greaves (6423, -0.11 DPS) [world_drop]; Skirmisher's Mail Boots (252564, -0.12 DPS, sim-verified) [crafted]; Skulker's Leather Shoes (252531, -0.24 DPS) [crafted] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.74 DPS) | yes | Legionnaire's Band (19512, -0.03 DPS) [rep]; Thunderbrow Ring (13097, -0.11 DPS) [world_drop]; Legionnaire's Band (19513, -0.21 DPS) [rep] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.74 DPS) | yes | Legionnaire's Band (19512, +0.00 DPS, sim-verified) [rep]; Thunderbrow Ring (13097, -0.11 DPS) [world_drop]; Legionnaire's Band (19513, -0.21 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (17.56 DPS) | yes | Ardent Custodian (868, +0.00 DPS, sim-verified) [world_drop]; Illusionary Rod (7713, +0.00 DPS) [dungeon]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor] |
| off_hand | Curve-bladed Ripper (2815) | World drop [world_drop] | sim-verified (81.7 DPS) | yes | Ardent Custodian (868, -2.98 DPS, sim-verified) [world_drop]; Savage Boar's Guard (10767, -15.46 DPS) [dungeon]; Skullance Shield (13081, -15.63 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; shoulder: Hard Gold Pauldrons; back: Dark Hooded Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Chain Girdle; legs: Scarlet Leggings; feet: Prowler's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Curve-bladed Ripper

No-known-source sample (15 of 537, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (orc, 5500000000000000-253130030005102051-0000000000000000)

Set DPS (verified): 127.0. Weights run: 1.4s. Verify run: 1.4s. 697 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.512 ± 0.061, crit=9.934 ± 0.268, hit=8.019 ± 0.121, melee_haste=3.814 ± 0.090

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 165.1 attack_power points (5.97 DPS) | yes | Blood Guard's Mail Helmet (220820, -0.46 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -0.94 DPS) [dungeon]; Blood Guard's Inscribed Skullcap (220842, -0.94 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-verified (115.8 DPS) | yes | Woven Ivy Necklace (19159, -0.00 DPS) [quest]; Ghostshard Talisman (7731, -0.10 DPS) [dungeon]; Zealous Shadowshard Pendant (17772, -2.13 DPS, sim-verified) [quest] |
| shoulder | Blood Guard's Mail Epaulets (220823) | Lady Palanseer [vendor] | 157.1 attack_power points (5.68 DPS) | yes | Blood Guard's Pulsing Shoulders (220849, -0.65 DPS) [vendor]; Blood Guard's Inscribed Shoulder Pads (220841, -1.18 DPS, sim-verified) [vendor]; Prowler's Leather Shoulder (252534, -4.72 DPS) [crafted] |
| back | Bloodlust Cape (14801) | World drop [world_drop] | sim-verified (114.8 DPS) | yes | Dark Hooded Cape (5257, -0.18 DPS) [world]; Pridelord Cape (14673, -0.27 DPS) [world_drop]; Blackveil Cape (11626, -1.18 DPS, sim-verified) [dungeon] |
| chest | Stone Guard's Mail Armor (220826) | Lady Palanseer [vendor] | 163.1 attack_power points (5.90 DPS) | yes | Stone Guard's Inscribed Chestpiece (220838, -1.57 DPS, sim-verified) [vendor]; Wildthorn Mail (12624, -3.00 DPS) [crafted]; Stone Guard's Pulsing Breastplate (220844, -3.00 DPS) [vendor] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 attack_power points (1.01 DPS) | yes | Branded Leather Bracers (19508, -0.29 DPS) [dungeon]; Skulker's Leather Bracers (252540, -0.34 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.42 DPS, sim-verified) [crafted] |
| hands | Fists of The Five Thunders (227022) | Mokvar [vendor] | sim-verified (114.9 DPS) | yes | First Sergeant's Mail Gauntlets (220831, -0.07 DPS) [vendor]; Dragonscale Gauntlets (8347, -0.54 DPS) [crafted]; Gloves of Holy Might (867, -1.25 DPS, sim-verified) [world_drop] |
| waist | Defiler's Chain Girdle (20151) (or Defiler's Leather Girdle (20193)) | The Defilers [rep] | 159.1 attack_power points (5.76 DPS) | yes | Defiler's Leather Girdle (20193, +0.00 DPS, sim-verified) [rep]; Defiler's Chain Girdle (20153, -0.43 DPS) [rep]; Highlander's Mail Girdle (20118, -0.72 DPS) [vendor] |
| legs | Stone Guard's Mail Legplates (220834) | Lady Palanseer [vendor] | sim-verified (115.6 DPS) | yes | Stone Guard's Inscribed Legplates (220839, -0.87 DPS) [vendor]; Stormshroud Pants (15057, -1.93 DPS, sim-verified) [crafted]; Stone Guard's Pulsing Legplates (220847, -3.00 DPS) [vendor] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 80.2 attack_power points (2.90 DPS) | yes | Prowler's Leather Boots (252468, -0.72 DPS, sim-verified) [crafted]; Skulker's Leather Boots (252469, -1.86 DPS) [crafted]; Skirmisher's Mail Sabatons (252578, -1.96 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 100.2 attack_power points (3.63 DPS) | yes | White Bone Band (11862, -2.76 DPS) [quest]; Mark of Kern (2262, -2.90 DPS) [dungeon]; Assault Band (13095, -2.90 DPS) [world_drop] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 24.6 attack_power points (0.89 DPS) | yes | Legionnaire's Band (19512, -0.16 DPS) [rep]; Mark of Kern (2262, -0.17 DPS) [dungeon]; White Bone Band (11862, -0.62 DPS, sim-verified) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 attack_power points (0.00 DPS) | yes | Ankh of Life (1713, -0.43 DPS, sim-verified) [world_drop] |
| main_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Executioner's Cleaver (13018, +0.00 DPS) [world_drop] |
| off_hand | Shadowblade (2163) | World drop [world_drop] | sim-verified (119.3 DPS) | yes | Claw of Celebras (17738, -2.27 DPS) [dungeon]; White Bone Shredder (11863, -3.56 DPS) [quest]; Hammer of the Northern Wind (810, -5.68 DPS, sim-verified) [world_drop] |
| ranged | - | - |  |  |  |

**New at 50:** neck: Skibi's Pendant; shoulder: Blood Guard's Mail Epaulets; back: Bloodlust Cape; chest: Stone Guard's Mail Armor; wrist: Bracers of the Stone Princess; hands: Fists of The Five Thunders; waist: Defiler's Chain Girdle; legs: Stone Guard's Mail Legplates; feet: Greaves of Withering Despair; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Dawn's Edge; off_hand: Shadowblade

No-known-source sample (15 of 697, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (orc, 5533220000000000-253130030005102051-0000000000000000)

Set DPS (verified): 217.8. Weights run: 1.4s. Verify run: 1.4s. 1477 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.519 ± 0.068, crit=10.313 ± 0.324, hit=9.771 ± 0.172, melee_haste=3.972 ± 0.142

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soulcrusher Headpiece (240096) | Leonid Barthalomew the Revered [vendor] | 433.2 attack_power points (15.62 DPS) | yes | Bloodvine Goggles (19999, -3.37 DPS) [crafted]; Ragefury Eyepatch (11735, -4.77 DPS) [dungeon]; Mask of the Unforgiven (13404, -5.21 DPS, sim-verified) [dungeon] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (217.8 DPS) | yes | Beads of Ogre Might (22150, -1.68 DPS) [quest]; Blazefury Medallion (17111, -2.62 DPS, sim-verified) [world]; Mark of Fordring (15411, -4.76 DPS) [quest] |
| shoulder | Warlord's Mail Pauldrons (231654) | Lady Palanseer [vendor] | 276.1 attack_power points (9.95 DPS) | yes | Soulcrusher Mantle (240125, +0.00 DPS, sim-verified) [vendor]; Champion's Mail Pauldrons (227154, -0.29 DPS) [vendor]; Soulcrusher Epaulets (240135, -4.05 DPS) [vendor] |
| back | Chromatic Cloak (18509) (or Fel Cape (279269)) | Leatherworking [crafted] | 144.4 attack_power points (5.21 DPS) | yes | Fel Cape (279269, +0.00 DPS, sim-verified) [crafted]; Howler's Furs (272414, -0.67 DPS) [vendor]; Earthweave Cloak (21187, -1.40 DPS) [quest] |
| chest | Soulcrusher Armor (240128) | Leonid Barthalomew the Revered [vendor] | sim-verified (217.8 DPS) | yes | Bloodsoul Breastplate (19690, -10.41 DPS) [crafted]; Stormshroud Armor (15056, -10.58 DPS) [crafted]; Tunic of Undead Slaying (23089, -24.46 DPS, sim-verified) [world] |
| wrist | Soulcrusher Vambraces (240137) | Leonid Barthalomew the Revered [vendor] | sim-verified (217.8 DPS) | yes | Soulcrusher Bracers (240108, -1.15 DPS) [vendor]; Primal Batskin Bracers (19687, -1.61 DPS) [crafted]; Wristwraps of Undead Slaying (23093, -10.45 DPS, sim-verified) [world] |
| hands | Soulcrusher Grips (240130) | Leonid Barthalomew the Revered [vendor] | 352.8 attack_power points (12.72 DPS) | yes | Soulcrusher Mitts (240122, -2.31 DPS) [vendor]; General's Mail Vices (231655, -2.84 DPS) [vendor]; Soulcrusher Handguards (240095, -3.92 DPS, sim-verified) [vendor] |
| waist | Soulcrusher Girdle (240099) | Leonid Barthalomew the Revered [vendor] | 288.8 attack_power points (10.41 DPS) | yes | Soulcrusher Waistguard (240107, -0.29 DPS, sim-verified) [vendor]; Soulcrusher Belt (240136, -2.61 DPS) [vendor]; Defiler's Chain Girdle (20150, -3.98 DPS) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 404.7 attack_power points (14.59 DPS) | yes | Soulcrusher Leggings (240134, +0.00 DPS, sim-verified) [vendor]; Sentinel's Leather Pants (237818, -3.67 DPS) [vendor]; General's Mail Legguards (231658, -3.84 DPS) [vendor] |
| feet | General's Mail Greaves (231656) | Lady Palanseer [vendor] | 131.7 attack_power points (4.75 DPS) | yes | Soulcrusher Sabatons (240102, +0.00 DPS, sim-verified) [vendor]; Blood Guard's Mail Greaves (227158, -0.22 DPS) [vendor]; Fine Dawn Treaders (227815, -1.02 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (217.8 DPS) | yes | Signet Ring of the Bronze Dragonflight (234034, -4.05 DPS) [vendor]; Band of the Penitent (13217, -4.10 DPS) [quest]; Wrath of Cenarius (21190, -7.23 DPS, sim-verified) [quest] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (217.8 DPS) | yes | Signet Ring of the Bronze Dragonflight (234034, -3.91 DPS) [vendor]; Band of the Penitent (13217, -3.96 DPS) [quest]; Wrath of Cenarius (21190, -6.91 DPS, sim-verified) [quest] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (217.8 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Earthstrike (21180) | Champion's Battlegear [quest] | sim-verified (217.8 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Rune of the Guard Captain (19120, -1.23 DPS, sim-verified) [quest] |
| main_hand | Ebon Hand (19170) | Blacksmithing [crafted] | sim-verified (217.8 DPS) | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [vendor]; High Warlord's War Staff (234549, +0.00 DPS) [vendor]; The Lobotomizer (19324, -16.28 DPS, sim-verified) [rep] |
| off_hand | Persuader (22384) | Blacksmithing [crafted] | sim-verified (217.8 DPS) | yes | High Warlord's Left Claw (18848, +0.00 DPS) [vendor]; High Warlord's Left Claw (234558, +0.00 DPS) [vendor]; Eskhandar's Left Claw (18202, -17.25 DPS, sim-verified) [world] |
| ranged | Totem of the Storm (23199) (or Totem of Thunder (228176), Tidal Totem (272431), Totem of the Storm (272432), Burning Totem (272433), Totem of Urgency (279249), Totem of Ancestral Protection (249443), Kajaric Icon (206387), Polished Driftwood Icon (249398), Tempest Icon (206382), Dyadic Icon (206381), Galvanic Icon (206386), Sulfurous Icon (206388), Voltaic Icon (225838)) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Totem of Thunder (228176, +0.00 DPS, sim-verified) [vendor] |

**New at 60:** head: Soulcrusher Headpiece; neck: Medallion of the Dawn; shoulder: Warlord's Mail Pauldrons; back: Chromatic Cloak; chest: Soulcrusher Armor; wrist: Soulcrusher Vambraces; hands: Soulcrusher Grips; waist: Soulcrusher Girdle; legs: Sentinel's Chain Leggings; feet: General's Mail Greaves; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Darkmoon Card: Maelstrom; trinket2: Earthstrike; main_hand: Ebon Hand; off_hand: Persuader; ranged: Totem of the Storm

No-known-source sample (15 of 1477, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

