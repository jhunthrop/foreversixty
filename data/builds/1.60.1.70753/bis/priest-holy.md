# Leveling BiS: Holy

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-03503000000000000-00000000000000000)

Set DPS (verified): 37.7. Weights run: 7.3s. Verify run: 25.3s. 190 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.008, intellect=0.895 ± 0.009, spirit=0.699 ± 0.006, mp5=1.741 ± 0.018, crit=0.073 ± 0.003 per rating point (14 rating = 1%, 1.027 per %), spell_haste=not significant (0.148 ± 0.052)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.46 DPS) | yes | Shadow Goggles (4373, -0.48 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.23 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 1.4 healing_power points (0.11 DPS) | yes | Tarnished Locket (279870, -0.13 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.1 healing_power points (0.65 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS, sim-verified) [dungeon]; Reinforced Woolen Shoulders (4315, -0.36 DPS) [crafted] |
| back | Sanguine Cape (14376) | World drop [world_drop] | 3.6 healing_power points (0.29 DPS) | yes | Seer's Cape (6378, +0.00 DPS, sim-verified) [dungeon]; Spirit Cloak (4792, -0.06 DPS) [vendor]; Sylvan Cloak (4793, -0.06 DPS) [vendor] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 19.6 healing_power points (1.59 DPS) | yes | Corsair's Overshirt (5202, -0.24 DPS) [dungeon]; Robe of the Moccasin (6465, -0.29 DPS) [dungeon]; Bloody Apron (6226, -0.53 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 9.7 healing_power points (0.79 DPS) | yes | Bright Bracers (3647, -0.48 DPS, sim-verified) [world_drop]; Repurposed Hair Band (281256, -0.53 DPS) [quest]; Seer's Cuffs (3645, -0.60 DPS) [dungeon] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 13.7 healing_power points (1.11 DPS) | yes | Magefist Gloves (12977, -0.34 DPS, sim-verified) [world_drop]; Bright Gloves (3066, -0.59 DPS) [world_drop]; Evergreen Gloves (7738, -0.78 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 14.6 healing_power points (1.18 DPS) | yes | Novice Ardent's Sash (253887, -0.16 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.60 DPS) [world_drop]; Relic Hunter Belt (11936, -0.78 DPS) [quest] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 26.2 healing_power points (2.13 DPS) | yes | Darkweave Breeches (12987, -1.09 DPS, sim-verified) [world_drop]; Filigreed Silky Leggings (253939, -1.46 DPS) [crafted]; Filigreed Flame Leggings (253941, -1.46 DPS) [crafted] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 11.7 healing_power points (0.95 DPS) | yes | Kimbra Boots (6191, -0.15 DPS, sim-verified) [quest]; Bluegill Sandals (1560, -0.58 DPS) [world]; Smoldering Boots (3076, -0.61 DPS) [world] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.4 healing_power points (0.44 DPS) | yes | Volcanic Rock Ring (12053, -0.22 DPS) [world_drop]; Furlfeather Ring (286745, -0.27 DPS) [world]; Minor Channeling Ring (1449, -0.29 DPS) [quest] |
| finger2 | Band of Purification (12996) | World drop [world_drop] | 4.2 healing_power points (0.34 DPS) | yes | Volcanic Rock Ring (12053, +0.00 DPS, sim-verified) [world_drop]; Furlfeather Ring (286745, -0.18 DPS) [world]; Minor Channeling Ring (1449, -0.20 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Westfall (2042) | The Defias Brotherhood [quest] | sim-verified (37.7 DPS) | yes | Staff of the Blessed Seer (2271, -0.14 DPS) [dungeon]; Lesser Staff of the Spire (1300, -0.27 DPS) [world]; Twisted Chanter's Staff (890, -5.18 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Flaring Baton (5326) (or Moonstone Wand (15204)) | The Escape [quest] | 1.8 healing_power points (0.15 DPS) | yes | Moonstone Wand (15204, +0.00 DPS) [quest]; Dwarven Flamestick (5241, -0.03 DPS) [quest]; Sable Wand (7607, -0.03 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Band of Purification; main_hand: Staff of Westfall; ranged: Flaring Baton

No-known-source sample (15 of 190, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes; 16607 Acolyte's Sacrificial Robes; 18851 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (gnome, 000000000000000000-03505003030110000-00000000000000000)

Set DPS (verified): 56.1. Weights run: 7.7s. Verify run: 28.2s. 321 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.017, intellect=1.066 ± 0.015, spirit=0.778 ± 0.010, mp5=2.083 ± 0.015, crit=0.113 ± 0.006 per rating point (14 rating = 1%, 1.588 per %), spell_haste=0.890 ± 0.106

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 37.7 healing_power points (3.03 DPS) | yes | Filigreed Pristine Circlet (253975, -1.50 DPS, sim-verified) [crafted]; Pristine Circlet (253949, -1.58 DPS) [crafted]; Nightsky Cowl (4039, -1.69 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 7.4 healing_power points (0.59 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.08 DPS) [world]; Crystal Starfire Medallion (5003, -0.13 DPS, sim-verified) [world_drop]; Pendant of Myzrael (4614, -0.22 DPS) [dungeon] |
| shoulder | Mantle of Honor (3560) | Bride of the Embalmer [quest] | 12.9 healing_power points (1.04 DPS) | yes | Nightsky Mantle (4718, -0.04 DPS) [world_drop]; Death Speaker Mantle (6685, -0.10 DPS) [dungeon]; Faerie Mantle (5820, -0.15 DPS) [quest] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 13.2 healing_power points (1.06 DPS) | yes | Repairman's Cape (9605, -0.01 DPS) [quest]; Prelacy Cape (7004, -0.31 DPS) [quest]; Sergeant Major's Cape (16315, -0.47 DPS) [pvp] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 31.7 healing_power points (2.55 DPS) | yes | Pristine Gown (253961, -0.11 DPS, sim-verified) [crafted]; Robes of Arugal (6324, -0.73 DPS) [dungeon]; Filigreed Pristine Gown (253901, -0.89 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 11.6 healing_power points (0.93 DPS) | yes | Glowing Magical Bracelets (13106, -0.25 DPS) [world_drop]; Nightsky Wristbands (6407, -0.38 DPS, sim-verified) [world_drop]; Spidertank Oilrag (9448, -0.43 DPS) [dungeon] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 18.2 healing_power points (1.46 DPS) | yes | Gloves of Old (9395, -0.12 DPS, sim-verified) [world_drop]; Pristine Gloves (253913, -0.32 DPS) [crafted]; Hotshot Pilot's Gloves (9491, -0.47 DPS) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 15.3 healing_power points (1.23 DPS) | yes | Novice Ardent's Sash (253887, -0.46 DPS, sim-verified) [crafted]; Resilient Cord (14406, -0.46 DPS) [world_drop]; Dreamer's Belt (4829, -0.50 DPS) [vendor] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 33.3 healing_power points (2.68 DPS) | yes | Filigreed Pristine Leggings (253937, -0.66 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -0.79 DPS) [crafted]; Darkweave Breeches (12987, -1.71 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 30.6 healing_power points (2.46 DPS) | yes | Acidic Walkers (9454, -1.52 DPS) [dungeon]; Frothing Slippers (254003, -1.61 DPS) [crafted]; Pristine Boots (253889, -1.83 DPS, sim-verified) [crafted] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 14.4 healing_power points (1.16 DPS) | yes | The Queen's Jewel (13094, -0.49 DPS) [world_drop]; Defias Mage Ring (1077, -0.63 DPS) [world]; Seal of Wrynn (2933, -0.63 DPS) [quest] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 14.4 healing_power points (1.16 DPS) | yes | The Queen's Jewel (13094, -0.59 DPS, sim-verified) [world_drop]; Defias Mage Ring (1077, -0.63 DPS) [world]; Seal of Wrynn (2933, -0.63 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Friar (3415) | Blackfathom Deeps [dungeon] | sim-verified (56.1 DPS) | yes | Glimmering Staff (249392, -0.06 DPS) [crafted]; Gnarled Ash Staff (791, -0.07 DPS) [world_drop]; Death Speaker Scepter (2816, -5.06 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 8.4 healing_power points (0.68 DPS) | yes | Lesser Mystic Wand (11289, -0.34 DPS) [crafted]; Starfaller (13063, -0.34 DPS) [world_drop]; Summoner's Wand (5245, -0.51 DPS, sim-verified) [world] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Mantle of Honor; back: Glowing Thresher Cape; chest: Death Speaker Robes; hands: Truefaith Gloves; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Staff of the Friar; ranged: Wand of Mana Concentration

No-known-source sample (15 of 321, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 40 (gnome, 000000000000000000-03505003030121431-00000000000000000)

Set DPS (verified): 107.2. Weights run: 8.1s. Verify run: 25.8s. 422 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.082, intellect=0.988 ± 0.031, spirit=1.558 ± 0.030, mp5=3.240 ± 0.033, crit=0.216 ± 0.010 per rating point (14 rating = 1%, 3.020 per %), spell_haste=-5.367 ± 0.507

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 52.8 healing_power points (4.26 DPS) | yes | Holy Shroud (2721, -0.84 DPS) [world_drop]; Whitemane's Chapeau (7720, -0.88 DPS, sim-verified) [dungeon]; Earthen Silk Hood (254015, -1.23 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 27.9 healing_power points (2.25 DPS) | yes | Amberglow Talisman (10824, -0.99 DPS) [quest]; Explorers' League Commendation (7746, -1.12 DPS) [quest]; Necklace of Calisea (1714, -2.02 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 34.8 healing_power points (2.81 DPS) | yes | Crimson Silk Shoulders (7059, -1.29 DPS) [crafted]; Mistscape Mantle (4734, -1.31 DPS) [dungeon]; Earthen Silk Shoulders (254033, -1.38 DPS, sim-verified) [crafted] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 19.5 healing_power points (1.57 DPS) | yes | Ceremonial Centaur Blanket (6789, -0.28 DPS) [quest]; Repairman's Cape (9605, -0.30 DPS, sim-verified) [quest]; Energy Cloak (9397, -0.31 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 70.5 healing_power points (5.69 DPS) | yes | Doomsayer's Robe (4746, -2.13 DPS) [quest]; Pristine Gown (253961, -3.01 DPS) [crafted]; Civinad Robes (9623, -3.80 DPS, sim-verified) [quest] |
| wrist | Earthen Silk Cuffs (254019) | Tailoring [crafted] | 17.8 healing_power points (1.44 DPS) | yes | Silkstream Cuffs (16791, -0.53 DPS) [quest]; Nightsky Wristbands (6407, -0.58 DPS) [world_drop]; Mindthrust Bracers (1974, -0.61 DPS, sim-verified) [dungeon] |
| hands | Earthen Silk Gloves (254017) | Tailoring [crafted] | 30.5 healing_power points (2.46 DPS) | yes | Gilded Handwraps (254021, -0.25 DPS, sim-verified) [crafted]; Gloves of Old (9395, -0.55 DPS) [world_drop]; Restorer's Fine Gloves (270059, -0.56 DPS) [quest] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 39.3 healing_power points (3.17 DPS) | yes | Earthen Silk Belt (7061, -1.52 DPS) [crafted]; Sutarn's Ring (13105, -1.62 DPS) [world_drop]; Windchaser Cinch (14435, -2.64 DPS, sim-verified) [world_drop] |
| legs | Stoneweaver Leggings (9407) | Uldaman: Ironaya [dungeon] | 46.6 healing_power points (3.76 DPS) | yes | Pristine Leggings (253987, -0.80 DPS) [crafted]; Earthen Leggings (253999, -1.17 DPS) [crafted]; Stormcloth Pants (10010, -1.47 DPS, sim-verified) [crafted] |
| feet | Thoughtcast Boots (10578) | Razorfen Downs [dungeon] | 33.3 healing_power points (2.68 DPS) | yes | Gilded Slippers (254001, -0.01 DPS) [crafted]; Furen's Boots (13100, -0.11 DPS) [world_drop]; Everlast Boots (10359, -0.75 DPS) [quest] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 19.9 healing_power points (1.61 DPS) | yes | Welken Ring (5011, -0.33 DPS) [world_drop]; The Queen's Jewel (13094, -0.44 DPS) [world_drop]; Blush Ember Ring (13093, -0.60 DPS) [world_drop] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 19.9 healing_power points (1.61 DPS) | yes | The Queen's Jewel (13094, -0.44 DPS) [world_drop]; Welken Ring (5011, -0.50 DPS, sim-verified) [world_drop]; Blush Ember Ring (13093, -0.60 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (107.1 DPS) | yes | Gnarled Ash Staff (791, -0.97 DPS) [world_drop]; Staff of the Friar (3415, -1.10 DPS) [dungeon]; Death Speaker Scepter (2816, -3.99 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Captain Rackmore's Tiller (16789) | Claim Rackmore's Treasure! [quest] | 10.7 healing_power points (0.86 DPS) | yes | Jaina's Firestarter (13064, -0.01 DPS) [world_drop]; Wand of Mana Concentration (273043, -0.02 DPS) [world]; Earthen Rod (9381, -0.14 DPS) [dungeon] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; chest: Stormcloth Vest; wrist: Earthen Silk Cuffs; hands: Earthen Silk Gloves; waist: Gilded Cord; legs: Stoneweaver Leggings; feet: Thoughtcast Boots; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Captain Rackmore's Tiller

No-known-source sample (15 of 422, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 50 (gnome, 025003000000000000-03505003030121431-00000000000000000)

Set DPS (verified): 180.9. Weights run: 7.9s. Verify run: 60.2s. 541 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.046, intellect=1.116 ± 0.032, spirit=2.216 ± 0.037, mp5=3.967 ± 0.037, crit=0.268 ± 0.013 per rating point (14 rating = 1%, 3.752 per %), spell_haste=-5.547 ± 0.626

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | sim-verified (180.9 DPS) | yes | Cassandra's Grace (13102, -0.37 DPS) [world_drop]; Knight-Lieutenant's Satin Cover (220896, -0.43 DPS) [vendor]; Papal Fez (9431, -0.66 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | sim-verified (180.9 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS) [dungeon]; Necklace of Calisea (1714, -0.11 DPS) [dungeon]; Amberglow Talisman (10824, -0.22 DPS) [quest] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | sim-verified (180.9 DPS) | yes | Knight-Lieutenant's Satin Pads (220894, +0.00 DPS) [vendor]; Inquisitor's Shawl (19507, -1.11 DPS) [dungeon]; Earthen Silk Shoulders (254033, -1.18 DPS) [crafted] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | sim-verified (180.9 DPS) | yes | Arcane Cloak (8286, +0.00 DPS) [world_drop]; Chainlink Towel (9648, +0.00 DPS) [quest]; Wingveil Cloak (10802, +0.00 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ghostweave Vest (14141, +0.00 DPS) [crafted]; Grimlok's Tribal Vestments (9415, -0.96 DPS) [dungeon]; Doomsayer's Robe (4746, -1.87 DPS) [quest] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 38.9 healing_power points (3.67 DPS) | yes | Venomshroud Armguards (14439, -1.58 DPS) [world_drop]; Earthen Silk Cuffs (254019, -1.68 DPS) [crafted]; Shizzle's Nozzle Wiper (11917, -1.78 DPS) [quest] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 57.9 healing_power points (5.46 DPS) | yes | Earthenweave Gloves (254075, -0.71 DPS) [crafted]; Earthen Silk Gloves (254017, -2.09 DPS) [crafted]; Sergeant Major's Satin Gloves (220897, -2.12 DPS) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | sim-verified (180.9 DPS) | yes | Gilded Cord (254037, +0.00 DPS) [crafted]; Earthenweave Cord (254077, +0.00 DPS) [crafted]; Gilded Waistcord (254081, +0.00 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | sim-verified (180.9 DPS) | yes | Senior Designer's Pantaloons (11841, +0.00 DPS) [dungeon]; Dalewind Trousers (13008, -0.05 DPS) [world_drop]; Stoneweaver Leggings (9407, -0.99 DPS) [dungeon] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | sim-verified (180.9 DPS) | yes | Furen's Boots (13100, +0.00 DPS) [world_drop]; Vinerot Sandals (17748, +0.00 DPS) [dungeon]; Earthenweave Boots (254093, +0.00 DPS) [crafted] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 33.3 healing_power points (3.14 DPS) | yes | Choking Band (11868, -0.42 DPS) [quest]; Electrocutioner Lagnut (9447, -0.83 DPS) [dungeon]; Runemender's Seal (273031, -0.83 DPS) [world] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | sim-verified (180.9 DPS) | yes | Electrocutioner Lagnut (9447, +0.00 DPS) [dungeon]; Choking Band (11868, +0.00 DPS) [quest]; Runemender's Seal (273031, +0.00 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (180.9 DPS) | yes | Evonice's Landin' Pilla (18951, -0.42 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.84 DPS) [quest]; Alchemist's Stone (13503, -1.67 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (180.9 DPS) | yes | Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.03 DPS) [quest]; Alchemist's Stone (13503, -0.86 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (180.9 DPS) | yes | Spellshifter Rod (9527, +0.00 DPS) [quest]; Spire of Hakkar (10844, +0.00 DPS) [world]; Charstone Dirk (17710, +0.00 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cairnstone Sliver (9654) | The Morrow Stone [quest] | sim-verified (180.9 DPS) | yes | Jaina's Firestarter (13064, -0.31 DPS) [world_drop]; Captain Rackmore's Tiller (16789, -0.38 DPS) [quest]; Wand of Mana Concentration (273043, -0.40 DPS) [world] |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Nethergeld Shoulders; back: Imperial Red Cloak; wrist: Nethergeld Cuffs; hands: Gilded Gloves; waist: Dawnspire Cord; legs: Kilt of the Atal'ai Prophet; feet: Gilded Sandals; finger1: Eye of Adaegus; finger2: Welken Ring; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Cairnstone Sliver

No-known-source sample (15 of 541, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60 (gnome, 025003031330000000-03505003030121431-00000000000000000)

Set DPS (verified): 310.2. Weights run: 6.8s. Verify run: 32.2s. 1226 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.055, intellect=0.871 ± 0.035, spirit=1.415 ± 0.033, mp5=2.375 ± 0.043, crit=0.205 ± 0.014 per rating point (14 rating = 1%, 2.873 per %), spell_haste=-25.565 ± 0.929

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (310.2 DPS) | yes | Mooncloth Circlet (14140, +0.00 DPS) [crafted]; Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor] |
| neck | Lady Alizabeth's Pendant (13002) | World drop [world_drop] | sim-verified (310.2 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS) [dungeon]; Horizon Choker (13085, -0.63 DPS) [world_drop]; Necklace of Calisea (1714, -0.97 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (310.2 DPS) | yes | Mooncloth Shoulders (14139, +0.00 DPS) [crafted]; Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp] |
| back | Wingveil Cloak (10802) | Sunken Temple: Spawn of Hakkar [dungeon] | 20.6 healing_power points (3.82 DPS) | yes | Ethereal Mist Cape (11873, -0.18 DPS) [quest]; Indomitable Cloak (14683, -0.40 DPS) [world_drop]; Glowing Thresher Cape (6901, -0.42 DPS) [dungeon] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 95.6 healing_power points (17.78 DPS) | yes | Field Marshal's Satin Tunic (231624, -3.25 DPS) [vendor]; Knight-Captain's Satin Tunic (227122, -4.23 DPS) [pvp]; Mooncloth Vest (14138, -4.71 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (310.2 DPS) | yes | Nethergeld Cuffs (254061, +0.00 DPS) [crafted]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted]; Marshal's Satin Bracers (17606, -0.26 DPS) [pvp] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (310.2 DPS) | yes | Gilded Gloves (254095, +0.00 DPS) [crafted]; Knight-Lieutenant's Satin Handwraps (227128, -0.63 DPS) [pvp]; Swarmtender's Gloves (275607, -0.64 DPS) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 55.9 healing_power points (10.40 DPS) | yes | Earthenweave Cord (254077, -2.67 DPS) [crafted]; Marshal's Satin Sash (17609, -2.73 DPS) [pvp]; Devout Belt (16696, -2.75 DPS) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 69.6 healing_power points (12.94 DPS) | yes | Knight-Captain's Satin Legguards (227125, +0.00 DPS) [pvp]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Mooncloth Leggings (14137, -1.65 DPS) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (310.2 DPS) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; Marshal's Satin Walkers (231627, +0.00 DPS) [vendor]; Swarmtender's Footpads (275609, -0.14 DPS) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 24.3 healing_power points (4.51 DPS) | yes | Electrocutioner Lagnut (9447, -1.00 DPS) [dungeon]; Runemender's Seal (273031, -1.00 DPS) [world]; Choking Band (11868, -1.09 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 22.2 healing_power points (4.13 DPS) | yes | Electrocutioner Lagnut (9447, -0.61 DPS) [dungeon]; Runemender's Seal (273031, -0.61 DPS) [world]; Choking Band (11868, -0.71 DPS) [quest] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.02 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.55 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -0.19 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.72 DPS) [quest]; Thunderbrew's Boot Flask (744, -1.24 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (310.2 DPS) | yes | Spellshifter Rod (9527, +0.00 DPS) [quest]; Spire of Hakkar (10844, +0.00 DPS) [world]; Charstone Dirk (17710, +0.00 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 32.6 healing_power points (6.05 DPS) | yes | Cairnstone Sliver (9654, -3.93 DPS) [quest]; Captain Rackmore's Tiller (16789, -4.15 DPS) [quest]; Jaina's Firestarter (13064, -4.29 DPS) [world_drop] |

**New at 60:** head: Devout Crown; neck: Lady Alizabeth's Pendant; shoulder: Devout Mantle; back: Wingveil Cloak; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Shard of the Splithooves; ranged: Torch of Light

No-known-source sample (15 of 1226, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60, raid preset (gnome, 025003031330000000-03505003030121431-00000000000000000)

Set DPS (verified): 641.2. Weights run: 4.7s. Verify run: 21.5s. 1226 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.156, intellect=0.592 ± 0.030, spirit=0.861 ± 0.030, mp5=1.616 ± 0.032, crit=0.334 ± 0.021 per rating point (14 rating = 1%, 4.681 per %), spell_haste=-3.179 ± 0.613

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (641.2 DPS) | yes | Mooncloth Circlet (14140, +0.00 DPS) [crafted]; Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 24.7 healing_power points (8.05 DPS) | yes | Lady Alizabeth's Pendant (13002, -3.85 DPS) [world_drop]; Horizon Choker (13085, -4.23 DPS) [world_drop]; Necklace of Calisea (1714, -4.74 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (641.2 DPS) | yes | Mooncloth Shoulders (14139, +0.00 DPS) [crafted]; Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp] |
| back | Prelacy Cape (7004) | Researching the Corruption [quest] | sim-verified (641.2 DPS) | yes | Glowing Thresher Cape (6901, +0.00 DPS) [dungeon]; Wingveil Cloak (10802, +0.00 DPS) [dungeon]; Ethereal Mist Cape (11873, +0.00 DPS) [quest] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 86.8 healing_power points (28.25 DPS) | yes | Field Marshal's Satin Tunic (231624, -6.98 DPS) [vendor]; Stormcloth Vest (10020, -7.13 DPS) [crafted]; Knight-Captain's Satin Tunic (227122, -8.41 DPS) [pvp] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (641.2 DPS) | yes | Nethergeld Cuffs (254061, +0.00 DPS) [crafted]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted]; Marshal's Satin Bracers (17606, -1.46 DPS) [pvp] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (641.2 DPS) | yes | Knight-Lieutenant's Satin Handwraps (227128, +0.00 DPS) [pvp]; Marshal's Satin Handwraps (231623, +0.00 DPS) [vendor]; Gilded Gloves (254095, +0.00 DPS) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 47.0 healing_power points (15.31 DPS) | yes | Earthenweave Cord (254077, -3.77 DPS) [crafted]; Gilded Waistcord (254081, -4.14 DPS) [crafted]; Gilded Cord (254037, -4.93 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 52.7 healing_power points (17.15 DPS) | yes | Knight-Captain's Satin Legguards (227125, +0.00 DPS) [pvp]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Earthenweave Leggings (254125, -0.62 DPS) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (641.2 DPS) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; Marshal's Satin Walkers (231627, +0.00 DPS) [vendor]; Swarmtender's Footpads (275609, +0.00 DPS) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 15.4 healing_power points (5.01 DPS) | yes | Electrocutioner Lagnut (9447, -0.12 DPS) [dungeon]; Runemender's Seal (273031, -0.12 DPS) [world]; Choking Band (11868, -1.37 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | sim-verified (641.2 DPS) | yes | Electrocutioner Lagnut (9447, +0.00 DPS) [dungeon]; Runemender's Seal (273031, +0.00 DPS) [world]; Choking Band (11868, -0.88 DPS) [quest] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.35 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.92 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -2.50 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -3.06 DPS) [quest]; Thunderbrew's Boot Flask (744, -3.62 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (641.2 DPS) | yes | Spire of Hakkar (10844, +0.00 DPS) [world]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Charstone Dirk (17710, +0.00 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 26.5 healing_power points (8.62 DPS) | yes | Earthen Rod (9381, -5.69 DPS) [dungeon]; Captain Rackmore's Tiller (16789, -5.82 DPS) [quest]; Cairnstone Sliver (9654, -6.25 DPS) [quest] |

**New at 60:** head: Devout Crown; neck: Glowing Eye of Mordresh; shoulder: Devout Mantle; back: Prelacy Cape; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Shard of the Splithooves; ranged: Torch of Light

No-known-source sample (15 of 1226, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

## Horde

### Band 20 (troll, 000000000000000000-03503000000000000-00000000000000000)

Set DPS (verified): 34.2. Weights run: 7.3s. Verify run: 25.5s. 183 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.008, intellect=0.895 ± 0.009, spirit=0.699 ± 0.006, mp5=1.741 ± 0.018, crit=0.073 ± 0.003 per rating point (14 rating = 1%, 1.027 per %), spell_haste=not significant (0.148 ± 0.052)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.46 DPS) | yes | Shadow Goggles (4373, -0.88 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.23 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 2.1 healing_power points (0.17 DPS) | yes | Scholarly Pendant (277203, -0.06 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.1 healing_power points (0.65 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS, sim-verified) [dungeon]; Reinforced Woolen Shoulders (4315, -0.36 DPS) [crafted] |
| back | Sanguine Cape (14376) | World drop [world_drop] | 3.6 healing_power points (0.29 DPS) | yes | Spirit Cloak (4792, -0.06 DPS) [vendor]; Antiquated Cloak (4799, -0.06 DPS) [vendor]; Seer's Cape (6378, -0.14 DPS, sim-verified) [dungeon] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 19.6 healing_power points (1.59 DPS) | yes | Robe of the Moccasin (6465, -0.29 DPS) [dungeon]; Corsair's Overshirt (5202, -0.42 DPS, sim-verified) [dungeon]; Bloody Apron (6226, -0.53 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 9.7 healing_power points (0.79 DPS) | yes | Bright Bracers (3647, -0.56 DPS, sim-verified) [world_drop]; Featherbead Bracers (15452, -0.57 DPS) [quest]; Seer's Cuffs (3645, -0.60 DPS) [dungeon] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 13.7 healing_power points (1.11 DPS) | yes | Magefist Gloves (12977, -0.58 DPS, sim-verified) [world_drop]; Bright Gloves (3066, -0.59 DPS) [world_drop]; Smoldering Gloves (3074, -0.83 DPS) [world] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 14.6 healing_power points (1.18 DPS) | yes | Novice Ardent's Sash (253887, -0.56 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.60 DPS) [world_drop]; Wise Man's Belt (4786, -0.82 DPS) [vendor] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 26.2 healing_power points (2.13 DPS) | yes | Filigreed Silky Leggings (253939, -1.46 DPS) [crafted]; Filigreed Flame Leggings (253941, -1.46 DPS) [crafted]; Darkweave Breeches (12987, -1.63 DPS, sim-verified) [world_drop] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 11.7 healing_power points (0.95 DPS) | yes | Smoldering Boots (3076, -0.61 DPS) [world]; Sanguine Sandals (14374, -0.66 DPS) [world_drop]; Bluegill Sandals (1560, -0.68 DPS, sim-verified) [world] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.4 healing_power points (0.44 DPS) | yes | Band of Purification (12996, -0.10 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.22 DPS) [world_drop]; Furlfeather Ring (286745, -0.27 DPS) [world] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 4.5 healing_power points (0.36 DPS) | yes | Band of Purification (12996, -0.09 DPS, sim-verified) [world_drop]; Volcanic Rock Ring (12053, -0.15 DPS) [world_drop]; Furlfeather Ring (286745, -0.20 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) | The Wrath of Rath'mael [quest] | sim-verified (34.2 DPS) | yes | Staff of the Blessed Seer (2271, +0.00 DPS) [dungeon]; Crescent Staff (6505, -0.06 DPS) [quest]; Twisted Chanter's Staff (890, -2.73 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Decay (5252) (or Flaring Baton (5326)) | Beren's Peril [quest] | 1.8 healing_power points (0.15 DPS) | yes | Flaring Baton (5326, +0.00 DPS) [quest]; Wisesight Wand (286750, -0.07 DPS) [world] |

**New at 20:** head: Pristine Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Gnarled Necromancer's Staff; ranged: Wand of Decay

No-known-source sample (15 of 183, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (troll, 000000000000000000-03505003030110000-00000000000000000)

Set DPS (verified): 53.9. Weights run: 7.7s. Verify run: 28.1s. 313 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.017, intellect=1.066 ± 0.015, spirit=0.778 ± 0.010, mp5=2.083 ± 0.015, crit=0.113 ± 0.006 per rating point (14 rating = 1%, 1.588 per %), spell_haste=0.890 ± 0.106

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 37.7 healing_power points (3.03 DPS) | yes | Filigreed Pristine Circlet (253975, -1.33 DPS, sim-verified) [crafted]; Pristine Circlet (253949, -1.58 DPS) [crafted]; Nightsky Cowl (4039, -1.69 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 7.4 healing_power points (0.59 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.08 DPS) [world]; Crystal Starfire Medallion (5003, -0.12 DPS, sim-verified) [world_drop]; Pendant of Myzrael (4614, -0.22 DPS) [dungeon] |
| shoulder | Nightsky Mantle (4718) | World drop [world_drop] | 12.4 healing_power points (1.00 DPS) | yes | Death Speaker Mantle (6685, +0.00 DPS, sim-verified) [dungeon]; Ghostly Mantle (3324, -0.14 DPS) [quest]; Resilient Mantle (14397, -0.23 DPS) [world_drop] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 13.2 healing_power points (1.06 DPS) | yes | Resilient Cape (14400, -0.20 DPS, sim-verified) [world_drop]; Silver-thread Cloak (4713, -0.60 DPS) [world_drop]; Amy's Blanket (13005, -0.63 DPS) [world_drop] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 31.7 healing_power points (2.55 DPS) | yes | Pristine Gown (253961, -0.06 DPS, sim-verified) [crafted]; Robes of Arugal (6324, -0.73 DPS) [dungeon]; Filigreed Pristine Gown (253901, -0.89 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 11.6 healing_power points (0.93 DPS) | yes | Nightsky Wristbands (6407, -0.07 DPS, sim-verified) [world_drop]; Glowing Magical Bracelets (13106, -0.25 DPS) [world_drop]; Spidertank Oilrag (9448, -0.43 DPS) [dungeon] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 18.2 healing_power points (1.46 DPS) | yes | Gloves of Old (9395, -0.23 DPS, sim-verified) [world_drop]; Pristine Gloves (253913, -0.32 DPS) [crafted]; Hotshot Pilot's Gloves (9491, -0.47 DPS) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 15.3 healing_power points (1.23 DPS) | yes | Lilac Sash (6780, -0.22 DPS, sim-verified) [quest]; Novice Ardent's Sash (253887, -0.30 DPS) [crafted]; Resilient Cord (14406, -0.46 DPS) [world_drop] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 33.3 healing_power points (2.68 DPS) | yes | Filigreed Pristine Leggings (253937, -0.36 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -0.79 DPS) [crafted]; Sacred Burial Trousers (6282, -1.69 DPS) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 30.6 healing_power points (2.46 DPS) | yes | Pristine Boots (253889, -1.50 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.52 DPS) [dungeon]; Frothing Slippers (254003, -1.61 DPS) [crafted] |
| finger1 | Solid Gold Band (279845) | Opportunistic Education [quest] | 14.9 healing_power points (1.20 DPS) | yes | Runemender's Seal (273031, -0.04 DPS) [world]; The Queen's Jewel (13094, -0.53 DPS) [world_drop]; Defias Mage Ring (1077, -0.67 DPS) [world] |
| finger2 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 14.4 healing_power points (1.16 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; The Queen's Jewel (13094, -0.49 DPS) [world_drop]; Defias Mage Ring (1077, -0.63 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Friar (3415) | Blackfathom Deeps [dungeon] | sim-verified (53.9 DPS) | yes | Glimmering Staff (249392, -0.06 DPS) [crafted]; Gnarled Ash Staff (791, -0.07 DPS) [world_drop]; Death Speaker Scepter (2816, -4.75 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 8.4 healing_power points (0.68 DPS) | yes | Summoner's Wand (5245, -0.09 DPS, sim-verified) [world]; Lesser Mystic Wand (11289, -0.34 DPS) [crafted]; Starfaller (13063, -0.34 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Nightsky Mantle; back: Glowing Thresher Cape; chest: Death Speaker Robes; hands: Truefaith Gloves; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Solid Gold Band; finger2: Electrocutioner Lagnut; main_hand: Staff of the Friar; ranged: Wand of Mana Concentration

No-known-source sample (15 of 313, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (troll, 000000000000000000-03505003030121431-00000000000000000)

Set DPS (verified): 104.8. Weights run: 8.1s. Verify run: 25.8s. 414 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.082, intellect=0.988 ± 0.031, spirit=1.558 ± 0.030, mp5=3.240 ± 0.033, crit=0.216 ± 0.010 per rating point (14 rating = 1%, 3.020 per %), spell_haste=-5.367 ± 0.507

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 52.8 healing_power points (4.26 DPS) | yes | Whitemane's Chapeau (7720, -0.61 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -0.84 DPS) [world_drop]; Earthen Silk Hood (254015, -1.23 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 27.9 healing_power points (2.25 DPS) | yes | Amberglow Talisman (10824, -0.99 DPS) [quest]; Triune Amulet (7722, -1.37 DPS) [dungeon]; Necklace of Calisea (1714, -1.77 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 34.8 healing_power points (2.81 DPS) | yes | Earthen Silk Shoulders (254033, -1.13 DPS, sim-verified) [crafted]; Naga Priestess's Mantle (279889, -1.25 DPS) [quest]; Crimson Silk Shoulders (7059, -1.29 DPS) [crafted] |
| back | Cloak of Blight (6832) | Nothing But The Truth [quest] | 20.3 healing_power points (1.63 DPS) | yes | Glowing Thresher Cape (6901, +0.00 DPS, sim-verified) [dungeon]; Ceremonial Centaur Blanket (6789, -0.34 DPS) [quest]; Stargazer Cloak (9660, -0.38 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 70.5 healing_power points (5.69 DPS) | yes | Doomsayer's Robe (4746, -2.13 DPS) [quest]; Pristine Gown (253961, -3.01 DPS) [crafted]; Civinad Robes (9623, -3.68 DPS, sim-verified) [quest] |
| wrist | Earthen Silk Cuffs (254019) | Tailoring [crafted] | 17.8 healing_power points (1.44 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Silkstream Cuffs (16791, -0.53 DPS) [quest]; Nightsky Wristbands (6407, -0.58 DPS) [world_drop] |
| hands | Earthen Silk Gloves (254017) | Tailoring [crafted] | 30.5 healing_power points (2.46 DPS) | yes | Gilded Handwraps (254021, +0.00 DPS, sim-verified) [crafted]; Gloves of Old (9395, -0.55 DPS) [world_drop]; Stormcloth Gloves (10011, -0.61 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 39.3 healing_power points (3.17 DPS) | yes | Earthen Silk Belt (7061, -1.52 DPS) [crafted]; Sutarn's Ring (13105, -1.62 DPS) [world_drop]; Windchaser Cinch (14435, -2.49 DPS, sim-verified) [world_drop] |
| legs | Stoneweaver Leggings (9407) | Uldaman: Ironaya [dungeon] | 46.6 healing_power points (3.76 DPS) | yes | Pristine Leggings (253987, -0.80 DPS) [crafted]; Earthen Leggings (253999, -1.17 DPS) [crafted]; Stormcloth Pants (10010, -1.27 DPS, sim-verified) [crafted] |
| feet | Thoughtcast Boots (10578) | Razorfen Downs [dungeon] | 33.3 healing_power points (2.68 DPS) | yes | Gilded Slippers (254001, +0.00 DPS, sim-verified) [crafted]; Furen's Boots (13100, -0.11 DPS) [world_drop]; Everlast Boots (10359, -0.75 DPS) [quest] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 19.9 healing_power points (1.61 DPS) | yes | Solid Gold Band (279845, -0.09 DPS) [quest]; Welken Ring (5011, -0.33 DPS) [world_drop]; The Queen's Jewel (13094, -0.44 DPS) [world_drop] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 19.9 healing_power points (1.61 DPS) | yes | Solid Gold Band (279845, -0.09 DPS) [quest]; Welken Ring (5011, -0.33 DPS) [world_drop]; The Queen's Jewel (13094, -0.44 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (104.8 DPS) | yes | Gnarled Ash Staff (791, -0.97 DPS) [world_drop]; Staff of the Friar (3415, -1.10 DPS) [dungeon]; Death Speaker Scepter (2816, -2.95 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Captain Rackmore's Tiller (16789) | Claim Rackmore's Treasure! [quest] | 10.7 healing_power points (0.86 DPS) | yes | Jaina's Firestarter (13064, +0.00 DPS, sim-verified) [world_drop]; Wand of Mana Concentration (273043, -0.02 DPS) [world]; Earthen Rod (9381, -0.14 DPS) [dungeon] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Cloak of Blight; chest: Stormcloth Vest; wrist: Earthen Silk Cuffs; hands: Earthen Silk Gloves; waist: Gilded Cord; legs: Stoneweaver Leggings; feet: Thoughtcast Boots; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Captain Rackmore's Tiller

No-known-source sample (15 of 414, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16336 Sergeant Major's Cape

### Band 50 (troll, 025003000000000000-03505003030121431-00000000000000000)

Set DPS (verified): 172.9. Weights run: 7.9s. Verify run: 33.2s. 533 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.046, intellect=1.116 ± 0.032, spirit=2.216 ± 0.037, mp5=3.967 ± 0.037, crit=0.268 ± 0.013 per rating point (14 rating = 1%, 3.752 per %), spell_haste=-5.547 ± 0.626

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 67.9 healing_power points (6.40 DPS) | yes | Cassandra's Grace (13102, -0.37 DPS) [world_drop]; Blood Guard's Satin Cover (220899, -0.43 DPS) [vendor]; Papal Fez (9431, -0.66 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | sim-verified (172.9 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS) [dungeon]; Necklace of Calisea (1714, -0.11 DPS) [dungeon]; Choker of the High Shaman (4112, -0.11 DPS) [quest] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | 48.2 healing_power points (4.55 DPS) | yes | Blood Guard's Satin Pads (220901, +0.00 DPS) [vendor]; Inquisitor's Shawl (19507, -1.11 DPS) [dungeon]; Earthen Silk Shoulders (254033, -1.18 DPS) [crafted] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | sim-verified (172.9 DPS) | yes | Cloak of Blight (6832, +0.00 DPS) [quest]; Arcane Cloak (8286, +0.00 DPS) [world_drop]; Wingveil Cloak (10802, +0.00 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ghostweave Vest (14141, +0.00 DPS) [crafted]; Grimlok's Tribal Vestments (9415, -0.96 DPS) [dungeon]; Doomsayer's Robe (4746, -1.87 DPS) [quest] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 38.9 healing_power points (3.67 DPS) | yes | Venomshroud Armguards (14439, -1.58 DPS) [world_drop]; Earthen Silk Cuffs (254019, -1.68 DPS) [crafted]; Shizzle's Nozzle Wiper (11917, -1.78 DPS) [quest] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 57.9 healing_power points (5.46 DPS) | yes | Earthenweave Gloves (254075, -0.71 DPS) [crafted]; Earthen Silk Gloves (254017, -2.09 DPS) [crafted]; First Sergeant's Satin Gloves (220898, -2.12 DPS) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | sim-verified (172.9 DPS) | yes | Gilded Cord (254037, +0.00 DPS) [crafted]; Earthenweave Cord (254077, +0.00 DPS) [crafted]; Gilded Waistcord (254081, +0.00 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | sim-verified (172.9 DPS) | yes | Senior Designer's Pantaloons (11841, +0.00 DPS) [dungeon]; Dalewind Trousers (13008, -0.05 DPS) [world_drop]; Stoneweaver Leggings (9407, -0.99 DPS) [dungeon] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | sim-verified (172.9 DPS) | yes | Furen's Boots (13100, +0.00 DPS) [world_drop]; Vinerot Sandals (17748, +0.00 DPS) [dungeon]; Earthenweave Boots (254093, +0.00 DPS) [crafted] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 33.3 healing_power points (3.14 DPS) | yes | Electrocutioner Lagnut (9447, -0.83 DPS) [dungeon]; Runemender's Seal (273031, -0.83 DPS) [world]; Coldwater Ring (4550, -0.94 DPS) [quest] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | sim-verified (172.9 DPS) | yes | Coldwater Ring (4550, +0.00 DPS) [quest]; Electrocutioner Lagnut (9447, +0.00 DPS) [dungeon]; Runemender's Seal (273031, +0.00 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.42 DPS) [quest]; Alchemist's Stone (13503, -1.67 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Alchemist's Stone (13503, -0.86 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (172.9 DPS) | yes | Spellshifter Rod (9527, +0.00 DPS) [quest]; Spire of Hakkar (10844, +0.00 DPS) [world]; Charstone Dirk (17710, +0.00 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 13.3 healing_power points (1.26 DPS) | yes | Chillnail Splinter (10704, -0.00 DPS) [quest]; Captain Rackmore's Tiller (16789, -0.07 DPS) [quest]; Wand of Mana Concentration (273043, -0.09 DPS) [world] |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Nethergeld Shoulders; back: Imperial Red Cloak; wrist: Nethergeld Cuffs; hands: Gilded Gloves; waist: Dawnspire Cord; legs: Kilt of the Atal'ai Prophet; feet: Gilded Sandals; finger1: Eye of Adaegus; finger2: Welken Ring; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Jaina's Firestarter

No-known-source sample (15 of 533, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60 (troll, 025003031330000000-03505003030121431-00000000000000000)

Set DPS (verified): 302.7. Weights run: 6.8s. Verify run: 32.6s. 1219 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.055, intellect=0.871 ± 0.035, spirit=1.415 ± 0.033, mp5=2.375 ± 0.043, crit=0.205 ± 0.014 per rating point (14 rating = 1%, 2.873 per %), spell_haste=-25.565 ± 0.929

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (302.7 DPS) | yes | Mooncloth Circlet (14140, +0.00 DPS) [crafted]; Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 27.0 healing_power points (5.01 DPS) | yes | Lady Alizabeth's Pendant (13002, -1.07 DPS) [world_drop]; Horizon Choker (13085, -1.69 DPS) [world_drop]; Necklace of Calisea (1714, -2.04 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (302.7 DPS) | yes | Mooncloth Shoulders (14139, +0.00 DPS) [crafted]; Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp] |
| back | Wingveil Cloak (10802) | Sunken Temple: Spawn of Hakkar [dungeon] | 20.6 healing_power points (3.82 DPS) | yes | Cloak of Blight (6832, -0.40 DPS) [quest]; Indomitable Cloak (14683, -0.40 DPS) [world_drop]; Glowing Thresher Cape (6901, -0.42 DPS) [dungeon] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 95.6 healing_power points (17.78 DPS) | yes | Warlord's Satin Tunic (231632, -3.25 DPS) [vendor]; Legionnaire's Satin Tunic (227124, -4.23 DPS) [pvp]; Mooncloth Vest (14138, -4.71 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (302.7 DPS) | yes | Nethergeld Cuffs (254061, +0.00 DPS) [crafted]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted]; General's Satin Bracers (17619, -0.26 DPS) [pvp] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (302.7 DPS) | yes | Gilded Gloves (254095, +0.00 DPS) [crafted]; Blood Guard's Satin Handwraps (227126, -0.63 DPS) [pvp]; Swarmtender's Gloves (275607, -0.64 DPS) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 55.9 healing_power points (10.40 DPS) | yes | Earthenweave Cord (254077, -2.67 DPS) [crafted]; General's Satin Cinch (17621, -2.73 DPS) [pvp]; Devout Belt (16696, -2.75 DPS) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 69.6 healing_power points (12.94 DPS) | yes | Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Mooncloth Leggings (14137, -1.65 DPS) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (302.7 DPS) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; General's Satin Walkers (231630, +0.00 DPS) [vendor]; Swarmtender's Footpads (275609, -0.14 DPS) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 24.3 healing_power points (4.51 DPS) | yes | Electrocutioner Lagnut (9447, -1.00 DPS) [dungeon]; Runemender's Seal (273031, -1.00 DPS) [world]; Solid Gold Band (279845, -1.15 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 22.2 healing_power points (4.13 DPS) | yes | Electrocutioner Lagnut (9447, -0.61 DPS) [dungeon]; Runemender's Seal (273031, -0.61 DPS) [world]; Solid Gold Band (279845, -0.77 DPS) [quest] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.02 DPS) [quest]; Alchemist's Stone (13503, -1.60 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -0.19 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.72 DPS) [quest]; Alchemist's Stone (13503, -2.29 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (302.7 DPS) | yes | Spellshifter Rod (9527, +0.00 DPS) [quest]; Spire of Hakkar (10844, +0.00 DPS) [world]; Charstone Dirk (17710, +0.00 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 32.6 healing_power points (6.05 DPS) | yes | Captain Rackmore's Tiller (16789, -4.15 DPS) [quest]; Jaina's Firestarter (13064, -4.29 DPS) [world_drop]; Earthen Rod (9381, -4.38 DPS) [dungeon] |

**New at 60:** head: Devout Crown; neck: Glowing Eye of Mordresh; shoulder: Devout Mantle; back: Wingveil Cloak; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Shard of the Splithooves; ranged: Torch of Light

No-known-source sample (15 of 1219, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60, raid preset (troll, 025003031330000000-03505003030121431-00000000000000000)

Set DPS (verified): 623.4. Weights run: 4.7s. Verify run: 21.8s. 1219 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.156, intellect=0.592 ± 0.030, spirit=0.861 ± 0.030, mp5=1.616 ± 0.032, crit=0.334 ± 0.021 per rating point (14 rating = 1%, 4.681 per %), spell_haste=-3.179 ± 0.613

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (623.5 DPS) | yes | Mooncloth Circlet (14140, +0.00 DPS) [crafted]; Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 24.7 healing_power points (8.05 DPS) | yes | Lady Alizabeth's Pendant (13002, -3.85 DPS) [world_drop]; Horizon Choker (13085, -4.23 DPS) [world_drop]; Necklace of Calisea (1714, -4.74 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (623.5 DPS) | yes | Mooncloth Shoulders (14139, +0.00 DPS) [crafted]; Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp] |
| back | Wingveil Cloak (10802) | Sunken Temple: Spawn of Hakkar [dungeon] | 14.5 healing_power points (4.71 DPS) | yes | Glowing Thresher Cape (6901, -0.19 DPS) [dungeon]; Cloak of Blight (6832, -1.07 DPS) [quest]; Indomitable Cloak (14683, -1.07 DPS) [world_drop] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 86.8 healing_power points (28.25 DPS) | yes | Warlord's Satin Tunic (231632, -6.98 DPS) [vendor]; Stormcloth Vest (10020, -7.13 DPS) [crafted]; Legionnaire's Satin Tunic (227124, -8.41 DPS) [pvp] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (623.5 DPS) | yes | Nethergeld Cuffs (254061, +0.00 DPS) [crafted]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted]; General's Satin Bracers (17619, -1.46 DPS) [pvp] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (623.5 DPS) | yes | Blood Guard's Satin Handwraps (227126, +0.00 DPS) [pvp]; General's Satin Handwraps (231633, +0.00 DPS) [vendor]; Gilded Gloves (254095, +0.00 DPS) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 47.0 healing_power points (15.31 DPS) | yes | Earthenweave Cord (254077, -3.77 DPS) [crafted]; Gilded Waistcord (254081, -4.14 DPS) [crafted]; Gilded Cord (254037, -4.93 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 52.7 healing_power points (17.15 DPS) | yes | Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Earthenweave Leggings (254125, -0.62 DPS) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (623.5 DPS) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; General's Satin Walkers (231630, +0.00 DPS) [vendor]; Swarmtender's Footpads (275609, +0.00 DPS) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 15.4 healing_power points (5.01 DPS) | yes | Solid Gold Band (279845, -0.03 DPS) [quest]; Electrocutioner Lagnut (9447, -0.12 DPS) [dungeon]; Runemender's Seal (273031, -0.12 DPS) [world] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | sim-verified (623.5 DPS) | yes | Electrocutioner Lagnut (9447, +0.00 DPS) [dungeon]; Runemender's Seal (273031, +0.00 DPS) [world]; Solid Gold Band (279845, +0.00 DPS) [quest] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.35 DPS) [quest]; Alchemist's Stone (13503, -2.04 DPS) [crafted] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (623.5 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.56 DPS) [quest]; Alchemist's Stone (13503, -2.24 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (623.5 DPS) | yes | Spire of Hakkar (10844, +0.00 DPS) [world]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Charstone Dirk (17710, +0.00 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 26.5 healing_power points (8.62 DPS) | yes | Earthen Rod (9381, -5.69 DPS) [dungeon]; Captain Rackmore's Tiller (16789, -5.82 DPS) [quest]; Jaina's Firestarter (13064, -6.62 DPS) [world_drop] |

**New at 60:** head: Devout Crown; neck: Glowing Eye of Mordresh; shoulder: Devout Mantle; back: Wingveil Cloak; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Shard of the Splithooves; trinket2: Ankh of Life; ranged: Torch of Light

No-known-source sample (15 of 1219, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

