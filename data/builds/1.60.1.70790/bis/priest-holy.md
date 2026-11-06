# Leveling BiS: Holy

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 000000000000000000-33500000000000000-00000000000000000)

Set DPS (verified): 41.9. Weights run: 15.4s. Verify run: 82.6s. 186 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.048, intellect=0.941 ± 0.010, spirit=0.947 ± 0.011, mp5=1.827 ± 0.013, crit=0.064 ± 0.004 per rating point (14 rating = 1%, 0.893 per %), spell_haste=not significant (-0.028 ± 0.086)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.71 DPS) | yes | Flying Tiger Goggles (4368, -1.35 DPS) [crafted]; Shadow Goggles (4373, -1.48 DPS, sim-verified) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 1.9 healing_power points (0.18 DPS) | yes | Tarnished Locket (279870, -0.65 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.5 healing_power points (1.28 DPS) | yes | Slime-encrusted Pads (6461, -0.76 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.90 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -1.14 DPS, sim-verified) [crafted] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.51 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -0.58 DPS) [crafted]; Spirit Cloak (4792, -0.68 DPS) [vendor] |
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

No-known-source sample (15 of 186, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes

### Band 30 (human, 000000000000000000-33505003020000000-00000000000000000)

Set DPS (verified): 66.2. Weights run: 15.0s. Verify run: 96.3s. 308 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.028, intellect=0.802 ± 0.011, spirit=0.458 ± 0.015, mp5=1.845 ± 0.027, crit=0.081 ± 0.004 per rating point (14 rating = 1%, 1.130 per %), spell_haste=not significant (0.060 ± 0.052)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 35.7 healing_power points (4.04 DPS) | yes | Filigreed Pristine Circlet (253975, -1.55 DPS) [crafted]; Pristine Circlet (253949, -2.01 DPS) [crafted]; Enchanter's Cowl (4322, -2.46 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 5.0 healing_power points (0.57 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.03 DPS) [world]; Crystal Starfire Medallion (5003, -0.05 DPS) [world_drop]; Seal of the Expedition (271920, -0.12 DPS) [rep] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | sim-verified (66.2 DPS) | yes | Death Speaker Mantle (6685, -0.16 DPS) [dungeon]; Magician's Mantle (12998, -0.45 DPS) [world_drop]; Invoker's Mantle (215365, -0.59 DPS) [crafted] |
| back | Caretaker's Cape (19533) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Repairman's Cape (9605, -0.47 DPS) [quest]; Glowing Thresher Cape (6901, -0.47 DPS) [dungeon]; Prelacy Cape (7004, -0.73 DPS) [quest] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 28.8 healing_power points (3.26 DPS) | yes | Pristine Gown (253961, -0.16 DPS) [crafted]; Beguiler Robes (7728, -0.74 DPS) [dungeon]; Robes of Arugal (6324, -0.91 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 14.5 healing_power points (1.64 DPS) | yes | Mindthrust Bracers (1974, -0.56 DPS) [dungeon]; Glowing Magical Bracelets (13106, -0.92 DPS) [world_drop]; Nightsky Wristbands (6407, -0.94 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | sim-verified (66.2 DPS) | yes | Truefaith Gloves (7049, +0.00 DPS) [crafted]; Pristine Gloves (253913, +0.00 DPS) [crafted]; Hotshot Pilot's Gloves (9491, -0.09 DPS) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 14.2 healing_power points (1.61 DPS) | yes | Belt of Arugal (6392, -0.06 DPS) [dungeon]; Highlander's Cloth Girdle (20099, -0.09 DPS) [rep]; Crimson Silk Belt (7055, -0.29 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (66.2 DPS) | yes | Filigreed Pristine Leggings (253937, -0.59 DPS) [crafted]; Earthen Leggings (253999, -1.11 DPS) [crafted]; Gaze Dreamer Pants (6903, -1.56 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 27.4 healing_power points (3.10 DPS) | yes | Acidic Walkers (9454, -1.60 DPS) [dungeon]; Pristine Boots (253889, -1.81 DPS) [crafted]; Spidersilk Boots (4320, -1.95 DPS) [crafted] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 12.2 healing_power points (1.38 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; Lorekeeper's Ring (19525, -0.17 DPS) [rep]; Snake Hoop (6750, -0.25 DPS) [quest] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | sim-verified (66.2 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; Lorekeeper's Ring (19525, -0.10 DPS) [rep]; Snake Hoop (6750, -0.17 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of the Friar (3415, -0.91 DPS) [dungeon]; Odo's Ley Staff (6318, -1.19 DPS) [dungeon]; Staff of Westfall (2042, -1.50 DPS) [quest] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 15.7 healing_power points (1.78 DPS) | yes | Eye of Paleth (2943, -0.31 DPS) [quest]; Orb of Mystic Insight (249394, -0.45 DPS) [crafted]; Orb of Mistmantle (13031, -0.56 DPS) [world_drop] |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 6.9 healing_power points (0.78 DPS) | yes | Thunderwood (13062, -0.10 DPS) [world_drop]; Wand of Eventide (5214, -0.21 DPS) [world_drop]; Greater Mystic Wand (217287, -0.21 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Caretaker's Cape; chest: Death Speaker Robes; wrist: Spidertank Oilrag; hands: Gloves of Old; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Ring of Power Regulation; main_hand: Death Speaker Scepter; off_hand: Orb of Souls; ranged: Wand of Mana Concentration

No-known-source sample (15 of 308, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (human, 000000000000000000-33505003030121131-00000000000000000)

Set DPS (verified): 168.2. Weights run: 26.8s. Verify run: 174.8s. 402 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.316, intellect=0.699 ± 0.042, spirit=0.914 ± 0.040, mp5=1.696 ± 0.047, crit=0.191 ± 0.033 per rating point (14 rating = 1%, 2.670 per %), spell_haste=not significant (0.197 ± 0.389)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 42.1 healing_power points (11.68 DPS) | yes | Whitemane's Chapeau (7720, -1.15 DPS) [dungeon]; Holy Shroud (2721, -1.83 DPS, sim-verified) [world_drop]; Earthen Silk Hood (254015, -3.04 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 25.6 healing_power points (7.10 DPS) | yes | Triune Amulet (7722, -2.82 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -3.97 DPS) [dungeon]; Amberglow Talisman (10824, -4.56 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 31.1 healing_power points (8.62 DPS) | yes | Earthen Silk Shoulders (254033, -2.55 DPS, sim-verified) [crafted]; Mantle of Doan (7712, -3.49 DPS) [dungeon]; Green Silken Shoulders (7057, -4.27 DPS) [crafted] |
| back | Caretaker's Cape (19532) | Silverwing Sentinels [rep] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Glowing Thresher Cape (6901, -2.29 DPS) [dungeon]; Repairman's Cape (9605, -2.75 DPS) [quest]; Icy Cloak (4327, -2.80 DPS) [crafted] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 65.3 healing_power points (18.11 DPS) | yes | Dreamweave Vest (10021, -5.59 DPS, sim-verified) [crafted]; Civinad Robes (9623, -7.86 DPS) [quest]; Doomsayer's Robe (4746, -9.11 DPS) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Earthen Silk Cuffs (254019, -0.60 DPS, sim-verified) [crafted]; Arcane Runed Bracers (4744, -1.41 DPS) [quest]; Mindthrust Bracers (1974, -1.53 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 27.2 healing_power points (7.54 DPS) | yes | Gilded Handwraps (254021, -0.08 DPS) [crafted]; Earthen Silk Gloves (254017, -0.52 DPS) [crafted]; Stormcloth Gloves (10011, -2.16 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 33.1 healing_power points (9.17 DPS) | yes | Deathmage Sash (10771, -3.04 DPS, sim-verified) [dungeon]; Highlander's Cloth Girdle (20098, -4.51 DPS) [rep]; Star Belt (4329, -4.55 DPS) [crafted] |
| legs | Stoneweaver Leggings (9407) | Uldaman: Ironaya [dungeon] | 36.6 healing_power points (10.14 DPS) | yes | Pristine Leggings (253987, -1.25 DPS, sim-verified) [crafted]; Stormcloth Pants (10010, -2.18 DPS) [crafted]; Filigreed Pristine Leggings (253937, -2.97 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 28.5 healing_power points (7.92 DPS) | yes | Earthen Silk Slippers (254013, -1.15 DPS) [crafted]; Thoughtcast Boots (10578, -1.70 DPS, sim-verified) [dungeon]; Furen's Boots (13100, -2.50 DPS) [world_drop] |
| finger1 | Electrocutioner Lagnut (9447) | Gnomeregan: Electrocutioner 6000 [dungeon] | 15.4 healing_power points (4.27 DPS) | yes | Philanthropist's Ring (281635, -0.53 DPS) [quest]; Ring of Power Regulation (273030, -1.20 DPS) [dungeon]; Runemender's Seal (273031, -3.18 DPS, sim-verified) [world] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Philanthropist's Ring (281635, -0.16 DPS) [quest]; Runemender's Seal (273031, -0.57 DPS, sim-verified) [world]; Ring of Power Regulation (273030, -0.83 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+4.4 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Jordan (873, -7.96 DPS, sim-verified) [world_drop]; Death Speaker Scepter (2816, -9.24 DPS) [dungeon]; Staff of the Friar (3415, -10.05 DPS) [dungeon] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | 25.7 healing_power points (7.11 DPS) | yes | Orb of Lorica (11262, -1.05 DPS, sim-verified) [quest]; Orb of Souls (249395, -1.99 DPS) [crafted]; Orb of the Forgotten Seer (7685, -3.23 DPS) [dungeon] |
| ranged | Earthen Rod (9381) | Uldaman [dungeon] | 9.0 healing_power points (2.50 DPS) | yes | Captain Rackmore's Tiller (16789, -0.07 DPS) [quest]; Jaina's Firestarter (13064, -0.57 DPS) [world_drop]; Wand of Mana Concentration (273043, -0.78 DPS) [world] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Caretaker's Cape; chest: Stormcloth Vest; hands: Dreamweave Gloves; waist: Gilded Cord; legs: Stoneweaver Leggings; finger2: Lorekeeper's Ring; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Hand of Righteousness; off_hand: Beacon of Hope; ranged: Earthen Rod

No-known-source sample (15 of 402, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (human, 005002000000000000-33505003030121431-00000000000000000)

Set DPS (verified): 247.5. Weights run: 28.4s. Verify run: 243.3s. 517 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.411, intellect=1.002 ± 0.062, spirit=0.434 ± 0.056, mp5=0.871 ± 0.058, crit=0.281 ± 0.047 per rating point (14 rating = 1%, 3.936 per %), spell_haste=not significant (0.813 ± 0.625)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Cassandra's Grace (13102) | World drop [world_drop] | 47.9 healing_power points (12.78 DPS) | yes | Knight-Lieutenant's Satin Cover (220896, -0.41 DPS) [vendor]; Papal Fez (9431, -0.42 DPS, sim-verified) [dungeon]; Red Mageweave Headband (10033, -2.37 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 28.0 healing_power points (7.47 DPS) | yes | Horizon Choker (13085, -1.82 DPS, sim-verified) [world_drop]; Mindburst Medallion (11196, -3.69 DPS) [quest]; Gemshard Heart (17707, -4.10 DPS) [dungeon] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | 36.6 healing_power points (9.77 DPS) | yes | Knight-Lieutenant's Satin Pads (220894, -0.30 DPS) [vendor]; Inquisitor's Shawl (19507, -0.89 DPS, sim-verified) [dungeon]; Rotgrip Mantle (17732, -1.49 DPS) [dungeon] |
| back | Caretaker's Cape (19531) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -0.53 DPS) [dungeon]; Runecloth Cloak (13860, -2.02 DPS) [crafted]; Expeditionary's Cape (271921, -2.24 DPS) [rep] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 61.5 healing_power points (16.39 DPS) | yes | Knight's Satin Armor (220892, -2.74 DPS, sim-verified) [vendor]; Grimlok's Tribal Vestments (9415, -5.70 DPS) [dungeon]; Acumen Robes (17775, -5.98 DPS) [quest] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 29.2 healing_power points (7.78 DPS) | yes | Aristocratic Cuffs (12546, -2.28 DPS, sim-verified) [dungeon]; Bloodband Bracers (11469, -3.70 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -4.23 DPS) [quest] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 42.5 healing_power points (11.33 DPS) | yes | Earthenweave Gloves (254075, -3.12 DPS) [crafted]; Gilded Handwraps (254021, -3.59 DPS) [crafted]; Sergeant Major's Satin Gloves (220897, -4.36 DPS, sim-verified) [vendor] |
| waist | Gilded Waistcord (254081) | Tailoring [crafted] | 38.0 healing_power points (10.14 DPS) | yes | Gilded Cord (254037, -0.49 DPS, sim-verified) [crafted]; Dawnspire Cord (12466, -1.86 DPS) [dungeon]; Earthenweave Cord (254077, -1.93 DPS) [crafted] |
| legs | Knight's Satin Leggings (220893) | Captain Dirgehammer [vendor] | 43.6 healing_power points (11.64 DPS) | yes | Senior Designer's Pantaloons (11841, -1.19 DPS, sim-verified) [dungeon]; Stoneweaver Leggings (9407, -2.78 DPS) [dungeon]; Spellshock Leggings (9484, -2.83 DPS) [dungeon] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 40.0 healing_power points (10.67 DPS) | yes | Sergeant Major's Satin Boots (220895, -1.48 DPS, sim-verified) [vendor]; Earthenweave Boots (254093, -2.46 DPS) [crafted]; Gilded Slippers (254001, -3.00 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cyclopean Band (11824, +0.00 DPS) [dungeon]; Philanthropist's Ring (281635, -0.13 DPS) [quest]; Band of the Unicorn (7553, -0.66 DPS) [world_drop] |
| finger2 | Brainlash (6440) | Zul'Farrak [dungeon] | 17.2 healing_power points (4.59 DPS) | yes | Cyclopean Band (11824, +0.00 DPS) [dungeon]; Philanthropist's Ring (281635, -0.58 DPS) [quest]; Band of the Unicorn (7553, -1.12 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+6.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.23 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.46 DPS) [quest] |
| trinket2 | Enriched Seal (271903) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.00 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.24 DPS) [quest]; Uther's Strength (11302, -0.60 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -8.43 DPS, sim-verified) [dungeon]; Resurgence Rod (17743, -16.37 DPS) [quest]; Dreamstaff (249454, -18.59 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Charged Lightning Rod (11860) | Ledger from Tanaris [quest] | 9.0 healing_power points (2.40 DPS) | yes | Lesser Eternal Wand (249232, -0.27 DPS) [crafted]; Jaina's Firestarter (13064, -0.45 DPS) [world_drop]; Earthen Rod (9381, -1.31 DPS, sim-verified) [dungeon] |

**New at 50:** head: Cassandra's Grace; shoulder: Nethergeld Shoulders; back: Caretaker's Cape; wrist: Nethergeld Cuffs; hands: Gilded Gloves; waist: Gilded Waistcord; legs: Knight's Satin Leggings; feet: Gilded Sandals; finger1: Lorekeeper's Ring; finger2: Brainlash; trinket2: Enriched Seal; main_hand: Glowing Brightwood Staff; ranged: Charged Lightning Rod

No-known-source sample (15 of 517, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (human, 005005031030000000-33505003030121431-00000000000000000)

Set DPS (verified): 403.8. Weights run: 27.1s. Verify run: 895.7s. 1359 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.284, intellect=1.082 ± 0.069, spirit=not significant (0.229 ± 0.074), mp5=3.512 ± 0.116, crit=0.403 ± 0.042 per rating point (14 rating = 1%, 5.639 per %), spell_haste=not significant (-0.472 ± 0.879)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Crown of the Penitent (13216, -10.73 DPS, sim-verified) [quest] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Animated Chain Necklace (18723, +0.00 DPS) [dungeon]; Tooth of Gnarr (13141, -0.77 DPS) [dungeon]; Stormpike Sage's Pendant (19098, -0.99 DPS) [rep] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Satin Epaulets (227136, +0.00 DPS) [vendor]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp]; Highlander's Epaulets (20061, -9.65 DPS, sim-verified) [rep] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gracious Cape (18743, -4.24 DPS, sim-verified) [dungeon]; Caretaker's Cape (19530, -4.49 DPS) [rep]; Arcanoweave Cloak (272411, -5.32 DPS) [vendor] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 76.7 healing_power points (19.99 DPS) | yes | Field Marshal's Satin Vestments (17605, +0.00 DPS) [vendor]; Field Marshal's Satin Tunic (231624, +0.00 DPS) [vendor]; Robes of the Exalted (13346, -2.10 DPS, sim-verified) [dungeon] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (+8.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [rep]; Nethergeld Cuffs (254061, +0.00 DPS) [crafted] |
| hands | Theramore Handwraps (271945) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Satin Handwraps (231623, +0.00 DPS) [vendor]; Hands of the Exalted Herald (12554, -0.89 DPS) [dungeon]; Knight-Lieutenant's Satin Handwraps (227128, -2.04 DPS) [pvp] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 48.2 healing_power points (12.57 DPS) | yes | Devout Belt (16696, -3.52 DPS) [dungeon]; Gilded Waistcord (254081, -3.80 DPS, sim-verified) [crafted]; Dawnspire Cord (12466, -4.08 DPS) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Satin Pants (17603, +0.00 DPS) [vendor]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Sentinel's Silk Leggings (237815, -14.05 DPS, sim-verified) [vendor] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Satin Walkers (231627, +0.00 DPS) [vendor]; Swarmtender's Footpads (275609, +0.00 DPS) [crafted]; Mooncloth Boots (15802, -8.14 DPS, sim-verified) [crafted] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, +0.00 DPS) [dungeon]; Fordring's Seal (16058, +0.00 DPS) [quest]; Lorekeeper's Ring (19522, -1.25 DPS) [rep] |
| finger2 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, +0.00 DPS) [dungeon]; Fordring's Seal (16058, -1.70 DPS, sim-verified) [quest]; Lorekeeper's Ring (19522, -3.10 DPS) [rep] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+4.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Second Wind (11819, +0.00 DPS) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Second Wind (11819, -5.22 DPS) [dungeon]; Weakness Analyzer (272438, -5.22 DPS) [vendor]; Shard of the Splithooves (10659, -11.90 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Warhammer (23454, +0.00 DPS) [vendor]; Resurgence Rod (17743, -8.56 DPS) [quest]; Hand of Edward the Odd (2243, -10.48 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 19.5 healing_power points (5.09 DPS) | yes | Bonecreeper Stylus (13938, -1.09 DPS) [dungeon]; Oblivion's Touch (18761, -1.98 DPS) [dungeon]; Brilliant Wand (249385, -2.36 DPS, sim-verified) [crafted] |

**New at 60:** head: Devout Crown; shoulder: Devout Mantle; back: Drape of Recovery; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Theramore Handwraps; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Devout Sandals; finger1: Blessed Band of Light; finger2: Band of Mending; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 1359, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (human, 005005031030000000-33505003030121431-00000000000000000)

Set DPS (verified): 867.3. Weights run: 17.6s. Verify run: 487.1s. 1359 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.193, intellect=0.513 ± 0.040, spirit=0.748 ± 0.039, mp5=1.366 ± 0.042, crit=0.410 ± 0.029 per rating point (14 rating = 1%, 5.737 per %), spell_haste=not significant (0.199 ± 0.743)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -18.41 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Animated Chain Necklace (18723, +0.00 DPS) [dungeon]; Lady Maye's Pendant (14558, -2.86 DPS) [world_drop]; Archlight Talisman (15856, -2.91 DPS) [quest] |
| shoulder | Mooncloth Shoulders (14139) | Tailoring [crafted] | 45.2 healing_power points (18.85 DPS) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp]; Earthenweave Mantle (254119, -5.48 DPS, sim-verified) [crafted] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Caretaker's Cape (19530, -3.86 DPS, sim-verified) [rep]; Crystalline Threaded Cape (20697, -6.01 DPS) [world] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 85.0 healing_power points (35.41 DPS) | yes | Field Marshal's Satin Tunic (231624, -0.78 DPS) [vendor]; Robes of the Exalted (13346, -5.36 DPS, sim-verified) [dungeon]; Stormcloth Vest (10020, -8.75 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (+16.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [rep]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Lieutenant's Satin Handwraps (227128, +0.00 DPS) [pvp]; Marshal's Satin Handwraps (231623, +0.00 DPS) [vendor]; Hands of the Exalted Herald (12554, -15.72 DPS, sim-verified) [dungeon] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 44.9 healing_power points (18.72 DPS) | yes | Penitent's Cinch (272394, -3.42 DPS, sim-verified) [vendor]; Earthenweave Cord (254077, -4.46 DPS) [crafted]; Gilded Waistcord (254081, -4.71 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Satin Legguards (227125, +0.00 DPS) [pvp]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Earthenweave Leggings (254125, -25.70 DPS, sim-verified) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Satin Walkers (231627, +0.00 DPS) [vendor]; Swarmtender's Footpads (275609, +0.00 DPS) [crafted]; Faith Healer's Boots (22247, -16.13 DPS, sim-verified) [dungeon] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fordring's Seal (16058, +0.00 DPS) [quest]; Rosewine Circle (13178, -0.48 DPS) [dungeon]; Maiden's Circle (13001, -4.54 DPS) [world_drop] |
| finger2 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-verified (+3.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Fordring's Seal (16058, +0.00 DPS) [quest]; Rosewine Circle (13178, -1.36 DPS) [dungeon]; Maiden's Circle (13001, -5.42 DPS) [world_drop] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -8.34 DPS) [vendor]; Second Wind (11819, -8.40 DPS, sim-verified) [dungeon]; Uther's Strength (11302, -10.00 DPS) [world_drop] |
| trinket2 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-verified (+10.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Second Wind (11819, +0.00 DPS) [dungeon]; Burst of Knowledge (11832, +0.00 DPS) [dungeon] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Resurgence Rod (17743, -21.13 DPS, sim-verified) [quest]; Dreamstaff (249454, -21.84 DPS) [crafted]; Grand Marshal's Warhammer (23454, -25.14 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 25.2 healing_power points (10.52 DPS) | yes | Brilliant Wand (249385, -5.27 DPS) [crafted]; Ritssyn's Wand of Bad Mojo (22408, -5.93 DPS) [dungeon]; Bonecreeper Stylus (13938, -7.13 DPS, sim-verified) [dungeon] |

**New at 60:** head: Devout Crown; shoulder: Mooncloth Shoulders; back: Drape of Recovery; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Devout Sandals; finger1: Blessed Band of Light; finger2: Band of Mending; trinket1: Serenity Field; trinket2: Dormant Heart of the Mountain; ranged: Torch of Light

No-known-source sample (15 of 1359, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

## Horde

### Band 20 (undead, 000000000000000000-33500000000000000-00000000000000000)

Set DPS (verified): 40.3. Weights run: 15.4s. Verify run: 84.3s. 179 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.107, intellect=1.074 ± 0.017, spirit=0.981 ± 0.017, mp5=2.493 ± 0.019, crit=0.110 ± 0.007 per rating point (14 rating = 1%, 1.534 per %), spell_haste=not significant (-0.017 ± 0.150)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (2.00 DPS) | yes | Shadow Goggles (4373, -1.49 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.56 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Roadwatcher's Confidence (281265, -0.18 DPS, sim-verified) [quest] |
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

No-known-source sample (15 of 179, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20441 Scout's Blade; 23173 Abomination Skin Leggings

### Band 30 (undead, 000000000000000000-33505003020000000-00000000000000000)

Set DPS (verified): 68.4. Weights run: 15.0s. Verify run: 97.4s. 301 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.028, intellect=0.802 ± 0.011, spirit=0.458 ± 0.015, mp5=1.845 ± 0.027, crit=0.081 ± 0.004 per rating point (14 rating = 1%, 1.130 per %), spell_haste=not significant (0.060 ± 0.052)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 35.7 healing_power points (4.04 DPS) | yes | Filigreed Pristine Circlet (253975, -1.55 DPS) [crafted]; Pristine Circlet (253949, -2.01 DPS) [crafted]; Enchanter's Cowl (4322, -2.46 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kaleidoscope Chain (13084, +0.00 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.03 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.09 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Mantle (6685, -0.16 DPS) [dungeon]; Magician's Mantle (12998, -0.45 DPS) [world_drop]; Invoker's Mantle (215365, -0.59 DPS) [crafted] |
| back | Battle Healer's Cloak (19529) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Thresher Cape (6901, -0.47 DPS) [dungeon]; Darkspear Raider's Cloak (272078, -0.80 DPS) [vendor]; Windsong Drape (15468, -0.90 DPS) [quest] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 28.8 healing_power points (3.26 DPS) | yes | Pristine Gown (253961, -0.16 DPS) [crafted]; Beguiler Robes (7728, -0.74 DPS) [dungeon]; Robes of Arugal (6324, -0.91 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 14.5 healing_power points (1.64 DPS) | yes | Mindthrust Bracers (1974, -0.56 DPS) [dungeon]; Glowing Magical Bracelets (13106, -0.92 DPS) [world_drop]; Nightsky Wristbands (6407, -0.94 DPS) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 17.4 healing_power points (1.97 DPS) | yes | Pristine Gloves (253913, -0.45 DPS) [crafted]; Jutebraid Gloves (10654, -0.84 DPS) [quest]; Gloves of Old (9395, -0.89 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 14.2 healing_power points (1.61 DPS) | yes | Belt of Arugal (6392, -0.06 DPS) [dungeon]; Defiler's Cloth Girdle (20164, -0.09 DPS) [rep]; Crimson Silk Belt (7055, -0.29 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 29.9 healing_power points (3.38 DPS) | yes | Filigreed Pristine Leggings (253937, -0.59 DPS) [crafted]; Earthen Leggings (253999, -1.11 DPS) [crafted]; Gaze Dreamer Pants (6903, -1.56 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 27.4 healing_power points (3.10 DPS) | yes | Acidic Walkers (9454, -1.60 DPS) [dungeon]; Pristine Boots (253889, -1.81 DPS) [crafted]; Spidersilk Boots (4320, -1.95 DPS) [crafted] |
| finger1 | Solid Gold Band (279845) | Opportunistic Education [quest] | 13.3 healing_power points (1.50 DPS) | yes | Electrocutioner Lagnut (9447, -0.12 DPS) [dungeon]; Runemender's Seal (273031, -0.12 DPS) [world]; Advisor's Ring (19521, -0.29 DPS) [rep] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Electrocutioner Lagnut (9447, +0.00 DPS) [dungeon]; Runemender's Seal (273031, +0.00 DPS) [world]; Advisor's Ring (19521, -0.10 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of the Friar (3415, -0.91 DPS) [dungeon]; Odo's Ley Staff (6318, -1.19 DPS) [dungeon]; Living Root (6631, -2.34 DPS) [dungeon] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Orb of Souls (249395, +0.00 DPS) [crafted]; Orb of Mistmantle (13031, -0.11 DPS) [world_drop]; Witch's Finger (16887, -0.61 DPS) [quest] |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 6.9 healing_power points (0.78 DPS) | yes | Thunderwood (13062, -0.10 DPS) [world_drop]; Wand of Eventide (5214, -0.21 DPS) [world_drop]; Greater Mystic Wand (217287, -0.21 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Battle Healer's Cloak; chest: Death Speaker Robes; wrist: Spidertank Oilrag; hands: Truefaith Gloves; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Solid Gold Band; finger2: Ring of Power Regulation; main_hand: Death Speaker Scepter; off_hand: Orb of Mystic Insight; ranged: Wand of Mana Concentration

No-known-source sample (15 of 301, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (undead, 000000000000000000-33505003030121131-00000000000000000)

Set DPS (verified): 178.1. Weights run: 26.8s. Verify run: 172.6s. 395 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.316, intellect=0.699 ± 0.042, spirit=0.914 ± 0.040, mp5=1.696 ± 0.047, crit=0.191 ± 0.033 per rating point (14 rating = 1%, 2.670 per %), spell_haste=not significant (0.197 ± 0.389)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 42.1 healing_power points (11.68 DPS) | yes | Holy Shroud (2721, -1.00 DPS) [world_drop]; Whitemane's Chapeau (7720, -1.15 DPS) [dungeon]; Earthen Silk Hood (254015, -3.04 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 25.6 healing_power points (7.10 DPS) | yes | Triune Amulet (7722, -3.10 DPS) [dungeon]; Necklace of Calisea (1714, -3.97 DPS) [dungeon]; Amberglow Talisman (10824, -4.56 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 31.1 healing_power points (8.62 DPS) | yes | Earthen Silk Shoulders (254033, -1.60 DPS) [crafted]; Mantle of Doan (7712, -3.49 DPS) [dungeon]; Green Silken Shoulders (7057, -4.27 DPS) [crafted] |
| back | Battle Healer's Cloak (19528) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Thresher Cape (6901, -2.29 DPS) [dungeon]; Icy Cloak (4327, -2.80 DPS) [crafted]; Cloak of Blight (6832, -2.96 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 65.3 healing_power points (18.11 DPS) | yes | Dreamweave Vest (10021, -7.83 DPS) [crafted]; Civinad Robes (9623, -7.86 DPS) [quest]; Doomsayer's Robe (4746, -9.11 DPS) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Earthen Silk Cuffs (254019, +0.00 DPS) [crafted]; Radiant Silver Bracers (4545, -1.25 DPS) [quest]; Mindthrust Bracers (1974, -1.53 DPS) [dungeon] |
| hands | Gilded Handwraps (254021) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamweave Gloves (10019, +0.00 DPS) [crafted]; Earthen Silk Gloves (254017, -0.44 DPS) [crafted]; Stormcloth Gloves (10011, -2.08 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 33.1 healing_power points (9.17 DPS) | yes | Deathmage Sash (10771, -4.32 DPS) [dungeon]; Defiler's Cloth Girdle (20166, -4.51 DPS) [rep]; Star Belt (4329, -4.55 DPS) [crafted] |
| legs | Stoneweaver Leggings (9407) | Uldaman: Ironaya [dungeon] | 36.6 healing_power points (10.14 DPS) | yes | Pristine Leggings (253987, -1.41 DPS) [crafted]; Stormcloth Pants (10010, -2.18 DPS) [crafted]; Filigreed Pristine Leggings (253937, -2.97 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 28.5 healing_power points (7.92 DPS) | yes | Thoughtcast Boots (10578, -1.02 DPS) [dungeon]; Earthen Silk Slippers (254013, -1.15 DPS) [crafted]; Furen's Boots (13100, -2.50 DPS) [world_drop] |
| finger1 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | sim-verified (178.1 DPS) | yes | Electrocutioner Lagnut (9447, +0.00 DPS) [dungeon]; Runemender's Seal (273031, +0.00 DPS) [world]; Solid Gold Band (279845, +0.00 DPS) [quest] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | sim-verified (178.1 DPS) | yes | Electrocutioner Lagnut (9447, +0.00 DPS) [dungeon]; Runemender's Seal (273031, +0.00 DPS) [world]; Solid Gold Band (279845, +0.00 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Jordan (873, -5.05 DPS) [world_drop]; Death Speaker Scepter (2816, -9.24 DPS) [dungeon]; Staff of the Friar (3415, -10.05 DPS) [dungeon] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | 25.7 healing_power points (7.11 DPS) | yes | Orb of Souls (249395, -1.99 DPS) [crafted]; Omega Orb (7749, -3.21 DPS) [quest]; Orb of the Forgotten Seer (7685, -3.23 DPS) [dungeon] |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | sim-verified (178.1 DPS) | yes | Earthen Rod (9381, +0.00 DPS) [dungeon]; Jaina's Firestarter (13064, +0.00 DPS) [world_drop]; Captain Rackmore's Tiller (16789, +0.00 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Battle Healer's Cloak; chest: Stormcloth Vest; hands: Gilded Handwraps; waist: Gilded Cord; legs: Stoneweaver Leggings; finger1: Darkspear Signet; finger2: Advisor's Ring; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Hand of Righteousness; off_hand: Beacon of Hope

No-known-source sample (15 of 395, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger; 5743 Prismstone Ring

### Band 50 (undead, 005002000000000000-33505003030121431-00000000000000000)

Set DPS (verified): 260.7. Weights run: 28.4s. Verify run: 251.7s. 510 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.411, intellect=1.002 ± 0.062, spirit=0.434 ± 0.056, mp5=0.871 ± 0.058, crit=0.281 ± 0.047 per rating point (14 rating = 1%, 3.936 per %), spell_haste=not significant (0.813 ± 0.625)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Cassandra's Grace (13102) | World drop [world_drop] | 47.9 healing_power points (12.78 DPS) | yes | Blood Guard's Satin Cover (220899, -0.41 DPS) [vendor]; Papal Fez (9431, -1.33 DPS) [dungeon]; Red Mageweave Headband (10033, -2.37 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 28.0 healing_power points (7.47 DPS) | yes | Horizon Choker (13085, -1.15 DPS, sim-verified) [world_drop]; Mindburst Medallion (11196, -3.69 DPS) [quest]; Gemshard Heart (17707, -4.10 DPS) [dungeon] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | 36.6 healing_power points (9.77 DPS) | yes | Blood Guard's Satin Pads (220901, -0.30 DPS) [vendor]; Rotgrip Mantle (17732, -1.49 DPS) [dungeon]; Inquisitor's Shawl (19507, -1.61 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Deep Woodlands Cloak (19121, -0.43 DPS) [quest]; Battle Healer's Cloak (19527, -0.64 DPS, sim-verified) [rep] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 61.5 healing_power points (16.39 DPS) | yes | Stone Guard's Satin Armor (220903, -2.62 DPS, sim-verified) [vendor]; Grimlok's Tribal Vestments (9415, -5.70 DPS) [dungeon]; Acumen Robes (17775, -5.98 DPS) [quest] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 29.2 healing_power points (7.78 DPS) | yes | Aristocratic Cuffs (12546, -1.73 DPS, sim-verified) [dungeon]; Bloodband Bracers (11469, -3.70 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -4.23 DPS) [quest] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 42.5 healing_power points (11.33 DPS) | yes | Greenleaf Handwraps (19116, -2.53 DPS) [quest]; Earthenweave Gloves (254075, -3.12 DPS) [crafted]; First Sergeant's Satin Gloves (220898, -3.67 DPS, sim-verified) [vendor] |
| waist | Gilded Waistcord (254081) | Tailoring [crafted] | 38.0 healing_power points (10.14 DPS) | yes | Gilded Cord (254037, -0.61 DPS, sim-verified) [crafted]; Dawnspire Cord (12466, -1.86 DPS) [dungeon]; Earthenweave Cord (254077, -1.93 DPS) [crafted] |
| legs | Stone Guard's Satin Leggings (220902) | Lady Palanseer [vendor] | 43.6 healing_power points (11.64 DPS) | yes | Senior Designer's Pantaloons (11841, -1.02 DPS) [dungeon]; Stoneweaver Leggings (9407, -2.78 DPS) [dungeon]; Spellshock Leggings (9484, -2.83 DPS) [dungeon] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 40.0 healing_power points (10.67 DPS) | yes | First Sergeant's Satin Boots (220900, -1.92 DPS, sim-verified) [vendor]; Earthenweave Boots (254093, -2.46 DPS) [crafted]; Gilded Slippers (254001, -3.00 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Brainlash (6440, +0.00 DPS) [dungeon]; Cyclopean Band (11824, +0.00 DPS) [dungeon]; Solid Gold Band (279845, -0.49 DPS) [quest] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Cyclopean Band (11824, +0.00 DPS) [dungeon]; Solid Gold Band (279845, -0.62 DPS) [quest]; Brainlash (6440, -0.68 DPS, sim-verified) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.23 DPS) [quest]; Alchemist's Stone (13503, -0.93 DPS) [crafted]; Uther's Strength (11302, -1.39 DPS, sim-verified) [world_drop] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.00 DPS) [quest]; Alchemist's Stone (13503, -0.70 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -7.81 DPS, sim-verified) [dungeon]; Resurgence Rod (17743, -16.37 DPS) [quest]; Dreamstaff (249454, -18.59 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Charged Lightning Rod (11860) | Ledger from Tanaris [quest] | 9.0 healing_power points (2.40 DPS) | yes | Dancing Flame (6806, -0.27 DPS) [quest]; Lesser Eternal Wand (249232, -0.27 DPS) [crafted]; Earthen Rod (9381, -0.55 DPS, sim-verified) [dungeon] |

**New at 50:** head: Cassandra's Grace; shoulder: Nethergeld Shoulders; back: Spritecaster Cape; wrist: Nethergeld Cuffs; hands: Gilded Gloves; waist: Gilded Waistcord; legs: Stone Guard's Satin Leggings; feet: Gilded Sandals; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket2: Darkspear Voodoo Seal; main_hand: Glowing Brightwood Staff; ranged: Charged Lightning Rod

No-known-source sample (15 of 510, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60 (undead, 005005031030000000-33505003030121431-00000000000000000)

Set DPS (verified): 426.0. Weights run: 27.1s. Verify run: 892.5s. 1354 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.284, intellect=1.082 ± 0.069, spirit=not significant (0.229 ± 0.074), mp5=3.512 ± 0.116, crit=0.403 ± 0.042 per rating point (14 rating = 1%, 5.639 per %), spell_haste=not significant (-0.472 ± 0.879)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -4.82 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Animated Chain Necklace (18723, +0.00 DPS) [dungeon]; Tooth of Gnarr (13141, -0.77 DPS) [dungeon]; Frostwolf Advisor's Pendant (19096, -0.99 DPS) [rep] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Satin Epaulets (227130, +0.00 DPS) [vendor]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp]; Defiler's Epaulets (20176, -6.26 DPS, sim-verified) [rep] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battle Healer's Cloak (19526, -4.49 DPS) [rep]; Arcanoweave Cloak (272411, -5.32 DPS) [vendor]; Gracious Cape (18743, -6.34 DPS, sim-verified) [dungeon] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 76.7 healing_power points (19.99 DPS) | yes | Warlord's Satin Robes (231612, +0.00 DPS) [pvp]; Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Robes of the Exalted (13346, -0.87 DPS, sim-verified) [dungeon] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (+2.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [rep]; Nethergeld Cuffs (254061, +0.00 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Satin Handwraps (231633, +0.00 DPS) [vendor]; Raider Handwraps (272097, +0.00 DPS) [vendor]; Hands of the Exalted Herald (12554, -1.15 DPS, sim-verified) [dungeon] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gilded Waistcord (254081, +0.00 DPS) [crafted]; Dawnspire Cord (12466, -0.56 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -2.01 DPS, sim-verified) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Satin Leggings (231614, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Outrider's Silk Leggings (22747, -8.83 DPS, sim-verified) [rep] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Satin Boots (17618, +0.00 DPS) [vendor]; General's Satin Walkers (231630, +0.00 DPS) [vendor]; Swarmtender's Footpads (275609, -4.07 DPS, sim-verified) [crafted] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, +0.00 DPS) [dungeon]; Fordring's Seal (16058, +0.00 DPS) [quest]; Darkspear Signet (272068, -0.32 DPS) [vendor] |
| finger2 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, +0.00 DPS) [dungeon]; Fordring's Seal (16058, +0.00 DPS) [quest]; Darkspear Signet (272068, -2.18 DPS) [vendor] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (+11.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Second Wind (11819, -5.22 DPS) [dungeon]; Weakness Analyzer (272438, -5.22 DPS) [vendor]; Shard of the Splithooves (10659, -5.46 DPS) [quest] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Second Wind (11819, -0.93 DPS, sim-verified) [dungeon] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Mace (23464, +0.00 DPS) [vendor]; Hand of Edward the Odd (2243, -6.72 DPS, sim-verified) [world_drop]; Resurgence Rod (17743, -8.56 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 19.5 healing_power points (5.09 DPS) | yes | Bonecreeper Stylus (13938, -1.09 DPS) [dungeon]; Oblivion's Touch (18761, -1.98 DPS) [dungeon]; Brilliant Wand (249385, -3.62 DPS, sim-verified) [crafted] |

**New at 60:** head: Devout Crown; shoulder: Devout Mantle; back: Drape of Recovery; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Blessed Band of Light; finger2: Band of Mending; trinket1: Serenity Field; trinket2: Ankh of Life; ranged: Torch of Light

No-known-source sample (15 of 1354, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60, raid preset (undead, 005005031030000000-33505003030121431-00000000000000000)

Set DPS (verified): 879.7. Weights run: 17.6s. Verify run: 496.1s. 1354 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.193, intellect=0.513 ± 0.040, spirit=0.748 ± 0.039, mp5=1.366 ± 0.042, crit=0.410 ± 0.029 per rating point (14 rating = 1%, 5.737 per %), spell_haste=not significant (0.199 ± 0.743)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -19.24 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Animated Chain Necklace (18723, +0.00 DPS) [dungeon]; Lady Maye's Pendant (14558, -2.86 DPS) [world_drop]; Archlight Talisman (15856, -2.91 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp]; Mooncloth Shoulders (14139, -18.37 DPS, sim-verified) [crafted] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battle Healer's Cloak (19526, -2.30 DPS, sim-verified) [rep]; Raincaster Drape (12110, -3.22 DPS) [quest] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 85.0 healing_power points (35.41 DPS) | yes | Warlord's Satin Tunic (231632, -0.78 DPS) [vendor]; Robes of the Exalted (13346, -4.71 DPS, sim-verified) [dungeon]; Stormcloth Vest (10020, -8.75 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (+18.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [rep]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of the Exalted Herald (12554, +0.00 DPS) [dungeon]; Blood Guard's Satin Handwraps (227126, +0.00 DPS) [pvp]; General's Satin Handwraps (231633, +0.00 DPS) [vendor] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 44.9 healing_power points (18.72 DPS) | yes | Penitent's Cinch (272394, -1.36 DPS, sim-verified) [vendor]; Earthenweave Cord (254077, -4.46 DPS) [crafted]; Gilded Waistcord (254081, -4.71 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Earthenweave Leggings (254125, -28.41 DPS, sim-verified) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Satin Walkers (231630, +0.00 DPS) [vendor]; Swarmtender's Footpads (275609, +0.00 DPS) [crafted]; Faith Healer's Boots (22247, -17.50 DPS, sim-verified) [dungeon] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fordring's Seal (16058, +0.00 DPS) [quest]; Rosewine Circle (13178, -0.48 DPS) [dungeon]; Maiden's Circle (13001, -4.54 DPS) [world_drop] |
| finger2 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-verified (+3.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Fordring's Seal (16058, +0.00 DPS) [quest]; Rosewine Circle (13178, -1.36 DPS) [dungeon]; Maiden's Circle (13001, -5.42 DPS) [world_drop] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dormant Heart of the Mountain (249473, -6.59 DPS, sim-verified) [crafted]; Weakness Analyzer (272438, -8.34 DPS) [vendor]; Uther's Strength (11302, -10.00 DPS) [world_drop] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (+10.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -1.67 DPS) [world_drop]; Burst of Knowledge (11832, -4.17 DPS) [dungeon]; Ankh of Life (1713, -5.43 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamstaff (249454, -21.84 DPS) [crafted]; Resurgence Rod (17743, -21.86 DPS, sim-verified) [quest]; High Warlord's Battle Mace (23464, -25.14 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 25.2 healing_power points (10.52 DPS) | yes | Brilliant Wand (249385, -5.27 DPS) [crafted]; Ritssyn's Wand of Bad Mojo (22408, -5.93 DPS) [dungeon]; Bonecreeper Stylus (13938, -9.22 DPS, sim-verified) [dungeon] |

**New at 60:** head: Devout Crown; shoulder: Devout Mantle; back: Drape of Recovery; chest: Truefaith Vestments; wrist: Devout Bracers; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Devout Sandals; finger1: Blessed Band of Light; finger2: Band of Mending; trinket1: Serenity Field; trinket2: Second Wind; ranged: Torch of Light

No-known-source sample (15 of 1354, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

