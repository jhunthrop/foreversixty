# Leveling BiS: Elemental

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 5510000000000000-000000000000000000-0000000000000000)

Set DPS (verified): 29.2. Weights run: 6.8s. Verify run: 11.1s. 250 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.950 ± 0.017, crit=0.094 ± 0.003 per rating point (14 rating = 1%, 1.317 per %), hit=0.332 ± 0.013 per rating point (10 rating = 1%, 3.320 per %), spell_haste=3.370 ± 0.218, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.869 ± 0.002, spirit=not significant (-0.194 ± 0.184), mp5=not significant (-0.029 ± 0.289)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Hood (252448) | Leatherworking [crafted] | 9.0 spell_power points (1.10 DPS) | yes | Trapper's Leather Hood (252505, -0.17 DPS) [crafted]; Wisdom's Leather Hood (252507, -0.37 DPS) [crafted]; Pristine Circlet (253949, -0.37 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.5 spell_power points (1.66 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.44 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.17 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Sanguine Cape (14376, -0.02 DPS) [world_drop]; Black Whelp Cloak (7283, -0.12 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.28 DPS, sim-verified) [crafted] |
| chest | Stormrider's Leather Armor (252492) | Leatherworking [crafted] | 12.7 spell_power points (1.56 DPS) | yes | Wisdom's Leather Armor (252493, -0.36 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -0.37 DPS) [crafted]; Totemic Leather Armor (252435, -0.48 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Bright Bracers (3647, -0.12 DPS) [world_drop]; Inscribed Leather Bracers (3205, -0.23 DPS) [dungeon]; Owl Bracers (4796, -0.76 DPS, sim-verified) [vendor] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 8.8 spell_power points (1.08 DPS) | yes | Serpent Gloves (5970, -0.22 DPS) [dungeon]; Windfelt Gloves (5630, -0.24 DPS) [quest]; Pristine Gloves (253913, -0.24 DPS) [crafted] |
| waist | Pristine Sash (253925) (or Stormrider's Leather Belt (252432)) | Tailoring [crafted] | 7.8 spell_power points (0.96 DPS) | yes | Stormrider's Leather Belt (252432, +0.00 DPS) [crafted]; Keller's Girdle (2911, -0.02 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.12 DPS) [crafted] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 15.7 spell_power points (1.92 DPS) | yes | Wisdom's Leather Pants (252503, -0.36 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.49 DPS) [crafted]; Totemic Leather Pants (252446, -0.60 DPS) [crafted] |
| feet | Stormrider's Leather Boots (252443) | Leatherworking [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Spidersilk Boots (4320, -0.12 DPS, sim-verified) [crafted]; Wisdom's Leather Boots (252444, -0.25 DPS) [crafted]; Black Whelp Slippers (252424, -0.48 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.9 spell_power points (0.85 DPS) | yes | Volcanic Rock Ring (12053, -0.50 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.54 DPS, sim-verified) [dungeon]; Furlfeather Ring (286745, -0.72 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Lavishly Jeweled Ring (1156, -0.24 DPS, sim-verified) [dungeon]; Volcanic Rock Ring (12053, -0.26 DPS) [world_drop]; Furlfeather Ring (286745, -0.49 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | 29.9 spell_power points (3.66 DPS) | yes | Emberstone Staff (5201, -0.72 DPS) [dungeon]; Staff of the Blessed Seer (2271, -0.95 DPS, sim-verified) [dungeon]; Buzzer Blade (2169, -0.97 DPS) [dungeon] |
| off_hand | Seedcloud Buckler (6630) | Wailing Caverns: Verdan the Everliving [dungeon] | 9.7 spell_power points (1.19 DPS) | yes | Mathystra Channeling Orb (281291, -0.61 DPS) [quest]; Tear of Grief (5611, -0.64 DPS, sim-verified) [quest]; Owl's Disk (4822, -0.73 DPS) [vendor] |
| ranged | - | - |  |  |  |

**New at 20:** head: Totemic Leather Hood; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Stormrider's Leather Armor; wrist: Mindthrust Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Stormrider's Leather Pants; feet: Stormrider's Leather Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Evocator's Blade; off_hand: Seedcloud Buckler

No-known-source sample (15 of 250, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18845 Insignia of the Horde

### Band 30 (dwarf, 5510331300000000-000000000000000000-0000000000000000)

Set DPS (verified): 56.1. Weights run: 6.5s. Verify run: 10.9s. 434 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.755 ± 0.012, crit=0.099 ± 0.003 per rating point (14 rating = 1%, 1.386 per %), hit=0.352 ± 0.012 per rating point (10 rating = 1%, 3.523 per %), spell_haste=not significant (-0.482 ± 0.227), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.850 ± 0.002, spirit=not significant (0.124 ± 0.081), mp5=2.198 ± 0.216

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 13.6 spell_power points (2.86 DPS) | yes | Enduring Cap (3020, -0.13 DPS) [world_drop]; Totemic Leather Helm (252456, -0.33 DPS) [crafted]; Holy Shroud (2721, -0.38 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 4.5 spell_power points (0.96 DPS) | yes | Kaleidoscope Chain (13084, -0.21 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.24 DPS) [world_drop]; Seal of the Expedition (271920, -0.39 DPS, sim-verified) [rep] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.8 spell_power points (3.33 DPS) | yes | Death Speaker Mantle (6685, -0.31 DPS) [dungeon]; Magician's Mantle (12998, -0.84 DPS) [world_drop]; Invoker's Mantle (215365, -1.06 DPS) [crafted] |
| back | Expeditionary's Cape (271923) | Theramore Expeditionary Force [rep] | sim-verified (56.1 DPS) | yes | Hillman's Cloak (3719, -0.30 DPS) [crafted]; Repairman's Cape (9605, -0.32 DPS) [quest]; Vine Pruner's Cloak (279835, -0.41 DPS, sim-verified) [quest] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 19.1 spell_power points (4.02 DPS) | yes | Green Silk Armor (7065, -0.05 DPS) [crafted]; Mechbuilder's Overalls (9508, -0.42 DPS) [dungeon]; Guardian Armor (4256, -0.53 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 15.6 spell_power points (3.29 DPS) | yes | Mindthrust Bracers (1974, -1.24 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -2.01 DPS) [world_drop]; Nightsky Wristbands (6407, -2.25 DPS) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 8.3 spell_power points (1.75 DPS) | yes | Stormrider's Leather Gloves (252498, -0.06 DPS) [crafted]; Truefaith Gloves (7049, -0.22 DPS) [crafted]; Gloves of Insight (9698, -0.27 DPS) [quest] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 15.0 spell_power points (3.17 DPS) | yes | Highlander's Cloth Girdle (20099, -0.58 DPS, sim-verified) [rep]; Moss Cinch (6911, -0.64 DPS) [dungeon]; Belt of Arugal (6392, -0.66 DPS) [dungeon] |
| legs | Leech Pants (6910) | Blackfathom Deeps: Aku'mai [dungeon] | 16.0 spell_power points (3.37 DPS) | yes | Stormrider's Leather Pants (252502, -0.20 DPS) [crafted]; Guardian Pants (5962, -0.41 DPS) [crafted]; Dark Ritual Leggings (270031, -0.42 DPS) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 12.8 spell_power points (2.70 DPS) | yes | Acidic Walkers (9454, -0.26 DPS) [dungeon]; Spidersilk Boots (4320, -0.58 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.63 DPS) [crafted] |
| finger1 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 12.6 spell_power points (2.66 DPS) | yes | Snake Hoop (6750, -0.59 DPS) [quest]; Minor Channeling Ring (1449, -1.28 DPS) [quest]; Lavishly Jeweled Ring (1156, -1.70 DPS) [dungeon] |
| finger2 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 11.4 spell_power points (2.40 DPS) | yes | Snake Hoop (6750, -0.34 DPS) [quest]; Minor Channeling Ring (1449, -1.03 DPS) [quest]; Lavishly Jeweled Ring (1156, -1.45 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-verified (56.1 DPS) | yes | Hydrocane (9452, -0.95 DPS) [dungeon]; Tendonscraper Dagger (273053, -0.95 DPS) [world]; Glimmering Staff (249392, -2.09 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 11.5 spell_power points (2.43 DPS) | yes | Seedcloud Buckler (6630, -0.82 DPS, sim-verified) [dungeon]; Heart of Agamaggan (6694, -1.38 DPS) [dungeon]; Satyr's Rod (15962, -1.42 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Expeditionary's Cape; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Skycaller's Leather Belt; legs: Leech Pants; feet: Gilded Slippers; finger1: Ring of Power Regulation; finger2: Lorekeeper's Ring; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight

No-known-source sample (15 of 434, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 40 (dwarf, 5510331300001351-000000000000000000-0000000000000000)

Set DPS (verified): 116.6. Weights run: 6.1s. Verify run: 13.5s. 661 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.946 ± 0.012, crit=0.232 ± 0.006 per rating point (14 rating = 1%, 3.251 per %), hit=0.456 ± 0.023 per rating point (10 rating = 1%, 4.558 per %), spell_haste=not significant (-0.651 ± 0.232), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.896 ± 0.002, spirit=not significant (0.190 ± 0.049), mp5=2.061 ± 0.164

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 25.8 spell_power points (8.12 DPS) | yes | Big Voodoo Mask (8201, -1.12 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.51 DPS) [crafted]; Augural Shroud (2620, -1.68 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 14.6 spell_power points (4.59 DPS) | yes | Triune Amulet (7722, -1.74 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -2.08 DPS) [dungeon]; Seal of the Expedition (271919, -2.50 DPS) [rep] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 19.3 spell_power points (6.07 DPS) | yes | Green Silken Shoulders (7057, -0.28 DPS) [crafted]; Bloodmage Mantle (7684, -0.56 DPS) [dungeon]; Berylline Pads (4197, -0.89 DPS) [quest] |
| back | Expeditionary's Cape (271922) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Long Silken Cloak (4326, -0.14 DPS) [crafted]; Guardian Cloak (5965, -0.14 DPS) [crafted]; Blackforge Cape (6424, -0.89 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Dreamweave Vest (10021, -0.37 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.55 DPS) [crafted]; Big Voodoo Robe (8200, -2.01 DPS) [crafted] |
| wrist | Green Whelp Bracers (7386) | Leatherworking [crafted] | 15.6 spell_power points (4.92 DPS) | yes | Spidertank Oilrag (9448, -0.14 DPS) [dungeon]; Guardian Leather Bracers (4260, -1.25 DPS) [crafted]; Mindthrust Bracers (1974, -1.49 DPS) [dungeon] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamweave Gloves (10019, -3.93 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -3.94 DPS) [crafted]; Skycaller's Leather Gloves (252529, -4.23 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 21.2 spell_power points (6.67 DPS) | yes | Highlander's Cloth Girdle (20098, -1.03 DPS, sim-verified) [rep]; Gilded Cord (254037, -1.41 DPS) [crafted]; Skycaller's Leather Belt (252522, -1.50 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 25.4 spell_power points (7.98 DPS) | yes | Crimson Silk Pantaloons (7062, -1.94 DPS, sim-verified) [crafted]; Stoneweaver Leggings (9407, -2.68 DPS) [dungeon]; Stormrider's Leather Pants (252502, -2.81 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.3 spell_power points (7.97 DPS) | yes | Skycaller's Mail Boots (252563, -1.50 DPS) [crafted]; Skycaller's Leather Shoes (252532, -1.76 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -2.95 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 15.2 spell_power points (4.78 DPS) | yes | Ring of Power Regulation (273030, -0.94 DPS) [dungeon]; Snake Hoop (6750, -1.40 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.26 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.7 spell_power points (4.64 DPS) | yes | Ring of Power Regulation (273030, -0.87 DPS, sim-verified) [dungeon]; Snake Hoop (6750, -1.26 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.12 DPS) [quest] |
| trinket1 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -10.85 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -11.18 DPS) [dungeon]; Glimmering Staff (249392, -13.61 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Expeditionary's Cape; chest: Robe of the Magi; wrist: Green Whelp Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Lorekeeper's Ring; finger2: Philanthropist's Ring; trinket1: Enriched Seal; trinket2: Ankh of Life; main_hand: Staff of Jordan

No-known-source sample (15 of 661, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 50 (dwarf, 5510331300001351-050000000000000000-0500000000000000)

Set DPS (verified): 173.5. Weights run: 5.6s. Verify run: 12.3s. 848 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.795 ± 0.009, crit=0.250 ± 0.006 per rating point (14 rating = 1%, 3.503 per %), hit=0.416 ± 0.023 per rating point (10 rating = 1%, 4.156 per %), spell_haste=not significant (0.408 ± 0.253), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.877 ± 0.002, spirit=not significant (0.025 ± 0.036), mp5=1.459 ± 0.121

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 35.1 spell_power points (14.25 DPS) | yes | Red Mageweave Headband (10033, -0.09 DPS) [crafted]; Dreamweave Circlet (10041, -2.38 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -3.30 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.4 spell_power points (5.42 DPS) | yes | Mindburst Medallion (11196, -1.00 DPS) [quest]; Horizon Choker (13085, -1.16 DPS, sim-verified) [world_drop]; Triune Amulet (7722, -2.10 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 31.1 spell_power points (12.61 DPS) | yes | Rotgrip Mantle (17732, -1.69 DPS, sim-verified) [dungeon]; Rockshard Pauldrons (9411, -3.49 DPS) [dungeon]; Skycaller's Leather Shoulders (252537, -4.38 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Runecloth Cloak (13860, -1.36 DPS, sim-verified) [crafted]; Big Voodoo Cloak (8216, -2.74 DPS) [crafted]; Expeditionary's Cape (271921, -3.11 DPS) [rep] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 34.9 spell_power points (14.16 DPS) | yes | Feathered Breastplate (8349, -2.44 DPS, sim-verified) [crafted]; Robe of the Magi (1716, -3.25 DPS) [world_drop]; Runecloth Tunic (13857, -3.71 DPS) [crafted] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 17.9 spell_power points (7.28 DPS) | yes | Skycaller's Leather Bracers (252542, -0.92 DPS) [crafted]; Skycaller's Mail Bracers (252571, -0.92 DPS) [crafted]; Green Whelp Bracers (7386, -1.50 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller's Mail Gauntlets (252585, -3.84 DPS) [crafted]; Elven Spirit Claws (2564, -4.12 DPS) [world_drop]; Skycaller's Leather Gauntlets (252550, -4.38 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 25.5 spell_power points (10.36 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Satyrmane Sash (17755, -1.46 DPS) [dungeon]; Dawnspire Cord (12466, -1.80 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.5 spell_power points (9.55 DPS) | yes | Big Voodoo Pants (8202, -0.24 DPS) [crafted]; Wizardweave Leggings (14132, -1.84 DPS) [crafted]; Turtle Scale Leggings (8185, -1.86 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.2 spell_power points (9.81 DPS) | yes | Skycaller's Leather Boots (252471, -0.17 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -0.17 DPS) [crafted]; Greaves of Withering Despair (22240, -0.43 DPS) [dungeon] |
| finger1 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 17.8 spell_power points (7.23 DPS) | yes | Band of the Unicorn (7553, -1.96 DPS) [world_drop]; Ring of Power Regulation (273030, -3.03 DPS) [dungeon]; Snake Hoop (6750, -3.19 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.0 spell_power points (5.67 DPS) | yes | Band of the Unicorn (7553, -0.40 DPS) [world_drop]; Ring of Power Regulation (273030, -1.46 DPS) [dungeon]; Snake Hoop (6750, -1.62 DPS) [quest] |
| trinket1 | Enriched Seal (271903) | Theramore Expeditionary Force [rep] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -1.44 DPS) [crafted]; Ankh of Life (1713, -2.84 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -2.86 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -0.92 DPS) [crafted]; Thunderbrew's Boot Flask (744, -1.11 DPS, sim-verified) [quest]; Ankh of Life (1713, -2.31 DPS) [world_drop] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -3.79 DPS) [dungeon]; Glowing Brightwood Staff (812, -4.08 DPS, sim-verified) [world_drop]; Spellshifter Rod (9527, -4.10 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; waist: Skycaller's Leather Waistguard; finger1: Lorekeeper's Ring; trinket1: Enriched Seal; trinket2: Uther's Strength

No-known-source sample (15 of 848, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60 (dwarf, 5510331300001351-050000000000000000-0530520000000000)

Set DPS (verified): 295.5. Weights run: 4.7s. Verify run: 45.4s. 1668 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.571 ± 0.011, crit=0.383 ± 0.008 per rating point (14 rating = 1%, 5.360 per %), hit=0.694 ± 0.036 per rating point (10 rating = 1%, 6.939 per %), spell_haste=3.040 ± 0.353, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.880 ± 0.001, spirit=0.394 ± 0.089, mp5=1.020 ± 0.157

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blue Dragonscale Helm (252604) | Leatherworking [crafted] | 42.2 spell_power points (19.10 DPS) | yes | Living Crown (252561, -0.64 DPS) [crafted]; Soothsayer's Headdress (17740, -3.39 DPS) [dungeon]; Dreadmist Mask (16698, -4.68 DPS) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 17.7 spell_power points (8.01 DPS) | yes | Mindburst Medallion (11196, -2.86 DPS) [quest]; Glowing Eye of Mordresh (10769, -3.06 DPS, sim-verified) [dungeon]; Triune Amulet (7722, -3.15 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-verified (295.5 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Argent Shoulders (19059, -0.82 DPS) [crafted]; Theramore Shoulderguards (272957, -8.13 DPS, sim-verified) [rep] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (295.5 DPS) | yes | Cloak of Earth and Sky (279270, -0.49 DPS) [crafted]; Stormpike Sage's Cloak (19086, -2.62 DPS) [rep]; Runecloth Cloak (13860, -2.81 DPS) [crafted] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | 38.9 spell_power points (17.59 DPS) | yes | Feathered Breastplate (8349, -2.13 DPS) [crafted]; Magister's Robes (16688, -3.50 DPS) [dungeon]; Vest of Elements (16666, -10.60 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 29.3 spell_power points (13.27 DPS) | yes | Runecloth Cuffs (254123, -5.71 DPS) [crafted]; Spitfire Bracers (20481, -5.72 DPS) [crafted]; Skycaller's Leather Bracers (252542, -6.04 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-verified (295.5 DPS) | yes | Theramore Handguards (271950, -2.57 DPS, sim-verified) [rep]; Storm Gauntlets (12632, -3.56 DPS) [crafted]; Elven Spirit Claws (2564, -3.58 DPS) [world_drop] |
| waist | Stormpike Cloth Girdle (19094) | Stormpike Guard [rep] | 23.7 spell_power points (10.73 DPS) | yes | Skycaller's Leather Waistguard (252476, -0.39 DPS) [crafted]; Skycaller's Mail Belt (252589, -0.39 DPS) [crafted]; Highlander's Cloth Girdle (20047, -0.42 DPS) [rep] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 53.6 spell_power points (24.26 DPS) | yes | Sentinel's Silk Leggings (237815, -2.34 DPS, sim-verified) [vendor]; Sentinel's Lizardhide Pants (237817, -2.36 DPS) [vendor] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 26.9 spell_power points (12.15 DPS) | yes | Earthen Silk Slippers (254013, -0.04 DPS) [crafted]; Venomspew Footpads (275606, -0.77 DPS) [crafted]; Bileblister Boots (275612, -0.77 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 24.4 spell_power points (11.02 DPS) | yes | Deep Rooted Ring (19109, -3.01 DPS) [world]; Band of the Unicorn (7553, -5.14 DPS) [world_drop]; Philanthropist's Ring (281635, -5.20 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19522) | Silverwing Sentinels [rep] | 18.1 spell_power points (8.18 DPS) | yes | Deep Rooted Ring (19109, -0.17 DPS) [world]; Band of the Unicorn (7553, -2.30 DPS) [world_drop]; Philanthropist's Ring (281635, -2.36 DPS) [quest] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (295.5 DPS) | yes | Uther's Strength (11302, -0.05 DPS) [world_drop]; Enriched Seal (271903, -0.46 DPS) [rep]; Ankh of Life (1713, -0.63 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.11 DPS) [world_drop]; Enriched Seal (271903, -0.52 DPS) [rep]; Ankh of Life (1713, -0.69 DPS) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (295.5 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Sageclaw (20070, -3.62 DPS) [rep]; Ironbark Staff (20069, -10.13 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | sim-verified (295.5 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Barricade (235473, +0.00 DPS) [vendor]; Spirit of Aquementas (11904, -3.70 DPS, sim-verified) [quest] |
| ranged | Totem of Thunder (228176) | Enchanting [crafted] | sim-verified (295.5 DPS) | yes | Totem of the Storm (272432, -1.83 DPS, sim-verified) [world_drop] |

**New at 60:** head: Blue Dragonscale Helm; neck: Archlight Talisman; chest: Ironfeather Breastplate; wrist: Dryad's Wrist Bindings; waist: Stormpike Cloth Girdle; legs: Ironfeather Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Lorekeeper's Ring; trinket1: Shard of the Splithooves; trinket2: Frozen Heart of the Mountain; main_hand: Hand of Edward the Odd; off_hand: Therazane's Touch; ranged: Totem of Thunder

No-known-source sample (15 of 1668, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60, raid preset (dwarf, 5510331300001351-055000000000000000-0530020000000000)

Set DPS (verified): 663.8. Weights run: 4.7s. Verify run: 48.9s. 1668 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.285 ± 0.007, crit=0.616 ± 0.014 per rating point (14 rating = 1%, 8.623 per %), hit=1.186 ± 0.073 per rating point (10 rating = 1%, 11.860 per %), spell_haste=8.756 ± 0.609, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.877 ± 0.001, spirit=not significant (0.049 ± 0.019), mp5=0.154 ± 0.025

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blue Dragonscale Helm (252604) | Leatherworking [crafted] | 42.0 spell_power points (23.45 DPS) | yes | Mask of the Unforgiven (13404, -5.11 DPS, sim-verified) [dungeon]; Living Crown (252561, -6.24 DPS) [crafted]; Soothsayer's Headdress (17740, -8.00 DPS) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 14.8 spell_power points (8.29 DPS) | yes | Glowing Eye of Mordresh (10769, -3.50 DPS, sim-verified) [dungeon]; Triune Amulet (7722, -3.63 DPS) [dungeon]; Mindburst Medallion (11196, -3.85 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-verified (663.8 DPS) | yes | Argent Shoulders (19059, +0.00 DPS) [crafted]; Theramore Shoulderguards (272957, +0.00 DPS) [rep]; Mantle of the Timbermaw (19050, -8.64 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (663.8 DPS) | yes | Stormpike Sage's Cloak (19086, -1.12 DPS) [rep]; Cloak of Earth and Sky (279270, -1.20 DPS) [crafted]; Cloak of Flames (3475, -2.31 DPS) [world_drop] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | sim-verified (663.8 DPS) | yes | Fernpulse Jerkin (15786, +0.00 DPS) [quest]; Acumen Robes (17775, -0.61 DPS) [quest]; Vest of Elements (16666, -13.25 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 24.6 spell_power points (13.75 DPS) | yes | Runecloth Cuffs (254123, -6.02 DPS) [crafted]; Runic Leather Bracers (15092, -6.58 DPS) [crafted]; Skycaller's Leather Bracers (252542, -6.92 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-verified (663.8 DPS) | yes | Storm Gauntlets (12632, -1.63 DPS) [crafted]; Earth Warder's Gloves (21318, -6.02 DPS) [quest]; Theramore Handguards (271950, -6.03 DPS) [rep] |
| waist | Highlander's Cloth Girdle (20047) | The League of Arathor [rep] | 24.3 spell_power points (13.59 DPS) | yes | Stormpike Cloth Girdle (19094, -1.95 DPS) [rep]; Skycaller's Leather Waistguard (252476, -2.75 DPS) [crafted]; Skycaller's Mail Belt (252589, -2.75 DPS) [crafted] |
| legs | Sentinel's Silk Leggings (237815) (or Sentinel's Lizardhide Pants (237817)) | Illiyana Moonblaze [vendor] | 52.1 spell_power points (29.10 DPS) | yes | Sentinel's Lizardhide Pants (237817, +0.00 DPS) [vendor]; Ironfeather Leggings (252486, -0.19 DPS) [crafted]; Pristine Scorpid Leggings (252606, -5.85 DPS) [crafted] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 25.7 spell_power points (14.36 DPS) | yes | Omnicast Boots (11822, -1.28 DPS) [dungeon]; Venomspew Footpads (275606, -1.85 DPS) [crafted]; Earthen Silk Slippers (254013, -4.60 DPS, sim-verified) [crafted] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 20.5 spell_power points (11.44 DPS) | yes | Deep Rooted Ring (19109, -3.15 DPS) [world]; Lorekeeper's Ring (19522, -3.28 DPS) [rep]; Band of the Unicorn (7553, -4.18 DPS) [world_drop] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | 20.3 spell_power points (11.33 DPS) | yes | Lorekeeper's Ring (19522, -3.17 DPS) [rep]; Deep Rooted Ring (19109, -4.00 DPS, sim-verified) [world]; Band of the Unicorn (7553, -4.07 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (663.8 DPS) | yes | Shard of the Splithooves (10659, -5.45 DPS) [quest]; Enriched Seal (271903, -5.53 DPS) [rep]; Uther's Strength (11302, -6.23 DPS, sim-verified) [world_drop] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Enriched Seal (271903, +0.00 DPS) [rep] |
| main_hand | Sageclaw (20070) | The League of Arathor [rep] | sim-verified (663.8 DPS) | yes | Ironbark Staff (20069, +0.00 DPS) [rep]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Hand of Edward the Odd (2243, -1.36 DPS) [world_drop] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | sim-verified (663.8 DPS) | yes | Grand Marshal's Tome of Power (234589, -0.17 DPS) [pvp]; Grand Marshal's Barricade (235473, -1.20 DPS) [vendor]; Spirit of Aquementas (11904, -4.62 DPS, sim-verified) [quest] |
| ranged | Totem of Thunder (228176) | Enchanting [crafted] | sim-verified (663.8 DPS) | yes | Totem of the Storm (272432, -3.71 DPS, sim-verified) [world_drop] |

**New at 60:** head: Blue Dragonscale Helm; neck: Archlight Talisman; chest: Ironfeather Breastplate; wrist: Dryad's Wrist Bindings; waist: Highlander's Cloth Girdle; legs: Sentinel's Silk Leggings; feet: Greaves of Withering Despair; finger1: Don Julio's Band; finger2: Maiden's Circle; trinket1: Frozen Heart of the Mountain; trinket2: Thunderbrew's Boot Flask; main_hand: Sageclaw; off_hand: Therazane's Touch; ranged: Totem of Thunder

No-known-source sample (15 of 1668, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

## Horde

### Band 20 (orc, 5510000000000000-000000000000000000-0000000000000000)

Set DPS (verified): 28.5. Weights run: 6.8s. Verify run: 9.9s. 239 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.785 ± 0.011, crit=0.061 ± 0.002 per rating point (14 rating = 1%, 0.852 per %), hit=0.211 ± 0.007 per rating point (10 rating = 1%, 2.114 per %), spell_haste=0.157 ± 0.029, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.885 ± 0.002, spirit=not significant (0.185 ± 0.115), mp5=2.116 ± 0.140

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Hood (252448) | Leatherworking [crafted] | 9.0 spell_power points (0.78 DPS) | yes | Wisdom's Leather Hood (252507, -0.26 DPS) [crafted]; Pristine Circlet (253949, -0.26 DPS) [crafted]; Trapper's Leather Hood (252505, -0.27 DPS, sim-verified) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.6 spell_power points (0.05 DPS) | yes | Scholarly Pendant (277203, -0.02 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.1 spell_power points (1.05 DPS) | yes | Slime-encrusted Pads (6461, -0.50 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.70 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.71 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pearl-clasped Cloak (5542, -0.03 DPS) [crafted]; Battle Healer's Cloak (20427, -0.12 DPS) [rep]; Sanguine Cape (14376, -0.14 DPS) [world_drop] |
| chest | Stormrider's Leather Armor (252492) | Leatherworking [crafted] | 12.5 spell_power points (1.09 DPS) | yes | Filigreed Pristine Gown (253901, -0.26 DPS) [crafted]; Totemic Leather Armor (252435, -0.36 DPS) [crafted]; Wisdom's Leather Armor (252493, -0.37 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 10.3 spell_power points (0.90 DPS) | yes | Owl Bracers (4796, -0.58 DPS, sim-verified) [vendor]; Bright Bracers (3647, -0.62 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.63 DPS) [quest] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 8.1 spell_power points (0.71 DPS) | yes | Pristine Gloves (253913, -0.16 DPS) [crafted]; Wisdom's Leather Gloves (252499, -0.17 DPS) [crafted]; Serpent Gloves (5970, -0.54 DPS, sim-verified) [dungeon] |
| waist | Pristine Sash (253925) (or Stormrider's Leather Belt (252432)) | Tailoring [crafted] | 7.1 spell_power points (0.62 DPS) | yes | Stormrider's Leather Belt (252432, +0.00 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.07 DPS) [crafted]; Keller's Girdle (2911, -0.08 DPS) [world_drop] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 15.4 spell_power points (1.35 DPS) | yes | Filigreed Pristine Leggings (253937, -0.35 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.37 DPS, sim-verified) [crafted]; Totemic Leather Pants (252446, -0.46 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.1 spell_power points (0.88 DPS) | yes | Wisdom's Leather Boots (252444, -0.19 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.26 DPS, sim-verified) [crafted]; Black Whelp Slippers (252424, -0.35 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 9.2 spell_power points (0.81 DPS) | yes | Loop of Sacrifice (281673, -0.46 DPS) [quest]; Volcanic Rock Ring (12053, -0.60 DPS) [world_drop]; Band of Purification (12996, -0.71 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 4.7 spell_power points (0.41 DPS) | yes | Volcanic Rock Ring (12053, -0.21 DPS) [world_drop]; Band of Purification (12996, -0.31 DPS) [world_drop]; Loop of Sacrifice (281673, -0.40 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-verified (28.5 DPS) | yes | Staff of the Blessed Seer (2271, -0.06 DPS) [dungeon]; Gnarled Necromancer's Staff (251534, -0.23 DPS) [quest]; Crescent Staff (6505, -0.74 DPS, sim-verified) [quest] |
| off_hand | Seedcloud Buckler (6630) | Wailing Caverns: Verdan the Everliving [dungeon] | 8.9 spell_power points (0.78 DPS) | yes | Worn Turtle Shell Shield (6447, -0.41 DPS) [dungeon]; Owl's Disk (4822, -0.50 DPS) [vendor]; Bristlebark Buckler (15894, -0.50 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 20:** head: Totemic Leather Hood; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Stormrider's Leather Armor; wrist: Mindthrust Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Stormrider's Leather Pants; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Evocator's Blade; off_hand: Seedcloud Buckler

No-known-source sample (15 of 239, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 20425 Advisor's Gnarled Staff; 20441 Scout's Blade; 23173 Abomination Skin Leggings; 206376 Bulwark Icon; 206383 Lithic Icon; 206384 Astral Icon; 206385 Clastic Icon

### Band 30 (orc, 5510331300000000-000000000000000000-0000000000000000)

Set DPS (verified): 56.0. Weights run: 6.5s. Verify run: 11.3s. 428 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.755 ± 0.012, crit=0.099 ± 0.003 per rating point (14 rating = 1%, 1.386 per %), hit=0.352 ± 0.012 per rating point (10 rating = 1%, 3.523 per %), spell_haste=not significant (-0.482 ± 0.227), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.850 ± 0.002, spirit=not significant (0.124 ± 0.081), mp5=2.198 ± 0.216

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 13.6 spell_power points (2.86 DPS) | yes | Enduring Cap (3020, -0.13 DPS) [world_drop]; Totemic Leather Helm (252456, -0.33 DPS) [crafted]; Holy Shroud (2721, -0.38 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 4.5 spell_power points (0.96 DPS) | yes | Darkspear Warding Pendant (272075, -0.16 DPS) [vendor]; Kaleidoscope Chain (13084, -0.21 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.24 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.8 spell_power points (3.33 DPS) | yes | Death Speaker Mantle (6685, -0.54 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.84 DPS) [world_drop]; Invoker's Mantle (215365, -1.06 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, -0.30 DPS) [crafted]; Resilient Cape (14400, -0.40 DPS) [world_drop]; Windsong Drape (15468, -0.59 DPS, sim-verified) [quest] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 19.1 spell_power points (4.02 DPS) | yes | Mechbuilder's Overalls (9508, -0.42 DPS) [dungeon]; Green Silk Armor (7065, -0.48 DPS, sim-verified) [crafted]; Guardian Armor (4256, -0.53 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 15.6 spell_power points (3.29 DPS) | yes | Mindthrust Bracers (1974, -1.07 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -2.01 DPS) [world_drop]; Nightsky Wristbands (6407, -2.25 DPS) [world_drop] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | 10.5 spell_power points (2.22 DPS) | yes | Jutebraid Gloves (10654, -0.15 DPS) [quest]; Stormrider's Leather Gloves (252498, -0.53 DPS) [crafted]; Truefaith Gloves (7049, -0.68 DPS) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 15.0 spell_power points (3.17 DPS) | yes | Defiler's Cloth Girdle (20164, -0.53 DPS, sim-verified) [rep]; Moss Cinch (6911, -0.64 DPS) [dungeon]; Belt of Arugal (6392, -0.66 DPS) [dungeon] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Guardian Pants (5962, -0.21 DPS) [crafted]; Dark Ritual Leggings (270031, -0.22 DPS) [quest]; Leech Pants (6910, -0.26 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 12.8 spell_power points (2.70 DPS) | yes | Acidic Walkers (9454, -0.26 DPS) [dungeon]; Spidersilk Boots (4320, -0.58 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.63 DPS) [crafted] |
| finger1 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 12.6 spell_power points (2.66 DPS) | yes | Snake Hoop (6750, -0.59 DPS) [quest]; Darkspear Signet (272071, -0.80 DPS) [vendor]; Solid Gold Band (279845, -1.68 DPS) [quest] |
| finger2 | Advisor's Ring (19521) | Warsong Outriders [rep] | 11.4 spell_power points (2.40 DPS) | yes | Snake Hoop (6750, -0.34 DPS) [quest]; Darkspear Signet (272071, -0.55 DPS) [vendor]; Solid Gold Band (279845, -1.43 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -0.95 DPS) [dungeon]; Tendonscraper Dagger (273053, -0.95 DPS) [world]; Glimmering Staff (249392, -1.88 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 11.5 spell_power points (2.43 DPS) | yes | Seedcloud Buckler (6630, -0.56 DPS, sim-verified) [dungeon]; Witch's Finger (16887, -1.16 DPS) [quest]; Heart of Agamaggan (6694, -1.38 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Darkspear Raider's Cloak; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; feet: Gilded Slippers; finger1: Ring of Power Regulation; finger2: Advisor's Ring; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight

No-known-source sample (15 of 428, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape

### Band 40 (orc, 5510331300001351-000000000000000000-0000000000000000)

Set DPS (verified): 115.7. Weights run: 6.1s. Verify run: 12.5s. 644 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.663 ± 0.009, crit=0.093 ± 0.002 per rating point (14 rating = 1%, 1.295 per %), hit=0.174 ± 0.007 per rating point (10 rating = 1%, 1.741 per %), spell_haste=-0.898 ± 0.089, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.861 ± 0.002, spirit=not significant (0.043 ± 0.052), mp5=not significant (0.236 ± 0.086)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | sim-verified (115.7 DPS) | yes | Big Voodoo Mask (8201, -0.29 DPS) [crafted]; Augural Shroud (2620, -0.43 DPS) [world]; Spellpower Goggles Xtreme (10502, -1.44 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.3 spell_power points (2.63 DPS) | yes | Necklace of Calisea (1714, -1.57 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.64 DPS) [vendor]; Triune Amulet (7722, -2.06 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.6 spell_power points (3.33 DPS) | yes | Green Silken Shoulders (7057, -0.07 DPS) [crafted]; Bloodmage Mantle (7684, -0.14 DPS) [dungeon]; Berylline Pads (4197, -0.42 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.40 DPS) [vendor]; Icy Cloak (4327, -0.44 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.2 spell_power points (5.59 DPS) | yes | Dreamweave Vest (10021, -0.35 DPS) [crafted]; Robe of Power (7054, -0.83 DPS) [crafted]; Elemental Raiment (9434, -1.11 DPS) [world_drop] |
| wrist | Green Whelp Bracers (7386) | Leatherworking [crafted] | 13.0 spell_power points (2.79 DPS) | yes | Spidertank Oilrag (9448, -0.71 DPS) [dungeon]; Turtle Scale Bracers (8198, -0.80 DPS) [crafted]; Guardian Leather Bracers (4260, -1.52 DPS, sim-verified) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Red Mageweave Gloves (10018, -3.20 DPS) [crafted]; Skycaller's Leather Gloves (252529, -3.34 DPS) [crafted]; Dreamweave Gloves (10019, -3.34 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 16.9 spell_power points (3.62 DPS) | yes | Skycaller's Leather Belt (252522, -0.60 DPS) [crafted]; Gilded Cord (254037, -0.72 DPS) [crafted]; Defiler's Cloth Girdle (20166, -1.90 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.0 spell_power points (4.69 DPS) | yes | Stormrider's Leather Pants (252502, -1.67 DPS) [crafted]; Dark Ritual Leggings (270031, -1.70 DPS) [quest]; Crimson Silk Pantaloons (7062, -1.78 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.3 spell_power points (5.19 DPS) | yes | Skycaller's Mail Boots (252563, -1.38 DPS) [crafted]; Skycaller's Leather Shoes (252532, -1.62 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -2.42 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.3 spell_power points (2.84 DPS) | yes | Snake Hoop (6750, -0.85 DPS) [quest]; Reedknot Ring (9622, -1.35 DPS) [quest]; Ring of Power Regulation (273030, -1.41 DPS) [dungeon] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.7 spell_power points (2.07 DPS) | yes | Reedknot Ring (9622, -0.58 DPS) [quest]; Ring of Power Regulation (273030, -0.64 DPS) [dungeon]; Snake Hoop (6750, -1.37 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hypnotic Blade (7714, -5.73 DPS) [dungeon]; Glimmering Staff (249392, -7.47 DPS) [crafted]; Illusionary Rod (7713, -10.87 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Green Whelp Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Staff of Jordan

No-known-source sample (15 of 644, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7748 Forcestone Buckler

### Band 50 (orc, 5510331300001351-050000000000000000-0500000000000000)

Set DPS (verified): 176.8. Weights run: 5.6s. Verify run: 14.6s. 819 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.414 ± 0.007, crit=0.101 ± 0.002 per rating point (14 rating = 1%, 1.420 per %), hit=0.173 ± 0.006 per rating point (10 rating = 1%, 1.729 per %), spell_haste=-0.353 ± 0.074, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.889 ± 0.002, spirit=0.233 ± 0.026, mp5=1.189 ± 0.065

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 31.1 spell_power points (8.79 DPS) | yes | Red Mageweave Headband (10033, -1.07 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -1.15 DPS) [crafted]; Dreamweave Circlet (10041, -2.37 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 10.3 spell_power points (2.92 DPS) | yes | Mindburst Medallion (11196, -0.19 DPS) [quest]; Triune Amulet (7722, -0.19 DPS) [dungeon]; Horizon Choker (13085, -1.01 DPS) [world_drop] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 25.1 spell_power points (7.11 DPS) | yes | Rockshard Pauldrons (9411, -1.80 DPS) [dungeon]; Rotgrip Mantle (17732, -1.88 DPS, sim-verified) [dungeon]; Skycaller's Leather Shoulders (252537, -1.99 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -0.61 DPS) [quest]; Runecloth Cloak (13860, -1.58 DPS) [crafted]; Nightfall Drape (12465, -2.51 DPS) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | sim-verified (+3.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Robe of the Magi (1716, -0.46 DPS) [world_drop]; Flamestrider Robes (11747, -0.54 DPS) [dungeon]; Feathered Breastplate (8349, -3.00 DPS, sim-verified) [crafted] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 14.1 spell_power points (4.00 DPS) | yes | Skycaller's Leather Bracers (252542, -0.02 DPS) [crafted]; Skycaller's Mail Bracers (252571, -0.02 DPS) [crafted]; Spidertank Oilrag (9448, -0.45 DPS) [dungeon] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamweave Gloves (10019, -3.28 DPS) [crafted]; Skycaller's Leather Gauntlets (252550, -3.33 DPS) [crafted]; Elven Spirit Claws (2564, -4.47 DPS, sim-verified) [world_drop] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 21.0 spell_power points (5.93 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Satyrmane Sash (17755, -0.80 DPS) [dungeon]; Ghostweave Cord (254073, -1.25 DPS) [crafted] |
| legs | Stone Guard's Pulsing Legplates (220847) | Lady Palanseer [vendor] | 24.7 spell_power points (6.99 DPS) | yes | Senior Designer's Pantaloons (11841, -1.59 DPS) [dungeon]; Wizardweave Leggings (14132, -1.61 DPS) [crafted]; Big Voodoo Pants (8202, -3.32 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.6 spell_power points (7.25 DPS) | yes | Stormcloth Boots (10039, -0.96 DPS) [crafted]; Skycaller's Leather Boots (252471, -1.72 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -1.72 DPS) [crafted] |
| finger1 | Advisor's Ring (19519) | Warsong Outriders [rep] | 16.8 spell_power points (4.74 DPS) | yes | Ring of Power Regulation (273030, -2.03 DPS) [dungeon]; Darkspear Signet (272069, -2.39 DPS) [vendor]; Band of the Unicorn (7553, -8.48 DPS, sim-verified) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Band of the Unicorn (7553, -0.52 DPS, sim-verified) [world_drop]; Ring of Power Regulation (273030, -0.71 DPS) [dungeon]; Darkspear Signet (272069, -1.06 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, -0.89 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.02 DPS) [quest]; Frozen Heart of the Mountain (249469, -1.24 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, -0.91 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -1.00 DPS, sim-verified) [crafted]; Evonice's Landin' Pilla (18951, -1.04 DPS) [quest] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spire of Hakkar (10844, -2.12 DPS) [world]; Glowing Brightwood Staff (812, -2.57 DPS) [world_drop]; Kindling Stave (11750, -4.69 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; waist: Skycaller's Leather Waistguard; legs: Stone Guard's Pulsing Legplates; finger1: Advisor's Ring; finger2: Philanthropist's Ring; trinket1: Darkspear Voodoo Seal; trinket2: Uther's Strength

No-known-source sample (15 of 819, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders

### Band 60 (orc, 5510331300001351-050000000000000000-0530520000000000)

Set DPS (verified): 296.0. Weights run: 4.7s. Verify run: 45.8s. 1628 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.571 ± 0.011, crit=0.383 ± 0.008 per rating point (14 rating = 1%, 5.360 per %), hit=0.694 ± 0.036 per rating point (10 rating = 1%, 6.939 per %), spell_haste=3.040 ± 0.353, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.880 ± 0.001, spirit=0.394 ± 0.089, mp5=1.020 ± 0.157

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blue Dragonscale Helm (252604) | Leatherworking [crafted] | 42.2 spell_power points (19.10 DPS) | yes | Warlord's Mail Helm (231663, -1.10 DPS) [pvp]; Soothsayer's Headdress (17740, -3.39 DPS) [dungeon]; Living Crown (252561, -3.41 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 17.7 spell_power points (8.01 DPS) | yes | Glowing Eye of Mordresh (10769, -2.75 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -2.86 DPS) [quest]; Triune Amulet (7722, -3.15 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-verified (296.0 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Warlord's Mail Spaulders (231659, +0.00 DPS) [pvp]; Darkspear Shoulderguards (272958, -8.24 DPS, sim-verified) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of Earth and Sky (279270, -0.49 DPS) [crafted]; Deep Woodlands Cloak (19121, -1.20 DPS) [quest]; Frostwolf Advisor's Cloak (19085, -2.62 DPS) [rep] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | 38.9 spell_power points (17.59 DPS) | yes | Warlord's Mail Breastplate (231662, +0.00 DPS) [vendor]; Legionnaire's Mail Breastplate (227165, -1.47 DPS) [vendor]; Vest of Elements (16666, -10.61 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 29.3 spell_power points (13.27 DPS) | yes | Runecloth Cuffs (254123, -5.71 DPS) [crafted]; Spitfire Bracers (20481, -5.72 DPS) [crafted]; Skycaller's Leather Bracers (252542, -6.04 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-verified (296.0 DPS) | yes | General's Mail Gauntlets (231660, +0.00 DPS) [pvp]; General's Mail Gloves (231666, -2.20 DPS) [vendor]; Raider Handguards (272101, -2.65 DPS, sim-verified) [vendor] |
| waist | Frostwolf Cloth Belt (19090) | Frostwolf Clan [rep] | 23.7 spell_power points (10.73 DPS) | yes | Skycaller's Leather Waistguard (252476, -0.39 DPS) [crafted]; Skycaller's Mail Belt (252589, -0.39 DPS) [crafted]; Defiler's Cloth Girdle (20163, -0.42 DPS) [rep] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 53.6 spell_power points (24.26 DPS) | yes | General's Mail Leggings (231664, -2.97 DPS) [pvp]; Outrider's Silk Leggings (22747, -5.38 DPS, sim-verified) [rep]; Pristine Scorpid Leggings (252606, -6.91 DPS) [crafted] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 26.9 spell_power points (12.15 DPS) | yes | General's Mail Sabatons (231661, +0.00 DPS) [vendor]; Earthen Silk Slippers (254013, -0.04 DPS) [crafted]; Venomspew Footpads (275606, -0.77 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 24.4 spell_power points (11.02 DPS) | yes | Deep Rooted Ring (19109, -3.01 DPS) [world]; Band of the Unicorn (7553, -5.14 DPS) [world_drop]; Philanthropist's Ring (281635, -5.20 DPS) [quest] |
| finger2 | Advisor's Ring (19518) | Warsong Outriders [rep] | 18.1 spell_power points (8.18 DPS) | yes | Deep Rooted Ring (19109, -0.17 DPS) [world]; Band of the Unicorn (7553, -2.30 DPS) [world_drop]; Philanthropist's Ring (281635, -2.36 DPS) [quest] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (296.0 DPS) | yes | Uther's Strength (11302, -0.05 DPS) [world_drop]; Darkspear Voodoo Seal (272061, -0.46 DPS) [vendor]; Rune of the Guard Captain (19120, -0.57 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (296.0 DPS) | yes | Uther's Strength (11302, -0.11 DPS) [world_drop]; Rune of the Guard Captain (19120, -0.63 DPS) [quest]; Darkspear Voodoo Seal (272061, -2.32 DPS, sim-verified) [vendor] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (296.0 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Mindfang (20214, -3.62 DPS) [rep]; Ironbark Staff (20220, -10.36 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | sim-verified (296.0 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Barricade (235474, +0.00 DPS) [vendor]; Spirit of Aquementas (11904, -3.73 DPS, sim-verified) [quest] |
| ranged | Totem of Thunder (228176) | Enchanting [crafted] | sim-verified (296.0 DPS) | yes | Totem of the Storm (272432, -1.63 DPS, sim-verified) [world_drop] |

**New at 60:** head: Blue Dragonscale Helm; neck: Archlight Talisman; chest: Ironfeather Breastplate; wrist: Dryad's Wrist Bindings; waist: Frostwolf Cloth Belt; legs: Ironfeather Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Advisor's Ring; trinket1: Shard of the Splithooves; trinket2: Frozen Heart of the Mountain; main_hand: Hand of Edward the Odd; off_hand: Therazane's Touch; ranged: Totem of Thunder

No-known-source sample (15 of 1628, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders

### Band 60, raid preset (orc, 5510331300001351-055000000000000000-0530020000000000)

Set DPS (verified): 660.2. Weights run: 4.7s. Verify run: 46.5s. 1628 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.285 ± 0.007, crit=0.616 ± 0.014 per rating point (14 rating = 1%, 8.623 per %), hit=1.186 ± 0.073 per rating point (10 rating = 1%, 11.860 per %), spell_haste=8.756 ± 0.609, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.877 ± 0.001, spirit=not significant (0.049 ± 0.019), mp5=0.154 ± 0.025

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blue Dragonscale Helm (252604) | Leatherworking [crafted] | 42.0 spell_power points (23.45 DPS) | yes | Warlord's Mail Helm (231663, -2.61 DPS) [pvp]; Mask of the Unforgiven (13404, -4.27 DPS, sim-verified) [dungeon]; Champion's Mail Helm (227162, -4.92 DPS) [pvp] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 14.8 spell_power points (8.29 DPS) | yes | Glowing Eye of Mordresh (10769, -3.42 DPS, sim-verified) [dungeon]; Triune Amulet (7722, -3.63 DPS) [dungeon]; Mindburst Medallion (11196, -3.85 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 29.3 spell_power points (16.38 DPS) | yes | Champion's Mail Spaulders (227160, +0.00 DPS) [vendor]; Warlord's Mail Spaulders (231659, +0.00 DPS) [pvp]; Argent Shoulders (19059, -3.35 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (660.2 DPS) | yes | Cloak of Flames (3475, +0.00 DPS) [world_drop]; Deep Woodlands Cloak (19121, -0.81 DPS) [quest]; Frostwolf Advisor's Cloak (19085, -1.12 DPS) [rep] |
| chest | Stone Guard's Pulsing Breastplate (220844) | Lady Palanseer [vendor] | 33.3 spell_power points (18.59 DPS) | yes | Legionnaire's Mail Breastplate (227165, +0.00 DPS) [vendor]; Warlord's Mail Breastplate (231662, +0.00 DPS) [vendor]; Vest of Elements (16666, -3.01 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 24.6 spell_power points (13.75 DPS) | yes | Runecloth Cuffs (254123, -6.02 DPS) [crafted]; Runic Leather Bracers (15092, -6.58 DPS) [crafted]; Skycaller's Leather Bracers (252542, -6.92 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Mail Gauntlets (231660, +0.00 DPS) [pvp]; Storm Gauntlets (12632, -1.56 DPS, sim-verified) [crafted]; General's Mail Gloves (231666, -3.29 DPS) [vendor] |
| waist | Defiler's Cloth Girdle (20163) | The Defilers [rep] | 24.3 spell_power points (13.59 DPS) | yes | Frostwolf Cloth Belt (19090, -1.95 DPS) [rep]; Skycaller's Leather Waistguard (252476, -2.75 DPS) [crafted]; Skycaller's Mail Belt (252589, -2.75 DPS) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 51.7 spell_power points (28.90 DPS) | yes | General's Mail Leggings (231664, -4.63 DPS) [pvp]; Pristine Scorpid Leggings (252606, -6.19 DPS, sim-verified) [crafted]; Legionnaire's Mail Leggings (227161, -9.65 DPS) [pvp] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 25.7 spell_power points (14.36 DPS) | yes | General's Mail Sabatons (231661, -0.48 DPS) [vendor]; Earthen Silk Slippers (254013, -0.76 DPS) [crafted]; Omnicast Boots (11822, -1.28 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 20.5 spell_power points (11.44 DPS) | yes | Deep Rooted Ring (19109, -3.15 DPS) [world]; Advisor's Ring (19518, -3.28 DPS) [rep]; Band of the Unicorn (7553, -4.18 DPS) [world_drop] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | 20.3 spell_power points (11.33 DPS) | yes | Advisor's Ring (19518, -3.17 DPS) [rep]; Band of the Unicorn (7553, -4.07 DPS) [world_drop]; Deep Rooted Ring (19109, -4.49 DPS, sim-verified) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (660.2 DPS) | yes | Uther's Strength (11302, -2.61 DPS) [world_drop]; Shard of the Splithooves (10659, -5.45 DPS) [quest]; Darkspear Voodoo Seal (272061, -5.53 DPS) [vendor] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (660.2 DPS) | yes | Uther's Strength (11302, -1.29 DPS) [world_drop]; Shard of the Splithooves (10659, -4.12 DPS) [quest]; Darkspear Voodoo Seal (272061, -4.21 DPS) [vendor] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (660.2 DPS) | yes | Mindfang (20214, +0.00 DPS) [rep]; Ironbark Staff (20220, +0.00 DPS) [rep]; High Warlord's War Staff (234549, +0.00 DPS) [pvp] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | sim-verified (660.2 DPS) | yes | High Warlord's Tome of Destruction (234563, -0.17 DPS) [pvp]; High Warlord's Barricade (235474, -1.20 DPS) [vendor]; Spirit of Aquementas (11904, -4.71 DPS, sim-verified) [quest] |
| ranged | Totem of Thunder (228176) | Enchanting [crafted] | sim-verified (660.2 DPS) | yes | Totem of the Storm (272432, -4.00 DPS, sim-verified) [world_drop] |

**New at 60:** head: Blue Dragonscale Helm; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; chest: Stone Guard's Pulsing Breastplate; wrist: Dryad's Wrist Bindings; waist: Defiler's Cloth Girdle; legs: Ironfeather Leggings; feet: Greaves of Withering Despair; finger1: Don Julio's Band; finger2: Maiden's Circle; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Hand of Edward the Odd; off_hand: Therazane's Touch; ranged: Totem of Thunder

No-known-source sample (15 of 1628, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders

