# Leveling BiS: Elemental

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 5500100000000000-000000000000000000-0000000000000000)

Set DPS (verified): 30.5. Weights run: 7.2s. Verify run: 6.1s. 260 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.986 ± 0.016, crit=0.094 ± 0.003 per rating point (14 rating = 1%, 1.309 per %), hit=0.325 ± 0.013 per rating point (10 rating = 1%, 3.246 per %), spell_haste=not significant (0.266 ± 0.138), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.867 ± 0.002, spirit=not significant (0.388 ± 0.179), mp5=1.929 ± 0.286

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Hood (252448) | Leatherworking [crafted] | 9.0 spell_power points (1.16 DPS) | yes | Shadow Goggles (4373, -0.22 DPS) [crafted]; Trapper's Leather Hood (252505, -0.32 DPS, sim-verified) [crafted]; Wisdom's Leather Hood (252507, -0.39 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.8 spell_power points (0.10 DPS) | yes | Tarnished Locket (279870, -0.62 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.9 spell_power points (1.78 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.59 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -1.04 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -1.27 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sanguine Cape (14376, -0.21 DPS) [world_drop]; Pearl-clasped Cloak (5542, -0.21 DPS, sim-verified) [crafted]; Caretaker's Cape (20428, -0.23 DPS) [rep] |
| chest | Stormrider's Leather Armor (252492) | Leatherworking [crafted] | 14.1 spell_power points (1.81 DPS) | yes | Wisdom's Leather Armor (252493, -0.38 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -0.39 DPS) [crafted]; Totemic Leather Armor (252435, -0.66 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 10.7 spell_power points (1.38 DPS) | yes | Bright Bracers (3647, -0.87 DPS) [world_drop]; Owl Bracers (4796, -0.96 DPS, sim-verified) [vendor]; Inscribed Leather Bracers (3205, -1.00 DPS) [dungeon] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 8.9 spell_power points (1.15 DPS) | yes | Serpent Gloves (5970, -0.25 DPS) [dungeon]; Windfelt Gloves (5630, -0.25 DPS) [quest]; Pristine Gloves (253913, -0.25 DPS) [crafted] |
| waist | Pristine Sash (253925) (or Stormrider's Leather Belt (252432)) | Tailoring [crafted] | 7.9 spell_power points (1.02 DPS) | yes | Stormrider's Leather Belt (252432, +0.00 DPS) [crafted]; Keller's Girdle (2911, -0.01 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.13 DPS) [crafted] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 17.5 spell_power points (2.24 DPS) | yes | Wisdom's Leather Pants (252503, -0.38 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.51 DPS) [crafted]; Totemic Leather Pants (252446, -0.84 DPS) [crafted] |
| feet | Stormrider's Leather Boots (252443) | Leatherworking [crafted] | sim-verified (30.5 DPS) | yes | Wisdom's Leather Boots (252444, -0.26 DPS) [crafted]; Spidersilk Boots (4320, -0.46 DPS, sim-verified) [crafted]; Black Whelp Slippers (252424, -0.51 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 8.9 spell_power points (1.14 DPS) | yes | Lavishly Jeweled Ring (1156, -0.38 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.76 DPS) [world_drop]; Band of Purification (12996, -0.84 DPS) [world_drop] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 7.0 spell_power points (0.90 DPS) | yes | Volcanic Rock Ring (12053, -0.52 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.55 DPS, sim-verified) [dungeon]; Band of Purification (12996, -0.60 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | 31.9 spell_power points (4.10 DPS) | yes | Emberstone Staff (5201, -1.02 DPS) [dungeon]; Staff of the Blessed Seer (2271, -1.09 DPS, sim-verified) [dungeon]; Staff of Westfall (2042, -1.11 DPS) [quest] |
| off_hand | Seedcloud Buckler (6630) | Wailing Caverns: Verdan the Everliving [dungeon] | 9.9 spell_power points (1.27 DPS) | yes | Mathystra Channeling Orb (281291, -0.64 DPS, sim-verified) [quest]; Owl's Disk (4822, -0.77 DPS) [vendor]; Bristlebark Buckler (15894, -0.77 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 20:** head: Totemic Leather Hood; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Stormrider's Leather Armor; wrist: Mindthrust Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Stormrider's Leather Pants; feet: Stormrider's Leather Boots; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Evocator's Blade; off_hand: Seedcloud Buckler

No-known-source sample (15 of 260, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers

### Band 30 (dwarf, 5500331220000000-000000000000000000-0000000000000000)

Set DPS (verified): 61.0. Weights run: 6.3s. Verify run: 12.0s. 448 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.675 ± 0.011, crit=0.042 ± 0.001 per rating point (14 rating = 1%, 0.591 per %), hit=0.145 ± 0.005 per rating point (10 rating = 1%, 1.450 per %), spell_haste=-1.542 ± 0.114, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.875 ± 0.002, spirit=not significant (-0.093 ± 0.049), mp5=not significant (0.121 ± 0.081)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 12.8 spell_power points (2.18 DPS) | yes | Holy Shroud (2721, -0.30 DPS) [world_drop]; Enduring Cap (3020, -0.33 DPS) [world_drop]; Totemic Leather Helm (252456, -0.51 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 4.1 spell_power points (0.69 DPS) | yes | Crystal Starfire Medallion (5003, -0.23 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.23 DPS) [world_drop]; Seal of the Expedition (271920, -0.52 DPS, sim-verified) [rep] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.1 spell_power points (2.57 DPS) | yes | Death Speaker Mantle (6685, -0.60 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.68 DPS) [world_drop]; Invoker's Mantle (215365, -0.80 DPS) [crafted] |
| back | Expeditionary's Cape (271923) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, -0.07 DPS) [crafted]; Resilient Cape (14400, -0.23 DPS) [world_drop]; Vine Pruner's Cloak (279835, -0.81 DPS, sim-verified) [quest] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Mechbuilder's Overalls (9508, -0.17 DPS) [dungeon]; Green Silk Armor (7065, -0.22 DPS, sim-verified) [crafted]; Guardian Armor (4256, -0.23 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.4 spell_power points (1.60 DPS) | yes | Nightsky Wristbands (6407, -0.91 DPS) [world_drop]; Mindthrust Bracers (1974, -0.96 DPS) [dungeon]; Glowing Magical Bracelets (13106, -1.99 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 7.7 spell_power points (1.32 DPS) | yes | Truefaith Gloves (7049, -0.12 DPS) [crafted]; Gloves of Insight (9698, -0.12 DPS) [quest]; Stormrider's Leather Gloves (252498, -0.47 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 14.1 spell_power points (2.40 DPS) | yes | Moss Cinch (6911, -0.35 DPS) [dungeon]; Belt of Arugal (6392, -0.52 DPS) [dungeon]; Highlander's Cloth Girdle (20099, -0.67 DPS, sim-verified) [rep] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 14.1 spell_power points (2.40 DPS) | yes | Guardian Pants (5962, -0.11 DPS) [crafted]; Totemic Leather Leggings (252458, -0.34 DPS) [crafted]; Dark Ritual Leggings (270031, -0.55 DPS, sim-verified) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.7 spell_power points (2.00 DPS) | yes | Spidersilk Boots (4320, -0.35 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.40 DPS) [crafted]; Acidic Walkers (9454, -0.44 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.4 spell_power points (1.60 DPS) | yes | Minor Channeling Ring (1449, -0.52 DPS) [quest]; Lorekeeper's Ring (20431, -0.71 DPS) [rep] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Minor Channeling Ring (1449, -0.00 DPS) [quest]; Lorekeeper's Ring (20431, -0.19 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -0.68 DPS) [dungeon]; Tendonscraper Dagger (273053, -0.68 DPS) [world]; Glimmering Staff (249392, -2.59 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 11.1 spell_power points (1.89 DPS) | yes | Seedcloud Buckler (6630, -1.01 DPS, sim-verified) [dungeon]; Heart of Agamaggan (6694, -1.03 DPS) [dungeon]; Satyr's Rod (15962, -1.19 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Expeditionary's Cape; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Skycaller's Leather Belt; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight

No-known-source sample (15 of 448, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 40 (dwarf, 5500331220001351-000000000000000000-0000000000000000)

Set DPS (verified): 109.7. Weights run: 7.8s. Verify run: 18.8s. 696 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.634 ± 0.022, crit=0.208 ± 0.006 per rating point (14 rating = 1%, 2.914 per %), hit=0.393 ± 0.018 per rating point (10 rating = 1%, 3.930 per %), spell_haste=-3.908 ± 0.320, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.885 ± 0.002, spirit=-1.302 ± 0.209, mp5=-4.053 ± 0.350

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellpower Goggles Xtreme (10502, +0.00 DPS) [crafted]; Big Voodoo Mask (8201, -0.22 DPS) [crafted]; Augural Shroud (2620, -0.35 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.1 spell_power points (2.92 DPS) | yes | Triune Amulet (7722, -0.98 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.39 DPS) [quest]; Seal of the Expedition (271919, -1.85 DPS) [rep] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.2 spell_power points (3.69 DPS) | yes | Green Silken Shoulders (7057, -0.06 DPS) [crafted]; Bloodmage Mantle (7684, -0.13 DPS) [dungeon]; Berylline Pads (4197, -0.46 DPS) [quest] |
| back | Expeditionary's Cape (271922) | Theramore Expeditionary Force [rep] | sim-verified (109.7 DPS) | yes | Long Silken Cloak (4326, +0.00 DPS) [crafted]; Icy Cloak (4327, +0.00 DPS) [crafted]; Guardian Cloak (5965, +0.00 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | sim-verified (109.7 DPS) | yes | Dreamweave Vest (10021, -0.51 DPS) [crafted]; Robe of Power (7054, -1.02 DPS) [crafted]; Elemental Raiment (9434, -1.16 DPS) [world_drop] |
| wrist | Green Whelp Bracers (7386) | Leatherworking [crafted] | 13.0 spell_power points (3.15 DPS) | yes | Guardian Leather Bracers (4260, -0.78 DPS) [crafted]; Turtle Scale Bracers (8198, -0.94 DPS) [crafted]; Arcane Runed Bracers (4744, -0.98 DPS) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamweave Gloves (10019, -2.98 DPS) [crafted]; Red Mageweave Gloves (10018, -3.75 DPS) [crafted]; Skycaller's Leather Gloves (252529, -3.91 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Highlander's Cloth Girdle (20098, +0.00 DPS) [rep]; Skycaller's Leather Belt (252522, -0.65 DPS) [crafted]; Gilded Cord (254037, -0.83 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.6 spell_power points (5.23 DPS) | yes | Crimson Silk Pantaloons (7062, -1.30 DPS) [crafted]; Dark Ritual Leggings (270031, -1.84 DPS) [quest]; Stormrider's Leather Pants (252502, -1.89 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.81 DPS) | yes | Skycaller's Leather Shoes (252532, -1.59 DPS) [crafted]; Skycaller's Mail Boots (252563, -1.59 DPS) [crafted]; Mender's Leather Shoes (252533, -2.80 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | sim-verified (109.7 DPS) | yes | Snake Hoop (6750, +0.00 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.24 DPS) [quest]; Reedknot Ring (9622, -0.48 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | sim-verified (109.7 DPS) | yes | Snake Hoop (6750, -0.97 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.25 DPS) [quest]; Reedknot Ring (9622, -1.49 DPS) [quest] |
| trinket1 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-verified (109.7 DPS) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (109.7 DPS) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -6.74 DPS) [dungeon]; Hypnotic Blade (7714, -7.14 DPS) [dungeon]; Glimmering Staff (249392, -9.10 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Expeditionary's Cape; chest: Robe of the Magi; wrist: Green Whelp Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Lorekeeper's Ring; finger2: Philanthropist's Ring; trinket1: Enriched Seal; trinket2: Ankh of Life; main_hand: Staff of Jordan

No-known-source sample (15 of 696, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (dwarf, 5500331220001351-050000000000000000-0500000000000000)

Set DPS (verified): 149.5. Weights run: 7.9s. Verify run: 21.1s. 893 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.732 ± 0.015, crit=0.202 ± 0.006 per rating point (14 rating = 1%, 2.824 per %), hit=0.343 ± 0.015 per rating point (10 rating = 1%, 3.433 per %), spell_haste=-2.978 ± 0.179, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.882 ± 0.002, spirit=0.420 ± 0.043, mp5=1.342 ± 0.109

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 37.3 spell_power points (6.56 DPS) | yes | Red Mageweave Headband (10033, -0.65 DPS) [crafted]; Dreamweave Circlet (10041, -0.70 DPS) [crafted]; Charged Scorpid Helm (252581, -1.66 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.9 spell_power points (2.26 DPS) | yes | Mindburst Medallion (11196, -0.06 DPS) [quest]; Horizon Choker (13085, -0.16 DPS) [world_drop]; Triune Amulet (7722, -0.34 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 33.0 spell_power points (5.80 DPS) | yes | Rotgrip Mantle (17732, -1.20 DPS) [dungeon]; Kentic Amice (11624, -1.22 DPS) [dungeon]; Lead Surveyor's Mantle (11842, -3.09 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Expeditionary's Cape (271921, -1.50 DPS) [rep]; Big Voodoo Cloak (8216, -1.64 DPS) [crafted]; Runecloth Cloak (13860, -1.68 DPS, sim-verified) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Dreamweave Vest (10021, -0.56 DPS) [crafted]; Flamestrider Robes (11747, -0.70 DPS) [dungeon]; Feathered Breastplate (8349, -0.83 DPS, sim-verified) [crafted] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 17.3 spell_power points (3.04 DPS) | yes | Skycaller's Leather Bracers (252542, -0.02 DPS) [crafted]; Skycaller's Mail Bracers (252571, -0.02 DPS) [crafted]; Nethergeld Cuffs (254061, -0.54 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodfire Talons (12464, -1.30 DPS) [dungeon]; Skycaller's Leather Gauntlets (252550, -1.35 DPS) [crafted]; Elven Spirit Claws (2564, -2.71 DPS, sim-verified) [world_drop] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 24.8 spell_power points (4.35 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Ban'thok Sash (11662, -0.23 DPS) [dungeon]; Satyrmane Sash (17755, -0.61 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | 30.3 spell_power points (5.33 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -1.16 DPS) [dungeon]; Jinxed Hoodoo Kilt (9474, -1.23 DPS) [dungeon]; Senior Designer's Pantaloons (11841, -4.17 DPS, sim-verified) [dungeon] |
| feet | Skycaller's Leather Boots (252471) | Leatherworking [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Skycaller's Mail Sabatons (252577, +0.00 DPS) [crafted]; Stormcloth Boots (10039, -0.17 DPS) [crafted]; Earthen Silk Slippers (254013, -0.74 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 17.4 spell_power points (3.05 DPS) | yes | Brainlash (6440, -0.75 DPS) [dungeon]; Band of the Unicorn (7553, -0.77 DPS) [world_drop]; Cyclopean Band (11824, -5.54 DPS, sim-verified) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | sim-verified (149.5 DPS) | yes | Cyclopean Band (11824, +0.00 DPS) [dungeon]; Brainlash (6440, -0.10 DPS) [dungeon]; Band of the Unicorn (7553, -0.12 DPS) [world_drop] |
| trinket1 | Enriched Seal (271903) | Theramore Expeditionary Force [rep] | sim-verified (+4.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -0.12 DPS) [world_drop]; Ankh of Life (1713, -0.29 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.44 DPS) [quest] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Uther's Strength (11302, -1.48 DPS, sim-verified) [world_drop] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -2.17 DPS) [dungeon]; Spellshifter Rod (9527, -2.31 DPS) [quest]; Glowing Brightwood Staff (812, -3.01 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; waist: Skycaller's Leather Waistguard; legs: Spellshock Leggings; feet: Skycaller's Leather Boots; finger1: Lorekeeper's Ring; trinket1: Enriched Seal; trinket2: Thunderbrew's Boot Flask

No-known-source sample (15 of 893, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (dwarf, 5500331220001351-050000000000000000-0530520000000000)

Set DPS (verified): 271.7. Weights run: 7.6s. Verify run: 66.0s. 2086 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.976 ± 0.035, crit=0.342 ± 0.009 per rating point (14 rating = 1%, 4.794 per %), hit=0.565 ± 0.033 per rating point (10 rating = 1%, 5.654 per %), spell_haste=not significant (-0.877 ± 0.386), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.863 ± 0.002, spirit=not significant (0.848 ± 0.288), mp5=not significant (1.574 ± 0.394)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 54.4 spell_power points (9.72 DPS) | yes | Blue Dragonscale Helm (252604, -1.09 DPS, sim-verified) [crafted]; The Postmaster's Band (13390, -1.35 DPS) [dungeon]; Dreadmist Mask (16698, -1.40 DPS) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jeweled Amulet of Cainwyn (1443, +0.00 DPS) [world_drop]; Lady Maye's Pendant (14558, +0.00 DPS) [world_drop]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Theramore Shoulderpads (271953, -0.53 DPS) [rep]; Theramore Shoulders (271954, -0.53 DPS) [rep]; Theramore Shoulderguards (272957, -5.14 DPS, sim-verified) [rep] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of Earth and Sky (279270, -0.83 DPS) [crafted]; Crystalline Threaded Cape (20697, -0.99 DPS) [world]; Spritecaster Cape (11623, -2.14 DPS, sim-verified) [dungeon] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Feathered Breastplate (8349, -1.31 DPS) [crafted]; Magister's Robes (16688, -1.85 DPS) [dungeon]; Alanna's Embrace (13314, -5.72 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 35.7 spell_power points (6.38 DPS) | yes | Spitfire Bracers (20481, -2.02 DPS) [crafted]; Bindings of Elements (16671, -2.28 DPS) [dungeon]; Wyrmthalak's Shackles (13958, -2.54 DPS) [quest] |
| hands | Theramore Handguards (271950) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of Power (13253, +0.00 DPS) [dungeon]; Gauntlets of Elements (16672, -0.37 DPS) [dungeon]; Dracorian Gauntlets (13344, -0.38 DPS) [dungeon] |
| waist | Stormseeker's Girdle (272399) | Pix Xizzix [vendor] | 39.4 spell_power points (7.04 DPS) | yes | Sash of the Windreaver (18676, -0.29 DPS) [world]; Barrage Girdle (18721, -0.98 DPS) [dungeon]; Crystallized Girdle (12606, -1.21 DPS) [dungeon] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Leggings of Arcana (12756, +0.00 DPS) [quest]; Sentinel's Silk Leggings (22752, -0.95 DPS) [rep] |
| feet | Wolfrunner Shoes (13101) | World drop [world_drop] | sim-verified (271.7 DPS) | yes | Dragonrider Boots (18102, +0.00 DPS) [dungeon]; Mooncloth Boots (15802, -0.07 DPS) [crafted]; Omnicast Boots (11822, -0.24 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.30 DPS) [quest]; Seal of Rivendare (13345, -1.32 DPS) [dungeon]; Naglering (11669, -9.21 DPS, sim-verified) [dungeon] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, +0.00 DPS) [quest]; Seal of Rivendare (13345, -0.64 DPS) [dungeon]; Blessed Band of Light (272407, -0.65 DPS) [vendor] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -0.33 DPS) [world_drop]; Shard of the Splithooves (10659, -0.46 DPS) [quest]; Second Wind (11819, -6.19 DPS, sim-verified) [dungeon] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Second Wind (11819, +0.00 DPS) [dungeon]; Serenity Field (272439, -1.43 DPS) [vendor]; Ankh of Life (1713, -2.11 DPS) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Staff of Jordan (873, -2.62 DPS) [world_drop]; Ironbark Staff (20069, -5.93 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tome of the Lost (22253, +0.00 DPS) [dungeon]; Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Barricade (235473, +0.00 DPS) [vendor] |
| ranged | Burning Totem (272433) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of Thunder (228176, -2.85 DPS, sim-verified) [vendor] |

**New at 60:** head: Living Crown; neck: Archlight Talisman; back: Arcanoweave Cloak; chest: Ironfeather Breastplate; wrist: Dryad's Wrist Bindings; hands: Theramore Handguards; waist: Stormseeker's Girdle; legs: Ironfeather Leggings; feet: Wolfrunner Shoes; finger1: Maiden's Circle; finger2: Channeler's Ring; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Hand of Edward the Odd; off_hand: Therazane's Touch; ranged: Burning Totem

No-known-source sample (15 of 2086, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60, raid preset (dwarf, 5500331220001351-055000000000000000-0530020000000000)

Set DPS (verified): 736.6. Weights run: 4.9s. Verify run: 55.8s. 2086 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.702 ± 0.026, crit=0.500 ± 0.011 per rating point (14 rating = 1%, 7.003 per %), hit=0.922 ± 0.053 per rating point (10 rating = 1%, 9.219 per %), spell_haste=3.771 ± 0.522, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.875 ± 0.002, spirit=not significant (0.546 ± 0.234), mp5=not significant (0.783 ± 0.322)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blue Dragonscale Helm (252604) | Leatherworking [crafted] | 46.8 spell_power points (21.76 DPS) | yes | Crimson Felt Hat (18727, -3.19 DPS) [dungeon]; Helm of the New Moon (22407, -3.80 DPS) [dungeon]; Living Crown (252561, -5.13 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon]; Lady Maye's Pendant (14558, -0.11 DPS) [world_drop] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mantle of the Timbermaw (19050, -0.13 DPS) [crafted]; Burial Shawl (18681, -1.01 DPS) [dungeon]; Theramore Shoulderguards (272957, -10.18 DPS, sim-verified) [rep] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -4.34 DPS) [dungeon]; Cloak of Earth and Sky (279270, -4.67 DPS) [crafted]; Crystalline Threaded Cape (20697, -5.51 DPS, sim-verified) [world] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Feathered Breastplate (8349, -2.60 DPS) [crafted]; Vest of Elements (16666, -3.56 DPS) [dungeon]; Alanna's Embrace (13314, -11.13 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 31.4 spell_power points (14.60 DPS) | yes | Spitfire Bracers (20481, -5.67 DPS) [crafted]; Bindings of Elements (16671, -5.85 DPS) [dungeon]; Runecloth Cuffs (254123, -6.23 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of Power (13253, +0.00 DPS) [dungeon]; Theramore Handguards (271950, -0.94 DPS) [rep]; Gauntlets of Elements (16672, -1.26 DPS) [dungeon] |
| waist | Stormseeker's Girdle (272399) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sash of the Windreaver (18676, +0.00 DPS) [world]; Barrage Girdle (18721, -1.09 DPS) [dungeon]; Ban'thok Sash (11662, -1.81 DPS) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentinel's Lizardhide Pants (237817, +0.00 DPS) [vendor]; Ironfeather Leggings (252486, +0.00 DPS) [crafted]; Leggings of Arcana (12756, -2.15 DPS) [quest] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (736.6 DPS) | yes | Dragonrider Boots (18102, +0.00 DPS) [dungeon]; Earthen Silk Slippers (254013, -0.28 DPS) [crafted]; Greaves of Withering Despair (22240, -0.55 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.51 DPS) [quest]; Channeler's Ring (272406, -1.29 DPS) [vendor]; Naglering (11669, -11.94 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, +0.00 DPS) [quest]; Channeler's Ring (272406, -0.08 DPS) [vendor]; Deep Rooted Ring (19109, -2.12 DPS) [world] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -4.57 DPS, sim-verified) [vendor]; Burst of Knowledge (11832, -4.65 DPS) [dungeon]; Ankh of Life (1713, -7.18 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Ankh of Life (1713, -0.81 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -6.18 DPS, sim-verified) [quest] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -0.64 DPS) [pvp]; Staff of Jordan (873, -16.28 DPS, sim-verified) [world_drop]; Kindling Stave (11750, -17.24 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Thunder (228176) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of the Storm (272432, -4.94 DPS, sim-verified) [world_drop] |

**New at 60:** head: Blue Dragonscale Helm; neck: Archlight Talisman; back: Arcanoweave Cloak; chest: Ironfeather Breastplate; wrist: Dryad's Wrist Bindings; waist: Stormseeker's Girdle; legs: Sentinel's Silk Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Blessed Band of Light; trinket1: Weakness Analyzer; trinket2: Frozen Heart of the Mountain; main_hand: Ironbark Staff; ranged: Totem of Thunder

No-known-source sample (15 of 2086, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

## Horde

### Band 20 (orc, 5500100000000000-000000000000000000-0000000000000000)

Set DPS (verified): 30.2. Weights run: 7.2s. Verify run: 6.5s. 249 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.986 ± 0.016, crit=0.094 ± 0.003 per rating point (14 rating = 1%, 1.309 per %), hit=0.325 ± 0.013 per rating point (10 rating = 1%, 3.246 per %), spell_haste=not significant (0.266 ± 0.138), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.867 ± 0.002, spirit=not significant (0.388 ± 0.179), mp5=1.929 ± 0.286

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Hood (252448) | Leatherworking [crafted] | 9.0 spell_power points (1.16 DPS) | yes | Trapper's Leather Hood (252505, -0.14 DPS) [crafted]; Shadow Goggles (4373, -0.22 DPS) [crafted]; Wisdom's Leather Hood (252507, -0.39 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 1.2 spell_power points (0.15 DPS) | yes | Scholarly Pendant (277203, -0.05 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.9 spell_power points (1.78 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.63 DPS) [crafted]; Slime-encrusted Pads (6461, -1.04 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -1.27 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pearl-clasped Cloak (5542, -0.08 DPS) [crafted]; Sanguine Cape (14376, -0.21 DPS) [world_drop]; Battle Healer's Cloak (20427, -0.23 DPS) [rep] |
| chest | Stormrider's Leather Armor (252492) | Leatherworking [crafted] | 14.1 spell_power points (1.81 DPS) | yes | Wisdom's Leather Armor (252493, -0.39 DPS) [crafted]; Filigreed Pristine Gown (253901, -0.39 DPS) [crafted]; Totemic Leather Armor (252435, -0.66 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 10.7 spell_power points (1.38 DPS) | yes | Owl Bracers (4796, -0.74 DPS) [vendor]; Bright Bracers (3647, -0.87 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.99 DPS) [quest] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 8.9 spell_power points (1.15 DPS) | yes | Serpent Gloves (5970, -0.25 DPS) [dungeon]; Pristine Gloves (253913, -0.25 DPS) [crafted]; Wisdom's Leather Gloves (252499, -0.26 DPS) [crafted] |
| waist | Pristine Sash (253925) (or Stormrider's Leather Belt (252432)) | Tailoring [crafted] | 7.9 spell_power points (1.02 DPS) | yes | Stormrider's Leather Belt (252432, +0.00 DPS) [crafted]; Keller's Girdle (2911, -0.01 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.13 DPS) [crafted] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 17.5 spell_power points (2.24 DPS) | yes | Wisdom's Leather Pants (252503, -0.39 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.51 DPS) [crafted]; Totemic Leather Pants (252446, -0.84 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | sim-verified (30.2 DPS) | yes | Stormrider's Leather Boots (252443, -0.00 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.26 DPS) [crafted]; Black Whelp Slippers (252424, -0.51 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 8.9 spell_power points (1.14 DPS) | yes | Loop of Sacrifice (281673, -0.50 DPS) [quest]; Volcanic Rock Ring (12053, -0.76 DPS) [world_drop]; Band of Purification (12996, -0.84 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | sim-verified (30.2 DPS) | yes | Loop of Sacrifice (281673, -0.13 DPS) [quest]; Volcanic Rock Ring (12053, -0.38 DPS) [world_drop]; Band of Purification (12996, -0.46 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crescent Staff (6505, +0.00 DPS) [quest]; Staff of the Blessed Seer (2271, -0.00 DPS) [dungeon]; Gnarled Necromancer's Staff (251534, -0.26 DPS) [quest] |
| off_hand | Seedcloud Buckler (6630) | Wailing Caverns: Verdan the Everliving [dungeon] | 9.9 spell_power points (1.27 DPS) | yes | Owl's Disk (4822, -0.77 DPS) [vendor]; Forest Buckler (6383, -0.77 DPS) [world_drop]; Bristlebark Buckler (15894, -0.77 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 20:** head: Totemic Leather Hood; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Stormrider's Leather Armor; wrist: Mindthrust Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Stormrider's Leather Pants; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Evocator's Blade; off_hand: Seedcloud Buckler

No-known-source sample (15 of 249, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield; 14389 Durability Shoulderpads

### Band 30 (orc, 5500331220000000-000000000000000000-0000000000000000)

Set DPS (verified): 61.1. Weights run: 6.3s. Verify run: 11.9s. 442 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.669 ± 0.011, crit=0.087 ± 0.002 per rating point (14 rating = 1%, 1.217 per %), hit=0.295 ± 0.010 per rating point (10 rating = 1%, 2.948 per %), spell_haste=-1.050 ± 0.191, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.853 ± 0.002, spirit=not significant (0.127 ± 0.102), mp5=1.557 ± 0.140

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 12.7 spell_power points (3.15 DPS) | yes | Totemic Leather Helm (252456, -0.17 DPS) [crafted]; Holy Shroud (2721, -0.23 DPS) [world_drop]; Enduring Cap (3020, -0.27 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 4.0 spell_power points (1.00 DPS) | yes | Darkspear Warding Pendant (272075, -0.17 DPS) [vendor]; Kaleidoscope Chain (13084, -0.21 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.24 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.0 spell_power points (3.73 DPS) | yes | Death Speaker Mantle (6685, -0.41 DPS) [dungeon]; Magician's Mantle (12998, -0.99 DPS) [world_drop]; Invoker's Mantle (215365, -1.16 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windsong Drape (15468, -0.06 DPS) [quest]; Hillman's Cloak (3719, -0.18 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.30 DPS) [crafted] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 18.0 spell_power points (4.48 DPS) | yes | Green Silk Armor (7065, -0.09 DPS) [crafted]; Mechbuilder's Overalls (9508, -0.50 DPS) [dungeon]; Guardian Armor (4256, -0.58 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 13.7 spell_power points (3.39 DPS) | yes | Mindthrust Bracers (1974, -1.49 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -2.07 DPS) [world_drop]; Nightsky Wristbands (6407, -2.30 DPS) [world_drop] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | 10.3 spell_power points (2.57 DPS) | yes | Jutebraid Gloves (10654, -0.25 DPS) [quest]; Stormrider's Leather Gloves (252498, -0.66 DPS) [crafted]; Truefaith Gloves (7049, -0.83 DPS) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 14.5 spell_power points (3.60 DPS) | yes | Defiler's Cloth Girdle (20164, -0.38 DPS) [rep]; Moss Cinch (6911, -0.63 DPS) [dungeon]; Belt of Arugal (6392, -0.71 DPS) [dungeon] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 14.5 spell_power points (3.60 DPS) | yes | Dark Ritual Leggings (270031, -0.13 DPS) [quest]; Guardian Pants (5962, -0.29 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.34 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 12.2 spell_power points (3.03 DPS) | yes | Acidic Walkers (9454, -0.33 DPS) [dungeon]; Spidersilk Boots (4320, -0.62 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.71 DPS) [crafted] |
| finger1 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 10.7 spell_power points (2.65 DPS) | yes | Advisor's Ring (20426, -0.63 DPS) [rep]; Darkspear Signet (272071, -1.10 DPS) [vendor] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Advisor's Ring (19521, -0.24 DPS, sim-verified) [rep]; Darkspear Signet (272071, -0.77 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.12 DPS) [dungeon]; Tendonscraper Dagger (273053, -1.12 DPS) [world]; Glimmering Staff (249392, -2.51 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 11.0 spell_power points (2.73 DPS) | yes | Seedcloud Buckler (6630, -0.53 DPS, sim-verified) [dungeon]; Witch's Finger (16887, -1.41 DPS) [quest]; Heart of Agamaggan (6694, -1.49 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Darkspear Raider's Cloak; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; feet: Gilded Slippers; finger1: Ring of Power Regulation; finger2: Snake Hoop; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight

No-known-source sample (15 of 442, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5255 Quilboar Tomahawk

### Band 40 (orc, 5500331220001351-000000000000000000-0000000000000000)

Set DPS (verified): 109.0. Weights run: 7.8s. Verify run: 18.1s. 679 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.634 ± 0.022, crit=0.208 ± 0.006 per rating point (14 rating = 1%, 2.914 per %), hit=0.393 ± 0.018 per rating point (10 rating = 1%, 3.930 per %), spell_haste=-3.908 ± 0.320, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.885 ± 0.002, spirit=-1.302 ± 0.209, mp5=-4.053 ± 0.350

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellpower Goggles Xtreme (10502, +0.00 DPS) [crafted]; Big Voodoo Mask (8201, -0.22 DPS) [crafted]; Augural Shroud (2620, -0.35 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (109.0 DPS) | yes | Triune Amulet (7722, -0.98 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.39 DPS) [quest]; Darkspear Warding Pendant (272074, -1.85 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.2 spell_power points (3.69 DPS) | yes | Green Silken Shoulders (7057, -0.06 DPS) [crafted]; Bloodmage Mantle (7684, -0.13 DPS) [dungeon]; Berylline Pads (4197, -0.46 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.52 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.53 DPS) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | sim-verified (109.0 DPS) | yes | Dreamweave Vest (10021, -0.51 DPS) [crafted]; Robe of Power (7054, -1.02 DPS) [crafted]; Elemental Raiment (9434, -1.16 DPS) [world_drop] |
| wrist | Green Whelp Bracers (7386) | Leatherworking [crafted] | 13.0 spell_power points (3.15 DPS) | yes | Guardian Leather Bracers (4260, -0.78 DPS) [crafted]; Turtle Scale Bracers (8198, -0.94 DPS) [crafted]; Radiant Silver Bracers (4545, -0.96 DPS) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamweave Gloves (10019, -2.98 DPS) [crafted]; Red Mageweave Gloves (10018, -3.75 DPS) [crafted]; Skycaller's Leather Gloves (252529, -3.91 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defiler's Cloth Girdle (20166, +0.00 DPS) [rep]; Skycaller's Leather Belt (252522, -0.65 DPS) [crafted]; Gilded Cord (254037, -0.83 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.6 spell_power points (5.23 DPS) | yes | Crimson Silk Pantaloons (7062, -1.30 DPS) [crafted]; Dark Ritual Leggings (270031, -1.84 DPS) [quest]; Stormrider's Leather Pants (252502, -1.89 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.81 DPS) | yes | Skycaller's Leather Shoes (252532, -1.59 DPS) [crafted]; Skycaller's Mail Boots (252563, -1.59 DPS) [crafted]; Mender's Leather Shoes (252533, -2.80 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.2 spell_power points (3.19 DPS) | yes | Snake Hoop (6750, -0.97 DPS) [quest]; Reedknot Ring (9622, -1.49 DPS) [quest]; Ring of Power Regulation (273030, -1.73 DPS) [dungeon] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Snake Hoop (6750, +0.00 DPS) [quest]; Reedknot Ring (9622, -0.48 DPS) [quest]; Ring of Power Regulation (273030, -0.73 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (109.0 DPS) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (109.0 DPS) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -6.74 DPS) [dungeon]; Hypnotic Blade (7714, -7.14 DPS) [dungeon]; Glimmering Staff (249392, -9.10 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Green Whelp Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Staff of Jordan

No-known-source sample (15 of 679, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 50 (orc, 5500331220001351-050000000000000000-0500000000000000)

Set DPS (verified): 149.2. Weights run: 7.9s. Verify run: 20.4s. 862 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.732 ± 0.015, crit=0.202 ± 0.006 per rating point (14 rating = 1%, 2.824 per %), hit=0.343 ± 0.015 per rating point (10 rating = 1%, 3.433 per %), spell_haste=-2.978 ± 0.179, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.882 ± 0.002, spirit=0.420 ± 0.043, mp5=1.342 ± 0.109

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 37.3 spell_power points (6.56 DPS) | yes | Dreamweave Circlet (10041, -0.70 DPS) [crafted]; Red Mageweave Headband (10033, -1.34 DPS, sim-verified) [crafted]; Charged Scorpid Helm (252581, -1.66 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.9 spell_power points (2.26 DPS) | yes | Mindburst Medallion (11196, -0.06 DPS) [quest]; Horizon Choker (13085, -0.16 DPS) [world_drop]; Triune Amulet (7722, -0.34 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 33.0 spell_power points (5.80 DPS) | yes | Rotgrip Mantle (17732, -1.20 DPS) [dungeon]; Kentic Amice (11624, -1.22 DPS) [dungeon]; Lead Surveyor's Mantle (11842, -1.83 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Runecloth Cloak (13860, -1.06 DPS) [crafted]; Darkspear Raider's Cloak (272076, -1.50 DPS) [vendor]; Deep Woodlands Cloak (19121, -1.51 DPS, sim-verified) [quest] |
| chest | Feathered Breastplate (8349) | Leatherworking [crafted] | 36.4 spell_power points (6.39 DPS) | yes | Acumen Robes (17775, -0.48 DPS) [quest]; Dreamweave Vest (10021, -1.04 DPS) [crafted]; Stone Guard's Pulsing Breastplate (220844, -1.08 DPS) [vendor] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 17.3 spell_power points (3.04 DPS) | yes | Skycaller's Mail Bracers (252571, -0.02 DPS) [crafted]; Nethergeld Cuffs (254061, -0.54 DPS) [crafted]; Skycaller's Leather Bracers (252542, -1.14 DPS, sim-verified) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodfire Talons (12464, -1.30 DPS) [dungeon]; Skycaller's Leather Gauntlets (252550, -1.35 DPS) [crafted]; Elven Spirit Claws (2564, -2.88 DPS, sim-verified) [world_drop] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 24.8 spell_power points (4.35 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Ban'thok Sash (11662, -0.23 DPS) [dungeon]; Satyrmane Sash (17755, -0.61 DPS) [dungeon] |
| legs | Stone Guard's Pulsing Legplates (220847) | Lady Palanseer [vendor] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Senior Designer's Pantaloons (11841, -0.97 DPS) [dungeon]; Spellshock Leggings (9484, -1.05 DPS, sim-verified) [dungeon]; Kilt of the Atal'ai Prophet (10807, -1.14 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 26.9 spell_power points (4.73 DPS) | yes | Skycaller's Leather Boots (252471, -0.68 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -0.68 DPS) [crafted]; Stormcloth Boots (10039, -0.85 DPS) [crafted] |
| finger1 | Advisor's Ring (19519) | Warsong Outriders [rep] | 17.4 spell_power points (3.05 DPS) | yes | Philanthropist's Ring (281635, -0.65 DPS) [quest]; Brainlash (6440, -0.75 DPS) [dungeon]; Band of the Unicorn (7553, -0.77 DPS) [world_drop] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.8 spell_power points (2.78 DPS) | yes | Brainlash (6440, -0.48 DPS) [dungeon]; Band of the Unicorn (7553, -0.49 DPS) [world_drop]; Philanthropist's Ring (281635, -1.55 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -0.29 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.44 DPS) [quest]; Uther's Strength (11302, -1.85 DPS, sim-verified) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -2.17 DPS) [dungeon]; Spellshifter Rod (9527, -2.31 DPS) [quest]; Glowing Brightwood Staff (812, -3.56 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Feathered Breastplate; wrist: Runic Leather Bracers; waist: Skycaller's Leather Waistguard; legs: Stone Guard's Pulsing Legplates; finger1: Advisor's Ring; finger2: Cyclopean Band; trinket1: Darkspear Voodoo Seal; trinket2: Frozen Heart of the Mountain

No-known-source sample (15 of 862, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 60 (orc, 5500331220001351-050000000000000000-0530520000000000)

Set DPS (verified): 272.2. Weights run: 7.6s. Verify run: 67.6s. 2043 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.976 ± 0.035, crit=0.342 ± 0.009 per rating point (14 rating = 1%, 4.794 per %), hit=0.565 ± 0.033 per rating point (10 rating = 1%, 5.654 per %), spell_haste=not significant (-0.877 ± 0.386), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.863 ± 0.002, spirit=not significant (0.848 ± 0.288), mp5=not significant (1.574 ± 0.394)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 54.4 spell_power points (9.72 DPS) | yes | Blue Dragonscale Helm (252604, -1.11 DPS) [crafted]; Warlord's Mail Helm (231663, -1.27 DPS) [pvp]; The Postmaster's Band (13390, -1.35 DPS) [dungeon] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | 27.0 spell_power points (4.83 DPS) | yes | Jeweled Amulet of Cainwyn (1443, -0.17 DPS) [world_drop]; Diana's Pearl Necklace (22403, -0.81 DPS) [dungeon]; Archlight Talisman (15856, -0.94 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Mail Spaulders (231659, +0.00 DPS) [pvp]; Darkspear Shoulderpads (272103, -0.53 DPS) [vendor]; Darkspear Shoulderguards (272958, -5.44 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of Earth and Sky (279270, -0.83 DPS) [crafted]; Crystalline Threaded Cape (20697, -0.99 DPS) [world]; Spritecaster Cape (11623, -4.05 DPS, sim-verified) [dungeon] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Feathered Breastplate (8349, -1.31 DPS) [crafted]; Warlord's Mail Breastplate (231662, -1.63 DPS) [vendor]; Alanna's Embrace (13314, -6.05 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 35.7 spell_power points (6.38 DPS) | yes | Spitfire Bracers (20481, -2.02 DPS) [crafted]; Bindings of Elements (16671, -2.28 DPS) [dungeon]; Wyrmthalak's Shackles (13958, -2.54 DPS) [quest] |
| hands | Raider Handguards (272101) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of Power (13253, +0.00 DPS) [dungeon]; General's Mail Gauntlets (231660, +0.00 DPS) [pvp]; Gauntlets of Elements (16672, -0.37 DPS) [dungeon] |
| waist | Stormseeker's Girdle (272399) | Pix Xizzix [vendor] | 39.4 spell_power points (7.04 DPS) | yes | Barrage Girdle (18721, -0.98 DPS) [dungeon]; General's Mail Waistband (16575, -1.17 DPS) [pvp]; Sash of the Windreaver (18676, -3.58 DPS, sim-verified) [world] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Mail Leggings (231664, -0.74 DPS) [pvp]; Outrider's Silk Leggings (22747, -0.95 DPS) [rep]; Leggings of Arcana (12756, -2.07 DPS, sim-verified) [quest] |
| feet | Wolfrunner Shoes (13101) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dragonrider Boots (18102, +0.00 DPS) [dungeon]; General's Mail Sabatons (231661, +0.00 DPS) [vendor]; Mooncloth Boots (15802, -0.07 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.30 DPS) [quest]; Seal of Rivendare (13345, -1.32 DPS) [dungeon]; Naglering (11669, -11.28 DPS, sim-verified) [dungeon] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-verified (272.2 DPS) | yes | Eye of Orgrimmar (12545, +0.00 DPS) [quest]; Seal of Rivendare (13345, -0.64 DPS) [dungeon]; Blessed Band of Light (272407, -0.65 DPS) [vendor] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -0.33 DPS) [world_drop]; Shard of the Splithooves (10659, -0.46 DPS) [quest]; Second Wind (11819, -6.74 DPS, sim-verified) [dungeon] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Second Wind (11819, +0.00 DPS) [dungeon]; Serenity Field (272439, -1.43 DPS) [vendor]; Ankh of Life (1713, -2.11 DPS) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Staff of Jordan (873, -2.62 DPS) [world_drop]; Ironbark Staff (20220, -8.30 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Barricade (235474, +0.00 DPS) [vendor]; Tome of the Lost (22253, -2.39 DPS, sim-verified) [dungeon] |
| ranged | Burning Totem (272433) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of Thunder (228176, -4.56 DPS, sim-verified) [vendor] |

**New at 60:** head: Living Crown; neck: Lady Maye's Pendant; back: Arcanoweave Cloak; chest: Ironfeather Breastplate; wrist: Dryad's Wrist Bindings; hands: Raider Handguards; waist: Stormseeker's Girdle; legs: Ironfeather Leggings; feet: Wolfrunner Shoes; finger1: Maiden's Circle; finger2: Channeler's Ring; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Hand of Edward the Odd; off_hand: Therazane's Touch; ranged: Burning Totem

No-known-source sample (15 of 2043, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 60, raid preset (orc, 5500331220001351-055000000000000000-0530020000000000)

Set DPS (verified): 735.2. Weights run: 4.9s. Verify run: 54.6s. 2043 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.234 ± 0.009, crit=0.702 ± 0.016 per rating point (14 rating = 1%, 9.827 per %), hit=1.288 ± 0.079 per rating point (10 rating = 1%, 12.879 per %), spell_haste=9.068 ± 0.644, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.879 ± 0.001, spirit=not significant (0.022 ± 0.022), mp5=not significant (0.056 ± 0.032)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blue Dragonscale Helm (252604) | Leatherworking [crafted] | 42.1 spell_power points (23.54 DPS) | yes | Warlord's Mail Helm (231663, -2.56 DPS) [pvp]; Champion's Mail Helm (227162, -4.76 DPS) [pvp]; Mask of the Unforgiven (13404, -6.09 DPS, sim-verified) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon]; Emberfury Talisman (12929, -2.46 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 29.9 spell_power points (16.70 DPS) | yes | Champion's Mail Spaulders (227160, +0.00 DPS) [vendor]; Warlord's Mail Spaulders (231659, +0.00 DPS) [pvp]; Wyrmhide Spaulders (12082, -6.20 DPS, sim-verified) [quest] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -5.98 DPS, sim-verified) [world]; Spritecaster Cape (11623, -8.51 DPS) [dungeon]; Shroud of Arcane Mastery (22330, -8.55 DPS) [dungeon] |
| chest | Stone Guard's Pulsing Breastplate (220844) | Lady Palanseer [vendor] | 33.7 spell_power points (18.84 DPS) | yes | Legionnaire's Mail Breastplate (227165, +0.00 DPS) [vendor]; Warlord's Mail Breastplate (231662, +0.00 DPS) [vendor]; Tunic of the Crescent Moon (22409, -2.94 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 24.0 spell_power points (13.44 DPS) | yes | Runecloth Cuffs (254123, -5.98 DPS) [crafted]; Slashclaw Bracers (13211, -6.16 DPS) [dungeon]; Blackmist Armguards (12966, -6.23 DPS) [dungeon] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Mail Gauntlets (231660, +0.00 DPS) [pvp]; Storm Gauntlets (12632, -1.07 DPS) [crafted]; Hands of Power (13253, -2.94 DPS) [dungeon] |
| waist | Defiler's Cloth Girdle (20163) | The Defilers [rep] | sim-verified (735.2 DPS) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Sash of the Windreaver (18676, +0.00 DPS) [world]; Barrage Girdle (18721, -0.39 DPS) [dungeon] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 52.0 spell_power points (29.10 DPS) | yes | General's Mail Leggings (231664, -4.81 DPS) [pvp]; Pristine Scorpid Leggings (252606, -7.56 DPS, sim-verified) [crafted]; Skyshroud Leggings (13170, -9.04 DPS) [dungeon] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 26.2 spell_power points (14.66 DPS) | yes | Earthen Silk Slippers (254013, -1.16 DPS) [crafted]; General's Mail Sabatons (231661, -1.25 DPS) [vendor]; Omnicast Boots (11822, -1.91 DPS) [dungeon] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -1.25 DPS) [dungeon]; Maiden's Circle (13001, -1.84 DPS) [world_drop]; Naglering (11669, -12.16 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -1.05 DPS) [dungeon]; Maiden's Circle (13001, -1.64 DPS) [world_drop]; Eye of Orgrimmar (12545, -1.66 DPS) [quest] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+8.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -0.23 DPS) [crafted]; Rune of the Guard Captain (19120, -1.67 DPS) [quest]; Uther's Strength (11302, -3.36 DPS) [world_drop] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -4.76 DPS, sim-verified) [vendor]; Frozen Heart of the Mountain (249469, -5.82 DPS) [crafted]; Rune of the Guard Captain (19120, -7.26 DPS) [quest] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -2.87 DPS) [pvp]; Staff of Jordan (873, -15.65 DPS, sim-verified) [world_drop]; Whiteout Staff (19101, -18.69 DPS) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Thunder (228176) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of the Storm (272432, -4.78 DPS, sim-verified) [world_drop] |

**New at 60:** head: Blue Dragonscale Helm; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Stone Guard's Pulsing Breastplate; wrist: Dryad's Wrist Bindings; waist: Defiler's Cloth Girdle; legs: Ironfeather Leggings; feet: Greaves of Withering Despair; finger1: Blessed Band of Light; finger2: Don Julio's Band; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Totem of Thunder

No-known-source sample (15 of 2043, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

