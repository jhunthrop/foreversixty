# Leveling BiS: Discipline

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 025003010000000000-00000000000000000-000000000000000000)

Set DPS (verified): 35.9. Weights run: 1.7s. Verify run: 5.2s. 148 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.006, intellect=2.353 ± 0.013, spirit=1.366 ± 0.007, mp5=3.060 ± 0.046, crit=0.161 ± 0.010 per rating point (14 rating = 1%, 2.260 per %), spell_haste=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Goggles (4373) | Engineering [crafted] | 20.0 healing_power points (0.94 DPS) | yes | Pristine Circlet (253949, +0.00 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -0.68 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 5.5 healing_power points (0.26 DPS) | yes | Scholarly Pendant (277203, +0.00 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 21.2 healing_power points (1.00 DPS) | yes | Slime-encrusted Pads (6461, -0.56 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -0.65 DPS, sim-verified) [crafted] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | 11.7 healing_power points (0.55 DPS) | yes | Seer's Cape (6378, -0.20 DPS) [dungeon]; Pearl-clasped Cloak (5542, -0.22 DPS) [crafted]; Sanguine Cape (14376, -0.35 DPS, sim-verified) [world_drop] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 28.9 healing_power points (1.36 DPS) | yes | Corsair's Overshirt (5202, -0.33 DPS) [dungeon]; Seer's Robe (2981, -0.50 DPS) [world_drop]; Robe of the Moccasin (6465, -0.79 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 20.9 healing_power points (0.99 DPS) | yes | Repurposed Hair Band (281256, -0.64 DPS) [quest]; Seer's Cuffs (3645, -0.75 DPS) [dungeon]; Bright Bracers (3647, -0.76 DPS, sim-verified) [world_drop] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 18.1 healing_power points (0.85 DPS) | yes | Bright Gloves (3066, -0.15 DPS) [world_drop]; Tomb Robber's Gloves (280096, -0.19 DPS) [quest]; Magefist Gloves (12977, -0.25 DPS, sim-verified) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 20.4 healing_power points (0.96 DPS) | yes | Novice Ardent's Sash (253887, -0.15 DPS) [crafted]; Keller's Girdle (2911, -0.28 DPS, sim-verified) [world_drop]; Tarantula Silk Sash (3229, -0.28 DPS) [world] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 37.6 healing_power points (1.77 DPS) | yes | Filigreed Silky Leggings (253939, -0.85 DPS) [crafted]; Filigreed Shadow Leggings (253943, -0.85 DPS) [crafted]; Darkweave Breeches (12987, -1.08 DPS, sim-verified) [world_drop] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 16.1 healing_power points (0.76 DPS) | yes | Bluegill Sandals (1560, -0.28 DPS) [world]; Sanguine Sandals (14374, -0.31 DPS) [world_drop]; Kimbra Boots (6191, -0.42 DPS, sim-verified) [quest] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 14.1 healing_power points (0.66 DPS) | yes | Band of Purification (12996, -0.28 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.33 DPS) [world_drop]; Deep Fathom Ring (6463, -0.34 DPS) [dungeon] |
| finger2 | Black Pearl Ring (6332) | Lady Vespira [world] | 12.9 healing_power points (0.61 DPS) | yes | Band of Purification (12996, -0.25 DPS, sim-verified) [world_drop]; Volcanic Rock Ring (12053, -0.28 DPS) [world_drop]; Deep Fathom Ring (6463, -0.29 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 23.5 healing_power points (1.11 DPS) | yes | Channeler's Staff (4437, -0.22 DPS) [world]; Lesser Staff of the Spire (1300, -0.44 DPS) [world]; Staff of Westfall (2042, -0.48 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Flaring Baton (5326) (or Moonstone Wand (15204)) | The Escape [quest] | 4.7 healing_power points (0.22 DPS) | yes | Moonstone Wand (15204, +0.00 DPS) [quest]; Sable Wand (7607, -0.03 DPS) [quest]; Dwarven Flamestick (5241, -0.09 DPS) [quest] |

**New at 20:** head: Shadow Goggles; neck: Tarnished Locket; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Black Pearl Ring; main_hand: Twisted Chanter's Staff; ranged: Flaring Baton

No-known-source sample (15 of 148, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes; 16607 Acolyte's Sacrificial Robes; 18851 Insignia of the Horde; 20434 Lorekeeper's Staff

### Band 30 (human, 025003031304000000-00000000000000000-000000000000000000)

Set DPS (verified): 78.6. Weights run: 1.6s. Verify run: 5.1s. 256 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.017, intellect=2.196 ± 0.016, spirit=1.734 ± 0.014, mp5=3.806 ± 0.015, crit=0.214 ± 0.012 per rating point (14 rating = 1%, 2.992 per %), spell_haste=not significant (0.275 ± 0.123)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 43.4 healing_power points (2.97 DPS) | yes | Nightsky Cowl (4039, -0.57 DPS) [world_drop]; Resilient Cap (14401, -0.84 DPS) [world_drop]; Embalmed Shroud (7691, -1.00 DPS, sim-verified) [dungeon] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.7 healing_power points (1.08 DPS) | yes | Scorn's Icy Choker (23169, -0.17 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -0.17 DPS) [world]; Crystal Starfire Medallion (5003, -0.20 DPS, sim-verified) [world_drop] |
| shoulder | Mantle of Honor (3560) | Bride of the Embalmer [quest] | 27.5 healing_power points (1.88 DPS) | yes | Nightsky Mantle (4718, -0.09 DPS) [world_drop]; Death Speaker Mantle (6685, -0.23 DPS) [dungeon]; Batwing Mantle (6697, -0.23 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | 22.8 healing_power points (1.56 DPS) | yes | Prelacy Cape (7004, -0.06 DPS, sim-verified) [quest]; Glowing Thresher Cape (6901, -0.13 DPS) [dungeon]; Caretaker's Cape (19533, -0.19 DPS) [rep] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 44.2 healing_power points (3.02 DPS) | yes | Pristine Gown (253961, -0.17 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -0.27 DPS) [dungeon]; Silver-thread Robe (4035, -0.89 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 22.4 healing_power points (1.53 DPS) | yes | Glowing Magical Bracelets (13106, -0.33 DPS) [world_drop]; Nightsky Wristbands (6407, -0.42 DPS, sim-verified) [world_drop]; Spidertank Oilrag (9448, -0.75 DPS) [dungeon] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 31.8 healing_power points (2.18 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.38 DPS) [dungeon]; Truefaith Gloves (7049, -0.70 DPS) [crafted]; Zodiac Gloves (7106, -0.74 DPS) [quest] |
| waist | Resilient Cord (14406) | World drop [world_drop] | 20.1 healing_power points (1.38 DPS) | yes | Pristine Sash (253925, +0.00 DPS, sim-verified) [crafted]; Dreamer's Belt (4829, -0.09 DPS) [vendor]; Novice Ardent's Sash (253887, -0.16 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 46.0 healing_power points (3.15 DPS) | yes | Earthen Leggings (253999, -0.82 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.82 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -1.39 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 42.3 healing_power points (2.90 DPS) | yes | Soggy Boots (274747, -1.34 DPS) [vendor]; Frothing Slippers (254003, -1.37 DPS) [crafted]; Acidic Walkers (9454, -1.88 DPS, sim-verified) [dungeon] |
| finger1 | The Queen's Jewel (13094) | World drop [world_drop] | 18.3 healing_power points (1.25 DPS) | yes | Runemender's Seal (273031, -0.01 DPS) [world]; Sea Giant's Toe Ring (274746, -0.02 DPS) [vendor]; Black Widow Band (6199, -0.20 DPS) [world] |
| finger2 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 18.1 healing_power points (1.24 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; Sea Giant's Toe Ring (274746, -0.01 DPS) [vendor]; Black Widow Band (6199, -0.19 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | 98.7 healing_power points (6.76 DPS) | yes | Royal Diplomatic Scepter (9457, -0.72 DPS, sim-verified) [dungeon]; Death Speaker Scepter (2816, -2.10 DPS) [dungeon]; Spine of the Devourer (273040, -4.29 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 16.4 healing_power points (1.12 DPS) | yes | Lesser Mystic Wand (11289, -0.52 DPS) [crafted]; Starfaller (13063, -0.52 DPS) [world_drop]; Summoner's Wand (5245, -0.76 DPS, sim-verified) [world] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Mantle of Honor; back: Darkspear Raider's Cloak; chest: Death Speaker Robes; hands: Gloves of Old; waist: Resilient Cord; legs: Pristine Leggings; feet: Gilded Slippers; finger1: The Queen's Jewel; finger2: Electrocutioner Lagnut; main_hand: Wind Spirit Staff; ranged: Wand of Mana Concentration

No-known-source sample (15 of 256, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 40 (human, 025003031305101520-00000000000000000-000000000000000000)

Set DPS (verified): 137.1. Weights run: 1.8s. Verify run: 14.1s. 336 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.023, intellect=2.189 ± 0.028, spirit=0.629 ± 0.042, mp5=2.446 ± 0.048, crit=0.446 ± 0.019 per rating point (14 rating = 1%, 6.248 per %), spell_haste=-4.866 ± 0.389

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 56.9 healing_power points (4.25 DPS) | yes | Holy Shroud (2721, -1.50 DPS) [world_drop]; Craftsman's Monocle (4393, -1.80 DPS) [crafted]; Thinking Cap (2624, -3.39 DPS, sim-verified) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 30.5 healing_power points (2.28 DPS) | yes | Necklace of Calisea (1714, -0.23 DPS, sim-verified) [dungeon]; Darkspear Warding Pendant (272074, -1.13 DPS) [vendor]; Archmagister's Faceted Pendant (273049, -1.30 DPS) [world] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 43.5 healing_power points (3.24 DPS) | yes | Mistscape Mantle (4734, -1.21 DPS) [dungeon]; Silksand Shoulder Pads (14423, -1.31 DPS) [world_drop]; Windchaser Amice (14432, -1.55 DPS, sim-verified) [world_drop] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 26.6 healing_power points (1.99 DPS) | yes | Caretaker's Cape (19532, -0.27 DPS, sim-verified) [rep]; Blackforge Cape (6424, -0.49 DPS) [dungeon]; Cloak of Rot (4462, -0.68 DPS) [world] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 63.0 healing_power points (4.71 DPS) | yes | Red Mageweave Vest (10007, -1.76 DPS) [crafted]; Silksand Tunic (14417, -1.86 DPS) [world_drop]; Death Speaker Robes (6682, -5.16 DPS, sim-verified) [dungeon] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 19.7 healing_power points (1.47 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Mistscape Bracers (4045, -0.16 DPS) [dungeon]; Enchanted Stonecloth Bracers (4979, -0.16 DPS) [quest] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (137.1 DPS) | yes | Earthen Silk Gloves (254017, -1.06 DPS) [crafted]; Red Mageweave Gloves (10018, -1.15 DPS) [crafted]; Gilded Handwraps (254021, -2.07 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 43.3 healing_power points (3.23 DPS) | yes | Razzeric's Customized Seatbelt (6726, -1.27 DPS) [quest]; Teacher's Sash (10747, -1.27 DPS) [quest]; Deathmage Sash (10771, -2.06 DPS, sim-verified) [dungeon] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-verified (137.1 DPS) | yes | Stoneweaver Leggings (9407, +0.00 DPS) [dungeon]; Filigreed Pristine Leggings (253937, +0.00 DPS) [crafted]; Pristine Leggings (253987, -1.59 DPS, sim-verified) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 37.8 healing_power points (2.83 DPS) | yes | Furen's Boots (13100, +0.00 DPS, sim-verified) [world_drop]; Kodo Rustler Boots (15697, -1.28 DPS) [quest]; Acidic Walkers (9454, -1.33 DPS) [dungeon] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.34 DPS) | yes | Ogremind Ring (1993, -0.06 DPS) [world_drop]; Mindbender Loop (5009, -0.11 DPS) [world_drop]; Welken Ring (5011, -0.20 DPS) [world_drop] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 17.2 healing_power points (1.28 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Mindbender Loop (5009, -0.05 DPS) [world_drop]; Welken Ring (5011, -0.14 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (137.1 DPS) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (137.1 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Wind Spirit Staff (6689, -1.34 DPS) [dungeon]; Royal Diplomatic Scepter (9457, -1.47 DPS) [dungeon] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 33.9 healing_power points (2.53 DPS) | yes | Aurora Sphere (7610, -0.92 DPS, sim-verified) [dungeon]; Orb of Souls (249395, -1.28 DPS) [crafted]; Beacon of Hope (9393, -1.30 DPS) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 15.0 healing_power points (1.12 DPS) | yes | Wand of Mana Concentration (273043, +0.00 DPS, sim-verified) [world]; Flash Wand (5248, -0.33 DPS) [quest]; Goblin Igniter (5253, -0.33 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Stormcloth Gloves; waist: Gilded Cord; legs: Stormcloth Pants; finger1: Sea Giant's Toe Ring; finger2: Voodoo Band; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Orb of Lorica; ranged: Jaina's Firestarter

No-known-source sample (15 of 336, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 50 (human, 025003031305101520-03502000000000000-000000000000000000)

Set DPS (verified): 197.4. Weights run: 1.8s. Verify run: 7.9s. 430 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.019, intellect=2.517 ± 0.048, spirit=1.688 ± 0.064, mp5=4.067 ± 0.108, crit=0.573 ± 0.024 per rating point (14 rating = 1%, 8.017 per %), spell_haste=-6.791 ± 0.586

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 82.2 healing_power points (6.59 DPS) | yes | Papal Fez (9431, -0.82 DPS) [dungeon]; Stormcloth Headband (10032, -1.29 DPS) [crafted]; Knight-Lieutenant's Satin Cover (220896, -3.20 DPS, sim-verified) [vendor] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 42.0 healing_power points (3.37 DPS) | yes | Glowing Eye of Mordresh (10769, -0.71 DPS) [dungeon]; Necklace of Calisea (1714, -1.01 DPS) [dungeon]; Darkmoon Necklace (19303, -4.17 DPS, sim-verified) [vendor] |
| shoulder | Knight-Lieutenant's Satin Pads (220894) | Captain Dirgehammer [vendor] | 63.4 healing_power points (5.09 DPS) | yes | Inquisitor's Shawl (19507, -1.26 DPS) [dungeon]; Rotgrip Mantle (17732, -1.45 DPS) [dungeon]; Nethergeld Shoulders (254049, -4.42 DPS, sim-verified) [crafted] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 43.7 healing_power points (3.50 DPS) | yes | Caretaker's Cape (19531, -0.93 DPS) [rep]; Highborne Cloak (14450, -1.15 DPS) [world_drop]; Imperial Red Cloak (8248, -3.08 DPS, sim-verified) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 71.5 healing_power points (5.74 DPS) | yes | Ghostweave Vest (14141, -0.42 DPS) [crafted]; Earth Warder's Vest (21311, -1.15 DPS) [quest]; Knight's Satin Armor (220892, -4.20 DPS, sim-verified) [vendor] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 46.1 healing_power points (3.70 DPS) | yes | Forgotten Wraps (9433, -1.27 DPS) [world_drop]; Bloodband Bracers (11469, -1.47 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -5.08 DPS, sim-verified) [quest] |
| hands | Virtuous Mitts (226950) | Mokvar [vendor] | 82.1 healing_power points (6.59 DPS) | yes | Gilded Gloves (254095, -1.16 DPS) [crafted]; Virtuous Hands (226958, -2.33 DPS) [vendor]; Raider Handwraps (272098, -5.63 DPS, sim-verified) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 59.8 healing_power points (4.80 DPS) | yes | Gilded Waistcord (254081, -0.66 DPS) [crafted]; Earthenweave Cord (254077, -1.22 DPS) [crafted]; Gilded Cord (254037, -4.28 DPS, sim-verified) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 81.7 healing_power points (6.55 DPS) | yes | Knight's Satin Leggings (220893, -0.83 DPS) [vendor]; Senior Designer's Pantaloons (11841, -1.30 DPS, sim-verified) [dungeon]; Dalewind Trousers (13008, -1.35 DPS) [world_drop] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 53.7 healing_power points (4.31 DPS) | yes | Sergeant Major's Satin Boots (220895, -0.28 DPS) [vendor]; Furen's Boots (13100, -0.72 DPS) [world_drop]; Vinerot Sandals (17748, -4.16 DPS, sim-verified) [dungeon] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 35.4 healing_power points (2.84 DPS) | yes | Welken Ring (5011, -0.88 DPS) [world_drop]; Coldwater Ring (4550, -1.01 DPS) [quest]; Voodoo Band (1996, -1.02 DPS) [world] |
| finger2 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 28.5 healing_power points (2.28 DPS) | yes | Welken Ring (5011, -0.33 DPS) [world_drop]; Coldwater Ring (4550, -0.46 DPS) [quest]; Voodoo Band (1996, -0.46 DPS) [world] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (197.4 DPS) | yes | Uther's Strength (11302, -3.45 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -3.54 DPS) [quest]; Thunderbrew's Boot Flask (744, -3.81 DPS) [quest] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (197.4 DPS) | yes | Evonice's Landin' Pilla (18951, -0.27 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.54 DPS) [quest]; Uther's Strength (11302, -3.67 DPS, sim-verified) [world_drop] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-verified (197.4 DPS) | yes | Hammer of the Northern Wind (810, +0.00 DPS) [world_drop]; Spellshifter Rod (9527, -3.20 DPS) [quest]; Hand of Righteousness (7721, -3.56 DPS) [dungeon] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 37.2 healing_power points (2.98 DPS) | yes | Skullspell Orb (10708, -0.56 DPS) [quest]; Aurora Sphere (7610, -1.03 DPS) [dungeon]; Twisting Essence Jar (249456, -3.85 DPS, sim-verified) [crafted] |
| ranged | Cairnstone Sliver (9654) | The Morrow Stone [quest] | 21.0 healing_power points (1.69 DPS) | yes | Wand of Mana Concentration (273043, -0.23 DPS) [world]; Goblin Igniter (5253, -0.47 DPS) [quest]; Jaina's Firestarter (13064, -3.07 DPS, sim-verified) [world_drop] |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Knight-Lieutenant's Satin Pads; back: Darkspear Raider's Cloak; wrist: Nethergeld Cuffs; hands: Virtuous Mitts; waist: Dawnspire Cord; legs: Kilt of the Atal'ai Prophet; feet: Gilded Sandals; finger1: Eye of Adaegus; finger2: Darkspear Signet; trinket1: Darkspear Voodoo Seal; main_hand: Charstone Dirk; ranged: Cairnstone Sliver

No-known-source sample (15 of 430, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60 (human, 025003031305101520-03505003030100000-000000000000000000)

Set DPS (verified): 320.7. Weights run: 1.8s. Verify run: 10.8s. 943 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.244, intellect=2.919 ± 0.088, spirit=4.482 ± 0.666, mp5=3.713 ± 0.283, crit=0.817 ± 0.053 per rating point (14 rating = 1%, 11.434 per %), spell_haste=not significant (-5.264 ± 1.542)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 165.8 healing_power points (12.51 DPS) | yes | Field Marshal's Satin Hood (231622, -0.20 DPS) [vendor]; Lieutenant Commander's Satin Hood (227121, -0.79 DPS) [pvp]; Devout Crown (16693, -13.38 DPS, sim-verified) [dungeon] |
| neck | The Eye of Zuldazar (19593) (or The All-Seeing Eye of Zuldazar (19594)) | The Eye of Zuldazar [quest] | 99.6 healing_power points (7.52 DPS) | yes | The All-Seeing Eye of Zuldazar (19594, +0.00 DPS) [quest]; Wavefront Necklace (20685, -2.14 DPS) [world]; Lady Alizabeth's Pendant (13002, -2.44 DPS) [world_drop] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 126.1 healing_power points (9.52 DPS) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Devout Mantle (16695, -1.02 DPS) [dungeon]; Virtuous Mantle (226951, -3.45 DPS, sim-verified) [quest] |
| back | Darkspear Raider's Cloak (272063) | Creeg Bothunk [vendor] | 73.6 healing_power points (5.55 DPS) | yes | Hide of the Wild (18510, +0.00 DPS, sim-verified) [crafted]; Ethereal Mist Cape (11873, -0.83 DPS) [quest]; Caretaker's Cape (19530, -0.89 DPS) [rep] |
| chest | Virtuous Robe (226945) | Saving the Best for Last [quest] | sim-verified (+6.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Field Marshal's Satin Tunic (231624, +0.00 DPS) [vendor]; Devout Robe (16690, -0.38 DPS) [dungeon]; Mooncloth Vest (14138, -6.72 DPS, sim-verified) [crafted] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-verified (+5.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Satin Bracers (17606, +0.00 DPS) [pvp]; Devout Bracers (16697, -1.09 DPS) [dungeon]; Bracers of Hope (22667, -5.51 DPS, sim-verified) [quest] |
| hands | Virtuous Mitts (226950) | Mokvar [vendor] | 136.8 healing_power points (10.32 DPS) | yes | Swarmtender's Gloves (275607, -2.35 DPS) [crafted]; Magister's Gloves (16684, -2.50 DPS) [dungeon]; Devout Gloves (16692, -13.34 DPS, sim-verified) [dungeon] |
| waist | Virtuous Belt (226948) | Just Compensation [quest] | sim-verified (+5.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Satin Sash (17609, +0.00 DPS) [pvp]; Wisdom of the Timbermaw (19047, -0.52 DPS) [crafted]; Penitent's Cinch (272394, -5.83 DPS, sim-verified) [vendor] |
| legs | Virtuous Skirt (226946) | Anthion's Parting Words [quest] | sim-verified (+4.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Captain's Satin Legguards (227125, -0.89 DPS) [pvp]; Marshal's Satin Legguards (231626, -1.34 DPS) [vendor]; Devout Skirt (16694, -4.57 DPS, sim-verified) [dungeon] |
| feet | Virtuous Sandals (226952) | Anthion's Parting Words [quest] | 155.0 healing_power points (11.70 DPS) | yes | Mooncloth Boots (15802, -2.41 DPS) [crafted]; Devout Sandals (16691, -2.92 DPS) [dungeon]; Incandescent Mooncloth Boots (227862, -9.35 DPS, sim-verified) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 79.8 healing_power points (6.03 DPS) | yes | Channeler's Ring (272406, -0.24 DPS) [vendor]; Eye of Adaegus (5266, -0.64 DPS) [world_drop]; Blessed Band of Light (272407, -0.98 DPS) [vendor] |
| finger2 | Band of the Hierophant (13096) | World drop [world_drop] | 78.5 healing_power points (5.92 DPS) | yes | Eye of Adaegus (5266, -0.54 DPS) [world_drop]; Blessed Band of Light (272407, -0.88 DPS) [vendor]; Channeler's Ring (272406, -1.22 DPS, sim-verified) [vendor] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor]; Uther's Strength (11302, -4.23 DPS, sim-verified) [world_drop] |
| trinket2 | Darkmoon Card: Blue Dragon (19288) | Darkmoon Beast Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor] |
| main_hand | Staff of Hale Magefire (13000) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -1.21 DPS) [dungeon]; Soulkeeper (1607, -1.73 DPS) [world_drop]; Hand of Edward the Odd (2243, -13.35 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 66.3 healing_power points (5.00 DPS) | yes | Jaina's Firestarter (13064, -2.67 DPS) [world_drop]; Goblin Igniter (5253, -3.11 DPS) [quest]; Cairnstone Sliver (9654, -6.60 DPS, sim-verified) [quest] |

**New at 60:** head: Virtuous Crown; neck: The Eye of Zuldazar; shoulder: Argent Elite Shoulders; back: Darkspear Raider's Cloak; chest: Virtuous Robe; wrist: Virtuous Bracers; waist: Virtuous Belt; legs: Virtuous Skirt; feet: Virtuous Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of the Hierophant; trinket1: Serenity Field; trinket2: Darkmoon Card: Blue Dragon; main_hand: Staff of Hale Magefire; ranged: Torch of Light

No-known-source sample (15 of 943, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60, raid preset (human, 025003031305101520-03505003030100000-000000000000000000)

Set DPS (verified): 575.2. Weights run: 1.5s. Verify run: 7.9s. 943 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.382, intellect=2.779 ± 0.085, spirit=2.149 ± 0.079, mp5=3.667 ± 0.090, crit=0.950 ± 0.061 per rating point (14 rating = 1%, 13.304 per %), spell_haste=not significant (-2.876 ± 1.662)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 127.9 healing_power points (16.17 DPS) | yes | Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Lieutenant Commander's Satin Hood (227121, -0.99 DPS) [pvp]; Gnomish Turban of Psychic Might (21517, -10.97 DPS, sim-verified) [quest] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 70.0 healing_power points (8.85 DPS) | yes | The All-Seeing Eye of Zuldazar (19594, -0.23 DPS) [quest]; Drake Tooth Necklace (21531, -1.35 DPS) [quest]; The Eye of Zuldazar (19593, -3.00 DPS, sim-verified) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 102.5 healing_power points (12.95 DPS) | yes | Lieutenant Commander's Satin Mantle (227119, -1.02 DPS) [pvp]; Field Marshal's Satin Mantle (231628, -2.14 DPS) [pvp]; Devout Mantle (16695, -10.67 DPS, sim-verified) [dungeon] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 69.8 healing_power points (8.82 DPS) | yes | Darkspear Raider's Cloak (272076, -2.55 DPS) [vendor]; Drape of Recovery (272413, -3.05 DPS) [vendor] |
| chest | Virtuous Robe (226945) | Saving the Best for Last [quest] | 127.0 healing_power points (16.05 DPS) | yes | Field Marshal's Satin Tunic (231624, +0.00 DPS) [vendor]; Devout Robe (16690, -1.02 DPS) [dungeon]; Mooncloth Vest (14138, -9.02 DPS, sim-verified) [crafted] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-verified (+11.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Satin Bracers (17606, +0.00 DPS) [pvp]; Devout Bracers (16697, -1.51 DPS) [dungeon]; Bracers of Hope (22667, -11.02 DPS, sim-verified) [quest] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 99.0 healing_power points (12.51 DPS) | yes | Virtuous Mitts (226950, +0.00 DPS, sim-verified) [vendor]; Devout Gloves (16692, -2.49 DPS) [dungeon]; Mooncloth Gloves (18409, -2.73 DPS) [crafted] |
| waist | Virtuous Belt (226948) | Just Compensation [quest] | 89.2 healing_power points (11.28 DPS) | yes | Marshal's Satin Sash (17609, +0.00 DPS) [pvp]; Devout Belt (16696, -0.42 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -6.09 DPS, sim-verified) [crafted] |
| legs | Virtuous Skirt (226946) | Anthion's Parting Words [quest] | sim-verified (+13.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Knight-Captain's Satin Legguards (227125, -0.30 DPS) [pvp]; Devout Skirt (16694, -13.07 DPS, sim-verified) [dungeon] |
| feet | Virtuous Sandals (226952) | Anthion's Parting Words [quest] | sim-verified (+6.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Mooncloth Boots (15802, -1.97 DPS) [crafted]; Marshal's Satin Walkers (231627, -3.46 DPS) [vendor]; Incandescent Mooncloth Boots (227862, -6.04 DPS, sim-verified) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 78.3 healing_power points (9.90 DPS) | yes | Channeler's Ring (272406, -2.81 DPS) [vendor]; Band of the Hierophant (13096, -3.40 DPS) [world_drop]; Blessed Band of Light (272407, -3.81 DPS) [vendor] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 64.5 healing_power points (8.15 DPS) | yes | Band of the Hierophant (13096, -1.65 DPS) [world_drop]; Blessed Band of Light (272407, -2.06 DPS) [vendor]; Channeler's Ring (272406, -5.86 DPS, sim-verified) [vendor] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (+12.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, -2.05 DPS) [world_drop]; Shard of the Splithooves (10659, -2.53 DPS) [quest]; Evonice's Landin' Pilla (18951, -2.59 DPS) [quest] |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -0.91 DPS) [world_drop]; Shard of the Splithooves (10659, -0.96 DPS, sim-verified) [quest]; Evonice's Landin' Pilla (18951, -1.45 DPS) [quest] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Edward the Odd (2243, +0.00 DPS) [world_drop]; Spellshifter Rod (9527, -4.55 DPS) [quest]; Staff of Hale Magefire (13000, -4.70 DPS) [world_drop] |
| off_hand | Lei of the Lifegiver (19312) | Stormpike Guard [rep] | 58.3 healing_power points (7.37 DPS) | yes | Grand Marshal's Tome of Restoration (234590, +0.00 DPS) [pvp]; Arcane Star (15931, -1.69 DPS) [world_drop]; Lapidis Tankard of Tidesippe (4696, -8.95 DPS, sim-verified) [world_drop] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 40.6 healing_power points (5.14 DPS) | yes | Jaina's Firestarter (13064, -2.21 DPS) [world_drop]; Brilliant Wand (249385, -2.68 DPS) [crafted]; Cairnstone Sliver (9654, -5.01 DPS, sim-verified) [quest] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Virtuous Robe; wrist: Virtuous Bracers; hands: Raider Handwraps; waist: Virtuous Belt; legs: Virtuous Skirt; feet: Virtuous Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket1: Serenity Field; trinket2: Darkspear Voodoo Seal; off_hand: Lei of the Lifegiver; ranged: Torch of Light

No-known-source sample (15 of 943, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

## Horde

### Band 20 (undead, 025003010000000000-00000000000000000-000000000000000000)

Set DPS (verified): 35.5. Weights run: 1.7s. Verify run: 5.3s. 138 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.006, intellect=2.353 ± 0.013, spirit=1.366 ± 0.007, mp5=3.060 ± 0.046, crit=0.161 ± 0.010 per rating point (14 rating = 1%, 2.260 per %), spell_haste=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Goggles (4373) | Engineering [crafted] | 20.0 healing_power points (0.94 DPS) | yes | Pristine Circlet (253949, +0.00 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -0.68 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 4.1 healing_power points (0.19 DPS) | yes | Scholarly Pendant (277203, -0.08 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 21.2 healing_power points (1.00 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.53 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.56 DPS) [dungeon] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | 11.7 healing_power points (0.55 DPS) | yes | Sanguine Cape (14376, -0.17 DPS, sim-verified) [world_drop]; Seer's Cape (6378, -0.20 DPS) [dungeon]; Pearl-clasped Cloak (5542, -0.22 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 28.9 healing_power points (1.36 DPS) | yes | Corsair's Overshirt (5202, -0.33 DPS) [dungeon]; Seer's Robe (2981, -0.50 DPS) [world_drop]; Robe of the Moccasin (6465, -0.60 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 20.9 healing_power points (0.99 DPS) | yes | Bright Bracers (3647, -0.54 DPS) [world_drop]; Crystalline Cuffs (14148, -0.57 DPS) [dungeon]; Featherbead Bracers (15452, -0.69 DPS, sim-verified) [quest] |
| hands | Blight Gloves (279877) | The New Plague [quest] | 23.3 healing_power points (1.10 DPS) | yes | Pristine Gloves (253913, -0.25 DPS) [crafted]; Magefist Gloves (12977, -0.29 DPS) [world_drop]; Bright Gloves (3066, -0.40 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 20.4 healing_power points (0.96 DPS) | yes | Novice Ardent's Sash (253887, -0.15 DPS) [crafted]; Keller's Girdle (2911, -0.20 DPS, sim-verified) [world_drop]; Tarantula Silk Sash (3229, -0.28 DPS) [world] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 37.6 healing_power points (1.77 DPS) | yes | Filigreed Silky Leggings (253939, -0.85 DPS) [crafted]; Filigreed Shadow Leggings (253943, -0.85 DPS) [crafted]; Darkweave Breeches (12987, -1.01 DPS, sim-verified) [world_drop] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 16.1 healing_power points (0.76 DPS) | yes | Walking Boots (4660, -0.31 DPS) [world]; Sanguine Sandals (14374, -0.31 DPS) [world_drop]; Bluegill Sandals (1560, -0.70 DPS, sim-verified) [world] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 14.1 healing_power points (0.66 DPS) | yes | Loop of Sacrifice (281673, -0.11 DPS) [quest]; Band of Purification (12996, -0.28 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.33 DPS) [world_drop] |
| finger2 | Black Pearl Ring (6332) | Lady Vespira [world] | 12.9 healing_power points (0.61 DPS) | yes | Loop of Sacrifice (281673, +0.00 DPS, sim-verified) [quest]; Band of Purification (12996, -0.22 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.28 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 23.5 healing_power points (1.11 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Staff of Orgrimmar (15444, -0.05 DPS) [quest]; Channeler's Staff (4437, -0.22 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Decay (5252) (or Flaring Baton (5326)) | Beren's Peril [quest] | 4.7 healing_power points (0.22 DPS) | yes | Flaring Baton (5326, +0.00 DPS) [quest]; Wisesight Wand (286750, -0.11 DPS) [world] |

**New at 20:** head: Shadow Goggles; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Blight Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Black Pearl Ring; main_hand: Gnarled Necromancer's Staff; ranged: Wand of Decay

No-known-source sample (15 of 138, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 209613 Insignia of the Alliance; 241089 Scarlet Dagger; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff; 263007 Skyseer's Vest; 263402 Sash of Sorrow

### Band 30 (undead, 025003031304000000-00000000000000000-000000000000000000)

Set DPS (verified): 78.0. Weights run: 1.6s. Verify run: 5.1s. 242 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.017, intellect=2.196 ± 0.016, spirit=1.734 ± 0.014, mp5=3.806 ± 0.015, crit=0.214 ± 0.012 per rating point (14 rating = 1%, 2.992 per %), spell_haste=not significant (0.275 ± 0.123)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 43.4 healing_power points (2.97 DPS) | yes | Nightsky Cowl (4039, -0.57 DPS) [world_drop]; Resilient Cap (14401, -0.84 DPS) [world_drop]; Embalmed Shroud (7691, -0.92 DPS, sim-verified) [dungeon] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.7 healing_power points (1.08 DPS) | yes | Scorn's Icy Choker (23169, -0.17 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -0.17 DPS) [world]; Crystal Starfire Medallion (5003, -0.31 DPS, sim-verified) [world_drop] |
| shoulder | Nightsky Mantle (4718) | World drop [world_drop] | 26.2 healing_power points (1.80 DPS) | yes | Ghostly Mantle (3324, -0.08 DPS, sim-verified) [quest]; Death Speaker Mantle (6685, -0.14 DPS) [dungeon]; Batwing Mantle (6697, -0.14 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | 22.8 healing_power points (1.56 DPS) | yes | Battle Healer's Cloak (19529, -0.19 DPS) [rep]; Glowing Thresher Cape (6901, -0.21 DPS, sim-verified) [dungeon]; Cloak of Rot (4462, -0.36 DPS) [world] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 44.2 healing_power points (3.02 DPS) | yes | Beguiler Robes (7728, -0.27 DPS) [dungeon]; Pristine Gown (253961, -0.32 DPS, sim-verified) [crafted]; Silver-thread Robe (4035, -0.89 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 22.4 healing_power points (1.53 DPS) | yes | Glowing Magical Bracelets (13106, -0.33 DPS) [world_drop]; Nightsky Wristbands (6407, -0.42 DPS, sim-verified) [world_drop]; Spidertank Oilrag (9448, -0.75 DPS) [dungeon] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 31.8 healing_power points (2.18 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.22 DPS, sim-verified) [dungeon]; Blight Gloves (279877, -0.53 DPS) [quest]; Truefaith Gloves (7049, -0.70 DPS) [crafted] |
| waist | Lilac Sash (6780) | Centaur Bounty [quest] | 25.0 healing_power points (1.71 DPS) | yes | Pristine Sash (253925, -0.35 DPS) [crafted]; Dreamer's Belt (4829, -0.42 DPS) [vendor]; Resilient Cord (14406, -0.48 DPS, sim-verified) [world_drop] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 46.0 healing_power points (3.15 DPS) | yes | Filigreed Pristine Leggings (253937, -0.77 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -0.82 DPS) [crafted]; Sacred Burial Trousers (6282, -1.33 DPS) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 42.3 healing_power points (2.90 DPS) | yes | Soggy Boots (274747, -1.34 DPS) [vendor]; Frothing Slippers (254003, -1.37 DPS) [crafted]; Acidic Walkers (9454, -1.57 DPS, sim-verified) [dungeon] |
| finger1 | The Queen's Jewel (13094) | World drop [world_drop] | 18.3 healing_power points (1.25 DPS) | yes | Runemender's Seal (273031, -0.01 DPS) [world]; Sea Giant's Toe Ring (274746, -0.02 DPS) [vendor]; Solid Gold Band (279845, -0.18 DPS) [quest] |
| finger2 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 18.1 healing_power points (1.24 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; Sea Giant's Toe Ring (274746, -0.01 DPS) [vendor]; Solid Gold Band (279845, -0.17 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | 98.7 healing_power points (6.76 DPS) | yes | Royal Diplomatic Scepter (9457, -0.50 DPS, sim-verified) [dungeon]; Death Speaker Scepter (2816, -2.10 DPS) [dungeon]; Spine of the Devourer (273040, -4.29 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 16.4 healing_power points (1.12 DPS) | yes | Summoner's Wand (5245, -0.50 DPS, sim-verified) [world]; Lesser Mystic Wand (11289, -0.52 DPS) [crafted]; Starfaller (13063, -0.52 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Nightsky Mantle; back: Darkspear Raider's Cloak; chest: Death Speaker Robes; hands: Gloves of Old; waist: Lilac Sash; legs: Pristine Leggings; feet: Gilded Slippers; finger1: The Queen's Jewel; finger2: Electrocutioner Lagnut; main_hand: Wind Spirit Staff; ranged: Wand of Mana Concentration

No-known-source sample (15 of 242, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (undead, 025003031305101520-00000000000000000-000000000000000000)

Set DPS (verified): 134.8. Weights run: 1.8s. Verify run: 14.4s. 319 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.023, intellect=2.189 ± 0.028, spirit=0.629 ± 0.042, mp5=2.446 ± 0.048, crit=0.446 ± 0.019 per rating point (14 rating = 1%, 6.248 per %), spell_haste=-4.866 ± 0.389

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 56.9 healing_power points (4.25 DPS) | yes | Holy Shroud (2721, -1.50 DPS) [world_drop]; Craftsman's Monocle (4393, -1.80 DPS) [crafted]; Thinking Cap (2624, -3.50 DPS, sim-verified) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 30.5 healing_power points (2.28 DPS) | yes | Necklace of Calisea (1714, -0.63 DPS, sim-verified) [dungeon]; Darkspear Warding Pendant (272074, -1.13 DPS) [vendor]; Archmagister's Faceted Pendant (273049, -1.30 DPS) [world] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 43.5 healing_power points (3.24 DPS) | yes | Mistscape Mantle (4734, -1.21 DPS) [dungeon]; Silksand Shoulder Pads (14423, -1.31 DPS) [world_drop]; Windchaser Amice (14432, -1.50 DPS, sim-verified) [world_drop] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 26.6 healing_power points (1.99 DPS) | yes | Battle Healer's Cloak (19528, +0.00 DPS, sim-verified) [rep]; Blackforge Cape (6424, -0.49 DPS) [dungeon]; Cloak of Rot (4462, -0.68 DPS) [world] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 63.0 healing_power points (4.71 DPS) | yes | Red Mageweave Vest (10007, -1.76 DPS) [crafted]; Silksand Tunic (14417, -1.86 DPS) [world_drop]; Death Speaker Robes (6682, -5.31 DPS, sim-verified) [dungeon] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 19.7 healing_power points (1.47 DPS) | yes | Mindthrust Bracers (1974, -0.14 DPS, sim-verified) [dungeon]; Mistscape Bracers (4045, -0.16 DPS) [dungeon]; Enchanted Stonecloth Bracers (4979, -0.16 DPS) [quest] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (134.8 DPS) | yes | Earthen Silk Gloves (254017, -1.06 DPS) [crafted]; Red Mageweave Gloves (10018, -1.15 DPS) [crafted]; Gilded Handwraps (254021, -2.24 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 43.3 healing_power points (3.23 DPS) | yes | Razzeric's Customized Seatbelt (6726, -1.27 DPS) [quest]; Sutarn's Ring (13105, -1.32 DPS) [world_drop]; Deathmage Sash (10771, -2.21 DPS, sim-verified) [dungeon] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-verified (134.8 DPS) | yes | Stoneweaver Leggings (9407, +0.00 DPS) [dungeon]; Filigreed Pristine Leggings (253937, +0.00 DPS) [crafted]; Pristine Leggings (253987, -1.81 DPS, sim-verified) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 37.8 healing_power points (2.83 DPS) | yes | Furen's Boots (13100, -0.17 DPS, sim-verified) [world_drop]; Boots of the Maharishi (9658, -1.07 DPS) [quest]; Kodo Rustler Boots (15697, -1.28 DPS) [quest] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.34 DPS) | yes | Ogremind Ring (1993, -0.06 DPS) [world_drop]; Mindbender Loop (5009, -0.11 DPS) [world_drop]; Welken Ring (5011, -0.20 DPS) [world_drop] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 17.2 healing_power points (1.28 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Mindbender Loop (5009, -0.05 DPS) [world_drop]; Welken Ring (5011, -0.14 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (134.8 DPS) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (134.8 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Wind Spirit Staff (6689, -1.34 DPS) [dungeon]; Royal Diplomatic Scepter (9457, -1.47 DPS) [dungeon] |
| off_hand | Aurora Sphere (7610) | Uldaman: Ancient Treasure [dungeon] | 17.8 healing_power points (1.33 DPS) | yes | Orb of Souls (249395, -0.08 DPS) [crafted]; Beacon of Hope (9393, -0.10 DPS) [dungeon]; Witch's Finger (16887, -0.81 DPS, sim-verified) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 15.0 healing_power points (1.12 DPS) | yes | Flash Wand (5248, -0.33 DPS) [quest]; Goblin Igniter (5253, -0.33 DPS) [quest]; Wand of Mana Concentration (273043, -0.34 DPS, sim-verified) [world] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Stormcloth Gloves; waist: Gilded Cord; legs: Stormcloth Pants; finger1: Sea Giant's Toe Ring; finger2: Voodoo Band; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Aurora Sphere; ranged: Jaina's Firestarter

No-known-source sample (15 of 319, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16336 Sergeant Major's Cape

### Band 50 (undead, 025003031305101520-03502000000000000-000000000000000000)

Set DPS (verified): 192.0. Weights run: 1.8s. Verify run: 7.7s. 410 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.019, intellect=2.517 ± 0.048, spirit=1.688 ± 0.064, mp5=4.067 ± 0.108, crit=0.573 ± 0.024 per rating point (14 rating = 1%, 8.017 per %), spell_haste=-6.791 ± 0.586

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 82.2 healing_power points (6.59 DPS) | yes | Blood Guard's Satin Cover (220899, -0.01 DPS) [vendor]; Stormcloth Headband (10032, -1.29 DPS) [crafted]; Papal Fez (9431, -1.72 DPS, sim-verified) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 42.0 healing_power points (3.37 DPS) | yes | Glowing Eye of Mordresh (10769, -0.71 DPS) [dungeon]; Necklace of Calisea (1714, -1.01 DPS) [dungeon]; Darkmoon Necklace (19303, -1.32 DPS, sim-verified) [vendor] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | 56.3 healing_power points (4.51 DPS) | yes | Blood Guard's Satin Pads (220901, +0.00 DPS) [vendor]; Rotgrip Mantle (17732, -0.88 DPS) [dungeon]; Inquisitor's Shawl (19507, -2.56 DPS, sim-verified) [dungeon] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 43.7 healing_power points (3.50 DPS) | yes | Battle Healer's Cloak (19527, -0.93 DPS) [rep]; Highborne Cloak (14450, -1.15 DPS) [world_drop]; Imperial Red Cloak (8248, -2.36 DPS, sim-verified) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 71.5 healing_power points (5.74 DPS) | yes | Ghostweave Vest (14141, -0.42 DPS) [crafted]; Earth Warder's Vest (21311, -1.15 DPS) [quest]; Stone Guard's Satin Armor (220903, -2.68 DPS, sim-verified) [vendor] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 46.1 healing_power points (3.70 DPS) | yes | Forgotten Wraps (9433, -1.27 DPS) [world_drop]; Bloodband Bracers (11469, -1.47 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -2.82 DPS, sim-verified) [quest] |
| hands | Virtuous Mitts (226950) | Mokvar [vendor] | 82.1 healing_power points (6.59 DPS) | yes | Gilded Gloves (254095, -1.16 DPS) [crafted]; Virtuous Hands (226958, -2.33 DPS) [vendor]; Raider Handwraps (272098, -2.69 DPS, sim-verified) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 59.8 healing_power points (4.80 DPS) | yes | Gilded Cord (254037, +0.00 DPS, sim-verified) [crafted]; Gilded Waistcord (254081, -0.66 DPS) [crafted]; Earthenweave Cord (254077, -1.22 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 81.7 healing_power points (6.55 DPS) | yes | Senior Designer's Pantaloons (11841, -0.49 DPS, sim-verified) [dungeon]; Stone Guard's Satin Leggings (220902, -0.83 DPS) [vendor]; Dalewind Trousers (13008, -1.35 DPS) [world_drop] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 53.7 healing_power points (4.31 DPS) | yes | First Sergeant's Satin Boots (220900, -0.28 DPS) [vendor]; Furen's Boots (13100, -0.72 DPS) [world_drop]; Vinerot Sandals (17748, -1.36 DPS, sim-verified) [dungeon] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 35.4 healing_power points (2.84 DPS) | yes | Welken Ring (5011, -0.88 DPS) [world_drop]; Coldwater Ring (4550, -1.01 DPS) [quest]; Voodoo Band (1996, -1.02 DPS) [world] |
| finger2 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 28.5 healing_power points (2.28 DPS) | yes | Welken Ring (5011, -0.33 DPS) [world_drop]; Coldwater Ring (4550, -0.46 DPS) [quest]; Voodoo Band (1996, -0.46 DPS) [world] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (192.0 DPS) | yes | Uther's Strength (11302, -3.45 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -3.54 DPS) [quest]; Alchemist's Stone (13503, -4.35 DPS) [crafted] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (192.0 DPS) | yes | Uther's Strength (11302, -0.18 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.27 DPS) [quest]; Alchemist's Stone (13503, -0.74 DPS, sim-verified) [crafted] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-verified (192.0 DPS) | yes | Hammer of the Northern Wind (810, +0.00 DPS) [world_drop]; Spellshifter Rod (9527, -3.20 DPS) [quest]; Hand of Righteousness (7721, -3.56 DPS) [dungeon] |
| off_hand | Twisting Essence Jar (249456) | Enchanting [crafted] | 32.5 healing_power points (2.61 DPS) | yes | Jademoon Orb (11859, -0.72 DPS) [quest]; Orb of Souls (249395, -0.75 DPS) [crafted]; Aurora Sphere (7610, -1.72 DPS, sim-verified) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 20.2 healing_power points (1.62 DPS) | yes | Flash Wand (5248, -0.40 DPS) [quest]; Goblin Igniter (5253, -0.40 DPS) [quest]; Wand of Mana Concentration (273043, -1.02 DPS, sim-verified) [world] |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Nethergeld Shoulders; back: Darkspear Raider's Cloak; wrist: Nethergeld Cuffs; hands: Virtuous Mitts; waist: Dawnspire Cord; legs: Kilt of the Atal'ai Prophet; feet: Gilded Sandals; finger1: Eye of Adaegus; finger2: Darkspear Signet; trinket1: Darkspear Voodoo Seal; main_hand: Charstone Dirk; off_hand: Twisting Essence Jar

No-known-source sample (15 of 410, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60 (undead, 025003031305101520-03505003030100000-000000000000000000)

Set DPS (verified): 318.7. Weights run: 1.8s. Verify run: 10.4s. 935 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.244, intellect=2.919 ± 0.088, spirit=4.482 ± 0.666, mp5=3.713 ± 0.283, crit=0.817 ± 0.053 per rating point (14 rating = 1%, 11.434 per %), spell_haste=not significant (-5.264 ± 1.542)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 165.8 healing_power points (12.51 DPS) | yes | Warlord's Satin Hood (231635, -0.20 DPS) [vendor]; Champion's Satin Hood (227118, -0.79 DPS) [pvp]; Devout Crown (16693, -13.03 DPS, sim-verified) [dungeon] |
| neck | The Eye of Zuldazar (19593) (or The All-Seeing Eye of Zuldazar (19594)) | The Eye of Zuldazar [quest] | 99.6 healing_power points (7.52 DPS) | yes | The All-Seeing Eye of Zuldazar (19594, +0.00 DPS) [quest]; Wavefront Necklace (20685, -2.14 DPS) [world]; Lady Alizabeth's Pendant (13002, -2.44 DPS) [world_drop] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 126.1 healing_power points (9.52 DPS) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Devout Mantle (16695, -1.02 DPS) [dungeon]; Virtuous Mantle (226951, -3.02 DPS, sim-verified) [quest] |
| back | Darkspear Raider's Cloak (272063) | Creeg Bothunk [vendor] | 73.6 healing_power points (5.55 DPS) | yes | Hide of the Wild (18510, +0.00 DPS, sim-verified) [crafted]; Battle Healer's Cloak (19526, -0.89 DPS) [rep]; Cloak of Blight (6832, -1.16 DPS) [quest] |
| chest | Virtuous Robe (226945) | Saving the Best for Last [quest] | sim-verified (+7.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Devout Robe (16690, -0.38 DPS) [dungeon]; Mooncloth Vest (14138, -7.12 DPS, sim-verified) [crafted] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-verified (+6.1 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Satin Bracers (17619, +0.00 DPS) [pvp]; Devout Bracers (16697, -1.09 DPS) [dungeon]; Bracers of Hope (22667, -6.11 DPS, sim-verified) [quest] |
| hands | Virtuous Mitts (226950) | Mokvar [vendor] | 136.8 healing_power points (10.32 DPS) | yes | Swarmtender's Gloves (275607, -2.35 DPS) [crafted]; Magister's Gloves (16684, -2.50 DPS) [dungeon]; Devout Gloves (16692, -12.97 DPS, sim-verified) [dungeon] |
| waist | Virtuous Belt (226948) | Just Compensation [quest] | sim-verified (+5.5 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Satin Cinch (17621, +0.00 DPS) [pvp]; Wisdom of the Timbermaw (19047, -0.52 DPS) [crafted]; Penitent's Cinch (272394, -5.54 DPS, sim-verified) [vendor] |
| legs | Virtuous Skirt (226946) | Anthion's Parting Words [quest] | sim-verified (+5.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Legionnaire's Satin Legguards (227123, -0.89 DPS) [pvp]; General's Satin Legguards (231634, -1.34 DPS) [vendor]; Devout Skirt (16694, -5.21 DPS, sim-verified) [dungeon] |
| feet | Virtuous Sandals (226952) | Anthion's Parting Words [quest] | 155.0 healing_power points (11.70 DPS) | yes | Mooncloth Boots (15802, -2.41 DPS) [crafted]; Devout Sandals (16691, -2.92 DPS) [dungeon]; Incandescent Mooncloth Boots (227862, -9.31 DPS, sim-verified) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 79.8 healing_power points (6.03 DPS) | yes | Channeler's Ring (272406, -0.24 DPS) [vendor]; Eye of Adaegus (5266, -0.64 DPS) [world_drop]; Blessed Band of Light (272407, -0.98 DPS) [vendor] |
| finger2 | Band of the Hierophant (13096) | World drop [world_drop] | 78.5 healing_power points (5.92 DPS) | yes | Eye of Adaegus (5266, -0.54 DPS) [world_drop]; Blessed Band of Light (272407, -0.88 DPS) [vendor]; Channeler's Ring (272406, -0.98 DPS, sim-verified) [vendor] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor]; Uther's Strength (11302, -4.21 DPS, sim-verified) [world_drop] |
| main_hand | Staff of Hale Magefire (13000) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -1.21 DPS) [dungeon]; Soulkeeper (1607, -1.73 DPS) [world_drop]; Hand of Edward the Odd (2243, -10.76 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 66.3 healing_power points (5.00 DPS) | yes | Chillnail Splinter (10704, -2.97 DPS) [quest]; Goblin Igniter (5253, -3.11 DPS) [quest]; Jaina's Firestarter (13064, -5.05 DPS, sim-verified) [world_drop] |

**New at 60:** head: Virtuous Crown; neck: The Eye of Zuldazar; shoulder: Argent Elite Shoulders; back: Darkspear Raider's Cloak; chest: Virtuous Robe; wrist: Virtuous Bracers; waist: Virtuous Belt; legs: Virtuous Skirt; feet: Virtuous Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of the Hierophant; trinket1: Shard of the Splithooves; trinket2: Serenity Field; main_hand: Staff of Hale Magefire; ranged: Torch of Light

No-known-source sample (15 of 935, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60, raid preset (undead, 025003031305101520-03505003030100000-000000000000000000)

Set DPS (verified): 577.7. Weights run: 1.5s. Verify run: 7.6s. 935 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.382, intellect=2.779 ± 0.085, spirit=2.149 ± 0.079, mp5=3.667 ± 0.090, crit=0.950 ± 0.061 per rating point (14 rating = 1%, 13.304 per %), spell_haste=not significant (-2.876 ± 1.662)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 127.9 healing_power points (16.17 DPS) | yes | Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Champion's Satin Hood (227118, -0.99 DPS) [pvp]; Gnomish Turban of Psychic Might (21517, -12.06 DPS, sim-verified) [quest] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 70.0 healing_power points (8.85 DPS) | yes | The All-Seeing Eye of Zuldazar (19594, -0.23 DPS) [quest]; Drake Tooth Necklace (21531, -1.35 DPS) [quest]; The Eye of Zuldazar (19593, -5.67 DPS, sim-verified) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 102.5 healing_power points (12.95 DPS) | yes | Champion's Satin Mantle (227120, -1.02 DPS) [pvp]; Warlord's Satin Mantle (231631, -2.14 DPS) [pvp]; Devout Mantle (16695, -11.53 DPS, sim-verified) [dungeon] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 69.8 healing_power points (8.82 DPS) | yes | Darkspear Raider's Cloak (272076, -2.55 DPS) [vendor]; Drape of Recovery (272413, -3.05 DPS) [vendor] |
| chest | Virtuous Robe (226945) | Saving the Best for Last [quest] | 127.0 healing_power points (16.05 DPS) | yes | Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Devout Robe (16690, -1.02 DPS) [dungeon]; Mooncloth Vest (14138, -10.35 DPS, sim-verified) [crafted] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-verified (+10.3 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Satin Bracers (17619, +0.00 DPS) [pvp]; Devout Bracers (16697, -1.51 DPS) [dungeon]; Bracers of Hope (22667, -10.34 DPS, sim-verified) [quest] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 99.0 healing_power points (12.51 DPS) | yes | Virtuous Mitts (226950, -0.73 DPS) [vendor]; Devout Gloves (16692, -2.49 DPS) [dungeon]; Mooncloth Gloves (18409, -2.73 DPS) [crafted] |
| waist | Virtuous Belt (226948) | Just Compensation [quest] | 89.2 healing_power points (11.28 DPS) | yes | General's Satin Cinch (17621, +0.00 DPS) [pvp]; Devout Belt (16696, -0.42 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -6.98 DPS, sim-verified) [crafted] |
| legs | Virtuous Skirt (226946) | Anthion's Parting Words [quest] | sim-verified (+11.4 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Satin Legguards (231634, +0.00 DPS) [vendor]; Legionnaire's Satin Legguards (227123, -0.30 DPS) [pvp]; Devout Skirt (16694, -11.40 DPS, sim-verified) [dungeon] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 113.0 healing_power points (14.28 DPS) | yes | Virtuous Sandals (226952, +0.00 DPS, sim-verified) [quest]; Mooncloth Boots (15802, -2.80 DPS) [crafted]; General's Satin Walkers (231630, -4.29 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 78.3 healing_power points (9.90 DPS) | yes | Channeler's Ring (272406, -2.81 DPS) [vendor]; Band of the Hierophant (13096, -3.40 DPS) [world_drop]; Blessed Band of Light (272407, -3.81 DPS) [vendor] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 64.5 healing_power points (8.15 DPS) | yes | Band of the Hierophant (13096, -1.65 DPS) [world_drop]; Blessed Band of Light (272407, -2.06 DPS) [vendor]; Channeler's Ring (272406, -5.38 DPS, sim-verified) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+15.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, -3.69 DPS) [world_drop]; Shard of the Splithooves (10659, -4.17 DPS) [quest]; Evonice's Landin' Pilla (18951, -4.24 DPS) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -2.05 DPS) [world_drop]; Shard of the Splithooves (10659, -2.53 DPS) [quest]; Uther's Strength (11302, -6.64 DPS, sim-verified) [world_drop] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Edward the Odd (2243, +0.00 DPS) [world_drop]; Spellshifter Rod (9527, -4.55 DPS) [quest]; Staff of Hale Magefire (13000, -4.70 DPS) [world_drop] |
| off_hand | Lei of the Lifegiver (19312) | Frostwolf Clan [rep] | 58.3 healing_power points (7.37 DPS) | yes | High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp]; Arcane Star (15931, -1.69 DPS) [world_drop]; Lapidis Tankard of Tidesippe (4696, -12.49 DPS, sim-verified) [world_drop] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 40.6 healing_power points (5.14 DPS) | yes | Brilliant Wand (249385, -2.68 DPS) [crafted]; Wand of Mana Concentration (273043, -2.80 DPS) [world]; Jaina's Firestarter (13064, -6.84 DPS, sim-verified) [world_drop] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Virtuous Robe; wrist: Virtuous Bracers; hands: Raider Handwraps; waist: Virtuous Belt; legs: Virtuous Skirt; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Serenity Field; off_hand: Lei of the Lifegiver; ranged: Torch of Light

No-known-source sample (15 of 935, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

