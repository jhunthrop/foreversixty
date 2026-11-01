# Leveling BiS: Discipline

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 025004000000000000-00000000000000000-00000000000000000)

Set DPS (verified): 35.6. Weights run: 10.5s. Verify run: 28.6s. 163 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.007, intellect=1.507 ± 0.007, spirit=0.697 ± 0.006, mp5=1.593 ± 0.011, crit=0.048 ± 0.003 per rating point (14 rating = 1%, 0.679 per %), spell_haste=not significant (0.024 ± 0.025)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.37 DPS) | yes | Shadow Goggles (4373, -0.99 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.16 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 2.8 healing_power points (0.21 DPS) | yes | Scholarly Pendant (277203, -0.09 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.6 healing_power points (1.03 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.63 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.67 DPS) [dungeon] |
| back | Sanguine Cape (14376) | World drop [world_drop] | 6.0 healing_power points (0.46 DPS) | yes | Seer's Cape (6378, -0.12 DPS) [dungeon]; Forest Cloak (4710, -0.23 DPS) [world_drop]; Pearl-clasped Cloak (5542, -0.31 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 22.6 healing_power points (1.72 DPS) | yes | Robe of the Moccasin (6465, -0.51 DPS) [dungeon]; Bloody Apron (6226, -0.73 DPS) [dungeon]; Corsair's Overshirt (5202, -0.89 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 12.3 healing_power points (0.94 DPS) | yes | Repurposed Hair Band (281256, -0.60 DPS) [quest]; Mystic's Bracelets (14366, -0.71 DPS) [world_drop]; Bright Bracers (3647, -0.99 DPS, sim-verified) [world_drop] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 15.5 healing_power points (1.18 DPS) | yes | Tomb Robber's Gloves (280096, -0.49 DPS) [quest]; Bright Gloves (3066, -0.51 DPS) [world_drop]; Magefist Gloves (12977, -0.76 DPS, sim-verified) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 17.0 healing_power points (1.30 DPS) | yes | Keller's Girdle (2911, -0.38 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.63 DPS, sim-verified) [crafted]; Wise Man's Belt (4786, -0.72 DPS) [vendor] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 29.8 healing_power points (2.27 DPS) | yes | Filigreed Silky Leggings (253939, -1.37 DPS) [crafted]; Filigreed Shadow Leggings (253943, -1.37 DPS) [crafted]; Darkweave Breeches (12987, -1.72 DPS, sim-verified) [world_drop] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 13.5 healing_power points (1.03 DPS) | yes | Spidersilk Boots (4320, -0.57 DPS) [crafted]; Sanguine Sandals (14374, -0.57 DPS) [world_drop]; Kimbra Boots (6191, -0.64 DPS, sim-verified) [quest] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 9.0 healing_power points (0.69 DPS) | yes | Band of Purification (12996, -0.37 DPS) [world_drop]; Deep Fathom Ring (6463, -0.42 DPS) [dungeon]; Minor Channeling Ring (1449, -0.46 DPS) [quest] |
| finger2 | Volcanic Rock Ring (12053) | World drop [world_drop] | 4.5 healing_power points (0.34 DPS) | yes | Deep Fathom Ring (6463, -0.08 DPS) [dungeon]; Minor Channeling Ring (1449, -0.11 DPS) [quest]; Band of Purification (12996, -0.29 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 15.1 healing_power points (1.15 DPS) | yes | Staff of Westfall (2042, -0.32 DPS, sim-verified) [quest]; Lesser Staff of the Spire (1300, -0.46 DPS) [world]; Witching Stave (1484, -0.57 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Flaring Baton (5326) (or Moonstone Wand (15204)) | The Escape [quest] | 3.0 healing_power points (0.23 DPS) | yes | Moonstone Wand (15204, +0.00 DPS) [quest]; Sable Wand (7607, -0.07 DPS) [quest]; Elven Wand (5604, -0.11 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Tarnished Locket; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Volcanic Rock Ring; main_hand: Twisted Chanter's Staff; ranged: Flaring Baton

No-known-source sample (15 of 163, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes; 16607 Acolyte's Sacrificial Robes; 18851 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (human, 025005031320000000-00000000000000000-00000000000000000)

Set DPS (verified): 65.5. Weights run: 11.0s. Verify run: 32.7s. 294 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.006, intellect=1.697 ± 0.010, spirit=1.138 ± 0.011, mp5=2.521 ± 0.009, crit=0.087 ± 0.005 per rating point (14 rating = 1%, 1.215 per %), spell_haste=not significant (0.045 ± 0.027)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 39.8 healing_power points (3.55 DPS) | yes | Resilient Cap (14401, -1.48 DPS) [world_drop]; Filigreed Pristine Circlet (253975, -1.59 DPS) [crafted]; Nightsky Cowl (4039, -2.47 DPS, sim-verified) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 11.3 healing_power points (1.01 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.10 DPS) [world]; Crystal Starfire Medallion (5003, -0.32 DPS, sim-verified) [world_drop]; Pendant of Myzrael (4614, -0.40 DPS) [dungeon] |
| shoulder | Mantle of Honor (3560) | Bride of the Embalmer [quest] | 19.8 healing_power points (1.77 DPS) | yes | Death Speaker Mantle (6685, -0.10 DPS) [dungeon]; Nightsky Mantle (4718, -0.22 DPS, sim-verified) [world_drop]; Faerie Mantle (5820, -0.25 DPS) [quest] |
| back | Prelacy Cape (7004) | Researching the Corruption [quest] | 17.8 healing_power points (1.59 DPS) | yes | Repairman's Cape (9605, -0.27 DPS) [quest]; Glowing Thresher Cape (6901, -0.30 DPS, sim-verified) [dungeon]; Sergeant Major's Cape (16315, -0.58 DPS) [pvp] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 38.7 healing_power points (3.44 DPS) | yes | Pristine Gown (253961, -0.37 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -0.82 DPS) [dungeon]; Filigreed Pristine Gown (253901, -1.23 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 16.0 healing_power points (1.43 DPS) | yes | Glowing Magical Bracelets (13106, -0.22 DPS) [world_drop]; Stonecloth Bindings (14416, -0.67 DPS) [world_drop]; Nightsky Wristbands (6407, -0.82 DPS, sim-verified) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 22.1 healing_power points (1.97 DPS) | yes | Truefaith Gloves (7049, +0.00 DPS, sim-verified) [crafted]; Hotshot Pilot's Gloves (9491, -0.26 DPS) [dungeon]; Pristine Gloves (253913, -0.54 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 17.8 healing_power points (1.58 DPS) | yes | Novice Ardent's Sash (253887, -0.29 DPS) [crafted]; Dreamer's Belt (4829, -0.32 DPS) [vendor]; Resilient Cord (14406, -0.70 DPS, sim-verified) [world_drop] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 39.6 healing_power points (3.52 DPS) | yes | Filigreed Pristine Leggings (253937, -0.90 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -1.07 DPS) [crafted]; Darkweave Breeches (12987, -1.86 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 36.4 healing_power points (3.24 DPS) | yes | Nimbus Boots (6998, -1.73 DPS) [quest]; Frothing Slippers (254003, -1.78 DPS) [crafted]; Acidic Walkers (9454, -2.43 DPS, sim-verified) [dungeon] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 14.0 healing_power points (1.24 DPS) | yes | The Queen's Jewel (13094, -0.13 DPS) [world_drop]; Defias Mage Ring (1077, -0.34 DPS) [world]; Seal of Wrynn (2933, -0.34 DPS) [quest] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 14.0 healing_power points (1.24 DPS) | yes | Defias Mage Ring (1077, -0.34 DPS) [world]; Seal of Wrynn (2933, -0.34 DPS) [quest]; The Queen's Jewel (13094, -0.60 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-verified (65.5 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -0.51 DPS) [quest]; Death Speaker Scepter (2816, -1.80 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 11.8 healing_power points (1.05 DPS) | yes | Lesser Mystic Wand (11289, -0.45 DPS) [crafted]; Starfaller (13063, -0.45 DPS) [world_drop]; Summoner's Wand (5245, -1.07 DPS, sim-verified) [world] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Mantle of Honor; back: Prelacy Cape; chest: Death Speaker Robes; hands: Gloves of Old; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Spine of the Devourer; ranged: Wand of Mana Concentration

No-known-source sample (15 of 294, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 40 (human, 025005031330011520-00000000000000000-00000000000000000)

Set DPS (verified): 127.9. Weights run: 13.0s. Verify run: 76.9s. 396 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.028, intellect=1.781 ± 0.022, spirit=0.280 ± 0.047, mp5=3.107 ± 0.033, crit=0.210 ± 0.008 per rating point (14 rating = 1%, 2.935 per %), spell_haste=-4.768 ± 0.266

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 46.8 healing_power points (4.60 DPS) | yes | Holy Shroud (2721, -1.24 DPS, sim-verified) [world_drop]; Whitemane's Chapeau (7720, -1.24 DPS) [dungeon]; Thinking Cap (2624, -1.62 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 27.2 healing_power points (2.68 DPS) | yes | Necklace of Calisea (1714, -0.52 DPS, sim-verified) [dungeon]; Archmagister's Faceted Pendant (273049, -1.63 DPS) [world]; Kaleidoscope Chain (13084, -1.87 DPS) [world_drop] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 38.2 healing_power points (3.75 DPS) | yes | Mistscape Mantle (4734, -1.69 DPS) [dungeon]; Silksand Shoulder Pads (14423, -1.74 DPS) [world_drop]; Windchaser Amice (14432, -2.02 DPS, sim-verified) [world_drop] |
| back | Blackforge Cape (6424) | Uldaman: Ancient Treasure [dungeon] | 15.4 healing_power points (1.51 DPS) | yes | Prelacy Cape (7004, +0.00 DPS, sim-verified) [quest]; Glyphed Cloak (4732, -0.28 DPS) [world_drop]; Sergeant Major's Cape (16336, -0.29 DPS) [pvp] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 60.2 healing_power points (5.91 DPS) | yes | Pristine Gown (253961, -2.62 DPS) [crafted]; Red Mageweave Vest (10007, -2.77 DPS) [crafted]; Death Speaker Robes (6682, -8.08 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 18.2 healing_power points (1.79 DPS) | yes | Mistscape Bracers (4045, -0.39 DPS) [dungeon]; Enchanted Stonecloth Bracers (4979, -0.39 DPS) [quest]; Windchaser Cuffs (14429, -0.48 DPS, sim-verified) [world_drop] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (127.9 DPS) | yes | Truefaith Gloves (7049, -1.18 DPS) [crafted]; Earthen Silk Gloves (254017, -1.19 DPS) [crafted]; Gilded Handwraps (254021, -1.75 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 37.9 healing_power points (3.72 DPS) | yes | Razzeric's Customized Seatbelt (6726, -1.63 DPS) [quest]; Teacher's Sash (10747, -1.63 DPS) [quest]; Deathmage Sash (10771, -3.42 DPS, sim-verified) [dungeon] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-verified (127.9 DPS) | yes | Stoneweaver Leggings (9407, +0.00 DPS) [dungeon]; Filigreed Pristine Leggings (253937, +0.00 DPS) [crafted]; Pristine Leggings (253987, -2.27 DPS, sim-verified) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 33.6 healing_power points (3.30 DPS) | yes | Furen's Boots (13100, -1.63 DPS) [world_drop]; Kodo Rustler Boots (15697, -1.76 DPS) [quest]; Nimbus Boots (6998, -2.21 DPS, sim-verified) [quest] |
| finger1 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 13.3 healing_power points (1.31 DPS) | yes | Mindbender Loop (5009, -0.03 DPS) [world_drop]; Welken Ring (5011, -0.24 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.26 DPS) [dungeon] |
| finger2 | Ogremind Ring (1993) | World drop [world_drop] | 13.3 healing_power points (1.31 DPS) | yes | Mindbender Loop (5009, -0.17 DPS, sim-verified) [world_drop]; Welken Ring (5011, -0.24 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.26 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (127.9 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Death Speaker Scepter (2816, -2.77 DPS) [dungeon]; Hypnotic Blade (7714, -3.65 DPS) [dungeon] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 29.8 healing_power points (2.93 DPS) | yes | Beacon of Hope (9393, -0.89 DPS, sim-verified) [dungeon]; Orb of Souls (249395, -1.49 DPS) [crafted]; Aurora Sphere (7610, -1.59 DPS) [dungeon] |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 13.3 healing_power points (1.31 DPS) | yes | Jaina's Firestarter (13064, +0.00 DPS, sim-verified) [world_drop]; Flash Wand (5248, -0.53 DPS) [quest]; Goblin Igniter (5253, -0.53 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Blackforge Cape; chest: Stormcloth Vest; hands: Stormcloth Gloves; waist: Gilded Cord; legs: Stormcloth Pants; finger1: Voodoo Band; finger2: Ogremind Ring; trinket1: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Orb of Lorica

No-known-source sample (15 of 396, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 50 (human, 025005031330011530-03501000000000000-00000000000000000)

Set DPS (verified): 209.1. Weights run: 13.0s. Verify run: 41.6s. 515 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.010, intellect=2.283 ± 0.022, spirit=1.260 ± 0.043, mp5=3.742 ± 0.054, crit=0.492 ± 0.016 per rating point (14 rating = 1%, 6.895 per %), spell_haste=-5.945 ± 0.376

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 70.5 healing_power points (5.95 DPS) | yes | Knight-Lieutenant's Satin Cover (220896, +0.00 DPS) [vendor]; Stormcloth Headband (10032, -0.93 DPS) [crafted]; Papal Fez (9431, -2.16 DPS, sim-verified) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 37.0 healing_power points (3.12 DPS) | yes | Glowing Eye of Mordresh (10769, -1.01 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -1.03 DPS) [dungeon]; Mindburst Medallion (11196, -1.44 DPS) [quest] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | 51.8 healing_power points (4.38 DPS) | yes | Knight-Lieutenant's Satin Pads (220894, +0.00 DPS) [vendor]; Rotgrip Mantle (17732, -0.91 DPS) [dungeon]; Inquisitor's Shawl (19507, -1.80 DPS, sim-verified) [dungeon] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 30.2 healing_power points (2.55 DPS) | yes | Blackforge Cape (6424, -0.58 DPS) [dungeon]; Keeper's Cloak (14665, -0.62 DPS) [world_drop]; Highborne Cloak (14450, -1.16 DPS, sim-verified) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 68.1 healing_power points (5.75 DPS) | yes | Ghostweave Vest (14141, -1.09 DPS) [crafted]; Earth Warder's Vest (21311, -1.35 DPS) [quest]; Knight's Satin Armor (220892, -1.75 DPS, sim-verified) [vendor] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 42.3 healing_power points (3.57 DPS) | yes | Forgotten Wraps (9433, -1.26 DPS) [world_drop]; Imperial Red Bracers (8247, -1.45 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -2.06 DPS, sim-verified) [quest] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 61.9 healing_power points (5.23 DPS) | yes | Earthenweave Gloves (254075, -1.86 DPS) [crafted]; Stormcloth Gloves (10011, -1.99 DPS) [crafted]; Sergeant Major's Satin Gloves (220897, -3.75 DPS, sim-verified) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 55.4 healing_power points (4.68 DPS) | yes | Gilded Waistcord (254081, -0.49 DPS) [crafted]; Gilded Cord (254037, -0.64 DPS) [crafted]; Earthenweave Cord (254077, -1.31 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 69.8 healing_power points (5.89 DPS) | yes | Senior Designer's Pantaloons (11841, -0.52 DPS) [dungeon]; Knight's Satin Leggings (220893, -0.89 DPS, sim-verified) [vendor]; Dalewind Trousers (13008, -1.50 DPS) [world_drop] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 51.5 healing_power points (4.35 DPS) | yes | Vinerot Sandals (17748, -0.54 DPS) [dungeon]; Sergeant Major's Satin Boots (220895, -0.72 DPS, sim-verified) [vendor]; Gilded Slippers (254001, -0.89 DPS) [crafted] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 28.8 healing_power points (2.43 DPS) | yes | Ogremind Ring (1993, -0.76 DPS) [world_drop]; Voodoo Band (1996, -0.76 DPS) [world]; Mindbender Loop (5009, -0.87 DPS) [world_drop] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | 20.2 healing_power points (1.71 DPS) | yes | Ogremind Ring (1993, -0.04 DPS) [world_drop]; Mindbender Loop (5009, -0.15 DPS) [world_drop]; Voodoo Band (1996, -0.38 DPS, sim-verified) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (209.1 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -0.43 DPS) [quest]; Alchemist's Stone (13503, -0.85 DPS) [crafted] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-verified (209.1 DPS) | yes | Thunderbrew's Boot Flask (744, -0.21 DPS) [quest]; Uther's Strength (11302, -0.40 DPS, sim-verified) [world_drop]; Alchemist's Stone (13503, -0.64 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (209.1 DPS) | yes | Charstone Dirk (17710, -6.67 DPS) [dungeon]; Spellshifter Rod (9527, -10.22 DPS) [quest]; Hammer of the Northern Wind (810, -30.86 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cairnstone Sliver (9654) | The Morrow Stone [quest] | 17.7 healing_power points (1.50 DPS) | yes | Wand of Mana Concentration (273043, -0.09 DPS) [world]; Jaina's Firestarter (13064, -0.28 DPS, sim-verified) [world_drop]; Goblin Igniter (5253, -0.41 DPS) [quest] |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Nethergeld Shoulders; back: Imperial Red Cloak; wrist: Nethergeld Cuffs; hands: Gilded Gloves; waist: Dawnspire Cord; legs: Kilt of the Atal'ai Prophet; feet: Gilded Sandals; finger1: Eye of Adaegus; finger2: Welken Ring; trinket2: Evonice's Landin' Pilla; main_hand: Glowing Brightwood Staff; ranged: Cairnstone Sliver

No-known-source sample (15 of 515, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60 (human, 025005031330011530-03505003030000000-00000000000000000)

Set DPS (verified): 301.9. Weights run: 12.1s. Verify run: 46.6s. 1200 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.063, intellect=1.486 ± 0.027, spirit=4.369 ± 0.040, mp5=1.706 ± 0.039, crit=0.322 ± 0.015 per rating point (14 rating = 1%, 4.511 per %), spell_haste=-9.472 ± 0.574

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | 121.2 healing_power points (17.95 DPS) | yes | Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Earthenweave Crown (254137, -5.69 DPS, sim-verified) [crafted] |
| neck | Lady Alizabeth's Pendant (13002) | World drop [world_drop] | 65.5 healing_power points (9.71 DPS) | yes | Necklace of Calisea (1714, -3.64 DPS) [dungeon]; Explorers' League Commendation (7746, -3.88 DPS) [quest]; Amberglow Talisman (10824, -6.04 DPS, sim-verified) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (+4.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Earthenweave Mantle (254119, -0.46 DPS) [crafted]; Mooncloth Shoulders (14139, -4.46 DPS, sim-verified) [crafted] |
| back | Ethereal Mist Cape (11873) | Ancient Spirit [quest] | 56.9 healing_power points (8.43 DPS) | yes | Wingveil Cloak (10802, -0.57 DPS) [dungeon]; Arcane Cloak (8286, -0.66 DPS) [world_drop]; Indomitable Cloak (14683, -2.43 DPS, sim-verified) [world_drop] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 142.9 healing_power points (21.17 DPS) | yes | Knight-Captain's Satin Tunic (227122, -2.88 DPS) [pvp]; Field Marshal's Satin Tunic (231624, -3.11 DPS) [vendor]; Mooncloth Vest (14138, -8.51 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | 66.6 healing_power points (9.86 DPS) | yes | Marshal's Satin Bracers (17606, +0.00 DPS) [pvp]; Earthenweave Cuffs (254131, -0.13 DPS) [crafted]; Nethergeld Cuffs (254061, -2.12 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | 104.1 healing_power points (15.42 DPS) | yes | Magister's Gloves (16684, -3.28 DPS) [dungeon]; Gilded Gloves (254095, -3.75 DPS) [crafted]; Swarmtender's Gloves (275607, -5.90 DPS, sim-verified) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Satin Sash (17609, +0.00 DPS) [pvp]; Mana-infused Cord (279267, -0.21 DPS) [crafted]; Wisdom of the Timbermaw (19047, -3.35 DPS, sim-verified) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 146.8 healing_power points (21.74 DPS) | yes | Knight-Captain's Satin Legguards (227125, -3.45 DPS) [pvp]; Dalewind Trousers (13008, -3.58 DPS) [world_drop]; Mooncloth Leggings (14137, -4.56 DPS, sim-verified) [crafted] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 101.6 healing_power points (15.05 DPS) | yes | Devout Sandals (16691, +0.00 DPS, sim-verified) [dungeon]; Elunarian Boots (14458, -2.73 DPS) [world_drop]; Furen's Boots (13100, -3.15 DPS) [world_drop] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 62.9 healing_power points (9.32 DPS) | yes | Choking Band (11868, -0.91 DPS) [quest]; Coldwater Ring (4550, -2.84 DPS) [quest]; Seedtime Hoop (9655, -2.85 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 61.3 healing_power points (9.09 DPS) | yes | Coldwater Ring (4550, -2.60 DPS) [quest]; Seedtime Hoop (9655, -2.61 DPS) [quest]; Choking Band (11868, -3.20 DPS, sim-verified) [quest] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Mark of the Chosen (17774, -3.74 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -14.32 DPS) [dungeon]; Staff of Hale Magefire (13000, -14.69 DPS) [world_drop]; Hand of Edward the Odd (2243, -18.63 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 65.1 healing_power points (9.64 DPS) | yes | Jaina's Firestarter (13064, -6.37 DPS) [world_drop]; Captain Rackmore's Tiller (16789, -6.81 DPS) [quest]; Cairnstone Sliver (9654, -8.68 DPS, sim-verified) [quest] |

**New at 60:** head: Devout Crown; neck: Lady Alizabeth's Pendant; shoulder: Devout Mantle; back: Ethereal Mist Cape; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Shard of the Splithooves; trinket2: Uther's Strength; ranged: Torch of Light

No-known-source sample (15 of 1200, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60, raid preset (human, 025005031330011530-03505003030000000-00000000000000000)

Set DPS (verified): 569.1. Weights run: 7.9s. Verify run: 30.6s. 1200 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.094, intellect=1.025 ± 0.024, spirit=0.993 ± 0.023, mp5=2.229 ± 0.026, crit=0.473 ± 0.019 per rating point (14 rating = 1%, 6.616 per %), spell_haste=-5.802 ± 0.482

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (569.0 DPS) | yes | Mooncloth Circlet (14140, +0.00 DPS) [crafted]; Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 21.2 healing_power points (5.74 DPS) | yes | Horizon Choker (13085, -0.78 DPS) [world_drop]; Lady Alizabeth's Pendant (13002, -1.71 DPS) [world_drop]; Necklace of Calisea (1714, -1.91 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (569.0 DPS) | yes | Mooncloth Shoulders (14139, +0.00 DPS) [crafted]; Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp] |
| back | Sergeant Major's Cape (16337) | PvP rank 9 · Sergeant Major · Alliance [pvp] | sim-verified (569.0 DPS) | yes | Prelacy Cape (7004, +0.00 DPS) [quest]; Wingveil Cloak (10802, -0.06 DPS) [dungeon]; Imperial Red Cloak (8248, -0.24 DPS) [world_drop] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 88.9 healing_power points (24.05 DPS) | yes | Field Marshal's Satin Tunic (231624, -2.77 DPS) [vendor]; Knight-Captain's Satin Tunic (227122, -5.31 DPS) [pvp]; Stormcloth Vest (10020, -6.21 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (569.0 DPS) | yes | Marshal's Satin Bracers (17606, +0.00 DPS) [pvp]; Nethergeld Cuffs (254061, +0.00 DPS) [crafted]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (569.0 DPS) | yes | Knight-Lieutenant's Satin Handwraps (227128, +0.00 DPS) [pvp]; Marshal's Satin Handwraps (231623, +0.00 DPS) [vendor]; Gilded Gloves (254095, +0.00 DPS) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 54.3 healing_power points (14.69 DPS) | yes | Devout Belt (16696, -3.75 DPS) [dungeon]; Marshal's Satin Sash (17609, -4.28 DPS) [pvp]; Gilded Waistcord (254081, -4.35 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 62.2 healing_power points (16.84 DPS) | yes | Knight-Captain's Satin Legguards (227125, +0.00 DPS) [pvp]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Earthenweave Leggings (254125, -2.52 DPS) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (569.0 DPS) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; Marshal's Satin Walkers (231627, +0.00 DPS) [vendor]; Swarmtender's Footpads (275609, +0.00 DPS) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 21.2 healing_power points (5.73 DPS) | yes | Maiden's Circle (13001, -2.18 DPS) [world_drop]; Electrocutioner Lagnut (9447, -2.23 DPS) [dungeon]; Runemender's Seal (273031, -2.23 DPS) [world] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 18.1 healing_power points (4.89 DPS) | yes | Maiden's Circle (13001, -1.33 DPS) [world_drop]; Electrocutioner Lagnut (9447, -1.38 DPS) [dungeon]; Runemender's Seal (273031, -1.38 DPS) [world] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (569.0 DPS) | yes | Shard of the Splithooves (10659, -1.25 DPS) [quest]; Evonice's Landin' Pilla (18951, -2.18 DPS) [quest]; Thunderbrew's Boot Flask (744, -2.72 DPS) [quest] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (569.0 DPS) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Evonice's Landin' Pilla (18951, -0.54 DPS) [quest]; Thunderbrew's Boot Flask (744, -1.07 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -15.41 DPS) [dungeon]; Hand of Edward the Odd (2243, -20.41 DPS) [world_drop]; Hand of Righteousness (7721, -23.47 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 27.9 healing_power points (7.56 DPS) | yes | Cairnstone Sliver (9654, -4.83 DPS) [quest]; Jaina's Firestarter (13064, -5.09 DPS) [world_drop]; Captain Rackmore's Tiller (16789, -5.13 DPS) [quest] |

**New at 60:** head: Devout Crown; neck: Glowing Eye of Mordresh; shoulder: Devout Mantle; back: Sergeant Major's Cape; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Uther's Strength; trinket2: Ankh of Life; ranged: Torch of Light

No-known-source sample (15 of 1200, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

## Horde

### Band 20 (undead, 025004000000000000-00000000000000000-00000000000000000)

Set DPS (verified): 38.1. Weights run: 10.5s. Verify run: 29.7s. 156 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.007, intellect=1.507 ± 0.007, spirit=0.697 ± 0.006, mp5=1.593 ± 0.011, crit=0.048 ± 0.003 per rating point (14 rating = 1%, 0.679 per %), spell_haste=not significant (0.024 ± 0.025)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.37 DPS) | yes | Shadow Goggles (4373, -1.07 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.16 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 2.1 healing_power points (0.16 DPS) | yes | Scholarly Pendant (277203, -0.36 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.6 healing_power points (1.03 DPS) | yes | Slime-encrusted Pads (6461, -0.67 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -0.88 DPS, sim-verified) [crafted] |
| back | Sanguine Cape (14376) | World drop [world_drop] | 6.0 healing_power points (0.46 DPS) | yes | Seer's Cape (6378, -0.12 DPS) [dungeon]; Forest Cloak (4710, -0.23 DPS) [world_drop]; Pearl-clasped Cloak (5542, -0.30 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 22.6 healing_power points (1.72 DPS) | yes | Robe of the Moccasin (6465, -0.51 DPS) [dungeon]; Bloody Apron (6226, -0.73 DPS) [dungeon]; Corsair's Overshirt (5202, -0.90 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 12.3 healing_power points (0.94 DPS) | yes | Bright Bracers (3647, -0.48 DPS) [world_drop]; Crystalline Cuffs (14148, -0.55 DPS) [dungeon]; Featherbead Bracers (15452, -0.86 DPS, sim-verified) [quest] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 15.5 healing_power points (1.18 DPS) | yes | Magefist Gloves (12977, -0.40 DPS) [world_drop]; Bright Gloves (3066, -0.51 DPS) [world_drop]; Blight Gloves (279877, -0.59 DPS, sim-verified) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 17.0 healing_power points (1.30 DPS) | yes | Keller's Girdle (2911, -0.38 DPS) [world_drop]; Wise Man's Belt (4786, -0.72 DPS) [vendor]; Novice Ardent's Sash (253887, -0.77 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 29.8 healing_power points (2.27 DPS) | yes | Filigreed Silky Leggings (253939, -1.37 DPS) [crafted]; Filigreed Shadow Leggings (253943, -1.37 DPS) [crafted]; Darkweave Breeches (12987, -2.00 DPS, sim-verified) [world_drop] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 13.5 healing_power points (1.03 DPS) | yes | Spidersilk Boots (4320, -0.57 DPS) [crafted]; Bluegill Sandals (1560, -0.59 DPS) [world]; Sanguine Sandals (14374, -1.20 DPS, sim-verified) [world_drop] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 9.0 healing_power points (0.69 DPS) | yes | Volcanic Rock Ring (12053, -0.34 DPS) [world_drop]; Band of Purification (12996, -0.37 DPS) [world_drop]; Deep Fathom Ring (6463, -0.42 DPS) [dungeon] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 7.5 healing_power points (0.57 DPS) | yes | Band of Purification (12996, -0.26 DPS) [world_drop]; Deep Fathom Ring (6463, -0.31 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.44 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 15.1 healing_power points (1.15 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Staff of Orgrimmar (15444, -0.18 DPS) [quest]; Lesser Staff of the Spire (1300, -0.46 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Decay (5252) (or Flaring Baton (5326)) | Beren's Peril [quest] | 3.0 healing_power points (0.23 DPS) | yes | Flaring Baton (5326, +0.00 DPS) [quest]; Wisesight Wand (286750, -0.11 DPS) [world] |

**New at 20:** head: Pristine Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Gnarled Necromancer's Staff; ranged: Wand of Decay

No-known-source sample (15 of 156, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (undead, 025005031320000000-00000000000000000-00000000000000000)

Set DPS (verified): 69.9. Weights run: 11.0s. Verify run: 37.2s. 286 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.006, intellect=1.697 ± 0.010, spirit=1.138 ± 0.011, mp5=2.521 ± 0.009, crit=0.087 ± 0.005 per rating point (14 rating = 1%, 1.215 per %), spell_haste=not significant (0.045 ± 0.027)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 39.8 healing_power points (3.55 DPS) | yes | Nightsky Cowl (4039, -1.23 DPS) [world_drop]; Resilient Cap (14401, -1.48 DPS) [world_drop]; Filigreed Pristine Circlet (253975, -1.59 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 11.3 healing_power points (1.01 DPS) | yes | Crystal Starfire Medallion (5003, -0.10 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.10 DPS) [world]; Pendant of Myzrael (4614, -0.40 DPS) [dungeon] |
| shoulder | Ghostly Mantle (3324) | Deathstalkers in Shadowfang [quest] | sim-verified (69.9 DPS) | yes | Nightsky Mantle (4718, +0.00 DPS) [world_drop]; Death Speaker Mantle (6685, -0.05 DPS) [dungeon]; Violet Sorcerer's Mantle (273051, -0.35 DPS) [world] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 16.1 healing_power points (1.43 DPS) | yes | Resilient Cape (14400, -0.53 DPS) [world_drop]; Silver-thread Cloak (4713, -0.63 DPS) [world_drop]; Soft Willow Cape (16661, -0.68 DPS) [quest] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 38.7 healing_power points (3.44 DPS) | yes | Pristine Gown (253961, -0.20 DPS) [crafted]; Beguiler Robes (7728, -0.82 DPS) [dungeon]; Filigreed Pristine Gown (253901, -1.23 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 16.0 healing_power points (1.43 DPS) | yes | Nightsky Wristbands (6407, -0.22 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.22 DPS) [world_drop]; Featherbead Bracers (15452, -0.67 DPS) [quest] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 22.1 healing_power points (1.97 DPS) | yes | Truefaith Gloves (7049, -0.18 DPS) [crafted]; Hotshot Pilot's Gloves (9491, -0.26 DPS) [dungeon]; Blight Gloves (279877, -0.41 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (69.9 DPS) | yes | Lilac Sash (6780, +0.00 DPS) [quest]; Resilient Cord (14406, -0.27 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.29 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 39.6 healing_power points (3.52 DPS) | yes | Filigreed Pristine Leggings (253937, -0.61 DPS) [crafted]; Earthen Leggings (253999, -1.07 DPS) [crafted]; Sacred Burial Trousers (6282, -1.86 DPS) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 36.4 healing_power points (3.24 DPS) | yes | Acidic Walkers (9454, -1.63 DPS) [dungeon]; Frothing Slippers (254003, -1.78 DPS) [crafted]; Black Slippers (254007, -1.78 DPS) [crafted] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 14.0 healing_power points (1.24 DPS) | yes | Solid Gold Band (279845, -0.11 DPS) [quest]; The Queen's Jewel (13094, -0.13 DPS) [world_drop]; Defias Mage Ring (1077, -0.34 DPS) [world] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 14.0 healing_power points (1.24 DPS) | yes | Solid Gold Band (279845, -0.11 DPS) [quest]; The Queen's Jewel (13094, -0.13 DPS) [world_drop]; Defias Mage Ring (1077, -0.34 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, +0.00 DPS) [dungeon]; Tendonscraper Dagger (273053, +0.00 DPS) [world]; Staff of the Friar (3415, -1.54 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 11.8 healing_power points (1.05 DPS) | yes | Summoner's Wand (5245, -0.45 DPS) [world]; Lesser Mystic Wand (11289, -0.45 DPS) [crafted]; Starfaller (13063, -0.45 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Ghostly Mantle; back: Glowing Thresher Cape; chest: Death Speaker Robes; hands: Gloves of Old; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Spine of the Devourer; ranged: Wand of Mana Concentration

No-known-source sample (15 of 286, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (undead, 025005031330011520-00000000000000000-00000000000000000)

Set DPS (verified): 133.5. Weights run: 13.0s. Verify run: 78.0s. 388 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.028, intellect=1.781 ± 0.022, spirit=0.280 ± 0.047, mp5=3.107 ± 0.033, crit=0.210 ± 0.008 per rating point (14 rating = 1%, 2.935 per %), spell_haste=-4.768 ± 0.266

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 46.8 healing_power points (4.60 DPS) | yes | Whitemane's Chapeau (7720, -1.24 DPS) [dungeon]; Holy Shroud (2721, -1.49 DPS, sim-verified) [world_drop]; Thinking Cap (2624, -1.62 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 27.2 healing_power points (2.68 DPS) | yes | Necklace of Calisea (1714, -0.64 DPS, sim-verified) [dungeon]; Archmagister's Faceted Pendant (273049, -1.63 DPS) [world]; Kaleidoscope Chain (13084, -1.87 DPS) [world_drop] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 38.2 healing_power points (3.75 DPS) | yes | Mistscape Mantle (4734, -1.69 DPS) [dungeon]; Silksand Shoulder Pads (14423, -1.74 DPS) [world_drop]; Windchaser Amice (14432, -2.11 DPS, sim-verified) [world_drop] |
| back | Blackforge Cape (6424) | Uldaman: Ancient Treasure [dungeon] | 15.4 healing_power points (1.51 DPS) | yes | Windchaser Cloak (14430, -0.35 DPS) [world_drop]; Resilient Cape (14400, -0.46 DPS) [world_drop]; Glyphed Cloak (4732, -1.07 DPS, sim-verified) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 60.2 healing_power points (5.91 DPS) | yes | Pristine Gown (253961, -2.62 DPS) [crafted]; Red Mageweave Vest (10007, -2.77 DPS) [crafted]; Death Speaker Robes (6682, -7.89 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 18.2 healing_power points (1.79 DPS) | yes | Mistscape Bracers (4045, -0.39 DPS) [dungeon]; Enchanted Stonecloth Bracers (4979, -0.39 DPS) [quest]; Windchaser Cuffs (14429, -0.43 DPS, sim-verified) [world_drop] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (133.5 DPS) | yes | Truefaith Gloves (7049, -1.18 DPS) [crafted]; Earthen Silk Gloves (254017, -1.19 DPS) [crafted]; Gilded Handwraps (254021, -1.48 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 37.9 healing_power points (3.72 DPS) | yes | Razzeric's Customized Seatbelt (6726, -1.63 DPS) [quest]; Mistscape Sash (4736, -1.80 DPS) [dungeon]; Deathmage Sash (10771, -3.07 DPS, sim-verified) [dungeon] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-verified (133.5 DPS) | yes | Stoneweaver Leggings (9407, +0.00 DPS) [dungeon]; Filigreed Pristine Leggings (253937, +0.00 DPS) [crafted]; Pristine Leggings (253987, -2.40 DPS, sim-verified) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 33.6 healing_power points (3.30 DPS) | yes | Furen's Boots (13100, -1.63 DPS) [world_drop]; Kodo Rustler Boots (15697, -1.76 DPS) [quest]; Boots of the Maharishi (9658, -1.81 DPS, sim-verified) [quest] |
| finger1 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 13.3 healing_power points (1.31 DPS) | yes | Mindbender Loop (5009, -0.03 DPS) [world_drop]; Welken Ring (5011, -0.24 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.26 DPS) [dungeon] |
| finger2 | Ogremind Ring (1993) | World drop [world_drop] | 13.3 healing_power points (1.31 DPS) | yes | Welken Ring (5011, -0.24 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.26 DPS) [dungeon]; Mindbender Loop (5009, -0.35 DPS, sim-verified) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (133.5 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Death Speaker Scepter (2816, -2.77 DPS) [dungeon]; Hypnotic Blade (7714, -3.65 DPS) [dungeon] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | 15.1 healing_power points (1.48 DPS) | yes | Orb of Souls (249395, +0.00 DPS, sim-verified) [crafted]; Witch's Finger (16887, -0.09 DPS) [quest]; Aurora Sphere (7610, -0.15 DPS) [dungeon] |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 13.3 healing_power points (1.31 DPS) | yes | Jaina's Firestarter (13064, +0.00 DPS, sim-verified) [world_drop]; Flash Wand (5248, -0.53 DPS) [quest]; Goblin Igniter (5253, -0.53 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Blackforge Cape; chest: Stormcloth Vest; hands: Stormcloth Gloves; waist: Gilded Cord; legs: Stormcloth Pants; finger1: Voodoo Band; finger2: Ogremind Ring; trinket1: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Beacon of Hope

No-known-source sample (15 of 388, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16336 Sergeant Major's Cape

### Band 50 (undead, 025005031330011530-03501000000000000-00000000000000000)

Set DPS (verified): 220.8. Weights run: 13.0s. Verify run: 42.9s. 507 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.010, intellect=2.283 ± 0.022, spirit=1.260 ± 0.043, mp5=3.742 ± 0.054, crit=0.492 ± 0.016 per rating point (14 rating = 1%, 6.895 per %), spell_haste=-5.945 ± 0.376

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 70.5 healing_power points (5.95 DPS) | yes | Blood Guard's Satin Cover (220899, +0.00 DPS) [vendor]; Stormcloth Headband (10032, -0.93 DPS) [crafted]; Papal Fez (9431, -2.24 DPS, sim-verified) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 37.0 healing_power points (3.12 DPS) | yes | Glowing Eye of Mordresh (10769, -0.59 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -1.03 DPS) [dungeon]; Choker of the High Shaman (4112, -1.31 DPS) [quest] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | 51.8 healing_power points (4.38 DPS) | yes | Blood Guard's Satin Pads (220901, +0.00 DPS) [vendor]; Rotgrip Mantle (17732, -0.91 DPS) [dungeon]; Inquisitor's Shawl (19507, -2.36 DPS, sim-verified) [dungeon] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 30.2 healing_power points (2.55 DPS) | yes | Highborne Cloak (14450, -0.57 DPS, sim-verified) [world_drop]; Blackforge Cape (6424, -0.58 DPS) [dungeon]; Keeper's Cloak (14665, -0.62 DPS) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 68.1 healing_power points (5.75 DPS) | yes | Ghostweave Vest (14141, -1.09 DPS) [crafted]; Earth Warder's Vest (21311, -1.35 DPS) [quest]; Stone Guard's Satin Armor (220903, -2.01 DPS, sim-verified) [vendor] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 42.3 healing_power points (3.57 DPS) | yes | Forgotten Wraps (9433, -1.26 DPS) [world_drop]; Imperial Red Bracers (8247, -1.45 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -2.54 DPS, sim-verified) [quest] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 61.9 healing_power points (5.23 DPS) | yes | First Sergeant's Satin Gloves (220898, -1.84 DPS) [vendor]; Earthenweave Gloves (254075, -1.86 DPS) [crafted]; Greenleaf Handwraps (19116, -3.51 DPS, sim-verified) [quest] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 55.4 healing_power points (4.68 DPS) | yes | Gilded Waistcord (254081, +0.00 DPS, sim-verified) [crafted]; Gilded Cord (254037, -0.64 DPS) [crafted]; Earthenweave Cord (254077, -1.31 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 69.8 healing_power points (5.89 DPS) | yes | Senior Designer's Pantaloons (11841, -0.52 DPS) [dungeon]; Stone Guard's Satin Leggings (220902, -1.23 DPS, sim-verified) [vendor]; Dalewind Trousers (13008, -1.50 DPS) [world_drop] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 51.5 healing_power points (4.35 DPS) | yes | Vinerot Sandals (17748, -0.54 DPS) [dungeon]; Gilded Slippers (254001, -0.89 DPS) [crafted]; First Sergeant's Satin Boots (220900, -1.05 DPS, sim-verified) [vendor] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 28.8 healing_power points (2.43 DPS) | yes | Ogremind Ring (1993, -0.76 DPS) [world_drop]; Voodoo Band (1996, -0.76 DPS) [world]; Lonetree's Circle (18586, -0.85 DPS) [quest] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | 20.2 healing_power points (1.71 DPS) | yes | Ogremind Ring (1993, -0.04 DPS) [world_drop]; Lonetree's Circle (18586, -0.13 DPS) [quest]; Voodoo Band (1996, -0.70 DPS, sim-verified) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (220.7 DPS) | yes | Evonice's Landin' Pilla (18951, -0.21 DPS) [quest]; Alchemist's Stone (13503, -0.85 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (220.7 DPS) | yes | Evonice's Landin' Pilla (18951, -0.39 DPS, sim-verified) [quest]; Alchemist's Stone (13503, -1.09 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (220.7 DPS) | yes | Charstone Dirk (17710, -6.67 DPS) [dungeon]; Spellshifter Rod (9527, -10.22 DPS) [quest]; Hammer of the Northern Wind (810, -33.41 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 17.5 healing_power points (1.48 DPS) | yes | Wand of Mana Concentration (273043, -0.27 DPS, sim-verified) [world]; Nature's Breath (19118, -0.32 DPS) [quest]; Goblin Igniter (5253, -0.39 DPS) [quest] |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Nethergeld Shoulders; back: Imperial Red Cloak; wrist: Nethergeld Cuffs; hands: Gilded Gloves; waist: Dawnspire Cord; legs: Kilt of the Atal'ai Prophet; feet: Gilded Sandals; finger1: Eye of Adaegus; finger2: Welken Ring; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Jaina's Firestarter

No-known-source sample (15 of 507, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60 (undead, 025005031330011530-03505003030000000-00000000000000000)

Set DPS (verified): 312.7. Weights run: 12.1s. Verify run: 48.3s. 1193 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.063, intellect=1.486 ± 0.027, spirit=4.369 ± 0.040, mp5=1.706 ± 0.039, crit=0.322 ± 0.015 per rating point (14 rating = 1%, 4.511 per %), spell_haste=-9.472 ± 0.574

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | 121.2 healing_power points (17.95 DPS) | yes | Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Earthenweave Crown (254137, +0.00 DPS, sim-verified) [crafted] |
| neck | Lady Alizabeth's Pendant (13002) | World drop [world_drop] | 65.5 healing_power points (9.71 DPS) | yes | Amberglow Talisman (10824, +0.00 DPS, sim-verified) [quest]; Choker of the High Shaman (4112, -3.43 DPS) [quest]; Necklace of Calisea (1714, -3.64 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (+10.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Earthenweave Mantle (254119, -0.46 DPS) [crafted]; Mooncloth Shoulders (14139, -10.85 DPS, sim-verified) [crafted] |
| back | Cloak of Blight (6832) (or Indomitable Cloak (14683)) | Nothing But The Truth [quest] | 56.8 healing_power points (8.41 DPS) | yes | Indomitable Cloak (14683, +0.00 DPS) [world_drop]; Wingveil Cloak (10802, -0.55 DPS) [dungeon]; Arcane Cloak (8286, -0.65 DPS) [world_drop] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 142.9 healing_power points (21.17 DPS) | yes | Mooncloth Vest (14138, -0.65 DPS, sim-verified) [crafted]; Legionnaire's Satin Tunic (227124, -2.88 DPS) [pvp]; Warlord's Satin Tunic (231632, -3.11 DPS) [vendor] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | 66.6 healing_power points (9.86 DPS) | yes | General's Satin Bracers (17619, +0.00 DPS) [pvp]; Earthenweave Cuffs (254131, +0.00 DPS, sim-verified) [crafted]; Nethergeld Cuffs (254061, -2.12 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | 104.1 healing_power points (15.42 DPS) | yes | Swarmtender's Gloves (275607, +0.00 DPS, sim-verified) [crafted]; Magister's Gloves (16684, -3.28 DPS) [dungeon]; Gilded Gloves (254095, -3.75 DPS) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (+7.6 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Satin Cinch (17621, +0.00 DPS) [pvp]; Mana-infused Cord (279267, -0.21 DPS) [crafted]; Wisdom of the Timbermaw (19047, -7.62 DPS, sim-verified) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 146.8 healing_power points (21.74 DPS) | yes | Legionnaire's Satin Legguards (227123, -3.45 DPS) [pvp]; Mooncloth Leggings (14137, -3.56 DPS) [crafted]; Dalewind Trousers (13008, -3.58 DPS) [world_drop] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (+9.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Elunarian Boots (14458, -2.51 DPS) [world_drop]; Furen's Boots (13100, -2.94 DPS) [world_drop]; Mooncloth Boots (15802, -9.08 DPS, sim-verified) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 62.9 healing_power points (9.32 DPS) | yes | Coldwater Ring (4550, -2.84 DPS) [quest]; Welken Ring (5011, -3.69 DPS) [world_drop]; The Queen's Jewel (13094, -3.70 DPS) [world_drop] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 61.3 healing_power points (9.09 DPS) | yes | Coldwater Ring (4550, +0.00 DPS, sim-verified) [quest]; Welken Ring (5011, -3.46 DPS) [world_drop]; The Queen's Jewel (13094, -3.47 DPS) [world_drop] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop]; Alchemist's Stone (13503, +0.00 DPS) [crafted]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+9.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Evonice's Landin' Pilla (18951, -1.29 DPS) [quest]; Alchemist's Stone (13503, -5.18 DPS) [crafted]; Uther's Strength (11302, -9.23 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -14.32 DPS) [dungeon]; Staff of Hale Magefire (13000, -14.69 DPS) [world_drop]; Hand of Edward the Odd (2243, -15.04 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 65.1 healing_power points (9.64 DPS) | yes | Chillnail Splinter (10704, -2.22 DPS, sim-verified) [quest]; Jaina's Firestarter (13064, -6.37 DPS) [world_drop]; Captain Rackmore's Tiller (16789, -6.81 DPS) [quest] |

**New at 60:** head: Devout Crown; neck: Lady Alizabeth's Pendant; shoulder: Devout Mantle; back: Cloak of Blight; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Shard of the Splithooves; trinket2: Ankh of Life; ranged: Torch of Light

No-known-source sample (15 of 1193, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60, raid preset (undead, 025005031330011530-03505003030000000-00000000000000000)

Set DPS (verified): 586.8. Weights run: 7.9s. Verify run: 29.7s. 1193 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.094, intellect=1.025 ± 0.024, spirit=0.993 ± 0.023, mp5=2.229 ± 0.026, crit=0.473 ± 0.019 per rating point (14 rating = 1%, 6.616 per %), spell_haste=-5.802 ± 0.482

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (586.8 DPS) | yes | Mooncloth Circlet (14140, +0.00 DPS) [crafted]; Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 21.2 healing_power points (5.74 DPS) | yes | Horizon Choker (13085, -0.78 DPS) [world_drop]; Lady Alizabeth's Pendant (13002, -1.71 DPS) [world_drop]; Necklace of Calisea (1714, -1.91 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (586.8 DPS) | yes | Mooncloth Shoulders (14139, +0.00 DPS) [crafted]; Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp] |
| back | Wingveil Cloak (10802) | Sunken Temple: Spawn of Hakkar [dungeon] | 15.9 healing_power points (4.31 DPS) | yes | Imperial Red Cloak (8248, -0.18 DPS) [world_drop]; Glowing Thresher Cape (6901, -0.26 DPS) [dungeon]; Highborne Cloak (14450, -0.74 DPS) [world_drop] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 88.9 healing_power points (24.05 DPS) | yes | Warlord's Satin Tunic (231632, -2.77 DPS) [vendor]; Legionnaire's Satin Tunic (227124, -5.31 DPS) [pvp]; Stormcloth Vest (10020, -6.21 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (586.8 DPS) | yes | General's Satin Bracers (17619, +0.00 DPS) [pvp]; Nethergeld Cuffs (254061, +0.00 DPS) [crafted]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (586.8 DPS) | yes | Blood Guard's Satin Handwraps (227126, +0.00 DPS) [pvp]; General's Satin Handwraps (231633, +0.00 DPS) [vendor]; Gilded Gloves (254095, +0.00 DPS) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 54.3 healing_power points (14.69 DPS) | yes | Devout Belt (16696, -3.75 DPS) [dungeon]; General's Satin Cinch (17621, -4.28 DPS) [pvp]; Gilded Waistcord (254081, -4.35 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 62.2 healing_power points (16.84 DPS) | yes | Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Earthenweave Leggings (254125, -2.52 DPS) [crafted] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 51.3 healing_power points (13.87 DPS) | yes | General's Satin Walkers (231630, -1.03 DPS) [vendor]; Swarmtender's Footpads (275609, -2.24 DPS) [crafted]; Blood Guard's Satin Walkers (227127, -2.67 DPS) [pvp] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 21.2 healing_power points (5.73 DPS) | yes | Chemist's Ring (15702, -1.88 DPS) [quest]; Maiden's Circle (13001, -2.18 DPS) [world_drop]; Electrocutioner Lagnut (9447, -2.23 DPS) [dungeon] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 18.1 healing_power points (4.89 DPS) | yes | Chemist's Ring (15702, -1.04 DPS) [quest]; Maiden's Circle (13001, -1.33 DPS) [world_drop]; Electrocutioner Lagnut (9447, -1.38 DPS) [dungeon] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (586.8 DPS) | yes | Ankh of Life (1713, -1.65 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -2.18 DPS) [quest]; Alchemist's Stone (13503, -3.80 DPS) [crafted] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (586.8 DPS) | yes | Ankh of Life (1713, -0.39 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.93 DPS) [quest]; Alchemist's Stone (13503, -2.54 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -15.41 DPS) [dungeon]; Hand of Edward the Odd (2243, -20.41 DPS) [world_drop]; Hand of Righteousness (7721, -23.47 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 27.9 healing_power points (7.56 DPS) | yes | Jaina's Firestarter (13064, -5.09 DPS) [world_drop]; Captain Rackmore's Tiller (16789, -5.13 DPS) [quest]; Wand of Mana Concentration (273043, -5.24 DPS) [world] |

**New at 60:** head: Devout Crown; neck: Glowing Eye of Mordresh; shoulder: Devout Mantle; back: Wingveil Cloak; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Uther's Strength; trinket2: Shard of the Splithooves; ranged: Torch of Light

No-known-source sample (15 of 1193, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

