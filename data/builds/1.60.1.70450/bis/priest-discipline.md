# Leveling BiS: Discipline

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 025004000000000000-00000000000000000-00000000000000000)

Set DPS (verified): 36.5. Weights run: 10.7s. Verify run: 29.4s. 146 eligible items had no known source.

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

Set DPS (verified): 67.6. Weights run: 10.9s. Verify run: 32.3s. 258 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.017, intellect=2.171 ± 0.014, spirit=1.721 ± 0.012, mp5=3.882 ± 0.013, crit=0.214 ± 0.009 per rating point (14 rating = 1%, 2.993 per %), spell_haste=not significant (0.202 ± 0.091)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 43.3 healing_power points (2.75 DPS) | yes | Resilient Cap (14401, -0.80 DPS) [world_drop]; Shadow Hood (4323, -1.23 DPS) [crafted]; Nightsky Cowl (4039, -1.35 DPS, sim-verified) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.6 healing_power points (0.99 DPS) | yes | Scorn's Icy Choker (23169, -0.16 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -0.16 DPS) [world]; Crystal Starfire Medallion (5003, -0.22 DPS, sim-verified) [world_drop] |
| shoulder | Mantle of Honor (3560) | Bride of the Embalmer [quest] | 27.2 healing_power points (1.73 DPS) | yes | Nightsky Mantle (4718, -0.08 DPS, sim-verified) [world_drop]; Death Speaker Mantle (6685, -0.21 DPS) [dungeon]; Faerie Mantle (5820, -0.25 DPS) [quest] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | 22.5 healing_power points (1.43 DPS) | yes | Glowing Thresher Cape (6901, -0.11 DPS) [dungeon]; Caretaker's Cape (19533, -0.17 DPS) [rep]; Prelacy Cape (7004, -0.19 DPS, sim-verified) [quest] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 43.9 healing_power points (2.78 DPS) | yes | Beguiler Robes (7728, -0.26 DPS) [dungeon]; Pristine Gown (253961, -0.40 DPS, sim-verified) [crafted]; Silver-thread Robe (4035, -0.83 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 22.5 healing_power points (1.43 DPS) | yes | Glowing Magical Bracelets (13106, -0.33 DPS) [world_drop]; Nightsky Wristbands (6407, -0.49 DPS, sim-verified) [world_drop]; Spidertank Oilrag (9448, -0.69 DPS) [dungeon] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 31.5 healing_power points (2.00 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.22 DPS, sim-verified) [dungeon]; Truefaith Gloves (7049, -0.63 DPS) [crafted]; Zodiac Gloves (7106, -0.68 DPS) [quest] |
| waist | Resilient Cord (14406) | World drop [world_drop] | 19.9 healing_power points (1.26 DPS) | yes | Pristine Sash (253925, +0.00 DPS, sim-verified) [crafted]; Dreamer's Belt (4829, -0.08 DPS) [vendor]; Novice Ardent's Sash (253887, -0.14 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 45.8 healing_power points (2.91 DPS) | yes | Earthen Leggings (253999, -0.75 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.84 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -1.29 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 42.1 healing_power points (2.67 DPS) | yes | Soggy Boots (274747, -1.25 DPS) [vendor]; Frothing Slippers (254003, -1.27 DPS) [crafted]; Acidic Walkers (9454, -1.78 DPS, sim-verified) [dungeon] |
| finger1 | The Queen's Jewel (13094) | World drop [world_drop] | 18.1 healing_power points (1.15 DPS) | yes | Runemender's Seal (273031, -0.00 DPS) [world]; Sea Giant's Toe Ring (274746, -0.01 DPS) [vendor]; Darkspear Signet (272071, -0.16 DPS) [vendor] |
| finger2 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 18.0 healing_power points (1.14 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; Sea Giant's Toe Ring (274746, -0.00 DPS) [vendor]; Darkspear Signet (272071, -0.16 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-verified (67.7 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -0.03 DPS) [quest]; Death Speaker Scepter (2816, -1.19 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 16.5 healing_power points (1.04 DPS) | yes | Lesser Mystic Wand (11289, -0.49 DPS) [crafted]; Starfaller (13063, -0.49 DPS) [world_drop]; Summoner's Wand (5245, -0.83 DPS, sim-verified) [world] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Mantle of Honor; back: Darkspear Raider's Cloak; chest: Death Speaker Robes; hands: Gloves of Old; waist: Resilient Cord; legs: Pristine Leggings; feet: Gilded Slippers; finger1: The Queen's Jewel; finger2: Electrocutioner Lagnut; main_hand: Spine of the Devourer; ranged: Wand of Mana Concentration

No-known-source sample (15 of 258, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 40 (human, 025005031330011520-00000000000000000-00000000000000000)

Set DPS (verified): 135.7. Weights run: 12.6s. Verify run: 79.7s. 339 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.013, intellect=2.181 ± 0.022, spirit=0.858 ± 0.037, mp5=2.390 ± 0.040, crit=0.420 ± 0.013 per rating point (14 rating = 1%, 5.876 per %), spell_haste=-12.136 ± 0.403

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 58.8 healing_power points (4.40 DPS) | yes | Whitemane's Chapeau (7720, -1.11 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -1.55 DPS) [world_drop]; Thinking Cap (2624, -1.63 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 30.4 healing_power points (2.28 DPS) | yes | Necklace of Calisea (1714, -0.59 DPS, sim-verified) [dungeon]; Darkspear Warding Pendant (272074, -1.14 DPS) [vendor]; Archmagister's Faceted Pendant (273049, -1.30 DPS) [world] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 43.3 healing_power points (3.25 DPS) | yes | Mistscape Mantle (4734, -1.13 DPS) [dungeon]; Silksand Shoulder Pads (14423, -1.26 DPS) [world_drop]; Windchaser Amice (14432, -1.49 DPS, sim-verified) [world_drop] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 27.4 healing_power points (2.05 DPS) | yes | Caretaker's Cape (19532, -0.38 DPS) [rep]; Blackforge Cape (6424, -0.49 DPS) [dungeon]; Sergeant Major's Cape (16336, -0.69 DPS) [pvp] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 64.9 healing_power points (4.86 DPS) | yes | Red Mageweave Vest (10007, -1.92 DPS) [crafted]; Silksand Tunic (14417, -1.92 DPS) [world_drop]; Death Speaker Robes (6682, -5.50 DPS, sim-verified) [dungeon] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 19.6 healing_power points (1.47 DPS) | yes | Mindthrust Bracers (1974, -0.15 DPS, sim-verified) [dungeon]; Mistscape Bracers (4045, -0.16 DPS) [dungeon]; Enchanted Stonecloth Bracers (4979, -0.16 DPS) [quest] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (135.7 DPS) | yes | Earthen Silk Gloves (254017, -0.92 DPS) [crafted]; Red Mageweave Gloves (10018, -1.15 DPS) [crafted]; Gilded Handwraps (254021, -2.02 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 44.6 healing_power points (3.34 DPS) | yes | Sutarn's Ring (13105, -1.32 DPS) [world_drop]; Razzeric's Customized Seatbelt (6726, -1.38 DPS) [quest]; Deathmage Sash (10771, -2.57 DPS, sim-verified) [dungeon] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-verified (135.7 DPS) | yes | Stoneweaver Leggings (9407, +0.00 DPS) [dungeon]; Filigreed Pristine Leggings (253937, +0.00 DPS) [crafted]; Pristine Leggings (253987, -1.99 DPS, sim-verified) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 38.7 healing_power points (2.90 DPS) | yes | Furen's Boots (13100, +0.00 DPS, sim-verified) [world_drop]; Kodo Rustler Boots (15697, -1.27 DPS) [quest]; Acidic Walkers (9454, -1.33 DPS) [dungeon] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.35 DPS) | yes | Ogremind Ring (1993, -0.01 DPS) [world_drop]; Mindbender Loop (5009, -0.08 DPS) [world_drop]; Welken Ring (5011, -0.08 DPS) [world_drop] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 17.8 healing_power points (1.34 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Mindbender Loop (5009, -0.06 DPS) [world_drop]; Welken Ring (5011, -0.07 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (135.7 DPS) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (135.7 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Death Speaker Scepter (2816, -2.46 DPS) [dungeon]; Hypnotic Blade (7714, -2.76 DPS) [dungeon] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 33.8 healing_power points (2.53 DPS) | yes | Orb of Souls (249395, -1.17 DPS) [crafted]; Beacon of Hope (9393, -1.23 DPS) [dungeon]; Aurora Sphere (7610, -1.46 DPS, sim-verified) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 15.7 healing_power points (1.17 DPS) | yes | Flash Wand (5248, -0.33 DPS) [quest]; Goblin Igniter (5253, -0.33 DPS) [quest]; Wand of Mana Concentration (273043, -0.38 DPS, sim-verified) [world] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Stormcloth Gloves; waist: Gilded Cord; legs: Stormcloth Pants; finger1: Sea Giant's Toe Ring; finger2: Voodoo Band; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Orb of Lorica; ranged: Jaina's Firestarter

No-known-source sample (15 of 339, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 50 (human, 025005031330011520-03502000000000000-00000000000000000)

Set DPS (verified): 214.8. Weights run: 12.5s. Verify run: 44.1s. 433 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.050, intellect=2.721 ± 0.033, spirit=2.154 ± 0.046, mp5=4.514 ± 0.065, crit=0.646 ± 0.021 per rating point (14 rating = 1%, 9.050 per %), spell_haste=-9.814 ± 0.496

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 94.0 healing_power points (6.72 DPS) | yes | Papal Fez (9431, -1.03 DPS) [dungeon]; Imperial Red Circlet (8254, -1.30 DPS) [world_drop]; Knight-Lieutenant's Satin Cover (220896, -1.49 DPS, sim-verified) [vendor] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 46.7 healing_power points (3.34 DPS) | yes | Glowing Eye of Mordresh (10769, -0.85 DPS) [dungeon]; Darkmoon Necklace (19303, -0.87 DPS, sim-verified) [vendor]; Necklace of Calisea (1714, -0.90 DPS) [dungeon] |
| shoulder | Knight-Lieutenant's Satin Pads (220894) | Captain Dirgehammer [vendor] | 67.3 healing_power points (4.81 DPS) | yes | Inquisitor's Shawl (19507, -1.21 DPS) [dungeon]; Rotgrip Mantle (17732, -1.31 DPS) [dungeon]; Nethergeld Shoulders (254049, -1.99 DPS, sim-verified) [crafted] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 48.9 healing_power points (3.49 DPS) | yes | Caretaker's Cape (19531, -1.00 DPS) [rep]; Highborne Cloak (14450, -1.13 DPS) [world_drop]; Imperial Red Cloak (8248, -2.57 DPS, sim-verified) [world_drop] |
| chest | Ghostweave Vest (14141) | Tailoring [crafted] | 79.2 healing_power points (5.66 DPS) | yes | Stormcloth Vest (10020, -0.28 DPS) [crafted]; Knight's Satin Armor (220892, -0.29 DPS) [vendor]; Grimlok's Tribal Vestments (9415, -1.17 DPS) [dungeon] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 49.8 healing_power points (3.56 DPS) | yes | Forgotten Wraps (9433, -1.23 DPS) [world_drop]; Bloodband Bracers (11469, -1.35 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -2.58 DPS, sim-verified) [quest] |
| hands | Virtuous Mitts (226950) | Mokvar [vendor] | 92.7 healing_power points (6.62 DPS) | yes | Gilded Gloves (254095, -1.38 DPS) [crafted]; Virtuous Hands (226958, -2.25 DPS) [vendor]; Raider Handwraps (272098, -2.64 DPS, sim-verified) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 63.7 healing_power points (4.55 DPS) | yes | Gilded Cord (254037, +0.00 DPS, sim-verified) [crafted]; Gilded Waistcord (254081, -0.73 DPS) [crafted]; Earthenweave Cord (254077, -1.00 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 93.8 healing_power points (6.70 DPS) | yes | Senior Designer's Pantaloons (11841, -0.38 DPS, sim-verified) [dungeon]; Dalewind Trousers (13008, -1.10 DPS) [world_drop]; Knight's Satin Leggings (220893, -1.33 DPS) [vendor] |
| feet | Vinerot Sandals (17748) | Maraudon: Razorlash [dungeon] | 58.0 healing_power points (4.15 DPS) | yes | Furen's Boots (13100, -0.32 DPS) [world_drop]; Sergeant Major's Satin Boots (220895, -0.38 DPS) [vendor]; Gilded Sandals (254107, -2.02 DPS, sim-verified) [crafted] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 42.2 healing_power points (3.01 DPS) | yes | Welken Ring (5011, -0.96 DPS) [world_drop]; Choking Band (11868, -1.01 DPS) [quest]; Coldwater Ring (4550, -1.05 DPS) [quest] |
| finger2 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 31.6 healing_power points (2.26 DPS) | yes | Choking Band (11868, -0.26 DPS) [quest]; Coldwater Ring (4550, -0.29 DPS) [quest]; Welken Ring (5011, -2.77 DPS, sim-verified) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (214.8 DPS) | yes | Ankh of Life (1713, -2.99 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -3.30 DPS) [quest]; Uther's Strength (11302, -3.55 DPS) [world_drop] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (214.8 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (214.8 DPS) | yes | Charstone Dirk (17710, -6.67 DPS) [dungeon]; Spellshifter Rod (9527, -9.41 DPS) [quest]; Hammer of the Northern Wind (810, -28.61 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cairnstone Sliver (9654) | The Morrow Stone [quest] | 24.4 healing_power points (1.74 DPS) | yes | Wand of Mana Concentration (273043, -0.32 DPS) [world]; Goblin Igniter (5253, -0.50 DPS) [quest]; Jaina's Firestarter (13064, -0.59 DPS, sim-verified) [world_drop] |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Knight-Lieutenant's Satin Pads; back: Darkspear Raider's Cloak; chest: Ghostweave Vest; wrist: Nethergeld Cuffs; hands: Virtuous Mitts; waist: Dawnspire Cord; legs: Kilt of the Atal'ai Prophet; feet: Vinerot Sandals; finger1: Eye of Adaegus; finger2: Darkspear Signet; trinket1: Darkspear Voodoo Seal; trinket2: Thunderbrew's Boot Flask; main_hand: Glowing Brightwood Staff; ranged: Cairnstone Sliver

No-known-source sample (15 of 433, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60 (human, 025005031330011520-03505003030100000-00000000000000000)

Set DPS (verified): 358.7. Weights run: 12.0s. Verify run: 53.2s. 947 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.107, intellect=2.093 ± 0.043, spirit=2.892 ± 0.313, mp5=2.333 ± 0.149, crit=0.663 ± 0.028 per rating point (14 rating = 1%, 9.278 per %), spell_haste=-5.624 ± 0.808

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 124.1 healing_power points (13.94 DPS) | yes | Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Lieutenant Commander's Satin Hood (227121, -0.67 DPS) [pvp]; Devout Crown (16693, -12.97 DPS, sim-verified) [dungeon] |
| neck | The Eye of Zuldazar (19593) (or The All-Seeing Eye of Zuldazar (19594)) | The Eye of Zuldazar [quest] | 72.3 healing_power points (8.12 DPS) | yes | The All-Seeing Eye of Zuldazar (19594, +0.00 DPS) [quest]; Wavefront Necklace (20685, -1.92 DPS) [world]; Drake Tooth Necklace (21531, -2.38 DPS) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 101.9 healing_power points (11.44 DPS) | yes | Lieutenant Commander's Satin Mantle (227119, -0.60 DPS) [pvp]; Virtuous Mantle (226951, -0.75 DPS, sim-verified) [quest]; Devout Mantle (16695, -2.35 DPS) [dungeon] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 62.9 healing_power points (7.07 DPS) | yes | Caretaker's Cape (19530, -1.55 DPS) [rep]; Darkspear Raider's Cloak (272076, -2.15 DPS) [vendor] |
| chest | Virtuous Robe (226945) | Saving the Best for Last [quest] | 123.1 healing_power points (13.83 DPS) | yes | Field Marshal's Satin Tunic (231624, +0.00 DPS) [vendor]; Truefaith Vestments (14154, -0.44 DPS) [crafted]; Mooncloth Vest (14138, -10.53 DPS, sim-verified) [crafted] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-verified (+7.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Satin Bracers (17606, +0.00 DPS) [pvp]; Devout Bracers (16697, -1.35 DPS) [dungeon]; Bracers of Hope (22667, -7.06 DPS, sim-verified) [quest] |
| hands | Virtuous Mitts (226950) | Mokvar [vendor] | 99.1 healing_power points (11.13 DPS) | yes | Raider Handwraps (272097, -2.08 DPS) [vendor]; Desert Bloom Gloves (20717, -2.60 DPS) [quest]; Devout Gloves (16692, -12.00 DPS, sim-verified) [dungeon] |
| waist | Virtuous Belt (226948) | Just Compensation [quest] | sim-verified (+8.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Satin Sash (17609, +0.00 DPS) [pvp]; Wisdom of the Timbermaw (19047, -0.03 DPS) [crafted]; Penitent's Cinch (272394, -7.98 DPS, sim-verified) [vendor] |
| legs | Virtuous Skirt (226946) | Anthion's Parting Words [quest] | sim-verified (+7.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Captain's Satin Legguards (227125, -0.52 DPS) [pvp]; Marshal's Satin Legguards (231626, -0.85 DPS) [vendor]; Devout Skirt (16694, -7.05 DPS, sim-verified) [dungeon] |
| feet | Virtuous Sandals (226952) | Anthion's Parting Words [quest] | sim-verified (+5.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Mooncloth Boots (15802, -2.11 DPS) [crafted]; Devout Sandals (16691, -3.21 DPS) [dungeon]; Incandescent Mooncloth Boots (227862, -5.03 DPS, sim-verified) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 65.5 healing_power points (7.36 DPS) | yes | Band of the Hierophant (13096, -1.43 DPS) [world_drop]; Channeler's Ring (272406, -1.47 DPS) [vendor]; Band of Piety (22681, -1.49 DPS) [quest] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | 54.1 healing_power points (6.08 DPS) | yes | Band of the Hierophant (13096, -0.16 DPS) [world_drop]; Channeler's Ring (272406, -0.19 DPS) [vendor]; Band of Piety (22681, -0.21 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+11.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, -0.03 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.68 DPS) [quest]; Thunderbrew's Boot Flask (744, -1.33 DPS) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -0.82 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.47 DPS) [quest]; Uther's Strength (11302, -4.28 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -10.45 DPS) [dungeon]; Hand of Righteousness (7721, -13.43 DPS) [dungeon]; Hand of Edward the Odd (2243, -13.96 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 48.8 healing_power points (5.48 DPS) | yes | Jaina's Firestarter (13064, -3.10 DPS) [world_drop]; Goblin Igniter (5253, -3.57 DPS) [quest]; Cairnstone Sliver (9654, -5.05 DPS, sim-verified) [quest] |

**New at 60:** head: Virtuous Crown; neck: The Eye of Zuldazar; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Virtuous Robe; wrist: Virtuous Bracers; waist: Virtuous Belt; legs: Virtuous Skirt; feet: Virtuous Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Blessed Band of Light; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 947, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60, raid preset (human, 025005031330011520-03505003030100000-00000000000000000)

Set DPS (verified): 620.9. Weights run: 7.9s. Verify run: 27.8s. 947 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.199, intellect=2.165 ± 0.050, spirit=1.751 ± 0.044, mp5=3.077 ± 0.054, crit=0.741 ± 0.037 per rating point (14 rating = 1%, 10.375 per %), spell_haste=not significant (-2.643 ± 0.950)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 108.2 healing_power points (16.25 DPS) | yes | Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Lieutenant Commander's Satin Hood (227121, -0.71 DPS) [pvp]; Devout Crown (16693, -10.83 DPS, sim-verified) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 61.6 healing_power points (9.25 DPS) | yes | The All-Seeing Eye of Zuldazar (19594, -0.53 DPS) [quest]; Drake Tooth Necklace (21531, -1.45 DPS) [quest]; The Eye of Zuldazar (19593, -1.98 DPS, sim-verified) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 95.0 healing_power points (14.26 DPS) | yes | Lieutenant Commander's Satin Mantle (227119, -2.03 DPS) [pvp]; Field Marshal's Satin Mantle (231628, -3.26 DPS) [pvp]; Devout Mantle (16695, -11.86 DPS, sim-verified) [dungeon] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 63.6 healing_power points (9.56 DPS) | yes | Drape of Recovery (272413, -3.05 DPS) [vendor]; Caretaker's Cape (19530, -3.55 DPS) [rep]; Darkspear Raider's Cloak (272063, -7.24 DPS, sim-verified) [vendor] |
| chest | Virtuous Robe (226945) | Saving the Best for Last [quest] | 108.3 healing_power points (16.26 DPS) | yes | Field Marshal's Satin Tunic (231624, +0.00 DPS) [vendor]; Truefaith Vestments (14154, -1.09 DPS) [crafted]; Mooncloth Vest (14138, -9.74 DPS, sim-verified) [crafted] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-verified (+11.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Satin Bracers (17606, +0.00 DPS) [pvp]; Devout Bracers (16697, -1.64 DPS) [dungeon]; Bracers of Hope (22667, -11.78 DPS, sim-verified) [quest] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 82.4 healing_power points (12.38 DPS) | yes | Virtuous Mitts (226950, +0.00 DPS, sim-verified) [vendor]; Desert Bloom Gloves (20717, -1.78 DPS) [quest]; Mooncloth Gloves (18409, -2.22 DPS) [crafted] |
| waist | Virtuous Belt (226948) | Just Compensation [quest] | sim-verified (+11.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Satin Sash (17609, +0.00 DPS) [pvp]; Devout Belt (16696, -0.79 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -11.30 DPS, sim-verified) [crafted] |
| legs | Virtuous Skirt (226946) | Anthion's Parting Words [quest] | 97.2 healing_power points (14.59 DPS) | yes | Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Knight-Captain's Satin Legguards (227125, -0.03 DPS) [pvp]; Devout Skirt (16694, -5.67 DPS, sim-verified) [dungeon] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 98.8 healing_power points (14.84 DPS) | yes | Virtuous Sandals (226952, +0.00 DPS, sim-verified) [quest]; Mooncloth Boots (15802, -3.27 DPS) [crafted]; Marshal's Satin Walkers (231627, -4.63 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 69.9 healing_power points (10.49 DPS) | yes | Blessed Band of Light (272407, -3.73 DPS) [vendor]; Channeler's Ring (272406, -3.84 DPS) [vendor]; Band of the Hierophant (13096, -4.35 DPS) [world_drop] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 56.0 healing_power points (8.40 DPS) | yes | Blessed Band of Light (272407, -1.10 DPS, sim-verified) [vendor]; Channeler's Ring (272406, -1.75 DPS) [vendor]; Band of the Hierophant (13096, -2.26 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+11.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, -3.78 DPS) [world_drop]; Shard of the Splithooves (10659, -4.16 DPS) [quest]; Uther's Strength (11302, -4.23 DPS) [world_drop] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -3.15 DPS) [world_drop]; Shard of the Splithooves (10659, -3.53 DPS) [quest]; Uther's Strength (11302, -8.41 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -12.40 DPS) [dungeon]; Hand of Righteousness (7721, -18.10 DPS) [dungeon]; Hand of Edward the Odd (2243, -23.90 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 36.3 healing_power points (5.45 DPS) | yes | Jaina's Firestarter (13064, -2.71 DPS) [world_drop]; Brilliant Wand (249385, -3.17 DPS) [crafted]; Cairnstone Sliver (9654, -4.97 DPS, sim-verified) [quest] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Virtuous Robe; wrist: Virtuous Bracers; hands: Raider Handwraps; waist: Virtuous Belt; legs: Virtuous Skirt; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 947, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

## Horde

### Band 20 (undead, 025004000000000000-00000000000000000-00000000000000000)

Set DPS (verified): 36.6. Weights run: 10.7s. Verify run: 29.2s. 136 eligible items had no known source.

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

Set DPS (verified): 67.4. Weights run: 10.9s. Verify run: 32.5s. 244 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.017, intellect=2.171 ± 0.014, spirit=1.721 ± 0.012, mp5=3.882 ± 0.013, crit=0.214 ± 0.009 per rating point (14 rating = 1%, 2.993 per %), spell_haste=not significant (0.202 ± 0.091)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 43.3 healing_power points (2.75 DPS) | yes | Resilient Cap (14401, -0.80 DPS) [world_drop]; Nightsky Cowl (4039, -1.01 DPS, sim-verified) [world_drop]; Shadow Hood (4323, -1.23 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.6 healing_power points (0.99 DPS) | yes | Scorn's Icy Choker (23169, -0.16 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -0.16 DPS) [world]; Crystal Starfire Medallion (5003, -0.33 DPS, sim-verified) [world_drop] |
| shoulder | Nightsky Mantle (4718) | World drop [world_drop] | 26.0 healing_power points (1.65 DPS) | yes | Ghostly Mantle (3324, -0.12 DPS, sim-verified) [quest]; Death Speaker Mantle (6685, -0.13 DPS) [dungeon]; Talbar Mantle (10657, -0.33 DPS) [quest] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | 22.5 healing_power points (1.43 DPS) | yes | Glowing Thresher Cape (6901, -0.11 DPS) [dungeon]; Battle Healer's Cloak (19529, -0.17 DPS) [rep]; Resilient Cape (14400, -0.60 DPS) [world_drop] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 43.9 healing_power points (2.78 DPS) | yes | Pristine Gown (253961, -0.25 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -0.26 DPS) [dungeon]; Silver-thread Robe (4035, -0.83 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 22.5 healing_power points (1.43 DPS) | yes | Nightsky Wristbands (6407, -0.28 DPS, sim-verified) [world_drop]; Glowing Magical Bracelets (13106, -0.33 DPS) [world_drop]; Spidertank Oilrag (9448, -0.69 DPS) [dungeon] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 31.5 healing_power points (2.00 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.13 DPS, sim-verified) [dungeon]; Blight Gloves (279877, -0.49 DPS) [quest]; Truefaith Gloves (7049, -0.63 DPS) [crafted] |
| waist | Lilac Sash (6780) | Centaur Bounty [quest] | 24.7 healing_power points (1.57 DPS) | yes | Resilient Cord (14406, -0.31 DPS, sim-verified) [world_drop]; Pristine Sash (253925, -0.32 DPS) [crafted]; Dreamer's Belt (4829, -0.38 DPS) [vendor] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 45.8 healing_power points (2.91 DPS) | yes | Filigreed Pristine Leggings (253937, -0.61 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -0.75 DPS) [crafted]; Sacred Burial Trousers (6282, -1.23 DPS) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 42.1 healing_power points (2.67 DPS) | yes | Soggy Boots (274747, -1.25 DPS) [vendor]; Frothing Slippers (254003, -1.27 DPS) [crafted]; Acidic Walkers (9454, -1.30 DPS, sim-verified) [dungeon] |
| finger1 | The Queen's Jewel (13094) | World drop [world_drop] | 18.1 healing_power points (1.15 DPS) | yes | Runemender's Seal (273031, -0.00 DPS) [world]; Sea Giant's Toe Ring (274746, -0.01 DPS) [vendor]; Solid Gold Band (279845, -0.16 DPS) [quest] |
| finger2 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 18.0 healing_power points (1.14 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; Sea Giant's Toe Ring (274746, -0.00 DPS) [vendor]; Solid Gold Band (279845, -0.15 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-verified (67.4 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Advisor's Gnarled Staff (19569, -0.33 DPS) [pvp]; Death Speaker Scepter (2816, -1.61 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 16.5 healing_power points (1.04 DPS) | yes | Summoner's Wand (5245, -0.40 DPS, sim-verified) [world]; Lesser Mystic Wand (11289, -0.49 DPS) [crafted]; Starfaller (13063, -0.49 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Nightsky Mantle; back: Darkspear Raider's Cloak; chest: Death Speaker Robes; hands: Gloves of Old; waist: Lilac Sash; legs: Pristine Leggings; feet: Gilded Slippers; finger1: The Queen's Jewel; finger2: Electrocutioner Lagnut; main_hand: Spine of the Devourer; ranged: Wand of Mana Concentration

No-known-source sample (15 of 244, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (undead, 025005031330011520-00000000000000000-00000000000000000)

Set DPS (verified): 133.2. Weights run: 12.6s. Verify run: 80.0s. 322 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.013, intellect=2.181 ± 0.022, spirit=0.858 ± 0.037, mp5=2.390 ± 0.040, crit=0.420 ± 0.013 per rating point (14 rating = 1%, 5.876 per %), spell_haste=-12.136 ± 0.403

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 58.8 healing_power points (4.40 DPS) | yes | Whitemane's Chapeau (7720, -1.02 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -1.55 DPS) [world_drop]; Thinking Cap (2624, -1.63 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 30.4 healing_power points (2.28 DPS) | yes | Necklace of Calisea (1714, -0.74 DPS, sim-verified) [dungeon]; Darkspear Warding Pendant (272074, -1.14 DPS) [vendor]; Archmagister's Faceted Pendant (273049, -1.30 DPS) [world] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 43.3 healing_power points (3.25 DPS) | yes | Mistscape Mantle (4734, -1.13 DPS) [dungeon]; Silksand Shoulder Pads (14423, -1.26 DPS) [world_drop]; Windchaser Amice (14432, -1.50 DPS, sim-verified) [world_drop] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 27.4 healing_power points (2.05 DPS) | yes | Battle Healer's Cloak (19528, -0.09 DPS, sim-verified) [rep]; Blackforge Cape (6424, -0.49 DPS) [dungeon]; Windchaser Cloak (14430, -0.82 DPS) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 64.9 healing_power points (4.86 DPS) | yes | Red Mageweave Vest (10007, -1.92 DPS) [crafted]; Silksand Tunic (14417, -1.92 DPS) [world_drop]; Death Speaker Robes (6682, -5.36 DPS, sim-verified) [dungeon] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 19.6 healing_power points (1.47 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Mistscape Bracers (4045, -0.16 DPS) [dungeon]; Enchanted Stonecloth Bracers (4979, -0.16 DPS) [quest] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (133.2 DPS) | yes | Earthen Silk Gloves (254017, -0.92 DPS) [crafted]; Red Mageweave Gloves (10018, -1.15 DPS) [crafted]; Gilded Handwraps (254021, -1.93 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 44.6 healing_power points (3.34 DPS) | yes | Sutarn's Ring (13105, -1.32 DPS) [world_drop]; Razzeric's Customized Seatbelt (6726, -1.38 DPS) [quest]; Deathmage Sash (10771, -2.38 DPS, sim-verified) [dungeon] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-verified (133.2 DPS) | yes | Stoneweaver Leggings (9407, +0.00 DPS) [dungeon]; Filigreed Pristine Leggings (253937, +0.00 DPS) [crafted]; Pristine Leggings (253987, -1.65 DPS, sim-verified) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 38.7 healing_power points (2.90 DPS) | yes | Furen's Boots (13100, -0.23 DPS, sim-verified) [world_drop]; Boots of the Maharishi (9658, -1.04 DPS) [quest]; Kodo Rustler Boots (15697, -1.27 DPS) [quest] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.35 DPS) | yes | Ogremind Ring (1993, -0.01 DPS) [world_drop]; Mindbender Loop (5009, -0.08 DPS) [world_drop]; Welken Ring (5011, -0.08 DPS) [world_drop] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 17.8 healing_power points (1.34 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Mindbender Loop (5009, -0.06 DPS) [world_drop]; Welken Ring (5011, -0.07 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (133.2 DPS) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (133.2 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Death Speaker Scepter (2816, -2.46 DPS) [dungeon]; Hypnotic Blade (7714, -2.76 DPS) [dungeon] |
| off_hand | Aurora Sphere (7610) | Uldaman: Ancient Treasure [dungeon] | 18.7 healing_power points (1.40 DPS) | yes | Orb of Souls (249395, +0.00 DPS, sim-verified) [crafted]; Witch's Finger (16887, -0.09 DPS) [quest]; Beacon of Hope (9393, -0.09 DPS) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 15.7 healing_power points (1.17 DPS) | yes | Wand of Mana Concentration (273043, -0.16 DPS) [world]; Flash Wand (5248, -0.33 DPS) [quest]; Goblin Igniter (5253, -0.33 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Stormcloth Gloves; waist: Gilded Cord; legs: Stormcloth Pants; finger1: Sea Giant's Toe Ring; finger2: Voodoo Band; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Aurora Sphere; ranged: Jaina's Firestarter

No-known-source sample (15 of 322, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16336 Sergeant Major's Cape

### Band 50 (undead, 025005031330011520-03502000000000000-00000000000000000)

Set DPS (verified): 210.5. Weights run: 12.5s. Verify run: 43.3s. 413 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.050, intellect=2.721 ± 0.033, spirit=2.154 ± 0.046, mp5=4.514 ± 0.065, crit=0.646 ± 0.021 per rating point (14 rating = 1%, 9.050 per %), spell_haste=-9.814 ± 0.496

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 94.0 healing_power points (6.72 DPS) | yes | Blood Guard's Satin Cover (220899, -0.51 DPS) [vendor]; Papal Fez (9431, -0.83 DPS, sim-verified) [dungeon]; Imperial Red Circlet (8254, -1.30 DPS) [world_drop] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 46.7 healing_power points (3.34 DPS) | yes | Darkmoon Necklace (19303, -0.55 DPS, sim-verified) [vendor]; Glowing Eye of Mordresh (10769, -0.85 DPS) [dungeon]; Necklace of Calisea (1714, -0.90 DPS) [dungeon] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | 60.7 healing_power points (4.34 DPS) | yes | Blood Guard's Satin Pads (220901, +0.00 DPS) [vendor]; Rotgrip Mantle (17732, -0.84 DPS) [dungeon]; Inquisitor's Shawl (19507, -1.48 DPS, sim-verified) [dungeon] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 48.9 healing_power points (3.49 DPS) | yes | Imperial Red Cloak (8248, +0.00 DPS, sim-verified) [world_drop]; Battle Healer's Cloak (19527, -1.00 DPS) [rep]; Highborne Cloak (14450, -1.13 DPS) [world_drop] |
| chest | Ghostweave Vest (14141) | Tailoring [crafted] | 79.2 healing_power points (5.66 DPS) | yes | Stormcloth Vest (10020, +0.00 DPS, sim-verified) [crafted]; Stone Guard's Satin Armor (220903, -0.29 DPS) [vendor]; Grimlok's Tribal Vestments (9415, -1.17 DPS) [dungeon] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 49.8 healing_power points (3.56 DPS) | yes | Forgotten Wraps (9433, -1.23 DPS) [world_drop]; Bloodband Bracers (11469, -1.35 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -1.64 DPS, sim-verified) [quest] |
| hands | Virtuous Mitts (226950) | Mokvar [vendor] | 92.7 healing_power points (6.62 DPS) | yes | Gilded Gloves (254095, -1.38 DPS) [crafted]; Raider Handwraps (272098, -1.93 DPS, sim-verified) [vendor]; Virtuous Hands (226958, -2.25 DPS) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 63.7 healing_power points (4.55 DPS) | yes | Gilded Cord (254037, +0.00 DPS, sim-verified) [crafted]; Gilded Waistcord (254081, -0.73 DPS) [crafted]; Earthenweave Cord (254077, -1.00 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 93.8 healing_power points (6.70 DPS) | yes | Senior Designer's Pantaloons (11841, -0.78 DPS) [dungeon]; Dalewind Trousers (13008, -1.10 DPS) [world_drop]; Stone Guard's Satin Leggings (220902, -1.33 DPS) [vendor] |
| feet | Vinerot Sandals (17748) | Maraudon: Razorlash [dungeon] | 58.0 healing_power points (4.15 DPS) | yes | Gilded Sandals (254107, -0.18 DPS) [crafted]; Furen's Boots (13100, -0.32 DPS) [world_drop]; First Sergeant's Satin Boots (220900, -0.38 DPS) [vendor] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 42.2 healing_power points (3.01 DPS) | yes | Welken Ring (5011, -0.96 DPS) [world_drop]; Coldwater Ring (4550, -1.05 DPS) [quest]; Voodoo Band (1996, -1.19 DPS) [world] |
| finger2 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 31.6 healing_power points (2.26 DPS) | yes | Coldwater Ring (4550, -0.29 DPS) [quest]; Voodoo Band (1996, -0.43 DPS) [world]; Welken Ring (5011, -0.44 DPS, sim-verified) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (210.5 DPS) | yes | Evonice's Landin' Pilla (18951, -2.69 DPS, sim-verified) [quest]; Ankh of Life (1713, -2.99 DPS) [world_drop]; Alchemist's Stone (13503, -4.22 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (210.5 DPS) | yes | Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Ankh of Life (1713, -0.21 DPS, sim-verified) [world_drop]; Alchemist's Stone (13503, -0.67 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (210.5 DPS) | yes | Charstone Dirk (17710, -6.67 DPS) [dungeon]; Spellshifter Rod (9527, -9.41 DPS) [quest]; Hammer of the Northern Wind (810, -26.10 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 22.8 healing_power points (1.63 DPS) | yes | Wand of Mana Concentration (273043, +0.00 DPS, sim-verified) [world]; Flash Wand (5248, -0.39 DPS) [quest]; Goblin Igniter (5253, -0.39 DPS) [quest] |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Nethergeld Shoulders; back: Darkspear Raider's Cloak; chest: Ghostweave Vest; wrist: Nethergeld Cuffs; hands: Virtuous Mitts; waist: Dawnspire Cord; legs: Kilt of the Atal'ai Prophet; feet: Vinerot Sandals; finger1: Eye of Adaegus; finger2: Darkspear Signet; trinket1: Darkspear Voodoo Seal; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 413, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60 (undead, 025005031330011520-03505003030100000-00000000000000000)

Set DPS (verified): 351.6. Weights run: 12.0s. Verify run: 51.5s. 939 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.107, intellect=2.093 ± 0.043, spirit=2.892 ± 0.313, mp5=2.333 ± 0.149, crit=0.663 ± 0.028 per rating point (14 rating = 1%, 9.278 per %), spell_haste=-5.624 ± 0.808

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 124.1 healing_power points (13.94 DPS) | yes | Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Champion's Satin Hood (227118, -0.67 DPS) [pvp]; Devout Crown (16693, -13.63 DPS, sim-verified) [dungeon] |
| neck | The Eye of Zuldazar (19593) (or The All-Seeing Eye of Zuldazar (19594)) | The Eye of Zuldazar [quest] | 72.3 healing_power points (8.12 DPS) | yes | The All-Seeing Eye of Zuldazar (19594, +0.00 DPS) [quest]; Wavefront Necklace (20685, -1.92 DPS) [world]; Drake Tooth Necklace (21531, -2.38 DPS) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 101.9 healing_power points (11.44 DPS) | yes | Champion's Satin Mantle (227120, -0.60 DPS) [pvp]; Devout Mantle (16695, -2.35 DPS) [dungeon]; Virtuous Mantle (226951, -3.24 DPS, sim-verified) [quest] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 62.9 healing_power points (7.07 DPS) | yes | Battle Healer's Cloak (19526, -1.55 DPS) [rep]; Darkspear Raider's Cloak (272076, -2.15 DPS) [vendor] |
| chest | Virtuous Robe (226945) | Saving the Best for Last [quest] | 123.1 healing_power points (13.83 DPS) | yes | Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Truefaith Vestments (14154, -0.44 DPS) [crafted]; Mooncloth Vest (14138, -11.52 DPS, sim-verified) [crafted] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-verified (+6.4 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Satin Bracers (17619, +0.00 DPS) [pvp]; Devout Bracers (16697, -1.35 DPS) [dungeon]; Bracers of Hope (22667, -6.39 DPS, sim-verified) [quest] |
| hands | Virtuous Mitts (226950) | Mokvar [vendor] | 99.1 healing_power points (11.13 DPS) | yes | Raider Handwraps (272097, -2.08 DPS) [vendor]; Desert Bloom Gloves (20717, -2.60 DPS) [quest]; Devout Gloves (16692, -13.03 DPS, sim-verified) [dungeon] |
| waist | Virtuous Belt (226948) | Just Compensation [quest] | sim-verified (+6.5 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Satin Cinch (17621, +0.00 DPS) [pvp]; Wisdom of the Timbermaw (19047, -0.03 DPS) [crafted]; Penitent's Cinch (272394, -6.51 DPS, sim-verified) [vendor] |
| legs | Virtuous Skirt (226946) | Anthion's Parting Words [quest] | sim-verified (+5.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Legionnaire's Satin Legguards (227123, -0.52 DPS) [pvp]; General's Satin Legguards (231634, -0.85 DPS) [vendor]; Devout Skirt (16694, -5.89 DPS, sim-verified) [dungeon] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 113.8 healing_power points (12.78 DPS) | yes | Virtuous Sandals (226952, +0.00 DPS, sim-verified) [quest]; Mooncloth Boots (15802, -2.57 DPS) [crafted]; Devout Sandals (16691, -3.67 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 65.5 healing_power points (7.36 DPS) | yes | Band of the Hierophant (13096, -1.43 DPS) [world_drop]; Channeler's Ring (272406, -1.47 DPS) [vendor]; Band of Piety (22681, -1.49 DPS) [quest] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | 54.1 healing_power points (6.08 DPS) | yes | Channeler's Ring (272406, -0.19 DPS) [vendor]; Band of Piety (22681, -0.21 DPS) [quest]; Band of the Hierophant (13096, -2.54 DPS, sim-verified) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -0.03 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.68 DPS) [quest]; Serenity Field (272439, -5.48 DPS, sim-verified) [vendor] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Serenity Field (272439, -0.37 DPS, sim-verified) [vendor] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -10.45 DPS) [dungeon]; Hand of Righteousness (7721, -13.43 DPS) [dungeon]; Hand of Edward the Odd (2243, -14.74 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 48.8 healing_power points (5.48 DPS) | yes | Chillnail Splinter (10704, -3.53 DPS) [quest]; Goblin Igniter (5253, -3.57 DPS) [quest]; Jaina's Firestarter (13064, -7.19 DPS, sim-verified) [world_drop] |

**New at 60:** head: Virtuous Crown; neck: The Eye of Zuldazar; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Virtuous Robe; wrist: Virtuous Bracers; waist: Virtuous Belt; legs: Virtuous Skirt; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Blessed Band of Light; trinket2: Shard of the Splithooves; ranged: Torch of Light

No-known-source sample (15 of 939, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60, raid preset (undead, 025005031330011520-03505003030100000-00000000000000000)

Set DPS (verified): 611.7. Weights run: 7.9s. Verify run: 26.7s. 939 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.199, intellect=2.165 ± 0.050, spirit=1.751 ± 0.044, mp5=3.077 ± 0.054, crit=0.741 ± 0.037 per rating point (14 rating = 1%, 10.375 per %), spell_haste=not significant (-2.643 ± 0.950)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 108.2 healing_power points (16.25 DPS) | yes | Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Champion's Satin Hood (227118, -0.71 DPS) [pvp]; Devout Crown (16693, -15.65 DPS, sim-verified) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 61.6 healing_power points (9.25 DPS) | yes | The All-Seeing Eye of Zuldazar (19594, -0.53 DPS) [quest]; Drake Tooth Necklace (21531, -1.45 DPS) [quest]; The Eye of Zuldazar (19593, -4.61 DPS, sim-verified) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 95.0 healing_power points (14.26 DPS) | yes | Champion's Satin Mantle (227120, -2.03 DPS) [pvp]; Warlord's Satin Mantle (231631, -3.26 DPS) [pvp]; Devout Mantle (16695, -15.54 DPS, sim-verified) [dungeon] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 63.6 healing_power points (9.56 DPS) | yes | Drape of Recovery (272413, -3.05 DPS) [vendor]; Battle Healer's Cloak (19526, -3.55 DPS) [rep]; Darkspear Raider's Cloak (272063, -11.22 DPS, sim-verified) [vendor] |
| chest | Virtuous Robe (226945) | Saving the Best for Last [quest] | 108.3 healing_power points (16.26 DPS) | yes | Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Truefaith Vestments (14154, -1.09 DPS) [crafted]; Mooncloth Vest (14138, -13.48 DPS, sim-verified) [crafted] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-verified (+11.8 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Satin Bracers (17619, +0.00 DPS) [pvp]; Devout Bracers (16697, -1.64 DPS) [dungeon]; Bracers of Hope (22667, -11.75 DPS, sim-verified) [quest] |
| hands | Virtuous Mitts (226950) | Mokvar [vendor] | sim-verified (+11.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Desert Bloom Gloves (20717, -1.31 DPS) [quest]; Raider Handwraps (272098, -1.46 DPS) [vendor] |
| waist | Virtuous Belt (226948) | Just Compensation [quest] | sim-verified (+9.4 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Satin Cinch (17621, +0.00 DPS) [pvp]; Devout Belt (16696, -0.79 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -9.43 DPS, sim-verified) [crafted] |
| legs | Virtuous Skirt (226946) | Anthion's Parting Words [quest] | 97.2 healing_power points (14.59 DPS) | yes | General's Satin Legguards (231634, +0.00 DPS) [vendor]; Legionnaire's Satin Legguards (227123, -0.03 DPS) [pvp]; Devout Skirt (16694, -9.92 DPS, sim-verified) [dungeon] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 98.8 healing_power points (14.84 DPS) | yes | Virtuous Sandals (226952, +0.00 DPS, sim-verified) [quest]; Mooncloth Boots (15802, -3.27 DPS) [crafted]; General's Satin Walkers (231630, -4.63 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 69.9 healing_power points (10.49 DPS) | yes | Blessed Band of Light (272407, -3.73 DPS) [vendor]; Channeler's Ring (272406, -3.84 DPS) [vendor]; Band of the Hierophant (13096, -4.35 DPS) [world_drop] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 56.0 healing_power points (8.40 DPS) | yes | Blessed Band of Light (272407, -1.27 DPS, sim-verified) [vendor]; Channeler's Ring (272406, -1.75 DPS) [vendor]; Band of the Hierophant (13096, -2.26 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+15.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, -3.78 DPS) [world_drop]; Shard of the Splithooves (10659, -4.16 DPS) [quest]; Uther's Strength (11302, -4.23 DPS) [world_drop] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -3.15 DPS) [world_drop]; Shard of the Splithooves (10659, -3.53 DPS) [quest]; Uther's Strength (11302, -7.96 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -12.40 DPS) [dungeon]; Hand of Righteousness (7721, -18.10 DPS) [dungeon]; Hand of Edward the Odd (2243, -24.51 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 36.3 healing_power points (5.45 DPS) | yes | Brilliant Wand (249385, -3.17 DPS) [crafted]; Wand of Mana Concentration (273043, -3.22 DPS) [world]; Jaina's Firestarter (13064, -6.36 DPS, sim-verified) [world_drop] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Virtuous Robe; wrist: Virtuous Bracers; waist: Virtuous Belt; legs: Virtuous Skirt; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 939, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

