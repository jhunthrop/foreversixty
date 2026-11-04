# Leveling BiS: Frost

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-253100000000000000)

Set DPS (verified): 48.8. Weights run: 2.7s. Verify run: 2.6s. 178 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.978 ± 0.007, crit=0.093 ± 0.003 per rating point (14 rating = 1%, 1.295 per %), hit=0.263 ± 0.008 per rating point (10 rating = 1%, 2.631 per %), spell_haste=0.623 ± 0.103, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002, spirit=not significant (0.000 ± 0.000), mp5=2.134 ± 0.106

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Silky Circlet (253951) | Tailoring [crafted] | 11.0 spell_power points (2.37 DPS) | yes | Pristine Circlet (253949, -1.08 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.32 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.8 spell_power points (2.98 DPS) | yes | Reinforced Woolen Shoulders (4315, -1.17 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -1.60 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -2.11 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Sanguine Cape (14376, -0.02 DPS) [world_drop]; Pearl-clasped Cloak (5542, -0.21 DPS, sim-verified) [crafted]; Black Whelp Cloak (7283, -0.22 DPS) [crafted] |
| chest | Filigreed Silky Gown (253903) | Tailoring [crafted] | 13.9 spell_power points (3.00 DPS) | yes | Filigreed Pristine Gown (253901, -0.86 DPS, sim-verified) [crafted]; Gray Woolen Robe (2585, -1.08 DPS) [crafted]; Mystic's Wrap (14369, -1.52 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 11.3 spell_power points (2.44 DPS) | yes | Bright Bracers (3647, -0.41 DPS, sim-verified) [world_drop]; Mystic's Bracelets (14366, -2.01 DPS) [world_drop]; Repurposed Hair Band (281256, -2.01 DPS) [quest] |
| hands | Silky Gloves (253915) | Tailoring [crafted] | 9.9 spell_power points (2.14 DPS) | yes | Serpent Gloves (5970, -0.38 DPS, sim-verified) [dungeon]; Pristine Gloves (253913, -0.65 DPS) [crafted]; Gnoll Casting Gloves (892, -0.85 DPS) [world] |
| waist | Silky Sash (253927) | Tailoring [crafted] | 10.9 spell_power points (2.35 DPS) | yes | Pristine Sash (253925, -0.65 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.67 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.86 DPS) [crafted] |
| legs | Filigreed Silky Leggings (253939) | Tailoring [crafted] | 16.9 spell_power points (3.64 DPS) | yes | Filigreed Pristine Leggings (253937, -1.08 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -2.13 DPS) [dungeon]; Darkweave Breeches (12987, -2.16 DPS) [world_drop] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.9 spell_power points (2.35 DPS) | yes | Silky Boots (253891, -0.38 DPS, sim-verified) [crafted]; Pristine Boots (253889, -1.07 DPS) [crafted]; Red Woolen Boots (4313, -1.49 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 9.3 spell_power points (2.00 DPS) | yes | Lavishly Jeweled Ring (1156, -0.73 DPS) [dungeon]; Volcanic Rock Ring (12053, -1.37 DPS) [world_drop]; Furlfeather Ring (286745, -1.78 DPS) [world] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 7.0 spell_power points (1.50 DPS) | yes | Lavishly Jeweled Ring (1156, -0.30 DPS, sim-verified) [dungeon]; Volcanic Rock Ring (12053, -0.87 DPS) [world_drop]; Furlfeather Ring (286745, -1.28 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | 30.0 spell_power points (6.46 DPS) | yes | Emberstone Staff (5201, -1.28 DPS) [dungeon]; Staff of the Blessed Seer (2271, -1.45 DPS, sim-verified) [dungeon]; Buzzer Blade (2169, -1.72 DPS) [dungeon] |
| off_hand | Antipodean Rod (2879) | World drop [world_drop] | 6.0 spell_power points (1.29 DPS) | yes | Mathystra Channeling Orb (281291, -0.24 DPS) [quest]; Tear of Grief (5611, -0.43 DPS, sim-verified) [quest]; Dwarven Tome (279898, -0.65 DPS) [quest] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller (12984, -0.47 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -0.68 DPS) [quest]; Greater Magic Wand (11288, -0.71 DPS) [crafted] |

**New at 20:** head: Silky Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Silky Gown; wrist: Mindthrust Bracers; hands: Silky Gloves; waist: Silky Sash; legs: Filigreed Silky Leggings; feet: Spidersilk Boots; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Evocator's Blade; off_hand: Antipodean Rod; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 178, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18850 Insignia of the Horde; 20434 Lorekeeper's Staff; 20440 Protector's Sword

### Band 30 (gnome, 000000000000000000-00000000000000000-253510130100000000)

Set DPS (verified): 82.7. Weights run: 2.6s. Verify run: 2.2s. 303 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.301 ± 0.005, crit=0.111 ± 0.004 per rating point (14 rating = 1%, 1.552 per %), hit=0.208 ± 0.006 per rating point (10 rating = 1%, 2.076 per %), spell_haste=1.103 ± 0.061, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002, spirit=0.179 ± 0.038, mp5=0.287 ± 0.042

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Silky Circlet (253977) | Tailoring [crafted] | 14.0 spell_power points (2.89 DPS) | yes | Holy Shroud (2721, -0.40 DPS) [world_drop]; Silky Circlet (253951, -0.62 DPS) [crafted]; Azure Silk Hood (7048, -0.83 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 1.9 spell_power points (0.40 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.02 DPS) [world]; Crystal Starfire Medallion (5003, -0.04 DPS) [world_drop]; Seal of the Expedition (271920, -0.09 DPS) [rep] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.7 spell_power points (2.42 DPS) | yes | Death Speaker Mantle (6685, -0.51 DPS, sim-verified) [dungeon]; Invoker's Mantle (215365, -0.66 DPS) [crafted]; Magician's Mantle (12998, -0.83 DPS) [world_drop] |
| back | Azure Silk Cloak (7053) | Tailoring [crafted] | sim-verified (82.7 DPS) | yes | Hillman's Cloak (3719, -1.14 DPS) [crafted]; Heavy Woolen Cloak (4311, -1.20 DPS) [crafted]; Vine Pruner's Cloak (279835, -1.61 DPS, sim-verified) [quest] |
| chest | Silky Gown (253963) | Tailoring [crafted] | 15.8 spell_power points (3.26 DPS) | yes | Beguiler Robes (7728, -0.37 DPS) [dungeon]; Tree Bark Jacket (1486, -0.58 DPS) [dungeon]; Green Silk Armor (7065, -0.60 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.9 spell_power points (2.03 DPS) | yes | Mindthrust Bracers (1974, -1.55 DPS) [dungeon]; Nightsky Wristbands (6407, -1.55 DPS) [world_drop]; Glowing Magical Bracelets (13106, -2.02 DPS, sim-verified) [world_drop] |
| hands | Azure Silk Gloves (4319) | Tailoring [crafted] | 10.5 spell_power points (2.17 DPS) | yes | Silky Gloves (253915, -0.63 DPS, sim-verified) [crafted]; Serpent Gloves (5970, -0.73 DPS) [dungeon]; Gloves of Meditation (4318, -0.88 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.9 spell_power points (2.46 DPS) | yes | Belt of Arugal (6392, -0.23 DPS) [dungeon]; Invoker's Cord (215366, -0.70 DPS) [crafted]; Silky Sash (253927, -0.76 DPS) [crafted] |
| legs | Silky Leggings (253989) | Tailoring [crafted] | 17.0 spell_power points (3.51 DPS) | yes | Filigreed Silky Leggings (253939, -0.72 DPS) [crafted]; Azure Silk Pants (7046, -1.07 DPS) [crafted]; Gaze Dreamer Pants (6903, -1.61 DPS, sim-verified) [dungeon] |
| feet | Frothing Slippers (254003) | Tailoring [crafted] | 15.8 spell_power points (3.26 DPS) | yes | Spidersilk Boots (4320, -1.57 DPS) [crafted]; Acidic Walkers (9454, -1.59 DPS) [dungeon]; Gilded Slippers (254001, -1.60 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.6 spell_power points (1.56 DPS) | yes | Ring of Power Regulation (273030, -0.15 DPS) [dungeon]; Minor Channeling Ring (1449, -0.41 DPS) [quest]; Electrocutioner Lagnut (9447, -0.69 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 7.5 spell_power points (1.55 DPS) | yes | Ring of Power Regulation (273030, -0.13 DPS) [dungeon]; Minor Channeling Ring (1449, -0.39 DPS) [quest]; Electrocutioner Lagnut (9447, -0.67 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-verified (82.7 DPS) | yes | Hydrocane (9452, -0.97 DPS) [dungeon]; Tendonscraper Dagger (273053, -0.97 DPS) [world]; Glimmering Staff (249392, -3.11 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.8 spell_power points (1.82 DPS) | yes | Orb of Souls (249395, -0.77 DPS) [crafted]; Eye of Paleth (2943, -0.99 DPS) [quest]; Antipodean Rod (2879, -1.26 DPS, sim-verified) [world_drop] |
| ranged | Thunderwood (13062) | World drop [world_drop] | sim-verified (82.7 DPS) | yes | Greater Mystic Wand (217287, -0.06 DPS) [crafted]; Wand of Eventide (5214, -0.31 DPS) [world_drop]; Starfaller (13063, -0.55 DPS) [world_drop] |

**New at 30:** head: Filigreed Silky Circlet; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Azure Silk Cloak; chest: Silky Gown; wrist: Spidertank Oilrag; hands: Azure Silk Gloves; waist: Highlander's Cloth Girdle; legs: Silky Leggings; feet: Frothing Slippers; finger1: Lorekeeper's Ring; finger2: Snake Hoop; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Thunderwood

No-known-source sample (15 of 303, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 40 (gnome, 000000000000000000-00000000000000000-253510130101301221)

Set DPS (verified): 195.5. Weights run: 2.7s. Verify run: 3.7s. 387 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.544 ± 0.012, crit=0.134 ± 0.005 per rating point (14 rating = 1%, 1.872 per %), hit=0.285 ± 0.014 per rating point (10 rating = 1%, 2.852 per %), spell_haste=1.264 ± 0.166, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004, spirit=not significant (0.325 ± 0.114), mp5=0.625 ± 0.153

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.57 DPS) | yes | Augural Shroud (2620, -1.43 DPS) [world]; Whitemane's Chapeau (7720, -1.46 DPS) [dungeon]; Papal Fez (9431, -2.83 DPS, sim-verified) [dungeon] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Glowing Eye of Mordresh (10769, -0.53 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -1.31 DPS) [dungeon]; Seal of the Expedition (271919, -2.02 DPS) [rep] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.1 spell_power points (4.40 DPS) | yes | Green Silken Shoulders (7057, -0.03 DPS) [crafted]; Bloodmage Mantle (7684, -0.05 DPS) [dungeon]; Azure Shoulders (7060, -0.33 DPS) [crafted] |
| back | Frost Leather Cloak (7377) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Azure Silk Cloak (7053, -0.01 DPS) [crafted]; Icy Cloak (4327, -0.64 DPS) [crafted]; Long Silken Cloak (4326, -0.71 DPS) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 27.4 spell_power points (8.58 DPS) | yes | Robe of Power (7054, -1.35 DPS) [crafted]; Robe of the Magi (1716, -1.71 DPS, sim-verified) [world_drop]; Elemental Raiment (9434, -2.02 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 10.9 spell_power points (3.40 DPS) | yes | Arcane Runed Bracers (4744, -0.59 DPS) [quest]; Condor Bracers (15864, -1.21 DPS) [quest]; Earthen Silk Cuffs (254019, -1.64 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.5 spell_power points (7.02 DPS) | yes | Red Mageweave Gloves (10018, -1.88 DPS) [crafted]; Black Mageweave Gloves (10003, -2.33 DPS) [crafted]; Frothing Handwraps (254023, -3.32 DPS, sim-verified) [crafted] |
| waist | Frothing Cord (254039) | Tailoring [crafted] | 20.3 spell_power points (6.35 DPS) | yes | Deathmage Sash (10771, -1.61 DPS) [dungeon]; Star Belt (4329, -1.88 DPS) [crafted]; Highlander's Cloth Girdle (20098, -2.83 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 20.5 spell_power points (6.42 DPS) | yes | Silky Leggings (253989, -0.34 DPS) [crafted]; Filigreed Silky Leggings (253939, -1.55 DPS) [crafted]; Stoneweaver Leggings (9407, -1.65 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 26.3 spell_power points (8.22 DPS) | yes | Frothing Slippers (254003, -3.98 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -4.43 DPS) [crafted]; Acidic Walkers (9454, -4.89 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.7 spell_power points (3.98 DPS) | yes | Snake Hoop (6750, -1.25 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.48 DPS) [quest]; Ring of Power Regulation (273030, -1.51 DPS) [dungeon] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 10.9 spell_power points (3.40 DPS) | yes | Ring of Forlorn Spirits (2043, -0.90 DPS) [quest]; Ring of Power Regulation (273030, -0.94 DPS) [dungeon]; Snake Hoop (6750, -1.56 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Celestial Stave (9517, -8.78 DPS) [quest]; Hypnotic Blade (7714, -9.10 DPS) [dungeon]; Illusionary Rod (7713, -13.45 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Icefury Wand (7514) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Twisted Nether Wand (249144, -1.55 DPS) [crafted]; Fizzle's Zippy Lighter (6729, -1.76 DPS) [quest]; Burning Sliver (5249, -1.97 DPS, sim-verified) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Triune Amulet; shoulder: Inquisitor's Shawl; back: Frost Leather Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Frothing Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Staff of Jordan; ranged: Icefury Wand

No-known-source sample (15 of 387, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 230005000000000000-00000000000000000-253510130101301221)

Set DPS (verified): 275.7. Weights run: 2.8s. Verify run: 6.2s. 491 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.548 ± 0.023, crit=0.244 ± 0.009 per rating point (14 rating = 1%, 3.420 per %), hit=0.566 ± 0.030 per rating point (10 rating = 1%, 5.661 per %), spell_haste=2.104 ± 0.354, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004, spirit=not significant (0.854 ± 0.312), mp5=not significant (0.899 ± 0.335)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreamweave Circlet (10041) | Tailoring [crafted] | 36.7 spell_power points (10.67 DPS) | yes | Red Mageweave Headband (10033, -1.97 DPS) [crafted]; Gemburst Circlet (10751, -2.33 DPS) [quest]; Spellpower Goggles Xtreme Plus (15999, -2.83 DPS) [crafted] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 14.0 spell_power points (4.06 DPS) | yes | Mindburst Medallion (11196, -0.12 DPS) [quest]; Glowing Eye of Mordresh (10769, -0.75 DPS) [dungeon]; Horizon Choker (13085, -0.84 DPS) [world_drop] |
| shoulder | Netherfroth Shoulders (254051) | Tailoring [crafted] | 26.5 spell_power points (7.70 DPS) | yes | Rotgrip Mantle (17732, -1.06 DPS) [dungeon]; Nethergeld Shoulders (254049, -2.32 DPS) [crafted]; Mantle of Doan (7712, -2.73 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Azure Silk Cloak (7053, -2.86 DPS) [crafted]; Runecloth Cloak (13860, -2.86 DPS, sim-verified) [crafted]; Caretaker's Cape (19531, -2.99 DPS) [rep] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 34.9 spell_power points (10.13 DPS) | yes | Frostweave Robe (13868, -0.44 DPS) [crafted]; Ghostweave Vest (14141, -0.47 DPS) [crafted]; Frostweave Tunic (13869, -1.41 DPS) [crafted] |
| wrist | Netherfroth Cuffs (254063) | Tailoring [crafted] | 21.1 spell_power points (6.13 DPS) | yes | Bloodband Bracers (11469, -2.50 DPS) [quest]; Nethergeld Cuffs (254061, -2.73 DPS, sim-verified) [crafted]; Spidertank Oilrag (9448, -2.73 DPS) [dungeon] |
| hands | Frothing Gloves (254097) | Tailoring [crafted] | 31.3 spell_power points (9.09 DPS) | yes | Dreamweave Gloves (10019, -1.49 DPS) [crafted]; Runecloth Gloves (13863, -1.94 DPS) [crafted]; Frostweave Gloves (13870, -2.04 DPS) [crafted] |
| waist | Frothing Waistcord (254083) | Tailoring [crafted] | 23.9 spell_power points (6.95 DPS) | yes | Frothing Cord (254039, -0.12 DPS) [crafted]; Ghostweave Cord (254073, -0.16 DPS) [crafted]; Satyrmane Sash (17755, -1.29 DPS) [dungeon] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 31.2 spell_power points (9.06 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -0.86 DPS) [dungeon]; Dalewind Trousers (13008, -1.43 DPS) [world_drop]; Silky Leggings (253989, -2.64 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 30.0 spell_power points (8.70 DPS) | yes | Frothing Sandals (254109, -1.46 DPS) [crafted]; Frothing Slippers (254003, -2.82 DPS) [crafted]; Vinerot Sandals (17748, -3.05 DPS) [dungeon] |
| finger1 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 15.6 spell_power points (4.53 DPS) | yes | Philanthropist's Ring (281635, -0.83 DPS) [quest]; Choking Band (11868, -1.30 DPS) [quest]; Eye of Adaegus (5266, -6.55 DPS, sim-verified) [world_drop] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Philanthropist's Ring (281635, -0.07 DPS) [quest]; Choking Band (11868, -0.55 DPS) [quest]; Eye of Adaegus (5266, -1.89 DPS, sim-verified) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Enriched Seal (271903, +0.00 DPS) [rep] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spire of Hakkar (10844, -3.58 DPS) [world]; Glowing Brightwood Staff (812, -4.18 DPS) [world_drop]; Kindling Stave (11750, -9.49 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Icefury Wand (7514) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lesser Eternal Wand (249232, -0.43 DPS) [crafted]; Cairnstone Sliver (9654, -0.89 DPS) [quest]; Dreambough Wand (249234, -0.98 DPS) [crafted] |

**New at 50:** head: Dreamweave Circlet; shoulder: Netherfroth Shoulders; back: Spritecaster Cape; wrist: Netherfroth Cuffs; hands: Frothing Gloves; waist: Frothing Waistcord; legs: Senior Designer's Pantaloons; finger1: Lorekeeper's Ring; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength

No-known-source sample (15 of 491, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 230005200000000000-51020000000000000-253510130101301221)

Set DPS (verified): 418.9. Weights run: 3.2s. Verify run: 16.9s. 1037 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.305 ± 0.027, crit=0.508 ± 0.016 per rating point (14 rating = 1%, 7.112 per %), hit=1.169 ± 0.064 per rating point (10 rating = 1%, 11.688 per %), spell_haste=4.892 ± 0.368, spell_penetration=not significant (0.000 ± 0.000), frost_power=0.847 ± 0.002, spirit=not significant (0.504 ± 0.187), mp5=not significant (0.713 ± 0.241)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreamweave Circlet (10041) | Tailoring [crafted] | 30.1 spell_power points (14.27 DPS) | yes | Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Ghostweave Hood (254135, -0.01 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 15.1 spell_power points (7.14 DPS) | yes | Mindburst Medallion (11196, -2.23 DPS) [quest]; Glowing Eye of Mordresh (10769, -2.66 DPS) [dungeon]; Triune Amulet (7722, -3.17 DPS, sim-verified) [dungeon] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 28.5 spell_power points (13.53 DPS) | yes | Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -0.21 DPS) [crafted]; Lieutenant Commander's Silk Mantle (227102, -0.49 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (418.9 DPS) | yes | Stormpike Sage's Cloak (19086, -2.30 DPS) [rep]; Caretaker's Cape (19530, -2.76 DPS) [rep]; Cloak of Flames (3475, -3.40 DPS) [world_drop] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-verified (418.9 DPS) | yes | Field Marshal's Silk Vestments (231603, -1.16 DPS) [pvp]; Robe of Winter Night (14136, -6.58 DPS, sim-verified) [crafted]; Knight-Captain's Silk Tunic (227108, -7.09 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 28.0 spell_power points (13.26 DPS) | yes | Netherfroth Cuffs (254063, -5.83 DPS) [crafted]; Runecloth Cuffs (254123, -6.60 DPS) [crafted]; Earthenweave Cuffs (254131, -7.08 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 26.3 spell_power points (12.46 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Frothing Gloves (254097, +0.00 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20047) | The League of Arathor [rep] | 22.9 spell_power points (10.88 DPS) | yes | Stormpike Cloth Girdle (19094, -0.90 DPS) [rep]; Wisdom of the Timbermaw (19047, -1.48 DPS) [crafted]; Ghostweave Cord (254073, -1.61 DPS) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 49.3 spell_power points (23.36 DPS) | yes | Marshal's Silk Leggings (231605, -0.48 DPS) [pvp]; Ghostweave Pants (14144, -6.25 DPS) [crafted]; Knight-Captain's Silk Legguards (227109, -6.38 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 27.5 spell_power points (13.05 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.71 DPS) [crafted]; Rime-encrusted Boots (279268, -0.72 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 23.2 spell_power points (10.98 DPS) | yes | Lorekeeper's Ring (19522, -2.99 DPS) [rep]; Deep Rooted Ring (19109, -3.85 DPS) [world]; Frigid Ring (18679, -3.91 DPS) [world] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | 18.8 spell_power points (8.91 DPS) | yes | Lorekeeper's Ring (19522, -0.92 DPS) [rep]; Deep Rooted Ring (19109, -1.78 DPS) [world]; Frigid Ring (18679, -1.84 DPS) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (418.9 DPS) | yes | Ankh of Life (1713, -2.12 DPS) [world_drop]; Uther's Strength (11302, -2.14 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -2.60 DPS) [quest] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (418.9 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Enriched Seal (271903, +0.00 DPS) [rep] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-verified (418.9 DPS) | yes | Grand Marshal's Stave (234571, -2.50 DPS) [pvp]; Sageclaw (20070, -5.88 DPS) [rep]; Ironbark Staff (20069, -14.04 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (13.75 DPS) | yes | Grand Marshal's Tome of Power (234589, -0.01 DPS) [pvp]; Tome of the Ice Lord (19310, -0.14 DPS) [rep]; Spirit of Aquementas (11904, -3.79 DPS) [quest] |
| ranged | Wand of Biting Cold (19108) | Korrak the Bloodrager [quest] | sim-verified (418.9 DPS) | yes | Torch of Light (279246, -0.76 DPS) [crafted]; Brilliant Wand (249385, -1.21 DPS) [crafted]; Greater Eternal Wand (249237, -2.32 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Highlander's Cloth Girdle; legs: Sentinel's Silk Leggings; finger1: Maiden's Circle; finger2: Don Julio's Band; trinket2: Shard of the Splithooves; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Wand of Biting Cold

No-known-source sample (15 of 1037, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60, raid preset (gnome, 230005200000000000-51020000000000000-253510130101301221)

Set DPS (verified): 693.0. Weights run: 3.4s. Verify run: 16.4s. 1037 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.195 ± 0.012, crit=0.562 ± 0.017 per rating point (14 rating = 1%, 7.869 per %), hit=1.465 ± 0.083 per rating point (10 rating = 1%, 14.651 per %), spell_haste=7.953 ± 0.439, spell_penetration=not significant (0.000 ± 0.000), frost_power=0.853 ± 0.002, spirit=not significant (0.152 ± 0.054), mp5=not significant (0.280 ± 0.089)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (16.42 DPS) | yes | Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Dreamweave Circlet (10041, -1.35 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 14.0 spell_power points (8.48 DPS) | yes | Triune Amulet (7722, -2.97 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -3.28 DPS) [dungeon]; Mindburst Medallion (11196, -3.66 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 27.4 spell_power points (16.67 DPS) | yes | Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Argent Shoulders (19059, -0.82 DPS) [crafted]; Lieutenant Commander's Silk Mantle (227102, -1.08 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (693.0 DPS) | yes | Cloak of Flames (3475, -0.87 DPS) [world_drop]; Stormpike Sage's Cloak (19086, -1.27 DPS) [rep]; Runecloth Cloak (13860, -3.36 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-verified (693.0 DPS) | yes | Field Marshal's Silk Vestments (231603, -3.11 DPS) [pvp]; Robe of Winter Night (14136, -8.57 DPS, sim-verified) [crafted]; Knight-Captain's Silk Tunic (227108, -10.50 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 24.6 spell_power points (14.97 DPS) | yes | Netherfroth Cuffs (254063, -6.93 DPS) [crafted]; Runecloth Cuffs (254123, -7.10 DPS) [crafted]; Spidertank Oilrag (9448, -8.99 DPS) [dungeon] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 26.7 spell_power points (16.23 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Earth Warder's Gloves (21318, +0.00 DPS) [quest]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Highlander's Cloth Girdle (20047) | The League of Arathor [rep] | 23.0 spell_power points (14.01 DPS) | yes | Stormpike Cloth Girdle (19094, -1.88 DPS) [rep]; Frothing Waistcord (254083, -3.08 DPS) [crafted]; Satyrmane Sash (17755, -4.31 DPS) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 49.7 spell_power points (30.22 DPS) | yes | Marshal's Silk Leggings (231605, -3.89 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -10.18 DPS) [pvp]; Ghostweave Pants (14144, -12.56 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.1 spell_power points (15.24 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Sergeant Major's Dreadweave Boots (220891, -0.40 DPS) [vendor]; Venomspew Footpads (275606, -1.12 DPS) [crafted] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 22.5 spell_power points (13.69 DPS) | yes | Lorekeeper's Ring (19522, -4.50 DPS) [rep]; Blackstone Ring (17713, -4.79 DPS) [dungeon]; Deep Rooted Ring (19109, -5.21 DPS) [world] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | 20.3 spell_power points (12.33 DPS) | yes | Lorekeeper's Ring (19522, -3.14 DPS) [rep]; Blackstone Ring (17713, -3.42 DPS) [dungeon]; Deep Rooted Ring (19109, -3.85 DPS) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (693.0 DPS) | yes | Ankh of Life (1713, -6.91 DPS) [world_drop]; Shard of the Splithooves (10659, -7.00 DPS) [quest]; Evonice's Landin' Pilla (18951, -7.10 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (693.0 DPS) | yes | Enriched Seal (271905, +0.00 DPS) [rep]; Ankh of Life (1713, -2.54 DPS) [world_drop]; Shard of the Splithooves (10659, -2.63 DPS) [quest] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-verified (693.0 DPS) | yes | Grand Marshal's Stave (234571, -5.23 DPS) [pvp]; Sageclaw (20070, -8.49 DPS) [rep]; Ironbark Staff (20069, -20.41 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (17.64 DPS) | yes | Tome of the Ice Lord (19310, -0.60 DPS) [rep]; Grand Marshal's Tome of Power (234589, -0.89 DPS) [pvp]; Spirit of Aquementas (11904, -4.86 DPS) [quest] |
| ranged | Wand of Biting Cold (19108) | Korrak the Bloodrager [quest] | sim-verified (693.0 DPS) | yes | Brilliant Wand (249385, -2.00 DPS) [crafted]; Greater Eternal Wand (249237, -2.83 DPS) [crafted]; Icefury Wand (7514, -3.28 DPS) [quest] |

**New at 60:** head: Spellpower Goggles Xtreme Plus; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Highlander's Cloth Girdle; legs: Sentinel's Silk Leggings; finger1: Don Julio's Band; finger2: Maiden's Circle; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Wand of Biting Cold

No-known-source sample (15 of 1037, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (troll, 000000000000000000-00000000000000000-253100000000000000)

Set DPS (verified): 45.9. Weights run: 2.7s. Verify run: 2.4s. 172 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.978 ± 0.007, crit=0.093 ± 0.003 per rating point (14 rating = 1%, 1.295 per %), hit=0.263 ± 0.008 per rating point (10 rating = 1%, 2.631 per %), spell_haste=0.623 ± 0.103, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002, spirit=not significant (0.000 ± 0.000), mp5=2.134 ± 0.106

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Silky Circlet (253951) | Tailoring [crafted] | 11.0 spell_power points (2.37 DPS) | yes | Pristine Circlet (253949, -1.05 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.32 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.8 spell_power points (2.98 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.91 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -1.60 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -2.11 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.20 DPS) [crafted]; Sanguine Cape (14376, -0.22 DPS) [world_drop]; Black Whelp Cloak (7283, -0.42 DPS) [crafted] |
| chest | Filigreed Silky Gown (253903) | Tailoring [crafted] | 13.9 spell_power points (3.00 DPS) | yes | Filigreed Pristine Gown (253901, -0.84 DPS, sim-verified) [crafted]; Gray Woolen Robe (2585, -1.08 DPS) [crafted]; Mystic's Wrap (14369, -1.52 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 11.3 spell_power points (2.44 DPS) | yes | Bright Bracers (3647, -0.70 DPS, sim-verified) [world_drop]; Tabitha's Cuffs (251486, -1.79 DPS) [quest]; Owlbeard Bracers (16981, -1.80 DPS) [quest] |
| hands | Silky Gloves (253915) | Tailoring [crafted] | 9.9 spell_power points (2.14 DPS) | yes | Pristine Gloves (253913, -0.65 DPS) [crafted]; Serpent Gloves (5970, -0.67 DPS, sim-verified) [dungeon]; Gnoll Casting Gloves (892, -0.85 DPS) [world] |
| waist | Silky Sash (253927) | Tailoring [crafted] | 10.9 spell_power points (2.35 DPS) | yes | Pristine Sash (253925, -0.63 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.67 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.86 DPS) [crafted] |
| legs | Filigreed Silky Leggings (253939) | Tailoring [crafted] | 16.9 spell_power points (3.64 DPS) | yes | Filigreed Pristine Leggings (253937, -1.05 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -2.13 DPS) [dungeon]; Darkweave Breeches (12987, -2.16 DPS) [world_drop] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.9 spell_power points (2.35 DPS) | yes | Silky Boots (253891, -0.37 DPS, sim-verified) [crafted]; Pristine Boots (253889, -1.07 DPS) [crafted]; Red Woolen Boots (4313, -1.49 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 9.3 spell_power points (2.00 DPS) | yes | Loop of Sacrifice (281673, -0.94 DPS) [quest]; Volcanic Rock Ring (12053, -1.37 DPS) [world_drop]; Furlfeather Ring (286745, -1.78 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.9 spell_power points (1.27 DPS) | yes | Loop of Sacrifice (281673, -0.21 DPS) [quest]; Volcanic Rock Ring (12053, -0.63 DPS) [world_drop]; Furlfeather Ring (286745, -1.05 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-verified (45.9 DPS) | yes | Staff of the Blessed Seer (2271, -0.42 DPS) [dungeon]; Crescent Staff (6505, -0.77 DPS, sim-verified) [quest]; Gnarled Necromancer's Staff (251534, -0.85 DPS) [quest] |
| off_hand | Antipodean Rod (2879) | World drop [world_drop] | 6.0 spell_power points (1.29 DPS) | yes | Seer's Fine Stein (7608, -0.66 DPS) [world_drop]; Spellbinder Orb (15926, -0.66 DPS) [dungeon]; Tork Wrench (11855, -0.87 DPS) [quest] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller (12984, -0.43 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -0.77 DPS) [quest]; Greater Magic Wand (11288, -0.81 DPS) [crafted] |

**New at 20:** head: Silky Circlet; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Silky Gown; wrist: Mindthrust Bracers; hands: Silky Gloves; waist: Silky Sash; legs: Filigreed Silky Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Evocator's Blade; off_hand: Antipodean Rod; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 172, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 23173 Abomination Skin Leggings; 209618 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade

### Band 30 (troll, 000000000000000000-00000000000000000-253510130100000000)

Set DPS (verified): 79.5. Weights run: 2.6s. Verify run: 3.1s. 295 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.301 ± 0.005, crit=0.111 ± 0.004 per rating point (14 rating = 1%, 1.552 per %), hit=0.208 ± 0.006 per rating point (10 rating = 1%, 2.076 per %), spell_haste=1.103 ± 0.061, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002, spirit=0.179 ± 0.038, mp5=0.287 ± 0.042

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Silky Circlet (253977) | Tailoring [crafted] | 14.0 spell_power points (2.89 DPS) | yes | Silky Circlet (253951, -0.62 DPS) [crafted]; Holy Shroud (2721, -0.70 DPS, sim-verified) [world_drop]; Azure Silk Hood (7048, -0.83 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Crystal Starfire Medallion (5003, -0.01 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.06 DPS) [vendor]; Kaleidoscope Chain (13084, -0.34 DPS, sim-verified) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.7 spell_power points (2.42 DPS) | yes | Death Speaker Mantle (6685, -0.49 DPS) [dungeon]; Invoker's Mantle (215365, -0.66 DPS) [crafted]; Magician's Mantle (12998, -0.83 DPS) [world_drop] |
| back | Azure Silk Cloak (7053) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, -1.14 DPS) [crafted]; Heavy Woolen Cloak (4311, -1.20 DPS) [crafted]; Windsong Drape (15468, -1.27 DPS, sim-verified) [quest] |
| chest | Silky Gown (253963) | Tailoring [crafted] | 15.8 spell_power points (3.26 DPS) | yes | Beguiler Robes (7728, -0.37 DPS) [dungeon]; Tree Bark Jacket (1486, -0.58 DPS) [dungeon]; Green Silk Armor (7065, -0.60 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.9 spell_power points (2.03 DPS) | yes | Glowing Magical Bracelets (13106, -1.54 DPS) [world_drop]; Mindthrust Bracers (1974, -1.55 DPS) [dungeon]; Tabitha's Cuffs (251486, -2.35 DPS, sim-verified) [quest] |
| hands | Azure Silk Gloves (4319) | Tailoring [crafted] | 10.5 spell_power points (2.17 DPS) | yes | Silky Gloves (253915, -0.54 DPS) [crafted]; Jutebraid Gloves (10654, -0.63 DPS) [quest]; Serpent Gloves (5970, -0.73 DPS) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.9 spell_power points (2.46 DPS) | yes | Belt of Arugal (6392, -0.23 DPS) [dungeon]; Invoker's Cord (215366, -0.70 DPS) [crafted]; Warsong Sash (16975, -0.83 DPS, sim-verified) [quest] |
| legs | Silky Leggings (253989) | Tailoring [crafted] | 17.0 spell_power points (3.51 DPS) | yes | Filigreed Silky Leggings (253939, -0.72 DPS) [crafted]; Azure Silk Pants (7046, -1.07 DPS) [crafted]; Gaze Dreamer Pants (6903, -2.09 DPS, sim-verified) [dungeon] |
| feet | Frothing Slippers (254003) | Tailoring [crafted] | 15.8 spell_power points (3.26 DPS) | yes | Gilded Slippers (254001, -1.53 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -1.57 DPS) [crafted]; Acidic Walkers (9454, -1.59 DPS) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.6 spell_power points (1.56 DPS) | yes | Ring of Power Regulation (273030, -0.15 DPS) [dungeon]; Solid Gold Band (279845, -0.55 DPS) [quest]; Electrocutioner Lagnut (9447, -0.69 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 7.5 spell_power points (1.55 DPS) | yes | Ring of Power Regulation (273030, -0.13 DPS) [dungeon]; Solid Gold Band (279845, -0.54 DPS) [quest]; Electrocutioner Lagnut (9447, -0.67 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -0.97 DPS) [dungeon]; Tendonscraper Dagger (273053, -0.97 DPS) [world]; Glimmering Staff (249392, -2.41 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.8 spell_power points (1.82 DPS) | yes | Orb of Souls (249395, -0.77 DPS) [crafted]; Witch's Finger (16887, -1.32 DPS) [quest]; Antipodean Rod (2879, -1.78 DPS, sim-verified) [world_drop] |
| ranged | Starfaller (13063) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderwood (13062, +0.00 DPS) [world_drop]; Greater Mystic Wand (217287, +0.00 DPS) [crafted]; Wand of Eventide (5214, -0.04 DPS) [world_drop] |

**New at 30:** head: Filigreed Silky Circlet; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Azure Silk Cloak; chest: Silky Gown; wrist: Spidertank Oilrag; hands: Azure Silk Gloves; waist: Defiler's Cloth Girdle; legs: Silky Leggings; feet: Frothing Slippers; finger1: Advisor's Ring; finger2: Snake Hoop; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Starfaller

No-known-source sample (15 of 295, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia; 18859 Insignia of the Alliance; 19545 Scout's Blade

### Band 40 (troll, 000000000000000000-00000000000000000-253510130101301221)

Set DPS (verified): 190.2. Weights run: 2.7s. Verify run: 3.5s. 379 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.544 ± 0.012, crit=0.134 ± 0.005 per rating point (14 rating = 1%, 1.872 per %), hit=0.285 ± 0.014 per rating point (10 rating = 1%, 2.852 per %), spell_haste=1.264 ± 0.166, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004, spirit=not significant (0.325 ± 0.114), mp5=0.625 ± 0.153

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.57 DPS) | yes | Papal Fez (9431, -0.26 DPS) [dungeon]; Augural Shroud (2620, -1.43 DPS) [world]; Whitemane's Chapeau (7720, -1.46 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.3 spell_power points (3.55 DPS) | yes | Triune Amulet (7722, -1.59 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -1.65 DPS) [dungeon]; Darkspear Warding Pendant (272074, -2.36 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.1 spell_power points (4.40 DPS) | yes | Green Silken Shoulders (7057, -0.03 DPS) [crafted]; Bloodmage Mantle (7684, -0.05 DPS) [dungeon]; Azure Shoulders (7060, -0.33 DPS) [crafted] |
| back | Frost Leather Cloak (7377) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Azure Silk Cloak (7053, -0.01 DPS) [crafted]; Icy Cloak (4327, -0.64 DPS) [crafted]; Long Silken Cloak (4326, -0.71 DPS) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 27.4 spell_power points (8.58 DPS) | yes | Robe of Power (7054, -1.35 DPS) [crafted]; Elemental Raiment (9434, -2.02 DPS) [world_drop]; Robe of the Magi (1716, -2.21 DPS, sim-verified) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 10.9 spell_power points (3.40 DPS) | yes | Radiant Silver Bracers (4545, -0.79 DPS) [quest]; Condor Bracers (15864, -1.21 DPS) [quest]; Earthen Silk Cuffs (254019, -1.64 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.5 spell_power points (7.02 DPS) | yes | Red Mageweave Gloves (10018, -1.88 DPS) [crafted]; Black Mageweave Gloves (10003, -2.33 DPS) [crafted]; Frothing Handwraps (254023, -3.00 DPS, sim-verified) [crafted] |
| waist | Frothing Cord (254039) | Tailoring [crafted] | 20.3 spell_power points (6.35 DPS) | yes | Deathmage Sash (10771, -1.61 DPS) [dungeon]; Star Belt (4329, -1.88 DPS) [crafted]; Defiler's Cloth Girdle (20166, -3.07 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 20.5 spell_power points (6.42 DPS) | yes | Silky Leggings (253989, -0.34 DPS) [crafted]; Filigreed Silky Leggings (253939, -1.55 DPS) [crafted]; Stoneweaver Leggings (9407, -1.65 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 26.3 spell_power points (8.22 DPS) | yes | Frothing Slippers (254003, -3.96 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -4.43 DPS) [crafted]; Acidic Walkers (9454, -4.89 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.7 spell_power points (3.98 DPS) | yes | Snake Hoop (6750, -1.25 DPS) [quest]; Ring of Power Regulation (273030, -1.51 DPS) [dungeon]; Reedknot Ring (9622, -1.79 DPS) [quest] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 10.9 spell_power points (3.40 DPS) | yes | Snake Hoop (6750, -0.68 DPS) [quest]; Ring of Power Regulation (273030, -0.94 DPS) [dungeon]; Reedknot Ring (9622, -1.21 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (190.2 DPS) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (190.2 DPS) | yes | Celestial Stave (9517, -8.78 DPS) [quest]; Hypnotic Blade (7714, -9.10 DPS) [dungeon]; Illusionary Rod (7713, -14.72 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Icefury Wand (7514) | Mage's Wand [quest] | sim-verified (190.2 DPS) | yes | Twisted Nether Wand (249144, -1.62 DPS) [crafted]; Jaina's Firestarter (13064, -1.89 DPS) [world_drop]; Dancing Flame (6806, -2.30 DPS, sim-verified) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Frost Leather Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Frothing Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Staff of Jordan; ranged: Icefury Wand

No-known-source sample (15 of 379, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 50 (troll, 230005000000000000-00000000000000000-253510130101301221)

Set DPS (verified): 266.4. Weights run: 2.8s. Verify run: 6.3s. 483 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.548 ± 0.023, crit=0.244 ± 0.009 per rating point (14 rating = 1%, 3.420 per %), hit=0.566 ± 0.030 per rating point (10 rating = 1%, 5.661 per %), spell_haste=2.104 ± 0.354, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004, spirit=not significant (0.854 ± 0.312), mp5=not significant (0.899 ± 0.335)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreamweave Circlet (10041) | Tailoring [crafted] | 36.7 spell_power points (10.67 DPS) | yes | Red Mageweave Headband (10033, -1.97 DPS) [crafted]; Gemburst Circlet (10751, -2.33 DPS) [quest]; Spellpower Goggles Xtreme Plus (15999, -2.83 DPS) [crafted] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 14.0 spell_power points (4.06 DPS) | yes | Mindburst Medallion (11196, -0.12 DPS) [quest]; Glowing Eye of Mordresh (10769, -0.75 DPS) [dungeon]; Horizon Choker (13085, -0.84 DPS) [world_drop] |
| shoulder | Netherfroth Shoulders (254051) | Tailoring [crafted] | 26.5 spell_power points (7.70 DPS) | yes | Rotgrip Mantle (17732, -1.06 DPS) [dungeon]; Nethergeld Shoulders (254049, -2.32 DPS) [crafted]; Mantle of Doan (7712, -2.73 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -1.59 DPS) [quest]; Runecloth Cloak (13860, -2.62 DPS) [crafted]; Azure Silk Cloak (7053, -2.86 DPS) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 34.9 spell_power points (10.13 DPS) | yes | Frostweave Robe (13868, -0.44 DPS) [crafted]; Ghostweave Vest (14141, -0.47 DPS) [crafted]; Frostweave Tunic (13869, -1.41 DPS) [crafted] |
| wrist | Netherfroth Cuffs (254063) | Tailoring [crafted] | 21.1 spell_power points (6.13 DPS) | yes | Bloodband Bracers (11469, -2.50 DPS) [quest]; Nethergeld Cuffs (254061, -2.66 DPS, sim-verified) [crafted]; Spidertank Oilrag (9448, -2.73 DPS) [dungeon] |
| hands | Frothing Gloves (254097) | Tailoring [crafted] | 31.3 spell_power points (9.09 DPS) | yes | Runecloth Gloves (13863, -1.94 DPS) [crafted]; Frostweave Gloves (13870, -2.04 DPS) [crafted]; Dreamweave Gloves (10019, -3.66 DPS, sim-verified) [crafted] |
| waist | Frothing Waistcord (254083) | Tailoring [crafted] | 23.9 spell_power points (6.95 DPS) | yes | Frothing Cord (254039, -0.12 DPS) [crafted]; Ghostweave Cord (254073, -0.16 DPS) [crafted]; Satyrmane Sash (17755, -1.29 DPS) [dungeon] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 31.2 spell_power points (9.06 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -0.86 DPS) [dungeon]; Dalewind Trousers (13008, -1.43 DPS) [world_drop]; Silky Leggings (253989, -2.64 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 30.0 spell_power points (8.70 DPS) | yes | Frothing Sandals (254109, -1.46 DPS) [crafted]; Frothing Slippers (254003, -2.82 DPS) [crafted]; Vinerot Sandals (17748, -3.05 DPS) [dungeon] |
| finger1 | Advisor's Ring (19519) | Warsong Outriders [rep] | 15.6 spell_power points (4.53 DPS) | yes | Philanthropist's Ring (281635, -0.83 DPS) [quest]; Coldwater Ring (4550, -1.82 DPS) [quest]; Eye of Adaegus (5266, -7.37 DPS, sim-verified) [world_drop] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Philanthropist's Ring (281635, -0.07 DPS) [quest]; Coldwater Ring (4550, -1.06 DPS) [quest]; Eye of Adaegus (5266, -1.28 DPS, sim-verified) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+3.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spire of Hakkar (10844, -3.58 DPS) [world]; Glowing Brightwood Staff (812, -4.18 DPS) [world_drop]; Kindling Stave (11750, -10.47 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Icefury Wand (7514) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lesser Eternal Wand (249232, -0.14 DPS) [crafted]; Noxious Shooter (17745, -0.69 DPS) [dungeon]; Dreambough Wand (249234, -0.91 DPS) [crafted] |

**New at 50:** head: Dreamweave Circlet; neck: Triune Amulet; shoulder: Netherfroth Shoulders; back: Spritecaster Cape; wrist: Netherfroth Cuffs; hands: Frothing Gloves; waist: Frothing Waistcord; legs: Senior Designer's Pantaloons; finger1: Advisor's Ring; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Darkspear Voodoo Seal

No-known-source sample (15 of 483, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60 (troll, 230005200000000000-51020000000000000-253510130101301221)

Set DPS (verified): 414.7. Weights run: 3.2s. Verify run: 17.7s. 1032 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.357 ± 0.024, crit=0.288 ± 0.009 per rating point (14 rating = 1%, 4.025 per %), hit=0.769 ± 0.041 per rating point (10 rating = 1%, 7.689 per %), spell_haste=2.570 ± 0.356, spell_penetration=not significant (0.000 ± 0.000), frost_power=0.843 ± 0.003, spirit=not significant (0.494 ± 0.270), mp5=not significant (0.824 ± 0.320)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreadmist Mask (16698) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (414.7 DPS) | yes | Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Dreamweave Circlet (10041, -5.92 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 15.6 spell_power points (5.89 DPS) | yes | Triune Amulet (7722, -1.56 DPS) [dungeon]; Mindburst Medallion (11196, -1.88 DPS) [quest]; Glowing Eye of Mordresh (10769, -2.16 DPS) [dungeon] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 28.5 spell_power points (10.77 DPS) | yes | Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -1.06 DPS) [crafted]; Champion's Silk Mantle (227104, -1.34 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (414.7 DPS) | yes | Cloak of Flames (3475, +0.00 DPS) [world_drop]; Deep Woodlands Cloak (19121, -1.47 DPS) [quest]; Frostwolf Advisor's Cloak (19085, -1.93 DPS) [rep] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-verified (414.7 DPS) | yes | Warlord's Silk Raiment (231596, -0.85 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -5.58 DPS) [pvp]; Robe of Winter Night (14136, -8.81 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 28.3 spell_power points (10.71 DPS) | yes | Netherfroth Cuffs (254063, -4.68 DPS) [crafted]; Runecloth Cuffs (254123, -5.20 DPS) [crafted]; Earthenweave Cuffs (254131, -5.82 DPS) [crafted] |
| hands | Dreadmist Wraps (16705) | Scholomance: Lorekeeper Polkelt [dungeon] | sim-verified (414.7 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Frothing Gloves (254097, -8.72 DPS, sim-verified) [crafted] |
| waist | Dreadmist Belt (16702) | Scholomance: Scholomance Necromancer [dungeon] | sim-verified (414.7 DPS) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Defiler's Cloth Girdle (20163, +0.00 DPS) [rep]; Frostwolf Cloth Belt (19090, -4.72 DPS, sim-verified) [rep] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 39.7 spell_power points (15.03 DPS) | yes | General's Silk Trousers (231595, +0.00 DPS) [pvp]; Ghostweave Pants (14144, -1.45 DPS) [crafted]; Legionnaire's Silk Legguards (227107, -2.33 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 27.5 spell_power points (10.39 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.57 DPS) [crafted]; Rime-encrusted Boots (279268, -0.64 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 23.5 spell_power points (8.88 DPS) | yes | Deep Rooted Ring (19109, -2.99 DPS) [world]; Frigid Ring (18679, -3.06 DPS) [world]; Band of the Unicorn (7553, -3.96 DPS) [world_drop] |
| finger2 | Advisor's Ring (19518) | Warsong Outriders [rep] | 17.3 spell_power points (6.54 DPS) | yes | Deep Rooted Ring (19109, -0.65 DPS) [world]; Frigid Ring (18679, -0.73 DPS) [world]; Band of the Unicorn (7553, -1.62 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (414.7 DPS) | yes | Ankh of Life (1713, -0.38 DPS) [world_drop]; Shard of the Splithooves (10659, -0.75 DPS) [quest]; Uther's Strength (11302, -4.11 DPS, sim-verified) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (414.7 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Shard of the Splithooves (10659, -5.70 DPS, sim-verified) [quest] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-verified (414.7 DPS) | yes | High Warlord's War Staff (234549, -2.49 DPS) [pvp]; Elemental Mage Staff (944, -3.87 DPS) [world_drop]; Ironbark Staff (20220, -17.46 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (10.97 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; Tome of the Ice Lord (19310, -0.00 DPS) [rep]; Spirit of Aquementas (11904, -3.03 DPS) [quest] |
| ranged | Wand of Biting Cold (19108) | The Legend of Korrak [quest] | sim-verified (414.7 DPS) | yes | Torch of Light (279246, -0.57 DPS) [crafted]; Brilliant Wand (249385, -0.83 DPS) [crafted]; Greater Eternal Wand (249237, -1.88 DPS) [crafted] |

**New at 60:** head: Dreadmist Mask; neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Dreadmist Wraps; waist: Dreadmist Belt; legs: Outrider's Silk Leggings; finger1: Maiden's Circle; finger2: Advisor's Ring; trinket2: Rune of the Guard Captain; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Wand of Biting Cold

No-known-source sample (15 of 1032, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60, raid preset (troll, 230005200000000000-51020000000000000-253510130101301221)

Set DPS (verified): 681.7. Weights run: 3.4s. Verify run: 16.6s. 1032 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.195 ± 0.012, crit=0.562 ± 0.017 per rating point (14 rating = 1%, 7.869 per %), hit=1.465 ± 0.083 per rating point (10 rating = 1%, 14.651 per %), spell_haste=7.953 ± 0.439, spell_penetration=not significant (0.000 ± 0.000), frost_power=0.853 ± 0.002, spirit=not significant (0.152 ± 0.054), mp5=not significant (0.280 ± 0.089)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (16.42 DPS) | yes | Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Dreamweave Circlet (10041, -1.35 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 14.0 spell_power points (8.48 DPS) | yes | Triune Amulet (7722, -2.97 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -3.28 DPS) [dungeon]; Mindburst Medallion (11196, -3.66 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 27.4 spell_power points (16.67 DPS) | yes | Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Argent Shoulders (19059, -0.82 DPS) [crafted]; Champion's Silk Mantle (227104, -1.08 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of Flames (3475, -0.87 DPS) [world_drop]; Frostwolf Advisor's Cloak (19085, -1.27 DPS) [rep]; Deep Woodlands Cloak (19121, -1.41 DPS) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-verified (681.7 DPS) | yes | Warlord's Silk Raiment (231596, -3.11 DPS) [pvp]; Robe of Winter Night (14136, -8.62 DPS, sim-verified) [crafted]; Legionnaire's Silk Tunic (227106, -10.50 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 24.6 spell_power points (14.97 DPS) | yes | Netherfroth Cuffs (254063, -6.93 DPS) [crafted]; Runecloth Cuffs (254123, -7.10 DPS) [crafted]; Spidertank Oilrag (9448, -8.99 DPS) [dungeon] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 26.7 spell_power points (16.23 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; Earth Warder's Gloves (21318, +0.00 DPS) [quest]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Defiler's Cloth Girdle (20163) | The Defilers [rep] | 23.0 spell_power points (14.01 DPS) | yes | Frostwolf Cloth Belt (19090, -1.88 DPS) [rep]; Frothing Waistcord (254083, -3.08 DPS) [crafted]; Satyrmane Sash (17755, -4.31 DPS) [dungeon] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 33.2 spell_power points (20.21 DPS) | yes | General's Silk Trousers (231595, +0.00 DPS) [pvp]; Legionnaire's Silk Legguards (227107, -0.17 DPS) [pvp]; Ghostweave Pants (14144, -2.55 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.1 spell_power points (15.24 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; First Sergeant's Dreadweave Boots (220909, -0.40 DPS) [vendor]; Venomspew Footpads (275606, -1.12 DPS) [crafted] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 22.5 spell_power points (13.69 DPS) | yes | Advisor's Ring (19518, -4.50 DPS) [rep]; Blackstone Ring (17713, -4.79 DPS) [dungeon]; Deep Rooted Ring (19109, -5.21 DPS) [world] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | 20.3 spell_power points (12.33 DPS) | yes | Advisor's Ring (19518, -3.14 DPS) [rep]; Blackstone Ring (17713, -3.42 DPS) [dungeon]; Deep Rooted Ring (19109, -3.85 DPS) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (681.7 DPS) | yes | Uther's Strength (11302, -5.14 DPS, sim-verified) [world_drop]; Ankh of Life (1713, -6.91 DPS) [world_drop]; Shard of the Splithooves (10659, -7.00 DPS) [quest] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (681.7 DPS) | yes | Uther's Strength (11302, -2.59 DPS) [world_drop]; Ankh of Life (1713, -5.13 DPS) [world_drop]; Shard of the Splithooves (10659, -5.21 DPS) [quest] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-verified (681.7 DPS) | yes | High Warlord's War Staff (234549, -5.23 DPS) [pvp]; Mindfang (20214, -8.49 DPS) [rep]; Ironbark Staff (20220, -20.00 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (17.64 DPS) | yes | Tome of the Ice Lord (19310, -0.60 DPS) [rep]; High Warlord's Tome of Destruction (234563, -0.89 DPS) [pvp]; Spirit of Aquementas (11904, -4.86 DPS) [quest] |
| ranged | Wand of Biting Cold (19108) | The Legend of Korrak [quest] | sim-verified (681.7 DPS) | yes | Brilliant Wand (249385, -2.00 DPS) [crafted]; Greater Eternal Wand (249237, -2.83 DPS) [crafted]; Icefury Wand (7514, -3.28 DPS) [quest] |

**New at 60:** head: Spellpower Goggles Xtreme Plus; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Defiler's Cloth Girdle; legs: Outrider's Silk Leggings; finger1: Don Julio's Band; finger2: Maiden's Circle; trinket2: Rune of the Guard Captain; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Wand of Biting Cold

No-known-source sample (15 of 1032, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

