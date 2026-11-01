# Leveling BiS: Holy

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-03503000000000000-00000000000000000)

Set DPS (verified): 40.1. Weights run: 10.4s. Verify run: 28.3s. 163 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.006, intellect=1.334 ± 0.006, spirit=0.605 ± 0.003, mp5=1.556 ± 0.011, crit=0.070 ± 0.003 per rating point (14 rating = 1%, 0.982 per %), spell_haste=not significant (-0.005 ± 0.014)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.63 DPS) | yes | Shadow Goggles (4373, -0.98 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.41 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 2.4 healing_power points (0.22 DPS) | yes | Scholarly Pendant (277203, -0.17 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.0 healing_power points (1.09 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.53 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.67 DPS) [dungeon] |
| back | Sanguine Cape (14376) | World drop [world_drop] | 5.3 healing_power points (0.48 DPS) | yes | Pearl-clasped Cloak (5542, -0.12 DPS, sim-verified) [crafted]; Seer's Cape (6378, -0.13 DPS) [dungeon]; Forest Cloak (4710, -0.24 DPS) [world_drop] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 21.5 healing_power points (1.95 DPS) | yes | Corsair's Overshirt (5202, -0.43 DPS, sim-verified) [dungeon]; Robe of the Moccasin (6465, -0.58 DPS) [dungeon]; Bloody Apron (6226, -0.77 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 11.3 healing_power points (1.03 DPS) | yes | Repurposed Hair Band (281256, -0.68 DPS) [quest]; Mystic's Bracelets (14366, -0.79 DPS) [world_drop]; Bright Bracers (3647, -0.99 DPS, sim-verified) [world_drop] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 15.0 healing_power points (1.36 DPS) | yes | Magefist Gloves (12977, -0.49 DPS, sim-verified) [world_drop]; Tomb Robber's Gloves (280096, -0.63 DPS) [quest]; Bright Gloves (3066, -0.66 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.3 healing_power points (1.48 DPS) | yes | Novice Ardent's Sash (253887, -0.51 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.51 DPS) [world_drop]; Wise Man's Belt (4786, -0.88 DPS) [vendor] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 28.4 healing_power points (2.58 DPS) | yes | Darkweave Breeches (12987, -1.58 DPS, sim-verified) [world_drop]; Filigreed Silky Leggings (253939, -1.63 DPS) [crafted]; Filigreed Shadow Leggings (253943, -1.63 DPS) [crafted] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 13.0 healing_power points (1.18 DPS) | yes | Kimbra Boots (6191, -0.54 DPS, sim-verified) [quest]; Spidersilk Boots (4320, -0.69 DPS) [crafted]; Sanguine Sandals (14374, -0.69 DPS) [world_drop] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 8.0 healing_power points (0.73 DPS) | yes | Band of Purification (12996, -0.40 DPS) [world_drop]; Deep Fathom Ring (6463, -0.45 DPS) [dungeon]; Minor Channeling Ring (1449, -0.48 DPS) [quest] |
| finger2 | Volcanic Rock Ring (12053) | World drop [world_drop] | 4.0 healing_power points (0.36 DPS) | yes | Band of Purification (12996, +0.00 DPS, sim-verified) [world_drop]; Deep Fathom Ring (6463, -0.09 DPS) [dungeon]; Minor Channeling Ring (1449, -0.12 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 13.3 healing_power points (1.21 DPS) | yes | Staff of Westfall (2042, -0.19 DPS, sim-verified) [quest]; Lesser Staff of the Spire (1300, -0.48 DPS) [world]; Witching Stave (1484, -0.60 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Flaring Baton (5326) (or Moonstone Wand (15204)) | The Escape [quest] | 2.7 healing_power points (0.24 DPS) | yes | Moonstone Wand (15204, +0.00 DPS) [quest]; Sable Wand (7607, -0.08 DPS) [quest]; Elven Wand (5604, -0.12 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Tarnished Locket; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Volcanic Rock Ring; main_hand: Twisted Chanter's Staff; ranged: Flaring Baton

No-known-source sample (15 of 163, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes; 16607 Acolyte's Sacrificial Robes; 18851 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (gnome, 000000000000000000-03505003030110000-00000000000000000)

Set DPS (verified): 63.8. Weights run: 11.0s. Verify run: 33.7s. 294 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.013, intellect=1.400 ± 0.009, spirit=0.888 ± 0.004, mp5=1.683 ± 0.017, crit=0.093 ± 0.005 per rating point (14 rating = 1%, 1.303 per %), spell_haste=0.255 ± 0.053

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 38.3 healing_power points (3.45 DPS) | yes | Nightsky Cowl (4039, -1.54 DPS) [world_drop]; Resilient Cap (14401, -1.75 DPS) [world_drop]; Filigreed Pristine Circlet (253975, -1.98 DPS, sim-verified) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 9.2 healing_power points (0.82 DPS) | yes | Crystal Starfire Medallion (5003, -0.08 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.13 DPS, sim-verified) [world]; Pendant of Myzrael (4614, -0.34 DPS) [dungeon] |
| shoulder | Mantle of Honor (3560) | Bride of the Embalmer [quest] | 16.0 healing_power points (1.44 DPS) | yes | Nightsky Mantle (4718, +0.00 DPS, sim-verified) [world_drop]; Death Speaker Mantle (6685, -0.06 DPS) [dungeon]; Faerie Mantle (5820, -0.21 DPS) [quest] |
| back | Prelacy Cape (7004) | Researching the Corruption [quest] | 16.3 healing_power points (1.47 DPS) | yes | Glowing Thresher Cape (6901, -0.12 DPS, sim-verified) [dungeon]; Repairman's Cape (9605, -0.33 DPS) [quest]; Sergeant Major's Cape (16315, -0.65 DPS) [pvp] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 35.4 healing_power points (3.19 DPS) | yes | Pristine Gown (253961, -0.28 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -1.04 DPS) [dungeon]; Robes of Arugal (6324, -1.09 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 12.0 healing_power points (1.09 DPS) | yes | Nightsky Wristbands (6407, -0.09 DPS) [world_drop]; Stonecloth Bindings (14416, -0.46 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.51 DPS, sim-verified) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 19.2 healing_power points (1.73 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.32 DPS) [dungeon]; Pristine Gloves (253913, -0.36 DPS) [crafted]; Gloves of Old (9395, -0.40 DPS, sim-verified) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.6 healing_power points (1.50 DPS) | yes | Novice Ardent's Sash (253887, -0.28 DPS, sim-verified) [crafted]; Resilient Cord (14406, -0.42 DPS) [world_drop]; Dreamer's Belt (4829, -0.45 DPS) [vendor] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 36.2 healing_power points (3.27 DPS) | yes | Filigreed Pristine Leggings (253937, -0.86 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -1.03 DPS) [crafted]; Necromancer Leggings (2277, -1.88 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 33.4 healing_power points (3.01 DPS) | yes | Acidic Walkers (9454, -1.68 DPS) [dungeon]; Nimbus Boots (6998, -1.79 DPS, sim-verified) [quest]; Frothing Slippers (254003, -1.80 DPS) [crafted] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 12.2 healing_power points (1.10 DPS) | yes | The Queen's Jewel (13094, -0.21 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.34 DPS) [dungeon]; Defias Mage Ring (1077, -0.36 DPS) [world] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 12.2 healing_power points (1.10 DPS) | yes | The Queen's Jewel (13094, -0.18 DPS, sim-verified) [world_drop]; Lavishly Jeweled Ring (1156, -0.34 DPS) [dungeon]; Defias Mage Ring (1077, -0.36 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-verified (63.8 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -0.72 DPS) [quest]; Death Speaker Scepter (2816, -1.94 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 9.0 healing_power points (0.81 DPS) | yes | Lesser Mystic Wand (11289, -0.30 DPS) [crafted]; Starfaller (13063, -0.30 DPS) [world_drop]; Summoner's Wand (5245, -0.84 DPS, sim-verified) [world] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Mantle of Honor; back: Prelacy Cape; chest: Death Speaker Robes; hands: Truefaith Gloves; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Spine of the Devourer; ranged: Wand of Mana Concentration

No-known-source sample (15 of 294, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 40 (gnome, 000000000000000000-03505003030121431-00000000000000000)

Set DPS (verified): 109.1. Weights run: 11.4s. Verify run: 33.8s. 396 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.028, intellect=1.950 ± 0.019, spirit=1.133 ± 0.011, mp5=2.488 ± 0.020, crit=0.150 ± 0.008 per rating point (14 rating = 1%, 2.095 per %), spell_haste=-3.940 ± 0.250

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 57.3 healing_power points (5.30 DPS) | yes | Whitemane's Chapeau (7720, -1.08 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -1.62 DPS) [world_drop]; Earthen Silk Hood (254015, -2.22 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 28.6 healing_power points (2.64 DPS) | yes | Necklace of Calisea (1714, -1.06 DPS, sim-verified) [dungeon]; Kaleidoscope Chain (13084, -1.50 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.56 DPS) [world] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 40.3 healing_power points (3.73 DPS) | yes | Earthen Silk Shoulders (254033, -1.23 DPS) [crafted]; Windchaser Amice (14432, -1.39 DPS) [world_drop]; Mistscape Mantle (4734, -1.85 DPS, sim-verified) [dungeon] |
| back | Blackforge Cape (6424) | Uldaman: Ancient Treasure [dungeon] | 20.1 healing_power points (1.86 DPS) | yes | Sergeant Major's Cape (16336, -0.15 DPS) [pvp]; Prelacy Cape (7004, -0.22 DPS) [quest]; Windchaser Cloak (14430, -0.36 DPS) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 67.1 healing_power points (6.20 DPS) | yes | Pristine Gown (253961, -2.67 DPS) [crafted]; Silksand Tunic (14417, -2.79 DPS) [world_drop]; Death Speaker Robes (6682, -3.53 DPS, sim-verified) [dungeon] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 17.5 healing_power points (1.62 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Earthen Silk Cuffs (254019, -0.17 DPS) [crafted]; Mistscape Bracers (4045, -0.18 DPS) [dungeon] |
| hands | Gilded Handwraps (254021) | Tailoring [crafted] | 35.6 healing_power points (3.30 DPS) | yes | Stormcloth Gloves (10011, -0.29 DPS, sim-verified) [crafted]; Earthen Silk Gloves (254017, -0.79 DPS) [crafted]; Gloves of Old (9395, -1.14 DPS) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 44.4 healing_power points (4.10 DPS) | yes | Sutarn's Ring (13105, -1.67 DPS) [world_drop]; Windchaser Cinch (14435, -1.72 DPS) [world_drop]; Deathmage Sash (10771, -2.08 DPS, sim-verified) [dungeon] |
| legs | Stoneweaver Leggings (9407) | Uldaman: Ironaya [dungeon] | 42.2 healing_power points (3.90 DPS) | yes | Pristine Leggings (253987, -0.08 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.74 DPS) [crafted]; Aurora Pants (4044, -0.87 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 38.2 healing_power points (3.53 DPS) | yes | Furen's Boots (13100, -0.99 DPS, sim-verified) [world_drop]; Thoughtcast Boots (10578, -1.43 DPS) [dungeon]; Kodo Rustler Boots (15697, -1.56 DPS) [quest] |
| finger1 | Welken Ring (5011) | World drop [world_drop] | 17.7 healing_power points (1.63 DPS) | yes | Ogremind Ring (1993, -0.06 DPS) [world_drop]; Mindbender Loop (5009, -0.16 DPS) [world_drop]; Electrocutioner Lagnut (9447, -0.35 DPS) [dungeon] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 17.0 healing_power points (1.58 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Mindbender Loop (5009, -0.10 DPS) [world_drop]; Electrocutioner Lagnut (9447, -0.29 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (109.1 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Death Speaker Scepter (2816, -3.24 DPS) [dungeon]; Illusionary Rod (7713, -3.65 DPS) [dungeon] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 31.5 healing_power points (2.91 DPS) | yes | Orb of Souls (249395, -1.17 DPS, sim-verified) [crafted]; Beacon of Hope (9393, -1.20 DPS) [dungeon]; Aurora Sphere (7610, -1.23 DPS) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 15.1 healing_power points (1.40 DPS) | yes | Wand of Mana Concentration (273043, -0.21 DPS) [world]; Flash Wand (5248, -0.36 DPS) [quest]; Goblin Igniter (5253, -0.36 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Blackforge Cape; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Gilded Handwraps; waist: Gilded Cord; legs: Stoneweaver Leggings; finger1: Welken Ring; finger2: Voodoo Band; trinket1: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Orb of Lorica; ranged: Jaina's Firestarter

No-known-source sample (15 of 396, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 50 (gnome, 025003000000000000-03505003030121431-00000000000000000)

Set DPS (verified): 213.5. Weights run: 11.0s. Verify run: 35.2s. 515 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.063, intellect=2.044 ± 0.030, spirit=1.608 ± 0.020, mp5=2.971 ± 0.025, crit=0.289 ± 0.017 per rating point (14 rating = 1%, 4.043 per %), spell_haste=not significant (0.460 ± 0.485)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 72.7 healing_power points (7.66 DPS) | yes | Knight-Lieutenant's Satin Cover (220896, -0.24 DPS) [vendor]; Stormcloth Headband (10032, -1.28 DPS) [crafted]; Papal Fez (9431, -3.42 DPS, sim-verified) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 35.0 healing_power points (3.70 DPS) | yes | Necklace of Calisea (1714, -1.00 DPS) [dungeon]; Mindburst Medallion (11196, -1.55 DPS) [quest]; Glowing Eye of Mordresh (10769, -3.79 DPS, sim-verified) [dungeon] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | 52.0 healing_power points (5.48 DPS) | yes | Knight-Lieutenant's Satin Pads (220894, +0.00 DPS) [vendor]; Rotgrip Mantle (17732, -1.60 DPS) [dungeon]; Inquisitor's Shawl (19507, -3.08 DPS, sim-verified) [dungeon] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 28.9 healing_power points (3.05 DPS) | yes | Blackforge Cape (6424, -0.65 DPS) [dungeon]; Wingveil Cloak (10802, -0.66 DPS) [dungeon]; Highborne Cloak (14450, -3.35 DPS, sim-verified) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 70.9 healing_power points (7.47 DPS) | yes | Ghostweave Vest (14141, -1.57 DPS) [crafted]; Grimlok's Tribal Vestments (9415, -2.02 DPS) [dungeon]; Knight's Satin Armor (220892, -2.55 DPS, sim-verified) [vendor] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 42.3 healing_power points (4.47 DPS) | yes | Forgotten Wraps (9433, -1.88 DPS) [world_drop]; Bloodband Bracers (11469, -2.02 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -3.19 DPS, sim-verified) [quest] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 62.3 healing_power points (6.57 DPS) | yes | Sergeant Major's Satin Gloves (220897, -2.44 DPS) [vendor]; Quintis' Research Gloves (11888, -2.50 DPS) [quest]; Earthenweave Gloves (254075, -7.10 DPS, sim-verified) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 50.8 healing_power points (5.36 DPS) | yes | Gilded Waistcord (254081, -0.36 DPS) [crafted]; Earthenweave Cord (254077, -0.75 DPS) [crafted]; Gilded Cord (254037, -3.34 DPS, sim-verified) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 71.7 healing_power points (7.56 DPS) | yes | Knight's Satin Leggings (220893, -0.98 DPS) [vendor]; Dalewind Trousers (13008, -1.38 DPS) [world_drop]; Senior Designer's Pantaloons (11841, -3.33 DPS, sim-verified) [dungeon] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 49.4 healing_power points (5.21 DPS) | yes | Sergeant Major's Satin Boots (220895, -0.51 DPS) [vendor]; Earthenweave Boots (254093, -0.60 DPS) [crafted]; Vinerot Sandals (17748, -1.91 DPS, sim-verified) [dungeon] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 31.6 healing_power points (3.33 DPS) | yes | Choking Band (11868, -1.12 DPS) [quest]; Coldwater Ring (4550, -1.16 DPS) [quest]; Voodoo Band (1996, -1.31 DPS) [world] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | 21.5 healing_power points (2.26 DPS) | yes | Coldwater Ring (4550, -0.09 DPS) [quest]; Voodoo Band (1996, -0.25 DPS) [world]; Choking Band (11868, -2.85 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (213.5 DPS) | yes | Evonice's Landin' Pilla (18951, -0.34 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.68 DPS) [quest]; Alchemist's Stone (13503, -1.36 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (213.5 DPS) | yes | Evonice's Landin' Pilla (18951, -0.20 DPS) [quest]; Alchemist's Stone (13503, -1.22 DPS) [crafted]; Thunderbrew's Boot Flask (744, -1.54 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (213.5 DPS) | yes | Charstone Dirk (17710, -8.36 DPS) [dungeon]; Hand of Righteousness (7721, -12.32 DPS) [dungeon]; Hammer of the Northern Wind (810, -43.50 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cairnstone Sliver (9654) | The Morrow Stone [quest] | 18.3 healing_power points (1.93 DPS) | yes | Jaina's Firestarter (13064, -0.41 DPS, sim-verified) [world_drop]; Wand of Mana Concentration (273043, -0.44 DPS) [world]; Goblin Igniter (5253, -0.55 DPS) [quest] |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Nethergeld Shoulders; back: Imperial Red Cloak; wrist: Nethergeld Cuffs; hands: Gilded Gloves; waist: Dawnspire Cord; legs: Kilt of the Atal'ai Prophet; feet: Gilded Sandals; finger1: Eye of Adaegus; finger2: Welken Ring; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Cairnstone Sliver

No-known-source sample (15 of 515, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60 (gnome, 025003031330000000-03505003030121431-00000000000000000)

Set DPS (verified): 350.0. Weights run: 9.4s. Verify run: 63.0s. 1200 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.080, intellect=1.353 ± 0.025, spirit=1.201 ± 0.024, mp5=2.077 ± 0.030, crit=0.208 ± 0.014 per rating point (14 rating = 1%, 2.912 per %), spell_haste=-10.174 ± 0.544

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | 70.5 healing_power points (13.77 DPS) | yes | Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -6.81 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 23.8 healing_power points (4.65 DPS) | yes | Horizon Choker (13085, +0.00 DPS, sim-verified) [world_drop]; Lady Alizabeth's Pendant (13002, -1.14 DPS) [world_drop]; Necklace of Calisea (1714, -1.16 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | 50.2 healing_power points (9.81 DPS) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp]; Mooncloth Shoulders (14139, -5.81 DPS, sim-verified) [crafted] |
| back | Sergeant Major's Cape (16337) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 20.4 healing_power points (3.99 DPS) | yes | Ethereal Mist Cape (11873, -0.38 DPS) [quest]; Wingveil Cloak (10802, -0.43 DPS) [dungeon]; Imperial Red Cloak (8248, -1.74 DPS, sim-verified) [world_drop] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 92.2 healing_power points (18.01 DPS) | yes | Field Marshal's Satin Tunic (231624, -0.49 DPS) [vendor]; Knight-Captain's Satin Tunic (227122, -2.95 DPS) [pvp]; Mooncloth Vest (14138, -3.95 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (350.0 DPS) | yes | Marshal's Satin Bracers (17606, +0.00 DPS) [pvp]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -9.53 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (350.0 DPS) | yes | Knight-Lieutenant's Satin Handwraps (227128, -0.43 DPS) [pvp]; Marshal's Satin Handwraps (231623, -0.46 DPS) [vendor]; Gilded Gloves (254095, -8.79 DPS, sim-verified) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 60.8 healing_power points (11.86 DPS) | yes | Devout Belt (16696, +0.00 DPS, sim-verified) [dungeon]; Marshal's Satin Sash (17609, -2.27 DPS) [pvp]; Gilded Waistcord (254081, -3.82 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 71.9 healing_power points (14.04 DPS) | yes | Knight-Captain's Satin Legguards (227125, +0.00 DPS) [pvp]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Mooncloth Leggings (14137, -13.16 DPS, sim-verified) [crafted] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 58.6 healing_power points (11.43 DPS) | yes | Devout Sandals (16691, +0.00 DPS, sim-verified) [dungeon]; Marshal's Satin Walkers (231627, -1.01 DPS) [vendor]; Knight-Lieutenant's Satin Walkers (227129, -2.39 DPS) [pvp] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 26.7 healing_power points (5.22 DPS) | yes | Maiden's Circle (13001, -1.97 DPS) [world_drop]; Glowing Crystal Ring (18402, -2.05 DPS) [quest]; Choking Band (11868, -2.17 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 22.5 healing_power points (4.40 DPS) | yes | Glowing Crystal Ring (18402, -1.23 DPS) [quest]; Choking Band (11868, -1.35 DPS) [quest]; Maiden's Circle (13001, -2.67 DPS, sim-verified) [world_drop] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (350.0 DPS) | yes | Evonice's Landin' Pilla (18951, -0.09 DPS) [quest]; Ankh of Life (1713, -0.41 DPS, sim-verified) [world_drop]; Thunderbrew's Boot Flask (744, -0.56 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (350.0 DPS) | yes | Ankh of Life (1713, -0.49 DPS, sim-verified) [world_drop]; Evonice's Landin' Pilla (18951, -1.17 DPS) [quest]; Thunderbrew's Boot Flask (744, -1.64 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (350.0 DPS) | yes | Charstone Dirk (17710, -12.70 DPS) [dungeon]; Hand of Righteousness (7721, -18.83 DPS) [dungeon]; Hand of Edward the Odd (2243, -20.79 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 30.2 healing_power points (5.90 DPS) | yes | Jaina's Firestarter (13064, -3.61 DPS) [world_drop]; Captain Rackmore's Tiller (16789, -4.02 DPS) [quest]; Cairnstone Sliver (9654, -6.07 DPS, sim-verified) [quest] |

**New at 60:** head: Devout Crown; neck: Glowing Eye of Mordresh; shoulder: Devout Mantle; back: Sergeant Major's Cape; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Shard of the Splithooves; ranged: Torch of Light

No-known-source sample (15 of 1200, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60, raid preset (gnome, 025003031330000000-03505003030121431-00000000000000000)

Set DPS (verified): 673.6. Weights run: 6.2s. Verify run: 29.2s. 1200 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.114, intellect=0.778 ± 0.017, spirit=0.765 ± 0.017, mp5=1.396 ± 0.019, crit=0.276 ± 0.014 per rating point (14 rating = 1%, 3.862 per %), spell_haste=-3.191 ± 0.367

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (673.6 DPS) | yes | Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -11.57 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 19.2 healing_power points (7.88 DPS) | yes | Horizon Choker (13085, -1.33 DPS, sim-verified) [world_drop]; Lady Alizabeth's Pendant (13002, -3.18 DPS) [world_drop]; Necklace of Calisea (1714, -3.46 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (673.6 DPS) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp]; Mooncloth Shoulders (14139, -11.33 DPS, sim-verified) [crafted] |
| back | Prelacy Cape (7004) | Researching the Corruption [quest] | 15.6 healing_power points (6.39 DPS) | yes | Glowing Thresher Cape (6901, -1.01 DPS) [dungeon]; Wingveil Cloak (10802, -1.06 DPS, sim-verified) [dungeon]; Sergeant Major's Cape (16337, -1.33 DPS) [pvp] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 85.2 healing_power points (34.95 DPS) | yes | Field Marshal's Satin Tunic (231624, -6.38 DPS) [vendor]; Stormcloth Vest (10020, -7.97 DPS, sim-verified) [crafted]; Knight-Captain's Satin Tunic (227122, -9.37 DPS) [pvp] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (673.6 DPS) | yes | Nethergeld Cuffs (254061, +0.00 DPS) [crafted]; Marshal's Satin Bracers (17606, -1.00 DPS) [pvp]; Earthenweave Cuffs (254131, -12.05 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (673.6 DPS) | yes | Knight-Lieutenant's Satin Handwraps (227128, +0.00 DPS) [pvp]; Marshal's Satin Handwraps (231623, +0.00 DPS) [vendor]; Gilded Gloves (254095, -11.01 DPS, sim-verified) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 48.8 healing_power points (20.00 DPS) | yes | Earthenweave Cord (254077, -5.89 DPS) [crafted]; Gilded Waistcord (254081, -5.92 DPS, sim-verified) [crafted]; Devout Belt (16696, -6.29 DPS) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 53.3 healing_power points (21.84 DPS) | yes | Knight-Captain's Satin Legguards (227125, +0.00 DPS) [pvp]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Earthenweave Leggings (254125, -16.86 DPS, sim-verified) [crafted] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 44.8 healing_power points (18.39 DPS) | yes | Marshal's Satin Walkers (231627, -0.75 DPS) [vendor]; Gilded Sandals (254107, -2.80 DPS) [crafted]; Swarmtender's Footpads (275609, -3.53 DPS, sim-verified) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 16.2 healing_power points (6.64 DPS) | yes | Electrocutioner Lagnut (9447, -1.99 DPS) [dungeon]; Runemender's Seal (273031, -1.99 DPS) [world]; Maiden's Circle (13001, -2.53 DPS) [world_drop] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 13.8 healing_power points (5.68 DPS) | yes | Runemender's Seal (273031, -1.02 DPS) [world]; Maiden's Circle (13001, -1.56 DPS) [world_drop]; Electrocutioner Lagnut (9447, -2.74 DPS, sim-verified) [dungeon] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (673.6 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.30 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.93 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (673.6 DPS) | yes | Ankh of Life (1713, -3.17 DPS, sim-verified) [world_drop]; Evonice's Landin' Pilla (18951, -4.25 DPS) [quest]; Thunderbrew's Boot Flask (744, -4.87 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (673.6 DPS) | yes | Charstone Dirk (17710, -21.37 DPS) [dungeon]; Hand of Edward the Odd (2243, -28.73 DPS, sim-verified) [world_drop]; Hand of Righteousness (7721, -32.54 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 25.4 healing_power points (10.42 DPS) | yes | Captain Rackmore's Tiller (16789, -6.38 DPS, sim-verified) [quest]; Cairnstone Sliver (9654, -7.26 DPS) [quest]; Jaina's Firestarter (13064, -7.57 DPS) [world_drop] |

**New at 60:** head: Devout Crown; neck: Glowing Eye of Mordresh; shoulder: Devout Mantle; back: Prelacy Cape; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Shard of the Splithooves; ranged: Torch of Light

No-known-source sample (15 of 1200, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

## Horde

### Band 20 (troll, 000000000000000000-03503000000000000-00000000000000000)

Set DPS (verified): 38.2. Weights run: 10.4s. Verify run: 28.7s. 156 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.006, intellect=1.334 ± 0.006, spirit=0.605 ± 0.003, mp5=1.556 ± 0.011, crit=0.070 ± 0.003 per rating point (14 rating = 1%, 0.982 per %), spell_haste=not significant (-0.005 ± 0.014)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.63 DPS) | yes | Shadow Goggles (4373, -0.78 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.41 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 1.8 healing_power points (0.16 DPS) | yes | Scholarly Pendant (277203, -0.19 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.0 healing_power points (1.09 DPS) | yes | Slime-encrusted Pads (6461, -0.67 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -0.80 DPS, sim-verified) [crafted] |
| back | Sanguine Cape (14376) | World drop [world_drop] | 5.3 healing_power points (0.48 DPS) | yes | Seer's Cape (6378, -0.13 DPS) [dungeon]; Forest Cloak (4710, -0.24 DPS) [world_drop]; Pearl-clasped Cloak (5542, -0.31 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 21.5 healing_power points (1.95 DPS) | yes | Corsair's Overshirt (5202, -0.43 DPS, sim-verified) [dungeon]; Robe of the Moccasin (6465, -0.58 DPS) [dungeon]; Bloody Apron (6226, -0.77 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 11.3 healing_power points (1.03 DPS) | yes | Bright Bracers (3647, -0.54 DPS) [world_drop]; Crystalline Cuffs (14148, -0.62 DPS) [dungeon]; Featherbead Bracers (15452, -0.73 DPS, sim-verified) [quest] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 15.0 healing_power points (1.36 DPS) | yes | Blight Gloves (279877, -0.29 DPS, sim-verified) [quest]; Magefist Gloves (12977, -0.54 DPS) [world_drop]; Bright Gloves (3066, -0.66 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.3 healing_power points (1.48 DPS) | yes | Keller's Girdle (2911, -0.51 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.61 DPS, sim-verified) [crafted]; Wise Man's Belt (4786, -0.88 DPS) [vendor] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 28.4 healing_power points (2.58 DPS) | yes | Darkweave Breeches (12987, -1.49 DPS, sim-verified) [world_drop]; Filigreed Silky Leggings (253939, -1.63 DPS) [crafted]; Filigreed Shadow Leggings (253943, -1.63 DPS) [crafted] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 13.0 healing_power points (1.18 DPS) | yes | Spidersilk Boots (4320, -0.69 DPS) [crafted]; Bluegill Sandals (1560, -0.72 DPS) [world]; Sanguine Sandals (14374, -0.88 DPS, sim-verified) [world_drop] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 8.0 healing_power points (0.73 DPS) | yes | Volcanic Rock Ring (12053, -0.36 DPS) [world_drop]; Band of Purification (12996, -0.40 DPS) [world_drop]; Deep Fathom Ring (6463, -0.45 DPS) [dungeon] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 6.7 healing_power points (0.60 DPS) | yes | Band of Purification (12996, -0.28 DPS) [world_drop]; Deep Fathom Ring (6463, -0.33 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.37 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 13.3 healing_power points (1.21 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Staff of Orgrimmar (15444, -0.17 DPS) [quest]; Lesser Staff of the Spire (1300, -0.48 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Decay (5252) (or Flaring Baton (5326)) | Beren's Peril [quest] | 2.7 healing_power points (0.24 DPS) | yes | Flaring Baton (5326, +0.00 DPS) [quest]; Wisesight Wand (286750, -0.12 DPS) [world] |

**New at 20:** head: Pristine Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Gnarled Necromancer's Staff; ranged: Wand of Decay

No-known-source sample (15 of 156, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (troll, 000000000000000000-03505003030110000-00000000000000000)

Set DPS (verified): 61.0. Weights run: 11.0s. Verify run: 34.0s. 286 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.013, intellect=1.400 ± 0.009, spirit=0.888 ± 0.004, mp5=1.683 ± 0.017, crit=0.093 ± 0.005 per rating point (14 rating = 1%, 1.303 per %), spell_haste=0.255 ± 0.053

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 38.3 healing_power points (3.45 DPS) | yes | Nightsky Cowl (4039, -1.54 DPS) [world_drop]; Filigreed Pristine Circlet (253975, -1.67 DPS, sim-verified) [crafted]; Resilient Cap (14401, -1.75 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 9.2 healing_power points (0.82 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.07 DPS) [world]; Crystal Starfire Medallion (5003, -0.08 DPS) [world_drop]; Pendant of Myzrael (4614, -0.34 DPS) [dungeon] |
| shoulder | Ghostly Mantle (3324) | Deathstalkers in Shadowfang [quest] | 17.0 healing_power points (1.53 DPS) | yes | Nightsky Mantle (4718, +0.00 DPS, sim-verified) [world_drop]; Death Speaker Mantle (6685, -0.14 DPS) [dungeon]; Violet Sorcerer's Mantle (273051, -0.40 DPS) [world] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 14.1 healing_power points (1.27 DPS) | yes | Resilient Cape (14400, -0.33 DPS, sim-verified) [world_drop]; Silver-thread Cloak (4713, -0.61 DPS) [world_drop]; Soft Willow Cape (16661, -0.64 DPS) [quest] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 35.4 healing_power points (3.19 DPS) | yes | Pristine Gown (253961, +0.00 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -1.04 DPS) [dungeon]; Robes of Arugal (6324, -1.09 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 12.0 healing_power points (1.09 DPS) | yes | Nightsky Wristbands (6407, -0.09 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.42 DPS, sim-verified) [world_drop]; Featherbead Bracers (15452, -0.46 DPS) [quest] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 19.2 healing_power points (1.73 DPS) | yes | Gloves of Old (9395, -0.14 DPS) [world_drop]; Hotshot Pilot's Gloves (9491, -0.32 DPS) [dungeon]; Tattered Mittens (270030, -0.35 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.6 healing_power points (1.50 DPS) | yes | Lilac Sash (6780, +0.00 DPS, sim-verified) [quest]; Novice Ardent's Sash (253887, -0.34 DPS) [crafted]; Resilient Cord (14406, -0.42 DPS) [world_drop] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 36.2 healing_power points (3.27 DPS) | yes | Filigreed Pristine Leggings (253937, -0.69 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -1.03 DPS) [crafted]; Necromancer Leggings (2277, -1.88 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 33.4 healing_power points (3.01 DPS) | yes | Frothing Slippers (254003, -1.80 DPS) [crafted]; Black Slippers (254007, -1.80 DPS) [crafted]; Acidic Walkers (9454, -2.06 DPS, sim-verified) [dungeon] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 12.2 healing_power points (1.10 DPS) | yes | Solid Gold Band (279845, -0.07 DPS) [quest]; The Queen's Jewel (13094, -0.21 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.34 DPS) [dungeon] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 12.2 healing_power points (1.10 DPS) | yes | Solid Gold Band (279845, -0.07 DPS) [quest]; The Queen's Jewel (13094, -0.21 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.34 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-verified (61.0 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Glimmering Staff (249392, -1.86 DPS) [crafted]; Death Speaker Scepter (2816, -2.08 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 9.0 healing_power points (0.81 DPS) | yes | Lesser Mystic Wand (11289, -0.30 DPS) [crafted]; Starfaller (13063, -0.30 DPS) [world_drop]; Summoner's Wand (5245, -0.53 DPS, sim-verified) [world] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Ghostly Mantle; back: Glowing Thresher Cape; chest: Death Speaker Robes; hands: Truefaith Gloves; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Spine of the Devourer; ranged: Wand of Mana Concentration

No-known-source sample (15 of 286, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (troll, 000000000000000000-03505003030121431-00000000000000000)

Set DPS (verified): 103.5. Weights run: 11.4s. Verify run: 34.4s. 388 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.028, intellect=1.950 ± 0.019, spirit=1.133 ± 0.011, mp5=2.488 ± 0.020, crit=0.150 ± 0.008 per rating point (14 rating = 1%, 2.095 per %), spell_haste=-3.940 ± 0.250

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 57.3 healing_power points (5.30 DPS) | yes | Whitemane's Chapeau (7720, -1.02 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -1.62 DPS) [world_drop]; Earthen Silk Hood (254015, -2.22 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 28.6 healing_power points (2.64 DPS) | yes | Necklace of Calisea (1714, -0.99 DPS, sim-verified) [dungeon]; Kaleidoscope Chain (13084, -1.50 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.56 DPS) [world] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 40.3 healing_power points (3.73 DPS) | yes | Earthen Silk Shoulders (254033, -1.23 DPS) [crafted]; Windchaser Amice (14432, -1.39 DPS) [world_drop]; Mistscape Mantle (4734, -1.80 DPS, sim-verified) [dungeon] |
| back | Blackforge Cape (6424) | Uldaman: Ancient Treasure [dungeon] | 20.1 healing_power points (1.86 DPS) | yes | Glowing Thresher Cape (6901, -0.38 DPS) [dungeon]; Windchaser Cloak (14430, -0.48 DPS, sim-verified) [world_drop]; Cloak of Blight (6832, -0.50 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 67.1 healing_power points (6.20 DPS) | yes | Pristine Gown (253961, -2.67 DPS) [crafted]; Silksand Tunic (14417, -2.79 DPS) [world_drop]; Death Speaker Robes (6682, -3.55 DPS, sim-verified) [dungeon] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 17.5 healing_power points (1.62 DPS) | yes | Mindthrust Bracers (1974, -0.03 DPS) [dungeon]; Earthen Silk Cuffs (254019, -0.17 DPS) [crafted]; Mistscape Bracers (4045, -0.18 DPS) [dungeon] |
| hands | Gilded Handwraps (254021) | Tailoring [crafted] | 35.6 healing_power points (3.30 DPS) | yes | Stormcloth Gloves (10011, -0.43 DPS, sim-verified) [crafted]; Earthen Silk Gloves (254017, -0.79 DPS) [crafted]; Gloves of Old (9395, -1.14 DPS) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 44.4 healing_power points (4.10 DPS) | yes | Sutarn's Ring (13105, -1.67 DPS) [world_drop]; Windchaser Cinch (14435, -1.72 DPS) [world_drop]; Deathmage Sash (10771, -2.15 DPS, sim-verified) [dungeon] |
| legs | Stoneweaver Leggings (9407) | Uldaman: Ironaya [dungeon] | 42.2 healing_power points (3.90 DPS) | yes | Pristine Leggings (253987, -0.22 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.74 DPS) [crafted]; Aurora Pants (4044, -0.87 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 38.2 healing_power points (3.53 DPS) | yes | Furen's Boots (13100, -1.08 DPS, sim-verified) [world_drop]; Boots of the Maharishi (9658, -1.28 DPS) [quest]; Thoughtcast Boots (10578, -1.43 DPS) [dungeon] |
| finger1 | Welken Ring (5011) | World drop [world_drop] | 17.7 healing_power points (1.63 DPS) | yes | Ogremind Ring (1993, -0.06 DPS) [world_drop]; Mindbender Loop (5009, -0.16 DPS) [world_drop]; Electrocutioner Lagnut (9447, -0.35 DPS) [dungeon] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 17.0 healing_power points (1.58 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Mindbender Loop (5009, -0.10 DPS) [world_drop]; Electrocutioner Lagnut (9447, -0.29 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (103.5 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Death Speaker Scepter (2816, -3.24 DPS) [dungeon]; Illusionary Rod (7713, -3.65 DPS) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 19.8 healing_power points (1.83 DPS) | yes | Beacon of Hope (9393, -0.14 DPS, sim-verified) [dungeon]; Aurora Sphere (7610, -0.15 DPS) [dungeon]; Witch's Finger (16887, -0.39 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 15.1 healing_power points (1.40 DPS) | yes | Wand of Mana Concentration (273043, -0.10 DPS, sim-verified) [world]; Flash Wand (5248, -0.36 DPS) [quest]; Goblin Igniter (5253, -0.36 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Blackforge Cape; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Gilded Handwraps; waist: Gilded Cord; legs: Stoneweaver Leggings; finger1: Welken Ring; finger2: Voodoo Band; trinket1: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Orb of Souls; ranged: Jaina's Firestarter

No-known-source sample (15 of 388, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16336 Sergeant Major's Cape

### Band 50 (troll, 025003000000000000-03505003030121431-00000000000000000)

Set DPS (verified): 202.8. Weights run: 11.0s. Verify run: 35.8s. 507 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.063, intellect=2.044 ± 0.030, spirit=1.608 ± 0.020, mp5=2.971 ± 0.025, crit=0.289 ± 0.017 per rating point (14 rating = 1%, 4.043 per %), spell_haste=not significant (0.460 ± 0.485)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 72.7 healing_power points (7.66 DPS) | yes | Blood Guard's Satin Cover (220899, -0.24 DPS) [vendor]; Stormcloth Headband (10032, -1.28 DPS) [crafted]; Papal Fez (9431, -5.11 DPS, sim-verified) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 35.0 healing_power points (3.70 DPS) | yes | Necklace of Calisea (1714, -1.00 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -1.08 DPS, sim-verified) [dungeon]; Choker of the High Shaman (4112, -1.26 DPS) [quest] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | 52.0 healing_power points (5.48 DPS) | yes | Blood Guard's Satin Pads (220901, +0.00 DPS) [vendor]; Rotgrip Mantle (17732, -1.60 DPS) [dungeon]; Inquisitor's Shawl (19507, -2.82 DPS, sim-verified) [dungeon] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 28.9 healing_power points (3.05 DPS) | yes | Blackforge Cape (6424, -0.65 DPS) [dungeon]; Wingveil Cloak (10802, -0.66 DPS) [dungeon]; Highborne Cloak (14450, -0.66 DPS, sim-verified) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 70.9 healing_power points (7.47 DPS) | yes | Ghostweave Vest (14141, -1.57 DPS) [crafted]; Grimlok's Tribal Vestments (9415, -2.02 DPS) [dungeon]; Stone Guard's Satin Armor (220903, -2.22 DPS, sim-verified) [vendor] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 42.3 healing_power points (4.47 DPS) | yes | Forgotten Wraps (9433, -1.88 DPS) [world_drop]; Bloodband Bracers (11469, -2.02 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -2.85 DPS, sim-verified) [quest] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 62.3 healing_power points (6.57 DPS) | yes | Earthenweave Gloves (254075, -1.96 DPS) [crafted]; First Sergeant's Satin Gloves (220898, -2.44 DPS) [vendor]; Greenleaf Handwraps (19116, -8.13 DPS, sim-verified) [quest] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 50.8 healing_power points (5.36 DPS) | yes | Gilded Cord (254037, -0.30 DPS) [crafted]; Gilded Waistcord (254081, -0.36 DPS) [crafted]; Earthenweave Cord (254077, -0.75 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 71.7 healing_power points (7.56 DPS) | yes | Senior Designer's Pantaloons (11841, +0.00 DPS, sim-verified) [dungeon]; Stone Guard's Satin Leggings (220902, -0.98 DPS) [vendor]; Dalewind Trousers (13008, -1.38 DPS) [world_drop] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 49.4 healing_power points (5.21 DPS) | yes | First Sergeant's Satin Boots (220900, -0.51 DPS) [vendor]; Earthenweave Boots (254093, -0.60 DPS) [crafted]; Vinerot Sandals (17748, -1.03 DPS, sim-verified) [dungeon] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 31.6 healing_power points (3.33 DPS) | yes | Coldwater Ring (4550, -1.16 DPS) [quest]; Ogremind Ring (1993, -1.31 DPS) [world_drop]; Voodoo Band (1996, -1.31 DPS) [world] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | 21.5 healing_power points (2.26 DPS) | yes | Ogremind Ring (1993, -0.25 DPS) [world_drop]; Voodoo Band (1996, -0.25 DPS) [world]; Coldwater Ring (4550, -0.32 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (202.8 DPS) | yes | Evonice's Landin' Pilla (18951, -0.56 DPS, sim-verified) [quest]; Alchemist's Stone (13503, -1.36 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (202.8 DPS) | yes | Evonice's Landin' Pilla (18951, -0.82 DPS, sim-verified) [quest]; Alchemist's Stone (13503, -1.22 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (202.8 DPS) | yes | Charstone Dirk (17710, -8.36 DPS) [dungeon]; Hand of Righteousness (7721, -12.32 DPS) [dungeon]; Hammer of the Northern Wind (810, -40.64 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 17.1 healing_power points (1.80 DPS) | yes | Wand of Mana Concentration (273043, -0.39 DPS, sim-verified) [world]; Flash Wand (5248, -0.43 DPS) [quest]; Goblin Igniter (5253, -0.43 DPS) [quest] |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Nethergeld Shoulders; back: Imperial Red Cloak; wrist: Nethergeld Cuffs; hands: Gilded Gloves; waist: Dawnspire Cord; legs: Kilt of the Atal'ai Prophet; feet: Gilded Sandals; finger1: Eye of Adaegus; finger2: Welken Ring; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 507, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60 (troll, 025003031330000000-03505003030121431-00000000000000000)

Set DPS (verified): 339.2. Weights run: 9.4s. Verify run: 62.6s. 1193 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.080, intellect=1.353 ± 0.025, spirit=1.201 ± 0.024, mp5=2.077 ± 0.030, crit=0.208 ± 0.014 per rating point (14 rating = 1%, 2.912 per %), spell_haste=-10.174 ± 0.544

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | 70.5 healing_power points (13.77 DPS) | yes | Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -6.59 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 23.8 healing_power points (4.65 DPS) | yes | Horizon Choker (13085, -0.69 DPS, sim-verified) [world_drop]; Lady Alizabeth's Pendant (13002, -1.14 DPS) [world_drop]; Necklace of Calisea (1714, -1.16 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | 50.2 healing_power points (9.81 DPS) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp]; Mooncloth Shoulders (14139, -6.89 DPS, sim-verified) [crafted] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 19.7 healing_power points (3.85 DPS) | yes | Wingveil Cloak (10802, +0.00 DPS, sim-verified) [dungeon]; Highborne Cloak (14450, -0.53 DPS) [world_drop]; Glowing Thresher Cape (6901, -0.60 DPS) [dungeon] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 92.2 healing_power points (18.01 DPS) | yes | Warlord's Satin Tunic (231632, -0.49 DPS) [vendor]; Legionnaire's Satin Tunic (227124, -2.95 DPS) [pvp]; Mooncloth Vest (14138, -6.84 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (339.2 DPS) | yes | General's Satin Bracers (17619, +0.00 DPS) [pvp]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -9.07 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (339.2 DPS) | yes | Blood Guard's Satin Handwraps (227126, -0.43 DPS) [pvp]; General's Satin Handwraps (231633, -0.46 DPS) [vendor]; Gilded Gloves (254095, -8.78 DPS, sim-verified) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 60.8 healing_power points (11.86 DPS) | yes | Devout Belt (16696, +0.00 DPS, sim-verified) [dungeon]; General's Satin Cinch (17621, -2.27 DPS) [pvp]; Gilded Waistcord (254081, -3.82 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 71.9 healing_power points (14.04 DPS) | yes | Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Mooncloth Leggings (14137, -11.59 DPS, sim-verified) [crafted] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 58.6 healing_power points (11.43 DPS) | yes | Devout Sandals (16691, +0.00 DPS, sim-verified) [dungeon]; General's Satin Walkers (231630, -1.01 DPS) [vendor]; Blood Guard's Satin Walkers (227127, -2.39 DPS) [pvp] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 26.7 healing_power points (5.22 DPS) | yes | Chemist's Ring (15702, -1.64 DPS) [quest]; Maiden's Circle (13001, -1.97 DPS) [world_drop]; Glowing Crystal Ring (18402, -2.05 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 22.5 healing_power points (4.40 DPS) | yes | Maiden's Circle (13001, -1.14 DPS) [world_drop]; Glowing Crystal Ring (18402, -1.23 DPS) [quest]; Chemist's Ring (15702, -2.19 DPS, sim-verified) [quest] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (339.2 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.09 DPS) [quest]; Alchemist's Stone (13503, -1.50 DPS) [crafted] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (339.2 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.25 DPS, sim-verified) [quest]; Alchemist's Stone (13503, -1.88 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (339.2 DPS) | yes | Charstone Dirk (17710, -12.70 DPS) [dungeon]; Hand of Righteousness (7721, -18.83 DPS) [dungeon]; Hand of Edward the Odd (2243, -21.37 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 30.2 healing_power points (5.90 DPS) | yes | Captain Rackmore's Tiller (16789, -4.02 DPS) [quest]; Wand of Mana Concentration (273043, -4.03 DPS) [world]; Jaina's Firestarter (13064, -6.43 DPS, sim-verified) [world_drop] |

**New at 60:** head: Devout Crown; neck: Glowing Eye of Mordresh; shoulder: Devout Mantle; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Shard of the Splithooves; trinket2: Ankh of Life; ranged: Torch of Light

No-known-source sample (15 of 1193, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60, raid preset (troll, 025003031330000000-03505003030121431-00000000000000000)

Set DPS (verified): 657.3. Weights run: 6.2s. Verify run: 31.5s. 1193 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.114, intellect=0.778 ± 0.017, spirit=0.765 ± 0.017, mp5=1.396 ± 0.019, crit=0.276 ± 0.014 per rating point (14 rating = 1%, 3.862 per %), spell_haste=-3.191 ± 0.367

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (657.3 DPS) | yes | Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -11.14 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 19.2 healing_power points (7.88 DPS) | yes | Horizon Choker (13085, -2.16 DPS) [world_drop]; Lady Alizabeth's Pendant (13002, -3.18 DPS) [world_drop]; Necklace of Calisea (1714, -3.46 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (657.3 DPS) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp]; Mooncloth Shoulders (14139, -10.64 DPS, sim-verified) [crafted] |
| back | Wingveil Cloak (10802) | Sunken Temple: Spawn of Hakkar [dungeon] | 13.4 healing_power points (5.50 DPS) | yes | Glowing Thresher Cape (6901, -0.12 DPS) [dungeon]; Imperial Red Cloak (8248, -0.73 DPS) [world_drop]; Highborne Cloak (14450, -1.37 DPS) [world_drop] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 85.2 healing_power points (34.95 DPS) | yes | Warlord's Satin Tunic (231632, -6.38 DPS) [vendor]; Stormcloth Vest (10020, -7.52 DPS, sim-verified) [crafted]; Legionnaire's Satin Tunic (227124, -9.37 DPS) [pvp] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (657.3 DPS) | yes | Nethergeld Cuffs (254061, +0.00 DPS) [crafted]; General's Satin Bracers (17619, -1.00 DPS) [pvp]; Earthenweave Cuffs (254131, -11.71 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (657.3 DPS) | yes | Blood Guard's Satin Handwraps (227126, +0.00 DPS) [pvp]; General's Satin Handwraps (231633, +0.00 DPS) [vendor]; Gilded Gloves (254095, -11.12 DPS, sim-verified) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 48.8 healing_power points (20.00 DPS) | yes | Earthenweave Cord (254077, -5.89 DPS) [crafted]; Gilded Waistcord (254081, -5.90 DPS, sim-verified) [crafted]; Devout Belt (16696, -6.29 DPS) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 53.3 healing_power points (21.84 DPS) | yes | Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Earthenweave Leggings (254125, -16.82 DPS, sim-verified) [crafted] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 44.8 healing_power points (18.39 DPS) | yes | General's Satin Walkers (231630, -0.75 DPS) [vendor]; Gilded Sandals (254107, -2.80 DPS) [crafted]; Swarmtender's Footpads (275609, -3.22 DPS, sim-verified) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 16.2 healing_power points (6.64 DPS) | yes | Electrocutioner Lagnut (9447, -1.99 DPS) [dungeon]; Runemender's Seal (273031, -1.99 DPS) [world]; Chemist's Ring (15702, -2.19 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 13.8 healing_power points (5.68 DPS) | yes | Runemender's Seal (273031, -1.02 DPS) [world]; Chemist's Ring (15702, -1.23 DPS) [quest]; Electrocutioner Lagnut (9447, -2.33 DPS, sim-verified) [dungeon] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (657.3 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.30 DPS) [quest]; Alchemist's Stone (13503, -2.18 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (657.3 DPS) | yes | Ankh of Life (1713, -2.23 DPS, sim-verified) [world_drop]; Evonice's Landin' Pilla (18951, -4.25 DPS) [quest]; Alchemist's Stone (13503, -6.13 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (657.3 DPS) | yes | Charstone Dirk (17710, -21.37 DPS) [dungeon]; Hand of Edward the Odd (2243, -27.71 DPS, sim-verified) [world_drop]; Hand of Righteousness (7721, -32.54 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 25.4 healing_power points (10.42 DPS) | yes | Captain Rackmore's Tiller (16789, -6.69 DPS, sim-verified) [quest]; Jaina's Firestarter (13064, -7.57 DPS) [world_drop]; Earthen Rod (9381, -7.96 DPS) [dungeon] |

**New at 60:** head: Devout Crown; neck: Glowing Eye of Mordresh; shoulder: Devout Mantle; back: Wingveil Cloak; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Shard of the Splithooves; ranged: Torch of Light

No-known-source sample (15 of 1193, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

