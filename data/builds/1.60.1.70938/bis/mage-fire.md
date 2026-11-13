# Leveling BiS: Fire

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 000000000000000000-51400001000000000-000000000000000000)

Set DPS (verified): 39.9. Weights run: 3.2s. Verify run: 2.8s. 198 eligible items had no known source.

Race sweep: gnome 40.4 ± 0.1, human 39.9 ± 0.1, high-order-skyborne 39.7 ± 0.1 (dps); gnome beats the guide race human by 0.5 (1.25%) ± 0.1, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.903 ± 0.016, crit=0.151 ± 0.007 per rating point (14 rating = 1%, 2.109 per %), hit=0.310 ± 0.016 per rating point (10 rating = 1%, 3.097 per %), spell_haste=not significant (0.700 ± 0.199), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.262 ± 0.079), mp5=1.438 ± 0.212

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Flame Circlet (253953) | Tailoring [crafted] | 11.0 spell_power points (1.55 DPS) | yes | Pristine Circlet (253949, -0.70 DPS) [crafted]; Shadow Goggles (4373, -1.13 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.40 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.5 spell_power points (0.07 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.1 spell_power points (1.85 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.74 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -1.24 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -1.28 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pearl-clasped Cloak (5542, -0.05 DPS) [crafted]; Sanguine Cape (14376, -0.20 DPS) [world_drop]; Caretaker's Cape (20428, -0.21 DPS) [rep] |
| chest | Filigreed Flame Gown (253905) | Tailoring [crafted] | 14.3 spell_power points (2.01 DPS) | yes | Filigreed Pristine Gown (253901, -0.57 DPS, sim-verified) [crafted]; Gray Woolen Robe (2585, -0.81 DPS) [crafted]; Mystic's Wrap (14369, -1.12 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 8.8 spell_power points (1.24 DPS) | yes | Bright Bracers (3647, -0.64 DPS, sim-verified) [world_drop]; Repurposed Hair Band (281256, -0.91 DPS) [quest]; Mystic's Bracelets (14366, -0.99 DPS) [world_drop] |
| hands | Flame Gloves (253917) | Tailoring [crafted] | 9.7 spell_power points (1.37 DPS) | yes | Phoenix Gloves (4331, -0.28 DPS, sim-verified) [crafted]; Serpent Gloves (5970, -0.38 DPS) [dungeon]; Pristine Gloves (253913, -0.42 DPS) [crafted] |
| waist | Flame Sash (253929) | Tailoring [crafted] | 10.6 spell_power points (1.49 DPS) | yes | Pristine Sash (253925, -0.43 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.48 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.55 DPS) [crafted] |
| legs | Filigreed Flame Leggings (253941) | Tailoring [crafted] | 17.5 spell_power points (2.46 DPS) | yes | Phoenix Pants (4317, -0.61 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.70 DPS) [crafted]; Darkweave Breeches (12987, -1.35 DPS) [world_drop] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.6 spell_power points (1.49 DPS) | yes | Flame Boots (253893, -0.43 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.69 DPS) [crafted]; Red Woolen Boots (4313, -0.75 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 7.9 spell_power points (1.11 DPS) | yes | Lavishly Jeweled Ring (1156, -0.35 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.73 DPS) [world_drop]; Band of Purification (12996, -0.89 DPS) [world_drop] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.8 spell_power points (0.96 DPS) | yes | Lavishly Jeweled Ring (1156, -0.51 DPS, sim-verified) [dungeon]; Volcanic Rock Ring (12053, -0.58 DPS) [world_drop]; Band of Purification (12996, -0.74 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | 31.1 spell_power points (4.38 DPS) | yes | Emberstone Staff (5201, -0.82 DPS) [dungeon]; Staff of the Blessed Seer (2271, -1.13 DPS, sim-verified) [dungeon]; Staff of Westfall (2042, -1.27 DPS) [quest] |
| off_hand | Antipodean Rod (2879) | World drop [world_drop] | 6.0 spell_power points (0.84 DPS) | yes | Tear of Grief (5611, -0.32 DPS) [quest]; Dwarven Tome (279898, -0.39 DPS) [quest]; Mathystra Channeling Orb (281291, -0.51 DPS, sim-verified) [quest] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller (12984, -0.30 DPS) [world_drop]; Deepblaze (279896, -1.51 DPS) [quest]; Sizzle Stick (8071, -1.58 DPS) [quest] |

**New at 20:** head: Flame Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Flame Gown; wrist: Mindthrust Bracers; hands: Flame Gloves; waist: Flame Sash; legs: Filigreed Flame Leggings; feet: Spidersilk Boots; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Evocator's Blade; off_hand: Antipodean Rod; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 198, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 8350 The 1 Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (human, 000000000000000000-54500001230100000-000000000000000000)

Set DPS (verified): 80.0. Weights run: 3.7s. Verify run: 5.2s. 326 eligible items had no known source.

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

Set DPS (verified): 160.1. Weights run: 3.8s. Verify run: 5.0s. 456 eligible items had no known source.

Race sweep: human 160.1 ± 0.3, gnome 159.3 ± 0.3, high-order-skyborne 155.4 ± 0.3 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.580 ± 0.015, crit=0.186 ± 0.012 per rating point (14 rating = 1%, 2.604 per %), hit=0.321 ± 0.020 per rating point (10 rating = 1%, 3.210 per %), spell_haste=not significant (0.283 ± 0.143), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.262 ± 0.154), mp5=not significant (0.522 ± 0.172)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Miner's Hat of the Deep (9429, +0.00 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, +0.00 DPS) [crafted]; Corpseshroud (10574, -0.55 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.6 spell_power points (2.89 DPS) | yes | Triune Amulet (7722, -0.45 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.66 DPS) [quest]; Necklace of Calisea (1714, -1.43 DPS) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.5 spell_power points (3.62 DPS) | yes | Green Silken Shoulders (7057, -0.04 DPS) [crafted]; Bloodmage Mantle (7684, -0.08 DPS) [dungeon]; Berylline Pads (4197, -0.43 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Long Silken Cloak (4326, +0.00 DPS) [crafted]; Guardian Cloak (5965, +0.00 DPS) [crafted]; Expeditionary's Cape (271922, -0.28 DPS) [rep] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cindercloth Robe (10042, +0.00 DPS) [crafted]; Robe of the Magi (1716, -0.03 DPS) [world_drop]; Robe of Power (7054, -0.95 DPS) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (3.23 DPS) | yes | Spidertank Oilrag (9448, -0.61 DPS) [dungeon]; Arcane Runed Bracers (4744, -0.99 DPS) [quest]; Condor Bracers (15864, -1.49 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.2 spell_power points (5.51 DPS) | yes | Fiery Handwraps (254025, -1.02 DPS) [crafted]; Crimson Silk Gloves (7064, -1.16 DPS) [crafted]; Red Mageweave Gloves (10018, -1.33 DPS) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | 20.2 spell_power points (5.03 DPS) | yes | Highlander's Cloth Girdle (20098, -0.97 DPS) [rep]; Deathmage Sash (10771, -1.12 DPS) [dungeon]; Scorching Sash (4117, -1.21 DPS) [quest] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Red Mageweave Pants (10009, +0.00 DPS) [crafted]; Crimson Silk Pantaloons (7062, -0.95 DPS) [crafted]; Filigreed Flame Leggings (253941, -0.96 DPS) [crafted] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 18.1 spell_power points (4.50 DPS) | yes | Gilded Slippers (254001, -1.49 DPS) [crafted]; Acidic Walkers (9454, -1.85 DPS) [dungeon]; Spidersilk Boots (4320, -2.19 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lorekeeper's Ring (19524, +0.00 DPS) [rep]; Runemender's Seal (273031, -0.21 DPS) [world] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lorekeeper's Ring (19524, -0.58 DPS) [rep]; Runemender's Seal (273031, -1.20 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -6.74 DPS) [dungeon]; Celestial Stave (9517, -7.06 DPS) [quest]; Hypnotic Blade (7714, -7.36 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Ragefire Wand (7513) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plaguerot Sprig (10766, -0.84 DPS) [dungeon]; Jaina's Firestarter (13064, -1.27 DPS) [world_drop]; Twisted Nether Wand (249144, -1.66 DPS) [crafted] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Icy Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Fiery Cord; finger2: Philanthropist's Ring; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Staff of Jordan; ranged: Ragefire Wand

No-known-source sample (15 of 456, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (human, 230002000000000000-54500001230133031-003000000000000000)

Set DPS (verified): 258.7. Weights run: 3.9s. Verify run: 8.3s. 600 eligible items had no known source.

Race sweep: human 258.7 ± 0.4, gnome 256.5 ± 0.4, high-order-skyborne 250.1 ± 0.4 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.650 ± 0.029, crit=0.270 ± 0.024 per rating point (14 rating = 1%, 3.782 per %), hit=0.525 ± 0.040 per rating point (10 rating = 1%, 5.252 per %), spell_haste=not significant (0.654 ± 0.288), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.674 ± 0.288), mp5=not significant (0.788 ± 0.325)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Flame (3075) | World drop [world_drop] | 56.2 spell_power points (13.60 DPS) | yes | Red Mageweave Headband (10033, -5.86 DPS) [crafted]; Gemburst Circlet (10751, -7.03 DPS) [quest]; Dreamweave Circlet (10041, -7.26 DPS, sim-verified) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 19.9 spell_power points (4.81 DPS) | yes | Triune Amulet (7722, -1.74 DPS) [dungeon]; Talvash's Enhancing Necklace (7673, -1.75 DPS) [quest]; Mindburst Medallion (11196, -2.62 DPS, sim-verified) [quest] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 26.5 spell_power points (6.41 DPS) | yes | Netherflame Shoulders (254053, -0.06 DPS) [crafted]; Rotgrip Mantle (17732, -0.43 DPS) [dungeon]; Flameseer Mantle (11310, -1.39 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cindercloth Cloak (14044, -0.90 DPS) [crafted]; Runecloth Cloak (13860, -1.87 DPS) [crafted]; Cloak of Fire (14134, -1.92 DPS) [crafted] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 37.7 spell_power points (9.12 DPS) | yes | Cindercloth Vest (14042, -1.52 DPS) [crafted]; Robe of the Magi (1716, -2.04 DPS) [world_drop]; Dreamweave Vest (10021, -3.35 DPS, sim-verified) [crafted] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 20.9 spell_power points (5.06 DPS) | yes | Aristocratic Cuffs (12546, -1.73 DPS) [dungeon]; Phoenix Bindings (210781, -1.92 DPS) [crafted]; Nethergeld Cuffs (254061, -2.27 DPS, sim-verified) [crafted] |
| hands | Fiery Gloves (254099) | Tailoring [crafted] | 30.9 spell_power points (7.47 DPS) | yes | Cindercloth Gloves (14043, -1.57 DPS) [crafted]; Runecloth Gloves (13863, -1.69 DPS) [crafted]; Dreamweave Gloves (10019, -2.83 DPS, sim-verified) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Fiery Cord (254041, -0.12 DPS) [crafted]; Ghostweave Cord (254073, -0.57 DPS) [crafted]; Fiery Waistcord (254085, -1.05 DPS, sim-verified) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | 29.5 spell_power points (7.13 DPS) | yes | Senior Designer's Pantaloons (11841, -0.20 DPS) [dungeon]; Kilt of the Atal'ai Prophet (10807, -0.65 DPS) [dungeon]; Cindercloth Leggings (12256, -3.23 DPS, sim-verified) [vendor] |
| feet | Fiery Sandals (254111) | Tailoring [crafted] | 25.9 spell_power points (6.25 DPS) | yes | Cindercloth Boots (10044, -1.17 DPS) [crafted]; Fiery Slippers (254005, -1.36 DPS) [crafted]; Mistwalker Boots (10629, -6.02 DPS, sim-verified) [dungeon] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 20.9 spell_power points (5.05 DPS) | yes | Lorekeeper's Ring (19523, -1.38 DPS) [rep]; Philanthropist's Ring (281635, -1.84 DPS) [quest]; Band of the Unicorn (7553, -1.90 DPS) [world_drop] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 16.2 spell_power points (3.93 DPS) | yes | Lorekeeper's Ring (19523, -0.26 DPS) [rep]; Philanthropist's Ring (281635, -0.72 DPS) [quest]; Band of the Unicorn (7553, -0.79 DPS) [world_drop] |
| trinket1 | Fire Ruby (20036) | Destroy Morphaz [quest] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, -1.70 DPS, sim-verified) [world_drop]; Evonice's Landin' Pilla (18951, -1.76 DPS) [quest]; Uther's Strength (11302, -1.93 DPS) [world_drop] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (-2.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Spire of Hakkar (10844, +0.00 DPS) [world]; Staff of Jordan (873, -6.94 DPS, sim-verified) [world_drop] |
| off_hand | Truesilver Conduit (249455) | Enchanting [crafted] | 15.2 spell_power points (3.68 DPS) | yes | Celestial Orb (7515, -0.06 DPS) [quest]; Orb of the Forgotten Seer (7685, -0.29 DPS) [dungeon]; Enthralled Sphere (11625, -0.32 DPS) [dungeon] |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Noxious Shooter (17745, -3.19 DPS) [dungeon]; Lesser Eternal Wand (249232, -4.57 DPS, sim-verified) [crafted]; Wand of Allistarj (13065, -4.88 DPS) [world_drop] |

**New at 50:** head: Eye of Flame; neck: Arcane Crystal Pendant; shoulder: Kentic Amice; back: Spritecaster Cape; chest: Embrace of the Wind Serpent; wrist: Netherflame Cuffs; hands: Fiery Gloves; waist: Ban'thok Sash; legs: Spellshock Leggings; feet: Fiery Sandals; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Fire Ruby; trinket2: Thunderbrew's Boot Flask; main_hand: Blade of Eternal Darkness; off_hand: Truesilver Conduit; ranged: Pyric Caduceus

No-known-source sample (15 of 600, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (human, 230105200100000000-55500201230133031-003000000000000000)

Set DPS (verified): 400.0. Weights run: 4.6s. Verify run: 34.1s. 1459 eligible items had no known source.

Race sweep: human 400.0 ± 0.7, gnome 395.0 ± 0.7, high-order-skyborne 387.3 ± 0.7 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.540 ± 0.054, crit=0.457 ± 0.035 per rating point (14 rating = 1%, 6.403 per %), hit=1.270 ± 0.074 per rating point (10 rating = 1%, 12.698 per %), spell_haste=not significant (-3.592 ± 0.965), spell_penetration=not significant (0.000 ± 0.000), fire_power=0.839 ± 0.003, spirit=not significant (1.490 ± 0.655), mp5=not significant (0.832 ± 0.666)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Coronet (231604, -1.69 DPS) [pvp]; Eye of Flame (3075, -1.94 DPS) [world_drop]; Crimson Felt Hat (18727, -3.31 DPS) [dungeon] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Dragon's Eye (10829, -0.96 DPS) [dungeon]; Nacreous Shell Necklace (22403, -1.03 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Nathrezim (18720, +0.00 DPS) [dungeon]; Argent Shoulders (19059, +0.00 DPS) [crafted]; Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, -1.16 DPS) [quest]; Spritecaster Cape (11623, -1.53 DPS) [dungeon]; Featherskin Cape (10843, -1.72 DPS) [world] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Embrace of the Wind Serpent (12462, +0.00 DPS) [world]; Robes of the Exalted (13346, +0.00 DPS) [dungeon]; Field Marshal's Silk Vestments (231603, +0.00 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 36.8 spell_power points (8.23 DPS) | yes | Wyrmthalak's Shackles (13958, -1.66 DPS) [quest]; Marshal's Silk Bracers (16438, -2.60 DPS) [pvp]; Devout Bracers (16697, -3.03 DPS) [dungeon] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Marshal's Silk Handwraps (231609, +0.00 DPS) [vendor] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sash of the Burning Heart (11807, +0.00 DPS) [dungeon]; Marshal's Silk Sash (16439, +0.00 DPS) [pvp]; Thuzadin Sash (18740, +0.00 DPS) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ghostweave Pants (14144, +0.00 DPS) [crafted]; Magister's Leggings (16687, +0.00 DPS) [dungeon]; Marshal's Silk Leggings (231605, +0.00 DPS) [pvp] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mistwalker Boots (10629, +0.00 DPS) [dungeon]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Ghostweave Boots (254129, +0.00 DPS) [crafted] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Postmaster's Seal (13392, -0.47 DPS) [dungeon]; Maiden's Circle (13001, -0.87 DPS) [world_drop]; Brainlash (6440, -2.41 DPS) [dungeon] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Postmaster's Seal (13392, +0.00 DPS) [dungeon]; Maiden's Circle (13001, -0.33 DPS) [world_drop]; Brainlash (6440, -1.87 DPS) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -0.03 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.70 DPS) [quest]; Smolderweb's Eye (13213, -1.19 DPS) [dungeon] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Headmaster's Charge (13937, +0.00 DPS) [dungeon]; Dancing Sliver (15854, +0.00 DPS) [quest]; Ironbark Staff (20069, -0.23 DPS) [rep] |
| off_hand | Tome of Fiery Arcana (19311) | Stormpike Guard [rep] | 31.1 spell_power points (6.96 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Therazane's Touch (19315, -0.46 DPS) [rep]; Penelope's Rose (15805, -0.72 DPS) [quest] |
| ranged | Oblivion's Touch (18761) | Scholomance: Death Knight Darkreaver [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, +0.00 DPS) [dungeon]; Brilliant Wand (249385, -1.10 DPS) [crafted] |

**New at 60:** head: Spellweaver's Turban; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Devout Sandals; finger1: Blessed Band of Light; finger2: Channeler's Ring; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Sageblade; off_hand: Tome of Fiery Arcana; ranged: Oblivion's Touch

No-known-source sample (15 of 1459, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (human, 050105200100000000-55500201230133031-003000000000000000)

Set DPS (verified): 813.1. Weights run: 5.1s. Verify run: 21.8s. 1459 eligible items had no known source.

Race sweep: gnome 814.5 ± 0.8, human 813.1 ± 0.8, high-order-skyborne 794.0 ± 0.8 (dps); gnome beats the guide race human by 1.4 (0.17%) ± 1.2, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.562 ± 0.048, crit=0.879 ± 0.057 per rating point (14 rating = 1%, 12.306 per %), hit=1.626 ± 0.110 per rating point (10 rating = 1%, 16.261 per %), spell_haste=not significant (-5.222 ± 1.315), spell_penetration=not significant (0.000 ± 0.000), fire_power=0.866 ± 0.002, spirit=not significant (-0.149 ± 0.243), mp5=not significant (-0.063 ± 0.350)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 49.3 spell_power points (29.23 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -3.46 DPS) [pvp]; Eye of Flame (3075, -4.05 DPS, sim-verified) [world_drop] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 29.8 spell_power points (17.66 DPS) | yes | Star of Mystaria (12103, -2.78 DPS) [dungeon]; Lady Maye's Pendant (14558, -4.79 DPS) [world_drop]; Arcane Crystal Pendant (20037, -6.16 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 36.6 spell_power points (21.73 DPS) | yes | Lieutenant Commander's Silk Mantle (227102, -1.85 DPS) [pvp]; Field Marshal's Silk Spaulders (231602, -1.89 DPS) [pvp]; Burial Shawl (18681, -5.51 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -4.60 DPS, sim-verified) [world]; Mageflame Cloak (13007, -11.02 DPS) [world_drop]; Shroud of the Exile (15421, -11.36 DPS) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Vestments (231603, -2.49 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -9.61 DPS) [pvp]; Alanna's Embrace (13314, -18.52 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 26.5 spell_power points (15.72 DPS) | yes | Runecloth Cuffs (254123, -5.86 DPS) [crafted]; Manacle Cuffs (11962, -6.57 DPS) [quest]; Netherflame Cuffs (254065, -6.71 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 39.2 spell_power points (23.28 DPS) | yes | Marshal's Silk Gloves (16440, -3.26 DPS) [vendor]; Marshal's Silk Gauntlets (231608, -3.26 DPS) [vendor]; Hands of Power (13253, -6.34 DPS, sim-verified) [dungeon] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 33.9 spell_power points (20.11 DPS) | yes | Magician's Cord (272393, -0.83 DPS) [vendor]; Highlander's Cloth Girdle (20047, -2.49 DPS) [rep]; Clutch of Andros (13956, -4.12 DPS) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 62.2 spell_power points (36.93 DPS) | yes | Marshal's Silk Leggings (231605, -5.16 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -11.50 DPS) [pvp]; Skyshroud Leggings (13170, -15.13 DPS) [dungeon] |
| feet | Kayser's Boots of Precision (22231) | Blackrock Spire: War Master Voone [dungeon] | 30.4 spell_power points (18.03 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Sergeant Major's Dreadweave Boots (220891, -0.63 DPS) [vendor]; Dragonrider Boots (18102, -2.01 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -2.63 DPS) [vendor]; Don Mauricio's Band of Domination (22433, -3.12 DPS) [dungeon]; Naglering (11669, -15.96 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -1.27 DPS) [vendor]; Don Mauricio's Band of Domination (22433, -1.75 DPS) [dungeon]; Maiden's Circle (13001, -2.57 DPS) [world_drop] |
| trinket1 | Eye of the Beast (13968) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Draconic Infused Emblem (22268, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, -1.03 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -2.00 DPS) [crafted]; Blackhand's Breadth (13965, -6.05 DPS, sim-verified) [quest] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -1.49 DPS) [pvp]; Elemental Mage Staff (944, -13.66 DPS, sim-verified) [world_drop]; Staff of Jordan (873, -15.08 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Oblivion's Touch (18761, +0.00 DPS) [dungeon]; Pyric Caduceus (11748, -1.07 DPS) [dungeon] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; legs: Sentinel's Silk Leggings; feet: Kayser's Boots of Precision; finger1: Don Julio's Band; finger2: Blessed Band of Light; trinket1: Eye of the Beast; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1459, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (undead, 000000000000000000-51400001000000000-000000000000000000)

Set DPS (verified): 40.3. Weights run: 3.2s. Verify run: 3.6s. 192 eligible items had no known source.

Race sweep: undead 40.3 ± 0.1, orc 39.2 ± 0.1, troll 39.0 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.903 ± 0.016, crit=0.151 ± 0.007 per rating point (14 rating = 1%, 2.109 per %), hit=0.310 ± 0.016 per rating point (10 rating = 1%, 3.097 per %), spell_haste=not significant (0.700 ± 0.199), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.262 ± 0.079), mp5=1.438 ± 0.212

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Flame Circlet (253953) | Tailoring [crafted] | 11.0 spell_power points (1.55 DPS) | yes | Pristine Circlet (253949, -0.70 DPS) [crafted]; Shadow Goggles (4373, -1.26 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.40 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.8 spell_power points (0.11 DPS) | yes | Scholarly Pendant (277203, -0.04 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.1 spell_power points (1.85 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.69 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -1.24 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -1.28 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sanguine Cape (14376, -0.20 DPS) [world_drop]; Battle Healer's Cloak (20427, -0.21 DPS) [rep]; Pearl-clasped Cloak (5542, -0.32 DPS, sim-verified) [crafted] |
| chest | Filigreed Flame Gown (253905) | Tailoring [crafted] | 14.3 spell_power points (2.01 DPS) | yes | Filigreed Pristine Gown (253901, -0.59 DPS, sim-verified) [crafted]; Gray Woolen Robe (2585, -0.81 DPS) [crafted]; Mystic's Wrap (14369, -1.12 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 8.8 spell_power points (1.24 DPS) | yes | Tabitha's Cuffs (251486, -0.82 DPS) [quest]; Bright Bracers (3647, -0.83 DPS, sim-verified) [world_drop]; Owlbeard Bracers (16981, -0.85 DPS) [quest] |
| hands | Flame Gloves (253917) | Tailoring [crafted] | 9.7 spell_power points (1.37 DPS) | yes | Serpent Gloves (5970, -0.38 DPS) [dungeon]; Pristine Gloves (253913, -0.42 DPS) [crafted]; Phoenix Gloves (4331, -0.43 DPS, sim-verified) [crafted] |
| waist | Flame Sash (253929) | Tailoring [crafted] | 10.6 spell_power points (1.49 DPS) | yes | Pristine Sash (253925, -0.44 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.48 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.55 DPS) [crafted] |
| legs | Filigreed Flame Leggings (253941) | Tailoring [crafted] | 17.5 spell_power points (2.46 DPS) | yes | Filigreed Pristine Leggings (253937, -0.70 DPS) [crafted]; Phoenix Pants (4317, -0.94 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -1.35 DPS) [world_drop] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.6 spell_power points (1.49 DPS) | yes | Flame Boots (253893, -0.22 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.69 DPS) [crafted]; Red Woolen Boots (4313, -0.75 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 7.9 spell_power points (1.11 DPS) | yes | Loop of Sacrifice (281673, -0.47 DPS) [quest]; Volcanic Rock Ring (12053, -0.73 DPS) [world_drop]; Band of Purification (12996, -0.89 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.4 spell_power points (0.76 DPS) | yes | Loop of Sacrifice (281673, -0.13 DPS) [quest]; Volcanic Rock Ring (12053, -0.38 DPS) [world_drop]; Band of Purification (12996, -0.54 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Staff of the Blessed Seer (2271, -0.07 DPS) [dungeon]; Gnarled Necromancer's Staff (251534, -0.35 DPS) [quest]; Crescent Staff (6505, -0.43 DPS, sim-verified) [quest] |
| off_hand | Antipodean Rod (2879) | World drop [world_drop] | 6.0 spell_power points (0.84 DPS) | yes | Seer's Fine Stein (7608, -0.43 DPS) [world_drop]; Spellbinder Orb (15926, -0.46 DPS) [dungeon]; Tork Wrench (11855, -0.59 DPS) [quest] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller (12984, -0.30 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -1.44 DPS) [quest]; Greater Magic Wand (11288, -1.50 DPS) [crafted] |

**New at 20:** head: Flame Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Flame Gown; wrist: Mindthrust Bracers; hands: Flame Gloves; waist: Flame Sash; legs: Filigreed Flame Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Evocator's Blade; off_hand: Antipodean Rod; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 192, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 8350 The 1 Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 23173 Abomination Skin Leggings

### Band 30 (undead, 000000000000000000-54500001230100000-000000000000000000)

Set DPS (verified): 80.2. Weights run: 3.7s. Verify run: 6.5s. 318 eligible items had no known source.

Race sweep: undead 80.2 ± 0.1, orc 78.5 ± 0.1, troll 77.4 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.857 ± 0.015, crit=0.220 ± 0.012 per rating point (14 rating = 1%, 3.082 per %), hit=0.388 ± 0.021 per rating point (10 rating = 1%, 3.883 per %), spell_haste=0.444 ± 0.108, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.011 ± 0.096), mp5=not significant (0.342 ± 0.169)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Flame Circlet (253979) | Tailoring [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Enchanter's Cowl (4322, -0.71 DPS, sim-verified) [crafted]; Holy Shroud (2721, -0.72 DPS) [world_drop]; Flame Circlet (253953, -0.74 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 5.1 spell_power points (1.27 DPS) | yes | Darkspear Warding Pendant (272075, -0.21 DPS) [vendor]; Kaleidoscope Chain (13084, -0.41 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.41 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 16.7 spell_power points (4.12 DPS) | yes | Mantle of Woe (7750, -0.35 DPS) [quest]; Death Speaker Mantle (6685, -0.80 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.99 DPS) [world_drop] |
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
| ranged | Scorching Wand (5213) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Starfaller (13063, -0.07 DPS) [world_drop]; Greater Mystic Wand (217287, -0.23 DPS) [crafted]; Wand of Mana Concentration (273043, -0.96 DPS, sim-verified) [world] |

**New at 30:** head: Filigreed Flame Circlet; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Darkspear Raider's Cloak; chest: Beguiler Robes; wrist: Phoenix Bindings; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Flame Leggings; feet: Fiery Slippers; finger1: Snake Hoop; finger2: Runemender's Seal; main_hand: Swinetusk Shank; off_hand: Rod of Molten Fire; ranged: Scorching Wand

No-known-source sample (15 of 318, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 8350 The 1 Ring; 9362 Brilliant Gold Ring

### Band 40 (undead, 000000000000000000-54500001230133031-000000000000000000)

Set DPS (verified): 158.8. Weights run: 3.8s. Verify run: 6.3s. 448 eligible items had no known source.

Race sweep: undead 158.8 ± 0.3, orc 156.4 ± 0.3, troll 154.5 ± 0.3 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.580 ± 0.015, crit=0.186 ± 0.012 per rating point (14 rating = 1%, 2.604 per %), hit=0.321 ± 0.020 per rating point (10 rating = 1%, 3.210 per %), spell_haste=not significant (0.283 ± 0.143), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.262 ± 0.154), mp5=not significant (0.522 ± 0.172)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Miner's Hat of the Deep (9429) | Uldaman [dungeon] | 21.9 spell_power points (5.44 DPS) | yes | Papal Fez (9431, -0.41 DPS) [dungeon]; Corpseshroud (10574, -0.95 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -1.22 DPS, sim-verified) [crafted] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Prodigious Shadowshard Pendant (17773, -0.21 DPS) [quest]; Glowing Eye of Mordresh (10769, -0.80 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -0.98 DPS) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.5 spell_power points (3.62 DPS) | yes | Green Silken Shoulders (7057, -0.04 DPS) [crafted]; Bloodmage Mantle (7684, -0.08 DPS) [dungeon]; Berylline Pads (4197, -0.43 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.08 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.37 DPS) [vendor] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Robe of the Magi (1716, -0.03 DPS) [world_drop]; Robe of Power (7054, -0.95 DPS) [crafted]; Cindercloth Robe (10042, -2.43 DPS, sim-verified) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (3.23 DPS) | yes | Spidertank Oilrag (9448, -0.61 DPS) [dungeon]; Radiant Silver Bracers (4545, -1.08 DPS) [quest]; Condor Bracers (15864, -1.49 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.2 spell_power points (5.51 DPS) | yes | Crimson Silk Gloves (7064, -1.16 DPS) [crafted]; Red Mageweave Gloves (10018, -1.33 DPS) [crafted]; Fiery Handwraps (254025, -2.97 DPS, sim-verified) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | 20.2 spell_power points (5.03 DPS) | yes | Deathmage Sash (10771, -1.12 DPS) [dungeon]; Scorching Sash (4117, -1.21 DPS) [quest]; Defiler's Cloth Girdle (20166, -3.48 DPS, sim-verified) [rep] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Red Mageweave Pants (10009, -0.62 DPS, sim-verified) [crafted]; Crimson Silk Pantaloons (7062, -0.95 DPS) [crafted]; Filigreed Flame Leggings (253941, -0.96 DPS) [crafted] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 18.1 spell_power points (4.50 DPS) | yes | Acidic Walkers (9454, -1.85 DPS) [dungeon]; Gilded Slippers (254001, -2.10 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -2.19 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.9 spell_power points (3.21 DPS) | yes | Snake Hoop (6750, -0.99 DPS) [quest]; Runemender's Seal (273031, -1.20 DPS) [world]; Ring of Power Regulation (273030, -1.33 DPS) [dungeon] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 10.6 spell_power points (2.63 DPS) | yes | Runemender's Seal (273031, -0.62 DPS) [world]; Ring of Power Regulation (273030, -0.75 DPS) [dungeon]; Snake Hoop (6750, -1.61 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+4.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Celestial Stave (9517, -7.06 DPS) [quest]; Hypnotic Blade (7714, -7.36 DPS) [dungeon]; Illusionary Rod (7713, -10.25 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Ragefire Wand (7513) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plaguerot Sprig (10766, -0.81 DPS) [dungeon]; Jaina's Firestarter (13064, -1.24 DPS) [world_drop]; Dancing Flame (6806, -1.59 DPS, sim-verified) [quest] |

**New at 40:** head: Miner's Hat of the Deep; neck: Triune Amulet; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Fiery Cord; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Staff of Jordan; ranged: Ragefire Wand

No-known-source sample (15 of 448, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (undead, 230002000000000000-54500001230133031-003000000000000000)

Set DPS (verified): 254.9. Weights run: 3.9s. Verify run: 5.7s. 592 eligible items had no known source.

Race sweep: undead 254.9 ± 0.4, orc 252.4 ± 0.4, troll 250.0 ± 0.4 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.690 ± 0.032, crit=0.513 ± 0.032 per rating point (14 rating = 1%, 7.182 per %), hit=0.772 ± 0.053 per rating point (10 rating = 1%, 7.718 per %), spell_haste=not significant (0.825 ± 0.280), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.858 ± 0.338), mp5=not significant (0.783 ± 0.388)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Flame (3075) | World drop [world_drop] | 58.5 spell_power points (21.40 DPS) | yes | Dreamweave Circlet (10041, -8.43 DPS, sim-verified) [crafted]; Red Mageweave Headband (10033, -9.40 DPS) [crafted]; Gemburst Circlet (10751, -9.99 DPS) [quest] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 20.1 spell_power points (7.37 DPS) | yes | Triune Amulet (7722, -2.24 DPS) [dungeon]; Horizon Choker (13085, -2.58 DPS) [world_drop]; Mindburst Medallion (11196, -3.63 DPS, sim-verified) [quest] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 28.1 spell_power points (10.29 DPS) | yes | Netherflame Shoulders (254053, -0.16 DPS) [crafted]; Rotgrip Mantle (17732, -0.99 DPS) [dungeon]; Flameseer Mantle (11310, -2.02 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cindercloth Cloak (14044, -1.85 DPS, sim-verified) [crafted]; Deep Woodlands Cloak (19121, -1.86 DPS) [quest]; Featherskin Cape (10843, -2.71 DPS) [world] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 43.6 spell_power points (15.97 DPS) | yes | Dreamweave Vest (10021, -2.39 DPS, sim-verified) [crafted]; Cindercloth Vest (14042, -3.73 DPS) [crafted]; Ghostweave Vest (14141, -4.10 DPS) [crafted] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 22.1 spell_power points (8.09 DPS) | yes | Nethergeld Cuffs (254061, -2.24 DPS, sim-verified) [crafted]; Aristocratic Cuffs (12546, -2.84 DPS) [dungeon]; Bloodband Bracers (11469, -3.05 DPS) [quest] |
| hands | Fiery Gloves (254099) | Tailoring [crafted] | 32.8 spell_power points (11.99 DPS) | yes | Cindercloth Gloves (14043, -2.31 DPS) [crafted]; Runecloth Gloves (13863, -2.50 DPS) [crafted]; Dreamweave Gloves (10019, -4.14 DPS, sim-verified) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 26.6 spell_power points (9.74 DPS) | yes | Fiery Waistcord (254085, -0.52 DPS) [crafted]; Fiery Cord (254041, -0.71 DPS) [crafted]; Ghostweave Cord (254073, -1.16 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 32.3 spell_power points (11.81 DPS) | yes | Cindercloth Leggings (12256, -0.15 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -0.52 DPS) [dungeon]; Spellshock Leggings (9484, -0.87 DPS) [dungeon] |
| feet | Fiery Sandals (254111) | Tailoring [crafted] | 26.2 spell_power points (9.59 DPS) | yes | First Sergeant's Dreadweave Boots (220909, -1.57 DPS) [vendor]; Fiery Slippers (254005, -1.81 DPS) [crafted]; Mistwalker Boots (10629, -5.43 DPS, sim-verified) [dungeon] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 22.2 spell_power points (8.12 DPS) | yes | Advisor's Ring (19519, -2.58 DPS) [rep]; Eye of Adaegus (5266, -2.84 DPS) [world_drop]; Philanthropist's Ring (281635, -3.20 DPS) [quest] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 17.3 spell_power points (6.32 DPS) | yes | Advisor's Ring (19519, -0.78 DPS) [rep]; Eye of Adaegus (5266, -1.03 DPS) [world_drop]; Philanthropist's Ring (281635, -1.39 DPS) [quest] |
| trinket1 | Fire Ruby (20036) | Destroy Morphaz [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -2.58 DPS) [crafted]; Uther's Strength (11302, -2.93 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -4.37 DPS, sim-verified) [quest] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.63 DPS) [quest]; Frozen Heart of the Mountain (249469, -1.23 DPS) [crafted]; Uther's Strength (11302, -1.57 DPS) [world_drop] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Spire of Hakkar (10844, +0.00 DPS) [world]; Staff of Jordan (873, -5.34 DPS, sim-verified) [world_drop] |
| off_hand | Truesilver Conduit (249455) | Enchanting [crafted] | 15.5 spell_power points (5.68 DPS) | yes | Celestial Orb (7515, -0.16 DPS) [quest]; Enthralled Sphere (11625, -0.27 DPS) [dungeon]; Flask of Forest Mojo (19115, -0.37 DPS) [quest] |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Noxious Shooter (17745, -4.64 DPS) [dungeon]; Lesser Eternal Wand (249232, -5.35 DPS, sim-verified) [crafted]; Ragefire Wand (7513, -6.35 DPS) [quest] |

**New at 50:** head: Eye of Flame; neck: Arcane Crystal Pendant; shoulder: Kentic Amice; back: Spritecaster Cape; chest: Embrace of the Wind Serpent; wrist: Netherflame Cuffs; hands: Fiery Gloves; waist: Ban'thok Sash; legs: Senior Designer's Pantaloons; feet: Fiery Sandals; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Fire Ruby; trinket2: Ankh of Life; main_hand: Blade of Eternal Darkness; off_hand: Truesilver Conduit; ranged: Pyric Caduceus

No-known-source sample (15 of 592, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (undead, 230105200100000000-55500201230133031-003000000000000000)

Set DPS (verified): 392.0. Weights run: 4.6s. Verify run: 31.6s. 1453 eligible items had no known source.

Race sweep: undead 392.0 ± 0.7, orc 387.4 ± 0.7, troll 384.9 ± 0.7 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.540 ± 0.054, crit=0.457 ± 0.035 per rating point (14 rating = 1%, 6.403 per %), hit=1.270 ± 0.074 per rating point (10 rating = 1%, 12.698 per %), spell_haste=not significant (-3.592 ± 0.965), spell_penetration=not significant (0.000 ± 0.000), fire_power=0.839 ± 0.003, spirit=not significant (1.490 ± 0.655), mp5=not significant (0.832 ± 0.666)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 65.1 spell_power points (14.58 DPS) | yes | Warlord's Silk Cowl (231601, -1.69 DPS) [pvp]; Eye of Flame (3075, -1.94 DPS) [world_drop]; Crimson Felt Hat (18727, -3.31 DPS) [dungeon] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Dragon's Eye (10829, -0.96 DPS) [dungeon]; Nacreous Shell Necklace (22403, -1.03 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Nathrezim (18720, +0.00 DPS) [dungeon]; Argent Shoulders (19059, +0.00 DPS) [crafted]; Warlord's Silk Amice (231594, +0.00 DPS) [pvp] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, -1.16 DPS) [quest]; Spritecaster Cape (11623, -1.53 DPS) [dungeon]; Featherskin Cape (10843, -1.72 DPS) [world] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Embrace of the Wind Serpent (12462, +0.00 DPS) [world]; Robes of the Exalted (13346, +0.00 DPS) [dungeon]; Warlord's Silk Raiment (231596, +0.00 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 36.8 spell_power points (8.23 DPS) | yes | Wyrmthalak's Shackles (13958, -1.66 DPS) [quest]; General's Silk Cuffs (16538, -2.60 DPS) [pvp]; Devout Bracers (16697, -3.03 DPS) [dungeon] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; General's Silk Handwraps (231600, +0.00 DPS) [vendor] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sash of the Burning Heart (11807, +0.00 DPS) [dungeon]; General's Silk Sash (16537, +0.00 DPS) [pvp]; Thuzadin Sash (18740, +0.00 DPS) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ghostweave Pants (14144, +0.00 DPS) [crafted]; Magister's Leggings (16687, +0.00 DPS) [dungeon]; General's Silk Trousers (231595, +0.00 DPS) [pvp] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mistwalker Boots (10629, +0.00 DPS) [dungeon]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Ghostweave Boots (254129, +0.00 DPS) [crafted] |
| finger1 | The Postmaster's Seal (13392) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, +0.00 DPS) [vendor]; Maiden's Circle (13001, -0.40 DPS) [world_drop]; Brainlash (6440, -1.94 DPS) [dungeon] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, +0.00 DPS) [vendor]; Maiden's Circle (13001, -0.33 DPS) [world_drop]; Brainlash (6440, -1.87 DPS) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -0.03 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.70 DPS) [quest]; Smolderweb's Eye (13213, -1.19 DPS) [dungeon] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Headmaster's Charge (13937, +0.00 DPS) [dungeon]; Dancing Sliver (15854, +0.00 DPS) [quest]; Ironbark Staff (20220, -0.23 DPS) [rep] |
| off_hand | Tome of Fiery Arcana (19311) | Frostwolf Clan [rep] | 31.1 spell_power points (6.96 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; Therazane's Touch (19315, -0.46 DPS) [rep]; Penelope's Rose (15805, -0.72 DPS) [quest] |
| ranged | Bonecreeper Stylus (13938) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ritssyn's Wand of Bad Mojo (22408, -0.13 DPS) [dungeon]; Oblivion's Touch (18761, -0.26 DPS) [dungeon]; Brilliant Wand (249385, -1.40 DPS) [crafted] |

**New at 60:** head: Spellweaver's Turban; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Devout Sandals; finger1: The Postmaster's Seal; finger2: Channeler's Ring; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Sageblade; off_hand: Tome of Fiery Arcana; ranged: Bonecreeper Stylus

No-known-source sample (15 of 1453, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60, raid preset (undead, 050105200100000000-55500201230133031-003000000000000000)

Set DPS (verified): 800.4. Weights run: 5.1s. Verify run: 22.6s. 1453 eligible items had no known source.

Race sweep: undead 800.4 ± 0.8, orc 797.7 ± 0.8, troll 791.4 ± 0.8 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.562 ± 0.048, crit=0.879 ± 0.057 per rating point (14 rating = 1%, 12.306 per %), hit=1.626 ± 0.110 per rating point (10 rating = 1%, 16.261 per %), spell_haste=not significant (-5.222 ± 1.315), spell_penetration=not significant (0.000 ± 0.000), fire_power=0.866 ± 0.002, spirit=not significant (-0.149 ± 0.243), mp5=not significant (-0.063 ± 0.350)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 49.3 spell_power points (29.23 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -3.46 DPS) [pvp]; Eye of Flame (3075, -3.81 DPS) [world_drop] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 29.8 spell_power points (17.66 DPS) | yes | Star of Mystaria (12103, -2.78 DPS) [dungeon]; Lady Maye's Pendant (14558, -4.79 DPS) [world_drop]; Arcane Crystal Pendant (20037, -6.16 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 36.6 spell_power points (21.73 DPS) | yes | Champion's Silk Mantle (227104, -1.85 DPS) [pvp]; Warlord's Silk Amice (231594, -1.89 DPS) [pvp]; Burial Shawl (18681, -4.95 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -5.27 DPS, sim-verified) [world]; Mageflame Cloak (13007, -11.02 DPS) [world_drop]; Shroud of the Exile (15421, -11.36 DPS) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Raiment (231596, -2.49 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -9.61 DPS) [pvp]; Alanna's Embrace (13314, -16.64 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 26.5 spell_power points (15.72 DPS) | yes | Runecloth Cuffs (254123, -5.86 DPS) [crafted]; Manacle Cuffs (11962, -6.57 DPS) [quest]; Netherflame Cuffs (254065, -6.71 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 39.2 spell_power points (23.28 DPS) | yes | General's Silk Handguards (16540, -3.26 DPS) [vendor]; General's Silk Gauntlets (231599, -3.26 DPS) [vendor]; Hands of Power (13253, -7.35 DPS, sim-verified) [dungeon] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 33.9 spell_power points (20.11 DPS) | yes | Magician's Cord (272393, -1.42 DPS, sim-verified) [vendor]; Defiler's Cloth Girdle (20163, -2.49 DPS) [rep]; Clutch of Andros (13956, -4.12 DPS) [dungeon] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 38.7 spell_power points (22.95 DPS) | yes | Legionnaire's Silk Legguards (227107, +0.00 DPS) [pvp]; General's Silk Trousers (231595, +0.00 DPS) [pvp]; Skyshroud Leggings (13170, -4.96 DPS, sim-verified) [dungeon] |
| feet | Kayser's Boots of Precision (22231) | Blackrock Spire: War Master Voone [dungeon] | 30.4 spell_power points (18.03 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -2.01 DPS) [dungeon]; First Sergeant's Dreadweave Boots (220909, -4.35 DPS, sim-verified) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -2.63 DPS) [vendor]; Don Mauricio's Band of Domination (22433, -3.12 DPS) [dungeon]; Naglering (11669, -17.94 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -1.27 DPS) [vendor]; Don Mauricio's Band of Domination (22433, -1.75 DPS) [dungeon]; Maiden's Circle (13001, -2.57 DPS) [world_drop] |
| trinket1 | Eye of the Beast (13968) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Fire Ruby (20036, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, -1.03 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -2.00 DPS) [crafted]; Fire Ruby (20036, -3.49 DPS) [quest] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -1.49 DPS) [pvp]; Elemental Mage Staff (944, -13.86 DPS, sim-verified) [world_drop]; Staff of Jordan (873, -15.08 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bonecreeper Stylus (13938) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Oblivion's Touch (18761, -0.52 DPS) [dungeon]; Brilliant Wand (249385, -0.54 DPS) [crafted]; Pyric Caduceus (11748, -1.59 DPS) [dungeon] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; legs: Outrider's Silk Leggings; feet: Kayser's Boots of Precision; finger1: Don Julio's Band; finger2: Blessed Band of Light; trinket1: Eye of the Beast; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Bonecreeper Stylus

No-known-source sample (15 of 1453, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

