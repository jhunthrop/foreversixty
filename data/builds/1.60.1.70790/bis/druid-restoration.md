# Leveling BiS: Restoration

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-00000000000000000000-2503000100000000)

Set DPS (verified): 75.9. Weights run: 14.0s. Verify run: 45.1s. 251 eligible items had no known source.

Race sweep: high-order-skyborne 75.9 ± 0.0, night-elf 75.8 ± 0.0 (heal_score); high-order-skyborne beats the guide race night-elf by 0.0 (0.01%) ± 0.0, inside error.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.029, intellect=0.665 ± 0.003, spirit=1.680 ± 0.010, mp5=1.940 ± 0.010, crit=0.043 ± 0.002 per rating point (14 rating = 1%, 0.603 per %), spell_haste=not significant (-0.047 ± 0.034)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Wisdom's Leather Hood (252507) | Leatherworking [crafted] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Circlet (253949, -0.04 DPS, sim-verified) [crafted]; Stormrider's Leather Hood (252506, -0.48 DPS) [crafted]; Shadow Goggles (4373, -0.92 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 3.4 healing_power points (0.86 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 11.8 healing_power points (3.02 DPS) | yes | Magician's Mantle (12998, -0.20 DPS) [world_drop]; Reinforced Woolen Shoulders (4315, -1.05 DPS) [crafted]; Forest Leather Mantle (4709, -1.29 DPS) [world_drop] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | sim-verified (75.8 DPS) | yes | Heavy Woolen Cloak (4311, -0.42 DPS) [crafted]; Spirit Cloak (4792, -1.45 DPS) [vendor]; Sylvan Cloak (4793, -1.45 DPS) [vendor] |
| chest | Robe of the Moccasin (6465) | Wailing Caverns: Lord Cobrahn [dungeon] | 25.8 healing_power points (6.62 DPS) | yes | Corsair's Overshirt (5202, -0.35 DPS) [dungeon]; Wisdom's Leather Armor (252493, -0.62 DPS) [crafted]; Scholarly Robes (2034, -0.68 DPS) [dungeon] |
| wrist | Drakewing Bands (12999) | World drop [world_drop] | 10.1 healing_power points (2.59 DPS) | yes | Mindthrust Bracers (1974, -0.24 DPS) [dungeon]; Repurposed Hair Band (281256, -1.38 DPS) [quest]; Seer's Cuffs (3645, -1.55 DPS) [dungeon] |
| hands | Wisdom's Leather Gloves (252499) | Leatherworking [crafted] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Gloves (253913, -0.04 DPS, sim-verified) [crafted]; Magefist Gloves (12977, -0.67 DPS) [world_drop]; Bright Gloves (3066, -0.84 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 13.7 healing_power points (3.50 DPS) | yes | Novice Ardent's Sash (253887, -0.16 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.77 DPS) [crafted]; Stormrider's Leather Belt (252432, -1.80 DPS) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 30.7 healing_power points (7.88 DPS) | yes | Filigreed Pristine Leggings (253937, -0.51 DPS) [crafted]; Stormrider's Leather Pants (252502, -2.56 DPS) [crafted]; Darkweave Breeches (12987, -4.10 DPS) [world_drop] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 14.3 healing_power points (3.67 DPS) | yes | Red Woolen Boots (4313, -0.49 DPS) [crafted]; Black Whelp Slippers (252424, -0.68 DPS) [crafted]; Pristine Boots (253889, -0.85 DPS) [crafted] |
| finger1 | Band of Purification (12996) | World drop [world_drop] | 10.1 healing_power points (2.59 DPS) | yes | Lorekeeper's Ring (20431, -0.13 DPS, sim-verified) [rep]; Lavishly Jeweled Ring (1156, -1.56 DPS) [dungeon]; Demon Band (12054, -1.72 DPS) [world_drop] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Lorekeeper's Ring (20431, -0.04 DPS, sim-verified) [rep]; Lavishly Jeweled Ring (1156, -0.60 DPS) [dungeon]; Demon Band (12054, -0.76 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Westfall (2042) | The Defias Brotherhood [quest] | 61.4 healing_power points (15.75 DPS) | yes | Living Root (6631, -0.09 DPS, sim-verified) [dungeon]; Staff of the Blessed Seer (2271, -4.26 DPS) [dungeon]; Evocator's Blade (2567, -6.07 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Wisdom's Leather Hood; neck: Scholarly Pendant; shoulder: Serpent's Shoulders; back: Caretaker's Cape; chest: Robe of the Moccasin; wrist: Drakewing Bands; hands: Wisdom's Leather Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Band of Purification; finger2: Minor Channeling Ring; main_hand: Staff of Westfall

No-known-source sample (15 of 251, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5968 Rugged Boots; 6360 Steelscale Crushfish; 6478 Rat Stompers

### Band 30 (night-elf, 0100000000000000-00000000000000000000-3503004103100000)

Set DPS (verified): 132.2. Weights run: 8.7s. Verify run: 81.8s. 424 eligible items had no known source.

Race sweep: high-order-skyborne 134.7 ± 0.0, night-elf 132.3 ± 0.1 (heal_score); high-order-skyborne beats the guide race night-elf by 2.4 (1.82%) ± 0.1, beyond error.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.147, intellect=1.000 ± 0.014, spirit=1.638 ± 0.015, mp5=2.523 ± 0.024, crit=0.169 ± 0.015 per rating point (14 rating = 1%, 2.372 per %), spell_haste=not significant (0.296 ± 0.650)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 42.8 healing_power points (9.16 DPS) | yes | Enduring Cap (3020, -2.10 DPS, sim-verified) [world_drop]; Filigreed Pristine Circlet (253975, -4.45 DPS) [crafted]; Wisdom's Leather Helm (252515, -4.67 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 10.6 healing_power points (2.26 DPS) | yes | Pendant of Myzrael (4614, -0.15 DPS) [dungeon]; Glowing Green Talisman (5002, -0.15 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.35 DPS) [world_drop] |
| shoulder | Mantle of Honor (3560) | Bride of the Embalmer [quest] | 18.5 healing_power points (3.95 DPS) | yes | Death Speaker Mantle (6685, -0.31 DPS) [dungeon]; Nightsky Mantle (4718, -0.49 DPS) [world_drop]; Bloodmage Mantle (7684, -0.58 DPS, sim-verified) [dungeon] |
| back | Prelacy Cape (7004) | Researching the Corruption [quest] | sim-verified (132.3 DPS) | yes | Glowing Thresher Cape (6901, -0.15 DPS) [dungeon]; Caretaker's Cape (19533, -0.27 DPS) [rep]; Repairman's Cape (9605, -0.78 DPS) [quest] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 34.1 healing_power points (7.29 DPS) | yes | Pristine Gown (253961, -0.12 DPS) [crafted]; Death Speaker Robes (6682, -0.66 DPS) [dungeon]; Wisdom's Leather Tunic (252511, -0.68 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 16.6 healing_power points (3.54 DPS) | yes | Mindthrust Bracers (1974, -0.48 DPS, sim-verified) [dungeon]; Nightsky Wristbands (6407, -1.21 DPS) [world_drop]; Drakewing Bands (12999, -1.44 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 24.7 healing_power points (5.27 DPS) | yes | Shilly Mitts (9609, -0.95 DPS, sim-verified) [quest]; Truefaith Gloves (7049, -1.42 DPS) [crafted]; Gloves of Meditation (4318, -1.75 DPS) [crafted] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 34.8 healing_power points (7.45 DPS) | yes | Skycaller's Leather Belt (252522, -2.32 DPS, sim-verified) [crafted]; Belt of Arugal (6392, -3.13 DPS) [dungeon]; Silver-lined Belt (13011, -4.12 DPS) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 39.5 healing_power points (8.44 DPS) | yes | Pristine Leggings (253987, -0.62 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -1.38 DPS) [crafted]; Wisdom's Leather Pants (252503, -1.48 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 33.6 healing_power points (7.17 DPS) | yes | Acidic Walkers (9454, -2.36 DPS, sim-verified) [dungeon]; Boots of the Enchanter (4325, -3.30 DPS) [crafted]; Glinteye Slippers (273024, -3.58 DPS) [dungeon] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 26.5 healing_power points (5.66 DPS) | yes | Electrocutioner Lagnut (9447, -1.28 DPS) [dungeon]; The Queen's Jewel (13094, -2.43 DPS) [world_drop]; Ring of Power Regulation (273030, -2.76 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 24.5 healing_power points (5.23 DPS) | yes | Electrocutioner Lagnut (9447, -1.24 DPS, sim-verified) [dungeon]; The Queen's Jewel (13094, -2.00 DPS) [world_drop]; Ring of Power Regulation (273030, -2.33 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Friar (3415) | Blackfathom Deeps [dungeon] | sim-verified (132.3 DPS) | yes | Odo's Ley Staff (6318, -0.65 DPS, sim-verified) [dungeon]; Death Speaker Scepter (2816, -1.42 DPS) [dungeon]; Staff of Westfall (2042, -2.53 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Mantle of Honor; back: Prelacy Cape; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Runemender's Seal; finger2: Snake Hoop; main_hand: Staff of the Friar

No-known-source sample (15 of 424, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 40 (night-elf, 0500200000000000-00000000000000000000-3503005103103000)

Set DPS (verified): 243.7. Weights run: 8.8s. Verify run: 83.2s. 607 eligible items had no known source.

Race sweep: high-order-skyborne 247.2 ± 0.1, night-elf 243.8 ± 0.1 (heal_score); high-order-skyborne beats the guide race night-elf by 3.3 (1.37%) ± 0.2, beyond error.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.218, intellect=1.391 ± 0.027, spirit=2.707 ± 0.037, mp5=3.469 ± 0.027, crit=0.345 ± 0.045 per rating point (14 rating = 1%, 4.832 per %), spell_haste=not significant (0.261 ± 0.455)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 70.0 healing_power points (15.11 DPS) | yes | Whitemane's Chapeau (7720, -0.47 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -4.48 DPS) [world_drop]; Earthen Silk Hood (254015, -4.52 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 31.1 healing_power points (6.72 DPS) | yes | Necklace of Calisea (1714, -0.32 DPS, sim-verified) [dungeon]; Amberglow Talisman (10824, -0.87 DPS) [quest]; Triune Amulet (7722, -0.90 DPS) [dungeon] |
| shoulder | Earthen Silk Shoulders (254033) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Sheepshear Mantle (13115, -0.25 DPS) [world_drop]; Inquisitor's Shawl (19507, -0.54 DPS, sim-verified) [dungeon]; Mantle of Doan (7712, -0.57 DPS) [dungeon] |
| back | Caretaker's Cape (19532) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Thresher Cape (6901, -0.48 DPS, sim-verified) [dungeon]; Prelacy Cape (7004, -0.93 DPS) [quest]; Ceremonial Centaur Blanket (6789, -0.95 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 79.7 healing_power points (17.19 DPS) | yes | Dreamweave Vest (10021, -2.78 DPS, sim-verified) [crafted]; Doomsayer's Robe (4746, -3.22 DPS) [quest]; Death's Head Vestment (10581, -3.31 DPS) [dungeon] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 32.6 healing_power points (7.04 DPS) | yes | Earthen Silk Cuffs (254019, -2.12 DPS, sim-verified) [crafted]; Spidertank Oilrag (9448, -2.86 DPS) [dungeon]; Silkstream Cuffs (16791, -3.21 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Earthen Silk Gloves (254017, -0.35 DPS, sim-verified) [crafted]; Gloves of Old (9395, -0.66 DPS) [world_drop]; Warden's Gloves (14606, -1.55 DPS) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 49.4 healing_power points (10.65 DPS) | yes | Mender's Leather Belt (252523, -1.02 DPS, sim-verified) [crafted]; Windchaser Cinch (14435, -2.99 DPS) [world_drop]; Sutarn's Ring (13105, -4.15 DPS) [world_drop] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 70.2 healing_power points (15.16 DPS) | yes | Stoneweaver Leggings (9407, -1.60 DPS, sim-verified) [dungeon]; Stormcloth Pants (10010, -3.16 DPS) [crafted]; Wisdom's Leather Leggings (252519, -4.44 DPS) [crafted] |
| feet | Furen's Boots (13100) | World drop [world_drop] | 53.1 healing_power points (11.45 DPS) | yes | Mender's Leather Shoes (252533, -0.62 DPS, sim-verified) [crafted]; Thoughtcast Boots (10578, -1.05 DPS) [dungeon]; Gilded Slippers (254001, -2.70 DPS) [crafted] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 36.3 healing_power points (7.83 DPS) | yes | Electrocutioner Lagnut (9447, -1.80 DPS) [dungeon]; Welken Ring (5011, -2.24 DPS) [world_drop]; The Queen's Jewel (13094, -2.56 DPS) [world_drop] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 34.7 healing_power points (7.49 DPS) | yes | Electrocutioner Lagnut (9447, -1.35 DPS, sim-verified) [dungeon]; Welken Ring (5011, -1.90 DPS) [world_drop]; The Queen's Jewel (13094, -2.21 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+10.3 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of the Friar (3415, -5.83 DPS) [dungeon]; Odo's Ley Staff (6318, -6.73 DPS) [dungeon]; Staff of Jordan (873, -8.51 DPS, sim-verified) [world_drop] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | 32.8 healing_power points (7.08 DPS) | yes | Orb of Souls (249395, -0.77 DPS) [crafted]; Orb of Lorica (11262, -0.78 DPS, sim-verified) [quest]; Silksand Star (15964, -1.83 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Earthen Silk Shoulders; back: Caretaker's Cape; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Dreamweave Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Furen's Boots; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Hand of Righteousness; off_hand: Beacon of Hope

No-known-source sample (15 of 607, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 50 (night-elf, 0530200100000000-00000000000000000000-5503005113103210)

Set DPS (verified): 377.0. Weights run: 7.4s. Verify run: 76.0s. 815 eligible items had no known source.

Race sweep: high-order-skyborne 385.6 ± 0.2, night-elf 377.3 ± 0.3 (heal_score); high-order-skyborne beats the guide race night-elf by 8.3 (2.20%) ± 0.4, beyond error.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.105, intellect=2.018 ± 0.081, spirit=4.090 ± 0.093, mp5=5.075 ± 0.099, crit=0.624 ± 0.035 per rating point (14 rating = 1%, 8.739 per %), spell_haste=not significant (0.284 ± 1.944)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Engineer's Guild Headpiece (9534) | Divino-matic Rod [quest] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Gemburst Circlet (10751, -1.70 DPS, sim-verified) [quest]; Knight-Lieutenant's Restored Leather Helm (220874, -2.99 DPS) [vendor]; Serpentskin Helm (8261, -4.87 DPS) [world_drop] |
| neck | Lei of Lilies (1315) | World drop [world_drop] | 61.3 healing_power points (11.78 DPS) | yes | Necklace of Calisea (1714, -3.57 DPS) [dungeon]; Amberglow Talisman (10824, -3.93 DPS) [quest]; Horizon Choker (13085, -4.01 DPS, sim-verified) [world_drop] |
| shoulder | Living Shoulders (15061) | Leatherworking [crafted] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Restored Leather Spaulders (220876, -0.21 DPS) [vendor]; Ironfeather Shoulders (15067, -1.49 DPS, sim-verified) [crafted]; Mender's Leather Shoulders (252538, -2.57 DPS) [crafted] |
| back | Featherskin Cape (10843) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wingveil Cloak (10802, -3.32 DPS, sim-verified) [dungeon]; Spritecaster Cape (11623, -3.78 DPS) [dungeon]; Arcane Cloak (8286, -4.08 DPS) [world_drop] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 175.9 healing_power points (33.78 DPS) | yes | Feathered Breastplate (8349, -7.21 DPS, sim-verified) [crafted]; Ghostweave Vest (14141, -8.54 DPS) [crafted]; Forest's Embrace (22272, -9.71 DPS) [quest] |
| wrist | Mender's Leather Bracers (252543) | Leatherworking [crafted] | 57.7 healing_power points (11.08 DPS) | yes | Nethergeld Cuffs (254061, -0.59 DPS) [crafted]; Enchanted Kodo Bracers (13119, -1.67 DPS) [world_drop]; Skycaller's Leather Bracers (252542, -2.51 DPS) [crafted] |
| hands | Gloves of the Atal'ai Prophet (10808) | Sunken Temple: Jammal'an the Prophet [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mender's Leather Gauntlets (252551, -1.78 DPS, sim-verified) [crafted]; Gilded Gloves (254095, -2.48 DPS) [crafted]; Ogreseer Fists (11665, -3.98 DPS) [dungeon] |
| waist | Earthenweave Cord (254077) | Tailoring [crafted] | 71.0 healing_power points (13.63 DPS) | yes | Tooled Leather Belt (252544, -1.58 DPS, sim-verified) [crafted]; Gilded Cord (254037, -1.59 DPS) [crafted]; Ghostweave Cord (254073, -2.30 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 124.7 healing_power points (23.96 DPS) | yes | Rainstrider Leggings (11123, -0.08 DPS) [quest]; Jinxed Hoodoo Kilt (9474, -0.55 DPS) [dungeon]; Dalewind Trousers (13008, -0.83 DPS) [world_drop] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 97.9 healing_power points (18.81 DPS) | yes | Mistwalker Boots (10629, -2.19 DPS, sim-verified) [dungeon]; Sergeant Major's Restored Leather Boots (220884, -3.24 DPS) [vendor]; Furen's Boots (13100, -3.53 DPS) [world_drop] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 61.2 healing_power points (11.75 DPS) | yes | Brainlash (6440, -1.84 DPS) [dungeon]; Runemender's Seal (273031, -2.20 DPS) [world]; Snake Hoop (6750, -2.39 DPS) [quest] |
| finger2 | Choking Band (11868) | Ogre Head On A Stick = Party [quest] | 53.2 healing_power points (10.21 DPS) | yes | Brainlash (6440, -0.29 DPS) [dungeon]; Runemender's Seal (273031, -0.66 DPS) [world]; Snake Hoop (6750, -0.85 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+3.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Thunderbrew's Boot Flask (744, -3.14 DPS) [quest]; Enriched Seal (271903, -4.55 DPS) [rep]; Uther's Strength (11302, -5.97 DPS) [world_drop] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Thunderbrew's Boot Flask (744, -1.77 DPS, sim-verified) [quest]; Enriched Seal (271903, -2.98 DPS) [rep]; Uther's Strength (11302, -4.40 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamstaff (249454, -9.92 DPS, sim-verified) [crafted]; Resurgence Rod (17743, -14.89 DPS) [quest]; Charstone Dirk (17710, -18.63 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Engineer's Guild Headpiece; neck: Lei of Lilies; shoulder: Living Shoulders; back: Featherskin Cape; chest: Embrace of the Wind Serpent; wrist: Mender's Leather Bracers; hands: Gloves of the Atal'ai Prophet; waist: Earthenweave Cord; legs: Senior Designer's Pantaloons; feet: Sandals of the Insurgent; finger1: Eye of Adaegus; finger2: Choking Band; trinket2: Evonice's Landin' Pilla; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 815, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60 (night-elf, 0530200100000000-00000000000000000000-5553005153113210)

Set DPS (verified): 496.1. Weights run: 8.2s. Verify run: 293.7s. 1986 eligible items had no known source.

Race sweep: high-order-skyborne 503.4 ± 0.3, night-elf 495.8 ± 0.3 (heal_score); high-order-skyborne beats the guide race night-elf by 7.6 (1.54%) ± 0.4, beyond error.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.048, intellect=2.232 ± 0.135, spirit=5.678 ± 0.117, mp5=6.515 ± 0.141, crit=1.014 ± 0.032 per rating point (14 rating = 1%, 14.190 per %), spell_haste=not significant (0.957 ± 1.903)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 192.7 healing_power points (28.36 DPS) | yes | Lieutenant Commander's Dragonhide Headdress (227199, -3.69 DPS) [vendor]; Field Marshal's Dragonhide Headdress (231701, -4.09 DPS) [vendor]; Wildheart Cowl (16720, -4.46 DPS, sim-verified) [dungeon] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 115.0 healing_power points (16.92 DPS) | yes | Dragon's Eye (10829, -1.43 DPS, sim-verified) [dungeon]; Lei of Lilies (1315, -4.38 DPS) [world_drop]; Lady Alizabeth's Pendant (13002, -4.38 DPS) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Dragonhide Pauldrons (227201, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Pauldrons (231705, +0.00 DPS) [vendor]; Mantle of the Scarlet Crusade (22405, -18.30 DPS, sim-verified) [dungeon] |
| back | Featherskin Cape (10843) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, -2.96 DPS) [quest]; Indomitable Cloak (14683, -3.00 DPS) [world_drop]; Ethereal Mist Cape (11873, -3.36 DPS, sim-verified) [quest] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 225.7 healing_power points (33.21 DPS) | yes | Ironfeather Breastplate (15066, -3.60 DPS, sim-verified) [crafted]; Robes of the Exalted (13346, -3.81 DPS) [dungeon]; Mooncloth Vest (14138, -6.94 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wyrmthalak's Shackles (13958, +0.00 DPS) [quest]; Marshal's Dragonhide Bracers (16445, +0.00 DPS) [pvp]; Bleak Howler Armguards (13208, -15.02 DPS, sim-verified) [dungeon] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swarmtender's Gloves (275607, -0.63 DPS) [crafted]; Hands of the Exalted Herald (12554, -0.84 DPS) [dungeon]; Wildheart Gloves (16717, -16.65 DPS, sim-verified) [dungeon] |
| waist | Caretaker's Cord (272398) | Pix Xizzix [vendor] | 159.3 healing_power points (23.44 DPS) | yes | Sash of Mercy (14553, -5.39 DPS, sim-verified) [world_drop]; Wisdom of the Timbermaw (19047, -6.76 DPS) [crafted]; Marshal's Dragonhide Waistguard (16447, -6.87 DPS) [pvp] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Haunting Specter Leggings (11929, -2.01 DPS) [dungeon]; Wolfshear Leggings (13206, -2.96 DPS) [dungeon]; Leggings of Arcana (12756, -13.00 DPS, sim-verified) [quest] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Boots (15802, -0.12 DPS) [crafted]; Wildheart Boots (16715, -0.33 DPS) [dungeon]; Sandals of the Insurgent (13111, -19.35 DPS, sim-verified) [world_drop] |
| finger1 | The Postmaster's Seal (13392) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Adaegus (5266, -0.93 DPS) [world_drop]; Band of Mending (22334, -1.05 DPS) [dungeon]; Naglering (11669, -13.25 DPS, sim-verified) [dungeon] |
| finger2 | Band of the Hierophant (13096) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Adaegus (5266, -0.48 DPS) [world_drop]; Band of Mending (22334, -0.60 DPS) [dungeon]; Channeler's Ring (272406, -1.19 DPS) [vendor] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+12.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Burst of Knowledge (11832, -2.36 DPS) [dungeon]; Thunderbrew's Boot Flask (744, -3.34 DPS) [quest]; Shard of the Splithooves (10659, -4.28 DPS) [quest] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burst of Knowledge (11832, -0.55 DPS, sim-verified) [dungeon]; Thunderbrew's Boot Flask (744, -1.67 DPS) [quest]; Shard of the Splithooves (10659, -2.60 DPS) [quest] |
| main_hand | Staff of Metanoia (22394) | Scholomance: Jandice Barov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hammer of the Grand Crusader (18717, -1.13 DPS, sim-verified) [dungeon]; Redemption (22406, -2.37 DPS) [dungeon]; Glowing Brightwood Staff (812, -4.06 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Idol of Synthesis (272429) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Idol of the Moon (23197, -0.73 DPS, sim-verified) [world_drop] |

**New at 60:** head: Living Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; wrist: Devout Bracers; hands: Devout Gloves; waist: Caretaker's Cord; legs: Devout Skirt; feet: Devout Sandals; finger1: The Postmaster's Seal; finger2: Band of the Hierophant; main_hand: Staff of Metanoia; ranged: Idol of Synthesis

No-known-source sample (15 of 1986, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60, raid preset (night-elf, 0530200100000000-00000000000000000000-5553005153113210)

Set DPS (verified): 1028.3. Weights run: 7.7s. Verify run: 253.9s. 1986 eligible items had no known source.

Race sweep: high-order-skyborne 1035.9 ± 0.4, night-elf 1028.4 ± 0.4 (heal_score); high-order-skyborne beats the guide race night-elf by 7.5 (0.73%) ± 0.6, beyond error.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.041, intellect=0.945 ± 0.037, spirit=2.066 ± 0.035, mp5=1.923 ± 0.040, crit=0.416 ± 0.012 per rating point (14 rating = 1%, 5.819 per %), spell_haste=not significant (0.357 ± 0.422)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 118.9 healing_power points (47.17 DPS) | yes | Field Marshal's Dragonhide Headdress (231701, -11.43 DPS) [vendor]; Lieutenant Commander's Dragonhide Headdress (227199, -11.89 DPS) [vendor]; Mooncloth Circlet (14140, -16.52 DPS, sim-verified) [crafted] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 55.7 healing_power points (22.08 DPS) | yes | Dragon's Eye (10829, -4.63 DPS) [dungeon]; Heart of the Fiend (13960, -5.16 DPS) [dungeon]; Animated Chain Necklace (18723, -5.82 DPS, sim-verified) [dungeon] |
| shoulder | Mooncloth Shoulders (14139) | Tailoring [crafted] | 59.7 healing_power points (23.69 DPS) | yes | Lieutenant Commander's Dragonhide Pauldrons (227201, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Pauldrons (231705, +0.00 DPS) [vendor]; Living Shoulders (15061, -0.74 DPS) [crafted] |
| back | Caretaker's Cape (19530) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Featherskin Cape (10843, -1.11 DPS, sim-verified) [world]; Drape of Recovery (272413, -1.52 DPS) [vendor]; Shroud of the Exile (15421, -2.63 DPS) [quest] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 104.4 healing_power points (41.42 DPS) | yes | Robes of the Exalted (13346, -3.18 DPS, sim-verified) [dungeon]; Field Marshal's Dragonhide Tunic (231702, -5.00 DPS) [vendor]; Ironfeather Breastplate (15066, -5.65 DPS) [crafted] |
| wrist | Bleak Howler Armguards (13208) | Blackrock Spire: Gizrul the Slavener [dungeon] | 52.4 healing_power points (20.79 DPS) | yes | Wyrmthalak's Shackles (13958, -2.38 DPS, sim-verified) [quest]; Dryad's Wrist Bindings (19595, -3.33 DPS) [rep]; Earthenweave Cuffs (254131, -3.87 DPS) [crafted] |
| hands | Hands of the Exalted Herald (12554) | Blackrock Depths: Princess Moira Bronzebeard  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Dragonhide Gauntlets (231706, -1.68 DPS) [pvp]; Wildheart Gloves (16717, -1.90 DPS, sim-verified) [dungeon]; Devout Gloves (16692, -3.37 DPS) [dungeon] |
| waist | Sash of Mercy (14553) | World drop [world_drop] | 77.8 healing_power points (30.85 DPS) | yes | Caretaker's Cord (272398, -1.09 DPS, sim-verified) [vendor]; Wisdom of the Timbermaw (19047, -5.94 DPS) [crafted]; Earthenweave Cord (254077, -11.53 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Leggings of Arcana (12756, +0.00 DPS) [quest]; Knight-Captain's Dragonhide Legguards (227200, +0.00 DPS) [vendor]; Marshal's Dragonhide Legguards (231703, +0.00 DPS) [pvp] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 64.1 healing_power points (25.42 DPS) | yes | Faith Healer's Boots (22247, -1.88 DPS, sim-verified) [dungeon]; Marshal's Dragonhide Boots (231698, -2.01 DPS) [pvp]; Devout Sandals (16691, -3.38 DPS) [dungeon] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fordring's Seal (16058, -2.91 DPS) [quest]; The Postmaster's Seal (13392, -3.72 DPS) [dungeon]; Naglering (11669, -20.89 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fordring's Seal (16058, -1.18 DPS, sim-verified) [quest]; The Postmaster's Seal (13392, -2.74 DPS) [dungeon]; Maiden's Circle (13001, -4.17 DPS) [world_drop] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Serenity Field (272439, +0.00 DPS) [vendor]; Evonice's Landin' Pilla (18951, -1.35 DPS, sim-verified) [quest] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hammer of the Grand Crusader (18717, -0.52 DPS) [dungeon]; Staff of Metanoia (22394, -0.54 DPS) [dungeon]; Argent Crusader (13249, -4.75 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Idol of Synthesis (272429) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Idol of the Moon (23197, -3.84 DPS, sim-verified) [world_drop] |

**New at 60:** head: Living Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Mooncloth Shoulders; back: Caretaker's Cape; wrist: Bleak Howler Armguards; hands: Hands of the Exalted Herald; waist: Sash of Mercy; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of Mending; finger2: Blessed Band of Light; trinket1: Dormant Heart of the Mountain; trinket2: Burst of Knowledge; ranged: Idol of Synthesis

No-known-source sample (15 of 1986, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

## Horde

### Band 20 (tauren, 0000000000000000-00000000000000000000-2503000100000000)

Set DPS (verified): 75.7. Weights run: 14.0s. Verify run: 44.7s. 247 eligible items had no known source.

Race sweep: windshaper-skyborne 75.8 ± 0.0, tauren 75.7 ± 0.0 (heal_score); windshaper-skyborne beats the guide race tauren by 0.1 (0.11%) ± 0.0, beyond error.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.029, intellect=0.665 ± 0.003, spirit=1.680 ± 0.010, mp5=1.940 ± 0.010, crit=0.043 ± 0.002 per rating point (14 rating = 1%, 0.603 per %), spell_haste=not significant (-0.047 ± 0.034)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (4.62 DPS) | yes | Wisdom's Leather Hood (252507, -0.07 DPS, sim-verified) [crafted]; Stormrider's Leather Hood (252506, -0.74 DPS) [crafted]; Shadow Goggles (4373, -1.18 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 5.0 healing_power points (1.29 DPS) | yes | Scholarly Pendant (277203, -0.12 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 11.8 healing_power points (3.02 DPS) | yes | Magician's Mantle (12998, -0.20 DPS) [world_drop]; Reinforced Woolen Shoulders (4315, -1.05 DPS) [crafted]; Forest Leather Mantle (4709, -1.29 DPS) [world_drop] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.10 DPS, sim-verified) [crafted]; Spirit Cloak (4792, -1.45 DPS) [vendor]; Antiquated Cloak (4799, -1.45 DPS) [vendor] |
| chest | Robe of the Moccasin (6465) | Wailing Caverns: Lord Cobrahn [dungeon] | 25.8 healing_power points (6.62 DPS) | yes | Corsair's Overshirt (5202, -0.13 DPS, sim-verified) [dungeon]; Wisdom's Leather Armor (252493, -0.62 DPS) [crafted]; Scholarly Robes (2034, -0.68 DPS) [dungeon] |
| wrist | Drakewing Bands (12999) | World drop [world_drop] | 10.1 healing_power points (2.59 DPS) | yes | Mindthrust Bracers (1974, -0.24 DPS) [dungeon]; Crystalline Cuffs (14148, -1.29 DPS) [dungeon]; Seer's Cuffs (3645, -1.55 DPS) [dungeon] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 13.0 healing_power points (3.33 DPS) | yes | Wisdom's Leather Gloves (252499, -0.09 DPS) [crafted]; Magefist Gloves (12977, -0.76 DPS) [world_drop]; Bright Gloves (3066, -0.93 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 13.7 healing_power points (3.50 DPS) | yes | Novice Ardent's Sash (253887, -0.16 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.77 DPS) [crafted]; Stormrider's Leather Belt (252432, -1.80 DPS) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 30.7 healing_power points (7.88 DPS) | yes | Filigreed Pristine Leggings (253937, -0.13 DPS, sim-verified) [crafted]; Stormrider's Leather Pants (252502, -2.56 DPS) [crafted]; Darkweave Breeches (12987, -4.10 DPS) [world_drop] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 14.3 healing_power points (3.67 DPS) | yes | Red Woolen Boots (4313, -0.11 DPS, sim-verified) [crafted]; Black Whelp Slippers (252424, -0.68 DPS) [crafted]; Pristine Boots (253889, -0.85 DPS) [crafted] |
| finger1 | Band of Purification (12996) | World drop [world_drop] | 10.1 healing_power points (2.59 DPS) | yes | Lavishly Jeweled Ring (1156, -1.56 DPS) [dungeon]; Demon Band (12054, -1.72 DPS) [world_drop]; Horn Ring (17692, -1.72 DPS) [quest] |
| finger2 | Advisor's Ring (20426) | Warsong Outriders [rep] | 8.9 healing_power points (2.28 DPS) | yes | Lavishly Jeweled Ring (1156, -0.23 DPS, sim-verified) [dungeon]; Demon Band (12054, -1.42 DPS) [world_drop]; Horn Ring (17692, -1.42 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Living Root (6631) | Wailing Caverns: Verdan the Everliving [dungeon] | 53.4 healing_power points (13.70 DPS) | yes | Staff of the Blessed Seer (2271, -0.59 DPS, sim-verified) [dungeon]; Gnarled Necromancer's Staff (251534, -2.72 DPS) [quest]; Evocator's Blade (2567, -4.02 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Roadwatcher's Confidence; shoulder: Serpent's Shoulders; back: Battle Healer's Cloak; chest: Robe of the Moccasin; wrist: Drakewing Bands; hands: Pristine Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Band of Purification; finger2: Advisor's Ring; main_hand: Living Root

No-known-source sample (15 of 247, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5968 Rugged Boots; 6360 Steelscale Crushfish; 6478 Rat Stompers; 7187 VanCleef's Boots; 8350 The 1 Ring; 14389 Durability Shoulderpads

### Band 30 (tauren, 0100000000000000-00000000000000000000-3503004103100000)

Set DPS (verified): 130.6. Weights run: 8.7s. Verify run: 83.5s. 425 eligible items had no known source.

Race sweep: windshaper-skyborne 131.7 ± 0.1, tauren 130.7 ± 0.1 (heal_score); windshaper-skyborne beats the guide race tauren by 1.0 (0.80%) ± 0.1, beyond error.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.147, intellect=1.000 ± 0.014, spirit=1.638 ± 0.015, mp5=2.523 ± 0.024, crit=0.169 ± 0.015 per rating point (14 rating = 1%, 2.372 per %), spell_haste=not significant (0.296 ± 0.650)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 42.8 healing_power points (9.16 DPS) | yes | Enduring Cap (3020, -1.62 DPS, sim-verified) [world_drop]; Filigreed Pristine Circlet (253975, -4.45 DPS) [crafted]; Wisdom's Leather Helm (252515, -4.67 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 10.6 healing_power points (2.26 DPS) | yes | Pendant of Myzrael (4614, -0.15 DPS) [dungeon]; Glowing Green Talisman (5002, -0.15 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.35 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 18.0 healing_power points (3.85 DPS) | yes | Death Speaker Mantle (6685, -0.21 DPS) [dungeon]; Nightsky Mantle (4718, -0.39 DPS) [world_drop]; Ghostly Mantle (3324, -0.46 DPS) [quest] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battle Healer's Cloak (19529, -0.33 DPS, sim-verified) [rep]; Darkspear Raider's Cloak (272078, -1.54 DPS) [vendor] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 34.1 healing_power points (7.29 DPS) | yes | Pristine Gown (253961, -0.12 DPS) [crafted]; Wisdom's Leather Tunic (252511, -0.44 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -0.66 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 16.6 healing_power points (3.54 DPS) | yes | Mindthrust Bracers (1974, -0.86 DPS) [dungeon]; Nightsky Wristbands (6407, -1.21 DPS) [world_drop]; Drakewing Bands (12999, -1.44 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 24.7 healing_power points (5.27 DPS) | yes | Gloves of Meditation (4318, -1.75 DPS) [crafted]; Hotshot Pilot's Gloves (9491, -1.81 DPS) [dungeon]; Truefaith Gloves (7049, -2.30 DPS, sim-verified) [crafted] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 34.8 healing_power points (7.45 DPS) | yes | Skycaller's Leather Belt (252522, -2.08 DPS, sim-verified) [crafted]; Belt of Arugal (6392, -3.13 DPS) [dungeon]; Silver-lined Belt (13011, -4.12 DPS) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 39.5 healing_power points (8.44 DPS) | yes | Pristine Leggings (253987, -0.59 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -1.38 DPS) [crafted]; Wisdom's Leather Pants (252503, -1.48 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 33.6 healing_power points (7.17 DPS) | yes | Acidic Walkers (9454, -1.91 DPS, sim-verified) [dungeon]; Boots of the Enchanter (4325, -3.30 DPS) [crafted]; Glinteye Slippers (273024, -3.58 DPS) [dungeon] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 26.5 healing_power points (5.66 DPS) | yes | Electrocutioner Lagnut (9447, -1.28 DPS) [dungeon]; Solid Gold Band (279845, -1.56 DPS) [quest]; The Queen's Jewel (13094, -2.43 DPS) [world_drop] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 24.5 healing_power points (5.23 DPS) | yes | Electrocutioner Lagnut (9447, -0.90 DPS, sim-verified) [dungeon]; Solid Gold Band (279845, -1.13 DPS) [quest]; The Queen's Jewel (13094, -2.00 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Friar (3415) | Blackfathom Deeps [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Odo's Ley Staff (6318, -0.75 DPS, sim-verified) [dungeon]; Death Speaker Scepter (2816, -1.42 DPS) [dungeon]; Living Root (6631, -4.59 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Glowing Thresher Cape; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Runemender's Seal; finger2: Snake Hoop; main_hand: Staff of the Friar

No-known-source sample (15 of 425, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5821 Darkstalker Boots

### Band 40 (tauren, 0500200000000000-00000000000000000000-3503005103103000)

Set DPS (verified): 243.2. Weights run: 8.8s. Verify run: 83.2s. 606 eligible items had no known source.

Race sweep: windshaper-skyborne 243.6 ± 0.1, tauren 243.2 ± 0.2 (heal_score); windshaper-skyborne beats the guide race tauren by 0.3 (0.14%) ± 0.2, beyond error.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.218, intellect=1.391 ± 0.027, spirit=2.707 ± 0.037, mp5=3.469 ± 0.027, crit=0.345 ± 0.045 per rating point (14 rating = 1%, 4.832 per %), spell_haste=not significant (0.261 ± 0.455)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 70.0 healing_power points (15.11 DPS) | yes | Whitemane's Chapeau (7720, -0.59 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -4.48 DPS) [world_drop]; Earthen Silk Hood (254015, -4.52 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 31.1 healing_power points (6.72 DPS) | yes | Necklace of Calisea (1714, -0.56 DPS, sim-verified) [dungeon]; Amberglow Talisman (10824, -0.87 DPS) [quest]; Triune Amulet (7722, -0.90 DPS) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 40.1 healing_power points (8.65 DPS) | yes | Earthen Silk Shoulders (254033, -0.09 DPS) [crafted]; Sheepshear Mantle (13115, -0.34 DPS) [world_drop]; Mantle of Doan (7712, -0.66 DPS) [dungeon] |
| back | Cloak of Blight (6832) | Nothing But The Truth [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battle Healer's Cloak (19528, -1.15 DPS, sim-verified) [rep]; Glowing Thresher Cape (6901, -1.41 DPS) [dungeon]; Ceremonial Centaur Blanket (6789, -1.74 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 79.7 healing_power points (17.19 DPS) | yes | Dreamweave Vest (10021, -2.67 DPS, sim-verified) [crafted]; Doomsayer's Robe (4746, -3.22 DPS) [quest]; Death's Head Vestment (10581, -3.31 DPS) [dungeon] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 32.6 healing_power points (7.04 DPS) | yes | Earthen Silk Cuffs (254019, -2.00 DPS, sim-verified) [crafted]; Spidertank Oilrag (9448, -2.86 DPS) [dungeon]; Silkstream Cuffs (16791, -3.21 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Earthen Silk Gloves (254017, -0.53 DPS, sim-verified) [crafted]; Gloves of Old (9395, -0.66 DPS) [world_drop]; Warden's Gloves (14606, -1.55 DPS) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 49.4 healing_power points (10.65 DPS) | yes | Mender's Leather Belt (252523, -1.01 DPS, sim-verified) [crafted]; Windchaser Cinch (14435, -2.99 DPS) [world_drop]; Sutarn's Ring (13105, -4.15 DPS) [world_drop] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 70.2 healing_power points (15.16 DPS) | yes | Stoneweaver Leggings (9407, -1.60 DPS, sim-verified) [dungeon]; Stormcloth Pants (10010, -3.16 DPS) [crafted]; Wisdom's Leather Leggings (252519, -4.44 DPS) [crafted] |
| feet | Furen's Boots (13100) | World drop [world_drop] | 53.1 healing_power points (11.45 DPS) | yes | Mender's Leather Shoes (252533, -0.44 DPS, sim-verified) [crafted]; Thoughtcast Boots (10578, -1.05 DPS) [dungeon]; Gilded Slippers (254001, -2.70 DPS) [crafted] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 36.3 healing_power points (7.83 DPS) | yes | Electrocutioner Lagnut (9447, -1.80 DPS) [dungeon]; Welken Ring (5011, -2.24 DPS) [world_drop]; Solid Gold Band (279845, -2.54 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 34.7 healing_power points (7.49 DPS) | yes | Electrocutioner Lagnut (9447, -1.45 DPS, sim-verified) [dungeon]; Welken Ring (5011, -1.90 DPS) [world_drop]; Solid Gold Band (279845, -2.19 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+9.9 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of the Friar (3415, -5.83 DPS) [dungeon]; Odo's Ley Staff (6318, -6.73 DPS) [dungeon]; Staff of Jordan (873, -8.18 DPS, sim-verified) [world_drop] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | 32.8 healing_power points (7.08 DPS) | yes | Orb of Souls (249395, -0.64 DPS, sim-verified) [crafted]; Prophetic Cane (6803, -0.80 DPS) [quest]; Silksand Star (15964, -1.83 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Cloak of Blight; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Dreamweave Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Furen's Boots; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Hand of Righteousness; off_hand: Beacon of Hope

No-known-source sample (15 of 606, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (tauren, 0530200100000000-00000000000000000000-5503005113103210)

Set DPS (verified): 374.6. Weights run: 7.4s. Verify run: 78.2s. 814 eligible items had no known source.

Race sweep: windshaper-skyborne 376.6 ± 0.2, tauren 374.8 ± 0.3 (heal_score); windshaper-skyborne beats the guide race tauren by 1.9 (0.50%) ± 0.4, beyond error.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.105, intellect=2.018 ± 0.081, spirit=4.090 ± 0.093, mp5=5.075 ± 0.099, crit=0.624 ± 0.035 per rating point (14 rating = 1%, 8.739 per %), spell_haste=not significant (0.284 ± 1.944)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Engineer's Guild Headpiece (9534) | Divino-matic Rod [quest] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Gemburst Circlet (10751, -0.99 DPS, sim-verified) [quest]; Blood Guard's Restored Leather Helm (220875, -2.99 DPS) [vendor]; Serpentskin Helm (8261, -4.87 DPS) [world_drop] |
| neck | Lei of Lilies (1315) | World drop [world_drop] | 61.3 healing_power points (11.78 DPS) | yes | Choker of the High Shaman (4112, -3.56 DPS) [quest]; Necklace of Calisea (1714, -3.57 DPS) [dungeon]; Horizon Choker (13085, -4.62 DPS, sim-verified) [world_drop] |
| shoulder | Living Shoulders (15061) | Leatherworking [crafted] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Blood Guard's Restored Leather Spaulders (220877, -0.21 DPS) [vendor]; Ironfeather Shoulders (15067, -0.65 DPS, sim-verified) [crafted]; Mender's Leather Shoulders (252538, -2.57 DPS) [crafted] |
| back | Featherskin Cape (10843) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of Blight (6832, -3.46 DPS, sim-verified) [quest]; Wingveil Cloak (10802, -3.53 DPS) [dungeon]; Spritecaster Cape (11623, -3.78 DPS) [dungeon] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 175.9 healing_power points (33.78 DPS) | yes | Feathered Breastplate (8349, -7.19 DPS, sim-verified) [crafted]; Ghostweave Vest (14141, -8.54 DPS) [crafted]; Forest's Embrace (22272, -9.71 DPS) [quest] |
| wrist | Mender's Leather Bracers (252543) | Leatherworking [crafted] | 57.7 healing_power points (11.08 DPS) | yes | Nethergeld Cuffs (254061, -0.84 DPS, sim-verified) [crafted]; Enchanted Kodo Bracers (13119, -1.67 DPS) [world_drop]; Skycaller's Leather Bracers (252542, -2.51 DPS) [crafted] |
| hands | Gloves of the Atal'ai Prophet (10808) | Sunken Temple: Jammal'an the Prophet [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gilded Gloves (254095, -2.48 DPS) [crafted]; Mender's Leather Gauntlets (252551, -2.50 DPS, sim-verified) [crafted]; Ogreseer Fists (11665, -3.98 DPS) [dungeon] |
| waist | Earthenweave Cord (254077) | Tailoring [crafted] | 71.0 healing_power points (13.63 DPS) | yes | Gilded Cord (254037, -1.59 DPS) [crafted]; Tooled Leather Belt (252544, -2.23 DPS, sim-verified) [crafted]; Ghostweave Cord (254073, -2.30 DPS) [crafted] |
| legs | Rainstrider Leggings (11123) | Jammal'an the Prophet [quest] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Jinxed Hoodoo Kilt (9474, -0.47 DPS) [dungeon]; Dalewind Trousers (13008, -0.75 DPS) [world_drop]; Senior Designer's Pantaloons (11841, -0.84 DPS, sim-verified) [dungeon] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 97.9 healing_power points (18.81 DPS) | yes | Mistwalker Boots (10629, -2.51 DPS, sim-verified) [dungeon]; First Sergeant's Restored Leather Boots (220885, -3.24 DPS) [vendor]; Furen's Boots (13100, -3.53 DPS) [world_drop] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 61.2 healing_power points (11.75 DPS) | yes | Runemender's Seal (273031, -2.20 DPS) [world]; Snake Hoop (6750, -2.39 DPS) [quest]; Coldwater Ring (4550, -3.52 DPS) [quest] |
| finger2 | Brainlash (6440) | Zul'Farrak [dungeon] | 51.6 healing_power points (9.92 DPS) | yes | Runemender's Seal (273031, -0.36 DPS) [world]; Snake Hoop (6750, -0.55 DPS) [quest]; Coldwater Ring (4550, -1.68 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+3.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Darkspear Voodoo Seal (272061, -4.55 DPS) [vendor]; Uther's Strength (11302, -5.97 DPS) [world_drop]; Alchemist's Stone (13503, -6.28 DPS) [crafted] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-verified (+3.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Darkspear Voodoo Seal (272061, -3.20 DPS, sim-verified) [vendor]; Uther's Strength (11302, -4.40 DPS) [world_drop]; Alchemist's Stone (13503, -4.71 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamstaff (249454, -9.61 DPS, sim-verified) [crafted]; Resurgence Rod (17743, -14.89 DPS) [quest]; Charstone Dirk (17710, -18.63 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Engineer's Guild Headpiece; neck: Lei of Lilies; shoulder: Living Shoulders; back: Featherskin Cape; chest: Embrace of the Wind Serpent; wrist: Mender's Leather Bracers; hands: Gloves of the Atal'ai Prophet; waist: Earthenweave Cord; legs: Rainstrider Leggings; feet: Sandals of the Insurgent; finger1: Eye of Adaegus; finger2: Brainlash; trinket2: Evonice's Landin' Pilla; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 814, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (tauren, 0530200100000000-00000000000000000000-5553005153113210)

Set DPS (verified): 494.3. Weights run: 8.2s. Verify run: 282.2s. 1990 eligible items had no known source.

Race sweep: windshaper-skyborne 494.6 ± 0.3, tauren 494.2 ± 0.3 (heal_score); windshaper-skyborne beats the guide race tauren by 0.4 (0.08%) ± 0.4, inside error.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.048, intellect=2.232 ± 0.135, spirit=5.678 ± 0.117, mp5=6.515 ± 0.141, crit=1.014 ± 0.032 per rating point (14 rating = 1%, 14.190 per %), spell_haste=not significant (0.957 ± 1.903)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 192.7 healing_power points (28.36 DPS) | yes | Champion's Dragonhide Headdress (227205, -3.69 DPS) [vendor]; Warlord's Dragonhide Headdress (231675, -4.09 DPS) [vendor]; Wildheart Cowl (16720, -4.33 DPS, sim-verified) [dungeon] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 115.0 healing_power points (16.92 DPS) | yes | Dragon's Eye (10829, -1.58 DPS, sim-verified) [dungeon]; Lei of Lilies (1315, -4.38 DPS) [world_drop]; Lady Alizabeth's Pendant (13002, -4.38 DPS) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Dragonhide Pauldrons (227207, +0.00 DPS) [vendor]; Warlord's Dragonhide Pauldrons (231672, +0.00 DPS) [vendor]; Mantle of the Scarlet Crusade (22405, -18.43 DPS, sim-verified) [dungeon] |
| back | Featherskin Cape (10843) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of Blight (6832, -3.00 DPS) [quest]; Indomitable Cloak (14683, -3.00 DPS) [world_drop]; Shroud of the Exile (15421, -3.68 DPS, sim-verified) [quest] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 225.7 healing_power points (33.21 DPS) | yes | Ironfeather Breastplate (15066, -3.38 DPS, sim-verified) [crafted]; Robes of the Exalted (13346, -3.81 DPS) [dungeon]; Mooncloth Vest (14138, -6.94 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wyrmthalak's Shackles (13958, +0.00 DPS) [quest]; General's Dragonhide Bracers (16553, +0.00 DPS) [pvp]; Bleak Howler Armguards (13208, -14.97 DPS, sim-verified) [dungeon] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swarmtender's Gloves (275607, -0.63 DPS) [crafted]; Hands of the Exalted Herald (12554, -0.84 DPS) [dungeon]; Wildheart Gloves (16717, -16.70 DPS, sim-verified) [dungeon] |
| waist | Caretaker's Cord (272398) | Pix Xizzix [vendor] | 159.3 healing_power points (23.44 DPS) | yes | Sash of Mercy (14553, -5.25 DPS, sim-verified) [world_drop]; Wisdom of the Timbermaw (19047, -6.76 DPS) [crafted]; General's Dragonhide Belt (16556, -6.87 DPS) [pvp] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Haunting Specter Leggings (11929, -2.01 DPS) [dungeon]; Wolfshear Leggings (13206, -2.96 DPS) [dungeon]; Leggings of Arcana (12756, -13.01 DPS, sim-verified) [quest] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Boots (15802, -0.12 DPS) [crafted]; Wildheart Boots (16715, -0.33 DPS) [dungeon]; Sandals of the Insurgent (13111, -19.41 DPS, sim-verified) [world_drop] |
| finger1 | The Postmaster's Seal (13392) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Adaegus (5266, -0.93 DPS) [world_drop]; Band of Mending (22334, -1.05 DPS) [dungeon]; Naglering (11669, -13.39 DPS, sim-verified) [dungeon] |
| finger2 | Band of the Hierophant (13096) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Adaegus (5266, -0.48 DPS) [world_drop]; Band of Mending (22334, -0.60 DPS) [dungeon]; Channeler's Ring (272406, -1.19 DPS) [vendor] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+12.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Burst of Knowledge (11832, -2.36 DPS) [dungeon]; Shard of the Splithooves (10659, -4.28 DPS) [quest]; Darkspear Voodoo Seal (272061, -5.23 DPS) [vendor] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burst of Knowledge (11832, -0.84 DPS, sim-verified) [dungeon]; Shard of the Splithooves (10659, -2.60 DPS) [quest]; Darkspear Voodoo Seal (272061, -3.56 DPS) [vendor] |
| main_hand | Staff of Metanoia (22394) | Scholomance: Jandice Barov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hammer of the Grand Crusader (18717, -1.46 DPS, sim-verified) [dungeon]; Redemption (22406, -2.37 DPS) [dungeon]; Glowing Brightwood Staff (812, -4.06 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Idol of Synthesis (272429) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Idol of the Moon (23197, -0.65 DPS, sim-verified) [world_drop] |

**New at 60:** head: Living Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; wrist: Devout Bracers; hands: Devout Gloves; waist: Caretaker's Cord; legs: Devout Skirt; feet: Devout Sandals; finger1: The Postmaster's Seal; finger2: Band of the Hierophant; main_hand: Staff of Metanoia; ranged: Idol of Synthesis

No-known-source sample (15 of 1990, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (tauren, 0530200100000000-00000000000000000000-5553005153113210)

Set DPS (verified): 1024.0. Weights run: 7.7s. Verify run: 248.3s. 1990 eligible items had no known source.

Race sweep: windshaper-skyborne 1025.0 ± 0.4, tauren 1024.3 ± 0.4 (heal_score); windshaper-skyborne beats the guide race tauren by 0.7 (0.07%) ± 0.6, beyond error.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.041, intellect=0.945 ± 0.037, spirit=2.066 ± 0.035, mp5=1.923 ± 0.040, crit=0.416 ± 0.012 per rating point (14 rating = 1%, 5.819 per %), spell_haste=not significant (0.357 ± 0.422)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 118.9 healing_power points (47.17 DPS) | yes | Warlord's Dragonhide Headdress (231675, -11.43 DPS) [vendor]; Champion's Dragonhide Headdress (227205, -11.89 DPS) [vendor]; Mooncloth Circlet (14140, -16.56 DPS, sim-verified) [crafted] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 55.7 healing_power points (22.08 DPS) | yes | Dragon's Eye (10829, -4.63 DPS) [dungeon]; Heart of the Fiend (13960, -5.16 DPS) [dungeon]; Animated Chain Necklace (18723, -5.78 DPS, sim-verified) [dungeon] |
| shoulder | Mooncloth Shoulders (14139) | Tailoring [crafted] | 59.7 healing_power points (23.69 DPS) | yes | Champion's Dragonhide Pauldrons (227207, +0.00 DPS) [vendor]; Warlord's Dragonhide Pauldrons (231672, +0.00 DPS) [vendor]; Living Shoulders (15061, -0.74 DPS) [crafted] |
| back | Battle Healer's Cloak (19526) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Featherskin Cape (10843, -1.01 DPS) [world]; Drape of Recovery (272413, -1.52 DPS) [vendor]; Shroud of the Exile (15421, -2.63 DPS) [quest] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 104.4 healing_power points (41.42 DPS) | yes | Robes of the Exalted (13346, -2.57 DPS, sim-verified) [dungeon]; Warlord's Dragonhide Tunic (231674, -5.00 DPS) [vendor]; Ironfeather Breastplate (15066, -5.65 DPS) [crafted] |
| wrist | Bleak Howler Armguards (13208) | Blackrock Spire: Gizrul the Slavener [dungeon] | 52.4 healing_power points (20.79 DPS) | yes | Wyrmthalak's Shackles (13958, -1.09 DPS, sim-verified) [quest]; Dryad's Wrist Bindings (19595, -3.33 DPS) [rep]; Earthenweave Cuffs (254131, -3.87 DPS) [crafted] |
| hands | Hands of the Exalted Herald (12554) | Blackrock Depths: Princess Moira Bronzebeard  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Dragonhide Gauntlets (231676, -1.68 DPS) [vendor]; Wildheart Gloves (16717, -2.09 DPS, sim-verified) [dungeon]; Devout Gloves (16692, -3.37 DPS) [dungeon] |
| waist | Sash of Mercy (14553) | World drop [world_drop] | 77.8 healing_power points (30.85 DPS) | yes | Caretaker's Cord (272398, -2.06 DPS) [vendor]; Wisdom of the Timbermaw (19047, -5.94 DPS) [crafted]; Earthenweave Cord (254077, -11.53 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Leggings of Arcana (12756, +0.00 DPS) [quest]; Legionnaire's Dragonhide Legguards (227206, +0.00 DPS) [vendor]; General's Dragonhide Legguards (231673, +0.00 DPS) [vendor] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 64.1 healing_power points (25.42 DPS) | yes | General's Dragonhide Boots (231682, -2.01 DPS) [pvp]; Faith Healer's Boots (22247, -2.24 DPS, sim-verified) [dungeon]; Devout Sandals (16691, -3.38 DPS) [dungeon] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fordring's Seal (16058, -2.91 DPS) [quest]; The Postmaster's Seal (13392, -3.72 DPS) [dungeon]; Naglering (11669, -20.33 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fordring's Seal (16058, -0.83 DPS, sim-verified) [quest]; The Postmaster's Seal (13392, -2.74 DPS) [dungeon]; Royal Seal of Alexis (18022, -3.57 DPS) [quest] |
| trinket1 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Serenity Field (272439, +0.00 DPS) [vendor]; Weakness Analyzer (272438, -1.06 DPS) [vendor] |
| trinket2 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Metanoia (22394, -0.54 DPS) [dungeon]; Hammer of the Grand Crusader (18717, -1.38 DPS, sim-verified) [dungeon]; Argent Crusader (13249, -4.75 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Idol of Synthesis (272429) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Idol of the Moon (23197, -4.11 DPS, sim-verified) [world_drop] |

**New at 60:** head: Living Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Mooncloth Shoulders; back: Battle Healer's Cloak; wrist: Bleak Howler Armguards; hands: Hands of the Exalted Herald; waist: Sash of Mercy; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of Mending; finger2: Blessed Band of Light; trinket1: Evonice's Landin' Pilla; trinket2: Dormant Heart of the Mountain; ranged: Idol of Synthesis

No-known-source sample (15 of 1990, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

