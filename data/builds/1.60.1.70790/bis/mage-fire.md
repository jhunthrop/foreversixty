# Leveling BiS: Fire

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 000000000000000000-51400001000000000-000000000000000000)

Set DPS (verified): 38.6. Weights run: 3.9s. Verify run: 2.2s. 181 eligible items had no known source.

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
| finger1 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 7.1 spell_power points (1.00 DPS) | yes | Lavishly Jeweled Ring (1156, -0.14 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.57 DPS) [world_drop]; The 1 Ring (8350, -0.84 DPS) [world] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 7.0 spell_power points (0.99 DPS) | yes | Volcanic Rock Ring (12053, -0.56 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.57 DPS, sim-verified) [dungeon]; The 1 Ring (8350, -0.83 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | 30.8 spell_power points (4.33 DPS) | yes | Emberstone Staff (5201, -0.95 DPS) [dungeon]; Staff of the Blessed Seer (2271, -1.01 DPS, sim-verified) [dungeon]; Buzzer Blade (2169, -1.23 DPS) [dungeon] |
| off_hand | Antipodean Rod (2879) | World drop [world_drop] | 6.0 spell_power points (0.84 DPS) | yes | Tear of Grief (5611, -0.19 DPS) [quest]; Seer's Fine Stein (7608, -0.39 DPS) [world_drop]; Mathystra Channeling Orb (281291, -0.58 DPS, sim-verified) [quest] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller (12984, -0.45 DPS, sim-verified) [world_drop]; Deepblaze (279896, -1.22 DPS) [quest]; Sizzle Stick (8071, -1.22 DPS) [quest] |

**New at 20:** head: Flame Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Flame Gown; wrist: Mindthrust Bracers; hands: Flame Gloves; waist: Flame Sash; legs: Filigreed Flame Leggings; feet: Spidersilk Boots; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Evocator's Blade; off_hand: Antipodean Rod; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 181, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 30 (human, 000000000000000000-54500001230100000-000000000000000000)

Set DPS (verified): 78.7. Weights run: 2.6s. Verify run: 3.9s. 309 eligible items had no known source.

Race sweep: gnome 79.2 ± 0.1, human 78.7 ± 0.1, high-order-skyborne 77.0 ± 0.1 (dps); gnome beats the guide race human by 0.5 (0.66%) ± 0.2, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.884 ± 0.016, crit=0.220 ± 0.012 per rating point (14 rating = 1%, 3.085 per %), hit=0.395 ± 0.022 per rating point (10 rating = 1%, 3.955 per %), spell_haste=0.501 ± 0.104, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (-0.090 ± 0.097), mp5=not significant (0.514 ± 0.168)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Flame Circlet (253979) | Tailoring [crafted] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Enchanter's Cowl (4322, -0.61 DPS, sim-verified) [crafted]; Holy Shroud (2721, -0.74 DPS) [world_drop]; Flame Circlet (253953, -0.74 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 5.3 spell_power points (1.31 DPS) | yes | Seal of the Expedition (271920, -0.22 DPS) [rep]; Crystal Starfire Medallion (5003, -0.44 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.44 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 17.0 spell_power points (4.19 DPS) | yes | Death Speaker Mantle (6685, -0.58 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.99 DPS) [world_drop]; Invoker's Mantle (215365, -1.37 DPS) [crafted] |
| back | Expeditionary's Cape (271923) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Resilient Cape (14400, -0.44 DPS) [world_drop]; Hillman's Cloak (3719, -0.51 DPS) [crafted]; Vine Pruner's Cloak (279835, -0.51 DPS) [quest] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Flame Gown (253965, -0.10 DPS) [crafted]; Mechbuilder's Overalls (9508, -0.25 DPS) [dungeon]; Green Silk Armor (7065, -1.05 DPS, sim-verified) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (3.21 DPS) | yes | Spidertank Oilrag (9448, -0.82 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -1.47 DPS) [world_drop]; Mindthrust Bracers (1974, -1.74 DPS) [dungeon] |
| hands | Flame Gloves (253917) | Tailoring [crafted] | 9.7 spell_power points (2.39 DPS) | yes | Phoenix Gloves (4331, -0.16 DPS) [crafted]; Truefaith Gloves (7049, -0.49 DPS) [crafted]; Town Clerk's Mittens (270029, -0.74 DPS, sim-verified) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 13.7 spell_power points (3.38 DPS) | yes | Belt of Arugal (6392, -0.49 DPS) [dungeon]; Invoker's Cord (215366, -0.55 DPS) [crafted]; Crimson Silk Belt (7055, -1.00 DPS, sim-verified) [crafted] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | 20.2 spell_power points (4.99 DPS) | yes | Filigreed Flame Leggings (253941, -0.73 DPS, sim-verified) [crafted]; Phoenix Pants (4317, -1.64 DPS) [crafted]; Pristine Leggings (253987, -1.73 DPS) [crafted] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 19.2 spell_power points (4.74 DPS) | yes | Gilded Slippers (254001, -1.57 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.76 DPS) [dungeon]; Spidersilk Boots (4320, -2.14 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 10.4 spell_power points (2.58 DPS) | yes | Ring of Power Regulation (273030, -0.71 DPS) [dungeon]; Minor Channeling Ring (1449, -0.90 DPS) [quest]; Lavishly Jeweled Ring (1156, -1.26 DPS) [dungeon] |
| finger2 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 8.0 spell_power points (1.98 DPS) | yes | Ring of Power Regulation (273030, -0.12 DPS) [dungeon]; Minor Channeling Ring (1449, -0.31 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.67 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -0.99 DPS) [dungeon]; Tendonscraper Dagger (273053, -0.99 DPS) [world]; Glimmering Staff (249392, -4.47 DPS, sim-verified) [crafted] |
| off_hand | Rod of Molten Fire (2565) | World drop [world_drop] | 13.0 spell_power points (3.21 DPS) | yes | Orb of Mystic Insight (249394, -0.79 DPS, sim-verified) [crafted]; Antipodean Rod (2879, -1.73 DPS) [world_drop]; Satyr's Rod (15962, -1.90 DPS) [world_drop] |
| ranged | Scorching Wand (5213) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wand of Mana Concentration (273043, +0.00 DPS) [world]; Greater Mystic Wand (217287, -0.22 DPS) [crafted]; Starfaller (13063, -1.17 DPS, sim-verified) [world_drop] |

**New at 30:** head: Filigreed Flame Circlet; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Expeditionary's Cape; chest: Beguiler Robes; wrist: Phoenix Bindings; waist: Highlander's Cloth Girdle; legs: Flame Leggings; feet: Fiery Slippers; finger1: Snake Hoop; finger2: Lorekeeper's Ring; main_hand: Swinetusk Shank; off_hand: Rod of Molten Fire; ranged: Scorching Wand

No-known-source sample (15 of 309, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (human, 000000000000000000-54500001230133031-000000000000000000)

Set DPS (verified): 164.2. Weights run: 2.9s. Verify run: 4.0s. 400 eligible items had no known source.

Race sweep: human 164.2 ± 0.3, gnome 163.6 ± 0.3, high-order-skyborne 159.1 ± 0.3 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.539 ± 0.017, crit=0.322 ± 0.017 per rating point (14 rating = 1%, 4.511 per %), hit=0.553 ± 0.030 per rating point (10 rating = 1%, 5.533 per %), spell_haste=0.577 ± 0.103, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=0.644 ± 0.120, mp5=0.841 ± 0.165

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 23.0 spell_power points (8.76 DPS) | yes | Whitemane's Chapeau (7720, -1.08 DPS) [dungeon]; Electromagnetic Gigaflux Reactivator (9492, -1.85 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -2.81 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Necklace of Calisea (1714, -1.15 DPS) [dungeon]; Triune Amulet (7722, -1.83 DPS, sim-verified) [dungeon]; Amberglow Talisman (10824, -1.86 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Green Silken Shoulders (7057, -0.03 DPS) [crafted]; Bloodmage Mantle (7684, -0.06 DPS) [dungeon]; Mantle of Doan (7712, -1.62 DPS, sim-verified) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Caretaker's Cape (19532, -0.63 DPS) [rep]; Long Silken Cloak (4326, -0.83 DPS) [crafted]; Guardian Cloak (5965, -0.83 DPS) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 31.9 spell_power points (12.15 DPS) | yes | Robe of the Magi (1716, -1.30 DPS) [world_drop]; Cindercloth Robe (10042, -1.86 DPS) [crafted]; Robe of Power (7054, -2.38 DPS) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (4.96 DPS) | yes | Spidertank Oilrag (9448, -0.56 DPS) [dungeon]; Arcane Runed Bracers (4744, -1.53 DPS) [quest]; Earthen Silk Cuffs (254019, -2.20 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 24.7 spell_power points (9.40 DPS) | yes | Fiery Handwraps (254025, -2.63 DPS) [crafted]; Crimson Silk Gloves (7064, -2.83 DPS) [crafted]; Red Mageweave Gloves (10018, -3.16 DPS) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | 22.2 spell_power points (8.46 DPS) | yes | Gilded Cord (254037, -2.29 DPS) [crafted]; Highlander's Cloth Girdle (20098, -2.30 DPS) [rep]; Scorching Sash (4117, -2.98 DPS, sim-verified) [quest] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Stoneweaver Leggings (9407, -0.54 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.02 DPS) [dungeon]; Flame Leggings (253991, -1.74 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 28.5 spell_power points (10.87 DPS) | yes | Fiery Slippers (254005, -1.92 DPS, sim-verified) [crafted]; Furen's Boots (13100, -5.50 DPS) [world_drop]; Thoughtcast Boots (10578, -5.77 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.7 spell_power points (4.84 DPS) | yes | Snake Hoop (6750, -1.53 DPS) [quest]; Ring of Power Regulation (273030, -1.59 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -1.79 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 11.5 spell_power points (4.39 DPS) | yes | Snake Hoop (6750, -1.08 DPS) [quest]; Ring of Power Regulation (273030, -1.14 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -1.34 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -13.37 DPS, sim-verified) [dungeon]; Celestial Stave (9517, -13.43 DPS) [quest]; Hypnotic Blade (7714, -13.45 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Ragefire Wand (7513) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jaina's Firestarter (13064, -1.85 DPS, sim-verified) [world_drop]; Twisted Nether Wand (249144, -2.03 DPS) [crafted]; Burning Sliver (5249, -2.15 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Icy Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Fiery Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Staff of Jordan; ranged: Ragefire Wand

No-known-source sample (15 of 400, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (human, 230002000000000000-54500001230133031-003000000000000000)

Set DPS (verified): 236.9. Weights run: 2.9s. Verify run: 4.0s. 512 eligible items had no known source.

Race sweep: human 236.9 ± 0.4, gnome 235.3 ± 0.4, high-order-skyborne 228.6 ± 0.4 (dps); the guide race human is best.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.846 ± 0.031, crit=0.394 ± 0.025 per rating point (14 rating = 1%, 5.511 per %), hit=0.684 ± 0.041 per rating point (10 rating = 1%, 6.836 per %), spell_haste=not significant (0.198 ± 0.244), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.740 ± 0.305), mp5=not significant (0.864 ± 0.323)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Flame (3075) | World drop [world_drop] | 58.9 spell_power points (12.24 DPS) | yes | Red Mageweave Headband (10033, -4.77 DPS) [crafted]; Gemburst Circlet (10751, -5.65 DPS) [quest]; Dreamweave Circlet (10041, -6.84 DPS, sim-verified) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 21.1 spell_power points (4.38 DPS) | yes | Horizon Choker (13085, -1.30 DPS) [world_drop]; Mindburst Medallion (11196, -1.31 DPS) [quest]; Glowing Eye of Mordresh (10769, -1.52 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kentic Amice (11624, +0.00 DPS) [dungeon]; Netherflame Shoulders (254053, -0.00 DPS) [crafted]; Flameseer Mantle (11310, -1.42 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cindercloth Cloak (14044, -0.78 DPS) [crafted]; Runecloth Cloak (13860, -1.61 DPS) [crafted]; Expeditionary's Cape (271921, -1.66 DPS) [rep] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Acumen Robes (17775, -0.01 DPS) [quest]; Embrace of the Wind Serpent (12462, -0.75 DPS, sim-verified) [world]; Robes of Insight (940, -0.77 DPS) [world_drop] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 22.6 spell_power points (4.71 DPS) | yes | Aristocratic Cuffs (12546, -1.14 DPS) [dungeon]; Nethergeld Cuffs (254061, -1.25 DPS) [crafted]; Bloodband Bracers (11469, -1.62 DPS) [quest] |
| hands | Fiery Gloves (254099) | Tailoring [crafted] | 33.4 spell_power points (6.94 DPS) | yes | Runecloth Gloves (13863, -1.48 DPS) [crafted]; Cindercloth Gloves (14043, -1.71 DPS) [crafted]; Dreamweave Gloves (10019, -2.89 DPS, sim-verified) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 28.1 spell_power points (5.85 DPS) | yes | Fiery Waistcord (254085, -0.32 DPS) [crafted]; Fiery Cord (254041, -0.61 DPS) [crafted]; Satyrmane Sash (17755, -1.18 DPS) [dungeon] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 31.6 spell_power points (6.56 DPS) | yes | Spellshock Leggings (9484, -0.02 DPS) [dungeon]; Senior Designer's Pantaloons (11841, -0.06 DPS) [dungeon]; Cindercloth Leggings (12256, -0.30 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 29.2 spell_power points (6.07 DPS) | yes | Fiery Sandals (254111, -0.33 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -1.40 DPS) [vendor]; Fiery Slippers (254005, -1.52 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cyclopean Band (11824, +0.00 DPS) [dungeon]; Philanthropist's Ring (281635, -0.26 DPS) [quest]; Eye of Adaegus (5266, -0.31 DPS) [world_drop] |
| finger2 | Brainlash (6440) | Zul'Farrak [dungeon] | 16.4 spell_power points (3.41 DPS) | yes | Cyclopean Band (11824, +0.00 DPS) [dungeon]; Philanthropist's Ring (281635, -0.45 DPS) [quest]; Eye of Adaegus (5266, -0.51 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+4.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Uther's Strength (11302, -0.03 DPS) [world_drop] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, -2.78 DPS) [world_drop]; Spire of Hakkar (10844, -4.31 DPS) [world]; Kindling Stave (11750, -6.28 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Noxious Shooter (17745, -2.84 DPS) [dungeon]; Wand of Allistarj (13065, -4.33 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.60 DPS, sim-verified) [crafted] |

**New at 50:** head: Eye of Flame; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Spritecaster Cape; wrist: Netherflame Cuffs; hands: Fiery Gloves; waist: Ban'thok Sash; legs: Kilt of the Atal'ai Prophet; finger1: Lorekeeper's Ring; finger2: Brainlash; trinket1: Frozen Heart of the Mountain; trinket2: Thunderbrew's Boot Flask; ranged: Pyric Caduceus

No-known-source sample (15 of 512, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (human, 230105200100000000-55500201230133031-003000000000000000)

Set DPS (verified): 406.1. Weights run: 3.2s. Verify run: 26.1s. 1299 eligible items had no known source.

Race sweep: human 406.1 ± 0.8, gnome 402.0 ± 0.8, high-order-skyborne 393.2 ± 0.8 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.599 ± 0.052, crit=0.463 ± 0.037 per rating point (14 rating = 1%, 6.476 per %), hit=1.239 ± 0.073 per rating point (10 rating = 1%, 12.390 per %), spell_haste=not significant (-2.546 ± 0.966), spell_penetration=not significant (0.000 ± 0.000), fire_power=0.847 ± 0.003, spirit=not significant (0.338 ± 0.669), mp5=not significant (-0.531 ± 0.679)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Flame (3075) | World drop [world_drop] | 45.8 spell_power points (10.43 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.25 DPS) [pvp]; Crimson Felt Hat (18727, -3.35 DPS, sim-verified) [dungeon] |
| neck | Star of Mystaria (12103) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon]; Arcane Crystal Pendant (20037, -0.20 DPS) [quest]; Archlight Talisman (15856, -0.57 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burial Shawl (18681, +0.00 DPS) [dungeon]; Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -14.01 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -3.09 DPS) [dungeon]; Shroud of Arcane Mastery (22330, -3.24 DPS) [dungeon]; Crystalline Threaded Cape (20697, -4.78 DPS, sim-verified) [world] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Vestments (231603, -0.45 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -3.26 DPS) [pvp]; Alanna's Embrace (13314, -10.61 DPS, sim-verified) [dungeon] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (+12.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [rep] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Hands of Power (13253, -16.08 DPS, sim-verified) [dungeon] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Highlander's Cloth Girdle (20047, +0.00 DPS) [rep]; Magician's Cord (272393, -14.03 DPS, sim-verified) [vendor] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 50.9 spell_power points (11.60 DPS) | yes | Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -2.64 DPS) [pvp]; Skyshroud Leggings (13170, -2.77 DPS) [dungeon] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Omnicast Boots (11822, +0.00 DPS) [dungeon]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -18.86 DPS, sim-verified) [dungeon] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.35 DPS) [quest]; Channeler's Ring (272406, -0.37 DPS) [vendor]; Naglering (11669, -12.21 DPS, sim-verified) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.15 DPS) [quest]; Channeler's Ring (272406, -0.17 DPS) [vendor]; Flaming Band (12926, -0.89 DPS) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -0.19 DPS) [crafted]; Uther's Strength (11302, -1.37 DPS) [world_drop]; Second Wind (11819, -12.16 DPS, sim-verified) [dungeon] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -1.82 DPS) [vendor]; Frozen Heart of the Mountain (249469, -2.47 DPS) [crafted]; Uther's Strength (11302, -3.64 DPS) [world_drop] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -0.51 DPS) [pvp]; Elemental Mage Staff (944, -2.22 DPS) [world_drop]; Ironbark Staff (20069, -10.79 DPS, sim-verified) [rep] |
| off_hand | Tome of Fiery Arcana (19311) | Stormpike Guard [rep] | 31.3 spell_power points (7.14 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Therazane's Touch (19315, -0.53 DPS) [rep]; Grand Marshal's Tome of Restoration (234590, -1.81 DPS) [pvp] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ritssyn's Wand of Bad Mojo (22408, -3.50 DPS) [dungeon]; Brilliant Wand (249385, -4.25 DPS) [crafted]; Bonecreeper Stylus (13938, -5.29 DPS, sim-verified) [dungeon] |

**New at 60:** neck: Star of Mystaria; shoulder: Devout Mantle; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Sentinel's Silk Leggings; feet: Devout Sandals; finger1: Blessed Band of Light; finger2: Maiden's Circle; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Sageblade; off_hand: Tome of Fiery Arcana; ranged: Torch of Light

No-known-source sample (15 of 1299, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (human, 050105200100000000-55500201230133031-003000000000000000)

Set DPS (verified): 810.9. Weights run: 3.8s. Verify run: 23.6s. 1299 eligible items had no known source.

Race sweep: gnome 813.6 ± 1.0, human 810.9 ± 1.0, high-order-skyborne 794.9 ± 1.0 (dps); gnome beats the guide race human by 2.7 (0.34%) ± 1.4, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.877 ± 0.051, crit=0.643 ± 0.049 per rating point (14 rating = 1%, 9.003 per %), hit=1.286 ± 0.087 per rating point (10 rating = 1%, 12.862 per %), spell_haste=not significant (-3.924 ± 1.060), spell_penetration=not significant (0.000 ± 0.000), fire_power=0.857 ± 0.002, spirit=not significant (0.916 ± 0.329), mp5=not significant (1.387 ± 0.501)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Flame (3075) | World drop [world_drop] | 54.8 spell_power points (26.32 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.68 DPS) [pvp]; Dreadmist Mask (16698, -12.18 DPS, sim-verified) [dungeon] |
| neck | Diana's Pearl Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 28.9 spell_power points (13.87 DPS) | yes | Star of Mystaria (12103, -0.38 DPS) [dungeon]; Lady Maye's Pendant (14558, -1.47 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -1.89 DPS) [world_drop] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 37.4 spell_power points (17.97 DPS) | yes | Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Shroud of the Nathrezim (18720, -0.02 DPS) [dungeon]; Lieutenant Commander's Silk Mantle (227102, -0.04 DPS) [pvp] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -5.94 DPS) [world]; Shroud of Arcane Mastery (22330, -6.42 DPS) [dungeon]; Spritecaster Cape (11623, -13.70 DPS, sim-verified) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Vestments (231603, +0.00 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -4.82 DPS) [pvp]; Alanna's Embrace (13314, -18.40 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 35.4 spell_power points (17.02 DPS) | yes | Marshal's Silk Bracers (16438, -4.62 DPS) [pvp]; Netherflame Cuffs (254065, -6.52 DPS) [crafted]; Wyrmthalak's Shackles (13958, -6.63 DPS) [quest] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | sim-verified (+4.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Hands of Power (13253, -4.42 DPS, sim-verified) [dungeon] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 40.0 spell_power points (19.24 DPS) | yes | Marshal's Silk Sash (16439, -2.95 DPS) [pvp]; Thuzadin Sash (18740, -4.05 DPS) [dungeon]; Ban'thok Sash (11662, -6.32 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 58.8 spell_power points (28.24 DPS) | yes | Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Magister's Leggings (16687, -3.75 DPS) [dungeon]; Knight-Captain's Silk Legguards (227109, -4.46 DPS) [pvp] |
| feet | Wolfrunner Shoes (13101) | World drop [world_drop] | 32.7 spell_power points (15.72 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Mooncloth Boots (15802, -0.26 DPS) [crafted]; Dragonrider Boots (18102, -0.33 DPS) [dungeon] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Maiden's Circle (13001, -1.36 DPS) [world_drop]; Songstone of Ironforge (12543, -2.24 DPS) [quest]; Naglering (11669, -18.35 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -1.15 DPS) [quest]; Don Julio's Band (19325, -4.00 DPS) [rep]; Maiden's Circle (13001, -6.01 DPS, sim-verified) [world_drop] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -0.20 DPS) [crafted]; Ankh of Life (1713, -0.48 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.36 DPS) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+3.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -3.53 DPS, sim-verified) [crafted]; Serenity Field (272439, -3.84 DPS) [vendor]; Ankh of Life (1713, -5.29 DPS) [world_drop] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Staff of Jordan (873, -6.33 DPS) [world_drop]; Elemental Mage Staff (944, -13.26 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Torch of Light (279246, +0.00 DPS) [crafted]; Bonecreeper Stylus (13938, -0.29 DPS) [dungeon]; Pyric Caduceus (11748, -0.64 DPS) [dungeon] |

**New at 60:** neck: Diana's Pearl Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Magician's Cord; legs: Sentinel's Silk Leggings; feet: Wolfrunner Shoes; finger1: Channeler's Ring; finger2: Blessed Band of Light; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1299, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (undead, 000000000000000000-51400001000000000-000000000000000000)

Set DPS (verified): 39.1. Weights run: 3.9s. Verify run: 2.9s. 175 eligible items had no known source.

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
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 7.1 spell_power points (1.00 DPS) | yes | Loop of Sacrifice (281673, -0.29 DPS) [quest]; Volcanic Rock Ring (12053, -0.57 DPS) [world_drop]; The 1 Ring (8350, -0.84 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 6.1 spell_power points (0.86 DPS) | yes | Loop of Sacrifice (281673, -0.14 DPS) [quest]; Volcanic Rock Ring (12053, -0.43 DPS) [world_drop]; The 1 Ring (8350, -0.70 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Staff of the Blessed Seer (2271, -0.18 DPS) [dungeon]; Crescent Staff (6505, -0.35 DPS, sim-verified) [quest]; Gnarled Necromancer's Staff (251534, -0.47 DPS) [quest] |
| off_hand | Antipodean Rod (2879) | World drop [world_drop] | 6.0 spell_power points (0.84 DPS) | yes | Seer's Fine Stein (7608, -0.39 DPS) [world_drop]; Spellbinder Orb (15926, -0.41 DPS) [dungeon]; Tork Wrench (11855, -0.56 DPS) [quest] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller (12984, -0.66 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -1.14 DPS) [quest]; Greater Magic Wand (11288, -1.18 DPS) [crafted] |

**New at 20:** head: Flame Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Flame Gown; wrist: Mindthrust Bracers; hands: Flame Gloves; waist: Flame Sash; legs: Filigreed Flame Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Evocator's Blade; off_hand: Antipodean Rod; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 175, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 23173 Abomination Skin Leggings; 209618 Insignia of the Alliance

### Band 30 (undead, 000000000000000000-54500001230100000-000000000000000000)

Set DPS (verified): 79.3. Weights run: 2.6s. Verify run: 3.7s. 301 eligible items had no known source.

Race sweep: undead 79.3 ± 0.1, orc 77.4 ± 0.1, troll 76.5 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.884 ± 0.016, crit=0.220 ± 0.012 per rating point (14 rating = 1%, 3.085 per %), hit=0.395 ± 0.022 per rating point (10 rating = 1%, 3.955 per %), spell_haste=0.501 ± 0.104, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (-0.090 ± 0.097), mp5=not significant (0.514 ± 0.168)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Flame Circlet (253979) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Holy Shroud (2721, -0.74 DPS) [world_drop]; Flame Circlet (253953, -0.74 DPS) [crafted]; Enchanter's Cowl (4322, -0.80 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 5.3 spell_power points (1.31 DPS) | yes | Darkspear Warding Pendant (272075, -0.22 DPS) [vendor]; Crystal Starfire Medallion (5003, -0.44 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.44 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 17.0 spell_power points (4.19 DPS) | yes | Death Speaker Mantle (6685, -0.94 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.99 DPS) [world_drop]; Invoker's Mantle (215365, -1.37 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, -0.51 DPS) [crafted]; Windsong Drape (15468, -0.51 DPS) [quest]; Resilient Cape (14400, -0.93 DPS, sim-verified) [world_drop] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Flame Gown (253965, -0.10 DPS) [crafted]; Mechbuilder's Overalls (9508, -0.25 DPS) [dungeon]; Green Silk Armor (7065, -0.69 DPS, sim-verified) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (3.21 DPS) | yes | Spidertank Oilrag (9448, -0.77 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -1.47 DPS) [world_drop]; Mindthrust Bracers (1974, -1.74 DPS) [dungeon] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 10.4 spell_power points (2.58 DPS) | yes | Flame Gloves (253917, -0.19 DPS) [crafted]; Phoenix Gloves (4331, -0.35 DPS) [crafted]; Truefaith Gloves (7049, -0.68 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 13.7 spell_power points (3.38 DPS) | yes | Belt of Arugal (6392, -0.49 DPS) [dungeon]; Invoker's Cord (215366, -0.55 DPS) [crafted]; Crimson Silk Belt (7055, -0.97 DPS, sim-verified) [crafted] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | 20.2 spell_power points (4.99 DPS) | yes | Filigreed Flame Leggings (253941, -1.33 DPS, sim-verified) [crafted]; Phoenix Pants (4317, -1.64 DPS) [crafted]; Pristine Leggings (253987, -1.73 DPS) [crafted] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 19.2 spell_power points (4.74 DPS) | yes | Gilded Slippers (254001, -1.55 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.76 DPS) [dungeon]; Spidersilk Boots (4320, -2.14 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 10.4 spell_power points (2.58 DPS) | yes | Ring of Power Regulation (273030, -0.71 DPS) [dungeon]; Lavishly Jeweled Ring (1156, -1.26 DPS) [dungeon]; Azora's Will (4999, -1.48 DPS) [world_drop] |
| finger2 | Advisor's Ring (19521) | Warsong Outriders [rep] | 8.0 spell_power points (1.98 DPS) | yes | Ring of Power Regulation (273030, -0.12 DPS) [dungeon]; Lavishly Jeweled Ring (1156, -0.67 DPS) [dungeon]; Azora's Will (4999, -0.89 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -0.99 DPS) [dungeon]; Tendonscraper Dagger (273053, -0.99 DPS) [world]; Glimmering Staff (249392, -4.08 DPS, sim-verified) [crafted] |
| off_hand | Rod of Molten Fire (2565) | World drop [world_drop] | 13.0 spell_power points (3.21 DPS) | yes | Orb of Mystic Insight (249394, -1.09 DPS, sim-verified) [crafted]; Witch's Finger (16887, -1.47 DPS) [quest]; Antipodean Rod (2879, -1.73 DPS) [world_drop] |
| ranged | Scorching Wand (5213) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wand of Mana Concentration (273043, +0.00 DPS) [world]; Greater Mystic Wand (217287, -0.22 DPS) [crafted]; Starfaller (13063, -1.23 DPS, sim-verified) [world_drop] |

**New at 30:** head: Filigreed Flame Circlet; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Darkspear Raider's Cloak; chest: Beguiler Robes; wrist: Phoenix Bindings; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Flame Leggings; feet: Fiery Slippers; finger1: Snake Hoop; finger2: Advisor's Ring; main_hand: Swinetusk Shank; off_hand: Rod of Molten Fire; ranged: Scorching Wand

No-known-source sample (15 of 301, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (undead, 000000000000000000-54500001230133031-000000000000000000)

Set DPS (verified): 162.1. Weights run: 2.9s. Verify run: 4.2s. 392 eligible items had no known source.

Race sweep: undead 162.1 ± 0.3, orc 160.5 ± 0.3, troll 158.3 ± 0.3 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.539 ± 0.017, crit=0.322 ± 0.017 per rating point (14 rating = 1%, 4.511 per %), hit=0.553 ± 0.030 per rating point (10 rating = 1%, 5.533 per %), spell_haste=0.577 ± 0.103, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=0.644 ± 0.120, mp5=0.841 ± 0.165

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 23.0 spell_power points (8.76 DPS) | yes | Whitemane's Chapeau (7720, -1.08 DPS) [dungeon]; Electromagnetic Gigaflux Reactivator (9492, -1.85 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -4.14 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Necklace of Calisea (1714, -1.15 DPS) [dungeon]; Triune Amulet (7722, -1.38 DPS, sim-verified) [dungeon]; Amberglow Talisman (10824, -1.86 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Green Silken Shoulders (7057, -0.03 DPS) [crafted]; Bloodmage Mantle (7684, -0.06 DPS) [dungeon]; Mantle of Doan (7712, -1.36 DPS, sim-verified) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battle Healer's Cloak (19528, -0.63 DPS) [rep]; Long Silken Cloak (4326, -0.83 DPS) [crafted]; Guardian Cloak (5965, -0.83 DPS) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 31.9 spell_power points (12.15 DPS) | yes | Robe of the Magi (1716, -1.28 DPS, sim-verified) [world_drop]; Cindercloth Robe (10042, -1.86 DPS) [crafted]; Robe of Power (7054, -2.38 DPS) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (4.96 DPS) | yes | Spidertank Oilrag (9448, -0.56 DPS) [dungeon]; Radiant Silver Bracers (4545, -1.79 DPS) [quest]; Earthen Silk Cuffs (254019, -2.20 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 24.7 spell_power points (9.40 DPS) | yes | Fiery Handwraps (254025, -1.44 DPS, sim-verified) [crafted]; Crimson Silk Gloves (7064, -2.83 DPS) [crafted]; Red Mageweave Gloves (10018, -3.16 DPS) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | 22.2 spell_power points (8.46 DPS) | yes | Gilded Cord (254037, -2.29 DPS) [crafted]; Defiler's Cloth Girdle (20166, -2.30 DPS) [rep]; Scorching Sash (4117, -4.04 DPS, sim-verified) [quest] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Stoneweaver Leggings (9407, -0.54 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.02 DPS) [dungeon]; Flame Leggings (253991, -1.05 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 28.5 spell_power points (10.87 DPS) | yes | Fiery Slippers (254005, -2.69 DPS, sim-verified) [crafted]; Furen's Boots (13100, -5.50 DPS) [world_drop]; Thoughtcast Boots (10578, -5.77 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.7 spell_power points (4.84 DPS) | yes | Snake Hoop (6750, -1.53 DPS) [quest]; Ring of Power Regulation (273030, -1.59 DPS) [dungeon]; Electrocutioner Lagnut (9447, -1.98 DPS) [dungeon] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 11.5 spell_power points (4.39 DPS) | yes | Ring of Power Regulation (273030, -1.14 DPS) [dungeon]; Snake Hoop (6750, -1.33 DPS, sim-verified) [quest]; Electrocutioner Lagnut (9447, -1.53 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -13.37 DPS, sim-verified) [dungeon]; Celestial Stave (9517, -13.43 DPS) [quest]; Hypnotic Blade (7714, -13.45 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Ragefire Wand (7513) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Twisted Nether Wand (249144, -2.06 DPS) [crafted]; Jaina's Firestarter (13064, -2.15 DPS, sim-verified) [world_drop]; Dancing Flame (6806, -2.28 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Icy Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Fiery Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Staff of Jordan; ranged: Ragefire Wand

No-known-source sample (15 of 392, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (undead, 230002000000000000-54500001230133031-003000000000000000)

Set DPS (verified): 232.5. Weights run: 2.9s. Verify run: 3.5s. 504 eligible items had no known source.

Race sweep: undead 232.5 ± 0.4, orc 229.1 ± 0.4, troll 226.6 ± 0.4 (dps); the guide race undead is best.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.846 ± 0.031, crit=0.394 ± 0.025 per rating point (14 rating = 1%, 5.511 per %), hit=0.684 ± 0.041 per rating point (10 rating = 1%, 6.836 per %), spell_haste=not significant (0.198 ± 0.244), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.740 ± 0.305), mp5=not significant (0.864 ± 0.323)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Flame (3075) | World drop [world_drop] | 58.9 spell_power points (12.24 DPS) | yes | Red Mageweave Headband (10033, -4.77 DPS) [crafted]; Gemburst Circlet (10751, -5.65 DPS) [quest]; Dreamweave Circlet (10041, -7.23 DPS, sim-verified) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 21.1 spell_power points (4.38 DPS) | yes | Mindburst Medallion (11196, -1.31 DPS) [quest]; Glowing Eye of Mordresh (10769, -1.52 DPS) [dungeon]; Horizon Choker (13085, -2.19 DPS, sim-verified) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kentic Amice (11624, +0.00 DPS) [dungeon]; Netherflame Shoulders (254053, -0.00 DPS) [crafted]; Flameseer Mantle (11310, -1.42 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -0.81 DPS) [quest]; Runecloth Cloak (13860, -1.61 DPS) [crafted]; Cindercloth Cloak (14044, -2.05 DPS, sim-verified) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Embrace of the Wind Serpent (12462, +0.00 DPS) [world]; Acumen Robes (17775, -0.01 DPS) [quest]; Robes of Insight (940, -0.77 DPS) [world_drop] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 22.6 spell_power points (4.71 DPS) | yes | Nethergeld Cuffs (254061, -1.25 DPS) [crafted]; Bloodband Bracers (11469, -1.62 DPS) [quest]; Aristocratic Cuffs (12546, -2.02 DPS, sim-verified) [dungeon] |
| hands | Fiery Gloves (254099) | Tailoring [crafted] | 33.4 spell_power points (6.94 DPS) | yes | Runecloth Gloves (13863, -1.48 DPS) [crafted]; Cindercloth Gloves (14043, -1.71 DPS) [crafted]; Dreamweave Gloves (10019, -2.05 DPS, sim-verified) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 28.1 spell_power points (5.85 DPS) | yes | Fiery Waistcord (254085, -0.32 DPS) [crafted]; Fiery Cord (254041, -0.61 DPS) [crafted]; Satyrmane Sash (17755, -1.18 DPS) [dungeon] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 31.6 spell_power points (6.56 DPS) | yes | Spellshock Leggings (9484, -0.02 DPS) [dungeon]; Senior Designer's Pantaloons (11841, -0.06 DPS) [dungeon]; Cindercloth Leggings (12256, -0.30 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 29.2 spell_power points (6.07 DPS) | yes | Fiery Sandals (254111, -0.33 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -1.40 DPS) [vendor]; Fiery Slippers (254005, -1.52 DPS) [crafted] |
| finger1 | Advisor's Ring (19519) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cyclopean Band (11824, +0.00 DPS) [dungeon]; Philanthropist's Ring (281635, -0.26 DPS) [quest]; Eye of Adaegus (5266, -0.31 DPS) [world_drop] |
| finger2 | Brainlash (6440) | Zul'Farrak [dungeon] | 16.4 spell_power points (3.41 DPS) | yes | Cyclopean Band (11824, +0.00 DPS) [dungeon]; Philanthropist's Ring (281635, -0.45 DPS) [quest]; Eye of Adaegus (5266, -0.51 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.31 DPS) [quest]; Frozen Heart of the Mountain (249469, -0.57 DPS) [crafted]; Uther's Strength (11302, -0.60 DPS) [world_drop] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, -2.78 DPS) [world_drop]; Spire of Hakkar (10844, -4.31 DPS) [world]; Kindling Stave (11750, -6.72 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lesser Eternal Wand (249232, -3.54 DPS) [crafted]; Wand of Allistarj (13065, -4.40 DPS) [world_drop]; Noxious Shooter (17745, -6.74 DPS, sim-verified) [dungeon] |

**New at 50:** head: Eye of Flame; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Spritecaster Cape; wrist: Netherflame Cuffs; hands: Fiery Gloves; waist: Ban'thok Sash; legs: Kilt of the Atal'ai Prophet; finger1: Advisor's Ring; finger2: Brainlash; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; ranged: Pyric Caduceus

No-known-source sample (15 of 504, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (undead, 230105200100000000-55500201230133031-003000000000000000)

Set DPS (verified): 395.5. Weights run: 3.2s. Verify run: 23.7s. 1293 eligible items had no known source.

Race sweep: undead 395.5 ± 0.7, orc 390.1 ± 0.7, troll 387.5 ± 0.7 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.599 ± 0.052, crit=0.463 ± 0.037 per rating point (14 rating = 1%, 6.476 per %), hit=1.239 ± 0.073 per rating point (10 rating = 1%, 12.390 per %), spell_haste=not significant (-2.546 ± 0.966), spell_penetration=not significant (0.000 ± 0.000), fire_power=0.847 ± 0.003, spirit=not significant (0.338 ± 0.669), mp5=not significant (-0.531 ± 0.679)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Flame (3075, +0.00 DPS) [world_drop]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Magister's Crown (16686, -22.92 DPS, sim-verified) [dungeon] |
| neck | Diana's Pearl Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 26.2 spell_power points (5.96 DPS) | yes | Star of Mystaria (12103, -1.30 DPS) [dungeon]; Arcane Crystal Pendant (20037, -1.50 DPS) [quest]; Archlight Talisman (15856, -1.87 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Magister's Mantle (16689, -22.16 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -3.09 DPS) [dungeon]; Shroud of Arcane Mastery (22330, -3.24 DPS) [dungeon]; Crystalline Threaded Cape (20697, -4.96 DPS, sim-verified) [world] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Raiment (231596, -0.45 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -3.26 DPS) [pvp]; Alanna's Embrace (13314, -8.76 DPS, sim-verified) [dungeon] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dryad's Wrist Bindings (19595, +0.00 DPS) [rep]; Magister's Bindings (16683, -21.70 DPS, sim-verified) [dungeon] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Magister's Gloves (16684, -18.75 DPS, sim-verified) [dungeon] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 33.4 spell_power points (7.60 DPS) | yes | Defiler's Cloth Girdle (20163, -2.12 DPS) [rep]; Frostwolf Cloth Belt (19090, -2.14 DPS) [rep]; Ban'thok Sash (11662, -5.96 DPS, sim-verified) [dungeon] |
| legs | Magister's Leggings (16687) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Outrider's Silk Leggings (22747, +0.00 DPS) [rep]; Legionnaire's Silk Legguards (227107, +0.00 DPS) [pvp]; General's Silk Trousers (231595, +0.00 DPS) [pvp] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Omnicast Boots (11822, +0.00 DPS) [dungeon]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -20.88 DPS, sim-verified) [dungeon] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.35 DPS) [quest]; Channeler's Ring (272406, -0.37 DPS) [vendor]; Naglering (11669, -12.22 DPS, sim-verified) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.15 DPS) [quest]; Channeler's Ring (272406, -0.17 DPS) [vendor]; Flaming Band (12926, -0.89 DPS) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -0.19 DPS) [crafted]; Rune of the Guard Captain (19120, -0.76 DPS) [quest]; Second Wind (11819, -12.00 DPS, sim-verified) [dungeon] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -2.47 DPS) [crafted]; Rune of the Guard Captain (19120, -3.04 DPS) [quest]; Serenity Field (272439, -3.25 DPS, sim-verified) [vendor] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -0.51 DPS) [pvp]; Elemental Mage Staff (944, -2.22 DPS) [world_drop]; Ironbark Staff (20220, -9.27 DPS, sim-verified) [rep] |
| off_hand | Tome of Fiery Arcana (19311) | Frostwolf Clan [rep] | 31.3 spell_power points (7.14 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; Therazane's Touch (19315, -0.53 DPS) [rep]; High Warlord's Tome of Mending (234564, -1.81 DPS) [pvp] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, -3.31 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -3.85 DPS, sim-verified) [dungeon]; Brilliant Wand (249385, -4.25 DPS) [crafted] |

**New at 60:** head: Devout Crown; neck: Diana's Pearl Necklace; shoulder: Devout Mantle; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Devout Bracers; hands: Devout Gloves; waist: Magician's Cord; legs: Magister's Leggings; feet: Devout Sandals; finger1: Blessed Band of Light; finger2: Maiden's Circle; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Sageblade; off_hand: Tome of Fiery Arcana; ranged: Torch of Light

No-known-source sample (15 of 1293, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60, raid preset (undead, 050105200100000000-55500201230133031-003000000000000000)

Set DPS (verified): 798.4. Weights run: 3.8s. Verify run: 23.6s. 1293 eligible items had no known source.

Race sweep: undead 798.4 ± 0.9, orc 796.2 ± 0.9, troll 788.8 ± 0.9 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.877 ± 0.051, crit=0.643 ± 0.049 per rating point (14 rating = 1%, 9.003 per %), hit=1.286 ± 0.087 per rating point (10 rating = 1%, 12.862 per %), spell_haste=not significant (-3.924 ± 1.060), spell_penetration=not significant (0.000 ± 0.000), fire_power=0.857 ± 0.002, spirit=not significant (0.916 ± 0.329), mp5=not significant (1.387 ± 0.501)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Flame (3075) | World drop [world_drop] | 54.8 spell_power points (26.32 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -1.68 DPS) [pvp]; Dreadmist Mask (16698, -9.38 DPS, sim-verified) [dungeon] |
| neck | Diana's Pearl Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 28.9 spell_power points (13.87 DPS) | yes | Lady Maye's Pendant (14558, -1.47 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -1.89 DPS) [world_drop]; Star of Mystaria (12103, -7.40 DPS, sim-verified) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 37.4 spell_power points (17.97 DPS) | yes | Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Champion's Silk Mantle (227104, -0.04 DPS) [pvp]; Shroud of the Nathrezim (18720, -6.03 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -5.34 DPS) [dungeon]; Crystalline Threaded Cape (20697, -5.94 DPS) [world]; Cloak of Flames (3475, -10.58 DPS, sim-verified) [world_drop] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Raiment (231596, +0.00 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -4.82 DPS) [pvp]; Alanna's Embrace (13314, -17.84 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 35.4 spell_power points (17.02 DPS) | yes | General's Silk Cuffs (16538, -4.62 DPS) [pvp]; Netherflame Cuffs (254065, -6.52 DPS) [crafted]; Wyrmthalak's Shackles (13958, -6.63 DPS) [quest] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | sim-verified (+7.1 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Hands of Power (13253, -7.05 DPS, sim-verified) [dungeon] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 40.0 spell_power points (19.24 DPS) | yes | Ban'thok Sash (11662, -2.66 DPS) [dungeon]; General's Silk Sash (16537, -2.95 DPS) [pvp]; Thuzadin Sash (18740, -4.05 DPS) [dungeon] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 53.8 spell_power points (25.86 DPS) | yes | General's Silk Trousers (231595, +0.00 DPS) [pvp]; Magister's Leggings (16687, -1.37 DPS) [dungeon]; Legionnaire's Silk Legguards (227107, -2.08 DPS) [pvp] |
| feet | Wolfrunner Shoes (13101) | World drop [world_drop] | 32.7 spell_power points (15.72 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Mooncloth Boots (15802, -0.26 DPS) [crafted]; Dragonrider Boots (18102, -0.33 DPS) [dungeon] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Maiden's Circle (13001, -1.36 DPS) [world_drop]; Eye of Orgrimmar (12545, -2.24 DPS) [quest]; Naglering (11669, -18.03 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -1.15 DPS) [quest]; Don Julio's Band (19325, -4.00 DPS) [rep]; Maiden's Circle (13001, -5.79 DPS, sim-verified) [world_drop] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+4.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Serenity Field (272439, -3.84 DPS) [vendor]; Frozen Heart of the Mountain (249469, -3.97 DPS, sim-verified) [crafted]; Ankh of Life (1713, -5.29 DPS) [world_drop] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (-8.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -0.20 DPS) [crafted]; Ankh of Life (1713, -0.48 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.36 DPS) [quest] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Staff of Jordan (873, -6.33 DPS) [world_drop]; Elemental Mage Staff (944, -14.06 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Torch of Light (279246, +0.00 DPS) [crafted]; Bonecreeper Stylus (13938, -0.29 DPS) [dungeon]; Pyric Caduceus (11748, -0.64 DPS) [dungeon] |

**New at 60:** neck: Diana's Pearl Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Magician's Cord; legs: Outrider's Silk Leggings; feet: Wolfrunner Shoes; finger1: Channeler's Ring; finger2: Blessed Band of Light; trinket1: Weakness Analyzer; trinket2: Burst of Knowledge; main_hand: Ironbark Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1293, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

