# Leveling BiS: Discipline

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 005005000010000000-00000000000000000-00000000000000000)

Set DPS (verified): 46.4. Weights run: 13.2s. Verify run: 71.2s. 203 eligible items had no known source.

Race sweep: human 46.4 ± 0.0, gnome 45.9 ± 0.0, night-elf 44.8 ± 0.0, dwarf 44.5 ± 0.0 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.027, intellect=0.700 ± 0.005, spirit=0.609 ± 0.006, mp5=2.069 ± 0.017, crit=0.056 ± 0.003 per rating point (14 rating = 1%, 0.781 per %), spell_haste=not significant (0.013 ± 0.043)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (2.21 DPS) | yes | Shadow Goggles (4373, -0.85 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.91 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 1.2 healing_power points (0.15 DPS) | yes | Tarnished Locket (279870, -0.32 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.3 healing_power points (1.39 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.56 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.63 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.90 DPS) [crafted] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.40 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -0.75 DPS) [crafted]; Black Whelp Cloak (7283, -0.89 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 18.3 healing_power points (2.25 DPS) | yes | Corsair's Overshirt (5202, -0.27 DPS, sim-verified) [dungeon]; Robe of the Moccasin (6465, -0.40 DPS) [dungeon]; Bloody Apron (6226, -0.65 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 9.7 healing_power points (1.19 DPS) | yes | Repurposed Hair Band (281256, -0.87 DPS) [quest]; Seer's Cuffs (3645, -0.96 DPS) [dungeon]; Bright Bracers (3647, -1.48 DPS, sim-verified) [world_drop] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 13.1 healing_power points (1.61 DPS) | yes | Serpent Gloves (5970, -0.80 DPS, sim-verified) [dungeon]; Gnoll Casting Gloves (892, -0.87 DPS) [world]; Magefist Gloves (12977, -0.88 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 13.8 healing_power points (1.70 DPS) | yes | Novice Ardent's Sash (253887, -0.44 DPS, sim-verified) [crafted]; Novice Arcanist's Sash (253885, -0.95 DPS) [crafted]; Keller's Girdle (2911, -1.01 DPS) [world_drop] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 24.6 healing_power points (3.03 DPS) | yes | Darkweave Breeches (12987, -1.85 DPS, sim-verified) [world_drop]; Colorful Kilt (10048, -2.04 DPS) [crafted]; Silk-threaded Trousers (1929, -2.17 DPS) [dungeon] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 11.1 healing_power points (1.36 DPS) | yes | Spidersilk Boots (4320, -0.14 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.50 DPS) [crafted]; Kimbra Boots (6191, -0.80 DPS) [quest] |
| finger1 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 9.1 healing_power points (1.12 DPS) | yes | Lavishly Jeweled Ring (1156, -0.61 DPS) [dungeon]; Band of Purification (12996, -0.67 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.86 DPS) [world_drop] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.4 healing_power points (0.79 DPS) | yes | Lavishly Jeweled Ring (1156, -0.15 DPS, sim-verified) [dungeon]; Band of Purification (12996, -0.34 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.53 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Westfall (2042) | The Defias Brotherhood [quest] | 55.2 healing_power points (6.77 DPS) | yes | Living Root (6631, -1.06 DPS, sim-verified) [dungeon]; Staff of the Blessed Seer (2271, -2.59 DPS) [dungeon]; Evocator's Blade (2567, -2.79 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Sizzle Stick (8071) (or Greater Magic Wand (11288), Torchlight Wand (5240)) | Deviate Eradication [quest] | 2.0 healing_power points (0.25 DPS) | yes | Torchlight Wand (5240, +0.00 DPS) [quest]; Greater Magic Wand (11288, +0.00 DPS) [crafted]; Flaring Baton (5326, -0.07 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Staff of Westfall; ranged: Sizzle Stick

No-known-source sample (15 of 203, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6360 Steelscale Crushfish; 6478 Rat Stompers; 8350 The 1 Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (human, 005005031330001000-00000000000000000-00000000000000000)

Set DPS (verified): 94.7. Weights run: 15.1s. Verify run: 52.9s. 325 eligible items had no known source.

Race sweep: human 94.7 ± 0.0, gnome 91.9 ± 0.0, night-elf 89.9 ± 0.0, dwarf 89.2 ± 0.0 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.145, intellect=0.897 ± 0.014, spirit=1.163 ± 0.013, mp5=2.851 ± 0.015, crit=0.168 ± 0.012 per rating point (14 rating = 1%, 2.349 per %), spell_haste=0.961 ± 0.226

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 40.0 healing_power points (6.95 DPS) | yes | Filigreed Pristine Circlet (253975, -3.05 DPS, sim-verified) [crafted]; Pristine Circlet (253949, -3.82 DPS) [crafted]; Nightsky Cowl (4039, -4.07 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 8.2 healing_power points (1.43 DPS) | yes | Crystal Starfire Medallion (5003, -0.20 DPS, sim-verified) [world_drop]; Pendant of Myzrael (4614, -0.22 DPS) [dungeon]; Glowing Green Talisman (5002, -0.22 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 17.1 healing_power points (2.97 DPS) | yes | Death Speaker Mantle (6685, -0.21 DPS, sim-verified) [dungeon]; Mantle of Honor (3560, -0.46 DPS) [quest]; Magician's Mantle (12998, -0.70 DPS) [world_drop] |
| back | Prelacy Cape (7004) | Researching the Corruption [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Caretaker's Cape (19533, -0.15 DPS, sim-verified) [rep]; Glowing Thresher Cape (6901, -0.29 DPS) [dungeon]; Repairman's Cape (9605, -0.60 DPS) [quest] |
| chest | Pristine Gown (253961) | Tailoring [crafted] | 30.9 healing_power points (5.38 DPS) | yes | Death Speaker Robes (6682, -0.26 DPS, sim-verified) [dungeon]; Beguiler Robes (7728, -0.32 DPS) [dungeon]; Robes of Arugal (6324, -1.04 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 17.6 healing_power points (3.05 DPS) | yes | Mindthrust Bracers (1974, -0.77 DPS, sim-verified) [dungeon]; Nightsky Wristbands (6407, -1.51 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.81 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 18.4 healing_power points (3.21 DPS) | yes | Shilly Mitts (9609, -0.17 DPS) [quest]; Truefaith Gloves (7049, -0.61 DPS, sim-verified) [crafted]; Pristine Gloves (253913, -0.83 DPS) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 17.5 healing_power points (3.05 DPS) | yes | Pristine Sash (253925, -0.65 DPS, sim-verified) [crafted]; Highlander's Cloth Girdle (20099, -0.66 DPS) [rep]; Crimson Silk Belt (7055, -0.91 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 34.1 healing_power points (5.93 DPS) | yes | Filigreed Pristine Leggings (253937, -1.08 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -1.10 DPS) [crafted]; Gaze Dreamer Pants (6903, -2.02 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 30.9 healing_power points (5.38 DPS) | yes | Acidic Walkers (9454, -2.33 DPS, sim-verified) [dungeon]; Boots of the Enchanter (4325, -2.89 DPS) [crafted]; Glinteye Slippers (273024, -2.95 DPS) [dungeon] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 22.5 healing_power points (3.92 DPS) | yes | Electrocutioner Lagnut (9447, -0.94 DPS) [dungeon]; Ring of Power Regulation (273030, -1.39 DPS) [dungeon]; Lorekeeper's Ring (19525, -1.71 DPS) [rep] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 20.4 healing_power points (3.55 DPS) | yes | Electrocutioner Lagnut (9447, -0.63 DPS, sim-verified) [dungeon]; Ring of Power Regulation (273030, -1.02 DPS) [dungeon]; Lorekeeper's Ring (19525, -1.34 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Odo's Ley Staff (6318, -0.36 DPS) [dungeon]; Staff of Westfall (2042, -1.48 DPS) [quest]; Staff of the Friar (3415, -3.11 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 20.0 healing_power points (3.48 DPS) | yes | Orb of Mistmantle (13031, -1.10 DPS, sim-verified) [world_drop]; Eye of Paleth (2943, -1.21 DPS) [quest]; Orb of Mystic Insight (249394, -1.32 DPS) [crafted] |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 9.3 healing_power points (1.62 DPS) | yes | Thunderwood (13062, -0.68 DPS, sim-verified) [world_drop]; Wand of Eventide (5214, -0.75 DPS) [world_drop]; Greater Mystic Wand (217287, -0.75 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Prelacy Cape; chest: Pristine Gown; wrist: Spidertank Oilrag; hands: Gloves of Old; waist: Belt of Arugal; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Runemender's Seal; finger2: Snake Hoop; main_hand: Death Speaker Scepter; off_hand: Orb of Souls; ranged: Wand of Mana Concentration

No-known-source sample (15 of 325, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6360 Steelscale Crushfish

### Band 40 (human, 005005031330011000-33300000000000000-00000000000000000)

Set DPS (verified): 144.5. Weights run: 25.3s. Verify run: 166.5s. 461 eligible items had no known source.

Race sweep: human 144.5 ± 0.1, gnome 141.1 ± 0.0, night-elf 138.4 ± 0.0, dwarf 137.6 ± 0.0 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.134, intellect=1.395 ± 0.026, spirit=1.807 ± 0.029, mp5=3.525 ± 0.027, crit=0.358 ± 0.019 per rating point (14 rating = 1%, 5.016 per %), spell_haste=3.998 ± 0.484

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 62.0 healing_power points (8.80 DPS) | yes | Whitemane's Chapeau (7720, -0.64 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -2.57 DPS) [world_drop]; Earthen Silk Hood (254015, -3.11 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 31.2 healing_power points (4.42 DPS) | yes | Necklace of Calisea (1714, -1.15 DPS, sim-verified) [dungeon]; Triune Amulet (7722, -1.49 DPS) [dungeon]; Amberglow Talisman (10824, -1.86 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 40.1 healing_power points (5.70 DPS) | yes | Earthen Silk Shoulders (254033, -0.81 DPS, sim-verified) [crafted]; Mantle of Doan (7712, -1.46 DPS) [dungeon]; Mistscape Mantle (4734, -2.24 DPS) [dungeon] |
| back | Caretaker's Cape (19532) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Prelacy Cape (7004, -0.74 DPS) [quest]; Glowing Thresher Cape (6901, -0.79 DPS) [dungeon]; Expeditionary's Cape (271922, -0.88 DPS, sim-verified) [rep] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 72.5 healing_power points (10.28 DPS) | yes | Dreamweave Vest (10021, -2.43 DPS, sim-verified) [crafted]; Death's Head Vestment (10581, -2.81 DPS) [dungeon]; Doomsayer's Robe (4746, -3.40 DPS) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 19.6 healing_power points (2.78 DPS) | yes | Earthen Silk Cuffs (254019, -0.13 DPS, sim-verified) [crafted]; Mindthrust Bracers (1974, -0.29 DPS) [dungeon]; Silkstream Cuffs (16791, -0.76 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 36.2 healing_power points (5.14 DPS) | yes | Earthen Silk Gloves (254017, -0.44 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -0.63 DPS) [crafted]; Gloves of Old (9395, -1.07 DPS) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 44.0 healing_power points (6.25 DPS) | yes | Deathmage Sash (10771, -1.82 DPS, sim-verified) [dungeon]; Windchaser Cinch (14435, -2.35 DPS) [world_drop]; Sutarn's Ring (13105, -2.73 DPS) [world_drop] |
| legs | Stoneweaver Leggings (9407) | Uldaman: Ironaya [dungeon] | 52.8 healing_power points (7.50 DPS) | yes | Stormcloth Pants (10010, -1.24 DPS, sim-verified) [crafted]; Pristine Leggings (253987, -1.71 DPS) [crafted]; Earthen Leggings (253999, -2.55 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Thoughtcast Boots (10578, -0.07 DPS) [dungeon]; Earthen Silk Slippers (254013, -1.04 DPS) [crafted]; Furen's Boots (13100, -1.23 DPS, sim-verified) [world_drop] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 30.0 healing_power points (4.26 DPS) | yes | Electrocutioner Lagnut (9447, -1.19 DPS) [dungeon]; Welken Ring (5011, -1.48 DPS) [world_drop]; Lorekeeper's Ring (19524, -1.48 DPS) [rep] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 28.4 healing_power points (4.03 DPS) | yes | Electrocutioner Lagnut (9447, -0.46 DPS, sim-verified) [dungeon]; Welken Ring (5011, -1.25 DPS) [world_drop]; Lorekeeper's Ring (19524, -1.25 DPS) [rep] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+4.3 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of the Friar (3415, -4.34 DPS) [dungeon]; Odo's Ley Staff (6318, -4.94 DPS) [dungeon]; Staff of Jordan (873, -5.13 DPS, sim-verified) [world_drop] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 31.9 healing_power points (4.53 DPS) | yes | Beacon of Hope (9393, -0.39 DPS) [dungeon]; Orb of Souls (249395, -1.15 DPS) [crafted]; Aurora Sphere (7610, -2.12 DPS) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 13.8 healing_power points (1.96 DPS) | yes | Wand of Mana Concentration (273043, -0.16 DPS) [world]; Captain Rackmore's Tiller (16789, -0.34 DPS) [quest]; Goblin Igniter (5253, -0.40 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Caretaker's Cape; chest: Stormcloth Vest; hands: Dreamweave Gloves; waist: Gilded Cord; legs: Stoneweaver Leggings; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Hand of Righteousness; off_hand: Orb of Lorica; ranged: Jaina's Firestarter

No-known-source sample (15 of 461, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (human, 005005031330011430-33501000000000000-00000000000000000)

Set DPS (verified): 207.0. Weights run: 26.6s. Verify run: 229.5s. 606 eligible items had no known source.

Race sweep: human 207.0 ± 0.1, gnome 199.6 ± 0.1, night-elf 196.2 ± 0.1, dwarf 194.7 ± 0.1 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.116, intellect=1.117 ± 0.027, spirit=0.995 ± 0.030, mp5=1.951 ± 0.032, crit=0.413 ± 0.020 per rating point (14 rating = 1%, 5.777 per %), spell_haste=not significant (1.835 ± 0.506)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Cassandra's Grace (13102) | World drop [world_drop] | 53.0 healing_power points (9.29 DPS) | yes | Knight-Lieutenant's Satin Cover (220896, +0.00 DPS) [vendor]; Papal Fez (9431, -0.18 DPS, sim-verified) [dungeon]; Gemburst Circlet (10751, -0.54 DPS) [quest] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 28.9 healing_power points (5.08 DPS) | yes | Gemshard Heart (17707, -0.62 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -1.63 DPS) [world_drop]; Mindburst Medallion (11196, -1.98 DPS) [quest] |
| shoulder | Mantle of Lost Hope (22234) | Blackrock Depths: Lord Roccor [dungeon] | 46.1 healing_power points (8.09 DPS) | yes | Knight-Lieutenant's Satin Pads (220894, -0.90 DPS) [vendor]; Nethergeld Shoulders (254049, -0.95 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -1.68 DPS) [dungeon] |
| back | Caretaker's Cape (19531) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -0.23 DPS) [dungeon]; Featherskin Cape (10843, -0.71 DPS) [world]; Expeditionary's Cape (271921, -1.29 DPS) [rep] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 74.0 healing_power points (12.99 DPS) | yes | Robes of Insight (940, -2.49 DPS) [world_drop]; Stormcloth Vest (10020, -4.24 DPS, sim-verified) [crafted]; Grimlok's Tribal Vestments (9415, -4.48 DPS) [dungeon] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 32.8 healing_power points (5.75 DPS) | yes | Aristocratic Cuffs (12546, -0.44 DPS, sim-verified) [dungeon]; Bloodband Bracers (11469, -2.59 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -2.88 DPS) [quest] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 48.1 healing_power points (8.45 DPS) | yes | Earthenweave Gloves (254075, -1.05 DPS, sim-verified) [crafted]; Sergeant Major's Satin Gloves (220897, -2.22 DPS) [vendor]; Jumanza Grips (18083, -2.55 DPS) [dungeon] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 39.2 healing_power points (6.88 DPS) | yes | Gilded Waistcord (254081, -0.17 DPS, sim-verified) [crafted]; Earthenweave Cord (254077, -0.40 DPS) [crafted]; Gilded Cord (254037, -0.41 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 62.7 healing_power points (11.01 DPS) | yes | Knight's Satin Leggings (220893, -2.51 DPS, sim-verified) [vendor]; Kilt of the Atal'ai Prophet (10807, -2.76 DPS) [dungeon]; Rainstrider Leggings (11123, -3.31 DPS) [quest] |
| feet | Vinerot Sandals (17748) | Maraudon: Razorlash [dungeon] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Earthenweave Boots (254093, -0.19 DPS) [crafted]; Sergeant Major's Satin Boots (220895, -0.81 DPS) [vendor]; Gilded Sandals (254107, -1.17 DPS, sim-verified) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 27.1 healing_power points (4.76 DPS) | yes | Cyclopean Band (11824, -1.11 DPS) [dungeon]; Snake Hoop (6750, -1.12 DPS) [quest]; Lorekeeper's Ring (19523, -1.29 DPS) [rep] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 22.7 healing_power points (3.98 DPS) | yes | Cyclopean Band (11824, -0.24 DPS, sim-verified) [dungeon]; Snake Hoop (6750, -0.33 DPS) [quest]; Lorekeeper's Ring (19523, -0.50 DPS) [rep] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Prayer Beads (19990, -0.16 DPS) [quest]; Evonice's Landin' Pilla (18951, -0.35 DPS) [quest]; Enriched Seal (271903, -4.75 DPS, sim-verified) [rep] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Enriched Seal (271903, -0.23 DPS, sim-verified) [rep]; Blessed Prayer Beads (19990, -1.23 DPS) [quest]; Evonice's Landin' Pilla (18951, -1.41 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Resurgence Rod (17743, -7.43 DPS, sim-verified) [quest]; Charstone Dirk (17710, -10.38 DPS) [dungeon]; Dreamstaff (249454, -11.84 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Cairnstone Sliver (9654) | The Morrow Stone [quest] | 10.6 healing_power points (1.85 DPS) | yes | Jaina's Firestarter (13064, -0.15 DPS) [world_drop]; Charged Lightning Rod (11860, -0.19 DPS) [quest]; Earthen Rod (9381, -0.27 DPS) [dungeon] |

**New at 50:** head: Cassandra's Grace; shoulder: Mantle of Lost Hope; back: Caretaker's Cape; chest: Embrace of the Wind Serpent; wrist: Nethergeld Cuffs; hands: Gilded Gloves; waist: Dawnspire Cord; legs: Senior Designer's Pantaloons; feet: Vinerot Sandals; finger1: Brainlash; finger2: Runemender's Seal; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Cairnstone Sliver

No-known-source sample (15 of 606, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (human, 005005031330011430-33505003030000000-00000000000000000)

Set DPS (verified): 287.6. Weights run: 27.2s. Verify run: 654.3s. 1523 eligible items had no known source.

Race sweep: human 287.4 ± 0.2, gnome 284.9 ± 0.2, night-elf 273.6 ± 0.2, dwarf 271.1 ± 0.2 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.246, intellect=1.801 ± 0.076, spirit=2.676 ± 0.069, mp5=5.066 ± 0.073, crit=0.946 ± 0.043 per rating point (14 rating = 1%, 13.240 per %), spell_haste=not significant (5.069 ± 1.445)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | 103.4 healing_power points (12.12 DPS) | yes | Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Crown of Caer Darrow (13986, -5.74 DPS, sim-verified) [quest] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 77.2 healing_power points (9.05 DPS) | yes | Dragon's Eye (10829, -2.78 DPS, sim-verified) [dungeon]; Animated Chain Necklace (18723, -2.99 DPS) [dungeon]; Tooth of Gnarr (13141, -3.19 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp]; Mantle of the Scarlet Crusade (22405, -4.95 DPS, sim-verified) [dungeon] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Featherskin Cape (10843, -0.25 DPS) [world]; Shroud of the Exile (15421, -0.30 DPS) [quest]; Gracious Cape (18743, -1.86 DPS, sim-verified) [dungeon] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 131.3 healing_power points (15.40 DPS) | yes | Field Marshal's Satin Tunic (231624, +0.00 DPS) [vendor]; Robes of the Exalted (13346, -1.65 DPS, sim-verified) [dungeon]; Truefaith Vestments (14154, -1.82 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Satin Bracers (17606, +0.00 DPS) [pvp]; Dryad's Wrist Bindings (19595, +0.00 DPS) [rep]; Wyrmthalak's Shackles (13958, -3.44 DPS, sim-verified) [quest] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Theramore Handwraps (271945, -0.69 DPS) [rep]; Magister's Gloves (16684, -1.03 DPS) [dungeon]; Hands of the Exalted Herald (12554, -2.69 DPS, sim-verified) [dungeon] |
| waist | Penitent's Cinch (272394) | Pix Xizzix [vendor] | 87.2 healing_power points (10.23 DPS) | yes | Wisdom of the Timbermaw (19047, -0.38 DPS, sim-verified) [crafted]; Marshal's Satin Sash (17609, -0.66 DPS) [pvp]; Devout Belt (16696, -1.89 DPS) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 112.6 healing_power points (13.20 DPS) | yes | Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Knight-Captain's Satin Legguards (227125, -0.95 DPS) [pvp]; Senior Designer's Pantaloons (11841, -5.73 DPS, sim-verified) [dungeon] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Faith Healer's Boots (22247, +0.00 DPS) [dungeon]; Marshal's Satin Walkers (231627, +0.00 DPS) [vendor]; Mooncloth Boots (15802, -2.86 DPS, sim-verified) [crafted] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -0.85 DPS) [vendor]; The Postmaster's Seal (13392, -1.07 DPS) [dungeon]; Naglering (11669, -9.37 DPS, sim-verified) [dungeon] |
| finger2 | Rosewine Circle (13178) | Blackrock Spire: Urok Doomhowl [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -0.23 DPS) [vendor]; The Postmaster's Seal (13392, -0.45 DPS) [dungeon]; Fordring's Seal (16058, -0.77 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -0.20 DPS) [quest]; Evonice's Landin' Pilla (18951, -0.63 DPS) [quest]; Serenity Field (272439, -2.75 DPS, sim-verified) [vendor] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -1.11 DPS, sim-verified) [vendor]; Shard of the Splithooves (10659, -1.19 DPS) [quest]; Evonice's Landin' Pilla (18951, -1.61 DPS) [quest] |
| main_hand | Redemption (22406) | Stratholme: Cannon Master Willey [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Grand Marshal's Warhammer (23454, +0.00 DPS) [vendor]; Staff of Metanoia (22394, -0.35 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Eternal Light (22254) | Blackrock Depths: Watchman Doomgrip [dungeon] | 29.9 healing_power points (3.51 DPS) | yes | Oblivion's Touch (18761, -0.98 DPS) [dungeon]; Brilliant Wand (249385, -0.98 DPS) [crafted]; Cairnstone Sliver (9654, -1.22 DPS, sim-verified) [quest] |

**New at 60:** head: Devout Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; back: Drape of Recovery; wrist: Devout Bracers; hands: Devout Gloves; waist: Penitent's Cinch; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of Mending; finger2: Rosewine Circle; trinket2: Burst of Knowledge; main_hand: Redemption; ranged: Wand of Eternal Light

No-known-source sample (15 of 1523, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (human, 005005031330011430-33505003030000000-00000000000000000)

Set DPS (verified): 604.2. Weights run: 18.8s. Verify run: 502.8s. 1523 eligible items had no known source.

Race sweep: human 603.9 ± 0.4, gnome 601.1 ± 0.3, night-elf 588.0 ± 0.3, dwarf 585.1 ± 0.3 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.351, intellect=1.387 ± 0.066, spirit=1.155 ± 0.067, mp5=3.168 ± 0.076, crit=1.061 ± 0.059 per rating point (14 rating = 1%, 14.852 per %), spell_haste=not significant (3.831 ± 1.488)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -13.22 DPS, sim-verified) [crafted] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 54.5 healing_power points (13.42 DPS) | yes | Animated Chain Necklace (18723, -3.32 DPS, sim-verified) [dungeon]; Tooth of Gnarr (13141, -3.63 DPS) [dungeon]; Lady Maye's Pendant (14558, -4.23 DPS) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp]; Mantle of Lost Hope (22234, -12.94 DPS, sim-verified) [dungeon] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Caretaker's Cape (19530, -2.08 DPS) [rep]; Shroud of the Exile (15421, -2.39 DPS) [quest]; Gracious Cape (18743, -3.84 DPS, sim-verified) [dungeon] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 91.5 healing_power points (22.53 DPS) | yes | Field Marshal's Satin Tunic (231624, +0.00 DPS) [vendor]; Field Marshal's Satin Vestments (17605, -1.72 DPS) [vendor]; Embrace of the Wind Serpent (12462, -6.26 DPS, sim-verified) [world] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Funeral Cuffs (12626, +0.00 DPS) [dungeon]; Marshal's Satin Bracers (17606, +0.00 DPS) [pvp]; Dryad's Wrist Bindings (19595, -12.05 DPS, sim-verified) [rep] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Satin Handwraps (231623, +0.00 DPS) [vendor]; Theramore Handwraps (271945, +0.00 DPS) [rep]; Hands of the Exalted Herald (12554, -10.42 DPS, sim-verified) [dungeon] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 60.8 healing_power points (14.97 DPS) | yes | Marshal's Satin Sash (17609, -2.85 DPS) [pvp]; Devout Belt (16696, -2.88 DPS) [dungeon]; Penitent's Cinch (272394, -4.95 DPS, sim-verified) [vendor] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Satin Legguards (227125, +0.00 DPS) [pvp]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Sentinel's Silk Leggings (237815, -15.21 DPS, sim-verified) [vendor] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 58.4 healing_power points (14.39 DPS) | yes | Marshal's Satin Walkers (231627, +0.00 DPS) [vendor]; Faith Healer's Boots (22247, -1.09 DPS, sim-verified) [dungeon]; Marshal's Satin Sandals (17607, -2.58 DPS) [vendor] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, -1.14 DPS) [dungeon]; Blessed Band of Light (272407, -1.82 DPS) [vendor]; Naglering (11669, -13.14 DPS, sim-verified) [dungeon] |
| finger2 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -1.25 DPS) [vendor]; Rosewine Circle (13178, -2.61 DPS, sim-verified) [dungeon]; Maiden's Circle (13001, -2.63 DPS) [world_drop] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (+7.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Burst of Knowledge (11832, +0.00 DPS) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Shard of the Splithooves (10659, -0.49 DPS) [quest]; Weakness Analyzer (272438, -1.16 DPS, sim-verified) [vendor] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Argent Crusader (13249, -3.77 DPS, sim-verified) [quest]; Redemption (22406, -4.37 DPS) [dungeon]; Staff of Metanoia (22394, -4.78 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Eternal Light (22254) | Blackrock Depths: Watchman Doomgrip [dungeon] | 23.7 healing_power points (5.82 DPS) | yes | Brilliant Wand (249385, -1.22 DPS) [crafted]; Oblivion's Touch (18761, -1.28 DPS, sim-verified) [dungeon]; Bonecreeper Stylus (13938, -1.41 DPS) [dungeon] |

**New at 60:** head: Devout Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; back: Drape of Recovery; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of Mending; finger2: Fordring's Seal; trinket1: Dormant Heart of the Mountain; trinket2: Serenity Field; ranged: Wand of Eternal Light

No-known-source sample (15 of 1523, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

## Horde

### Band 20 (undead, 005005000010000000-00000000000000000-00000000000000000)

Set DPS (verified): 45.4. Weights run: 13.2s. Verify run: 68.6s. 196 eligible items had no known source.

Race sweep: undead 45.4 ± 0.0, troll 43.4 ± 0.0 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.027, intellect=0.700 ± 0.005, spirit=0.609 ± 0.006, mp5=2.069 ± 0.017, crit=0.056 ± 0.003 per rating point (14 rating = 1%, 0.781 per %), spell_haste=not significant (0.013 ± 0.043)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (2.21 DPS) | yes | Shadow Goggles (4373, -1.30 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.91 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 1.8 healing_power points (0.22 DPS) | yes | Scholarly Pendant (277203, -0.07 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.3 healing_power points (1.39 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.42 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.63 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.90 DPS) [crafted] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.44 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -0.75 DPS) [crafted]; Black Whelp Cloak (7283, -0.89 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 18.3 healing_power points (2.25 DPS) | yes | Corsair's Overshirt (5202, -0.30 DPS) [dungeon]; Robe of the Moccasin (6465, -0.40 DPS) [dungeon]; Bloody Apron (6226, -0.65 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 9.7 healing_power points (1.19 DPS) | yes | Bright Bracers (3647, -0.85 DPS) [world_drop]; Owlbeard Bracers (16981, -0.90 DPS) [quest]; Tabitha's Cuffs (251486, -1.27 DPS, sim-verified) [quest] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 13.1 healing_power points (1.61 DPS) | yes | Serpent Gloves (5970, -0.81 DPS, sim-verified) [dungeon]; Gnoll Casting Gloves (892, -0.87 DPS) [world]; Magefist Gloves (12977, -0.88 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 13.8 healing_power points (1.70 DPS) | yes | Novice Ardent's Sash (253887, -0.51 DPS, sim-verified) [crafted]; Novice Arcanist's Sash (253885, -0.95 DPS) [crafted]; Keller's Girdle (2911, -1.01 DPS) [world_drop] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 24.6 healing_power points (3.03 DPS) | yes | Darkweave Breeches (12987, -1.93 DPS, sim-verified) [world_drop]; Colorful Kilt (10048, -2.04 DPS) [crafted]; Silk-threaded Trousers (1929, -2.17 DPS) [dungeon] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 11.1 healing_power points (1.36 DPS) | yes | Spidersilk Boots (4320, -0.07 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.50 DPS) [crafted]; Woolen Boots (2583, -0.80 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 9.1 healing_power points (1.12 DPS) | yes | Loop of Sacrifice (281673, -0.69 DPS) [quest]; Volcanic Rock Ring (12053, -0.86 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -1.48 DPS, sim-verified) [dungeon] |
| finger2 | Band of Purification (12996) | World drop [world_drop] | sim-verified (45.4 DPS) | yes | Loop of Sacrifice (281673, -0.02 DPS) [quest]; Volcanic Rock Ring (12053, -0.19 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.21 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Living Root (6631) | Wailing Caverns: Verdan the Everliving [dungeon] | 48.0 healing_power points (5.90 DPS) | yes | Staff of the Blessed Seer (2271, -1.25 DPS, sim-verified) [dungeon]; Kris of Orgrimmar (15443, -1.85 DPS) [quest]; Evocator's Blade (2567, -1.92 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Sizzle Stick (8071) (or Greater Magic Wand (11288), Torchlight Wand (5240)) | Deviate Eradication [quest] | 2.0 healing_power points (0.25 DPS) | yes | Torchlight Wand (5240, +0.00 DPS) [quest]; Greater Magic Wand (11288, +0.00 DPS) [crafted]; Wand of Decay (5252, -0.07 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Advisor's Ring; finger2: Band of Purification; main_hand: Living Root; ranged: Sizzle Stick

No-known-source sample (15 of 196, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6360 Steelscale Crushfish; 6478 Rat Stompers; 8350 The 1 Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 20425 Advisor's Gnarled Staff

### Band 30 (undead, 005005031330001000-00000000000000000-00000000000000000)

Set DPS (verified): 96.6. Weights run: 15.1s. Verify run: 49.3s. 318 eligible items had no known source.

Race sweep: undead 96.5 ± 0.0, troll 89.0 ± 0.0 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.145, intellect=0.897 ± 0.014, spirit=1.163 ± 0.013, mp5=2.851 ± 0.015, crit=0.168 ± 0.012 per rating point (14 rating = 1%, 2.349 per %), spell_haste=0.961 ± 0.226

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 40.0 healing_power points (6.95 DPS) | yes | Filigreed Pristine Circlet (253975, -2.98 DPS, sim-verified) [crafted]; Pristine Circlet (253949, -3.82 DPS) [crafted]; Nightsky Cowl (4039, -4.07 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 8.2 healing_power points (1.43 DPS) | yes | Crystal Starfire Medallion (5003, -0.16 DPS, sim-verified) [world_drop]; Pendant of Myzrael (4614, -0.22 DPS) [dungeon]; Glowing Green Talisman (5002, -0.22 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 17.1 healing_power points (2.97 DPS) | yes | Death Speaker Mantle (6685, -0.18 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.70 DPS) [world_drop]; Nightsky Mantle (4718, -0.71 DPS) [world_drop] |
| back | Battle Healer's Cloak (19529) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Thresher Cape (6901, -0.25 DPS, sim-verified) [dungeon]; Darkspear Raider's Cloak (272078, -1.22 DPS) [vendor]; Windsong Drape (15468, -1.39 DPS) [quest] |
| chest | Pristine Gown (253961) | Tailoring [crafted] | 30.9 healing_power points (5.38 DPS) | yes | Death Speaker Robes (6682, -0.13 DPS, sim-verified) [dungeon]; Beguiler Robes (7728, -0.32 DPS) [dungeon]; Robes of Arugal (6324, -1.04 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 17.6 healing_power points (3.05 DPS) | yes | Mindthrust Bracers (1974, -0.65 DPS, sim-verified) [dungeon]; Nightsky Wristbands (6407, -1.51 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.81 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 18.4 healing_power points (3.21 DPS) | yes | Truefaith Gloves (7049, -0.56 DPS, sim-verified) [crafted]; Pristine Gloves (253913, -0.83 DPS) [crafted]; Gloves of Meditation (4318, -0.92 DPS) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 17.5 healing_power points (3.05 DPS) | yes | Pristine Sash (253925, -0.49 DPS, sim-verified) [crafted]; Defiler's Cloth Girdle (20164, -0.66 DPS) [rep]; Crimson Silk Belt (7055, -0.91 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 34.1 healing_power points (5.93 DPS) | yes | Filigreed Pristine Leggings (253937, -1.09 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -1.10 DPS) [crafted]; Gaze Dreamer Pants (6903, -2.02 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 30.9 healing_power points (5.38 DPS) | yes | Acidic Walkers (9454, -2.35 DPS, sim-verified) [dungeon]; Boots of the Enchanter (4325, -2.89 DPS) [crafted]; Glinteye Slippers (273024, -2.95 DPS) [dungeon] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 22.5 healing_power points (3.92 DPS) | yes | Electrocutioner Lagnut (9447, -0.94 DPS) [dungeon]; Solid Gold Band (279845, -0.99 DPS) [quest]; Ring of Power Regulation (273030, -1.39 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 20.4 healing_power points (3.55 DPS) | yes | Solid Gold Band (279845, -0.63 DPS) [quest]; Electrocutioner Lagnut (9447, -0.77 DPS, sim-verified) [dungeon]; Ring of Power Regulation (273030, -1.02 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Odo's Ley Staff (6318, -0.36 DPS) [dungeon]; Living Root (6631, -2.99 DPS) [dungeon]; Staff of the Friar (3415, -3.25 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 20.0 healing_power points (3.48 DPS) | yes | Orb of Mistmantle (13031, -1.05 DPS, sim-verified) [world_drop]; Orb of Mystic Insight (249394, -1.32 DPS) [crafted]; Satyr's Rod (15962, -2.13 DPS) [world_drop] |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 9.3 healing_power points (1.62 DPS) | yes | Wand of Eventide (5214, -0.75 DPS) [world_drop]; Greater Mystic Wand (217287, -0.75 DPS) [crafted]; Thunderwood (13062, -0.78 DPS, sim-verified) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Battle Healer's Cloak; chest: Pristine Gown; wrist: Spidertank Oilrag; hands: Gloves of Old; waist: Belt of Arugal; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Runemender's Seal; finger2: Snake Hoop; main_hand: Death Speaker Scepter; off_hand: Orb of Souls; ranged: Wand of Mana Concentration

No-known-source sample (15 of 318, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6360 Steelscale Crushfish; 6478 Rat Stompers; 8350 The 1 Ring

### Band 40 (undead, 005005031330011000-33300000000000000-00000000000000000)

Set DPS (verified): 149.0. Weights run: 25.3s. Verify run: 153.5s. 454 eligible items had no known source.

Race sweep: undead 149.0 ± 0.1, troll 136.7 ± 0.0 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.134, intellect=1.395 ± 0.026, spirit=1.807 ± 0.029, mp5=3.525 ± 0.027, crit=0.358 ± 0.019 per rating point (14 rating = 1%, 5.016 per %), spell_haste=3.998 ± 0.484

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 62.0 healing_power points (8.80 DPS) | yes | Whitemane's Chapeau (7720, -1.46 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -2.57 DPS) [world_drop]; Earthen Silk Hood (254015, -3.11 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 31.2 healing_power points (4.42 DPS) | yes | Triune Amulet (7722, -1.49 DPS) [dungeon]; Amberglow Talisman (10824, -1.86 DPS) [quest]; Necklace of Calisea (1714, -1.90 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 40.1 healing_power points (5.70 DPS) | yes | Mantle of Doan (7712, -1.46 DPS) [dungeon]; Earthen Silk Shoulders (254033, -2.19 DPS, sim-verified) [crafted]; Mistscape Mantle (4734, -2.24 DPS) [dungeon] |
| back | Battle Healer's Cloak (19528) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Raider's Cloak (272077, -0.63 DPS) [vendor]; Glowing Thresher Cape (6901, -0.79 DPS) [dungeon]; Cloak of Blight (6832, -1.76 DPS, sim-verified) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 72.5 healing_power points (10.28 DPS) | yes | Death's Head Vestment (10581, -2.81 DPS) [dungeon]; Doomsayer's Robe (4746, -3.40 DPS) [quest]; Dreamweave Vest (10021, -3.74 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 19.6 healing_power points (2.78 DPS) | yes | Mindthrust Bracers (1974, -0.29 DPS) [dungeon]; Radiant Silver Bracers (4545, -0.63 DPS) [quest]; Earthen Silk Cuffs (254019, -0.79 DPS, sim-verified) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 36.2 healing_power points (5.14 DPS) | yes | Gilded Handwraps (254021, -0.63 DPS) [crafted]; Earthen Silk Gloves (254017, -0.98 DPS, sim-verified) [crafted]; Gloves of Old (9395, -1.07 DPS) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 44.0 healing_power points (6.25 DPS) | yes | Windchaser Cinch (14435, -2.35 DPS) [world_drop]; Deathmage Sash (10771, -2.56 DPS, sim-verified) [dungeon]; Sutarn's Ring (13105, -2.73 DPS) [world_drop] |
| legs | Stoneweaver Leggings (9407) | Uldaman: Ironaya [dungeon] | 52.8 healing_power points (7.50 DPS) | yes | Pristine Leggings (253987, -1.71 DPS) [crafted]; Stormcloth Pants (10010, -2.14 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -2.55 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Thoughtcast Boots (10578, -0.07 DPS) [dungeon]; Earthen Silk Slippers (254013, -1.04 DPS) [crafted]; Furen's Boots (13100, -1.16 DPS, sim-verified) [world_drop] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 30.0 healing_power points (4.26 DPS) | yes | Electrocutioner Lagnut (9447, -1.19 DPS) [dungeon]; Darkspear Signet (272070, -1.26 DPS) [vendor]; Solid Gold Band (279845, -1.42 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 28.4 healing_power points (4.03 DPS) | yes | Electrocutioner Lagnut (9447, -0.98 DPS, sim-verified) [dungeon]; Darkspear Signet (272070, -1.03 DPS) [vendor]; Solid Gold Band (279845, -1.19 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+4.2 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of the Friar (3415, -4.34 DPS) [dungeon]; Odo's Ley Staff (6318, -4.94 DPS) [dungeon]; Staff of Jordan (873, -5.88 DPS, sim-verified) [world_drop] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | 29.2 healing_power points (4.15 DPS) | yes | Prophetic Cane (6803, -0.01 DPS) [quest]; Orb of Souls (249395, -0.76 DPS) [crafted]; Omega Orb (7749, -1.37 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 13.8 healing_power points (1.96 DPS) | yes | Captain Rackmore's Tiller (16789, -0.34 DPS) [quest]; Goblin Igniter (5253, -0.40 DPS) [quest]; Wand of Mana Concentration (273043, -0.50 DPS, sim-verified) [world] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Battle Healer's Cloak; chest: Stormcloth Vest; hands: Dreamweave Gloves; waist: Gilded Cord; legs: Stoneweaver Leggings; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Hand of Righteousness; off_hand: Beacon of Hope; ranged: Jaina's Firestarter

No-known-source sample (15 of 454, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger; 5743 Prismstone Ring

### Band 50 (undead, 005005031330011430-33501000000000000-00000000000000000)

Set DPS (verified): 210.6. Weights run: 26.6s. Verify run: 219.4s. 599 eligible items had no known source.

Race sweep: undead 210.6 ± 0.1, troll 195.0 ± 0.1 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.116, intellect=1.117 ± 0.027, spirit=0.995 ± 0.030, mp5=1.951 ± 0.032, crit=0.413 ± 0.020 per rating point (14 rating = 1%, 5.777 per %), spell_haste=not significant (1.835 ± 0.506)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Cassandra's Grace (13102) | World drop [world_drop] | 53.0 healing_power points (9.29 DPS) | yes | Blood Guard's Satin Cover (220899, +0.00 DPS) [vendor]; Papal Fez (9431, -0.53 DPS) [dungeon]; Gemburst Circlet (10751, -0.54 DPS) [quest] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 28.9 healing_power points (5.08 DPS) | yes | Gemshard Heart (17707, -0.46 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -1.63 DPS) [world_drop]; Mindburst Medallion (11196, -1.98 DPS) [quest] |
| shoulder | Mantle of Lost Hope (22234) | Blackrock Depths: Lord Roccor [dungeon] | 46.1 healing_power points (8.09 DPS) | yes | Blood Guard's Satin Pads (220901, -0.90 DPS) [vendor]; Nethergeld Shoulders (254049, -1.06 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -1.68 DPS) [dungeon] |
| back | Battle Healer's Cloak (19527) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -0.23 DPS) [dungeon]; Featherskin Cape (10843, -0.71 DPS) [world]; Deep Woodlands Cloak (19121, -1.04 DPS) [quest] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 74.0 healing_power points (12.99 DPS) | yes | Robes of Insight (940, -2.49 DPS) [world_drop]; Grimlok's Tribal Vestments (9415, -4.48 DPS) [dungeon]; Stormcloth Vest (10020, -4.57 DPS, sim-verified) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 32.8 healing_power points (5.75 DPS) | yes | Aristocratic Cuffs (12546, -1.27 DPS, sim-verified) [dungeon]; Bloodband Bracers (11469, -2.59 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -2.88 DPS) [quest] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 48.1 healing_power points (8.45 DPS) | yes | Earthenweave Gloves (254075, -1.27 DPS, sim-verified) [crafted]; First Sergeant's Satin Gloves (220898, -2.22 DPS) [vendor]; Greenleaf Handwraps (19116, -2.43 DPS) [quest] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 39.2 healing_power points (6.88 DPS) | yes | Gilded Waistcord (254081, -0.03 DPS) [crafted]; Earthenweave Cord (254077, -0.40 DPS) [crafted]; Gilded Cord (254037, -0.41 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 62.7 healing_power points (11.01 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -2.76 DPS) [dungeon]; Stone Guard's Satin Leggings (220902, -3.12 DPS, sim-verified) [vendor]; Rainstrider Leggings (11123, -3.31 DPS) [quest] |
| feet | Vinerot Sandals (17748) | Maraudon: Razorlash [dungeon] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Earthenweave Boots (254093, -0.19 DPS) [crafted]; First Sergeant's Satin Boots (220900, -0.81 DPS) [vendor]; Gilded Sandals (254107, -1.15 DPS, sim-verified) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 27.1 healing_power points (4.76 DPS) | yes | Cyclopean Band (11824, -1.11 DPS) [dungeon]; Snake Hoop (6750, -1.12 DPS) [quest]; Advisor's Ring (19519, -1.29 DPS) [rep] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 22.7 healing_power points (3.98 DPS) | yes | Snake Hoop (6750, -0.33 DPS) [quest]; Advisor's Ring (19519, -0.50 DPS) [rep]; Cyclopean Band (11824, -0.67 DPS, sim-verified) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+5.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Blessed Prayer Beads (19990, -0.16 DPS) [quest]; Evonice's Landin' Pilla (18951, -0.35 DPS) [quest] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Blessed Prayer Beads (19990, +0.00 DPS) [quest]; Uther's Strength (11302, -0.37 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Resurgence Rod (17743, -7.76 DPS, sim-verified) [quest]; Charstone Dirk (17710, -10.38 DPS) [dungeon]; Dreamstaff (249454, -11.84 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 9.7 healing_power points (1.70 DPS) | yes | Earthen Rod (9381, -0.12 DPS) [dungeon]; Captain Rackmore's Tiller (16789, -0.12 DPS) [quest]; Charged Lightning Rod (11860, -0.23 DPS, sim-verified) [quest] |

**New at 50:** head: Cassandra's Grace; shoulder: Mantle of Lost Hope; back: Battle Healer's Cloak; chest: Embrace of the Wind Serpent; wrist: Nethergeld Cuffs; hands: Gilded Gloves; waist: Dawnspire Cord; legs: Senior Designer's Pantaloons; feet: Vinerot Sandals; finger1: Brainlash; finger2: Runemender's Seal; trinket2: Darkspear Voodoo Seal; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 599, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60 (undead, 005005031330011430-33505003030000000-00000000000000000)

Set DPS (verified): 286.9. Weights run: 27.2s. Verify run: 681.3s. 1518 eligible items had no known source.

Race sweep: undead 287.2 ± 0.2, troll 271.5 ± 0.2 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.246, intellect=1.801 ± 0.076, spirit=2.676 ± 0.069, mp5=5.066 ± 0.073, crit=0.946 ± 0.043 per rating point (14 rating = 1%, 13.240 per %), spell_haste=not significant (5.069 ± 1.445)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | 103.4 healing_power points (12.12 DPS) | yes | Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Crown of Caer Darrow (13986, -9.46 DPS, sim-verified) [quest] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 77.2 healing_power points (9.05 DPS) | yes | Animated Chain Necklace (18723, -2.99 DPS) [dungeon]; Tooth of Gnarr (13141, -3.19 DPS) [dungeon]; Dragon's Eye (10829, -3.21 DPS, sim-verified) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp]; Mantle of the Scarlet Crusade (22405, -6.79 DPS, sim-verified) [dungeon] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Featherskin Cape (10843, -0.25 DPS) [world]; Shroud of the Exile (15421, -0.30 DPS) [quest]; Gracious Cape (18743, -2.09 DPS, sim-verified) [dungeon] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 131.3 healing_power points (15.40 DPS) | yes | Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Truefaith Vestments (14154, -1.82 DPS) [crafted]; Robes of the Exalted (13346, -3.06 DPS, sim-verified) [dungeon] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Satin Bracers (17619, +0.00 DPS) [pvp]; Dryad's Wrist Bindings (19595, +0.00 DPS) [rep]; Wyrmthalak's Shackles (13958, -6.99 DPS, sim-verified) [quest] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Raider Handwraps (272097, -0.69 DPS) [vendor]; Magister's Gloves (16684, -1.03 DPS) [dungeon]; Hands of the Exalted Herald (12554, -5.80 DPS, sim-verified) [dungeon] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Satin Cinch (17621, +0.00 DPS) [pvp]; Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Penitent's Cinch (272394, -3.23 DPS, sim-verified) [vendor] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 112.6 healing_power points (13.20 DPS) | yes | General's Satin Legguards (231634, +0.00 DPS) [vendor]; Legionnaire's Satin Legguards (227123, -0.95 DPS) [pvp]; Senior Designer's Pantaloons (11841, -8.16 DPS, sim-verified) [dungeon] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 84.0 healing_power points (9.85 DPS) | yes | Faith Healer's Boots (22247, -0.74 DPS) [dungeon]; General's Satin Walkers (231630, -0.89 DPS) [vendor]; Devout Sandals (16691, -1.11 DPS) [dungeon] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, -0.62 DPS) [dungeon]; The Postmaster's Seal (13392, -1.07 DPS) [dungeon]; Naglering (11669, -10.53 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, +0.00 DPS) [dungeon]; The Postmaster's Seal (13392, -0.22 DPS) [dungeon]; Fordring's Seal (16058, -0.54 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.63 DPS) [quest]; Darkspear Voodoo Seal (272061, -0.80 DPS) [vendor]; Shard of the Splithooves (10659, -1.01 DPS, sim-verified) [quest] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -1.61 DPS) [quest]; Darkspear Voodoo Seal (272061, -1.78 DPS) [vendor]; Shard of the Splithooves (10659, -3.03 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Mace (23464, +0.00 DPS) [vendor]; Staff of Metanoia (22394, -0.64 DPS) [dungeon]; Redemption (22406, -3.10 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Eternal Light (22254) | Blackrock Depths: Watchman Doomgrip [dungeon] | 29.9 healing_power points (3.51 DPS) | yes | Brilliant Wand (249385, -0.98 DPS) [crafted]; Bonecreeper Stylus (13938, -1.16 DPS) [dungeon]; Oblivion's Touch (18761, -2.81 DPS, sim-verified) [dungeon] |

**New at 60:** head: Devout Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; back: Drape of Recovery; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of Mending; finger2: Blessed Band of Light; trinket2: Burst of Knowledge; ranged: Wand of Eternal Light

No-known-source sample (15 of 1518, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60, raid preset (undead, 005005031330011430-33505003030000000-00000000000000000)

Set DPS (verified): 596.6. Weights run: 18.8s. Verify run: 501.7s. 1518 eligible items had no known source.

Race sweep: undead 596.5 ± 0.4, troll 584.3 ± 0.3 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.351, intellect=1.387 ± 0.066, spirit=1.155 ± 0.067, mp5=3.168 ± 0.076, crit=1.061 ± 0.059 per rating point (14 rating = 1%, 14.852 per %), spell_haste=not significant (3.831 ± 1.488)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | 70.6 healing_power points (17.39 DPS) | yes | Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -12.00 DPS, sim-verified) [crafted] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 54.5 healing_power points (13.42 DPS) | yes | Animated Chain Necklace (18723, -2.66 DPS, sim-verified) [dungeon]; Tooth of Gnarr (13141, -3.63 DPS) [dungeon]; Lady Maye's Pendant (14558, -4.23 DPS) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp]; Mantle of Lost Hope (22234, -11.42 DPS, sim-verified) [dungeon] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battle Healer's Cloak (19526, -2.08 DPS) [rep]; Shroud of the Exile (15421, -2.39 DPS) [quest]; Gracious Cape (18743, -4.17 DPS, sim-verified) [dungeon] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 91.5 healing_power points (22.53 DPS) | yes | Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Warlord's Satin Robes (231612, -1.72 DPS) [pvp]; Embrace of the Wind Serpent (12462, -6.22 DPS, sim-verified) [world] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Funeral Cuffs (12626, +0.00 DPS) [dungeon]; General's Satin Bracers (17619, +0.00 DPS) [pvp]; Dryad's Wrist Bindings (19595, -10.45 DPS, sim-verified) [rep] |
| hands | Hands of the Exalted Herald (12554) | Blackrock Depths: Princess Moira Bronzebeard  [dungeon] | 62.9 healing_power points (15.49 DPS) | yes | General's Satin Handwraps (231633, -0.21 DPS) [vendor]; Raider Handwraps (272097, -1.90 DPS, sim-verified) [vendor] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Satin Cinch (17621, +0.00 DPS) [pvp]; Penitent's Cinch (272394, +0.00 DPS) [vendor]; Wisdom of the Timbermaw (19047, -9.93 DPS, sim-verified) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 71.4 healing_power points (17.57 DPS) | yes | Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Senior Designer's Pantaloons (11841, -14.77 DPS, sim-verified) [dungeon] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 58.4 healing_power points (14.39 DPS) | yes | General's Satin Walkers (231630, +0.00 DPS) [vendor]; Faith Healer's Boots (22247, -0.97 DPS) [dungeon]; General's Satin Boots (17618, -2.58 DPS) [vendor] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, -1.14 DPS) [dungeon]; Blessed Band of Light (272407, -1.82 DPS) [vendor]; Naglering (11669, -14.67 DPS, sim-verified) [dungeon] |
| finger2 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -1.25 DPS) [vendor]; Maiden's Circle (13001, -2.63 DPS) [world_drop]; Rosewine Circle (13178, -3.18 DPS, sim-verified) [dungeon] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (+6.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Burst of Knowledge (11832, +0.00 DPS) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Shard of the Splithooves (10659, -0.49 DPS) [quest]; Weakness Analyzer (272438, -0.90 DPS, sim-verified) [vendor] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Argent Crusader (13249, -3.69 DPS, sim-verified) [quest]; Redemption (22406, -4.37 DPS) [dungeon]; Staff of Metanoia (22394, -4.78 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Eternal Light (22254) | Blackrock Depths: Watchman Doomgrip [dungeon] | 23.7 healing_power points (5.82 DPS) | yes | Brilliant Wand (249385, -1.22 DPS) [crafted]; Bonecreeper Stylus (13938, -1.41 DPS) [dungeon]; Oblivion's Touch (18761, -1.54 DPS, sim-verified) [dungeon] |

**New at 60:** head: Devout Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; back: Drape of Recovery; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Hands of the Exalted Herald; waist: Devout Belt; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of Mending; finger2: Fordring's Seal; trinket1: Dormant Heart of the Mountain; trinket2: Serenity Field; ranged: Wand of Eternal Light

No-known-source sample (15 of 1518, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

