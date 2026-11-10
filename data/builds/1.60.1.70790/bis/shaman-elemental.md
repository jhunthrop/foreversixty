# Leveling BiS: Elemental

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 5300300000000000-000000000000000000-0000000000000000)

Set DPS (verified): 31.5. Weights run: 3.9s. Verify run: 7.0s. 286 eligible items had no known source.

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

No-known-source sample (15 of 286, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6360 Steelscale Crushfish

### Band 30 (dwarf, 5500331210001000-000000000000000000-0000000000000000)

Set DPS (verified): 62.9. Weights run: 3.8s. Verify run: 8.6s. 475 eligible items had no known source.

Race sweep: dwarf 62.9 ± 0.1 (dps); the guide race dwarf is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.702 ± 0.010, crit=0.079 ± 0.002 per rating point (14 rating = 1%, 1.113 per %), hit=0.273 ± 0.011 per rating point (10 rating = 1%, 2.733 per %), spell_haste=-1.132 ± 0.186, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.856 ± 0.002, spirit=not significant (-0.048 ± 0.089), mp5=1.398 ± 0.139

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 13.0 spell_power points (3.35 DPS) | yes | Totemic Leather Helm (252456, -0.26 DPS) [crafted]; Enduring Cap (3020, -0.46 DPS) [world_drop]; Holy Shroud (2721, -0.52 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 4.2 spell_power points (1.08 DPS) | yes | Seal of the Expedition (271920, -0.18 DPS) [rep]; Crystal Starfire Medallion (5003, -0.36 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.36 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.3 spell_power points (3.94 DPS) | yes | Death Speaker Mantle (6685, -0.41 DPS) [dungeon]; Magician's Mantle (12998, -1.03 DPS) [world_drop]; Invoker's Mantle (215365, -1.24 DPS) [crafted] |
| back | Expeditionary's Cape (271923) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, -0.16 DPS) [crafted]; Vine Pruner's Cloak (279835, -0.16 DPS) [quest]; Resilient Cape (14400, -0.36 DPS) [world_drop] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Mechbuilder's Overalls (9508, -0.26 DPS) [dungeon]; Guardian Armor (4256, -0.36 DPS) [crafted]; Green Silk Armor (7065, -0.44 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 13.2 spell_power points (3.40 DPS) | yes | Mindthrust Bracers (1974, -1.20 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -1.95 DPS) [world_drop]; Nightsky Wristbands (6407, -2.31 DPS) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 7.9 spell_power points (2.04 DPS) | yes | Stormrider's Leather Gloves (252498, -0.03 DPS) [crafted]; Truefaith Gloves (7049, -0.21 DPS) [crafted]; Gloves of Insight (9698, -0.23 DPS) [quest] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 14.2 spell_power points (3.66 DPS) | yes | Highlander's Cloth Girdle (20099, -0.28 DPS) [rep]; Moss Cinch (6911, -0.57 DPS) [dungeon]; Belt of Arugal (6392, -0.80 DPS) [dungeon] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 14.2 spell_power points (3.66 DPS) | yes | Dark Ritual Leggings (270031, -0.05 DPS) [quest]; Guardian Pants (5962, -0.15 DPS) [crafted]; Totemic Leather Leggings (252458, -0.52 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.9 spell_power points (3.07 DPS) | yes | Acidic Walkers (9454, -0.22 DPS, sim-verified) [dungeon]; Spidersilk Boots (4320, -0.54 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.62 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 10.9 spell_power points (2.81 DPS) | yes | Lorekeeper's Ring (19525, -0.29 DPS) [rep]; Runemender's Seal (273031, -0.95 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 10.2 spell_power points (2.62 DPS) | yes | Lorekeeper's Ring (19525, -0.10 DPS) [rep]; Runemender's Seal (273031, -0.77 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.03 DPS) [dungeon]; Tendonscraper Dagger (273053, -1.03 DPS) [world]; Glimmering Staff (249392, -2.33 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 11.2 spell_power points (2.89 DPS) | yes | Seedcloud Buckler (6630, -0.75 DPS, sim-verified) [dungeon]; Heart of Agamaggan (6694, -1.60 DPS) [dungeon]; Satyr's Rod (15962, -1.80 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Expeditionary's Cape; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Skycaller's Leather Belt; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight

No-known-source sample (15 of 475, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 40 (dwarf, 5500331220001351-000000000000000000-0000000000000000)

Set DPS (verified): 110.3. Weights run: 4.6s. Verify run: 12.8s. 804 eligible items had no known source.

Race sweep: dwarf 110.3 ± 0.2 (dps); the guide race dwarf is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.793 ± 0.019, crit=0.203 ± 0.006 per rating point (14 rating = 1%, 2.847 per %), hit=0.347 ± 0.018 per rating point (10 rating = 1%, 3.468 per %), spell_haste=-1.270 ± 0.233, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.885 ± 0.002, spirit=-0.790 ± 0.169, mp5=not significant (0.118 ± 0.346)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Miner's Hat of the Deep (9429) | Uldaman [dungeon] | 25.5 spell_power points (6.20 DPS) | yes | Corpseshroud (10574, -0.83 DPS) [dungeon]; Papal Fez (9431, -0.97 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -1.09 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.3 spell_power points (3.25 DPS) | yes | Triune Amulet (7722, -1.30 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.58 DPS, sim-verified) [quest]; Seal of the Expedition (271919, -1.90 DPS) [rep] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 17.3 spell_power points (4.21 DPS) | yes | Green Silken Shoulders (7057, -0.14 DPS) [crafted]; Bloodmage Mantle (7684, -0.28 DPS) [dungeon]; Berylline Pads (4197, -0.58 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Expeditionary's Cape (271922, -0.30 DPS) [rep]; Icy Cloak (4327, -0.72 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.8 spell_power points (6.51 DPS) | yes | Dreamweave Vest (10021, -0.39 DPS) [crafted]; Robe of Power (7054, -0.79 DPS) [crafted]; Elemental Raiment (9434, -1.40 DPS) [world_drop] |
| wrist | Green Whelp Bracers (7386) | Leatherworking [crafted] | 14.3 spell_power points (3.48 DPS) | yes | Guardian Leather Bracers (4260, -1.03 DPS, sim-verified) [crafted]; Turtle Scale Bracers (8198, -1.06 DPS) [crafted]; Spidertank Oilrag (9448, -1.21 DPS) [dungeon] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamweave Gloves (10019, -1.96 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -3.39 DPS) [crafted]; Skycaller's Leather Gloves (252529, -3.58 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 18.9 spell_power points (4.60 DPS) | yes | Skycaller's Leather Belt (252522, -1.01 DPS) [crafted]; Gilded Cord (254037, -1.11 DPS) [crafted]; Highlander's Cloth Girdle (20098, -1.48 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.5 spell_power points (5.72 DPS) | yes | Crimson Silk Pantaloons (7062, -0.93 DPS, sim-verified) [crafted]; Stormrider's Leather Pants (252502, -2.13 DPS) [crafted]; Guardian Pants (5962, -2.23 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.84 DPS) | yes | Skycaller's Leather Shoes (252532, -0.84 DPS, sim-verified) [crafted]; Skycaller's Mail Boots (252563, -1.33 DPS) [crafted]; Mender's Leather Shoes (252533, -2.54 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.0 spell_power points (3.40 DPS) | yes | Lorekeeper's Ring (19524, -1.12 DPS) [rep]; Ring of Forlorn Spirits (2043, -1.45 DPS) [quest]; Runemender's Seal (273031, -1.51 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 11.5 spell_power points (2.81 DPS) | yes | Lorekeeper's Ring (19524, -0.53 DPS) [rep]; Ring of Forlorn Spirits (2043, -0.86 DPS) [quest]; Runemender's Seal (273031, -0.92 DPS) [world] |
| trinket1 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hypnotic Blade (7714, -7.47 DPS) [dungeon]; Illusionary Rod (7713, -8.39 DPS, sim-verified) [dungeon]; Glimmering Staff (249392, -9.33 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Miner's Hat of the Deep; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Green Whelp Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; trinket1: Enriched Seal; trinket2: Ankh of Life; main_hand: Staff of Jordan

No-known-source sample (15 of 804, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (dwarf, 5500331220001351-050000000000000000-0500000000000000)

Set DPS (verified): 164.9. Weights run: 4.8s. Verify run: 15.8s. 1050 eligible items had no known source.

Race sweep: dwarf 164.9 ± 0.3 (dps); the guide race dwarf is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.796 ± 0.015, crit=0.177 ± 0.006 per rating point (14 rating = 1%, 2.483 per %), hit=0.308 ± 0.016 per rating point (10 rating = 1%, 3.078 per %), spell_haste=-4.210 ± 0.167, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.881 ± 0.002, spirit=0.467 ± 0.043, mp5=1.405 ± 0.110

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 38.7 spell_power points (6.76 DPS) | yes | Red Mageweave Headband (10033, -0.66 DPS) [crafted]; Dreamweave Circlet (10041, -0.72 DPS) [crafted]; Charged Scorpid Helm (252581, -1.62 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.4 spell_power points (2.34 DPS) | yes | Mindburst Medallion (11196, -0.04 DPS) [quest]; Horizon Choker (13085, -0.06 DPS) [world_drop]; Talvash's Enhancing Necklace (7673, -0.10 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 34.6 spell_power points (6.06 DPS) | yes | Lead Surveyor's Mantle (11842, -0.75 DPS) [dungeon]; Rotgrip Mantle (17732, -1.28 DPS) [dungeon]; Kentic Amice (11624, -1.31 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Runecloth Cloak (13860, -1.09 DPS) [crafted]; Expeditionary's Cape (271921, -1.42 DPS) [rep]; Big Voodoo Cloak (8216, -1.65 DPS) [crafted] |
| chest | Feathered Breastplate (8349) | Leatherworking [crafted] | 38.2 spell_power points (6.67 DPS) | yes | Dreamweave Vest (10021, -1.13 DPS) [crafted]; Flamestrider Robes (11747, -1.31 DPS) [dungeon]; Robes of Insight (940, -1.41 DPS) [world_drop] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 18.0 spell_power points (3.14 DPS) | yes | Skycaller's Leather Bracers (252542, -0.01 DPS) [crafted]; Skycaller's Mail Bracers (252571, -0.01 DPS) [crafted]; Aristocratic Cuffs (12546, -0.35 DPS) [dungeon] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Elven Spirit Claws (2564, -1.07 DPS) [world_drop]; Skycaller's Leather Gauntlets (252550, -1.20 DPS) [crafted]; Skycaller's Mail Gauntlets (252585, -1.20 DPS) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 25.5 spell_power points (4.47 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Ban'thok Sash (11662, -0.44 DPS) [dungeon]; Satyrmane Sash (17755, -0.63 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | 31.0 spell_power points (5.41 DPS) | yes | Brusslehide Leggings (17751, -0.42 DPS) [dungeon]; Senior Designer's Pantaloons (11841, -0.87 DPS) [dungeon]; Kilt of the Atal'ai Prophet (10807, -0.91 DPS) [dungeon] |
| feet | Skycaller's Leather Boots (252471) | Leatherworking [crafted] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Skycaller's Mail Sabatons (252577, +0.00 DPS) [crafted]; Stormcloth Boots (10039, -0.29 DPS) [crafted]; Earthen Silk Slippers (254013, -1.18 DPS, sim-verified) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 21.3 spell_power points (3.72 DPS) | yes | Cyclopean Band (11824, -0.85 DPS) [dungeon]; Snake Hoop (6750, -1.13 DPS) [quest]; Philanthropist's Ring (281635, -1.28 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 17.6 spell_power points (3.08 DPS) | yes | Cyclopean Band (11824, -0.21 DPS) [dungeon]; Snake Hoop (6750, -0.49 DPS) [quest]; Philanthropist's Ring (281635, -0.64 DPS) [quest] |
| trinket1 | Enriched Seal (271903) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -0.18 DPS) [world_drop]; Ankh of Life (1713, -0.25 DPS) [world_drop] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Staff of Jordan (873, -10.72 DPS, sim-verified) [world_drop] |
| off_hand | Gizlock's Hypertech Buckler (17718) | Maraudon: Tinkerer Gizlock [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Truesilver Conduit (249455, -0.39 DPS) [crafted]; Enthralled Sphere (11625, -0.70 DPS) [dungeon]; Orb of the Forgotten Seer (7685, -0.80 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Feathered Breastplate; wrist: Runic Leather Bracers; waist: Skycaller's Leather Waistguard; legs: Spellshock Leggings; feet: Skycaller's Leather Boots; finger1: Brainlash; finger2: Lorekeeper's Ring; trinket1: Enriched Seal; trinket2: Thunderbrew's Boot Flask; main_hand: Blade of Eternal Darkness; off_hand: Gizlock's Hypertech Buckler

No-known-source sample (15 of 1050, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (dwarf, 5500331220001351-050000000000000000-0530520000000000)

Set DPS (verified): 282.6. Weights run: 4.4s. Verify run: 49.0s. 2350 eligible items had no known source.

Race sweep: dwarf 282.6 ± 0.5 (dps); the guide race dwarf is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.020 ± 0.035, crit=0.312 ± 0.009 per rating point (14 rating = 1%, 4.364 per %), hit=0.518 ± 0.033 per rating point (10 rating = 1%, 5.179 per %), spell_haste=-2.019 ± 0.380, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.862 ± 0.002, spirit=not significant (0.961 ± 0.289), mp5=1.817 ± 0.393

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | The Postmaster's Band (13390) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crimson Felt Hat (18727, -0.05 DPS) [dungeon]; Dreadmist Mask (16698, -0.20 DPS) [dungeon]; Living Crown (252561, -6.08 DPS, sim-verified) [crafted] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 34.0 spell_power points (6.04 DPS) | yes | Lady Maye's Pendant (14558, -0.64 DPS) [world_drop]; Nacreous Shell Necklace (22403, -2.07 DPS) [dungeon]; Archlight Talisman (15856, -2.09 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Theramore Shoulderpads (271953, -0.63 DPS) [rep]; Theramore Shoulders (271954, -0.63 DPS) [rep]; Theramore Shoulderguards (272957, -6.57 DPS, sim-verified) [rep] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -0.61 DPS) [dungeon]; Crystalline Threaded Cape (20697, -0.93 DPS) [world]; Shroud of the Exile (15421, -3.43 DPS, sim-verified) [quest] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | 60.2 spell_power points (10.69 DPS) | yes | Magister's Robes (16688, -2.12 DPS) [dungeon]; The Postmaster's Tunic (13388, -2.16 DPS) [dungeon]; Feathered Breastplate (8349, -9.37 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 36.9 spell_power points (6.56 DPS) | yes | Bleak Howler Armguards (13208, -1.64 DPS) [dungeon]; Spitfire Bracers (20481, -1.97 DPS) [crafted]; Bindings of Elements (16671, -2.23 DPS) [dungeon] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dracorian Gauntlets (13344, -0.18 DPS) [dungeon]; Theramore Handguards (271950, -0.44 DPS) [rep]; Gauntlets of Elements (16672, -0.64 DPS) [dungeon] |
| waist | Stormseeker's Girdle (272399) | Pix Xizzix [vendor] | 40.5 spell_power points (7.20 DPS) | yes | Sash of the Windreaver (18676, -0.29 DPS) [world]; Crystallized Girdle (12606, -0.41 DPS) [dungeon]; Barrage Girdle (18721, -0.63 DPS) [dungeon] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Leggings of Arcana (12756, +0.00 DPS) [quest]; Sentinel's Silk Leggings (22752, -0.66 DPS) [rep]; Red Dragonscale Leggings (252603, -1.44 DPS) [crafted] |
| feet | The Postmaster's Treads (13391) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; Dragonrider Boots (18102, +0.00 DPS) [dungeon]; Wolfrunner Shoes (13101, -6.83 DPS, sim-verified) [world_drop] |
| finger1 | The Postmaster's Seal (13392) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Brainlash (6440, -0.28 DPS) [dungeon]; Seal of Rivendare (13345, -0.40 DPS) [dungeon]; Maiden's Circle (13001, -7.02 DPS, sim-verified) [world_drop] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-verified (+10.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Maiden's Circle (13001, +0.00 DPS) [world_drop]; Brainlash (6440, -0.20 DPS) [dungeon]; Seal of Rivendare (13345, -0.32 DPS) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+12.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Ankh of Life (1713, -0.53 DPS) [world_drop]; Shard of the Splithooves (10659, -0.65 DPS) [quest] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Eye of the Beast (13968, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironbark Staff (20069, +0.00 DPS) [rep]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Hand of Edward the Odd (2243, -3.48 DPS, sim-verified) [world_drop] |
| off_hand | Rhombeard Protector (13205) | Blackrock Spire: Gizrul the Slavener [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Therazane's Touch (19315, +0.00 DPS) [rep]; Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Barricade (235473, +0.00 DPS) [vendor] |
| ranged | Totem of Thunder (228176) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of Rage (22395, +0.00 DPS) [dungeon] |

**New at 60:** head: The Postmaster's Band; neck: Jeweled Amulet of Cainwyn; back: Arcanoweave Cloak; chest: Ironfeather Breastplate; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Stormseeker's Girdle; legs: Ironfeather Leggings; feet: The Postmaster's Treads; finger1: The Postmaster's Seal; finger2: Channeler's Ring; trinket1: Burst of Knowledge; trinket2: Second Wind; off_hand: Rhombeard Protector; ranged: Totem of Thunder

No-known-source sample (15 of 2350, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60, raid preset (dwarf, 5500331220001351-055000000000000000-0530020000000000)

Set DPS (verified): 739.0. Weights run: 3.4s. Verify run: 42.3s. 2350 eligible items had no known source.

Race sweep: dwarf 739.0 ± 0.5 (dps); the guide race dwarf is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.202 ± 0.007, crit=0.687 ± 0.015 per rating point (14 rating = 1%, 9.615 per %), hit=1.236 ± 0.076 per rating point (10 rating = 1%, 12.357 per %), spell_haste=9.427 ± 0.611, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.877 ± 0.001, spirit=not significant (0.000 ± 0.000), mp5=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blue Dragonscale Helm (252604) | Leatherworking [crafted] | 41.0 spell_power points (22.74 DPS) | yes | Crimson Felt Hat (18727, -4.98 DPS) [dungeon]; Living Crown (252561, -6.86 DPS) [crafted]; Mask of the Unforgiven (13404, -15.44 DPS, sim-verified) [dungeon] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 23.0 spell_power points (12.75 DPS) | yes | Lady Maye's Pendant (14558, -4.52 DPS) [world_drop]; Archlight Talisman (15856, -4.97 DPS) [quest]; Star of Mystaria (12103, -4.97 DPS, sim-verified) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 29.2 spell_power points (16.22 DPS) | yes | Burial Shawl (18681, -3.33 DPS) [dungeon]; Argent Shoulders (19059, -4.25 DPS, sim-verified) [crafted]; Theramore Shoulderguards (272957, -4.32 DPS) [rep] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -8.19 DPS) [dungeon]; Shroud of Arcane Mastery (22330, -8.54 DPS) [dungeon]; Crystalline Threaded Cape (20697, -10.22 DPS, sim-verified) [world] |
| chest | Fernpulse Jerkin (15786) | High Chief Winterfall [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tunic of the Crescent Moon (22409, -0.32 DPS) [dungeon]; Alanna's Embrace (13314, -0.32 DPS) [dungeon]; Vest of Elements (16666, -0.32 DPS) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 23.6 spell_power points (13.10 DPS) | yes | Runecloth Cuffs (254123, -5.88 DPS) [crafted]; Blackmist Armguards (12966, -6.25 DPS) [dungeon]; Slashclaw Bracers (13211, -6.25 DPS) [dungeon] |
| hands | Trueaim Gauntlets (13255) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of the Greatfather (17721, -2.39 DPS) [crafted]; Storm Gauntlets (12632, -3.73 DPS) [crafted]; Hands of Power (13253, -6.80 DPS) [dungeon] |
| waist | Sash of the Windreaver (18676) | The Windreaver [world] | 26.8 spell_power points (14.90 DPS) | yes | Highlander's Cloth Girdle (20047, -1.12 DPS) [rep]; Barrage Girdle (18721, -1.24 DPS) [dungeon]; Ban'thok Sash (11662, -9.63 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) (or Sentinel's Lizardhide Pants (237817)) | Illiyana Moonblaze [vendor] | 53.2 spell_power points (29.55 DPS) | yes | Sentinel's Lizardhide Pants (237817, +0.00 DPS) [vendor]; Ironfeather Leggings (252486, -1.11 DPS) [crafted]; Pristine Scorpid Leggings (252606, -5.90 DPS) [crafted] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 25.4 spell_power points (14.08 DPS) | yes | Omnicast Boots (11822, -1.64 DPS) [dungeon]; Stormcloth Boots (10039, -1.86 DPS) [crafted]; Earthen Silk Slippers (254013, -6.85 DPS, sim-verified) [crafted] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -0.97 DPS) [dungeon]; Maiden's Circle (13001, -1.63 DPS) [world_drop]; Naglering (11669, -17.63 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -1.42 DPS) [quest]; Maiden's Circle (13001, -1.42 DPS) [world_drop]; Don Mauricio's Band of Domination (22433, -5.44 DPS, sim-verified) [dungeon] |
| trinket1 | Smolderweb's Eye (13213) | Blackrock Spire: Mother Smolderweb [dungeon] | sim-verified (+12.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -0.69 DPS) [crafted]; Cannonball Runner (13382, -3.43 DPS) [dungeon]; Uther's Strength (11302, -3.53 DPS) [world_drop] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -3.82 DPS) [crafted]; Serenity Field (272439, -6.10 DPS) [vendor]; Draconic Infused Emblem (22268, -8.87 DPS, sim-verified) [dungeon] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -2.99 DPS) [pvp]; Headmaster's Charge (13937, -14.06 DPS, sim-verified) [dungeon]; Staff of Jordan (873, -15.36 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Thunder (228176) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burning Totem (272433, -2.99 DPS, sim-verified) [vendor] |

**New at 60:** head: Blue Dragonscale Helm; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Fernpulse Jerkin; wrist: Dryad's Wrist Bindings; hands: Trueaim Gauntlets; waist: Sash of the Windreaver; legs: Sentinel's Silk Leggings; feet: Greaves of Withering Despair; finger1: Blessed Band of Light; finger2: Don Julio's Band; trinket1: Smolderweb's Eye; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Totem of Thunder

No-known-source sample (15 of 2350, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

## Horde

### Band 20 (orc, 5300300000000000-000000000000000000-0000000000000000)

Set DPS (verified): 30.9. Weights run: 3.9s. Verify run: 7.9s. 275 eligible items had no known source.

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

No-known-source sample (15 of 275, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6360 Steelscale Crushfish; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield

### Band 30 (orc, 5500331210001000-000000000000000000-0000000000000000)

Set DPS (verified): 62.9. Weights run: 3.8s. Verify run: 9.3s. 469 eligible items had no known source.

Race sweep: orc 62.9 ± 0.1, tauren 62.2 ± 0.1, troll 62.1 ± 0.1, windshaper-skyborne 60.7 ± 0.1 (dps); the guide race orc is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.716 ± 0.011, crit=0.040 ± 0.001 per rating point (14 rating = 1%, 0.561 per %), hit=0.130 ± 0.006 per rating point (10 rating = 1%, 1.303 per %), spell_haste=-2.592 ± 0.109, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.877 ± 0.002, spirit=not significant (-0.104 ± 0.050), mp5=not significant (0.075 ± 0.081)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 13.2 spell_power points (2.28 DPS) | yes | Totemic Leather Helm (252456, -0.20 DPS) [crafted]; Enduring Cap (3020, -0.30 DPS) [world_drop]; Holy Shroud (2721, -0.37 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 4.3 spell_power points (0.74 DPS) | yes | Darkspear Warding Pendant (272075, -0.12 DPS) [vendor]; Crystal Starfire Medallion (5003, -0.25 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.25 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.4 spell_power points (2.68 DPS) | yes | Death Speaker Mantle (6685, -0.27 DPS) [dungeon]; Magician's Mantle (12998, -0.69 DPS) [world_drop]; Invoker's Mantle (215365, -0.84 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windsong Drape (15468, -0.13 DPS) [quest]; Hillman's Cloak (3719, -0.23 DPS, sim-verified) [crafted]; Resilient Cape (14400, -0.25 DPS) [world_drop] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Mechbuilder's Overalls (9508, -0.17 DPS) [dungeon]; Guardian Armor (4256, -0.25 DPS) [crafted]; Green Silk Armor (7065, -0.47 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.2 spell_power points (1.60 DPS) | yes | Nightsky Wristbands (6407, -0.85 DPS) [world_drop]; Mindthrust Bracers (1974, -0.94 DPS) [dungeon]; Glowing Magical Bracelets (13106, -1.93 DPS, sim-verified) [world_drop] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | 10.4 spell_power points (1.81 DPS) | yes | Jutebraid Gloves (10654, -0.15 DPS) [quest]; Stormrider's Leather Gloves (252498, -0.45 DPS) [crafted]; Truefaith Gloves (7049, -0.57 DPS) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 14.3 spell_power points (2.48 DPS) | yes | Defiler's Cloth Girdle (20164, -0.20 DPS) [rep]; Moss Cinch (6911, -0.40 DPS) [dungeon]; Belt of Arugal (6392, -0.55 DPS) [dungeon] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 14.3 spell_power points (2.48 DPS) | yes | Dark Ritual Leggings (270031, -0.05 DPS) [quest]; Guardian Pants (5962, -0.10 DPS) [crafted]; Totemic Leather Leggings (252458, -0.35 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 12.0 spell_power points (2.08 DPS) | yes | Acidic Walkers (9454, -0.22 DPS) [dungeon]; Spidersilk Boots (4320, -0.37 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.42 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 11.0 spell_power points (1.91 DPS) | yes | Ring of Power Regulation (273030, -0.83 DPS) [dungeon]; Lavishly Jeweled Ring (1156, -1.16 DPS) [dungeon]; Runemender's Seal (273031, -2.90 DPS, sim-verified) [world] |
| finger2 | Advisor's Ring (19521) | Warsong Outriders [rep] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Ring of Power Regulation (273030, -0.16 DPS) [dungeon]; Runemender's Seal (273031, -0.48 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.49 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -0.69 DPS) [dungeon]; Tendonscraper Dagger (273053, -0.69 DPS) [world]; Glimmering Staff (249392, -2.48 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 11.3 spell_power points (1.96 DPS) | yes | Seedcloud Buckler (6630, -0.68 DPS, sim-verified) [dungeon]; Witch's Finger (16887, -0.97 DPS) [quest]; Heart of Agamaggan (6694, -1.09 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Darkspear Raider's Cloak; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Advisor's Ring; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight

No-known-source sample (15 of 469, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5255 Quilboar Tomahawk

### Band 40 (orc, 5500331220001351-000000000000000000-0000000000000000)

Set DPS (verified): 109.8. Weights run: 4.6s. Verify run: 13.8s. 787 eligible items had no known source.

Race sweep: orc 109.8 ± 0.2, troll 108.8 ± 0.2, tauren 108.3 ± 0.2, windshaper-skyborne 104.1 ± 0.2 (dps); the guide race orc is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.793 ± 0.019, crit=0.203 ± 0.006 per rating point (14 rating = 1%, 2.847 per %), hit=0.347 ± 0.018 per rating point (10 rating = 1%, 3.468 per %), spell_haste=-1.270 ± 0.233, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.885 ± 0.002, spirit=-0.790 ± 0.169, mp5=not significant (0.118 ± 0.346)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Miner's Hat of the Deep (9429) | Uldaman [dungeon] | 25.5 spell_power points (6.20 DPS) | yes | Corpseshroud (10574, -0.83 DPS) [dungeon]; Papal Fez (9431, -0.97 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -1.09 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.3 spell_power points (3.25 DPS) | yes | Triune Amulet (7722, -1.30 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.46 DPS, sim-verified) [quest]; Darkspear Warding Pendant (272074, -1.90 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 17.3 spell_power points (4.21 DPS) | yes | Green Silken Shoulders (7057, -0.14 DPS) [crafted]; Bloodmage Mantle (7684, -0.28 DPS) [dungeon]; Berylline Pads (4197, -0.58 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.30 DPS) [vendor]; Icy Cloak (4327, -0.72 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.8 spell_power points (6.51 DPS) | yes | Dreamweave Vest (10021, -0.39 DPS) [crafted]; Robe of Power (7054, -0.79 DPS) [crafted]; Elemental Raiment (9434, -1.40 DPS) [world_drop] |
| wrist | Green Whelp Bracers (7386) | Leatherworking [crafted] | 14.3 spell_power points (3.48 DPS) | yes | Guardian Leather Bracers (4260, -0.86 DPS) [crafted]; Radiant Silver Bracers (4545, -0.97 DPS) [quest]; Turtle Scale Bracers (8198, -1.06 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamweave Gloves (10019, -2.08 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -3.39 DPS) [crafted]; Skycaller's Leather Gloves (252529, -3.58 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 18.9 spell_power points (4.60 DPS) | yes | Defiler's Cloth Girdle (20166, -0.42 DPS) [rep]; Skycaller's Leather Belt (252522, -1.01 DPS) [crafted]; Gilded Cord (254037, -1.11 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.5 spell_power points (5.72 DPS) | yes | Crimson Silk Pantaloons (7062, -1.27 DPS) [crafted]; Stormrider's Leather Pants (252502, -2.13 DPS) [crafted]; Guardian Pants (5962, -2.23 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.84 DPS) | yes | Skycaller's Leather Shoes (252532, -0.83 DPS, sim-verified) [crafted]; Skycaller's Mail Boots (252563, -1.33 DPS) [crafted]; Mender's Leather Shoes (252533, -2.54 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.0 spell_power points (3.40 DPS) | yes | Advisor's Ring (19520, -1.12 DPS) [rep]; Runemender's Seal (273031, -1.51 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 11.5 spell_power points (2.81 DPS) | yes | Advisor's Ring (19520, -0.53 DPS) [rep]; Runemender's Seal (273031, -0.92 DPS) [world] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hypnotic Blade (7714, -7.47 DPS) [dungeon]; Illusionary Rod (7713, -8.21 DPS, sim-verified) [dungeon]; Glimmering Staff (249392, -9.33 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Miner's Hat of the Deep; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Green Whelp Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Staff of Jordan

No-known-source sample (15 of 787, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 50 (orc, 5500331220001351-050000000000000000-0500000000000000)

Set DPS (verified): 164.5. Weights run: 4.8s. Verify run: 17.2s. 1018 eligible items had no known source.

Race sweep: orc 164.5 ± 0.3, tauren 162.5 ± 0.3, troll 162.4 ± 0.3, windshaper-skyborne 158.6 ± 0.3 (dps); the guide race orc is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.796 ± 0.015, crit=0.177 ± 0.006 per rating point (14 rating = 1%, 2.483 per %), hit=0.308 ± 0.016 per rating point (10 rating = 1%, 3.078 per %), spell_haste=-4.210 ± 0.167, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.881 ± 0.002, spirit=0.467 ± 0.043, mp5=1.405 ± 0.110

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 38.7 spell_power points (6.76 DPS) | yes | Red Mageweave Headband (10033, -0.66 DPS) [crafted]; Dreamweave Circlet (10041, -0.72 DPS) [crafted]; Charged Scorpid Helm (252581, -1.62 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.4 spell_power points (2.34 DPS) | yes | Horizon Choker (13085, -0.06 DPS) [world_drop]; Jarkal's Enhancing Necklace (7888, -0.10 DPS) [quest]; Mindburst Medallion (11196, -1.88 DPS, sim-verified) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 34.6 spell_power points (6.06 DPS) | yes | Rotgrip Mantle (17732, -1.28 DPS) [dungeon]; Kentic Amice (11624, -1.31 DPS) [dungeon]; Lead Surveyor's Mantle (11842, -2.15 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -0.42 DPS) [quest]; Runecloth Cloak (13860, -1.09 DPS) [crafted]; Darkspear Raider's Cloak (272076, -1.42 DPS) [vendor] |
| chest | Feathered Breastplate (8349) | Leatherworking [crafted] | 38.2 spell_power points (6.67 DPS) | yes | Flamestrider Robes (11747, -1.31 DPS) [dungeon]; Stone Guard's Pulsing Breastplate (220844, -1.32 DPS) [vendor]; Dreamweave Vest (10021, -1.83 DPS, sim-verified) [crafted] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 18.0 spell_power points (3.14 DPS) | yes | Skycaller's Leather Bracers (252542, -0.01 DPS) [crafted]; Skycaller's Mail Bracers (252571, -0.01 DPS) [crafted]; Aristocratic Cuffs (12546, -0.35 DPS) [dungeon] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller's Leather Gauntlets (252550, -1.20 DPS) [crafted]; Skycaller's Mail Gauntlets (252585, -1.20 DPS) [crafted]; Elven Spirit Claws (2564, -2.86 DPS, sim-verified) [world_drop] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 25.5 spell_power points (4.47 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Ban'thok Sash (11662, -0.44 DPS) [dungeon]; Satyrmane Sash (17755, -0.63 DPS) [dungeon] |
| legs | Stone Guard's Pulsing Legplates (220847) | PvP rank 6 · Stone Guard · Horde [vendor] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Brusslehide Leggings (17751, -0.36 DPS) [dungeon]; Senior Designer's Pantaloons (11841, -0.81 DPS) [dungeon]; Spellshock Leggings (9484, -1.10 DPS, sim-verified) [dungeon] |
| feet | Skycaller's Leather Boots (252471) | Leatherworking [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Skycaller's Mail Sabatons (252577, +0.00 DPS) [crafted]; Stormcloth Boots (10039, -0.29 DPS) [crafted]; Earthen Silk Slippers (254013, -0.66 DPS, sim-verified) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 21.3 spell_power points (3.72 DPS) | yes | Cyclopean Band (11824, -0.85 DPS) [dungeon]; Snake Hoop (6750, -1.13 DPS) [quest]; Philanthropist's Ring (281635, -1.28 DPS) [quest] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 17.6 spell_power points (3.08 DPS) | yes | Snake Hoop (6750, -0.49 DPS) [quest]; Philanthropist's Ring (281635, -0.64 DPS) [quest]; Cyclopean Band (11824, -1.41 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Enamored Water Spirit (20503, -0.25 DPS) [quest]; Ankh of Life (1713, -0.25 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.41 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Enamored Water Spirit (20503, -0.07 DPS) [quest]; Ankh of Life (1713, -0.07 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.95 DPS, sim-verified) [quest] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Wildstaff (20556, +0.00 DPS) [quest]; Staff of Jordan (873, -12.83 DPS, sim-verified) [world_drop] |
| off_hand | Gizlock's Hypertech Buckler (17718) | Maraudon: Tinkerer Gizlock [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Flask of Forest Mojo (19115, -0.56 DPS) [quest]; Enthralled Sphere (11625, -0.70 DPS) [dungeon]; Truesilver Conduit (249455, -2.00 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Feathered Breastplate; wrist: Runic Leather Bracers; waist: Skycaller's Leather Waistguard; legs: Stone Guard's Pulsing Legplates; feet: Skycaller's Leather Boots; finger1: Brainlash; finger2: Advisor's Ring; trinket1: Darkspear Voodoo Seal; trinket2: Uther's Strength; main_hand: Blade of Eternal Darkness; off_hand: Gizlock's Hypertech Buckler

No-known-source sample (15 of 1018, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 60 (orc, 5500331220001351-050000000000000000-0530520000000000)

Set DPS (verified): 282.5. Weights run: 4.4s. Verify run: 50.2s. 2307 eligible items had no known source.

Race sweep: windshaper-skyborne 282.6 ± 0.5, orc 282.5 ± 0.5, troll 280.8 ± 0.5, tauren 279.9 ± 0.5 (dps); windshaper-skyborne beats the guide race orc by 0.1 (0.02%) ± 0.7, inside error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.020 ± 0.035, crit=0.312 ± 0.009 per rating point (14 rating = 1%, 4.364 per %), hit=0.518 ± 0.033 per rating point (10 rating = 1%, 5.179 per %), spell_haste=-2.019 ± 0.380, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.862 ± 0.002, spirit=not significant (0.961 ± 0.289), mp5=1.817 ± 0.393

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | The Postmaster's Band (13390) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crimson Felt Hat (18727, -0.05 DPS) [dungeon]; Dreadmist Mask (16698, -0.20 DPS) [dungeon]; Living Crown (252561, -7.43 DPS, sim-verified) [crafted] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 34.0 spell_power points (6.04 DPS) | yes | Lady Maye's Pendant (14558, -0.64 DPS) [world_drop]; Nacreous Shell Necklace (22403, -2.07 DPS) [dungeon]; Archlight Talisman (15856, -2.09 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Mail Spaulders (231659, -0.25 DPS) [pvp]; Darkspear Shoulderpads (272103, -0.63 DPS) [vendor]; Darkspear Shoulderguards (272958, -6.63 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -0.61 DPS) [dungeon]; Crystalline Threaded Cape (20697, -0.93 DPS) [world]; Shroud of the Exile (15421, -2.83 DPS, sim-verified) [quest] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | 60.2 spell_power points (10.69 DPS) | yes | Magister's Robes (16688, -2.12 DPS) [dungeon]; The Postmaster's Tunic (13388, -2.16 DPS) [dungeon]; Feathered Breastplate (8349, -9.71 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 36.9 spell_power points (6.56 DPS) | yes | Bleak Howler Armguards (13208, -1.64 DPS) [dungeon]; Spitfire Bracers (20481, -1.97 DPS) [crafted]; Bindings of Elements (16671, -2.23 DPS) [dungeon] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Mail Gauntlets (231660, +0.00 DPS) [pvp]; Dracorian Gauntlets (13344, -0.18 DPS) [dungeon]; Raider Handguards (272101, -0.44 DPS) [vendor] |
| waist | Stormseeker's Girdle (272399) | Pix Xizzix [vendor] | 40.5 spell_power points (7.20 DPS) | yes | Sash of the Windreaver (18676, -0.29 DPS) [world]; Crystallized Girdle (12606, -0.41 DPS) [dungeon]; Barrage Girdle (18721, -0.63 DPS) [dungeon] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Outrider's Silk Leggings (22747, -0.66 DPS) [rep]; General's Mail Leggings (231664, -0.69 DPS) [pvp]; Leggings of Arcana (12756, -2.10 DPS, sim-verified) [quest] |
| feet | The Postmaster's Treads (13391) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; General's Mail Sabatons (231661, +0.00 DPS) [vendor]; Wolfrunner Shoes (13101, -6.96 DPS, sim-verified) [world_drop] |
| finger1 | The Postmaster's Seal (13392) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Brainlash (6440, -0.28 DPS) [dungeon]; Seal of Rivendare (13345, -0.40 DPS) [dungeon]; Maiden's Circle (13001, -6.05 DPS, sim-verified) [world_drop] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-verified (+11.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Maiden's Circle (13001, +0.00 DPS) [world_drop]; Brainlash (6440, -0.20 DPS) [dungeon]; Seal of Rivendare (13345, -0.32 DPS) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+12.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Ankh of Life (1713, -0.53 DPS) [world_drop]; Shard of the Splithooves (10659, -0.65 DPS) [quest] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Shard of the Splithooves (10659, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Edward the Odd (2243, +0.00 DPS) [world_drop]; Ironbark Staff (20220, +0.00 DPS) [rep]; High Warlord's War Staff (234549, +0.00 DPS) [pvp] |
| off_hand | Rhombeard Protector (13205) | Blackrock Spire: Gizrul the Slavener [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Therazane's Touch (19315, +0.00 DPS) [rep]; High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Barricade (235474, +0.00 DPS) [vendor] |
| ranged | Burning Totem (272433) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of Thunder (228176, +0.00 DPS) [vendor] |

**New at 60:** head: The Postmaster's Band; neck: Jeweled Amulet of Cainwyn; back: Arcanoweave Cloak; chest: Ironfeather Breastplate; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Stormseeker's Girdle; legs: Ironfeather Leggings; feet: The Postmaster's Treads; finger1: The Postmaster's Seal; finger2: Channeler's Ring; trinket1: Burst of Knowledge; trinket2: Second Wind; off_hand: Rhombeard Protector; ranged: Burning Totem

No-known-source sample (15 of 2307, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 60, raid preset (orc, 5500331220001351-055000000000000000-0530020000000000)

Set DPS (verified): 733.6. Weights run: 3.4s. Verify run: 42.5s. 2307 eligible items had no known source.

Race sweep: windshaper-skyborne 737.2 ± 0.5, orc 733.6 ± 0.5, troll 730.4 ± 0.5, tauren 726.4 ± 0.5 (dps); windshaper-skyborne beats the guide race orc by 3.6 (0.49%) ± 0.8, beyond error.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.202 ± 0.007, crit=0.687 ± 0.015 per rating point (14 rating = 1%, 9.615 per %), hit=1.236 ± 0.076 per rating point (10 rating = 1%, 12.357 per %), spell_haste=9.427 ± 0.611, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.877 ± 0.001, spirit=not significant (0.000 ± 0.000), mp5=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blue Dragonscale Helm (252604) | Leatherworking [crafted] | 41.0 spell_power points (22.74 DPS) | yes | Warlord's Mail Helm (231663, -2.41 DPS) [pvp]; Champion's Mail Helm (227162, -4.52 DPS) [pvp]; Mask of the Unforgiven (13404, -13.12 DPS, sim-verified) [dungeon] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 23.0 spell_power points (12.75 DPS) | yes | Star of Mystaria (12103, -3.79 DPS, sim-verified) [dungeon]; Lady Maye's Pendant (14558, -4.52 DPS) [world_drop]; Archlight Talisman (15856, -4.97 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 29.2 spell_power points (16.22 DPS) | yes | Champion's Mail Spaulders (227160, +0.00 DPS) [vendor]; Warlord's Mail Spaulders (231659, +0.00 DPS) [pvp]; Argent Shoulders (19059, -4.24 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -8.19 DPS) [dungeon]; Shroud of Arcane Mastery (22330, -8.54 DPS) [dungeon]; Crystalline Threaded Cape (20697, -10.50 DPS, sim-verified) [world] |
| chest | Stone Guard's Pulsing Breastplate (220844) | PvP rank 6 · Stone Guard · Horde [vendor] | 32.8 spell_power points (18.19 DPS) | yes | Legionnaire's Mail Breastplate (227165, +0.00 DPS) [vendor]; Warlord's Mail Breastplate (231662, +0.00 DPS) [vendor]; Fernpulse Jerkin (15786, -8.10 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 23.6 spell_power points (13.10 DPS) | yes | Runecloth Cuffs (254123, -5.88 DPS) [crafted]; Blackmist Armguards (12966, -6.25 DPS) [dungeon]; Slashclaw Bracers (13211, -6.25 DPS) [dungeon] |
| hands | Trueaim Gauntlets (13255) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Mail Gauntlets (231660, -1.90 DPS) [pvp]; Gloves of the Greatfather (17721, -3.50 DPS, sim-verified) [crafted]; Storm Gauntlets (12632, -3.73 DPS) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sash of the Windreaver (18676, +0.00 DPS) [world]; Defiler's Cloth Girdle (20163, -0.86 DPS) [rep]; Barrage Girdle (18721, -0.98 DPS) [dungeon] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 51.2 spell_power points (28.44 DPS) | yes | General's Mail Leggings (231664, -4.88 DPS) [pvp]; Pristine Scorpid Leggings (252606, -7.58 DPS, sim-verified) [crafted]; Legionnaire's Mail Leggings (227161, -9.55 DPS) [pvp] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 25.4 spell_power points (14.08 DPS) | yes | General's Mail Sabatons (231661, -1.08 DPS) [vendor]; Omnicast Boots (11822, -1.64 DPS) [dungeon]; Earthen Silk Slippers (254013, -7.29 DPS, sim-verified) [crafted] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -0.97 DPS) [dungeon]; Maiden's Circle (13001, -1.63 DPS) [world_drop]; Naglering (11669, -17.61 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -1.42 DPS) [quest]; Maiden's Circle (13001, -1.42 DPS) [world_drop]; Don Mauricio's Band of Domination (22433, -4.87 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+9.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Cannonball Runner (13382, -1.37 DPS) [dungeon] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, -3.13 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -3.82 DPS) [crafted]; Draconic Infused Emblem (22268, -9.83 DPS, sim-verified) [dungeon] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -2.99 DPS) [pvp]; Headmaster's Charge (13937, -13.02 DPS, sim-verified) [dungeon]; Wildstaff (20556, -14.58 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Thunder (228176) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of Rage (22395, -2.93 DPS, sim-verified) [dungeon] |

**New at 60:** head: Blue Dragonscale Helm; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Stone Guard's Pulsing Breastplate; wrist: Dryad's Wrist Bindings; hands: Trueaim Gauntlets; waist: Ban'thok Sash; legs: Ironfeather Leggings; feet: Greaves of Withering Despair; finger1: Blessed Band of Light; finger2: Don Julio's Band; trinket1: Rune of the Guard Captain; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Totem of Thunder

No-known-source sample (15 of 2307, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

