# Leveling BiS: Holy

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 000000000000000000-03502001000000000-00000000000000000)

Set DPS (verified): 41.8. Weights run: 24.9s. Verify run: 147.0s. 203 eligible items had no known source.

Race sweep: human 41.8 ± 0.0, gnome 40.4 ± 0.0, night-elf 39.1 ± 0.0, dwarf 38.7 ± 0.0 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.059, intellect=0.970 ± 0.011, spirit=1.046 ± 0.015, mp5=2.052 ± 0.012, crit=0.070 ± 0.005 per rating point (14 rating = 1%, 0.977 per %), spell_haste=not significant (-0.118 ± 0.066)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.69 DPS) | yes | Flying Tiger Goggles (4368, -1.30 DPS) [crafted]; Shadow Goggles (4373, -1.35 DPS, sim-verified) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 2.1 healing_power points (0.20 DPS) | yes | Tarnished Locket (279870, -0.54 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.7 healing_power points (1.29 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.68 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.71 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.91 DPS) [crafted] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | sim-verified (41.8 DPS) | yes | Heavy Woolen Cloak (4311, -0.52 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -0.58 DPS) [crafted]; Spirit Cloak (4792, -0.65 DPS) [vendor] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 21.0 healing_power points (1.97 DPS) | yes | Corsair's Overshirt (5202, -0.15 DPS) [dungeon]; Scholarly Robes (2034, -0.51 DPS) [dungeon]; Robe of the Moccasin (6465, -0.59 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 11.0 healing_power points (1.03 DPS) | yes | Bright Bracers (3647, -0.67 DPS) [world_drop]; Seer's Cuffs (3645, -0.75 DPS) [dungeon]; Repurposed Hair Band (281256, -1.67 DPS, sim-verified) [quest] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 13.9 healing_power points (1.31 DPS) | yes | Bright Gloves (3066, -0.55 DPS) [world_drop]; Serpent Gloves (5970, -0.65 DPS) [dungeon]; Magefist Gloves (12977, -0.88 DPS, sim-verified) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 14.9 healing_power points (1.40 DPS) | yes | Novice Ardent's Sash (253887, -0.49 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.67 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.75 DPS) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 28.0 healing_power points (2.63 DPS) | yes | Colorful Kilt (10048, -1.67 DPS) [crafted]; Filigreed Silky Leggings (253939, -1.69 DPS) [crafted]; Darkweave Breeches (12987, -2.23 DPS, sim-verified) [world_drop] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 11.9 healing_power points (1.12 DPS) | yes | Red Woolen Boots (4313, -0.25 DPS) [crafted]; Spidersilk Boots (4320, -0.28 DPS, sim-verified) [crafted]; Kimbra Boots (6191, -0.46 DPS) [quest] |
| finger1 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 9.1 healing_power points (0.86 DPS) | yes | Band of Purification (12996, -0.27 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.31 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.58 DPS) [world_drop] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.9 healing_power points (0.65 DPS) | yes | Lavishly Jeweled Ring (1156, -0.11 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.38 DPS) [world_drop]; Band of Purification (12996, -0.43 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Westfall (2042) | The Defias Brotherhood [quest] | 59.1 healing_power points (5.56 DPS) | yes | Living Root (6631, -1.11 DPS, sim-verified) [dungeon]; Staff of the Blessed Seer (2271, -1.94 DPS) [dungeon]; Evocator's Blade (2567, -2.25 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Sable Wand (7607) (or Dwarven Flamestick (5241)) | Underground Assault [quest] | 2.1 healing_power points (0.20 DPS) | yes | Dwarven Flamestick (5241, +0.00 DPS) [quest]; Sizzle Stick (8071, -0.01 DPS) [quest]; Greater Magic Wand (11288, -0.01 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Staff of Westfall; ranged: Sable Wand

No-known-source sample (15 of 203, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6360 Steelscale Crushfish; 6478 Rat Stompers; 8350 The 1 Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (human, 000000000000000000-33502003030110000-00000000000000000)

Set DPS (verified): 76.2. Weights run: 25.1s. Verify run: 167.1s. 325 eligible items had no known source.

Race sweep: human 76.2 ± 0.0, gnome 72.2 ± 0.0, night-elf 70.4 ± 0.0, dwarf 69.8 ± 0.0 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.139, intellect=1.045 ± 0.018, spirit=0.496 ± 0.016, mp5=2.794 ± 0.036, crit=0.186 ± 0.012 per rating point (14 rating = 1%, 2.603 per %), spell_haste=not significant (0.220 ± 0.165)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 36.0 healing_power points (4.94 DPS) | yes | Filigreed Pristine Circlet (253975, -1.15 DPS, sim-verified) [crafted]; Pristine Circlet (253949, -2.47 DPS) [crafted]; Enchanter's Cowl (4322, -2.68 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 6.3 healing_power points (0.86 DPS) | yes | Crystal Starfire Medallion (5003, -0.08 DPS) [world_drop]; Seal of the Expedition (271920, -0.14 DPS) [rep]; Kaleidoscope Chain (13084, -0.21 DPS, sim-verified) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 18.4 healing_power points (2.53 DPS) | yes | Death Speaker Mantle (6685, -0.11 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.55 DPS) [world_drop]; Feline Mantle (3748, -0.83 DPS) [dungeon] |
| back | Caretaker's Cape (19533) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Prelacy Cape (7004, -0.17 DPS, sim-verified) [quest]; Repairman's Cape (9605, -0.46 DPS) [quest]; Glowing Thresher Cape (6901, -0.55 DPS) [dungeon] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 31.5 healing_power points (4.33 DPS) | yes | Pristine Gown (253961, -0.42 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -0.82 DPS) [dungeon]; Green Silk Armor (7065, -1.22 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 17.4 healing_power points (2.39 DPS) | yes | Mindthrust Bracers (1974, -0.49 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -1.24 DPS) [world_drop]; Nightsky Wristbands (6407, -1.32 DPS) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 18.1 healing_power points (2.49 DPS) | yes | Pristine Gloves (253913, -0.59 DPS, sim-verified) [crafted]; Shilly Mitts (9609, -0.92 DPS) [quest]; Gloves of Old (9395, -0.96 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 15.2 healing_power points (2.09 DPS) | yes | Highlander's Cloth Girdle (20099, -0.14 DPS) [rep]; Crimson Silk Belt (7055, -0.26 DPS) [crafted]; Belt of Arugal (6392, -0.36 DPS, sim-verified) [dungeon] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 31.8 healing_power points (4.37 DPS) | yes | Filigreed Pristine Leggings (253937, -0.58 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -1.56 DPS) [crafted]; Leech Pants (6910, -1.76 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 29.3 healing_power points (4.03 DPS) | yes | Acidic Walkers (9454, -1.99 DPS, sim-verified) [dungeon]; Pristine Boots (253889, -2.36 DPS) [crafted]; Spidersilk Boots (4320, -2.49 DPS) [crafted] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 18.2 healing_power points (2.51 DPS) | yes | Lorekeeper's Ring (19525, -0.78 DPS) [rep]; Electrocutioner Lagnut (9447, -0.79 DPS) [dungeon]; Snake Hoop (6750, -0.97 DPS) [quest] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 14.4 healing_power points (1.98 DPS) | yes | Lorekeeper's Ring (19525, -0.25 DPS) [rep]; Electrocutioner Lagnut (9447, -0.26 DPS) [dungeon]; Snake Hoop (6750, -0.43 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Odo's Ley Staff (6318, -1.38 DPS) [dungeon]; Staff of Westfall (2042, -1.62 DPS) [quest]; Staff of the Friar (3415, -3.72 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Eye of Paleth (2943, -0.04 DPS) [quest]; Orb of Mistmantle (13031, -0.31 DPS) [world_drop]; Orb of Souls (249395, -0.40 DPS, sim-verified) [crafted] |
| ranged | Thunderwood (13062) | World drop [world_drop] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Wand of Eventide (5214, -0.14 DPS) [world_drop]; Greater Mystic Wand (217287, -0.14 DPS) [crafted]; Wand of Mana Concentration (273043, -0.46 DPS, sim-verified) [world] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Caretaker's Cape; chest: Death Speaker Robes; wrist: Spidertank Oilrag; hands: Truefaith Gloves; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Runemender's Seal; finger2: Ring of Power Regulation; main_hand: Death Speaker Scepter; off_hand: Orb of Mystic Insight; ranged: Thunderwood

No-known-source sample (15 of 325, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6360 Steelscale Crushfish

### Band 40 (human, 000000000000000000-33502003030121431-00000000000000000)

Set DPS (verified): 175.1. Weights run: 112.3s. Verify run: 258.8s. 461 eligible items had no known source.

Race sweep: human 175.1 ± 0.2, gnome 171.2 ± 0.2, night-elf 168.8 ± 0.2, dwarf 167.8 ± 0.2 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.140, intellect=1.054 ± 0.031, spirit=1.065 ± 0.030, mp5=1.830 ± 0.033, crit=0.159 ± 0.015 per rating point (14 rating = 1%, 2.230 per %), spell_haste=not significant (0.171 ± 0.146)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 49.5 healing_power points (7.06 DPS) | yes | Whitemane's Chapeau (7720, -0.50 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -1.44 DPS) [world_drop]; Earthen Silk Hood (254015, -2.40 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 28.4 healing_power points (4.06 DPS) | yes | Necklace of Calisea (1714, -1.94 DPS) [dungeon]; Triune Amulet (7722, -1.95 DPS, sim-verified) [dungeon]; Amberglow Talisman (10824, -2.54 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 35.7 healing_power points (5.09 DPS) | yes | Earthen Silk Shoulders (254033, -1.33 DPS, sim-verified) [crafted]; Mantle of Doan (7712, -1.98 DPS) [dungeon]; Green Silken Shoulders (7057, -2.30 DPS) [crafted] |
| back | Caretaker's Cape (19532) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Expeditionary's Cape (271922, -1.07 DPS) [rep]; Glowing Thresher Cape (6901, -1.11 DPS) [dungeon]; Prelacy Cape (7004, -1.29 DPS, sim-verified) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 66.5 healing_power points (9.49 DPS) | yes | Death's Head Vestment (10581, -3.75 DPS) [dungeon]; Civinad Robes (9623, -3.96 DPS) [quest]; Dreamweave Vest (10021, -5.22 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Mindthrust Bracers (1974, -0.53 DPS) [dungeon]; Earthen Silk Cuffs (254019, -0.66 DPS, sim-verified) [crafted]; Silkstream Cuffs (16791, -0.71 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 29.7 healing_power points (4.23 DPS) | yes | Gilded Handwraps (254021, -0.04 DPS) [crafted]; Earthen Silk Gloves (254017, -0.45 DPS) [crafted]; Stormcloth Gloves (10011, -0.86 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 36.8 healing_power points (5.25 DPS) | yes | Highlander's Cloth Girdle (20098, -2.65 DPS) [rep]; Windchaser Cinch (14435, -2.68 DPS) [world_drop]; Deathmage Sash (10771, -2.77 DPS, sim-verified) [dungeon] |
| legs | Stoneweaver Leggings (9407) | Uldaman: Ironaya [dungeon] | 41.2 healing_power points (5.88 DPS) | yes | Pristine Leggings (253987, -1.20 DPS, sim-verified) [crafted]; Stormcloth Pants (10010, -1.46 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.80 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 31.6 healing_power points (4.51 DPS) | yes | Thoughtcast Boots (10578, -0.78 DPS, sim-verified) [dungeon]; Earthen Silk Slippers (254013, -1.02 DPS) [crafted]; Furen's Boots (13100, -1.03 DPS) [world_drop] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 21.7 healing_power points (3.10 DPS) | yes | Philanthropist's Ring (281635, -0.92 DPS) [quest]; Lorekeeper's Ring (19524, -1.03 DPS) [rep]; Welken Ring (5011, -1.28 DPS) [world_drop] |
| finger2 | Electrocutioner Lagnut (9447) | Gnomeregan: Electrocutioner 6000 [dungeon] | 16.5 healing_power points (2.35 DPS) | yes | Lorekeeper's Ring (19524, -0.28 DPS) [rep]; Philanthropist's Ring (281635, -0.29 DPS, sim-verified) [quest]; Welken Ring (5011, -0.53 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+4.5 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, -4.93 DPS) [dungeon]; Staff of the Friar (3415, -4.93 DPS) [dungeon]; Staff of Jordan (873, -7.91 DPS, sim-verified) [world_drop] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 28.5 healing_power points (4.07 DPS) | yes | Beacon of Hope (9393, -0.32 DPS) [dungeon]; Orb of Souls (249395, -1.31 DPS) [crafted]; Orb of the Forgotten Seer (7685, -2.07 DPS) [dungeon] |
| ranged | Captain Rackmore's Tiller (16789) | Claim Rackmore's Treasure! [quest] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Earthen Rod (9381, -0.03 DPS) [dungeon]; Wand of Mana Concentration (273043, -0.19 DPS) [world]; Jaina's Firestarter (13064, -0.70 DPS, sim-verified) [world_drop] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Caretaker's Cape; chest: Stormcloth Vest; hands: Dreamweave Gloves; waist: Gilded Cord; legs: Stoneweaver Leggings; finger2: Electrocutioner Lagnut; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Hand of Righteousness; off_hand: Orb of Lorica; ranged: Captain Rackmore's Tiller

No-known-source sample (15 of 461, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (human, 005005000000000000-33502003030121431-00000000000000000)

Set DPS (verified): 257.8. Weights run: 42.5s. Verify run: 383.6s. 608 eligible items had no known source.

Race sweep: human 257.6 ± 0.2, gnome 251.2 ± 0.2, night-elf 246.8 ± 0.2, dwarf 245.4 ± 0.2 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.127, intellect=1.123 ± 0.042, spirit=2.396 ± 0.053, mp5=3.923 ± 0.056, crit=0.214 ± 0.018 per rating point (14 rating = 1%, 2.990 per %), spell_haste=not significant (0.325 ± 0.237)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Cassandra's Grace (13102) | World drop [world_drop] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Satin Cover (220896, -0.37 DPS) [vendor]; Papal Fez (9431, -0.46 DPS) [dungeon]; Gemburst Circlet (10751, -1.99 DPS, sim-verified) [quest] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (+5.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Horizon Choker (13085, -0.58 DPS) [world_drop]; Gemshard Heart (17707, -0.61 DPS) [dungeon]; Lei of Lilies (1315, -5.47 DPS, sim-verified) [world_drop] |
| shoulder | Mantle of Lost Hope (22234) | Blackrock Depths: Lord Roccor [dungeon] | 54.0 healing_power points (8.53 DPS) | yes | Knight-Lieutenant's Satin Pads (220894, -0.81 DPS) [vendor]; Nethergeld Shoulders (254049, -1.59 DPS, sim-verified) [crafted]; Kentic Amice (11624, -1.75 DPS) [dungeon] |
| back | Caretaker's Cape (19531) | Silverwing Sentinels [rep] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Spritecaster Cape (11623, -0.20 DPS) [dungeon]; Wingveil Cloak (10802, -0.48 DPS) [dungeon]; Featherskin Cape (10843, -2.40 DPS, sim-verified) [world] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 116.1 healing_power points (18.33 DPS) | yes | Stormcloth Vest (10020, -6.15 DPS) [crafted]; Ghostweave Vest (14141, -7.03 DPS, sim-verified) [crafted]; Grimlok's Tribal Vestments (9415, -7.57 DPS) [dungeon] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 39.8 healing_power points (6.29 DPS) | yes | Aristocratic Cuffs (12546, -0.99 DPS, sim-verified) [dungeon]; Venomshroud Armguards (14439, -2.55 DPS) [world_drop]; Bloodband Bracers (11469, -2.77 DPS) [quest] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | sim-verified (+5.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Earthenweave Gloves (254075, -1.11 DPS) [crafted]; Ghostweave Gloves (14142, -2.28 DPS) [crafted]; Gloves of the Atal'ai Prophet (10808, -5.02 DPS, sim-verified) [dungeon] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Earthenweave Cord (254077, -0.52 DPS, sim-verified) [crafted]; Ghostweave Cord (254073, -0.79 DPS) [crafted]; Dawnspire Cord (12466, -0.95 DPS) [dungeon] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 88.0 healing_power points (13.89 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -2.47 DPS) [dungeon]; Dalewind Trousers (13008, -2.84 DPS) [world_drop]; Rainstrider Leggings (11123, -5.76 DPS, sim-verified) [quest] |
| feet | Earthenweave Boots (254093) | Tailoring [crafted] | 52.4 healing_power points (8.27 DPS) | yes | Vinerot Sandals (17748, -0.05 DPS) [dungeon]; Furen's Boots (13100, -0.97 DPS) [world_drop]; Mistwalker Boots (10629, -3.42 DPS, sim-verified) [dungeon] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 35.5 healing_power points (5.60 DPS) | yes | Choking Band (11868, -0.69 DPS) [quest]; Runemender's Seal (273031, -0.85 DPS) [world]; Lorekeeper's Ring (19523, -1.23 DPS) [rep] |
| finger2 | Brainlash (6440) | Zul'Farrak [dungeon] | 34.2 healing_power points (5.40 DPS) | yes | Runemender's Seal (273031, -0.65 DPS) [world]; Lorekeeper's Ring (19523, -1.03 DPS) [rep]; Choking Band (11868, -4.83 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.76 DPS) [quest]; Enriched Seal (271903, -1.44 DPS) [rep]; Thunderbrew's Boot Flask (744, -1.51 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Enriched Seal (271903, +0.00 DPS) [rep]; Dormant Heart of the Mountain (249473, -1.24 DPS, sim-verified) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamstaff (249454, -8.46 DPS) [crafted]; Charstone Dirk (17710, -10.73 DPS) [dungeon]; Resurgence Rod (17743, -13.15 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cairnstone Sliver (9654) | The Morrow Stone [quest] | 17.6 healing_power points (2.78 DPS) | yes | Jaina's Firestarter (13064, -0.58 DPS) [world_drop]; Captain Rackmore's Tiller (16789, -0.70 DPS) [quest]; Wand of Mana Concentration (273043, -0.83 DPS) [world] |

**New at 50:** head: Cassandra's Grace; shoulder: Mantle of Lost Hope; back: Caretaker's Cape; chest: Embrace of the Wind Serpent; wrist: Nethergeld Cuffs; hands: Gilded Gloves; legs: Senior Designer's Pantaloons; feet: Earthenweave Boots; finger1: Eye of Adaegus; finger2: Brainlash; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Cairnstone Sliver

No-known-source sample (15 of 608, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (human, 005005031030000000-33505003030121431-00000000000000000)

Set DPS (verified): 419.3. Weights run: 39.9s. Verify run: 1106.4s. 1525 eligible items had no known source.

Race sweep: human 419.1 ± 0.3, gnome 399.5 ± 0.3, night-elf 387.1 ± 0.3, dwarf 383.7 ± 0.3 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.107, intellect=1.156 ± 0.057, spirit=3.420 ± 0.079, mp5=5.941 ± 0.096, crit=0.329 ± 0.019 per rating point (14 rating = 1%, 4.612 per %), spell_haste=not significant (0.573 ± 0.393)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Circlet (14140, +0.00 DPS) [crafted]; Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 73.0 healing_power points (13.28 DPS) | yes | Dragon's Eye (10829, -1.58 DPS) [dungeon]; Animated Chain Necklace (18723, -2.92 DPS) [dungeon]; Heart of the Fiend (13960, -3.06 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Shoulders (14139, +0.00 DPS) [crafted]; Mantle of the Scarlet Crusade (22405, +0.00 DPS) [dungeon]; Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Featherskin Cape (10843, +0.00 DPS) [world]; Caretaker's Cape (19530, -0.25 DPS) [rep]; Gracious Cape (18743, -0.73 DPS) [dungeon] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Embrace of the Wind Serpent (12462, +0.00 DPS) [world]; Robes of the Exalted (13346, +0.00 DPS) [dungeon]; Field Marshal's Satin Tunic (231624, -0.47 DPS) [vendor] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wyrmthalak's Shackles (13958, +0.00 DPS) [quest]; Marshal's Satin Bracers (17606, +0.00 DPS) [pvp]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of the Exalted Herald (12554, +0.00 DPS) [dungeon]; Swarmtender's Gloves (275607, -0.05 DPS) [crafted]; Gloves of the Atal'ai Prophet (10808, -0.60 DPS) [dungeon] |
| waist | Penitent's Cinch (272394) | Pix Xizzix [vendor] | 105.1 healing_power points (19.11 DPS) | yes | Wisdom of the Timbermaw (19047, -4.93 DPS) [crafted]; Marshal's Satin Sash (17609, -4.95 DPS) [pvp]; Grimgore Noose (13403, -7.11 DPS) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 120.0 healing_power points (21.82 DPS) | yes | Marshal's Satin Legguards (231626, -1.08 DPS) [vendor]; Wolfshear Leggings (13206, -1.55 DPS) [dungeon]; Rainstrider Leggings (11123, -2.38 DPS) [quest] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; Faith Healer's Boots (22247, -0.33 DPS) [dungeon]; Mistwalker Boots (10629, -1.84 DPS) [dungeon] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, +0.00 DPS) [dungeon]; The Postmaster's Seal (13392, -0.38 DPS) [dungeon]; Band of the Hierophant (13096, -1.90 DPS) [world_drop] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, +0.00 DPS) [dungeon]; The Postmaster's Seal (13392, -0.15 DPS) [dungeon]; Band of the Hierophant (13096, -1.67 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -0.98 DPS) [quest]; Evonice's Landin' Pilla (18951, -1.24 DPS) [quest]; Enriched Seal (271903, -2.06 DPS) [rep] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -2.16 DPS) [quest]; Evonice's Landin' Pilla (18951, -2.42 DPS) [quest]; Enriched Seal (271903, -3.24 DPS) [rep] |
| main_hand | Redemption (22406) | Stratholme: Cannon Master Willey [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Warhammer (23454, +0.00 DPS) [vendor]; Staff of Metanoia (22394, -2.86 DPS) [dungeon]; Glowing Brightwood Staff (812, -5.28 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Eternal Light (22254) | Blackrock Depths: Watchman Doomgrip [dungeon] | 27.8 healing_power points (5.06 DPS) | yes | Cairnstone Sliver (9654, -0.90 DPS) [quest]; Jaina's Firestarter (13064, -1.93 DPS) [world_drop]; Oblivion's Touch (18761, -1.95 DPS) [dungeon] |

**New at 60:** head: Devout Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; back: Drape of Recovery; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Penitent's Cinch; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of Mending; finger2: Blessed Band of Light; trinket2: Burst of Knowledge; main_hand: Redemption; ranged: Wand of Eternal Light

No-known-source sample (15 of 1525, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (human, 005005031030000000-33505003030121431-00000000000000000)

Set DPS (verified): 880.7. Weights run: 63.4s. Verify run: 568.2s. 1525 eligible items had no known source.

Race sweep: human 880.9 ± 0.5, gnome 867.4 ± 0.5, night-elf 851.2 ± 0.5, dwarf 848.9 ± 0.5 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.206, intellect=0.820 ± 0.052, spirit=0.922 ± 0.053, mp5=1.989 ± 0.054, crit=0.469 ± 0.037 per rating point (14 rating = 1%, 6.569 per %), spell_haste=not significant (0.961 ± 0.901)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Circlet (14140, +0.00 DPS) [crafted]; Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 42.0 healing_power points (13.65 DPS) | yes | Animated Chain Necklace (18723, -0.82 DPS) [dungeon]; Heart of the Fiend (13960, -3.50 DPS) [dungeon]; Tooth of Gnarr (13141, -3.90 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Shoulders (14139, +0.00 DPS) [crafted]; Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Caretaker's Cape (19530, -1.82 DPS) [rep]; Shroud of the Exile (15421, -4.00 DPS) [quest] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 87.8 healing_power points (28.54 DPS) | yes | Field Marshal's Satin Tunic (231624, +0.00 DPS) [vendor]; Robes of the Exalted (13346, -5.90 DPS) [dungeon]; Embrace of the Wind Serpent (12462, -6.15 DPS) [world] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wyrmthalak's Shackles (13958, +0.00 DPS) [quest]; Dryad's Wrist Bindings (19595, +0.00 DPS) [rep]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of the Exalted Herald (12554, +0.00 DPS) [dungeon]; Marshal's Satin Handwraps (231623, +0.00 DPS) [vendor]; Theramore Handwraps (271945, +0.00 DPS) [rep] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 50.8 healing_power points (16.52 DPS) | yes | Penitent's Cinch (272394, -1.84 DPS) [vendor]; Gilded Waistcord (254081, -4.68 DPS) [crafted]; Earthenweave Cord (254077, -4.76 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Senior Designer's Pantaloons (11841, +0.00 DPS) [dungeon]; Knight-Captain's Satin Legguards (227125, +0.00 DPS) [pvp]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; Faith Healer's Boots (22247, +0.00 DPS) [dungeon]; Marshal's Satin Walkers (231627, +0.00 DPS) [vendor] |
| finger1 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, +0.00 DPS) [dungeon]; Rosewine Circle (13178, -1.95 DPS) [dungeon]; Maiden's Circle (13001, -4.11 DPS) [world_drop] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, +0.00 DPS) [dungeon]; Rosewine Circle (13178, -0.79 DPS) [dungeon]; Maiden's Circle (13001, -2.96 DPS) [world_drop] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Argent Crusader (13249, -4.31 DPS) [quest]; Staff of Metanoia (22394, -4.71 DPS) [dungeon]; Redemption (22406, -5.25 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Eternal Light (22254) | Blackrock Depths: Watchman Doomgrip [dungeon] | 17.9 healing_power points (5.82 DPS) | yes | Bonecreeper Stylus (13938, -0.91 DPS) [dungeon]; Oblivion's Touch (18761, -1.03 DPS) [dungeon]; Brilliant Wand (249385, -1.03 DPS) [crafted] |

**New at 60:** head: Devout Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; back: Drape of Recovery; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Devout Sandals; finger1: Fordring's Seal; finger2: Blessed Band of Light; trinket1: Burst of Knowledge; trinket2: Dormant Heart of the Mountain; ranged: Wand of Eternal Light

No-known-source sample (15 of 1525, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

## Horde

### Band 20 (undead, 000000000000000000-03502001000000000-00000000000000000)

Set DPS (verified): 41.0. Weights run: 24.9s. Verify run: 136.2s. 196 eligible items had no known source.

Race sweep: undead 41.0 ± 0.0, troll 37.3 ± 0.0 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.059, intellect=0.970 ± 0.011, spirit=1.046 ± 0.015, mp5=2.052 ± 0.012, crit=0.070 ± 0.005 per rating point (14 rating = 1%, 0.977 per %), spell_haste=not significant (-0.118 ± 0.066)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.69 DPS) | yes | Flying Tiger Goggles (4368, -1.30 DPS) [crafted]; Shadow Goggles (4373, -1.94 DPS, sim-verified) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Roadwatcher's Confidence (281265, -0.18 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.7 healing_power points (1.29 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.68 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.71 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.91 DPS) [crafted] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pearl-clasped Cloak (5542, -0.58 DPS) [crafted]; Spirit Cloak (4792, -0.65 DPS) [vendor]; Heavy Woolen Cloak (4311, -0.71 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 21.0 healing_power points (1.97 DPS) | yes | Corsair's Overshirt (5202, -0.15 DPS) [dungeon]; Scholarly Robes (2034, -0.51 DPS) [dungeon]; Robe of the Moccasin (6465, -0.75 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 11.0 healing_power points (1.03 DPS) | yes | Crystalline Cuffs (14148, -0.74 DPS) [dungeon]; Seer's Cuffs (3645, -0.75 DPS) [dungeon]; Bright Bracers (3647, -1.68 DPS, sim-verified) [world_drop] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 13.9 healing_power points (1.31 DPS) | yes | Bright Gloves (3066, -0.55 DPS) [world_drop]; Serpent Gloves (5970, -0.65 DPS) [dungeon]; Magefist Gloves (12977, -1.11 DPS, sim-verified) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 14.9 healing_power points (1.40 DPS) | yes | Keller's Girdle (2911, -0.67 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.75 DPS) [crafted]; Novice Ardent's Sash (253887, -0.82 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 28.0 healing_power points (2.63 DPS) | yes | Colorful Kilt (10048, -1.67 DPS) [crafted]; Filigreed Silky Leggings (253939, -1.69 DPS) [crafted]; Darkweave Breeches (12987, -2.06 DPS, sim-verified) [world_drop] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 11.9 healing_power points (1.12 DPS) | yes | Red Woolen Boots (4313, -0.25 DPS) [crafted]; Spidersilk Boots (4320, -0.36 DPS, sim-verified) [crafted]; Smoldering Boots (3076, -0.53 DPS) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 9.1 healing_power points (0.86 DPS) | yes | Lavishly Jeweled Ring (1156, -0.31 DPS) [dungeon]; Loop of Sacrifice (281673, -0.40 DPS) [quest]; Volcanic Rock Ring (12053, -0.58 DPS) [world_drop] |
| finger2 | Band of Purification (12996) | World drop [world_drop] | 6.3 healing_power points (0.59 DPS) | yes | Loop of Sacrifice (281673, -0.13 DPS) [quest]; Volcanic Rock Ring (12053, -0.32 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.75 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Living Root (6631) | Wailing Caverns: Verdan the Everliving [dungeon] | 50.2 healing_power points (4.72 DPS) | yes | Gnarled Necromancer's Staff (251534, -1.29 DPS) [quest]; Staff of the Blessed Seer (2271, -1.36 DPS, sim-verified) [dungeon]; Evocator's Blade (2567, -1.42 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Sizzle Stick (8071) (or Greater Magic Wand (11288), Torchlight Wand (5240)) | Deviate Eradication [quest] | 2.0 healing_power points (0.19 DPS) | yes | Torchlight Wand (5240, +0.00 DPS) [quest]; Greater Magic Wand (11288, +0.00 DPS) [crafted]; Wand of Decay (5252, -0.01 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Advisor's Ring; finger2: Band of Purification; main_hand: Living Root; ranged: Sizzle Stick

No-known-source sample (15 of 196, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6360 Steelscale Crushfish; 6478 Rat Stompers; 8350 The 1 Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 20425 Advisor's Gnarled Staff

### Band 30 (undead, 000000000000000000-33502003030110000-00000000000000000)

Set DPS (verified): 79.3. Weights run: 25.1s. Verify run: 162.1s. 318 eligible items had no known source.

Race sweep: undead 79.4 ± 0.0, troll 70.8 ± 0.0 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.139, intellect=1.045 ± 0.018, spirit=0.496 ± 0.016, mp5=2.794 ± 0.036, crit=0.186 ± 0.012 per rating point (14 rating = 1%, 2.603 per %), spell_haste=not significant (0.220 ± 0.165)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 36.0 healing_power points (4.94 DPS) | yes | Filigreed Pristine Circlet (253975, -1.09 DPS, sim-verified) [crafted]; Pristine Circlet (253949, -2.47 DPS) [crafted]; Enchanter's Cowl (4322, -2.68 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 6.3 healing_power points (0.86 DPS) | yes | Crystal Starfire Medallion (5003, -0.08 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.14 DPS) [vendor]; Kaleidoscope Chain (13084, -0.27 DPS, sim-verified) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 18.4 healing_power points (2.53 DPS) | yes | Mantle of Woe (7750, -0.09 DPS, sim-verified) [quest]; Death Speaker Mantle (6685, -0.13 DPS) [dungeon]; Magician's Mantle (12998, -0.55 DPS) [world_drop] |
| back | Battle Healer's Cloak (19529) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Raider's Cloak (272078, -0.71 DPS) [vendor]; Glowing Thresher Cape (6901, -0.71 DPS, sim-verified) [dungeon]; Windsong Drape (15468, -1.10 DPS) [quest] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 31.5 healing_power points (4.33 DPS) | yes | Pristine Gown (253961, -0.56 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -0.82 DPS) [dungeon]; Green Silk Armor (7065, -1.22 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 17.4 healing_power points (2.39 DPS) | yes | Mindthrust Bracers (1974, -0.43 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -1.24 DPS) [world_drop]; Nightsky Wristbands (6407, -1.32 DPS) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 18.1 healing_power points (2.49 DPS) | yes | Pristine Gloves (253913, -0.49 DPS, sim-verified) [crafted]; Jutebraid Gloves (10654, -0.95 DPS) [quest]; Gloves of Old (9395, -0.96 DPS) [world_drop] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Defiler's Cloth Girdle (20164, -0.07 DPS) [rep]; Crimson Silk Belt (7055, -0.18 DPS) [crafted]; Pristine Sash (253925, -0.40 DPS, sim-verified) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 31.8 healing_power points (4.37 DPS) | yes | Filigreed Pristine Leggings (253937, -0.68 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -1.56 DPS) [crafted]; Leech Pants (6910, -1.76 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 29.3 healing_power points (4.03 DPS) | yes | Acidic Walkers (9454, -1.67 DPS, sim-verified) [dungeon]; Pristine Boots (253889, -2.36 DPS) [crafted]; Spidersilk Boots (4320, -2.49 DPS) [crafted] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 18.2 healing_power points (2.51 DPS) | yes | Advisor's Ring (19521, -0.78 DPS) [rep]; Electrocutioner Lagnut (9447, -0.79 DPS) [dungeon]; Ring of Power Regulation (273030, -1.66 DPS, sim-verified) [dungeon] |
| finger2 | Solid Gold Band (279845) | Opportunistic Education [quest] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Advisor's Ring (19521, -0.12 DPS) [rep]; Electrocutioner Lagnut (9447, -0.14 DPS) [dungeon]; Ring of Power Regulation (273030, -0.38 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Odo's Ley Staff (6318, -1.38 DPS) [dungeon]; Staff of the Friar (3415, -1.74 DPS, sim-verified) [dungeon]; Living Root (6631, -2.82 DPS) [dungeon] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Orb of Mistmantle (13031, -0.31 DPS) [world_drop]; Orb of Souls (249395, -0.41 DPS, sim-verified) [crafted]; Witch's Finger (16887, -0.67 DPS) [quest] |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 9.8 healing_power points (1.34 DPS) | yes | Thunderwood (13062, -0.52 DPS) [world_drop]; Wand of Eventide (5214, -0.66 DPS) [world_drop]; Greater Mystic Wand (217287, -0.66 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Battle Healer's Cloak; chest: Death Speaker Robes; wrist: Spidertank Oilrag; hands: Truefaith Gloves; waist: Belt of Arugal; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Runemender's Seal; finger2: Solid Gold Band; main_hand: Death Speaker Scepter; off_hand: Orb of Mystic Insight; ranged: Wand of Mana Concentration

No-known-source sample (15 of 318, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6360 Steelscale Crushfish; 6478 Rat Stompers; 8350 The 1 Ring

### Band 40 (undead, 000000000000000000-33502003030121431-00000000000000000)

Set DPS (verified): 182.0. Weights run: 112.3s. Verify run: 246.1s. 454 eligible items had no known source.

Race sweep: undead 181.8 ± 0.2, troll 166.7 ± 0.2 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.140, intellect=1.054 ± 0.031, spirit=1.065 ± 0.030, mp5=1.830 ± 0.033, crit=0.159 ± 0.015 per rating point (14 rating = 1%, 2.230 per %), spell_haste=not significant (0.171 ± 0.146)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 49.5 healing_power points (7.06 DPS) | yes | Whitemane's Chapeau (7720, -1.36 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -1.44 DPS) [world_drop]; Earthen Silk Hood (254015, -2.40 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 28.4 healing_power points (4.06 DPS) | yes | Necklace of Calisea (1714, -1.94 DPS) [dungeon]; Triune Amulet (7722, -2.43 DPS, sim-verified) [dungeon]; Amberglow Talisman (10824, -2.54 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 35.7 healing_power points (5.09 DPS) | yes | Mantle of Doan (7712, -1.98 DPS) [dungeon]; Earthen Silk Shoulders (254033, -2.18 DPS, sim-verified) [crafted]; Green Silken Shoulders (7057, -2.30 DPS) [crafted] |
| back | Battle Healer's Cloak (19528) | Warsong Outriders [rep] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Darkspear Raider's Cloak (272077, -1.07 DPS) [vendor]; Glowing Thresher Cape (6901, -1.11 DPS) [dungeon]; Cloak of Blight (6832, -1.35 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 66.5 healing_power points (9.49 DPS) | yes | Death's Head Vestment (10581, -3.75 DPS) [dungeon]; Civinad Robes (9623, -3.96 DPS) [quest]; Dreamweave Vest (10021, -5.30 DPS, sim-verified) [crafted] |
| wrist | Earthen Silk Cuffs (254019) | Tailoring [crafted] | 15.3 healing_power points (2.19 DPS) | yes | Spidertank Oilrag (9448, -0.12 DPS) [dungeon]; Radiant Silver Bracers (4545, -0.41 DPS) [quest]; Mindthrust Bracers (1974, -0.65 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 29.7 healing_power points (4.23 DPS) | yes | Gilded Handwraps (254021, -0.04 DPS) [crafted]; Earthen Silk Gloves (254017, -0.45 DPS) [crafted]; Stormcloth Gloves (10011, -0.86 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 36.8 healing_power points (5.25 DPS) | yes | Defiler's Cloth Girdle (20166, -2.65 DPS) [rep]; Windchaser Cinch (14435, -2.68 DPS) [world_drop]; Deathmage Sash (10771, -3.36 DPS, sim-verified) [dungeon] |
| legs | Stoneweaver Leggings (9407) | Uldaman: Ironaya [dungeon] | 41.2 healing_power points (5.88 DPS) | yes | Stormcloth Pants (10010, -1.46 DPS) [crafted]; Pristine Leggings (253987, -1.74 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -1.80 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 31.6 healing_power points (4.51 DPS) | yes | Earthen Silk Slippers (254013, -1.02 DPS) [crafted]; Furen's Boots (13100, -1.03 DPS) [world_drop]; Thoughtcast Boots (10578, -1.17 DPS, sim-verified) [dungeon] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 21.7 healing_power points (3.10 DPS) | yes | Solid Gold Band (279845, -0.77 DPS) [quest]; Philanthropist's Ring (281635, -0.92 DPS) [quest]; Advisor's Ring (19520, -1.03 DPS) [rep] |
| finger2 | Electrocutioner Lagnut (9447) | Gnomeregan: Electrocutioner 6000 [dungeon] | 16.5 healing_power points (2.35 DPS) | yes | Philanthropist's Ring (281635, -0.17 DPS) [quest]; Advisor's Ring (19520, -0.28 DPS) [rep]; Solid Gold Band (279845, -0.33 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+4.2 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, -4.93 DPS) [dungeon]; Staff of the Friar (3415, -4.93 DPS) [dungeon]; Staff of Jordan (873, -8.12 DPS, sim-verified) [world_drop] |
| off_hand | Prophetic Cane (6803) | Into The Scarlet Monastery [quest] | 26.4 healing_power points (3.77 DPS) | yes | Beacon of Hope (9393, -0.66 DPS, sim-verified) [dungeon]; Orb of Souls (249395, -1.00 DPS) [crafted]; Omega Orb (7749, -1.70 DPS) [quest] |
| ranged | Dancing Flame (6806) | Final Passage [quest] | 10.3 healing_power points (1.47 DPS) | yes | Captain Rackmore's Tiller (16789, -0.15 DPS) [quest]; Earthen Rod (9381, -0.18 DPS) [dungeon]; Jaina's Firestarter (13064, -1.09 DPS, sim-verified) [world_drop] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Battle Healer's Cloak; chest: Stormcloth Vest; wrist: Earthen Silk Cuffs; hands: Dreamweave Gloves; waist: Gilded Cord; legs: Stoneweaver Leggings; finger2: Electrocutioner Lagnut; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Hand of Righteousness; off_hand: Prophetic Cane; ranged: Dancing Flame

No-known-source sample (15 of 454, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger; 5743 Prismstone Ring

### Band 50 (undead, 005005000000000000-33502003030121431-00000000000000000)

Set DPS (verified): 272.3. Weights run: 42.5s. Verify run: 349.1s. 601 eligible items had no known source.

Race sweep: undead 272.1 ± 0.2, troll 244.9 ± 0.2 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.401, intellect=1.043 ± 0.063, spirit=0.455 ± 0.056, mp5=1.149 ± 0.060, crit=0.273 ± 0.045 per rating point (14 rating = 1%, 3.827 per %), spell_haste=not significant (0.163 ± 0.619)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Cassandra's Grace (13102) | World drop [world_drop] | 48.1 healing_power points (13.04 DPS) | yes | Blood Guard's Satin Cover (220899, +0.00 DPS) [vendor]; Papal Fez (9431, -1.16 DPS) [dungeon]; Soulcatcher Halo (10630, -1.36 DPS, sim-verified) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 28.3 healing_power points (7.69 DPS) | yes | Gemshard Heart (17707, -0.79 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -3.23 DPS) [world_drop]; Mindburst Medallion (11196, -3.75 DPS) [quest] |
| shoulder | Mantle of Lost Hope (22234) | Blackrock Depths: Lord Roccor [dungeon] | 42.1 healing_power points (11.41 DPS) | yes | Blood Guard's Satin Pads (220901, -1.37 DPS) [vendor]; Inquisitor's Shawl (19507, -1.77 DPS) [dungeon]; Nethergeld Shoulders (254049, -3.18 DPS, sim-verified) [crafted] |
| back | Battle Healer's Cloak (19527) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -0.47 DPS) [dungeon]; Deep Woodlands Cloak (19121, -0.91 DPS) [quest]; Runecloth Cloak (13860, -2.00 DPS) [crafted] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 61.6 healing_power points (16.72 DPS) | yes | Robes of Insight (940, -1.08 DPS, sim-verified) [world_drop]; Embrace of the Wind Serpent (12462, -1.24 DPS) [world]; Stone Guard's Satin Armor (220903, -4.52 DPS) [vendor] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 29.6 healing_power points (8.02 DPS) | yes | Aristocratic Cuffs (12546, -1.63 DPS, sim-verified) [dungeon]; Bloodband Bracers (11469, -3.75 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -4.26 DPS) [quest] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 43.1 healing_power points (11.68 DPS) | yes | Greenleaf Handwraps (19116, -2.60 DPS) [quest]; Earthenweave Gloves (254075, -3.27 DPS) [crafted]; First Sergeant's Satin Gloves (220898, -3.85 DPS, sim-verified) [vendor] |
| waist | Gilded Waistcord (254081) | Tailoring [crafted] | 38.4 healing_power points (10.41 DPS) | yes | Dawnspire Cord (12466, -0.15 DPS) [dungeon]; Gilded Cord (254037, -1.44 DPS) [crafted]; Earthenweave Cord (254077, -2.00 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 52.5 healing_power points (14.23 DPS) | yes | Stone Guard's Satin Leggings (220902, -2.92 DPS, sim-verified) [vendor]; Kilt of the Atal'ai Prophet (10807, -4.48 DPS) [dungeon]; Stoneweaver Leggings (9407, -5.07 DPS) [dungeon] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 40.4 healing_power points (10.95 DPS) | yes | First Sergeant's Satin Boots (220900, -1.91 DPS, sim-verified) [vendor]; Vinerot Sandals (17748, -2.29 DPS) [dungeon]; Earthenweave Boots (254093, -2.54 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 23.7 healing_power points (6.43 DPS) | yes | Runemender's Seal (273031, -1.55 DPS) [world]; Advisor's Ring (19519, -1.93 DPS) [rep]; Woodseed Hoop (17768, -1.97 DPS) [quest] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 18.1 healing_power points (4.91 DPS) | yes | Runemender's Seal (273031, -0.04 DPS) [world]; Advisor's Ring (19519, -0.41 DPS) [rep]; Woodseed Hoop (17768, -0.46 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+5.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Blessed Prayer Beads (19990, +0.00 DPS) [quest]; Evonice's Landin' Pilla (18951, -0.25 DPS) [quest] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Blessed Prayer Beads (19990, +0.00 DPS) [quest]; Evonice's Landin' Pilla (18951, -0.32 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -8.85 DPS, sim-verified) [dungeon]; Resurgence Rod (17743, -16.19 DPS) [quest]; Dreamstaff (249454, -19.17 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Dancing Flame (6806) | Final Passage [quest] | 10.2 healing_power points (2.77 DPS) | yes | Charged Lightning Rod (11860, -0.28 DPS) [quest]; Earthen Rod (9381, -0.33 DPS) [dungeon]; Lesser Eternal Wand (249232, -0.60 DPS) [crafted] |

**New at 50:** head: Cassandra's Grace; shoulder: Mantle of Lost Hope; back: Battle Healer's Cloak; wrist: Nethergeld Cuffs; hands: Gilded Gloves; waist: Gilded Waistcord; legs: Senior Designer's Pantaloons; feet: Gilded Sandals; finger1: Brainlash; finger2: Cyclopean Band; trinket2: Darkspear Voodoo Seal; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 601, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60 (undead, 005005031030000000-33505003030121431-00000000000000000)

Set DPS (verified): 429.6. Weights run: 39.9s. Verify run: 1077.8s. 1520 eligible items had no known source.

Race sweep: undead 429.6 ± 0.3, troll 391.9 ± 0.3 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.107, intellect=1.156 ± 0.057, spirit=3.420 ± 0.079, mp5=5.941 ± 0.096, crit=0.329 ± 0.019 per rating point (14 rating = 1%, 4.612 per %), spell_haste=not significant (0.573 ± 0.393)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -9.49 DPS, sim-verified) [crafted] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 73.0 healing_power points (13.28 DPS) | yes | Animated Chain Necklace (18723, -2.92 DPS) [dungeon]; Heart of the Fiend (13960, -3.06 DPS) [dungeon]; Dragon's Eye (10829, -3.85 DPS, sim-verified) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Shoulders (14139, +0.00 DPS) [crafted]; Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Mantle of the Scarlet Crusade (22405, -8.84 DPS, sim-verified) [dungeon] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-verified (+3.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Battle Healer's Cloak (19526, -0.25 DPS) [rep]; Gracious Cape (18743, -0.73 DPS) [dungeon]; Featherskin Cape (10843, -3.74 DPS, sim-verified) [world] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 147.2 healing_power points (26.76 DPS) | yes | Robes of the Exalted (13346, -2.26 DPS) [dungeon]; Truefaith Vestments (14154, -3.54 DPS) [crafted]; Warlord's Satin Tunic (231632, -4.00 DPS) [vendor] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Satin Bracers (17619, +0.00 DPS) [pvp]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted]; Wyrmthalak's Shackles (13958, -6.85 DPS, sim-verified) [quest] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swarmtender's Gloves (275607, -0.05 DPS) [crafted]; Gloves of the Atal'ai Prophet (10808, -0.60 DPS) [dungeon]; Hands of the Exalted Herald (12554, -9.45 DPS, sim-verified) [dungeon] |
| waist | Penitent's Cinch (272394) | Pix Xizzix [vendor] | 105.1 healing_power points (19.11 DPS) | yes | Wisdom of the Timbermaw (19047, -1.72 DPS, sim-verified) [crafted]; General's Satin Cinch (17621, -4.95 DPS) [pvp]; Grimgore Noose (13403, -7.11 DPS) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 120.0 healing_power points (21.82 DPS) | yes | General's Satin Legguards (231634, -1.08 DPS) [vendor]; Rainstrider Leggings (11123, -2.38 DPS) [quest]; Wolfshear Leggings (13206, -13.83 DPS, sim-verified) [dungeon] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Faith Healer's Boots (22247, -0.33 DPS) [dungeon]; Mistwalker Boots (10629, -1.84 DPS) [dungeon]; Mooncloth Boots (15802, -8.54 DPS, sim-verified) [crafted] |
| finger1 | Rosewine Circle (13178) | Blackrock Spire: Urok Doomhowl [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -0.23 DPS) [vendor]; The Postmaster's Seal (13392, -0.39 DPS) [dungeon]; Naglering (11669, -14.36 DPS, sim-verified) [dungeon] |
| finger2 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -0.23 DPS) [vendor]; The Postmaster's Seal (13392, -0.38 DPS) [dungeon]; Darkspear Signet (272068, -1.12 DPS) [vendor] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -2.16 DPS) [quest]; Evonice's Landin' Pilla (18951, -2.42 DPS) [quest]; Dormant Heart of the Mountain (249473, -4.95 DPS, sim-verified) [crafted] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -0.98 DPS) [quest]; Evonice's Landin' Pilla (18951, -1.24 DPS) [quest]; Darkspear Voodoo Seal (272061, -2.06 DPS) [vendor] |
| main_hand | Redemption (22406) | Stratholme: Cannon Master Willey [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Mace (23464, +0.00 DPS) [vendor]; Staff of Metanoia (22394, -3.81 DPS, sim-verified) [dungeon]; Glowing Brightwood Staff (812, -5.28 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Eternal Light (22254) | Blackrock Depths: Watchman Doomgrip [dungeon] | 27.8 healing_power points (5.06 DPS) | yes | Jaina's Firestarter (13064, -1.93 DPS) [world_drop]; Oblivion's Touch (18761, -1.95 DPS) [dungeon]; Chillnail Splinter (10704, -2.39 DPS, sim-verified) [quest] |

**New at 60:** head: Devout Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; back: Drape of Recovery; chest: Embrace of the Wind Serpent; wrist: Devout Bracers; hands: Devout Gloves; waist: Penitent's Cinch; legs: Devout Skirt; feet: Devout Sandals; finger1: Rosewine Circle; finger2: Band of Mending; trinket1: Burst of Knowledge; trinket2: Ankh of Life; main_hand: Redemption; ranged: Wand of Eternal Light

No-known-source sample (15 of 1520, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60, raid preset (undead, 005005031030000000-33505003030121431-00000000000000000)

Set DPS (verified): 893.4. Weights run: 63.4s. Verify run: 565.0s. 1520 eligible items had no known source.

Race sweep: undead 893.4 ± 0.5, troll 846.6 ± 0.5 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.206, intellect=0.820 ± 0.052, spirit=0.922 ± 0.053, mp5=1.989 ± 0.054, crit=0.469 ± 0.037 per rating point (14 rating = 1%, 6.569 per %), spell_haste=not significant (0.961 ± 0.901)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -19.26 DPS, sim-verified) [crafted] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 42.0 healing_power points (13.65 DPS) | yes | Heart of the Fiend (13960, -3.50 DPS) [dungeon]; Tooth of Gnarr (13141, -3.90 DPS) [dungeon]; Animated Chain Necklace (18723, -4.40 DPS, sim-verified) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp]; Mooncloth Shoulders (14139, -18.72 DPS, sim-verified) [crafted] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battle Healer's Cloak (19526, -2.27 DPS, sim-verified) [rep]; Shroud of the Exile (15421, -4.00 DPS) [quest] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 87.8 healing_power points (28.54 DPS) | yes | Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Embrace of the Wind Serpent (12462, -6.15 DPS) [world]; Robes of the Exalted (13346, -7.37 DPS, sim-verified) [dungeon] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wyrmthalak's Shackles (13958, +0.00 DPS) [quest]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted]; Dryad's Wrist Bindings (19595, -16.64 DPS, sim-verified) [rep] |
| hands | Hands of the Exalted Herald (12554) | Blackrock Depths: Princess Moira Bronzebeard  [dungeon] | 52.7 healing_power points (17.15 DPS) | yes | General's Satin Handwraps (231633, +0.00 DPS) [vendor]; Gilded Gloves (254095, -2.65 DPS) [crafted]; Raider Handwraps (272097, -3.68 DPS, sim-verified) [vendor] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gilded Waistcord (254081, +0.00 DPS) [crafted]; Penitent's Cinch (272394, +0.00 DPS) [vendor]; Wisdom of the Timbermaw (19047, -15.92 DPS, sim-verified) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Senior Designer's Pantaloons (11841, -21.59 DPS, sim-verified) [dungeon] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 47.5 healing_power points (15.44 DPS) | yes | General's Satin Walkers (231630, +0.00 DPS) [vendor]; Faith Healer's Boots (22247, -0.83 DPS) [dungeon]; Swarmtender's Footpads (275609, -1.45 DPS) [crafted] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -1.27 DPS) [vendor]; Rosewine Circle (13178, -2.06 DPS) [dungeon]; Naglering (11669, -19.09 DPS, sim-verified) [dungeon] |
| finger2 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -1.52 DPS, sim-verified) [vendor]; Rosewine Circle (13178, -1.95 DPS) [dungeon]; Royal Seal of Alexis (18022, -4.08 DPS) [quest] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+11.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, -3.88 DPS, sim-verified) [vendor] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Metanoia (22394, -4.71 DPS) [dungeon]; Redemption (22406, -5.25 DPS) [dungeon]; Argent Crusader (13249, -5.95 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Eternal Light (22254) | Blackrock Depths: Watchman Doomgrip [dungeon] | 17.9 healing_power points (5.82 DPS) | yes | Oblivion's Touch (18761, -1.03 DPS) [dungeon]; Brilliant Wand (249385, -1.03 DPS) [crafted]; Bonecreeper Stylus (13938, -2.77 DPS, sim-verified) [dungeon] |

**New at 60:** head: Devout Crown; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; back: Drape of Recovery; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Hands of the Exalted Herald; waist: Devout Belt; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of Mending; finger2: Fordring's Seal; trinket1: Burst of Knowledge; trinket2: Dormant Heart of the Mountain; ranged: Wand of Eternal Light

No-known-source sample (15 of 1520, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

