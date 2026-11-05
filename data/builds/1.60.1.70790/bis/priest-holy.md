# Leveling BiS: Holy

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 000000000000000000-33500000000000000-00000000000000000)

Set DPS (verified): 41.9. Weights run: 15.3s. Verify run: 80.1s. 183 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.048, intellect=0.941 ± 0.010, spirit=0.947 ± 0.011, mp5=1.827 ± 0.013, crit=0.064 ± 0.004 per rating point (14 rating = 1%, 0.893 per %), spell_haste=not significant (-0.028 ± 0.086)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.71 DPS) | yes | Flying Tiger Goggles (4368, -1.35 DPS) [crafted]; Shadow Goggles (4373, -1.48 DPS, sim-verified) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 1.9 healing_power points (0.18 DPS) | yes | Tarnished Locket (279870, -0.65 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.5 healing_power points (1.28 DPS) | yes | Slime-encrusted Pads (6461, -0.76 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.90 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -1.14 DPS, sim-verified) [crafted] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | sim-verified (41.9 DPS) | yes | Heavy Woolen Cloak (4311, -0.51 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -0.58 DPS) [crafted]; Spirit Cloak (4792, -0.68 DPS) [vendor] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 20.5 healing_power points (1.96 DPS) | yes | Robe of the Moccasin (6465, -0.20 DPS) [dungeon]; Scholarly Robes (2034, -0.59 DPS) [dungeon]; Corsair's Overshirt (5202, -0.79 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 10.2 healing_power points (0.97 DPS) | yes | Bright Bracers (3647, -0.61 DPS) [world_drop]; Seer's Cuffs (3645, -0.70 DPS) [dungeon]; Repurposed Hair Band (281256, -1.98 DPS, sim-verified) [quest] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 13.8 healing_power points (1.32 DPS) | yes | Bright Gloves (3066, -0.60 DPS) [world_drop]; Serpent Gloves (5970, -0.65 DPS) [dungeon]; Magefist Gloves (12977, -1.02 DPS, sim-verified) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 14.8 healing_power points (1.41 DPS) | yes | Novice Ardent's Sash (253887, -0.58 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.69 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.76 DPS) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 27.4 healing_power points (2.61 DPS) | yes | Colorful Kilt (10048, -1.69 DPS) [crafted]; Filigreed Silky Leggings (253939, -1.71 DPS) [crafted]; Darkweave Breeches (12987, -2.01 DPS, sim-verified) [world_drop] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 11.8 healing_power points (1.13 DPS) | yes | Spidersilk Boots (4320, -0.12 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.29 DPS) [crafted]; Kimbra Boots (6191, -0.50 DPS) [quest] |
| finger1 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 8.7 healing_power points (0.82 DPS) | yes | Band of Purification (12996, -0.28 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.29 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.56 DPS) [world_drop] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.9 healing_power points (0.66 DPS) | yes | Lavishly Jeweled Ring (1156, -0.12 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.39 DPS) [world_drop]; Band of Purification (12996, -0.48 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Westfall (2042) | The Defias Brotherhood [quest] | 58.4 healing_power points (5.56 DPS) | yes | Staff of the Blessed Seer (2271, -1.99 DPS) [dungeon]; Living Root (6631, -2.00 DPS, sim-verified) [dungeon]; Evocator's Blade (2567, -2.26 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Sizzle Stick (8071) (or Greater Magic Wand (11288), Torchlight Wand (5240)) | Deviate Eradication [quest] | 2.0 healing_power points (0.19 DPS) | yes | Torchlight Wand (5240, +0.00 DPS) [quest]; Greater Magic Wand (11288, +0.00 DPS) [crafted]; Sable Wand (7607, -0.01 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Staff of Westfall; ranged: Sizzle Stick

No-known-source sample (15 of 183, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes; 16607 Acolyte's Sacrificial Robes; 18851 Insignia of the Horde; 20434 Lorekeeper's Staff

### Band 30 (human, 000000000000000000-33505003020000000-00000000000000000)

Set DPS (verified): 67.3. Weights run: 16.7s. Verify run: 100.1s. 302 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.040, intellect=1.199 ± 0.014, spirit=0.973 ± 0.012, mp5=2.251 ± 0.017, crit=0.087 ± 0.005 per rating point (14 rating = 1%, 1.224 per %), spell_haste=not significant (0.056 ± 0.096)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 38.8 healing_power points (3.79 DPS) | yes | Filigreed Pristine Circlet (253975, -1.45 DPS, sim-verified) [crafted]; Nightsky Cowl (4039, -1.91 DPS) [world_drop]; Pristine Circlet (253949, -2.03 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 8.7 healing_power points (0.85 DPS) | yes | Crystal Starfire Medallion (5003, -0.09 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.15 DPS) [world]; Seal of the Expedition (271920, -0.26 DPS) [rep] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 19.8 healing_power points (1.93 DPS) | yes | Death Speaker Mantle (6685, -0.06 DPS) [dungeon]; Magician's Mantle (12998, -0.39 DPS) [world_drop]; Mantle of Honor (3560, -0.45 DPS) [quest] |
| back | Caretaker's Cape (19533) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Thresher Cape (6901, -0.18 DPS, sim-verified) [dungeon]; Repairman's Cape (9605, -0.24 DPS) [quest]; Expeditionary's Cape (271923, -0.43 DPS) [rep] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 33.2 healing_power points (3.24 DPS) | yes | Pristine Gown (253961, -0.09 DPS) [crafted]; Beguiler Robes (7728, -0.20 DPS) [dungeon]; Green Silk Armor (7065, -0.84 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 15.8 healing_power points (1.54 DPS) | yes | Mindthrust Bracers (1974, -0.31 DPS, sim-verified) [dungeon]; Nightsky Wristbands (6407, -0.55 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.60 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Gloves (253913, -0.30 DPS) [crafted]; Hotshot Pilot's Gloves (9491, -0.31 DPS) [dungeon]; Truefaith Gloves (7049, -0.53 DPS, sim-verified) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Highlander's Cloth Girdle (20099, -0.12 DPS) [rep]; Crimson Silk Belt (7055, -0.14 DPS) [crafted]; Belt of Arugal (6392, -0.15 DPS, sim-verified) [dungeon] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 35.3 healing_power points (3.44 DPS) | yes | Filigreed Pristine Leggings (253937, -0.75 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -0.93 DPS) [crafted]; Gaze Dreamer Pants (6903, -1.41 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 32.3 healing_power points (3.15 DPS) | yes | Acidic Walkers (9454, -1.86 DPS, sim-verified) [dungeon]; Glinteye Slippers (273024, -1.90 DPS) [dungeon]; Boots of the Enchanter (4325, -1.90 DPS) [crafted] |
| finger1 | Electrocutioner Lagnut (9447) | Gnomeregan: Electrocutioner 6000 [dungeon] | 15.8 healing_power points (1.54 DPS) | yes | Snake Hoop (6750, -0.37 DPS) [quest]; Lorekeeper's Ring (19525, -0.42 DPS) [rep]; Runemender's Seal (273031, -1.37 DPS, sim-verified) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Snake Hoop (6750, -0.07 DPS) [quest]; Lorekeeper's Ring (19525, -0.12 DPS) [rep]; Runemender's Seal (273031, -0.28 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Odo's Ley Staff (6318, -0.42 DPS) [dungeon]; Staff of Westfall (2042, -0.80 DPS) [quest]; Staff of the Friar (3415, -2.14 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 18.8 healing_power points (1.84 DPS) | yes | Orb of Mystic Insight (249394, -0.45 DPS) [crafted]; Eye of Paleth (2943, -0.57 DPS) [quest]; Orb of Mistmantle (13031, -0.58 DPS) [world_drop] |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 9.3 healing_power points (0.91 DPS) | yes | Thunderwood (13062, -0.12 DPS, sim-verified) [world_drop]; Wand of Eventide (5214, -0.42 DPS) [world_drop]; Greater Mystic Wand (217287, -0.42 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Caretaker's Cape; chest: Death Speaker Robes; wrist: Spidertank Oilrag; hands: Gloves of Old; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Ring of Power Regulation; main_hand: Death Speaker Scepter; off_hand: Orb of Souls; ranged: Wand of Mana Concentration

No-known-source sample (15 of 302, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 40 (human, 000000000000000000-33505003030121131-00000000000000000)

Set DPS (verified): 178.0. Weights run: 16.4s. Verify run: 99.1s. 389 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.946, intellect=1.733 ± 0.109, spirit=1.214 ± 0.098, mp5=4.108 ± 0.123, crit=not significant (0.315 ± 0.086) per rating point (14 rating = 1%, 4.404 per %), spell_haste=not significant (0.784 ± 1.211)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 62.4 healing_power points (8.04 DPS) | yes | Whitemane's Chapeau (7720, -1.58 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -2.85 DPS) [world_drop]; Earthen Silk Hood (254015, -3.64 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 33.9 healing_power points (4.36 DPS) | yes | Triune Amulet (7722, -2.24 DPS) [dungeon]; Amberglow Talisman (10824, -2.80 DPS) [quest]; Necklace of Calisea (1714, -3.70 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 44.5 healing_power points (5.74 DPS) | yes | Mantle of Doan (7712, -2.24 DPS) [dungeon]; Green Silken Shoulders (7057, -2.25 DPS) [crafted]; Earthen Silk Shoulders (254033, -2.68 DPS, sim-verified) [crafted] |
| back | Caretaker's Cape (19532) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackforge Cape (6424, -0.69 DPS) [dungeon]; Sergeant Major's Cape (16336, -0.82 DPS) [pvp]; Expeditionary's Cape (271922, -2.24 DPS, sim-verified) [rep] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 67.7 healing_power points (8.72 DPS) | yes | Robe of Power (7054, -2.99 DPS) [crafted]; Civinad Robes (9623, -3.50 DPS) [quest]; Dreamweave Vest (10021, -7.00 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 21.3 healing_power points (2.75 DPS) | yes | Earthen Silk Cuffs (254019, -0.68 DPS) [crafted]; Windchaser Cuffs (14429, -0.74 DPS) [world_drop]; Mindthrust Bracers (1974, -1.07 DPS, sim-verified) [dungeon] |
| hands | Gilded Handwraps (254021) | Tailoring [crafted] | 34.1 healing_power points (4.40 DPS) | yes | Stormcloth Gloves (10011, -0.30 DPS) [crafted]; Dreamweave Gloves (10019, -0.58 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -0.75 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 43.1 healing_power points (5.56 DPS) | yes | Windchaser Cinch (14435, -2.37 DPS) [world_drop]; Sutarn's Ring (13105, -2.39 DPS) [world_drop]; Deathmage Sash (10771, -3.18 DPS, sim-verified) [dungeon] |
| legs | Stoneweaver Leggings (9407) | Uldaman: Ironaya [dungeon] | 48.4 healing_power points (6.24 DPS) | yes | Pristine Leggings (253987, -1.40 DPS, sim-verified) [crafted]; Red Mageweave Pants (10009, -1.76 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.96 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 37.0 healing_power points (4.77 DPS) | yes | Thoughtcast Boots (10578, -1.06 DPS) [dungeon]; Earthen Silk Slippers (254013, -1.35 DPS) [crafted]; Furen's Boots (13100, -2.39 DPS, sim-verified) [world_drop] |
| finger1 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 21.3 healing_power points (2.75 DPS) | yes | Electrocutioner Lagnut (9447, -0.49 DPS) [dungeon]; Runemender's Seal (273031, -0.49 DPS) [world]; Philanthropist's Ring (281635, -4.78 DPS, sim-verified) [quest] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | sim-verified (178.0 DPS) | yes | Electrocutioner Lagnut (9447, -0.11 DPS) [dungeon]; Runemender's Seal (273031, -0.11 DPS) [world]; Philanthropist's Ring (281635, -0.49 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+5.2 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of the Friar (3415, -4.12 DPS) [dungeon]; Death Speaker Scepter (2816, -4.60 DPS) [dungeon]; Staff of Jordan (873, -8.60 DPS, sim-verified) [world_drop] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 29.3 healing_power points (3.78 DPS) | yes | Beacon of Hope (9393, -0.32 DPS) [dungeon]; Orb of Souls (249395, -1.17 DPS) [crafted]; Orb of Mystic Insight (249394, -1.54 DPS) [crafted] |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 15.1 healing_power points (1.95 DPS) | yes | Flash Wand (5248, -0.59 DPS) [quest]; Goblin Igniter (5253, -0.59 DPS) [quest]; Jaina's Firestarter (13064, -0.71 DPS, sim-verified) [world_drop] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Caretaker's Cape; chest: Stormcloth Vest; hands: Gilded Handwraps; waist: Gilded Cord; legs: Stoneweaver Leggings; finger1: Lorekeeper's Ring; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Hand of Righteousness; off_hand: Orb of Lorica

No-known-source sample (15 of 389, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 50 (human, 005002000000000000-33505003030121431-00000000000000000)

Set DPS (verified): 264.4. Weights run: 18.0s. Verify run: 138.6s. 496 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.371, intellect=0.676 ± 0.051, spirit=0.419 ± 0.049, mp5=1.594 ± 0.062, crit=0.223 ± 0.041 per rating point (14 rating = 1%, 3.117 per %), spell_haste=not significant (0.229 ± 0.585)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Cassandra's Grace (13102) | World drop [world_drop] | 47.8 healing_power points (16.99 DPS) | yes | Knight-Lieutenant's Satin Cover (220896, -0.72 DPS) [vendor]; Papal Fez (9431, -2.28 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -4.36 DPS) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 25.4 healing_power points (9.04 DPS) | yes | Mindburst Medallion (11196, -3.79 DPS, sim-verified) [quest]; Horizon Choker (13085, -5.08 DPS) [world_drop]; Triune Amulet (7722, -5.15 DPS) [dungeon] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | 33.9 healing_power points (12.06 DPS) | yes | Knight-Lieutenant's Satin Pads (220894, +0.00 DPS) [vendor]; Inquisitor's Shawl (19507, -1.10 DPS, sim-verified) [dungeon]; Rotgrip Mantle (17732, -3.11 DPS) [dungeon] |
| back | Caretaker's Cape (19531) | Silverwing Sentinels [rep] | sim-verified (264.4 DPS) | yes | Spritecaster Cape (11623, -0.76 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -3.59 DPS) [crafted]; Expeditionary's Cape (271921, -4.61 DPS) [rep] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 61.3 healing_power points (21.82 DPS) | yes | Knight's Satin Armor (220892, -4.35 DPS, sim-verified) [vendor]; Grimlok's Tribal Vestments (9415, -8.22 DPS) [dungeon]; Acumen Robes (17775, -10.25 DPS) [quest] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 26.8 healing_power points (9.54 DPS) | yes | Spidertank Oilrag (9448, -3.65 DPS, sim-verified) [dungeon]; Bloodband Bracers (11469, -5.15 DPS) [quest]; Earthen Silk Cuffs (254019, -5.24 DPS) [crafted] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 39.1 healing_power points (13.91 DPS) | yes | Sergeant Major's Satin Gloves (220897, -2.25 DPS, sim-verified) [vendor]; Earthenweave Gloves (254075, -3.02 DPS) [crafted]; Gilded Handwraps (254021, -4.40 DPS) [crafted] |
| waist | Gilded Waistcord (254081) | Tailoring [crafted] | 35.1 healing_power points (12.48 DPS) | yes | Gilded Cord (254037, -1.84 DPS) [crafted]; Earthenweave Cord (254077, -2.47 DPS, sim-verified) [crafted]; Dawnspire Cord (12466, -3.64 DPS) [dungeon] |
| legs | Knight's Satin Leggings (220893) | Captain Dirgehammer [vendor] | 41.9 healing_power points (14.90 DPS) | yes | Senior Designer's Pantaloons (11841, -1.94 DPS, sim-verified) [dungeon]; Stoneweaver Leggings (9407, -4.08 DPS) [dungeon]; Pristine Leggings (253987, -4.65 DPS) [crafted] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 37.1 healing_power points (13.19 DPS) | yes | Earthenweave Boots (254093, -2.83 DPS, sim-verified) [crafted]; Sergeant Major's Satin Boots (220895, -3.19 DPS) [vendor]; Gilded Slippers (254001, -3.80 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 18.4 healing_power points (6.54 DPS) | yes | Band of the Unicorn (7553, -1.91 DPS) [world_drop]; Electrocutioner Lagnut (9447, -2.29 DPS) [dungeon]; Runemender's Seal (273031, -2.29 DPS) [world] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.4 healing_power points (4.76 DPS) | yes | Electrocutioner Lagnut (9447, -0.52 DPS) [dungeon]; Runemender's Seal (273031, -0.52 DPS) [world]; Band of the Unicorn (7553, -1.08 DPS, sim-verified) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (264.4 DPS) | yes | Enriched Seal (271903, +0.00 DPS) [rep]; Evonice's Landin' Pilla (18951, -0.30 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (264.4 DPS) | yes | Enriched Seal (271903, -3.22 DPS, sim-verified) [rep]; Evonice's Landin' Pilla (18951, -4.91 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (264.4 DPS) | yes | Resurgence Rod (17743, -15.57 DPS, sim-verified) [quest]; Charstone Dirk (17710, -16.62 DPS) [dungeon]; Dreamstaff (249454, -21.49 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Earthen Rod (9381) | Uldaman [dungeon] | 9.0 healing_power points (3.20 DPS) | yes | Lesser Eternal Wand (249232, -0.36 DPS) [crafted]; Charged Lightning Rod (11860, -0.46 DPS) [quest]; Captain Rackmore's Tiller (16789, -0.62 DPS) [quest] |

**New at 50:** head: Cassandra's Grace; shoulder: Nethergeld Shoulders; back: Caretaker's Cape; wrist: Nethergeld Cuffs; hands: Gilded Gloves; waist: Gilded Waistcord; legs: Knight's Satin Leggings; feet: Gilded Sandals; finger1: Lorekeeper's Ring; finger2: Philanthropist's Ring; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Earthen Rod

No-known-source sample (15 of 496, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60 (human, 005005031030000000-33505003030121431-00000000000000000)

Set DPS (verified): 431.8. Weights run: 16.1s. Verify run: 331.0s. 1107 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.106, intellect=0.734 ± 0.054, spirit=1.946 ± 0.052, mp5=3.914 ± 0.054, crit=0.230 ± 0.016 per rating point (14 rating = 1%, 3.225 per %), spell_haste=not significant (0.307 ± 0.542)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mooncloth Circlet (14140) | Tailoring [crafted] | 80.2 healing_power points (19.44 DPS) | yes | Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Earthenweave Crown (254137, -3.56 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (431.8 DPS) | yes | Stormpike Sage's Pendant (19098, -0.70 DPS) [rep]; Triune Amulet (7722, -1.03 DPS) [dungeon]; Lady Alizabeth's Pendant (13002, -7.88 DPS, sim-verified) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (431.8 DPS) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp]; Mooncloth Shoulders (14139, -10.20 DPS, sim-verified) [crafted] |
| back | Caretaker's Cape (19530) | Silverwing Sentinels [rep] | sim-verified (431.8 DPS) | yes | Spritecaster Cape (11623, -2.78 DPS) [dungeon]; Wingveil Cloak (10802, -3.67 DPS) [dungeon]; Ethereal Mist Cape (11873, -3.88 DPS) [quest] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 104.1 healing_power points (25.24 DPS) | yes | Field Marshal's Satin Tunic (231624, -0.99 DPS) [vendor]; Knight-Captain's Satin Tunic (227122, -6.27 DPS) [pvp]; Ghostweave Vest (14141, -20.56 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (431.8 DPS) | yes | Earthenweave Cuffs (254131, +0.00 DPS) [crafted]; Dryad's Wrist Bindings (19596, -0.06 DPS) [rep] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | 55.4 healing_power points (13.44 DPS) | yes | Marshal's Satin Handwraps (231623, -0.46 DPS) [vendor]; Gilded Gloves (254095, -0.85 DPS) [crafted]; Swarmtender's Gloves (275607, -15.63 DPS, sim-verified) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 58.8 healing_power points (14.25 DPS) | yes | Marshal's Satin Sash (17609, -3.08 DPS) [pvp]; Devout Belt (16696, -3.78 DPS) [dungeon]; Earthenweave Cord (254077, -5.42 DPS, sim-verified) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 79.8 healing_power points (19.34 DPS) | yes | Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Knight-Captain's Satin Legguards (227125, -0.37 DPS) [pvp]; Mooncloth Leggings (14137, -17.07 DPS, sim-verified) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (431.8 DPS) | yes | Marshal's Satin Walkers (231627, +0.00 DPS) [vendor]; Earthenweave Boots (254093, -0.97 DPS) [crafted]; Mooncloth Boots (15802, -9.07 DPS, sim-verified) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 34.8 healing_power points (8.44 DPS) | yes | Band of the Hierophant (13096, -1.47 DPS) [world_drop]; Eye of Adaegus (5266, -1.71 DPS) [world_drop]; Choking Band (11868, -2.31 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19522) | Silverwing Sentinels [rep] | 29.7 healing_power points (7.19 DPS) | yes | Eye of Adaegus (5266, -0.46 DPS) [world_drop]; Choking Band (11868, -1.06 DPS) [quest]; Band of the Hierophant (13096, -1.96 DPS, sim-verified) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (431.8 DPS) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Enriched Seal (271903, -0.92 DPS) [rep]; Evonice's Landin' Pilla (18951, -0.94 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (431.8 DPS) | yes | Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Enriched Seal (271903, +0.00 DPS) [rep]; Shard of the Splithooves (10659, -6.51 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (431.8 DPS) | yes | Grand Marshal's Warhammer (23454, +0.00 DPS) [vendor]; Resurgence Rod (17743, -8.18 DPS) [quest]; Hand of Edward the Odd (2243, -13.40 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 38.4 healing_power points (9.31 DPS) | yes | Brilliant Wand (249385, -5.18 DPS, sim-verified) [crafted]; Cairnstone Sliver (9654, -6.06 DPS) [quest]; Captain Rackmore's Tiller (16789, -6.44 DPS) [quest] |

**New at 60:** head: Mooncloth Circlet; shoulder: Devout Mantle; back: Caretaker's Cape; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Devout Sandals; finger1: Maiden's Circle; finger2: Lorekeeper's Ring; ranged: Torch of Light

No-known-source sample (15 of 1107, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60, raid preset (human, 005005031030000000-33505003030121431-00000000000000000)

Set DPS (verified): 847.3. Weights run: 5.1s. Verify run: 82.1s. 1107 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.252, intellect=0.631 ± 0.042, spirit=1.100 ± 0.042, mp5=2.174 ± 0.045, crit=0.452 ± 0.043 per rating point (14 rating = 1%, 6.327 per %), spell_haste=not significant (0.973 ± 1.032)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mooncloth Circlet (14140) | Tailoring [crafted] | 67.5 healing_power points (37.02 DPS) | yes | Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Lieutenant Commander's Satin Hood (227121, -0.41 DPS) [pvp]; Earthenweave Crown (254137, -5.18 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 25.0 healing_power points (13.74 DPS) | yes | Archlight Talisman (15856, -2.91 DPS, sim-verified) [quest]; Lady Alizabeth's Pendant (13002, -4.69 DPS) [world_drop]; Triune Amulet (7722, -5.13 DPS) [dungeon] |
| shoulder | Mooncloth Shoulders (14139) | Tailoring [crafted] | 49.1 healing_power points (26.93 DPS) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp]; Earthenweave Mantle (254119, -5.28 DPS, sim-verified) [crafted] |
| back | Caretaker's Cape (19530) | Silverwing Sentinels [rep] | sim-verified (+3.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Spritecaster Cape (11623, -5.71 DPS) [dungeon]; Wingveil Cloak (10802, -9.71 DPS) [dungeon]; Expeditionary's Cape (271907, -9.93 DPS) [rep] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 90.6 healing_power points (49.69 DPS) | yes | Field Marshal's Satin Tunic (231624, -0.57 DPS) [vendor]; Stormcloth Vest (10020, -11.12 DPS, sim-verified) [crafted]; Knight-Captain's Satin Tunic (227122, -14.12 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 34.8 healing_power points (19.06 DPS) | yes | Earthenweave Cuffs (254131, -1.76 DPS, sim-verified) [crafted]; Nethergeld Cuffs (254061, -2.65 DPS) [crafted]; Devout Bracers (16697, -4.22 DPS) [dungeon] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 44.1 healing_power points (24.19 DPS) | yes | Marshal's Satin Handwraps (231623, +0.00 DPS) [vendor]; Knight-Lieutenant's Satin Handwraps (227128, -0.53 DPS) [pvp]; Theramore Handwraps (271945, -1.69 DPS) [rep] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 49.7 healing_power points (27.28 DPS) | yes | Earthenweave Cord (254077, -6.10 DPS, sim-verified) [crafted]; Gilded Waistcord (254081, -8.26 DPS) [crafted]; Gilded Cord (254037, -8.82 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 58.8 healing_power points (32.23 DPS) | yes | Knight-Captain's Satin Legguards (227125, +0.00 DPS) [pvp]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Earthenweave Leggings (254125, -2.30 DPS, sim-verified) [crafted] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 47.1 healing_power points (25.85 DPS) | yes | Marshal's Satin Walkers (231627, +0.00 DPS) [vendor]; Swarmtender's Footpads (275609, -2.27 DPS) [crafted]; Earthenweave Boots (254093, -4.96 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 29.0 healing_power points (15.92 DPS) | yes | Band of the Hierophant (13096, -5.82 DPS) [world_drop]; Deep Rooted Ring (19109, -5.87 DPS) [world]; Eye of Adaegus (5266, -6.60 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (19522) | Silverwing Sentinels [rep] | 22.7 healing_power points (12.45 DPS) | yes | Band of the Hierophant (13096, -2.35 DPS) [world_drop]; Deep Rooted Ring (19109, -2.41 DPS) [world]; Eye of Adaegus (5266, -3.13 DPS) [world_drop] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (847.4 DPS) | yes | Ankh of Life (1713, -3.46 DPS, sim-verified) [world_drop]; Evonice's Landin' Pilla (18951, -3.84 DPS) [quest]; Enriched Seal (271903, -3.91 DPS) [rep] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (-4.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.12 DPS) [quest]; Enriched Seal (271903, -1.19 DPS) [rep] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Warhammer (23454, +0.00 DPS) [vendor]; Hand of Edward the Odd (2243, -12.55 DPS, sim-verified) [world_drop]; Resurgence Rod (17743, -23.19 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 29.1 healing_power points (15.96 DPS) | yes | Brilliant Wand (249385, -6.94 DPS, sim-verified) [crafted]; Captain Rackmore's Tiller (16789, -10.86 DPS) [quest]; Greater Eternal Wand (249237, -11.03 DPS) [crafted] |

**New at 60:** head: Mooncloth Circlet; shoulder: Mooncloth Shoulders; back: Caretaker's Cape; chest: Truefaith Vestments; wrist: Dryad's Wrist Bindings; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Maiden's Circle; finger2: Lorekeeper's Ring; trinket1: Uther's Strength; trinket2: Shard of the Splithooves; ranged: Torch of Light

No-known-source sample (15 of 1107, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

## Horde

### Band 20 (undead, 000000000000000000-33500000000000000-00000000000000000)

Set DPS (verified): 40.3. Weights run: 15.3s. Verify run: 81.8s. 176 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.107, intellect=1.074 ± 0.017, spirit=0.981 ± 0.017, mp5=2.493 ± 0.019, crit=0.110 ± 0.007 per rating point (14 rating = 1%, 1.534 per %), spell_haste=not significant (-0.017 ± 0.150)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (2.00 DPS) | yes | Shadow Goggles (4373, -1.49 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.56 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | sim-verified (40.3 DPS) | yes | Roadwatcher's Confidence (281265, -0.18 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 14.7 healing_power points (1.63 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.68 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.80 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -1.18 DPS) [crafted] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.36 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -0.64 DPS) [crafted]; Sanguine Cape (14376, -0.74 DPS) [world_drop] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 21.3 healing_power points (2.37 DPS) | yes | Robe of the Moccasin (6465, -0.28 DPS) [dungeon]; Scholarly Robes (2034, -0.73 DPS) [dungeon]; Corsair's Overshirt (5202, -0.73 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 12.9 healing_power points (1.43 DPS) | yes | Featherbead Bracers (15452, -1.07 DPS) [quest]; Owlbeard Bracers (16981, -1.08 DPS) [quest]; Bright Bracers (3647, -1.36 DPS, sim-verified) [world_drop] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 14.2 healing_power points (1.58 DPS) | yes | Bright Gloves (3066, -0.67 DPS) [world_drop]; Magefist Gloves (12977, -0.67 DPS, sim-verified) [world_drop]; Serpent Gloves (5970, -0.80 DPS) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 15.3 healing_power points (1.70 DPS) | yes | Novice Ardent's Sash (253887, -0.45 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.74 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.90 DPS) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 28.4 healing_power points (3.15 DPS) | yes | Darkweave Breeches (12987, -1.51 DPS, sim-verified) [world_drop]; Filigreed Silky Leggings (253939, -2.00 DPS) [crafted]; Filigreed Flame Leggings (253941, -2.00 DPS) [crafted] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 12.2 healing_power points (1.36 DPS) | yes | Spidersilk Boots (4320, -0.10 DPS) [crafted]; Red Woolen Boots (4313, -0.37 DPS) [crafted]; Woolen Boots (2583, -0.68 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 10.0 healing_power points (1.11 DPS) | yes | Band of Purification (12996, -0.46 DPS) [world_drop]; Loop of Sacrifice (281673, -0.51 DPS) [quest]; Volcanic Rock Ring (12053, -0.75 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 6.4 healing_power points (0.72 DPS) | yes | Loop of Sacrifice (281673, -0.12 DPS) [quest]; Band of Purification (12996, -0.20 DPS, sim-verified) [world_drop]; Volcanic Rock Ring (12053, -0.36 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Living Root (6631) | Wailing Caverns: Verdan the Everliving [dungeon] | 49.9 healing_power points (5.54 DPS) | yes | Staff of the Blessed Seer (2271, -0.91 DPS, sim-verified) [dungeon]; Gnarled Necromancer's Staff (251534, -1.56 DPS) [quest]; Evocator's Blade (2567, -1.65 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Decay (5252) (or Flaring Baton (5326)) | Beren's Peril [quest] | 2.1 healing_power points (0.24 DPS) | yes | Flaring Baton (5326, +0.00 DPS) [quest]; Sizzle Stick (8071, -0.02 DPS) [quest]; Greater Magic Wand (11288, -0.02 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Living Root; ranged: Wand of Decay

No-known-source sample (15 of 176, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20441 Scout's Blade; 23173 Abomination Skin Leggings; 209613 Insignia of the Alliance; 241089 Scarlet Dagger; 248007 Militia Shortblade

### Band 30 (undead, 000000000000000000-33505003020000000-00000000000000000)

Set DPS (verified): 73.0. Weights run: 16.7s. Verify run: 97.4s. 295 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.143, intellect=1.279 ± 0.021, spirit=0.429 ± 0.027, mp5=2.269 ± 0.031, crit=0.163 ± 0.013 per rating point (14 rating = 1%, 2.283 per %), spell_haste=not significant (0.245 ± 0.208)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 35.6 healing_power points (4.81 DPS) | yes | Filigreed Pristine Circlet (253975, -1.76 DPS, sim-verified) [crafted]; Enchanter's Cowl (4322, -2.27 DPS) [crafted]; Pristine Circlet (253949, -2.38 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 7.7 healing_power points (1.04 DPS) | yes | Crystal Starfire Medallion (5003, -0.17 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.17 DPS) [vendor]; Kaleidoscope Chain (13084, -0.71 DPS, sim-verified) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 20.5 healing_power points (2.77 DPS) | yes | Death Speaker Mantle (6685, -0.21 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.54 DPS) [world_drop]; Feline Mantle (3748, -0.85 DPS) [dungeon] |
| back | Battle Healer's Cloak (19529) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Thresher Cape (6901, -0.58 DPS) [dungeon]; Darkspear Raider's Cloak (272078, -0.62 DPS, sim-verified) [vendor]; Resilient Cape (14400, -0.95 DPS) [world_drop] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 34.1 healing_power points (4.61 DPS) | yes | Pristine Gown (253961, -0.73 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -0.85 DPS) [dungeon]; Green Silk Armor (7065, -1.14 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 15.8 healing_power points (2.14 DPS) | yes | Mindthrust Bracers (1974, -0.34 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -0.75 DPS) [world_drop]; Nightsky Wristbands (6407, -0.93 DPS) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 18.8 healing_power points (2.55 DPS) | yes | Pristine Gloves (253913, -0.52 DPS, sim-verified) [crafted]; Jutebraid Gloves (10654, -0.87 DPS) [quest]; Hotshot Pilot's Gloves (9491, -0.87 DPS) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.1 healing_power points (2.18 DPS) | yes | Crimson Silk Belt (7055, -0.16 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.17 DPS) [rep]; Belt of Arugal (6392, -1.03 DPS, sim-verified) [dungeon] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 33.1 healing_power points (4.48 DPS) | yes | Filigreed Pristine Leggings (253937, -0.89 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -1.81 DPS) [crafted]; Leech Pants (6910, -2.27 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 30.7 healing_power points (4.15 DPS) | yes | Acidic Walkers (9454, -1.90 DPS, sim-verified) [dungeon]; Pristine Boots (253889, -2.41 DPS) [crafted]; Spidersilk Boots (4320, -2.51 DPS) [crafted] |
| finger1 | Solid Gold Band (279845) | Opportunistic Education [quest] | 13.1 healing_power points (1.78 DPS) | yes | Snake Hoop (6750, -0.10 DPS) [quest]; Electrocutioner Lagnut (9447, -0.15 DPS) [dungeon]; Runemender's Seal (273031, -0.15 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 12.8 healing_power points (1.73 DPS) | yes | Electrocutioner Lagnut (9447, -0.11 DPS) [dungeon]; Runemender's Seal (273031, -0.11 DPS) [world]; Snake Hoop (6750, -1.44 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Odo's Ley Staff (6318, -1.47 DPS) [dungeon]; Glimmering Staff (249392, -2.70 DPS) [crafted]; Staff of the Friar (3415, -3.74 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | sim-verified (73.0 DPS) | yes | Orb of Souls (249395, -0.36 DPS, sim-verified) [crafted]; Orb of Mistmantle (13031, -0.54 DPS) [world_drop]; Witch's Finger (16887, -0.60 DPS) [quest] |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 9.7 healing_power points (1.31 DPS) | yes | Summoner's Wand (5245, -0.61 DPS) [world]; Lesser Mystic Wand (11289, -0.61 DPS) [crafted]; Thunderwood (13062, -0.70 DPS, sim-verified) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Battle Healer's Cloak; chest: Death Speaker Robes; wrist: Spidertank Oilrag; hands: Truefaith Gloves; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Solid Gold Band; finger2: Ring of Power Regulation; main_hand: Death Speaker Scepter; off_hand: Orb of Mystic Insight; ranged: Wand of Mana Concentration

No-known-source sample (15 of 295, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (undead, 000000000000000000-33505003030121131-00000000000000000)

Set DPS (verified): 188.4. Weights run: 16.4s. Verify run: 97.4s. 382 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.946, intellect=1.733 ± 0.109, spirit=1.214 ± 0.098, mp5=4.108 ± 0.123, crit=not significant (0.315 ± 0.086) per rating point (14 rating = 1%, 4.404 per %), spell_haste=not significant (0.784 ± 1.211)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 62.4 healing_power points (8.04 DPS) | yes | Whitemane's Chapeau (7720, -1.79 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -2.85 DPS) [world_drop]; Earthen Silk Hood (254015, -3.64 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 33.9 healing_power points (4.36 DPS) | yes | Triune Amulet (7722, -2.24 DPS) [dungeon]; Amberglow Talisman (10824, -2.80 DPS) [quest]; Necklace of Calisea (1714, -3.27 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 44.5 healing_power points (5.74 DPS) | yes | Mantle of Doan (7712, -2.24 DPS) [dungeon]; Green Silken Shoulders (7057, -2.25 DPS) [crafted]; Earthen Silk Shoulders (254033, -2.88 DPS, sim-verified) [crafted] |
| back | Battle Healer's Cloak (19528) | Warsong Outriders [rep] | sim-verified (188.5 DPS) | yes | Blackforge Cape (6424, -0.69 DPS) [dungeon]; Darkspear Raider's Cloak (272078, -0.85 DPS) [vendor] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 67.7 healing_power points (8.72 DPS) | yes | Robe of Power (7054, -2.99 DPS) [crafted]; Civinad Robes (9623, -3.50 DPS) [quest]; Dreamweave Vest (10021, -6.17 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 21.3 healing_power points (2.75 DPS) | yes | Radiant Silver Bracers (4545, -0.45 DPS) [quest]; Earthen Silk Cuffs (254019, -0.68 DPS) [crafted]; Mindthrust Bracers (1974, -1.18 DPS, sim-verified) [dungeon] |
| hands | Gilded Handwraps (254021) | Tailoring [crafted] | 34.1 healing_power points (4.40 DPS) | yes | Stormcloth Gloves (10011, -0.30 DPS) [crafted]; Dreamweave Gloves (10019, -0.50 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -0.75 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 43.1 healing_power points (5.56 DPS) | yes | Windchaser Cinch (14435, -2.37 DPS) [world_drop]; Sutarn's Ring (13105, -2.39 DPS) [world_drop]; Deathmage Sash (10771, -2.99 DPS, sim-verified) [dungeon] |
| legs | Stoneweaver Leggings (9407) | Uldaman: Ironaya [dungeon] | 48.4 healing_power points (6.24 DPS) | yes | Pristine Leggings (253987, -1.21 DPS, sim-verified) [crafted]; Red Mageweave Pants (10009, -1.76 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.96 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 37.0 healing_power points (4.77 DPS) | yes | Thoughtcast Boots (10578, -1.06 DPS) [dungeon]; Earthen Silk Slippers (254013, -1.35 DPS) [crafted]; Furen's Boots (13100, -2.81 DPS, sim-verified) [world_drop] |
| finger1 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 24.6 healing_power points (3.18 DPS) | yes | Philanthropist's Ring (281635, -0.77 DPS) [quest]; Ring of Power Regulation (273030, -0.81 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.92 DPS) [dungeon] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 21.3 healing_power points (2.75 DPS) | yes | Ring of Power Regulation (273030, -0.39 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.49 DPS) [dungeon]; Philanthropist's Ring (281635, -0.87 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (188.5 DPS) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (188.5 DPS) | yes | Staff of the Friar (3415, -4.12 DPS) [dungeon]; Death Speaker Scepter (2816, -4.60 DPS) [dungeon]; Staff of Jordan (873, -9.08 DPS, sim-verified) [world_drop] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | 26.9 healing_power points (3.46 DPS) | yes | Orb of Souls (249395, -0.85 DPS) [crafted]; Orb of Mystic Insight (249394, -1.22 DPS) [crafted]; Omega Orb (7749, -1.74 DPS, sim-verified) [quest] |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 15.1 healing_power points (1.95 DPS) | yes | Jaina's Firestarter (13064, -0.47 DPS, sim-verified) [world_drop]; Dancing Flame (6806, -0.54 DPS) [quest]; Goblin Igniter (5253, -0.59 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Battle Healer's Cloak; chest: Stormcloth Vest; hands: Gilded Handwraps; waist: Gilded Cord; legs: Stoneweaver Leggings; finger1: Darkspear Signet; finger2: Advisor's Ring; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Hand of Righteousness; off_hand: Beacon of Hope

No-known-source sample (15 of 382, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16336 Sergeant Major's Cape

### Band 50 (undead, 005002000000000000-33505003030121431-00000000000000000)

Set DPS (verified): 282.1. Weights run: 18.0s. Verify run: 139.4s. 489 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.371, intellect=0.676 ± 0.051, spirit=0.419 ± 0.049, mp5=1.594 ± 0.062, crit=0.223 ± 0.041 per rating point (14 rating = 1%, 3.117 per %), spell_haste=not significant (0.229 ± 0.585)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Cassandra's Grace (13102) | World drop [world_drop] | 47.8 healing_power points (16.99 DPS) | yes | Blood Guard's Satin Cover (220899, -0.72 DPS) [vendor]; Papal Fez (9431, -1.69 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -4.36 DPS) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 25.4 healing_power points (9.04 DPS) | yes | Mindburst Medallion (11196, -4.21 DPS, sim-verified) [quest]; Horizon Choker (13085, -5.08 DPS) [world_drop]; Triune Amulet (7722, -5.15 DPS) [dungeon] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | 33.9 healing_power points (12.06 DPS) | yes | Blood Guard's Satin Pads (220901, +0.00 DPS) [vendor]; Inquisitor's Shawl (19507, -0.94 DPS, sim-verified) [dungeon]; Rotgrip Mantle (17732, -3.11 DPS) [dungeon] |
| back | Battle Healer's Cloak (19527) | Warsong Outriders [rep] | sim-verified (282.2 DPS) | yes | Spritecaster Cape (11623, -1.62 DPS, sim-verified) [dungeon]; Deep Woodlands Cloak (19121, -2.29 DPS) [quest]; Runecloth Cloak (13860, -3.59 DPS) [crafted] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 61.3 healing_power points (21.82 DPS) | yes | Stone Guard's Satin Armor (220903, -5.21 DPS, sim-verified) [vendor]; Grimlok's Tribal Vestments (9415, -8.22 DPS) [dungeon]; Acumen Robes (17775, -10.25 DPS) [quest] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 26.8 healing_power points (9.54 DPS) | yes | Bloodband Bracers (11469, -5.15 DPS) [quest]; Spidertank Oilrag (9448, -5.20 DPS, sim-verified) [dungeon]; Earthen Silk Cuffs (254019, -5.24 DPS) [crafted] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 39.1 healing_power points (13.91 DPS) | yes | First Sergeant's Satin Gloves (220898, -2.36 DPS, sim-verified) [vendor]; Earthenweave Gloves (254075, -3.02 DPS) [crafted]; Greenleaf Handwraps (19116, -3.44 DPS) [quest] |
| waist | Gilded Waistcord (254081) | Tailoring [crafted] | 35.1 healing_power points (12.48 DPS) | yes | Gilded Cord (254037, -1.84 DPS) [crafted]; Earthenweave Cord (254077, -3.12 DPS, sim-verified) [crafted]; Dawnspire Cord (12466, -3.64 DPS) [dungeon] |
| legs | Stone Guard's Satin Leggings (220902) | Lady Palanseer [vendor] | 41.9 healing_power points (14.90 DPS) | yes | Senior Designer's Pantaloons (11841, -2.35 DPS, sim-verified) [dungeon]; Stoneweaver Leggings (9407, -4.08 DPS) [dungeon]; Pristine Leggings (253987, -4.65 DPS) [crafted] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 37.1 healing_power points (13.19 DPS) | yes | First Sergeant's Satin Boots (220900, -3.19 DPS) [vendor]; Earthenweave Boots (254093, -3.30 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -3.80 DPS) [crafted] |
| finger1 | Advisor's Ring (19519) | Warsong Outriders [rep] | 18.4 healing_power points (6.54 DPS) | yes | Solid Gold Band (279845, -1.88 DPS) [quest]; Band of the Unicorn (7553, -1.91 DPS) [world_drop]; Electrocutioner Lagnut (9447, -2.29 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.4 healing_power points (4.76 DPS) | yes | Band of the Unicorn (7553, -0.13 DPS) [world_drop]; Electrocutioner Lagnut (9447, -0.52 DPS) [dungeon]; Solid Gold Band (279845, -0.75 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (282.2 DPS) | yes | Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor]; Evonice's Landin' Pilla (18951, -0.30 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (282.2 DPS) | yes | Darkspear Voodoo Seal (272061, -3.47 DPS, sim-verified) [vendor]; Evonice's Landin' Pilla (18951, -4.91 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (282.2 DPS) | yes | Resurgence Rod (17743, -16.50 DPS, sim-verified) [quest]; Charstone Dirk (17710, -16.62 DPS) [dungeon]; Dreamstaff (249454, -21.49 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Earthen Rod (9381) | Uldaman [dungeon] | 9.0 healing_power points (3.20 DPS) | yes | Charged Lightning Rod (11860, -0.46 DPS) [quest]; Captain Rackmore's Tiller (16789, -0.62 DPS) [quest]; Lesser Eternal Wand (249232, -0.63 DPS, sim-verified) [crafted] |

**New at 50:** head: Cassandra's Grace; shoulder: Nethergeld Shoulders; back: Battle Healer's Cloak; wrist: Nethergeld Cuffs; hands: Gilded Gloves; waist: Gilded Waistcord; legs: Stone Guard's Satin Leggings; feet: Gilded Sandals; finger1: Advisor's Ring; finger2: Philanthropist's Ring; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Earthen Rod

No-known-source sample (15 of 489, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60 (undead, 005005031030000000-33505003030121431-00000000000000000)

Set DPS (verified): 442.1. Weights run: 16.1s. Verify run: 341.9s. 1103 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.106, intellect=0.734 ± 0.054, spirit=1.946 ± 0.052, mp5=3.914 ± 0.054, crit=0.230 ± 0.016 per rating point (14 rating = 1%, 3.225 per %), spell_haste=not significant (0.307 ± 0.542)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (442.1 DPS) | yes | Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -8.62 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (442.1 DPS) | yes | Frostwolf Advisor's Pendant (19096, -0.70 DPS) [rep]; Triune Amulet (7722, -1.03 DPS) [dungeon]; Lady Alizabeth's Pendant (13002, -3.55 DPS, sim-verified) [world_drop] |
| shoulder | Mooncloth Shoulders (14139) | Tailoring [crafted] | 58.4 healing_power points (14.16 DPS) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp]; Earthenweave Mantle (254119, -2.86 DPS, sim-verified) [crafted] |
| back | Battle Healer's Cloak (19526) | Warsong Outriders [rep] | sim-verified (442.1 DPS) | yes | Spritecaster Cape (11623, -2.78 DPS) [dungeon]; Wingveil Cloak (10802, -3.67 DPS) [dungeon]; Cloak of Blight (6832, -3.94 DPS) [quest] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 104.1 healing_power points (25.24 DPS) | yes | Warlord's Satin Tunic (231632, -0.99 DPS) [vendor]; Legionnaire's Satin Tunic (227124, -6.27 DPS) [pvp]; Ghostweave Vest (14141, -17.94 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (442.1 DPS) | yes | Earthenweave Cuffs (254131, +0.00 DPS) [crafted]; Dryad's Wrist Bindings (19596, -0.06 DPS) [rep] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | 55.4 healing_power points (13.44 DPS) | yes | General's Satin Handwraps (231633, -0.46 DPS) [vendor]; Gilded Gloves (254095, -0.85 DPS) [crafted]; Swarmtender's Gloves (275607, -12.67 DPS, sim-verified) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 58.8 healing_power points (14.25 DPS) | yes | General's Satin Cinch (17621, -3.08 DPS) [pvp]; Devout Belt (16696, -3.78 DPS) [dungeon]; Earthenweave Cord (254077, -5.06 DPS, sim-verified) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 79.8 healing_power points (19.34 DPS) | yes | General's Satin Legguards (231634, +0.00 DPS) [vendor]; Legionnaire's Satin Legguards (227123, -0.37 DPS) [pvp]; Mooncloth Leggings (14137, -14.06 DPS, sim-verified) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (442.1 DPS) | yes | General's Satin Walkers (231630, +0.00 DPS) [vendor]; Earthenweave Boots (254093, -0.97 DPS) [crafted]; Mooncloth Boots (15802, -6.54 DPS, sim-verified) [crafted] |
| finger1 | Darkspear Signet (272068) | Creeg Bothunk [vendor] | 35.2 healing_power points (8.54 DPS) | yes | Advisor's Ring (19518, -1.35 DPS) [rep]; Band of the Hierophant (13096, -1.57 DPS) [world_drop]; Eye of Adaegus (5266, -1.81 DPS) [world_drop] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | 34.8 healing_power points (8.44 DPS) | yes | Band of the Hierophant (13096, -1.47 DPS) [world_drop]; Eye of Adaegus (5266, -1.71 DPS) [world_drop]; Advisor's Ring (19518, -1.77 DPS, sim-verified) [rep] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (442.1 DPS) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Darkspear Voodoo Seal (272061, -0.92 DPS) [vendor]; Evonice's Landin' Pilla (18951, -0.94 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (442.1 DPS) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor]; Evonice's Landin' Pilla (18951, -2.76 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (442.1 DPS) | yes | High Warlord's Battle Mace (23464, +0.00 DPS) [vendor]; Resurgence Rod (17743, -8.18 DPS) [quest]; Hand of Edward the Odd (2243, -8.49 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 38.4 healing_power points (9.31 DPS) | yes | Brilliant Wand (249385, -3.22 DPS, sim-verified) [crafted]; Captain Rackmore's Tiller (16789, -6.44 DPS) [quest]; Chillnail Splinter (10704, -6.48 DPS) [quest] |

**New at 60:** head: Devout Crown; shoulder: Mooncloth Shoulders; back: Battle Healer's Cloak; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Devout Sandals; finger1: Darkspear Signet; finger2: Maiden's Circle; ranged: Torch of Light

No-known-source sample (15 of 1103, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60, raid preset (undead, 005005031030000000-33505003030121431-00000000000000000)

Set DPS (verified): 857.9. Weights run: 5.1s. Verify run: 82.9s. 1103 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.070, intellect=0.574 ± 0.018, spirit=0.817 ± 0.019, mp5=1.495 ± 0.019, crit=0.270 ± 0.014 per rating point (14 rating = 1%, 3.787 per %), spell_haste=not significant (0.508 ± 0.499)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mooncloth Circlet (14140) | Tailoring [crafted] | 63.3 healing_power points (30.65 DPS) | yes | Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Champion's Satin Hood (227118, -0.72 DPS) [pvp]; Earthenweave Crown (254137, -4.32 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 24.6 healing_power points (11.91 DPS) | yes | Archlight Talisman (15856, -3.16 DPS, sim-verified) [quest]; Triune Amulet (7722, -5.27 DPS) [dungeon]; Mindburst Medallion (11196, -5.36 DPS) [quest] |
| shoulder | Mooncloth Shoulders (14139) | Tailoring [crafted] | 46.0 healing_power points (22.28 DPS) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp]; Earthenweave Mantle (254119, -4.97 DPS, sim-verified) [crafted] |
| back | Battle Healer's Cloak (19526) | Warsong Outriders [rep] | sim-verified (+2.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Spritecaster Cape (11623, -4.94 DPS) [dungeon]; Deep Woodlands Cloak (19121, -7.45 DPS) [quest]; Darkspear Raider's Cloak (272063, -8.94 DPS) [vendor] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 86.1 healing_power points (41.70 DPS) | yes | Warlord's Satin Tunic (231632, -0.32 DPS) [vendor]; Legionnaire's Satin Tunic (227124, -12.61 DPS) [pvp]; Stormcloth Vest (10020, -12.71 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 32.3 healing_power points (15.65 DPS) | yes | Nethergeld Cuffs (254061, -2.04 DPS) [crafted]; Earthenweave Cuffs (254131, -3.47 DPS, sim-verified) [crafted]; Devout Bracers (16697, -4.14 DPS) [dungeon] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 41.3 healing_power points (20.00 DPS) | yes | Blood Guard's Satin Handwraps (227126, +0.00 DPS) [pvp]; General's Satin Handwraps (231633, +0.00 DPS) [vendor]; Raider Handwraps (272097, -0.87 DPS) [vendor] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 46.4 healing_power points (22.47 DPS) | yes | Gilded Waistcord (254081, -5.92 DPS) [crafted]; Gilded Cord (254037, -7.22 DPS) [crafted]; Earthenweave Cord (254077, -7.81 DPS, sim-verified) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 51.4 healing_power points (24.90 DPS) | yes | Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Earthenweave Leggings (254125, -5.32 DPS, sim-verified) [crafted] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Satin Walkers (231630, +0.00 DPS) [vendor]; Swarmtender's Footpads (275609, -1.47 DPS, sim-verified) [crafted]; Gilded Sandals (254107, -3.15 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 26.9 healing_power points (13.04 DPS) | yes | Deep Rooted Ring (19109, -4.45 DPS) [world]; Solid Gold Band (279845, -5.73 DPS) [quest]; Band of the Hierophant (13096, -5.91 DPS) [world_drop] |
| finger2 | Advisor's Ring (19518) | Warsong Outriders [rep] | 20.0 healing_power points (9.68 DPS) | yes | Deep Rooted Ring (19109, -1.09 DPS) [world]; Solid Gold Band (279845, -2.37 DPS) [quest]; Band of the Hierophant (13096, -2.55 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -0.41 DPS) [quest]; Evonice's Landin' Pilla (18951, -0.79 DPS) [quest]; Darkspear Voodoo Seal (272061, -1.13 DPS) [vendor] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Shard of the Splithooves (10659, -2.35 DPS, sim-verified) [quest]; Evonice's Landin' Pilla (18951, -4.76 DPS) [quest]; Darkspear Voodoo Seal (272061, -5.10 DPS) [vendor] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Mace (23464, +0.00 DPS) [vendor]; Hand of Edward the Odd (2243, -13.88 DPS, sim-verified) [world_drop]; Resurgence Rod (17743, -22.07 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 26.0 healing_power points (12.59 DPS) | yes | Brilliant Wand (249385, -7.13 DPS, sim-verified) [crafted]; Earthen Rod (9381, -8.23 DPS) [dungeon]; Greater Eternal Wand (249237, -8.23 DPS) [crafted] |

**New at 60:** head: Mooncloth Circlet; shoulder: Mooncloth Shoulders; back: Battle Healer's Cloak; chest: Truefaith Vestments; wrist: Dryad's Wrist Bindings; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Maiden's Circle; finger2: Advisor's Ring; ranged: Torch of Light

No-known-source sample (15 of 1103, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

