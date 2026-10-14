# Leveling BiS: Enhancement

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-254000000000000000-0000000000000000)

Set DPS (verified): 37.5. Weights run: 1.5s. Verify run: 4.6s. 217 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.006, agility=1.355 ± 0.041, crit=0.913 ± 0.041 per rating point (14 rating = 1%, 12.786 per %), hit=1.669 ± 0.111 per rating point (10 rating = 1%, 16.690 per %), melee_haste=9.111 ± 0.568

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (0.57 DPS) | yes | Brawler's Leather Hood (252504, -0.31 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 8.1 attack_power points (0.29 DPS) | yes | Erudite's Amulet (277204, -0.19 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 6.8 attack_power points (0.24 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS) [dungeon] |
| back | Grave Shroud (279865) | Abominable Creatures [quest] | 8.7 attack_power points (0.31 DPS) | yes | Glowing Lizardscale Cloak (6449, -0.02 DPS) [dungeon]; Dark Leather Cloak (2316, -0.02 DPS) [crafted]; Lambent Scale Cloak (4706, -0.03 DPS) [world_drop] |
| chest | Armor of the Fang (6473) | Wailing Caverns: Lord Pythas [dungeon] | sim-verified (37.5 DPS) | yes | Tunic of Westfall (2041, +0.00 DPS) [quest]; Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -0.34 DPS, sim-verified) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 9.4 attack_power points (0.34 DPS) | yes | Bristlebark Bindings (14569, -0.05 DPS) [world_drop]; Forest Leather Bracers (3202, -0.09 DPS) [world_drop]; Wolf Bracers (4794, -0.14 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 16.1 attack_power points (0.57 DPS) | yes | Bristlebark Gloves (14572, -0.10 DPS) [world_drop]; Brawler's Leather Gloves (252494, -0.10 DPS) [crafted]; Gold-flecked Gloves (5195, -0.38 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.64 DPS) | yes | Deviate Scale Belt (6468, -0.19 DPS) [crafted]; Brawler's Leather Belt (252428, -0.24 DPS, sim-verified) [crafted]; Murloc Scale Belt (5780, -0.35 DPS) [crafted] |
| legs | Leggings of the Fang (10410) | Wailing Caverns: Lord Cobrahn [dungeon] | sim-verified (37.5 DPS) | yes | Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 16.8 attack_power points (0.60 DPS) | yes | Brawler's Leather Boots (252439, -0.00 DPS) [crafted]; Defender's Leather Boots (252441, -0.24 DPS) [crafted]; Totemic Leather Boots (252442, -0.24 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 13.4 attack_power points (0.48 DPS) | yes | The 1 Ring (8350, -0.36 DPS) [world]; Lavishly Jeweled Ring (1156, -0.38 DPS) [dungeon]; Ring of the Moon (12052, -0.41 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | The 1 Ring (8350, -0.17 DPS) [world]; Lavishly Jeweled Ring (1156, -0.19 DPS) [dungeon]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 303.4 attack_power points (10.81 DPS) | yes | Night Reaver (1318, -0.53 DPS) [dungeon]; Staff of Westfall (2042, -0.58 DPS) [quest]; Living Root (6631, -0.85 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Grave Shroud; chest: Armor of the Fang; wrist: Bravo's Armbands; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Demon Band; main_hand: Smite's Mighty Hammer

No-known-source sample (15 of 217, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 30 (dwarf, 0000000000000000-255130030002000000-0000000000000000)

Set DPS (verified): 90.6. Weights run: 1.5s. Verify run: 2.6s. 382 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.004, strength=2.000 ± 0.008, agility=1.278 ± 0.044, crit=1.076 ± 0.051 per rating point (14 rating = 1%, 15.061 per %), hit=2.134 ± 0.122 per rating point (10 rating = 1%, 21.341 per %), melee_haste=9.457 ± 0.684

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (1.19 DPS) | yes | Azure Gustwoven Hood (277050, -0.30 DPS) [crafted]; Defender's Leather Hood (252447, -0.40 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.47 DPS, sim-verified) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 13.1 attack_power points (0.65 DPS) | yes | Sentinel's Medallion (19541, -0.14 DPS) [rep]; River Pride Choker (13087, -0.25 DPS) [world_drop]; Fallen Guard's Pendant (279837, -0.72 DPS, sim-verified) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 24.1 attack_power points (1.19 DPS) | yes | Bristlebark Amice (14573, -0.52 DPS) [world_drop]; Mantle of Thieves (2264, -0.56 DPS) [dungeon]; Barbaric Shoulders (5964, -0.74 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 14.9 attack_power points (0.74 DPS) | yes | Sergeant Major's Cape (16315, -0.09 DPS) [pvp]; Tigerstrike Mantle (13108, -0.23 DPS) [world_drop]; Wolfmaster Cape (6314, -0.25 DPS) [dungeon] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 22.2 attack_power points (1.10 DPS) | yes | Brawler's Leather Armor (252490, -0.16 DPS) [crafted]; Thick Murloc Armor (5782, -0.21 DPS) [crafted]; Defender's Leather Tunic (252450, -1.14 DPS, sim-verified) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 15.7 attack_power points (0.78 DPS) | yes | Barbaric Bracers (18948, -0.13 DPS) [crafted]; Bands of Serra'kis (6902, -0.18 DPS) [dungeon]; Demonhide Bracers (270033, -0.23 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 21.7 attack_power points (1.08 DPS) | yes | Toughened Leather Gloves (4253, -0.10 DPS) [crafted]; Wolfclaw Gloves (1978, -0.20 DPS) [dungeon]; Heavy Earthen Gloves (7359, -0.28 DPS) [crafted] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 25.7 attack_power points (1.27 DPS) | yes | Highlander's Leather Girdle (20117, -0.08 DPS) [rep]; Skulker's Leather Belt (252520, -0.11 DPS) [crafted]; Highlander's Chain Girdle (20090, -1.28 DPS, sim-verified) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.29 DPS) | yes | Brawler's Leather Pants (252500, -0.12 DPS) [crafted]; Trapper's Leather Pants (252501, -0.12 DPS) [crafted]; Brawler's Leather Legguards (252516, -1.95 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 16.4 attack_power points (0.81 DPS) | yes | Feet of the Lynx (1121, -0.01 DPS) [world_drop]; Disjointed Shoes (277226, -0.22 DPS) [quest]; Insignia Boots (4055, -0.31 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.8 attack_power points (0.98 DPS) | yes | Tiger Band (6749, -0.29 DPS) [quest]; Silverlaine's Family Seal (6321, -0.49 DPS) [dungeon]; Seal of Wrynn (2933, -0.50 DPS) [quest] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 19.7 attack_power points (0.98 DPS) | yes | Silverlaine's Family Seal (6321, -0.48 DPS) [dungeon]; Seal of Wrynn (2933, -0.49 DPS) [quest]; Tiger Band (6749, -0.57 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (90.6 DPS) | yes | Corpsemaker (6687, -0.45 DPS) [dungeon]; Darkspear Raider's Reaper (272082, -1.08 DPS) [vendor]; Viscous Hammer (13045, -22.20 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Ferine Leggings; feet: Brawler's Leather Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 382, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots

### Band 40 (dwarf, 0000000000000000-255130030005102031-0000000000000000)

Set DPS (verified): 119.0. Weights run: 1.5s. Verify run: 3.2s. 590 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.007, agility=1.249 ± 0.055, crit=1.234 ± 0.068 per rating point (14 rating = 1%, 17.279 per %), hit=1.963 ± 0.127 per rating point (10 rating = 1%, 19.629 per %), melee_haste=7.876 ± 0.968

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 35.7 attack_power points (2.52 DPS) | yes | Barbaric Iron Helm (7915, -0.46 DPS) [crafted]; Hard Gold Coif (250537, -0.55 DPS) [crafted]; Tusken Helm (6686, -0.69 DPS) [dungeon] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | sim-verified (119.0 DPS) | yes | Fallen Guard's Pendant (279837, -0.07 DPS) [quest]; Sentinel's Medallion (19541, -0.21 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 25.7 attack_power points (1.82 DPS) | yes | Forest Tracker Epaulets (2278, -0.14 DPS) [world_drop]; Flintrock Shoulders (7755, -0.23 DPS) [dungeon]; Hard Gold Pauldrons (250539, -0.26 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 19.5 attack_power points (1.38 DPS) | yes | Yeti Fur Cloak (2805, -0.42 DPS) [quest]; Scorpashi Cape (14656, -0.65 DPS) [world_drop]; Hawkeye's Cloak (14593, -1.38 DPS, sim-verified) [world_drop] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 36.7 attack_power points (2.59 DPS) | yes | Avenger's Armor (1488, -0.48 DPS) [dungeon]; Shining Silver Breastplate (2870, -0.62 DPS) [crafted]; Golden Scale Cuirass (3845, -0.62 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.41 DPS) | yes | Ravager's Armguards (14770, -0.07 DPS) [world_drop]; Pugilist Bracers (4438, -0.28 DPS) [dungeon]; Yorgen Bracers (13012, -0.30 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 37.3 attack_power points (2.63 DPS) | yes | Scarlet Gauntlets (10331, -0.06 DPS) [dungeon]; Gauntlets of Divinity (7724, -0.37 DPS) [dungeon]; Prowler's Leather Gloves (252524, -0.43 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (2.12 DPS) | yes | Ogron's Sash (13117, -0.05 DPS) [world_drop]; Officer's Belt (250556, -0.18 DPS) [crafted]; Prowler's Leather Belt (252459, -0.32 DPS) [crafted] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.96 DPS) | yes | Firemane Leggings (13129, -0.28 DPS) [world_drop]; Orcish War Leggings (7929, -0.56 DPS) [crafted]; Legguards of the Vault (9396, -0.69 DPS) [dungeon] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 30.7 attack_power points (2.17 DPS) | yes | Blackforge Greaves (6423, -0.02 DPS) [dungeon]; Skulker's Leather Shoes (252531, -0.21 DPS) [crafted]; Excelsior Boots (4109, -0.39 DPS) [quest] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 26.0 attack_power points (1.83 DPS) | yes | Assault Band (13095, -0.42 DPS) [world_drop]; Thunderbrow Ring (13097, -0.44 DPS) [world_drop]; Falcon's Hook (7552, -0.62 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.41 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.02 DPS) [world_drop]; Falcon's Hook (7552, -0.19 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonebiter (6830, +0.00 DPS) [quest]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor]; Ravager (7717, -3.48 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; shoulder: Sunburn Spaulders; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; legs: Scarlet Leggings; feet: Prowler's Leather Shoes; finger1: Protector's Band; finger2: Mark of Kern

No-known-source sample (15 of 590, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4833 Glorious Shoulders

### Band 50 (dwarf, 3230000000000000-255130030005102051-0000000000000000)

Set DPS (verified): 158.3. Weights run: 1.6s. Verify run: 5.9s. 753 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.006, agility=1.268 ± 0.063, crit=1.548 ± 0.083 per rating point (14 rating = 1%, 21.672 per %), hit=2.655 ± 0.154 per rating point (10 rating = 1%, 26.554 per %), melee_haste=9.183 ± 1.144

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Mail Helmet (223075) | Captain Dirgehammer [vendor] | 43.7 attack_power points (3.12 DPS) | yes | Barbaric Iron Helm (7915, -1.02 DPS) [crafted]; Hard Gold Coif (250537, -1.12 DPS) [crafted]; White Bandit Mask (10008, -5.22 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 26.5 attack_power points (1.89 DPS) | yes | Sentinel's Medallion (19540, -0.89 DPS) [rep]; Kaleidoscope Chain (13084, -0.96 DPS) [world_drop] |
| shoulder | Knight-Lieutenant's Mail Epaulets (223073) | Captain Dirgehammer [vendor] | 39.7 attack_power points (2.83 DPS) | yes | Failed Flying Experiment (9647, -0.50 DPS) [quest]; Skulker's Leather Shoulder (252535, -0.55 DPS) [crafted]; Prowler's Leather Shoulder (252534, -2.76 DPS, sim-verified) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 19.6 attack_power points (1.40 DPS) | yes | Dark Phantom Cape (13122, -0.04 DPS) [world_drop]; Bloodlust Cape (14801, -0.11 DPS) [world_drop]; Pridelord Cape (14673, -0.16 DPS) [world_drop] |
| chest | Knight's Mail Armor (223078) | Captain Dirgehammer [vendor] | 45.7 attack_power points (3.26 DPS) | yes | Kolkar Marauder Chain (6773, -0.63 DPS) [quest]; Blazewind Breastplate (11193, -0.75 DPS) [quest]; Warbear Harness (15064, -2.61 DPS, sim-verified) [crafted] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | sim-verified (158.3 DPS) | yes | Deepfury Bracers (13120, -0.07 DPS) [world_drop]; Prowler's Leather Bracers (252539, -0.08 DPS) [crafted]; Arena Bands (18711, -1.69 DPS, sim-verified) [world] |
| hands | Sergeant Major's Mail Gauntlets (223076) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; Fists of The Five Thunders (227022, +0.00 DPS) [vendor]; Raider Gloves (272100, -2.58 DPS, sim-verified) [vendor] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 43.2 attack_power points (3.08 DPS) | yes | Skulker's Leather Waistguard (252474, -0.10 DPS) [crafted]; Highlander's Chain Girdle (20088, -0.11 DPS) [rep]; Highlander's Leather Girdle (20115, -0.11 DPS) [rep] |
| legs | Knight's Mail Legplates (223074) | Captain Dirgehammer [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpentskin Leggings (8262, +0.00 DPS) [world_drop]; Stormshroud Pants (15057, -0.17 DPS) [crafted]; Gryphon Rider's Leggings (9652, -2.41 DPS, sim-verified) [quest] |
| feet | Sergeant Major's Mail Sabatons (223077) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Prowler's Leather Shoes (252465, +0.00 DPS) [crafted]; Prowler's Leather Boots (252468, +0.00 DPS) [crafted]; Skulker's Leather Boots (252469, +0.00 DPS) [crafted] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 31.4 attack_power points (2.24 DPS) | yes | Mark of Kern (2262, -0.81 DPS) [dungeon]; Assault Band (13095, -0.81 DPS) [world_drop]; Thunderbrow Ring (13097, -0.83 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) (or Mark of Kern (2262), Assault Band (13095)) | Maraudon: Princess Theradras [dungeon] | 20.0 attack_power points (1.43 DPS) | yes | Mark of Kern (2262, +0.00 DPS) [dungeon]; Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.01 DPS) [world_drop] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smoking Heart of the Mountain (11811, +0.00 DPS) [crafted]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thorium Greatmace (250613, -0.93 DPS) [crafted]; Darkspear Raider's Reaper (272080, -1.62 DPS) [vendor]; Ravager (7717, -10.65 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Knight-Lieutenant's Mail Helmet; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Mail Epaulets; chest: Knight's Mail Armor; wrist: Bracers of the Stone Princess; hands: Sergeant Major's Mail Gauntlets; waist: Prowler's Leather Waistguard; legs: Knight's Mail Legplates; feet: Sergeant Major's Mail Sabatons; finger1: Protector's Band; finger2: Blackstone Ring; trinket1: Molten Heart of the Mountain; trinket2: Mark of the Chosen; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 753, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4833 Glorious Shoulders

### Band 60 (dwarf, 3230300000000000-255130030005102051-0520000000000000)

Set DPS (verified): 236.4. Weights run: 1.6s. Verify run: 13.8s. 1459 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.006, agility=1.713 ± 0.089, crit=2.399 ± 0.127 per rating point (14 rating = 1%, 33.582 per %), hit=4.035 ± 0.229 per rating point (10 rating = 1%, 40.347 per %), melee_haste=16.679 ± 1.975

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Face of The Five Thunders (227021, -0.29 DPS) [vendor]; Blue Suede Hat (252482, -0.59 DPS) [crafted]; Outlaw's Collar (279253, -6.11 DPS, sim-verified) [crafted] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 64.3 attack_power points (4.64 DPS) | yes | Medallion of the Dawn (22659, -0.49 DPS) [quest]; Amulet of the Darkmoon (19491, -0.85 DPS) [quest]; Skibi's Pendant (13089, -2.31 DPS) [world_drop] |
| shoulder | Highlander's Mail Pauldrons (20056) | Samuel Hawke [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Mantle of the Dawn (227818, +0.00 DPS) [vendor]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, -4.24 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 68.3 attack_power points (4.93 DPS) | yes | Cape of the Black Baron (13340, -1.63 DPS) [dungeon]; Cloak of the Honor Guard (20073, -1.86 DPS) [rep]; Stalwart Cloak (272415, -2.02 DPS) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 137.3 attack_power points (9.90 DPS) | yes | Timbermaw Tunic (252484, -0.50 DPS) [crafted]; Obsidian Mail Tunic (22191, -2.00 DPS) [crafted]; Shadowcraft Tunic (16721, -4.36 DPS) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 54.5 attack_power points (3.93 DPS) | yes | Bracers of Subterfuge (22668, -0.64 DPS) [quest]; Windtalker's Wristguards (19582, -1.19 DPS) [rep] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | 73.9 attack_power points (5.33 DPS) | yes | Raider Gloves (272099, -0.41 DPS) [vendor]; Timbermaw Brawlers (19049, -0.57 DPS) [crafted]; Devilsaur Gauntlets (15063, -0.89 DPS) [crafted] |
| waist | Highlander's Mail Girdle (20044) | Samuel Hawke [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Highlander's Chain Girdle (20043, +0.00 DPS) [rep]; Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Ferocity of the Timbermaw (227805, +0.00 DPS) [vendor] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 167.4 attack_power points (12.07 DPS) | yes | Sentinel's Leather Pants (237818, -3.90 DPS) [vendor]; Warbear Woolies (15065, -5.81 DPS) [crafted] |
| feet | Highlander's Mail Greaves (20051) | Samuel Hawke [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Drudge Boots (21532, +0.00 DPS) [quest]; Fine Dawn Treaders (227815, +0.00 DPS) [vendor]; Scalegut Treaders (275618, +0.00 DPS) [crafted] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 89.9 attack_power points (6.48 DPS) | yes | Cutthroat's Signet (272408, -1.85 DPS) [vendor]; Protector's Band (19514, -3.40 DPS) [rep]; Channeler's Ring (272406, -3.58 DPS) [vendor] |
| finger2 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | 88.3 attack_power points (6.37 DPS) | yes | Cutthroat's Signet (272408, -1.73 DPS) [vendor]; Protector's Band (19514, -3.28 DPS) [rep]; Channeler's Ring (272406, -3.46 DPS) [vendor] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkmoon Card: Heroism (19287, -10.79 DPS, sim-verified) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (236.4 DPS) | yes | Darkmoon Card: Blue Dragon (19288, -2.93 DPS, sim-verified) [quest] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Annihilator (12798, -51.20 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Black Dragonscale Helm; neck: Beads of Ogre Might; shoulder: Highlander's Mail Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Highlander's Mail Girdle; legs: Sentinel's Chain Leggings; feet: Highlander's Mail Greaves; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Frozen Heart of the Mountain; main_hand: The Unstoppable Force

No-known-source sample (15 of 1459, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4833 Glorious Shoulders

### Band 60, raid preset (dwarf, 3230031000000000-255030031005102031-0530000000000000)

Set DPS (verified): 653.9. Weights run: 1.5s. Verify run: 11.7s. 1459 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.200 ± 0.006, agility=2.271 ± 0.120, crit=2.991 ± 0.162 per rating point (14 rating = 1%, 41.867 per %), hit=5.688 ± 0.407 per rating point (10 rating = 1%, 56.878 per %), melee_haste=22.765 ± 3.000

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (653.9 DPS) | yes | Face of The Five Thunders (227021, -1.73 DPS) [vendor]; Blue Suede Hat (252482, -2.20 DPS) [crafted]; Outlaw's Collar (279253, -12.83 DPS, sim-verified) [crafted] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 80.9 attack_power points (11.23 DPS) | yes | Medallion of the Dawn (22659, -2.08 DPS) [quest]; Amulet of the Darkmoon (19491, -2.18 DPS) [quest]; Skibi's Pendant (13089, -5.60 DPS) [world_drop] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 103.2 attack_power points (14.33 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Glowing Mantle of the Dawn (227818, -2.01 DPS) [vendor]; Highlander's Leather Shoulders (20059, -4.49 DPS) [rep] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 84.9 attack_power points (11.79 DPS) | yes | Arcanoweave Cloak (272411, -3.89 DPS) [vendor]; Cape of the Black Baron (13340, -4.28 DPS) [dungeon]; Stalwart Cloak (272415, -5.70 DPS, sim-verified) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 182.8 attack_power points (25.38 DPS) | yes | Timbermaw Tunic (252484, +0.00 DPS, sim-verified) [crafted]; Obsidian Mail Tunic (22191, -9.01 DPS) [crafted]; Shadowcraft Tunic (16721, -10.86 DPS) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 67.3 attack_power points (9.35 DPS) | yes | Bracers of Subterfuge (22668, -1.57 DPS) [quest]; Deepfury Bracers (13120, -3.40 DPS) [world_drop]; Arena Wristguards (18709, -3.54 DPS) [world] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | 98.7 attack_power points (13.71 DPS) | yes | Raider Gloves (272099, -1.84 DPS) [vendor]; Timbermaw Brawlers (19049, -2.88 DPS) [crafted] |
| waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | 121.7 attack_power points (16.90 DPS) | yes | Ferocity of the Timbermaw (227805, -3.60 DPS) [vendor]; Girdle of The Five Thunders (227018, -5.95 DPS) [vendor]; Might of the Timbermaw (19044, -6.07 DPS) [crafted] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 220.1 attack_power points (30.56 DPS) | yes | Sentinel's Leather Pants (237818, -10.42 DPS) [vendor]; Warbear Woolies (15065, -16.33 DPS) [crafted] |
| feet | Fine Dawn Treaders (227815) | Argent Quartermaster Hasana [vendor] | 81.9 attack_power points (11.37 DPS) | yes | Drudge Boots (21532, -2.03 DPS) [quest]; Shadowcraft Boots (16711, -2.38 DPS) [dungeon]; Argent Elite Boots (227816, -3.47 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 114.7 attack_power points (15.93 DPS) | yes | Cutthroat's Signet (272408, -3.62 DPS) [vendor]; Channeler's Ring (272406, -8.03 DPS) [vendor]; Blessed Band of Light (272407, -8.03 DPS) [vendor] |
| finger2 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 111.4 attack_power points (15.46 DPS) | yes | Cutthroat's Signet (272408, -3.15 DPS) [vendor]; Channeler's Ring (272406, -7.57 DPS) [vendor]; Blessed Band of Light (272407, -7.57 DPS) [vendor] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Darkmoon Card: Heroism (19287, -9.62 DPS, sim-verified) [quest] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkmoon Card: Heroism (19287, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Annihilator (12798, -120.97 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Black Dragonscale Helm; neck: Beads of Ogre Might; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Belt of Preserved Heads; legs: Sentinel's Chain Leggings; feet: Fine Dawn Treaders; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Talisman of Ascendance; main_hand: The Unstoppable Force

No-known-source sample (15 of 1459, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4833 Glorious Shoulders

## Horde

### Band 20 (orc, 0000000000000000-254000000000000000-0000000000000000)

Set DPS (verified): 37.9. Weights run: 1.5s. Verify run: 4.5s. 197 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.006, agility=1.355 ± 0.041, crit=0.913 ± 0.041 per rating point (14 rating = 1%, 12.786 per %), hit=1.669 ± 0.111 per rating point (10 rating = 1%, 16.690 per %), melee_haste=9.111 ± 0.568

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (0.57 DPS) | yes | Brawler's Leather Hood (252504, -0.31 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 8.1 attack_power points (0.29 DPS) | yes | Erudite's Amulet (277204, -0.19 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 6.8 attack_power points (0.24 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS) [dungeon] |
| back | Grave Shroud (279865) | Unending Torment [quest] | 8.7 attack_power points (0.31 DPS) | yes | Glowing Lizardscale Cloak (6449, -0.02 DPS) [dungeon]; Dark Leather Cloak (2316, -0.02 DPS) [crafted]; Lambent Scale Cloak (4706, -0.03 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (37.9 DPS) | yes | Murloc Scale Breastplate (5781, +0.00 DPS) [crafted]; Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -0.98 DPS, sim-verified) [crafted] |
| wrist | Bristlebark Bindings (14569) | World drop [world_drop] | 8.1 attack_power points (0.29 DPS) | yes | Forest Leather Bracers (3202, -0.05 DPS) [world_drop]; Wolf Bracers (4794, -0.09 DPS) [vendor]; Ratchet Wristwraps (274742, -0.14 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (37.9 DPS) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Gloves of the Fang (10413, -1.17 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.64 DPS) | yes | Deviate Scale Belt (6468, -0.19 DPS) [crafted]; Murloc Scale Belt (5780, -0.35 DPS) [crafted]; Brawler's Leather Belt (252428, -1.47 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (37.9 DPS) | yes | Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -1.04 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 16.8 attack_power points (0.60 DPS) | yes | Brawler's Leather Boots (252439, -0.00 DPS) [crafted]; Defender's Leather Boots (252441, -0.24 DPS) [crafted]; Totemic Leather Boots (252442, -0.24 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 13.4 attack_power points (0.48 DPS) | yes | Loop of Sacrifice (281673, -0.26 DPS) [quest]; Bounty Hunter's Ring (5351, -0.33 DPS) [quest]; The 1 Ring (8350, -0.36 DPS) [world] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Loop of Sacrifice (281673, -0.07 DPS) [quest]; Bounty Hunter's Ring (5351, -0.14 DPS) [quest]; The 1 Ring (8350, -0.17 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hammerbone (270018) | Leaders of the Fang [quest] | 306.2 attack_power points (10.90 DPS) | yes | Forsaken Greataxe (251533, -0.11 DPS) [quest]; Living Root (6631, -0.33 DPS) [dungeon]; Smite's Mighty Hammer (7230, -0.37 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Grave Shroud; chest: Blackened Defias Armor; wrist: Bristlebark Bindings; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: Hammerbone

No-known-source sample (15 of 197, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 206376 Bulwark Icon; 206383 Lithic Icon; 206384 Astral Icon; 206385 Clastic Icon; 206389 Brimstone Icon; 206390 Ichthyan Icon; 206391 Syzygic Icon; 241089 Scarlet Dagger

### Band 30 (orc, 0000000000000000-255130030002000000-0000000000000000)

Set DPS (verified): 90.7. Weights run: 1.5s. Verify run: 2.4s. 364 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.004, strength=2.000 ± 0.008, agility=1.278 ± 0.044, crit=1.076 ± 0.051 per rating point (14 rating = 1%, 15.061 per %), hit=2.134 ± 0.122 per rating point (10 rating = 1%, 21.341 per %), melee_haste=9.457 ± 0.684

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (1.19 DPS) | yes | Cloudy Gustwoven Hood (277042, -0.30 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.30 DPS) [crafted]; Defender's Leather Hood (252447, -0.40 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 13.1 attack_power points (0.65 DPS) | yes | River Pride Choker (13087, -0.25 DPS) [world_drop]; Scout's Medallion (20442, -0.27 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 24.1 attack_power points (1.19 DPS) | yes | Bristlebark Amice (14573, -0.52 DPS) [world_drop]; Mantle of Thieves (2264, -0.56 DPS) [dungeon]; Barbaric Shoulders (5964, -0.62 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 14.9 attack_power points (0.74 DPS) | yes | Construct Cloak (279848, -0.05 DPS) [quest]; Tigerstrike Mantle (13108, -0.23 DPS) [world_drop]; Wildhunter Cloak (16658, -0.25 DPS) [quest] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 22.2 attack_power points (1.10 DPS) | yes | Dreamer's Chestguard (270043, -0.11 DPS) [quest]; Defender's Leather Tunic (252450, -0.13 DPS) [crafted]; Brawler's Leather Armor (252490, -0.16 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 15.7 attack_power points (0.78 DPS) | yes | Barbaric Bracers (18948, -0.13 DPS) [crafted]; Bands of Serra'kis (6902, -0.18 DPS) [dungeon]; Cultist's Armguards (270032, -0.28 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 21.7 attack_power points (1.08 DPS) | yes | Toughened Leather Gloves (4253, -0.10 DPS) [crafted]; Wolfclaw Gloves (1978, -0.20 DPS) [dungeon]; Heavy Earthen Gloves (7359, -0.28 DPS) [crafted] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 25.7 attack_power points (1.27 DPS) | yes | Defiler's Leather Girdle (20191, -0.08 DPS) [rep]; Skulker's Leather Belt (252520, -0.11 DPS) [crafted]; Defiler's Chain Girdle (20152, -0.98 DPS, sim-verified) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.29 DPS) | yes | Brawler's Leather Pants (252500, -0.12 DPS) [crafted]; Trapper's Leather Pants (252501, -0.12 DPS) [crafted]; Brawler's Leather Legguards (252516, -2.13 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 16.4 attack_power points (0.81 DPS) | yes | Feet of the Lynx (1121, -0.01 DPS) [world_drop]; Stomping Boots (3741, -0.16 DPS) [quest]; Vorrel's Boots (7751, -0.18 DPS) [quest] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.8 attack_power points (0.98 DPS) | yes | Tiger Band (6749, -0.29 DPS) [quest]; Band of the Fist (17694, -0.33 DPS) [quest]; Silverlaine's Family Seal (6321, -0.49 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 19.7 attack_power points (0.98 DPS) | yes | Tiger Band (6749, -0.28 DPS) [quest]; Band of the Fist (17694, -0.33 DPS) [quest]; Silverlaine's Family Seal (6321, -0.48 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (90.7 DPS) | yes | Corpsemaker (6687, -0.45 DPS) [dungeon]; Darkspear Raider's Reaper (272082, -1.08 DPS) [vendor]; Viscous Hammer (13045, -22.06 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Ferine Leggings; feet: Brawler's Leather Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 364, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (orc, 0000000000000000-255130030005102031-0000000000000000)

Set DPS (verified): 117.5. Weights run: 1.5s. Verify run: 3.3s. 553 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.007, agility=1.249 ± 0.055, crit=1.234 ± 0.068 per rating point (14 rating = 1%, 17.279 per %), hit=1.963 ± 0.127 per rating point (10 rating = 1%, 19.629 per %), melee_haste=7.876 ± 0.968

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 35.7 attack_power points (2.52 DPS) | yes | Hard Gold Coif (250537, -0.55 DPS) [crafted]; Tusken Helm (6686, -0.69 DPS) [dungeon]; Barbaric Iron Helm (7915, -0.92 DPS, sim-verified) [crafted] |
| neck | Ethereal Talisman (4430) | The Crown of Will [quest] | 15.0 attack_power points (1.06 DPS) | yes | Scout's Medallion (19536, -0.09 DPS) [rep]; Kaleidoscope Chain (13084, -0.14 DPS) [world_drop] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 25.7 attack_power points (1.82 DPS) | yes | Forest Tracker Epaulets (2278, -0.14 DPS) [world_drop]; Flintrock Shoulders (7755, -0.23 DPS) [dungeon]; Hard Gold Pauldrons (250539, -0.26 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 14.7 attack_power points (1.04 DPS) | yes | Construct Cloak (279848, -0.05 DPS) [quest]; Scorpashi Cape (14656, -0.32 DPS) [world_drop]; Wildhunter Cloak (16658, -0.33 DPS) [quest] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 36.7 attack_power points (2.59 DPS) | yes | Shining Silver Breastplate (2870, -0.62 DPS) [crafted]; Golden Scale Cuirass (3845, -0.62 DPS) [crafted]; Avenger's Armor (1488, -1.33 DPS, sim-verified) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.41 DPS) | yes | Windtalker's Wristguards (19584, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Ravager's Armguards (14770, -0.07 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 37.3 attack_power points (2.63 DPS) | yes | Scarlet Gauntlets (10331, -0.06 DPS) [dungeon]; Gauntlets of Divinity (7724, -0.37 DPS) [dungeon]; Prowler's Leather Gloves (252524, -0.43 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (2.12 DPS) | yes | Ogron's Sash (13117, -0.05 DPS) [world_drop]; Officer's Belt (250556, -0.18 DPS) [crafted]; Tharg's Shoelace (9705, -0.28 DPS) [quest] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.96 DPS) | yes | Firemane Leggings (13129, -0.28 DPS) [world_drop]; Orcish War Leggings (7929, -0.56 DPS) [crafted]; Legguards of the Vault (9396, -0.69 DPS) [dungeon] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 30.7 attack_power points (2.17 DPS) | yes | Skulker's Leather Shoes (252531, -0.21 DPS) [crafted]; Excelsior Boots (4109, -0.39 DPS) [quest]; Blackforge Greaves (6423, -1.40 DPS, sim-verified) [dungeon] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 26.0 attack_power points (1.83 DPS) | yes | Assault Band (13095, -0.42 DPS) [world_drop]; Thunderbrow Ring (13097, -0.44 DPS) [world_drop]; Falcon's Hook (7552, -0.62 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.41 DPS) | yes | Thunderbrow Ring (13097, -0.02 DPS) [world_drop]; Falcon's Hook (7552, -0.19 DPS) [dungeon]; Assault Band (13095, -1.72 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (117.5 DPS) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor]; Ravager (7717, -6.14 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; neck: Ethereal Talisman; shoulder: Sunburn Spaulders; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; legs: Scarlet Leggings; feet: Prowler's Leather Shoes; finger1: Legionnaire's Band; finger2: Mark of Kern

No-known-source sample (15 of 553, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7748 Forcestone Buckler; 8708 Hammer of Expertise

### Band 50 (orc, 3230000000000000-255130030005102051-0000000000000000)

Set DPS (verified): 154.7. Weights run: 1.6s. Verify run: 4.1s. 698 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.006, agility=1.268 ± 0.063, crit=1.548 ± 0.083 per rating point (14 rating = 1%, 21.672 per %), hit=2.655 ± 0.154 per rating point (10 rating = 1%, 26.554 per %), melee_haste=9.183 ± 1.144

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 35.9 attack_power points (2.56 DPS) | yes | Blood Guard's Mail Helmet (220820, +0.00 DPS) [vendor]; Undercity Reservist's Cap (20643, -0.14 DPS) [quest]; Barbaric Iron Helm (7915, -0.47 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 26.5 attack_power points (1.89 DPS) | yes | Woven Ivy Necklace (19159, -0.22 DPS) [quest]; Scout's Medallion (19535, -0.80 DPS) [rep]; Ethereal Talisman (4430, -0.81 DPS) [quest] |
| shoulder | Prowler's Leather Shoulder (252534) | Leatherworking [crafted] | 33.4 attack_power points (2.38 DPS) | yes | Blood Guard's Mail Epaulets (220823, +0.00 DPS) [vendor]; Skulker's Leather Shoulder (252535, -0.10 DPS) [crafted]; Failed Flying Experiment (9647, -4.33 DPS, sim-verified) [quest] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 19.0 attack_power points (1.36 DPS) | yes | Bloodlust Cape (14801, -0.07 DPS) [world_drop]; Pridelord Cape (14673, -0.11 DPS) [world_drop]; Blackflame Cape (13109, -0.27 DPS) [world_drop] |
| chest | Stone Guard's Mail Armor (220826) | Lady Palanseer [vendor] | 45.7 attack_power points (3.26 DPS) | yes | Kolkar Marauder Chain (6773, -0.63 DPS) [quest]; Blazewind Breastplate (11193, -0.75 DPS) [quest]; Warbear Harness (15064, -6.57 DPS, sim-verified) [crafted] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Windtalker's Wristguards (19583, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Arena Bands (18711, -2.34 DPS, sim-verified) [world] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 47.9 attack_power points (3.42 DPS) | yes | Gloves of Holy Might (867, -0.44 DPS) [world_drop]; Fists of The Five Thunders (227022, -0.59 DPS) [vendor]; Prowler's Leather Gauntlets (252547, -0.60 DPS) [crafted] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 43.2 attack_power points (3.08 DPS) | yes | Skulker's Leather Waistguard (252474, -0.10 DPS) [crafted]; Defiler's Chain Girdle (20151, -0.11 DPS) [rep]; Defiler's Leather Girdle (20193, -0.11 DPS) [rep] |
| legs | Stone Guard's Mail Legplates (220834) | Lady Palanseer [vendor] | sim-verified (+2.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Stormshroud Pants (15057, -0.17 DPS) [crafted]; Scarlet Leggings (10330, -0.26 DPS) [dungeon]; Serpentskin Leggings (8262, -2.81 DPS, sim-verified) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 39.9 attack_power points (2.85 DPS) | yes | Skulker's Leather Boots (252469, -0.10 DPS) [crafted]; Prowler's Leather Shoes (252465, -0.65 DPS) [crafted]; Blackforge Greaves (6423, -0.66 DPS) [dungeon] |
| finger1 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 31.4 attack_power points (2.24 DPS) | yes | Mark of Kern (2262, -0.81 DPS) [dungeon]; Assault Band (13095, -0.81 DPS) [world_drop]; Blackstone Ring (17713, -0.81 DPS) [dungeon] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.71 DPS) | yes | Mark of Kern (2262, -0.29 DPS) [dungeon]; Assault Band (13095, -0.29 DPS) [world_drop]; Blackstone Ring (17713, -0.29 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+5.2 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -1.61 DPS, sim-verified) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thorium Greatmace (250613, -0.93 DPS) [crafted]; Darkspear Raider's Reaper (272080, -1.62 DPS) [vendor]; Ravager (7717, -7.67 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** neck: Skibi's Pendant; shoulder: Prowler's Leather Shoulder; back: Dark Phantom Cape; chest: Stone Guard's Mail Armor; wrist: Bracers of the Stone Princess; hands: Raider Gloves; waist: Prowler's Leather Waistguard; legs: Stone Guard's Mail Legplates; feet: Prowler's Leather Boots; finger1: Legionnaire's Band; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 698, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60 (orc, 3230300000000000-255130030005102051-0520000000000000)

Set DPS (verified): 243.0. Weights run: 1.6s. Verify run: 13.6s. 1384 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.006, agility=1.713 ± 0.089, crit=2.399 ± 0.127 per rating point (14 rating = 1%, 33.582 per %), hit=4.035 ± 0.229 per rating point (10 rating = 1%, 40.347 per %), melee_haste=16.679 ± 1.975

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (243.0 DPS) | yes | Champion's Mail Headguard (227155, +0.00 DPS) [pvp]; Face of The Five Thunders (227021, -0.29 DPS) [vendor]; Outlaw's Collar (279253, -10.26 DPS, sim-verified) [crafted] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 64.3 attack_power points (4.64 DPS) | yes | Amulet of the Darkmoon (19491, -0.85 DPS) [quest]; Skibi's Pendant (13089, -2.31 DPS) [world_drop]; Medallion of the Dawn (22659, -3.11 DPS, sim-verified) [quest] |
| shoulder | Defiler's Lizardhide Shoulders (20175) | The Defilers [rep] | sim-verified (243.0 DPS) | yes | Champion's Mail Pauldrons (227154, +0.00 DPS) [pvp]; Warlord's Mail Pauldrons (231654, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, -8.75 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 68.3 attack_power points (4.93 DPS) | yes | Deathguard's Cloak (20068, -1.86 DPS) [rep]; Stalwart Cloak (272415, -2.02 DPS) [vendor]; Cape of the Black Baron (13340, -3.50 DPS, sim-verified) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (243.0 DPS) | yes | Legionnaire's Mail Hauberk (227157, -0.90 DPS) [pvp]; Warlord's Mail Hauberk (231653, -0.90 DPS) [vendor]; Dawn Armor (252483, -2.78 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 54.5 attack_power points (3.93 DPS) | yes | Bracers of Subterfuge (22668, -0.64 DPS) [quest]; Windtalker's Wristguards (19582, -1.19 DPS) [rep] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | 73.9 attack_power points (5.33 DPS) | yes | Blood Guard's Mail Vices (227159, +0.00 DPS) [pvp]; General's Mail Vices (231655, +0.00 DPS) [vendor]; Raider Gloves (272099, -2.06 DPS, sim-verified) [vendor] |
| waist | Defiler's Lizardhide Girdle (20171) | The Defilers [rep] | sim-verified (243.0 DPS) | yes | Defiler's Chain Girdle (20150, +0.00 DPS) [rep]; Ferocity of the Timbermaw (227805, +0.00 DPS) [vendor]; Belt of Preserved Heads (20216, -3.43 DPS, sim-verified) [quest] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 167.4 attack_power points (12.07 DPS) | yes | General's Mail Legguards (231658, -2.71 DPS) [vendor]; Outrider's Chain Leggings (22673, -3.26 DPS, sim-verified) [rep]; Legionnaire's Mail Legguards (227156, -3.57 DPS) [pvp] |
| feet | Defiler's Lizardhide Boots (20167) | The Defilers [rep] | sim-verified (243.0 DPS) | yes | Blood Guard's Mail Greaves (227158, +0.00 DPS) [pvp]; General's Mail Greaves (231656, +0.00 DPS) [vendor]; Fine Dawn Treaders (227815, -5.76 DPS, sim-verified) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 89.9 attack_power points (6.48 DPS) | yes | Cutthroat's Signet (272408, -1.85 DPS) [vendor]; Legionnaire's Band (19510, -3.40 DPS) [rep]; Channeler's Ring (272406, -3.58 DPS) [vendor] |
| finger2 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | 88.3 attack_power points (6.37 DPS) | yes | Cutthroat's Signet (272408, -1.73 DPS) [vendor]; Legionnaire's Band (19510, -3.28 DPS) [rep]; Channeler's Ring (272406, -3.46 DPS) [vendor] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (243.0 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (243.0 DPS) | yes | Frozen Heart of the Mountain (249469, -4.81 DPS, sim-verified) [crafted] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (243.0 DPS) | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; Annihilator (12798, -52.48 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Black Dragonscale Helm; neck: Beads of Ogre Might; shoulder: Defiler's Lizardhide Shoulders; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Defiler's Lizardhide Girdle; legs: Sentinel's Chain Leggings; feet: Defiler's Lizardhide Boots; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: The Unstoppable Force

No-known-source sample (15 of 1384, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60, raid preset (orc, 3230031000000000-255030031005102031-0530000000000000)

Set DPS (verified): 671.2. Weights run: 1.5s. Verify run: 11.5s. 1384 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.200 ± 0.006, agility=2.271 ± 0.120, crit=2.991 ± 0.162 per rating point (14 rating = 1%, 41.867 per %), hit=5.688 ± 0.407 per rating point (10 rating = 1%, 56.878 per %), melee_haste=22.765 ± 3.000

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (+13.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Champion's Mail Headguard (227155, +0.00 DPS) [pvp]; Face of The Five Thunders (227021, -1.73 DPS) [vendor]; Outlaw's Collar (279253, -13.79 DPS, sim-verified) [crafted] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 80.9 attack_power points (11.23 DPS) | yes | Medallion of the Dawn (22659, -2.08 DPS) [quest]; Amulet of the Darkmoon (19491, -2.18 DPS) [quest]; Skibi's Pendant (13089, -5.60 DPS) [world_drop] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 103.2 attack_power points (14.33 DPS) | yes | Champion's Mail Pauldrons (227154, +0.00 DPS) [pvp]; Warlord's Mail Pauldrons (231654, +0.00 DPS) [vendor]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 84.9 attack_power points (11.79 DPS) | yes | Arcanoweave Cloak (272411, -3.89 DPS) [vendor]; Cape of the Black Baron (13340, -4.28 DPS) [dungeon]; Stalwart Cloak (272415, -5.76 DPS, sim-verified) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (+8.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Legionnaire's Mail Hauberk (227157, -1.14 DPS) [pvp]; Warlord's Mail Hauberk (231653, -1.14 DPS) [vendor]; Dawn Armor (252483, -8.95 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 67.3 attack_power points (9.35 DPS) | yes | Bracers of Subterfuge (22668, -1.57 DPS) [quest]; Deepfury Bracers (13120, -3.40 DPS) [world_drop]; Arena Wristguards (18709, -3.54 DPS) [world] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | 98.7 attack_power points (13.71 DPS) | yes | Blood Guard's Mail Vices (227159, +0.00 DPS) [pvp]; General's Mail Vices (231655, +0.00 DPS) [vendor]; Raider Gloves (272099, +0.00 DPS, sim-verified) [vendor] |
| waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | 121.7 attack_power points (16.90 DPS) | yes | Ferocity of the Timbermaw (227805, -3.60 DPS) [vendor]; Girdle of The Five Thunders (227018, -5.95 DPS) [vendor]; Might of the Timbermaw (19044, -6.07 DPS) [crafted] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 220.1 attack_power points (30.56 DPS) | yes | Outrider's Chain Leggings (22673, -5.81 DPS) [rep]; General's Mail Legguards (231658, -8.30 DPS) [vendor]; Legionnaire's Mail Legguards (227156, -10.13 DPS) [pvp] |
| feet | Drudge Boots (21532) | The Nightmare Manifests [quest] | sim-verified (+7.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Blood Guard's Mail Greaves (227158, +0.00 DPS) [pvp]; General's Mail Greaves (231656, +0.00 DPS) [vendor]; Fine Dawn Treaders (227815, -6.96 DPS, sim-verified) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 114.7 attack_power points (15.93 DPS) | yes | Cutthroat's Signet (272408, -3.62 DPS) [vendor]; Channeler's Ring (272406, -8.03 DPS) [vendor]; Blessed Band of Light (272407, -8.03 DPS) [vendor] |
| finger2 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 111.4 attack_power points (15.46 DPS) | yes | Cutthroat's Signet (272408, -3.15 DPS) [vendor]; Channeler's Ring (272406, -7.57 DPS) [vendor]; Blessed Band of Light (272407, -7.57 DPS) [vendor] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Darkmoon Card: Heroism (19287, -7.81 DPS, sim-verified) [quest] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -4.25 DPS) [crafted]; Talisman of Ascendance (22678, -6.64 DPS, sim-verified) [quest] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; Annihilator (12798, -118.43 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Black Dragonscale Helm; neck: Beads of Ogre Might; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Belt of Preserved Heads; legs: Sentinel's Chain Leggings; feet: Drudge Boots; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: The Unstoppable Force

No-known-source sample (15 of 1384, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

