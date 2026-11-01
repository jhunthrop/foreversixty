# Leveling BiS: Discipline

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 025004000000000000-00000000000000000-00000000000000000)

Set DPS (verified): 36.6. Weights run: 10.8s. Verify run: 29.2s. 157 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.010, intellect=2.707 ± 0.012, spirit=1.309 ± 0.007, mp5=2.960 ± 0.024, crit=0.139 ± 0.006 per rating point (14 rating = 1%, 1.952 per %), spell_haste=0.674 ± 0.086

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Goggles (4373) | Engineering [crafted] | 21.4 healing_power points (1.04 DPS) | yes | Pristine Circlet (253949, +0.00 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -0.78 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 5.2 healing_power points (0.25 DPS) | yes | Scholarly Pendant (277203, -0.20 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 24.4 healing_power points (1.18 DPS) | yes | Slime-encrusted Pads (6461, -0.75 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -0.81 DPS, sim-verified) [crafted] |
| back | Sanguine Cape (14376) | World drop [world_drop] | 10.8 healing_power points (0.53 DPS) | yes | Seer's Cape (6378, -0.14 DPS) [dungeon]; Forest Cloak (4710, -0.26 DPS) [world_drop]; Pearl-clasped Cloak (5542, -0.29 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 30.5 healing_power points (1.48 DPS) | yes | Corsair's Overshirt (5202, -0.44 DPS) [dungeon]; Seer's Robe (2981, -0.50 DPS) [world_drop]; Robe of the Moccasin (6465, -0.73 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 22.4 healing_power points (1.09 DPS) | yes | Repurposed Hair Band (281256, -0.70 DPS) [quest]; Mystic's Bracelets (14366, -0.83 DPS) [world_drop]; Bright Bracers (3647, -1.07 DPS, sim-verified) [world_drop] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 19.1 healing_power points (0.93 DPS) | yes | Tomb Robber's Gloves (280096, -0.14 DPS) [quest]; Bright Gloves (3066, -0.15 DPS) [world_drop]; Magefist Gloves (12977, -0.39 DPS, sim-verified) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 21.8 healing_power points (1.06 DPS) | yes | Novice Ardent's Sash (253887, -0.18 DPS) [crafted]; Keller's Girdle (2911, -0.33 DPS, sim-verified) [world_drop]; Wise Man's Belt (4786, -0.40 DPS) [vendor] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 39.5 healing_power points (1.92 DPS) | yes | Filigreed Silky Leggings (253939, -0.87 DPS) [crafted]; Filigreed Shadow Leggings (253943, -0.87 DPS) [crafted]; Darkweave Breeches (12987, -1.15 DPS, sim-verified) [world_drop] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 17.1 healing_power points (0.83 DPS) | yes | Kimbra Boots (6191, -0.28 DPS, sim-verified) [quest]; Spidersilk Boots (4320, -0.31 DPS) [crafted]; Sanguine Sandals (14374, -0.31 DPS) [world_drop] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 16.2 healing_power points (0.79 DPS) | yes | Band of Purification (12996, -0.41 DPS) [world_drop]; Deep Fathom Ring (6463, -0.47 DPS) [dungeon]; Minor Channeling Ring (1449, -0.53 DPS) [quest] |
| finger2 | Volcanic Rock Ring (12053) | World drop [world_drop] | 8.1 healing_power points (0.39 DPS) | yes | Deep Fathom Ring (6463, -0.08 DPS) [dungeon]; Minor Channeling Ring (1449, -0.13 DPS) [quest]; Band of Purification (12996, -0.36 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 27.1 healing_power points (1.31 DPS) | yes | Lesser Staff of the Spire (1300, -0.53 DPS) [world]; Staff of Westfall (2042, -0.54 DPS, sim-verified) [quest]; Witching Stave (1484, -0.66 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Flaring Baton (5326) (or Moonstone Wand (15204)) | The Escape [quest] | 5.4 healing_power points (0.26 DPS) | yes | Moonstone Wand (15204, +0.00 DPS) [quest]; Sable Wand (7607, -0.07 DPS) [quest]; Elven Wand (5604, -0.13 DPS) [quest] |

**New at 20:** head: Shadow Goggles; neck: Tarnished Locket; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Volcanic Rock Ring; main_hand: Twisted Chanter's Staff; ranged: Flaring Baton

No-known-source sample (15 of 157, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes; 16607 Acolyte's Sacrificial Robes; 18851 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (human, 025005031320000000-00000000000000000-00000000000000000)

Set DPS (verified): 67.7. Weights run: 11.1s. Verify run: 33.1s. 288 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.016, intellect=2.127 ± 0.014, spirit=1.715 ± 0.011, mp5=3.813 ± 0.013, crit=0.203 ± 0.008 per rating point (14 rating = 1%, 2.836 per %), spell_haste=not significant (-0.105 ± 0.086)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 43.3 healing_power points (2.75 DPS) | yes | Resilient Cap (14401, -0.83 DPS) [world_drop]; Nightsky Cowl (4039, -1.11 DPS, sim-verified) [world_drop]; Shadow Hood (4323, -1.26 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.4 healing_power points (0.98 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.17 DPS) [world]; Crystal Starfire Medallion (5003, -0.21 DPS, sim-verified) [world_drop]; Pendant of Myzrael (4614, -0.32 DPS) [dungeon] |
| shoulder | Mantle of Honor (3560) | Bride of the Embalmer [quest] | 26.9 healing_power points (1.71 DPS) | yes | Nightsky Mantle (4718, -0.15 DPS, sim-verified) [world_drop]; Death Speaker Mantle (6685, -0.22 DPS) [dungeon]; Faerie Mantle (5820, -0.24 DPS) [quest] |
| back | Prelacy Cape (7004) | Researching the Corruption [quest] | 21.3 healing_power points (1.35 DPS) | yes | Glowing Thresher Cape (6901, -0.12 DPS, sim-verified) [dungeon]; Repairman's Cape (9605, -0.15 DPS) [quest]; Sergeant Major's Cape (16315, -0.38 DPS) [pvp] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 43.4 healing_power points (2.75 DPS) | yes | Beguiler Robes (7728, -0.26 DPS) [dungeon]; Pristine Gown (253961, -0.38 DPS, sim-verified) [crafted]; Silver-thread Robe (4035, -0.83 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 22.1 healing_power points (1.40 DPS) | yes | Glowing Magical Bracelets (13106, -0.32 DPS) [world_drop]; Nightsky Wristbands (6407, -0.37 DPS, sim-verified) [world_drop]; Spidertank Oilrag (9448, -0.68 DPS) [dungeon] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 31.2 healing_power points (1.98 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.17 DPS, sim-verified) [dungeon]; Truefaith Gloves (7049, -0.62 DPS) [crafted]; Zodiac Gloves (7106, -0.68 DPS) [quest] |
| waist | Resilient Cord (14406) | World drop [world_drop] | 19.6 healing_power points (1.25 DPS) | yes | Pristine Sash (253925, +0.00 DPS, sim-verified) [crafted]; Dreamer's Belt (4829, -0.08 DPS) [vendor]; Novice Ardent's Sash (253887, -0.13 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 45.5 healing_power points (2.89 DPS) | yes | Filigreed Pristine Leggings (253937, -0.62 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -0.74 DPS) [crafted]; Darkweave Breeches (12987, -1.29 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 41.7 healing_power points (2.65 DPS) | yes | Frothing Slippers (254003, -1.27 DPS) [crafted]; Black Slippers (254007, -1.27 DPS) [crafted]; Acidic Walkers (9454, -1.47 DPS, sim-verified) [dungeon] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 18.0 healing_power points (1.14 DPS) | yes | The Queen's Jewel (13094, -0.00 DPS) [world_drop]; Defias Mage Ring (1077, -0.28 DPS) [world]; Seal of Wrynn (2933, -0.28 DPS) [quest] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 18.0 healing_power points (1.14 DPS) | yes | The Queen's Jewel (13094, -0.00 DPS) [world_drop]; Defias Mage Ring (1077, -0.28 DPS) [world]; Seal of Wrynn (2933, -0.28 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-verified (67.7 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -0.04 DPS) [quest]; Death Speaker Scepter (2816, -1.43 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 16.1 healing_power points (1.02 DPS) | yes | Lesser Mystic Wand (11289, -0.48 DPS) [crafted]; Starfaller (13063, -0.48 DPS) [world_drop]; Summoner's Wand (5245, -0.57 DPS, sim-verified) [world] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Mantle of Honor; back: Prelacy Cape; chest: Death Speaker Robes; hands: Gloves of Old; waist: Resilient Cord; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Spine of the Devourer; ranged: Wand of Mana Concentration

No-known-source sample (15 of 288, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 40 (human, 025005031330011520-00000000000000000-00000000000000000)

Set DPS (verified): 130.9. Weights run: 12.8s. Verify run: 77.7s. 390 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.020, intellect=2.098 ± 0.021, spirit=0.809 ± 0.036, mp5=2.286 ± 0.038, crit=0.418 ± 0.013 per rating point (14 rating = 1%, 5.857 per %), spell_haste=-10.189 ± 0.391

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 57.0 healing_power points (4.38 DPS) | yes | Whitemane's Chapeau (7720, -0.61 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -1.47 DPS) [world_drop]; Thinking Cap (2624, -1.64 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 29.8 healing_power points (2.29 DPS) | yes | Necklace of Calisea (1714, -0.15 DPS, sim-verified) [dungeon]; Archmagister's Faceted Pendant (273049, -1.32 DPS) [world]; Kaleidoscope Chain (13084, -1.40 DPS) [world_drop] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 42.3 healing_power points (3.25 DPS) | yes | Mistscape Mantle (4734, -1.17 DPS) [dungeon]; Silksand Shoulder Pads (14423, -1.29 DPS) [world_drop]; Windchaser Amice (14432, -1.48 DPS, sim-verified) [world_drop] |
| back | Blackforge Cape (6424) | Uldaman: Ancient Treasure [dungeon] | 20.0 healing_power points (1.54 DPS) | yes | Sergeant Major's Cape (16336, +0.00 DPS, sim-verified) [pvp]; Prelacy Cape (7004, -0.32 DPS) [quest]; Windchaser Cloak (14430, -0.32 DPS) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 64.5 healing_power points (4.96 DPS) | yes | Pristine Gown (253961, -2.04 DPS) [crafted]; Red Mageweave Vest (10007, -2.06 DPS) [crafted]; Death Speaker Robes (6682, -5.43 DPS, sim-verified) [dungeon] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 18.9 healing_power points (1.45 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Mistscape Bracers (4045, -0.16 DPS) [dungeon]; Enchanted Stonecloth Bracers (4979, -0.16 DPS) [quest] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (130.9 DPS) | yes | Earthen Silk Gloves (254017, -0.90 DPS) [crafted]; Truefaith Gloves (7049, -1.15 DPS) [crafted]; Gilded Handwraps (254021, -1.70 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 43.6 healing_power points (3.36 DPS) | yes | Sutarn's Ring (13105, -1.37 DPS) [world_drop]; Razzeric's Customized Seatbelt (6726, -1.42 DPS) [quest]; Deathmage Sash (10771, -2.03 DPS, sim-verified) [dungeon] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-verified (130.9 DPS) | yes | Stoneweaver Leggings (9407, +0.00 DPS) [dungeon]; Filigreed Pristine Leggings (253937, +0.00 DPS) [crafted]; Pristine Leggings (253987, -2.05 DPS, sim-verified) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 37.9 healing_power points (2.92 DPS) | yes | Furen's Boots (13100, +0.00 DPS, sim-verified) [world_drop]; Kodo Rustler Boots (15697, -1.32 DPS) [quest]; Acidic Walkers (9454, -1.38 DPS) [dungeon] |
| finger1 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 17.1 healing_power points (1.32 DPS) | yes | Mindbender Loop (5009, -0.06 DPS) [world_drop]; Welken Ring (5011, -0.07 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.35 DPS) [dungeon] |
| finger2 | Ogremind Ring (1993) | World drop [world_drop] | 17.1 healing_power points (1.32 DPS) | yes | Mindbender Loop (5009, -0.06 DPS) [world_drop]; Welken Ring (5011, -0.07 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.35 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (130.9 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Death Speaker Scepter (2816, -2.50 DPS) [dungeon]; Hypnotic Blade (7714, -2.87 DPS) [dungeon] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 33.0 healing_power points (2.54 DPS) | yes | Aurora Sphere (7610, -0.84 DPS, sim-verified) [dungeon]; Orb of Souls (249395, -1.16 DPS) [crafted]; Beacon of Hope (9393, -1.21 DPS) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 15.0 healing_power points (1.16 DPS) | yes | Wand of Mana Concentration (273043, +0.00 DPS, sim-verified) [world]; Flash Wand (5248, -0.32 DPS) [quest]; Goblin Igniter (5253, -0.32 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Blackforge Cape; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Stormcloth Gloves; waist: Gilded Cord; legs: Stormcloth Pants; finger1: Voodoo Band; finger2: Ogremind Ring; trinket1: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Orb of Lorica; ranged: Jaina's Firestarter

No-known-source sample (15 of 390, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 50 (human, 025005031330011520-03502000000000000-00000000000000000)

Set DPS (verified): 209.7. Weights run: 12.7s. Verify run: 43.4s. 509 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.031, intellect=2.648 ± 0.032, spirit=1.981 ± 0.049, mp5=4.589 ± 0.072, crit=0.639 ± 0.021 per rating point (14 rating = 1%, 8.942 per %), spell_haste=-11.820 ± 0.573

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 89.7 healing_power points (6.55 DPS) | yes | Knight-Lieutenant's Satin Cover (220896, -0.24 DPS) [vendor]; Papal Fez (9431, -1.00 DPS, sim-verified) [dungeon]; Imperial Red Circlet (8254, -1.28 DPS) [world_drop] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 45.0 healing_power points (3.29 DPS) | yes | Glowing Eye of Mordresh (10769, -0.35 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -0.92 DPS) [dungeon]; Mindburst Medallion (11196, -1.40 DPS) [quest] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | 59.1 healing_power points (4.31 DPS) | yes | Knight-Lieutenant's Satin Pads (220894, +0.00 DPS) [vendor]; Rotgrip Mantle (17732, -0.83 DPS) [dungeon]; Inquisitor's Shawl (19507, -1.15 DPS, sim-verified) [dungeon] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 37.1 healing_power points (2.71 DPS) | yes | Highborne Cloak (14450, -0.38 DPS, sim-verified) [world_drop]; Blackforge Cape (6424, -0.58 DPS) [dungeon]; Spritecaster Cape (11623, -0.68 DPS) [dungeon] |
| chest | Ghostweave Vest (14141) | Tailoring [crafted] | 76.3 healing_power points (5.57 DPS) | yes | Stormcloth Vest (10020, -0.18 DPS) [crafted]; Knight's Satin Armor (220892, -0.77 DPS, sim-verified) [vendor]; Earth Warder's Vest (21311, -1.05 DPS) [quest] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 48.4 healing_power points (3.54 DPS) | yes | Forgotten Wraps (9433, -1.22 DPS) [world_drop]; Bloodband Bracers (11469, -1.36 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -2.82 DPS, sim-verified) [quest] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 71.3 healing_power points (5.21 DPS) | yes | Earthenweave Gloves (254075, -1.72 DPS) [crafted]; Highborne Gloves (14451, -1.83 DPS) [world_drop]; Quintis' Research Gloves (11888, -4.02 DPS, sim-verified) [quest] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | sim-verified (209.6 DPS) | yes | Gilded Waistcord (254081, -0.16 DPS) [crafted]; Earthenweave Cord (254077, -0.53 DPS) [crafted]; Dawnspire Cord (12466, -4.88 DPS, sim-verified) [dungeon] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 89.3 healing_power points (6.52 DPS) | yes | Senior Designer's Pantaloons (11841, -0.25 DPS, sim-verified) [dungeon]; Knight's Satin Leggings (220893, -1.08 DPS) [vendor]; Dalewind Trousers (13008, -1.17 DPS) [world_drop] |
| feet | Vinerot Sandals (17748) | Maraudon: Razorlash [dungeon] | 55.6 healing_power points (4.06 DPS) | yes | Gilded Sandals (254107, +0.00 DPS, sim-verified) [crafted]; Sergeant Major's Satin Boots (220895, -0.28 DPS) [vendor]; Furen's Boots (13100, -0.40 DPS) [world_drop] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 39.7 healing_power points (2.90 DPS) | yes | Coldwater Ring (4550, -1.01 DPS) [quest]; Choking Band (11868, -1.02 DPS) [quest]; Voodoo Band (1996, -1.11 DPS) [world] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | 27.1 healing_power points (1.98 DPS) | yes | Choking Band (11868, -0.10 DPS) [quest]; Voodoo Band (1996, -0.19 DPS) [world]; Coldwater Ring (4550, -0.89 DPS, sim-verified) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -0.16 DPS) [quest]; Alchemist's Stone (13503, -0.74 DPS) [crafted] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, -0.29 DPS) [quest]; Ankh of Life (1713, -0.60 DPS, sim-verified) [world_drop]; Alchemist's Stone (13503, -0.87 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -6.60 DPS) [dungeon]; Spellshifter Rod (9527, -9.47 DPS) [quest]; Hammer of the Northern Wind (810, -25.93 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cairnstone Sliver (9654) | The Morrow Stone [quest] | 23.1 healing_power points (1.69 DPS) | yes | Wand of Mana Concentration (273043, -0.25 DPS) [world]; Goblin Igniter (5253, -0.48 DPS) [quest]; Jaina's Firestarter (13064, -0.96 DPS, sim-verified) [world_drop] |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Nethergeld Shoulders; back: Imperial Red Cloak; chest: Ghostweave Vest; wrist: Nethergeld Cuffs; hands: Gilded Gloves; legs: Kilt of the Atal'ai Prophet; feet: Vinerot Sandals; finger1: Eye of Adaegus; finger2: Welken Ring; trinket1: Uther's Strength; trinket2: Evonice's Landin' Pilla; main_hand: Glowing Brightwood Staff; ranged: Cairnstone Sliver

No-known-source sample (15 of 509, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60 (human, 025005031330011520-03505003030100000-00000000000000000)

Set DPS (verified): 309.3. Weights run: 12.3s. Verify run: 46.7s. 1194 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.133, intellect=2.713 ± 0.071, spirit=3.888 ± 0.434, mp5=3.203 ± 0.192, crit=0.850 ± 0.041 per rating point (14 rating = 1%, 11.897 per %), spell_haste=-12.342 ± 1.234

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | 143.4 healing_power points (11.76 DPS) | yes | Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Lieutenant Commander's Satin Hood (227121, -0.01 DPS) [pvp]; Gemburst Circlet (10751, -10.62 DPS, sim-verified) [quest] |
| neck | Lady Alizabeth's Pendant (13002) | World drop [world_drop] | 58.3 healing_power points (4.78 DPS) | yes | Necklace of Calisea (1714, -0.99 DPS) [dungeon]; Amberglow Talisman (10824, -1.59 DPS) [quest]; Horizon Choker (13085, -2.69 DPS, sim-verified) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | 103.0 healing_power points (8.44 DPS) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, -1.53 DPS) [pvp]; Mooncloth Shoulders (14139, -8.02 DPS, sim-verified) [crafted] |
| back | Ethereal Mist Cape (11873) | Ancient Spirit [quest] | 54.8 healing_power points (4.49 DPS) | yes | Indomitable Cloak (14683, -0.35 DPS) [world_drop]; Sergeant Major's Cape (16337, -0.42 DPS, sim-verified) [pvp]; Wingveil Cloak (10802, -0.58 DPS) [dungeon] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | 154.1 healing_power points (12.64 DPS) | yes | Field Marshal's Satin Tunic (231624, -0.18 DPS) [vendor]; Devout Robe (16690, -0.44 DPS, sim-verified) [dungeon]; Truefaith Vestments (14154, -1.55 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | 71.6 healing_power points (5.87 DPS) | yes | Marshal's Satin Bracers (17606, +0.00 DPS) [pvp]; Magister's Bindings (16683, -1.07 DPS) [dungeon]; Earthenweave Cuffs (254131, -6.33 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | 108.2 healing_power points (8.87 DPS) | yes | Magister's Gloves (16684, -1.30 DPS) [dungeon]; Gilded Gloves (254095, -1.72 DPS) [crafted]; Swarmtender's Gloves (275607, -9.54 DPS, sim-verified) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 104.0 healing_power points (8.53 DPS) | yes | Marshal's Satin Sash (17609, +0.00 DPS) [pvp]; Elunarian Belt (14465, -0.89 DPS) [world_drop]; Devout Belt (16696, -2.02 DPS, sim-verified) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 154.1 healing_power points (12.64 DPS) | yes | Knight-Captain's Satin Legguards (227125, -1.56 DPS) [pvp]; Marshal's Satin Legguards (231626, -1.90 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -12.36 DPS, sim-verified) [dungeon] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 112.5 healing_power points (9.23 DPS) | yes | Elunarian Boots (14458, -2.47 DPS) [world_drop]; Furen's Boots (13100, -2.57 DPS) [world_drop]; Devout Sandals (16691, -3.33 DPS, sim-verified) [dungeon] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 69.9 healing_power points (5.73 DPS) | yes | Choking Band (11868, -1.59 DPS) [quest]; Coldwater Ring (4550, -2.19 DPS) [quest]; Maiden's Circle (13001, -2.26 DPS) [world_drop] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 62.9 healing_power points (5.16 DPS) | yes | Choking Band (11868, -1.52 DPS, sim-verified) [quest]; Coldwater Ring (4550, -1.62 DPS) [quest]; Maiden's Circle (13001, -1.69 DPS) [world_drop] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (309.3 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Uther's Strength (11302, -3.88 DPS, sim-verified) [world_drop] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (309.3 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.36 DPS, sim-verified) [quest]; Uther's Strength (11302, -1.07 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (309.3 DPS) | yes | Staff of Hale Magefire (13000, -9.05 DPS) [world_drop]; Charstone Dirk (17710, -9.14 DPS) [dungeon]; Hand of Edward the Odd (2243, -14.84 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 59.8 healing_power points (4.90 DPS) | yes | Jaina's Firestarter (13064, -2.61 DPS) [world_drop]; Goblin Igniter (5253, -3.05 DPS) [quest]; Cairnstone Sliver (9654, -5.84 DPS, sim-verified) [quest] |

**New at 60:** head: Devout Crown; neck: Lady Alizabeth's Pendant; shoulder: Devout Mantle; back: Ethereal Mist Cape; chest: Mooncloth Vest; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Shard of the Splithooves; trinket2: Thunderbrew's Boot Flask; ranged: Torch of Light

No-known-source sample (15 of 1194, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60, raid preset (human, 025005031330011520-03505003030100000-00000000000000000)

Set DPS (verified): 569.8. Weights run: 8.3s. Verify run: 45.7s. 1194 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.128, intellect=1.212 ± 0.028, spirit=1.313 ± 0.026, mp5=2.328 ± 0.030, crit=0.593 ± 0.023 per rating point (14 rating = 1%, 8.308 per %), spell_haste=-5.523 ± 0.525

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (569.9 DPS) | yes | Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -8.55 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 22.7 healing_power points (5.28 DPS) | yes | Horizon Choker (13085, -0.11 DPS) [world_drop]; Lady Alizabeth's Pendant (13002, -0.70 DPS) [world_drop]; Necklace of Calisea (1714, -1.17 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (569.9 DPS) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp]; Mooncloth Shoulders (14139, -7.73 DPS, sim-verified) [crafted] |
| back | Sergeant Major's Cape (16337) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 20.2 healing_power points (4.70 DPS) | yes | Ethereal Mist Cape (11873, -0.19 DPS) [quest]; Prelacy Cape (7004, -0.31 DPS) [quest]; Wingveil Cloak (10802, -1.95 DPS, sim-verified) [dungeon] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 94.0 healing_power points (21.85 DPS) | yes | Field Marshal's Satin Tunic (231624, -1.69 DPS) [vendor]; Knight-Captain's Satin Tunic (227122, -4.06 DPS) [pvp]; Mooncloth Vest (14138, -5.87 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (569.9 DPS) | yes | Marshal's Satin Bracers (17606, +0.00 DPS) [pvp]; Nethergeld Cuffs (254061, +0.00 DPS) [crafted]; Earthenweave Cuffs (254131, -8.09 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (569.9 DPS) | yes | Knight-Lieutenant's Satin Handwraps (227128, -0.79 DPS) [pvp]; Marshal's Satin Handwraps (231623, -1.09 DPS) [vendor]; Gilded Gloves (254095, -3.64 DPS, sim-verified) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 59.8 healing_power points (13.90 DPS) | yes | Marshal's Satin Sash (17609, -2.84 DPS) [pvp]; Devout Belt (16696, -3.43 DPS, sim-verified) [dungeon]; Earthenweave Cord (254077, -4.50 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 72.4 healing_power points (16.83 DPS) | yes | Knight-Captain's Satin Legguards (227125, +0.00 DPS) [pvp]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Mooncloth Leggings (14137, -12.71 DPS, sim-verified) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (569.9 DPS) | yes | Marshal's Satin Walkers (231627, +0.00 DPS) [vendor]; Knight-Lieutenant's Satin Walkers (227129, -0.29 DPS) [pvp]; Mooncloth Boots (15802, -5.64 DPS, sim-verified) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 26.6 healing_power points (6.18 DPS) | yes | Choking Band (11868, -2.21 DPS) [quest]; Maiden's Circle (13001, -2.37 DPS) [world_drop]; Coldwater Ring (4550, -2.58 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 23.0 healing_power points (5.35 DPS) | yes | Maiden's Circle (13001, -1.55 DPS) [world_drop]; Coldwater Ring (4550, -1.76 DPS) [quest]; Choking Band (11868, -1.89 DPS, sim-verified) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (569.9 DPS) | yes | Shard of the Splithooves (10659, -0.94 DPS) [quest]; Evonice's Landin' Pilla (18951, -1.13 DPS) [quest]; Thunderbrew's Boot Flask (744, -1.74 DPS) [quest] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (569.9 DPS) | yes | Shard of the Splithooves (10659, -0.42 DPS) [quest]; Evonice's Landin' Pilla (18951, -1.06 DPS, sim-verified) [quest]; Thunderbrew's Boot Flask (744, -1.22 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (569.9 DPS) | yes | Charstone Dirk (17710, -14.64 DPS) [dungeon]; Hand of Righteousness (7721, -21.49 DPS) [dungeon]; Hand of Edward the Odd (2243, -23.35 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 31.4 healing_power points (7.31 DPS) | yes | Jaina's Firestarter (13064, -4.70 DPS) [world_drop]; Captain Rackmore's Tiller (16789, -5.00 DPS) [quest]; Cairnstone Sliver (9654, -6.35 DPS, sim-verified) [quest] |

**New at 60:** head: Devout Crown; neck: Glowing Eye of Mordresh; shoulder: Devout Mantle; back: Sergeant Major's Cape; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket2: Ankh of Life; ranged: Torch of Light

No-known-source sample (15 of 1194, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

## Horde

### Band 20 (undead, 025004000000000000-00000000000000000-00000000000000000)

Set DPS (verified): 36.7. Weights run: 10.8s. Verify run: 31.0s. 150 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.010, intellect=2.707 ± 0.012, spirit=1.309 ± 0.007, mp5=2.960 ± 0.024, crit=0.139 ± 0.006 per rating point (14 rating = 1%, 1.952 per %), spell_haste=0.674 ± 0.086

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | sim-verified (36.7 DPS) | yes | Shadow Goggles (4373, -0.43 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -0.62 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 3.9 healing_power points (0.19 DPS) | yes | Scholarly Pendant (277203, -0.05 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 24.4 healing_power points (1.18 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.69 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.75 DPS) [dungeon] |
| back | Sanguine Cape (14376) | World drop [world_drop] | 10.8 healing_power points (0.53 DPS) | yes | Seer's Cape (6378, -0.14 DPS) [dungeon]; Pearl-clasped Cloak (5542, -0.19 DPS, sim-verified) [crafted]; Forest Cloak (4710, -0.26 DPS) [world_drop] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 30.5 healing_power points (1.48 DPS) | yes | Corsair's Overshirt (5202, -0.44 DPS) [dungeon]; Seer's Robe (2981, -0.50 DPS) [world_drop]; Robe of the Moccasin (6465, -0.61 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 22.4 healing_power points (1.09 DPS) | yes | Featherbead Bracers (15452, -0.53 DPS, sim-verified) [quest]; Bright Bracers (3647, -0.56 DPS) [world_drop]; Crystalline Cuffs (14148, -0.63 DPS) [dungeon] |
| hands | Blight Gloves (279877) | The New Plague [quest] | 25.5 healing_power points (1.24 DPS) | yes | Pristine Gloves (253913, -0.31 DPS) [crafted]; Magefist Gloves (12977, -0.33 DPS) [world_drop]; Bright Gloves (3066, -0.46 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 21.8 healing_power points (1.06 DPS) | yes | Novice Ardent's Sash (253887, -0.18 DPS) [crafted]; Keller's Girdle (2911, -0.22 DPS, sim-verified) [world_drop]; Wise Man's Belt (4786, -0.40 DPS) [vendor] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 39.5 healing_power points (1.92 DPS) | yes | Filigreed Silky Leggings (253939, -0.87 DPS) [crafted]; Filigreed Shadow Leggings (253943, -0.87 DPS) [crafted]; Darkweave Breeches (12987, -0.97 DPS, sim-verified) [world_drop] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 17.1 healing_power points (0.83 DPS) | yes | Spidersilk Boots (4320, -0.31 DPS) [crafted]; Bluegill Sandals (1560, -0.31 DPS) [world]; Sanguine Sandals (14374, -0.41 DPS, sim-verified) [world_drop] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 16.2 healing_power points (0.79 DPS) | yes | Volcanic Rock Ring (12053, -0.39 DPS) [world_drop]; Band of Purification (12996, -0.41 DPS) [world_drop]; Deep Fathom Ring (6463, -0.47 DPS) [dungeon] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 13.5 healing_power points (0.66 DPS) | yes | Volcanic Rock Ring (12053, -0.22 DPS, sim-verified) [world_drop]; Band of Purification (12996, -0.28 DPS) [world_drop]; Deep Fathom Ring (6463, -0.34 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 27.1 healing_power points (1.31 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Staff of Orgrimmar (15444, -0.18 DPS) [quest]; Lesser Staff of the Spire (1300, -0.53 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Decay (5252) (or Flaring Baton (5326)) | Beren's Peril [quest] | 5.4 healing_power points (0.26 DPS) | yes | Flaring Baton (5326, +0.00 DPS) [quest]; Wisesight Wand (286750, -0.13 DPS) [world] |

**New at 20:** head: Pristine Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Blight Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Gnarled Necromancer's Staff; ranged: Wand of Decay

No-known-source sample (15 of 150, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (undead, 025005031320000000-00000000000000000-00000000000000000)

Set DPS (verified): 67.0. Weights run: 11.1s. Verify run: 33.3s. 280 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.016, intellect=2.127 ± 0.014, spirit=1.715 ± 0.011, mp5=3.813 ± 0.013, crit=0.203 ± 0.008 per rating point (14 rating = 1%, 2.836 per %), spell_haste=not significant (-0.105 ± 0.086)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 43.3 healing_power points (2.75 DPS) | yes | Resilient Cap (14401, -0.83 DPS) [world_drop]; Nightsky Cowl (4039, -0.88 DPS, sim-verified) [world_drop]; Shadow Hood (4323, -1.26 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.4 healing_power points (0.98 DPS) | yes | Crystal Starfire Medallion (5003, -0.15 DPS, sim-verified) [world_drop]; Archmagister's Faceted Pendant (273049, -0.17 DPS) [world]; Pendant of Myzrael (4614, -0.32 DPS) [dungeon] |
| shoulder | Nightsky Mantle (4718) | World drop [world_drop] | 25.6 healing_power points (1.62 DPS) | yes | Ghostly Mantle (3324, -0.07 DPS) [quest]; Death Speaker Mantle (6685, -0.14 DPS) [dungeon]; Talbar Mantle (10657, -0.33 DPS) [quest] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 20.7 healing_power points (1.32 DPS) | yes | Resilient Cape (14400, -0.16 DPS, sim-verified) [world_drop]; Amy's Blanket (13005, -0.55 DPS) [world_drop]; Silver-thread Cloak (4713, -0.56 DPS) [world_drop] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 43.4 healing_power points (2.75 DPS) | yes | Pristine Gown (253961, -0.13 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -0.26 DPS) [dungeon]; Silver-thread Robe (4035, -0.83 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 22.1 healing_power points (1.40 DPS) | yes | Nightsky Wristbands (6407, -0.15 DPS, sim-verified) [world_drop]; Glowing Magical Bracelets (13106, -0.32 DPS) [world_drop]; Spidertank Oilrag (9448, -0.68 DPS) [dungeon] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 31.2 healing_power points (1.98 DPS) | yes | Hotshot Pilot's Gloves (9491, +0.00 DPS, sim-verified) [dungeon]; Blight Gloves (279877, -0.49 DPS) [quest]; Truefaith Gloves (7049, -0.62 DPS) [crafted] |
| waist | Lilac Sash (6780) | Centaur Bounty [quest] | 24.3 healing_power points (1.54 DPS) | yes | Resilient Cord (14406, -0.29 DPS, sim-verified) [world_drop]; Pristine Sash (253925, -0.30 DPS) [crafted]; Dreamer's Belt (4829, -0.38 DPS) [vendor] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 45.5 healing_power points (2.89 DPS) | yes | Filigreed Pristine Leggings (253937, -0.48 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -0.74 DPS) [crafted]; Sacred Burial Trousers (6282, -1.23 DPS) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 41.7 healing_power points (2.65 DPS) | yes | Frothing Slippers (254003, -1.27 DPS) [crafted]; Black Slippers (254007, -1.27 DPS) [crafted]; Acidic Walkers (9454, -1.62 DPS, sim-verified) [dungeon] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 18.0 healing_power points (1.14 DPS) | yes | The Queen's Jewel (13094, -0.00 DPS) [world_drop]; Solid Gold Band (279845, -0.15 DPS) [quest]; Defias Mage Ring (1077, -0.28 DPS) [world] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 18.0 healing_power points (1.14 DPS) | yes | The Queen's Jewel (13094, -0.00 DPS) [world_drop]; Solid Gold Band (279845, -0.15 DPS) [quest]; Defias Mage Ring (1077, -0.28 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-verified (67.1 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Staff of the Friar (3415, -0.57 DPS) [dungeon]; Death Speaker Scepter (2816, -1.72 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 16.1 healing_power points (1.02 DPS) | yes | Summoner's Wand (5245, -0.18 DPS, sim-verified) [world]; Lesser Mystic Wand (11289, -0.48 DPS) [crafted]; Starfaller (13063, -0.48 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Nightsky Mantle; back: Glowing Thresher Cape; chest: Death Speaker Robes; hands: Gloves of Old; waist: Lilac Sash; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Spine of the Devourer; ranged: Wand of Mana Concentration

No-known-source sample (15 of 280, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (undead, 025005031330011520-00000000000000000-00000000000000000)

Set DPS (verified): 128.4. Weights run: 12.8s. Verify run: 82.6s. 382 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.020, intellect=2.098 ± 0.021, spirit=0.809 ± 0.036, mp5=2.286 ± 0.038, crit=0.418 ± 0.013 per rating point (14 rating = 1%, 5.857 per %), spell_haste=-10.189 ± 0.391

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 57.0 healing_power points (4.38 DPS) | yes | Whitemane's Chapeau (7720, -1.02 DPS) [dungeon]; Holy Shroud (2721, -1.47 DPS) [world_drop]; Thinking Cap (2624, -1.64 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 29.8 healing_power points (2.29 DPS) | yes | Necklace of Calisea (1714, -0.73 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.32 DPS) [world]; Kaleidoscope Chain (13084, -1.40 DPS) [world_drop] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 42.3 healing_power points (3.25 DPS) | yes | Windchaser Amice (14432, -1.15 DPS) [world_drop]; Mistscape Mantle (4734, -1.17 DPS) [dungeon]; Silksand Shoulder Pads (14423, -1.29 DPS) [world_drop] |
| back | Blackforge Cape (6424) | Uldaman: Ancient Treasure [dungeon] | 20.0 healing_power points (1.54 DPS) | yes | Windchaser Cloak (14430, -0.32 DPS) [world_drop]; Glyphed Cloak (4732, -0.41 DPS) [world_drop]; Glowing Thresher Cape (6901, -0.50 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 64.5 healing_power points (4.96 DPS) | yes | Death Speaker Robes (6682, -1.65 DPS) [dungeon]; Pristine Gown (253961, -2.04 DPS) [crafted]; Red Mageweave Vest (10007, -2.06 DPS) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 18.9 healing_power points (1.45 DPS) | yes | Mindthrust Bracers (1974, -0.12 DPS) [dungeon]; Mistscape Bracers (4045, -0.16 DPS) [dungeon]; Enchanted Stonecloth Bracers (4979, -0.16 DPS) [quest] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gilded Handwraps (254021, +0.00 DPS) [crafted]; Earthen Silk Gloves (254017, -0.90 DPS) [crafted]; Truefaith Gloves (7049, -1.15 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 43.6 healing_power points (3.36 DPS) | yes | Deathmage Sash (10771, -0.94 DPS) [dungeon]; Sutarn's Ring (13105, -1.37 DPS) [world_drop]; Razzeric's Customized Seatbelt (6726, -1.42 DPS) [quest] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stoneweaver Leggings (9407, +0.00 DPS) [dungeon]; Filigreed Pristine Leggings (253937, +0.00 DPS) [crafted]; Pristine Leggings (253987, +0.00 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 37.9 healing_power points (2.92 DPS) | yes | Furen's Boots (13100, -0.79 DPS) [world_drop]; Boots of the Maharishi (9658, -1.09 DPS) [quest]; Kodo Rustler Boots (15697, -1.32 DPS) [quest] |
| finger1 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 17.1 healing_power points (1.32 DPS) | yes | Mindbender Loop (5009, -0.06 DPS) [world_drop]; Welken Ring (5011, -0.07 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.35 DPS) [dungeon] |
| finger2 | Ogremind Ring (1993) | World drop [world_drop] | 17.1 healing_power points (1.32 DPS) | yes | Mindbender Loop (5009, -0.06 DPS) [world_drop]; Welken Ring (5011, -0.07 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.35 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, -2.50 DPS) [dungeon]; Hypnotic Blade (7714, -2.87 DPS) [dungeon]; Illusionary Rod (7713, -2.94 DPS) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | sim-verified (128.4 DPS) | yes | Aurora Sphere (7610, +0.00 DPS) [dungeon]; Beacon of Hope (9393, -0.05 DPS) [dungeon]; Witch's Finger (16887, -0.08 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 15.0 healing_power points (1.16 DPS) | yes | Wand of Mana Concentration (273043, -0.16 DPS) [world]; Flash Wand (5248, -0.32 DPS) [quest]; Goblin Igniter (5253, -0.32 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Blackforge Cape; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Stormcloth Gloves; waist: Gilded Cord; legs: Stormcloth Pants; finger1: Voodoo Band; finger2: Ogremind Ring; trinket1: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Orb of Souls; ranged: Jaina's Firestarter

No-known-source sample (15 of 382, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16336 Sergeant Major's Cape

### Band 50 (undead, 025005031330011520-03502000000000000-00000000000000000)

Set DPS (verified): 205.1. Weights run: 12.7s. Verify run: 46.1s. 501 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.031, intellect=2.648 ± 0.032, spirit=1.981 ± 0.049, mp5=4.589 ± 0.072, crit=0.639 ± 0.021 per rating point (14 rating = 1%, 8.942 per %), spell_haste=-11.820 ± 0.573

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 89.7 healing_power points (6.55 DPS) | yes | Blood Guard's Satin Cover (220899, -0.24 DPS) [vendor]; Papal Fez (9431, -0.94 DPS) [dungeon]; Imperial Red Circlet (8254, -1.28 DPS) [world_drop] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 45.0 healing_power points (3.29 DPS) | yes | Glowing Eye of Mordresh (10769, -0.79 DPS) [dungeon]; Necklace of Calisea (1714, -0.92 DPS) [dungeon]; Choker of the High Shaman (4112, -1.16 DPS) [quest] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | 59.1 healing_power points (4.31 DPS) | yes | Blood Guard's Satin Pads (220901, +0.00 DPS) [vendor]; Inquisitor's Shawl (19507, -0.70 DPS) [dungeon]; Rotgrip Mantle (17732, -0.83 DPS) [dungeon] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 37.1 healing_power points (2.71 DPS) | yes | Highborne Cloak (14450, -0.39 DPS) [world_drop]; Blackforge Cape (6424, -0.58 DPS) [dungeon]; Spritecaster Cape (11623, -0.68 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-verified (205.2 DPS) | yes | Ghostweave Vest (14141, +0.00 DPS) [crafted]; Stone Guard's Satin Armor (220903, +0.00 DPS) [vendor]; Earth Warder's Vest (21311, -0.87 DPS) [quest] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 48.4 healing_power points (3.54 DPS) | yes | Shizzle's Nozzle Wiper (11917, -0.78 DPS) [quest]; Forgotten Wraps (9433, -1.22 DPS) [world_drop]; Bloodband Bracers (11469, -1.36 DPS) [quest] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 71.3 healing_power points (5.21 DPS) | yes | Greenleaf Handwraps (19116, -1.48 DPS) [quest]; Earthenweave Gloves (254075, -1.72 DPS) [crafted]; Highborne Gloves (14451, -1.83 DPS) [world_drop] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 62.3 healing_power points (4.55 DPS) | yes | Gilded Cord (254037, -0.53 DPS) [crafted]; Gilded Waistcord (254081, -0.69 DPS) [crafted]; Earthenweave Cord (254077, -1.06 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 89.3 healing_power points (6.52 DPS) | yes | Senior Designer's Pantaloons (11841, -0.74 DPS) [dungeon]; Stone Guard's Satin Leggings (220902, -1.08 DPS) [vendor]; Dalewind Trousers (13008, -1.17 DPS) [world_drop] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | sim-verified (205.2 DPS) | yes | Vinerot Sandals (17748, +0.00 DPS) [dungeon]; First Sergeant's Satin Boots (220900, -0.22 DPS) [vendor]; Furen's Boots (13100, -0.34 DPS) [world_drop] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 39.7 healing_power points (2.90 DPS) | yes | Welken Ring (5011, -0.92 DPS) [world_drop]; Ogremind Ring (1993, -1.11 DPS) [world_drop]; Voodoo Band (1996, -1.11 DPS) [world] |
| finger2 | Coldwater Ring (4550) | Death From Below [quest] | sim-verified (205.2 DPS) | yes | Welken Ring (5011, +0.00 DPS) [world_drop]; Ogremind Ring (1993, -0.09 DPS) [world_drop]; Voodoo Band (1996, -0.09 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.29 DPS) [quest]; Alchemist's Stone (13503, -1.16 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Alchemist's Stone (13503, -0.74 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -6.60 DPS) [dungeon]; Spellshifter Rod (9527, -9.47 DPS) [quest]; Hand of Righteousness (7721, -9.84 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 21.8 healing_power points (1.59 DPS) | yes | Wand of Mana Concentration (273043, -0.15 DPS) [world]; Flash Wand (5248, -0.39 DPS) [quest]; Goblin Igniter (5253, -0.39 DPS) [quest] |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Nethergeld Shoulders; back: Imperial Red Cloak; wrist: Nethergeld Cuffs; hands: Gilded Gloves; waist: Dawnspire Cord; legs: Kilt of the Atal'ai Prophet; feet: Gilded Sandals; finger1: Eye of Adaegus; finger2: Coldwater Ring; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 501, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60 (undead, 025005031330011520-03505003030100000-00000000000000000)

Set DPS (verified): 303.8. Weights run: 12.3s. Verify run: 48.8s. 1187 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.133, intellect=2.713 ± 0.071, spirit=3.888 ± 0.434, mp5=3.203 ± 0.192, crit=0.850 ± 0.041 per rating point (14 rating = 1%, 11.897 per %), spell_haste=-12.342 ± 1.234

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | 143.4 healing_power points (11.76 DPS) | yes | Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Champion's Satin Hood (227118, -0.01 DPS) [pvp]; Gemburst Circlet (10751, -10.00 DPS, sim-verified) [quest] |
| neck | Lady Alizabeth's Pendant (13002) | World drop [world_drop] | 58.3 healing_power points (4.78 DPS) | yes | Horizon Choker (13085, +0.00 DPS, sim-verified) [world_drop]; Necklace of Calisea (1714, -0.99 DPS) [dungeon]; Choker of the High Shaman (4112, -1.12 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | 103.0 healing_power points (8.44 DPS) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, -1.53 DPS) [pvp]; Mooncloth Shoulders (14139, -7.54 DPS, sim-verified) [crafted] |
| back | Cloak of Blight (6832) (or Indomitable Cloak (14683)) | Nothing But The Truth [quest] | 50.5 healing_power points (4.14 DPS) | yes | Indomitable Cloak (14683, +0.00 DPS) [world_drop]; Wingveil Cloak (10802, -0.23 DPS) [dungeon]; Chainlink Towel (9648, -0.29 DPS) [quest] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (+4.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Truefaith Vestments (14154, -0.67 DPS) [crafted]; Mooncloth Vest (14138, -4.43 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | 71.6 healing_power points (5.87 DPS) | yes | General's Satin Bracers (17619, +0.00 DPS) [pvp]; Magister's Bindings (16683, -1.07 DPS) [dungeon]; Earthenweave Cuffs (254131, -6.89 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | 108.2 healing_power points (8.87 DPS) | yes | Magister's Gloves (16684, -1.30 DPS) [dungeon]; Gilded Gloves (254095, -1.72 DPS) [crafted]; Swarmtender's Gloves (275607, -8.04 DPS, sim-verified) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (+5.9 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Satin Cinch (17621, +0.00 DPS) [pvp]; Elunarian Belt (14465, -0.58 DPS) [world_drop]; Wisdom of the Timbermaw (19047, -5.90 DPS, sim-verified) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 154.1 healing_power points (12.64 DPS) | yes | Legionnaire's Satin Legguards (227123, -1.56 DPS) [pvp]; General's Satin Legguards (231634, -1.90 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -11.53 DPS, sim-verified) [dungeon] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (+4.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Elunarian Boots (14458, -1.79 DPS) [world_drop]; Furen's Boots (13100, -1.89 DPS) [world_drop]; Mooncloth Boots (15802, -4.41 DPS, sim-verified) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 69.9 healing_power points (5.73 DPS) | yes | Coldwater Ring (4550, -2.19 DPS) [quest]; Chemist's Ring (15702, -2.23 DPS) [quest]; Maiden's Circle (13001, -2.26 DPS) [world_drop] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 62.9 healing_power points (5.16 DPS) | yes | Chemist's Ring (15702, -1.66 DPS) [quest]; Maiden's Circle (13001, -1.69 DPS) [world_drop]; Coldwater Ring (4550, -1.72 DPS, sim-verified) [quest] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Alchemist's Stone (13503, -0.30 DPS) [crafted]; Ankh of Life (1713, -0.83 DPS, sim-verified) [world_drop] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Alchemist's Stone (13503, -0.20 DPS) [crafted]; Evonice's Landin' Pilla (18951, -0.49 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Hale Magefire (13000, -9.05 DPS) [world_drop]; Charstone Dirk (17710, -9.14 DPS) [dungeon]; Hand of Edward the Odd (2243, -14.80 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 59.8 healing_power points (4.90 DPS) | yes | Chillnail Splinter (10704, -2.99 DPS) [quest]; Goblin Igniter (5253, -3.05 DPS) [quest]; Jaina's Firestarter (13064, -3.41 DPS, sim-verified) [world_drop] |

**New at 60:** head: Devout Crown; neck: Lady Alizabeth's Pendant; shoulder: Devout Mantle; back: Cloak of Blight; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Shard of the Splithooves; ranged: Torch of Light

No-known-source sample (15 of 1187, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60, raid preset (undead, 025005031330011520-03505003030100000-00000000000000000)

Set DPS (verified): 561.2. Weights run: 8.3s. Verify run: 46.9s. 1187 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.128, intellect=1.212 ± 0.028, spirit=1.313 ± 0.026, mp5=2.328 ± 0.030, crit=0.593 ± 0.023 per rating point (14 rating = 1%, 8.308 per %), spell_haste=-5.523 ± 0.525

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (561.3 DPS) | yes | Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -11.09 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 22.7 healing_power points (5.28 DPS) | yes | Lady Alizabeth's Pendant (13002, -0.70 DPS) [world_drop]; Necklace of Calisea (1714, -1.17 DPS) [dungeon]; Horizon Choker (13085, -2.57 DPS, sim-verified) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (561.3 DPS) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp]; Mooncloth Shoulders (14139, -9.90 DPS, sim-verified) [crafted] |
| back | Wingveil Cloak (10802) | Sunken Temple: Spawn of Hakkar [dungeon] | 19.4 healing_power points (4.52 DPS) | yes | Imperial Red Cloak (8248, +0.00 DPS, sim-verified) [world_drop]; Glowing Thresher Cape (6901, -0.45 DPS) [dungeon]; Cloak of Blight (6832, -0.55 DPS) [quest] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 94.0 healing_power points (21.85 DPS) | yes | Warlord's Satin Tunic (231632, -1.69 DPS) [vendor]; Legionnaire's Satin Tunic (227124, -4.06 DPS) [pvp]; Mooncloth Vest (14138, -6.35 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (561.3 DPS) | yes | General's Satin Bracers (17619, +0.00 DPS) [pvp]; Nethergeld Cuffs (254061, +0.00 DPS) [crafted]; Earthenweave Cuffs (254131, -10.58 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (561.3 DPS) | yes | Blood Guard's Satin Handwraps (227126, -0.79 DPS) [pvp]; General's Satin Handwraps (231633, -1.09 DPS) [vendor]; Gilded Gloves (254095, -8.59 DPS, sim-verified) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 59.8 healing_power points (13.90 DPS) | yes | Devout Belt (16696, -2.29 DPS, sim-verified) [dungeon]; General's Satin Cinch (17621, -2.84 DPS) [pvp]; Earthenweave Cord (254077, -4.50 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 72.4 healing_power points (16.83 DPS) | yes | Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Mooncloth Leggings (14137, -16.00 DPS, sim-verified) [crafted] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 58.0 healing_power points (13.49 DPS) | yes | General's Satin Walkers (231630, -1.68 DPS) [vendor]; Devout Sandals (16691, -2.40 DPS, sim-verified) [dungeon]; Blood Guard's Satin Walkers (227127, -3.22 DPS) [pvp] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 26.6 healing_power points (6.18 DPS) | yes | Chemist's Ring (15702, -2.14 DPS) [quest]; Maiden's Circle (13001, -2.37 DPS) [world_drop]; Coldwater Ring (4550, -2.58 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 23.0 healing_power points (5.35 DPS) | yes | Chemist's Ring (15702, -1.32 DPS) [quest]; Maiden's Circle (13001, -1.55 DPS) [world_drop]; Coldwater Ring (4550, -1.76 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (561.3 DPS) | yes | Evonice's Landin' Pilla (18951, -1.13 DPS) [quest]; Alchemist's Stone (13503, -2.96 DPS) [crafted]; Ankh of Life (1713, -7.23 DPS, sim-verified) [world_drop] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (561.3 DPS) | yes | Evonice's Landin' Pilla (18951, -0.19 DPS) [quest]; Ankh of Life (1713, -1.89 DPS, sim-verified) [world_drop]; Alchemist's Stone (13503, -2.03 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (561.3 DPS) | yes | Charstone Dirk (17710, -14.64 DPS) [dungeon]; Hand of Righteousness (7721, -21.49 DPS) [dungeon]; Hand of Edward the Odd (2243, -23.22 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 31.4 healing_power points (7.31 DPS) | yes | Jaina's Firestarter (13064, -4.10 DPS, sim-verified) [world_drop]; Captain Rackmore's Tiller (16789, -5.00 DPS) [quest]; Wand of Mana Concentration (273043, -5.10 DPS) [world] |

**New at 60:** head: Devout Crown; neck: Glowing Eye of Mordresh; shoulder: Devout Mantle; back: Wingveil Cloak; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Uther's Strength; trinket2: Shard of the Splithooves; ranged: Torch of Light

No-known-source sample (15 of 1187, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

