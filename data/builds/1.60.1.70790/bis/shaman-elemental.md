# Leveling BiS: Elemental

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 5300300000000000-000000000000000000-0000000000000000)

Set DPS (verified): 31.5. Weights run: 4.5s. Verify run: 8.1s. 260 eligible items had no known source.

Race sweep: dwarf 31.5 ± 0.0 (dps); the guide race dwarf is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.588 ± 0.007, crit=0.040 ± 0.001 per rating point (14 rating = 1%, 0.564 per %), hit=0.147 ± 0.005 per rating point (10 rating = 1%, 1.468 per %), spell_haste=-1.903 ± 0.108, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.902 ± 0.002, spirit=not significant (-0.180 ± 0.082), mp5=2.071 ± 0.126

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Hood (252448) | Leatherworking [crafted] | 9.0 spell_power points (0.95 DPS) | yes | Wisdom's Leather Hood (252507, -0.32 DPS) [crafted]; Pristine Circlet (253949, -0.32 DPS) [crafted]; Trapper's Leather Hood (252505, -0.45 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | sim-verified (31.5 DPS) | yes | - |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.3 spell_power points (1.08 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.31 DPS) [crafted]; Slime-encrusted Pads (6461, -0.43 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.66 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pearl-clasped Cloak (5542, -0.02 DPS) [crafted]; Black Whelp Cloak (7283, -0.11 DPS) [crafted]; Caretaker's Cape (20428, -0.11 DPS) [rep] |
| chest | Stormrider's Leather Armor (252492) | Leatherworking [crafted] | 10.9 spell_power points (1.15 DPS) | yes | Wisdom's Leather Armor (252493, -0.32 DPS) [crafted]; Filigreed Pristine Gown (253901, -0.32 DPS) [crafted]; Totemic Leather Armor (252435, -0.33 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 9.1 spell_power points (0.96 DPS) | yes | Owl Bracers (4796, -0.65 DPS) [vendor]; Bright Bracers (3647, -0.71 DPS) [world_drop]; Inscribed Leather Bracers (3205, -0.78 DPS) [dungeon] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.4 spell_power points (0.77 DPS) | yes | Serpent Gloves (5970, -0.04 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.14 DPS) [world]; Windfelt Gloves (5630, -0.17 DPS) [quest] |
| waist | Pristine Sash (253925) (or Stormrider's Leather Belt (252432)) | Tailoring [crafted] | 6.4 spell_power points (0.67 DPS) | yes | Stormrider's Leather Belt (252432, +0.00 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.06 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.11 DPS) [crafted] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 13.5 spell_power points (1.42 DPS) | yes | Wisdom's Leather Pants (252503, -0.32 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.42 DPS) [crafted]; Totemic Leather Pants (252446, -0.44 DPS) [crafted] |
| feet | Stormrider's Leather Boots (252443) | Leatherworking [crafted] | sim-verified (31.5 DPS) | yes | Spidersilk Boots (4320, +0.00 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.21 DPS) [crafted]; Totemic Leather Boots (252442, -0.31 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 9.1 spell_power points (0.96 DPS) | yes | Lavishly Jeweled Ring (1156, -0.59 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.77 DPS) [world_drop]; Furlfeather Ring (286745, -0.86 DPS) [world] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.2 spell_power points (0.65 DPS) | yes | Lavishly Jeweled Ring (1156, -0.28 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.46 DPS) [world_drop]; Furlfeather Ring (286745, -0.54 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | 29.2 spell_power points (3.06 DPS) | yes | Staff of the Blessed Seer (2271, -0.12 DPS) [dungeon]; Emberstone Staff (5201, -0.54 DPS) [dungeon]; Buzzer Blade (2169, -0.75 DPS) [dungeon] |
| off_hand | Seedcloud Buckler (6630) | Wailing Caverns: Verdan the Everliving [dungeon] | 7.9 spell_power points (0.83 DPS) | yes | Worn Turtle Shell Shield (6447, -0.40 DPS) [dungeon]; Dwarven Tome (279898, -0.52 DPS) [quest]; Tear of Grief (5611, -0.53 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 20:** head: Totemic Leather Hood; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Stormrider's Leather Armor; wrist: Mindthrust Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Stormrider's Leather Pants; feet: Stormrider's Leather Boots; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Evocator's Blade; off_hand: Seedcloud Buckler

No-known-source sample (15 of 260, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers

### Band 30 (dwarf, 5500331210001000-000000000000000000-0000000000000000)

Set DPS (verified): 62.2. Weights run: 4.2s. Verify run: 9.5s. 448 eligible items had no known source.

Race sweep: dwarf 62.2 ± 0.1 (dps); the guide race dwarf is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.750 ± 0.010, crit=0.080 ± 0.002 per rating point (14 rating = 1%, 1.113 per %), hit=0.278 ± 0.011 per rating point (10 rating = 1%, 2.784 per %), spell_haste=-1.248 ± 0.189, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.856 ± 0.002, spirit=not significant (0.152 ± 0.093), mp5=1.742 ± 0.133

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 13.5 spell_power points (3.41 DPS) | yes | Enduring Cap (3020, -0.11 DPS) [world_drop]; Totemic Leather Helm (252456, -0.38 DPS) [crafted]; Holy Shroud (2721, -0.40 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 4.5 spell_power points (1.14 DPS) | yes | Seal of the Expedition (271920, -0.19 DPS) [rep]; Kaleidoscope Chain (13084, -0.23 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.26 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.8 spell_power points (3.97 DPS) | yes | Death Speaker Mantle (6685, -0.38 DPS) [dungeon]; Magician's Mantle (12998, -1.01 DPS) [world_drop]; Invoker's Mantle (215365, -1.26 DPS) [crafted] |
| back | Expeditionary's Cape (271923) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Repairman's Cape (9605, -0.37 DPS) [quest]; Hillman's Cloak (3719, -0.37 DPS) [crafted]; Vine Pruner's Cloak (279835, -0.37 DPS) [quest] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 19.2 spell_power points (4.85 DPS) | yes | Green Silk Armor (7065, -0.19 DPS, sim-verified) [crafted]; Mechbuilder's Overalls (9508, -0.56 DPS) [dungeon]; Guardian Armor (4256, -0.69 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 14.2 spell_power points (3.59 DPS) | yes | Mindthrust Bracers (1974, -1.02 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -2.07 DPS) [world_drop]; Nightsky Wristbands (6407, -2.34 DPS) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 8.3 spell_power points (2.08 DPS) | yes | Stormrider's Leather Gloves (252498, -0.06 DPS) [crafted]; Truefaith Gloves (7049, -0.25 DPS) [crafted]; Gloves of Insight (9698, -0.32 DPS) [quest] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 15.1 spell_power points (3.81 DPS) | yes | Highlander's Cloth Girdle (20099, -0.47 DPS) [rep]; Belt of Arugal (6392, -0.78 DPS) [dungeon]; Moss Cinch (6911, -0.78 DPS) [dungeon] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 15.1 spell_power points (3.81 DPS) | yes | Guardian Pants (5962, -0.27 DPS, sim-verified) [crafted]; Dark Ritual Leggings (270031, -0.28 DPS) [quest]; Leech Pants (6910, -0.35 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 12.9 spell_power points (3.25 DPS) | yes | Acidic Walkers (9454, -0.32 DPS) [dungeon]; Spidersilk Boots (4320, -0.72 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.78 DPS) [crafted] |
| finger1 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 11.2 spell_power points (2.83 DPS) | yes | Snake Hoop (6750, -0.37 DPS) [quest]; Minor Channeling Ring (1449, -1.19 DPS) [quest]; Lavishly Jeweled Ring (1156, -1.70 DPS) [dungeon] |
| finger2 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 10.5 spell_power points (2.65 DPS) | yes | Snake Hoop (6750, -0.18 DPS) [quest]; Minor Channeling Ring (1449, -1.00 DPS) [quest]; Lavishly Jeweled Ring (1156, -1.51 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.16 DPS) [dungeon]; Tendonscraper Dagger (273053, -1.16 DPS) [world]; Glimmering Staff (249392, -2.29 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 11.5 spell_power points (2.90 DPS) | yes | Seedcloud Buckler (6630, -0.58 DPS, sim-verified) [dungeon]; Heart of Agamaggan (6694, -1.64 DPS) [dungeon]; Orb of Souls (249395, -1.66 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Expeditionary's Cape; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Skycaller's Leather Belt; feet: Gilded Slippers; finger1: Ring of Power Regulation; finger2: Lorekeeper's Ring; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight

No-known-source sample (15 of 448, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 40 (dwarf, 5500331220001351-000000000000000000-0000000000000000)

Set DPS (verified): 109.6. Weights run: 5.3s. Verify run: 14.9s. 696 eligible items had no known source.

Race sweep: dwarf 109.6 ± 0.2 (dps); the guide race dwarf is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.672 ± 0.021, crit=0.203 ± 0.005 per rating point (14 rating = 1%, 2.840 per %), hit=0.358 ± 0.018 per rating point (10 rating = 1%, 3.581 per %), spell_haste=-2.772 ± 0.300, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.884 ± 0.002, spirit=not significant (-0.486 ± 0.188), mp5=not significant (-1.174 ± 0.330)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellpower Goggles Xtreme (10502, +0.00 DPS) [crafted]; Big Voodoo Mask (8201, -0.24 DPS) [crafted]; Augural Shroud (2620, -0.41 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Triune Amulet (7722, -1.05 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.36 DPS) [quest]; Seal of the Expedition (271919, -1.85 DPS) [rep] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.7 spell_power points (3.79 DPS) | yes | Green Silken Shoulders (7057, -0.08 DPS) [crafted]; Bloodmage Mantle (7684, -0.17 DPS) [dungeon]; Berylline Pads (4197, -0.49 DPS) [quest] |
| back | Expeditionary's Cape (271922) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Long Silken Cloak (4326, +0.00 DPS) [crafted]; Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.09 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamweave Vest (10021, -0.48 DPS) [crafted]; Robe of Power (7054, -0.96 DPS) [crafted]; Elemental Raiment (9434, -1.21 DPS) [world_drop] |
| wrist | Green Whelp Bracers (7386) | Leatherworking [crafted] | 13.3 spell_power points (3.21 DPS) | yes | Guardian Leather Bracers (4260, -0.80 DPS) [crafted]; Turtle Scale Bracers (8198, -0.96 DPS) [crafted]; Arcane Runed Bracers (4744, -1.04 DPS) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamweave Gloves (10019, -2.93 DPS) [crafted]; Red Mageweave Gloves (10018, -3.65 DPS) [crafted]; Skycaller's Leather Gloves (252529, -3.81 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 17.1 spell_power points (4.11 DPS) | yes | Highlander's Cloth Girdle (20098, -0.09 DPS) [rep]; Skycaller's Leather Belt (252522, -0.73 DPS) [crafted]; Gilded Cord (254037, -0.89 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.1 spell_power points (5.31 DPS) | yes | Crimson Silk Pantaloons (7062, -1.28 DPS) [crafted]; Stormrider's Leather Pants (252502, -1.93 DPS) [crafted]; Dark Ritual Leggings (270031, -1.94 DPS) [quest] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.78 DPS) | yes | Skycaller's Leather Shoes (252532, -1.52 DPS) [crafted]; Skycaller's Mail Boots (252563, -1.52 DPS) [crafted]; Mender's Leather Shoes (252533, -2.72 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Snake Hoop (6750, +0.00 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.24 DPS) [quest]; Reedknot Ring (9622, -0.48 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Snake Hoop (6750, -0.96 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.29 DPS) [quest]; Reedknot Ring (9622, -1.53 DPS) [quest] |
| trinket1 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -6.73 DPS) [dungeon]; Hypnotic Blade (7714, -7.09 DPS) [dungeon]; Glimmering Staff (249392, -9.01 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Expeditionary's Cape; chest: Robe of the Magi; wrist: Green Whelp Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Lorekeeper's Ring; finger2: Philanthropist's Ring; trinket1: Enriched Seal; trinket2: Ankh of Life; main_hand: Staff of Jordan

No-known-source sample (15 of 696, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (dwarf, 5500331220001351-050000000000000000-0500000000000000)

Set DPS (verified): 149.2. Weights run: 5.3s. Verify run: 16.7s. 893 eligible items had no known source.

Race sweep: dwarf 149.2 ± 0.3 (dps); the guide race dwarf is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.796 ± 0.015, crit=0.177 ± 0.006 per rating point (14 rating = 1%, 2.483 per %), hit=0.308 ± 0.016 per rating point (10 rating = 1%, 3.078 per %), spell_haste=-4.210 ± 0.167, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.881 ± 0.002, spirit=0.467 ± 0.043, mp5=1.405 ± 0.110

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 38.7 spell_power points (6.76 DPS) | yes | Red Mageweave Headband (10033, -0.66 DPS) [crafted]; Dreamweave Circlet (10041, -0.72 DPS) [crafted]; Charged Scorpid Helm (252581, -1.62 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.4 spell_power points (2.34 DPS) | yes | Mindburst Medallion (11196, -0.04 DPS) [quest]; Horizon Choker (13085, -0.06 DPS) [world_drop]; Triune Amulet (7722, -0.37 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 34.6 spell_power points (6.06 DPS) | yes | Rotgrip Mantle (17732, -1.28 DPS) [dungeon]; Kentic Amice (11624, -1.31 DPS) [dungeon]; Lead Surveyor's Mantle (11842, -3.09 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Expeditionary's Cape (271921, -1.42 DPS) [rep]; Big Voodoo Cloak (8216, -1.65 DPS) [crafted]; Runecloth Cloak (13860, -1.68 DPS, sim-verified) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Dreamweave Vest (10021, -0.56 DPS) [crafted]; Flamestrider Robes (11747, -0.74 DPS) [dungeon]; Feathered Breastplate (8349, -0.83 DPS, sim-verified) [crafted] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 18.0 spell_power points (3.14 DPS) | yes | Skycaller's Leather Bracers (252542, -0.01 DPS) [crafted]; Skycaller's Mail Bracers (252571, -0.01 DPS) [crafted]; Nethergeld Cuffs (254061, -0.53 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodfire Talons (12464, -1.16 DPS) [dungeon]; Skycaller's Leather Gauntlets (252550, -1.20 DPS) [crafted]; Elven Spirit Claws (2564, -2.71 DPS, sim-verified) [world_drop] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 25.5 spell_power points (4.47 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Ban'thok Sash (11662, -0.30 DPS) [dungeon]; Satyrmane Sash (17755, -0.63 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | 31.0 spell_power points (5.41 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -0.91 DPS) [dungeon]; Jinxed Hoodoo Kilt (9474, -1.07 DPS) [dungeon]; Senior Designer's Pantaloons (11841, -4.17 DPS, sim-verified) [dungeon] |
| feet | Skycaller's Leather Boots (252471) | Leatherworking [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Skycaller's Mail Sabatons (252577, +0.00 DPS) [crafted]; Stormcloth Boots (10039, -0.29 DPS) [crafted]; Earthen Silk Slippers (254013, -0.74 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 17.6 spell_power points (3.08 DPS) | yes | Brainlash (6440, -0.59 DPS) [dungeon]; Band of the Unicorn (7553, -0.81 DPS) [world_drop]; Cyclopean Band (11824, -5.54 DPS, sim-verified) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Brainlash (6440, +0.00 DPS) [dungeon]; Cyclopean Band (11824, +0.00 DPS) [dungeon]; Band of the Unicorn (7553, -0.17 DPS) [world_drop] |
| trinket1 | Enriched Seal (271903) | Theramore Expeditionary Force [rep] | sim-verified (+4.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -0.18 DPS) [world_drop]; Ankh of Life (1713, -0.25 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.41 DPS) [quest] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Uther's Strength (11302, -1.48 DPS, sim-verified) [world_drop] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, -2.25 DPS) [quest]; Kindling Stave (11750, -2.30 DPS) [dungeon]; Glowing Brightwood Staff (812, -3.01 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; waist: Skycaller's Leather Waistguard; legs: Spellshock Leggings; feet: Skycaller's Leather Boots; finger1: Lorekeeper's Ring; trinket1: Enriched Seal; trinket2: Thunderbrew's Boot Flask

No-known-source sample (15 of 893, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (dwarf, 5500331220001351-050000000000000000-0530520000000000)

Set DPS (verified): 273.0. Weights run: 5.1s. Verify run: 57.7s. 2086 eligible items had no known source.

Race sweep: dwarf 273.0 ± 0.5 (dps); the guide race dwarf is best.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.123 ± 0.032, crit=0.477 ± 0.012 per rating point (14 rating = 1%, 6.675 per %), hit=0.847 ± 0.052 per rating point (10 rating = 1%, 8.467 per %), spell_haste=-3.513 ± 0.510, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.876 ± 0.002, spirit=not significant (0.563 ± 0.369), mp5=2.895 ± 0.379

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blue Dragonscale Helm (252604) | Leatherworking [crafted] | 53.7 spell_power points (16.39 DPS) | yes | Living Crown (252561, -0.18 DPS) [crafted]; Black Dragonscale Helm (252605, -1.32 DPS) [crafted]; Crown of the Penitent (13216, -2.86 DPS, sim-verified) [quest] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | 27.0 spell_power points (8.24 DPS) | yes | Diana's Pearl Necklace (22403, -0.16 DPS) [dungeon]; Jeweled Amulet of Cainwyn (1443, -0.34 DPS) [world_drop]; Tooth of Gnarr (13141, -0.78 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Highlander's Epaulets (20061, +0.00 DPS) [rep]; Theramore Shoulderpads (271953, +0.00 DPS) [rep]; Theramore Shoulderguards (272957, -5.50 DPS, sim-verified) [rep] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -2.74 DPS) [world]; Spritecaster Cape (11623, -2.85 DPS) [dungeon]; Gracious Cape (18743, -3.67 DPS, sim-verified) [dungeon] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Magister's Robes (16688, -0.70 DPS) [dungeon]; Vest of Elements (16666, -1.46 DPS) [dungeon]; Alanna's Embrace (13314, -5.39 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 34.9 spell_power points (10.66 DPS) | yes | Spitfire Bracers (20481, -3.59 DPS) [crafted]; Runecloth Cuffs (254123, -3.88 DPS) [crafted]; Bindings of Elements (16671, -3.90 DPS) [dungeon] |
| hands | Theramore Handguards (271950) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Theramore Handwraps (271945, -0.38 DPS) [rep]; Hands of Power (13253, -1.06 DPS) [dungeon]; Gauntlets of Elements (16672, -1.75 DPS) [dungeon] |
| waist | Stormseeker's Girdle (272399) | Pix Xizzix [vendor] | 43.0 spell_power points (13.12 DPS) | yes | Feralsurge Girdle (18104, -2.62 DPS) [dungeon]; Crystallized Girdle (12606, -2.82 DPS) [dungeon]; Sash of the Windreaver (18676, -2.82 DPS, sim-verified) [world] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 64.9 spell_power points (19.81 DPS) | yes | Leggings of Arcana (12756, -2.47 DPS, sim-verified) [quest]; Sentinel's Silk Leggings (237815, -2.54 DPS) [vendor]; Sentinel's Lizardhide Pants (237817, -2.54 DPS) [vendor] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dragonrider Boots (18102, +0.00 DPS) [dungeon]; Wolfrunner Shoes (13101, -0.59 DPS) [world_drop]; Mooncloth Boots (15802, -0.74 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.34 DPS) [quest]; Lorekeeper's Ring (19522, -1.12 DPS) [rep]; Naglering (11669, -10.10 DPS, sim-verified) [dungeon] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.18 DPS) [quest]; Lorekeeper's Ring (19522, -0.95 DPS) [rep] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Enriched Seal (271903, +0.00 DPS) [rep]; Second Wind (11819, -12.91 DPS, sim-verified) [dungeon] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -1.41 DPS) [quest]; Enriched Seal (271903, -2.30 DPS) [rep]; Serenity Field (272439, -2.44 DPS) [vendor] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Staff of Jordan (873, -2.88 DPS) [world_drop]; Ironbark Staff (20069, -7.36 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lei of the Lifegiver (19312, +0.00 DPS) [rep]; Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Barricade (235473, +0.00 DPS) [vendor] |
| ranged | Burning Totem (272433) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of Thunder (228176, -2.68 DPS, sim-verified) [vendor] |

**New at 60:** head: Blue Dragonscale Helm; neck: Lady Maye's Pendant; back: Arcanoweave Cloak; chest: Ironfeather Breastplate; wrist: Dryad's Wrist Bindings; hands: Theramore Handguards; waist: Stormseeker's Girdle; legs: Ironfeather Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Channeler's Ring; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Hand of Edward the Odd; off_hand: Therazane's Touch; ranged: Burning Totem

No-known-source sample (15 of 2086, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60, raid preset (dwarf, 5500331220001351-055000000000000000-0530020000000000)

Set DPS (verified): 731.4. Weights run: 3.8s. Verify run: 45.2s. 2086 eligible items had no known source.

Race sweep: dwarf 731.4 ± 0.6 (dps); the guide race dwarf is best.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.198 ± 0.007, crit=0.667 ± 0.014 per rating point (14 rating = 1%, 9.332 per %), hit=1.238 ± 0.076 per rating point (10 rating = 1%, 12.376 per %), spell_haste=9.517 ± 0.588, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.879 ± 0.001, spirit=not significant (0.000 ± 0.000), mp5=not significant (0.003 ± 0.009)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blue Dragonscale Helm (252604) | Leatherworking [crafted] | 40.9 spell_power points (23.10 DPS) | yes | Crimson Felt Hat (18727, -5.28 DPS) [dungeon]; Living Crown (252561, -6.98 DPS) [crafted]; Mask of the Unforgiven (13404, -7.69 DPS, sim-verified) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon]; Mark of Fordring (15411, -2.62 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 28.9 spell_power points (16.31 DPS) | yes | Wyrmhide Spaulders (12082, -2.34 DPS) [quest]; Truestrike Shoulders (12927, -2.34 DPS) [dungeon]; Argent Shoulders (19059, -4.45 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -7.53 DPS, sim-verified) [world]; Spritecaster Cape (11623, -8.34 DPS) [dungeon]; Shroud of Arcane Mastery (22330, -8.69 DPS) [dungeon] |
| chest | Vest of Elements (16666) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Tunic of the Crescent Moon (22409, -0.14 DPS) [dungeon]; Alanna's Embrace (13314, -1.69 DPS) [dungeon]; Fernpulse Jerkin (15786, -2.17 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 23.6 spell_power points (13.31 DPS) | yes | Runecloth Cuffs (254123, -5.98 DPS) [crafted]; Blackmist Armguards (12966, -6.32 DPS) [dungeon]; Slashclaw Bracers (13211, -6.32 DPS) [dungeon] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Storm Gauntlets (12632, -1.36 DPS) [crafted]; Hands of Power (13253, -3.16 DPS) [dungeon]; Voone's Vice Grips (13963, -4.53 DPS) [quest] |
| waist | Highlander's Cloth Girdle (20047) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Sash of the Windreaver (18676, +0.00 DPS) [world]; Barrage Girdle (18721, -0.19 DPS) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) (or Sentinel's Lizardhide Pants (237817)) | Illiyana Moonblaze [vendor] | 52.6 spell_power points (29.71 DPS) | yes | Sentinel's Lizardhide Pants (237817, +0.00 DPS) [vendor]; Ironfeather Leggings (252486, -0.99 DPS) [crafted]; Pristine Scorpid Leggings (252606, -5.82 DPS) [crafted] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 25.4 spell_power points (14.31 DPS) | yes | Earthen Silk Slippers (254013, -0.76 DPS) [crafted]; Omnicast Boots (11822, -1.68 DPS) [dungeon]; Stormcloth Boots (10039, -1.86 DPS) [crafted] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -1.15 DPS) [dungeon]; Maiden's Circle (13001, -1.69 DPS) [world_drop]; Naglering (11669, -14.33 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -0.78 DPS) [dungeon]; Songstone of Ironforge (12543, -1.31 DPS) [quest]; Maiden's Circle (13001, -1.31 DPS) [world_drop] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -4.69 DPS, sim-verified) [vendor]; Burst of Knowledge (11832, -5.64 DPS) [dungeon]; Uther's Strength (11302, -9.03 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Uther's Strength (11302, -2.90 DPS) [world_drop]; Shard of the Splithooves (10659, -6.27 DPS) [quest] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -3.06 DPS) [pvp]; Staff of Jordan (873, -17.46 DPS, sim-verified) [world_drop]; Kindling Stave (11750, -21.40 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Thunder (228176) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burning Totem (272433, -4.33 DPS, sim-verified) [vendor] |

**New at 60:** head: Blue Dragonscale Helm; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Vest of Elements; wrist: Dryad's Wrist Bindings; waist: Highlander's Cloth Girdle; legs: Sentinel's Silk Leggings; feet: Greaves of Withering Despair; finger1: Blessed Band of Light; finger2: Don Julio's Band; trinket1: Weakness Analyzer; trinket2: Frozen Heart of the Mountain; main_hand: Ironbark Staff; ranged: Totem of Thunder

No-known-source sample (15 of 2086, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

## Horde

### Band 20 (orc, 5300300000000000-000000000000000000-0000000000000000)

Set DPS (verified): 30.9. Weights run: 4.5s. Verify run: 9.1s. 249 eligible items had no known source.

Race sweep: orc 30.9 ± 0.0, troll 30.4 ± 0.0, tauren 29.9 ± 0.0, windshaper-skyborne 28.7 ± 0.0 (dps); the guide race orc is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.588 ± 0.007, crit=0.040 ± 0.001 per rating point (14 rating = 1%, 0.564 per %), hit=0.147 ± 0.005 per rating point (10 rating = 1%, 1.468 per %), spell_haste=-1.903 ± 0.108, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.902 ± 0.002, spirit=not significant (-0.180 ± 0.082), mp5=2.071 ± 0.126

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Hood (252448) | Leatherworking [crafted] | 9.0 spell_power points (0.95 DPS) | yes | Pristine Circlet (253949, -0.32 DPS) [crafted]; Trapper's Leather Hood (252505, -0.45 DPS) [crafted]; Wisdom's Leather Hood (252507, -0.47 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.3 spell_power points (1.08 DPS) | yes | Slime-encrusted Pads (6461, -0.43 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.66 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.72 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Black Whelp Cloak (7283, -0.11 DPS) [crafted]; Battle Healer's Cloak (20427, -0.11 DPS) [rep]; Pearl-clasped Cloak (5542, -0.41 DPS, sim-verified) [crafted] |
| chest | Stormrider's Leather Armor (252492) | Leatherworking [crafted] | 10.9 spell_power points (1.15 DPS) | yes | Filigreed Pristine Gown (253901, -0.32 DPS) [crafted]; Totemic Leather Armor (252435, -0.33 DPS) [crafted]; Wisdom's Leather Armor (252493, -0.47 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 9.1 spell_power points (0.96 DPS) | yes | Owl Bracers (4796, -0.65 DPS) [vendor]; Bright Bracers (3647, -0.71 DPS) [world_drop]; Tabitha's Cuffs (251486, -1.40 DPS, sim-verified) [quest] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.4 spell_power points (0.77 DPS) | yes | Gnoll Casting Gloves (892, -0.14 DPS) [world]; Pristine Gloves (253913, -0.17 DPS) [crafted]; Serpent Gloves (5970, -0.18 DPS, sim-verified) [dungeon] |
| waist | Pristine Sash (253925) (or Stormrider's Leather Belt (252432)) | Tailoring [crafted] | 6.4 spell_power points (0.67 DPS) | yes | Stormrider's Leather Belt (252432, +0.00 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.06 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.11 DPS) [crafted] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 13.5 spell_power points (1.42 DPS) | yes | Filigreed Pristine Leggings (253937, -0.42 DPS) [crafted]; Totemic Leather Pants (252446, -0.44 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.47 DPS, sim-verified) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.4 spell_power points (0.98 DPS) | yes | Stormrider's Leather Boots (252443, -0.04 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.25 DPS) [crafted]; Totemic Leather Boots (252442, -0.35 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 9.1 spell_power points (0.96 DPS) | yes | Loop of Sacrifice (281673, -0.65 DPS) [quest]; Volcanic Rock Ring (12053, -0.77 DPS) [world_drop]; Furlfeather Ring (286745, -0.86 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 3.5 spell_power points (0.37 DPS) | yes | Loop of Sacrifice (281673, -0.06 DPS) [quest]; Volcanic Rock Ring (12053, -0.19 DPS) [world_drop]; Furlfeather Ring (286745, -0.27 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Staff of the Blessed Seer (2271, -0.12 DPS) [dungeon]; Gnarled Necromancer's Staff (251534, -0.33 DPS) [quest]; Crescent Staff (6505, -1.08 DPS, sim-verified) [quest] |
| off_hand | Seedcloud Buckler (6630) | Wailing Caverns: Verdan the Everliving [dungeon] | 7.9 spell_power points (0.83 DPS) | yes | Worn Turtle Shell Shield (6447, -0.40 DPS) [dungeon]; Owl's Disk (4822, -0.59 DPS) [vendor]; Bristlebark Buckler (15894, -0.59 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 20:** head: Totemic Leather Hood; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Stormrider's Leather Armor; wrist: Mindthrust Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Stormrider's Leather Pants; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Evocator's Blade; off_hand: Seedcloud Buckler

No-known-source sample (15 of 249, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield; 14389 Durability Shoulderpads

### Band 30 (orc, 5500331210001000-000000000000000000-0000000000000000)

Set DPS (verified): 62.3. Weights run: 4.2s. Verify run: 10.0s. 442 eligible items had no known source.

Race sweep: orc 62.3 ± 0.1, troll 61.7 ± 0.1, tauren 61.3 ± 0.1, windshaper-skyborne 60.1 ± 0.1 (dps); the guide race orc is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.750 ± 0.010, crit=0.080 ± 0.002 per rating point (14 rating = 1%, 1.113 per %), hit=0.278 ± 0.011 per rating point (10 rating = 1%, 2.784 per %), spell_haste=-1.248 ± 0.189, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.856 ± 0.002, spirit=not significant (0.152 ± 0.093), mp5=1.742 ± 0.133

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 13.5 spell_power points (3.41 DPS) | yes | Enduring Cap (3020, -0.11 DPS) [world_drop]; Totemic Leather Helm (252456, -0.38 DPS) [crafted]; Holy Shroud (2721, -0.40 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 4.5 spell_power points (1.14 DPS) | yes | Darkspear Warding Pendant (272075, -0.19 DPS) [vendor]; Kaleidoscope Chain (13084, -0.23 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.26 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.8 spell_power points (3.97 DPS) | yes | Death Speaker Mantle (6685, -0.38 DPS) [dungeon]; Magician's Mantle (12998, -1.01 DPS) [world_drop]; Invoker's Mantle (215365, -1.26 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windsong Drape (15468, -0.21 DPS) [quest]; Hillman's Cloak (3719, -0.37 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.47 DPS) [crafted] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 19.2 spell_power points (4.85 DPS) | yes | Green Silk Armor (7065, -0.12 DPS) [crafted]; Mechbuilder's Overalls (9508, -0.56 DPS) [dungeon]; Guardian Armor (4256, -0.69 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 14.2 spell_power points (3.59 DPS) | yes | Mindthrust Bracers (1974, -1.54 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -2.07 DPS) [world_drop]; Nightsky Wristbands (6407, -2.34 DPS) [world_drop] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | 10.5 spell_power points (2.65 DPS) | yes | Jutebraid Gloves (10654, -0.19 DPS) [quest]; Stormrider's Leather Gloves (252498, -0.63 DPS) [crafted]; Truefaith Gloves (7049, -0.82 DPS) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 15.1 spell_power points (3.81 DPS) | yes | Defiler's Cloth Girdle (20164, -0.47 DPS) [rep]; Belt of Arugal (6392, -0.78 DPS) [dungeon]; Moss Cinch (6911, -0.78 DPS) [dungeon] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 15.1 spell_power points (3.81 DPS) | yes | Guardian Pants (5962, -0.28 DPS) [crafted]; Dark Ritual Leggings (270031, -0.28 DPS) [quest]; Leech Pants (6910, -0.35 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 12.9 spell_power points (3.25 DPS) | yes | Acidic Walkers (9454, -0.32 DPS) [dungeon]; Spidersilk Boots (4320, -0.72 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.78 DPS) [crafted] |
| finger1 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 11.2 spell_power points (2.83 DPS) | yes | Advisor's Ring (20426, -0.69 DPS) [rep]; Darkspear Signet (272071, -1.07 DPS) [vendor] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Advisor's Ring (19521, -0.23 DPS, sim-verified) [rep]; Darkspear Signet (272071, -0.70 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.16 DPS) [dungeon]; Tendonscraper Dagger (273053, -1.16 DPS) [world]; Glimmering Staff (249392, -2.57 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 11.5 spell_power points (2.90 DPS) | yes | Seedcloud Buckler (6630, -0.55 DPS, sim-verified) [dungeon]; Witch's Finger (16887, -1.39 DPS) [quest]; Heart of Agamaggan (6694, -1.64 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Darkspear Raider's Cloak; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; feet: Gilded Slippers; finger1: Ring of Power Regulation; finger2: Snake Hoop; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight

No-known-source sample (15 of 442, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5255 Quilboar Tomahawk

### Band 40 (orc, 5500331220001351-000000000000000000-0000000000000000)

Set DPS (verified): 108.8. Weights run: 5.3s. Verify run: 15.8s. 679 eligible items had no known source.

Race sweep: orc 108.8 ± 0.2, troll 107.6 ± 0.2, tauren 107.3 ± 0.2, windshaper-skyborne 104.2 ± 0.2 (dps); the guide race orc is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.672 ± 0.021, crit=0.203 ± 0.005 per rating point (14 rating = 1%, 2.840 per %), hit=0.358 ± 0.018 per rating point (10 rating = 1%, 3.581 per %), spell_haste=-2.772 ± 0.300, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.884 ± 0.002, spirit=not significant (-0.486 ± 0.188), mp5=not significant (-1.174 ± 0.330)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellpower Goggles Xtreme (10502, +0.00 DPS) [crafted]; Big Voodoo Mask (8201, -0.24 DPS) [crafted]; Augural Shroud (2620, -0.41 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Triune Amulet (7722, -1.05 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.36 DPS) [quest]; Darkspear Warding Pendant (272074, -1.85 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.7 spell_power points (3.79 DPS) | yes | Green Silken Shoulders (7057, -0.08 DPS) [crafted]; Bloodmage Mantle (7684, -0.17 DPS) [dungeon]; Berylline Pads (4197, -0.49 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.47 DPS) [vendor]; Icy Cloak (4327, -0.57 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamweave Vest (10021, -0.48 DPS) [crafted]; Robe of Power (7054, -0.96 DPS) [crafted]; Elemental Raiment (9434, -1.21 DPS) [world_drop] |
| wrist | Green Whelp Bracers (7386) | Leatherworking [crafted] | 13.3 spell_power points (3.21 DPS) | yes | Guardian Leather Bracers (4260, -0.80 DPS) [crafted]; Radiant Silver Bracers (4545, -0.95 DPS) [quest]; Turtle Scale Bracers (8198, -0.96 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamweave Gloves (10019, -2.93 DPS) [crafted]; Red Mageweave Gloves (10018, -3.65 DPS) [crafted]; Skycaller's Leather Gloves (252529, -3.81 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 17.1 spell_power points (4.11 DPS) | yes | Defiler's Cloth Girdle (20166, -0.09 DPS) [rep]; Skycaller's Leather Belt (252522, -0.73 DPS) [crafted]; Gilded Cord (254037, -0.89 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.1 spell_power points (5.31 DPS) | yes | Crimson Silk Pantaloons (7062, -1.28 DPS) [crafted]; Stormrider's Leather Pants (252502, -1.93 DPS) [crafted]; Dark Ritual Leggings (270031, -1.94 DPS) [quest] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.78 DPS) | yes | Skycaller's Leather Shoes (252532, -1.52 DPS) [crafted]; Skycaller's Mail Boots (252563, -1.52 DPS) [crafted]; Mender's Leather Shoes (252533, -2.72 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.4 spell_power points (3.22 DPS) | yes | Snake Hoop (6750, -0.96 DPS) [quest]; Reedknot Ring (9622, -1.53 DPS) [quest]; Ring of Power Regulation (273030, -1.77 DPS) [dungeon] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Snake Hoop (6750, +0.00 DPS) [quest]; Reedknot Ring (9622, -0.48 DPS) [quest]; Ring of Power Regulation (273030, -0.72 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -6.73 DPS) [dungeon]; Hypnotic Blade (7714, -7.09 DPS) [dungeon]; Glimmering Staff (249392, -9.01 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Green Whelp Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Staff of Jordan

No-known-source sample (15 of 679, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 50 (orc, 5500331220001351-050000000000000000-0500000000000000)

Set DPS (verified): 149.8. Weights run: 5.3s. Verify run: 17.2s. 862 eligible items had no known source.

Race sweep: orc 149.8 ± 0.2, troll 147.7 ± 0.2, tauren 147.7 ± 0.2, windshaper-skyborne 142.6 ± 0.2 (dps); the guide race orc is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.796 ± 0.015, crit=0.177 ± 0.006 per rating point (14 rating = 1%, 2.483 per %), hit=0.308 ± 0.016 per rating point (10 rating = 1%, 3.078 per %), spell_haste=-4.210 ± 0.167, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.881 ± 0.002, spirit=0.467 ± 0.043, mp5=1.405 ± 0.110

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 38.7 spell_power points (6.76 DPS) | yes | Dreamweave Circlet (10041, -0.72 DPS) [crafted]; Red Mageweave Headband (10033, -1.34 DPS, sim-verified) [crafted]; Charged Scorpid Helm (252581, -1.62 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.4 spell_power points (2.34 DPS) | yes | Mindburst Medallion (11196, -0.04 DPS) [quest]; Horizon Choker (13085, -0.06 DPS) [world_drop]; Triune Amulet (7722, -0.37 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 34.6 spell_power points (6.06 DPS) | yes | Rotgrip Mantle (17732, -1.28 DPS) [dungeon]; Kentic Amice (11624, -1.31 DPS) [dungeon]; Lead Surveyor's Mantle (11842, -1.83 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Runecloth Cloak (13860, -1.09 DPS) [crafted]; Darkspear Raider's Cloak (272076, -1.42 DPS) [vendor]; Deep Woodlands Cloak (19121, -1.51 DPS, sim-verified) [quest] |
| chest | Feathered Breastplate (8349) | Leatherworking [crafted] | 38.2 spell_power points (6.67 DPS) | yes | Acumen Robes (17775, -0.57 DPS) [quest]; Dreamweave Vest (10021, -1.13 DPS) [crafted]; Flamestrider Robes (11747, -1.31 DPS) [dungeon] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 18.0 spell_power points (3.14 DPS) | yes | Skycaller's Mail Bracers (252571, -0.01 DPS) [crafted]; Nethergeld Cuffs (254061, -0.53 DPS) [crafted]; Skycaller's Leather Bracers (252542, -1.14 DPS, sim-verified) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodfire Talons (12464, -1.16 DPS) [dungeon]; Skycaller's Leather Gauntlets (252550, -1.20 DPS) [crafted]; Elven Spirit Claws (2564, -2.88 DPS, sim-verified) [world_drop] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 25.5 spell_power points (4.47 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Ban'thok Sash (11662, -0.30 DPS) [dungeon]; Satyrmane Sash (17755, -0.63 DPS) [dungeon] |
| legs | Stone Guard's Pulsing Legplates (220847) | Lady Palanseer [vendor] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Senior Designer's Pantaloons (11841, -0.81 DPS) [dungeon]; Kilt of the Atal'ai Prophet (10807, -0.86 DPS) [dungeon]; Spellshock Leggings (9484, -1.05 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 27.3 spell_power points (4.77 DPS) | yes | Skycaller's Leather Boots (252471, -0.61 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -0.61 DPS) [crafted]; Stormcloth Boots (10039, -0.90 DPS) [crafted] |
| finger1 | Advisor's Ring (19519) | Warsong Outriders [rep] | 17.6 spell_power points (3.08 DPS) | yes | Brainlash (6440, -0.59 DPS) [dungeon]; Philanthropist's Ring (281635, -0.64 DPS) [quest]; Band of the Unicorn (7553, -0.81 DPS) [world_drop] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 16.4 spell_power points (2.87 DPS) | yes | Brainlash (6440, -0.38 DPS) [dungeon]; Philanthropist's Ring (281635, -0.43 DPS) [quest]; Band of the Unicorn (7553, -0.60 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -0.25 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.41 DPS) [quest]; Uther's Strength (11302, -1.85 DPS, sim-verified) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, -2.25 DPS) [quest]; Kindling Stave (11750, -2.30 DPS) [dungeon]; Glowing Brightwood Staff (812, -3.56 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Feathered Breastplate; wrist: Runic Leather Bracers; waist: Skycaller's Leather Waistguard; legs: Stone Guard's Pulsing Legplates; finger1: Advisor's Ring; finger2: Cyclopean Band; trinket1: Darkspear Voodoo Seal; trinket2: Frozen Heart of the Mountain

No-known-source sample (15 of 862, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 60 (orc, 5500331220001351-050000000000000000-0530520000000000)

Set DPS (verified): 272.5. Weights run: 5.1s. Verify run: 54.4s. 2043 eligible items had no known source.

Race sweep: orc 272.5 ± 0.5, windshaper-skyborne 270.4 ± 0.5, troll 269.2 ± 0.5, tauren 269.1 ± 0.5 (dps); the guide race orc is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.020 ± 0.035, crit=0.312 ± 0.009 per rating point (14 rating = 1%, 4.364 per %), hit=0.518 ± 0.033 per rating point (10 rating = 1%, 5.179 per %), spell_haste=-2.019 ± 0.380, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.862 ± 0.002, spirit=not significant (0.961 ± 0.289), mp5=1.817 ± 0.393

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 56.8 spell_power points (10.10 DPS) | yes | The Postmaster's Band (13390, -1.37 DPS) [dungeon]; Dreadmist Mask (16698, -1.39 DPS) [dungeon]; Blue Dragonscale Helm (252604, -1.47 DPS) [crafted] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | 29.0 spell_power points (5.15 DPS) | yes | Jeweled Amulet of Cainwyn (1443, -0.18 DPS) [world_drop]; Diana's Pearl Necklace (22403, -1.18 DPS) [dungeon]; Archlight Talisman (15856, -1.21 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-verified (+5.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Mail Spaulders (231659, -0.25 DPS) [pvp]; Darkspear Shoulderpads (272103, -0.63 DPS) [vendor]; Darkspear Shoulderguards (272958, -5.58 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -0.61 DPS) [dungeon]; Crystalline Threaded Cape (20697, -0.93 DPS) [world]; Cloak of Earth and Sky (279270, -3.13 DPS, sim-verified) [crafted] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | 60.2 spell_power points (10.69 DPS) | yes | Alanna's Embrace (13314, -0.09 DPS) [dungeon]; Feathered Breastplate (8349, -1.40 DPS) [crafted]; Magister's Robes (16688, -2.12 DPS) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 36.9 spell_power points (6.56 DPS) | yes | Spitfire Bracers (20481, -1.97 DPS) [crafted]; Bindings of Elements (16671, -2.23 DPS) [dungeon]; Wyrmthalak's Shackles (13958, -2.36 DPS) [quest] |
| hands | Raider Handguards (272101) | Creeg Bothunk [vendor] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Mail Gauntlets (231660, +0.00 DPS) [pvp]; Gauntlets of Elements (16672, -0.21 DPS) [dungeon]; Hands of Power (13253, -0.90 DPS, sim-verified) [dungeon] |
| waist | Stormseeker's Girdle (272399) | Pix Xizzix [vendor] | 40.5 spell_power points (7.20 DPS) | yes | General's Mail Waistband (16575, -0.95 DPS) [pvp]; Barrage Girdle (18721, -0.99 DPS) [dungeon]; Sash of the Windreaver (18676, -2.09 DPS, sim-verified) [world] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Outrider's Silk Leggings (22747, -0.66 DPS) [rep]; General's Mail Leggings (231664, -0.69 DPS) [pvp]; Leggings of Arcana (12756, -2.26 DPS, sim-verified) [quest] |
| feet | Wolfrunner Shoes (13101) | World drop [world_drop] | 34.8 spell_power points (6.18 DPS) | yes | General's Mail Sabatons (231661, +0.00 DPS) [vendor]; Mooncloth Boots (15802, -0.00 DPS) [crafted]; Dragonrider Boots (18102, -0.08 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.34 DPS) [quest]; Seal of Rivendare (13345, -1.22 DPS) [dungeon]; Naglering (11669, -9.56 DPS, sim-verified) [dungeon] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, +0.00 DPS) [quest]; Seal of Rivendare (13345, -0.55 DPS) [dungeon]; Blessed Band of Light (272407, -0.76 DPS) [vendor] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -0.08 DPS) [world_drop]; Shard of the Splithooves (10659, -0.19 DPS) [quest]; Second Wind (11819, -11.51 DPS, sim-verified) [dungeon] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Serenity Field (272439, -1.42 DPS) [vendor]; Second Wind (11819, -1.54 DPS, sim-verified) [dungeon]; Ankh of Life (1713, -1.86 DPS) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Staff of Jordan (873, -2.68 DPS) [world_drop]; Ironbark Staff (20220, -7.39 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tome of the Lost (22253, +0.00 DPS) [dungeon]; High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Barricade (235474, +0.00 DPS) [vendor] |
| ranged | Burning Totem (272433) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of Thunder (228176, -3.63 DPS, sim-verified) [vendor] |

**New at 60:** head: Living Crown; neck: Lady Maye's Pendant; back: Arcanoweave Cloak; chest: Ironfeather Breastplate; wrist: Dryad's Wrist Bindings; hands: Raider Handguards; waist: Stormseeker's Girdle; legs: Ironfeather Leggings; feet: Wolfrunner Shoes; finger1: Maiden's Circle; finger2: Channeler's Ring; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Hand of Edward the Odd; off_hand: Therazane's Touch; ranged: Burning Totem

No-known-source sample (15 of 2043, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 60, raid preset (orc, 5500331220001351-055000000000000000-0530020000000000)

Set DPS (verified): 732.9. Weights run: 3.8s. Verify run: 46.4s. 2043 eligible items had no known source.

Race sweep: windshaper-skyborne 734.9 ± 0.6, orc 732.9 ± 0.6, troll 729.1 ± 0.6, tauren 725.3 ± 0.6 (dps); windshaper-skyborne beats the guide race orc by 2.0 (0.27%) ± 0.9, beyond error.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.198 ± 0.007, crit=0.667 ± 0.014 per rating point (14 rating = 1%, 9.332 per %), hit=1.238 ± 0.076 per rating point (10 rating = 1%, 12.376 per %), spell_haste=9.517 ± 0.588, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.879 ± 0.001, spirit=not significant (0.000 ± 0.000), mp5=not significant (0.003 ± 0.009)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blue Dragonscale Helm (252604) | Leatherworking [crafted] | 40.9 spell_power points (23.10 DPS) | yes | Warlord's Mail Helm (231663, -2.62 DPS) [pvp]; Champion's Mail Helm (227162, -4.76 DPS) [pvp]; Mask of the Unforgiven (13404, -6.34 DPS, sim-verified) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon]; Mark of Fordring (15411, -2.62 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 28.9 spell_power points (16.31 DPS) | yes | Champion's Mail Spaulders (227160, +0.00 DPS) [vendor]; Warlord's Mail Spaulders (231659, +0.00 DPS) [pvp]; Argent Shoulders (19059, -4.46 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -8.04 DPS, sim-verified) [world]; Spritecaster Cape (11623, -8.34 DPS) [dungeon]; Shroud of Arcane Mastery (22330, -8.69 DPS) [dungeon] |
| chest | Stone Guard's Pulsing Breastplate (220844) | Lady Palanseer [vendor] | 32.8 spell_power points (18.48 DPS) | yes | Legionnaire's Mail Breastplate (227165, +0.00 DPS) [vendor]; Warlord's Mail Breastplate (231662, +0.00 DPS) [vendor]; Fernpulse Jerkin (15786, -5.72 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 23.6 spell_power points (13.31 DPS) | yes | Runecloth Cuffs (254123, -5.98 DPS) [crafted]; Blackmist Armguards (12966, -6.32 DPS) [dungeon]; Slashclaw Bracers (13211, -6.32 DPS) [dungeon] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Mail Gauntlets (231660, +0.00 DPS) [pvp]; Storm Gauntlets (12632, -1.36 DPS) [crafted]; Hands of Power (13253, -3.16 DPS) [dungeon] |
| waist | Defiler's Cloth Girdle (20163) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Sash of the Windreaver (18676, +0.00 DPS) [world]; Barrage Girdle (18721, -0.19 DPS) [dungeon] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 50.9 spell_power points (28.72 DPS) | yes | General's Mail Leggings (231664, -4.97 DPS) [pvp]; Pristine Scorpid Leggings (252606, -7.51 DPS, sim-verified) [crafted]; Skyshroud Leggings (13170, -8.64 DPS) [dungeon] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 25.4 spell_power points (14.31 DPS) | yes | General's Mail Sabatons (231661, -1.12 DPS) [vendor]; Omnicast Boots (11822, -1.68 DPS) [dungeon]; Earthen Silk Slippers (254013, -4.64 DPS, sim-verified) [crafted] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -1.15 DPS) [dungeon]; Maiden's Circle (13001, -1.69 DPS) [world_drop]; Naglering (11669, -16.29 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -1.31 DPS) [quest]; Maiden's Circle (13001, -1.31 DPS) [world_drop]; Don Mauricio's Band of Domination (22433, -3.19 DPS, sim-verified) [dungeon] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -4.71 DPS, sim-verified) [vendor]; Burst of Knowledge (11832, -5.64 DPS) [dungeon]; Rune of the Guard Captain (19120, -7.53 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, -1.40 DPS) [quest]; Uther's Strength (11302, -2.90 DPS) [world_drop] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -3.06 DPS) [pvp]; Staff of Jordan (873, -15.63 DPS, sim-verified) [world_drop]; Whiteout Staff (19101, -18.32 DPS) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Thunder (228176) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burning Totem (272433, -3.72 DPS, sim-verified) [vendor] |

**New at 60:** head: Blue Dragonscale Helm; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Stone Guard's Pulsing Breastplate; wrist: Dryad's Wrist Bindings; waist: Defiler's Cloth Girdle; legs: Ironfeather Leggings; feet: Greaves of Withering Despair; finger1: Blessed Band of Light; finger2: Don Julio's Band; trinket1: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Totem of Thunder

No-known-source sample (15 of 2043, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

