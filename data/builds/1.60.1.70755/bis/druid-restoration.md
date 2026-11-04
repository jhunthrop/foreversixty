# Leveling BiS: Restoration

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-00000000000000000000-2503000100000000)

Set DPS (verified): 73.2. Weights run: 7.9s. Verify run: 63.5s. 222 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.155, intellect=1.196 ± 0.012, spirit=2.028 ± 0.020, mp5=2.968 ± 0.041, crit=0.147 ± 0.012 per rating point (14 rating = 1%, 2.056 per %), spell_haste=not significant (-0.203 ± 0.179)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Stormrider's Leather Hood (252506) | Leatherworking [crafted] | 18.2 healing_power points (2.46 DPS) | yes | Pristine Circlet (253949, -0.03 DPS) [crafted]; Shadow Goggles (4373, -0.05 DPS, sim-verified) [crafted]; Wisdom's Leather Hood (252507, -0.17 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 4.1 healing_power points (0.55 DPS) | yes | Tarnished Locket (279870, -0.62 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 15.8 healing_power points (2.12 DPS) | yes | Serpent's Shoulders (5404, -0.21 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -0.81 DPS) [crafted]; Slime-encrusted Pads (6461, -0.92 DPS) [dungeon] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | sim-verified (73.2 DPS) | yes | Heavy Woolen Cloak (4311, +0.00 DPS, sim-verified) [crafted]; Spirit Cloak (4792, -0.67 DPS) [vendor]; Sylvan Cloak (4793, -0.67 DPS) [vendor] |
| chest | Robe of the Moccasin (6465) | Wailing Caverns: Lord Cobrahn [dungeon] | 29.3 healing_power points (3.94 DPS) | yes | Scholarly Robes (2034, -0.19 DPS, sim-verified) [dungeon]; Corsair's Overshirt (5202, -0.28 DPS) [dungeon]; Wisdom's Leather Armor (252493, -0.30 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 14.9 healing_power points (2.00 DPS) | yes | Drakewing Bands (12999, -0.34 DPS, sim-verified) [world_drop]; Repurposed Hair Band (281256, -1.14 DPS) [quest]; Owl Bracers (4796, -1.20 DPS) [vendor] |
| hands | Wisdom's Leather Gloves (252499) | Leatherworking [crafted] | 14.8 healing_power points (1.99 DPS) | yes | Pristine Gloves (253913, -0.03 DPS) [crafted]; Magefist Gloves (12977, -0.09 DPS) [world_drop]; Bright Gloves (3066, -0.25 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 15.8 healing_power points (2.13 DPS) | yes | Novice Ardent's Sash (253887, -0.02 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.40 DPS) [crafted]; Keller's Girdle (2911, -0.84 DPS) [world_drop] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 35.3 healing_power points (4.75 DPS) | yes | Filigreed Pristine Leggings (253937, -0.27 DPS, sim-verified) [crafted]; Stormrider's Leather Pants (252502, -1.35 DPS) [crafted]; Darkweave Breeches (12987, -1.99 DPS) [world_drop] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 17.0 healing_power points (2.29 DPS) | yes | Red Woolen Boots (4313, -0.23 DPS, sim-verified) [crafted]; Black Whelp Slippers (252424, -0.43 DPS) [crafted]; Pristine Boots (253889, -0.59 DPS) [crafted] |
| finger1 | Band of Purification (12996) | World drop [world_drop] | 12.2 healing_power points (1.64 DPS) | yes | Minor Channeling Ring (1449, -0.64 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.67 DPS) [dungeon]; Demon Band (12054, -1.09 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 10.9 healing_power points (1.47 DPS) | yes | Lavishly Jeweled Ring (1156, -0.51 DPS) [dungeon]; Minor Channeling Ring (1449, -0.59 DPS, sim-verified) [quest]; Demon Band (12054, -0.93 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Westfall (2042) | The Defias Brotherhood [quest] | 66.1 healing_power points (8.91 DPS) | yes | Living Root (6631, -1.49 DPS, sim-verified) [dungeon]; Staff of the Blessed Seer (2271, -2.41 DPS) [dungeon]; Evocator's Blade (2567, -3.45 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Stormrider's Leather Hood; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Robe of the Moccasin; wrist: Mindthrust Bracers; hands: Wisdom's Leather Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Band of Purification; finger2: Lorekeeper's Ring; main_hand: Staff of Westfall

No-known-source sample (15 of 222, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18853 Insignia of the Horde

### Band 30 (night-elf, 0000000000000000-00000000000000000000-5553001100010000)

Set DPS (verified): 118.4. Weights run: 12.0s. Verify run: 113.1s. 390 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.228, intellect=1.563 ± 0.024, spirit=2.794 ± 0.044, mp5=5.141 ± 0.032, crit=0.256 ± 0.021 per rating point (14 rating = 1%, 3.577 per %), spell_haste=not significant (-0.662 ± 0.696)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 49.8 healing_power points (6.16 DPS) | yes | Enduring Cap (3020, -1.49 DPS, sim-verified) [world_drop]; Nightsky Cowl (4039, -2.11 DPS) [world_drop]; Dokebi Hat (14584, -2.30 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 17.4 healing_power points (2.16 DPS) | yes | Pendant of Myzrael (4614, -0.08 DPS) [dungeon]; Glowing Green Talisman (5002, -0.08 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.35 DPS) [world_drop] |
| shoulder | Mantle of Honor (3560) | Bride of the Embalmer [quest] | 30.5 healing_power points (3.78 DPS) | yes | Faerie Mantle (5820, -0.54 DPS) [quest]; Nightsky Mantle (4718, -0.57 DPS, sim-verified) [world_drop]; Death Speaker Mantle (6685, -0.90 DPS) [dungeon] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | sim-verified (118.4 DPS) | yes | Caretaker's Cape (19533, -0.64 DPS) [rep]; Repairman's Cape (9605, -0.74 DPS, sim-verified) [quest]; Expeditionary's Cape (271923, -1.05 DPS) [rep] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 50.1 healing_power points (6.20 DPS) | yes | Wisdom's Leather Tunic (252511, -0.52 DPS, sim-verified) [crafted]; Pristine Gown (253961, -0.99 DPS) [crafted]; Pressed Felt Robe (1997, -1.47 DPS) [world] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 24.4 healing_power points (3.02 DPS) | yes | Mindthrust Bracers (1974, -0.30 DPS, sim-verified) [dungeon]; Nightsky Wristbands (6407, -0.82 DPS) [world_drop]; Drakewing Bands (12999, -0.95 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 41.3 healing_power points (5.12 DPS) | yes | Hotshot Pilot's Gloves (9491, -1.84 DPS, sim-verified) [dungeon]; Zodiac Gloves (7106, -1.92 DPS) [quest]; Gloves of Meditation (4318, -2.08 DPS) [crafted] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 45.1 healing_power points (5.59 DPS) | yes | Skycaller's Leather Belt (252522, -1.96 DPS, sim-verified) [crafted]; Belt of Arugal (6392, -2.16 DPS) [dungeon]; Silver-lined Belt (13011, -2.46 DPS) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 51.5 healing_power points (6.37 DPS) | yes | Pristine Leggings (253987, -0.53 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -0.71 DPS) [crafted]; Stormrider's Leather Kilt (252518, -1.22 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 42.1 healing_power points (5.21 DPS) | yes | Boots of the Enchanter (4325, -1.83 DPS) [crafted]; Acidic Walkers (9454, -1.93 DPS, sim-verified) [dungeon]; Glinteye Slippers (273024, -2.27 DPS) [dungeon] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 28.6 healing_power points (3.53 DPS) | yes | The Queen's Jewel (13094, -0.38 DPS) [world_drop]; Ring of Power Regulation (273030, -0.88 DPS) [dungeon]; Ring of Calm (6790, -1.11 DPS) [quest] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 28.6 healing_power points (3.53 DPS) | yes | The Queen's Jewel (13094, -0.66 DPS, sim-verified) [world_drop]; Ring of Power Regulation (273030, -0.88 DPS) [dungeon]; Ring of Calm (6790, -1.11 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Friar (3415) | Blackfathom Deeps [dungeon] | sim-verified (118.4 DPS) | yes | Odo's Ley Staff (6318, -0.47 DPS, sim-verified) [dungeon]; Staff of Westfall (2042, -2.18 DPS) [quest]; Death Speaker Scepter (2816, -2.75 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Mantle of Honor; back: Glowing Thresher Cape; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Staff of the Friar

No-known-source sample (15 of 390, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 40 (night-elf, 0000000000000000-00000000000000000000-5553005151010000)

Set DPS (verified): 231.9. Weights run: 11.2s. Verify run: 109.4s. 509 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.189, intellect=1.028 ± 0.022, spirit=1.963 ± 0.022, mp5=3.030 ± 0.032, crit=0.232 ± 0.021 per rating point (14 rating = 1%, 3.254 per %), spell_haste=not significant (-0.162 ± 0.351)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 57.1 healing_power points (14.24 DPS) | yes | Whitemane's Chapeau (7720, -0.62 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -3.08 DPS) [world_drop]; Earthen Silk Hood (254015, -3.87 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 28.2 healing_power points (7.03 DPS) | yes | Triune Amulet (7722, -0.72 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -1.82 DPS) [dungeon]; Amberglow Talisman (10824, -2.14 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 35.4 healing_power points (8.81 DPS) | yes | Earthen Silk Shoulders (254033, -0.41 DPS) [crafted]; Mantle of Doan (7712, -1.62 DPS) [dungeon]; Sheepshear Mantle (13115, -1.78 DPS) [world_drop] |
| back | Caretaker's Cape (19532) | Silverwing Sentinels [rep] | sim-verified (231.9 DPS) | yes | Glowing Thresher Cape (6901, -0.35 DPS, sim-verified) [dungeon]; Ceremonial Centaur Blanket (6789, -2.02 DPS) [quest]; Energy Cloak (9397, -2.04 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 73.7 healing_power points (18.37 DPS) | yes | Dreamweave Vest (10021, -2.02 DPS, sim-verified) [crafted]; Doomsayer's Robe (4746, -5.58 DPS) [quest]; Civinad Robes (9623, -6.02 DPS) [quest] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 23.7 healing_power points (5.92 DPS) | yes | Spidertank Oilrag (9448, -1.41 DPS) [dungeon]; Earthen Silk Cuffs (254019, -1.56 DPS, sim-verified) [crafted]; Mindthrust Bracers (1974, -2.37 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | sim-verified (231.9 DPS) | yes | Earthen Silk Gloves (254017, -0.29 DPS, sim-verified) [crafted]; Mender's Leather Gloves (252530, -1.40 DPS) [crafted]; Gilded Handwraps (254021, -1.66 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 42.0 healing_power points (10.47 DPS) | yes | Mender's Leather Belt (252523, -0.95 DPS, sim-verified) [crafted]; Windchaser Cinch (14435, -4.02 DPS) [world_drop]; Skycaller's Leather Belt (252522, -4.48 DPS) [crafted] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 57.6 healing_power points (14.37 DPS) | yes | Stoneweaver Leggings (9407, -1.12 DPS, sim-verified) [dungeon]; Stormcloth Pants (10010, -3.29 DPS) [crafted]; Wisdom's Leather Leggings (252519, -3.92 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) | Leatherworking [crafted] | 44.9 healing_power points (11.20 DPS) | yes | Furen's Boots (13100, -1.58 DPS) [world_drop]; Thoughtcast Boots (10578, -1.60 DPS) [dungeon]; Gilded Slippers (254001, -2.46 DPS) [crafted] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 22.7 healing_power points (5.67 DPS) | yes | Welken Ring (5011, -0.96 DPS) [world_drop]; Lorekeeper's Ring (19524, -1.16 DPS) [rep]; The Queen's Jewel (13094, -1.24 DPS) [world_drop] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 22.7 healing_power points (5.67 DPS) | yes | Welken Ring (5011, -0.45 DPS, sim-verified) [world_drop]; Lorekeeper's Ring (19524, -1.16 DPS) [rep]; The Queen's Jewel (13094, -1.24 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (231.9 DPS) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (231.9 DPS) | yes | Staff of Jordan (873, -6.45 DPS, sim-verified) [world_drop]; Staff of the Friar (3415, -7.74 DPS) [dungeon]; Odo's Ley Staff (6318, -8.51 DPS) [dungeon] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | 29.9 healing_power points (7.44 DPS) | yes | Orb of Souls (249395, -0.71 DPS, sim-verified) [crafted]; Orb of Lorica (11262, -1.89 DPS) [quest]; Silksand Star (15964, -3.04 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Caretaker's Cape; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Dreamweave Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Hand of Righteousness; off_hand: Beacon of Hope

No-known-source sample (15 of 509, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

### Band 50 (night-elf, 0200000000000000-00000000000000000000-5553005153113200)

Set DPS (verified): 396.3. Weights run: 28.7s. Verify run: 114.6s. 665 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.006, intellect=0.609 ± 0.018, spirit=1.982 ± 0.020, mp5=2.010 ± 0.021, crit=0.096 ± 0.004 per rating point (14 rating = 1%, 1.342 per %), spell_haste=not significant (-0.043 ± 0.318)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Engineer's Guild Headpiece (9534) | Divino-matic Rod [quest] | 62.6 healing_power points (27.81 DPS) | yes | Cassandra's Grace (13102, -1.97 DPS, sim-verified) [world_drop]; Knight-Lieutenant's Restored Leather Helm (220874, -2.18 DPS) [vendor]; Gemburst Circlet (10751, -3.36 DPS) [quest] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 24.9 healing_power points (11.05 DPS) | yes | Triune Amulet (7722, -1.33 DPS) [dungeon]; Amberglow Talisman (10824, -2.24 DPS) [quest]; Mindburst Medallion (11196, -2.36 DPS) [quest] |
| shoulder | Living Shoulders (15061) | Leatherworking [crafted] | 56.8 healing_power points (25.22 DPS) | yes | Mender's Leather Shoulders (252538, -2.73 DPS, sim-verified) [crafted]; Knight-Lieutenant's Restored Leather Spaulders (220876, -3.93 DPS) [vendor]; Ironfeather Shoulders (15067, -6.10 DPS) [crafted] |
| back | Caretaker's Cape (19531) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -1.93 DPS) [dungeon]; Wingveil Cloak (10802, -3.15 DPS) [dungeon]; Arcane Cloak (8286, -4.49 DPS) [world_drop] |
| chest | Feathered Breastplate (8349) | Leatherworking [crafted] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Stormcloth Vest (10020, -3.40 DPS, sim-verified) [crafted]; Ghostweave Vest (14141, -3.53 DPS) [crafted]; Grimlok's Tribal Vestments (9415, -5.72 DPS) [dungeon] |
| wrist | Mender's Leather Bracers (252543) | Leatherworking [crafted] | 35.2 healing_power points (15.62 DPS) | yes | Nethergeld Cuffs (254061, -0.44 DPS) [crafted]; Skycaller's Leather Bracers (252542, -4.88 DPS) [crafted]; Enchanted Kodo Bracers (13119, -5.73 DPS) [world_drop] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gilded Gloves (254095, -1.50 DPS) [crafted]; Earthenweave Gloves (254075, -2.90 DPS) [crafted]; Ghostweave Gloves (14142, -6.93 DPS) [crafted] |
| waist | Earthenweave Cord (254077) | Tailoring [crafted] | 47.8 healing_power points (21.24 DPS) | yes | Gilded Cord (254037, -1.39 DPS, sim-verified) [crafted]; Mender's Leather Waistguard (252477, -4.66 DPS) [crafted]; Ghostweave Cord (254073, -5.33 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 64.9 healing_power points (28.85 DPS) | yes | Jinxed Hoodoo Kilt (9474, -2.31 DPS) [dungeon]; Warchief Kilt (7760, -4.24 DPS) [dungeon]; Dalewind Trousers (13008, -4.40 DPS) [world_drop] |
| feet | Sergeant Major's Restored Leather Boots (220884) | PvP reward · Alliance [vendor] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Sandals of the Insurgent (13111, -0.62 DPS) [world_drop]; Earthenweave Boots (254093, -1.09 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -1.68 DPS) [crafted] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 27.4 healing_power points (12.19 DPS) | yes | Electrocutioner Lagnut (9447, -2.03 DPS) [dungeon]; Runemender's Seal (273031, -2.03 DPS) [world]; Lorekeeper's Ring (19523, -3.29 DPS) [rep] |
| finger2 | Choking Band (11868) | Ogre Head On A Stick = Party [quest] | 25.8 healing_power points (11.45 DPS) | yes | Runemender's Seal (273031, -1.29 DPS) [world]; Electrocutioner Lagnut (9447, -1.47 DPS, sim-verified) [dungeon]; Lorekeeper's Ring (19523, -2.55 DPS) [rep] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+4.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -2.57 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -3.52 DPS) [quest]; Enriched Seal (271903, -6.10 DPS) [rep] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -0.81 DPS) [world_drop]; Enriched Seal (271903, -1.67 DPS, sim-verified) [rep]; Thunderbrew's Boot Flask (744, -1.76 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamstaff (249454, -11.18 DPS, sim-verified) [crafted]; Resurgence Rod (17743, -22.83 DPS) [quest]; Charstone Dirk (17710, -26.12 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Engineer's Guild Headpiece; shoulder: Living Shoulders; back: Caretaker's Cape; chest: Feathered Breastplate; wrist: Mender's Leather Bracers; hands: Mender's Leather Gauntlets; waist: Earthenweave Cord; legs: Senior Designer's Pantaloons; feet: Sergeant Major's Restored Leather Boots; finger1: Eye of Adaegus; finger2: Choking Band; trinket2: Evonice's Landin' Pilla; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 665, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (night-elf, 0531200100000000-00000000000000000000-5553005153113200)

Set DPS (verified): 627.6. Weights run: 9.6s. Verify run: 363.1s. 1443 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.019, intellect=0.591 ± 0.020, spirit=1.646 ± 0.019, mp5=1.932 ± 0.022, crit=0.107 ± 0.004 per rating point (14 rating = 1%, 1.503 per %), spell_haste=not significant (0.020 ± 0.163)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 106.7 healing_power points (57.22 DPS) | yes | Field Marshal's Dragonhide Headdress (231701, -15.31 DPS) [vendor]; Lieutenant Commander's Dragonhide Headdress (227199, -16.50 DPS) [vendor]; Mooncloth Circlet (14140, -16.76 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 24.7 healing_power points (13.26 DPS) | yes | Lady Alizabeth's Pendant (13002, -1.14 DPS, sim-verified) [world_drop]; Triune Amulet (7722, -2.79 DPS) [dungeon]; Archlight Talisman (15856, -3.66 DPS) [quest] |
| shoulder | Mooncloth Shoulders (14139) | Tailoring [crafted] | 55.1 healing_power points (29.56 DPS) | yes | Lieutenant Commander's Dragonhide Pauldrons (227201, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Pauldrons (231705, +0.00 DPS) [vendor]; Living Shoulders (15061, -1.23 DPS, sim-verified) [crafted] |
| back | Caretaker's Cape (19530) | Silverwing Sentinels [rep] | sim-verified (627.8 DPS) | yes | Spritecaster Cape (11623, -6.30 DPS) [dungeon]; Wingveil Cloak (10802, -8.62 DPS) [dungeon]; Cloak of Earth and Sky (279270, -8.65 DPS) [crafted] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (627.8 DPS) | yes | Knight-Captain's Dragonhide Tunic (227196, +0.00 DPS) [pvp]; Field Marshal's Dragonhide Tunic (231702, +0.00 DPS) [vendor]; Ironfeather Breastplate (15066, -13.97 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (627.8 DPS) | yes | Dryad's Wrist Bindings (19595, +0.00 DPS) [rep]; Earthenweave Cuffs (254131, -18.28 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (627.8 DPS) | yes | Wildheart Gloves (16717, +0.00 DPS) [dungeon]; Marshal's Dragonhide Gauntlets (231706, +0.00 DPS) [pvp]; Mender's Leather Gauntlets (252551, -19.78 DPS, sim-verified) [crafted] |
| waist | Elderwild Waistcord (279252) | Leatherworking [crafted] | 59.6 healing_power points (31.98 DPS) | yes | Wisdom of the Timbermaw (19047, -1.46 DPS, sim-verified) [crafted]; Earthenweave Cord (254077, -8.33 DPS) [crafted]; Devout Belt (16696, -11.80 DPS) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-verified (627.8 DPS) | yes | Knight-Captain's Dragonhide Legguards (227200, +0.00 DPS) [vendor]; Marshal's Dragonhide Legguards (231703, +0.00 DPS) [pvp]; Elderwild Pants (279254, -18.12 DPS, sim-verified) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (627.8 DPS) | yes | Marshal's Dragonhide Boots (231698, +0.00 DPS) [pvp]; Wildheart Boots (16715, -0.32 DPS) [dungeon]; Mooncloth Boots (15802, -16.53 DPS, sim-verified) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 32.0 healing_power points (17.17 DPS) | yes | Eye of Adaegus (5266, -4.68 DPS) [world_drop]; Lorekeeper's Ring (19522, -5.52 DPS) [rep]; Choking Band (11868, -5.69 DPS) [quest] |
| finger2 | Band of the Hierophant (13096) | World drop [world_drop] | 24.0 healing_power points (12.88 DPS) | yes | Eye of Adaegus (5266, -0.72 DPS, sim-verified) [world_drop]; Lorekeeper's Ring (19522, -1.22 DPS) [rep]; Choking Band (11868, -1.40 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (627.8 DPS) | yes | Evonice's Landin' Pilla (18951, -1.77 DPS) [quest]; Thunderbrew's Boot Flask (744, -3.53 DPS) [quest]; Shard of the Splithooves (10659, -4.38 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (627.8 DPS) | yes | Evonice's Landin' Pilla (18951, -1.16 DPS, sim-verified) [quest]; Thunderbrew's Boot Flask (744, -2.59 DPS) [quest]; Shard of the Splithooves (10659, -3.44 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (627.8 DPS) | yes | Grand Marshal's Warhammer (23454, +0.00 DPS) [vendor]; Hand of Edward the Odd (2243, -10.24 DPS, sim-verified) [world_drop]; Dreamstaff (249454, -24.50 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Idol of Swiftness (279250) | Enchanting [crafted] | sim-verified (627.8 DPS) | yes | Idol of the Moon (23197, -1.81 DPS, sim-verified) [world_drop] |

**New at 60:** head: Living Crown; shoulder: Mooncloth Shoulders; back: Caretaker's Cape; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Elderwild Waistcord; legs: Devout Skirt; feet: Devout Sandals; finger1: Maiden's Circle; finger2: Band of the Hierophant; trinket2: Uther's Strength; ranged: Idol of Swiftness

No-known-source sample (15 of 1443, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60, raid preset (night-elf, 0531200100000000-00000000000000000000-5553005153113200)

Set DPS (verified): 1065.0. Weights run: 6.9s. Verify run: 174.6s. 1443 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.383, intellect=1.101 ± 0.085, spirit=1.402 ± 0.086, mp5=1.242 ± 0.091, crit=0.667 ± 0.095 per rating point (14 rating = 1%, 9.341 per %), spell_haste=not significant (1.321 ± 1.565)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 112.4 healing_power points (48.67 DPS) | yes | Field Marshal's Dragonhide Headdress (231701, -12.90 DPS) [vendor]; Lieutenant Commander's Dragonhide Headdress (227199, -13.20 DPS) [vendor]; Mooncloth Circlet (14140, -16.08 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 28.8 healing_power points (12.47 DPS) | yes | Archlight Talisman (15856, -2.09 DPS, sim-verified) [quest]; Lady Alizabeth's Pendant (13002, -3.36 DPS) [world_drop]; Horizon Choker (13085, -3.37 DPS) [world_drop] |
| shoulder | Mooncloth Shoulders (14139) | Tailoring [crafted] | 52.4 healing_power points (22.69 DPS) | yes | Lieutenant Commander's Dragonhide Pauldrons (227201, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Pauldrons (231705, +0.00 DPS) [vendor]; Wildheart Spaulders (16718, -1.91 DPS, sim-verified) [dungeon] |
| back | Caretaker's Cape (19530) | Silverwing Sentinels [rep] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Spritecaster Cape (11623, -3.55 DPS) [dungeon]; Expeditionary's Cape (271907, -4.85 DPS) [rep] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | 74.6 healing_power points (32.31 DPS) | yes | Knight-Captain's Dragonhide Tunic (227196, +0.00 DPS) [pvp]; Field Marshal's Dragonhide Tunic (231702, +0.00 DPS) [vendor]; Ironfeather Breastplate (15066, -0.51 DPS) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 40.6 healing_power points (17.58 DPS) | yes | Earthenweave Cuffs (254131, -1.99 DPS) [crafted]; Mender's Leather Bracers (252543, -2.38 DPS) [crafted]; Marshal's Dragonhide Bracers (16445, -2.46 DPS) [pvp] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | sim-verified (1064.9 DPS) | yes | Marshal's Dragonhide Gauntlets (231706, +0.00 DPS) [pvp]; Gilded Gloves (254095, -1.00 DPS) [crafted]; Theramore Handwraps (271945, -2.54 DPS, sim-verified) [rep] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 59.0 healing_power points (25.55 DPS) | yes | Elderwild Waistcord (279252, -1.64 DPS) [crafted]; Devout Belt (16696, -5.80 DPS) [dungeon]; Wildheart Belt (16716, -6.62 DPS) [dungeon] |
| legs | Elderwild Pants (279254) | Leatherworking [crafted] | 73.6 healing_power points (31.88 DPS) | yes | Knight-Captain's Dragonhide Legguards (227200, +0.00 DPS) [vendor]; Marshal's Dragonhide Legguards (231703, +0.00 DPS) [pvp]; Devout Skirt (16694, -1.05 DPS, sim-verified) [dungeon] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 57.6 healing_power points (24.95 DPS) | yes | Marshal's Dragonhide Boots (231698, -1.91 DPS) [pvp]; Devout Sandals (16691, -4.29 DPS, sim-verified) [dungeon]; Sergeant Major's Restored Leather Boots (220884, -5.46 DPS) [vendor] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 34.1 healing_power points (14.77 DPS) | yes | Eye of Adaegus (5266, -4.63 DPS) [world_drop]; Deep Rooted Ring (19109, -4.81 DPS) [world]; Lorekeeper's Ring (19522, -6.56 DPS) [rep] |
| finger2 | Band of the Hierophant (13096) | World drop [world_drop] | 26.4 healing_power points (11.44 DPS) | yes | Eye of Adaegus (5266, -1.30 DPS) [world_drop]; Deep Rooted Ring (19109, -1.48 DPS) [world]; Lorekeeper's Ring (19522, -3.23 DPS) [rep] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (+12.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.21 DPS) [quest]; Thunderbrew's Boot Flask (744, -2.43 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Warhammer (23454, +0.00 DPS) [vendor]; Hand of Edward the Odd (2243, -12.21 DPS, sim-verified) [world_drop]; Dreamstaff (249454, -27.22 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Idol of Swiftness (279250) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Idol of the Moon (23197, -7.99 DPS, sim-verified) [world_drop] |

**New at 60:** head: Living Crown; shoulder: Mooncloth Shoulders; back: Caretaker's Cape; chest: Mooncloth Vest; wrist: Dryad's Wrist Bindings; waist: Wisdom of the Timbermaw; legs: Elderwild Pants; feet: Mooncloth Boots; finger1: Maiden's Circle; finger2: Band of the Hierophant; trinket1: Dormant Heart of the Mountain; trinket2: Ankh of Life; ranged: Idol of Swiftness

No-known-source sample (15 of 1443, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

## Horde

### Band 20 (tauren, 0000000000000000-00000000000000000000-2503000100000000)

Set DPS (verified): 71.7. Weights run: 7.9s. Verify run: 64.6s. 218 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.155, intellect=1.196 ± 0.012, spirit=2.028 ± 0.020, mp5=2.968 ± 0.041, crit=0.147 ± 0.012 per rating point (14 rating = 1%, 2.056 per %), spell_haste=not significant (-0.203 ± 0.179)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Stormrider's Leather Hood (252506) | Leatherworking [crafted] | 18.2 healing_power points (2.46 DPS) | yes | Pristine Circlet (253949, -0.03 DPS) [crafted]; Shadow Goggles (4373, -0.14 DPS, sim-verified) [crafted]; Wisdom's Leather Hood (252507, -0.17 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 6.1 healing_power points (0.82 DPS) | yes | Scholarly Pendant (277203, -0.34 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 15.8 healing_power points (2.12 DPS) | yes | Serpent's Shoulders (5404, -0.21 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -0.81 DPS) [crafted]; Slime-encrusted Pads (6461, -0.92 DPS) [dungeon] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.06 DPS, sim-verified) [crafted]; Spirit Cloak (4792, -0.67 DPS) [vendor]; Antiquated Cloak (4799, -0.67 DPS) [vendor] |
| chest | Robe of the Moccasin (6465) | Wailing Caverns: Lord Cobrahn [dungeon] | 29.3 healing_power points (3.94 DPS) | yes | Scholarly Robes (2034, -0.23 DPS, sim-verified) [dungeon]; Corsair's Overshirt (5202, -0.28 DPS) [dungeon]; Wisdom's Leather Armor (252493, -0.30 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 14.9 healing_power points (2.00 DPS) | yes | Drakewing Bands (12999, -0.36 DPS, sim-verified) [world_drop]; Crystalline Cuffs (14148, -1.19 DPS) [dungeon]; Owl Bracers (4796, -1.20 DPS) [vendor] |
| hands | Wisdom's Leather Gloves (252499) | Leatherworking [crafted] | 14.8 healing_power points (1.99 DPS) | yes | Pristine Gloves (253913, -0.03 DPS) [crafted]; Magefist Gloves (12977, -0.09 DPS) [world_drop]; Bright Gloves (3066, -0.25 DPS) [world_drop] |
| waist | Novice Ardent's Sash (253887) | Tailoring [crafted] | sim-verified (71.7 DPS) | yes | Pristine Sash (253925, -0.07 DPS, sim-verified) [crafted]; Wisdom's Leather Belt (252433, -0.39 DPS) [crafted]; Keller's Girdle (2911, -0.82 DPS) [world_drop] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 35.3 healing_power points (4.75 DPS) | yes | Filigreed Pristine Leggings (253937, -0.31 DPS, sim-verified) [crafted]; Stormrider's Leather Pants (252502, -1.35 DPS) [crafted]; Darkweave Breeches (12987, -1.99 DPS) [world_drop] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 17.0 healing_power points (2.29 DPS) | yes | Red Woolen Boots (4313, -0.22 DPS, sim-verified) [crafted]; Black Whelp Slippers (252424, -0.43 DPS) [crafted]; Pristine Boots (253889, -0.59 DPS) [crafted] |
| finger1 | Band of Purification (12996) | World drop [world_drop] | 12.2 healing_power points (1.64 DPS) | yes | Lavishly Jeweled Ring (1156, -0.67 DPS) [dungeon]; Loop of Sacrifice (281673, -0.83 DPS) [quest]; Horn Ring (17692, -1.09 DPS) [quest] |
| finger2 | Advisor's Ring (20426) | Warsong Outriders [rep] | 10.9 healing_power points (1.47 DPS) | yes | Loop of Sacrifice (281673, -0.67 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.73 DPS, sim-verified) [dungeon]; Horn Ring (17692, -0.93 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Living Root (6631) | Wailing Caverns: Verdan the Everliving [dungeon] | 55.1 healing_power points (7.43 DPS) | yes | Staff of the Blessed Seer (2271, -1.02 DPS, sim-verified) [dungeon]; Gnarled Necromancer's Staff (251534, -1.19 DPS) [quest]; Evocator's Blade (2567, -1.97 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Stormrider's Leather Hood; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Robe of the Moccasin; wrist: Mindthrust Bracers; hands: Wisdom's Leather Gloves; waist: Novice Ardent's Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Band of Purification; finger2: Advisor's Ring; main_hand: Living Root

No-known-source sample (15 of 218, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20441 Scout's Blade; 23173 Abomination Skin Leggings; 209617 Insignia of the Alliance; 241089 Scarlet Dagger; 248005 Northshire Hammer

### Band 30 (tauren, 0000000000000000-00000000000000000000-5553001100010000)

Set DPS (verified): 117.3. Weights run: 12.0s. Verify run: 116.3s. 391 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.228, intellect=1.563 ± 0.024, spirit=2.794 ± 0.044, mp5=5.141 ± 0.032, crit=0.256 ± 0.021 per rating point (14 rating = 1%, 3.577 per %), spell_haste=not significant (-0.662 ± 0.696)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 49.8 healing_power points (6.16 DPS) | yes | Enduring Cap (3020, -1.41 DPS, sim-verified) [world_drop]; Nightsky Cowl (4039, -2.11 DPS) [world_drop]; Dokebi Hat (14584, -2.30 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 17.4 healing_power points (2.16 DPS) | yes | Pendant of Myzrael (4614, -0.08 DPS) [dungeon]; Glowing Green Talisman (5002, -0.08 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.35 DPS) [world_drop] |
| shoulder | Nightsky Mantle (4718) | World drop [world_drop] | 26.5 healing_power points (3.28 DPS) | yes | Death Speaker Mantle (6685, -0.40 DPS, sim-verified) [dungeon]; Bloodmage Mantle (7684, -0.42 DPS) [dungeon]; Ghostly Mantle (3324, -0.46 DPS) [quest] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | sim-verified (117.4 DPS) | yes | Battle Healer's Cloak (19529, -0.67 DPS, sim-verified) [rep]; Darkspear Raider's Cloak (272078, -1.05 DPS) [vendor]; Amy's Blanket (13005, -1.21 DPS) [world_drop] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 50.1 healing_power points (6.20 DPS) | yes | Wisdom's Leather Tunic (252511, -0.70 DPS, sim-verified) [crafted]; Pristine Gown (253961, -0.99 DPS) [crafted]; Pressed Felt Robe (1997, -1.47 DPS) [world] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 24.4 healing_power points (3.02 DPS) | yes | Mindthrust Bracers (1974, -0.32 DPS, sim-verified) [dungeon]; Nightsky Wristbands (6407, -0.82 DPS) [world_drop]; Drakewing Bands (12999, -0.95 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 41.3 healing_power points (5.12 DPS) | yes | Hotshot Pilot's Gloves (9491, -1.81 DPS, sim-verified) [dungeon]; Gloves of Meditation (4318, -2.08 DPS) [crafted]; Stonecloth Gloves (14411, -2.35 DPS) [world_drop] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 45.1 healing_power points (5.59 DPS) | yes | Skycaller's Leather Belt (252522, -1.86 DPS, sim-verified) [crafted]; Belt of Arugal (6392, -2.16 DPS) [dungeon]; Silver-lined Belt (13011, -2.46 DPS) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 51.5 healing_power points (6.37 DPS) | yes | Pristine Leggings (253987, -0.51 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -0.71 DPS) [crafted]; Stormrider's Leather Kilt (252518, -1.22 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 42.1 healing_power points (5.21 DPS) | yes | Boots of the Enchanter (4325, -1.83 DPS) [crafted]; Acidic Walkers (9454, -2.02 DPS, sim-verified) [dungeon]; Glinteye Slippers (273024, -2.27 DPS) [dungeon] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 28.6 healing_power points (3.53 DPS) | yes | The Queen's Jewel (13094, -0.38 DPS) [world_drop]; Solid Gold Band (279845, -0.44 DPS) [quest]; Ring of Power Regulation (273030, -0.88 DPS) [dungeon] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 28.6 healing_power points (3.53 DPS) | yes | Solid Gold Band (279845, -0.44 DPS) [quest]; The Queen's Jewel (13094, -0.54 DPS, sim-verified) [world_drop]; Ring of Power Regulation (273030, -0.88 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Friar (3415) | Blackfathom Deeps [dungeon] | sim-verified (117.4 DPS) | yes | Odo's Ley Staff (6318, -0.48 DPS, sim-verified) [dungeon]; Death Speaker Scepter (2816, -2.75 DPS) [dungeon]; Soulstaff (249393, -3.15 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Nightsky Mantle; back: Glowing Thresher Cape; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Staff of the Friar

No-known-source sample (15 of 391, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape

### Band 40 (tauren, 0000000000000000-00000000000000000000-5553005151010000)

Set DPS (verified): 231.3. Weights run: 11.2s. Verify run: 109.4s. 508 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.189, intellect=1.028 ± 0.022, spirit=1.963 ± 0.022, mp5=3.030 ± 0.032, crit=0.232 ± 0.021 per rating point (14 rating = 1%, 3.254 per %), spell_haste=not significant (-0.162 ± 0.351)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 57.1 healing_power points (14.24 DPS) | yes | Whitemane's Chapeau (7720, -0.52 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -3.08 DPS) [world_drop]; Earthen Silk Hood (254015, -3.87 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 28.2 healing_power points (7.03 DPS) | yes | Triune Amulet (7722, -0.65 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -1.82 DPS) [dungeon]; Amberglow Talisman (10824, -2.14 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 35.4 healing_power points (8.81 DPS) | yes | Earthen Silk Shoulders (254033, -0.41 DPS) [crafted]; Mantle of Doan (7712, -1.62 DPS) [dungeon]; Sheepshear Mantle (13115, -1.78 DPS) [world_drop] |
| back | Cloak of Blight (6832) | Nothing But The Truth [quest] | sim-verified (231.4 DPS) | yes | Battle Healer's Cloak (19528, -0.51 DPS, sim-verified) [rep]; Glowing Thresher Cape (6901, -0.70 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 73.7 healing_power points (18.37 DPS) | yes | Dreamweave Vest (10021, -2.45 DPS, sim-verified) [crafted]; Doomsayer's Robe (4746, -5.58 DPS) [quest]; Civinad Robes (9623, -6.02 DPS) [quest] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 23.7 healing_power points (5.92 DPS) | yes | Earthen Silk Cuffs (254019, -1.39 DPS, sim-verified) [crafted]; Spidertank Oilrag (9448, -1.41 DPS) [dungeon]; Mindthrust Bracers (1974, -2.37 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Earthen Silk Gloves (254017, -0.41 DPS, sim-verified) [crafted]; Mender's Leather Gloves (252530, -1.40 DPS) [crafted]; Gilded Handwraps (254021, -1.66 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 42.0 healing_power points (10.47 DPS) | yes | Mender's Leather Belt (252523, -0.86 DPS, sim-verified) [crafted]; Windchaser Cinch (14435, -4.02 DPS) [world_drop]; Skycaller's Leather Belt (252522, -4.48 DPS) [crafted] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 57.6 healing_power points (14.37 DPS) | yes | Stoneweaver Leggings (9407, -1.19 DPS, sim-verified) [dungeon]; Stormcloth Pants (10010, -3.29 DPS) [crafted]; Wisdom's Leather Leggings (252519, -3.92 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) | Leatherworking [crafted] | 44.9 healing_power points (11.20 DPS) | yes | Furen's Boots (13100, -1.58 DPS) [world_drop]; Thoughtcast Boots (10578, -1.60 DPS) [dungeon]; Gilded Slippers (254001, -2.46 DPS) [crafted] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 22.7 healing_power points (5.67 DPS) | yes | Solid Gold Band (279845, -0.48 DPS) [quest]; Welken Ring (5011, -0.96 DPS) [world_drop]; Darkspear Signet (272070, -1.14 DPS) [vendor] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 22.7 healing_power points (5.67 DPS) | yes | Solid Gold Band (279845, -0.49 DPS, sim-verified) [quest]; Welken Ring (5011, -0.96 DPS) [world_drop]; Darkspear Signet (272070, -1.14 DPS) [vendor] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+8.5 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Jordan (873, -6.37 DPS, sim-verified) [world_drop]; Staff of the Friar (3415, -7.74 DPS) [dungeon]; Odo's Ley Staff (6318, -8.51 DPS) [dungeon] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | 29.9 healing_power points (7.44 DPS) | yes | Orb of Souls (249395, -0.65 DPS, sim-verified) [crafted]; Omega Orb (7749, -2.93 DPS) [quest]; Silksand Star (15964, -3.04 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Cloak of Blight; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Dreamweave Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Hand of Righteousness; off_hand: Beacon of Hope

No-known-source sample (15 of 508, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 50 (tauren, 0200000000000000-00000000000000000000-5553005153113200)

Set DPS (verified): 392.7. Weights run: 28.7s. Verify run: 113.8s. 664 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.006, intellect=0.609 ± 0.018, spirit=1.982 ± 0.020, mp5=2.010 ± 0.021, crit=0.096 ± 0.004 per rating point (14 rating = 1%, 1.342 per %), spell_haste=not significant (-0.043 ± 0.318)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Engineer's Guild Headpiece (9534) | Divino-matic Rod [quest] | 62.6 healing_power points (27.81 DPS) | yes | Blood Guard's Restored Leather Helm (220875, -2.18 DPS) [vendor]; Cassandra's Grace (13102, -3.17 DPS, sim-verified) [world_drop]; Gemburst Circlet (10751, -3.36 DPS) [quest] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 24.9 healing_power points (11.05 DPS) | yes | Triune Amulet (7722, -1.33 DPS) [dungeon]; Amberglow Talisman (10824, -2.24 DPS) [quest]; Mindburst Medallion (11196, -2.36 DPS) [quest] |
| shoulder | Living Shoulders (15061) | Leatherworking [crafted] | 56.8 healing_power points (25.22 DPS) | yes | Mender's Leather Shoulders (252538, -2.61 DPS, sim-verified) [crafted]; Blood Guard's Restored Leather Spaulders (220877, -3.93 DPS) [vendor]; Ironfeather Shoulders (15067, -6.10 DPS) [crafted] |
| back | Battle Healer's Cloak (19527) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -1.93 DPS) [dungeon]; Wingveil Cloak (10802, -3.15 DPS) [dungeon]; Cloak of Blight (6832, -3.61 DPS) [quest] |
| chest | Feathered Breastplate (8349) | Leatherworking [crafted] | sim-verified (392.6 DPS) | yes | Stormcloth Vest (10020, -2.86 DPS, sim-verified) [crafted]; Ghostweave Vest (14141, -3.53 DPS) [crafted]; Grimlok's Tribal Vestments (9415, -5.72 DPS) [dungeon] |
| wrist | Mender's Leather Bracers (252543) | Leatherworking [crafted] | 35.2 healing_power points (15.62 DPS) | yes | Nethergeld Cuffs (254061, -0.44 DPS) [crafted]; Skycaller's Leather Bracers (252542, -4.88 DPS) [crafted]; Enchanted Kodo Bracers (13119, -5.73 DPS) [world_drop] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gilded Gloves (254095, -0.67 DPS, sim-verified) [crafted]; Earthenweave Gloves (254075, -2.90 DPS) [crafted]; Ghostweave Gloves (14142, -6.93 DPS) [crafted] |
| waist | Earthenweave Cord (254077) | Tailoring [crafted] | 47.8 healing_power points (21.24 DPS) | yes | Gilded Cord (254037, -1.84 DPS, sim-verified) [crafted]; Mender's Leather Waistguard (252477, -4.66 DPS) [crafted]; Ghostweave Cord (254073, -5.33 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 64.9 healing_power points (28.85 DPS) | yes | Jinxed Hoodoo Kilt (9474, -2.31 DPS) [dungeon]; Warchief Kilt (7760, -4.24 DPS) [dungeon]; Dalewind Trousers (13008, -4.40 DPS) [world_drop] |
| feet | Earthenweave Boots (254093) | Tailoring [crafted] | 47.8 healing_power points (21.24 DPS) | yes | First Sergeant's Restored Leather Boots (220885, -0.84 DPS) [vendor]; Sandals of the Insurgent (13111, -1.46 DPS) [world_drop]; Mender's Leather Shoes (252533, -2.52 DPS) [crafted] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 27.4 healing_power points (12.19 DPS) | yes | Runemender's Seal (273031, -2.03 DPS) [world]; Solid Gold Band (279845, -2.90 DPS) [quest]; Advisor's Ring (19519, -3.29 DPS) [rep] |
| finger2 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 22.9 healing_power points (10.16 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; Solid Gold Band (279845, -0.87 DPS) [quest]; Advisor's Ring (19519, -1.26 DPS) [rep] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+10.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -2.57 DPS) [world_drop]; Darkspear Voodoo Seal (272061, -6.10 DPS) [vendor]; Alchemist's Stone (13503, -7.04 DPS) [crafted] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -2.15 DPS, sim-verified) [world_drop]; Darkspear Voodoo Seal (272061, -4.34 DPS) [vendor]; Alchemist's Stone (13503, -5.28 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamstaff (249454, -12.46 DPS, sim-verified) [crafted]; Resurgence Rod (17743, -22.83 DPS) [quest]; Charstone Dirk (17710, -26.12 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Engineer's Guild Headpiece; shoulder: Living Shoulders; back: Battle Healer's Cloak; chest: Feathered Breastplate; wrist: Mender's Leather Bracers; hands: Mender's Leather Gauntlets; waist: Earthenweave Cord; legs: Senior Designer's Pantaloons; feet: Earthenweave Boots; finger1: Eye of Adaegus; finger2: Electrocutioner Lagnut; trinket2: Evonice's Landin' Pilla; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 664, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise

### Band 60 (tauren, 0531200100000000-00000000000000000000-5553005153113200)

Set DPS (verified): 624.9. Weights run: 9.6s. Verify run: 344.2s. 1448 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.019, intellect=0.591 ± 0.020, spirit=1.646 ± 0.019, mp5=1.932 ± 0.022, crit=0.107 ± 0.004 per rating point (14 rating = 1%, 1.503 per %), spell_haste=not significant (0.020 ± 0.163)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 106.7 healing_power points (57.22 DPS) | yes | Warlord's Dragonhide Headdress (231675, -15.31 DPS) [vendor]; Mooncloth Circlet (14140, -15.92 DPS, sim-verified) [crafted]; Champion's Dragonhide Headdress (227205, -16.50 DPS) [vendor] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 24.7 healing_power points (13.26 DPS) | yes | Lady Alizabeth's Pendant (13002, -0.02 DPS) [world_drop]; Triune Amulet (7722, -2.79 DPS) [dungeon]; Archlight Talisman (15856, -3.66 DPS) [quest] |
| shoulder | Mooncloth Shoulders (14139) | Tailoring [crafted] | 55.1 healing_power points (29.56 DPS) | yes | Champion's Dragonhide Pauldrons (227207, +0.00 DPS) [vendor]; Warlord's Dragonhide Pauldrons (231672, +0.00 DPS) [vendor]; Living Shoulders (15061, -1.06 DPS, sim-verified) [crafted] |
| back | Battle Healer's Cloak (19526) | Warsong Outriders [rep] | sim-verified (624.8 DPS) | yes | Spritecaster Cape (11623, -6.30 DPS) [dungeon]; Wingveil Cloak (10802, -8.62 DPS) [dungeon]; Cloak of Earth and Sky (279270, -8.65 DPS) [crafted] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (624.8 DPS) | yes | Legionnaire's Dragonhide Tunic (227202, +0.00 DPS) [vendor]; Warlord's Dragonhide Tunic (231674, +0.00 DPS) [vendor]; Ironfeather Breastplate (15066, -13.39 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (624.8 DPS) | yes | Dryad's Wrist Bindings (19595, +0.00 DPS) [rep]; Earthenweave Cuffs (254131, -17.97 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (624.8 DPS) | yes | Wildheart Gloves (16717, +0.00 DPS) [dungeon]; General's Dragonhide Gauntlets (231676, +0.00 DPS) [vendor]; Mender's Leather Gauntlets (252551, -19.47 DPS, sim-verified) [crafted] |
| waist | Elderwild Waistcord (279252) | Leatherworking [crafted] | 59.6 healing_power points (31.98 DPS) | yes | Wisdom of the Timbermaw (19047, -1.40 DPS, sim-verified) [crafted]; Earthenweave Cord (254077, -8.33 DPS) [crafted]; Devout Belt (16696, -11.80 DPS) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-verified (624.8 DPS) | yes | Legionnaire's Dragonhide Legguards (227206, +0.00 DPS) [vendor]; General's Dragonhide Legguards (231673, +0.00 DPS) [vendor]; Elderwild Pants (279254, -17.03 DPS, sim-verified) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (624.8 DPS) | yes | General's Dragonhide Boots (231682, +0.00 DPS) [pvp]; Wildheart Boots (16715, -0.32 DPS) [dungeon]; Mooncloth Boots (15802, -16.21 DPS, sim-verified) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 32.0 healing_power points (17.17 DPS) | yes | Eye of Adaegus (5266, -4.68 DPS) [world_drop]; Advisor's Ring (19518, -5.52 DPS) [rep]; Electrocutioner Lagnut (9447, -6.16 DPS) [dungeon] |
| finger2 | Band of the Hierophant (13096) | World drop [world_drop] | 24.0 healing_power points (12.88 DPS) | yes | Eye of Adaegus (5266, -0.38 DPS) [world_drop]; Advisor's Ring (19518, -1.22 DPS) [rep]; Electrocutioner Lagnut (9447, -1.87 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (624.8 DPS) | yes | Evonice's Landin' Pilla (18951, -1.77 DPS) [quest]; Shard of the Splithooves (10659, -4.38 DPS) [quest]; Darkspear Voodoo Seal (272061, -5.41 DPS) [vendor] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (624.8 DPS) | yes | Evonice's Landin' Pilla (18951, -0.80 DPS, sim-verified) [quest]; Shard of the Splithooves (10659, -3.44 DPS) [quest]; Darkspear Voodoo Seal (272061, -4.47 DPS) [vendor] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (624.8 DPS) | yes | High Warlord's Battle Mace (23464, +0.00 DPS) [vendor]; Hand of Edward the Odd (2243, -10.36 DPS, sim-verified) [world_drop]; Dreamstaff (249454, -24.50 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Idol of Swiftness (279250) | Enchanting [crafted] | sim-verified (624.8 DPS) | yes | Idol of the Moon (23197, -1.13 DPS, sim-verified) [world_drop] |

**New at 60:** head: Living Crown; shoulder: Mooncloth Shoulders; back: Battle Healer's Cloak; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Elderwild Waistcord; legs: Devout Skirt; feet: Devout Sandals; finger1: Maiden's Circle; finger2: Band of the Hierophant; trinket2: Uther's Strength; ranged: Idol of Swiftness

No-known-source sample (15 of 1448, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise

### Band 60, raid preset (tauren, 0531200100000000-00000000000000000000-5553005153113200)

Set DPS (verified): 1061.1. Weights run: 6.9s. Verify run: 176.7s. 1448 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.383, intellect=1.101 ± 0.085, spirit=1.402 ± 0.086, mp5=1.242 ± 0.091, crit=0.667 ± 0.095 per rating point (14 rating = 1%, 9.341 per %), spell_haste=not significant (1.321 ± 1.565)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 112.4 healing_power points (48.67 DPS) | yes | Warlord's Dragonhide Headdress (231675, -12.90 DPS) [vendor]; Champion's Dragonhide Headdress (227205, -13.20 DPS) [vendor]; Mooncloth Circlet (14140, -17.52 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 28.8 healing_power points (12.47 DPS) | yes | Archlight Talisman (15856, -3.32 DPS, sim-verified) [quest]; Lady Alizabeth's Pendant (13002, -3.36 DPS) [world_drop]; Horizon Choker (13085, -3.37 DPS) [world_drop] |
| shoulder | Mooncloth Shoulders (14139) | Tailoring [crafted] | 52.4 healing_power points (22.69 DPS) | yes | Champion's Dragonhide Pauldrons (227207, +0.00 DPS) [vendor]; Warlord's Dragonhide Pauldrons (231672, +0.00 DPS) [vendor]; Wildheart Spaulders (16718, -2.87 DPS, sim-verified) [dungeon] |
| back | Battle Healer's Cloak (19526) | Warsong Outriders [rep] | sim-verified (+4.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Spritecaster Cape (11623, -3.55 DPS) [dungeon]; Darkspear Raider's Cloak (272063, -4.85 DPS) [vendor] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | 74.6 healing_power points (32.31 DPS) | yes | Legionnaire's Dragonhide Tunic (227202, +0.00 DPS) [vendor]; Warlord's Dragonhide Tunic (231674, +0.00 DPS) [vendor]; Ironfeather Breastplate (15066, -0.51 DPS) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 40.6 healing_power points (17.58 DPS) | yes | Mender's Leather Bracers (252543, -2.38 DPS) [crafted]; General's Dragonhide Bracers (16553, -2.46 DPS) [pvp]; Earthenweave Cuffs (254131, -2.91 DPS, sim-verified) [crafted] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | sim-verified (1060.9 DPS) | yes | General's Dragonhide Gauntlets (231676, +0.00 DPS) [vendor]; Raider Handwraps (272097, -0.98 DPS, sim-verified) [vendor]; Gilded Gloves (254095, -1.00 DPS) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 59.0 healing_power points (25.55 DPS) | yes | Elderwild Waistcord (279252, -1.07 DPS, sim-verified) [crafted]; Devout Belt (16696, -5.80 DPS) [dungeon]; Wildheart Belt (16716, -6.62 DPS) [dungeon] |
| legs | Elderwild Pants (279254) | Leatherworking [crafted] | 73.6 healing_power points (31.88 DPS) | yes | Legionnaire's Dragonhide Legguards (227206, +0.00 DPS) [vendor]; General's Dragonhide Legguards (231673, +0.00 DPS) [vendor]; Devout Skirt (16694, -1.77 DPS, sim-verified) [dungeon] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 57.6 healing_power points (24.95 DPS) | yes | General's Dragonhide Boots (231682, -1.91 DPS) [pvp]; Devout Sandals (16691, -5.03 DPS, sim-verified) [dungeon]; First Sergeant's Restored Leather Boots (220885, -5.46 DPS) [vendor] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 34.1 healing_power points (14.77 DPS) | yes | Eye of Adaegus (5266, -4.63 DPS) [world_drop]; Deep Rooted Ring (19109, -4.81 DPS) [world]; Advisor's Ring (19518, -6.56 DPS) [rep] |
| finger2 | Band of the Hierophant (13096) | World drop [world_drop] | 26.4 healing_power points (11.44 DPS) | yes | Deep Rooted Ring (19109, -1.48 DPS) [world]; Eye of Adaegus (5266, -1.53 DPS, sim-verified) [world_drop]; Advisor's Ring (19518, -3.23 DPS) [rep] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Uther's Strength (11302, -5.41 DPS, sim-verified) [world_drop] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -1.21 DPS) [quest]; Uther's Strength (11302, -1.72 DPS, sim-verified) [world_drop]; Shard of the Splithooves (10659, -4.06 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Mace (23464, +0.00 DPS) [vendor]; Hand of Edward the Odd (2243, -13.68 DPS, sim-verified) [world_drop]; Dreamstaff (249454, -27.22 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Idol of Swiftness (279250) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Idol of the Moon (23197, -9.32 DPS, sim-verified) [world_drop] |

**New at 60:** head: Living Crown; shoulder: Mooncloth Shoulders; back: Battle Healer's Cloak; chest: Mooncloth Vest; wrist: Dryad's Wrist Bindings; waist: Wisdom of the Timbermaw; legs: Elderwild Pants; feet: Mooncloth Boots; finger1: Maiden's Circle; finger2: Band of the Hierophant; trinket1: Dormant Heart of the Mountain; trinket2: Ankh of Life; ranged: Idol of Swiftness

No-known-source sample (15 of 1448, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise

