# Leveling BiS: Holy

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 000000000000000000-03502001000000000-00000000000000000)

Set DPS (verified): 42.0. Weights run: 14.7s. Verify run: 68.5s. 203 eligible items had no known source.

Race sweep: human 42.0 ± 0.0, gnome 40.3 ± 0.0, night-elf 39.4 ± 0.0, dwarf 39.0 ± 0.0 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.096, intellect=0.985 ± 0.015, spirit=0.901 ± 0.013, mp5=2.282 ± 0.019, crit=0.102 ± 0.007 per rating point (14 rating = 1%, 1.425 per %), spell_haste=not significant (-0.078 ± 0.114)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (2.24 DPS) | yes | Shadow Goggles (4373, -1.42 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.79 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 1.8 healing_power points (0.22 DPS) | yes | Tarnished Locket (279870, -0.44 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.9 healing_power points (1.72 DPS) | yes | Slime-encrusted Pads (6461, -0.87 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -1.02 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.23 DPS) [crafted] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | sim-verified (42.0 DPS) | yes | Heavy Woolen Cloak (4311, -0.33 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -0.73 DPS) [crafted]; Sanguine Cape (14376, -0.85 DPS) [world_drop] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 20.6 healing_power points (2.57 DPS) | yes | Robe of the Moccasin (6465, -0.33 DPS) [dungeon]; Corsair's Overshirt (5202, -0.73 DPS, sim-verified) [dungeon]; Scholarly Robes (2034, -0.85 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 11.8 healing_power points (1.46 DPS) | yes | Repurposed Hair Band (281256, -0.99 DPS) [quest]; Seer's Cuffs (3645, -1.12 DPS) [dungeon]; Bright Bracers (3647, -2.02 DPS, sim-verified) [world_drop] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 14.0 healing_power points (1.74 DPS) | yes | Bright Gloves (3066, -0.80 DPS) [world_drop]; Serpent Gloves (5970, -0.86 DPS) [dungeon]; Magefist Gloves (12977, -0.92 DPS, sim-verified) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 14.9 healing_power points (1.86 DPS) | yes | Novice Ardent's Sash (253887, -0.42 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.88 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.99 DPS) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 27.5 healing_power points (3.42 DPS) | yes | Darkweave Breeches (12987, -1.92 DPS, sim-verified) [world_drop]; Filigreed Silky Leggings (253939, -2.24 DPS) [crafted]; Filigreed Flame Leggings (253941, -2.24 DPS) [crafted] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 12.0 healing_power points (1.49 DPS) | yes | Spidersilk Boots (4320, -0.19 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.43 DPS) [crafted]; Kimbra Boots (6191, -0.66 DPS) [quest] |
| finger1 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 9.6 healing_power points (1.19 DPS) | yes | Lavishly Jeweled Ring (1156, -0.45 DPS) [dungeon]; Band of Purification (12996, -0.52 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.82 DPS) [world_drop] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 7.0 healing_power points (0.87 DPS) | yes | Band of Purification (12996, -0.19 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.25 DPS, sim-verified) [dungeon]; Volcanic Rock Ring (12053, -0.50 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Westfall (2042) | The Defias Brotherhood [quest] | 58.3 healing_power points (7.25 DPS) | yes | Living Root (6631, -1.77 DPS, sim-verified) [dungeon]; Staff of the Blessed Seer (2271, -2.65 DPS) [dungeon]; Evocator's Blade (2567, -2.97 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Sizzle Stick (8071) (or Greater Magic Wand (11288), Torchlight Wand (5240)) | Deviate Eradication [quest] | 2.0 healing_power points (0.25 DPS) | yes | Torchlight Wand (5240, +0.00 DPS) [quest]; Greater Magic Wand (11288, +0.00 DPS) [crafted]; Flaring Baton (5326, -0.00 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Staff of Westfall; ranged: Sizzle Stick

No-known-source sample (15 of 203, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6360 Steelscale Crushfish; 6478 Rat Stompers; 8350 The 1 Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (human, 000000000000000000-33502003030110000-00000000000000000)

Set DPS (verified): 74.3. Weights run: 14.9s. Verify run: 88.6s. 325 eligible items had no known source.

Race sweep: gnome 74.3 ± 0.0, human 74.3 ± 0.0, night-elf 72.7 ± 0.0, dwarf 72.3 ± 0.0 (heal_score); gnome beats the guide race human by 0.0 (0.07%) ± 0.0, beyond error.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.143, intellect=1.105 ± 0.017, spirit=0.506 ± 0.016, mp5=2.482 ± 0.032, crit=0.162 ± 0.012 per rating point (14 rating = 1%, 2.271 per %), spell_haste=not significant (-0.082 ± 0.158)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 36.0 healing_power points (4.96 DPS) | yes | Filigreed Pristine Circlet (253975, -1.37 DPS, sim-verified) [crafted]; Pristine Circlet (253949, -2.48 DPS) [crafted]; Enchanter's Cowl (4322, -2.61 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 6.6 healing_power points (0.91 DPS) | yes | Crystal Starfire Medallion (5003, -0.10 DPS) [world_drop]; Seal of the Expedition (271920, -0.15 DPS) [rep]; Kaleidoscope Chain (13084, -0.22 DPS, sim-verified) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 18.9 healing_power points (2.61 DPS) | yes | Death Speaker Mantle (6685, -0.11 DPS) [dungeon]; Magician's Mantle (12998, -0.55 DPS) [world_drop]; Feline Mantle (3748, -0.84 DPS) [dungeon] |
| back | Caretaker's Cape (19533) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Prelacy Cape (7004, -0.29 DPS, sim-verified) [quest]; Repairman's Cape (9605, -0.44 DPS) [quest]; Glowing Thresher Cape (6901, -0.55 DPS) [dungeon] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 32.1 healing_power points (4.42 DPS) | yes | Pristine Gown (253961, -0.61 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -0.80 DPS) [dungeon]; Green Silk Armor (7065, -1.21 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 16.4 healing_power points (2.26 DPS) | yes | Mindthrust Bracers (1974, -0.54 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -1.05 DPS) [world_drop]; Nightsky Wristbands (6407, -1.14 DPS) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 18.3 healing_power points (2.52 DPS) | yes | Pristine Gloves (253913, -0.56 DPS, sim-verified) [crafted]; Gloves of Old (9395, -0.92 DPS) [world_drop]; Shilly Mitts (9609, -0.93 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 15.4 healing_power points (2.12 DPS) | yes | Highlander's Cloth Girdle (20099, -0.15 DPS) [rep]; Crimson Silk Belt (7055, -0.23 DPS) [crafted]; Belt of Arugal (6392, -0.37 DPS, sim-verified) [dungeon] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 32.3 healing_power points (4.44 DPS) | yes | Filigreed Pristine Leggings (253937, -0.54 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -1.61 DPS) [crafted]; Leech Pants (6910, -2.04 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 29.8 healing_power points (4.09 DPS) | yes | Acidic Walkers (9454, -1.96 DPS, sim-verified) [dungeon]; Pristine Boots (253889, -2.40 DPS) [crafted]; Spidersilk Boots (4320, -2.52 DPS) [crafted] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 19.2 healing_power points (2.64 DPS) | yes | Ring of Power Regulation (273030, -0.79 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.91 DPS) [dungeon]; Lorekeeper's Ring (19525, -0.99 DPS) [rep] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 17.3 healing_power points (2.38 DPS) | yes | Electrocutioner Lagnut (9447, -0.65 DPS) [dungeon]; Lorekeeper's Ring (19525, -0.73 DPS) [rep]; Ring of Power Regulation (273030, -0.76 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Odo's Ley Staff (6318, -1.37 DPS) [dungeon]; Staff of Westfall (2042, -1.57 DPS) [quest]; Staff of the Friar (3415, -3.56 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | sim-verified (74.3 DPS) | yes | Eye of Paleth (2943, -0.09 DPS) [quest]; Orb of Souls (249395, -0.10 DPS, sim-verified) [crafted]; Orb of Mistmantle (13031, -0.36 DPS) [world_drop] |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 9.4 healing_power points (1.29 DPS) | yes | Thunderwood (13062, -0.07 DPS, sim-verified) [world_drop]; Wand of Eventide (5214, -0.60 DPS) [world_drop]; Greater Mystic Wand (217287, -0.60 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Caretaker's Cape; chest: Death Speaker Robes; wrist: Spidertank Oilrag; hands: Truefaith Gloves; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Runemender's Seal; finger2: Snake Hoop; main_hand: Death Speaker Scepter; off_hand: Orb of Mystic Insight; ranged: Wand of Mana Concentration

No-known-source sample (15 of 325, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6360 Steelscale Crushfish

### Band 40 (human, 000000000000000000-33502003030121431-00000000000000000)

Set DPS (verified): 175.9. Weights run: 26.2s. Verify run: 175.5s. 461 eligible items had no known source.

Race sweep: human 175.9 ± 0.2, gnome 171.1 ± 0.2, night-elf 168.7 ± 0.2, dwarf 168.1 ± 0.2 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.333, intellect=0.628 ± 0.041, spirit=0.993 ± 0.040, mp5=1.492 ± 0.044, crit=0.124 ± 0.031 per rating point (14 rating = 1%, 1.736 per %), spell_haste=not significant (0.034 ± 0.351)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 41.6 healing_power points (12.25 DPS) | yes | Whitemane's Chapeau (7720, -1.01 DPS) [dungeon]; Holy Shroud (2721, -2.05 DPS, sim-verified) [world_drop]; Earthen Silk Hood (254015, -2.85 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 25.0 healing_power points (7.37 DPS) | yes | Triune Amulet (7722, -2.73 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -4.03 DPS) [dungeon]; Amberglow Talisman (10824, -4.44 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 30.2 healing_power points (8.88 DPS) | yes | Earthen Silk Shoulders (254033, -1.95 DPS, sim-verified) [crafted]; Mantle of Doan (7712, -3.37 DPS) [dungeon]; Green Silken Shoulders (7057, -4.49 DPS) [crafted] |
| back | Caretaker's Cape (19532) | Silverwing Sentinels [rep] | sim-verified (175.9 DPS) | yes | Prelacy Cape (7004, -1.77 DPS) [quest]; Glowing Thresher Cape (6901, -2.36 DPS) [dungeon]; Icy Cloak (4327, -2.95 DPS) [crafted] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 65.9 healing_power points (19.42 DPS) | yes | Civinad Robes (9623, -5.49 DPS, sim-verified) [quest]; Dreamweave Vest (10021, -8.36 DPS) [crafted]; Death's Head Vestment (10581, -8.84 DPS) [dungeon] |
| wrist | Earthen Silk Cuffs (254019) | Tailoring [crafted] | 15.0 healing_power points (4.41 DPS) | yes | Spidertank Oilrag (9448, -0.44 DPS) [dungeon]; Arcane Runed Bracers (4744, -1.76 DPS) [quest]; Mindthrust Bracers (1974, -2.17 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 27.5 healing_power points (8.09 DPS) | yes | Gilded Handwraps (254021, -0.30 DPS, sim-verified) [crafted]; Earthen Silk Gloves (254017, -0.45 DPS) [crafted]; Restorer's Fine Gloves (270059, -2.51 DPS) [quest] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 33.0 healing_power points (9.71 DPS) | yes | Star Belt (4329, -4.63 DPS, sim-verified) [crafted]; Highlander's Cloth Girdle (20098, -4.85 DPS) [rep]; Deathmage Sash (10771, -4.88 DPS) [dungeon] |
| legs | Stoneweaver Leggings (9407) | Uldaman: Ironaya [dungeon] | 36.9 healing_power points (10.88 DPS) | yes | Stormcloth Pants (10010, -2.07 DPS) [crafted]; Pristine Leggings (253987, -2.12 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -3.24 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 28.4 healing_power points (8.35 DPS) | yes | Earthen Silk Slippers (254013, -1.01 DPS) [crafted]; Thoughtcast Boots (10578, -1.02 DPS, sim-verified) [dungeon]; Furen's Boots (13100, -2.38 DPS) [world_drop] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 19.7 healing_power points (5.81 DPS) | yes | Electrocutioner Lagnut (9447, -1.11 DPS) [dungeon]; Lorekeeper's Ring (19524, -1.84 DPS) [rep]; Philanthropist's Ring (281635, -1.94 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 17.3 healing_power points (5.11 DPS) | yes | Lorekeeper's Ring (19524, -1.14 DPS) [rep]; Electrocutioner Lagnut (9447, -1.17 DPS, sim-verified) [dungeon]; Philanthropist's Ring (281635, -1.24 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (175.9 DPS) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (175.9 DPS) | yes | Staff of Jordan (873, -8.73 DPS, sim-verified) [world_drop]; Death Speaker Scepter (2816, -10.00 DPS) [dungeon]; Staff of the Friar (3415, -10.64 DPS) [dungeon] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | 26.0 healing_power points (7.65 DPS) | yes | Orb of Lorica (11262, -0.46 DPS, sim-verified) [quest]; Orb of Souls (249395, -2.07 DPS) [crafted]; Orb of the Forgotten Seer (7685, -3.53 DPS) [dungeon] |
| ranged | Earthen Rod (9381) | Uldaman [dungeon] | 9.0 healing_power points (2.65 DPS) | yes | Captain Rackmore's Tiller (16789, -0.01 DPS) [quest]; Jaina's Firestarter (13064, -0.66 DPS) [world_drop]; Twisted Nether Wand (249144, -0.88 DPS) [crafted] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Caretaker's Cape; chest: Stormcloth Vest; wrist: Earthen Silk Cuffs; hands: Dreamweave Gloves; waist: Gilded Cord; legs: Stoneweaver Leggings; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Hand of Righteousness; off_hand: Beacon of Hope; ranged: Earthen Rod

No-known-source sample (15 of 461, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (human, 005005000000000000-33502003030121431-00000000000000000)

Set DPS (verified): 258.4. Weights run: 28.0s. Verify run: 246.4s. 606 eligible items had no known source.

Race sweep: human 258.0 ± 0.2, gnome 250.4 ± 0.2, night-elf 246.3 ± 0.2, dwarf 244.9 ± 0.2 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.285, intellect=0.783 ± 0.045, spirit=0.222 ± 0.043, mp5=0.845 ± 0.045, crit=0.184 ± 0.033 per rating point (14 rating = 1%, 2.574 per %), spell_haste=not significant (-0.085 ± 0.480)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Cassandra's Grace (13102) | World drop [world_drop] | 46.0 healing_power points (16.02 DPS) | yes | Knight-Lieutenant's Satin Cover (220896, -0.90 DPS) [vendor]; Soulcatcher Halo (10630, -1.95 DPS, sim-verified) [dungeon]; Papal Fez (9431, -3.03 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 26.3 healing_power points (9.15 DPS) | yes | Gemshard Heart (17707, -0.97 DPS, sim-verified) [dungeon]; Talvash's Enhancing Necklace (7673, -4.70 DPS) [quest]; Horizon Choker (13085, -5.02 DPS) [world_drop] |
| shoulder | Mantle of Lost Hope (22234) | Blackrock Depths: Lord Roccor [dungeon] | 38.0 healing_power points (13.23 DPS) | yes | Knight-Lieutenant's Satin Pads (220894, -1.67 DPS) [vendor]; Inquisitor's Shawl (19507, -2.03 DPS) [dungeon]; Nethergeld Shoulders (254049, -4.12 DPS, sim-verified) [crafted] |
| back | Caretaker's Cape (19531) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -1.15 DPS) [dungeon]; Runecloth Cloak (13860, -2.81 DPS) [crafted]; Prelacy Cape (7004, -3.83 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 59.8 healing_power points (20.82 DPS) | yes | Robes of Insight (940, -2.18 DPS, sim-verified) [world_drop]; Knight's Satin Armor (220892, -6.57 DPS) [vendor]; Funeral Pyre Vestment (12542, -8.22 DPS) [dungeon] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 26.6 healing_power points (9.26 DPS) | yes | Aristocratic Cuffs (12546, -1.32 DPS, sim-verified) [dungeon]; Bloodband Bracers (11469, -4.83 DPS) [quest]; Spidertank Oilrag (9448, -5.24 DPS) [dungeon] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 38.6 healing_power points (13.45 DPS) | yes | Earthenweave Gloves (254075, -3.54 DPS) [crafted]; Gilded Handwraps (254021, -3.87 DPS) [crafted]; Sergeant Major's Satin Gloves (220897, -5.10 DPS, sim-verified) [vendor] |
| waist | Gilded Waistcord (254081) | Tailoring [crafted] | 36.0 healing_power points (12.56 DPS) | yes | Dawnspire Cord (12466, -0.50 DPS, sim-verified) [dungeon]; Gilded Cord (254037, -2.25 DPS) [crafted]; Earthenweave Cord (254077, -2.65 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 46.5 healing_power points (16.19 DPS) | yes | Knight's Satin Leggings (220893, -2.65 DPS, sim-verified) [vendor]; Spellshock Leggings (9484, -5.45 DPS) [dungeon]; Stoneweaver Leggings (9407, -6.11 DPS) [dungeon] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 38.0 healing_power points (13.25 DPS) | yes | Sergeant Major's Satin Boots (220895, -1.77 DPS, sim-verified) [vendor]; Earthenweave Boots (254093, -3.35 DPS) [crafted]; Vinerot Sandals (17748, -3.76 DPS) [dungeon] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 19.9 healing_power points (6.94 DPS) | yes | Cyclopean Band (11824, -1.59 DPS) [dungeon]; Runemender's Seal (273031, -1.63 DPS) [world]; Philanthropist's Ring (281635, -2.10 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 15.4 healing_power points (5.36 DPS) | yes | Runemender's Seal (273031, -0.05 DPS) [world]; Philanthropist's Ring (281635, -0.51 DPS) [quest]; Cyclopean Band (11824, -0.89 DPS, sim-verified) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Prayer Beads (19990, +0.00 DPS) [quest]; Evonice's Landin' Pilla (18951, -0.15 DPS) [quest]; Uther's Strength (11302, -3.58 DPS, sim-verified) [world_drop] |
| trinket2 | Enriched Seal (271903) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Prayer Beads (19990, +0.00 DPS) [quest]; Evonice's Landin' Pilla (18951, -0.70 DPS) [quest]; Uther's Strength (11302, -1.60 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -9.05 DPS, sim-verified) [dungeon]; Resurgence Rod (17743, -18.60 DPS) [quest]; Dreamstaff (249454, -22.81 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Earthen Rod (9381) | Uldaman [dungeon] | 9.0 healing_power points (3.14 DPS) | yes | Charged Lightning Rod (11860, -0.30 DPS) [quest]; Lesser Eternal Wand (249232, -0.35 DPS) [crafted]; Dreambough Wand (249234, -0.70 DPS) [crafted] |

**New at 50:** head: Cassandra's Grace; shoulder: Mantle of Lost Hope; back: Caretaker's Cape; wrist: Nethergeld Cuffs; hands: Gilded Gloves; waist: Gilded Waistcord; legs: Senior Designer's Pantaloons; feet: Gilded Sandals; finger1: Brainlash; finger2: Lorekeeper's Ring; trinket2: Enriched Seal; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 606, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (human, 005005031030000000-33505003030121431-00000000000000000)

Set DPS (verified): 417.7. Weights run: 26.7s. Verify run: 724.0s. 1508 eligible items had no known source.

Race sweep: human 417.7 ± 0.3, gnome 398.9 ± 0.3, night-elf 386.0 ± 0.3, dwarf 381.7 ± 0.3 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.183, intellect=1.355 ± 0.065, spirit=3.297 ± 0.086, mp5=6.024 ± 0.110, crit=0.316 ± 0.025 per rating point (14 rating = 1%, 4.418 per %), spell_haste=not significant (0.245 ± 0.466)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | 102.0 healing_power points (17.85 DPS) | yes | Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Crown of Caer Darrow (13986, -8.99 DPS, sim-verified) [quest] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 75.4 healing_power points (13.19 DPS) | yes | Animated Chain Necklace (18723, -3.38 DPS) [dungeon]; Dragon's Eye (10829, -3.50 DPS, sim-verified) [dungeon]; Heart of the Fiend (13960, -3.57 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Shoulders (14139, +0.00 DPS) [crafted]; Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Mantle of the Scarlet Crusade (22405, -6.79 DPS, sim-verified) [dungeon] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Caretaker's Cape (19530, -0.48 DPS) [rep]; Shroud of the Exile (15421, -1.13 DPS) [quest]; Featherskin Cape (10843, -1.19 DPS, sim-verified) [world] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Satin Tunic (231624, +0.00 DPS) [vendor]; Ghostweave Vest (14141, -0.46 DPS) [crafted]; Robes of the Exalted (13346, -1.38 DPS, sim-verified) [dungeon] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Satin Bracers (17606, +0.00 DPS) [pvp]; Dryad's Wrist Bindings (19595, +0.00 DPS) [rep]; Wyrmthalak's Shackles (13958, -4.96 DPS, sim-verified) [quest] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swarmtender's Gloves (275607, -0.47 DPS) [crafted]; Gloves of the Atal'ai Prophet (10808, -0.99 DPS) [dungeon]; Hands of the Exalted Herald (12554, -6.82 DPS, sim-verified) [dungeon] |
| waist | Penitent's Cinch (272394) | Pix Xizzix [vendor] | 102.1 healing_power points (17.87 DPS) | yes | Wisdom of the Timbermaw (19047, -1.70 DPS, sim-verified) [crafted]; Marshal's Satin Sash (17609, -3.76 DPS) [pvp]; Grimgore Noose (13403, -5.91 DPS) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 120.2 healing_power points (21.03 DPS) | yes | Marshal's Satin Legguards (231626, -0.53 DPS) [vendor]; Knight-Captain's Satin Legguards (227125, -2.39 DPS) [pvp]; Wolfshear Leggings (13206, -10.55 DPS, sim-verified) [dungeon] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Faith Healer's Boots (22247, +0.00 DPS) [dungeon]; Marshal's Satin Walkers (231627, -2.13 DPS) [vendor]; Mooncloth Boots (15802, -7.02 DPS, sim-verified) [crafted] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, -0.02 DPS) [dungeon]; The Postmaster's Seal (13392, -0.64 DPS) [dungeon]; Naglering (11669, -12.28 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | The Postmaster's Seal (13392, -0.13 DPS) [dungeon]; Rosewine Circle (13178, -0.99 DPS, sim-verified) [dungeon]; Band of the Hierophant (13096, -1.32 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+8.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Shard of the Splithooves (10659, -0.60 DPS) [quest]; Evonice's Landin' Pilla (18951, -1.15 DPS) [quest]; Enriched Seal (271903, -1.65 DPS) [rep] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -1.42 DPS, sim-verified) [quest]; Evonice's Landin' Pilla (18951, -2.67 DPS) [quest]; Enriched Seal (271903, -3.16 DPS) [rep] |
| main_hand | Redemption (22406) | Stratholme: Cannon Master Willey [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Warhammer (23454, +0.00 DPS) [vendor]; Staff of Metanoia (22394, -2.26 DPS, sim-verified) [dungeon]; Glowing Brightwood Staff (812, -4.18 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Eternal Light (22254) | Blackrock Depths: Watchman Doomgrip [dungeon] | 29.2 healing_power points (5.11 DPS) | yes | Brilliant Wand (249385, -1.87 DPS) [crafted]; Jaina's Firestarter (13064, -1.95 DPS) [world_drop]; Cairnstone Sliver (9654, -2.30 DPS, sim-verified) [quest] |

**New at 60:** head: Devout Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; back: Drape of Recovery; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Penitent's Cinch; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of Mending; finger2: Blessed Band of Light; trinket2: Burst of Knowledge; main_hand: Redemption; ranged: Wand of Eternal Light

No-known-source sample (15 of 1508, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (human, 005005031030000000-33505003030121431-00000000000000000)

Set DPS (verified): 871.3. Weights run: 17.6s. Verify run: 429.5s. 1508 eligible items had no known source.

Race sweep: human 870.8 ± 0.5, gnome 856.8 ± 0.5, night-elf 841.7 ± 0.5, dwarf 838.8 ± 0.5 (heal_score); the guide race human is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.193, intellect=0.513 ± 0.040, spirit=0.748 ± 0.039, mp5=1.366 ± 0.042, crit=0.410 ± 0.029 per rating point (14 rating = 1%, 5.737 per %), spell_haste=not significant (0.199 ± 0.743)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -9.70 DPS, sim-verified) [crafted] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heart of the Fiend (13960, -2.18 DPS) [dungeon]; Animated Chain Necklace (18723, -3.63 DPS, sim-verified) [dungeon]; Dragon's Eye (10829, -4.38 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp]; Mooncloth Shoulders (14139, -8.07 DPS, sim-verified) [crafted] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Caretaker's Cape (19530, -1.89 DPS, sim-verified) [rep]; Shroud of the Exile (15421, -5.87 DPS) [quest] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 85.0 healing_power points (35.41 DPS) | yes | Field Marshal's Satin Tunic (231624, -0.78 DPS) [vendor]; Robes of the Exalted (13346, -9.23 DPS) [dungeon]; Stormcloth Vest (10020, -10.33 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wyrmthalak's Shackles (13958, +0.00 DPS) [quest]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted]; Dryad's Wrist Bindings (19595, -9.37 DPS, sim-verified) [rep] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Lieutenant's Satin Handwraps (227128, +0.00 DPS) [pvp]; Marshal's Satin Handwraps (231623, +0.00 DPS) [vendor]; Hands of the Exalted Herald (12554, -5.98 DPS, sim-verified) [dungeon] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 44.9 healing_power points (18.72 DPS) | yes | Penitent's Cinch (272394, -2.26 DPS, sim-verified) [vendor]; Earthenweave Cord (254077, -4.46 DPS) [crafted]; Gilded Waistcord (254081, -4.71 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Satin Legguards (227125, +0.00 DPS) [pvp]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Senior Designer's Pantaloons (11841, -12.39 DPS, sim-verified) [dungeon] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Satin Walkers (231627, +0.00 DPS) [vendor]; Swarmtender's Footpads (275609, +0.00 DPS) [crafted]; Faith Healer's Boots (22247, -7.49 DPS, sim-verified) [dungeon] |
| finger1 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, -0.42 DPS) [dungeon]; Rosewine Circle (13178, -3.29 DPS) [dungeon]; Naglering (11669, -17.12 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, +0.00 DPS) [dungeon]; Rosewine Circle (13178, -2.00 DPS) [dungeon]; Maiden's Circle (13001, -4.54 DPS) [world_drop] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor]; Shard of the Splithooves (10659, -1.50 DPS, sim-verified) [quest] |
| trinket2 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, -3.05 DPS, sim-verified) [vendor] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Argent Crusader (13249, -5.73 DPS) [quest]; Redemption (22406, -6.36 DPS) [dungeon]; Staff of Metanoia (22394, -6.67 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Eternal Light (22254) | Blackrock Depths: Watchman Doomgrip [dungeon] | 14.8 healing_power points (6.17 DPS) | yes | Brilliant Wand (249385, -0.92 DPS) [crafted]; Ritssyn's Wand of Bad Mojo (22408, -1.59 DPS) [dungeon]; Bonecreeper Stylus (13938, -2.34 DPS, sim-verified) [dungeon] |

**New at 60:** head: Devout Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; back: Drape of Recovery; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Devout Sandals; finger1: Fordring's Seal; finger2: Blessed Band of Light; trinket1: Burst of Knowledge; trinket2: Dormant Heart of the Mountain; ranged: Wand of Eternal Light

No-known-source sample (15 of 1508, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

## Horde

### Band 20 (undead, 000000000000000000-03502001000000000-00000000000000000)

Set DPS (verified): 40.6. Weights run: 14.7s. Verify run: 67.3s. 196 eligible items had no known source.

Race sweep: undead 40.6 ± 0.0, troll 37.5 ± 0.0 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.058, intellect=0.883 ± 0.011, spirit=0.887 ± 0.012, mp5=1.794 ± 0.011, crit=0.060 ± 0.004 per rating point (14 rating = 1%, 0.840 per %), spell_haste=not significant (-0.118 ± 0.066)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.80 DPS) | yes | Shadow Goggles (4373, -0.83 DPS) [crafted]; Flying Tiger Goggles (4368, -1.45 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Roadwatcher's Confidence (281265, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.9 healing_power points (1.30 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.44 DPS) [crafted]; Slime-encrusted Pads (6461, -0.76 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.90 DPS) [crafted] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.32 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.61 DPS) [crafted]; Spirit Cloak (4792, -0.72 DPS) [vendor] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Corsair's Overshirt (5202, -0.20 DPS) [dungeon]; Robe of the Moccasin (6465, -0.22 DPS) [dungeon]; Scholarly Robes (2034, -0.64 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 9.8 healing_power points (0.98 DPS) | yes | Bright Bracers (3647, -0.63 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.68 DPS) [quest]; Owlbeard Bracers (16981, -0.70 DPS) [quest] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 13.6 healing_power points (1.37 DPS) | yes | Magefist Gloves (12977, -0.57 DPS) [world_drop]; Bright Gloves (3066, -0.66 DPS) [world_drop]; Serpent Gloves (5970, -0.67 DPS) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 14.5 healing_power points (1.46 DPS) | yes | Novice Ardent's Sash (253887, -0.32 DPS) [crafted]; Keller's Girdle (2911, -0.75 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.79 DPS) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 26.8 healing_power points (2.69 DPS) | yes | Darkweave Breeches (12987, -1.54 DPS) [world_drop]; Colorful Kilt (10048, -1.74 DPS) [crafted]; Filigreed Silky Leggings (253939, -1.80 DPS) [crafted] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 11.6 healing_power points (1.17 DPS) | yes | Spidersilk Boots (4320, -0.11 DPS) [crafted]; Red Woolen Boots (4313, -0.32 DPS) [crafted]; Woolen Boots (2583, -0.61 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 8.6 healing_power points (0.86 DPS) | yes | Band of Purification (12996, -0.33 DPS) [world_drop]; Loop of Sacrifice (281673, -0.42 DPS) [quest]; Volcanic Rock Ring (12053, -0.60 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Purification (12996, +0.00 DPS) [world_drop]; Loop of Sacrifice (281673, -0.09 DPS) [quest]; Volcanic Rock Ring (12053, -0.27 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Living Root (6631) | Wailing Caverns: Verdan the Everliving [dungeon] | 49.4 healing_power points (4.95 DPS) | yes | Staff of the Blessed Seer (2271, -1.26 DPS) [dungeon]; Gnarled Necromancer's Staff (251534, -1.46 DPS) [quest]; Evocator's Blade (2567, -1.53 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Decay (5252) | Beren's Peril [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Torchlight Wand (5240, +0.00 DPS) [quest]; Sizzle Stick (8071, +0.00 DPS) [quest]; Greater Magic Wand (11288, +0.00 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Living Root; ranged: Wand of Decay

No-known-source sample (15 of 196, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6360 Steelscale Crushfish; 6478 Rat Stompers; 8350 The 1 Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 20425 Advisor's Gnarled Staff

### Band 30 (undead, 000000000000000000-33502003030110000-00000000000000000)

Set DPS (verified): 79.9. Weights run: 14.9s. Verify run: 87.5s. 318 eligible items had no known source.

Race sweep: undead 79.9 ± 0.0, troll 71.1 ± 0.0 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.143, intellect=1.105 ± 0.017, spirit=0.506 ± 0.016, mp5=2.482 ± 0.032, crit=0.162 ± 0.012 per rating point (14 rating = 1%, 2.271 per %), spell_haste=not significant (-0.082 ± 0.158)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 36.0 healing_power points (4.96 DPS) | yes | Filigreed Pristine Circlet (253975, -1.91 DPS, sim-verified) [crafted]; Pristine Circlet (253949, -2.48 DPS) [crafted]; Enchanter's Cowl (4322, -2.61 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 6.6 healing_power points (0.91 DPS) | yes | Crystal Starfire Medallion (5003, -0.10 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.15 DPS) [vendor]; Kaleidoscope Chain (13084, -0.85 DPS, sim-verified) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 18.9 healing_power points (2.61 DPS) | yes | Death Speaker Mantle (6685, -0.11 DPS) [dungeon]; Magician's Mantle (12998, -0.55 DPS) [world_drop]; Feline Mantle (3748, -0.84 DPS) [dungeon] |
| back | Battle Healer's Cloak (19529) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Raider's Cloak (272078, -0.64 DPS) [vendor]; Windsong Drape (15468, -1.10 DPS) [quest]; Glowing Thresher Cape (6901, -1.32 DPS, sim-verified) [dungeon] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 32.1 healing_power points (4.42 DPS) | yes | Beguiler Robes (7728, -0.80 DPS) [dungeon]; Green Silk Armor (7065, -1.21 DPS) [crafted]; Pristine Gown (253961, -1.22 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 16.4 healing_power points (2.26 DPS) | yes | Mindthrust Bracers (1974, -0.51 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -1.05 DPS) [world_drop]; Nightsky Wristbands (6407, -1.14 DPS) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 18.3 healing_power points (2.52 DPS) | yes | Pristine Gloves (253913, -0.52 DPS, sim-verified) [crafted]; Gloves of Old (9395, -0.92 DPS) [world_drop]; Jutebraid Gloves (10654, -0.93 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 15.4 healing_power points (2.12 DPS) | yes | Defiler's Cloth Girdle (20164, -0.15 DPS) [rep]; Crimson Silk Belt (7055, -0.23 DPS) [crafted]; Belt of Arugal (6392, -1.07 DPS, sim-verified) [dungeon] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 32.3 healing_power points (4.44 DPS) | yes | Filigreed Pristine Leggings (253937, -0.85 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -1.61 DPS) [crafted]; Leech Pants (6910, -2.04 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 29.8 healing_power points (4.09 DPS) | yes | Acidic Walkers (9454, -1.80 DPS, sim-verified) [dungeon]; Pristine Boots (253889, -2.40 DPS) [crafted]; Spidersilk Boots (4320, -2.52 DPS) [crafted] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 19.2 healing_power points (2.64 DPS) | yes | Solid Gold Band (279845, -0.78 DPS) [quest]; Ring of Power Regulation (273030, -0.79 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.91 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 17.3 healing_power points (2.38 DPS) | yes | Ring of Power Regulation (273030, -0.53 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.65 DPS) [dungeon]; Solid Gold Band (279845, -0.97 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Odo's Ley Staff (6318, -1.37 DPS) [dungeon]; Living Root (6631, -2.82 DPS) [dungeon]; Staff of the Friar (3415, -3.94 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 16.0 healing_power points (2.21 DPS) | yes | Orb of Mystic Insight (249394, -0.43 DPS, sim-verified) [crafted]; Orb of Mistmantle (13031, -0.69 DPS) [world_drop]; Witch's Finger (16887, -0.99 DPS) [quest] |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 9.4 healing_power points (1.29 DPS) | yes | Wand of Eventide (5214, -0.60 DPS) [world_drop]; Greater Mystic Wand (217287, -0.60 DPS) [crafted]; Thunderwood (13062, -0.86 DPS, sim-verified) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Battle Healer's Cloak; chest: Death Speaker Robes; wrist: Spidertank Oilrag; hands: Truefaith Gloves; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Runemender's Seal; finger2: Snake Hoop; main_hand: Death Speaker Scepter; off_hand: Orb of Souls; ranged: Wand of Mana Concentration

No-known-source sample (15 of 318, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6360 Steelscale Crushfish; 6478 Rat Stompers; 8350 The 1 Ring

### Band 40 (undead, 000000000000000000-33502003030121431-00000000000000000)

Set DPS (verified): 184.6. Weights run: 26.2s. Verify run: 167.8s. 454 eligible items had no known source.

Race sweep: undead 184.7 ± 0.1, troll 168.4 ± 0.2 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.333, intellect=0.628 ± 0.041, spirit=0.993 ± 0.040, mp5=1.492 ± 0.044, crit=0.124 ± 0.031 per rating point (14 rating = 1%, 1.736 per %), spell_haste=not significant (0.034 ± 0.351)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 41.6 healing_power points (12.25 DPS) | yes | Whitemane's Chapeau (7720, -1.01 DPS) [dungeon]; Earthen Silk Hood (254015, -2.85 DPS) [crafted]; Holy Shroud (2721, -3.81 DPS, sim-verified) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 25.0 healing_power points (7.37 DPS) | yes | Triune Amulet (7722, -3.44 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -4.03 DPS) [dungeon]; Amberglow Talisman (10824, -4.44 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 30.2 healing_power points (8.88 DPS) | yes | Earthen Silk Shoulders (254033, -3.01 DPS, sim-verified) [crafted]; Mantle of Doan (7712, -3.37 DPS) [dungeon]; Naga Priestess's Mantle (279889, -4.18 DPS) [quest] |
| back | Battle Healer's Cloak (19528) | Warsong Outriders [rep] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Glowing Thresher Cape (6901, -2.36 DPS) [dungeon]; Icy Cloak (4327, -2.95 DPS) [crafted]; Cloak of Blight (6832, -2.96 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 65.9 healing_power points (19.42 DPS) | yes | Civinad Robes (9623, -5.71 DPS, sim-verified) [quest]; Dreamweave Vest (10021, -8.36 DPS) [crafted]; Death's Head Vestment (10581, -8.84 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (184.6 DPS) | yes | Earthen Silk Cuffs (254019, -0.64 DPS, sim-verified) [crafted]; Radiant Silver Bracers (4545, -1.31 DPS) [quest]; Mindthrust Bracers (1974, -1.73 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 27.5 healing_power points (8.09 DPS) | yes | Gilded Handwraps (254021, -0.28 DPS, sim-verified) [crafted]; Earthen Silk Gloves (254017, -0.45 DPS) [crafted]; Stormcloth Gloves (10011, -2.63 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 33.0 healing_power points (9.71 DPS) | yes | Defiler's Cloth Girdle (20166, -4.85 DPS) [rep]; Deathmage Sash (10771, -4.88 DPS) [dungeon]; Star Belt (4329, -4.88 DPS, sim-verified) [crafted] |
| legs | Stoneweaver Leggings (9407) | Uldaman: Ironaya [dungeon] | 36.9 healing_power points (10.88 DPS) | yes | Stormcloth Pants (10010, -2.07 DPS) [crafted]; Pristine Leggings (253987, -2.22 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -3.24 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 28.4 healing_power points (8.35 DPS) | yes | Earthen Silk Slippers (254013, -1.01 DPS) [crafted]; Thoughtcast Boots (10578, -1.88 DPS, sim-verified) [dungeon]; Furen's Boots (13100, -2.38 DPS) [world_drop] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 19.7 healing_power points (5.81 DPS) | yes | Solid Gold Band (279845, -1.11 DPS) [quest]; Electrocutioner Lagnut (9447, -1.11 DPS) [dungeon]; Advisor's Ring (19520, -1.84 DPS) [rep] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 17.3 healing_power points (5.11 DPS) | yes | Electrocutioner Lagnut (9447, -0.41 DPS) [dungeon]; Advisor's Ring (19520, -1.14 DPS) [rep]; Solid Gold Band (279845, -1.84 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+5.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Jordan (873, -8.33 DPS, sim-verified) [world_drop]; Death Speaker Scepter (2816, -10.00 DPS) [dungeon]; Staff of the Friar (3415, -10.64 DPS) [dungeon] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | 26.0 healing_power points (7.65 DPS) | yes | Prophetic Cane (6803, -0.51 DPS, sim-verified) [quest]; Orb of Souls (249395, -2.07 DPS) [crafted]; Orb of the Forgotten Seer (7685, -3.53 DPS) [dungeon] |
| ranged | Earthen Rod (9381) | Uldaman [dungeon] | 9.0 healing_power points (2.65 DPS) | yes | Captain Rackmore's Tiller (16789, -0.01 DPS) [quest]; Jaina's Firestarter (13064, -0.66 DPS) [world_drop]; Dancing Flame (6806, -0.73 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Battle Healer's Cloak; chest: Stormcloth Vest; hands: Dreamweave Gloves; waist: Gilded Cord; legs: Stoneweaver Leggings; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Hand of Righteousness; off_hand: Beacon of Hope; ranged: Earthen Rod

No-known-source sample (15 of 454, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger; 5743 Prismstone Ring

### Band 50 (undead, 005005000000000000-33502003030121431-00000000000000000)

Set DPS (verified): 272.4. Weights run: 28.0s. Verify run: 246.3s. 599 eligible items had no known source.

Race sweep: undead 272.4 ± 0.2, troll 245.8 ± 0.2 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.285, intellect=0.783 ± 0.045, spirit=0.222 ± 0.043, mp5=0.845 ± 0.045, crit=0.184 ± 0.033 per rating point (14 rating = 1%, 2.574 per %), spell_haste=not significant (-0.085 ± 0.480)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Cassandra's Grace (13102) | World drop [world_drop] | 46.0 healing_power points (16.02 DPS) | yes | Soulcatcher Halo (10630, -0.85 DPS, sim-verified) [dungeon]; Blood Guard's Satin Cover (220899, -0.90 DPS) [vendor]; Papal Fez (9431, -3.03 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 26.3 healing_power points (9.15 DPS) | yes | Gemshard Heart (17707, -0.87 DPS, sim-verified) [dungeon]; Jarkal's Enhancing Necklace (7888, -4.70 DPS) [quest]; Horizon Choker (13085, -5.02 DPS) [world_drop] |
| shoulder | Mantle of Lost Hope (22234) | Blackrock Depths: Lord Roccor [dungeon] | 38.0 healing_power points (13.23 DPS) | yes | Blood Guard's Satin Pads (220901, -1.67 DPS) [vendor]; Inquisitor's Shawl (19507, -2.03 DPS) [dungeon]; Nethergeld Shoulders (254049, -2.19 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (272.4 DPS) | yes | Battle Healer's Cloak (19528, -0.32 DPS) [rep]; Deep Woodlands Cloak (19121, -0.34 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 59.8 healing_power points (20.82 DPS) | yes | Robes of Insight (940, -2.78 DPS) [world_drop]; Stone Guard's Satin Armor (220903, -6.57 DPS) [vendor]; Funeral Pyre Vestment (12542, -8.22 DPS) [dungeon] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 26.6 healing_power points (9.26 DPS) | yes | Aristocratic Cuffs (12546, -1.34 DPS) [dungeon]; Bloodband Bracers (11469, -4.83 DPS) [quest]; Spidertank Oilrag (9448, -5.24 DPS) [dungeon] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 38.6 healing_power points (13.45 DPS) | yes | Greenleaf Handwraps (19116, -2.78 DPS) [quest]; First Sergeant's Satin Gloves (220898, -3.24 DPS, sim-verified) [vendor]; Earthenweave Gloves (254075, -3.54 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Gilded Cord (254037, -1.14 DPS) [crafted]; Gilded Waistcord (254081, -1.27 DPS, sim-verified) [crafted]; Earthenweave Cord (254077, -1.54 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 46.5 healing_power points (16.19 DPS) | yes | Stone Guard's Satin Leggings (220902, -2.21 DPS, sim-verified) [vendor]; Spellshock Leggings (9484, -5.45 DPS) [dungeon]; Stoneweaver Leggings (9407, -6.11 DPS) [dungeon] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 38.0 healing_power points (13.25 DPS) | yes | First Sergeant's Satin Boots (220900, -1.30 DPS, sim-verified) [vendor]; Earthenweave Boots (254093, -3.35 DPS) [crafted]; Vinerot Sandals (17748, -3.76 DPS) [dungeon] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 19.9 healing_power points (6.94 DPS) | yes | Cyclopean Band (11824, -1.59 DPS) [dungeon]; Runemender's Seal (273031, -1.63 DPS) [world]; Philanthropist's Ring (281635, -2.10 DPS) [quest] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 15.4 healing_power points (5.36 DPS) | yes | Cyclopean Band (11824, -0.01 DPS) [dungeon]; Runemender's Seal (273031, -0.05 DPS) [world]; Philanthropist's Ring (281635, -0.51 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Prayer Beads (19990, +0.00 DPS) [quest]; Evonice's Landin' Pilla (18951, -0.15 DPS) [quest]; Uther's Strength (11302, -1.15 DPS, sim-verified) [world_drop] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Blessed Prayer Beads (19990, +0.00 DPS) [quest]; Evonice's Landin' Pilla (18951, -0.70 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -6.85 DPS, sim-verified) [dungeon]; Resurgence Rod (17743, -18.60 DPS) [quest]; Dreamstaff (249454, -22.81 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Charged Lightning Rod (11860) | Ledger from Tanaris [quest] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Lesser Eternal Wand (249232, -0.05 DPS) [crafted]; Dancing Flame (6806, -0.35 DPS) [quest]; Earthen Rod (9381, -0.59 DPS, sim-verified) [dungeon] |

**New at 50:** head: Cassandra's Grace; shoulder: Mantle of Lost Hope; back: Spritecaster Cape; wrist: Nethergeld Cuffs; hands: Gilded Gloves; waist: Dawnspire Cord; legs: Senior Designer's Pantaloons; feet: Gilded Sandals; finger1: Brainlash; finger2: Advisor's Ring; trinket2: Darkspear Voodoo Seal; main_hand: Glowing Brightwood Staff; ranged: Charged Lightning Rod

No-known-source sample (15 of 599, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60 (undead, 005005031030000000-33505003030121431-00000000000000000)

Set DPS (verified): 424.1. Weights run: 26.7s. Verify run: 733.7s. 1503 eligible items had no known source.

Race sweep: undead 424.2 ± 0.3, troll 387.1 ± 0.3 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.183, intellect=1.355 ± 0.065, spirit=3.297 ± 0.086, mp5=6.024 ± 0.110, crit=0.316 ± 0.025 per rating point (14 rating = 1%, 4.418 per %), spell_haste=not significant (0.245 ± 0.466)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Crown of Caer Darrow (13986, -12.66 DPS, sim-verified) [quest] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 75.4 healing_power points (13.19 DPS) | yes | Animated Chain Necklace (18723, -3.38 DPS) [dungeon]; Heart of the Fiend (13960, -3.57 DPS) [dungeon]; Dragon's Eye (10829, -3.82 DPS, sim-verified) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Shoulders (14139, +0.00 DPS) [crafted]; Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Mantle of the Scarlet Crusade (22405, -9.56 DPS, sim-verified) [dungeon] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-verified (+3.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Battle Healer's Cloak (19526, -0.48 DPS) [rep]; Shroud of the Exile (15421, -1.13 DPS) [quest]; Featherskin Cape (10843, -2.96 DPS, sim-verified) [world] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Ghostweave Vest (14141, -0.46 DPS) [crafted]; Robes of the Exalted (13346, -1.59 DPS, sim-verified) [dungeon] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Satin Bracers (17619, +0.00 DPS) [pvp]; Dryad's Wrist Bindings (19595, +0.00 DPS) [rep]; Wyrmthalak's Shackles (13958, -8.27 DPS, sim-verified) [quest] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swarmtender's Gloves (275607, -0.47 DPS) [crafted]; Gloves of the Atal'ai Prophet (10808, -0.99 DPS) [dungeon]; Hands of the Exalted Herald (12554, -12.23 DPS, sim-verified) [dungeon] |
| waist | Penitent's Cinch (272394) | Pix Xizzix [vendor] | 102.1 healing_power points (17.87 DPS) | yes | Wisdom of the Timbermaw (19047, -3.17 DPS, sim-verified) [crafted]; General's Satin Cinch (17621, -3.76 DPS) [pvp]; Grimgore Noose (13403, -5.91 DPS) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 120.2 healing_power points (21.03 DPS) | yes | General's Satin Legguards (231634, -0.53 DPS) [vendor]; Legionnaire's Satin Legguards (227123, -2.39 DPS) [pvp]; Wolfshear Leggings (13206, -14.53 DPS, sim-verified) [dungeon] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Faith Healer's Boots (22247, +0.00 DPS) [dungeon]; General's Satin Walkers (231630, -2.13 DPS) [vendor]; Mooncloth Boots (15802, -11.51 DPS, sim-verified) [crafted] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -0.51 DPS) [vendor]; The Postmaster's Seal (13392, -0.64 DPS) [dungeon]; Naglering (11669, -14.41 DPS, sim-verified) [dungeon] |
| finger2 | Rosewine Circle (13178) | Blackrock Spire: Urok Doomhowl [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Postmaster's Seal (13392, -0.62 DPS) [dungeon]; Darkspear Signet (272068, -1.04 DPS) [vendor]; Blessed Band of Light (272407, -1.16 DPS, sim-verified) [vendor] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+10.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Shard of the Splithooves (10659, -0.60 DPS) [quest]; Evonice's Landin' Pilla (18951, -1.15 DPS) [quest]; Darkspear Voodoo Seal (272061, -1.65 DPS) [vendor] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -2.48 DPS, sim-verified) [quest]; Evonice's Landin' Pilla (18951, -2.67 DPS) [quest]; Darkspear Voodoo Seal (272061, -3.16 DPS) [vendor] |
| main_hand | Redemption (22406) | Stratholme: Cannon Master Willey [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Mace (23464, +0.00 DPS) [vendor]; Glowing Brightwood Staff (812, -4.18 DPS) [world_drop]; Staff of Metanoia (22394, -5.93 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Eternal Light (22254) | Blackrock Depths: Watchman Doomgrip [dungeon] | 29.2 healing_power points (5.11 DPS) | yes | Brilliant Wand (249385, -1.87 DPS) [crafted]; Jaina's Firestarter (13064, -1.95 DPS) [world_drop]; Chillnail Splinter (10704, -3.59 DPS, sim-verified) [quest] |

**New at 60:** head: Devout Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; back: Drape of Recovery; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Penitent's Cinch; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of Mending; finger2: Rosewine Circle; trinket2: Burst of Knowledge; main_hand: Redemption; ranged: Wand of Eternal Light

No-known-source sample (15 of 1503, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60, raid preset (undead, 005005031030000000-33505003030121431-00000000000000000)

Set DPS (verified): 882.0. Weights run: 17.6s. Verify run: 428.4s. 1503 eligible items had no known source.

Race sweep: undead 882.1 ± 0.5, troll 837.1 ± 0.5 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.425, intellect=0.724 ± 0.058, spirit=1.088 ± 0.058, mp5=1.696 ± 0.062, crit=0.442 ± 0.055 per rating point (14 rating = 1%, 6.188 per %), spell_haste=not significant (-0.610 ± 1.037)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -17.07 DPS, sim-verified) [crafted] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 41.9 healing_power points (23.53 DPS) | yes | Animated Chain Necklace (18723, -2.71 DPS, sim-verified) [dungeon]; Heart of the Fiend (13960, -5.07 DPS) [dungeon]; Dragon's Eye (10829, -7.07 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp]; Mooncloth Shoulders (14139, -18.34 DPS, sim-verified) [crafted] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battle Healer's Cloak (19526, -2.12 DPS, sim-verified) [rep]; Shroud of the Exile (15421, -6.13 DPS) [quest] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 90.4 healing_power points (50.75 DPS) | yes | Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Robes of the Exalted (13346, -5.78 DPS, sim-verified) [dungeon]; Stormcloth Vest (10020, -13.31 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dryad's Wrist Bindings (19595, +0.00 DPS) [rep]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted]; Wyrmthalak's Shackles (13958, -16.90 DPS, sim-verified) [quest] |
| hands | Hands of the Exalted Herald (12554) | Blackrock Depths: Princess Moira Bronzebeard  [dungeon] | 53.5 healing_power points (30.02 DPS) | yes | General's Satin Handwraps (231633, -0.03 DPS) [vendor]; Gilded Gloves (254095, -4.13 DPS, sim-verified) [crafted]; Blood Guard's Satin Handwraps (227126, -5.53 DPS) [pvp] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 50.9 healing_power points (28.59 DPS) | yes | Penitent's Cinch (272394, -1.54 DPS, sim-verified) [vendor]; Earthenweave Cord (254077, -7.28 DPS) [crafted]; Gilded Waistcord (254081, -8.65 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Senior Designer's Pantaloons (11841, -19.42 DPS, sim-verified) [dungeon] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Faith Healer's Boots (22247, +0.00 DPS) [dungeon]; General's Satin Walkers (231630, +0.00 DPS) [vendor]; Mooncloth Boots (15802, -15.79 DPS, sim-verified) [crafted] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -1.62 DPS) [vendor]; Rosewine Circle (13178, -4.73 DPS) [dungeon]; Naglering (11669, -18.03 DPS, sim-verified) [dungeon] |
| finger2 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -0.92 DPS) [vendor]; Rosewine Circle (13178, -4.03 DPS) [dungeon]; Royal Seal of Alexis (18022, -5.97 DPS) [quest] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor]; Shard of the Splithooves (10659, -4.53 DPS, sim-verified) [quest] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Metanoia (22394, -6.86 DPS, sim-verified) [dungeon]; Argent Crusader (13249, -7.40 DPS) [quest]; Redemption (22406, -9.20 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Eternal Light (22254) | Blackrock Depths: Watchman Doomgrip [dungeon] | 16.7 healing_power points (9.39 DPS) | yes | Brilliant Wand (249385, -1.49 DPS, sim-verified) [crafted]; Bonecreeper Stylus (13938, -1.59 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -3.22 DPS) [dungeon] |

**New at 60:** head: Devout Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; back: Drape of Recovery; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Hands of the Exalted Herald; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of Mending; finger2: Fordring's Seal; trinket1: Dormant Heart of the Mountain; trinket2: Burst of Knowledge; ranged: Wand of Eternal Light

No-known-source sample (15 of 1503, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

