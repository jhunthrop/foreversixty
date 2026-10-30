# Leveling BiS: Restoration

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-00000000000000000000-5050010000000000)

Set DPS (verified): 46.4. Weights run: 4.4s. Verify run: 15.1s. 186 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.000, intellect=1.289 ± 0.004, spirit=1.552 ± 0.008, mp5=3.307 ± 0.015, crit=0.044 ± 0.003 per rating point (14 rating = 1%, 0.609 per %), spell_haste=-5.588 ± 0.106

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.85 DPS) | yes | Wisdom's Leather Hood (252507, -0.13 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.23 DPS) [crafted]; Stormrider's Leather Hood (252506, -0.42 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 6.2 healing_power points (0.64 DPS) | yes | Scholarly Pendant (277203, +0.00 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.6 healing_power points (1.20 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS, sim-verified) [dungeon]; Forest Leather Mantle (4709, -0.56 DPS) [world_drop]; Prospector's Pads (14566, -0.56 DPS) [world_drop] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | 12.1 healing_power points (1.25 DPS) | yes | Heavy Woolen Cloak (4311, -0.61 DPS) [crafted]; Sylvan Cloak (4793, -0.61 DPS) [vendor]; Spirit Cloak (4792, -0.62 DPS, sim-verified) [vendor] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 26.1 healing_power points (2.69 DPS) | yes | Robe of the Moccasin (6465, +0.00 DPS, sim-verified) [dungeon]; Filigreed Pristine Gown (253901, -0.21 DPS) [crafted]; Corsair's Overshirt (5202, -0.28 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 16.4 healing_power points (1.69 DPS) | yes | Drakewing Bands (12999, -0.11 DPS, sim-verified) [world_drop]; Owl Bracers (4796, -1.02 DPS) [vendor]; Bravo's Armbands (270015, -1.05 DPS) [quest] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Magefist Gloves (12977, -0.23 DPS) [world_drop]; Bright Gloves (3066, -0.36 DPS) [world_drop]; Wisdom's Leather Gloves (252499, -1.18 DPS, sim-verified) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.2 healing_power points (1.66 DPS) | yes | Novice Ardent's Sash (253887, +0.00 DPS, sim-verified) [crafted]; Wisdom's Leather Belt (252433, -0.31 DPS) [crafted]; Keller's Girdle (2911, -0.60 DPS) [world_drop] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 33.9 healing_power points (3.50 DPS) | yes | Filigreed Pristine Leggings (253937, -0.26 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -1.61 DPS) [world_drop]; Dreamer's Leggings (270016, -2.04 DPS) [quest] |
| feet | Black Whelp Slippers (252424) | Leatherworking [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Boots (253889, -0.13 DPS) [crafted]; Kimbra Boots (6191, -0.45 DPS) [quest]; Wisdom's Leather Boots (252444, -0.79 DPS, sim-verified) [crafted] |
| finger1 | Band of Purification (12996) | World drop [world_drop] | 9.3 healing_power points (0.96 DPS) | yes | Lavishly Jeweled Ring (1156, -0.16 DPS) [dungeon]; Lorekeeper's Ring (20431, -0.28 DPS) [rep]; Volcanic Rock Ring (12053, -0.56 DPS) [world_drop] |
| finger2 | Deep Fathom Ring (6463) | Wailing Caverns: Mutanus the Devourer [dungeon] | 7.8 healing_power points (0.80 DPS) | yes | Lorekeeper's Ring (20431, -0.12 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.30 DPS, sim-verified) [dungeon]; Volcanic Rock Ring (12053, -0.40 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps [dungeon] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Twisted Chanter's Staff (890, -0.27 DPS) [world_drop]; Evocator's Blade (2567, -0.53 DPS) [dungeon]; Staff of Westfall (2042, -1.01 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Tarnished Locket; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Black Whelp Slippers; finger1: Band of Purification; finger2: Deep Fathom Ring; main_hand: Staff of the Blessed Seer

No-known-source sample (15 of 186, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 30 (night-elf, 0000000000000000-00000000000000000000-5050035110010000)

Set DPS (verified): 89.0. Weights run: 6.2s. Verify run: 21.2s. 346 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.454 ± 0.014, spirit=2.917 ± 0.018, mp5=4.441 ± 0.027, crit=0.052 ± 0.005 per rating point (14 rating = 1%, 0.735 per %), spell_haste=not significant (-1.023 ± 0.259)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 50.5 healing_power points (5.14 DPS) | yes | Nightsky Cowl (4039, -1.88 DPS) [world_drop]; Dokebi Hat (14584, -2.03 DPS) [world_drop]; Enduring Cap (3020, -2.46 DPS, sim-verified) [world_drop] |
| neck | Pendant of Myzrael (4614) (or Glowing Green Talisman (5002)) | Razorfen Downs: Splinterbone Captain [dungeon] | 17.5 healing_power points (1.78 DPS) | yes | Glowing Green Talisman (5002, +0.00 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.00 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.30 DPS) [world_drop] |
| shoulder | Mantle of Honor (3560) | Bride of the Embalmer [quest] | 30.6 healing_power points (3.12 DPS) | yes | Faerie Mantle (5820, -0.42 DPS, sim-verified) [quest]; Nightsky Mantle (4718, -0.45 DPS) [world_drop]; Insignia Mantle (4721, -1.04 DPS) [world_drop] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 30.3 healing_power points (3.09 DPS) | yes | Prelacy Cape (7004, +0.00 DPS, sim-verified) [quest]; Caretaker's Cape (19533, -0.58 DPS) [rep]; Repairman's Cape (9605, -0.75 DPS) [quest] |
| chest | Wisdom's Leather Tunic (252511) | Leatherworking [crafted] | 44.2 healing_power points (4.50 DPS) | yes | Pristine Gown (253961, -0.24 DPS) [crafted]; Beguiler Robes (7728, -0.35 DPS) [dungeon]; Pressed Felt Robe (1997, -0.49 DPS) [world] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 20.6 healing_power points (2.10 DPS) | yes | Nightsky Wristbands (6407, -0.32 DPS) [world_drop]; Dokebi Bracers (14580, -0.46 DPS) [world_drop]; Drakewing Bands (12999, -1.41 DPS, sim-verified) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 42.3 healing_power points (4.30 DPS) | yes | Shilly Mitts (9609, -1.55 DPS, sim-verified) [quest]; Zodiac Gloves (7106, -1.63 DPS) [quest]; Hotshot Pilot's Gloves (9491, -1.63 DPS) [dungeon] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 45.2 healing_power points (4.60 DPS) | yes | Dokebi Cord (14578, -2.23 DPS) [world_drop]; Silver-lined Belt (13011, -2.48 DPS, sim-verified) [world_drop]; Resilient Cord (14406, -2.53 DPS) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 51.6 healing_power points (5.25 DPS) | yes | Pristine Leggings (253987, -0.49 DPS) [crafted]; Stormrider's Leather Kilt (252518, -0.95 DPS) [crafted]; Earthen Leggings (253999, -1.45 DPS, sim-verified) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 41.8 healing_power points (4.26 DPS) | yes | Boots of the Enchanter (4325, -1.88 DPS) [crafted]; Stonecloth Boots (14408, -1.88 DPS) [world_drop]; Soggy Boots (274747, -2.08 DPS, sim-verified) [vendor] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 26.4 healing_power points (2.69 DPS) | yes | The Queen's Jewel (13094, -0.02 DPS) [world_drop]; Ring of Calm (6790, -0.61 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.86 DPS) [vendor] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 26.4 healing_power points (2.69 DPS) | yes | The Queen's Jewel (13094, -0.07 DPS, sim-verified) [world_drop]; Ring of Calm (6790, -0.61 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.86 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (89.0 DPS) | yes | Hardened Root Staff (1317, -2.21 DPS) [quest]; Gnarled Ash Staff (791, -2.47 DPS) [world_drop]; Manual Crowd Pummeler (9449, -3.60 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 30.5 healing_power points (3.11 DPS) | yes | Orb of Mistmantle (13031, -1.00 DPS) [world_drop]; Defective Samophlange (274743, -1.12 DPS, sim-verified) [vendor]; Strength of Will (4837, -1.32 DPS) [vendor] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Pendant of Myzrael; shoulder: Mantle of Honor; back: Glowing Thresher Cape; chest: Wisdom's Leather Tunic; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 346, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (night-elf, 0000000000000000-00000000000000000000-5050035153112000)

Set DPS (verified): 134.0. Weights run: 7.0s. Verify run: 23.9s. 459 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.000, intellect=1.347 ± 0.016, spirit=2.509 ± 0.011, mp5=4.854 ± 0.034, crit=0.056 ± 0.005 per rating point (14 rating = 1%, 0.782 per %), spell_haste=not significant (0.538 ± 0.314)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 59.5 healing_power points (6.11 DPS) | yes | Whitemane's Chapeau (7720, +0.00 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -1.17 DPS) [world_drop]; Earthen Silk Hood (254015, -1.27 DPS) [crafted] |
| neck | Necklace of Calisea (1714) | Zul'Farrak: Witch Doctor's Chest [dungeon] | 27.0 healing_power points (2.77 DPS) | yes | Amberglow Talisman (10824, +0.00 DPS, sim-verified) [quest]; Glowing Eye of Mordresh (10769, -0.33 DPS) [dungeon]; Explorers' League Commendation (7746, -0.45 DPS) [quest] |
| shoulder | Earthen Silk Shoulders (254033) | Tailoring [crafted] | 38.1 healing_power points (3.91 DPS) | yes | Inquisitor's Shawl (19507, -0.57 DPS) [dungeon]; Crimson Silk Shoulders (7059, -1.00 DPS) [crafted]; Sheepshear Mantle (13115, -1.03 DPS, sim-verified) [world_drop] |
| back | Caretaker's Cape (19532) | Silverwing Sentinels [rep] | 30.5 healing_power points (3.14 DPS) | yes | Prelacy Cape (7004, -0.46 DPS) [quest]; Ceremonial Centaur Blanket (6789, -0.54 DPS) [quest]; Glowing Thresher Cape (6901, -0.99 DPS, sim-verified) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 78.1 healing_power points (8.02 DPS) | yes | Dreamweave Vest (10021, -3.17 DPS) [crafted]; Civinad Robes (9623, -3.18 DPS) [quest]; Doomsayer's Robe (4746, -4.75 DPS, sim-verified) [quest] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 30.5 healing_power points (3.13 DPS) | yes | Mindthrust Bracers (1974, -0.94 DPS) [dungeon]; Earthen Silk Cuffs (254019, -0.99 DPS, sim-verified) [crafted]; Silkstream Cuffs (16791, -1.41 DPS) [quest] |
| hands | Earthen Silk Gloves (254017) | Tailoring [crafted] | 38.1 healing_power points (3.91 DPS) | yes | Mender's Leather Gloves (252530, -0.51 DPS) [crafted]; Warden's Gloves (14606, -0.52 DPS) [world_drop]; Gloves of Old (9395, -1.73 DPS, sim-verified) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 47.8 healing_power points (4.91 DPS) | yes | Mender's Leather Belt (252523, -1.00 DPS, sim-verified) [crafted]; Windchaser Cinch (14435, -1.49 DPS) [world_drop]; Sutarn's Ring (13105, -1.98 DPS) [world_drop] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 59.2 healing_power points (6.08 DPS) | yes | Stoneweaver Leggings (9407, -0.66 DPS, sim-verified) [dungeon]; Stormcloth Pants (10010, -0.67 DPS) [crafted]; Wisdom's Leather Leggings (252519, -1.15 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) | Leatherworking [crafted] | 51.0 healing_power points (5.24 DPS) | yes | Thoughtcast Boots (10578, -1.07 DPS) [dungeon]; Gilded Slippers (254001, -1.18 DPS) [crafted]; Furen's Boots (13100, -1.90 DPS, sim-verified) [world_drop] |
| finger1 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 29.1 healing_power points (2.99 DPS) | yes | Electrocutioner Lagnut (9447, -0.57 DPS) [dungeon]; Runemender's Seal (273031, -0.57 DPS) [world]; The Queen's Jewel (13094, -0.65 DPS) [world_drop] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | 24.3 healing_power points (2.50 DPS) | yes | Electrocutioner Lagnut (9447, -0.08 DPS) [dungeon]; Runemender's Seal (273031, -0.08 DPS) [world]; The Queen's Jewel (13094, -0.16 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (134.0 DPS) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (134.0 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -4.56 DPS) [dungeon]; Death Speaker Scepter (2816, -4.73 DPS) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 28.1 healing_power points (2.88 DPS) | yes | Beacon of Hope (9393, -0.41 DPS) [dungeon]; Orb of Lorica (11262, -0.54 DPS, sim-verified) [quest]; Silksand Star (15964, -0.56 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Necklace of Calisea; shoulder: Earthen Silk Shoulders; back: Caretaker's Cape; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Earthen Silk Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; finger1: Darkspear Signet; finger2: Welken Ring; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Hand of Righteousness

No-known-source sample (15 of 459, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 50 (night-elf, 4300000000000000-00000000000000000000-5050035153113200)

Set DPS (verified): 227.8. Weights run: 7.1s. Verify run: 25.6s. 598 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.003, intellect=1.528 ± 0.017, spirit=2.753 ± 0.017, mp5=4.398 ± 0.016, crit=0.072 ± 0.007 per rating point (14 rating = 1%, 1.009 per %), spell_haste=0.983 ± 0.178

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 84.5 healing_power points (8.38 DPS) | yes | Knight-Lieutenant's Restored Leather Helm (220874, +0.00 DPS, sim-verified) [vendor]; Engineer's Guild Headpiece (9534, -0.49 DPS) [quest]; Cassandra's Grace (13102, -1.56 DPS) [world_drop] |
| neck | Darkmoon Necklace (19303) | Lhara [vendor] | 35.6 healing_power points (3.53 DPS) | yes | Necklace of Calisea (1714, -0.55 DPS) [dungeon]; Horizon Choker (13085, -0.74 DPS, sim-verified) [world_drop]; Amberglow Talisman (10824, -0.80 DPS) [quest] |
| shoulder | Living Shoulders (15061) | Leatherworking [crafted] | 66.8 healing_power points (6.62 DPS) | yes | Mender's Leather Shoulders (252538, -0.97 DPS) [crafted]; Knight-Lieutenant's Restored Leather Spaulders (220876, -1.04 DPS, sim-verified) [vendor]; Nethergeld Shoulders (254049, -1.19 DPS) [crafted] |
| back | Caretaker's Cape (19531) | Silverwing Sentinels [rep] | 38.5 healing_power points (3.82 DPS) | yes | Darkspear Raider's Cloak (272076, -0.33 DPS) [vendor]; Arcane Cloak (8286, -0.54 DPS) [world_drop]; Wingveil Cloak (10802, -1.47 DPS, sim-verified) [dungeon] |
| chest | Ghostweave Vest (14141) | Tailoring [crafted] | 90.2 healing_power points (8.95 DPS) | yes | Feathered Breastplate (8349, -0.96 DPS, sim-verified) [crafted]; Stormcloth Vest (10020, -1.01 DPS) [crafted]; Knight's Restored Leather Jerkin (220878, -1.45 DPS) [vendor] |
| wrist | Mender's Leather Bracers (252543) | Leatherworking [crafted] | 46.2 healing_power points (4.58 DPS) | yes | Nethergeld Cuffs (254061, -0.17 DPS) [crafted]; Enchanted Kodo Bracers (13119, -1.25 DPS) [world_drop]; Venomshroud Armguards (14439, -1.79 DPS) [world_drop] |
| hands | Feralheart Gauntlets (226784) | Mokvar [vendor] | 90.4 healing_power points (8.96 DPS) | yes | Gilded Gloves (254095, -2.39 DPS) [crafted]; Mender's Leather Gauntlets (252551, -2.55 DPS, sim-verified) [crafted]; Earthenweave Gloves (254075, -3.38 DPS) [crafted] |
| waist | Earthenweave Cord (254077) | Tailoring [crafted] | 56.3 healing_power points (5.58 DPS) | yes | Mender's Leather Waistguard (252477, -0.79 DPS) [crafted]; Mender's Leather Belt (252523, -1.15 DPS) [crafted]; Gilded Cord (254037, -1.36 DPS, sim-verified) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 85.2 healing_power points (8.45 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -0.22 DPS) [dungeon]; Dalewind Trousers (13008, -0.27 DPS) [world_drop]; Jinxed Hoodoo Kilt (9474, -1.13 DPS, sim-verified) [dungeon] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 67.3 healing_power points (6.67 DPS) | yes | Sergeant Major's Restored Leather Boots (220884, +0.00 DPS, sim-verified) [vendor]; Earthenweave Boots (254093, -1.09 DPS) [crafted]; Furen's Boots (13100, -1.24 DPS) [world_drop] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 42.2 healing_power points (4.18 DPS) | yes | Darkspear Signet (272069, -1.13 DPS) [vendor]; Coldwater Ring (4550, -1.27 DPS) [quest]; Seedtime Hoop (9655, -1.45 DPS) [quest] |
| finger2 | Choking Band (11868) | Ogre Head On A Stick = Party [quest] | 35.8 healing_power points (3.55 DPS) | yes | Darkspear Signet (272069, -0.57 DPS, sim-verified) [vendor]; Coldwater Ring (4550, -0.64 DPS) [quest]; Seedtime Hoop (9655, -0.82 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (227.8 DPS) | yes | Evonice's Landin' Pilla (18951, -3.81 DPS) [quest]; Thunderbrew's Boot Flask (744, -4.36 DPS) [quest]; Uther's Strength (11302, -4.76 DPS) [world_drop] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (227.8 DPS) | yes | Evonice's Landin' Pilla (18951, -0.64 DPS, sim-verified) [quest]; Thunderbrew's Boot Flask (744, -1.09 DPS) [quest]; Uther's Strength (11302, -1.49 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (227.8 DPS) | yes | Charstone Dirk (17710, -7.68 DPS) [dungeon]; Hand of Righteousness (7721, -10.22 DPS) [dungeon]; Hammer of the Northern Wind (810, -33.72 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Darkmoon Necklace; shoulder: Living Shoulders; back: Caretaker's Cape; chest: Ghostweave Vest; wrist: Mender's Leather Bracers; hands: Feralheart Gauntlets; waist: Earthenweave Cord; legs: Senior Designer's Pantaloons; feet: Sandals of the Insurgent; finger1: Eye of Adaegus; finger2: Choking Band; trinket1: Darkspear Voodoo Seal; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 598, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (night-elf, 4532200100000000-00000000000000000000-5050035153113200)

Set DPS (verified): 370.2. Weights run: 7.2s. Verify run: 52.0s. 1244 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.022, intellect=2.234 ± 0.044, spirit=4.276 ± 0.034, mp5=6.865 ± 0.033, crit=0.098 ± 0.010 per rating point (14 rating = 1%, 1.366 per %), spell_haste=not significant (-0.039 ± 0.480)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 173.1 healing_power points (17.46 DPS) | yes | Lieutenant Commander's Dragonhide Headdress (227199, -2.67 DPS) [vendor]; Feralheart Headdress (226786, -3.26 DPS) [vendor]; Wildheart Cowl (16720, -5.64 DPS, sim-verified) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 92.3 healing_power points (9.31 DPS) | yes | Enchanted South Seas Kelp (19612, -3.18 DPS) [quest]; Pristine Enchanted South Seas Kelp (19613, -3.18 DPS) [quest]; Lady Alizabeth's Pendant (13002, -4.52 DPS, sim-verified) [world_drop] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 143.0 healing_power points (14.42 DPS) | yes | Feralheart Mantle (226785, +0.00 DPS, sim-verified) [vendor]; Lieutenant Commander's Dragonhide Pauldrons (227201, -2.32 DPS) [vendor]; Field Marshal's Dragonhide Pauldrons (231705, -3.38 DPS) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 64.3 healing_power points (6.49 DPS) | yes | Caretaker's Cape (19530, -0.42 DPS) [rep]; Cloak of Earth and Sky (279270, -0.45 DPS) [crafted]; Darkspear Raider's Cloak (272063, -2.10 DPS, sim-verified) [vendor] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | 151.9 healing_power points (15.33 DPS) | yes | Wildheart Vest (16706, -1.08 DPS) [dungeon]; Ghostweave Vest (14141, -1.16 DPS) [crafted]; Ironfeather Breastplate (15066, -1.20 DPS, sim-verified) [crafted] |
| wrist | Feralheart Bindings (226782) | Mokvar [vendor] | sim-verified (+6.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Dragonhide Bracers (16445, +0.00 DPS) [pvp]; Devout Bracers (16697, -1.36 DPS) [dungeon]; Bracers of Hope (22667, -6.60 DPS, sim-verified) [quest] |
| hands | Feralheart Gauntlets (226784) | Mokvar [vendor] | 125.5 healing_power points (12.66 DPS) | yes | Wildheart Gloves (16717, -0.76 DPS, sim-verified) [dungeon]; Devout Gloves (16692, -1.56 DPS) [dungeon]; Swarmtender's Gloves (275607, -2.42 DPS) [crafted] |
| waist | Feralheart Cord (226780) | Mokvar [vendor] | sim-verified (+5.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Elderwild Waistcord (279252, -0.72 DPS) [crafted]; Wisdom of the Timbermaw (19047, -1.34 DPS) [crafted]; Caretaker's Cord (272398, -5.54 DPS, sim-verified) [vendor] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 155.9 healing_power points (15.72 DPS) | yes | Knight-Captain's Dragonhide Legguards (227200, -0.25 DPS, sim-verified) [vendor]; Feralheart Pants (226787, -1.23 DPS) [vendor]; Elderwild Pants (279254, -1.33 DPS) [crafted] |
| feet | Feralheart Sandals (226781) | Mokvar [vendor] | 148.3 healing_power points (14.96 DPS) | yes | Incandescent Mooncloth Boots (227862, -0.63 DPS, sim-verified) [vendor]; Mooncloth Boots (15802, -3.78 DPS) [crafted]; Devout Sandals (16691, -4.26 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 89.4 healing_power points (9.02 DPS) | yes | Band of the Hierophant (13096, -2.02 DPS) [world_drop]; Channeler's Ring (272406, -2.42 DPS) [vendor]; Blessed Band of Light (272407, -2.44 DPS) [vendor] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 71.8 healing_power points (7.24 DPS) | yes | Channeler's Ring (272406, -0.64 DPS) [vendor]; Blessed Band of Light (272407, -0.66 DPS) [vendor]; Band of the Hierophant (13096, -1.10 DPS, sim-verified) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+4.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, -5.21 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -6.07 DPS) [quest]; Shard of the Splithooves (10659, -6.23 DPS) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Shard of the Splithooves (10659, -0.08 DPS) [quest]; Ankh of Life (1713, -0.25 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -9.98 DPS) [dungeon]; Staff of Hale Magefire (13000, -10.43 DPS) [world_drop]; Hand of Edward the Odd (2243, -12.57 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Mooncloth Vest; wrist: Feralheart Bindings; waist: Feralheart Cord; legs: Devout Skirt; feet: Feralheart Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Serenity Field

No-known-source sample (15 of 1244, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60, raid preset (night-elf, 4532200100000000-00000000000000000000-5050035153113200)

Set DPS (verified): 681.4. Weights run: 5.3s. Verify run: 49.0s. 1244 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.204, intellect=2.451 ± 0.080, spirit=2.281 ± 0.075, mp5=3.336 ± 0.099, crit=0.327 ± 0.030 per rating point (14 rating = 1%, 4.573 per %), spell_haste=not significant (-0.743 ± 1.216)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 149.0 healing_power points (24.64 DPS) | yes | Lieutenant Commander's Dragonhide Headdress (227199, -4.67 DPS) [vendor]; Sanctified Leather Helm (22689, -5.80 DPS) [quest]; Feralheart Headdress (226786, -8.35 DPS, sim-verified) [vendor] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 65.4 healing_power points (10.81 DPS) | yes | Enchanted South Seas Kelp (19612, -3.37 DPS) [quest]; Pristine Enchanted South Seas Kelp (19613, -3.37 DPS) [quest]; Drake Tooth Necklace (21531, -4.21 DPS, sim-verified) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 101.8 healing_power points (16.83 DPS) | yes | Lieutenant Commander's Dragonhide Pauldrons (227201, -1.16 DPS) [vendor]; Field Marshal's Dragonhide Pauldrons (231705, -1.32 DPS) [vendor]; Devout Mantle (16695, -12.81 DPS, sim-verified) [dungeon] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 66.5 healing_power points (10.99 DPS) | yes | Darkspear Raider's Cloak (272076, -3.44 DPS) [vendor]; Drape of Recovery (272413, -3.66 DPS) [vendor] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | 118.4 healing_power points (19.57 DPS) | yes | Feralheart Embrace (226783, -1.31 DPS) [vendor]; Field Marshal's Dragonhide Tunic (231702, -1.39 DPS) [vendor]; Devout Robe (16690, -2.06 DPS, sim-verified) [dungeon] |
| wrist | Feralheart Bindings (226782) | Mokvar [vendor] | sim-verified (681.4 DPS) | yes | Marshal's Dragonhide Bracers (16445, +0.00 DPS) [pvp]; Devout Bracers (16697, -1.94 DPS) [dungeon]; Bracers of Hope (22667, -9.68 DPS, sim-verified) [quest] |
| hands | Feralheart Gauntlets (226784) | Mokvar [vendor] | 92.0 healing_power points (15.21 DPS) | yes | Wildheart Gloves (16717, -1.83 DPS) [dungeon]; Marshal's Dragonhide Gauntlets (231706, -2.22 DPS) [pvp]; Raider Handwraps (272097, -15.37 DPS, sim-verified) [vendor] |
| waist | Feralheart Cord (226780) | Mokvar [vendor] | 85.9 healing_power points (14.20 DPS) | yes | Marshal's Dragonhide Waistguard (16447, -0.30 DPS) [pvp]; Devout Belt (16696, -0.89 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -10.79 DPS, sim-verified) [crafted] |
| legs | Knight-Captain's Dragonhide Legguards (227200) | Captain Dirgehammer [vendor] | 119.4 healing_power points (19.74 DPS) | yes | Feralheart Pants (226787, -1.35 DPS) [vendor]; Devout Skirt (16694, -1.73 DPS, sim-verified) [dungeon]; Marshal's Dragonhide Legguards (231703, -3.19 DPS) [pvp] |
| feet | Feralheart Sandals (226781) | Mokvar [vendor] | sim-verified (681.4 DPS) | yes | Mooncloth Boots (15802, -2.68 DPS) [crafted]; Devout Sandals (16691, -4.94 DPS) [dungeon]; Incandescent Mooncloth Boots (227862, -9.28 DPS, sim-verified) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 73.7 healing_power points (12.19 DPS) | yes | Channeler's Ring (272406, -3.50 DPS) [vendor]; Band of the Hierophant (13096, -3.99 DPS) [world_drop]; Blessed Band of Light (272407, -4.05 DPS) [vendor] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 59.8 healing_power points (9.89 DPS) | yes | Band of the Hierophant (13096, -1.69 DPS) [world_drop]; Blessed Band of Light (272407, -1.75 DPS) [vendor]; Channeler's Ring (272406, -3.42 DPS, sim-verified) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (681.4 DPS) | yes | Ankh of Life (1713, -3.75 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -4.50 DPS) [quest]; Shard of the Splithooves (10659, -4.96 DPS) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (681.4 DPS) | yes | Ankh of Life (1713, -2.42 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -3.17 DPS) [quest]; Darkmoon Card: Blue Dragon (19288, -4.45 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (681.4 DPS) | yes | Charstone Dirk (17710, -15.20 DPS) [dungeon]; Hand of Edward the Odd (2243, -20.11 DPS, sim-verified) [world_drop]; Hand of Righteousness (7721, -21.38 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Mooncloth Vest; wrist: Feralheart Bindings; waist: Feralheart Cord; legs: Knight-Captain's Dragonhide Legguards; feet: Feralheart Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Serenity Field

No-known-source sample (15 of 1244, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

## Horde

### Band 20 (tauren, 0000000000000000-00000000000000000000-5050010000000000)

Set DPS (verified): 47.6. Weights run: 4.4s. Verify run: 14.2s. 176 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.000, intellect=1.289 ± 0.004, spirit=1.552 ± 0.008, mp5=3.307 ± 0.015, crit=0.044 ± 0.003 per rating point (14 rating = 1%, 0.609 per %), spell_haste=-5.588 ± 0.106

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.85 DPS) | yes | Wisdom's Leather Hood (252507, -0.13 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.23 DPS) [crafted]; Stormrider's Leather Hood (252506, -0.42 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 4.7 healing_power points (0.48 DPS) | yes | Scholarly Pendant (277203, -1.98 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.6 healing_power points (1.20 DPS) | yes | Forest Leather Mantle (4709, -0.56 DPS) [world_drop]; Prospector's Pads (14566, -0.56 DPS) [world_drop]; Slime-encrusted Pads (6461, -1.80 DPS, sim-verified) [dungeon] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | 12.1 healing_power points (1.25 DPS) | yes | Sylvan Cloak (4793, -0.61 DPS) [vendor]; Antiquated Cloak (4799, -0.61 DPS) [vendor]; Spirit Cloak (4792, -1.32 DPS, sim-verified) [vendor] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 26.1 healing_power points (2.69 DPS) | yes | Filigreed Pristine Gown (253901, -0.21 DPS) [crafted]; Corsair's Overshirt (5202, -0.28 DPS) [dungeon]; Robe of the Moccasin (6465, -2.45 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 16.4 healing_power points (1.69 DPS) | yes | Crystalline Cuffs (14148, -0.94 DPS) [dungeon]; Owl Bracers (4796, -1.02 DPS) [vendor]; Drakewing Bands (12999, -1.63 DPS, sim-verified) [world_drop] |
| hands | Blight Gloves (279877) | The New Plague [quest] | 16.8 healing_power points (1.73 DPS) | yes | Pristine Gloves (253913, -0.20 DPS) [crafted]; Magefist Gloves (12977, -0.43 DPS) [world_drop]; Wisdom's Leather Gloves (252499, -1.10 DPS, sim-verified) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.2 healing_power points (1.66 DPS) | yes | Wisdom's Leather Belt (252433, -0.31 DPS) [crafted]; Keller's Girdle (2911, -0.60 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.73 DPS, sim-verified) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 33.9 healing_power points (3.50 DPS) | yes | Filigreed Pristine Leggings (253937, -0.27 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -1.61 DPS) [world_drop]; Filigreed Silky Leggings (253939, -2.06 DPS) [crafted] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 17.4 healing_power points (1.80 DPS) | yes | Pristine Boots (253889, -0.47 DPS) [crafted]; Smoldering Boots (3076, -0.84 DPS) [world]; Black Whelp Slippers (252424, -2.22 DPS, sim-verified) [crafted] |
| finger1 | Band of Purification (12996) | World drop [world_drop] | 9.3 healing_power points (0.96 DPS) | yes | Lavishly Jeweled Ring (1156, -0.16 DPS) [dungeon]; Advisor's Ring (20426, -0.28 DPS) [rep]; Loop of Sacrifice (281673, -0.30 DPS) [quest] |
| finger2 | Deep Fathom Ring (6463) | Wailing Caverns: Mutanus the Devourer [dungeon] | 7.8 healing_power points (0.80 DPS) | yes | Advisor's Ring (20426, -0.12 DPS) [rep]; Loop of Sacrifice (281673, -0.14 DPS) [quest]; Lavishly Jeweled Ring (1156, -2.04 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Orgrimmar (15444) | Hidden Enemies [quest] | 19.5 healing_power points (2.01 DPS) | yes | Staff of the Blessed Seer (2271, -0.07 DPS, sim-verified) [dungeon]; Advisor's Gnarled Staff (20425, -0.46 DPS) [pvp]; Gnarled Necromancer's Staff (251534, -0.68 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Blight Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Band of Purification; finger2: Deep Fathom Ring; main_hand: Staff of Orgrimmar

No-known-source sample (15 of 176, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 209617 Insignia of the Alliance; 241089 Scarlet Dagger; 248005 Northshire Hammer; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff; 263006 Scout Ranger's Tunic; 263007 Skyseer's Vest

### Band 30 (tauren, 0000000000000000-00000000000000000000-5050035110010000)

Set DPS (verified): 88.3. Weights run: 6.2s. Verify run: 21.4s. 336 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.454 ± 0.014, spirit=2.917 ± 0.018, mp5=4.441 ± 0.027, crit=0.052 ± 0.005 per rating point (14 rating = 1%, 0.735 per %), spell_haste=not significant (-1.023 ± 0.259)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 50.5 healing_power points (5.14 DPS) | yes | Enduring Cap (3020, -0.69 DPS) [world_drop]; Nightsky Cowl (4039, -1.88 DPS) [world_drop]; Dokebi Hat (14584, -2.03 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | sim-verified (88.3 DPS) | yes | Pendant of Myzrael (4614, +0.00 DPS) [dungeon]; Glowing Green Talisman (5002, +0.00 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.30 DPS) [world_drop] |
| shoulder | Ghostly Mantle (3324) | Deathstalkers in Shadowfang [quest] | 35.3 healing_power points (3.59 DPS) | yes | Nightsky Mantle (4718, -0.92 DPS) [world_drop]; Desert Shoulders (15457, -1.36 DPS) [quest]; Insignia Mantle (4721, -1.51 DPS) [world_drop] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 30.3 healing_power points (3.09 DPS) | yes | Battle Healer's Cloak (19529, -0.58 DPS) [rep]; Amy's Blanket (13005, -1.01 DPS) [world_drop]; Darkspear Raider's Cloak (272078, -1.01 DPS) [vendor] |
| chest | Wisdom's Leather Tunic (252511) | Leatherworking [crafted] | 44.2 healing_power points (4.50 DPS) | yes | Pristine Gown (253961, -0.24 DPS) [crafted]; Beguiler Robes (7728, -0.35 DPS) [dungeon]; Pressed Felt Robe (1997, -0.49 DPS) [world] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 20.6 healing_power points (2.10 DPS) | yes | Drakewing Bands (12999, -0.31 DPS) [world_drop]; Nightsky Wristbands (6407, -0.32 DPS) [world_drop]; Dokebi Bracers (14580, -0.46 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 42.3 healing_power points (4.30 DPS) | yes | Tattered Mittens (270030, -1.50 DPS) [quest]; Hotshot Pilot's Gloves (9491, -1.63 DPS) [dungeon]; Blight Gloves (279877, -1.78 DPS) [quest] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 45.2 healing_power points (4.60 DPS) | yes | Silver-lined Belt (13011, -2.08 DPS) [world_drop]; Dokebi Cord (14578, -2.23 DPS) [world_drop]; Lilac Sash (6780, -2.38 DPS) [quest] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 51.6 healing_power points (5.25 DPS) | yes | Earthen Leggings (253999, -0.46 DPS) [crafted]; Pristine Leggings (253987, -0.49 DPS) [crafted]; Stormrider's Leather Kilt (252518, -0.95 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 41.8 healing_power points (4.26 DPS) | yes | Soggy Boots (274747, -1.29 DPS) [vendor]; Boots of the Enchanter (4325, -1.88 DPS) [crafted]; Stonecloth Boots (14408, -1.88 DPS) [world_drop] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 26.4 healing_power points (2.69 DPS) | yes | The Queen's Jewel (13094, -0.02 DPS) [world_drop]; Solid Gold Band (279845, -0.49 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.86 DPS) [vendor] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 26.4 healing_power points (2.69 DPS) | yes | The Queen's Jewel (13094, -0.02 DPS) [world_drop]; Solid Gold Band (279845, -0.49 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.86 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gnarled Ash Staff (791, -2.47 DPS) [world_drop]; Staff of the Friar (3415, -2.92 DPS) [dungeon]; Spine of the Devourer (273040, -3.26 DPS) [world] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 30.5 healing_power points (3.11 DPS) | yes | Defective Samophlange (274743, -0.88 DPS) [vendor]; Orb of Mistmantle (13031, -1.00 DPS) [world_drop]; Strength of Will (4837, -1.32 DPS) [vendor] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Ghostly Mantle; back: Glowing Thresher Cape; chest: Wisdom's Leather Tunic; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 336, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (tauren, 0000000000000000-00000000000000000000-5050035153112000)

Set DPS (verified): 132.7. Weights run: 7.0s. Verify run: 25.9s. 443 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.000, intellect=1.347 ± 0.016, spirit=2.509 ± 0.011, mp5=4.854 ± 0.034, crit=0.056 ± 0.005 per rating point (14 rating = 1%, 0.782 per %), spell_haste=not significant (0.538 ± 0.314)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 59.5 healing_power points (6.11 DPS) | yes | Whitemane's Chapeau (7720, -0.90 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -1.17 DPS) [world_drop]; Earthen Silk Hood (254015, -1.27 DPS) [crafted] |
| neck | Necklace of Calisea (1714) | Zul'Farrak: Witch Doctor's Chest [dungeon] | 27.0 healing_power points (2.77 DPS) | yes | Glowing Eye of Mordresh (10769, -0.33 DPS) [dungeon]; Amberglow Talisman (10824, -0.72 DPS, sim-verified) [quest]; Triune Amulet (7722, -0.97 DPS) [dungeon] |
| shoulder | Earthen Silk Shoulders (254033) | Tailoring [crafted] | 38.1 healing_power points (3.91 DPS) | yes | Inquisitor's Shawl (19507, -0.57 DPS) [dungeon]; Ghostly Mantle (3324, -0.67 DPS) [quest]; Sheepshear Mantle (13115, -1.69 DPS, sim-verified) [world_drop] |
| back | Cloak of Blight (6832) | Nothing But The Truth [quest] | 32.6 healing_power points (3.35 DPS) | yes | Battle Healer's Cloak (19528, +0.00 DPS, sim-verified) [rep]; Glowing Thresher Cape (6901, -0.57 DPS) [dungeon]; Ceremonial Centaur Blanket (6789, -0.75 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 78.1 healing_power points (8.02 DPS) | yes | Dreamweave Vest (10021, -3.17 DPS) [crafted]; Civinad Robes (9623, -3.18 DPS) [quest]; Doomsayer's Robe (4746, -4.95 DPS, sim-verified) [quest] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 30.5 healing_power points (3.13 DPS) | yes | Earthen Silk Cuffs (254019, -0.18 DPS, sim-verified) [crafted]; Mindthrust Bracers (1974, -0.94 DPS) [dungeon]; Dryad's Wrist Bindings (19597, -1.01 DPS) [pvp] |
| hands | Earthen Silk Gloves (254017) | Tailoring [crafted] | 38.1 healing_power points (3.91 DPS) | yes | Mender's Leather Gloves (252530, -0.51 DPS) [crafted]; Warden's Gloves (14606, -0.52 DPS) [world_drop]; Gloves of Old (9395, -1.25 DPS, sim-verified) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 47.8 healing_power points (4.91 DPS) | yes | Mender's Leather Belt (252523, -1.24 DPS, sim-verified) [crafted]; Windchaser Cinch (14435, -1.49 DPS) [world_drop]; Sutarn's Ring (13105, -1.98 DPS) [world_drop] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 59.2 healing_power points (6.08 DPS) | yes | Stormcloth Pants (10010, -0.67 DPS) [crafted]; Stoneweaver Leggings (9407, -0.74 DPS, sim-verified) [dungeon]; Wisdom's Leather Leggings (252519, -1.15 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) | Leatherworking [crafted] | 51.0 healing_power points (5.24 DPS) | yes | Thoughtcast Boots (10578, -1.07 DPS) [dungeon]; Gilded Slippers (254001, -1.18 DPS) [crafted]; Furen's Boots (13100, -1.58 DPS, sim-verified) [world_drop] |
| finger1 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 29.1 healing_power points (2.99 DPS) | yes | Electrocutioner Lagnut (9447, -0.57 DPS) [dungeon]; Runemender's Seal (273031, -0.57 DPS) [world]; The Queen's Jewel (13094, -0.65 DPS) [world_drop] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | 24.3 healing_power points (2.50 DPS) | yes | Electrocutioner Lagnut (9447, -0.08 DPS) [dungeon]; Runemender's Seal (273031, -0.08 DPS) [world]; The Queen's Jewel (13094, -0.16 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (132.7 DPS) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (132.7 DPS) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (132.7 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -4.56 DPS) [dungeon]; Death Speaker Scepter (2816, -4.73 DPS) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 28.1 healing_power points (2.88 DPS) | yes | Silksand Star (15964, -0.56 DPS) [world_drop]; Aurora Sphere (7610, -0.88 DPS) [dungeon]; Beacon of Hope (9393, -1.00 DPS, sim-verified) [dungeon] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Necklace of Calisea; shoulder: Earthen Silk Shoulders; back: Cloak of Blight; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Earthen Silk Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; finger1: Darkspear Signet; finger2: Welken Ring; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Hand of Righteousness

No-known-source sample (15 of 443, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 50 (tauren, 4300000000000000-00000000000000000000-5050035153113200)

Set DPS (verified): 226.4. Weights run: 7.1s. Verify run: 24.3s. 578 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.003, intellect=1.528 ± 0.017, spirit=2.753 ± 0.017, mp5=4.398 ± 0.016, crit=0.072 ± 0.007 per rating point (14 rating = 1%, 1.009 per %), spell_haste=0.983 ± 0.178

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 84.5 healing_power points (8.38 DPS) | yes | Engineer's Guild Headpiece (9534, +0.00 DPS, sim-verified) [quest]; Blood Guard's Restored Leather Helm (220875, -0.48 DPS) [vendor]; Cassandra's Grace (13102, -1.56 DPS) [world_drop] |
| neck | Darkmoon Necklace (19303) | Lhara [vendor] | 35.6 healing_power points (3.53 DPS) | yes | Horizon Choker (13085, -0.49 DPS, sim-verified) [world_drop]; Necklace of Calisea (1714, -0.55 DPS) [dungeon]; Choker of the High Shaman (4112, -0.58 DPS) [quest] |
| shoulder | Living Shoulders (15061) | Leatherworking [crafted] | 66.8 healing_power points (6.62 DPS) | yes | Blood Guard's Restored Leather Spaulders (220877, -0.20 DPS) [vendor]; Nethergeld Shoulders (254049, -1.19 DPS) [crafted]; Mender's Leather Shoulders (252538, -2.00 DPS, sim-verified) [crafted] |
| back | Battle Healer's Cloak (19527) | Warsong Outriders [rep] | 38.5 healing_power points (3.82 DPS) | yes | Wingveil Cloak (10802, -0.32 DPS) [dungeon]; Darkspear Raider's Cloak (272076, -0.33 DPS) [vendor]; Cloak of Blight (6832, -1.75 DPS, sim-verified) [quest] |
| chest | Ghostweave Vest (14141) | Tailoring [crafted] | 90.2 healing_power points (8.95 DPS) | yes | Stormcloth Vest (10020, -1.01 DPS) [crafted]; Feathered Breastplate (8349, -1.35 DPS, sim-verified) [crafted]; Stone Guard's Restored Leather Jerkin (220879, -1.45 DPS) [vendor] |
| wrist | Mender's Leather Bracers (252543) | Leatherworking [crafted] | 46.2 healing_power points (4.58 DPS) | yes | Nethergeld Cuffs (254061, -0.17 DPS) [crafted]; Enchanted Kodo Bracers (13119, -1.25 DPS) [world_drop]; Venomshroud Armguards (14439, -1.79 DPS) [world_drop] |
| hands | Feralheart Gauntlets (226784) | Mokvar [vendor] | 90.4 healing_power points (8.96 DPS) | yes | Mender's Leather Gauntlets (252551, -1.94 DPS, sim-verified) [crafted]; Gilded Gloves (254095, -2.39 DPS) [crafted]; Earthenweave Gloves (254075, -3.38 DPS) [crafted] |
| waist | Earthenweave Cord (254077) | Tailoring [crafted] | 56.3 healing_power points (5.58 DPS) | yes | Mender's Leather Waistguard (252477, -0.79 DPS) [crafted]; Mender's Leather Belt (252523, -1.15 DPS) [crafted]; Gilded Cord (254037, -1.60 DPS, sim-verified) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 85.2 healing_power points (8.45 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -0.22 DPS) [dungeon]; Dalewind Trousers (13008, -0.27 DPS) [world_drop]; Jinxed Hoodoo Kilt (9474, -0.73 DPS, sim-verified) [dungeon] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 67.3 healing_power points (6.67 DPS) | yes | First Sergeant's Restored Leather Boots (220885, +0.00 DPS, sim-verified) [vendor]; Earthenweave Boots (254093, -1.09 DPS) [crafted]; Furen's Boots (13100, -1.24 DPS) [world_drop] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 42.2 healing_power points (4.18 DPS) | yes | Coldwater Ring (4550, -1.27 DPS) [quest]; Welken Ring (5011, -1.52 DPS) [world_drop]; Electrocutioner Lagnut (9447, -1.68 DPS) [dungeon] |
| finger2 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 30.8 healing_power points (3.05 DPS) | yes | Coldwater Ring (4550, -0.14 DPS) [quest]; Welken Ring (5011, -0.38 DPS) [world_drop]; Electrocutioner Lagnut (9447, -0.55 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (226.4 DPS) | yes | Evonice's Landin' Pilla (18951, -3.81 DPS) [quest]; Uther's Strength (11302, -4.76 DPS) [world_drop]; Alchemist's Stone (13503, -5.45 DPS) [crafted] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (226.4 DPS) | yes | Evonice's Landin' Pilla (18951, -0.50 DPS, sim-verified) [quest]; Uther's Strength (11302, -1.49 DPS) [world_drop]; Alchemist's Stone (13503, -2.18 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (226.4 DPS) | yes | Charstone Dirk (17710, -7.68 DPS) [dungeon]; Hand of Righteousness (7721, -10.22 DPS) [dungeon]; Hammer of the Northern Wind (810, -33.94 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Darkmoon Necklace; shoulder: Living Shoulders; back: Battle Healer's Cloak; chest: Ghostweave Vest; wrist: Mender's Leather Bracers; hands: Feralheart Gauntlets; waist: Earthenweave Cord; legs: Senior Designer's Pantaloons; feet: Sandals of the Insurgent; finger1: Eye of Adaegus; finger2: Darkspear Signet; trinket1: Darkspear Voodoo Seal; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 578, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60 (tauren, 4532200100000000-00000000000000000000-5050035153113200)

Set DPS (verified): 368.1. Weights run: 7.2s. Verify run: 49.5s. 1238 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.022, intellect=2.234 ± 0.044, spirit=4.276 ± 0.034, mp5=6.865 ± 0.033, crit=0.098 ± 0.010 per rating point (14 rating = 1%, 1.366 per %), spell_haste=not significant (-0.039 ± 0.480)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 173.1 healing_power points (17.46 DPS) | yes | Champion's Dragonhide Headdress (227205, -2.67 DPS) [vendor]; Feralheart Headdress (226786, -3.26 DPS) [vendor]; Wildheart Cowl (16720, -6.63 DPS, sim-verified) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 92.3 healing_power points (9.31 DPS) | yes | Enchanted South Seas Kelp (19612, -3.18 DPS) [quest]; Pristine Enchanted South Seas Kelp (19613, -3.18 DPS) [quest]; Lady Alizabeth's Pendant (13002, -4.18 DPS, sim-verified) [world_drop] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 143.0 healing_power points (14.42 DPS) | yes | Feralheart Mantle (226785, -0.32 DPS, sim-verified) [vendor]; Champion's Dragonhide Pauldrons (227207, -2.32 DPS) [vendor]; Warlord's Dragonhide Pauldrons (231672, -3.38 DPS) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 64.3 healing_power points (6.49 DPS) | yes | Battle Healer's Cloak (19526, -0.42 DPS) [rep]; Cloak of Earth and Sky (279270, -0.45 DPS) [crafted]; Darkspear Raider's Cloak (272063, -2.60 DPS, sim-verified) [vendor] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | 151.9 healing_power points (15.33 DPS) | yes | Wildheart Vest (16706, -1.08 DPS) [dungeon]; Ghostweave Vest (14141, -1.16 DPS) [crafted]; Ironfeather Breastplate (15066, -1.30 DPS, sim-verified) [crafted] |
| wrist | Feralheart Bindings (226782) | Mokvar [vendor] | sim-verified (+6.0 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Dragonhide Bracers (16553, +0.00 DPS) [pvp]; Devout Bracers (16697, -1.36 DPS) [dungeon]; Bracers of Hope (22667, -5.98 DPS, sim-verified) [quest] |
| hands | Feralheart Gauntlets (226784) | Mokvar [vendor] | 125.5 healing_power points (12.66 DPS) | yes | Wildheart Gloves (16717, -1.05 DPS, sim-verified) [dungeon]; Devout Gloves (16692, -1.56 DPS) [dungeon]; Swarmtender's Gloves (275607, -2.42 DPS) [crafted] |
| waist | Feralheart Cord (226780) | Mokvar [vendor] | sim-verified (+5.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Elderwild Waistcord (279252, -0.72 DPS) [crafted]; Wisdom of the Timbermaw (19047, -1.34 DPS) [crafted]; Caretaker's Cord (272398, -4.97 DPS, sim-verified) [vendor] |
| legs | Feralheart Pants (226787) | Mokvar [vendor] | sim-verified (+5.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Legionnaire's Dragonhide Legguards (227206, +0.00 DPS) [vendor]; Elderwild Pants (279254, -0.11 DPS) [crafted]; Devout Skirt (16694, -5.73 DPS, sim-verified) [dungeon] |
| feet | Feralheart Sandals (226781) | Mokvar [vendor] | 148.3 healing_power points (14.96 DPS) | yes | Incandescent Mooncloth Boots (227862, -1.38 DPS, sim-verified) [vendor]; Mooncloth Boots (15802, -3.78 DPS) [crafted]; Devout Sandals (16691, -4.26 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 89.4 healing_power points (9.02 DPS) | yes | Band of the Hierophant (13096, -2.02 DPS) [world_drop]; Channeler's Ring (272406, -2.42 DPS) [vendor]; Blessed Band of Light (272407, -2.44 DPS) [vendor] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 71.8 healing_power points (7.24 DPS) | yes | Channeler's Ring (272406, -0.64 DPS) [vendor]; Blessed Band of Light (272407, -0.66 DPS) [vendor]; Band of the Hierophant (13096, -1.14 DPS, sim-verified) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+5.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, -5.21 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -6.07 DPS) [quest]; Shard of the Splithooves (10659, -6.23 DPS) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Shard of the Splithooves (10659, -0.08 DPS) [quest]; Ankh of Life (1713, -0.88 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -9.98 DPS) [dungeon]; Staff of Hale Magefire (13000, -10.43 DPS) [world_drop]; Hand of Edward the Odd (2243, -12.93 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Mooncloth Vest; wrist: Feralheart Bindings; waist: Feralheart Cord; legs: Feralheart Pants; feet: Feralheart Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Serenity Field

No-known-source sample (15 of 1238, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60, raid preset (tauren, 4532200100000000-00000000000000000000-5050035153113200)

Set DPS (verified): 678.4. Weights run: 5.3s. Verify run: 47.3s. 1238 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.204, intellect=2.451 ± 0.080, spirit=2.281 ± 0.075, mp5=3.336 ± 0.099, crit=0.327 ± 0.030 per rating point (14 rating = 1%, 4.573 per %), spell_haste=not significant (-0.743 ± 1.216)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 149.0 healing_power points (24.64 DPS) | yes | Champion's Dragonhide Headdress (227205, -4.67 DPS) [vendor]; Sanctified Leather Helm (22689, -5.80 DPS) [quest]; Feralheart Headdress (226786, -9.01 DPS, sim-verified) [vendor] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 65.4 healing_power points (10.81 DPS) | yes | Enchanted South Seas Kelp (19612, -3.37 DPS) [quest]; Pristine Enchanted South Seas Kelp (19613, -3.37 DPS) [quest]; Drake Tooth Necklace (21531, -4.52 DPS, sim-verified) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 101.8 healing_power points (16.83 DPS) | yes | Champion's Dragonhide Pauldrons (227207, -1.16 DPS) [vendor]; Warlord's Dragonhide Pauldrons (231672, -1.32 DPS) [vendor]; Devout Mantle (16695, -12.61 DPS, sim-verified) [dungeon] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 66.5 healing_power points (10.99 DPS) | yes | Darkspear Raider's Cloak (272076, -3.44 DPS) [vendor]; Drape of Recovery (272413, -3.66 DPS) [vendor] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | 118.4 healing_power points (19.57 DPS) | yes | Feralheart Embrace (226783, -1.31 DPS) [vendor]; Warlord's Dragonhide Tunic (231674, -1.39 DPS) [vendor]; Devout Robe (16690, -2.52 DPS, sim-verified) [dungeon] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 72.3 healing_power points (11.96 DPS) | yes | General's Dragonhide Bracers (16553, -0.49 DPS) [pvp]; Feralheart Bindings (226782, -0.98 DPS, sim-verified) [vendor]; Devout Bracers (16697, -3.13 DPS) [dungeon] |
| hands | Feralheart Gauntlets (226784) | Mokvar [vendor] | 92.0 healing_power points (15.21 DPS) | yes | Wildheart Gloves (16717, -1.83 DPS) [dungeon]; General's Dragonhide Gauntlets (231676, -2.22 DPS) [vendor]; Raider Handwraps (272097, -15.77 DPS, sim-verified) [vendor] |
| waist | Feralheart Cord (226780) | Mokvar [vendor] | 85.9 healing_power points (14.20 DPS) | yes | General's Dragonhide Belt (16556, -0.30 DPS) [pvp]; Devout Belt (16696, -0.89 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -10.49 DPS, sim-verified) [crafted] |
| legs | Feralheart Pants (226787) | Mokvar [vendor] | sim-verified (678.4 DPS) | yes | Legionnaire's Dragonhide Legguards (227206, +0.00 DPS) [vendor]; General's Dragonhide Legguards (231673, -1.84 DPS) [vendor]; Devout Skirt (16694, -10.25 DPS, sim-verified) [dungeon] |
| feet | Feralheart Sandals (226781) | Mokvar [vendor] | sim-verified (678.4 DPS) | yes | Mooncloth Boots (15802, -2.68 DPS) [crafted]; Devout Sandals (16691, -4.94 DPS) [dungeon]; Incandescent Mooncloth Boots (227862, -9.22 DPS, sim-verified) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 73.7 healing_power points (12.19 DPS) | yes | Channeler's Ring (272406, -3.50 DPS) [vendor]; Band of the Hierophant (13096, -3.99 DPS) [world_drop]; Blessed Band of Light (272407, -4.05 DPS) [vendor] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 59.8 healing_power points (9.89 DPS) | yes | Band of the Hierophant (13096, -1.69 DPS) [world_drop]; Blessed Band of Light (272407, -1.75 DPS) [vendor]; Channeler's Ring (272406, -3.89 DPS, sim-verified) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (678.4 DPS) | yes | Ankh of Life (1713, -3.75 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -4.50 DPS) [quest]; Shard of the Splithooves (10659, -4.96 DPS) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (678.4 DPS) | yes | Evonice's Landin' Pilla (18951, -3.17 DPS) [quest]; Shard of the Splithooves (10659, -3.63 DPS) [quest]; Ankh of Life (1713, -4.82 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (678.4 DPS) | yes | Charstone Dirk (17710, -15.20 DPS) [dungeon]; Hand of Edward the Odd (2243, -20.65 DPS, sim-verified) [world_drop]; Hand of Righteousness (7721, -21.38 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Mooncloth Vest; wrist: Bracers of Hope; waist: Feralheart Cord; legs: Feralheart Pants; feet: Feralheart Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Serenity Field

No-known-source sample (15 of 1238, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

