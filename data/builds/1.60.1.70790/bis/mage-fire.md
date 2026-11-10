# Leveling BiS: Fire

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 000000000000000000-51400001000000000-000000000000000000)

Set DPS (verified): 38.6. Weights run: 2.1s. Verify run: 2.1s. 198 eligible items had no known source.

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

Set DPS (verified): 79.8. Weights run: 2.5s. Verify run: 4.5s. 326 eligible items had no known source.

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

Set DPS (verified): 164.0. Weights run: 2.6s. Verify run: 3.5s. 456 eligible items had no known source.

Race sweep: gnome 165.0 ± 0.3, human 164.0 ± 0.3, high-order-skyborne 160.6 ± 0.3 (dps); gnome beats the guide race human by 1.0 (0.60%) ± 0.4, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.605 ± 0.017, crit=0.339 ± 0.019 per rating point (14 rating = 1%, 4.746 per %), hit=0.529 ± 0.031 per rating point (10 rating = 1%, 5.286 per %), spell_haste=0.576 ± 0.113, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=0.824 ± 0.140, mp5=0.978 ± 0.176

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 25.7 spell_power points (9.62 DPS) | yes | Miner's Hat of the Deep (9429, -1.28 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -1.76 DPS) [crafted]; Whitemane's Chapeau (7720, -2.37 DPS, sim-verified) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Necklace of Calisea (1714, -0.69 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.06 DPS) [quest]; Triune Amulet (7722, -2.38 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Green Silken Shoulders (7057, -0.08 DPS) [crafted]; Bloodmage Mantle (7684, -0.16 DPS) [dungeon]; Mantle of Doan (7712, -2.27 DPS, sim-verified) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Caretaker's Cape (19532, -0.68 DPS) [rep]; Expeditionary's Cape (271922, -0.75 DPS) [rep]; Glowing Thresher Cape (6901, -0.88 DPS) [dungeon] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 35.0 spell_power points (13.09 DPS) | yes | Robe of the Magi (1716, -1.96 DPS) [world_drop]; Robe of Power (7054, -2.67 DPS) [crafted]; Cindercloth Robe (10042, -2.99 DPS) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (4.87 DPS) | yes | Spidertank Oilrag (9448, -1.02 DPS, sim-verified) [dungeon]; Arcane Runed Bracers (4744, -1.50 DPS) [quest]; Earthen Silk Cuffs (254019, -1.83 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 26.2 spell_power points (9.80 DPS) | yes | Fiery Handwraps (254025, -2.98 DPS) [crafted]; Crimson Silk Gloves (7064, -3.20 DPS) [crafted]; Red Mageweave Gloves (10018, -3.42 DPS) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | 23.8 spell_power points (8.90 DPS) | yes | Gilded Cord (254037, -2.25 DPS) [crafted]; Scorching Sash (4117, -2.72 DPS, sim-verified) [quest]; Highlander's Cloth Girdle (20098, -2.76 DPS) [rep] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | 22.4 spell_power points (8.37 DPS) | yes | Stoneweaver Leggings (9407, -0.23 DPS) [dungeon]; Red Mageweave Pants (10009, -0.41 DPS) [crafted]; Black Mageweave Leggings (9999, -1.05 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 29.8 spell_power points (11.14 DPS) | yes | Fiery Slippers (254005, -1.29 DPS, sim-verified) [crafted]; Furen's Boots (13100, -4.62 DPS) [world_drop]; Thoughtcast Boots (10578, -5.26 DPS) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 16.0 spell_power points (5.99 DPS) | yes | Runemender's Seal (273031, -1.35 DPS) [world]; Lorekeeper's Ring (19524, -1.52 DPS) [rep] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.0 spell_power points (4.88 DPS) | yes | Runemender's Seal (273031, -0.23 DPS) [world]; Lorekeeper's Ring (19524, -0.41 DPS) [rep] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -13.08 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -13.83 DPS) [dungeon]; Celestial Stave (9517, -14.01 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ragefire Wand (7513) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jaina's Firestarter (13064, -1.60 DPS) [world_drop]; Plaguerot Sprig (10766, -1.64 DPS) [dungeon]; Twisted Nether Wand (249144, -2.04 DPS) [crafted] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Icy Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Fiery Cord; feet: Earthen Silk Slippers; finger2: Philanthropist's Ring; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Staff of Jordan; ranged: Ragefire Wand

No-known-source sample (15 of 456, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (human, 230002000000000000-54500001230133031-003000000000000000)

Set DPS (verified): 248.8. Weights run: 2.5s. Verify run: 6.1s. 598 eligible items had no known source.

Race sweep: human 248.8 ± 0.4, gnome 246.5 ± 0.4, high-order-skyborne 240.0 ± 0.4 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.846 ± 0.031, crit=0.394 ± 0.025 per rating point (14 rating = 1%, 5.511 per %), hit=0.684 ± 0.041 per rating point (10 rating = 1%, 6.836 per %), spell_haste=not significant (0.198 ± 0.244), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.740 ± 0.305), mp5=not significant (0.864 ± 0.323)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Flame (3075) | World drop [world_drop] | 58.9 spell_power points (12.24 DPS) | yes | Red Mageweave Headband (10033, -4.77 DPS) [crafted]; Gemburst Circlet (10751, -5.65 DPS) [quest]; Dreamweave Circlet (10041, -7.87 DPS, sim-verified) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 21.1 spell_power points (4.38 DPS) | yes | Mindburst Medallion (11196, -1.31 DPS) [quest]; Glowing Eye of Mordresh (10769, -1.52 DPS) [dungeon]; Horizon Choker (13085, -1.52 DPS, sim-verified) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Netherflame Shoulders (254053, -0.00 DPS) [crafted]; Kentic Amice (11624, -0.75 DPS, sim-verified) [dungeon]; Mantle of Lost Hope (22234, -1.34 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cindercloth Cloak (14044, -0.78 DPS) [crafted]; Runecloth Cloak (13860, -1.61 DPS) [crafted]; Expeditionary's Cape (271921, -1.66 DPS) [rep] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 36.6 spell_power points (7.61 DPS) | yes | Dreamweave Vest (10021, -0.13 DPS) [crafted]; Cindercloth Vest (14042, -0.93 DPS) [crafted]; Robes of Insight (940, -1.10 DPS) [world_drop] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 22.6 spell_power points (4.71 DPS) | yes | Aristocratic Cuffs (12546, -0.76 DPS, sim-verified) [dungeon]; Nethergeld Cuffs (254061, -1.25 DPS) [crafted]; Bloodband Bracers (11469, -1.62 DPS) [quest] |
| hands | Fiery Gloves (254099) | Tailoring [crafted] | 33.4 spell_power points (6.94 DPS) | yes | Runecloth Gloves (13863, -1.48 DPS) [crafted]; Cindercloth Gloves (14043, -1.71 DPS) [crafted]; Dreamweave Gloves (10019, -2.59 DPS, sim-verified) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 27.3 spell_power points (5.68 DPS) | yes | Fiery Waistcord (254085, -0.14 DPS) [crafted]; Fiery Cord (254041, -0.43 DPS) [crafted]; Satyrmane Sash (17755, -1.01 DPS) [dungeon] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 31.6 spell_power points (6.56 DPS) | yes | Spellshock Leggings (9484, -0.02 DPS) [dungeon]; Senior Designer's Pantaloons (11841, -0.06 DPS) [dungeon]; Cindercloth Leggings (12256, -0.30 DPS) [vendor] |
| feet | Fiery Sandals (254111) | Tailoring [crafted] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Earthen Silk Slippers (254013, -0.89 DPS, sim-verified) [crafted]; Sergeant Major's Dreadweave Boots (220891, -1.07 DPS) [vendor]; Mistwalker Boots (10629, -1.10 DPS) [dungeon] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 23.2 spell_power points (4.82 DPS) | yes | Snake Hoop (6750, -1.26 DPS) [quest]; Lorekeeper's Ring (19523, -1.60 DPS) [rep]; Woodseed Hoop (17768, -1.81 DPS) [quest] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 17.9 spell_power points (3.72 DPS) | yes | Snake Hoop (6750, -0.16 DPS) [quest]; Lorekeeper's Ring (19523, -0.50 DPS) [rep]; Woodseed Hoop (17768, -0.71 DPS) [quest] |
| trinket1 | Fire Ruby (20036) | Destroy Morphaz [quest] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, -1.06 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.37 DPS) [quest]; Frozen Heart of the Mountain (249469, -3.40 DPS, sim-verified) [crafted] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (-2.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Staff of Jordan (873, -5.20 DPS, sim-verified) [world_drop] |
| off_hand | Truesilver Conduit (249455) | Enchanting [crafted] | 16.8 spell_power points (3.49 DPS) | yes | Celestial Orb (7515, -0.26 DPS) [quest]; Orb of Lorica (11262, -0.48 DPS) [quest]; Enthralled Sphere (11625, -0.87 DPS, sim-verified) [dungeon] |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lesser Eternal Wand (249232, -3.54 DPS) [crafted]; Wand of Allistarj (13065, -4.40 DPS) [world_drop]; Noxious Shooter (17745, -5.66 DPS, sim-verified) [dungeon] |

**New at 50:** head: Eye of Flame; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Embrace of the Wind Serpent; wrist: Netherflame Cuffs; hands: Fiery Gloves; waist: Ban'thok Sash; legs: Kilt of the Atal'ai Prophet; feet: Fiery Sandals; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Fire Ruby; trinket2: Thunderbrew's Boot Flask; main_hand: Blade of Eternal Darkness; off_hand: Truesilver Conduit; ranged: Pyric Caduceus

No-known-source sample (15 of 598, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (human, 230105200100000000-55500201230133031-003000000000000000)

Set DPS (verified): 399.2. Weights run: 2.9s. Verify run: 24.0s. 1442 eligible items had no known source.

Race sweep: human 399.2 ± 0.7, gnome 396.3 ± 0.7, high-order-skyborne 387.7 ± 0.7 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.599 ± 0.052, crit=0.463 ± 0.037 per rating point (14 rating = 1%, 6.476 per %), hit=1.239 ± 0.073 per rating point (10 rating = 1%, 12.390 per %), spell_haste=not significant (-2.546 ± 0.966), spell_penetration=not significant (0.000 ± 0.000), fire_power=0.847 ± 0.003, spirit=not significant (0.338 ± 0.669), mp5=not significant (-0.531 ± 0.679)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Flame (3075, +0.00 DPS) [world_drop]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; The Postmaster's Band (13390, -9.15 DPS, sim-verified) [dungeon] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 26.2 spell_power points (5.96 DPS) | yes | Star of Mystaria (12103, -0.80 DPS) [dungeon]; Lady Maye's Pendant (14558, -0.87 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -1.37 DPS) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burial Shawl (18681, +0.00 DPS) [dungeon]; Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -11.47 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -2.46 DPS) [world]; Shroud of the Exile (15421, -2.92 DPS) [quest]; Spritecaster Cape (11623, -3.09 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Vestments (231603, -0.45 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -3.26 DPS) [pvp]; Alanna's Embrace (13314, -12.95 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 29.2 spell_power points (6.64 DPS) | yes | Runecloth Cuffs (254123, -2.77 DPS) [crafted]; Netherflame Cuffs (254065, -2.79 DPS) [crafted]; Manacle Cuffs (11962, -3.00 DPS) [quest] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 31.1 spell_power points (7.08 DPS) | yes | Gloves of Spell Mastery (14146, +0.00 DPS) [crafted]; Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Highlander's Cloth Girdle (20047, +0.00 DPS) [rep]; Magician's Cord (272393, -9.68 DPS, sim-verified) [vendor] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Sentinel's Silk Leggings (237815, +0.00 DPS) [vendor]; The Postmaster's Trousers (13389, -12.41 DPS, sim-verified) [dungeon] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Omnicast Boots (11822, +0.00 DPS) [dungeon]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -10.87 DPS, sim-verified) [dungeon] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Maiden's Circle (13001, -0.20 DPS) [world_drop]; Songstone of Ironforge (12543, -0.66 DPS) [quest]; The Postmaster's Seal (13392, -1.80 DPS, sim-verified) [dungeon] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Songstone of Ironforge (12543, -0.29 DPS) [quest]; Flaming Band (12926, -0.72 DPS) [dungeon]; Maiden's Circle (13001, -1.90 DPS, sim-verified) [world_drop] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Second Wind (11819, -5.58 DPS, sim-verified) [dungeon] |
| trinket2 | Eye of the Beast (13968) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Second Wind (11819, +0.00 DPS) [dungeon]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironbark Staff (20069, -0.01 DPS) [rep]; Grand Marshal's Stave (234571, -0.51 DPS) [pvp]; Blade of Eternal Darkness (17780, -1.53 DPS, sim-verified) [dungeon] |
| off_hand | Tome of Fiery Arcana (19311) | Stormpike Guard [rep] | 31.3 spell_power points (7.14 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Therazane's Touch (19315, -0.53 DPS) [rep]; Grand Marshal's Tome of Restoration (234590, -1.81 DPS) [pvp] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, +0.00 DPS) [dungeon]; Oblivion's Touch (18761, -0.49 DPS) [dungeon] |

**New at 60:** head: Devout Crown; neck: Nacreous Shell Necklace; shoulder: Devout Mantle; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Devout Belt; legs: Devout Skirt; feet: Devout Sandals; finger1: Blessed Band of Light; finger2: Channeler's Ring; trinket1: Burst of Knowledge; trinket2: Eye of the Beast; main_hand: Sageblade; off_hand: Tome of Fiery Arcana; ranged: Brilliant Wand

No-known-source sample (15 of 1442, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (human, 050105200100000000-55500201230133031-003000000000000000)

Set DPS (verified): 820.3. Weights run: 3.1s. Verify run: 19.6s. 1442 eligible items had no known source.

Race sweep: gnome 823.4 ± 0.9, human 820.3 ± 0.9, high-order-skyborne 801.7 ± 0.9 (dps); gnome beats the guide race human by 3.1 (0.37%) ± 1.3, beyond error.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.567 ± 0.046, crit=0.888 ± 0.060 per rating point (14 rating = 1%, 12.438 per %), hit=1.499 ± 0.120 per rating point (10 rating = 1%, 14.994 per %), spell_haste=-5.974 ± 1.389, spell_penetration=not significant (0.000 ± 0.000), fire_power=0.865 ± 0.002, spirit=not significant (0.431 ± 0.208), mp5=not significant (0.515 ± 0.372)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Flame (3075) | World drop [world_drop] | 47.2 spell_power points (27.39 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -0.56 DPS) [pvp]; Crimson Felt Hat (18727, -4.51 DPS, sim-verified) [dungeon] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 28.5 spell_power points (16.56 DPS) | yes | Star of Mystaria (12103, -1.19 DPS) [dungeon]; Lady Maye's Pendant (14558, -3.92 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -4.65 DPS) [world_drop] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 36.8 spell_power points (21.36 DPS) | yes | Field Marshal's Silk Spaulders (231602, -0.67 DPS) [pvp]; Lieutenant Commander's Silk Mantle (227102, -0.82 DPS) [pvp]; Shroud of the Nathrezim (18720, -5.85 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -7.68 DPS, sim-verified) [world]; Shroud of Arcane Mastery (22330, -8.30 DPS) [dungeon]; Shroud of the Exile (15421, -8.61 DPS) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Vestments (231603, -0.92 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -8.13 DPS) [pvp]; Alanna's Embrace (13314, -18.48 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 29.6 spell_power points (17.15 DPS) | yes | Netherflame Cuffs (254065, -7.07 DPS) [crafted]; Runecloth Cuffs (254123, -7.48 DPS) [crafted]; Manacle Cuffs (11962, -8.15 DPS) [quest] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 39.5 spell_power points (22.95 DPS) | yes | Marshal's Silk Gloves (16440, -2.08 DPS) [vendor]; Marshal's Silk Gauntlets (231608, -2.08 DPS) [vendor]; Hands of Power (13253, -8.22 DPS, sim-verified) [dungeon] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 32.7 spell_power points (18.96 DPS) | yes | Magician's Cord (272393, -0.03 DPS) [vendor]; Highlander's Cloth Girdle (20047, -1.64 DPS) [rep]; Clutch of Andros (13956, -4.00 DPS) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 62.5 spell_power points (36.30 DPS) | yes | Marshal's Silk Leggings (231605, -2.59 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -10.05 DPS) [pvp]; Magister's Leggings (16687, -14.22 DPS) [dungeon] |
| feet | Kayser's Boots of Precision (22231) | Blackrock Spire: War Master Voone [dungeon] | 29.2 spell_power points (16.95 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Sergeant Major's Dreadweave Boots (220891, -0.64 DPS) [vendor]; Fire Striders (13369, -1.13 DPS) [dungeon] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -1.20 DPS) [vendor]; Maiden's Circle (13001, -2.26 DPS) [world_drop]; Naglering (11669, -15.54 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Maiden's Circle (13001, -1.67 DPS) [world_drop]; Don Mauricio's Band of Domination (22433, -2.32 DPS) [dungeon]; Channeler's Ring (272406, -5.15 DPS, sim-verified) [vendor] |
| trinket1 | Eye of the Beast (13968) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Fire Ruby (20036, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+5.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, -1.74 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -2.61 DPS) [crafted]; Thunderbrew's Boot Flask (744, -5.08 DPS, sim-verified) [quest] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -1.43 DPS) [pvp]; Headmaster's Charge (13937, -13.21 DPS) [dungeon]; Elemental Mage Staff (944, -13.69 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Pyric Caduceus (11748, -0.30 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -1.11 DPS) [dungeon] |

**New at 60:** neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; legs: Sentinel's Silk Leggings; feet: Kayser's Boots of Precision; finger1: Blessed Band of Light; finger2: Don Julio's Band; trinket1: Eye of the Beast; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1442, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (undead, 000000000000000000-51400001000000000-000000000000000000)

Set DPS (verified): 39.1. Weights run: 2.1s. Verify run: 2.6s. 192 eligible items had no known source.

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

Set DPS (verified): 80.4. Weights run: 2.5s. Verify run: 4.9s. 318 eligible items had no known source.

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

Set DPS (verified): 162.3. Weights run: 2.6s. Verify run: 4.5s. 448 eligible items had no known source.

Race sweep: undead 162.3 ± 0.3, orc 161.1 ± 0.3, troll 159.0 ± 0.3 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.646 ± 0.014, crit=0.189 ± 0.011 per rating point (14 rating = 1%, 2.648 per %), hit=0.295 ± 0.019 per rating point (10 rating = 1%, 2.950 per %), spell_haste=not significant (0.233 ± 0.131), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.237 ± 0.117), mp5=0.780 ± 0.145

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Spellpower Goggles Xtreme (10502, -0.03 DPS) [crafted]; Corpseshroud (10574, -0.46 DPS) [dungeon]; Miner's Hat of the Deep (9429, -0.61 DPS, sim-verified) [dungeon] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Prodigious Shadowshard Pendant (17773, -0.17 DPS) [quest]; Glowing Eye of Mordresh (10769, -0.76 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -0.87 DPS) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.4 spell_power points (3.85 DPS) | yes | Green Silken Shoulders (7057, -0.07 DPS) [crafted]; Bloodmage Mantle (7684, -0.15 DPS) [dungeon]; Berylline Pads (4197, -0.48 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.20 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.29 DPS) [vendor] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 27.1 spell_power points (6.78 DPS) | yes | Cindercloth Robe (10042, -0.04 DPS) [crafted]; Robe of Power (7054, -0.87 DPS) [crafted]; Robe of the Magi (1716, -2.07 DPS, sim-verified) [world_drop] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (3.25 DPS) | yes | Spidertank Oilrag (9448, -0.41 DPS) [dungeon]; Radiant Silver Bracers (4545, -0.96 DPS) [quest]; Condor Bracers (15864, -1.50 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.2 spell_power points (5.55 DPS) | yes | Crimson Silk Gloves (7064, -1.09 DPS) [crafted]; Red Mageweave Gloves (10018, -1.19 DPS) [crafted]; Fiery Handwraps (254025, -2.87 DPS, sim-verified) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | 20.6 spell_power points (5.14 DPS) | yes | Defiler's Cloth Girdle (20166, -1.00 DPS) [rep]; Scorching Sash (4117, -1.36 DPS) [quest]; Deathmage Sash (10771, -3.96 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.8 spell_power points (5.43 DPS) | yes | Flame Leggings (253991, -0.51 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.34 DPS) [crafted]; Filigreed Flame Leggings (253941, -1.48 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.7 spell_power points (6.41 DPS) | yes | Gilded Slippers (254001, -3.29 DPS) [crafted]; Acidic Walkers (9454, -3.63 DPS) [dungeon]; Fiery Slippers (254005, -4.02 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.2 spell_power points (3.30 DPS) | yes | Advisor's Ring (19520, -0.47 DPS) [rep]; Runemender's Seal (273031, -1.17 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 12.2 spell_power points (3.04 DPS) | yes | Advisor's Ring (19521, -0.91 DPS) [rep]; Runemender's Seal (273031, -0.91 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+4.2 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Celestial Stave (9517, -7.13 DPS) [quest]; Hypnotic Blade (7714, -7.45 DPS) [dungeon]; Illusionary Rod (7713, -11.42 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Ragefire Wand (7513) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jaina's Firestarter (13064, -1.39 DPS) [world_drop]; Twisted Nether Wand (249144, -1.51 DPS) [crafted]; Plaguerot Sprig (10766, -2.13 DPS, sim-verified) [dungeon] |

**New at 40:** head: Papal Fez; neck: Triune Amulet; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Fiery Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Staff of Jordan; ranged: Ragefire Wand

No-known-source sample (15 of 448, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (undead, 230002000000000000-54500001230133031-003000000000000000)

Set DPS (verified): 244.3. Weights run: 2.5s. Verify run: 5.4s. 590 eligible items had no known source.

Race sweep: undead 244.3 ± 0.4, orc 242.3 ± 0.4, troll 239.6 ± 0.4 (dps); the guide race undead is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.846 ± 0.031, crit=0.394 ± 0.025 per rating point (14 rating = 1%, 5.511 per %), hit=0.684 ± 0.041 per rating point (10 rating = 1%, 6.836 per %), spell_haste=not significant (0.198 ± 0.244), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.740 ± 0.305), mp5=not significant (0.864 ± 0.323)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Flame (3075) | World drop [world_drop] | 58.9 spell_power points (12.24 DPS) | yes | Red Mageweave Headband (10033, -4.77 DPS) [crafted]; Gemburst Circlet (10751, -5.65 DPS) [quest]; Dreamweave Circlet (10041, -8.47 DPS, sim-verified) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 21.1 spell_power points (4.38 DPS) | yes | Mindburst Medallion (11196, -1.31 DPS) [quest]; Glowing Eye of Mordresh (10769, -1.52 DPS) [dungeon]; Horizon Choker (13085, -3.49 DPS, sim-verified) [world_drop] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 29.4 spell_power points (6.12 DPS) | yes | Rotgrip Mantle (17732, -0.25 DPS) [dungeon]; Netherflame Shoulders (254053, -0.26 DPS) [crafted]; Mantle of Lost Hope (22234, -1.60 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -0.81 DPS) [quest]; Runecloth Cloak (13860, -1.61 DPS) [crafted]; Cindercloth Cloak (14044, -2.38 DPS, sim-verified) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Embrace of the Wind Serpent (12462, +0.00 DPS) [world]; Cindercloth Vest (14042, -0.80 DPS) [crafted]; Robes of Insight (940, -0.97 DPS) [world_drop] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 22.6 spell_power points (4.71 DPS) | yes | Nethergeld Cuffs (254061, -1.25 DPS) [crafted]; Bloodband Bracers (11469, -1.62 DPS) [quest]; Aristocratic Cuffs (12546, -2.85 DPS, sim-verified) [dungeon] |
| hands | Fiery Gloves (254099) | Tailoring [crafted] | 33.4 spell_power points (6.94 DPS) | yes | Runecloth Gloves (13863, -1.48 DPS) [crafted]; Cindercloth Gloves (14043, -1.71 DPS) [crafted]; Dreamweave Gloves (10019, -4.55 DPS, sim-verified) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 27.3 spell_power points (5.68 DPS) | yes | Fiery Waistcord (254085, -0.14 DPS) [crafted]; Fiery Cord (254041, -0.43 DPS) [crafted]; Satyrmane Sash (17755, -1.01 DPS) [dungeon] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 31.6 spell_power points (6.56 DPS) | yes | Senior Designer's Pantaloons (11841, -0.06 DPS) [dungeon]; Cindercloth Leggings (12256, -0.30 DPS) [vendor]; Spellshock Leggings (9484, -3.13 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 29.2 spell_power points (6.07 DPS) | yes | Fiery Sandals (254111, -0.33 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -1.40 DPS) [vendor]; Mistwalker Boots (10629, -1.43 DPS) [dungeon] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 23.2 spell_power points (4.82 DPS) | yes | Snake Hoop (6750, -1.26 DPS) [quest]; Advisor's Ring (19519, -1.60 DPS) [rep]; Woodseed Hoop (17768, -1.81 DPS) [quest] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 17.9 spell_power points (3.72 DPS) | yes | Advisor's Ring (19519, -0.50 DPS) [rep]; Woodseed Hoop (17768, -0.71 DPS) [quest]; Snake Hoop (6750, -2.03 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.31 DPS) [quest]; Frozen Heart of the Mountain (249469, -0.57 DPS) [crafted]; Uther's Strength (11302, -0.60 DPS) [world_drop] |
| trinket2 | Fire Ruby (20036) | Destroy Morphaz [quest] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Evonice's Landin' Pilla (18951, -1.37 DPS) [quest]; Frozen Heart of the Mountain (249469, -1.63 DPS) [crafted]; Darkspear Voodoo Seal (272061, -2.23 DPS, sim-verified) [vendor] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Staff of Jordan (873, -7.82 DPS, sim-verified) [world_drop] |
| off_hand | Truesilver Conduit (249455) | Enchanting [crafted] | 16.8 spell_power points (3.49 DPS) | yes | Flask of Forest Mojo (19115, -0.21 DPS) [quest]; Celestial Orb (7515, -0.26 DPS) [quest]; Enthralled Sphere (11625, -2.09 DPS, sim-verified) [dungeon] |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lesser Eternal Wand (249232, -3.65 DPS) [crafted]; Wand of Allistarj (13065, -4.47 DPS) [world_drop]; Noxious Shooter (17745, -6.21 DPS, sim-verified) [dungeon] |

**New at 50:** head: Eye of Flame; neck: Arcane Crystal Pendant; shoulder: Kentic Amice; back: Spritecaster Cape; wrist: Netherflame Cuffs; hands: Fiery Gloves; waist: Ban'thok Sash; legs: Kilt of the Atal'ai Prophet; finger1: Brainlash; finger2: Cyclopean Band; trinket2: Fire Ruby; main_hand: Blade of Eternal Darkness; off_hand: Truesilver Conduit; ranged: Pyric Caduceus

No-known-source sample (15 of 590, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (undead, 230105200100000000-55500201230133031-003000000000000000)

Set DPS (verified): 393.5. Weights run: 2.9s. Verify run: 23.3s. 1436 eligible items had no known source.

Race sweep: undead 393.5 ± 0.7, orc 388.6 ± 0.7, troll 385.7 ± 0.7 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.599 ± 0.052, crit=0.463 ± 0.037 per rating point (14 rating = 1%, 6.476 per %), hit=1.239 ± 0.073 per rating point (10 rating = 1%, 12.390 per %), spell_haste=not significant (-2.546 ± 0.966), spell_penetration=not significant (0.000 ± 0.000), fire_power=0.847 ± 0.003, spirit=not significant (0.338 ± 0.669), mp5=not significant (-0.531 ± 0.679)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Eye of Flame (3075, -15.79 DPS, sim-verified) [world_drop] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 26.2 spell_power points (5.96 DPS) | yes | Star of Mystaria (12103, -0.80 DPS) [dungeon]; Lady Maye's Pendant (14558, -0.87 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -1.37 DPS) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burial Shawl (18681, +0.00 DPS) [dungeon]; Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -16.74 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -2.46 DPS) [world]; Shroud of the Exile (15421, -2.92 DPS) [quest]; Cloak of Flames (3475, -10.84 DPS, sim-verified) [world_drop] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Raiment (231596, -0.45 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -3.26 DPS) [pvp]; Alanna's Embrace (13314, -14.73 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 29.2 spell_power points (6.64 DPS) | yes | Runecloth Cuffs (254123, -2.77 DPS) [crafted]; Netherflame Cuffs (254065, -2.79 DPS) [crafted]; Manacle Cuffs (11962, -3.00 DPS) [quest] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Hands of Power (13253, -15.37 DPS, sim-verified) [dungeon] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 33.4 spell_power points (7.60 DPS) | yes | Defiler's Cloth Girdle (20163, -2.12 DPS) [rep]; Frostwolf Cloth Belt (19090, -2.14 DPS) [rep]; Ban'thok Sash (11662, -3.82 DPS, sim-verified) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Outrider's Silk Leggings (22747, +0.00 DPS) [rep]; General's Silk Trousers (231595, +0.00 DPS) [pvp]; The Postmaster's Trousers (13389, -20.51 DPS, sim-verified) [dungeon] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dragonrider Boots (18102, +0.00 DPS) [dungeon]; General's Silk Boots (231597, +0.00 DPS) [pvp]; The Postmaster's Treads (13391, -21.65 DPS, sim-verified) [dungeon] |
| finger1 | The Postmaster's Seal (13392) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, +0.00 DPS) [quest]; Channeler's Ring (272406, +0.00 DPS) [vendor]; Blessed Band of Light (272407, +0.00 DPS) [vendor] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, +0.00 DPS) [vendor]; Channeler's Ring (272406, -0.17 DPS) [vendor]; Eye of Orgrimmar (12545, -0.46 DPS) [quest] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Second Wind (11819, -5.03 DPS, sim-verified) [dungeon] |
| trinket2 | Eye of the Beast (13968) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Fire Ruby (20036, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironbark Staff (20220, -0.01 DPS) [rep]; High Warlord's War Staff (234549, -0.51 DPS) [pvp]; Blade of Eternal Darkness (17780, -6.10 DPS, sim-verified) [dungeon] |
| off_hand | Tome of Fiery Arcana (19311) | Frostwolf Clan [rep] | 31.3 spell_power points (7.14 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; Therazane's Touch (19315, -0.53 DPS) [rep]; High Warlord's Tome of Mending (234564, -1.81 DPS) [pvp] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, +0.00 DPS) [dungeon]; Oblivion's Touch (18761, -0.50 DPS) [dungeon] |

**New at 60:** head: Devout Crown; neck: Nacreous Shell Necklace; shoulder: Devout Mantle; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Devout Gloves; waist: Magician's Cord; legs: Devout Skirt; feet: Devout Sandals; finger1: The Postmaster's Seal; finger2: Maiden's Circle; trinket1: Burst of Knowledge; trinket2: Eye of the Beast; main_hand: Sageblade; off_hand: Tome of Fiery Arcana; ranged: Brilliant Wand

No-known-source sample (15 of 1436, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60, raid preset (undead, 050105200100000000-55500201230133031-003000000000000000)

Set DPS (verified): 807.4. Weights run: 3.1s. Verify run: 18.8s. 1436 eligible items had no known source.

Race sweep: undead 807.4 ± 0.9, orc 803.6 ± 0.9, troll 797.0 ± 0.9 (dps); the guide race undead is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.567 ± 0.046, crit=0.888 ± 0.060 per rating point (14 rating = 1%, 12.438 per %), hit=1.499 ± 0.120 per rating point (10 rating = 1%, 14.994 per %), spell_haste=-5.974 ± 1.389, spell_penetration=not significant (0.000 ± 0.000), fire_power=0.865 ± 0.002, spirit=not significant (0.431 ± 0.208), mp5=not significant (0.515 ± 0.372)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Flame (3075) | World drop [world_drop] | 47.2 spell_power points (27.39 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -0.56 DPS) [pvp]; Crimson Felt Hat (18727, -4.22 DPS, sim-verified) [dungeon] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 28.5 spell_power points (16.56 DPS) | yes | Lady Maye's Pendant (14558, -3.92 DPS) [world_drop]; Star of Mystaria (12103, -4.20 DPS, sim-verified) [dungeon]; Jeweled Amulet of Cainwyn (1443, -4.65 DPS) [world_drop] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 36.8 spell_power points (21.36 DPS) | yes | Warlord's Silk Amice (231594, -0.67 DPS) [pvp]; Champion's Silk Mantle (227104, -0.82 DPS) [pvp]; Shroud of the Nathrezim (18720, -6.96 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -7.70 DPS) [world]; Shroud of Arcane Mastery (22330, -8.30 DPS) [dungeon]; Cloak of Flames (3475, -11.28 DPS, sim-verified) [world_drop] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Raiment (231596, -0.92 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -8.13 DPS) [pvp]; Alanna's Embrace (13314, -21.98 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 29.6 spell_power points (17.15 DPS) | yes | Netherflame Cuffs (254065, -7.07 DPS) [crafted]; Runecloth Cuffs (254123, -7.48 DPS) [crafted]; Manacle Cuffs (11962, -8.15 DPS) [quest] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 39.5 spell_power points (22.95 DPS) | yes | General's Silk Handguards (16540, -2.08 DPS) [vendor]; General's Silk Gauntlets (231599, -2.08 DPS) [vendor]; Hands of Power (13253, -4.88 DPS, sim-verified) [dungeon] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 32.7 spell_power points (18.96 DPS) | yes | Defiler's Cloth Girdle (20163, -1.64 DPS) [rep]; Clutch of Andros (13956, -4.00 DPS) [dungeon]; Magician's Cord (272393, -6.55 DPS, sim-verified) [vendor] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 43.1 spell_power points (25.00 DPS) | yes | Legionnaire's Silk Legguards (227107, +0.00 DPS) [pvp]; General's Silk Trousers (231595, +0.00 DPS) [pvp]; Magister's Leggings (16687, -7.41 DPS, sim-verified) [dungeon] |
| feet | Kayser's Boots of Precision (22231) | Blackrock Spire: War Master Voone [dungeon] | 29.2 spell_power points (16.95 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; First Sergeant's Dreadweave Boots (220909, -0.64 DPS) [vendor]; Fire Striders (13369, -1.13 DPS) [dungeon] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -1.20 DPS) [vendor]; Maiden's Circle (13001, -2.26 DPS) [world_drop]; Naglering (11669, -17.07 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Maiden's Circle (13001, -1.67 DPS) [world_drop]; Don Mauricio's Band of Domination (22433, -2.32 DPS) [dungeon]; Channeler's Ring (272406, -4.47 DPS, sim-verified) [vendor] |
| trinket1 | Eye of the Beast (13968) | For The Horde! [quest] | sim-verified (+9.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Fire Ruby (20036, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, -1.74 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -2.61 DPS) [crafted]; Fire Ruby (20036, -3.41 DPS) [quest] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -1.43 DPS) [pvp]; Elemental Mage Staff (944, -12.17 DPS, sim-verified) [world_drop]; Headmaster's Charge (13937, -13.21 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Pyric Caduceus (11748, -0.29 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -1.12 DPS) [dungeon] |

**New at 60:** neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; legs: Outrider's Silk Leggings; feet: Kayser's Boots of Precision; finger1: Blessed Band of Light; finger2: Don Julio's Band; trinket1: Eye of the Beast; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1436, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

