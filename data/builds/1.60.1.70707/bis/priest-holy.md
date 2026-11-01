# Leveling BiS: Holy

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-03503000000000000-00000000000000000)

Set DPS (verified): 32.4. Weights run: 7.8s. Verify run: 23.7s. 163 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.008, intellect=0.946 ± 0.005, spirit=0.694 ± 0.004, mp5=1.729 ± 0.014, crit=0.061 ± 0.003 per rating point (14 rating = 1%, 0.859 per %), spell_haste=not significant (0.019 ± 0.032)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.43 DPS) | yes | Shadow Goggles (4373, -0.76 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.21 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 2.8 healing_power points (0.22 DPS) | yes | Scholarly Pendant (277203, -0.24 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.5 healing_power points (0.68 DPS) | yes | Slime-encrusted Pads (6461, -0.16 DPS, sim-verified) [dungeon]; Reinforced Woolen Shoulders (4315, -0.38 DPS) [crafted] |
| back | Sanguine Cape (14376) | World drop [world_drop] | 3.8 healing_power points (0.30 DPS) | yes | Seer's Cape (6378, +0.00 DPS, sim-verified) [dungeon]; Pearl-clasped Cloak (5542, -0.08 DPS) [crafted]; Spirit Cloak (4792, -0.08 DPS) [vendor] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 19.8 healing_power points (1.58 DPS) | yes | Corsair's Overshirt (5202, -0.24 DPS, sim-verified) [dungeon]; Robe of the Moccasin (6465, -0.31 DPS) [dungeon]; Bloody Apron (6226, -0.54 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 9.9 healing_power points (0.79 DPS) | yes | Repurposed Hair Band (281256, -0.53 DPS) [quest]; Seer's Cuffs (3645, -0.60 DPS) [dungeon]; Bright Bracers (3647, -1.13 DPS, sim-verified) [world_drop] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 13.8 healing_power points (1.10 DPS) | yes | Magefist Gloves (12977, -0.43 DPS, sim-verified) [world_drop]; Bright Gloves (3066, -0.58 DPS) [world_drop]; Tomb Robber's Gloves (280096, -0.65 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 14.8 healing_power points (1.18 DPS) | yes | Novice Ardent's Sash (253887, -0.28 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.57 DPS) [world_drop]; Relic Hunter Belt (11936, -0.77 DPS) [quest] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 26.5 healing_power points (2.11 DPS) | yes | Filigreed Silky Leggings (253939, -1.43 DPS) [crafted]; Filigreed Shadow Leggings (253943, -1.43 DPS) [crafted]; Darkweave Breeches (12987, -1.53 DPS, sim-verified) [world_drop] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 11.8 healing_power points (0.94 DPS) | yes | Kimbra Boots (6191, -0.51 DPS, sim-verified) [quest]; Bluegill Sandals (1560, -0.57 DPS) [world]; Smoldering Boots (3076, -0.61 DPS) [world] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.7 healing_power points (0.45 DPS) | yes | Deep Fathom Ring (6463, -0.18 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.23 DPS) [world_drop]; Minor Channeling Ring (1449, -0.30 DPS) [quest] |
| finger2 | Band of Purification (12996) | World drop [world_drop] | 4.2 healing_power points (0.33 DPS) | yes | Deep Fathom Ring (6463, -0.06 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.11 DPS) [world_drop]; Minor Channeling Ring (1449, -0.18 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 9.5 healing_power points (0.75 DPS) | yes | Staff of Westfall (2042, +0.00 DPS, sim-verified) [quest]; Staff of the Blessed Seer (2271, -0.20 DPS) [dungeon]; Lesser Staff of the Spire (1300, -0.30 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Sable Wand (7607) | Underground Assault [quest] | 2.1 healing_power points (0.17 DPS) | yes | Moonstone Wand (15204, -0.02 DPS) [quest]; Dwarven Flamestick (5241, -0.06 DPS) [quest]; Flaring Baton (5326, -0.09 DPS, sim-verified) [quest] |

**New at 20:** head: Pristine Circlet; neck: Tarnished Locket; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Band of Purification; main_hand: Twisted Chanter's Staff; ranged: Sable Wand

No-known-source sample (15 of 163, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes; 16607 Acolyte's Sacrificial Robes; 18851 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (gnome, 000000000000000000-03505003030110000-00000000000000000)

Set DPS (verified): 51.6. Weights run: 8.1s. Verify run: 27.6s. 294 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.013, intellect=1.040 ± 0.011, spirit=0.786 ± 0.006, mp5=2.046 ± 0.014, crit=0.081 ± 0.005 per rating point (14 rating = 1%, 1.131 per %), spell_haste=not significant (0.106 ± 0.047)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 37.7 healing_power points (3.05 DPS) | yes | Filigreed Pristine Circlet (253975, -1.39 DPS, sim-verified) [crafted]; Pristine Circlet (253949, -1.59 DPS) [crafted]; Nightsky Cowl (4039, -1.72 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 7.3 healing_power points (0.59 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.09 DPS) [world]; Crystal Starfire Medallion (5003, -0.12 DPS, sim-verified) [world_drop]; Pendant of Myzrael (4614, -0.21 DPS) [dungeon] |
| shoulder | Mantle of Honor (3560) | Bride of the Embalmer [quest] | 12.8 healing_power points (1.03 DPS) | yes | Nightsky Mantle (4718, -0.04 DPS) [world_drop]; Death Speaker Mantle (6685, -0.11 DPS) [dungeon]; Faerie Mantle (5820, -0.15 DPS) [quest] |
| back | Prelacy Cape (7004) | Researching the Corruption [quest] | 15.7 healing_power points (1.27 DPS) | yes | Glowing Thresher Cape (6901, -0.31 DPS, sim-verified) [dungeon]; Repairman's Cape (9605, -0.38 DPS) [quest]; Sergeant Major's Cape (16315, -0.68 DPS) [pvp] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 31.4 healing_power points (2.54 DPS) | yes | Pristine Gown (253961, -0.19 DPS, sim-verified) [crafted]; Robes of Arugal (6324, -0.71 DPS) [dungeon]; Filigreed Pristine Gown (253901, -0.88 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 11.3 healing_power points (0.92 DPS) | yes | Glowing Magical Bracelets (13106, -0.24 DPS) [world_drop]; Nightsky Wristbands (6407, -0.30 DPS, sim-verified) [world_drop]; Spidertank Oilrag (9448, -0.42 DPS) [dungeon] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 18.1 healing_power points (1.46 DPS) | yes | Gloves of Old (9395, -0.17 DPS, sim-verified) [world_drop]; Pristine Gloves (253913, -0.32 DPS) [crafted]; Hotshot Pilot's Gloves (9491, -0.47 DPS) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 15.2 healing_power points (1.22 DPS) | yes | Resilient Cord (14406, -0.47 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.48 DPS, sim-verified) [crafted]; Dreamer's Belt (4829, -0.51 DPS) [vendor] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 33.2 healing_power points (2.68 DPS) | yes | Filigreed Pristine Leggings (253937, -0.62 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -0.77 DPS) [crafted]; Darkweave Breeches (12987, -1.71 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 30.4 healing_power points (2.46 DPS) | yes | Nimbus Boots (6998, -1.15 DPS, sim-verified) [quest]; Pristine Boots (253889, -1.48 DPS) [crafted]; Acidic Walkers (9454, -1.53 DPS) [dungeon] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 11.5 healing_power points (0.93 DPS) | yes | The Queen's Jewel (13094, -0.25 DPS) [world_drop]; Defias Mage Ring (1077, -0.40 DPS) [world]; Seal of Wrynn (2933, -0.40 DPS) [quest] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 11.5 healing_power points (0.93 DPS) | yes | The Queen's Jewel (13094, -0.31 DPS, sim-verified) [world_drop]; Defias Mage Ring (1077, -0.40 DPS) [world]; Seal of Wrynn (2933, -0.40 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-verified (51.6 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -0.72 DPS) [quest]; Death Speaker Scepter (2816, -1.99 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 8.3 healing_power points (0.67 DPS) | yes | Lesser Mystic Wand (11289, -0.33 DPS) [crafted]; Starfaller (13063, -0.33 DPS) [world_drop]; Summoner's Wand (5245, -0.39 DPS, sim-verified) [world] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Mantle of Honor; back: Prelacy Cape; chest: Death Speaker Robes; hands: Truefaith Gloves; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Spine of the Devourer; ranged: Wand of Mana Concentration

No-known-source sample (15 of 294, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 40 (gnome, 000000000000000000-03505003030121431-00000000000000000)

Set DPS (verified): 94.8. Weights run: 8.5s. Verify run: 32.0s. 396 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.039, intellect=0.836 ± 0.024, spirit=1.255 ± 0.026, mp5=2.662 ± 0.025, crit=0.130 ± 0.008 per rating point (14 rating = 1%, 1.821 per %), spell_haste=-6.644 ± 0.430

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | sim-verified (94.8 DPS) | yes | Holy Shroud (2721, +0.00 DPS) [world_drop]; Whitemane's Chapeau (7720, -0.29 DPS) [dungeon]; Earthen Silk Hood (254015, -0.47 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 19.7 healing_power points (1.87 DPS) | yes | Necklace of Calisea (1714, -0.48 DPS) [dungeon]; Amberglow Talisman (10824, -0.68 DPS) [quest]; Explorers' League Commendation (7746, -0.80 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | sim-verified (94.8 DPS) | yes | Earthen Silk Shoulders (254033, +0.00 DPS) [crafted]; Crimson Silk Shoulders (7059, -0.99 DPS) [crafted]; Mistscape Mantle (4734, -0.99 DPS) [dungeon] |
| back | Blackforge Cape (6424) | Uldaman: Ancient Treasure [dungeon] | sim-verified (94.8 DPS) | yes | Ceremonial Centaur Blanket (6789, +0.00 DPS) [quest]; Glowing Thresher Cape (6901, +0.00 DPS) [dungeon]; Prelacy Cape (7004, +0.00 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 68.0 healing_power points (6.48 DPS) | yes | Doomsayer's Robe (4746, -3.37 DPS) [quest]; Civinad Robes (9623, -3.42 DPS) [quest]; Pristine Gown (253961, -3.54 DPS) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | sim-verified (94.8 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS) [dungeon]; Silkstream Cuffs (16791, +0.00 DPS) [quest]; Earthen Silk Cuffs (254019, +0.00 DPS) [crafted] |
| hands | Gilded Handwraps (254021) | Tailoring [crafted] | sim-verified (94.8 DPS) | yes | Earthen Silk Gloves (254017, +0.00 DPS) [crafted]; Stormcloth Gloves (10011, -0.65 DPS) [crafted]; Gloves of Old (9395, -0.82 DPS) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 36.2 healing_power points (3.45 DPS) | yes | Earthen Silk Belt (7061, -1.73 DPS) [crafted]; Windchaser Cinch (14435, -1.74 DPS) [world_drop]; Sutarn's Ring (13105, -1.93 DPS) [world_drop] |
| legs | Stoneweaver Leggings (9407) | Uldaman: Ironaya [dungeon] | 34.8 healing_power points (3.31 DPS) | yes | Pristine Leggings (253987, -0.06 DPS) [crafted]; Stormcloth Pants (10010, -0.09 DPS) [crafted]; Earthen Leggings (253999, -0.57 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 30.9 healing_power points (2.94 DPS) | yes | Furen's Boots (13100, -0.47 DPS) [world_drop]; Thoughtcast Boots (10578, -0.62 DPS) [dungeon]; Everlast Boots (10359, -1.20 DPS) [quest] |
| finger1 | Welken Ring (5011) | World drop [world_drop] | sim-verified (94.8 DPS) | yes | Electrocutioner Lagnut (9447, +0.00 DPS) [dungeon]; Runemender's Seal (273031, +0.00 DPS) [world]; The Queen's Jewel (13094, -0.12 DPS) [world_drop] |
| finger2 | Voodoo Band (1996) | Bloodscalp Witch Doctor [world] | sim-verified (94.8 DPS) | yes | Electrocutioner Lagnut (9447, +0.00 DPS) [dungeon]; The Queen's Jewel (13094, +0.00 DPS) [world_drop]; Runemender's Seal (273031, +0.00 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, -3.43 DPS) [dungeon]; Illusionary Rod (7713, -4.40 DPS) [dungeon]; Hypnotic Blade (7714, -4.72 DPS) [dungeon] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | sim-verified (94.8 DPS) | yes | Orb of Souls (249395, +0.00 DPS) [crafted]; Beacon of Hope (9393, -0.13 DPS) [dungeon]; Orb of Mistmantle (13031, -0.60 DPS) [world_drop] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | sim-verified (94.8 DPS) | yes | Captain Rackmore's Tiller (16789, +0.00 DPS) [quest]; Wand of Mana Concentration (273043, -0.01 DPS) [world]; Goblin Igniter (5253, -0.16 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Blackforge Cape; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Gilded Handwraps; waist: Gilded Cord; legs: Stoneweaver Leggings; finger1: Welken Ring; finger2: Voodoo Band; trinket1: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Orb of Lorica; ranged: Jaina's Firestarter

No-known-source sample (15 of 396, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 50 (gnome, 025003000000000000-03505003030121431-00000000000000000)

Set DPS (verified): 179.9. Weights run: 8.3s. Verify run: 31.2s. 515 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.069, intellect=1.115 ± 0.037, spirit=2.197 ± 0.039, mp5=3.970 ± 0.040, crit=0.258 ± 0.015 per rating point (14 rating = 1%, 3.617 per %), spell_haste=-6.978 ± 0.657

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 67.5 healing_power points (6.16 DPS) | yes | Cassandra's Grace (13102, -0.34 DPS) [world_drop]; Knight-Lieutenant's Satin Cover (220896, -0.38 DPS) [vendor]; Stormcloth Headband (10032, -0.95 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 24.4 healing_power points (2.23 DPS) | yes | Necklace of Calisea (1714, -0.11 DPS) [dungeon]; Amberglow Talisman (10824, -0.22 DPS) [quest]; Glowing Eye of Mordresh (10769, -0.23 DPS) [dungeon] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | 48.1 healing_power points (4.39 DPS) | yes | Knight-Lieutenant's Satin Pads (220894, +0.00 DPS) [vendor]; Earthen Silk Shoulders (254033, -1.14 DPS) [crafted]; Hibernal Mantle (8111, -1.37 DPS) [world_drop] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | sim-verified (180.0 DPS) | yes | Arcane Cloak (8286, +0.00 DPS) [world_drop]; Chainlink Towel (9648, +0.00 DPS) [quest]; Wingveil Cloak (10802, +0.00 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ghostweave Vest (14141, +0.00 DPS) [crafted]; Grimlok's Tribal Vestments (9415, -1.85 DPS) [dungeon]; Knight's Satin Armor (220892, -1.94 DPS) [vendor] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 38.8 healing_power points (3.54 DPS) | yes | Venomshroud Armguards (14439, -1.53 DPS) [world_drop]; Earthen Silk Cuffs (254019, -1.62 DPS) [crafted]; Shizzle's Nozzle Wiper (11917, -1.72 DPS) [quest] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 57.7 healing_power points (5.27 DPS) | yes | Earthenweave Gloves (254075, -0.69 DPS) [crafted]; Earthen Silk Gloves (254017, -2.02 DPS) [crafted]; Sergeant Major's Satin Gloves (220897, -2.03 DPS) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | sim-verified (180.0 DPS) | yes | Gilded Cord (254037, +0.00 DPS) [crafted]; Earthenweave Cord (254077, +0.00 DPS) [crafted]; Gilded Waistcord (254081, +0.00 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | sim-verified (180.0 DPS) | yes | Senior Designer's Pantaloons (11841, +0.00 DPS) [dungeon]; Dalewind Trousers (13008, -0.06 DPS) [world_drop]; Knight's Satin Leggings (220893, -1.03 DPS) [vendor] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | sim-verified (180.0 DPS) | yes | Furen's Boots (13100, +0.00 DPS) [world_drop]; Vinerot Sandals (17748, +0.00 DPS) [dungeon]; Earthenweave Boots (254093, +0.00 DPS) [crafted] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 33.1 healing_power points (3.01 DPS) | yes | Choking Band (11868, -0.41 DPS) [quest]; Coldwater Ring (4550, -0.91 DPS) [quest]; Seedtime Hoop (9655, -1.01 DPS) [quest] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | sim-verified (180.0 DPS) | yes | Coldwater Ring (4550, +0.00 DPS) [quest]; Seedtime Hoop (9655, +0.00 DPS) [quest]; Choking Band (11868, +0.00 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (180.0 DPS) | yes | Evonice's Landin' Pilla (18951, -0.40 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.80 DPS) [quest]; Alchemist's Stone (13503, -1.60 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (180.0 DPS) | yes | Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.04 DPS) [quest]; Alchemist's Stone (13503, -0.84 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -6.01 DPS) [dungeon]; Hand of Righteousness (7721, -8.26 DPS) [dungeon]; Spire of Hakkar (10844, -9.57 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cairnstone Sliver (9654) | The Morrow Stone [quest] | 16.6 healing_power points (1.51 DPS) | yes | Jaina's Firestarter (13064, -0.30 DPS) [world_drop]; Captain Rackmore's Tiller (16789, -0.36 DPS) [quest]; Wand of Mana Concentration (273043, -0.38 DPS) [world] |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Nethergeld Shoulders; back: Imperial Red Cloak; wrist: Nethergeld Cuffs; hands: Gilded Gloves; waist: Dawnspire Cord; legs: Kilt of the Atal'ai Prophet; feet: Gilded Sandals; finger1: Eye of Adaegus; finger2: Welken Ring; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Cairnstone Sliver

No-known-source sample (15 of 515, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60 (gnome, 025003031330000000-03505003030121431-00000000000000000)

Set DPS (verified): 310.0. Weights run: 7.2s. Verify run: 50.7s. 1200 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.078, intellect=0.945 ± 0.036, spirit=1.560 ± 0.033, mp5=2.714 ± 0.042, crit=0.188 ± 0.015 per rating point (14 rating = 1%, 2.637 per %), spell_haste=-18.305 ± 0.861

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (310.0 DPS) | yes | Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -2.18 DPS, sim-verified) [crafted] |
| neck | Lady Alizabeth's Pendant (13002) | World drop [world_drop] | 23.4 healing_power points (3.97 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -0.67 DPS) [world_drop]; Necklace of Calisea (1714, -1.00 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (310.0 DPS) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp]; Mooncloth Shoulders (14139, -2.14 DPS, sim-verified) [crafted] |
| back | Wingveil Cloak (10802) | Sunken Temple: Spawn of Hakkar [dungeon] | 22.2 healing_power points (3.76 DPS) | yes | Prelacy Cape (7004, -0.31 DPS) [quest]; Indomitable Cloak (14683, -0.32 DPS) [world_drop]; Ethereal Mist Cape (11873, -0.70 DPS, sim-verified) [quest] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 98.0 healing_power points (16.64 DPS) | yes | Field Marshal's Satin Tunic (231624, -2.80 DPS) [vendor]; Knight-Captain's Satin Tunic (227122, -3.73 DPS) [pvp]; Mooncloth Vest (14138, -7.13 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (310.0 DPS) | yes | Nethergeld Cuffs (254061, +0.00 DPS) [crafted]; Marshal's Satin Bracers (17606, -0.10 DPS) [pvp]; Earthenweave Cuffs (254131, -2.40 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | 51.0 healing_power points (8.66 DPS) | yes | Swarmtender's Gloves (275607, -0.64 DPS) [crafted]; Knight-Lieutenant's Satin Handwraps (227128, -1.06 DPS) [pvp]; Gilded Gloves (254095, -3.00 DPS, sim-verified) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 58.3 healing_power points (9.90 DPS) | yes | Marshal's Satin Sash (17609, -2.23 DPS) [pvp]; Earthenweave Cord (254077, -2.57 DPS) [crafted]; Devout Belt (16696, -4.31 DPS, sim-verified) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 74.0 healing_power points (12.58 DPS) | yes | Knight-Captain's Satin Legguards (227125, +0.00 DPS) [pvp]; Marshal's Satin Legguards (231626, -0.20 DPS) [vendor]; Mooncloth Leggings (14137, -6.53 DPS, sim-verified) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (310.0 DPS) | yes | Marshal's Satin Walkers (231627, -0.16 DPS) [vendor]; Mooncloth Boots (15802, -0.57 DPS, sim-verified) [crafted]; Earthenweave Boots (254093, -0.65 DPS) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 26.6 healing_power points (4.52 DPS) | yes | Choking Band (11868, -1.07 DPS) [quest]; Electrocutioner Lagnut (9447, -1.65 DPS) [dungeon]; Runemender's Seal (273031, -1.65 DPS) [world] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 24.4 healing_power points (4.14 DPS) | yes | Choking Band (11868, -1.00 DPS, sim-verified) [quest]; Electrocutioner Lagnut (9447, -1.27 DPS) [dungeon]; Runemender's Seal (273031, -1.27 DPS) [world] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (310.0 DPS) | yes | Ankh of Life (1713, +0.00 DPS, sim-verified) [world_drop]; Evonice's Landin' Pilla (18951, -0.12 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.65 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (310.0 DPS) | yes | Evonice's Landin' Pilla (18951, -0.41 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.94 DPS) [quest]; Ankh of Life (1713, -1.58 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (310.0 DPS) | yes | Charstone Dirk (17710, -10.13 DPS) [dungeon]; Hand of Righteousness (7721, -14.43 DPS) [dungeon]; Hand of Edward the Odd (2243, -19.76 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 34.2 healing_power points (5.80 DPS) | yes | Captain Rackmore's Tiller (16789, -3.99 DPS) [quest]; Jaina's Firestarter (13064, -4.04 DPS) [world_drop]; Cairnstone Sliver (9654, -4.90 DPS, sim-verified) [quest] |

**New at 60:** head: Devout Crown; neck: Lady Alizabeth's Pendant; shoulder: Devout Mantle; back: Wingveil Cloak; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Shard of the Splithooves; ranged: Torch of Light

No-known-source sample (15 of 1200, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60, raid preset (gnome, 025003031330000000-03505003030121431-00000000000000000)

Set DPS (verified): 656.7. Weights run: 4.8s. Verify run: 31.0s. 1200 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.156, intellect=0.665 ± 0.029, spirit=0.866 ± 0.028, mp5=1.700 ± 0.032, crit=0.302 ± 0.016 per rating point (14 rating = 1%, 4.233 per %), spell_haste=-2.810 ± 0.559

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (656.8 DPS) | yes | Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -9.27 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 18.3 healing_power points (5.91 DPS) | yes | Horizon Choker (13085, -1.79 DPS) [world_drop]; Necklace of Calisea (1714, -2.45 DPS) [dungeon]; Lady Alizabeth's Pendant (13002, -2.52 DPS, sim-verified) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (656.8 DPS) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp]; Mooncloth Shoulders (14139, -8.96 DPS, sim-verified) [crafted] |
| back | Prelacy Cape (7004) | Researching the Corruption [quest] | 16.2 healing_power points (5.23 DPS) | yes | Wingveil Cloak (10802, -0.54 DPS) [dungeon]; Glowing Thresher Cape (6901, -0.73 DPS) [dungeon]; Ethereal Mist Cape (11873, -1.23 DPS) [quest] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 86.9 healing_power points (28.03 DPS) | yes | Field Marshal's Satin Tunic (231624, -6.27 DPS) [vendor]; Knight-Captain's Satin Tunic (227122, -7.99 DPS) [pvp]; Stormcloth Vest (10020, -14.74 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (656.8 DPS) | yes | Nethergeld Cuffs (254061, +0.00 DPS) [crafted]; Marshal's Satin Bracers (17606, -1.19 DPS) [pvp]; Earthenweave Cuffs (254131, -11.15 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (656.8 DPS) | yes | Knight-Lieutenant's Satin Handwraps (227128, +0.00 DPS) [pvp]; Marshal's Satin Handwraps (231623, +0.00 DPS) [vendor]; Gilded Gloves (254095, -11.20 DPS, sim-verified) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 48.1 healing_power points (15.52 DPS) | yes | Gilded Waistcord (254081, -4.23 DPS) [crafted]; Gilded Cord (254037, -5.03 DPS) [crafted]; Earthenweave Cord (254077, -10.84 DPS, sim-verified) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 53.9 healing_power points (17.39 DPS) | yes | Knight-Captain's Satin Legguards (227125, +0.00 DPS) [pvp]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Earthenweave Leggings (254125, -18.02 DPS, sim-verified) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (656.8 DPS) | yes | Marshal's Satin Walkers (231627, +0.00 DPS) [vendor]; Swarmtender's Footpads (275609, +0.00 DPS) [crafted]; Mooncloth Boots (15802, -5.78 DPS, sim-verified) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 16.2 healing_power points (5.22 DPS) | yes | Electrocutioner Lagnut (9447, -1.33 DPS) [dungeon]; Runemender's Seal (273031, -1.33 DPS) [world]; Choking Band (11868, -1.59 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 14.4 healing_power points (4.64 DPS) | yes | Runemender's Seal (273031, -0.75 DPS) [world]; Choking Band (11868, -1.01 DPS) [quest]; Electrocutioner Lagnut (9447, -2.95 DPS, sim-verified) [dungeon] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (656.8 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.49 DPS) [quest]; Thunderbrew's Boot Flask (744, -1.05 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (656.8 DPS) | yes | Evonice's Landin' Pilla (18951, -3.01 DPS) [quest]; Thunderbrew's Boot Flask (744, -3.57 DPS) [quest]; Ankh of Life (1713, -4.00 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (656.8 DPS) | yes | Charstone Dirk (17710, -16.25 DPS) [dungeon]; Hand of Righteousness (7721, -24.57 DPS) [dungeon]; Hand of Edward the Odd (2243, -43.86 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 26.5 healing_power points (8.56 DPS) | yes | Cairnstone Sliver (9654, -6.09 DPS) [quest]; Jaina's Firestarter (13064, -6.43 DPS) [world_drop]; Captain Rackmore's Tiller (16789, -13.77 DPS, sim-verified) [quest] |

**New at 60:** head: Devout Crown; neck: Glowing Eye of Mordresh; shoulder: Devout Mantle; back: Prelacy Cape; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Shard of the Splithooves; ranged: Torch of Light

No-known-source sample (15 of 1200, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

## Horde

### Band 20 (troll, 000000000000000000-03503000000000000-00000000000000000)

Set DPS (verified): 31.2. Weights run: 7.8s. Verify run: 23.8s. 156 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.008, intellect=0.946 ± 0.005, spirit=0.694 ± 0.004, mp5=1.729 ± 0.014, crit=0.061 ± 0.003 per rating point (14 rating = 1%, 0.859 per %), spell_haste=not significant (0.019 ± 0.032)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.43 DPS) | yes | Shadow Goggles (4373, -0.93 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.21 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 2.1 healing_power points (0.17 DPS) | yes | Scholarly Pendant (277203, -0.06 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.5 healing_power points (0.68 DPS) | yes | Slime-encrusted Pads (6461, -0.27 DPS, sim-verified) [dungeon]; Reinforced Woolen Shoulders (4315, -0.38 DPS) [crafted] |
| back | Sanguine Cape (14376) | World drop [world_drop] | 3.8 healing_power points (0.30 DPS) | yes | Pearl-clasped Cloak (5542, -0.08 DPS) [crafted]; Spirit Cloak (4792, -0.08 DPS) [vendor]; Seer's Cape (6378, -0.17 DPS, sim-verified) [dungeon] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 19.8 healing_power points (1.58 DPS) | yes | Robe of the Moccasin (6465, -0.31 DPS) [dungeon]; Corsair's Overshirt (5202, -0.44 DPS, sim-verified) [dungeon]; Bloody Apron (6226, -0.54 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 9.9 healing_power points (0.79 DPS) | yes | Crystalline Cuffs (14148, -0.47 DPS) [dungeon]; Bright Bracers (3647, -0.49 DPS) [world_drop]; Featherbead Bracers (15452, -1.05 DPS, sim-verified) [quest] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 13.8 healing_power points (1.10 DPS) | yes | Blight Gloves (279877, -0.44 DPS, sim-verified) [quest]; Magefist Gloves (12977, -0.50 DPS) [world_drop]; Bright Gloves (3066, -0.58 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 14.8 healing_power points (1.18 DPS) | yes | Novice Ardent's Sash (253887, -0.40 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.57 DPS) [world_drop]; Wise Man's Belt (4786, -0.80 DPS) [vendor] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 26.5 healing_power points (2.11 DPS) | yes | Filigreed Silky Leggings (253939, -1.43 DPS) [crafted]; Filigreed Shadow Leggings (253943, -1.43 DPS) [crafted]; Darkweave Breeches (12987, -1.49 DPS, sim-verified) [world_drop] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 11.8 healing_power points (0.94 DPS) | yes | Smoldering Boots (3076, -0.61 DPS) [world]; Bluegill Sandals (1560, -0.64 DPS, sim-verified) [world]; Sanguine Sandals (14374, -0.64 DPS) [world_drop] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.7 healing_power points (0.45 DPS) | yes | Band of Purification (12996, -0.12 DPS) [world_drop]; Deep Fathom Ring (6463, -0.18 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.23 DPS) [world_drop] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 4.7 healing_power points (0.38 DPS) | yes | Deep Fathom Ring (6463, -0.10 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.15 DPS) [world_drop]; Band of Purification (12996, -0.23 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Orgrimmar (15444) | Hidden Enemies [quest] | 10.7 healing_power points (0.85 DPS) | yes | Twisted Chanter's Staff (890, -0.10 DPS) [world_drop]; Staff of the Blessed Seer (2271, -0.30 DPS) [dungeon]; Gnarled Necromancer's Staff (251534, -0.55 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Decay (5252) (or Flaring Baton (5326)) | Beren's Peril [quest] | 1.9 healing_power points (0.15 DPS) | yes | Flaring Baton (5326, +0.00 DPS) [quest]; Wisesight Wand (286750, -0.08 DPS) [world] |

**New at 20:** head: Pristine Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Staff of Orgrimmar; ranged: Wand of Decay

No-known-source sample (15 of 156, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (troll, 000000000000000000-03505003030110000-00000000000000000)

Set DPS (verified): 49.7. Weights run: 8.1s. Verify run: 28.1s. 286 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.013, intellect=1.040 ± 0.011, spirit=0.786 ± 0.006, mp5=2.046 ± 0.014, crit=0.081 ± 0.005 per rating point (14 rating = 1%, 1.131 per %), spell_haste=not significant (0.106 ± 0.047)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 37.7 healing_power points (3.05 DPS) | yes | Filigreed Pristine Circlet (253975, -1.38 DPS, sim-verified) [crafted]; Pristine Circlet (253949, -1.59 DPS) [crafted]; Nightsky Cowl (4039, -1.72 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 7.3 healing_power points (0.59 DPS) | yes | Crystal Starfire Medallion (5003, -0.08 DPS, sim-verified) [world_drop]; Archmagister's Faceted Pendant (273049, -0.09 DPS) [world]; Pendant of Myzrael (4614, -0.21 DPS) [dungeon] |
| shoulder | Ghostly Mantle (3324) | Deathstalkers in Shadowfang [quest] | 16.1 healing_power points (1.30 DPS) | yes | Death Speaker Mantle (6685, -0.37 DPS) [dungeon]; Talbar Mantle (10657, -0.46 DPS) [quest]; Nightsky Mantle (4718, -0.55 DPS, sim-verified) [world_drop] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 13.3 healing_power points (1.07 DPS) | yes | Resilient Cape (14400, -0.51 DPS, sim-verified) [world_drop]; Silver-thread Cloak (4713, -0.61 DPS) [world_drop]; Amy's Blanket (13005, -0.63 DPS) [world_drop] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 31.4 healing_power points (2.54 DPS) | yes | Pristine Gown (253961, -0.07 DPS, sim-verified) [crafted]; Robes of Arugal (6324, -0.71 DPS) [dungeon]; Filigreed Pristine Gown (253901, -0.88 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 11.3 healing_power points (0.92 DPS) | yes | Glowing Magical Bracelets (13106, -0.24 DPS) [world_drop]; Nightsky Wristbands (6407, -0.31 DPS, sim-verified) [world_drop]; Spidertank Oilrag (9448, -0.42 DPS) [dungeon] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 18.1 healing_power points (1.46 DPS) | yes | Tattered Mittens (270030, -0.15 DPS, sim-verified) [quest]; Gloves of Old (9395, -0.28 DPS) [world_drop]; Pristine Gloves (253913, -0.32 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 15.2 healing_power points (1.22 DPS) | yes | Novice Ardent's Sash (253887, -0.30 DPS) [crafted]; Lilac Sash (6780, -0.36 DPS, sim-verified) [quest]; Resilient Cord (14406, -0.47 DPS) [world_drop] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 33.2 healing_power points (2.68 DPS) | yes | Filigreed Pristine Leggings (253937, -0.58 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -0.77 DPS) [crafted]; Sacred Burial Trousers (6282, -1.69 DPS) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 30.4 healing_power points (2.46 DPS) | yes | Acidic Walkers (9454, -1.53 DPS) [dungeon]; Pristine Boots (253889, -1.57 DPS, sim-verified) [crafted]; Frothing Slippers (254003, -1.62 DPS) [crafted] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 11.5 healing_power points (0.93 DPS) | yes | Solid Gold Band (279845, -0.05 DPS) [quest]; The Queen's Jewel (13094, -0.25 DPS) [world_drop]; Defias Mage Ring (1077, -0.40 DPS) [world] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 11.5 healing_power points (0.93 DPS) | yes | Solid Gold Band (279845, -0.05 DPS) [quest]; The Queen's Jewel (13094, -0.25 DPS) [world_drop]; Defias Mage Ring (1077, -0.40 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-verified (49.7 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Staff of the Friar (3415, -1.89 DPS) [dungeon]; Death Speaker Scepter (2816, -1.91 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 8.3 healing_power points (0.67 DPS) | yes | Lesser Mystic Wand (11289, -0.33 DPS) [crafted]; Starfaller (13063, -0.33 DPS) [world_drop]; Summoner's Wand (5245, -0.36 DPS, sim-verified) [world] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Ghostly Mantle; back: Glowing Thresher Cape; chest: Death Speaker Robes; hands: Truefaith Gloves; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Spine of the Devourer; ranged: Wand of Mana Concentration

No-known-source sample (15 of 286, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (troll, 000000000000000000-03505003030121431-00000000000000000)

Set DPS (verified): 91.5. Weights run: 8.5s. Verify run: 33.9s. 388 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.039, intellect=0.836 ± 0.024, spirit=1.255 ± 0.026, mp5=2.662 ± 0.025, crit=0.130 ± 0.008 per rating point (14 rating = 1%, 1.821 per %), spell_haste=-6.644 ± 0.430

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | sim-verified (91.5 DPS) | yes | Whitemane's Chapeau (7720, -0.29 DPS) [dungeon]; Earthen Silk Hood (254015, -0.47 DPS) [crafted]; Holy Shroud (2721, -0.98 DPS, sim-verified) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 19.7 healing_power points (1.87 DPS) | yes | Necklace of Calisea (1714, -0.55 DPS, sim-verified) [dungeon]; Amberglow Talisman (10824, -0.68 DPS) [quest]; Triune Amulet (7722, -1.04 DPS) [dungeon] |
| shoulder | Earthen Silk Shoulders (254033) | Tailoring [crafted] | 28.0 healing_power points (2.67 DPS) | yes | Inquisitor's Shawl (19507, -0.21 DPS) [dungeon]; Ghostly Mantle (3324, -0.74 DPS) [quest]; Crimson Silk Shoulders (7059, -1.20 DPS) [crafted] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 17.0 healing_power points (1.62 DPS) | yes | Cloak of Blight (6832, -0.35 DPS, sim-verified) [quest]; Ceremonial Centaur Blanket (6789, -0.39 DPS) [quest]; Stargazer Cloak (9660, -0.43 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 68.0 healing_power points (6.48 DPS) | yes | Civinad Robes (9623, -3.42 DPS) [quest]; Pristine Gown (253961, -3.54 DPS) [crafted]; Doomsayer's Robe (4746, -4.02 DPS, sim-verified) [quest] |
| wrist | Earthen Silk Cuffs (254019) | Tailoring [crafted] | 16.3 healing_power points (1.55 DPS) | yes | Mindthrust Bracers (1974, -0.39 DPS) [dungeon]; Silkstream Cuffs (16791, -0.67 DPS) [quest]; Nightsky Wristbands (6407, -0.71 DPS) [world_drop] |
| hands | Earthen Silk Gloves (254017) | Tailoring [crafted] | 28.0 healing_power points (2.67 DPS) | yes | Gilded Handwraps (254021, -0.26 DPS, sim-verified) [crafted]; Stormcloth Gloves (10011, -0.67 DPS) [crafted]; Gloves of Old (9395, -0.84 DPS) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 36.2 healing_power points (3.45 DPS) | yes | Windchaser Cinch (14435, -1.74 DPS) [world_drop]; Sutarn's Ring (13105, -1.93 DPS) [world_drop]; Earthen Silk Belt (7061, -2.14 DPS, sim-verified) [crafted] |
| legs | Stoneweaver Leggings (9407) | Uldaman: Ironaya [dungeon] | 34.8 healing_power points (3.31 DPS) | yes | Stormcloth Pants (10010, -0.09 DPS) [crafted]; Pristine Leggings (253987, -0.21 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -0.57 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 30.9 healing_power points (2.94 DPS) | yes | Thoughtcast Boots (10578, -0.62 DPS) [dungeon]; Furen's Boots (13100, -0.63 DPS, sim-verified) [world_drop]; Everlast Boots (10359, -1.20 DPS) [quest] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 14.8 healing_power points (1.41 DPS) | yes | Solid Gold Band (279845, -0.14 DPS) [quest]; Welken Ring (5011, -0.17 DPS) [world_drop]; The Queen's Jewel (13094, -0.29 DPS) [world_drop] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 14.8 healing_power points (1.41 DPS) | yes | Solid Gold Band (279845, -0.14 DPS) [quest]; Welken Ring (5011, -0.17 DPS) [world_drop]; The Queen's Jewel (13094, -0.29 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Death Speaker Scepter (2816, -3.43 DPS) [dungeon]; Illusionary Rod (7713, -4.40 DPS) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 20.5 healing_power points (1.95 DPS) | yes | Beacon of Hope (9393, -0.14 DPS) [dungeon]; Orb of Mistmantle (13031, -0.62 DPS) [world_drop]; Silksand Star (15964, -0.88 DPS) [world_drop] |
| ranged | Captain Rackmore's Tiller (16789) | Claim Rackmore's Treasure! [quest] | 9.8 healing_power points (0.93 DPS) | yes | Jaina's Firestarter (13064, -0.09 DPS) [world_drop]; Wand of Mana Concentration (273043, -0.10 DPS) [world]; Goblin Igniter (5253, -0.25 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Earthen Silk Shoulders; chest: Stormcloth Vest; wrist: Earthen Silk Cuffs; hands: Earthen Silk Gloves; waist: Gilded Cord; legs: Stoneweaver Leggings; trinket1: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Orb of Souls; ranged: Captain Rackmore's Tiller

No-known-source sample (15 of 388, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16336 Sergeant Major's Cape

### Band 50 (troll, 025003000000000000-03505003030121431-00000000000000000)

Set DPS (verified): 172.6. Weights run: 8.3s. Verify run: 31.1s. 507 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.069, intellect=1.115 ± 0.037, spirit=2.197 ± 0.039, mp5=3.970 ± 0.040, crit=0.258 ± 0.015 per rating point (14 rating = 1%, 3.617 per %), spell_haste=-6.978 ± 0.657

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 67.5 healing_power points (6.16 DPS) | yes | Cassandra's Grace (13102, -0.34 DPS) [world_drop]; Blood Guard's Satin Cover (220899, -0.38 DPS) [vendor]; Stormcloth Headband (10032, -0.95 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 24.4 healing_power points (2.23 DPS) | yes | Necklace of Calisea (1714, -0.11 DPS) [dungeon]; Choker of the High Shaman (4112, -0.11 DPS) [quest]; Amberglow Talisman (10824, -0.22 DPS) [quest] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | 48.1 healing_power points (4.39 DPS) | yes | Blood Guard's Satin Pads (220901, +0.00 DPS) [vendor]; Earthen Silk Shoulders (254033, -1.14 DPS) [crafted]; Hibernal Mantle (8111, -1.37 DPS) [world_drop] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | sim-verified (172.6 DPS) | yes | Cloak of Blight (6832, +0.00 DPS) [quest]; Arcane Cloak (8286, +0.00 DPS) [world_drop]; Wingveil Cloak (10802, +0.00 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ghostweave Vest (14141, +0.00 DPS) [crafted]; Grimlok's Tribal Vestments (9415, -1.85 DPS) [dungeon]; Stone Guard's Satin Armor (220903, -1.94 DPS) [vendor] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 38.8 healing_power points (3.54 DPS) | yes | Venomshroud Armguards (14439, -1.53 DPS) [world_drop]; Earthen Silk Cuffs (254019, -1.62 DPS) [crafted]; Shizzle's Nozzle Wiper (11917, -1.72 DPS) [quest] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 57.7 healing_power points (5.27 DPS) | yes | Earthenweave Gloves (254075, -0.69 DPS) [crafted]; Earthen Silk Gloves (254017, -2.02 DPS) [crafted]; First Sergeant's Satin Gloves (220898, -2.03 DPS) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | sim-verified (172.6 DPS) | yes | Gilded Cord (254037, +0.00 DPS) [crafted]; Earthenweave Cord (254077, +0.00 DPS) [crafted]; Gilded Waistcord (254081, +0.00 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | sim-verified (172.6 DPS) | yes | Senior Designer's Pantaloons (11841, +0.00 DPS) [dungeon]; Dalewind Trousers (13008, -0.06 DPS) [world_drop]; Stone Guard's Satin Leggings (220902, -1.03 DPS) [vendor] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | sim-verified (172.6 DPS) | yes | Furen's Boots (13100, +0.00 DPS) [world_drop]; Vinerot Sandals (17748, +0.00 DPS) [dungeon]; Earthenweave Boots (254093, +0.00 DPS) [crafted] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 33.1 healing_power points (3.01 DPS) | yes | Coldwater Ring (4550, -0.91 DPS) [quest]; Electrocutioner Lagnut (9447, -1.06 DPS) [dungeon]; Runemender's Seal (273031, -1.06 DPS) [world] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | sim-verified (172.6 DPS) | yes | Coldwater Ring (4550, +0.00 DPS) [quest]; Electrocutioner Lagnut (9447, +0.00 DPS) [dungeon]; Runemender's Seal (273031, +0.00 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (172.6 DPS) | yes | Evonice's Landin' Pilla (18951, -0.40 DPS) [quest]; Alchemist's Stone (13503, -1.60 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (172.6 DPS) | yes | Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Alchemist's Stone (13503, -0.84 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -6.01 DPS) [dungeon]; Hand of Righteousness (7721, -8.26 DPS) [dungeon]; Spire of Hakkar (10844, -9.57 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 13.3 healing_power points (1.21 DPS) | yes | Chillnail Splinter (10704, -0.01 DPS) [quest]; Captain Rackmore's Tiller (16789, -0.06 DPS) [quest]; Wand of Mana Concentration (273043, -0.08 DPS) [world] |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Nethergeld Shoulders; back: Imperial Red Cloak; wrist: Nethergeld Cuffs; hands: Gilded Gloves; waist: Dawnspire Cord; legs: Kilt of the Atal'ai Prophet; feet: Gilded Sandals; finger1: Eye of Adaegus; finger2: Welken Ring; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Jaina's Firestarter

No-known-source sample (15 of 507, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60 (troll, 025003031330000000-03505003030121431-00000000000000000)

Set DPS (verified): 301.3. Weights run: 7.2s. Verify run: 51.5s. 1193 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.078, intellect=0.945 ± 0.036, spirit=1.560 ± 0.033, mp5=2.714 ± 0.042, crit=0.188 ± 0.015 per rating point (14 rating = 1%, 2.637 per %), spell_haste=-18.305 ± 0.861

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Circlet (14140, +0.00 DPS) [crafted]; Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (301.2 DPS) | yes | Lady Alizabeth's Pendant (13002, +0.00 DPS) [world_drop]; Horizon Choker (13085, -0.19 DPS) [world_drop]; Necklace of Calisea (1714, -0.51 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Shoulders (14139, +0.00 DPS) [crafted]; Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp] |
| back | Wingveil Cloak (10802) | Sunken Temple: Spawn of Hakkar [dungeon] | 22.2 healing_power points (3.76 DPS) | yes | Cloak of Blight (6832, -0.32 DPS) [quest]; Indomitable Cloak (14683, -0.32 DPS) [world_drop]; Glowing Thresher Cape (6901, -0.46 DPS) [dungeon] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 98.0 healing_power points (16.64 DPS) | yes | Warlord's Satin Tunic (231632, -2.80 DPS) [vendor]; Legionnaire's Satin Tunic (227124, -3.73 DPS) [pvp]; Mooncloth Vest (14138, -3.98 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nethergeld Cuffs (254061, +0.00 DPS) [crafted]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted]; General's Satin Bracers (17619, -0.10 DPS) [pvp] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | 51.0 healing_power points (8.66 DPS) | yes | Gilded Gloves (254095, -0.01 DPS) [crafted]; Swarmtender's Gloves (275607, -0.64 DPS) [crafted]; Blood Guard's Satin Handwraps (227126, -1.06 DPS) [pvp] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 58.3 healing_power points (9.90 DPS) | yes | General's Satin Cinch (17621, -2.23 DPS) [pvp]; Devout Belt (16696, -2.43 DPS) [dungeon]; Earthenweave Cord (254077, -2.57 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 74.0 healing_power points (12.58 DPS) | yes | Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Legguards (231634, -0.20 DPS) [vendor]; Mooncloth Leggings (14137, -1.75 DPS) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; General's Satin Walkers (231630, -0.16 DPS) [vendor]; Earthenweave Boots (254093, -0.65 DPS) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 26.6 healing_power points (4.52 DPS) | yes | Electrocutioner Lagnut (9447, -1.65 DPS) [dungeon]; Runemender's Seal (273031, -1.65 DPS) [world]; Coldwater Ring (4550, -1.65 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 24.4 healing_power points (4.14 DPS) | yes | Electrocutioner Lagnut (9447, -1.27 DPS) [dungeon]; Runemender's Seal (273031, -1.27 DPS) [world]; Coldwater Ring (4550, -1.28 DPS) [quest] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.12 DPS) [quest]; Alchemist's Stone (13503, -1.71 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (301.2 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.41 DPS) [quest]; Alchemist's Stone (13503, -2.00 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -10.13 DPS) [dungeon]; Hand of Edward the Odd (2243, -12.50 DPS) [world_drop]; Hand of Righteousness (7721, -14.43 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 34.2 healing_power points (5.80 DPS) | yes | Captain Rackmore's Tiller (16789, -3.99 DPS) [quest]; Jaina's Firestarter (13064, -4.04 DPS) [world_drop]; Chillnail Splinter (10704, -4.21 DPS) [quest] |

**New at 60:** head: Devout Crown; neck: Glowing Eye of Mordresh; shoulder: Devout Mantle; back: Wingveil Cloak; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Shard of the Splithooves; ranged: Torch of Light

No-known-source sample (15 of 1193, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60, raid preset (troll, 025003031330000000-03505003030121431-00000000000000000)

Set DPS (verified): 631.0. Weights run: 4.8s. Verify run: 32.0s. 1193 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.156, intellect=0.665 ± 0.029, spirit=0.866 ± 0.028, mp5=1.700 ± 0.032, crit=0.302 ± 0.016 per rating point (14 rating = 1%, 4.233 per %), spell_haste=-2.810 ± 0.559

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (631.1 DPS) | yes | Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -6.99 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 18.3 healing_power points (5.91 DPS) | yes | Horizon Choker (13085, -1.79 DPS) [world_drop]; Necklace of Calisea (1714, -2.45 DPS) [dungeon]; Lady Alizabeth's Pendant (13002, -3.37 DPS, sim-verified) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (631.1 DPS) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp]; Mooncloth Shoulders (14139, -6.77 DPS, sim-verified) [crafted] |
| back | Wingveil Cloak (10802) | Sunken Temple: Spawn of Hakkar [dungeon] | 14.5 healing_power points (4.69 DPS) | yes | Glowing Thresher Cape (6901, -0.19 DPS) [dungeon]; Cloak of Blight (6832, -1.05 DPS) [quest]; Indomitable Cloak (14683, -1.05 DPS) [world_drop] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 86.9 healing_power points (28.03 DPS) | yes | Warlord's Satin Tunic (231632, -6.27 DPS) [vendor]; Legionnaire's Satin Tunic (227124, -7.99 DPS) [pvp]; Stormcloth Vest (10020, -14.64 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (631.1 DPS) | yes | Nethergeld Cuffs (254061, +0.00 DPS) [crafted]; General's Satin Bracers (17619, -1.19 DPS) [pvp]; Earthenweave Cuffs (254131, -10.01 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (631.1 DPS) | yes | Blood Guard's Satin Handwraps (227126, +0.00 DPS) [pvp]; General's Satin Handwraps (231633, +0.00 DPS) [vendor]; Gilded Gloves (254095, -9.12 DPS, sim-verified) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 48.1 healing_power points (15.52 DPS) | yes | Gilded Waistcord (254081, -4.23 DPS) [crafted]; Gilded Cord (254037, -5.03 DPS) [crafted]; Earthenweave Cord (254077, -9.91 DPS, sim-verified) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 53.9 healing_power points (17.39 DPS) | yes | Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Earthenweave Leggings (254125, -15.13 DPS, sim-verified) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (631.1 DPS) | yes | General's Satin Walkers (231630, +0.00 DPS) [vendor]; Swarmtender's Footpads (275609, +0.00 DPS) [crafted]; Mooncloth Boots (15802, -4.69 DPS, sim-verified) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 16.2 healing_power points (5.22 DPS) | yes | Electrocutioner Lagnut (9447, -1.33 DPS) [dungeon]; Runemender's Seal (273031, -1.33 DPS) [world]; Solid Gold Band (279845, -1.56 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 14.4 healing_power points (4.64 DPS) | yes | Runemender's Seal (273031, -0.75 DPS) [world]; Solid Gold Band (279845, -0.99 DPS) [quest]; Electrocutioner Lagnut (9447, -4.61 DPS, sim-verified) [dungeon] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (631.1 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.49 DPS) [quest]; Alchemist's Stone (13503, -2.17 DPS) [crafted] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (631.1 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop]; Evonice's Landin' Pilla (18951, -0.56 DPS) [quest]; Alchemist's Stone (13503, -2.24 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (631.1 DPS) | yes | Charstone Dirk (17710, -16.25 DPS) [dungeon]; Hand of Righteousness (7721, -24.57 DPS) [dungeon]; Hand of Edward the Odd (2243, -40.42 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 26.5 healing_power points (8.56 DPS) | yes | Jaina's Firestarter (13064, -6.43 DPS) [world_drop]; Wand of Mana Concentration (273043, -6.61 DPS) [world]; Captain Rackmore's Tiller (16789, -12.00 DPS, sim-verified) [quest] |

**New at 60:** head: Devout Crown; neck: Glowing Eye of Mordresh; shoulder: Devout Mantle; back: Wingveil Cloak; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Shard of the Splithooves; trinket2: Ankh of Life; ranged: Torch of Light

No-known-source sample (15 of 1193, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

