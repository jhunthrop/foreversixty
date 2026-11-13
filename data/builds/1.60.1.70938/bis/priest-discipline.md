# Leveling BiS: Discipline

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 005005000010000000-00000000000000000-00000000000000000)

Set DPS (verified): 46.4. Weights run: 23.6s. Verify run: 104.8s. 203 eligible items had no known source.

Race sweep: human 46.4 ± 0.0, gnome 45.9 ± 0.0, night-elf 44.8 ± 0.0, dwarf 44.5 ± 0.0 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.128, intellect=0.879 ± 0.012, spirit=0.951 ± 0.012, mp5=2.615 ± 0.012, crit=0.091 ± 0.008 per rating point (14 rating = 1%, 1.273 per %), spell_haste=not significant (0.171 ± 0.150)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (2.32 DPS) | yes | Shadow Goggles (4373, -0.90 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.83 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 1.9 healing_power points (0.25 DPS) | yes | Tarnished Locket (279870, -0.30 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.9 healing_power points (1.67 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.56 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.65 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -1.15 DPS) [crafted] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.43 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -0.81 DPS) [crafted]; Spirit Cloak (4792, -0.92 DPS) [vendor] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 20.3 healing_power points (2.62 DPS) | yes | Robe of the Moccasin (6465, -0.22 DPS) [dungeon]; Corsair's Overshirt (5202, -0.24 DPS, sim-verified) [dungeon]; Scholarly Robes (2034, -0.75 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 12.2 healing_power points (1.58 DPS) | yes | Bright Bracers (3647, -1.13 DPS) [world_drop]; Seer's Cuffs (3645, -1.22 DPS) [dungeon]; Repurposed Hair Band (281256, -1.34 DPS, sim-verified) [quest] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 13.6 healing_power points (1.76 DPS) | yes | Magefist Gloves (12977, -0.61 DPS, sim-verified) [world_drop]; Bright Gloves (3066, -0.82 DPS) [world_drop]; Serpent Gloves (5970, -0.86 DPS) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 14.5 healing_power points (1.87 DPS) | yes | Novice Ardent's Sash (253887, -0.41 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.97 DPS) [world_drop]; Novice Arcanist's Sash (253885, -1.02 DPS) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 27.1 healing_power points (3.50 DPS) | yes | Darkweave Breeches (12987, -1.79 DPS, sim-verified) [world_drop]; Colorful Kilt (10048, -2.24 DPS) [crafted]; Filigreed Silky Leggings (253939, -2.32 DPS) [crafted] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 11.6 healing_power points (1.50 DPS) | yes | Spidersilk Boots (4320, -0.15 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.37 DPS) [crafted]; Kimbra Boots (6191, -0.68 DPS) [quest] |
| finger1 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 10.2 healing_power points (1.32 DPS) | yes | Band of Purification (12996, -0.58 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.64 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.98 DPS) [world_drop] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.8 healing_power points (0.87 DPS) | yes | Band of Purification (12996, -0.14 DPS, sim-verified) [world_drop]; Lavishly Jeweled Ring (1156, -0.19 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.53 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Westfall (2042) | The Defias Brotherhood [quest] | 58.1 healing_power points (7.50 DPS) | yes | Living Root (6631, -1.08 DPS, sim-verified) [dungeon]; Staff of the Blessed Seer (2271, -2.66 DPS) [dungeon]; Evocator's Blade (2567, -3.05 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Sizzle Stick (8071) (or Greater Magic Wand (11288), Torchlight Wand (5240)) | Deviate Eradication [quest] | 2.0 healing_power points (0.26 DPS) | yes | Torchlight Wand (5240, +0.00 DPS) [quest]; Greater Magic Wand (11288, +0.00 DPS) [crafted]; Sable Wand (7607, -0.01 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Staff of Westfall; ranged: Sizzle Stick

No-known-source sample (15 of 203, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6360 Steelscale Crushfish; 6478 Rat Stompers; 8350 The 1 Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (human, 005005031330001000-00000000000000000-00000000000000000)

Set DPS (verified): 93.7. Weights run: 24.2s. Verify run: 136.0s. 325 eligible items had no known source.

Race sweep: human 93.7 ± 0.0, gnome 90.2 ± 0.0, night-elf 88.2 ± 0.0, dwarf 87.7 ± 0.0 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.040, intellect=0.708 ± 0.008, spirit=1.158 ± 0.008, mp5=2.476 ± 0.009, crit=0.086 ± 0.004 per rating point (14 rating = 1%, 1.202 per %), spell_haste=not significant (0.242 ± 0.087)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 40.0 healing_power points (5.60 DPS) | yes | Pristine Circlet (253949, -3.07 DPS) [crafted]; Filigreed Pristine Circlet (253975, -3.18 DPS, sim-verified) [crafted]; Nightsky Cowl (4039, -3.59 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 7.5 healing_power points (1.05 DPS) | yes | Glowing Green Talisman (5002, -0.07 DPS) [world_drop]; Pendant of Myzrael (4614, -0.15 DPS, sim-verified) [dungeon]; Crystal Starfire Medallion (5003, -0.16 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.4 healing_power points (2.15 DPS) | yes | Death Speaker Mantle (6685, -0.11 DPS, sim-verified) [dungeon]; Mantle of Honor (3560, -0.32 DPS) [quest]; Nightsky Mantle (4718, -0.55 DPS) [world_drop] |
| back | Caretaker's Cape (19533) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Thresher Cape (6901, -0.06 DPS, sim-verified) [dungeon]; Repairman's Cape (9605, -0.52 DPS) [quest]; Prelacy Cape (7004, -1.00 DPS) [quest] |
| chest | Pristine Gown (253961) | Tailoring [crafted] | 29.6 healing_power points (4.14 DPS) | yes | Death Speaker Robes (6682, -0.30 DPS, sim-verified) [dungeon]; Beguiler Robes (7728, -0.40 DPS) [dungeon]; Robes of Arugal (6324, -0.65 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 16.4 healing_power points (2.30 DPS) | yes | Mindthrust Bracers (1974, -0.67 DPS, sim-verified) [dungeon]; Nightsky Wristbands (6407, -1.22 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.51 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 17.4 healing_power points (2.44 DPS) | yes | Pristine Gloves (253913, -0.60 DPS) [crafted]; Gloves of Meditation (4318, -0.61 DPS) [crafted]; Truefaith Gloves (7049, -0.76 DPS, sim-verified) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 16.9 healing_power points (2.37 DPS) | yes | Highlander's Cloth Girdle (20099, -0.53 DPS) [rep]; Novice Ardent's Sash (253887, -0.74 DPS) [crafted]; Pristine Sash (253925, -0.75 DPS, sim-verified) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 32.7 healing_power points (4.59 DPS) | yes | Filigreed Pristine Leggings (253937, -0.82 DPS) [crafted]; Earthen Leggings (253999, -0.87 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -1.45 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 29.6 healing_power points (4.14 DPS) | yes | Boots of the Enchanter (4325, -2.15 DPS) [crafted]; Glinteye Slippers (273024, -2.19 DPS) [dungeon]; Acidic Walkers (9454, -2.30 DPS, sim-verified) [dungeon] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 20.2 healing_power points (2.83 DPS) | yes | Ring of Power Regulation (273030, -0.95 DPS) [dungeon]; Lorekeeper's Ring (19525, -1.16 DPS) [rep]; The Queen's Jewel (13094, -1.33 DPS) [world_drop] |
| finger2 | Electrocutioner Lagnut (9447) | Gnomeregan: Electrocutioner 6000 [dungeon] | 17.1 healing_power points (2.40 DPS) | yes | Ring of Power Regulation (273030, -0.51 DPS, sim-verified) [dungeon]; Lorekeeper's Ring (19525, -0.72 DPS) [rep]; The Queen's Jewel (13094, -0.90 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Odo's Ley Staff (6318, -0.29 DPS) [dungeon]; Staff of Westfall (2042, -1.33 DPS) [quest]; Staff of the Friar (3415, -2.94 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 20.0 healing_power points (2.79 DPS) | yes | Eye of Paleth (2943, -0.97 DPS) [quest]; Orb of Mistmantle (13031, -1.09 DPS, sim-verified) [world_drop]; Orb of Mystic Insight (249394, -1.22 DPS) [crafted] |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 7.8 healing_power points (1.09 DPS) | yes | Wand of Eventide (5214, -0.39 DPS) [world_drop]; Greater Mystic Wand (217287, -0.39 DPS) [crafted]; Thunderwood (13062, -0.87 DPS, sim-verified) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Caretaker's Cape; chest: Pristine Gown; wrist: Spidertank Oilrag; hands: Gloves of Old; waist: Belt of Arugal; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Runemender's Seal; finger2: Electrocutioner Lagnut; main_hand: Death Speaker Scepter; off_hand: Orb of Souls; ranged: Wand of Mana Concentration

No-known-source sample (15 of 325, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6360 Steelscale Crushfish

### Band 40 (human, 005005031330011000-33300000000000000-00000000000000000)

Set DPS (verified): 143.2. Weights run: 38.6s. Verify run: 205.7s. 461 eligible items had no known source.

Race sweep: human 143.2 ± 0.1, gnome 136.6 ± 0.1, dwarf 134.8 ± 0.1, night-elf 132.4 ± 0.1 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.114, intellect=1.181 ± 0.021, spirit=1.706 ± 0.029, mp5=3.034 ± 0.022, crit=0.276 ± 0.015 per rating point (14 rating = 1%, 3.868 per %), spell_haste=not significant (0.097 ± 0.386)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 57.4 healing_power points (9.46 DPS) | yes | Whitemane's Chapeau (7720, -0.85 DPS) [dungeon]; Holy Shroud (2721, -2.34 DPS) [world_drop]; Earthen Silk Hood (254015, -3.03 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 29.4 healing_power points (4.85 DPS) | yes | Necklace of Calisea (1714, -1.52 DPS) [dungeon]; Triune Amulet (7722, -1.57 DPS) [dungeon]; Amberglow Talisman (10824, -2.04 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 37.4 healing_power points (6.16 DPS) | yes | Earthen Silk Shoulders (254033, -0.94 DPS) [crafted]; Mantle of Doan (7712, -1.59 DPS) [dungeon]; Mistscape Mantle (4734, -2.61 DPS) [dungeon] |
| back | Caretaker's Cape (19532) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Thresher Cape (6901, -0.97 DPS) [dungeon]; Expeditionary's Cape (271922, -1.11 DPS) [rep]; Repairman's Cape (9605, -1.39 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 71.6 healing_power points (11.81 DPS) | yes | Dreamweave Vest (10021, -3.15 DPS) [crafted]; Death's Head Vestment (10581, -3.63 DPS) [dungeon]; Doomsayer's Robe (4746, -4.11 DPS) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Earthen Silk Cuffs (254019, +0.00 DPS) [crafted]; Mindthrust Bracers (1974, -0.51 DPS) [dungeon]; Silkstream Cuffs (16791, -0.89 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 34.7 healing_power points (5.71 DPS) | yes | Earthen Silk Gloves (254017, -0.50 DPS) [crafted]; Gilded Handwraps (254021, -0.72 DPS) [crafted]; Gloves of Old (9395, -1.37 DPS) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 41.7 healing_power points (6.87 DPS) | yes | Windchaser Cinch (14435, -2.78 DPS) [world_drop]; Deathmage Sash (10771, -2.80 DPS) [dungeon]; Sutarn's Ring (13105, -3.24 DPS) [world_drop] |
| legs | Stoneweaver Leggings (9407) | Uldaman: Ironaya [dungeon] | 49.9 healing_power points (8.23 DPS) | yes | Stormcloth Pants (10010, -1.54 DPS) [crafted]; Pristine Leggings (253987, -1.83 DPS) [crafted]; Earthen Leggings (253999, -2.66 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thoughtcast Boots (10578, +0.00 DPS) [dungeon]; Furen's Boots (13100, +0.00 DPS) [world_drop]; Earthen Silk Slippers (254013, -1.01 DPS) [crafted] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 26.3 healing_power points (4.34 DPS) | yes | Electrocutioner Lagnut (9447, -0.89 DPS) [dungeon]; Lorekeeper's Ring (19524, -1.35 DPS) [rep]; Welken Ring (5011, -1.40 DPS) [world_drop] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Welken Ring (5011, +0.00 DPS) [world_drop]; Electrocutioner Lagnut (9447, +0.00 DPS) [dungeon]; Lorekeeper's Ring (19524, +0.00 DPS) [rep] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Jordan (873, -1.15 DPS) [world_drop]; Staff of the Friar (3415, -5.21 DPS) [dungeon]; Odo's Ley Staff (6318, -5.80 DPS) [dungeon] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 29.8 healing_power points (4.91 DPS) | yes | Beacon of Hope (9393, -0.16 DPS) [dungeon]; Orb of Souls (249395, -1.08 DPS) [crafted]; Orb of Mistmantle (13031, -2.31 DPS) [world_drop] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 12.2 healing_power points (2.01 DPS) | yes | Captain Rackmore's Tiller (16789, -0.18 DPS) [quest]; Wand of Mana Concentration (273043, -0.23 DPS) [world]; Goblin Igniter (5253, -0.39 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Caretaker's Cape; chest: Stormcloth Vest; hands: Dreamweave Gloves; waist: Gilded Cord; legs: Stoneweaver Leggings; finger2: Snake Hoop; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Hand of Righteousness; off_hand: Orb of Lorica; ranged: Jaina's Firestarter

No-known-source sample (15 of 461, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (human, 005005031330011430-33501000000000000-00000000000000000)

Set DPS (verified): 209.3. Weights run: 40.3s. Verify run: 349.4s. 608 eligible items had no known source.

Race sweep: human 209.3 ± 0.1, gnome 201.5 ± 0.1, night-elf 197.4 ± 0.1, dwarf 195.9 ± 0.1 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.046, intellect=1.334 ± 0.030, spirit=2.441 ± 0.027, mp5=5.744 ± 0.031, crit=0.361 ± 0.014 per rating point (14 rating = 1%, 5.047 per %), spell_haste=not significant (-1.420 ± 0.496)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 79.6 healing_power points (7.41 DPS) | yes | Knight-Lieutenant's Satin Cover (220896, -0.42 DPS) [vendor]; Papal Fez (9431, -1.21 DPS) [dungeon]; Cassandra's Grace (13102, -1.27 DPS) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Horizon Choker (13085, -0.21 DPS) [world_drop]; Gemshard Heart (17707, -0.34 DPS) [dungeon]; Lei of Lilies (1315, -2.66 DPS, sim-verified) [world_drop] |
| shoulder | Mantle of Lost Hope (22234) | Blackrock Depths: Lord Roccor [dungeon] | 63.6 healing_power points (5.92 DPS) | yes | Knight-Lieutenant's Satin Pads (220894, -0.50 DPS) [vendor]; Kentic Amice (11624, -1.64 DPS) [dungeon]; Nethergeld Shoulders (254049, -1.80 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Caretaker's Cape (19531, -0.00 DPS) [rep]; Wingveil Cloak (10802, -0.26 DPS) [dungeon]; Featherskin Cape (10843, -0.54 DPS, sim-verified) [world] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 119.6 healing_power points (11.13 DPS) | yes | Stormcloth Vest (10020, -3.91 DPS) [crafted]; Grimlok's Tribal Vestments (9415, -4.63 DPS) [dungeon]; Ghostweave Vest (14141, -5.53 DPS, sim-verified) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 41.5 healing_power points (3.87 DPS) | yes | Spidertank Oilrag (9448, -1.42 DPS) [dungeon]; Venomshroud Armguards (14439, -1.55 DPS) [world_drop]; Aristocratic Cuffs (12546, -1.57 DPS, sim-verified) [dungeon] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 61.9 healing_power points (5.76 DPS) | yes | Earthenweave Gloves (254075, -0.84 DPS) [crafted]; Jumanza Grips (18083, -1.08 DPS) [dungeon]; Gloves of the Atal'ai Prophet (10808, -3.87 DPS, sim-verified) [dungeon] |
| waist | Earthenweave Cord (254077) | Tailoring [crafted] | 52.8 healing_power points (4.92 DPS) | yes | Gilded Cord (254037, -0.63 DPS, sim-verified) [crafted]; Dawnspire Cord (12466, -0.88 DPS) [dungeon]; Gilded Waistcord (254081, -1.10 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 90.3 healing_power points (8.40 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -1.24 DPS) [dungeon]; Dalewind Trousers (13008, -1.61 DPS) [world_drop]; Rainstrider Leggings (11123, -4.36 DPS, sim-verified) [quest] |
| feet | Vinerot Sandals (17748) | Maraudon: Razorlash [dungeon] | 54.4 healing_power points (5.06 DPS) | yes | Earthenweave Boots (254093, -0.15 DPS) [crafted]; Furen's Boots (13100, -0.56 DPS) [world_drop]; Mistwalker Boots (10629, -3.13 DPS, sim-verified) [dungeon] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 37.3 healing_power points (3.47 DPS) | yes | Lorekeeper's Ring (19523, -0.22 DPS) [rep]; Choking Band (11868, -0.52 DPS) [quest]; Runemender's Seal (273031, -0.53 DPS) [world] |
| finger2 | Brainlash (6440) | Zul'Farrak [dungeon] | 36.5 healing_power points (3.40 DPS) | yes | Choking Band (11868, -0.45 DPS) [quest]; Runemender's Seal (273031, -0.46 DPS) [world]; Lorekeeper's Ring (19523, -1.38 DPS, sim-verified) [rep] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Enriched Seal (271903, -0.05 DPS) [rep]; Evonice's Landin' Pilla (18951, -0.45 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.91 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Enriched Seal (271903, +0.00 DPS) [rep]; Dormant Heart of the Mountain (249473, -0.97 DPS, sim-verified) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamstaff (249454, -5.52 DPS) [crafted]; Charstone Dirk (17710, -6.37 DPS) [dungeon]; Resurgence Rod (17743, -9.49 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cairnstone Sliver (9654) | The Morrow Stone [quest] | 18.9 healing_power points (1.76 DPS) | yes | Jaina's Firestarter (13064, -0.33 DPS) [world_drop]; Captain Rackmore's Tiller (16789, -0.52 DPS) [quest]; Wand of Mana Concentration (273043, -1.29 DPS, sim-verified) [world] |

**New at 50:** head: Gemburst Circlet; shoulder: Mantle of Lost Hope; back: Spritecaster Cape; chest: Embrace of the Wind Serpent; wrist: Nethergeld Cuffs; hands: Gilded Gloves; waist: Earthenweave Cord; legs: Senior Designer's Pantaloons; feet: Vinerot Sandals; finger1: Eye of Adaegus; finger2: Brainlash; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Cairnstone Sliver

No-known-source sample (15 of 608, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (human, 005005031330011430-33505003030000000-00000000000000000)

Set DPS (verified): 293.6. Weights run: 41.0s. Verify run: 1118.6s. 1525 eligible items had no known source.

Race sweep: human 293.7 ± 0.2, gnome 292.3 ± 0.2, night-elf 281.3 ± 0.2, dwarf 278.8 ± 0.2 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.111, intellect=2.138 ± 0.069, spirit=2.706 ± 0.050, mp5=5.192 ± 0.079, crit=0.506 ± 0.025 per rating point (14 rating = 1%, 7.090 per %), spell_haste=not significant (-2.397 ± 1.017)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | 111.9 healing_power points (10.90 DPS) | yes | Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Crown of Caer Darrow (13986, -8.47 DPS, sim-verified) [quest] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 83.5 healing_power points (8.14 DPS) | yes | Dragon's Eye (10829, -2.92 DPS) [dungeon]; Animated Chain Necklace (18723, -3.08 DPS) [dungeon]; Tooth of Gnarr (13141, -3.56 DPS, sim-verified) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | 80.2 healing_power points (7.82 DPS) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp]; Mantle of the Scarlet Crusade (22405, -6.53 DPS, sim-verified) [dungeon] |
| back | Gracious Cape (18743) | Stratholme: Ghostly Citizen [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Drape of Recovery (272413, -0.13 DPS) [vendor]; Expeditionary's Cape (271907, -0.26 DPS) [rep]; Shroud of the Exile (15421, -1.30 DPS, sim-verified) [quest] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 135.5 healing_power points (13.21 DPS) | yes | Field Marshal's Satin Tunic (231624, +0.00 DPS) [vendor]; Mooncloth Vest (14138, -1.50 DPS) [crafted]; Robes of the Exalted (13346, -1.77 DPS, sim-verified) [dungeon] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Satin Bracers (17606, +0.00 DPS) [pvp]; Dryad's Wrist Bindings (19595, +0.00 DPS) [rep]; Wyrmthalak's Shackles (13958, -6.04 DPS, sim-verified) [quest] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Theramore Handwraps (271945, -0.06 DPS) [rep]; Magister's Gloves (16684, -0.74 DPS) [dungeon]; Hands of the Exalted Herald (12554, -3.71 DPS, sim-verified) [dungeon] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Satin Sash (17609, +0.00 DPS) [pvp]; Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Penitent's Cinch (272394, -3.01 DPS, sim-verified) [vendor] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 118.3 healing_power points (11.53 DPS) | yes | Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Knight-Captain's Satin Legguards (227125, -0.81 DPS) [pvp]; Senior Designer's Pantaloons (11841, -7.22 DPS, sim-verified) [dungeon] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 89.1 healing_power points (8.68 DPS) | yes | Marshal's Satin Walkers (231627, -0.65 DPS) [vendor]; Faith Healer's Boots (22247, -0.68 DPS) [dungeon]; Devout Sandals (16691, -1.05 DPS) [dungeon] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -0.93 DPS) [vendor]; The Postmaster's Seal (13392, -1.01 DPS) [dungeon]; Naglering (11669, -9.62 DPS, sim-verified) [dungeon] |
| finger2 | Rosewine Circle (13178) | Blackrock Spire: Urok Doomhowl [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Postmaster's Seal (13392, -0.32 DPS) [dungeon]; Channeler's Ring (272406, -0.35 DPS) [vendor]; Blessed Band of Light (272407, -1.16 DPS, sim-verified) [vendor] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -0.13 DPS) [quest]; Evonice's Landin' Pilla (18951, -0.53 DPS) [quest]; Serenity Field (272439, -4.67 DPS, sim-verified) [vendor] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -1.01 DPS) [quest]; Evonice's Landin' Pilla (18951, -1.41 DPS) [quest]; Serenity Field (272439, -2.75 DPS, sim-verified) [vendor] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Warhammer (23454, +0.00 DPS) [vendor]; Argent Crusader (13249, -0.99 DPS) [quest]; Staff of Metanoia (22394, -3.07 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Eternal Light (22254) | Blackrock Depths: Watchman Doomgrip [dungeon] | 32.2 healing_power points (3.14 DPS) | yes | Oblivion's Touch (18761, -0.80 DPS) [dungeon]; Brilliant Wand (249385, -0.80 DPS) [crafted]; Cairnstone Sliver (9654, -1.18 DPS, sim-verified) [quest] |

**New at 60:** head: Devout Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; back: Gracious Cape; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of Mending; finger2: Rosewine Circle; trinket2: Burst of Knowledge; ranged: Wand of Eternal Light

No-known-source sample (15 of 1525, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (human, 005005031330011430-33505003030000000-00000000000000000)

Set DPS (verified): 624.5. Weights run: 27.4s. Verify run: 716.8s. 1525 eligible items had no known source.

Race sweep: human 624.0 ± 0.3, gnome 621.8 ± 0.4, night-elf 607.2 ± 0.3, dwarf 604.6 ± 0.3 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.233, intellect=1.323 ± 0.061, spirit=0.985 ± 0.063, mp5=2.811 ± 0.069, crit=0.867 ± 0.049 per rating point (14 rating = 1%, 12.136 per %), spell_haste=not significant (0.840 ± 1.352)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | 66.5 healing_power points (18.37 DPS) | yes | Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -14.94 DPS, sim-verified) [crafted] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 51.7 healing_power points (14.27 DPS) | yes | Tooth of Gnarr (13141, -3.78 DPS) [dungeon]; Lady Maye's Pendant (14558, -4.29 DPS) [world_drop]; Animated Chain Necklace (18723, -4.52 DPS, sim-verified) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp]; Mantle of Lost Hope (22234, -13.99 DPS, sim-verified) [dungeon] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Caretaker's Cape (19530, -2.41 DPS, sim-verified) [rep]; Gracious Cape (18743, -2.52 DPS) [dungeon]; Shroud of the Exile (15421, -2.79 DPS) [quest] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 88.8 healing_power points (24.51 DPS) | yes | Field Marshal's Satin Tunic (231624, +0.00 DPS) [vendor]; Field Marshal's Satin Vestments (17605, -2.06 DPS) [vendor]; Embrace of the Wind Serpent (12462, -4.57 DPS, sim-verified) [world] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Funeral Cuffs (12626, +0.00 DPS) [dungeon]; Marshal's Satin Bracers (17606, +0.00 DPS) [pvp]; Dryad's Wrist Bindings (19595, -12.66 DPS, sim-verified) [rep] |
| hands | Hands of the Exalted Herald (12554) | Blackrock Depths: Princess Moira Bronzebeard  [dungeon] | 60.0 healing_power points (16.58 DPS) | yes | Marshal's Satin Handwraps (231623, +0.00 DPS) [vendor]; Theramore Handwraps (271945, -1.87 DPS, sim-verified) [rep] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 58.4 healing_power points (16.13 DPS) | yes | Devout Belt (16696, -3.33 DPS) [dungeon]; Marshal's Satin Sash (17609, -3.64 DPS) [pvp]; Penitent's Cinch (272394, -4.58 DPS, sim-verified) [vendor] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Satin Legguards (227125, +0.00 DPS) [pvp]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Sentinel's Silk Leggings (237815, -14.92 DPS, sim-verified) [vendor] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Faith Healer's Boots (22247, +0.00 DPS) [dungeon]; Marshal's Satin Walkers (231627, +0.00 DPS) [vendor]; Mooncloth Boots (15802, -11.58 DPS, sim-verified) [crafted] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, -1.46 DPS) [dungeon]; Blessed Band of Light (272407, -2.01 DPS) [vendor]; Naglering (11669, -13.37 DPS, sim-verified) [dungeon] |
| finger2 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -1.67 DPS) [vendor]; Maiden's Circle (13001, -3.25 DPS) [world_drop]; Rosewine Circle (13178, -3.70 DPS, sim-verified) [dungeon] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Uther's Strength (11302, -0.83 DPS) [world_drop]; Weakness Analyzer (272438, -0.83 DPS) [vendor] |
| trinket2 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Uther's Strength (11302, -3.33 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Argent Crusader (13249, -4.40 DPS, sim-verified) [quest]; Redemption (22406, -5.71 DPS) [dungeon]; Staff of Metanoia (22394, -5.72 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Eternal Light (22254) | Blackrock Depths: Watchman Doomgrip [dungeon] | 22.6 healing_power points (6.23 DPS) | yes | Brilliant Wand (249385, -1.19 DPS) [crafted]; Bonecreeper Stylus (13938, -1.37 DPS) [dungeon]; Oblivion's Touch (18761, -1.98 DPS, sim-verified) [dungeon] |

**New at 60:** head: Devout Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; back: Drape of Recovery; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Hands of the Exalted Herald; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of Mending; finger2: Fordring's Seal; trinket1: Serenity Field; trinket2: Dormant Heart of the Mountain; ranged: Wand of Eternal Light

No-known-source sample (15 of 1525, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

## Horde

### Band 20 (undead, 005005000010000000-00000000000000000-00000000000000000)

Set DPS (verified): 45.6. Weights run: 23.6s. Verify run: 93.9s. 196 eligible items had no known source.

Race sweep: undead 45.6 ± 0.0, troll 43.4 ± 0.0 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.128, intellect=0.879 ± 0.012, spirit=0.951 ± 0.012, mp5=2.615 ± 0.012, crit=0.091 ± 0.008 per rating point (14 rating = 1%, 1.273 per %), spell_haste=not significant (0.171 ± 0.150)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (2.32 DPS) | yes | Shadow Goggles (4373, -1.58 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.83 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 2.9 healing_power points (0.37 DPS) | yes | Scholarly Pendant (277203, -0.12 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.9 healing_power points (1.67 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.34 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.65 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -1.15 DPS) [crafted] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | sim-verified (45.6 DPS) | yes | Heavy Woolen Cloak (4311, -0.52 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -0.81 DPS) [crafted]; Spirit Cloak (4792, -0.92 DPS) [vendor] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 20.3 healing_power points (2.62 DPS) | yes | Corsair's Overshirt (5202, -0.21 DPS) [dungeon]; Robe of the Moccasin (6465, -0.22 DPS) [dungeon]; Scholarly Robes (2034, -0.75 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 12.2 healing_power points (1.58 DPS) | yes | Tabitha's Cuffs (251486, -1.19 DPS) [quest]; Crystalline Cuffs (14148, -1.21 DPS) [dungeon]; Bright Bracers (3647, -1.21 DPS, sim-verified) [world_drop] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 13.6 healing_power points (1.76 DPS) | yes | Magefist Gloves (12977, -0.77 DPS, sim-verified) [world_drop]; Bright Gloves (3066, -0.82 DPS) [world_drop]; Serpent Gloves (5970, -0.86 DPS) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 14.5 healing_power points (1.87 DPS) | yes | Novice Ardent's Sash (253887, -0.55 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.97 DPS) [world_drop]; Novice Arcanist's Sash (253885, -1.02 DPS) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 27.1 healing_power points (3.50 DPS) | yes | Darkweave Breeches (12987, -1.95 DPS, sim-verified) [world_drop]; Colorful Kilt (10048, -2.24 DPS) [crafted]; Filigreed Silky Leggings (253939, -2.32 DPS) [crafted] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 11.6 healing_power points (1.50 DPS) | yes | Spidersilk Boots (4320, -0.18 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.37 DPS) [crafted]; Smoldering Boots (3076, -0.77 DPS) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 10.2 healing_power points (1.32 DPS) | yes | Lavishly Jeweled Ring (1156, -0.64 DPS) [dungeon]; Loop of Sacrifice (281673, -0.75 DPS) [quest]; Volcanic Rock Ring (12053, -0.98 DPS) [world_drop] |
| finger2 | Band of Purification (12996) | World drop [world_drop] | 5.7 healing_power points (0.74 DPS) | yes | Loop of Sacrifice (281673, -0.17 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.31 DPS, sim-verified) [dungeon]; Volcanic Rock Ring (12053, -0.40 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Living Root (6631) | Wailing Caverns: Verdan the Everliving [dungeon] | 49.8 healing_power points (6.43 DPS) | yes | Staff of the Blessed Seer (2271, -1.35 DPS, sim-verified) [dungeon]; Gnarled Necromancer's Staff (251534, -1.84 DPS) [quest]; Evocator's Blade (2567, -1.97 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Sizzle Stick (8071) (or Greater Magic Wand (11288), Torchlight Wand (5240)) | Deviate Eradication [quest] | 2.0 healing_power points (0.26 DPS) | yes | Torchlight Wand (5240, +0.00 DPS) [quest]; Greater Magic Wand (11288, +0.00 DPS) [crafted]; Wand of Decay (5252, -0.03 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Advisor's Ring; finger2: Band of Purification; main_hand: Living Root; ranged: Sizzle Stick

No-known-source sample (15 of 196, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6360 Steelscale Crushfish; 6478 Rat Stompers; 8350 The 1 Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 20425 Advisor's Gnarled Staff

### Band 30 (undead, 005005031330001000-00000000000000000-00000000000000000)

Set DPS (verified): 95.0. Weights run: 24.2s. Verify run: 128.4s. 318 eligible items had no known source.

Race sweep: undead 95.0 ± 0.0, troll 87.5 ± 0.0 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.040, intellect=0.708 ± 0.008, spirit=1.158 ± 0.008, mp5=2.476 ± 0.009, crit=0.086 ± 0.004 per rating point (14 rating = 1%, 1.202 per %), spell_haste=not significant (0.242 ± 0.087)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 40.0 healing_power points (5.60 DPS) | yes | Filigreed Pristine Circlet (253975, -3.03 DPS, sim-verified) [crafted]; Pristine Circlet (253949, -3.07 DPS) [crafted]; Nightsky Cowl (4039, -3.59 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 7.5 healing_power points (1.05 DPS) | yes | Glowing Green Talisman (5002, -0.07 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.16 DPS) [world_drop]; Pendant of Myzrael (4614, -0.27 DPS, sim-verified) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.4 healing_power points (2.15 DPS) | yes | Death Speaker Mantle (6685, -0.09 DPS, sim-verified) [dungeon]; Ghostly Mantle (3324, -0.34 DPS) [quest]; Nightsky Mantle (4718, -0.55 DPS) [world_drop] |
| back | Battle Healer's Cloak (19529) | Warsong Outriders [rep] | sim-verified (95.0 DPS) | yes | Glowing Thresher Cape (6901, -0.07 DPS, sim-verified) [dungeon]; Windsong Drape (15468, -1.12 DPS) [quest]; Darkspear Raider's Cloak (272078, -1.19 DPS) [vendor] |
| chest | Pristine Gown (253961) | Tailoring [crafted] | 29.6 healing_power points (4.14 DPS) | yes | Death Speaker Robes (6682, -0.25 DPS) [dungeon]; Beguiler Robes (7728, -0.40 DPS) [dungeon]; Robes of Arugal (6324, -0.65 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 16.4 healing_power points (2.30 DPS) | yes | Mindthrust Bracers (1974, -0.60 DPS, sim-verified) [dungeon]; Nightsky Wristbands (6407, -1.22 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.51 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 17.4 healing_power points (2.44 DPS) | yes | Pristine Gloves (253913, -0.60 DPS) [crafted]; Gloves of Meditation (4318, -0.61 DPS) [crafted]; Truefaith Gloves (7049, -0.63 DPS, sim-verified) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 16.9 healing_power points (2.37 DPS) | yes | Pristine Sash (253925, -0.39 DPS, sim-verified) [crafted]; Defiler's Cloth Girdle (20164, -0.53 DPS) [rep]; Novice Ardent's Sash (253887, -0.74 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 32.7 healing_power points (4.59 DPS) | yes | Filigreed Pristine Leggings (253937, -0.82 DPS) [crafted]; Earthen Leggings (253999, -1.06 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -1.45 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 29.6 healing_power points (4.14 DPS) | yes | Boots of the Enchanter (4325, -2.15 DPS) [crafted]; Acidic Walkers (9454, -2.18 DPS, sim-verified) [dungeon]; Glinteye Slippers (273024, -2.19 DPS) [dungeon] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 20.2 healing_power points (2.83 DPS) | yes | Solid Gold Band (279845, -0.48 DPS) [quest]; Ring of Power Regulation (273030, -0.95 DPS) [dungeon]; Advisor's Ring (19521, -1.16 DPS) [rep] |
| finger2 | Electrocutioner Lagnut (9447) | Gnomeregan: Electrocutioner 6000 [dungeon] | 17.1 healing_power points (2.40 DPS) | yes | Solid Gold Band (279845, -0.04 DPS) [quest]; Ring of Power Regulation (273030, -0.52 DPS) [dungeon]; Advisor's Ring (19521, -0.72 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (95.0 DPS) | yes | Odo's Ley Staff (6318, -0.29 DPS) [dungeon]; Living Root (6631, -2.41 DPS) [dungeon]; Staff of the Friar (3415, -2.91 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 20.0 healing_power points (2.79 DPS) | yes | Orb of Mistmantle (13031, -1.02 DPS, sim-verified) [world_drop]; Orb of Mystic Insight (249394, -1.22 DPS) [crafted]; Strength of Will (4837, -1.82 DPS) [vendor] |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 7.8 healing_power points (1.09 DPS) | yes | Wand of Eventide (5214, -0.39 DPS) [world_drop]; Greater Mystic Wand (217287, -0.39 DPS) [crafted]; Thunderwood (13062, -0.61 DPS, sim-verified) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Battle Healer's Cloak; chest: Pristine Gown; wrist: Spidertank Oilrag; hands: Gloves of Old; waist: Belt of Arugal; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Runemender's Seal; finger2: Electrocutioner Lagnut; main_hand: Death Speaker Scepter; off_hand: Orb of Souls; ranged: Wand of Mana Concentration

No-known-source sample (15 of 318, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6360 Steelscale Crushfish; 6478 Rat Stompers; 8350 The 1 Ring

### Band 40 (undead, 005005031330011000-33300000000000000-00000000000000000)

Set DPS (verified): 147.1. Weights run: 38.6s. Verify run: 195.7s. 454 eligible items had no known source.

Race sweep: undead 147.2 ± 0.1, troll 135.4 ± 0.0 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.114, intellect=1.181 ± 0.021, spirit=1.706 ± 0.029, mp5=3.034 ± 0.022, crit=0.276 ± 0.015 per rating point (14 rating = 1%, 3.868 per %), spell_haste=not significant (0.097 ± 0.386)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 57.4 healing_power points (9.46 DPS) | yes | Whitemane's Chapeau (7720, -0.64 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -2.34 DPS) [world_drop]; Earthen Silk Hood (254015, -3.03 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 29.4 healing_power points (4.85 DPS) | yes | Triune Amulet (7722, -1.57 DPS) [dungeon]; Necklace of Calisea (1714, -1.74 DPS, sim-verified) [dungeon]; Amberglow Talisman (10824, -2.04 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 37.4 healing_power points (6.16 DPS) | yes | Earthen Silk Shoulders (254033, -1.29 DPS, sim-verified) [crafted]; Mantle of Doan (7712, -1.59 DPS) [dungeon]; Mistscape Mantle (4734, -2.61 DPS) [dungeon] |
| back | Battle Healer's Cloak (19528) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Thresher Cape (6901, -0.97 DPS) [dungeon]; Cloak of Blight (6832, -0.97 DPS, sim-verified) [quest]; Darkspear Raider's Cloak (272077, -1.11 DPS) [vendor] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 71.6 healing_power points (11.81 DPS) | yes | Dreamweave Vest (10021, -3.00 DPS, sim-verified) [crafted]; Death's Head Vestment (10581, -3.63 DPS) [dungeon]; Doomsayer's Robe (4746, -4.11 DPS) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Mindthrust Bracers (1974, -0.51 DPS) [dungeon]; Radiant Silver Bracers (4545, -0.77 DPS) [quest]; Earthen Silk Cuffs (254019, -0.93 DPS, sim-verified) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 34.7 healing_power points (5.71 DPS) | yes | Earthen Silk Gloves (254017, -0.35 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -0.72 DPS) [crafted]; Gloves of Old (9395, -1.37 DPS) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 41.7 healing_power points (6.87 DPS) | yes | Windchaser Cinch (14435, -2.47 DPS, sim-verified) [world_drop]; Deathmage Sash (10771, -2.80 DPS) [dungeon]; Sutarn's Ring (13105, -3.24 DPS) [world_drop] |
| legs | Stoneweaver Leggings (9407) | Uldaman: Ironaya [dungeon] | 49.9 healing_power points (8.23 DPS) | yes | Stormcloth Pants (10010, -1.26 DPS, sim-verified) [crafted]; Pristine Leggings (253987, -1.83 DPS) [crafted]; Earthen Leggings (253999, -2.66 DPS) [crafted] |
| feet | Thoughtcast Boots (10578) | Razorfen Downs [dungeon] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Gilded Slippers (254001, -0.01 DPS) [crafted]; Furen's Boots (13100, -0.35 DPS, sim-verified) [world_drop]; Earthen Silk Slippers (254013, -1.03 DPS) [crafted] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 26.3 healing_power points (4.34 DPS) | yes | Darkspear Signet (272070, -1.34 DPS) [vendor]; Advisor's Ring (19520, -1.35 DPS) [rep]; Electrocutioner Lagnut (9447, -2.28 DPS, sim-verified) [dungeon] |
| finger2 | Solid Gold Band (279845) | Opportunistic Education [quest] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Darkspear Signet (272070, -0.22 DPS) [vendor]; Advisor's Ring (19520, -0.23 DPS) [rep]; Electrocutioner Lagnut (9447, -0.27 DPS, sim-verified) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+3.9 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Jordan (873, -4.70 DPS, sim-verified) [world_drop]; Staff of the Friar (3415, -5.21 DPS) [dungeon]; Odo's Ley Staff (6318, -5.80 DPS) [dungeon] |
| off_hand | Prophetic Cane (6803) | Into The Scarlet Monastery [quest] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Orb of Souls (249395, -0.69 DPS) [crafted]; Beacon of Hope (9393, -0.86 DPS, sim-verified) [dungeon]; Omega Orb (7749, -1.54 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 12.2 healing_power points (2.01 DPS) | yes | Wand of Mana Concentration (273043, -0.23 DPS) [world]; Goblin Igniter (5253, -0.39 DPS) [quest]; Captain Rackmore's Tiller (16789, -0.45 DPS, sim-verified) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Battle Healer's Cloak; chest: Stormcloth Vest; hands: Dreamweave Gloves; waist: Gilded Cord; legs: Stoneweaver Leggings; feet: Thoughtcast Boots; finger2: Solid Gold Band; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Hand of Righteousness; off_hand: Prophetic Cane; ranged: Jaina's Firestarter

No-known-source sample (15 of 454, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger; 5743 Prismstone Ring

### Band 50 (undead, 005005031330011430-33501000000000000-00000000000000000)

Set DPS (verified): 211.5. Weights run: 40.3s. Verify run: 346.4s. 601 eligible items had no known source.

Race sweep: undead 211.5 ± 0.1, troll 195.2 ± 0.1 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.097, intellect=1.102 ± 0.030, spirit=0.589 ± 0.028, mp5=2.457 ± 0.029, crit=0.489 ± 0.022 per rating point (14 rating = 1%, 6.849 per %), spell_haste=3.056 ± 0.554

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Blood Guard's Satin Cover (220899, +0.00 DPS) [vendor]; Soulcatcher Halo (10630, -0.08 DPS) [dungeon]; Cassandra's Grace (13102, -0.31 DPS, sim-verified) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 28.8 healing_power points (4.80 DPS) | yes | Gemshard Heart (17707, -0.65 DPS) [dungeon]; Horizon Choker (13085, -1.84 DPS) [world_drop]; Mindburst Medallion (11196, -2.21 DPS) [quest] |
| shoulder | Mantle of Lost Hope (22234) | Blackrock Depths: Lord Roccor [dungeon] | 47.9 healing_power points (7.98 DPS) | yes | Blood Guard's Satin Pads (220901, -0.85 DPS) [vendor]; Nethergeld Shoulders (254049, -1.09 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -1.93 DPS) [dungeon] |
| back | Battle Healer's Cloak (19527) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -0.23 DPS) [dungeon]; Deep Woodlands Cloak (19121, -0.60 DPS) [quest]; Darkspear Raider's Cloak (272076, -1.19 DPS) [vendor] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | sim-verified (+5.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Robes of Insight (940, -0.38 DPS) [world_drop]; Stone Guard's Satin Armor (220903, -2.02 DPS) [vendor]; Stormcloth Vest (10020, -5.91 DPS, sim-verified) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 30.7 healing_power points (5.10 DPS) | yes | Aristocratic Cuffs (12546, -1.48 DPS, sim-verified) [dungeon]; Bloodband Bracers (11469, -2.33 DPS) [quest]; Spidertank Oilrag (9448, -2.38 DPS) [dungeon] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 44.7 healing_power points (7.45 DPS) | yes | First Sergeant's Satin Gloves (220898, -1.55 DPS) [vendor]; Greenleaf Handwraps (19116, -1.77 DPS) [quest]; Jumanza Grips (18083, -1.94 DPS, sim-verified) [dungeon] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 38.9 healing_power points (6.48 DPS) | yes | Gilded Waistcord (254081, -0.00 DPS) [crafted]; Gilded Cord (254037, -0.76 DPS) [crafted]; Earthenweave Cord (254077, -1.07 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 55.3 healing_power points (9.21 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -2.64 DPS) [dungeon]; Stone Guard's Satin Leggings (220902, -2.70 DPS, sim-verified) [vendor]; Stoneweaver Leggings (9407, -3.23 DPS) [dungeon] |
| feet | Vinerot Sandals (17748) | Maraudon: Razorlash [dungeon] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | First Sergeant's Satin Boots (220900, -0.10 DPS) [vendor]; Earthenweave Boots (254093, -0.22 DPS) [crafted]; Gilded Sandals (254107, -1.79 DPS, sim-verified) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 25.0 healing_power points (4.16 DPS) | yes | Runemender's Seal (273031, -0.97 DPS) [world]; Cyclopean Band (11824, -0.98 DPS) [dungeon]; Darkspear Signet (272069, -1.29 DPS) [vendor] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 21.8 healing_power points (3.63 DPS) | yes | Runemender's Seal (273031, -0.36 DPS, sim-verified) [world]; Cyclopean Band (11824, -0.46 DPS) [dungeon]; Darkspear Signet (272069, -0.77 DPS) [vendor] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+5.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Blessed Prayer Beads (19990, +0.00 DPS) [quest]; Evonice's Landin' Pilla (18951, -0.20 DPS) [quest] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Blessed Prayer Beads (19990, -0.21 DPS) [quest]; Evonice's Landin' Pilla (18951, -0.27 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Resurgence Rod (17743, -5.81 DPS, sim-verified) [quest]; Charstone Dirk (17710, -9.03 DPS) [dungeon]; Dreamstaff (249454, -11.83 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Earthen Rod (9381, -0.05 DPS) [dungeon]; Dancing Flame (6806, -0.15 DPS) [quest]; Charged Lightning Rod (11860, -0.59 DPS, sim-verified) [quest] |

**New at 50:** shoulder: Mantle of Lost Hope; back: Battle Healer's Cloak; chest: Embrace of the Wind Serpent; wrist: Nethergeld Cuffs; hands: Gilded Gloves; waist: Dawnspire Cord; legs: Senior Designer's Pantaloons; feet: Vinerot Sandals; finger1: Brainlash; finger2: Advisor's Ring; trinket2: Darkspear Voodoo Seal; main_hand: Glowing Brightwood Staff; ranged: Wand of Mana Concentration

No-known-source sample (15 of 601, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60 (undead, 005005031330011430-33505003030000000-00000000000000000)

Set DPS (verified): 296.1. Weights run: 41.0s. Verify run: 1143.6s. 1520 eligible items had no known source.

Race sweep: undead 296.1 ± 0.2, troll 277.7 ± 0.2 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.111, intellect=2.138 ± 0.069, spirit=2.706 ± 0.050, mp5=5.192 ± 0.079, crit=0.506 ± 0.025 per rating point (14 rating = 1%, 7.090 per %), spell_haste=not significant (-2.397 ± 1.017)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | 111.9 healing_power points (10.90 DPS) | yes | Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 83.5 healing_power points (8.14 DPS) | yes | Tooth of Gnarr (13141, -2.87 DPS) [dungeon]; Dragon's Eye (10829, -2.92 DPS) [dungeon]; Animated Chain Necklace (18723, -3.08 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | 80.2 healing_power points (7.82 DPS) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp]; Mantle of the Scarlet Crusade (22405, -0.21 DPS) [dungeon] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, +0.00 DPS) [quest]; Gracious Cape (18743, +0.00 DPS) [dungeon]; Darkspear Raider's Cloak (272063, -0.13 DPS) [vendor] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 135.5 healing_power points (13.21 DPS) | yes | Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Robes of the Exalted (13346, -1.05 DPS) [dungeon]; Mooncloth Vest (14138, -1.50 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wyrmthalak's Shackles (13958, +0.00 DPS) [quest]; General's Satin Bracers (17619, +0.00 DPS) [pvp]; Dryad's Wrist Bindings (19595, +0.00 DPS) [rep] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of the Exalted Herald (12554, +0.00 DPS) [dungeon]; Raider Handwraps (272097, -0.06 DPS) [vendor]; Magister's Gloves (16684, -0.74 DPS) [dungeon] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Satin Cinch (17621, +0.00 DPS) [pvp]; Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Penitent's Cinch (272394, +0.00 DPS) [vendor] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 118.3 healing_power points (11.53 DPS) | yes | General's Satin Legguards (231634, +0.00 DPS) [vendor]; Legionnaire's Satin Legguards (227123, -0.81 DPS) [pvp]; Senior Designer's Pantaloons (11841, -1.72 DPS) [dungeon] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Satin Walkers (231630, -0.65 DPS) [vendor]; Faith Healer's Boots (22247, -0.68 DPS) [dungeon]; Devout Sandals (16691, -1.05 DPS) [dungeon] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, -0.69 DPS) [dungeon]; The Postmaster's Seal (13392, -1.01 DPS) [dungeon]; Channeler's Ring (272406, -1.04 DPS) [vendor] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, +0.00 DPS) [dungeon]; The Postmaster's Seal (13392, -0.07 DPS) [dungeon]; Channeler's Ring (272406, -0.10 DPS) [vendor] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -0.13 DPS) [quest]; Evonice's Landin' Pilla (18951, -0.53 DPS) [quest]; Darkspear Voodoo Seal (272061, -0.63 DPS) [vendor] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -1.01 DPS) [quest]; Evonice's Landin' Pilla (18951, -1.41 DPS) [quest]; Darkspear Voodoo Seal (272061, -1.52 DPS) [vendor] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Mace (23464, +0.00 DPS) [vendor]; Staff of Metanoia (22394, -0.95 DPS) [dungeon]; Argent Crusader (13249, -0.99 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Eternal Light (22254) | Blackrock Depths: Watchman Doomgrip [dungeon] | 32.2 healing_power points (3.14 DPS) | yes | Oblivion's Touch (18761, -0.80 DPS) [dungeon]; Brilliant Wand (249385, -0.80 DPS) [crafted]; Bonecreeper Stylus (13938, -1.03 DPS) [dungeon] |

**New at 60:** head: Devout Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; back: Drape of Recovery; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of Mending; finger2: Blessed Band of Light; trinket2: Burst of Knowledge; ranged: Wand of Eternal Light

No-known-source sample (15 of 1520, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60, raid preset (undead, 005005031330011430-33505003030000000-00000000000000000)

Set DPS (verified): 620.3. Weights run: 27.4s. Verify run: 707.8s. 1520 eligible items had no known source.

Race sweep: undead 619.9 ± 0.4, troll 603.8 ± 0.3 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.233, intellect=1.323 ± 0.061, spirit=0.985 ± 0.063, mp5=2.811 ± 0.069, crit=0.867 ± 0.049 per rating point (14 rating = 1%, 12.136 per %), spell_haste=not significant (0.840 ± 1.352)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | 66.5 healing_power points (18.37 DPS) | yes | Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -13.98 DPS, sim-verified) [crafted] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 51.7 healing_power points (14.27 DPS) | yes | Tooth of Gnarr (13141, -3.78 DPS) [dungeon]; Lady Maye's Pendant (14558, -4.29 DPS) [world_drop]; Animated Chain Necklace (18723, -4.52 DPS, sim-verified) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp]; Mantle of Lost Hope (22234, -11.91 DPS, sim-verified) [dungeon] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gracious Cape (18743, -2.52 DPS) [dungeon]; Battle Healer's Cloak (19526, -2.71 DPS, sim-verified) [rep]; Shroud of the Exile (15421, -2.79 DPS) [quest] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 88.8 healing_power points (24.51 DPS) | yes | Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Warlord's Satin Robes (231612, -2.06 DPS) [pvp]; Embrace of the Wind Serpent (12462, -5.54 DPS, sim-verified) [world] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Funeral Cuffs (12626, +0.00 DPS) [dungeon]; General's Satin Bracers (17619, +0.00 DPS) [pvp]; Dryad's Wrist Bindings (19595, -11.41 DPS, sim-verified) [rep] |
| hands | Hands of the Exalted Herald (12554) | Blackrock Depths: Princess Moira Bronzebeard  [dungeon] | 60.0 healing_power points (16.58 DPS) | yes | General's Satin Handwraps (231633, +0.00 DPS) [vendor]; Raider Handwraps (272098, -2.09 DPS) [vendor] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Penitent's Cinch (272394, +0.00 DPS) [vendor]; General's Satin Cinch (17621, -0.31 DPS) [pvp]; Wisdom of the Timbermaw (19047, -10.72 DPS, sim-verified) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 66.5 healing_power points (18.37 DPS) | yes | Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Senior Designer's Pantaloons (11841, -15.92 DPS, sim-verified) [dungeon] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 55.3 healing_power points (15.28 DPS) | yes | General's Satin Walkers (231630, +0.00 DPS) [vendor]; Faith Healer's Boots (22247, -1.44 DPS, sim-verified) [dungeon]; General's Satin Boots (17618, -2.35 DPS) [vendor] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, -1.46 DPS) [dungeon]; Blessed Band of Light (272407, -2.01 DPS) [vendor]; Naglering (11669, -14.76 DPS, sim-verified) [dungeon] |
| finger2 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -1.67 DPS) [vendor]; Rosewine Circle (13178, -2.69 DPS, sim-verified) [dungeon]; Maiden's Circle (13001, -3.25 DPS) [world_drop] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (+8.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Burst of Knowledge (11832, +0.00 DPS) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Uther's Strength (11302, -0.83 DPS) [world_drop]; Weakness Analyzer (272438, -0.91 DPS, sim-verified) [vendor] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Argent Crusader (13249, -3.73 DPS, sim-verified) [quest]; Redemption (22406, -5.71 DPS) [dungeon]; Staff of Metanoia (22394, -5.72 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Eternal Light (22254) | Blackrock Depths: Watchman Doomgrip [dungeon] | 22.6 healing_power points (6.23 DPS) | yes | Brilliant Wand (249385, -1.19 DPS) [crafted]; Bonecreeper Stylus (13938, -1.37 DPS) [dungeon]; Oblivion's Touch (18761, -2.43 DPS, sim-verified) [dungeon] |

**New at 60:** head: Devout Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; back: Drape of Recovery; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Hands of the Exalted Herald; waist: Devout Belt; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of Mending; finger2: Fordring's Seal; trinket1: Dormant Heart of the Mountain; trinket2: Serenity Field; ranged: Wand of Eternal Light

No-known-source sample (15 of 1520, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

