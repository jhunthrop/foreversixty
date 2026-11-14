# Leveling BiS: Elemental

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 5300300000000000-000000000000000000-0000000000000000)

Set DPS (verified): 31.5. Weights run: 3.7s. Verify run: 6.9s. 286 eligible items had no known source.

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

Set DPS (verified): 62.2. Weights run: 3.6s. Verify run: 6.9s. 475 eligible items had no known source.

Race sweep: dwarf 62.2 ± 0.1 (dps); the guide race dwarf is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.724 ± 0.009, crit=0.077 ± 0.002 per rating point (14 rating = 1%, 1.083 per %), hit=0.270 ± 0.011 per rating point (10 rating = 1%, 2.699 per %), spell_haste=-0.923 ± 0.179, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.857 ± 0.002, spirit=not significant (-0.028 ± 0.073), mp5=1.529 ± 0.124

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 13.2 spell_power points (3.39 DPS) | yes | Totemic Leather Helm (252456, -0.32 DPS) [crafted]; Enduring Cap (3020, -0.42 DPS) [world_drop]; Holy Shroud (2721, -0.57 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 4.3 spell_power points (1.11 DPS) | yes | Seal of the Expedition (271920, -0.19 DPS) [rep]; Crystal Starfire Medallion (5003, -0.37 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.37 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.5 spell_power points (3.98 DPS) | yes | Death Speaker Mantle (6685, -0.40 DPS) [dungeon]; Magician's Mantle (12998, -1.02 DPS) [world_drop]; Invoker's Mantle (215365, -1.25 DPS) [crafted] |
| back | Expeditionary's Cape (271923) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, -0.20 DPS) [crafted]; Vine Pruner's Cloak (279835, -0.20 DPS) [quest]; Resilient Cape (14400, -0.37 DPS) [world_drop] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Green Silk Armor (7065, +0.00 DPS) [crafted]; Mechbuilder's Overalls (9508, -0.26 DPS) [dungeon]; Guardian Armor (4256, -0.37 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 13.6 spell_power points (3.48 DPS) | yes | Mindthrust Bracers (1974, -1.38 DPS) [dungeon]; Glowing Magical Bracelets (13106, -2.00 DPS) [world_drop]; Nightsky Wristbands (6407, -2.37 DPS) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 8.1 spell_power points (2.07 DPS) | yes | Stormrider's Leather Gloves (252498, -0.04 DPS) [crafted]; Truefaith Gloves (7049, -0.23 DPS) [crafted]; Gloves of Insight (9698, -0.27 DPS) [quest] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 14.3 spell_power points (3.68 DPS) | yes | Highlander's Cloth Girdle (20099, -0.30 DPS) [rep]; Moss Cinch (6911, -0.60 DPS) [dungeon]; Belt of Arugal (6392, -0.81 DPS) [dungeon] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 14.3 spell_power points (3.68 DPS) | yes | Dark Ritual Leggings (270031, -0.09 DPS) [quest]; Guardian Pants (5962, -0.14 DPS) [crafted]; Leech Pants (6910, -0.44 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 12.1 spell_power points (3.09 DPS) | yes | Acidic Walkers (9454, -0.33 DPS) [dungeon]; Spidersilk Boots (4320, -0.56 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.63 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lorekeeper's Ring (19525, +0.00 DPS) [rep]; Runemender's Seal (273031, -0.58 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lorekeeper's Ring (19525, -0.14 DPS) [rep]; Runemender's Seal (273031, -0.83 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glimmering Staff (249392, +0.00 DPS) [crafted]; Hydrocane (9452, -1.02 DPS) [dungeon]; Tendonscraper Dagger (273053, -1.02 DPS) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 11.3 spell_power points (2.91 DPS) | yes | Seedcloud Buckler (6630, -0.70 DPS) [dungeon]; Heart of Agamaggan (6694, -1.63 DPS) [dungeon]; Satyr's Rod (15962, -1.79 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Expeditionary's Cape; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Skycaller's Leather Belt; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight

No-known-source sample (15 of 475, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 40 (dwarf, 5500331220001351-000000000000000000-0000000000000000)

Set DPS (verified): 109.5. Weights run: 4.3s. Verify run: 10.4s. 804 eligible items had no known source.

Race sweep: dwarf 109.5 ± 0.2 (dps); the guide race dwarf is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.768 ± 0.020, crit=0.204 ± 0.006 per rating point (14 rating = 1%, 2.861 per %), hit=0.345 ± 0.018 per rating point (10 rating = 1%, 3.448 per %), spell_haste=-1.877 ± 0.239, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.886 ± 0.002, spirit=-0.742 ± 0.170, mp5=not significant (-0.393 ± 0.348)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Miner's Hat of the Deep (9429) | Uldaman [dungeon] | 25.1 spell_power points (6.04 DPS) | yes | Corpseshroud (10574, -0.83 DPS) [dungeon]; Papal Fez (9431, -0.96 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -0.98 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.1 spell_power points (3.17 DPS) | yes | Triune Amulet (7722, -1.24 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.77 DPS, sim-verified) [quest]; Seal of the Expedition (271919, -1.87 DPS) [rep] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 17.0 spell_power points (4.09 DPS) | yes | Green Silken Shoulders (7057, -0.13 DPS) [crafted]; Bloodmage Mantle (7684, -0.26 DPS) [dungeon]; Berylline Pads (4197, -0.56 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Expeditionary's Cape (271922, -0.34 DPS) [rep]; Icy Cloak (4327, -0.68 DPS) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Robe of Power (7054, -0.41 DPS) [crafted]; Elemental Raiment (9434, -0.94 DPS) [world_drop]; Robe of the Magi (1716, -2.03 DPS, sim-verified) [world_drop] |
| wrist | Green Whelp Bracers (7386) | Leatherworking [crafted] | 14.1 spell_power points (3.40 DPS) | yes | Turtle Scale Bracers (8198, -1.03 DPS) [crafted]; Guardian Leather Bracers (4260, -1.15 DPS, sim-verified) [crafted]; Arcane Runed Bracers (4744, -1.23 DPS) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamweave Gloves (10019, -2.84 DPS) [crafted]; Red Mageweave Gloves (10018, -3.42 DPS) [crafted]; Skycaller's Leather Gloves (252529, -3.60 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 18.5 spell_power points (4.46 DPS) | yes | Skycaller's Leather Belt (252522, -0.94 DPS) [crafted]; Gilded Cord (254037, -1.05 DPS) [crafted]; Highlander's Cloth Girdle (20098, -1.30 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.2 spell_power points (5.60 DPS) | yes | Crimson Silk Pantaloons (7062, -1.14 DPS, sim-verified) [crafted]; Stormrider's Leather Pants (252502, -2.07 DPS) [crafted]; Guardian Pants (5962, -2.19 DPS) [crafted] |
| feet | Skycaller's Leather Shoes (252532) (or Skycaller's Mail Boots (252563)) | Leatherworking [crafted] | 18.4 spell_power points (4.43 DPS) | yes | Skycaller's Mail Boots (252563, +0.00 DPS) [crafted]; Mender's Leather Shoes (252533, -1.20 DPS) [crafted]; Mender's Mail Boots (252565, -1.20 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.8 spell_power points (3.34 DPS) | yes | Ring of Forlorn Spirits (2043, -1.41 DPS) [quest]; Runemender's Seal (273031, -1.50 DPS) [world]; Snake Hoop (6750, -3.89 DPS, sim-verified) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | sim-verified (+3.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Ring of Forlorn Spirits (2043, -0.24 DPS) [quest]; Runemender's Seal (273031, -0.34 DPS) [world]; Snake Hoop (6750, -3.03 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hypnotic Blade (7714, -7.32 DPS) [dungeon]; Illusionary Rod (7713, -8.10 DPS, sim-verified) [dungeon]; Glimmering Staff (249392, -9.18 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Miner's Hat of the Deep; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Dreamweave Vest; wrist: Green Whelp Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Skycaller's Leather Shoes; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan

No-known-source sample (15 of 804, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (dwarf, 5500331220001351-030000000000000000-0520000000000000)

Set DPS (verified): 187.2. Weights run: 4.1s. Verify run: 10.7s. 1053 eligible items had no known source.

Race sweep: dwarf 187.2 ± 0.3 (dps); the guide race dwarf is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.081 ± 0.022, crit=0.350 ± 0.009 per rating point (14 rating = 1%, 4.893 per %), hit=0.567 ± 0.034 per rating point (10 rating = 1%, 5.668 per %), spell_haste=-2.244 ± 0.395, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.873 ± 0.002, spirit=not significant (0.421 ± 0.330), mp5=3.085 ± 0.333

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 42.6 spell_power points (12.39 DPS) | yes | Red Mageweave Headband (10033, -0.57 DPS) [crafted]; Dreamweave Circlet (10041, -1.66 DPS) [crafted]; Charged Scorpid Helm (252581, -2.64 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 16.8 spell_power points (4.90 DPS) | yes | Glowing Eye of Mordresh (10769, -0.34 DPS) [dungeon]; Gemshard Heart (17707, -0.61 DPS) [dungeon]; Mindburst Medallion (11196, -0.65 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 40.0 spell_power points (11.64 DPS) | yes | Lead Surveyor's Mantle (11842, -1.93 DPS, sim-verified) [dungeon]; Mantle of Lost Hope (22234, -1.97 DPS) [dungeon]; Rotgrip Mantle (17732, -2.19 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Expeditionary's Cape (271921, -1.68 DPS) [rep]; Runecloth Cloak (13860, -1.71 DPS, sim-verified) [crafted]; Big Voodoo Cloak (8216, -2.41 DPS) [crafted] |
| chest | Feathered Breastplate (8349) | Leatherworking [crafted] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Robes of Insight (940, -0.87 DPS) [world_drop]; Embrace of the Wind Serpent (12462, -1.59 DPS) [world]; Ghostweave Vest (14141, -2.03 DPS, sim-verified) [crafted] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 20.8 spell_power points (6.06 DPS) | yes | Aristocratic Cuffs (12546, -0.17 DPS) [dungeon]; Skycaller's Leather Bracers (252542, -0.33 DPS) [crafted]; Skycaller's Mail Bracers (252571, -0.33 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller's Mail Gauntlets (252585, -1.30 DPS) [crafted]; Ogreseer Fists (11665, -1.37 DPS) [dungeon]; Skycaller's Leather Gauntlets (252550, -1.49 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 29.0 spell_power points (8.43 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Ban'thok Sash (11662, -0.14 DPS) [dungeon]; Dawnspire Cord (12466, -0.71 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | 33.8 spell_power points (9.84 DPS) | yes | Brusslehide Leggings (17751, -0.94 DPS) [dungeon]; Kilt of the Atal'ai Prophet (10807, -1.10 DPS) [dungeon]; Senior Designer's Pantaloons (11841, -1.94 DPS) [dungeon] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 27.5 spell_power points (8.00 DPS) | yes | Skycaller's Leather Boots (252471, -0.17 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -0.17 DPS) [crafted]; Coldstone Slippers (18697, -2.87 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 24.3 spell_power points (7.08 DPS) | yes | Cyclopean Band (11824, -1.77 DPS) [dungeon]; Woodseed Hoop (17768, -2.19 DPS) [quest]; Mindseye Circle (10634, -2.48 DPS) [dungeon] |
| finger2 | Brainlash (6440) | Zul'Farrak [dungeon] | 23.9 spell_power points (6.96 DPS) | yes | Woodseed Hoop (17768, -2.07 DPS) [quest]; Cyclopean Band (11824, -2.26 DPS, sim-verified) [dungeon]; Mindseye Circle (10634, -2.36 DPS) [dungeon] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Enriched Seal (271903) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.84 DPS, sim-verified) [world_drop]; Frozen Heart of the Mountain (249469, -3.00 DPS) [crafted]; Ankh of Life (1713, -3.02 DPS) [world_drop] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Staff of Jordan (873, -16.89 DPS, sim-verified) [world_drop] |
| off_hand | Gizlock's Hypertech Buckler (17718) | Maraudon: Tinkerer Gizlock [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Truesilver Conduit (249455, -2.96 DPS, sim-verified) [crafted]; Enthralled Sphere (11625, -3.25 DPS) [dungeon]; Orb of Lorica (11262, -3.30 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Feathered Breastplate; wrist: Runic Leather Bracers; waist: Skycaller's Leather Waistguard; legs: Spellshock Leggings; feet: Greaves of Withering Despair; finger1: Lorekeeper's Ring; finger2: Brainlash; trinket1: Thunderbrew's Boot Flask; trinket2: Enriched Seal; main_hand: Blade of Eternal Darkness; off_hand: Gizlock's Hypertech Buckler

No-known-source sample (15 of 1053, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (dwarf, 5500331220001351-055000010000000000-0520020000000000)

Set DPS (verified): 305.1. Weights run: 4.1s. Verify run: 42.8s. 2377 eligible items had no known source.

Race sweep: dwarf 305.1 ± 0.5 (dps); the guide race dwarf is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.129 ± 0.028, crit=0.448 ± 0.010 per rating point (14 rating = 1%, 6.276 per %), hit=0.758 ± 0.043 per rating point (10 rating = 1%, 7.577 per %), spell_haste=-2.389 ± 0.464, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.840 ± 0.002, spirit=not significant (0.664 ± 0.357), mp5=2.607 ± 0.455

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | The Postmaster's Band (13390) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Living Crown (252561, +0.00 DPS) [crafted]; Blue Dragonscale Helm (252604, +0.00 DPS) [crafted]; Spellweaver's Turban (22267, -0.21 DPS) [dungeon] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lady Maye's Pendant (14558, -0.16 DPS) [world_drop]; Nacreous Shell Necklace (22403, -2.35 DPS) [dungeon]; Tooth of Gnarr (13141, -2.46 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Theramore Shoulderguards (272957, +0.00 DPS) [rep]; Theramore Shoulderpads (271953, -0.13 DPS) [rep]; Theramore Shoulders (271954, -0.13 DPS) [rep] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, -1.49 DPS) [quest]; Gracious Cape (18743, -1.94 DPS) [dungeon]; Spritecaster Cape (11623, -2.51 DPS) [dungeon] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Magister's Robes (16688, -1.34 DPS) [dungeon]; The Postmaster's Tunic (13388, -1.89 DPS) [dungeon]; Feathered Breastplate (8349, -2.21 DPS) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 35.7 spell_power points (11.39 DPS) | yes | Brazecore Armguards (13179, -2.62 DPS) [dungeon]; Loomguard Armbraces (13969, -3.18 DPS) [dungeon]; Bleak Howler Armguards (13208, -3.58 DPS) [dungeon] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Theramore Handwraps (271945, +0.00 DPS) [rep]; Theramore Handguards (271950, +0.00 DPS) [rep]; Dracorian Gauntlets (13344, -0.34 DPS) [dungeon] |
| waist | Stormseeker's Girdle (272399) | Pix Xizzix [vendor] | 43.1 spell_power points (13.76 DPS) | yes | Crystallized Girdle (12606, -1.49 DPS) [dungeon]; Sash of the Windreaver (18676, -1.71 DPS) [world]; Barrage Girdle (18721, -2.26 DPS) [dungeon] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 64.6 spell_power points (20.62 DPS) | yes | Leggings of Arcana (12756, -1.31 DPS) [quest]; Ghostloom Leggings (14545, -2.69 DPS) [dungeon]; Red Dragonscale Leggings (252603, -2.70 DPS) [crafted] |
| feet | The Postmaster's Treads (13391) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Omnicast Boots (11822, +0.00 DPS) [dungeon]; Wolfrunner Shoes (13101, +0.00 DPS) [world_drop]; Dragonrider Boots (18102, +0.00 DPS) [dungeon] |
| finger1 | The Postmaster's Seal (13392) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, +0.00 DPS) [quest]; Maiden's Circle (13001, +0.00 DPS) [world_drop]; Seal of Rivendare (13345, +0.00 DPS) [dungeon] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Maiden's Circle (13001, -2.49 DPS) [world_drop]; Seal of Rivendare (13345, -3.35 DPS) [dungeon]; Songstone of Ironforge (12543, -3.76 DPS) [quest] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -0.91 DPS) [vendor]; Shard of the Splithooves (10659, -1.66 DPS) [quest]; Enriched Seal (271903, -2.50 DPS) [rep] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Enriched Seal (271903, +0.00 DPS) [rep]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Trindlehaven Staff (13161, +0.00 DPS) [dungeon]; Ironbark Staff (20069, +0.00 DPS) [rep]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp] |
| off_hand | Rhombeard Protector (13205) | Blackrock Spire: Gizrul the Slavener [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Barricade (235473, +0.00 DPS) [vendor]; Grand Marshal's Tome of Restoration (234590, -0.53 DPS) [pvp] |
| ranged | Totem of Thunder (228176) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 60:** head: The Postmaster's Band; neck: Jeweled Amulet of Cainwyn; back: Arcanoweave Cloak; chest: Ironfeather Breastplate; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Stormseeker's Girdle; legs: Ironfeather Leggings; feet: The Postmaster's Treads; finger1: The Postmaster's Seal; finger2: Channeler's Ring; trinket1: Burst of Knowledge; trinket2: Second Wind; off_hand: Rhombeard Protector; ranged: Totem of Thunder

No-known-source sample (15 of 2377, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60, raid preset (dwarf, 5500331220001351-055000000000000000-0530020000000000)

Set DPS (verified): 759.5. Weights run: 3.1s. Verify run: 34.6s. 2377 eligible items had no known source.

Race sweep: dwarf 759.5 ± 0.6 (dps); the guide race dwarf is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.207 ± 0.007, crit=0.706 ± 0.015 per rating point (14 rating = 1%, 9.883 per %), hit=1.281 ± 0.076 per rating point (10 rating = 1%, 12.812 per %), spell_haste=9.779 ± 0.617, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.877 ± 0.001, spirit=not significant (0.000 ± 0.000), mp5=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 45.8 spell_power points (25.42 DPS) | yes | Blue Dragonscale Helm (252604, -2.37 DPS) [crafted]; Mask of the Unforgiven (13404, -5.72 DPS) [dungeon]; Crimson Felt Hat (18727, -7.62 DPS) [dungeon] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 23.5 spell_power points (13.02 DPS) | yes | Emberfury Talisman (12929, -4.21 DPS) [dungeon]; Lady Maye's Pendant (14558, -4.73 DPS) [world_drop]; Star of Mystaria (12103, -5.59 DPS, sim-verified) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 29.6 spell_power points (16.41 DPS) | yes | Burial Shawl (18681, -3.47 DPS) [dungeon]; Theramore Shoulderguards (272957, -4.43 DPS) [rep]; Argent Shoulders (19059, -5.34 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -8.45 DPS) [dungeon]; Shroud of the Exile (15421, -9.10 DPS) [quest]; Crystalline Threaded Cape (20697, -11.75 DPS, sim-verified) [world] |
| chest | Fernpulse Jerkin (15786) | High Chief Winterfall [quest] | 27.7 spell_power points (15.35 DPS) | yes | Tunic of the Crescent Moon (22409, -0.16 DPS) [dungeon]; Alanna's Embrace (13314, -0.28 DPS) [dungeon]; Vest of Elements (16666, -0.28 DPS) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 23.7 spell_power points (13.13 DPS) | yes | Runecloth Cuffs (254123, -5.87 DPS) [crafted]; Blackmist Armguards (12966, -6.02 DPS) [dungeon]; Slashclaw Bracers (13211, -6.02 DPS) [dungeon] |
| hands | Trueaim Gauntlets (13255) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of the Greatfather (17721, -3.14 DPS) [crafted]; Storm Gauntlets (12632, -4.06 DPS, sim-verified) [crafted]; Hands of Power (13253, -7.53 DPS) [dungeon] |
| waist | Sash of the Windreaver (18676) | The Windreaver [world] | 27.8 spell_power points (15.40 DPS) | yes | Highlander's Cloth Girdle (20047, -1.46 DPS) [rep]; Barrage Girdle (18721, -1.72 DPS) [dungeon]; Ban'thok Sash (11662, -7.87 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) (or Sentinel's Lizardhide Pants (237817)) | Illiyana Moonblaze [vendor] | 53.8 spell_power points (29.88 DPS) | yes | Sentinel's Lizardhide Pants (237817, +0.00 DPS) [vendor]; Ironfeather Leggings (252486, -1.23 DPS) [crafted]; Pristine Scorpid Leggings (252606, -3.31 DPS) [crafted] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 25.9 spell_power points (14.36 DPS) | yes | Stormcloth Boots (10039, -2.14 DPS) [crafted]; Venomspew Footpads (275606, -2.16 DPS) [crafted]; Omnicast Boots (11822, -6.86 DPS, sim-verified) [dungeon] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -2.13 DPS) [dungeon]; Maiden's Circle (13001, -2.92 DPS) [world_drop]; Naglering (11669, -22.70 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -1.80 DPS) [quest]; Maiden's Circle (13001, -1.80 DPS) [world_drop]; Don Mauricio's Band of Domination (22433, -5.93 DPS, sim-verified) [dungeon] |
| trinket1 | Smolderweb's Eye (13213) | Blackrock Spire: Mother Smolderweb [dungeon] | sim-verified (+12.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -0.71 DPS) [crafted]; Cannonball Runner (13382, -3.55 DPS) [dungeon]; Uther's Strength (11302, -3.78 DPS) [world_drop] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -3.59 DPS) [crafted]; Serenity Field (272439, -6.10 DPS) [vendor]; Draconic Infused Emblem (22268, -9.20 DPS, sim-verified) [dungeon] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -2.96 DPS) [pvp]; Headmaster's Charge (13937, -14.16 DPS, sim-verified) [dungeon]; Staff of Jordan (873, -15.27 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Thunder (228176) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burning Totem (272433, -3.32 DPS, sim-verified) [vendor] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Fernpulse Jerkin; wrist: Dryad's Wrist Bindings; hands: Trueaim Gauntlets; waist: Sash of the Windreaver; legs: Sentinel's Silk Leggings; finger1: Channeler's Ring; finger2: Don Julio's Band; trinket1: Smolderweb's Eye; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Totem of Thunder

No-known-source sample (15 of 2377, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

## Horde

### Band 20 (orc, 5300300000000000-000000000000000000-0000000000000000)

Set DPS (verified): 30.9. Weights run: 3.7s. Verify run: 7.6s. 275 eligible items had no known source.

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

Set DPS (verified): 62.1. Weights run: 3.6s. Verify run: 7.8s. 469 eligible items had no known source.

Race sweep: orc 62.1 ± 0.1, troll 61.3 ± 0.1, tauren 61.2 ± 0.1, windshaper-skyborne 59.7 ± 0.1 (dps); the guide race orc is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.724 ± 0.009, crit=0.077 ± 0.002 per rating point (14 rating = 1%, 1.083 per %), hit=0.270 ± 0.011 per rating point (10 rating = 1%, 2.699 per %), spell_haste=-0.923 ± 0.179, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.857 ± 0.002, spirit=not significant (-0.028 ± 0.073), mp5=1.529 ± 0.124

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 13.2 spell_power points (3.39 DPS) | yes | Totemic Leather Helm (252456, -0.32 DPS) [crafted]; Enduring Cap (3020, -0.42 DPS) [world_drop]; Holy Shroud (2721, -0.57 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 4.3 spell_power points (1.11 DPS) | yes | Darkspear Warding Pendant (272075, -0.19 DPS) [vendor]; Crystal Starfire Medallion (5003, -0.37 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.37 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.5 spell_power points (3.98 DPS) | yes | Death Speaker Mantle (6685, -0.40 DPS) [dungeon]; Magician's Mantle (12998, -1.02 DPS) [world_drop]; Invoker's Mantle (215365, -1.25 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windsong Drape (15468, -0.20 DPS) [quest]; Resilient Cape (14400, -0.37 DPS) [world_drop]; Hillman's Cloak (3719, -0.74 DPS, sim-verified) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 18.4 spell_power points (4.72 DPS) | yes | Beguiler Robes (7728, -0.19 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.44 DPS) [dungeon]; Guardian Armor (4256, -0.56 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 13.6 spell_power points (3.48 DPS) | yes | Mindthrust Bracers (1974, -1.47 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -2.00 DPS) [world_drop]; Nightsky Wristbands (6407, -2.37 DPS) [world_drop] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | 10.4 spell_power points (2.68 DPS) | yes | Jutebraid Gloves (10654, -0.21 DPS) [quest]; Stormrider's Leather Gloves (252498, -0.65 DPS) [crafted]; Truefaith Gloves (7049, -0.84 DPS) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 14.3 spell_power points (3.68 DPS) | yes | Defiler's Cloth Girdle (20164, -0.46 DPS, sim-verified) [rep]; Moss Cinch (6911, -0.60 DPS) [dungeon]; Belt of Arugal (6392, -0.81 DPS) [dungeon] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 14.3 spell_power points (3.68 DPS) | yes | Guardian Pants (5962, -0.14 DPS) [crafted]; Leech Pants (6910, -0.44 DPS) [dungeon]; Dark Ritual Leggings (270031, -0.71 DPS, sim-verified) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 12.1 spell_power points (3.09 DPS) | yes | Acidic Walkers (9454, -0.33 DPS) [dungeon]; Spidersilk Boots (4320, -0.56 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.63 DPS) [crafted] |
| finger1 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 10.6 spell_power points (2.71 DPS) | yes | Snake Hoop (6750, -0.25 DPS) [quest]; Runemender's Seal (273031, -0.83 DPS) [world]; Lavishly Jeweled Ring (1156, -1.60 DPS) [dungeon] |
| finger2 | Advisor's Ring (19521) | Warsong Outriders [rep] | 10.1 spell_power points (2.58 DPS) | yes | Snake Hoop (6750, -0.11 DPS) [quest]; Runemender's Seal (273031, -0.70 DPS) [world]; Lavishly Jeweled Ring (1156, -1.46 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.02 DPS) [dungeon]; Tendonscraper Dagger (273053, -1.02 DPS) [world]; Glimmering Staff (249392, -2.49 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 11.3 spell_power points (2.91 DPS) | yes | Seedcloud Buckler (6630, -0.66 DPS, sim-verified) [dungeon]; Witch's Finger (16887, -1.42 DPS) [quest]; Heart of Agamaggan (6694, -1.63 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Darkspear Raider's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; feet: Gilded Slippers; finger1: Ring of Power Regulation; finger2: Advisor's Ring; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight

No-known-source sample (15 of 469, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5255 Quilboar Tomahawk

### Band 40 (orc, 5500331220001351-000000000000000000-0000000000000000)

Set DPS (verified): 109.1. Weights run: 4.3s. Verify run: 11.2s. 787 eligible items had no known source.

Race sweep: orc 109.1 ± 0.2, troll 108.0 ± 0.2, tauren 107.5 ± 0.2, windshaper-skyborne 103.3 ± 0.2 (dps); the guide race orc is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.768 ± 0.020, crit=0.204 ± 0.006 per rating point (14 rating = 1%, 2.861 per %), hit=0.345 ± 0.018 per rating point (10 rating = 1%, 3.448 per %), spell_haste=-1.877 ± 0.239, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.886 ± 0.002, spirit=-0.742 ± 0.170, mp5=not significant (-0.393 ± 0.348)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Miner's Hat of the Deep (9429) | Uldaman [dungeon] | 25.1 spell_power points (6.04 DPS) | yes | Corpseshroud (10574, -0.83 DPS) [dungeon]; Papal Fez (9431, -0.96 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -0.98 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.1 spell_power points (3.17 DPS) | yes | Triune Amulet (7722, -1.24 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.41 DPS, sim-verified) [quest]; Darkspear Warding Pendant (272074, -1.87 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 17.0 spell_power points (4.09 DPS) | yes | Green Silken Shoulders (7057, -0.13 DPS) [crafted]; Bloodmage Mantle (7684, -0.26 DPS) [dungeon]; Berylline Pads (4197, -0.56 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.34 DPS) [vendor]; Icy Cloak (4327, -0.68 DPS) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Robe of Power (7054, -0.41 DPS) [crafted]; Elemental Raiment (9434, -0.94 DPS) [world_drop]; Robe of the Magi (1716, -2.67 DPS, sim-verified) [world_drop] |
| wrist | Green Whelp Bracers (7386) | Leatherworking [crafted] | 14.1 spell_power points (3.40 DPS) | yes | Radiant Silver Bracers (4545, -0.96 DPS) [quest]; Turtle Scale Bracers (8198, -1.03 DPS) [crafted]; Guardian Leather Bracers (4260, -1.18 DPS, sim-verified) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamweave Gloves (10019, -2.84 DPS) [crafted]; Red Mageweave Gloves (10018, -3.42 DPS) [crafted]; Skycaller's Leather Gloves (252529, -3.60 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 18.5 spell_power points (4.46 DPS) | yes | Defiler's Cloth Girdle (20166, -0.93 DPS, sim-verified) [rep]; Skycaller's Leather Belt (252522, -0.94 DPS) [crafted]; Gilded Cord (254037, -1.05 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.2 spell_power points (5.60 DPS) | yes | Crimson Silk Pantaloons (7062, -0.94 DPS, sim-verified) [crafted]; Stormrider's Leather Pants (252502, -2.07 DPS) [crafted]; Guardian Pants (5962, -2.19 DPS) [crafted] |
| feet | Skycaller's Leather Shoes (252532) (or Skycaller's Mail Boots (252563)) | Leatherworking [crafted] | 18.4 spell_power points (4.43 DPS) | yes | Skycaller's Mail Boots (252563, +0.00 DPS) [crafted]; Mender's Leather Shoes (252533, -1.20 DPS) [crafted]; Mender's Mail Boots (252565, -1.20 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.8 spell_power points (3.34 DPS) | yes | Runemender's Seal (273031, -1.50 DPS) [world]; Reedknot Ring (9622, -1.65 DPS) [quest]; Snake Hoop (6750, -3.81 DPS, sim-verified) [quest] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Runemender's Seal (273031, -0.34 DPS) [world]; Reedknot Ring (9622, -0.48 DPS) [quest]; Snake Hoop (6750, -3.45 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hypnotic Blade (7714, -7.32 DPS) [dungeon]; Illusionary Rod (7713, -8.38 DPS, sim-verified) [dungeon]; Glimmering Staff (249392, -9.18 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Miner's Hat of the Deep; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Dreamweave Vest; wrist: Green Whelp Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Skycaller's Leather Shoes; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan

No-known-source sample (15 of 787, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 50 (orc, 5500331220001351-030000000000000000-0520000000000000)

Set DPS (verified): 186.7. Weights run: 4.1s. Verify run: 11.6s. 1021 eligible items had no known source.

Race sweep: orc 186.7 ± 0.3, troll 184.5 ± 0.3, tauren 184.5 ± 0.3, windshaper-skyborne 183.5 ± 0.3 (dps); the guide race orc is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.081 ± 0.022, crit=0.350 ± 0.009 per rating point (14 rating = 1%, 4.893 per %), hit=0.567 ± 0.034 per rating point (10 rating = 1%, 5.668 per %), spell_haste=-2.244 ± 0.395, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.873 ± 0.002, spirit=not significant (0.421 ± 0.330), mp5=3.085 ± 0.333

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 42.6 spell_power points (12.39 DPS) | yes | Red Mageweave Headband (10033, -0.57 DPS) [crafted]; Dreamweave Circlet (10041, -1.66 DPS) [crafted]; Charged Scorpid Helm (252581, -2.64 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 16.8 spell_power points (4.90 DPS) | yes | Glowing Eye of Mordresh (10769, -0.34 DPS) [dungeon]; Gemshard Heart (17707, -0.61 DPS) [dungeon]; Mindburst Medallion (11196, -0.65 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 40.0 spell_power points (11.64 DPS) | yes | Lead Surveyor's Mantle (11842, -1.52 DPS, sim-verified) [dungeon]; Mantle of Lost Hope (22234, -1.97 DPS) [dungeon]; Rotgrip Mantle (17732, -2.19 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -0.37 DPS) [quest]; Runecloth Cloak (13860, -1.56 DPS) [crafted]; Darkspear Raider's Cloak (272076, -1.68 DPS) [vendor] |
| chest | Feathered Breastplate (8349) | Leatherworking [crafted] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Robes of Insight (940, -0.87 DPS) [world_drop]; Stone Guard's Pulsing Breastplate (220844, -0.96 DPS) [vendor]; Ghostweave Vest (14141, -2.06 DPS, sim-verified) [crafted] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 20.8 spell_power points (6.06 DPS) | yes | Aristocratic Cuffs (12546, -0.17 DPS) [dungeon]; Skycaller's Leather Bracers (252542, -0.33 DPS) [crafted]; Skycaller's Mail Bracers (252571, -0.33 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller's Mail Gauntlets (252585, -1.30 DPS) [crafted]; Ogreseer Fists (11665, -1.37 DPS) [dungeon]; Skycaller's Leather Gauntlets (252550, -1.54 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 29.0 spell_power points (8.43 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Ban'thok Sash (11662, -0.14 DPS) [dungeon]; Dawnspire Cord (12466, -0.71 DPS) [dungeon] |
| legs | Stone Guard's Pulsing Legplates (220847) | PvP rank 6 · Stone Guard · Horde [vendor] | 36.6 spell_power points (10.66 DPS) | yes | Spellshock Leggings (9484, -0.82 DPS) [dungeon]; Brusslehide Leggings (17751, -1.77 DPS) [dungeon]; Kilt of the Atal'ai Prophet (10807, -1.92 DPS) [dungeon] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 27.5 spell_power points (8.00 DPS) | yes | Skycaller's Leather Boots (252471, -0.17 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -0.17 DPS) [crafted]; Coldstone Slippers (18697, -2.16 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19519) | Warsong Outriders [rep] | 24.3 spell_power points (7.08 DPS) | yes | Cyclopean Band (11824, -1.77 DPS) [dungeon]; Woodseed Hoop (17768, -2.19 DPS) [quest]; Mindseye Circle (10634, -2.48 DPS) [dungeon] |
| finger2 | Brainlash (6440) | Zul'Farrak [dungeon] | 23.9 spell_power points (6.96 DPS) | yes | Cyclopean Band (11824, -1.97 DPS, sim-verified) [dungeon]; Woodseed Hoop (17768, -2.07 DPS) [quest]; Mindseye Circle (10634, -2.36 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Enamored Water Spirit (20503, +0.00 DPS) [quest]; Ankh of Life (1713, -0.01 DPS) [world_drop] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Enamored Water Spirit (20503, -0.90 DPS) [quest]; Uther's Strength (11302, -2.74 DPS) [world_drop]; Ankh of Life (1713, -3.02 DPS) [world_drop] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Wildstaff (20556, +0.00 DPS) [quest]; Staff of Jordan (873, -18.24 DPS, sim-verified) [world_drop] |
| off_hand | Gizlock's Hypertech Buckler (17718) | Maraudon: Tinkerer Gizlock [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Truesilver Conduit (249455, -1.71 DPS, sim-verified) [crafted]; Omega Orb (7749, -2.88 DPS) [quest]; Flask of Forest Mojo (19115, -3.06 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Feathered Breastplate; wrist: Runic Leather Bracers; waist: Skycaller's Leather Waistguard; legs: Stone Guard's Pulsing Legplates; feet: Greaves of Withering Despair; finger1: Advisor's Ring; finger2: Brainlash; trinket1: Frozen Heart of the Mountain; trinket2: Darkspear Voodoo Seal; main_hand: Blade of Eternal Darkness; off_hand: Gizlock's Hypertech Buckler

No-known-source sample (15 of 1021, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 60 (orc, 5500331220001351-055000010000000000-0520020000000000)

Set DPS (verified): 303.1. Weights run: 4.1s. Verify run: 43.9s. 2334 eligible items had no known source.

Race sweep: windshaper-skyborne 304.2 ± 0.5, orc 303.1 ± 0.5, troll 300.7 ± 0.5, tauren 299.9 ± 0.5 (dps); windshaper-skyborne beats the guide race orc by 1.1 (0.36%) ± 0.7, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.129 ± 0.028, crit=0.448 ± 0.010 per rating point (14 rating = 1%, 6.276 per %), hit=0.758 ± 0.043 per rating point (10 rating = 1%, 7.577 per %), spell_haste=-2.389 ± 0.464, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.840 ± 0.002, spirit=not significant (0.664 ± 0.357), mp5=2.607 ± 0.455

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | The Postmaster's Band (13390) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Mail Helm (231663, +0.00 DPS) [pvp]; Living Crown (252561, +0.00 DPS) [crafted]; Blue Dragonscale Helm (252604, +0.00 DPS) [crafted] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lady Maye's Pendant (14558, -0.16 DPS) [world_drop]; Nacreous Shell Necklace (22403, -2.35 DPS) [dungeon]; Tooth of Gnarr (13141, -2.46 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Mail Spaulders (231659, +0.00 DPS) [pvp]; Darkspear Shoulderguards (272958, +0.00 DPS) [vendor]; Darkspear Shoulderpads (272103, -0.13 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, -1.49 DPS) [quest]; Gracious Cape (18743, -1.94 DPS) [dungeon]; Spritecaster Cape (11623, -2.51 DPS) [dungeon] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Mail Breastplate (231662, -0.41 DPS) [vendor]; Magister's Robes (16688, -1.34 DPS) [dungeon]; The Postmaster's Tunic (13388, -1.89 DPS) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 35.7 spell_power points (11.39 DPS) | yes | Brazecore Armguards (13179, -2.62 DPS) [dungeon]; Loomguard Armbraces (13969, -3.18 DPS) [dungeon]; Bleak Howler Armguards (13208, -3.58 DPS) [dungeon] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Mail Gauntlets (231660, +0.00 DPS) [pvp]; Raider Handwraps (272097, +0.00 DPS) [vendor]; Raider Handguards (272101, +0.00 DPS) [vendor] |
| waist | Stormseeker's Girdle (272399) | Pix Xizzix [vendor] | 43.1 spell_power points (13.76 DPS) | yes | Crystallized Girdle (12606, -1.49 DPS) [dungeon]; Sash of the Windreaver (18676, -1.71 DPS) [world]; Barrage Girdle (18721, -2.26 DPS) [dungeon] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 64.6 spell_power points (20.62 DPS) | yes | General's Mail Leggings (231664, -1.03 DPS) [pvp]; Leggings of Arcana (12756, -1.31 DPS) [quest]; Ghostloom Leggings (14545, -2.69 DPS) [dungeon] |
| feet | The Postmaster's Treads (13391) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Omnicast Boots (11822, +0.00 DPS) [dungeon]; Dragonrider Boots (18102, +0.00 DPS) [dungeon]; General's Mail Sabatons (231661, +0.00 DPS) [vendor] |
| finger1 | The Postmaster's Seal (13392) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, +0.00 DPS) [quest]; Maiden's Circle (13001, +0.00 DPS) [world_drop]; Seal of Rivendare (13345, +0.00 DPS) [dungeon] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Maiden's Circle (13001, -2.49 DPS) [world_drop]; Seal of Rivendare (13345, -3.35 DPS) [dungeon]; Eye of Orgrimmar (12545, -3.76 DPS) [quest] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -0.91 DPS) [vendor]; Shard of the Splithooves (10659, -1.66 DPS) [quest]; Darkspear Voodoo Seal (272061, -2.50 DPS) [vendor] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Trindlehaven Staff (13161, +0.00 DPS) [dungeon]; Ironbark Staff (20220, +0.00 DPS) [rep]; High Warlord's War Staff (234549, +0.00 DPS) [pvp] |
| off_hand | Rhombeard Protector (13205) | Blackrock Spire: Gizrul the Slavener [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Barricade (235474, +0.00 DPS) [vendor]; High Warlord's Tome of Mending (234564, -0.53 DPS) [pvp] |
| ranged | Burning Totem (272433) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 60:** head: The Postmaster's Band; neck: Jeweled Amulet of Cainwyn; back: Arcanoweave Cloak; chest: Ironfeather Breastplate; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Stormseeker's Girdle; legs: Ironfeather Leggings; feet: The Postmaster's Treads; finger1: The Postmaster's Seal; finger2: Channeler's Ring; trinket1: Burst of Knowledge; trinket2: Second Wind; off_hand: Rhombeard Protector; ranged: Burning Totem

No-known-source sample (15 of 2334, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 60, raid preset (orc, 5500331220001351-055000000000000000-0530020000000000)

Set DPS (verified): 761.9. Weights run: 3.1s. Verify run: 34.5s. 2334 eligible items had no known source.

Race sweep: windshaper-skyborne 765.8 ± 0.6, orc 761.9 ± 0.6, troll 758.8 ± 0.6, tauren 754.6 ± 0.6 (dps); windshaper-skyborne beats the guide race orc by 3.9 (0.52%) ± 0.8, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.207 ± 0.007, crit=0.706 ± 0.015 per rating point (14 rating = 1%, 9.883 per %), hit=1.281 ± 0.076 per rating point (10 rating = 1%, 12.812 per %), spell_haste=9.779 ± 0.617, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.877 ± 0.001, spirit=not significant (0.000 ± 0.000), mp5=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 45.8 spell_power points (25.42 DPS) | yes | Blue Dragonscale Helm (252604, -2.37 DPS) [crafted]; Warlord's Mail Helm (231663, -4.87 DPS) [pvp]; Mask of the Unforgiven (13404, -5.72 DPS) [dungeon] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 23.5 spell_power points (13.02 DPS) | yes | Emberfury Talisman (12929, -4.21 DPS) [dungeon]; Lady Maye's Pendant (14558, -4.73 DPS) [world_drop]; Star of Mystaria (12103, -5.07 DPS, sim-verified) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 29.6 spell_power points (16.41 DPS) | yes | Champion's Mail Spaulders (227160, +0.00 DPS) [vendor]; Warlord's Mail Spaulders (231659, +0.00 DPS) [pvp]; Argent Shoulders (19059, -4.34 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -8.45 DPS) [dungeon]; Shroud of the Exile (15421, -9.10 DPS) [quest]; Crystalline Threaded Cape (20697, -9.64 DPS, sim-verified) [world] |
| chest | Stone Guard's Pulsing Breastplate (220844) | PvP rank 6 · Stone Guard · Horde [vendor] | 33.3 spell_power points (18.48 DPS) | yes | Legionnaire's Mail Breastplate (227165, +0.00 DPS) [vendor]; Warlord's Mail Breastplate (231662, +0.00 DPS) [vendor]; Fernpulse Jerkin (15786, -7.23 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 23.7 spell_power points (13.13 DPS) | yes | Runecloth Cuffs (254123, -5.87 DPS) [crafted]; Blackmist Armguards (12966, -6.02 DPS) [dungeon]; Slashclaw Bracers (13211, -6.02 DPS) [dungeon] |
| hands | Trueaim Gauntlets (13255) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Mail Gauntlets (231660, -2.46 DPS) [pvp]; Gloves of the Greatfather (17721, -3.14 DPS) [crafted]; Storm Gauntlets (12632, -4.23 DPS) [crafted] |
| waist | Sash of the Windreaver (18676) | The Windreaver [world] | 27.8 spell_power points (15.40 DPS) | yes | Defiler's Cloth Girdle (20163, -1.46 DPS) [rep]; Barrage Girdle (18721, -1.72 DPS) [dungeon]; Ban'thok Sash (11662, -8.04 DPS, sim-verified) [dungeon] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 51.6 spell_power points (28.64 DPS) | yes | General's Mail Leggings (231664, -4.86 DPS) [pvp]; Pristine Scorpid Leggings (252606, -5.18 DPS, sim-verified) [crafted]; Legionnaire's Mail Leggings (227161, -9.55 DPS) [pvp] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 25.9 spell_power points (14.36 DPS) | yes | General's Mail Sabatons (231661, -1.31 DPS) [vendor]; Stormcloth Boots (10039, -2.14 DPS) [crafted]; Omnicast Boots (11822, -5.83 DPS, sim-verified) [dungeon] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -2.13 DPS) [dungeon]; Maiden's Circle (13001, -2.92 DPS) [world_drop]; Naglering (11669, -22.25 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -1.80 DPS) [quest]; Maiden's Circle (13001, -1.80 DPS) [world_drop]; Don Mauricio's Band of Domination (22433, -4.88 DPS, sim-verified) [dungeon] |
| trinket1 | Smolderweb's Eye (13213) | Blackrock Spire: Mother Smolderweb [dungeon] | sim-verified (+11.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -0.71 DPS) [crafted]; Rune of the Guard Captain (19120, -2.13 DPS) [quest]; Cannonball Runner (13382, -3.55 DPS) [dungeon] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -3.59 DPS) [crafted]; Rune of the Guard Captain (19120, -5.01 DPS) [quest]; Draconic Infused Emblem (22268, -7.92 DPS, sim-verified) [dungeon] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -2.96 DPS) [pvp]; Headmaster's Charge (13937, -13.26 DPS, sim-verified) [dungeon]; Wildstaff (20556, -14.49 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Thunder (228176) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burning Totem (272433, -3.18 DPS, sim-verified) [vendor] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Stone Guard's Pulsing Breastplate; wrist: Dryad's Wrist Bindings; hands: Trueaim Gauntlets; waist: Sash of the Windreaver; legs: Ironfeather Leggings; finger1: Channeler's Ring; finger2: Don Julio's Band; trinket1: Smolderweb's Eye; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Totem of Thunder

No-known-source sample (15 of 2334, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

