# Leveling BiS: Restoration

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-00000000000000000000-5050010000000000)

Set DPS (verified): 44.7. Weights run: 5.7s. Verify run: 18.0s. 196 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.000, intellect=1.289 ± 0.004, spirit=1.552 ± 0.008, mp5=3.307 ± 0.015, crit=0.044 ± 0.003 per rating point (14 rating = 1%, 0.609 per %), spell_haste=-5.588 ± 0.106

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.85 DPS) | yes | Wisdom's Leather Hood (252507, -0.13 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.23 DPS) [crafted]; Stormrider's Leather Hood (252506, -0.42 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 6.2 healing_power points (0.64 DPS) | yes | Scholarly Pendant (277203, -0.55 DPS, sim-verified) [quest] |
| shoulder | Slime-encrusted Pads (6461) | Wailing Caverns: Mutanus the Devourer [dungeon] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Forest Leather Mantle (4709, -0.38 DPS) [world_drop]; Prospector's Pads (14566, -0.38 DPS) [world_drop]; Magician's Mantle (12998, -0.76 DPS, sim-verified) [world_drop] |
| back | Spirit Cloak (4792) (or Sylvan Cloak (4793), Heavy Woolen Cloak (4311)) | Jennabink Powerseam [vendor] | 6.2 healing_power points (0.64 DPS) | yes | Heavy Woolen Cloak (4311, +0.00 DPS) [crafted]; Sylvan Cloak (4793, +0.00 DPS) [vendor]; Seer's Cape (6378, -0.05 DPS) [dungeon] |
| chest | Robe of the Moccasin (6465) | Wailing Caverns: Lord Cobrahn [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, -0.04 DPS) [crafted]; Corsair's Overshirt (5202, -0.11 DPS) [dungeon]; Wisdom's Leather Armor (252493, -0.57 DPS, sim-verified) [crafted] |
| wrist | Drakewing Bands (12999) | World drop [world_drop] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Owl Bracers (4796, -0.30 DPS) [vendor]; Bravo's Armbands (270015, -0.32 DPS) [quest]; Mindthrust Bracers (1974, -0.49 DPS, sim-verified) [dungeon] |
| hands | Wisdom's Leather Gloves (252499) | Leatherworking [crafted] | 15.2 healing_power points (1.56 DPS) | yes | Pristine Gloves (253913, +0.00 DPS, sim-verified) [crafted]; Magefist Gloves (12977, -0.26 DPS) [world_drop]; Bright Gloves (3066, -0.39 DPS) [world_drop] |
| waist | Novice Ardent's Sash (253887) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Wisdom's Leather Belt (252433, -0.14 DPS) [crafted]; Keller's Girdle (2911, -0.43 DPS) [world_drop]; Pristine Sash (253925, -0.82 DPS, sim-verified) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 33.9 healing_power points (3.50 DPS) | yes | Filigreed Pristine Leggings (253937, -0.26 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -1.61 DPS) [world_drop]; Dreamer's Leggings (270016, -2.04 DPS) [quest] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 17.4 healing_power points (1.80 DPS) | yes | Black Whelp Slippers (252424, +0.00 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.47 DPS) [crafted]; Kimbra Boots (6191, -0.79 DPS) [quest] |
| finger1 | Band of Purification (12996) | World drop [world_drop] | 9.3 healing_power points (0.96 DPS) | yes | Lavishly Jeweled Ring (1156, -0.16 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.56 DPS) [world_drop]; Demon Band (12054, -0.64 DPS) [world_drop] |
| finger2 | Deep Fathom Ring (6463) | Wailing Caverns: Mutanus the Devourer [dungeon] | 7.8 healing_power points (0.80 DPS) | yes | Lavishly Jeweled Ring (1156, -0.11 DPS, sim-verified) [dungeon]; Volcanic Rock Ring (12053, -0.40 DPS) [world_drop]; Demon Band (12054, -0.48 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Westfall (2042) | The Defias Brotherhood [quest] | 15.8 healing_power points (1.62 DPS) | yes | Staff of the Blessed Seer (2271, +0.00 DPS, sim-verified) [dungeon]; Twisted Chanter's Staff (890, -0.30 DPS) [world_drop]; Evocator's Blade (2567, -0.56 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Tarnished Locket; shoulder: Slime-encrusted Pads; back: Spirit Cloak; chest: Robe of the Moccasin; wrist: Drakewing Bands; hands: Wisdom's Leather Gloves; waist: Novice Ardent's Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Band of Purification; finger2: Deep Fathom Ring; main_hand: Staff of Westfall

No-known-source sample (15 of 196, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18853 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (night-elf, 0000000000000000-00000000000000000000-5050035110010000)

Set DPS (verified): 89.0. Weights run: 8.0s. Verify run: 27.5s. 380 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.454 ± 0.014, spirit=2.917 ± 0.018, mp5=4.441 ± 0.027, crit=0.052 ± 0.005 per rating point (14 rating = 1%, 0.735 per %), spell_haste=not significant (-1.023 ± 0.259)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 50.5 healing_power points (5.14 DPS) | yes | Nightsky Cowl (4039, -1.88 DPS) [world_drop]; Dokebi Hat (14584, -2.03 DPS) [world_drop]; Enduring Cap (3020, -2.46 DPS, sim-verified) [world_drop] |
| neck | Pendant of Myzrael (4614) (or Glowing Green Talisman (5002)) | Razorfen Downs: Splinterbone Captain [dungeon] | 17.5 healing_power points (1.78 DPS) | yes | Glowing Green Talisman (5002, +0.00 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.00 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.30 DPS) [world_drop] |
| shoulder | Mantle of Honor (3560) | Bride of the Embalmer [quest] | 30.6 healing_power points (3.12 DPS) | yes | Faerie Mantle (5820, -0.42 DPS, sim-verified) [quest]; Nightsky Mantle (4718, -0.45 DPS) [world_drop]; Insignia Mantle (4721, -1.04 DPS) [world_drop] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 30.3 healing_power points (3.09 DPS) | yes | Prelacy Cape (7004, +0.00 DPS, sim-verified) [quest]; Repairman's Cape (9605, -0.75 DPS) [quest]; Amy's Blanket (13005, -1.01 DPS) [world_drop] |
| chest | Wisdom's Leather Tunic (252511) | Leatherworking [crafted] | 44.2 healing_power points (4.50 DPS) | yes | Pristine Gown (253961, -0.24 DPS) [crafted]; Beguiler Robes (7728, -0.35 DPS) [dungeon]; Pressed Felt Robe (1997, -0.49 DPS) [world] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 20.6 healing_power points (2.10 DPS) | yes | Nightsky Wristbands (6407, -0.32 DPS) [world_drop]; Dokebi Bracers (14580, -0.46 DPS) [world_drop]; Drakewing Bands (12999, -1.41 DPS, sim-verified) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 42.3 healing_power points (4.30 DPS) | yes | Shilly Mitts (9609, -1.55 DPS, sim-verified) [quest]; Zodiac Gloves (7106, -1.63 DPS) [quest]; Hotshot Pilot's Gloves (9491, -1.63 DPS) [dungeon] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 45.2 healing_power points (4.60 DPS) | yes | Dokebi Cord (14578, -2.23 DPS) [world_drop]; Silver-lined Belt (13011, -2.48 DPS, sim-verified) [world_drop]; Resilient Cord (14406, -2.53 DPS) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 51.6 healing_power points (5.25 DPS) | yes | Pristine Leggings (253987, -0.49 DPS) [crafted]; Stormrider's Leather Kilt (252518, -0.95 DPS) [crafted]; Earthen Leggings (253999, -1.45 DPS, sim-verified) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 41.8 healing_power points (4.26 DPS) | yes | Resilient Boots (14399, -1.88 DPS) [world_drop]; Stonecloth Boots (14408, -1.88 DPS) [world_drop]; Boots of the Enchanter (4325, -3.50 DPS, sim-verified) [crafted] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 26.4 healing_power points (2.69 DPS) | yes | The Queen's Jewel (13094, -0.02 DPS) [world_drop]; Ring of Calm (6790, -0.61 DPS) [quest]; Band of Thorns (5007, -0.91 DPS) [world_drop] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 26.4 healing_power points (2.69 DPS) | yes | The Queen's Jewel (13094, -0.07 DPS, sim-verified) [world_drop]; Ring of Calm (6790, -0.61 DPS) [quest]; Band of Thorns (5007, -0.91 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (89.0 DPS) | yes | Hardened Root Staff (1317, -2.21 DPS) [quest]; Gnarled Ash Staff (791, -2.47 DPS) [world_drop]; Manual Crowd Pummeler (9449, -3.60 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 30.5 healing_power points (3.11 DPS) | yes | Orb of Mistmantle (13031, -1.01 DPS, sim-verified) [world_drop]; Strength of Will (4837, -1.32 DPS) [vendor]; Torch of Holy Flame (7344, -1.62 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Pendant of Myzrael; shoulder: Mantle of Honor; back: Glowing Thresher Cape; chest: Wisdom's Leather Tunic; wrist: Mindthrust Bracers; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 380, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 40 (night-elf, 0000000000000000-00000000000000000000-5050035153112000)

Set DPS (verified): 128.2. Weights run: 8.9s. Verify run: 26.9s. 516 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.000, intellect=1.347 ± 0.016, spirit=2.509 ± 0.011, mp5=4.854 ± 0.034, crit=0.056 ± 0.005 per rating point (14 rating = 1%, 0.782 per %), spell_haste=not significant (0.538 ± 0.314)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 59.5 healing_power points (6.11 DPS) | yes | Holy Shroud (2721, -1.17 DPS) [world_drop]; Whitemane's Chapeau (7720, -1.18 DPS, sim-verified) [dungeon]; Earthen Silk Hood (254015, -1.27 DPS) [crafted] |
| neck | Necklace of Calisea (1714) | Zul'Farrak: Witch Doctor's Chest [dungeon] | 27.0 healing_power points (2.77 DPS) | yes | Glowing Eye of Mordresh (10769, -0.33 DPS) [dungeon]; Explorers' League Commendation (7746, -0.45 DPS) [quest]; Amberglow Talisman (10824, -1.01 DPS, sim-verified) [quest] |
| shoulder | Earthen Silk Shoulders (254033) | Tailoring [crafted] | 38.1 healing_power points (3.91 DPS) | yes | Inquisitor's Shawl (19507, -0.57 DPS) [dungeon]; Crimson Silk Shoulders (7059, -1.00 DPS) [crafted]; Sheepshear Mantle (13115, -1.35 DPS, sim-verified) [world_drop] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 27.1 healing_power points (2.78 DPS) | yes | Prelacy Cape (7004, -0.10 DPS) [quest]; Ceremonial Centaur Blanket (6789, -0.18 DPS) [quest]; Energy Cloak (9397, -0.20 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 78.1 healing_power points (8.02 DPS) | yes | Dreamweave Vest (10021, -3.17 DPS) [crafted]; Civinad Robes (9623, -3.18 DPS) [quest]; Doomsayer's Robe (4746, -4.01 DPS, sim-verified) [quest] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 30.5 healing_power points (3.13 DPS) | yes | Earthen Silk Cuffs (254019, -0.83 DPS, sim-verified) [crafted]; Mindthrust Bracers (1974, -0.94 DPS) [dungeon]; Silkstream Cuffs (16791, -1.41 DPS) [quest] |
| hands | Earthen Silk Gloves (254017) | Tailoring [crafted] | 38.1 healing_power points (3.91 DPS) | yes | Gloves of Old (9395, -0.46 DPS, sim-verified) [world_drop]; Mender's Leather Gloves (252530, -0.51 DPS) [crafted]; Warden's Gloves (14606, -0.52 DPS) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 47.8 healing_power points (4.91 DPS) | yes | Mender's Leather Belt (252523, -1.38 DPS, sim-verified) [crafted]; Windchaser Cinch (14435, -1.49 DPS) [world_drop]; Sutarn's Ring (13105, -1.98 DPS) [world_drop] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 59.2 healing_power points (6.08 DPS) | yes | Stoneweaver Leggings (9407, -0.49 DPS, sim-verified) [dungeon]; Stormcloth Pants (10010, -0.67 DPS) [crafted]; Wisdom's Leather Leggings (252519, -1.15 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) | Leatherworking [crafted] | 51.0 healing_power points (5.24 DPS) | yes | Furen's Boots (13100, -0.93 DPS, sim-verified) [world_drop]; Thoughtcast Boots (10578, -1.07 DPS) [dungeon]; Gilded Slippers (254001, -1.18 DPS) [crafted] |
| finger1 | Welken Ring (5011) | World drop [world_drop] | 24.3 healing_power points (2.50 DPS) | yes | Runemender's Seal (273031, -0.08 DPS) [world]; The Queen's Jewel (13094, -0.16 DPS) [world_drop]; Blush Ember Ring (13093, -0.43 DPS) [world_drop] |
| finger2 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 23.6 healing_power points (2.42 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; The Queen's Jewel (13094, -0.08 DPS) [world_drop]; Blush Ember Ring (13093, -0.36 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (128.2 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -4.56 DPS) [dungeon]; Death Speaker Scepter (2816, -4.73 DPS) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 28.1 healing_power points (2.88 DPS) | yes | Orb of Lorica (11262, -0.38 DPS, sim-verified) [quest]; Beacon of Hope (9393, -0.41 DPS) [dungeon]; Silksand Star (15964, -0.56 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Necklace of Calisea; shoulder: Earthen Silk Shoulders; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Earthen Silk Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; finger1: Welken Ring; finger2: Electrocutioner Lagnut; trinket1: Ankh of Life; main_hand: Hand of Righteousness

No-known-source sample (15 of 516, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 50 (night-elf, 4300000000000000-00000000000000000000-5050035153113200)

Set DPS (verified): 219.5. Weights run: 9.0s. Verify run: 29.0s. 686 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.003, intellect=1.528 ± 0.017, spirit=2.753 ± 0.017, mp5=4.398 ± 0.016, crit=0.072 ± 0.007 per rating point (14 rating = 1%, 1.009 per %), spell_haste=0.983 ± 0.178

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 84.5 healing_power points (8.38 DPS) | yes | Knight-Lieutenant's Restored Leather Helm (220874, -0.48 DPS) [vendor]; Engineer's Guild Headpiece (9534, -0.49 DPS) [quest]; Cassandra's Grace (13102, -1.56 DPS) [world_drop] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 32.4 healing_power points (3.21 DPS) | yes | Amberglow Talisman (10824, -0.48 DPS) [quest]; Necklace of Calisea (1714, -0.53 DPS, sim-verified) [dungeon]; Glowing Eye of Mordresh (10769, -0.71 DPS) [dungeon] |
| shoulder | Living Shoulders (15061) | Leatherworking [crafted] | 66.8 healing_power points (6.62 DPS) | yes | Knight-Lieutenant's Restored Leather Spaulders (220876, -0.20 DPS) [vendor]; Nethergeld Shoulders (254049, -1.19 DPS) [crafted]; Mender's Leather Shoulders (252538, -1.70 DPS, sim-verified) [crafted] |
| back | Wingveil Cloak (10802) | Sunken Temple: Spawn of Hakkar [dungeon] | 35.3 healing_power points (3.50 DPS) | yes | Chainlink Towel (9648, -0.31 DPS) [quest]; Traveler's Cloak (8297, -0.50 DPS) [world_drop]; Arcane Cloak (8286, -1.01 DPS, sim-verified) [world_drop] |
| chest | Ghostweave Vest (14141) | Tailoring [crafted] | 90.2 healing_power points (8.95 DPS) | yes | Feathered Breastplate (8349, -0.84 DPS, sim-verified) [crafted]; Stormcloth Vest (10020, -1.01 DPS) [crafted]; Knight's Restored Leather Jerkin (220878, -1.45 DPS) [vendor] |
| wrist | Mender's Leather Bracers (252543) | Leatherworking [crafted] | 46.2 healing_power points (4.58 DPS) | yes | Nethergeld Cuffs (254061, -0.55 DPS, sim-verified) [crafted]; Enchanted Kodo Bracers (13119, -1.25 DPS) [world_drop]; Venomshroud Armguards (14439, -1.79 DPS) [world_drop] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | 69.5 healing_power points (6.89 DPS) | yes | Gilded Gloves (254095, -0.55 DPS, sim-verified) [crafted]; Earthenweave Gloves (254075, -1.31 DPS) [crafted]; Keeper's Gloves (14666, -2.50 DPS) [world_drop] |
| waist | Earthenweave Cord (254077) | Tailoring [crafted] | 56.3 healing_power points (5.58 DPS) | yes | Mender's Leather Waistguard (252477, -0.79 DPS) [crafted]; Mender's Leather Belt (252523, -1.15 DPS) [crafted]; Gilded Cord (254037, -1.25 DPS, sim-verified) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 85.2 healing_power points (8.45 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -0.22 DPS) [dungeon]; Dalewind Trousers (13008, -0.27 DPS) [world_drop]; Jinxed Hoodoo Kilt (9474, -0.98 DPS, sim-verified) [dungeon] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 67.3 healing_power points (6.67 DPS) | yes | Sergeant Major's Restored Leather Boots (220884, -0.44 DPS) [vendor]; Earthenweave Boots (254093, -1.09 DPS) [crafted]; Furen's Boots (13100, -1.24 DPS) [world_drop] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 42.2 healing_power points (4.18 DPS) | yes | Coldwater Ring (4550, -1.27 DPS) [quest]; Seedtime Hoop (9655, -1.45 DPS) [quest]; Welken Ring (5011, -1.52 DPS) [world_drop] |
| finger2 | Choking Band (11868) | Ogre Head On A Stick = Party [quest] | 35.8 healing_power points (3.55 DPS) | yes | Coldwater Ring (4550, -0.57 DPS, sim-verified) [quest]; Seedtime Hoop (9655, -0.82 DPS) [quest]; Welken Ring (5011, -0.88 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (219.5 DPS) | yes | Thunderbrew's Boot Flask (744, -1.09 DPS) [quest]; Uther's Strength (11302, -1.49 DPS) [world_drop]; Alchemist's Stone (13503, -2.18 DPS) [crafted] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-verified (219.5 DPS) | yes | Uther's Strength (11302, -0.48 DPS, sim-verified) [world_drop]; Thunderbrew's Boot Flask (744, -0.55 DPS) [quest]; Alchemist's Stone (13503, -1.64 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (219.5 DPS) | yes | Charstone Dirk (17710, -7.68 DPS) [dungeon]; Hand of Righteousness (7721, -10.22 DPS) [dungeon]; Hammer of the Northern Wind (810, -33.39 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Living Shoulders; back: Wingveil Cloak; chest: Ghostweave Vest; wrist: Mender's Leather Bracers; hands: Mender's Leather Gauntlets; waist: Earthenweave Cord; legs: Senior Designer's Pantaloons; feet: Sandals of the Insurgent; finger1: Eye of Adaegus; finger2: Choking Band; trinket2: Evonice's Landin' Pilla; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 686, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60 (night-elf, 4532200100000000-00000000000000000000-5050035153113200)

Set DPS (verified): 325.5. Weights run: 9.3s. Verify run: 81.6s. 1564 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.022, intellect=2.234 ± 0.044, spirit=4.276 ± 0.034, mp5=6.865 ± 0.033, crit=0.098 ± 0.010 per rating point (14 rating = 1%, 1.366 per %), spell_haste=not significant (-0.039 ± 0.480)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 173.1 healing_power points (17.46 DPS) | yes | Lieutenant Commander's Dragonhide Headdress (227199, -2.67 DPS) [vendor]; Devout Crown (16693, -3.56 DPS) [dungeon]; Wildheart Cowl (16720, -6.76 DPS, sim-verified) [dungeon] |
| neck | Lady Alizabeth's Pendant (13002) | World drop [world_drop] | 64.1 healing_power points (6.47 DPS) | yes | Horizon Choker (13085, -1.72 DPS, sim-verified) [world_drop]; Necklace of Calisea (1714, -1.87 DPS) [dungeon]; Amberglow Talisman (10824, -2.16 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | 96.4 healing_power points (9.72 DPS) | yes | Lieutenant Commander's Dragonhide Pauldrons (227201, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Pauldrons (231705, +0.00 DPS) [vendor]; Wildheart Spaulders (16718, -5.85 DPS, sim-verified) [dungeon] |
| back | Cloak of Earth and Sky (279270) | Tailoring [crafted] | 59.9 healing_power points (6.04 DPS) | yes | Ethereal Mist Cape (11873, -0.19 DPS) [quest]; Indomitable Cloak (14683, -0.43 DPS) [world_drop]; Sergeant Major's Cape (16337, -0.79 DPS) [pvp] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | 151.9 healing_power points (15.33 DPS) | yes | Wildheart Vest (16706, -1.08 DPS) [dungeon]; Ghostweave Vest (14141, -1.16 DPS) [crafted]; Ironfeather Breastplate (15066, -2.30 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | 71.6 healing_power points (7.23 DPS) | yes | Marshal's Dragonhide Bracers (16445, +0.00 DPS) [pvp]; Mender's Leather Bracers (252543, -1.14 DPS) [crafted]; Earthenweave Cuffs (254131, -6.12 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (325.5 DPS) | yes | Swarmtender's Gloves (275607, -0.86 DPS) [crafted]; Marshal's Dragonhide Gauntlets (231706, -1.49 DPS) [pvp]; Wildheart Gloves (16717, -5.06 DPS, sim-verified) [dungeon] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (325.5 DPS) | yes | Marshal's Dragonhide Waistguard (16447, +0.00 DPS) [pvp]; Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Elderwild Waistcord (279252, -2.96 DPS, sim-verified) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 155.9 healing_power points (15.72 DPS) | yes | Knight-Captain's Dragonhide Legguards (227200, -1.01 DPS) [vendor]; Living Leggings (15060, -2.32 DPS) [crafted]; Elderwild Pants (279254, -6.01 DPS, sim-verified) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (325.5 DPS) | yes | Wildheart Boots (16715, -0.23 DPS) [dungeon]; Sandals of the Insurgent (13111, -0.27 DPS) [world_drop]; Mooncloth Boots (15802, -5.07 DPS, sim-verified) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 69.4 healing_power points (7.00 DPS) | yes | Choking Band (11868, -1.39 DPS) [quest]; Coldwater Ring (4550, -2.44 DPS) [quest]; Seedtime Hoop (9655, -2.68 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 64.7 healing_power points (6.53 DPS) | yes | Choking Band (11868, -0.69 DPS, sim-verified) [quest]; Coldwater Ring (4550, -1.97 DPS) [quest]; Seedtime Hoop (9655, -2.21 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (325.5 DPS) | yes | Shard of the Splithooves (10659, -1.02 DPS) [quest]; Thunderbrew's Boot Flask (744, -1.73 DPS) [quest]; Uther's Strength (11302, -3.36 DPS) [world_drop] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-verified (325.5 DPS) | yes | Shard of the Splithooves (10659, -0.16 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.86 DPS, sim-verified) [quest]; Uther's Strength (11302, -2.50 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (325.5 DPS) | yes | Charstone Dirk (17710, -9.98 DPS) [dungeon]; Staff of Hale Magefire (13000, -10.43 DPS) [world_drop]; Hand of Edward the Odd (2243, -13.11 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Lady Alizabeth's Pendant; shoulder: Devout Mantle; back: Cloak of Earth and Sky; chest: Mooncloth Vest; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60, raid preset (night-elf, 4532200100000000-00000000000000000000-5050035153113200)

Set DPS (verified): 614.3. Weights run: 6.7s. Verify run: 54.1s. 1564 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.204, intellect=2.451 ± 0.080, spirit=2.281 ± 0.075, mp5=3.336 ± 0.099, crit=0.327 ± 0.030 per rating point (14 rating = 1%, 4.573 per %), spell_haste=not significant (-0.743 ± 1.216)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 149.0 healing_power points (24.64 DPS) | yes | Lieutenant Commander's Dragonhide Headdress (227199, -4.67 DPS) [vendor]; Wildheart Cowl (16720, -6.52 DPS) [dungeon]; Devout Crown (16693, -12.27 DPS, sim-verified) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 43.4 healing_power points (7.18 DPS) | yes | Lady Alizabeth's Pendant (13002, +0.00 DPS, sim-verified) [world_drop]; Necklace of Calisea (1714, -1.70 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -1.79 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | 83.0 healing_power points (13.72 DPS) | yes | Lieutenant Commander's Dragonhide Pauldrons (227201, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Pauldrons (231705, +0.00 DPS) [vendor]; Wildheart Spaulders (16718, -3.94 DPS, sim-verified) [dungeon] |
| back | Sergeant Major's Cape (16337) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 37.9 healing_power points (6.26 DPS) | yes | Ethereal Mist Cape (11873, -0.52 DPS) [quest]; Cloak of Earth and Sky (279270, -0.98 DPS) [crafted]; Imperial Red Cloak (8248, -1.75 DPS, sim-verified) [world_drop] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (614.4 DPS) | yes | Field Marshal's Dragonhide Tunic (231702, -0.51 DPS) [vendor]; Knight-Captain's Dragonhide Tunic (227196, -0.72 DPS) [pvp]; Mooncloth Vest (14138, -2.17 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | 53.4 healing_power points (8.83 DPS) | yes | Marshal's Dragonhide Bracers (16445, +0.00 DPS) [pvp]; Nethergeld Cuffs (254061, -0.80 DPS) [crafted]; Mender's Leather Bracers (252543, -4.71 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (614.4 DPS) | yes | Marshal's Dragonhide Gauntlets (231706, +0.00 DPS) [pvp]; Mender's Leather Gauntlets (252551, -0.78 DPS) [crafted]; Wildheart Gloves (16717, -5.04 DPS, sim-verified) [dungeon] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (614.4 DPS) | yes | Marshal's Dragonhide Waistguard (16447, +0.00 DPS) [pvp]; Wildheart Belt (16716, -0.84 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -1.46 DPS, sim-verified) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 113.2 healing_power points (18.72 DPS) | yes | Knight-Captain's Dragonhide Legguards (227200, +0.00 DPS) [vendor]; Marshal's Dragonhide Legguards (231703, -2.16 DPS) [pvp]; Elderwild Pants (279254, -4.39 DPS, sim-verified) [crafted] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 88.0 healing_power points (14.54 DPS) | yes | Devout Sandals (16691, -2.39 DPS, sim-verified) [dungeon]; Wildheart Boots (16715, -2.67 DPS) [dungeon]; Marshal's Dragonhide Greaves (231704, -2.67 DPS) [vendor] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 49.6 healing_power points (8.20 DPS) | yes | Maiden's Circle (13001, -3.10 DPS) [world_drop]; Choking Band (11868, -3.30 DPS) [quest]; Glowing Crystal Ring (18402, -3.34 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 42.1 healing_power points (6.96 DPS) | yes | Choking Band (11868, -2.05 DPS) [quest]; Glowing Crystal Ring (18402, -2.09 DPS) [quest]; Maiden's Circle (13001, -2.27 DPS, sim-verified) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (614.4 DPS) | yes | Shard of the Splithooves (10659, -1.22 DPS) [quest]; Thunderbrew's Boot Flask (744, -1.51 DPS) [quest]; Uther's Strength (11302, -1.55 DPS) [world_drop] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-verified (614.4 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop]; Shard of the Splithooves (10659, -0.46 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.75 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (614.4 DPS) | yes | Charstone Dirk (17710, -15.20 DPS) [dungeon]; Hand of Edward the Odd (2243, -18.09 DPS, sim-verified) [world_drop]; Hand of Righteousness (7721, -21.38 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; shoulder: Devout Mantle; back: Sergeant Major's Cape; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of the Hierophant; finger2: Eye of Adaegus

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

## Horde

### Band 20 (tauren, 0000000000000000-00000000000000000000-5050010000000000)

Set DPS (verified): 46.3. Weights run: 5.7s. Verify run: 16.8s. 192 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.000, intellect=1.289 ± 0.004, spirit=1.552 ± 0.008, mp5=3.307 ± 0.015, crit=0.044 ± 0.003 per rating point (14 rating = 1%, 0.609 per %), spell_haste=-5.588 ± 0.106

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.85 DPS) | yes | Wisdom's Leather Hood (252507, -0.14 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.23 DPS) [crafted]; Stormrider's Leather Hood (252506, -0.42 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 4.7 healing_power points (0.48 DPS) | yes | Scholarly Pendant (277203, -3.03 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.6 healing_power points (1.20 DPS) | yes | Forest Leather Mantle (4709, -0.56 DPS) [world_drop]; Prospector's Pads (14566, -0.56 DPS) [world_drop]; Slime-encrusted Pads (6461, -0.84 DPS, sim-verified) [dungeon] |
| back | Spirit Cloak (4792) (or Antiquated Cloak (4799), Sylvan Cloak (4793), Heavy Woolen Cloak (4311)) | Jennabink Powerseam [vendor] | 6.2 healing_power points (0.64 DPS) | yes | Heavy Woolen Cloak (4311, +0.00 DPS) [crafted]; Sylvan Cloak (4793, +0.00 DPS) [vendor]; Antiquated Cloak (4799, +0.00 DPS) [vendor] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 26.1 healing_power points (2.69 DPS) | yes | Filigreed Pristine Gown (253901, -0.21 DPS) [crafted]; Corsair's Overshirt (5202, -0.28 DPS) [dungeon]; Robe of the Moccasin (6465, -1.93 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 16.4 healing_power points (1.69 DPS) | yes | Crystalline Cuffs (14148, -0.94 DPS) [dungeon]; Drakewing Bands (12999, -1.02 DPS, sim-verified) [world_drop]; Owl Bracers (4796, -1.02 DPS) [vendor] |
| hands | Blight Gloves (279877) | The New Plague [quest] | 16.8 healing_power points (1.73 DPS) | yes | Pristine Gloves (253913, -0.20 DPS) [crafted]; Magefist Gloves (12977, -0.43 DPS) [world_drop]; Wisdom's Leather Gloves (252499, -0.80 DPS, sim-verified) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.2 healing_power points (1.66 DPS) | yes | Wisdom's Leather Belt (252433, -0.31 DPS) [crafted]; Novice Ardent's Sash (253887, -0.44 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.60 DPS) [world_drop] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 33.9 healing_power points (3.50 DPS) | yes | Filigreed Pristine Leggings (253937, -0.27 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -1.61 DPS) [world_drop]; Filigreed Silky Leggings (253939, -2.06 DPS) [crafted] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 17.4 healing_power points (1.80 DPS) | yes | Pristine Boots (253889, -0.47 DPS) [crafted]; Smoldering Boots (3076, -0.84 DPS) [world]; Black Whelp Slippers (252424, -2.12 DPS, sim-verified) [crafted] |
| finger1 | Band of Purification (12996) | World drop [world_drop] | 9.3 healing_power points (0.96 DPS) | yes | Lavishly Jeweled Ring (1156, -0.16 DPS) [dungeon]; Loop of Sacrifice (281673, -0.30 DPS) [quest]; Volcanic Rock Ring (12053, -0.56 DPS) [world_drop] |
| finger2 | Deep Fathom Ring (6463) | Wailing Caverns: Mutanus the Devourer [dungeon] | 7.8 healing_power points (0.80 DPS) | yes | Loop of Sacrifice (281673, -0.14 DPS) [quest]; Volcanic Rock Ring (12053, -0.40 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -2.60 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Orgrimmar (15444) | Hidden Enemies [quest] | 19.5 healing_power points (2.01 DPS) | yes | Staff of the Blessed Seer (2271, -0.68 DPS, sim-verified) [dungeon]; Twisted Chanter's Staff (890, -0.68 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.68 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Spirit Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Blight Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Band of Purification; finger2: Deep Fathom Ring; main_hand: Staff of Orgrimmar

No-known-source sample (15 of 192, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion; 21566 Rune of Perfection

### Band 30 (tauren, 0000000000000000-00000000000000000000-5050035110010000)

Set DPS (verified): 88.3. Weights run: 8.0s. Verify run: 27.8s. 380 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.454 ± 0.014, spirit=2.917 ± 0.018, mp5=4.441 ± 0.027, crit=0.052 ± 0.005 per rating point (14 rating = 1%, 0.735 per %), spell_haste=not significant (-1.023 ± 0.259)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 50.5 healing_power points (5.14 DPS) | yes | Enduring Cap (3020, -0.69 DPS) [world_drop]; Nightsky Cowl (4039, -1.88 DPS) [world_drop]; Dokebi Hat (14584, -2.03 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | sim-verified (88.3 DPS) | yes | Pendant of Myzrael (4614, +0.00 DPS) [dungeon]; Glowing Green Talisman (5002, +0.00 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.30 DPS) [world_drop] |
| shoulder | Ghostly Mantle (3324) | Deathstalkers in Shadowfang [quest] | 35.3 healing_power points (3.59 DPS) | yes | Nightsky Mantle (4718, -0.92 DPS) [world_drop]; Desert Shoulders (15457, -1.36 DPS) [quest]; Insignia Mantle (4721, -1.51 DPS) [world_drop] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 30.3 healing_power points (3.09 DPS) | yes | Amy's Blanket (13005, -1.01 DPS) [world_drop]; Nightsky Cloak (4719, -1.31 DPS) [world_drop]; Insignia Cloak (4722, -1.31 DPS) [world_drop] |
| chest | Wisdom's Leather Tunic (252511) | Leatherworking [crafted] | 44.2 healing_power points (4.50 DPS) | yes | Pristine Gown (253961, -0.24 DPS) [crafted]; Beguiler Robes (7728, -0.35 DPS) [dungeon]; Pressed Felt Robe (1997, -0.49 DPS) [world] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 20.6 healing_power points (2.10 DPS) | yes | Drakewing Bands (12999, -0.31 DPS) [world_drop]; Nightsky Wristbands (6407, -0.32 DPS) [world_drop]; Dokebi Bracers (14580, -0.46 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 42.3 healing_power points (4.30 DPS) | yes | Tattered Mittens (270030, -1.50 DPS) [quest]; Hotshot Pilot's Gloves (9491, -1.63 DPS) [dungeon]; Blight Gloves (279877, -1.78 DPS) [quest] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 45.2 healing_power points (4.60 DPS) | yes | Silver-lined Belt (13011, -2.08 DPS) [world_drop]; Dokebi Cord (14578, -2.23 DPS) [world_drop]; Lilac Sash (6780, -2.38 DPS) [quest] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 51.6 healing_power points (5.25 DPS) | yes | Earthen Leggings (253999, -0.46 DPS) [crafted]; Pristine Leggings (253987, -0.49 DPS) [crafted]; Stormrider's Leather Kilt (252518, -0.95 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 41.8 healing_power points (4.26 DPS) | yes | Boots of the Enchanter (4325, -1.88 DPS) [crafted]; Resilient Boots (14399, -1.88 DPS) [world_drop]; Stonecloth Boots (14408, -1.88 DPS) [world_drop] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 26.4 healing_power points (2.69 DPS) | yes | The Queen's Jewel (13094, -0.02 DPS) [world_drop]; Solid Gold Band (279845, -0.49 DPS) [quest]; Band of Thorns (5007, -0.91 DPS) [world_drop] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 26.4 healing_power points (2.69 DPS) | yes | The Queen's Jewel (13094, -0.02 DPS) [world_drop]; Solid Gold Band (279845, -0.49 DPS) [quest]; Band of Thorns (5007, -0.91 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gnarled Ash Staff (791, -2.47 DPS) [world_drop]; Staff of the Friar (3415, -2.92 DPS) [dungeon]; Spine of the Devourer (273040, -3.26 DPS) [world] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 30.5 healing_power points (3.11 DPS) | yes | Orb of Mistmantle (13031, -1.00 DPS) [world_drop]; Strength of Will (4837, -1.32 DPS) [vendor]; Silver-thread Rod (15928, -1.62 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Ghostly Mantle; back: Glowing Thresher Cape; chest: Wisdom's Leather Tunic; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 380, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (tauren, 0000000000000000-00000000000000000000-5050035153112000)

Set DPS (verified): 127.8. Weights run: 8.9s. Verify run: 26.9s. 514 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.000, intellect=1.347 ± 0.016, spirit=2.509 ± 0.011, mp5=4.854 ± 0.034, crit=0.056 ± 0.005 per rating point (14 rating = 1%, 0.782 per %), spell_haste=not significant (0.538 ± 0.314)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 59.5 healing_power points (6.11 DPS) | yes | Holy Shroud (2721, -1.17 DPS) [world_drop]; Earthen Silk Hood (254015, -1.27 DPS) [crafted]; Whitemane's Chapeau (7720, -1.83 DPS, sim-verified) [dungeon] |
| neck | Necklace of Calisea (1714) | Zul'Farrak: Witch Doctor's Chest [dungeon] | 27.0 healing_power points (2.77 DPS) | yes | Glowing Eye of Mordresh (10769, -0.33 DPS) [dungeon]; Triune Amulet (7722, -0.97 DPS) [dungeon]; Amberglow Talisman (10824, -1.65 DPS, sim-verified) [quest] |
| shoulder | Earthen Silk Shoulders (254033) | Tailoring [crafted] | 38.1 healing_power points (3.91 DPS) | yes | Inquisitor's Shawl (19507, -0.57 DPS) [dungeon]; Ghostly Mantle (3324, -0.67 DPS) [quest]; Sheepshear Mantle (13115, -0.82 DPS, sim-verified) [world_drop] |
| back | Cloak of Blight (6832) | Nothing But The Truth [quest] | 32.6 healing_power points (3.35 DPS) | yes | Glowing Thresher Cape (6901, -0.67 DPS, sim-verified) [dungeon]; Ceremonial Centaur Blanket (6789, -0.75 DPS) [quest]; Stargazer Cloak (9660, -0.77 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 78.1 healing_power points (8.02 DPS) | yes | Dreamweave Vest (10021, -3.17 DPS) [crafted]; Civinad Robes (9623, -3.18 DPS) [quest]; Doomsayer's Robe (4746, -4.43 DPS, sim-verified) [quest] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 30.5 healing_power points (3.13 DPS) | yes | Mindthrust Bracers (1974, -0.94 DPS) [dungeon]; Silkstream Cuffs (16791, -1.41 DPS) [quest]; Earthen Silk Cuffs (254019, -1.65 DPS, sim-verified) [crafted] |
| hands | Earthen Silk Gloves (254017) | Tailoring [crafted] | 38.1 healing_power points (3.91 DPS) | yes | Mender's Leather Gloves (252530, -0.51 DPS) [crafted]; Warden's Gloves (14606, -0.52 DPS) [world_drop]; Gloves of Old (9395, -0.72 DPS, sim-verified) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 47.8 healing_power points (4.91 DPS) | yes | Windchaser Cinch (14435, -1.49 DPS) [world_drop]; Mender's Leather Belt (252523, -1.72 DPS, sim-verified) [crafted]; Sutarn's Ring (13105, -1.98 DPS) [world_drop] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 59.2 healing_power points (6.08 DPS) | yes | Stormcloth Pants (10010, -0.67 DPS) [crafted]; Stoneweaver Leggings (9407, -0.79 DPS, sim-verified) [dungeon]; Wisdom's Leather Leggings (252519, -1.15 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) | Leatherworking [crafted] | 51.0 healing_power points (5.24 DPS) | yes | Thoughtcast Boots (10578, -1.07 DPS) [dungeon]; Gilded Slippers (254001, -1.18 DPS) [crafted]; Furen's Boots (13100, -1.20 DPS, sim-verified) [world_drop] |
| finger1 | Welken Ring (5011) | World drop [world_drop] | 24.3 healing_power points (2.50 DPS) | yes | Runemender's Seal (273031, -0.08 DPS) [world]; The Queen's Jewel (13094, -0.16 DPS) [world_drop]; Blush Ember Ring (13093, -0.43 DPS) [world_drop] |
| finger2 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 23.6 healing_power points (2.42 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; The Queen's Jewel (13094, -0.08 DPS) [world_drop]; Blush Ember Ring (13093, -0.36 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (127.8 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -4.56 DPS) [dungeon]; Death Speaker Scepter (2816, -4.73 DPS) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 28.1 healing_power points (2.88 DPS) | yes | Beacon of Hope (9393, -0.42 DPS, sim-verified) [dungeon]; Silksand Star (15964, -0.56 DPS) [world_drop]; Aurora Sphere (7610, -0.88 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Necklace of Calisea; shoulder: Earthen Silk Shoulders; back: Cloak of Blight; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Earthen Silk Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; finger1: Welken Ring; finger2: Electrocutioner Lagnut; trinket1: Ankh of Life; main_hand: Hand of Righteousness

No-known-source sample (15 of 514, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 50 (tauren, 4300000000000000-00000000000000000000-5050035153113200)

Set DPS (verified): 217.4. Weights run: 9.0s. Verify run: 29.1s. 684 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.003, intellect=1.528 ± 0.017, spirit=2.753 ± 0.017, mp5=4.398 ± 0.016, crit=0.072 ± 0.007 per rating point (14 rating = 1%, 1.009 per %), spell_haste=0.983 ± 0.178

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 84.5 healing_power points (8.38 DPS) | yes | Blood Guard's Restored Leather Helm (220875, -0.48 DPS) [vendor]; Engineer's Guild Headpiece (9534, -0.49 DPS) [quest]; Cassandra's Grace (13102, -1.56 DPS) [world_drop] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 32.4 healing_power points (3.21 DPS) | yes | Necklace of Calisea (1714, +0.00 DPS, sim-verified) [dungeon]; Choker of the High Shaman (4112, -0.27 DPS) [quest]; Amberglow Talisman (10824, -0.48 DPS) [quest] |
| shoulder | Living Shoulders (15061) | Leatherworking [crafted] | 66.8 healing_power points (6.62 DPS) | yes | Blood Guard's Restored Leather Spaulders (220877, -0.20 DPS) [vendor]; Mender's Leather Shoulders (252538, -1.13 DPS, sim-verified) [crafted]; Nethergeld Shoulders (254049, -1.19 DPS) [crafted] |
| back | Cloak of Blight (6832) | Nothing But The Truth [quest] | 35.8 healing_power points (3.55 DPS) | yes | Wingveil Cloak (10802, +0.00 DPS, sim-verified) [dungeon]; Arcane Cloak (8286, -0.27 DPS) [world_drop]; Chainlink Towel (9648, -0.36 DPS) [quest] |
| chest | Ghostweave Vest (14141) | Tailoring [crafted] | 90.2 healing_power points (8.95 DPS) | yes | Feathered Breastplate (8349, -0.94 DPS, sim-verified) [crafted]; Stormcloth Vest (10020, -1.01 DPS) [crafted]; Stone Guard's Restored Leather Jerkin (220879, -1.45 DPS) [vendor] |
| wrist | Mender's Leather Bracers (252543) | Leatherworking [crafted] | 46.2 healing_power points (4.58 DPS) | yes | Nethergeld Cuffs (254061, -0.17 DPS) [crafted]; Enchanted Kodo Bracers (13119, -1.25 DPS) [world_drop]; Venomshroud Armguards (14439, -1.79 DPS) [world_drop] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | 69.5 healing_power points (6.89 DPS) | yes | Gilded Gloves (254095, -0.37 DPS, sim-verified) [crafted]; Earthenweave Gloves (254075, -1.31 DPS) [crafted]; Keeper's Gloves (14666, -2.50 DPS) [world_drop] |
| waist | Earthenweave Cord (254077) | Tailoring [crafted] | 56.3 healing_power points (5.58 DPS) | yes | Mender's Leather Waistguard (252477, -0.79 DPS) [crafted]; Gilded Cord (254037, -0.90 DPS, sim-verified) [crafted]; Mender's Leather Belt (252523, -1.15 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 85.2 healing_power points (8.45 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -0.22 DPS) [dungeon]; Dalewind Trousers (13008, -0.27 DPS) [world_drop]; Jinxed Hoodoo Kilt (9474, -0.80 DPS, sim-verified) [dungeon] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 67.3 healing_power points (6.67 DPS) | yes | First Sergeant's Restored Leather Boots (220885, +0.00 DPS, sim-verified) [vendor]; Earthenweave Boots (254093, -1.09 DPS) [crafted]; Furen's Boots (13100, -1.24 DPS) [world_drop] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 42.2 healing_power points (4.18 DPS) | yes | Welken Ring (5011, -1.52 DPS) [world_drop]; Electrocutioner Lagnut (9447, -1.68 DPS) [dungeon]; Runemender's Seal (273031, -1.68 DPS) [world] |
| finger2 | Coldwater Ring (4550) | Death From Below [quest] | 29.4 healing_power points (2.91 DPS) | yes | Welken Ring (5011, -0.15 DPS, sim-verified) [world_drop]; Electrocutioner Lagnut (9447, -0.41 DPS) [dungeon]; Runemender's Seal (273031, -0.41 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (217.4 DPS) | yes | Uther's Strength (11302, -1.49 DPS) [world_drop]; Alchemist's Stone (13503, -2.18 DPS) [crafted] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-verified (217.4 DPS) | yes | Uther's Strength (11302, -0.58 DPS, sim-verified) [world_drop]; Alchemist's Stone (13503, -1.64 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (217.4 DPS) | yes | Charstone Dirk (17710, -7.68 DPS) [dungeon]; Hand of Righteousness (7721, -10.22 DPS) [dungeon]; Hammer of the Northern Wind (810, -33.46 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Living Shoulders; chest: Ghostweave Vest; wrist: Mender's Leather Bracers; hands: Mender's Leather Gauntlets; waist: Earthenweave Cord; legs: Senior Designer's Pantaloons; feet: Sandals of the Insurgent; finger1: Eye of Adaegus; finger2: Coldwater Ring; trinket2: Evonice's Landin' Pilla; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 684, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60 (tauren, 4532200100000000-00000000000000000000-5050035153113200)

Set DPS (verified): 321.8. Weights run: 9.3s. Verify run: 82.0s. 1564 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.022, intellect=2.234 ± 0.044, spirit=4.276 ± 0.034, mp5=6.865 ± 0.033, crit=0.098 ± 0.010 per rating point (14 rating = 1%, 1.366 per %), spell_haste=not significant (-0.039 ± 0.480)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 173.1 healing_power points (17.46 DPS) | yes | Champion's Dragonhide Headdress (227205, -2.67 DPS) [vendor]; Devout Crown (16693, -3.56 DPS) [dungeon]; Wildheart Cowl (16720, -6.93 DPS, sim-verified) [dungeon] |
| neck | Lady Alizabeth's Pendant (13002) | World drop [world_drop] | 64.1 healing_power points (6.47 DPS) | yes | Necklace of Calisea (1714, -1.87 DPS) [dungeon]; Choker of the High Shaman (4112, -1.89 DPS) [quest]; Horizon Choker (13085, -2.26 DPS, sim-verified) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | 96.4 healing_power points (9.72 DPS) | yes | Champion's Dragonhide Pauldrons (227207, +0.00 DPS) [vendor]; Warlord's Dragonhide Pauldrons (231672, +0.00 DPS) [vendor]; Wildheart Spaulders (16718, -4.55 DPS, sim-verified) [dungeon] |
| back | Cloak of Earth and Sky (279270) | Tailoring [crafted] | 59.9 healing_power points (6.04 DPS) | yes | Indomitable Cloak (14683, -0.43 DPS) [world_drop]; Cloak of Blight (6832, -0.63 DPS, sim-verified) [quest]; Wingveil Cloak (10802, -0.79 DPS) [dungeon] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | 151.9 healing_power points (15.33 DPS) | yes | Wildheart Vest (16706, -1.08 DPS) [dungeon]; Ghostweave Vest (14141, -1.16 DPS) [crafted]; Ironfeather Breastplate (15066, -2.27 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | 71.6 healing_power points (7.23 DPS) | yes | General's Dragonhide Bracers (16553, +0.00 DPS) [pvp]; Mender's Leather Bracers (252543, -1.14 DPS) [crafted]; Earthenweave Cuffs (254131, -5.33 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (321.7 DPS) | yes | Swarmtender's Gloves (275607, -0.86 DPS) [crafted]; General's Dragonhide Gauntlets (231676, -1.49 DPS) [vendor]; Wildheart Gloves (16717, -5.31 DPS, sim-verified) [dungeon] |
| waist | Elderwild Waistcord (279252) | Leatherworking [crafted] | 107.0 healing_power points (10.79 DPS) | yes | Wisdom of the Timbermaw (19047, -0.74 DPS, sim-verified) [crafted]; General's Dragonhide Belt (16556, -0.84 DPS) [pvp]; Devout Belt (16696, -1.29 DPS) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 155.9 healing_power points (15.72 DPS) | yes | Legionnaire's Dragonhide Legguards (227206, -1.01 DPS) [vendor]; Living Leggings (15060, -2.32 DPS) [crafted]; Elderwild Pants (279254, -5.11 DPS, sim-verified) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (321.7 DPS) | yes | Wildheart Boots (16715, -0.23 DPS) [dungeon]; Sandals of the Insurgent (13111, -0.27 DPS) [world_drop]; Mooncloth Boots (15802, -3.76 DPS, sim-verified) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 69.4 healing_power points (7.00 DPS) | yes | Coldwater Ring (4550, -2.44 DPS) [quest]; Maiden's Circle (13001, -2.83 DPS) [world_drop]; Welken Ring (5011, -2.85 DPS) [world_drop] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 64.7 healing_power points (6.53 DPS) | yes | Coldwater Ring (4550, -2.10 DPS, sim-verified) [quest]; Maiden's Circle (13001, -2.36 DPS) [world_drop]; Welken Ring (5011, -2.38 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (321.7 DPS) | yes | Evonice's Landin' Pilla (18951, -0.99 DPS, sim-verified) [quest]; Uther's Strength (11302, -3.36 DPS) [world_drop]; Alchemist's Stone (13503, -3.45 DPS) [crafted] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (321.7 DPS) | yes | Evonice's Landin' Pilla (18951, -0.27 DPS, sim-verified) [quest]; Uther's Strength (11302, -2.34 DPS) [world_drop]; Alchemist's Stone (13503, -2.43 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (321.7 DPS) | yes | Charstone Dirk (17710, -9.98 DPS) [dungeon]; Staff of Hale Magefire (13000, -10.43 DPS) [world_drop]; Hand of Edward the Odd (2243, -12.89 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Lady Alizabeth's Pendant; shoulder: Devout Mantle; back: Cloak of Earth and Sky; chest: Mooncloth Vest; wrist: Devout Bracers; hands: Devout Gloves; waist: Elderwild Waistcord; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket2: Shard of the Splithooves

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60, raid preset (tauren, 4532200100000000-00000000000000000000-5050035153113200)

Set DPS (verified): 611.7. Weights run: 6.7s. Verify run: 54.5s. 1564 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.204, intellect=2.451 ± 0.080, spirit=2.281 ± 0.075, mp5=3.336 ± 0.099, crit=0.327 ± 0.030 per rating point (14 rating = 1%, 4.573 per %), spell_haste=not significant (-0.743 ± 1.216)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 149.0 healing_power points (24.64 DPS) | yes | Champion's Dragonhide Headdress (227205, -4.67 DPS) [vendor]; Wildheart Cowl (16720, -6.52 DPS) [dungeon]; Devout Crown (16693, -11.42 DPS, sim-verified) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 43.4 healing_power points (7.18 DPS) | yes | Lady Alizabeth's Pendant (13002, +0.00 DPS, sim-verified) [world_drop]; Necklace of Calisea (1714, -1.70 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -1.79 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | 83.0 healing_power points (13.72 DPS) | yes | Champion's Dragonhide Pauldrons (227207, +0.00 DPS) [vendor]; Warlord's Dragonhide Pauldrons (231672, +0.00 DPS) [vendor]; Wildheart Spaulders (16718, -4.20 DPS, sim-verified) [dungeon] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 36.1 healing_power points (5.96 DPS) | yes | Cloak of Earth and Sky (279270, +0.00 DPS, sim-verified) [crafted]; Highborne Cloak (14450, -0.81 DPS) [world_drop]; Chainlink Towel (9648, -0.98 DPS) [quest] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (611.7 DPS) | yes | Warlord's Dragonhide Tunic (231674, -0.51 DPS) [vendor]; Legionnaire's Dragonhide Tunic (227202, -0.72 DPS) [vendor]; Mooncloth Vest (14138, -2.83 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | 53.4 healing_power points (8.83 DPS) | yes | General's Dragonhide Bracers (16553, +0.00 DPS) [pvp]; Nethergeld Cuffs (254061, -0.80 DPS) [crafted]; Mender's Leather Bracers (252543, -5.07 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (611.7 DPS) | yes | General's Dragonhide Gauntlets (231676, +0.00 DPS) [vendor]; Mender's Leather Gauntlets (252551, -0.78 DPS) [crafted]; Wildheart Gloves (16717, -4.36 DPS, sim-verified) [dungeon] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (611.7 DPS) | yes | General's Dragonhide Belt (16556, +0.00 DPS) [pvp]; Wildheart Belt (16716, -0.84 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -1.72 DPS, sim-verified) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 113.2 healing_power points (18.72 DPS) | yes | Legionnaire's Dragonhide Legguards (227206, +0.00 DPS) [vendor]; General's Dragonhide Legguards (231673, -2.16 DPS) [vendor]; Elderwild Pants (279254, -4.98 DPS, sim-verified) [crafted] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 88.0 healing_power points (14.54 DPS) | yes | Devout Sandals (16691, -2.29 DPS, sim-verified) [dungeon]; Wildheart Boots (16715, -2.67 DPS) [dungeon]; General's Dragonhide Greaves (231671, -2.67 DPS) [vendor] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 49.6 healing_power points (8.20 DPS) | yes | Chemist's Ring (15702, -2.64 DPS) [quest]; Maiden's Circle (13001, -3.10 DPS) [world_drop]; Glowing Crystal Ring (18402, -3.34 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 42.1 healing_power points (6.96 DPS) | yes | Maiden's Circle (13001, -1.86 DPS) [world_drop]; Glowing Crystal Ring (18402, -2.09 DPS) [quest]; Chemist's Ring (15702, -2.67 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (611.7 DPS) | yes | Evonice's Landin' Pilla (18951, -0.75 DPS) [quest]; Shard of the Splithooves (10659, -1.22 DPS) [quest]; Alchemist's Stone (13503, -3.02 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (611.7 DPS) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Evonice's Landin' Pilla (18951, -0.89 DPS, sim-verified) [quest]; Alchemist's Stone (13503, -1.47 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (611.7 DPS) | yes | Charstone Dirk (17710, -15.20 DPS) [dungeon]; Hand of Edward the Odd (2243, -17.65 DPS, sim-verified) [world_drop]; Hand of Righteousness (7721, -21.38 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; shoulder: Devout Mantle; back: Imperial Red Cloak; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket2: Uther's Strength

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

