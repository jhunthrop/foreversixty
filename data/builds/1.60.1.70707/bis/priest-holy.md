# Leveling BiS: Holy

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-03503000000000000-00000000000000000)

Set DPS (verified): 37.4. Weights run: 9.9s. Verify run: 27.1s. 163 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.004, intellect=1.464 ± 0.006, spirit=0.700 ± 0.004, mp5=1.665 ± 0.017, crit=0.064 ± 0.003 per rating point (14 rating = 1%, 0.897 per %), spell_haste=not significant (0.004 ± 0.019)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.52 DPS) | yes | Shadow Goggles (4373, -0.84 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.29 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 2.8 healing_power points (0.24 DPS) | yes | Scholarly Pendant (277203, -0.08 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.2 healing_power points (1.11 DPS) | yes | Slime-encrusted Pads (6461, -0.69 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -0.80 DPS, sim-verified) [crafted] |
| back | Sanguine Cape (14376) | World drop [world_drop] | 5.9 healing_power points (0.50 DPS) | yes | Pearl-clasped Cloak (5542, +0.00 DPS, sim-verified) [crafted]; Seer's Cape (6378, -0.13 DPS) [dungeon]; Forest Cloak (4710, -0.25 DPS) [world_drop] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 22.4 healing_power points (1.90 DPS) | yes | Robe of the Moccasin (6465, -0.54 DPS) [dungeon]; Corsair's Overshirt (5202, -0.63 DPS, sim-verified) [dungeon]; Bloody Apron (6226, -0.80 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 12.3 healing_power points (1.04 DPS) | yes | Repurposed Hair Band (281256, -0.68 DPS) [quest]; Mystic's Bracelets (14366, -0.79 DPS) [world_drop]; Bright Bracers (3647, -1.20 DPS, sim-verified) [world_drop] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 15.4 healing_power points (1.30 DPS) | yes | Magefist Gloves (12977, -0.49 DPS, sim-verified) [world_drop]; Tomb Robber's Gloves (280096, -0.56 DPS) [quest]; Bright Gloves (3066, -0.57 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.9 healing_power points (1.43 DPS) | yes | Keller's Girdle (2911, -0.44 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.48 DPS, sim-verified) [crafted]; Wise Man's Belt (4786, -0.81 DPS) [vendor] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 29.6 healing_power points (2.50 DPS) | yes | Filigreed Silky Leggings (253939, -1.52 DPS) [crafted]; Filigreed Shadow Leggings (253943, -1.52 DPS) [crafted]; Darkweave Breeches (12987, -1.78 DPS, sim-verified) [world_drop] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 13.4 healing_power points (1.13 DPS) | yes | Kimbra Boots (6191, -0.51 DPS, sim-verified) [quest]; Spidersilk Boots (4320, -0.64 DPS) [crafted]; Sanguine Sandals (14374, -0.64 DPS) [world_drop] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 8.8 healing_power points (0.74 DPS) | yes | Band of Purification (12996, -0.39 DPS) [world_drop]; Deep Fathom Ring (6463, -0.45 DPS) [dungeon]; Minor Channeling Ring (1449, -0.50 DPS) [quest] |
| finger2 | Volcanic Rock Ring (12053) | World drop [world_drop] | 4.4 healing_power points (0.37 DPS) | yes | Band of Purification (12996, +0.00 DPS, sim-verified) [world_drop]; Deep Fathom Ring (6463, -0.08 DPS) [dungeon]; Minor Channeling Ring (1449, -0.12 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 14.6 healing_power points (1.24 DPS) | yes | Staff of Westfall (2042, -0.33 DPS, sim-verified) [quest]; Lesser Staff of the Spire (1300, -0.50 DPS) [world]; Witching Stave (1484, -0.62 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Flaring Baton (5326) (or Moonstone Wand (15204)) | The Escape [quest] | 2.9 healing_power points (0.25 DPS) | yes | Moonstone Wand (15204, +0.00 DPS) [quest]; Sable Wand (7607, -0.07 DPS) [quest]; Elven Wand (5604, -0.12 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Tarnished Locket; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Volcanic Rock Ring; main_hand: Twisted Chanter's Staff; ranged: Flaring Baton

No-known-source sample (15 of 163, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes; 16607 Acolyte's Sacrificial Robes; 18851 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (gnome, 000000000000000000-03505003030110000-00000000000000000)

Set DPS (verified): 58.9. Weights run: 10.6s. Verify run: 36.5s. 294 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.017, intellect=1.646 ± 0.010, spirit=0.718 ± 0.006, mp5=1.931 ± 0.014, crit=0.091 ± 0.005 per rating point (14 rating = 1%, 1.270 per %), spell_haste=0.331 ± 0.056

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 37.3 healing_power points (3.18 DPS) | yes | Filigreed Pristine Circlet (253975, -1.31 DPS) [crafted]; Resilient Cap (14401, -1.39 DPS) [world_drop]; Nightsky Cowl (4039, -1.74 DPS, sim-verified) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 9.9 healing_power points (0.84 DPS) | yes | Crystal Starfire Medallion (5003, -0.10 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.10 DPS, sim-verified) [world_drop]; Pendant of Myzrael (4614, -0.47 DPS) [dungeon] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 18.1 healing_power points (1.54 DPS) | yes | Mantle of Honor (3560, -0.13 DPS) [quest]; Nightsky Mantle (4718, -0.27 DPS, sim-verified) [world_drop]; Violet Sorcerer's Mantle (273051, -0.28 DPS) [world] |
| back | Prelacy Cape (7004) | Researching the Corruption [quest] | 15.3 healing_power points (1.31 DPS) | yes | Repairman's Cape (9605, -0.24 DPS) [quest]; Glowing Thresher Cape (6901, -0.36 DPS, sim-verified) [dungeon]; Resilient Cape (14400, -0.46 DPS) [world_drop] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 38.1 healing_power points (3.25 DPS) | yes | Pristine Gown (253961, -0.48 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -1.08 DPS) [dungeon]; Filigreed Pristine Gown (253901, -1.26 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 14.0 healing_power points (1.20 DPS) | yes | Glowing Magical Bracelets (13106, -0.09 DPS, sim-verified) [world_drop]; Nightsky Wristbands (6407, -0.17 DPS) [world_drop]; Stonecloth Bindings (14416, -0.49 DPS) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 19.9 healing_power points (1.70 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.27 DPS) [dungeon]; Pristine Gloves (253913, -0.34 DPS) [crafted]; Gloves of Old (9395, -0.44 DPS, sim-verified) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 17.6 healing_power points (1.50 DPS) | yes | Keller's Girdle (2911, -0.38 DPS) [world_drop]; Nightsky Sash (4720, -0.38 DPS) [world_drop]; Wizard's Belt (4827, -0.56 DPS, sim-verified) [vendor] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 37.1 healing_power points (3.17 DPS) | yes | Filigreed Pristine Leggings (253937, -0.67 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -1.21 DPS) [crafted]; Necromancer Leggings (2277, -1.62 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 34.4 healing_power points (2.93 DPS) | yes | Nimbus Boots (6998, -1.51 DPS, sim-verified) [quest]; Acidic Walkers (9454, -1.57 DPS) [dungeon]; Frothing Slippers (254003, -1.71 DPS) [crafted] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 11.0 healing_power points (0.94 DPS) | yes | Lavishly Jeweled Ring (1156, -0.10 DPS) [dungeon]; The Queen's Jewel (13094, -0.17 DPS) [world_drop]; Defias Mage Ring (1077, -0.20 DPS) [world] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 11.0 healing_power points (0.94 DPS) | yes | Lavishly Jeweled Ring (1156, +0.00 DPS, sim-verified) [dungeon]; The Queen's Jewel (13094, -0.17 DPS) [world_drop]; Defias Mage Ring (1077, -0.20 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-verified (58.9 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -0.81 DPS) [quest]; Death Speaker Scepter (2816, -2.16 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 10.4 healing_power points (0.89 DPS) | yes | Lesser Mystic Wand (11289, -0.33 DPS) [crafted]; Starfaller (13063, -0.33 DPS) [world_drop]; Summoner's Wand (5245, -0.39 DPS, sim-verified) [world] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Death Speaker Mantle; back: Prelacy Cape; chest: Death Speaker Robes; hands: Truefaith Gloves; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Spine of the Devourer; ranged: Wand of Mana Concentration

No-known-source sample (15 of 294, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 40 (gnome, 000000000000000000-03505003030121431-00000000000000000)

Set DPS (verified): 108.7. Weights run: 11.0s. Verify run: 37.1s. 396 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.008, intellect=1.580 ± 0.021, spirit=1.179 ± 0.023, mp5=2.623 ± 0.025, crit=0.142 ± 0.007 per rating point (14 rating = 1%, 1.994 per %), spell_haste=-5.749 ± 0.401

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 51.5 healing_power points (4.95 DPS) | yes | Whitemane's Chapeau (7720, -0.82 DPS) [dungeon]; Holy Shroud (2721, -1.10 DPS) [world_drop]; Earthen Silk Hood (254015, -1.70 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 25.6 healing_power points (2.47 DPS) | yes | Necklace of Calisea (1714, -0.61 DPS) [dungeon]; Amberglow Talisman (10824, -1.33 DPS) [quest]; Kaleidoscope Chain (13084, -1.40 DPS) [world_drop] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 35.5 healing_power points (3.42 DPS) | yes | Earthen Silk Shoulders (254033, -0.78 DPS) [crafted]; Mistscape Mantle (4734, -1.18 DPS) [dungeon]; Silksand Shoulder Pads (14423, -1.41 DPS) [world_drop] |
| back | Blackforge Cape (6424) | Uldaman: Ancient Treasure [dungeon] | sim-verified (108.7 DPS) | yes | Prelacy Cape (7004, +0.00 DPS) [quest]; Sergeant Major's Cape (16336, -0.08 DPS) [pvp]; Glowing Thresher Cape (6901, -0.09 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 67.4 healing_power points (6.49 DPS) | yes | Death Speaker Robes (6682, -2.89 DPS) [dungeon]; Pristine Gown (253961, -3.05 DPS) [crafted]; Doomsayer's Robe (4746, -3.48 DPS) [quest] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | sim-verified (108.7 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS) [dungeon]; Earthen Silk Cuffs (254019, +0.00 DPS) [crafted]; Nightsky Wristbands (6407, -0.12 DPS) [world_drop] |
| hands | Gilded Handwraps (254021) | Tailoring [crafted] | 33.1 healing_power points (3.18 DPS) | yes | Stormcloth Gloves (10011, -0.30 DPS) [crafted]; Earthen Silk Gloves (254017, -0.54 DPS) [crafted]; Gloves of Old (9395, -1.06 DPS) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 41.7 healing_power points (4.01 DPS) | yes | Deathmage Sash (10771, -1.73 DPS) [dungeon]; Windchaser Cinch (14435, -1.78 DPS) [world_drop]; Sutarn's Ring (13105, -1.81 DPS) [world_drop] |
| legs | Stoneweaver Leggings (9407) | Uldaman: Ironaya [dungeon] | 39.8 healing_power points (3.83 DPS) | yes | Pristine Leggings (253987, -0.08 DPS) [crafted]; Stormcloth Pants (10010, -0.68 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.73 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 35.8 healing_power points (3.44 DPS) | yes | Furen's Boots (13100, -0.56 DPS) [world_drop]; Thoughtcast Boots (10578, -1.20 DPS) [dungeon]; Kodo Rustler Boots (15697, -1.66 DPS) [quest] |
| finger1 | Welken Ring (5011) | World drop [world_drop] | 16.1 healing_power points (1.55 DPS) | yes | Ogremind Ring (1993, -0.15 DPS) [world_drop]; Electrocutioner Lagnut (9447, -0.18 DPS) [dungeon]; Runemender's Seal (273031, -0.18 DPS) [world] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 14.6 healing_power points (1.40 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Electrocutioner Lagnut (9447, -0.03 DPS) [dungeon]; Runemender's Seal (273031, -0.03 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, -3.41 DPS) [dungeon]; Illusionary Rod (7713, -4.01 DPS) [dungeon]; Hypnotic Blade (7714, -4.16 DPS) [dungeon] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 27.8 healing_power points (2.67 DPS) | yes | Orb of Souls (249395, -0.74 DPS) [crafted]; Beacon of Hope (9393, -0.87 DPS) [dungeon]; Aurora Sphere (7610, -1.16 DPS) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 13.0 healing_power points (1.25 DPS) | yes | Wand of Mana Concentration (273043, -0.14 DPS) [world]; Flash Wand (5248, -0.30 DPS) [quest]; Goblin Igniter (5253, -0.30 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Blackforge Cape; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Gilded Handwraps; waist: Gilded Cord; legs: Stoneweaver Leggings; finger1: Welken Ring; finger2: Voodoo Band; trinket1: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Orb of Lorica; ranged: Jaina's Firestarter

No-known-source sample (15 of 396, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 50 (gnome, 025003000000000000-03505003030121431-00000000000000000)

Set DPS (verified): 203.7. Weights run: 10.8s. Verify run: 38.0s. 515 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.052, intellect=1.516 ± 0.036, spirit=1.815 ± 0.034, mp5=3.535 ± 0.037, crit=0.265 ± 0.014 per rating point (14 rating = 1%, 3.706 per %), spell_haste=-6.355 ± 0.548

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 67.4 healing_power points (7.09 DPS) | yes | Knight-Lieutenant's Satin Cover (220896, -0.11 DPS) [vendor]; Cassandra's Grace (13102, -0.75 DPS) [world_drop]; Stormcloth Headband (10032, -1.06 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 28.5 healing_power points (2.99 DPS) | yes | Glowing Eye of Mordresh (10769, -0.35 DPS) [dungeon]; Necklace of Calisea (1714, -0.54 DPS) [dungeon]; Mindburst Medallion (11196, -1.08 DPS) [quest] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | 49.0 healing_power points (5.15 DPS) | yes | Knight-Lieutenant's Satin Pads (220894, +0.00 DPS) [vendor]; Inquisitor's Shawl (19507, -1.50 DPS) [dungeon]; Hibernal Mantle (8111, -1.65 DPS) [world_drop] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | sim-verified (203.9 DPS) | yes | Wingveil Cloak (10802, +0.00 DPS) [dungeon]; Chainlink Towel (9648, -0.13 DPS) [quest]; Prelacy Cape (7004, -0.22 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 72.5 healing_power points (7.62 DPS) | yes | Ghostweave Vest (14141, -0.83 DPS) [crafted]; Knight's Satin Armor (220892, -1.55 DPS) [vendor]; Grimlok's Tribal Vestments (9415, -2.16 DPS) [dungeon] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 39.7 healing_power points (4.17 DPS) | yes | Shizzle's Nozzle Wiper (11917, -1.69 DPS) [quest]; Venomshroud Armguards (14439, -2.01 DPS) [world_drop]; Netherfroth Cuffs (254063, -2.10 DPS) [crafted] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 58.7 healing_power points (6.17 DPS) | yes | Earthenweave Gloves (254075, -1.34 DPS) [crafted]; Sergeant Major's Satin Gloves (220897, -2.27 DPS) [vendor]; Quintis' Research Gloves (11888, -2.51 DPS) [quest] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | sim-verified (203.9 DPS) | yes | Gilded Cord (254037, +0.00 DPS) [crafted]; Earthenweave Cord (254077, +0.00 DPS) [crafted]; Gilded Waistcord (254081, +0.00 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | sim-verified (203.9 DPS) | yes | Senior Designer's Pantaloons (11841, +0.00 DPS) [dungeon]; Dalewind Trousers (13008, -0.73 DPS) [world_drop]; Knight's Satin Leggings (220893, -0.86 DPS) [vendor] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | sim-verified (203.9 DPS) | yes | Earthenweave Boots (254093, +0.00 DPS) [crafted]; Vinerot Sandals (17748, -0.09 DPS) [dungeon]; Furen's Boots (13100, -0.53 DPS) [world_drop] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 30.9 healing_power points (3.24 DPS) | yes | Choking Band (11868, -0.77 DPS) [quest]; Coldwater Ring (4550, -1.05 DPS) [quest]; Electrocutioner Lagnut (9447, -1.28 DPS) [dungeon] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | sim-verified (203.9 DPS) | yes | Coldwater Ring (4550, +0.00 DPS) [quest]; Choking Band (11868, +0.00 DPS) [quest]; Electrocutioner Lagnut (9447, -0.17 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (203.9 DPS) | yes | Evonice's Landin' Pilla (18951, -0.38 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.76 DPS) [quest]; Alchemist's Stone (13503, -1.53 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (203.9 DPS) | yes | Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.37 DPS) [quest]; Alchemist's Stone (13503, -1.13 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -7.41 DPS) [dungeon]; Hand of Righteousness (7721, -10.69 DPS) [dungeon]; Spire of Hakkar (10844, -12.53 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cairnstone Sliver (9654) | The Morrow Stone [quest] | 16.7 healing_power points (1.75 DPS) | yes | Jaina's Firestarter (13064, -0.22 DPS) [world_drop]; Wand of Mana Concentration (273043, -0.37 DPS) [world]; Goblin Igniter (5253, -0.54 DPS) [quest] |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Nethergeld Shoulders; back: Imperial Red Cloak; wrist: Nethergeld Cuffs; hands: Gilded Gloves; waist: Dawnspire Cord; legs: Kilt of the Atal'ai Prophet; feet: Gilded Sandals; finger1: Eye of Adaegus; finger2: Welken Ring; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Cairnstone Sliver

No-known-source sample (15 of 515, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60 (gnome, 025003031330000000-03505003030121431-00000000000000000)

Set DPS (verified): 337.4. Weights run: 9.2s. Verify run: 65.8s. 1200 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.112, intellect=0.958 ± 0.028, spirit=1.169 ± 0.027, mp5=2.088 ± 0.037, crit=0.184 ± 0.014 per rating point (14 rating = 1%, 2.580 per %), spell_haste=-16.256 ± 0.701

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Circlet (14140, +0.00 DPS) [crafted]; Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 20.7 healing_power points (4.45 DPS) | yes | Horizon Choker (13085, -0.55 DPS) [world_drop]; Lady Alizabeth's Pendant (13002, -0.67 DPS) [world_drop]; Necklace of Calisea (1714, -1.24 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Shoulders (14139, +0.00 DPS) [crafted]; Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp] |
| back | Sergeant Major's Cape (16337) | PvP rank 9 · Sergeant Major · Alliance [pvp] | sim-verified (337.4 DPS) | yes | Prelacy Cape (7004, +0.00 DPS) [quest]; Wingveil Cloak (10802, +0.00 DPS) [dungeon]; Ethereal Mist Cape (11873, -0.02 DPS) [quest] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 91.7 healing_power points (19.73 DPS) | yes | Field Marshal's Satin Tunic (231624, -2.88 DPS) [vendor]; Knight-Captain's Satin Tunic (227122, -4.51 DPS) [pvp]; Mooncloth Vest (14138, -5.24 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nethergeld Cuffs (254061, +0.00 DPS) [crafted]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted]; Marshal's Satin Bracers (17606, -0.10 DPS) [pvp] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Lieutenant's Satin Handwraps (227128, +0.00 DPS) [pvp]; Gilded Gloves (254095, +0.00 DPS) [crafted]; Marshal's Satin Handwraps (231623, -0.65 DPS) [vendor] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 54.9 healing_power points (11.82 DPS) | yes | Devout Belt (16696, -3.07 DPS) [dungeon]; Marshal's Satin Sash (17609, -3.31 DPS) [pvp]; Earthenweave Cord (254077, -3.46 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 65.3 healing_power points (14.04 DPS) | yes | Knight-Captain's Satin Legguards (227125, +0.00 DPS) [pvp]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Earthenweave Leggings (254125, -2.06 DPS) [crafted] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 52.6 healing_power points (11.32 DPS) | yes | Marshal's Satin Walkers (231627, -1.37 DPS) [vendor]; Swarmtender's Footpads (275609, -2.07 DPS) [crafted]; Devout Sandals (16691, -2.62 DPS) [dungeon] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 22.4 healing_power points (4.83 DPS) | yes | Choking Band (11868, -1.56 DPS) [quest]; Electrocutioner Lagnut (9447, -1.78 DPS) [dungeon]; Runemender's Seal (273031, -1.78 DPS) [world] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 19.8 healing_power points (4.26 DPS) | yes | Choking Band (11868, -0.99 DPS) [quest]; Electrocutioner Lagnut (9447, -1.20 DPS) [dungeon]; Runemender's Seal (273031, -1.20 DPS) [world] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.18 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.68 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -0.85 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.36 DPS) [quest]; Thunderbrew's Boot Flask (744, -1.86 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -12.39 DPS) [dungeon]; Hand of Edward the Odd (2243, -15.96 DPS) [world_drop]; Hand of Righteousness (7721, -18.28 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 29.9 healing_power points (6.43 DPS) | yes | Cairnstone Sliver (9654, -4.14 DPS) [quest]; Captain Rackmore's Tiller (16789, -4.38 DPS) [quest]; Jaina's Firestarter (13064, -4.43 DPS) [world_drop] |

**New at 60:** head: Devout Crown; neck: Glowing Eye of Mordresh; shoulder: Devout Mantle; back: Sergeant Major's Cape; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Shard of the Splithooves; ranged: Torch of Light

No-known-source sample (15 of 1200, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60, raid preset (gnome, 025003031330000000-03505003030121431-00000000000000000)

Set DPS (verified): 771.4. Weights run: 5.9s. Verify run: 30.7s. 1200 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.173, intellect=1.011 ± 0.033, spirit=0.924 ± 0.031, mp5=1.887 ± 0.034, crit=0.379 ± 0.020 per rating point (14 rating = 1%, 5.311 per %), spell_haste=-3.552 ± 0.630

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Circlet (14140, +0.00 DPS) [crafted]; Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 21.1 healing_power points (6.22 DPS) | yes | Horizon Choker (13085, -0.96 DPS) [world_drop]; Lady Alizabeth's Pendant (13002, -2.13 DPS) [world_drop]; Necklace of Calisea (1714, -2.22 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Shoulders (14139, +0.00 DPS) [crafted]; Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp] |
| back | Prelacy Cape (7004) | Researching the Corruption [quest] | 16.5 healing_power points (4.88 DPS) | yes | Sergeant Major's Cape (16337, -0.31 DPS) [pvp]; Wingveil Cloak (10802, -0.41 DPS) [dungeon]; Imperial Red Cloak (8248, -0.51 DPS) [world_drop] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 87.8 healing_power points (25.88 DPS) | yes | Field Marshal's Satin Tunic (231624, -3.00 DPS) [vendor]; Knight-Captain's Satin Tunic (227122, -5.81 DPS) [pvp]; Stormcloth Vest (10020, -6.60 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Satin Bracers (17606, +0.00 DPS) [pvp]; Nethergeld Cuffs (254061, +0.00 DPS) [crafted]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Lieutenant's Satin Handwraps (227128, +0.00 DPS) [pvp]; Marshal's Satin Handwraps (231623, +0.00 DPS) [vendor]; Gilded Gloves (254095, +0.00 DPS) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 53.5 healing_power points (15.76 DPS) | yes | Devout Belt (16696, -4.11 DPS) [dungeon]; Gilded Waistcord (254081, -4.53 DPS) [crafted]; Marshal's Satin Sash (17609, -4.82 DPS) [pvp] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 60.4 healing_power points (17.81 DPS) | yes | Knight-Captain's Satin Legguards (227125, +0.00 DPS) [pvp]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Earthenweave Leggings (254125, -2.54 DPS) [crafted] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 50.2 healing_power points (14.79 DPS) | yes | Marshal's Satin Walkers (231627, -0.87 DPS) [vendor]; Swarmtender's Footpads (275609, -2.11 DPS) [crafted]; Knight-Lieutenant's Satin Walkers (227129, -2.65 DPS) [pvp] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 20.3 healing_power points (5.98 DPS) | yes | Maiden's Circle (13001, -2.26 DPS) [world_drop]; Electrocutioner Lagnut (9447, -2.30 DPS) [dungeon]; Runemender's Seal (273031, -2.30 DPS) [world] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 17.1 healing_power points (5.06 DPS) | yes | Maiden's Circle (13001, -1.34 DPS) [world_drop]; Electrocutioner Lagnut (9447, -1.38 DPS) [dungeon]; Runemender's Seal (273031, -1.38 DPS) [world] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (771.6 DPS) | yes | Ankh of Life (1713, -0.07 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.62 DPS) [quest]; Thunderbrew's Boot Flask (744, -1.16 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (771.6 DPS) | yes | Ankh of Life (1713, -2.04 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -2.58 DPS) [quest]; Thunderbrew's Boot Flask (744, -3.13 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -16.72 DPS) [dungeon]; Hand of Edward the Odd (2243, -22.18 DPS) [world_drop]; Hand of Righteousness (7721, -25.42 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 27.2 healing_power points (8.01 DPS) | yes | Cairnstone Sliver (9654, -5.16 DPS) [quest]; Jaina's Firestarter (13064, -5.40 DPS) [world_drop]; Captain Rackmore's Tiller (16789, -5.42 DPS) [quest] |

**New at 60:** head: Devout Crown; neck: Glowing Eye of Mordresh; shoulder: Devout Mantle; back: Prelacy Cape; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Shard of the Splithooves; ranged: Torch of Light

No-known-source sample (15 of 1200, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

## Horde

### Band 20 (troll, 000000000000000000-03503000000000000-00000000000000000)

Set DPS (verified): 35.5. Weights run: 9.9s. Verify run: 27.3s. 156 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.004, intellect=1.464 ± 0.006, spirit=0.700 ± 0.004, mp5=1.665 ± 0.017, crit=0.064 ± 0.003 per rating point (14 rating = 1%, 0.897 per %), spell_haste=not significant (0.004 ± 0.019)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.52 DPS) | yes | Shadow Goggles (4373, -0.73 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.29 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 2.1 healing_power points (0.18 DPS) | yes | Scholarly Pendant (277203, -0.16 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.2 healing_power points (1.11 DPS) | yes | Slime-encrusted Pads (6461, -0.69 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -0.89 DPS, sim-verified) [crafted] |
| back | Sanguine Cape (14376) | World drop [world_drop] | 5.9 healing_power points (0.50 DPS) | yes | Seer's Cape (6378, -0.13 DPS) [dungeon]; Pearl-clasped Cloak (5542, -0.22 DPS, sim-verified) [crafted]; Forest Cloak (4710, -0.25 DPS) [world_drop] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 22.4 healing_power points (1.90 DPS) | yes | Robe of the Moccasin (6465, -0.54 DPS) [dungeon]; Corsair's Overshirt (5202, -0.67 DPS, sim-verified) [dungeon]; Bloody Apron (6226, -0.80 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 12.3 healing_power points (1.04 DPS) | yes | Bright Bracers (3647, -0.55 DPS) [world_drop]; Crystalline Cuffs (14148, -0.62 DPS) [dungeon]; Featherbead Bracers (15452, -0.85 DPS, sim-verified) [quest] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 15.4 healing_power points (1.30 DPS) | yes | Blight Gloves (279877, -0.16 DPS, sim-verified) [quest]; Magefist Gloves (12977, -0.45 DPS) [world_drop]; Bright Gloves (3066, -0.57 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.9 healing_power points (1.43 DPS) | yes | Keller's Girdle (2911, -0.44 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.51 DPS, sim-verified) [crafted]; Wise Man's Belt (4786, -0.81 DPS) [vendor] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 29.6 healing_power points (2.50 DPS) | yes | Filigreed Silky Leggings (253939, -1.52 DPS) [crafted]; Filigreed Shadow Leggings (253943, -1.52 DPS) [crafted]; Darkweave Breeches (12987, -1.65 DPS, sim-verified) [world_drop] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 13.4 healing_power points (1.13 DPS) | yes | Spidersilk Boots (4320, -0.64 DPS) [crafted]; Bluegill Sandals (1560, -0.65 DPS) [world]; Sanguine Sandals (14374, -0.93 DPS, sim-verified) [world_drop] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 8.8 healing_power points (0.74 DPS) | yes | Volcanic Rock Ring (12053, -0.37 DPS) [world_drop]; Band of Purification (12996, -0.39 DPS) [world_drop]; Deep Fathom Ring (6463, -0.45 DPS) [dungeon] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 7.3 healing_power points (0.62 DPS) | yes | Band of Purification (12996, -0.26 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.27 DPS, sim-verified) [world_drop]; Deep Fathom Ring (6463, -0.32 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 14.6 healing_power points (1.24 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Staff of Orgrimmar (15444, -0.17 DPS) [quest]; Lesser Staff of the Spire (1300, -0.50 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Decay (5252) (or Flaring Baton (5326)) | Beren's Peril [quest] | 2.9 healing_power points (0.25 DPS) | yes | Flaring Baton (5326, +0.00 DPS) [quest]; Wisesight Wand (286750, -0.12 DPS) [world] |

**New at 20:** head: Pristine Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Gnarled Necromancer's Staff; ranged: Wand of Decay

No-known-source sample (15 of 156, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (troll, 000000000000000000-03505003030110000-00000000000000000)

Set DPS (verified): 56.0. Weights run: 10.6s. Verify run: 37.0s. 286 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.017, intellect=1.646 ± 0.010, spirit=0.718 ± 0.006, mp5=1.931 ± 0.014, crit=0.091 ± 0.005 per rating point (14 rating = 1%, 1.270 per %), spell_haste=0.331 ± 0.056

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 37.3 healing_power points (3.18 DPS) | yes | Filigreed Pristine Circlet (253975, -1.31 DPS) [crafted]; Resilient Cap (14401, -1.39 DPS) [world_drop]; Nightsky Cowl (4039, -1.89 DPS, sim-verified) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 9.9 healing_power points (0.84 DPS) | yes | Crystal Starfire Medallion (5003, -0.10 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.16 DPS, sim-verified) [world_drop]; Pendant of Myzrael (4614, -0.47 DPS) [dungeon] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 18.1 healing_power points (1.54 DPS) | yes | Ghostly Mantle (3324, -0.23 DPS) [quest]; Violet Sorcerer's Mantle (273051, -0.28 DPS) [world]; Nightsky Mantle (4718, -0.48 DPS, sim-verified) [world_drop] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 12.7 healing_power points (1.09 DPS) | yes | Resilient Cape (14400, -0.35 DPS, sim-verified) [world_drop]; Soft Willow Cape (16661, -0.39 DPS) [quest]; Silver-thread Cloak (4713, -0.40 DPS) [world_drop] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 38.1 healing_power points (3.25 DPS) | yes | Pristine Gown (253961, -0.57 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -1.08 DPS) [dungeon]; Filigreed Pristine Gown (253901, -1.26 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 14.0 healing_power points (1.20 DPS) | yes | Nightsky Wristbands (6407, -0.17 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.19 DPS, sim-verified) [world_drop]; Featherbead Bracers (15452, -0.49 DPS) [quest] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 19.9 healing_power points (1.70 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.27 DPS) [dungeon]; Pristine Gloves (253913, -0.34 DPS) [crafted]; Gloves of Old (9395, -0.57 DPS, sim-verified) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 17.6 healing_power points (1.50 DPS) | yes | Nightsky Sash (4720, -0.38 DPS) [world_drop]; Wizard's Belt (4827, -0.38 DPS) [vendor]; Lilac Sash (6780, -0.43 DPS, sim-verified) [quest] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 37.1 healing_power points (3.17 DPS) | yes | Filigreed Pristine Leggings (253937, -0.56 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -1.21 DPS) [crafted]; Necromancer Leggings (2277, -1.62 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 34.4 healing_power points (2.93 DPS) | yes | Frothing Slippers (254003, -1.71 DPS) [crafted]; Black Slippers (254007, -1.71 DPS) [crafted]; Acidic Walkers (9454, -1.82 DPS, sim-verified) [dungeon] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 11.0 healing_power points (0.94 DPS) | yes | Solid Gold Band (279845, -0.04 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.10 DPS) [dungeon]; The Queen's Jewel (13094, -0.17 DPS) [world_drop] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 11.0 healing_power points (0.94 DPS) | yes | Solid Gold Band (279845, -0.04 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.10 DPS) [dungeon]; The Queen's Jewel (13094, -0.17 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-verified (56.0 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Glimmering Staff (249392, -1.53 DPS) [crafted]; Death Speaker Scepter (2816, -2.11 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 10.4 healing_power points (0.89 DPS) | yes | Summoner's Wand (5245, -0.32 DPS, sim-verified) [world]; Lesser Mystic Wand (11289, -0.33 DPS) [crafted]; Starfaller (13063, -0.33 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Death Speaker Mantle; back: Glowing Thresher Cape; chest: Death Speaker Robes; hands: Truefaith Gloves; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Spine of the Devourer; ranged: Wand of Mana Concentration

No-known-source sample (15 of 286, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (troll, 000000000000000000-03505003030121431-00000000000000000)

Set DPS (verified): 102.5. Weights run: 11.0s. Verify run: 37.5s. 388 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.008, intellect=1.580 ± 0.021, spirit=1.179 ± 0.023, mp5=2.623 ± 0.025, crit=0.142 ± 0.007 per rating point (14 rating = 1%, 1.994 per %), spell_haste=-5.749 ± 0.401

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 51.5 healing_power points (4.95 DPS) | yes | Holy Shroud (2721, -1.10 DPS) [world_drop]; Earthen Silk Hood (254015, -1.70 DPS) [crafted]; Whitemane's Chapeau (7720, -1.85 DPS, sim-verified) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 25.6 healing_power points (2.47 DPS) | yes | Necklace of Calisea (1714, -0.71 DPS, sim-verified) [dungeon]; Amberglow Talisman (10824, -1.33 DPS) [quest]; Kaleidoscope Chain (13084, -1.40 DPS) [world_drop] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 35.5 healing_power points (3.42 DPS) | yes | Mistscape Mantle (4734, -1.18 DPS) [dungeon]; Earthen Silk Shoulders (254033, -1.39 DPS, sim-verified) [crafted]; Silksand Shoulder Pads (14423, -1.41 DPS) [world_drop] |
| back | Blackforge Cape (6424) | Uldaman: Ancient Treasure [dungeon] | 17.4 healing_power points (1.67 DPS) | yes | Cloak of Blight (6832, -0.19 DPS) [quest]; Windchaser Cloak (14430, -0.30 DPS) [world_drop]; Glowing Thresher Cape (6901, -0.76 DPS, sim-verified) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 67.4 healing_power points (6.49 DPS) | yes | Pristine Gown (253961, -3.05 DPS) [crafted]; Doomsayer's Robe (4746, -3.48 DPS) [quest]; Death Speaker Robes (6682, -4.00 DPS, sim-verified) [dungeon] |
| wrist | Earthen Silk Cuffs (254019) | Tailoring [crafted] | 15.9 healing_power points (1.53 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Windchaser Cuffs (14429, -0.16 DPS) [world_drop]; Nightsky Wristbands (6407, -0.28 DPS) [world_drop] |
| hands | Gilded Handwraps (254021) | Tailoring [crafted] | 33.1 healing_power points (3.18 DPS) | yes | Stormcloth Gloves (10011, -0.33 DPS, sim-verified) [crafted]; Earthen Silk Gloves (254017, -0.54 DPS) [crafted]; Gloves of Old (9395, -1.06 DPS) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 41.7 healing_power points (4.01 DPS) | yes | Windchaser Cinch (14435, -1.78 DPS) [world_drop]; Sutarn's Ring (13105, -1.81 DPS) [world_drop]; Deathmage Sash (10771, -2.33 DPS, sim-verified) [dungeon] |
| legs | Stoneweaver Leggings (9407) | Uldaman: Ironaya [dungeon] | 39.8 healing_power points (3.83 DPS) | yes | Pristine Leggings (253987, -0.65 DPS, sim-verified) [crafted]; Stormcloth Pants (10010, -0.68 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.73 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 35.8 healing_power points (3.44 DPS) | yes | Furen's Boots (13100, -0.72 DPS, sim-verified) [world_drop]; Thoughtcast Boots (10578, -1.20 DPS) [dungeon]; Boots of the Maharishi (9658, -1.39 DPS) [quest] |
| finger1 | Welken Ring (5011) | World drop [world_drop] | 16.1 healing_power points (1.55 DPS) | yes | Ogremind Ring (1993, -0.15 DPS) [world_drop]; Electrocutioner Lagnut (9447, -0.18 DPS) [dungeon]; Runemender's Seal (273031, -0.18 DPS) [world] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 14.6 healing_power points (1.40 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Electrocutioner Lagnut (9447, -0.03 DPS) [dungeon]; Runemender's Seal (273031, -0.03 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (102.5 DPS) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (102.5 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Death Speaker Scepter (2816, -3.41 DPS) [dungeon]; Illusionary Rod (7713, -4.01 DPS) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 20.1 healing_power points (1.93 DPS) | yes | Aurora Sphere (7610, -0.41 DPS) [dungeon]; Beacon of Hope (9393, -0.56 DPS, sim-verified) [dungeon]; Traitor's Finger (6774, -0.60 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 13.0 healing_power points (1.25 DPS) | yes | Wand of Mana Concentration (273043, -0.14 DPS) [world]; Flash Wand (5248, -0.30 DPS) [quest]; Goblin Igniter (5253, -0.30 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Blackforge Cape; chest: Stormcloth Vest; wrist: Earthen Silk Cuffs; hands: Gilded Handwraps; waist: Gilded Cord; legs: Stoneweaver Leggings; finger1: Welken Ring; finger2: Voodoo Band; trinket1: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Orb of Souls; ranged: Jaina's Firestarter

No-known-source sample (15 of 388, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16336 Sergeant Major's Cape

### Band 50 (troll, 025003000000000000-03505003030121431-00000000000000000)

Set DPS (verified): 195.5. Weights run: 10.8s. Verify run: 38.5s. 507 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.052, intellect=1.516 ± 0.036, spirit=1.815 ± 0.034, mp5=3.535 ± 0.037, crit=0.265 ± 0.014 per rating point (14 rating = 1%, 3.706 per %), spell_haste=-6.355 ± 0.548

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 67.4 healing_power points (7.09 DPS) | yes | Blood Guard's Satin Cover (220899, -0.11 DPS) [vendor]; Cassandra's Grace (13102, -0.75 DPS) [world_drop]; Stormcloth Headband (10032, -1.06 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 28.5 healing_power points (2.99 DPS) | yes | Glowing Eye of Mordresh (10769, -0.35 DPS) [dungeon]; Necklace of Calisea (1714, -0.54 DPS) [dungeon]; Choker of the High Shaman (4112, -0.67 DPS) [quest] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | 49.0 healing_power points (5.15 DPS) | yes | Blood Guard's Satin Pads (220901, +0.00 DPS) [vendor]; Inquisitor's Shawl (19507, -1.50 DPS) [dungeon]; Hibernal Mantle (8111, -1.65 DPS) [world_drop] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | sim-verified (195.5 DPS) | yes | Wingveil Cloak (10802, +0.00 DPS) [dungeon]; Cloak of Blight (6832, -0.04 DPS) [quest]; Chainlink Towel (9648, -0.13 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 72.5 healing_power points (7.62 DPS) | yes | Ghostweave Vest (14141, -0.83 DPS) [crafted]; Stone Guard's Satin Armor (220903, -1.55 DPS) [vendor]; Grimlok's Tribal Vestments (9415, -2.16 DPS) [dungeon] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 39.7 healing_power points (4.17 DPS) | yes | Shizzle's Nozzle Wiper (11917, -1.69 DPS) [quest]; Venomshroud Armguards (14439, -2.01 DPS) [world_drop]; Netherfroth Cuffs (254063, -2.10 DPS) [crafted] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 58.7 healing_power points (6.17 DPS) | yes | Earthenweave Gloves (254075, -1.34 DPS) [crafted]; Greenleaf Handwraps (19116, -2.10 DPS) [quest]; First Sergeant's Satin Gloves (220898, -2.27 DPS) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | sim-verified (195.5 DPS) | yes | Gilded Cord (254037, +0.00 DPS) [crafted]; Earthenweave Cord (254077, +0.00 DPS) [crafted]; Gilded Waistcord (254081, +0.00 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | sim-verified (195.5 DPS) | yes | Senior Designer's Pantaloons (11841, +0.00 DPS) [dungeon]; Dalewind Trousers (13008, -0.73 DPS) [world_drop]; Stone Guard's Satin Leggings (220902, -0.86 DPS) [vendor] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | sim-verified (195.5 DPS) | yes | Earthenweave Boots (254093, +0.00 DPS) [crafted]; Vinerot Sandals (17748, -0.09 DPS) [dungeon]; Furen's Boots (13100, -0.53 DPS) [world_drop] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 30.9 healing_power points (3.24 DPS) | yes | Coldwater Ring (4550, -1.05 DPS) [quest]; Electrocutioner Lagnut (9447, -1.28 DPS) [dungeon]; Runemender's Seal (273031, -1.28 DPS) [world] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | sim-verified (195.5 DPS) | yes | Coldwater Ring (4550, +0.00 DPS) [quest]; Electrocutioner Lagnut (9447, -0.17 DPS) [dungeon]; Runemender's Seal (273031, -0.17 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (195.5 DPS) | yes | Evonice's Landin' Pilla (18951, -0.38 DPS) [quest]; Alchemist's Stone (13503, -1.53 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (195.5 DPS) | yes | Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Alchemist's Stone (13503, -1.13 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -7.41 DPS) [dungeon]; Hand of Righteousness (7721, -10.69 DPS) [dungeon]; Spire of Hakkar (10844, -12.53 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 14.5 healing_power points (1.53 DPS) | yes | Wand of Mana Concentration (273043, -0.15 DPS) [world]; Flash Wand (5248, -0.32 DPS) [quest]; Goblin Igniter (5253, -0.32 DPS) [quest] |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Nethergeld Shoulders; back: Imperial Red Cloak; wrist: Nethergeld Cuffs; hands: Gilded Gloves; waist: Dawnspire Cord; legs: Kilt of the Atal'ai Prophet; feet: Gilded Sandals; finger1: Eye of Adaegus; finger2: Welken Ring; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 507, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60 (troll, 025003031330000000-03505003030121431-00000000000000000)

Set DPS (verified): 327.7. Weights run: 9.2s. Verify run: 66.9s. 1193 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.112, intellect=0.958 ± 0.028, spirit=1.169 ± 0.027, mp5=2.088 ± 0.037, crit=0.184 ± 0.014 per rating point (14 rating = 1%, 2.580 per %), spell_haste=-16.256 ± 0.701

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (327.7 DPS) | yes | Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -1.64 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 20.7 healing_power points (4.45 DPS) | yes | Horizon Choker (13085, -0.64 DPS, sim-verified) [world_drop]; Lady Alizabeth's Pendant (13002, -0.67 DPS) [world_drop]; Necklace of Calisea (1714, -1.24 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (327.7 DPS) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp]; Mooncloth Shoulders (14139, -3.19 DPS, sim-verified) [crafted] |
| back | Wingveil Cloak (10802) | Sunken Temple: Spawn of Hakkar [dungeon] | 17.9 healing_power points (3.84 DPS) | yes | Glowing Thresher Cape (6901, -0.36 DPS, sim-verified) [dungeon]; Imperial Red Cloak (8248, -0.57 DPS) [world_drop]; Cloak of Blight (6832, -0.57 DPS) [quest] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 91.7 healing_power points (19.73 DPS) | yes | Warlord's Satin Tunic (231632, -2.88 DPS) [vendor]; Legionnaire's Satin Tunic (227124, -4.51 DPS) [pvp]; Mooncloth Vest (14138, -5.14 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (327.7 DPS) | yes | Nethergeld Cuffs (254061, +0.00 DPS) [crafted]; General's Satin Bracers (17619, -0.10 DPS) [pvp]; Earthenweave Cuffs (254131, -3.80 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (327.7 DPS) | yes | Blood Guard's Satin Handwraps (227126, +0.00 DPS) [pvp]; General's Satin Handwraps (231633, -0.65 DPS) [vendor]; Gilded Gloves (254095, -3.41 DPS, sim-verified) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 54.9 healing_power points (11.82 DPS) | yes | General's Satin Cinch (17621, -3.31 DPS) [pvp]; Earthenweave Cord (254077, -3.46 DPS) [crafted]; Devout Belt (16696, -4.23 DPS, sim-verified) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 65.3 healing_power points (14.04 DPS) | yes | Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Earthenweave Leggings (254125, -7.52 DPS, sim-verified) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (327.7 DPS) | yes | General's Satin Walkers (231630, +0.00 DPS) [vendor]; Swarmtender's Footpads (275609, +0.00 DPS) [crafted]; Mooncloth Boots (15802, -0.96 DPS, sim-verified) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 22.4 healing_power points (4.83 DPS) | yes | Chemist's Ring (15702, -1.76 DPS) [quest]; Electrocutioner Lagnut (9447, -1.78 DPS) [dungeon]; Runemender's Seal (273031, -1.78 DPS) [world] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 19.8 healing_power points (4.26 DPS) | yes | Chemist's Ring (15702, -0.63 DPS, sim-verified) [quest]; Electrocutioner Lagnut (9447, -1.20 DPS) [dungeon]; Runemender's Seal (273031, -1.20 DPS) [world] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (327.7 DPS) | yes | Evonice's Landin' Pilla (18951, -0.18 DPS) [quest]; Ankh of Life (1713, -0.39 DPS, sim-verified) [world_drop]; Alchemist's Stone (13503, -1.69 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (327.7 DPS) | yes | Ankh of Life (1713, -0.85 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.36 DPS) [quest]; Mark of the Chosen (17774, -5.21 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (327.7 DPS) | yes | Charstone Dirk (17710, -12.39 DPS) [dungeon]; Hand of Righteousness (7721, -18.28 DPS) [dungeon]; Hand of Edward the Odd (2243, -20.28 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 29.9 healing_power points (6.43 DPS) | yes | Jaina's Firestarter (13064, -4.43 DPS) [world_drop]; Wand of Mana Concentration (273043, -4.70 DPS) [world]; Captain Rackmore's Tiller (16789, -4.84 DPS, sim-verified) [quest] |

**New at 60:** head: Devout Crown; neck: Glowing Eye of Mordresh; shoulder: Devout Mantle; back: Wingveil Cloak; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Shard of the Splithooves; ranged: Torch of Light

No-known-source sample (15 of 1193, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60, raid preset (troll, 025003031330000000-03505003030121431-00000000000000000)

Set DPS (verified): 760.8. Weights run: 5.9s. Verify run: 31.5s. 1193 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.173, intellect=1.011 ± 0.033, spirit=0.924 ± 0.031, mp5=1.887 ± 0.034, crit=0.379 ± 0.020 per rating point (14 rating = 1%, 5.311 per %), spell_haste=-3.552 ± 0.630

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (760.9 DPS) | yes | Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -19.00 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 21.1 healing_power points (6.22 DPS) | yes | Horizon Choker (13085, -0.93 DPS, sim-verified) [world_drop]; Lady Alizabeth's Pendant (13002, -2.13 DPS) [world_drop]; Necklace of Calisea (1714, -2.22 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (760.9 DPS) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp]; Mooncloth Shoulders (14139, -17.63 DPS, sim-verified) [crafted] |
| back | Wingveil Cloak (10802) | Sunken Temple: Spawn of Hakkar [dungeon] | 15.2 healing_power points (4.47 DPS) | yes | Imperial Red Cloak (8248, -0.10 DPS) [world_drop]; Glowing Thresher Cape (6901, -0.23 DPS) [dungeon]; Highborne Cloak (14450, -0.70 DPS) [world_drop] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 87.8 healing_power points (25.88 DPS) | yes | Warlord's Satin Tunic (231632, -3.00 DPS) [vendor]; Legionnaire's Satin Tunic (227124, -5.81 DPS) [pvp]; Stormcloth Vest (10020, -9.26 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (760.9 DPS) | yes | General's Satin Bracers (17619, +0.00 DPS) [pvp]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -16.16 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (760.9 DPS) | yes | Blood Guard's Satin Handwraps (227126, +0.00 DPS) [pvp]; General's Satin Handwraps (231633, +0.00 DPS) [vendor]; Gilded Gloves (254095, -16.38 DPS, sim-verified) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 53.5 healing_power points (15.76 DPS) | yes | Devout Belt (16696, +0.00 DPS, sim-verified) [dungeon]; Gilded Waistcord (254081, -4.53 DPS) [crafted]; General's Satin Cinch (17621, -4.82 DPS) [pvp] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 60.4 healing_power points (17.81 DPS) | yes | Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Earthenweave Leggings (254125, -25.69 DPS, sim-verified) [crafted] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 50.2 healing_power points (14.79 DPS) | yes | General's Satin Walkers (231630, -0.87 DPS) [vendor]; Blood Guard's Satin Walkers (227127, -2.65 DPS) [pvp]; Swarmtender's Footpads (275609, -4.49 DPS, sim-verified) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 20.3 healing_power points (5.98 DPS) | yes | Chemist's Ring (15702, -1.91 DPS) [quest]; Maiden's Circle (13001, -2.26 DPS) [world_drop]; Electrocutioner Lagnut (9447, -2.30 DPS) [dungeon] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 17.1 healing_power points (5.06 DPS) | yes | Maiden's Circle (13001, -1.34 DPS) [world_drop]; Electrocutioner Lagnut (9447, -1.38 DPS) [dungeon]; Chemist's Ring (15702, -2.05 DPS, sim-verified) [quest] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (760.9 DPS) | yes | Ankh of Life (1713, -0.07 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.62 DPS) [quest]; Alchemist's Stone (13503, -2.25 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (760.9 DPS) | yes | Ankh of Life (1713, -1.84 DPS, sim-verified) [world_drop]; Evonice's Landin' Pilla (18951, -2.58 DPS) [quest]; Alchemist's Stone (13503, -4.22 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (760.9 DPS) | yes | Charstone Dirk (17710, -16.72 DPS) [dungeon]; Hand of Righteousness (7721, -25.42 DPS) [dungeon]; Hand of Edward the Odd (2243, -40.13 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 27.2 healing_power points (8.01 DPS) | yes | Captain Rackmore's Tiller (16789, -5.42 DPS) [quest]; Wand of Mana Concentration (273043, -5.70 DPS) [world]; Jaina's Firestarter (13064, -6.83 DPS, sim-verified) [world_drop] |

**New at 60:** head: Devout Crown; neck: Glowing Eye of Mordresh; shoulder: Devout Mantle; back: Wingveil Cloak; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Shard of the Splithooves; ranged: Torch of Light

No-known-source sample (15 of 1193, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

