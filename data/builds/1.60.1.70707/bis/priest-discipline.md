# Leveling BiS: Discipline

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 025004000000000000-00000000000000000-00000000000000000)

Set DPS (verified): 38.0. Weights run: 6.6s. Verify run: 18.0s. 163 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.004, intellect=1.596 ± 0.009, spirit=0.717 ± 0.004, mp5=1.469 ± 0.011, crit=0.055 ± 0.003 per rating point (14 rating = 1%, 0.765 per %), spell_haste=0.120 ± 0.021

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.46 DPS) | yes | Shadow Goggles (4373, -0.73 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.23 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 2.9 healing_power points (0.23 DPS) | yes | Scholarly Pendant (277203, -0.06 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 14.4 healing_power points (1.17 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.81 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.81 DPS) [dungeon] |
| back | Sanguine Cape (14376) | World drop [world_drop] | 6.4 healing_power points (0.52 DPS) | yes | Pearl-clasped Cloak (5542, -0.13 DPS) [crafted]; Seer's Cape (6378, -0.14 DPS) [dungeon]; Forest Cloak (4710, -0.26 DPS) [world_drop] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 23.1 healing_power points (1.88 DPS) | yes | Corsair's Overshirt (5202, -0.43 DPS, sim-verified) [dungeon]; Robe of the Moccasin (6465, -0.57 DPS) [dungeon]; Bloody Apron (6226, -0.82 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 12.4 healing_power points (1.01 DPS) | yes | Repurposed Hair Band (281256, -0.63 DPS) [quest]; Bright Bracers (3647, -0.74 DPS, sim-verified) [world_drop]; Mystic's Bracelets (14366, -0.75 DPS) [world_drop] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 15.8 healing_power points (1.28 DPS) | yes | Tomb Robber's Gloves (280096, -0.51 DPS) [quest]; Bright Gloves (3066, -0.53 DPS) [world_drop]; Magefist Gloves (12977, -0.65 DPS, sim-verified) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 17.4 healing_power points (1.41 DPS) | yes | Keller's Girdle (2911, -0.38 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.39 DPS, sim-verified) [crafted]; Wise Man's Belt (4786, -0.77 DPS) [vendor] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 30.4 healing_power points (2.48 DPS) | yes | Filigreed Silky Leggings (253939, -1.46 DPS) [crafted]; Filigreed Shadow Leggings (253943, -1.46 DPS) [crafted]; Darkweave Breeches (12987, -1.48 DPS, sim-verified) [world_drop] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 13.8 healing_power points (1.12 DPS) | yes | Spidersilk Boots (4320, -0.60 DPS) [crafted]; Sanguine Sandals (14374, -0.60 DPS) [world_drop]; Kimbra Boots (6191, -0.78 DPS, sim-verified) [quest] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 9.6 healing_power points (0.78 DPS) | yes | Band of Purification (12996, -0.43 DPS) [world_drop]; Deep Fathom Ring (6463, -0.49 DPS) [dungeon]; Minor Channeling Ring (1449, -0.52 DPS) [quest] |
| finger2 | Volcanic Rock Ring (12053) | World drop [world_drop] | 4.8 healing_power points (0.39 DPS) | yes | Band of Purification (12996, -0.04 DPS) [world_drop]; Deep Fathom Ring (6463, -0.10 DPS) [dungeon]; Minor Channeling Ring (1449, -0.13 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 16.0 healing_power points (1.30 DPS) | yes | Staff of Westfall (2042, -0.25 DPS, sim-verified) [quest]; Lesser Staff of the Spire (1300, -0.52 DPS) [world]; Witching Stave (1484, -0.65 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Flaring Baton (5326) (or Moonstone Wand (15204)) | The Escape [quest] | 3.2 healing_power points (0.26 DPS) | yes | Moonstone Wand (15204, +0.00 DPS) [quest]; Sable Wand (7607, -0.08 DPS) [quest]; Elven Wand (5604, -0.13 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Tarnished Locket; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Volcanic Rock Ring; main_hand: Twisted Chanter's Staff; ranged: Flaring Baton

No-known-source sample (15 of 163, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes; 16607 Acolyte's Sacrificial Robes; 18851 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (human, 025005031320000000-00000000000000000-00000000000000000)

Set DPS (verified): 71.1. Weights run: 6.9s. Verify run: 20.2s. 294 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.413 ± 0.008, spirit=1.048 ± 0.007, mp5=2.333 ± 0.009, crit=0.091 ± 0.004 per rating point (14 rating = 1%, 1.274 per %), spell_haste=-0.201 ± 0.039

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 39.3 healing_power points (3.78 DPS) | yes | Filigreed Pristine Circlet (253975, -1.66 DPS) [crafted]; Resilient Cap (14401, -1.88 DPS) [world_drop]; Nightsky Cowl (4039, -2.57 DPS, sim-verified) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 9.8 healing_power points (0.95 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.13 DPS) [world]; Crystal Starfire Medallion (5003, -0.20 DPS, sim-verified) [world_drop]; Pendant of Myzrael (4614, -0.34 DPS) [dungeon] |
| shoulder | Mantle of Honor (3560) | Bride of the Embalmer [quest] | 17.2 healing_power points (1.66 DPS) | yes | Nightsky Mantle (4718, -0.09 DPS, sim-verified) [world_drop]; Death Speaker Mantle (6685, -0.16 DPS) [dungeon]; Faerie Mantle (5820, -0.24 DPS) [quest] |
| back | Prelacy Cape (7004) | Researching the Corruption [quest] | 17.3 healing_power points (1.66 DPS) | yes | Glowing Thresher Cape (6901, -0.16 DPS, sim-verified) [dungeon]; Repairman's Cape (9605, -0.37 DPS) [quest]; Sergeant Major's Cape (16315, -0.72 DPS) [pvp] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 35.5 healing_power points (3.42 DPS) | yes | Pristine Gown (253961, -0.20 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -0.98 DPS) [dungeon]; Robes of Arugal (6324, -1.08 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 14.1 healing_power points (1.35 DPS) | yes | Glowing Magical Bracelets (13106, -0.27 DPS) [world_drop]; Nightsky Wristbands (6407, -0.58 DPS, sim-verified) [world_drop]; Stonecloth Bindings (14416, -0.67 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 19.6 healing_power points (1.89 DPS) | yes | Truefaith Gloves (7049, +0.00 DPS, sim-verified) [crafted]; Hotshot Pilot's Gloves (9491, -0.30 DPS) [dungeon]; Pristine Gloves (253913, -0.42 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.7 healing_power points (1.60 DPS) | yes | Novice Ardent's Sash (253887, -0.29 DPS, sim-verified) [crafted]; Resilient Cord (14406, -0.38 DPS) [world_drop]; Dreamer's Belt (4829, -0.45 DPS) [vendor] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 37.1 healing_power points (3.57 DPS) | yes | Filigreed Pristine Leggings (253937, -0.76 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -1.02 DPS) [crafted]; Darkweave Breeches (12987, -2.02 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 34.1 healing_power points (3.28 DPS) | yes | Acidic Walkers (9454, -1.79 DPS) [dungeon]; Frothing Slippers (254003, -1.93 DPS) [crafted]; Nimbus Boots (6998, -2.05 DPS, sim-verified) [quest] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 13.3 healing_power points (1.28 DPS) | yes | The Queen's Jewel (13094, -0.20 DPS) [world_drop]; Defias Mage Ring (1077, -0.44 DPS) [world]; Seal of Wrynn (2933, -0.44 DPS) [quest] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 13.3 healing_power points (1.28 DPS) | yes | The Queen's Jewel (13094, -0.38 DPS, sim-verified) [world_drop]; Defias Mage Ring (1077, -0.44 DPS) [world]; Seal of Wrynn (2933, -0.44 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-verified (71.1 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -0.63 DPS) [quest]; Death Speaker Scepter (2816, -2.09 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 10.3 healing_power points (0.99 DPS) | yes | Lesser Mystic Wand (11289, -0.45 DPS) [crafted]; Starfaller (13063, -0.45 DPS) [world_drop]; Summoner's Wand (5245, -0.81 DPS, sim-verified) [world] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Mantle of Honor; back: Prelacy Cape; chest: Death Speaker Robes; hands: Gloves of Old; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Spine of the Devourer; ranged: Wand of Mana Concentration

No-known-source sample (15 of 294, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 40 (human, 025005031330011520-00000000000000000-00000000000000000)

Set DPS (verified): 135.7. Weights run: 8.0s. Verify run: 28.1s. 396 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.018, intellect=1.579 ± 0.023, spirit=0.376 ± 0.036, mp5=2.986 ± 0.025, crit=0.238 ± 0.008 per rating point (14 rating = 1%, 3.335 per %), spell_haste=-2.938 ± 0.299

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 44.2 healing_power points (4.54 DPS) | yes | Holy Shroud (2721, -0.92 DPS) [world_drop]; Whitemane's Chapeau (7720, -1.12 DPS) [dungeon]; Thinking Cap (2624, -1.78 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 25.6 healing_power points (2.63 DPS) | yes | Necklace of Calisea (1714, -1.23 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.66 DPS) [world]; Kaleidoscope Chain (13084, -1.83 DPS) [world_drop] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 35.5 healing_power points (3.65 DPS) | yes | Earthen Silk Shoulders (254033, -1.49 DPS) [crafted]; Windchaser Amice (14432, -1.54 DPS) [world_drop]; Mistscape Mantle (4734, -1.67 DPS) [dungeon] |
| back | Blackforge Cape (6424) | Uldaman: Ancient Treasure [dungeon] | 14.1 healing_power points (1.45 DPS) | yes | Prelacy Cape (7004, -0.09 DPS) [quest]; Sergeant Major's Cape (16336, -0.25 DPS) [pvp]; Glyphed Cloak (4732, -0.32 DPS) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 61.0 healing_power points (6.26 DPS) | yes | Death Speaker Robes (6682, -2.43 DPS) [dungeon]; Pristine Gown (253961, -2.92 DPS) [crafted]; Red Mageweave Vest (10007, -3.34 DPS) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | sim-verified (135.7 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS) [dungeon]; Mistscape Bracers (4045, -0.16 DPS) [dungeon]; Enchanted Stonecloth Bracers (4979, -0.16 DPS) [quest] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (135.7 DPS) | yes | Gilded Handwraps (254021, +0.00 DPS) [crafted]; Earthen Silk Gloves (254017, -0.92 DPS) [crafted]; Truefaith Gloves (7049, -1.05 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 36.9 healing_power points (3.79 DPS) | yes | Deathmage Sash (10771, -1.35 DPS) [dungeon]; Razzeric's Customized Seatbelt (6726, -1.84 DPS) [quest]; Teacher's Sash (10747, -1.84 DPS) [quest] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-verified (135.7 DPS) | yes | Stoneweaver Leggings (9407, +0.00 DPS) [dungeon]; Filigreed Pristine Leggings (253937, +0.00 DPS) [crafted]; Pristine Leggings (253987, +0.00 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 32.6 healing_power points (3.34 DPS) | yes | Furen's Boots (13100, -1.59 DPS) [world_drop]; Nimbus Boots (6998, -1.60 DPS) [quest]; Kodo Rustler Boots (15697, -1.85 DPS) [quest] |
| finger1 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 12.2 healing_power points (1.25 DPS) | yes | Mindbender Loop (5009, -0.04 DPS) [world_drop]; Welken Ring (5011, -0.17 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.28 DPS) [dungeon] |
| finger2 | Ogremind Ring (1993) | World drop [world_drop] | 12.2 healing_power points (1.25 DPS) | yes | Mindbender Loop (5009, -0.04 DPS) [world_drop]; Welken Ring (5011, -0.17 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.28 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, -2.98 DPS) [dungeon]; Hypnotic Blade (7714, -4.03 DPS) [dungeon]; Illusionary Rod (7713, -4.20 DPS) [dungeon] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 27.8 healing_power points (2.85 DPS) | yes | Beacon of Hope (9393, -1.26 DPS) [dungeon]; Orb of Souls (249395, -1.29 DPS) [crafted]; Eye of Paleth (2943, -1.52 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | sim-verified (135.7 DPS) | yes | Wand of Mana Concentration (273043, +0.00 DPS) [world]; Flash Wand (5248, -0.32 DPS) [quest]; Goblin Igniter (5253, -0.32 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Blackforge Cape; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Stormcloth Gloves; waist: Gilded Cord; legs: Stormcloth Pants; finger1: Voodoo Band; finger2: Ogremind Ring; trinket1: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Orb of Lorica; ranged: Jaina's Firestarter

No-known-source sample (15 of 396, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 50 (human, 025005031330011530-03501000000000000-00000000000000000)

Set DPS (verified): 224.2. Weights run: 8.1s. Verify run: 26.0s. 515 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.028, intellect=1.601 ± 0.027, spirit=1.154 ± 0.038, mp5=3.342 ± 0.054, crit=0.513 ± 0.016 per rating point (14 rating = 1%, 7.178 per %), spell_haste=-9.501 ± 0.457

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 57.0 healing_power points (5.28 DPS) | yes | Knight-Lieutenant's Satin Cover (220896, +0.00 DPS) [vendor]; Papal Fez (9431, -0.50 DPS) [dungeon]; Cassandra's Grace (13102, -1.87 DPS, sim-verified) [world_drop] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 27.0 healing_power points (2.50 DPS) | yes | Necklace of Calisea (1714, -0.72 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -0.80 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -1.08 DPS) [quest] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | 45.7 healing_power points (4.24 DPS) | yes | Knight-Lieutenant's Satin Pads (220894, +0.00 DPS) [vendor]; Rotgrip Mantle (17732, -1.57 DPS) [dungeon]; Inquisitor's Shawl (19507, -2.34 DPS, sim-verified) [dungeon] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 22.2 healing_power points (2.06 DPS) | yes | Prelacy Cape (7004, -0.40 DPS) [quest]; Wingveil Cloak (10802, -0.42 DPS) [dungeon]; Highborne Cloak (14450, -0.54 DPS, sim-verified) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 67.2 healing_power points (6.23 DPS) | yes | Ghostweave Vest (14141, -1.61 DPS) [crafted]; Grimlok's Tribal Vestments (9415, -2.23 DPS) [dungeon]; Knight's Satin Armor (220892, -2.51 DPS, sim-verified) [vendor] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 37.0 healing_power points (3.43 DPS) | yes | Forgotten Wraps (9433, -1.65 DPS) [world_drop]; Mindthrust Bracers (1974, -1.76 DPS) [dungeon]; Shizzle's Nozzle Wiper (11917, -2.68 DPS, sim-verified) [quest] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 54.2 healing_power points (5.03 DPS) | yes | Sergeant Major's Satin Gloves (220897, -1.56 DPS) [vendor]; Gilded Handwraps (254021, -1.95 DPS) [crafted]; Earthenweave Gloves (254075, -2.79 DPS, sim-verified) [crafted] |
| waist | Gilded Waistcord (254081) | Tailoring [crafted] | 43.4 healing_power points (4.02 DPS) | yes | Gilded Cord (254037, -0.16 DPS) [crafted]; Dawnspire Cord (12466, -0.27 DPS, sim-verified) [dungeon]; Earthenweave Cord (254077, -0.44 DPS) [crafted] |
| legs | Knight's Satin Leggings (220893) | Captain Dirgehammer [vendor] | 58.2 healing_power points (5.40 DPS) | yes | Senior Designer's Pantaloons (11841, +0.00 DPS, sim-verified) [dungeon]; Kilt of the Atal'ai Prophet (10807, -0.24 DPS) [dungeon]; Dalewind Trousers (13008, -1.39 DPS) [world_drop] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 45.4 healing_power points (4.21 DPS) | yes | Earthenweave Boots (254093, -0.62 DPS) [crafted]; Vinerot Sandals (17748, -0.69 DPS) [dungeon]; Sergeant Major's Satin Boots (220895, -1.02 DPS, sim-verified) [vendor] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 23.5 healing_power points (2.17 DPS) | yes | Coldwater Ring (4550, -0.77 DPS) [quest]; Choking Band (11868, -0.78 DPS) [quest]; Voodoo Band (1996, -0.81 DPS) [world] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | 16.1 healing_power points (1.49 DPS) | yes | Coldwater Ring (4550, -0.08 DPS) [quest]; Choking Band (11868, -0.10 DPS) [quest]; Voodoo Band (1996, -0.13 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (224.2 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -0.43 DPS) [quest]; Alchemist's Stone (13503, -0.86 DPS) [crafted] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-verified (224.2 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -0.52 DPS, sim-verified) [quest]; Alchemist's Stone (13503, -0.64 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (224.2 DPS) | yes | Charstone Dirk (17710, -6.16 DPS) [dungeon]; Hand of Righteousness (7721, -9.60 DPS) [dungeon]; Hammer of the Northern Wind (810, -33.43 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cairnstone Sliver (9654) | The Morrow Stone [quest] | 13.8 healing_power points (1.28 DPS) | yes | Jaina's Firestarter (13064, -0.07 DPS) [world_drop]; Goblin Igniter (5253, -0.36 DPS) [quest]; Wand of Mana Concentration (273043, -0.41 DPS, sim-verified) [world] |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Nethergeld Shoulders; back: Imperial Red Cloak; wrist: Nethergeld Cuffs; hands: Gilded Gloves; waist: Gilded Waistcord; legs: Knight's Satin Leggings; feet: Gilded Sandals; finger1: Eye of Adaegus; finger2: Welken Ring; trinket2: Evonice's Landin' Pilla; main_hand: Glowing Brightwood Staff; ranged: Cairnstone Sliver

No-known-source sample (15 of 515, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60 (human, 025005031330011530-03505003030000000-00000000000000000)

Set DPS (verified): 323.4. Weights run: 7.8s. Verify run: 29.5s. 1200 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.073, intellect=1.442 ± 0.028, spirit=4.145 ± 0.036, mp5=1.827 ± 0.037, crit=0.364 ± 0.017 per rating point (14 rating = 1%, 5.091 per %), spell_haste=-6.962 ± 0.558

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | 116.8 healing_power points (18.74 DPS) | yes | Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Earthenweave Crown (254137, -5.40 DPS, sim-verified) [crafted] |
| neck | Lady Alizabeth's Pendant (13002) | World drop [world_drop] | 62.2 healing_power points (9.98 DPS) | yes | Necklace of Calisea (1714, -3.70 DPS) [dungeon]; Explorers' League Commendation (7746, -3.99 DPS) [quest]; Amberglow Talisman (10824, -5.67 DPS, sim-verified) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (+4.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Earthenweave Mantle (254119, -0.46 DPS) [crafted]; Mooncloth Shoulders (14139, -4.75 DPS, sim-verified) [crafted] |
| back | Ethereal Mist Cape (11873) | Ancient Spirit [quest] | 54.1 healing_power points (8.68 DPS) | yes | Wingveil Cloak (10802, -0.56 DPS) [dungeon]; Arcane Cloak (8286, -0.69 DPS) [world_drop]; Indomitable Cloak (14683, -1.90 DPS, sim-verified) [world_drop] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 139.3 healing_power points (22.36 DPS) | yes | Knight-Captain's Satin Tunic (227122, -3.16 DPS) [pvp]; Field Marshal's Satin Tunic (231624, -3.31 DPS) [vendor]; Mooncloth Vest (14138, -8.21 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | 64.0 healing_power points (10.27 DPS) | yes | Marshal's Satin Bracers (17606, +0.00 DPS) [pvp]; Earthenweave Cuffs (254131, +0.00 DPS, sim-verified) [crafted]; Nethergeld Cuffs (254061, -2.11 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | 99.9 healing_power points (16.03 DPS) | yes | Magister's Gloves (16684, -3.48 DPS) [dungeon]; Gilded Gloves (254095, -3.74 DPS) [crafted]; Swarmtender's Gloves (275607, -6.28 DPS, sim-verified) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (+3.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Satin Sash (17609, +0.00 DPS) [pvp]; Mana-infused Cord (279267, -0.41 DPS) [crafted]; Wisdom of the Timbermaw (19047, -3.48 DPS, sim-verified) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 141.0 healing_power points (22.62 DPS) | yes | Knight-Captain's Satin Legguards (227125, -3.42 DPS) [pvp]; Dalewind Trousers (13008, -3.91 DPS) [world_drop]; Mooncloth Leggings (14137, -4.94 DPS, sim-verified) [crafted] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 98.1 healing_power points (15.74 DPS) | yes | Devout Sandals (16691, +0.00 DPS, sim-verified) [dungeon]; Elunarian Boots (14458, -3.04 DPS) [world_drop]; Furen's Boots (13100, -3.48 DPS) [world_drop] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 60.0 healing_power points (9.63 DPS) | yes | Choking Band (11868, -0.98 DPS) [quest]; Coldwater Ring (4550, -2.95 DPS) [quest]; Seedtime Hoop (9655, -2.98 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 58.4 healing_power points (9.37 DPS) | yes | Coldwater Ring (4550, -2.69 DPS) [quest]; Seedtime Hoop (9655, -2.72 DPS) [quest]; Choking Band (11868, -2.89 DPS, sim-verified) [quest] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Mark of the Chosen (17774, -4.17 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -15.02 DPS) [dungeon]; Spire of Hakkar (10844, -16.17 DPS) [world]; Hand of Edward the Odd (2243, -19.19 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 62.6 healing_power points (10.05 DPS) | yes | Jaina's Firestarter (13064, -6.66 DPS) [world_drop]; Captain Rackmore's Tiller (16789, -7.09 DPS) [quest]; Cairnstone Sliver (9654, -8.99 DPS, sim-verified) [quest] |

**New at 60:** head: Devout Crown; neck: Lady Alizabeth's Pendant; shoulder: Devout Mantle; back: Ethereal Mist Cape; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Shard of the Splithooves; trinket2: Uther's Strength; ranged: Torch of Light

No-known-source sample (15 of 1200, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60, raid preset (human, 025005031330011530-03505003030000000-00000000000000000)

Set DPS (verified): 593.3. Weights run: 5.2s. Verify run: 17.6s. 1200 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.109, intellect=1.218 ± 0.030, spirit=1.118 ± 0.026, mp5=2.481 ± 0.033, crit=0.604 ± 0.022 per rating point (14 rating = 1%, 8.453 per %), spell_haste=-4.887 ± 0.556

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (593.3 DPS) | yes | Mooncloth Circlet (14140, +0.00 DPS) [crafted]; Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 22.7 healing_power points (5.50 DPS) | yes | Horizon Choker (13085, -0.29 DPS) [world_drop]; Lady Alizabeth's Pendant (13002, -1.44 DPS) [world_drop]; Necklace of Calisea (1714, -1.55 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (593.3 DPS) | yes | Mooncloth Shoulders (14139, +0.00 DPS) [crafted]; Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp] |
| back | Sergeant Major's Cape (16337) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 18.7 healing_power points (4.52 DPS) | yes | Imperial Red Cloak (8248, -0.20 DPS) [world_drop]; Prelacy Cape (7004, -0.24 DPS) [quest]; Wingveil Cloak (10802, -0.34 DPS) [dungeon] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 90.9 healing_power points (21.98 DPS) | yes | Field Marshal's Satin Tunic (231624, -1.38 DPS) [vendor]; Knight-Captain's Satin Tunic (227122, -4.10 DPS) [pvp]; Mooncloth Vest (14138, -4.66 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (593.3 DPS) | yes | Marshal's Satin Bracers (17606, +0.00 DPS) [pvp]; Nethergeld Cuffs (254061, +0.00 DPS) [crafted]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (593.3 DPS) | yes | Gilded Gloves (254095, +0.00 DPS) [crafted]; Knight-Lieutenant's Satin Handwraps (227128, -0.02 DPS) [pvp]; Marshal's Satin Handwraps (231623, -0.33 DPS) [vendor] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 58.1 healing_power points (14.05 DPS) | yes | Devout Belt (16696, -3.07 DPS) [dungeon]; Marshal's Satin Sash (17609, -3.22 DPS) [pvp]; Gilded Waistcord (254081, -4.39 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 68.0 healing_power points (16.44 DPS) | yes | Knight-Captain's Satin Legguards (227125, +0.00 DPS) [pvp]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Earthenweave Leggings (254125, -3.17 DPS) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (593.3 DPS) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; Knight-Lieutenant's Satin Walkers (227129, +0.00 DPS) [pvp]; Marshal's Satin Walkers (231627, +0.00 DPS) [vendor] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 24.5 healing_power points (5.92 DPS) | yes | Maiden's Circle (13001, -2.24 DPS) [world_drop]; Glowing Crystal Ring (18402, -2.39 DPS) [quest]; Choking Band (11868, -2.41 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 20.7 healing_power points (5.01 DPS) | yes | Maiden's Circle (13001, -1.33 DPS) [world_drop]; Glowing Crystal Ring (18402, -1.48 DPS) [quest]; Choking Band (11868, -1.50 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (593.3 DPS) | yes | Shard of the Splithooves (10659, -0.75 DPS) [quest]; Evonice's Landin' Pilla (18951, -1.65 DPS) [quest]; Thunderbrew's Boot Flask (744, -2.19 DPS) [quest] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (593.3 DPS) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Evonice's Landin' Pilla (18951, -0.54 DPS) [quest]; Thunderbrew's Boot Flask (744, -1.08 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -14.76 DPS) [dungeon]; Hand of Edward the Odd (2243, -18.95 DPS) [world_drop]; Hand of Righteousness (7721, -22.35 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 29.3 healing_power points (7.08 DPS) | yes | Cairnstone Sliver (9654, -4.26 DPS) [quest]; Jaina's Firestarter (13064, -4.51 DPS) [world_drop]; Wand of Mana Concentration (273043, -4.71 DPS) [world] |

**New at 60:** head: Devout Crown; neck: Glowing Eye of Mordresh; shoulder: Devout Mantle; back: Sergeant Major's Cape; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Uther's Strength; trinket2: Ankh of Life; ranged: Torch of Light

No-known-source sample (15 of 1200, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

## Horde

### Band 20 (undead, 025004000000000000-00000000000000000-00000000000000000)

Set DPS (verified): 38.1. Weights run: 6.6s. Verify run: 17.9s. 156 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.004, intellect=1.596 ± 0.009, spirit=0.717 ± 0.004, mp5=1.469 ± 0.011, crit=0.055 ± 0.003 per rating point (14 rating = 1%, 0.765 per %), spell_haste=0.120 ± 0.021

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.46 DPS) | yes | Shadow Goggles (4373, -1.07 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.23 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 2.1 healing_power points (0.17 DPS) | yes | Scholarly Pendant (277203, -0.36 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 14.4 healing_power points (1.17 DPS) | yes | Slime-encrusted Pads (6461, -0.81 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -0.88 DPS, sim-verified) [crafted] |
| back | Sanguine Cape (14376) | World drop [world_drop] | 6.4 healing_power points (0.52 DPS) | yes | Seer's Cape (6378, -0.14 DPS) [dungeon]; Forest Cloak (4710, -0.26 DPS) [world_drop]; Pearl-clasped Cloak (5542, -0.30 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 23.1 healing_power points (1.88 DPS) | yes | Robe of the Moccasin (6465, -0.57 DPS) [dungeon]; Bloody Apron (6226, -0.82 DPS) [dungeon]; Corsair's Overshirt (5202, -0.90 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 12.4 healing_power points (1.01 DPS) | yes | Bright Bracers (3647, -0.49 DPS) [world_drop]; Crystalline Cuffs (14148, -0.57 DPS) [dungeon]; Featherbead Bracers (15452, -0.86 DPS, sim-verified) [quest] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 15.8 healing_power points (1.28 DPS) | yes | Magefist Gloves (12977, -0.40 DPS) [world_drop]; Bright Gloves (3066, -0.53 DPS) [world_drop]; Blight Gloves (279877, -0.59 DPS, sim-verified) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 17.4 healing_power points (1.41 DPS) | yes | Keller's Girdle (2911, -0.38 DPS) [world_drop]; Wise Man's Belt (4786, -0.77 DPS) [vendor]; Novice Ardent's Sash (253887, -0.77 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 30.4 healing_power points (2.48 DPS) | yes | Filigreed Silky Leggings (253939, -1.46 DPS) [crafted]; Filigreed Shadow Leggings (253943, -1.46 DPS) [crafted]; Darkweave Breeches (12987, -2.00 DPS, sim-verified) [world_drop] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 13.8 healing_power points (1.12 DPS) | yes | Spidersilk Boots (4320, -0.60 DPS) [crafted]; Bluegill Sandals (1560, -0.63 DPS) [world]; Sanguine Sandals (14374, -1.20 DPS, sim-verified) [world_drop] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 9.6 healing_power points (0.78 DPS) | yes | Volcanic Rock Ring (12053, -0.39 DPS) [world_drop]; Band of Purification (12996, -0.43 DPS) [world_drop]; Deep Fathom Ring (6463, -0.49 DPS) [dungeon] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 8.0 healing_power points (0.65 DPS) | yes | Band of Purification (12996, -0.30 DPS) [world_drop]; Deep Fathom Ring (6463, -0.36 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.44 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 16.0 healing_power points (1.30 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Staff of Orgrimmar (15444, -0.25 DPS) [quest]; Lesser Staff of the Spire (1300, -0.52 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Decay (5252) (or Flaring Baton (5326)) | Beren's Peril [quest] | 3.2 healing_power points (0.26 DPS) | yes | Flaring Baton (5326, +0.00 DPS) [quest]; Wisesight Wand (286750, -0.13 DPS) [world] |

**New at 20:** head: Pristine Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Gnarled Necromancer's Staff; ranged: Wand of Decay

No-known-source sample (15 of 156, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (undead, 025005031320000000-00000000000000000-00000000000000000)

Set DPS (verified): 69.9. Weights run: 6.9s. Verify run: 20.2s. 286 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.413 ± 0.008, spirit=1.048 ± 0.007, mp5=2.333 ± 0.009, crit=0.091 ± 0.004 per rating point (14 rating = 1%, 1.274 per %), spell_haste=-0.201 ± 0.039

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 39.3 healing_power points (3.78 DPS) | yes | Filigreed Pristine Circlet (253975, -1.66 DPS) [crafted]; Resilient Cap (14401, -1.88 DPS) [world_drop]; Nightsky Cowl (4039, -2.21 DPS, sim-verified) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 9.8 healing_power points (0.95 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.13 DPS) [world]; Crystal Starfire Medallion (5003, -0.25 DPS, sim-verified) [world_drop]; Pendant of Myzrael (4614, -0.34 DPS) [dungeon] |
| shoulder | Ghostly Mantle (3324) | Deathstalkers in Shadowfang [quest] | 18.4 healing_power points (1.77 DPS) | yes | Nightsky Mantle (4718, -0.19 DPS, sim-verified) [world_drop]; Death Speaker Mantle (6685, -0.28 DPS) [dungeon]; Talbar Mantle (10657, -0.51 DPS) [quest] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 15.4 healing_power points (1.48 DPS) | yes | Resilient Cape (14400, -0.62 DPS, sim-verified) [world_drop]; Silver-thread Cloak (4713, -0.73 DPS) [world_drop]; Amy's Blanket (13005, -0.77 DPS) [world_drop] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 35.5 healing_power points (3.42 DPS) | yes | Pristine Gown (253961, -0.06 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -0.98 DPS) [dungeon]; Robes of Arugal (6324, -1.08 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 14.1 healing_power points (1.35 DPS) | yes | Glowing Magical Bracelets (13106, -0.27 DPS) [world_drop]; Nightsky Wristbands (6407, -0.39 DPS, sim-verified) [world_drop]; Featherbead Bracers (15452, -0.67 DPS) [quest] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 19.6 healing_power points (1.89 DPS) | yes | Truefaith Gloves (7049, +0.00 DPS, sim-verified) [crafted]; Hotshot Pilot's Gloves (9491, -0.30 DPS) [dungeon]; Tattered Mittens (270030, -0.32 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.7 healing_power points (1.60 DPS) | yes | Lilac Sash (6780, -0.19 DPS, sim-verified) [quest]; Novice Ardent's Sash (253887, -0.31 DPS) [crafted]; Resilient Cord (14406, -0.38 DPS) [world_drop] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 37.1 healing_power points (3.57 DPS) | yes | Filigreed Pristine Leggings (253937, -0.72 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -1.02 DPS) [crafted]; Sacred Burial Trousers (6282, -1.99 DPS) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 34.1 healing_power points (3.28 DPS) | yes | Frothing Slippers (254003, -1.93 DPS) [crafted]; Black Slippers (254007, -1.93 DPS) [crafted]; Acidic Walkers (9454, -2.03 DPS, sim-verified) [dungeon] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 13.3 healing_power points (1.28 DPS) | yes | Solid Gold Band (279845, -0.11 DPS) [quest]; The Queen's Jewel (13094, -0.20 DPS) [world_drop]; Defias Mage Ring (1077, -0.44 DPS) [world] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 13.3 healing_power points (1.28 DPS) | yes | Solid Gold Band (279845, -0.09 DPS, sim-verified) [quest]; The Queen's Jewel (13094, -0.20 DPS) [world_drop]; Defias Mage Ring (1077, -0.44 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-verified (69.9 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Staff of the Friar (3415, -1.85 DPS) [dungeon]; Death Speaker Scepter (2816, -2.20 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 10.3 healing_power points (0.99 DPS) | yes | Lesser Mystic Wand (11289, -0.45 DPS) [crafted]; Starfaller (13063, -0.45 DPS) [world_drop]; Summoner's Wand (5245, -0.62 DPS, sim-verified) [world] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Ghostly Mantle; back: Glowing Thresher Cape; chest: Death Speaker Robes; hands: Gloves of Old; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Spine of the Devourer; ranged: Wand of Mana Concentration

No-known-source sample (15 of 286, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (undead, 025005031330011520-00000000000000000-00000000000000000)

Set DPS (verified): 133.5. Weights run: 8.0s. Verify run: 46.6s. 388 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.018, intellect=1.579 ± 0.023, spirit=0.376 ± 0.036, mp5=2.986 ± 0.025, crit=0.238 ± 0.008 per rating point (14 rating = 1%, 3.335 per %), spell_haste=-2.938 ± 0.299

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 44.2 healing_power points (4.54 DPS) | yes | Whitemane's Chapeau (7720, -1.12 DPS) [dungeon]; Holy Shroud (2721, -1.49 DPS, sim-verified) [world_drop]; Thinking Cap (2624, -1.78 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 25.6 healing_power points (2.63 DPS) | yes | Necklace of Calisea (1714, -0.64 DPS, sim-verified) [dungeon]; Archmagister's Faceted Pendant (273049, -1.66 DPS) [world]; Kaleidoscope Chain (13084, -1.83 DPS) [world_drop] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 35.5 healing_power points (3.65 DPS) | yes | Earthen Silk Shoulders (254033, -0.89 DPS, sim-verified) [crafted]; Windchaser Amice (14432, -1.54 DPS) [world_drop]; Mistscape Mantle (4734, -1.67 DPS) [dungeon] |
| back | Blackforge Cape (6424) | Uldaman: Ancient Treasure [dungeon] | 14.1 healing_power points (1.45 DPS) | yes | Windchaser Cloak (14430, -0.32 DPS) [world_drop]; Glowing Thresher Cape (6901, -0.42 DPS) [dungeon]; Glyphed Cloak (4732, -1.07 DPS, sim-verified) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 61.0 healing_power points (6.26 DPS) | yes | Pristine Gown (253961, -2.92 DPS) [crafted]; Red Mageweave Vest (10007, -3.34 DPS) [crafted]; Death Speaker Robes (6682, -7.89 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 16.9 healing_power points (1.73 DPS) | yes | Windchaser Cuffs (14429, -0.43 DPS, sim-verified) [world_drop]; Mistscape Bracers (4045, -0.43 DPS) [dungeon]; Enchanted Stonecloth Bracers (4979, -0.43 DPS) [quest] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (133.5 DPS) | yes | Earthen Silk Gloves (254017, -0.92 DPS) [crafted]; Truefaith Gloves (7049, -1.05 DPS) [crafted]; Gilded Handwraps (254021, -1.48 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 36.9 healing_power points (3.79 DPS) | yes | Razzeric's Customized Seatbelt (6726, -1.84 DPS) [quest]; Sutarn's Ring (13105, -1.93 DPS) [world_drop]; Deathmage Sash (10771, -3.07 DPS, sim-verified) [dungeon] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-verified (133.5 DPS) | yes | Stoneweaver Leggings (9407, +0.00 DPS) [dungeon]; Filigreed Pristine Leggings (253937, +0.00 DPS) [crafted]; Pristine Leggings (253987, -2.40 DPS, sim-verified) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 32.6 healing_power points (3.34 DPS) | yes | Furen's Boots (13100, -0.73 DPS, sim-verified) [world_drop]; Boots of the Maharishi (9658, -1.65 DPS) [quest]; Kodo Rustler Boots (15697, -1.85 DPS) [quest] |
| finger1 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 12.2 healing_power points (1.25 DPS) | yes | Mindbender Loop (5009, -0.04 DPS) [world_drop]; Welken Ring (5011, -0.17 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.28 DPS) [dungeon] |
| finger2 | Ogremind Ring (1993) | World drop [world_drop] | 12.2 healing_power points (1.25 DPS) | yes | Welken Ring (5011, -0.17 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.28 DPS) [dungeon]; Mindbender Loop (5009, -0.35 DPS, sim-verified) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (133.5 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Death Speaker Scepter (2816, -2.98 DPS) [dungeon]; Hypnotic Blade (7714, -4.03 DPS) [dungeon] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | 15.5 healing_power points (1.59 DPS) | yes | Orb of Souls (249395, +0.00 DPS, sim-verified) [crafted]; Witch's Finger (16887, -0.29 DPS) [quest]; Aurora Sphere (7610, -0.30 DPS) [dungeon] |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 12.3 healing_power points (1.26 DPS) | yes | Jaina's Firestarter (13064, +0.00 DPS, sim-verified) [world_drop]; Flash Wand (5248, -0.50 DPS) [quest]; Goblin Igniter (5253, -0.50 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Blackforge Cape; chest: Stormcloth Vest; hands: Stormcloth Gloves; waist: Gilded Cord; legs: Stormcloth Pants; finger1: Voodoo Band; finger2: Ogremind Ring; trinket1: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Beacon of Hope

No-known-source sample (15 of 388, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16336 Sergeant Major's Cape

### Band 50 (undead, 025005031330011530-03501000000000000-00000000000000000)

Set DPS (verified): 220.8. Weights run: 8.1s. Verify run: 29.4s. 507 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.028, intellect=1.601 ± 0.027, spirit=1.154 ± 0.038, mp5=3.342 ± 0.054, crit=0.513 ± 0.016 per rating point (14 rating = 1%, 7.178 per %), spell_haste=-9.501 ± 0.457

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 57.0 healing_power points (5.28 DPS) | yes | Blood Guard's Satin Cover (220899, +0.00 DPS) [vendor]; Cassandra's Grace (13102, -0.24 DPS) [world_drop]; Papal Fez (9431, -0.50 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 27.0 healing_power points (2.50 DPS) | yes | Glowing Eye of Mordresh (10769, -0.11 DPS) [dungeon]; Necklace of Calisea (1714, -0.72 DPS) [dungeon]; Choker of the High Shaman (4112, -0.91 DPS) [quest] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | 45.7 healing_power points (4.24 DPS) | yes | Blood Guard's Satin Pads (220901, +0.00 DPS) [vendor]; Inquisitor's Shawl (19507, -0.92 DPS) [dungeon]; Rotgrip Mantle (17732, -1.57 DPS) [dungeon] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 22.2 healing_power points (2.06 DPS) | yes | Highborne Cloak (14450, -0.30 DPS) [world_drop]; Wingveil Cloak (10802, -0.42 DPS) [dungeon]; Blackforge Cape (6424, -0.45 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 67.2 healing_power points (6.23 DPS) | yes | Stone Guard's Satin Armor (220903, -0.83 DPS) [vendor]; Ghostweave Vest (14141, -1.61 DPS) [crafted]; Grimlok's Tribal Vestments (9415, -2.23 DPS) [dungeon] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 37.0 healing_power points (3.43 DPS) | yes | Shizzle's Nozzle Wiper (11917, -1.33 DPS) [quest]; Forgotten Wraps (9433, -1.65 DPS) [world_drop]; Mindthrust Bracers (1974, -1.76 DPS) [dungeon] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 54.2 healing_power points (5.03 DPS) | yes | Greenleaf Handwraps (19116, -1.36 DPS) [quest]; Earthenweave Gloves (254075, -1.44 DPS) [crafted]; First Sergeant's Satin Gloves (220898, -1.56 DPS) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | sim-verified (220.8 DPS) | yes | Gilded Waistcord (254081, +0.00 DPS) [crafted]; Gilded Cord (254037, -0.06 DPS) [crafted]; Earthenweave Cord (254077, -0.34 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | sim-verified (220.8 DPS) | yes | Senior Designer's Pantaloons (11841, +0.00 DPS) [dungeon]; Stone Guard's Satin Leggings (220902, +0.00 DPS) [vendor]; Dalewind Trousers (13008, -1.14 DPS) [world_drop] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 45.4 healing_power points (4.21 DPS) | yes | First Sergeant's Satin Boots (220900, -0.57 DPS) [vendor]; Earthenweave Boots (254093, -0.62 DPS) [crafted]; Vinerot Sandals (17748, -0.69 DPS) [dungeon] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 23.5 healing_power points (2.17 DPS) | yes | Welken Ring (5011, -0.68 DPS) [world_drop]; Ogremind Ring (1993, -0.81 DPS) [world_drop]; Voodoo Band (1996, -0.81 DPS) [world] |
| finger2 | Coldwater Ring (4550) | Death From Below [quest] | sim-verified (220.8 DPS) | yes | Welken Ring (5011, +0.00 DPS) [world_drop]; Ogremind Ring (1993, -0.05 DPS) [world_drop]; Voodoo Band (1996, -0.05 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.21 DPS) [quest]; Alchemist's Stone (13503, -0.86 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (220.8 DPS) | yes | Evonice's Landin' Pilla (18951, -0.60 DPS) [quest]; Alchemist's Stone (13503, -1.24 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -6.16 DPS) [dungeon]; Hand of Righteousness (7721, -9.60 DPS) [dungeon]; Spellshifter Rod (9527, -10.75 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | sim-verified (220.8 DPS) | yes | Wand of Mana Concentration (273043, +0.00 DPS) [world]; Flash Wand (5248, -0.30 DPS) [quest]; Goblin Igniter (5253, -0.30 DPS) [quest] |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Nethergeld Shoulders; back: Imperial Red Cloak; wrist: Nethergeld Cuffs; hands: Gilded Gloves; waist: Dawnspire Cord; legs: Kilt of the Atal'ai Prophet; feet: Gilded Sandals; finger1: Eye of Adaegus; finger2: Coldwater Ring; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Jaina's Firestarter

No-known-source sample (15 of 507, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60 (undead, 025005031330011530-03505003030000000-00000000000000000)

Set DPS (verified): 312.7. Weights run: 7.8s. Verify run: 30.3s. 1193 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.073, intellect=1.442 ± 0.028, spirit=4.145 ± 0.036, mp5=1.827 ± 0.037, crit=0.364 ± 0.017 per rating point (14 rating = 1%, 5.091 per %), spell_haste=-6.962 ± 0.558

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | 116.8 healing_power points (18.74 DPS) | yes | Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Earthenweave Crown (254137, +0.00 DPS, sim-verified) [crafted] |
| neck | Lady Alizabeth's Pendant (13002) | World drop [world_drop] | 62.2 healing_power points (9.98 DPS) | yes | Amberglow Talisman (10824, +0.00 DPS, sim-verified) [quest]; Choker of the High Shaman (4112, -3.50 DPS) [quest]; Necklace of Calisea (1714, -3.70 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (+10.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Earthenweave Mantle (254119, -0.46 DPS) [crafted]; Mooncloth Shoulders (14139, -10.85 DPS, sim-verified) [crafted] |
| back | Cloak of Blight (6832) (or Indomitable Cloak (14683)) | Nothing But The Truth [quest] | 53.9 healing_power points (8.65 DPS) | yes | Indomitable Cloak (14683, +0.00 DPS) [world_drop]; Wingveil Cloak (10802, -0.53 DPS) [dungeon]; Arcane Cloak (8286, -0.67 DPS) [world_drop] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 139.3 healing_power points (22.36 DPS) | yes | Mooncloth Vest (14138, -0.65 DPS, sim-verified) [crafted]; Legionnaire's Satin Tunic (227124, -3.16 DPS) [pvp]; Warlord's Satin Tunic (231632, -3.31 DPS) [vendor] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | 64.0 healing_power points (10.27 DPS) | yes | General's Satin Bracers (17619, +0.00 DPS) [pvp]; Earthenweave Cuffs (254131, +0.00 DPS, sim-verified) [crafted]; Nethergeld Cuffs (254061, -2.11 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | 99.9 healing_power points (16.03 DPS) | yes | Swarmtender's Gloves (275607, +0.00 DPS, sim-verified) [crafted]; Magister's Gloves (16684, -3.48 DPS) [dungeon]; Gilded Gloves (254095, -3.74 DPS) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (+7.6 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Satin Cinch (17621, +0.00 DPS) [pvp]; Mana-infused Cord (279267, -0.41 DPS) [crafted]; Wisdom of the Timbermaw (19047, -7.62 DPS, sim-verified) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 141.0 healing_power points (22.62 DPS) | yes | Legionnaire's Satin Legguards (227123, -3.42 DPS) [pvp]; Mooncloth Leggings (14137, -3.68 DPS) [crafted]; Dalewind Trousers (13008, -3.91 DPS) [world_drop] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (+9.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Elunarian Boots (14458, -2.69 DPS) [world_drop]; Furen's Boots (13100, -3.12 DPS) [world_drop]; Mooncloth Boots (15802, -9.08 DPS, sim-verified) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 60.0 healing_power points (9.63 DPS) | yes | Coldwater Ring (4550, -2.95 DPS) [quest]; Welken Ring (5011, -3.82 DPS) [world_drop]; The Queen's Jewel (13094, -3.85 DPS) [world_drop] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 58.4 healing_power points (9.37 DPS) | yes | Coldwater Ring (4550, +0.00 DPS, sim-verified) [quest]; Welken Ring (5011, -3.56 DPS) [world_drop]; The Queen's Jewel (13094, -3.59 DPS) [world_drop] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop]; Alchemist's Stone (13503, +0.00 DPS) [crafted]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+9.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Evonice's Landin' Pilla (18951, -1.33 DPS) [quest]; Alchemist's Stone (13503, -5.32 DPS) [crafted]; Uther's Strength (11302, -9.23 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -15.02 DPS) [dungeon]; Hand of Edward the Odd (2243, -15.04 DPS, sim-verified) [world_drop]; Spire of Hakkar (10844, -16.17 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 62.6 healing_power points (10.05 DPS) | yes | Chillnail Splinter (10704, -2.22 DPS, sim-verified) [quest]; Jaina's Firestarter (13064, -6.66 DPS) [world_drop]; Captain Rackmore's Tiller (16789, -7.09 DPS) [quest] |

**New at 60:** head: Devout Crown; neck: Lady Alizabeth's Pendant; shoulder: Devout Mantle; back: Cloak of Blight; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Shard of the Splithooves; trinket2: Ankh of Life; ranged: Torch of Light

No-known-source sample (15 of 1193, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60, raid preset (undead, 025005031330011530-03505003030000000-00000000000000000)

Set DPS (verified): 586.8. Weights run: 5.2s. Verify run: 19.1s. 1193 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.109, intellect=1.218 ± 0.030, spirit=1.118 ± 0.026, mp5=2.481 ± 0.033, crit=0.604 ± 0.022 per rating point (14 rating = 1%, 8.453 per %), spell_haste=-4.887 ± 0.556

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (586.8 DPS) | yes | Mooncloth Circlet (14140, +0.00 DPS) [crafted]; Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 22.7 healing_power points (5.50 DPS) | yes | Horizon Choker (13085, -0.29 DPS) [world_drop]; Lady Alizabeth's Pendant (13002, -1.44 DPS) [world_drop]; Necklace of Calisea (1714, -1.55 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (586.8 DPS) | yes | Mooncloth Shoulders (14139, +0.00 DPS) [crafted]; Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp] |
| back | Wingveil Cloak (10802) | Sunken Temple: Spawn of Hakkar [dungeon] | sim-verified (586.8 DPS) | yes | Imperial Red Cloak (8248, +0.00 DPS) [world_drop]; Glowing Thresher Cape (6901, -0.33 DPS) [dungeon]; Highborne Cloak (14450, -0.45 DPS) [world_drop] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 90.9 healing_power points (21.98 DPS) | yes | Warlord's Satin Tunic (231632, -1.38 DPS) [vendor]; Legionnaire's Satin Tunic (227124, -4.10 DPS) [pvp]; Mooncloth Vest (14138, -4.66 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (586.8 DPS) | yes | General's Satin Bracers (17619, +0.00 DPS) [pvp]; Nethergeld Cuffs (254061, +0.00 DPS) [crafted]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (586.8 DPS) | yes | Gilded Gloves (254095, +0.00 DPS) [crafted]; Blood Guard's Satin Handwraps (227126, -0.02 DPS) [pvp]; General's Satin Handwraps (231633, -0.33 DPS) [vendor] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 58.1 healing_power points (14.05 DPS) | yes | Devout Belt (16696, -3.07 DPS) [dungeon]; General's Satin Cinch (17621, -3.22 DPS) [pvp]; Gilded Waistcord (254081, -4.39 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 68.0 healing_power points (16.44 DPS) | yes | Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Earthenweave Leggings (254125, -3.17 DPS) [crafted] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 55.6 healing_power points (13.44 DPS) | yes | General's Satin Walkers (231630, -1.13 DPS) [vendor]; Blood Guard's Satin Walkers (227127, -2.74 DPS) [pvp]; Swarmtender's Footpads (275609, -3.05 DPS) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 24.5 healing_power points (5.92 DPS) | yes | Chemist's Ring (15702, -1.89 DPS) [quest]; Maiden's Circle (13001, -2.24 DPS) [world_drop]; Glowing Crystal Ring (18402, -2.39 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 20.7 healing_power points (5.01 DPS) | yes | Chemist's Ring (15702, -0.98 DPS) [quest]; Maiden's Circle (13001, -1.33 DPS) [world_drop]; Glowing Crystal Ring (18402, -1.48 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (586.8 DPS) | yes | Ankh of Life (1713, -1.11 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.65 DPS) [quest]; Alchemist's Stone (13503, -3.27 DPS) [crafted] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (586.8 DPS) | yes | Ankh of Life (1713, -0.35 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.90 DPS) [quest]; Alchemist's Stone (13503, -2.52 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -14.76 DPS) [dungeon]; Hand of Edward the Odd (2243, -18.95 DPS) [world_drop]; Hand of Righteousness (7721, -22.35 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 29.3 healing_power points (7.08 DPS) | yes | Jaina's Firestarter (13064, -4.51 DPS) [world_drop]; Wand of Mana Concentration (273043, -4.71 DPS) [world]; Captain Rackmore's Tiller (16789, -4.82 DPS) [quest] |

**New at 60:** head: Devout Crown; neck: Glowing Eye of Mordresh; shoulder: Devout Mantle; back: Wingveil Cloak; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Uther's Strength; trinket2: Shard of the Splithooves; ranged: Torch of Light

No-known-source sample (15 of 1193, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

