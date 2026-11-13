# Leveling BiS: Fire

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 000000000000000000-51400001000000000-000000000000000000)

Set DPS (verified): 40.0. Weights run: 2.6s. Verify run: 2.5s. 198 eligible items had no known source.

Race sweep: gnome 40.5 ± 0.1, human 40.0 ± 0.1, high-order-skyborne 39.8 ± 0.1 (dps); gnome beats the guide race human by 0.6 (1.43%) ± 0.1, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.908 ± 0.016, crit=0.155 ± 0.007 per rating point (14 rating = 1%, 2.175 per %), hit=0.310 ± 0.016 per rating point (10 rating = 1%, 3.096 per %), spell_haste=not significant (0.696 ± 0.200), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.264 ± 0.079), mp5=1.437 ± 0.214

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Flame Circlet (253953) | Tailoring [crafted] | 11.0 spell_power points (1.55 DPS) | yes | Pristine Circlet (253949, -0.70 DPS) [crafted]; Shadow Goggles (4373, -0.81 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.40 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.5 spell_power points (0.07 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.2 spell_power points (1.85 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.79 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -1.25 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -1.29 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pearl-clasped Cloak (5542, -0.16 DPS, sim-verified) [crafted]; Sanguine Cape (14376, -0.20 DPS) [world_drop]; Caretaker's Cape (20428, -0.22 DPS) [rep] |
| chest | Filigreed Flame Gown (253905) | Tailoring [crafted] | 14.3 spell_power points (2.02 DPS) | yes | Filigreed Pristine Gown (253901, -0.57 DPS, sim-verified) [crafted]; Gray Woolen Robe (2585, -0.82 DPS) [crafted]; Mystic's Wrap (14369, -1.12 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 8.9 spell_power points (1.25 DPS) | yes | Bright Bracers (3647, -0.71 DPS, sim-verified) [world_drop]; Repurposed Hair Band (281256, -0.92 DPS) [quest]; Mystic's Bracelets (14366, -0.99 DPS) [world_drop] |
| hands | Flame Gloves (253917) | Tailoring [crafted] | 9.7 spell_power points (1.37 DPS) | yes | Phoenix Gloves (4331, -0.34 DPS, sim-verified) [crafted]; Serpent Gloves (5970, -0.38 DPS) [dungeon]; Pristine Gloves (253913, -0.42 DPS) [crafted] |
| waist | Flame Sash (253929) | Tailoring [crafted] | 10.6 spell_power points (1.50 DPS) | yes | Pristine Sash (253925, -0.43 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.47 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.55 DPS) [crafted] |
| legs | Filigreed Flame Leggings (253941) | Tailoring [crafted] | 17.5 spell_power points (2.46 DPS) | yes | Phoenix Pants (4317, -0.66 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.70 DPS) [crafted]; Darkweave Breeches (12987, -1.35 DPS) [world_drop] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.6 spell_power points (1.50 DPS) | yes | Flame Boots (253893, -0.52 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.69 DPS) [crafted]; Red Woolen Boots (4313, -0.75 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 7.9 spell_power points (1.11 DPS) | yes | Lavishly Jeweled Ring (1156, -0.34 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.73 DPS) [world_drop]; Band of Purification (12996, -0.89 DPS) [world_drop] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.8 spell_power points (0.96 DPS) | yes | Lavishly Jeweled Ring (1156, -0.19 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.58 DPS) [world_drop]; Band of Purification (12996, -0.74 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | 31.1 spell_power points (4.38 DPS) | yes | Emberstone Staff (5201, -1.00 DPS) [dungeon]; Staff of the Blessed Seer (2271, -1.13 DPS, sim-verified) [dungeon]; Staff of Westfall (2042, -1.27 DPS) [quest] |
| off_hand | Antipodean Rod (2879) | World drop [world_drop] | 6.0 spell_power points (0.84 DPS) | yes | Mathystra Channeling Orb (281291, -0.24 DPS, sim-verified) [quest]; Tear of Grief (5611, -0.32 DPS) [quest]; Dwarven Tome (279898, -0.39 DPS) [quest] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller (12984, -0.30 DPS) [world_drop]; Deepblaze (279896, -1.51 DPS) [quest]; Sizzle Stick (8071, -1.58 DPS) [quest] |

**New at 20:** head: Flame Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Flame Gown; wrist: Mindthrust Bracers; hands: Flame Gloves; waist: Flame Sash; legs: Filigreed Flame Leggings; feet: Spidersilk Boots; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Evocator's Blade; off_hand: Antipodean Rod; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 198, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 8350 The 1 Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (human, 000000000000000000-54500001230100000-000000000000000000)

Set DPS (verified): 80.9. Weights run: 2.8s. Verify run: 5.2s. 326 eligible items had no known source.

Race sweep: gnome 82.0 ± 0.1, human 80.9 ± 0.1, high-order-skyborne 79.6 ± 0.1 (dps); gnome beats the guide race human by 1.1 (1.41%) ± 0.2, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.938 ± 0.015, crit=0.227 ± 0.012 per rating point (14 rating = 1%, 3.184 per %), hit=0.408 ± 0.021 per rating point (10 rating = 1%, 4.077 per %), spell_haste=0.452 ± 0.103, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (-0.107 ± 0.050), mp5=0.815 ± 0.168

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Flame Circlet (253979) | Tailoring [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Nightsky Cowl (4039, -0.68 DPS) [world_drop]; Flame Circlet (253953, -0.75 DPS) [crafted]; Enchanter's Cowl (4322, -1.01 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 5.6 spell_power points (1.40 DPS) | yes | Seal of the Expedition (271920, -0.23 DPS) [rep]; Crystal Starfire Medallion (5003, -0.47 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.47 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 17.4 spell_power points (4.35 DPS) | yes | Death Speaker Mantle (6685, -0.84 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -1.00 DPS) [world_drop]; Invoker's Mantle (215365, -1.43 DPS) [crafted] |
| back | Expeditionary's Cape (271923) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Resilient Cape (14400, -0.62 DPS, sim-verified) [world_drop]; Hillman's Cloak (3719, -0.63 DPS) [crafted]; Vine Pruner's Cloak (279835, -0.63 DPS) [quest] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Flame Gown (253965, -0.17 DPS) [crafted]; Mechbuilder's Overalls (9508, -0.25 DPS) [dungeon]; Green Silk Armor (7065, -1.51 DPS, sim-verified) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (3.24 DPS) | yes | Spidertank Oilrag (9448, -0.68 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -1.37 DPS) [world_drop]; Mindthrust Bracers (1974, -1.46 DPS) [dungeon] |
| hands | Flame Gloves (253917) | Tailoring [crafted] | 9.8 spell_power points (2.45 DPS) | yes | Phoenix Gloves (4331, -0.20 DPS) [crafted]; Truefaith Gloves (7049, -0.50 DPS) [crafted]; Town Clerk's Mittens (270029, -1.06 DPS, sim-verified) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 13.8 spell_power points (3.44 DPS) | yes | Belt of Arugal (6392, -0.50 DPS) [dungeon]; Invoker's Cord (215366, -0.53 DPS) [crafted]; Crimson Silk Belt (7055, -1.34 DPS, sim-verified) [crafted] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | 20.6 spell_power points (5.13 DPS) | yes | Filigreed Flame Leggings (253941, -0.85 DPS, sim-verified) [crafted]; Phoenix Pants (4317, -1.70 DPS) [crafted]; Pristine Leggings (253987, -1.75 DPS) [crafted] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 19.6 spell_power points (4.88 DPS) | yes | Gilded Slippers (254001, -1.64 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.76 DPS) [dungeon]; Spidersilk Boots (4320, -2.20 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 12.6 spell_power points (3.13 DPS) | yes | Lorekeeper's Ring (19525, -0.98 DPS) [rep]; Ring of Power Regulation (273030, -1.03 DPS) [dungeon]; Minor Channeling Ring (1449, -1.42 DPS) [quest] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 8.6 spell_power points (2.15 DPS) | yes | Lorekeeper's Ring (19525, -0.00 DPS) [rep]; Ring of Power Regulation (273030, -0.05 DPS) [dungeon]; Minor Channeling Ring (1449, -0.44 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.00 DPS) [dungeon]; Tendonscraper Dagger (273053, -1.00 DPS) [world]; Glimmering Staff (249392, -4.02 DPS, sim-verified) [crafted] |
| off_hand | Rod of Molten Fire (2565) | World drop [world_drop] | 13.0 spell_power points (3.24 DPS) | yes | Orb of Mystic Insight (249394, -0.66 DPS, sim-verified) [crafted]; Antipodean Rod (2879, -1.75 DPS) [world_drop]; Satyr's Rod (15962, -1.84 DPS) [world_drop] |
| ranged | Scorching Wand (5213) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Starfaller (13063, +0.00 DPS) [world_drop]; Greater Mystic Wand (217287, -0.22 DPS) [crafted]; Wand of Mana Concentration (273043, -0.82 DPS, sim-verified) [world] |

**New at 30:** head: Filigreed Flame Circlet; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Expeditionary's Cape; chest: Beguiler Robes; wrist: Phoenix Bindings; waist: Highlander's Cloth Girdle; legs: Flame Leggings; feet: Fiery Slippers; finger1: Snake Hoop; finger2: Runemender's Seal; main_hand: Swinetusk Shank; off_hand: Rod of Molten Fire; ranged: Scorching Wand

No-known-source sample (15 of 326, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (human, 000000000000000000-54500001230133031-000000000000000000)

Set DPS (verified): 162.1. Weights run: 3.0s. Verify run: 4.6s. 456 eligible items had no known source.

Race sweep: human 162.1 ± 0.3, gnome 161.6 ± 0.3, high-order-skyborne 157.6 ± 0.3 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.621 ± 0.019, crit=0.357 ± 0.021 per rating point (14 rating = 1%, 4.995 per %), hit=0.520 ± 0.032 per rating point (10 rating = 1%, 5.199 per %), spell_haste=0.620 ± 0.151, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=0.803 ± 0.185, mp5=0.929 ± 0.218

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 25.8 spell_power points (9.20 DPS) | yes | Miner's Hat of the Deep (9429, -1.15 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -1.71 DPS) [crafted]; Whitemane's Chapeau (7720, -3.10 DPS, sim-verified) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Necklace of Calisea (1714, -0.71 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.06 DPS) [quest]; Triune Amulet (7722, -1.52 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Green Silken Shoulders (7057, -0.09 DPS) [crafted]; Bloodmage Mantle (7684, -0.17 DPS) [dungeon]; Mantle of Doan (7712, -1.61 DPS, sim-verified) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Expeditionary's Cape (271922, -0.63 DPS) [rep]; Caretaker's Cape (19532, -0.64 DPS) [rep]; Glowing Thresher Cape (6901, -0.85 DPS) [dungeon] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 34.8 spell_power points (12.42 DPS) | yes | Robe of the Magi (1716, -1.81 DPS) [world_drop]; Robe of Power (7054, -2.48 DPS) [crafted]; Cindercloth Robe (10042, -2.79 DPS) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (4.64 DPS) | yes | Spidertank Oilrag (9448, -0.43 DPS) [dungeon]; Arcane Runed Bracers (4744, -1.43 DPS) [quest]; Earthen Silk Cuffs (254019, -1.78 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 26.1 spell_power points (9.31 DPS) | yes | Fiery Handwraps (254025, -1.58 DPS, sim-verified) [crafted]; Crimson Silk Gloves (7064, -2.99 DPS) [crafted]; Red Mageweave Gloves (10018, -3.17 DPS) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | 23.8 spell_power points (8.48 DPS) | yes | Gilded Cord (254037, -2.14 DPS) [crafted]; Highlander's Cloth Girdle (20098, -2.60 DPS) [rep]; Scorching Sash (4117, -3.12 DPS, sim-verified) [quest] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | 22.4 spell_power points (7.98 DPS) | yes | Red Mageweave Pants (10009, -0.32 DPS) [crafted]; Black Mageweave Leggings (9999, -1.11 DPS) [crafted]; Stoneweaver Leggings (9407, -1.79 DPS, sim-verified) [dungeon] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 20.6 spell_power points (7.33 DPS) | yes | Thoughtcast Boots (10578, -1.83 DPS) [dungeon]; Gilded Slippers (254001, -2.14 DPS) [crafted]; Furen's Boots (13100, -3.42 DPS, sim-verified) [world_drop] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 16.0 spell_power points (5.70 DPS) | yes | Runemender's Seal (273031, -1.29 DPS) [world]; Lorekeeper's Ring (19524, -1.49 DPS) [rep] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.1 spell_power points (4.67 DPS) | yes | Lorekeeper's Ring (19524, -0.47 DPS) [rep]; Runemender's Seal (273031, -1.79 DPS, sim-verified) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -12.38 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -13.15 DPS) [dungeon]; Celestial Stave (9517, -13.30 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ragefire Wand (7513) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plaguerot Sprig (10766, -1.51 DPS) [dungeon]; Jaina's Firestarter (13064, -1.74 DPS, sim-verified) [world_drop]; Twisted Nether Wand (249144, -2.05 DPS) [crafted] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Icy Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Fiery Cord; finger2: Philanthropist's Ring; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Staff of Jordan; ranged: Ragefire Wand

No-known-source sample (15 of 456, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (human, 230002000000000000-54500001230133031-003000000000000000)

Set DPS (verified): 259.7. Weights run: 3.0s. Verify run: 4.8s. 598 eligible items had no known source.

Race sweep: human 259.7 ± 0.4, gnome 257.1 ± 0.4, high-order-skyborne 250.7 ± 0.4 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.655 ± 0.030, crit=0.283 ± 0.024 per rating point (14 rating = 1%, 3.968 per %), hit=0.525 ± 0.040 per rating point (10 rating = 1%, 5.246 per %), spell_haste=not significant (0.645 ± 0.290), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.690 ± 0.289), mp5=not significant (0.791 ± 0.326)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Flame (3075) | World drop [world_drop] | 56.5 spell_power points (13.67 DPS) | yes | Red Mageweave Headband (10033, -5.90 DPS) [crafted]; Gemburst Circlet (10751, -7.00 DPS) [quest]; Dreamweave Circlet (10041, -7.31 DPS, sim-verified) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 19.9 spell_power points (4.83 DPS) | yes | Mindburst Medallion (11196, -1.59 DPS) [quest]; Triune Amulet (7722, -1.72 DPS) [dungeon]; Talvash's Enhancing Necklace (7673, -1.76 DPS) [quest] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 26.7 spell_power points (6.46 DPS) | yes | Netherflame Shoulders (254053, -0.07 DPS) [crafted]; Rotgrip Mantle (17732, -0.45 DPS) [dungeon]; Flameseer Mantle (11310, -1.39 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cindercloth Cloak (14044, -0.93 DPS) [crafted]; Runecloth Cloak (13860, -1.90 DPS) [crafted]; Cloak of Fire (14134, -1.96 DPS) [crafted] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 38.3 spell_power points (9.27 DPS) | yes | Cindercloth Vest (14042, -1.62 DPS) [crafted]; Robe of the Magi (1716, -2.15 DPS) [world_drop]; Dreamweave Vest (10021, -4.24 DPS, sim-verified) [crafted] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 21.0 spell_power points (5.10 DPS) | yes | Aristocratic Cuffs (12546, -1.75 DPS) [dungeon]; Phoenix Bindings (210781, -1.95 DPS) [crafted]; Nethergeld Cuffs (254061, -2.22 DPS, sim-verified) [crafted] |
| hands | Fiery Gloves (254099) | Tailoring [crafted] | 31.1 spell_power points (7.53 DPS) | yes | Cindercloth Gloves (14043, -1.57 DPS) [crafted]; Runecloth Gloves (13863, -1.69 DPS) [crafted]; Dreamweave Gloves (10019, -3.30 DPS, sim-verified) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Fiery Cord (254041, -0.10 DPS) [crafted]; Ghostweave Cord (254073, -0.53 DPS) [crafted]; Fiery Waistcord (254085, -1.32 DPS, sim-verified) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | 29.6 spell_power points (7.16 DPS) | yes | Cindercloth Leggings (12256, -0.05 DPS) [vendor]; Senior Designer's Pantaloons (11841, -0.13 DPS) [dungeon]; Kilt of the Atal'ai Prophet (10807, -0.56 DPS) [dungeon] |
| feet | Fiery Sandals (254111) | Tailoring [crafted] | 25.9 spell_power points (6.27 DPS) | yes | Cindercloth Boots (10044, -1.19 DPS) [crafted]; Fiery Slippers (254005, -1.34 DPS) [crafted]; Mistwalker Boots (10629, -2.65 DPS, sim-verified) [dungeon] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 21.0 spell_power points (5.09 DPS) | yes | Snake Hoop (6750, -1.35 DPS) [quest]; Lorekeeper's Ring (19523, -1.41 DPS) [rep]; Philanthropist's Ring (281635, -1.87 DPS) [quest] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 16.3 spell_power points (3.96 DPS) | yes | Snake Hoop (6750, -0.22 DPS) [quest]; Lorekeeper's Ring (19523, -0.29 DPS) [rep]; Philanthropist's Ring (281635, -0.74 DPS) [quest] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| trinket2 | Fire Ruby (20036) | Destroy Morphaz [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -1.38 DPS) [world_drop]; Uther's Strength (11302, -1.94 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -2.04 DPS, sim-verified) [quest] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Spire of Hakkar (10844, +0.00 DPS) [world]; Staff of Jordan (873, -7.17 DPS, sim-verified) [world_drop] |
| off_hand | Truesilver Conduit (249455) | Enchanting [crafted] | 15.2 spell_power points (3.69 DPS) | yes | Celestial Orb (7515, -0.07 DPS) [quest]; Orb of the Forgotten Seer (7685, -0.30 DPS) [dungeon]; Enthralled Sphere (11625, -0.31 DPS) [dungeon] |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Noxious Shooter (17745, -3.19 DPS) [dungeon]; Wand of Allistarj (13065, -4.88 DPS) [world_drop]; Lesser Eternal Wand (249232, -6.25 DPS, sim-verified) [crafted] |

**New at 50:** head: Eye of Flame; neck: Arcane Crystal Pendant; shoulder: Kentic Amice; back: Spritecaster Cape; chest: Embrace of the Wind Serpent; wrist: Netherflame Cuffs; hands: Fiery Gloves; waist: Ban'thok Sash; legs: Spellshock Leggings; feet: Fiery Sandals; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Thunderbrew's Boot Flask; trinket2: Fire Ruby; main_hand: Blade of Eternal Darkness; off_hand: Truesilver Conduit; ranged: Pyric Caduceus

No-known-source sample (15 of 598, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (human, 230105200100000000-55500201230133031-003000000000000000)

Set DPS (verified): 400.8. Weights run: 3.5s. Verify run: 27.0s. 1456 eligible items had no known source.

Race sweep: human 400.8 ± 0.7, gnome 396.9 ± 0.7, high-order-skyborne 387.8 ± 0.7 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.544 ± 0.054, crit=0.476 ± 0.035 per rating point (14 rating = 1%, 6.659 per %), hit=1.266 ± 0.074 per rating point (10 rating = 1%, 12.658 per %), spell_haste=not significant (-3.573 ± 0.971), spell_penetration=not significant (0.000 ± 0.000), fire_power=0.840 ± 0.003, spirit=not significant (1.493 ± 0.657), mp5=not significant (0.821 ± 0.668)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (+2.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Field Marshal's Coronet (231604, -1.62 DPS) [pvp]; Eye of Flame (3075, -2.52 DPS, sim-verified) [world_drop]; Crimson Felt Hat (18727, -3.30 DPS) [dungeon] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 30.7 spell_power points (6.89 DPS) | yes | Star of Mystaria (12103, -0.00 DPS) [dungeon]; Dragon's Eye (10829, -0.97 DPS) [dungeon]; Nacreous Shell Necklace (22403, -1.06 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Argent Shoulders (19059, +0.00 DPS) [crafted]; Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Shroud of the Nathrezim (18720, -12.98 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, -1.15 DPS) [quest]; Spritecaster Cape (11623, -1.52 DPS) [dungeon]; Featherskin Cape (10843, -1.71 DPS) [world] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Robes of the Exalted (13346, +0.00 DPS) [dungeon]; Field Marshal's Silk Vestments (231603, +0.00 DPS) [pvp]; Embrace of the Wind Serpent (12462, -11.04 DPS, sim-verified) [world] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 36.8 spell_power points (8.26 DPS) | yes | Wyrmthalak's Shackles (13958, -1.66 DPS) [quest]; Marshal's Silk Bracers (16438, -2.59 DPS) [pvp]; Devout Bracers (16697, -3.03 DPS) [dungeon] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Hands of Power (13253, -13.43 DPS, sim-verified) [dungeon] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sash of the Burning Heart (11807, +0.00 DPS) [dungeon]; Marshal's Silk Sash (16439, +0.00 DPS) [pvp]; Thuzadin Sash (18740, -12.14 DPS, sim-verified) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Magister's Leggings (16687, +0.00 DPS) [dungeon]; Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; The Postmaster's Trousers (13389, -13.66 DPS, sim-verified) [dungeon] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mistwalker Boots (10629, +0.00 DPS) [dungeon]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; The Postmaster's Treads (13391, -14.99 DPS, sim-verified) [dungeon] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Maiden's Circle (13001, -0.86 DPS) [world_drop]; The Postmaster's Seal (13392, -1.71 DPS, sim-verified) [dungeon]; Brainlash (6440, -2.40 DPS) [dungeon] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Maiden's Circle (13001, -0.32 DPS) [world_drop]; Brainlash (6440, -1.86 DPS) [dungeon]; The Postmaster's Seal (13392, -11.63 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Eye of the Beast (13968, -12.08 DPS, sim-verified) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cannonball Runner (13382, +0.00 DPS) [dungeon]; Ankh of Life (1713, -0.02 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.69 DPS) [quest] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Headmaster's Charge (13937, +0.00 DPS) [dungeon]; Ironbark Staff (20069, -0.10 DPS) [rep]; Blade of Eternal Darkness (17780, -3.19 DPS, sim-verified) [dungeon] |
| off_hand | Tome of Fiery Arcana (19311) | Stormpike Guard [rep] | 31.1 spell_power points (6.97 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Therazane's Touch (19315, -0.46 DPS) [rep]; Penelope's Rose (15805, -0.71 DPS) [quest] |
| ranged | Oblivion's Touch (18761) | Scholomance: Death Knight Darkreaver [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, +0.00 DPS) [dungeon]; Brilliant Wand (249385, -1.03 DPS) [crafted] |

**New at 60:** head: Spellweaver's Turban; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Devout Sandals; finger1: Blessed Band of Light; finger2: Channeler's Ring; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Sageblade; off_hand: Tome of Fiery Arcana; ranged: Oblivion's Touch

No-known-source sample (15 of 1456, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (human, 050105200100000000-55500201230133031-003000000000000000)

Set DPS (verified): 821.3. Weights run: 4.0s. Verify run: 22.3s. 1456 eligible items had no known source.

Race sweep: gnome 823.1 ± 0.9, human 821.3 ± 0.9, high-order-skyborne 801.1 ± 0.9 (dps); gnome beats the guide race human by 1.8 (0.22%) ± 1.2, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.582 ± 0.049, crit=0.954 ± 0.059 per rating point (14 rating = 1%, 13.362 per %), hit=1.603 ± 0.112 per rating point (10 rating = 1%, 16.025 per %), spell_haste=not significant (-5.374 ± 1.346), spell_penetration=not significant (0.000 ± 0.000), fire_power=0.867 ± 0.002, spirit=not significant (-0.165 ± 0.244), mp5=not significant (-0.040 ± 0.355)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 49.0 spell_power points (29.39 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -2.51 DPS) [pvp]; Eye of Flame (3075, -5.12 DPS, sim-verified) [world_drop] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 29.7 spell_power points (17.80 DPS) | yes | Star of Mystaria (12103, -4.38 DPS, sim-verified) [dungeon]; Lady Maye's Pendant (14558, -4.57 DPS) [world_drop]; Arcane Crystal Pendant (20037, -6.11 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 37.9 spell_power points (22.74 DPS) | yes | Lieutenant Commander's Silk Mantle (227102, -1.90 DPS) [pvp]; Field Marshal's Silk Spaulders (231602, -2.52 DPS) [pvp]; Shroud of the Nathrezim (18720, -5.82 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -8.61 DPS) [world]; Shroud of Arcane Mastery (22330, -9.81 DPS, sim-verified) [dungeon]; Mageflame Cloak (13007, -11.08 DPS) [world_drop] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Vestments (231603, -2.45 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -9.65 DPS) [pvp]; Alanna's Embrace (13314, -17.57 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 26.7 spell_power points (15.98 DPS) | yes | Runecloth Cuffs (254123, -5.90 DPS) [crafted]; Manacle Cuffs (11962, -6.55 DPS) [quest]; Netherflame Cuffs (254065, -6.78 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 41.5 spell_power points (24.91 DPS) | yes | Marshal's Silk Gloves (16440, -4.53 DPS) [vendor]; Marshal's Silk Gauntlets (231608, -4.53 DPS) [vendor]; Hands of Power (13253, -10.72 DPS, sim-verified) [dungeon] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 33.8 spell_power points (20.29 DPS) | yes | Magician's Cord (272393, -0.52 DPS) [vendor]; Highlander's Cloth Girdle (20047, -1.79 DPS) [rep]; Clutch of Andros (13956, -4.05 DPS) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 64.5 spell_power points (38.70 DPS) | yes | Marshal's Silk Leggings (231605, -5.72 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -12.16 DPS) [pvp]; Skyshroud Leggings (13170, -16.52 DPS) [dungeon] |
| feet | Kayser's Boots of Precision (22231) | Blackrock Spire: War Master Voone [dungeon] | 30.5 spell_power points (18.29 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Sergeant Major's Dreadweave Boots (220891, -0.74 DPS) [vendor]; Dragonrider Boots (18102, -1.91 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -3.01 DPS) [dungeon]; Channeler's Ring (272406, -3.12 DPS) [vendor]; Naglering (11669, -17.45 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -1.00 DPS) [dungeon]; Channeler's Ring (272406, -1.11 DPS) [vendor]; Maiden's Circle (13001, -2.37 DPS) [world_drop] |
| trinket1 | Eye of the Beast (13968) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fire Ruby (20036, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Smolderweb's Eye (13213, -10.44 DPS, sim-verified) [dungeon] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+2.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -2.14 DPS) [crafted]; Smolderweb's Eye (13213, -2.61 DPS, sim-verified) [dungeon]; Fire Ruby (20036, -3.51 DPS) [quest] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -1.40 DPS) [pvp]; Elemental Mage Staff (944, -16.51 DPS, sim-verified) [world_drop]; Staff of Jordan (873, -16.65 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Oblivion's Touch (18761, +0.00 DPS) [dungeon]; Pyric Caduceus (11748, -1.14 DPS) [dungeon] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; legs: Sentinel's Silk Leggings; feet: Kayser's Boots of Precision; finger1: Don Julio's Band; finger2: Blessed Band of Light; trinket1: Eye of the Beast; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1456, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (undead, 000000000000000000-51400001000000000-000000000000000000)

Set DPS (verified): 40.4. Weights run: 2.6s. Verify run: 2.4s. 192 eligible items had no known source.

Race sweep: undead 40.4 ± 0.1, orc 39.3 ± 0.1, troll 39.0 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.908 ± 0.016, crit=0.155 ± 0.007 per rating point (14 rating = 1%, 2.175 per %), hit=0.310 ± 0.016 per rating point (10 rating = 1%, 3.096 per %), spell_haste=not significant (0.696 ± 0.200), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.264 ± 0.079), mp5=1.437 ± 0.214

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Flame Circlet (253953) | Tailoring [crafted] | 11.0 spell_power points (1.55 DPS) | yes | Pristine Circlet (253949, -0.70 DPS) [crafted]; Shadow Goggles (4373, -1.06 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.40 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.8 spell_power points (0.11 DPS) | yes | Scholarly Pendant (277203, -0.04 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.2 spell_power points (1.85 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.38 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -1.25 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -1.29 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pearl-clasped Cloak (5542, -0.05 DPS) [crafted]; Sanguine Cape (14376, -0.20 DPS) [world_drop]; Battle Healer's Cloak (20427, -0.22 DPS) [rep] |
| chest | Filigreed Flame Gown (253905) | Tailoring [crafted] | 14.3 spell_power points (2.02 DPS) | yes | Filigreed Pristine Gown (253901, -0.59 DPS, sim-verified) [crafted]; Gray Woolen Robe (2585, -0.82 DPS) [crafted]; Mystic's Wrap (14369, -1.12 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 8.9 spell_power points (1.25 DPS) | yes | Bright Bracers (3647, -0.62 DPS, sim-verified) [world_drop]; Tabitha's Cuffs (251486, -0.82 DPS) [quest]; Owlbeard Bracers (16981, -0.85 DPS) [quest] |
| hands | Flame Gloves (253917) | Tailoring [crafted] | 9.7 spell_power points (1.37 DPS) | yes | Phoenix Gloves (4331, -0.29 DPS, sim-verified) [crafted]; Serpent Gloves (5970, -0.38 DPS) [dungeon]; Pristine Gloves (253913, -0.42 DPS) [crafted] |
| waist | Flame Sash (253929) | Tailoring [crafted] | 10.6 spell_power points (1.50 DPS) | yes | Pristine Sash (253925, -0.44 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.47 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.55 DPS) [crafted] |
| legs | Filigreed Flame Leggings (253941) | Tailoring [crafted] | 17.5 spell_power points (2.46 DPS) | yes | Phoenix Pants (4317, -0.68 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.70 DPS) [crafted]; Darkweave Breeches (12987, -1.35 DPS) [world_drop] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.6 spell_power points (1.50 DPS) | yes | Flame Boots (253893, -0.27 DPS) [crafted]; Pristine Boots (253889, -0.69 DPS) [crafted]; Red Woolen Boots (4313, -0.75 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 7.9 spell_power points (1.11 DPS) | yes | Loop of Sacrifice (281673, -0.47 DPS) [quest]; Volcanic Rock Ring (12053, -0.73 DPS) [world_drop]; Band of Purification (12996, -0.89 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.4 spell_power points (0.77 DPS) | yes | Loop of Sacrifice (281673, -0.13 DPS) [quest]; Volcanic Rock Ring (12053, -0.38 DPS) [world_drop]; Band of Purification (12996, -0.54 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Staff of the Blessed Seer (2271, -0.07 DPS) [dungeon]; Gnarled Necromancer's Staff (251534, -0.35 DPS) [quest]; Crescent Staff (6505, -0.75 DPS, sim-verified) [quest] |
| off_hand | Antipodean Rod (2879) | World drop [world_drop] | 6.0 spell_power points (0.84 DPS) | yes | Seer's Fine Stein (7608, -0.42 DPS) [world_drop]; Spellbinder Orb (15926, -0.46 DPS) [dungeon]; Tork Wrench (11855, -0.59 DPS) [quest] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller (12984, -0.27 DPS) [world_drop]; Sizzle Stick (8071, -1.43 DPS) [quest]; Greater Magic Wand (11288, -1.49 DPS) [crafted] |

**New at 20:** head: Flame Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Flame Gown; wrist: Mindthrust Bracers; hands: Flame Gloves; waist: Flame Sash; legs: Filigreed Flame Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Evocator's Blade; off_hand: Antipodean Rod; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 192, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 8350 The 1 Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 23173 Abomination Skin Leggings

### Band 30 (undead, 000000000000000000-54500001230100000-000000000000000000)

Set DPS (verified): 81.5. Weights run: 2.8s. Verify run: 5.4s. 318 eligible items had no known source.

Race sweep: undead 81.5 ± 0.1, orc 80.0 ± 0.1, troll 79.0 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.938 ± 0.015, crit=0.227 ± 0.012 per rating point (14 rating = 1%, 3.184 per %), hit=0.408 ± 0.021 per rating point (10 rating = 1%, 4.077 per %), spell_haste=0.452 ± 0.103, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (-0.107 ± 0.050), mp5=0.815 ± 0.168

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Flame Circlet (253979) | Tailoring [crafted] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Nightsky Cowl (4039, -0.68 DPS) [world_drop]; Flame Circlet (253953, -0.75 DPS) [crafted]; Enchanter's Cowl (4322, -0.87 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 5.6 spell_power points (1.40 DPS) | yes | Darkspear Warding Pendant (272075, -0.23 DPS) [vendor]; Crystal Starfire Medallion (5003, -0.47 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.47 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 17.4 spell_power points (4.35 DPS) | yes | Death Speaker Mantle (6685, -0.76 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -1.00 DPS) [world_drop]; Invoker's Mantle (215365, -1.43 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, -0.63 DPS) [crafted]; Windsong Drape (15468, -0.63 DPS) [quest]; Resilient Cape (14400, -0.72 DPS, sim-verified) [world_drop] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Flame Gown (253965, -0.17 DPS) [crafted]; Mechbuilder's Overalls (9508, -0.25 DPS) [dungeon]; Green Silk Armor (7065, -0.90 DPS, sim-verified) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (3.24 DPS) | yes | Spidertank Oilrag (9448, -0.56 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -1.37 DPS) [world_drop]; Mindthrust Bracers (1974, -1.46 DPS) [dungeon] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 10.7 spell_power points (2.67 DPS) | yes | Flame Gloves (253917, -0.22 DPS) [crafted]; Phoenix Gloves (4331, -0.42 DPS) [crafted]; Truefaith Gloves (7049, -0.72 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 13.8 spell_power points (3.44 DPS) | yes | Belt of Arugal (6392, -0.50 DPS) [dungeon]; Invoker's Cord (215366, -0.53 DPS) [crafted]; Crimson Silk Belt (7055, -1.06 DPS, sim-verified) [crafted] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | 20.6 spell_power points (5.13 DPS) | yes | Filigreed Flame Leggings (253941, -1.17 DPS, sim-verified) [crafted]; Phoenix Pants (4317, -1.70 DPS) [crafted]; Pristine Leggings (253987, -1.75 DPS) [crafted] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 19.6 spell_power points (4.88 DPS) | yes | Gilded Slippers (254001, -1.61 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.76 DPS) [dungeon]; Spidersilk Boots (4320, -2.20 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 12.6 spell_power points (3.13 DPS) | yes | Advisor's Ring (19521, -0.98 DPS) [rep]; Ring of Power Regulation (273030, -1.03 DPS) [dungeon] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 8.6 spell_power points (2.15 DPS) | yes | Advisor's Ring (19521, -0.00 DPS) [rep]; Ring of Power Regulation (273030, -0.05 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.00 DPS) [dungeon]; Tendonscraper Dagger (273053, -1.00 DPS) [world]; Glimmering Staff (249392, -4.21 DPS, sim-verified) [crafted] |
| off_hand | Rod of Molten Fire (2565) | World drop [world_drop] | 13.0 spell_power points (3.24 DPS) | yes | Orb of Mystic Insight (249394, -0.59 DPS, sim-verified) [crafted]; Witch's Finger (16887, -1.37 DPS) [quest]; Antipodean Rod (2879, -1.75 DPS) [world_drop] |
| ranged | Scorching Wand (5213) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Starfaller (13063, +0.00 DPS) [world_drop]; Greater Mystic Wand (217287, -0.22 DPS) [crafted]; Wand of Mana Concentration (273043, -0.71 DPS, sim-verified) [world] |

**New at 30:** head: Filigreed Flame Circlet; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Darkspear Raider's Cloak; chest: Beguiler Robes; wrist: Phoenix Bindings; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Flame Leggings; feet: Fiery Slippers; finger1: Snake Hoop; finger2: Runemender's Seal; main_hand: Swinetusk Shank; off_hand: Rod of Molten Fire; ranged: Scorching Wand

No-known-source sample (15 of 318, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 8350 The 1 Ring; 9362 Brilliant Gold Ring

### Band 40 (undead, 000000000000000000-54500001230133031-000000000000000000)

Set DPS (verified): 160.4. Weights run: 3.0s. Verify run: 4.4s. 448 eligible items had no known source.

Race sweep: undead 160.4 ± 0.3, orc 159.0 ± 0.3, troll 156.9 ± 0.3 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.621 ± 0.019, crit=0.357 ± 0.021 per rating point (14 rating = 1%, 4.995 per %), hit=0.520 ± 0.032 per rating point (10 rating = 1%, 5.199 per %), spell_haste=0.620 ± 0.151, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=0.803 ± 0.185, mp5=0.929 ± 0.218

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 25.8 spell_power points (9.20 DPS) | yes | Miner's Hat of the Deep (9429, -1.15 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -1.71 DPS) [crafted]; Whitemane's Chapeau (7720, -2.52 DPS, sim-verified) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Necklace of Calisea (1714, -0.71 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.06 DPS) [quest]; Triune Amulet (7722, -1.91 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Green Silken Shoulders (7057, -0.09 DPS) [crafted]; Bloodmage Mantle (7684, -0.17 DPS) [dungeon]; Mantle of Doan (7712, -1.32 DPS, sim-verified) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Raider's Cloak (272077, -0.63 DPS) [vendor]; Battle Healer's Cloak (19528, -0.64 DPS) [rep]; Cloak of Blight (6832, -2.05 DPS, sim-verified) [quest] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 34.8 spell_power points (12.42 DPS) | yes | Robe of the Magi (1716, -1.81 DPS) [world_drop]; Robe of Power (7054, -2.48 DPS) [crafted]; Cindercloth Robe (10042, -2.79 DPS) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (4.64 DPS) | yes | Spidertank Oilrag (9448, -0.43 DPS) [dungeon]; Radiant Silver Bracers (4545, -1.44 DPS) [quest]; Earthen Silk Cuffs (254019, -1.78 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 26.1 spell_power points (9.31 DPS) | yes | Fiery Handwraps (254025, -1.24 DPS, sim-verified) [crafted]; Crimson Silk Gloves (7064, -2.99 DPS) [crafted]; Red Mageweave Gloves (10018, -3.17 DPS) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | 23.8 spell_power points (8.48 DPS) | yes | Gilded Cord (254037, -2.14 DPS) [crafted]; Defiler's Cloth Girdle (20166, -2.60 DPS) [rep]; Scorching Sash (4117, -3.56 DPS, sim-verified) [quest] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | 22.4 spell_power points (7.98 DPS) | yes | Red Mageweave Pants (10009, -0.32 DPS) [crafted]; Black Mageweave Leggings (9999, -1.11 DPS) [crafted]; Stoneweaver Leggings (9407, -1.76 DPS, sim-verified) [dungeon] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 20.6 spell_power points (7.33 DPS) | yes | Thoughtcast Boots (10578, -1.83 DPS) [dungeon]; Gilded Slippers (254001, -2.14 DPS) [crafted]; Furen's Boots (13100, -3.86 DPS, sim-verified) [world_drop] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 16.0 spell_power points (5.70 DPS) | yes | Runemender's Seal (273031, -1.29 DPS) [world]; Advisor's Ring (19520, -1.49 DPS) [rep] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.1 spell_power points (4.67 DPS) | yes | Advisor's Ring (19520, -0.47 DPS) [rep]; Runemender's Seal (273031, -1.72 DPS, sim-verified) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -12.08 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -13.15 DPS) [dungeon]; Celestial Stave (9517, -13.30 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ragefire Wand (7513) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plaguerot Sprig (10766, -1.49 DPS) [dungeon]; Jaina's Firestarter (13064, -1.66 DPS, sim-verified) [world_drop]; Twisted Nether Wand (249144, -2.09 DPS) [crafted] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Icy Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Fiery Cord; finger2: Philanthropist's Ring; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Staff of Jordan; ranged: Ragefire Wand

No-known-source sample (15 of 448, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (undead, 230002000000000000-54500001230133031-003000000000000000)

Set DPS (verified): 254.8. Weights run: 3.0s. Verify run: 4.7s. 590 eligible items had no known source.

Race sweep: undead 254.8 ± 0.4, orc 251.4 ± 0.4, troll 248.4 ± 0.4 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.699 ± 0.032, crit=0.537 ± 0.032 per rating point (14 rating = 1%, 7.523 per %), hit=0.770 ± 0.053 per rating point (10 rating = 1%, 7.697 per %), spell_haste=not significant (0.824 ± 0.281), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.850 ± 0.340), mp5=not significant (0.781 ± 0.390)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Flame (3075) | World drop [world_drop] | 58.5 spell_power points (21.46 DPS) | yes | Dreamweave Circlet (10041, -8.03 DPS, sim-verified) [crafted]; Red Mageweave Headband (10033, -9.36 DPS) [crafted]; Gemburst Circlet (10751, -10.02 DPS) [quest] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 20.2 spell_power points (7.41 DPS) | yes | Triune Amulet (7722, -2.29 DPS) [dungeon]; Horizon Choker (13085, -2.57 DPS) [world_drop]; Mindburst Medallion (11196, -3.57 DPS, sim-verified) [quest] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 28.2 spell_power points (10.34 DPS) | yes | Netherflame Shoulders (254053, -0.18 DPS) [crafted]; Rotgrip Mantle (17732, -0.96 DPS) [dungeon]; Flameseer Mantle (11310, -2.09 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -1.84 DPS) [quest]; Cindercloth Cloak (14044, -2.61 DPS, sim-verified) [crafted]; Featherskin Cape (10843, -2.77 DPS) [world] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 43.5 spell_power points (15.95 DPS) | yes | Dreamweave Vest (10021, -2.65 DPS, sim-verified) [crafted]; Cindercloth Vest (14042, -3.72 DPS) [crafted]; Ghostweave Vest (14141, -4.12 DPS) [crafted] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 22.1 spell_power points (8.12 DPS) | yes | Nethergeld Cuffs (254061, -2.23 DPS, sim-verified) [crafted]; Aristocratic Cuffs (12546, -2.81 DPS) [dungeon]; Bloodband Bracers (11469, -3.05 DPS) [quest] |
| hands | Fiery Gloves (254099) | Tailoring [crafted] | 32.8 spell_power points (12.03 DPS) | yes | Dreamweave Gloves (10019, -2.22 DPS) [crafted]; Cindercloth Gloves (14043, -2.36 DPS) [crafted]; Runecloth Gloves (13863, -2.51 DPS) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 26.7 spell_power points (9.79 DPS) | yes | Fiery Waistcord (254085, -0.51 DPS) [crafted]; Fiery Cord (254041, -0.73 DPS) [crafted]; Ghostweave Cord (254073, -1.22 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 32.2 spell_power points (11.81 DPS) | yes | Cindercloth Leggings (12256, -0.16 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -0.48 DPS) [dungeon]; Spellshock Leggings (9484, -0.81 DPS) [dungeon] |
| feet | Fiery Sandals (254111) | Tailoring [crafted] | 26.3 spell_power points (9.65 DPS) | yes | First Sergeant's Dreadweave Boots (220909, -1.58 DPS) [vendor]; Fiery Slippers (254005, -1.83 DPS) [crafted]; Mistwalker Boots (10629, -5.24 DPS, sim-verified) [dungeon] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 22.2 spell_power points (8.16 DPS) | yes | Snake Hoop (6750, -1.98 DPS) [quest]; Advisor's Ring (19519, -2.61 DPS) [rep]; Eye of Adaegus (5266, -2.88 DPS) [world_drop] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 17.3 spell_power points (6.34 DPS) | yes | Snake Hoop (6750, -0.17 DPS) [quest]; Advisor's Ring (19519, -0.79 DPS) [rep]; Eye of Adaegus (5266, -1.06 DPS) [world_drop] |
| trinket1 | Fire Ruby (20036) | Destroy Morphaz [quest] | sim-verified (+5.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, -1.40 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -2.02 DPS) [quest]; Frozen Heart of the Mountain (249469, -2.60 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Staff of Jordan (873, -11.32 DPS, sim-verified) [world_drop] |
| off_hand | Truesilver Conduit (249455) | Enchanting [crafted] | 15.6 spell_power points (5.72 DPS) | yes | Celestial Orb (7515, -0.18 DPS) [quest]; Enthralled Sphere (11625, -0.28 DPS) [dungeon]; Flask of Forest Mojo (19115, -0.37 DPS) [quest] |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Noxious Shooter (17745, -4.63 DPS) [dungeon]; Ragefire Wand (7513, -6.31 DPS) [quest]; Lesser Eternal Wand (249232, -7.31 DPS, sim-verified) [crafted] |

**New at 50:** head: Eye of Flame; neck: Arcane Crystal Pendant; shoulder: Kentic Amice; back: Spritecaster Cape; chest: Embrace of the Wind Serpent; wrist: Netherflame Cuffs; hands: Fiery Gloves; waist: Ban'thok Sash; legs: Senior Designer's Pantaloons; feet: Fiery Sandals; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Fire Ruby; trinket2: Uther's Strength; main_hand: Blade of Eternal Darkness; off_hand: Truesilver Conduit; ranged: Pyric Caduceus

No-known-source sample (15 of 590, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (undead, 230105200100000000-55500201230133031-003000000000000000)

Set DPS (verified): 393.6. Weights run: 3.5s. Verify run: 26.1s. 1450 eligible items had no known source.

Race sweep: undead 393.6 ± 0.7, orc 389.4 ± 0.7, troll 386.1 ± 0.7 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.544 ± 0.054, crit=0.476 ± 0.035 per rating point (14 rating = 1%, 6.659 per %), hit=1.266 ± 0.074 per rating point (10 rating = 1%, 12.658 per %), spell_haste=not significant (-3.573 ± 0.971), spell_penetration=not significant (0.000 ± 0.000), fire_power=0.840 ± 0.003, spirit=not significant (1.493 ± 0.657), mp5=not significant (0.821 ± 0.668)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 65.1 spell_power points (14.60 DPS) | yes | Warlord's Silk Cowl (231601, -1.62 DPS) [pvp]; Eye of Flame (3075, -1.93 DPS) [world_drop]; Crimson Felt Hat (18727, -3.30 DPS) [dungeon] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 30.7 spell_power points (6.89 DPS) | yes | Dragon's Eye (10829, -0.97 DPS) [dungeon]; Nacreous Shell Necklace (22403, -1.06 DPS) [dungeon]; Star of Mystaria (12103, -1.29 DPS, sim-verified) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Argent Shoulders (19059, +0.00 DPS) [crafted]; Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Shroud of the Nathrezim (18720, -11.34 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -1.52 DPS) [dungeon]; Featherskin Cape (10843, -1.71 DPS) [world]; Shroud of the Exile (15421, -3.52 DPS, sim-verified) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Embrace of the Wind Serpent (12462, +0.00 DPS) [world]; Warlord's Silk Raiment (231596, +0.00 DPS) [pvp]; Robes of the Exalted (13346, -17.28 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 36.8 spell_power points (8.26 DPS) | yes | Wyrmthalak's Shackles (13958, -1.66 DPS) [quest]; General's Silk Cuffs (16538, -2.59 DPS) [pvp]; Devout Bracers (16697, -3.03 DPS) [dungeon] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Hands of Power (13253, -10.56 DPS, sim-verified) [dungeon] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sash of the Burning Heart (11807, +0.00 DPS) [dungeon]; General's Silk Sash (16537, +0.00 DPS) [pvp]; Thuzadin Sash (18740, -11.04 DPS, sim-verified) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Magister's Leggings (16687, +0.00 DPS) [dungeon]; General's Silk Trousers (231595, +0.00 DPS) [pvp]; The Postmaster's Trousers (13389, -14.18 DPS, sim-verified) [dungeon] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mistwalker Boots (10629, +0.00 DPS) [dungeon]; General's Silk Boots (231597, +0.00 DPS) [pvp]; The Postmaster's Treads (13391, -14.95 DPS, sim-verified) [dungeon] |
| finger1 | The Postmaster's Seal (13392) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, +0.00 DPS) [vendor]; Maiden's Circle (13001, -0.40 DPS) [world_drop]; Brainlash (6440, -1.94 DPS) [dungeon] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-verified (+10.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Blessed Band of Light (272407, +0.00 DPS) [vendor]; Maiden's Circle (13001, -0.32 DPS) [world_drop]; Brainlash (6440, -1.86 DPS) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Eye of the Beast (13968, -9.70 DPS, sim-verified) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Second Wind (11819, +0.00 DPS) [dungeon]; Ankh of Life (1713, -0.02 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.69 DPS) [quest] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Headmaster's Charge (13937, +0.00 DPS) [dungeon]; Ironbark Staff (20220, -0.10 DPS) [rep]; Blade of Eternal Darkness (17780, -2.23 DPS, sim-verified) [dungeon] |
| off_hand | Tome of Fiery Arcana (19311) | Frostwolf Clan [rep] | 31.1 spell_power points (6.97 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; Therazane's Touch (19315, -0.46 DPS) [rep]; Penelope's Rose (15805, -0.71 DPS) [quest] |
| ranged | Bonecreeper Stylus (13938) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ritssyn's Wand of Bad Mojo (22408, -0.16 DPS) [dungeon]; Oblivion's Touch (18761, -0.25 DPS) [dungeon]; Brilliant Wand (249385, -1.34 DPS) [crafted] |

**New at 60:** head: Spellweaver's Turban; neck: Jeweled Amulet of Cainwyn; shoulder: Devout Mantle; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Devout Sandals; finger1: The Postmaster's Seal; finger2: Channeler's Ring; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Sageblade; off_hand: Tome of Fiery Arcana; ranged: Bonecreeper Stylus

No-known-source sample (15 of 1450, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60, raid preset (undead, 050105200100000000-55500201230133031-003000000000000000)

Set DPS (verified): 808.7. Weights run: 4.0s. Verify run: 22.2s. 1450 eligible items had no known source.

Race sweep: undead 808.7 ± 0.8, orc 803.7 ± 0.8, troll 797.9 ± 0.8 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.582 ± 0.049, crit=0.954 ± 0.059 per rating point (14 rating = 1%, 13.362 per %), hit=1.603 ± 0.112 per rating point (10 rating = 1%, 16.025 per %), spell_haste=not significant (-5.374 ± 1.346), spell_penetration=not significant (0.000 ± 0.000), fire_power=0.867 ± 0.002, spirit=not significant (-0.165 ± 0.244), mp5=not significant (-0.040 ± 0.355)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 49.0 spell_power points (29.39 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -2.51 DPS) [pvp]; Eye of Flame (3075, -5.96 DPS, sim-verified) [world_drop] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 29.7 spell_power points (17.80 DPS) | yes | Star of Mystaria (12103, -2.86 DPS) [dungeon]; Lady Maye's Pendant (14558, -4.57 DPS) [world_drop]; Arcane Crystal Pendant (20037, -6.11 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 37.9 spell_power points (22.74 DPS) | yes | Champion's Silk Mantle (227104, -1.90 DPS) [pvp]; Warlord's Silk Amice (231594, -2.52 DPS) [pvp]; Shroud of the Nathrezim (18720, -5.32 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -8.61 DPS) [world]; Shroud of Arcane Mastery (22330, -9.14 DPS, sim-verified) [dungeon]; Mageflame Cloak (13007, -11.08 DPS) [world_drop] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Raiment (231596, -2.45 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -9.65 DPS) [pvp]; Alanna's Embrace (13314, -21.12 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 26.7 spell_power points (15.98 DPS) | yes | Runecloth Cuffs (254123, -5.90 DPS) [crafted]; Manacle Cuffs (11962, -6.55 DPS) [quest]; Netherflame Cuffs (254065, -6.78 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 41.5 spell_power points (24.91 DPS) | yes | General's Silk Handguards (16540, -4.53 DPS) [vendor]; General's Silk Gauntlets (231599, -4.53 DPS) [vendor]; Hands of Power (13253, -7.83 DPS, sim-verified) [dungeon] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 33.8 spell_power points (20.29 DPS) | yes | Magician's Cord (272393, -0.52 DPS) [vendor]; Defiler's Cloth Girdle (20163, -1.79 DPS) [rep]; Clutch of Andros (13956, -4.05 DPS) [dungeon] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 39.1 spell_power points (23.42 DPS) | yes | Legionnaire's Silk Legguards (227107, +0.00 DPS) [pvp]; General's Silk Trousers (231595, +0.00 DPS) [pvp]; Skyshroud Leggings (13170, -1.24 DPS) [dungeon] |
| feet | Kayser's Boots of Precision (22231) | Blackrock Spire: War Master Voone [dungeon] | 30.5 spell_power points (18.29 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; First Sergeant's Dreadweave Boots (220909, -0.74 DPS) [vendor]; Dragonrider Boots (18102, -1.91 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -3.01 DPS) [dungeon]; Channeler's Ring (272406, -3.12 DPS) [vendor]; Naglering (11669, -18.30 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -1.11 DPS) [vendor]; Don Mauricio's Band of Domination (22433, -1.92 DPS, sim-verified) [dungeon]; Maiden's Circle (13001, -2.37 DPS) [world_drop] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -2.14 DPS) [crafted]; Eye of the Beast (13968, -2.19 DPS, sim-verified) [quest]; Fire Ruby (20036, -3.51 DPS) [quest] |
| trinket2 | Smolderweb's Eye (13213) | Blackrock Spire: Mother Smolderweb [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -0.96 DPS) [crafted]; Fire Ruby (20036, -2.33 DPS) [quest]; Rune of the Guard Captain (19120, -2.88 DPS) [quest] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -1.40 DPS) [pvp]; Elemental Mage Staff (944, -13.44 DPS, sim-verified) [world_drop]; Staff of Jordan (873, -16.65 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bonecreeper Stylus (13938) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Oblivion's Touch (18761, -0.50 DPS) [dungeon]; Brilliant Wand (249385, -0.53 DPS) [crafted]; Pyric Caduceus (11748, -1.68 DPS) [dungeon] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; legs: Outrider's Silk Leggings; feet: Kayser's Boots of Precision; finger1: Don Julio's Band; finger2: Blessed Band of Light; trinket1: Weakness Analyzer; trinket2: Smolderweb's Eye; main_hand: Ironbark Staff; ranged: Bonecreeper Stylus

No-known-source sample (15 of 1450, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

