# Leveling BiS: Holy

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-03503000000000000-000000000000000000)

Set DPS (verified): 40.9. Weights run: 7.5s. Verify run: 20.8s. 148 eligible items had no known source.

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

No-known-source sample (15 of 148, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes; 16607 Acolyte's Sacrificial Robes; 18851 Insignia of the Horde; 20434 Lorekeeper's Staff

### Band 30 (gnome, 000000000000000000-03505003030110000-000000000000000000)

Set DPS (verified): 65.1. Weights run: 7.8s. Verify run: 23.2s. 256 eligible items had no known source.

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

No-known-source sample (15 of 256, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 40 (gnome, 000000000000000000-03505003030121431-000000000000000000)

Set DPS (verified): 86.0. Weights run: 7.7s. Verify run: 49.9s. 336 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.075, intellect=4.127 ± 0.033, spirit=0.660 ± 0.020, mp5=0.606 ± 0.023, crit=0.213 ± 0.017 per rating point (14 rating = 1%, 2.984 per %), spell_haste=not significant (0.974 ± 0.519)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 90.1 healing_power points (4.49 DPS) | yes | Cowl of Necromancy (2621, -1.41 DPS) [world]; Craftsman's Monocle (4393, -1.41 DPS) [crafted]; Thinking Cap (2624, -1.74 DPS, sim-verified) [world] |
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

No-known-source sample (15 of 336, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 50 (gnome, 025003000000000000-03505003030121431-000000000000000000)

Set DPS (verified): 104.4. Weights run: 8.0s. Verify run: 59.3s. 430 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.042, intellect=4.816 ± 0.060, spirit=0.208 ± 0.028, mp5=1.346 ± 0.060, crit=0.289 ± 0.021 per rating point (14 rating = 1%, 4.052 per %), spell_haste=-25.020 ± 1.049

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Satin Cover (220896) | Captain Dirgehammer [vendor] | sim-verified (104.4 DPS) | yes | Papal Fez (9431, -0.03 DPS) [dungeon]; Red Mageweave Headband (10033, -0.09 DPS) [crafted]; Bad Mojo Mask (9470, -6.32 DPS, sim-verified) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 68.3 healing_power points (3.18 DPS) | yes | Glowing Eye of Mordresh (10769, -0.78 DPS) [dungeon]; Darkspear Warding Pendant (272073, -1.16 DPS) [vendor]; Darkmoon Necklace (19303, -1.46 DPS) [vendor] |
| shoulder | Knight-Lieutenant's Satin Pads (220894) | Captain Dirgehammer [vendor] | sim-verified (104.4 DPS) | yes | Inquisitor's Shawl (19507, +0.00 DPS) [dungeon]; Imperial Red Mantle (8250, -0.15 DPS) [world_drop]; Rotgrip Mantle (17732, -5.52 DPS, sim-verified) [dungeon] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 68.5 healing_power points (3.19 DPS) | yes | Keeper's Cloak (14665, -0.95 DPS) [world_drop]; Imperial Red Cloak (8248, -1.08 DPS, sim-verified) [world_drop]; Highborne Cloak (14450, -1.13 DPS) [world_drop] |
| chest | Knight's Satin Armor (220892) | Captain Dirgehammer [vendor] | sim-verified (104.4 DPS) | yes | Acumen Robes (17775, +0.00 DPS) [quest]; Red Mageweave Vest (10007, -0.19 DPS) [crafted]; Hibernal Robe (8113, -5.91 DPS, sim-verified) [world_drop] |
| wrist | Shizzle's Nozzle Wiper (11917) | Shizzle's Flyer [quest] | 58.4 healing_power points (2.72 DPS) | yes | Forgotten Wraps (9433, -0.03 DPS) [world_drop]; Nethergeld Cuffs (254061, -0.17 DPS) [crafted]; Imperial Red Bracers (8247, -0.25 DPS) [world_drop] |
| hands | Sergeant Major's Satin Gloves (220897) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-verified (104.4 DPS) | yes | Virtuous Mitts (226950, +0.00 DPS) [vendor]; Gilded Gloves (254095, +0.00 DPS) [crafted]; Raider Handwraps (272098, -0.75 DPS, sim-verified) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 103.5 healing_power points (4.82 DPS) | yes | Gilded Waistcord (254081, -1.45 DPS) [crafted]; Imperial Red Sash (8253, -1.46 DPS) [world_drop]; Serenity Belt (13144, -1.89 DPS, sim-verified) [world_drop] |
| legs | Knight's Satin Leggings (220893) | Captain Dirgehammer [vendor] | sim-verified (104.4 DPS) | yes | Venomshroud Leggings (14444, -0.75 DPS) [world_drop]; Imperial Red Pants (8251, -0.87 DPS) [world_drop]; Kilt of the Atal'ai Prophet (10807, -4.82 DPS, sim-verified) [dungeon] |
| feet | Sergeant Major's Satin Boots (220895) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 77.8 healing_power points (3.62 DPS) | yes | Highborne Footpads (14447, -0.93 DPS) [world_drop]; Vinerot Sandals (17748, -0.95 DPS) [dungeon]; Gilded Sandals (254107, -3.29 DPS, sim-verified) [crafted] |
| finger1 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 34.3 healing_power points (1.60 DPS) | yes | Mindbender Loop (5009, -0.01 DPS) [world_drop]; Eye of Adaegus (5266, -0.14 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.25 DPS) [dungeon] |
| finger2 | Ogremind Ring (1993) | World drop [world_drop] | 34.3 healing_power points (1.60 DPS) | yes | Mindbender Loop (5009, -0.01 DPS) [world_drop]; Eye of Adaegus (5266, -0.14 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.25 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (104.4 DPS) | yes | Ankh of Life (1713, -0.82 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.84 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.86 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (104.4 DPS) | yes | Evonice's Landin' Pilla (18951, -0.74 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.76 DPS) [quest]; Ankh of Life (1713, -1.52 DPS, sim-verified) [world_drop] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-verified (104.4 DPS) | yes | Hammer of the Northern Wind (810, +0.00 DPS) [world_drop]; Spellshifter Rod (9527, -0.32 DPS) [quest]; Glowing Brightwood Staff (812, -1.40 DPS) [world_drop] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 60.2 healing_power points (2.80 DPS) | yes | Mistscape Stave (7611, -0.78 DPS) [world_drop]; Truesilver Conduit (249455, -1.01 DPS) [crafted]; Skullspell Orb (10708, -1.28 DPS, sim-verified) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 29.5 healing_power points (1.37 DPS) | yes | Wand of Mana Concentration (273043, -0.35 DPS) [world]; Goblin Igniter (5253, -0.45 DPS) [quest]; Cairnstone Sliver (9654, -0.58 DPS, sim-verified) [quest] |

**New at 50:** head: Knight-Lieutenant's Satin Cover; neck: Horizon Choker; shoulder: Knight-Lieutenant's Satin Pads; back: Darkspear Raider's Cloak; chest: Knight's Satin Armor; wrist: Shizzle's Nozzle Wiper; hands: Sergeant Major's Satin Gloves; waist: Dawnspire Cord; legs: Knight's Satin Leggings; feet: Sergeant Major's Satin Boots; trinket1: Darkspear Voodoo Seal; trinket2: Uther's Strength; main_hand: Charstone Dirk

No-known-source sample (15 of 430, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60 (gnome, 025003031303000000-03505003030121431-000000000000000000)

Set DPS (verified): 383.0. Weights run: 7.0s. Verify run: 52.9s. 943 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.113, intellect=1.171 ± 0.049, spirit=0.482 ± 0.073, mp5=0.612 ± 0.049, crit=0.180 ± 0.022 per rating point (14 rating = 1%, 2.527 per %), spell_haste=-6.962 ± 1.017

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 63.2 healing_power points (13.61 DPS) | yes | Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -16.74 DPS, sim-verified) [crafted] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | sim-verified (383.0 DPS) | yes | The Eye of Zuldazar (19593, -0.49 DPS) [quest]; The All-Seeing Eye of Zuldazar (19594, -0.49 DPS) [quest]; Drake Tooth Necklace (21531, -4.04 DPS, sim-verified) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 67.5 healing_power points (14.55 DPS) | yes | Field Marshal's Satin Mantle (231628, -3.03 DPS) [pvp]; Lieutenant Commander's Satin Mantle (227119, -3.38 DPS) [pvp]; Virtuous Mantle (226951, -9.30 DPS, sim-verified) [quest] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 53.7 healing_power points (11.58 DPS) | yes | Drape of Recovery (272413, -4.41 DPS, sim-verified) [vendor]; Caretaker's Cape (19530, -5.14 DPS) [rep] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 80.7 healing_power points (17.40 DPS) | yes | Field Marshal's Satin Tunic (231624, -0.56 DPS) [vendor]; Virtuous Robe (226945, -3.35 DPS, sim-verified) [quest]; Knight-Captain's Satin Tunic (227122, -3.53 DPS) [vendor] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nethergeld Cuffs (254061, -0.70 DPS) [crafted]; Marshal's Satin Bracers (17606, -1.46 DPS) [pvp]; Bracers of Hope (22667, -11.91 DPS, sim-verified) [quest] |
| hands | Desert Bloom Gloves (20717) | Armaments of War [quest] | 59.3 healing_power points (12.78 DPS) | yes | Raider Handwraps (272097, -0.60 DPS, sim-verified) [vendor]; Mooncloth Gloves (18409, -1.62 DPS) [crafted] |
| waist | Virtuous Belt (226948) | Just Compensation [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gilded Waistcord (254081, -0.58 DPS) [crafted]; Devout Belt (16696, -0.75 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -9.27 DPS, sim-verified) [crafted] |
| legs | Virtuous Skirt (226946) | Anthion's Parting Words [quest] | 58.4 healing_power points (12.59 DPS) | yes | Knight-Captain's Satin Legguards (227125, +0.00 DPS) [vendor]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Devout Skirt (16694, -13.79 DPS, sim-verified) [dungeon] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 67.2 healing_power points (14.48 DPS) | yes | Marshal's Satin Walkers (231627, -3.68 DPS) [vendor]; Virtuous Sandals (226952, -4.54 DPS) [quest]; Mooncloth Boots (15802, -6.21 DPS, sim-verified) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 48.6 healing_power points (10.48 DPS) | yes | Blessed Band of Light (272407, -2.96 DPS) [vendor]; Channeler's Ring (272406, -6.11 DPS) [vendor]; Sea Giant's Toe Ring (274746, -6.60 DPS) [vendor] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 36.2 healing_power points (7.80 DPS) | yes | Blessed Band of Light (272407, -1.68 DPS, sim-verified) [vendor]; Channeler's Ring (272406, -3.43 DPS) [vendor]; Sea Giant's Toe Ring (274746, -3.91 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Ankh of Life (1713, -0.73 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.94 DPS) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -5.17 DPS) [world_drop]; Ankh of Life (1713, -7.81 DPS) [world_drop]; Shard of the Splithooves (10659, -9.41 DPS, sim-verified) [quest] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Edward the Odd (2243, +0.00 DPS) [world_drop]; Hand of Righteousness (7721, -6.95 DPS) [dungeon]; Spellshifter Rod (9527, -10.60 DPS) [quest] |
| off_hand | Lei of the Lifegiver (19312) | Stormpike Guard [rep] | 43.1 healing_power points (9.28 DPS) | yes | Grand Marshal's Tome of Restoration (234590, +0.00 DPS) [pvp]; Twisting Essence Jar (249456, -4.35 DPS) [crafted]; Orb of Lorica (11262, -7.67 DPS, sim-verified) [quest] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 22.3 healing_power points (4.81 DPS) | yes | Cairnstone Sliver (9654, -3.03 DPS) [quest]; Brilliant Wand (249385, -3.04 DPS) [crafted]; Jaina's Firestarter (13064, -5.00 DPS, sim-verified) [world_drop] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Virtuous Bracers; hands: Desert Bloom Gloves; waist: Virtuous Belt; legs: Virtuous Skirt; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Serenity Field; off_hand: Lei of the Lifegiver; ranged: Torch of Light

No-known-source sample (15 of 943, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60, raid preset (gnome, 025003031303000000-03505003030121431-000000000000000000)

Set DPS (verified): 758.4. Weights run: 4.6s. Verify run: 18.7s. 943 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.299, intellect=1.359 ± 0.049, spirit=1.201 ± 0.048, mp5=2.409 ± 0.054, crit=0.496 ± 0.035 per rating point (14 rating = 1%, 6.939 per %), spell_haste=not significant (-1.646 ± 1.051)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 82.3 healing_power points (21.25 DPS) | yes | Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Lieutenant Commander's Satin Hood (227121, -0.25 DPS) [pvp]; Devout Crown (16693, -3.01 DPS) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 51.4 healing_power points (13.29 DPS) | yes | The Eye of Zuldazar (19593, -1.79 DPS) [quest]; The All-Seeing Eye of Zuldazar (19594, -1.79 DPS) [quest]; Drake Tooth Necklace (21531, -2.36 DPS) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 85.5 healing_power points (22.08 DPS) | yes | Lieutenant Commander's Satin Mantle (227119, -5.52 DPS) [pvp]; Field Marshal's Satin Mantle (231628, -7.31 DPS) [pvp]; Virtuous Mantle (226951, -8.11 DPS) [quest] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 55.6 healing_power points (14.36 DPS) | yes | Drape of Recovery (272413, -3.86 DPS) [vendor]; Caretaker's Cape (19530, -5.16 DPS) [rep] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 92.2 healing_power points (23.83 DPS) | yes | Field Marshal's Satin Tunic (231624, -0.61 DPS) [vendor]; Virtuous Robe (226945, -2.20 DPS) [quest]; Knight-Captain's Satin Tunic (227122, -3.88 DPS) [vendor] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bracers of Hope (22667, +0.00 DPS) [quest]; Marshal's Satin Bracers (17606, -1.45 DPS) [pvp]; Nethergeld Cuffs (254061, -2.05 DPS) [crafted] |
| hands | Desert Bloom Gloves (20717) | Armaments of War [quest] | sim-verified (758.5 DPS) | yes | Raider Handwraps (272097, -0.81 DPS) [vendor]; Virtuous Mitts (226950, -0.84 DPS) [vendor]; Mooncloth Gloves (18409, -2.34 DPS) [crafted] |
| waist | Virtuous Belt (226948) | Just Compensation [quest] | sim-verified (758.5 DPS) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Penitent's Cinch (272394, -1.31 DPS) [vendor]; Marshal's Satin Sash (17609, -1.97 DPS) [pvp] |
| legs | Virtuous Skirt (226946) | Anthion's Parting Words [quest] | 74.4 healing_power points (19.21 DPS) | yes | Knight-Captain's Satin Legguards (227125, +0.00 DPS) [vendor]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Devout Skirt (16694, -0.61 DPS) [dungeon] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 79.8 healing_power points (20.63 DPS) | yes | Virtuous Sandals (226952, -3.45 DPS) [quest]; Mooncloth Boots (15802, -5.48 DPS) [crafted]; Marshal's Satin Walkers (231627, -6.82 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 59.3 healing_power points (15.32 DPS) | yes | Blessed Band of Light (272407, -4.82 DPS) [vendor]; Channeler's Ring (272406, -7.92 DPS) [vendor]; Band of the Hierophant (13096, -8.39 DPS) [world_drop] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 45.2 healing_power points (11.68 DPS) | yes | Blessed Band of Light (272407, -1.19 DPS) [vendor]; Channeler's Ring (272406, -4.29 DPS) [vendor]; Band of the Hierophant (13096, -4.76 DPS) [world_drop] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -6.20 DPS) [world_drop]; Shard of the Splithooves (10659, -7.12 DPS) [quest]; Ankh of Life (1713, -7.13 DPS) [world_drop] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -4.68 DPS) [world_drop]; Shard of the Splithooves (10659, -5.60 DPS) [quest]; Ankh of Life (1713, -5.61 DPS) [world_drop] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Righteousness (7721, -8.31 DPS) [dungeon]; Spellshifter Rod (9527, -13.05 DPS) [quest]; Spire of Hakkar (10844, -14.09 DPS) [world] |
| off_hand | Lei of the Lifegiver (19312) | Stormpike Guard [rep] | 52.0 healing_power points (13.45 DPS) | yes | Grand Marshal's Tome of Restoration (234590, -1.04 DPS) [pvp]; Twisting Essence Jar (249456, -6.06 DPS) [crafted]; Lapidis Tankard of Tidesippe (4696, -6.73 DPS) [world_drop] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 30.2 healing_power points (7.81 DPS) | yes | Cairnstone Sliver (9654, -4.50 DPS) [quest]; Jaina's Firestarter (13064, -4.77 DPS) [world_drop]; Wand of Mana Concentration (273043, -5.16 DPS) [world] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Virtuous Bracers; hands: Desert Bloom Gloves; waist: Virtuous Belt; legs: Virtuous Skirt; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket1: Serenity Field; trinket2: Darkspear Voodoo Seal; off_hand: Lei of the Lifegiver; ranged: Torch of Light

No-known-source sample (15 of 943, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

## Horde

### Band 20 (troll, 000000000000000000-03503000000000000-000000000000000000)

Set DPS (verified): 38.8. Weights run: 7.5s. Verify run: 20.9s. 138 eligible items had no known source.

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

No-known-source sample (15 of 138, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 209613 Insignia of the Alliance; 241089 Scarlet Dagger; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff; 263007 Skyseer's Vest; 263402 Sash of Sorrow

### Band 30 (troll, 000000000000000000-03505003030110000-000000000000000000)

Set DPS (verified): 63.0. Weights run: 7.8s. Verify run: 23.4s. 242 eligible items had no known source.

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

No-known-source sample (15 of 242, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (troll, 000000000000000000-03505003030121431-000000000000000000)

Set DPS (verified): 87.5. Weights run: 7.7s. Verify run: 30.7s. 319 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.075, intellect=4.127 ± 0.033, spirit=0.660 ± 0.020, mp5=0.606 ± 0.023, crit=0.213 ± 0.017 per rating point (14 rating = 1%, 2.984 per %), spell_haste=not significant (0.974 ± 0.519)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 90.1 healing_power points (4.49 DPS) | yes | Thinking Cap (2624, -0.99 DPS) [world]; Cowl of Necromancy (2621, -1.41 DPS) [world]; Craftsman's Monocle (4393, -1.41 DPS) [crafted] |
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

No-known-source sample (15 of 319, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16336 Sergeant Major's Cape

### Band 50 (troll, 025003000000000000-03505003030121431-000000000000000000)

Set DPS (verified): 93.0. Weights run: 8.0s. Verify run: 56.1s. 410 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.042, intellect=4.816 ± 0.060, spirit=0.208 ± 0.028, mp5=1.346 ± 0.060, crit=0.289 ± 0.021 per rating point (14 rating = 1%, 4.052 per %), spell_haste=-25.020 ± 1.049

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bad Mojo Mask (9470) | Zul'Farrak: Shadowpriest Sezz'ziz [dungeon] | 101.1 healing_power points (4.71 DPS) | yes | Papal Fez (9431, +0.00 DPS, sim-verified) [dungeon]; Blood Guard's Satin Cover (220899, -0.13 DPS) [vendor]; Red Mageweave Headband (10033, -0.22 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 68.3 healing_power points (3.18 DPS) | yes | Glowing Eye of Mordresh (10769, -0.78 DPS) [dungeon]; Darkspear Warding Pendant (272073, -1.16 DPS) [vendor]; Darkmoon Necklace (19303, -1.46 DPS) [vendor] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 86.7 healing_power points (4.04 DPS) | yes | Inquisitor's Shawl (19507, +0.00 DPS, sim-verified) [dungeon]; Blood Guard's Satin Pads (220901, -0.52 DPS) [vendor]; Imperial Red Mantle (8250, -0.67 DPS) [world_drop] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 68.5 healing_power points (3.19 DPS) | yes | Imperial Red Cloak (8248, -0.48 DPS, sim-verified) [world_drop]; Keeper's Cloak (14665, -0.95 DPS) [world_drop]; Highborne Cloak (14450, -1.13 DPS) [world_drop] |
| chest | Stone Guard's Satin Armor (220903) | Lady Palanseer [vendor] | sim-verified (93.0 DPS) | yes | Acumen Robes (17775, +0.00 DPS) [quest]; Red Mageweave Vest (10007, -0.19 DPS) [crafted]; Hibernal Robe (8113, -1.08 DPS, sim-verified) [world_drop] |
| wrist | Shizzle's Nozzle Wiper (11917) | Shizzle's Flyer [quest] | 58.4 healing_power points (2.72 DPS) | yes | Forgotten Wraps (9433, -0.03 DPS) [world_drop]; Nethergeld Cuffs (254061, -0.17 DPS) [crafted]; Imperial Red Bracers (8247, -0.25 DPS) [world_drop] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 127.9 healing_power points (5.96 DPS) | yes | Virtuous Mitts (226950, -1.45 DPS, sim-verified) [vendor]; Gilded Gloves (254095, -2.29 DPS) [crafted]; Greenleaf Handwraps (19116, -2.47 DPS) [quest] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 103.5 healing_power points (4.82 DPS) | yes | Serenity Belt (13144, -1.37 DPS, sim-verified) [world_drop]; Gilded Waistcord (254081, -1.45 DPS) [crafted]; Imperial Red Sash (8253, -1.46 DPS) [world_drop] |
| legs | Stone Guard's Satin Leggings (220902) | Lady Palanseer [vendor] | sim-verified (93.0 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -0.48 DPS, sim-verified) [dungeon]; Venomshroud Leggings (14444, -0.75 DPS) [world_drop]; Imperial Red Pants (8251, -0.87 DPS) [world_drop] |
| feet | First Sergeant's Satin Boots (220900) | PvP rank 9 · First Sergeant · Horde [vendor] | 77.8 healing_power points (3.62 DPS) | yes | Gilded Sandals (254107, +0.00 DPS, sim-verified) [crafted]; Highborne Footpads (14447, -0.93 DPS) [world_drop]; Vinerot Sandals (17748, -0.95 DPS) [dungeon] |
| finger1 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 34.3 healing_power points (1.60 DPS) | yes | Mindbender Loop (5009, -0.01 DPS) [world_drop]; Eye of Adaegus (5266, -0.14 DPS) [world_drop]; Lonetree's Circle (18586, -0.21 DPS) [quest] |
| finger2 | Ogremind Ring (1993) | World drop [world_drop] | 34.3 healing_power points (1.60 DPS) | yes | Mindbender Loop (5009, -0.01 DPS) [world_drop]; Eye of Adaegus (5266, -0.14 DPS) [world_drop]; Lonetree's Circle (18586, -0.21 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (93.0 DPS) | yes | Ankh of Life (1713, -0.82 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.84 DPS) [quest]; Alchemist's Stone (13503, -0.90 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (93.0 DPS) | yes | Evonice's Landin' Pilla (18951, -0.74 DPS) [quest]; Alchemist's Stone (13503, -0.80 DPS) [crafted]; Ankh of Life (1713, -1.27 DPS, sim-verified) [world_drop] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-verified (93.0 DPS) | yes | Hammer of the Northern Wind (810, +0.00 DPS) [world_drop]; Spellshifter Rod (9527, -0.32 DPS) [quest]; Glowing Brightwood Staff (812, -1.40 DPS) [world_drop] |
| off_hand | Mistscape Stave (7611) | World drop [world_drop] | 43.3 healing_power points (2.02 DPS) | yes | Witch's Finger (16887, -0.22 DPS) [quest]; Truesilver Conduit (249455, -0.22 DPS) [crafted]; Flask of Forest Mojo (19115, -0.27 DPS, sim-verified) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 29.5 healing_power points (1.37 DPS) | yes | Nature's Breath (19118, -0.03 DPS) [quest]; Wand of Mana Concentration (273043, -0.35 DPS) [world]; Goblin Igniter (5253, -0.45 DPS) [quest] |

**New at 50:** head: Bad Mojo Mask; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Darkspear Raider's Cloak; chest: Stone Guard's Satin Armor; wrist: Shizzle's Nozzle Wiper; hands: Raider Handwraps; waist: Dawnspire Cord; legs: Stone Guard's Satin Leggings; feet: First Sergeant's Satin Boots; finger2: Ogremind Ring; trinket1: Darkspear Voodoo Seal; trinket2: Uther's Strength; main_hand: Charstone Dirk; off_hand: Mistscape Stave

No-known-source sample (15 of 410, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60 (troll, 025003031303000000-03505003030121431-000000000000000000)

Set DPS (verified): 369.3. Weights run: 7.0s. Verify run: 50.5s. 935 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.113, intellect=1.171 ± 0.049, spirit=0.482 ± 0.073, mp5=0.612 ± 0.049, crit=0.180 ± 0.022 per rating point (14 rating = 1%, 2.527 per %), spell_haste=-6.962 ± 1.017

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 63.2 healing_power points (13.61 DPS) | yes | Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -15.69 DPS, sim-verified) [crafted] |
| neck | Drake Tooth Necklace (21531) | The Nightmare Manifests [quest] | 40.1 healing_power points (8.63 DPS) | yes | Wavefront Necklace (20685, +0.00 DPS, sim-verified) [world]; The Eye of Zuldazar (19593, -1.38 DPS) [quest]; The All-Seeing Eye of Zuldazar (19594, -1.38 DPS) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 67.5 healing_power points (14.55 DPS) | yes | Warlord's Satin Mantle (231631, -3.03 DPS) [pvp]; Champion's Satin Mantle (227120, -3.38 DPS) [pvp]; Virtuous Mantle (226951, -11.01 DPS, sim-verified) [quest] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 53.7 healing_power points (11.58 DPS) | yes | Drape of Recovery (272413, -2.83 DPS, sim-verified) [vendor]; Battle Healer's Cloak (19526, -5.14 DPS) [rep] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 80.7 healing_power points (17.40 DPS) | yes | Warlord's Satin Tunic (231632, -0.56 DPS) [vendor]; Legionnaire's Satin Tunic (227124, -3.53 DPS) [pvp]; Virtuous Robe (226945, -3.94 DPS, sim-verified) [quest] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nethergeld Cuffs (254061, -0.70 DPS) [crafted]; General's Satin Bracers (17619, -1.46 DPS) [pvp]; Bracers of Hope (22667, -9.96 DPS, sim-verified) [quest] |
| hands | Desert Bloom Gloves (20717) | Armaments of War [quest] | 59.3 healing_power points (12.78 DPS) | yes | Mooncloth Gloves (18409, -1.62 DPS) [crafted]; Raider Handwraps (272097, -1.66 DPS, sim-verified) [vendor] |
| waist | Virtuous Belt (226948) | Just Compensation [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gilded Waistcord (254081, -0.58 DPS) [crafted]; Devout Belt (16696, -0.75 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -10.35 DPS, sim-verified) [crafted] |
| legs | Virtuous Skirt (226946) | Anthion's Parting Words [quest] | 58.4 healing_power points (12.59 DPS) | yes | Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Devout Skirt (16694, -12.36 DPS, sim-verified) [dungeon] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 67.2 healing_power points (14.48 DPS) | yes | General's Satin Walkers (231630, -3.68 DPS) [vendor]; Virtuous Sandals (226952, -4.54 DPS) [quest]; Mooncloth Boots (15802, -6.33 DPS, sim-verified) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 48.6 healing_power points (10.48 DPS) | yes | Blessed Band of Light (272407, -2.96 DPS) [vendor]; Channeler's Ring (272406, -6.11 DPS) [vendor]; Sea Giant's Toe Ring (274746, -6.60 DPS) [vendor] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 36.2 healing_power points (7.80 DPS) | yes | Blessed Band of Light (272407, -1.69 DPS, sim-verified) [vendor]; Channeler's Ring (272406, -3.43 DPS) [vendor]; Sea Giant's Toe Ring (274746, -3.91 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+3.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Ankh of Life (1713, -0.73 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.94 DPS) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (369.2 DPS) | yes | Uther's Strength (11302, -5.17 DPS) [world_drop]; Ankh of Life (1713, -7.81 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -8.89 DPS, sim-verified) [quest] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Edward the Odd (2243, +0.00 DPS) [world_drop]; Hand of Righteousness (7721, -6.95 DPS) [dungeon]; Spellshifter Rod (9527, -10.60 DPS) [quest] |
| off_hand | Lei of the Lifegiver (19312) | Frostwolf Clan [rep] | 43.1 healing_power points (9.28 DPS) | yes | High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp]; Elunarian Sphere (15968, -5.75 DPS) [world_drop]; Twisting Essence Jar (249456, -7.52 DPS, sim-verified) [crafted] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 22.3 healing_power points (4.81 DPS) | yes | Brilliant Wand (249385, -3.04 DPS) [crafted]; Captain Rackmore's Tiller (16789, -3.20 DPS) [quest]; Jaina's Firestarter (13064, -5.10 DPS, sim-verified) [world_drop] |

**New at 60:** head: Virtuous Crown; neck: Drake Tooth Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Virtuous Bracers; hands: Desert Bloom Gloves; waist: Virtuous Belt; legs: Virtuous Skirt; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Serenity Field; off_hand: Lei of the Lifegiver; ranged: Torch of Light

No-known-source sample (15 of 935, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60, raid preset (troll, 025003031303000000-03505003030121431-000000000000000000)

Set DPS (verified): 745.5. Weights run: 4.6s. Verify run: 20.1s. 935 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.299, intellect=1.359 ± 0.049, spirit=1.201 ± 0.048, mp5=2.409 ± 0.054, crit=0.496 ± 0.035 per rating point (14 rating = 1%, 6.939 per %), spell_haste=not significant (-1.646 ± 1.051)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 82.3 healing_power points (21.25 DPS) | yes | Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Champion's Satin Hood (227118, -0.25 DPS) [pvp]; Devout Crown (16693, -3.01 DPS) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 51.4 healing_power points (13.29 DPS) | yes | The Eye of Zuldazar (19593, -1.79 DPS) [quest]; The All-Seeing Eye of Zuldazar (19594, -1.79 DPS) [quest]; Drake Tooth Necklace (21531, -2.36 DPS) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 85.5 healing_power points (22.08 DPS) | yes | Champion's Satin Mantle (227120, -5.52 DPS) [pvp]; Warlord's Satin Mantle (231631, -7.31 DPS) [pvp]; Virtuous Mantle (226951, -8.11 DPS) [quest] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 55.6 healing_power points (14.36 DPS) | yes | Drape of Recovery (272413, -3.86 DPS) [vendor]; Battle Healer's Cloak (19526, -5.16 DPS) [rep] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 92.2 healing_power points (23.83 DPS) | yes | Warlord's Satin Tunic (231632, -0.61 DPS) [vendor]; Virtuous Robe (226945, -2.20 DPS) [quest]; Legionnaire's Satin Tunic (227124, -3.88 DPS) [pvp] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bracers of Hope (22667, +0.00 DPS) [quest]; General's Satin Bracers (17619, -1.45 DPS) [pvp]; Nethergeld Cuffs (254061, -2.05 DPS) [crafted] |
| hands | Desert Bloom Gloves (20717) | Armaments of War [quest] | sim-verified (745.6 DPS) | yes | Raider Handwraps (272097, -0.81 DPS) [vendor]; Virtuous Mitts (226950, -0.84 DPS) [vendor]; Mooncloth Gloves (18409, -2.34 DPS) [crafted] |
| waist | Virtuous Belt (226948) | Just Compensation [quest] | sim-verified (745.6 DPS) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Penitent's Cinch (272394, -1.31 DPS) [vendor]; General's Satin Cinch (17621, -1.97 DPS) [pvp] |
| legs | Virtuous Skirt (226946) | Anthion's Parting Words [quest] | 74.4 healing_power points (19.21 DPS) | yes | Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Devout Skirt (16694, -0.61 DPS) [dungeon] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 79.8 healing_power points (20.63 DPS) | yes | Virtuous Sandals (226952, -3.45 DPS) [quest]; Mooncloth Boots (15802, -5.48 DPS) [crafted]; General's Satin Walkers (231630, -6.82 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 59.3 healing_power points (15.32 DPS) | yes | Blessed Band of Light (272407, -4.82 DPS) [vendor]; Channeler's Ring (272406, -7.92 DPS) [vendor]; Band of the Hierophant (13096, -8.39 DPS) [world_drop] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 45.2 healing_power points (11.68 DPS) | yes | Blessed Band of Light (272407, -1.19 DPS) [vendor]; Channeler's Ring (272406, -4.29 DPS) [vendor]; Band of the Hierophant (13096, -4.76 DPS) [world_drop] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -6.20 DPS) [world_drop]; Shard of the Splithooves (10659, -7.12 DPS) [quest]; Ankh of Life (1713, -7.13 DPS) [world_drop] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -4.68 DPS) [world_drop]; Shard of the Splithooves (10659, -5.60 DPS) [quest]; Ankh of Life (1713, -5.61 DPS) [world_drop] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Righteousness (7721, -8.31 DPS) [dungeon]; Spellshifter Rod (9527, -13.05 DPS) [quest]; Spire of Hakkar (10844, -14.09 DPS) [world] |
| off_hand | Lei of the Lifegiver (19312) | Frostwolf Clan [rep] | 52.0 healing_power points (13.45 DPS) | yes | High Warlord's Tome of Mending (234564, -1.04 DPS) [pvp]; Twisting Essence Jar (249456, -6.06 DPS) [crafted]; Lapidis Tankard of Tidesippe (4696, -6.73 DPS) [world_drop] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 30.2 healing_power points (7.81 DPS) | yes | Jaina's Firestarter (13064, -4.77 DPS) [world_drop]; Wand of Mana Concentration (273043, -5.16 DPS) [world]; Captain Rackmore's Tiller (16789, -5.33 DPS) [quest] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Virtuous Bracers; hands: Desert Bloom Gloves; waist: Virtuous Belt; legs: Virtuous Skirt; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket1: Serenity Field; trinket2: Darkspear Voodoo Seal; off_hand: Lei of the Lifegiver; ranged: Torch of Light

No-known-source sample (15 of 935, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

