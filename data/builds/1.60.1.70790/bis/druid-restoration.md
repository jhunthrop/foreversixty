# Leveling BiS: Restoration

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-00000000000000000000-2503000100000000)

Set DPS (verified): 75.8. Weights run: 6.1s. Verify run: 44.2s. 228 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.021, intellect=0.702 ± 0.003, spirit=1.744 ± 0.007, mp5=1.987 ± 0.009, crit=0.043 ± 0.001 per rating point (14 rating = 1%, 0.599 per %), spell_haste=not significant (-0.094 ± 0.026)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (4.15 DPS) | yes | Wisdom's Leather Hood (252507, -0.06 DPS, sim-verified) [crafted]; Stormrider's Leather Hood (252506, -0.53 DPS) [crafted]; Shadow Goggles (4373, -0.93 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 3.5 healing_power points (0.80 DPS) | yes | Tarnished Locket (279870, -0.19 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 12.2 healing_power points (2.82 DPS) | yes | Magician's Mantle (12998, -0.20 DPS) [world_drop]; Reinforced Woolen Shoulders (4315, -1.01 DPS) [crafted]; Forest Leather Mantle (4709, -1.21 DPS) [world_drop] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | sim-verified (75.8 DPS) | yes | Heavy Woolen Cloak (4311, -0.35 DPS) [crafted]; Spirit Cloak (4792, -1.27 DPS) [vendor]; Sylvan Cloak (4793, -1.27 DPS) [vendor] |
| chest | Robe of the Moccasin (6465) | Wailing Caverns: Lord Cobrahn [dungeon] | 26.4 healing_power points (6.10 DPS) | yes | Corsair's Overshirt (5202, -0.12 DPS, sim-verified) [dungeon]; Scholarly Robes (2034, -0.58 DPS) [dungeon]; Wisdom's Leather Armor (252493, -0.62 DPS) [crafted] |
| wrist | Drakewing Bands (12999) | World drop [world_drop] | 10.5 healing_power points (2.41 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Repurposed Hair Band (281256, -1.28 DPS) [quest]; Seer's Cuffs (3645, -1.45 DPS) [dungeon] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 13.1 healing_power points (3.02 DPS) | yes | Wisdom's Leather Gloves (252499, -0.07 DPS) [crafted]; Magefist Gloves (12977, -0.60 DPS) [world_drop]; Bright Gloves (3066, -0.77 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 13.8 healing_power points (3.18 DPS) | yes | Novice Ardent's Sash (253887, -0.11 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.69 DPS) [crafted]; Stormrider's Leather Belt (252432, -1.61 DPS) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 31.2 healing_power points (7.19 DPS) | yes | Filigreed Pristine Leggings (253937, -0.09 DPS, sim-verified) [crafted]; Stormrider's Leather Pants (252502, -2.31 DPS) [crafted]; Darkweave Breeches (12987, -3.65 DPS) [world_drop] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 14.5 healing_power points (3.35 DPS) | yes | Red Woolen Boots (4313, -0.09 DPS, sim-verified) [crafted]; Black Whelp Slippers (252424, -0.62 DPS) [crafted]; Pristine Boots (253889, -0.79 DPS) [crafted] |
| finger1 | Band of Purification (12996) | World drop [world_drop] | 10.5 healing_power points (2.41 DPS) | yes | Minor Channeling Ring (1449, -0.94 DPS) [quest]; Lavishly Jeweled Ring (1156, -1.44 DPS) [dungeon]; Demon Band (12054, -1.61 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 9.0 healing_power points (2.07 DPS) | yes | Minor Channeling Ring (1449, -0.13 DPS, sim-verified) [quest]; Lavishly Jeweled Ring (1156, -1.10 DPS) [dungeon]; Demon Band (12054, -1.27 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Westfall (2042) | The Defias Brotherhood [quest] | 62.0 healing_power points (14.29 DPS) | yes | Living Root (6631, -0.61 DPS, sim-verified) [dungeon]; Staff of the Blessed Seer (2271, -3.81 DPS) [dungeon]; Evocator's Blade (2567, -5.50 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Scholarly Pendant; shoulder: Serpent's Shoulders; back: Caretaker's Cape; chest: Robe of the Moccasin; wrist: Drakewing Bands; hands: Pristine Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Band of Purification; finger2: Lorekeeper's Ring; main_hand: Staff of Westfall

No-known-source sample (15 of 228, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (night-elf, 0000000000000000-00000000000000000000-5553001100010000)

Set DPS (verified): 121.5. Weights run: 8.4s. Verify run: 77.6s. 400 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.049, intellect=1.337 ± 0.019, spirit=2.833 ± 0.018, mp5=3.849 ± 0.023, crit=0.105 ± 0.005 per rating point (14 rating = 1%, 1.470 per %), spell_haste=not significant (-0.066 ± 0.171)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 50.0 healing_power points (7.91 DPS) | yes | Enduring Cap (3020, -1.83 DPS, sim-verified) [world_drop]; Nightsky Cowl (4039, -3.13 DPS) [world_drop]; Dokebi Hat (14584, -3.34 DPS) [world_drop] |
| neck | Pendant of Myzrael (4614) (or Glowing Green Talisman (5002)) | Razorfen Downs: Splinterbone Captain [dungeon] | 17.0 healing_power points (2.69 DPS) | yes | Glowing Green Talisman (5002, +0.00 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.05 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.50 DPS) [world_drop] |
| shoulder | Mantle of Honor (3560) | Bride of the Embalmer [quest] | 29.2 healing_power points (4.62 DPS) | yes | Faerie Mantle (5820, -0.61 DPS, sim-verified) [quest]; Nightsky Mantle (4718, -0.68 DPS) [world_drop]; Bloodmage Mantle (7684, -1.29 DPS) [dungeon] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | sim-verified (121.5 DPS) | yes | Caretaker's Cape (19533, -0.53 DPS, sim-verified) [rep]; Repairman's Cape (9605, -0.87 DPS) [quest]; Amy's Blanket (13005, -1.55 DPS) [world_drop] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 47.7 healing_power points (7.54 DPS) | yes | Wisdom's Leather Tunic (252511, -0.53 DPS, sim-verified) [crafted]; Pristine Gown (253961, -1.11 DPS) [crafted]; Pressed Felt Robe (1997, -1.53 DPS) [world] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 20.5 healing_power points (3.25 DPS) | yes | Drakewing Bands (12999, -0.56 DPS) [world_drop]; Mindthrust Bracers (1974, -0.57 DPS, sim-verified) [dungeon]; Nightsky Wristbands (6407, -0.64 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 40.7 healing_power points (6.43 DPS) | yes | Zodiac Gloves (7106, -2.04 DPS, sim-verified) [quest]; Hotshot Pilot's Gloves (9491, -2.50 DPS) [dungeon]; Gloves of Meditation (4318, -2.51 DPS) [crafted] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 44.0 healing_power points (6.96 DPS) | yes | Skycaller's Leather Belt (252522, -1.96 DPS, sim-verified) [crafted]; Belt of Arugal (6392, -2.66 DPS) [dungeon]; Silver-lined Belt (13011, -3.27 DPS) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 50.2 healing_power points (7.94 DPS) | yes | Pristine Leggings (253987, -0.74 DPS) [crafted]; Earthen Leggings (253999, -0.74 DPS, sim-verified) [crafted]; Stormrider's Leather Kilt (252518, -1.53 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 40.7 healing_power points (6.43 DPS) | yes | Boots of the Enchanter (4325, -2.05 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -2.16 DPS) [dungeon]; Glinteye Slippers (273024, -2.64 DPS) [dungeon] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 28.8 healing_power points (4.56 DPS) | yes | The Queen's Jewel (13094, -0.55 DPS) [world_drop]; Ring of Calm (6790, -1.42 DPS) [quest]; Ring of Power Regulation (273030, -1.78 DPS) [dungeon] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 28.8 healing_power points (4.56 DPS) | yes | The Queen's Jewel (13094, -0.70 DPS, sim-verified) [world_drop]; Ring of Calm (6790, -1.42 DPS) [quest]; Ring of Power Regulation (273030, -1.78 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Friar (3415) | Blackfathom Deeps [dungeon] | sim-verified (121.5 DPS) | yes | Odo's Ley Staff (6318, -0.45 DPS, sim-verified) [dungeon]; Staff of Westfall (2042, -2.90 DPS) [quest]; Death Speaker Scepter (2816, -3.48 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Pendant of Myzrael; shoulder: Mantle of Honor; back: Glowing Thresher Cape; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Staff of the Friar

No-known-source sample (15 of 400, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 40 (night-elf, 0000000000000000-00000000000000000000-5553005151010000)

Set DPS (verified): 213.3. Weights run: 8.6s. Verify run: 78.9s. 526 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.170, intellect=1.132 ± 0.033, spirit=2.488 ± 0.023, mp5=3.698 ± 0.026, crit=0.252 ± 0.019 per rating point (14 rating = 1%, 3.534 per %), spell_haste=not significant (0.594 ± 0.341)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 63.6 healing_power points (13.37 DPS) | yes | Whitemane's Chapeau (7720, -0.69 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -3.30 DPS) [world_drop]; Earthen Silk Hood (254015, -3.52 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 29.1 healing_power points (6.10 DPS) | yes | Triune Amulet (7722, -0.54 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -0.78 DPS) [dungeon]; Amberglow Talisman (10824, -0.88 DPS) [quest] |
| shoulder | Earthen Silk Shoulders (254033) | Tailoring [crafted] | 37.9 healing_power points (7.96 DPS) | yes | Inquisitor's Shawl (19507, -0.46 DPS, sim-verified) [dungeon]; Mantle of Doan (7712, -0.88 DPS) [dungeon]; Sheepshear Mantle (13115, -1.02 DPS) [world_drop] |
| back | Caretaker's Cape (19532) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Thresher Cape (6901, -0.74 DPS) [dungeon]; Energy Cloak (9397, -1.17 DPS) [dungeon]; Ceremonial Centaur Blanket (6789, -1.21 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 77.9 healing_power points (16.36 DPS) | yes | Dreamweave Vest (10021, -1.24 DPS, sim-verified) [crafted]; Doomsayer's Robe (4746, -3.60 DPS) [quest]; Civinad Robes (9623, -4.63 DPS) [quest] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 29.4 healing_power points (6.18 DPS) | yes | Spidertank Oilrag (9448, -1.96 DPS) [dungeon]; Earthen Silk Cuffs (254019, -2.25 DPS, sim-verified) [crafted]; Mindthrust Bracers (1974, -2.66 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Earthen Silk Gloves (254017, -0.50 DPS, sim-verified) [crafted]; Gloves of Old (9395, -0.93 DPS) [world_drop]; Windchaser Handguards (14431, -1.60 DPS) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 46.0 healing_power points (9.66 DPS) | yes | Mender's Leather Belt (252523, -1.07 DPS, sim-verified) [crafted]; Windchaser Cinch (14435, -3.05 DPS) [world_drop]; Earthen Silk Belt (7061, -3.80 DPS) [crafted] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 65.6 healing_power points (13.78 DPS) | yes | Stoneweaver Leggings (9407, -1.25 DPS, sim-verified) [dungeon]; Stormcloth Pants (10010, -2.79 DPS) [crafted]; Wisdom's Leather Leggings (252519, -4.05 DPS) [crafted] |
| feet | Furen's Boots (13100) | World drop [world_drop] | sim-verified (213.4 DPS) | yes | Thoughtcast Boots (10578, -0.50 DPS) [dungeon]; Mender's Leather Shoes (252533, -0.99 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -2.07 DPS) [crafted] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 26.4 healing_power points (5.55 DPS) | yes | Welken Ring (5011, -0.70 DPS) [world_drop]; The Queen's Jewel (13094, -0.89 DPS) [world_drop]; Lorekeeper's Ring (19524, -1.33 DPS) [rep] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 26.4 healing_power points (5.55 DPS) | yes | Welken Ring (5011, -0.33 DPS, sim-verified) [world_drop]; The Queen's Jewel (13094, -0.89 DPS) [world_drop]; Lorekeeper's Ring (19524, -1.33 DPS) [rep] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+9.6 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of the Friar (3415, -6.02 DPS) [dungeon]; Staff of Jordan (873, -6.05 DPS, sim-verified) [world_drop]; Odo's Ley Staff (6318, -6.73 DPS) [dungeon] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | 32.0 healing_power points (6.71 DPS) | yes | Orb of Souls (249395, -0.43 DPS, sim-verified) [crafted]; Orb of Lorica (11262, -1.81 DPS) [quest]; Silksand Star (15964, -2.01 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Earthen Silk Shoulders; back: Caretaker's Cape; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Dreamweave Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Furen's Boots; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Hand of Righteousness; off_hand: Beacon of Hope

No-known-source sample (15 of 526, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 50 (night-elf, 0100000000000000-00000000000000000000-5553005153113210)

Set DPS (verified): 343.1. Weights run: 7.1s. Verify run: 83.2s. 690 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.118, intellect=2.737 ± 0.092, spirit=12.203 ± 0.088, mp5=11.091 ± 0.108, crit=0.586 ± 0.025 per rating point (14 rating = 1%, 8.200 per %), spell_haste=not significant (-0.389 ± 2.346)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Engineer's Guild Headpiece (9534) | Divino-matic Rod [quest] | 287.5 healing_power points (19.57 DPS) | yes | Gemburst Circlet (10751, -0.84 DPS) [quest]; Imperial Red Circlet (8254, -4.13 DPS) [world_drop]; Serpentskin Helm (8261, -4.13 DPS) [world_drop] |
| neck | Lei of Lilies (1315) | World drop [world_drop] | 183.0 healing_power points (12.46 DPS) | yes | Talvash's Enhancing Necklace (7673, -3.35 DPS, sim-verified) [quest]; Amberglow Talisman (10824, -4.15 DPS) [quest]; Explorers' League Commendation (7746, -4.98 DPS) [quest] |
| shoulder | Living Shoulders (15061) | Leatherworking [crafted] | 189.6 healing_power points (12.91 DPS) | yes | Ironfeather Shoulders (15067, -1.11 DPS, sim-verified) [crafted]; Knight-Lieutenant's Restored Leather Spaulders (220876, -1.24 DPS) [vendor]; Serpentskin Spaulders (8263, -1.54 DPS) [world_drop] |
| back | Arcane Cloak (8286) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Featherskin Cape (10843, +0.00 DPS) [world]; Wingveil Cloak (10802, -0.49 DPS) [dungeon]; Traveler's Cloak (8297, -0.83 DPS) [world_drop] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 412.6 healing_power points (28.09 DPS) | yes | Vestments of the Atal'ai Prophet (10806, -3.32 DPS, sim-verified) [dungeon]; Ghostweave Vest (14141, -4.82 DPS) [crafted]; Feathered Breastplate (8349, -5.00 DPS) [crafted] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 133.0 healing_power points (9.05 DPS) | yes | Aristocratic Cuffs (12546, -1.27 DPS) [dungeon]; Mender's Leather Bracers (252543, -1.47 DPS) [crafted]; Venomshroud Armguards (14439, -1.66 DPS) [world_drop] |
| hands | Ghostweave Gloves (14142) | Tailoring [crafted] | sim-verified (343.2 DPS) | yes | Gloves of the Atal'ai Prophet (10808, +0.00 DPS) [dungeon]; Witherseed Gloves (16738, -1.73 DPS) [quest]; Mender's Leather Gauntlets (252551, -1.81 DPS) [crafted] |
| waist | Earthenweave Cord (254077) | Tailoring [crafted] | 160.2 healing_power points (10.91 DPS) | yes | Tooled Leather Belt (252544, -1.60 DPS) [crafted]; Windchaser Cinch (14435, -1.94 DPS) [world_drop]; Ghostweave Cord (254073, -2.32 DPS, sim-verified) [crafted] |
| legs | Dalewind Trousers (13008) | World drop [world_drop] | 329.7 healing_power points (22.45 DPS) | yes | Rainstrider Leggings (11123, -1.35 DPS, sim-verified) [quest]; Jinxed Hoodoo Kilt (9474, -3.71 DPS) [dungeon]; Kilt of the Atal'ai Prophet (10807, -3.73 DPS) [dungeon] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 265.9 healing_power points (18.11 DPS) | yes | Mistwalker Boots (10629, -1.65 DPS, sim-verified) [dungeon]; Furen's Boots (13100, -3.51 DPS) [world_drop]; Imperial Red Boots (8246, -5.64 DPS) [world_drop] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 162.9 healing_power points (11.09 DPS) | yes | Seedtime Hoop (9655, -2.78 DPS) [quest]; Coldwater Ring (4550, -3.05 DPS) [quest]; The Queen's Jewel (13094, -4.07 DPS) [world_drop] |
| finger2 | Choking Band (11868) | Ogre Head On A Stick = Party [quest] | 158.6 healing_power points (10.80 DPS) | yes | Seedtime Hoop (9655, -2.59 DPS, sim-verified) [quest]; Coldwater Ring (4550, -2.76 DPS) [quest]; The Queen's Jewel (13094, -3.78 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+4.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Thunderbrew's Boot Flask (744, -3.32 DPS) [quest]; Enriched Seal (271903, -6.19 DPS) [rep]; Alchemist's Stone (13503, -6.65 DPS) [crafted] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Thunderbrew's Boot Flask (744, -1.32 DPS, sim-verified) [quest]; Enriched Seal (271903, -4.53 DPS) [rep]; Alchemist's Stone (13503, -4.98 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Soulkeeper (1607, +0.00 DPS) [world_drop]; Dreamstaff (249454, -0.16 DPS) [crafted]; Spire of Hakkar (10844, -17.49 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Engineer's Guild Headpiece; neck: Lei of Lilies; shoulder: Living Shoulders; back: Arcane Cloak; chest: Embrace of the Wind Serpent; hands: Ghostweave Gloves; waist: Earthenweave Cord; legs: Dalewind Trousers; feet: Sandals of the Insurgent; finger1: Eye of Adaegus; finger2: Choking Band; trinket2: Evonice's Landin' Pilla; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 690, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60 (night-elf, 0530200100000000-00000000000000000000-5553005153113210)

Set DPS (verified): 483.0. Weights run: 8.0s. Verify run: 330.6s. 1755 eligible items had no known source.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.056, intellect=3.976 ± 0.155, spirit=12.514 ± 0.161, mp5=13.013 ± 0.178, crit=0.815 ± 0.029 per rating point (14 rating = 1%, 11.409 per %), spell_haste=not significant (-0.647 ± 1.505)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Feathermoon Headdress (13113, +0.00 DPS) [world_drop]; Crown of Caer Darrow (13986, +0.00 DPS) [quest]; Wildheart Cowl (16720, -15.19 DPS, sim-verified) [dungeon] |
| neck | Lady Alizabeth's Pendant (13002) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jeweled Amulet of Cainwyn (1443, +0.00 DPS) [world_drop]; Heart of the Fiend (13960, +0.00 DPS) [dungeon]; Lady Maye's Pendant (14558, +0.00 DPS) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Dragonhide Pauldrons (227201, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Pauldrons (231705, +0.00 DPS) [vendor]; Mantle of the Scarlet Crusade (22405, -16.37 DPS, sim-verified) [dungeon] |
| back | Cloak of Earth and Sky (279270) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Featherskin Cape (10843, +0.00 DPS) [world]; Frostweaver Cape (12968, +0.00 DPS) [dungeon]; Indomitable Cloak (14683, -0.87 DPS) [world_drop] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Embrace of the Wind Serpent (12462, -1.31 DPS, sim-verified) [world]; Vestments of the Atal'ai Prophet (10806, -2.60 DPS) [dungeon]; Ghostweave Vest (14141, -3.86 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wyrmthalak's Shackles (13958, +0.00 DPS) [quest]; Marshal's Dragonhide Bracers (16445, +0.00 DPS) [pvp]; Bleak Howler Armguards (13208, -13.36 DPS, sim-verified) [dungeon] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swarmtender's Gloves (275607, -0.08 DPS) [crafted]; Gloves of the Atal'ai Prophet (10808, -1.19 DPS) [dungeon]; Wildheart Gloves (16717, -14.69 DPS, sim-verified) [dungeon] |
| waist | Caretaker's Cord (272398) | Pix Xizzix [vendor] | 323.3 healing_power points (22.40 DPS) | yes | Elderwild Waistcord (279252, -3.65 DPS, sim-verified) [crafted]; Mana-infused Cord (279267, -6.79 DPS) [crafted]; Marshal's Dragonhide Waistguard (16447, -6.84 DPS) [pvp] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Haunting Specter Leggings (11929, +0.00 DPS) [dungeon]; Leggings of Arcana (12756, +0.00 DPS) [quest]; Earthborn Kilt (9402, -0.59 DPS) [world_drop] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wildheart Boots (16715, -0.28 DPS) [dungeon]; Mistwalker Boots (10629, -0.45 DPS) [dungeon]; Sandals of the Insurgent (13111, -17.39 DPS, sim-verified) [world_drop] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | sim-verified (483.0 DPS) | yes | The Postmaster's Seal (13392, +0.00 DPS) [dungeon]; Choking Band (11868, -0.79 DPS) [quest]; Channeler's Ring (272406, -1.26 DPS) [vendor] |
| finger2 | Band of the Hierophant (13096) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Postmaster's Seal (13392, +0.00 DPS) [dungeon]; Choking Band (11868, -1.02 DPS) [quest]; Channeler's Ring (272406, -1.50 DPS) [vendor] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+3.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Thunderbrew's Boot Flask (744, -3.47 DPS) [quest]; Shard of the Splithooves (10659, -4.99 DPS) [quest]; Enriched Seal (271903, -5.89 DPS) [rep] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Thunderbrew's Boot Flask (744, -1.73 DPS) [quest]; Serenity Field (272439, -2.05 DPS, sim-verified) [vendor]; Shard of the Splithooves (10659, -3.26 DPS) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Staff of Hale Magefire (13000, +0.00 DPS) [world_drop]; Dancing Sliver (15854, +0.00 DPS) [quest] |
| off_hand | Lapidis Tankard of Tidesippe (4696) | World drop [world_drop] | 220.1 healing_power points (15.25 DPS) | yes | Penelope's Rose (15805, -1.99 DPS, sim-verified) [quest]; Thaurissan's Royal Scepter (11928, -3.68 DPS) [dungeon]; Cloud Stone (17737, -3.82 DPS) [dungeon] |
| ranged | Idol of Swiftness (279250) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Idol of the Moon (23197, -1.36 DPS, sim-verified) [world_drop] |

**New at 60:** head: Devout Crown; neck: Lady Alizabeth's Pendant; shoulder: Devout Mantle; back: Cloak of Earth and Sky; chest: Ironfeather Breastplate; wrist: Devout Bracers; hands: Devout Gloves; waist: Caretaker's Cord; legs: Devout Skirt; feet: Devout Sandals; finger2: Band of the Hierophant; main_hand: Hand of Edward the Odd; off_hand: Lapidis Tankard of Tidesippe; ranged: Idol of Swiftness

No-known-source sample (15 of 1755, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60, raid preset (night-elf, 0530200100000000-00000000000000000000-5553005153113210)

Set DPS (verified): 1033.6. Weights run: 7.5s. Verify run: 302.0s. 1755 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.041, intellect=0.945 ± 0.037, spirit=2.066 ± 0.035, mp5=1.923 ± 0.040, crit=0.416 ± 0.012 per rating point (14 rating = 1%, 5.819 per %), spell_haste=not significant (0.357 ± 0.422)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 118.9 healing_power points (47.17 DPS) | yes | Field Marshal's Dragonhide Headdress (231701, -11.43 DPS) [vendor]; Lieutenant Commander's Dragonhide Headdress (227199, -11.89 DPS) [vendor]; Mooncloth Circlet (14140, -15.45 DPS, sim-verified) [crafted] |
| neck | Lady Alizabeth's Pendant (13002) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jeweled Amulet of Cainwyn (1443, +0.00 DPS) [world_drop]; Lady Maye's Pendant (14558, +0.00 DPS) [world_drop]; Animated Chain Necklace (18723, +0.00 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Dragonhide Pauldrons (227201, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Pauldrons (231705, +0.00 DPS) [vendor]; Mooncloth Shoulders (14139, -5.50 DPS, sim-verified) [crafted] |
| back | Caretaker's Cape (19530) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Drape of Recovery (272413, -1.52 DPS) [vendor]; Frostweaver Cape (12968, -2.54 DPS) [dungeon]; Featherskin Cape (10843, -3.08 DPS) [world] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Robes of the Exalted (13346, +0.00 DPS) [dungeon]; Field Marshal's Dragonhide Tunic (231702, +0.00 DPS) [vendor]; Ironfeather Breastplate (15066, -1.88 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dryad's Wrist Bindings (19595, +0.00 DPS) [rep]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted]; Bleak Howler Armguards (13208, -2.49 DPS, sim-verified) [dungeon] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wildheart Gloves (16717, +0.00 DPS) [dungeon]; Marshal's Dragonhide Gauntlets (231706, +0.00 DPS) [pvp]; Hands of the Exalted Herald (12554, -3.83 DPS, sim-verified) [dungeon] |
| waist | Caretaker's Cord (272398) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sash of Mercy (14553, +0.00 DPS) [world_drop]; Elderwild Waistcord (279252, -2.14 DPS) [crafted]; Wisdom of the Timbermaw (19047, -3.87 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Dragonhide Legguards (227200, +0.00 DPS) [vendor]; Elderwild Pants (279254, +0.00 DPS) [crafted]; Leggings of Arcana (12756, -1.15 DPS, sim-verified) [quest] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Faith Healer's Boots (22247, +0.00 DPS) [dungeon]; Marshal's Dragonhide Boots (231698, +0.00 DPS) [pvp]; Mooncloth Boots (15802, -3.85 DPS, sim-verified) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-verified (1033.7 DPS) | yes | Rosewine Circle (13178, +0.00 DPS) [dungeon]; Fordring's Seal (16058, +0.00 DPS) [quest]; Band of Mending (22334, +0.00 DPS) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, +0.00 DPS) [dungeon]; Fordring's Seal (16058, -1.41 DPS, sim-verified) [quest]; Rosewine Circle (13178, -3.53 DPS) [dungeon] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -6.82 DPS) [world_drop]; Second Wind (11819, -7.93 DPS) [dungeon]; Evonice's Landin' Pilla (18951, -16.74 DPS, sim-verified) [quest] |
| trinket2 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Second Wind (11819, +0.00 DPS) [dungeon]; Evonice's Landin' Pilla (18951, -1.51 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamstaff (249454, -21.33 DPS, sim-verified) [crafted]; Resurgence Rod (17743, -24.92 DPS) [quest]; Grand Marshal's Warhammer (23454, -30.89 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Idol of Swiftness (279250) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Idol of the Moon (23197, -9.17 DPS, sim-verified) [world_drop] |

**New at 60:** head: Living Crown; neck: Lady Alizabeth's Pendant; shoulder: Devout Mantle; back: Caretaker's Cape; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Caretaker's Cord; legs: Devout Skirt; feet: Devout Sandals; finger1: Maiden's Circle; finger2: Blessed Band of Light; trinket1: Serenity Field; trinket2: Dormant Heart of the Mountain; ranged: Idol of Swiftness

No-known-source sample (15 of 1755, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

## Horde

### Band 20 (tauren, 0000000000000000-00000000000000000000-2503000100000000)

Set DPS (verified): 75.5. Weights run: 6.1s. Verify run: 45.4s. 224 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.021, intellect=0.702 ± 0.003, spirit=1.744 ± 0.007, mp5=1.987 ± 0.009, crit=0.043 ± 0.001 per rating point (14 rating = 1%, 0.599 per %), spell_haste=not significant (-0.094 ± 0.026)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (4.15 DPS) | yes | Wisdom's Leather Hood (252507, -0.11 DPS, sim-verified) [crafted]; Stormrider's Leather Hood (252506, -0.53 DPS) [crafted]; Shadow Goggles (4373, -0.93 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 5.2 healing_power points (1.21 DPS) | yes | Scholarly Pendant (277203, -0.21 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Serpent's Shoulders (5404, -0.12 DPS, sim-verified) [dungeon]; Reinforced Woolen Shoulders (4315, -0.81 DPS) [crafted]; Forest Leather Mantle (4709, -1.00 DPS) [world_drop] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.07 DPS, sim-verified) [crafted]; Spirit Cloak (4792, -1.27 DPS) [vendor]; Antiquated Cloak (4799, -1.27 DPS) [vendor] |
| chest | Robe of the Moccasin (6465) | Wailing Caverns: Lord Cobrahn [dungeon] | 26.4 healing_power points (6.10 DPS) | yes | Corsair's Overshirt (5202, -0.21 DPS, sim-verified) [dungeon]; Scholarly Robes (2034, -0.58 DPS) [dungeon]; Wisdom's Leather Armor (252493, -0.62 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Drakewing Bands (12999, -0.24 DPS, sim-verified) [world_drop]; Crystalline Cuffs (14148, -0.98 DPS) [dungeon]; Seer's Cuffs (3645, -1.22 DPS) [dungeon] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 13.1 healing_power points (3.02 DPS) | yes | Wisdom's Leather Gloves (252499, -0.07 DPS) [crafted]; Magefist Gloves (12977, -0.60 DPS) [world_drop]; Bright Gloves (3066, -0.77 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 13.8 healing_power points (3.18 DPS) | yes | Novice Ardent's Sash (253887, -0.11 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.69 DPS) [crafted]; Stormrider's Leather Belt (252432, -1.61 DPS) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 31.2 healing_power points (7.19 DPS) | yes | Filigreed Pristine Leggings (253937, -0.21 DPS, sim-verified) [crafted]; Stormrider's Leather Pants (252502, -2.31 DPS) [crafted]; Darkweave Breeches (12987, -3.65 DPS) [world_drop] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 14.5 healing_power points (3.35 DPS) | yes | Red Woolen Boots (4313, -0.23 DPS, sim-verified) [crafted]; Black Whelp Slippers (252424, -0.62 DPS) [crafted]; Pristine Boots (253889, -0.79 DPS) [crafted] |
| finger1 | Band of Purification (12996) | World drop [world_drop] | 10.5 healing_power points (2.41 DPS) | yes | Lavishly Jeweled Ring (1156, -1.44 DPS) [dungeon]; Loop of Sacrifice (281673, -1.60 DPS) [quest]; Horn Ring (17692, -1.61 DPS) [quest] |
| finger2 | Advisor's Ring (20426) | Warsong Outriders [rep] | 9.0 healing_power points (2.07 DPS) | yes | Lavishly Jeweled Ring (1156, -0.47 DPS, sim-verified) [dungeon]; Loop of Sacrifice (281673, -1.26 DPS) [quest]; Horn Ring (17692, -1.27 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Living Root (6631) | Wailing Caverns: Verdan the Everliving [dungeon] | 53.7 healing_power points (12.39 DPS) | yes | Staff of the Blessed Seer (2271, -0.93 DPS, sim-verified) [dungeon]; Gnarled Necromancer's Staff (251534, -2.37 DPS) [quest]; Evocator's Blade (2567, -3.60 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Robe of the Moccasin; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Band of Purification; finger2: Advisor's Ring; main_hand: Living Root

No-known-source sample (15 of 224, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance

### Band 30 (tauren, 0000000000000000-00000000000000000000-5553001100010000)

Set DPS (verified): 120.5. Weights run: 8.4s. Verify run: 77.6s. 401 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.049, intellect=1.337 ± 0.019, spirit=2.833 ± 0.018, mp5=3.849 ± 0.023, crit=0.105 ± 0.005 per rating point (14 rating = 1%, 1.470 per %), spell_haste=not significant (-0.066 ± 0.171)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 50.0 healing_power points (7.91 DPS) | yes | Enduring Cap (3020, -2.08 DPS, sim-verified) [world_drop]; Nightsky Cowl (4039, -3.13 DPS) [world_drop]; Dokebi Hat (14584, -3.34 DPS) [world_drop] |
| neck | Pendant of Myzrael (4614) (or Glowing Green Talisman (5002)) | Razorfen Downs: Splinterbone Captain [dungeon] | 17.0 healing_power points (2.69 DPS) | yes | Glowing Green Talisman (5002, +0.00 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.05 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.50 DPS) [world_drop] |
| shoulder | Nightsky Mantle (4718) | World drop [world_drop] | 24.9 healing_power points (3.93 DPS) | yes | Ghostly Mantle (3324, -0.31 DPS, sim-verified) [quest]; Bloodmage Mantle (7684, -0.61 DPS) [dungeon]; Desert Shoulders (15457, -0.63 DPS) [quest] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | sim-verified (120.5 DPS) | yes | Battle Healer's Cloak (19529, -0.36 DPS, sim-verified) [rep]; Amy's Blanket (13005, -1.55 DPS) [world_drop]; Darkspear Raider's Cloak (272078, -1.66 DPS) [vendor] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 47.7 healing_power points (7.54 DPS) | yes | Wisdom's Leather Tunic (252511, -0.65 DPS, sim-verified) [crafted]; Pristine Gown (253961, -1.11 DPS) [crafted]; Pressed Felt Robe (1997, -1.53 DPS) [world] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 20.5 healing_power points (3.25 DPS) | yes | Drakewing Bands (12999, -0.56 DPS) [world_drop]; Mindthrust Bracers (1974, -0.63 DPS, sim-verified) [dungeon]; Nightsky Wristbands (6407, -0.64 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 40.7 healing_power points (6.43 DPS) | yes | Hotshot Pilot's Gloves (9491, -2.01 DPS, sim-verified) [dungeon]; Gloves of Meditation (4318, -2.51 DPS) [crafted]; Stonecloth Gloves (14411, -2.85 DPS) [world_drop] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 44.0 healing_power points (6.96 DPS) | yes | Skycaller's Leather Belt (252522, -2.13 DPS, sim-verified) [crafted]; Belt of Arugal (6392, -2.66 DPS) [dungeon]; Silver-lined Belt (13011, -3.27 DPS) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 50.2 healing_power points (7.94 DPS) | yes | Pristine Leggings (253987, -0.74 DPS) [crafted]; Earthen Leggings (253999, -0.83 DPS, sim-verified) [crafted]; Stormrider's Leather Kilt (252518, -1.53 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 40.7 healing_power points (6.43 DPS) | yes | Boots of the Enchanter (4325, -2.10 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -2.16 DPS) [dungeon]; Glinteye Slippers (273024, -2.64 DPS) [dungeon] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 28.8 healing_power points (4.56 DPS) | yes | The Queen's Jewel (13094, -0.55 DPS) [world_drop]; Solid Gold Band (279845, -0.58 DPS) [quest]; Ring of Power Regulation (273030, -1.78 DPS) [dungeon] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 28.8 healing_power points (4.56 DPS) | yes | Solid Gold Band (279845, -0.58 DPS) [quest]; The Queen's Jewel (13094, -0.77 DPS, sim-verified) [world_drop]; Ring of Power Regulation (273030, -1.78 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Friar (3415) | Blackfathom Deeps [dungeon] | sim-verified (120.5 DPS) | yes | Odo's Ley Staff (6318, -0.54 DPS, sim-verified) [dungeon]; Death Speaker Scepter (2816, -3.48 DPS) [dungeon]; Soulstaff (249393, -3.93 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Pendant of Myzrael; shoulder: Nightsky Mantle; back: Glowing Thresher Cape; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Staff of the Friar

No-known-source sample (15 of 401, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5821 Darkstalker Boots

### Band 40 (tauren, 0000000000000000-00000000000000000000-5553005151010000)

Set DPS (verified): 213.3. Weights run: 8.6s. Verify run: 80.7s. 525 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.170, intellect=1.132 ± 0.033, spirit=2.488 ± 0.023, mp5=3.698 ± 0.026, crit=0.252 ± 0.019 per rating point (14 rating = 1%, 3.534 per %), spell_haste=not significant (0.594 ± 0.341)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 63.6 healing_power points (13.37 DPS) | yes | Whitemane's Chapeau (7720, -0.69 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -3.30 DPS) [world_drop]; Earthen Silk Hood (254015, -3.52 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 29.1 healing_power points (6.10 DPS) | yes | Triune Amulet (7722, -0.63 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -0.78 DPS) [dungeon]; Amberglow Talisman (10824, -0.88 DPS) [quest] |
| shoulder | Earthen Silk Shoulders (254033) | Tailoring [crafted] | 37.9 healing_power points (7.96 DPS) | yes | Inquisitor's Shawl (19507, -0.25 DPS) [dungeon]; Mantle of Doan (7712, -0.88 DPS) [dungeon]; Sheepshear Mantle (13115, -1.02 DPS) [world_drop] |
| back | Cloak of Blight (6832) | Nothing But The Truth [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Thresher Cape (6901, -1.14 DPS) [dungeon]; Battle Healer's Cloak (19528, -1.17 DPS, sim-verified) [rep]; Stargazer Cloak (9660, -1.57 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 77.9 healing_power points (16.36 DPS) | yes | Dreamweave Vest (10021, -1.16 DPS, sim-verified) [crafted]; Doomsayer's Robe (4746, -3.60 DPS) [quest]; Civinad Robes (9623, -4.63 DPS) [quest] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 29.4 healing_power points (6.18 DPS) | yes | Earthen Silk Cuffs (254019, -1.92 DPS, sim-verified) [crafted]; Spidertank Oilrag (9448, -1.96 DPS) [dungeon]; Mindthrust Bracers (1974, -2.66 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Earthen Silk Gloves (254017, -0.65 DPS, sim-verified) [crafted]; Gloves of Old (9395, -0.93 DPS) [world_drop]; Windchaser Handguards (14431, -1.60 DPS) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 46.0 healing_power points (9.66 DPS) | yes | Mender's Leather Belt (252523, -1.02 DPS, sim-verified) [crafted]; Windchaser Cinch (14435, -3.05 DPS) [world_drop]; Earthen Silk Belt (7061, -3.80 DPS) [crafted] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 65.6 healing_power points (13.78 DPS) | yes | Stoneweaver Leggings (9407, -1.16 DPS, sim-verified) [dungeon]; Stormcloth Pants (10010, -2.79 DPS) [crafted]; Wisdom's Leather Leggings (252519, -4.05 DPS) [crafted] |
| feet | Furen's Boots (13100) | World drop [world_drop] | sim-verified (213.3 DPS) | yes | Thoughtcast Boots (10578, -0.50 DPS) [dungeon]; Mender's Leather Shoes (252533, -0.66 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -2.07 DPS) [crafted] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 26.4 healing_power points (5.55 DPS) | yes | Solid Gold Band (279845, -0.63 DPS) [quest]; Welken Ring (5011, -0.70 DPS) [world_drop]; Darkspear Signet (272070, -0.89 DPS) [vendor] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 26.4 healing_power points (5.55 DPS) | yes | Solid Gold Band (279845, -0.63 DPS, sim-verified) [quest]; Welken Ring (5011, -0.70 DPS) [world_drop]; Darkspear Signet (272070, -0.89 DPS) [vendor] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+9.2 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of the Friar (3415, -6.02 DPS) [dungeon]; Staff of Jordan (873, -6.16 DPS, sim-verified) [world_drop]; Odo's Ley Staff (6318, -6.73 DPS) [dungeon] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | 32.0 healing_power points (6.71 DPS) | yes | Orb of Souls (249395, -0.39 DPS, sim-verified) [crafted]; Silksand Star (15964, -2.01 DPS) [world_drop]; Omega Orb (7749, -2.49 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Earthen Silk Shoulders; back: Cloak of Blight; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Dreamweave Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Furen's Boots; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Hand of Righteousness; off_hand: Beacon of Hope

No-known-source sample (15 of 525, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (tauren, 0100000000000000-00000000000000000000-5553005153113210)

Set DPS (verified): 344.2. Weights run: 7.1s. Verify run: 83.4s. 689 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.129, intellect=1.972 ± 0.187, spirit=4.317 ± 0.132, mp5=4.990 ± 0.190, crit=0.735 ± 0.092 per rating point (14 rating = 1%, 10.290 per %), spell_haste=not significant (1.811 ± 3.415)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Engineer's Guild Headpiece (9534) | Divino-matic Rod [quest] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Gemburst Circlet (10751, -0.87 DPS, sim-verified) [quest]; Blood Guard's Restored Leather Helm (220875, -1.18 DPS) [vendor]; Serpentskin Helm (8261, -2.53 DPS) [world_drop] |
| neck | Lei of Lilies (1315) | World drop [world_drop] | 64.8 healing_power points (9.27 DPS) | yes | Jarkal's Enhancing Necklace (7888, -2.81 DPS) [quest]; Horizon Choker (13085, -2.84 DPS) [world_drop]; Gemshard Heart (17707, -4.79 DPS, sim-verified) [dungeon] |
| shoulder | Living Shoulders (15061) | Leatherworking [crafted] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Blood Guard's Restored Leather Spaulders (220877, -0.32 DPS) [vendor]; Mender's Leather Shoulders (252538, -2.17 DPS) [crafted]; Ironfeather Shoulders (15067, -2.18 DPS, sim-verified) [crafted] |
| back | Cloak of Blight (6832) | Nothing But The Truth [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Featherskin Cape (10843, +0.00 DPS) [world]; Wingveil Cloak (10802, -0.52 DPS) [dungeon]; Arcane Cloak (8286, -0.62 DPS) [world_drop] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 163.0 healing_power points (23.33 DPS) | yes | Feathered Breastplate (8349, -2.82 DPS, sim-verified) [crafted]; Vestments of the Atal'ai Prophet (10806, -3.55 DPS) [dungeon]; Ghostweave Vest (14141, -3.97 DPS) [crafted] |
| wrist | Mender's Leather Bracers (252543) | Leatherworking [crafted] | 58.7 healing_power points (8.40 DPS) | yes | Nethergeld Cuffs (254061, -0.47 DPS) [crafted]; Enchanted Kodo Bracers (13119, -1.09 DPS) [world_drop]; Aristocratic Cuffs (12546, -1.79 DPS, sim-verified) [dungeon] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of the Atal'ai Prophet (10808, -0.18 DPS) [dungeon]; Gilded Gloves (254095, -0.62 DPS) [crafted]; Ghostweave Gloves (14142, -1.98 DPS) [crafted] |
| waist | Earthenweave Cord (254077) | Tailoring [crafted] | 73.5 healing_power points (10.52 DPS) | yes | Gilded Cord (254037, -1.40 DPS) [crafted]; Ghostweave Cord (254073, -1.72 DPS) [crafted]; Tooled Leather Belt (252544, -3.30 DPS, sim-verified) [crafted] |
| legs | Dalewind Trousers (13008) | World drop [world_drop] | 125.7 healing_power points (17.98 DPS) | yes | Jinxed Hoodoo Kilt (9474, -1.13 DPS) [dungeon]; Rainstrider Leggings (11123, -1.30 DPS) [quest]; Kilt of the Atal'ai Prophet (10807, -2.84 DPS, sim-verified) [dungeon] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 102.1 healing_power points (14.61 DPS) | yes | Mistwalker Boots (10629, -1.64 DPS, sim-verified) [dungeon]; First Sergeant's Restored Leather Boots (220885, -2.75 DPS) [vendor]; Furen's Boots (13100, -2.75 DPS) [world_drop] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 63.6 healing_power points (9.11 DPS) | yes | Welken Ring (5011, -3.37 DPS) [world_drop]; Cyclopean Band (11824, -3.37 DPS) [dungeon]; Brainlash (6440, -10.05 DPS, sim-verified) [dungeon] |
| finger2 | Coldwater Ring (4550) | Death From Below [quest] | sim-verified (344.2 DPS) | yes | Brainlash (6440, +0.00 DPS) [dungeon]; Welken Ring (5011, -0.67 DPS) [world_drop]; Cyclopean Band (11824, -0.67 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+5.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Darkspear Voodoo Seal (272061, -3.84 DPS) [vendor]; Uther's Strength (11302, -4.84 DPS) [world_drop]; Alchemist's Stone (13503, -4.94 DPS) [crafted] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-verified (+3.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Darkspear Voodoo Seal (272061, -3.14 DPS, sim-verified) [vendor]; Uther's Strength (11302, -3.60 DPS) [world_drop]; Alchemist's Stone (13503, -3.71 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamstaff (249454, -5.61 DPS, sim-verified) [crafted]; Resurgence Rod (17743, -11.33 DPS) [quest]; Charstone Dirk (17710, -14.08 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Engineer's Guild Headpiece; neck: Lei of Lilies; shoulder: Living Shoulders; chest: Embrace of the Wind Serpent; wrist: Mender's Leather Bracers; hands: Mender's Leather Gauntlets; waist: Earthenweave Cord; legs: Dalewind Trousers; feet: Sandals of the Insurgent; finger1: Eye of Adaegus; finger2: Coldwater Ring; trinket2: Evonice's Landin' Pilla; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 689, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (tauren, 0530200100000000-00000000000000000000-5553005153113210)

Set DPS (verified): 481.9. Weights run: 8.0s. Verify run: 330.2s. 1759 eligible items had no known source.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.056, intellect=3.976 ± 0.155, spirit=12.514 ± 0.161, mp5=13.013 ± 0.178, crit=0.815 ± 0.029 per rating point (14 rating = 1%, 11.409 per %), spell_haste=not significant (-0.647 ± 1.505)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Feathermoon Headdress (13113, +0.00 DPS) [world_drop]; Crown of Caer Darrow (13986, +0.00 DPS) [quest]; Wildheart Cowl (16720, -15.37 DPS, sim-verified) [dungeon] |
| neck | Lady Alizabeth's Pendant (13002) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jeweled Amulet of Cainwyn (1443, +0.00 DPS) [world_drop]; Heart of the Fiend (13960, +0.00 DPS) [dungeon]; Lady Maye's Pendant (14558, +0.00 DPS) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Dragonhide Pauldrons (227207, +0.00 DPS) [vendor]; Warlord's Dragonhide Pauldrons (231672, +0.00 DPS) [vendor]; Mantle of the Scarlet Crusade (22405, -16.62 DPS, sim-verified) [dungeon] |
| back | Cloak of Earth and Sky (279270) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Featherskin Cape (10843, +0.00 DPS) [world]; Frostweaver Cape (12968, +0.00 DPS) [dungeon]; Cloak of Blight (6832, -0.87 DPS) [quest] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Embrace of the Wind Serpent (12462, +0.00 DPS) [world]; Vestments of the Atal'ai Prophet (10806, -2.60 DPS) [dungeon]; Ghostweave Vest (14141, -3.86 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wyrmthalak's Shackles (13958, +0.00 DPS) [quest]; General's Dragonhide Bracers (16553, +0.00 DPS) [pvp]; Bleak Howler Armguards (13208, -13.55 DPS, sim-verified) [dungeon] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swarmtender's Gloves (275607, -0.08 DPS) [crafted]; Gloves of the Atal'ai Prophet (10808, -1.19 DPS) [dungeon]; Wildheart Gloves (16717, -15.28 DPS, sim-verified) [dungeon] |
| waist | Caretaker's Cord (272398) | Pix Xizzix [vendor] | 323.3 healing_power points (22.40 DPS) | yes | Elderwild Waistcord (279252, -2.39 DPS, sim-verified) [crafted]; Mana-infused Cord (279267, -6.79 DPS) [crafted]; General's Dragonhide Belt (16556, -6.84 DPS) [pvp] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Haunting Specter Leggings (11929, +0.00 DPS) [dungeon]; Leggings of Arcana (12756, +0.00 DPS) [quest]; Earthborn Kilt (9402, -0.59 DPS) [world_drop] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wildheart Boots (16715, -0.28 DPS) [dungeon]; Mistwalker Boots (10629, -0.45 DPS) [dungeon]; Sandals of the Insurgent (13111, -17.40 DPS, sim-verified) [world_drop] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | sim-verified (481.9 DPS) | yes | The Postmaster's Seal (13392, +0.00 DPS) [dungeon]; Channeler's Ring (272406, -1.26 DPS) [vendor]; Seal of Rivendare (13345, -1.30 DPS) [dungeon] |
| finger2 | Band of the Hierophant (13096) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Postmaster's Seal (13392, +0.00 DPS) [dungeon]; Channeler's Ring (272406, -1.50 DPS) [vendor]; Seal of Rivendare (13345, -1.54 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+11.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Shard of the Splithooves (10659, -4.99 DPS) [quest]; Darkspear Voodoo Seal (272061, -5.89 DPS) [vendor]; Alchemist's Stone (13503, -6.93 DPS) [crafted] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, +0.00 DPS) [vendor]; Shard of the Splithooves (10659, -3.26 DPS) [quest]; Darkspear Voodoo Seal (272061, -4.16 DPS) [vendor] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Staff of Hale Magefire (13000, +0.00 DPS) [world_drop]; Dancing Sliver (15854, +0.00 DPS) [quest] |
| off_hand | Lapidis Tankard of Tidesippe (4696) | World drop [world_drop] | 220.1 healing_power points (15.25 DPS) | yes | Penelope's Rose (15805, -2.01 DPS, sim-verified) [quest]; Thaurissan's Royal Scepter (11928, -3.68 DPS) [dungeon]; Cloud Stone (17737, -3.82 DPS) [dungeon] |
| ranged | Idol of Swiftness (279250) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Idol of the Moon (23197, -1.08 DPS, sim-verified) [world_drop] |

**New at 60:** head: Devout Crown; neck: Lady Alizabeth's Pendant; shoulder: Devout Mantle; back: Cloak of Earth and Sky; chest: Ironfeather Breastplate; wrist: Devout Bracers; hands: Devout Gloves; waist: Caretaker's Cord; legs: Devout Skirt; feet: Devout Sandals; finger2: Band of the Hierophant; main_hand: Hand of Edward the Odd; off_hand: Lapidis Tankard of Tidesippe; ranged: Idol of Swiftness

No-known-source sample (15 of 1759, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (tauren, 0530200100000000-00000000000000000000-5553005153113210)

Set DPS (verified): 1030.9. Weights run: 7.5s. Verify run: 299.6s. 1759 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.041, intellect=0.945 ± 0.037, spirit=2.066 ± 0.035, mp5=1.923 ± 0.040, crit=0.416 ± 0.012 per rating point (14 rating = 1%, 5.819 per %), spell_haste=not significant (0.357 ± 0.422)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 118.9 healing_power points (47.17 DPS) | yes | Warlord's Dragonhide Headdress (231675, -11.43 DPS) [vendor]; Champion's Dragonhide Headdress (227205, -11.89 DPS) [vendor]; Mooncloth Circlet (14140, -15.23 DPS, sim-verified) [crafted] |
| neck | Lady Alizabeth's Pendant (13002) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jeweled Amulet of Cainwyn (1443, +0.00 DPS) [world_drop]; Lady Maye's Pendant (14558, +0.00 DPS) [world_drop]; Animated Chain Necklace (18723, +0.00 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Dragonhide Pauldrons (227207, +0.00 DPS) [vendor]; Warlord's Dragonhide Pauldrons (231672, +0.00 DPS) [vendor]; Mooncloth Shoulders (14139, -5.79 DPS, sim-verified) [crafted] |
| back | Battle Healer's Cloak (19526) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Raincaster Drape (12110, -1.45 DPS, sim-verified) [quest]; Drape of Recovery (272413, -1.52 DPS) [vendor]; Frostweaver Cape (12968, -2.54 DPS) [dungeon] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Robes of the Exalted (13346, +0.00 DPS) [dungeon]; Warlord's Dragonhide Tunic (231674, +0.00 DPS) [vendor]; Ironfeather Breastplate (15066, -3.32 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dryad's Wrist Bindings (19595, +0.00 DPS) [rep]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted]; Bleak Howler Armguards (13208, -2.74 DPS, sim-verified) [dungeon] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wildheart Gloves (16717, +0.00 DPS) [dungeon]; General's Dragonhide Gauntlets (231676, +0.00 DPS) [vendor]; Hands of the Exalted Herald (12554, -4.09 DPS, sim-verified) [dungeon] |
| waist | Caretaker's Cord (272398) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sash of Mercy (14553, +0.00 DPS) [world_drop]; Elderwild Waistcord (279252, -2.14 DPS) [crafted]; Wisdom of the Timbermaw (19047, -3.87 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Dragonhide Legguards (227206, +0.00 DPS) [vendor]; Elderwild Pants (279254, +0.00 DPS) [crafted]; Leggings of Arcana (12756, -1.89 DPS, sim-verified) [quest] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Faith Healer's Boots (22247, +0.00 DPS) [dungeon]; General's Dragonhide Boots (231682, +0.00 DPS) [pvp]; Mooncloth Boots (15802, -4.23 DPS, sim-verified) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-verified (1030.9 DPS) | yes | Rosewine Circle (13178, +0.00 DPS) [dungeon]; Fordring's Seal (16058, +0.00 DPS) [quest]; Band of Mending (22334, +0.00 DPS) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, +0.00 DPS) [dungeon]; Fordring's Seal (16058, -1.42 DPS, sim-verified) [quest]; Rosewine Circle (13178, -3.53 DPS) [dungeon] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -6.82 DPS) [world_drop]; Second Wind (11819, -7.93 DPS) [dungeon]; Evonice's Landin' Pilla (18951, -15.98 DPS, sim-verified) [quest] |
| trinket2 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (+2.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Second Wind (11819, +0.00 DPS) [dungeon]; Evonice's Landin' Pilla (18951, -2.58 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamstaff (249454, -22.52 DPS, sim-verified) [crafted]; Resurgence Rod (17743, -24.92 DPS) [quest]; High Warlord's Battle Mace (23464, -30.89 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Idol of Swiftness (279250) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Idol of the Moon (23197, -10.17 DPS, sim-verified) [world_drop] |

**New at 60:** head: Living Crown; neck: Lady Alizabeth's Pendant; shoulder: Devout Mantle; back: Battle Healer's Cloak; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Caretaker's Cord; legs: Devout Skirt; feet: Devout Sandals; finger1: Maiden's Circle; finger2: Blessed Band of Light; trinket1: Serenity Field; trinket2: Dormant Heart of the Mountain; ranged: Idol of Swiftness

No-known-source sample (15 of 1759, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

