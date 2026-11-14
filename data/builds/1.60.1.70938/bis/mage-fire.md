# Leveling BiS: Fire

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 000000000000000000-51400001000000000-000000000000000000)

Set DPS (verified): 39.9. Weights run: 2.4s. Verify run: 2.2s. 198 eligible items had no known source.

Race sweep: gnome 40.4 ± 0.1, human 39.9 ± 0.1, high-order-skyborne 39.7 ± 0.1 (dps); gnome beats the guide race human by 0.5 (1.25%) ± 0.1, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.920 ± 0.016, crit=0.147 ± 0.007 per rating point (14 rating = 1%, 2.061 per %), hit=0.305 ± 0.016 per rating point (10 rating = 1%, 3.054 per %), spell_haste=not significant (0.734 ± 0.198), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.316 ± 0.086), mp5=1.108 ± 0.243

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Flame Circlet (253953) | Tailoring [crafted] | 11.0 spell_power points (1.54 DPS) | yes | Pristine Circlet (253949, -0.70 DPS) [crafted]; Shadow Goggles (4373, -1.13 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.36 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.6 spell_power points (0.09 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.3 spell_power points (1.86 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.74 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.30 DPS) [crafted]; Slime-encrusted Pads (6461, -1.39 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pearl-clasped Cloak (5542, -0.07 DPS) [crafted]; Sanguine Cape (14376, -0.22 DPS) [world_drop]; Caretaker's Cape (20428, -0.23 DPS) [rep] |
| chest | Filigreed Flame Gown (253905) | Tailoring [crafted] | 14.5 spell_power points (2.04 DPS) | yes | Filigreed Pristine Gown (253901, -0.57 DPS, sim-verified) [crafted]; Gray Woolen Robe (2585, -0.83 DPS) [crafted]; Scholarly Robes (2034, -1.09 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 7.9 spell_power points (1.11 DPS) | yes | Bright Bracers (3647, -0.64 DPS, sim-verified) [world_drop]; Repurposed Hair Band (281256, -0.76 DPS) [quest]; Mystic's Bracelets (14366, -0.85 DPS) [world_drop] |
| hands | Flame Gloves (253917) | Tailoring [crafted] | 9.8 spell_power points (1.37 DPS) | yes | Phoenix Gloves (4331, -0.28 DPS, sim-verified) [crafted]; Serpent Gloves (5970, -0.39 DPS) [dungeon]; Pristine Gloves (253913, -0.42 DPS) [crafted] |
| waist | Flame Sash (253929) | Tailoring [crafted] | 10.7 spell_power points (1.49 DPS) | yes | Pristine Sash (253925, -0.43 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.46 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.55 DPS) [crafted] |
| legs | Filigreed Flame Leggings (253941) | Tailoring [crafted] | 17.8 spell_power points (2.49 DPS) | yes | Phoenix Pants (4317, -0.61 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.70 DPS) [crafted]; Darkweave Breeches (12987, -1.32 DPS) [world_drop] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.7 spell_power points (1.49 DPS) | yes | Flame Boots (253893, -0.43 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.69 DPS) [crafted]; Red Woolen Boots (4313, -0.71 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 7.2 spell_power points (1.01 DPS) | yes | Lavishly Jeweled Ring (1156, -0.24 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.62 DPS) [world_drop]; Band of Purification (12996, -0.74 DPS) [world_drop] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.8 spell_power points (0.96 DPS) | yes | Lavishly Jeweled Ring (1156, -0.51 DPS, sim-verified) [dungeon]; Volcanic Rock Ring (12053, -0.57 DPS) [world_drop]; Band of Purification (12996, -0.69 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | 31.4 spell_power points (4.40 DPS) | yes | Emberstone Staff (5201, -1.04 DPS) [dungeon]; Staff of the Blessed Seer (2271, -1.13 DPS, sim-verified) [dungeon]; Staff of Westfall (2042, -1.25 DPS) [quest] |
| off_hand | Antipodean Rod (2879) | World drop [world_drop] | 6.0 spell_power points (0.84 DPS) | yes | Tear of Grief (5611, -0.33 DPS) [quest]; Dwarven Tome (279898, -0.38 DPS) [quest]; Mathystra Channeling Orb (281291, -0.51 DPS, sim-verified) [quest] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller (12984, -0.30 DPS) [world_drop]; Deepblaze (279896, -1.51 DPS) [quest]; Sizzle Stick (8071, -1.59 DPS) [quest] |

**New at 20:** head: Flame Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Flame Gown; wrist: Mindthrust Bracers; hands: Flame Gloves; waist: Flame Sash; legs: Filigreed Flame Leggings; feet: Spidersilk Boots; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Evocator's Blade; off_hand: Antipodean Rod; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 198, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 8350 The 1 Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (human, 000000000000000000-54500001230100000-000000000000000000)

Set DPS (verified): 80.0. Weights run: 2.8s. Verify run: 3.9s. 326 eligible items had no known source.

Race sweep: gnome 80.4 ± 0.1, human 80.0 ± 0.1, high-order-skyborne 78.0 ± 0.1 (dps); gnome beats the guide race human by 0.4 (0.49%) ± 0.2, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.857 ± 0.015, crit=0.220 ± 0.012 per rating point (14 rating = 1%, 3.082 per %), hit=0.388 ± 0.021 per rating point (10 rating = 1%, 3.883 per %), spell_haste=0.444 ± 0.108, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.011 ± 0.096), mp5=not significant (0.342 ± 0.169)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Flame Circlet (253979) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Enchanter's Cowl (4322, -0.48 DPS, sim-verified) [crafted]; Holy Shroud (2721, -0.72 DPS) [world_drop]; Flame Circlet (253953, -0.74 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 5.1 spell_power points (1.27 DPS) | yes | Seal of the Expedition (271920, -0.21 DPS) [rep]; Kaleidoscope Chain (13084, -0.41 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.41 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 16.7 spell_power points (4.12 DPS) | yes | Death Speaker Mantle (6685, -0.85 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.99 DPS) [world_drop]; Invoker's Mantle (215365, -1.34 DPS) [crafted] |
| back | Expeditionary's Cape (271923) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, -0.47 DPS) [crafted]; Vine Pruner's Cloak (279835, -0.47 DPS) [quest]; Resilient Cape (14400, -1.06 DPS, sim-verified) [world_drop] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Flame Gown (253965, -0.08 DPS) [crafted]; Mechbuilder's Overalls (9508, -0.27 DPS) [dungeon]; Green Silk Armor (7065, -0.85 DPS, sim-verified) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (3.21 DPS) | yes | Spidertank Oilrag (9448, -0.73 DPS) [dungeon]; Glowing Magical Bracelets (13106, -1.52 DPS) [world_drop]; Mindthrust Bracers (1974, -1.90 DPS) [dungeon] |
| hands | Flame Gloves (253917) | Tailoring [crafted] | 9.6 spell_power points (2.36 DPS) | yes | Phoenix Gloves (4331, -0.14 DPS) [crafted]; Town Clerk's Mittens (270029, -0.14 DPS) [quest]; Truefaith Gloves (7049, -0.49 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 13.6 spell_power points (3.35 DPS) | yes | Belt of Arugal (6392, -0.48 DPS) [dungeon]; Invoker's Cord (215366, -0.56 DPS) [crafted]; Crimson Silk Belt (7055, -0.78 DPS, sim-verified) [crafted] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | 20.1 spell_power points (4.95 DPS) | yes | Filigreed Flame Leggings (253941, -1.33 DPS, sim-verified) [crafted]; Phoenix Pants (4317, -1.64 DPS) [crafted]; Pristine Leggings (253987, -1.73 DPS) [crafted] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 19.0 spell_power points (4.70 DPS) | yes | Gilded Slippers (254001, -1.60 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.76 DPS) [dungeon]; Spidersilk Boots (4320, -2.13 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 10.3 spell_power points (2.54 DPS) | yes | Lorekeeper's Ring (19525, -0.64 DPS) [rep]; Ring of Power Regulation (273030, -0.80 DPS) [dungeon]; Minor Channeling Ring (1449, -0.88 DPS) [quest] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 8.2 spell_power points (2.03 DPS) | yes | Ring of Power Regulation (273030, -0.29 DPS) [dungeon]; Minor Channeling Ring (1449, -0.37 DPS) [quest]; Lorekeeper's Ring (19525, -0.65 DPS, sim-verified) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.00 DPS) [dungeon]; Tendonscraper Dagger (273053, -1.00 DPS) [world]; Glimmering Staff (249392, -4.54 DPS, sim-verified) [crafted] |
| off_hand | Rod of Molten Fire (2565) | World drop [world_drop] | 13.0 spell_power points (3.21 DPS) | yes | Orb of Mystic Insight (249394, -1.01 DPS, sim-verified) [crafted]; Antipodean Rod (2879, -1.73 DPS) [world_drop]; Satyr's Rod (15962, -1.93 DPS) [world_drop] |
| ranged | Scorching Wand (5213) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Starfaller (13063, +0.00 DPS) [world_drop]; Greater Mystic Wand (217287, -0.22 DPS) [crafted]; Wand of Mana Concentration (273043, -0.82 DPS, sim-verified) [world] |

**New at 30:** head: Filigreed Flame Circlet; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Expeditionary's Cape; chest: Beguiler Robes; wrist: Phoenix Bindings; waist: Highlander's Cloth Girdle; legs: Flame Leggings; feet: Fiery Slippers; finger1: Snake Hoop; finger2: Runemender's Seal; main_hand: Swinetusk Shank; off_hand: Rod of Molten Fire; ranged: Scorching Wand

No-known-source sample (15 of 326, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (human, 000000000000000000-54500001230133031-000000000000000000)

Set DPS (verified): 159.9. Weights run: 3.0s. Verify run: 5.2s. 456 eligible items had no known source.

Race sweep: human 159.9 ± 0.3, gnome 159.0 ± 0.3, high-order-skyborne 155.1 ± 0.3 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.606 ± 0.015, crit=0.188 ± 0.013 per rating point (14 rating = 1%, 2.631 per %), hit=0.320 ± 0.020 per rating point (10 rating = 1%, 3.204 per %), spell_haste=not significant (0.373 ± 0.132), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.228 ± 0.144), mp5=not significant (0.527 ± 0.173)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Miner's Hat of the Deep (9429, +0.00 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, +0.00 DPS) [crafted]; Corpseshroud (10574, -0.46 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Triune Amulet (7722, -0.56 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.71 DPS) [quest]; Necklace of Calisea (1714, -1.50 DPS) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.9 spell_power points (3.71 DPS) | yes | Green Silken Shoulders (7057, -0.05 DPS) [crafted]; Bloodmage Mantle (7684, -0.11 DPS) [dungeon]; Berylline Pads (4197, -0.45 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Long Silken Cloak (4326, +0.00 DPS) [crafted]; Guardian Cloak (5965, +0.00 DPS) [crafted]; Expeditionary's Cape (271922, -0.20 DPS) [rep] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Robe of the Magi (1716, +0.00 DPS) [world_drop]; Cindercloth Robe (10042, +0.00 DPS) [crafted]; Robe of Power (7054, -0.89 DPS) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (3.24 DPS) | yes | Spidertank Oilrag (9448, -0.60 DPS) [dungeon]; Arcane Runed Bracers (4744, -1.00 DPS) [quest]; Condor Bracers (15864, -1.50 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.0 spell_power points (5.49 DPS) | yes | Fiery Handwraps (254025, -0.94 DPS) [crafted]; Crimson Silk Gloves (7064, -1.09 DPS) [crafted]; Red Mageweave Gloves (10018, -1.24 DPS) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | 20.2 spell_power points (5.04 DPS) | yes | Highlander's Cloth Girdle (20098, -0.94 DPS) [rep]; Deathmage Sash (10771, -1.03 DPS) [dungeon]; Scorching Sash (4117, -1.29 DPS) [quest] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Red Mageweave Pants (10009, +0.00 DPS) [crafted]; Crimson Silk Pantaloons (7062, -0.87 DPS) [crafted]; Filigreed Flame Leggings (253941, -0.96 DPS) [crafted] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 18.2 spell_power points (4.53 DPS) | yes | Gilded Slippers (254001, -1.50 DPS) [crafted]; Acidic Walkers (9454, -1.84 DPS) [dungeon]; Spidersilk Boots (4320, -2.18 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lorekeeper's Ring (19524, +0.00 DPS) [rep]; Runemender's Seal (273031, -0.26 DPS) [world] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lorekeeper's Ring (19524, -0.61 DPS) [rep]; Runemender's Seal (273031, -1.25 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -6.82 DPS) [dungeon]; Celestial Stave (9517, -7.07 DPS) [quest]; Hypnotic Blade (7714, -7.40 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Ragefire Wand (7513) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plaguerot Sprig (10766, -0.86 DPS) [dungeon]; Jaina's Firestarter (13064, -1.28 DPS) [world_drop]; Twisted Nether Wand (249144, -1.68 DPS) [crafted] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Icy Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Fiery Cord; finger2: Philanthropist's Ring; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Staff of Jordan; ranged: Ragefire Wand

No-known-source sample (15 of 456, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (human, 230002000000000000-54500001230133031-003000000000000000)

Set DPS (verified): 261.1. Weights run: 3.1s. Verify run: 4.2s. 600 eligible items had no known source.

Race sweep: human 261.1 ± 0.4, gnome 258.9 ± 0.4, high-order-skyborne 252.4 ± 0.4 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.661 ± 0.028, crit=0.300 ± 0.025 per rating point (14 rating = 1%, 4.198 per %), hit=0.512 ± 0.039 per rating point (10 rating = 1%, 5.124 per %), spell_haste=not significant (0.758 ± 0.259), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.901 ± 0.285), mp5=not significant (0.950 ± 0.327)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Flame (3075) | World drop [world_drop] | 58.6 spell_power points (14.48 DPS) | yes | Dreamweave Circlet (10041, -4.99 DPS) [crafted]; Red Mageweave Headband (10033, -6.52 DPS) [crafted]; Gemburst Circlet (10751, -6.71 DPS) [quest] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 20.0 spell_power points (4.93 DPS) | yes | Mindburst Medallion (11196, -1.36 DPS) [quest]; Triune Amulet (7722, -1.40 DPS) [dungeon]; Lei of Lilies (1315, -1.60 DPS) [world_drop] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 28.0 spell_power points (6.92 DPS) | yes | Netherflame Shoulders (254053, -0.08 DPS) [crafted]; Rotgrip Mantle (17732, -0.77 DPS) [dungeon]; Flameseer Mantle (11310, -1.23 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cindercloth Cloak (14044, -1.26 DPS) [crafted]; Featherskin Cape (10843, -1.69 DPS) [world]; Runecloth Cloak (13860, -2.24 DPS) [crafted] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 44.6 spell_power points (11.02 DPS) | yes | Dreamweave Vest (10021, -1.99 DPS) [crafted]; Ghostweave Vest (14141, -2.47 DPS) [crafted]; Cindercloth Vest (14042, -2.65 DPS) [crafted] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 22.1 spell_power points (5.47 DPS) | yes | Nethergeld Cuffs (254061, -1.48 DPS) [crafted]; Aristocratic Cuffs (12546, -2.03 DPS) [dungeon]; Bloodband Bracers (11469, -2.09 DPS) [quest] |
| hands | Fiery Gloves (254099) | Tailoring [crafted] | 32.8 spell_power points (8.11 DPS) | yes | Dreamweave Gloves (10019, -1.45 DPS) [crafted]; Cindercloth Gloves (14043, -1.46 DPS) [crafted]; Runecloth Gloves (13863, -1.67 DPS) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fiery Cord (254041, +0.00 DPS) [crafted]; Ghostweave Cord (254073, +0.00 DPS) [crafted]; Fiery Waistcord (254085, +0.00 DPS) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kilt of the Atal'ai Prophet (10807, +0.00 DPS) [dungeon]; Senior Designer's Pantaloons (11841, +0.00 DPS) [dungeon]; Cindercloth Leggings (12256, +0.00 DPS) [vendor] |
| feet | Fiery Sandals (254111) | Tailoring [crafted] | 26.0 spell_power points (6.41 DPS) | yes | Mistwalker Boots (10629, -0.18 DPS) [dungeon]; Fiery Slippers (254005, -1.17 DPS) [crafted]; Cindercloth Boots (10044, -1.22 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 22.1 spell_power points (5.46 DPS) | yes | Lorekeeper's Ring (19523, -1.56 DPS) [rep]; Eye of Adaegus (5266, -1.81 DPS) [world_drop]; Philanthropist's Ring (281635, -2.18 DPS) [quest] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 17.2 spell_power points (4.26 DPS) | yes | Lorekeeper's Ring (19523, -0.35 DPS) [rep]; Eye of Adaegus (5266, -0.61 DPS) [world_drop]; Philanthropist's Ring (281635, -0.97 DPS) [quest] |
| trinket1 | Fire Ruby (20036) | Destroy Morphaz [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -0.79 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.23 DPS) [quest]; Uther's Strength (11302, -1.98 DPS) [world_drop] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Uther's Strength (11302, -0.30 DPS) [world_drop] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Staff of Jordan (873, +0.00 DPS) [world_drop]; Spire of Hakkar (10844, +0.00 DPS) [world] |
| off_hand | Truesilver Conduit (249455) | Enchanting [crafted] | 15.3 spell_power points (3.78 DPS) | yes | Celestial Orb (7515, -0.08 DPS) [quest]; Enthralled Sphere (11625, -0.16 DPS) [dungeon]; Orb of the Forgotten Seer (7685, -0.32 DPS) [dungeon] |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Noxious Shooter (17745, -3.36 DPS) [dungeon]; Lesser Eternal Wand (249232, -3.84 DPS) [crafted]; Wand of Allistarj (13065, -5.05 DPS) [world_drop] |

**New at 50:** head: Eye of Flame; neck: Arcane Crystal Pendant; shoulder: Kentic Amice; back: Spritecaster Cape; chest: Embrace of the Wind Serpent; wrist: Netherflame Cuffs; hands: Fiery Gloves; waist: Ban'thok Sash; legs: Spellshock Leggings; feet: Fiery Sandals; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Fire Ruby; trinket2: Thunderbrew's Boot Flask; main_hand: Blade of Eternal Darkness; off_hand: Truesilver Conduit; ranged: Pyric Caduceus

No-known-source sample (15 of 600, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (human, 230105200100000000-55500201230133031-003000000000000000)

Set DPS (verified): 408.4. Weights run: 3.5s. Verify run: 24.5s. 1459 eligible items had no known source.

Race sweep: human 408.4 ± 0.7, gnome 404.6 ± 0.7, high-order-skyborne 394.2 ± 0.7 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.581 ± 0.049, crit=0.441 ± 0.035 per rating point (14 rating = 1%, 6.177 per %), hit=1.207 ± 0.072 per rating point (10 rating = 1%, 12.073 per %), spell_haste=not significant (-3.764 ± 0.955), spell_penetration=not significant (0.000 ± 0.000), fire_power=0.836 ± 0.003, spirit=not significant (1.228 ± 0.618), mp5=not significant (1.065 ± 0.645)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 61.0 spell_power points (14.22 DPS) | yes | Field Marshal's Coronet (231604, -1.08 DPS) [pvp]; Eye of Flame (3075, -1.63 DPS) [world_drop]; Crimson Felt Hat (18727, -3.02 DPS) [dungeon] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 28.7 spell_power points (6.69 DPS) | yes | Star of Mystaria (12103, -0.09 DPS) [dungeon]; Nacreous Shell Necklace (22403, -0.70 DPS) [dungeon]; Dragon's Eye (10829, -1.47 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Argent Shoulders (19059, +0.00 DPS) [crafted]; Lieutenant Commander's Silk Mantle (227102, +0.00 DPS) [pvp]; Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, -1.47 DPS) [quest]; Spritecaster Cape (11623, -1.83 DPS) [dungeon]; Crystalline Threaded Cape (20697, -2.42 DPS) [world] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Embrace of the Wind Serpent (12462, +0.00 DPS) [world]; Field Marshal's Silk Vestments (231603, +0.00 DPS) [pvp]; Robes of the Exalted (13346, -1.12 DPS) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 35.2 spell_power points (8.21 DPS) | yes | Wyrmthalak's Shackles (13958, -2.29 DPS) [quest]; Marshal's Silk Bracers (16438, -2.78 DPS) [pvp]; Netherflame Cuffs (254065, -3.30 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Marshal's Silk Handwraps (231609, +0.00 DPS) [vendor] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Silk Sash (16439, +0.00 DPS) [pvp]; Thuzadin Sash (18740, +0.00 DPS) [dungeon]; Magician's Cord (272393, +0.00 DPS) [vendor] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Magister's Leggings (16687, +0.00 DPS) [dungeon]; Sentinel's Silk Leggings (22752, +0.00 DPS) [rep]; Marshal's Silk Leggings (231605, +0.00 DPS) [pvp] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wolfrunner Shoes (13101, +0.00 DPS) [world_drop]; Mooncloth Boots (15802, +0.00 DPS) [crafted]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Brainlash (6440, +0.00 DPS) [dungeon]; Maiden's Circle (13001, +0.00 DPS) [world_drop]; The Postmaster's Seal (13392, +0.00 DPS) [dungeon] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Postmaster's Seal (13392, -2.22 DPS) [dungeon]; Maiden's Circle (13001, -2.23 DPS) [world_drop]; Brainlash (6440, -3.75 DPS) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -0.76 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.33 DPS) [quest]; Smolderweb's Eye (13213, -1.38 DPS) [dungeon] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Headmaster's Charge (13937, +0.00 DPS) [dungeon]; Dancing Sliver (15854, +0.00 DPS) [quest]; Ironbark Staff (20069, -0.11 DPS) [rep] |
| off_hand | Tome of Fiery Arcana (19311) | Stormpike Guard [rep] | 30.9 spell_power points (7.21 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Therazane's Touch (19315, -0.45 DPS) [rep]; Penelope's Rose (15805, -1.23 DPS) [quest] |
| ranged | Oblivion's Touch (18761) | Scholomance: Death Knight Darkreaver [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -0.08 DPS) [dungeon]; Brilliant Wand (249385, -0.86 DPS) [crafted] |

**New at 60:** head: Spellweaver's Turban; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Devout Sandals; finger1: Blessed Band of Light; finger2: Channeler's Ring; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Sageblade; off_hand: Tome of Fiery Arcana; ranged: Oblivion's Touch

No-known-source sample (15 of 1459, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (human, 050105200100000000-55500201230133031-003000000000000000)

Set DPS (verified): 818.9. Weights run: 3.8s. Verify run: 18.3s. 1459 eligible items had no known source.

Race sweep: gnome 820.3 ± 0.8, human 818.9 ± 0.8, high-order-skyborne 799.4 ± 0.8 (dps); gnome beats the guide race human by 1.4 (0.18%) ± 1.2, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.570 ± 0.047, crit=0.903 ± 0.062 per rating point (14 rating = 1%, 12.643 per %), hit=1.549 ± 0.117 per rating point (10 rating = 1%, 15.488 per %), spell_haste=not significant (-5.209 ± 1.383), spell_penetration=not significant (0.000 ± 0.000), fire_power=0.866 ± 0.002, spirit=not significant (0.271 ± 0.251), mp5=not significant (0.388 ± 0.379)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 52.0 spell_power points (30.53 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -3.80 DPS) [pvp]; Eye of Flame (3075, -7.80 DPS, sim-verified) [world_drop] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 29.0 spell_power points (17.05 DPS) | yes | Star of Mystaria (12103, -1.78 DPS) [dungeon]; Lady Maye's Pendant (14558, -4.24 DPS) [world_drop]; Arcane Crystal Pendant (20037, -5.65 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 37.1 spell_power points (21.75 DPS) | yes | Lieutenant Commander's Silk Mantle (227102, -1.21 DPS) [pvp]; Field Marshal's Silk Spaulders (231602, -1.26 DPS) [pvp]; Shroud of the Nathrezim (18720, -6.47 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, -9.66 DPS) [quest]; Spritecaster Cape (11623, -9.98 DPS) [dungeon]; Crystalline Threaded Cape (20697, -11.95 DPS, sim-verified) [world] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Vestments (231603, -1.48 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -8.69 DPS) [pvp]; Alanna's Embrace (13314, -13.37 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 28.5 spell_power points (16.70 DPS) | yes | Runecloth Cuffs (254123, -6.90 DPS) [crafted]; Netherflame Cuffs (254065, -6.96 DPS) [crafted]; Manacle Cuffs (11962, -7.58 DPS) [quest] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 40.0 spell_power points (23.47 DPS) | yes | Marshal's Silk Gloves (16440, -2.81 DPS) [vendor]; Marshal's Silk Gauntlets (231608, -2.81 DPS) [vendor]; Hands of Power (13253, -6.46 DPS, sim-verified) [dungeon] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 33.2 spell_power points (19.48 DPS) | yes | Magician's Cord (272393, -0.30 DPS) [vendor]; Highlander's Cloth Girdle (20047, -1.83 DPS) [rep]; Clutch of Andros (13956, -4.03 DPS) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 63.0 spell_power points (36.98 DPS) | yes | Marshal's Silk Leggings (231605, -3.66 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -10.74 DPS) [pvp]; Skyshroud Leggings (13170, -15.35 DPS) [dungeon] |
| feet | Kayser's Boots of Precision (22231) | Blackrock Spire: War Master Voone [dungeon] | 29.7 spell_power points (17.46 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Sergeant Major's Dreadweave Boots (220891, -0.66 DPS) [vendor]; Dragonrider Boots (18102, -1.54 DPS) [dungeon] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -6.45 DPS) [dungeon]; Maiden's Circle (13001, -6.47 DPS) [world_drop]; Naglering (11669, -22.12 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Maiden's Circle (13001, -2.65 DPS) [world_drop]; Songstone of Ironforge (12543, -3.61 DPS) [quest]; Don Mauricio's Band of Domination (22433, -6.43 DPS, sim-verified) [dungeon] |
| trinket1 | Smolderweb's Eye (13213) | Blackrock Spire: Mother Smolderweb [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -0.91 DPS) [crafted]; Fire Ruby (20036, -1.98 DPS) [quest]; Cannonball Runner (13382, -4.55 DPS) [dungeon] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -2.38 DPS) [crafted]; Fire Ruby (20036, -3.45 DPS) [quest]; Cannonball Runner (13382, -6.02 DPS) [dungeon] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -1.43 DPS) [pvp]; Elemental Mage Staff (944, -12.20 DPS, sim-verified) [world_drop]; Staff of Jordan (873, -14.38 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bonecreeper Stylus (13938) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Oblivion's Touch (18761, -0.51 DPS) [dungeon]; Brilliant Wand (249385, -0.53 DPS) [crafted]; Pyric Caduceus (11748, -1.14 DPS) [dungeon] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; legs: Sentinel's Silk Leggings; feet: Kayser's Boots of Precision; finger1: Channeler's Ring; finger2: Don Julio's Band; trinket1: Smolderweb's Eye; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Bonecreeper Stylus

No-known-source sample (15 of 1459, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (undead, 000000000000000000-51400001000000000-000000000000000000)

Set DPS (verified): 40.3. Weights run: 2.4s. Verify run: 3.2s. 192 eligible items had no known source.

Race sweep: undead 40.3 ± 0.1, orc 39.2 ± 0.1, troll 39.0 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.920 ± 0.016, crit=0.147 ± 0.007 per rating point (14 rating = 1%, 2.061 per %), hit=0.305 ± 0.016 per rating point (10 rating = 1%, 3.054 per %), spell_haste=not significant (0.734 ± 0.198), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.316 ± 0.086), mp5=1.108 ± 0.243

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Flame Circlet (253953) | Tailoring [crafted] | 11.0 spell_power points (1.54 DPS) | yes | Pristine Circlet (253949, -0.70 DPS) [crafted]; Shadow Goggles (4373, -1.26 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.36 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.9 spell_power points (0.13 DPS) | yes | Scholarly Pendant (277203, -0.04 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.3 spell_power points (1.86 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.69 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.30 DPS) [crafted]; Slime-encrusted Pads (6461, -1.39 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sanguine Cape (14376, -0.22 DPS) [world_drop]; Battle Healer's Cloak (20427, -0.23 DPS) [rep]; Pearl-clasped Cloak (5542, -0.32 DPS, sim-verified) [crafted] |
| chest | Filigreed Flame Gown (253905) | Tailoring [crafted] | 14.5 spell_power points (2.04 DPS) | yes | Filigreed Pristine Gown (253901, -0.59 DPS, sim-verified) [crafted]; Gray Woolen Robe (2585, -0.83 DPS) [crafted]; Scholarly Robes (2034, -1.09 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 7.9 spell_power points (1.11 DPS) | yes | Tabitha's Cuffs (251486, -0.69 DPS) [quest]; Owlbeard Bracers (16981, -0.71 DPS) [quest]; Bright Bracers (3647, -0.83 DPS, sim-verified) [world_drop] |
| hands | Flame Gloves (253917) | Tailoring [crafted] | 9.8 spell_power points (1.37 DPS) | yes | Serpent Gloves (5970, -0.39 DPS) [dungeon]; Pristine Gloves (253913, -0.42 DPS) [crafted]; Phoenix Gloves (4331, -0.43 DPS, sim-verified) [crafted] |
| waist | Flame Sash (253929) | Tailoring [crafted] | 10.7 spell_power points (1.49 DPS) | yes | Pristine Sash (253925, -0.44 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.46 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.55 DPS) [crafted] |
| legs | Filigreed Flame Leggings (253941) | Tailoring [crafted] | 17.8 spell_power points (2.49 DPS) | yes | Filigreed Pristine Leggings (253937, -0.70 DPS) [crafted]; Phoenix Pants (4317, -0.94 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -1.32 DPS) [world_drop] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.7 spell_power points (1.49 DPS) | yes | Flame Boots (253893, -0.22 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.69 DPS) [crafted]; Red Woolen Boots (4313, -0.71 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 7.2 spell_power points (1.01 DPS) | yes | Loop of Sacrifice (281673, -0.37 DPS) [quest]; Volcanic Rock Ring (12053, -0.62 DPS) [world_drop]; Band of Purification (12996, -0.74 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.5 spell_power points (0.77 DPS) | yes | Loop of Sacrifice (281673, -0.13 DPS) [quest]; Volcanic Rock Ring (12053, -0.39 DPS) [world_drop]; Band of Purification (12996, -0.51 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Staff of the Blessed Seer (2271, -0.04 DPS) [dungeon]; Gnarled Necromancer's Staff (251534, -0.32 DPS) [quest]; Crescent Staff (6505, -0.43 DPS, sim-verified) [quest] |
| off_hand | Antipodean Rod (2879) | World drop [world_drop] | 6.0 spell_power points (0.84 DPS) | yes | Seer's Fine Stein (7608, -0.41 DPS) [world_drop]; Spellbinder Orb (15926, -0.45 DPS) [dungeon]; Tork Wrench (11855, -0.58 DPS) [quest] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller (12984, -0.30 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -1.44 DPS) [quest]; Greater Magic Wand (11288, -1.50 DPS) [crafted] |

**New at 20:** head: Flame Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Flame Gown; wrist: Mindthrust Bracers; hands: Flame Gloves; waist: Flame Sash; legs: Filigreed Flame Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Evocator's Blade; off_hand: Antipodean Rod; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 192, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 8350 The 1 Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 23173 Abomination Skin Leggings

### Band 30 (undead, 000000000000000000-54500001230100000-000000000000000000)

Set DPS (verified): 80.2. Weights run: 2.8s. Verify run: 4.7s. 318 eligible items had no known source.

Race sweep: undead 80.2 ± 0.1, orc 78.5 ± 0.1, troll 77.4 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.857 ± 0.015, crit=0.220 ± 0.012 per rating point (14 rating = 1%, 3.082 per %), hit=0.388 ± 0.021 per rating point (10 rating = 1%, 3.883 per %), spell_haste=0.444 ± 0.108, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.011 ± 0.096), mp5=not significant (0.342 ± 0.169)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Flame Circlet (253979) | Tailoring [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Enchanter's Cowl (4322, -0.71 DPS, sim-verified) [crafted]; Holy Shroud (2721, -0.72 DPS) [world_drop]; Flame Circlet (253953, -0.74 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 5.1 spell_power points (1.27 DPS) | yes | Darkspear Warding Pendant (272075, -0.21 DPS) [vendor]; Kaleidoscope Chain (13084, -0.41 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.41 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 16.7 spell_power points (4.12 DPS) | yes | Death Speaker Mantle (6685, -0.80 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.99 DPS) [world_drop]; Invoker's Mantle (215365, -1.34 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windsong Drape (15468, -0.45 DPS) [quest]; Hillman's Cloak (3719, -0.47 DPS) [crafted]; Resilient Cape (14400, -0.93 DPS, sim-verified) [world_drop] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Flame Gown (253965, -0.08 DPS) [crafted]; Mechbuilder's Overalls (9508, -0.27 DPS) [dungeon]; Green Silk Armor (7065, -0.84 DPS, sim-verified) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (3.21 DPS) | yes | Spidertank Oilrag (9448, -0.61 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -1.52 DPS) [world_drop]; Mindthrust Bracers (1974, -1.90 DPS) [dungeon] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 10.3 spell_power points (2.54 DPS) | yes | Flame Gloves (253917, -0.18 DPS) [crafted]; Phoenix Gloves (4331, -0.32 DPS) [crafted]; Truefaith Gloves (7049, -0.67 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 13.6 spell_power points (3.35 DPS) | yes | Belt of Arugal (6392, -0.48 DPS) [dungeon]; Invoker's Cord (215366, -0.56 DPS) [crafted]; Crimson Silk Belt (7055, -1.30 DPS, sim-verified) [crafted] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | 20.1 spell_power points (4.95 DPS) | yes | Filigreed Flame Leggings (253941, -1.35 DPS, sim-verified) [crafted]; Phoenix Pants (4317, -1.64 DPS) [crafted]; Pristine Leggings (253987, -1.73 DPS) [crafted] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 19.0 spell_power points (4.70 DPS) | yes | Gilded Slippers (254001, -1.58 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.76 DPS) [dungeon]; Spidersilk Boots (4320, -2.13 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 10.3 spell_power points (2.54 DPS) | yes | Advisor's Ring (19521, -0.64 DPS) [rep]; Ring of Power Regulation (273030, -0.80 DPS) [dungeon] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 8.2 spell_power points (2.03 DPS) | yes | Advisor's Ring (19521, -0.13 DPS) [rep]; Ring of Power Regulation (273030, -0.29 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.00 DPS) [dungeon]; Tendonscraper Dagger (273053, -1.00 DPS) [world]; Glimmering Staff (249392, -4.70 DPS, sim-verified) [crafted] |
| off_hand | Rod of Molten Fire (2565) | World drop [world_drop] | 13.0 spell_power points (3.21 DPS) | yes | Orb of Mystic Insight (249394, -1.09 DPS, sim-verified) [crafted]; Witch's Finger (16887, -1.52 DPS) [quest]; Antipodean Rod (2879, -1.73 DPS) [world_drop] |
| ranged | Scorching Wand (5213) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Starfaller (13063, +0.00 DPS) [world_drop]; Greater Mystic Wand (217287, -0.22 DPS) [crafted]; Wand of Mana Concentration (273043, -0.96 DPS, sim-verified) [world] |

**New at 30:** head: Filigreed Flame Circlet; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Darkspear Raider's Cloak; chest: Beguiler Robes; wrist: Phoenix Bindings; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Flame Leggings; feet: Fiery Slippers; finger1: Snake Hoop; finger2: Runemender's Seal; main_hand: Swinetusk Shank; off_hand: Rod of Molten Fire; ranged: Scorching Wand

No-known-source sample (15 of 318, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 8350 The 1 Ring; 9362 Brilliant Gold Ring

### Band 40 (undead, 000000000000000000-54500001230133031-000000000000000000)

Set DPS (verified): 158.8. Weights run: 3.0s. Verify run: 5.4s. 448 eligible items had no known source.

Race sweep: undead 158.8 ± 0.3, orc 156.7 ± 0.3, troll 154.5 ± 0.3 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.606 ± 0.015, crit=0.188 ± 0.013 per rating point (14 rating = 1%, 2.631 per %), hit=0.320 ± 0.020 per rating point (10 rating = 1%, 3.204 per %), spell_haste=not significant (0.373 ± 0.132), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.228 ± 0.144), mp5=not significant (0.527 ± 0.173)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Miner's Hat of the Deep (9429) | Uldaman [dungeon] | 22.3 spell_power points (5.56 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.32 DPS) [crafted]; Papal Fez (9431, -0.49 DPS) [dungeon]; Corpseshroud (10574, -0.94 DPS) [dungeon] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.15 DPS) [quest]; Necklace of Calisea (1714, -0.94 DPS) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.9 spell_power points (3.71 DPS) | yes | Green Silken Shoulders (7057, -0.05 DPS) [crafted]; Bloodmage Mantle (7684, -0.11 DPS) [dungeon]; Berylline Pads (4197, -0.45 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.16 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.36 DPS) [vendor] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Robe of the Magi (1716, +0.00 DPS) [world_drop]; Cindercloth Robe (10042, +0.00 DPS) [crafted]; Robe of Power (7054, -0.89 DPS) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (3.24 DPS) | yes | Spidertank Oilrag (9448, -0.60 DPS) [dungeon]; Radiant Silver Bracers (4545, -1.04 DPS) [quest]; Condor Bracers (15864, -1.50 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.0 spell_power points (5.49 DPS) | yes | Fiery Handwraps (254025, -0.94 DPS) [crafted]; Crimson Silk Gloves (7064, -1.09 DPS) [crafted]; Red Mageweave Gloves (10018, -1.24 DPS) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | 20.2 spell_power points (5.04 DPS) | yes | Defiler's Cloth Girdle (20166, -0.94 DPS) [rep]; Deathmage Sash (10771, -1.03 DPS) [dungeon]; Scorching Sash (4117, -1.29 DPS) [quest] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Red Mageweave Pants (10009, +0.00 DPS) [crafted]; Crimson Silk Pantaloons (7062, -0.87 DPS) [crafted]; Filigreed Flame Leggings (253941, -0.96 DPS) [crafted] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 18.2 spell_power points (4.53 DPS) | yes | Gilded Slippers (254001, -1.50 DPS) [crafted]; Acidic Walkers (9454, -1.84 DPS) [dungeon]; Spidersilk Boots (4320, -2.18 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.0 spell_power points (3.25 DPS) | yes | Snake Hoop (6750, -1.00 DPS) [quest]; Runemender's Seal (273031, -1.25 DPS) [world]; Ring of Power Regulation (273030, -1.36 DPS) [dungeon] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 10.6 spell_power points (2.64 DPS) | yes | Snake Hoop (6750, -0.39 DPS) [quest]; Runemender's Seal (273031, -0.64 DPS) [world]; Ring of Power Regulation (273030, -0.75 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -6.82 DPS) [dungeon]; Celestial Stave (9517, -7.07 DPS) [quest]; Hypnotic Blade (7714, -7.40 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Ragefire Wand (7513) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plaguerot Sprig (10766, -0.83 DPS) [dungeon]; Jaina's Firestarter (13064, -1.25 DPS) [world_drop]; Twisted Nether Wand (249144, -1.73 DPS) [crafted] |

**New at 40:** head: Miner's Hat of the Deep; neck: Triune Amulet; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Fiery Cord; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Staff of Jordan; ranged: Ragefire Wand

No-known-source sample (15 of 448, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (undead, 230002000000000000-54500001230133031-003000000000000000)

Set DPS (verified): 257.9. Weights run: 3.1s. Verify run: 4.3s. 592 eligible items had no known source.

Race sweep: undead 257.9 ± 0.4, orc 255.9 ± 0.4, troll 252.2 ± 0.4 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.661 ± 0.028, crit=0.300 ± 0.025 per rating point (14 rating = 1%, 4.198 per %), hit=0.512 ± 0.039 per rating point (10 rating = 1%, 5.124 per %), spell_haste=not significant (0.758 ± 0.259), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.901 ± 0.285), mp5=not significant (0.950 ± 0.327)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Flame (3075) | World drop [world_drop] | 58.6 spell_power points (14.48 DPS) | yes | Dreamweave Circlet (10041, -4.99 DPS) [crafted]; Red Mageweave Headband (10033, -6.52 DPS) [crafted]; Gemburst Circlet (10751, -6.71 DPS) [quest] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 20.0 spell_power points (4.93 DPS) | yes | Mindburst Medallion (11196, -1.36 DPS) [quest]; Triune Amulet (7722, -1.40 DPS) [dungeon]; Lei of Lilies (1315, -1.60 DPS) [world_drop] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 28.0 spell_power points (6.92 DPS) | yes | Netherflame Shoulders (254053, -0.08 DPS) [crafted]; Rotgrip Mantle (17732, -0.77 DPS) [dungeon]; Flameseer Mantle (11310, -1.23 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cindercloth Cloak (14044, -1.26 DPS) [crafted]; Deep Woodlands Cloak (19121, -1.34 DPS) [quest]; Featherskin Cape (10843, -1.69 DPS) [world] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 44.6 spell_power points (11.02 DPS) | yes | Dreamweave Vest (10021, -1.99 DPS) [crafted]; Ghostweave Vest (14141, -2.47 DPS) [crafted]; Cindercloth Vest (14042, -2.65 DPS) [crafted] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 22.1 spell_power points (5.47 DPS) | yes | Nethergeld Cuffs (254061, -1.48 DPS) [crafted]; Aristocratic Cuffs (12546, -2.03 DPS) [dungeon]; Bloodband Bracers (11469, -2.09 DPS) [quest] |
| hands | Fiery Gloves (254099) | Tailoring [crafted] | 32.8 spell_power points (8.11 DPS) | yes | Dreamweave Gloves (10019, -1.45 DPS) [crafted]; Cindercloth Gloves (14043, -1.46 DPS) [crafted]; Runecloth Gloves (13863, -1.67 DPS) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fiery Cord (254041, +0.00 DPS) [crafted]; Ghostweave Cord (254073, +0.00 DPS) [crafted]; Fiery Waistcord (254085, +0.00 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 32.8 spell_power points (8.11 DPS) | yes | Cindercloth Leggings (12256, -0.08 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -0.43 DPS) [dungeon]; Spellshock Leggings (9484, -0.80 DPS) [dungeon] |
| feet | Fiery Sandals (254111) | Tailoring [crafted] | 26.0 spell_power points (6.41 DPS) | yes | Mistwalker Boots (10629, -0.18 DPS) [dungeon]; Fiery Slippers (254005, -1.17 DPS) [crafted]; Cindercloth Boots (10044, -1.22 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 22.1 spell_power points (5.46 DPS) | yes | Advisor's Ring (19519, -1.56 DPS) [rep]; Eye of Adaegus (5266, -1.81 DPS) [world_drop]; Philanthropist's Ring (281635, -2.18 DPS) [quest] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 17.2 spell_power points (4.26 DPS) | yes | Advisor's Ring (19519, -0.35 DPS) [rep]; Eye of Adaegus (5266, -0.61 DPS) [world_drop]; Philanthropist's Ring (281635, -0.97 DPS) [quest] |
| trinket1 | Fire Ruby (20036) | Destroy Morphaz [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -1.23 DPS) [quest]; Uther's Strength (11302, -1.98 DPS) [world_drop]; Darkspear Voodoo Seal (272061, -2.28 DPS) [vendor] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.44 DPS) [quest]; Uther's Strength (11302, -1.19 DPS) [world_drop]; Darkspear Voodoo Seal (272061, -1.50 DPS) [vendor] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Staff of Jordan (873, +0.00 DPS) [world_drop]; Spire of Hakkar (10844, +0.00 DPS) [world] |
| off_hand | Truesilver Conduit (249455) | Enchanting [crafted] | 15.3 spell_power points (3.78 DPS) | yes | Celestial Orb (7515, -0.08 DPS) [quest]; Enthralled Sphere (11625, -0.16 DPS) [dungeon]; Flask of Forest Mojo (19115, -0.25 DPS) [quest] |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Noxious Shooter (17745, -3.38 DPS) [dungeon]; Lesser Eternal Wand (249232, -3.90 DPS) [crafted]; Wand of Allistarj (13065, -5.09 DPS) [world_drop] |

**New at 50:** head: Eye of Flame; neck: Arcane Crystal Pendant; shoulder: Kentic Amice; back: Spritecaster Cape; chest: Embrace of the Wind Serpent; wrist: Netherflame Cuffs; hands: Fiery Gloves; waist: Ban'thok Sash; legs: Senior Designer's Pantaloons; feet: Fiery Sandals; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Fire Ruby; trinket2: Ankh of Life; main_hand: Blade of Eternal Darkness; off_hand: Truesilver Conduit; ranged: Pyric Caduceus

No-known-source sample (15 of 592, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (undead, 230105200100000000-55500201230133031-003000000000000000)

Set DPS (verified): 401.8. Weights run: 3.5s. Verify run: 25.0s. 1453 eligible items had no known source.

Race sweep: undead 401.8 ± 0.7, orc 396.4 ± 0.7, troll 394.4 ± 0.7 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.581 ± 0.049, crit=0.441 ± 0.035 per rating point (14 rating = 1%, 6.177 per %), hit=1.207 ± 0.072 per rating point (10 rating = 1%, 12.073 per %), spell_haste=not significant (-3.764 ± 0.955), spell_penetration=not significant (0.000 ± 0.000), fire_power=0.836 ± 0.003, spirit=not significant (1.228 ± 0.618), mp5=not significant (1.065 ± 0.645)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 61.0 spell_power points (14.22 DPS) | yes | Warlord's Silk Cowl (231601, -1.08 DPS) [pvp]; Eye of Flame (3075, -1.63 DPS) [world_drop]; Crimson Felt Hat (18727, -3.02 DPS) [dungeon] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 28.7 spell_power points (6.69 DPS) | yes | Star of Mystaria (12103, -0.09 DPS) [dungeon]; Nacreous Shell Necklace (22403, -0.70 DPS) [dungeon]; Dragon's Eye (10829, -1.47 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Argent Shoulders (19059, +0.00 DPS) [crafted]; Champion's Silk Mantle (227104, +0.00 DPS) [pvp]; Warlord's Silk Amice (231594, +0.00 DPS) [pvp] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, -1.47 DPS) [quest]; Spritecaster Cape (11623, -1.83 DPS) [dungeon]; Crystalline Threaded Cape (20697, -2.42 DPS) [world] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Embrace of the Wind Serpent (12462, +0.00 DPS) [world]; Warlord's Silk Raiment (231596, +0.00 DPS) [pvp]; Robes of the Exalted (13346, -1.12 DPS) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 35.2 spell_power points (8.21 DPS) | yes | Wyrmthalak's Shackles (13958, -2.29 DPS) [quest]; General's Silk Cuffs (16538, -2.78 DPS) [pvp]; Netherflame Cuffs (254065, -3.30 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; General's Silk Handwraps (231600, +0.00 DPS) [vendor] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Silk Sash (16537, +0.00 DPS) [pvp]; Thuzadin Sash (18740, +0.00 DPS) [dungeon]; Magician's Cord (272393, +0.00 DPS) [vendor] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Magister's Leggings (16687, +0.00 DPS) [dungeon]; Outrider's Silk Leggings (22747, +0.00 DPS) [rep]; General's Silk Trousers (231595, +0.00 DPS) [pvp] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wolfrunner Shoes (13101, +0.00 DPS) [world_drop]; Mooncloth Boots (15802, +0.00 DPS) [crafted]; General's Silk Boots (231597, +0.00 DPS) [pvp] |
| finger1 | The Postmaster's Seal (13392) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Maiden's Circle (13001, -0.01 DPS) [world_drop]; Brainlash (6440, -1.52 DPS) [dungeon]; Eye of Orgrimmar (12545, -1.73 DPS) [quest] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Maiden's Circle (13001, -2.23 DPS) [world_drop]; Brainlash (6440, -3.75 DPS) [dungeon]; Eye of Orgrimmar (12545, -3.95 DPS) [quest] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -0.76 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.33 DPS) [quest]; Smolderweb's Eye (13213, -1.38 DPS) [dungeon] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Headmaster's Charge (13937, +0.00 DPS) [dungeon]; Dancing Sliver (15854, +0.00 DPS) [quest]; Ironbark Staff (20220, -0.11 DPS) [rep] |
| off_hand | Tome of Fiery Arcana (19311) | Frostwolf Clan [rep] | 30.9 spell_power points (7.21 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; Therazane's Touch (19315, -0.45 DPS) [rep]; Penelope's Rose (15805, -1.23 DPS) [quest] |
| ranged | Bonecreeper Stylus (13938) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Oblivion's Touch (18761, -0.24 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -0.31 DPS) [dungeon]; Brilliant Wand (249385, -1.12 DPS) [crafted] |

**New at 60:** head: Spellweaver's Turban; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Devout Sandals; finger1: The Postmaster's Seal; finger2: Channeler's Ring; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Sageblade; off_hand: Tome of Fiery Arcana; ranged: Bonecreeper Stylus

No-known-source sample (15 of 1453, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60, raid preset (undead, 050105200100000000-55500201230133031-003000000000000000)

Set DPS (verified): 805.9. Weights run: 3.8s. Verify run: 20.0s. 1453 eligible items had no known source.

Race sweep: undead 805.9 ± 0.8, orc 802.5 ± 0.8, troll 797.2 ± 0.8 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.570 ± 0.047, crit=0.903 ± 0.062 per rating point (14 rating = 1%, 12.643 per %), hit=1.549 ± 0.117 per rating point (10 rating = 1%, 15.488 per %), spell_haste=not significant (-5.209 ± 1.383), spell_penetration=not significant (0.000 ± 0.000), fire_power=0.866 ± 0.002, spirit=not significant (0.271 ± 0.251), mp5=not significant (0.388 ± 0.379)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 52.0 spell_power points (30.53 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -3.80 DPS) [pvp]; Eye of Flame (3075, -6.08 DPS, sim-verified) [world_drop] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 29.0 spell_power points (17.05 DPS) | yes | Lady Maye's Pendant (14558, -4.24 DPS) [world_drop]; Star of Mystaria (12103, -4.43 DPS, sim-verified) [dungeon]; Arcane Crystal Pendant (20037, -5.65 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 37.1 spell_power points (21.75 DPS) | yes | Champion's Silk Mantle (227104, -1.21 DPS) [pvp]; Warlord's Silk Amice (231594, -1.26 DPS) [pvp]; Shroud of the Nathrezim (18720, -5.48 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -9.06 DPS, sim-verified) [world]; Shroud of the Exile (15421, -9.66 DPS) [quest]; Spritecaster Cape (11623, -9.98 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Raiment (231596, -1.48 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -8.69 DPS) [pvp]; Alanna's Embrace (13314, -20.25 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 28.5 spell_power points (16.70 DPS) | yes | Runecloth Cuffs (254123, -6.90 DPS) [crafted]; Netherflame Cuffs (254065, -6.96 DPS) [crafted]; Manacle Cuffs (11962, -7.58 DPS) [quest] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 40.0 spell_power points (23.47 DPS) | yes | General's Silk Handguards (16540, -2.81 DPS) [vendor]; General's Silk Gauntlets (231599, -2.81 DPS) [vendor]; Hands of Power (13253, -5.98 DPS, sim-verified) [dungeon] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 33.2 spell_power points (19.48 DPS) | yes | Defiler's Cloth Girdle (20163, -1.83 DPS) [rep]; Clutch of Andros (13956, -4.03 DPS) [dungeon]; Magician's Cord (272393, -4.37 DPS, sim-verified) [vendor] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 41.5 spell_power points (24.38 DPS) | yes | Legionnaire's Silk Legguards (227107, +0.00 DPS) [pvp]; General's Silk Trousers (231595, +0.00 DPS) [pvp]; Skyshroud Leggings (13170, -5.10 DPS, sim-verified) [dungeon] |
| feet | Kayser's Boots of Precision (22231) | Blackrock Spire: War Master Voone [dungeon] | 29.7 spell_power points (17.46 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; First Sergeant's Dreadweave Boots (220909, -0.66 DPS) [vendor]; Dragonrider Boots (18102, -1.54 DPS) [dungeon] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -6.45 DPS) [dungeon]; Maiden's Circle (13001, -6.47 DPS) [world_drop]; Naglering (11669, -19.29 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -2.63 DPS) [dungeon]; Maiden's Circle (13001, -2.65 DPS) [world_drop]; Eye of Orgrimmar (12545, -3.61 DPS) [quest] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+3.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -2.99 DPS, sim-verified) [crafted]; Fire Ruby (20036, -3.45 DPS) [quest]; Rune of the Guard Captain (19120, -4.20 DPS) [quest] |
| trinket2 | Smolderweb's Eye (13213) | Blackrock Spire: Mother Smolderweb [dungeon] | sim-verified (-1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -0.91 DPS) [crafted]; Fire Ruby (20036, -1.98 DPS) [quest]; Rune of the Guard Captain (19120, -2.73 DPS) [quest] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -1.43 DPS) [pvp]; Elemental Mage Staff (944, -14.30 DPS, sim-verified) [world_drop]; Staff of Jordan (873, -14.38 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bonecreeper Stylus (13938) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Oblivion's Touch (18761, -0.51 DPS) [dungeon]; Brilliant Wand (249385, -0.53 DPS) [crafted]; Pyric Caduceus (11748, -1.15 DPS) [dungeon] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; legs: Outrider's Silk Leggings; feet: Kayser's Boots of Precision; finger1: Channeler's Ring; finger2: Don Julio's Band; trinket1: Weakness Analyzer; trinket2: Smolderweb's Eye; main_hand: Ironbark Staff; ranged: Bonecreeper Stylus

No-known-source sample (15 of 1453, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

