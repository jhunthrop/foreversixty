# Leveling BiS: Restoration

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-00000000000000000000-5050010000000000)

Set DPS (verified): 48.5. Weights run: 5.7s. Verify run: 19.6s. 196 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.280 ± 0.010, spirit=1.607 ± 0.011, mp5=3.265 ± 0.016, crit=0.046 ± 0.003 per rating point (14 rating = 1%, 0.646 per %), spell_haste=-3.341 ± 0.170

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.90 DPS) | yes | Wisdom's Leather Hood (252507, -0.14 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.21 DPS) [crafted]; Stormrider's Leather Hood (252506, -0.37 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 6.4 healing_power points (0.68 DPS) | yes | Scholarly Pendant (277203, -0.21 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.5 healing_power points (1.22 DPS) | yes | Slime-encrusted Pads (6461, -0.18 DPS) [dungeon]; Forest Leather Mantle (4709, -0.54 DPS) [world_drop]; Prospector's Pads (14566, -0.54 DPS) [world_drop] |
| back | Spirit Cloak (4792) (or Sylvan Cloak (4793), Heavy Woolen Cloak (4311)) | Jennabink Powerseam [vendor] | 6.4 healing_power points (0.68 DPS) | yes | Heavy Woolen Cloak (4311, +0.00 DPS) [crafted]; Sylvan Cloak (4793, +0.00 DPS) [vendor]; Seer's Cape (6378, -0.07 DPS) [dungeon] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 26.2 healing_power points (2.77 DPS) | yes | Robe of the Moccasin (6465, +0.00 DPS, sim-verified) [dungeon]; Filigreed Pristine Gown (253901, -0.21 DPS) [crafted]; Corsair's Overshirt (5202, -0.25 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 16.2 healing_power points (1.71 DPS) | yes | Drakewing Bands (12999, -0.60 DPS, sim-verified) [world_drop]; Bravo's Armbands (270015, -1.03 DPS) [quest]; Owl Bracers (4796, -1.03 DPS) [vendor] |
| hands | Wisdom's Leather Gloves (252499) | Leatherworking [crafted] | 15.1 healing_power points (1.60 DPS) | yes | Magefist Gloves (12977, -0.24 DPS) [world_drop]; Pristine Gloves (253913, -0.29 DPS, sim-verified) [crafted]; Bright Gloves (3066, -0.38 DPS) [world_drop] |
| waist | Novice Ardent's Sash (253887) | Tailoring [crafted] | sim-verified (48.6 DPS) | yes | Wisdom's Leather Belt (252433, -0.16 DPS) [crafted]; Keller's Girdle (2911, -0.47 DPS) [world_drop]; Pristine Sash (253925, -0.54 DPS, sim-verified) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 34.1 healing_power points (3.60 DPS) | yes | Filigreed Pristine Leggings (253937, -0.27 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -1.64 DPS) [world_drop]; Filigreed Silky Leggings (253939, -2.11 DPS) [crafted] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 17.4 healing_power points (1.84 DPS) | yes | Pristine Boots (253889, -0.48 DPS) [crafted]; Black Whelp Slippers (252424, -0.70 DPS, sim-verified) [crafted]; Kimbra Boots (6191, -0.79 DPS) [quest] |
| finger1 | Band of Purification (12996) | World drop [world_drop] | 9.6 healing_power points (1.02 DPS) | yes | Lavishly Jeweled Ring (1156, -0.21 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.61 DPS) [world_drop]; Demon Band (12054, -0.68 DPS) [world_drop] |
| finger2 | Deep Fathom Ring (6463) | Wailing Caverns: Mutanus the Devourer [dungeon] | 8.0 healing_power points (0.85 DPS) | yes | Lavishly Jeweled Ring (1156, -0.39 DPS, sim-verified) [dungeon]; Volcanic Rock Ring (12053, -0.44 DPS) [world_drop]; Demon Band (12054, -0.51 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps [dungeon] | 16.1 healing_power points (1.70 DPS) | yes | Staff of Westfall (2042, -0.30 DPS, sim-verified) [quest]; Twisted Chanter's Staff (890, -0.35 DPS) [world_drop]; Evocator's Blade (2567, -0.58 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Tarnished Locket; shoulder: Magician's Mantle; back: Spirit Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Wisdom's Leather Gloves; waist: Novice Ardent's Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Band of Purification; finger2: Deep Fathom Ring; main_hand: Staff of the Blessed Seer

No-known-source sample (15 of 196, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18853 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (night-elf, 0000000000000000-00000000000000000000-5050035110010000)

Set DPS (verified): 92.6. Weights run: 8.1s. Verify run: 27.2s. 380 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.459 ± 0.017, spirit=2.982 ± 0.023, mp5=4.518 ± 0.031, crit=0.055 ± 0.005 per rating point (14 rating = 1%, 0.766 per %), spell_haste=not significant (-0.089 ± 0.301)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 50.9 healing_power points (5.29 DPS) | yes | Nightsky Cowl (4039, -1.92 DPS) [world_drop]; Dokebi Hat (14584, -2.07 DPS) [world_drop]; Enduring Cap (3020, -2.20 DPS, sim-verified) [world_drop] |
| neck | Pendant of Myzrael (4614) (or Glowing Green Talisman (5002)) | Razorfen Downs: Splinterbone Captain [dungeon] | 17.9 healing_power points (1.86 DPS) | yes | Glowing Green Talisman (5002, +0.00 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.01 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.32 DPS) [world_drop] |
| shoulder | Mantle of Honor (3560) | Bride of the Embalmer [quest] | 31.1 healing_power points (3.23 DPS) | yes | Faerie Mantle (5820, -0.39 DPS, sim-verified) [quest]; Nightsky Mantle (4718, -0.47 DPS) [world_drop]; Insignia Mantle (4721, -1.06 DPS) [world_drop] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 30.9 healing_power points (3.21 DPS) | yes | Prelacy Cape (7004, -0.20 DPS) [quest]; Repairman's Cape (9605, -0.79 DPS) [quest]; Amy's Blanket (13005, -1.04 DPS) [world_drop] |
| chest | Wisdom's Leather Tunic (252511) | Leatherworking [crafted] | 44.6 healing_power points (4.64 DPS) | yes | Pristine Gown (253961, -0.26 DPS) [crafted]; Beguiler Robes (7728, -0.34 DPS) [dungeon]; Pressed Felt Robe (1997, -0.47 DPS) [world] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 20.8 healing_power points (2.17 DPS) | yes | Nightsky Wristbands (6407, -0.33 DPS) [world_drop]; Dokebi Bracers (14580, -0.47 DPS) [world_drop]; Drakewing Bands (12999, -1.11 DPS, sim-verified) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 43.1 healing_power points (4.47 DPS) | yes | Zodiac Gloves (7106, -1.70 DPS) [quest]; Hotshot Pilot's Gloves (9491, -1.71 DPS) [dungeon]; Shilly Mitts (9609, -1.82 DPS, sim-verified) [quest] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 45.6 healing_power points (4.74 DPS) | yes | Dokebi Cord (14578, -2.28 DPS) [world_drop]; Resilient Cord (14406, -2.59 DPS) [world_drop]; Silver-lined Belt (13011, -2.81 DPS, sim-verified) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 52.1 healing_power points (5.41 DPS) | yes | Pristine Leggings (253987, -0.52 DPS) [crafted]; Stormrider's Leather Kilt (252518, -0.94 DPS) [crafted]; Earthen Leggings (253999, -1.32 DPS, sim-verified) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 42.1 healing_power points (4.38 DPS) | yes | Resilient Boots (14399, -1.90 DPS) [world_drop]; Stonecloth Boots (14408, -1.90 DPS) [world_drop]; Boots of the Enchanter (4325, -3.42 DPS, sim-verified) [crafted] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 26.9 healing_power points (2.79 DPS) | yes | The Queen's Jewel (13094, -0.01 DPS) [world_drop]; Ring of Calm (6790, -0.62 DPS) [quest]; Band of Thorns (5007, -0.93 DPS) [world_drop] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 26.9 healing_power points (2.79 DPS) | yes | The Queen's Jewel (13094, -0.17 DPS, sim-verified) [world_drop]; Ring of Calm (6790, -0.62 DPS) [quest]; Band of Thorns (5007, -0.93 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (92.7 DPS) | yes | Hardened Root Staff (1317, -2.20 DPS) [quest]; Gnarled Ash Staff (791, -2.42 DPS) [world_drop]; Manual Crowd Pummeler (9449, -3.45 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 30.9 healing_power points (3.21 DPS) | yes | Orb of Mistmantle (13031, -1.06 DPS, sim-verified) [world_drop]; Strength of Will (4837, -1.35 DPS) [vendor]; Torch of Holy Flame (7344, -1.66 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Pendant of Myzrael; shoulder: Mantle of Honor; back: Glowing Thresher Cape; chest: Wisdom's Leather Tunic; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 380, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 40 (night-elf, 0000000000000000-00000000000000000000-5050035153112000)

Set DPS (verified): 132.9. Weights run: 8.9s. Verify run: 32.1s. 516 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.376 ± 0.021, spirit=2.703 ± 0.020, mp5=5.010 ± 0.038, crit=0.059 ± 0.005 per rating point (14 rating = 1%, 0.822 per %), spell_haste=not significant (1.040 ± 0.404)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 61.7 healing_power points (6.45 DPS) | yes | Whitemane's Chapeau (7720, -0.26 DPS) [dungeon]; Holy Shroud (2721, -1.31 DPS) [world_drop]; Earthen Silk Hood (254015, -1.33 DPS) [crafted] |
| neck | Necklace of Calisea (1714) | Zul'Farrak: Witch Doctor's Chest [dungeon] | 28.6 healing_power points (2.98 DPS) | yes | Amberglow Talisman (10824, -0.16 DPS) [quest]; Explorers' League Commendation (7746, -0.44 DPS) [quest]; Glowing Eye of Mordresh (10769, -0.47 DPS) [dungeon] |
| shoulder | Earthen Silk Shoulders (254033) | Tailoring [crafted] | 39.6 healing_power points (4.14 DPS) | yes | Sheepshear Mantle (13115, -0.14 DPS) [world_drop]; Inquisitor's Shawl (19507, -0.70 DPS) [dungeon]; Crimson Silk Shoulders (7059, -1.01 DPS) [crafted] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 28.6 healing_power points (2.99 DPS) | yes | Prelacy Cape (7004, -0.15 DPS) [quest]; Ceremonial Centaur Blanket (6789, -0.16 DPS) [quest]; Energy Cloak (9397, -0.17 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 79.6 healing_power points (8.32 DPS) | yes | Doomsayer's Robe (4746, -2.19 DPS) [quest]; Dreamweave Vest (10021, -3.07 DPS) [crafted]; Civinad Robes (9623, -3.15 DPS) [quest] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 32.5 healing_power points (3.40 DPS) | yes | Earthen Silk Cuffs (254019, -0.94 DPS) [crafted]; Mindthrust Bracers (1974, -1.11 DPS) [dungeon]; Silkstream Cuffs (16791, -1.55 DPS) [quest] |
| hands | Earthen Silk Gloves (254017) | Tailoring [crafted] | 39.6 healing_power points (4.14 DPS) | yes | Gloves of Old (9395, -0.03 DPS) [world_drop]; Warden's Gloves (14606, -0.46 DPS) [world_drop]; Windchaser Handguards (14431, -0.47 DPS) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 49.2 healing_power points (5.14 DPS) | yes | Mender's Leather Belt (252523, -0.60 DPS) [crafted]; Windchaser Cinch (14435, -1.45 DPS) [world_drop]; Sutarn's Ring (13105, -2.01 DPS) [world_drop] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 62.1 healing_power points (6.49 DPS) | yes | Stoneweaver Leggings (9407, -0.59 DPS) [dungeon]; Stormcloth Pants (10010, -0.68 DPS) [crafted]; Wisdom's Leather Leggings (252519, -1.31 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) | Leatherworking [crafted] | sim-verified (132.9 DPS) | yes | Furen's Boots (13100, +0.00 DPS) [world_drop]; Thoughtcast Boots (10578, -0.98 DPS) [dungeon]; Gilded Slippers (254001, -1.26 DPS) [crafted] |
| finger1 | Welken Ring (5011) | World drop [world_drop] | 25.8 healing_power points (2.70 DPS) | yes | Runemender's Seal (273031, -0.09 DPS) [world]; The Queen's Jewel (13094, -0.15 DPS) [world_drop]; Blush Ember Ring (13093, -0.44 DPS) [world_drop] |
| finger2 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 24.9 healing_power points (2.60 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; The Queen's Jewel (13094, -0.06 DPS) [world_drop]; Blush Ember Ring (13093, -0.34 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -4.64 DPS) [dungeon]; Death Speaker Scepter (2816, -4.97 DPS) [dungeon]; Hypnotic Blade (7714, -5.48 DPS) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 29.2 healing_power points (3.05 DPS) | yes | Orb of Lorica (11262, -0.36 DPS) [quest]; Beacon of Hope (9393, -0.46 DPS) [dungeon]; Silksand Star (15964, -0.51 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Necklace of Calisea; shoulder: Earthen Silk Shoulders; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Earthen Silk Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; finger1: Welken Ring; finger2: Electrocutioner Lagnut; trinket1: Ankh of Life; main_hand: Hand of Righteousness

No-known-source sample (15 of 516, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 50 (night-elf, 4300000000000000-00000000000000000000-5050035153113200)

Set DPS (verified): 227.1. Weights run: 9.1s. Verify run: 31.5s. 686 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.560 ± 0.022, spirit=2.872 ± 0.023, mp5=4.645 ± 0.032, crit=0.077 ± 0.008 per rating point (14 rating = 1%, 1.079 per %), spell_haste=not significant (0.637 ± 0.270)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 87.2 healing_power points (8.76 DPS) | yes | Engineer's Guild Headpiece (9534, -0.51 DPS) [quest]; Knight-Lieutenant's Restored Leather Helm (220874, -0.56 DPS) [vendor]; Cassandra's Grace (13102, -1.75 DPS) [world_drop] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 33.3 healing_power points (3.35 DPS) | yes | Necklace of Calisea (1714, -0.29 DPS, sim-verified) [dungeon]; Amberglow Talisman (10824, -0.46 DPS) [quest]; Explorers' League Commendation (7746, -0.75 DPS) [quest] |
| shoulder | Living Shoulders (15061) | Leatherworking [crafted] | 68.3 healing_power points (6.87 DPS) | yes | Knight-Lieutenant's Restored Leather Spaulders (220876, -0.20 DPS) [vendor]; Nethergeld Shoulders (254049, -1.27 DPS) [crafted]; Mender's Leather Shoulders (252538, -1.73 DPS, sim-verified) [crafted] |
| back | Wingveil Cloak (10802) | Sunken Temple: Spawn of Hakkar [dungeon] | 36.6 healing_power points (3.68 DPS) | yes | Chainlink Towel (9648, -0.32 DPS) [quest]; Traveler's Cloak (8297, -0.50 DPS) [world_drop]; Arcane Cloak (8286, -0.76 DPS, sim-verified) [world_drop] |
| chest | Ghostweave Vest (14141) | Tailoring [crafted] | 94.6 healing_power points (9.51 DPS) | yes | Feathered Breastplate (8349, -0.86 DPS, sim-verified) [crafted]; Stormcloth Vest (10020, -1.37 DPS) [crafted]; Knight's Restored Leather Jerkin (220878, -1.71 DPS) [vendor] |
| wrist | Mender's Leather Bracers (252543) | Leatherworking [crafted] | 47.2 healing_power points (4.74 DPS) | yes | Nethergeld Cuffs (254061, -0.39 DPS, sim-verified) [crafted]; Enchanted Kodo Bracers (13119, -1.22 DPS) [world_drop]; Venomshroud Armguards (14439, -1.80 DPS) [world_drop] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | 70.9 healing_power points (7.12 DPS) | yes | Gilded Gloves (254095, -0.44 DPS, sim-verified) [crafted]; Earthenweave Gloves (254075, -1.34 DPS) [crafted]; Keeper's Gloves (14666, -2.51 DPS) [world_drop] |
| waist | Earthenweave Cord (254077) | Tailoring [crafted] | 57.6 healing_power points (5.79 DPS) | yes | Mender's Leather Waistguard (252477, -0.89 DPS) [crafted]; Gilded Cord (254037, -1.17 DPS, sim-verified) [crafted]; Mender's Leather Belt (252523, -1.21 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 87.6 healing_power points (8.80 DPS) | yes | Dalewind Trousers (13008, -0.18 DPS) [world_drop]; Kilt of the Atal'ai Prophet (10807, -0.18 DPS) [dungeon]; Jinxed Hoodoo Kilt (9474, -0.83 DPS, sim-verified) [dungeon] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 69.9 healing_power points (7.03 DPS) | yes | Sergeant Major's Restored Leather Boots (220884, -0.56 DPS) [vendor]; Earthenweave Boots (254093, -1.24 DPS) [crafted]; Furen's Boots (13100, -1.31 DPS) [world_drop] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 43.8 healing_power points (4.40 DPS) | yes | Coldwater Ring (4550, -1.34 DPS) [quest]; Seedtime Hoop (9655, -1.52 DPS) [quest]; Welken Ring (5011, -1.60 DPS) [world_drop] |
| finger2 | Choking Band (11868) | Ogre Head On A Stick = Party [quest] | 37.3 healing_power points (3.75 DPS) | yes | Seedtime Hoop (9655, -0.87 DPS) [quest]; Coldwater Ring (4550, -0.88 DPS, sim-verified) [quest]; Welken Ring (5011, -0.95 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (227.1 DPS) | yes | Thunderbrew's Boot Flask (744, -1.15 DPS) [quest]; Uther's Strength (11302, -1.65 DPS) [world_drop]; Alchemist's Stone (13503, -2.31 DPS) [crafted] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-verified (227.1 DPS) | yes | Thunderbrew's Boot Flask (744, -0.58 DPS) [quest]; Uther's Strength (11302, -0.65 DPS, sim-verified) [world_drop]; Alchemist's Stone (13503, -1.73 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (227.1 DPS) | yes | Charstone Dirk (17710, -7.90 DPS) [dungeon]; Hand of Righteousness (7721, -10.46 DPS) [dungeon]; Hammer of the Northern Wind (810, -34.70 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Living Shoulders; back: Wingveil Cloak; chest: Ghostweave Vest; wrist: Mender's Leather Bracers; hands: Mender's Leather Gauntlets; waist: Earthenweave Cord; legs: Senior Designer's Pantaloons; feet: Sandals of the Insurgent; finger1: Eye of Adaegus; finger2: Choking Band; trinket2: Evonice's Landin' Pilla; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 686, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60 (night-elf, 4532200100000000-00000000000000000000-5050035153113200)

Set DPS (verified): 337.4. Weights run: 9.3s. Verify run: 84.2s. 1564 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.021, intellect=2.344 ± 0.054, spirit=4.387 ± 0.047, mp5=7.227 ± 0.055, crit=0.114 ± 0.012 per rating point (14 rating = 1%, 1.591 per %), spell_haste=not significant (0.465 ± 0.667)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 176.6 healing_power points (18.14 DPS) | yes | Lieutenant Commander's Dragonhide Headdress (227199, -2.70 DPS) [vendor]; Wildheart Cowl (16720, -2.77 DPS) [dungeon]; Devout Crown (16693, -3.55 DPS) [dungeon] |
| neck | Lady Alizabeth's Pendant (13002) | World drop [world_drop] | 65.8 healing_power points (6.76 DPS) | yes | Horizon Choker (13085, -1.59 DPS) [world_drop]; Necklace of Calisea (1714, -1.92 DPS) [dungeon]; Amberglow Talisman (10824, -2.25 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | 99.7 healing_power points (10.24 DPS) | yes | Lieutenant Commander's Dragonhide Pauldrons (227201, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Pauldrons (231705, +0.00 DPS) [vendor]; Wildheart Spaulders (16718, -0.25 DPS) [dungeon] |
| back | Cloak of Earth and Sky (279270) | Tailoring [crafted] | 61.4 healing_power points (6.31 DPS) | yes | Ethereal Mist Cape (11873, -0.18 DPS) [quest]; Indomitable Cloak (14683, -0.45 DPS) [world_drop]; Sergeant Major's Cape (16337, -0.78 DPS) [pvp] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | 156.2 healing_power points (16.05 DPS) | yes | Ironfeather Breastplate (15066, -0.54 DPS) [crafted]; Wildheart Vest (16706, -1.09 DPS) [dungeon]; Ghostweave Vest (14141, -1.10 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | 73.6 healing_power points (7.56 DPS) | yes | Marshal's Dragonhide Bracers (16445, +0.00 DPS) [pvp]; Earthenweave Cuffs (254131, -0.80 DPS) [crafted]; Mender's Leather Bracers (252543, -1.22 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wildheart Gloves (16717, +0.00 DPS) [dungeon]; Swarmtender's Gloves (275607, -0.95 DPS) [crafted]; Marshal's Dragonhide Gauntlets (231706, -1.54 DPS) [pvp] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (337.4 DPS) | yes | Marshal's Dragonhide Waistguard (16447, +0.00 DPS) [pvp]; Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Elderwild Waistcord (279252, +0.00 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 160.1 healing_power points (16.44 DPS) | yes | Knight-Captain's Dragonhide Legguards (227200, -1.10 DPS) [vendor]; Elderwild Pants (279254, -1.52 DPS) [crafted]; Living Leggings (15060, -2.50 DPS) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; Wildheart Boots (16715, -0.24 DPS) [dungeon]; Sandals of the Insurgent (13111, -0.26 DPS) [world_drop] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 71.7 healing_power points (7.36 DPS) | yes | Choking Band (11868, -1.51 DPS) [quest]; Coldwater Ring (4550, -2.59 DPS) [quest]; Seedtime Hoop (9655, -2.86 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 66.7 healing_power points (6.85 DPS) | yes | Choking Band (11868, -0.99 DPS) [quest]; Coldwater Ring (4550, -2.07 DPS) [quest]; Seedtime Hoop (9655, -2.35 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -0.95 DPS) [quest]; Thunderbrew's Boot Flask (744, -1.80 DPS) [quest]; Uther's Strength (11302, -3.56 DPS) [world_drop] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -0.05 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.90 DPS) [quest]; Uther's Strength (11302, -2.66 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Edward the Odd (2243, -9.56 DPS) [world_drop]; Charstone Dirk (17710, -10.40 DPS) [dungeon]; Staff of Hale Magefire (13000, -10.51 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Lady Alizabeth's Pendant; shoulder: Devout Mantle; back: Cloak of Earth and Sky; chest: Mooncloth Vest; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60, raid preset (night-elf, 4532200100000000-00000000000000000000-5050035153113200)

Set DPS (verified): 638.8. Weights run: 6.7s. Verify run: 54.9s. 1564 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.169, intellect=2.204 ± 0.090, spirit=2.354 ± 0.086, mp5=3.646 ± 0.107, crit=0.324 ± 0.034 per rating point (14 rating = 1%, 4.536 per %), spell_haste=not significant (0.379 ± 1.406)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 145.6 healing_power points (24.28 DPS) | yes | Lieutenant Commander's Dragonhide Headdress (227199, -4.74 DPS) [vendor]; Wildheart Cowl (16720, -6.58 DPS) [dungeon]; Devout Crown (16693, -12.88 DPS, sim-verified) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 40.3 healing_power points (6.72 DPS) | yes | Lady Alizabeth's Pendant (13002, +0.00 DPS, sim-verified) [world_drop]; Necklace of Calisea (1714, -1.40 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -1.61 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (638.8 DPS) | yes | Lieutenant Commander's Dragonhide Pauldrons (227201, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Pauldrons (231705, +0.00 DPS) [vendor]; Wildheart Spaulders (16718, -4.53 DPS, sim-verified) [dungeon] |
| back | Sergeant Major's Cape (16337) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 36.5 healing_power points (6.08 DPS) | yes | Ethereal Mist Cape (11873, -0.27 DPS) [quest]; Imperial Red Cloak (8248, -0.47 DPS) [world_drop]; Cloak of Earth and Sky (279270, -0.59 DPS) [crafted] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (638.8 DPS) | yes | Field Marshal's Dragonhide Tunic (231702, -0.36 DPS) [vendor]; Knight-Captain's Dragonhide Tunic (227196, -0.36 DPS) [pvp]; Mooncloth Vest (14138, -3.80 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | 52.2 healing_power points (8.70 DPS) | yes | Marshal's Dragonhide Bracers (16445, +0.00 DPS) [pvp]; Nethergeld Cuffs (254061, -0.83 DPS) [crafted]; Mender's Leather Bracers (252543, -5.83 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (638.8 DPS) | yes | Marshal's Dragonhide Gauntlets (231706, -0.24 DPS) [pvp]; Mender's Leather Gauntlets (252551, -0.84 DPS) [crafted]; Wildheart Gloves (16717, -5.70 DPS, sim-verified) [dungeon] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (638.8 DPS) | yes | Marshal's Dragonhide Waistguard (16447, +0.00 DPS) [pvp]; Elderwild Waistcord (279252, -0.65 DPS) [crafted]; Wisdom of the Timbermaw (19047, -2.89 DPS, sim-verified) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 111.2 healing_power points (18.54 DPS) | yes | Knight-Captain's Dragonhide Legguards (227200, +0.00 DPS) [vendor]; Marshal's Dragonhide Legguards (231703, -2.33 DPS) [pvp]; Elderwild Pants (279254, -5.18 DPS, sim-verified) [crafted] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 85.4 healing_power points (14.25 DPS) | yes | Wildheart Boots (16715, -2.44 DPS) [dungeon]; Marshal's Dragonhide Greaves (231704, -2.97 DPS) [vendor]; Devout Sandals (16691, -3.07 DPS, sim-verified) [dungeon] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 47.9 healing_power points (7.99 DPS) | yes | Choking Band (11868, -2.89 DPS) [quest]; Maiden's Circle (13001, -3.06 DPS) [world_drop]; Coldwater Ring (4550, -3.36 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 41.5 healing_power points (6.91 DPS) | yes | Choking Band (11868, -1.70 DPS, sim-verified) [quest]; Maiden's Circle (13001, -1.99 DPS) [world_drop]; Coldwater Ring (4550, -2.28 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (638.8 DPS) | yes | Evonice's Landin' Pilla (18951, -0.78 DPS) [quest]; Shard of the Splithooves (10659, -1.06 DPS) [quest]; Thunderbrew's Boot Flask (744, -1.57 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (638.8 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Shard of the Splithooves (10659, +0.00 DPS) [quest]; Evonice's Landin' Pilla (18951, -0.71 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (638.8 DPS) | yes | Charstone Dirk (17710, -14.60 DPS) [dungeon]; Hand of Righteousness (7721, -20.39 DPS) [dungeon]; Hand of Edward the Odd (2243, -20.71 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; shoulder: Devout Mantle; back: Sergeant Major's Cape; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket2: Uther's Strength

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

## Horde

### Band 20 (tauren, 0000000000000000-00000000000000000000-5050010000000000)

Set DPS (verified): 47.8. Weights run: 5.7s. Verify run: 18.5s. 192 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.280 ± 0.010, spirit=1.607 ± 0.011, mp5=3.265 ± 0.016, crit=0.046 ± 0.003 per rating point (14 rating = 1%, 0.646 per %), spell_haste=-3.341 ± 0.170

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.90 DPS) | yes | Wisdom's Leather Hood (252507, -0.14 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.21 DPS) [crafted]; Stormrider's Leather Hood (252506, -0.37 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 4.8 healing_power points (0.51 DPS) | yes | Scholarly Pendant (277203, -1.28 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.5 healing_power points (1.22 DPS) | yes | Slime-encrusted Pads (6461, -0.28 DPS, sim-verified) [dungeon]; Forest Leather Mantle (4709, -0.54 DPS) [world_drop]; Prospector's Pads (14566, -0.54 DPS) [world_drop] |
| back | Spirit Cloak (4792) (or Antiquated Cloak (4799), Sylvan Cloak (4793), Heavy Woolen Cloak (4311)) | Jennabink Powerseam [vendor] | 6.4 healing_power points (0.68 DPS) | yes | Heavy Woolen Cloak (4311, +0.00 DPS) [crafted]; Sylvan Cloak (4793, +0.00 DPS) [vendor]; Antiquated Cloak (4799, +0.00 DPS) [vendor] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 26.2 healing_power points (2.77 DPS) | yes | Filigreed Pristine Gown (253901, -0.21 DPS) [crafted]; Corsair's Overshirt (5202, -0.25 DPS) [dungeon]; Robe of the Moccasin (6465, -0.78 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 16.2 healing_power points (1.71 DPS) | yes | Drakewing Bands (12999, -0.80 DPS, sim-verified) [world_drop]; Crystalline Cuffs (14148, -0.93 DPS) [dungeon]; Owl Bracers (4796, -1.03 DPS) [vendor] |
| hands | Blight Gloves (279877) | The New Plague [quest] | 17.0 healing_power points (1.79 DPS) | yes | Wisdom's Leather Gloves (252499, -0.22 DPS, sim-verified) [crafted]; Pristine Gloves (253913, -0.23 DPS) [crafted]; Magefist Gloves (12977, -0.44 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.1 healing_power points (1.70 DPS) | yes | Novice Ardent's Sash (253887, -0.25 DPS, sim-verified) [crafted]; Wisdom's Leather Belt (252433, -0.32 DPS) [crafted]; Keller's Girdle (2911, -0.62 DPS) [world_drop] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 34.1 healing_power points (3.60 DPS) | yes | Filigreed Pristine Leggings (253937, -0.28 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -1.64 DPS) [world_drop]; Filigreed Silky Leggings (253939, -2.11 DPS) [crafted] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 17.4 healing_power points (1.84 DPS) | yes | Pristine Boots (253889, -0.48 DPS) [crafted]; Smoldering Boots (3076, -0.82 DPS) [world]; Black Whelp Slippers (252424, -0.98 DPS, sim-verified) [crafted] |
| finger1 | Band of Purification (12996) | World drop [world_drop] | 9.6 healing_power points (1.02 DPS) | yes | Lavishly Jeweled Ring (1156, -0.21 DPS) [dungeon]; Loop of Sacrifice (281673, -0.34 DPS) [quest]; Volcanic Rock Ring (12053, -0.61 DPS) [world_drop] |
| finger2 | Deep Fathom Ring (6463) | Wailing Caverns: Mutanus the Devourer [dungeon] | 8.0 healing_power points (0.85 DPS) | yes | Loop of Sacrifice (281673, -0.17 DPS) [quest]; Volcanic Rock Ring (12053, -0.44 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -1.26 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Orgrimmar (15444) | Hidden Enemies [quest] | 19.7 healing_power points (2.08 DPS) | yes | Staff of the Blessed Seer (2271, -0.52 DPS, sim-verified) [dungeon]; Twisted Chanter's Staff (890, -0.73 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.73 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Spirit Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Blight Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Band of Purification; finger2: Deep Fathom Ring; main_hand: Staff of Orgrimmar

No-known-source sample (15 of 192, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion; 21566 Rune of Perfection

### Band 30 (tauren, 0000000000000000-00000000000000000000-5050035110010000)

Set DPS (verified): 92.2. Weights run: 8.1s. Verify run: 27.6s. 380 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.459 ± 0.017, spirit=2.982 ± 0.023, mp5=4.518 ± 0.031, crit=0.055 ± 0.005 per rating point (14 rating = 1%, 0.766 per %), spell_haste=not significant (-0.089 ± 0.301)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 50.9 healing_power points (5.29 DPS) | yes | Enduring Cap (3020, -0.69 DPS) [world_drop]; Nightsky Cowl (4039, -1.92 DPS) [world_drop]; Dokebi Hat (14584, -2.07 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | sim-verified (92.2 DPS) | yes | Pendant of Myzrael (4614, +0.00 DPS) [dungeon]; Glowing Green Talisman (5002, +0.00 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.31 DPS) [world_drop] |
| shoulder | Ghostly Mantle (3324) | Deathstalkers in Shadowfang [quest] | 35.8 healing_power points (3.72 DPS) | yes | Nightsky Mantle (4718, -0.96 DPS) [world_drop]; Desert Shoulders (15457, -1.42 DPS) [quest]; Insignia Mantle (4721, -1.55 DPS) [world_drop] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 30.9 healing_power points (3.21 DPS) | yes | Amy's Blanket (13005, -1.04 DPS) [world_drop]; Nightsky Cloak (4719, -1.35 DPS) [world_drop]; Insignia Cloak (4722, -1.35 DPS) [world_drop] |
| chest | Wisdom's Leather Tunic (252511) | Leatherworking [crafted] | 44.6 healing_power points (4.64 DPS) | yes | Pristine Gown (253961, -0.26 DPS) [crafted]; Beguiler Robes (7728, -0.34 DPS) [dungeon]; Pressed Felt Robe (1997, -0.47 DPS) [world] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 20.8 healing_power points (2.17 DPS) | yes | Drakewing Bands (12999, -0.31 DPS) [world_drop]; Nightsky Wristbands (6407, -0.33 DPS) [world_drop]; Dokebi Bracers (14580, -0.47 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 43.1 healing_power points (4.47 DPS) | yes | Tattered Mittens (270030, -1.58 DPS) [quest]; Hotshot Pilot's Gloves (9491, -1.71 DPS) [dungeon]; Blight Gloves (279877, -1.87 DPS) [quest] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 45.6 healing_power points (4.74 DPS) | yes | Silver-lined Belt (13011, -2.14 DPS) [world_drop]; Dokebi Cord (14578, -2.28 DPS) [world_drop]; Lilac Sash (6780, -2.45 DPS) [quest] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 52.1 healing_power points (5.41 DPS) | yes | Earthen Leggings (253999, -0.44 DPS) [crafted]; Pristine Leggings (253987, -0.52 DPS) [crafted]; Stormrider's Leather Kilt (252518, -0.94 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 42.1 healing_power points (4.38 DPS) | yes | Boots of the Enchanter (4325, -1.90 DPS) [crafted]; Resilient Boots (14399, -1.90 DPS) [world_drop]; Stonecloth Boots (14408, -1.90 DPS) [world_drop] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 26.9 healing_power points (2.79 DPS) | yes | The Queen's Jewel (13094, -0.01 DPS) [world_drop]; Solid Gold Band (279845, -0.52 DPS) [quest]; Band of Thorns (5007, -0.93 DPS) [world_drop] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 26.9 healing_power points (2.79 DPS) | yes | The Queen's Jewel (13094, -0.01 DPS) [world_drop]; Solid Gold Band (279845, -0.52 DPS) [quest]; Band of Thorns (5007, -0.93 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gnarled Ash Staff (791, -2.42 DPS) [world_drop]; Staff of the Friar (3415, -2.89 DPS) [dungeon]; Spine of the Devourer (273040, -3.32 DPS) [world] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 30.9 healing_power points (3.21 DPS) | yes | Orb of Mistmantle (13031, -1.04 DPS) [world_drop]; Strength of Will (4837, -1.35 DPS) [vendor]; Silver-thread Rod (15928, -1.66 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Ghostly Mantle; back: Glowing Thresher Cape; chest: Wisdom's Leather Tunic; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 380, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (tauren, 0000000000000000-00000000000000000000-5050035153112000)

Set DPS (verified): 132.5. Weights run: 8.9s. Verify run: 32.1s. 514 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.376 ± 0.021, spirit=2.703 ± 0.020, mp5=5.010 ± 0.038, crit=0.059 ± 0.005 per rating point (14 rating = 1%, 0.822 per %), spell_haste=not significant (1.040 ± 0.404)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 61.7 healing_power points (6.45 DPS) | yes | Whitemane's Chapeau (7720, -0.26 DPS) [dungeon]; Holy Shroud (2721, -1.31 DPS) [world_drop]; Earthen Silk Hood (254015, -1.33 DPS) [crafted] |
| neck | Necklace of Calisea (1714) | Zul'Farrak: Witch Doctor's Chest [dungeon] | 28.6 healing_power points (2.98 DPS) | yes | Amberglow Talisman (10824, -0.16 DPS) [quest]; Glowing Eye of Mordresh (10769, -0.47 DPS) [dungeon]; Triune Amulet (7722, -1.01 DPS) [dungeon] |
| shoulder | Earthen Silk Shoulders (254033) | Tailoring [crafted] | 39.6 healing_power points (4.14 DPS) | yes | Sheepshear Mantle (13115, -0.14 DPS) [world_drop]; Ghostly Mantle (3324, -0.66 DPS) [quest]; Inquisitor's Shawl (19507, -0.70 DPS) [dungeon] |
| back | Cloak of Blight (6832) | Nothing But The Truth [quest] | 35.1 healing_power points (3.67 DPS) | yes | Glowing Thresher Cape (6901, -0.68 DPS) [dungeon]; Ceremonial Centaur Blanket (6789, -0.84 DPS) [quest]; Stargazer Cloak (9660, -0.85 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 79.6 healing_power points (8.32 DPS) | yes | Doomsayer's Robe (4746, -2.19 DPS) [quest]; Dreamweave Vest (10021, -3.07 DPS) [crafted]; Civinad Robes (9623, -3.15 DPS) [quest] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 32.5 healing_power points (3.40 DPS) | yes | Earthen Silk Cuffs (254019, -0.94 DPS) [crafted]; Mindthrust Bracers (1974, -1.11 DPS) [dungeon]; Silkstream Cuffs (16791, -1.55 DPS) [quest] |
| hands | Earthen Silk Gloves (254017) | Tailoring [crafted] | 39.6 healing_power points (4.14 DPS) | yes | Gloves of Old (9395, -0.03 DPS) [world_drop]; Warden's Gloves (14606, -0.46 DPS) [world_drop]; Windchaser Handguards (14431, -0.47 DPS) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 49.2 healing_power points (5.14 DPS) | yes | Mender's Leather Belt (252523, -0.60 DPS) [crafted]; Windchaser Cinch (14435, -1.45 DPS) [world_drop]; Sutarn's Ring (13105, -2.01 DPS) [world_drop] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 62.1 healing_power points (6.49 DPS) | yes | Stoneweaver Leggings (9407, -0.59 DPS) [dungeon]; Stormcloth Pants (10010, -0.68 DPS) [crafted]; Wisdom's Leather Leggings (252519, -1.31 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) | Leatherworking [crafted] | sim-verified (132.6 DPS) | yes | Furen's Boots (13100, +0.00 DPS) [world_drop]; Thoughtcast Boots (10578, -0.98 DPS) [dungeon]; Gilded Slippers (254001, -1.26 DPS) [crafted] |
| finger1 | Welken Ring (5011) | World drop [world_drop] | 25.8 healing_power points (2.70 DPS) | yes | Runemender's Seal (273031, -0.09 DPS) [world]; The Queen's Jewel (13094, -0.15 DPS) [world_drop]; Blush Ember Ring (13093, -0.44 DPS) [world_drop] |
| finger2 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 24.9 healing_power points (2.60 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; The Queen's Jewel (13094, -0.06 DPS) [world_drop]; Blush Ember Ring (13093, -0.34 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -4.64 DPS) [dungeon]; Death Speaker Scepter (2816, -4.97 DPS) [dungeon]; Hypnotic Blade (7714, -5.48 DPS) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 29.2 healing_power points (3.05 DPS) | yes | Beacon of Hope (9393, -0.46 DPS) [dungeon]; Silksand Star (15964, -0.51 DPS) [world_drop]; Aurora Sphere (7610, -0.92 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Necklace of Calisea; shoulder: Earthen Silk Shoulders; back: Cloak of Blight; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Earthen Silk Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; finger1: Welken Ring; finger2: Electrocutioner Lagnut; trinket1: Ankh of Life; main_hand: Hand of Righteousness

No-known-source sample (15 of 514, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 50 (tauren, 4300000000000000-00000000000000000000-5050035153113200)

Set DPS (verified): 225.0. Weights run: 9.1s. Verify run: 31.4s. 684 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.560 ± 0.022, spirit=2.872 ± 0.023, mp5=4.645 ± 0.032, crit=0.077 ± 0.008 per rating point (14 rating = 1%, 1.079 per %), spell_haste=not significant (0.637 ± 0.270)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 87.2 healing_power points (8.76 DPS) | yes | Engineer's Guild Headpiece (9534, -0.51 DPS) [quest]; Blood Guard's Restored Leather Helm (220875, -0.56 DPS) [vendor]; Cassandra's Grace (13102, -1.75 DPS) [world_drop] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 33.3 healing_power points (3.35 DPS) | yes | Necklace of Calisea (1714, -0.23 DPS) [dungeon]; Choker of the High Shaman (4112, -0.26 DPS) [quest]; Amberglow Talisman (10824, -0.46 DPS) [quest] |
| shoulder | Living Shoulders (15061) | Leatherworking [crafted] | 68.3 healing_power points (6.87 DPS) | yes | Blood Guard's Restored Leather Spaulders (220877, -0.20 DPS) [vendor]; Nethergeld Shoulders (254049, -1.27 DPS) [crafted]; Mender's Leather Shoulders (252538, -1.48 DPS, sim-verified) [crafted] |
| back | Cloak of Blight (6832) | Nothing But The Truth [quest] | 37.3 healing_power points (3.75 DPS) | yes | Wingveil Cloak (10802, -0.07 DPS) [dungeon]; Arcane Cloak (8286, -0.29 DPS) [world_drop]; Chainlink Towel (9648, -0.40 DPS) [quest] |
| chest | Ghostweave Vest (14141) | Tailoring [crafted] | 94.6 healing_power points (9.51 DPS) | yes | Feathered Breastplate (8349, -1.04 DPS, sim-verified) [crafted]; Stormcloth Vest (10020, -1.37 DPS) [crafted]; Stone Guard's Restored Leather Jerkin (220879, -1.71 DPS) [vendor] |
| wrist | Mender's Leather Bracers (252543) | Leatherworking [crafted] | 47.2 healing_power points (4.74 DPS) | yes | Nethergeld Cuffs (254061, -0.19 DPS) [crafted]; Enchanted Kodo Bracers (13119, -1.22 DPS) [world_drop]; Venomshroud Armguards (14439, -1.80 DPS) [world_drop] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | 70.9 healing_power points (7.12 DPS) | yes | Gilded Gloves (254095, -0.35 DPS, sim-verified) [crafted]; Earthenweave Gloves (254075, -1.34 DPS) [crafted]; Keeper's Gloves (14666, -2.51 DPS) [world_drop] |
| waist | Earthenweave Cord (254077) | Tailoring [crafted] | 57.6 healing_power points (5.79 DPS) | yes | Mender's Leather Waistguard (252477, -0.89 DPS) [crafted]; Gilded Cord (254037, -1.01 DPS, sim-verified) [crafted]; Mender's Leather Belt (252523, -1.21 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 87.6 healing_power points (8.80 DPS) | yes | Dalewind Trousers (13008, -0.18 DPS) [world_drop]; Kilt of the Atal'ai Prophet (10807, -0.18 DPS) [dungeon]; Jinxed Hoodoo Kilt (9474, -0.56 DPS, sim-verified) [dungeon] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 69.9 healing_power points (7.03 DPS) | yes | First Sergeant's Restored Leather Boots (220885, -0.56 DPS) [vendor]; Earthenweave Boots (254093, -1.24 DPS) [crafted]; Furen's Boots (13100, -1.31 DPS) [world_drop] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 43.8 healing_power points (4.40 DPS) | yes | Welken Ring (5011, -1.60 DPS) [world_drop]; Electrocutioner Lagnut (9447, -1.78 DPS) [dungeon]; Runemender's Seal (273031, -1.78 DPS) [world] |
| finger2 | Coldwater Ring (4550) | Death From Below [quest] | 30.5 healing_power points (3.07 DPS) | yes | Welken Ring (5011, -0.26 DPS) [world_drop]; Electrocutioner Lagnut (9447, -0.44 DPS) [dungeon]; Runemender's Seal (273031, -0.44 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (225.0 DPS) | yes | Uther's Strength (11302, -1.65 DPS) [world_drop]; Alchemist's Stone (13503, -2.31 DPS) [crafted] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-verified (225.0 DPS) | yes | Uther's Strength (11302, -0.45 DPS, sim-verified) [world_drop]; Alchemist's Stone (13503, -1.73 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (225.0 DPS) | yes | Charstone Dirk (17710, -7.90 DPS) [dungeon]; Hand of Righteousness (7721, -10.46 DPS) [dungeon]; Hammer of the Northern Wind (810, -34.42 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Living Shoulders; chest: Ghostweave Vest; wrist: Mender's Leather Bracers; hands: Mender's Leather Gauntlets; waist: Earthenweave Cord; legs: Senior Designer's Pantaloons; feet: Sandals of the Insurgent; finger1: Eye of Adaegus; finger2: Coldwater Ring; trinket2: Evonice's Landin' Pilla; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 684, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60 (tauren, 4532200100000000-00000000000000000000-5050035153113200)

Set DPS (verified): 333.2. Weights run: 9.3s. Verify run: 84.2s. 1564 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.021, intellect=2.344 ± 0.054, spirit=4.387 ± 0.047, mp5=7.227 ± 0.055, crit=0.114 ± 0.012 per rating point (14 rating = 1%, 1.591 per %), spell_haste=not significant (0.465 ± 0.667)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 176.6 healing_power points (18.14 DPS) | yes | Champion's Dragonhide Headdress (227205, -2.70 DPS) [vendor]; Devout Crown (16693, -3.55 DPS) [dungeon]; Wildheart Cowl (16720, -7.35 DPS, sim-verified) [dungeon] |
| neck | Lady Alizabeth's Pendant (13002) | World drop [world_drop] | 65.8 healing_power points (6.76 DPS) | yes | Necklace of Calisea (1714, -1.92 DPS) [dungeon]; Choker of the High Shaman (4112, -1.95 DPS) [quest]; Horizon Choker (13085, -2.17 DPS, sim-verified) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | 99.7 healing_power points (10.24 DPS) | yes | Champion's Dragonhide Pauldrons (227207, +0.00 DPS) [vendor]; Warlord's Dragonhide Pauldrons (231672, +0.00 DPS) [vendor]; Wildheart Spaulders (16718, -4.87 DPS, sim-verified) [dungeon] |
| back | Cloak of Earth and Sky (279270) | Tailoring [crafted] | 61.4 healing_power points (6.31 DPS) | yes | Indomitable Cloak (14683, -0.45 DPS) [world_drop]; Cloak of Blight (6832, -0.59 DPS, sim-verified) [quest]; Wingveil Cloak (10802, -0.84 DPS) [dungeon] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | 156.2 healing_power points (16.05 DPS) | yes | Wildheart Vest (16706, -1.09 DPS) [dungeon]; Ghostweave Vest (14141, -1.10 DPS) [crafted]; Ironfeather Breastplate (15066, -2.15 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | 73.6 healing_power points (7.56 DPS) | yes | General's Dragonhide Bracers (16553, +0.00 DPS) [pvp]; Mender's Leather Bracers (252543, -1.22 DPS) [crafted]; Earthenweave Cuffs (254131, -5.96 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (333.2 DPS) | yes | Swarmtender's Gloves (275607, -0.95 DPS) [crafted]; General's Dragonhide Gauntlets (231676, -1.54 DPS) [vendor]; Wildheart Gloves (16717, -5.67 DPS, sim-verified) [dungeon] |
| waist | Elderwild Waistcord (279252) | Leatherworking [crafted] | 109.0 healing_power points (11.19 DPS) | yes | General's Dragonhide Belt (16556, -0.67 DPS) [pvp]; Wisdom of the Timbermaw (19047, -0.89 DPS, sim-verified) [crafted]; Devout Belt (16696, -1.19 DPS) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 160.1 healing_power points (16.44 DPS) | yes | Legionnaire's Dragonhide Legguards (227206, -1.10 DPS) [vendor]; Living Leggings (15060, -2.50 DPS) [crafted]; Elderwild Pants (279254, -6.03 DPS, sim-verified) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (333.2 DPS) | yes | Wildheart Boots (16715, -0.24 DPS) [dungeon]; Sandals of the Insurgent (13111, -0.26 DPS) [world_drop]; Mooncloth Boots (15802, -4.37 DPS, sim-verified) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 71.7 healing_power points (7.36 DPS) | yes | Coldwater Ring (4550, -2.59 DPS) [quest]; Maiden's Circle (13001, -2.98 DPS) [world_drop]; Welken Ring (5011, -3.01 DPS) [world_drop] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 66.7 healing_power points (6.85 DPS) | yes | Coldwater Ring (4550, -1.90 DPS, sim-verified) [quest]; Maiden's Circle (13001, -2.46 DPS) [world_drop]; Welken Ring (5011, -2.49 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (333.2 DPS) | yes | Evonice's Landin' Pilla (18951, -1.01 DPS, sim-verified) [quest]; Uther's Strength (11302, -3.56 DPS) [world_drop]; Alchemist's Stone (13503, -3.60 DPS) [crafted] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (333.2 DPS) | yes | Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Uther's Strength (11302, -2.60 DPS) [world_drop]; Alchemist's Stone (13503, -2.65 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (333.2 DPS) | yes | Charstone Dirk (17710, -10.40 DPS) [dungeon]; Staff of Hale Magefire (13000, -10.51 DPS) [world_drop]; Hand of Edward the Odd (2243, -13.69 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Lady Alizabeth's Pendant; shoulder: Devout Mantle; back: Cloak of Earth and Sky; chest: Mooncloth Vest; wrist: Devout Bracers; hands: Devout Gloves; waist: Elderwild Waistcord; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket2: Shard of the Splithooves

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60, raid preset (tauren, 4532200100000000-00000000000000000000-5050035153113200)

Set DPS (verified): 635.3. Weights run: 6.7s. Verify run: 54.9s. 1564 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.169, intellect=2.204 ± 0.090, spirit=2.354 ± 0.086, mp5=3.646 ± 0.107, crit=0.324 ± 0.034 per rating point (14 rating = 1%, 4.536 per %), spell_haste=not significant (0.379 ± 1.406)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 145.6 healing_power points (24.28 DPS) | yes | Champion's Dragonhide Headdress (227205, -4.74 DPS) [vendor]; Devout Crown (16693, -6.24 DPS) [dungeon]; Wildheart Cowl (16720, -6.58 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 40.3 healing_power points (6.72 DPS) | yes | Lady Alizabeth's Pendant (13002, -0.83 DPS) [world_drop]; Necklace of Calisea (1714, -1.40 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -1.61 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wildheart Spaulders (16718, +0.00 DPS) [dungeon]; Champion's Dragonhide Pauldrons (227207, +0.00 DPS) [vendor]; Warlord's Dragonhide Pauldrons (231672, +0.00 DPS) [vendor] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 33.7 healing_power points (5.61 DPS) | yes | Cloak of Earth and Sky (279270, -0.12 DPS) [crafted]; Wingveil Cloak (10802, -0.46 DPS) [dungeon]; Cloak of Blight (6832, -0.51 DPS) [quest] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Vest (14138, +0.00 DPS) [crafted]; Warlord's Dragonhide Tunic (231674, -0.36 DPS) [vendor]; Legionnaire's Dragonhide Tunic (227202, -0.36 DPS) [vendor] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | 52.2 healing_power points (8.70 DPS) | yes | General's Dragonhide Bracers (16553, +0.00 DPS) [pvp]; Mender's Leather Bracers (252543, -0.60 DPS) [crafted]; Nethergeld Cuffs (254061, -0.83 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wildheart Gloves (16717, +0.00 DPS) [dungeon]; General's Dragonhide Gauntlets (231676, -0.24 DPS) [vendor]; Mender's Leather Gauntlets (252551, -0.84 DPS) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Dragonhide Belt (16556, +0.00 DPS) [pvp]; Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Elderwild Waistcord (279252, -0.65 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 111.2 healing_power points (18.54 DPS) | yes | Legionnaire's Dragonhide Legguards (227206, +0.00 DPS) [vendor]; General's Dragonhide Legguards (231673, -2.33 DPS) [vendor]; Elderwild Pants (279254, -2.45 DPS) [crafted] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 85.4 healing_power points (14.25 DPS) | yes | Devout Sandals (16691, -2.07 DPS) [dungeon]; Wildheart Boots (16715, -2.44 DPS) [dungeon]; General's Dragonhide Greaves (231671, -2.97 DPS) [vendor] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 47.9 healing_power points (7.99 DPS) | yes | Chemist's Ring (15702, -2.75 DPS) [quest]; Maiden's Circle (13001, -3.06 DPS) [world_drop]; Coldwater Ring (4550, -3.36 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 41.5 healing_power points (6.91 DPS) | yes | Chemist's Ring (15702, -1.67 DPS) [quest]; Maiden's Circle (13001, -1.99 DPS) [world_drop]; Coldwater Ring (4550, -2.28 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (635.3 DPS) | yes | Evonice's Landin' Pilla (18951, -0.78 DPS) [quest]; Shard of the Splithooves (10659, -1.06 DPS) [quest]; Alchemist's Stone (13503, -3.14 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Alchemist's Stone (13503, -1.43 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -14.60 DPS) [dungeon]; Hand of Edward the Odd (2243, -15.49 DPS) [world_drop]; Hand of Righteousness (7721, -20.39 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; shoulder: Devout Mantle; back: Imperial Red Cloak; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket2: Uther's Strength

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

