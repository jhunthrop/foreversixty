# Leveling BiS: Holy

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-03503000000000000-00000000000000000)

Set DPS (verified): 40.9. Weights run: 10.2s. Verify run: 28.1s. 146 eligible items had no known source.

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

Set DPS (verified): 64.8. Weights run: 10.8s. Verify run: 33.2s. 258 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.013, intellect=1.400 ± 0.009, spirit=0.888 ± 0.004, mp5=1.683 ± 0.017, crit=0.093 ± 0.005 per rating point (14 rating = 1%, 1.303 per %), spell_haste=0.255 ± 0.053

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 38.3 healing_power points (3.45 DPS) | yes | Nightsky Cowl (4039, -1.54 DPS) [world_drop]; Filigreed Pristine Circlet (253975, -1.63 DPS, sim-verified) [crafted]; Resilient Cap (14401, -1.75 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 9.2 healing_power points (0.82 DPS) | yes | Scorn's Icy Choker (23169, -0.07 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -0.07 DPS) [world]; Crystal Starfire Medallion (5003, -0.08 DPS) [world_drop] |
| shoulder | Mantle of Honor (3560) | Bride of the Embalmer [quest] | 16.0 healing_power points (1.44 DPS) | yes | Nightsky Mantle (4718, +0.00 DPS, sim-verified) [world_drop]; Death Speaker Mantle (6685, -0.06 DPS) [dungeon]; Faerie Mantle (5820, -0.21 DPS) [quest] |
| back | Caretaker's Cape (19533) | Silverwing Sentinels [rep] | 16.6 healing_power points (1.49 DPS) | yes | Prelacy Cape (7004, -0.10 DPS, sim-verified) [quest]; Glowing Thresher Cape (6901, -0.22 DPS) [dungeon]; Darkspear Raider's Cloak (272078, -0.24 DPS) [vendor] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 35.4 healing_power points (3.19 DPS) | yes | Pristine Gown (253961, -0.21 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -1.04 DPS) [dungeon]; Robes of Arugal (6324, -1.09 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 12.0 healing_power points (1.09 DPS) | yes | Nightsky Wristbands (6407, -0.09 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.18 DPS, sim-verified) [world_drop]; Stonecloth Bindings (14416, -0.46 DPS) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 19.2 healing_power points (1.73 DPS) | yes | Gloves of Old (9395, -0.15 DPS, sim-verified) [world_drop]; Hotshot Pilot's Gloves (9491, -0.32 DPS) [dungeon]; Pristine Gloves (253913, -0.36 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.6 healing_power points (1.50 DPS) | yes | Resilient Cord (14406, -0.42 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.44 DPS, sim-verified) [crafted]; Dreamer's Belt (4829, -0.45 DPS) [vendor] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 36.2 healing_power points (3.27 DPS) | yes | Filigreed Pristine Leggings (253937, -0.63 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -1.03 DPS) [crafted]; Necromancer Leggings (2277, -1.88 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 33.4 healing_power points (3.01 DPS) | yes | Nimbus Boots (6998, -1.62 DPS, sim-verified) [quest]; Acidic Walkers (9454, -1.68 DPS) [dungeon]; Frothing Slippers (254003, -1.80 DPS) [crafted] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.62 DPS) | yes | Runemender's Seal (273031, -0.52 DPS) [world]; The Queen's Jewel (13094, -0.73 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.87 DPS) [dungeon] |
| finger2 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 12.2 healing_power points (1.10 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; The Queen's Jewel (13094, -0.21 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.34 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-verified (64.8 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -0.72 DPS) [quest]; Death Speaker Scepter (2816, -2.24 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 9.0 healing_power points (0.81 DPS) | yes | Lesser Mystic Wand (11289, -0.30 DPS) [crafted]; Starfaller (13063, -0.30 DPS) [world_drop]; Summoner's Wand (5245, -0.44 DPS, sim-verified) [world] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Mantle of Honor; back: Caretaker's Cape; chest: Death Speaker Robes; hands: Truefaith Gloves; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Sea Giant's Toe Ring; finger2: Electrocutioner Lagnut; main_hand: Spine of the Devourer; ranged: Wand of Mana Concentration

No-known-source sample (15 of 258, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 40 (gnome, 000000000000000000-03505003030121431-00000000000000000)

Set DPS (verified): 113.8. Weights run: 11.2s. Verify run: 35.0s. 340 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.028, intellect=1.950 ± 0.019, spirit=1.133 ± 0.011, mp5=2.488 ± 0.020, crit=0.150 ± 0.008 per rating point (14 rating = 1%, 2.095 per %), spell_haste=-3.940 ± 0.250

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 57.3 healing_power points (5.30 DPS) | yes | Whitemane's Chapeau (7720, -1.22 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -1.62 DPS) [world_drop]; Earthen Silk Hood (254015, -2.22 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 28.6 healing_power points (2.64 DPS) | yes | Necklace of Calisea (1714, -1.14 DPS, sim-verified) [dungeon]; Darkspear Warding Pendant (272074, -1.38 DPS) [vendor]; Kaleidoscope Chain (13084, -1.50 DPS) [world_drop] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 40.3 healing_power points (3.73 DPS) | yes | Earthen Silk Shoulders (254033, -1.23 DPS) [crafted]; Windchaser Amice (14432, -1.39 DPS) [world_drop]; Mistscape Mantle (4734, -1.90 DPS, sim-verified) [dungeon] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 26.0 healing_power points (2.40 DPS) | yes | Caretaker's Cape (19532, -0.33 DPS, sim-verified) [rep]; Blackforge Cape (6424, -0.54 DPS) [dungeon]; Sergeant Major's Cape (16336, -0.69 DPS) [pvp] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 67.1 healing_power points (6.20 DPS) | yes | Pristine Gown (253961, -2.67 DPS) [crafted]; Silksand Tunic (14417, -2.79 DPS) [world_drop]; Death Speaker Robes (6682, -4.03 DPS, sim-verified) [dungeon] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 17.5 healing_power points (1.62 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Earthen Silk Cuffs (254019, -0.17 DPS) [crafted]; Mistscape Bracers (4045, -0.18 DPS) [dungeon] |
| hands | Gilded Handwraps (254021) | Tailoring [crafted] | 35.6 healing_power points (3.30 DPS) | yes | Stormcloth Gloves (10011, -0.36 DPS, sim-verified) [crafted]; Earthen Silk Gloves (254017, -0.79 DPS) [crafted]; Gloves of Old (9395, -1.14 DPS) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 44.4 healing_power points (4.10 DPS) | yes | Sutarn's Ring (13105, -1.67 DPS) [world_drop]; Windchaser Cinch (14435, -1.72 DPS) [world_drop]; Deathmage Sash (10771, -2.54 DPS, sim-verified) [dungeon] |
| legs | Stoneweaver Leggings (9407) | Uldaman: Ironaya [dungeon] | 42.2 healing_power points (3.90 DPS) | yes | Pristine Leggings (253987, -0.33 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.74 DPS) [crafted]; Aurora Pants (4044, -0.87 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 38.2 healing_power points (3.53 DPS) | yes | Furen's Boots (13100, -1.08 DPS, sim-verified) [world_drop]; Thoughtcast Boots (10578, -1.43 DPS) [dungeon]; Kodo Rustler Boots (15697, -1.56 DPS) [quest] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.66 DPS) | yes | Ogremind Ring (1993, -0.09 DPS) [world_drop]; Voodoo Band (1996, -0.09 DPS) [world]; Mindbender Loop (5009, -0.19 DPS) [world_drop] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | 17.7 healing_power points (1.63 DPS) | yes | Ogremind Ring (1993, -0.06 DPS) [world_drop]; Mindbender Loop (5009, -0.16 DPS) [world_drop]; Voodoo Band (1996, -0.23 DPS, sim-verified) [world] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (113.8 DPS) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (113.8 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Death Speaker Scepter (2816, -3.24 DPS) [dungeon]; Illusionary Rod (7713, -3.65 DPS) [dungeon] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 31.5 healing_power points (2.91 DPS) | yes | Beacon of Hope (9393, -1.20 DPS) [dungeon]; Aurora Sphere (7610, -1.23 DPS) [dungeon]; Orb of Souls (249395, -1.47 DPS, sim-verified) [crafted] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 15.1 healing_power points (1.40 DPS) | yes | Wand of Mana Concentration (273043, -0.21 DPS) [world]; Flash Wand (5248, -0.36 DPS) [quest]; Goblin Igniter (5253, -0.36 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Gilded Handwraps; waist: Gilded Cord; legs: Stoneweaver Leggings; finger2: Welken Ring; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Orb of Lorica; ranged: Jaina's Firestarter

No-known-source sample (15 of 340, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 50 (gnome, 025003000000000000-03505003030121431-00000000000000000)

Set DPS (verified): 221.7. Weights run: 10.9s. Verify run: 37.5s. 434 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.063, intellect=2.044 ± 0.030, spirit=1.608 ± 0.020, mp5=2.971 ± 0.025, crit=0.289 ± 0.017 per rating point (14 rating = 1%, 4.043 per %), spell_haste=not significant (0.460 ± 0.485)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Satin Cover (220896) | Captain Dirgehammer [vendor] | sim-verified (221.7 DPS) | yes | Papal Fez (9431, -0.76 DPS) [dungeon]; Stormcloth Headband (10032, -1.04 DPS) [crafted]; Gemburst Circlet (10751, -2.79 DPS, sim-verified) [quest] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 35.0 healing_power points (3.70 DPS) | yes | Glowing Eye of Mordresh (10769, -0.60 DPS) [dungeon]; Necklace of Calisea (1714, -1.00 DPS) [dungeon]; Darkmoon Necklace (19303, -1.04 DPS, sim-verified) [vendor] |
| shoulder | Knight-Lieutenant's Satin Pads (220894) | Captain Dirgehammer [vendor] | 54.3 healing_power points (5.73 DPS) | yes | Nethergeld Shoulders (254049, -0.43 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -1.34 DPS) [dungeon]; Rotgrip Mantle (17732, -1.85 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 36.7 healing_power points (3.86 DPS) | yes | Imperial Red Cloak (8248, -0.82 DPS) [world_drop]; Caretaker's Cape (19532, -1.12 DPS) [rep] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 70.9 healing_power points (7.47 DPS) | yes | Ghostweave Vest (14141, -1.57 DPS) [crafted]; Grimlok's Tribal Vestments (9415, -2.02 DPS) [dungeon]; Knight's Satin Armor (220892, -3.21 DPS, sim-verified) [vendor] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 42.3 healing_power points (4.47 DPS) | yes | Forgotten Wraps (9433, -1.88 DPS) [world_drop]; Bloodband Bracers (11469, -2.02 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -2.38 DPS, sim-verified) [quest] |
| hands | Virtuous Mitts (226950) | Mokvar [vendor] | 75.4 healing_power points (7.95 DPS) | yes | Raider Handwraps (272098, -0.90 DPS, sim-verified) [vendor]; Gilded Gloves (254095, -1.38 DPS) [crafted]; Virtuous Hands (226958, -3.12 DPS) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 50.8 healing_power points (5.36 DPS) | yes | Gilded Waistcord (254081, -0.36 DPS) [crafted]; Earthenweave Cord (254077, -0.75 DPS) [crafted]; Gilded Cord (254037, -2.27 DPS, sim-verified) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 71.7 healing_power points (7.56 DPS) | yes | Senior Designer's Pantaloons (11841, +0.00 DPS, sim-verified) [dungeon]; Knight's Satin Leggings (220893, -0.98 DPS) [vendor]; Dalewind Trousers (13008, -1.38 DPS) [world_drop] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 49.4 healing_power points (5.21 DPS) | yes | Sergeant Major's Satin Boots (220895, -0.51 DPS) [vendor]; Earthenweave Boots (254093, -0.60 DPS) [crafted]; Vinerot Sandals (17748, -3.05 DPS, sim-verified) [dungeon] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 31.6 healing_power points (3.33 DPS) | yes | Choking Band (11868, -1.12 DPS) [quest]; Darkspear Signet (272069, -1.13 DPS) [vendor]; Coldwater Ring (4550, -1.16 DPS) [quest] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | 21.5 healing_power points (2.26 DPS) | yes | Darkspear Signet (272069, -0.07 DPS) [vendor]; Coldwater Ring (4550, -0.09 DPS) [quest]; Choking Band (11868, -2.32 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, -0.78 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.12 DPS) [quest]; Thunderbrew's Boot Flask (744, -1.46 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -0.54 DPS) [quest]; Evonice's Landin' Pilla (18951, -2.81 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -8.36 DPS) [dungeon]; Hand of Righteousness (7721, -12.32 DPS) [dungeon]; Hammer of the Northern Wind (810, -42.58 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cairnstone Sliver (9654) | The Morrow Stone [quest] | 18.3 healing_power points (1.93 DPS) | yes | Jaina's Firestarter (13064, +0.00 DPS, sim-verified) [world_drop]; Wand of Mana Concentration (273043, -0.44 DPS) [world]; Goblin Igniter (5253, -0.55 DPS) [quest] |

**New at 50:** head: Knight-Lieutenant's Satin Cover; neck: Horizon Choker; shoulder: Knight-Lieutenant's Satin Pads; back: Darkspear Raider's Cloak; wrist: Nethergeld Cuffs; hands: Virtuous Mitts; waist: Dawnspire Cord; legs: Kilt of the Atal'ai Prophet; feet: Gilded Sandals; finger1: Eye of Adaegus; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Cairnstone Sliver

No-known-source sample (15 of 434, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60 (gnome, 025003031330000000-03505003030121431-00000000000000000)

Set DPS (verified): 430.9. Weights run: 9.3s. Verify run: 61.7s. 948 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.080, intellect=1.353 ± 0.025, spirit=1.201 ± 0.024, mp5=2.077 ± 0.030, crit=0.208 ± 0.014 per rating point (14 rating = 1%, 2.912 per %), spell_haste=-10.174 ± 0.544

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 78.1 healing_power points (15.26 DPS) | yes | Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Devout Crown (16693, -17.19 DPS, sim-verified) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 48.7 healing_power points (9.52 DPS) | yes | The All-Seeing Eye of Zuldazar (19594, -0.84 DPS) [quest]; Drake Tooth Necklace (21531, -1.27 DPS) [quest]; The Eye of Zuldazar (19593, -2.62 DPS, sim-verified) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 83.5 healing_power points (16.30 DPS) | yes | Lieutenant Commander's Satin Mantle (227119, -3.79 DPS) [pvp]; Field Marshal's Satin Mantle (231628, -5.16 DPS) [pvp]; Virtuous Mantle (226951, -8.90 DPS, sim-verified) [quest] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 55.5 healing_power points (10.84 DPS) | yes | Caretaker's Cape (19530, -3.89 DPS) [rep]; Drape of Recovery (272413, -4.42 DPS, sim-verified) [vendor] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 92.2 healing_power points (18.01 DPS) | yes | Field Marshal's Satin Tunic (231624, -0.49 DPS) [vendor]; Knight-Captain's Satin Tunic (227122, -2.95 DPS) [pvp]; Virtuous Robe (226945, -3.77 DPS, sim-verified) [quest] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-verified (430.9 DPS) | yes | Marshal's Satin Bracers (17606, -1.11 DPS) [pvp]; Nethergeld Cuffs (254061, -1.54 DPS) [crafted]; Bracers of Hope (22667, -13.03 DPS, sim-verified) [quest] |
| hands | Virtuous Mitts (226950) | Mokvar [vendor] | sim-verified (430.9 DPS) | yes | Raider Handwraps (272097, +0.00 DPS) [vendor]; Mooncloth Gloves (18409, -1.14 DPS) [crafted]; Desert Bloom Gloves (20717, -11.49 DPS, sim-verified) [quest] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 60.8 healing_power points (11.86 DPS) | yes | Virtuous Belt (226948, -0.92 DPS, sim-verified) [quest]; Penitent's Cinch (272394, -1.75 DPS) [vendor]; Marshal's Satin Sash (17609, -2.27 DPS) [pvp] |
| legs | Virtuous Skirt (226946) | Anthion's Parting Words [quest] | 74.3 healing_power points (14.50 DPS) | yes | Knight-Captain's Satin Legguards (227125, +0.00 DPS) [pvp]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Devout Skirt (16694, -15.49 DPS, sim-verified) [dungeon] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 79.8 healing_power points (15.58 DPS) | yes | Virtuous Sandals (226952, -3.94 DPS, sim-verified) [quest]; Mooncloth Boots (15802, -4.14 DPS) [crafted]; Marshal's Satin Walkers (231627, -5.15 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 57.6 healing_power points (11.24 DPS) | yes | Blessed Band of Light (272407, -3.31 DPS) [vendor]; Channeler's Ring (272406, -5.67 DPS) [vendor]; Band of the Hierophant (13096, -6.02 DPS) [world_drop] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 43.8 healing_power points (8.56 DPS) | yes | Blessed Band of Light (272407, -2.27 DPS, sim-verified) [vendor]; Channeler's Ring (272406, -2.99 DPS) [vendor]; Band of the Hierophant (13096, -3.34 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (430.9 DPS) | yes | Uther's Strength (11302, -2.57 DPS) [world_drop]; Ankh of Life (1713, -3.27 DPS) [world_drop]; Shard of the Splithooves (10659, -3.65 DPS) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (430.9 DPS) | yes | Ankh of Life (1713, -5.39 DPS) [world_drop]; Shard of the Splithooves (10659, -5.77 DPS) [quest]; Uther's Strength (11302, -7.11 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (430.9 DPS) | yes | Charstone Dirk (17710, -12.70 DPS) [dungeon]; Hand of Righteousness (7721, -18.83 DPS) [dungeon]; Hand of Edward the Odd (2243, -22.47 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 30.2 healing_power points (5.90 DPS) | yes | Jaina's Firestarter (13064, -3.61 DPS) [world_drop]; Captain Rackmore's Tiller (16789, -4.02 DPS) [quest]; Cairnstone Sliver (9654, -5.02 DPS, sim-verified) [quest] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Virtuous Bracers; waist: Wisdom of the Timbermaw; legs: Virtuous Skirt; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket1: Darkspear Voodoo Seal; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 948, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60, raid preset (gnome, 025003031330000000-03505003030121431-00000000000000000)

Set DPS (verified): 753.5. Weights run: 6.1s. Verify run: 28.9s. 948 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.114, intellect=0.778 ± 0.017, spirit=0.765 ± 0.017, mp5=1.396 ± 0.019, crit=0.276 ± 0.014 per rating point (14 rating = 1%, 3.862 per %), spell_haste=-3.191 ± 0.367

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | sim-verified (753.4 DPS) | yes | Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Devout Crown (16693, -14.05 DPS, sim-verified) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 39.8 healing_power points (16.34 DPS) | yes | The Eye of Zuldazar (19593, -2.33 DPS) [quest]; The All-Seeing Eye of Zuldazar (19594, -2.33 DPS) [quest]; Drake Tooth Necklace (21531, -2.78 DPS, sim-verified) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 75.0 healing_power points (30.77 DPS) | yes | Mooncloth Shoulders (14139, -8.63 DPS, sim-verified) [crafted]; Lieutenant Commander's Satin Mantle (227119, -9.83 DPS) [pvp]; Field Marshal's Satin Mantle (231628, -12.08 DPS) [pvp] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 49.8 healing_power points (20.42 DPS) | yes | Drape of Recovery (272413, -3.26 DPS, sim-verified) [vendor]; Caretaker's Cape (19530, -7.25 DPS) [rep] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 85.2 healing_power points (34.95 DPS) | yes | Field Marshal's Satin Tunic (231624, -6.38 DPS) [vendor]; Virtuous Robe (226945, -8.39 DPS) [quest]; Stormcloth Vest (10020, -8.66 DPS) [crafted] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-verified (753.4 DPS) | yes | Bracers of Hope (22667, +0.00 DPS) [quest]; Earthenweave Cuffs (254131, -1.55 DPS) [crafted]; Devout Bracers (16697, -14.68 DPS, sim-verified) [dungeon] |
| hands | Virtuous Mitts (226950) | Mokvar [vendor] | sim-verified (753.4 DPS) | yes | Desert Bloom Gloves (20717, +0.00 DPS) [quest]; Mooncloth Gloves (18409, -0.36 DPS) [crafted]; Devout Gloves (16692, -15.40 DPS, sim-verified) [dungeon] |
| waist | Virtuous Belt (226948) | Just Compensation [quest] | sim-verified (753.4 DPS) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Penitent's Cinch (272394, -0.30 DPS) [vendor]; Devout Belt (16696, -13.72 DPS, sim-verified) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-verified (753.4 DPS) | yes | Virtuous Skirt (226946, +0.00 DPS) [quest]; Knight-Captain's Satin Legguards (227125, +0.00 DPS) [pvp]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 65.6 healing_power points (26.90 DPS) | yes | Virtuous Sandals (226952, -3.27 DPS, sim-verified) [quest]; Mooncloth Boots (15802, -8.52 DPS) [crafted]; Marshal's Satin Walkers (231627, -9.26 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 49.0 healing_power points (20.09 DPS) | yes | Band of Piety (22681, -5.58 DPS) [quest]; Sea Giant's Toe Ring (274746, -12.71 DPS) [vendor]; Channeler's Ring (272406, -13.11 DPS) [vendor] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | 37.1 healing_power points (15.22 DPS) | yes | Band of Piety (22681, +0.00 DPS, sim-verified) [quest]; Sea Giant's Toe Ring (274746, -7.84 DPS) [vendor]; Channeler's Ring (272406, -8.24 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (753.4 DPS) | yes | Uther's Strength (11302, -1.21 DPS) [world_drop]; Ankh of Life (1713, -4.83 DPS) [world_drop]; Shard of the Splithooves (10659, -5.15 DPS) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (753.4 DPS) | yes | Evonice's Landin' Pilla (18951, -6.25 DPS, sim-verified) [quest]; Uther's Strength (11302, -9.84 DPS) [world_drop]; Ankh of Life (1713, -13.46 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (753.4 DPS) | yes | Charstone Dirk (17710, -21.37 DPS) [dungeon]; Hand of Edward the Odd (2243, -22.03 DPS, sim-verified) [world_drop]; Hand of Righteousness (7721, -32.54 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 25.4 healing_power points (10.42 DPS) | yes | Captain Rackmore's Tiller (16789, -4.47 DPS, sim-verified) [quest]; Cairnstone Sliver (9654, -7.26 DPS) [quest]; Jaina's Firestarter (13064, -7.57 DPS) [world_drop] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Virtuous Bracers; waist: Virtuous Belt; legs: Devout Skirt; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Blessed Band of Light; trinket1: Darkspear Voodoo Seal; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 948, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

## Horde

### Band 20 (troll, 000000000000000000-03503000000000000-00000000000000000)

Set DPS (verified): 38.8. Weights run: 10.2s. Verify run: 28.3s. 136 eligible items had no known source.

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

Set DPS (verified): 62.3. Weights run: 10.8s. Verify run: 33.6s. 244 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.013, intellect=1.400 ± 0.009, spirit=0.888 ± 0.004, mp5=1.683 ± 0.017, crit=0.093 ± 0.005 per rating point (14 rating = 1%, 1.303 per %), spell_haste=0.255 ± 0.053

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 38.3 healing_power points (3.45 DPS) | yes | Nightsky Cowl (4039, -1.54 DPS) [world_drop]; Filigreed Pristine Circlet (253975, -1.72 DPS, sim-verified) [crafted]; Resilient Cap (14401, -1.75 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 9.2 healing_power points (0.82 DPS) | yes | Scorn's Icy Choker (23169, -0.07 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -0.07 DPS) [world]; Crystal Starfire Medallion (5003, -0.08 DPS) [world_drop] |
| shoulder | Ghostly Mantle (3324) | Deathstalkers in Shadowfang [quest] | 17.0 healing_power points (1.53 DPS) | yes | Death Speaker Mantle (6685, -0.14 DPS) [dungeon]; Violet Sorcerer's Mantle (273051, -0.40 DPS) [world]; Nightsky Mantle (4718, -0.50 DPS, sim-verified) [world_drop] |
| back | Battle Healer's Cloak (19529) | Warsong Outriders [rep] | 16.6 healing_power points (1.49 DPS) | yes | Darkspear Raider's Cloak (272078, -0.24 DPS) [vendor]; Glowing Thresher Cape (6901, -0.56 DPS, sim-verified) [dungeon]; Resilient Cape (14400, -0.73 DPS) [world_drop] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 35.4 healing_power points (3.19 DPS) | yes | Pristine Gown (253961, -0.57 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -1.04 DPS) [dungeon]; Robes of Arugal (6324, -1.09 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 12.0 healing_power points (1.09 DPS) | yes | Glowing Magical Bracelets (13106, -0.06 DPS, sim-verified) [world_drop]; Nightsky Wristbands (6407, -0.09 DPS) [world_drop]; Featherbead Bracers (15452, -0.46 DPS) [quest] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 19.2 healing_power points (1.73 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.32 DPS) [dungeon]; Tattered Mittens (270030, -0.35 DPS) [quest]; Gloves of Old (9395, -0.55 DPS, sim-verified) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.6 healing_power points (1.50 DPS) | yes | Novice Ardent's Sash (253887, -0.34 DPS) [crafted]; Lilac Sash (6780, -0.35 DPS, sim-verified) [quest]; Resilient Cord (14406, -0.42 DPS) [world_drop] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 36.2 healing_power points (3.27 DPS) | yes | Filigreed Pristine Leggings (253937, -0.88 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -1.03 DPS) [crafted]; Necromancer Leggings (2277, -1.88 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 33.4 healing_power points (3.01 DPS) | yes | Frothing Slippers (254003, -1.80 DPS) [crafted]; Black Slippers (254007, -1.80 DPS) [crafted]; Acidic Walkers (9454, -2.14 DPS, sim-verified) [dungeon] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.62 DPS) | yes | Runemender's Seal (273031, -0.52 DPS) [world]; Solid Gold Band (279845, -0.59 DPS) [quest]; The Queen's Jewel (13094, -0.73 DPS) [world_drop] |
| finger2 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 12.2 healing_power points (1.10 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; Solid Gold Band (279845, -0.07 DPS) [quest]; The Queen's Jewel (13094, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-verified (62.3 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Advisor's Gnarled Staff (19569, -1.76 DPS) [pvp]; Death Speaker Scepter (2816, -1.99 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 9.0 healing_power points (0.81 DPS) | yes | Lesser Mystic Wand (11289, -0.30 DPS) [crafted]; Starfaller (13063, -0.30 DPS) [world_drop]; Summoner's Wand (5245, -0.46 DPS, sim-verified) [world] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Ghostly Mantle; back: Battle Healer's Cloak; chest: Death Speaker Robes; hands: Truefaith Gloves; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Sea Giant's Toe Ring; finger2: Electrocutioner Lagnut; main_hand: Spine of the Devourer; ranged: Wand of Mana Concentration

No-known-source sample (15 of 244, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (troll, 000000000000000000-03505003030121431-00000000000000000)

Set DPS (verified): 108.0. Weights run: 11.2s. Verify run: 35.6s. 323 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.028, intellect=1.950 ± 0.019, spirit=1.133 ± 0.011, mp5=2.488 ± 0.020, crit=0.150 ± 0.008 per rating point (14 rating = 1%, 2.095 per %), spell_haste=-3.940 ± 0.250

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 57.3 healing_power points (5.30 DPS) | yes | Whitemane's Chapeau (7720, -1.40 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -1.62 DPS) [world_drop]; Earthen Silk Hood (254015, -2.22 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 28.6 healing_power points (2.64 DPS) | yes | Necklace of Calisea (1714, -1.15 DPS, sim-verified) [dungeon]; Darkspear Warding Pendant (272074, -1.38 DPS) [vendor]; Kaleidoscope Chain (13084, -1.50 DPS) [world_drop] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 40.3 healing_power points (3.73 DPS) | yes | Earthen Silk Shoulders (254033, -1.23 DPS) [crafted]; Windchaser Amice (14432, -1.39 DPS) [world_drop]; Mistscape Mantle (4734, -2.04 DPS, sim-verified) [dungeon] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 26.0 healing_power points (2.40 DPS) | yes | Battle Healer's Cloak (19528, +0.00 DPS, sim-verified) [rep]; Blackforge Cape (6424, -0.54 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 67.1 healing_power points (6.20 DPS) | yes | Pristine Gown (253961, -2.67 DPS) [crafted]; Silksand Tunic (14417, -2.79 DPS) [world_drop]; Death Speaker Robes (6682, -3.93 DPS, sim-verified) [dungeon] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 17.5 healing_power points (1.62 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Dryad's Wrist Bindings (19597, -0.02 DPS) [pvp]; Earthen Silk Cuffs (254019, -0.17 DPS) [crafted] |
| hands | Gilded Handwraps (254021) | Tailoring [crafted] | 35.6 healing_power points (3.30 DPS) | yes | Stormcloth Gloves (10011, -0.39 DPS, sim-verified) [crafted]; Earthen Silk Gloves (254017, -0.79 DPS) [crafted]; Gloves of Old (9395, -1.14 DPS) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 44.4 healing_power points (4.10 DPS) | yes | Sutarn's Ring (13105, -1.67 DPS) [world_drop]; Windchaser Cinch (14435, -1.72 DPS) [world_drop]; Deathmage Sash (10771, -2.21 DPS, sim-verified) [dungeon] |
| legs | Stoneweaver Leggings (9407) | Uldaman: Ironaya [dungeon] | 42.2 healing_power points (3.90 DPS) | yes | Pristine Leggings (253987, -0.08 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.74 DPS) [crafted]; Aurora Pants (4044, -0.87 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 38.2 healing_power points (3.53 DPS) | yes | Furen's Boots (13100, -1.15 DPS, sim-verified) [world_drop]; Boots of the Maharishi (9658, -1.28 DPS) [quest]; Thoughtcast Boots (10578, -1.43 DPS) [dungeon] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.66 DPS) | yes | Ogremind Ring (1993, -0.09 DPS) [world_drop]; Voodoo Band (1996, -0.09 DPS) [world]; Mindbender Loop (5009, -0.19 DPS) [world_drop] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | 17.7 healing_power points (1.63 DPS) | yes | Ogremind Ring (1993, -0.06 DPS) [world_drop]; Voodoo Band (1996, -0.14 DPS, sim-verified) [world]; Mindbender Loop (5009, -0.16 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (108.0 DPS) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (108.0 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Death Speaker Scepter (2816, -3.24 DPS) [dungeon]; Illusionary Rod (7713, -3.65 DPS) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 19.8 healing_power points (1.83 DPS) | yes | Beacon of Hope (9393, -0.14 DPS, sim-verified) [dungeon]; Aurora Sphere (7610, -0.15 DPS) [dungeon]; Witch's Finger (16887, -0.39 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 15.1 healing_power points (1.40 DPS) | yes | Wand of Mana Concentration (273043, -0.25 DPS, sim-verified) [world]; Flash Wand (5248, -0.36 DPS) [quest]; Goblin Igniter (5253, -0.36 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Gilded Handwraps; waist: Gilded Cord; legs: Stoneweaver Leggings; finger2: Welken Ring; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Orb of Souls; ranged: Jaina's Firestarter

No-known-source sample (15 of 323, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16336 Sergeant Major's Cape

### Band 50 (troll, 025003000000000000-03505003030121431-00000000000000000)

Set DPS (verified): 211.6. Weights run: 10.9s. Verify run: 36.9s. 414 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.063, intellect=2.044 ± 0.030, spirit=1.608 ± 0.020, mp5=2.971 ± 0.025, crit=0.289 ± 0.017 per rating point (14 rating = 1%, 4.043 per %), spell_haste=not significant (0.460 ± 0.485)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | sim-verified (211.6 DPS) | yes | Blood Guard's Satin Cover (220899, +0.00 DPS) [vendor]; Stormcloth Headband (10032, -0.28 DPS) [crafted]; Gemburst Circlet (10751, -2.68 DPS, sim-verified) [quest] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 35.0 healing_power points (3.70 DPS) | yes | Darkmoon Necklace (19303, -0.52 DPS) [vendor]; Glowing Eye of Mordresh (10769, -0.60 DPS) [dungeon]; Necklace of Calisea (1714, -1.00 DPS) [dungeon] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | 52.0 healing_power points (5.48 DPS) | yes | Blood Guard's Satin Pads (220901, +0.00 DPS) [vendor]; Inquisitor's Shawl (19507, -0.89 DPS, sim-verified) [dungeon]; Rotgrip Mantle (17732, -1.60 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 36.7 healing_power points (3.86 DPS) | yes | Imperial Red Cloak (8248, -0.82 DPS) [world_drop]; Battle Healer's Cloak (19527, -0.95 DPS, sim-verified) [rep] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 70.9 healing_power points (7.47 DPS) | yes | Stone Guard's Satin Armor (220903, +0.00 DPS, sim-verified) [vendor]; Ghostweave Vest (14141, -1.57 DPS) [crafted]; Grimlok's Tribal Vestments (9415, -2.02 DPS) [dungeon] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 42.3 healing_power points (4.47 DPS) | yes | Forgotten Wraps (9433, -1.88 DPS) [world_drop]; Bloodband Bracers (11469, -2.02 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -3.79 DPS, sim-verified) [quest] |
| hands | Virtuous Mitts (226950) | Mokvar [vendor] | 75.4 healing_power points (7.95 DPS) | yes | Raider Handwraps (272098, +0.00 DPS, sim-verified) [vendor]; Gilded Gloves (254095, -1.38 DPS) [crafted]; Virtuous Hands (226958, -3.12 DPS) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 50.8 healing_power points (5.36 DPS) | yes | Gilded Waistcord (254081, -0.36 DPS) [crafted]; Earthenweave Cord (254077, -0.75 DPS) [crafted]; Gilded Cord (254037, -2.20 DPS, sim-verified) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 71.7 healing_power points (7.56 DPS) | yes | Senior Designer's Pantaloons (11841, -0.37 DPS) [dungeon]; Stone Guard's Satin Leggings (220902, -0.98 DPS) [vendor]; Dalewind Trousers (13008, -1.38 DPS) [world_drop] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 49.4 healing_power points (5.21 DPS) | yes | First Sergeant's Satin Boots (220900, -0.51 DPS) [vendor]; Earthenweave Boots (254093, -0.60 DPS) [crafted]; Vinerot Sandals (17748, -3.27 DPS, sim-verified) [dungeon] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 31.6 healing_power points (3.33 DPS) | yes | Darkspear Signet (272069, -1.13 DPS) [vendor]; Coldwater Ring (4550, -1.16 DPS) [quest]; Voodoo Band (1996, -1.31 DPS) [world] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | 21.5 healing_power points (2.26 DPS) | yes | Darkspear Signet (272069, +0.00 DPS, sim-verified) [vendor]; Coldwater Ring (4550, -0.09 DPS) [quest]; Voodoo Band (1996, -0.25 DPS) [world] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, -2.66 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -3.00 DPS) [quest]; Alchemist's Stone (13503, -4.02 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.20 DPS) [quest]; Alchemist's Stone (13503, -1.22 DPS) [crafted]; Ankh of Life (1713, -2.95 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -8.36 DPS) [dungeon]; Hand of Righteousness (7721, -12.32 DPS) [dungeon]; Hammer of the Northern Wind (810, -35.10 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 17.1 healing_power points (1.80 DPS) | yes | Wand of Mana Concentration (273043, -0.29 DPS, sim-verified) [world]; Flash Wand (5248, -0.43 DPS) [quest]; Goblin Igniter (5253, -0.43 DPS) [quest] |

**New at 50:** neck: Horizon Choker; shoulder: Nethergeld Shoulders; back: Darkspear Raider's Cloak; wrist: Nethergeld Cuffs; hands: Virtuous Mitts; waist: Dawnspire Cord; legs: Kilt of the Atal'ai Prophet; feet: Gilded Sandals; finger1: Eye of Adaegus; trinket1: Darkspear Voodoo Seal; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 414, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60 (troll, 025003031330000000-03505003030121431-00000000000000000)

Set DPS (verified): 419.6. Weights run: 9.3s. Verify run: 62.7s. 940 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.080, intellect=1.353 ± 0.025, spirit=1.201 ± 0.024, mp5=2.077 ± 0.030, crit=0.208 ± 0.014 per rating point (14 rating = 1%, 2.912 per %), spell_haste=-10.174 ± 0.544

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 78.1 healing_power points (15.26 DPS) | yes | Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Devout Crown (16693, -14.44 DPS, sim-verified) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 48.7 healing_power points (9.52 DPS) | yes | The All-Seeing Eye of Zuldazar (19594, -0.84 DPS) [quest]; Drake Tooth Necklace (21531, -1.27 DPS) [quest]; The Eye of Zuldazar (19593, -1.79 DPS, sim-verified) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 83.5 healing_power points (16.30 DPS) | yes | Champion's Satin Mantle (227120, -3.79 DPS) [pvp]; Warlord's Satin Mantle (231631, -5.16 DPS) [pvp]; Virtuous Mantle (226951, -8.75 DPS, sim-verified) [quest] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 55.5 healing_power points (10.84 DPS) | yes | Battle Healer's Cloak (19526, -3.89 DPS) [rep]; Drape of Recovery (272413, -4.42 DPS, sim-verified) [vendor] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 92.2 healing_power points (18.01 DPS) | yes | Warlord's Satin Tunic (231632, -0.49 DPS) [vendor]; Legionnaire's Satin Tunic (227124, -2.95 DPS) [pvp]; Virtuous Robe (226945, -2.96 DPS, sim-verified) [quest] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-verified (419.7 DPS) | yes | General's Satin Bracers (17619, -1.11 DPS) [pvp]; Nethergeld Cuffs (254061, -1.54 DPS) [crafted]; Bracers of Hope (22667, -11.43 DPS, sim-verified) [quest] |
| hands | Virtuous Mitts (226950) | Mokvar [vendor] | sim-verified (419.7 DPS) | yes | Raider Handwraps (272097, +0.00 DPS) [vendor]; Mooncloth Gloves (18409, -1.14 DPS) [crafted]; Desert Bloom Gloves (20717, -9.95 DPS, sim-verified) [quest] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 60.8 healing_power points (11.86 DPS) | yes | Virtuous Belt (226948, -1.23 DPS, sim-verified) [quest]; Penitent's Cinch (272394, -1.75 DPS) [vendor]; General's Satin Cinch (17621, -2.27 DPS) [pvp] |
| legs | Virtuous Skirt (226946) | Anthion's Parting Words [quest] | 74.3 healing_power points (14.50 DPS) | yes | Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Devout Skirt (16694, -13.17 DPS, sim-verified) [dungeon] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 79.8 healing_power points (15.58 DPS) | yes | Virtuous Sandals (226952, -4.03 DPS, sim-verified) [quest]; Mooncloth Boots (15802, -4.14 DPS) [crafted]; General's Satin Walkers (231630, -5.15 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 57.6 healing_power points (11.24 DPS) | yes | Blessed Band of Light (272407, -3.31 DPS) [vendor]; Channeler's Ring (272406, -5.67 DPS) [vendor]; Band of the Hierophant (13096, -6.02 DPS) [world_drop] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 43.8 healing_power points (8.56 DPS) | yes | Blessed Band of Light (272407, -1.36 DPS, sim-verified) [vendor]; Channeler's Ring (272406, -2.99 DPS) [vendor]; Band of the Hierophant (13096, -3.34 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (419.7 DPS) | yes | Uther's Strength (11302, -2.57 DPS) [world_drop]; Ankh of Life (1713, -3.27 DPS) [world_drop]; Shard of the Splithooves (10659, -3.65 DPS) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (419.7 DPS) | yes | Uther's Strength (11302, -4.69 DPS) [world_drop]; Shard of the Splithooves (10659, -5.77 DPS) [quest]; Ankh of Life (1713, -7.81 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (419.7 DPS) | yes | Charstone Dirk (17710, -12.70 DPS) [dungeon]; Hand of Righteousness (7721, -18.83 DPS) [dungeon]; Hand of Edward the Odd (2243, -22.87 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 30.2 healing_power points (5.90 DPS) | yes | Captain Rackmore's Tiller (16789, -4.02 DPS) [quest]; Wand of Mana Concentration (273043, -4.03 DPS) [world]; Jaina's Firestarter (13064, -4.80 DPS, sim-verified) [world_drop] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Virtuous Bracers; waist: Wisdom of the Timbermaw; legs: Virtuous Skirt; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 940, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60, raid preset (troll, 025003031330000000-03505003030121431-00000000000000000)

Set DPS (verified): 737.4. Weights run: 6.1s. Verify run: 30.5s. 940 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.114, intellect=0.778 ± 0.017, spirit=0.765 ± 0.017, mp5=1.396 ± 0.019, crit=0.276 ± 0.014 per rating point (14 rating = 1%, 3.862 per %), spell_haste=-3.191 ± 0.367

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | sim-verified (737.4 DPS) | yes | Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Devout Crown (16693, -16.29 DPS, sim-verified) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 39.8 healing_power points (16.34 DPS) | yes | The Eye of Zuldazar (19593, -2.33 DPS) [quest]; The All-Seeing Eye of Zuldazar (19594, -2.33 DPS) [quest]; Drake Tooth Necklace (21531, -2.58 DPS, sim-verified) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 75.0 healing_power points (30.77 DPS) | yes | Mooncloth Shoulders (14139, -8.86 DPS, sim-verified) [crafted]; Champion's Satin Mantle (227120, -9.83 DPS) [pvp]; Warlord's Satin Mantle (231631, -12.08 DPS) [pvp] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 49.8 healing_power points (20.42 DPS) | yes | Drape of Recovery (272413, -3.86 DPS, sim-verified) [vendor]; Battle Healer's Cloak (19526, -7.25 DPS) [rep] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 85.2 healing_power points (34.95 DPS) | yes | Warlord's Satin Tunic (231632, -6.38 DPS) [vendor]; Virtuous Robe (226945, -8.39 DPS) [quest]; Stormcloth Vest (10020, -8.66 DPS) [crafted] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-verified (737.4 DPS) | yes | Bracers of Hope (22667, +0.00 DPS) [quest]; Earthenweave Cuffs (254131, -1.55 DPS) [crafted]; Devout Bracers (16697, -16.20 DPS, sim-verified) [dungeon] |
| hands | Virtuous Mitts (226950) | Mokvar [vendor] | sim-verified (737.4 DPS) | yes | Desert Bloom Gloves (20717, +0.00 DPS) [quest]; Mooncloth Gloves (18409, -0.36 DPS) [crafted]; Devout Gloves (16692, -16.51 DPS, sim-verified) [dungeon] |
| waist | Virtuous Belt (226948) | Just Compensation [quest] | sim-verified (737.4 DPS) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Penitent's Cinch (272394, -0.30 DPS) [vendor]; Devout Belt (16696, -14.80 DPS, sim-verified) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-verified (737.4 DPS) | yes | Virtuous Skirt (226946, +0.00 DPS, sim-verified) [quest]; Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 65.6 healing_power points (26.90 DPS) | yes | Virtuous Sandals (226952, -2.70 DPS, sim-verified) [quest]; Mooncloth Boots (15802, -8.52 DPS) [crafted]; General's Satin Walkers (231630, -9.26 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 49.0 healing_power points (20.09 DPS) | yes | Band of Piety (22681, -5.58 DPS) [quest]; Sea Giant's Toe Ring (274746, -12.71 DPS) [vendor]; Channeler's Ring (272406, -13.11 DPS) [vendor] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | 37.1 healing_power points (15.22 DPS) | yes | Band of Piety (22681, +0.00 DPS, sim-verified) [quest]; Sea Giant's Toe Ring (274746, -7.84 DPS) [vendor]; Channeler's Ring (272406, -8.24 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (737.4 DPS) | yes | Uther's Strength (11302, -1.21 DPS) [world_drop]; Ankh of Life (1713, -4.83 DPS) [world_drop]; Shard of the Splithooves (10659, -5.15 DPS) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (737.4 DPS) | yes | Darkmoon Card: Blue Dragon (19288, -4.77 DPS, sim-verified) [quest]; Uther's Strength (11302, -9.84 DPS) [world_drop]; Ankh of Life (1713, -13.46 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (737.4 DPS) | yes | Charstone Dirk (17710, -21.37 DPS) [dungeon]; Hand of Edward the Odd (2243, -23.14 DPS, sim-verified) [world_drop]; Hand of Righteousness (7721, -32.54 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 25.4 healing_power points (10.42 DPS) | yes | Captain Rackmore's Tiller (16789, -5.61 DPS, sim-verified) [quest]; Jaina's Firestarter (13064, -7.57 DPS) [world_drop]; Earthen Rod (9381, -7.96 DPS) [dungeon] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Virtuous Bracers; waist: Virtuous Belt; legs: Devout Skirt; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Blessed Band of Light; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 940, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

