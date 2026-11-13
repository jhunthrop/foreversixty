# Leveling BiS: Arcane

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 140005000100000000-00000000000000000-000000000000000000)

Set DPS (verified): 42.7. Weights run: 2.8s. Verify run: 2.7s. 198 eligible items had no known source.

Race sweep: gnome 42.7 ± 0.1, high-order-skyborne 41.3 ± 0.0, human 41.3 ± 0.0 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.927 ± 0.011, crit=0.086 ± 0.002 per rating point (14 rating = 1%, 1.208 per %), hit=0.306 ± 0.012 per rating point (10 rating = 1%, 3.056 per %), spell_haste=not significant (0.277 ± 0.104), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.870 ± 0.003, spirit=not significant (0.145 ± 0.064), mp5=0.531 ± 0.128

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pearly Circlet (253957) | Tailoring [crafted] | 9.6 spell_power points (1.55 DPS) | yes | Pristine Circlet (253949, -0.65 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.66 DPS) [crafted]; Flying Tiger Goggles (4368, -1.45 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.3 spell_power points (0.05 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.3 spell_power points (2.16 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.94 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.51 DPS) [crafted]; Slime-encrusted Pads (6461, -1.90 DPS) [dungeon] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.03 DPS) [crafted]; Sanguine Cape (14376, -0.17 DPS) [world_drop]; Caretaker's Cape (20428, -0.24 DPS) [rep] |
| chest | Filigreed Pearly Gown (253909) | Tailoring [crafted] | 12.9 spell_power points (2.09 DPS) | yes | Filigreed Pristine Gown (253901, -0.51 DPS, sim-verified) [crafted]; Gray Woolen Robe (2585, -0.69 DPS) [crafted]; Mystic's Wrap (14369, -1.04 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 6.2 spell_power points (1.01 DPS) | yes | Repurposed Hair Band (281256, -0.66 DPS) [quest]; Mystic's Bracelets (14366, -0.71 DPS) [world_drop]; Bright Bracers (3647, -0.75 DPS, sim-verified) [world_drop] |
| hands | Pearly Gloves (253921) | Tailoring [crafted] | 8.9 spell_power points (1.43 DPS) | yes | Pristine Gloves (253913, -0.34 DPS) [crafted]; Gnoll Casting Gloves (892, -0.46 DPS) [world]; Serpent Gloves (5970, -0.50 DPS, sim-verified) [dungeon] |
| waist | Pearly Sash (253933) | Tailoring [crafted] | 9.8 spell_power points (1.58 DPS) | yes | Pristine Sash (253925, -0.38 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.39 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.49 DPS) [crafted] |
| legs | Filigreed Pearly Leggings (253945) | Tailoring [crafted] | 15.7 spell_power points (2.54 DPS) | yes | Filigreed Pristine Leggings (253937, -0.65 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -1.35 DPS) [world_drop]; Silk-threaded Trousers (1929, -1.41 DPS) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.7 spell_power points (1.73 DPS) | yes | Pearly Boots (253897, -0.52 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.80 DPS) [crafted]; Red Woolen Boots (4313, -0.97 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.9 spell_power points (1.11 DPS) | yes | Lavishly Jeweled Ring (1156, -0.21 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.66 DPS) [world_drop]; Furlfeather Ring (286745, -0.95 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 6.1 spell_power points (0.98 DPS) | yes | Lavishly Jeweled Ring (1156, -0.37 DPS, sim-verified) [dungeon]; Volcanic Rock Ring (12053, -0.53 DPS) [world_drop]; Furlfeather Ring (286745, -0.82 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | 30.6 spell_power points (4.94 DPS) | yes | Staff of the Blessed Seer (2271, -1.04 DPS, sim-verified) [dungeon]; Emberstone Staff (5201, -1.06 DPS) [dungeon]; Buzzer Blade (2169, -1.39 DPS) [dungeon] |
| off_hand | Mathystra Channeling Orb (281291) | Baron Marinous [quest] | 4.6 spell_power points (0.75 DPS) | yes | Tear of Grief (5611, -0.07 DPS) [quest]; Dwarven Tome (279898, -0.24 DPS) [quest]; Seer's Fine Stein (7608, -0.28 DPS) [world_drop] |
| ranged | Skycaller (12984) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cookie's Stirring Rod (5198, -0.98 DPS, sim-verified) [dungeon]; Sizzle Stick (8071, -1.47 DPS) [quest]; Greater Magic Wand (11288, -1.51 DPS) [crafted] |

**New at 20:** head: Pearly Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Pearly Gown; wrist: Mindthrust Bracers; hands: Pearly Gloves; waist: Pearly Sash; legs: Filigreed Pearly Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Evocator's Blade; off_hand: Mathystra Channeling Orb; ranged: Skycaller

No-known-source sample (15 of 198, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 8350 The 1 Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (gnome, 141005203100301000-00000000000000000-000000000000000000)

Set DPS (verified): 90.6. Weights run: 2.9s. Verify run: 3.7s. 326 eligible items had no known source.

Race sweep: gnome 90.6 ± 0.1, human 89.1 ± 0.1, high-order-skyborne 88.6 ± 0.1 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.499 ± 0.009, crit=0.072 ± 0.002 per rating point (14 rating = 1%, 1.007 per %), hit=0.236 ± 0.012 per rating point (10 rating = 1%, 2.363 per %), spell_haste=0.952 ± 0.142, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.563 ± 0.003, spirit=not significant (0.120 ± 0.087), mp5=not significant (0.455 ± 0.114)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.7 spell_power points (2.62 DPS) | yes | Enchanter's Cowl (4322, -0.16 DPS) [crafted]; Silk Headband (7050, -0.61 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.83 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.0 spell_power points (0.67 DPS) | yes | Seal of the Expedition (271920, -0.11 DPS) [rep]; Kaleidoscope Chain (13084, -0.12 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.14 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.5 spell_power points (3.01 DPS) | yes | Death Speaker Mantle (6685, -0.45 DPS) [dungeon]; Invoker's Mantle (215365, -0.89 DPS) [crafted]; Magician's Mantle (12998, -0.89 DPS) [world_drop] |
| back | Vine Pruner's Cloak (279835) | A Green Sample [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Prelacy Cape (7004, -0.06 DPS) [quest]; Heavy Woolen Cloak (4311, -0.12 DPS) [crafted] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 15.9 spell_power points (3.56 DPS) | yes | Green Silk Armor (7065, -0.10 DPS) [crafted]; Mechbuilder's Overalls (9508, -0.44 DPS) [dungeon]; Tree Bark Jacket (1486, -0.66 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 10.4 spell_power points (2.32 DPS) | yes | Mindthrust Bracers (1974, -1.45 DPS) [dungeon]; Glowing Magical Bracelets (13106, -1.54 DPS, sim-verified) [world_drop]; Nightsky Wristbands (6407, -1.57 DPS) [world_drop] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 8.6 spell_power points (1.91 DPS) | yes | Serpent Gloves (5970, -0.35 DPS) [dungeon]; Truefaith Gloves (7049, -0.46 DPS) [crafted]; Shilly Mitts (9609, -0.62 DPS, sim-verified) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.5 spell_power points (2.79 DPS) | yes | Belt of Arugal (6392, -0.31 DPS) [dungeon]; Invoker's Cord (215366, -0.67 DPS) [crafted]; Crimson Silk Belt (7055, -0.67 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 13.1 spell_power points (2.92 DPS) | yes | Pearly Leggings (253995, -0.25 DPS) [crafted]; Pristine Leggings (253987, -0.44 DPS) [crafted]; Filigreed Pearly Leggings (253945, -0.76 DPS) [crafted] |
| feet | Golden Slippers (254009) | Tailoring [crafted] | 11.3 spell_power points (2.52 DPS) | yes | Gilded Slippers (254001, -0.07 DPS) [crafted]; Acidic Walkers (9454, -0.41 DPS) [dungeon]; Spidersilk Boots (4320, -0.51 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 10.3 spell_power points (2.31 DPS) | yes | Ring of Power Regulation (273030, -0.66 DPS) [dungeon]; Runemender's Seal (273031, -0.78 DPS) [world]; Minor Channeling Ring (1449, -0.97 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.9 spell_power points (1.77 DPS) | yes | Ring of Power Regulation (273030, -0.12 DPS) [dungeon]; Runemender's Seal (273031, -0.24 DPS) [world]; Minor Channeling Ring (1449, -0.43 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.00 DPS) [dungeon]; Tendonscraper Dagger (273053, -1.00 DPS) [world]; Glimmering Staff (249392, -2.94 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.0 spell_power points (2.23 DPS) | yes | Orb of Souls (249395, -1.18 DPS) [crafted]; Arcane Infused Rod (279838, -1.20 DPS, sim-verified) [quest]; Eye of Paleth (2943, -1.34 DPS) [quest] |
| ranged | Starfaller (13063) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderwood (13062, -0.25 DPS) [world_drop]; Greater Mystic Wand (217287, -0.35 DPS) [crafted]; Wand of Eventide (5214, -0.57 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Golden Slippers; finger1: Snake Hoop; finger2: Lorekeeper's Ring; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Starfaller

No-known-source sample (15 of 326, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 141005213100301531-00000000000000000-000000000000000000)

Set DPS (verified): 167.6. Weights run: 3.0s. Verify run: 3.1s. 456 eligible items had no known source.

Race sweep: gnome 167.6 ± 0.3, human 166.9 ± 0.3, high-order-skyborne 165.2 ± 0.3 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.587 ± 0.022, crit=0.261 ± 0.006 per rating point (14 rating = 1%, 3.650 per %), hit=0.604 ± 0.034 per rating point (10 rating = 1%, 6.039 per %), spell_haste=2.133 ± 0.376, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.628 ± 0.002, spirit=not significant (0.277 ± 0.203), mp5=not significant (0.370 ± 0.231)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Miner's Hat of the Deep (9429) | Uldaman [dungeon] | 22.0 spell_power points (7.02 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.31 DPS) [crafted]; Papal Fez (9431, -0.48 DPS) [dungeon]; Corpseshroud (10574, -1.22 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.7 spell_power points (3.74 DPS) | yes | Triune Amulet (7722, -0.56 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.86 DPS) [quest]; Necklace of Calisea (1714, -1.80 DPS) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.6 spell_power points (4.68 DPS) | yes | Green Silken Shoulders (7057, -0.06 DPS) [crafted]; Bloodmage Mantle (7684, -0.11 DPS) [dungeon]; Berylline Pads (4197, -0.56 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.09 DPS) [crafted]; Expeditionary's Cape (271922, -0.44 DPS) [rep] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 27.2 spell_power points (8.68 DPS) | yes | Robe of the Magi (1716, -0.08 DPS) [world_drop]; Robe of Power (7054, -1.25 DPS) [crafted]; Elemental Raiment (9434, -1.97 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 10.1 spell_power points (3.23 DPS) | yes | Arcane Runed Bracers (4744, -0.35 DPS) [quest]; Condor Bracers (15864, -0.99 DPS) [quest]; Earthen Silk Cuffs (254019, -1.51 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.3 spell_power points (7.12 DPS) | yes | Red Mageweave Gloves (10018, -2.25 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -2.33 DPS) [crafted]; Golden Handwraps (254029, -3.00 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 16.3 spell_power points (5.22 DPS) | yes | Deathmage Sash (10771, -0.17 DPS) [dungeon]; Golden Cord (254045, -0.38 DPS) [crafted]; Gilded Cord (254037, -0.63 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.0 spell_power points (6.72 DPS) | yes | Crimson Silk Pantaloons (7062, -1.80 DPS, sim-verified) [crafted]; Stoneweaver Leggings (9407, -1.92 DPS) [dungeon]; Gaze Dreamer Pants (6903, -2.09 DPS) [dungeon] |
| feet | Golden Slippers (254009) | Tailoring [crafted] | 13.4 spell_power points (4.28 DPS) | yes | Gilded Slippers (254001, -0.37 DPS) [crafted]; Acidic Walkers (9454, -0.82 DPS) [dungeon]; Spidersilk Boots (4320, -1.29 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.9 spell_power points (4.13 DPS) | yes | Lorekeeper's Ring (19524, -0.90 DPS) [rep]; Runemender's Seal (273031, -1.43 DPS) [world]; Ring of Forlorn Spirits (2043, -1.58 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 12.1 spell_power points (3.85 DPS) | yes | Lorekeeper's Ring (19524, -0.62 DPS) [rep]; Runemender's Seal (273031, -1.15 DPS) [world]; Ring of Forlorn Spirits (2043, -1.29 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+2.6 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -9.74 DPS, sim-verified) [dungeon]; Celestial Stave (9517, -10.50 DPS) [quest]; Hypnotic Blade (7714, -10.88 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Nether Force Wand (11263) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plaguerot Sprig (10766, -0.33 DPS) [dungeon]; Twisted Nether Wand (249144, -0.65 DPS) [crafted]; Starfaller (13063, -0.68 DPS) [world_drop] |

**New at 40:** head: Miner's Hat of the Deep; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; finger1: Philanthropist's Ring; finger2: Snake Hoop; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Staff of Jordan; ranged: Nether Force Wand

No-known-source sample (15 of 456, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (gnome, 141005213100311531-30000000000000000-033000000000000000)

Set DPS (verified): 277.3. Weights run: 3.1s. Verify run: 7.3s. 598 eligible items had no known source.

Race sweep: gnome 277.3 ± 0.4, human 276.4 ± 0.4, high-order-skyborne 273.0 ± 0.4 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.838 ± 0.031, crit=0.407 ± 0.011 per rating point (14 rating = 1%, 5.695 per %), hit=0.811 ± 0.055 per rating point (10 rating = 1%, 8.112 per %), spell_haste=3.087 ± 0.533, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.805 ± 0.002, spirit=not significant (0.136 ± 0.278), mp5=not significant (0.275 ± 0.311)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 35.8 spell_power points (12.72 DPS) | yes | Dreamweave Circlet (10041, -1.69 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -2.14 DPS) [vendor]; Chief Architect's Monocle (11839, -2.53 DPS) [dungeon] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 21.0 spell_power points (7.48 DPS) | yes | Talvash's Enhancing Necklace (7673, -2.91 DPS) [quest]; Glowing Eye of Mordresh (10769, -2.94 DPS, sim-verified) [dungeon]; Gemshard Heart (17707, -3.02 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 28.1 spell_power points (9.99 DPS) | yes | Kentic Amice (11624, -0.84 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Mantle (220887, -2.44 DPS) [vendor]; Nethershine Shoulders (254057, -2.45 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Big Voodoo Cloak (8216, -2.60 DPS) [crafted]; Runecloth Cloak (13860, -2.63 DPS, sim-verified) [crafted]; Expeditionary's Cape (271921, -2.64 DPS) [rep] |
| chest | Robes of Insight (940) | World drop [world_drop] | 31.1 spell_power points (11.06 DPS) | yes | Robe of the Magi (1716, -1.21 DPS) [world_drop]; Dreamweave Vest (10021, -1.30 DPS) [crafted]; Knight's Dreadweave Vest (220886, -1.49 DPS) [vendor] |
| wrist | Nethershine Cuffs (254069) | Tailoring [crafted] | 17.0 spell_power points (6.05 DPS) | yes | Aristocratic Cuffs (12546, -0.16 DPS) [dungeon]; Nethergeld Cuffs (254061, -1.23 DPS) [crafted]; Bloodband Bracers (11469, -1.44 DPS) [quest] |
| hands | Golden Gloves (254103) | Tailoring [crafted] | 24.8 spell_power points (8.81 DPS) | yes | Dreamweave Gloves (10019, -0.88 DPS) [crafted]; Runecloth Gloves (13863, -1.42 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.50 DPS) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 28.5 spell_power points (10.13 DPS) | yes | Satyrmane Sash (17755, -2.17 DPS) [dungeon]; Dawnspire Cord (12466, -2.33 DPS) [dungeon]; Golden Waistcord (254089, -2.96 DPS, sim-verified) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | 31.4 spell_power points (11.16 DPS) | yes | Knight's Dreadweave Leggings (220888, -1.29 DPS) [vendor]; Red Mageweave Pants (10009, -2.60 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -3.85 DPS) [dungeon] |
| feet | Sergeant Major's Dreadweave Boots (220891) | PvP rank 5 · Sergeant Major · Alliance [vendor] | 23.7 spell_power points (8.41 DPS) | yes | Golden Sandals (254115, -0.00 DPS) [crafted]; Gilded Sandals (254107, -1.82 DPS) [crafted]; Southsea Mojo Boots (20641, -2.29 DPS) [quest] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 20.1 spell_power points (7.13 DPS) | yes | Woodseed Hoop (17768, -2.02 DPS) [quest]; Philanthropist's Ring (281635, -2.09 DPS) [quest]; Mindseye Circle (10634, -2.11 DPS) [dungeon] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.4 spell_power points (5.48 DPS) | yes | Woodseed Hoop (17768, -0.37 DPS) [quest]; Philanthropist's Ring (281635, -0.43 DPS) [quest]; Mindseye Circle (10634, -0.46 DPS) [dungeon] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Enriched Seal (271903, +0.00 DPS) [rep]; Ankh of Life (1713, -0.85 DPS, sim-verified) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.46 DPS) [world_drop]; Ankh of Life (1713, -2.01 DPS) [world_drop]; Enriched Seal (271903, -2.11 DPS) [rep] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-verified (+8.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Staff of Jordan (873, -8.13 DPS, sim-verified) [world_drop] |
| off_hand | Truesilver Conduit (249455) | Enchanting [crafted] | 16.7 spell_power points (5.94 DPS) | yes | Celestial Orb (7515, -0.42 DPS) [quest]; Orb of Lorica (11262, -0.83 DPS) [quest]; Orb of the Forgotten Seer (7685, -0.96 DPS) [dungeon] |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nether Force Wand (11263, -0.44 DPS) [quest]; Noxious Shooter (17745, -0.53 DPS) [dungeon]; Charged Lightning Rod (11860, -0.94 DPS) [quest] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Robes of Insight; wrist: Nethershine Cuffs; hands: Golden Gloves; waist: Ban'thok Sash; legs: Spellshock Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Thunderbrew's Boot Flask; trinket2: Frozen Heart of the Mountain; main_hand: Blade of Eternal Darkness; off_hand: Truesilver Conduit; ranged: Lesser Eternal Wand

No-known-source sample (15 of 598, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (gnome, 141005213100311531-30000000000000000-053500030000000000)

Set DPS (verified): 412.4. Weights run: 3.3s. Verify run: 19.4s. 1456 eligible items had no known source.

Race sweep: gnome 412.4 ± 0.6, human 411.8 ± 0.6, high-order-skyborne 407.2 ± 0.6 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.206 ± 0.050, crit=0.611 ± 0.017 per rating point (14 rating = 1%, 8.549 per %), hit=1.260 ± 0.084 per rating point (10 rating = 1%, 12.596 per %), spell_haste=not significant (3.613 ± 0.939), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.751 ± 0.003, spirit=not significant (0.683 ± 0.479), mp5=not significant (1.037 ± 0.486)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | The Postmaster's Band (13390) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellweaver's Turban (22267, +0.00 DPS) [dungeon]; Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jeweled Amulet of Cainwyn (1443, +0.00 DPS) [world_drop]; Nacreous Shell Necklace (22403, -0.96 DPS) [dungeon]; Star of Mystaria (12103, -1.56 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Theramore Shoulderpads (271953, +0.00 DPS) [rep]; Shroud of the Nathrezim (18720, -0.33 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, -3.41 DPS) [quest]; Shroud of Arcane Mastery (22330, -4.48 DPS) [dungeon]; Spritecaster Cape (11623, -4.68 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Vestments (231603, +0.00 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -3.46 DPS) [pvp]; Magister's Robes (16688, -4.27 DPS) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 36.4 spell_power points (13.19 DPS) | yes | Marshal's Silk Bracers (16438, -2.42 DPS) [pvp]; Manacle Cuffs (11962, -4.11 DPS) [quest]; Runecloth Cuffs (254123, -4.84 DPS) [crafted] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Marshal's Silk Handwraps (231609, +0.00 DPS) [vendor] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 47.9 spell_power points (17.36 DPS) | yes | Marshal's Silk Sash (16439, -3.61 DPS) [pvp]; Ban'thok Sash (11662, -4.09 DPS) [dungeon]; Clutch of Andros (13956, -4.50 DPS) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 61.2 spell_power points (22.14 DPS) | yes | Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -2.79 DPS) [pvp]; Magister's Leggings (16687, -3.77 DPS) [dungeon] |
| feet | The Postmaster's Treads (13391) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dragonrider Boots (18102, +0.00 DPS) [dungeon]; Kayser's Boots of Precision (22231, +0.00 DPS) [dungeon]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp] |
| finger1 | The Postmaster's Seal (13392) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Maiden's Circle (13001, +0.00 DPS) [world_drop]; Seal of Rivendare (13345, +0.00 DPS) [dungeon]; Blessed Band of Light (272407, +0.00 DPS) [vendor] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Maiden's Circle (13001, -1.59 DPS) [world_drop]; Seal of Rivendare (13345, -2.33 DPS) [dungeon]; Blessed Band of Light (272407, -2.49 DPS) [vendor] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Ankh of Life (1713, -0.04 DPS) [world_drop] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, -1.96 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -2.41 DPS) [crafted]; Ankh of Life (1713, -3.55 DPS) [world_drop] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironbark Staff (20069, +0.00 DPS) [rep]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Staff of Jordan (873, -1.45 DPS) [world_drop] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (10.50 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Tome of Restoration (234590, +0.00 DPS) [pvp]; Penelope's Rose (15805, -0.69 DPS) [quest] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Oblivion's Touch (18761, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -1.90 DPS) [dungeon] |

**New at 60:** head: The Postmaster's Band; neck: Lady Maye's Pendant; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Magician's Cord; legs: Sentinel's Silk Leggings; feet: The Postmaster's Treads; finger1: The Postmaster's Seal; finger2: Channeler's Ring; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Brilliant Wand

No-known-source sample (15 of 1456, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (gnome, 131205203100311531-52500001200100000-003000000000000000)

Set DPS (verified): 800.5. Weights run: 3.9s. Verify run: 19.6s. 1456 eligible items had no known source.

Race sweep: gnome 800.5 ± 0.8, human 795.4 ± 0.8, high-order-skyborne 787.5 ± 0.8 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.886 ± 0.046, crit=0.881 ± 0.039 per rating point (14 rating = 1%, 12.339 per %), hit=1.856 ± 0.115 per rating point (10 rating = 1%, 18.556 per %), spell_haste=3.858 ± 0.806, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.561 ± 0.002, spirit=not significant (0.324 ± 0.365), mp5=not significant (0.562 ± 0.408)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 55.8 spell_power points (29.48 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -2.40 DPS) [pvp]; Cap of the Scarlet Savant (12752, -11.90 DPS, sim-verified) [quest] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 34.6 spell_power points (18.32 DPS) | yes | Star of Mystaria (12103, -1.09 DPS) [dungeon]; Lady Maye's Pendant (14558, -3.60 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -5.00 DPS) [world_drop] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 40.9 spell_power points (21.60 DPS) | yes | Field Marshal's Silk Spaulders (231602, -0.50 DPS) [pvp]; Lieutenant Commander's Silk Mantle (227102, -1.31 DPS) [pvp]; Soulstealer Mantle (13374, -1.48 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of Arcane Mastery (22330, -8.55 DPS, sim-verified) [dungeon]; Crystalline Threaded Cape (20697, -9.57 DPS) [world]; Shroud of the Exile (15421, -9.79 DPS) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Vestments (231603, -0.33 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -6.85 DPS) [pvp]; Alanna's Embrace (13314, -22.17 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 31.4 spell_power points (16.58 DPS) | yes | Manacle Cuffs (11962, -5.85 DPS) [quest]; Marshal's Silk Bracers (16438, -5.96 DPS) [pvp]; Runecloth Cuffs (254123, -6.08 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 42.5 spell_power points (22.49 DPS) | yes | Marshal's Silk Gloves (16440, -1.74 DPS) [vendor]; Marshal's Silk Gauntlets (231608, -1.74 DPS) [vendor]; Hands of Power (13253, -5.52 DPS, sim-verified) [dungeon] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 40.3 spell_power points (21.29 DPS) | yes | Clutch of Andros (13956, -2.58 DPS) [dungeon]; Highlander's Cloth Girdle (20047, -4.55 DPS) [rep]; Ban'thok Sash (11662, -7.36 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 65.5 spell_power points (34.65 DPS) | yes | Marshal's Silk Leggings (231605, -1.18 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -8.20 DPS) [pvp]; Magister's Leggings (16687, -11.50 DPS) [dungeon] |
| feet | Kayser's Boots of Precision (22231) | Blackrock Spire: War Master Voone [dungeon] | 38.5 spell_power points (20.36 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -3.35 DPS) [dungeon]; Sergeant Major's Dreadweave Boots (220891, -5.60 DPS, sim-verified) [vendor] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Julio's Band (19325, -1.41 DPS) [rep]; Maiden's Circle (13001, -3.92 DPS) [world_drop]; Naglering (11669, -19.47 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Maiden's Circle (13001, -2.64 DPS) [world_drop]; Songstone of Ironforge (12543, -3.67 DPS) [quest]; Don Julio's Band (19325, -4.86 DPS, sim-verified) [rep] |
| trinket1 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Eye of the Beast (13968) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, -5.37 DPS, sim-verified) [vendor] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Staff of Jordan (873, -13.94 DPS, sim-verified) [world_drop]; Headmaster's Charge (13937, -15.59 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Oblivion's Touch (18761, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -2.14 DPS) [dungeon] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Magician's Cord; legs: Sentinel's Silk Leggings; feet: Kayser's Boots of Precision; finger1: Channeler's Ring; finger2: Blessed Band of Light; trinket1: Second Wind; trinket2: Eye of the Beast; main_hand: Ironbark Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1456, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (undead, 140005000100000000-00000000000000000-000000000000000000)

Set DPS (verified): 41.9. Weights run: 2.8s. Verify run: 2.9s. 192 eligible items had no known source.

Race sweep: undead 41.9 ± 0.0, orc 40.3 ± 0.0, troll 39.9 ± 0.0 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.927 ± 0.011, crit=0.086 ± 0.002 per rating point (14 rating = 1%, 1.208 per %), hit=0.306 ± 0.012 per rating point (10 rating = 1%, 3.056 per %), spell_haste=not significant (0.277 ± 0.104), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.870 ± 0.003, spirit=not significant (0.145 ± 0.064), mp5=0.531 ± 0.128

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pearly Circlet (253957) | Tailoring [crafted] | 9.6 spell_power points (1.55 DPS) | yes | Pristine Circlet (253949, -0.61 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.66 DPS) [crafted]; Flying Tiger Goggles (4368, -1.45 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.4 spell_power points (0.07 DPS) | yes | Scholarly Pendant (277203, -0.02 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.3 spell_power points (2.16 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.84 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.51 DPS) [crafted]; Slime-encrusted Pads (6461, -1.90 DPS) [dungeon] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.03 DPS) [crafted]; Sanguine Cape (14376, -0.17 DPS) [world_drop]; Battle Healer's Cloak (20427, -0.24 DPS) [rep] |
| chest | Filigreed Pearly Gown (253909) | Tailoring [crafted] | 12.9 spell_power points (2.09 DPS) | yes | Filigreed Pristine Gown (253901, -0.48 DPS, sim-verified) [crafted]; Gray Woolen Robe (2585, -0.69 DPS) [crafted]; Mystic's Wrap (14369, -1.04 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 6.2 spell_power points (1.01 DPS) | yes | Tabitha's Cuffs (251486, -0.52 DPS) [quest]; Owlbeard Bracers (16981, -0.55 DPS) [quest]; Bright Bracers (3647, -0.55 DPS, sim-verified) [world_drop] |
| hands | Pearly Gloves (253921) | Tailoring [crafted] | 8.9 spell_power points (1.43 DPS) | yes | Serpent Gloves (5970, -0.32 DPS, sim-verified) [dungeon]; Pristine Gloves (253913, -0.34 DPS) [crafted]; Gnoll Casting Gloves (892, -0.46 DPS) [world] |
| waist | Pearly Sash (253933) | Tailoring [crafted] | 9.8 spell_power points (1.58 DPS) | yes | Pristine Sash (253925, -0.36 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.39 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.49 DPS) [crafted] |
| legs | Filigreed Pearly Leggings (253945) | Tailoring [crafted] | 15.7 spell_power points (2.54 DPS) | yes | Filigreed Pristine Leggings (253937, -0.61 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -1.35 DPS) [world_drop]; Silk-threaded Trousers (1929, -1.41 DPS) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.7 spell_power points (1.73 DPS) | yes | Pearly Boots (253897, -0.48 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.80 DPS) [crafted]; Red Woolen Boots (4313, -0.97 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 6.1 spell_power points (0.98 DPS) | yes | Loop of Sacrifice (281673, -0.23 DPS) [quest]; Volcanic Rock Ring (12053, -0.53 DPS) [world_drop]; Furlfeather Ring (286745, -0.82 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.6 spell_power points (0.90 DPS) | yes | Loop of Sacrifice (281673, -0.15 DPS) [quest]; Volcanic Rock Ring (12053, -0.45 DPS) [world_drop]; Furlfeather Ring (286745, -0.74 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Crescent Staff (6505) | Leaders of the Fang [quest] | 32.5 spell_power points (5.25 DPS) | yes | Evocator's Blade (2567, -0.31 DPS) [dungeon]; Staff of the Blessed Seer (2271, -0.49 DPS) [dungeon]; Gnarled Necromancer's Staff (251534, -0.81 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Skycaller (12984) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cookie's Stirring Rod (5198, -1.02 DPS, sim-verified) [dungeon]; Sizzle Stick (8071, -1.59 DPS) [quest]; Greater Magic Wand (11288, -1.63 DPS) [crafted] |

**New at 20:** head: Pearly Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Pearly Gown; wrist: Mindthrust Bracers; hands: Pearly Gloves; waist: Pearly Sash; legs: Filigreed Pearly Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Crescent Staff; ranged: Skycaller

No-known-source sample (15 of 192, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 8350 The 1 Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 23173 Abomination Skin Leggings

### Band 30 (undead, 141005203100301000-00000000000000000-000000000000000000)

Set DPS (verified): 90.7. Weights run: 2.9s. Verify run: 4.2s. 318 eligible items had no known source.

Race sweep: undead 90.7 ± 0.1, orc 89.3 ± 0.1, troll 88.1 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.499 ± 0.009, crit=0.072 ± 0.002 per rating point (14 rating = 1%, 1.007 per %), hit=0.236 ± 0.012 per rating point (10 rating = 1%, 2.363 per %), spell_haste=0.952 ± 0.142, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.563 ± 0.003, spirit=not significant (0.120 ± 0.087), mp5=not significant (0.455 ± 0.114)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.7 spell_power points (2.62 DPS) | yes | Enchanter's Cowl (4322, -0.16 DPS) [crafted]; Silk Headband (7050, -0.61 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.83 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.0 spell_power points (0.67 DPS) | yes | Darkspear Warding Pendant (272075, -0.11 DPS) [vendor]; Kaleidoscope Chain (13084, -0.12 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.14 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.5 spell_power points (3.01 DPS) | yes | Death Speaker Mantle (6685, -0.45 DPS) [dungeon]; Invoker's Mantle (215365, -0.89 DPS) [crafted]; Magician's Mantle (12998, -0.89 DPS) [world_drop] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, +0.00 DPS) [crafted]; Windsong Drape (15468, +0.00 DPS) [quest] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 15.9 spell_power points (3.56 DPS) | yes | Green Silk Armor (7065, -0.10 DPS) [crafted]; Mechbuilder's Overalls (9508, -0.44 DPS) [dungeon]; Tree Bark Jacket (1486, -0.66 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 10.4 spell_power points (2.32 DPS) | yes | Glowing Magical Bracelets (13106, -1.42 DPS) [world_drop]; Mindthrust Bracers (1974, -1.45 DPS) [dungeon]; Nightsky Wristbands (6407, -1.57 DPS) [world_drop] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 8.6 spell_power points (1.91 DPS) | yes | Jutebraid Gloves (10654, -0.01 DPS) [quest]; Serpent Gloves (5970, -0.35 DPS) [dungeon]; Truefaith Gloves (7049, -0.46 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.5 spell_power points (2.79 DPS) | yes | Belt of Arugal (6392, -0.31 DPS) [dungeon]; Warsong Sash (16975, -0.33 DPS) [quest]; Invoker's Cord (215366, -0.67 DPS) [crafted] |
| legs | Pearly Leggings (253995) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gaze Dreamer Pants (6903, +0.00 DPS) [dungeon]; Pristine Leggings (253987, -0.20 DPS) [crafted]; Filigreed Pearly Leggings (253945, -0.52 DPS) [crafted] |
| feet | Golden Slippers (254009) | Tailoring [crafted] | 11.3 spell_power points (2.52 DPS) | yes | Gilded Slippers (254001, -0.07 DPS) [crafted]; Acidic Walkers (9454, -0.41 DPS) [dungeon]; Spidersilk Boots (4320, -0.51 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 10.3 spell_power points (2.31 DPS) | yes | Advisor's Ring (19521, -0.54 DPS) [rep]; Ring of Power Regulation (273030, -0.66 DPS) [dungeon] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Advisor's Ring (19521, +0.00 DPS) [rep]; Ring of Power Regulation (273030, +0.00 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glimmering Staff (249392, -0.22 DPS) [crafted]; Hydrocane (9452, -1.00 DPS) [dungeon]; Tendonscraper Dagger (273053, -1.00 DPS) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.0 spell_power points (2.23 DPS) | yes | Orb of Souls (249395, -1.18 DPS) [crafted]; Witch's Finger (16887, -1.34 DPS) [quest]; Satyr's Rod (15962, -1.51 DPS) [world_drop] |
| ranged | Starfaller (13063) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderwood (13062, -0.33 DPS) [world_drop]; Greater Mystic Wand (217287, -0.40 DPS) [crafted]; Wand of Eventide (5214, -0.67 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Darkspear Raider's Cloak; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Defiler's Cloth Girdle; legs: Pearly Leggings; feet: Golden Slippers; finger1: Snake Hoop; finger2: Runemender's Seal; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Starfaller

No-known-source sample (15 of 318, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 8350 The 1 Ring; 9362 Brilliant Gold Ring

### Band 40 (undead, 141005213100301531-00000000000000000-000000000000000000)

Set DPS (verified): 166.8. Weights run: 3.0s. Verify run: 3.0s. 448 eligible items had no known source.

Race sweep: undead 166.8 ± 0.3, orc 165.4 ± 0.3, troll 163.4 ± 0.3 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.587 ± 0.022, crit=0.261 ± 0.006 per rating point (14 rating = 1%, 3.650 per %), hit=0.604 ± 0.034 per rating point (10 rating = 1%, 6.039 per %), spell_haste=2.133 ± 0.376, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.628 ± 0.002, spirit=not significant (0.277 ± 0.203), mp5=not significant (0.370 ± 0.231)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Miner's Hat of the Deep (9429) | Uldaman [dungeon] | 22.0 spell_power points (7.02 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.31 DPS) [crafted]; Papal Fez (9431, -0.48 DPS) [dungeon]; Corpseshroud (10574, -1.22 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.7 spell_power points (3.74 DPS) | yes | Triune Amulet (7722, -0.56 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.86 DPS) [quest]; Necklace of Calisea (1714, -1.80 DPS) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.6 spell_power points (4.68 DPS) | yes | Green Silken Shoulders (7057, -0.06 DPS) [crafted]; Bloodmage Mantle (7684, -0.11 DPS) [dungeon]; Berylline Pads (4197, -0.56 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.09 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.44 DPS) [vendor] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 27.2 spell_power points (8.68 DPS) | yes | Robe of Power (7054, -1.25 DPS) [crafted]; Elemental Raiment (9434, -1.97 DPS) [world_drop]; Robe of the Magi (1716, -2.83 DPS, sim-verified) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 10.1 spell_power points (3.23 DPS) | yes | Condor Bracers (15864, -0.99 DPS) [quest]; Earthen Silk Cuffs (254019, -1.51 DPS) [crafted]; Radiant Silver Bracers (4545, -1.73 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.3 spell_power points (7.12 DPS) | yes | Black Mageweave Gloves (10003, -2.33 DPS) [crafted]; Red Mageweave Gloves (10018, -2.41 DPS, sim-verified) [crafted]; Golden Handwraps (254029, -3.00 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 16.3 spell_power points (5.22 DPS) | yes | Deathmage Sash (10771, -0.17 DPS) [dungeon]; Golden Cord (254045, -0.38 DPS) [crafted]; Gilded Cord (254037, -0.63 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.0 spell_power points (6.72 DPS) | yes | Stoneweaver Leggings (9407, -1.92 DPS) [dungeon]; Gaze Dreamer Pants (6903, -2.09 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.45 DPS, sim-verified) [crafted] |
| feet | Golden Slippers (254009) | Tailoring [crafted] | 13.4 spell_power points (4.28 DPS) | yes | Gilded Slippers (254001, -0.37 DPS) [crafted]; Acidic Walkers (9454, -0.82 DPS) [dungeon]; Spidersilk Boots (4320, -1.29 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.9 spell_power points (4.13 DPS) | yes | Advisor's Ring (19520, -0.90 DPS) [rep]; Runemender's Seal (273031, -1.43 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 12.1 spell_power points (3.85 DPS) | yes | Runemender's Seal (273031, -1.15 DPS) [world]; Advisor's Ring (19521, -1.38 DPS) [rep] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+3.3 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Celestial Stave (9517, -10.50 DPS) [quest]; Illusionary Rod (7713, -10.60 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -10.88 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Nether Force Wand (11263) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plaguerot Sprig (10766, -0.32 DPS) [dungeon]; Twisted Nether Wand (249144, -0.66 DPS) [crafted]; Starfaller (13063, -0.70 DPS) [world_drop] |

**New at 40:** head: Miner's Hat of the Deep; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; finger1: Philanthropist's Ring; finger2: Snake Hoop; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Staff of Jordan; ranged: Nether Force Wand

No-known-source sample (15 of 448, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (undead, 141005213100311531-30000000000000000-033000000000000000)

Set DPS (verified): 275.4. Weights run: 3.1s. Verify run: 5.3s. 590 eligible items had no known source.

Race sweep: undead 275.4 ± 0.4, orc 272.0 ± 0.4, troll 269.1 ± 0.4 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.838 ± 0.031, crit=0.407 ± 0.011 per rating point (14 rating = 1%, 5.695 per %), hit=0.811 ± 0.055 per rating point (10 rating = 1%, 8.112 per %), spell_haste=3.087 ± 0.533, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.805 ± 0.002, spirit=not significant (0.136 ± 0.278), mp5=not significant (0.275 ± 0.311)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 35.8 spell_power points (12.72 DPS) | yes | Dreamweave Circlet (10041, -1.69 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -2.14 DPS) [vendor]; Chief Architect's Monocle (11839, -2.53 DPS) [dungeon] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 21.0 spell_power points (7.48 DPS) | yes | Jarkal's Enhancing Necklace (7888, -2.91 DPS) [quest]; Gemshard Heart (17707, -3.02 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -3.26 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 28.1 spell_power points (9.99 DPS) | yes | Kentic Amice (11624, -0.84 DPS) [dungeon]; Blood Guard's Dreadweave Mantle (220905, -2.44 DPS) [vendor]; Nethershine Shoulders (254057, -2.45 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -0.11 DPS) [quest]; Runecloth Cloak (13860, -1.47 DPS) [crafted]; Big Voodoo Cloak (8216, -2.60 DPS) [crafted] |
| chest | Robes of Insight (940) | World drop [world_drop] | 31.1 spell_power points (11.06 DPS) | yes | Robe of the Magi (1716, -1.21 DPS) [world_drop]; Dreamweave Vest (10021, -1.30 DPS) [crafted]; Stone Guard's Dreadweave Vest (220904, -1.49 DPS) [vendor] |
| wrist | Nethershine Cuffs (254069) | Tailoring [crafted] | 17.0 spell_power points (6.05 DPS) | yes | Aristocratic Cuffs (12546, -0.16 DPS) [dungeon]; Nethergeld Cuffs (254061, -1.23 DPS) [crafted]; Bloodband Bracers (11469, -1.44 DPS) [quest] |
| hands | Golden Gloves (254103) | Tailoring [crafted] | 24.8 spell_power points (8.81 DPS) | yes | Runecloth Gloves (13863, -1.42 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.50 DPS) [vendor]; Dreamweave Gloves (10019, -1.96 DPS, sim-verified) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 28.5 spell_power points (10.13 DPS) | yes | Satyrmane Sash (17755, -2.17 DPS) [dungeon]; Dawnspire Cord (12466, -2.33 DPS) [dungeon]; Golden Waistcord (254089, -3.45 DPS, sim-verified) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | 31.4 spell_power points (11.16 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -1.29 DPS) [vendor]; Red Mageweave Pants (10009, -2.60 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -3.85 DPS) [dungeon] |
| feet | First Sergeant's Dreadweave Boots (220909) | PvP rank 5 · First Sergeant · Horde [vendor] | 23.7 spell_power points (8.41 DPS) | yes | Golden Sandals (254115, -0.00 DPS) [crafted]; Gilded Sandals (254107, -1.82 DPS) [crafted]; Southsea Mojo Boots (20641, -2.29 DPS) [quest] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 20.1 spell_power points (7.13 DPS) | yes | Woodseed Hoop (17768, -2.02 DPS) [quest]; Philanthropist's Ring (281635, -2.09 DPS) [quest]; Mindseye Circle (10634, -2.11 DPS) [dungeon] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.4 spell_power points (5.48 DPS) | yes | Philanthropist's Ring (281635, -0.43 DPS) [quest]; Mindseye Circle (10634, -0.46 DPS) [dungeon]; Woodseed Hoop (17768, -2.25 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.46 DPS) [world_drop]; Rune of the Guard Captain (19120, -0.58 DPS) [quest]; Darkspear Voodoo Seal (272061, -2.11 DPS) [vendor] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Darkspear Voodoo Seal (272061, -2.17 DPS, sim-verified) [vendor] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Staff of Jordan (873, -6.96 DPS, sim-verified) [world_drop] |
| off_hand | Truesilver Conduit (249455) | Enchanting [crafted] | 16.7 spell_power points (5.94 DPS) | yes | Flask of Forest Mojo (19115, -0.36 DPS) [quest]; Celestial Orb (7515, -0.42 DPS) [quest]; Orb of the Forgotten Seer (7685, -0.96 DPS) [dungeon] |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Noxious Shooter (17745, -0.49 DPS) [dungeon]; Nether Force Wand (11263, -0.51 DPS) [quest]; Charged Lightning Rod (11860, -1.01 DPS) [quest] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Robes of Insight; wrist: Nethershine Cuffs; hands: Golden Gloves; waist: Ban'thok Sash; legs: Spellshock Leggings; feet: First Sergeant's Dreadweave Boots; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Ankh of Life; main_hand: Blade of Eternal Darkness; off_hand: Truesilver Conduit; ranged: Lesser Eternal Wand

No-known-source sample (15 of 590, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (undead, 141005213100311531-30000000000000000-053500030000000000)

Set DPS (verified): 407.7. Weights run: 3.3s. Verify run: 19.4s. 1450 eligible items had no known source.

Race sweep: undead 407.7 ± 0.5, orc 404.0 ± 0.5, troll 401.8 ± 0.5 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.206 ± 0.050, crit=0.611 ± 0.017 per rating point (14 rating = 1%, 8.549 per %), hit=1.260 ± 0.084 per rating point (10 rating = 1%, 12.596 per %), spell_haste=not significant (3.613 ± 0.939), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.751 ± 0.003, spirit=not significant (0.683 ± 0.479), mp5=not significant (1.037 ± 0.486)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 54.5 spell_power points (19.72 DPS) | yes | Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; The Postmaster's Band (13390, -0.91 DPS) [dungeon] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 34.5 spell_power points (12.50 DPS) | yes | Lady Maye's Pendant (14558, -0.23 DPS) [world_drop]; Nacreous Shell Necklace (22403, -1.19 DPS) [dungeon]; Star of Mystaria (12103, -1.79 DPS) [dungeon] |
| shoulder | Darkspear Shoulderpads (272103) | Creeg Bothunk [vendor] | 44.5 spell_power points (16.13 DPS) | yes | Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -1.21 DPS) [crafted]; Shroud of the Nathrezim (18720, -1.54 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, -3.41 DPS) [quest]; Shroud of Arcane Mastery (22330, -4.48 DPS) [dungeon]; Spritecaster Cape (11623, -4.68 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Raiment (231596, +0.00 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -3.46 DPS) [pvp]; Magister's Robes (16688, -7.74 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 36.4 spell_power points (13.19 DPS) | yes | General's Silk Cuffs (16538, -2.42 DPS) [pvp]; Manacle Cuffs (11962, -4.11 DPS) [quest]; Runecloth Cuffs (254123, -4.84 DPS) [crafted] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Raider Handwraps (272097, -1.57 DPS, sim-verified) [vendor] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 47.9 spell_power points (17.36 DPS) | yes | General's Silk Sash (16537, -3.61 DPS) [pvp]; Clutch of Andros (13956, -4.50 DPS) [dungeon]; Ban'thok Sash (11662, -5.39 DPS, sim-verified) [dungeon] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 57.7 spell_power points (20.91 DPS) | yes | General's Silk Trousers (231595, +0.00 DPS) [pvp]; Legionnaire's Silk Legguards (227107, -1.55 DPS) [pvp]; Magister's Leggings (16687, -2.84 DPS, sim-verified) [dungeon] |
| feet | Dragonrider Boots (18102) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Omnicast Boots (11822, -1.02 DPS) [dungeon]; Kayser's Boots of Precision (22231, -1.83 DPS, sim-verified) [dungeon] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Seal of Rivendare (13345, -2.33 DPS) [dungeon]; Blessed Band of Light (272407, -2.49 DPS) [vendor]; Naglering (11669, -6.67 DPS, sim-verified) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -0.90 DPS) [vendor]; Seal of Rivendare (13345, -1.25 DPS, sim-verified) [dungeon]; Brainlash (6440, -1.47 DPS) [dungeon] |
| trinket1 | Eye of the Beast (13968) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Staff of Jordan (873, -1.45 DPS) [world_drop]; Ironbark Staff (20220, -6.21 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (10.50 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp]; Penelope's Rose (15805, -0.69 DPS) [quest] |
| ranged | Bonecreeper Stylus (13938) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Oblivion's Touch (18761, +0.00 DPS) [dungeon]; Brilliant Wand (249385, -0.13 DPS) [crafted]; Ritssyn's Wand of Bad Mojo (22408, -2.07 DPS) [dungeon] |

**New at 60:** head: Spellweaver's Turban; neck: Jeweled Amulet of Cainwyn; shoulder: Darkspear Shoulderpads; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Magician's Cord; legs: Outrider's Silk Leggings; feet: Dragonrider Boots; finger1: Channeler's Ring; finger2: Maiden's Circle; trinket1: Eye of the Beast; trinket2: Burst of Knowledge; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Bonecreeper Stylus

No-known-source sample (15 of 1450, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60, raid preset (undead, 131205203100311531-52500001200100000-003000000000000000)

Set DPS (verified): 789.0. Weights run: 3.9s. Verify run: 21.3s. 1450 eligible items had no known source.

Race sweep: undead 789.0 ± 0.8, orc 787.2 ± 0.8, troll 779.9 ± 0.8 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.886 ± 0.046, crit=0.881 ± 0.039 per rating point (14 rating = 1%, 12.339 per %), hit=1.856 ± 0.115 per rating point (10 rating = 1%, 18.556 per %), spell_haste=3.858 ± 0.806, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.561 ± 0.002, spirit=not significant (0.324 ± 0.365), mp5=not significant (0.562 ± 0.408)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 55.8 spell_power points (29.48 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -2.40 DPS) [pvp]; Cap of the Scarlet Savant (12752, -4.87 DPS, sim-verified) [quest] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 34.6 spell_power points (18.32 DPS) | yes | Star of Mystaria (12103, -1.58 DPS, sim-verified) [dungeon]; Lady Maye's Pendant (14558, -3.60 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -5.00 DPS) [world_drop] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 40.9 spell_power points (21.60 DPS) | yes | Warlord's Silk Amice (231594, -0.50 DPS) [pvp]; Champion's Silk Mantle (227104, -1.31 DPS) [pvp]; Soulstealer Mantle (13374, -6.17 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of Arcane Mastery (22330, -6.71 DPS, sim-verified) [dungeon]; Crystalline Threaded Cape (20697, -9.57 DPS) [world]; Shroud of the Exile (15421, -9.79 DPS) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Raiment (231596, -0.33 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -6.85 DPS) [pvp]; Alanna's Embrace (13314, -15.72 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 31.4 spell_power points (16.58 DPS) | yes | Manacle Cuffs (11962, -5.85 DPS) [quest]; General's Silk Cuffs (16538, -5.96 DPS) [pvp]; Runecloth Cuffs (254123, -6.08 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 42.5 spell_power points (22.49 DPS) | yes | General's Silk Handguards (16540, -1.74 DPS) [vendor]; General's Silk Gauntlets (231599, -1.74 DPS) [vendor]; Hands of Power (13253, -3.87 DPS, sim-verified) [dungeon] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 40.3 spell_power points (21.29 DPS) | yes | Ban'thok Sash (11662, -0.45 DPS) [dungeon]; Clutch of Andros (13956, -2.58 DPS) [dungeon]; Defiler's Cloth Girdle (20163, -4.55 DPS) [rep] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 48.1 spell_power points (25.42 DPS) | yes | Legionnaire's Silk Legguards (227107, +0.00 DPS) [pvp]; General's Silk Trousers (231595, +0.00 DPS) [pvp]; Magister's Leggings (16687, -4.67 DPS, sim-verified) [dungeon] |
| feet | Kayser's Boots of Precision (22231) | Blackrock Spire: War Master Voone [dungeon] | 38.5 spell_power points (20.36 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; First Sergeant's Dreadweave Boots (220909, -2.10 DPS) [vendor]; Dragonrider Boots (18102, -3.35 DPS) [dungeon] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Julio's Band (19325, -1.41 DPS) [rep]; Maiden's Circle (13001, -3.92 DPS) [world_drop]; Naglering (11669, -18.10 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Julio's Band (19325, -0.13 DPS) [rep]; Maiden's Circle (13001, -2.64 DPS) [world_drop]; Eye of Orgrimmar (12545, -3.67 DPS) [quest] |
| trinket1 | Eye of the Beast (13968) | For The Horde! [quest] | sim-verified (+5.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+3.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -0.69 DPS) [crafted]; Second Wind (11819, -3.57 DPS, sim-verified) [dungeon] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Staff of Jordan (873, -14.37 DPS, sim-verified) [world_drop]; Headmaster's Charge (13937, -15.59 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Oblivion's Touch (18761, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -2.16 DPS) [dungeon] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Magician's Cord; legs: Outrider's Silk Leggings; feet: Kayser's Boots of Precision; finger1: Channeler's Ring; finger2: Blessed Band of Light; trinket1: Eye of the Beast; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1450, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

