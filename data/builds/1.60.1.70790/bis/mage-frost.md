# Leveling BiS: Frost

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-053200010000000000)

Set DPS (verified): 50.6. Weights run: 2.2s. Verify run: 1.9s. 198 eligible items had no known source.

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

Set DPS (verified): 89.6. Weights run: 2.2s. Verify run: 1.8s. 326 eligible items had no known source.

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

Set DPS (verified): 173.9. Weights run: 2.4s. Verify run: 4.8s. 456 eligible items had no known source.

Race sweep: gnome 173.9 ± 0.3, human 172.0 ± 0.3, high-order-skyborne 170.8 ± 0.3 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.239 ± 0.009, crit=0.266 ± 0.008 per rating point (14 rating = 1%, 3.722 per %), hit=0.542 ± 0.022 per rating point (10 rating = 1%, 5.421 per %), spell_haste=4.344 ± 0.144, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002, spirit=not significant (-0.014 ± 0.073), mp5=not significant (0.058 ± 0.084)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (7.97 DPS) | yes | Miner's Hat of the Deep (9429, -1.88 DPS) [dungeon]; Filigreed Silky Circlet (253977, -2.66 DPS) [crafted]; Augural Shroud (2620, -2.89 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Triune Amulet (7722, -0.35 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.88 DPS, sim-verified) [quest]; Seal of the Expedition (271919, -2.75 DPS) [rep] |
| shoulder | Azure Shoulders (7060) | Tailoring [crafted] | 13.0 spell_power points (4.94 DPS) | yes | Bloodmage Mantle (7684, -0.70 DPS) [dungeon]; Green Silken Shoulders (7057, -0.90 DPS) [crafted]; Inquisitor's Shawl (19507, -1.10 DPS) [dungeon] |
| back | Azure Silk Cloak (7053) | Tailoring [crafted] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Frost Leather Cloak (7377, -0.58 DPS, sim-verified) [crafted]; Long Silken Cloak (4326, -1.07 DPS) [crafted]; Guardian Cloak (5965, -1.07 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.4 spell_power points (8.90 DPS) | yes | Dreamweave Vest (10021, -1.25 DPS) [crafted]; Robe of Power (7054, -2.49 DPS) [crafted]; Elemental Raiment (9434, -3.26 DPS, sim-verified) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.2 spell_power points (3.48 DPS) | yes | Condor Bracers (15864, -0.83 DPS) [quest]; Arcane Runed Bracers (4744, -1.83 DPS, sim-verified) [quest]; Earthen Silk Cuffs (254019, -1.97 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.0 spell_power points (7.20 DPS) | yes | Black Mageweave Gloves (10003, -1.50 DPS) [crafted]; Red Mageweave Gloves (10018, -2.11 DPS) [crafted]; Frothing Handwraps (254023, -2.94 DPS, sim-verified) [crafted] |
| waist | Frothing Cord (254039) | Tailoring [crafted] | 15.9 spell_power points (6.04 DPS) | yes | Star Belt (4329, -1.11 DPS) [crafted]; Highlander's Cloth Girdle (20099, -1.59 DPS) [rep] |
| legs | Silky Leggings (253989) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Red Mageweave Pants (10009, -0.50 DPS, sim-verified) [crafted]; Filigreed Silky Leggings (253939, -1.23 DPS) [crafted]; Gaze Dreamer Pants (6903, -1.39 DPS) [dungeon] |
| feet | Frothing Slippers (254003) | Tailoring [crafted] | 14.7 spell_power points (5.57 DPS) | yes | Gilded Slippers (254001, -2.13 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -2.55 DPS) [crafted]; Glinteye Slippers (273024, -2.91 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.2 spell_power points (4.25 DPS) | yes | Ring of Forlorn Spirits (2043, -1.21 DPS) [quest]; Snake Hoop (6750, -1.34 DPS) [quest]; Reedknot Ring (9622, -1.59 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.2 spell_power points (3.48 DPS) | yes | Snake Hoop (6750, -0.57 DPS) [quest]; Reedknot Ring (9622, -0.83 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.18 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+3.9 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Celestial Stave (9517, -9.49 DPS) [quest]; Hypnotic Blade (7714, -10.25 DPS) [dungeon]; Illusionary Rod (7713, -11.18 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Icefury Wand (7514) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Freezing Shard (10572, -0.03 DPS) [dungeon]; Plaguerot Sprig (10766, -1.20 DPS) [dungeon]; Twisted Nether Wand (249144, -1.72 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Azure Shoulders; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Frothing Cord; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Staff of Jordan; ranged: Icefury Wand

No-known-source sample (15 of 456, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (gnome, 230005000000000000-00000000000000000-253510130101301221)

Set DPS (verified): 260.9. Weights run: 2.3s. Verify run: 5.9s. 598 eligible items had no known source.

Race sweep: human 261.3 ± 0.4, gnome 260.9 ± 0.4, high-order-skyborne 256.0 ± 0.4 (dps); human beats the guide race gnome by 0.4 (0.17%) ± 0.6, inside error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.429 ± 0.015, crit=0.374 ± 0.012 per rating point (14 rating = 1%, 5.242 per %), hit=0.802 ± 0.033 per rating point (10 rating = 1%, 8.016 per %), spell_haste=4.821 ± 0.216, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002, spirit=not significant (0.119 ± 0.063), mp5=0.454 ± 0.099

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 27.6 spell_power points (10.48 DPS) | yes | Spellpower Goggles Xtreme Plus (15999, -0.22 DPS) [crafted]; Dreamweave Circlet (10041, -0.33 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -1.21 DPS) [vendor] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 18.6 spell_power points (7.06 DPS) | yes | Talvash's Enhancing Necklace (7673, -2.58 DPS, sim-verified) [quest]; Glowing Eye of Mordresh (10769, -3.09 DPS) [dungeon]; Mindburst Medallion (11196, -3.57 DPS) [quest] |
| shoulder | Netherfroth Shoulders (254051) | Tailoring [crafted] | 21.1 spell_power points (8.04 DPS) | yes | Rotgrip Mantle (17732, -0.16 DPS) [dungeon]; Kentic Amice (11624, -0.33 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.54 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frost Leather Cloak (7377, -2.39 DPS) [crafted]; Azure Silk Cloak (7053, -2.63 DPS) [crafted]; Runecloth Cloak (13860, -3.43 DPS, sim-verified) [crafted] |
| chest | Frostweave Robe (13868) | Tailoring [crafted] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Robe of the Magi (1716, -0.05 DPS) [world_drop]; Dreamweave Vest (10021, -0.68 DPS) [crafted]; Frostweave Tunic (13869, -1.14 DPS, sim-verified) [crafted] |
| wrist | Netherfroth Cuffs (254063) | Tailoring [crafted] | 16.6 spell_power points (6.31 DPS) | yes | Nethergeld Cuffs (254061, -2.20 DPS, sim-verified) [crafted]; Aristocratic Cuffs (12546, -2.34 DPS) [dungeon]; Spidertank Oilrag (9448, -2.37 DPS) [dungeon] |
| hands | Frothing Gloves (254097) | Tailoring [crafted] | 24.2 spell_power points (9.21 DPS) | yes | Dreamweave Gloves (10019, -1.40 DPS) [crafted]; Frothing Handwraps (254023, -2.75 DPS) [crafted]; Frostweave Gloves (13870, -2.87 DPS, sim-verified) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 24.3 spell_power points (9.24 DPS) | yes | Frothing Waistcord (254083, -0.55 DPS) [crafted]; Satyrmane Sash (17755, -2.29 DPS) [dungeon]; Frothing Cord (254039, -2.34 DPS) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | 27.3 spell_power points (10.37 DPS) | yes | Knight's Dreadweave Leggings (220888, -1.86 DPS) [vendor]; Red Mageweave Pants (10009, -3.09 DPS) [crafted]; Wizardweave Leggings (14132, -3.15 DPS) [crafted] |
| feet | Frothing Sandals (254109) | Tailoring [crafted] | 23.9 spell_power points (9.07 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -1.51 DPS) [vendor]; Coldstone Slippers (18697, -1.92 DPS) [dungeon]; Frothing Slippers (254003, -2.81 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 15.9 spell_power points (6.04 DPS) | yes | Freezing Band (942, -1.10 DPS) [world_drop]; Band of the Unicorn (7553, -1.10 DPS) [world_drop]; Cyclopean Band (11824, -1.29 DPS) [dungeon] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 13.8 spell_power points (5.25 DPS) | yes | Freezing Band (942, -0.31 DPS) [world_drop]; Band of the Unicorn (7553, -0.31 DPS) [world_drop]; Cyclopean Band (11824, -0.51 DPS) [dungeon] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Enriched Seal (271903, +0.00 DPS) [rep] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Enriched Seal (271903, +0.00 DPS) [rep] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Kindling Stave (11750, +0.00 DPS) [dungeon]; Arbiter's Blade (11784, +0.00 DPS) [dungeon]; Staff of Jordan (873, -2.68 DPS, sim-verified) [world_drop] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 14.3 spell_power points (5.43 DPS) | yes | Orb of the Forgotten Seer (7685, -0.11 DPS) [dungeon]; Truesilver Conduit (249455, -0.33 DPS) [crafted]; Swampchill Fetish (1992, -1.25 DPS) [dungeon] |
| ranged | Icefury Wand (7514) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lesser Eternal Wand (249232, +0.00 DPS) [crafted]; Freezing Shard (10572, -0.26 DPS) [dungeon]; Noxious Shooter (17745, -0.60 DPS) [dungeon] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Netherfroth Shoulders; back: Spritecaster Cape; chest: Frostweave Robe; wrist: Netherfroth Cuffs; hands: Frothing Gloves; waist: Ban'thok Sash; legs: Spellshock Leggings; feet: Frothing Sandals; finger1: Brainlash; finger2: Lorekeeper's Ring; trinket1: Thunderbrew's Boot Flask; trinket2: Ankh of Life; main_hand: Blade of Eternal Darkness; off_hand: Celestial Orb

No-known-source sample (15 of 598, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (gnome, 230005203000000000-50000000000000000-253510130101301221)

Set DPS (verified): 414.8. Weights run: 2.7s. Verify run: 16.3s. 1456 eligible items had no known source.

Race sweep: gnome 414.8 ± 0.5, human 414.3 ± 0.5, high-order-skyborne 413.9 ± 0.5 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.405 ± 0.019, crit=0.565 ± 0.017 per rating point (14 rating = 1%, 7.915 per %), hit=1.151 ± 0.059 per rating point (10 rating = 1%, 11.515 per %), spell_haste=5.968 ± 0.285, spell_penetration=not significant (0.000 ± 0.000), frost_power=0.810 ± 0.002, spirit=0.539 ± 0.129, mp5=0.677 ± 0.158

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 51.5 spell_power points (21.16 DPS) | yes | Field Marshal's Coronet (231604, -0.20 DPS) [pvp]; Crimson Felt Hat (18727, -4.91 DPS, sim-verified) [dungeon]; Lieutenant Commander's Silk Cowl (227103, -4.96 DPS) [pvp] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 23.8 spell_power points (9.76 DPS) | yes | Star of Mystaria (12103, -0.89 DPS) [dungeon]; Lady Maye's Pendant (14558, -2.08 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -2.09 DPS) [world_drop] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 30.2 spell_power points (12.39 DPS) | yes | Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Mantle (227102, -0.27 DPS) [pvp]; Argent Shoulders (19059, -0.58 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -2.26 DPS, sim-verified) [world]; Shroud of the Exile (15421, -4.49 DPS) [quest]; Spritecaster Cape (11623, -4.56 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Vestments (231603, -0.72 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -5.87 DPS) [pvp]; Robe of Winter Night (14136, -10.31 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 29.0 spell_power points (11.91 DPS) | yes | Netherfroth Cuffs (254063, -5.32 DPS) [crafted]; Wyrmthalak's Shackles (13958, -5.72 DPS) [quest]; Runecloth Cuffs (254123, -5.73 DPS) [crafted] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 31.0 spell_power points (12.71 DPS) | yes | Gloves of Spell Mastery (14146, +0.00 DPS) [crafted]; Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 28.7 spell_power points (11.79 DPS) | yes | Ban'thok Sash (11662, -0.47 DPS) [dungeon]; Highlander's Cloth Girdle (20047, -1.80 DPS) [rep]; Thuzadin Sash (18740, -2.44 DPS) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 51.9 spell_power points (21.31 DPS) | yes | Marshal's Silk Leggings (231605, -0.20 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -5.50 DPS) [pvp]; Ghostweave Pants (14144, -6.20 DPS) [crafted] |
| feet | Venomspew Footpads (275606) | Tailoring [crafted] | 26.3 spell_power points (10.80 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Knight-Lieutenant's Silk Walkers (227112, -0.52 DPS) [pvp]; Omnicast Boots (11822, -0.60 DPS) [dungeon] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -1.78 DPS) [vendor]; Songstone of Ironforge (12543, -2.05 DPS) [quest]; Naglering (11669, -9.94 DPS, sim-verified) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -1.06 DPS) [vendor]; Songstone of Ironforge (12543, -1.33 DPS) [quest]; The Postmaster's Seal (13392, -1.39 DPS) [dungeon] |
| trinket1 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (+4.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, -2.66 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -3.14 DPS) [crafted]; Serenity Field (272439, -4.61 DPS, sim-verified) [vendor] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironbark Staff (20069, +0.00 DPS) [rep]; Grand Marshal's Stave (234571, -0.74 DPS) [pvp]; Blade of Eternal Darkness (17780, -7.58 DPS, sim-verified) [dungeon] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (11.91 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Tome of the Ice Lord (19310, -0.27 DPS) [rep]; Spirit of Aquementas (11904, -3.29 DPS) [quest] |
| ranged | Wand of Biting Cold (19108) | Korrak the Bloodrager [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, -0.01 DPS) [dungeon]; Oblivion's Touch (18761, -0.50 DPS) [dungeon]; Brilliant Wand (249385, -0.59 DPS) [crafted] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Magician's Cord; legs: Sentinel's Silk Leggings; feet: Venomspew Footpads; finger1: Blessed Band of Light; finger2: Maiden's Circle; trinket1: Second Wind; trinket2: Weakness Analyzer; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Wand of Biting Cold

No-known-source sample (15 of 1456, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (gnome, 230005203000000000-50000000000000000-253510130101301221)

Set DPS (verified): 696.4. Weights run: 2.9s. Verify run: 18.3s. 1456 eligible items had no known source.

Race sweep: high-order-skyborne 701.1 ± 0.6, gnome 696.4 ± 0.7, human 695.2 ± 0.6 (dps); high-order-skyborne beats the guide race gnome by 4.8 (0.68%) ± 0.9, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.172 ± 0.009, crit=0.674 ± 0.020 per rating point (14 rating = 1%, 9.431 per %), hit=1.582 ± 0.071 per rating point (10 rating = 1%, 15.820 per %), spell_haste=10.486 ± 0.314, spell_penetration=not significant (0.000 ± 0.000), frost_power=0.815 ± 0.002, spirit=not significant (0.000 ± 0.000), mp5=not significant (-0.001 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 48.8 spell_power points (24.94 DPS) | yes | Field Marshal's Coronet (231604, -1.77 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -7.81 DPS) [pvp]; Crimson Felt Hat (18727, -8.05 DPS, sim-verified) [dungeon] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 26.2 spell_power points (13.38 DPS) | yes | Star of Mystaria (12103, -3.39 DPS, sim-verified) [dungeon]; Arcane Crystal Pendant (20037, -4.68 DPS) [quest]; Emberfury Talisman (12929, -5.50 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 28.7 spell_power points (14.64 DPS) | yes | Field Marshal's Silk Spaulders (231602, -0.56 DPS) [pvp]; Lieutenant Commander's Silk Mantle (227102, -1.20 DPS) [pvp]; Argent Shoulders (19059, -1.87 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -6.17 DPS, sim-verified) [world]; Shroud of Arcane Mastery (22330, -7.91 DPS) [dungeon]; Cloak of Flames (3475, -8.87 DPS) [world_drop] |
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

No-known-source sample (15 of 1456, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (undead, 000000000000000000-00000000000000000-053200010000000000)

Set DPS (verified): 49.2. Weights run: 2.2s. Verify run: 2.3s. 192 eligible items had no known source.

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

Set DPS (verified): 88.0. Weights run: 2.2s. Verify run: 1.7s. 318 eligible items had no known source.

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

Set DPS (verified): 172.4. Weights run: 2.4s. Verify run: 4.4s. 448 eligible items had no known source.

Race sweep: undead 172.4 ± 0.2, orc 171.2 ± 0.2, troll 170.3 ± 0.2 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.684 ± 0.016, crit=0.164 ± 0.006 per rating point (14 rating = 1%, 2.297 per %), hit=0.343 ± 0.017 per rating point (10 rating = 1%, 3.430 per %), spell_haste=1.030 ± 0.198, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003, spirit=not significant (0.538 ± 0.144), mp5=0.734 ± 0.183

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Miner's Hat of the Deep (9429) | Uldaman [dungeon] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Spellpower Goggles Xtreme (10502, -0.67 DPS) [crafted]; Papal Fez (9431, -0.83 DPS, sim-verified) [dungeon]; Whitemane's Chapeau (7720, -0.86 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.5 spell_power points (3.20 DPS) | yes | Triune Amulet (7722, -0.18 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.89 DPS) [quest]; Necklace of Calisea (1714, -1.01 DPS) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.9 spell_power points (4.08 DPS) | yes | Green Silken Shoulders (7057, -0.09 DPS) [crafted]; Mantle of Doan (7712, -0.12 DPS) [dungeon]; Bloodmage Mantle (7684, -0.19 DPS) [dungeon] |
| back | Azure Silk Cloak (7053) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frost Leather Cloak (7377, -0.16 DPS) [crafted]; Icy Cloak (4327, -0.36 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.50 DPS) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Robe of Power (7054, -0.59 DPS) [crafted]; Dreamweave Vest (10021, -0.71 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -1.87 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 11.2 spell_power points (2.88 DPS) | yes | Condor Bracers (15864, -1.08 DPS) [quest]; Radiant Silver Bracers (4545, -1.08 DPS, sim-verified) [quest]; Earthen Silk Cuffs (254019, -1.16 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 24.5 spell_power points (6.29 DPS) | yes | Red Mageweave Gloves (10018, -1.71 DPS) [crafted]; Frothing Handwraps (254023, -1.75 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -2.44 DPS) [crafted] |
| waist | Frothing Cord (254039) | Tailoring [crafted] | 22.7 spell_power points (5.83 DPS) | yes | Defiler's Cloth Girdle (20166, -1.53 DPS) [rep]; Gilded Cord (254037, -1.54 DPS) [crafted]; Deathmage Sash (10771, -2.39 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.2 spell_power points (5.70 DPS) | yes | Silky Leggings (253989, -0.19 DPS) [crafted]; Stoneweaver Leggings (9407, -0.84 DPS) [dungeon]; Filigreed Silky Leggings (253939, -1.27 DPS) [crafted] |
| feet | Frothing Slippers (254003) | Tailoring [crafted] | 19.9 spell_power points (5.12 DPS) | yes | Furen's Boots (13100, -1.68 DPS) [world_drop]; Acidic Walkers (9454, -1.88 DPS) [dungeon]; Gilded Slippers (254001, -2.27 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 14.5 spell_power points (3.74 DPS) | yes | Advisor's Ring (19520, -0.86 DPS) [rep]; Runemender's Seal (273031, -0.95 DPS) [world] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.4 spell_power points (3.45 DPS) | yes | Advisor's Ring (19520, -0.57 DPS) [rep]; Runemender's Seal (273031, -0.66 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -8.25 DPS, sim-verified) [dungeon]; Celestial Stave (9517, -8.41 DPS) [quest]; Hypnotic Blade (7714, -8.51 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Freezing Shard (10572) | Razorfen Downs [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Icefury Wand (7514, +0.00 DPS) [quest]; Dancing Flame (6806, -0.94 DPS) [quest]; Plaguerot Sprig (10766, -0.95 DPS) [dungeon] |

**New at 40:** head: Miner's Hat of the Deep; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Frothing Cord; legs: Red Mageweave Pants; finger2: Philanthropist's Ring; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Staff of Jordan; ranged: Freezing Shard

No-known-source sample (15 of 448, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (undead, 230005000000000000-00000000000000000-253510130101301221)

Set DPS (verified): 258.7. Weights run: 2.3s. Verify run: 5.9s. 590 eligible items had no known source.

Race sweep: undead 258.7 ± 0.4, orc 256.7 ± 0.4, troll 256.5 ± 0.4 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.638 ± 0.027, crit=0.279 ± 0.011 per rating point (14 rating = 1%, 3.908 per %), hit=0.696 ± 0.036 per rating point (10 rating = 1%, 6.960 per %), spell_haste=not significant (1.410 ± 0.386), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003, spirit=not significant (1.150 ± 0.366), mp5=not significant (1.193 ± 0.391)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreamweave Circlet (10041) | Tailoring [crafted] | 41.2 spell_power points (9.95 DPS) | yes | Red Mageweave Headband (10033, -2.28 DPS) [crafted]; Papal Fez (9431, -2.90 DPS) [dungeon]; Gemburst Circlet (10751, -4.68 DPS, sim-verified) [quest] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 19.8 spell_power points (4.79 DPS) | yes | Triune Amulet (7722, -0.91 DPS) [dungeon]; Mindburst Medallion (11196, -1.03 DPS) [quest]; Lei of Lilies (1315, -5.88 DPS, sim-verified) [world_drop] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 29.2 spell_power points (7.06 DPS) | yes | Netherfroth Shoulders (254051, -0.05 DPS) [crafted]; Rotgrip Mantle (17732, -1.14 DPS) [dungeon]; Nethergeld Shoulders (254049, -1.98 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -1.69 DPS) [quest]; Cloak of Blight (6832, -2.36 DPS) [quest]; Featherskin Cape (10843, -4.94 DPS, sim-verified) [world] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 51.9 spell_power points (12.54 DPS) | yes | Ghostweave Vest (14141, -2.73 DPS, sim-verified) [crafted]; Dreamweave Vest (10021, -2.91 DPS) [crafted]; Vestments of the Atal'ai Prophet (10806, -3.28 DPS) [dungeon] |
| wrist | Netherfroth Cuffs (254063) | Tailoring [crafted] | 23.2 spell_power points (5.61 DPS) | yes | Bloodband Bracers (11469, -2.18 DPS) [quest]; Nethergeld Cuffs (254061, -2.22 DPS, sim-verified) [crafted]; Aristocratic Cuffs (12546, -2.33 DPS) [dungeon] |
| hands | Frothing Gloves (254097) | Tailoring [crafted] | 34.6 spell_power points (8.36 DPS) | yes | Runecloth Gloves (13863, -1.57 DPS) [crafted]; Gloves of the Atal'ai Prophet (10808, -1.83 DPS) [dungeon]; Dreamweave Gloves (10019, -2.94 DPS, sim-verified) [crafted] |
| waist | Frothing Cord (254039) | Tailoring [crafted] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Ban'thok Sash (11662, -0.16 DPS) [dungeon]; Frothing Waistcord (254083, -0.31 DPS) [crafted]; Ghostweave Cord (254073, -1.77 DPS, sim-verified) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 37.2 spell_power points (8.98 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -0.48 DPS) [dungeon]; Dalewind Trousers (13008, -0.65 DPS) [world_drop]; Rainstrider Leggings (11123, -3.81 DPS, sim-verified) [quest] |
| feet | Frothing Sandals (254109) | Tailoring [crafted] | sim-verified (+5.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Coldstone Slippers (18697, -0.25 DPS) [dungeon]; Vinerot Sandals (17748, -0.60 DPS) [dungeon]; Mistwalker Boots (10629, -5.52 DPS, sim-verified) [dungeon] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 23.1 spell_power points (5.59 DPS) | yes | Eye of Adaegus (5266, -1.33 DPS) [world_drop]; Advisor's Ring (19519, -1.54 DPS) [rep]; Snake Hoop (6750, -7.33 DPS, sim-verified) [quest] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Eye of Adaegus (5266, -0.10 DPS) [world_drop]; Advisor's Ring (19519, -0.31 DPS) [rep]; Snake Hoop (6750, -1.11 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Uther's Strength (11302, -2.01 DPS, sim-verified) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Spire of Hakkar (10844, +0.00 DPS) [world]; Staff of Jordan (873, -1.90 DPS, sim-verified) [world_drop] |
| off_hand | Enthralled Sphere (11625) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Truesilver Conduit (249455, -0.02 DPS) [crafted]; Celestial Orb (7515, -0.07 DPS) [quest]; Twisting Essence Jar (249456, -1.61 DPS, sim-verified) [crafted] |
| ranged | Icefury Wand (7514) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Freezing Shard (10572, -0.27 DPS) [dungeon]; Lesser Eternal Wand (249232, -0.41 DPS) [crafted]; Jaina's Firestarter (13064, -0.80 DPS) [world_drop] |

**New at 50:** head: Dreamweave Circlet; neck: Arcane Crystal Pendant; shoulder: Kentic Amice; back: Spritecaster Cape; chest: Embrace of the Wind Serpent; wrist: Netherfroth Cuffs; hands: Frothing Gloves; legs: Senior Designer's Pantaloons; feet: Frothing Sandals; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Blade of Eternal Darkness; off_hand: Enthralled Sphere; ranged: Icefury Wand

No-known-source sample (15 of 590, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (undead, 230005203000000000-50000000000000000-253510130101301221)

Set DPS (verified): 410.2. Weights run: 2.7s. Verify run: 15.3s. 1450 eligible items had no known source.

Race sweep: undead 410.2 ± 0.5, orc 406.8 ± 0.5, troll 406.1 ± 0.5 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.418 ± 0.030, crit=0.359 ± 0.011 per rating point (14 rating = 1%, 5.028 per %), hit=0.894 ± 0.050 per rating point (10 rating = 1%, 8.942 per %), spell_haste=2.984 ± 0.415, spell_penetration=not significant (0.000 ± 0.000), frost_power=0.806 ± 0.002, spirit=not significant (0.671 ± 0.353), mp5=not significant (1.303 ± 0.403)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 50.7 spell_power points (15.11 DPS) | yes | Warlord's Silk Cowl (231601, -0.45 DPS) [pvp]; Champion's Silk Cowl (227105, -3.90 DPS) [pvp]; Crimson Felt Hat (18727, -5.50 DPS, sim-verified) [dungeon] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 21.3 spell_power points (6.35 DPS) | yes | Star of Mystaria (12103, -0.61 DPS) [dungeon]; Lady Maye's Pendant (14558, -0.70 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -3.19 DPS, sim-verified) [world_drop] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 29.7 spell_power points (8.86 DPS) | yes | Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -0.67 DPS) [crafted]; Champion's Silk Mantle (227104, -0.71 DPS) [pvp] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, -2.21 DPS) [quest]; Spritecaster Cape (11623, -2.31 DPS) [dungeon]; Crystalline Threaded Cape (20697, -4.40 DPS, sim-verified) [world] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Raiment (231596, -0.26 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -4.04 DPS) [pvp]; Robe of Winter Night (14136, -7.60 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 30.0 spell_power points (8.96 DPS) | yes | Wyrmthalak's Shackles (13958, -3.87 DPS) [quest]; Netherfroth Cuffs (254063, -3.96 DPS) [crafted]; Runecloth Cuffs (254123, -4.43 DPS) [crafted] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 32.1 spell_power points (9.58 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Mana Shaping Handwraps (22256, -2.35 DPS, sim-verified) [dungeon] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 29.0 spell_power points (8.66 DPS) | yes | Ban'thok Sash (11662, -1.17 DPS) [dungeon]; Thuzadin Sash (18740, -1.38 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -1.83 DPS) [crafted] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 42.7 spell_power points (12.72 DPS) | yes | General's Silk Trousers (231595, +0.00 DPS) [pvp]; Legionnaire's Silk Legguards (227107, -1.84 DPS) [pvp]; Ghostweave Pants (14144, -3.57 DPS, sim-verified) [crafted] |
| feet | Venomspew Footpads (275606) | Tailoring [crafted] | 27.4 spell_power points (8.16 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Blood Guard's Silk Walkers (227110, -0.65 DPS) [pvp]; Omnicast Boots (11822, -0.70 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Postmaster's Seal (13392, -0.79 DPS) [dungeon]; Eye of Orgrimmar (12545, -1.20 DPS) [quest]; Naglering (11669, -8.54 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -1.01 DPS) [quest]; Channeler's Ring (272406, -1.24 DPS) [vendor]; The Postmaster's Seal (13392, -3.75 DPS, sim-verified) [dungeon] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Burst of Knowledge (11832, -2.26 DPS) [dungeon]; Smolderweb's Eye (13213, -2.70 DPS) [dungeon]; Second Wind (11819, -3.42 DPS, sim-verified) [dungeon] |
| trinket2 | Eye of the Beast (13968) | For The Horde! [quest] | sim-verified (-1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironbark Staff (20220, -0.34 DPS) [rep]; Headmaster's Charge (13937, -0.60 DPS) [dungeon]; Blade of Eternal Darkness (17780, -8.70 DPS, sim-verified) [dungeon] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (8.65 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; Tome of the Ice Lord (19310, -0.20 DPS) [rep]; Spirit of Aquementas (11904, -2.39 DPS) [quest] |
| ranged | Wand of Biting Cold (19108) | The Legend of Korrak [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Oblivion's Touch (18761, -0.31 DPS) [dungeon]; Brilliant Wand (249385, -0.35 DPS) [crafted] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Argent Shoulders; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Magician's Cord; legs: Outrider's Silk Leggings; feet: Venomspew Footpads; finger1: Maiden's Circle; finger2: Blessed Band of Light; trinket1: Weakness Analyzer; trinket2: Eye of the Beast; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Wand of Biting Cold

No-known-source sample (15 of 1450, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60, raid preset (undead, 230005203000000000-50000000000000000-253510130101301221)

Set DPS (verified): 688.8. Weights run: 2.9s. Verify run: 18.7s. 1450 eligible items had no known source.

Race sweep: undead 688.8 ± 0.6, troll 685.1 ± 0.6, orc 684.9 ± 0.6 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.172 ± 0.009, crit=0.674 ± 0.020 per rating point (14 rating = 1%, 9.431 per %), hit=1.582 ± 0.071 per rating point (10 rating = 1%, 15.820 per %), spell_haste=10.486 ± 0.314, spell_penetration=not significant (0.000 ± 0.000), frost_power=0.815 ± 0.002, spirit=not significant (0.000 ± 0.000), mp5=not significant (-0.001 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 48.8 spell_power points (24.94 DPS) | yes | Warlord's Silk Cowl (231601, -1.77 DPS) [pvp]; Champion's Silk Cowl (227105, -7.81 DPS) [pvp]; Crimson Felt Hat (18727, -8.86 DPS, sim-verified) [dungeon] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 26.2 spell_power points (13.38 DPS) | yes | Star of Mystaria (12103, -3.19 DPS, sim-verified) [dungeon]; Arcane Crystal Pendant (20037, -4.68 DPS) [quest]; Emberfury Talisman (12929, -5.50 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 28.7 spell_power points (14.64 DPS) | yes | Warlord's Silk Amice (231594, -0.56 DPS) [pvp]; Champion's Silk Mantle (227104, -1.20 DPS) [pvp]; Argent Shoulders (19059, -1.87 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -6.15 DPS, sim-verified) [world]; Shroud of Arcane Mastery (22330, -7.91 DPS) [dungeon]; Cloak of Flames (3475, -8.87 DPS) [world_drop] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Raiment (231596, -3.14 DPS) [pvp]; Robe of Winter Night (14136, -8.86 DPS, sim-verified) [crafted]; Legionnaire's Silk Tunic (227106, -9.27 DPS) [pvp] |
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

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Ban'thok Sash; legs: Skyshroud Leggings; feet: First Sergeant's Dreadweave Boots; finger1: Blessed Band of Light; finger2: Don Julio's Band; trinket1: Weakness Analyzer; trinket2: Smolderweb's Eye; main_hand: Ironbark Staff; ranged: Wand of Biting Cold

No-known-source sample (15 of 1450, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

