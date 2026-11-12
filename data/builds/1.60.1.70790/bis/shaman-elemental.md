# Leveling BiS: Elemental

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 5300300000000000-000000000000000000-0000000000000000)

Set DPS (verified): 31.5. Weights run: 4.8s. Verify run: 7.9s. 286 eligible items had no known source.

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

Set DPS (verified): 62.9. Weights run: 4.3s. Verify run: 9.3s. 475 eligible items had no known source.

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

Set DPS (verified): 107.3. Weights run: 5.3s. Verify run: 14.4s. 804 eligible items had no known source.

Race sweep: dwarf 107.3 ± 0.2 (dps); the guide race dwarf is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.853 ± 0.015, crit=0.195 ± 0.005 per rating point (14 rating = 1%, 2.737 per %), hit=0.322 ± 0.017 per rating point (10 rating = 1%, 3.225 per %), spell_haste=-3.045 ± 0.231, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.888 ± 0.002, spirit=1.324 ± 0.101, mp5=7.813 ± 0.251

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Miner's Hat of the Deep (9429) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Whitemane's Chapeau (7720, +0.00 DPS) [dungeon]; Papal Fez (9431, +0.00 DPS) [dungeon]; Electromagnetic Gigaflux Reactivator (9492, +0.00 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Necklace of Calisea (1714, +0.00 DPS) [dungeon]; Triune Amulet (7722, +0.00 DPS) [dungeon]; Amberglow Talisman (10824, -0.15 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slime-encrusted Pads (6461, +0.00 DPS) [dungeon]; Mantle of Doan (7712, +0.00 DPS) [dungeon]; Sheepshear Mantle (13115, +0.00 DPS) [world_drop] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Icy Cloak (4327, +0.00 DPS) [crafted]; Ceremonial Centaur Blanket (6789, +0.00 DPS) [quest]; Expeditionary's Cape (271922, +0.00 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Robe of Power (7054, +0.00 DPS) [crafted]; Dreamweave Vest (10021, +0.00 DPS) [crafted]; Deathchill Armor (10764, +0.00 DPS) [dungeon] |
| wrist | Green Whelp Bracers (7386) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mindthrust Bracers (1974, +0.00 DPS) [dungeon]; Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Enchanted Kodo Bracers (13119, +0.00 DPS) [world_drop] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonefingers (10765, +0.00 DPS) [dungeon]; Dreamweave Gloves (10019, -0.55 DPS) [crafted]; Gloves of Old (9395, -3.19 DPS) [world_drop] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller's Leather Belt (252522, +0.00 DPS) [crafted]; Gilded Cord (254037, +0.00 DPS) [crafted]; Mender's Leather Belt (252523, -0.18 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Leech Pants (6910, +0.00 DPS) [dungeon]; Warchief Kilt (7760, +0.00 DPS) [dungeon]; Stoneweaver Leggings (9407, +0.00 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Furen's Boots (13100, +0.00 DPS) [world_drop]; Skycaller's Leather Shoes (252532, +0.00 DPS) [crafted]; Skycaller's Mail Boots (252563, +0.00 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lorekeeper's Ring (19524, +0.00 DPS) [rep]; Ring of Power Regulation (273030, +0.00 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lorekeeper's Ring (19524, +0.00 DPS) [rep]; Ring of Power Regulation (273030, +0.00 DPS) [dungeon] |
| trinket1 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -8.54 DPS) [dungeon]; Hypnotic Blade (7714, -10.13 DPS) [dungeon]; Soulstaff (249393, -11.69 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Miner's Hat of the Deep; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Green Whelp Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; trinket1: Enriched Seal; trinket2: Ankh of Life; main_hand: Staff of Jordan

No-known-source sample (15 of 804, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (dwarf, 5500331220001351-050000000000000000-0500000000000000)

Set DPS (verified): 166.6. Weights run: 5.4s. Verify run: 14.7s. 1050 eligible items had no known source.

Race sweep: dwarf 166.6 ± 0.3 (dps); the guide race dwarf is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.816 ± 0.015, crit=0.181 ± 0.006 per rating point (14 rating = 1%, 2.535 per %), hit=0.313 ± 0.015 per rating point (10 rating = 1%, 3.126 per %), spell_haste=-4.296 ± 0.168, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.881 ± 0.002, spirit=0.480 ± 0.044, mp5=1.446 ± 0.109

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 39.1 spell_power points (6.83 DPS) | yes | Red Mageweave Headband (10033, -0.66 DPS) [crafted]; Dreamweave Circlet (10041, -0.73 DPS) [crafted]; Charged Scorpid Helm (252581, -1.61 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.5 spell_power points (2.37 DPS) | yes | Horizon Choker (13085, -0.03 DPS) [world_drop]; Mindburst Medallion (11196, -0.04 DPS) [quest]; Talvash's Enhancing Necklace (7673, -0.13 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 35.2 spell_power points (6.15 DPS) | yes | Rotgrip Mantle (17732, -1.31 DPS) [dungeon]; Kentic Amice (11624, -1.34 DPS) [dungeon]; Lead Surveyor's Mantle (11842, -1.99 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Expeditionary's Cape (271921, -1.39 DPS) [rep]; Big Voodoo Cloak (8216, -1.65 DPS) [crafted]; Runecloth Cloak (13860, -1.75 DPS, sim-verified) [crafted] |
| chest | Feathered Breastplate (8349) | Leatherworking [crafted] | 38.7 spell_power points (6.76 DPS) | yes | Dreamweave Vest (10021, -1.16 DPS) [crafted]; Flamestrider Robes (11747, -1.35 DPS) [dungeon]; Embrace of the Wind Serpent (12462, -2.37 DPS, sim-verified) [world] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 18.2 spell_power points (3.18 DPS) | yes | Skycaller's Leather Bracers (252542, -0.01 DPS) [crafted]; Skycaller's Mail Bracers (252571, -0.01 DPS) [crafted]; Aristocratic Cuffs (12546, -0.34 DPS) [dungeon] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller's Leather Gauntlets (252550, -1.15 DPS) [crafted]; Skycaller's Mail Gauntlets (252585, -1.15 DPS) [crafted]; Elven Spirit Claws (2564, -2.87 DPS, sim-verified) [world_drop] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 25.8 spell_power points (4.51 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Ban'thok Sash (11662, -0.44 DPS) [dungeon]; Satyrmane Sash (17755, -0.64 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | 31.2 spell_power points (5.45 DPS) | yes | Brusslehide Leggings (17751, -0.43 DPS) [dungeon]; Senior Designer's Pantaloons (11841, -0.84 DPS) [dungeon]; Kilt of the Atal'ai Prophet (10807, -0.84 DPS) [dungeon] |
| feet | Skycaller's Leather Boots (252471) (or Skycaller's Mail Sabatons (252577)) | Leatherworking [crafted] | 24.0 spell_power points (4.19 DPS) | yes | Skycaller's Mail Sabatons (252577, +0.00 DPS) [crafted]; Greaves of Withering Despair (22240, -0.30 DPS) [dungeon]; Stormcloth Boots (10039, -0.33 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 21.6 spell_power points (3.77 DPS) | yes | Cyclopean Band (11824, -0.86 DPS) [dungeon]; Snake Hoop (6750, -1.13 DPS) [quest]; Woodseed Hoop (17768, -1.29 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 17.8 spell_power points (3.11 DPS) | yes | Cyclopean Band (11824, -0.20 DPS) [dungeon]; Snake Hoop (6750, -0.47 DPS) [quest]; Woodseed Hoop (17768, -0.63 DPS) [quest] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.38 DPS, sim-verified) [quest] |
| trinket2 | Enriched Seal (271903) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.21 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.42 DPS) [quest]; Ankh of Life (1713, -2.28 DPS, sim-verified) [world_drop] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Staff of Jordan (873, -11.79 DPS, sim-verified) [world_drop] |
| off_hand | Gizlock's Hypertech Buckler (17718) | Maraudon: Tinkerer Gizlock [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Enthralled Sphere (11625, -0.73 DPS) [dungeon]; Orb of Lorica (11262, -0.84 DPS) [quest]; Truesilver Conduit (249455, -2.41 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Feathered Breastplate; wrist: Runic Leather Bracers; waist: Skycaller's Leather Waistguard; legs: Spellshock Leggings; feet: Skycaller's Leather Boots; finger1: Brainlash; finger2: Lorekeeper's Ring; trinket1: Thunderbrew's Boot Flask; trinket2: Enriched Seal; main_hand: Blade of Eternal Darkness; off_hand: Gizlock's Hypertech Buckler

No-known-source sample (15 of 1050, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (dwarf, 5500331220001351-050000000000000000-0530520000000000)

Set DPS (verified): 287.7. Weights run: 5.0s. Verify run: 49.6s. 2374 eligible items had no known source.

Race sweep: dwarf 287.7 ± 0.5 (dps); the guide race dwarf is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.050 ± 0.035, crit=0.319 ± 0.009 per rating point (14 rating = 1%, 4.461 per %), hit=0.533 ± 0.033 per rating point (10 rating = 1%, 5.326 per %), spell_haste=-2.168 ± 0.379, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.862 ± 0.002, spirit=not significant (0.981 ± 0.286), mp5=1.858 ± 0.391

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | The Postmaster's Band (13390) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellweaver's Turban (22267, +0.00 DPS) [dungeon]; Crimson Felt Hat (18727, -0.13 DPS) [dungeon]; Living Crown (252561, -6.02 DPS, sim-verified) [crafted] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 34.7 spell_power points (6.17 DPS) | yes | Lady Maye's Pendant (14558, -0.67 DPS) [world_drop]; Nacreous Shell Necklace (22403, -2.13 DPS) [dungeon]; Archlight Talisman (15856, -2.17 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Theramore Shoulderpads (271953, -0.62 DPS) [rep]; Theramore Shoulders (271954, -0.62 DPS) [rep]; Theramore Shoulderguards (272957, -6.50 DPS, sim-verified) [rep] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -0.63 DPS) [dungeon]; Crystalline Threaded Cape (20697, -0.98 DPS) [world]; Shroud of the Exile (15421, -3.36 DPS, sim-verified) [quest] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | 61.1 spell_power points (10.85 DPS) | yes | Embrace of the Wind Serpent (12462, -1.80 DPS) [world]; Magister's Robes (16688, -2.10 DPS) [dungeon]; Feathered Breastplate (8349, -9.32 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 37.3 spell_power points (6.62 DPS) | yes | Bleak Howler Armguards (13208, -1.61 DPS) [dungeon]; Spitfire Bracers (20481, -1.95 DPS) [crafted]; Bindings of Elements (16671, -2.21 DPS) [dungeon] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dracorian Gauntlets (13344, -0.16 DPS) [dungeon]; Theramore Handguards (271950, -0.40 DPS) [rep]; Gauntlets of Elements (16672, -0.59 DPS) [dungeon] |
| waist | Stormseeker's Girdle (272399) | Pix Xizzix [vendor] | 41.2 spell_power points (7.32 DPS) | yes | Sash of the Windreaver (18676, -0.35 DPS) [world]; Crystallized Girdle (12606, -0.42 DPS) [dungeon]; Barrage Girdle (18721, -0.70 DPS) [dungeon] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Leggings of Arcana (12756, +0.00 DPS) [quest]; Sentinel's Silk Leggings (22752, -0.64 DPS) [rep]; Red Dragonscale Leggings (252603, -1.38 DPS) [crafted] |
| feet | The Postmaster's Treads (13391) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wolfrunner Shoes (13101, +0.00 DPS) [world_drop]; Dragonrider Boots (18102, +0.00 DPS) [dungeon]; Mooncloth Boots (15802, -8.28 DPS, sim-verified) [crafted] |
| finger1 | The Postmaster's Seal (13392) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Brainlash (6440, -0.27 DPS) [dungeon]; Seal of Rivendare (13345, -0.37 DPS) [dungeon]; Maiden's Circle (13001, -6.95 DPS, sim-verified) [world_drop] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-verified (+10.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Maiden's Circle (13001, +0.00 DPS) [world_drop]; Brainlash (6440, -0.26 DPS) [dungeon]; Seal of Rivendare (13345, -0.36 DPS) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+12.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Ankh of Life (1713, -0.55 DPS) [world_drop]; Shard of the Splithooves (10659, -0.66 DPS) [quest] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Eye of the Beast (13968, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironbark Staff (20069, +0.00 DPS) [rep]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Hand of Edward the Odd (2243, -3.50 DPS, sim-verified) [world_drop] |
| off_hand | Rhombeard Protector (13205) | Blackrock Spire: Gizrul the Slavener [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Therazane's Touch (19315, +0.00 DPS) [rep]; Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Barricade (235473, +0.00 DPS) [vendor] |
| ranged | Totem of Thunder (228176) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of Rage (22395, +0.00 DPS) [dungeon] |

**New at 60:** head: The Postmaster's Band; neck: Jeweled Amulet of Cainwyn; back: Arcanoweave Cloak; chest: Ironfeather Breastplate; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Stormseeker's Girdle; legs: Ironfeather Leggings; feet: The Postmaster's Treads; finger1: The Postmaster's Seal; finger2: Channeler's Ring; trinket1: Burst of Knowledge; trinket2: Second Wind; off_hand: Rhombeard Protector; ranged: Totem of Thunder

No-known-source sample (15 of 2374, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60, raid preset (dwarf, 5500331220001351-055000000000000000-0530020000000000)

Set DPS (verified): 748.7. Weights run: 4.0s. Verify run: 40.6s. 2374 eligible items had no known source.

Race sweep: dwarf 748.7 ± 0.6 (dps); the guide race dwarf is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.206 ± 0.007, crit=0.701 ± 0.015 per rating point (14 rating = 1%, 9.813 per %), hit=1.272 ± 0.076 per rating point (10 rating = 1%, 12.718 per %), spell_haste=9.709 ± 0.612, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.877 ± 0.001, spirit=not significant (0.000 ± 0.000), mp5=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 45.7 spell_power points (25.37 DPS) | yes | Blue Dragonscale Helm (252604, -2.38 DPS) [crafted]; Mask of the Unforgiven (13404, -5.81 DPS) [dungeon]; Crimson Felt Hat (18727, -7.58 DPS) [dungeon] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 23.4 spell_power points (12.97 DPS) | yes | Emberfury Talisman (12929, -4.19 DPS) [dungeon]; Lady Maye's Pendant (14558, -4.69 DPS) [world_drop]; Star of Mystaria (12103, -5.99 DPS, sim-verified) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 29.5 spell_power points (16.36 DPS) | yes | Burial Shawl (18681, -3.44 DPS) [dungeon]; Theramore Shoulderguards (272957, -4.40 DPS) [rep]; Argent Shoulders (19059, -7.72 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -8.40 DPS) [dungeon]; Shroud of Arcane Mastery (22330, -8.54 DPS) [dungeon]; Crystalline Threaded Cape (20697, -11.66 DPS, sim-verified) [world] |
| chest | Tunic of the Crescent Moon (22409) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Alanna's Embrace (13314, -0.09 DPS) [dungeon]; Vest of Elements (16666, -0.09 DPS) [dungeon]; Fernpulse Jerkin (15786, -1.28 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 23.6 spell_power points (13.12 DPS) | yes | Runecloth Cuffs (254123, -5.88 DPS) [crafted]; Blackmist Armguards (12966, -6.06 DPS) [dungeon]; Slashclaw Bracers (13211, -6.06 DPS) [dungeon] |
| hands | Trueaim Gauntlets (13255) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of the Greatfather (17721, -2.99 DPS) [crafted]; Storm Gauntlets (12632, -4.39 DPS, sim-verified) [crafted]; Hands of Power (13253, -7.38 DPS) [dungeon] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sash of the Windreaver (18676, +0.00 DPS) [world]; Highlander's Cloth Girdle (20047, -0.96 DPS) [rep]; Barrage Girdle (18721, -1.18 DPS) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) (or Sentinel's Lizardhide Pants (237817)) | Illiyana Moonblaze [vendor] | 53.7 spell_power points (29.79 DPS) | yes | Sentinel's Lizardhide Pants (237817, +0.00 DPS) [vendor]; Ironfeather Leggings (252486, -1.20 DPS) [crafted]; Pristine Scorpid Leggings (252606, -6.03 DPS) [crafted] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 25.8 spell_power points (14.30 DPS) | yes | Stormcloth Boots (10039, -2.08 DPS) [crafted]; Venomspew Footpads (275606, -2.10 DPS) [crafted]; Omnicast Boots (11822, -6.57 DPS, sim-verified) [dungeon] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -1.06 DPS) [dungeon]; Maiden's Circle (13001, -1.82 DPS) [world_drop]; Naglering (11669, -18.59 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -1.71 DPS) [quest]; Maiden's Circle (13001, -1.71 DPS) [world_drop]; Don Mauricio's Band of Domination (22433, -6.17 DPS, sim-verified) [dungeon] |
| trinket1 | Smolderweb's Eye (13213) | Blackrock Spire: Mother Smolderweb [dungeon] | sim-verified (+12.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -0.71 DPS) [crafted]; Cannonball Runner (13382, -3.53 DPS) [dungeon]; Uther's Strength (11302, -3.73 DPS) [world_drop] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -3.64 DPS) [crafted]; Serenity Field (272439, -6.10 DPS) [vendor]; Draconic Infused Emblem (22268, -8.73 DPS, sim-verified) [dungeon] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -2.97 DPS) [pvp]; Headmaster's Charge (13937, -15.01 DPS, sim-verified) [dungeon]; Staff of Jordan (873, -15.27 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Thunder (228176) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burning Totem (272433, +0.00 DPS) [vendor] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Tunic of the Crescent Moon; wrist: Dryad's Wrist Bindings; hands: Trueaim Gauntlets; waist: Ban'thok Sash; legs: Sentinel's Silk Leggings; feet: Greaves of Withering Despair; finger1: Blessed Band of Light; finger2: Don Julio's Band; trinket1: Smolderweb's Eye; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Totem of Thunder

No-known-source sample (15 of 2374, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

## Horde

### Band 20 (orc, 5300300000000000-000000000000000000-0000000000000000)

Set DPS (verified): 30.9. Weights run: 4.8s. Verify run: 8.9s. 275 eligible items had no known source.

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

Set DPS (verified): 62.9. Weights run: 4.3s. Verify run: 10.2s. 469 eligible items had no known source.

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

Set DPS (verified): 106.7. Weights run: 5.3s. Verify run: 15.2s. 787 eligible items had no known source.

Race sweep: orc 106.7 ± 0.2, troll 105.8 ± 0.2, tauren 105.4 ± 0.2, windshaper-skyborne 101.3 ± 0.2 (dps); the guide race orc is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.853 ± 0.015, crit=0.195 ± 0.005 per rating point (14 rating = 1%, 2.737 per %), hit=0.322 ± 0.017 per rating point (10 rating = 1%, 3.225 per %), spell_haste=-3.045 ± 0.231, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.888 ± 0.002, spirit=1.324 ± 0.101, mp5=7.813 ± 0.251

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Miner's Hat of the Deep (9429) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Whitemane's Chapeau (7720, +0.00 DPS) [dungeon]; Papal Fez (9431, +0.00 DPS) [dungeon]; Electromagnetic Gigaflux Reactivator (9492, +0.00 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Necklace of Calisea (1714, +0.00 DPS) [dungeon]; Triune Amulet (7722, +0.00 DPS) [dungeon]; Amberglow Talisman (10824, -0.15 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slime-encrusted Pads (6461, +0.00 DPS) [dungeon]; Mantle of Doan (7712, +0.00 DPS) [dungeon]; Sheepshear Mantle (13115, +0.00 DPS) [world_drop] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Icy Cloak (4327, +0.00 DPS) [crafted]; Cloak of Blight (6832, +0.00 DPS) [quest]; Darkspear Raider's Cloak (272077, +0.00 DPS) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Robe of Power (7054, +0.00 DPS) [crafted]; Dreamweave Vest (10021, +0.00 DPS) [crafted]; Deathchill Armor (10764, +0.00 DPS) [dungeon] |
| wrist | Green Whelp Bracers (7386) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mindthrust Bracers (1974, +0.00 DPS) [dungeon]; Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Enchanted Kodo Bracers (13119, +0.00 DPS) [world_drop] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonefingers (10765, +0.00 DPS) [dungeon]; Dreamweave Gloves (10019, -0.55 DPS) [crafted]; Gloves of Old (9395, -3.19 DPS) [world_drop] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller's Leather Belt (252522, +0.00 DPS) [crafted]; Gilded Cord (254037, +0.00 DPS) [crafted]; Mender's Leather Belt (252523, -0.18 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Leech Pants (6910, +0.00 DPS) [dungeon]; Warchief Kilt (7760, +0.00 DPS) [dungeon]; Stoneweaver Leggings (9407, +0.00 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Furen's Boots (13100, +0.00 DPS) [world_drop]; Skycaller's Leather Shoes (252532, +0.00 DPS) [crafted]; Skycaller's Mail Boots (252563, +0.00 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Advisor's Ring (19520, +0.00 DPS) [rep]; Darkspear Signet (272070, +0.00 DPS) [vendor] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Advisor's Ring (19520, +0.00 DPS) [rep]; Darkspear Signet (272070, +0.00 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -8.54 DPS) [dungeon]; Hypnotic Blade (7714, -10.13 DPS) [dungeon]; Soulstaff (249393, -11.69 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Miner's Hat of the Deep; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Green Whelp Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Staff of Jordan

No-known-source sample (15 of 787, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 50 (orc, 5500331220001351-050000000000000000-0500000000000000)

Set DPS (verified): 166.8. Weights run: 5.4s. Verify run: 16.8s. 1018 eligible items had no known source.

Race sweep: orc 166.8 ± 0.3, tauren 164.7 ± 0.3, troll 164.7 ± 0.3, windshaper-skyborne 160.9 ± 0.3 (dps); the guide race orc is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.816 ± 0.015, crit=0.181 ± 0.006 per rating point (14 rating = 1%, 2.535 per %), hit=0.313 ± 0.015 per rating point (10 rating = 1%, 3.126 per %), spell_haste=-4.296 ± 0.168, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.881 ± 0.002, spirit=0.480 ± 0.044, mp5=1.446 ± 0.109

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 39.1 spell_power points (6.83 DPS) | yes | Red Mageweave Headband (10033, -0.66 DPS) [crafted]; Dreamweave Circlet (10041, -0.73 DPS) [crafted]; Charged Scorpid Helm (252581, -1.61 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.5 spell_power points (2.37 DPS) | yes | Horizon Choker (13085, -0.03 DPS) [world_drop]; Mindburst Medallion (11196, -0.04 DPS) [quest]; Jarkal's Enhancing Necklace (7888, -0.13 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 35.2 spell_power points (6.15 DPS) | yes | Rotgrip Mantle (17732, -1.31 DPS) [dungeon]; Kentic Amice (11624, -1.34 DPS) [dungeon]; Lead Surveyor's Mantle (11842, -1.78 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -0.42 DPS) [quest]; Runecloth Cloak (13860, -1.09 DPS) [crafted]; Darkspear Raider's Cloak (272076, -1.39 DPS) [vendor] |
| chest | Feathered Breastplate (8349) | Leatherworking [crafted] | 38.7 spell_power points (6.76 DPS) | yes | Dreamweave Vest (10021, -1.16 DPS) [crafted]; Flamestrider Robes (11747, -1.35 DPS) [dungeon]; Embrace of the Wind Serpent (12462, -1.58 DPS, sim-verified) [world] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 18.2 spell_power points (3.18 DPS) | yes | Skycaller's Leather Bracers (252542, -0.01 DPS) [crafted]; Skycaller's Mail Bracers (252571, -0.01 DPS) [crafted]; Aristocratic Cuffs (12546, -0.34 DPS) [dungeon] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller's Leather Gauntlets (252550, -1.15 DPS) [crafted]; Skycaller's Mail Gauntlets (252585, -1.15 DPS) [crafted]; Elven Spirit Claws (2564, -1.88 DPS, sim-verified) [world_drop] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 25.8 spell_power points (4.51 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Ban'thok Sash (11662, -0.44 DPS) [dungeon]; Satyrmane Sash (17755, -0.64 DPS) [dungeon] |
| legs | Stone Guard's Pulsing Legplates (220847) | PvP rank 6 · Stone Guard · Horde [vendor] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Brusslehide Leggings (17751, -0.39 DPS) [dungeon]; Senior Designer's Pantaloons (11841, -0.80 DPS) [dungeon]; Spellshock Leggings (9484, -1.12 DPS, sim-verified) [dungeon] |
| feet | Skycaller's Leather Boots (252471) (or Skycaller's Mail Sabatons (252577)) | Leatherworking [crafted] | 24.0 spell_power points (4.19 DPS) | yes | Skycaller's Mail Sabatons (252577, +0.00 DPS) [crafted]; Greaves of Withering Despair (22240, -0.30 DPS) [dungeon]; Stormcloth Boots (10039, -0.33 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 21.6 spell_power points (3.77 DPS) | yes | Cyclopean Band (11824, -0.86 DPS) [dungeon]; Snake Hoop (6750, -1.13 DPS) [quest]; Woodseed Hoop (17768, -1.29 DPS) [quest] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 17.8 spell_power points (3.11 DPS) | yes | Cyclopean Band (11824, -0.20 DPS) [dungeon]; Snake Hoop (6750, -0.47 DPS) [quest]; Woodseed Hoop (17768, -0.63 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.21 DPS) [world_drop]; Ankh of Life (1713, -0.26 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.42 DPS) [quest] |
| trinket2 | Enamored Water Spirit (20503) | Da Voodoo [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Ankh of Life (1713, -0.00 DPS) [world_drop]; Alchemist's Stone (13503, -2.40 DPS, sim-verified) [crafted] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Wildstaff (20556, +0.00 DPS) [quest]; Staff of Jordan (873, -12.56 DPS, sim-verified) [world_drop] |
| off_hand | Gizlock's Hypertech Buckler (17718) | Maraudon: Tinkerer Gizlock [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Truesilver Conduit (249455, -0.42 DPS) [crafted]; Flask of Forest Mojo (19115, -0.60 DPS) [quest]; Enthralled Sphere (11625, -0.73 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Feathered Breastplate; wrist: Runic Leather Bracers; waist: Skycaller's Leather Waistguard; legs: Stone Guard's Pulsing Legplates; feet: Skycaller's Leather Boots; finger1: Brainlash; finger2: Advisor's Ring; trinket1: Darkspear Voodoo Seal; trinket2: Enamored Water Spirit; main_hand: Blade of Eternal Darkness; off_hand: Gizlock's Hypertech Buckler

No-known-source sample (15 of 1018, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 60 (orc, 5500331220001351-050000000000000000-0530520000000000)

Set DPS (verified): 287.6. Weights run: 5.0s. Verify run: 51.6s. 2331 eligible items had no known source.

Race sweep: windshaper-skyborne 287.8 ± 0.5, orc 287.6 ± 0.5, troll 285.9 ± 0.5, tauren 285.0 ± 0.5 (dps); windshaper-skyborne beats the guide race orc by 0.2 (0.07%) ± 0.7, inside error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.050 ± 0.035, crit=0.319 ± 0.009 per rating point (14 rating = 1%, 4.461 per %), hit=0.533 ± 0.033 per rating point (10 rating = 1%, 5.326 per %), spell_haste=-2.168 ± 0.379, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.862 ± 0.002, spirit=not significant (0.981 ± 0.286), mp5=1.858 ± 0.391

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | The Postmaster's Band (13390) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellweaver's Turban (22267, +0.00 DPS) [dungeon]; Crimson Felt Hat (18727, -0.13 DPS) [dungeon]; Living Crown (252561, -7.52 DPS, sim-verified) [crafted] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 34.7 spell_power points (6.17 DPS) | yes | Lady Maye's Pendant (14558, -0.67 DPS) [world_drop]; Nacreous Shell Necklace (22403, -2.13 DPS) [dungeon]; Archlight Talisman (15856, -2.17 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Mail Spaulders (231659, -0.28 DPS) [pvp]; Darkspear Shoulderpads (272103, -0.62 DPS) [vendor]; Darkspear Shoulderguards (272958, -6.64 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -0.63 DPS) [dungeon]; Crystalline Threaded Cape (20697, -0.98 DPS) [world]; Shroud of the Exile (15421, -2.87 DPS, sim-verified) [quest] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | 61.1 spell_power points (10.85 DPS) | yes | Embrace of the Wind Serpent (12462, -1.80 DPS) [world]; Magister's Robes (16688, -2.10 DPS) [dungeon]; Feathered Breastplate (8349, -9.73 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 37.3 spell_power points (6.62 DPS) | yes | Bleak Howler Armguards (13208, -1.61 DPS) [dungeon]; Spitfire Bracers (20481, -1.95 DPS) [crafted]; Bindings of Elements (16671, -2.21 DPS) [dungeon] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Mail Gauntlets (231660, +0.00 DPS) [pvp]; Dracorian Gauntlets (13344, -0.16 DPS) [dungeon]; Raider Handguards (272101, -0.40 DPS) [vendor] |
| waist | Stormseeker's Girdle (272399) | Pix Xizzix [vendor] | 41.2 spell_power points (7.32 DPS) | yes | Sash of the Windreaver (18676, -0.35 DPS) [world]; Crystallized Girdle (12606, -0.42 DPS) [dungeon]; Barrage Girdle (18721, -0.70 DPS) [dungeon] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Leggings of Arcana (12756, +0.00 DPS) [quest]; Outrider's Silk Leggings (22747, -0.64 DPS) [rep]; General's Mail Leggings (231664, -0.66 DPS) [pvp] |
| feet | The Postmaster's Treads (13391) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wolfrunner Shoes (13101, +0.00 DPS) [world_drop]; General's Mail Sabatons (231661, +0.00 DPS) [vendor]; Mooncloth Boots (15802, -7.83 DPS, sim-verified) [crafted] |
| finger1 | The Postmaster's Seal (13392) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Brainlash (6440, -0.27 DPS) [dungeon]; Seal of Rivendare (13345, -0.37 DPS) [dungeon]; Maiden's Circle (13001, -6.03 DPS, sim-verified) [world_drop] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-verified (+11.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Maiden's Circle (13001, +0.00 DPS) [world_drop]; Brainlash (6440, -0.26 DPS) [dungeon]; Seal of Rivendare (13345, -0.36 DPS) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+12.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Ankh of Life (1713, -0.55 DPS) [world_drop]; Shard of the Splithooves (10659, -0.66 DPS) [quest] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Shard of the Splithooves (10659, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Edward the Odd (2243, +0.00 DPS) [world_drop]; Ironbark Staff (20220, +0.00 DPS) [rep]; High Warlord's War Staff (234549, +0.00 DPS) [pvp] |
| off_hand | Rhombeard Protector (13205) | Blackrock Spire: Gizrul the Slavener [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Therazane's Touch (19315, +0.00 DPS) [rep]; High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Barricade (235474, +0.00 DPS) [vendor] |
| ranged | Burning Totem (272433) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of Thunder (228176, +0.00 DPS) [vendor] |

**New at 60:** head: The Postmaster's Band; neck: Jeweled Amulet of Cainwyn; back: Arcanoweave Cloak; chest: Ironfeather Breastplate; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Stormseeker's Girdle; legs: Ironfeather Leggings; feet: The Postmaster's Treads; finger1: The Postmaster's Seal; finger2: Channeler's Ring; trinket1: Burst of Knowledge; trinket2: Second Wind; off_hand: Rhombeard Protector; ranged: Burning Totem

No-known-source sample (15 of 2331, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 60, raid preset (orc, 5500331220001351-055000000000000000-0530020000000000)

Set DPS (verified): 750.4. Weights run: 4.0s. Verify run: 42.4s. 2331 eligible items had no known source.

Race sweep: windshaper-skyborne 753.5 ± 0.6, orc 750.4 ± 0.6, troll 747.1 ± 0.6, tauren 742.9 ± 0.6 (dps); windshaper-skyborne beats the guide race orc by 3.1 (0.41%) ± 0.8, beyond error.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.206 ± 0.007, crit=0.701 ± 0.015 per rating point (14 rating = 1%, 9.813 per %), hit=1.272 ± 0.076 per rating point (10 rating = 1%, 12.718 per %), spell_haste=9.709 ± 0.612, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.877 ± 0.001, spirit=not significant (0.000 ± 0.000), mp5=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 45.7 spell_power points (25.37 DPS) | yes | Blue Dragonscale Helm (252604, -2.38 DPS) [crafted]; Warlord's Mail Helm (231663, -4.88 DPS) [pvp]; Mask of the Unforgiven (13404, -5.81 DPS) [dungeon] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 23.4 spell_power points (12.97 DPS) | yes | Emberfury Talisman (12929, -4.19 DPS) [dungeon]; Lady Maye's Pendant (14558, -4.69 DPS) [world_drop]; Star of Mystaria (12103, -5.19 DPS, sim-verified) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 29.5 spell_power points (16.36 DPS) | yes | Champion's Mail Spaulders (227160, +0.00 DPS) [vendor]; Warlord's Mail Spaulders (231659, +0.00 DPS) [pvp]; Argent Shoulders (19059, -5.50 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -8.40 DPS) [dungeon]; Shroud of Arcane Mastery (22330, -8.54 DPS) [dungeon]; Crystalline Threaded Cape (20697, -10.71 DPS, sim-verified) [world] |
| chest | Stone Guard's Pulsing Breastplate (220844) | PvP rank 6 · Stone Guard · Horde [vendor] | 33.2 spell_power points (18.42 DPS) | yes | Legionnaire's Mail Breastplate (227165, +0.00 DPS) [vendor]; Warlord's Mail Breastplate (231662, +0.00 DPS) [vendor]; Fernpulse Jerkin (15786, -8.31 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 23.6 spell_power points (13.12 DPS) | yes | Runecloth Cuffs (254123, -5.88 DPS) [crafted]; Blackmist Armguards (12966, -6.06 DPS) [dungeon]; Slashclaw Bracers (13211, -6.06 DPS) [dungeon] |
| hands | Trueaim Gauntlets (13255) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Mail Gauntlets (231660, -2.36 DPS) [pvp]; Gloves of the Greatfather (17721, -2.71 DPS, sim-verified) [crafted]; Storm Gauntlets (12632, -4.13 DPS) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sash of the Windreaver (18676, +0.00 DPS) [world]; Defiler's Cloth Girdle (20163, -0.96 DPS) [rep]; Barrage Girdle (18721, -1.18 DPS) [dungeon] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 51.5 spell_power points (28.59 DPS) | yes | General's Mail Leggings (231664, -4.86 DPS) [pvp]; Pristine Scorpid Leggings (252606, -9.26 DPS, sim-verified) [crafted]; Legionnaire's Mail Leggings (227161, -9.55 DPS) [pvp] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 25.8 spell_power points (14.30 DPS) | yes | General's Mail Sabatons (231661, -1.26 DPS) [vendor]; Stormcloth Boots (10039, -2.08 DPS) [crafted]; Omnicast Boots (11822, -5.69 DPS, sim-verified) [dungeon] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -1.06 DPS) [dungeon]; Maiden's Circle (13001, -1.82 DPS) [world_drop]; Naglering (11669, -18.75 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -1.71 DPS) [quest]; Maiden's Circle (13001, -1.71 DPS) [world_drop]; Don Mauricio's Band of Domination (22433, -5.43 DPS, sim-verified) [dungeon] |
| trinket1 | Smolderweb's Eye (13213) | Blackrock Spire: Mother Smolderweb [dungeon] | sim-verified (+12.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -0.71 DPS) [crafted]; Rune of the Guard Captain (19120, -2.12 DPS) [quest]; Cannonball Runner (13382, -3.53 DPS) [dungeon] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -3.64 DPS) [crafted]; Rune of the Guard Captain (19120, -5.05 DPS) [quest]; Draconic Infused Emblem (22268, -8.19 DPS, sim-verified) [dungeon] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -2.97 DPS) [pvp]; Headmaster's Charge (13937, -13.30 DPS, sim-verified) [dungeon]; Wildstaff (20556, -14.50 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Thunder (228176) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burning Totem (272433, +0.00 DPS) [vendor] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Stone Guard's Pulsing Breastplate; wrist: Dryad's Wrist Bindings; hands: Trueaim Gauntlets; waist: Ban'thok Sash; legs: Ironfeather Leggings; feet: Greaves of Withering Despair; finger1: Blessed Band of Light; finger2: Don Julio's Band; trinket1: Smolderweb's Eye; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Totem of Thunder

No-known-source sample (15 of 2331, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

