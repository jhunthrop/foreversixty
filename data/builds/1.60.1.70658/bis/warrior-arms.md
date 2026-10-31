# Leveling BiS: Arms

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 03323000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 25.5. Weights run: 1.5s. Verify run: 2.5s. 293 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.209 ± 0.030, crit=0.828 ± 0.027 per rating point (14 rating = 1%, 11.592 per %), hit=1.232 ± 0.083 per rating point (10 rating = 1%, 12.321 per %), melee_haste=8.730 ± 0.630

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.71 DPS) | yes | Defender's Leather Hood (252447, -0.15 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.28 DPS) [crafted]; Brawler's Leather Hood (252504, -0.37 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 2.4 attack_power points (0.09 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 6.0 attack_power points (0.21 DPS) | yes | Rough Bronze Shoulders (3480, -0.00 DPS) [crafted]; Silvered Bronze Shoulders (3481, -0.00 DPS) [crafted] |
| back | Grave Shroud (279865) | Abominable Creatures [quest] | 8.4 attack_power points (0.30 DPS) | yes | Lambent Scale Cloak (4706, -0.01 DPS) [world_drop]; Dark Leather Cloak (2316, -0.03 DPS) [crafted]; Glowing Lizardscale Cloak (6449, -0.04 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (25.5 DPS) | yes | Veteran's Chain Shirt (250488, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, +0.00 DPS) [crafted]; Mutant Scale Breastplate (6627, -1.01 DPS, sim-verified) [dungeon] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bravo's Armbands (270015, -0.04 DPS) [quest]; Cryptwalker Bracers (280095, -0.07 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (25.5 DPS) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Thorbia's Gauntlets (12994, -1.18 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.64 DPS) | yes | Brawler's Leather Belt (252428, -0.18 DPS) [crafted]; Deviate Scale Belt (6468, -0.21 DPS) [crafted]; Cobrahn's Grasp (6460, -1.37 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (25.5 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Veteran's Chain Leggings (250493, -1.15 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 16.0 attack_power points (0.57 DPS) | yes | Feet of the Lynx (1121, -0.01 DPS) [world_drop]; Veteran's Boots (250503, -0.04 DPS) [crafted]; Defender's Leather Boots (252441, -0.21 DPS) [crafted] |
| finger1 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Lavishly Jeweled Ring (1156, -0.20 DPS) [dungeon]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| finger2 | The 1 Ring (8350) | Fishing [world] | 3.2 attack_power points (0.11 DPS) | yes | Lavishly Jeweled Ring (1156, -0.03 DPS) [dungeon]; Ring of the Moon (12052, -0.04 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 302.8 attack_power points (10.71 DPS) | yes | Duskbringer (2205, -0.30 DPS) [dungeon]; Night Reaver (1318, -0.50 DPS) [dungeon]; Living Root (6631, -0.94 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.8 attack_power points (0.17 DPS) | yes | Dwarven Fishing Pole (3567, -0.03 DPS) [quest]; Fine Longbow (11304, -0.03 DPS) [vendor]; Cracked Blacksmith Hammer (285279, -0.03 DPS) [crafted] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Grave Shroud; chest: Blackened Defias Armor; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Brawler's Leather Boots; finger1: Demon Band; finger2: The 1 Ring; main_hand: Smite's Mighty Hammer; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 293, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6478 Rat Stompers; 9602 Brushwood Blade; 14145 Cursed Felblade; 14147 Cavedweller Bracers

### Band 30 (human, 03325213020000000-00000000000000000-000000000000000000)

Set DPS (verified): 72.5. Weights run: 1.6s. Verify run: 1.7s. 551 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=2.138 ± 0.053, crit=1.699 ± 0.057 per rating point (14 rating = 1%, 23.793 per %), hit=1.891 ± 0.186 per rating point (10 rating = 1%, 18.912 per %), melee_haste=10.934 ± 1.331

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 37.2 attack_power points (1.43 DPS) | yes | Tusken Helm (6686, -0.35 DPS, sim-verified) [dungeon]; Guard's Chain Helm (250499, -0.45 DPS) [crafted]; Slayer's Skullcap (14753, -0.49 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 16.6 attack_power points (0.64 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.18 DPS) [world]; River Pride Choker (13087, -0.33 DPS) [world_drop]; Fallen Guard's Pendant (279837, -0.41 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 33.5 attack_power points (1.29 DPS) | yes | Barbaric Iron Shoulders (7913, -0.37 DPS, sim-verified) [crafted]; Mantle of Thieves (2264, -0.47 DPS) [dungeon]; Barbaric Shoulders (5964, -0.49 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 21.0 attack_power points (0.81 DPS) | yes | Sergeant Major's Cape (16315, -0.17 DPS) [pvp]; Tigerstrike Mantle (13108, -0.31 DPS, sim-verified) [world_drop]; Fenrus' Hide (6340, -0.31 DPS) [dungeon] |
| chest | Veteran's Silvered Chain Shirt (250518) | Blacksmithing [crafted] | 32.8 attack_power points (1.26 DPS) | yes | Avenger's Armor (1488, -0.11 DPS) [dungeon]; Dusky Leather Armor (7374, -0.11 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.14 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 20.8 attack_power points (0.80 DPS) | yes | Yorgen Bracers (13012, -0.09 DPS) [world_drop]; Barbaric Bracers (18948, -0.16 DPS) [crafted]; Pugilist Bracers (4438, -0.19 DPS) [dungeon] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 26.8 attack_power points (1.03 DPS) | yes | Toughened Leather Gloves (4253, -0.08 DPS) [crafted]; Mail Combat Gauntlets (4075, -0.09 DPS) [world_drop]; Fletcher's Gloves (7348, -0.12 DPS) [crafted] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 32.8 attack_power points (1.26 DPS) | yes | Skulker's Leather Belt (252520, -0.06 DPS) [crafted]; Prowler's Leather Belt (252459, -0.08 DPS) [crafted]; Warden's Leather Belt (252460, -0.31 DPS) [crafted] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 37.0 attack_power points (1.42 DPS) | yes | Brawler's Leather Pants (252500, -0.22 DPS) [crafted]; Trapper's Leather Pants (252501, -0.22 DPS) [crafted]; Brawler's Leather Legguards (252516, -0.34 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 23.1 attack_power points (0.89 DPS) | yes | Brawler's Leather Boots (252439, -0.09 DPS) [crafted]; Alacritous Treads (277234, -0.09 DPS) [quest]; Veteran's Boots (250503, -0.18 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 22.4 attack_power points (0.86 DPS) | yes | Seal of Wrynn (2933, -0.38 DPS) [quest]; Defias Renegade Ring (1076, -0.45 DPS) [dungeon]; Monkey Ring (6748, -0.45 DPS) [quest] |
| finger2 | Ring of Precision (1491) | Blackfathom Deeps [dungeon] | 12.8 attack_power points (0.49 DPS) | yes | Seal of Wrynn (2933, -0.02 DPS) [quest]; Defias Renegade Ring (1076, -0.08 DPS) [dungeon]; Monkey Ring (6748, -0.08 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (72.5 DPS) | yes | Corpsemaker (6687, -0.52 DPS) [dungeon]; Felbough Arm (273037, -0.91 DPS) [world]; Viscous Hammer (13045, -20.20 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 12.8 attack_power points (0.49 DPS) | yes | Double-barreled Shotgun (2098, -0.02 DPS) [world_drop]; BKP "Sparrow" Smallbore (3042, -0.16 DPS) [world_drop]; Silver Star (3463, -0.16 DPS) [quest] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Veteran's Silvered Chain Shirt; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Feet of the Lynx; finger1: Thunderbrow Ring; finger2: Ring of Precision; main_hand: Manual Crowd Pummeler; ranged: Golemsight Long Gun

No-known-source sample (15 of 551, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 40 (human, 03325213032511000-00000000000000000-000000000000000000)

Set DPS (verified): 95.1. Weights run: 1.8s. Verify run: 1.8s. 781 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.005, agility=1.924 ± 0.047, crit=1.887 ± 0.057 per rating point (14 rating = 1%, 26.420 per %), hit=2.156 ± 0.180 per rating point (10 rating = 1%, 21.561 per %), melee_haste=13.332 ± 1.509

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chromite Barbute (8142) | Uldaman: Ancient Treasure [dungeon] | 47.1 attack_power points (2.41 DPS) | yes | White Bandit Mask (10008, -0.20 DPS) [crafted]; Barbaric Iron Helm (7915, -0.60 DPS) [crafted]; Hawkeye's Helm (14591, -0.92 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.7 attack_power points (0.80 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.19 DPS) [world]; River Pride Choker (13087, -0.39 DPS) [world_drop]; Gazlowe's Charm (13088, -0.39 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 31.2 attack_power points (1.59 DPS) | yes | Flintrock Shoulders (7755, -0.10 DPS) [dungeon]; Barbaric Iron Shoulders (7913, -0.39 DPS) [crafted]; Hard Gold Pauldrons (250539, -0.47 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 23.5 attack_power points (1.20 DPS) | yes | Hawkeye's Cloak (14593, -0.21 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.31 DPS) [quest]; Parachute Cloak (10518, -0.42 DPS) [crafted] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 46.2 attack_power points (2.36 DPS) | yes | Kolkar Marauder Chain (6773, -0.24 DPS) [quest]; Carapace of Tuten'kash (10775, -0.55 DPS) [dungeon]; Wolffear Harness (13110, -0.69 DPS) [world_drop] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 21.7 attack_power points (1.11 DPS) | yes | Branded Leather Bracers (19508, -0.09 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.11 DPS) [world_drop]; Yorgen Bracers (13012, -0.20 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 46.4 attack_power points (2.38 DPS) | yes | Scarlet Gauntlets (10331, -0.16 DPS) [dungeon]; Plated Fist of Hakoo (13071, -0.37 DPS) [world_drop]; Dragonscale Gauntlets (8347, -0.43 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 35.3 attack_power points (1.81 DPS) | yes | Officer's Belt (250556, -0.19 DPS) [crafted]; Prowler's Leather Belt (252459, -0.30 DPS) [crafted]; Skulker's Leather Belt (252520, -0.31 DPS) [crafted] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.15 DPS) | yes | Symbolic Legplates (14829, -0.02 DPS) [world_drop]; Triprunner Dungarees (9624, -0.07 DPS) [quest]; Basilisk Hide Pants (1718, -0.08 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 39.5 attack_power points (2.02 DPS) | yes | Blackforge Greaves (6423, -0.11 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.20 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.22 DPS) [crafted] |
| finger1 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 23.3 attack_power points (1.19 DPS) | yes | Thunderbrow Ring (13097, -0.08 DPS) [world_drop]; Mark of Kern (2262, -0.17 DPS) [dungeon]; Assault Band (13095, -0.17 DPS) [world_drop] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 23.2 attack_power points (1.19 DPS) | yes | Thunderbrow Ring (13097, -0.08 DPS) [world_drop]; Mark of Kern (2262, -0.17 DPS) [dungeon]; Assault Band (13095, -0.17 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Nightblade (1982) | World drop [world_drop] | sim-verified (95.1 DPS) | yes | Primitive Fishing Pole (276203, -1.05 DPS) [vendor]; Bonebiter (6830, -1.23 DPS) [quest]; Fiery War Axe (870, -4.03 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Monolithic Bow (9426) | Uldaman [dungeon] | sim-verified (95.1 DPS) | yes | Master Hunter's Rifle (17687, -0.11 DPS) [quest]; The Silencer (13138, -0.19 DPS) [world_drop]; Bow of Searing Arrows (2825, -0.98 DPS, sim-verified) [world_drop] |

**New at 40:** head: Chromite Barbute; back: Sergeant Major's Cape; chest: Icemail Jerkin; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Falcon's Hook; finger2: Ring of the Underwood; main_hand: Nightblade; ranged: Monolithic Bow

No-known-source sample (15 of 781, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band

### Band 50 (human, 03325213032515001-05000000000000000-000000000000000000)

Set DPS (verified): 176.7. Weights run: 1.8s. Verify run: 2.1s. 1004 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.140 ± 0.059, crit=1.339 ± 0.079 per rating point (14 rating = 1%, 18.740 per %), hit=2.101 ± 0.194 per rating point (10 rating = 1%, 21.014 per %), melee_haste=12.996 ± 1.717

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sunscale Helmet (14849) | World drop [world_drop] | 38.8 attack_power points (3.13 DPS) | yes | Knight-Lieutenant's Plate Helm (220804, +0.00 DPS) [vendor]; Ornate Mithril Helm (7937, -0.01 DPS) [crafted]; Chromite Barbute (8142, -0.09 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 24.8 attack_power points (2.00 DPS) | yes | Kaleidoscope Chain (13084, -0.99 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.03 DPS) [world]; River Pride Choker (13087, -1.35 DPS) [world_drop] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 35.1 attack_power points (2.83 DPS) | yes | Knight-Lieutenant's Plate Pauldrons (220795, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.16 DPS) [world_drop]; Prowler's Leather Shoulders (252534, -0.23 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 18.8 attack_power points (1.52 DPS) | yes | Bloodlust Cape (14801, -0.07 DPS) [world_drop]; Dark Phantom Cape (13122, -0.14 DPS) [world_drop]; Pridelord Cape (14673, -0.21 DPS) [world_drop] |
| chest | Knight's Plate Hauberk (220794) | Captain Dirgehammer [vendor] | 48.7 attack_power points (3.93 DPS) | yes | Warforged Chestplate (11195, -0.06 DPS) [quest]; Warbear Harness (15064, -0.50 DPS) [crafted]; Valorous Chestguard (8274, -0.70 DPS) [world_drop] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 28.8 attack_power points (2.32 DPS) | yes | Bracers of the Stone Princess (17714, -0.07 DPS) [dungeon]; Arena Bands (18711, -0.07 DPS) [world]; Prowler's Leather Bracers (252539, -0.23 DPS) [crafted] |
| hands | Officer's Gloves (250551) | Blacksmithing [crafted] | 42.3 attack_power points (3.40 DPS) | yes | Stonerender Gauntlets (17007, -0.18 DPS) [world_drop]; Gloves of Holy Might (867, -0.28 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, -0.32 DPS) [crafted] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 41.7 attack_power points (3.36 DPS) | yes | Skulker's Leather Waistguard (252474, -0.14 DPS) [crafted]; Officer's Waistguard (250571, -0.18 DPS) [crafted]; Belt of the Gladiator (13134, -0.46 DPS) [world_drop] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | 54.1 attack_power points (4.36 DPS) | yes | Gryphon Rider's Leggings (9652, -0.57 DPS) [quest]; Centurion Legplates (10740, -0.57 DPS) [quest]; Serpentskin Leggings (8262, -0.80 DPS) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 38.5 attack_power points (3.10 DPS) | yes | Skulker's Leather Boots (252469, -0.14 DPS) [crafted]; Officer's Sabatons (250561, -0.18 DPS) [crafted]; Officer's Boots (250546, -0.37 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) (or Mark of Kern (2262), Assault Band (13095)) | Maraudon: Princess Theradras [dungeon] | 20.0 attack_power points (1.61 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Falcon's Hook (7552, -0.30 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.61 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Falcon's Hook (7552, -0.30 DPS) [dungeon] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (176.7 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (176.7 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Blight (7959) | Blacksmithing [crafted] | sim-verified (176.7 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Nightblade (1982, -3.07 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (176.7 DPS) | yes | Houndmaster's Bow (11628, -0.16 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.33 DPS) [world_drop]; Dark Iron Rifle (16004, -1.80 DPS, sim-verified) [crafted] |

**New at 50:** head: Sunscale Helmet; neck: Skibi's Pendant; shoulder: Officer's Pauldrons; chest: Knight's Plate Hauberk; wrist: Officer's Wristguards; hands: Officer's Gloves; waist: Prowler's Leather Waistguard; legs: Knight's Plate Leggings; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: Mark of Kern; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain; main_hand: Blight; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 1004, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60 (human, 03325213032515001-05050000000000000-005000000000000000)

Set DPS (verified): 235.5. Weights run: 1.9s. Verify run: 7.0s. 2015 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.462 ± 0.068, crit=2.089 ± 0.097 per rating point (14 rating = 1%, 29.240 per %), hit=3.886 ± 0.268 per rating point (10 rating = 1%, 38.862 per %), melee_haste=19.865 ± 2.553

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 172.2 attack_power points (14.12 DPS) | yes | Stalwart Helm (250599, -4.71 DPS, sim-verified) [crafted]; Lieutenant Commander's Plate Helm (23314, -5.09 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 29.0 attack_power points (2.38 DPS) | yes | Sunprism Pendant (20649, -1.10 DPS, sim-verified) [quest]; Kaleidoscope Chain (13084, -1.24 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.39 DPS) [world] |
| shoulder | Lightforge Spaulders (16729) | Blackrock Spire: The Beast [dungeon] | 60.3 attack_power points (4.94 DPS) | yes | Lieutenant Commander's Plate Shoulders (227045, +0.00 DPS) [pvp]; Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp]; Dark Warder's Pauldrons (22241, -2.23 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 41.9 attack_power points (3.44 DPS) | yes | Sergeant Major's Cape (16337, -1.17 DPS) [pvp]; Fel Cape (279269, -1.43 DPS, sim-verified) [crafted]; Dark Phantom Cape (13122, -1.64 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 122.2 attack_power points (10.02 DPS) | yes | Timbermaw Tunic (252484, -0.05 DPS) [crafted]; Field Marshal's Plate Armor (231536, -3.32 DPS) [pvp]; Knight-Captain's Plate Hauberk (227046, -4.18 DPS) [pvp] |
| wrist | Tranquil Wristguards (279256) | Leatherworking [crafted] | 32.0 attack_power points (2.62 DPS) | yes | Marshal's Plate Bracers (16481, +0.00 DPS) [pvp]; Officer's Wristguards (250581, -0.10 DPS) [crafted]; Deepfury Bracers (13120, -0.17 DPS) [world_drop] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Plate Gauntlets (231541, +0.00 DPS) [pvp]; Gloves of Holy Might (867, -0.66 DPS) [world_drop]; Timbermaw Brawlers (19049, -4.64 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 62.5 attack_power points (5.12 DPS) | yes | Marshal's Plate Girdle (16482, +0.00 DPS) [pvp]; Knight-Captain's Plate Girdle (16407, -1.29 DPS) [pvp]; Prowler's Leather Waistguard (252473, -1.53 DPS, sim-verified) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Plate Leggings (23301, +0.00 DPS) [vendor]; Marshal's Plate Legguards (231540, +0.00 DPS) [pvp]; Warbear Woolies (15065, -4.84 DPS, sim-verified) [crafted] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 49.7 attack_power points (4.07 DPS) | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Prowler's Leather Boots (252468, -0.62 DPS) [crafted]; Shadowcraft Boots (16711, -1.52 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 20.5 attack_power points (1.68 DPS) | yes | Mark of Kern (2262, -0.04 DPS) [dungeon]; Blackstone Ring (17713, -0.04 DPS) [dungeon]; Ring of Living Stone (18400, -0.04 DPS) [quest] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.4 attack_power points (1.67 DPS) | yes | Mark of Kern (2262, -0.03 DPS) [dungeon]; Blackstone Ring (17713, -0.03 DPS) [dungeon]; Ring of Living Stone (18400, -0.03 DPS) [quest] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smoking Heart of the Mountain (11811, +0.00 DPS) [crafted]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Arcanite Champion (12790) | Blacksmithing [crafted] | sim-verified (235.5 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Forest Defender's Axe (279260, -2.77 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skull Splitting Crossbow (13039, -0.04 DPS) [world_drop]; Houndmaster's Bow (11628, -0.20 DPS) [dungeon]; Dark Iron Rifle (16004, -1.26 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; shoulder: Lightforge Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Tranquil Wristguards; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Scalegut Treaders; finger1: Masons Fraternity Ring; finger2: Thunderbrow Ring; trinket2: Mark of the Chosen; main_hand: Arcanite Champion; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 2015, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60, raid preset (human, 02325213032515201-05050020000000000-200000000000000000)

Set DPS (verified): 719.3. Weights run: 1.6s. Verify run: 6.4s. 2015 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.689 ± 0.118, crit=2.129 ± 0.161 per rating point (14 rating = 1%, 29.805 per %), hit=5.097 ± 0.436 per rating point (10 rating = 1%, 50.970 per %), melee_haste=18.897 ± 3.662

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 201.1 attack_power points (34.15 DPS) | yes | Lieutenant Commander's Plate Helm (23314, -12.59 DPS) [vendor]; Stalwart Helm (250599, -13.13 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 33.0 attack_power points (5.60 DPS) | yes | Sunprism Pendant (20649, -2.15 DPS) [quest]; Kaleidoscope Chain (13084, -2.95 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -3.56 DPS) [world] |
| shoulder | Lightforge Spaulders (16729) | Blackrock Spire: The Beast [dungeon] | 74.1 attack_power points (12.58 DPS) | yes | Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp]; Lieutenant Commander's Plate Shoulders (227045, -1.17 DPS) [pvp]; Dark Warder's Pauldrons (22241, -5.33 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 45.3 attack_power points (7.70 DPS) | yes | Sergeant Major's Cape (16337, -2.42 DPS) [pvp]; Fel Cape (279269, -2.64 DPS) [crafted]; Dark Phantom Cape (13122, -3.40 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 143.3 attack_power points (24.33 DPS) | yes | Timbermaw Tunic (252484, -1.78 DPS) [crafted]; Field Marshal's Plate Armor (231536, -9.27 DPS) [pvp]; Shadowcraft Tunic (16721, -9.95 DPS) [dungeon] |
| wrist | Tranquil Wristguards (279256) | Leatherworking [crafted] | 35.2 attack_power points (5.98 DPS) | yes | Marshal's Plate Bracers (16481, +0.00 DPS) [pvp]; Officer's Wristguards (250581, -0.15 DPS) [crafted]; Deepfury Bracers (13120, -0.18 DPS) [world_drop] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-verified (719.3 DPS) | yes | Marshal's Plate Gauntlets (231541, +0.00 DPS) [pvp]; Storm Gauntlets (12632, -1.16 DPS) [crafted]; Timbermaw Brawlers (19049, -21.96 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 69.9 attack_power points (11.86 DPS) | yes | Marshal's Plate Girdle (16482, +0.00 DPS) [pvp]; Knight-Captain's Plate Girdle (16407, -3.04 DPS) [pvp]; Prowler's Leather Waistguard (252473, -3.19 DPS) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (719.3 DPS) | yes | Knight-Captain's Plate Leggings (23301, +0.00 DPS) [vendor]; Marshal's Plate Legguards (231540, +0.00 DPS) [pvp]; Warbear Woolies (15065, -21.38 DPS, sim-verified) [crafted] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 51.5 attack_power points (8.75 DPS) | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Shadowcraft Boots (16711, -0.13 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.74 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 23.7 attack_power points (4.02 DPS) | yes | Ring of Living Stone (18400, -0.28 DPS) [quest]; Falcon's Hook (7552, -0.31 DPS) [dungeon]; Ring of the Underwood (2951, -0.40 DPS) [world_drop] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 22.7 attack_power points (3.85 DPS) | yes | Ring of Living Stone (18400, -0.11 DPS) [quest]; Falcon's Hook (7552, -0.15 DPS) [dungeon]; Ring of the Underwood (2951, -0.23 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (719.3 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (719.3 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Arcanite Champion (12790) | Blacksmithing [crafted] | sim-verified (719.3 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Forest Defender's Axe (279260, -30.15 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (719.3 DPS) | yes | Skull Splitting Crossbow (13039, -0.52 DPS) [world_drop]; Guttbuster (13139, -0.60 DPS) [world_drop]; Dark Iron Rifle (16004, -4.36 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; shoulder: Lightforge Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Tranquil Wristguards; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Scalegut Treaders; finger1: Masons Fraternity Ring; finger2: Thunderbrow Ring; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Arcanite Champion; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 2015, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band

## Horde

### Band 20 (orc, 03323000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 25.5. Weights run: 1.5s. Verify run: 2.5s. 276 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.209 ± 0.030, crit=0.828 ± 0.027 per rating point (14 rating = 1%, 11.592 per %), hit=1.232 ± 0.083 per rating point (10 rating = 1%, 12.321 per %), melee_haste=8.730 ± 0.630

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.71 DPS) | yes | Defender's Leather Hood (252447, -0.15 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.28 DPS) [crafted]; Brawler's Leather Hood (252504, -0.37 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 2.4 attack_power points (0.09 DPS) | yes | Roadwatcher's Confidence (281265, +0.00 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 6.0 attack_power points (0.21 DPS) | yes | Rough Bronze Shoulders (3480, -0.00 DPS) [crafted]; Silvered Bronze Shoulders (3481, -0.00 DPS) [crafted] |
| back | Grave Shroud (279865) | Unending Torment [quest] | 8.4 attack_power points (0.30 DPS) | yes | Lambent Scale Cloak (4706, -0.01 DPS) [world_drop]; Dark Leather Cloak (2316, -0.03 DPS) [crafted]; Glowing Lizardscale Cloak (6449, -0.04 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (25.5 DPS) | yes | Veteran's Chain Shirt (250488, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, +0.00 DPS) [crafted]; Mutant Scale Breastplate (6627, -0.78 DPS, sim-verified) [dungeon] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bristlebark Bindings (14569, -0.08 DPS) [world_drop]; Forest Leather Bracers (3202, -0.14 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (25.5 DPS) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Thorbia's Gauntlets (12994, -0.94 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.64 DPS) | yes | Brawler's Leather Belt (252428, -0.18 DPS) [crafted]; Deviate Scale Belt (6468, -0.21 DPS) [crafted]; Cobrahn's Grasp (6460, -1.09 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (25.5 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Veteran's Chain Leggings (250493, -0.90 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 16.0 attack_power points (0.57 DPS) | yes | Feet of the Lynx (1121, -0.01 DPS) [world_drop]; Veteran's Boots (250503, -0.04 DPS) [crafted]; Defender's Leather Boots (252441, -0.21 DPS) [crafted] |
| finger1 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Bounty Hunter's Ring (5351, -0.15 DPS) [quest]; The 1 Ring (8350, -0.17 DPS) [world]; Lavishly Jeweled Ring (1156, -0.20 DPS) [dungeon] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 6.0 attack_power points (0.21 DPS) | yes | Bounty Hunter's Ring (5351, -0.08 DPS) [quest]; The 1 Ring (8350, -0.10 DPS) [world]; Lavishly Jeweled Ring (1156, -0.13 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hammerbone (270018) | Leaders of the Fang [quest] | 306.2 attack_power points (10.83 DPS) | yes | Forsaken Greataxe (251533, +0.00 DPS, sim-verified) [quest]; Smite's Mighty Hammer (7230, -0.12 DPS) [dungeon]; Living Root (6631, -0.33 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.8 attack_power points (0.17 DPS) | yes | Fine Longbow (11304, -0.03 DPS) [vendor]; Cracked Blacksmith Hammer (285279, -0.03 DPS) [crafted]; Deadly Blunderbuss (4369, -0.09 DPS) [crafted] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Grave Shroud; chest: Blackened Defias Armor; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Brawler's Leather Boots; finger1: Demon Band; finger2: Loop of Sacrifice; main_hand: Hammerbone; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 276, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 15402 Noosegrip Gauntlets; 18854 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring

### Band 30 (orc, 03325213020000000-00000000000000000-000000000000000000)

Set DPS (verified): 72.4. Weights run: 1.6s. Verify run: 1.7s. 534 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=2.138 ± 0.053, crit=1.699 ± 0.057 per rating point (14 rating = 1%, 23.793 per %), hit=1.891 ± 0.186 per rating point (10 rating = 1%, 18.912 per %), melee_haste=10.934 ± 1.331

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 37.2 attack_power points (1.43 DPS) | yes | Tusken Helm (6686, -0.41 DPS, sim-verified) [dungeon]; Guard's Chain Helm (250499, -0.45 DPS) [crafted]; Slayer's Skullcap (14753, -0.49 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 16.6 attack_power points (0.64 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.18 DPS) [world]; River Pride Choker (13087, -0.33 DPS) [world_drop]; Erudite's Amulet (277204, -0.47 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 33.5 attack_power points (1.29 DPS) | yes | Barbaric Iron Shoulders (7913, -0.40 DPS, sim-verified) [crafted]; Mantle of Thieves (2264, -0.47 DPS) [dungeon]; Barbaric Shoulders (5964, -0.49 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 21.0 attack_power points (0.81 DPS) | yes | Fenrus' Hide (6340, -0.31 DPS) [dungeon]; Swiftrunner Cape (6745, -0.31 DPS) [quest]; Tigerstrike Mantle (13108, -0.33 DPS, sim-verified) [world_drop] |
| chest | Veteran's Silvered Chain Shirt (250518) | Blacksmithing [crafted] | 32.8 attack_power points (1.26 DPS) | yes | Avenger's Armor (1488, -0.11 DPS) [dungeon]; Dusky Leather Armor (7374, -0.11 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.14 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 20.8 attack_power points (0.80 DPS) | yes | Yorgen Bracers (13012, -0.09 DPS) [world_drop]; Barbaric Bracers (18948, -0.16 DPS) [crafted]; Pugilist Bracers (4438, -0.19 DPS) [dungeon] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 26.8 attack_power points (1.03 DPS) | yes | Toughened Leather Gloves (4253, -0.08 DPS) [crafted]; Mail Combat Gauntlets (4075, -0.09 DPS) [world_drop]; Fletcher's Gloves (7348, -0.12 DPS) [crafted] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 32.8 attack_power points (1.26 DPS) | yes | Skulker's Leather Belt (252520, -0.06 DPS) [crafted]; Prowler's Leather Belt (252459, -0.08 DPS) [crafted]; Warden's Leather Belt (252460, -0.31 DPS) [crafted] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 37.0 attack_power points (1.42 DPS) | yes | Brawler's Leather Pants (252500, -0.22 DPS) [crafted]; Trapper's Leather Pants (252501, -0.22 DPS) [crafted]; Brawler's Leather Legguards (252516, -0.39 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 23.1 attack_power points (0.89 DPS) | yes | Brawler's Leather Boots (252439, -0.09 DPS) [crafted]; Veteran's Boots (250503, -0.18 DPS) [crafted]; Insignia Boots (4055, -0.23 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 22.4 attack_power points (0.86 DPS) | yes | Ring of Precision (1491, -0.37 DPS) [dungeon]; Defias Renegade Ring (1076, -0.45 DPS) [dungeon]; Monkey Ring (6748, -0.45 DPS) [quest] |
| finger2 | Band of the Fist (17694) | Allegiance to the Old Gods [quest] | 16.6 attack_power points (0.64 DPS) | yes | Defias Renegade Ring (1076, -0.23 DPS) [dungeon]; Monkey Ring (6748, -0.23 DPS) [quest]; Ring of Precision (1491, -0.39 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (72.4 DPS) | yes | Corpsemaker (6687, -0.52 DPS) [dungeon]; Felbough Arm (273037, -0.91 DPS) [world]; Viscous Hammer (13045, -20.41 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 12.8 attack_power points (0.49 DPS) | yes | Double-barreled Shotgun (2098, -0.02 DPS) [world_drop]; BKP "Sparrow" Smallbore (3042, -0.16 DPS) [world_drop]; Silver Star (3463, -0.16 DPS) [quest] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Veteran's Silvered Chain Shirt; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Feet of the Lynx; finger1: Thunderbrow Ring; finger2: Band of the Fist; main_hand: Manual Crowd Pummeler; ranged: Golemsight Long Gun

No-known-source sample (15 of 534, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads

### Band 40 (orc, 03325213032511000-00000000000000000-000000000000000000)

Set DPS (verified): 93.6. Weights run: 1.8s. Verify run: 1.7s. 761 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.005, agility=1.924 ± 0.047, crit=1.887 ± 0.057 per rating point (14 rating = 1%, 26.420 per %), hit=2.156 ± 0.180 per rating point (10 rating = 1%, 21.561 per %), melee_haste=13.332 ± 1.509

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chromite Barbute (8142) | Uldaman: Ancient Treasure [dungeon] | 47.1 attack_power points (2.41 DPS) | yes | White Bandit Mask (10008, -0.20 DPS) [crafted]; Barbaric Iron Helm (7915, -0.60 DPS) [crafted]; Hawkeye's Helm (14591, -0.92 DPS) [world_drop] |
| neck | Ethereal Talisman (4430) | The Crown of Will [quest] | 17.7 attack_power points (0.91 DPS) | yes | Kaleidoscope Chain (13084, -0.10 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.29 DPS) [world]; River Pride Choker (13087, -0.50 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 31.2 attack_power points (1.59 DPS) | yes | Flintrock Shoulders (7755, -0.10 DPS) [dungeon]; Barbaric Iron Shoulders (7913, -0.39 DPS) [crafted]; Hard Gold Pauldrons (250539, -0.47 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 19.5 attack_power points (1.00 DPS) | yes | Imperial Cloak (6432, -0.21 DPS) [dungeon]; Parachute Cloak (10518, -0.21 DPS) [crafted]; Tigerstrike Mantle (13108, -0.21 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 46.2 attack_power points (2.36 DPS) | yes | Kolkar Marauder Chain (6773, -0.24 DPS) [quest]; Carapace of Tuten'kash (10775, -0.55 DPS) [dungeon]; Wolffear Harness (13110, -0.69 DPS) [world_drop] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 21.7 attack_power points (1.11 DPS) | yes | Branded Leather Bracers (19508, -0.09 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.11 DPS) [world_drop]; Yorgen Bracers (13012, -0.20 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 46.4 attack_power points (2.38 DPS) | yes | Scarlet Gauntlets (10331, -0.16 DPS) [dungeon]; Plated Fist of Hakoo (13071, -0.37 DPS) [world_drop]; Dragonscale Gauntlets (8347, -0.43 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 35.3 attack_power points (1.81 DPS) | yes | Officer's Belt (250556, -0.19 DPS) [crafted]; Prowler's Leather Belt (252459, -0.30 DPS) [crafted]; Skulker's Leather Belt (252520, -0.31 DPS) [crafted] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.15 DPS) | yes | Symbolic Legplates (14829, -0.02 DPS) [world_drop]; Triprunner Dungarees (9624, -0.07 DPS) [quest]; Basilisk Hide Pants (1718, -0.08 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 39.5 attack_power points (2.02 DPS) | yes | Blackforge Greaves (6423, -0.11 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.20 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.22 DPS) [crafted] |
| finger1 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 23.3 attack_power points (1.19 DPS) | yes | Thunderbrow Ring (13097, -0.08 DPS) [world_drop]; Mark of Kern (2262, -0.17 DPS) [dungeon]; Assault Band (13095, -0.17 DPS) [world_drop] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 23.2 attack_power points (1.19 DPS) | yes | Thunderbrow Ring (13097, -0.08 DPS) [world_drop]; Mark of Kern (2262, -0.17 DPS) [dungeon]; Assault Band (13095, -0.17 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Nightblade (1982) | World drop [world_drop] | sim-verified (93.6 DPS) | yes | Primitive Fishing Pole (276203, -1.05 DPS) [vendor]; Staff of Jordan (873, -2.14 DPS) [world_drop]; Fiery War Axe (870, -3.50 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Monolithic Bow (9426) | Uldaman [dungeon] | sim-verified (93.6 DPS) | yes | Master Hunter's Rifle (17687, -0.11 DPS) [quest]; The Silencer (13138, -0.19 DPS) [world_drop]; Bow of Searing Arrows (2825, -0.92 DPS, sim-verified) [world_drop] |

**New at 40:** head: Chromite Barbute; neck: Ethereal Talisman; chest: Icemail Jerkin; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Falcon's Hook; finger2: Ring of the Underwood; main_hand: Nightblade; ranged: Monolithic Bow

No-known-source sample (15 of 761, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 50 (orc, 03325213032515001-05000000000000000-000000000000000000)

Set DPS (verified): 179.9. Weights run: 1.8s. Verify run: 2.2s. 983 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.140 ± 0.059, crit=1.339 ± 0.079 per rating point (14 rating = 1%, 18.740 per %), hit=2.101 ± 0.194 per rating point (10 rating = 1%, 21.014 per %), melee_haste=12.996 ± 1.717

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sunscale Helmet (14849) | World drop [world_drop] | 38.8 attack_power points (3.13 DPS) | yes | Blood Guard's Plate Helm (220803, +0.00 DPS) [vendor]; Ornate Mithril Helm (7937, -0.01 DPS) [crafted]; Chromite Barbute (8142, -0.09 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 24.8 attack_power points (2.00 DPS) | yes | Woven Ivy Necklace (19159, -0.21 DPS) [quest]; Ethereal Talisman (4430, -0.83 DPS) [quest]; Kaleidoscope Chain (13084, -0.99 DPS) [world_drop] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 35.1 attack_power points (2.83 DPS) | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.16 DPS) [world_drop]; Prowler's Leather Shoulders (252534, -0.23 DPS) [crafted] |
| back | Bloodlust Cape (14801) | World drop [world_drop] | 18.0 attack_power points (1.45 DPS) | yes | Dark Phantom Cape (13122, -0.07 DPS) [world_drop]; Pridelord Cape (14673, -0.14 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.32 DPS) [world_drop] |
| chest | Stone Guard's Plate Armor (220801) | PvP rank 10 · Stone Guard · Horde [vendor] | 48.7 attack_power points (3.93 DPS) | yes | Warforged Chestplate (11195, -0.06 DPS) [quest]; Warbear Harness (15064, -0.50 DPS) [crafted]; Valorous Chestguard (8274, -0.70 DPS) [world_drop] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 28.8 attack_power points (2.32 DPS) | yes | Bracers of the Stone Princess (17714, -0.07 DPS) [dungeon]; Arena Bands (18711, -0.07 DPS) [world]; Prowler's Leather Bracers (252539, -0.23 DPS) [crafted] |
| hands | Officer's Gloves (250551) | Blacksmithing [crafted] | 42.3 attack_power points (3.40 DPS) | yes | Stonerender Gauntlets (17007, -0.18 DPS) [world_drop]; Gloves of Holy Might (867, -0.28 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, -0.32 DPS) [crafted] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 41.7 attack_power points (3.36 DPS) | yes | Skulker's Leather Waistguard (252474, -0.14 DPS) [crafted]; Officer's Waistguard (250571, -0.18 DPS) [crafted]; Belt of the Gladiator (13134, -0.46 DPS) [world_drop] |
| legs | Stone Guard's Plate Leggings (220798) | PvP rank 10 · Stone Guard · Horde [vendor] | 54.1 attack_power points (4.36 DPS) | yes | Serpentskin Leggings (8262, -0.80 DPS) [world_drop]; Golem Shard Leggings (13074, -0.82 DPS) [world_drop]; Scarlet Leggings (10330, -0.98 DPS) [dungeon] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 38.5 attack_power points (3.10 DPS) | yes | Skulker's Leather Boots (252469, -0.14 DPS) [crafted]; Officer's Sabatons (250561, -0.18 DPS) [crafted]; Officer's Boots (250546, -0.37 DPS) [crafted] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.93 DPS) | yes | Mark of Kern (2262, -0.32 DPS) [dungeon]; Assault Band (13095, -0.32 DPS) [world_drop]; Thunderbrow Ring (13097, -0.37 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) (or Mark of Kern (2262), Assault Band (13095)) | Maraudon: Princess Theradras [dungeon] | 20.0 attack_power points (1.61 DPS) | yes | Mark of Kern (2262, +0.00 DPS) [dungeon]; Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.05 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (179.9 DPS) | yes | Molten Heart of the Mountain (249470, -4.60 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (179.9 DPS) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Blight (7959) | Blacksmithing [crafted] | sim-verified (179.9 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Nightblade (1982, -1.71 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (179.9 DPS) | yes | Houndmaster's Bow (11628, -0.16 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.33 DPS) [world_drop]; Dark Iron Rifle (16004, -1.80 DPS, sim-verified) [crafted] |

**New at 50:** head: Sunscale Helmet; neck: Skibi's Pendant; shoulder: Officer's Pauldrons; back: Bloodlust Cape; chest: Stone Guard's Plate Armor; wrist: Officer's Wristguards; hands: Officer's Gloves; waist: Prowler's Leather Waistguard; legs: Stone Guard's Plate Leggings; feet: Prowler's Leather Boots; finger1: White Bone Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Blight; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 983, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (orc, 03325213032515001-05050000000000000-005000000000000000)

Set DPS (verified): 243.1. Weights run: 1.9s. Verify run: 6.9s. 1998 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.462 ± 0.068, crit=2.089 ± 0.097 per rating point (14 rating = 1%, 29.240 per %), hit=3.886 ± 0.268 per rating point (10 rating = 1%, 38.862 per %), melee_haste=19.865 ± 2.553

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 172.2 attack_power points (14.12 DPS) | yes | Champion's Plate Helm (227043, -5.09 DPS) [pvp]; Stalwart Helm (250599, -6.09 DPS, sim-verified) [crafted]; Blood Guard's Plate Helm (220803, -6.40 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 29.0 attack_power points (2.38 DPS) | yes | Woven Ivy Necklace (19159, -0.32 DPS) [quest]; Sunprism Pendant (20649, -0.94 DPS) [quest]; Ethereal Talisman (4430, -1.08 DPS) [quest] |
| shoulder | Lightforge Spaulders (16729) | Blackrock Spire: The Beast [dungeon] | 60.3 attack_power points (4.94 DPS) | yes | Champion's Plate Shoulders (227042, +0.00 DPS) [pvp]; Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp]; Dark Warder's Pauldrons (22241, -2.72 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 41.9 attack_power points (3.44 DPS) | yes | Fel Cape (279269, -1.51 DPS, sim-verified) [crafted]; Dark Phantom Cape (13122, -1.64 DPS) [world_drop]; Pridelord Cape (14673, -1.87 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 122.2 attack_power points (10.02 DPS) | yes | Timbermaw Tunic (252484, -0.05 DPS) [crafted]; Warlord's Plate Armor (231530, -3.32 DPS) [pvp]; Legionnaire's Plate Hauberk (227049, -4.18 DPS) [pvp] |
| wrist | Tranquil Wristguards (279256) | Leatherworking [crafted] | 32.0 attack_power points (2.62 DPS) | yes | General's Plate Armguards (16546, +0.00 DPS) [pvp]; Officer's Wristguards (250581, -0.10 DPS) [crafted]; Deepfury Bracers (13120, -0.17 DPS) [world_drop] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-verified (243.1 DPS) | yes | General's Plate Gauntlets (231532, +0.00 DPS) [pvp]; Gloves of Holy Might (867, -0.66 DPS) [world_drop]; Timbermaw Brawlers (19049, -6.33 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 62.5 attack_power points (5.12 DPS) | yes | General's Plate Girdle (16547, +0.00 DPS) [pvp]; Legionnaire's Plate Cinch (16511, -1.29 DPS) [pvp]; Prowler's Leather Waistguard (252473, -1.52 DPS, sim-verified) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (243.1 DPS) | yes | Legionnaire's Plate Leggings (227048, +0.00 DPS) [pvp]; General's Plate Leggings (231533, +0.00 DPS) [pvp]; Warbear Woolies (15065, -6.05 DPS, sim-verified) [crafted] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 49.7 attack_power points (4.07 DPS) | yes | General's Plate Boots (231531, +0.00 DPS) [pvp]; Prowler's Leather Boots (252468, -0.62 DPS) [crafted]; Shadowcraft Boots (16711, -2.29 DPS, sim-verified) [dungeon] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.97 DPS) | yes | Thunderbrow Ring (13097, -0.30 DPS) [world_drop]; Blackstone Ring (17713, -0.33 DPS) [dungeon]; Ring of Living Stone (18400, -0.33 DPS) [quest] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 20.5 attack_power points (1.68 DPS) | yes | Thunderbrow Ring (13097, -0.01 DPS) [world_drop]; Blackstone Ring (17713, -0.04 DPS) [dungeon]; Ring of Living Stone (18400, -0.04 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (243.1 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (243.1 DPS) | yes | Molten Heart of the Mountain (249470, -3.80 DPS, sim-verified) [crafted] |
| main_hand | Arcanite Champion (12790) | Blacksmithing [crafted] | sim-verified (243.1 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Forest Defender's Axe (279260, -2.65 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (243.1 DPS) | yes | Skull Splitting Crossbow (13039, -0.04 DPS) [world_drop]; Houndmaster's Bow (11628, -0.20 DPS) [dungeon]; Dark Iron Rifle (16004, -1.70 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; shoulder: Lightforge Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Tranquil Wristguards; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Scalegut Treaders; finger2: Masons Fraternity Ring; main_hand: Arcanite Champion; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1998, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60, raid preset (orc, 02325213032515201-05050020000000000-200000000000000000)

Set DPS (verified): 733.1. Weights run: 1.6s. Verify run: 6.5s. 1998 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.689 ± 0.118, crit=2.129 ± 0.161 per rating point (14 rating = 1%, 29.805 per %), hit=5.097 ± 0.436 per rating point (10 rating = 1%, 50.970 per %), melee_haste=18.897 ± 3.662

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 201.1 attack_power points (34.15 DPS) | yes | Champion's Plate Helm (227043, -12.59 DPS) [pvp]; Stalwart Helm (250599, -14.95 DPS, sim-verified) [crafted]; Blood Guard's Plate Helm (220803, -15.58 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 33.0 attack_power points (5.60 DPS) | yes | Woven Ivy Necklace (19159, -0.77 DPS) [quest]; Sunprism Pendant (20649, -2.15 DPS) [quest]; Ethereal Talisman (4430, -2.58 DPS) [quest] |
| shoulder | Lightforge Spaulders (16729) | Blackrock Spire: The Beast [dungeon] | 74.1 attack_power points (12.58 DPS) | yes | Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp]; Champion's Plate Shoulders (227042, -1.17 DPS) [pvp]; Dark Warder's Pauldrons (22241, -7.01 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 45.3 attack_power points (7.70 DPS) | yes | Dark Phantom Cape (13122, -3.40 DPS) [world_drop]; Pridelord Cape (14673, -4.00 DPS) [world_drop]; Fel Cape (279269, -4.40 DPS, sim-verified) [crafted] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 143.3 attack_power points (24.33 DPS) | yes | Timbermaw Tunic (252484, -3.81 DPS, sim-verified) [crafted]; Warlord's Plate Armor (231530, -9.27 DPS) [pvp]; Shadowcraft Tunic (16721, -9.95 DPS) [dungeon] |
| wrist | Tranquil Wristguards (279256) | Leatherworking [crafted] | 35.2 attack_power points (5.98 DPS) | yes | General's Plate Armguards (16546, +0.00 DPS) [pvp]; Officer's Wristguards (250581, -0.15 DPS) [crafted]; Deepfury Bracers (13120, -0.18 DPS) [world_drop] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-verified (733.1 DPS) | yes | General's Plate Gauntlets (231532, +0.00 DPS) [pvp]; Storm Gauntlets (12632, -1.16 DPS) [crafted]; Timbermaw Brawlers (19049, -23.24 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 69.9 attack_power points (11.86 DPS) | yes | General's Plate Girdle (16547, +0.00 DPS) [pvp]; Legionnaire's Plate Cinch (16511, -3.04 DPS) [pvp]; Prowler's Leather Waistguard (252473, -3.71 DPS, sim-verified) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (733.1 DPS) | yes | Legionnaire's Plate Leggings (227048, +0.00 DPS) [pvp]; General's Plate Leggings (231533, +0.00 DPS) [pvp]; Warbear Woolies (15065, -21.68 DPS, sim-verified) [crafted] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 51.5 attack_power points (8.75 DPS) | yes | General's Plate Boots (231531, +0.00 DPS) [pvp]; Shadowcraft Boots (16711, -0.13 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.74 DPS) [crafted] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (4.07 DPS) | yes | Thunderbrow Ring (13097, -0.23 DPS) [world_drop]; Ring of Living Stone (18400, -0.34 DPS) [quest]; Falcon's Hook (7552, -0.37 DPS) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 23.7 attack_power points (4.02 DPS) | yes | Thunderbrow Ring (13097, -0.17 DPS) [world_drop]; Ring of Living Stone (18400, -0.28 DPS) [quest]; Falcon's Hook (7552, -0.31 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (733.1 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (733.1 DPS) | yes | Molten Heart of the Mountain (249470, -10.13 DPS, sim-verified) [crafted] |
| main_hand | Arcanite Champion (12790) | Blacksmithing [crafted] | sim-verified (733.1 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Forest Defender's Axe (279260, -31.85 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (733.1 DPS) | yes | Skull Splitting Crossbow (13039, -0.52 DPS) [world_drop]; Guttbuster (13139, -0.60 DPS) [world_drop]; Dark Iron Rifle (16004, -4.01 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; shoulder: Lightforge Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Tranquil Wristguards; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Scalegut Treaders; finger2: Masons Fraternity Ring; main_hand: Arcanite Champion; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1998, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

