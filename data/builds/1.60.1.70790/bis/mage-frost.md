# Leveling BiS: Frost

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-053200010000000000)

Set DPS (verified): 50.6. Weights run: 2.0s. Verify run: 1.8s. 198 eligible items had no known source.

Race sweep: gnome 50.6 ± 0.1, high-order-skyborne 49.1 ± 0.1, human 48.7 ± 0.1 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.977 ± 0.007, crit=0.105 ± 0.003 per rating point (14 rating = 1%, 1.470 per %), hit=0.255 ± 0.008 per rating point (10 rating = 1%, 2.550 per %), spell_haste=1.090 ± 0.098, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002, spirit=not significant (0.051 ± 0.036), mp5=0.856 ± 0.080

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Silky Circlet (253951) | Tailoring [crafted] | 11.0 spell_power points (2.46 DPS) | yes | Pristine Circlet (253949, -1.13 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.30 DPS) [crafted]; Flying Tiger Goggles (4368, -2.42 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.1 spell_power points (0.02 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.8 spell_power points (3.09 DPS) | yes | Reinforced Woolen Shoulders (4315, -1.46 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -2.19 DPS) [crafted]; Slime-encrusted Pads (6461, -2.51 DPS) [dungeon] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.16 DPS) [crafted]; Sanguine Cape (14376, -0.23 DPS) [world_drop]; Caretaker's Cape (20428, -0.41 DPS) [rep] |
| chest | Filigreed Silky Gown (253903) | Tailoring [crafted] | 14.0 spell_power points (3.14 DPS) | yes | Filigreed Pristine Gown (253901, -0.91 DPS, sim-verified) [crafted]; Gray Woolen Robe (2585, -1.15 DPS) [crafted]; Mystic's Wrap (14369, -1.61 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 7.5 spell_power points (1.67 DPS) | yes | Bright Bracers (3647, -0.85 DPS, sim-verified) [world_drop]; Repurposed Hair Band (281256, -1.21 DPS) [quest]; Mystic's Bracelets (14366, -1.23 DPS) [world_drop] |
| hands | Silky Gloves (253915) | Tailoring [crafted] | 9.9 spell_power points (2.22 DPS) | yes | Pristine Gloves (253913, -0.67 DPS) [crafted]; Serpent Gloves (5970, -0.75 DPS, sim-verified) [dungeon]; Gnoll Casting Gloves (892, -0.88 DPS) [world] |
| waist | Silky Sash (253927) | Tailoring [crafted] | 10.9 spell_power points (2.44 DPS) | yes | Pristine Sash (253925, -0.68 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.69 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.89 DPS) [crafted] |
| legs | Filigreed Silky Leggings (253939) | Tailoring [crafted] | 17.1 spell_power points (3.82 DPS) | yes | Filigreed Pristine Leggings (253937, -1.13 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -2.22 DPS) [world_drop]; Silk-threaded Trousers (1929, -2.25 DPS) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.9 spell_power points (2.44 DPS) | yes | Silky Boots (253891, -0.75 DPS, sim-verified) [crafted]; Pristine Boots (253889, -1.11 DPS) [crafted]; Red Woolen Boots (4313, -1.49 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 7.0 spell_power points (1.56 DPS) | yes | Lavishly Jeweled Ring (1156, -0.24 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.90 DPS) [world_drop]; Furlfeather Ring (286745, -1.33 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 6.7 spell_power points (1.50 DPS) | yes | Lavishly Jeweled Ring (1156, -0.19 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.85 DPS) [world_drop]; Furlfeather Ring (286745, -1.28 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | 30.2 spell_power points (6.76 DPS) | yes | Emberstone Staff (5201, -1.39 DPS) [dungeon]; Buzzer Blade (2169, -1.84 DPS) [dungeon]; Staff of the Blessed Seer (2271, -1.87 DPS, sim-verified) [dungeon] |
| off_hand | Antipodean Rod (2879) | World drop [world_drop] | 6.0 spell_power points (1.34 DPS) | yes | Mathystra Channeling Orb (281291, -0.25 DPS) [quest]; Tear of Grief (5611, -0.28 DPS) [quest]; Dwarven Tome (279898, -0.66 DPS) [quest] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller (12984, -0.15 DPS) [world_drop]; Sizzle Stick (8071, -0.51 DPS) [quest]; Greater Magic Wand (11288, -0.54 DPS) [crafted] |

**New at 20:** head: Silky Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Silky Gown; wrist: Mindthrust Bracers; hands: Silky Gloves; waist: Silky Sash; legs: Filigreed Silky Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Evocator's Blade; off_hand: Antipodean Rod; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 198, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 8350 The 1 Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (gnome, 000000000000000000-00000000000000000-053510030100200100)

Set DPS (verified): 89.6. Weights run: 2.0s. Verify run: 1.7s. 326 eligible items had no known source.

Race sweep: gnome 89.6 ± 0.1, high-order-skyborne 88.0 ± 0.1, human 87.3 ± 0.1 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.495 ± 0.010, crit=0.193 ± 0.006 per rating point (14 rating = 1%, 2.708 per %), hit=0.371 ± 0.014 per rating point (10 rating = 1%, 3.710 per %), spell_haste=1.927 ± 0.088, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002, spirit=0.310 ± 0.065, mp5=0.676 ± 0.083

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Silky Circlet (253977) | Tailoring [crafted] | 14.0 spell_power points (4.09 DPS) | yes | Holy Shroud (2721, -0.84 DPS, sim-verified) [world_drop]; Silky Circlet (253951, -0.88 DPS) [crafted]; Enchanter's Cowl (4322, -0.89 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 3.2 spell_power points (0.94 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.07 DPS) [world]; Crystal Starfire Medallion (5003, -0.09 DPS) [world_drop]; Seal of the Expedition (271920, -0.22 DPS) [rep] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.5 spell_power points (3.93 DPS) | yes | Death Speaker Mantle (6685, -0.66 DPS, sim-verified) [dungeon]; Invoker's Mantle (215365, -1.16 DPS) [crafted]; Magician's Mantle (12998, -1.17 DPS) [world_drop] |
| back | Azure Silk Cloak (7053) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Thresher Cape (6901, -1.59 DPS) [dungeon]; Heavy Woolen Cloak (4311, -1.66 DPS) [crafted]; Prelacy Cape (7004, -1.70 DPS, sim-verified) [quest] |
| chest | Silky Gown (253963) | Tailoring [crafted] | 17.7 spell_power points (5.17 DPS) | yes | Beguiler Robes (7728, -0.08 DPS) [dungeon]; Green Silk Armor (7065, -0.66 DPS) [crafted]; Mechbuilder's Overalls (9508, -1.10 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 11.0 spell_power points (3.22 DPS) | yes | Glowing Magical Bracelets (13106, -2.06 DPS) [world_drop]; Nightsky Wristbands (6407, -2.08 DPS) [world_drop]; Mindthrust Bracers (1974, -2.13 DPS, sim-verified) [dungeon] |
| hands | Azure Silk Gloves (4319) | Tailoring [crafted] | 10.9 spell_power points (3.19 DPS) | yes | Silky Gloves (253915, -0.71 DPS) [crafted]; Shilly Mitts (9609, -0.84 DPS, sim-verified) [quest]; Gloves of Meditation (4318, -1.10 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.5 spell_power points (3.65 DPS) | yes | Belt of Arugal (6392, -0.13 DPS) [dungeon]; Invoker's Cord (215366, -0.88 DPS) [crafted]; Crimson Silk Belt (7055, -0.88 DPS) [crafted] |
| legs | Silky Leggings (253989) | Tailoring [crafted] | 19.0 spell_power points (5.56 DPS) | yes | Filigreed Silky Leggings (253939, -1.09 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -1.24 DPS) [dungeon]; Azure Silk Pants (7046, -1.77 DPS) [crafted] |
| feet | Frothing Slippers (254003) | Tailoring [crafted] | 17.7 spell_power points (5.17 DPS) | yes | Gilded Slippers (254001, -1.72 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -2.19 DPS) [dungeon]; Spidersilk Boots (4320, -2.55 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 11.6 spell_power points (3.40 DPS) | yes | Runemender's Seal (273031, -1.02 DPS) [world]; Ring of Power Regulation (273030, -1.05 DPS) [dungeon]; Minor Channeling Ring (1449, -1.65 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 8.4 spell_power points (2.44 DPS) | yes | Ring of Power Regulation (273030, -0.09 DPS) [dungeon]; Runemender's Seal (273031, -0.67 DPS, sim-verified) [world]; Minor Channeling Ring (1449, -0.69 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Soulstaff (249393, -1.12 DPS) [crafted]; Tendonscraper Dagger (273053, -1.53 DPS) [world]; Glimmering Staff (249392, -3.61 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.0 spell_power points (2.91 DPS) | yes | Orb of Souls (249395, -1.20 DPS) [crafted]; Eye of Paleth (2943, -1.75 DPS) [quest]; Antipodean Rod (2879, -1.75 DPS, sim-verified) [world_drop] |
| ranged | Thunderwood (13062) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Greater Mystic Wand (217287, -0.17 DPS) [crafted]; Wand of Eventide (5214, -0.38 DPS) [world_drop]; Wand of Mana Concentration (273043, -0.52 DPS) [world] |

**New at 30:** head: Filigreed Silky Circlet; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Azure Silk Cloak; chest: Silky Gown; wrist: Spidertank Oilrag; hands: Azure Silk Gloves; waist: Highlander's Cloth Girdle; legs: Silky Leggings; feet: Frothing Slippers; finger1: Snake Hoop; finger2: Lorekeeper's Ring; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Thunderwood

No-known-source sample (15 of 326, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 000000000000000000-00000000000000000-253510130101301221)

Set DPS (verified): 178.4. Weights run: 2.1s. Verify run: 3.3s. 456 eligible items had no known source.

Race sweep: gnome 178.4 ± 0.2, human 177.2 ± 0.2, high-order-skyborne 176.1 ± 0.3 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.684 ± 0.016, crit=0.164 ± 0.006 per rating point (14 rating = 1%, 2.297 per %), hit=0.343 ± 0.017 per rating point (10 rating = 1%, 3.430 per %), spell_haste=1.030 ± 0.198, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003, spirit=not significant (0.538 ± 0.144), mp5=0.734 ± 0.183

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Miner's Hat of the Deep (9429, +0.00 DPS) [dungeon]; Papal Fez (9431, +0.00 DPS) [dungeon]; Whitemane's Chapeau (7720, -0.18 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Triune Amulet (7722, -0.18 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.89 DPS) [quest]; Necklace of Calisea (1714, -1.01 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Green Silken Shoulders (7057, +0.00 DPS) [crafted]; Mantle of Doan (7712, +0.00 DPS) [dungeon]; Inquisitor's Shawl (19507, +0.00 DPS) [dungeon] |
| back | Azure Silk Cloak (7053) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frost Leather Cloak (7377, -0.16 DPS) [crafted]; Icy Cloak (4327, -0.36 DPS) [crafted]; Expeditionary's Cape (271922, -0.50 DPS) [rep] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 31.7 spell_power points (8.14 DPS) | yes | Robe of the Magi (1716, -0.74 DPS) [world_drop]; Robe of Power (7054, -1.33 DPS) [crafted]; Beguiler Robes (7728, -2.61 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 11.2 spell_power points (2.88 DPS) | yes | Arcane Runed Bracers (4744, -0.57 DPS) [quest]; Condor Bracers (15864, -1.08 DPS) [quest]; Earthen Silk Cuffs (254019, -1.16 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 24.5 spell_power points (6.29 DPS) | yes | Frothing Handwraps (254023, -1.47 DPS) [crafted]; Red Mageweave Gloves (10018, -1.71 DPS) [crafted]; Black Mageweave Gloves (10003, -2.44 DPS) [crafted] |
| waist | Frothing Cord (254039) | Tailoring [crafted] | 22.7 spell_power points (5.83 DPS) | yes | Deathmage Sash (10771, -1.40 DPS) [dungeon]; Highlander's Cloth Girdle (20098, -1.53 DPS) [rep]; Gilded Cord (254037, -1.54 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.2 spell_power points (5.70 DPS) | yes | Silky Leggings (253989, -0.19 DPS) [crafted]; Stoneweaver Leggings (9407, -0.84 DPS) [dungeon]; Filigreed Silky Leggings (253939, -1.27 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 27.8 spell_power points (7.13 DPS) | yes | Frothing Slippers (254003, -2.01 DPS) [crafted]; Gilded Slippers (254001, -3.55 DPS) [crafted]; Furen's Boots (13100, -3.69 DPS) [world_drop] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Snake Hoop (6750, +0.00 DPS) [quest]; Runemender's Seal (273031, -0.66 DPS) [world]; Ring of Power Regulation (273030, -1.34 DPS) [dungeon] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Snake Hoop (6750, +0.00 DPS) [quest]; Runemender's Seal (273031, -0.09 DPS) [world]; Ring of Power Regulation (273030, -0.77 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -7.72 DPS) [dungeon]; Celestial Stave (9517, -8.41 DPS) [quest]; Hypnotic Blade (7714, -8.51 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Icefury Wand (7514) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Freezing Shard (10572, -0.29 DPS) [dungeon]; Plaguerot Sprig (10766, -1.27 DPS) [dungeon]; Jaina's Firestarter (13064, -1.34 DPS) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Frothing Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Staff of Jordan; ranged: Icefury Wand

No-known-source sample (15 of 456, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (gnome, 230005000000000000-00000000000000000-253510130101301221)

Set DPS (verified): 260.2. Weights run: 2.1s. Verify run: 5.4s. 598 eligible items had no known source.

Race sweep: gnome 260.2 ± 0.4, human 259.4 ± 0.4, high-order-skyborne 256.2 ± 0.4 (dps); the guide race gnome is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.372 ± 0.015, crit=0.372 ± 0.012 per rating point (14 rating = 1%, 5.213 per %), hit=0.811 ± 0.033 per rating point (10 rating = 1%, 8.106 per %), spell_haste=4.974 ± 0.206, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002, spirit=not significant (0.167 ± 0.064), mp5=0.455 ± 0.099

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreamweave Circlet (10041) | Tailoring [crafted] | sim-verified (+3.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Red Mageweave Headband (10033, -0.11 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -1.16 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -3.32 DPS, sim-verified) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 18.2 spell_power points (6.95 DPS) | yes | Glowing Eye of Mordresh (10769, -3.15 DPS) [dungeon]; Talvash's Enhancing Necklace (7673, -3.31 DPS, sim-verified) [quest]; Triune Amulet (7722, -3.46 DPS) [dungeon] |
| shoulder | Netherfroth Shoulders (254051) | Tailoring [crafted] | 21.0 spell_power points (8.00 DPS) | yes | Kentic Amice (11624, -0.43 DPS) [dungeon]; Rotgrip Mantle (17732, -0.49 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.68 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frost Leather Cloak (7377, -2.38 DPS) [crafted]; Azure Silk Cloak (7053, -2.57 DPS) [crafted]; Runecloth Cloak (13860, -3.62 DPS, sim-verified) [crafted] |
| chest | Frostweave Robe (13868) | Tailoring [crafted] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Robe of the Magi (1716, -0.29 DPS) [world_drop]; Dreamweave Vest (10021, -0.82 DPS) [crafted]; Frostweave Tunic (13869, -1.08 DPS, sim-verified) [crafted] |
| wrist | Netherfroth Cuffs (254063) | Tailoring [crafted] | 16.4 spell_power points (6.27 DPS) | yes | Nethergeld Cuffs (254061, -2.16 DPS, sim-verified) [crafted]; Spidertank Oilrag (9448, -2.32 DPS) [dungeon]; Aristocratic Cuffs (12546, -2.61 DPS) [dungeon] |
| hands | Frothing Gloves (254097) | Tailoring [crafted] | 24.1 spell_power points (9.17 DPS) | yes | Dreamweave Gloves (10019, -1.30 DPS) [crafted]; Runecloth Gloves (13863, -2.75 DPS) [crafted]; Frostweave Gloves (13870, -3.71 DPS, sim-verified) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 23.8 spell_power points (9.08 DPS) | yes | Frothing Waistcord (254083, -0.56 DPS) [crafted]; Frothing Cord (254039, -2.23 DPS) [crafted]; Satyrmane Sash (17755, -2.33 DPS) [dungeon] |
| legs | Knight's Dreadweave Leggings (220888) | PvP rank 6 · Knight · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshock Leggings (9484, +0.00 DPS) [dungeon]; Wizardweave Leggings (14132, -1.02 DPS) [crafted]; Red Mageweave Pants (10009, -1.22 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.2 spell_power points (9.59 DPS) | yes | Frothing Sandals (254109, -0.69 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -2.18 DPS) [vendor]; Frothing Slippers (254003, -3.39 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 15.6 spell_power points (5.93 DPS) | yes | Freezing Band (942, -0.97 DPS) [world_drop]; Band of the Unicorn (7553, -0.97 DPS) [world_drop]; Cyclopean Band (11824, -1.25 DPS) [dungeon] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 13.8 spell_power points (5.27 DPS) | yes | Band of the Unicorn (7553, -0.31 DPS) [world_drop]; Cyclopean Band (11824, -0.59 DPS) [dungeon]; Freezing Band (942, -2.31 DPS, sim-verified) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Enriched Seal (271903, +0.00 DPS) [rep] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blade of Eternal Darkness (17780, +0.00 DPS) [dungeon]; Kindling Stave (11750, -4.42 DPS) [dungeon]; Arbiter's Blade (11784, -5.24 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Icefury Wand (7514) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lesser Eternal Wand (249232, +0.00 DPS) [crafted]; Freezing Shard (10572, -0.23 DPS) [dungeon]; Noxious Shooter (17745, -0.29 DPS) [dungeon] |

**New at 50:** head: Dreamweave Circlet; neck: Arcane Crystal Pendant; shoulder: Netherfroth Shoulders; back: Spritecaster Cape; chest: Frostweave Robe; wrist: Netherfroth Cuffs; hands: Frothing Gloves; waist: Ban'thok Sash; legs: Knight's Dreadweave Leggings; finger1: Brainlash; finger2: Lorekeeper's Ring; trinket2: Thunderbrew's Boot Flask

No-known-source sample (15 of 598, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (gnome, 230005203000000000-50000000000000000-253510130101301221)

Set DPS (verified): 414.7. Weights run: 2.3s. Verify run: 16.2s. 1442 eligible items had no known source.

Race sweep: gnome 414.7 ± 0.5, human 413.7 ± 0.5, high-order-skyborne 412.8 ± 0.5 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.265 ± 0.020, crit=0.520 ± 0.016 per rating point (14 rating = 1%, 7.285 per %), hit=1.163 ± 0.056 per rating point (10 rating = 1%, 11.633 per %), spell_haste=6.690 ± 0.310, spell_penetration=not significant (0.000 ± 0.000), frost_power=0.808 ± 0.002, spirit=not significant (0.165 ± 0.085), mp5=not significant (0.427 ± 0.131)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 34.3 spell_power points (13.54 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -0.10 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -4.26 DPS, sim-verified) [crafted] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 22.8 spell_power points (8.98 DPS) | yes | Star of Mystaria (12103, -1.82 DPS) [dungeon]; Arcane Crystal Pendant (20037, -2.04 DPS) [quest]; Lady Maye's Pendant (14558, -2.65 DPS) [world_drop] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 27.7 spell_power points (10.95 DPS) | yes | Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Argent Shoulders (19059, -0.62 DPS) [crafted]; Lieutenant Commander's Silk Mantle (227102, -0.74 DPS) [pvp] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -4.05 DPS, sim-verified) [world]; Spritecaster Cape (11623, -5.20 DPS) [dungeon]; Shroud of the Exile (15421, -5.51 DPS) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Vestments (231603, -1.85 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -6.65 DPS) [pvp]; Robe of Winter Night (14136, -8.93 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 25.3 spell_power points (9.98 DPS) | yes | Runecloth Cuffs (254123, -4.59 DPS) [crafted]; Netherfroth Cuffs (254063, -4.77 DPS) [crafted]; Manacle Cuffs (11962, -5.65 DPS) [quest] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 26.7 spell_power points (10.54 DPS) | yes | Gloves of Spell Mastery (14146, +0.00 DPS) [crafted]; Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 26.3 spell_power points (10.38 DPS) | yes | Magician's Cord (272393, -0.37 DPS) [vendor]; Highlander's Cloth Girdle (20047, -1.35 DPS) [rep]; Stormpike Cloth Girdle (19094, -2.22 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 49.2 spell_power points (19.43 DPS) | yes | Marshal's Silk Leggings (231605, -1.97 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -6.16 DPS) [pvp]; Skyshroud Leggings (13170, -6.33 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.2 spell_power points (9.93 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Knight-Lieutenant's Silk Walkers (227112, -0.59 DPS) [pvp]; Venomspew Footpads (275606, -0.72 DPS) [crafted] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -1.22 DPS) [quest]; Don Julio's Band (19325, -1.59 DPS) [rep]; Naglering (11669, -11.10 DPS, sim-verified) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.39 DPS) [quest]; Don Julio's Band (19325, -0.76 DPS) [rep]; Don Mauricio's Band of Domination (22433, -1.01 DPS) [dungeon] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+4.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, -2.51 DPS) [dungeon]; Serenity Field (272439, -4.34 DPS) [vendor]; Draconic Infused Emblem (22268, -4.59 DPS, sim-verified) [dungeon] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Uther's Strength (11302, -1.76 DPS) [world_drop]; Cannonball Runner (13382, -1.84 DPS) [dungeon] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironbark Staff (20069, -0.27 DPS) [rep]; Grand Marshal's Stave (234571, -2.20 DPS) [pvp]; Blade of Eternal Darkness (17780, -10.71 DPS, sim-verified) [dungeon] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (11.45 DPS) | yes | Grand Marshal's Tome of Power (234589, -0.22 DPS) [pvp]; Tome of the Ice Lord (19310, -0.73 DPS) [rep]; Spirit of Aquementas (11904, -3.16 DPS) [quest] |
| ranged | Wand of Biting Cold (19108) | Korrak the Bloodrager [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, -0.37 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -0.76 DPS) [dungeon]; Brilliant Wand (249385, -0.90 DPS) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Hands of Power; legs: Sentinel's Silk Leggings; finger1: Blessed Band of Light; finger2: Maiden's Circle; trinket1: Weakness Analyzer; trinket2: Frozen Heart of the Mountain; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Wand of Biting Cold

No-known-source sample (15 of 1442, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (gnome, 230005203000000000-50000000000000000-253510130101301221)

Set DPS (verified): 685.3. Weights run: 2.4s. Verify run: 17.0s. 1442 eligible items had no known source.

Race sweep: high-order-skyborne 689.5 ± 0.7, gnome 685.3 ± 0.7, human 683.4 ± 0.7 (dps); high-order-skyborne beats the guide race gnome by 4.3 (0.62%) ± 1.0, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.207 ± 0.011, crit=0.470 ± 0.014 per rating point (14 rating = 1%, 6.581 per %), hit=1.046 ± 0.049 per rating point (10 rating = 1%, 10.464 per %), spell_haste=6.307 ± 0.230, spell_penetration=not significant (0.000 ± 0.000), frost_power=0.810 ± 0.002, spirit=not significant (0.045 ± 0.023), mp5=not significant (0.066 ± 0.079)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 32.5 spell_power points (14.84 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -0.43 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -2.52 DPS) [crafted] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 21.1 spell_power points (9.64 DPS) | yes | Arcane Crystal Pendant (20037, -1.77 DPS) [quest]; Star of Mystaria (12103, -2.62 DPS) [dungeon]; Lady Maye's Pendant (14558, -2.82 DPS) [world_drop] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 26.3 spell_power points (11.99 DPS) | yes | Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Argent Shoulders (19059, -0.44 DPS) [crafted]; Lieutenant Commander's Silk Mantle (227102, -1.02 DPS) [pvp] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -3.66 DPS, sim-verified) [world]; Spritecaster Cape (11623, -5.75 DPS) [dungeon]; Shroud of the Exile (15421, -6.27 DPS) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Vestments (231603, -2.60 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -8.09 DPS) [pvp]; Robe of Winter Night (14136, -8.53 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 24.0 spell_power points (10.94 DPS) | yes | Runecloth Cuffs (254123, -4.98 DPS) [crafted]; Netherfroth Cuffs (254063, -5.37 DPS) [crafted]; Manacle Cuffs (11962, -6.33 DPS) [quest] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Hands of Power (13253, -2.40 DPS, sim-verified) [dungeon] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 24.5 spell_power points (11.20 DPS) | yes | Magician's Cord (272393, -0.26 DPS) [vendor]; Highlander's Cloth Girdle (20047, -1.24 DPS) [rep]; Stormpike Cloth Girdle (19094, -2.04 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 47.2 spell_power points (21.56 DPS) | yes | Marshal's Silk Leggings (231605, -2.76 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -7.26 DPS) [pvp]; Skyshroud Leggings (13170, -11.10 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.3 spell_power points (11.10 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Knight-Lieutenant's Silk Walkers (227112, -0.57 DPS) [pvp]; Omnicast Boots (11822, -0.83 DPS) [dungeon] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.63 DPS) [quest]; Don Mauricio's Band of Domination (22433, -1.48 DPS) [dungeon]; Naglering (11669, -10.58 DPS, sim-verified) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.12 DPS) [quest]; Don Mauricio's Band of Domination (22433, -0.98 DPS) [dungeon]; Don Julio's Band (19325, -1.22 DPS) [rep] |
| trinket1 | Eye of the Beast (13968) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Draconic Infused Emblem (22268, -3.94 DPS, sim-verified) [dungeon] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Draconic Infused Emblem (22268, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -2.44 DPS) [pvp]; Elemental Mage Staff (944, -8.18 DPS) [world_drop]; Headmaster's Charge (13937, -10.92 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Biting Cold (19108) | Korrak the Bloodrager [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, -0.51 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -0.89 DPS) [dungeon]; Brilliant Wand (249385, -1.14 DPS) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; legs: Sentinel's Silk Leggings; finger1: Blessed Band of Light; finger2: Maiden's Circle; trinket1: Eye of the Beast; trinket2: Frozen Heart of the Mountain; main_hand: Ironbark Staff; ranged: Wand of Biting Cold

No-known-source sample (15 of 1442, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (undead, 000000000000000000-00000000000000000-053200010000000000)

Set DPS (verified): 49.2. Weights run: 2.0s. Verify run: 2.2s. 192 eligible items had no known source.

Race sweep: undead 49.2 ± 0.1, orc 48.0 ± 0.1, troll 47.4 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.977 ± 0.007, crit=0.105 ± 0.003 per rating point (14 rating = 1%, 1.470 per %), hit=0.255 ± 0.008 per rating point (10 rating = 1%, 2.550 per %), spell_haste=1.090 ± 0.098, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002, spirit=not significant (0.051 ± 0.036), mp5=0.856 ± 0.080

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Silky Circlet (253951) | Tailoring [crafted] | 11.0 spell_power points (2.46 DPS) | yes | Pristine Circlet (253949, -1.10 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.30 DPS) [crafted]; Flying Tiger Goggles (4368, -2.42 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.2 spell_power points (0.03 DPS) | yes | Scholarly Pendant (277203, -0.01 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.8 spell_power points (3.09 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.96 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -2.19 DPS) [crafted]; Slime-encrusted Pads (6461, -2.51 DPS) [dungeon] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.16 DPS) [crafted]; Sanguine Cape (14376, -0.23 DPS) [world_drop]; Battle Healer's Cloak (20427, -0.41 DPS) [rep] |
| chest | Filigreed Silky Gown (253903) | Tailoring [crafted] | 14.0 spell_power points (3.14 DPS) | yes | Filigreed Pristine Gown (253901, -0.88 DPS, sim-verified) [crafted]; Gray Woolen Robe (2585, -1.15 DPS) [crafted]; Mystic's Wrap (14369, -1.61 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 7.5 spell_power points (1.67 DPS) | yes | Bright Bracers (3647, -0.76 DPS, sim-verified) [world_drop]; Tabitha's Cuffs (251486, -1.00 DPS) [quest]; Owlbeard Bracers (16981, -1.01 DPS) [quest] |
| hands | Silky Gloves (253915) | Tailoring [crafted] | 9.9 spell_power points (2.22 DPS) | yes | Pristine Gloves (253913, -0.67 DPS) [crafted]; Serpent Gloves (5970, -0.77 DPS, sim-verified) [dungeon]; Gnoll Casting Gloves (892, -0.88 DPS) [world] |
| waist | Silky Sash (253927) | Tailoring [crafted] | 10.9 spell_power points (2.44 DPS) | yes | Pristine Sash (253925, -0.66 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.69 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.89 DPS) [crafted] |
| legs | Filigreed Silky Leggings (253939) | Tailoring [crafted] | 17.1 spell_power points (3.82 DPS) | yes | Filigreed Pristine Leggings (253937, -1.10 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -2.22 DPS) [world_drop]; Silk-threaded Trousers (1929, -2.25 DPS) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.9 spell_power points (2.44 DPS) | yes | Silky Boots (253891, -0.34 DPS, sim-verified) [crafted]; Pristine Boots (253889, -1.11 DPS) [crafted]; Red Woolen Boots (4313, -1.49 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 6.7 spell_power points (1.50 DPS) | yes | Loop of Sacrifice (281673, -0.41 DPS) [quest]; Volcanic Rock Ring (12053, -0.85 DPS) [world_drop]; Furlfeather Ring (286745, -1.28 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.9 spell_power points (1.31 DPS) | yes | Loop of Sacrifice (281673, -0.22 DPS) [quest]; Volcanic Rock Ring (12053, -0.66 DPS) [world_drop]; Furlfeather Ring (286745, -1.09 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Staff of the Blessed Seer (2271, -0.38 DPS) [dungeon]; Crescent Staff (6505, -0.65 DPS, sim-verified) [quest]; Gnarled Necromancer's Staff (251534, -0.83 DPS) [quest] |
| off_hand | Antipodean Rod (2879) | World drop [world_drop] | 6.0 spell_power points (1.34 DPS) | yes | Seer's Fine Stein (7608, -0.68 DPS) [world_drop]; Spellbinder Orb (15926, -0.69 DPS) [dungeon]; Tork Wrench (11855, -0.91 DPS) [quest] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller (12984, -0.29 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -0.58 DPS) [quest]; Greater Magic Wand (11288, -0.61 DPS) [crafted] |

**New at 20:** head: Silky Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Silky Gown; wrist: Mindthrust Bracers; hands: Silky Gloves; waist: Silky Sash; legs: Filigreed Silky Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Evocator's Blade; off_hand: Antipodean Rod; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 192, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 8350 The 1 Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 23173 Abomination Skin Leggings

### Band 30 (undead, 000000000000000000-00000000000000000-053510030100200100)

Set DPS (verified): 88.0. Weights run: 2.0s. Verify run: 1.8s. 318 eligible items had no known source.

Race sweep: undead 88.0 ± 0.1, troll 87.2 ± 0.1, orc 86.4 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.322 ± 0.007, crit=0.107 ± 0.004 per rating point (14 rating = 1%, 1.499 per %), hit=0.216 ± 0.009 per rating point (10 rating = 1%, 2.163 per %), spell_haste=0.963 ± 0.089, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003, spirit=not significant (0.093 ± 0.039), mp5=0.311 ± 0.065

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Silky Circlet (253977) | Tailoring [crafted] | 14.0 spell_power points (3.04 DPS) | yes | Silky Circlet (253951, -0.65 DPS) [crafted]; Azure Silk Hood (7048, -0.87 DPS) [crafted]; Holy Shroud (2721, -1.07 DPS, sim-verified) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 1.9 spell_power points (0.42 DPS) | yes | Kaleidoscope Chain (13084, -0.06 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.07 DPS) [vendor]; Crystal Starfire Medallion (5003, -0.08 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.9 spell_power points (2.58 DPS) | yes | Death Speaker Mantle (6685, -0.51 DPS) [dungeon]; Invoker's Mantle (215365, -0.71 DPS) [crafted]; Magician's Mantle (12998, -0.87 DPS) [world_drop] |
| back | Azure Silk Cloak (7053) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, -1.15 DPS) [crafted]; Heavy Woolen Cloak (4311, -1.28 DPS) [crafted]; Windsong Drape (15468, -1.36 DPS, sim-verified) [quest] |
| chest | Silky Gown (253963) | Tailoring [crafted] | 15.6 spell_power points (3.39 DPS) | yes | Beguiler Robes (7728, -0.44 DPS) [dungeon]; Green Silk Armor (7065, -0.53 DPS) [crafted]; Tree Bark Jacket (1486, -0.57 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.9 spell_power points (2.16 DPS) | yes | Glowing Magical Bracelets (13106, -1.60 DPS) [world_drop]; Mindthrust Bracers (1974, -1.60 DPS) [dungeon]; Tabitha's Cuffs (251486, -2.35 DPS, sim-verified) [quest] |
| hands | Azure Silk Gloves (4319) | Tailoring [crafted] | 10.3 spell_power points (2.23 DPS) | yes | Silky Gloves (253915, -0.50 DPS) [crafted]; Jutebraid Gloves (10654, -0.58 DPS) [quest]; Serpent Gloves (5970, -0.71 DPS) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.0 spell_power points (2.60 DPS) | yes | Belt of Arugal (6392, -0.33 DPS) [dungeon]; Invoker's Cord (215366, -0.73 DPS) [crafted]; Warsong Sash (16975, -0.78 DPS, sim-verified) [quest] |
| legs | Silky Leggings (253989) | Tailoring [crafted] | 16.7 spell_power points (3.63 DPS) | yes | Gaze Dreamer Pants (6903, -0.84 DPS) [dungeon]; Azure Silk Pants (7046, -1.04 DPS) [crafted]; Filigreed Silky Leggings (253939, -1.10 DPS, sim-verified) [crafted] |
| feet | Frothing Slippers (254003) | Tailoring [crafted] | 15.6 spell_power points (3.39 DPS) | yes | Spidersilk Boots (4320, -1.59 DPS) [crafted]; Gilded Slippers (254001, -1.63 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.67 DPS) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.9 spell_power points (1.93 DPS) | yes | Ring of Power Regulation (273030, -0.43 DPS) [dungeon]; Runemender's Seal (273031, -0.72 DPS) [world]; Solid Gold Band (279845, -0.96 DPS) [quest] |
| finger2 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.6 spell_power points (1.65 DPS) | yes | Ring of Power Regulation (273030, -0.15 DPS) [dungeon]; Runemender's Seal (273031, -0.44 DPS) [world]; Solid Gold Band (279845, -0.69 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -0.95 DPS) [dungeon]; Tendonscraper Dagger (273053, -0.95 DPS) [world]; Glimmering Staff (249392, -2.51 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.9 spell_power points (1.94 DPS) | yes | Orb of Souls (249395, -0.95 DPS) [crafted]; Witch's Finger (16887, -1.38 DPS) [quest]; Antipodean Rod (2879, -1.88 DPS, sim-verified) [world_drop] |
| ranged | Thunderwood (13062) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Greater Mystic Wand (217287, -0.04 DPS) [crafted]; Wand of Eventide (5214, -0.34 DPS) [world_drop]; Starfaller (13063, -0.50 DPS) [world_drop] |

**New at 30:** head: Filigreed Silky Circlet; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Azure Silk Cloak; chest: Silky Gown; wrist: Spidertank Oilrag; hands: Azure Silk Gloves; waist: Defiler's Cloth Girdle; legs: Silky Leggings; feet: Frothing Slippers; finger1: Snake Hoop; finger2: Advisor's Ring; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Thunderwood

No-known-source sample (15 of 318, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 8350 The 1 Ring; 9362 Brilliant Gold Ring

### Band 40 (undead, 000000000000000000-00000000000000000-253510130101301221)

Set DPS (verified): 176.7. Weights run: 2.1s. Verify run: 3.2s. 448 eligible items had no known source.

Race sweep: undead 176.7 ± 0.2, orc 175.4 ± 0.3, troll 174.3 ± 0.3 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.684 ± 0.016, crit=0.164 ± 0.006 per rating point (14 rating = 1%, 2.297 per %), hit=0.343 ± 0.017 per rating point (10 rating = 1%, 3.430 per %), spell_haste=1.030 ± 0.198, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003, spirit=not significant (0.538 ± 0.144), mp5=0.734 ± 0.183

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Miner's Hat of the Deep (9429, +0.00 DPS) [dungeon]; Papal Fez (9431, +0.00 DPS) [dungeon]; Whitemane's Chapeau (7720, -0.18 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.5 spell_power points (3.20 DPS) | yes | Triune Amulet (7722, -0.18 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.89 DPS) [quest]; Necklace of Calisea (1714, -1.01 DPS) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.9 spell_power points (4.08 DPS) | yes | Green Silken Shoulders (7057, -0.09 DPS) [crafted]; Mantle of Doan (7712, -0.12 DPS) [dungeon]; Bloodmage Mantle (7684, -0.19 DPS) [dungeon] |
| back | Frost Leather Cloak (7377) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Azure Silk Cloak (7053, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.20 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.34 DPS) [vendor] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Robe of the Magi (1716, -0.74 DPS) [world_drop]; Robe of Power (7054, -1.33 DPS) [crafted]; Beguiler Robes (7728, -2.61 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 11.2 spell_power points (2.88 DPS) | yes | Radiant Silver Bracers (4545, -0.44 DPS) [quest]; Condor Bracers (15864, -1.08 DPS) [quest]; Earthen Silk Cuffs (254019, -1.16 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 24.5 spell_power points (6.29 DPS) | yes | Frothing Handwraps (254023, -1.47 DPS) [crafted]; Red Mageweave Gloves (10018, -1.71 DPS) [crafted]; Black Mageweave Gloves (10003, -2.44 DPS) [crafted] |
| waist | Frothing Cord (254039) | Tailoring [crafted] | 22.7 spell_power points (5.83 DPS) | yes | Deathmage Sash (10771, -1.40 DPS) [dungeon]; Defiler's Cloth Girdle (20166, -1.53 DPS) [rep]; Gilded Cord (254037, -1.54 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.2 spell_power points (5.70 DPS) | yes | Silky Leggings (253989, -0.19 DPS) [crafted]; Stoneweaver Leggings (9407, -0.84 DPS) [dungeon]; Filigreed Silky Leggings (253939, -1.27 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 27.8 spell_power points (7.13 DPS) | yes | Frothing Slippers (254003, -2.01 DPS) [crafted]; Gilded Slippers (254001, -3.55 DPS) [crafted]; Furen's Boots (13100, -3.69 DPS) [world_drop] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Snake Hoop (6750, +0.00 DPS) [quest]; Runemender's Seal (273031, -0.66 DPS) [world]; Ring of Power Regulation (273030, -1.34 DPS) [dungeon] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Snake Hoop (6750, +0.00 DPS) [quest]; Runemender's Seal (273031, -0.09 DPS) [world]; Ring of Power Regulation (273030, -0.77 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -7.72 DPS) [dungeon]; Celestial Stave (9517, -8.41 DPS) [quest]; Hypnotic Blade (7714, -8.51 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Icefury Wand (7514) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Freezing Shard (10572, -0.30 DPS) [dungeon]; Dancing Flame (6806, -1.25 DPS) [quest]; Plaguerot Sprig (10766, -1.25 DPS) [dungeon] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Frost Leather Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Frothing Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Staff of Jordan; ranged: Icefury Wand

No-known-source sample (15 of 448, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (undead, 230005000000000000-00000000000000000-253510130101301221)

Set DPS (verified): 258.9. Weights run: 2.1s. Verify run: 5.7s. 590 eligible items had no known source.

Race sweep: undead 258.9 ± 0.4, orc 256.8 ± 0.4, troll 255.7 ± 0.4 (dps); the guide race undead is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.372 ± 0.015, crit=0.372 ± 0.012 per rating point (14 rating = 1%, 5.213 per %), hit=0.811 ± 0.033 per rating point (10 rating = 1%, 8.106 per %), spell_haste=4.974 ± 0.206, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002, spirit=not significant (0.167 ± 0.064), mp5=0.455 ± 0.099

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreamweave Circlet (10041) | Tailoring [crafted] | sim-verified (+4.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Red Mageweave Headband (10033, -0.11 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -1.16 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -4.03 DPS, sim-verified) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 18.2 spell_power points (6.95 DPS) | yes | Jarkal's Enhancing Necklace (7888, -3.07 DPS, sim-verified) [quest]; Glowing Eye of Mordresh (10769, -3.15 DPS) [dungeon]; Triune Amulet (7722, -3.46 DPS) [dungeon] |
| shoulder | Netherfroth Shoulders (254051) | Tailoring [crafted] | 21.0 spell_power points (8.00 DPS) | yes | Kentic Amice (11624, -0.43 DPS) [dungeon]; Rotgrip Mantle (17732, -0.49 DPS) [dungeon]; Blood Guard's Dreadweave Mantle (220905, -1.68 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Runecloth Cloak (13860, -2.00 DPS) [crafted]; Frost Leather Cloak (7377, -2.38 DPS) [crafted]; Deep Woodlands Cloak (19121, -3.28 DPS, sim-verified) [quest] |
| chest | Frostweave Robe (13868) | Tailoring [crafted] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Robe of the Magi (1716, -0.29 DPS) [world_drop]; Dreamweave Vest (10021, -0.82 DPS) [crafted]; Frostweave Tunic (13869, -1.41 DPS, sim-verified) [crafted] |
| wrist | Netherfroth Cuffs (254063) | Tailoring [crafted] | 16.4 spell_power points (6.27 DPS) | yes | Nethergeld Cuffs (254061, -2.29 DPS) [crafted]; Spidertank Oilrag (9448, -2.32 DPS) [dungeon]; Aristocratic Cuffs (12546, -2.61 DPS) [dungeon] |
| hands | Frothing Gloves (254097) | Tailoring [crafted] | 24.1 spell_power points (9.17 DPS) | yes | Dreamweave Gloves (10019, -1.30 DPS) [crafted]; Runecloth Gloves (13863, -2.75 DPS) [crafted]; Frostweave Gloves (13870, -4.61 DPS, sim-verified) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 23.8 spell_power points (9.08 DPS) | yes | Frothing Cord (254039, -2.23 DPS) [crafted]; Satyrmane Sash (17755, -2.33 DPS) [dungeon]; Frothing Waistcord (254083, -2.45 DPS, sim-verified) [crafted] |
| legs | Stone Guard's Dreadweave Leggings (220906) | PvP rank 6 · Stone Guard · Horde [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshock Leggings (9484, +0.00 DPS) [dungeon]; Wizardweave Leggings (14132, -1.02 DPS) [crafted]; Red Mageweave Pants (10009, -1.22 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.2 spell_power points (9.59 DPS) | yes | Frothing Sandals (254109, -0.69 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -2.18 DPS) [vendor]; Frothing Slippers (254003, -3.39 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 15.6 spell_power points (5.93 DPS) | yes | Freezing Band (942, -0.97 DPS) [world_drop]; Band of the Unicorn (7553, -0.97 DPS) [world_drop]; Cyclopean Band (11824, -1.25 DPS) [dungeon] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 13.8 spell_power points (5.27 DPS) | yes | Band of the Unicorn (7553, -0.31 DPS) [world_drop]; Cyclopean Band (11824, -0.59 DPS) [dungeon]; Freezing Band (942, -2.46 DPS, sim-verified) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Darkspear Voodoo Seal (272061, -1.41 DPS, sim-verified) [vendor] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, +0.00 DPS) [dungeon]; Arbiter's Blade (11784, +0.00 DPS) [dungeon]; Staff of Jordan (873, -3.73 DPS, sim-verified) [world_drop] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 14.1 spell_power points (5.38 DPS) | yes | Orb of the Forgotten Seer (7685, -0.04 DPS) [dungeon]; Truesilver Conduit (249455, -0.43 DPS) [crafted]; Flask of Forest Mojo (19115, -0.82 DPS) [quest] |
| ranged | Icefury Wand (7514) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lesser Eternal Wand (249232, +0.00 DPS) [crafted]; Noxious Shooter (17745, -0.04 DPS) [dungeon]; Freezing Shard (10572, -0.26 DPS) [dungeon] |

**New at 50:** head: Dreamweave Circlet; neck: Arcane Crystal Pendant; shoulder: Netherfroth Shoulders; back: Spritecaster Cape; chest: Frostweave Robe; wrist: Netherfroth Cuffs; hands: Frothing Gloves; waist: Ban'thok Sash; legs: Stone Guard's Dreadweave Leggings; finger1: Brainlash; finger2: Advisor's Ring; trinket2: Evonice's Landin' Pilla; main_hand: Blade of Eternal Darkness; off_hand: Celestial Orb

No-known-source sample (15 of 590, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (undead, 230005203000000000-50000000000000000-253510130101301221)

Set DPS (verified): 410.6. Weights run: 2.3s. Verify run: 16.2s. 1436 eligible items had no known source.

Race sweep: undead 410.6 ± 0.5, orc 406.9 ± 0.5, troll 406.9 ± 0.5 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.265 ± 0.020, crit=0.520 ± 0.016 per rating point (14 rating = 1%, 7.285 per %), hit=1.163 ± 0.056 per rating point (10 rating = 1%, 11.633 per %), spell_haste=6.690 ± 0.310, spell_penetration=not significant (0.000 ± 0.000), frost_power=0.808 ± 0.002, spirit=not significant (0.165 ± 0.085), mp5=not significant (0.427 ± 0.131)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 34.3 spell_power points (13.54 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -0.10 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -4.00 DPS, sim-verified) [crafted] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 22.8 spell_power points (8.98 DPS) | yes | Arcane Crystal Pendant (20037, -2.04 DPS) [quest]; Star of Mystaria (12103, -2.14 DPS, sim-verified) [dungeon]; Lady Maye's Pendant (14558, -2.65 DPS) [world_drop] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 27.7 spell_power points (10.95 DPS) | yes | Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Argent Shoulders (19059, -0.62 DPS) [crafted]; Champion's Silk Mantle (227104, -0.74 DPS) [pvp] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -3.80 DPS, sim-verified) [world]; Spritecaster Cape (11623, -5.20 DPS) [dungeon]; Shroud of the Exile (15421, -5.51 DPS) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Raiment (231596, -1.85 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -6.65 DPS) [pvp]; Robe of Winter Night (14136, -9.44 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 25.3 spell_power points (9.98 DPS) | yes | Runecloth Cuffs (254123, -4.59 DPS) [crafted]; Netherfroth Cuffs (254063, -4.77 DPS) [crafted]; Manacle Cuffs (11962, -5.65 DPS) [quest] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 26.7 spell_power points (10.54 DPS) | yes | Gloves of Spell Mastery (14146, +0.00 DPS) [crafted]; General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 26.3 spell_power points (10.38 DPS) | yes | Magician's Cord (272393, -0.37 DPS) [vendor]; Defiler's Cloth Girdle (20163, -1.35 DPS) [rep]; Frostwolf Cloth Belt (19090, -2.22 DPS) [rep] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 34.7 spell_power points (13.69 DPS) | yes | General's Silk Trousers (231595, +0.00 DPS) [pvp]; Legionnaire's Silk Legguards (227107, -0.42 DPS) [pvp]; Skyshroud Leggings (13170, -0.59 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.2 spell_power points (9.93 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Blood Guard's Silk Walkers (227110, -0.59 DPS) [pvp]; Venomspew Footpads (275606, -0.72 DPS) [crafted] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -1.22 DPS) [quest]; Don Julio's Band (19325, -1.59 DPS) [rep]; Naglering (11669, -10.31 DPS, sim-verified) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.39 DPS) [quest]; Don Julio's Band (19325, -0.76 DPS) [rep]; Don Mauricio's Band of Domination (22433, -1.01 DPS) [dungeon] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+4.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, -2.51 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -2.97 DPS) [crafted]; Burst of Knowledge (11832, -4.32 DPS, sim-verified) [dungeon] |
| trinket2 | Eye of the Beast (13968) | For The Horde! [quest] | sim-verified (-1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironbark Staff (20220, -0.27 DPS) [rep]; High Warlord's War Staff (234549, -2.20 DPS) [pvp]; Blade of Eternal Darkness (17780, -9.39 DPS, sim-verified) [dungeon] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (11.45 DPS) | yes | High Warlord's Tome of Destruction (234563, -0.22 DPS) [pvp]; Tome of the Ice Lord (19310, -2.43 DPS, sim-verified) [rep]; Spirit of Aquementas (11904, -3.16 DPS) [quest] |
| ranged | Wand of Biting Cold (19108) | The Legend of Korrak [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, -0.37 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -0.76 DPS) [dungeon]; Brilliant Wand (249385, -0.89 DPS) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Hands of Power; legs: Outrider's Silk Leggings; finger1: Blessed Band of Light; finger2: Maiden's Circle; trinket1: Weakness Analyzer; trinket2: Eye of the Beast; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Wand of Biting Cold

No-known-source sample (15 of 1436, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60, raid preset (undead, 230005203000000000-50000000000000000-253510130101301221)

Set DPS (verified): 679.8. Weights run: 2.4s. Verify run: 17.3s. 1436 eligible items had no known source.

Race sweep: undead 679.8 ± 0.7, troll 677.2 ± 0.7, orc 676.3 ± 0.7 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.207 ± 0.011, crit=0.470 ± 0.014 per rating point (14 rating = 1%, 6.581 per %), hit=1.046 ± 0.049 per rating point (10 rating = 1%, 10.464 per %), spell_haste=6.307 ± 0.230, spell_penetration=not significant (0.000 ± 0.000), frost_power=0.810 ± 0.002, spirit=not significant (0.045 ± 0.023), mp5=not significant (0.066 ± 0.079)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 32.5 spell_power points (14.84 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -0.43 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -2.52 DPS) [crafted] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 21.1 spell_power points (9.64 DPS) | yes | Arcane Crystal Pendant (20037, -1.77 DPS) [quest]; Star of Mystaria (12103, -2.62 DPS) [dungeon]; Lady Maye's Pendant (14558, -2.82 DPS) [world_drop] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 26.3 spell_power points (11.99 DPS) | yes | Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Argent Shoulders (19059, -0.44 DPS) [crafted]; Champion's Silk Mantle (227104, -1.02 DPS) [pvp] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -4.07 DPS, sim-verified) [world]; Spritecaster Cape (11623, -5.75 DPS) [dungeon]; Shroud of the Exile (15421, -6.27 DPS) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Raiment (231596, -2.60 DPS) [pvp]; Robe of Winter Night (14136, -4.36 DPS, sim-verified) [crafted]; Legionnaire's Silk Tunic (227106, -8.09 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 24.0 spell_power points (10.94 DPS) | yes | Runecloth Cuffs (254123, -4.98 DPS) [crafted]; Netherfroth Cuffs (254063, -5.37 DPS) [crafted]; Manacle Cuffs (11962, -6.33 DPS) [quest] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Hands of Power (13253, -2.44 DPS, sim-verified) [dungeon] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 24.5 spell_power points (11.20 DPS) | yes | Magician's Cord (272393, -0.26 DPS) [vendor]; Defiler's Cloth Girdle (20163, -1.24 DPS) [rep]; Frostwolf Cloth Belt (19090, -2.04 DPS) [rep] |
| legs | Skyshroud Leggings (13170) | Blackrock Spire: Highlord Omokk [dungeon] | 32.5 spell_power points (14.83 DPS) | yes | General's Silk Trousers (231595, +0.00 DPS) [pvp]; Outrider's Silk Leggings (22747, -0.04 DPS) [rep]; Legionnaire's Silk Legguards (227107, -0.53 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.3 spell_power points (11.10 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Blood Guard's Silk Walkers (227110, -0.57 DPS) [pvp]; Omnicast Boots (11822, -0.83 DPS) [dungeon] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.63 DPS) [quest]; Don Mauricio's Band of Domination (22433, -1.48 DPS) [dungeon]; Naglering (11669, -10.66 DPS, sim-verified) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.12 DPS) [quest]; Don Mauricio's Band of Domination (22433, -0.98 DPS) [dungeon]; Don Julio's Band (19325, -1.22 DPS) [rep] |
| trinket1 | Eye of the Beast (13968) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -2.40 DPS, sim-verified) [crafted]; Smolderweb's Eye (13213, -3.44 DPS) [dungeon]; Rune of the Guard Captain (19120, -4.87 DPS) [quest] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -2.44 DPS) [pvp]; Headmaster's Charge (13937, -8.13 DPS, sim-verified) [dungeon]; Elemental Mage Staff (944, -8.18 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Biting Cold (19108) | The Legend of Korrak [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, -0.51 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -0.89 DPS) [dungeon]; Brilliant Wand (249385, -1.14 DPS) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; legs: Skyshroud Leggings; finger1: Blessed Band of Light; finger2: Maiden's Circle; trinket1: Eye of the Beast; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Wand of Biting Cold

No-known-source sample (15 of 1436, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

