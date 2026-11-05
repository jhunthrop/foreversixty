# Leveling BiS: Discipline

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 005005000010000000-00000000000000000-00000000000000000)

Set DPS (verified): 43.2. Weights run: 12.4s. Verify run: 33.9s. 183 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.119, intellect=0.830 ± 0.010, spirit=0.953 ± 0.011, mp5=2.535 ± 0.011, crit=0.098 ± 0.008 per rating point (14 rating = 1%, 1.379 per %), spell_haste=not significant (-0.026 ± 0.143)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (2.18 DPS) | yes | Shadow Goggles (4373, -0.96 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.72 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 1.9 healing_power points (0.23 DPS) | yes | Tarnished Locket (279870, -0.33 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.5 healing_power points (1.51 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.54 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.59 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -1.03 DPS) [crafted] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | sim-verified (43.2 DPS) | yes | Heavy Woolen Cloak (4311, -0.42 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -0.78 DPS) [crafted]; Spirit Cloak (4792, -0.86 DPS) [vendor] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 20.0 healing_power points (2.42 DPS) | yes | Robe of the Moccasin (6465, -0.18 DPS) [dungeon]; Corsair's Overshirt (5202, -0.23 DPS, sim-verified) [dungeon]; Scholarly Robes (2034, -0.68 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 11.8 healing_power points (1.42 DPS) | yes | Bright Bracers (3647, -1.02 DPS) [world_drop]; Seer's Cuffs (3645, -1.09 DPS) [dungeon]; Repurposed Hair Band (281256, -1.33 DPS, sim-verified) [quest] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 13.5 healing_power points (1.63 DPS) | yes | Magefist Gloves (12977, -0.64 DPS, sim-verified) [world_drop]; Bright Gloves (3066, -0.77 DPS) [world_drop]; Serpent Gloves (5970, -0.79 DPS) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 14.3 healing_power points (1.74 DPS) | yes | Novice Ardent's Sash (253887, -0.41 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.93 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.95 DPS) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 26.8 healing_power points (3.25 DPS) | yes | Darkweave Breeches (12987, -1.75 DPS, sim-verified) [world_drop]; Colorful Kilt (10048, -2.06 DPS) [crafted]; Filigreed Silky Leggings (253939, -2.18 DPS) [crafted] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 11.5 healing_power points (1.39 DPS) | yes | Spidersilk Boots (4320, -0.17 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.33 DPS) [crafted]; Kimbra Boots (6191, -0.64 DPS) [quest] |
| finger1 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 10.1 healing_power points (1.22 DPS) | yes | Band of Purification (12996, -0.53 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.62 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.92 DPS) [world_drop] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.7 healing_power points (0.81 DPS) | yes | Band of Purification (12996, -0.16 DPS, sim-verified) [world_drop]; Lavishly Jeweled Ring (1156, -0.20 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.51 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Westfall (2042) | The Defias Brotherhood [quest] | 57.9 healing_power points (7.01 DPS) | yes | Living Root (6631, -1.01 DPS, sim-verified) [dungeon]; Staff of the Blessed Seer (2271, -2.46 DPS) [dungeon]; Evocator's Blade (2567, -2.84 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Sizzle Stick (8071) (or Greater Magic Wand (11288), Torchlight Wand (5240)) | Deviate Eradication [quest] | 2.0 healing_power points (0.24 DPS) | yes | Torchlight Wand (5240, +0.00 DPS) [quest]; Greater Magic Wand (11288, +0.00 DPS) [crafted]; Sable Wand (7607, -0.01 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Staff of Westfall; ranged: Sizzle Stick

No-known-source sample (15 of 183, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes; 16607 Acolyte's Sacrificial Robes; 18851 Insignia of the Horde; 20434 Lorekeeper's Staff

### Band 30 (human, 005005031330001000-00000000000000000-00000000000000000)

Set DPS (verified): 93.0. Weights run: 15.1s. Verify run: 79.1s. 302 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.033, intellect=0.641 ± 0.008, spirit=0.833 ± 0.008, mp5=2.285 ± 0.009, crit=0.082 ± 0.003 per rating point (14 rating = 1%, 1.146 per %), spell_haste=not significant (0.059 ± 0.061)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 38.0 healing_power points (5.78 DPS) | yes | Filigreed Pristine Circlet (253975, -2.83 DPS, sim-verified) [crafted]; Pristine Circlet (253949, -3.04 DPS) [crafted]; Enchanter's Cowl (4322, -3.89 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 5.9 healing_power points (0.90 DPS) | yes | Pendant of Myzrael (4614, -0.14 DPS) [dungeon]; Glowing Green Talisman (5002, -0.14 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.17 DPS, sim-verified) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.8 healing_power points (2.25 DPS) | yes | Death Speaker Mantle (6685, -0.11 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.61 DPS) [world_drop]; Mantle of Honor (3560, -0.68 DPS) [quest] |
| back | Caretaker's Cape (19533) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Thresher Cape (6901, -0.12 DPS, sim-verified) [dungeon]; Repairman's Cape (9605, -0.65 DPS) [quest]; Prelacy Cape (7004, -1.04 DPS) [quest] |
| chest | Pristine Gown (253961) | Tailoring [crafted] | 27.8 healing_power points (4.23 DPS) | yes | Death Speaker Robes (6682, -0.15 DPS, sim-verified) [dungeon]; Beguiler Robes (7728, -0.68 DPS) [dungeon]; Robes of Arugal (6324, -0.73 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 15.9 healing_power points (2.41 DPS) | yes | Mindthrust Bracers (1974, -0.76 DPS, sim-verified) [dungeon]; Nightsky Wristbands (6407, -1.45 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.63 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | sim-verified (93.0 DPS) | yes | Pristine Gloves (253913, -0.04 DPS) [crafted]; Truefaith Gloves (7049, -0.19 DPS, sim-verified) [crafted]; Gloves of Meditation (4318, -0.36 DPS) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 15.1 healing_power points (2.30 DPS) | yes | Highlander's Cloth Girdle (20099, -0.33 DPS) [rep]; Pristine Sash (253925, -0.44 DPS, sim-verified) [crafted]; Crimson Silk Belt (7055, -0.70 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 30.6 healing_power points (4.67 DPS) | yes | Filigreed Pristine Leggings (253937, -0.79 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -0.99 DPS) [crafted]; Gaze Dreamer Pants (6903, -1.70 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 27.8 healing_power points (4.23 DPS) | yes | Acidic Walkers (9454, -2.13 DPS, sim-verified) [dungeon]; Glinteye Slippers (273024, -2.41 DPS) [dungeon]; Boots of the Enchanter (4325, -2.46 DPS) [crafted] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 14.8 healing_power points (2.26 DPS) | yes | Ring of Power Regulation (273030, -0.30 DPS) [dungeon]; Lorekeeper's Ring (19525, -0.50 DPS) [rep] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 14.8 healing_power points (2.26 DPS) | yes | Ring of Power Regulation (273030, -0.37 DPS, sim-verified) [dungeon]; Lorekeeper's Ring (19525, -0.50 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Odo's Ley Staff (6318, -0.91 DPS) [dungeon]; Staff of Westfall (2042, -1.80 DPS) [quest]; Staff of the Friar (3415, -2.98 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 18.0 healing_power points (2.74 DPS) | yes | Orb of Mistmantle (13031, -0.86 DPS) [world_drop]; Orb of Mystic Insight (249394, -1.09 DPS) [crafted]; Eye of Paleth (2943, -1.11 DPS, sim-verified) [quest] |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 7.1 healing_power points (1.09 DPS) | yes | Wand of Eventide (5214, -0.32 DPS) [world_drop]; Greater Mystic Wand (217287, -0.32 DPS) [crafted]; Thunderwood (13062, -0.53 DPS, sim-verified) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Caretaker's Cape; chest: Pristine Gown; wrist: Spidertank Oilrag; hands: Gloves of Old; waist: Belt of Arugal; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Death Speaker Scepter; off_hand: Orb of Souls; ranged: Wand of Mana Concentration

No-known-source sample (15 of 302, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 40 (human, 005005031330011430-20000000000000000-00000000000000000)

Set DPS (verified): 155.9. Weights run: 15.4s. Verify run: 89.1s. 389 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.188, intellect=1.157 ± 0.028, spirit=1.530 ± 0.028, mp5=3.678 ± 0.029, crit=0.377 ± 0.023 per rating point (14 rating = 1%, 5.279 per %), spell_haste=1.581 ± 0.384

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 55.4 healing_power points (8.57 DPS) | yes | Whitemane's Chapeau (7720, -0.73 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -2.05 DPS) [world_drop]; Earthen Silk Hood (254015, -2.81 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 29.3 healing_power points (4.53 DPS) | yes | Triune Amulet (7722, -1.63 DPS) [dungeon]; Necklace of Calisea (1714, -1.71 DPS, sim-verified) [dungeon]; Amberglow Talisman (10824, -2.16 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 37.0 healing_power points (5.73 DPS) | yes | Earthen Silk Shoulders (254033, -1.05 DPS, sim-verified) [crafted]; Mantle of Doan (7712, -1.68 DPS) [dungeon]; Green Silken Shoulders (7057, -2.52 DPS) [crafted] |
| back | Caretaker's Cape (19532) | Silverwing Sentinels [rep] | sim-verified (155.9 DPS) | yes | Glowing Thresher Cape (6901, -0.83 DPS, sim-verified) [dungeon]; Expeditionary's Cape (271922, -1.05 DPS) [rep]; Repairman's Cape (9605, -1.32 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 70.2 healing_power points (10.86 DPS) | yes | Dreamweave Vest (10021, -3.58 DPS, sim-verified) [crafted]; Civinad Robes (9623, -4.00 DPS) [quest]; Doomsayer's Robe (4746, -4.13 DPS) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 20.0 healing_power points (3.10 DPS) | yes | Earthen Silk Cuffs (254019, -0.33 DPS, sim-verified) [crafted]; Mindthrust Bracers (1974, -0.50 DPS) [dungeon]; Silkstream Cuffs (16791, -1.26 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 33.3 healing_power points (5.16 DPS) | yes | Earthen Silk Gloves (254017, -0.38 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -0.50 DPS) [crafted]; Stormcloth Gloves (10011, -1.31 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 40.4 healing_power points (6.25 DPS) | yes | Deathmage Sash (10771, -2.61 DPS, sim-verified) [dungeon]; Windchaser Cinch (14435, -2.69 DPS) [world_drop]; Sutarn's Ring (13105, -3.04 DPS) [world_drop] |
| legs | Stoneweaver Leggings (9407) | Uldaman: Ironaya [dungeon] | 47.6 healing_power points (7.36 DPS) | yes | Stormcloth Pants (10010, -1.40 DPS, sim-verified) [crafted]; Pristine Leggings (253987, -1.53 DPS) [crafted]; Earthen Leggings (253999, -2.44 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 34.2 healing_power points (5.29 DPS) | yes | Thoughtcast Boots (10578, -0.20 DPS, sim-verified) [dungeon]; Furen's Boots (13100, -0.25 DPS) [world_drop]; Earthen Silk Slippers (254013, -0.85 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 20.0 healing_power points (3.10 DPS) | yes | Runemender's Seal (273031, -0.05 DPS) [world]; Ring of Power Regulation (273030, -0.46 DPS) [dungeon]; Welken Ring (5011, -0.55 DPS) [world_drop] |
| finger2 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 19.7 healing_power points (3.05 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; Ring of Power Regulation (273030, -0.41 DPS) [dungeon]; Welken Ring (5011, -0.50 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (155.9 DPS) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (155.9 DPS) | yes | Staff of the Friar (3415, -5.01 DPS) [dungeon]; Odo's Ley Staff (6318, -5.55 DPS) [dungeon]; Staff of Jordan (873, -6.52 DPS, sim-verified) [world_drop] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | 28.1 healing_power points (4.35 DPS) | yes | Orb of Lorica (11262, -0.71 DPS, sim-verified) [quest]; Orb of Souls (249395, -0.92 DPS) [crafted]; Orb of Mistmantle (13031, -2.01 DPS) [world_drop] |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 12.0 healing_power points (1.85 DPS) | yes | Jaina's Firestarter (13064, -0.07 DPS) [world_drop]; Captain Rackmore's Tiller (16789, -0.22 DPS) [quest]; Goblin Igniter (5253, -0.43 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Caretaker's Cape; chest: Stormcloth Vest; hands: Dreamweave Gloves; waist: Gilded Cord; legs: Stoneweaver Leggings; finger1: Lorekeeper's Ring; finger2: Electrocutioner Lagnut; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Hand of Righteousness; off_hand: Beacon of Hope

No-known-source sample (15 of 389, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 50 (human, 005005031330011430-33501000000000000-00000000000000000)

Set DPS (verified): 235.7. Weights run: 13.4s. Verify run: 79.1s. 496 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.112, intellect=0.745 ± 0.022, spirit=1.123 ± 0.018, mp5=3.192 ± 0.022, crit=0.327 ± 0.016 per rating point (14 rating = 1%, 4.581 per %), spell_haste=not significant (0.697 ± 0.263)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Cassandra's Grace (13102) | World drop [world_drop] | 54.1 healing_power points (14.59 DPS) | yes | Knight-Lieutenant's Satin Cover (220896, +0.00 DPS) [vendor]; Papal Fez (9431, -2.29 DPS, sim-verified) [dungeon]; Stormcloth Headband (10032, -3.08 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 26.0 healing_power points (7.00 DPS) | yes | Mindburst Medallion (11196, -2.34 DPS, sim-verified) [quest]; Triune Amulet (7722, -2.72 DPS) [dungeon]; Horizon Choker (13085, -2.98 DPS) [world_drop] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | 38.7 healing_power points (10.44 DPS) | yes | Knight-Lieutenant's Satin Pads (220894, +0.00 DPS) [vendor]; Inquisitor's Shawl (19507, -2.07 DPS, sim-verified) [dungeon]; Earthen Silk Shoulders (254033, -3.16 DPS) [crafted] |
| back | Caretaker's Cape (19531) | Silverwing Sentinels [rep] | sim-verified (235.7 DPS) | yes | Spritecaster Cape (11623, -1.19 DPS, sim-verified) [dungeon]; Wingveil Cloak (10802, -3.07 DPS) [dungeon]; Expeditionary's Cape (271921, -3.42 DPS) [rep] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 67.0 healing_power points (18.06 DPS) | yes | Ghostweave Vest (14141, -1.56 DPS, sim-verified) [crafted]; Grimlok's Tribal Vestments (9415, -5.00 DPS) [dungeon]; Knight's Satin Armor (220892, -5.25 DPS) [vendor] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 30.8 healing_power points (8.32 DPS) | yes | Spidertank Oilrag (9448, -3.24 DPS, sim-verified) [dungeon]; Earthen Silk Cuffs (254019, -4.10 DPS) [crafted]; Bloodband Bracers (11469, -4.25 DPS) [quest] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 45.4 healing_power points (12.25 DPS) | yes | Earthenweave Gloves (254075, -1.80 DPS, sim-verified) [crafted]; Sergeant Major's Satin Gloves (220897, -3.09 DPS) [vendor]; Dreamweave Gloves (10019, -4.48 DPS) [crafted] |
| waist | Earthenweave Cord (254077) | Tailoring [crafted] | 38.4 healing_power points (10.34 DPS) | yes | Gilded Waistcord (254081, -0.91 DPS, sim-verified) [crafted]; Gilded Cord (254037, -0.99 DPS) [crafted]; Ghostweave Cord (254073, -3.24 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 50.4 healing_power points (13.60 DPS) | yes | Knight's Satin Leggings (220893, -1.03 DPS, sim-verified) [vendor]; Kilt of the Atal'ai Prophet (10807, -2.91 DPS) [dungeon]; Stoneweaver Leggings (9407, -2.97 DPS) [dungeon] |
| feet | Earthenweave Boots (254093) | Tailoring [crafted] | 38.4 healing_power points (10.34 DPS) | yes | Gilded Sandals (254107, -0.37 DPS, sim-verified) [crafted]; Vinerot Sandals (17748, -2.27 DPS) [dungeon]; Gilded Slippers (254001, -2.33 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 24.8 healing_power points (6.68 DPS) | yes | Electrocutioner Lagnut (9447, -2.13 DPS) [dungeon]; Runemender's Seal (273031, -2.13 DPS) [world]; Ring of Power Regulation (273030, -2.48 DPS) [dungeon] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 18.0 healing_power points (4.84 DPS) | yes | Runemender's Seal (273031, -0.29 DPS) [world]; Electrocutioner Lagnut (9447, -0.62 DPS, sim-verified) [dungeon]; Ring of Power Regulation (273030, -0.64 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (235.7 DPS) | yes | Evonice's Landin' Pilla (18951, -0.61 DPS) [quest]; Thunderbrew's Boot Flask (744, -1.21 DPS) [quest]; Uther's Strength (11302, -1.56 DPS, sim-verified) [world_drop] |
| trinket2 | Enriched Seal (271903) | Theramore Expeditionary Force [rep] | sim-verified (235.7 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.27 DPS) [quest]; Thunderbrew's Boot Flask (744, -1.88 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (235.7 DPS) | yes | Resurgence Rod (17743, -8.34 DPS, sim-verified) [quest]; Charstone Dirk (17710, -13.79 DPS) [dungeon]; Dreamstaff (249454, -14.94 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Captain Rackmore's Tiller (16789) | Claim Rackmore's Treasure! [quest] | 9.4 healing_power points (2.53 DPS) | yes | Wand of Mana Concentration (273043, -0.00 DPS) [world]; Cairnstone Sliver (9654, -0.01 DPS) [quest]; Earthen Rod (9381, -0.10 DPS) [dungeon] |

**New at 50:** head: Cassandra's Grace; shoulder: Nethergeld Shoulders; back: Caretaker's Cape; wrist: Nethergeld Cuffs; hands: Gilded Gloves; waist: Earthenweave Cord; legs: Senior Designer's Pantaloons; feet: Earthenweave Boots; finger1: Lorekeeper's Ring; finger2: Eye of Adaegus; trinket2: Enriched Seal; main_hand: Glowing Brightwood Staff; ranged: Captain Rackmore's Tiller

No-known-source sample (15 of 496, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60 (human, 005005031330011430-33505003030000000-00000000000000000)

Set DPS (verified): 348.1. Weights run: 6.9s. Verify run: 113.8s. 1107 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.082, intellect=0.746 ± 0.024, spirit=0.922 ± 0.018, mp5=2.629 ± 0.024, crit=0.438 ± 0.015 per rating point (14 rating = 1%, 6.133 per %), spell_haste=1.977 ± 0.302

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mooncloth Circlet (14140) | Tailoring [crafted] | 64.8 healing_power points (21.00 DPS) | yes | Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Earthenweave Crown (254137, -2.97 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 26.0 healing_power points (8.41 DPS) | yes | Archlight Talisman (15856, -2.46 DPS, sim-verified) [quest]; Stormpike Sage's Pendant (19098, -2.59 DPS) [rep]; Mindburst Medallion (11196, -3.52 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (348.1 DPS) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp]; Mooncloth Shoulders (14139, -16.78 DPS, sim-verified) [crafted] |
| back | Caretaker's Cape (19530) | Silverwing Sentinels [rep] | sim-verified (348.1 DPS) | yes | Spritecaster Cape (11623, -3.04 DPS) [dungeon]; Expeditionary's Cape (271907, -5.15 DPS) [rep]; Wingveil Cloak (10802, -5.91 DPS) [dungeon] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 87.8 healing_power points (28.43 DPS) | yes | Field Marshal's Satin Tunic (231624, +0.00 DPS) [vendor]; Stormcloth Vest (10020, -6.86 DPS, sim-verified) [crafted]; Field Marshal's Satin Vestments (17605, -7.57 DPS) [vendor] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (348.1 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [rep]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (348.1 DPS) | yes | Marshal's Satin Handwraps (231623, +0.00 DPS) [vendor]; Theramore Handwraps (271945, +0.00 DPS) [rep]; Gilded Gloves (254095, -20.15 DPS, sim-verified) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 49.7 healing_power points (16.11 DPS) | yes | Earthenweave Cord (254077, -3.98 DPS, sim-verified) [crafted]; Gilded Waistcord (254081, -4.54 DPS) [crafted]; Devout Belt (16696, -5.03 DPS) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 56.4 healing_power points (18.27 DPS) | yes | Knight-Captain's Satin Legguards (227125, +0.00 DPS) [pvp]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Earthenweave Leggings (254125, -22.85 DPS, sim-verified) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (348.1 DPS) | yes | Marshal's Satin Walkers (231627, +0.00 DPS) [vendor]; Swarmtender's Footpads (275609, +0.00 DPS) [crafted]; Mooncloth Boots (15802, -17.78 DPS, sim-verified) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 28.8 healing_power points (9.31 DPS) | yes | Deep Rooted Ring (19109, -3.01 DPS) [world]; Band of the Hierophant (13096, -3.61 DPS) [world_drop]; Eye of Adaegus (5266, -4.28 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (19522) | Silverwing Sentinels [rep] | 24.5 healing_power points (7.94 DPS) | yes | Deep Rooted Ring (19109, -1.64 DPS) [world]; Band of the Hierophant (13096, -2.24 DPS) [world_drop]; Eye of Adaegus (5266, -2.91 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (348.1 DPS) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Enriched Seal (271903, +0.00 DPS) [rep]; Evonice's Landin' Pilla (18951, -0.60 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (348.1 DPS) | yes | Enriched Seal (271903, -1.57 DPS) [rep]; Shard of the Splithooves (10659, -1.73 DPS, sim-verified) [quest]; Evonice's Landin' Pilla (18951, -2.84 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (348.1 DPS) | yes | Grand Marshal's Warhammer (23454, +0.00 DPS) [vendor]; Hand of Edward the Odd (2243, -6.99 DPS, sim-verified) [world_drop]; Resurgence Rod (17743, -12.64 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 27.1 healing_power points (8.79 DPS) | yes | Brilliant Wand (249385, -3.85 DPS, sim-verified) [crafted]; Earthen Rod (9381, -5.88 DPS) [dungeon]; Greater Eternal Wand (249237, -5.88 DPS) [crafted] |

**New at 60:** head: Mooncloth Circlet; shoulder: Devout Mantle; back: Caretaker's Cape; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Devout Sandals; finger1: Maiden's Circle; finger2: Lorekeeper's Ring; trinket2: Uther's Strength; ranged: Torch of Light

No-known-source sample (15 of 1107, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60, raid preset (human, 005005031330011430-33505003030000000-00000000000000000)

Set DPS (verified): 611.8. Weights run: 5.2s. Verify run: 76.2s. 1107 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.110, intellect=0.622 ± 0.035, spirit=0.439 ± 0.032, mp5=1.173 ± 0.037, crit=0.606 ± 0.027 per rating point (14 rating = 1%, 8.489 per %), spell_haste=3.121 ± 0.665

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mooncloth Circlet (14140) | Tailoring [crafted] | 57.6 healing_power points (26.71 DPS) | yes | Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Lieutenant Commander's Satin Hood (227121, -0.11 DPS) [pvp]; Cassandra's Grace (13102, -5.10 DPS, sim-verified) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 25.0 healing_power points (11.59 DPS) | yes | Archlight Talisman (15856, -3.04 DPS, sim-verified) [quest]; Mindburst Medallion (11196, -6.05 DPS) [quest]; Triune Amulet (7722, -6.45 DPS) [dungeon] |
| shoulder | Mooncloth Shoulders (14139) | Tailoring [crafted] | 41.8 healing_power points (19.40 DPS) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp]; Nethergeld Shoulders (254049, -2.89 DPS, sim-verified) [crafted] |
| back | Caretaker's Cape (19530) | Silverwing Sentinels [rep] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Spritecaster Cape (11623, -4.24 DPS) [dungeon]; Stormpike Sage's Cloak (19086, -7.19 DPS) [rep]; Runecloth Cloak (13860, -7.20 DPS) [crafted] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 80.0 healing_power points (37.12 DPS) | yes | Field Marshal's Satin Tunic (231624, +0.00 DPS) [vendor]; Stormcloth Vest (10020, -9.04 DPS, sim-verified) [crafted]; Knight-Captain's Satin Tunic (227122, -11.39 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 30.0 healing_power points (13.94 DPS) | yes | Earthenweave Cuffs (254131, -1.70 DPS) [crafted]; Nethergeld Cuffs (254061, -2.14 DPS, sim-verified) [crafted]; Devout Bracers (16697, -4.49 DPS) [dungeon] |
| hands | Theramore Handwraps (271945) | Theramore Expeditionary Force [rep] | 40.8 healing_power points (18.92 DPS) | yes | Knight-Lieutenant's Satin Handwraps (227128, +0.00 DPS) [pvp]; Marshal's Satin Handwraps (231623, +0.00 DPS) [vendor]; Gilded Gloves (254095, -1.61 DPS, sim-verified) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 43.7 healing_power points (20.25 DPS) | yes | Gilded Waistcord (254081, -4.99 DPS, sim-verified) [crafted]; Earthenweave Cord (254077, -5.95 DPS) [crafted]; Gilded Cord (254037, -6.52 DPS) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 55.2 healing_power points (25.60 DPS) | yes | Knight-Captain's Satin Legguards (227125, +0.00 DPS) [pvp]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Marshal's Satin Pants (17603, -2.52 DPS) [vendor] |
| feet | Swarmtender's Footpads (275609) | Tailoring [crafted] | 43.0 healing_power points (19.95 DPS) | yes | Marshal's Satin Walkers (231627, +0.00 DPS) [vendor]; Mooncloth Boots (15802, -1.02 DPS, sim-verified) [crafted]; Gilded Sandals (254107, -2.97 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 25.0 healing_power points (11.59 DPS) | yes | Deep Rooted Ring (19109, -3.14 DPS) [world]; Philanthropist's Ring (281635, -5.51 DPS) [quest]; Band of the Unicorn (7553, -5.56 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (19522) | Silverwing Sentinels [rep] | 18.7 healing_power points (8.67 DPS) | yes | Deep Rooted Ring (19109, -0.61 DPS, sim-verified) [world]; Philanthropist's Ring (281635, -2.59 DPS) [quest]; Band of the Unicorn (7553, -2.64 DPS) [world_drop] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (611.7 DPS) | yes | Ankh of Life (1713, -1.28 DPS, sim-verified) [world_drop]; Enriched Seal (271903, -5.63 DPS) [rep]; Evonice's Landin' Pilla (18951, -6.31 DPS) [quest] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (-3.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Enriched Seal (271903, -0.54 DPS) [rep]; Ankh of Life (1713, -0.82 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.23 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Warhammer (23454, +0.00 DPS) [vendor]; Charstone Dirk (17710, -2.53 DPS, sim-verified) [dungeon]; Resurgence Rod (17743, -21.84 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 21.8 healing_power points (10.12 DPS) | yes | Brilliant Wand (249385, -4.37 DPS, sim-verified) [crafted]; Earthen Rod (9381, -5.95 DPS) [dungeon]; Greater Eternal Wand (249237, -5.95 DPS) [crafted] |

**New at 60:** head: Mooncloth Circlet; shoulder: Mooncloth Shoulders; back: Caretaker's Cape; chest: Truefaith Vestments; wrist: Dryad's Wrist Bindings; hands: Theramore Handwraps; waist: Wisdom of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Swarmtender's Footpads; finger1: Maiden's Circle; finger2: Lorekeeper's Ring; trinket1: Uther's Strength; trinket2: Shard of the Splithooves; ranged: Torch of Light

No-known-source sample (15 of 1107, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

## Horde

### Band 20 (undead, 005005000010000000-00000000000000000-00000000000000000)

Set DPS (verified): 41.7. Weights run: 12.4s. Verify run: 36.2s. 176 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.119, intellect=0.830 ± 0.010, spirit=0.953 ± 0.011, mp5=2.535 ± 0.011, crit=0.098 ± 0.008 per rating point (14 rating = 1%, 1.379 per %), spell_haste=not significant (-0.026 ± 0.143)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (2.18 DPS) | yes | Shadow Goggles (4373, -1.31 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.72 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 2.9 healing_power points (0.35 DPS) | yes | Scholarly Pendant (277203, -0.12 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.5 healing_power points (1.51 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.46 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.59 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -1.03 DPS) [crafted] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | sim-verified (41.7 DPS) | yes | Heavy Woolen Cloak (4311, -0.48 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -0.78 DPS) [crafted]; Spirit Cloak (4792, -0.86 DPS) [vendor] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 20.0 healing_power points (2.42 DPS) | yes | Corsair's Overshirt (5202, -0.15 DPS, sim-verified) [dungeon]; Robe of the Moccasin (6465, -0.18 DPS) [dungeon]; Scholarly Robes (2034, -0.68 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 11.8 healing_power points (1.42 DPS) | yes | Tabitha's Cuffs (251486, -1.06 DPS) [quest]; Crystalline Cuffs (14148, -1.08 DPS) [dungeon]; Bright Bracers (3647, -1.27 DPS, sim-verified) [world_drop] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 13.5 healing_power points (1.63 DPS) | yes | Magefist Gloves (12977, -0.68 DPS, sim-verified) [world_drop]; Bright Gloves (3066, -0.77 DPS) [world_drop]; Serpent Gloves (5970, -0.79 DPS) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 14.3 healing_power points (1.74 DPS) | yes | Novice Ardent's Sash (253887, -0.55 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.93 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.95 DPS) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 26.8 healing_power points (3.25 DPS) | yes | Darkweave Breeches (12987, -1.67 DPS, sim-verified) [world_drop]; Colorful Kilt (10048, -2.06 DPS) [crafted]; Filigreed Silky Leggings (253939, -2.18 DPS) [crafted] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 11.5 healing_power points (1.39 DPS) | yes | Spidersilk Boots (4320, -0.20 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.33 DPS) [crafted]; Smoldering Boots (3076, -0.70 DPS) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 10.1 healing_power points (1.22 DPS) | yes | Lavishly Jeweled Ring (1156, -0.62 DPS) [dungeon]; Loop of Sacrifice (281673, -0.72 DPS) [quest]; Volcanic Rock Ring (12053, -0.92 DPS) [world_drop] |
| finger2 | Band of Purification (12996) | World drop [world_drop] | 5.7 healing_power points (0.69 DPS) | yes | Lavishly Jeweled Ring (1156, -0.18 DPS, sim-verified) [dungeon]; Loop of Sacrifice (281673, -0.19 DPS) [quest]; Volcanic Rock Ring (12053, -0.39 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Living Root (6631) | Wailing Caverns: Verdan the Everliving [dungeon] | 49.8 healing_power points (6.03 DPS) | yes | Staff of the Blessed Seer (2271, -1.08 DPS, sim-verified) [dungeon]; Gnarled Necromancer's Staff (251534, -1.73 DPS) [quest]; Evocator's Blade (2567, -1.86 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Sizzle Stick (8071) (or Greater Magic Wand (11288), Torchlight Wand (5240)) | Deviate Eradication [quest] | 2.0 healing_power points (0.24 DPS) | yes | Torchlight Wand (5240, +0.00 DPS) [quest]; Greater Magic Wand (11288, +0.00 DPS) [crafted]; Wand of Decay (5252, -0.04 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Advisor's Ring; finger2: Band of Purification; main_hand: Living Root; ranged: Sizzle Stick

No-known-source sample (15 of 176, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20441 Scout's Blade; 23173 Abomination Skin Leggings; 209613 Insignia of the Alliance; 241089 Scarlet Dagger; 248007 Militia Shortblade

### Band 30 (undead, 005005031330001000-00000000000000000-00000000000000000)

Set DPS (verified): 93.9. Weights run: 15.1s. Verify run: 80.5s. 295 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.033, intellect=0.641 ± 0.008, spirit=0.833 ± 0.008, mp5=2.285 ± 0.009, crit=0.082 ± 0.003 per rating point (14 rating = 1%, 1.146 per %), spell_haste=not significant (0.059 ± 0.061)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 38.0 healing_power points (5.78 DPS) | yes | Filigreed Pristine Circlet (253975, -2.73 DPS, sim-verified) [crafted]; Pristine Circlet (253949, -3.04 DPS) [crafted]; Enchanter's Cowl (4322, -3.89 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 5.9 healing_power points (0.90 DPS) | yes | Crystal Starfire Medallion (5003, -0.11 DPS, sim-verified) [world_drop]; Pendant of Myzrael (4614, -0.14 DPS) [dungeon]; Glowing Green Talisman (5002, -0.14 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.8 healing_power points (2.25 DPS) | yes | Death Speaker Mantle (6685, -0.20 DPS, sim-verified) [dungeon]; Ghostly Mantle (3324, -0.57 DPS) [quest]; Magician's Mantle (12998, -0.61 DPS) [world_drop] |
| back | Battle Healer's Cloak (19529) | Warsong Outriders [rep] | sim-verified (93.9 DPS) | yes | Glowing Thresher Cape (6901, -0.39 DPS, sim-verified) [dungeon]; Windsong Drape (15468, -1.22 DPS) [quest]; Darkspear Raider's Cloak (272078, -1.33 DPS) [vendor] |
| chest | Pristine Gown (253961) | Tailoring [crafted] | 27.8 healing_power points (4.23 DPS) | yes | Death Speaker Robes (6682, -0.24 DPS, sim-verified) [dungeon]; Beguiler Robes (7728, -0.68 DPS) [dungeon]; Robes of Arugal (6324, -0.73 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 15.9 healing_power points (2.41 DPS) | yes | Mindthrust Bracers (1974, -0.72 DPS, sim-verified) [dungeon]; Nightsky Wristbands (6407, -1.45 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.63 DPS) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 16.9 healing_power points (2.58 DPS) | yes | Gloves of Old (9395, -0.57 DPS) [world_drop]; Pristine Gloves (253913, -0.61 DPS) [crafted]; Gloves of Meditation (4318, -0.93 DPS) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 15.1 healing_power points (2.30 DPS) | yes | Defiler's Cloth Girdle (20164, -0.33 DPS) [rep]; Pristine Sash (253925, -0.41 DPS, sim-verified) [crafted]; Warsong Sash (16975, -0.62 DPS) [quest] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 30.6 healing_power points (4.67 DPS) | yes | Earthen Leggings (253999, -0.99 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.01 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -1.70 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 27.8 healing_power points (4.23 DPS) | yes | Acidic Walkers (9454, -2.19 DPS, sim-verified) [dungeon]; Glinteye Slippers (273024, -2.41 DPS) [dungeon]; Boots of the Enchanter (4325, -2.46 DPS) [crafted] |
| finger1 | Solid Gold Band (279845) | Opportunistic Education [quest] | 15.2 healing_power points (2.31 DPS) | yes | Runemender's Seal (273031, -0.05 DPS) [world]; Ring of Power Regulation (273030, -0.35 DPS) [dungeon]; Advisor's Ring (19521, -0.55 DPS) [rep] |
| finger2 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 14.8 healing_power points (2.26 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; Ring of Power Regulation (273030, -0.30 DPS) [dungeon]; Advisor's Ring (19521, -0.50 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (93.9 DPS) | yes | Odo's Ley Staff (6318, -0.91 DPS) [dungeon]; Living Root (6631, -2.87 DPS) [dungeon]; Staff of the Friar (3415, -3.17 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 18.0 healing_power points (2.74 DPS) | yes | Orb of Mistmantle (13031, -0.95 DPS, sim-verified) [world_drop]; Orb of Mystic Insight (249394, -1.09 DPS) [crafted]; Satyr's Rod (15962, -1.90 DPS) [world_drop] |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 7.1 healing_power points (1.09 DPS) | yes | Wand of Eventide (5214, -0.32 DPS) [world_drop]; Greater Mystic Wand (217287, -0.32 DPS) [crafted]; Thunderwood (13062, -0.46 DPS, sim-verified) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Battle Healer's Cloak; chest: Pristine Gown; wrist: Spidertank Oilrag; hands: Truefaith Gloves; waist: Belt of Arugal; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Solid Gold Band; finger2: Electrocutioner Lagnut; main_hand: Death Speaker Scepter; off_hand: Orb of Souls; ranged: Wand of Mana Concentration

No-known-source sample (15 of 295, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (undead, 005005031330011430-20000000000000000-00000000000000000)

Set DPS (verified): 159.4. Weights run: 15.4s. Verify run: 89.4s. 382 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.188, intellect=1.157 ± 0.028, spirit=1.530 ± 0.028, mp5=3.678 ± 0.029, crit=0.377 ± 0.023 per rating point (14 rating = 1%, 5.279 per %), spell_haste=1.581 ± 0.384

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 55.4 healing_power points (8.57 DPS) | yes | Whitemane's Chapeau (7720, -0.99 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -2.05 DPS) [world_drop]; Earthen Silk Hood (254015, -2.81 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 29.3 healing_power points (4.53 DPS) | yes | Triune Amulet (7722, -1.63 DPS) [dungeon]; Necklace of Calisea (1714, -1.94 DPS, sim-verified) [dungeon]; Amberglow Talisman (10824, -2.16 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 37.0 healing_power points (5.73 DPS) | yes | Earthen Silk Shoulders (254033, -1.32 DPS, sim-verified) [crafted]; Mantle of Doan (7712, -1.68 DPS) [dungeon]; Green Silken Shoulders (7057, -2.52 DPS) [crafted] |
| back | Battle Healer's Cloak (19528) | Warsong Outriders [rep] | sim-verified (159.4 DPS) | yes | Glowing Thresher Cape (6901, -0.99 DPS) [dungeon]; Darkspear Raider's Cloak (272077, -1.05 DPS) [vendor]; Cloak of Blight (6832, -1.34 DPS, sim-verified) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 70.2 healing_power points (10.86 DPS) | yes | Dreamweave Vest (10021, -3.89 DPS, sim-verified) [crafted]; Civinad Robes (9623, -4.00 DPS) [quest]; Doomsayer's Robe (4746, -4.13 DPS) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 20.0 healing_power points (3.10 DPS) | yes | Earthen Silk Cuffs (254019, -0.38 DPS, sim-verified) [crafted]; Mindthrust Bracers (1974, -0.50 DPS) [dungeon]; Radiant Silver Bracers (4545, -1.05 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 33.3 healing_power points (5.16 DPS) | yes | Gilded Handwraps (254021, -0.50 DPS) [crafted]; Earthen Silk Gloves (254017, -0.61 DPS, sim-verified) [crafted]; Stormcloth Gloves (10011, -1.31 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 40.4 healing_power points (6.25 DPS) | yes | Deathmage Sash (10771, -2.57 DPS, sim-verified) [dungeon]; Windchaser Cinch (14435, -2.69 DPS) [world_drop]; Sutarn's Ring (13105, -3.04 DPS) [world_drop] |
| legs | Stoneweaver Leggings (9407) | Uldaman: Ironaya [dungeon] | 47.6 healing_power points (7.36 DPS) | yes | Pristine Leggings (253987, -1.53 DPS) [crafted]; Stormcloth Pants (10010, -1.94 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -2.44 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 34.2 healing_power points (5.29 DPS) | yes | Furen's Boots (13100, -0.25 DPS) [world_drop]; Thoughtcast Boots (10578, -0.44 DPS, sim-verified) [dungeon]; Earthen Silk Slippers (254013, -0.85 DPS) [crafted] |
| finger1 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 22.1 healing_power points (3.41 DPS) | yes | Electrocutioner Lagnut (9447, -0.37 DPS) [dungeon]; Runemender's Seal (273031, -0.37 DPS) [world]; Solid Gold Band (279845, -0.53 DPS) [quest] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 20.0 healing_power points (3.10 DPS) | yes | Runemender's Seal (273031, -0.05 DPS) [world]; Electrocutioner Lagnut (9447, -0.15 DPS, sim-verified) [dungeon]; Solid Gold Band (279845, -0.21 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (159.4 DPS) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (159.4 DPS) | yes | Staff of the Friar (3415, -5.01 DPS) [dungeon]; Odo's Ley Staff (6318, -5.55 DPS) [dungeon]; Staff of Jordan (873, -6.58 DPS, sim-verified) [world_drop] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | 28.1 healing_power points (4.35 DPS) | yes | Orb of Souls (249395, -1.08 DPS, sim-verified) [crafted]; Omega Orb (7749, -1.25 DPS) [quest]; Orb of Mistmantle (13031, -2.01 DPS) [world_drop] |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 12.0 healing_power points (1.85 DPS) | yes | Jaina's Firestarter (13064, -0.07 DPS) [world_drop]; Captain Rackmore's Tiller (16789, -0.22 DPS) [quest]; Goblin Igniter (5253, -0.43 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Battle Healer's Cloak; chest: Stormcloth Vest; hands: Dreamweave Gloves; waist: Gilded Cord; legs: Stoneweaver Leggings; finger1: Darkspear Signet; finger2: Advisor's Ring; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Hand of Righteousness; off_hand: Beacon of Hope

No-known-source sample (15 of 382, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16336 Sergeant Major's Cape

### Band 50 (undead, 005005031330011430-33501000000000000-00000000000000000)

Set DPS (verified): 238.8. Weights run: 13.4s. Verify run: 84.6s. 489 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.112, intellect=0.745 ± 0.022, spirit=1.123 ± 0.018, mp5=3.192 ± 0.022, crit=0.327 ± 0.016 per rating point (14 rating = 1%, 4.581 per %), spell_haste=not significant (0.697 ± 0.263)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Cassandra's Grace (13102) | World drop [world_drop] | 54.1 healing_power points (14.59 DPS) | yes | Blood Guard's Satin Cover (220899, +0.00 DPS) [vendor]; Papal Fez (9431, -2.53 DPS, sim-verified) [dungeon]; Stormcloth Headband (10032, -3.08 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 26.0 healing_power points (7.00 DPS) | yes | Mindburst Medallion (11196, -2.47 DPS, sim-verified) [quest]; Triune Amulet (7722, -2.72 DPS) [dungeon]; Horizon Choker (13085, -2.98 DPS) [world_drop] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | 38.7 healing_power points (10.44 DPS) | yes | Blood Guard's Satin Pads (220901, +0.00 DPS) [vendor]; Inquisitor's Shawl (19507, -2.01 DPS, sim-verified) [dungeon]; Earthen Silk Shoulders (254033, -3.16 DPS) [crafted] |
| back | Battle Healer's Cloak (19527) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -0.80 DPS, sim-verified) [dungeon]; Deep Woodlands Cloak (19121, -2.71 DPS) [quest]; Wingveil Cloak (10802, -3.07 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 67.0 healing_power points (18.06 DPS) | yes | Ghostweave Vest (14141, -1.76 DPS, sim-verified) [crafted]; Grimlok's Tribal Vestments (9415, -5.00 DPS) [dungeon]; Stone Guard's Satin Armor (220903, -5.25 DPS) [vendor] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 30.8 healing_power points (8.32 DPS) | yes | Spidertank Oilrag (9448, -3.12 DPS, sim-verified) [dungeon]; Earthen Silk Cuffs (254019, -4.10 DPS) [crafted]; Bloodband Bracers (11469, -4.25 DPS) [quest] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 45.4 healing_power points (12.25 DPS) | yes | Earthenweave Gloves (254075, -0.78 DPS, sim-verified) [crafted]; First Sergeant's Satin Gloves (220898, -3.09 DPS) [vendor]; Greenleaf Handwraps (19116, -4.11 DPS) [quest] |
| waist | Gilded Waistcord (254081) | Tailoring [crafted] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Earthenweave Cord (254077, -0.20 DPS, sim-verified) [crafted]; Gilded Cord (254037, -0.27 DPS) [crafted]; Ghostweave Cord (254073, -2.52 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 50.4 healing_power points (13.60 DPS) | yes | Stone Guard's Satin Leggings (220902, -0.79 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -2.91 DPS) [dungeon]; Stoneweaver Leggings (9407, -2.97 DPS) [dungeon] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Earthenweave Boots (254093, -0.67 DPS, sim-verified) [crafted]; Vinerot Sandals (17748, -2.09 DPS) [dungeon]; Gilded Slippers (254001, -2.16 DPS) [crafted] |
| finger1 | Advisor's Ring (19519) | Warsong Outriders [rep] | 24.8 healing_power points (6.68 DPS) | yes | Eye of Adaegus (5266, -1.84 DPS) [world_drop]; Electrocutioner Lagnut (9447, -2.13 DPS) [dungeon]; Runemender's Seal (273031, -2.13 DPS) [world] |
| finger2 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 22.3 healing_power points (6.02 DPS) | yes | Eye of Adaegus (5266, -1.18 DPS) [world_drop]; Electrocutioner Lagnut (9447, -1.48 DPS) [dungeon]; Runemender's Seal (273031, -1.48 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor]; Evonice's Landin' Pilla (18951, -0.61 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Darkspear Voodoo Seal (272061, -0.46 DPS, sim-verified) [vendor]; Evonice's Landin' Pilla (18951, -1.82 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Resurgence Rod (17743, -6.72 DPS, sim-verified) [quest]; Charstone Dirk (17710, -13.79 DPS) [dungeon]; Dreamstaff (249454, -14.94 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Earthen Rod (9381, -0.10 DPS) [dungeon]; Captain Rackmore's Tiller (16789, -0.35 DPS, sim-verified) [quest]; Lesser Eternal Wand (249232, -0.37 DPS) [crafted] |

**New at 50:** head: Cassandra's Grace; shoulder: Nethergeld Shoulders; back: Battle Healer's Cloak; wrist: Nethergeld Cuffs; hands: Gilded Gloves; waist: Gilded Waistcord; legs: Senior Designer's Pantaloons; feet: Gilded Sandals; finger1: Advisor's Ring; finger2: Darkspear Signet; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 489, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60 (undead, 005005031330011430-33505003030000000-00000000000000000)

Set DPS (verified): 348.8. Weights run: 6.9s. Verify run: 114.7s. 1103 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.014, intellect=0.424 ± 0.011, spirit=0.508 ± 0.009, mp5=1.712 ± 0.013, crit=0.121 ± 0.007 per rating point (14 rating = 1%, 1.692 per %), spell_haste=not significant (0.011 ± 0.104)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mooncloth Circlet (14140) | Tailoring [crafted] | 58.6 healing_power points (18.74 DPS) | yes | Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Champion's Satin Hood (227118, -1.24 DPS) [pvp]; Earthenweave Crown (254137, -3.27 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 23.4 healing_power points (7.48 DPS) | yes | Archlight Talisman (15856, -2.36 DPS, sim-verified) [quest]; Triune Amulet (7722, -3.78 DPS) [dungeon]; Frostwolf Advisor's Pendant (19096, -3.93 DPS) [rep] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (348.8 DPS) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp]; Mooncloth Shoulders (14139, -18.23 DPS, sim-verified) [crafted] |
| back | Battle Healer's Cloak (19526) | Warsong Outriders [rep] | sim-verified (348.8 DPS) | yes | Spritecaster Cape (11623, -3.35 DPS) [dungeon]; Deep Woodlands Cloak (19121, -4.56 DPS) [quest]; Frostwolf Advisor's Cloak (19085, -5.14 DPS) [rep] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 81.1 healing_power points (25.93 DPS) | yes | Warlord's Satin Tunic (231632, -0.86 DPS) [vendor]; Stormcloth Vest (10020, -7.40 DPS, sim-verified) [crafted]; Legionnaire's Satin Tunic (227124, -8.84 DPS) [pvp] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (348.8 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [rep]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (348.8 DPS) | yes | Blood Guard's Satin Handwraps (227126, +0.00 DPS) [pvp]; General's Satin Handwraps (231633, +0.00 DPS) [vendor]; Gilded Gloves (254095, -20.22 DPS, sim-verified) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 41.5 healing_power points (13.27 DPS) | yes | Earthenweave Cord (254077, -3.17 DPS) [crafted]; Gilded Waistcord (254081, -4.16 DPS, sim-verified) [crafted]; Gilded Cord (254037, -4.18 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-verified (348.8 DPS) | yes | Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Earthenweave Leggings (254125, -23.38 DPS, sim-verified) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (348.8 DPS) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; General's Satin Walkers (231630, +0.00 DPS) [vendor]; Swarmtender's Footpads (275609, -18.32 DPS, sim-verified) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 24.0 healing_power points (7.68 DPS) | yes | Deep Rooted Ring (19109, -2.49 DPS) [world]; Darkspear Signet (272068, -2.75 DPS) [vendor]; Solid Gold Band (279845, -3.35 DPS) [quest] |
| finger2 | Advisor's Ring (19518) | Warsong Outriders [rep] | 20.8 healing_power points (6.66 DPS) | yes | Deep Rooted Ring (19109, -1.47 DPS) [world]; Darkspear Signet (272068, -1.74 DPS) [vendor]; Solid Gold Band (279845, -2.34 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (348.8 DPS) | yes | Ankh of Life (1713, -1.78 DPS, sim-verified) [world_drop]; Darkspear Voodoo Seal (272061, -3.02 DPS) [vendor] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (348.8 DPS) | yes | Darkspear Voodoo Seal (272061, -0.55 DPS) [vendor]; Ankh of Life (1713, -1.33 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (348.8 DPS) | yes | High Warlord's Battle Mace (23464, +0.00 DPS) [vendor]; Hand of Edward the Odd (2243, -8.01 DPS, sim-verified) [world_drop]; Resurgence Rod (17743, -11.51 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 22.6 healing_power points (7.22 DPS) | yes | Earthen Rod (9381, -4.34 DPS) [dungeon]; Greater Eternal Wand (249237, -4.34 DPS) [crafted]; Brilliant Wand (249385, -4.41 DPS, sim-verified) [crafted] |

**New at 60:** head: Mooncloth Circlet; shoulder: Devout Mantle; back: Battle Healer's Cloak; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Devout Sandals; finger1: Maiden's Circle; finger2: Advisor's Ring; trinket1: Uther's Strength; trinket2: Shard of the Splithooves; ranged: Torch of Light

No-known-source sample (15 of 1103, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60, raid preset (undead, 005005031330011430-33505003030000000-00000000000000000)

Set DPS (verified): 598.6. Weights run: 5.2s. Verify run: 81.9s. 1103 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.110, intellect=0.622 ± 0.035, spirit=0.439 ± 0.032, mp5=1.173 ± 0.037, crit=0.606 ± 0.027 per rating point (14 rating = 1%, 8.489 per %), spell_haste=3.121 ± 0.665

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mooncloth Circlet (14140) | Tailoring [crafted] | 57.6 healing_power points (26.71 DPS) | yes | Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Champion's Satin Hood (227118, -0.11 DPS) [pvp]; Cassandra's Grace (13102, -4.10 DPS, sim-verified) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 25.0 healing_power points (11.59 DPS) | yes | Archlight Talisman (15856, -2.20 DPS, sim-verified) [quest]; Mindburst Medallion (11196, -6.05 DPS) [quest]; Triune Amulet (7722, -6.45 DPS) [dungeon] |
| shoulder | Mooncloth Shoulders (14139) | Tailoring [crafted] | 41.8 healing_power points (19.40 DPS) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp]; Nethergeld Shoulders (254049, -2.70 DPS, sim-verified) [crafted] |
| back | Battle Healer's Cloak (19526) | Warsong Outriders [rep] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Spritecaster Cape (11623, -4.24 DPS) [dungeon]; Deep Woodlands Cloak (19121, -5.52 DPS) [quest]; Frostwolf Advisor's Cloak (19085, -7.19 DPS) [rep] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 80.0 healing_power points (37.12 DPS) | yes | Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Stormcloth Vest (10020, -7.92 DPS, sim-verified) [crafted]; Legionnaire's Satin Tunic (227124, -11.39 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 30.0 healing_power points (13.94 DPS) | yes | Nethergeld Cuffs (254061, -1.45 DPS, sim-verified) [crafted]; Earthenweave Cuffs (254131, -1.70 DPS) [crafted]; Devout Bracers (16697, -4.49 DPS) [dungeon] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 40.8 healing_power points (18.92 DPS) | yes | Blood Guard's Satin Handwraps (227126, +0.00 DPS) [pvp]; General's Satin Handwraps (231633, +0.00 DPS) [vendor]; Gilded Gloves (254095, -0.93 DPS, sim-verified) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 43.7 healing_power points (20.25 DPS) | yes | Gilded Waistcord (254081, -3.49 DPS, sim-verified) [crafted]; Earthenweave Cord (254077, -5.95 DPS) [crafted]; Gilded Cord (254037, -6.52 DPS) [crafted] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 44.2 healing_power points (20.51 DPS) | yes | Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Earthenweave Leggings (254125, +0.00 DPS) [crafted] |
| feet | Swarmtender's Footpads (275609) | Tailoring [crafted] | 43.0 healing_power points (19.95 DPS) | yes | General's Satin Walkers (231630, +0.00 DPS) [vendor]; Mooncloth Boots (15802, -2.13 DPS) [crafted]; Gilded Sandals (254107, -2.97 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 25.0 healing_power points (11.59 DPS) | yes | Deep Rooted Ring (19109, -3.14 DPS) [world]; Solid Gold Band (279845, -5.47 DPS) [quest]; Philanthropist's Ring (281635, -5.51 DPS) [quest] |
| finger2 | Advisor's Ring (19518) | Warsong Outriders [rep] | 18.7 healing_power points (8.67 DPS) | yes | Deep Rooted Ring (19109, -0.22 DPS) [world]; Solid Gold Band (279845, -2.55 DPS) [quest]; Philanthropist's Ring (281635, -2.59 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (598.7 DPS) | yes | Ankh of Life (1713, -2.81 DPS, sim-verified) [world_drop]; Shard of the Splithooves (10659, -5.08 DPS) [quest]; Darkspear Voodoo Seal (272061, -5.63 DPS) [vendor] |
| trinket2 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (-3.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Shard of the Splithooves (10659, +0.00 DPS) [quest]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Mace (23464, +0.00 DPS) [vendor]; Charstone Dirk (17710, -1.62 DPS, sim-verified) [dungeon]; Resurgence Rod (17743, -21.84 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 21.8 healing_power points (10.12 DPS) | yes | Brilliant Wand (249385, -3.31 DPS, sim-verified) [crafted]; Earthen Rod (9381, -5.95 DPS) [dungeon]; Greater Eternal Wand (249237, -5.95 DPS) [crafted] |

**New at 60:** head: Mooncloth Circlet; shoulder: Mooncloth Shoulders; back: Battle Healer's Cloak; chest: Truefaith Vestments; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Wisdom of the Timbermaw; legs: Outrider's Silk Leggings; feet: Swarmtender's Footpads; finger1: Maiden's Circle; finger2: Advisor's Ring; trinket1: Uther's Strength; trinket2: Dormant Heart of the Mountain; ranged: Torch of Light

No-known-source sample (15 of 1103, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

