# Leveling BiS: Holy

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-03503000000000000-00000000000000000)

Set DPS (verified): 40.9. Weights run: 10.3s. Verify run: 28.4s. 146 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.006, intellect=1.334 ± 0.006, spirit=0.605 ± 0.003, mp5=1.556 ± 0.011, crit=0.070 ± 0.003 per rating point (14 rating = 1%, 0.982 per %), spell_haste=not significant (-0.005 ± 0.014)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.63 DPS) | yes | Shadow Goggles (4373, -1.12 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.41 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 2.4 healing_power points (0.22 DPS) | yes | Scholarly Pendant (277203, +0.00 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.0 healing_power points (1.09 DPS) | yes | Slime-encrusted Pads (6461, -0.67 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -0.94 DPS, sim-verified) [crafted] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | 10.2 healing_power points (0.93 DPS) | yes | Pearl-clasped Cloak (5542, -0.56 DPS) [crafted]; Seer's Cape (6378, -0.57 DPS) [dungeon]; Sanguine Cape (14376, -0.83 DPS, sim-verified) [world_drop] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 21.5 healing_power points (1.95 DPS) | yes | Robe of the Moccasin (6465, -0.58 DPS) [dungeon]; Bloody Apron (6226, -0.77 DPS) [dungeon]; Corsair's Overshirt (5202, -0.82 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 11.3 healing_power points (1.03 DPS) | yes | Repurposed Hair Band (281256, -0.68 DPS) [quest]; Mystic's Bracelets (14366, -0.79 DPS) [world_drop]; Bright Bracers (3647, -1.14 DPS, sim-verified) [world_drop] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 15.0 healing_power points (1.36 DPS) | yes | Tomb Robber's Gloves (280096, -0.63 DPS) [quest]; Bright Gloves (3066, -0.66 DPS) [world_drop]; Magefist Gloves (12977, -1.04 DPS, sim-verified) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.3 healing_power points (1.48 DPS) | yes | Keller's Girdle (2911, -0.51 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.58 DPS, sim-verified) [crafted]; Wise Man's Belt (4786, -0.88 DPS) [vendor] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 28.4 healing_power points (2.58 DPS) | yes | Abomination Skin Leggings (23173, -1.61 DPS) [dungeon]; Filigreed Silky Leggings (253939, -1.63 DPS) [crafted]; Darkweave Breeches (12987, -1.99 DPS, sim-verified) [world_drop] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 13.0 healing_power points (1.18 DPS) | yes | Spidersilk Boots (4320, -0.69 DPS) [crafted]; Sanguine Sandals (14374, -0.69 DPS) [world_drop]; Kimbra Boots (6191, -0.92 DPS, sim-verified) [quest] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 8.0 healing_power points (0.73 DPS) | yes | Band of Purification (12996, -0.40 DPS) [world_drop]; Lorekeeper's Ring (20431, -0.44 DPS) [rep]; Deep Fathom Ring (6463, -0.45 DPS) [dungeon] |
| finger2 | Volcanic Rock Ring (12053) | World drop [world_drop] | 4.0 healing_power points (0.36 DPS) | yes | Lorekeeper's Ring (20431, -0.08 DPS) [rep]; Deep Fathom Ring (6463, -0.09 DPS) [dungeon]; Band of Purification (12996, -0.21 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 13.3 healing_power points (1.21 DPS) | yes | Lesser Staff of the Spire (1300, -0.48 DPS) [world]; Staff of Westfall (2042, -0.53 DPS, sim-verified) [quest]; Witching Stave (1484, -0.60 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Flaring Baton (5326) (or Moonstone Wand (15204)) | The Escape [quest] | 2.7 healing_power points (0.24 DPS) | yes | Moonstone Wand (15204, +0.00 DPS) [quest]; Sable Wand (7607, -0.08 DPS) [quest]; Elven Wand (5604, -0.12 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Tarnished Locket; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Volcanic Rock Ring; main_hand: Twisted Chanter's Staff; ranged: Flaring Baton

No-known-source sample (15 of 146, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes; 16607 Acolyte's Sacrificial Robes; 18851 Insignia of the Horde; 20434 Lorekeeper's Staff

### Band 30 (gnome, 000000000000000000-03505003030110000-00000000000000000)

Set DPS (verified): 65.1. Weights run: 10.3s. Verify run: 31.7s. 258 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.013, intellect=1.413 ± 0.007, spirit=0.841 ± 0.004, mp5=1.628 ± 0.015, crit=0.098 ± 0.005 per rating point (14 rating = 1%, 1.379 per %), spell_haste=0.351 ± 0.047

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 38.0 healing_power points (3.46 DPS) | yes | Nightsky Cowl (4039, -1.54 DPS) [world_drop]; Filigreed Pristine Circlet (253975, -1.73 DPS, sim-verified) [crafted]; Resilient Cap (14401, -1.74 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 9.0 healing_power points (0.82 DPS) | yes | Scorn's Icy Choker (23169, -0.05 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -0.05 DPS) [world]; Crystal Starfire Medallion (5003, -0.08 DPS) [world_drop] |
| shoulder | Mantle of Honor (3560) | Bride of the Embalmer [quest] | 15.8 healing_power points (1.43 DPS) | yes | Death Speaker Mantle (6685, +0.00 DPS, sim-verified) [dungeon]; Nightsky Mantle (4718, -0.02 DPS) [world_drop]; Faerie Mantle (5820, -0.20 DPS) [quest] |
| back | Caretaker's Cape (19533) | Silverwing Sentinels [rep] | 16.4 healing_power points (1.49 DPS) | yes | Prelacy Cape (7004, -0.03 DPS) [quest]; Darkspear Raider's Cloak (272078, -0.23 DPS) [vendor]; Glowing Thresher Cape (6901, -0.24 DPS) [dungeon] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 35.5 healing_power points (3.23 DPS) | yes | Pristine Gown (253961, -0.34 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -1.08 DPS) [dungeon]; Robes of Arugal (6324, -1.14 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 11.9 healing_power points (1.09 DPS) | yes | Nightsky Wristbands (6407, -0.09 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.26 DPS, sim-verified) [world_drop]; Stonecloth Bindings (14416, -0.44 DPS) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 19.2 healing_power points (1.75 DPS) | yes | Gloves of Old (9395, -0.19 DPS) [world_drop]; Hotshot Pilot's Gloves (9491, -0.34 DPS) [dungeon]; Pristine Gloves (253913, -0.36 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.7 healing_power points (1.51 DPS) | yes | Novice Ardent's Sash (253887, -0.39 DPS, sim-verified) [crafted]; Resilient Cord (14406, -0.44 DPS) [world_drop]; Dreamer's Belt (4829, -0.46 DPS) [vendor] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 36.1 healing_power points (3.28 DPS) | yes | Filigreed Pristine Leggings (253937, -0.65 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -1.08 DPS) [crafted]; Necromancer Leggings (2277, -1.87 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 33.3 healing_power points (3.02 DPS) | yes | Nimbus Boots (6998, -1.67 DPS, sim-verified) [quest]; Acidic Walkers (9454, -1.69 DPS) [dungeon]; Frothing Slippers (254003, -1.82 DPS) [crafted] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.64 DPS) | yes | Runemender's Seal (273031, -0.56 DPS) [world]; The Queen's Jewel (13094, -0.77 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.87 DPS) [dungeon] |
| finger2 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 11.9 healing_power points (1.08 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; The Queen's Jewel (13094, -0.21 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.31 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-verified (65.1 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -0.77 DPS) [quest]; Death Speaker Scepter (2816, -2.13 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 8.9 healing_power points (0.81 DPS) | yes | Lesser Mystic Wand (11289, -0.30 DPS) [crafted]; Starfaller (13063, -0.30 DPS) [world_drop]; Summoner's Wand (5245, -0.57 DPS, sim-verified) [world] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Mantle of Honor; back: Caretaker's Cape; chest: Death Speaker Robes; hands: Truefaith Gloves; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Sea Giant's Toe Ring; finger2: Electrocutioner Lagnut; main_hand: Spine of the Devourer; ranged: Wand of Mana Concentration

No-known-source sample (15 of 258, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 40 (gnome, 000000000000000000-03505003030121431-00000000000000000)

Set DPS (verified): 86.0. Weights run: 10.5s. Verify run: 68.6s. 339 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.075, intellect=4.127 ± 0.033, spirit=0.660 ± 0.020, mp5=0.606 ± 0.023, crit=0.213 ± 0.017 per rating point (14 rating = 1%, 2.984 per %), spell_haste=not significant (0.974 ± 0.519)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 90.1 healing_power points (4.49 DPS) | yes | Whitemane's Chapeau (7720, -1.39 DPS) [dungeon]; Craftsman's Monocle (4393, -1.41 DPS) [crafted]; Thinking Cap (2624, -1.74 DPS, sim-verified) [world] |
| neck | Necklace of Calisea (1714) | Zul'Farrak: Witch Doctor's Chest [dungeon] | sim-verified (86.0 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS, sim-verified) [dungeon]; Darkspear Warding Pendant (272074, -0.23 DPS) [vendor]; Archmagister's Faceted Pendant (273049, -0.44 DPS) [world] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 68.6 healing_power points (3.42 DPS) | yes | Mistscape Mantle (4734, -1.00 DPS) [dungeon]; Silksand Shoulder Pads (14423, -1.06 DPS) [world_drop]; Windchaser Amice (14432, -1.69 DPS, sim-verified) [world_drop] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 48.0 healing_power points (2.40 DPS) | yes | Blackforge Cape (6424, -0.44 DPS, sim-verified) [dungeon]; Glyphed Cloak (4732, -0.96 DPS) [world_drop]; Sergeant Major's Cape (16336, -0.96 DPS) [pvp] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-verified (86.0 DPS) | yes | Red Mageweave Vest (10007, +0.00 DPS) [crafted]; Silksand Wraps (14425, +0.00 DPS) [world_drop]; Silksand Tunic (14417, -5.04 DPS, sim-verified) [world_drop] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 37.1 healing_power points (1.85 DPS) | yes | Mistscape Bracers (4045, +0.00 DPS, sim-verified) [dungeon]; Aurora Bracers (4043, -0.21 DPS) [world_drop]; Enchanted Stonecloth Bracers (4979, -0.21 DPS) [quest] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | 60.5 healing_power points (3.02 DPS) | yes | Red Mageweave Gloves (10018, -0.96 DPS) [crafted]; Mistscape Gloves (6428, -1.17 DPS) [dungeon]; Gilded Handwraps (254021, -1.23 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | sim-verified (86.0 DPS) | yes | Razzeric's Customized Seatbelt (6726, -0.47 DPS) [quest]; Teacher's Sash (10747, -0.47 DPS) [quest]; Deathmage Sash (10771, -0.65 DPS, sim-verified) [dungeon] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-verified (86.0 DPS) | yes | Crimson Silk Pantaloons (7062, +0.00 DPS) [crafted]; Stoneweaver Leggings (9407, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, +0.00 DPS) [crafted] |
| feet | Furen's Boots (13100) | World drop [world_drop] | sim-verified (86.0 DPS) | yes | Gilded Slippers (254001, +0.00 DPS, sim-verified) [crafted]; Kodo Rustler Boots (15697, -0.16 DPS) [quest]; Acidic Walkers (9454, -0.19 DPS) [dungeon] |
| finger1 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 30.9 healing_power points (1.54 DPS) | yes | Mindbender Loop (5009, -0.03 DPS) [world_drop]; Welken Ring (5011, -0.28 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.30 DPS) [dungeon] |
| finger2 | Ogremind Ring (1993) | World drop [world_drop] | 30.9 healing_power points (1.54 DPS) | yes | Welken Ring (5011, -0.28 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.30 DPS) [dungeon]; Mindbender Loop (5009, -0.34 DPS, sim-verified) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (86.0 DPS) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (86.0 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Hypnotic Blade (7714, -1.01 DPS) [dungeon]; Illusionary Rod (7713, -1.29 DPS) [dungeon] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 53.3 healing_power points (2.66 DPS) | yes | Skull of Impending Doom (4984, -1.22 DPS) [quest]; Nightsky Orb (15929, -1.32 DPS) [world_drop]; Aurora Sphere (7610, -1.88 DPS, sim-verified) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 26.7 healing_power points (1.33 DPS) | yes | Goblin Igniter (5253, +0.00 DPS, sim-verified) [quest]; Flash Wand (5248, -0.41 DPS) [quest]; Wand of Mana Concentration (273043, -0.45 DPS) [world] |

**New at 40:** head: Papal Fez; neck: Necklace of Calisea; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Stormcloth Gloves; waist: Gilded Cord; legs: Stormcloth Pants; feet: Furen's Boots; finger1: Voodoo Band; finger2: Ogremind Ring; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Orb of Lorica; ranged: Jaina's Firestarter

No-known-source sample (15 of 339, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 50 (gnome, 025003000000000000-03505003030121431-00000000000000000)

Set DPS (verified): 150.4. Weights run: 10.7s. Verify run: 36.2s. 433 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.100, intellect=3.827 ± 0.050, spirit=0.984 ± 0.039, mp5=1.330 ± 0.040, crit=0.303 ± 0.024 per rating point (14 rating = 1%, 4.245 per %), spell_haste=not significant (-2.915 ± 0.745)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 91.8 healing_power points (5.04 DPS) | yes | Knight-Lieutenant's Satin Cover (220896, -0.35 DPS) [vendor]; Eye of Theradras (17715, -0.51 DPS) [dungeon]; Papal Fez (9431, -1.46 DPS, sim-verified) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 57.5 healing_power points (3.16 DPS) | yes | Glowing Eye of Mordresh (10769, -0.70 DPS, sim-verified) [dungeon]; Darkspear Warding Pendant (272073, -1.27 DPS) [vendor]; Necklace of Calisea (1714, -1.31 DPS) [dungeon] |
| shoulder | Knight-Lieutenant's Satin Pads (220894) | Captain Dirgehammer [vendor] | sim-verified (+4.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Inquisitor's Shawl (19507, -0.05 DPS) [dungeon]; Nethergeld Shoulders (254049, -0.17 DPS) [crafted]; Rotgrip Mantle (17732, -4.09 DPS, sim-verified) [dungeon] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 58.5 healing_power points (3.21 DPS) | yes | Highborne Cloak (14450, -1.10 DPS) [world_drop]; Keeper's Cloak (14665, -1.11 DPS) [world_drop]; Imperial Red Cloak (8248, -1.27 DPS, sim-verified) [world_drop] |
| chest | Knight's Satin Armor (220892) | Captain Dirgehammer [vendor] | 78.9 healing_power points (4.33 DPS) | yes | Acumen Robes (17775, -0.13 DPS) [quest]; Red Mageweave Vest (10007, -0.55 DPS) [crafted]; Hibernal Robe (8113, -4.52 DPS, sim-verified) [world_drop] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 51.7 healing_power points (2.84 DPS) | yes | Forgotten Wraps (9433, -0.32 DPS) [world_drop]; Imperial Red Bracers (8247, -0.53 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -2.84 DPS, sim-verified) [quest] |
| hands | Virtuous Mitts (226950) | Mokvar [vendor] | sim-verified (+4.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Gilded Gloves (254095, -0.48 DPS) [crafted]; Virtuous Hands (226958, -1.22 DPS) [vendor]; Raider Handwraps (272098, -4.47 DPS, sim-verified) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 84.7 healing_power points (4.65 DPS) | yes | Gilded Waistcord (254081, -1.17 DPS) [crafted]; Gilded Cord (254037, -1.44 DPS) [crafted]; Serenity Belt (13144, -2.56 DPS, sim-verified) [world_drop] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 92.6 healing_power points (5.08 DPS) | yes | Knight's Satin Leggings (220893, -0.75 DPS) [vendor]; Senior Designer's Pantaloons (11841, -1.27 DPS) [dungeon]; Venomshroud Leggings (14444, -1.28 DPS) [world_drop] |
| feet | Sergeant Major's Satin Boots (220895) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 65.9 healing_power points (3.62 DPS) | yes | Gilded Sandals (254107, +0.00 DPS, sim-verified) [crafted]; Vinerot Sandals (17748, -0.53 DPS) [dungeon]; Gilded Slippers (254001, -0.83 DPS) [crafted] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 34.8 healing_power points (1.91 DPS) | yes | Ogremind Ring (1993, -0.28 DPS) [world_drop]; Mindbender Loop (5009, -0.33 DPS) [world_drop]; Welken Ring (5011, -0.48 DPS) [world_drop] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 29.7 healing_power points (1.63 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Mindbender Loop (5009, -0.05 DPS) [world_drop]; Welken Ring (5011, -0.20 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+5.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, -0.45 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.56 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.66 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.45 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.56 DPS) [quest]; Ankh of Life (1713, -0.64 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -5.99 DPS) [dungeon]; Spellshifter Rod (9527, -7.02 DPS) [quest]; Hammer of the Northern Wind (810, -29.99 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 25.9 healing_power points (1.42 DPS) | yes | Cairnstone Sliver (9654, -0.33 DPS, sim-verified) [quest]; Flash Wand (5248, -0.42 DPS) [quest]; Goblin Igniter (5253, -0.42 DPS) [quest] |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Knight-Lieutenant's Satin Pads; back: Darkspear Raider's Cloak; chest: Knight's Satin Armor; wrist: Nethergeld Cuffs; hands: Virtuous Mitts; waist: Dawnspire Cord; legs: Kilt of the Atal'ai Prophet; feet: Sergeant Major's Satin Boots; finger1: Eye of Adaegus; finger2: Voodoo Band; trinket1: Darkspear Voodoo Seal; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 433, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60 (gnome, 025003031330000000-03505003030121431-00000000000000000)

Set DPS (verified): 425.1. Weights run: 9.6s. Verify run: 65.4s. 947 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.156, intellect=1.688 ± 0.065, spirit=0.435 ± 0.082, mp5=0.712 ± 0.061, crit=0.205 ± 0.029 per rating point (14 rating = 1%, 2.870 per %), spell_haste=-13.005 ± 1.372

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 71.5 healing_power points (12.07 DPS) | yes | Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Devout Crown (16693, -16.03 DPS, sim-verified) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | sim-verified (+4.2 DPS vs the runner-up, not corroborated against the finished set) | yes | The Eye of Zuldazar (19593, -0.45 DPS) [quest]; The All-Seeing Eye of Zuldazar (19594, -0.45 DPS) [quest]; Drake Tooth Necklace (21531, -4.18 DPS, sim-verified) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 67.6 healing_power points (11.41 DPS) | yes | Field Marshal's Satin Mantle (231628, -0.65 DPS) [pvp]; Lieutenant Commander's Satin Mantle (227119, -1.73 DPS) [pvp]; Devout Mantle (16695, -11.63 DPS, sim-verified) [dungeon] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 58.9 healing_power points (9.94 DPS) | yes | Drape of Recovery (272413, -4.06 DPS, sim-verified) [vendor]; Darkspear Raider's Cloak (272063, -4.94 DPS) [vendor]; Caretaker's Cape (19530, -4.96 DPS) [rep] |
| chest | Virtuous Robe (226945) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Satin Tunic (231624, +0.00 DPS) [vendor]; Knight-Captain's Satin Tunic (227122, -0.34 DPS) [pvp]; Truefaith Vestments (14154, -10.94 DPS, sim-verified) [crafted] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Satin Bracers (17606, -0.26 DPS) [pvp]; Nethergeld Cuffs (254061, -0.67 DPS) [crafted]; Bracers of Hope (22667, -12.52 DPS, sim-verified) [quest] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 69.6 healing_power points (11.74 DPS) | yes | Desert Bloom Gloves (20717, +0.00 DPS, sim-verified) [quest]; Mooncloth Gloves (18409, -1.61 DPS) [crafted]; Marshal's Satin Handwraps (231623, -3.14 DPS) [vendor] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 58.5 healing_power points (9.88 DPS) | yes | Virtuous Belt (226948, -1.68 DPS) [quest]; Marshal's Satin Sash (17609, -2.23 DPS) [pvp]; Devout Belt (16696, -4.15 DPS, sim-verified) [dungeon] |
| legs | Virtuous Skirt (226946) | Anthion's Parting Words [quest] | 65.8 healing_power points (11.11 DPS) | yes | Knight-Captain's Satin Legguards (227125, +0.00 DPS) [pvp]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Devout Skirt (16694, -14.29 DPS, sim-verified) [dungeon] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 73.7 healing_power points (12.44 DPS) | yes | Marshal's Satin Walkers (231627, -2.42 DPS) [vendor]; Virtuous Sandals (226952, -3.65 DPS) [quest]; Mooncloth Boots (15802, -6.65 DPS, sim-verified) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 53.8 healing_power points (9.07 DPS) | yes | Blessed Band of Light (272407, -3.25 DPS) [vendor]; Channeler's Ring (272406, -4.50 DPS) [vendor]; Band of the Hierophant (13096, -5.42 DPS) [world_drop] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 41.7 healing_power points (7.04 DPS) | yes | Blessed Band of Light (272407, -1.44 DPS, sim-verified) [vendor]; Channeler's Ring (272406, -2.47 DPS) [vendor]; Band of the Hierophant (13096, -3.39 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+4.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Ankh of Life (1713, -0.92 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.07 DPS) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (+9.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -4.05 DPS) [world_drop]; Ankh of Life (1713, -6.21 DPS) [world_drop]; Shard of the Splithooves (10659, -9.95 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -11.29 DPS) [dungeon]; Hand of Righteousness (7721, -17.79 DPS) [dungeon]; Hand of Edward the Odd (2243, -26.74 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 21.8 healing_power points (3.68 DPS) | yes | Jaina's Firestarter (13064, -1.75 DPS) [world_drop]; Cairnstone Sliver (9654, -1.88 DPS) [quest]; Brilliant Wand (249385, -6.04 DPS, sim-verified) [crafted] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Virtuous Robe; wrist: Virtuous Bracers; hands: Raider Handwraps; waist: Wisdom of the Timbermaw; legs: Virtuous Skirt; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 947, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60, raid preset (gnome, 025003031330000000-03505003030121431-00000000000000000)

Set DPS (verified): 771.1. Weights run: 6.2s. Verify run: 13.5s. 947 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.231, intellect=0.960 ± 0.032, spirit=0.829 ± 0.032, mp5=1.644 ± 0.037, crit=0.312 ± 0.025 per rating point (14 rating = 1%, 4.363 per %), spell_haste=not significant (-0.884 ± 0.705)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 66.9 healing_power points (25.10 DPS) | yes | Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -1.32 DPS) [crafted] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 42.9 healing_power points (16.09 DPS) | yes | Drake Tooth Necklace (21531, -0.47 DPS, sim-verified) [quest]; The Eye of Zuldazar (19593, -2.42 DPS) [quest]; The All-Seeing Eye of Zuldazar (19594, -2.42 DPS) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 77.1 healing_power points (28.93 DPS) | yes | Mooncloth Shoulders (14139, -1.81 DPS, sim-verified) [crafted]; Lieutenant Commander's Satin Mantle (227119, -8.63 DPS) [pvp]; Field Marshal's Satin Mantle (231628, -10.48 DPS) [pvp] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 51.6 healing_power points (19.35 DPS) | yes | Drape of Recovery (272413, -1.00 DPS, sim-verified) [vendor]; Caretaker's Cape (19530, -7.11 DPS) [rep] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 86.3 healing_power points (32.34 DPS) | yes | Virtuous Robe (226945, -0.61 DPS, sim-verified) [quest]; Field Marshal's Satin Tunic (231624, -4.10 DPS) [vendor]; Knight-Captain's Satin Tunic (227122, -7.60 DPS) [pvp] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 38.5 healing_power points (14.44 DPS) | yes | Virtuous Bracers (226949, -1.03 DPS) [quest]; Nethergeld Cuffs (254061, -2.86 DPS) [crafted]; Earthenweave Cuffs (254131, -3.08 DPS) [crafted] |
| hands | Desert Bloom Gloves (20717) | Armaments of War [quest] | 59.9 healing_power points (22.48 DPS) | yes | Raider Handwraps (272097, -1.09 DPS, sim-verified) [vendor]; Virtuous Mitts (226950, -3.93 DPS) [vendor]; Mooncloth Gloves (18409, -4.34 DPS) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 51.9 healing_power points (19.46 DPS) | yes | Virtuous Belt (226948, -2.32 DPS) [quest]; Penitent's Cinch (272394, -3.38 DPS) [vendor]; Devout Belt (16696, -5.34 DPS) [dungeon] |
| legs | Virtuous Skirt (226946) | Anthion's Parting Words [quest] | 61.3 healing_power points (22.98 DPS) | yes | Knight-Captain's Satin Legguards (227125, +0.00 DPS) [pvp]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Devout Skirt (16694, -0.43 DPS, sim-verified) [dungeon] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 69.0 healing_power points (25.89 DPS) | yes | Virtuous Sandals (226952, -0.63 DPS, sim-verified) [quest]; Mooncloth Boots (15802, -7.81 DPS) [crafted]; Marshal's Satin Walkers (231627, -8.53 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 51.9 healing_power points (19.45 DPS) | yes | Blessed Band of Light (272407, -5.34 DPS) [vendor]; Channeler's Ring (272406, -11.92 DPS) [vendor]; Band of the Hierophant (13096, -12.43 DPS) [world_drop] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 38.2 healing_power points (14.31 DPS) | yes | Blessed Band of Light (272407, -0.20 DPS) [vendor]; Channeler's Ring (272406, -6.79 DPS) [vendor]; Band of the Hierophant (13096, -7.30 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (771.1 DPS) | yes | Uther's Strength (11302, -2.50 DPS) [world_drop]; Ankh of Life (1713, -5.52 DPS) [world_drop]; Shard of the Splithooves (10659, -5.55 DPS) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (771.1 DPS) | yes | Darkmoon Card: Blue Dragon (19288, -1.77 DPS, sim-verified) [quest]; Uther's Strength (11302, -9.00 DPS) [world_drop]; Ankh of Life (1713, -12.02 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (771.1 DPS) | yes | Hand of Edward the Odd (2243, -8.92 DPS, sim-verified) [world_drop]; Charstone Dirk (17710, -20.79 DPS) [dungeon]; Hand of Righteousness (7721, -31.75 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 26.1 healing_power points (9.79 DPS) | yes | Cairnstone Sliver (9654, -0.86 DPS, sim-verified) [quest]; Captain Rackmore's Tiller (16789, -6.61 DPS) [quest]; Jaina's Firestarter (13064, -6.70 DPS) [world_drop] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Bracers of Hope; hands: Desert Bloom Gloves; waist: Wisdom of the Timbermaw; legs: Virtuous Skirt; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 947, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

## Horde

### Band 20 (troll, 000000000000000000-03503000000000000-00000000000000000)

Set DPS (verified): 38.8. Weights run: 10.3s. Verify run: 28.5s. 136 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.006, intellect=1.334 ± 0.006, spirit=0.605 ± 0.003, mp5=1.556 ± 0.011, crit=0.070 ± 0.003 per rating point (14 rating = 1%, 0.982 per %), spell_haste=not significant (-0.005 ± 0.014)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.63 DPS) | yes | Shadow Goggles (4373, -0.84 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.41 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 1.8 healing_power points (0.16 DPS) | yes | Scholarly Pendant (277203, -0.07 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.0 healing_power points (1.09 DPS) | yes | Slime-encrusted Pads (6461, -0.67 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -0.91 DPS, sim-verified) [crafted] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | 10.2 healing_power points (0.93 DPS) | yes | Pearl-clasped Cloak (5542, -0.56 DPS) [crafted]; Seer's Cape (6378, -0.57 DPS) [dungeon]; Sanguine Cape (14376, -0.66 DPS, sim-verified) [world_drop] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 21.5 healing_power points (1.95 DPS) | yes | Corsair's Overshirt (5202, -0.36 DPS, sim-verified) [dungeon]; Robe of the Moccasin (6465, -0.58 DPS) [dungeon]; Bloody Apron (6226, -0.77 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 11.3 healing_power points (1.03 DPS) | yes | Bright Bracers (3647, -0.54 DPS) [world_drop]; Crystalline Cuffs (14148, -0.62 DPS) [dungeon]; Featherbead Bracers (15452, -1.02 DPS, sim-verified) [quest] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 15.0 healing_power points (1.36 DPS) | yes | Blight Gloves (279877, -0.11 DPS, sim-verified) [quest]; Magefist Gloves (12977, -0.54 DPS) [world_drop]; Bright Gloves (3066, -0.66 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.3 healing_power points (1.48 DPS) | yes | Novice Ardent's Sash (253887, -0.44 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.51 DPS) [world_drop]; Wise Man's Belt (4786, -0.88 DPS) [vendor] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 28.4 healing_power points (2.58 DPS) | yes | Abomination Skin Leggings (23173, -1.61 DPS) [dungeon]; Filigreed Silky Leggings (253939, -1.63 DPS) [crafted]; Darkweave Breeches (12987, -1.78 DPS, sim-verified) [world_drop] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 13.0 healing_power points (1.18 DPS) | yes | Spidersilk Boots (4320, -0.69 DPS) [crafted]; Bluegill Sandals (1560, -0.72 DPS) [world]; Sanguine Sandals (14374, -1.02 DPS, sim-verified) [world_drop] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 8.0 healing_power points (0.73 DPS) | yes | Volcanic Rock Ring (12053, -0.36 DPS) [world_drop]; Band of Purification (12996, -0.40 DPS) [world_drop]; Advisor's Ring (20426, -0.44 DPS) [rep] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 6.7 healing_power points (0.60 DPS) | yes | Band of Purification (12996, -0.28 DPS) [world_drop]; Advisor's Ring (20426, -0.32 DPS) [rep]; Volcanic Rock Ring (12053, -0.45 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 13.3 healing_power points (1.21 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Staff of Orgrimmar (15444, -0.17 DPS) [quest]; Advisor's Gnarled Staff (20425, -0.30 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Decay (5252) (or Flaring Baton (5326)) | Beren's Peril [quest] | 2.7 healing_power points (0.24 DPS) | yes | Flaring Baton (5326, +0.00 DPS) [quest]; Wisesight Wand (286750, -0.12 DPS) [world] |

**New at 20:** head: Pristine Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Gnarled Necromancer's Staff; ranged: Wand of Decay

No-known-source sample (15 of 136, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 209613 Insignia of the Alliance; 241089 Scarlet Dagger; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff; 263007 Skyseer's Vest; 263402 Sash of Sorrow

### Band 30 (troll, 000000000000000000-03505003030110000-00000000000000000)

Set DPS (verified): 63.0. Weights run: 10.3s. Verify run: 32.2s. 244 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.013, intellect=1.413 ± 0.007, spirit=0.841 ± 0.004, mp5=1.628 ± 0.015, crit=0.098 ± 0.005 per rating point (14 rating = 1%, 1.379 per %), spell_haste=0.351 ± 0.047

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 38.0 healing_power points (3.46 DPS) | yes | Filigreed Pristine Circlet (253975, -1.47 DPS, sim-verified) [crafted]; Nightsky Cowl (4039, -1.54 DPS) [world_drop]; Resilient Cap (14401, -1.74 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 9.0 healing_power points (0.82 DPS) | yes | Scorn's Icy Choker (23169, -0.05 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -0.05 DPS) [world]; Crystal Starfire Medallion (5003, -0.08 DPS) [world_drop] |
| shoulder | Ghostly Mantle (3324) | Deathstalkers in Shadowfang [quest] | 16.6 healing_power points (1.51 DPS) | yes | Nightsky Mantle (4718, -0.10 DPS) [world_drop]; Death Speaker Mantle (6685, -0.27 DPS, sim-verified) [dungeon]; Violet Sorcerer's Mantle (273051, -0.35 DPS) [world] |
| back | Battle Healer's Cloak (19529) | Warsong Outriders [rep] | 16.4 healing_power points (1.49 DPS) | yes | Glowing Thresher Cape (6901, -0.24 DPS) [dungeon]; Resilient Cape (14400, -0.72 DPS) [world_drop]; Darkspear Raider's Cloak (272078, -1.01 DPS, sim-verified) [vendor] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 35.5 healing_power points (3.23 DPS) | yes | Pristine Gown (253961, -0.63 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -1.08 DPS) [dungeon]; Robes of Arugal (6324, -1.14 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 11.9 healing_power points (1.09 DPS) | yes | Nightsky Wristbands (6407, -0.09 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.14 DPS, sim-verified) [world_drop]; Featherbead Bracers (15452, -0.44 DPS) [quest] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 19.2 healing_power points (1.75 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.34 DPS) [dungeon]; Pristine Gloves (253913, -0.36 DPS) [crafted]; Gloves of Old (9395, -0.43 DPS, sim-verified) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.7 healing_power points (1.51 DPS) | yes | Lilac Sash (6780, -0.32 DPS, sim-verified) [quest]; Novice Ardent's Sash (253887, -0.35 DPS) [crafted]; Resilient Cord (14406, -0.44 DPS) [world_drop] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 36.1 healing_power points (3.28 DPS) | yes | Filigreed Pristine Leggings (253937, -0.78 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -1.08 DPS) [crafted]; Necromancer Leggings (2277, -1.87 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 33.3 healing_power points (3.02 DPS) | yes | Frothing Slippers (254003, -1.82 DPS) [crafted]; Black Slippers (254007, -1.82 DPS) [crafted]; Acidic Walkers (9454, -2.18 DPS, sim-verified) [dungeon] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.64 DPS) | yes | Runemender's Seal (273031, -0.56 DPS) [world]; Solid Gold Band (279845, -0.62 DPS) [quest]; The Queen's Jewel (13094, -0.77 DPS) [world_drop] |
| finger2 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 11.9 healing_power points (1.08 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; Solid Gold Band (279845, -0.06 DPS) [quest]; The Queen's Jewel (13094, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-verified (63.0 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Advisor's Gnarled Staff (19569, -1.78 DPS) [pvp]; Death Speaker Scepter (2816, -2.26 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 8.9 healing_power points (0.81 DPS) | yes | Lesser Mystic Wand (11289, -0.30 DPS) [crafted]; Starfaller (13063, -0.30 DPS) [world_drop]; Summoner's Wand (5245, -0.31 DPS, sim-verified) [world] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Ghostly Mantle; back: Battle Healer's Cloak; chest: Death Speaker Robes; hands: Truefaith Gloves; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Sea Giant's Toe Ring; finger2: Electrocutioner Lagnut; main_hand: Spine of the Devourer; ranged: Wand of Mana Concentration

No-known-source sample (15 of 244, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (troll, 000000000000000000-03505003030121431-00000000000000000)

Set DPS (verified): 87.5. Weights run: 10.5s. Verify run: 42.1s. 322 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.075, intellect=4.127 ± 0.033, spirit=0.660 ± 0.020, mp5=0.606 ± 0.023, crit=0.213 ± 0.017 per rating point (14 rating = 1%, 2.984 per %), spell_haste=not significant (0.974 ± 0.519)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 90.1 healing_power points (4.49 DPS) | yes | Thinking Cap (2624, -0.99 DPS) [world]; Whitemane's Chapeau (7720, -1.39 DPS) [dungeon]; Craftsman's Monocle (4393, -1.41 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (87.5 DPS) | yes | Necklace of Calisea (1714, -0.62 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.85 DPS) [vendor]; Archmagister's Faceted Pendant (273049, -1.06 DPS) [world] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 68.6 healing_power points (3.42 DPS) | yes | Windchaser Amice (14432, -0.75 DPS) [world_drop]; Mistscape Mantle (4734, -1.00 DPS) [dungeon]; Silksand Shoulder Pads (14423, -1.06 DPS) [world_drop] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 48.0 healing_power points (2.40 DPS) | yes | Blackforge Cape (6424, -0.62 DPS) [dungeon]; Glyphed Cloak (4732, -0.96 DPS) [world_drop]; Windchaser Cloak (14430, -1.03 DPS) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-verified (87.5 DPS) | yes | Red Mageweave Vest (10007, +0.00 DPS) [crafted]; Silksand Tunic (14417, +0.00 DPS) [world_drop]; Silksand Wraps (14425, +0.00 DPS) [world_drop] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 37.1 healing_power points (1.85 DPS) | yes | Mistscape Bracers (4045, -0.21 DPS) [dungeon]; Radiant Silver Bracers (4545, -0.21 DPS) [quest]; Enchanted Stonecloth Bracers (4979, -0.21 DPS) [quest] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | 60.5 healing_power points (3.02 DPS) | yes | Gilded Handwraps (254021, -0.48 DPS) [crafted]; Red Mageweave Gloves (10018, -0.96 DPS) [crafted]; Mistscape Gloves (6428, -1.17 DPS) [dungeon] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deathmage Sash (10771, +0.00 DPS) [dungeon]; Razzeric's Customized Seatbelt (6726, -0.47 DPS) [quest]; Mistscape Sash (4736, -0.68 DPS) [dungeon] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-verified (87.5 DPS) | yes | Crimson Silk Pantaloons (7062, +0.00 DPS) [crafted]; Stoneweaver Leggings (9407, +0.00 DPS) [dungeon]; Pristine Leggings (253987, +0.00 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 51.5 healing_power points (2.57 DPS) | yes | Boots of the Maharishi (9658, -0.52 DPS) [quest]; Furen's Boots (13100, -0.60 DPS) [world_drop]; Kodo Rustler Boots (15697, -0.76 DPS) [quest] |
| finger1 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 30.9 healing_power points (1.54 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Welken Ring (5011, -0.28 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.30 DPS) [dungeon] |
| finger2 | Mindbender Loop (5009) | World drop [world_drop] | sim-verified (87.5 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Welken Ring (5011, -0.25 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.27 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hypnotic Blade (7714, -1.01 DPS) [dungeon]; Illusionary Rod (7713, -1.29 DPS) [dungeon]; Death Speaker Scepter (2816, -1.56 DPS) [dungeon] |
| off_hand | Witch's Finger (16887) | Baron Aquanis [quest] | sim-verified (87.5 DPS) | yes | Aurora Sphere (7610, -0.07 DPS) [dungeon]; Skull of Impending Doom (4984, -0.21 DPS) [quest]; Nightsky Orb (15929, -0.31 DPS) [world_drop] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 26.7 healing_power points (1.33 DPS) | yes | Flash Wand (5248, -0.41 DPS) [quest]; Goblin Igniter (5253, -0.41 DPS) [quest]; Wand of Mana Concentration (273043, -0.45 DPS) [world] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Stormcloth Gloves; waist: Gilded Cord; legs: Stormcloth Pants; finger1: Voodoo Band; finger2: Mindbender Loop; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Witch's Finger; ranged: Jaina's Firestarter

No-known-source sample (15 of 322, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16336 Sergeant Major's Cape

### Band 50 (troll, 025003000000000000-03505003030121431-00000000000000000)

Set DPS (verified): 123.4. Weights run: 10.7s. Verify run: 34.3s. 413 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.100, intellect=3.827 ± 0.050, spirit=0.984 ± 0.039, mp5=1.330 ± 0.040, crit=0.303 ± 0.024 per rating point (14 rating = 1%, 4.245 per %), spell_haste=not significant (-2.915 ± 0.745)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 91.8 healing_power points (5.04 DPS) | yes | Papal Fez (9431, -0.21 DPS) [dungeon]; Blood Guard's Satin Cover (220899, -0.35 DPS) [vendor]; Eye of Theradras (17715, -0.51 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 57.5 healing_power points (3.16 DPS) | yes | Glowing Eye of Mordresh (10769, -0.76 DPS) [dungeon]; Darkspear Warding Pendant (272073, -1.27 DPS) [vendor]; Necklace of Calisea (1714, -1.31 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 68.9 healing_power points (3.78 DPS) | yes | Inquisitor's Shawl (19507, +0.00 DPS, sim-verified) [dungeon]; Blood Guard's Satin Pads (220901, -0.18 DPS) [vendor]; Nethergeld Shoulders (254049, -0.35 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 58.5 healing_power points (3.21 DPS) | yes | Imperial Red Cloak (8248, -0.97 DPS, sim-verified) [world_drop]; Highborne Cloak (14450, -1.10 DPS) [world_drop]; Keeper's Cloak (14665, -1.11 DPS) [world_drop] |
| chest | Stone Guard's Satin Armor (220903) | Lady Palanseer [vendor] | 78.9 healing_power points (4.33 DPS) | yes | Acumen Robes (17775, -0.13 DPS) [quest]; Red Mageweave Vest (10007, -0.55 DPS) [crafted]; Hibernal Robe (8113, -3.10 DPS, sim-verified) [world_drop] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 51.7 healing_power points (2.84 DPS) | yes | Forgotten Wraps (9433, -0.32 DPS) [world_drop]; Imperial Red Bracers (8247, -0.53 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -1.94 DPS, sim-verified) [quest] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 106.2 healing_power points (5.83 DPS) | yes | Virtuous Mitts (226950, -1.08 DPS, sim-verified) [vendor]; Gilded Gloves (254095, -1.71 DPS) [crafted]; Greenleaf Handwraps (19116, -2.31 DPS) [quest] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 84.7 healing_power points (4.65 DPS) | yes | Gilded Waistcord (254081, -1.17 DPS) [crafted]; Gilded Cord (254037, -1.44 DPS) [crafted]; Serenity Belt (13144, -2.32 DPS, sim-verified) [world_drop] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 92.6 healing_power points (5.08 DPS) | yes | Stone Guard's Satin Leggings (220902, +0.00 DPS, sim-verified) [vendor]; Senior Designer's Pantaloons (11841, -1.27 DPS) [dungeon]; Venomshroud Leggings (14444, -1.28 DPS) [world_drop] |
| feet | First Sergeant's Satin Boots (220900) | PvP rank 9 · First Sergeant · Horde [vendor] | 65.9 healing_power points (3.62 DPS) | yes | Gilded Sandals (254107, +0.00 DPS, sim-verified) [crafted]; Vinerot Sandals (17748, -0.53 DPS) [dungeon]; Gilded Slippers (254001, -0.83 DPS) [crafted] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 34.8 healing_power points (1.91 DPS) | yes | Ogremind Ring (1993, -0.28 DPS) [world_drop]; Mindbender Loop (5009, -0.33 DPS) [world_drop]; Lonetree's Circle (18586, -0.43 DPS) [quest] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 29.7 healing_power points (1.63 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Mindbender Loop (5009, -0.05 DPS) [world_drop]; Lonetree's Circle (18586, -0.16 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (123.3 DPS) | yes | Ankh of Life (1713, -0.45 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.56 DPS) [quest]; Alchemist's Stone (13503, -0.88 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (123.3 DPS) | yes | Evonice's Landin' Pilla (18951, -0.45 DPS) [quest]; Alchemist's Stone (13503, -0.77 DPS) [crafted]; Ankh of Life (1713, -1.62 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (123.3 DPS) | yes | Charstone Dirk (17710, -5.99 DPS) [dungeon]; Spellshifter Rod (9527, -7.02 DPS) [quest]; Hammer of the Northern Wind (810, -27.67 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 25.9 healing_power points (1.42 DPS) | yes | Nature's Breath (19118, -0.41 DPS, sim-verified) [quest]; Flash Wand (5248, -0.42 DPS) [quest]; Goblin Igniter (5253, -0.42 DPS) [quest] |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Darkspear Raider's Cloak; chest: Stone Guard's Satin Armor; wrist: Nethergeld Cuffs; hands: Raider Handwraps; waist: Dawnspire Cord; legs: Kilt of the Atal'ai Prophet; feet: First Sergeant's Satin Boots; finger1: Eye of Adaegus; finger2: Voodoo Band; trinket1: Darkspear Voodoo Seal; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 413, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60 (troll, 025003031330000000-03505003030121431-00000000000000000)

Set DPS (verified): 405.1. Weights run: 9.6s. Verify run: 65.5s. 939 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.156, intellect=1.688 ± 0.065, spirit=0.435 ± 0.082, mp5=0.712 ± 0.061, crit=0.205 ± 0.029 per rating point (14 rating = 1%, 2.870 per %), spell_haste=-13.005 ± 1.372

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 71.5 healing_power points (12.07 DPS) | yes | Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Devout Crown (16693, -16.39 DPS, sim-verified) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | sim-verified (405.2 DPS) | yes | The Eye of Zuldazar (19593, -0.45 DPS) [quest]; The All-Seeing Eye of Zuldazar (19594, -0.45 DPS) [quest]; Drake Tooth Necklace (21531, -4.72 DPS, sim-verified) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 67.6 healing_power points (11.41 DPS) | yes | Warlord's Satin Mantle (231631, -0.65 DPS) [pvp]; Champion's Satin Mantle (227120, -1.73 DPS) [pvp]; Devout Mantle (16695, -11.59 DPS, sim-verified) [dungeon] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 58.9 healing_power points (9.94 DPS) | yes | Drape of Recovery (272413, -3.70 DPS, sim-verified) [vendor]; Darkspear Raider's Cloak (272063, -4.94 DPS) [vendor]; Battle Healer's Cloak (19526, -4.96 DPS) [rep] |
| chest | Virtuous Robe (226945) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Legionnaire's Satin Tunic (227124, -0.34 DPS) [pvp]; Truefaith Vestments (14154, -9.84 DPS, sim-verified) [crafted] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Satin Bracers (17619, -0.26 DPS) [pvp]; Nethergeld Cuffs (254061, -0.67 DPS) [crafted]; Bracers of Hope (22667, -13.15 DPS, sim-verified) [quest] |
| hands | Desert Bloom Gloves (20717) | Armaments of War [quest] | sim-verified (+3.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Mooncloth Gloves (18409, -0.27 DPS) [crafted]; Raider Handwraps (272098, -0.42 DPS) [vendor] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 58.5 healing_power points (9.88 DPS) | yes | Virtuous Belt (226948, -1.68 DPS) [quest]; General's Satin Cinch (17621, -2.23 DPS) [pvp]; Devout Belt (16696, -3.85 DPS, sim-verified) [dungeon] |
| legs | Virtuous Skirt (226946) | Anthion's Parting Words [quest] | 65.8 healing_power points (11.11 DPS) | yes | Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Devout Skirt (16694, -14.89 DPS, sim-verified) [dungeon] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 73.7 healing_power points (12.44 DPS) | yes | General's Satin Walkers (231630, -2.42 DPS) [vendor]; Virtuous Sandals (226952, -3.65 DPS) [quest]; Mooncloth Boots (15802, -6.46 DPS, sim-verified) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 53.8 healing_power points (9.07 DPS) | yes | Blessed Band of Light (272407, -3.25 DPS) [vendor]; Channeler's Ring (272406, -4.50 DPS) [vendor]; Band of the Hierophant (13096, -5.42 DPS) [world_drop] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 41.7 healing_power points (7.04 DPS) | yes | Blessed Band of Light (272407, -1.29 DPS, sim-verified) [vendor]; Channeler's Ring (272406, -2.47 DPS) [vendor]; Band of the Hierophant (13096, -3.39 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+4.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Serenity Field (272439, +0.00 DPS) [vendor]; Ankh of Life (1713, -0.92 DPS) [world_drop] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -11.29 DPS) [dungeon]; Hand of Righteousness (7721, -17.79 DPS) [dungeon]; Hand of Edward the Odd (2243, -24.15 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 21.8 healing_power points (3.68 DPS) | yes | Jaina's Firestarter (13064, -1.75 DPS) [world_drop]; Nature's Breath (19118, -1.97 DPS) [quest]; Brilliant Wand (249385, -6.78 DPS, sim-verified) [crafted] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Virtuous Robe; wrist: Virtuous Bracers; hands: Desert Bloom Gloves; waist: Wisdom of the Timbermaw; legs: Virtuous Skirt; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Shard of the Splithooves; ranged: Torch of Light

No-known-source sample (15 of 939, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60, raid preset (troll, 025003031330000000-03505003030121431-00000000000000000)

Set DPS (verified): 768.6. Weights run: 6.2s. Verify run: 13.7s. 939 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.231, intellect=0.960 ± 0.032, spirit=0.829 ± 0.032, mp5=1.644 ± 0.037, crit=0.312 ± 0.025 per rating point (14 rating = 1%, 4.363 per %), spell_haste=not significant (-0.884 ± 0.705)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 66.9 healing_power points (25.10 DPS) | yes | Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -0.70 DPS, sim-verified) [crafted] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 42.9 healing_power points (16.09 DPS) | yes | Drake Tooth Necklace (21531, -1.20 DPS, sim-verified) [quest]; The Eye of Zuldazar (19593, -2.42 DPS) [quest]; The All-Seeing Eye of Zuldazar (19594, -2.42 DPS) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 77.1 healing_power points (28.93 DPS) | yes | Mooncloth Shoulders (14139, -5.53 DPS, sim-verified) [crafted]; Champion's Satin Mantle (227120, -8.63 DPS) [pvp]; Warlord's Satin Mantle (231631, -10.48 DPS) [pvp] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 51.6 healing_power points (19.35 DPS) | yes | Drape of Recovery (272413, -1.66 DPS, sim-verified) [vendor]; Battle Healer's Cloak (19526, -7.11 DPS) [rep] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 86.3 healing_power points (32.34 DPS) | yes | Warlord's Satin Tunic (231632, -4.10 DPS) [vendor]; Virtuous Robe (226945, -6.41 DPS) [quest]; Legionnaire's Satin Tunic (227124, -7.60 DPS) [pvp] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 38.5 healing_power points (14.44 DPS) | yes | Virtuous Bracers (226949, +0.00 DPS, sim-verified) [quest]; Nethergeld Cuffs (254061, -2.86 DPS) [crafted]; Earthenweave Cuffs (254131, -3.08 DPS) [crafted] |
| hands | Desert Bloom Gloves (20717) | Armaments of War [quest] | 59.9 healing_power points (22.48 DPS) | yes | Raider Handwraps (272097, -0.60 DPS, sim-verified) [vendor]; Virtuous Mitts (226950, -3.93 DPS) [vendor]; Mooncloth Gloves (18409, -4.34 DPS) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 51.9 healing_power points (19.46 DPS) | yes | Virtuous Belt (226948, +0.00 DPS, sim-verified) [quest]; Penitent's Cinch (272394, -3.38 DPS) [vendor]; Devout Belt (16696, -5.34 DPS) [dungeon] |
| legs | Virtuous Skirt (226946) | Anthion's Parting Words [quest] | 61.3 healing_power points (22.98 DPS) | yes | Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Devout Skirt (16694, -0.44 DPS, sim-verified) [dungeon] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 69.0 healing_power points (25.89 DPS) | yes | Virtuous Sandals (226952, -6.26 DPS) [quest]; Mooncloth Boots (15802, -7.81 DPS) [crafted]; General's Satin Walkers (231630, -8.53 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 51.9 healing_power points (19.45 DPS) | yes | Blessed Band of Light (272407, -5.34 DPS) [vendor]; Channeler's Ring (272406, -11.92 DPS) [vendor]; Band of the Hierophant (13096, -12.43 DPS) [world_drop] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 38.2 healing_power points (14.31 DPS) | yes | Blessed Band of Light (272407, -0.20 DPS) [vendor]; Channeler's Ring (272406, -6.79 DPS) [vendor]; Band of the Hierophant (13096, -7.30 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (768.6 DPS) | yes | Uther's Strength (11302, -2.50 DPS) [world_drop]; Ankh of Life (1713, -5.52 DPS) [world_drop]; Shard of the Splithooves (10659, -5.55 DPS) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (768.6 DPS) | yes | Uther's Strength (11302, -2.92 DPS, sim-verified) [world_drop]; Ankh of Life (1713, -12.02 DPS) [world_drop]; Shard of the Splithooves (10659, -12.05 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (768.6 DPS) | yes | Hand of Edward the Odd (2243, -17.91 DPS, sim-verified) [world_drop]; Charstone Dirk (17710, -20.79 DPS) [dungeon]; Hand of Righteousness (7721, -31.75 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 26.1 healing_power points (9.79 DPS) | yes | Captain Rackmore's Tiller (16789, -2.33 DPS, sim-verified) [quest]; Jaina's Firestarter (13064, -6.70 DPS) [world_drop]; Wand of Mana Concentration (273043, -7.12 DPS) [world] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Bracers of Hope; hands: Desert Bloom Gloves; waist: Wisdom of the Timbermaw; legs: Virtuous Skirt; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 939, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

