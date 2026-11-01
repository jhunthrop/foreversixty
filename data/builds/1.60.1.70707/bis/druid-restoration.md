# Leveling BiS: Restoration

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-00000000000000000000-5051000000000000)

Set DPS (verified): 48.5. Weights run: 5.7s. Verify run: 19.6s. 203 eligible items had no known source.

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

No-known-source sample (15 of 203, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18853 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (night-elf, 0000000000000000-00000000000000000000-5053005110010000)

Set DPS (verified): 114.7. Weights run: 7.5s. Verify run: 25.1s. 387 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.658 ± 0.023, spirit=3.736 ± 0.027, mp5=4.831 ± 0.033, crit=0.056 ± 0.004 per rating point (14 rating = 1%, 0.779 per %), spell_haste=not significant (0.459 ± 0.416)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 55.4 healing_power points (6.06 DPS) | yes | Nightsky Cowl (4039, -1.84 DPS) [world_drop]; Enduring Cap (3020, -1.93 DPS, sim-verified) [world_drop]; Stormrider's Leather Helm (252514, -1.98 DPS) [crafted] |
| neck | Pendant of Myzrael (4614) (or Glowing Green Talisman (5002)) | Razorfen Downs: Splinterbone Captain [dungeon] | 22.4 healing_power points (2.45 DPS) | yes | Glowing Green Talisman (5002, +0.00 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.09 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.50 DPS) [world_drop] |
| shoulder | Mantle of Honor (3560) | Bride of the Embalmer [quest] | 37.8 healing_power points (4.13 DPS) | yes | Faerie Mantle (5820, -0.54 DPS, sim-verified) [quest]; Nightsky Mantle (4718, -0.64 DPS) [world_drop]; Insignia Mantle (4721, -1.27 DPS) [world_drop] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 36.9 healing_power points (4.04 DPS) | yes | Prelacy Cape (7004, -0.30 DPS, sim-verified) [quest]; Repairman's Cape (9605, -1.01 DPS) [quest]; Amy's Blanket (13005, -1.17 DPS) [world_drop] |
| chest | Wisdom's Leather Tunic (252511) | Leatherworking [crafted] | 50.4 healing_power points (5.51 DPS) | yes | Beguiler Robes (7728, -0.06 DPS) [dungeon]; Silver-thread Armor (7110, -0.29 DPS) [world_drop]; Pressed Felt Robe (1997, -0.90 DPS, sim-verified) [world] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 22.8 healing_power points (2.49 DPS) | yes | Nightsky Wristbands (6407, -0.18 DPS) [world_drop]; Dokebi Bracers (14580, -0.31 DPS) [world_drop]; Drakewing Bands (12999, -0.66 DPS, sim-verified) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 53.1 healing_power points (5.81 DPS) | yes | Zodiac Gloves (7106, -2.23 DPS) [quest]; Hotshot Pilot's Gloves (9491, -2.32 DPS) [dungeon]; Shilly Mitts (9609, -2.47 DPS, sim-verified) [quest] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 51.4 healing_power points (5.62 DPS) | yes | Dokebi Cord (14578, -2.44 DPS) [world_drop]; Resilient Cord (14406, -2.90 DPS) [world_drop]; Silver-lined Belt (13011, -3.30 DPS, sim-verified) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 58.8 healing_power points (6.43 DPS) | yes | Stormrider's Leather Kilt (252518, -0.66 DPS) [crafted]; Pristine Leggings (253987, -0.71 DPS) [crafted]; Earthen Leggings (253999, -0.80 DPS, sim-verified) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 46.6 healing_power points (5.09 DPS) | yes | Resilient Boots (14399, -1.82 DPS) [world_drop]; Stonecloth Boots (14408, -1.82 DPS) [world_drop]; Boots of the Enchanter (4325, -3.02 DPS, sim-verified) [crafted] |
| finger1 | The Queen's Jewel (13094) | World drop [world_drop] | 33.2 healing_power points (3.63 DPS) | yes | Runemender's Seal (273031, -0.12 DPS) [world]; Ring of Calm (6790, -0.77 DPS) [quest]; Band of Thorns (5007, -1.18 DPS) [world_drop] |
| finger2 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 32.2 healing_power points (3.52 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; Ring of Calm (6790, -0.66 DPS) [quest]; Band of Thorns (5007, -1.07 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (114.6 DPS) | yes | Gnarled Ash Staff (791, -1.31 DPS) [world_drop]; Hardened Root Staff (1317, -1.57 DPS) [quest]; Manual Crowd Pummeler (9449, -4.58 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 35.4 healing_power points (3.88 DPS) | yes | Strength of Will (4837, -1.42 DPS) [vendor]; Orb of Mistmantle (13031, -1.57 DPS, sim-verified) [world_drop]; Torch of Holy Flame (7344, -1.83 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Pendant of Myzrael; shoulder: Mantle of Honor; back: Glowing Thresher Cape; chest: Wisdom's Leather Tunic; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: The Queen's Jewel; finger2: Electrocutioner Lagnut; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 387, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 40 (night-elf, 0000000000000000-00000000000000000000-5053005153112000)

Set DPS (verified): 156.5. Weights run: 8.4s. Verify run: 29.1s. 523 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.611 ± 0.030, spirit=3.632 ± 0.030, mp5=4.696 ± 0.036, crit=0.065 ± 0.006 per rating point (14 rating = 1%, 0.912 per %), spell_haste=not significant (-0.026 ± 0.385)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 74.1 healing_power points (8.06 DPS) | yes | Whitemane's Chapeau (7720, -0.15 DPS) [dungeon]; Imperial Leather Helm (6433, -1.61 DPS) [dungeon]; Earthen Silk Hood (254015, -1.71 DPS) [crafted] |
| neck | Necklace of Calisea (1714) | Zul'Farrak: Witch Doctor's Chest [dungeon] | 36.7 healing_power points (3.99 DPS) | yes | Amberglow Talisman (10824, -0.04 DPS) [quest]; Explorers' League Commendation (7746, -0.44 DPS) [quest]; Glowing Eye of Mordresh (10769, -1.18 DPS) [dungeon] |
| shoulder | Earthen Silk Shoulders (254033) | Tailoring [crafted] | sim-verified (156.5 DPS) | yes | Sheepshear Mantle (13115, +0.00 DPS) [world_drop]; Mantle of Doan (7712, -0.91 DPS) [dungeon]; Crimson Silk Shoulders (7059, -0.95 DPS) [crafted] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | sim-verified (156.5 DPS) | yes | Energy Cloak (9397, +0.00 DPS) [dungeon]; Ceremonial Centaur Blanket (6789, -0.02 DPS) [quest]; Prelacy Cape (7004, -0.35 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 87.1 healing_power points (9.47 DPS) | yes | Doomsayer's Robe (4746, -1.27 DPS) [quest]; Dreamweave Vest (10021, -2.36 DPS) [crafted]; Black Mageweave Robe (10001, -2.54 DPS) [crafted] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 42.8 healing_power points (4.65 DPS) | yes | Earthen Silk Cuffs (254019, -1.59 DPS) [crafted]; Silkstream Cuffs (16791, -2.20 DPS) [quest]; Mindthrust Bracers (1974, -2.24 DPS) [dungeon] |
| hands | Earthen Silk Gloves (254017) | Tailoring [crafted] | sim-verified (156.5 DPS) | yes | Gloves of Old (9395, +0.00 DPS) [world_drop]; Windchaser Handguards (14431, +0.00 DPS) [world_drop]; Warden's Gloves (14606, -0.07 DPS) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 56.7 healing_power points (6.17 DPS) | yes | Mender's Leather Belt (252523, -0.68 DPS) [crafted]; Windchaser Cinch (14435, -1.21 DPS) [world_drop]; Sutarn's Ring (13105, -2.04 DPS) [world_drop] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 76.5 healing_power points (8.32 DPS) | yes | Stormcloth Pants (10010, -0.77 DPS) [crafted]; Stoneweaver Leggings (9407, -0.77 DPS) [dungeon]; Wisdom's Leather Leggings (252519, -2.05 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) | Leatherworking [crafted] | sim-verified (156.5 DPS) | yes | Furen's Boots (13100, +0.00 DPS) [world_drop]; Thoughtcast Boots (10578, -0.60 DPS) [dungeon]; Gilded Slippers (254001, -1.62 DPS) [crafted] |
| finger1 | Welken Ring (5011) | World drop [world_drop] | 33.5 healing_power points (3.64 DPS) | yes | The Queen's Jewel (13094, -0.13 DPS) [world_drop]; Runemender's Seal (273031, -0.22 DPS) [world]; Blush Ember Ring (13093, -0.48 DPS) [world_drop] |
| finger2 | Electrocutioner Lagnut (9447) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (156.5 DPS) | yes | The Queen's Jewel (13094, +0.00 DPS) [world_drop]; Runemender's Seal (273031, +0.00 DPS) [world]; Blush Ember Ring (13093, -0.26 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -4.78 DPS) [dungeon]; Death Speaker Scepter (2816, -5.99 DPS) [dungeon]; Hypnotic Blade (7714, -6.01 DPS) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 34.8 healing_power points (3.78 DPS) | yes | Silksand Star (15964, -0.23 DPS) [world_drop]; Beacon of Hope (9393, -0.68 DPS) [dungeon]; Orb of Lorica (11262, -0.73 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Necklace of Calisea; shoulder: Earthen Silk Shoulders; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Earthen Silk Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; finger1: Welken Ring; trinket1: Ankh of Life; main_hand: Hand of Righteousness

No-known-source sample (15 of 523, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 50 (night-elf, 4300000000000000-00000000000000000000-5053005153113200)

Set DPS (verified): 268.7. Weights run: 8.6s. Verify run: 29.1s. 693 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.011, intellect=1.817 ± 0.054, spirit=3.794 ± 0.051, mp5=5.068 ± 0.065, crit=0.092 ± 0.009 per rating point (14 rating = 1%, 1.286 per %), spell_haste=not significant (-0.235 ± 0.531)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 108.2 healing_power points (11.50 DPS) | yes | Engineer's Guild Headpiece (9534, -0.61 DPS) [quest]; Knight-Lieutenant's Restored Leather Helm (220874, -1.19 DPS) [vendor]; Serpentskin Helm (8261, -2.36 DPS) [world_drop] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 40.6 healing_power points (4.32 DPS) | yes | Necklace of Calisea (1714, -0.14 DPS) [dungeon]; Amberglow Talisman (10824, -0.28 DPS) [quest]; Explorers' League Commendation (7746, -0.69 DPS) [quest] |
| shoulder | Living Shoulders (15061) | Leatherworking [crafted] | 80.3 healing_power points (8.54 DPS) | yes | Knight-Lieutenant's Restored Leather Spaulders (220876, -0.23 DPS) [vendor]; Mender's Leather Shoulders (252538, -1.43 DPS) [crafted]; Ironfeather Shoulders (15067, -1.45 DPS) [crafted] |
| back | Wingveil Cloak (10802) | Sunken Temple: Spawn of Hakkar [dungeon] | 46.7 healing_power points (4.97 DPS) | yes | Arcane Cloak (8286, -0.13 DPS) [world_drop]; Chainlink Towel (9648, -0.36 DPS) [quest]; Traveler's Cloak (8297, -0.53 DPS) [world_drop] |
| chest | Ghostweave Vest (14141) | Tailoring [crafted] | 116.4 healing_power points (12.38 DPS) | yes | Feathered Breastplate (8349, -0.77 DPS) [crafted]; Knight's Restored Leather Jerkin (220878, -2.50 DPS) [vendor]; Stormcloth Vest (10020, -2.99 DPS) [crafted] |
| wrist | Mender's Leather Bracers (252543) | Leatherworking [crafted] | 54.5 healing_power points (5.79 DPS) | yes | Nethergeld Cuffs (254061, -0.30 DPS) [crafted]; Enchanted Kodo Bracers (13119, -0.99 DPS) [world_drop]; Venomshroud Armguards (14439, -1.79 DPS) [world_drop] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | 81.5 healing_power points (8.67 DPS) | yes | Gilded Gloves (254095, -0.42 DPS) [crafted]; Earthenweave Gloves (254075, -1.46 DPS) [crafted]; Keeper's Gloves (14666, -2.51 DPS) [world_drop] |
| waist | Earthenweave Cord (254077) | Tailoring [crafted] | 67.7 healing_power points (7.20 DPS) | yes | Gilded Cord (254037, -0.90 DPS) [crafted]; Mender's Leather Belt (252523, -1.60 DPS) [crafted]; Mender's Leather Waistguard (252477, -1.69 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | sim-verified (268.7 DPS) | yes | Jinxed Hoodoo Kilt (9474, +0.00 DPS) [dungeon]; Kilt of the Atal'ai Prophet (10807, +0.00 DPS) [dungeon]; Dalewind Trousers (13008, +0.00 DPS) [world_drop] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 90.4 healing_power points (9.61 DPS) | yes | Sergeant Major's Restored Leather Boots (220884, -1.52 DPS) [vendor]; Furen's Boots (13100, -1.81 DPS) [world_drop]; Earthenweave Boots (254093, -2.41 DPS) [crafted] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 56.4 healing_power points (6.00 DPS) | yes | Coldwater Ring (4550, -1.79 DPS) [quest]; Seedtime Hoop (9655, -1.97 DPS) [quest]; Welken Ring (5011, -2.21 DPS) [world_drop] |
| finger2 | Choking Band (11868) | Ogre Head On A Stick = Party [quest] | 49.3 healing_power points (5.24 DPS) | yes | Coldwater Ring (4550, -1.03 DPS) [quest]; Seedtime Hoop (9655, -1.21 DPS) [quest]; Welken Ring (5011, -1.45 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, -1.61 DPS) [quest]; Uther's Strength (11302, -2.93 DPS) [world_drop]; Alchemist's Stone (13503, -3.23 DPS) [crafted] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, -0.81 DPS) [quest]; Uther's Strength (11302, -2.12 DPS) [world_drop]; Alchemist's Stone (13503, -2.42 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -9.65 DPS) [dungeon]; Hand of Righteousness (7721, -11.96 DPS) [dungeon]; Spire of Hakkar (10844, -12.14 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Living Shoulders; back: Wingveil Cloak; chest: Ghostweave Vest; wrist: Mender's Leather Bracers; hands: Mender's Leather Gauntlets; waist: Earthenweave Cord; legs: Senior Designer's Pantaloons; feet: Sandals of the Insurgent; finger1: Eye of Adaegus; finger2: Choking Band; trinket2: Evonice's Landin' Pilla; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 693, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60 (night-elf, 4532200100000000-00000000000000000000-5053005153113200)

Set DPS (verified): 387.6. Weights run: 8.8s. Verify run: 78.6s. 1571 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=2.528 ± 0.081, spirit=4.842 ± 0.074, mp5=7.243 ± 0.095, crit=0.112 ± 0.011 per rating point (14 rating = 1%, 1.567 per %), spell_haste=not significant (0.037 ± 0.815)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 186.3 healing_power points (21.18 DPS) | yes | Wildheart Cowl (16720, -2.72 DPS) [dungeon]; Lieutenant Commander's Dragonhide Headdress (227199, -2.91 DPS) [vendor]; Devout Crown (16693, -3.75 DPS) [dungeon] |
| neck | Lady Alizabeth's Pendant (13002) | World drop [world_drop] | 72.6 healing_power points (8.26 DPS) | yes | Horizon Choker (13085, -2.03 DPS) [world_drop]; Necklace of Calisea (1714, -2.39 DPS) [dungeon]; Amberglow Talisman (10824, -2.75 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | 107.7 healing_power points (12.24 DPS) | yes | Lieutenant Commander's Dragonhide Pauldrons (227201, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Pauldrons (231705, +0.00 DPS) [vendor]; Wildheart Spaulders (16718, -0.39 DPS) [dungeon] |
| back | Cloak of Earth and Sky (279270) | Tailoring [crafted] | 67.8 healing_power points (7.71 DPS) | yes | Ethereal Mist Cape (11873, -0.24 DPS) [quest]; Indomitable Cloak (14683, -0.55 DPS) [world_drop]; Sergeant Major's Cape (16337, -1.00 DPS) [pvp] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | 168.6 healing_power points (19.16 DPS) | yes | Ironfeather Breastplate (15066, -0.30 DPS) [crafted]; Wildheart Vest (16706, -1.15 DPS) [dungeon]; Ghostweave Vest (14141, -1.56 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | 79.6 healing_power points (9.05 DPS) | yes | Marshal's Dragonhide Bracers (16445, +0.00 DPS) [pvp]; Earthenweave Cuffs (254131, -1.05 DPS) [crafted]; Mender's Leather Bracers (252543, -1.58 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wildheart Gloves (16717, +0.00 DPS) [dungeon]; Swarmtender's Gloves (275607, -1.11 DPS) [crafted]; Marshal's Dragonhide Gauntlets (231706, -1.96 DPS) [pvp] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Dragonhide Waistguard (16447, +0.00 DPS) [pvp]; Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Elderwild Waistcord (279252, +0.00 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 173.3 healing_power points (19.70 DPS) | yes | Knight-Captain's Dragonhide Legguards (227200, -1.62 DPS) [vendor]; Elderwild Pants (279254, -1.94 DPS) [crafted]; Living Leggings (15060, -2.98 DPS) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; Sandals of the Insurgent (13111, -0.17 DPS) [world_drop]; Wildheart Boots (16715, -0.29 DPS) [dungeon] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 78.5 healing_power points (8.93 DPS) | yes | Choking Band (11868, -1.77 DPS) [quest]; Coldwater Ring (4550, -3.11 DPS) [quest]; Seedtime Hoop (9655, -3.42 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 73.3 healing_power points (8.33 DPS) | yes | Choking Band (11868, -1.17 DPS) [quest]; Coldwater Ring (4550, -2.51 DPS) [quest]; Seedtime Hoop (9655, -2.83 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -1.67 DPS) [quest]; Thunderbrew's Boot Flask (744, -2.20 DPS) [quest]; Alchemist's Stone (13503, -4.40 DPS) [crafted] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-verified (387.4 DPS) | yes | Shard of the Splithooves (10659, -0.56 DPS) [quest]; Thunderbrew's Boot Flask (744, -1.10 DPS) [quest]; Alchemist's Stone (13503, -3.30 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Edward the Odd (2243, -10.87 DPS) [world_drop]; Staff of Hale Magefire (13000, -11.05 DPS) [world_drop]; Charstone Dirk (17710, -12.35 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Lady Alizabeth's Pendant; shoulder: Devout Mantle; back: Cloak of Earth and Sky; chest: Mooncloth Vest; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus

No-known-source sample (15 of 1571, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60, raid preset (night-elf, 4532200100000000-00000000000000000000-5053005153113200)

Set DPS (verified): 722.8. Weights run: 6.1s. Verify run: 46.6s. 1571 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.176, intellect=1.756 ± 0.076, spirit=2.753 ± 0.075, mp5=3.408 ± 0.086, crit=0.373 ± 0.035 per rating point (14 rating = 1%, 5.221 per %), spell_haste=not significant (2.231 ± 1.326)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 143.2 healing_power points (27.51 DPS) | yes | Lieutenant Commander's Dragonhide Headdress (227199, -5.47 DPS) [vendor]; Devout Crown (16693, -7.63 DPS) [dungeon]; Wildheart Cowl (16720, -8.56 DPS, sim-verified) [dungeon] |
| neck | Lady Alizabeth's Pendant (13002) | World drop [world_drop] | 41.3 healing_power points (7.94 DPS) | yes | Necklace of Calisea (1714, -1.87 DPS) [dungeon]; Amberglow Talisman (10824, -2.65 DPS) [quest]; Horizon Choker (13085, -2.65 DPS, sim-verified) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (722.9 DPS) | yes | Lieutenant Commander's Dragonhide Pauldrons (227201, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Pauldrons (231705, +0.00 DPS) [vendor]; Wildheart Spaulders (16718, -9.84 DPS, sim-verified) [dungeon] |
| back | Cloak of Earth and Sky (279270) | Tailoring [crafted] | 38.5 healing_power points (7.41 DPS) | yes | Sergeant Major's Cape (16337, -0.47 DPS) [pvp]; Indomitable Cloak (14683, -0.53 DPS) [world_drop]; Ethereal Mist Cape (11873, -0.83 DPS, sim-verified) [quest] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | 113.4 healing_power points (21.80 DPS) | yes | Devout Robe (16690, +0.00 DPS, sim-verified) [dungeon]; Knight-Captain's Dragonhide Tunic (227196, -1.64 DPS) [pvp]; Field Marshal's Dragonhide Tunic (231702, -2.07 DPS) [vendor] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | 52.6 healing_power points (10.10 DPS) | yes | Marshal's Dragonhide Bracers (16445, +0.00 DPS) [pvp]; Mender's Leather Bracers (252543, -0.92 DPS) [crafted]; Earthenweave Cuffs (254131, -9.81 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (722.9 DPS) | yes | Marshal's Dragonhide Gauntlets (231706, -1.25 DPS) [pvp]; Mender's Leather Gauntlets (252551, -1.50 DPS) [crafted]; Wildheart Gloves (16717, -9.81 DPS, sim-verified) [dungeon] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 80.4 healing_power points (15.44 DPS) | yes | Elderwild Waistcord (279252, -0.16 DPS) [crafted]; Marshal's Dragonhide Waistguard (16447, -1.72 DPS) [pvp]; Devout Belt (16696, -1.82 DPS) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 113.7 healing_power points (21.84 DPS) | yes | Knight-Captain's Dragonhide Legguards (227200, +0.00 DPS) [vendor]; Marshal's Dragonhide Legguards (231703, -3.46 DPS) [pvp]; Elderwild Pants (279254, -9.87 DPS, sim-verified) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (722.9 DPS) | yes | Wildheart Boots (16715, -0.34 DPS) [dungeon]; Sandals of the Insurgent (13111, -1.20 DPS) [world_drop]; Mooncloth Boots (15802, -9.34 DPS, sim-verified) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 47.8 healing_power points (9.19 DPS) | yes | Choking Band (11868, -2.32 DPS) [quest]; Coldwater Ring (4550, -3.42 DPS) [quest]; Maiden's Circle (13001, -3.66 DPS) [world_drop] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 43.6 healing_power points (8.37 DPS) | yes | Choking Band (11868, -1.04 DPS, sim-verified) [quest]; Coldwater Ring (4550, -2.60 DPS) [quest]; Maiden's Circle (13001, -2.84 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (722.9 DPS) | yes | Evonice's Landin' Pilla (18951, -1.06 DPS) [quest]; Uther's Strength (11302, -1.98 DPS, sim-verified) [world_drop]; Thunderbrew's Boot Flask (744, -2.12 DPS) [quest] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (722.9 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Uther's Strength (11302, -0.47 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (722.9 DPS) | yes | Hand of Edward the Odd (2243, -14.99 DPS, sim-verified) [world_drop]; Charstone Dirk (17710, -16.06 DPS) [dungeon]; Hand of Righteousness (7721, -21.08 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Lady Alizabeth's Pendant; shoulder: Devout Mantle; back: Cloak of Earth and Sky; chest: Mooncloth Vest; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket2: Shard of the Splithooves

No-known-source sample (15 of 1571, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

## Horde

### Band 20 (tauren, 0000000000000000-00000000000000000000-5051000000000000)

Set DPS (verified): 47.8. Weights run: 5.7s. Verify run: 18.5s. 199 eligible items had no known source.

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

No-known-source sample (15 of 199, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion; 21566 Rune of Perfection

### Band 30 (tauren, 0000000000000000-00000000000000000000-5053005110010000)

Set DPS (verified): 114.1. Weights run: 7.5s. Verify run: 25.5s. 387 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.658 ± 0.023, spirit=3.736 ± 0.027, mp5=4.831 ± 0.033, crit=0.056 ± 0.004 per rating point (14 rating = 1%, 0.779 per %), spell_haste=not significant (0.459 ± 0.416)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 55.4 healing_power points (6.06 DPS) | yes | Enduring Cap (3020, -0.30 DPS) [world_drop]; Nightsky Cowl (4039, -1.84 DPS) [world_drop]; Stormrider's Leather Helm (252514, -1.98 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | sim-verified (114.1 DPS) | yes | Pendant of Myzrael (4614, +0.00 DPS) [dungeon]; Glowing Green Talisman (5002, +0.00 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.41 DPS) [world_drop] |
| shoulder | Ghostly Mantle (3324) | Deathstalkers in Shadowfang [quest] | 42.6 healing_power points (4.66 DPS) | yes | Nightsky Mantle (4718, -1.17 DPS) [world_drop]; Desert Shoulders (15457, -1.71 DPS) [quest]; Insignia Mantle (4721, -1.80 DPS) [world_drop] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 36.9 healing_power points (4.04 DPS) | yes | Amy's Blanket (13005, -1.17 DPS) [world_drop]; Nightsky Cloak (4719, -1.58 DPS) [world_drop]; Insignia Cloak (4722, -1.58 DPS) [world_drop] |
| chest | Wisdom's Leather Tunic (252511) | Leatherworking [crafted] | 50.4 healing_power points (5.51 DPS) | yes | Pressed Felt Robe (1997, -0.06 DPS) [world]; Beguiler Robes (7728, -0.06 DPS) [dungeon]; Silver-thread Armor (7110, -0.29 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 22.8 healing_power points (2.49 DPS) | yes | Drakewing Bands (12999, -0.04 DPS) [world_drop]; Nightsky Wristbands (6407, -0.18 DPS) [world_drop]; Dokebi Bracers (14580, -0.31 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 53.1 healing_power points (5.81 DPS) | yes | Tattered Mittens (270030, -2.27 DPS) [quest]; Hotshot Pilot's Gloves (9491, -2.32 DPS) [dungeon]; Blight Gloves (279877, -2.50 DPS) [quest] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 51.4 healing_power points (5.62 DPS) | yes | Silver-lined Belt (13011, -2.35 DPS) [world_drop]; Dokebi Cord (14578, -2.44 DPS) [world_drop]; Lilac Sash (6780, -2.76 DPS) [quest] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 58.8 healing_power points (6.43 DPS) | yes | Earthen Leggings (253999, -0.29 DPS) [crafted]; Stormrider's Leather Kilt (252518, -0.66 DPS) [crafted]; Pristine Leggings (253987, -0.71 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 46.6 healing_power points (5.09 DPS) | yes | Boots of the Enchanter (4325, -1.82 DPS) [crafted]; Resilient Boots (14399, -1.82 DPS) [world_drop]; Stonecloth Boots (14408, -1.82 DPS) [world_drop] |
| finger1 | Electrocutioner Lagnut (9447) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (114.1 DPS) | yes | The Queen's Jewel (13094, +0.00 DPS) [world_drop]; Solid Gold Band (279845, -0.71 DPS) [quest]; Band of Thorns (5007, -1.07 DPS) [world_drop] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | sim-verified (114.1 DPS) | yes | The Queen's Jewel (13094, +0.00 DPS) [world_drop]; Solid Gold Band (279845, -0.71 DPS) [quest]; Band of Thorns (5007, -1.07 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gnarled Ash Staff (791, -1.31 DPS) [world_drop]; Staff of the Friar (3415, -1.99 DPS) [dungeon]; Odo's Ley Staff (6318, -2.53 DPS) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 35.4 healing_power points (3.88 DPS) | yes | Orb of Mistmantle (13031, -1.26 DPS) [world_drop]; Strength of Will (4837, -1.42 DPS) [vendor]; Silver-thread Rod (15928, -1.83 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Ghostly Mantle; back: Glowing Thresher Cape; chest: Wisdom's Leather Tunic; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 387, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (tauren, 0000000000000000-00000000000000000000-5053005153112000)

Set DPS (verified): 156.6. Weights run: 8.4s. Verify run: 29.2s. 521 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.611 ± 0.030, spirit=3.632 ± 0.030, mp5=4.696 ± 0.036, crit=0.065 ± 0.006 per rating point (14 rating = 1%, 0.912 per %), spell_haste=not significant (-0.026 ± 0.385)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 74.1 healing_power points (8.06 DPS) | yes | Whitemane's Chapeau (7720, -0.15 DPS) [dungeon]; Imperial Leather Helm (6433, -1.61 DPS) [dungeon]; Earthen Silk Hood (254015, -1.71 DPS) [crafted] |
| neck | Necklace of Calisea (1714) | Zul'Farrak: Witch Doctor's Chest [dungeon] | 36.7 healing_power points (3.99 DPS) | yes | Amberglow Talisman (10824, -0.04 DPS) [quest]; Glowing Eye of Mordresh (10769, -1.18 DPS) [dungeon]; Triune Amulet (7722, -1.23 DPS) [dungeon] |
| shoulder | Earthen Silk Shoulders (254033) | Tailoring [crafted] | sim-verified (156.6 DPS) | yes | Sheepshear Mantle (13115, +0.00 DPS) [world_drop]; Ghostly Mantle (3324, -0.58 DPS) [quest]; Mantle of Doan (7712, -0.91 DPS) [dungeon] |
| back | Cloak of Blight (6832) | Nothing But The Truth [quest] | 47.2 healing_power points (5.14 DPS) | yes | Energy Cloak (9397, -1.19 DPS) [dungeon]; Stargazer Cloak (9660, -1.19 DPS) [quest]; Glowing Thresher Cape (6901, -1.21 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 87.1 healing_power points (9.47 DPS) | yes | Doomsayer's Robe (4746, -1.27 DPS) [quest]; Dreamweave Vest (10021, -2.36 DPS) [crafted]; Black Mageweave Robe (10001, -2.54 DPS) [crafted] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 42.8 healing_power points (4.65 DPS) | yes | Earthen Silk Cuffs (254019, -1.59 DPS) [crafted]; Silkstream Cuffs (16791, -2.20 DPS) [quest]; Mindthrust Bracers (1974, -2.24 DPS) [dungeon] |
| hands | Earthen Silk Gloves (254017) | Tailoring [crafted] | sim-verified (156.6 DPS) | yes | Gloves of Old (9395, +0.00 DPS) [world_drop]; Windchaser Handguards (14431, +0.00 DPS) [world_drop]; Warden's Gloves (14606, -0.07 DPS) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 56.7 healing_power points (6.17 DPS) | yes | Mender's Leather Belt (252523, -0.68 DPS) [crafted]; Windchaser Cinch (14435, -1.21 DPS) [world_drop]; Sutarn's Ring (13105, -2.04 DPS) [world_drop] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 76.5 healing_power points (8.32 DPS) | yes | Stormcloth Pants (10010, -0.77 DPS) [crafted]; Stoneweaver Leggings (9407, -0.77 DPS) [dungeon]; Mistspray Kilt (4976, -1.61 DPS) [quest] |
| feet | Mender's Leather Shoes (252533) | Leatherworking [crafted] | sim-verified (156.6 DPS) | yes | Furen's Boots (13100, +0.00 DPS) [world_drop]; Thoughtcast Boots (10578, -0.60 DPS) [dungeon]; Gilded Slippers (254001, -1.62 DPS) [crafted] |
| finger1 | Welken Ring (5011) | World drop [world_drop] | 33.5 healing_power points (3.64 DPS) | yes | The Queen's Jewel (13094, -0.13 DPS) [world_drop]; Runemender's Seal (273031, -0.22 DPS) [world]; Blush Ember Ring (13093, -0.48 DPS) [world_drop] |
| finger2 | Electrocutioner Lagnut (9447) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (156.6 DPS) | yes | The Queen's Jewel (13094, +0.00 DPS) [world_drop]; Runemender's Seal (273031, +0.00 DPS) [world]; Blush Ember Ring (13093, -0.26 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -4.78 DPS) [dungeon]; Death Speaker Scepter (2816, -5.99 DPS) [dungeon]; Hypnotic Blade (7714, -6.01 DPS) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 34.8 healing_power points (3.78 DPS) | yes | Silksand Star (15964, -0.23 DPS) [world_drop]; Beacon of Hope (9393, -0.68 DPS) [dungeon]; Traitor's Finger (6774, -0.93 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Necklace of Calisea; shoulder: Earthen Silk Shoulders; back: Cloak of Blight; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Earthen Silk Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; finger1: Welken Ring; finger2: Electrocutioner Lagnut; trinket1: Ankh of Life; main_hand: Hand of Righteousness

No-known-source sample (15 of 521, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 50 (tauren, 4300000000000000-00000000000000000000-5053005153113200)

Set DPS (verified): 266.5. Weights run: 8.6s. Verify run: 29.0s. 691 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.011, intellect=1.817 ± 0.054, spirit=3.794 ± 0.051, mp5=5.068 ± 0.065, crit=0.092 ± 0.009 per rating point (14 rating = 1%, 1.286 per %), spell_haste=not significant (-0.235 ± 0.531)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 108.2 healing_power points (11.50 DPS) | yes | Engineer's Guild Headpiece (9534, -0.61 DPS) [quest]; Blood Guard's Restored Leather Helm (220875, -1.19 DPS) [vendor]; Serpentskin Helm (8261, -2.36 DPS) [world_drop] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 40.6 healing_power points (4.32 DPS) | yes | Choker of the High Shaman (4112, -0.13 DPS) [quest]; Necklace of Calisea (1714, -0.14 DPS) [dungeon]; Amberglow Talisman (10824, -0.28 DPS) [quest] |
| shoulder | Living Shoulders (15061) | Leatherworking [crafted] | 80.3 healing_power points (8.54 DPS) | yes | Blood Guard's Restored Leather Spaulders (220877, -0.23 DPS) [vendor]; Mender's Leather Shoulders (252538, -1.43 DPS) [crafted]; Ironfeather Shoulders (15067, -1.45 DPS) [crafted] |
| back | Cloak of Blight (6832) | Nothing But The Truth [quest] | 49.3 healing_power points (5.24 DPS) | yes | Wingveil Cloak (10802, -0.28 DPS) [dungeon]; Arcane Cloak (8286, -0.40 DPS) [world_drop]; Chainlink Towel (9648, -0.63 DPS) [quest] |
| chest | Ghostweave Vest (14141) | Tailoring [crafted] | 116.4 healing_power points (12.38 DPS) | yes | Feathered Breastplate (8349, -0.77 DPS) [crafted]; Stone Guard's Restored Leather Jerkin (220879, -2.50 DPS) [vendor]; Stormcloth Vest (10020, -2.99 DPS) [crafted] |
| wrist | Mender's Leather Bracers (252543) | Leatherworking [crafted] | 54.5 healing_power points (5.79 DPS) | yes | Nethergeld Cuffs (254061, -0.30 DPS) [crafted]; Enchanted Kodo Bracers (13119, -0.99 DPS) [world_drop]; Venomshroud Armguards (14439, -1.79 DPS) [world_drop] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | 81.5 healing_power points (8.67 DPS) | yes | Gilded Gloves (254095, -0.42 DPS) [crafted]; Earthenweave Gloves (254075, -1.46 DPS) [crafted]; Keeper's Gloves (14666, -2.51 DPS) [world_drop] |
| waist | Earthenweave Cord (254077) | Tailoring [crafted] | 67.7 healing_power points (7.20 DPS) | yes | Gilded Cord (254037, -0.90 DPS) [crafted]; Mender's Leather Belt (252523, -1.60 DPS) [crafted]; Mender's Leather Waistguard (252477, -1.69 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | sim-verified (266.6 DPS) | yes | Jinxed Hoodoo Kilt (9474, +0.00 DPS) [dungeon]; Kilt of the Atal'ai Prophet (10807, +0.00 DPS) [dungeon]; Dalewind Trousers (13008, +0.00 DPS) [world_drop] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 90.4 healing_power points (9.61 DPS) | yes | First Sergeant's Restored Leather Boots (220885, -1.52 DPS) [vendor]; Furen's Boots (13100, -1.81 DPS) [world_drop]; Earthenweave Boots (254093, -2.41 DPS) [crafted] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 56.4 healing_power points (6.00 DPS) | yes | Welken Ring (5011, -2.21 DPS) [world_drop]; The Queen's Jewel (13094, -2.39 DPS) [world_drop]; Electrocutioner Lagnut (9447, -2.54 DPS) [dungeon] |
| finger2 | Coldwater Ring (4550) | Death From Below [quest] | 39.6 healing_power points (4.21 DPS) | yes | Welken Ring (5011, -0.42 DPS) [world_drop]; The Queen's Jewel (13094, -0.60 DPS) [world_drop]; Electrocutioner Lagnut (9447, -0.75 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -2.93 DPS) [world_drop]; Alchemist's Stone (13503, -3.23 DPS) [crafted] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -2.12 DPS) [world_drop]; Alchemist's Stone (13503, -2.42 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -9.65 DPS) [dungeon]; Hand of Righteousness (7721, -11.96 DPS) [dungeon]; Spire of Hakkar (10844, -12.14 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Living Shoulders; chest: Ghostweave Vest; wrist: Mender's Leather Bracers; hands: Mender's Leather Gauntlets; waist: Earthenweave Cord; legs: Senior Designer's Pantaloons; feet: Sandals of the Insurgent; finger1: Eye of Adaegus; finger2: Coldwater Ring; trinket2: Evonice's Landin' Pilla; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 691, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60 (tauren, 4532200100000000-00000000000000000000-5053005153113200)

Set DPS (verified): 384.9. Weights run: 8.8s. Verify run: 78.4s. 1571 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=2.528 ± 0.081, spirit=4.842 ± 0.074, mp5=7.243 ± 0.095, crit=0.112 ± 0.011 per rating point (14 rating = 1%, 1.567 per %), spell_haste=not significant (0.037 ± 0.815)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 186.3 healing_power points (21.18 DPS) | yes | Champion's Dragonhide Headdress (227205, -2.91 DPS) [vendor]; Devout Crown (16693, -3.75 DPS) [dungeon]; Wildheart Cowl (16720, -7.53 DPS, sim-verified) [dungeon] |
| neck | Lady Alizabeth's Pendant (13002) | World drop [world_drop] | 72.6 healing_power points (8.26 DPS) | yes | Necklace of Calisea (1714, -2.39 DPS) [dungeon]; Choker of the High Shaman (4112, -2.42 DPS) [quest]; Horizon Choker (13085, -3.04 DPS, sim-verified) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | 107.7 healing_power points (12.24 DPS) | yes | Champion's Dragonhide Pauldrons (227207, +0.00 DPS) [vendor]; Warlord's Dragonhide Pauldrons (231672, +0.00 DPS) [vendor]; Wildheart Spaulders (16718, -6.02 DPS, sim-verified) [dungeon] |
| back | Cloak of Earth and Sky (279270) | Tailoring [crafted] | 67.8 healing_power points (7.71 DPS) | yes | Cloak of Blight (6832, -0.54 DPS, sim-verified) [quest]; Indomitable Cloak (14683, -0.55 DPS) [world_drop]; Wingveil Cloak (10802, -1.08 DPS) [dungeon] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | 168.6 healing_power points (19.16 DPS) | yes | Wildheart Vest (16706, -1.15 DPS) [dungeon]; Ghostweave Vest (14141, -1.56 DPS) [crafted]; Ironfeather Breastplate (15066, -1.90 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | 79.6 healing_power points (9.05 DPS) | yes | General's Dragonhide Bracers (16553, +0.00 DPS) [pvp]; Mender's Leather Bracers (252543, -1.58 DPS) [crafted]; Earthenweave Cuffs (254131, -6.23 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (384.8 DPS) | yes | Swarmtender's Gloves (275607, -1.11 DPS) [crafted]; General's Dragonhide Gauntlets (231676, -1.96 DPS) [vendor]; Wildheart Gloves (16717, -4.95 DPS, sim-verified) [dungeon] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (384.8 DPS) | yes | General's Dragonhide Belt (16556, +0.00 DPS) [pvp]; Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Elderwild Waistcord (279252, -1.70 DPS, sim-verified) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 173.3 healing_power points (19.70 DPS) | yes | Legionnaire's Dragonhide Legguards (227206, -1.62 DPS) [vendor]; Living Leggings (15060, -2.98 DPS) [crafted]; Elderwild Pants (279254, -6.42 DPS, sim-verified) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (384.8 DPS) | yes | Sandals of the Insurgent (13111, -0.17 DPS) [world_drop]; Wildheart Boots (16715, -0.29 DPS) [dungeon]; Mooncloth Boots (15802, -5.21 DPS, sim-verified) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 78.5 healing_power points (8.93 DPS) | yes | Coldwater Ring (4550, -3.11 DPS) [quest]; Maiden's Circle (13001, -3.61 DPS) [world_drop]; Welken Ring (5011, -3.64 DPS) [world_drop] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 73.3 healing_power points (8.33 DPS) | yes | Coldwater Ring (4550, -1.97 DPS, sim-verified) [quest]; Maiden's Circle (13001, -3.02 DPS) [world_drop]; Welken Ring (5011, -3.04 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (384.8 DPS) | yes | Evonice's Landin' Pilla (18951, -1.13 DPS, sim-verified) [quest]; Alchemist's Stone (13503, -4.40 DPS) [crafted]; Uther's Strength (11302, -4.56 DPS) [world_drop] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (384.8 DPS) | yes | Evonice's Landin' Pilla (18951, +0.00 DPS, sim-verified) [quest]; Alchemist's Stone (13503, -2.74 DPS) [crafted]; Uther's Strength (11302, -2.89 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (384.8 DPS) | yes | Staff of Hale Magefire (13000, -11.05 DPS) [world_drop]; Charstone Dirk (17710, -12.35 DPS) [dungeon]; Hand of Edward the Odd (2243, -15.11 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Lady Alizabeth's Pendant; shoulder: Devout Mantle; back: Cloak of Earth and Sky; chest: Mooncloth Vest; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket2: Shard of the Splithooves

No-known-source sample (15 of 1571, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60, raid preset (tauren, 4532200100000000-00000000000000000000-5053005153113200)

Set DPS (verified): 721.4. Weights run: 6.1s. Verify run: 46.3s. 1571 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.176, intellect=1.756 ± 0.076, spirit=2.753 ± 0.075, mp5=3.408 ± 0.086, crit=0.373 ± 0.035 per rating point (14 rating = 1%, 5.221 per %), spell_haste=not significant (2.231 ± 1.326)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 143.2 healing_power points (27.51 DPS) | yes | Champion's Dragonhide Headdress (227205, -5.47 DPS) [vendor]; Devout Crown (16693, -7.63 DPS) [dungeon]; Wildheart Cowl (16720, -7.82 DPS, sim-verified) [dungeon] |
| neck | Lady Alizabeth's Pendant (13002) | World drop [world_drop] | 41.3 healing_power points (7.94 DPS) | yes | Necklace of Calisea (1714, -1.87 DPS) [dungeon]; Choker of the High Shaman (4112, -2.02 DPS) [quest]; Horizon Choker (13085, -2.27 DPS, sim-verified) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (721.5 DPS) | yes | Champion's Dragonhide Pauldrons (227207, +0.00 DPS) [vendor]; Warlord's Dragonhide Pauldrons (231672, +0.00 DPS) [vendor]; Wildheart Spaulders (16718, -9.29 DPS, sim-verified) [dungeon] |
| back | Cloak of Earth and Sky (279270) | Tailoring [crafted] | 38.5 healing_power points (7.41 DPS) | yes | Cloak of Blight (6832, -0.53 DPS) [quest]; Indomitable Cloak (14683, -0.53 DPS) [world_drop]; Wingveil Cloak (10802, -0.63 DPS) [dungeon] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | 113.4 healing_power points (21.80 DPS) | yes | Devout Robe (16690, +0.00 DPS, sim-verified) [dungeon]; Legionnaire's Dragonhide Tunic (227202, -1.64 DPS) [vendor]; Warlord's Dragonhide Tunic (231674, -2.07 DPS) [vendor] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | 52.6 healing_power points (10.10 DPS) | yes | General's Dragonhide Bracers (16553, +0.00 DPS) [pvp]; Mender's Leather Bracers (252543, -0.92 DPS) [crafted]; Earthenweave Cuffs (254131, -9.41 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (721.5 DPS) | yes | General's Dragonhide Gauntlets (231676, -1.25 DPS) [vendor]; Mender's Leather Gauntlets (252551, -1.50 DPS) [crafted]; Wildheart Gloves (16717, -9.22 DPS, sim-verified) [dungeon] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 80.4 healing_power points (15.44 DPS) | yes | Elderwild Waistcord (279252, +0.00 DPS, sim-verified) [crafted]; General's Dragonhide Belt (16556, -1.72 DPS) [pvp]; Devout Belt (16696, -1.82 DPS) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 113.7 healing_power points (21.84 DPS) | yes | Legionnaire's Dragonhide Legguards (227206, +0.00 DPS) [vendor]; General's Dragonhide Legguards (231673, -3.46 DPS) [vendor]; Elderwild Pants (279254, -9.38 DPS, sim-verified) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (721.5 DPS) | yes | Wildheart Boots (16715, -0.34 DPS) [dungeon]; Sandals of the Insurgent (13111, -1.20 DPS) [world_drop]; Mooncloth Boots (15802, -8.62 DPS, sim-verified) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 47.8 healing_power points (9.19 DPS) | yes | Coldwater Ring (4550, -3.42 DPS) [quest]; Maiden's Circle (13001, -3.66 DPS) [world_drop]; Chemist's Ring (15702, -3.70 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 43.6 healing_power points (8.37 DPS) | yes | Coldwater Ring (4550, -1.14 DPS, sim-verified) [quest]; Maiden's Circle (13001, -2.84 DPS) [world_drop]; Chemist's Ring (15702, -2.88 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (721.5 DPS) | yes | Evonice's Landin' Pilla (18951, -0.68 DPS, sim-verified) [quest]; Uther's Strength (11302, -2.89 DPS) [world_drop]; Alchemist's Stone (13503, -4.23 DPS) [crafted] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (721.5 DPS) | yes | Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Uther's Strength (11302, -0.47 DPS) [world_drop]; Alchemist's Stone (13503, -1.81 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (721.5 DPS) | yes | Hand of Edward the Odd (2243, -14.07 DPS, sim-verified) [world_drop]; Charstone Dirk (17710, -16.06 DPS) [dungeon]; Hand of Righteousness (7721, -21.08 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Lady Alizabeth's Pendant; shoulder: Devout Mantle; back: Cloak of Earth and Sky; chest: Mooncloth Vest; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket2: Shard of the Splithooves

No-known-source sample (15 of 1571, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

