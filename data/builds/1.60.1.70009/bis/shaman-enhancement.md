# Leveling BiS: Enhancement

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-253100000000000000-0000000000000000)

Set DPS (verified): 42.8. Weights run: 0.9s. Verify run: 1.0s. 215 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=0.235 ± 0.025, crit=4.333 ± 0.103, hit=4.241 ± 0.303, melee_haste=2.440 ± 0.609

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 | yes | Defender's Leather Hood (252447, -0.17 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.61 DPS) [crafted]; Brawler's Leather Hood (252504, -0.63 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 1.4 | yes | Scholarly Pendant (277203, -0.05 DPS) [quest]; Tarnished Locket (279870, -0.05 DPS) [quest]; Erudite's Amulet (277204, -0.05 DPS, sim-verified) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 | yes | Silvered Bronze Shoulders (3481, +0.00 DPS, sim-verified) [crafted]; Serpent's Shoulders (5404, -0.17 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.21 DPS) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 | yes | Catacomb Cloak (279899, -0.07 DPS) [quest]; Dark Leather Cloak (2316, -0.11 DPS) [crafted]; Grave Shroud (279865, -0.13 DPS, sim-verified) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 | yes | Veteran's Chain Shirt (250488, -0.16 DPS, sim-verified) [crafted]; Defender's Leather Armor (252434, -0.18 DPS) [crafted]; Totemic Leather Armor (252435, -0.21 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 | yes | Cryptwalker Bracers (280095, -0.08 DPS, sim-verified) [quest]; Bravo's Armbands (270015, -0.17 DPS) [quest]; Bristlebark Bindings (14569, -0.18 DPS) [world_drop] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 60.7 | yes | Thorbia's Gauntlets (12994, +0.00 DPS, sim-verified) [world_drop]; Gold-flecked Gloves (5195, -1.61 DPS) [dungeon]; Polar Gauntlets (7606, -1.68 DPS) [quest] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 | yes | Cobrahn's Grasp (6460, -0.09 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.21 DPS) [world]; Support Girdle (1215, -0.28 DPS) [world] |
| legs | Chausses of Westfall (6087) | The Defias Brotherhood [quest] | 22.0 | yes | Veteran's Chain Leggings (250493, -0.04 DPS, sim-verified) [crafted]; Defender's Leather Pants (252445, -0.11 DPS) [crafted]; Totemic Leather Pants (252446, -0.14 DPS) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 11.2 | yes | Veteran's Boots (250503, -0.02 DPS, sim-verified) [crafted]; Guard's Boots (250504, -0.04 DPS) [crafted]; Defender's Leather Boots (252441, -0.04 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.9 | yes | Loop of Sacrifice (281673, -0.10 DPS) [quest]; The 1 Ring (8350, -0.23 DPS) [world]; Ring of the Moon (12052, -0.24 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 | yes | Loop of Sacrifice (281673, +0.00 DPS, sim-verified) [quest]; The 1 Ring (8350, -0.20 DPS) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 237.0 | yes | Living Root (6631, +0.00 DPS) [dungeon]; Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest] |
| off_hand | Diamond Hammer (2194) | World drop [world_drop] | 232.8 | yes | Blackfang (2236, -1.95 DPS, sim-verified) [world_drop]; Redbeard Crest (12997, -7.62 DPS) [world_drop]; Bear Buckler (4821, -7.82 DPS) [vendor] |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; trinket1: Rune of Duty; trinket2: Rune of Perfection; main_hand: Butcher's Cleaver; off_hand: Diamond Hammer; ranged: Kajaric Icon

No-known-source sample (15 of 215, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6478 Rat Stompers; 7956 Bronze Warhammer; 10047 Simple Kilt; 10421 Rough Copper Vest; 14147 Cavedweller Bracers

### Band 30 (dwarf, 0000000000000000-253130030004000000-0000000000000000)

Set DPS (verified): 71.5. Weights run: 0.9s. Verify run: 1.1s. 373 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.432 ± 0.100, crit=7.348 ± 0.337, hit=5.948 ± 0.471, melee_haste=2.835 ± 0.417

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 | yes | Defender's Leather Helm (252455, -0.07 DPS) [crafted]; Veteran's Chain Helm (250498, -0.12 DPS, sim-verified) [crafted]; Crusader's Chain Helm (250502, -0.14 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | sim-verified (71.1 DPS) | yes | River Pride Choker (13087, -0.06 DPS) [world_drop]; Sentinel's Medallion (19541, -0.22 DPS) [rep]; Ghostshard Talisman (7731, -1.11 DPS, sim-verified) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 14.7 | yes | Golden Scale Shoulders (3841, -0.03 DPS) [crafted]; Mail Combat Spaulders (6404, -0.03 DPS) [world_drop]; Barbaric Iron Shoulders (7913, -0.10 DPS, sim-verified) [crafted] |
| back | Sergeant Major's Cape (16315) | PvP rank 9 · Sergeant Major · Alliance [pvp] | sim-verified (71.3 DPS) | yes | Hawkeye's Cloak (14593, -0.02 DPS) [world_drop]; Lambent Scale Cloak (4706, -0.06 DPS) [world_drop]; Wolfmaster Cape (6314, -1.34 DPS, sim-verified) [dungeon] |
| chest | Shining Silver Breastplate (2870) | Blacksmithing [crafted] | 28.0 | yes | Veteran's Silvered Chain Shirt (250518, -0.19 DPS) [crafted]; Hard Gold Cuirass (250533, -0.21 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.23 DPS, sim-verified) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 | yes | Yorgen Bracers (13012, -0.08 DPS, sim-verified) [world_drop]; Bands of Serra'kis (6902, -0.14 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.19 DPS) [world_drop] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 102.9 | yes | Gauntlets of Ogre Strength (3341, +0.00 DPS, sim-verified) [world]; Bonefist Gauntlets (4465, -3.00 DPS) [world]; Mail Combat Gauntlets (4075, -3.01 DPS) [world_drop] |
| waist | Girdle of Golem Strength (9405) (or Highlander's Plate Girdle (20126)) | World drop [world_drop] | 24.0 | yes | Highlander's Plate Girdle (20126, +0.00 DPS, sim-verified) [rep]; Highlander's Chain Girdle (20090, -0.00 DPS) [rep]; Highlander's Leather Girdle (20117, -0.00 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Golden Scale Leggings (3843, -0.14 DPS) [crafted]; Chausses of Westfall (6087, -0.14 DPS) [quest]; Veteran's Silvered Chain Leggings (250523, -0.81 DPS, sim-verified) [crafted] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 17.0 | yes | Hard Gold Boots (250534, +0.00 DPS, sim-verified) [crafted]; Brawler's Leather Boots (252439, -0.17 DPS) [crafted]; Alacritous Treads (277234, -0.17 DPS) [quest] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 17.3 | yes | Ironspine's Eye (7686, -0.19 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.26 DPS) [dungeon]; Insurgent's Band (272067, -0.29 DPS) [vendor] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 14.6 | yes | Ironspine's Eye (7686, -0.08 DPS, sim-verified) [dungeon]; Silverlaine's Family Seal (6321, -0.16 DPS) [dungeon]; Protector's Band (20439, -0.17 DPS) [rep] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (70.3 DPS) | yes | Minor Recombobulator (4381, +0.00 DPS) [crafted]; Gnomish Universal Remote (7506, +0.00 DPS) [crafted]; Talisman of Arathor (21119, -1.31 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 345.4 | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 334.9 | yes | Shoni's Disarming Tool (9608, -3.89 DPS) [quest]; Shield of Thorsen (13079, -11.34 DPS) [world_drop]; Swinetusk Shank (6691, -11.82 DPS, sim-verified) [dungeon] |
| ranged | - | - |  |  |  |

**New at 30:** head: Tusken Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Sergeant Major's Cape; chest: Shining Silver Breastplate; wrist: Pugilist Bracers; waist: Girdle of Golem Strength; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Pronged Reaver; off_hand: Ironspine's Fist

No-known-source sample (15 of 373, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (dwarf, 0000000000000000-253130030005102051-0000000000000000)

Set DPS (verified): 93.3. Weights run: 1.1s. Verify run: 1.1s. 512 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.386 ± 0.051, crit=7.710 ± 0.213, hit=5.187 ± 0.448, melee_haste=3.244 ± 0.122

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 133.9 | yes | Hard Gold Coif (250537, -1.02 DPS, sim-verified) [crafted]; White Bandit Mask (10008, -3.99 DPS) [crafted]; Tusken Helm (6686, -4.00 DPS) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; River Pride Choker (13087, -0.22 DPS) [world_drop]; Gazlowe's Charm (13088, -0.22 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 | yes | Imperial Leather Spaulders (4737, -0.14 DPS, sim-verified) [world_drop]; Wrangling Spaulders (15698, -0.19 DPS) [quest]; Sunburn Spaulders (274751, -0.21 DPS) [vendor] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 14.3 | yes | Sergeant Major's Cape (16315, -0.18 DPS) [pvp]; Hawkeye's Cloak (14593, -0.21 DPS) [world_drop]; Wolfmaster Cape (6314, -1.92 DPS, sim-verified) [dungeon] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 30.7 | yes | Golden Scale Cuirass (3845, -0.10 DPS) [crafted]; Mail Combat Armor (4074, -0.17 DPS) [world_drop]; Shining Silver Breastplate (2870, -0.46 DPS, sim-verified) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Ravager's Armguards (14770, -0.17 DPS) [world_drop]; Pugilist Bracers (4438, -0.23 DPS, sim-verified) [dungeon]; Yorgen Bracers (13012, -0.25 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 127.9 | yes | Fletcher's Gloves (7348, -0.74 DPS) [crafted]; Shadowskin Gloves (18238, -0.74 DPS) [crafted]; Dragonscale Gauntlets (8347, -0.84 DPS, sim-verified) [crafted] |
| waist | Highlander's Chain Girdle (20089) | The League of Arathor [rep] | 115.9 | yes | Highlander's Leather Girdle (20116, -0.15 DPS, sim-verified) [rep]; Girdle of Golem Strength (9405, -3.40 DPS) [world_drop]; Scarlet Belt (10329, -3.40 DPS) [dungeon] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 | yes | Firemane Leggings (13129, -0.23 DPS, sim-verified) [world_drop]; Orcish War Leggings (7929, -0.30 DPS) [crafted]; Ferine Leggings (6690, -0.59 DPS) [dungeon] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 24.7 | yes | Blackforge Greaves (6423, -0.11 DPS) [world_drop]; Skirmisher's Mail Boots (252564, -0.12 DPS, sim-verified) [crafted]; Ironheel Boots (4653, -0.17 DPS) [quest] |
| finger1 | Assault Band (13095) | World drop [world_drop] | 20.0 | yes | Thunderbrow Ring (13097, -0.11 DPS) [world_drop]; Suspicious Spare Part (274754, -0.22 DPS) [vendor]; Insurgent's Band (272066, -0.30 DPS) [vendor] |
| finger2 | Protector's Band (19515) | Silverwing Sentinels [rep] | 19.1 | yes | Protector's Band (19517, -0.18 DPS) [rep]; Suspicious Spare Part (274754, -0.19 DPS) [vendor]; Thunderbrow Ring (13097, -0.36 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (80.0 DPS) | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Rune of Duty (21567, -0.97 DPS, sim-verified) [rep] |
| main_hand | Ardent Custodian (868) | World drop [world_drop] | 460.0 | yes | Curve-bladed Ripper (2815, +0.00 DPS, sim-verified) [world_drop]; Bonebiter (6830, +0.00 DPS) [quest]; Illusionary Rod (7713, +0.00 DPS) [dungeon] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-verified (91.8 DPS) | yes | Shoni's Disarming Tool (9608, -7.94 DPS) [quest]; Curve-bladed Ripper (2815, -12.75 DPS, sim-verified) [world_drop]; Salbac Shield (4652, -15.59 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; neck: Ghostshard Talisman; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Chain Girdle; legs: Scarlet Leggings; feet: Prowler's Leather Shoes; finger1: Assault Band; finger2: Protector's Band; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Ardent Custodian; off_hand: Jhordy's Misplaced Screwdriver; ranged: Totem of Ancestral Protection

No-known-source sample (15 of 512, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 50 (dwarf, 5500000000000000-253130030005102051-0000000000000000)

Set DPS (verified): 125.6. Weights run: 1.2s. Verify run: 1.2s. 674 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.512 ± 0.061, crit=9.934 ± 0.268, hit=7.977 ± 0.662, melee_haste=3.814 ± 0.090

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 165.1 | yes | Knight-Lieutenant's Mail Helmet (223075, -0.03 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -0.94 DPS) [dungeon]; Hard Gold Coif (250537, -4.96 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 16.7 | yes | Kaleidoscope Chain (13084, -0.24 DPS) [world_drop]; River Pride Choker (13087, -0.31 DPS) [world_drop]; Ghostshard Talisman (7731, -0.96 DPS, sim-verified) [dungeon] |
| shoulder | Knight-Lieutenant's Mail Epaulets (223073) | Captain Dirgehammer [vendor] | 157.1 | yes | Prowler's Leather Shoulder (252534, -0.62 DPS, sim-verified) [crafted]; Failed Flying Experiment (9647, -4.77 DPS) [quest]; Skulker's Leather Shoulder (252535, -4.83 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | sim-verified (123.5 DPS) | yes | Pridelord Cape (14673, -0.16 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.18 DPS) [pvp]; Bloodlust Cape (14801, -1.66 DPS, sim-verified) [world_drop] |
| chest | Knight's Mail Armor (223078) | Captain Dirgehammer [vendor] | 163.1 | yes | Wildthorn Mail (12624, -1.92 DPS, sim-verified) [crafted]; Kolkar Marauder Chain (6773, -4.76 DPS) [quest]; Warbear Harness (15064, -4.77 DPS) [crafted] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 | yes | Branded Leather Bracers (19508, -0.29 DPS) [dungeon]; Skulker's Leather Bracers (252540, -0.34 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.58 DPS, sim-verified) [crafted] |
| hands | Fists of The Five Thunders (227022) | Mokvar [vendor] | sim-verified (123.1 DPS) | yes | Sergeant Major's Mail Gauntlets (223076, -0.07 DPS) [vendor]; Dragonscale Gauntlets (8347, -0.54 DPS) [crafted]; Gloves of Holy Might (867, -1.23 DPS, sim-verified) [world_drop] |
| waist | Highlander's Chain Girdle (20088) (or Highlander's Leather Girdle (20115)) | The League of Arathor [rep] | 159.1 | yes | Highlander's Leather Girdle (20115, +0.00 DPS, sim-verified) [rep]; Highlander's Chain Girdle (20089, -0.43 DPS) [rep]; Highlander's Cloth Girdle (20097, -0.72 DPS) [rep] |
| legs | Knight's Mail Legplates (223074) | Captain Dirgehammer [vendor] | sim-verified (124.0 DPS) | yes | Stormshroud Pants (15057, -2.14 DPS, sim-verified) [crafted]; Scarlet Leggings (10330, -4.38 DPS) [dungeon]; Firemane Leggings (13129, -4.53 DPS) [world_drop] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 79.8 | yes | Prowler's Leather Boots (252468, -1.33 DPS, sim-verified) [crafted]; Skulker's Leather Boots (252469, -1.85 DPS) [crafted]; Skirmisher's Mail Sabatons (252578, -1.95 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 99.8 | yes | Assault Band (13095, -2.89 DPS) [world_drop]; Thunderbrow Ring (13097, -2.98 DPS) [world_drop]; Insurgent's Band (272065, -3.07 DPS) [vendor] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 24.6 | yes | Assault Band (13095, -0.17 DPS) [world_drop]; Thunderbrow Ring (13097, -0.26 DPS) [world_drop]; Protector's Band (19515, -0.37 DPS, sim-verified) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (123.1 DPS) | yes | Ankh of Life (1713, -0.69 DPS, sim-verified) [world_drop]; Tidal Charm (1404, -2.60 DPS) [vendor]; Guardian Talisman (1490, -2.60 DPS) [quest] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (123.1 DPS) | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, -0.50 DPS, sim-verified) [world_drop] |
| main_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | sim-verified (123.1 DPS) | yes | Flurry Axe (871, +0.00 DPS, sim-verified) [world_drop]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Executioner's Cleaver (13018, +0.00 DPS) [world_drop] |
| off_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | 553.3 | yes | Claw of Celebras (17738, -2.57 DPS) [dungeon]; Might of Hakkar (10838, -3.10 DPS, sim-verified) [world]; Shoni's Disarming Tool (9608, -11.89 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 50:** neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Mail Epaulets; chest: Knight's Mail Armor; wrist: Bracers of the Stone Princess; hands: Fists of The Five Thunders; waist: Highlander's Chain Girdle; legs: Knight's Mail Legplates; feet: Greaves of Withering Despair; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Frozen Heart of the Mountain; trinket2: Thunderbrew's Boot Flask; main_hand: Dawn's Edge; off_hand: Hammer of the Northern Wind

No-known-source sample (15 of 674, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (dwarf, 5533220000000000-253130030005102051-0000000000000000)

Set DPS (verified): 229.4. Weights run: 1.1s. Verify run: 1.3s. 1303 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.519 ± 0.068, crit=10.313 ± 0.324, hit=11.159 ± 1.127, melee_haste=3.972 ± 0.142

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soulcrusher Headpiece (240096) | Leonid Barthalomew the Revered [vendor] | 433.2 | yes | Bloodvine Goggles (19999, -2.36 DPS) [crafted]; Mask of the Unforgiven (13404, -2.82 DPS, sim-verified) [dungeon]; Ragefury Eyepatch (11735, -4.77 DPS) [dungeon] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (220.5 DPS) | yes | Blazefury Medallion (17111, -0.39 DPS, sim-verified) [world]; Beads of Ogre Might (22150, -1.18 DPS) [quest]; Amulet of the Darkmoon (19491, -4.99 DPS) [quest] |
| shoulder | Soulcrusher Mantle (240125) | Leonid Barthalomew the Revered [vendor] | 256.0 | yes | Soulcrusher Epaulets (240135, +0.00 DPS, sim-verified) [vendor]; Knight-Lieutenant's Mail Epaulets (223073, -3.37 DPS) [vendor]; Darkspear Pauldrons (272105, -3.52 DPS) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-verified (225.3 DPS) | yes | Earthweave Cloak (21187, -0.73 DPS) [quest]; Arcanoweave Cloak (272411, -1.01 DPS) [vendor]; Chromatic Cloak (18509, -3.85 DPS, sim-verified) [crafted] |
| chest | Soulcrusher Armor (240128) | Leonid Barthalomew the Revered [vendor] | 609.9 | yes | Stormshroud Armor (15056, -11.58 DPS) [crafted]; Dawn Armor (252483, -12.07 DPS) [crafted]; Bloodsoul Breastplate (19690, -15.34 DPS, sim-verified) [crafted] |
| wrist | Soulcrusher Vambraces (240137) | Leonid Barthalomew the Revered [vendor] | 163.6 | yes | Primal Batskin Bracers (19687, -1.61 DPS) [crafted]; Rockfury Bracers (21186, -1.87 DPS) [quest]; Soulcrusher Bracers (240108, -3.56 DPS, sim-verified) [vendor] |
| hands | Soulcrusher Grips (240130) | Leonid Barthalomew the Revered [vendor] | 352.8 | yes | Soulcrusher Mitts (240122, -2.31 DPS) [vendor]; Stormshroud Gloves (21278, -3.49 DPS) [crafted]; Soulcrusher Handguards (240095, -4.02 DPS, sim-verified) [vendor] |
| waist | Soulcrusher Girdle (240099) | Leonid Barthalomew the Revered [vendor] | 288.8 | yes | Soulcrusher Waistguard (240107, +0.00 DPS, sim-verified) [vendor]; Soulcrusher Belt (240136, -2.61 DPS) [vendor]; Highlander's Chain Girdle (20043, -3.98 DPS) [rep] |
| legs | Soulcrusher Leggings (240134) | Leonid Barthalomew the Revered [vendor] | sim-verified (226.5 DPS) | yes | Sentinel's Leather Pants (237818, -3.10 DPS) [vendor]; Stormshroud Pants (15057, -3.61 DPS) [crafted]; Sentinel's Chain Leggings (237819, -5.07 DPS, sim-verified) [vendor] |
| feet | Soulcrusher Sabatons (240102) | Leonid Barthalomew the Revered [vendor] | 135.6 | yes | Bloodvine Boots (19684, -0.87 DPS) [crafted]; Greaves of Withering Despair (22240, -0.87 DPS) [dungeon]; Fine Dawn Treaders (227815, -3.13 DPS, sim-verified) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (219.8 DPS) | yes | Signet Ring of the Bronze Dragonflight (234034, -4.05 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234030, -4.20 DPS) [vendor]; Wrath of Cenarius (21190, -4.25 DPS, sim-verified) [quest] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (219.8 DPS) | yes | Wrath of Cenarius (21190, -3.86 DPS, sim-verified) [quest]; Signet Ring of the Bronze Dragonflight (234034, -3.91 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234030, -4.05 DPS) [vendor] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (218.7 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -8.30 DPS, sim-verified) [crafted] |
| trinket2 | - | - |  |  |  |
| main_hand | Ebon Hand (19170) | Blacksmithing [crafted] | sim-verified (220.5 DPS) | yes | Grand Marshal's Sunderer (234566, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor]; Persuader (22384, -4.21 DPS, sim-verified) [crafted] |
| off_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (220.5 DPS) | yes | Grand Marshal's Left Hand Blade (18847, +0.00 DPS) [vendor]; Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [vendor]; Force Reactive Disk (18168, -16.74 DPS, sim-verified) [crafted] |
| ranged | Totem of the Storm (23199) (or Totem of Thunder (228176), Tidal Totem (272431), Totem of the Storm (272432), Burning Totem (272433), Totem of Urgency (279249), Totem of Ancestral Protection (249443), Kajaric Icon (206387), Polished Driftwood Icon (249398), Tempest Icon (206382), Dyadic Icon (206381), Galvanic Icon (206386), Sulfurous Icon (206388), Voltaic Icon (225838)) | World drop [world_drop] | 0.0 | yes | Totem of Thunder (228176, +0.00 DPS, sim-verified) [vendor]; Tidal Totem (272431, +0.00 DPS) [vendor]; Totem of the Storm (272432, +0.00 DPS) [world_drop] |

**New at 60:** head: Soulcrusher Headpiece; neck: Medallion of the Dawn; shoulder: Soulcrusher Mantle; back: Howler's Furs; chest: Soulcrusher Armor; wrist: Soulcrusher Vambraces; hands: Soulcrusher Grips; waist: Soulcrusher Girdle; legs: Soulcrusher Leggings; feet: Soulcrusher Sabatons; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Darkmoon Card: Maelstrom; trinket2: Guardian Talisman; main_hand: Ebon Hand; off_hand: Shadowsong's Sorrow; ranged: Totem of the Storm

No-known-source sample (15 of 1303, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

## Horde

### Band 20 (orc, 0000000000000000-253100000000000000-0000000000000000)

Set DPS (verified): 45.8. Weights run: 0.9s. Verify run: 1.0s. 211 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=0.235 ± 0.025, crit=4.333 ± 0.103, hit=4.241 ± 0.303, melee_haste=2.440 ± 0.609

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 | yes | Defender's Leather Hood (252447, -0.19 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.61 DPS) [crafted]; Brawler's Leather Hood (252504, -0.63 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 1.4 | yes | Scholarly Pendant (277203, -0.05 DPS) [quest]; Tarnished Locket (279870, -0.05 DPS) [quest]; Erudite's Amulet (277204, -0.05 DPS, sim-verified) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 | yes | Silvered Bronze Shoulders (3481, +0.00 DPS, sim-verified) [crafted]; Serpent's Shoulders (5404, -0.17 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.21 DPS) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 | yes | Catacomb Cloak (279899, -0.07 DPS) [quest]; Subterranean Cape (14149, -0.07 DPS) [dungeon]; Grave Shroud (279865, -0.18 DPS, sim-verified) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 | yes | Veteran's Chain Shirt (250488, -0.17 DPS, sim-verified) [crafted]; Defender's Leather Armor (252434, -0.18 DPS) [crafted]; Totemic Leather Armor (252435, -0.21 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 | yes | Cryptwalker Bracers (280095, -0.09 DPS, sim-verified) [quest]; Raptorcrest Bracers (270010, -0.14 DPS) [quest]; Bristlebark Bindings (14569, -0.18 DPS) [world_drop] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 60.7 | yes | Thorbia's Gauntlets (12994, +0.00 DPS, sim-verified) [world_drop]; Gold-flecked Gloves (5195, -1.61 DPS) [dungeon]; Blackened Defias Gloves (10401, -1.68 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 | yes | Cobrahn's Grasp (6460, -0.06 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.21 DPS) [world]; Support Girdle (1215, -0.28 DPS) [world] |
| legs | Veteran's Chain Leggings (250493) | Blacksmithing [crafted] | 19.2 | yes | Defender's Leather Pants (252445, -0.03 DPS, sim-verified) [crafted]; Totemic Leather Pants (252446, -0.04 DPS) [crafted]; Hulking Leggings (14748, -0.09 DPS) [world_drop] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 11.2 | yes | Veteran's Boots (250503, -0.03 DPS, sim-verified) [crafted]; Guard's Boots (250504, -0.04 DPS) [crafted]; Defender's Leather Boots (252441, -0.04 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.9 | yes | The 1 Ring (8350, -0.23 DPS) [world]; Ring of the Moon (12052, -0.24 DPS) [world_drop]; Demon Band (12054, -0.48 DPS, sim-verified) [world_drop] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | sim-verified (45.8 DPS) | yes | The 1 Ring (8350, -0.13 DPS) [world]; Ring of the Moon (12052, -0.14 DPS) [world_drop]; Demon Band (12054, -0.47 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 237.0 | yes | Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest]; Hammerbone (270018, +0.00 DPS) [quest] |
| off_hand | Diamond Hammer (2194) | World drop [world_drop] | 232.8 | yes | Blackfang (2236, -0.23 DPS, sim-verified) [world_drop]; Redbeard Crest (12997, -7.62 DPS) [world_drop]; Bear Buckler (4821, -7.82 DPS) [vendor] |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Loop of Sacrifice; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Butcher's Cleaver; off_hand: Diamond Hammer; ranged: Kajaric Icon

No-known-source sample (15 of 211, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7956 Bronze Warhammer; 10047 Simple Kilt; 10421 Rough Copper Vest; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 15402 Noosegrip Gauntlets; 20425 Advisor's Gnarled Staff; 20441 Scout's Blade

### Band 30 (orc, 0000000000000000-253130030004000000-0000000000000000)

Set DPS (verified): 74.9. Weights run: 0.9s. Verify run: 1.0s. 370 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.432 ± 0.100, crit=7.348 ± 0.337, hit=5.948 ± 0.471, melee_haste=2.835 ± 0.417

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 | yes | Defender's Leather Helm (252455, -0.07 DPS) [crafted]; Veteran's Chain Helm (250498, -0.13 DPS, sim-verified) [crafted]; Crusader's Chain Helm (250502, -0.14 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Kaleidoscope Chain (13084, -0.10 DPS, sim-verified) [world_drop]; River Pride Choker (13087, -0.21 DPS) [world_drop]; Scout's Medallion (19537, -0.37 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 14.7 | yes | Golden Scale Shoulders (3841, -0.03 DPS) [crafted]; Mail Combat Spaulders (6404, -0.03 DPS) [world_drop]; Barbaric Iron Shoulders (7913, -0.14 DPS, sim-verified) [crafted] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Hawkeye's Cloak (14593, -0.03 DPS) [world_drop]; Lambent Scale Cloak (4706, -0.07 DPS) [world_drop] |
| chest | Shining Silver Breastplate (2870) | Blacksmithing [crafted] | 28.0 | yes | Veteran's Silvered Chain Shirt (250518, -0.19 DPS) [crafted]; Hard Gold Cuirass (250533, -0.21 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.26 DPS, sim-verified) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 | yes | Bands of Serra'kis (6902, -0.14 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.19 DPS) [world_drop]; Yorgen Bracers (13012, -0.47 DPS, sim-verified) [world_drop] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 102.9 | yes | Gauntlets of Ogre Strength (3341, +0.00 DPS, sim-verified) [world]; Warsong Gauntlets (16978, -2.93 DPS) [quest]; Bonefist Gauntlets (4465, -3.00 DPS) [world] |
| waist | Girdle of Golem Strength (9405) (or Defiler's Plate Girdle (20207)) | World drop [world_drop] | 24.0 | yes | Defiler's Plate Girdle (20207, +0.00 DPS, sim-verified) [rep]; Defiler's Chain Girdle (20152, -0.00 DPS) [rep]; Defiler's Leather Girdle (20191, -0.00 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Golden Scale Leggings (3843, -0.14 DPS) [crafted]; Slayer's Pants (14757, -0.14 DPS) [world_drop]; Veteran's Silvered Chain Leggings (250523, -1.38 DPS, sim-verified) [crafted] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 17.0 | yes | Hard Gold Boots (250534, +0.00 DPS, sim-verified) [crafted]; Brawler's Leather Boots (252439, -0.17 DPS) [crafted]; Alacritous Treads (277234, -0.17 DPS) [quest] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 17.3 | yes | Ironspine's Eye (7686, -0.19 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.26 DPS) [dungeon]; Band of the Fist (17694, -0.27 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 14.6 | yes | Silverlaine's Family Seal (6321, -0.16 DPS) [dungeon]; Band of the Fist (17694, -0.17 DPS) [quest]; Ironspine's Eye (7686, -0.47 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (74.1 DPS) | yes | Minor Recombobulator (4381, +0.00 DPS) [crafted]; Gnomish Universal Remote (7506, +0.00 DPS) [crafted]; Defiler's Talisman (21120, -1.51 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 345.4 | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 334.9 | yes | Shield of Thorsen (13079, -11.34 DPS) [world_drop]; Slayer's Shield (15892, -11.36 DPS) [world_drop]; Swinetusk Shank (6691, -16.12 DPS, sim-verified) [dungeon] |
| ranged | - | - |  |  |  |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Shining Silver Breastplate; wrist: Pugilist Bracers; waist: Girdle of Golem Strength; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Pronged Reaver; off_hand: Ironspine's Fist

No-known-source sample (15 of 370, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7956 Bronze Warhammer; 7958 Bronze Battle Axe; 9362 Brilliant Gold Ring; 10047 Simple Kilt

### Band 40 (orc, 0000000000000000-253130030005102051-0000000000000000)

Set DPS (verified): 84.9. Weights run: 1.1s. Verify run: 1.0s. 510 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.386 ± 0.051, crit=7.710 ± 0.213, hit=5.187 ± 0.448, melee_haste=3.244 ± 0.122

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 133.9 | yes | Hard Gold Coif (250537, -1.00 DPS, sim-verified) [crafted]; White Bandit Mask (10008, -3.99 DPS) [crafted]; Tusken Helm (6686, -4.00 DPS) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Ethereal Talisman (4430, +0.00 DPS, sim-verified) [quest]; Kaleidoscope Chain (13084, -0.17 DPS) [world_drop]; River Pride Choker (13087, -0.22 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 | yes | Wrangling Spaulders (15698, -0.19 DPS) [quest]; Sunburn Spaulders (274751, -0.21 DPS) [vendor]; Imperial Leather Spaulders (4737, -0.67 DPS, sim-verified) [world_drop] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Hawkeye's Cloak (14593, -0.05 DPS) [world_drop]; Lambent Scale Cloak (4706, -0.07 DPS) [world_drop] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 30.7 | yes | Golden Scale Cuirass (3845, -0.10 DPS) [crafted]; Mail Combat Armor (4074, -0.17 DPS) [world_drop]; Shining Silver Breastplate (2870, -0.52 DPS, sim-verified) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Ravager's Armguards (14770, -0.17 DPS) [world_drop]; Pugilist Bracers (4438, -0.21 DPS, sim-verified) [dungeon]; Darkspear Armsplints (4132, -0.22 DPS) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 127.9 | yes | Fletcher's Gloves (7348, -0.74 DPS) [crafted]; Shadowskin Gloves (18238, -0.74 DPS) [crafted]; Dragonscale Gauntlets (8347, -0.90 DPS, sim-verified) [crafted] |
| waist | Defiler's Chain Girdle (20153) | The Defilers [rep] | 115.9 | yes | Defiler's Leather Girdle (20192, -0.00 DPS, sim-verified) [rep]; Tharg's Shoelace (9705, -3.33 DPS) [quest]; Girdle of Golem Strength (9405, -3.40 DPS) [world_drop] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 | yes | Firemane Leggings (13129, -0.26 DPS, sim-verified) [world_drop]; Orcish War Leggings (7929, -0.30 DPS) [crafted]; Ferine Leggings (6690, -0.59 DPS) [dungeon] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 24.7 | yes | Blackforge Greaves (6423, -0.11 DPS) [world_drop]; Skirmisher's Mail Boots (252564, -0.21 DPS, sim-verified) [crafted]; Skulker's Leather Shoes (252531, -0.24 DPS) [crafted] |
| finger1 | Assault Band (13095) | World drop [world_drop] | 20.0 | yes | Thunderbrow Ring (13097, -0.11 DPS) [world_drop]; Suspicious Spare Part (274754, -0.22 DPS) [vendor]; Insurgent's Band (272066, -0.30 DPS) [vendor] |
| finger2 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 19.1 | yes | Legionnaire's Band (19513, -0.18 DPS) [rep]; Suspicious Spare Part (274754, -0.19 DPS) [vendor]; Thunderbrow Ring (13097, -0.43 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ardent Custodian (868) | World drop [world_drop] | 460.0 | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Illusionary Rod (7713, +0.00 DPS) [dungeon]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor] |
| off_hand | Curve-bladed Ripper (2815) | World drop [world_drop] | 439.6 | yes | Jhordy's Misplaced Screwdriver (274753, -0.33 DPS, sim-verified) [vendor]; Skullance Shield (13081, -15.63 DPS) [world_drop]; Pit Fighter's Shield (4507, -15.68 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; shoulder: Hard Gold Pauldrons; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Chain Girdle; legs: Scarlet Leggings; feet: Prowler's Leather Shoes; finger1: Assault Band; finger2: Legionnaire's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Ardent Custodian; off_hand: Curve-bladed Ripper; ranged: Totem of Ancestral Protection

No-known-source sample (15 of 510, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7748 Forcestone Buckler; 7956 Bronze Warhammer

### Band 50 (orc, 5500000000000000-253130030005102051-0000000000000000)

Set DPS (verified): 120.9. Weights run: 1.2s. Verify run: 1.1s. 660 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.512 ± 0.061, crit=9.934 ± 0.268, hit=7.977 ± 0.662, melee_haste=3.814 ± 0.090

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 165.1 | yes | Blood Guard's Mail Helmet (220820, -0.54 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -0.94 DPS) [dungeon]; Blood Guard's Inscribed Skullcap (220842, -0.94 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 16.7 | yes | Ghostshard Talisman (7731, -0.10 DPS) [dungeon]; Ethereal Talisman (4430, -0.17 DPS) [quest]; Woven Ivy Necklace (19159, -0.95 DPS, sim-verified) [quest] |
| shoulder | Blood Guard's Mail Epaulets (220823) | Lady Palanseer [vendor] | 157.1 | yes | Blood Guard's Pulsing Shoulders (220849, -0.65 DPS) [vendor]; Blood Guard's Inscribed Shoulder Pads (220841, -1.20 DPS, sim-verified) [vendor]; Prowler's Leather Shoulder (252534, -4.72 DPS) [crafted] |
| back | Bloodlust Cape (14801) | World drop [world_drop] | 18.0 | yes | Wolfmaster Cape (6314, -0.29 DPS) [dungeon]; Battlehard Cape (11858, -0.29 DPS) [quest]; Pridelord Cape (14673, -0.41 DPS, sim-verified) [world_drop] |
| chest | Stone Guard's Mail Armor (220826) | Lady Palanseer [vendor] | 163.1 | yes | Stone Guard's Inscribed Chestpiece (220838, -1.59 DPS, sim-verified) [vendor]; Wildthorn Mail (12624, -3.01 DPS) [crafted]; Stone Guard's Pulsing Breastplate (220844, -3.01 DPS) [vendor] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 | yes | Branded Leather Bracers (19508, -0.29 DPS) [dungeon]; Skulker's Leather Bracers (252540, -0.34 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.50 DPS, sim-verified) [crafted] |
| hands | Fists of The Five Thunders (227022) | Mokvar [vendor] | sim-verified (118.6 DPS) | yes | First Sergeant's Mail Gauntlets (220831, -0.07 DPS) [vendor]; Dragonscale Gauntlets (8347, -0.54 DPS) [crafted]; Gloves of Holy Might (867, -1.61 DPS, sim-verified) [world_drop] |
| waist | Defiler's Chain Girdle (20151) (or Defiler's Leather Girdle (20193)) | The Defilers [rep] | 159.1 | yes | Defiler's Leather Girdle (20193, +0.00 DPS, sim-verified) [rep]; Defiler's Chain Girdle (20153, -0.43 DPS) [rep]; Highlander's Mail Girdle (20118, -0.72 DPS) [vendor] |
| legs | Stone Guard's Mail Legplates (220834) | Lady Palanseer [vendor] | sim-verified (118.6 DPS) | yes | Stone Guard's Inscribed Legplates (220839, -0.87 DPS) [vendor]; Stormshroud Pants (15057, -1.63 DPS, sim-verified) [crafted]; Stone Guard's Pulsing Legplates (220847, -3.01 DPS) [vendor] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 79.8 | yes | Prowler's Leather Boots (252468, -0.28 DPS, sim-verified) [crafted]; Skulker's Leather Boots (252469, -1.85 DPS) [crafted]; Skirmisher's Mail Sabatons (252578, -1.95 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 99.8 | yes | White Bone Band (11862, -2.74 DPS) [quest]; Assault Band (13095, -2.89 DPS) [world_drop]; Band of Allegiance (18585, -2.96 DPS) [quest] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 24.6 | yes | Legionnaire's Band (19512, -0.16 DPS) [rep]; Assault Band (13095, -0.17 DPS) [world_drop]; White Bone Band (11862, -0.79 DPS, sim-verified) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (117.5 DPS) | yes | Ankh of Life (1713, -2.26 DPS, sim-verified) [world_drop]; Tidal Charm (1404, -3.54 DPS) [vendor]; Guardian Talisman (1490, -3.54 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (117.5 DPS) | yes | Ankh of Life (1713, -0.35 DPS, sim-verified) [world_drop]; Tidal Charm (1404, -2.60 DPS) [vendor]; Guardian Talisman (1490, -2.60 DPS) [quest] |
| main_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | sim-verified (117.5 DPS) | yes | Flurry Axe (871, +0.00 DPS, sim-verified) [world_drop]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Executioner's Cleaver (13018, +0.00 DPS) [world_drop] |
| off_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | 553.3 | yes | Claw of Celebras (17738, -2.57 DPS) [dungeon]; White Bone Shredder (11863, -3.86 DPS) [quest]; Might of Hakkar (10838, -4.73 DPS, sim-verified) [world] |
| ranged | - | - |  |  |  |

**New at 50:** neck: Skibi's Pendant; shoulder: Blood Guard's Mail Epaulets; back: Bloodlust Cape; chest: Stone Guard's Mail Armor; wrist: Bracers of the Stone Princess; hands: Fists of The Five Thunders; waist: Defiler's Chain Girdle; legs: Stone Guard's Mail Legplates; feet: Greaves of Withering Despair; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Dawn's Edge; off_hand: Hammer of the Northern Wind

No-known-source sample (15 of 660, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60 (orc, 5533220000000000-253130030005102051-0000000000000000)

Set DPS (verified): 225.6. Weights run: 1.1s. Verify run: 1.3s. 1242 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.519 ± 0.068, crit=10.313 ± 0.324, hit=11.159 ± 1.127, melee_haste=3.972 ± 0.142

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soulcrusher Headpiece (240096) | Leonid Barthalomew the Revered [vendor] | 433.2 | yes | Bloodvine Goggles (19999, -2.36 DPS) [crafted]; Ragefury Eyepatch (11735, -4.77 DPS) [dungeon]; Mask of the Unforgiven (13404, -5.38 DPS, sim-verified) [dungeon] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (218.6 DPS) | yes | Beads of Ogre Might (22150, -1.18 DPS) [quest]; Blazefury Medallion (17111, -2.26 DPS, sim-verified) [world]; Amulet of the Darkmoon (19491, -4.99 DPS) [quest] |
| shoulder | Warlord's Mail Pauldrons (231654) | Lady Palanseer [vendor] | 290.0 | yes | Soulcrusher Mantle (240125, +0.00 DPS, sim-verified) [vendor]; Champion's Mail Pauldrons (227154, -0.29 DPS) [vendor]; Soulcrusher Epaulets (240135, -4.05 DPS) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-verified (222.8 DPS) | yes | Earthweave Cloak (21187, -0.73 DPS) [quest]; Arcanoweave Cloak (272411, -1.01 DPS) [vendor]; Chromatic Cloak (18509, -3.83 DPS, sim-verified) [crafted] |
| chest | Soulcrusher Armor (240128) | Leonid Barthalomew the Revered [vendor] | 609.9 | yes | Warlord's Mail Hauberk (231653, -9.00 DPS, sim-verified) [vendor]; Legionnaire's Mail Hauberk (227157, -11.18 DPS) [vendor]; Bloodsoul Breastplate (19690, -11.41 DPS) [crafted] |
| wrist | Soulcrusher Vambraces (240137) | Leonid Barthalomew the Revered [vendor] | 163.6 | yes | Primal Batskin Bracers (19687, -1.61 DPS) [crafted]; Rockfury Bracers (21186, -1.87 DPS) [quest]; Soulcrusher Bracers (240108, -3.25 DPS, sim-verified) [vendor] |
| hands | Soulcrusher Grips (240130) | Leonid Barthalomew the Revered [vendor] | 352.8 | yes | Soulcrusher Mitts (240122, -2.31 DPS) [vendor]; General's Mail Vices (231655, -2.34 DPS) [vendor]; Soulcrusher Handguards (240095, -3.32 DPS, sim-verified) [vendor] |
| waist | Soulcrusher Girdle (240099) | Leonid Barthalomew the Revered [vendor] | 288.8 | yes | Soulcrusher Waistguard (240107, +0.00 DPS, sim-verified) [vendor]; Soulcrusher Belt (240136, -2.61 DPS) [vendor]; Defiler's Chain Girdle (20150, -3.98 DPS) [rep] |
| legs | Soulcrusher Leggings (240134) | Leonid Barthalomew the Revered [vendor] | sim-verified (223.9 DPS) | yes | General's Mail Legguards (231658, -2.77 DPS) [vendor]; Sentinel's Leather Pants (237818, -3.10 DPS) [vendor]; Sentinel's Chain Leggings (237819, -4.97 DPS, sim-verified) [vendor] |
| feet | General's Mail Greaves (231656) | Lady Palanseer [vendor] | 145.6 | yes | Soulcrusher Sabatons (240102, +0.00 DPS, sim-verified) [vendor]; Blood Guard's Mail Greaves (227158, -0.22 DPS) [vendor]; Fine Dawn Treaders (227815, -1.02 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (215.6 DPS) | yes | Signet Ring of the Bronze Dragonflight (234034, -4.05 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234030, -4.20 DPS) [vendor]; Wrath of Cenarius (21190, -5.43 DPS, sim-verified) [quest] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (215.6 DPS) | yes | Signet Ring of the Bronze Dragonflight (234034, -3.91 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234030, -4.05 DPS) [vendor]; Wrath of Cenarius (21190, -5.09 DPS, sim-verified) [quest] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (212.0 DPS) | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (215.6 DPS) | yes | Frozen Heart of the Mountain (249469, -2.91 DPS, sim-verified) [crafted]; Tidal Charm (1404, -4.33 DPS) [vendor]; Guardian Talisman (1490, -4.33 DPS) [quest] |
| main_hand | Annihilator (12798) | Blacksmithing [crafted] | sim-verified (218.6 DPS) | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [vendor]; High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Ebon Hand (19170, -1.31 DPS, sim-verified) [crafted] |
| off_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (218.6 DPS) | yes | High Warlord's Left Claw (18848, +0.00 DPS) [vendor]; High Warlord's Left Claw (234558, +0.00 DPS) [vendor]; Force Reactive Disk (18168, -22.69 DPS, sim-verified) [crafted] |
| ranged | Totem of the Storm (23199) (or Totem of Thunder (228176), Tidal Totem (272431), Totem of the Storm (272432), Burning Totem (272433), Totem of Urgency (279249), Totem of Ancestral Protection (249443), Kajaric Icon (206387), Polished Driftwood Icon (249398), Tempest Icon (206382), Dyadic Icon (206381), Galvanic Icon (206386), Sulfurous Icon (206388), Voltaic Icon (225838)) | World drop [world_drop] | 0.0 | yes | Totem of Thunder (228176, +0.00 DPS, sim-verified) [vendor]; Tidal Totem (272431, +0.00 DPS) [vendor]; Totem of the Storm (272432, +0.00 DPS) [world_drop] |

**New at 60:** head: Soulcrusher Headpiece; neck: Medallion of the Dawn; shoulder: Warlord's Mail Pauldrons; back: Howler's Furs; chest: Soulcrusher Armor; wrist: Soulcrusher Vambraces; hands: Soulcrusher Grips; waist: Soulcrusher Girdle; legs: Soulcrusher Leggings; feet: General's Mail Greaves; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Annihilator; off_hand: Shadowsong's Sorrow; ranged: Totem of the Storm

No-known-source sample (15 of 1242, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

