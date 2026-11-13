# Leveling BiS: Arcane

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 140005000100000000-00000000000000000-000000000000000000)

Set DPS (verified): 42.7. Weights run: 3.6s. Verify run: 3.1s. 198 eligible items had no known source.

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
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | 30.6 spell_power points (4.94 DPS) | yes | Emberstone Staff (5201, -0.95 DPS) [dungeon]; Staff of the Blessed Seer (2271, -1.04 DPS, sim-verified) [dungeon]; Buzzer Blade (2169, -1.39 DPS) [dungeon] |
| off_hand | Mathystra Channeling Orb (281291) | Baron Marinous [quest] | 4.6 spell_power points (0.75 DPS) | yes | Tear of Grief (5611, -0.07 DPS) [quest]; Dwarven Tome (279898, -0.24 DPS) [quest]; Seer's Fine Stein (7608, -0.28 DPS) [world_drop] |
| ranged | Skycaller (12984) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cookie's Stirring Rod (5198, -0.98 DPS, sim-verified) [dungeon]; Sizzle Stick (8071, -1.47 DPS) [quest]; Greater Magic Wand (11288, -1.51 DPS) [crafted] |

**New at 20:** head: Pearly Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Pearly Gown; wrist: Mindthrust Bracers; hands: Pearly Gloves; waist: Pearly Sash; legs: Filigreed Pearly Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Evocator's Blade; off_hand: Mathystra Channeling Orb; ranged: Skycaller

