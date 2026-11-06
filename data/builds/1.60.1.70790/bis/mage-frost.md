# Leveling BiS: Frost

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-253100000000000000)

Set DPS (verified): 49.0. Weights run: 1.8s. Verify run: 2.2s. 181 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.981 ± 0.007, crit=0.092 ± 0.003 per rating point (14 rating = 1%, 1.294 per %), hit=0.262 ± 0.008 per rating point (10 rating = 1%, 2.623 per %), spell_haste=0.663 ± 0.102, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002, spirit=not significant (0.000 ± 0.000), mp5=2.089 ± 0.105

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Silky Circlet (253951) | Tailoring [crafted] | 11.0 spell_power points (2.37 DPS) | yes | Pristine Circlet (253949, -1.08 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.31 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.8 spell_power points (2.98 DPS) | yes | Reinforced Woolen Shoulders (4315, -1.17 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -1.63 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -2.12 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Sanguine Cape (14376, -0.02 DPS) [world_drop]; Pearl-clasped Cloak (5542, -0.21 DPS, sim-verified) [crafted]; Black Whelp Cloak (7283, -0.22 DPS) [crafted] |
| chest | Filigreed Silky Gown (253903) | Tailoring [crafted] | 13.9 spell_power points (2.99 DPS) | yes | Filigreed Pristine Gown (253901, -0.86 DPS, sim-verified) [crafted]; Gray Woolen Robe (2585, -1.08 DPS) [crafted]; Mystic's Wrap (14369, -1.51 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 11.2 spell_power points (2.40 DPS) | yes | Bright Bracers (3647, -0.41 DPS, sim-verified) [world_drop]; Mystic's Bracelets (14366, -1.98 DPS) [world_drop]; Repurposed Hair Band (281256, -1.98 DPS) [quest] |
| hands | Silky Gloves (253915) | Tailoring [crafted] | 9.9 spell_power points (2.14 DPS) | yes | Serpent Gloves (5970, -0.38 DPS, sim-verified) [dungeon]; Pristine Gloves (253913, -0.65 DPS) [crafted]; Gnoll Casting Gloves (892, -0.85 DPS) [world] |
| waist | Silky Sash (253927) | Tailoring [crafted] | 10.9 spell_power points (2.35 DPS) | yes | Pristine Sash (253925, -0.65 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.66 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.86 DPS) [crafted] |
| legs | Filigreed Silky Leggings (253939) | Tailoring [crafted] | 16.9 spell_power points (3.63 DPS) | yes | Filigreed Pristine Leggings (253937, -1.08 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -2.13 DPS) [dungeon]; Darkweave Breeches (12987, -2.16 DPS) [world_drop] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.9 spell_power points (2.35 DPS) | yes | Silky Boots (253891, -0.38 DPS, sim-verified) [crafted]; Pristine Boots (253889, -1.07 DPS) [crafted]; Red Woolen Boots (4313, -1.49 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 9.2 spell_power points (1.98 DPS) | yes | Lavishly Jeweled Ring (1156, -0.71 DPS) [dungeon]; Volcanic Rock Ring (12053, -1.34 DPS) [world_drop]; Furlfeather Ring (286745, -1.76 DPS) [world] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 7.0 spell_power points (1.50 DPS) | yes | Lavishly Jeweled Ring (1156, -0.30 DPS, sim-verified) [dungeon]; Volcanic Rock Ring (12053, -0.87 DPS) [world_drop]; Furlfeather Ring (286745, -1.28 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | 30.0 spell_power points (6.45 DPS) | yes | Emberstone Staff (5201, -1.28 DPS) [dungeon]; Staff of the Blessed Seer (2271, -1.45 DPS, sim-verified) [dungeon]; Buzzer Blade (2169, -1.71 DPS) [dungeon] |
| off_hand | Antipodean Rod (2879) | World drop [world_drop] | 6.0 spell_power points (1.29 DPS) | yes | Mathystra Channeling Orb (281291, -0.24 DPS) [quest]; Tear of Grief (5611, -0.43 DPS, sim-verified) [quest]; Dwarven Tome (279898, -0.65 DPS) [quest] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller (12984, -0.47 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -0.68 DPS) [quest]; Greater Magic Wand (11288, -0.71 DPS) [crafted] |

**New at 20:** head: Silky Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Silky Gown; wrist: Mindthrust Bracers; hands: Silky Gloves; waist: Silky Sash; legs: Filigreed Silky Leggings; feet: Spidersilk Boots; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Evocator's Blade; off_hand: Antipodean Rod; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 181, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 30 (gnome, 000000000000000000-00000000000000000-253510130100000000)

Set DPS (verified): 83.4. Weights run: 1.8s. Verify run: 1.5s. 309 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.300 ± 0.005, crit=0.111 ± 0.004 per rating point (14 rating = 1%, 1.551 per %), hit=0.207 ± 0.006 per rating point (10 rating = 1%, 2.068 per %), spell_haste=1.058 ± 0.060, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002, spirit=0.197 ± 0.037, mp5=0.288 ± 0.041

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Silky Circlet (253977) | Tailoring [crafted] | 14.0 spell_power points (2.88 DPS) | yes | Holy Shroud (2721, -0.37 DPS) [world_drop]; Silky Circlet (253951, -0.62 DPS) [crafted]; Azure Silk Hood (7048, -0.82 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 2.0 spell_power points (0.41 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.04 DPS) [world]; Crystal Starfire Medallion (5003, -0.04 DPS) [world_drop]; Seal of the Expedition (271920, -0.10 DPS) [rep] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.7 spell_power points (2.41 DPS) | yes | Death Speaker Mantle (6685, -0.51 DPS, sim-verified) [dungeon]; Invoker's Mantle (215365, -0.66 DPS) [crafted]; Magician's Mantle (12998, -0.82 DPS) [world_drop] |
| back | Azure Silk Cloak (7053) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, -1.15 DPS) [crafted]; Heavy Woolen Cloak (4311, -1.19 DPS) [crafted]; Vine Pruner's Cloak (279835, -1.61 DPS, sim-verified) [quest] |
| chest | Silky Gown (253963) | Tailoring [crafted] | 15.9 spell_power points (3.27 DPS) | yes | Beguiler Robes (7728, -0.35 DPS) [dungeon]; Tree Bark Jacket (1486, -0.59 DPS) [dungeon]; Green Silk Armor (7065, -0.62 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.9 spell_power points (2.03 DPS) | yes | Nightsky Wristbands (6407, -1.54 DPS) [world_drop]; Mindthrust Bracers (1974, -1.54 DPS) [dungeon]; Glowing Magical Bracelets (13106, -2.02 DPS, sim-verified) [world_drop] |
| hands | Azure Silk Gloves (4319) | Tailoring [crafted] | 10.6 spell_power points (2.18 DPS) | yes | Silky Gloves (253915, -0.63 DPS, sim-verified) [crafted]; Serpent Gloves (5970, -0.74 DPS) [dungeon]; Gloves of Meditation (4318, -0.87 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.9 spell_power points (2.45 DPS) | yes | Belt of Arugal (6392, -0.21 DPS) [dungeon]; Invoker's Cord (215366, -0.70 DPS) [crafted]; Silky Sash (253927, -0.76 DPS) [crafted] |
| legs | Silky Leggings (253989) | Tailoring [crafted] | 17.1 spell_power points (3.52 DPS) | yes | Filigreed Silky Leggings (253939, -0.72 DPS) [crafted]; Azure Silk Pants (7046, -1.09 DPS) [crafted]; Gaze Dreamer Pants (6903, -1.61 DPS, sim-verified) [dungeon] |
| feet | Frothing Slippers (254003) | Tailoring [crafted] | 15.9 spell_power points (3.27 DPS) | yes | Spidersilk Boots (4320, -1.58 DPS) [crafted]; Acidic Walkers (9454, -1.59 DPS) [dungeon]; Gilded Slippers (254001, -1.60 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.6 spell_power points (1.56 DPS) | yes | Ring of Power Regulation (273030, -0.15 DPS) [dungeon]; Minor Channeling Ring (1449, -0.41 DPS) [quest]; Electrocutioner Lagnut (9447, -0.66 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 7.5 spell_power points (1.54 DPS) | yes | Ring of Power Regulation (273030, -0.13 DPS) [dungeon]; Minor Channeling Ring (1449, -0.39 DPS) [quest]; Electrocutioner Lagnut (9447, -0.64 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Soulstaff (249393, -0.95 DPS) [crafted]; Tendonscraper Dagger (273053, -0.99 DPS) [world]; Glimmering Staff (249392, -3.11 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.8 spell_power points (1.81 DPS) | yes | Orb of Souls (249395, -0.74 DPS) [crafted]; Eye of Paleth (2943, -0.99 DPS) [quest]; Antipodean Rod (2879, -1.26 DPS, sim-verified) [world_drop] |
| ranged | Thunderwood (13062) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Greater Mystic Wand (217287, -0.06 DPS) [crafted]; Wand of Eventide (5214, -0.31 DPS) [world_drop]; Starfaller (13063, -0.55 DPS) [world_drop] |

**New at 30:** head: Filigreed Silky Circlet; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Azure Silk Cloak; chest: Silky Gown; wrist: Spidertank Oilrag; hands: Azure Silk Gloves; waist: Highlander's Cloth Girdle; legs: Silky Leggings; feet: Frothing Slippers; finger1: Lorekeeper's Ring; finger2: Snake Hoop; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Thunderwood

No-known-source sample (15 of 309, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 000000000000000000-00000000000000000-253510130101301221)

Set DPS (verified): 178.4. Weights run: 2.0s. Verify run: 2.0s. 400 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.684 ± 0.016, crit=0.164 ± 0.006 per rating point (14 rating = 1%, 2.297 per %), hit=0.343 ± 0.017 per rating point (10 rating = 1%, 3.430 per %), spell_haste=1.030 ± 0.198, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003, spirit=not significant (0.538 ± 0.144), mp5=0.734 ± 0.183

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Papal Fez (9431, +0.00 DPS) [dungeon]; Whitemane's Chapeau (7720, -0.18 DPS) [dungeon]; Electromagnetic Gigaflux Reactivator (9492, -0.69 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Triune Amulet (7722, -0.18 DPS) [dungeon]; Necklace of Calisea (1714, -1.01 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.45 DPS) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Green Silken Shoulders (7057, +0.00 DPS) [crafted]; Mantle of Doan (7712, +0.00 DPS) [dungeon]; Inquisitor's Shawl (19507, +0.00 DPS) [dungeon] |
| back | Azure Silk Cloak (7053) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frost Leather Cloak (7377, -0.16 DPS) [crafted]; Icy Cloak (4327, -0.36 DPS) [crafted]; Expeditionary's Cape (271922, -0.50 DPS) [rep] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Robe of the Magi (1716, -0.74 DPS) [world_drop]; Robe of Power (7054, -1.33 DPS) [crafted]; Beguiler Robes (7728, -2.61 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 11.2 spell_power points (2.88 DPS) | yes | Arcane Runed Bracers (4744, -0.57 DPS) [quest]; Condor Bracers (15864, -1.08 DPS) [quest]; Earthen Silk Cuffs (254019, -1.16 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 24.5 spell_power points (6.29 DPS) | yes | Frothing Handwraps (254023, -1.47 DPS) [crafted]; Red Mageweave Gloves (10018, -1.71 DPS) [crafted]; Black Mageweave Gloves (10003, -2.44 DPS) [crafted] |
| waist | Frothing Cord (254039) | Tailoring [crafted] | 22.7 spell_power points (5.83 DPS) | yes | Deathmage Sash (10771, -1.40 DPS) [dungeon]; Highlander's Cloth Girdle (20098, -1.53 DPS) [rep]; Gilded Cord (254037, -1.54 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.2 spell_power points (5.70 DPS) | yes | Silky Leggings (253989, -0.19 DPS) [crafted]; Stoneweaver Leggings (9407, -0.84 DPS) [dungeon]; Filigreed Silky Leggings (253939, -1.27 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 27.8 spell_power points (7.13 DPS) | yes | Frothing Slippers (254003, -2.01 DPS) [crafted]; Gilded Slippers (254001, -3.55 DPS) [crafted]; Furen's Boots (13100, -3.69 DPS) [world_drop] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.4 spell_power points (3.45 DPS) | yes | Snake Hoop (6750, -1.03 DPS) [quest]; Ring of Power Regulation (273030, -1.34 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -1.39 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 11.2 spell_power points (2.88 DPS) | yes | Snake Hoop (6750, -0.46 DPS) [quest]; Ring of Power Regulation (273030, -0.77 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -0.82 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -7.72 DPS) [dungeon]; Celestial Stave (9517, -8.41 DPS) [quest]; Hypnotic Blade (7714, -8.51 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Icefury Wand (7514) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Freezing Shard (10572, -0.29 DPS) [dungeon]; Jaina's Firestarter (13064, -1.34 DPS) [world_drop]; Twisted Nether Wand (249144, -1.35 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Frothing Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Staff of Jordan; ranged: Icefury Wand

No-known-source sample (15 of 400, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (gnome, 230005000000000000-00000000000000000-253510130101301221)

Set DPS (verified): 259.2. Weights run: 2.0s. Verify run: 5.1s. 512 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.638 ± 0.027, crit=0.279 ± 0.011 per rating point (14 rating = 1%, 3.908 per %), hit=0.696 ± 0.036 per rating point (10 rating = 1%, 6.960 per %), spell_haste=not significant (1.410 ± 0.386), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003, spirit=not significant (1.150 ± 0.366), mp5=not significant (1.193 ± 0.391)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreamweave Circlet (10041) | Tailoring [crafted] | 41.2 spell_power points (9.95 DPS) | yes | Red Mageweave Headband (10033, -2.28 DPS) [crafted]; Papal Fez (9431, -2.90 DPS) [dungeon]; Gemburst Circlet (10751, -4.57 DPS, sim-verified) [quest] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 19.8 spell_power points (4.79 DPS) | yes | Triune Amulet (7722, -0.91 DPS) [dungeon]; Mindburst Medallion (11196, -1.03 DPS) [quest]; Lei of Lilies (1315, -6.60 DPS, sim-verified) [world_drop] |
| shoulder | Netherfroth Shoulders (254051) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kentic Amice (11624, +0.00 DPS) [dungeon]; Rotgrip Mantle (17732, -1.09 DPS) [dungeon]; Nethergeld Shoulders (254049, -1.93 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Expeditionary's Cape (271921, -2.43 DPS) [rep]; Wingveil Cloak (10802, -2.44 DPS) [dungeon]; Featherskin Cape (10843, -5.04 DPS, sim-verified) [world] |
| chest | Ghostweave Vest (14141) | Tailoring [crafted] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Dreamweave Vest (10021, -0.41 DPS) [crafted]; Vestments of the Atal'ai Prophet (10806, -0.84 DPS) [dungeon]; Embrace of the Wind Serpent (12462, -0.87 DPS, sim-verified) [world] |
| wrist | Netherfroth Cuffs (254063) | Tailoring [crafted] | 23.2 spell_power points (5.61 DPS) | yes | Aristocratic Cuffs (12546, -1.63 DPS) [dungeon]; Bloodband Bracers (11469, -2.18 DPS) [quest]; Nethergeld Cuffs (254061, -2.29 DPS, sim-verified) [crafted] |
| hands | Frothing Gloves (254097) | Tailoring [crafted] | 34.6 spell_power points (8.36 DPS) | yes | Runecloth Gloves (13863, -1.57 DPS) [crafted]; Dreamweave Gloves (10019, -1.91 DPS, sim-verified) [crafted]; Ghostweave Gloves (14142, -2.01 DPS) [crafted] |
| waist | Frothing Cord (254039) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Ban'thok Sash (11662, -0.01 DPS) [dungeon]; Frothing Waistcord (254083, -0.31 DPS) [crafted]; Ghostweave Cord (254073, -0.77 DPS, sim-verified) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 37.2 spell_power points (8.98 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -0.48 DPS) [dungeon]; Dalewind Trousers (13008, -0.65 DPS) [world_drop]; Rainstrider Leggings (11123, -1.48 DPS) [quest] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 32.1 spell_power points (7.75 DPS) | yes | Frothing Sandals (254109, -1.52 DPS) [crafted]; Vinerot Sandals (17748, -2.13 DPS) [dungeon]; Mistwalker Boots (10629, -6.80 DPS, sim-verified) [dungeon] |
| finger1 | Freezing Band (942) | World drop [world_drop] | 21.0 spell_power points (5.08 DPS) | yes | Eye of Adaegus (5266, -0.81 DPS) [world_drop]; Lorekeeper's Ring (19523, -1.02 DPS) [rep]; Brainlash (6440, -1.37 DPS) [dungeon] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 18.1 spell_power points (4.37 DPS) | yes | Lorekeeper's Ring (19523, -0.31 DPS) [rep]; Brainlash (6440, -0.66 DPS) [dungeon]; Eye of Adaegus (5266, -3.02 DPS, sim-verified) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Uther's Strength (11302, -0.06 DPS) [world_drop] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Uther's Strength (11302, -0.81 DPS, sim-verified) [world_drop] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spire of Hakkar (10844, -3.46 DPS) [world]; Glowing Brightwood Staff (812, -3.83 DPS) [world_drop]; Kindling Stave (11750, -5.48 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Icefury Wand (7514) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Freezing Shard (10572, -0.26 DPS) [dungeon]; Cairnstone Sliver (9654, -0.36 DPS) [quest]; Lesser Eternal Wand (249232, -0.47 DPS) [crafted] |

**New at 50:** head: Dreamweave Circlet; neck: Arcane Crystal Pendant; shoulder: Netherfroth Shoulders; back: Spritecaster Cape; chest: Ghostweave Vest; wrist: Netherfroth Cuffs; hands: Frothing Gloves; legs: Senior Designer's Pantaloons; finger1: Freezing Band; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Thunderbrew's Boot Flask

No-known-source sample (15 of 512, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (gnome, 230005203000000000-50000000000000000-253510130101301221)

Set DPS (verified): 408.7. Weights run: 2.2s. Verify run: 15.2s. 1299 eligible items had no known source.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.222 ± 0.016, crit=0.521 ± 0.015 per rating point (14 rating = 1%, 7.292 per %), hit=1.140 ± 0.053 per rating point (10 rating = 1%, 11.397 per %), spell_haste=6.781 ± 0.247, spell_penetration=not significant (0.000 ± 0.000), frost_power=0.808 ± 0.001, spirit=not significant (0.027 ± 0.055), mp5=not significant (0.152 ± 0.100)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crimson Felt Hat (18727, +0.00 DPS) [dungeon]; Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arcane Crystal Pendant (20037, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon]; Star of Mystaria (12103, -0.25 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 27.2 spell_power points (11.00 DPS) | yes | Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Boreal Mantle (11782, -0.80 DPS) [dungeon]; Argent Shoulders (19059, -0.81 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -3.89 DPS, sim-verified) [world]; Spritecaster Cape (11623, -5.53 DPS) [dungeon]; Stormpike Sage's Cloak (19086, -6.14 DPS) [rep] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Vestments (231603, -2.32 DPS) [pvp]; Freezing Lich Robes (14340, -6.16 DPS) [dungeon]; Robe of Winter Night (14136, -7.78 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 24.0 spell_power points (9.70 DPS) | yes | Runecloth Cuffs (254123, -4.35 DPS) [crafted]; Netherfroth Cuffs (254063, -4.76 DPS) [crafted]; Spidertank Oilrag (9448, -5.87 DPS) [dungeon] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Marshal's Silk Handwraps (231609, +0.00 DPS) [vendor] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Highlander's Cloth Girdle (20047, -0.69 DPS) [rep]; Stormpike Cloth Girdle (19094, -1.66 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 48.8 spell_power points (19.75 DPS) | yes | Marshal's Silk Leggings (231605, -2.75 DPS) [pvp]; Skyshroud Leggings (13170, -6.20 DPS, sim-verified) [dungeon]; Knight-Captain's Silk Legguards (227109, -6.72 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.2 spell_power points (9.79 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Knight-Lieutenant's Silk Walkers (227112, -0.39 DPS) [pvp]; Omnicast Boots (11822, -0.62 DPS) [dungeon] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.79 DPS) [quest]; Don Julio's Band (19325, -1.18 DPS) [rep]; Naglering (11669, -10.80 DPS, sim-verified) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.02 DPS) [quest]; Don Julio's Band (19325, -0.42 DPS) [rep]; Don Mauricio's Band of Domination (22433, -0.58 DPS) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+8.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -0.71 DPS) [crafted]; Uther's Strength (11302, -2.43 DPS) [world_drop]; Shard of the Splithooves (10659, -4.49 DPS) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -3.35 DPS, sim-verified) [vendor]; Frozen Heart of the Mountain (249469, -4.75 DPS) [crafted]; Uther's Strength (11302, -6.47 DPS) [world_drop] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -2.44 DPS) [pvp]; Sageclaw (20070, -4.63 DPS) [rep]; Ironbark Staff (20069, -13.33 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (11.73 DPS) | yes | Grand Marshal's Tome of Power (234589, -0.45 DPS) [pvp]; Tome of the Ice Lord (19310, -0.88 DPS) [rep]; Spirit of Aquementas (11904, -3.24 DPS) [quest] |
| ranged | Wand of Biting Cold (19108) | Korrak the Bloodrager [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, -0.44 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -0.78 DPS) [dungeon]; Brilliant Wand (249385, -1.02 DPS) [crafted] |

**New at 60:** head: Spellpower Goggles Xtreme Plus; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Magician's Cord; legs: Sentinel's Silk Leggings; finger1: Blessed Band of Light; finger2: Maiden's Circle; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Wand of Biting Cold

No-known-source sample (15 of 1299, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (gnome, 230005203000000000-50000000000000000-253510130101301221)

Set DPS (verified): 679.1. Weights run: 2.3s. Verify run: 15.8s. 1299 eligible items had no known source.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.168 ± 0.008, crit=0.660 ± 0.020 per rating point (14 rating = 1%, 9.245 per %), hit=1.559 ± 0.071 per rating point (10 rating = 1%, 15.591 per %), spell_haste=10.391 ± 0.311, spell_penetration=not significant (0.000 ± 0.000), frost_power=0.815 ± 0.002, spirit=not significant (-0.000 ± 0.001), mp5=not significant (0.001 ± 0.001)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crimson Felt Hat (18727, +0.00 DPS) [dungeon]; Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Arcane Crystal Pendant (20037, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 28.4 spell_power points (14.49 DPS) | yes | Field Marshal's Silk Spaulders (231602, -0.46 DPS) [pvp]; Lieutenant Commander's Silk Mantle (227102, -1.19 DPS) [pvp]; Argent Shoulders (19059, -1.75 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -5.30 DPS, sim-verified) [world]; Shroud of Arcane Mastery (22330, -7.90 DPS) [dungeon]; Cloak of Flames (3475, -8.84 DPS) [world_drop] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Vestments (231603, -3.14 DPS) [pvp]; Freezing Lich Robes (14340, -4.35 DPS, sim-verified) [dungeon]; Robe of Winter Night (14136, -8.49 DPS) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 23.3 spell_power points (11.89 DPS) | yes | Runecloth Cuffs (254123, -5.43 DPS) [crafted]; Netherfroth Cuffs (254063, -5.90 DPS) [crafted]; Arena Wristguards (18709, -7.18 DPS) [world] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 29.2 spell_power points (14.86 DPS) | yes | Hands of Power (13253, +0.00 DPS) [dungeon]; Marshal's Silk Gloves (16440, -0.08 DPS) [vendor]; Marshal's Silk Gauntlets (231608, -0.08 DPS) [vendor] |
| waist | Highlander's Cloth Girdle (20047) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Magician's Cord (272393, -0.62 DPS) [vendor]; Stormpike Cloth Girdle (19094, -2.33 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 52.2 spell_power points (26.58 DPS) | yes | Marshal's Silk Leggings (231605, -4.87 DPS) [pvp]; Skyshroud Leggings (13170, -8.58 DPS, sim-verified) [dungeon]; Knight-Captain's Silk Legguards (227109, -9.72 DPS) [pvp] |
| feet | Sergeant Major's Dreadweave Boots (220891) | PvP reward · Alliance [vendor] | 25.1 spell_power points (12.79 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Earthen Silk Slippers (254013, -0.56 DPS) [crafted]; Knight-Lieutenant's Silk Walkers (227112, -1.23 DPS) [pvp] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -2.72 DPS) [dungeon]; Maiden's Circle (13001, -3.27 DPS) [world_drop]; Naglering (11669, -12.22 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -2.34 DPS) [dungeon]; Songstone of Ironforge (12543, -2.88 DPS) [quest]; Maiden's Circle (13001, -2.88 DPS) [world_drop] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -4.26 DPS, sim-verified) [vendor]; Burst of Knowledge (11832, -5.10 DPS) [dungeon]; Uther's Strength (11302, -8.15 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burst of Knowledge (11832, -1.04 DPS) [dungeon]; Uther's Strength (11302, -4.09 DPS) [world_drop]; Shard of the Splithooves (10659, -7.15 DPS) [quest] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -2.88 DPS) [pvp]; Staff of Jordan (873, -10.42 DPS, sim-verified) [world_drop]; Elemental Mage Staff (944, -11.81 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Biting Cold (19108) | Korrak the Bloodrager [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, -0.69 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -1.04 DPS) [dungeon]; Brilliant Wand (249385, -1.46 DPS) [crafted] |

**New at 60:** head: Spellpower Goggles Xtreme Plus; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Highlander's Cloth Girdle; legs: Sentinel's Silk Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blessed Band of Light; finger2: Don Julio's Band; trinket1: Weakness Analyzer; trinket2: Frozen Heart of the Mountain; main_hand: Ironbark Staff; ranged: Wand of Biting Cold

No-known-source sample (15 of 1299, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (undead, 000000000000000000-00000000000000000-253100000000000000)

Set DPS (verified): 47.8. Weights run: 1.8s. Verify run: 2.1s. 175 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.981 ± 0.007, crit=0.092 ± 0.003 per rating point (14 rating = 1%, 1.294 per %), hit=0.262 ± 0.008 per rating point (10 rating = 1%, 2.623 per %), spell_haste=0.663 ± 0.102, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002, spirit=not significant (0.000 ± 0.000), mp5=2.089 ± 0.105

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Silky Circlet (253951) | Tailoring [crafted] | 11.0 spell_power points (2.37 DPS) | yes | Pristine Circlet (253949, -1.06 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.31 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.8 spell_power points (2.98 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.98 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -1.63 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -2.12 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.20 DPS) [crafted]; Sanguine Cape (14376, -0.22 DPS) [world_drop]; Black Whelp Cloak (7283, -0.42 DPS) [crafted] |
| chest | Filigreed Silky Gown (253903) | Tailoring [crafted] | 13.9 spell_power points (2.99 DPS) | yes | Filigreed Pristine Gown (253901, -0.85 DPS, sim-verified) [crafted]; Gray Woolen Robe (2585, -1.08 DPS) [crafted]; Mystic's Wrap (14369, -1.51 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 11.2 spell_power points (2.40 DPS) | yes | Bright Bracers (3647, -0.74 DPS, sim-verified) [world_drop]; Tabitha's Cuffs (251486, -1.76 DPS) [quest]; Owlbeard Bracers (16981, -1.77 DPS) [quest] |
| hands | Silky Gloves (253915) | Tailoring [crafted] | 9.9 spell_power points (2.14 DPS) | yes | Pristine Gloves (253913, -0.65 DPS) [crafted]; Serpent Gloves (5970, -0.69 DPS, sim-verified) [dungeon]; Gnoll Casting Gloves (892, -0.85 DPS) [world] |
| waist | Silky Sash (253927) | Tailoring [crafted] | 10.9 spell_power points (2.35 DPS) | yes | Pristine Sash (253925, -0.63 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.66 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.86 DPS) [crafted] |
| legs | Filigreed Silky Leggings (253939) | Tailoring [crafted] | 16.9 spell_power points (3.63 DPS) | yes | Filigreed Pristine Leggings (253937, -1.06 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -2.13 DPS) [dungeon]; Darkweave Breeches (12987, -2.16 DPS) [world_drop] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.9 spell_power points (2.35 DPS) | yes | Silky Boots (253891, -0.54 DPS, sim-verified) [crafted]; Pristine Boots (253889, -1.07 DPS) [crafted]; Red Woolen Boots (4313, -1.49 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 9.2 spell_power points (1.98 DPS) | yes | Loop of Sacrifice (281673, -0.92 DPS) [quest]; Volcanic Rock Ring (12053, -1.34 DPS) [world_drop]; Furlfeather Ring (286745, -1.76 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.9 spell_power points (1.27 DPS) | yes | Loop of Sacrifice (281673, -0.33 DPS, sim-verified) [quest]; Volcanic Rock Ring (12053, -0.63 DPS) [world_drop]; Furlfeather Ring (286745, -1.05 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Staff of the Blessed Seer (2271, -0.42 DPS) [dungeon]; Crescent Staff (6505, -0.57 DPS, sim-verified) [quest]; Gnarled Necromancer's Staff (251534, -0.85 DPS) [quest] |
| off_hand | Antipodean Rod (2879) | World drop [world_drop] | 6.0 spell_power points (1.29 DPS) | yes | Seer's Fine Stein (7608, -0.66 DPS) [world_drop]; Spellbinder Orb (15926, -0.66 DPS) [dungeon]; Tork Wrench (11855, -0.87 DPS) [quest] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller (12984, -0.41 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -0.68 DPS) [quest]; Greater Magic Wand (11288, -0.72 DPS) [crafted] |

**New at 20:** head: Silky Circlet; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Silky Gown; wrist: Mindthrust Bracers; hands: Silky Gloves; waist: Silky Sash; legs: Filigreed Silky Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Evocator's Blade; off_hand: Antipodean Rod; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 175, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 23173 Abomination Skin Leggings; 209618 Insignia of the Alliance

### Band 30 (undead, 000000000000000000-00000000000000000-253510130100000000)

Set DPS (verified): 82.5. Weights run: 1.8s. Verify run: 1.8s. 301 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.300 ± 0.005, crit=0.111 ± 0.004 per rating point (14 rating = 1%, 1.551 per %), hit=0.207 ± 0.006 per rating point (10 rating = 1%, 2.068 per %), spell_haste=1.058 ± 0.060, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002, spirit=0.197 ± 0.037, mp5=0.288 ± 0.041

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Silky Circlet (253977) | Tailoring [crafted] | 14.0 spell_power points (2.88 DPS) | yes | Silky Circlet (253951, -0.62 DPS) [crafted]; Azure Silk Hood (7048, -0.82 DPS) [crafted]; Holy Shroud (2721, -0.91 DPS, sim-verified) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | sim-verified (82.5 DPS) | yes | Crystal Starfire Medallion (5003, -0.00 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.06 DPS) [vendor]; Kaleidoscope Chain (13084, -0.41 DPS, sim-verified) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.7 spell_power points (2.41 DPS) | yes | Death Speaker Mantle (6685, -0.49 DPS) [dungeon]; Invoker's Mantle (215365, -0.66 DPS) [crafted]; Magician's Mantle (12998, -0.82 DPS) [world_drop] |
| back | Azure Silk Cloak (7053) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, -1.15 DPS) [crafted]; Heavy Woolen Cloak (4311, -1.19 DPS) [crafted]; Windsong Drape (15468, -1.27 DPS, sim-verified) [quest] |
| chest | Silky Gown (253963) | Tailoring [crafted] | 15.9 spell_power points (3.27 DPS) | yes | Beguiler Robes (7728, -0.35 DPS) [dungeon]; Tree Bark Jacket (1486, -0.59 DPS) [dungeon]; Green Silk Armor (7065, -0.62 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.9 spell_power points (2.03 DPS) | yes | Glowing Magical Bracelets (13106, -1.54 DPS) [world_drop]; Nightsky Wristbands (6407, -1.54 DPS) [world_drop]; Tabitha's Cuffs (251486, -2.43 DPS, sim-verified) [quest] |
| hands | Azure Silk Gloves (4319) | Tailoring [crafted] | 10.6 spell_power points (2.18 DPS) | yes | Silky Gloves (253915, -0.55 DPS) [crafted]; Jutebraid Gloves (10654, -0.64 DPS) [quest]; Serpent Gloves (5970, -0.74 DPS) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.9 spell_power points (2.45 DPS) | yes | Belt of Arugal (6392, -0.21 DPS) [dungeon]; Invoker's Cord (215366, -0.70 DPS) [crafted]; Warsong Sash (16975, -0.99 DPS, sim-verified) [quest] |
| legs | Silky Leggings (253989) | Tailoring [crafted] | 17.1 spell_power points (3.52 DPS) | yes | Filigreed Silky Leggings (253939, -0.72 DPS) [crafted]; Azure Silk Pants (7046, -1.09 DPS) [crafted]; Gaze Dreamer Pants (6903, -2.69 DPS, sim-verified) [dungeon] |
| feet | Frothing Slippers (254003) | Tailoring [crafted] | 15.9 spell_power points (3.27 DPS) | yes | Gilded Slippers (254001, -1.52 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -1.58 DPS) [crafted]; Acidic Walkers (9454, -1.59 DPS) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.6 spell_power points (1.56 DPS) | yes | Ring of Power Regulation (273030, -0.15 DPS) [dungeon]; Solid Gold Band (279845, -0.53 DPS) [quest]; Electrocutioner Lagnut (9447, -0.66 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 7.5 spell_power points (1.54 DPS) | yes | Solid Gold Band (279845, -0.52 DPS) [quest]; Ring of Power Regulation (273030, -0.56 DPS, sim-verified) [dungeon]; Electrocutioner Lagnut (9447, -0.64 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Soulstaff (249393, -0.95 DPS) [crafted]; Tendonscraper Dagger (273053, -0.99 DPS) [world]; Glimmering Staff (249392, -2.22 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.8 spell_power points (1.81 DPS) | yes | Orb of Souls (249395, -0.74 DPS) [crafted]; Witch's Finger (16887, -1.32 DPS) [quest]; Antipodean Rod (2879, -2.12 DPS, sim-verified) [world_drop] |
| ranged | Greater Mystic Wand (217287) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderwood (13062, +0.00 DPS) [world_drop]; Wand of Eventide (5214, -0.33 DPS) [world_drop]; Starfaller (13063, -0.39 DPS) [world_drop] |

**New at 30:** head: Filigreed Silky Circlet; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Azure Silk Cloak; chest: Silky Gown; wrist: Spidertank Oilrag; hands: Azure Silk Gloves; waist: Defiler's Cloth Girdle; legs: Silky Leggings; feet: Frothing Slippers; finger1: Advisor's Ring; finger2: Snake Hoop; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Greater Mystic Wand

No-known-source sample (15 of 301, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (undead, 000000000000000000-00000000000000000-253510130101301221)

Set DPS (verified): 176.7. Weights run: 2.0s. Verify run: 2.3s. 392 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.684 ± 0.016, crit=0.164 ± 0.006 per rating point (14 rating = 1%, 2.297 per %), hit=0.343 ± 0.017 per rating point (10 rating = 1%, 3.430 per %), spell_haste=1.030 ± 0.198, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003, spirit=not significant (0.538 ± 0.144), mp5=0.734 ± 0.183

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Papal Fez (9431, +0.00 DPS) [dungeon]; Whitemane's Chapeau (7720, -0.18 DPS) [dungeon]; Electromagnetic Gigaflux Reactivator (9492, -0.69 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.5 spell_power points (3.20 DPS) | yes | Triune Amulet (7722, -0.18 DPS) [dungeon]; Necklace of Calisea (1714, -1.01 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.45 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.9 spell_power points (4.08 DPS) | yes | Green Silken Shoulders (7057, -0.09 DPS) [crafted]; Mantle of Doan (7712, -0.12 DPS) [dungeon]; Bloodmage Mantle (7684, -0.19 DPS) [dungeon] |
| back | Frost Leather Cloak (7377) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Azure Silk Cloak (7053, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.20 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.34 DPS) [vendor] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Robe of the Magi (1716, -0.74 DPS) [world_drop]; Robe of Power (7054, -1.33 DPS) [crafted]; Beguiler Robes (7728, -2.61 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 11.2 spell_power points (2.88 DPS) | yes | Radiant Silver Bracers (4545, -0.44 DPS) [quest]; Condor Bracers (15864, -1.08 DPS) [quest]; Earthen Silk Cuffs (254019, -1.16 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 24.5 spell_power points (6.29 DPS) | yes | Frothing Handwraps (254023, -1.47 DPS) [crafted]; Red Mageweave Gloves (10018, -1.71 DPS) [crafted]; Black Mageweave Gloves (10003, -2.44 DPS) [crafted] |
| waist | Frothing Cord (254039) | Tailoring [crafted] | 22.7 spell_power points (5.83 DPS) | yes | Deathmage Sash (10771, -1.40 DPS) [dungeon]; Defiler's Cloth Girdle (20166, -1.53 DPS) [rep]; Gilded Cord (254037, -1.54 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.2 spell_power points (5.70 DPS) | yes | Silky Leggings (253989, -0.19 DPS) [crafted]; Stoneweaver Leggings (9407, -0.84 DPS) [dungeon]; Filigreed Silky Leggings (253939, -1.27 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 27.8 spell_power points (7.13 DPS) | yes | Frothing Slippers (254003, -2.01 DPS) [crafted]; Gilded Slippers (254001, -3.55 DPS) [crafted]; Furen's Boots (13100, -3.69 DPS) [world_drop] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.4 spell_power points (3.45 DPS) | yes | Snake Hoop (6750, -1.03 DPS) [quest]; Ring of Power Regulation (273030, -1.34 DPS) [dungeon]; Welken Ring (5011, -1.60 DPS) [world_drop] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 11.2 spell_power points (2.88 DPS) | yes | Snake Hoop (6750, -0.46 DPS) [quest]; Ring of Power Regulation (273030, -0.77 DPS) [dungeon]; Welken Ring (5011, -1.03 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -7.72 DPS) [dungeon]; Celestial Stave (9517, -8.41 DPS) [quest]; Hypnotic Blade (7714, -8.51 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Icefury Wand (7514) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Freezing Shard (10572, -0.30 DPS) [dungeon]; Dancing Flame (6806, -1.25 DPS) [quest]; Jaina's Firestarter (13064, -1.32 DPS) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Frost Leather Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Frothing Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Staff of Jordan; ranged: Icefury Wand

No-known-source sample (15 of 392, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (undead, 230005000000000000-00000000000000000-253510130101301221)

Set DPS (verified): 258.7. Weights run: 2.0s. Verify run: 4.5s. 504 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.449 ± 0.015, crit=0.393 ± 0.012 per rating point (14 rating = 1%, 5.501 per %), hit=0.817 ± 0.034 per rating point (10 rating = 1%, 8.174 per %), spell_haste=5.065 ± 0.227, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002, spirit=not significant (0.164 ± 0.069), mp5=0.432 ± 0.102

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreamweave Circlet (10041) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Red Mageweave Headband (10033, +0.00 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -0.17 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -0.97 DPS) [vendor] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 18.7 spell_power points (7.05 DPS) | yes | Glowing Eye of Mordresh (10769, -3.06 DPS) [dungeon]; Mindburst Medallion (11196, -3.46 DPS) [quest]; Triune Amulet (7722, -3.60 DPS) [dungeon] |
| shoulder | Netherfroth Shoulders (254051) | Tailoring [crafted] | 21.6 spell_power points (8.14 DPS) | yes | Rotgrip Mantle (17732, -0.19 DPS) [dungeon]; Kentic Amice (11624, -0.28 DPS) [dungeon]; Blood Guard's Dreadweave Mantle (220905, -1.52 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -0.62 DPS) [quest]; Runecloth Cloak (13860, -1.92 DPS) [crafted]; Frost Leather Cloak (7377, -2.52 DPS) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frostweave Robe (13868, +0.00 DPS) [crafted]; Frostweave Tunic (13869, +0.00 DPS) [crafted]; Acumen Robes (17775, +0.00 DPS) [quest] |
| wrist | Netherfroth Cuffs (254063) | Tailoring [crafted] | 17.0 spell_power points (6.40 DPS) | yes | Nethergeld Cuffs (254061, -2.26 DPS) [crafted]; Spidertank Oilrag (9448, -2.52 DPS) [dungeon]; Bloodband Bracers (11469, -2.80 DPS) [quest] |
| hands | Frothing Gloves (254097) | Tailoring [crafted] | 24.8 spell_power points (9.36 DPS) | yes | Frostweave Gloves (13870, -0.42 DPS) [crafted]; Dreamweave Gloves (10019, -1.46 DPS) [crafted]; Runecloth Gloves (13863, -2.75 DPS) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 25.1 spell_power points (9.47 DPS) | yes | Frothing Waistcord (254083, -0.78 DPS) [crafted]; Frothing Cord (254039, -2.47 DPS) [crafted]; Satyrmane Sash (17755, -2.50 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | 27.5 spell_power points (10.37 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -1.74 DPS) [vendor]; Red Mageweave Pants (10009, -3.06 DPS) [crafted]; Wizardweave Leggings (14132, -3.20 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.2 spell_power points (9.49 DPS) | yes | Frothing Sandals (254109, -0.42 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -1.86 DPS) [vendor]; Frothing Slippers (254003, -3.15 DPS) [crafted] |
| finger1 | Freezing Band (942) | World drop [world_drop] | 21.0 spell_power points (7.92 DPS) | yes | Band of the Unicorn (7553, -3.02 DPS) [world_drop]; Cyclopean Band (11824, -3.09 DPS) [dungeon]; Philanthropist's Ring (281635, -3.30 DPS) [quest] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 13.7 spell_power points (5.18 DPS) | yes | Band of the Unicorn (7553, -0.27 DPS) [world_drop]; Cyclopean Band (11824, -0.35 DPS) [dungeon]; Philanthropist's Ring (281635, -0.56 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -4.53 DPS) [dungeon]; Inventor's Focal Sword (17719, -6.56 DPS) [dungeon]; Spire of Hakkar (10844, -6.89 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Icefury Wand (7514) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lesser Eternal Wand (249232, +0.00 DPS) [crafted]; Noxious Shooter (17745, -0.16 DPS) [dungeon]; Freezing Shard (10572, -0.34 DPS) [dungeon] |

**New at 50:** head: Dreamweave Circlet; neck: Arcane Crystal Pendant; shoulder: Netherfroth Shoulders; back: Spritecaster Cape; wrist: Netherfroth Cuffs; hands: Frothing Gloves; waist: Ban'thok Sash; legs: Spellshock Leggings; finger1: Freezing Band; finger2: Advisor's Ring; trinket2: Darkspear Voodoo Seal

No-known-source sample (15 of 504, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (undead, 230005203000000000-50000000000000000-253510130101301221)

Set DPS (verified): 407.1. Weights run: 2.2s. Verify run: 15.7s. 1293 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.438 ± 0.027, crit=0.343 ± 0.011 per rating point (14 rating = 1%, 4.795 per %), hit=0.857 ± 0.044 per rating point (10 rating = 1%, 8.567 per %), spell_haste=3.422 ± 0.372, spell_penetration=not significant (0.000 ± 0.000), frost_power=0.806 ± 0.002, spirit=not significant (1.086 ± 0.320), mp5=not significant (1.375 ± 0.360)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ghostweave Hood (254135) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crimson Felt Hat (18727, +0.00 DPS) [dungeon]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -0.11 DPS) [pvp] |
| neck | Diana's Pearl Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Lady Maye's Pendant (14558, -0.59 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -0.73 DPS) [world_drop]; Star of Mystaria (12103, -2.73 DPS, sim-verified) [dungeon] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 32.6 spell_power points (10.24 DPS) | yes | Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Champion's Silk Mantle (227104, -1.15 DPS) [pvp]; Ghostweave Mantle (254127, -3.77 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -1.55 DPS) [dungeon]; Crystalline Threaded Cape (20697, -1.98 DPS) [world]; Cloak of Flames (3475, -7.43 DPS, sim-verified) [world_drop] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Raiment (231596, +0.00 DPS) [pvp]; Ghostweave Vest (14141, -2.61 DPS) [crafted]; Alanna's Embrace (13314, -10.53 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 33.1 spell_power points (10.40 DPS) | yes | Wyrmthalak's Shackles (13958, -4.04 DPS) [quest]; General's Silk Cuffs (16538, -4.37 DPS) [pvp]; Netherfroth Cuffs (254063, -4.44 DPS) [crafted] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 35.1 spell_power points (11.04 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Shivery Handwraps (18693, -4.38 DPS, sim-verified) [dungeon] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 29.5 spell_power points (9.27 DPS) | yes | Wisdom of the Timbermaw (19047, -0.82 DPS) [crafted]; General's Silk Sash (16537, -0.99 DPS) [pvp]; Thuzadin Sash (18740, -3.36 DPS, sim-verified) [dungeon] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Silk Trousers (231595, +0.00 DPS) [pvp]; Magister's Leggings (16687, -0.97 DPS) [dungeon]; Ghostweave Pants (14144, -2.66 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.29 DPS) [crafted]; Rime-encrusted Boots (279268, -1.20 DPS, sim-verified) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.68 DPS) [quest]; Channeler's Ring (272406, -1.32 DPS) [vendor]; Naglering (11669, -8.54 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -1.22 DPS) [vendor]; The Postmaster's Seal (13392, -2.35 DPS) [dungeon]; Eye of Orgrimmar (12545, -3.39 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+7.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -2.82 DPS) [world_drop]; Burst of Knowledge (11832, -3.14 DPS) [dungeon]; Serenity Field (272439, -3.26 DPS, sim-verified) [vendor] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -1.42 DPS) [pvp]; Mindfang (20214, -3.89 DPS) [rep]; Ironbark Staff (20220, -13.17 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (9.11 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; Tome of the Ice Lord (19310, -0.16 DPS) [rep]; High Warlord's Tome of Mending (234564, -2.47 DPS) [pvp] |
| ranged | Wand of Biting Cold (19108) | The Legend of Korrak [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Torch of Light (279246, +0.00 DPS) [crafted]; Bonecreeper Stylus (13938, -0.05 DPS) [dungeon]; Brilliant Wand (249385, -0.28 DPS) [crafted] |

**New at 60:** head: Ghostweave Hood; neck: Diana's Pearl Necklace; shoulder: Argent Shoulders; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Magician's Cord; legs: Outrider's Silk Leggings; finger1: Maiden's Circle; finger2: Blessed Band of Light; trinket1: Frozen Heart of the Mountain; trinket2: Weakness Analyzer; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Wand of Biting Cold

No-known-source sample (15 of 1293, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60, raid preset (undead, 230005203000000000-50000000000000000-253510130101301221)

Set DPS (verified): 675.6. Weights run: 2.3s. Verify run: 16.8s. 1293 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.207 ± 0.011, crit=0.470 ± 0.014 per rating point (14 rating = 1%, 6.581 per %), hit=1.046 ± 0.049 per rating point (10 rating = 1%, 10.464 per %), spell_haste=6.307 ± 0.230, spell_penetration=not significant (0.000 ± 0.000), frost_power=0.810 ± 0.002, spirit=not significant (0.045 ± 0.023), mp5=not significant (0.066 ± 0.079)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 32.0 spell_power points (14.61 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -0.20 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -2.83 DPS, sim-verified) [crafted] |
| neck | Diana's Pearl Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 21.1 spell_power points (9.64 DPS) | yes | Archlight Talisman (15856, -3.22 DPS) [quest]; Star of Mystaria (12103, -3.85 DPS) [dungeon]; Arcane Crystal Pendant (20037, -7.01 DPS, sim-verified) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 26.3 spell_power points (11.99 DPS) | yes | Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Argent Shoulders (19059, -0.44 DPS) [crafted]; Boreal Mantle (11782, -0.52 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -5.75 DPS) [dungeon]; Frostwolf Advisor's Cloak (19085, -6.45 DPS) [rep]; Crystalline Threaded Cape (20697, -8.28 DPS, sim-verified) [world] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Raiment (231596, -2.60 DPS) [pvp]; Robe of Winter Night (14136, -5.09 DPS, sim-verified) [crafted]; Freezing Lich Robes (14340, -6.51 DPS) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 24.0 spell_power points (10.94 DPS) | yes | Runecloth Cuffs (254123, -4.98 DPS) [crafted]; Netherfroth Cuffs (254063, -5.37 DPS) [crafted]; Spidertank Oilrag (9448, -6.74 DPS) [dungeon] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Hands of Power (13253, -1.46 DPS, sim-verified) [dungeon] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Defiler's Cloth Girdle (20163, -0.98 DPS) [rep]; Frostwolf Cloth Belt (19090, -1.78 DPS) [rep] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skyshroud Leggings (13170, +0.00 DPS) [dungeon]; General's Silk Trousers (231595, +0.00 DPS) [pvp]; Legionnaire's Silk Legguards (227107, -0.48 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.3 spell_power points (11.10 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Blood Guard's Silk Walkers (227110, -0.57 DPS) [pvp]; Omnicast Boots (11822, -0.83 DPS) [dungeon] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.54 DPS) [quest]; Don Mauricio's Band of Domination (22433, -1.48 DPS) [dungeon]; Naglering (11669, -14.88 DPS, sim-verified) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.04 DPS) [quest]; Don Mauricio's Band of Domination (22433, -0.98 DPS) [dungeon]; Don Julio's Band (19325, -1.22 DPS) [rep] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -4.00 DPS, sim-verified) [vendor]; Burst of Knowledge (11832, -4.56 DPS) [dungeon]; Rune of the Guard Captain (19120, -6.70 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Uther's Strength (11302, -1.56 DPS) [world_drop]; Rune of the Guard Captain (19120, -2.27 DPS, sim-verified) [quest] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -2.44 DPS) [pvp]; Staff of Jordan (873, -10.94 DPS) [world_drop]; Elemental Mage Staff (944, -12.47 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Biting Cold (19108) | The Legend of Korrak [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, -0.51 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -0.89 DPS) [dungeon]; Brilliant Wand (249385, -1.14 DPS) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Diana's Pearl Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Magician's Cord; legs: Outrider's Silk Leggings; finger1: Blessed Band of Light; finger2: Maiden's Circle; trinket1: Weakness Analyzer; trinket2: Frozen Heart of the Mountain; main_hand: Ironbark Staff; ranged: Wand of Biting Cold

No-known-source sample (15 of 1293, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

