# Leveling BiS: Discipline

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 025004000000000000-00000000000000000-00000000000000000)

Set DPS (verified): 36.5. Weights run: 10.7s. Verify run: 29.1s. 146 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.006, intellect=2.689 ± 0.012, spirit=1.300 ± 0.007, mp5=2.938 ± 0.024, crit=0.146 ± 0.007 per rating point (14 rating = 1%, 2.043 per %), spell_haste=0.647 ± 0.084

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Goggles (4373) | Engineering [crafted] | 21.2 healing_power points (1.03 DPS) | yes | Pristine Circlet (253949, +0.00 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -0.78 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 5.2 healing_power points (0.25 DPS) | yes | Scholarly Pendant (277203, +0.00 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 24.2 healing_power points (1.18 DPS) | yes | Slime-encrusted Pads (6461, -0.75 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -0.85 DPS, sim-verified) [crafted] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | 11.6 healing_power points (0.56 DPS) | yes | Sanguine Cape (14376, -0.06 DPS, sim-verified) [world_drop]; Pearl-clasped Cloak (5542, -0.17 DPS) [crafted]; Seer's Cape (6378, -0.18 DPS) [dungeon] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 30.3 healing_power points (1.48 DPS) | yes | Corsair's Overshirt (5202, -0.44 DPS) [dungeon]; Seer's Robe (2981, -0.50 DPS) [world_drop]; Robe of the Moccasin (6465, -0.70 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 22.3 healing_power points (1.08 DPS) | yes | Bright Bracers (3647, -0.69 DPS, sim-verified) [world_drop]; Repurposed Hair Band (281256, -0.70 DPS) [quest]; Mystic's Bracelets (14366, -0.82 DPS) [world_drop] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 19.1 healing_power points (0.93 DPS) | yes | Tomb Robber's Gloves (280096, -0.14 DPS) [quest]; Bright Gloves (3066, -0.15 DPS) [world_drop]; Magefist Gloves (12977, -0.30 DPS, sim-verified) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 21.8 healing_power points (1.06 DPS) | yes | Keller's Girdle (2911, -0.09 DPS, sim-verified) [world_drop]; Novice Ardent's Sash (253887, -0.18 DPS) [crafted]; Wise Man's Belt (4786, -0.40 DPS) [vendor] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 39.3 healing_power points (1.91 DPS) | yes | Abomination Skin Leggings (23173, -0.87 DPS) [dungeon]; Filigreed Silky Leggings (253939, -0.88 DPS) [crafted]; Darkweave Breeches (12987, -1.08 DPS, sim-verified) [world_drop] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 17.1 healing_power points (0.83 DPS) | yes | Spidersilk Boots (4320, -0.31 DPS) [crafted]; Sanguine Sandals (14374, -0.31 DPS) [world_drop]; Kimbra Boots (6191, -0.43 DPS, sim-verified) [quest] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 16.1 healing_power points (0.79 DPS) | yes | Band of Purification (12996, -0.41 DPS) [world_drop]; Deep Fathom Ring (6463, -0.47 DPS) [dungeon]; Lorekeeper's Ring (20431, -0.50 DPS) [rep] |
| finger2 | Volcanic Rock Ring (12053) | World drop [world_drop] | 8.1 healing_power points (0.39 DPS) | yes | Band of Purification (12996, -0.05 DPS, sim-verified) [world_drop]; Deep Fathom Ring (6463, -0.08 DPS) [dungeon]; Lorekeeper's Ring (20431, -0.11 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 26.9 healing_power points (1.31 DPS) | yes | Staff of Westfall (2042, -0.44 DPS, sim-verified) [quest]; Lesser Staff of the Spire (1300, -0.52 DPS) [world]; Witching Stave (1484, -0.65 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Flaring Baton (5326) (or Moonstone Wand (15204)) | The Escape [quest] | 5.4 healing_power points (0.26 DPS) | yes | Moonstone Wand (15204, +0.00 DPS) [quest]; Sable Wand (7607, -0.07 DPS) [quest]; Elven Wand (5604, -0.13 DPS) [quest] |

**New at 20:** head: Shadow Goggles; neck: Tarnished Locket; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Volcanic Rock Ring; main_hand: Twisted Chanter's Staff; ranged: Flaring Baton

No-known-source sample (15 of 146, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes; 16607 Acolyte's Sacrificial Robes; 18851 Insignia of the Horde; 20434 Lorekeeper's Staff

### Band 30 (human, 025005031320000000-00000000000000000-00000000000000000)

Set DPS (verified): 67.6. Weights run: 11.0s. Verify run: 32.8s. 258 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.016, intellect=2.126 ± 0.014, spirit=1.720 ± 0.011, mp5=3.816 ± 0.013, crit=0.212 ± 0.009 per rating point (14 rating = 1%, 2.965 per %), spell_haste=not significant (-0.112 ± 0.086)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 43.3 healing_power points (2.74 DPS) | yes | Resilient Cap (14401, -0.82 DPS) [world_drop]; Nightsky Cowl (4039, -0.85 DPS, sim-verified) [world_drop]; Shadow Hood (4323, -1.26 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.4 healing_power points (0.97 DPS) | yes | Scorn's Icy Choker (23169, -0.17 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -0.17 DPS) [world]; Crystal Starfire Medallion (5003, -0.23 DPS, sim-verified) [world_drop] |
| shoulder | Mantle of Honor (3560) | Bride of the Embalmer [quest] | 26.9 healing_power points (1.70 DPS) | yes | Nightsky Mantle (4718, -0.08 DPS) [world_drop]; Death Speaker Mantle (6685, -0.22 DPS) [dungeon]; Faerie Mantle (5820, -0.24 DPS) [quest] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | 22.2 healing_power points (1.40 DPS) | yes | Prelacy Cape (7004, +0.00 DPS, sim-verified) [quest]; Glowing Thresher Cape (6901, -0.09 DPS) [dungeon]; Caretaker's Cape (19533, -0.14 DPS) [rep] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 43.4 healing_power points (2.74 DPS) | yes | Pristine Gown (253961, -0.18 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -0.26 DPS) [dungeon]; Silver-thread Robe (4035, -0.83 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 22.1 healing_power points (1.40 DPS) | yes | Nightsky Wristbands (6407, -0.21 DPS, sim-verified) [world_drop]; Glowing Magical Bracelets (13106, -0.32 DPS) [world_drop]; Spidertank Oilrag (9448, -0.67 DPS) [dungeon] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 31.3 healing_power points (1.98 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.36 DPS) [dungeon]; Truefaith Gloves (7049, -0.62 DPS) [crafted]; Zodiac Gloves (7106, -0.68 DPS) [quest] |
| waist | Resilient Cord (14406) | World drop [world_drop] | 19.6 healing_power points (1.24 DPS) | yes | Pristine Sash (253925, +0.00 DPS, sim-verified) [crafted]; Dreamer's Belt (4829, -0.08 DPS) [vendor]; Novice Ardent's Sash (253887, -0.13 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 45.5 healing_power points (2.87 DPS) | yes | Filigreed Pristine Leggings (253937, -0.55 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -0.73 DPS) [crafted]; Darkweave Breeches (12987, -1.28 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 41.8 healing_power points (2.64 DPS) | yes | Soggy Boots (274747, -1.23 DPS) [vendor]; Frothing Slippers (254003, -1.26 DPS) [crafted]; Acidic Walkers (9454, -1.54 DPS, sim-verified) [dungeon] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 18.0 healing_power points (1.14 DPS) | yes | The Queen's Jewel (13094, -0.00 DPS) [world_drop]; Sea Giant's Toe Ring (274746, -0.00 DPS) [vendor]; Darkspear Signet (272071, -0.18 DPS) [vendor] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 18.0 healing_power points (1.14 DPS) | yes | The Queen's Jewel (13094, +0.00 DPS, sim-verified) [world_drop]; Sea Giant's Toe Ring (274746, -0.00 DPS) [vendor]; Darkspear Signet (272071, -0.18 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-verified (67.6 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -0.03 DPS) [quest]; Death Speaker Scepter (2816, -1.76 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 16.1 healing_power points (1.02 DPS) | yes | Summoner's Wand (5245, -0.22 DPS, sim-verified) [world]; Lesser Mystic Wand (11289, -0.48 DPS) [crafted]; Starfaller (13063, -0.48 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Mantle of Honor; back: Darkspear Raider's Cloak; chest: Death Speaker Robes; hands: Gloves of Old; waist: Resilient Cord; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Spine of the Devourer; ranged: Wand of Mana Concentration

No-known-source sample (15 of 258, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 40 (human, 025005031330011520-00000000000000000-00000000000000000)

Set DPS (verified): 134.9. Weights run: 12.6s. Verify run: 80.2s. 340 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.021, intellect=2.089 ± 0.021, spirit=0.812 ± 0.036, mp5=2.297 ± 0.038, crit=0.411 ± 0.013 per rating point (14 rating = 1%, 5.761 per %), spell_haste=-10.189 ± 0.388

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 56.8 healing_power points (4.35 DPS) | yes | Whitemane's Chapeau (7720, -0.96 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -1.45 DPS) [world_drop]; Thinking Cap (2624, -1.63 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 29.7 healing_power points (2.27 DPS) | yes | Necklace of Calisea (1714, -0.27 DPS, sim-verified) [dungeon]; Darkspear Warding Pendant (272074, -1.15 DPS) [vendor]; Archmagister's Faceted Pendant (273049, -1.31 DPS) [world] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 42.2 healing_power points (3.23 DPS) | yes | Mistscape Mantle (4734, -1.16 DPS) [dungeon]; Silksand Shoulder Pads (14423, -1.28 DPS) [world_drop]; Windchaser Amice (14432, -1.48 DPS, sim-verified) [world_drop] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 26.2 healing_power points (2.01 DPS) | yes | Caretaker's Cape (19532, -0.32 DPS) [rep]; Blackforge Cape (6424, -0.48 DPS) [dungeon]; Sergeant Major's Cape (16336, -0.68 DPS) [pvp] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 64.5 healing_power points (4.94 DPS) | yes | Pristine Gown (253961, -2.04 DPS) [crafted]; Red Mageweave Vest (10007, -2.06 DPS) [crafted]; Death Speaker Robes (6682, -5.34 DPS, sim-verified) [dungeon] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 18.8 healing_power points (1.44 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Mistscape Bracers (4045, -0.16 DPS) [dungeon]; Enchanted Stonecloth Bracers (4979, -0.16 DPS) [quest] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (134.9 DPS) | yes | Earthen Silk Gloves (254017, -0.89 DPS) [crafted]; Truefaith Gloves (7049, -1.13 DPS) [crafted]; Gilded Handwraps (254021, -1.83 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 43.6 healing_power points (3.34 DPS) | yes | Sutarn's Ring (13105, -1.36 DPS) [world_drop]; Razzeric's Customized Seatbelt (6726, -1.42 DPS) [quest]; Deathmage Sash (10771, -2.14 DPS, sim-verified) [dungeon] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-verified (134.9 DPS) | yes | Stoneweaver Leggings (9407, +0.00 DPS) [dungeon]; Filigreed Pristine Leggings (253937, +0.00 DPS) [crafted]; Pristine Leggings (253987, -1.89 DPS, sim-verified) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 37.9 healing_power points (2.90 DPS) | yes | Furen's Boots (13100, +0.00 DPS, sim-verified) [world_drop]; Kodo Rustler Boots (15697, -1.31 DPS) [quest]; Acidic Walkers (9454, -1.37 DPS) [dungeon] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.38 DPS) | yes | Ogremind Ring (1993, -0.07 DPS) [world_drop]; Mindbender Loop (5009, -0.13 DPS) [world_drop]; Welken Ring (5011, -0.14 DPS) [world_drop] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 17.1 healing_power points (1.31 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Mindbender Loop (5009, -0.06 DPS) [world_drop]; Welken Ring (5011, -0.07 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (134.9 DPS) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (134.9 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Death Speaker Scepter (2816, -2.49 DPS) [dungeon]; Hypnotic Blade (7714, -2.86 DPS) [dungeon] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 32.9 healing_power points (2.52 DPS) | yes | Orb of Souls (249395, -0.69 DPS, sim-verified) [crafted]; Aurora Sphere (7610, -1.15 DPS) [dungeon]; Beacon of Hope (9393, -1.20 DPS) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 15.0 healing_power points (1.15 DPS) | yes | Wand of Mana Concentration (273043, -0.15 DPS) [world]; Flash Wand (5248, -0.32 DPS) [quest]; Goblin Igniter (5253, -0.32 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Stormcloth Gloves; waist: Gilded Cord; legs: Stormcloth Pants; finger1: Sea Giant's Toe Ring; finger2: Voodoo Band; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Orb of Lorica; ranged: Jaina's Firestarter

No-known-source sample (15 of 340, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 50 (human, 025005031330011520-03502000000000000-00000000000000000)

Set DPS (verified): 216.4. Weights run: 12.5s. Verify run: 41.8s. 434 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.039, intellect=2.560 ± 0.032, spirit=1.896 ± 0.047, mp5=4.383 ± 0.068, crit=0.615 ± 0.019 per rating point (14 rating = 1%, 8.612 per %), spell_haste=-11.402 ± 0.547

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 86.6 healing_power points (6.60 DPS) | yes | Papal Fez (9431, -0.92 DPS) [dungeon]; Imperial Red Circlet (8254, -1.31 DPS) [world_drop]; Knight-Lieutenant's Satin Cover (220896, -5.30 DPS, sim-verified) [vendor] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 43.4 healing_power points (3.31 DPS) | yes | Glowing Eye of Mordresh (10769, -0.76 DPS) [dungeon]; Necklace of Calisea (1714, -0.93 DPS) [dungeon]; Darkmoon Necklace (19303, -4.00 DPS, sim-verified) [vendor] |
| shoulder | Knight-Lieutenant's Satin Pads (220894) | Captain Dirgehammer [vendor] | 65.1 healing_power points (4.96 DPS) | yes | Inquisitor's Shawl (19507, -1.28 DPS) [dungeon]; Rotgrip Mantle (17732, -1.45 DPS) [dungeon]; Nethergeld Shoulders (254049, -4.11 DPS, sim-verified) [crafted] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 45.3 healing_power points (3.45 DPS) | yes | Caretaker's Cape (19531, -0.91 DPS) [rep]; Highborne Cloak (14450, -1.12 DPS) [world_drop]; Imperial Red Cloak (8248, -1.67 DPS, sim-verified) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 73.2 healing_power points (5.57 DPS) | yes | Knight's Satin Armor (220892, -0.02 DPS) [vendor]; Earth Warder's Vest (21311, -1.03 DPS) [quest]; Ghostweave Vest (14141, -4.83 DPS, sim-verified) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 47.4 healing_power points (3.61 DPS) | yes | Forgotten Wraps (9433, -1.27 DPS) [world_drop]; Bloodband Bracers (11469, -1.42 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -4.59 DPS, sim-verified) [quest] |
| hands | Virtuous Mitts (226950) | Mokvar [vendor] | 86.3 healing_power points (6.57 DPS) | yes | Gilded Gloves (254095, -1.26 DPS) [crafted]; Virtuous Hands (226958, -2.30 DPS) [vendor]; Raider Handwraps (272098, -6.18 DPS, sim-verified) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 60.6 healing_power points (4.62 DPS) | yes | Gilded Waistcord (254081, -0.65 DPS) [crafted]; Earthenweave Cord (254077, -1.05 DPS) [crafted]; Gilded Cord (254037, -4.93 DPS, sim-verified) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 86.2 healing_power points (6.56 DPS) | yes | Senior Designer's Pantaloons (11841, -0.58 DPS, sim-verified) [dungeon]; Knight's Satin Leggings (220893, -1.02 DPS) [vendor]; Dalewind Trousers (13008, -1.20 DPS) [world_drop] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 54.0 healing_power points (4.11 DPS) | yes | Sergeant Major's Satin Boots (220895, -0.25 DPS) [vendor]; Furen's Boots (13100, -0.44 DPS) [world_drop]; Vinerot Sandals (17748, -4.51 DPS, sim-verified) [dungeon] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 38.1 healing_power points (2.90 DPS) | yes | Welken Ring (5011, -0.92 DPS) [world_drop]; Coldwater Ring (4550, -1.02 DPS) [quest]; Choking Band (11868, -1.02 DPS) [quest] |
| finger2 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 30.7 healing_power points (2.34 DPS) | yes | Welken Ring (5011, -0.35 DPS) [world_drop]; Coldwater Ring (4550, -0.45 DPS) [quest]; Choking Band (11868, -0.46 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (216.4 DPS) | yes | Ankh of Life (1713, -3.27 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -3.56 DPS) [quest]; Uther's Strength (11302, -3.63 DPS) [world_drop] |
| trinket2 | Alchemist's Stone (13503) | Alchemy [crafted] | sim-verified (216.4 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Uther's Strength (11302, -2.09 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (216.4 DPS) | yes | Charstone Dirk (17710, -6.73 DPS) [dungeon]; Spellshifter Rod (9527, -9.78 DPS) [quest]; Hammer of the Northern Wind (810, -31.21 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cairnstone Sliver (9654) | The Morrow Stone [quest] | 22.3 healing_power points (1.70 DPS) | yes | Wand of Mana Concentration (273043, -0.25 DPS) [world]; Goblin Igniter (5253, -0.48 DPS) [quest]; Jaina's Firestarter (13064, -3.18 DPS, sim-verified) [world_drop] |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Knight-Lieutenant's Satin Pads; back: Darkspear Raider's Cloak; wrist: Nethergeld Cuffs; hands: Virtuous Mitts; waist: Dawnspire Cord; legs: Kilt of the Atal'ai Prophet; feet: Gilded Sandals; finger1: Eye of Adaegus; finger2: Darkspear Signet; trinket1: Darkspear Voodoo Seal; trinket2: Alchemist's Stone; main_hand: Glowing Brightwood Staff; ranged: Cairnstone Sliver

No-known-source sample (15 of 434, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60 (human, 025005031330011520-03505003030100000-00000000000000000)

Set DPS (verified): 359.7. Weights run: 12.2s. Verify run: 54.9s. 948 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.138, intellect=2.726 ± 0.073, spirit=3.896 ± 0.436, mp5=3.200 ± 0.193, crit=0.872 ± 0.038 per rating point (14 rating = 1%, 12.211 per %), spell_haste=-11.957 ± 1.227

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 153.9 healing_power points (12.49 DPS) | yes | Field Marshal's Satin Hood (231622, -0.12 DPS) [vendor]; Devout Crown (16693, -0.81 DPS) [dungeon]; Lieutenant Commander's Satin Hood (227121, -0.83 DPS) [pvp] |
| neck | The Eye of Zuldazar (19593) (or The All-Seeing Eye of Zuldazar (19594)) | The Eye of Zuldazar [quest] | 90.4 healing_power points (7.34 DPS) | yes | The All-Seeing Eye of Zuldazar (19594, +0.00 DPS) [quest]; Wavefront Necklace (20685, -1.99 DPS) [world]; Drake Tooth Necklace (21531, -2.58 DPS) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 117.2 healing_power points (9.51 DPS) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Virtuous Mantle (226951, -1.05 DPS) [quest]; Devout Mantle (16695, -1.12 DPS) [dungeon] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 69.3 healing_power points (5.62 DPS) | yes | Darkspear Raider's Cloak (272063, -0.18 DPS) [vendor]; Caretaker's Cape (19530, -0.98 DPS) [rep] |
| chest | Virtuous Robe (226945) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Vest (14138, +0.00 DPS) [crafted]; Field Marshal's Satin Tunic (231624, +0.00 DPS) [vendor]; Devout Robe (16690, -0.60 DPS) [dungeon] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Satin Bracers (17606, +0.00 DPS) [pvp]; Bracers of Hope (22667, +0.00 DPS) [quest]; Devout Bracers (16697, -1.11 DPS) [dungeon] |
| hands | Virtuous Mitts (226950) | Mokvar [vendor] | 124.1 healing_power points (10.07 DPS) | yes | Devout Gloves (16692, -1.27 DPS) [dungeon]; Raider Handwraps (272097, -2.15 DPS) [vendor]; Swarmtender's Gloves (275607, -2.45 DPS) [crafted] |
| waist | Virtuous Belt (226948) | Just Compensation [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Satin Sash (17609, +0.00 DPS) [pvp]; Penitent's Cinch (272394, +0.00 DPS) [vendor]; Wisdom of the Timbermaw (19047, -0.35 DPS) [crafted] |
| legs | Virtuous Skirt (226946) | Anthion's Parting Words [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Devout Skirt (16694, +0.00 DPS) [dungeon]; Knight-Captain's Satin Legguards (227125, -0.76 DPS) [pvp]; Marshal's Satin Legguards (231626, -1.09 DPS) [vendor] |
| feet | Virtuous Sandals (226952) | Anthion's Parting Words [quest] | 139.7 healing_power points (11.34 DPS) | yes | Incandescent Mooncloth Boots (227862, -0.25 DPS) [vendor]; Mooncloth Boots (15802, -2.19 DPS) [crafted]; Devout Sandals (16691, -2.86 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 75.5 healing_power points (6.13 DPS) | yes | Band of the Hierophant (13096, -0.44 DPS) [world_drop]; Channeler's Ring (272406, -0.50 DPS) [vendor]; Eye of Adaegus (5266, -1.01 DPS) [world_drop] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-verified (359.7 DPS) | yes | Eye of Adaegus (5266, +0.00 DPS) [world_drop]; Band of the Hierophant (13096, +0.00 DPS) [world_drop]; Channeler's Ring (272406, +0.00 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (359.7 DPS) | yes | Ankh of Life (1713, -0.10 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.73 DPS) [quest]; Thunderbrew's Boot Flask (744, -1.37 DPS) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.25 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.88 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Edward the Odd (2243, -8.09 DPS) [world_drop]; Staff of Hale Magefire (13000, -8.95 DPS) [world_drop]; Charstone Dirk (17710, -9.07 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 59.9 healing_power points (4.86 DPS) | yes | Cairnstone Sliver (9654, -2.17 DPS) [quest]; Jaina's Firestarter (13064, -2.58 DPS) [world_drop]; Goblin Igniter (5253, -3.02 DPS) [quest] |

**New at 60:** head: Virtuous Crown; neck: The Eye of Zuldazar; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Virtuous Robe; wrist: Virtuous Bracers; waist: Virtuous Belt; legs: Virtuous Skirt; feet: Virtuous Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Blessed Band of Light; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 948, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60, raid preset (human, 025005031330011520-03505003030100000-00000000000000000)

Set DPS (verified): 641.3. Weights run: 8.2s. Verify run: 27.6s. 948 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.152, intellect=1.369 ± 0.031, spirit=1.520 ± 0.029, mp5=2.683 ± 0.035, crit=0.708 ± 0.025 per rating point (14 rating = 1%, 9.909 per %), spell_haste=-6.464 ± 0.590

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 90.5 healing_power points (18.22 DPS) | yes | Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Lieutenant Commander's Satin Hood (227121, -0.92 DPS) [pvp]; Devout Crown (16693, -2.99 DPS) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 53.7 healing_power points (10.81 DPS) | yes | The Eye of Zuldazar (19593, -1.00 DPS) [quest]; The All-Seeing Eye of Zuldazar (19594, -1.00 DPS) [quest]; Drake Tooth Necklace (21531, -2.26 DPS) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 90.3 healing_power points (18.18 DPS) | yes | Lieutenant Commander's Satin Mantle (227119, -4.35 DPS) [pvp]; Virtuous Mantle (226951, -6.49 DPS) [quest]; Field Marshal's Satin Mantle (231628, -6.63 DPS) [pvp] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 55.7 healing_power points (11.21 DPS) | yes | Drape of Recovery (272413, -2.81 DPS) [vendor]; Caretaker's Cape (19530, -3.53 DPS) [rep] |
| chest | Virtuous Robe (226945) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Truefaith Vestments (14154, +0.00 DPS) [crafted]; Field Marshal's Satin Tunic (231624, +0.00 DPS) [vendor]; Mooncloth Vest (14138, -1.88 DPS) [crafted] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bracers of Hope (22667, +0.00 DPS) [quest]; Marshal's Satin Bracers (17606, -1.17 DPS) [pvp]; Earthenweave Cuffs (254131, -1.98 DPS) [crafted] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-verified (641.4 DPS) | yes | Desert Bloom Gloves (20717, +0.00 DPS) [quest]; Virtuous Mitts (226950, +0.00 DPS) [vendor]; Mooncloth Gloves (18409, -1.22 DPS) [crafted] |
| waist | Virtuous Belt (226948) | Just Compensation [quest] | sim-verified (641.4 DPS) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Penitent's Cinch (272394, -0.39 DPS) [vendor]; Marshal's Satin Sash (17609, -1.43 DPS) [pvp] |
| legs | Virtuous Skirt (226946) | Anthion's Parting Words [quest] | 80.3 healing_power points (16.16 DPS) | yes | Knight-Captain's Satin Legguards (227125, +0.00 DPS) [pvp]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Devout Skirt (16694, -0.15 DPS) [dungeon] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 84.5 healing_power points (17.00 DPS) | yes | Virtuous Sandals (226952, -2.19 DPS) [quest]; Mooncloth Boots (15802, -4.33 DPS) [crafted]; Marshal's Satin Walkers (231627, -6.20 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 60.7 healing_power points (12.23 DPS) | yes | Blessed Band of Light (272407, -3.54 DPS) [vendor]; Channeler's Ring (272406, -5.92 DPS) [vendor]; Band of the Hierophant (13096, -6.10 DPS) [world_drop] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 46.4 healing_power points (9.35 DPS) | yes | Blessed Band of Light (272407, -0.66 DPS) [vendor]; Channeler's Ring (272406, -3.04 DPS) [vendor]; Band of the Hierophant (13096, -3.22 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (641.4 DPS) | yes | Ankh of Life (1713, -4.43 DPS) [world_drop]; Uther's Strength (11302, -4.48 DPS) [world_drop]; Shard of the Splithooves (10659, -4.86 DPS) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -4.78 DPS) [world_drop]; Uther's Strength (11302, -4.83 DPS) [world_drop]; Shard of the Splithooves (10659, -5.21 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -13.48 DPS) [dungeon]; Hand of Edward the Odd (2243, -16.18 DPS) [world_drop]; Hand of Righteousness (7721, -19.57 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 33.7 healing_power points (6.79 DPS) | yes | Cairnstone Sliver (9654, -3.88 DPS) [quest]; Jaina's Firestarter (13064, -4.22 DPS) [world_drop]; Wand of Mana Concentration (273043, -4.61 DPS) [world] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Virtuous Robe; wrist: Virtuous Bracers; hands: Raider Handwraps; waist: Virtuous Belt; legs: Virtuous Skirt; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 948, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

## Horde

### Band 20 (undead, 025004000000000000-00000000000000000-00000000000000000)

Set DPS (verified): 36.6. Weights run: 10.7s. Verify run: 28.8s. 136 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.006, intellect=2.689 ± 0.012, spirit=1.300 ± 0.007, mp5=2.938 ± 0.024, crit=0.146 ± 0.007 per rating point (14 rating = 1%, 2.043 per %), spell_haste=0.647 ± 0.084

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Goggles (4373) | Engineering [crafted] | 21.2 healing_power points (1.03 DPS) | yes | Pristine Circlet (253949, +0.00 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -0.78 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 3.9 healing_power points (0.19 DPS) | yes | Scholarly Pendant (277203, -0.06 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 24.2 healing_power points (1.18 DPS) | yes | Slime-encrusted Pads (6461, -0.75 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -0.80 DPS, sim-verified) [crafted] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | 11.6 healing_power points (0.56 DPS) | yes | Pearl-clasped Cloak (5542, -0.17 DPS) [crafted]; Seer's Cape (6378, -0.18 DPS) [dungeon]; Sanguine Cape (14376, -0.27 DPS, sim-verified) [world_drop] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 30.3 healing_power points (1.48 DPS) | yes | Corsair's Overshirt (5202, -0.44 DPS) [dungeon]; Seer's Robe (2981, -0.50 DPS) [world_drop]; Robe of the Moccasin (6465, -0.74 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 22.3 healing_power points (1.08 DPS) | yes | Bright Bracers (3647, -0.56 DPS) [world_drop]; Crystalline Cuffs (14148, -0.63 DPS) [dungeon]; Featherbead Bracers (15452, -0.64 DPS, sim-verified) [quest] |
| hands | Blight Gloves (279877) | The New Plague [quest] | 25.3 healing_power points (1.23 DPS) | yes | Pristine Gloves (253913, -0.30 DPS) [crafted]; Magefist Gloves (12977, -0.32 DPS) [world_drop]; Bright Gloves (3066, -0.46 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 21.8 healing_power points (1.06 DPS) | yes | Novice Ardent's Sash (253887, -0.18 DPS) [crafted]; Keller's Girdle (2911, -0.29 DPS, sim-verified) [world_drop]; Wise Man's Belt (4786, -0.40 DPS) [vendor] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 39.3 healing_power points (1.91 DPS) | yes | Abomination Skin Leggings (23173, -0.87 DPS) [dungeon]; Filigreed Silky Leggings (253939, -0.88 DPS) [crafted]; Darkweave Breeches (12987, -1.17 DPS, sim-verified) [world_drop] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 17.1 healing_power points (0.83 DPS) | yes | Spidersilk Boots (4320, -0.31 DPS) [crafted]; Bluegill Sandals (1560, -0.32 DPS) [world]; Sanguine Sandals (14374, -0.52 DPS, sim-verified) [world_drop] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 16.1 healing_power points (0.79 DPS) | yes | Volcanic Rock Ring (12053, -0.39 DPS) [world_drop]; Band of Purification (12996, -0.41 DPS) [world_drop]; Deep Fathom Ring (6463, -0.47 DPS) [dungeon] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 13.4 healing_power points (0.65 DPS) | yes | Band of Purification (12996, -0.27 DPS) [world_drop]; Deep Fathom Ring (6463, -0.34 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.36 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 26.9 healing_power points (1.31 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Staff of Orgrimmar (15444, -0.18 DPS) [quest]; Advisor's Gnarled Staff (20425, -0.36 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Decay (5252) (or Flaring Baton (5326)) | Beren's Peril [quest] | 5.4 healing_power points (0.26 DPS) | yes | Flaring Baton (5326, +0.00 DPS) [quest]; Wisesight Wand (286750, -0.13 DPS) [world] |

**New at 20:** head: Shadow Goggles; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Blight Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Gnarled Necromancer's Staff; ranged: Wand of Decay

No-known-source sample (15 of 136, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 209613 Insignia of the Alliance; 241089 Scarlet Dagger; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff; 263007 Skyseer's Vest; 263402 Sash of Sorrow

### Band 30 (undead, 025005031320000000-00000000000000000-00000000000000000)

Set DPS (verified): 67.1. Weights run: 11.0s. Verify run: 36.6s. 244 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.016, intellect=2.126 ± 0.014, spirit=1.720 ± 0.011, mp5=3.816 ± 0.013, crit=0.212 ± 0.009 per rating point (14 rating = 1%, 2.965 per %), spell_haste=not significant (-0.112 ± 0.086)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 43.3 healing_power points (2.74 DPS) | yes | Nightsky Cowl (4039, -0.58 DPS) [world_drop]; Resilient Cap (14401, -0.82 DPS) [world_drop]; Shadow Hood (4323, -1.26 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.4 healing_power points (0.97 DPS) | yes | Crystal Starfire Medallion (5003, -0.11 DPS) [world_drop]; Scorn's Icy Choker (23169, -0.17 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -0.17 DPS) [world] |
| shoulder | Nightsky Mantle (4718) | World drop [world_drop] | 25.6 healing_power points (1.62 DPS) | yes | Ghostly Mantle (3324, -0.07 DPS) [quest]; Death Speaker Mantle (6685, -0.14 DPS) [dungeon]; Talbar Mantle (10657, -0.33 DPS) [quest] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | 22.2 healing_power points (1.40 DPS) | yes | Glowing Thresher Cape (6901, -0.09 DPS) [dungeon]; Battle Healer's Cloak (19529, -0.14 DPS) [rep]; Resilient Cape (14400, -0.59 DPS) [world_drop] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 43.4 healing_power points (2.74 DPS) | yes | Pristine Gown (253961, -0.10 DPS) [crafted]; Beguiler Robes (7728, -0.26 DPS) [dungeon]; Silver-thread Robe (4035, -0.83 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 22.1 healing_power points (1.40 DPS) | yes | Nightsky Wristbands (6407, -0.26 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.32 DPS) [world_drop]; Spidertank Oilrag (9448, -0.67 DPS) [dungeon] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 31.3 healing_power points (1.98 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.36 DPS) [dungeon]; Blight Gloves (279877, -0.49 DPS) [quest]; Truefaith Gloves (7049, -0.62 DPS) [crafted] |
| waist | Lilac Sash (6780) | Centaur Bounty [quest] | 24.3 healing_power points (1.54 DPS) | yes | Resilient Cord (14406, -0.29 DPS) [world_drop]; Pristine Sash (253925, -0.30 DPS) [crafted]; Dreamer's Belt (4829, -0.38 DPS) [vendor] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 45.5 healing_power points (2.87 DPS) | yes | Filigreed Pristine Leggings (253937, -0.50 DPS) [crafted]; Earthen Leggings (253999, -0.73 DPS) [crafted]; Sacred Burial Trousers (6282, -1.22 DPS) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 41.8 healing_power points (2.64 DPS) | yes | Acidic Walkers (9454, -1.13 DPS) [dungeon]; Soggy Boots (274747, -1.23 DPS) [vendor]; Frothing Slippers (254003, -1.26 DPS) [crafted] |
| finger1 | The Queen's Jewel (13094) | World drop [world_drop] | sim-verified (67.1 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; Sea Giant's Toe Ring (274746, -0.00 DPS) [vendor]; Solid Gold Band (279845, -0.15 DPS) [quest] |
| finger2 | Electrocutioner Lagnut (9447) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (67.1 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; Sea Giant's Toe Ring (274746, -0.00 DPS) [vendor]; Solid Gold Band (279845, -0.15 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, +0.00 DPS) [dungeon]; Tendonscraper Dagger (273053, +0.00 DPS) [world]; Advisor's Gnarled Staff (19569, -0.37 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 16.1 healing_power points (1.02 DPS) | yes | Summoner's Wand (5245, -0.48 DPS) [world]; Lesser Mystic Wand (11289, -0.48 DPS) [crafted]; Starfaller (13063, -0.48 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Nightsky Mantle; back: Darkspear Raider's Cloak; chest: Death Speaker Robes; hands: Gloves of Old; waist: Lilac Sash; legs: Pristine Leggings; feet: Gilded Slippers; finger1: The Queen's Jewel; finger2: Electrocutioner Lagnut; main_hand: Spine of the Devourer; ranged: Wand of Mana Concentration

No-known-source sample (15 of 244, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (undead, 025005031330011520-00000000000000000-00000000000000000)

Set DPS (verified): 132.9. Weights run: 12.6s. Verify run: 80.7s. 323 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.021, intellect=2.089 ± 0.021, spirit=0.812 ± 0.036, mp5=2.297 ± 0.038, crit=0.411 ± 0.013 per rating point (14 rating = 1%, 5.761 per %), spell_haste=-10.189 ± 0.388

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 56.8 healing_power points (4.35 DPS) | yes | Whitemane's Chapeau (7720, -1.08 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -1.45 DPS) [world_drop]; Thinking Cap (2624, -1.63 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 29.7 healing_power points (2.27 DPS) | yes | Necklace of Calisea (1714, -0.36 DPS, sim-verified) [dungeon]; Darkspear Warding Pendant (272074, -1.15 DPS) [vendor]; Archmagister's Faceted Pendant (273049, -1.31 DPS) [world] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 42.2 healing_power points (3.23 DPS) | yes | Mistscape Mantle (4734, -1.16 DPS) [dungeon]; Silksand Shoulder Pads (14423, -1.28 DPS) [world_drop]; Windchaser Amice (14432, -1.50 DPS, sim-verified) [world_drop] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 26.2 healing_power points (2.01 DPS) | yes | Battle Healer's Cloak (19528, -0.19 DPS, sim-verified) [rep]; Blackforge Cape (6424, -0.48 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 64.5 healing_power points (4.94 DPS) | yes | Pristine Gown (253961, -2.04 DPS) [crafted]; Red Mageweave Vest (10007, -2.06 DPS) [crafted]; Death Speaker Robes (6682, -5.31 DPS, sim-verified) [dungeon] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 18.8 healing_power points (1.44 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Mistscape Bracers (4045, -0.16 DPS) [dungeon]; Enchanted Stonecloth Bracers (4979, -0.16 DPS) [quest] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (132.9 DPS) | yes | Earthen Silk Gloves (254017, -0.89 DPS) [crafted]; Truefaith Gloves (7049, -1.13 DPS) [crafted]; Gilded Handwraps (254021, -1.65 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 43.6 healing_power points (3.34 DPS) | yes | Sutarn's Ring (13105, -1.36 DPS) [world_drop]; Razzeric's Customized Seatbelt (6726, -1.42 DPS) [quest]; Deathmage Sash (10771, -2.01 DPS, sim-verified) [dungeon] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-verified (132.9 DPS) | yes | Stoneweaver Leggings (9407, +0.00 DPS) [dungeon]; Filigreed Pristine Leggings (253937, +0.00 DPS) [crafted]; Pristine Leggings (253987, -1.72 DPS, sim-verified) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 37.9 healing_power points (2.90 DPS) | yes | Furen's Boots (13100, -0.78 DPS) [world_drop]; Boots of the Maharishi (9658, -1.09 DPS) [quest]; Kodo Rustler Boots (15697, -1.31 DPS) [quest] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.38 DPS) | yes | Ogremind Ring (1993, -0.07 DPS) [world_drop]; Mindbender Loop (5009, -0.13 DPS) [world_drop]; Welken Ring (5011, -0.14 DPS) [world_drop] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 17.1 healing_power points (1.31 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Mindbender Loop (5009, -0.06 DPS) [world_drop]; Welken Ring (5011, -0.07 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (132.9 DPS) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (132.9 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Death Speaker Scepter (2816, -2.49 DPS) [dungeon]; Hypnotic Blade (7714, -2.86 DPS) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 17.9 healing_power points (1.37 DPS) | yes | Beacon of Hope (9393, -0.05 DPS) [dungeon]; Witch's Finger (16887, -0.09 DPS) [quest]; Aurora Sphere (7610, -0.36 DPS, sim-verified) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 15.0 healing_power points (1.15 DPS) | yes | Wand of Mana Concentration (273043, -0.21 DPS, sim-verified) [world]; Flash Wand (5248, -0.32 DPS) [quest]; Goblin Igniter (5253, -0.32 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Stormcloth Gloves; waist: Gilded Cord; legs: Stormcloth Pants; finger1: Sea Giant's Toe Ring; finger2: Voodoo Band; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Orb of Souls; ranged: Jaina's Firestarter

No-known-source sample (15 of 323, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16336 Sergeant Major's Cape

### Band 50 (undead, 025005031330011520-03502000000000000-00000000000000000)

Set DPS (verified): 211.5. Weights run: 12.5s. Verify run: 40.8s. 414 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.039, intellect=2.560 ± 0.032, spirit=1.896 ± 0.047, mp5=4.383 ± 0.068, crit=0.615 ± 0.019 per rating point (14 rating = 1%, 8.612 per %), spell_haste=-11.402 ± 0.547

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 86.6 healing_power points (6.60 DPS) | yes | Blood Guard's Satin Cover (220899, -0.19 DPS) [vendor]; Papal Fez (9431, -0.35 DPS, sim-verified) [dungeon]; Imperial Red Circlet (8254, -1.31 DPS) [world_drop] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 43.4 healing_power points (3.31 DPS) | yes | Glowing Eye of Mordresh (10769, -0.76 DPS) [dungeon]; Necklace of Calisea (1714, -0.93 DPS) [dungeon]; Darkmoon Necklace (19303, -1.00 DPS, sim-verified) [vendor] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | 57.9 healing_power points (4.40 DPS) | yes | Blood Guard's Satin Pads (220901, +0.00 DPS) [vendor]; Rotgrip Mantle (17732, -0.90 DPS) [dungeon]; Inquisitor's Shawl (19507, -2.68 DPS, sim-verified) [dungeon] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 45.3 healing_power points (3.45 DPS) | yes | Battle Healer's Cloak (19527, -0.91 DPS) [rep]; Imperial Red Cloak (8248, -1.01 DPS, sim-verified) [world_drop]; Highborne Cloak (14450, -1.12 DPS) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 73.2 healing_power points (5.57 DPS) | yes | Stone Guard's Satin Armor (220903, -0.02 DPS) [vendor]; Ghostweave Vest (14141, -0.27 DPS, sim-verified) [crafted]; Earth Warder's Vest (21311, -1.03 DPS) [quest] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 47.4 healing_power points (3.61 DPS) | yes | Forgotten Wraps (9433, -1.27 DPS) [world_drop]; Bloodband Bracers (11469, -1.42 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -3.39 DPS, sim-verified) [quest] |
| hands | Virtuous Mitts (226950) | Mokvar [vendor] | 86.3 healing_power points (6.57 DPS) | yes | Gilded Gloves (254095, -1.26 DPS) [crafted]; Virtuous Hands (226958, -2.30 DPS) [vendor]; Raider Handwraps (272098, -3.67 DPS, sim-verified) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 60.6 healing_power points (4.62 DPS) | yes | Gilded Waistcord (254081, -0.65 DPS) [crafted]; Earthenweave Cord (254077, -1.05 DPS) [crafted]; Gilded Cord (254037, -1.12 DPS, sim-verified) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 86.2 healing_power points (6.56 DPS) | yes | Senior Designer's Pantaloons (11841, -0.70 DPS) [dungeon]; Stone Guard's Satin Leggings (220902, -1.02 DPS) [vendor]; Dalewind Trousers (13008, -1.20 DPS) [world_drop] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 54.0 healing_power points (4.11 DPS) | yes | First Sergeant's Satin Boots (220900, -0.25 DPS) [vendor]; Furen's Boots (13100, -0.44 DPS) [world_drop]; Vinerot Sandals (17748, -1.56 DPS, sim-verified) [dungeon] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 38.1 healing_power points (2.90 DPS) | yes | Welken Ring (5011, -0.92 DPS) [world_drop]; Coldwater Ring (4550, -1.02 DPS) [quest]; Voodoo Band (1996, -1.10 DPS) [world] |
| finger2 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 30.7 healing_power points (2.34 DPS) | yes | Welken Ring (5011, -0.35 DPS) [world_drop]; Coldwater Ring (4550, -0.45 DPS) [quest]; Voodoo Band (1996, -0.54 DPS) [world] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (211.5 DPS) | yes | Evonice's Landin' Pilla (18951, -3.56 DPS) [quest]; Uther's Strength (11302, -3.63 DPS) [world_drop]; Alchemist's Stone (13503, -4.43 DPS) [crafted] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (211.5 DPS) | yes | Evonice's Landin' Pilla (18951, -0.29 DPS) [quest]; Uther's Strength (11302, -0.94 DPS, sim-verified) [world_drop]; Alchemist's Stone (13503, -1.15 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (211.5 DPS) | yes | Charstone Dirk (17710, -6.73 DPS) [dungeon]; Spellshifter Rod (9527, -9.78 DPS) [quest]; Hammer of the Northern Wind (810, -25.49 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 21.0 healing_power points (1.60 DPS) | yes | Wand of Mana Concentration (273043, -0.38 DPS, sim-verified) [world]; Flash Wand (5248, -0.39 DPS) [quest]; Goblin Igniter (5253, -0.39 DPS) [quest] |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Nethergeld Shoulders; back: Darkspear Raider's Cloak; wrist: Nethergeld Cuffs; hands: Virtuous Mitts; waist: Dawnspire Cord; legs: Kilt of the Atal'ai Prophet; feet: Gilded Sandals; finger1: Eye of Adaegus; finger2: Darkspear Signet; trinket1: Darkspear Voodoo Seal; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 414, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60 (undead, 025005031330011520-03505003030100000-00000000000000000)

Set DPS (verified): 352.2. Weights run: 12.2s. Verify run: 49.6s. 940 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.138, intellect=2.726 ± 0.073, spirit=3.896 ± 0.436, mp5=3.200 ± 0.193, crit=0.872 ± 0.038 per rating point (14 rating = 1%, 12.211 per %), spell_haste=-11.957 ± 1.227

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 153.9 healing_power points (12.49 DPS) | yes | Warlord's Satin Hood (231635, -0.12 DPS) [vendor]; Champion's Satin Hood (227118, -0.83 DPS) [pvp]; Devout Crown (16693, -11.55 DPS, sim-verified) [dungeon] |
| neck | The Eye of Zuldazar (19593) (or The All-Seeing Eye of Zuldazar (19594)) | The Eye of Zuldazar [quest] | 90.4 healing_power points (7.34 DPS) | yes | The All-Seeing Eye of Zuldazar (19594, +0.00 DPS) [quest]; Wavefront Necklace (20685, -1.99 DPS) [world]; Drake Tooth Necklace (21531, -2.58 DPS) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 117.2 healing_power points (9.51 DPS) | yes | Virtuous Mantle (226951, +0.00 DPS, sim-verified) [quest]; Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Devout Mantle (16695, -1.12 DPS) [dungeon] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 69.3 healing_power points (5.62 DPS) | yes | Darkspear Raider's Cloak (272076, -0.94 DPS) [vendor]; Battle Healer's Cloak (19526, -0.98 DPS) [rep] |
| chest | Virtuous Robe (226945) | Saving the Best for Last [quest] | sim-verified (+10.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Devout Robe (16690, -0.60 DPS) [dungeon]; Mooncloth Vest (14138, -10.52 DPS, sim-verified) [crafted] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-verified (+7.7 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Satin Bracers (17619, +0.00 DPS) [pvp]; Devout Bracers (16697, -1.11 DPS) [dungeon]; Bracers of Hope (22667, -7.73 DPS, sim-verified) [quest] |
| hands | Virtuous Mitts (226950) | Mokvar [vendor] | 124.1 healing_power points (10.07 DPS) | yes | Raider Handwraps (272097, -2.15 DPS) [vendor]; Swarmtender's Gloves (275607, -2.45 DPS) [crafted]; Devout Gloves (16692, -11.02 DPS, sim-verified) [dungeon] |
| waist | Virtuous Belt (226948) | Just Compensation [quest] | sim-verified (+8.8 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Satin Cinch (17621, +0.00 DPS) [pvp]; Wisdom of the Timbermaw (19047, -0.35 DPS) [crafted]; Penitent's Cinch (272394, -8.82 DPS, sim-verified) [vendor] |
| legs | Virtuous Skirt (226946) | Anthion's Parting Words [quest] | sim-verified (+8.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Legionnaire's Satin Legguards (227123, -0.76 DPS) [pvp]; General's Satin Legguards (231634, -1.09 DPS) [vendor]; Devout Skirt (16694, -8.47 DPS, sim-verified) [dungeon] |
| feet | Virtuous Sandals (226952) | Anthion's Parting Words [quest] | 139.7 healing_power points (11.34 DPS) | yes | Mooncloth Boots (15802, -2.19 DPS) [crafted]; Devout Sandals (16691, -2.86 DPS) [dungeon]; Incandescent Mooncloth Boots (227862, -8.70 DPS, sim-verified) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 75.5 healing_power points (6.13 DPS) | yes | Channeler's Ring (272406, -0.50 DPS) [vendor]; Eye of Adaegus (5266, -1.01 DPS) [world_drop]; Blessed Band of Light (272407, -1.09 DPS) [vendor] |
| finger2 | Band of the Hierophant (13096) | World drop [world_drop] | 70.1 healing_power points (5.69 DPS) | yes | Channeler's Ring (272406, +0.00 DPS, sim-verified) [vendor]; Eye of Adaegus (5266, -0.57 DPS) [world_drop]; Blessed Band of Light (272407, -0.64 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+13.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, -0.10 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.73 DPS) [quest]; Shard of the Splithooves (10659, -2.34 DPS) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.25 DPS) [quest]; Uther's Strength (11302, -4.30 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Hale Magefire (13000, -8.95 DPS) [world_drop]; Charstone Dirk (17710, -9.07 DPS) [dungeon]; Hand of Edward the Odd (2243, -15.83 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 59.9 healing_power points (4.86 DPS) | yes | Chillnail Splinter (10704, -2.96 DPS) [quest]; Goblin Igniter (5253, -3.02 DPS) [quest]; Jaina's Firestarter (13064, -4.72 DPS, sim-verified) [world_drop] |

**New at 60:** head: Virtuous Crown; neck: The Eye of Zuldazar; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Virtuous Robe; wrist: Virtuous Bracers; waist: Virtuous Belt; legs: Virtuous Skirt; feet: Virtuous Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of the Hierophant; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 940, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60, raid preset (undead, 025005031330011520-03505003030100000-00000000000000000)

Set DPS (verified): 637.6. Weights run: 8.2s. Verify run: 25.2s. 940 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.152, intellect=1.369 ± 0.031, spirit=1.520 ± 0.029, mp5=2.683 ± 0.035, crit=0.708 ± 0.025 per rating point (14 rating = 1%, 9.909 per %), spell_haste=-6.464 ± 0.590

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 90.5 healing_power points (18.22 DPS) | yes | Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Champion's Satin Hood (227118, -0.92 DPS) [pvp]; Devout Crown (16693, -8.72 DPS, sim-verified) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 53.7 healing_power points (10.81 DPS) | yes | The All-Seeing Eye of Zuldazar (19594, -1.00 DPS) [quest]; The Eye of Zuldazar (19593, -2.17 DPS, sim-verified) [quest]; Drake Tooth Necklace (21531, -2.26 DPS) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 90.3 healing_power points (18.18 DPS) | yes | Virtuous Mantle (226951, +0.00 DPS, sim-verified) [quest]; Champion's Satin Mantle (227120, -4.35 DPS) [pvp]; Warlord's Satin Mantle (231631, -6.63 DPS) [pvp] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 55.7 healing_power points (11.21 DPS) | yes | Battle Healer's Cloak (19526, -3.53 DPS) [rep]; Drape of Recovery (272413, -13.74 DPS, sim-verified) [vendor] |
| chest | Virtuous Robe (226945) | Saving the Best for Last [quest] | sim-verified (+8.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Mooncloth Vest (14138, -1.88 DPS) [crafted]; Truefaith Vestments (14154, -8.87 DPS, sim-verified) [crafted] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-verified (+8.6 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Satin Bracers (17619, -1.17 DPS) [pvp]; Earthenweave Cuffs (254131, -1.98 DPS) [crafted]; Bracers of Hope (22667, -8.59 DPS, sim-verified) [quest] |
| hands | Virtuous Mitts (226950) | Mokvar [vendor] | 66.4 healing_power points (13.37 DPS) | yes | Raider Handwraps (272097, -1.10 DPS) [vendor]; Mooncloth Gloves (18409, -2.32 DPS) [crafted]; Desert Bloom Gloves (20717, -5.16 DPS, sim-verified) [quest] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 63.9 healing_power points (12.85 DPS) | yes | Virtuous Belt (226948, +0.00 DPS, sim-verified) [quest]; Penitent's Cinch (272394, -0.88 DPS) [vendor]; General's Satin Cinch (17621, -1.92 DPS) [pvp] |
| legs | Virtuous Skirt (226946) | Anthion's Parting Words [quest] | 80.3 healing_power points (16.16 DPS) | yes | Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Devout Skirt (16694, -5.33 DPS, sim-verified) [dungeon] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 84.5 healing_power points (17.00 DPS) | yes | Virtuous Sandals (226952, -2.49 DPS, sim-verified) [quest]; Mooncloth Boots (15802, -4.33 DPS) [crafted]; General's Satin Walkers (231630, -6.20 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 60.7 healing_power points (12.23 DPS) | yes | Blessed Band of Light (272407, -3.54 DPS) [vendor]; Channeler's Ring (272406, -5.92 DPS) [vendor]; Band of the Hierophant (13096, -6.10 DPS) [world_drop] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 46.4 healing_power points (9.35 DPS) | yes | Blessed Band of Light (272407, -0.88 DPS, sim-verified) [vendor]; Channeler's Ring (272406, -3.04 DPS) [vendor]; Band of the Hierophant (13096, -3.22 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+11.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, -4.43 DPS) [world_drop]; Uther's Strength (11302, -4.48 DPS) [world_drop]; Shard of the Splithooves (10659, -4.86 DPS) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -4.78 DPS) [world_drop]; Shard of the Splithooves (10659, -5.21 DPS) [quest]; Uther's Strength (11302, -6.09 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -13.48 DPS) [dungeon]; Hand of Righteousness (7721, -19.57 DPS) [dungeon]; Hand of Edward the Odd (2243, -21.40 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 33.7 healing_power points (6.79 DPS) | yes | Wand of Mana Concentration (273043, -4.61 DPS) [world]; Captain Rackmore's Tiller (16789, -4.66 DPS) [quest]; Jaina's Firestarter (13064, -5.19 DPS, sim-verified) [world_drop] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Virtuous Robe; wrist: Virtuous Bracers; waist: Wisdom of the Timbermaw; legs: Virtuous Skirt; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 940, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

