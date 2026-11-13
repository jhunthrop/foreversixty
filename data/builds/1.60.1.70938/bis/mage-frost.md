# Leveling BiS: Frost

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-053200010000000000)

Set DPS (verified): 51.2. Weights run: 2.2s. Verify run: 2.0s. 198 eligible items had no known source.

Race sweep: gnome 51.2 ± 0.1, high-order-skyborne 49.8 ± 0.1, human 49.4 ± 0.1 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.982 ± 0.007, crit=0.104 ± 0.003 per rating point (14 rating = 1%, 1.458 per %), hit=0.256 ± 0.008 per rating point (10 rating = 1%, 2.564 per %), spell_haste=0.783 ± 0.086, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002, spirit=not significant (0.023 ± 0.007), mp5=1.328 ± 0.104

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Silky Circlet (253951) | Tailoring [crafted] | 11.0 spell_power points (2.32 DPS) | yes | Pristine Circlet (253949, -1.14 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.26 DPS) [crafted]; Flying Tiger Goggles (4368, -2.30 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.0 spell_power points (0.01 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.8 spell_power points (2.92 DPS) | yes | Reinforced Woolen Shoulders (4315, -1.25 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -2.08 DPS) [crafted]; Slime-encrusted Pads (6461, -2.08 DPS) [dungeon] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.18 DPS) [crafted]; Sanguine Cape (14376, -0.22 DPS) [world_drop]; Caretaker's Cape (20428, -0.40 DPS) [rep] |
| chest | Filigreed Silky Gown (253903) | Tailoring [crafted] | 14.0 spell_power points (2.95 DPS) | yes | Filigreed Pristine Gown (253901, -0.91 DPS, sim-verified) [crafted]; Gray Woolen Robe (2585, -1.07 DPS) [crafted]; Mystic's Wrap (14369, -1.50 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 8.9 spell_power points (1.88 DPS) | yes | Bright Bracers (3647, -0.84 DPS, sim-verified) [world_drop]; Repurposed Hair Band (281256, -1.45 DPS) [quest]; Mystic's Bracelets (14366, -1.46 DPS) [world_drop] |
| hands | Silky Gloves (253915) | Tailoring [crafted] | 9.9 spell_power points (2.10 DPS) | yes | Pristine Gloves (253913, -0.63 DPS) [crafted]; Serpent Gloves (5970, -0.74 DPS, sim-verified) [dungeon]; Gnoll Casting Gloves (892, -0.83 DPS) [world] |
| waist | Silky Sash (253927) | Tailoring [crafted] | 10.9 spell_power points (2.31 DPS) | yes | Keller's Girdle (2911, -0.65 DPS) [world_drop]; Pristine Sash (253925, -0.68 DPS, sim-verified) [crafted]; Novice Arcanist's Sash (253885, -0.84 DPS) [crafted] |
| legs | Filigreed Silky Leggings (253939) | Tailoring [crafted] | 17.0 spell_power points (3.59 DPS) | yes | Filigreed Pristine Leggings (253937, -1.14 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -2.11 DPS) [world_drop]; Silk-threaded Trousers (1929, -2.11 DPS) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.9 spell_power points (2.31 DPS) | yes | Silky Boots (253891, -0.75 DPS, sim-verified) [crafted]; Pristine Boots (253889, -1.05 DPS) [crafted]; Red Woolen Boots (4313, -1.44 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 7.7 spell_power points (1.62 DPS) | yes | Lavishly Jeweled Ring (1156, -0.37 DPS) [dungeon]; Volcanic Rock Ring (12053, -1.00 DPS) [world_drop]; Furlfeather Ring (286745, -1.41 DPS) [world] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 7.0 spell_power points (1.47 DPS) | yes | Lavishly Jeweled Ring (1156, -0.23 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.85 DPS) [world_drop]; Furlfeather Ring (286745, -1.26 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | 30.1 spell_power points (6.35 DPS) | yes | Emberstone Staff (5201, -1.28 DPS) [dungeon]; Buzzer Blade (2169, -1.71 DPS) [dungeon]; Staff of the Blessed Seer (2271, -1.85 DPS, sim-verified) [dungeon] |
| off_hand | Antipodean Rod (2879) | World drop [world_drop] | 6.0 spell_power points (1.27 DPS) | yes | Mathystra Channeling Orb (281291, -0.23 DPS) [quest]; Tear of Grief (5611, -0.25 DPS) [quest]; Dwarven Tome (279898, -0.63 DPS) [quest] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller (12984, -0.17 DPS) [world_drop]; Sizzle Stick (8071, -0.66 DPS) [quest]; Greater Magic Wand (11288, -0.69 DPS) [crafted] |

**New at 20:** head: Silky Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Silky Gown; wrist: Mindthrust Bracers; hands: Silky Gloves; waist: Silky Sash; legs: Filigreed Silky Leggings; feet: Spidersilk Boots; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Evocator's Blade; off_hand: Antipodean Rod; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 198, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 8350 The 1 Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (gnome, 000000000000000000-00000000000000000-053510030100200100)

Set DPS (verified): 88.9. Weights run: 2.3s. Verify run: 2.3s. 326 eligible items had no known source.

Race sweep: gnome 88.9 ± 0.1, high-order-skyborne 88.0 ± 0.1, human 87.3 ± 0.1 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.270 ± 0.008, crit=0.108 ± 0.004 per rating point (14 rating = 1%, 1.515 per %), hit=0.228 ± 0.009 per rating point (10 rating = 1%, 2.279 per %), spell_haste=1.125 ± 0.086, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003, spirit=not significant (-0.045 ± 0.035), mp5=not significant (0.191 ± 0.064)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Silky Circlet (253977) | Tailoring [crafted] | 14.0 spell_power points (3.02 DPS) | yes | Holy Shroud (2721, -0.65 DPS) [world_drop]; Silky Circlet (253951, -0.84 DPS, sim-verified) [crafted]; Azure Silk Hood (7048, -0.86 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 1.6 spell_power points (0.35 DPS) | yes | Seal of the Expedition (271920, -0.06 DPS) [rep]; Crystal Starfire Medallion (5003, -0.12 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.12 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.4 spell_power points (2.46 DPS) | yes | Death Speaker Mantle (6685, -0.65 DPS, sim-verified) [dungeon]; Invoker's Mantle (215365, -0.66 DPS) [crafted]; Magician's Mantle (12998, -0.86 DPS) [world_drop] |
| back | Azure Silk Cloak (7053) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, -1.08 DPS) [crafted]; Heavy Woolen Cloak (4311, -1.29 DPS) [crafted]; Vine Pruner's Cloak (279835, -1.79 DPS, sim-verified) [quest] |
| chest | Silky Gown (253963) | Tailoring [crafted] | 14.9 spell_power points (3.21 DPS) | yes | Green Silk Armor (7065, -0.51 DPS) [crafted]; Beguiler Robes (7728, -0.57 DPS) [dungeon]; Tree Bark Jacket (1486, -1.50 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.6 spell_power points (2.06 DPS) | yes | Mindthrust Bracers (1974, -1.65 DPS) [dungeon]; Nightsky Wristbands (6407, -1.72 DPS) [world_drop]; Glowing Magical Bracelets (13106, -2.45 DPS, sim-verified) [world_drop] |
| hands | Azure Silk Gloves (4319) | Tailoring [crafted] | 10.0 spell_power points (2.16 DPS) | yes | Serpent Gloves (5970, -0.65 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.86 DPS) [world]; Silky Gloves (253915, -0.90 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.8 spell_power points (2.55 DPS) | yes | Belt of Arugal (6392, -0.43 DPS) [dungeon]; Invoker's Cord (215366, -0.75 DPS) [crafted]; Silky Sash (253927, -0.80 DPS) [crafted] |
| legs | Silky Leggings (253989) | Tailoring [crafted] | 15.9 spell_power points (3.43 DPS) | yes | Gaze Dreamer Pants (6903, -0.84 DPS) [dungeon]; Azure Silk Pants (7046, -0.92 DPS) [crafted]; Filigreed Silky Leggings (253939, -1.09 DPS, sim-verified) [crafted] |
| feet | Frothing Slippers (254003) | Tailoring [crafted] | 14.9 spell_power points (3.21 DPS) | yes | Spidersilk Boots (4320, -1.47 DPS) [crafted]; Acidic Walkers (9454, -1.67 DPS) [dungeon]; Gilded Slippers (254001, -1.69 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.4 spell_power points (1.59 DPS) | yes | Ring of Power Regulation (273030, -0.17 DPS) [dungeon]; Minor Channeling Ring (1449, -0.40 DPS) [quest]; Runemender's Seal (273031, -0.60 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 7.3 spell_power points (1.58 DPS) | yes | Ring of Power Regulation (273030, -0.17 DPS) [dungeon]; Minor Channeling Ring (1449, -0.39 DPS) [quest]; Runemender's Seal (273031, -0.59 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -0.86 DPS) [dungeon]; Tendonscraper Dagger (273053, -0.86 DPS) [world]; Glimmering Staff (249392, -3.61 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.6 spell_power points (1.86 DPS) | yes | Eye of Paleth (2943, -1.00 DPS) [quest]; Orb of Souls (249395, -1.00 DPS) [crafted]; Antipodean Rod (2879, -1.34 DPS, sim-verified) [world_drop] |
| ranged | Thunderwood (13062) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Greater Mystic Wand (217287, -0.04 DPS) [crafted]; Wand of Eventide (5214, -0.34 DPS) [world_drop]; Starfaller (13063, -0.54 DPS) [world_drop] |

**New at 30:** head: Filigreed Silky Circlet; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Azure Silk Cloak; chest: Silky Gown; wrist: Spidertank Oilrag; hands: Azure Silk Gloves; waist: Highlander's Cloth Girdle; legs: Silky Leggings; feet: Frothing Slippers; finger1: Lorekeeper's Ring; finger2: Snake Hoop; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Thunderwood

No-known-source sample (15 of 326, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 000000000000000000-00000000000000000-253510130101301221)

Set DPS (verified): 175.1. Weights run: 2.4s. Verify run: 4.5s. 456 eligible items had no known source.

Race sweep: gnome 175.1 ± 0.2, human 172.6 ± 0.3, high-order-skyborne 171.7 ± 0.3 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.466 ± 0.012, crit=0.273 ± 0.009 per rating point (14 rating = 1%, 3.829 per %), hit=0.553 ± 0.022 per rating point (10 rating = 1%, 5.534 per %), spell_haste=2.375 ± 0.155, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002, spirit=0.644 ± 0.108, mp5=1.061 ± 0.119

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Miner's Hat of the Deep (9429, -0.40 DPS) [dungeon]; Whitemane's Chapeau (7720, -0.59 DPS) [dungeon]; Papal Fez (9431, -1.79 DPS, sim-verified) [dungeon] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 12.5 spell_power points (4.57 DPS) | yes | Glowing Eye of Mordresh (10769, -0.65 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.28 DPS) [quest]; Necklace of Calisea (1714, -1.73 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Green Silken Shoulders (7057, -0.03 DPS) [crafted]; Inquisitor's Shawl (19507, -0.05 DPS) [dungeon]; Mantle of Doan (7712, -0.55 DPS, sim-verified) [dungeon] |
| back | Azure Silk Cloak (7053) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frost Leather Cloak (7377, -0.34 DPS) [crafted]; Icy Cloak (4327, -0.39 DPS) [crafted]; Caretaker's Cape (19532, -0.99 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Dreamweave Vest (10021, -0.42 DPS, sim-verified) [crafted]; Robe of Power (7054, -1.20 DPS) [crafted]; Elemental Raiment (9434, -2.56 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 12.2 spell_power points (4.45 DPS) | yes | Arcane Runed Bracers (4744, -1.16 DPS) [quest]; Earthen Silk Cuffs (254019, -1.81 DPS) [crafted]; Condor Bracers (15864, -1.89 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 24.4 spell_power points (8.91 DPS) | yes | Frothing Handwraps (254023, -1.65 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -3.19 DPS) [crafted]; Black Mageweave Gloves (10003, -3.43 DPS) [crafted] |
| waist | Frothing Cord (254039) | Tailoring [crafted] | 21.6 spell_power points (7.89 DPS) | yes | Highlander's Cloth Girdle (20098, -2.09 DPS) [rep]; Gilded Cord (254037, -2.19 DPS) [crafted]; Star Belt (4329, -2.20 DPS) [crafted] |
| legs | Silky Leggings (253989) | Tailoring [crafted] | 20.5 spell_power points (7.49 DPS) | yes | Red Mageweave Pants (10009, -0.33 DPS) [crafted]; Stoneweaver Leggings (9407, -0.74 DPS) [dungeon]; Gaze Dreamer Pants (6903, -0.98 DPS) [dungeon] |
| feet | Frothing Slippers (254003) | Tailoring [crafted] | 18.8 spell_power points (6.89 DPS) | yes | Thoughtcast Boots (10578, -2.00 DPS) [dungeon]; Gilded Slippers (254001, -2.19 DPS) [crafted]; Furen's Boots (13100, -4.34 DPS, sim-verified) [world_drop] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.3 spell_power points (4.51 DPS) | yes | Runemender's Seal (273031, -0.98 DPS) [world]; Ring of Power Regulation (273030, -1.15 DPS) [dungeon]; Snake Hoop (6750, -1.46 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 12.2 spell_power points (4.45 DPS) | yes | Ring of Power Regulation (273030, -1.10 DPS) [dungeon]; Snake Hoop (6750, -1.41 DPS) [quest]; Runemender's Seal (273031, -1.46 DPS, sim-verified) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -10.27 DPS, sim-verified) [dungeon]; Celestial Stave (9517, -12.55 DPS) [quest]; Hypnotic Blade (7714, -12.57 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Icefury Wand (7514) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Freezing Shard (10572, -0.18 DPS) [dungeon]; Plaguerot Sprig (10766, -1.56 DPS) [dungeon]; Twisted Nether Wand (249144, -1.68 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Triune Amulet; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Frothing Cord; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Staff of Jordan; ranged: Icefury Wand

No-known-source sample (15 of 456, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (gnome, 230005000000000000-00000000000000000-253510130101301221)

Set DPS (verified): 265.8. Weights run: 2.4s. Verify run: 5.2s. 600 eligible items had no known source.

Race sweep: human 266.0 ± 0.4, gnome 265.8 ± 0.4, high-order-skyborne 263.1 ± 0.4 (dps); human beats the guide race gnome by 0.1 (0.04%) ± 0.5, inside error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.589 ± 0.017, crit=0.400 ± 0.013 per rating point (14 rating = 1%, 5.607 per %), hit=0.859 ± 0.036 per rating point (10 rating = 1%, 8.593 per %), spell_haste=3.497 ± 0.267, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002, spirit=0.737 ± 0.175, mp5=1.148 ± 0.192

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreamweave Circlet (10041) | Tailoring [crafted] | 35.7 spell_power points (12.98 DPS) | yes | Red Mageweave Headband (10033, -1.80 DPS) [crafted]; Gemburst Circlet (10751, -3.07 DPS) [quest]; Spellpower Goggles Xtreme Plus (15999, -3.17 DPS) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 19.5 spell_power points (7.09 DPS) | yes | Triune Amulet (7722, -2.31 DPS) [dungeon]; Talvash's Enhancing Necklace (7673, -2.52 DPS) [quest]; Mindburst Medallion (11196, -3.23 DPS, sim-verified) [quest] |
| shoulder | Netherfroth Shoulders (254051) | Tailoring [crafted] | 26.1 spell_power points (9.49 DPS) | yes | Kentic Amice (11624, -0.02 DPS) [dungeon]; Rotgrip Mantle (17732, -0.92 DPS) [dungeon]; Mantle of Lost Hope (22234, -2.20 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Runecloth Cloak (13860, -2.99 DPS) [crafted]; Azure Silk Cloak (7053, -3.54 DPS) [crafted]; Featherskin Cape (10843, -4.03 DPS, sim-verified) [world] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 39.0 spell_power points (14.17 DPS) | yes | Dreamweave Vest (10021, -1.96 DPS) [crafted]; Ghostweave Vest (14141, -2.21 DPS) [crafted]; Frostweave Robe (13868, -2.51 DPS) [crafted] |
| wrist | Netherfroth Cuffs (254063) | Tailoring [crafted] | 20.8 spell_power points (7.56 DPS) | yes | Nethergeld Cuffs (254061, -2.32 DPS, sim-verified) [crafted]; Aristocratic Cuffs (12546, -2.90 DPS) [dungeon]; Bloodband Bracers (11469, -3.01 DPS) [quest] |
| hands | Frothing Gloves (254097) | Tailoring [crafted] | 30.8 spell_power points (11.18 DPS) | yes | Dreamweave Gloves (10019, -1.91 DPS) [crafted]; Frostweave Gloves (13870, -2.27 DPS) [crafted]; Runecloth Gloves (13863, -2.49 DPS) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 26.5 spell_power points (9.62 DPS) | yes | Frothing Waistcord (254083, -0.79 DPS) [crafted]; Frothing Cord (254039, -1.22 DPS) [crafted]; Ghostweave Cord (254073, -1.59 DPS) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Kilt of the Atal'ai Prophet (10807, -0.73 DPS) [dungeon]; Rainstrider Leggings (11123, -1.35 DPS) [quest]; Senior Designer's Pantaloons (11841, -1.66 DPS, sim-verified) [dungeon] |
| feet | Frothing Sandals (254109) | Tailoring [crafted] | 25.3 spell_power points (9.19 DPS) | yes | Mistwalker Boots (10629, -1.10 DPS) [dungeon]; Sergeant Major's Dreadweave Boots (220891, -1.24 DPS) [vendor]; Coldstone Slippers (18697, -2.39 DPS, sim-verified) [dungeon] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 20.6 spell_power points (7.47 DPS) | yes | Cyclopean Band (11824, -1.64 DPS) [dungeon]; Freezing Band (942, -2.75 DPS) [world_drop]; Band of the Unicorn (7553, -2.75 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 16.6 spell_power points (6.02 DPS) | yes | Cyclopean Band (11824, -0.19 DPS) [dungeon]; Freezing Band (942, -1.30 DPS) [world_drop]; Band of the Unicorn (7553, -1.30 DPS) [world_drop] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (+3.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.13 DPS) [quest]; Uther's Strength (11302, -1.91 DPS, sim-verified) [world_drop] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-verified (+2.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Spire of Hakkar (10844, +0.00 DPS) [world]; Staff of Jordan (873, -2.84 DPS, sim-verified) [world_drop] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 14.8 spell_power points (5.36 DPS) | yes | Truesilver Conduit (249455, -0.02 DPS) [crafted]; Orb of the Forgotten Seer (7685, -0.28 DPS) [dungeon]; Enthralled Sphere (11625, -0.46 DPS) [dungeon] |
| ranged | Freezing Shard (10572) | Razorfen Downs [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Icefury Wand (7514, +0.00 DPS) [quest]; Lesser Eternal Wand (249232, -0.36 DPS) [crafted]; Charged Lightning Rod (11860, -0.90 DPS) [quest] |

**New at 50:** head: Dreamweave Circlet; neck: Arcane Crystal Pendant; shoulder: Netherfroth Shoulders; back: Spritecaster Cape; chest: Embrace of the Wind Serpent; wrist: Netherfroth Cuffs; hands: Frothing Gloves; waist: Ban'thok Sash; legs: Spellshock Leggings; feet: Frothing Sandals; finger1: Brainlash; finger2: Lorekeeper's Ring; trinket1: Thunderbrew's Boot Flask; trinket2: Frozen Heart of the Mountain; main_hand: Blade of Eternal Darkness; off_hand: Celestial Orb; ranged: Freezing Shard

No-known-source sample (15 of 600, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (gnome, 250005203000300000-00000000000000000-253510130101301221)

Set DPS (verified): 425.8. Weights run: 2.7s. Verify run: 17.7s. 1459 eligible items had no known source.

Race sweep: gnome 425.8 ± 0.5, high-order-skyborne 424.7 ± 0.5, human 424.6 ± 0.5 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.393 ± 0.021, crit=0.567 ± 0.017 per rating point (14 rating = 1%, 7.941 per %), hit=1.156 ± 0.059 per rating point (10 rating = 1%, 11.559 per %), spell_haste=5.503 ± 0.329, spell_penetration=not significant (0.000 ± 0.000), frost_power=0.809 ± 0.001, spirit=not significant (0.468 ± 0.138), mp5=0.732 ± 0.160

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 50.6 spell_power points (21.37 DPS) | yes | Field Marshal's Coronet (231604, -0.09 DPS) [pvp]; Crimson Felt Hat (18727, -2.82 DPS, sim-verified) [dungeon]; Lieutenant Commander's Silk Cowl (227103, -4.98 DPS) [pvp] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 23.7 spell_power points (10.00 DPS) | yes | Star of Mystaria (12103, -1.12 DPS) [dungeon]; Lady Maye's Pendant (14558, -2.21 DPS) [world_drop]; Arcane Crystal Pendant (20037, -2.26 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 30.1 spell_power points (12.68 DPS) | yes | Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Mantle (227102, -0.39 DPS) [pvp]; Argent Shoulders (19059, -0.75 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -3.85 DPS) [world]; Shroud of the Exile (15421, -4.85 DPS) [quest]; Spritecaster Cape (11623, -4.87 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Vestments (231603, -0.94 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -6.20 DPS) [pvp]; Robe of Winter Night (14136, -9.34 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 28.4 spell_power points (11.99 DPS) | yes | Netherfroth Cuffs (254063, -5.41 DPS) [crafted]; Runecloth Cuffs (254123, -5.69 DPS) [crafted]; Wyrmthalak's Shackles (13958, -6.08 DPS) [quest] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 30.3 spell_power points (12.78 DPS) | yes | Gloves of Spell Mastery (14146, +0.00 DPS) [crafted]; Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 28.4 spell_power points (12.00 DPS) | yes | Highlander's Cloth Girdle (20047, -1.75 DPS) [rep]; Ban'thok Sash (11662, -2.24 DPS, sim-verified) [dungeon]; Stormpike Cloth Girdle (19094, -2.75 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 51.8 spell_power points (21.86 DPS) | yes | Marshal's Silk Leggings (231605, -0.56 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -5.84 DPS) [pvp]; Ghostweave Pants (14144, -6.94 DPS) [crafted] |
| feet | Venomspew Footpads (275606) | Tailoring [crafted] | 25.7 spell_power points (10.86 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Knight-Lieutenant's Silk Walkers (227112, -0.34 DPS) [pvp]; Omnicast Boots (11822, -0.43 DPS) [dungeon] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -1.90 DPS) [vendor]; Songstone of Ironforge (12543, -1.92 DPS) [quest]; Naglering (11669, -9.45 DPS, sim-verified) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -1.16 DPS) [vendor]; Songstone of Ironforge (12543, -1.18 DPS) [quest]; The Postmaster's Seal (13392, -1.59 DPS) [dungeon] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+3.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, -2.72 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -3.65 DPS, sim-verified) [crafted]; Serenity Field (272439, -4.64 DPS) [vendor] |
| trinket2 | Eye of the Beast (13968) | General Drakkisath's Demise [quest] | sim-verified (-2.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironbark Staff (20069, +0.00 DPS) [rep]; Grand Marshal's Stave (234571, -0.84 DPS) [pvp]; Blade of Eternal Darkness (17780, -6.82 DPS, sim-verified) [dungeon] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (12.24 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Tome of the Ice Lord (19310, -0.33 DPS) [rep]; Spirit of Aquementas (11904, -3.38 DPS) [quest] |
| ranged | Bonecreeper Stylus (13938) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wand of Biting Cold (19108, +0.00 DPS) [quest]; Oblivion's Touch (18761, -0.51 DPS) [dungeon]; Brilliant Wand (249385, -0.57 DPS) [crafted] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Magician's Cord; legs: Sentinel's Silk Leggings; feet: Venomspew Footpads; finger1: Blessed Band of Light; finger2: Maiden's Circle; trinket1: Weakness Analyzer; trinket2: Eye of the Beast; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Bonecreeper Stylus

No-known-source sample (15 of 1459, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (gnome, 230005203000000000-50000000000000000-253510130101301221)

Set DPS (verified): 696.4. Weights run: 2.9s. Verify run: 18.4s. 1459 eligible items had no known source.

Race sweep: high-order-skyborne 701.1 ± 0.6, gnome 696.4 ± 0.7, human 695.1 ± 0.6 (dps); high-order-skyborne beats the guide race gnome by 4.8 (0.68%) ± 0.9, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.172 ± 0.009, crit=0.674 ± 0.020 per rating point (14 rating = 1%, 9.431 per %), hit=1.582 ± 0.071 per rating point (10 rating = 1%, 15.820 per %), spell_haste=10.486 ± 0.314, spell_penetration=not significant (0.000 ± 0.000), frost_power=0.815 ± 0.002, spirit=not significant (0.000 ± 0.000), mp5=not significant (-0.001 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 48.8 spell_power points (24.94 DPS) | yes | Field Marshal's Coronet (231604, -1.77 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -7.81 DPS) [pvp]; Crimson Felt Hat (18727, -8.05 DPS, sim-verified) [dungeon] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 26.2 spell_power points (13.38 DPS) | yes | Star of Mystaria (12103, -3.39 DPS, sim-verified) [dungeon]; Arcane Crystal Pendant (20037, -4.68 DPS) [quest]; Emberfury Talisman (12929, -5.50 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 28.7 spell_power points (14.64 DPS) | yes | Field Marshal's Silk Spaulders (231602, -0.56 DPS) [pvp]; Lieutenant Commander's Silk Mantle (227102, -1.20 DPS) [pvp]; Argent Shoulders (19059, -1.87 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -6.17 DPS, sim-verified) [world]; Cloak of Flames (3475, -8.87 DPS) [world_drop]; Stalwart Cloak (272415, -8.87 DPS) [vendor] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Vestments (231603, -3.14 DPS) [pvp]; Robe of Winter Night (14136, -8.21 DPS, sim-verified) [crafted]; Knight-Captain's Silk Tunic (227108, -9.27 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 23.4 spell_power points (11.94 DPS) | yes | Runecloth Cuffs (254123, -5.44 DPS) [crafted]; Netherfroth Cuffs (254063, -5.92 DPS) [crafted]; Manacle Cuffs (11962, -7.05 DPS) [quest] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 29.6 spell_power points (15.11 DPS) | yes | Hands of Power (13253, +0.00 DPS) [dungeon]; Marshal's Silk Gloves (16440, -0.26 DPS) [vendor]; Marshal's Silk Gauntlets (231608, -0.26 DPS) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 29.5 spell_power points (15.09 DPS) | yes | Highlander's Cloth Girdle (20047, -2.59 DPS) [rep]; Magician's Cord (272393, -3.28 DPS) [vendor]; Stormpike Cloth Girdle (19094, -5.02 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 52.6 spell_power points (26.86 DPS) | yes | Marshal's Silk Leggings (231605, -4.96 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -9.82 DPS) [pvp]; Skyshroud Leggings (13170, -10.54 DPS, sim-verified) [dungeon] |
| feet | Sergeant Major's Dreadweave Boots (220891) | PvP rank 5 · Sergeant Major · Alliance [vendor] | 25.4 spell_power points (12.96 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Kayser's Boots of Precision (22231, -1.25 DPS) [dungeon]; Knight-Lieutenant's Silk Walkers (227112, -1.35 DPS) [pvp] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -2.75 DPS) [dungeon]; Maiden's Circle (13001, -3.38 DPS) [world_drop]; Naglering (11669, -12.68 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -2.46 DPS) [dungeon]; Songstone of Ironforge (12543, -3.09 DPS) [quest]; Maiden's Circle (13001, -3.09 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cannonball Runner (13382, -3.23 DPS) [dungeon]; Uther's Strength (11302, -4.21 DPS) [world_drop]; Smolderweb's Eye (13213, -7.43 DPS, sim-verified) [dungeon] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+2.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, -2.87 DPS, sim-verified) [dungeon]; Cannonball Runner (13382, -5.15 DPS) [dungeon]; Serenity Field (272439, -5.62 DPS) [vendor] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -2.87 DPS) [pvp]; Headmaster's Charge (13937, -9.11 DPS, sim-verified) [dungeon]; Staff of Jordan (873, -11.38 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Biting Cold (19108) | Korrak the Bloodrager [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, -0.60 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -1.04 DPS) [dungeon]; Oblivion's Touch (18761, -1.45 DPS) [dungeon] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; legs: Sentinel's Silk Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blessed Band of Light; finger2: Don Julio's Band; trinket1: Frozen Heart of the Mountain; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Wand of Biting Cold

No-known-source sample (15 of 1459, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (undead, 000000000000000000-00000000000000000-053200010000000000)

Set DPS (verified): 50.0. Weights run: 2.2s. Verify run: 2.9s. 192 eligible items had no known source.

Race sweep: undead 50.0 ± 0.1, orc 48.7 ± 0.1, troll 48.2 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.982 ± 0.007, crit=0.104 ± 0.003 per rating point (14 rating = 1%, 1.458 per %), hit=0.256 ± 0.008 per rating point (10 rating = 1%, 2.564 per %), spell_haste=0.783 ± 0.086, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002, spirit=not significant (0.023 ± 0.007), mp5=1.328 ± 0.104

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Silky Circlet (253951) | Tailoring [crafted] | 11.0 spell_power points (2.32 DPS) | yes | Pristine Circlet (253949, -1.11 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.26 DPS) [crafted]; Flying Tiger Goggles (4368, -2.30 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.1 spell_power points (0.01 DPS) | yes | Scholarly Pendant (277203, -0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.8 spell_power points (2.92 DPS) | yes | Reinforced Woolen Shoulders (4315, -1.01 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -2.08 DPS) [crafted]; Slime-encrusted Pads (6461, -2.08 DPS) [dungeon] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sanguine Cape (14376, -0.22 DPS) [world_drop]; Heavy Woolen Cloak (4311, -0.39 DPS, sim-verified) [crafted]; Battle Healer's Cloak (20427, -0.40 DPS) [rep] |
| chest | Filigreed Silky Gown (253903) | Tailoring [crafted] | 14.0 spell_power points (2.95 DPS) | yes | Filigreed Pristine Gown (253901, -0.89 DPS, sim-verified) [crafted]; Gray Woolen Robe (2585, -1.07 DPS) [crafted]; Mystic's Wrap (14369, -1.50 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 8.9 spell_power points (1.88 DPS) | yes | Bright Bracers (3647, -0.96 DPS, sim-verified) [world_drop]; Tabitha's Cuffs (251486, -1.24 DPS) [quest]; Owlbeard Bracers (16981, -1.25 DPS) [quest] |
| hands | Silky Gloves (253915) | Tailoring [crafted] | 9.9 spell_power points (2.10 DPS) | yes | Pristine Gloves (253913, -0.63 DPS) [crafted]; Gnoll Casting Gloves (892, -0.83 DPS) [world]; Serpent Gloves (5970, -0.99 DPS, sim-verified) [dungeon] |
| waist | Silky Sash (253927) | Tailoring [crafted] | 10.9 spell_power points (2.31 DPS) | yes | Keller's Girdle (2911, -0.65 DPS) [world_drop]; Pristine Sash (253925, -0.66 DPS, sim-verified) [crafted]; Novice Arcanist's Sash (253885, -0.84 DPS) [crafted] |
| legs | Filigreed Silky Leggings (253939) | Tailoring [crafted] | 17.0 spell_power points (3.59 DPS) | yes | Filigreed Pristine Leggings (253937, -1.11 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -2.11 DPS) [world_drop]; Silk-threaded Trousers (1929, -2.11 DPS) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.9 spell_power points (2.31 DPS) | yes | Silky Boots (253891, -0.47 DPS, sim-verified) [crafted]; Pristine Boots (253889, -1.05 DPS) [crafted]; Red Woolen Boots (4313, -1.44 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 7.7 spell_power points (1.62 DPS) | yes | Loop of Sacrifice (281673, -0.58 DPS) [quest]; Volcanic Rock Ring (12053, -1.00 DPS) [world_drop]; Furlfeather Ring (286745, -1.41 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.9 spell_power points (1.24 DPS) | yes | Loop of Sacrifice (281673, -0.25 DPS, sim-verified) [quest]; Volcanic Rock Ring (12053, -0.62 DPS) [world_drop]; Furlfeather Ring (286745, -1.03 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Staff of the Blessed Seer (2271, -0.39 DPS) [dungeon]; Crescent Staff (6505, -0.68 DPS, sim-verified) [quest]; Gnarled Necromancer's Staff (251534, -0.81 DPS) [quest] |
| off_hand | Antipodean Rod (2879) | World drop [world_drop] | 6.0 spell_power points (1.27 DPS) | yes | Seer's Fine Stein (7608, -0.64 DPS) [world_drop]; Spellbinder Orb (15926, -0.65 DPS) [dungeon]; Tork Wrench (11855, -0.85 DPS) [quest] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller (12984, -0.29 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -0.73 DPS) [quest]; Greater Magic Wand (11288, -0.76 DPS) [crafted] |

**New at 20:** head: Silky Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Silky Gown; wrist: Mindthrust Bracers; hands: Silky Gloves; waist: Silky Sash; legs: Filigreed Silky Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Evocator's Blade; off_hand: Antipodean Rod; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 192, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 8350 The 1 Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 23173 Abomination Skin Leggings

### Band 30 (undead, 000000000000000000-00000000000000000-053510030100200100)

Set DPS (verified): 88.0. Weights run: 2.3s. Verify run: 2.4s. 318 eligible items had no known source.

Race sweep: undead 88.0 ± 0.1, troll 86.6 ± 0.1, orc 86.5 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.270 ± 0.008, crit=0.108 ± 0.004 per rating point (14 rating = 1%, 1.515 per %), hit=0.228 ± 0.009 per rating point (10 rating = 1%, 2.279 per %), spell_haste=1.125 ± 0.086, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003, spirit=not significant (-0.045 ± 0.035), mp5=not significant (0.191 ± 0.064)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Silky Circlet (253977) | Tailoring [crafted] | 14.0 spell_power points (3.02 DPS) | yes | Holy Shroud (2721, -0.65 DPS) [world_drop]; Silky Circlet (253951, -0.85 DPS, sim-verified) [crafted]; Azure Silk Hood (7048, -0.86 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 1.6 spell_power points (0.35 DPS) | yes | Darkspear Warding Pendant (272075, -0.06 DPS) [vendor]; Crystal Starfire Medallion (5003, -0.12 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.12 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.4 spell_power points (2.46 DPS) | yes | Death Speaker Mantle (6685, -0.53 DPS) [dungeon]; Invoker's Mantle (215365, -0.66 DPS) [crafted]; Magician's Mantle (12998, -0.86 DPS) [world_drop] |
| back | Azure Silk Cloak (7053) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windsong Drape (15468, -1.08 DPS) [quest]; Heavy Woolen Cloak (4311, -1.29 DPS) [crafted]; Hillman's Cloak (3719, -1.31 DPS, sim-verified) [crafted] |
| chest | Silky Gown (253963) | Tailoring [crafted] | 14.9 spell_power points (3.21 DPS) | yes | Green Silk Armor (7065, -0.51 DPS) [crafted]; Beguiler Robes (7728, -0.57 DPS) [dungeon]; Tree Bark Jacket (1486, -1.96 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.6 spell_power points (2.06 DPS) | yes | Glowing Magical Bracelets (13106, -1.60 DPS) [world_drop]; Mindthrust Bracers (1974, -1.65 DPS) [dungeon]; Tabitha's Cuffs (251486, -2.12 DPS, sim-verified) [quest] |
| hands | Azure Silk Gloves (4319) | Tailoring [crafted] | 10.0 spell_power points (2.16 DPS) | yes | Silky Gloves (253915, -0.47 DPS) [crafted]; Jutebraid Gloves (10654, -0.57 DPS) [quest]; Serpent Gloves (5970, -0.65 DPS) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.8 spell_power points (2.55 DPS) | yes | Belt of Arugal (6392, -0.43 DPS) [dungeon]; Warsong Sash (16975, -0.59 DPS, sim-verified) [quest]; Invoker's Cord (215366, -0.75 DPS) [crafted] |
| legs | Silky Leggings (253989) | Tailoring [crafted] | 15.9 spell_power points (3.43 DPS) | yes | Gaze Dreamer Pants (6903, -0.84 DPS) [dungeon]; Azure Silk Pants (7046, -0.92 DPS) [crafted]; Filigreed Silky Leggings (253939, -0.97 DPS, sim-verified) [crafted] |
| feet | Frothing Slippers (254003) | Tailoring [crafted] | 14.9 spell_power points (3.21 DPS) | yes | Spidersilk Boots (4320, -1.47 DPS) [crafted]; Gilded Slippers (254001, -1.62 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.67 DPS) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.4 spell_power points (1.59 DPS) | yes | Ring of Power Regulation (273030, -0.17 DPS) [dungeon]; Runemender's Seal (273031, -0.60 DPS) [world]; Solid Gold Band (279845, -0.73 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 7.3 spell_power points (1.58 DPS) | yes | Ring of Power Regulation (273030, -0.17 DPS) [dungeon]; Runemender's Seal (273031, -0.59 DPS) [world]; Solid Gold Band (279845, -0.72 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -0.86 DPS) [dungeon]; Tendonscraper Dagger (273053, -0.86 DPS) [world]; Glimmering Staff (249392, -2.64 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.6 spell_power points (1.86 DPS) | yes | Orb of Souls (249395, -1.00 DPS) [crafted]; Witch's Finger (16887, -1.39 DPS) [quest]; Antipodean Rod (2879, -1.73 DPS, sim-verified) [world_drop] |
| ranged | Greater Mystic Wand (217287) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderwood (13062, -0.00 DPS) [world_drop]; Wand of Eventide (5214, -0.37 DPS) [world_drop]; Starfaller (13063, -0.40 DPS) [world_drop] |

**New at 30:** head: Filigreed Silky Circlet; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Azure Silk Cloak; chest: Silky Gown; wrist: Spidertank Oilrag; hands: Azure Silk Gloves; waist: Defiler's Cloth Girdle; legs: Silky Leggings; feet: Frothing Slippers; finger1: Advisor's Ring; finger2: Snake Hoop; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Greater Mystic Wand

No-known-source sample (15 of 318, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 8350 The 1 Ring; 9362 Brilliant Gold Ring

### Band 40 (undead, 000000000000000000-00000000000000000-253510130101301221)

Set DPS (verified): 173.3. Weights run: 2.4s. Verify run: 4.8s. 448 eligible items had no known source.

Race sweep: undead 173.3 ± 0.2, orc 171.5 ± 0.2, troll 170.8 ± 0.2 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.622 ± 0.016, crit=0.166 ± 0.006 per rating point (14 rating = 1%, 2.330 per %), hit=0.374 ± 0.018 per rating point (10 rating = 1%, 3.735 per %), spell_haste=0.943 ± 0.207, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003, spirit=not significant (0.229 ± 0.158), mp5=not significant (0.791 ± 0.200)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Papal Fez (9431, -0.09 DPS) [dungeon]; Miner's Hat of the Deep (9429, -0.44 DPS, sim-verified) [dungeon]; Corpseshroud (10574, -0.56 DPS) [dungeon] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Prodigious Shadowshard Pendant (17773, -0.16 DPS) [quest]; Necklace of Calisea (1714, -0.94 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -1.11 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.1 spell_power points (3.88 DPS) | yes | Green Silken Shoulders (7057, -0.06 DPS) [crafted]; Bloodmage Mantle (7684, -0.13 DPS) [dungeon]; Berylline Pads (4197, -0.48 DPS) [quest] |
| back | Frost Leather Cloak (7377) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Azure Silk Cloak (7053, -0.08 DPS) [crafted]; Long Silken Cloak (4326, -0.49 DPS) [crafted]; Guardian Cloak (5965, -0.49 DPS) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Robe of Power (7054, -0.90 DPS) [crafted]; Robe of the Magi (1716, -1.11 DPS, sim-verified) [world_drop]; Elemental Raiment (9434, -1.50 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 11.4 spell_power points (2.93 DPS) | yes | Condor Bracers (15864, -1.13 DPS) [quest]; Windchaser Cuffs (14429, -1.49 DPS) [world_drop]; Radiant Silver Bracers (4545, -1.50 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.1 spell_power points (5.69 DPS) | yes | Red Mageweave Gloves (10018, -1.25 DPS) [crafted]; Black Mageweave Gloves (10003, -1.83 DPS) [crafted]; Frothing Handwraps (254023, -2.67 DPS, sim-verified) [crafted] |
| waist | Frothing Cord (254039) | Tailoring [crafted] | 20.4 spell_power points (5.24 DPS) | yes | Deathmage Sash (10771, -1.04 DPS) [dungeon]; Gilded Cord (254037, -1.54 DPS) [crafted]; Defiler's Cloth Girdle (20166, -2.31 DPS, sim-verified) [rep] |
| legs | Silky Leggings (253989) | Tailoring [crafted] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Red Mageweave Pants (10009, -0.56 DPS, sim-verified) [crafted]; Crimson Silk Pantaloons (7062, -0.88 DPS) [crafted]; Filigreed Silky Leggings (253939, -0.99 DPS) [crafted] |
| feet | Frothing Slippers (254003) | Tailoring [crafted] | 18.3 spell_power points (4.70 DPS) | yes | Acidic Walkers (9454, -1.90 DPS) [dungeon]; Gilded Slippers (254001, -2.20 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -2.26 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.1 spell_power points (3.38 DPS) | yes | Snake Hoop (6750, -1.03 DPS) [quest]; Ring of Power Regulation (273030, -1.22 DPS) [dungeon]; Runemender's Seal (273031, -1.29 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 11.4 spell_power points (2.93 DPS) | yes | Ring of Power Regulation (273030, -0.77 DPS) [dungeon]; Runemender's Seal (273031, -0.84 DPS) [world]; Snake Hoop (6750, -1.44 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+2.8 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Celestial Stave (9517, -7.56 DPS) [quest]; Hypnotic Blade (7714, -7.90 DPS) [dungeon]; Illusionary Rod (7713, -10.60 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Icefury Wand (7514) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Freezing Shard (10572, -0.31 DPS) [dungeon]; Plaguerot Sprig (10766, -1.11 DPS) [dungeon]; Twisted Nether Wand (249144, -1.46 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Triune Amulet; shoulder: Inquisitor's Shawl; back: Frost Leather Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Frothing Cord; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Staff of Jordan; ranged: Icefury Wand

No-known-source sample (15 of 448, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (undead, 230005000000000000-00000000000000000-253510130101301221)

Set DPS (verified): 262.7. Weights run: 2.4s. Verify run: 3.9s. 592 eligible items had no known source.

Race sweep: undead 262.7 ± 0.4, orc 260.7 ± 0.4, troll 260.2 ± 0.4 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.589 ± 0.017, crit=0.400 ± 0.013 per rating point (14 rating = 1%, 5.607 per %), hit=0.859 ± 0.036 per rating point (10 rating = 1%, 8.593 per %), spell_haste=3.497 ± 0.267, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002, spirit=0.737 ± 0.175, mp5=1.148 ± 0.192

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreamweave Circlet (10041) | Tailoring [crafted] | 35.7 spell_power points (12.98 DPS) | yes | Red Mageweave Headband (10033, -1.94 DPS, sim-verified) [crafted]; Gemburst Circlet (10751, -3.07 DPS) [quest]; Spellpower Goggles Xtreme Plus (15999, -3.17 DPS) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 19.5 spell_power points (7.09 DPS) | yes | Triune Amulet (7722, -2.31 DPS) [dungeon]; Jarkal's Enhancing Necklace (7888, -2.52 DPS) [quest]; Mindburst Medallion (11196, -3.69 DPS, sim-verified) [quest] |
| shoulder | Netherfroth Shoulders (254051) | Tailoring [crafted] | 26.1 spell_power points (9.49 DPS) | yes | Kentic Amice (11624, -0.02 DPS) [dungeon]; Rotgrip Mantle (17732, -0.92 DPS) [dungeon]; Mantle of Lost Hope (22234, -2.20 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -1.75 DPS, sim-verified) [quest]; Featherskin Cape (10843, -2.87 DPS) [world]; Runecloth Cloak (13860, -2.99 DPS) [crafted] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 39.0 spell_power points (14.17 DPS) | yes | Ghostweave Vest (14141, -2.21 DPS) [crafted]; Dreamweave Vest (10021, -2.25 DPS, sim-verified) [crafted]; Frostweave Robe (13868, -2.51 DPS) [crafted] |
| wrist | Netherfroth Cuffs (254063) | Tailoring [crafted] | 20.8 spell_power points (7.56 DPS) | yes | Nethergeld Cuffs (254061, -2.25 DPS, sim-verified) [crafted]; Aristocratic Cuffs (12546, -2.90 DPS) [dungeon]; Bloodband Bracers (11469, -3.01 DPS) [quest] |
| hands | Frothing Gloves (254097) | Tailoring [crafted] | 30.8 spell_power points (11.18 DPS) | yes | Frostweave Gloves (13870, -2.27 DPS) [crafted]; Runecloth Gloves (13863, -2.49 DPS) [crafted]; Dreamweave Gloves (10019, -3.21 DPS, sim-verified) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 26.5 spell_power points (9.62 DPS) | yes | Frothing Cord (254039, -1.22 DPS) [crafted]; Ghostweave Cord (254073, -1.59 DPS) [crafted]; Frothing Waistcord (254083, -1.67 DPS, sim-verified) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 29.4 spell_power points (10.67 DPS) | yes | Spellshock Leggings (9484, -0.18 DPS) [dungeon]; Kilt of the Atal'ai Prophet (10807, -0.91 DPS) [dungeon]; Rainstrider Leggings (11123, -1.53 DPS) [quest] |
| feet | Frothing Sandals (254109) | Tailoring [crafted] | 25.3 spell_power points (9.19 DPS) | yes | Mistwalker Boots (10629, -1.10 DPS) [dungeon]; First Sergeant's Dreadweave Boots (220909, -1.24 DPS) [vendor]; Coldstone Slippers (18697, -2.60 DPS, sim-verified) [dungeon] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 20.6 spell_power points (7.47 DPS) | yes | Cyclopean Band (11824, -1.64 DPS) [dungeon]; Freezing Band (942, -2.75 DPS) [world_drop]; Band of the Unicorn (7553, -2.75 DPS) [world_drop] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 16.6 spell_power points (6.02 DPS) | yes | Cyclopean Band (11824, -0.19 DPS) [dungeon]; Freezing Band (942, -1.30 DPS) [world_drop]; Band of the Unicorn (7553, -1.30 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.13 DPS) [quest]; Uther's Strength (11302, -2.06 DPS, sim-verified) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, -0.01 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -2.10 DPS, sim-verified) [quest] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Spire of Hakkar (10844, +0.00 DPS) [world]; Staff of Jordan (873, -6.04 DPS, sim-verified) [world_drop] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 14.8 spell_power points (5.36 DPS) | yes | Truesilver Conduit (249455, -0.02 DPS) [crafted]; Orb of the Forgotten Seer (7685, -0.28 DPS) [dungeon]; Flask of Forest Mojo (19115, -0.38 DPS) [quest] |
| ranged | Icefury Wand (7514) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Freezing Shard (10572, -0.36 DPS) [dungeon]; Lesser Eternal Wand (249232, -0.51 DPS) [crafted]; Charged Lightning Rod (11860, -1.23 DPS) [quest] |

**New at 50:** head: Dreamweave Circlet; neck: Arcane Crystal Pendant; shoulder: Netherfroth Shoulders; back: Spritecaster Cape; chest: Embrace of the Wind Serpent; wrist: Netherfroth Cuffs; hands: Frothing Gloves; waist: Ban'thok Sash; legs: Senior Designer's Pantaloons; feet: Frothing Sandals; finger1: Brainlash; finger2: Advisor's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Blade of Eternal Darkness; off_hand: Celestial Orb

No-known-source sample (15 of 592, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (undead, 250005203000300000-00000000000000000-253510130101301221)

Set DPS (verified): 420.5. Weights run: 2.7s. Verify run: 17.6s. 1453 eligible items had no known source.

Race sweep: undead 420.5 ± 0.5, troll 417.3 ± 0.5, orc 417.0 ± 0.5 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.393 ± 0.021, crit=0.567 ± 0.017 per rating point (14 rating = 1%, 7.941 per %), hit=1.156 ± 0.059 per rating point (10 rating = 1%, 11.559 per %), spell_haste=5.503 ± 0.329, spell_penetration=not significant (0.000 ± 0.000), frost_power=0.809 ± 0.001, spirit=not significant (0.468 ± 0.138), mp5=0.732 ± 0.160

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 50.6 spell_power points (21.37 DPS) | yes | Warlord's Silk Cowl (231601, -0.09 DPS) [pvp]; Crimson Felt Hat (18727, -2.06 DPS, sim-verified) [dungeon]; Champion's Silk Cowl (227105, -4.98 DPS) [pvp] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 23.7 spell_power points (10.00 DPS) | yes | Star of Mystaria (12103, -2.14 DPS, sim-verified) [dungeon]; Lady Maye's Pendant (14558, -2.21 DPS) [world_drop]; Arcane Crystal Pendant (20037, -2.26 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 30.1 spell_power points (12.68 DPS) | yes | Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Champion's Silk Mantle (227104, -0.39 DPS) [pvp]; Argent Shoulders (19059, -0.75 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -3.85 DPS) [world]; Shroud of the Exile (15421, -4.85 DPS) [quest]; Spritecaster Cape (11623, -4.87 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Raiment (231596, -0.94 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -6.20 DPS) [pvp]; Robe of Winter Night (14136, -8.60 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 28.4 spell_power points (11.99 DPS) | yes | Netherfroth Cuffs (254063, -5.41 DPS) [crafted]; Runecloth Cuffs (254123, -5.69 DPS) [crafted]; Wyrmthalak's Shackles (13958, -6.08 DPS) [quest] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 30.3 spell_power points (12.78 DPS) | yes | Gloves of Spell Mastery (14146, +0.00 DPS) [crafted]; General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 28.4 spell_power points (12.00 DPS) | yes | Ban'thok Sash (11662, -0.40 DPS) [dungeon]; Defiler's Cloth Girdle (20163, -1.75 DPS) [rep]; Frostwolf Cloth Belt (19090, -2.75 DPS) [rep] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 40.2 spell_power points (16.94 DPS) | yes | General's Silk Trousers (231595, +0.00 DPS) [pvp]; Legionnaire's Silk Legguards (227107, -0.92 DPS) [pvp]; Ghostweave Pants (14144, -3.61 DPS, sim-verified) [crafted] |
| feet | Venomspew Footpads (275606) | Tailoring [crafted] | 25.7 spell_power points (10.86 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Blood Guard's Silk Walkers (227110, -0.34 DPS) [pvp]; Omnicast Boots (11822, -0.43 DPS) [dungeon] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -1.90 DPS) [vendor]; Eye of Orgrimmar (12545, -1.92 DPS) [quest]; Naglering (11669, -7.59 DPS, sim-verified) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -1.18 DPS) [quest]; The Postmaster's Seal (13392, -1.59 DPS) [dungeon]; Channeler's Ring (272406, -3.50 DPS, sim-verified) [vendor] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+5.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -3.21 DPS) [crafted]; Rune of the Guard Captain (19120, -4.18 DPS) [quest]; Smolderweb's Eye (13213, -5.04 DPS, sim-verified) [dungeon] |
| trinket2 | Eye of the Beast (13968) | For The Horde! [quest] | sim-verified (-2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -0.84 DPS) [pvp]; Headmaster's Charge (13937, -3.61 DPS) [dungeon]; Ironbark Staff (20220, -10.13 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (12.24 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; Tome of the Ice Lord (19310, -0.33 DPS) [rep]; Spirit of Aquementas (11904, -3.38 DPS) [quest] |
| ranged | Wand of Biting Cold (19108) | The Legend of Korrak [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, -0.01 DPS) [dungeon]; Oblivion's Touch (18761, -0.53 DPS) [dungeon]; Brilliant Wand (249385, -0.58 DPS) [crafted] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Magician's Cord; legs: Outrider's Silk Leggings; feet: Venomspew Footpads; finger1: Blessed Band of Light; finger2: Maiden's Circle; trinket1: Weakness Analyzer; trinket2: Eye of the Beast; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Wand of Biting Cold

No-known-source sample (15 of 1453, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60, raid preset (undead, 230005203000000000-50000000000000000-253510130101301221)

Set DPS (verified): 688.8. Weights run: 2.9s. Verify run: 18.2s. 1453 eligible items had no known source.

Race sweep: undead 688.8 ± 0.6, troll 685.1 ± 0.6, orc 684.9 ± 0.6 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.172 ± 0.009, crit=0.674 ± 0.020 per rating point (14 rating = 1%, 9.431 per %), hit=1.582 ± 0.071 per rating point (10 rating = 1%, 15.820 per %), spell_haste=10.486 ± 0.314, spell_penetration=not significant (0.000 ± 0.000), frost_power=0.815 ± 0.002, spirit=not significant (0.000 ± 0.000), mp5=not significant (-0.001 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 48.8 spell_power points (24.94 DPS) | yes | Warlord's Silk Cowl (231601, -1.77 DPS) [pvp]; Champion's Silk Cowl (227105, -7.81 DPS) [pvp]; Crimson Felt Hat (18727, -8.86 DPS, sim-verified) [dungeon] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 26.2 spell_power points (13.38 DPS) | yes | Star of Mystaria (12103, -3.19 DPS, sim-verified) [dungeon]; Arcane Crystal Pendant (20037, -4.68 DPS) [quest]; Emberfury Talisman (12929, -5.50 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 28.7 spell_power points (14.64 DPS) | yes | Warlord's Silk Amice (231594, -0.56 DPS) [pvp]; Champion's Silk Mantle (227104, -1.20 DPS) [pvp]; Argent Shoulders (19059, -1.87 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -6.15 DPS, sim-verified) [world]; Cloak of Flames (3475, -8.87 DPS) [world_drop]; Stalwart Cloak (272415, -8.87 DPS) [vendor] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Raiment (231596, -3.14 DPS) [pvp]; Robe of Winter Night (14136, -8.87 DPS, sim-verified) [crafted]; Legionnaire's Silk Tunic (227106, -9.27 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 23.4 spell_power points (11.94 DPS) | yes | Runecloth Cuffs (254123, -5.44 DPS) [crafted]; Netherfroth Cuffs (254063, -5.92 DPS) [crafted]; Manacle Cuffs (11962, -7.05 DPS) [quest] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 29.6 spell_power points (15.11 DPS) | yes | General's Silk Handguards (16540, -0.26 DPS) [vendor]; General's Silk Gauntlets (231599, -0.26 DPS) [vendor]; Hands of Power (13253, -2.96 DPS, sim-verified) [dungeon] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 29.5 spell_power points (15.09 DPS) | yes | Defiler's Cloth Girdle (20163, -3.17 DPS, sim-verified) [rep]; Magician's Cord (272393, -3.28 DPS) [vendor]; Frostwolf Cloth Belt (19090, -5.02 DPS) [rep] |
| legs | Skyshroud Leggings (13170) | Blackrock Spire: Highlord Omokk [dungeon] | 32.1 spell_power points (16.38 DPS) | yes | Legionnaire's Silk Legguards (227107, +0.00 DPS) [pvp]; General's Silk Trousers (231595, +0.00 DPS) [pvp]; Outrider's Silk Leggings (22747, -0.41 DPS) [rep] |
| feet | First Sergeant's Dreadweave Boots (220909) | PvP rank 5 · First Sergeant · Horde [vendor] | 25.4 spell_power points (12.96 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Kayser's Boots of Precision (22231, -1.25 DPS) [dungeon]; Blood Guard's Silk Walkers (227110, -1.35 DPS) [pvp] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -2.75 DPS) [dungeon]; Maiden's Circle (13001, -3.38 DPS) [world_drop]; Naglering (11669, -13.34 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -2.46 DPS) [dungeon]; Eye of Orgrimmar (12545, -3.09 DPS) [quest]; Maiden's Circle (13001, -3.09 DPS) [world_drop] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+3.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -1.92 DPS) [crafted]; Rune of the Guard Captain (19120, -3.54 DPS) [quest]; Eye of the Beast (13968, -3.57 DPS, sim-verified) [quest] |
| trinket2 | Smolderweb's Eye (13213) | Blackrock Spire: Mother Smolderweb [dungeon] | sim-verified (-3.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -0.81 DPS) [crafted]; Rune of the Guard Captain (19120, -2.42 DPS) [quest]; Cannonball Runner (13382, -4.04 DPS) [dungeon] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -2.87 DPS) [pvp]; Headmaster's Charge (13937, -10.63 DPS, sim-verified) [dungeon]; Staff of Jordan (873, -11.38 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Biting Cold (19108) | The Legend of Korrak [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, -0.60 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -1.04 DPS) [dungeon]; Oblivion's Touch (18761, -1.45 DPS) [dungeon] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; legs: Skyshroud Leggings; feet: First Sergeant's Dreadweave Boots; finger1: Blessed Band of Light; finger2: Don Julio's Band; trinket1: Weakness Analyzer; trinket2: Smolderweb's Eye; main_hand: Ironbark Staff; ranged: Wand of Biting Cold

No-known-source sample (15 of 1453, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

