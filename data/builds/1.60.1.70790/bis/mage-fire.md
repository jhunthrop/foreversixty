# Leveling BiS: Fire

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 000000000000000000-51400001000000000-000000000000000000)

Set DPS (verified): 38.7. Weights run: 3.0s. Verify run: 2.6s. 181 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.018 ± 0.016, crit=0.143 ± 0.007 per rating point (14 rating = 1%, 2.004 per %), hit=0.311 ± 0.016 per rating point (10 rating = 1%, 3.108 per %), spell_haste=not significant (0.577 ± 0.204), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.157 ± 0.081), mp5=1.049 ± 0.217

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Flame Circlet (253953) | Tailoring [crafted] | 11.0 spell_power points (1.55 DPS) | yes | Pristine Circlet (253949, -0.70 DPS) [crafted]; Shadow Goggles (4373, -1.24 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.46 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.3 spell_power points (0.04 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 14.2 spell_power points (2.00 DPS) | yes | Reinforced Woolen Shoulders (4315, -1.17 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.43 DPS) [crafted]; Slime-encrusted Pads (6461, -1.55 DPS) [dungeon] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | sim-verified (38.7 DPS) | yes | Heavy Woolen Cloak (4311, -0.06 DPS) [crafted]; Sanguine Cape (14376, -0.14 DPS) [world_drop]; Caretaker's Cape (20428, -0.25 DPS) [rep] |
| chest | Filigreed Flame Gown (253905) | Tailoring [crafted] | 14.6 spell_power points (2.05 DPS) | yes | Filigreed Pristine Gown (253901, -0.59 DPS, sim-verified) [crafted]; Gray Woolen Robe (2585, -0.77 DPS) [crafted]; Mystic's Wrap (14369, -1.05 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 8.2 spell_power points (1.16 DPS) | yes | Repurposed Hair Band (281256, -0.83 DPS) [quest]; Mystic's Bracelets (14366, -0.87 DPS) [world_drop]; Bright Bracers (3647, -0.91 DPS, sim-verified) [world_drop] |
| hands | Flame Gloves (253917) | Tailoring [crafted] | 10.1 spell_power points (1.42 DPS) | yes | Phoenix Gloves (4331, -0.30 DPS, sim-verified) [crafted]; Pristine Gloves (253913, -0.42 DPS) [crafted]; Serpent Gloves (5970, -0.43 DPS) [dungeon] |
| waist | Flame Sash (253929) | Tailoring [crafted] | 11.1 spell_power points (1.56 DPS) | yes | Pristine Sash (253925, -0.42 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.57 DPS) [crafted]; Keller's Girdle (2911, -0.87 DPS, sim-verified) [world_drop] |
| legs | Filigreed Flame Leggings (253941) | Tailoring [crafted] | 17.7 spell_power points (2.50 DPS) | yes | Phoenix Pants (4317, -0.53 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.70 DPS) [crafted]; Darkweave Breeches (12987, -1.36 DPS) [world_drop] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 11.1 spell_power points (1.56 DPS) | yes | Flame Boots (253893, -0.28 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.71 DPS) [crafted]; Red Woolen Boots (4313, -0.89 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 7.1 spell_power points (1.00 DPS) | yes | Lavishly Jeweled Ring (1156, -0.14 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.57 DPS) [world_drop]; The 1 Ring (8350, -0.83 DPS) [world] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 7.0 spell_power points (0.99 DPS) | yes | Volcanic Rock Ring (12053, -0.56 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.57 DPS, sim-verified) [dungeon]; The 1 Ring (8350, -0.83 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | 30.8 spell_power points (4.34 DPS) | yes | Emberstone Staff (5201, -0.96 DPS) [dungeon]; Staff of the Blessed Seer (2271, -1.01 DPS, sim-verified) [dungeon]; Staff of Westfall (2042, -1.24 DPS) [quest] |
| off_hand | Antipodean Rod (2879) | World drop [world_drop] | 6.0 spell_power points (0.85 DPS) | yes | Tear of Grief (5611, -0.19 DPS) [quest]; Seer's Fine Stein (7608, -0.39 DPS) [world_drop]; Mathystra Channeling Orb (281291, -0.58 DPS, sim-verified) [quest] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | sim-verified (38.7 DPS) | yes | Skycaller (12984, -0.45 DPS, sim-verified) [world_drop]; Deepblaze (279896, -1.22 DPS) [quest]; Sizzle Stick (8071, -1.22 DPS) [quest] |

**New at 20:** head: Flame Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Flame Gown; wrist: Mindthrust Bracers; hands: Flame Gloves; waist: Flame Sash; legs: Filigreed Flame Leggings; feet: Spidersilk Boots; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Evocator's Blade; off_hand: Antipodean Rod; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 181, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 30 (human, 000000000000000000-51520101230100000-000000000000000000)

Set DPS (verified): 76.7. Weights run: 3.5s. Verify run: 3.8s. 309 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.638 ± 0.012, crit=0.139 ± 0.008 per rating point (14 rating = 1%, 1.951 per %), hit=0.281 ± 0.015 per rating point (10 rating = 1%, 2.807 per %), spell_haste=0.366 ± 0.058, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (-0.038 ± 0.050), mp5=not significant (0.077 ± 0.082)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Flame Circlet (253979) | Tailoring [crafted] | 14.0 spell_power points (2.27 DPS) | yes | Enchanter's Cowl (4322, -0.26 DPS) [crafted]; Holy Shroud (2721, -0.49 DPS) [world_drop]; Flame Circlet (253953, -0.49 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.8 spell_power points (0.62 DPS) | yes | Seal of the Expedition (271920, -0.10 DPS) [rep]; Crystal Starfire Medallion (5003, -0.21 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.21 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.7 spell_power points (2.39 DPS) | yes | Death Speaker Mantle (6685, -0.62 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.65 DPS) [world_drop]; Invoker's Mantle (215365, -0.74 DPS) [crafted] |
| back | Expeditionary's Cape (271923) | Theramore Expeditionary Force [rep] | sim-verified (76.7 DPS) | yes | Hillman's Cloak (3719, -0.02 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.18 DPS) [crafted]; Vine Pruner's Cloak (279835, -1.20 DPS, sim-verified) [quest] |
| chest | Flame Gown (253965) | Tailoring [crafted] | 17.5 spell_power points (2.84 DPS) | yes | Beguiler Robes (7728, -0.13 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.29 DPS) [dungeon]; Green Silk Armor (7065, -0.56 DPS, sim-verified) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (2.11 DPS) | yes | Spidertank Oilrag (9448, -0.72 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -1.28 DPS) [world_drop]; Nightsky Wristbands (6407, -1.49 DPS) [world_drop] |
| hands | Phoenix Gloves (4331) | Tailoring [crafted] | 9.0 spell_power points (1.46 DPS) | yes | Flame Gloves (253917, -0.01 DPS) [crafted]; Town Clerk's Mittens (270029, -0.25 DPS) [quest]; Serpent Gloves (5970, -0.32 DPS) [dungeon] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.9 spell_power points (2.10 DPS) | yes | Belt of Arugal (6392, -0.32 DPS) [dungeon]; Crimson Silk Belt (7055, -0.40 DPS) [crafted]; Invoker's Cord (215366, -0.44 DPS) [crafted] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | 18.5 spell_power points (3.00 DPS) | yes | Smoldering Pants (3073, -0.89 DPS) [world]; Phoenix Pants (4317, -0.96 DPS) [crafted]; Filigreed Flame Leggings (253941, -1.10 DPS, sim-verified) [crafted] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 17.5 spell_power points (2.84 DPS) | yes | Acidic Walkers (9454, -1.20 DPS) [dungeon]; Spidersilk Boots (4320, -1.29 DPS) [crafted]; Gilded Slippers (254001, -1.32 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.2 spell_power points (1.49 DPS) | yes | Minor Channeling Ring (1449, -0.47 DPS) [quest]; Ring of Power Regulation (273030, -0.48 DPS) [dungeon]; Lavishly Jeweled Ring (1156, -0.87 DPS) [dungeon] |
| finger2 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.2 spell_power points (1.16 DPS) | yes | Minor Channeling Ring (1449, -0.14 DPS) [quest]; Ring of Power Regulation (273030, -0.15 DPS) [dungeon]; Lavishly Jeweled Ring (1156, -0.54 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-verified (76.7 DPS) | yes | Hydrocane (9452, -0.65 DPS) [dungeon]; Tendonscraper Dagger (273053, -0.65 DPS) [world]; Glimmering Staff (249392, -3.05 DPS, sim-verified) [crafted] |
| off_hand | Rod of Molten Fire (2565) | World drop [world_drop] | 13.0 spell_power points (2.11 DPS) | yes | Orb of Mystic Insight (249394, -0.64 DPS, sim-verified) [crafted]; Antipodean Rod (2879, -1.14 DPS) [world_drop]; Eye of Paleth (2943, -1.46 DPS) [quest] |
| ranged | Scorching Wand (5213) | World drop [world_drop] | sim-verified (76.7 DPS) | yes | Wand of Mana Concentration (273043, +0.00 DPS) [world]; Greater Mystic Wand (217287, -0.12 DPS) [crafted]; Starfaller (13063, -0.86 DPS, sim-verified) [world_drop] |

**New at 30:** head: Filigreed Flame Circlet; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Expeditionary's Cape; chest: Flame Gown; wrist: Phoenix Bindings; hands: Phoenix Gloves; waist: Highlander's Cloth Girdle; legs: Flame Leggings; feet: Fiery Slippers; finger1: Snake Hoop; finger2: Lorekeeper's Ring; main_hand: Swinetusk Shank; off_hand: Rod of Molten Fire; ranged: Scorching Wand

No-known-source sample (15 of 309, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (human, 000000000000000000-51520101230133031-000000000000000000)

Set DPS (verified): 162.4. Weights run: 3.9s. Verify run: 5.9s. 400 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.631 ± 0.019, crit=0.344 ± 0.021 per rating point (14 rating = 1%, 4.811 per %), hit=0.539 ± 0.034 per rating point (10 rating = 1%, 5.393 per %), spell_haste=not significant (0.519 ± 0.171), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=0.879 ± 0.182, mp5=0.934 ± 0.221

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 26.6 spell_power points (9.45 DPS) | yes | Spellpower Goggles Xtreme (10502, -2.00 DPS) [crafted]; Electromagnetic Gigaflux Reactivator (9492, -2.01 DPS) [dungeon]; Whitemane's Chapeau (7720, -2.38 DPS, sim-verified) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Necklace of Calisea (1714, -0.52 DPS) [dungeon]; Amberglow Talisman (10824, -1.16 DPS) [quest]; Triune Amulet (7722, -2.16 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Green Silken Shoulders (7057, -0.09 DPS) [crafted]; Bloodmage Mantle (7684, -0.19 DPS) [dungeon]; Mantle of Doan (7712, -1.67 DPS, sim-verified) [dungeon] |
| back | Expeditionary's Cape (271922) | Theramore Expeditionary Force [rep] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Caretaker's Cape (19532, -0.02 DPS) [rep]; Glowing Thresher Cape (6901, -0.15 DPS) [dungeon]; Icy Cloak (4327, -1.53 DPS, sim-verified) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 36.0 spell_power points (12.77 DPS) | yes | Robe of the Magi (1716, -2.06 DPS) [world_drop]; Robe of Power (7054, -2.62 DPS) [crafted]; Cindercloth Robe (10042, -3.19 DPS) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (4.61 DPS) | yes | Spidertank Oilrag (9448, -0.43 DPS) [dungeon]; Arcane Runed Bracers (4744, -1.42 DPS) [quest]; Earthen Silk Cuffs (254019, -1.63 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 26.7 spell_power points (9.47 DPS) | yes | Fiery Handwraps (254025, -2.93 DPS) [crafted]; Crimson Silk Gloves (7064, -3.16 DPS) [crafted]; Red Mageweave Gloves (10018, -3.32 DPS) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | 24.3 spell_power points (8.63 DPS) | yes | Gilded Cord (254037, -2.13 DPS) [crafted]; Scorching Sash (4117, -2.74 DPS, sim-verified) [quest]; Highlander's Cloth Girdle (20098, -2.77 DPS) [rep] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | 22.8 spell_power points (8.10 DPS) | yes | Stoneweaver Leggings (9407, -0.08 DPS) [dungeon]; Red Mageweave Pants (10009, -0.44 DPS) [crafted]; Black Mageweave Leggings (9999, -0.89 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 30.2 spell_power points (10.70 DPS) | yes | Fiery Slippers (254005, -1.66 DPS, sim-verified) [crafted]; Furen's Boots (13100, -4.14 DPS) [world_drop]; Thoughtcast Boots (10578, -4.87 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.2 spell_power points (4.67 DPS) | yes | Welken Ring (5011, -1.37 DPS) [world_drop]; Snake Hoop (6750, -1.42 DPS) [quest]; Electrocutioner Lagnut (9447, -1.42 DPS) [dungeon] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 11.8 spell_power points (4.19 DPS) | yes | Snake Hoop (6750, -0.94 DPS) [quest]; Electrocutioner Lagnut (9447, -0.94 DPS) [dungeon]; Welken Ring (5011, -1.77 DPS, sim-verified) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -13.16 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -13.42 DPS) [dungeon]; Celestial Stave (9517, -13.65 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ragefire Wand (7513) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jaina's Firestarter (13064, -1.62 DPS, sim-verified) [world_drop]; Twisted Nether Wand (249144, -2.10 DPS) [crafted]; Burning Sliver (5249, -2.24 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Expeditionary's Cape; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Fiery Cord; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Staff of Jordan; ranged: Ragefire Wand

No-known-source sample (15 of 400, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (human, 252010000000000000-51520101230133031-000000000000000000)

Set DPS (verified): 213.4. Weights run: 3.6s. Verify run: 6.3s. 512 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.830 ± 0.031, crit=0.352 ± 0.024 per rating point (14 rating = 1%, 4.932 per %), hit=0.608 ± 0.040 per rating point (10 rating = 1%, 6.076 per %), spell_haste=not significant (0.840 ± 0.248), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.904 ± 0.283), mp5=not significant (0.534 ± 0.308)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Flame (3075) | World drop [world_drop] | 60.3 spell_power points (11.28 DPS) | yes | Red Mageweave Headband (10033, -4.62 DPS) [crafted]; Gemburst Circlet (10751, -4.85 DPS) [quest]; Dreamweave Circlet (10041, -7.23 DPS, sim-verified) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 21.0 spell_power points (3.92 DPS) | yes | Horizon Choker (13085, -1.07 DPS) [world_drop]; Triune Amulet (7722, -1.24 DPS) [dungeon]; Mindburst Medallion (11196, -3.54 DPS, sim-verified) [quest] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 30.2 spell_power points (5.65 DPS) | yes | Netherflame Shoulders (254053, -0.21 DPS) [crafted]; Rotgrip Mantle (17732, -0.42 DPS) [dungeon]; Flameseer Mantle (11310, -1.34 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (213.4 DPS) | yes | Cindercloth Cloak (14044, -0.89 DPS) [crafted]; Featherskin Cape (10843, -1.41 DPS) [world]; Expeditionary's Cape (271921, -1.55 DPS) [rep] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 41.2 spell_power points (7.70 DPS) | yes | Dreamweave Vest (10021, -0.58 DPS) [crafted]; Acumen Robes (17775, -1.05 DPS) [quest]; Robes of Insight (940, -1.29 DPS) [world_drop] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 23.3 spell_power points (4.36 DPS) | yes | Aristocratic Cuffs (12546, -1.02 DPS) [dungeon]; Nethergeld Cuffs (254061, -1.12 DPS) [crafted]; Bloodband Bracers (11469, -1.52 DPS) [quest] |
| hands | Fiery Gloves (254099) | Tailoring [crafted] | 34.5 spell_power points (6.45 DPS) | yes | Runecloth Gloves (13863, -1.29 DPS) [crafted]; Cindercloth Gloves (14043, -1.42 DPS) [crafted]; Dreamweave Gloves (10019, -2.64 DPS, sim-verified) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 27.2 spell_power points (5.09 DPS) | yes | Fiery Waistcord (254085, -0.14 DPS) [crafted]; Fiery Cord (254041, -0.21 DPS) [crafted]; Ghostweave Cord (254073, -0.61 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 34.2 spell_power points (6.39 DPS) | yes | Senior Designer's Pantaloons (11841, -0.02 DPS) [dungeon]; Cindercloth Leggings (12256, -0.31 DPS) [vendor]; Spellshock Leggings (9484, -0.54 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 30.3 spell_power points (5.67 DPS) | yes | Fiery Sandals (254111, -0.53 DPS) [crafted]; Mistwalker Boots (10629, -1.39 DPS) [dungeon]; Vinerot Sandals (17748, -1.46 DPS) [dungeon] |
| finger1 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 18.4 spell_power points (3.44 DPS) | yes | Eye of Adaegus (5266, -0.49 DPS) [world_drop]; Philanthropist's Ring (281635, -0.80 DPS) [quest]; Lorekeeper's Ring (19523, -0.80 DPS) [rep] |
| finger2 | Brainlash (6440) | Zul'Farrak [dungeon] | 17.0 spell_power points (3.17 DPS) | yes | Philanthropist's Ring (281635, -0.53 DPS) [quest]; Lorekeeper's Ring (19523, -0.53 DPS) [rep]; Eye of Adaegus (5266, -2.87 DPS, sim-verified) [world_drop] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (213.4 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Uther's Strength (11302, -0.23 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (213.4 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (213.4 DPS) | yes | Glowing Brightwood Staff (812, -2.35 DPS) [world_drop]; Spire of Hakkar (10844, -3.42 DPS) [world]; Kindling Stave (11750, -4.58 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (213.4 DPS) | yes | Noxious Shooter (17745, -2.72 DPS) [dungeon]; Wand of Allistarj (13065, -4.11 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.19 DPS, sim-verified) [crafted] |

**New at 50:** head: Eye of Flame; neck: Arcane Crystal Pendant; shoulder: Kentic Amice; back: Spritecaster Cape; chest: Embrace of the Wind Serpent; wrist: Netherflame Cuffs; hands: Fiery Gloves; waist: Ban'thok Sash; legs: Kilt of the Atal'ai Prophet; finger1: Cyclopean Band; finger2: Brainlash; trinket1: Thunderbrew's Boot Flask; trinket2: Frozen Heart of the Mountain; ranged: Pyric Caduceus

No-known-source sample (15 of 512, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (human, 252015000000200000-51520101230133031-003000000000000000)

Set DPS (verified): 375.4. Weights run: 4.0s. Verify run: 30.2s. 1299 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.542 ± 0.039, crit=0.558 ± 0.033 per rating point (14 rating = 1%, 7.814 per %), hit=1.170 ± 0.061 per rating point (10 rating = 1%, 11.699 per %), spell_haste=1.945 ± 0.436, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.261 ± 0.521), mp5=not significant (0.177 ± 0.513)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Flame (3075) | World drop [world_drop] | 51.0 spell_power points (11.64 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -2.48 DPS) [pvp]; Crimson Felt Hat (18727, -5.53 DPS, sim-verified) [dungeon] |
| neck | Diana's Pearl Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 25.0 spell_power points (5.71 DPS) | yes | Arcane Crystal Pendant (20037, -1.32 DPS) [quest]; Star of Mystaria (12103, -1.45 DPS) [dungeon]; Archlight Talisman (15856, -1.74 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Silk Mantle (227102, +0.00 DPS) [pvp]; Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -11.08 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -2.25 DPS) [world]; Mageflame Cloak (13007, -2.52 DPS) [world_drop]; Spritecaster Cape (11623, -3.01 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Vestments (231603, -0.62 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -3.42 DPS) [pvp]; Robes of Fiery Devastation (279266, -11.42 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (+3.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [rep] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Hands of Power (13253, -1.50 DPS, sim-verified) [dungeon] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Highlander's Cloth Girdle (20047, +0.00 DPS) [rep]; Magician's Cord (272393, -8.80 DPS, sim-verified) [vendor] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Sentinel's Silk Leggings (237815, +0.00 DPS) [vendor]; The Postmaster's Trousers (13389, -13.04 DPS, sim-verified) [dungeon] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fire Striders (13369, +0.00 DPS) [dungeon]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; The Postmaster's Treads (13391, -13.78 DPS, sim-verified) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-verified (375.4 DPS) | yes | Flaming Band (12926, +0.00 DPS) [dungeon]; Songstone of Ironforge (12543, -0.12 DPS) [quest]; Channeler's Ring (272406, -0.45 DPS) [vendor] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Flaming Band (12926, +0.00 DPS) [dungeon]; Songstone of Ironforge (12543, -0.22 DPS) [quest]; Channeler's Ring (272406, -0.55 DPS) [vendor] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -0.34 DPS) [crafted]; Uther's Strength (11302, -1.37 DPS) [world_drop]; Second Wind (11819, -6.81 DPS, sim-verified) [dungeon] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -1.82 DPS) [vendor]; Frozen Heart of the Mountain (249469, -2.62 DPS) [crafted]; Enriched Seal (271903, -7.61 DPS, sim-verified) [rep] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Stave (234571, -0.61 DPS) [pvp]; Sageblade (22383, -0.64 DPS) [crafted]; Elemental Mage Staff (944, -1.45 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, -3.18 DPS) [dungeon]; Brilliant Wand (249385, -4.16 DPS) [crafted]; Ritssyn's Wand of Bad Mojo (22408, -5.87 DPS, sim-verified) [dungeon] |

**New at 60:** neck: Diana's Pearl Necklace; shoulder: Devout Mantle; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Devout Bracers; hands: Gloves of Spell Mastery; waist: Devout Belt; legs: Devout Skirt; feet: Devout Sandals; finger1: Maiden's Circle; finger2: Blessed Band of Light; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Torch of Light

No-known-source sample (15 of 1299, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (human, 050105200100000000-55500201230133031-003000000000000000)

Set DPS (verified): 808.4. Weights run: 4.7s. Verify run: 22.6s. 1299 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.878 ± 0.049, crit=0.635 ± 0.047 per rating point (14 rating = 1%, 8.888 per %), hit=1.231 ± 0.085 per rating point (10 rating = 1%, 12.309 per %), spell_haste=not significant (-4.039 ± 1.073), spell_penetration=not significant (0.000 ± 0.000), fire_power=0.856 ± 0.002, spirit=not significant (0.621 ± 0.338), mp5=not significant (0.223 ± 0.496)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Flame (3075) | World drop [world_drop] | 51.8 spell_power points (24.90 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.15 DPS) [pvp]; The Postmaster's Band (13390, -9.50 DPS, sim-verified) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Lady Maye's Pendant (14558, +0.00 DPS) [world_drop]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 37.3 spell_power points (17.93 DPS) | yes | Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Mantle (227102, -0.61 DPS) [pvp]; Theramore Shoulderpads (271953, -10.53 DPS, sim-verified) [rep] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -5.93 DPS) [dungeon]; Shroud of Arcane Mastery (22330, -6.42 DPS) [dungeon]; Crystalline Threaded Cape (20697, -8.88 DPS, sim-verified) [world] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Vestments (231603, +0.00 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -5.53 DPS) [pvp]; Alanna's Embrace (13314, -21.57 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 33.4 spell_power points (16.04 DPS) | yes | Marshal's Silk Bracers (16438, -5.04 DPS) [pvp]; Netherflame Cuffs (254065, -6.24 DPS) [crafted]; Runecloth Cuffs (254123, -6.53 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 35.6 spell_power points (17.09 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Hands of Power (13253, -11.71 DPS, sim-verified) [dungeon] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 40.1 spell_power points (19.26 DPS) | yes | Ban'thok Sash (11662, -2.93 DPS) [dungeon]; Marshal's Silk Sash (16439, -5.08 DPS) [pvp]; Clutch of Andros (13956, -5.32 DPS) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 58.6 spell_power points (28.14 DPS) | yes | Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -5.11 DPS) [pvp]; Magister's Leggings (16687, -6.86 DPS) [dungeon] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (808.4 DPS) | yes | Dragonrider Boots (18102, +0.00 DPS) [dungeon]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Knight-Lieutenant's Silk Walkers (227112, -0.36 DPS) [pvp] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -1.10 DPS) [vendor]; Songstone of Ironforge (12543, -1.41 DPS) [quest]; Naglering (11669, -14.89 DPS, sim-verified) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -0.29 DPS) [vendor]; Songstone of Ironforge (12543, -0.60 DPS) [quest]; Don Julio's Band (19325, -3.21 DPS) [rep] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -4.80 DPS, sim-verified) [vendor]; Burst of Knowledge (11832, -4.81 DPS) [dungeon]; Ankh of Life (1713, -6.99 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Ankh of Life (1713, -1.74 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -2.34 DPS) [quest] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Staff of Jordan (873, -8.03 DPS) [world_drop]; Elemental Mage Staff (944, -17.00 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, -0.29 DPS) [dungeon]; Torch of Light (279246, -1.02 DPS) [crafted]; Pyric Caduceus (11748, -1.07 DPS) [dungeon] |

**New at 60:** neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Magician's Cord; legs: Sentinel's Silk Leggings; feet: Omnicast Boots; finger1: Channeler's Ring; finger2: Maiden's Circle; trinket1: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1299, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (undead, 000000000000000000-51400001000000000-000000000000000000)

Set DPS (verified): 38.9. Weights run: 3.0s. Verify run: 3.7s. 175 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.018 ± 0.016, crit=0.143 ± 0.007 per rating point (14 rating = 1%, 2.004 per %), hit=0.311 ± 0.016 per rating point (10 rating = 1%, 3.108 per %), spell_haste=not significant (0.577 ± 0.204), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.157 ± 0.081), mp5=1.049 ± 0.217

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Flame Circlet (253953) | Tailoring [crafted] | 11.0 spell_power points (1.55 DPS) | yes | Pristine Circlet (253949, -0.70 DPS) [crafted]; Shadow Goggles (4373, -1.21 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.46 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.5 spell_power points (0.07 DPS) | yes | Scholarly Pendant (277203, -0.02 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 14.2 spell_power points (2.00 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.81 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.43 DPS) [crafted]; Slime-encrusted Pads (6461, -1.55 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Sanguine Cape (14376, -0.08 DPS) [world_drop]; Pearl-clasped Cloak (5542, -0.10 DPS, sim-verified) [crafted]; Battle Healer's Cloak (20427, -0.19 DPS) [rep] |
| chest | Filigreed Flame Gown (253905) | Tailoring [crafted] | 14.6 spell_power points (2.05 DPS) | yes | Filigreed Pristine Gown (253901, -0.60 DPS, sim-verified) [crafted]; Gray Woolen Robe (2585, -0.77 DPS) [crafted]; Mystic's Wrap (14369, -1.05 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 8.2 spell_power points (1.16 DPS) | yes | Bright Bracers (3647, -0.41 DPS, sim-verified) [world_drop]; Featherbead Bracers (15452, -0.73 DPS) [quest]; Owlbeard Bracers (16981, -0.73 DPS) [quest] |
| hands | Flame Gloves (253917) | Tailoring [crafted] | 10.1 spell_power points (1.42 DPS) | yes | Phoenix Gloves (4331, -0.15 DPS) [crafted]; Pristine Gloves (253913, -0.42 DPS) [crafted]; Serpent Gloves (5970, -0.43 DPS) [dungeon] |
| waist | Flame Sash (253929) | Tailoring [crafted] | 11.1 spell_power points (1.56 DPS) | yes | Pristine Sash (253925, -0.42 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.57 DPS) [crafted]; Keller's Girdle (2911, -0.77 DPS, sim-verified) [world_drop] |
| legs | Filigreed Flame Leggings (253941) | Tailoring [crafted] | 17.7 spell_power points (2.50 DPS) | yes | Phoenix Pants (4317, -0.52 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.70 DPS) [crafted]; Darkweave Breeches (12987, -1.36 DPS) [world_drop] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 11.1 spell_power points (1.56 DPS) | yes | Flame Boots (253893, -0.30 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.71 DPS) [crafted]; Red Woolen Boots (4313, -0.89 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 7.1 spell_power points (1.00 DPS) | yes | Loop of Sacrifice (281673, -0.28 DPS) [quest]; Volcanic Rock Ring (12053, -0.57 DPS) [world_drop]; The 1 Ring (8350, -0.83 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 6.1 spell_power points (0.86 DPS) | yes | Loop of Sacrifice (281673, -0.14 DPS) [quest]; Volcanic Rock Ring (12053, -0.43 DPS) [world_drop]; The 1 Ring (8350, -0.70 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Staff of the Blessed Seer (2271, -0.18 DPS) [dungeon]; Crescent Staff (6505, -0.35 DPS, sim-verified) [quest]; Gnarled Necromancer's Staff (251534, -0.46 DPS) [quest] |
| off_hand | Antipodean Rod (2879) | World drop [world_drop] | 6.0 spell_power points (0.85 DPS) | yes | Seer's Fine Stein (7608, -0.39 DPS) [world_drop]; Spellbinder Orb (15926, -0.42 DPS) [dungeon]; Tork Wrench (11855, -0.56 DPS) [quest] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller (12984, -0.66 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -1.14 DPS) [quest]; Greater Magic Wand (11288, -1.18 DPS) [crafted] |

**New at 20:** head: Flame Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Flame Gown; wrist: Mindthrust Bracers; hands: Flame Gloves; waist: Flame Sash; legs: Filigreed Flame Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Evocator's Blade; off_hand: Antipodean Rod; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 175, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 23173 Abomination Skin Leggings; 209618 Insignia of the Alliance

### Band 30 (undead, 000000000000000000-51520101230100000-000000000000000000)

Set DPS (verified): 77.7. Weights run: 3.5s. Verify run: 4.0s. 301 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.638 ± 0.012, crit=0.139 ± 0.008 per rating point (14 rating = 1%, 1.951 per %), hit=0.281 ± 0.015 per rating point (10 rating = 1%, 2.807 per %), spell_haste=0.366 ± 0.058, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (-0.038 ± 0.050), mp5=not significant (0.077 ± 0.082)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Flame Circlet (253979) | Tailoring [crafted] | 14.0 spell_power points (2.27 DPS) | yes | Holy Shroud (2721, -0.49 DPS) [world_drop]; Flame Circlet (253953, -0.49 DPS) [crafted]; Enchanter's Cowl (4322, -0.79 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.8 spell_power points (0.62 DPS) | yes | Darkspear Warding Pendant (272075, -0.10 DPS) [vendor]; Crystal Starfire Medallion (5003, -0.21 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.21 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.7 spell_power points (2.39 DPS) | yes | Death Speaker Mantle (6685, -0.28 DPS) [dungeon]; Magician's Mantle (12998, -0.65 DPS) [world_drop]; Invoker's Mantle (215365, -0.74 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | sim-verified (77.7 DPS) | yes | Hillman's Cloak (3719, -0.02 DPS) [crafted]; Windsong Drape (15468, -0.02 DPS) [quest]; Heavy Woolen Cloak (4311, -0.18 DPS) [crafted] |
| chest | Flame Gown (253965) | Tailoring [crafted] | 17.5 spell_power points (2.84 DPS) | yes | Beguiler Robes (7728, -0.13 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.29 DPS) [dungeon]; Green Silk Armor (7065, -1.07 DPS, sim-verified) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (2.11 DPS) | yes | Spidertank Oilrag (9448, -0.86 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -1.28 DPS) [world_drop]; Nightsky Wristbands (6407, -1.49 DPS) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 9.2 spell_power points (1.49 DPS) | yes | Phoenix Gloves (4331, -0.03 DPS) [crafted]; Flame Gloves (253917, -0.04 DPS) [crafted]; Serpent Gloves (5970, -0.36 DPS) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.9 spell_power points (2.10 DPS) | yes | Warsong Sash (16975, -0.31 DPS) [quest]; Belt of Arugal (6392, -0.32 DPS) [dungeon]; Crimson Silk Belt (7055, -0.40 DPS) [crafted] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | 18.5 spell_power points (3.00 DPS) | yes | Smoldering Pants (3073, -0.89 DPS) [world]; Phoenix Pants (4317, -0.96 DPS) [crafted]; Filigreed Flame Leggings (253941, -1.16 DPS, sim-verified) [crafted] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 17.5 spell_power points (2.84 DPS) | yes | Acidic Walkers (9454, -1.20 DPS) [dungeon]; Spidersilk Boots (4320, -1.29 DPS) [crafted]; Gilded Slippers (254001, -1.31 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.2 spell_power points (1.49 DPS) | yes | Ring of Power Regulation (273030, -0.48 DPS) [dungeon]; Solid Gold Band (279845, -0.84 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.87 DPS) [dungeon] |
| finger2 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.2 spell_power points (1.16 DPS) | yes | Ring of Power Regulation (273030, -0.15 DPS) [dungeon]; Solid Gold Band (279845, -0.51 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.54 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-verified (77.7 DPS) | yes | Hydrocane (9452, -0.65 DPS) [dungeon]; Tendonscraper Dagger (273053, -0.65 DPS) [world]; Glimmering Staff (249392, -3.18 DPS, sim-verified) [crafted] |
| off_hand | Rod of Molten Fire (2565) | World drop [world_drop] | 13.0 spell_power points (2.11 DPS) | yes | Orb of Mystic Insight (249394, -0.83 DPS, sim-verified) [crafted]; Antipodean Rod (2879, -1.14 DPS) [world_drop]; Witch's Finger (16887, -1.28 DPS) [quest] |
| ranged | Scorching Wand (5213) | World drop [world_drop] | sim-verified (77.7 DPS) | yes | Wand of Mana Concentration (273043, +0.00 DPS) [world]; Greater Mystic Wand (217287, -0.12 DPS) [crafted]; Starfaller (13063, -1.25 DPS, sim-verified) [world_drop] |

**New at 30:** head: Filigreed Flame Circlet; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Darkspear Raider's Cloak; chest: Flame Gown; wrist: Phoenix Bindings; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Flame Leggings; feet: Fiery Slippers; finger1: Snake Hoop; finger2: Advisor's Ring; main_hand: Swinetusk Shank; off_hand: Rod of Molten Fire; ranged: Scorching Wand

No-known-source sample (15 of 301, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (undead, 000000000000000000-51520101230133031-000000000000000000)

Set DPS (verified): 159.2. Weights run: 3.9s. Verify run: 5.6s. 392 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.631 ± 0.019, crit=0.344 ± 0.021 per rating point (14 rating = 1%, 4.811 per %), hit=0.539 ± 0.034 per rating point (10 rating = 1%, 5.393 per %), spell_haste=not significant (0.519 ± 0.171), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=0.879 ± 0.182, mp5=0.934 ± 0.221

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 26.6 spell_power points (9.45 DPS) | yes | Whitemane's Chapeau (7720, -1.87 DPS, sim-verified) [dungeon]; Spellpower Goggles Xtreme (10502, -2.00 DPS) [crafted]; Electromagnetic Gigaflux Reactivator (9492, -2.01 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Necklace of Calisea (1714, -0.52 DPS) [dungeon]; Amberglow Talisman (10824, -1.16 DPS) [quest]; Triune Amulet (7722, -1.68 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Green Silken Shoulders (7057, -0.09 DPS) [crafted]; Bloodmage Mantle (7684, -0.19 DPS) [dungeon]; Mantle of Doan (7712, -1.56 DPS, sim-verified) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Raider's Cloak (272077, -0.65 DPS) [vendor]; Battle Healer's Cloak (19528, -0.67 DPS) [rep]; Cloak of Blight (6832, -1.57 DPS, sim-verified) [quest] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 36.0 spell_power points (12.77 DPS) | yes | Robe of the Magi (1716, -2.06 DPS) [world_drop]; Robe of Power (7054, -2.62 DPS) [crafted]; Cindercloth Robe (10042, -3.19 DPS) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (4.61 DPS) | yes | Spidertank Oilrag (9448, -0.43 DPS) [dungeon]; Radiant Silver Bracers (4545, -1.40 DPS) [quest]; Earthen Silk Cuffs (254019, -1.63 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 26.7 spell_power points (9.47 DPS) | yes | Fiery Handwraps (254025, -1.28 DPS, sim-verified) [crafted]; Crimson Silk Gloves (7064, -3.16 DPS) [crafted]; Red Mageweave Gloves (10018, -3.32 DPS) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | 24.3 spell_power points (8.63 DPS) | yes | Gilded Cord (254037, -2.13 DPS) [crafted]; Defiler's Cloth Girdle (20166, -2.77 DPS) [rep]; Scorching Sash (4117, -3.00 DPS, sim-verified) [quest] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | 22.8 spell_power points (8.10 DPS) | yes | Red Mageweave Pants (10009, -0.44 DPS) [crafted]; Black Mageweave Leggings (9999, -0.89 DPS) [crafted]; Stoneweaver Leggings (9407, -1.37 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 30.2 spell_power points (10.70 DPS) | yes | Fiery Slippers (254005, -2.11 DPS, sim-verified) [crafted]; Furen's Boots (13100, -4.14 DPS) [world_drop]; Thoughtcast Boots (10578, -4.87 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.2 spell_power points (4.67 DPS) | yes | Welken Ring (5011, -1.37 DPS) [world_drop]; Snake Hoop (6750, -1.42 DPS) [quest]; Electrocutioner Lagnut (9447, -1.42 DPS) [dungeon] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 11.8 spell_power points (4.19 DPS) | yes | Snake Hoop (6750, -0.94 DPS) [quest]; Electrocutioner Lagnut (9447, -0.94 DPS) [dungeon]; Welken Ring (5011, -2.46 DPS, sim-verified) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -12.59 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -13.42 DPS) [dungeon]; Celestial Stave (9517, -13.65 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ragefire Wand (7513) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jaina's Firestarter (13064, -1.79 DPS, sim-verified) [world_drop]; Twisted Nether Wand (249144, -2.13 DPS) [crafted]; Dancing Flame (6806, -2.31 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Icy Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Fiery Cord; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Staff of Jordan; ranged: Ragefire Wand

No-known-source sample (15 of 392, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (undead, 252010000000000000-51520101230133031-000000000000000000)

Set DPS (verified): 206.8. Weights run: 3.6s. Verify run: 5.7s. 504 eligible items had no known source.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.830 ± 0.031, crit=0.352 ± 0.024 per rating point (14 rating = 1%, 4.932 per %), hit=0.608 ± 0.040 per rating point (10 rating = 1%, 6.076 per %), spell_haste=not significant (0.840 ± 0.248), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.904 ± 0.283), mp5=not significant (0.534 ± 0.308)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Flame (3075) | World drop [world_drop] | 60.3 spell_power points (11.28 DPS) | yes | Red Mageweave Headband (10033, -4.62 DPS) [crafted]; Gemburst Circlet (10751, -4.85 DPS) [quest]; Dreamweave Circlet (10041, -6.61 DPS, sim-verified) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 21.0 spell_power points (3.92 DPS) | yes | Horizon Choker (13085, -1.07 DPS) [world_drop]; Triune Amulet (7722, -1.24 DPS) [dungeon]; Mindburst Medallion (11196, -2.90 DPS, sim-verified) [quest] |
| shoulder | Netherflame Shoulders (254053) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kentic Amice (11624, +0.00 DPS) [dungeon]; Rotgrip Mantle (17732, -0.21 DPS) [dungeon]; Flameseer Mantle (11310, -1.13 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cindercloth Cloak (14044, -0.89 DPS) [crafted]; Deep Woodlands Cloak (19121, -0.92 DPS) [quest]; Featherskin Cape (10843, -1.41 DPS) [world] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Embrace of the Wind Serpent (12462, +0.00 DPS) [world]; Acumen Robes (17775, -0.47 DPS) [quest]; Robes of Insight (940, -0.71 DPS) [world_drop] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 23.3 spell_power points (4.36 DPS) | yes | Nethergeld Cuffs (254061, -1.12 DPS) [crafted]; Bloodband Bracers (11469, -1.52 DPS) [quest]; Aristocratic Cuffs (12546, -1.80 DPS, sim-verified) [dungeon] |
| hands | Fiery Gloves (254099) | Tailoring [crafted] | 34.5 spell_power points (6.45 DPS) | yes | Runecloth Gloves (13863, -1.29 DPS) [crafted]; Cindercloth Gloves (14043, -1.42 DPS) [crafted]; Dreamweave Gloves (10019, -3.75 DPS, sim-verified) [crafted] |
| waist | Fiery Waistcord (254085) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Fiery Cord (254041, -0.08 DPS) [crafted]; Ghostweave Cord (254073, -0.47 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 34.2 spell_power points (6.39 DPS) | yes | Cindercloth Leggings (12256, -0.31 DPS) [vendor]; Spellshock Leggings (9484, -0.54 DPS) [dungeon]; Senior Designer's Pantaloons (11841, -2.01 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 30.3 spell_power points (5.67 DPS) | yes | Fiery Sandals (254111, -0.53 DPS) [crafted]; Mistwalker Boots (10629, -1.39 DPS) [dungeon]; Vinerot Sandals (17748, -1.46 DPS) [dungeon] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | sim-verified (206.8 DPS) | yes | Cyclopean Band (11824, +0.00 DPS) [dungeon]; Philanthropist's Ring (281635, -0.31 DPS) [quest]; Advisor's Ring (19519, -0.31 DPS) [rep] |
| finger2 | Brainlash (6440) | Zul'Farrak [dungeon] | 17.0 spell_power points (3.17 DPS) | yes | Cyclopean Band (11824, +0.00 DPS) [dungeon]; Philanthropist's Ring (281635, -0.53 DPS) [quest]; Advisor's Ring (19519, -0.53 DPS) [rep] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -0.91 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -1.00 DPS) [crafted]; Rune of the Guard Captain (19120, -1.23 DPS) [quest] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.57 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -0.67 DPS) [crafted]; Rune of the Guard Captain (19120, -0.89 DPS) [quest] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spire of Hakkar (10844, -3.42 DPS) [world]; Kindling Stave (11750, -3.71 DPS) [dungeon]; Glowing Brightwood Staff (812, -5.12 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lesser Eternal Wand (249232, -3.49 DPS) [crafted]; Wand of Allistarj (13065, -4.18 DPS) [world_drop]; Noxious Shooter (17745, -6.36 DPS, sim-verified) [dungeon] |

**New at 50:** head: Eye of Flame; neck: Arcane Crystal Pendant; shoulder: Netherflame Shoulders; back: Spritecaster Cape; wrist: Netherflame Cuffs; hands: Fiery Gloves; waist: Fiery Waistcord; legs: Kilt of the Atal'ai Prophet; finger1: Eye of Adaegus; finger2: Brainlash; trinket2: Evonice's Landin' Pilla; ranged: Pyric Caduceus

No-known-source sample (15 of 504, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (undead, 252015000000200000-51520101230133031-003000000000000000)

Set DPS (verified): 374.3. Weights run: 4.0s. Verify run: 27.3s. 1293 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.901 ± 0.050, crit=0.802 ± 0.050 per rating point (14 rating = 1%, 11.234 per %), hit=1.142 ± 0.081 per rating point (10 rating = 1%, 11.418 per %), spell_haste=not significant (0.973 ± 0.620), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (2.248 ± 0.649), mp5=not significant (2.383 ± 0.608)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Flame (3075) | World drop [world_drop] | 74.5 spell_power points (24.95 DPS) | yes | Warlord's Silk Cowl (231601, -0.49 DPS) [pvp]; Devout Crown (16693, -4.07 DPS) [dungeon]; Crown of Caer Darrow (13986, -10.01 DPS, sim-verified) [quest] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | 39.6 spell_power points (13.26 DPS) | yes | Jeweled Amulet of Cainwyn (1443, -0.30 DPS) [world_drop]; Heart of the Fiend (13960, -0.46 DPS) [dungeon]; Star of Mystaria (12103, -0.70 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Silk Mantle (227104, +0.00 DPS) [pvp]; Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Shroud of the Nathrezim (18720, -5.51 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-verified (374.3 DPS) | yes | Featherskin Cape (10843, +0.00 DPS) [world]; Frostweaver Cape (12968, +0.00 DPS) [dungeon]; Spritecaster Cape (11623, -0.58 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Embrace of the Wind Serpent (12462, +0.00 DPS) [world]; Warlord's Silk Raiment (231596, +0.00 DPS) [pvp]; Alanna's Embrace (13314, -4.49 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 44.9 spell_power points (15.06 DPS) | yes | General's Silk Cuffs (16538, -1.79 DPS) [pvp]; Devout Bracers (16697, -3.77 DPS) [dungeon]; Wyrmthalak's Shackles (13958, -4.32 DPS, sim-verified) [quest] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | 52.2 spell_power points (17.50 DPS) | yes | Dreadmist Wraps (16705, -0.41 DPS) [dungeon]; Swarmtender's Gloves (275607, -0.76 DPS) [crafted]; Magister's Gloves (16684, -8.78 DPS, sim-verified) [dungeon] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sash of the Burning Heart (11807, +0.00 DPS) [dungeon]; General's Silk Sash (16537, +0.00 DPS) [pvp]; Thuzadin Sash (18740, -7.99 DPS, sim-verified) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Haunting Specter Leggings (11929, +0.00 DPS) [dungeon]; General's Silk Trousers (231595, +0.00 DPS) [pvp]; Magister's Leggings (16687, -7.37 DPS, sim-verified) [dungeon] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | 51.2 spell_power points (17.16 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Mistwalker Boots (10629, -1.19 DPS) [dungeon]; Mooncloth Boots (15802, -9.04 DPS, sim-verified) [crafted] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Postmaster's Seal (13392, -0.37 DPS) [dungeon]; Maiden's Circle (13001, -1.41 DPS) [world_drop]; Naglering (11669, -11.07 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Maiden's Circle (13001, -0.54 DPS) [world_drop]; Band of the Hierophant (13096, -1.90 DPS) [world_drop]; The Postmaster's Seal (13392, -4.35 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+13.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Shard of the Splithooves (10659, +0.00 DPS) [quest]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Serenity Field (272439, -3.34 DPS, sim-verified) [vendor] |
| main_hand | Elemental Mage Staff (944) | World drop [world_drop] | sim-verified (+4.6 DPS vs the runner-up, not corroborated against the finished set) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Staff of Jordan (873, -1.94 DPS) [world_drop]; Ironbark Staff (20220, -4.65 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pyric Caduceus (11748, -7.18 DPS) [dungeon]; Bonecreeper Stylus (13938, -7.44 DPS) [dungeon]; Brilliant Wand (249385, -7.65 DPS) [crafted] |

**New at 60:** neck: Lady Maye's Pendant; shoulder: Devout Mantle; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Devout Sandals; finger1: Channeler's Ring; finger2: Blessed Band of Light; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Elemental Mage Staff; ranged: Torch of Light

No-known-source sample (15 of 1293, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60, raid preset (undead, 050105200100000000-55500201230133031-003000000000000000)

Set DPS (verified): 799.7. Weights run: 4.7s. Verify run: 26.8s. 1293 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.626 ± 0.046, crit=0.979 ± 0.060 per rating point (14 rating = 1%, 13.701 per %), hit=1.486 ± 0.120 per rating point (10 rating = 1%, 14.859 per %), spell_haste=-6.080 ± 1.413, spell_penetration=not significant (0.000 ± 0.000), fire_power=0.866 ± 0.002, spirit=not significant (0.327 ± 0.203), mp5=not significant (0.256 ± 0.326)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Flame (3075) | World drop [world_drop] | 46.8 spell_power points (27.69 DPS) | yes | Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Cap of the Scarlet Savant (12752, -4.61 DPS, sim-verified) [quest] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-verified (799.7 DPS) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Arcane Crystal Pendant (20037, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 38.8 spell_power points (22.98 DPS) | yes | Champion's Silk Mantle (227104, -1.15 DPS) [pvp]; Warlord's Silk Amice (231594, -1.67 DPS) [pvp]; Shroud of the Nathrezim (18720, -5.04 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -7.91 DPS) [world]; Shroud of Arcane Mastery (22330, -8.36 DPS) [dungeon]; Cloak of Flames (3475, -13.42 DPS, sim-verified) [world_drop] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Raiment (231596, -1.13 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -8.43 DPS) [pvp]; Alanna's Embrace (13314, -20.86 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 29.3 spell_power points (17.34 DPS) | yes | Netherflame Cuffs (254065, -7.11 DPS) [crafted]; Runecloth Cuffs (254123, -7.12 DPS) [crafted]; Arena Wristguards (18709, -8.07 DPS) [world] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 42.7 spell_power points (25.25 DPS) | yes | General's Silk Handguards (16540, -3.86 DPS) [vendor]; General's Silk Gauntlets (231599, -3.86 DPS) [vendor]; Hands of Power (13253, -4.74 DPS, sim-verified) [dungeon] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 34.0 spell_power points (20.13 DPS) | yes | Ban'thok Sash (11662, -0.16 DPS) [dungeon]; Defiler's Cloth Girdle (20163, -1.52 DPS) [rep]; Clutch of Andros (13956, -4.30 DPS) [dungeon] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 43.2 spell_power points (25.54 DPS) | yes | Legionnaire's Silk Legguards (227107, +0.00 DPS) [pvp]; General's Silk Trousers (231595, +0.00 DPS) [pvp]; Skyshroud Leggings (13170, -2.46 DPS) [dungeon] |
| feet | First Sergeant's Dreadweave Boots (220909) | PvP reward · Horde [vendor] | 28.5 spell_power points (16.86 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.58 DPS) [dungeon]; Dragonrider Boots (18102, -3.07 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -1.38 DPS) [vendor]; Don Mauricio's Band of Domination (22433, -2.28 DPS) [dungeon]; Naglering (11669, -19.76 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -0.73 DPS) [vendor]; Don Mauricio's Band of Domination (22433, -1.64 DPS) [dungeon]; Maiden's Circle (13001, -1.85 DPS) [world_drop] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -4.77 DPS, sim-verified) [vendor]; Burst of Knowledge (11832, -5.92 DPS) [dungeon]; Rune of the Guard Captain (19120, -6.86 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burst of Knowledge (11832, -0.81 DPS) [dungeon]; Rune of the Guard Captain (19120, -1.76 DPS) [quest]; Uther's Strength (11302, -4.36 DPS) [world_drop] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -1.18 DPS) [pvp]; Elemental Mage Staff (944, -12.18 DPS, sim-verified) [world_drop]; Staff of Jordan (873, -15.81 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Pyric Caduceus (11748, -0.71 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -1.38 DPS) [dungeon] |

**New at 60:** neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Magician's Cord; legs: Outrider's Silk Leggings; feet: First Sergeant's Dreadweave Boots; finger1: Don Julio's Band; finger2: Blessed Band of Light; trinket1: Weakness Analyzer; trinket2: Frozen Heart of the Mountain; main_hand: Ironbark Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1293, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

