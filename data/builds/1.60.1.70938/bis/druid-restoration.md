# Leveling BiS: Restoration

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-00000000000000000000-2503000100000000)

Set DPS (verified): 78.0. Weights run: 6.6s. Verify run: 50.7s. 251 eligible items had no known source.

Race sweep: high-order-skyborne 78.4 ± 0.0, night-elf 78.0 ± 0.0 (heal_score); high-order-skyborne beats the guide race night-elf by 0.3 (0.43%) ± 0.0, beyond error.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.245, intellect=0.903 ± 0.019, spirit=1.690 ± 0.032, mp5=2.982 ± 0.064, crit=0.121 ± 0.017 per rating point (14 rating = 1%, 1.690 per %), spell_haste=not significant (0.060 ± 0.243)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.55 DPS) | yes | Wisdom's Leather Hood (252507, -0.09 DPS, sim-verified) [crafted]; Stormrider's Leather Hood (252506, -0.24 DPS) [crafted]; Shadow Goggles (4373, -0.29 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 3.4 healing_power points (0.29 DPS) | yes | Tarnished Locket (279870, -0.21 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.1 healing_power points (1.13 DPS) | yes | Serpent's Shoulders (5404, -0.12 DPS, sim-verified) [dungeon]; Slime-encrusted Pads (6461, -0.36 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -0.39 DPS) [crafted] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | sim-verified (78.0 DPS) | yes | Heavy Woolen Cloak (4311, -0.07 DPS, sim-verified) [crafted]; Spirit Cloak (4792, -0.48 DPS) [vendor]; Sylvan Cloak (4793, -0.48 DPS) [vendor] |
| chest | Robe of the Moccasin (6465) | Wailing Caverns: Lord Cobrahn [dungeon] | 25.9 healing_power points (2.23 DPS) | yes | Wisdom's Leather Armor (252493, -0.07 DPS, sim-verified) [crafted]; Corsair's Overshirt (5202, -0.12 DPS) [dungeon]; Scholarly Robes (2034, -0.23 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 13.5 healing_power points (1.16 DPS) | yes | Drakewing Bands (12999, -0.40 DPS, sim-verified) [world_drop]; Repurposed Hair Band (281256, -0.71 DPS) [quest]; Owl Bracers (4796, -0.77 DPS) [vendor] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 13.7 healing_power points (1.18 DPS) | yes | Wisdom's Leather Gloves (252499, -0.01 DPS) [crafted]; Magefist Gloves (12977, -0.21 DPS) [world_drop]; Bright Gloves (3066, -0.29 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 14.6 healing_power points (1.26 DPS) | yes | Novice Ardent's Sash (253887, -0.07 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.26 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.60 DPS) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 32.2 healing_power points (2.77 DPS) | yes | Filigreed Pristine Leggings (253937, -0.10 DPS, sim-verified) [crafted]; Stormrider's Leather Pants (252502, -0.86 DPS) [crafted]; Darkweave Breeches (12987, -1.35 DPS) [world_drop] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 15.5 healing_power points (1.34 DPS) | yes | Black Whelp Slippers (252424, -0.24 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.26 DPS) [crafted]; Pristine Boots (253889, -0.33 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 11.0 healing_power points (0.94 DPS) | yes | Minor Channeling Ring (1449, -0.36 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.48 DPS) [dungeon]; Demon Band (12054, -0.65 DPS) [world_drop] |
| finger2 | Band of Purification (12996) | World drop [world_drop] | 10.1 healing_power points (0.87 DPS) | yes | Minor Channeling Ring (1449, -0.33 DPS, sim-verified) [quest]; Lavishly Jeweled Ring (1156, -0.41 DPS) [dungeon]; Demon Band (12054, -0.58 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Westfall (2042) | The Defias Brotherhood [quest] | 62.7 healing_power points (5.40 DPS) | yes | Living Root (6631, -0.81 DPS, sim-verified) [dungeon]; Staff of the Blessed Seer (2271, -1.53 DPS) [dungeon]; Evocator's Blade (2567, -2.10 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Robe of the Moccasin; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Lorekeeper's Ring; finger2: Band of Purification; main_hand: Staff of Westfall

No-known-source sample (15 of 251, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5968 Rugged Boots; 6360 Steelscale Crushfish; 6478 Rat Stompers

### Band 30 (night-elf, 0100000000000000-00000000000000000000-3503004103100000)

Set DPS (verified): 130.0. Weights run: 8.9s. Verify run: 88.5s. 424 eligible items had no known source.

Race sweep: high-order-skyborne 132.9 ± 0.1, night-elf 129.9 ± 0.1 (heal_score); high-order-skyborne beats the guide race night-elf by 3.0 (2.27%) ± 0.1, beyond error.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.033, intellect=1.104 ± 0.014, spirit=2.686 ± 0.014, mp5=3.672 ± 0.021, crit=0.098 ± 0.004 per rating point (14 rating = 1%, 1.368 per %), spell_haste=not significant (0.007 ± 0.246)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 49.1 healing_power points (8.48 DPS) | yes | Enduring Cap (3020, -1.78 DPS, sim-verified) [world_drop]; Stormrider's Leather Helm (252514, -3.84 DPS) [crafted]; Nightsky Cowl (4039, -3.87 DPS) [world_drop] |
| neck | Pendant of Myzrael (4614) (or Glowing Green Talisman (5002)) | Razorfen Downs: Splinterbone Captain [dungeon] | 16.1 healing_power points (2.78 DPS) | yes | Glowing Green Talisman (5002, +0.00 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.16 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.63 DPS) [world_drop] |
| shoulder | Mantle of Honor (3560) | Bride of the Embalmer [quest] | 26.5 healing_power points (4.58 DPS) | yes | Faerie Mantle (5820, -0.67 DPS, sim-verified) [quest]; Nightsky Mantle (4718, -0.74 DPS) [world_drop]; Bloodmage Mantle (7684, -1.31 DPS) [dungeon] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Prelacy Cape (7004, -0.24 DPS) [quest]; Caretaker's Cape (19533, -0.82 DPS) [rep]; Repairman's Cape (9605, -0.99 DPS) [quest] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 43.7 healing_power points (7.55 DPS) | yes | Wisdom's Leather Tunic (252511, -0.54 DPS, sim-verified) [crafted]; Pristine Gown (253961, -0.91 DPS) [crafted]; Robe of the Moccasin (6465, -1.36 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 20.0 healing_power points (3.46 DPS) | yes | Mindthrust Bracers (1974, -0.35 DPS, sim-verified) [dungeon]; Drakewing Bands (12999, -0.67 DPS) [world_drop]; Nightsky Wristbands (6407, -0.92 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 37.8 healing_power points (6.52 DPS) | yes | Shilly Mitts (9609, -0.83 DPS, sim-verified) [quest]; Gloves of Meditation (4318, -2.41 DPS) [crafted]; Zodiac Gloves (7106, -2.51 DPS) [quest] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 41.7 healing_power points (7.21 DPS) | yes | Skycaller's Leather Belt (252522, -2.36 DPS, sim-verified) [crafted]; Belt of Arugal (6392, -2.76 DPS) [dungeon]; Silver-lined Belt (13011, -3.64 DPS) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 47.5 healing_power points (8.21 DPS) | yes | Pristine Leggings (253987, -0.76 DPS) [crafted]; Earthen Leggings (253999, -0.93 DPS, sim-verified) [crafted]; Wisdom's Leather Pants (252503, -1.76 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 38.5 healing_power points (6.64 DPS) | yes | Acidic Walkers (9454, -2.40 DPS) [dungeon]; Boots of the Enchanter (4325, -2.48 DPS, sim-verified) [crafted]; Glinteye Slippers (273024, -2.65 DPS) [dungeon] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 31.7 healing_power points (5.48 DPS) | yes | The Queen's Jewel (13094, -1.39 DPS) [world_drop]; Monkey Ring (6748, -2.23 DPS) [quest]; Ring of Calm (6790, -2.23 DPS) [quest] |
| finger2 | Electrocutioner Lagnut (9447) | Gnomeregan: Electrocutioner 6000 [dungeon] | 27.8 healing_power points (4.80 DPS) | yes | The Queen's Jewel (13094, -0.68 DPS, sim-verified) [world_drop]; Monkey Ring (6748, -1.55 DPS) [quest]; Ring of Calm (6790, -1.55 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Friar (3415) | Blackfathom Deeps [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Odo's Ley Staff (6318, -0.82 DPS, sim-verified) [dungeon]; Staff of Westfall (2042, -3.09 DPS) [quest]; Death Speaker Scepter (2816, -3.38 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Pendant of Myzrael; shoulder: Mantle of Honor; back: Glowing Thresher Cape; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Runemender's Seal; finger2: Electrocutioner Lagnut; main_hand: Staff of the Friar

No-known-source sample (15 of 424, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 40 (night-elf, 0500200000000000-00000000000000000000-3503005103103000)

Set DPS (verified): 241.9. Weights run: 8.6s. Verify run: 76.0s. 607 eligible items had no known source.

Race sweep: high-order-skyborne 246.7 ± 0.1, night-elf 241.7 ± 0.2 (heal_score); high-order-skyborne beats the guide race night-elf by 5.0 (2.08%) ± 0.2, beyond error.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.023, intellect=1.393 ± 0.023, spirit=3.479 ± 0.029, mp5=3.800 ± 0.024, crit=0.158 ± 0.006 per rating point (14 rating = 1%, 2.209 per %), spell_haste=not significant (-0.111 ± 0.603)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 77.0 healing_power points (11.72 DPS) | yes | Whitemane's Chapeau (7720, -1.68 DPS, sim-verified) [dungeon]; Earthen Silk Hood (254015, -3.07 DPS) [crafted]; Imperial Leather Helm (6433, -3.35 DPS) [dungeon] |
| neck | Necklace of Calisea (1714) | Zul'Farrak: Witch Doctor's Chest [dungeon] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Triune Amulet (7722, -0.27 DPS) [dungeon]; Explorers' League Commendation (7746, -0.42 DPS) [quest]; Amberglow Talisman (10824, -1.04 DPS, sim-verified) [quest] |
| shoulder | Earthen Silk Shoulders (254033) | Tailoring [crafted] | 45.8 healing_power points (6.97 DPS) | yes | Mantle of Doan (7712, -0.40 DPS) [dungeon]; Sheepshear Mantle (13115, -0.41 DPS) [world_drop]; Inquisitor's Shawl (19507, -0.87 DPS) [dungeon] |
| back | Caretaker's Cape (19532) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Energy Cloak (9397, -0.09 DPS) [dungeon]; Ceremonial Centaur Blanket (6789, -0.20 DPS) [quest]; Glowing Thresher Cape (6901, -1.23 DPS, sim-verified) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 85.8 healing_power points (13.06 DPS) | yes | Doomsayer's Robe (4746, -1.10 DPS) [quest]; Death's Head Vestment (10581, -1.74 DPS) [dungeon]; Dreamweave Vest (10021, -3.21 DPS, sim-verified) [crafted] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 40.4 healing_power points (6.14 DPS) | yes | Earthen Silk Cuffs (254019, -1.72 DPS, sim-verified) [crafted]; Silkstream Cuffs (16791, -2.96 DPS) [quest]; Drakewing Bands (12999, -2.97 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Earthen Silk Gloves (254017, -0.32 DPS) [crafted]; Windchaser Handguards (14431, -0.41 DPS) [world_drop]; Gloves of Old (9395, -0.89 DPS, sim-verified) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 54.0 healing_power points (8.22 DPS) | yes | Windchaser Cinch (14435, -1.76 DPS) [world_drop]; Mender's Leather Belt (252523, -1.82 DPS, sim-verified) [crafted]; Earthen Silk Belt (7061, -2.77 DPS) [crafted] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 81.1 healing_power points (12.33 DPS) | yes | Stoneweaver Leggings (9407, -0.82 DPS, sim-verified) [dungeon]; Stormcloth Pants (10010, -2.11 DPS) [crafted]; Black Mageweave Leggings (9999, -3.71 DPS) [crafted] |
| feet | Furen's Boots (13100) | World drop [world_drop] | 65.4 healing_power points (9.95 DPS) | yes | Mender's Leather Shoes (252533, -1.11 DPS) [crafted]; Thoughtcast Boots (10578, -1.33 DPS, sim-verified) [dungeon]; Gilded Slippers (254001, -3.31 DPS) [crafted] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 38.2 healing_power points (5.82 DPS) | yes | Welken Ring (5011, -1.05 DPS) [world_drop]; The Queen's Jewel (13094, -1.16 DPS) [world_drop]; Blush Ember Ring (13093, -1.58 DPS) [world_drop] |
| finger2 | Electrocutioner Lagnut (9447) | Gnomeregan: Electrocutioner 6000 [dungeon] | 33.4 healing_power points (5.08 DPS) | yes | Welken Ring (5011, -0.31 DPS) [world_drop]; The Queen's Jewel (13094, -0.42 DPS) [world_drop]; Blush Ember Ring (13093, -0.84 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+9.3 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of the Friar (3415, -3.64 DPS) [dungeon]; Odo's Ley Staff (6318, -4.27 DPS) [dungeon]; Staff of Jordan (873, -7.95 DPS, sim-verified) [world_drop] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | 35.9 healing_power points (5.47 DPS) | yes | Orb of Lorica (11262, -0.61 DPS) [quest]; Silksand Star (15964, -0.70 DPS) [world_drop]; Orb of Souls (249395, -0.91 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Necklace of Calisea; shoulder: Earthen Silk Shoulders; back: Caretaker's Cape; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Dreamweave Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Furen's Boots; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Hand of Righteousness; off_hand: Beacon of Hope

No-known-source sample (15 of 607, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 50 (night-elf, 0530200100000000-00000000000000000000-5503005113103210)

Set DPS (verified): 378.6. Weights run: 7.5s. Verify run: 69.9s. 817 eligible items had no known source.

Race sweep: high-order-skyborne 384.8 ± 0.2, night-elf 378.4 ± 0.2 (heal_score); high-order-skyborne beats the guide race night-elf by 6.4 (1.69%) ± 0.3, beyond error.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.154, intellect=2.793 ± 0.101, spirit=9.082 ± 0.105, mp5=10.466 ± 0.103, crit=0.477 ± 0.029 per rating point (14 rating = 1%, 6.681 per %), spell_haste=not significant (-0.373 ± 1.122)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Engineer's Guild Headpiece (9534) | Divino-matic Rod [quest] | 228.8 healing_power points (18.02 DPS) | yes | Gemburst Circlet (10751, -0.38 DPS) [quest]; Imperial Red Circlet (8254, -3.77 DPS) [world_drop]; Serpentskin Helm (8261, -3.77 DPS) [world_drop] |
| neck | Lei of Lilies (1315) | World drop [world_drop] | 136.2 healing_power points (10.73 DPS) | yes | Amberglow Talisman (10824, -4.00 DPS, sim-verified) [quest]; Necklace of Calisea (1714, -4.18 DPS) [dungeon]; Explorers' League Commendation (7746, -4.29 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Restored Leather Spaulders (220876, -0.22 DPS) [vendor]; Living Shoulders (15061, -0.39 DPS, sim-verified) [crafted]; Serpentskin Spaulders (8263, -0.80 DPS) [world_drop] |
| back | Featherskin Cape (10843) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wingveil Cloak (10802, -3.02 DPS) [dungeon]; Traveler's Cloak (8297, -3.57 DPS) [world_drop]; Arcane Cloak (8286, -3.84 DPS, sim-verified) [world_drop] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 333.4 healing_power points (26.25 DPS) | yes | Feathered Breastplate (8349, -5.39 DPS) [crafted]; Ghostweave Vest (14141, -9.02 DPS, sim-verified) [crafted]; Vestments of the Atal'ai Prophet (10806, -9.50 DPS) [dungeon] |
| wrist | Mender's Leather Bracers (252543) | Leatherworking [crafted] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Nethergeld Cuffs (254061, -0.64 DPS) [crafted]; Venomshroud Armguards (14439, -0.73 DPS) [world_drop]; Enchanted Kodo Bracers (13119, -1.90 DPS, sim-verified) [world_drop] |
| hands | Gloves of the Atal'ai Prophet (10808) | Sunken Temple: Jammal'an the Prophet [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mender's Leather Gauntlets (252551, -4.47 DPS) [crafted]; Ghostweave Gloves (14142, -4.82 DPS, sim-verified) [crafted]; Ogreseer Fists (11665, -4.95 DPS) [dungeon] |
| waist | Earthenweave Cord (254077) | Tailoring [crafted] | 125.9 healing_power points (9.91 DPS) | yes | Tooled Leather Belt (252544, -1.06 DPS) [crafted]; Windchaser Cinch (14435, -1.72 DPS) [world_drop]; Ghostweave Cord (254073, -2.28 DPS, sim-verified) [crafted] |
| legs | Rainstrider Leggings (11123) | Jammal'an the Prophet [quest] | 254.1 healing_power points (20.01 DPS) | yes | Dalewind Trousers (13008, -1.58 DPS, sim-verified) [world_drop]; Jinxed Hoodoo Kilt (9474, -2.33 DPS) [dungeon]; Kilt of the Atal'ai Prophet (10807, -2.47 DPS) [dungeon] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 204.0 healing_power points (16.06 DPS) | yes | Furen's Boots (13100, -3.08 DPS) [world_drop]; Mistwalker Boots (10629, -3.65 DPS, sim-verified) [dungeon]; Sergeant Major's Restored Leather Boots (220884, -5.14 DPS) [vendor] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 125.7 healing_power points (9.90 DPS) | yes | Seedtime Hoop (9655, -2.75 DPS) [quest]; Coldwater Ring (4550, -2.81 DPS) [quest]; Brainlash (6440, -3.26 DPS) [dungeon] |
| finger2 | Choking Band (11868) | Ogre Head On A Stick = Party [quest] | 118.1 healing_power points (9.30 DPS) | yes | Coldwater Ring (4550, -2.20 DPS) [quest]; Seedtime Hoop (9655, -2.34 DPS, sim-verified) [quest]; Brainlash (6440, -2.66 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+4.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Thunderbrew's Boot Flask (744, -2.86 DPS) [quest]; Enriched Seal (271903, -4.46 DPS) [rep]; Alchemist's Stone (13503, -5.72 DPS) [crafted] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Thunderbrew's Boot Flask (744, -1.10 DPS, sim-verified) [quest]; Enriched Seal (271903, -3.03 DPS) [rep]; Alchemist's Stone (13503, -4.29 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Soulkeeper (1607, -3.61 DPS) [world_drop]; Resurgence Rod (17743, -6.74 DPS) [quest]; Dreamstaff (249454, -16.74 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Engineer's Guild Headpiece; neck: Lei of Lilies; shoulder: Ironfeather Shoulders; back: Featherskin Cape; chest: Embrace of the Wind Serpent; wrist: Mender's Leather Bracers; hands: Gloves of the Atal'ai Prophet; waist: Earthenweave Cord; legs: Rainstrider Leggings; feet: Sandals of the Insurgent; finger1: Eye of Adaegus; finger2: Choking Band; trinket2: Evonice's Landin' Pilla; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 817, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60 (night-elf, 0530200100000000-00000000000000000000-5553005153113210)

Set DPS (verified): 497.8. Weights run: 7.9s. Verify run: 283.7s. 1988 eligible items had no known source.

Race sweep: high-order-skyborne 504.1 ± 0.3, night-elf 498.0 ± 0.3 (heal_score); high-order-skyborne beats the guide race night-elf by 6.1 (1.22%) ± 0.4, beyond error.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.147, intellect=2.924 ± 0.120, spirit=10.015 ± 0.121, mp5=10.704 ± 0.145, crit=0.602 ± 0.037 per rating point (14 rating = 1%, 8.430 per %), spell_haste=not significant (1.084 ± 1.204)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Feathermoon Headdress (13113, +0.00 DPS) [world_drop]; Living Crown (252561, +0.00 DPS) [crafted]; Wildheart Cowl (16720, -16.34 DPS, sim-verified) [dungeon] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 170.8 healing_power points (14.47 DPS) | yes | Dragon's Eye (10829, -1.46 DPS, sim-verified) [dungeon]; Lei of Lilies (1315, -1.74 DPS) [world_drop]; Lady Alizabeth's Pendant (13002, -1.74 DPS) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Dragonhide Pauldrons (227201, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Pauldrons (231705, +0.00 DPS) [vendor]; Mantle of the Scarlet Crusade (22405, -18.05 DPS, sim-verified) [dungeon] |
| back | Featherskin Cape (10843) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ethereal Mist Cape (11873, -2.57 DPS) [quest]; Indomitable Cloak (14683, -3.31 DPS, sim-verified) [world_drop]; Arcane Cloak (8286, -3.31 DPS) [world_drop] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 362.7 healing_power points (30.74 DPS) | yes | Robes of the Exalted (13346, -4.15 DPS) [dungeon]; Ironfeather Breastplate (15066, -4.15 DPS, sim-verified) [crafted]; Ghostweave Vest (14141, -5.74 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wyrmthalak's Shackles (13958, +0.00 DPS) [quest]; Marshal's Dragonhide Bracers (16445, +0.00 DPS) [pvp]; Bleak Howler Armguards (13208, -14.32 DPS, sim-verified) [dungeon] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swarmtender's Gloves (275607, +0.00 DPS) [crafted]; Gloves of the Atal'ai Prophet (10808, -0.10 DPS) [dungeon]; Wildheart Gloves (16717, -16.29 DPS, sim-verified) [dungeon] |
| waist | Caretaker's Cord (272398) | Pix Xizzix [vendor] | 263.4 healing_power points (22.32 DPS) | yes | Sash of Mercy (14553, -5.40 DPS, sim-verified) [world_drop]; Marshal's Dragonhide Waistguard (16447, -7.64 DPS) [pvp]; Grimgore Noose (13403, -8.44 DPS) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Haunting Specter Leggings (11929, -0.23 DPS) [dungeon]; Earthborn Kilt (9402, -0.66 DPS) [world_drop]; Leggings of Arcana (12756, -1.23 DPS, sim-verified) [quest] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wildheart Boots (16715, -0.25 DPS) [dungeon]; Mooncloth Boots (15802, -1.30 DPS) [crafted]; Sandals of the Insurgent (13111, -18.65 DPS, sim-verified) [world_drop] |
| finger1 | The Postmaster's Seal (13392) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Adaegus (5266, -0.36 DPS) [world_drop]; Choking Band (11868, -1.00 DPS) [quest]; Naglering (11669, -13.76 DPS, sim-verified) [dungeon] |
| finger2 | Band of the Hierophant (13096) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Adaegus (5266, -0.67 DPS, sim-verified) [world_drop]; Choking Band (11868, -0.78 DPS) [quest]; Channeler's Ring (272406, -1.56 DPS) [vendor] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+12.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Burst of Knowledge (11832, -2.93 DPS) [dungeon]; Thunderbrew's Boot Flask (744, -3.40 DPS) [quest]; Shard of the Splithooves (10659, -4.74 DPS) [quest] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burst of Knowledge (11832, -0.95 DPS, sim-verified) [dungeon]; Thunderbrew's Boot Flask (744, -1.70 DPS) [quest]; Shard of the Splithooves (10659, -3.05 DPS) [quest] |
| main_hand | Staff of Metanoia (22394) | Scholomance: Jandice Barov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dancing Sliver (15854, -0.48 DPS) [quest]; Hammer of the Grand Crusader (18717, -1.52 DPS) [dungeon]; Hand of Edward the Odd (2243, -4.02 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Idol of Synthesis (272429) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Idol of the Moon (23197, -1.12 DPS, sim-verified) [world_drop] |

**New at 60:** head: Devout Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; wrist: Devout Bracers; hands: Devout Gloves; waist: Caretaker's Cord; legs: Devout Skirt; feet: Devout Sandals; finger1: The Postmaster's Seal; finger2: Band of the Hierophant; main_hand: Staff of Metanoia; ranged: Idol of Synthesis

No-known-source sample (15 of 1988, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60, raid preset (night-elf, 0530200100000000-00000000000000000000-5553005153113210)

Set DPS (verified): 1025.6. Weights run: 7.5s. Verify run: 243.2s. 1988 eligible items had no known source.

Race sweep: high-order-skyborne 1033.2 ± 0.4, night-elf 1025.3 ± 0.5 (heal_score); high-order-skyborne beats the guide race night-elf by 7.9 (0.77%) ± 0.6, beyond error.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.251, intellect=0.831 ± 0.039, spirit=1.408 ± 0.038, mp5=1.887 ± 0.041, crit=0.435 ± 0.040 per rating point (14 rating = 1%, 6.087 per %), spell_haste=not significant (-0.497 ± 0.728)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 107.7 healing_power points (61.73 DPS) | yes | Mooncloth Circlet (14140, -16.33 DPS, sim-verified) [crafted]; Field Marshal's Dragonhide Headdress (231701, -16.77 DPS) [vendor]; Lieutenant Commander's Dragonhide Headdress (227199, -17.64 DPS) [vendor] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 47.0 healing_power points (26.97 DPS) | yes | Animated Chain Necklace (18723, -5.15 DPS, sim-verified) [dungeon]; Heart of the Fiend (13960, -6.29 DPS) [dungeon]; Dragon's Eye (10829, -7.41 DPS) [dungeon] |
| shoulder | Mooncloth Shoulders (14139) | Tailoring [crafted] | 52.5 healing_power points (30.09 DPS) | yes | Lieutenant Commander's Dragonhide Pauldrons (227201, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Pauldrons (231705, +0.00 DPS) [vendor]; Living Shoulders (15061, -0.97 DPS, sim-verified) [crafted] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Caretaker's Cape (19530, -0.74 DPS) [rep]; Shroud of the Exile (15421, -4.81 DPS) [quest] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 83.5 healing_power points (47.90 DPS) | yes | Field Marshal's Dragonhide Tunic (231702, -1.84 DPS) [vendor]; Robes of the Exalted (13346, -2.60 DPS, sim-verified) [dungeon]; Forest's Embrace (22272, -5.43 DPS) [quest] |
| wrist | Bleak Howler Armguards (13208) | Blackrock Spire: Gizrul the Slavener [dungeon] | 42.2 healing_power points (24.19 DPS) | yes | Wyrmthalak's Shackles (13958, -1.58 DPS, sim-verified) [quest]; Dryad's Wrist Bindings (19595, -2.11 DPS) [rep]; Earthenweave Cuffs (254131, -3.50 DPS) [crafted] |
| hands | Hands of the Exalted Herald (12554) | Blackrock Depths: Princess Moira Bronzebeard  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Dragonhide Gauntlets (231706, -1.81 DPS) [pvp]; Mender's Leather Gauntlets (252551, -4.31 DPS, sim-verified) [crafted]; Gilded Gloves (254095, -5.80 DPS) [crafted] |
| waist | Sash of Mercy (14553) | World drop [world_drop] | 69.9 healing_power points (40.07 DPS) | yes | Caretaker's Cord (272398, -7.52 DPS) [vendor]; Wisdom of the Timbermaw (19047, -8.36 DPS) [crafted]; Earthenweave Cord (254077, -16.29 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Leggings of Arcana (12756, +0.00 DPS) [quest]; Knight-Captain's Dragonhide Legguards (227200, +0.00 DPS) [vendor]; Marshal's Dragonhide Legguards (231703, +0.00 DPS) [pvp] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 53.9 healing_power points (30.92 DPS) | yes | Faith Healer's Boots (22247, -1.08 DPS, sim-verified) [dungeon]; Marshal's Dragonhide Boots (231698, -2.53 DPS) [pvp]; Devout Sandals (16691, -6.13 DPS) [dungeon] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fordring's Seal (16058, -1.88 DPS) [quest]; Rosewine Circle (13178, -5.70 DPS) [dungeon]; Naglering (11669, -18.62 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fordring's Seal (16058, -1.58 DPS, sim-verified) [quest]; Rosewine Circle (13178, -3.98 DPS) [dungeon]; The Postmaster's Seal (13392, -5.66 DPS) [dungeon] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hammer of the Grand Crusader (18717, -3.57 DPS) [dungeon]; Staff of Metanoia (22394, -4.78 DPS) [dungeon]; Argent Crusader (13249, -7.32 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Idol of Synthesis (272429) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Idol of the Moon (23197, -3.22 DPS, sim-verified) [world_drop] |

**New at 60:** head: Living Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Mooncloth Shoulders; back: Drape of Recovery; wrist: Bleak Howler Armguards; hands: Hands of the Exalted Herald; waist: Sash of Mercy; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of Mending; finger2: Blessed Band of Light; trinket1: Dormant Heart of the Mountain; trinket2: Burst of Knowledge; ranged: Idol of Synthesis

No-known-source sample (15 of 1988, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

## Horde

### Band 20 (tauren, 0000000000000000-00000000000000000000-2503000100000000)

Set DPS (verified): 77.2. Weights run: 6.6s. Verify run: 52.9s. 247 eligible items had no known source.

Race sweep: windshaper-skyborne 77.4 ± 0.0, tauren 77.2 ± 0.0 (heal_score); windshaper-skyborne beats the guide race tauren by 0.2 (0.27%) ± 0.0, beyond error.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.245, intellect=0.903 ± 0.019, spirit=1.690 ± 0.032, mp5=2.982 ± 0.064, crit=0.121 ± 0.017 per rating point (14 rating = 1%, 1.690 per %), spell_haste=not significant (0.060 ± 0.243)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.55 DPS) | yes | Wisdom's Leather Hood (252507, -0.09 DPS, sim-verified) [crafted]; Stormrider's Leather Hood (252506, -0.24 DPS) [crafted]; Shadow Goggles (4373, -0.29 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 5.1 healing_power points (0.44 DPS) | yes | Scholarly Pendant (277203, -0.17 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.1 healing_power points (1.13 DPS) | yes | Serpent's Shoulders (5404, -0.22 DPS, sim-verified) [dungeon]; Slime-encrusted Pads (6461, -0.36 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -0.39 DPS) [crafted] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | sim-verified (77.2 DPS) | yes | Heavy Woolen Cloak (4311, -0.18 DPS, sim-verified) [crafted]; Spirit Cloak (4792, -0.48 DPS) [vendor]; Antiquated Cloak (4799, -0.48 DPS) [vendor] |
| chest | Robe of the Moccasin (6465) | Wailing Caverns: Lord Cobrahn [dungeon] | 25.9 healing_power points (2.23 DPS) | yes | Corsair's Overshirt (5202, -0.12 DPS) [dungeon]; Wisdom's Leather Armor (252493, -0.19 DPS, sim-verified) [crafted]; Scholarly Robes (2034, -0.23 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 13.5 healing_power points (1.16 DPS) | yes | Drakewing Bands (12999, -0.54 DPS, sim-verified) [world_drop]; Crystalline Cuffs (14148, -0.72 DPS) [dungeon]; Owl Bracers (4796, -0.77 DPS) [vendor] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 13.7 healing_power points (1.18 DPS) | yes | Wisdom's Leather Gloves (252499, -0.01 DPS) [crafted]; Magefist Gloves (12977, -0.21 DPS) [world_drop]; Bright Gloves (3066, -0.29 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 14.6 healing_power points (1.26 DPS) | yes | Novice Ardent's Sash (253887, -0.11 DPS, sim-verified) [crafted]; Wisdom's Leather Belt (252433, -0.26 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.60 DPS) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 32.2 healing_power points (2.77 DPS) | yes | Filigreed Pristine Leggings (253937, -0.22 DPS, sim-verified) [crafted]; Stormrider's Leather Pants (252502, -0.86 DPS) [crafted]; Darkweave Breeches (12987, -1.35 DPS) [world_drop] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 15.5 healing_power points (1.34 DPS) | yes | Red Woolen Boots (4313, -0.26 DPS) [crafted]; Pristine Boots (253889, -0.33 DPS) [crafted]; Black Whelp Slippers (252424, -0.39 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 11.0 healing_power points (0.94 DPS) | yes | Lavishly Jeweled Ring (1156, -0.48 DPS) [dungeon]; Loop of Sacrifice (281673, -0.56 DPS) [quest]; Horn Ring (17692, -0.65 DPS) [quest] |
| finger2 | Band of Purification (12996) | World drop [world_drop] | 10.1 healing_power points (0.87 DPS) | yes | Loop of Sacrifice (281673, -0.48 DPS) [quest]; Horn Ring (17692, -0.58 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.58 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Living Root (6631) | Wailing Caverns: Verdan the Everliving [dungeon] | 53.5 healing_power points (4.60 DPS) | yes | Staff of the Blessed Seer (2271, -0.80 DPS, sim-verified) [dungeon]; Gnarled Necromancer's Staff (251534, -0.91 DPS) [quest]; Evocator's Blade (2567, -1.31 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Robe of the Moccasin; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Advisor's Ring; finger2: Band of Purification; main_hand: Living Root

No-known-source sample (15 of 247, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5968 Rugged Boots; 6360 Steelscale Crushfish; 6478 Rat Stompers; 7187 VanCleef's Boots; 8350 The 1 Ring; 14389 Durability Shoulderpads

### Band 30 (tauren, 0100000000000000-00000000000000000000-3503004103100000)

Set DPS (verified): 128.8. Weights run: 8.9s. Verify run: 87.4s. 425 eligible items had no known source.

Race sweep: windshaper-skyborne 129.4 ± 0.1, tauren 128.8 ± 0.1 (heal_score); windshaper-skyborne beats the guide race tauren by 0.6 (0.47%) ± 0.1, beyond error.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.033, intellect=1.104 ± 0.014, spirit=2.686 ± 0.014, mp5=3.672 ± 0.021, crit=0.098 ± 0.004 per rating point (14 rating = 1%, 1.368 per %), spell_haste=not significant (0.007 ± 0.246)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 49.1 healing_power points (8.48 DPS) | yes | Enduring Cap (3020, -1.84 DPS, sim-verified) [world_drop]; Stormrider's Leather Helm (252514, -3.84 DPS) [crafted]; Nightsky Cowl (4039, -3.87 DPS) [world_drop] |
| neck | Pendant of Myzrael (4614) (or Glowing Green Talisman (5002)) | Razorfen Downs: Splinterbone Captain [dungeon] | 16.1 healing_power points (2.78 DPS) | yes | Glowing Green Talisman (5002, +0.00 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.16 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.63 DPS) [world_drop] |
| shoulder | Nightsky Mantle (4718) | World drop [world_drop] | 22.3 healing_power points (3.85 DPS) | yes | Ghostly Mantle (3324, -0.27 DPS, sim-verified) [quest]; Desert Shoulders (15457, -0.57 DPS) [quest]; Bloodmage Mantle (7684, -0.57 DPS) [dungeon] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battle Healer's Cloak (19529, -0.35 DPS, sim-verified) [rep]; Amy's Blanket (13005, -1.67 DPS) [world_drop]; Darkspear Raider's Cloak (272078, -2.00 DPS) [vendor] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 43.7 healing_power points (7.55 DPS) | yes | Wisdom's Leather Tunic (252511, -0.57 DPS, sim-verified) [crafted]; Pristine Gown (253961, -0.91 DPS) [crafted]; Robe of the Moccasin (6465, -1.36 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 20.0 healing_power points (3.46 DPS) | yes | Mindthrust Bracers (1974, -0.27 DPS, sim-verified) [dungeon]; Drakewing Bands (12999, -0.67 DPS) [world_drop]; Nightsky Wristbands (6407, -0.92 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 37.8 healing_power points (6.52 DPS) | yes | Gloves of Meditation (4318, -1.81 DPS, sim-verified) [crafted]; Hotshot Pilot's Gloves (9491, -2.68 DPS) [dungeon]; Stonecloth Gloves (14411, -2.81 DPS) [world_drop] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 41.7 healing_power points (7.21 DPS) | yes | Skycaller's Leather Belt (252522, -2.10 DPS, sim-verified) [crafted]; Belt of Arugal (6392, -2.76 DPS) [dungeon]; Silver-lined Belt (13011, -3.64 DPS) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 47.5 healing_power points (8.21 DPS) | yes | Pristine Leggings (253987, -0.76 DPS) [crafted]; Earthen Leggings (253999, -0.81 DPS, sim-verified) [crafted]; Wisdom's Leather Pants (252503, -1.76 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 38.5 healing_power points (6.64 DPS) | yes | Boots of the Enchanter (4325, -2.38 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -2.40 DPS) [dungeon]; Glinteye Slippers (273024, -2.65 DPS) [dungeon] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 31.7 healing_power points (5.48 DPS) | yes | Solid Gold Band (279845, -1.26 DPS) [quest]; The Queen's Jewel (13094, -1.39 DPS) [world_drop]; Monkey Ring (6748, -2.23 DPS) [quest] |
| finger2 | Electrocutioner Lagnut (9447) | Gnomeregan: Electrocutioner 6000 [dungeon] | 27.8 healing_power points (4.80 DPS) | yes | Solid Gold Band (279845, -0.33 DPS, sim-verified) [quest]; The Queen's Jewel (13094, -0.71 DPS) [world_drop]; Monkey Ring (6748, -1.55 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Friar (3415) | Blackfathom Deeps [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Odo's Ley Staff (6318, -0.50 DPS, sim-verified) [dungeon]; Death Speaker Scepter (2816, -3.38 DPS) [dungeon]; Soulstaff (249393, -4.14 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Pendant of Myzrael; shoulder: Nightsky Mantle; back: Glowing Thresher Cape; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Runemender's Seal; finger2: Electrocutioner Lagnut; main_hand: Staff of the Friar

No-known-source sample (15 of 425, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5821 Darkstalker Boots

### Band 40 (tauren, 0500200000000000-00000000000000000000-3503005103103000)

Set DPS (verified): 241.3. Weights run: 8.6s. Verify run: 79.7s. 606 eligible items had no known source.

Race sweep: windshaper-skyborne 241.8 ± 0.1, tauren 240.9 ± 0.2 (heal_score); windshaper-skyborne beats the guide race tauren by 0.8 (0.35%) ± 0.2, beyond error.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.218, intellect=1.308 ± 0.060, spirit=3.096 ± 0.068, mp5=3.897 ± 0.092, crit=0.338 ± 0.026 per rating point (14 rating = 1%, 4.726 per %), spell_haste=not significant (-0.417 ± 0.964)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 72.1 healing_power points (13.47 DPS) | yes | Whitemane's Chapeau (7720, -1.39 DPS, sim-verified) [dungeon]; Earthen Silk Hood (254015, -3.58 DPS) [crafted]; Holy Shroud (2721, -3.84 DPS) [world_drop] |
| neck | Necklace of Calisea (1714) | Zul'Farrak: Witch Doctor's Chest [dungeon] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Glowing Eye of Mordresh (10769, -0.07 DPS) [dungeon]; Triune Amulet (7722, -0.22 DPS) [dungeon]; Amberglow Talisman (10824, -1.27 DPS, sim-verified) [quest] |
| shoulder | Earthen Silk Shoulders (254033) | Tailoring [crafted] | 42.8 healing_power points (7.99 DPS) | yes | Mantle of Doan (7712, -0.59 DPS) [dungeon]; Sheepshear Mantle (13115, -0.61 DPS) [world_drop]; Inquisitor's Shawl (19507, -0.70 DPS) [dungeon] |
| back | Cloak of Blight (6832) | Nothing But The Truth [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battle Healer's Cloak (19528, -1.26 DPS) [rep]; Glowing Thresher Cape (6901, -1.58 DPS) [dungeon]; Stargazer Cloak (9660, -1.74 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 82.8 healing_power points (15.46 DPS) | yes | Doomsayer's Robe (4746, -2.06 DPS) [quest]; Death's Head Vestment (10581, -2.63 DPS) [dungeon]; Dreamweave Vest (10021, -2.74 DPS, sim-verified) [crafted] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 36.2 healing_power points (6.76 DPS) | yes | Earthen Silk Cuffs (254019, -1.76 DPS, sim-verified) [crafted]; Spidertank Oilrag (9448, -2.90 DPS) [dungeon]; Silkstream Cuffs (16791, -3.23 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Earthen Silk Gloves (254017, -0.40 DPS) [crafted]; Windchaser Handguards (14431, -0.87 DPS) [world_drop]; Gloves of Old (9395, -1.03 DPS, sim-verified) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 51.0 healing_power points (9.54 DPS) | yes | Mender's Leather Belt (252523, -1.66 DPS, sim-verified) [crafted]; Windchaser Cinch (14435, -2.38 DPS) [world_drop]; Earthen Silk Belt (7061, -3.41 DPS) [crafted] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 75.2 healing_power points (14.05 DPS) | yes | Stoneweaver Leggings (9407, -0.78 DPS, sim-verified) [dungeon]; Stormcloth Pants (10010, -2.57 DPS) [crafted]; Mistspray Kilt (4976, -4.22 DPS) [quest] |
| feet | Mender's Leather Shoes (252533) | Leatherworking [crafted] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Thoughtcast Boots (10578, -0.30 DPS) [dungeon]; Furen's Boots (13100, -0.93 DPS, sim-verified) [world_drop]; Gilded Slippers (254001, -2.48 DPS) [crafted] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 35.4 healing_power points (6.62 DPS) | yes | The Queen's Jewel (13094, -1.50 DPS) [world_drop]; Solid Gold Band (279845, -1.67 DPS) [quest]; Electrocutioner Lagnut (9447, -7.02 DPS, sim-verified) [dungeon] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | The Queen's Jewel (13094, -0.15 DPS) [world_drop]; Solid Gold Band (279845, -0.32 DPS) [quest]; Electrocutioner Lagnut (9447, -0.44 DPS, sim-verified) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+8.4 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of the Friar (3415, -4.80 DPS) [dungeon]; Odo's Ley Staff (6318, -5.53 DPS) [dungeon]; Staff of Jordan (873, -6.69 DPS, sim-verified) [world_drop] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | 34.4 healing_power points (6.42 DPS) | yes | Orb of Souls (249395, -0.85 DPS, sim-verified) [crafted]; Prophetic Cane (6803, -1.11 DPS) [quest]; Silksand Star (15964, -1.22 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Necklace of Calisea; shoulder: Earthen Silk Shoulders; back: Cloak of Blight; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Dreamweave Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; finger2: Welken Ring; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Hand of Righteousness; off_hand: Beacon of Hope

No-known-source sample (15 of 606, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (tauren, 0530200100000000-00000000000000000000-5503005113103210)

Set DPS (verified): 376.0. Weights run: 7.5s. Verify run: 68.9s. 816 eligible items had no known source.

Race sweep: windshaper-skyborne 376.9 ± 0.2, tauren 376.0 ± 0.2 (heal_score); windshaper-skyborne beats the guide race tauren by 0.9 (0.24%) ± 0.3, beyond error.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.158, intellect=2.378 ± 0.078, spirit=4.511 ± 0.076, mp5=5.726 ± 0.100, crit=0.676 ± 0.039 per rating point (14 rating = 1%, 9.469 per %), spell_haste=not significant (1.158 ± 1.290)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Engineer's Guild Headpiece (9534) | Divino-matic Rod [quest] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Gemburst Circlet (10751, -0.76 DPS, sim-verified) [quest]; Blood Guard's Restored Leather Helm (220875, -2.47 DPS) [vendor]; Serpentskin Helm (8261, -3.79 DPS) [world_drop] |
| neck | Lei of Lilies (1315) | World drop [world_drop] | 67.7 healing_power points (11.39 DPS) | yes | Necklace of Calisea (1714, -3.27 DPS) [dungeon]; Choker of the High Shaman (4112, -3.31 DPS) [quest]; Horizon Choker (13085, -4.44 DPS, sim-verified) [world_drop] |
| shoulder | Living Shoulders (15061) | Leatherworking [crafted] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Blood Guard's Restored Leather Spaulders (220877, +0.00 DPS) [vendor]; Ironfeather Shoulders (15067, -0.89 DPS, sim-verified) [crafted]; Serpentskin Spaulders (8263, -1.93 DPS) [world_drop] |
| back | Featherskin Cape (10843) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wingveil Cloak (10802, -3.37 DPS) [dungeon]; Darkspear Raider's Cloak (272076, -3.50 DPS) [vendor]; Cloak of Blight (6832, -3.58 DPS, sim-verified) [quest] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 192.1 healing_power points (32.33 DPS) | yes | Ghostweave Vest (14141, -7.92 DPS) [crafted]; Feathered Breastplate (8349, -7.96 DPS, sim-verified) [crafted]; Forest's Embrace (22272, -9.47 DPS) [quest] |
| wrist | Mender's Leather Bracers (252543) | Leatherworking [crafted] | 62.7 healing_power points (10.55 DPS) | yes | Nethergeld Cuffs (254061, -0.91 DPS, sim-verified) [crafted]; Enchanted Kodo Bracers (13119, -1.36 DPS) [world_drop]; Skycaller's Leather Bracers (252542, -2.27 DPS) [crafted] |
| hands | Gloves of the Atal'ai Prophet (10808) | Sunken Temple: Jammal'an the Prophet [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gilded Gloves (254095, -2.42 DPS) [crafted]; Mender's Leather Gauntlets (252551, -2.68 DPS, sim-verified) [crafted]; Ogreseer Fists (11665, -3.59 DPS) [dungeon] |
| waist | Earthenweave Cord (254077) | Tailoring [crafted] | 75.6 healing_power points (12.73 DPS) | yes | Gilded Cord (254037, -1.27 DPS) [crafted]; Ghostweave Cord (254073, -2.02 DPS) [crafted]; Tooled Leather Belt (252544, -2.31 DPS, sim-verified) [crafted] |
| legs | Rainstrider Leggings (11123) | Jammal'an the Prophet [quest] | 135.3 healing_power points (22.77 DPS) | yes | Dalewind Trousers (13008, -0.19 DPS) [world_drop]; Jinxed Hoodoo Kilt (9474, -0.30 DPS) [dungeon]; Senior Designer's Pantaloons (11841, -0.72 DPS, sim-verified) [dungeon] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 109.2 healing_power points (18.39 DPS) | yes | Mistwalker Boots (10629, -2.08 DPS, sim-verified) [dungeon]; First Sergeant's Restored Leather Boots (220885, -3.43 DPS) [vendor]; Furen's Boots (13100, -3.44 DPS) [world_drop] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 68.4 healing_power points (11.51 DPS) | yes | Runemender's Seal (273031, -3.04 DPS) [world]; Coldwater Ring (4550, -3.48 DPS) [quest]; Cyclopean Band (11824, -4.16 DPS) [dungeon] |
| finger2 | Brainlash (6440) | Zul'Farrak [dungeon] | 57.3 healing_power points (9.65 DPS) | yes | Runemender's Seal (273031, -1.16 DPS, sim-verified) [world]; Coldwater Ring (4550, -1.62 DPS) [quest]; Cyclopean Band (11824, -2.30 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+4.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Darkspear Voodoo Seal (272061, -4.29 DPS) [vendor]; Alchemist's Stone (13503, -6.07 DPS) [crafted]; Uther's Strength (11302, -6.08 DPS) [world_drop] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-verified (+3.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Darkspear Voodoo Seal (272061, -3.01 DPS, sim-verified) [vendor]; Alchemist's Stone (13503, -4.56 DPS) [crafted]; Uther's Strength (11302, -4.56 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamstaff (249454, -10.72 DPS, sim-verified) [crafted]; Resurgence Rod (17743, -14.24 DPS) [quest]; Charstone Dirk (17710, -17.83 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Engineer's Guild Headpiece; neck: Lei of Lilies; shoulder: Living Shoulders; back: Featherskin Cape; chest: Embrace of the Wind Serpent; wrist: Mender's Leather Bracers; hands: Gloves of the Atal'ai Prophet; waist: Earthenweave Cord; legs: Rainstrider Leggings; feet: Sandals of the Insurgent; finger1: Eye of Adaegus; finger2: Brainlash; trinket2: Evonice's Landin' Pilla; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 816, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (tauren, 0530200100000000-00000000000000000000-5553005153113210)

Set DPS (verified): 495.4. Weights run: 7.9s. Verify run: 281.9s. 1992 eligible items had no known source.

Race sweep: tauren 495.6 ± 0.3, windshaper-skyborne 495.5 ± 0.3 (heal_score); the guide race tauren is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=2.083 ± 0.124, spirit=5.428 ± 0.117, mp5=6.133 ± 0.125, crit=0.934 ± 0.038 per rating point (14 rating = 1%, 13.077 per %), spell_haste=not significant (1.089 ± 1.765)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 186.5 healing_power points (29.50 DPS) | yes | Champion's Dragonhide Headdress (227205, -4.03 DPS) [vendor]; Wildheart Cowl (16720, -4.29 DPS, sim-verified) [dungeon]; Warlord's Dragonhide Headdress (231675, -4.39 DPS) [vendor] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 109.8 healing_power points (17.36 DPS) | yes | Dragon's Eye (10829, -1.60 DPS, sim-verified) [dungeon]; Lei of Lilies (1315, -4.49 DPS) [world_drop]; Lady Alizabeth's Pendant (13002, -4.49 DPS) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Dragonhide Pauldrons (227207, +0.00 DPS) [vendor]; Warlord's Dragonhide Pauldrons (231672, +0.00 DPS) [vendor]; Mantle of the Scarlet Crusade (22405, -18.36 DPS, sim-verified) [dungeon] |
| back | Featherskin Cape (10843) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of Blight (6832, -3.14 DPS) [quest]; Indomitable Cloak (14683, -3.14 DPS) [world_drop]; Shroud of the Exile (15421, -3.90 DPS, sim-verified) [quest] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 216.7 healing_power points (34.27 DPS) | yes | Ironfeather Breastplate (15066, -2.96 DPS, sim-verified) [crafted]; Robes of the Exalted (13346, -3.84 DPS) [dungeon]; Mooncloth Vest (14138, -7.26 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wyrmthalak's Shackles (13958, +0.00 DPS) [quest]; General's Dragonhide Bracers (16553, +0.00 DPS) [pvp]; Bleak Howler Armguards (13208, -14.95 DPS, sim-verified) [dungeon] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swarmtender's Gloves (275607, -0.56 DPS) [crafted]; Hands of the Exalted Herald (12554, -0.77 DPS) [dungeon]; Wildheart Gloves (16717, -17.00 DPS, sim-verified) [dungeon] |
| waist | Caretaker's Cord (272398) | Pix Xizzix [vendor] | 153.3 healing_power points (24.24 DPS) | yes | Sash of Mercy (14553, -5.48 DPS, sim-verified) [world_drop]; Wisdom of the Timbermaw (19047, -7.00 DPS) [crafted]; General's Dragonhide Belt (16556, -7.42 DPS) [pvp] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Haunting Specter Leggings (11929, -2.21 DPS) [dungeon]; Wolfshear Leggings (13206, -2.91 DPS) [dungeon]; Leggings of Arcana (12756, -13.22 DPS, sim-verified) [quest] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Boots (15802, -0.06 DPS) [crafted]; Wildheart Boots (16715, -0.33 DPS) [dungeon]; Sandals of the Insurgent (13111, -19.27 DPS, sim-verified) [world_drop] |
| finger1 | The Postmaster's Seal (13392) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of the Hierophant (13096, -0.61 DPS) [world_drop]; Eye of Adaegus (5266, -1.07 DPS) [world_drop]; Naglering (11669, -13.43 DPS, sim-verified) [dungeon] |
| finger2 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-verified (495.7 DPS) | yes | Eye of Adaegus (5266, -0.08 DPS) [world_drop]; Band of the Hierophant (13096, -0.57 DPS, sim-verified) [world_drop]; Blessed Band of Light (272407, -0.59 DPS) [vendor] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+13.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Burst of Knowledge (11832, -2.54 DPS) [dungeon]; Shard of the Splithooves (10659, -4.48 DPS) [quest]; Darkspear Voodoo Seal (272061, -5.45 DPS) [vendor] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burst of Knowledge (11832, -0.99 DPS, sim-verified) [dungeon]; Shard of the Splithooves (10659, -2.76 DPS) [quest]; Darkspear Voodoo Seal (272061, -3.73 DPS) [vendor] |
| main_hand | Staff of Metanoia (22394) | Scholomance: Jandice Barov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hammer of the Grand Crusader (18717, -2.07 DPS, sim-verified) [dungeon]; Redemption (22406, -2.52 DPS) [dungeon]; Glowing Brightwood Staff (812, -4.17 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Idol of Synthesis (272429) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Idol of the Moon (23197, -0.96 DPS, sim-verified) [world_drop] |

**New at 60:** head: Living Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; wrist: Devout Bracers; hands: Devout Gloves; waist: Caretaker's Cord; legs: Devout Skirt; feet: Devout Sandals; finger1: The Postmaster's Seal; finger2: Band of Mending; main_hand: Staff of Metanoia; ranged: Idol of Synthesis

No-known-source sample (15 of 1992, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (tauren, 0530200100000000-00000000000000000000-5553005153113210)

Set DPS (verified): 1021.4. Weights run: 7.5s. Verify run: 243.3s. 1992 eligible items had no known source.

Race sweep: windshaper-skyborne 1023.0 ± 0.4, tauren 1021.8 ± 0.6 (heal_score); windshaper-skyborne beats the guide race tauren by 1.2 (0.12%) ± 0.7, beyond error.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.053, intellect=0.933 ± 0.040, spirit=2.126 ± 0.034, mp5=2.063 ± 0.041, crit=0.420 ± 0.014 per rating point (14 rating = 1%, 5.874 per %), spell_haste=not significant (0.034 ± 0.417)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 119.6 healing_power points (46.15 DPS) | yes | Warlord's Dragonhide Headdress (231675, -11.09 DPS) [vendor]; Champion's Dragonhide Headdress (227205, -11.56 DPS) [vendor]; Mooncloth Circlet (14140, -15.87 DPS, sim-verified) [crafted] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 56.1 healing_power points (21.64 DPS) | yes | Dragon's Eye (10829, -4.31 DPS) [dungeon]; Heart of the Fiend (13960, -4.94 DPS) [dungeon]; Animated Chain Necklace (18723, -5.75 DPS, sim-verified) [dungeon] |
| shoulder | Mooncloth Shoulders (14139) | Tailoring [crafted] | 60.4 healing_power points (23.31 DPS) | yes | Champion's Dragonhide Pauldrons (227207, +0.00 DPS) [vendor]; Warlord's Dragonhide Pauldrons (231672, +0.00 DPS) [vendor]; Living Shoulders (15061, -0.67 DPS) [crafted] |
| back | Battle Healer's Cloak (19526) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Featherskin Cape (10843, -0.82 DPS) [world]; Drape of Recovery (272413, -1.45 DPS) [vendor]; Shroud of the Exile (15421, -2.62 DPS) [quest] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 106.1 healing_power points (40.96 DPS) | yes | Robes of the Exalted (13346, -2.68 DPS, sim-verified) [dungeon]; Warlord's Dragonhide Tunic (231674, -5.28 DPS) [vendor]; Ironfeather Breastplate (15066, -5.55 DPS) [crafted] |
| wrist | Bleak Howler Armguards (13208) | Blackrock Spire: Gizrul the Slavener [dungeon] | 53.2 healing_power points (20.52 DPS) | yes | Wyrmthalak's Shackles (13958, -2.39 DPS, sim-verified) [quest]; Dryad's Wrist Bindings (19595, -3.40 DPS) [rep]; Earthenweave Cuffs (254131, -3.82 DPS) [crafted] |
| hands | Hands of the Exalted Herald (12554) | Blackrock Depths: Princess Moira Bronzebeard  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Dragonhide Gauntlets (231676, -1.69 DPS) [vendor]; Wildheart Gloves (16717, -2.18 DPS, sim-verified) [dungeon]; Devout Gloves (16692, -3.15 DPS) [dungeon] |
| waist | Sash of Mercy (14553) | World drop [world_drop] | 78.5 healing_power points (30.30 DPS) | yes | Caretaker's Cord (272398, -1.34 DPS, sim-verified) [vendor]; Wisdom of the Timbermaw (19047, -5.91 DPS) [crafted]; Earthenweave Cord (254077, -11.24 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Leggings of Arcana (12756, +0.00 DPS) [quest]; Legionnaire's Dragonhide Legguards (227206, +0.00 DPS) [vendor]; General's Dragonhide Legguards (231673, +0.00 DPS) [vendor] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 64.7 healing_power points (24.97 DPS) | yes | Faith Healer's Boots (22247, -1.38 DPS, sim-verified) [dungeon]; General's Dragonhide Boots (231682, -1.98 DPS) [pvp]; Devout Sandals (16691, -3.18 DPS) [dungeon] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fordring's Seal (16058, -2.97 DPS) [quest]; The Postmaster's Seal (13392, -3.46 DPS) [dungeon]; Naglering (11669, -20.14 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fordring's Seal (16058, -1.42 DPS, sim-verified) [quest]; The Postmaster's Seal (13392, -2.58 DPS) [dungeon]; Royal Seal of Alexis (18022, -3.47 DPS) [quest] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hammer of the Grand Crusader (18717, -0.31 DPS) [dungeon]; Staff of Metanoia (22394, -2.14 DPS, sim-verified) [dungeon]; Argent Crusader (13249, -4.61 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Idol of Synthesis (272429) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Idol of the Moon (23197, -4.31 DPS, sim-verified) [world_drop] |

**New at 60:** head: Living Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Mooncloth Shoulders; back: Battle Healer's Cloak; wrist: Bleak Howler Armguards; hands: Hands of the Exalted Herald; waist: Sash of Mercy; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of Mending; finger2: Blessed Band of Light; trinket1: Burst of Knowledge; trinket2: Dormant Heart of the Mountain; ranged: Idol of Synthesis

No-known-source sample (15 of 1992, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

