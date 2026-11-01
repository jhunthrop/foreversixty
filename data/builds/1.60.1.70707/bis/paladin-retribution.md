# Leveling BiS: Retribution

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-00000000000000000-05231000000000000)

Set DPS (verified): 39.8. Weights run: 1.4s. Verify run: 2.4s. 232 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.491 ± 0.026, crit=0.844 ± 0.024 per rating point (14 rating = 1%, 11.816 per %), hit=1.363 ± 0.065 per rating point (10 rating = 1%, 13.628 per %), melee_haste=8.675 ± 0.564

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.87 DPS) | yes | Defender's Leather Hood (252447, -0.17 DPS) [crafted]; Guard's Silvered Chain Helm (250529, -0.22 DPS) [crafted]; Brawler's Leather Hood (252504, -0.35 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 3.0 attack_power points (0.13 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 7.5 attack_power points (0.32 DPS) | yes | Rough Bronze Shoulders (3480, -0.06 DPS) [crafted]; Silvered Bronze Shoulders (3481, -0.06 DPS) [crafted] |
| back | Grave Shroud (279865) | Abominable Creatures [quest] | 9.0 attack_power points (0.39 DPS) | yes | Glowing Lizardscale Cloak (6449, -0.00 DPS) [dungeon]; Dark Leather Cloak (2316, -0.02 DPS) [crafted]; Lambent Scale Cloak (4706, -0.04 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (39.8 DPS) | yes | Mutant Scale Breastplate (6627, +0.00 DPS) [dungeon]; Veteran's Chain Shirt (250488, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -1.81 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.43 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bravo's Armbands (270015, -0.00 DPS) [quest]; Bristlebark Bindings (14569, -0.07 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (39.8 DPS) | yes | Dagmire Gauntlets (6481, +0.00 DPS) [quest]; Thorbia's Gauntlets (12994, +0.00 DPS) [world_drop]; Gloves of the Fang (10413, -2.09 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (39.8 DPS) | yes | Brawler's Leather Belt (252428, -0.18 DPS) [crafted]; Deviate Scale Belt (6468, -0.20 DPS) [crafted]; Cobrahn's Grasp (6460, -2.23 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (39.8 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Veteran's Chain Leggings (250493, -1.96 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 17.9 attack_power points (0.78 DPS) | yes | Brawler's Leather Boots (252439, -0.02 DPS) [crafted]; Veteran's Boots (250503, -0.09 DPS) [crafted]; Defender's Leather Boots (252441, -0.34 DPS) [crafted] |
| finger1 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.35 DPS) | yes | The 1 Ring (8350, -0.20 DPS) [world]; Ring of the Moon (12052, -0.26 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | sim-verified (39.8 DPS) | yes | Ring of the Moon (12052, -0.04 DPS) [world_drop]; The 1 Ring (8350, -0.48 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Verigan's Fist (6953) | The Test of Righteousness [quest] | 358.7 attack_power points (15.60 DPS) | yes | Duskbringer (2205, -2.79 DPS) [dungeon]; Sword of the Fallen (286977, -3.08 DPS) [dungeon]; Smite's Mighty Hammer (7230, -3.31 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Grave Shroud; chest: Blackened Defias Armor; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Demon Band; finger2: Lavishly Jeweled Ring; main_hand: Verigan's Fist

No-known-source sample (15 of 232, see the JSON for more): 1189 Overseer's Ring; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6478 Rat Stompers; 9602 Brushwood Blade; 14145 Cursed Felblade; 14147 Cavedweller Bracers; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 30 (human, 00000000000000000-00000000000000000-05233301001300000)

Set DPS (verified): 82.7. Weights run: 1.3s. Verify run: 2.5s. 454 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.267 ± 0.023, crit=0.964 ± 0.025 per rating point (14 rating = 1%, 13.497 per %), hit=1.585 ± 0.069 per rating point (10 rating = 1%, 15.846 per %), melee_haste=9.455 ± 0.568

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 29.4 attack_power points (1.42 DPS) | yes | Tusken Helm (6686, -0.16 DPS) [dungeon]; Veteran's Chain Helm (250498, -0.26 DPS) [crafted]; Defender's Leather Helm (252455, -0.26 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 13.1 attack_power points (0.63 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.05 DPS) [world]; River Pride Choker (13087, -0.24 DPS) [world_drop]; Fallen Guard's Pendant (279837, -0.34 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.9 attack_power points (1.15 DPS) | yes | Barbaric Iron Shoulders (7913, -0.21 DPS) [crafted]; Barbaric Shoulders (5964, -0.37 DPS) [crafted]; Golden Scale Shoulders (3841, -0.48 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 14.9 attack_power points (0.72 DPS) | yes | Sergeant Major's Cape (16315, -0.09 DPS) [pvp]; Tigerstrike Mantle (13108, -0.23 DPS) [world_drop]; Wolfmaster Cape (6314, -0.23 DPS) [dungeon] |
| chest | Avenger's Armor (1488) | Razorfen Kraul [dungeon] | 30.0 attack_power points (1.45 DPS) | yes | Shining Silver Breastplate (2870, -0.10 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.12 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.29 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul [dungeon] | 16.0 attack_power points (0.77 DPS) | yes | Yorgen Bracers (13012, -0.01 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.02 DPS) [world_drop]; Barbaric Bracers (18948, -0.14 DPS) [crafted] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (82.7 DPS) | yes | Mail Combat Gauntlets (4075, +0.00 DPS) [world_drop]; Insignia Gloves (6408, +0.00 DPS) [world_drop]; Gauntlets of Ogre Strength (3341, -4.42 DPS, sim-verified) [world] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (82.7 DPS) | yes | Girdle of Golem Strength (9405, +0.00 DPS) [world_drop]; Prowler's Leather Belt (252459, +0.00 DPS) [crafted]; Officer's Belt (250556, -4.58 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (82.7 DPS) | yes | Ferine Leggings (6690, +0.00 DPS) [dungeon]; Brawler's Leather Legguards (252516, +0.00 DPS) [crafted]; Veteran's Silvered Chain Leggings (250523, -4.36 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (82.7 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Alacritous Treads (277234, -4.48 DPS, sim-verified) [quest] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.8 attack_power points (0.96 DPS) | yes | Silverlaine's Family Seal (6321, -0.47 DPS) [dungeon]; Bonebinder's Signet (273032, -0.47 DPS) [world]; Seal of Wrynn (2933, -0.48 DPS) [quest] |
| finger2 | Tiger Band (6749) (or Bonebinder's Signet (273032), Silverlaine's Family Seal (6321)) | Willix the Importer [quest] | 10.0 attack_power points (0.48 DPS) | yes | Silverlaine's Family Seal (6321, +0.00 DPS) [dungeon]; Bonebinder's Signet (273032, +0.00 DPS) [world]; Seal of Wrynn (2933, -0.01 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (82.7 DPS) | yes | Felbough Arm (273037, -1.35 DPS) [world]; Cobalt Crusher (7730, -1.91 DPS) [dungeon]; Viscous Hammer (13045, -20.36 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; feet: Blackened Defias Boots; finger1: Thunderbrow Ring; finger2: Tiger Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 454, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (human, 01000000000000000-00000000000000000-05233301001301350)

Set DPS (verified): 124.8. Weights run: 1.3s. Verify run: 2.5s. 643 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.040 ± 0.004, agility=1.541 ± 0.030, crit=1.490 ± 0.036 per rating point (14 rating = 1%, 20.865 per %), hit=2.006 ± 0.088 per rating point (10 rating = 1%, 20.059 per %), melee_haste=11.290 ± 0.657

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chromite Barbute (8142) | Uldaman: Ancient Treasure [dungeon] | 43.0 attack_power points (2.24 DPS) | yes | White Bandit Mask (10008, -0.19 DPS) [crafted]; Barbaric Iron Helm (7915, -0.56 DPS) [crafted]; Hard Gold Coif (250537, -0.75 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 14.3 attack_power points (0.75 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.12 DPS) [world]; River Pride Choker (13087, -0.32 DPS) [world_drop]; Gazlowe's Charm (13088, -0.32 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 27.2 attack_power points (1.41 DPS) | yes | Flintrock Shoulders (7755, -0.08 DPS) [dungeon]; Hard Gold Pauldrons (250539, -0.25 DPS) [crafted]; Barbaric Iron Shoulders (7913, -0.30 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 21.5 attack_power points (1.12 DPS) | yes | Yeti Fur Cloak (2805, -0.32 DPS) [quest]; Parachute Cloak (10518, -0.48 DPS) [crafted]; Hawkeye's Cloak (14593, -1.74 DPS, sim-verified) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-verified (124.8 DPS) | yes | Icemail Jerkin (1981, +0.00 DPS) [world_drop]; Carapace of Tuten'kash (10775, +0.00 DPS) [dungeon]; Kolkar Marauder Chain (6773, -4.85 DPS, sim-verified) [quest] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 20.4 attack_power points (1.06 DPS) | yes | Branded Leather Bracers (19508, -0.02 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.16 DPS) [world_drop]; Yorgen Bracers (13012, -0.19 DPS) [world_drop] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (124.8 DPS) | yes | Scarlet Gauntlets (10331, +0.00 DPS) [dungeon]; Plated Fist of Hakoo (13071, +0.00 DPS) [world_drop]; Gloves of Holy Might (867, -2.29 DPS, sim-verified) [world_drop] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 32.2 attack_power points (1.68 DPS) | yes | Officer's Belt (250556, -0.13 DPS) [crafted]; Prowler's Leather Belt (252459, -0.24 DPS) [crafted]; Skulker's Leather Belt (252520, -0.32 DPS) [crafted] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-verified (124.8 DPS) | yes | Scarlet Leggings (10330, +0.00 DPS) [dungeon]; Firemane Leggings (13129, +0.00 DPS) [world_drop]; Symbolic Legplates (14829, +0.00 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 37.3 attack_power points (1.94 DPS) | yes | Blackforge Greaves (6423, -0.18 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.21 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.32 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.9 attack_power points (1.09 DPS) | yes | Assault Band (13095, -0.05 DPS) [world_drop]; Falcon's Hook (7552, -0.05 DPS) [dungeon]; Ring of the Underwood (2951, -0.08 DPS) [world_drop] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.04 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Falcon's Hook (7552, -0.00 DPS) [dungeon]; Ring of the Underwood (2951, -0.03 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (124.8 DPS) | yes | Thornstone Sledgehammer (1722, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Nightblade (1982, -19.18 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Chromite Barbute; back: Sergeant Major's Cape; chest: Stormcloth Vest; wrist: Ravager's Armguards; hands: Stormcloth Gloves; waist: Ogron's Sash; legs: Stormcloth Pants; feet: Officer's Boots; finger2: Mark of Kern

No-known-source sample (15 of 643, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 50 (human, 05024000000000000-00000000000000000-05233301001301350)

Set DPS (verified): 151.7. Weights run: 1.4s. Verify run: 3.1s. 840 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.509 ± 0.029, crit=1.792 ± 0.039 per rating point (14 rating = 1%, 25.091 per %), hit=2.283 ± 0.095 per rating point (10 rating = 1%, 22.826 per %), melee_haste=11.006 ± 0.993

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ornate Mithril Helm (7937) | Blacksmithing [crafted] | 47.1 attack_power points (2.49 DPS) | yes | Knight-Lieutenant's Plate Helm (220804, +0.00 DPS) [vendor]; Chromite Barbute (8142, -0.14 DPS) [dungeon]; Sunscale Helmet (14849, -0.15 DPS) [world_drop] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 30.6 attack_power points (1.62 DPS) | yes | Kaleidoscope Chain (13084, -0.83 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.98 DPS) [world]; River Pride Choker (13087, -1.15 DPS) [world_drop] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 40.7 attack_power points (2.15 DPS) | yes | Knight-Lieutenant's Plate Pauldrons (220795, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.12 DPS) [world_drop]; Prowler's Leather Shoulders (252534, -0.15 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | sim-verified (151.7 DPS) | yes | Pridelord Cape (14673, -0.11 DPS) [world_drop]; Bloodlust Cape (14801, -0.13 DPS) [world_drop]; Dark Phantom Cape (13122, -2.43 DPS, sim-verified) [world_drop] |
| chest | Knight's Lamellar Chestplate (220815) | Captain Dirgehammer [vendor] | sim-verified (151.7 DPS) | yes | Warforged Chestplate (11195, +0.00 DPS) [quest]; Knight's Plate Hauberk (220794, +0.00 DPS) [vendor]; Knight's Imbued Armor (220813, -1.10 DPS, sim-verified) [vendor] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 33.3 attack_power points (1.76 DPS) | yes | Deepfury Bracers (13120, -0.10 DPS) [world_drop]; Prowler's Leather Bracers (252539, -0.15 DPS) [crafted]; Giantslayer Bracers (13076, -0.20 DPS) [world_drop] |
| hands | Sergeant Major's Lamellar Gauntlets (220817) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-verified (151.7 DPS) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; Sergeant Major's Imbued Gauntlets (220812, +0.00 DPS) [vendor]; Officer's Gloves (250551, +0.00 DPS) [crafted] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 48.9 attack_power points (2.59 DPS) | yes | Skulker's Leather Waistguard (252474, -0.07 DPS) [crafted]; Officer's Waistguard (250571, -0.16 DPS) [crafted]; Belt of the Gladiator (13134, -0.49 DPS) [world_drop] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | 66.6 attack_power points (3.52 DPS) | yes | Gryphon Rider's Leggings (9652, -0.58 DPS) [quest]; Centurion Legplates (10740, -0.58 DPS) [quest]; Serpentskin Leggings (8262, -0.73 DPS) [world_drop] |
| feet | Sergeant Major's Lamellar Boots (220814) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-verified (151.7 DPS) | yes | Sergeant Major's Imbued Greaves (220811, +0.00 DPS) [vendor]; Prowler's Leather Boots (252468, +0.00 DPS) [crafted]; Skulker's Leather Boots (252469, +0.00 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 22.1 attack_power points (1.17 DPS) | yes | Falcon's Hook (7552, -0.10 DPS) [dungeon]; Mark of Kern (2262, -0.11 DPS) [dungeon]; Blackstone Ring (17713, -0.11 DPS) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 21.1 attack_power points (1.12 DPS) | yes | Falcon's Hook (7552, -0.05 DPS) [dungeon]; Mark of Kern (2262, -0.06 DPS) [dungeon]; Blackstone Ring (17713, -0.06 DPS) [dungeon] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (151.7 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (151.7 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-verified (151.7 DPS) | yes | Warmonger (13052, -0.22 DPS) [world_drop]; Stoneraven (13059, -1.75 DPS) [world_drop]; Blight (7959, -2.30 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Ornate Mithril Helm; neck: Skibi's Pendant; shoulder: Officer's Pauldrons; chest: Knight's Lamellar Chestplate; wrist: Officer's Wristguards; hands: Sergeant Major's Lamellar Gauntlets; waist: Prowler's Leather Waistguard; legs: Knight's Plate Leggings; feet: Sergeant Major's Lamellar Boots; finger2: Masons Fraternity Ring; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain; main_hand: Thorium Greatmace

No-known-source sample (15 of 840, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (human, 05025000200000000-50200000000000000-05233301001301350)

Set DPS (verified): 250.3. Weights run: 1.4s. Verify run: 7.2s. 1762 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=2.015 ± 0.048, crit=2.840 ± 0.069 per rating point (14 rating = 1%, 39.761 per %), hit=4.097 ± 0.173 per rating point (10 rating = 1%, 40.969 per %), melee_haste=19.815 ± 1.945

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Stalwart Helm (250599) | Blacksmithing [crafted] | sim-verified (250.3 DPS) | yes | Knight-Lieutenant's Plate Helm (220804, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -0.35 DPS) [crafted]; Lionheart Helm (12640, -5.16 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 37.2 attack_power points (2.02 DPS) | yes | Sunprism Pendant (20649, -0.71 DPS) [quest]; Kaleidoscope Chain (13084, -1.10 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.37 DPS) [world] |
| shoulder | Magister's Mantle (16689) | Scholomance: Ras Frostwhisper [dungeon] | sim-verified (250.3 DPS) | yes | Ironfeather Shoulders (15067, +0.00 DPS, sim-verified) [crafted]; Lightforge Spaulders (16729, +0.00 DPS) [dungeon]; Knight-Lieutenant's Plate Pauldrons (220795, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 50.2 attack_power points (2.72 DPS) | yes | Fel Cape (279269, -0.57 DPS) [crafted]; Sergeant Major's Cape (16337, -0.89 DPS) [pvp]; Dark Phantom Cape (13122, -1.08 DPS) [world_drop] |
| chest | Magister's Robes (16688) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (250.3 DPS) | yes | Dawn Armor (252483, +0.00 DPS) [crafted]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Ironfeather Breastplate (15066, -0.79 DPS, sim-verified) [crafted] |
| wrist | Magister's Bindings (16683) | Blackrock Spire: Rage Talon Fire Tongue [dungeon] | sim-verified (250.3 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS) [world_drop]; Arena Wristguards (18709, +0.00 DPS) [world]; Officer's Wristguards (250581, +0.00 DPS) [crafted] |
| hands | Gauntlets of Elements (16672) | Blackrock Spire: Pyroguard Emberseer [dungeon] | sim-verified (250.3 DPS) | yes | Timbermaw Brawlers (19049, +0.00 DPS) [crafted]; Marshal's Lamellar Gauntlets (231650, +0.00 DPS) [vendor]; Magister's Gloves (16684, -3.82 DPS, sim-verified) [dungeon] |
| waist | Cord of Elements (16673) | Blackrock Spire: Smolderthorn Witch Doctor [dungeon] | sim-verified (250.3 DPS) | yes | Might of the Timbermaw (19044, +0.00 DPS) [crafted]; Prowler's Leather Waistguard (252473, +0.00 DPS) [crafted]; Magister's Belt (16685, -1.03 DPS, sim-verified) [dungeon] |
| legs | Magister's Leggings (16687) | Stratholme: Baron Rivendare [dungeon] | sim-verified (250.3 DPS) | yes | Knight-Captain's Lamellar Leggings (16435, +0.00 DPS) [pvp]; Azerothium Legplates (279265, +0.00 DPS) [crafted]; Warbear Woolies (15065, -3.02 DPS, sim-verified) [crafted] |
| feet | Boots of Elements (16670) | Blackrock Spire: Highlord Omokk [dungeon] | sim-verified (250.3 DPS) | yes | Shadowcraft Boots (16711, +0.00 DPS) [dungeon]; Knight-Lieutenant's Lamellar Sabatons (227146, +0.00 DPS) [vendor]; Deathbone Sabatons (14621, -6.29 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 28.2 attack_power points (1.53 DPS) | yes | Ring of the Underwood (2951, -0.20 DPS) [world_drop]; Thunderbrow Ring (13097, -0.25 DPS) [world_drop]; Ring of Living Stone (18400, -0.34 DPS) [quest] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 24.7 attack_power points (1.34 DPS) | yes | Ring of the Underwood (2951, -0.01 DPS) [world_drop]; Thunderbrow Ring (13097, -0.06 DPS) [world_drop]; Ring of Living Stone (18400, -0.15 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (250.3 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (250.3 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Blade of Hanna (2801) | World drop [world_drop] | sim-verified (250.3 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Arcanite Champion (12790, -7.50 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-verified (250.3 DPS) | yes | Libram of Fervor (23203, +0.00 DPS) [world_drop] |

**New at 60:** head: Stalwart Helm; shoulder: Magister's Mantle; back: Cape of the Black Baron; chest: Magister's Robes; wrist: Magister's Bindings; hands: Gauntlets of Elements; waist: Cord of Elements; legs: Magister's Leggings; feet: Boots of Elements; finger1: Masons Fraternity Ring; finger2: Falcon's Hook; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Blade of Hanna

No-known-source sample (15 of 1762, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60, raid preset (human, 05015000200000000-50300000000000000-05233301001301350)

Set DPS (verified): 546.4. Weights run: 1.6s. Verify run: 7.5s. 1762 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.420 ± 0.005, agility=2.013 ± 0.042, crit=2.593 ± 0.058 per rating point (14 rating = 1%, 36.300 per %), hit=4.456 ± 0.257 per rating point (10 rating = 1%, 44.557 per %), melee_haste=33.789 ± 2.838

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Stalwart Helm (250599) | Blacksmithing [crafted] | sim-verified (546.4 DPS) | yes | Knight-Lieutenant's Plate Helm (220804, -0.64 DPS) [vendor]; Outlaw's Collar (279253, -2.34 DPS) [crafted]; Lionheart Helm (12640, -5.43 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 38.3 attack_power points (4.18 DPS) | yes | Sunprism Pendant (20649, -1.54 DPS) [quest]; Kaleidoscope Chain (13084, -2.24 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -2.87 DPS) [world] |
| shoulder | Lightforge Spaulders (16729) | Blackrock Spire: The Beast [dungeon] | sim-verified (546.4 DPS) | yes | Knight-Lieutenant's Plate Pauldrons (220795, -0.99 DPS) [vendor]; Dark Warder's Pauldrons (22241, -1.21 DPS) [dungeon]; Magister's Mantle (16689, -6.24 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 50.2 attack_power points (5.48 DPS) | yes | Fel Cape (279269, -1.52 DPS) [crafted]; Sergeant Major's Cape (16337, -1.61 DPS) [pvp]; Dark Phantom Cape (13122, -2.18 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 155.3 attack_power points (16.95 DPS) | yes | Timbermaw Tunic (252484, -0.32 DPS) [crafted]; Field Marshal's Lamellar Breastplate (231649, -7.18 DPS) [vendor]; Shadowcraft Tunic (16721, -7.35 DPS) [dungeon] |
| wrist | Lightforge Bracers (16722) | Scholomance: Lord Alexei Barov [dungeon] | sim-verified (546.4 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS) [world_drop]; Tranquil Wristguards (279256, +0.00 DPS) [crafted]; Magister's Bindings (16683, -6.12 DPS, sim-verified) [dungeon] |
| hands | Timbermaw Brawlers (19049) | Leatherworking [crafted] | 79.0 attack_power points (8.62 DPS) | yes | Devilsaur Gauntlets (15063, +0.00 DPS) [crafted]; Marshal's Lamellar Gauntlets (231650, -0.17 DPS) [vendor]; Knight-Lieutenant's Lamellar Gauntlets (23274, -1.49 DPS) [vendor] |
| waist | Lightforge Belt (16723) | Stratholme: Rockwing Gargoyle [dungeon] | sim-verified (546.4 DPS) | yes | Prowler's Leather Waistguard (252473, +0.00 DPS) [crafted]; Skulker's Leather Waistguard (252474, +0.00 DPS) [crafted]; Might of the Timbermaw (19044, -9.14 DPS, sim-verified) [crafted] |
| legs | Lightforge Legplates (16728) | Stratholme: Baron Rivendare [dungeon] | sim-verified (546.4 DPS) | yes | Warbear Woolies (15065, +0.00 DPS) [crafted]; Azerothium Legplates (279265, +0.00 DPS) [crafted]; Magister's Leggings (16687, -5.07 DPS, sim-verified) [dungeon] |
| feet | Lightforge Boots (16725) | Stratholme: Balnazzar [dungeon] | sim-verified (546.4 DPS) | yes | Knight-Lieutenant's Lamellar Sabatons (227146, +0.00 DPS) [vendor]; Scalegut Treaders (275618, +0.00 DPS) [crafted]; Shadowcraft Boots (16711, -2.25 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 28.2 attack_power points (3.08 DPS) | yes | Falcon's Hook (7552, -0.31 DPS) [dungeon]; Ring of the Underwood (2951, -0.35 DPS) [world_drop]; Ring of Living Stone (18400, -0.43 DPS) [quest] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 25.4 attack_power points (2.77 DPS) | yes | Falcon's Hook (7552, -0.00 DPS) [dungeon]; Ring of the Underwood (2951, -0.05 DPS) [world_drop]; Ring of Living Stone (18400, -0.13 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (546.4 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (546.4 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Blade of Hanna (2801) | World drop [world_drop] | sim-verified (546.4 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Arcanite Champion (12790, -6.52 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Stalwart Helm; shoulder: Lightforge Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Lightforge Bracers; hands: Timbermaw Brawlers; waist: Lightforge Belt; legs: Lightforge Legplates; feet: Lightforge Boots; finger1: Masons Fraternity Ring; finger2: Thunderbrow Ring; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Blade of Hanna

No-known-source sample (15 of 1762, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

## Horde

### Band 20 (undead, 00000000000000000-00000000000000000-05231000000000000)

Set DPS (verified): 39.2. Weights run: 1.4s. Verify run: 2.3s. 222 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.491 ± 0.026, crit=0.844 ± 0.024 per rating point (14 rating = 1%, 11.816 per %), hit=1.363 ± 0.065 per rating point (10 rating = 1%, 13.628 per %), melee_haste=8.675 ± 0.564

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.87 DPS) | yes | Defender's Leather Hood (252447, -0.17 DPS) [crafted]; Guard's Silvered Chain Helm (250529, -0.22 DPS) [crafted]; Brawler's Leather Hood (252504, -0.35 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 3.0 attack_power points (0.13 DPS) | yes | Roadwatcher's Confidence (281265, +0.00 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 7.5 attack_power points (0.32 DPS) | yes | Rough Bronze Shoulders (3480, -0.06 DPS) [crafted]; Silvered Bronze Shoulders (3481, -0.06 DPS) [crafted] |
| back | Grave Shroud (279865) | Unending Torment [quest] | 9.0 attack_power points (0.39 DPS) | yes | Glowing Lizardscale Cloak (6449, -0.00 DPS) [dungeon]; Dark Leather Cloak (2316, -0.02 DPS) [crafted]; Lambent Scale Cloak (4706, -0.04 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (39.2 DPS) | yes | Mutant Scale Breastplate (6627, +0.00 DPS) [dungeon]; Veteran's Chain Shirt (250488, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -2.55 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.43 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bristlebark Bindings (14569, -0.07 DPS) [world_drop]; Forest Leather Bracers (3202, -0.11 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (39.2 DPS) | yes | Dagmire Gauntlets (6481, +0.00 DPS) [quest]; Thorbia's Gauntlets (12994, +0.00 DPS) [world_drop]; Gloves of the Fang (10413, -2.68 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (39.2 DPS) | yes | Brawler's Leather Belt (252428, -0.18 DPS) [crafted]; Deviate Scale Belt (6468, -0.20 DPS) [crafted]; Cobrahn's Grasp (6460, -2.77 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (39.2 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Veteran's Chain Leggings (250493, -2.28 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 17.9 attack_power points (0.78 DPS) | yes | Brawler's Leather Boots (252439, -0.02 DPS) [crafted]; Veteran's Boots (250503, -0.09 DPS) [crafted]; Defender's Leather Boots (252441, -0.34 DPS) [crafted] |
| finger1 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.35 DPS) | yes | Bounty Hunter's Ring (5351, -0.15 DPS) [quest]; The 1 Ring (8350, -0.20 DPS) [world]; Lavishly Jeweled Ring (1156, -0.22 DPS) [dungeon] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 6.0 attack_power points (0.26 DPS) | yes | The 1 Ring (8350, -0.11 DPS) [world]; Lavishly Jeweled Ring (1156, -0.13 DPS) [dungeon]; Bounty Hunter's Ring (5351, -0.73 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hammerbone (270018) | Leaders of the Fang [quest] | 306.2 attack_power points (13.31 DPS) | yes | Duskbringer (2205, -0.51 DPS) [dungeon]; Sword of the Fallen (286977, -0.79 DPS) [dungeon]; Smite's Mighty Hammer (7230, -1.21 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Grave Shroud; chest: Blackened Defias Armor; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Demon Band; finger2: Loop of Sacrifice; main_hand: Hammerbone

No-known-source sample (15 of 222, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 15402 Noosegrip Gauntlets; 18864 Insignia of the Alliance; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword

### Band 30 (undead, 00000000000000000-00000000000000000-05233301001300000)

Set DPS (verified): 88.0. Weights run: 1.3s. Verify run: 2.5s. 443 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.267 ± 0.023, crit=0.964 ± 0.025 per rating point (14 rating = 1%, 13.497 per %), hit=1.585 ± 0.069 per rating point (10 rating = 1%, 15.846 per %), melee_haste=9.455 ± 0.568

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 29.4 attack_power points (1.42 DPS) | yes | Tusken Helm (6686, -0.16 DPS) [dungeon]; Veteran's Chain Helm (250498, -0.26 DPS) [crafted]; Defender's Leather Helm (252455, -0.26 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 13.1 attack_power points (0.63 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.05 DPS) [world]; River Pride Choker (13087, -0.24 DPS) [world_drop]; Erudite's Amulet (277204, -0.51 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.9 attack_power points (1.15 DPS) | yes | Barbaric Iron Shoulders (7913, -0.21 DPS) [crafted]; Barbaric Shoulders (5964, -0.37 DPS) [crafted]; Golden Scale Shoulders (3841, -0.48 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 14.9 attack_power points (0.72 DPS) | yes | Tigerstrike Mantle (13108, -0.23 DPS) [world_drop]; Wildhunter Cloak (16658, -0.23 DPS) [quest]; Construct Cloak (279848, -0.23 DPS) [quest] |
| chest | Avenger's Armor (1488) | Razorfen Kraul [dungeon] | 30.0 attack_power points (1.45 DPS) | yes | Shining Silver Breastplate (2870, -0.10 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.12 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.29 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul [dungeon] | 16.0 attack_power points (0.77 DPS) | yes | Yorgen Bracers (13012, -0.01 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.02 DPS) [world_drop]; Barbaric Bracers (18948, -0.14 DPS) [crafted] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (88.0 DPS) | yes | Mail Combat Gauntlets (4075, +0.00 DPS) [world_drop]; Insignia Gloves (6408, +0.00 DPS) [world_drop]; Gauntlets of Ogre Strength (3341, -4.74 DPS, sim-verified) [world] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (88.0 DPS) | yes | Girdle of Golem Strength (9405, +0.00 DPS) [world_drop]; Prowler's Leather Belt (252459, +0.00 DPS) [crafted]; Officer's Belt (250556, -4.26 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (88.0 DPS) | yes | Ferine Leggings (6690, +0.00 DPS) [dungeon]; Brawler's Leather Legguards (252516, +0.00 DPS) [crafted]; Veteran's Silvered Chain Leggings (250523, -4.01 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (88.0 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Veteran's Boots (250503, +0.00 DPS) [crafted]; Brawler's Leather Boots (252439, -4.68 DPS, sim-verified) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.8 attack_power points (0.96 DPS) | yes | Silverlaine's Family Seal (6321, -0.47 DPS) [dungeon]; Tiger Band (6749, -0.47 DPS) [quest]; Bonebinder's Signet (273032, -0.47 DPS) [world] |
| finger2 | Band of the Fist (17694) | Allegiance to the Old Gods [quest] | 13.1 attack_power points (0.63 DPS) | yes | Silverlaine's Family Seal (6321, -0.15 DPS) [dungeon]; Tiger Band (6749, -0.15 DPS) [quest]; Bonebinder's Signet (273032, -0.15 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (88.0 DPS) | yes | Felbough Arm (273037, -1.35 DPS) [world]; Cobalt Crusher (7730, -1.91 DPS) [dungeon]; Viscous Hammer (13045, -22.80 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; feet: Blackened Defias Boots; finger1: Thunderbrow Ring; finger2: Band of the Fist; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 443, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 15402 Noosegrip Gauntlets

### Band 40 (undead, 01000000000000000-00000000000000000-05233301001301350)

Set DPS (verified): 131.3. Weights run: 1.3s. Verify run: 2.6s. 629 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.040 ± 0.004, agility=1.541 ± 0.030, crit=1.490 ± 0.036 per rating point (14 rating = 1%, 20.865 per %), hit=2.006 ± 0.088 per rating point (10 rating = 1%, 20.059 per %), melee_haste=11.290 ± 0.657

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chromite Barbute (8142) | Uldaman: Ancient Treasure [dungeon] | 43.0 attack_power points (2.24 DPS) | yes | White Bandit Mask (10008, -0.19 DPS) [crafted]; Barbaric Iron Helm (7915, -0.56 DPS) [crafted]; Hard Gold Coif (250537, -0.75 DPS) [crafted] |
| neck | Ethereal Talisman (4430) | The Crown of Will [quest] | 16.4 attack_power points (0.85 DPS) | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; Archmagister's Faceted Pendant (273049, -0.23 DPS) [world]; River Pride Choker (13087, -0.43 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 27.2 attack_power points (1.41 DPS) | yes | Flintrock Shoulders (7755, +0.00 DPS, sim-verified) [dungeon]; Hard Gold Pauldrons (250539, -0.25 DPS) [crafted]; Barbaric Iron Shoulders (7913, -0.30 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 16.9 attack_power points (0.88 DPS) | yes | Imperial Cloak (6432, -0.24 DPS) [dungeon]; Parachute Cloak (10518, -0.24 DPS) [crafted]; Tigerstrike Mantle (13108, -0.24 DPS) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Icemail Jerkin (1981, +0.00 DPS) [world_drop]; Carapace of Tuten'kash (10775, +0.00 DPS) [dungeon]; Kolkar Marauder Chain (6773, -4.88 DPS, sim-verified) [quest] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 20.4 attack_power points (1.06 DPS) | yes | Branded Leather Bracers (19508, +0.00 DPS, sim-verified) [dungeon]; Hawkeye's Bracers (14590, -0.16 DPS) [world_drop]; Yorgen Bracers (13012, -0.19 DPS) [world_drop] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scarlet Gauntlets (10331, +0.00 DPS) [dungeon]; Plated Fist of Hakoo (13071, +0.00 DPS) [world_drop]; Gloves of Holy Might (867, -1.27 DPS, sim-verified) [world_drop] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 32.2 attack_power points (1.68 DPS) | yes | Officer's Belt (250556, -0.13 DPS) [crafted]; Prowler's Leather Belt (252459, -0.24 DPS) [crafted]; Tharg's Shoelace (9705, -0.30 DPS) [quest] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | sim-verified (131.3 DPS) | yes | Symbolic Legplates (14829, -0.16 DPS) [world_drop]; Firemane Leggings (13129, -0.21 DPS) [world_drop]; Stormcloth Pants (10010, -1.52 DPS, sim-verified) [crafted] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 37.3 attack_power points (1.94 DPS) | yes | Prowler's Leather Shoes (252465, -0.21 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.32 DPS) [crafted]; Blackforge Greaves (6423, -0.98 DPS, sim-verified) [dungeon] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.9 attack_power points (1.09 DPS) | yes | Assault Band (13095, -0.05 DPS) [world_drop]; Falcon's Hook (7552, -0.05 DPS) [dungeon]; Ring of the Underwood (2951, -0.08 DPS) [world_drop] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.04 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Falcon's Hook (7552, -0.00 DPS) [dungeon]; Ring of the Underwood (2951, -0.03 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thornstone Sledgehammer (1722, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Nightblade (1982, -22.08 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Chromite Barbute; neck: Ethereal Talisman; chest: Stormcloth Vest; wrist: Ravager's Armguards; hands: Stormcloth Gloves; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Officer's Boots; finger2: Mark of Kern

No-known-source sample (15 of 629, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7748 Forcestone Buckler; 8708 Hammer of Expertise

### Band 50 (undead, 05024000000000000-00000000000000000-05233301001301350)

Set DPS (verified): 157.6. Weights run: 1.4s. Verify run: 2.9s. 839 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.509 ± 0.029, crit=1.792 ± 0.039 per rating point (14 rating = 1%, 25.091 per %), hit=2.283 ± 0.095 per rating point (10 rating = 1%, 22.826 per %), melee_haste=11.006 ± 0.993

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Stormcloth Headband (10032) | Tailoring [crafted] | sim-verified (157.6 DPS) | yes | Chromite Barbute (8142, +0.00 DPS) [dungeon]; Blood Guard's Plate Helm (220803, +0.00 DPS) [vendor]; Ornate Mithril Helm (7937, -5.14 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 30.6 attack_power points (1.62 DPS) | yes | Woven Ivy Necklace (19159, -0.20 DPS) [quest]; Ethereal Talisman (4430, -0.72 DPS) [quest]; Kaleidoscope Chain (13084, -0.83 DPS) [world_drop] |
| shoulder | Stormcloth Shoulders (10038) | Tailoring [crafted] | sim-verified (157.6 DPS) | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Officer's Pauldrons (250576, +0.00 DPS) [crafted]; Blessed Plate Pauldrons (250586, -1.90 DPS, sim-verified) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 22.6 attack_power points (1.20 DPS) | yes | Pridelord Cape (14673, -0.13 DPS) [world_drop]; Bloodlust Cape (14801, -0.15 DPS) [world_drop]; Blackflame Cape (13109, -0.24 DPS) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-verified (157.6 DPS) | yes | Warforged Chestplate (11195, +0.00 DPS) [quest]; Warbear Harness (15064, +0.00 DPS) [crafted]; Stone Guard's Plate Armor (220801, -3.42 DPS, sim-verified) [vendor] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 33.3 attack_power points (1.76 DPS) | yes | Deepfury Bracers (13120, -0.10 DPS) [world_drop]; Prowler's Leather Bracers (252539, -0.15 DPS) [crafted]; Giantslayer Bracers (13076, -0.20 DPS) [world_drop] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (157.6 DPS) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; Officer's Gloves (250551, +0.00 DPS) [crafted]; Blessed Plate Gauntlet (250588, -3.74 DPS, sim-verified) [crafted] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 48.9 attack_power points (2.59 DPS) | yes | Skulker's Leather Waistguard (252474, -0.07 DPS) [crafted]; Officer's Waistguard (250571, -0.16 DPS) [crafted]; Belt of the Gladiator (13134, -0.49 DPS) [world_drop] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-verified (157.6 DPS) | yes | Serpentskin Leggings (8262, +0.00 DPS) [world_drop]; Stormshroud Pants (15057, +0.00 DPS) [crafted]; Stone Guard's Plate Leggings (220798, +0.00 DPS) [vendor] |
| feet | Stormcloth Boots (10039) | Tailoring [crafted] | sim-verified (157.6 DPS) | yes | Prowler's Leather Boots (252468, +0.00 DPS) [crafted]; Skulker's Leather Boots (252469, +0.00 DPS) [crafted]; Blessed Plate Boots (250587, -2.69 DPS, sim-verified) [crafted] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.27 DPS) | yes | Masons Fraternity Ring (9533, -0.15 DPS) [quest]; Falcon's Hook (7552, -0.20 DPS) [dungeon]; Blackstone Ring (17713, -0.21 DPS) [dungeon] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 22.1 attack_power points (1.17 DPS) | yes | Masons Fraternity Ring (9533, -0.05 DPS) [quest]; Falcon's Hook (7552, -0.10 DPS) [dungeon]; Blackstone Ring (17713, -0.11 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (157.6 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (157.6 DPS) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-verified (157.6 DPS) | yes | Warmonger (13052, -0.22 DPS) [world_drop]; Blight (7959, -1.42 DPS, sim-verified) [crafted]; Stoneraven (13059, -1.75 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Stormcloth Headband; neck: Skibi's Pendant; shoulder: Stormcloth Shoulders; back: Dark Phantom Cape; wrist: Officer's Wristguards; waist: Prowler's Leather Waistguard; legs: Stormcloth Pants; feet: Stormcloth Boots; finger1: White Bone Band; finger2: Thunderbrow Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Thorium Greatmace

No-known-source sample (15 of 839, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60 (undead, 05025000200000000-50200000000000000-05233301001301350)

Set DPS (verified): 265.7. Weights run: 1.4s. Verify run: 7.1s. 1781 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=2.015 ± 0.048, crit=2.840 ± 0.069 per rating point (14 rating = 1%, 39.761 per %), hit=4.097 ± 0.173 per rating point (10 rating = 1%, 40.969 per %), melee_haste=19.815 ± 1.945

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Stalwart Helm (250599) | Blacksmithing [crafted] | sim-verified (265.7 DPS) | yes | Blood Guard's Plate Helm (220803, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -0.35 DPS) [crafted]; Lionheart Helm (12640, -5.07 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 37.2 attack_power points (2.02 DPS) | yes | Woven Ivy Necklace (19159, -0.32 DPS) [quest]; Sunprism Pendant (20649, -0.71 DPS) [quest]; Ethereal Talisman (4430, -0.98 DPS) [quest] |
| shoulder | Magister's Mantle (16689) | Scholomance: Ras Frostwhisper [dungeon] | sim-verified (265.7 DPS) | yes | Lightforge Spaulders (16729, +0.00 DPS) [dungeon]; Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Ironfeather Shoulders (15067, -3.62 DPS, sim-verified) [crafted] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 50.2 attack_power points (2.72 DPS) | yes | Fel Cape (279269, -0.57 DPS) [crafted]; Dark Phantom Cape (13122, -1.08 DPS) [world_drop]; Pridelord Cape (14673, -1.38 DPS) [world_drop] |
| chest | Magister's Robes (16688) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (265.7 DPS) | yes | Dawn Armor (252483, +0.00 DPS) [crafted]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Ironfeather Breastplate (15066, -7.29 DPS, sim-verified) [crafted] |
| wrist | Bindings of Elements (16671) | Stratholme: Crypt Beast [dungeon] | sim-verified (265.7 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS) [world_drop]; Arena Wristguards (18709, +0.00 DPS) [world]; Magister's Bindings (16683, -5.52 DPS, sim-verified) [dungeon] |
| hands | Gauntlets of Elements (16672) | Blackrock Spire: Pyroguard Emberseer [dungeon] | sim-verified (265.7 DPS) | yes | Devilsaur Gauntlets (15063, +0.00 DPS) [crafted]; Timbermaw Brawlers (19049, +0.00 DPS) [crafted]; Magister's Gloves (16684, -5.55 DPS, sim-verified) [dungeon] |
| waist | Cord of Elements (16673) | Blackrock Spire: Smolderthorn Witch Doctor [dungeon] | sim-verified (265.7 DPS) | yes | Might of the Timbermaw (19044, +0.00 DPS) [crafted]; Prowler's Leather Waistguard (252473, +0.00 DPS) [crafted]; Magister's Belt (16685, -3.39 DPS, sim-verified) [dungeon] |
| legs | Magister's Leggings (16687) | Stratholme: Baron Rivendare [dungeon] | sim-verified (265.7 DPS) | yes | Warbear Woolies (15065, +0.00 DPS) [crafted]; Azerothium Legplates (279265, +0.00 DPS) [crafted]; Savage Gladiator Leggings (11728, -10.59 DPS, sim-verified) [dungeon] |
| feet | Savage Gladiator Greaves (11731) | Blackrock Depths: Anub'shiah [dungeon] | sim-verified (265.7 DPS) | yes | Shadowcraft Boots (16711, +0.00 DPS) [dungeon]; Prowler's Leather Boots (252468, +0.00 DPS) [crafted]; Scalegut Treaders (275618, +0.00 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 28.2 attack_power points (1.53 DPS) | yes | Ring of the Underwood (2951, -0.20 DPS) [world_drop]; White Bone Band (11862, -0.23 DPS) [quest]; Thunderbrow Ring (13097, -0.25 DPS) [world_drop] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 24.7 attack_power points (1.34 DPS) | yes | Ring of the Underwood (2951, -0.01 DPS) [world_drop]; White Bone Band (11862, -0.04 DPS) [quest]; Thunderbrow Ring (13097, -0.06 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (265.7 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (265.7 DPS) | yes | Molten Heart of the Mountain (249470, -1.89 DPS, sim-verified) [crafted] |
| main_hand | Blade of Hanna (2801) | World drop [world_drop] | sim-verified (265.7 DPS) | yes | Hand of Edward the Odd (2243, +0.00 DPS, sim-verified) [world_drop]; High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Stalwart Helm; shoulder: Magister's Mantle; back: Cape of the Black Baron; chest: Magister's Robes; wrist: Bindings of Elements; hands: Gauntlets of Elements; waist: Cord of Elements; legs: Magister's Leggings; feet: Savage Gladiator Greaves; finger1: Masons Fraternity Ring; finger2: Falcon's Hook; main_hand: Blade of Hanna

No-known-source sample (15 of 1781, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60, raid preset (undead, 05015000200000000-50300000000000000-05233301001301350)

Set DPS (verified): 569.5. Weights run: 1.6s. Verify run: 7.3s. 1781 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.420 ± 0.005, agility=2.013 ± 0.042, crit=2.593 ± 0.058 per rating point (14 rating = 1%, 36.300 per %), hit=4.456 ± 0.257 per rating point (10 rating = 1%, 44.557 per %), melee_haste=33.789 ± 2.838

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 205.3 attack_power points (22.41 DPS) | yes | Stalwart Helm (250599, +0.00 DPS, sim-verified) [crafted]; Blood Guard's Plate Helm (220803, -10.15 DPS) [vendor]; Outlaw's Collar (279253, -11.85 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 38.3 attack_power points (4.18 DPS) | yes | Woven Ivy Necklace (19159, -0.61 DPS) [quest]; Sunprism Pendant (20649, -1.54 DPS) [quest]; Ethereal Talisman (4430, -1.98 DPS) [quest] |
| shoulder | Lightforge Spaulders (16729) | Blackrock Spire: The Beast [dungeon] | sim-verified (569.5 DPS) | yes | Blood Guard's Plate Pauldrons (220796, -0.99 DPS) [vendor]; Dark Warder's Pauldrons (22241, -1.21 DPS) [dungeon]; Magister's Mantle (16689, -5.39 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 50.2 attack_power points (5.48 DPS) | yes | Fel Cape (279269, -1.52 DPS) [crafted]; Dark Phantom Cape (13122, -2.18 DPS) [world_drop]; Pridelord Cape (14673, -2.71 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 155.3 attack_power points (16.95 DPS) | yes | Timbermaw Tunic (252484, -0.32 DPS) [crafted]; Shadowcraft Tunic (16721, -7.35 DPS) [dungeon]; Stone Guard's Plate Armor (220801, -9.03 DPS) [vendor] |
| wrist | Lightforge Bracers (16722) | Scholomance: Lord Alexei Barov [dungeon] | sim-verified (569.5 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS) [world_drop]; Tranquil Wristguards (279256, +0.00 DPS) [crafted]; Magister's Bindings (16683, -5.74 DPS, sim-verified) [dungeon] |
| hands | Timbermaw Brawlers (19049) | Leatherworking [crafted] | 79.0 attack_power points (8.62 DPS) | yes | Devilsaur Gauntlets (15063, -1.60 DPS) [crafted]; Officer's Gloves (250551, -2.42 DPS) [crafted]; Gloves of Holy Might (867, -2.48 DPS) [world_drop] |
| waist | Lightforge Belt (16723) | Stratholme: Rockwing Gargoyle [dungeon] | sim-verified (569.5 DPS) | yes | Prowler's Leather Waistguard (252473, +0.00 DPS) [crafted]; Skulker's Leather Waistguard (252474, +0.00 DPS) [crafted]; Might of the Timbermaw (19044, -9.07 DPS, sim-verified) [crafted] |
| legs | Lightforge Legplates (16728) | Stratholme: Baron Rivendare [dungeon] | sim-verified (569.5 DPS) | yes | Warbear Woolies (15065, +0.00 DPS) [crafted]; Azerothium Legplates (279265, +0.00 DPS) [crafted]; Magister's Leggings (16687, -8.50 DPS, sim-verified) [dungeon] |
| feet | Lightforge Boots (16725) | Stratholme: Balnazzar [dungeon] | sim-verified (569.5 DPS) | yes | Prowler's Leather Boots (252468, +0.00 DPS) [crafted]; Scalegut Treaders (275618, +0.00 DPS) [crafted]; Shadowcraft Boots (16711, -3.02 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 28.2 attack_power points (3.08 DPS) | yes | Falcon's Hook (7552, -0.31 DPS) [dungeon]; Ring of the Underwood (2951, -0.35 DPS) [world_drop]; Ring of Living Stone (18400, -0.43 DPS) [quest] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 25.4 attack_power points (2.77 DPS) | yes | Falcon's Hook (7552, -0.00 DPS) [dungeon]; Ring of the Underwood (2951, -0.05 DPS) [world_drop]; Ring of Living Stone (18400, -0.13 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (569.5 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (569.5 DPS) | yes | Molten Heart of the Mountain (249470, -4.42 DPS, sim-verified) [crafted] |
| main_hand | Blade of Hanna (2801) | World drop [world_drop] | sim-verified (569.5 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Arcanite Champion (12790, -14.73 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Lionheart Helm; shoulder: Lightforge Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Lightforge Bracers; hands: Timbermaw Brawlers; waist: Lightforge Belt; legs: Lightforge Legplates; feet: Lightforge Boots; finger1: Masons Fraternity Ring; main_hand: Blade of Hanna

No-known-source sample (15 of 1781, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

