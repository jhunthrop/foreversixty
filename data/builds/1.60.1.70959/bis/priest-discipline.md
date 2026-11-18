# Leveling BiS: Discipline

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 005005000010000000-00000000000000000-00000000000000000)

Set DPS (verified): 46.4. Weights run: 14.7s. Verify run: 63.5s. 203 eligible items had no known source.

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

Set DPS (verified): 93.7. Weights run: 15.1s. Verify run: 81.6s. 325 eligible items had no known source.

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

Set DPS (verified): 143.2. Weights run: 24.0s. Verify run: 152.7s. 461 eligible items had no known source.

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

Set DPS (verified): 209.3. Weights run: 25.4s. Verify run: 218.7s. 607 eligible items had no known source.

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

No-known-source sample (15 of 607, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (human, 005005031330011430-33505003030000000-00000000000000000)

Set DPS (verified): 299.3. Weights run: 25.3s. Verify run: 652.5s. 1503 eligible items had no known source.

Race sweep: human 299.3 ± 0.2, gnome 295.5 ± 0.2, night-elf 284.1 ± 0.2, dwarf 281.4 ± 0.2 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.111, intellect=2.138 ± 0.069, spirit=2.706 ± 0.050, mp5=5.192 ± 0.079, crit=0.506 ± 0.025 per rating point (14 rating = 1%, 7.090 per %), spell_haste=not significant (-2.397 ± 1.017)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 119.7 healing_power points (11.66 DPS) | yes | Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Lieutenant Commander's Satin Hood (227121, -0.33 DPS) [pvp]; Devout Crown (16693, -5.34 DPS, sim-verified) [dungeon] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 83.5 healing_power points (8.14 DPS) | yes | Amulet of the Redeemed (22327, -2.76 DPS, sim-verified) [dungeon]; Tooth of Gnarr (13141, -2.87 DPS) [dungeon]; Dragon's Eye (10829, -2.92 DPS) [dungeon] |
| shoulder | Mantle of the Scarlet Crusade (22405) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp]; Devout Mantle (16695, -0.36 DPS, sim-verified) [dungeon] |
| back | Gracious Cape (18743) | Stratholme: Ghostly Citizen [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Drape of Recovery (272413, -0.13 DPS) [vendor]; Expeditionary's Cape (271907, -0.26 DPS) [rep]; Shroud of the Exile (15421, -0.60 DPS, sim-verified) [quest] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 135.5 healing_power points (13.21 DPS) | yes | Field Marshal's Satin Tunic (231624, +0.00 DPS) [vendor]; Robes of the Exalted (13346, -1.14 DPS, sim-verified) [dungeon]; Mooncloth Vest (14138, -1.50 DPS) [crafted] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | 67.0 healing_power points (6.53 DPS) | yes | Marshal's Satin Bracers (17606, +0.00 DPS) [pvp]; Dryad's Wrist Bindings (19595, -0.87 DPS) [rep]; Wyrmthalak's Shackles (13958, -6.40 DPS, sim-verified) [quest] |
| hands | Virtuous Mitts (226950) | Just Compensation [quest] | 96.2 healing_power points (9.37 DPS) | yes | Devout Gloves (16692, -1.35 DPS) [dungeon]; Theramore Handwraps (271945, -1.41 DPS) [rep]; Hands of the Exalted Herald (12554, -7.27 DPS, sim-verified) [dungeon] |
| waist | Virtuous Belt (226948) | Just Compensation [quest] | 90.6 healing_power points (8.83 DPS) | yes | Marshal's Satin Sash (17609, -0.09 DPS) [pvp]; Wisdom of the Timbermaw (19047, -0.52 DPS) [crafted]; Penitent's Cinch (272394, -7.64 DPS, sim-verified) [vendor] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 118.3 healing_power points (11.53 DPS) | yes | Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Knight-Captain's Satin Legguards (227125, -0.81 DPS) [pvp]; Senior Designer's Pantaloons (11841, -1.29 DPS, sim-verified) [dungeon] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 89.1 healing_power points (8.68 DPS) | yes | Marshal's Satin Walkers (231627, -0.65 DPS) [vendor]; Faith Healer's Boots (22247, -0.68 DPS) [dungeon]; Devout Sandals (16691, -1.05 DPS) [dungeon] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, -0.69 DPS) [dungeon]; Blessed Band of Light (272407, -0.93 DPS) [vendor]; Naglering (11669, -8.70 DPS, sim-verified) [dungeon] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, -0.53 DPS) [dungeon]; Blessed Band of Light (272407, -0.77 DPS) [vendor]; The Postmaster's Seal (13392, -0.85 DPS) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+7.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Shard of the Splithooves (10659, -1.01 DPS) [quest]; Evonice's Landin' Pilla (18951, -1.41 DPS) [quest]; Enriched Seal (271903, -1.52 DPS) [rep] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.53 DPS) [quest]; Enriched Seal (271903, -0.63 DPS) [rep]; Shard of the Splithooves (10659, -1.88 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Warhammer (23454, +0.00 DPS) [vendor]; Argent Crusader (13249, -0.99 DPS) [quest]; Staff of Metanoia (22394, -2.94 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Eternal Light (22254) | Blackrock Depths: Watchman Doomgrip [dungeon] | 32.2 healing_power points (3.14 DPS) | yes | Oblivion's Touch (18761, -0.80 DPS) [dungeon]; Brilliant Wand (249385, -0.80 DPS) [crafted]; Cairnstone Sliver (9654, -1.33 DPS, sim-verified) [quest] |

