# Leveling BiS: Fire

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 000000000000000000-51400001000000000-000000000000000000)

Set DPS (verified): 38.6. Weights run: 2.6s. Verify run: 2.5s. 198 eligible items had no known source.

Race sweep: gnome 39.6 ± 0.1, human 38.6 ± 0.1, high-order-skyborne 38.3 ± 0.1 (dps); gnome beats the guide race human by 1.1 (2.73%) ± 0.1, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.018 ± 0.016, crit=0.151 ± 0.007 per rating point (14 rating = 1%, 2.117 per %), hit=0.305 ± 0.016 per rating point (10 rating = 1%, 3.051 per %), spell_haste=not significant (0.571 ± 0.209), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.144 ± 0.082), mp5=1.059 ± 0.219

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Flame Circlet (253953) | Tailoring [crafted] | 11.0 spell_power points (1.55 DPS) | yes | Shadow Goggles (4373, -0.71 DPS) [crafted]; Pristine Circlet (253949, -0.74 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.47 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.3 spell_power points (0.04 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 14.2 spell_power points (1.99 DPS) | yes | Reinforced Woolen Shoulders (4315, -1.17 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.43 DPS) [crafted]; Slime-encrusted Pads (6461, -1.55 DPS) [dungeon] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.07 DPS) [crafted]; Sanguine Cape (14376, -0.14 DPS) [world_drop]; Caretaker's Cape (20428, -0.25 DPS) [rep] |
| chest | Filigreed Flame Gown (253905) | Tailoring [crafted] | 14.5 spell_power points (2.04 DPS) | yes | Filigreed Pristine Gown (253901, -0.59 DPS, sim-verified) [crafted]; Gray Woolen Robe (2585, -0.76 DPS) [crafted]; Mystic's Wrap (14369, -1.04 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 8.3 spell_power points (1.16 DPS) | yes | Repurposed Hair Band (281256, -0.84 DPS) [quest]; Mystic's Bracelets (14366, -0.88 DPS) [world_drop]; Bright Bracers (3647, -0.91 DPS, sim-verified) [world_drop] |
| hands | Flame Gloves (253917) | Tailoring [crafted] | 10.1 spell_power points (1.41 DPS) | yes | Phoenix Gloves (4331, -0.30 DPS, sim-verified) [crafted]; Pristine Gloves (253913, -0.42 DPS) [crafted]; Serpent Gloves (5970, -0.43 DPS) [dungeon] |
| waist | Flame Sash (253929) | Tailoring [crafted] | 11.1 spell_power points (1.56 DPS) | yes | Pristine Sash (253925, -0.42 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.57 DPS) [crafted]; Keller's Girdle (2911, -0.87 DPS, sim-verified) [world_drop] |
| legs | Filigreed Flame Leggings (253941) | Tailoring [crafted] | 17.7 spell_power points (2.49 DPS) | yes | Phoenix Pants (4317, -0.53 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.70 DPS) [crafted]; Darkweave Breeches (12987, -1.36 DPS) [world_drop] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 11.1 spell_power points (1.56 DPS) | yes | Flame Boots (253893, -0.28 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.71 DPS) [crafted]; Red Woolen Boots (4313, -0.89 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 7.1 spell_power points (1.00 DPS) | yes | Lavishly Jeweled Ring (1156, -0.14 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.57 DPS) [world_drop]; Furlfeather Ring (286745, -0.86 DPS) [world] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 7.0 spell_power points (0.99 DPS) | yes | Volcanic Rock Ring (12053, -0.56 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.57 DPS, sim-verified) [dungeon]; Furlfeather Ring (286745, -0.85 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | 30.8 spell_power points (4.33 DPS) | yes | Emberstone Staff (5201, -0.95 DPS) [dungeon]; Staff of the Blessed Seer (2271, -1.01 DPS, sim-verified) [dungeon]; Buzzer Blade (2169, -1.23 DPS) [dungeon] |
| off_hand | Antipodean Rod (2879) | World drop [world_drop] | 6.0 spell_power points (0.84 DPS) | yes | Tear of Grief (5611, -0.19 DPS) [quest]; Seer's Fine Stein (7608, -0.39 DPS) [world_drop]; Mathystra Channeling Orb (281291, -0.58 DPS, sim-verified) [quest] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller (12984, -0.45 DPS, sim-verified) [world_drop]; Deepblaze (279896, -1.22 DPS) [quest]; Sizzle Stick (8071, -1.22 DPS) [quest] |

**New at 20:** head: Flame Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Flame Gown; wrist: Mindthrust Bracers; hands: Flame Gloves; waist: Flame Sash; legs: Filigreed Flame Leggings; feet: Spidersilk Boots; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Evocator's Blade; off_hand: Antipodean Rod; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 198, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 8350 The 1 Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (human, 000000000000000000-54500001230100000-000000000000000000)

Set DPS (verified): 79.8. Weights run: 2.9s. Verify run: 5.4s. 326 eligible items had no known source.

Race sweep: gnome 81.2 ± 0.1, human 79.8 ± 0.1, high-order-skyborne 78.6 ± 0.1 (dps); gnome beats the guide race human by 1.3 (1.67%) ± 0.2, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.968 ± 0.015, crit=0.232 ± 0.012 per rating point (14 rating = 1%, 3.246 per %), hit=0.395 ± 0.022 per rating point (10 rating = 1%, 3.953 per %), spell_haste=0.504 ± 0.091, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (-0.072 ± 0.049), mp5=not significant (0.595 ± 0.163)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Flame Circlet (253979) | Tailoring [crafted] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Nightsky Cowl (4039, -0.60 DPS) [world_drop]; Flame Circlet (253953, -0.75 DPS) [crafted]; Enchanter's Cowl (4322, -0.85 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 5.8 spell_power points (1.45 DPS) | yes | Seal of the Expedition (271920, -0.24 DPS) [rep]; Crystal Starfire Medallion (5003, -0.48 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.48 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 17.7 spell_power points (4.42 DPS) | yes | Death Speaker Mantle (6685, -0.76 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -1.00 DPS) [world_drop]; Invoker's Mantle (215365, -1.47 DPS) [crafted] |
| back | Expeditionary's Cape (271923) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Resilient Cape (14400, -0.66 DPS, sim-verified) [world_drop]; Hillman's Cloak (3719, -0.68 DPS) [crafted]; Vine Pruner's Cloak (279835, -0.68 DPS) [quest] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Flame Gown (253965, -0.21 DPS) [crafted]; Mechbuilder's Overalls (9508, -0.25 DPS) [dungeon]; Green Silk Armor (7065, -1.34 DPS, sim-verified) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (3.25 DPS) | yes | Spidertank Oilrag (9448, -0.84 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -1.31 DPS) [world_drop]; Mindthrust Bracers (1974, -1.59 DPS) [dungeon] |
| hands | Flame Gloves (253917) | Tailoring [crafted] | 9.9 spell_power points (2.47 DPS) | yes | Phoenix Gloves (4331, -0.23 DPS) [crafted]; Truefaith Gloves (7049, -0.50 DPS) [crafted]; Town Clerk's Mittens (270029, -1.07 DPS, sim-verified) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 13.9 spell_power points (3.47 DPS) | yes | Belt of Arugal (6392, -0.50 DPS) [dungeon]; Invoker's Cord (215366, -0.52 DPS) [crafted]; Crimson Silk Belt (7055, -1.34 DPS, sim-verified) [crafted] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | 20.8 spell_power points (5.19 DPS) | yes | Filigreed Flame Leggings (253941, -1.01 DPS, sim-verified) [crafted]; Phoenix Pants (4317, -1.72 DPS) [crafted]; Pristine Leggings (253987, -1.75 DPS) [crafted] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 19.8 spell_power points (4.94 DPS) | yes | Gilded Slippers (254001, -1.64 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.76 DPS) [dungeon]; Spidersilk Boots (4320, -2.22 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 12.8 spell_power points (3.19 DPS) | yes | Lorekeeper's Ring (19525, -1.14 DPS) [rep]; Ring of Power Regulation (273030, -1.25 DPS) [dungeon]; Minor Channeling Ring (1449, -1.46 DPS) [quest] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 8.8 spell_power points (2.20 DPS) | yes | Lorekeeper's Ring (19525, -0.15 DPS) [rep]; Ring of Power Regulation (273030, -0.25 DPS) [dungeon]; Minor Channeling Ring (1449, -0.47 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.00 DPS) [dungeon]; Tendonscraper Dagger (273053, -1.00 DPS) [world]; Glimmering Staff (249392, -3.99 DPS, sim-verified) [crafted] |
| off_hand | Rod of Molten Fire (2565) | World drop [world_drop] | 13.0 spell_power points (3.25 DPS) | yes | Orb of Mystic Insight (249394, -0.69 DPS, sim-verified) [crafted]; Antipodean Rod (2879, -1.75 DPS) [world_drop]; Satyr's Rod (15962, -1.80 DPS) [world_drop] |
| ranged | Scorching Wand (5213) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Starfaller (13063, +0.00 DPS) [world_drop]; Greater Mystic Wand (217287, -0.22 DPS) [crafted]; Wand of Mana Concentration (273043, -0.92 DPS, sim-verified) [world] |

**New at 30:** head: Filigreed Flame Circlet; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Expeditionary's Cape; chest: Beguiler Robes; wrist: Phoenix Bindings; waist: Highlander's Cloth Girdle; legs: Flame Leggings; feet: Fiery Slippers; finger1: Snake Hoop; finger2: Runemender's Seal; main_hand: Swinetusk Shank; off_hand: Rod of Molten Fire; ranged: Scorching Wand

No-known-source sample (15 of 326, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (human, 000000000000000000-54500001230133031-000000000000000000)

Set DPS (verified): 161.5. Weights run: 3.0s. Verify run: 5.2s. 456 eligible items had no known source.

Race sweep: human 161.5 ± 0.3, gnome 160.6 ± 0.3, high-order-skyborne 156.2 ± 0.3 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.646 ± 0.014, crit=0.189 ± 0.011 per rating point (14 rating = 1%, 2.648 per %), hit=0.295 ± 0.019 per rating point (10 rating = 1%, 2.950 per %), spell_haste=not significant (0.233 ± 0.131), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.237 ± 0.117), mp5=0.780 ± 0.145

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Spellpower Goggles Xtreme (10502, -0.03 DPS) [crafted]; Corpseshroud (10574, -0.46 DPS) [dungeon]; Miner's Hat of the Deep (9429, -0.67 DPS, sim-verified) [dungeon] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Prodigious Shadowshard Pendant (17773, -0.17 DPS) [quest]; Glowing Eye of Mordresh (10769, -0.60 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -0.87 DPS) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.4 spell_power points (3.85 DPS) | yes | Green Silken Shoulders (7057, -0.07 DPS) [crafted]; Bloodmage Mantle (7684, -0.15 DPS) [dungeon]; Berylline Pads (4197, -0.48 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.20 DPS) [crafted]; Expeditionary's Cape (271922, -0.29 DPS) [rep] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 27.1 spell_power points (6.78 DPS) | yes | Cindercloth Robe (10042, -0.04 DPS) [crafted]; Robe of Power (7054, -0.87 DPS) [crafted]; Robe of the Magi (1716, -2.15 DPS, sim-verified) [world_drop] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (3.25 DPS) | yes | Spidertank Oilrag (9448, -0.41 DPS) [dungeon]; Arcane Runed Bracers (4744, -1.00 DPS) [quest]; Condor Bracers (15864, -1.50 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.2 spell_power points (5.55 DPS) | yes | Crimson Silk Gloves (7064, -1.09 DPS) [crafted]; Red Mageweave Gloves (10018, -1.19 DPS) [crafted]; Fiery Handwraps (254025, -2.85 DPS, sim-verified) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | 20.6 spell_power points (5.14 DPS) | yes | Highlander's Cloth Girdle (20098, -1.00 DPS) [rep]; Scorching Sash (4117, -1.36 DPS) [quest]; Deathmage Sash (10771, -2.37 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.8 spell_power points (5.43 DPS) | yes | Flame Leggings (253991, -0.51 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.34 DPS) [crafted]; Filigreed Flame Leggings (253941, -1.48 DPS) [crafted] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 18.5 spell_power points (4.61 DPS) | yes | Acidic Walkers (9454, -1.84 DPS) [dungeon]; Spidersilk Boots (4320, -2.22 DPS) [crafted]; Gilded Slippers (254001, -2.32 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.2 spell_power points (3.30 DPS) | yes | Lorekeeper's Ring (19524, -0.47 DPS) [rep]; Runemender's Seal (273031, -1.17 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 12.2 spell_power points (3.04 DPS) | yes | Lorekeeper's Ring (19524, -0.21 DPS) [rep]; Runemender's Seal (273031, -0.91 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+3.9 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Celestial Stave (9517, -7.13 DPS) [quest]; Hypnotic Blade (7714, -7.45 DPS) [dungeon]; Illusionary Rod (7713, -10.48 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Ragefire Wand (7513) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plaguerot Sprig (10766, -1.07 DPS) [dungeon]; Twisted Nether Wand (249144, -1.46 DPS) [crafted]; Jaina's Firestarter (13064, -1.56 DPS, sim-verified) [world_drop] |

**New at 40:** head: Papal Fez; neck: Triune Amulet; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Fiery Cord; legs: Red Mageweave Pants; finger1: Philanthropist's Ring; finger2: Snake Hoop; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Staff of Jordan; ranged: Ragefire Wand

No-known-source sample (15 of 456, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (human, 230002000000000000-54500001230133031-003000000000000000)

Set DPS (verified): 255.6. Weights run: 3.0s. Verify run: 6.4s. 598 eligible items had no known source.

Race sweep: human 255.6 ± 0.4, gnome 253.2 ± 0.5, high-order-skyborne 246.1 ± 0.4 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.702 ± 0.028, crit=0.294 ± 0.022 per rating point (14 rating = 1%, 4.116 per %), hit=0.489 ± 0.037 per rating point (10 rating = 1%, 4.888 per %), spell_haste=not significant (0.373 ± 0.223), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.333 ± 0.243), mp5=not significant (0.578 ± 0.288)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Flame (3075) | World drop [world_drop] | 53.3 spell_power points (13.03 DPS) | yes | Red Mageweave Headband (10033, -4.96 DPS) [crafted]; Dreamweave Circlet (10041, -5.21 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -6.43 DPS) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 20.2 spell_power points (4.94 DPS) | yes | Talvash's Enhancing Necklace (7673, -1.83 DPS) [quest]; Glowing Eye of Mordresh (10769, -1.86 DPS) [dungeon]; Mindburst Medallion (11196, -2.04 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kentic Amice (11624, -0.13 DPS) [dungeon]; Netherflame Shoulders (254053, -0.25 DPS) [crafted]; Mantle of Lost Hope (22234, -1.61 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cindercloth Cloak (14044, -0.39 DPS) [crafted]; Runecloth Cloak (13860, -1.37 DPS) [crafted]; Cloak of Fire (14134, -1.52 DPS) [crafted] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamweave Vest (10021, +0.00 DPS) [crafted]; Robe of the Magi (1716, -0.03 DPS) [world_drop]; Robes of Insight (940, -0.04 DPS) [world_drop] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 19.6 spell_power points (4.78 DPS) | yes | Aristocratic Cuffs (12546, -1.23 DPS) [dungeon]; Nethergeld Cuffs (254061, -1.47 DPS) [crafted]; Phoenix Bindings (210781, -1.61 DPS) [crafted] |
| hands | Fiery Gloves (254099) | Tailoring [crafted] | 28.7 spell_power points (7.00 DPS) | yes | Dreamweave Gloves (10019, -1.35 DPS) [crafted]; Runecloth Gloves (13863, -1.80 DPS) [crafted]; Cindercloth Gloves (14043, -1.96 DPS) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fiery Waistcord (254085, +0.00 DPS) [crafted]; Fiery Cord (254041, -0.56 DPS) [crafted]; Satyrmane Sash (17755, -0.71 DPS) [dungeon] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshock Leggings (9484, +0.00 DPS) [dungeon]; Cindercloth Leggings (12256, +0.00 DPS) [vendor]; Knight's Dreadweave Leggings (220888, +0.00 DPS) [vendor] |
| feet | Fiery Sandals (254111) | Tailoring [crafted] | 26.3 spell_power points (6.43 DPS) | yes | Cindercloth Boots (10044, -1.30 DPS) [crafted]; Fiery Slippers (254005, -1.73 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -1.74 DPS) [vendor] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 19.7 spell_power points (4.81 DPS) | yes | Lorekeeper's Ring (19523, -1.31 DPS) [rep]; Philanthropist's Ring (281635, -1.51 DPS) [quest]; Snake Hoop (6750, -1.57 DPS) [quest] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.2 spell_power points (3.72 DPS) | yes | Lorekeeper's Ring (19523, -0.23 DPS) [rep]; Philanthropist's Ring (281635, -0.42 DPS) [quest]; Snake Hoop (6750, -0.49 DPS) [quest] |
| trinket1 | Fire Ruby (20036) | Destroy Morphaz [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.95 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -2.34 DPS) [crafted]; Ankh of Life (1713, -2.44 DPS) [world_drop] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Staff of Jordan (873, +0.00 DPS) [world_drop]; Kindling Stave (11750, +0.00 DPS) [dungeon] |
| off_hand | Truesilver Conduit (249455) | Enchanting [crafted] | 15.6 spell_power points (3.81 DPS) | yes | Celestial Orb (7515, -0.12 DPS) [quest]; Orb of the Forgotten Seer (7685, -0.39 DPS) [dungeon]; Enthralled Sphere (11625, -0.56 DPS) [dungeon] |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Noxious Shooter (17745, -2.86 DPS) [dungeon]; Lesser Eternal Wand (249232, -3.27 DPS) [crafted]; Wand of Allistarj (13065, -4.51 DPS) [world_drop] |

**New at 50:** head: Eye of Flame; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Embrace of the Wind Serpent; wrist: Netherflame Cuffs; hands: Fiery Gloves; waist: Ban'thok Sash; legs: Kilt of the Atal'ai Prophet; feet: Fiery Sandals; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Fire Ruby; trinket2: Thunderbrew's Boot Flask; main_hand: Blade of Eternal Darkness; off_hand: Truesilver Conduit; ranged: Pyric Caduceus

No-known-source sample (15 of 598, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (human, 230105200100000000-55500201230133031-003000000000000000)

Set DPS (verified): 398.4. Weights run: 3.4s. Verify run: 26.5s. 1456 eligible items had no known source.

Race sweep: human 398.4 ± 0.7, gnome 394.8 ± 0.8, high-order-skyborne 385.8 ± 0.7 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.018 ± 0.073, crit=0.689 ± 0.058 per rating point (14 rating = 1%, 9.645 per %), hit=1.230 ± 0.114 per rating point (10 rating = 1%, 12.299 per %), spell_haste=not significant (2.655 ± 1.465), spell_penetration=not significant (0.000 ± 0.000), fire_power=0.850 ± 0.003, spirit=not significant (0.707 ± 0.902), mp5=not significant (0.284 ± 0.946)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 54.5 spell_power points (18.04 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Eye of Flame (3075, -0.22 DPS) [world_drop]; Lieutenant Commander's Silk Cowl (227103, -0.42 DPS) [pvp] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 31.4 spell_power points (10.39 DPS) | yes | Lady Maye's Pendant (14558, -0.35 DPS) [world_drop]; Nacreous Shell Necklace (22403, -0.65 DPS) [dungeon]; Star of Mystaria (12103, -1.10 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Theramore Shoulderpads (271953, +0.00 DPS) [rep]; Mantle of the Timbermaw (19050, -13.43 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -4.00 DPS) [dungeon]; Crystalline Threaded Cape (20697, -4.10 DPS) [world]; Shroud of the Exile (15421, -4.45 DPS, sim-verified) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Vestments (231603, +0.00 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -3.43 DPS) [pvp]; Magister's Robes (16688, -13.50 DPS, sim-verified) [dungeon] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (+11.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Silk Bracers (16438, +0.00 DPS) [pvp]; Dryad's Wrist Bindings (19596, +0.00 DPS) [rep] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Hands of Power (13253, -12.40 DPS, sim-verified) [dungeon] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Marshal's Silk Sash (16439, +0.00 DPS) [pvp]; Magician's Cord (272393, -14.15 DPS, sim-verified) [vendor] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 61.5 spell_power points (20.35 DPS) | yes | Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -3.30 DPS) [pvp]; Magister's Leggings (16687, -4.40 DPS) [dungeon] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dragonrider Boots (18102, +0.00 DPS) [dungeon]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Kayser's Boots of Precision (22231, -14.98 DPS, sim-verified) [dungeon] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Maiden's Circle (13001, -0.94 DPS) [world_drop]; Seal of Rivendare (13345, -2.28 DPS) [dungeon]; Naglering (11669, -12.48 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Seal of Rivendare (13345, -0.88 DPS) [dungeon]; Songstone of Ironforge (12543, -0.94 DPS) [quest]; Maiden's Circle (13001, -1.77 DPS, sim-verified) [world_drop] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Fire Ruby (20036, +0.00 DPS) [quest]; Second Wind (11819, -4.70 DPS, sim-verified) [dungeon] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+3.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, -1.89 DPS) [dungeon]; Fire Ruby (20036, -2.02 DPS) [quest]; Ankh of Life (1713, -3.62 DPS, sim-verified) [world_drop] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-verified (+4.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Ironbark Staff (20069, +0.00 DPS) [rep]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Blade of Eternal Darkness (17780, -4.53 DPS, sim-verified) [dungeon] |
| off_hand | Tome of Fiery Arcana (19311) | Stormpike Guard [rep] | 31.5 spell_power points (10.42 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Tome of Restoration (234590, -0.73 DPS) [pvp]; Therazane's Touch (19315, -0.81 DPS) [rep] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Oblivion's Touch (18761, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -0.25 DPS) [dungeon] |

**New at 60:** head: Spellweaver's Turban; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Sentinel's Silk Leggings; feet: Devout Sandals; finger1: Channeler's Ring; finger2: Blessed Band of Light; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Sageblade; off_hand: Tome of Fiery Arcana; ranged: Brilliant Wand

No-known-source sample (15 of 1456, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (human, 050105200100000000-55500201230133031-003000000000000000)

Set DPS (verified): 820.7. Weights run: 3.9s. Verify run: 18.7s. 1456 eligible items had no known source.

Race sweep: gnome 822.8 ± 0.9, human 820.7 ± 0.9, high-order-skyborne 800.8 ± 0.9 (dps); gnome beats the guide race human by 2.1 (0.26%) ± 1.2, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.574 ± 0.049, crit=0.954 ± 0.058 per rating point (14 rating = 1%, 13.356 per %), hit=1.554 ± 0.114 per rating point (10 rating = 1%, 15.543 per %), spell_haste=-5.870 ± 1.350, spell_penetration=not significant (0.000 ± 0.000), fire_power=0.867 ± 0.002, spirit=not significant (-0.204 ± 0.271), mp5=not significant (-0.006 ± 0.381)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 48.5 spell_power points (29.10 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -2.31 DPS) [pvp]; Eye of Flame (3075, -4.90 DPS, sim-verified) [world_drop] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 29.1 spell_power points (17.47 DPS) | yes | Star of Mystaria (12103, -2.94 DPS) [dungeon]; Lady Maye's Pendant (14558, -4.33 DPS) [world_drop]; Arcane Crystal Pendant (20037, -5.81 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 37.8 spell_power points (22.67 DPS) | yes | Lieutenant Commander's Silk Mantle (227102, -1.89 DPS) [pvp]; Field Marshal's Silk Spaulders (231602, -2.52 DPS) [pvp]; Shroud of the Nathrezim (18720, -5.03 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -8.28 DPS, sim-verified) [world]; Shroud of Arcane Mastery (22330, -8.56 DPS) [dungeon]; Mageflame Cloak (13007, -10.75 DPS) [world_drop] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Vestments (231603, -2.47 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -9.67 DPS) [pvp]; Alanna's Embrace (13314, -17.20 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 26.6 spell_power points (15.94 DPS) | yes | Runecloth Cuffs (254123, -5.91 DPS) [crafted]; Manacle Cuffs (11962, -6.58 DPS) [quest]; Netherflame Cuffs (254065, -6.77 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 41.5 spell_power points (24.85 DPS) | yes | Marshal's Silk Gloves (16440, -4.53 DPS) [vendor]; Marshal's Silk Gauntlets (231608, -4.53 DPS) [vendor]; Hands of Power (13253, -8.83 DPS, sim-verified) [dungeon] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 33.3 spell_power points (19.95 DPS) | yes | Magician's Cord (272393, -0.30 DPS) [vendor]; Highlander's Cloth Girdle (20047, -1.49 DPS) [rep]; Clutch of Andros (13956, -4.10 DPS) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 64.5 spell_power points (38.64 DPS) | yes | Marshal's Silk Leggings (231605, -5.76 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -12.19 DPS) [pvp]; Skyshroud Leggings (13170, -16.52 DPS) [dungeon] |
| feet | Kayser's Boots of Precision (22231) | Blackrock Spire: War Master Voone [dungeon] | 29.9 spell_power points (17.91 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Sergeant Major's Dreadweave Boots (220891, -0.70 DPS) [vendor]; Dragonrider Boots (18102, -1.61 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -2.72 DPS) [dungeon]; Channeler's Ring (272406, -3.19 DPS) [vendor]; Naglering (11669, -17.42 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -0.71 DPS) [dungeon]; Channeler's Ring (272406, -1.17 DPS) [vendor]; Maiden's Circle (13001, -2.11 DPS) [world_drop] |
| trinket1 | Eye of the Beast (13968) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fire Ruby (20036, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Smolderweb's Eye (13213, -10.63 DPS, sim-verified) [dungeon] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, -2.29 DPS, sim-verified) [dungeon]; Frozen Heart of the Mountain (249469, -2.40 DPS) [crafted]; Fire Ruby (20036, -3.51 DPS) [quest] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -1.44 DPS) [pvp]; Elemental Mage Staff (944, -15.47 DPS, sim-verified) [world_drop]; Staff of Jordan (873, -17.10 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Oblivion's Touch (18761, +0.00 DPS) [dungeon]; Pyric Caduceus (11748, -1.08 DPS) [dungeon] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; legs: Sentinel's Silk Leggings; feet: Kayser's Boots of Precision; finger1: Don Julio's Band; finger2: Blessed Band of Light; trinket1: Eye of the Beast; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1456, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (undead, 000000000000000000-51400001000000000-000000000000000000)

Set DPS (verified): 39.1. Weights run: 2.6s. Verify run: 3.1s. 192 eligible items had no known source.

Race sweep: undead 39.1 ± 0.1, orc 37.9 ± 0.1, troll 37.6 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.018 ± 0.016, crit=0.151 ± 0.007 per rating point (14 rating = 1%, 2.117 per %), hit=0.305 ± 0.016 per rating point (10 rating = 1%, 3.051 per %), spell_haste=not significant (0.571 ± 0.209), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.144 ± 0.082), mp5=1.059 ± 0.219

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Flame Circlet (253953) | Tailoring [crafted] | 11.0 spell_power points (1.55 DPS) | yes | Shadow Goggles (4373, -0.71 DPS) [crafted]; Pristine Circlet (253949, -0.75 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.47 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.4 spell_power points (0.06 DPS) | yes | Scholarly Pendant (277203, -0.02 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 14.2 spell_power points (1.99 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.81 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.43 DPS) [crafted]; Slime-encrusted Pads (6461, -1.55 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Sanguine Cape (14376, -0.07 DPS) [world_drop]; Pearl-clasped Cloak (5542, -0.10 DPS, sim-verified) [crafted]; Battle Healer's Cloak (20427, -0.18 DPS) [rep] |
| chest | Filigreed Flame Gown (253905) | Tailoring [crafted] | 14.5 spell_power points (2.04 DPS) | yes | Filigreed Pristine Gown (253901, -0.60 DPS, sim-verified) [crafted]; Gray Woolen Robe (2585, -0.76 DPS) [crafted]; Mystic's Wrap (14369, -1.04 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 8.3 spell_power points (1.16 DPS) | yes | Bright Bracers (3647, -0.41 DPS, sim-verified) [world_drop]; Featherbead Bracers (15452, -0.73 DPS) [quest]; Owlbeard Bracers (16981, -0.74 DPS) [quest] |
| hands | Flame Gloves (253917) | Tailoring [crafted] | 10.1 spell_power points (1.41 DPS) | yes | Phoenix Gloves (4331, -0.15 DPS) [crafted]; Pristine Gloves (253913, -0.42 DPS) [crafted]; Serpent Gloves (5970, -0.43 DPS) [dungeon] |
| waist | Flame Sash (253929) | Tailoring [crafted] | 11.1 spell_power points (1.56 DPS) | yes | Pristine Sash (253925, -0.42 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.57 DPS) [crafted]; Keller's Girdle (2911, -0.77 DPS, sim-verified) [world_drop] |
| legs | Filigreed Flame Leggings (253941) | Tailoring [crafted] | 17.7 spell_power points (2.49 DPS) | yes | Phoenix Pants (4317, -0.52 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.70 DPS) [crafted]; Darkweave Breeches (12987, -1.36 DPS) [world_drop] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 11.1 spell_power points (1.56 DPS) | yes | Flame Boots (253893, -0.30 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.71 DPS) [crafted]; Red Woolen Boots (4313, -0.89 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 7.1 spell_power points (1.00 DPS) | yes | Loop of Sacrifice (281673, -0.29 DPS) [quest]; Volcanic Rock Ring (12053, -0.57 DPS) [world_drop]; Furlfeather Ring (286745, -0.86 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 6.1 spell_power points (0.86 DPS) | yes | Loop of Sacrifice (281673, -0.14 DPS) [quest]; Volcanic Rock Ring (12053, -0.43 DPS) [world_drop]; Furlfeather Ring (286745, -0.72 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Staff of the Blessed Seer (2271, -0.18 DPS) [dungeon]; Crescent Staff (6505, -0.35 DPS, sim-verified) [quest]; Gnarled Necromancer's Staff (251534, -0.47 DPS) [quest] |
| off_hand | Antipodean Rod (2879) | World drop [world_drop] | 6.0 spell_power points (0.84 DPS) | yes | Seer's Fine Stein (7608, -0.39 DPS) [world_drop]; Spellbinder Orb (15926, -0.41 DPS) [dungeon]; Tork Wrench (11855, -0.56 DPS) [quest] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller (12984, -0.66 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -1.14 DPS) [quest]; Greater Magic Wand (11288, -1.18 DPS) [crafted] |

**New at 20:** head: Flame Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Flame Gown; wrist: Mindthrust Bracers; hands: Flame Gloves; waist: Flame Sash; legs: Filigreed Flame Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Evocator's Blade; off_hand: Antipodean Rod; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 192, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 8350 The 1 Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 23173 Abomination Skin Leggings

### Band 30 (undead, 000000000000000000-54500001230100000-000000000000000000)

Set DPS (verified): 80.4. Weights run: 2.9s. Verify run: 5.5s. 318 eligible items had no known source.

Race sweep: undead 80.4 ± 0.1, orc 78.9 ± 0.1, troll 78.0 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.968 ± 0.015, crit=0.232 ± 0.012 per rating point (14 rating = 1%, 3.246 per %), hit=0.395 ± 0.022 per rating point (10 rating = 1%, 3.953 per %), spell_haste=0.504 ± 0.091, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (-0.072 ± 0.049), mp5=not significant (0.595 ± 0.163)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Flame Circlet (253979) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Nightsky Cowl (4039, -0.60 DPS) [world_drop]; Flame Circlet (253953, -0.75 DPS) [crafted]; Enchanter's Cowl (4322, -0.77 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 5.8 spell_power points (1.45 DPS) | yes | Darkspear Warding Pendant (272075, -0.24 DPS) [vendor]; Crystal Starfire Medallion (5003, -0.48 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.48 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 17.7 spell_power points (4.42 DPS) | yes | Death Speaker Mantle (6685, -0.74 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -1.00 DPS) [world_drop]; Invoker's Mantle (215365, -1.47 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, -0.68 DPS) [crafted]; Windsong Drape (15468, -0.68 DPS) [quest]; Resilient Cape (14400, -0.80 DPS, sim-verified) [world_drop] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Flame Gown (253965, -0.21 DPS) [crafted]; Mechbuilder's Overalls (9508, -0.25 DPS) [dungeon]; Green Silk Armor (7065, -0.79 DPS, sim-verified) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (3.25 DPS) | yes | Spidertank Oilrag (9448, -0.70 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -1.31 DPS) [world_drop]; Mindthrust Bracers (1974, -1.59 DPS) [dungeon] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 10.8 spell_power points (2.71 DPS) | yes | Flame Gloves (253917, -0.23 DPS) [crafted]; Phoenix Gloves (4331, -0.46 DPS) [crafted]; Truefaith Gloves (7049, -0.73 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 13.9 spell_power points (3.47 DPS) | yes | Belt of Arugal (6392, -0.50 DPS) [dungeon]; Invoker's Cord (215366, -0.52 DPS) [crafted]; Crimson Silk Belt (7055, -1.17 DPS, sim-verified) [crafted] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | 20.8 spell_power points (5.19 DPS) | yes | Filigreed Flame Leggings (253941, -1.26 DPS, sim-verified) [crafted]; Phoenix Pants (4317, -1.72 DPS) [crafted]; Pristine Leggings (253987, -1.75 DPS) [crafted] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 19.8 spell_power points (4.94 DPS) | yes | Gilded Slippers (254001, -1.61 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.76 DPS) [dungeon]; Spidersilk Boots (4320, -2.22 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 12.8 spell_power points (3.19 DPS) | yes | Advisor's Ring (19521, -1.14 DPS) [rep]; Ring of Power Regulation (273030, -1.25 DPS) [dungeon] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 8.8 spell_power points (2.20 DPS) | yes | Ring of Power Regulation (273030, -0.25 DPS) [dungeon]; Advisor's Ring (19521, -0.58 DPS, sim-verified) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.00 DPS) [dungeon]; Tendonscraper Dagger (273053, -1.00 DPS) [world]; Glimmering Staff (249392, -4.37 DPS, sim-verified) [crafted] |
| off_hand | Rod of Molten Fire (2565) | World drop [world_drop] | 13.0 spell_power points (3.25 DPS) | yes | Orb of Mystic Insight (249394, -0.83 DPS, sim-verified) [crafted]; Witch's Finger (16887, -1.31 DPS) [quest]; Antipodean Rod (2879, -1.75 DPS) [world_drop] |
| ranged | Scorching Wand (5213) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Starfaller (13063, +0.00 DPS) [world_drop]; Greater Mystic Wand (217287, -0.22 DPS) [crafted]; Wand of Mana Concentration (273043, -1.05 DPS, sim-verified) [world] |

**New at 30:** head: Filigreed Flame Circlet; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Darkspear Raider's Cloak; chest: Beguiler Robes; wrist: Phoenix Bindings; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Flame Leggings; feet: Fiery Slippers; finger1: Snake Hoop; finger2: Runemender's Seal; main_hand: Swinetusk Shank; off_hand: Rod of Molten Fire; ranged: Scorching Wand

No-known-source sample (15 of 318, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 8350 The 1 Ring; 9362 Brilliant Gold Ring

### Band 40 (undead, 000000000000000000-54500001230133031-000000000000000000)

Set DPS (verified): 159.7. Weights run: 3.0s. Verify run: 5.0s. 448 eligible items had no known source.

Race sweep: undead 159.7 ± 0.3, orc 158.2 ± 0.3, troll 156.0 ± 0.3 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.557 ± 0.017, crit=0.328 ± 0.019 per rating point (14 rating = 1%, 4.593 per %), hit=0.549 ± 0.031 per rating point (10 rating = 1%, 5.494 per %), spell_haste=0.466 ± 0.103, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=0.971 ± 0.137, mp5=1.154 ± 0.163

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 26.2 spell_power points (9.84 DPS) | yes | Miner's Hat of the Deep (9429, -1.78 DPS) [dungeon]; Electromagnetic Gigaflux Reactivator (9492, -1.88 DPS) [dungeon]; Whitemane's Chapeau (7720, -2.73 DPS, sim-verified) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Necklace of Calisea (1714, -0.29 DPS) [dungeon]; Amberglow Talisman (10824, -0.65 DPS) [quest]; Triune Amulet (7722, -1.10 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Green Silken Shoulders (7057, -0.04 DPS) [crafted]; Bloodmage Mantle (7684, -0.09 DPS) [dungeon]; Mantle of Doan (7712, -1.37 DPS, sim-verified) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battle Healer's Cloak (19528, -0.74 DPS) [rep]; Glowing Thresher Cape (6901, -0.77 DPS) [dungeon]; Cloak of Blight (6832, -2.22 DPS, sim-verified) [quest] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 36.6 spell_power points (13.74 DPS) | yes | Robe of the Magi (1716, -2.41 DPS) [world_drop]; Robe of Power (7054, -3.06 DPS) [crafted]; Cindercloth Robe (10042, -3.61 DPS) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (4.88 DPS) | yes | Spidertank Oilrag (9448, -0.20 DPS) [dungeon]; Earthen Silk Cuffs (254019, -1.56 DPS) [crafted]; Radiant Silver Bracers (4545, -1.71 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 27.0 spell_power points (10.14 DPS) | yes | Fiery Handwraps (254025, -1.11 DPS, sim-verified) [crafted]; Crimson Silk Gloves (7064, -3.63 DPS) [crafted]; Red Mageweave Gloves (10018, -3.92 DPS) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | 24.3 spell_power points (9.11 DPS) | yes | Gilded Cord (254037, -2.25 DPS) [crafted]; Star Belt (4329, -2.78 DPS) [crafted]; Scorching Sash (4117, -3.57 DPS, sim-verified) [quest] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Black Mageweave Leggings (9999, -0.43 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.76 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.80 DPS, sim-verified) [dungeon] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 20.8 spell_power points (7.80 DPS) | yes | Thoughtcast Boots (10578, -1.19 DPS) [dungeon]; Gilded Slippers (254001, -2.25 DPS) [crafted]; Furen's Boots (13100, -4.40 DPS, sim-verified) [world_drop] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 16.7 spell_power points (6.27 DPS) | yes | Advisor's Ring (19520, -1.59 DPS) [rep]; Electrocutioner Lagnut (9447, -2.59 DPS) [dungeon]; Runemender's Seal (273031, -5.45 DPS, sim-verified) [world] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Advisor's Ring (19520, -0.12 DPS) [rep]; Electrocutioner Lagnut (9447, -1.12 DPS) [dungeon]; Runemender's Seal (273031, -1.76 DPS, sim-verified) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -12.47 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -14.16 DPS) [dungeon]; Celestial Stave (9517, -14.50 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ragefire Wand (7513) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jaina's Firestarter (13064, -1.51 DPS) [world_drop]; Twisted Nether Wand (249144, -1.97 DPS) [crafted]; Plaguerot Sprig (10766, -2.88 DPS, sim-verified) [dungeon] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Icy Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Fiery Cord; finger2: Philanthropist's Ring; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Staff of Jordan; ranged: Ragefire Wand

No-known-source sample (15 of 448, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (undead, 230002000000000000-54500001230133031-003000000000000000)

Set DPS (verified): 250.9. Weights run: 3.0s. Verify run: 6.7s. 590 eligible items had no known source.

Race sweep: undead 250.9 ± 0.4, orc 247.6 ± 0.4, troll 244.7 ± 0.4 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.743 ± 0.031, crit=0.559 ± 0.032 per rating point (14 rating = 1%, 7.819 per %), hit=0.794 ± 0.053 per rating point (10 rating = 1%, 7.938 per %), spell_haste=not significant (0.146 ± 0.284), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.989 ± 0.321), mp5=not significant (1.116 ± 0.366)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Flame (3075) | World drop [world_drop] | 60.3 spell_power points (21.21 DPS) | yes | Dreamweave Circlet (10041, -8.08 DPS, sim-verified) [crafted]; Gemburst Circlet (10751, -9.10 DPS) [quest]; Red Mageweave Headband (10033, -9.30 DPS) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 20.5 spell_power points (7.19 DPS) | yes | Triune Amulet (7722, -1.95 DPS) [dungeon]; Lei of Lilies (1315, -1.98 DPS) [world_drop]; Mindburst Medallion (11196, -2.86 DPS, sim-verified) [quest] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 29.6 spell_power points (10.41 DPS) | yes | Rotgrip Mantle (17732, -1.13 DPS) [dungeon]; Flameseer Mantle (11310, -2.01 DPS) [dungeon]; Netherflame Shoulders (254053, -2.83 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cindercloth Cloak (14044, -1.91 DPS) [crafted]; Deep Woodlands Cloak (19121, -2.00 DPS) [quest]; Featherskin Cape (10843, -2.31 DPS) [world] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 48.1 spell_power points (16.91 DPS) | yes | Dreamweave Vest (10021, -3.36 DPS) [crafted]; Ghostweave Vest (14141, -3.65 DPS) [crafted]; Cindercloth Vest (14042, -4.65 DPS) [crafted] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 23.1 spell_power points (8.14 DPS) | yes | Nethergeld Cuffs (254061, -2.26 DPS, sim-verified) [crafted]; Aristocratic Cuffs (12546, -2.81 DPS) [dungeon]; Bloodband Bracers (11469, -2.99 DPS) [quest] |
| hands | Fiery Gloves (254099) | Tailoring [crafted] | 34.3 spell_power points (12.08 DPS) | yes | Cindercloth Gloves (14043, -2.27 DPS) [crafted]; Runecloth Gloves (13863, -2.38 DPS) [crafted]; Dreamweave Gloves (10019, -4.14 DPS, sim-verified) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 27.4 spell_power points (9.63 DPS) | yes | Fiery Waistcord (254085, -0.59 DPS) [crafted]; Ghostweave Cord (254073, -0.88 DPS) [crafted]; Fiery Cord (254041, -3.46 DPS, sim-verified) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 35.0 spell_power points (12.31 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -0.29 DPS) [dungeon]; Cindercloth Leggings (12256, -0.41 DPS) [vendor]; Rainstrider Leggings (11123, -1.16 DPS) [quest] |
| feet | Fiery Sandals (254111) | Tailoring [crafted] | sim-verified (+3.6 DPS vs the runner-up, not corroborated against the finished set) | yes | First Sergeant's Dreadweave Boots (220909, -1.43 DPS) [vendor]; Vinerot Sandals (17748, -1.45 DPS) [dungeon]; Mistwalker Boots (10629, -3.60 DPS, sim-verified) [dungeon] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 23.4 spell_power points (8.22 DPS) | yes | Snake Hoop (6750, -1.85 DPS) [quest]; Advisor's Ring (19519, -2.43 DPS) [rep]; Eye of Adaegus (5266, -2.48 DPS) [world_drop] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 18.2 spell_power points (6.39 DPS) | yes | Snake Hoop (6750, -0.01 DPS) [quest]; Advisor's Ring (19519, -0.60 DPS) [rep]; Eye of Adaegus (5266, -0.65 DPS) [world_drop] |
| trinket1 | Fire Ruby (20036) | Destroy Morphaz [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -0.75 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.45 DPS) [quest]; Uther's Strength (11302, -3.02 DPS, sim-verified) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Uther's Strength (11302, -0.40 DPS) [world_drop] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Spire of Hakkar (10844, +0.00 DPS) [world]; Staff of Jordan (873, -11.60 DPS, sim-verified) [world_drop] |
| off_hand | Truesilver Conduit (249455) | Enchanting [crafted] | 15.9 spell_power points (5.61 DPS) | yes | Enthralled Sphere (11625, -0.10 DPS) [dungeon]; Celestial Orb (7515, -0.25 DPS) [quest]; Flask of Forest Mojo (19115, -0.35 DPS) [quest] |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Noxious Shooter (17745, -3.82 DPS, sim-verified) [dungeon]; Lesser Eternal Wand (249232, -4.21 DPS) [crafted]; Ragefire Wand (7513, -4.87 DPS) [quest] |

**New at 50:** head: Eye of Flame; neck: Arcane Crystal Pendant; shoulder: Kentic Amice; back: Spritecaster Cape; chest: Embrace of the Wind Serpent; wrist: Netherflame Cuffs; hands: Fiery Gloves; waist: Ban'thok Sash; legs: Senior Designer's Pantaloons; feet: Fiery Sandals; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Fire Ruby; trinket2: Frozen Heart of the Mountain; main_hand: Blade of Eternal Darkness; off_hand: Truesilver Conduit; ranged: Pyric Caduceus

No-known-source sample (15 of 590, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (undead, 230105200100000000-55500201230133031-003000000000000000)

Set DPS (verified): 387.7. Weights run: 3.4s. Verify run: 24.1s. 1450 eligible items had no known source.

Race sweep: undead 387.7 ± 0.7, orc 381.5 ± 0.7, troll 378.6 ± 0.7 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.575 ± 0.054, crit=0.492 ± 0.038 per rating point (14 rating = 1%, 6.887 per %), hit=1.173 ± 0.072 per rating point (10 rating = 1%, 11.729 per %), spell_haste=not significant (-2.898 ± 0.936), spell_penetration=not significant (0.000 ± 0.000), fire_power=0.846 ± 0.003, spirit=not significant (0.556 ± 0.647), mp5=not significant (0.504 ± 0.655)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 52.0 spell_power points (11.62 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -2.32 DPS) [pvp]; Eye of Flame (3075, -3.13 DPS, sim-verified) [world_drop] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 25.3 spell_power points (5.66 DPS) | yes | Star of Mystaria (12103, -0.48 DPS) [dungeon]; Lady Maye's Pendant (14558, -0.76 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -0.77 DPS) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Silk Mantle (227104, +0.00 DPS) [pvp]; Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -13.10 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -2.24 DPS) [world]; Shroud of the Exile (15421, -2.39 DPS) [quest]; Spritecaster Cape (11623, -2.58 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Raiment (231596, -0.18 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -2.99 DPS) [pvp]; The Postmaster's Tunic (13388, -15.57 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 30.5 spell_power points (6.82 DPS) | yes | Netherflame Cuffs (254065, -2.84 DPS) [crafted]; Runecloth Cuffs (254123, -3.07 DPS) [crafted]; General's Silk Cuffs (16538, -3.13 DPS) [pvp] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Hands of Power (13253, -14.12 DPS, sim-verified) [dungeon] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Thuzadin Sash (18740, +0.00 DPS) [dungeon]; Magician's Cord (272393, -13.84 DPS, sim-verified) [vendor] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Outrider's Silk Leggings (22747, +0.00 DPS) [rep]; General's Silk Trousers (231595, +0.00 DPS) [pvp]; The Postmaster's Trousers (13389, -16.82 DPS, sim-verified) [dungeon] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fire Striders (13369, +0.00 DPS) [dungeon]; General's Silk Boots (231597, +0.00 DPS) [pvp]; The Postmaster's Treads (13391, -18.20 DPS, sim-verified) [dungeon] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Maiden's Circle (13001, -0.18 DPS) [world_drop]; Eye of Orgrimmar (12545, -0.93 DPS) [quest]; The Postmaster's Seal (13392, -2.18 DPS, sim-verified) [dungeon] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Eye of Orgrimmar (12545, -0.49 DPS) [quest]; The Postmaster's Seal (13392, -0.63 DPS) [dungeon]; Maiden's Circle (13001, -1.53 DPS, sim-verified) [world_drop] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fire Ruby (20036, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Second Wind (11819, -4.64 DPS, sim-verified) [dungeon] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Fire Ruby (20036, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-verified (+3.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Ironbark Staff (20220, +0.00 DPS) [rep]; High Warlord's War Staff (234549, -0.27 DPS) [pvp]; Blade of Eternal Darkness (17780, -3.17 DPS, sim-verified) [dungeon] |
| off_hand | Tome of Fiery Arcana (19311) | Frostwolf Clan [rep] | 31.3 spell_power points (7.00 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; Therazane's Touch (19315, -0.52 DPS) [rep]; High Warlord's Tome of Mending (234564, -1.85 DPS) [pvp] |
| ranged | Oblivion's Touch (18761) | Scholomance: Death Knight Darkreaver [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, +0.00 DPS) [dungeon]; Brilliant Wand (249385, -1.03 DPS) [crafted] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Devout Mantle; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Devout Sandals; finger1: Blessed Band of Light; finger2: Channeler's Ring; trinket1: Burst of Knowledge; trinket2: Ankh of Life; main_hand: Sageblade; off_hand: Tome of Fiery Arcana; ranged: Oblivion's Touch

No-known-source sample (15 of 1450, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60, raid preset (undead, 050105200100000000-55500201230133031-003000000000000000)

Set DPS (verified): 808.4. Weights run: 3.9s. Verify run: 18.1s. 1450 eligible items had no known source.

Race sweep: undead 808.4 ± 0.8, orc 803.0 ± 0.8, troll 797.3 ± 0.8 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.574 ± 0.049, crit=0.954 ± 0.058 per rating point (14 rating = 1%, 13.356 per %), hit=1.554 ± 0.114 per rating point (10 rating = 1%, 15.543 per %), spell_haste=-5.870 ± 1.350, spell_penetration=not significant (0.000 ± 0.000), fire_power=0.867 ± 0.002, spirit=not significant (-0.204 ± 0.271), mp5=not significant (-0.006 ± 0.381)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 48.5 spell_power points (29.10 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -2.31 DPS) [pvp]; Eye of Flame (3075, -6.30 DPS, sim-verified) [world_drop] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 29.1 spell_power points (17.47 DPS) | yes | Star of Mystaria (12103, -2.94 DPS) [dungeon]; Lady Maye's Pendant (14558, -4.33 DPS) [world_drop]; Arcane Crystal Pendant (20037, -5.81 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 37.8 spell_power points (22.67 DPS) | yes | Champion's Silk Mantle (227104, -1.89 DPS) [pvp]; Warlord's Silk Amice (231594, -2.52 DPS) [pvp]; Shroud of the Nathrezim (18720, -4.43 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of Arcane Mastery (22330, -8.56 DPS) [dungeon]; Crystalline Threaded Cape (20697, -9.07 DPS, sim-verified) [world]; Mageflame Cloak (13007, -10.75 DPS) [world_drop] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Raiment (231596, -2.47 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -9.67 DPS) [pvp]; Alanna's Embrace (13314, -22.71 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 26.6 spell_power points (15.94 DPS) | yes | Runecloth Cuffs (254123, -5.91 DPS) [crafted]; Manacle Cuffs (11962, -6.58 DPS) [quest]; Netherflame Cuffs (254065, -6.77 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 41.5 spell_power points (24.85 DPS) | yes | General's Silk Handguards (16540, -4.53 DPS) [vendor]; General's Silk Gauntlets (231599, -4.53 DPS) [vendor]; Hands of Power (13253, -4.73 DPS, sim-verified) [dungeon] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 33.3 spell_power points (19.95 DPS) | yes | Defiler's Cloth Girdle (20163, -1.49 DPS) [rep]; Clutch of Andros (13956, -4.10 DPS) [dungeon]; Magician's Cord (272393, -4.98 DPS, sim-verified) [vendor] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 38.9 spell_power points (23.32 DPS) | yes | Legionnaire's Silk Legguards (227107, +0.00 DPS) [pvp]; General's Silk Trousers (231595, +0.00 DPS) [pvp]; Skyshroud Leggings (13170, -1.21 DPS) [dungeon] |
| feet | Kayser's Boots of Precision (22231) | Blackrock Spire: War Master Voone [dungeon] | 29.9 spell_power points (17.91 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; First Sergeant's Dreadweave Boots (220909, -0.70 DPS) [vendor]; Dragonrider Boots (18102, -1.61 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -2.72 DPS) [dungeon]; Channeler's Ring (272406, -3.19 DPS) [vendor]; Naglering (11669, -18.58 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -1.17 DPS) [vendor]; Maiden's Circle (13001, -2.11 DPS) [world_drop]; Don Mauricio's Band of Domination (22433, -6.21 DPS, sim-verified) [dungeon] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -2.40 DPS) [crafted]; Fire Ruby (20036, -3.51 DPS) [quest]; Blackhand's Breadth (13965, -4.91 DPS, sim-verified) [quest] |
| trinket2 | Smolderweb's Eye (13213) | Blackrock Spire: Mother Smolderweb [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -0.93 DPS) [crafted]; Fire Ruby (20036, -2.04 DPS) [quest]; Blackhand's Breadth (13965, -6.78 DPS, sim-verified) [quest] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -1.44 DPS) [pvp]; Elemental Mage Staff (944, -13.00 DPS, sim-verified) [world_drop]; Staff of Jordan (873, -17.10 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bonecreeper Stylus (13938) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Oblivion's Touch (18761, -0.51 DPS) [dungeon]; Brilliant Wand (249385, -0.53 DPS) [crafted]; Pyric Caduceus (11748, -1.61 DPS) [dungeon] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; legs: Outrider's Silk Leggings; feet: Kayser's Boots of Precision; finger1: Don Julio's Band; finger2: Blessed Band of Light; trinket1: Weakness Analyzer; trinket2: Smolderweb's Eye; main_hand: Ironbark Staff; ranged: Bonecreeper Stylus

No-known-source sample (15 of 1450, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