No-known-source sample (15 of 198, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 8350 The 1 Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (gnome, 141005203100301000-00000000000000000-000000000000000000)

Set DPS (verified): 90.1. Weights run: 3.8s. Verify run: 7.0s. 326 eligible items had no known source.

Race sweep: gnome 90.1 ± 0.1, human 88.6 ± 0.1, high-order-skyborne 88.0 ± 0.1 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.641 ± 0.015, crit=0.108 ± 0.003 per rating point (14 rating = 1%, 1.509 per %), hit=0.420 ± 0.023 per rating point (10 rating = 1%, 4.199 per %), spell_haste=1.336 ± 0.251, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.623 ± 0.002, spirit=not significant (0.471 ± 0.166), mp5=not significant (0.762 ± 0.201)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 13.8 spell_power points (3.73 DPS) | yes | Enchanter's Cowl (4322, -0.38 DPS) [crafted]; Nightsky Cowl (4039, -1.02 DPS) [world_drop]; Silk Headband (7050, -1.30 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 4.4 spell_power points (1.20 DPS) | yes | Crystal Starfire Medallion (5003, -0.13 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.16 DPS) [world]; Seal of the Expedition (271920, -0.34 DPS) [rep] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.8 spell_power points (3.99 DPS) | yes | Death Speaker Mantle (6685, -0.46 DPS) [dungeon]; Magician's Mantle (12998, -1.08 DPS) [world_drop]; Invoker's Mantle (215365, -1.23 DPS) [crafted] |
| back | Prelacy Cape (7004) | Researching the Corruption [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Thresher Cape (6901, -0.02 DPS) [dungeon]; Expeditionary's Cape (271923, -0.08 DPS) [rep]; Repairman's Cape (9605, -0.15 DPS) [quest] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 20.5 spell_power points (5.53 DPS) | yes | Green Silk Armor (7065, -0.84 DPS) [crafted]; Mechbuilder's Overalls (9508, -1.29 DPS) [dungeon]; Pearly Gown (253969, -1.62 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 11.3 spell_power points (3.05 DPS) | yes | Mindthrust Bracers (1974, -1.21 DPS, sim-verified) [dungeon]; Nightsky Wristbands (6407, -1.63 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.66 DPS) [world_drop] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Gloves of Old (9395, -0.34 DPS) [world_drop]; Gloves of Meditation (4318, -0.49 DPS) [crafted]; Shilly Mitts (9609, -0.92 DPS, sim-verified) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Belt of Arugal (6392, -0.43 DPS, sim-verified) [dungeon]; Crimson Silk Belt (7055, -0.66 DPS) [crafted]; Invoker's Cord (215366, -0.73 DPS) [crafted] |
| legs | Pearly Leggings (253995) | Tailoring [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Gaze Dreamer Pants (6903, -0.42 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.47 DPS) [crafted]; Filigreed Pearly Leggings (253945, -0.81 DPS) [crafted] |
| feet | Golden Slippers (254009) | Tailoring [crafted] | 14.5 spell_power points (3.91 DPS) | yes | Gilded Slippers (254001, -0.30 DPS) [crafted]; Acidic Walkers (9454, -0.66 DPS) [dungeon]; Glinteye Slippers (273024, -1.26 DPS) [dungeon] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 9.7 spell_power points (2.61 DPS) | yes | Lorekeeper's Ring (19525, -0.31 DPS) [rep]; Ring of Power Regulation (273030, -0.37 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.2 spell_power points (2.49 DPS) | yes | Lorekeeper's Ring (19525, -0.18 DPS) [rep]; Ring of Power Regulation (273030, -0.25 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Soulstaff (249393, -0.73 DPS) [crafted]; Rod of the Sleepwalker (1155, -1.27 DPS) [dungeon]; Glimmering Staff (249392, -2.51 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.8 spell_power points (2.93 DPS) | yes | Orb of Souls (249395, -1.09 DPS) [crafted]; Arcane Infused Rod (279838, -1.40 DPS, sim-verified) [quest]; Satyr's Rod (15962, -1.64 DPS) [world_drop] |
| ranged | Starfaller (13063) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderwood (13062, -0.44 DPS) [world_drop]; Greater Mystic Wand (217287, -0.59 DPS) [crafted]; Wand of Mana Concentration (273043, -0.70 DPS) [world] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Prelacy Cape; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Highlander's Cloth Girdle; legs: Pearly Leggings; feet: Golden Slippers; finger1: Runemender's Seal; finger2: Snake Hoop; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Starfaller

No-known-source sample (15 of 326, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 141005213100301531-00000000000000000-000000000000000000)

Set DPS (verified): 166.9. Weights run: 3.9s. Verify run: 6.4s. 456 eligible items had no known source.

Race sweep: gnome 166.9 ± 0.3, human 166.0 ± 0.3, high-order-skyborne 164.0 ± 0.3 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.600 ± 0.021, crit=0.264 ± 0.006 per rating point (14 rating = 1%, 3.691 per %), hit=0.611 ± 0.033 per rating point (10 rating = 1%, 6.114 per %), spell_haste=2.373 ± 0.379, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.627 ± 0.002, spirit=not significant (0.176 ± 0.220), mp5=not significant (0.236 ± 0.237)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Miner's Hat of the Deep (9429) | Uldaman [dungeon] | 22.2 spell_power points (7.04 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.38 DPS) [crafted]; Papal Fez (9431, -0.77 DPS) [dungeon]; Corpseshroud (10574, -1.20 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.8 spell_power points (3.74 DPS) | yes | Triune Amulet (7722, -0.81 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.89 DPS) [quest]; Necklace of Calisea (1714, -2.02 DPS) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.8 spell_power points (4.69 DPS) | yes | Green Silken Shoulders (7057, -0.06 DPS) [crafted]; Bloodmage Mantle (7684, -0.13 DPS) [dungeon]; Berylline Pads (4197, -0.57 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.30 DPS) [crafted]; Expeditionary's Cape (271922, -0.54 DPS) [rep] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Robe of the Magi (1716, -1.03 DPS, sim-verified) [world_drop]; Robe of Power (7054, -1.03 DPS) [crafted]; Elemental Raiment (9434, -1.54 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.7 spell_power points (3.08 DPS) | yes | Arcane Runed Bracers (4744, -0.22 DPS) [quest]; Condor Bracers (15864, -0.86 DPS) [quest]; Windchaser Cuffs (14429, -1.36 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.6 spell_power points (6.86 DPS) | yes | Red Mageweave Gloves (10018, -1.94 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -2.10 DPS) [crafted]; Golden Handwraps (254029, -2.74 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 16.4 spell_power points (5.20 DPS) | yes | Deathmage Sash (10771, -0.13 DPS) [dungeon]; Golden Cord (254045, -0.56 DPS) [crafted]; Gilded Cord (254037, -0.81 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.2 spell_power points (6.72 DPS) | yes | Crimson Silk Pantaloons (7062, -1.84 DPS, sim-verified) [crafted]; Stoneweaver Leggings (9407, -2.31 DPS) [dungeon]; Pearly Leggings (253995, -2.33 DPS) [crafted] |
| feet | Golden Slippers (254009) | Tailoring [crafted] | 13.1 spell_power points (4.14 DPS) | yes | Gilded Slippers (254001, -0.36 DPS) [crafted]; Acidic Walkers (9454, -0.81 DPS) [dungeon]; Spidersilk Boots (4320, -1.16 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.0 spell_power points (4.12 DPS) | yes | Snake Hoop (6750, -1.27 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.59 DPS) [quest]; Runemender's Seal (273031, -1.69 DPS) [world] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.7 spell_power points (3.08 DPS) | yes | Snake Hoop (6750, -0.22 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.54 DPS) [quest]; Runemender's Seal (273031, -0.65 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+2.5 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -9.29 DPS, sim-verified) [dungeon]; Celestial Stave (9517, -10.18 DPS) [quest]; Hypnotic Blade (7714, -10.65 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Nether Force Wand (11263) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plaguerot Sprig (10766, -0.30 DPS) [dungeon]; Twisted Nether Wand (249144, -0.70 DPS) [crafted]; Starfaller (13063, -0.74 DPS) [world_drop] |

**New at 40:** head: Miner's Hat of the Deep; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Staff of Jordan; ranged: Nether Force Wand

No-known-source sample (15 of 456, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (gnome, 141005213100311531-30000000000000000-033000000000000000)

Set DPS (verified): 280.3. Weights run: 4.0s. Verify run: 6.5s. 600 eligible items had no known source.

Race sweep: gnome 280.3 ± 0.4, human 279.9 ± 0.4, high-order-skyborne 276.2 ± 0.4 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.851 ± 0.031, crit=0.410 ± 0.011 per rating point (14 rating = 1%, 5.739 per %), hit=0.862 ± 0.054 per rating point (10 rating = 1%, 8.624 per %), spell_haste=3.157 ± 0.546, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.805 ± 0.002, spirit=not significant (0.722 ± 0.235), mp5=not significant (0.727 ± 0.282)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreamweave Circlet (10041) | Tailoring [crafted] | 38.2 spell_power points (13.54 DPS) | yes | Red Mageweave Headband (10033, -0.76 DPS) [crafted]; Gemburst Circlet (10751, -2.38 DPS) [quest]; Chief Architect's Monocle (11839, -2.47 DPS) [dungeon] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 21.1 spell_power points (7.48 DPS) | yes | Mindburst Medallion (11196, -2.27 DPS) [quest]; Glowing Eye of Mordresh (10769, -2.59 DPS) [dungeon]; Horizon Choker (13085, -4.17 DPS, sim-verified) [world_drop] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 29.4 spell_power points (10.42 DPS) | yes | Nethershine Shoulders (254057, -1.62 DPS) [crafted]; Rotgrip Mantle (17732, -2.05 DPS, sim-verified) [dungeon]; Knight-Lieutenant's Dreadweave Mantle (220887, -2.84 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Expeditionary's Cape (271921, -2.81 DPS) [rep]; Featherskin Cape (10843, -3.41 DPS) [world]; Runecloth Cloak (13860, -3.74 DPS, sim-verified) [crafted] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 41.2 spell_power points (14.60 DPS) | yes | Dreamweave Vest (10021, -1.92 DPS) [crafted]; Robes of Insight (940, -3.46 DPS) [world_drop]; Robe of the Magi (1716, -3.71 DPS) [world_drop] |
| wrist | Nethershine Cuffs (254069) | Tailoring [crafted] | 20.0 spell_power points (7.10 DPS) | yes | Nethergeld Cuffs (254061, -1.23 DPS) [crafted]; Bloodband Bracers (11469, -1.85 DPS) [quest]; Aristocratic Cuffs (12546, -2.12 DPS, sim-verified) [dungeon] |
| hands | Golden Gloves (254103) | Tailoring [crafted] | 29.6 spell_power points (10.49 DPS) | yes | Runecloth Gloves (13863, -1.22 DPS) [crafted]; Gilded Gloves (254095, -1.88 DPS) [crafted]; Dreamweave Gloves (10019, -2.84 DPS, sim-verified) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 29.1 spell_power points (10.33 DPS) | yes | Satyrmane Sash (17755, -2.35 DPS) [dungeon]; Golden Cord (254045, -2.38 DPS) [crafted]; Golden Waistcord (254089, -4.31 DPS, sim-verified) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | 31.5 spell_power points (11.17 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -0.07 DPS) [dungeon]; Senior Designer's Pantaloons (11841, -0.20 DPS) [dungeon]; Knight's Dreadweave Leggings (220888, -1.26 DPS) [vendor] |
| feet | Sergeant Major's Dreadweave Boots (220891) | PvP rank 5 · Sergeant Major · Alliance [vendor] | 24.3 spell_power points (8.61 DPS) | yes | Golden Sandals (254115, -0.18 DPS) [crafted]; Mistwalker Boots (10629, -0.81 DPS) [dungeon]; Vinerot Sandals (17748, -1.21 DPS) [dungeon] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 23.1 spell_power points (8.20 DPS) | yes | Lorekeeper's Ring (19523, -2.91 DPS) [rep]; Mindseye Circle (10634, -3.05 DPS) [dungeon]; Woodseed Hoop (17768, -3.05 DPS) [quest] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 17.8 spell_power points (6.33 DPS) | yes | Mindseye Circle (10634, -1.18 DPS) [dungeon]; Woodseed Hoop (17768, -1.18 DPS) [quest]; Lorekeeper's Ring (19523, -3.22 DPS, sim-verified) [rep] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Uther's Strength (11302, -2.83 DPS, sim-verified) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.19 DPS) [quest]; Uther's Strength (11302, -2.01 DPS, sim-verified) [world_drop] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Staff of Jordan (873, -8.56 DPS, sim-verified) [world_drop] |
| off_hand | Truesilver Conduit (249455) | Enchanting [crafted] | 16.8 spell_power points (5.96 DPS) | yes | Celestial Orb (7515, -0.44 DPS) [quest]; Orb of Lorica (11262, -0.82 DPS) [quest]; Enthralled Sphere (11625, -1.90 DPS, sim-verified) [dungeon] |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nether Force Wand (11263, -0.06 DPS) [quest]; Cairnstone Sliver (9654, -0.41 DPS) [quest]; Charged Lightning Rod (11860, -0.56 DPS) [quest] |

**New at 50:** head: Dreamweave Circlet; neck: Arcane Crystal Pendant; shoulder: Kentic Amice; back: Spritecaster Cape; chest: Embrace of the Wind Serpent; wrist: Nethershine Cuffs; hands: Golden Gloves; waist: Ban'thok Sash; legs: Spellshock Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Thunderbrew's Boot Flask; trinket2: Frozen Heart of the Mountain; main_hand: Blade of Eternal Darkness; off_hand: Truesilver Conduit; ranged: Lesser Eternal Wand

No-known-source sample (15 of 600, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (gnome, 141005213100311531-30000000000000000-053500030000000000)

Set DPS (verified): 412.4. Weights run: 4.1s. Verify run: 24.8s. 1459 eligible items had no known source.

Race sweep: gnome 412.4 ± 0.6, human 411.8 ± 0.6, high-order-skyborne 407.2 ± 0.6 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.206 ± 0.050, crit=0.611 ± 0.017 per rating point (14 rating = 1%, 8.549 per %), hit=1.260 ± 0.084 per rating point (10 rating = 1%, 12.596 per %), spell_haste=not significant (3.613 ± 0.939), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.751 ± 0.003, spirit=not significant (0.683 ± 0.479), mp5=not significant (1.037 ± 0.486)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | The Postmaster's Band (13390) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellweaver's Turban (22267, +0.00 DPS) [dungeon]; Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jeweled Amulet of Cainwyn (1443, +0.00 DPS) [world_drop]; Nacreous Shell Necklace (22403, -0.96 DPS) [dungeon]; Star of Mystaria (12103, -1.56 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Theramore Shoulderpads (271953, +0.00 DPS) [rep]; Shroud of the Nathrezim (18720, -0.33 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, -3.41 DPS) [quest]; Spritecaster Cape (11623, -4.68 DPS) [dungeon]; Crystalline Threaded Cape (20697, -4.86 DPS) [world] |
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

No-known-source sample (15 of 1459, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (gnome, 131205203100311531-52500001200100000-003000000000000000)

Set DPS (verified): 796.6. Weights run: 5.0s. Verify run: 28.3s. 1459 eligible items had no known source.

Race sweep: gnome 796.6 ± 0.8, human 792.9 ± 0.8, high-order-skyborne 786.3 ± 0.8 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.877 ± 0.046, crit=0.839 ± 0.039 per rating point (14 rating = 1%, 11.740 per %), hit=1.854 ± 0.113 per rating point (10 rating = 1%, 18.536 per %), spell_haste=3.861 ± 0.798, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.563 ± 0.002, spirit=not significant (0.347 ± 0.362), mp5=not significant (0.576 ± 0.405)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 56.0 spell_power points (29.51 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -2.87 DPS) [pvp]; Crimson Felt Hat (18727, -7.77 DPS, sim-verified) [dungeon] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 34.6 spell_power points (18.19 DPS) | yes | Star of Mystaria (12103, -0.99 DPS) [dungeon]; Lady Maye's Pendant (14558, -3.63 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -4.90 DPS) [world_drop] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 40.1 spell_power points (21.13 DPS) | yes | Field Marshal's Silk Spaulders (231602, -0.13 DPS) [pvp]; Soulstealer Mantle (13374, -1.22 DPS) [dungeon]; Lieutenant Commander's Silk Mantle (227102, -1.25 DPS) [pvp] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -9.51 DPS, sim-verified) [world]; Shroud of the Exile (15421, -9.66 DPS) [quest]; Spritecaster Cape (11623, -10.64 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Vestments (231603, -0.28 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -6.78 DPS) [pvp]; Alanna's Embrace (13314, -16.14 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 31.4 spell_power points (16.56 DPS) | yes | Manacle Cuffs (11962, -5.94 DPS) [quest]; Marshal's Silk Bracers (16438, -5.96 DPS) [pvp]; Runecloth Cuffs (254123, -6.15 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 41.2 spell_power points (21.72 DPS) | yes | Marshal's Silk Gloves (16440, -1.05 DPS) [vendor]; Marshal's Silk Gauntlets (231608, -1.05 DPS) [vendor]; Hands of Power (13253, -5.31 DPS, sim-verified) [dungeon] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 40.0 spell_power points (21.08 DPS) | yes | Ban'thok Sash (11662, -0.39 DPS) [dungeon]; Clutch of Andros (13956, -2.55 DPS) [dungeon]; Highlander's Cloth Girdle (20047, -4.76 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 64.2 spell_power points (33.83 DPS) | yes | Marshal's Silk Leggings (231605, -0.79 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -7.83 DPS) [pvp]; Magister's Leggings (16687, -10.60 DPS) [dungeon] |
| feet | Kayser's Boots of Precision (22231) | Blackrock Spire: War Master Voone [dungeon] | 38.3 spell_power points (20.18 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Sergeant Major's Dreadweave Boots (220891, -2.62 DPS, sim-verified) [vendor]; Dragonrider Boots (18102, -3.31 DPS) [dungeon] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Julio's Band (19325, -1.74 DPS) [rep]; Maiden's Circle (13001, -3.88 DPS) [world_drop]; Naglering (11669, -18.66 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Julio's Band (19325, -0.55 DPS) [rep]; Maiden's Circle (13001, -2.68 DPS) [world_drop]; Songstone of Ironforge (12543, -3.78 DPS) [quest] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+8.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Thunderbrew's Boot Flask (744, -8.59 DPS, sim-verified) [quest] |
| trinket2 | Eye of the Beast (13968) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Staff of Jordan (873, -13.55 DPS, sim-verified) [world_drop]; Headmaster's Charge (13937, -14.58 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Oblivion's Touch (18761, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -2.09 DPS) [dungeon] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Magician's Cord; legs: Sentinel's Silk Leggings; feet: Kayser's Boots of Precision; finger1: Channeler's Ring; finger2: Blessed Band of Light; trinket1: Burst of Knowledge; trinket2: Eye of the Beast; main_hand: Ironbark Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1459, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (undead, 140005000100000000-00000000000000000-000000000000000000)

Set DPS (verified): 41.9. Weights run: 3.6s. Verify run: 3.8s. 192 eligible items had no known source.

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

Set DPS (verified): 89.5. Weights run: 3.8s. Verify run: 8.1s. 318 eligible items had no known source.

Race sweep: undead 89.5 ± 0.1, orc 88.2 ± 0.1, troll 87.1 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.641 ± 0.015, crit=0.108 ± 0.003 per rating point (14 rating = 1%, 1.509 per %), hit=0.420 ± 0.023 per rating point (10 rating = 1%, 4.199 per %), spell_haste=1.336 ± 0.251, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.623 ± 0.002, spirit=not significant (0.471 ± 0.166), mp5=not significant (0.762 ± 0.201)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Holy Shroud (2721, -0.32 DPS, sim-verified) [world_drop]; Nightsky Cowl (4039, -0.64 DPS) [world_drop]; Silk Headband (7050, -0.92 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 4.4 spell_power points (1.20 DPS) | yes | Crystal Starfire Medallion (5003, -0.13 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.16 DPS) [world]; Darkspear Warding Pendant (272075, -0.34 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.8 spell_power points (3.99 DPS) | yes | Death Speaker Mantle (6685, -0.46 DPS) [dungeon]; Mantle of Woe (7750, -0.56 DPS) [quest]; Magician's Mantle (12998, -1.08 DPS) [world_drop] |
| back | Windsong Drape (15468) | Free at Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Thresher Cape (6901, -0.03 DPS) [dungeon]; Darkspear Raider's Cloak (272078, -0.09 DPS) [vendor]; Heavy Woolen Cloak (4311, -0.27 DPS) [crafted] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 20.5 spell_power points (5.53 DPS) | yes | Green Silk Armor (7065, -0.84 DPS) [crafted]; Mechbuilder's Overalls (9508, -1.29 DPS) [dungeon]; Pearly Gown (253969, -1.62 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 11.3 spell_power points (3.05 DPS) | yes | Mindthrust Bracers (1974, -1.07 DPS, sim-verified) [dungeon]; Nightsky Wristbands (6407, -1.63 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.66 DPS) [world_drop] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 10.1 spell_power points (2.73 DPS) | yes | Jutebraid Gloves (10654, -0.24 DPS) [quest]; Gloves of Old (9395, -0.34 DPS) [world_drop]; Gloves of Meditation (4318, -0.49 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Belt of Arugal (6392, -0.36 DPS, sim-verified) [dungeon]; Warsong Sash (16975, -0.52 DPS) [quest]; Crimson Silk Belt (7055, -0.66 DPS) [crafted] |
| legs | Pearly Leggings (253995) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Leggings (253987, -0.47 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.50 DPS, sim-verified) [dungeon]; Filigreed Pearly Leggings (253945, -0.81 DPS) [crafted] |
| feet | Golden Slippers (254009) | Tailoring [crafted] | 14.5 spell_power points (3.91 DPS) | yes | Gilded Slippers (254001, -0.30 DPS) [crafted]; Acidic Walkers (9454, -0.66 DPS) [dungeon]; Glinteye Slippers (273024, -1.26 DPS) [dungeon] |
| finger1 | Runemender's Seal (273031) | Arcane Anomaly [world] | 9.7 spell_power points (2.61 DPS) | yes | Advisor's Ring (19521, -0.31 DPS) [rep]; Ring of Power Regulation (273030, -0.37 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.2 spell_power points (2.49 DPS) | yes | Advisor's Ring (19521, -0.18 DPS) [rep]; Ring of Power Regulation (273030, -0.25 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Soulstaff (249393, -0.73 DPS) [crafted]; Rod of the Sleepwalker (1155, -1.27 DPS) [dungeon]; Glimmering Staff (249392, -2.37 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.8 spell_power points (2.93 DPS) | yes | Witch's Finger (16887, -1.54 DPS) [quest]; Satyr's Rod (15962, -1.64 DPS) [world_drop]; Orb of Souls (249395, -1.76 DPS, sim-verified) [crafted] |
| ranged | Starfaller (13063) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderwood (13062, -0.52 DPS) [world_drop]; Greater Mystic Wand (217287, -0.65 DPS) [crafted]; Wand of Mana Concentration (273043, -0.73 DPS) [world] |

**New at 30:** head: Enchanter's Cowl; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Windsong Drape; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Defiler's Cloth Girdle; legs: Pearly Leggings; feet: Golden Slippers; finger1: Runemender's Seal; finger2: Snake Hoop; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Starfaller

No-known-source sample (15 of 318, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 8350 The 1 Ring; 9362 Brilliant Gold Ring

### Band 40 (undead, 141005213100301531-00000000000000000-000000000000000000)

Set DPS (verified): 165.7. Weights run: 3.9s. Verify run: 6.2s. 448 eligible items had no known source.

Race sweep: undead 165.7 ± 0.3, orc 164.1 ± 0.3, troll 162.6 ± 0.3 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.600 ± 0.021, crit=0.264 ± 0.006 per rating point (14 rating = 1%, 3.691 per %), hit=0.611 ± 0.033 per rating point (10 rating = 1%, 6.114 per %), spell_haste=2.373 ± 0.379, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.627 ± 0.002, spirit=not significant (0.176 ± 0.220), mp5=not significant (0.236 ± 0.237)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Miner's Hat of the Deep (9429) | Uldaman [dungeon] | 22.2 spell_power points (7.04 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.38 DPS) [crafted]; Papal Fez (9431, -0.77 DPS) [dungeon]; Corpseshroud (10574, -1.20 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.8 spell_power points (3.74 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.89 DPS) [quest]; Triune Amulet (7722, -1.16 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -2.02 DPS) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.8 spell_power points (4.69 DPS) | yes | Green Silken Shoulders (7057, -0.06 DPS) [crafted]; Bloodmage Mantle (7684, -0.13 DPS) [dungeon]; Berylline Pads (4197, -0.57 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.30 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.54 DPS) [vendor] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Robe of the Magi (1716, -0.85 DPS, sim-verified) [world_drop]; Robe of Power (7054, -1.03 DPS) [crafted]; Elemental Raiment (9434, -1.54 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.7 spell_power points (3.08 DPS) | yes | Radiant Silver Bracers (4545, -0.29 DPS) [quest]; Condor Bracers (15864, -0.86 DPS) [quest]; Windchaser Cuffs (14429, -1.36 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.6 spell_power points (6.86 DPS) | yes | Black Mageweave Gloves (10003, -2.10 DPS) [crafted]; Red Mageweave Gloves (10018, -2.38 DPS, sim-verified) [crafted]; Golden Handwraps (254029, -2.74 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 16.4 spell_power points (5.20 DPS) | yes | Deathmage Sash (10771, -0.13 DPS) [dungeon]; Golden Cord (254045, -0.56 DPS) [crafted]; Gilded Cord (254037, -0.81 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.2 spell_power points (6.72 DPS) | yes | Crimson Silk Pantaloons (7062, -1.86 DPS, sim-verified) [crafted]; Stoneweaver Leggings (9407, -2.31 DPS) [dungeon]; Pearly Leggings (253995, -2.33 DPS) [crafted] |
| feet | Golden Slippers (254009) | Tailoring [crafted] | 13.1 spell_power points (4.14 DPS) | yes | Gilded Slippers (254001, -0.36 DPS) [crafted]; Acidic Walkers (9454, -0.81 DPS) [dungeon]; Spidersilk Boots (4320, -1.16 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.0 spell_power points (4.12 DPS) | yes | Snake Hoop (6750, -1.27 DPS) [quest]; Runemender's Seal (273031, -1.69 DPS) [world]; Reedknot Ring (9622, -1.90 DPS) [quest] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.7 spell_power points (3.08 DPS) | yes | Snake Hoop (6750, -0.22 DPS) [quest]; Runemender's Seal (273031, -0.65 DPS) [world]; Reedknot Ring (9622, -0.86 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+2.8 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Celestial Stave (9517, -10.18 DPS) [quest]; Hypnotic Blade (7714, -10.65 DPS) [dungeon]; Illusionary Rod (7713, -11.63 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Nether Force Wand (11263) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dancing Flame (6806, +0.00 DPS) [quest]; Plaguerot Sprig (10766, -0.29 DPS) [dungeon]; Twisted Nether Wand (249144, -0.73 DPS) [crafted] |

**New at 40:** head: Miner's Hat of the Deep; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Staff of Jordan; ranged: Nether Force Wand

No-known-source sample (15 of 448, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (undead, 141005213100311531-30000000000000000-033000000000000000)

Set DPS (verified): 277.2. Weights run: 4.0s. Verify run: 6.4s. 592 eligible items had no known source.

Race sweep: undead 277.2 ± 0.4, orc 274.2 ± 0.4, troll 272.1 ± 0.4 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.851 ± 0.031, crit=0.410 ± 0.011 per rating point (14 rating = 1%, 5.739 per %), hit=0.862 ± 0.054 per rating point (10 rating = 1%, 8.624 per %), spell_haste=3.157 ± 0.546, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.805 ± 0.002, spirit=not significant (0.722 ± 0.235), mp5=not significant (0.727 ± 0.282)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreamweave Circlet (10041) | Tailoring [crafted] | 38.2 spell_power points (13.54 DPS) | yes | Red Mageweave Headband (10033, -0.76 DPS) [crafted]; Gemburst Circlet (10751, -2.38 DPS) [quest]; Chief Architect's Monocle (11839, -2.47 DPS) [dungeon] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 21.1 spell_power points (7.48 DPS) | yes | Mindburst Medallion (11196, -2.27 DPS) [quest]; Glowing Eye of Mordresh (10769, -2.59 DPS) [dungeon]; Horizon Choker (13085, -3.08 DPS, sim-verified) [world_drop] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 29.4 spell_power points (10.42 DPS) | yes | Rotgrip Mantle (17732, -0.38 DPS) [dungeon]; Nethershine Shoulders (254057, -1.62 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -2.84 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -1.34 DPS) [quest]; Runecloth Cloak (13860, -2.71 DPS) [crafted]; Darkspear Raider's Cloak (272076, -2.81 DPS) [vendor] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 41.2 spell_power points (14.60 DPS) | yes | Dreamweave Vest (10021, -1.92 DPS) [crafted]; Robes of Insight (940, -3.46 DPS) [world_drop]; Robe of the Magi (1716, -3.71 DPS) [world_drop] |
| wrist | Nethershine Cuffs (254069) | Tailoring [crafted] | 20.0 spell_power points (7.10 DPS) | yes | Nethergeld Cuffs (254061, -1.23 DPS) [crafted]; Bloodband Bracers (11469, -1.85 DPS) [quest]; Aristocratic Cuffs (12546, -1.88 DPS, sim-verified) [dungeon] |
| hands | Golden Gloves (254103) | Tailoring [crafted] | 29.6 spell_power points (10.49 DPS) | yes | Dreamweave Gloves (10019, -1.11 DPS) [crafted]; Runecloth Gloves (13863, -1.22 DPS) [crafted]; Gilded Gloves (254095, -1.88 DPS) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 29.1 spell_power points (10.33 DPS) | yes | Satyrmane Sash (17755, -2.35 DPS) [dungeon]; Golden Cord (254045, -2.38 DPS) [crafted]; Golden Waistcord (254089, -3.65 DPS, sim-verified) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | 31.5 spell_power points (11.17 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -0.07 DPS) [dungeon]; Senior Designer's Pantaloons (11841, -0.20 DPS) [dungeon]; Stone Guard's Dreadweave Leggings (220906, -1.26 DPS) [vendor] |
| feet | First Sergeant's Dreadweave Boots (220909) | PvP rank 5 · First Sergeant · Horde [vendor] | 24.3 spell_power points (8.61 DPS) | yes | Golden Sandals (254115, -0.18 DPS) [crafted]; Mistwalker Boots (10629, -0.81 DPS) [dungeon]; Vinerot Sandals (17748, -1.21 DPS) [dungeon] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 23.1 spell_power points (8.20 DPS) | yes | Advisor's Ring (19519, -2.91 DPS) [rep]; Mindseye Circle (10634, -3.05 DPS) [dungeon]; Woodseed Hoop (17768, -3.05 DPS) [quest] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 17.8 spell_power points (6.33 DPS) | yes | Advisor's Ring (19519, -1.04 DPS) [rep]; Mindseye Circle (10634, -1.18 DPS) [dungeon]; Woodseed Hoop (17768, -1.18 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.19 DPS) [quest]; Uther's Strength (11302, -0.62 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Uther's Strength (11302, -0.01 DPS) [world_drop] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Staff of Jordan (873, -7.00 DPS, sim-verified) [world_drop] |
| off_hand | Truesilver Conduit (249455) | Enchanting [crafted] | 16.8 spell_power points (5.96 DPS) | yes | Enthralled Sphere (11625, -0.35 DPS) [dungeon]; Flask of Forest Mojo (19115, -0.35 DPS) [quest]; Celestial Orb (7515, -0.44 DPS) [quest] |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nether Force Wand (11263, -0.22 DPS) [quest]; Dancing Flame (6806, -0.61 DPS) [quest]; Noxious Shooter (17745, -0.63 DPS) [dungeon] |

**New at 50:** head: Dreamweave Circlet; neck: Arcane Crystal Pendant; shoulder: Kentic Amice; back: Spritecaster Cape; chest: Embrace of the Wind Serpent; wrist: Nethershine Cuffs; hands: Golden Gloves; waist: Ban'thok Sash; legs: Spellshock Leggings; feet: First Sergeant's Dreadweave Boots; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Blade of Eternal Darkness; off_hand: Truesilver Conduit; ranged: Lesser Eternal Wand

No-known-source sample (15 of 592, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (undead, 141005213100311531-30000000000000000-053500030000000000)

Set DPS (verified): 407.7. Weights run: 4.1s. Verify run: 24.8s. 1453 eligible items had no known source.

Race sweep: undead 407.7 ± 0.5, orc 404.0 ± 0.5, troll 401.8 ± 0.5 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.206 ± 0.050, crit=0.611 ± 0.017 per rating point (14 rating = 1%, 8.549 per %), hit=1.260 ± 0.084 per rating point (10 rating = 1%, 12.596 per %), spell_haste=not significant (3.613 ± 0.939), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.751 ± 0.003, spirit=not significant (0.683 ± 0.479), mp5=not significant (1.037 ± 0.486)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 54.5 spell_power points (19.72 DPS) | yes | Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; The Postmaster's Band (13390, -0.91 DPS) [dungeon] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 34.5 spell_power points (12.50 DPS) | yes | Lady Maye's Pendant (14558, -0.23 DPS) [world_drop]; Nacreous Shell Necklace (22403, -1.19 DPS) [dungeon]; Star of Mystaria (12103, -1.79 DPS) [dungeon] |
| shoulder | Darkspear Shoulderpads (272103) | Creeg Bothunk [vendor] | 44.5 spell_power points (16.13 DPS) | yes | Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -1.21 DPS) [crafted]; Shroud of the Nathrezim (18720, -1.54 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, -3.41 DPS) [quest]; Spritecaster Cape (11623, -4.68 DPS) [dungeon]; Crystalline Threaded Cape (20697, -4.86 DPS) [world] |
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

No-known-source sample (15 of 1453, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60, raid preset (undead, 131205203100311531-52500001200100000-003000000000000000)

Set DPS (verified): 785.2. Weights run: 5.0s. Verify run: 25.1s. 1453 eligible items had no known source.

Race sweep: undead 785.2 ± 0.8, orc 782.5 ± 0.8, troll 774.8 ± 0.8 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.877 ± 0.046, crit=0.839 ± 0.039 per rating point (14 rating = 1%, 11.740 per %), hit=1.854 ± 0.113 per rating point (10 rating = 1%, 18.536 per %), spell_haste=3.861 ± 0.798, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.563 ± 0.002, spirit=not significant (0.347 ± 0.362), mp5=not significant (0.576 ± 0.405)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 56.0 spell_power points (29.51 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -2.87 DPS) [pvp]; Crimson Felt Hat (18727, -9.01 DPS, sim-verified) [dungeon] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 34.6 spell_power points (18.19 DPS) | yes | Star of Mystaria (12103, -0.99 DPS) [dungeon]; Lady Maye's Pendant (14558, -3.63 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -4.90 DPS) [world_drop] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 40.1 spell_power points (21.13 DPS) | yes | Warlord's Silk Amice (231594, -0.13 DPS) [pvp]; Champion's Silk Mantle (227104, -1.25 DPS) [pvp]; Soulstealer Mantle (13374, -6.51 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, -9.66 DPS) [quest]; Crystalline Threaded Cape (20697, -9.91 DPS, sim-verified) [world]; Spritecaster Cape (11623, -10.64 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Raiment (231596, -0.28 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -6.78 DPS) [pvp]; Alanna's Embrace (13314, -15.10 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 31.4 spell_power points (16.56 DPS) | yes | Manacle Cuffs (11962, -5.94 DPS) [quest]; General's Silk Cuffs (16538, -5.96 DPS) [pvp]; Runecloth Cuffs (254123, -6.15 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 41.2 spell_power points (21.72 DPS) | yes | General's Silk Handguards (16540, -1.05 DPS) [vendor]; General's Silk Gauntlets (231599, -1.05 DPS) [vendor]; Hands of Power (13253, -7.74 DPS, sim-verified) [dungeon] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 40.0 spell_power points (21.08 DPS) | yes | Ban'thok Sash (11662, -0.39 DPS) [dungeon]; Clutch of Andros (13956, -2.55 DPS) [dungeon]; Defiler's Cloth Girdle (20163, -4.76 DPS) [rep] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 48.1 spell_power points (25.34 DPS) | yes | Legionnaire's Silk Legguards (227107, +0.00 DPS) [pvp]; General's Silk Trousers (231595, +0.00 DPS) [pvp]; Magister's Leggings (16687, -5.65 DPS, sim-verified) [dungeon] |
| feet | Kayser's Boots of Precision (22231) | Blackrock Spire: War Master Voone [dungeon] | 38.3 spell_power points (20.18 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; First Sergeant's Dreadweave Boots (220909, -2.05 DPS) [vendor]; Dragonrider Boots (18102, -3.31 DPS) [dungeon] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Julio's Band (19325, -1.74 DPS) [rep]; Maiden's Circle (13001, -3.88 DPS) [world_drop]; Naglering (11669, -17.09 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Julio's Band (19325, -0.55 DPS) [rep]; Maiden's Circle (13001, -2.68 DPS) [world_drop]; Eye of Orgrimmar (12545, -3.78 DPS) [quest] |
| trinket1 | Eye of the Beast (13968) | For The Horde! [quest] | sim-verified (+8.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -0.69 DPS) [crafted]; Serenity Field (272439, -5.94 DPS, sim-verified) [vendor] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Staff of Jordan (873, -12.82 DPS, sim-verified) [world_drop]; Headmaster's Charge (13937, -14.58 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Oblivion's Touch (18761, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -2.10 DPS) [dungeon] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Magician's Cord; legs: Outrider's Silk Leggings; feet: Kayser's Boots of Precision; finger1: Channeler's Ring; finger2: Blessed Band of Light; trinket1: Eye of the Beast; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1453, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