**New at 60:** head: Virtuous Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Mantle of the Scarlet Crusade; back: Gracious Cape; wrist: Virtuous Bracers; hands: Virtuous Mitts; waist: Virtuous Belt; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of Mending; finger2: Channeler's Ring; trinket1: Burst of Knowledge; trinket2: Ankh of Life; ranged: Wand of Eternal Light

No-known-source sample (15 of 1503, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (human, 005005031330011430-33505003030000000-00000000000000000)

Set DPS (verified): 633.8. Weights run: 17.4s. Verify run: 313.0s. 1503 eligible items had no known source.

Race sweep: gnome 633.4 ± 0.4, human 633.4 ± 0.4, night-elf 618.8 ± 0.4, dwarf 616.2 ± 0.4 (heal_score); gnome beats the guide race human by 0.0 (0.00%) ± 0.5, inside error.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.231, intellect=1.015 ± 0.050, spirit=0.834 ± 0.050, mp5=2.345 ± 0.056, crit=0.644 ± 0.041 per rating point (14 rating = 1%, 9.011 per %), spell_haste=not significant (1.527 ± 1.105)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 72.6 healing_power points (24.55 DPS) | yes | Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Field Marshal's Headdress (17602, -0.61 DPS) [vendor]; Mooncloth Circlet (14140, -5.48 DPS, sim-verified) [crafted] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 44.6 healing_power points (15.08 DPS) | yes | Tooth of Gnarr (13141, -3.86 DPS) [dungeon]; Animated Chain Necklace (18723, -4.32 DPS, sim-verified) [dungeon]; Amulet of the Redeemed (22327, -4.32 DPS) [dungeon] |
| shoulder | Mantle of Lost Hope (22234) | Blackrock Depths: Lord Roccor [dungeon] | 46.5 healing_power points (15.74 DPS) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp]; Mooncloth Shoulders (14139, -0.76 DPS, sim-verified) [crafted] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Caretaker's Cape (19530, -1.70 DPS, sim-verified) [rep]; Gracious Cape (18743, -4.13 DPS) [dungeon]; Shroud of the Exile (15421, -4.19 DPS) [quest] |
| chest | Virtuous Robe (226945) | Saving the Best for Last [quest] | sim-verified (+5.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Field Marshal's Satin Tunic (231624, +0.00 DPS) [vendor]; Field Marshal's Satin Vestments (17605, -1.34 DPS) [vendor]; Truefaith Vestments (14154, -5.61 DPS, sim-verified) [crafted] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | 36.3 healing_power points (12.28 DPS) | yes | Dryad's Wrist Bindings (19595, -1.19 DPS, sim-verified) [rep]; Wyrmthalak's Shackles (13958, -1.28 DPS) [quest]; Funeral Cuffs (12626, -1.39 DPS) [dungeon] |
| hands | Hands of the Exalted Herald (12554) | Blackrock Depths: Princess Moira Bronzebeard  [dungeon] | 54.2 healing_power points (18.33 DPS) | yes | Marshal's Satin Handwraps (231623, +0.00 DPS) [vendor]; Virtuous Mitts (226950, -1.36 DPS) [quest]; Theramore Handwraps (271945, -2.03 DPS, sim-verified) [rep] |
| waist | Virtuous Belt (226948) | Just Compensation [quest] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Penitent's Cinch (272394, -1.65 DPS) [vendor]; Wisdom of the Timbermaw (19047, -2.41 DPS, sim-verified) [crafted]; Devout Belt (16696, -3.07 DPS) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 60.2 healing_power points (20.35 DPS) | yes | Knight-Captain's Satin Legguards (227125, +0.00 DPS) [pvp]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Senior Designer's Pantaloons (11841, -2.12 DPS, sim-verified) [dungeon] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 49.0 healing_power points (16.59 DPS) | yes | Marshal's Satin Walkers (231627, +0.00 DPS) [vendor]; Faith Healer's Boots (22247, -1.05 DPS, sim-verified) [dungeon]; Swarmtender's Footpads (275609, -2.04 DPS) [crafted] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, -1.70 DPS) [dungeon]; Blessed Band of Light (272407, -1.84 DPS) [vendor]; Naglering (11669, -13.33 DPS, sim-verified) [dungeon] |
| finger2 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -1.83 DPS) [vendor]; Rosewine Circle (13178, -2.12 DPS, sim-verified) [dungeon]; Maiden's Circle (13001, -4.39 DPS) [world_drop] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.26 DPS) [world_drop]; Weakness Analyzer (272438, -0.26 DPS) [vendor]; Serenity Field (272439, -1.57 DPS, sim-verified) [vendor] |
| trinket2 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, -2.09 DPS, sim-verified) [vendor] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Argent Crusader (13249, -2.99 DPS, sim-verified) [quest]; Redemption (22406, -6.01 DPS) [dungeon]; Staff of Metanoia (22394, -6.21 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Eternal Light (22254) | Blackrock Depths: Watchman Doomgrip [dungeon] | 19.8 healing_power points (6.69 DPS) | yes | Brilliant Wand (249385, -1.24 DPS) [crafted]; Bonecreeper Stylus (13938, -1.25 DPS) [dungeon]; Oblivion's Touch (18761, -1.61 DPS, sim-verified) [dungeon] |

**New at 60:** head: Virtuous Crown; neck: Jeweled Amulet of Cainwyn; back: Drape of Recovery; chest: Virtuous Robe; wrist: Virtuous Bracers; hands: Hands of the Exalted Herald; waist: Virtuous Belt; legs: Sentinel's Silk Leggings; feet: Mooncloth Boots; finger1: Band of Mending; finger2: Fordring's Seal; trinket1: Burst of Knowledge; trinket2: Dormant Heart of the Mountain; ranged: Wand of Eternal Light

No-known-source sample (15 of 1503, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

## Horde

### Band 20 (undead, 005005000010000000-00000000000000000-00000000000000000)

Set DPS (verified): 45.6. Weights run: 14.7s. Verify run: 58.0s. 196 eligible items had no known source.

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

Set DPS (verified): 95.0. Weights run: 15.1s. Verify run: 78.3s. 318 eligible items had no known source.

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

Set DPS (verified): 147.1. Weights run: 24.0s. Verify run: 148.9s. 454 eligible items had no known source.

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
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 26.3 healing_power points (4.34 DPS) | yes | Advisor's Ring (19520, -1.35 DPS) [rep]; Welken Ring (5011, -1.40 DPS) [world_drop]; Electrocutioner Lagnut (9447, -2.28 DPS, sim-verified) [dungeon] |
| finger2 | Solid Gold Band (279845) | Opportunistic Education [quest] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Advisor's Ring (19520, -0.23 DPS) [rep]; Electrocutioner Lagnut (9447, -0.27 DPS, sim-verified) [dungeon]; Welken Ring (5011, -0.28 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+3.9 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Jordan (873, -4.70 DPS, sim-verified) [world_drop]; Staff of the Friar (3415, -5.21 DPS) [dungeon]; Odo's Ley Staff (6318, -5.80 DPS) [dungeon] |
| off_hand | Prophetic Cane (6803) | Into The Scarlet Monastery [quest] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Orb of Souls (249395, -0.69 DPS) [crafted]; Beacon of Hope (9393, -0.86 DPS, sim-verified) [dungeon]; Omega Orb (7749, -1.54 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 12.2 healing_power points (2.01 DPS) | yes | Wand of Mana Concentration (273043, -0.23 DPS) [world]; Goblin Igniter (5253, -0.39 DPS) [quest]; Captain Rackmore's Tiller (16789, -0.45 DPS, sim-verified) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Battle Healer's Cloak; chest: Stormcloth Vest; hands: Dreamweave Gloves; waist: Gilded Cord; legs: Stoneweaver Leggings; feet: Thoughtcast Boots; finger2: Solid Gold Band; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Hand of Righteousness; off_hand: Prophetic Cane; ranged: Jaina's Firestarter

No-known-source sample (15 of 454, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger; 5743 Prismstone Ring

### Band 50 (undead, 005005031330011430-33501000000000000-00000000000000000)

Set DPS (verified): 212.0. Weights run: 25.4s. Verify run: 212.8s. 600 eligible items had no known source.

Race sweep: undead 212.0 ± 0.1, troll 196.7 ± 0.1 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.046, intellect=1.334 ± 0.030, spirit=2.441 ± 0.027, mp5=5.744 ± 0.031, crit=0.361 ± 0.014 per rating point (14 rating = 1%, 5.047 per %), spell_haste=not significant (-1.420 ± 0.496)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 79.6 healing_power points (7.41 DPS) | yes | Blood Guard's Satin Cover (220899, -0.42 DPS) [vendor]; Papal Fez (9431, -0.52 DPS, sim-verified) [dungeon]; Cassandra's Grace (13102, -1.27 DPS) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Horizon Choker (13085, -0.21 DPS) [world_drop]; Gemshard Heart (17707, -0.34 DPS) [dungeon]; Lei of Lilies (1315, -0.97 DPS, sim-verified) [world_drop] |
| shoulder | Mantle of Lost Hope (22234) | Blackrock Depths: Lord Roccor [dungeon] | 63.6 healing_power points (5.92 DPS) | yes | Blood Guard's Satin Pads (220901, -0.50 DPS) [vendor]; Nethergeld Shoulders (254049, -1.23 DPS, sim-verified) [crafted]; Kentic Amice (11624, -1.64 DPS) [dungeon] |
| back | Featherskin Cape (10843) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -0.83 DPS) [dungeon]; Battle Healer's Cloak (19527, -0.83 DPS) [rep]; Wingveil Cloak (10802, -1.09 DPS) [dungeon] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 119.6 healing_power points (11.13 DPS) | yes | Stormcloth Vest (10020, -3.91 DPS) [crafted]; Grimlok's Tribal Vestments (9415, -4.63 DPS) [dungeon]; Ghostweave Vest (14141, -4.77 DPS, sim-verified) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 41.5 healing_power points (3.87 DPS) | yes | Spidertank Oilrag (9448, -1.42 DPS) [dungeon]; Aristocratic Cuffs (12546, -1.46 DPS, sim-verified) [dungeon]; Venomshroud Armguards (14439, -1.55 DPS) [world_drop] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 61.9 healing_power points (5.76 DPS) | yes | Earthenweave Gloves (254075, -0.84 DPS) [crafted]; Jumanza Grips (18083, -1.08 DPS) [dungeon]; Gloves of the Atal'ai Prophet (10808, -3.91 DPS, sim-verified) [dungeon] |
| waist | Earthenweave Cord (254077) | Tailoring [crafted] | 52.8 healing_power points (4.92 DPS) | yes | Gilded Cord (254037, -0.27 DPS, sim-verified) [crafted]; Dawnspire Cord (12466, -0.88 DPS) [dungeon]; Gilded Waistcord (254081, -1.10 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 90.3 healing_power points (8.40 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -1.24 DPS) [dungeon]; Dalewind Trousers (13008, -1.61 DPS) [world_drop]; Rainstrider Leggings (11123, -3.70 DPS, sim-verified) [quest] |
| feet | Vinerot Sandals (17748) | Maraudon: Razorlash [dungeon] | 54.4 healing_power points (5.06 DPS) | yes | Earthenweave Boots (254093, -0.15 DPS) [crafted]; Furen's Boots (13100, -0.56 DPS) [world_drop]; Mistwalker Boots (10629, -2.27 DPS, sim-verified) [dungeon] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 37.3 healing_power points (3.47 DPS) | yes | Advisor's Ring (19519, -0.22 DPS) [rep]; Runemender's Seal (273031, -0.53 DPS) [world]; Cyclopean Band (11824, -0.86 DPS) [dungeon] |
| finger2 | Brainlash (6440) | Zul'Farrak [dungeon] | 36.5 healing_power points (3.40 DPS) | yes | Advisor's Ring (19519, -0.37 DPS, sim-verified) [rep]; Runemender's Seal (273031, -0.46 DPS) [world]; Cyclopean Band (11824, -0.79 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Darkspear Voodoo Seal (272061, -0.05 DPS) [vendor]; Evonice's Landin' Pilla (18951, -0.45 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor]; Dormant Heart of the Mountain (249473, -0.49 DPS, sim-verified) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamstaff (249454, -5.52 DPS) [crafted]; Charstone Dirk (17710, -6.37 DPS) [dungeon]; Resurgence Rod (17743, -11.06 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 16.8 healing_power points (1.57 DPS) | yes | Chillnail Splinter (10704, -0.20 DPS) [quest]; Jaina's Firestarter (13064, -0.24 DPS, sim-verified) [world_drop]; Captain Rackmore's Tiller (16789, -0.33 DPS) [quest] |

**New at 50:** head: Gemburst Circlet; shoulder: Mantle of Lost Hope; back: Featherskin Cape; chest: Embrace of the Wind Serpent; wrist: Nethergeld Cuffs; hands: Gilded Gloves; waist: Earthenweave Cord; legs: Senior Designer's Pantaloons; feet: Vinerot Sandals; finger1: Eye of Adaegus; finger2: Brainlash; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Wand of Mana Concentration

No-known-source sample (15 of 600, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60 (undead, 005005031330011430-33505003030000000-00000000000000000)

Set DPS (verified): 299.6. Weights run: 25.3s. Verify run: 671.0s. 1498 eligible items had no known source.

Race sweep: undead 299.5 ± 0.2, troll 281.0 ± 0.2 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.106, intellect=1.274 ± 0.047, spirit=1.854 ± 0.043, mp5=3.407 ± 0.048, crit=0.643 ± 0.032 per rating point (14 rating = 1%, 9.003 per %), spell_haste=not significant (2.762 ± 1.006)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 93.3 healing_power points (17.42 DPS) | yes | Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Champion's Satin Hood (227118, -0.83 DPS) [pvp]; Mooncloth Circlet (14140, -9.36 DPS, sim-verified) [crafted] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 59.5 healing_power points (11.10 DPS) | yes | Animated Chain Necklace (18723, -2.06 DPS, sim-verified) [dungeon]; Amulet of the Redeemed (22327, -3.15 DPS) [dungeon]; Dragon's Eye (10829, -3.48 DPS) [dungeon] |
| shoulder | Mooncloth Shoulders (14139) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp]; Mantle of the Scarlet Crusade (22405, -0.51 DPS, sim-verified) [dungeon] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battle Healer's Cloak (19526, -0.71 DPS) [rep]; Shroud of the Exile (15421, -1.29 DPS) [quest]; Featherskin Cape (10843, -1.46 DPS) [world] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 102.7 healing_power points (19.16 DPS) | yes | Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Robes of the Exalted (13346, -1.14 DPS) [dungeon]; Embrace of the Wind Serpent (12462, -2.23 DPS, sim-verified) [world] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | 49.9 healing_power points (9.31 DPS) | yes | Dryad's Wrist Bindings (19595, -0.88 DPS) [rep]; General's Satin Bracers (17619, -1.33 DPS) [pvp]; Wyrmthalak's Shackles (13958, -7.56 DPS, sim-verified) [quest] |
| hands | Virtuous Mitts (226950) | Just Compensation [quest] | 71.4 healing_power points (13.32 DPS) | yes | General's Satin Handwraps (231633, -2.02 DPS) [vendor]; Devout Gloves (16692, -2.26 DPS) [dungeon]; Hands of the Exalted Herald (12554, -7.75 DPS, sim-verified) [dungeon] |
| waist | Virtuous Belt (226948) | Just Compensation [quest] | sim-verified (+7.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Wisdom of the Timbermaw (19047, -0.02 DPS) [crafted]; General's Satin Cinch (17621, -1.59 DPS) [pvp]; Penitent's Cinch (272394, -7.55 DPS, sim-verified) [vendor] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 85.7 healing_power points (16.00 DPS) | yes | General's Satin Legguards (231634, +0.00 DPS) [vendor]; Legionnaire's Satin Legguards (227123, -0.13 DPS) [pvp]; Senior Designer's Pantaloons (11841, -1.21 DPS) [dungeon] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 65.9 healing_power points (12.31 DPS) | yes | General's Satin Walkers (231630, +0.00 DPS) [vendor]; Faith Healer's Boots (22247, -0.82 DPS) [dungeon]; Devout Sandals (16691, -1.99 DPS) [dungeon] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -0.97 DPS) [vendor]; Rosewine Circle (13178, -1.23 DPS) [dungeon]; Naglering (11669, -7.01 DPS, sim-verified) [dungeon] |
| finger2 | Fordring's Seal (16058) | In Dreams [quest] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Rosewine Circle (13178, -0.04 DPS) [dungeon]; Channeler's Ring (272406, -0.56 DPS) [vendor]; Blessed Band of Light (272407, -1.78 DPS, sim-verified) [vendor] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+6.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Serenity Field (272439, -0.23 DPS) [vendor]; Shard of the Splithooves (10659, -0.34 DPS) [quest]; Evonice's Landin' Pilla (18951, -0.69 DPS) [quest] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -0.42 DPS, sim-verified) [quest]; Serenity Field (272439, -1.17 DPS) [vendor]; Evonice's Landin' Pilla (18951, -1.63 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Mace (23464, +0.00 DPS) [vendor]; Redemption (22406, -1.77 DPS) [dungeon]; Staff of Metanoia (22394, -4.08 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Eternal Light (22254) | Blackrock Depths: Watchman Doomgrip [dungeon] | 23.5 healing_power points (4.38 DPS) | yes | Oblivion's Touch (18761, -1.03 DPS) [dungeon]; Brilliant Wand (249385, -1.03 DPS) [crafted]; Bonecreeper Stylus (13938, -1.14 DPS) [dungeon] |

**New at 60:** head: Virtuous Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Mooncloth Shoulders; back: Drape of Recovery; chest: Truefaith Vestments; wrist: Virtuous Bracers; hands: Virtuous Mitts; waist: Virtuous Belt; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of Mending; finger2: Fordring's Seal; trinket2: Burst of Knowledge; ranged: Wand of Eternal Light

No-known-source sample (15 of 1498, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60, raid preset (undead, 005005031330011430-33505003030000000-00000000000000000)

Set DPS (verified): 631.3. Weights run: 17.4s. Verify run: 302.1s. 1498 eligible items had no known source.

Race sweep: undead 631.0 ± 0.4, troll 613.9 ± 0.3 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.231, intellect=1.015 ± 0.050, spirit=0.834 ± 0.050, mp5=2.345 ± 0.056, crit=0.644 ± 0.041 per rating point (14 rating = 1%, 9.011 per %), spell_haste=not significant (1.527 ± 1.105)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 72.6 healing_power points (24.55 DPS) | yes | Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Warlord's Satin Cowl (17623, -0.61 DPS) [vendor]; Mooncloth Circlet (14140, -6.80 DPS, sim-verified) [crafted] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 44.6 healing_power points (15.08 DPS) | yes | Tooth of Gnarr (13141, -3.86 DPS) [dungeon]; Amulet of the Redeemed (22327, -4.32 DPS) [dungeon]; Animated Chain Necklace (18723, -5.36 DPS, sim-verified) [dungeon] |
| shoulder | Mantle of Lost Hope (22234) | Blackrock Depths: Lord Roccor [dungeon] | 46.5 healing_power points (15.74 DPS) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp]; Mooncloth Shoulders (14139, -1.79 DPS, sim-verified) [crafted] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battle Healer's Cloak (19526, -2.82 DPS, sim-verified) [rep]; Gracious Cape (18743, -4.13 DPS) [dungeon]; Shroud of the Exile (15421, -4.19 DPS) [quest] |
| chest | Virtuous Robe (226945) | Saving the Best for Last [quest] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Warlord's Satin Robes (231612, -1.34 DPS) [pvp]; Truefaith Vestments (14154, -3.42 DPS, sim-verified) [crafted] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | 36.3 healing_power points (12.28 DPS) | yes | Wyrmthalak's Shackles (13958, -1.28 DPS) [quest]; Funeral Cuffs (12626, -1.39 DPS) [dungeon]; Dryad's Wrist Bindings (19595, -2.54 DPS, sim-verified) [rep] |
| hands | Hands of the Exalted Herald (12554) | Blackrock Depths: Princess Moira Bronzebeard  [dungeon] | 54.2 healing_power points (18.33 DPS) | yes | General's Satin Handwraps (231633, +0.00 DPS) [vendor]; Virtuous Mitts (226950, -1.36 DPS) [quest]; Raider Handwraps (272097, -4.07 DPS, sim-verified) [vendor] |
| waist | Virtuous Belt (226948) | Just Compensation [quest] | sim-verified (+3.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Penitent's Cinch (272394, -1.65 DPS) [vendor]; Devout Belt (16696, -3.07 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -3.71 DPS, sim-verified) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Senior Designer's Pantaloons (11841, -1.11 DPS, sim-verified) [dungeon] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 49.0 healing_power points (16.59 DPS) | yes | General's Satin Walkers (231630, +0.00 DPS) [vendor]; Faith Healer's Boots (22247, -1.16 DPS, sim-verified) [dungeon]; Swarmtender's Footpads (275609, -2.04 DPS) [crafted] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, -1.70 DPS) [dungeon]; Blessed Band of Light (272407, -1.84 DPS) [vendor]; Naglering (11669, -15.86 DPS, sim-verified) [dungeon] |
| finger2 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -1.83 DPS) [vendor]; Rosewine Circle (13178, -4.17 DPS, sim-verified) [dungeon]; Maiden's Circle (13001, -4.39 DPS) [world_drop] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (+7.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Second Wind (11819, +0.00 DPS) [dungeon]; Burst of Knowledge (11832, +0.00 DPS) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -0.75 DPS, sim-verified) [vendor]; Burst of Knowledge (11832, -0.76 DPS) [dungeon]; Uther's Strength (11302, -1.01 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Argent Crusader (13249, -3.63 DPS, sim-verified) [quest]; Redemption (22406, -6.01 DPS) [dungeon]; Staff of Metanoia (22394, -6.21 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Eternal Light (22254) | Blackrock Depths: Watchman Doomgrip [dungeon] | 19.8 healing_power points (6.69 DPS) | yes | Brilliant Wand (249385, -1.24 DPS) [crafted]; Bonecreeper Stylus (13938, -1.25 DPS) [dungeon]; Oblivion's Touch (18761, -3.95 DPS, sim-verified) [dungeon] |

**New at 60:** head: Virtuous Crown; neck: Jeweled Amulet of Cainwyn; back: Drape of Recovery; chest: Virtuous Robe; wrist: Virtuous Bracers; hands: Hands of the Exalted Herald; waist: Virtuous Belt; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of Mending; finger2: Fordring's Seal; trinket1: Dormant Heart of the Mountain; trinket2: Serenity Field; ranged: Wand of Eternal Light

No-known-source sample (15 of 1498, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

