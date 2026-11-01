# Leveling BiS: Discipline

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 025004000000000000-00000000000000000-00000000000000000)

Set DPS (verified): 35.8. Weights run: 7.7s. Verify run: 26.9s. 190 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.005, intellect=0.999 ± 0.007, spirit=0.773 ± 0.006, mp5=1.703 ± 0.013, crit=0.072 ± 0.003 per rating point (14 rating = 1%, 1.015 per %), spell_haste=not significant (0.083 ± 0.029)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.35 DPS) | yes | Shadow Goggles (4373, -1.02 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.12 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 1.5 healing_power points (0.12 DPS) | yes | Tarnished Locket (279870, -0.10 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 9.0 healing_power points (0.67 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.37 DPS) [crafted]; Slime-encrusted Pads (6461, -0.56 DPS, sim-verified) [dungeon] |
| back | Sanguine Cape (14376) | World drop [world_drop] | 4.0 healing_power points (0.30 DPS) | yes | Spirit Cloak (4792, -0.07 DPS) [vendor]; Sylvan Cloak (4793, -0.07 DPS) [vendor]; Seer's Cape (6378, -0.31 DPS, sim-verified) [dungeon] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 20.3 healing_power points (1.52 DPS) | yes | Robe of the Moccasin (6465, -0.27 DPS) [dungeon]; Bloody Apron (6226, -0.55 DPS) [dungeon]; Corsair's Overshirt (5202, -0.58 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 10.1 healing_power points (0.76 DPS) | yes | Repurposed Hair Band (281256, -0.49 DPS) [quest]; Seer's Cuffs (3645, -0.57 DPS) [dungeon]; Bright Bracers (3647, -0.88 DPS, sim-verified) [world_drop] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 14.0 healing_power points (1.05 DPS) | yes | Bright Gloves (3066, -0.52 DPS) [world_drop]; Evergreen Gloves (7738, -0.71 DPS) [quest]; Magefist Gloves (12977, -0.99 DPS, sim-verified) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 15.0 healing_power points (1.12 DPS) | yes | Novice Ardent's Sash (253887, -0.43 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.53 DPS) [world_drop]; Relic Hunter Belt (11936, -0.71 DPS) [quest] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 27.1 healing_power points (2.03 DPS) | yes | Filigreed Silky Leggings (253939, -1.35 DPS) [crafted]; Filigreed Flame Leggings (253941, -1.35 DPS) [crafted]; Darkweave Breeches (12987, -1.62 DPS, sim-verified) [world_drop] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 12.0 healing_power points (0.90 DPS) | yes | Bluegill Sandals (1560, -0.52 DPS) [world]; Smoldering Boots (3076, -0.55 DPS) [world]; Kimbra Boots (6191, -0.91 DPS, sim-verified) [quest] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 6.0 healing_power points (0.45 DPS) | yes | Volcanic Rock Ring (12053, -0.22 DPS) [world_drop]; Furlfeather Ring (286745, -0.30 DPS) [world]; Minor Channeling Ring (1449, -0.30 DPS) [quest] |
| finger2 | Band of Purification (12996) | World drop [world_drop] | 4.6 healing_power points (0.35 DPS) | yes | Volcanic Rock Ring (12053, -0.12 DPS, sim-verified) [world_drop]; Furlfeather Ring (286745, -0.20 DPS) [world]; Minor Channeling Ring (1449, -0.20 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Westfall (2042) | The Defias Brotherhood [quest] | sim-verified (35.8 DPS) | yes | Staff of the Blessed Seer (2271, -0.14 DPS) [dungeon]; Lesser Staff of the Spire (1300, -0.27 DPS) [world]; Twisted Chanter's Staff (890, -4.37 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Flaring Baton (5326) (or Moonstone Wand (15204)) | The Escape [quest] | 2.0 healing_power points (0.15 DPS) | yes | Moonstone Wand (15204, +0.00 DPS) [quest]; Dwarven Flamestick (5241, -0.03 DPS) [quest]; Sable Wand (7607, -0.03 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Band of Purification; main_hand: Staff of Westfall; ranged: Flaring Baton

No-known-source sample (15 of 190, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes; 16607 Acolyte's Sacrificial Robes; 18851 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (human, 025005031320000000-00000000000000000-00000000000000000)

Set DPS (verified): 59.6. Weights run: 8.0s. Verify run: 25.3s. 321 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.008, intellect=1.080 ± 0.010, spirit=1.156 ± 0.009, mp5=2.537 ± 0.009, crit=0.112 ± 0.005 per rating point (14 rating = 1%, 1.567 per %), spell_haste=not significant (-0.022 ± 0.060)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 39.9 healing_power points (3.46 DPS) | yes | Filigreed Pristine Circlet (253975, -1.55 DPS) [crafted]; Nightsky Cowl (4039, -1.83 DPS) [world_drop]; Pristine Circlet (253949, -1.90 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 8.9 healing_power points (0.77 DPS) | yes | Crystal Starfire Medallion (5003, -0.10 DPS) [world_drop]; Pendant of Myzrael (4614, -0.17 DPS) [dungeon]; Glowing Green Talisman (5002, -0.17 DPS) [world_drop] |
| shoulder | Mantle of Honor (3560) | Bride of the Embalmer [quest] | 15.7 healing_power points (1.36 DPS) | yes | Nightsky Mantle (4718, -0.11 DPS) [world_drop]; Faerie Mantle (5820, -0.19 DPS) [quest]; Death Speaker Mantle (6685, -0.33 DPS) [dungeon] |
| back | Prelacy Cape (7004) | Researching the Corruption [quest] | sim-verified (59.6 DPS) | yes | Glowing Thresher Cape (6901, +0.00 DPS) [dungeon]; Repairman's Cape (9605, +0.00 DPS) [quest]; Sergeant Major's Cape (16315, -0.13 DPS) [pvp] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (59.6 DPS) | yes | Pristine Gown (253961, +0.00 DPS) [crafted]; Robes of Arugal (6324, -0.60 DPS) [dungeon]; Beguiler Robes (7728, -0.84 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 13.0 healing_power points (1.13 DPS) | yes | Nightsky Wristbands (6407, -0.27 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.38 DPS) [world_drop]; Spidertank Oilrag (9448, -0.47 DPS) [dungeon] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 19.3 healing_power points (1.67 DPS) | yes | Truefaith Gloves (7049, -0.09 DPS) [crafted]; Hotshot Pilot's Gloves (9491, -0.42 DPS) [dungeon]; Pristine Gloves (253913, -0.44 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 15.3 healing_power points (1.33 DPS) | yes | Novice Ardent's Sash (253887, -0.23 DPS) [crafted]; Resilient Cord (14406, -0.37 DPS) [world_drop]; Dreamer's Belt (4829, -0.47 DPS) [vendor] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 35.3 healing_power points (3.06 DPS) | yes | Filigreed Pristine Leggings (253937, -0.54 DPS) [crafted]; Earthen Leggings (253999, -0.66 DPS) [crafted]; Darkweave Breeches (12987, -1.80 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 32.2 healing_power points (2.79 DPS) | yes | Acidic Walkers (9454, -1.64 DPS) [dungeon]; Pristine Boots (253889, -1.73 DPS) [crafted]; Frothing Slippers (254003, -1.73 DPS) [crafted] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 17.1 healing_power points (1.48 DPS) | yes | The Queen's Jewel (13094, -0.49 DPS) [world_drop]; Ring of Calm (6790, -0.78 DPS) [quest]; Defias Mage Ring (1077, -0.81 DPS) [world] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 17.1 healing_power points (1.48 DPS) | yes | The Queen's Jewel (13094, -0.49 DPS) [world_drop]; Ring of Calm (6790, -0.78 DPS) [quest]; Defias Mage Ring (1077, -0.81 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-verified (59.6 DPS) | yes | Gnarled Ash Staff (791, +0.00 DPS) [world_drop]; Death Speaker Scepter (2816, +0.00 DPS) [dungeon]; Staff of the Friar (3415, +0.00 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 9.4 healing_power points (0.81 DPS) | yes | Consecrated Wand (5244, -0.43 DPS) [quest]; Summoner's Wand (5245, -0.44 DPS) [world]; Lesser Mystic Wand (11289, -0.44 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Mantle of Honor; back: Prelacy Cape; chest: Death Speaker Robes; hands: Gloves of Old; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Spine of the Devourer; ranged: Wand of Mana Concentration

No-known-source sample (15 of 321, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 40 (human, 025005031330011520-00000000000000000-00000000000000000)

Set DPS (verified): 131.0. Weights run: 9.7s. Verify run: 29.0s. 422 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.007, intellect=1.124 ± 0.021, spirit=1.841 ± 0.020, mp5=3.394 ± 0.029, crit=0.265 ± 0.009 per rating point (14 rating = 1%, 3.710 per %), spell_haste=-4.206 ± 0.355

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 57.7 healing_power points (5.38 DPS) | yes | Holy Shroud (2721, -1.27 DPS) [world_drop]; Whitemane's Chapeau (7720, -1.37 DPS, sim-verified) [dungeon]; Earthen Silk Hood (254015, -1.61 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 29.0 healing_power points (2.71 DPS) | yes | Amberglow Talisman (10824, -0.99 DPS) [quest]; Explorers' League Commendation (7746, -1.16 DPS) [quest]; Necklace of Calisea (1714, -2.03 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 36.6 healing_power points (3.42 DPS) | yes | Earthen Silk Shoulders (254033, -1.18 DPS, sim-verified) [crafted]; Crimson Silk Shoulders (7059, -1.37 DPS) [crafted]; Mistscape Mantle (4734, -1.40 DPS) [dungeon] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 21.7 healing_power points (2.03 DPS) | yes | Repairman's Cape (9605, -0.29 DPS) [quest]; Energy Cloak (9397, -0.31 DPS) [dungeon]; Ceremonial Centaur Blanket (6789, -0.54 DPS, sim-verified) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 72.7 healing_power points (6.79 DPS) | yes | Civinad Robes (9623, -2.30 DPS) [quest]; Dreamweave Vest (10021, -3.44 DPS) [crafted]; Doomsayer's Robe (4746, -4.37 DPS, sim-verified) [quest] |
| wrist | Earthen Silk Cuffs (254019) | Tailoring [crafted] | 19.2 healing_power points (1.79 DPS) | yes | Mindthrust Bracers (1974, -0.32 DPS) [dungeon]; Silkstream Cuffs (16791, -0.58 DPS) [quest]; Nightsky Wristbands (6407, -0.65 DPS) [world_drop] |
| hands | Earthen Silk Gloves (254017) | Tailoring [crafted] | 32.7 healing_power points (3.05 DPS) | yes | Gloves of Old (9395, -0.47 DPS) [world_drop]; Restorer's Fine Gloves (270059, -0.65 DPS) [quest]; Gilded Handwraps (254021, -2.02 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 42.0 healing_power points (3.92 DPS) | yes | Earthen Silk Belt (7061, -1.80 DPS) [crafted]; Sutarn's Ring (13105, -1.84 DPS) [world_drop]; Windchaser Cinch (14435, -2.76 DPS, sim-verified) [world_drop] |
| legs | Stoneweaver Leggings (9407) | Uldaman: Ironaya [dungeon] | 51.1 healing_power points (4.77 DPS) | yes | Pristine Leggings (253987, -1.12 DPS) [crafted]; Earthen Leggings (253999, -1.48 DPS) [crafted]; Stormcloth Pants (10010, -1.97 DPS, sim-verified) [crafted] |
| feet | Furen's Boots (13100) | World drop [world_drop] | 37.3 healing_power points (3.48 DPS) | yes | Gilded Slippers (254001, -0.20 DPS) [crafted]; Everlast Boots (10359, -1.00 DPS) [quest]; Thoughtcast Boots (10578, -1.79 DPS, sim-verified) [dungeon] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 21.9 healing_power points (2.04 DPS) | yes | Welken Ring (5011, -0.32 DPS) [world_drop]; The Queen's Jewel (13094, -0.46 DPS) [world_drop]; Blush Ember Ring (13093, -0.67 DPS) [world_drop] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 21.9 healing_power points (2.04 DPS) | yes | Welken Ring (5011, -0.39 DPS, sim-verified) [world_drop]; The Queen's Jewel (13094, -0.46 DPS) [world_drop]; Blush Ember Ring (13093, -0.67 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (131.0 DPS) | yes | Gnarled Ash Staff (791, -1.24 DPS) [world_drop]; Staff of the Friar (3415, -1.44 DPS) [dungeon]; Death Speaker Scepter (2816, -5.19 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 12.3 healing_power points (1.14 DPS) | yes | Wand of Mana Concentration (273043, -0.09 DPS) [world]; Captain Rackmore's Tiller (16789, -0.14 DPS, sim-verified) [quest]; Goblin Igniter (5253, -0.21 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Glowing Thresher Cape; chest: Stormcloth Vest; wrist: Earthen Silk Cuffs; hands: Earthen Silk Gloves; waist: Gilded Cord; legs: Stoneweaver Leggings; feet: Furen's Boots; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Jaina's Firestarter

No-known-source sample (15 of 422, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 50 (human, 025005031330011530-03501000000000000-00000000000000000)

Set DPS (verified): 196.5. Weights run: 9.6s. Verify run: 38.8s. 541 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.033, intellect=1.609 ± 0.019, spirit=1.396 ± 0.021, mp5=3.771 ± 0.028, crit=0.553 ± 0.016 per rating point (14 rating = 1%, 7.738 per %), spell_haste=-16.540 ± 0.565

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | sim-verified (196.5 DPS) | yes | Papal Fez (9431, +0.00 DPS) [dungeon]; Knight-Lieutenant's Satin Cover (220896, +0.00 DPS) [vendor]; Cassandra's Grace (13102, -0.39 DPS) [world_drop] |
| neck | Horizon Choker (13085) | World drop [world_drop] | sim-verified (196.5 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS) [dungeon]; Necklace of Calisea (1714, -0.56 DPS) [dungeon]; Mindburst Medallion (11196, -0.91 DPS) [quest] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | 47.2 healing_power points (3.75 DPS) | yes | Knight-Lieutenant's Satin Pads (220894, +0.00 DPS) [vendor]; Inquisitor's Shawl (19507, -0.34 DPS) [dungeon]; Hibernal Mantle (8111, -1.36 DPS) [world_drop] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 23.3 healing_power points (1.85 DPS) | yes | Wingveil Cloak (10802, -0.23 DPS) [dungeon]; Highborne Cloak (14450, -0.26 DPS) [world_drop]; Chainlink Towel (9648, -0.36 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 69.2 healing_power points (5.48 DPS) | yes | Knight's Satin Armor (220892, -0.76 DPS) [vendor]; Ghostweave Vest (14141, -0.88 DPS) [crafted]; Grimlok's Tribal Vestments (9415, -1.00 DPS) [dungeon] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 38.2 healing_power points (3.03 DPS) | yes | Shizzle's Nozzle Wiper (11917, -1.17 DPS) [quest]; Mindthrust Bracers (1974, -1.50 DPS) [dungeon]; Forgotten Wraps (9433, -1.50 DPS) [world_drop] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 56.3 healing_power points (4.46 DPS) | yes | Earthenweave Gloves (254075, -1.18 DPS) [crafted]; Sergeant Major's Satin Gloves (220897, -1.49 DPS) [vendor]; Gilded Handwraps (254021, -1.82 DPS) [crafted] |
| waist | Gilded Waistcord (254081) | Tailoring [crafted] | 43.5 healing_power points (3.45 DPS) | yes | Gilded Cord (254037, -0.02 DPS) [crafted]; Dawnspire Cord (12466, -0.07 DPS) [dungeon]; Earthenweave Cord (254077, -0.17 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 61.4 healing_power points (4.87 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -0.10 DPS) [dungeon]; Knight's Satin Leggings (220893, -0.14 DPS) [vendor]; Stoneweaver Leggings (9407, -0.93 DPS) [dungeon] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 45.5 healing_power points (3.61 DPS) | yes | Earthenweave Boots (254093, -0.33 DPS) [crafted]; Vinerot Sandals (17748, -0.40 DPS) [dungeon]; Sergeant Major's Satin Boots (220895, -0.49 DPS) [vendor] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 26.4 healing_power points (2.09 DPS) | yes | Electrocutioner Lagnut (9447, -0.61 DPS) [dungeon]; Runemender's Seal (273031, -0.61 DPS) [world]; Choking Band (11868, -0.65 DPS) [quest] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | sim-verified (196.5 DPS) | yes | Electrocutioner Lagnut (9447, +0.00 DPS) [dungeon]; Choking Band (11868, +0.00 DPS) [quest]; Runemender's Seal (273031, +0.00 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -0.44 DPS) [quest]; Alchemist's Stone (13503, -0.89 DPS) [crafted] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -0.22 DPS) [quest]; Alchemist's Stone (13503, -0.66 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (196.5 DPS) | yes | Spellshifter Rod (9527, +0.00 DPS) [quest]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Charstone Dirk (17710, +0.00 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cairnstone Sliver (9654) | The Morrow Stone [quest] | 15.0 healing_power points (1.19 DPS) | yes | Wand of Mana Concentration (273043, -0.08 DPS) [world]; Jaina's Firestarter (13064, -0.09 DPS) [world_drop]; Goblin Igniter (5253, -0.35 DPS) [quest] |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Nethergeld Shoulders; back: Imperial Red Cloak; wrist: Nethergeld Cuffs; hands: Gilded Gloves; waist: Gilded Waistcord; legs: Senior Designer's Pantaloons; feet: Gilded Sandals; finger1: Eye of Adaegus; finger2: Welken Ring; trinket2: Evonice's Landin' Pilla; main_hand: Glowing Brightwood Staff; ranged: Cairnstone Sliver

No-known-source sample (15 of 541, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60 (human, 025005031330011530-03505003030000000-00000000000000000)

Set DPS (verified): 282.5. Weights run: 9.2s. Verify run: 71.1s. 1226 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.041, intellect=1.062 ± 0.039, spirit=3.412 ± 0.055, mp5=4.095 ± 0.064, crit=0.343 ± 0.017 per rating point (14 rating = 1%, 4.798 per %), spell_haste=-13.003 ± 0.795

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Circlet (14140, +0.00 DPS) [crafted]; Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor] |
| neck | Lady Alizabeth's Pendant (13002) | World drop [world_drop] | 51.2 healing_power points (7.58 DPS) | yes | Amberglow Talisman (10824, -2.53 DPS) [quest]; Necklace of Calisea (1714, -2.94 DPS) [dungeon]; Explorers' League Commendation (7746, -3.03 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Shoulders (14139, +0.00 DPS) [crafted]; Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Earthenweave Mantle (254119, +0.00 DPS) [crafted] |
| back | Ethereal Mist Cape (11873) | Ancient Spirit [quest] | sim-verified (282.5 DPS) | yes | Indomitable Cloak (14683, +0.00 DPS) [world_drop]; Wingveil Cloak (10802, -0.24 DPS) [dungeon]; Arcane Cloak (8286, -0.47 DPS) [world_drop] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 127.6 healing_power points (18.88 DPS) | yes | Mooncloth Vest (14138, -2.30 DPS) [crafted]; Knight-Captain's Satin Tunic (227122, -3.54 DPS) [pvp]; Field Marshal's Satin Tunic (231624, -3.87 DPS) [vendor] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Satin Bracers (17606, +0.00 DPS) [pvp]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -1.35 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swarmtender's Gloves (275607, +0.00 DPS) [crafted]; Gilded Gloves (254095, -2.47 DPS) [crafted]; Earthenweave Gloves (254075, -2.97 DPS) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (282.5 DPS) | yes | Marshal's Satin Sash (17609, +0.00 DPS) [pvp]; Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Earthenweave Cord (254077, +0.00 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 118.4 healing_power points (17.52 DPS) | yes | Knight-Captain's Satin Legguards (227125, -2.18 DPS) [pvp]; Mooncloth Leggings (14137, -2.33 DPS) [crafted]; Dalewind Trousers (13008, -3.48 DPS) [world_drop] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | sim-verified (282.5 DPS) | yes | Devout Sandals (16691, -0.53 DPS) [dungeon]; Elunarian Boots (14458, -2.79 DPS) [world_drop]; Earthenweave Boots (254093, -2.91 DPS) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 48.2 healing_power points (7.13 DPS) | yes | Choking Band (11868, -0.56 DPS) [quest]; Seedtime Hoop (9655, -2.08 DPS) [quest]; Coldwater Ring (4550, -2.11 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 47.3 healing_power points (7.00 DPS) | yes | Choking Band (11868, -0.44 DPS) [quest]; Seedtime Hoop (9655, -1.95 DPS) [quest]; Coldwater Ring (4550, -1.99 DPS) [quest] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (282.5 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (282.5 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (282.5 DPS) | yes | Spire of Hakkar (10844, +0.00 DPS) [world]; Staff of Hale Magefire (13000, +0.00 DPS) [world_drop]; Charstone Dirk (17710, +0.00 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 54.5 healing_power points (8.07 DPS) | yes | Cairnstone Sliver (9654, -4.76 DPS) [quest]; Jaina's Firestarter (13064, -5.61 DPS) [world_drop]; Captain Rackmore's Tiller (16789, -5.67 DPS) [quest] |

**New at 60:** head: Devout Crown; neck: Lady Alizabeth's Pendant; shoulder: Devout Mantle; back: Ethereal Mist Cape; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Shard of the Splithooves; trinket2: Uther's Strength; ranged: Torch of Light

No-known-source sample (15 of 1226, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60, raid preset (human, 025005031330011530-03505003030000000-00000000000000000)

Set DPS (verified): 538.3. Weights run: 6.1s. Verify run: 26.7s. 1226 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.117, intellect=0.856 ± 0.039, spirit=1.111 ± 0.039, mp5=2.155 ± 0.042, crit=0.573 ± 0.024 per rating point (14 rating = 1%, 8.026 per %), spell_haste=-6.873 ± 0.846

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (538.3 DPS) | yes | Mooncloth Circlet (14140, +0.00 DPS) [crafted]; Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 26.9 healing_power points (6.49 DPS) | yes | Lady Alizabeth's Pendant (13002, -2.46 DPS) [world_drop]; Horizon Choker (13085, -2.52 DPS) [world_drop]; Necklace of Calisea (1714, -3.16 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (538.3 DPS) | yes | Mooncloth Shoulders (14139, +0.00 DPS) [crafted]; Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp] |
| back | Sergeant Major's Cape (16337) | PvP rank 9 · Sergeant Major · Alliance [pvp] | sim-verified (538.3 DPS) | yes | Glowing Thresher Cape (6901, +0.00 DPS) [dungeon]; Wingveil Cloak (10802, +0.00 DPS) [dungeon]; Ethereal Mist Cape (11873, +0.00 DPS) [quest] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 90.8 healing_power points (21.95 DPS) | yes | Field Marshal's Satin Tunic (231624, -3.82 DPS) [vendor]; Knight-Captain's Satin Tunic (227122, -5.41 DPS) [pvp]; Stormcloth Vest (10020, -5.78 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (538.3 DPS) | yes | Nethergeld Cuffs (254061, +0.00 DPS) [crafted]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted]; Marshal's Satin Bracers (17606, -0.38 DPS) [pvp] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (538.3 DPS) | yes | Knight-Lieutenant's Satin Handwraps (227128, +0.00 DPS) [pvp]; Gilded Gloves (254095, +0.00 DPS) [crafted]; Marshal's Satin Handwraps (231623, -0.56 DPS) [vendor] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 53.0 healing_power points (12.82 DPS) | yes | Earthenweave Cord (254077, -3.57 DPS) [crafted]; Devout Belt (16696, -3.59 DPS) [dungeon]; Gilded Waistcord (254081, -3.94 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 62.4 healing_power points (15.09 DPS) | yes | Knight-Captain's Satin Legguards (227125, +0.00 DPS) [pvp]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Earthenweave Leggings (254125, -1.84 DPS) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (538.3 DPS) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; Marshal's Satin Walkers (231627, +0.00 DPS) [vendor]; Swarmtender's Footpads (275609, +0.00 DPS) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 20.8 healing_power points (5.03 DPS) | yes | Electrocutioner Lagnut (9447, -0.97 DPS) [dungeon]; Runemender's Seal (273031, -0.97 DPS) [world]; Choking Band (11868, -1.53 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 18.5 healing_power points (4.47 DPS) | yes | Electrocutioner Lagnut (9447, -0.41 DPS) [dungeon]; Runemender's Seal (273031, -0.41 DPS) [world]; Choking Band (11868, -0.97 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -1.23 DPS) [quest]; Evonice's Landin' Pilla (18951, -1.67 DPS) [quest]; Thunderbrew's Boot Flask (744, -2.20 DPS) [quest] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (538.3 DPS) | yes | Shard of the Splithooves (10659, -0.10 DPS) [quest]; Evonice's Landin' Pilla (18951, -0.54 DPS) [quest]; Thunderbrew's Boot Flask (744, -1.07 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (538.3 DPS) | yes | Spire of Hakkar (10844, +0.00 DPS) [world]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Charstone Dirk (17710, +0.00 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 29.2 healing_power points (7.07 DPS) | yes | Cairnstone Sliver (9654, -4.69 DPS) [quest]; Captain Rackmore's Tiller (16789, -4.81 DPS) [quest]; Earthen Rod (9381, -4.89 DPS) [dungeon] |

**New at 60:** head: Devout Crown; neck: Glowing Eye of Mordresh; shoulder: Devout Mantle; back: Sergeant Major's Cape; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Uther's Strength; trinket2: Ankh of Life; ranged: Torch of Light

No-known-source sample (15 of 1226, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

## Horde

### Band 20 (undead, 025004000000000000-00000000000000000-00000000000000000)

Set DPS (verified): 35.2. Weights run: 7.7s. Verify run: 26.3s. 183 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.005, intellect=0.999 ± 0.007, spirit=0.773 ± 0.006, mp5=1.703 ± 0.013, crit=0.072 ± 0.003 per rating point (14 rating = 1%, 1.015 per %), spell_haste=not significant (0.083 ± 0.029)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.35 DPS) | yes | Shadow Goggles (4373, -1.04 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.12 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 2.3 healing_power points (0.17 DPS) | yes | Scholarly Pendant (277203, -0.25 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 9.0 healing_power points (0.67 DPS) | yes | Slime-encrusted Pads (6461, -0.30 DPS, sim-verified) [dungeon]; Reinforced Woolen Shoulders (4315, -0.37 DPS) [crafted] |
| back | Sanguine Cape (14376) | World drop [world_drop] | 4.0 healing_power points (0.30 DPS) | yes | Spirit Cloak (4792, -0.07 DPS) [vendor]; Antiquated Cloak (4799, -0.07 DPS) [vendor]; Seer's Cape (6378, -0.10 DPS, sim-verified) [dungeon] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 20.3 healing_power points (1.52 DPS) | yes | Robe of the Moccasin (6465, -0.27 DPS) [dungeon]; Corsair's Overshirt (5202, -0.37 DPS, sim-verified) [dungeon]; Bloody Apron (6226, -0.55 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 10.1 healing_power points (0.76 DPS) | yes | Featherbead Bracers (15452, -0.53 DPS) [quest]; Seer's Cuffs (3645, -0.57 DPS) [dungeon]; Bright Bracers (3647, -0.96 DPS, sim-verified) [world_drop] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 14.0 healing_power points (1.05 DPS) | yes | Bright Gloves (3066, -0.52 DPS) [world_drop]; Smoldering Gloves (3074, -0.76 DPS) [world]; Magefist Gloves (12977, -0.88 DPS, sim-verified) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 15.0 healing_power points (1.12 DPS) | yes | Novice Ardent's Sash (253887, -0.35 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.53 DPS) [world_drop]; Wise Man's Belt (4786, -0.75 DPS) [vendor] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 27.1 healing_power points (2.03 DPS) | yes | Filigreed Silky Leggings (253939, -1.35 DPS) [crafted]; Filigreed Flame Leggings (253941, -1.35 DPS) [crafted]; Darkweave Breeches (12987, -1.56 DPS, sim-verified) [world_drop] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 12.0 healing_power points (0.90 DPS) | yes | Smoldering Boots (3076, -0.55 DPS) [world]; Sanguine Sandals (14374, -0.60 DPS) [world_drop]; Bluegill Sandals (1560, -0.96 DPS, sim-verified) [world] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 6.0 healing_power points (0.45 DPS) | yes | Band of Purification (12996, -0.10 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.22 DPS) [world_drop]; Furlfeather Ring (286745, -0.30 DPS) [world] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 5.0 healing_power points (0.37 DPS) | yes | Volcanic Rock Ring (12053, -0.15 DPS) [world_drop]; Furlfeather Ring (286745, -0.22 DPS) [world]; Band of Purification (12996, -0.24 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) | The Wrath of Rath'mael [quest] | sim-verified (35.2 DPS) | yes | Staff of the Blessed Seer (2271, +0.00 DPS) [dungeon]; Crescent Staff (6505, -0.06 DPS) [quest]; Twisted Chanter's Staff (890, -1.63 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Decay (5252) (or Flaring Baton (5326)) | Beren's Peril [quest] | 2.0 healing_power points (0.15 DPS) | yes | Flaring Baton (5326, +0.00 DPS) [quest]; Wisesight Wand (286750, -0.07 DPS) [world] |

**New at 20:** head: Pristine Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Gnarled Necromancer's Staff; ranged: Wand of Decay

No-known-source sample (15 of 183, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (undead, 025005031320000000-00000000000000000-00000000000000000)

Set DPS (verified): 63.7. Weights run: 8.0s. Verify run: 26.1s. 313 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.008, intellect=1.080 ± 0.010, spirit=1.156 ± 0.009, mp5=2.537 ± 0.009, crit=0.112 ± 0.005 per rating point (14 rating = 1%, 1.567 per %), spell_haste=not significant (-0.022 ± 0.060)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 39.9 healing_power points (3.46 DPS) | yes | Filigreed Pristine Circlet (253975, -1.55 DPS) [crafted]; Nightsky Cowl (4039, -1.83 DPS) [world_drop]; Pristine Circlet (253949, -1.90 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 8.9 healing_power points (0.77 DPS) | yes | Crystal Starfire Medallion (5003, -0.10 DPS) [world_drop]; Pendant of Myzrael (4614, -0.17 DPS) [dungeon]; Glowing Green Talisman (5002, -0.17 DPS) [world_drop] |
| shoulder | Ghostly Mantle (3324) | Deathstalkers in Shadowfang [quest] | sim-verified (63.7 DPS) | yes | Nightsky Mantle (4718, +0.00 DPS) [world_drop]; Death Speaker Mantle (6685, -0.09 DPS) [dungeon]; Desert Shoulders (15457, -0.15 DPS) [quest] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 16.2 healing_power points (1.41 DPS) | yes | Amy's Blanket (13005, -0.71 DPS) [world_drop]; Nightsky Cloak (4719, -0.81 DPS) [world_drop]; Insignia Cloak (4722, -0.81 DPS) [world_drop] |
| chest | Pristine Gown (253961) | Tailoring [crafted] | 32.2 healing_power points (2.79 DPS) | yes | Death Speaker Robes (6682, -0.03 DPS) [dungeon]; Robes of Arugal (6324, -0.63 DPS) [dungeon]; Beguiler Robes (7728, -0.86 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 13.0 healing_power points (1.13 DPS) | yes | Nightsky Wristbands (6407, -0.27 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.38 DPS) [world_drop]; Spidertank Oilrag (9448, -0.47 DPS) [dungeon] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 19.3 healing_power points (1.67 DPS) | yes | Truefaith Gloves (7049, -0.09 DPS) [crafted]; Hotshot Pilot's Gloves (9491, -0.42 DPS) [dungeon]; Pristine Gloves (253913, -0.44 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 15.3 healing_power points (1.33 DPS) | yes | Lilac Sash (6780, -0.18 DPS) [quest]; Novice Ardent's Sash (253887, -0.23 DPS) [crafted]; Resilient Cord (14406, -0.37 DPS) [world_drop] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 35.3 healing_power points (3.06 DPS) | yes | Filigreed Pristine Leggings (253937, -0.54 DPS) [crafted]; Earthen Leggings (253999, -0.66 DPS) [crafted]; Sacred Burial Trousers (6282, -1.69 DPS) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 32.2 healing_power points (2.79 DPS) | yes | Acidic Walkers (9454, -1.64 DPS) [dungeon]; Pristine Boots (253889, -1.73 DPS) [crafted]; Frothing Slippers (254003, -1.73 DPS) [crafted] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 17.1 healing_power points (1.48 DPS) | yes | Solid Gold Band (279845, -0.03 DPS) [quest]; The Queen's Jewel (13094, -0.49 DPS) [world_drop]; Defias Mage Ring (1077, -0.81 DPS) [world] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 17.1 healing_power points (1.48 DPS) | yes | Solid Gold Band (279845, -0.03 DPS) [quest]; The Queen's Jewel (13094, -0.49 DPS) [world_drop]; Defias Mage Ring (1077, -0.81 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-verified (63.7 DPS) | yes | Gnarled Ash Staff (791, +0.00 DPS) [world_drop]; Death Speaker Scepter (2816, +0.00 DPS) [dungeon]; Staff of the Friar (3415, +0.00 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 9.4 healing_power points (0.81 DPS) | yes | Summoner's Wand (5245, -0.44 DPS) [world]; Lesser Mystic Wand (11289, -0.44 DPS) [crafted]; Starfaller (13063, -0.44 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Ghostly Mantle; back: Glowing Thresher Cape; chest: Pristine Gown; hands: Gloves of Old; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Spine of the Devourer; ranged: Wand of Mana Concentration

No-known-source sample (15 of 313, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (undead, 025005031330011520-00000000000000000-00000000000000000)

Set DPS (verified): 134.8. Weights run: 9.7s. Verify run: 29.4s. 414 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.007, intellect=1.124 ± 0.021, spirit=1.841 ± 0.020, mp5=3.394 ± 0.029, crit=0.265 ± 0.009 per rating point (14 rating = 1%, 3.710 per %), spell_haste=-4.206 ± 0.355

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 57.7 healing_power points (5.38 DPS) | yes | Holy Shroud (2721, -1.27 DPS) [world_drop]; Whitemane's Chapeau (7720, -1.36 DPS, sim-verified) [dungeon]; Earthen Silk Hood (254015, -1.61 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 29.0 healing_power points (2.71 DPS) | yes | Amberglow Talisman (10824, -0.99 DPS) [quest]; Triune Amulet (7722, -1.50 DPS) [dungeon]; Necklace of Calisea (1714, -2.30 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 36.6 healing_power points (3.42 DPS) | yes | Crimson Silk Shoulders (7059, -1.37 DPS) [crafted]; Mistscape Mantle (4734, -1.40 DPS) [dungeon]; Earthen Silk Shoulders (254033, -1.53 DPS, sim-verified) [crafted] |
| back | Cloak of Blight (6832) | Nothing But The Truth [quest] | 23.9 healing_power points (2.23 DPS) | yes | Glowing Thresher Cape (6901, +0.00 DPS, sim-verified) [dungeon]; Ceremonial Centaur Blanket (6789, -0.48 DPS) [quest]; Stargazer Cloak (9660, -0.52 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 72.7 healing_power points (6.79 DPS) | yes | Civinad Robes (9623, -2.30 DPS) [quest]; Doomsayer's Robe (4746, -3.00 DPS, sim-verified) [quest]; Dreamweave Vest (10021, -3.44 DPS) [crafted] |
| wrist | Earthen Silk Cuffs (254019) | Tailoring [crafted] | 19.2 healing_power points (1.79 DPS) | yes | Mindthrust Bracers (1974, -0.46 DPS, sim-verified) [dungeon]; Silkstream Cuffs (16791, -0.58 DPS) [quest]; Nightsky Wristbands (6407, -0.65 DPS) [world_drop] |
| hands | Earthen Silk Gloves (254017) | Tailoring [crafted] | 32.7 healing_power points (3.05 DPS) | yes | Gilded Handwraps (254021, +0.00 DPS, sim-verified) [crafted]; Gloves of Old (9395, -0.47 DPS) [world_drop]; Stormcloth Gloves (10011, -0.77 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 42.0 healing_power points (3.92 DPS) | yes | Earthen Silk Belt (7061, -1.80 DPS) [crafted]; Sutarn's Ring (13105, -1.84 DPS) [world_drop]; Windchaser Cinch (14435, -3.11 DPS, sim-verified) [world_drop] |
| legs | Stoneweaver Leggings (9407) | Uldaman: Ironaya [dungeon] | 51.1 healing_power points (4.77 DPS) | yes | Pristine Leggings (253987, -1.12 DPS) [crafted]; Earthen Leggings (253999, -1.48 DPS) [crafted]; Stormcloth Pants (10010, -1.73 DPS, sim-verified) [crafted] |
| feet | Furen's Boots (13100) | World drop [world_drop] | 37.3 healing_power points (3.48 DPS) | yes | Gilded Slippers (254001, -0.20 DPS) [crafted]; Thoughtcast Boots (10578, -0.53 DPS, sim-verified) [dungeon]; Everlast Boots (10359, -1.00 DPS) [quest] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 21.9 healing_power points (2.04 DPS) | yes | Solid Gold Band (279845, -0.16 DPS) [quest]; Welken Ring (5011, -0.32 DPS) [world_drop]; The Queen's Jewel (13094, -0.46 DPS) [world_drop] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 21.9 healing_power points (2.04 DPS) | yes | Welken Ring (5011, -0.32 DPS) [world_drop]; The Queen's Jewel (13094, -0.46 DPS) [world_drop]; Solid Gold Band (279845, -0.53 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (134.7 DPS) | yes | Gnarled Ash Staff (791, -1.24 DPS) [world_drop]; Staff of the Friar (3415, -1.44 DPS) [dungeon]; Death Speaker Scepter (2816, -1.92 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 12.3 healing_power points (1.14 DPS) | yes | Wand of Mana Concentration (273043, -0.09 DPS) [world]; Goblin Igniter (5253, -0.21 DPS) [quest]; Captain Rackmore's Tiller (16789, -0.52 DPS, sim-verified) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Cloak of Blight; chest: Stormcloth Vest; wrist: Earthen Silk Cuffs; hands: Earthen Silk Gloves; waist: Gilded Cord; legs: Stoneweaver Leggings; feet: Furen's Boots; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Jaina's Firestarter

No-known-source sample (15 of 414, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16336 Sergeant Major's Cape

### Band 50 (undead, 025005031330011530-03501000000000000-00000000000000000)

Set DPS (verified): 201.2. Weights run: 9.6s. Verify run: 39.8s. 533 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.033, intellect=1.609 ± 0.019, spirit=1.396 ± 0.021, mp5=3.771 ± 0.028, crit=0.553 ± 0.016 per rating point (14 rating = 1%, 7.738 per %), spell_haste=-16.540 ± 0.565

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | sim-verified (201.2 DPS) | yes | Papal Fez (9431, +0.00 DPS) [dungeon]; Blood Guard's Satin Cover (220899, +0.00 DPS) [vendor]; Cassandra's Grace (13102, -0.39 DPS) [world_drop] |
| neck | Horizon Choker (13085) | World drop [world_drop] | sim-verified (201.2 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS) [dungeon]; Necklace of Calisea (1714, -0.56 DPS) [dungeon]; Choker of the High Shaman (4112, -0.71 DPS) [quest] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | 47.2 healing_power points (3.75 DPS) | yes | Blood Guard's Satin Pads (220901, +0.00 DPS) [vendor]; Inquisitor's Shawl (19507, -0.34 DPS) [dungeon]; Hibernal Mantle (8111, -1.36 DPS) [world_drop] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 23.3 healing_power points (1.85 DPS) | yes | Wingveil Cloak (10802, -0.23 DPS) [dungeon]; Highborne Cloak (14450, -0.26 DPS) [world_drop]; Chainlink Towel (9648, -0.36 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 69.2 healing_power points (5.48 DPS) | yes | Stone Guard's Satin Armor (220903, -0.76 DPS) [vendor]; Ghostweave Vest (14141, -0.88 DPS) [crafted]; Grimlok's Tribal Vestments (9415, -1.00 DPS) [dungeon] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 38.2 healing_power points (3.03 DPS) | yes | Shizzle's Nozzle Wiper (11917, -1.17 DPS) [quest]; Mindthrust Bracers (1974, -1.50 DPS) [dungeon]; Forgotten Wraps (9433, -1.50 DPS) [world_drop] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 56.3 healing_power points (4.46 DPS) | yes | Earthenweave Gloves (254075, -1.18 DPS) [crafted]; Greenleaf Handwraps (19116, -1.31 DPS) [quest]; First Sergeant's Satin Gloves (220898, -1.49 DPS) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | sim-verified (201.2 DPS) | yes | Gilded Cord (254037, +0.00 DPS) [crafted]; Gilded Waistcord (254081, +0.00 DPS) [crafted]; Earthenweave Cord (254077, -0.10 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | sim-verified (201.2 DPS) | yes | Senior Designer's Pantaloons (11841, +0.00 DPS) [dungeon]; Stone Guard's Satin Leggings (220902, -0.04 DPS) [vendor]; Stoneweaver Leggings (9407, -0.83 DPS) [dungeon] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | sim-verified (201.2 DPS) | yes | Earthenweave Boots (254093, -0.33 DPS) [crafted]; Vinerot Sandals (17748, -0.40 DPS) [dungeon]; First Sergeant's Satin Boots (220900, -0.49 DPS) [vendor] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 26.4 healing_power points (2.09 DPS) | yes | Electrocutioner Lagnut (9447, -0.61 DPS) [dungeon]; Runemender's Seal (273031, -0.61 DPS) [world]; Solid Gold Band (279845, -0.67 DPS) [quest] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | sim-verified (201.2 DPS) | yes | Electrocutioner Lagnut (9447, +0.00 DPS) [dungeon]; Runemender's Seal (273031, +0.00 DPS) [world]; Solid Gold Band (279845, +0.00 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.22 DPS) [quest]; Alchemist's Stone (13503, -0.89 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (201.2 DPS) | yes | Evonice's Landin' Pilla (18951, -0.32 DPS) [quest]; Alchemist's Stone (13503, -0.98 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (201.2 DPS) | yes | Spellshifter Rod (9527, +0.00 DPS) [quest]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Charstone Dirk (17710, +0.00 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | sim-verified (201.2 DPS) | yes | Wand of Mana Concentration (273043, +0.00 DPS) [world]; Flash Wand (5248, -0.26 DPS) [quest]; Goblin Igniter (5253, -0.26 DPS) [quest] |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Nethergeld Shoulders; back: Imperial Red Cloak; wrist: Nethergeld Cuffs; hands: Gilded Gloves; waist: Dawnspire Cord; legs: Kilt of the Atal'ai Prophet; feet: Gilded Sandals; finger1: Eye of Adaegus; finger2: Welken Ring; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 533, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60 (undead, 025005031330011530-03505003030000000-00000000000000000)

Set DPS (verified): 294.2. Weights run: 9.2s. Verify run: 44.5s. 1219 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.041, intellect=1.062 ± 0.039, spirit=3.412 ± 0.055, mp5=4.095 ± 0.064, crit=0.343 ± 0.017 per rating point (14 rating = 1%, 4.798 per %), spell_haste=-13.003 ± 0.795

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (294.2 DPS) | yes | Mooncloth Circlet (14140, +0.00 DPS) [crafted]; Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor] |
| neck | Lady Alizabeth's Pendant (13002) | World drop [world_drop] | 51.2 healing_power points (7.58 DPS) | yes | Amberglow Talisman (10824, -2.53 DPS) [quest]; Choker of the High Shaman (4112, -2.75 DPS) [quest]; Necklace of Calisea (1714, -2.94 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (294.2 DPS) | yes | Mooncloth Shoulders (14139, +0.00 DPS) [crafted]; Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Earthenweave Mantle (254119, +0.00 DPS) [crafted] |
| back | Cloak of Blight (6832) (or Indomitable Cloak (14683)) | Nothing But The Truth [quest] | 44.4 healing_power points (6.57 DPS) | yes | Indomitable Cloak (14683, +0.00 DPS) [world_drop]; Wingveil Cloak (10802, -0.27 DPS) [dungeon]; Arcane Cloak (8286, -0.51 DPS) [world_drop] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 127.6 healing_power points (18.88 DPS) | yes | Mooncloth Vest (14138, -2.30 DPS) [crafted]; Legionnaire's Satin Tunic (227124, -3.54 DPS) [pvp]; Warlord's Satin Tunic (231632, -3.87 DPS) [vendor] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (294.2 DPS) | yes | General's Satin Bracers (17619, +0.00 DPS) [pvp]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -1.35 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (294.2 DPS) | yes | Swarmtender's Gloves (275607, +0.00 DPS) [crafted]; Gilded Gloves (254095, -2.47 DPS) [crafted]; Earthenweave Gloves (254075, -2.97 DPS) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 76.6 healing_power points (11.33 DPS) | yes | General's Satin Cinch (17621, -0.14 DPS) [pvp]; Earthenweave Cord (254077, -1.93 DPS) [crafted]; Devout Belt (16696, -2.02 DPS) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 118.4 healing_power points (17.52 DPS) | yes | Legionnaire's Satin Legguards (227123, -2.18 DPS) [pvp]; Mooncloth Leggings (14137, -2.33 DPS) [crafted]; Dalewind Trousers (13008, -3.48 DPS) [world_drop] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (294.2 DPS) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; Elunarian Boots (14458, -2.26 DPS) [world_drop]; Earthenweave Boots (254093, -2.38 DPS) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 48.2 healing_power points (7.13 DPS) | yes | Coldwater Ring (4550, -2.11 DPS) [quest]; Electrocutioner Lagnut (9447, -2.26 DPS) [dungeon]; Runemender's Seal (273031, -2.26 DPS) [world] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 47.3 healing_power points (7.00 DPS) | yes | Coldwater Ring (4550, -1.99 DPS) [quest]; Electrocutioner Lagnut (9447, -2.14 DPS) [dungeon]; Runemender's Seal (273031, -2.14 DPS) [world] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Alchemist's Stone (13503, -1.62 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Alchemist's Stone (13503, -0.64 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (294.2 DPS) | yes | Spire of Hakkar (10844, +0.00 DPS) [world]; Staff of Hale Magefire (13000, +0.00 DPS) [world_drop]; Charstone Dirk (17710, +0.00 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 54.5 healing_power points (8.07 DPS) | yes | Chillnail Splinter (10704, -5.04 DPS) [quest]; Jaina's Firestarter (13064, -5.61 DPS) [world_drop]; Captain Rackmore's Tiller (16789, -5.67 DPS) [quest] |

**New at 60:** head: Devout Crown; neck: Lady Alizabeth's Pendant; shoulder: Devout Mantle; back: Cloak of Blight; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Shard of the Splithooves; ranged: Torch of Light

No-known-source sample (15 of 1219, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60, raid preset (undead, 025005031330011530-03505003030000000-00000000000000000)

Set DPS (verified): 547.8. Weights run: 6.1s. Verify run: 28.7s. 1219 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.117, intellect=0.856 ± 0.039, spirit=1.111 ± 0.039, mp5=2.155 ± 0.042, crit=0.573 ± 0.024 per rating point (14 rating = 1%, 8.026 per %), spell_haste=-6.873 ± 0.846

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (547.8 DPS) | yes | Mooncloth Circlet (14140, +0.00 DPS) [crafted]; Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 26.9 healing_power points (6.49 DPS) | yes | Lady Alizabeth's Pendant (13002, -2.46 DPS) [world_drop]; Horizon Choker (13085, -2.52 DPS) [world_drop]; Necklace of Calisea (1714, -3.16 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (547.8 DPS) | yes | Mooncloth Shoulders (14139, +0.00 DPS) [crafted]; Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp] |
| back | Wingveil Cloak (10802) | Sunken Temple: Spawn of Hakkar [dungeon] | 17.2 healing_power points (4.17 DPS) | yes | Glowing Thresher Cape (6901, -0.32 DPS) [dungeon]; Cloak of Blight (6832, -0.67 DPS) [quest]; Indomitable Cloak (14683, -0.67 DPS) [world_drop] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 90.8 healing_power points (21.95 DPS) | yes | Warlord's Satin Tunic (231632, -3.82 DPS) [vendor]; Legionnaire's Satin Tunic (227124, -5.41 DPS) [pvp]; Stormcloth Vest (10020, -5.78 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (547.8 DPS) | yes | Nethergeld Cuffs (254061, +0.00 DPS) [crafted]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted]; General's Satin Bracers (17619, -0.38 DPS) [pvp] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (547.8 DPS) | yes | Blood Guard's Satin Handwraps (227126, +0.00 DPS) [pvp]; Gilded Gloves (254095, +0.00 DPS) [crafted]; General's Satin Handwraps (231633, -0.56 DPS) [vendor] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 53.0 healing_power points (12.82 DPS) | yes | Earthenweave Cord (254077, -3.57 DPS) [crafted]; Devout Belt (16696, -3.59 DPS) [dungeon]; Gilded Waistcord (254081, -3.94 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 62.4 healing_power points (15.09 DPS) | yes | Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Earthenweave Leggings (254125, -1.84 DPS) [crafted] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 50.4 healing_power points (12.20 DPS) | yes | General's Satin Walkers (231630, -1.46 DPS) [vendor]; Swarmtender's Footpads (275609, -1.80 DPS) [crafted]; Blood Guard's Satin Walkers (227127, -2.80 DPS) [pvp] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 20.8 healing_power points (5.03 DPS) | yes | Electrocutioner Lagnut (9447, -0.97 DPS) [dungeon]; Runemender's Seal (273031, -0.97 DPS) [world]; Solid Gold Band (279845, -1.02 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 18.5 healing_power points (4.47 DPS) | yes | Electrocutioner Lagnut (9447, -0.41 DPS) [dungeon]; Runemender's Seal (273031, -0.41 DPS) [world]; Solid Gold Band (279845, -0.46 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (547.8 DPS) | yes | Ankh of Life (1713, -1.13 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.67 DPS) [quest]; Alchemist's Stone (13503, -3.28 DPS) [crafted] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (547.8 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.44 DPS) [quest]; Alchemist's Stone (13503, -2.05 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (547.8 DPS) | yes | Spire of Hakkar (10844, +0.00 DPS) [world]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Charstone Dirk (17710, +0.00 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 29.2 healing_power points (7.07 DPS) | yes | Captain Rackmore's Tiller (16789, -4.81 DPS) [quest]; Earthen Rod (9381, -4.89 DPS) [dungeon]; Jaina's Firestarter (13064, -5.02 DPS) [world_drop] |

**New at 60:** head: Devout Crown; neck: Glowing Eye of Mordresh; shoulder: Devout Mantle; back: Wingveil Cloak; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Uther's Strength; trinket2: Shard of the Splithooves; ranged: Torch of Light

No-known-source sample (15 of 1219, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

