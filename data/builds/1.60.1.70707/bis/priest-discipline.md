# Leveling BiS: Discipline

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 025004000000000000-00000000000000000-00000000000000000)

Set DPS (verified): 30.6. Weights run: 8.2s. Verify run: 25.2s. 163 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.007, intellect=0.905 ± 0.009, spirit=0.775 ± 0.007, mp5=1.649 ± 0.011, crit=0.044 ± 0.003 per rating point (14 rating = 1%, 0.611 per %), spell_haste=not significant (0.046 ± 0.027)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.32 DPS) | yes | Shadow Goggles (4373, -0.65 DPS) [crafted]; Flying Tiger Goggles (4368, -1.09 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 3.1 healing_power points (0.23 DPS) | yes | Scholarly Pendant (277203, -0.11 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.1 healing_power points (0.60 DPS) | yes | Slime-encrusted Pads (6461, -0.23 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -0.33 DPS) [crafted] |
| back | Sanguine Cape (14376) | World drop [world_drop] | 3.6 healing_power points (0.27 DPS) | yes | Seer's Cape (6378, -0.02 DPS) [dungeon]; Spirit Cloak (4792, -0.04 DPS) [vendor]; Sylvan Cloak (4793, -0.04 DPS) [vendor] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 19.9 healing_power points (1.45 DPS) | yes | Corsair's Overshirt (5202, -0.19 DPS) [dungeon]; Robe of the Moccasin (6465, -0.23 DPS) [dungeon]; Bloody Apron (6226, -0.50 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 9.5 healing_power points (0.69 DPS) | yes | Bright Bracers (3647, -0.43 DPS) [world_drop]; Repurposed Hair Band (281256, -0.45 DPS) [quest]; Seer's Cuffs (3645, -0.51 DPS) [dungeon] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 13.7 healing_power points (1.00 DPS) | yes | Magefist Gloves (12977, -0.45 DPS) [world_drop]; Bright Gloves (3066, -0.51 DPS) [world_drop]; Tomb Robber's Gloves (280096, -0.61 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 14.6 healing_power points (1.07 DPS) | yes | Novice Ardent's Sash (253887, -0.26 DPS) [crafted]; Keller's Girdle (2911, -0.54 DPS) [world_drop]; Relic Hunter Belt (11936, -0.69 DPS) [quest] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 26.5 healing_power points (1.94 DPS) | yes | Darkweave Breeches (12987, -1.14 DPS) [world_drop]; Filigreed Silky Leggings (253939, -1.32 DPS) [crafted]; Filigreed Shadow Leggings (253943, -1.32 DPS) [crafted] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 11.7 healing_power points (0.86 DPS) | yes | Kimbra Boots (6191, -0.42 DPS) [quest]; Bluegill Sandals (1560, -0.50 DPS) [world]; Smoldering Boots (3076, -0.52 DPS) [world] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.4 healing_power points (0.40 DPS) | yes | Band of Purification (12996, -0.06 DPS) [world_drop]; Deep Fathom Ring (6463, -0.11 DPS) [dungeon]; Minor Channeling Ring (1449, -0.27 DPS) [quest] |
| finger2 | Volcanic Rock Ring (12053) | World drop [world_drop] | sim-verified (30.6 DPS) | yes | Deep Fathom Ring (6463, +0.00 DPS) [dungeon]; Band of Purification (12996, +0.00 DPS) [world_drop]; Minor Channeling Ring (1449, -0.07 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | sim-verified (30.6 DPS) | yes | Staff of Westfall (2042, +0.00 DPS) [quest]; Staff of the Blessed Seer (2271, -0.09 DPS) [dungeon]; Evocator's Blade (2567, -0.25 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Flaring Baton (5326) | The Escape [quest] | sim-verified (30.6 DPS) | yes | Sable Wand (7607, +0.00 DPS) [quest]; Moonstone Wand (15204, +0.00 DPS) [quest]; Dwarven Flamestick (5241, -0.02 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Tarnished Locket; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Volcanic Rock Ring; main_hand: Twisted Chanter's Staff; ranged: Flaring Baton

No-known-source sample (15 of 163, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes; 16607 Acolyte's Sacrificial Robes; 18851 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (human, 025005031320000000-00000000000000000-00000000000000000)

Set DPS (verified): 58.2. Weights run: 8.4s. Verify run: 28.4s. 294 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.135 ± 0.008, spirit=1.214 ± 0.008, mp5=2.593 ± 0.008, crit=0.078 ± 0.005 per rating point (14 rating = 1%, 1.097 per %), spell_haste=not significant (0.004 ± 0.023)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 40.3 healing_power points (3.41 DPS) | yes | Filigreed Pristine Circlet (253975, -1.55 DPS) [crafted]; Nightsky Cowl (4039, -1.74 DPS) [world_drop]; Pristine Circlet (253949, -1.88 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 9.4 healing_power points (0.79 DPS) | yes | Crystal Starfire Medallion (5003, -0.10 DPS) [world_drop]; Pendant of Myzrael (4614, -0.18 DPS) [dungeon]; Glowing Green Talisman (5002, -0.18 DPS) [world_drop] |
| shoulder | Mantle of Honor (3560) | Bride of the Embalmer [quest] | 16.4 healing_power points (1.39 DPS) | yes | Nightsky Mantle (4718, -0.11 DPS) [world_drop]; Faerie Mantle (5820, -0.20 DPS) [quest]; Death Speaker Mantle (6685, -0.33 DPS) [dungeon] |
| back | Prelacy Cape (7004) | Researching the Corruption [quest] | 18.3 healing_power points (1.55 DPS) | yes | Glowing Thresher Cape (6901, -0.13 DPS) [dungeon]; Repairman's Cape (9605, -0.41 DPS) [quest]; Sergeant Major's Cape (16315, -0.75 DPS) [pvp] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (58.2 DPS) | yes | Pristine Gown (253961, +0.00 DPS) [crafted]; Robes of Arugal (6324, -0.61 DPS) [dungeon]; Beguiler Robes (7728, -0.77 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 13.5 healing_power points (1.14 DPS) | yes | Nightsky Wristbands (6407, -0.25 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.37 DPS) [world_drop]; Spidertank Oilrag (9448, -0.48 DPS) [dungeon] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 20.2 healing_power points (1.71 DPS) | yes | Truefaith Gloves (7049, -0.16 DPS) [crafted]; Hotshot Pilot's Gloves (9491, -0.43 DPS) [dungeon]; Pristine Gloves (253913, -0.49 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 15.5 healing_power points (1.31 DPS) | yes | Novice Ardent's Sash (253887, -0.21 DPS) [crafted]; Resilient Cord (14406, -0.33 DPS) [world_drop]; Dreamer's Belt (4829, -0.44 DPS) [vendor] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 36.0 healing_power points (3.05 DPS) | yes | Filigreed Pristine Leggings (253937, -0.54 DPS) [crafted]; Earthen Leggings (253999, -0.65 DPS) [crafted]; Darkweave Breeches (12987, -1.76 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 32.8 healing_power points (2.77 DPS) | yes | Nimbus Boots (6998, -1.34 DPS) [quest]; Acidic Walkers (9454, -1.60 DPS) [dungeon]; Frothing Slippers (254003, -1.69 DPS) [crafted] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 14.5 healing_power points (1.23 DPS) | yes | The Queen's Jewel (13094, -0.21 DPS) [world_drop]; Ring of Calm (6790, -0.51 DPS) [quest]; Defias Mage Ring (1077, -0.53 DPS) [world] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 14.5 healing_power points (1.23 DPS) | yes | The Queen's Jewel (13094, -0.21 DPS) [world_drop]; Ring of Calm (6790, -0.51 DPS) [quest]; Defias Mage Ring (1077, -0.53 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, +0.00 DPS) [dungeon]; Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -0.43 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 9.7 healing_power points (0.82 DPS) | yes | Consecrated Wand (5244, -0.43 DPS) [quest]; Summoner's Wand (5245, -0.44 DPS) [world]; Starfaller (13063, -0.44 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Mantle of Honor; back: Prelacy Cape; chest: Death Speaker Robes; hands: Gloves of Old; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Spine of the Devourer; ranged: Wand of Mana Concentration

No-known-source sample (15 of 294, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 40 (human, 025005031330011520-00000000000000000-00000000000000000)

Set DPS (verified): 115.6. Weights run: 10.0s. Verify run: 64.9s. 396 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.024, intellect=1.196 ± 0.025, spirit=0.461 ± 0.063, mp5=3.509 ± 0.041, crit=0.208 ± 0.009 per rating point (14 rating = 1%, 2.907 per %), spell_haste=-9.273 ± 0.388

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 38.5 healing_power points (3.35 DPS) | yes | Whitemane's Chapeau (7720, -0.69 DPS) [dungeon]; Holy Shroud (2721, -0.81 DPS, sim-verified) [world_drop]; Earthen Silk Hood (254015, -1.03 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 22.6 healing_power points (1.96 DPS) | yes | Archmagister's Faceted Pendant (273049, -1.34 DPS) [world]; Kaleidoscope Chain (13084, -1.39 DPS) [world_drop]; Necklace of Calisea (1714, -2.69 DPS, sim-verified) [dungeon] |
| shoulder | Earthen Silk Shoulders (254033) | Tailoring [crafted] | sim-verified (115.6 DPS) | yes | Inquisitor's Shawl (19507, -0.39 DPS, sim-verified) [dungeon]; Windchaser Amice (14432, -0.53 DPS) [world_drop]; Mistscape Mantle (4734, -0.54 DPS) [dungeon] |
| back | Prelacy Cape (7004) | Researching the Corruption [quest] | 13.8 healing_power points (1.20 DPS) | yes | Glowing Thresher Cape (6901, -0.27 DPS) [dungeon]; Sergeant Major's Cape (16336, -0.33 DPS) [pvp]; Blackforge Cape (6424, -0.79 DPS, sim-verified) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 61.7 healing_power points (5.36 DPS) | yes | Pristine Gown (253961, -2.74 DPS) [crafted]; Civinad Robes (9623, -3.40 DPS) [quest]; Death Speaker Robes (6682, -7.07 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 16.5 healing_power points (1.43 DPS) | yes | Earthen Silk Cuffs (254019, +0.00 DPS, sim-verified) [crafted]; Windchaser Cuffs (14429, -0.50 DPS) [world_drop]; Spidertank Oilrag (9448, -0.52 DPS) [dungeon] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (115.6 DPS) | yes | Earthen Silk Gloves (254017, -0.32 DPS) [crafted]; Truefaith Gloves (7049, -0.59 DPS) [crafted]; Gilded Handwraps (254021, -1.29 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 34.3 healing_power points (2.98 DPS) | yes | Pristine Sash (253925, -1.61 DPS) [crafted]; Sutarn's Ring (13105, -1.70 DPS) [world_drop]; Deathmage Sash (10771, -2.96 DPS, sim-verified) [dungeon] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-verified (115.6 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS) [crafted]; Pristine Leggings (253987, +0.00 DPS) [crafted]; Stoneweaver Leggings (9407, -1.84 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 30.2 healing_power points (2.63 DPS) | yes | Furen's Boots (13100, -1.26 DPS) [world_drop]; Thoughtcast Boots (10578, -1.41 DPS) [dungeon]; Nimbus Boots (6998, -1.61 DPS, sim-verified) [quest] |
| finger1 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 10.5 healing_power points (0.92 DPS) | yes | Ogremind Ring (1993, -0.07 DPS) [world_drop]; Mindbender Loop (5009, -0.11 DPS) [world_drop]; Electrocutioner Lagnut (9447, -0.11 DPS) [dungeon] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 9.8 healing_power points (0.85 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Mindbender Loop (5009, -0.04 DPS) [world_drop]; Electrocutioner Lagnut (9447, -0.05 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (115.6 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Death Speaker Scepter (2816, -2.58 DPS) [dungeon]; Hypnotic Blade (7714, -3.72 DPS) [dungeon] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 24.0 healing_power points (2.08 DPS) | yes | Orb of Souls (249395, -0.71 DPS) [crafted]; Eye of Paleth (2943, -0.95 DPS) [quest]; Beacon of Hope (9393, -2.24 DPS, sim-verified) [dungeon] |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 11.8 healing_power points (1.03 DPS) | yes | Jaina's Firestarter (13064, +0.00 DPS, sim-verified) [world_drop]; Captain Rackmore's Tiller (16789, -0.38 DPS) [quest]; Goblin Igniter (5253, -0.49 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Earthen Silk Shoulders; chest: Stormcloth Vest; hands: Stormcloth Gloves; waist: Gilded Cord; legs: Stormcloth Pants; finger1: Ring of Power Regulation; finger2: Voodoo Band; trinket1: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Orb of Lorica

No-known-source sample (15 of 396, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 50 (human, 025005031330011530-03501000000000000-00000000000000000)

Set DPS (verified): 188.6. Weights run: 10.0s. Verify run: 37.9s. 515 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.012, intellect=1.564 ± 0.014, spirit=1.239 ± 0.042, mp5=3.865 ± 0.063, crit=0.490 ± 0.016 per rating point (14 rating = 1%, 6.861 per %), spell_haste=-10.860 ± 0.492

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 57.9 healing_power points (4.56 DPS) | yes | Knight-Lieutenant's Satin Cover (220896, +0.00 DPS) [vendor]; Stormcloth Headband (10032, -0.46 DPS) [crafted]; Cassandra's Grace (13102, -1.03 DPS, sim-verified) [world_drop] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 26.8 healing_power points (2.12 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -0.57 DPS) [dungeon]; Mindburst Medallion (11196, -0.89 DPS) [quest] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | 45.9 healing_power points (3.62 DPS) | yes | Knight-Lieutenant's Satin Pads (220894, +0.00 DPS) [vendor]; Rotgrip Mantle (17732, -1.40 DPS) [dungeon]; Inquisitor's Shawl (19507, -1.89 DPS, sim-verified) [dungeon] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 22.2 healing_power points (1.75 DPS) | yes | Highborne Cloak (14450, -0.17 DPS, sim-verified) [world_drop]; Wingveil Cloak (10802, -0.28 DPS) [dungeon]; Prelacy Cape (7004, -0.29 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 67.9 healing_power points (5.35 DPS) | yes | Ghostweave Vest (14141, -0.96 DPS) [crafted]; Knight's Satin Armor (220892, -1.06 DPS, sim-verified) [vendor]; Grimlok's Tribal Vestments (9415, -1.87 DPS) [dungeon] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 37.1 healing_power points (2.93 DPS) | yes | Mindthrust Bracers (1974, -1.40 DPS) [dungeon]; Forgotten Wraps (9433, -1.45 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -1.72 DPS, sim-verified) [quest] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 54.5 healing_power points (4.30 DPS) | yes | Earthenweave Gloves (254075, -0.78 DPS, sim-verified) [crafted]; Sergeant Major's Satin Gloves (220897, -1.36 DPS) [vendor]; Gilded Handwraps (254021, -1.70 DPS) [crafted] |
| waist | Gilded Waistcord (254081) | Tailoring [crafted] | 43.1 healing_power points (3.39 DPS) | yes | Gilded Cord (254037, +0.00 DPS, sim-verified) [crafted]; Dawnspire Cord (12466, -0.11 DPS) [dungeon]; Earthenweave Cord (254077, -0.27 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | sim-verified (188.6 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -0.14 DPS) [dungeon]; Dalewind Trousers (13008, -1.04 DPS) [world_drop]; Knight's Satin Leggings (220893, -2.06 DPS, sim-verified) [vendor] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 45.1 healing_power points (3.55 DPS) | yes | Earthenweave Boots (254093, +0.00 DPS, sim-verified) [crafted]; Sergeant Major's Satin Boots (220895, -0.50 DPS) [vendor]; Vinerot Sandals (17748, -0.52 DPS) [dungeon] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 24.2 healing_power points (1.91 DPS) | yes | Choking Band (11868, -0.64 DPS) [quest]; Coldwater Ring (4550, -0.66 DPS) [quest]; Electrocutioner Lagnut (9447, -0.75 DPS) [dungeon] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | 16.5 healing_power points (1.30 DPS) | yes | Choking Band (11868, +0.00 DPS, sim-verified) [quest]; Coldwater Ring (4550, -0.05 DPS) [quest]; Electrocutioner Lagnut (9447, -0.14 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+4.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -0.39 DPS) [quest]; Alchemist's Stone (13503, -0.78 DPS) [crafted] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -0.20 DPS) [quest]; Alchemist's Stone (13503, -0.59 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -5.17 DPS) [dungeon]; Hand of Righteousness (7721, -8.08 DPS) [dungeon]; Hammer of the Northern Wind (810, -27.34 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cairnstone Sliver (9654) | The Morrow Stone [quest] | 14.0 healing_power points (1.10 DPS) | yes | Wand of Mana Concentration (273043, -0.00 DPS) [world]; Jaina's Firestarter (13064, -0.07 DPS) [world_drop]; Goblin Igniter (5253, -0.32 DPS) [quest] |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Nethergeld Shoulders; back: Imperial Red Cloak; wrist: Nethergeld Cuffs; hands: Gilded Gloves; waist: Gilded Waistcord; legs: Senior Designer's Pantaloons; feet: Gilded Sandals; finger1: Eye of Adaegus; finger2: Welken Ring; trinket2: Evonice's Landin' Pilla; main_hand: Glowing Brightwood Staff; ranged: Cairnstone Sliver

No-known-source sample (15 of 515, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60 (human, 025005031330011530-03505003030000000-00000000000000000)

Set DPS (verified): 276.8. Weights run: 9.3s. Verify run: 67.9s. 1200 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.055, intellect=1.268 ± 0.030, spirit=4.614 ± 0.045, mp5=1.760 ± 0.039, crit=0.301 ± 0.015 per rating point (14 rating = 1%, 4.220 per %), spell_haste=-5.782 ± 0.568

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Circlet (14140, +0.00 DPS) [crafted]; Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Earthenweave Crown (254137, +0.00 DPS) [crafted] |
| neck | Lady Alizabeth's Pendant (13002) | World drop [world_drop] | 69.2 healing_power points (9.60 DPS) | yes | Amberglow Talisman (10824, -3.20 DPS) [quest]; Explorers' League Commendation (7746, -3.84 DPS) [quest]; Necklace of Calisea (1714, -3.89 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Shoulders (14139, +0.00 DPS) [crafted]; Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Earthenweave Mantle (254119, +0.00 DPS) [crafted] |
| back | Ethereal Mist Cape (11873) | Ancient Spirit [quest] | sim-verified (276.8 DPS) | yes | Indomitable Cloak (14683, +0.00 DPS) [world_drop]; Wingveil Cloak (10802, -0.47 DPS) [dungeon]; Arcane Cloak (8286, -0.53 DPS) [world_drop] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 146.8 healing_power points (20.37 DPS) | yes | Mooncloth Vest (14138, -1.08 DPS) [crafted]; Knight-Captain's Satin Tunic (227122, -3.22 DPS) [pvp]; Devout Robe (16690, -3.77 DPS) [dungeon] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Satin Bracers (17606, +0.00 DPS) [pvp]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -2.13 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swarmtender's Gloves (275607, +0.00 DPS) [crafted]; Magister's Gloves (16684, -3.30 DPS) [dungeon]; Gilded Gloves (254095, -3.82 DPS) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (276.8 DPS) | yes | Marshal's Satin Sash (17609, +0.00 DPS) [pvp]; Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Mana-infused Cord (279267, +0.00 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 149.1 healing_power points (20.70 DPS) | yes | Mooncloth Leggings (14137, -2.95 DPS) [crafted]; Dalewind Trousers (13008, -3.11 DPS) [world_drop]; Knight-Captain's Satin Legguards (227125, -3.54 DPS) [pvp] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | sim-verified (276.8 DPS) | yes | Devout Sandals (16691, +0.00 DPS) [dungeon]; Elunarian Boots (14458, -2.18 DPS) [world_drop]; Rime-encrusted Boots (279268, -2.59 DPS) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 63.4 healing_power points (8.80 DPS) | yes | Choking Band (11868, -0.48 DPS) [quest]; Seedtime Hoop (9655, -2.40 DPS) [quest]; Coldwater Ring (4550, -2.51 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 63.0 healing_power points (8.74 DPS) | yes | Choking Band (11868, -0.42 DPS) [quest]; Seedtime Hoop (9655, -2.34 DPS) [quest]; Coldwater Ring (4550, -2.45 DPS) [quest] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (276.8 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (276.8 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Edward the Odd (2243, -10.50 DPS) [world_drop]; Spire of Hakkar (10844, -12.83 DPS) [world]; Soulkeeper (1607, -13.10 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 67.8 healing_power points (9.40 DPS) | yes | Cairnstone Sliver (9654, -5.32 DPS) [quest]; Jaina's Firestarter (13064, -6.42 DPS) [world_drop]; Captain Rackmore's Tiller (16789, -6.65 DPS) [quest] |

**New at 60:** head: Devout Crown; neck: Lady Alizabeth's Pendant; shoulder: Devout Mantle; back: Ethereal Mist Cape; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Shard of the Splithooves; trinket2: Uther's Strength; ranged: Torch of Light

No-known-source sample (15 of 1200, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60, raid preset (human, 025005031330011530-03505003030000000-00000000000000000)

Set DPS (verified): 523.0. Weights run: 6.4s. Verify run: 24.6s. 1200 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.129, intellect=0.865 ± 0.027, spirit=1.189 ± 0.027, mp5=2.585 ± 0.031, crit=0.510 ± 0.020 per rating point (14 rating = 1%, 7.133 per %), spell_haste=-6.383 ± 0.518

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (523.1 DPS) | yes | Mooncloth Circlet (14140, +0.00 DPS) [crafted]; Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 19.9 healing_power points (4.77 DPS) | yes | Lady Alizabeth's Pendant (13002, -0.50 DPS) [world_drop]; Horizon Choker (13085, -0.73 DPS) [world_drop]; Necklace of Calisea (1714, -1.33 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (523.1 DPS) | yes | Mooncloth Shoulders (14139, +0.00 DPS) [crafted]; Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp] |
| back | Sergeant Major's Cape (16337) | PvP rank 9 · Sergeant Major · Alliance [pvp] | sim-verified (523.1 DPS) | yes | Prelacy Cape (7004, +0.00 DPS) [quest]; Wingveil Cloak (10802, +0.00 DPS) [dungeon]; Ethereal Mist Cape (11873, +0.00 DPS) [quest] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 92.0 healing_power points (22.04 DPS) | yes | Field Marshal's Satin Tunic (231624, -3.85 DPS) [vendor]; Knight-Captain's Satin Tunic (227122, -5.37 DPS) [pvp]; Stormcloth Vest (10020, -5.87 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (523.1 DPS) | yes | Nethergeld Cuffs (254061, +0.00 DPS) [crafted]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted]; Marshal's Satin Bracers (17606, -0.35 DPS) [pvp] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (523.1 DPS) | yes | Knight-Lieutenant's Satin Handwraps (227128, +0.00 DPS) [pvp]; Gilded Gloves (254095, +0.00 DPS) [crafted]; Marshal's Satin Handwraps (231623, -0.86 DPS) [vendor] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 53.8 healing_power points (12.89 DPS) | yes | Earthenweave Cord (254077, -3.53 DPS) [crafted]; Devout Belt (16696, -3.55 DPS) [dungeon]; Marshal's Satin Sash (17609, -3.85 DPS) [pvp] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 64.3 healing_power points (15.41 DPS) | yes | Knight-Captain's Satin Legguards (227125, +0.00 DPS) [pvp]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Earthenweave Leggings (254125, -1.99 DPS) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (523.1 DPS) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; Marshal's Satin Walkers (231627, +0.00 DPS) [vendor]; Swarmtender's Footpads (275609, +0.00 DPS) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 21.7 healing_power points (5.20 DPS) | yes | Choking Band (11868, -1.50 DPS) [quest]; Electrocutioner Lagnut (9447, -1.77 DPS) [dungeon]; Runemender's Seal (273031, -1.77 DPS) [world] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 19.5 healing_power points (4.66 DPS) | yes | Choking Band (11868, -0.96 DPS) [quest]; Electrocutioner Lagnut (9447, -1.23 DPS) [dungeon]; Runemender's Seal (273031, -1.23 DPS) [world] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (523.1 DPS) | yes | Shard of the Splithooves (10659, -0.60 DPS) [quest]; Evonice's Landin' Pilla (18951, -1.46 DPS) [quest]; Thunderbrew's Boot Flask (744, -2.03 DPS) [quest] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (523.1 DPS) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Evonice's Landin' Pilla (18951, -0.57 DPS) [quest]; Thunderbrew's Boot Flask (744, -1.14 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -13.20 DPS) [dungeon]; Hand of Edward the Odd (2243, -17.40 DPS) [world_drop]; Hand of Righteousness (7721, -19.71 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 30.1 healing_power points (7.20 DPS) | yes | Cairnstone Sliver (9654, -4.74 DPS) [quest]; Captain Rackmore's Tiller (16789, -4.91 DPS) [quest]; Jaina's Firestarter (13064, -5.11 DPS) [world_drop] |

**New at 60:** head: Devout Crown; neck: Glowing Eye of Mordresh; shoulder: Devout Mantle; back: Sergeant Major's Cape; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Uther's Strength; trinket2: Ankh of Life; ranged: Torch of Light

No-known-source sample (15 of 1200, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

## Horde

### Band 20 (undead, 025004000000000000-00000000000000000-00000000000000000)

Set DPS (verified): 33.0. Weights run: 8.2s. Verify run: 26.3s. 156 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.007, intellect=0.905 ± 0.009, spirit=0.775 ± 0.007, mp5=1.649 ± 0.011, crit=0.044 ± 0.003 per rating point (14 rating = 1%, 0.611 per %), spell_haste=not significant (0.046 ± 0.027)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.32 DPS) | yes | Shadow Goggles (4373, -0.65 DPS) [crafted]; Flying Tiger Goggles (4368, -1.09 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 2.3 healing_power points (0.17 DPS) | yes | Scholarly Pendant (277203, -0.06 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.1 healing_power points (0.60 DPS) | yes | Slime-encrusted Pads (6461, -0.23 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -0.33 DPS) [crafted] |
| back | Sanguine Cape (14376) | World drop [world_drop] | 3.6 healing_power points (0.27 DPS) | yes | Seer's Cape (6378, -0.02 DPS) [dungeon]; Spirit Cloak (4792, -0.04 DPS) [vendor]; Antiquated Cloak (4799, -0.04 DPS) [vendor] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 19.9 healing_power points (1.45 DPS) | yes | Corsair's Overshirt (5202, -0.19 DPS) [dungeon]; Robe of the Moccasin (6465, -0.23 DPS) [dungeon]; Bloody Apron (6226, -0.50 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 9.5 healing_power points (0.69 DPS) | yes | Featherbead Bracers (15452, -0.36 DPS) [quest]; Crystalline Cuffs (14148, -0.39 DPS) [dungeon]; Bright Bracers (3647, -0.43 DPS) [world_drop] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 13.7 healing_power points (1.00 DPS) | yes | Blight Gloves (279877, -0.26 DPS) [quest]; Magefist Gloves (12977, -0.45 DPS) [world_drop]; Bright Gloves (3066, -0.51 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 14.6 healing_power points (1.07 DPS) | yes | Novice Ardent's Sash (253887, -0.26 DPS) [crafted]; Keller's Girdle (2911, -0.54 DPS) [world_drop]; Wise Man's Belt (4786, -0.74 DPS) [vendor] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 26.5 healing_power points (1.94 DPS) | yes | Darkweave Breeches (12987, -1.14 DPS) [world_drop]; Filigreed Silky Leggings (253939, -1.32 DPS) [crafted]; Filigreed Shadow Leggings (253943, -1.32 DPS) [crafted] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 11.7 healing_power points (0.86 DPS) | yes | Bluegill Sandals (1560, -0.50 DPS) [world]; Smoldering Boots (3076, -0.52 DPS) [world]; Red Woolen Boots (4313, -0.57 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.4 healing_power points (0.40 DPS) | yes | Band of Purification (12996, -0.06 DPS) [world_drop]; Deep Fathom Ring (6463, -0.11 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.20 DPS) [world_drop] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | sim-verified (33.0 DPS) | yes | Band of Purification (12996, +0.00 DPS) [world_drop]; Deep Fathom Ring (6463, -0.05 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.13 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) | The Wrath of Rath'mael [quest] | sim-verified (33.0 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Staff of Orgrimmar (15444, +0.00 DPS) [quest]; Staff of the Blessed Seer (2271, -0.09 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Decay (5252) (or Flaring Baton (5326)) | Beren's Peril [quest] | 1.8 healing_power points (0.13 DPS) | yes | Flaring Baton (5326, +0.00 DPS) [quest]; Wisesight Wand (286750, -0.07 DPS) [world] |

**New at 20:** head: Pristine Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Gnarled Necromancer's Staff; ranged: Wand of Decay

No-known-source sample (15 of 156, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (undead, 025005031320000000-00000000000000000-00000000000000000)

Set DPS (verified): 63.1. Weights run: 8.4s. Verify run: 29.5s. 286 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.135 ± 0.008, spirit=1.214 ± 0.008, mp5=2.593 ± 0.008, crit=0.078 ± 0.005 per rating point (14 rating = 1%, 1.097 per %), spell_haste=not significant (0.004 ± 0.023)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 40.3 healing_power points (3.41 DPS) | yes | Nightsky Cowl (4039, -1.74 DPS) [world_drop]; Pristine Circlet (253949, -1.88 DPS) [crafted]; Filigreed Pristine Circlet (253975, -1.91 DPS, sim-verified) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 9.4 healing_power points (0.79 DPS) | yes | Crystal Starfire Medallion (5003, -0.16 DPS, sim-verified) [world_drop]; Pendant of Myzrael (4614, -0.18 DPS) [dungeon]; Glowing Green Talisman (5002, -0.18 DPS) [world_drop] |
| shoulder | Ghostly Mantle (3324) | Deathstalkers in Shadowfang [quest] | 19.9 healing_power points (1.68 DPS) | yes | Nightsky Mantle (4718, -0.54 DPS, sim-verified) [world_drop]; Death Speaker Mantle (6685, -0.63 DPS) [dungeon]; Talbar Mantle (10657, -0.67 DPS) [quest] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 16.7 healing_power points (1.41 DPS) | yes | Nightsky Cloak (4719, -0.80 DPS) [world_drop]; Insignia Cloak (4722, -0.80 DPS) [world_drop]; Amy's Blanket (13005, -0.94 DPS, sim-verified) [world_drop] |
| chest | Pristine Gown (253961) | Tailoring [crafted] | 32.8 healing_power points (2.77 DPS) | yes | Death Speaker Robes (6682, -0.16 DPS, sim-verified) [dungeon]; Robes of Arugal (6324, -0.64 DPS) [dungeon]; Beguiler Robes (7728, -0.80 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 13.5 healing_power points (1.14 DPS) | yes | Glowing Magical Bracelets (13106, -0.37 DPS) [world_drop]; Spidertank Oilrag (9448, -0.48 DPS) [dungeon]; Nightsky Wristbands (6407, -0.48 DPS, sim-verified) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 20.2 healing_power points (1.71 DPS) | yes | Truefaith Gloves (7049, +0.00 DPS, sim-verified) [crafted]; Tattered Mittens (270030, -0.25 DPS) [quest]; Hotshot Pilot's Gloves (9491, -0.43 DPS) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 15.5 healing_power points (1.31 DPS) | yes | Lilac Sash (6780, -0.08 DPS, sim-verified) [quest]; Novice Ardent's Sash (253887, -0.21 DPS) [crafted]; Resilient Cord (14406, -0.33 DPS) [world_drop] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 36.0 healing_power points (3.05 DPS) | yes | Earthen Leggings (253999, -0.65 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.74 DPS, sim-verified) [crafted]; Sacred Burial Trousers (6282, -1.64 DPS) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 32.8 healing_power points (2.77 DPS) | yes | Frothing Slippers (254003, -1.69 DPS) [crafted]; Black Slippers (254007, -1.69 DPS) [crafted]; Acidic Walkers (9454, -2.27 DPS, sim-verified) [dungeon] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 14.5 healing_power points (1.23 DPS) | yes | Solid Gold Band (279845, -0.12 DPS) [quest]; The Queen's Jewel (13094, -0.21 DPS) [world_drop]; Defias Mage Ring (1077, -0.53 DPS) [world] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 14.5 healing_power points (1.23 DPS) | yes | Solid Gold Band (279845, -0.11 DPS, sim-verified) [quest]; The Queen's Jewel (13094, -0.21 DPS) [world_drop]; Defias Mage Ring (1077, -0.53 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-verified (63.1 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Gnarled Ash Staff (791, -1.50 DPS) [world_drop]; Death Speaker Scepter (2816, -2.05 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 9.7 healing_power points (0.82 DPS) | yes | Lesser Mystic Wand (11289, -0.44 DPS) [crafted]; Starfaller (13063, -0.44 DPS) [world_drop]; Summoner's Wand (5245, -0.62 DPS, sim-verified) [world] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Ghostly Mantle; back: Glowing Thresher Cape; chest: Pristine Gown; hands: Gloves of Old; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Spine of the Devourer; ranged: Wand of Mana Concentration

No-known-source sample (15 of 286, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (undead, 025005031330011520-00000000000000000-00000000000000000)

Set DPS (verified): 120.9. Weights run: 10.0s. Verify run: 38.3s. 388 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.024, intellect=1.196 ± 0.025, spirit=0.461 ± 0.063, mp5=3.509 ± 0.041, crit=0.208 ± 0.009 per rating point (14 rating = 1%, 2.907 per %), spell_haste=-9.273 ± 0.388

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 38.5 healing_power points (3.35 DPS) | yes | Holy Shroud (2721, -0.24 DPS) [world_drop]; Whitemane's Chapeau (7720, -0.69 DPS) [dungeon]; Earthen Silk Hood (254015, -1.03 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 22.6 healing_power points (1.96 DPS) | yes | Necklace of Calisea (1714, -0.95 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.34 DPS) [world]; Kaleidoscope Chain (13084, -1.39 DPS) [world_drop] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | sim-verified (120.9 DPS) | yes | Earthen Silk Shoulders (254033, -0.77 DPS) [crafted]; Windchaser Amice (14432, -1.30 DPS) [world_drop]; Mistscape Mantle (4734, -1.31 DPS) [dungeon] |
| back | Blackforge Cape (6424) | Uldaman: Ancient Treasure [dungeon] | 11.4 healing_power points (0.99 DPS) | yes | Glowing Thresher Cape (6901, -0.06 DPS) [dungeon]; Windchaser Cloak (14430, -0.21 DPS) [world_drop]; Glyphed Cloak (4732, -0.26 DPS) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 61.7 healing_power points (5.36 DPS) | yes | Death Speaker Robes (6682, -2.48 DPS) [dungeon]; Pristine Gown (253961, -2.74 DPS) [crafted]; Civinad Robes (9623, -3.40 DPS) [quest] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 16.5 healing_power points (1.43 DPS) | yes | Earthen Silk Cuffs (254019, -0.37 DPS) [crafted]; Windchaser Cuffs (14429, -0.50 DPS) [world_drop]; Spidertank Oilrag (9448, -0.52 DPS) [dungeon] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (120.9 DPS) | yes | Gilded Handwraps (254021, +0.00 DPS) [crafted]; Earthen Silk Gloves (254017, -0.32 DPS) [crafted]; Truefaith Gloves (7049, -0.59 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 34.3 healing_power points (2.98 DPS) | yes | Deathmage Sash (10771, -1.43 DPS) [dungeon]; Pristine Sash (253925, -1.61 DPS) [crafted]; Sutarn's Ring (13105, -1.70 DPS) [world_drop] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-verified (120.9 DPS) | yes | Stoneweaver Leggings (9407, +0.00 DPS) [dungeon]; Filigreed Pristine Leggings (253937, +0.00 DPS) [crafted]; Pristine Leggings (253987, +0.00 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 30.2 healing_power points (2.63 DPS) | yes | Furen's Boots (13100, -1.26 DPS) [world_drop]; Thoughtcast Boots (10578, -1.41 DPS) [dungeon]; Boots of the Maharishi (9658, -1.45 DPS) [quest] |
| finger1 | Voodoo Band (1996) | Bloodscalp Witch Doctor [world] | sim-verified (120.9 DPS) | yes | Ring of Power Regulation (273030, +0.00 DPS) [dungeon]; Solid Gold Band (279845, -0.04 DPS) [quest]; Mindbender Loop (5009, -0.04 DPS) [world_drop] |
| finger2 | Ogremind Ring (1993) | World drop [world_drop] | sim-verified (120.9 DPS) | yes | Ring of Power Regulation (273030, +0.00 DPS) [dungeon]; Solid Gold Band (279845, -0.04 DPS) [quest]; Mindbender Loop (5009, -0.04 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, -2.58 DPS) [dungeon]; Hypnotic Blade (7714, -3.72 DPS) [dungeon]; Illusionary Rod (7713, -3.76 DPS) [dungeon] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | 15.8 healing_power points (1.38 DPS) | yes | Orb of Souls (249395, -0.01 DPS) [crafted]; Orb of Mistmantle (13031, -0.43 DPS) [world_drop]; Omega Orb (7749, -0.46 DPS) [quest] |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 11.8 healing_power points (1.03 DPS) | yes | Jaina's Firestarter (13064, -0.28 DPS) [world_drop]; Captain Rackmore's Tiller (16789, -0.38 DPS) [quest]; Goblin Igniter (5253, -0.49 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Blackforge Cape; chest: Stormcloth Vest; hands: Stormcloth Gloves; waist: Gilded Cord; legs: Stormcloth Pants; finger1: Voodoo Band; finger2: Ogremind Ring; trinket1: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Beacon of Hope

No-known-source sample (15 of 388, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16336 Sergeant Major's Cape

### Band 50 (undead, 025005031330011530-03501000000000000-00000000000000000)

Set DPS (verified): 198.4. Weights run: 10.0s. Verify run: 37.5s. 507 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.012, intellect=1.564 ± 0.014, spirit=1.239 ± 0.042, mp5=3.865 ± 0.063, crit=0.490 ± 0.016 per rating point (14 rating = 1%, 6.861 per %), spell_haste=-10.860 ± 0.492

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 57.9 healing_power points (4.56 DPS) | yes | Blood Guard's Satin Cover (220899, +0.00 DPS) [vendor]; Cassandra's Grace (13102, -0.21 DPS) [world_drop]; Stormcloth Headband (10032, -0.46 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 26.8 healing_power points (2.12 DPS) | yes | Glowing Eye of Mordresh (10769, -0.11 DPS) [dungeon]; Necklace of Calisea (1714, -0.57 DPS) [dungeon]; Choker of the High Shaman (4112, -0.72 DPS) [quest] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | 45.9 healing_power points (3.62 DPS) | yes | Blood Guard's Satin Pads (220901, +0.00 DPS) [vendor]; Inquisitor's Shawl (19507, -0.84 DPS) [dungeon]; Rotgrip Mantle (17732, -1.40 DPS) [dungeon] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 22.2 healing_power points (1.75 DPS) | yes | Highborne Cloak (14450, -0.25 DPS) [world_drop]; Wingveil Cloak (10802, -0.28 DPS) [dungeon]; Blackforge Cape (6424, -0.37 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 67.9 healing_power points (5.35 DPS) | yes | Stone Guard's Satin Armor (220903, -0.67 DPS) [vendor]; Ghostweave Vest (14141, -0.96 DPS) [crafted]; Grimlok's Tribal Vestments (9415, -1.87 DPS) [dungeon] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 37.1 healing_power points (2.93 DPS) | yes | Shizzle's Nozzle Wiper (11917, -1.16 DPS) [quest]; Mindthrust Bracers (1974, -1.40 DPS) [dungeon]; Forgotten Wraps (9433, -1.45 DPS) [world_drop] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 54.5 healing_power points (4.30 DPS) | yes | Earthenweave Gloves (254075, -1.18 DPS) [crafted]; Greenleaf Handwraps (19116, -1.21 DPS) [quest]; First Sergeant's Satin Gloves (220898, -1.36 DPS) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | sim-verified (198.3 DPS) | yes | Gilded Cord (254037, +0.00 DPS) [crafted]; Gilded Waistcord (254081, +0.00 DPS) [crafted]; Earthenweave Cord (254077, -0.16 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | sim-verified (198.3 DPS) | yes | Senior Designer's Pantaloons (11841, +0.00 DPS) [dungeon]; Stone Guard's Satin Leggings (220902, +0.00 DPS) [vendor]; Dalewind Trousers (13008, -0.90 DPS) [world_drop] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 45.1 healing_power points (3.55 DPS) | yes | Earthenweave Boots (254093, -0.43 DPS) [crafted]; First Sergeant's Satin Boots (220900, -0.50 DPS) [vendor]; Vinerot Sandals (17748, -0.52 DPS) [dungeon] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 24.2 healing_power points (1.91 DPS) | yes | Coldwater Ring (4550, -0.66 DPS) [quest]; Ring of Subtlety (19038, -0.69 DPS) [quest]; Electrocutioner Lagnut (9447, -0.75 DPS) [dungeon] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | 16.5 healing_power points (1.30 DPS) | yes | Coldwater Ring (4550, -0.05 DPS) [quest]; Ring of Subtlety (19038, -0.08 DPS) [quest]; Electrocutioner Lagnut (9447, -0.14 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.20 DPS) [quest]; Alchemist's Stone (13503, -0.78 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (198.3 DPS) | yes | Evonice's Landin' Pilla (18951, -0.44 DPS) [quest]; Alchemist's Stone (13503, -1.03 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -5.17 DPS) [dungeon]; Hand of Righteousness (7721, -8.08 DPS) [dungeon]; Spellshifter Rod (9527, -9.18 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | sim-verified (198.3 DPS) | yes | Wand of Mana Concentration (273043, +0.00 DPS) [world]; Flash Wand (5248, -0.25 DPS) [quest]; Goblin Igniter (5253, -0.25 DPS) [quest] |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Nethergeld Shoulders; back: Imperial Red Cloak; wrist: Nethergeld Cuffs; hands: Gilded Gloves; waist: Dawnspire Cord; legs: Kilt of the Atal'ai Prophet; feet: Gilded Sandals; finger1: Eye of Adaegus; finger2: Welken Ring; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Jaina's Firestarter

No-known-source sample (15 of 507, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60 (undead, 025005031330011530-03505003030000000-00000000000000000)

Set DPS (verified): 293.4. Weights run: 9.3s. Verify run: 69.2s. 1193 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.055, intellect=1.268 ± 0.030, spirit=4.614 ± 0.045, mp5=1.760 ± 0.039, crit=0.301 ± 0.015 per rating point (14 rating = 1%, 4.220 per %), spell_haste=-5.782 ± 0.568

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (293.4 DPS) | yes | Mooncloth Circlet (14140, +0.00 DPS) [crafted]; Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Earthenweave Crown (254137, -6.81 DPS, sim-verified) [crafted] |
| neck | Lady Alizabeth's Pendant (13002) | World drop [world_drop] | 69.2 healing_power points (9.60 DPS) | yes | Amberglow Talisman (10824, -0.72 DPS, sim-verified) [quest]; Choker of the High Shaman (4112, -3.60 DPS) [quest]; Necklace of Calisea (1714, -3.89 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (293.4 DPS) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Earthenweave Mantle (254119, +0.00 DPS) [crafted]; Mooncloth Shoulders (14139, -6.40 DPS, sim-verified) [crafted] |
| back | Cloak of Blight (6832) (or Indomitable Cloak (14683)) | Nothing But The Truth [quest] | 60.0 healing_power points (8.32 DPS) | yes | Indomitable Cloak (14683, +0.00 DPS) [world_drop]; Wingveil Cloak (10802, -0.59 DPS) [dungeon]; Arcane Cloak (8286, -0.64 DPS) [world_drop] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 146.8 healing_power points (20.37 DPS) | yes | Legionnaire's Satin Tunic (227124, -3.22 DPS) [pvp]; Devout Robe (16690, -3.77 DPS) [dungeon]; Mooncloth Vest (14138, -6.95 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (293.4 DPS) | yes | General's Satin Bracers (17619, +0.00 DPS) [pvp]; Nethergeld Cuffs (254061, -2.13 DPS) [crafted]; Earthenweave Cuffs (254131, -8.77 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (293.4 DPS) | yes | Magister's Gloves (16684, -3.30 DPS) [dungeon]; Gilded Gloves (254095, -3.82 DPS) [crafted]; Swarmtender's Gloves (275607, -8.23 DPS, sim-verified) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 90.3 healing_power points (12.53 DPS) | yes | General's Satin Cinch (17621, +0.00 DPS) [pvp]; Devout Belt (16696, -1.72 DPS) [dungeon]; Mana-infused Cord (279267, -10.36 DPS, sim-verified) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 149.1 healing_power points (20.70 DPS) | yes | Dalewind Trousers (13008, -3.11 DPS) [world_drop]; Legionnaire's Satin Legguards (227123, -3.54 DPS) [pvp]; Mooncloth Leggings (14137, -12.55 DPS, sim-verified) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | 102.1 healing_power points (14.17 DPS) | yes | Elunarian Boots (14458, -2.23 DPS) [world_drop]; Rime-encrusted Boots (279268, -2.65 DPS) [crafted]; Mooncloth Boots (15802, -6.47 DPS, sim-verified) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 63.4 healing_power points (8.80 DPS) | yes | Coldwater Ring (4550, -2.51 DPS) [quest]; The Queen's Jewel (13094, -3.33 DPS) [world_drop]; Welken Ring (5011, -3.44 DPS) [world_drop] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 63.0 healing_power points (8.74 DPS) | yes | Coldwater Ring (4550, -1.63 DPS, sim-verified) [quest]; The Queen's Jewel (13094, -3.27 DPS) [world_drop]; Welken Ring (5011, -3.38 DPS) [world_drop] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (293.4 DPS) | yes | Alchemist's Stone (13503, +0.00 DPS) [crafted]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Ankh of Life (1713, -0.88 DPS, sim-verified) [world_drop] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (293.4 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Mark of the Chosen (17774, -4.13 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (293.4 DPS) | yes | Spire of Hakkar (10844, -12.83 DPS) [world]; Soulkeeper (1607, -13.10 DPS) [world_drop]; Hand of Edward the Odd (2243, -16.28 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 67.8 healing_power points (9.40 DPS) | yes | Chillnail Splinter (10704, -4.62 DPS, sim-verified) [quest]; Jaina's Firestarter (13064, -6.42 DPS) [world_drop]; Captain Rackmore's Tiller (16789, -6.65 DPS) [quest] |

**New at 60:** head: Devout Crown; neck: Lady Alizabeth's Pendant; shoulder: Devout Mantle; back: Cloak of Blight; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Shard of the Splithooves; ranged: Torch of Light

No-known-source sample (15 of 1193, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60, raid preset (undead, 025005031330011530-03505003030000000-00000000000000000)

Set DPS (verified): 541.1. Weights run: 6.4s. Verify run: 25.0s. 1193 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.129, intellect=0.865 ± 0.027, spirit=1.189 ± 0.027, mp5=2.585 ± 0.031, crit=0.510 ± 0.020 per rating point (14 rating = 1%, 7.133 per %), spell_haste=-6.383 ± 0.518

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (541.0 DPS) | yes | Mooncloth Circlet (14140, +0.00 DPS) [crafted]; Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 19.9 healing_power points (4.77 DPS) | yes | Lady Alizabeth's Pendant (13002, -0.50 DPS) [world_drop]; Horizon Choker (13085, -0.73 DPS) [world_drop]; Necklace of Calisea (1714, -1.33 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (541.0 DPS) | yes | Mooncloth Shoulders (14139, +0.00 DPS) [crafted]; Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp] |
| back | Wingveil Cloak (10802) | Sunken Temple: Spawn of Hakkar [dungeon] | 18.1 healing_power points (4.33 DPS) | yes | Glowing Thresher Cape (6901, -0.38 DPS) [dungeon]; Cloak of Blight (6832, -0.63 DPS) [quest]; Indomitable Cloak (14683, -0.63 DPS) [world_drop] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 92.0 healing_power points (22.04 DPS) | yes | Warlord's Satin Tunic (231632, -3.85 DPS) [vendor]; Legionnaire's Satin Tunic (227124, -5.37 DPS) [pvp]; Stormcloth Vest (10020, -5.87 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (541.0 DPS) | yes | Nethergeld Cuffs (254061, +0.00 DPS) [crafted]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted]; General's Satin Bracers (17619, -0.35 DPS) [pvp] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (541.0 DPS) | yes | Blood Guard's Satin Handwraps (227126, +0.00 DPS) [pvp]; Gilded Gloves (254095, +0.00 DPS) [crafted]; General's Satin Handwraps (231633, -0.86 DPS) [vendor] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 53.8 healing_power points (12.89 DPS) | yes | Earthenweave Cord (254077, -3.53 DPS) [crafted]; Devout Belt (16696, -3.55 DPS) [dungeon]; General's Satin Cinch (17621, -3.85 DPS) [pvp] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 64.3 healing_power points (15.41 DPS) | yes | Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Earthenweave Leggings (254125, -1.99 DPS) [crafted] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 51.6 healing_power points (12.35 DPS) | yes | General's Satin Walkers (231630, -1.67 DPS) [vendor]; Swarmtender's Footpads (275609, -2.05 DPS) [crafted]; Devout Sandals (16691, -2.80 DPS) [dungeon] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 21.7 healing_power points (5.20 DPS) | yes | Electrocutioner Lagnut (9447, -1.77 DPS) [dungeon]; Runemender's Seal (273031, -1.77 DPS) [world]; Chemist's Ring (15702, -1.99 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 19.5 healing_power points (4.66 DPS) | yes | Electrocutioner Lagnut (9447, -1.23 DPS) [dungeon]; Runemender's Seal (273031, -1.23 DPS) [world]; Chemist's Ring (15702, -1.45 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (541.0 DPS) | yes | Ankh of Life (1713, -0.90 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.46 DPS) [quest]; Alchemist's Stone (13503, -3.17 DPS) [crafted] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (541.0 DPS) | yes | Ankh of Life (1713, -0.30 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.87 DPS) [quest]; Alchemist's Stone (13503, -2.58 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -13.20 DPS) [dungeon]; Hand of Edward the Odd (2243, -17.40 DPS) [world_drop]; Hand of Righteousness (7721, -19.71 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 30.1 healing_power points (7.20 DPS) | yes | Captain Rackmore's Tiller (16789, -4.91 DPS) [quest]; Jaina's Firestarter (13064, -5.11 DPS) [world_drop]; Wand of Mana Concentration (273043, -5.14 DPS) [world] |

**New at 60:** head: Devout Crown; neck: Glowing Eye of Mordresh; shoulder: Devout Mantle; back: Wingveil Cloak; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Uther's Strength; trinket2: Shard of the Splithooves; ranged: Torch of Light

No-known-source sample (15 of 1193, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

