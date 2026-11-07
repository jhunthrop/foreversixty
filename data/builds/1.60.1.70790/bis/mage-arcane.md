# Leveling BiS: Arcane

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 140005000100000000-00000000000000000-000000000000000000)

Set DPS (verified): 48.0. Weights run: 2.6s. Verify run: 3.0s. 181 eligible items had no known source.

Race sweep: gnome 48.0 ± 0.1, human 46.5 ± 0.1, high-order-skyborne 46.3 ± 0.1 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.207 ± 0.006, crit=0.064 ± 0.002 per rating point (14 rating = 1%, 0.890 per %), hit=0.225 ± 0.007 per rating point (10 rating = 1%, 2.252 per %), spell_haste=0.354 ± 0.051, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.855 ± 0.005, spirit=not significant (0.030 ± 0.040), mp5=not significant (-0.007 ± 0.056)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pearly Circlet (253957) | Tailoring [crafted] | 9.4 spell_power points (1.33 DPS) | yes | Pristine Circlet (253949, -0.88 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.16 DPS) [crafted]; Flying Tiger Goggles (4368, -1.31 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.1 spell_power points (0.01 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 6.9 spell_power points (0.97 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.41 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.86 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Caretaker's Cape (20428, -0.15 DPS) [rep]; Black Whelp Cloak (7283, -0.16 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.21 DPS) [crafted] |
| chest | Filigreed Pearly Gown (253909) | Tailoring [crafted] | 8.8 spell_power points (1.25 DPS) | yes | Gray Woolen Robe (2585, -0.54 DPS) [crafted]; Green Woolen Vest (2582, -0.67 DPS) [crafted]; Filigreed Pristine Gown (253901, -0.70 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 1.0 spell_power points (0.15 DPS) | yes | Bright Bracers (3647, -0.03 DPS) [world_drop]; Repurposed Hair Band (281256, -0.08 DPS) [quest]; Windsong Bangles (263336, -1.30 DPS, sim-verified) [quest] |
| hands | Pearly Gloves (253921) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Gnoll Casting Gloves (892, -0.09 DPS) [world]; Pristine Gloves (253913, -0.28 DPS) [crafted]; Serpent Gloves (5970, -0.53 DPS, sim-verified) [dungeon] |
| waist | Pearly Sash (253933) | Tailoring [crafted] | 6.8 spell_power points (0.96 DPS) | yes | Novice Arcanist's Sash (253885, -0.31 DPS) [crafted]; Pristine Sash (253925, -0.52 DPS, sim-verified) [crafted]; Novice Ardent's Sash (253887, -0.58 DPS) [crafted] |
| legs | Filigreed Pearly Leggings (253945) | Tailoring [crafted] | 10.8 spell_power points (1.52 DPS) | yes | Silk-threaded Trousers (1929, -0.53 DPS) [dungeon]; Colorful Kilt (10048, -0.80 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.88 DPS, sim-verified) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.8 spell_power points (1.11 DPS) | yes | Pearly Boots (253897, -0.37 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.52 DPS) [crafted]; Pristine Boots (253889, -0.60 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.4 spell_power points (0.77 DPS) | yes | Lavishly Jeweled Ring (1156, -0.59 DPS) [dungeon]; Furlfeather Ring (286745, -0.62 DPS) [world]; Volcanic Rock Ring (12053, -0.68 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.71 DPS) | yes | Lavishly Jeweled Ring (1156, -0.53 DPS) [dungeon]; Furlfeather Ring (286745, -0.57 DPS) [world]; Volcanic Rock Ring (12053, -0.62 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | 28.6 spell_power points (4.04 DPS) | yes | Emberstone Staff (5201, -0.65 DPS) [dungeon]; Staff of the Blessed Seer (2271, -0.74 DPS, sim-verified) [dungeon]; Buzzer Blade (2169, -0.93 DPS) [dungeon] |
| off_hand | Mathystra Channeling Orb (281291) | Baron Marinous [quest] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Tear of Grief (5611, -0.01 DPS) [quest]; Seer's Fine Stein (7608, -0.05 DPS) [world_drop]; Dwarven Tome (279898, -0.16 DPS, sim-verified) [quest] |
| ranged | Skycaller (12984) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sizzle Stick (8071, -0.83 DPS) [quest]; Greater Magic Wand (11288, -0.85 DPS) [crafted]; Cookie's Stirring Rod (5198, -1.06 DPS, sim-verified) [dungeon] |

**New at 20:** head: Pearly Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pearly Gown; wrist: Mindthrust Bracers; hands: Pearly Gloves; waist: Pearly Sash; legs: Filigreed Pearly Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Evocator's Blade; off_hand: Mathystra Channeling Orb; ranged: Skycaller

No-known-source sample (15 of 181, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 30 (gnome, 141005203100301000-00000000000000000-000000000000000000)

Set DPS (verified): 92.1. Weights run: 2.9s. Verify run: 3.2s. 309 eligible items had no known source.

Race sweep: gnome 92.1 ± 0.1, human 90.7 ± 0.1, high-order-skyborne 90.6 ± 0.1 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.555 ± 0.010, crit=0.066 ± 0.002 per rating point (14 rating = 1%, 0.924 per %), hit=0.235 ± 0.012 per rating point (10 rating = 1%, 2.352 per %), spell_haste=0.844 ± 0.143, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.583 ± 0.003, spirit=not significant (0.101 ± 0.088), mp5=not significant (0.232 ± 0.126)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.6 spell_power points (2.73 DPS) | yes | Enchanter's Cowl (4322, -0.01 DPS) [crafted]; Silk Headband (7050, -0.61 DPS) [crafted]; Filigreed Pearly Circlet (253983, -0.81 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.3 spell_power points (0.78 DPS) | yes | Seal of the Expedition (271920, -0.13 DPS) [rep]; Kaleidoscope Chain (13084, -0.17 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.19 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.29 DPS) | yes | Death Speaker Mantle (6685, -0.44 DPS) [dungeon]; Magician's Mantle (12998, -0.94 DPS) [world_drop]; Invoker's Mantle (215365, -0.99 DPS) [crafted] |
| back | Vine Pruner's Cloak (279835) | A Green Sample [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Expeditionary's Cape (271923, -0.06 DPS) [rep]; Heavy Woolen Cloak (4311, -0.14 DPS) [crafted] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 16.5 spell_power points (3.87 DPS) | yes | Green Silk Armor (7065, -0.06 DPS) [crafted]; Mechbuilder's Overalls (9508, -0.43 DPS) [dungeon]; Death Speaker Robes (6682, -0.79 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.7 spell_power points (2.28 DPS) | yes | Glowing Magical Bracelets (13106, -1.25 DPS, sim-verified) [world_drop]; Nightsky Wristbands (6407, -1.43 DPS) [world_drop]; Mindthrust Bracers (1974, -1.46 DPS) [dungeon] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 9.1 spell_power points (2.15 DPS) | yes | Town Clerk's Mittens (270029, -0.53 DPS) [quest]; Truefaith Gloves (7049, -0.58 DPS) [crafted]; Serpent Gloves (5970, -0.63 DPS, sim-verified) [dungeon] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.7 spell_power points (2.98 DPS) | yes | Belt of Arugal (6392, -0.35 DPS) [dungeon]; Crimson Silk Belt (7055, -0.65 DPS) [crafted]; Invoker's Cord (215366, -0.68 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.9 spell_power points (3.04 DPS) | yes | Pearly Leggings (253995, -0.09 DPS) [crafted]; Pristine Leggings (253987, -0.36 DPS) [crafted]; Filigreed Pearly Leggings (253945, -0.65 DPS) [crafted] |
| feet | Golden Slippers (254009) | Tailoring [crafted] | 11.9 spell_power points (2.79 DPS) | yes | Gilded Slippers (254001, -0.14 DPS) [crafted]; Acidic Walkers (9454, -0.47 DPS) [dungeon]; Spidersilk Boots (4320, -0.62 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.8 spell_power points (2.06 DPS) | yes | Ring of Power Regulation (273030, -0.49 DPS) [dungeon]; Minor Channeling Ring (1449, -0.63 DPS) [quest]; Electrocutioner Lagnut (9447, -1.19 DPS) [dungeon] |
| finger2 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.5 spell_power points (1.75 DPS) | yes | Ring of Power Regulation (273030, -0.18 DPS) [dungeon]; Minor Channeling Ring (1449, -0.32 DPS) [quest]; Electrocutioner Lagnut (9447, -0.88 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.04 DPS) [dungeon]; Tendonscraper Dagger (273053, -1.04 DPS) [world]; Glimmering Staff (249392, -2.61 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.3 spell_power points (2.43 DPS) | yes | Arcane Infused Rod (279838, -0.92 DPS) [quest]; Orb of Souls (249395, -1.35 DPS) [crafted]; Eye of Paleth (2943, -1.49 DPS) [quest] |
| ranged | Starfaller (13063) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderwood (13062, -0.15 DPS) [world_drop]; Greater Mystic Wand (217287, -0.31 DPS) [crafted]; Wand of Eventide (5214, -0.43 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Golden Slippers; finger1: Snake Hoop; finger2: Lorekeeper's Ring; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Starfaller

No-known-source sample (15 of 309, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 141005213100301531-00000000000000000-000000000000000000)

Set DPS (verified): 174.0. Weights run: 3.0s. Verify run: 3.6s. 400 eligible items had no known source.

Race sweep: gnome 174.0 ± 0.3, human 173.9 ± 0.3, high-order-skyborne 172.5 ± 0.3 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.553 ± 0.025, crit=0.270 ± 0.006 per rating point (14 rating = 1%, 3.779 per %), hit=0.632 ± 0.034 per rating point (10 rating = 1%, 6.323 per %), spell_haste=2.432 ± 0.406, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.638 ± 0.002, spirit=not significant (0.651 ± 0.231), mp5=not significant (0.797 ± 0.255)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 23.3 spell_power points (7.55 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.73 DPS) [crafted]; Whitemane's Chapeau (7720, -0.95 DPS) [dungeon]; Electromagnetic Gigaflux Reactivator (9492, -1.62 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Triune Amulet (7722, -0.46 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -0.97 DPS) [dungeon]; Amberglow Talisman (10824, -1.60 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Green Silken Shoulders (7057, -0.03 DPS) [crafted]; Bloodmage Mantle (7684, -0.07 DPS) [dungeon]; Mantle of Doan (7712, -0.59 DPS, sim-verified) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Caretaker's Cape (19532, -0.54 DPS) [rep]; Long Silken Cloak (4326, -0.69 DPS) [crafted]; Guardian Cloak (5965, -0.69 DPS) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 32.1 spell_power points (10.42 DPS) | yes | Robe of the Magi (1716, -1.14 DPS) [world_drop]; Robe of Power (7054, -2.03 DPS) [crafted]; Elemental Raiment (9434, -3.60 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 11.4 spell_power points (3.70 DPS) | yes | Arcane Runed Bracers (4744, -0.54 DPS, sim-verified) [quest]; Earthen Silk Cuffs (254019, -1.34 DPS) [crafted]; Condor Bracers (15864, -1.43 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 24.8 spell_power points (8.04 DPS) | yes | Red Mageweave Gloves (10018, -1.51 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -3.17 DPS) [crafted]; Golden Handwraps (254029, -3.88 DPS) [crafted] |
| waist | Golden Cord (254045) | Tailoring [crafted] | 17.3 spell_power points (5.60 DPS) | yes | Gilded Cord (254037, -0.30 DPS) [crafted]; Highlander's Cloth Girdle (20098, -0.34 DPS) [rep]; Star Belt (4329, -0.54 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 20.6 spell_power points (6.70 DPS) | yes | Gaze Dreamer Pants (6903, -0.90 DPS) [dungeon]; Black Mageweave Leggings (9999, -1.15 DPS) [crafted]; Stoneweaver Leggings (9407, -2.73 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 28.6 spell_power points (9.27 DPS) | yes | Golden Slippers (254009, -3.42 DPS, sim-verified) [crafted]; Furen's Boots (13100, -4.63 DPS) [world_drop]; Gilded Slippers (254001, -4.89 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.8 spell_power points (4.14 DPS) | yes | Snake Hoop (6750, -1.30 DPS) [quest]; Ring of Power Regulation (273030, -1.42 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -1.55 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 11.4 spell_power points (3.70 DPS) | yes | Snake Hoop (6750, -0.85 DPS) [quest]; Ring of Power Regulation (273030, -0.97 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -1.10 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -9.37 DPS, sim-verified) [dungeon]; Celestial Stave (9517, -11.89 DPS) [quest]; Hypnotic Blade (7714, -11.91 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Nether Force Wand (11263) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Starfaller (13063, -0.52 DPS) [world_drop]; Burning Sliver (5249, -0.54 DPS) [quest]; Twisted Nether Wand (249144, -0.60 DPS, sim-verified) [crafted] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Icy Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Golden Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Staff of Jordan; ranged: Nether Force Wand

No-known-source sample (15 of 400, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (gnome, 141005213100301531-40000000000000000-033000000000000000)

Set DPS (verified): 263.6. Weights run: 3.0s. Verify run: 4.5s. 512 eligible items had no known source.

Race sweep: gnome 263.6 ± 0.4, human 261.8 ± 0.4, high-order-skyborne 258.6 ± 0.4 (dps); the guide race gnome is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.574 ± 0.025, crit=0.347 ± 0.008 per rating point (14 rating = 1%, 4.856 per %), hit=0.760 ± 0.038 per rating point (10 rating = 1%, 7.603 per %), spell_haste=3.201 ± 0.402, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.636 ± 0.002, spirit=not significant (0.283 ± 0.192), mp5=not significant (0.503 ± 0.228)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 30.5 spell_power points (11.39 DPS) | yes | Dreamweave Circlet (10041, -0.13 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -1.30 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -1.77 DPS) [vendor] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 19.4 spell_power points (7.26 DPS) | yes | Glowing Eye of Mordresh (10769, -2.43 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -3.21 DPS) [quest]; Triune Amulet (7722, -3.54 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 23.3 spell_power points (8.72 DPS) | yes | Kentic Amice (11624, -0.06 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.98 DPS) [vendor]; Nethershine Shoulders (254057, -2.33 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Big Voodoo Cloak (8216, -3.35 DPS) [crafted]; Expeditionary's Cape (271921, -3.62 DPS) [rep]; Runecloth Cloak (13860, -4.94 DPS, sim-verified) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 30.5 spell_power points (11.39 DPS) | yes | Robe of the Magi (1716, -1.35 DPS) [world_drop]; Runecloth Tunic (13857, -2.68 DPS) [crafted]; Dreamweave Vest (10021, -2.76 DPS, sim-verified) [crafted] |
| wrist | Nethershine Cuffs (254069) | Tailoring [crafted] | 13.7 spell_power points (5.12 DPS) | yes | Nethergeld Cuffs (254061, -0.47 DPS) [crafted]; Bloodband Bracers (11469, -1.00 DPS) [quest]; Spidertank Oilrag (9448, -1.19 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.3 spell_power points (8.32 DPS) | yes | Golden Gloves (254103, -0.82 DPS) [crafted]; Runecloth Gloves (13863, -0.96 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.54 DPS) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 25.9 spell_power points (9.68 DPS) | yes | Golden Waistcord (254089, -3.24 DPS) [crafted]; Ghostweave Cord (254073, -3.29 DPS) [crafted]; Satyrmane Sash (17755, -3.69 DPS, sim-verified) [dungeon] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshock Leggings (9484, +0.00 DPS) [dungeon]; Senior Designer's Pantaloons (11841, -0.98 DPS) [dungeon]; Red Mageweave Pants (10009, -1.07 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 26.0 spell_power points (9.71 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -2.40 DPS, sim-verified) [vendor]; Golden Sandals (254115, -3.03 DPS) [crafted]; Gilded Sandals (254107, -3.67 DPS) [crafted] |
| finger1 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 14.1 spell_power points (5.29 DPS) | yes | Band of the Unicorn (7553, -0.43 DPS) [world_drop]; Philanthropist's Ring (281635, -0.48 DPS) [quest]; Brainlash (6440, -1.54 DPS) [dungeon] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 14.0 spell_power points (5.24 DPS) | yes | Philanthropist's Ring (281635, -0.43 DPS) [quest]; Brainlash (6440, -1.49 DPS) [dungeon]; Band of the Unicorn (7553, -1.93 DPS, sim-verified) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Evonice's Landin' Pilla (18951, -0.21 DPS) [quest]; Enriched Seal (271903, -0.33 DPS) [rep] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Evonice's Landin' Pilla (18951, -1.19 DPS) [quest]; Enriched Seal (271903, -1.30 DPS) [rep] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -5.22 DPS) [dungeon]; Glowing Brightwood Staff (812, -6.17 DPS) [world_drop]; Spire of Hakkar (10844, -9.19 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Noxious Shooter (17745, -0.77 DPS) [dungeon]; Nether Force Wand (11263, -0.90 DPS) [quest]; Dreambough Wand (249234, -0.91 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethershine Cuffs; waist: Ban'thok Sash; legs: Knight's Dreadweave Leggings; finger1: Cyclopean Band; finger2: Lorekeeper's Ring; trinket2: Uther's Strength; ranged: Lesser Eternal Wand

No-known-source sample (15 of 512, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (gnome, 141005213100301531-40000000000000000-053500030000000000)

Set DPS (verified): 417.3. Weights run: 3.1s. Verify run: 20.2s. 1299 eligible items had no known source.

Race sweep: gnome 417.3 ± 0.6, human 415.5 ± 0.6, high-order-skyborne 412.6 ± 0.6 (dps); the guide race gnome is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.843 ± 0.042, crit=0.572 ± 0.014 per rating point (14 rating = 1%, 8.010 per %), hit=1.114 ± 0.071 per rating point (10 rating = 1%, 11.139 per %), spell_haste=6.611 ± 0.874, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.613 ± 0.002, spirit=not significant (-0.003 ± 0.431), mp5=not significant (0.110 ± 0.461)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 36.7 spell_power points (14.34 DPS) | yes | Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Magister's Crown (16686, -0.17 DPS) [dungeon] |
| neck | Diana's Pearl Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 26.9 spell_power points (10.50 DPS) | yes | Arcane Crystal Pendant (20037, -2.27 DPS) [quest]; Archlight Talisman (15856, -2.52 DPS) [quest]; Star of Mystaria (12103, -3.18 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 36.0 spell_power points (14.04 DPS) | yes | Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Theramore Shoulderpads (271953, -0.47 DPS) [rep]; Burial Shawl (18681, -0.97 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -3.80 DPS, sim-verified) [world]; Shroud of Arcane Mastery (22330, -5.26 DPS) [dungeon]; Spritecaster Cape (11623, -5.79 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Vestments (231603, -1.09 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -5.77 DPS) [pvp]; Alanna's Embrace (13314, -6.88 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 28.7 spell_power points (11.22 DPS) | yes | Runecloth Cuffs (254123, -3.64 DPS) [crafted]; Marshal's Silk Bracers (16438, -4.97 DPS) [pvp]; Magiskull Cuffs (13107, -5.63 DPS) [world_drop] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 33.5 spell_power points (13.06 DPS) | yes | Hands of Power (13253, +0.00 DPS) [dungeon]; Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 39.2 spell_power points (15.32 DPS) | yes | Ban'thok Sash (11662, -2.66 DPS) [dungeon]; Clutch of Andros (13956, -4.71 DPS) [dungeon]; Highlander's Cloth Girdle (20047, -4.75 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 56.5 spell_power points (22.04 DPS) | yes | Marshal's Silk Leggings (231605, -0.62 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -5.12 DPS) [pvp]; Skyshroud Leggings (13170, -6.13 DPS) [dungeon] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dragonrider Boots (18102, +0.00 DPS) [dungeon]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Knight-Lieutenant's Silk Walkers (227112, -0.27 DPS) [pvp] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, +0.00 DPS) [quest]; Blessed Band of Light (272407, -1.08 DPS) [vendor]; Naglering (11669, -10.49 DPS, sim-verified) [dungeon] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Blessed Band of Light (272407, -0.71 DPS) [vendor]; Deep Rooted Ring (19109, -0.98 DPS) [world]; Songstone of Ironforge (12543, -1.39 DPS, sim-verified) [quest] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+5.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Serenity Field (272439, -3.12 DPS) [vendor]; Frozen Heart of the Mountain (249469, -4.67 DPS) [crafted]; Thunderbrew's Boot Flask (744, -4.96 DPS, sim-verified) [quest] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -0.77 DPS) [crafted]; Uther's Strength (11302, -2.34 DPS) [world_drop]; Shard of the Splithooves (10659, -4.43 DPS) [quest] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Sageclaw (20070, -5.54 DPS) [rep]; Ironbark Staff (20069, -6.57 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (11.32 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Tome of Restoration (234590, -0.86 DPS) [pvp]; Tome of the Lost (22253, -1.99 DPS) [dungeon] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, -0.01 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -1.25 DPS) [dungeon]; Oblivion's Touch (18761, -2.01 DPS) [dungeon] |

**New at 60:** head: Crimson Felt Hat; neck: Diana's Pearl Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Magician's Cord; legs: Sentinel's Silk Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Channeler's Ring; trinket1: Weakness Analyzer; trinket2: Burst of Knowledge; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Brilliant Wand

No-known-source sample (15 of 1299, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (gnome, 131205203100311531-52500001200100000-003000000000000000)

Set DPS (verified): 803.9. Weights run: 3.7s. Verify run: 22.4s. 1299 eligible items had no known source.

Race sweep: gnome 803.9 ± 0.9, human 800.0 ± 0.9, high-order-skyborne 794.0 ± 0.9 (dps); the guide race gnome is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.911 ± 0.048, crit=0.914 ± 0.038 per rating point (14 rating = 1%, 12.801 per %), hit=1.653 ± 0.108 per rating point (10 rating = 1%, 16.532 per %), spell_haste=4.035 ± 0.657, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.569 ± 0.002, spirit=not significant (0.267 ± 0.372), mp5=not significant (0.637 ± 0.432)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cap of the Scarlet Savant (12752, +0.00 DPS) [quest]; Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp] |
| neck | Diana's Pearl Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 32.8 spell_power points (17.22 DPS) | yes | Star of Mystaria (12103, -3.12 DPS) [dungeon]; Arcane Crystal Pendant (20037, -5.96 DPS) [quest]; Archlight Talisman (15856, -6.14 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 41.6 spell_power points (21.85 DPS) | yes | Field Marshal's Silk Spaulders (231602, -0.86 DPS) [pvp]; Lieutenant Commander's Silk Mantle (227102, -1.44 DPS) [pvp]; Theramore Shoulderpads (271953, -2.64 DPS) [rep] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of Arcane Mastery (22330, -6.14 DPS, sim-verified) [dungeon]; Crystalline Threaded Cape (20697, -8.49 DPS) [world]; Spritecaster Cape (11623, -9.84 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Vestments (231603, -0.44 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -6.87 DPS) [pvp]; Alanna's Embrace (13314, -11.38 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 31.2 spell_power points (16.35 DPS) | yes | Runecloth Cuffs (254123, -5.80 DPS) [crafted]; Marshal's Silk Bracers (16438, -5.86 DPS) [pvp]; Dreadmist Bracers (16703, -7.76 DPS) [dungeon] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 43.7 spell_power points (22.93 DPS) | yes | Marshal's Silk Gloves (16440, -2.33 DPS) [vendor]; Marshal's Silk Gauntlets (231608, -2.33 DPS) [vendor]; Hands of Power (13253, -4.18 DPS, sim-verified) [dungeon] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 40.9 spell_power points (21.44 DPS) | yes | Ban'thok Sash (11662, -1.22 DPS) [dungeon]; Clutch of Andros (13956, -3.68 DPS) [dungeon]; Highlander's Cloth Girdle (20047, -4.51 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 66.7 spell_power points (35.00 DPS) | yes | Marshal's Silk Leggings (231605, -1.58 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -8.44 DPS) [pvp]; Skyshroud Leggings (13170, -13.34 DPS) [dungeon] |
| feet | Sergeant Major's Dreadweave Boots (220891) | PvP reward · Alliance [vendor] | 32.7 spell_power points (17.17 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.08 DPS) [dungeon]; Omnicast Boots (11822, -0.94 DPS) [dungeon] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -1.45 DPS) [vendor]; Maiden's Circle (13001, -2.86 DPS) [world_drop]; Naglering (11669, -21.54 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -0.35 DPS) [vendor]; Maiden's Circle (13001, -1.76 DPS) [world_drop]; Songstone of Ironforge (12543, -2.04 DPS) [quest] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -3.15 DPS) [world_drop]; Shard of the Splithooves (10659, -4.29 DPS) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+2.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -2.84 DPS, sim-verified) [crafted]; Serenity Field (272439, -4.20 DPS) [vendor]; Uther's Strength (11302, -8.39 DPS) [world_drop] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Staff of Jordan (873, -14.04 DPS, sim-verified) [world_drop]; Kindling Stave (11750, -22.84 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, -0.36 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -2.26 DPS) [dungeon]; Oblivion's Touch (18761, -2.78 DPS) [dungeon] |

**New at 60:** head: Magister's Crown; neck: Diana's Pearl Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Magician's Cord; legs: Sentinel's Silk Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Channeler's Ring; finger2: Don Julio's Band; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1299, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (undead, 140005000100000000-00000000000000000-000000000000000000)

Set DPS (verified): 46.9. Weights run: 2.6s. Verify run: 2.5s. 175 eligible items had no known source.

Race sweep: undead 46.9 ± 0.1, orc 45.4 ± 0.1, troll 44.8 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.814 ± 0.013, crit=0.079 ± 0.002 per rating point (14 rating = 1%, 1.101 per %), hit=0.344 ± 0.014 per rating point (10 rating = 1%, 3.438 per %), spell_haste=0.262 ± 0.061, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.951 ± 0.003, spirit=not significant (0.051 ± 0.048), mp5=0.659 ± 0.090

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pearly Circlet (253957) | Tailoring [crafted] | 10.5 spell_power points (2.10 DPS) | yes | Pristine Circlet (253949, -0.89 DPS) [crafted]; Shadow Goggles (4373, -1.22 DPS) [crafted]; Flying Tiger Goggles (4368, -2.06 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.2 spell_power points (0.03 DPS) | yes | Scholarly Pendant (277203, -0.01 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.3 spell_power points (2.47 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.82 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.67 DPS) [crafted]; Slime-encrusted Pads (6461, -2.07 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pearl-clasped Cloak (5542, +0.00 DPS) [crafted]; Sanguine Cape (14376, -0.19 DPS) [world_drop]; Battle Healer's Cloak (20427, -0.22 DPS) [rep] |
| chest | Filigreed Pearly Gown (253909) | Tailoring [crafted] | 12.8 spell_power points (2.56 DPS) | yes | Filigreed Pristine Gown (253901, -0.71 DPS) [crafted]; Gray Woolen Robe (2585, -0.94 DPS) [crafted]; Mystic's Wrap (14369, -1.42 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 6.0 spell_power points (1.21 DPS) | yes | Bright Bracers (3647, -0.56 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.61 DPS) [quest]; Owlbeard Bracers (16981, -0.69 DPS) [quest] |
| hands | Pearly Gloves (253921) | Tailoring [crafted] | 9.1 spell_power points (1.82 DPS) | yes | Serpent Gloves (5970, -0.42 DPS) [dungeon]; Pristine Gloves (253913, -0.53 DPS) [crafted]; Gnoll Casting Gloves (892, -0.62 DPS) [world] |
| waist | Pearly Sash (253933) | Tailoring [crafted] | 9.9 spell_power points (1.99 DPS) | yes | Pristine Sash (253925, -0.53 DPS) [crafted]; Keller's Girdle (2911, -0.68 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.70 DPS) [crafted] |
| legs | Filigreed Pearly Leggings (253945) | Tailoring [crafted] | 15.5 spell_power points (3.12 DPS) | yes | Filigreed Pristine Leggings (253937, -0.89 DPS) [crafted]; Silk-threaded Trousers (1929, -1.71 DPS) [dungeon]; Darkweave Breeches (12987, -1.91 DPS) [world_drop] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.3 spell_power points (2.06 DPS) | yes | Pearly Boots (253897, -0.42 DPS) [crafted]; Pristine Boots (253889, -0.96 DPS) [crafted]; Red Woolen Boots (4313, -1.20 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 6.3 spell_power points (1.27 DPS) | yes | Loop of Sacrifice (281673, -0.45 DPS) [quest]; Volcanic Rock Ring (12053, -0.78 DPS) [world_drop]; Furlfeather Ring (286745, -1.07 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 4.9 spell_power points (0.98 DPS) | yes | Loop of Sacrifice (281673, -0.16 DPS) [quest]; Volcanic Rock Ring (12053, -0.49 DPS) [world_drop]; Furlfeather Ring (286745, -0.78 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crescent Staff (6505, +0.00 DPS) [quest]; Staff of the Blessed Seer (2271, -0.28 DPS) [dungeon]; Gnarled Necromancer's Staff (251534, -0.68 DPS) [quest] |
| off_hand | Seer's Fine Stein (7608) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellbinder Orb (15926, -0.01 DPS) [dungeon]; Tork Wrench (11855, -0.17 DPS) [quest]; Ancestral Orb (15944, -0.34 DPS) [world_drop] |
| ranged | Skycaller (12984) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sizzle Stick (8071, -1.07 DPS) [quest]; Cookie's Stirring Rod (5198, -1.08 DPS) [dungeon]; Greater Magic Wand (11288, -1.09 DPS) [crafted] |

**New at 20:** head: Pearly Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pearly Gown; wrist: Mindthrust Bracers; hands: Pearly Gloves; waist: Pearly Sash; legs: Filigreed Pearly Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Evocator's Blade; off_hand: Seer's Fine Stein; ranged: Skycaller

No-known-source sample (15 of 175, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 23173 Abomination Skin Leggings; 209618 Insignia of the Alliance

### Band 30 (undead, 141005203100301000-00000000000000000-000000000000000000)

Set DPS (verified): 92.2. Weights run: 2.9s. Verify run: 4.0s. 301 eligible items had no known source.

Race sweep: undead 92.2 ± 0.1, orc 90.8 ± 0.1, troll 89.8 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.606 ± 0.016, crit=0.107 ± 0.003 per rating point (14 rating = 1%, 1.502 per %), hit=0.423 ± 0.023 per rating point (10 rating = 1%, 4.232 per %), spell_haste=1.318 ± 0.243, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.641 ± 0.002, spirit=not significant (0.192 ± 0.160), mp5=not significant (0.672 ± 0.210)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 12.2 spell_power points (3.39 DPS) | yes | Enchanter's Cowl (4322, -0.30 DPS, sim-verified) [crafted]; Silk Headband (7050, -0.88 DPS) [crafted]; Filigreed Pearly Circlet (253983, -0.89 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.6 spell_power points (1.01 DPS) | yes | Kaleidoscope Chain (13084, -0.12 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.17 DPS) [vendor]; Crystal Starfire Medallion (5003, -0.18 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.4 spell_power points (4.03 DPS) | yes | Death Speaker Mantle (6685, -0.50 DPS) [dungeon]; Magician's Mantle (12998, -1.12 DPS) [world_drop]; Invoker's Mantle (215365, -1.23 DPS) [crafted] |
| back | Windsong Drape (15468) | Free at Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Raider's Cloak (272078, -0.10 DPS) [vendor]; Hillman's Cloak (3719, -0.21 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.28 DPS) [crafted] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 17.8 spell_power points (4.97 DPS) | yes | Green Silk Armor (7065, -0.26 DPS) [crafted]; Mechbuilder's Overalls (9508, -0.71 DPS) [dungeon]; Death Speaker Robes (6682, -1.16 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 11.0 spell_power points (3.07 DPS) | yes | Mindthrust Bracers (1974, -1.40 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -1.72 DPS) [world_drop]; Nightsky Wristbands (6407, -1.90 DPS) [world_drop] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 10.0 spell_power points (2.79 DPS) | yes | Jutebraid Gloves (10654, -0.27 DPS) [quest]; Serpent Gloves (5970, -0.84 DPS) [dungeon]; Truefaith Gloves (7049, -0.89 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.8 spell_power points (3.58 DPS) | yes | Belt of Arugal (6392, -0.29 DPS) [dungeon]; Warsong Sash (16975, -0.51 DPS) [quest]; Crimson Silk Belt (7055, -0.72 DPS) [crafted] |
| legs | Pearly Leggings (253995) | Tailoring [crafted] | 14.2 spell_power points (3.96 DPS) | yes | Pristine Leggings (253987, -0.55 DPS) [crafted]; Filigreed Pearly Leggings (253945, -0.76 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.88 DPS, sim-verified) [dungeon] |
| feet | Golden Slippers (254009) | Tailoring [crafted] | 13.3 spell_power points (3.72 DPS) | yes | Gilded Slippers (254001, -0.37 DPS) [crafted]; Acidic Walkers (9454, -0.76 DPS) [dungeon]; Spidersilk Boots (4320, -1.09 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.0 spell_power points (2.52 DPS) | yes | Ring of Power Regulation (273030, -0.28 DPS) [dungeon]; Solid Gold Band (279845, -1.14 DPS) [quest]; Electrocutioner Lagnut (9447, -1.31 DPS) [dungeon] |
| finger2 | Advisor's Ring (19521) | Warsong Outriders [rep] | 8.3 spell_power points (2.33 DPS) | yes | Ring of Power Regulation (273030, -0.09 DPS) [dungeon]; Solid Gold Band (279845, -0.94 DPS) [quest]; Electrocutioner Lagnut (9447, -1.12 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Soulstaff (249393, -1.30 DPS) [crafted]; Tendonscraper Dagger (273053, -1.33 DPS) [world]; Glimmering Staff (249392, -3.20 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.6 spell_power points (2.97 DPS) | yes | Witch's Finger (16887, -1.62 DPS) [quest]; Satyr's Rod (15962, -1.85 DPS) [world_drop]; Orb of Souls (249395, -2.17 DPS, sim-verified) [crafted] |
| ranged | Starfaller (13063) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Greater Mystic Wand (217287, -0.50 DPS) [crafted]; Wand of Eventide (5214, -0.62 DPS) [world_drop]; Thunderwood (13062, -0.99 DPS, sim-verified) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Windsong Drape; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Defiler's Cloth Girdle; legs: Pearly Leggings; feet: Golden Slippers; finger1: Snake Hoop; finger2: Advisor's Ring; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Starfaller

No-known-source sample (15 of 301, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (undead, 141005213100301531-00000000000000000-000000000000000000)

Set DPS (verified): 173.3. Weights run: 3.0s. Verify run: 3.2s. 392 eligible items had no known source.

Race sweep: undead 173.3 ± 0.3, orc 171.6 ± 0.3, troll 170.4 ± 0.3 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.553 ± 0.025, crit=0.270 ± 0.006 per rating point (14 rating = 1%, 3.779 per %), hit=0.632 ± 0.034 per rating point (10 rating = 1%, 6.323 per %), spell_haste=2.432 ± 0.406, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.638 ± 0.002, spirit=not significant (0.651 ± 0.231), mp5=not significant (0.797 ± 0.255)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 23.3 spell_power points (7.55 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.73 DPS) [crafted]; Whitemane's Chapeau (7720, -0.95 DPS) [dungeon]; Electromagnetic Gigaflux Reactivator (9492, -1.62 DPS) [dungeon] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 12.6 spell_power points (4.07 DPS) | yes | Glowing Eye of Mordresh (10769, -0.37 DPS) [dungeon]; Necklace of Calisea (1714, -1.34 DPS) [dungeon]; Amberglow Talisman (10824, -1.96 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Green Silken Shoulders (7057, -0.03 DPS) [crafted]; Bloodmage Mantle (7684, -0.07 DPS) [dungeon]; Mantle of Doan (7712, -0.71 DPS, sim-verified) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battle Healer's Cloak (19528, -0.54 DPS) [rep]; Long Silken Cloak (4326, -0.69 DPS) [crafted]; Guardian Cloak (5965, -0.69 DPS) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 32.1 spell_power points (10.42 DPS) | yes | Robe of the Magi (1716, -1.14 DPS) [world_drop]; Robe of Power (7054, -2.03 DPS) [crafted]; Elemental Raiment (9434, -3.60 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 11.4 spell_power points (3.70 DPS) | yes | Radiant Silver Bracers (4545, -1.23 DPS, sim-verified) [quest]; Earthen Silk Cuffs (254019, -1.34 DPS) [crafted]; Condor Bracers (15864, -1.43 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 24.8 spell_power points (8.04 DPS) | yes | Red Mageweave Gloves (10018, -2.00 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -3.17 DPS) [crafted]; Golden Handwraps (254029, -3.88 DPS) [crafted] |
| waist | Golden Cord (254045) | Tailoring [crafted] | 17.3 spell_power points (5.60 DPS) | yes | Gilded Cord (254037, -0.30 DPS) [crafted]; Defiler's Cloth Girdle (20166, -0.34 DPS) [rep]; Star Belt (4329, -0.54 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 20.6 spell_power points (6.70 DPS) | yes | Gaze Dreamer Pants (6903, -0.90 DPS) [dungeon]; Black Mageweave Leggings (9999, -1.15 DPS) [crafted]; Stoneweaver Leggings (9407, -1.85 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 28.6 spell_power points (9.27 DPS) | yes | Furen's Boots (13100, -4.63 DPS) [world_drop]; Gilded Slippers (254001, -4.89 DPS) [crafted]; Golden Slippers (254009, -5.10 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.8 spell_power points (4.14 DPS) | yes | Snake Hoop (6750, -1.30 DPS) [quest]; Ring of Power Regulation (273030, -1.42 DPS) [dungeon]; Electrocutioner Lagnut (9447, -1.69 DPS) [dungeon] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 11.4 spell_power points (3.70 DPS) | yes | Snake Hoop (6750, -0.85 DPS) [quest]; Ring of Power Regulation (273030, -0.97 DPS) [dungeon]; Electrocutioner Lagnut (9447, -1.25 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Celestial Stave (9517, -11.89 DPS) [quest]; Hypnotic Blade (7714, -11.91 DPS) [dungeon]; Illusionary Rod (7713, -12.15 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Nether Force Wand (11263) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Twisted Nether Wand (249144, -0.54 DPS) [crafted]; Starfaller (13063, -0.56 DPS) [world_drop]; Dancing Flame (6806, -1.15 DPS, sim-verified) [quest] |

**New at 40:** head: Papal Fez; neck: Triune Amulet; shoulder: Inquisitor's Shawl; back: Icy Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Golden Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Staff of Jordan; ranged: Nether Force Wand

No-known-source sample (15 of 392, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (undead, 141005213100301531-40000000000000000-033000000000000000)

Set DPS (verified): 260.3. Weights run: 3.0s. Verify run: 5.4s. 504 eligible items had no known source.

Race sweep: undead 260.3 ± 0.4, orc 257.6 ± 0.4, troll 256.4 ± 0.4 (dps); the guide race undead is best.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.574 ± 0.025, crit=0.347 ± 0.008 per rating point (14 rating = 1%, 4.856 per %), hit=0.760 ± 0.038 per rating point (10 rating = 1%, 7.603 per %), spell_haste=3.201 ± 0.402, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.636 ± 0.002, spirit=not significant (0.283 ± 0.192), mp5=not significant (0.503 ± 0.228)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreamweave Circlet (10041) | Tailoring [crafted] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Red Mageweave Headband (10033, -0.90 DPS, sim-verified) [crafted]; Spellpower Goggles Xtreme Plus (15999, -1.17 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -1.64 DPS) [vendor] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 19.4 spell_power points (7.26 DPS) | yes | Glowing Eye of Mordresh (10769, -2.20 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -3.21 DPS) [quest]; Triune Amulet (7722, -3.54 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kentic Amice (11624, -0.06 DPS) [dungeon]; Blood Guard's Dreadweave Mantle (220905, -1.98 DPS) [vendor]; Nethershine Shoulders (254057, -2.33 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -0.74 DPS) [quest]; Runecloth Cloak (13860, -2.07 DPS) [crafted]; Big Voodoo Cloak (8216, -3.35 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 30.5 spell_power points (11.39 DPS) | yes | Dreamweave Vest (10021, -1.25 DPS) [crafted]; Robe of the Magi (1716, -1.35 DPS) [world_drop]; Runecloth Tunic (13857, -2.68 DPS) [crafted] |
| wrist | Nethershine Cuffs (254069) | Tailoring [crafted] | 13.7 spell_power points (5.12 DPS) | yes | Nethergeld Cuffs (254061, -0.47 DPS) [crafted]; Bloodband Bracers (11469, -1.00 DPS) [quest]; Spidertank Oilrag (9448, -1.19 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.3 spell_power points (8.32 DPS) | yes | Golden Gloves (254103, -0.82 DPS) [crafted]; Runecloth Gloves (13863, -0.96 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.54 DPS) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Golden Waistcord (254089, -0.93 DPS) [crafted]; Ghostweave Cord (254073, -0.98 DPS) [crafted] |
| legs | Stone Guard's Dreadweave Leggings (220906) | Lady Palanseer [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshock Leggings (9484, +0.00 DPS) [dungeon]; Senior Designer's Pantaloons (11841, -0.98 DPS) [dungeon]; Red Mageweave Pants (10009, -1.07 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 26.0 spell_power points (9.71 DPS) | yes | First Sergeant's Dreadweave Boots (220909, -2.74 DPS, sim-verified) [vendor]; Golden Sandals (254115, -3.03 DPS) [crafted]; Gilded Sandals (254107, -3.67 DPS) [crafted] |
| finger1 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 14.1 spell_power points (5.29 DPS) | yes | Band of the Unicorn (7553, -0.43 DPS) [world_drop]; Philanthropist's Ring (281635, -0.48 DPS) [quest]; Brainlash (6440, -1.54 DPS) [dungeon] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 14.0 spell_power points (5.24 DPS) | yes | Band of the Unicorn (7553, -0.38 DPS) [world_drop]; Philanthropist's Ring (281635, -0.43 DPS) [quest]; Brainlash (6440, -1.49 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Evonice's Landin' Pilla (18951, -0.21 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -0.32 DPS) [world_drop]; Rune of the Guard Captain (19120, -0.57 DPS) [quest]; Darkspear Voodoo Seal (272061, -1.21 DPS, sim-verified) [vendor] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -4.04 DPS, sim-verified) [dungeon]; Glowing Brightwood Staff (812, -6.17 DPS) [world_drop]; Spellshifter Rod (9527, -6.92 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Noxious Shooter (17745, -0.73 DPS) [dungeon]; Dreambough Wand (249234, -0.98 DPS) [crafted]; Nether Force Wand (11263, -0.98 DPS) [quest] |

**New at 50:** head: Dreamweave Circlet; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethershine Cuffs; waist: Satyrmane Sash; legs: Stone Guard's Dreadweave Leggings; finger1: Cyclopean Band; finger2: Advisor's Ring; trinket1: Ankh of Life; trinket2: Frozen Heart of the Mountain; ranged: Lesser Eternal Wand

No-known-source sample (15 of 504, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (undead, 141005213100301531-40000000000000000-053500030000000000)

Set DPS (verified): 410.8. Weights run: 3.1s. Verify run: 18.3s. 1293 eligible items had no known source.

Race sweep: undead 410.8 ± 0.6, orc 406.5 ± 0.6, troll 403.6 ± 0.6 (dps); the guide race undead is best.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.843 ± 0.042, crit=0.572 ± 0.014 per rating point (14 rating = 1%, 8.010 per %), hit=1.114 ± 0.071 per rating point (10 rating = 1%, 11.139 per %), spell_haste=6.611 ± 0.874, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.613 ± 0.002, spirit=not significant (-0.003 ± 0.431), mp5=not significant (0.110 ± 0.461)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crimson Felt Hat (18727, +0.00 DPS) [dungeon]; Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp] |
| neck | Diana's Pearl Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 26.9 spell_power points (10.50 DPS) | yes | Arcane Crystal Pendant (20037, -2.27 DPS) [quest]; Archlight Talisman (15856, -2.52 DPS) [quest]; Star of Mystaria (12103, -3.18 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 36.0 spell_power points (14.04 DPS) | yes | Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Darkspear Shoulderpads (272103, -0.47 DPS) [vendor]; Burial Shawl (18681, -0.97 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -3.63 DPS, sim-verified) [world]; Shroud of Arcane Mastery (22330, -5.26 DPS) [dungeon]; Deep Woodlands Cloak (19121, -5.58 DPS) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Raiment (231596, -1.09 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -5.77 DPS) [pvp]; Alanna's Embrace (13314, -6.09 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 28.7 spell_power points (11.22 DPS) | yes | Runecloth Cuffs (254123, -3.64 DPS) [crafted]; General's Silk Cuffs (16538, -4.97 DPS) [pvp]; Magiskull Cuffs (13107, -5.63 DPS) [world_drop] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 33.5 spell_power points (13.06 DPS) | yes | Hands of Power (13253, +0.00 DPS) [dungeon]; General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 39.2 spell_power points (15.32 DPS) | yes | Ban'thok Sash (11662, -2.66 DPS) [dungeon]; Clutch of Andros (13956, -4.71 DPS) [dungeon]; Defiler's Cloth Girdle (20163, -4.75 DPS) [rep] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 44.0 spell_power points (17.19 DPS) | yes | General's Silk Trousers (231595, +0.00 DPS) [pvp]; Legionnaire's Silk Legguards (227107, -0.26 DPS) [pvp]; Skyshroud Leggings (13170, -4.91 DPS, sim-verified) [dungeon] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dragonrider Boots (18102, +0.00 DPS) [dungeon]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Blood Guard's Silk Walkers (227110, -0.27 DPS) [pvp] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, +0.00 DPS) [quest]; Blessed Band of Light (272407, -1.08 DPS) [vendor]; Naglering (11669, -10.17 DPS, sim-verified) [dungeon] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Blessed Band of Light (272407, -0.71 DPS) [vendor]; Deep Rooted Ring (19109, -0.98 DPS) [world]; Eye of Orgrimmar (12545, -1.24 DPS, sim-verified) [quest] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+9.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -0.77 DPS) [crafted]; Rune of the Guard Captain (19120, -1.64 DPS) [quest]; Uther's Strength (11302, -2.34 DPS) [world_drop] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -3.06 DPS, sim-verified) [vendor]; Frozen Heart of the Mountain (249469, -4.67 DPS) [crafted]; Rune of the Guard Captain (19120, -5.54 DPS) [quest] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Mindfang (20214, -5.54 DPS) [rep]; Ironbark Staff (20220, -6.74 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (11.32 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Tome of Mending (234564, -0.86 DPS) [pvp]; Tome of the Lost (22253, -2.72 DPS, sim-verified) [dungeon] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, -0.03 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -1.28 DPS) [dungeon]; Oblivion's Touch (18761, -2.03 DPS) [dungeon] |

**New at 60:** head: Magister's Crown; neck: Diana's Pearl Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Magician's Cord; legs: Outrider's Silk Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Channeler's Ring; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Brilliant Wand

No-known-source sample (15 of 1293, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60, raid preset (undead, 131205203100311531-52500001200100000-003000000000000000)

Set DPS (verified): 788.7. Weights run: 3.7s. Verify run: 21.5s. 1293 eligible items had no known source.

Race sweep: undead 788.7 ± 0.9, orc 785.8 ± 0.9, troll 777.2 ± 0.9 (dps); the guide race undead is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.911 ± 0.048, crit=0.914 ± 0.038 per rating point (14 rating = 1%, 12.801 per %), hit=1.653 ± 0.108 per rating point (10 rating = 1%, 16.532 per %), spell_haste=4.035 ± 0.657, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.569 ± 0.002, spirit=not significant (0.267 ± 0.372), mp5=not significant (0.637 ± 0.432)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Cap of the Scarlet Savant (12752) | Cap of the Scarlet Savant [quest] | 43.8 spell_power points (22.99 DPS) | yes | Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; The Postmaster's Band (13390, -2.29 DPS) [dungeon] |
| neck | Diana's Pearl Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 32.8 spell_power points (17.22 DPS) | yes | Star of Mystaria (12103, -3.12 DPS) [dungeon]; Arcane Crystal Pendant (20037, -5.96 DPS) [quest]; Archlight Talisman (15856, -6.14 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 41.6 spell_power points (21.85 DPS) | yes | Warlord's Silk Amice (231594, -0.86 DPS) [pvp]; Champion's Silk Mantle (227104, -1.44 DPS) [pvp]; Darkspear Shoulderpads (272103, -2.64 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -8.49 DPS) [world]; Shroud of Arcane Mastery (22330, -8.82 DPS, sim-verified) [dungeon]; Spritecaster Cape (11623, -9.84 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Raiment (231596, -0.44 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -6.87 DPS) [pvp]; Alanna's Embrace (13314, -9.15 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 31.2 spell_power points (16.35 DPS) | yes | Runecloth Cuffs (254123, -5.80 DPS) [crafted]; General's Silk Cuffs (16538, -5.86 DPS) [pvp]; Dreadmist Bracers (16703, -7.76 DPS) [dungeon] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 43.7 spell_power points (22.93 DPS) | yes | General's Silk Handguards (16540, -2.33 DPS) [vendor]; General's Silk Gauntlets (231599, -2.33 DPS) [vendor]; Hands of Power (13253, -9.01 DPS, sim-verified) [dungeon] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 40.9 spell_power points (21.44 DPS) | yes | Ban'thok Sash (11662, -1.22 DPS) [dungeon]; Clutch of Andros (13956, -3.68 DPS) [dungeon]; Defiler's Cloth Girdle (20163, -4.51 DPS) [rep] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 48.0 spell_power points (25.17 DPS) | yes | Legionnaire's Silk Legguards (227107, +0.00 DPS) [pvp]; General's Silk Trousers (231595, +0.00 DPS) [pvp]; Skyshroud Leggings (13170, -3.51 DPS) [dungeon] |
| feet | First Sergeant's Dreadweave Boots (220909) | PvP reward · Horde [vendor] | 32.7 spell_power points (17.17 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.08 DPS) [dungeon]; Omnicast Boots (11822, -0.94 DPS) [dungeon] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -1.45 DPS) [vendor]; Maiden's Circle (13001, -2.86 DPS) [world_drop]; Naglering (11669, -10.82 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -0.35 DPS) [vendor]; Maiden's Circle (13001, -1.76 DPS) [world_drop]; Eye of Orgrimmar (12545, -2.04 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burst of Knowledge (11832, -1.51 DPS) [dungeon]; Rune of the Guard Captain (19120, -1.73 DPS) [quest]; Uther's Strength (11302, -4.66 DPS) [world_drop] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -4.38 DPS, sim-verified) [vendor]; Burst of Knowledge (11832, -5.25 DPS) [dungeon]; Rune of the Guard Captain (19120, -5.47 DPS) [quest] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Staff of Jordan (873, -9.64 DPS, sim-verified) [world_drop]; Whiteout Staff (19101, -20.30 DPS) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, -0.36 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -2.26 DPS) [dungeon]; Oblivion's Touch (18761, -2.78 DPS) [dungeon] |

**New at 60:** head: Cap of the Scarlet Savant; neck: Diana's Pearl Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Magician's Cord; legs: Outrider's Silk Leggings; feet: First Sergeant's Dreadweave Boots; finger1: Channeler's Ring; finger2: Don Julio's Band; trinket1: Frozen Heart of the Mountain; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1293, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

