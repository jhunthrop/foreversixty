# Leveling BiS: Enhancement

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-254000000000000000-0000000000000000)

Set DPS (verified): 37.0. Weights run: 1.5s. Verify run: 1.4s. 220 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.358 ± 0.029, crit=0.918 ± 0.029 per rating point (14 rating = 1%, 12.853 per %), hit=1.472 ± 0.077 per rating point (10 rating = 1%, 14.724 per %), melee_haste=8.999 ± 0.401

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (0.57 DPS) | yes | Brawler's Leather Hood (252504, -0.29 DPS, sim-verified) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 2.7 attack_power points (0.10 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Slime-encrusted Pads (6461) | Wailing Caverns: Mutanus the Devourer [dungeon] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Serpent's Shoulders (5404, -0.39 DPS, sim-verified) [dungeon] |
| back | Grave Shroud (279865) | Abominable Creatures [quest] | 8.7 attack_power points (0.31 DPS) | yes | Glowing Lizardscale Cloak (6449, -0.02 DPS) [dungeon]; Dark Leather Cloak (2316, -0.02 DPS) [crafted]; Lambent Scale Cloak (4706, -0.03 DPS) [world_drop] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 19.5 attack_power points (0.69 DPS) | yes | Defender's Leather Armor (252434, -0.05 DPS) [crafted]; Tunic of Westfall (2041, -0.16 DPS) [quest]; Murloc Scale Breastplate (5781, -0.19 DPS) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 9.4 attack_power points (0.34 DPS) | yes | Bristlebark Bindings (14569, -0.05 DPS) [world_drop]; Forest Leather Bracers (3202, -0.09 DPS) [world_drop]; Wolf Bracers (4794, -0.14 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 16.1 attack_power points (0.57 DPS) | yes | Gold-flecked Gloves (5195, -0.08 DPS) [dungeon]; Bristlebark Gloves (14572, -0.10 DPS) [world_drop]; Brawler's Leather Gloves (252494, -0.10 DPS) [crafted] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.64 DPS) | yes | Deviate Scale Belt (6468, -0.19 DPS) [crafted]; Brawler's Leather Belt (252428, -0.23 DPS, sim-verified) [crafted]; Murloc Scale Belt (5780, -0.35 DPS) [crafted] |
| legs | Trapper's Leather Pants (252501) | Leatherworking [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Defender's Leather Pants (252445, -0.03 DPS) [crafted]; Leggings of the Fang (10410, -0.07 DPS) [dungeon]; Brawler's Leather Pants (252500, -0.40 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 16.9 attack_power points (0.60 DPS) | yes | Brawler's Leather Boots (252439, -0.00 DPS) [crafted]; Defender's Leather Boots (252441, -0.24 DPS) [crafted]; Totemic Leather Boots (252442, -0.24 DPS) [crafted] |
| finger1 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | The 1 Ring (8350, -0.17 DPS) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Ring of the Moon (12052, -0.03 DPS) [world_drop]; The 1 Ring (8350, -0.44 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 303.4 attack_power points (10.80 DPS) | yes | Night Reaver (1318, -0.53 DPS) [dungeon]; Staff of Westfall (2042, -0.58 DPS) [quest]; Living Root (6631, -0.65 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Erudite's Amulet; shoulder: Slime-encrusted Pads; back: Grave Shroud; chest: Brawler's Leather Armor; wrist: Bravo's Armbands; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Trapper's Leather Pants; feet: Feet of the Lynx; finger1: Demon Band; finger2: Lavishly Jeweled Ring; main_hand: Smite's Mighty Hammer

No-known-source sample (15 of 220, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18845 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (dwarf, 0000000000000000-255130030002000000-0000000000000000)

Set DPS (verified): 89.7. Weights run: 1.6s. Verify run: 1.6s. 420 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.006, agility=1.266 ± 0.030, crit=1.041 ± 0.034 per rating point (14 rating = 1%, 14.570 per %), hit=1.909 ± 0.087 per rating point (10 rating = 1%, 19.089 per %), melee_haste=8.928 ± 0.473

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (1.19 DPS) | yes | Cloudy Gustwoven Hood (277042, -0.30 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.30 DPS) [crafted]; Defender's Leather Hood (252447, -0.40 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 13.1 attack_power points (0.65 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.05 DPS) [world]; River Pride Choker (13087, -0.25 DPS) [world_drop]; Fallen Guard's Pendant (279837, -0.35 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.9 attack_power points (1.19 DPS) | yes | Bristlebark Amice (14573, -0.51 DPS) [world_drop]; Mantle of Thieves (2264, -0.56 DPS) [dungeon]; Barbaric Shoulders (5964, -0.78 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 14.9 attack_power points (0.74 DPS) | yes | Sergeant Major's Cape (16315, -0.09 DPS) [pvp]; Tigerstrike Mantle (13108, -0.24 DPS) [world_drop]; Wolfmaster Cape (6314, -0.24 DPS) [dungeon] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 22.1 attack_power points (1.10 DPS) | yes | Defender's Leather Tunic (252450, -0.13 DPS) [crafted]; Brawler's Leather Armor (252490, -0.16 DPS) [crafted]; Thick Murloc Armor (5782, -0.21 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 15.6 attack_power points (0.78 DPS) | yes | Barbaric Bracers (18948, -0.13 DPS) [crafted]; Bands of Serra'kis (6902, -0.18 DPS) [dungeon]; Demonhide Bracers (270033, -0.23 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 21.6 attack_power points (1.07 DPS) | yes | Toughened Leather Gloves (4253, -0.10 DPS) [crafted]; Wolfclaw Gloves (1978, -0.20 DPS) [dungeon]; Heavy Earthen Gloves (7359, -0.28 DPS) [crafted] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 25.6 attack_power points (1.27 DPS) | yes | Warden's Leather Belt (252460, -0.30 DPS) [crafted]; Blackened Defias Belt (10403, -0.38 DPS) [dungeon]; Skulker's Leather Belt (252520, -0.58 DPS, sim-verified) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.29 DPS) | yes | Brawler's Leather Pants (252500, -0.13 DPS) [crafted]; Trapper's Leather Pants (252501, -0.13 DPS) [crafted]; Brawler's Leather Legguards (252516, -1.13 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 16.3 attack_power points (0.81 DPS) | yes | Feet of the Lynx (1121, -0.01 DPS) [world_drop]; Disjointed Shoes (277226, -0.22 DPS) [quest]; Insignia Boots (4055, -0.31 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.8 attack_power points (0.99 DPS) | yes | Silverlaine's Family Seal (6321, -0.49 DPS) [dungeon]; Bonebinder's Signet (273032, -0.49 DPS) [world]; Seal of Wrynn (2933, -0.50 DPS) [quest] |
| finger2 | Tiger Band (6749) (or Bonebinder's Signet (273032), Silverlaine's Family Seal (6321)) | Willix the Importer [quest] | 10.0 attack_power points (0.50 DPS) | yes | Silverlaine's Family Seal (6321, +0.00 DPS) [dungeon]; Bonebinder's Signet (273032, +0.00 DPS) [world]; Seal of Wrynn (2933, -0.01 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (89.7 DPS) | yes | Corpsemaker (6687, -0.45 DPS) [dungeon]; Cobalt Crusher (7730, -1.97 DPS) [dungeon]; Viscous Hammer (13045, -21.58 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Ferine Leggings; feet: Brawler's Leather Boots; finger1: Thunderbrow Ring; finger2: Tiger Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 420, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 40 (dwarf, 0000000000000000-255130030005102031-0000000000000000)

Set DPS (verified): 116.8. Weights run: 2.0s. Verify run: 2.1s. 669 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.304 ± 0.042, crit=1.281 ± 0.050 per rating point (14 rating = 1%, 17.938 per %), hit=2.017 ± 0.102 per rating point (10 rating = 1%, 20.174 per %), melee_haste=7.027 ± 0.801

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 36.3 attack_power points (2.50 DPS) | yes | Barbaric Iron Helm (7915, -0.45 DPS) [crafted]; Hard Gold Coif (250537, -0.57 DPS) [crafted]; Tusken Helm (6686, -0.71 DPS) [dungeon] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | River Pride Choker (13087, -0.27 DPS) [world_drop]; Gazlowe's Charm (13088, -0.27 DPS) [dungeon]; Kaleidoscope Chain (13084, -1.19 DPS, sim-verified) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 24.3 attack_power points (1.67 DPS) | yes | Flintrock Shoulders (7755, -0.09 DPS) [dungeon]; Hard Gold Pauldrons (250539, -0.16 DPS) [crafted]; Barbaric Iron Shoulders (7913, -0.31 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 19.8 attack_power points (1.36 DPS) | yes | Yeti Fur Cloak (2805, -0.41 DPS) [quest]; Scorpashi Cape (14656, -0.64 DPS) [world_drop]; Hawkeye's Cloak (14593, -2.20 DPS, sim-verified) [world_drop] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 37.1 attack_power points (2.55 DPS) | yes | Icemail Jerkin (1981, -0.40 DPS) [world_drop]; Avenger's Armor (1488, -0.49 DPS) [dungeon]; Golden Scale Cuirass (3845, -0.63 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.37 DPS) | yes | Ravager's Armguards (14770, -0.05 DPS) [world_drop]; Pugilist Bracers (4438, -0.27 DPS) [dungeon]; Yorgen Bracers (13012, -0.28 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 37.9 attack_power points (2.61 DPS) | yes | Scarlet Gauntlets (10331, -0.06 DPS) [dungeon]; Gauntlets of Divinity (7724, -0.41 DPS) [dungeon]; Prowler's Leather Gloves (252524, -0.43 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 29.7 attack_power points (2.04 DPS) | yes | Officer's Belt (250556, -0.13 DPS) [crafted]; Prowler's Leather Belt (252459, -0.27 DPS) [crafted]; Scarlet Belt (10329, -0.39 DPS) [dungeon] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.88 DPS) | yes | Firemane Leggings (13129, -0.27 DPS) [world_drop]; Orcish War Leggings (7929, -0.55 DPS) [crafted]; Legguards of the Vault (9396, -0.65 DPS) [dungeon] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 31.1 attack_power points (2.14 DPS) | yes | Blackforge Greaves (6423, -0.01 DPS) [dungeon]; Skulker's Leather Shoes (252531, -0.19 DPS) [crafted]; Excelsior Boots (4109, -0.37 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.37 DPS) | yes | Thunderbrow Ring (13097, -0.01 DPS) [world_drop]; Falcon's Hook (7552, -0.16 DPS) [dungeon]; Ring of the Underwood (2951, -0.20 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.37 DPS) | yes | Thunderbrow Ring (13097, -0.01 DPS) [world_drop]; Falcon's Hook (7552, -0.16 DPS) [dungeon]; Ring of the Underwood (2951, -0.20 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ravager (7717) | Scarlet Monastery: Herod [dungeon] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Bonebiter (6830, +0.00 DPS) [quest]; Pendulum of Doom (9425, -1.47 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; neck: Archmagister's Faceted Pendant; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Prowler's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Ravager

No-known-source sample (15 of 669, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 50 (dwarf, 3200300000000000-255130030005102051-0000000000000000)

Set DPS (verified): 157.3. Weights run: 1.9s. Verify run: 3.7s. 872 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.233 ± 0.046, crit=1.450 ± 0.060 per rating point (14 rating = 1%, 20.302 per %), hit=2.732 ± 0.136 per rating point (10 rating = 1%, 27.324 per %), melee_haste=7.835 ± 1.063

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 35.6 attack_power points (2.58 DPS) | yes | Knight-Lieutenant's Mail Helmet (223075, +0.00 DPS) [vendor]; Barbaric Iron Helm (7915, -0.47 DPS) [crafted]; Hard Gold Coif (250537, -0.55 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 26.0 attack_power points (1.89 DPS) | yes | Kaleidoscope Chain (13084, -0.95 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.02 DPS) [world]; River Pride Choker (13087, -1.31 DPS) [world_drop] |
| shoulder | Prowler's Leather Shoulders (252534) | Leatherworking [crafted] | 33.1 attack_power points (2.40 DPS) | yes | Knight-Lieutenant's Mail Epaulets (223073, +0.00 DPS) [vendor]; Skulker's Leather Shoulder (252535, -0.11 DPS) [crafted]; Failed Flying Experiment (9647, -3.87 DPS, sim-verified) [quest] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 19.4 attack_power points (1.41 DPS) | yes | Bloodlust Cape (14801, -0.10 DPS) [world_drop]; Pridelord Cape (14673, -0.17 DPS) [world_drop]; Dark Phantom Cape (13122, -3.86 DPS, sim-verified) [world_drop] |
| chest | Knight's Mail Armor (223078) | Captain Dirgehammer [vendor] | 44.3 attack_power points (3.21 DPS) | yes | Kolkar Marauder Chain (6773, -0.56 DPS) [quest]; Blazewind Breastplate (11193, -0.72 DPS) [quest]; Warbear Harness (15064, -4.80 DPS, sim-verified) [crafted] |
| wrist | Arena Bands (18711) (or Bracers of the Stone Princess (17714)) | Arena Treasure Chest [world] | 28.0 attack_power points (2.03 DPS) | yes | Bracers of the Stone Princess (17714, +0.00 DPS) [dungeon]; Prowler's Leather Bracers (252539, -0.10 DPS) [crafted]; Deepfury Bracers (13120, -0.11 DPS) [world_drop] |
| hands | Stonerender Gauntlets (17007) | World drop [world_drop] | sim-verified (157.3 DPS) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, -0.07 DPS) [crafted]; Sergeant Major's Mail Gauntlets (223076, -1.88 DPS, sim-verified) [vendor] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 42.8 attack_power points (3.10 DPS) | yes | Skulker's Leather Waistguard (252474, -0.11 DPS) [crafted]; Belt of the Gladiator (13134, -0.49 DPS) [world_drop]; Ogron's Sash (13117, -0.99 DPS) [world_drop] |
| legs | Knight's Mail Legplates (223074) | Captain Dirgehammer [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpentskin Leggings (8262, +0.00 DPS) [world_drop]; Scarlet Leggings (10330, -0.17 DPS) [dungeon]; Gryphon Rider's Leggings (9652, -4.36 DPS, sim-verified) [quest] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 39.6 attack_power points (2.87 DPS) | yes | Skulker's Leather Boots (252469, -0.11 DPS) [crafted]; Prowler's Leather Shoes (252465, -0.65 DPS) [crafted]; Blackforge Greaves (6423, -0.67 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) (or Mark of Kern (2262), Assault Band (13095)) | Maraudon: Princess Theradras [dungeon] | 20.0 attack_power points (1.45 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.02 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.20 DPS) [quest] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.45 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.02 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.20 DPS) [quest] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thorium Greatmace (250613, -0.94 DPS) [crafted]; Pendulum of Doom (9425, -1.75 DPS, sim-verified) [dungeon]; Bleakwood Hew (12769, -3.29 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** neck: Skibi's Pendant; shoulder: Prowler's Leather Shoulders; chest: Knight's Mail Armor; wrist: Arena Bands; hands: Stonerender Gauntlets; waist: Prowler's Leather Waistguard; legs: Knight's Mail Legplates; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: Mark of Kern; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 872, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (dwarf, 3200330000000000-255130030005102051-0520000000000000)

Set DPS (verified): 224.7. Weights run: 1.8s. Verify run: 6.0s. 1819 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.767 ± 0.068, crit=2.477 ± 0.096 per rating point (14 rating = 1%, 34.680 per %), hit=3.912 ± 0.195 per rating point (10 rating = 1%, 39.115 per %), melee_haste=17.063 ± 1.483

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (224.7 DPS) | yes | Blue Suede Hat (252482, -0.71 DPS) [crafted]; Warbear Helm (252485, -0.83 DPS) [crafted]; Outlaw's Collar (279253, -9.47 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 33.0 attack_power points (2.47 DPS) | yes | Sunprism Pendant (20649, -0.88 DPS) [quest]; Kaleidoscope Chain (13084, -1.34 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.57 DPS) [world] |
| shoulder | Dark Warder's Pauldrons (22241) | Blackrock Depths: Warder Stilgiss [dungeon] | 52.0 attack_power points (3.90 DPS) | yes | Knight-Lieutenant's Mail Epaulets (223073, +0.00 DPS) [vendor]; Shadowcraft Spaulders (16708, -0.40 DPS) [dungeon]; Golden Mantle of the Dawn (19058, -0.50 DPS) [crafted] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 46.5 attack_power points (3.49 DPS) | yes | Fel Cape (279269, -0.89 DPS) [crafted]; Sergeant Major's Cape (16337, -1.23 DPS) [pvp]; Dark Phantom Cape (13122, -1.50 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 139.2 attack_power points (10.44 DPS) | yes | Timbermaw Tunic (252484, -0.50 DPS) [crafted]; Shadowcraft Tunic (16721, -4.65 DPS) [dungeon]; Stormshroud Armor (15056, -5.24 DPS) [crafted] |
| wrist | Arena Wristguards (18709) | Arena Treasure Chest [world] | 34.7 attack_power points (2.60 DPS) | yes | Deepfury Bracers (13120, -0.01 DPS) [world_drop]; Tranquil Wristguards (279256, -0.20 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.32 DPS) [crafted] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-verified (224.7 DPS) | yes | Gloves of Holy Might (867, -0.60 DPS) [world_drop]; Sergeant Major's Mail Gauntlets (223076, -0.90 DPS) [vendor]; Timbermaw Brawlers (19049, -8.72 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 66.7 attack_power points (5.01 DPS) | yes | Skulker's Leather Waistguard (252474, -1.35 DPS) [crafted]; Serpentine Sash (13118, -1.80 DPS) [world_drop]; Prowler's Leather Waistguard (252473, -2.08 DPS, sim-verified) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (224.7 DPS) | yes | Pristine Scorpid Leggings (252606, -0.75 DPS) [crafted]; Stormshroud Pants (15057, -0.85 DPS) [crafted]; Warbear Woolies (15065, -8.88 DPS, sim-verified) [crafted] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 52.1 attack_power points (3.91 DPS) | yes | Prowler's Leather Boots (252468, -0.50 DPS) [crafted]; Skulker's Leather Boots (252469, -0.54 DPS) [crafted]; Shadowcraft Boots (16711, -1.66 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 24.7 attack_power points (1.86 DPS) | yes | Ring of the Underwood (2951, -0.23 DPS) [world_drop]; Thunderbrow Ring (13097, -0.26 DPS) [world_drop]; Ring of Living Stone (18400, -0.36 DPS) [quest] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 21.9 attack_power points (1.64 DPS) | yes | Ring of the Underwood (2951, -0.02 DPS) [world_drop]; Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Ring of Living Stone (18400, -0.14 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (224.7 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (224.7 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Forest Defender's Axe (279260) | Blacksmithing [crafted] | sim-verified (224.7 DPS) | yes | Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Fiery War Axe (870, -32.50 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Black Dragonscale Helm; shoulder: Dark Warder's Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Arena Wristguards; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Scalegut Treaders; finger1: Masons Fraternity Ring; finger2: Falcon's Hook; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Forest Defender's Axe

No-known-source sample (15 of 1819, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60, raid preset (dwarf, 0533001000000000-253030031005102051-0530000000000000)

Set DPS (verified): 623.5. Weights run: 1.8s. Verify run: 5.9s. 1819 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.005, agility=2.277 ± 0.085, crit=2.942 ± 0.114 per rating point (14 rating = 1%, 41.192 per %), hit=5.248 ± 0.296 per rating point (10 rating = 1%, 52.475 per %), melee_haste=23.353 ± 2.569

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (623.5 DPS) | yes | Blue Suede Hat (252482, -2.20 DPS) [crafted]; Warbear Helm (252485, -2.54 DPS) [crafted]; Outlaw's Collar (279253, -22.66 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 40.6 attack_power points (5.88 DPS) | yes | Sunprism Pendant (20649, -1.92 DPS) [quest]; Kaleidoscope Chain (13084, -3.29 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -4.14 DPS) [world] |
| shoulder | Dark Warder's Pauldrons (22241) | Blackrock Depths: Warder Stilgiss [dungeon] | 62.9 attack_power points (9.12 DPS) | yes | Knight-Lieutenant's Mail Epaulets (223073, -0.28 DPS) [vendor]; Shadowcraft Spaulders (16708, -0.34 DPS) [dungeon]; Clouddrift Mantle (11874, -1.33 DPS) [quest] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 54.2 attack_power points (7.85 DPS) | yes | Fel Cape (279269, -1.88 DPS) [crafted]; Sergeant Major's Cape (16337, -2.66 DPS) [pvp]; Dark Phantom Cape (13122, -2.90 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 177.9 attack_power points (25.78 DPS) | yes | Timbermaw Tunic (252484, -3.35 DPS) [crafted]; Shadowcraft Tunic (16721, -11.12 DPS) [dungeon]; Stormshroud Armor (15056, -13.84 DPS) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 43.0 attack_power points (6.22 DPS) | yes | Arena Wristguards (18709, -0.26 DPS) [world]; Skulker's Leather Bracers (252540, -1.02 DPS) [crafted]; Prowler's Leather Bracers (252539, -1.05 DPS) [crafted] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-verified (623.5 DPS) | yes | Gloves of Holy Might (867, -1.16 DPS) [world_drop]; Sergeant Major's Mail Gauntlets (223076, -1.51 DPS) [vendor]; Timbermaw Brawlers (19049, -16.31 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 78.1 attack_power points (11.31 DPS) | yes | Skulker's Leather Waistguard (252474, -2.87 DPS) [crafted]; Prowler's Leather Waistguard (252473, -2.89 DPS) [crafted]; Serpentine Sash (13118, -3.83 DPS) [world_drop] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (623.5 DPS) | yes | Stormshroud Pants (15057, -0.70 DPS) [crafted]; Pristine Scorpid Leggings (252606, -0.93 DPS) [crafted]; Warbear Woolies (15065, -15.59 DPS, sim-verified) [crafted] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 63.6 attack_power points (9.21 DPS) | yes | Scalegut Treaders (275618, -1.06 DPS) [crafted]; Savage Gladiator Greaves (11731, -1.39 DPS) [dungeon]; Skulker's Leather Boots (252469, -1.41 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 31.9 attack_power points (4.62 DPS) | yes | Falcon's Hook (7552, -0.69 DPS) [dungeon]; Thunderbrow Ring (13097, -1.08 DPS) [world_drop]; Ring of Living Stone (18400, -1.43 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 27.2 attack_power points (3.94 DPS) | yes | Falcon's Hook (7552, -0.01 DPS) [dungeon]; Thunderbrow Ring (13097, -0.40 DPS) [world_drop]; Ring of Living Stone (18400, -0.75 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (623.5 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (623.5 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Forest Defender's Axe (279260) | Blacksmithing [crafted] | sim-verified (623.5 DPS) | yes | Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Fiery War Axe (870, -68.62 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Black Dragonscale Helm; shoulder: Dark Warder's Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Deepfury Bracers; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Shadowcraft Boots; finger1: Masons Fraternity Ring; finger2: Ring of the Underwood; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Forest Defender's Axe

No-known-source sample (15 of 1819, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

## Horde

### Band 20 (orc, 0000000000000000-254000000000000000-0000000000000000)

Set DPS (verified): 37.5. Weights run: 1.5s. Verify run: 1.6s. 209 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.358 ± 0.029, crit=0.918 ± 0.029 per rating point (14 rating = 1%, 12.853 per %), hit=1.472 ± 0.077 per rating point (10 rating = 1%, 14.724 per %), melee_haste=8.999 ± 0.401

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (0.57 DPS) | yes | Brawler's Leather Hood (252504, -0.27 DPS, sim-verified) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 2.7 attack_power points (0.10 DPS) | yes | Roadwatcher's Confidence (281265, +0.00 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 6.8 attack_power points (0.24 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS, sim-verified) [dungeon] |
| back | Grave Shroud (279865) | Unending Torment [quest] | 8.7 attack_power points (0.31 DPS) | yes | Glowing Lizardscale Cloak (6449, -0.02 DPS) [dungeon]; Dark Leather Cloak (2316, -0.02 DPS) [crafted]; Lambent Scale Cloak (4706, -0.03 DPS) [world_drop] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 19.5 attack_power points (0.69 DPS) | yes | Defender's Leather Armor (252434, -0.05 DPS) [crafted]; Murloc Scale Breastplate (5781, -0.19 DPS) [crafted]; Totemic Leather Armor (252435, -0.20 DPS) [crafted] |
| wrist | Bristlebark Bindings (14569) | World drop [world_drop] | 8.1 attack_power points (0.29 DPS) | yes | Forest Leather Bracers (3202, -0.05 DPS) [world_drop]; Wolf Bracers (4794, -0.09 DPS) [vendor]; Spare Part Bindings (279875, -0.14 DPS) [quest] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 16.1 attack_power points (0.57 DPS) | yes | Gold-flecked Gloves (5195, -0.08 DPS) [dungeon]; Bristlebark Gloves (14572, -0.10 DPS) [world_drop]; Brawler's Leather Gloves (252494, -0.10 DPS) [crafted] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.64 DPS) | yes | Deviate Scale Belt (6468, -0.19 DPS) [crafted]; Brawler's Leather Belt (252428, -0.21 DPS, sim-verified) [crafted]; Murloc Scale Belt (5780, -0.35 DPS) [crafted] |
| legs | Trapper's Leather Pants (252501) | Leatherworking [crafted] | sim-verified (37.5 DPS) | yes | Defender's Leather Pants (252445, -0.03 DPS) [crafted]; Leggings of the Fang (10410, -0.07 DPS) [dungeon]; Brawler's Leather Pants (252500, -0.54 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 16.9 attack_power points (0.60 DPS) | yes | Brawler's Leather Boots (252439, -0.00 DPS) [crafted]; Defender's Leather Boots (252441, -0.24 DPS) [crafted]; Totemic Leather Boots (252442, -0.24 DPS) [crafted] |
| finger1 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Bounty Hunter's Ring (5351, -0.14 DPS) [quest]; The 1 Ring (8350, -0.17 DPS) [world]; Lavishly Jeweled Ring (1156, -0.19 DPS) [dungeon] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 6.0 attack_power points (0.21 DPS) | yes | The 1 Ring (8350, -0.09 DPS) [world]; Lavishly Jeweled Ring (1156, -0.12 DPS) [dungeon]; Bounty Hunter's Ring (5351, -0.23 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hammerbone (270018) | Leaders of the Fang [quest] | 306.2 attack_power points (10.90 DPS) | yes | Smite's Mighty Hammer (7230, -0.10 DPS) [dungeon]; Forsaken Greataxe (251533, -0.11 DPS) [quest]; Living Root (6631, -0.33 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Grave Shroud; chest: Brawler's Leather Armor; wrist: Bristlebark Bindings; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Trapper's Leather Pants; feet: Feet of the Lynx; finger1: Demon Band; finger2: Loop of Sacrifice; main_hand: Hammerbone

No-known-source sample (15 of 209, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion; 21566 Rune of Perfection; 21568 Rune of Duty

### Band 30 (orc, 0000000000000000-255130030002000000-0000000000000000)

Set DPS (verified): 90.0. Weights run: 1.6s. Verify run: 1.5s. 413 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.006, agility=1.266 ± 0.030, crit=1.041 ± 0.034 per rating point (14 rating = 1%, 14.570 per %), hit=1.909 ± 0.087 per rating point (10 rating = 1%, 19.089 per %), melee_haste=8.928 ± 0.473

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (1.19 DPS) | yes | Cloudy Gustwoven Hood (277042, -0.30 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.30 DPS) [crafted]; Defender's Leather Hood (252447, -0.40 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 13.1 attack_power points (0.65 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.05 DPS) [world]; River Pride Choker (13087, -0.25 DPS) [world_drop]; Erudite's Amulet (277204, -0.52 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.9 attack_power points (1.19 DPS) | yes | Bristlebark Amice (14573, -0.51 DPS) [world_drop]; Mantle of Thieves (2264, -0.56 DPS) [dungeon]; Barbaric Shoulders (5964, -0.63 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 14.9 attack_power points (0.74 DPS) | yes | Tigerstrike Mantle (13108, -0.24 DPS) [world_drop]; Wildhunter Cloak (16658, -0.24 DPS) [quest]; Construct Cloak (279848, -0.24 DPS) [quest] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 22.1 attack_power points (1.10 DPS) | yes | Dreamer's Chestguard (270043, +0.00 DPS, sim-verified) [quest]; Defender's Leather Tunic (252450, -0.13 DPS) [crafted]; Brawler's Leather Armor (252490, -0.16 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 15.6 attack_power points (0.78 DPS) | yes | Barbaric Bracers (18948, -0.13 DPS) [crafted]; Bands of Serra'kis (6902, -0.18 DPS) [dungeon]; Cultist's Armguards (270032, -0.28 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 21.6 attack_power points (1.07 DPS) | yes | Toughened Leather Gloves (4253, -0.10 DPS) [crafted]; Wolfclaw Gloves (1978, -0.20 DPS) [dungeon]; Heavy Earthen Gloves (7359, -0.28 DPS) [crafted] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 25.6 attack_power points (1.27 DPS) | yes | Warden's Leather Belt (252460, -0.30 DPS) [crafted]; Blackened Defias Belt (10403, -0.38 DPS) [dungeon]; Skulker's Leather Belt (252520, -0.85 DPS, sim-verified) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.29 DPS) | yes | Brawler's Leather Pants (252500, -0.13 DPS) [crafted]; Trapper's Leather Pants (252501, -0.13 DPS) [crafted]; Brawler's Leather Legguards (252516, -1.74 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 16.3 attack_power points (0.81 DPS) | yes | Feet of the Lynx (1121, -0.01 DPS) [world_drop]; Stomping Boots (3741, -0.16 DPS) [quest]; Insignia Boots (4055, -0.31 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.8 attack_power points (0.99 DPS) | yes | Silverlaine's Family Seal (6321, -0.49 DPS) [dungeon]; Tiger Band (6749, -0.49 DPS) [quest]; Bonebinder's Signet (273032, -0.49 DPS) [world] |
| finger2 | Band of the Fist (17694) | Allegiance to the Old Gods [quest] | 13.1 attack_power points (0.65 DPS) | yes | Silverlaine's Family Seal (6321, -0.15 DPS) [dungeon]; Tiger Band (6749, -0.15 DPS) [quest]; Bonebinder's Signet (273032, -0.15 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (90.0 DPS) | yes | Corpsemaker (6687, -0.45 DPS) [dungeon]; Cobalt Crusher (7730, -1.97 DPS) [dungeon]; Viscous Hammer (13045, -21.78 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Ferine Leggings; feet: Brawler's Leather Boots; finger1: Thunderbrow Ring; finger2: Band of the Fist; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 413, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (orc, 0000000000000000-255130030005102031-0000000000000000)

Set DPS (verified): 114.4. Weights run: 2.0s. Verify run: 1.7s. 651 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.304 ± 0.042, crit=1.281 ± 0.050 per rating point (14 rating = 1%, 17.938 per %), hit=2.017 ± 0.102 per rating point (10 rating = 1%, 20.174 per %), melee_haste=7.027 ± 0.801

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 36.3 attack_power points (2.50 DPS) | yes | Barbaric Iron Helm (7915, -0.45 DPS) [crafted]; Hard Gold Coif (250537, -0.57 DPS) [crafted]; Tusken Helm (6686, -0.71 DPS) [dungeon] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | sim-verified (114.4 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.08 DPS) [world]; River Pride Choker (13087, -0.36 DPS) [world_drop]; Ethereal Talisman (4430, -1.45 DPS, sim-verified) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 24.3 attack_power points (1.67 DPS) | yes | Flintrock Shoulders (7755, -0.09 DPS) [dungeon]; Hard Gold Pauldrons (250539, -0.16 DPS) [crafted]; Barbaric Iron Shoulders (7913, -0.31 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 15.1 attack_power points (1.04 DPS) | yes | Scorpashi Cape (14656, -0.32 DPS) [world_drop]; Imperial Cloak (6432, -0.32 DPS) [dungeon]; Parachute Cloak (10518, -0.32 DPS) [crafted] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 37.1 attack_power points (2.55 DPS) | yes | Icemail Jerkin (1981, -0.40 DPS) [world_drop]; Avenger's Armor (1488, -0.49 DPS) [dungeon]; Golden Scale Cuirass (3845, -0.63 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.37 DPS) | yes | Ravager's Armguards (14770, -0.05 DPS) [world_drop]; Pugilist Bracers (4438, -0.27 DPS) [dungeon]; Yorgen Bracers (13012, -0.28 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 37.9 attack_power points (2.61 DPS) | yes | Scarlet Gauntlets (10331, -0.06 DPS) [dungeon]; Gauntlets of Divinity (7724, -0.41 DPS) [dungeon]; Prowler's Leather Gloves (252524, -0.43 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 29.7 attack_power points (2.04 DPS) | yes | Officer's Belt (250556, -0.13 DPS) [crafted]; Tharg's Shoelace (9705, -0.26 DPS) [quest]; Prowler's Leather Belt (252459, -0.27 DPS) [crafted] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.88 DPS) | yes | Firemane Leggings (13129, -0.27 DPS) [world_drop]; Orcish War Leggings (7929, -0.55 DPS) [crafted]; Legguards of the Vault (9396, -0.65 DPS) [dungeon] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 31.1 attack_power points (2.14 DPS) | yes | Skulker's Leather Shoes (252531, -0.19 DPS) [crafted]; Excelsior Boots (4109, -0.37 DPS) [quest]; Blackforge Greaves (6423, -2.20 DPS, sim-verified) [dungeon] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.37 DPS) | yes | Thunderbrow Ring (13097, -0.01 DPS) [world_drop]; Falcon's Hook (7552, -0.16 DPS) [dungeon]; Ring of the Underwood (2951, -0.20 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.37 DPS) | yes | Thunderbrow Ring (13097, -0.01 DPS) [world_drop]; Falcon's Hook (7552, -0.16 DPS) [dungeon]; Ring of the Underwood (2951, -0.20 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ravager (7717) | Scarlet Monastery: Herod [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fiery War Axe (870, +0.00 DPS) [world_drop]; Staff of Jordan (873, +0.00 DPS) [world_drop]; Pendulum of Doom (9425, +0.00 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Prowler's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Ravager

No-known-source sample (15 of 651, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7748 Forcestone Buckler; 8708 Hammer of Expertise

### Band 50 (orc, 3200300000000000-255130030005102051-0000000000000000)

Set DPS (verified): 158.2. Weights run: 1.9s. Verify run: 2.2s. 842 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.233 ± 0.046, crit=1.450 ± 0.060 per rating point (14 rating = 1%, 20.302 per %), hit=2.732 ± 0.136 per rating point (10 rating = 1%, 27.324 per %), melee_haste=7.835 ± 1.063

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 35.6 attack_power points (2.58 DPS) | yes | Blood Guard's Mail Helmet (220820, +0.00 DPS) [vendor]; Undercity Reservist's Cap (20643, -0.14 DPS) [quest]; Barbaric Iron Helm (7915, -0.47 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 26.0 attack_power points (1.89 DPS) | yes | Woven Ivy Necklace (19159, -0.21 DPS) [quest]; Ethereal Talisman (4430, -0.80 DPS) [quest]; Kaleidoscope Chain (13084, -0.95 DPS) [world_drop] |
| shoulder | Prowler's Leather Shoulders (252534) | Leatherworking [crafted] | 33.1 attack_power points (2.40 DPS) | yes | Blood Guard's Mail Epaulets (220823, +0.00 DPS) [vendor]; Skulker's Leather Shoulder (252535, -0.11 DPS) [crafted]; Failed Flying Experiment (9647, -1.65 DPS, sim-verified) [quest] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 18.5 attack_power points (1.34 DPS) | yes | Bloodlust Cape (14801, -0.04 DPS) [world_drop]; Pridelord Cape (14673, -0.10 DPS) [world_drop]; Blackflame Cape (13109, -0.27 DPS) [world_drop] |
| chest | Stone Guard's Mail Armor (220826) | PvP rank 10 · Stone Guard · Horde [vendor] | 44.3 attack_power points (3.21 DPS) | yes | Kolkar Marauder Chain (6773, -0.56 DPS) [quest]; Blazewind Breastplate (11193, -0.72 DPS) [quest]; Warbear Harness (15064, -2.20 DPS, sim-verified) [crafted] |
| wrist | Arena Bands (18711) (or Bracers of the Stone Princess (17714)) | Arena Treasure Chest [world] | 28.0 attack_power points (2.03 DPS) | yes | Bracers of the Stone Princess (17714, +0.00 DPS, sim-verified) [dungeon]; Prowler's Leather Bracers (252539, -0.10 DPS) [crafted]; Deepfury Bracers (13120, -0.11 DPS) [world_drop] |
| hands | Stonerender Gauntlets (17007) | World drop [world_drop] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Prowler's Leather Gauntlets (252547, -0.07 DPS) [crafted]; Scarlet Gauntlets (10331, -0.27 DPS) [dungeon]; Gloves of Holy Might (867, -1.98 DPS, sim-verified) [world_drop] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 42.8 attack_power points (3.10 DPS) | yes | Skulker's Leather Waistguard (252474, -0.11 DPS) [crafted]; Belt of the Gladiator (13134, -0.49 DPS) [world_drop]; Ogron's Sash (13117, -0.99 DPS) [world_drop] |
| legs | Stone Guard's Mail Legplates (220834) | PvP rank 10 · Stone Guard · Horde [vendor] | sim-verified (+3.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Scarlet Leggings (10330, -0.17 DPS) [dungeon]; Stormshroud Pants (15057, -0.27 DPS) [crafted]; Serpentskin Leggings (8262, -3.57 DPS, sim-verified) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 39.6 attack_power points (2.87 DPS) | yes | Skulker's Leather Boots (252469, -0.11 DPS) [crafted]; Prowler's Leather Shoes (252465, -0.65 DPS) [crafted]; Blackforge Greaves (6423, -0.67 DPS) [dungeon] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.74 DPS) | yes | Mark of Kern (2262, -0.29 DPS) [dungeon]; Assault Band (13095, -0.29 DPS) [world_drop]; Thunderbrow Ring (13097, -0.31 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) (or Mark of Kern (2262), Assault Band (13095)) | Maraudon: Princess Theradras [dungeon] | 20.0 attack_power points (1.45 DPS) | yes | Mark of Kern (2262, +0.00 DPS) [dungeon]; Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.02 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -4.39 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -1.49 DPS, sim-verified) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pendulum of Doom (9425, +0.00 DPS) [dungeon]; Thorium Greatmace (250613, -0.94 DPS) [crafted]; Bleakwood Hew (12769, -3.29 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** neck: Skibi's Pendant; shoulder: Prowler's Leather Shoulders; back: Dark Phantom Cape; chest: Stone Guard's Mail Armor; wrist: Arena Bands; hands: Stonerender Gauntlets; waist: Prowler's Leather Waistguard; legs: Stone Guard's Mail Legplates; feet: Prowler's Leather Boots; finger1: White Bone Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 842, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60 (orc, 3200330000000000-255130030005102051-0520000000000000)

Set DPS (verified): 229.9. Weights run: 1.8s. Verify run: 5.7s. 1775 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.767 ± 0.068, crit=2.477 ± 0.096 per rating point (14 rating = 1%, 34.680 per %), hit=3.912 ± 0.195 per rating point (10 rating = 1%, 39.115 per %), melee_haste=17.063 ± 1.483

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (229.9 DPS) | yes | Champion's Mail Headguard (227155, +0.00 DPS) [pvp]; Warlord's Mail Headguard (231657, -0.68 DPS) [vendor]; Outlaw's Collar (279253, -8.48 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 33.0 attack_power points (2.47 DPS) | yes | Woven Ivy Necklace (19159, -0.38 DPS) [quest]; Sunprism Pendant (20649, -0.88 DPS) [quest]; Ethereal Talisman (4430, -1.19 DPS) [quest] |
| shoulder | Dark Warder's Pauldrons (22241) | Blackrock Depths: Warder Stilgiss [dungeon] | 52.0 attack_power points (3.90 DPS) | yes | Shadowcraft Spaulders (16708, +0.00 DPS) [dungeon]; Champion's Mail Pauldrons (227154, +0.00 DPS) [pvp]; Warlord's Mail Pauldrons (231654, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 46.5 attack_power points (3.49 DPS) | yes | Fel Cape (279269, -0.89 DPS) [crafted]; Dark Phantom Cape (13122, -1.50 DPS) [world_drop]; Pridelord Cape (14673, -1.85 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 139.2 attack_power points (10.44 DPS) | yes | Timbermaw Tunic (252484, -0.50 DPS) [crafted]; Legionnaire's Mail Hauberk (227157, -1.60 DPS) [pvp]; Warlord's Mail Hauberk (231653, -1.60 DPS) [vendor] |
| wrist | Arena Wristguards (18709) | Arena Treasure Chest [world] | 34.7 attack_power points (2.60 DPS) | yes | Deepfury Bracers (13120, -0.01 DPS) [world_drop]; Tranquil Wristguards (279256, -0.20 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.32 DPS) [crafted] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-verified (229.9 DPS) | yes | Blood Guard's Mail Vices (227159, +0.00 DPS) [pvp]; General's Mail Vices (231655, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -4.52 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 66.7 attack_power points (5.01 DPS) | yes | Prowler's Leather Waistguard (252473, -1.32 DPS) [crafted]; Skulker's Leather Waistguard (252474, -1.35 DPS) [crafted]; Serpentine Sash (13118, -1.80 DPS) [world_drop] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (229.9 DPS) | yes | Legionnaire's Mail Legguards (227156, +0.00 DPS) [pvp]; General's Mail Legguards (231658, +0.00 DPS) [vendor]; Warbear Woolies (15065, -4.78 DPS, sim-verified) [crafted] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 52.1 attack_power points (3.91 DPS) | yes | Blood Guard's Mail Greaves (227158, +0.00 DPS) [pvp]; General's Mail Greaves (231656, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -0.25 DPS) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 24.7 attack_power points (1.86 DPS) | yes | Falcon's Hook (7552, -0.21 DPS) [dungeon]; Ring of the Underwood (2951, -0.23 DPS) [world_drop]; Thunderbrow Ring (13097, -0.26 DPS) [world_drop] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.80 DPS) | yes | Falcon's Hook (7552, -0.16 DPS) [dungeon]; Ring of the Underwood (2951, -0.18 DPS) [world_drop]; Thunderbrow Ring (13097, -0.20 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (229.9 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (229.9 DPS) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Forest Defender's Axe (279260) | Blacksmithing [crafted] | sim-verified (229.9 DPS) | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; Fiery War Axe (870, -32.63 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Black Dragonscale Helm; shoulder: Dark Warder's Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Arena Wristguards; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Scalegut Treaders; finger1: Masons Fraternity Ring; finger2: White Bone Band; main_hand: Forest Defender's Axe

No-known-source sample (15 of 1775, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60, raid preset (orc, 0533001000000000-253030031005102051-0530000000000000)

Set DPS (verified): 617.0. Weights run: 1.8s. Verify run: 4.5s. 1775 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.005, agility=2.277 ± 0.085, crit=2.942 ± 0.114 per rating point (14 rating = 1%, 41.192 per %), hit=5.248 ± 0.296 per rating point (10 rating = 1%, 52.475 per %), melee_haste=23.353 ± 2.569

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (617.0 DPS) | yes | Champion's Mail Headguard (227155, +0.00 DPS) [pvp]; Blue Suede Hat (252482, -2.20 DPS) [crafted]; Outlaw's Collar (279253, -19.75 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 40.6 attack_power points (5.88 DPS) | yes | Woven Ivy Necklace (19159, -1.00 DPS) [quest]; Sunprism Pendant (20649, -1.92 DPS) [quest]; Ethereal Talisman (4430, -2.97 DPS) [quest] |
| shoulder | Dark Warder's Pauldrons (22241) | Blackrock Depths: Warder Stilgiss [dungeon] | 62.9 attack_power points (9.12 DPS) | yes | Shadowcraft Spaulders (16708, +0.00 DPS) [dungeon]; Champion's Mail Pauldrons (227154, +0.00 DPS) [pvp]; Warlord's Mail Pauldrons (231654, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 54.2 attack_power points (7.85 DPS) | yes | Fel Cape (279269, -1.88 DPS) [crafted]; Dark Phantom Cape (13122, -2.90 DPS) [world_drop]; Blackflame Cape (13109, -3.89 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 177.9 attack_power points (25.78 DPS) | yes | Timbermaw Tunic (252484, -3.35 DPS) [crafted]; Legionnaire's Mail Hauberk (227157, -5.20 DPS) [pvp]; Warlord's Mail Hauberk (231653, -5.20 DPS) [vendor] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 43.0 attack_power points (6.22 DPS) | yes | Arena Wristguards (18709, -0.26 DPS) [world]; Skulker's Leather Bracers (252540, -1.02 DPS) [crafted]; Prowler's Leather Bracers (252539, -1.05 DPS) [crafted] |
| hands | Timbermaw Brawlers (19049) | Leatherworking [crafted] | 78.1 attack_power points (11.31 DPS) | yes | Blood Guard's Mail Vices (227159, +0.00 DPS) [pvp]; General's Mail Vices (231655, +0.00 DPS) [vendor]; Devilsaur Gauntlets (15063, -1.29 DPS) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 78.1 attack_power points (11.31 DPS) | yes | Skulker's Leather Waistguard (252474, -2.87 DPS) [crafted]; Prowler's Leather Waistguard (252473, -2.89 DPS) [crafted]; Serpentine Sash (13118, -3.83 DPS) [world_drop] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 102.6 attack_power points (14.87 DPS) | yes | Legionnaire's Mail Legguards (227156, +0.00 DPS) [pvp]; General's Mail Legguards (231658, +0.00 DPS) [vendor]; Devilsaur Leggings (15062, -2.23 DPS) [crafted] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 63.6 attack_power points (9.21 DPS) | yes | Blood Guard's Mail Greaves (227158, +0.00 DPS) [pvp]; General's Mail Greaves (231656, +0.00 DPS) [vendor]; Scalegut Treaders (275618, -1.06 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 31.9 attack_power points (4.62 DPS) | yes | Falcon's Hook (7552, -0.69 DPS) [dungeon]; Thunderbrow Ring (13097, -1.08 DPS) [world_drop]; White Bone Band (11862, -1.14 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 27.2 attack_power points (3.94 DPS) | yes | Falcon's Hook (7552, -0.01 DPS) [dungeon]; Thunderbrow Ring (13097, -0.40 DPS) [world_drop]; White Bone Band (11862, -0.46 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -7.12 DPS, sim-verified) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Forest Defender's Axe (279260) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; Fiery War Axe (870, -64.03 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Black Dragonscale Helm; shoulder: Dark Warder's Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Deepfury Bracers; hands: Timbermaw Brawlers; waist: Might of the Timbermaw; legs: Warbear Woolies; feet: Shadowcraft Boots; finger1: Masons Fraternity Ring; finger2: Ring of the Underwood; trinket2: Molten Heart of the Mountain; main_hand: Forest Defender's Axe

No-known-source sample (15 of 1775, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

