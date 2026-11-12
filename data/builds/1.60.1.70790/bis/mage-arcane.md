# Leveling BiS: Arcane

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 140005000100000000-00000000000000000-000000000000000000)

Set DPS (verified): 41.8. Weights run: 2.6s. Verify run: 2.6s. 198 eligible items had no known source.

Race sweep: gnome 41.8 ± 0.1, high-order-skyborne 40.3 ± 0.1, human 40.2 ± 0.1 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.992 ± 0.010, crit=0.085 ± 0.002 per rating point (14 rating = 1%, 1.185 per %), hit=0.316 ± 0.012 per rating point (10 rating = 1%, 3.164 per %), spell_haste=not significant (0.185 ± 0.083), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.888 ± 0.003, spirit=not significant (0.165 ± 0.064), mp5=1.228 ± 0.124

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pearly Circlet (253957) | Tailoring [crafted] | 9.8 spell_power points (1.61 DPS) | yes | Shadow Goggles (4373, -0.63 DPS) [crafted]; Pristine Circlet (253949, -0.66 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.50 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.3 spell_power points (0.05 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.9 spell_power points (2.29 DPS) | yes | Reinforced Woolen Shoulders (4315, -1.15 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.63 DPS) [crafted]; Slime-encrusted Pads (6461, -1.69 DPS) [dungeon] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sanguine Cape (14376, -0.17 DPS) [world_drop]; Caretaker's Cape (20428, -0.27 DPS) [rep]; Heavy Woolen Cloak (4311, -0.35 DPS, sim-verified) [crafted] |
| chest | Filigreed Pearly Gown (253909) | Tailoring [crafted] | 13.4 spell_power points (2.21 DPS) | yes | Filigreed Pristine Gown (253901, -0.52 DPS, sim-verified) [crafted]; Gray Woolen Robe (2585, -0.74 DPS) [crafted]; Mystic's Wrap (14369, -1.07 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 8.6 spell_power points (1.42 DPS) | yes | Bright Bracers (3647, -0.74 DPS, sim-verified) [world_drop]; Repurposed Hair Band (281256, -1.04 DPS) [quest]; Mystic's Bracelets (14366, -1.10 DPS) [world_drop] |
| hands | Pearly Gloves (253921) | Tailoring [crafted] | 9.2 spell_power points (1.51 DPS) | yes | Pristine Gloves (253913, -0.36 DPS) [crafted]; Gnoll Casting Gloves (892, -0.53 DPS) [world]; Serpent Gloves (5970, -0.54 DPS, sim-verified) [dungeon] |
| waist | Pearly Sash (253933) | Tailoring [crafted] | 10.2 spell_power points (1.68 DPS) | yes | Keller's Girdle (2911, -0.37 DPS) [world_drop]; Pristine Sash (253925, -0.39 DPS, sim-verified) [crafted]; Novice Arcanist's Sash (253885, -0.53 DPS) [crafted] |
| legs | Filigreed Pearly Leggings (253945) | Tailoring [crafted] | 16.4 spell_power points (2.70 DPS) | yes | Filigreed Pristine Leggings (253937, -0.66 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -1.39 DPS) [world_drop]; Silk-threaded Trousers (1929, -1.54 DPS) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 11.0 spell_power points (1.81 DPS) | yes | Pearly Boots (253897, -0.55 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.82 DPS) [crafted]; Red Woolen Boots (4313, -1.01 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 7.5 spell_power points (1.23 DPS) | yes | Lavishly Jeweled Ring (1156, -0.25 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.74 DPS) [world_drop]; Furlfeather Ring (286745, -1.06 DPS) [world] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 7.0 spell_power points (1.15 DPS) | yes | Lavishly Jeweled Ring (1156, -0.33 DPS, sim-verified) [dungeon]; Volcanic Rock Ring (12053, -0.66 DPS) [world_drop]; Furlfeather Ring (286745, -0.99 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | 30.8 spell_power points (5.07 DPS) | yes | Emberstone Staff (5201, -1.12 DPS) [dungeon]; Staff of the Blessed Seer (2271, -1.27 DPS, sim-verified) [dungeon]; Buzzer Blade (2169, -1.45 DPS) [dungeon] |
| off_hand | Mathystra Channeling Orb (281291) | Baron Marinous [quest] | 5.0 spell_power points (0.82 DPS) | yes | Tear of Grief (5611, -0.08 DPS) [quest]; Dwarven Tome (279898, -0.30 DPS) [quest]; Seer's Fine Stein (7608, -0.30 DPS) [world_drop] |
| ranged | Skycaller (12984) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cookie's Stirring Rod (5198, -0.86 DPS, sim-verified) [dungeon]; Sizzle Stick (8071, -1.37 DPS) [quest]; Greater Magic Wand (11288, -1.40 DPS) [crafted] |

**New at 20:** head: Pearly Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Pearly Gown; wrist: Mindthrust Bracers; hands: Pearly Gloves; waist: Pearly Sash; legs: Filigreed Pearly Leggings; feet: Spidersilk Boots; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Evocator's Blade; off_hand: Mathystra Channeling Orb; ranged: Skycaller

No-known-source sample (15 of 198, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 8350 The 1 Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (gnome, 141005203100301000-00000000000000000-000000000000000000)

Set DPS (verified): 90.3. Weights run: 3.1s. Verify run: 4.7s. 326 eligible items had no known source.

Race sweep: gnome 90.3 ± 0.1, human 88.9 ± 0.1, high-order-skyborne 88.0 ± 0.1 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.564 ± 0.010, crit=0.070 ± 0.002 per rating point (14 rating = 1%, 0.975 per %), hit=0.219 ± 0.013 per rating point (10 rating = 1%, 2.188 per %), spell_haste=0.744 ± 0.149, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.565 ± 0.003, spirit=0.375 ± 0.085, mp5=0.503 ± 0.115

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 13.2 spell_power points (2.95 DPS) | yes | Enchanter's Cowl (4322, -0.36 DPS) [crafted]; Silk Headband (7050, -0.95 DPS) [crafted]; Nightsky Cowl (4039, -1.03 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 3.8 spell_power points (0.84 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.08 DPS) [world]; Crystal Starfire Medallion (5003, -0.08 DPS) [world_drop]; Seal of the Expedition (271920, -0.21 DPS) [rep] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.1 spell_power points (3.14 DPS) | yes | Death Speaker Mantle (6685, -0.42 DPS) [dungeon]; Magician's Mantle (12998, -0.89 DPS) [world_drop]; Invoker's Mantle (215365, -0.95 DPS) [crafted] |
| back | Prelacy Cape (7004) | Researching the Corruption [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Thresher Cape (6901, -0.06 DPS) [dungeon]; Expeditionary's Cape (271923, -0.14 DPS) [rep]; Repairman's Cape (9605, -0.15 DPS) [quest] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 18.8 spell_power points (4.18 DPS) | yes | Green Silk Armor (7065, -0.54 DPS) [crafted]; Mechbuilder's Overalls (9508, -0.89 DPS) [dungeon]; Death Speaker Robes (6682, -1.24 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 10.5 spell_power points (2.34 DPS) | yes | Glowing Magical Bracelets (13106, -1.25 DPS, sim-verified) [world_drop]; Nightsky Wristbands (6407, -1.34 DPS) [world_drop]; Mindthrust Bracers (1974, -1.38 DPS) [dungeon] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Gloves of Meditation (4318, -0.31 DPS) [crafted]; Gloves of Old (9395, -0.38 DPS) [world_drop]; Shilly Mitts (9609, -0.39 DPS, sim-verified) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.7 spell_power points (2.83 DPS) | yes | Belt of Arugal (6392, -0.03 DPS) [dungeon]; Crimson Silk Belt (7055, -0.61 DPS) [crafted]; Invoker's Cord (215366, -0.64 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 15.4 spell_power points (3.43 DPS) | yes | Pearly Leggings (253995, -0.37 DPS) [crafted]; Pristine Leggings (253987, -0.57 DPS) [crafted]; Filigreed Pearly Leggings (253945, -0.95 DPS) [crafted] |
| feet | Golden Slippers (254009) | Tailoring [crafted] | 12.8 spell_power points (2.85 DPS) | yes | Gilded Slippers (254001, -0.08 DPS) [crafted]; Acidic Walkers (9454, -0.40 DPS) [dungeon]; Spidersilk Boots (4320, -0.79 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 12.6 spell_power points (2.80 DPS) | yes | Lorekeeper's Ring (19525, -1.02 DPS) [rep]; Ring of Power Regulation (273030, -1.13 DPS) [dungeon]; Minor Channeling Ring (1449, -1.44 DPS) [quest] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 9.0 spell_power points (2.01 DPS) | yes | Lorekeeper's Ring (19525, -0.22 DPS) [rep]; Ring of Power Regulation (273030, -0.33 DPS) [dungeon]; Minor Channeling Ring (1449, -0.64 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Soulstaff (249393, -0.75 DPS) [crafted]; Rod of the Sleepwalker (1155, -1.20 DPS) [dungeon]; Glimmering Staff (249392, -2.63 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.4 spell_power points (2.31 DPS) | yes | Arcane Infused Rod (279838, -0.93 DPS) [quest]; Satyr's Rod (15962, -1.39 DPS) [world_drop]; Orb of Souls (249395, -1.82 DPS, sim-verified) [crafted] |
| ranged | Starfaller (13063) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderwood (13062, -0.15 DPS) [world_drop]; Greater Mystic Wand (217287, -0.29 DPS) [crafted]; Wand of Eventide (5214, -0.42 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Prelacy Cape; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Golden Slippers; finger1: Snake Hoop; finger2: Runemender's Seal; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Starfaller

No-known-source sample (15 of 326, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 141005213100301531-00000000000000000-000000000000000000)

Set DPS (verified): 167.2. Weights run: 3.1s. Verify run: 4.0s. 456 eligible items had no known source.

Race sweep: gnome 167.2 ± 0.3, human 166.8 ± 0.3, high-order-skyborne 164.8 ± 0.3 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.603 ± 0.022, crit=0.260 ± 0.006 per rating point (14 rating = 1%, 3.640 per %), hit=0.611 ± 0.034 per rating point (10 rating = 1%, 6.115 per %), spell_haste=1.979 ± 0.384, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.628 ± 0.002, spirit=not significant (0.387 ± 0.200), mp5=not significant (0.469 ± 0.230)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Miner's Hat of the Deep (9429) | Uldaman [dungeon] | 22.3 spell_power points (7.11 DPS) | yes | Papal Fez (9431, -0.16 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -0.40 DPS) [crafted]; Corpseshroud (10574, -1.21 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.8 spell_power points (3.78 DPS) | yes | Triune Amulet (7722, -0.36 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.90 DPS) [quest]; Necklace of Calisea (1714, -1.56 DPS) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.8 spell_power points (4.74 DPS) | yes | Green Silken Shoulders (7057, -0.07 DPS) [crafted]; Bloodmage Mantle (7684, -0.13 DPS) [dungeon]; Mantle of Doan (7712, -0.36 DPS) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Long Silken Cloak (4326, -0.10 DPS) [crafted]; Guardian Cloak (5965, -0.10 DPS) [crafted]; Expeditionary's Cape (271922, -0.36 DPS) [rep] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 28.8 spell_power points (9.22 DPS) | yes | Robe of the Magi (1716, -0.41 DPS) [world_drop]; Robe of Power (7054, -1.44 DPS) [crafted]; Elemental Raiment (9434, -2.51 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 10.4 spell_power points (3.33 DPS) | yes | Arcane Runed Bracers (4744, -0.45 DPS) [quest]; Condor Bracers (15864, -1.09 DPS) [quest]; Earthen Silk Cuffs (254019, -1.43 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 23.1 spell_power points (7.39 DPS) | yes | Red Mageweave Gloves (10018, -1.55 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -2.60 DPS) [crafted]; Golden Handwraps (254029, -3.23 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 16.4 spell_power points (5.25 DPS) | yes | Deathmage Sash (10771, -0.12 DPS) [dungeon]; Golden Cord (254045, -0.15 DPS) [crafted]; Gilded Cord (254037, -0.40 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.2 spell_power points (6.79 DPS) | yes | Stoneweaver Leggings (9407, -1.42 DPS, sim-verified) [dungeon]; Crimson Silk Pantaloons (7062, -1.72 DPS) [crafted]; Gaze Dreamer Pants (6903, -1.84 DPS) [dungeon] |
| feet | Golden Slippers (254009) | Tailoring [crafted] | 13.9 spell_power points (4.45 DPS) | yes | Gilded Slippers (254001, -0.37 DPS) [crafted]; Acidic Walkers (9454, -0.82 DPS) [dungeon]; Furen's Boots (13100, -1.12 DPS) [world_drop] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.0 spell_power points (4.16 DPS) | yes | Lorekeeper's Ring (19524, -0.83 DPS) [rep]; Runemender's Seal (273031, -1.18 DPS) [world]; Ring of Forlorn Spirits (2043, -1.60 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 12.9 spell_power points (4.13 DPS) | yes | Lorekeeper's Ring (19524, -0.81 DPS) [rep]; Runemender's Seal (273031, -1.15 DPS) [world]; Ring of Forlorn Spirits (2043, -1.58 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -9.10 DPS, sim-verified) [dungeon]; Celestial Stave (9517, -10.96 DPS) [quest]; Hypnotic Blade (7714, -11.22 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Nether Force Wand (11263) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plaguerot Sprig (10766, -0.38 DPS) [dungeon]; Starfaller (13063, -0.60 DPS) [world_drop]; Twisted Nether Wand (249144, -0.61 DPS) [crafted] |

**New at 40:** head: Miner's Hat of the Deep; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Icy Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; finger1: Philanthropist's Ring; finger2: Snake Hoop; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Staff of Jordan; ranged: Nether Force Wand

No-known-source sample (15 of 456, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (gnome, 141005213100311531-30000000000000000-033000000000000000)

Set DPS (verified): 276.3. Weights run: 3.1s. Verify run: 6.2s. 598 eligible items had no known source.

Race sweep: gnome 276.3 ± 0.4, human 274.6 ± 0.4, high-order-skyborne 271.3 ± 0.5 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.871 ± 0.032, crit=0.406 ± 0.011 per rating point (14 rating = 1%, 5.683 per %), hit=0.830 ± 0.056 per rating point (10 rating = 1%, 8.304 per %), spell_haste=2.825 ± 0.549, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.805 ± 0.002, spirit=not significant (-0.082 ± 0.282), mp5=not significant (0.109 ± 0.321)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 36.4 spell_power points (12.95 DPS) | yes | Knight-Lieutenant's Dreadweave Hat (220889, -2.24 DPS) [vendor]; Dreamweave Circlet (10041, -2.39 DPS) [crafted]; Chief Architect's Monocle (11839, -2.56 DPS) [dungeon] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 21.2 spell_power points (7.55 DPS) | yes | Glowing Eye of Mordresh (10769, -2.33 DPS, sim-verified) [dungeon]; Talvash's Enhancing Necklace (7673, -2.97 DPS) [quest]; Gemshard Heart (17707, -2.98 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 28.7 spell_power points (10.20 DPS) | yes | Kentic Amice (11624, -1.19 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Mantle (220887, -2.55 DPS) [vendor]; Nethershine Shoulders (254057, -2.85 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Big Voodoo Cloak (8216, -2.27 DPS) [crafted]; Expeditionary's Cape (271921, -2.50 DPS) [rep]; Runecloth Cloak (13860, -3.81 DPS, sim-verified) [crafted] |
| chest | Robes of Insight (940) | World drop [world_drop] | 31.9 spell_power points (11.35 DPS) | yes | Robe of the Magi (1716, -1.66 DPS) [world_drop]; Acumen Robes (17775, -1.72 DPS) [quest]; Knight's Dreadweave Vest (220886, -2.31 DPS, sim-verified) [vendor] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths [dungeon] | 17.1 spell_power points (6.07 DPS) | yes | Nethershine Cuffs (254069, -0.18 DPS) [crafted]; Nethergeld Cuffs (254061, -1.41 DPS) [crafted]; Bloodband Bracers (11469, -1.50 DPS) [quest] |
| hands | Golden Gloves (254103) | Tailoring [crafted] | 24.0 spell_power points (8.54 DPS) | yes | Dreamweave Gloves (10019, -0.90 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.13 DPS) [vendor]; Runecloth Gloves (13863, -1.48 DPS) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 29.0 spell_power points (10.32 DPS) | yes | Satyrmane Sash (17755, -2.24 DPS) [dungeon]; Dawnspire Cord (12466, -2.30 DPS) [dungeon]; Golden Waistcord (254089, -2.92 DPS, sim-verified) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | 31.7 spell_power points (11.28 DPS) | yes | Knight's Dreadweave Leggings (220888, -1.27 DPS) [vendor]; Red Mageweave Pants (10009, -2.58 DPS) [crafted]; Crimson Silk Pantaloons (7062, -4.40 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) | PvP rank 5 · Sergeant Major · Alliance [vendor] | 24.1 spell_power points (8.59 DPS) | yes | Golden Sandals (254115, -0.07 DPS) [crafted]; Gilded Sandals (254107, -1.89 DPS) [crafted]; Southsea Mojo Boots (20641, -2.33 DPS) [quest] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 19.7 spell_power points (7.01 DPS) | yes | Woodseed Hoop (17768, -1.78 DPS) [quest]; Mindseye Circle (10634, -1.82 DPS) [dungeon]; Philanthropist's Ring (281635, -1.90 DPS) [quest] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.1 spell_power points (5.37 DPS) | yes | Woodseed Hoop (17768, -0.14 DPS) [quest]; Mindseye Circle (10634, -0.17 DPS) [dungeon]; Philanthropist's Ring (281635, -0.26 DPS) [quest] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Enriched Seal (271903, +0.00 DPS) [rep] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, -0.52 DPS) [world_drop]; Enriched Seal (271903, -2.46 DPS) [rep] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Staff of Jordan (873, -8.33 DPS, sim-verified) [world_drop] |
| off_hand | Truesilver Conduit (249455) | Enchanting [crafted] | 17.0 spell_power points (6.03 DPS) | yes | Celestial Orb (7515, -0.48 DPS) [quest]; Orb of Lorica (11262, -0.80 DPS) [quest]; Orb of the Forgotten Seer (7685, -1.06 DPS) [dungeon] |
| ranged | Nether Force Wand (11263) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lesser Eternal Wand (249232, +0.00 DPS) [crafted]; Charged Lightning Rod (11860, -0.49 DPS) [quest]; Noxious Shooter (17745, -0.52 DPS) [dungeon] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Robes of Insight; wrist: Aristocratic Cuffs; hands: Golden Gloves; waist: Ban'thok Sash; legs: Spellshock Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Thunderbrew's Boot Flask; trinket2: Frozen Heart of the Mountain; main_hand: Blade of Eternal Darkness; off_hand: Truesilver Conduit

No-known-source sample (15 of 598, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (gnome, 141005213100311531-30000000000000000-053500030000000000)

Set DPS (verified): 411.8. Weights run: 3.2s. Verify run: 20.1s. 1456 eligible items had no known source.

Race sweep: gnome 411.8 ± 0.6, human 411.1 ± 0.6, high-order-skyborne 406.4 ± 0.6 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.807 ± 0.036, crit=0.448 ± 0.012 per rating point (14 rating = 1%, 6.266 per %), hit=0.874 ± 0.049 per rating point (10 rating = 1%, 8.737 per %), spell_haste=2.974 ± 0.617, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.728 ± 0.003, spirit=not significant (1.160 ± 0.444), mp5=not significant (1.086 ± 0.419)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | The Postmaster's Band (13390) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crimson Felt Hat (18727, +0.00 DPS) [dungeon]; Spellweaver's Turban (22267, +0.00 DPS) [dungeon]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jeweled Amulet of Cainwyn (1443, +0.00 DPS) [world_drop]; Star of Mystaria (12103, -0.31 DPS) [dungeon]; Nacreous Shell Necklace (22403, -0.56 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Nathrezim (18720, +0.00 DPS) [dungeon]; Lieutenant Commander's Silk Mantle (227102, +0.00 DPS) [pvp]; Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, -0.78 DPS) [quest]; Spritecaster Cape (11623, -1.40 DPS) [dungeon]; Crystalline Threaded Cape (20697, -2.07 DPS) [world] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Vestments (231603, +0.00 DPS) [pvp]; Embrace of the Wind Serpent (12462, -0.54 DPS) [world]; Robes of the Exalted (13346, -1.94 DPS) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 36.6 spell_power points (9.49 DPS) | yes | Marshal's Silk Bracers (16438, -2.50 DPS) [pvp]; Wyrmthalak's Shackles (13958, -3.16 DPS) [quest]; Devout Bracers (16697, -3.76 DPS) [dungeon] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 39.5 spell_power points (10.26 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Marshal's Silk Handwraps (231609, +0.00 DPS) [vendor] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 38.4 spell_power points (9.95 DPS) | yes | Marshal's Silk Sash (16439, -0.62 DPS) [pvp]; Thuzadin Sash (18740, -1.02 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -1.46 DPS) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Magister's Leggings (16687, -0.33 DPS) [dungeon]; Spiritshroud Leggings (12965, -0.55 DPS) [dungeon] |
| feet | The Postmaster's Treads (13391) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wolfrunner Shoes (13101, +0.00 DPS) [world_drop]; Mooncloth Boots (15802, +0.00 DPS) [crafted]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp] |
| finger1 | The Postmaster's Seal (13392) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Maiden's Circle (13001, +0.00 DPS) [world_drop]; Blessed Band of Light (272407, -0.34 DPS) [vendor]; Brainlash (6440, -1.16 DPS) [dungeon] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Maiden's Circle (13001, +0.00 DPS) [world_drop]; Blessed Band of Light (272407, -0.34 DPS) [vendor]; Brainlash (6440, -1.15 DPS) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -1.06 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.66 DPS) [quest]; Thunderbrew's Boot Flask (744, -2.26 DPS) [quest] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Headmaster's Charge (13937, +0.00 DPS) [dungeon]; Ironbark Staff (20069, +0.00 DPS) [rep]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (7.52 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Penelope's Rose (15805, -0.39 DPS) [quest]; Grand Marshal's Tome of Restoration (234590, -0.70 DPS) [pvp] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Oblivion's Touch (18761, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -0.71 DPS) [dungeon] |

**New at 60:** head: The Postmaster's Band; neck: Lady Maye's Pendant; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Magician's Cord; legs: Sentinel's Silk Leggings; feet: The Postmaster's Treads; finger1: The Postmaster's Seal; finger2: Channeler's Ring; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Brilliant Wand

No-known-source sample (15 of 1456, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (gnome, 131205203100311531-52500001200100000-003000000000000000)

Set DPS (verified): 799.8. Weights run: 3.8s. Verify run: 20.0s. 1456 eligible items had no known source.

Race sweep: gnome 799.8 ± 0.8, human 795.0 ± 0.8, high-order-skyborne 787.0 ± 0.8 (dps); the guide race gnome is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.884 ± 0.047, crit=0.885 ± 0.040 per rating point (14 rating = 1%, 12.387 per %), hit=1.852 ± 0.114 per rating point (10 rating = 1%, 18.521 per %), spell_haste=3.621 ± 0.810, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.561 ± 0.002, spirit=not significant (0.392 ± 0.361), mp5=not significant (0.473 ± 0.411)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 56.6 spell_power points (29.93 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -2.62 DPS) [pvp]; Crimson Felt Hat (18727, -8.31 DPS, sim-verified) [dungeon] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 34.6 spell_power points (18.29 DPS) | yes | Star of Mystaria (12103, -0.84 DPS) [dungeon]; Lady Maye's Pendant (14558, -3.59 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -4.63 DPS) [world_drop] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 40.9 spell_power points (21.61 DPS) | yes | Field Marshal's Silk Spaulders (231602, -0.35 DPS) [pvp]; Lieutenant Commander's Silk Mantle (227102, -1.16 DPS) [pvp]; Soulstealer Mantle (13374, -1.54 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of Arcane Mastery (22330, -7.58 DPS, sim-verified) [dungeon]; Shroud of the Exile (15421, -9.52 DPS) [quest]; Crystalline Threaded Cape (20697, -9.55 DPS) [world] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Vestments (231603, -0.12 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -6.67 DPS) [pvp]; Alanna's Embrace (13314, -21.03 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 31.8 spell_power points (16.82 DPS) | yes | Marshal's Silk Bracers (16438, -5.86 DPS) [pvp]; Manacle Cuffs (11962, -6.11 DPS) [quest]; Runecloth Cuffs (254123, -6.33 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 42.6 spell_power points (22.53 DPS) | yes | Marshal's Silk Gloves (16440, -1.61 DPS) [vendor]; Marshal's Silk Gauntlets (231608, -1.61 DPS) [vendor]; Hands of Power (13253, -4.42 DPS, sim-verified) [dungeon] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 40.2 spell_power points (21.26 DPS) | yes | Clutch of Andros (13956, -2.59 DPS) [dungeon]; Highlander's Cloth Girdle (20047, -4.51 DPS) [rep]; Ban'thok Sash (11662, -7.22 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 65.6 spell_power points (34.69 DPS) | yes | Marshal's Silk Leggings (231605, -0.86 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -8.05 DPS) [pvp]; Magister's Leggings (16687, -10.81 DPS) [dungeon] |
| feet | Kayser's Boots of Precision (22231) | Blackrock Spire: War Master Voone [dungeon] | 38.4 spell_power points (20.32 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -3.32 DPS) [dungeon]; Sergeant Major's Dreadweave Boots (220891, -4.94 DPS, sim-verified) [vendor] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Julio's Band (19325, -1.65 DPS) [rep]; Maiden's Circle (13001, -3.96 DPS) [world_drop]; Naglering (11669, -19.61 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Maiden's Circle (13001, -2.70 DPS) [world_drop]; Songstone of Ironforge (12543, -3.95 DPS) [quest]; Don Julio's Band (19325, -4.69 DPS, sim-verified) [rep] |
| trinket1 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Thunderbrew's Boot Flask (744, -8.40 DPS, sim-verified) [quest] |
| trinket2 | Eye of the Beast (13968) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, -4.53 DPS, sim-verified) [vendor] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Staff of Jordan (873, -13.17 DPS, sim-verified) [world_drop]; Headmaster's Charge (13937, -14.91 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Oblivion's Touch (18761, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -2.17 DPS) [dungeon] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Magician's Cord; legs: Sentinel's Silk Leggings; feet: Kayser's Boots of Precision; finger1: Channeler's Ring; finger2: Blessed Band of Light; trinket1: Second Wind; trinket2: Eye of the Beast; main_hand: Ironbark Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1456, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (undead, 140005000100000000-00000000000000000-000000000000000000)

Set DPS (verified): 40.9. Weights run: 2.6s. Verify run: 2.6s. 192 eligible items had no known source.

Race sweep: undead 40.9 ± 0.1, orc 39.3 ± 0.0, troll 38.8 ± 0.0 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.992 ± 0.010, crit=0.085 ± 0.002 per rating point (14 rating = 1%, 1.185 per %), hit=0.316 ± 0.012 per rating point (10 rating = 1%, 3.164 per %), spell_haste=not significant (0.185 ± 0.083), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.888 ± 0.003, spirit=not significant (0.165 ± 0.064), mp5=1.228 ± 0.124

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pearly Circlet (253957) | Tailoring [crafted] | 9.8 spell_power points (1.61 DPS) | yes | Pristine Circlet (253949, -0.61 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.63 DPS) [crafted]; Flying Tiger Goggles (4368, -1.50 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.5 spell_power points (0.08 DPS) | yes | Scholarly Pendant (277203, -0.03 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.9 spell_power points (2.29 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.86 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.63 DPS) [crafted]; Slime-encrusted Pads (6461, -1.69 DPS) [dungeon] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.05 DPS) [crafted]; Sanguine Cape (14376, -0.17 DPS) [world_drop]; Battle Healer's Cloak (20427, -0.27 DPS) [rep] |
| chest | Filigreed Pearly Gown (253909) | Tailoring [crafted] | 13.4 spell_power points (2.21 DPS) | yes | Filigreed Pristine Gown (253901, -0.49 DPS, sim-verified) [crafted]; Gray Woolen Robe (2585, -0.74 DPS) [crafted]; Mystic's Wrap (14369, -1.07 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 8.6 spell_power points (1.42 DPS) | yes | Bright Bracers (3647, -0.61 DPS, sim-verified) [world_drop]; Tabitha's Cuffs (251486, -0.93 DPS) [quest]; Owlbeard Bracers (16981, -0.93 DPS) [quest] |
| hands | Pearly Gloves (253921) | Tailoring [crafted] | 9.2 spell_power points (1.51 DPS) | yes | Pristine Gloves (253913, -0.36 DPS) [crafted]; Gnoll Casting Gloves (892, -0.53 DPS) [world]; Serpent Gloves (5970, -0.55 DPS, sim-verified) [dungeon] |
| waist | Pearly Sash (253933) | Tailoring [crafted] | 10.2 spell_power points (1.68 DPS) | yes | Pristine Sash (253925, -0.36 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.37 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.53 DPS) [crafted] |
| legs | Filigreed Pearly Leggings (253945) | Tailoring [crafted] | 16.4 spell_power points (2.70 DPS) | yes | Filigreed Pristine Leggings (253937, -0.61 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -1.39 DPS) [world_drop]; Silk-threaded Trousers (1929, -1.54 DPS) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 11.0 spell_power points (1.81 DPS) | yes | Pearly Boots (253897, -0.43 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.82 DPS) [crafted]; Red Woolen Boots (4313, -1.01 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 7.5 spell_power points (1.23 DPS) | yes | Loop of Sacrifice (281673, -0.41 DPS) [quest]; Volcanic Rock Ring (12053, -0.74 DPS) [world_drop]; Furlfeather Ring (286745, -1.06 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 6.0 spell_power points (0.98 DPS) | yes | Loop of Sacrifice (281673, -0.16 DPS) [quest]; Volcanic Rock Ring (12053, -0.49 DPS) [world_drop]; Furlfeather Ring (286745, -0.82 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Crescent Staff (6505) | Leaders of the Fang [quest] | 32.9 spell_power points (5.42 DPS) | yes | Evocator's Blade (2567, -0.25 DPS, sim-verified) [dungeon]; Staff of the Blessed Seer (2271, -0.54 DPS) [dungeon]; Gnarled Necromancer's Staff (251534, -0.87 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Skycaller (12984) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cookie's Stirring Rod (5198, -0.80 DPS, sim-verified) [dungeon]; Sizzle Stick (8071, -1.46 DPS) [quest]; Greater Magic Wand (11288, -1.49 DPS) [crafted] |

**New at 20:** head: Pearly Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Pearly Gown; wrist: Mindthrust Bracers; hands: Pearly Gloves; waist: Pearly Sash; legs: Filigreed Pearly Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Crescent Staff; ranged: Skycaller

No-known-source sample (15 of 192, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 8350 The 1 Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 23173 Abomination Skin Leggings

### Band 30 (undead, 141005203100301000-00000000000000000-000000000000000000)

Set DPS (verified): 90.2. Weights run: 3.1s. Verify run: 4.6s. 318 eligible items had no known source.

Race sweep: undead 90.2 ± 0.1, orc 88.8 ± 0.1, troll 87.6 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.705 ± 0.016, crit=0.107 ± 0.003 per rating point (14 rating = 1%, 1.503 per %), hit=0.406 ± 0.024 per rating point (10 rating = 1%, 4.056 per %), spell_haste=1.400 ± 0.250, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.629 ± 0.002, spirit=not significant (0.185 ± 0.115), mp5=not significant (0.662 ± 0.180)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Enchanter's Cowl (4322, -0.27 DPS, sim-verified) [crafted]; Nightsky Cowl (4039, -0.75 DPS) [world_drop]; Silk Headband (7050, -0.86 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 4.2 spell_power points (1.17 DPS) | yes | Kaleidoscope Chain (13084, -0.19 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.19 DPS) [vendor]; Crystal Starfire Medallion (5003, -0.24 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.3 spell_power points (4.23 DPS) | yes | Death Speaker Mantle (6685, -0.74 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -1.10 DPS) [world_drop]; Invoker's Mantle (215365, -1.33 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windsong Drape (15468, -0.13 DPS) [quest]; Hillman's Cloak (3719, -0.33 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.40 DPS) [crafted] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 18.9 spell_power points (5.23 DPS) | yes | Green Silk Armor (7065, -0.21 DPS) [crafted]; Mechbuilder's Overalls (9508, -0.68 DPS) [dungeon]; Death Speaker Robes (6682, -1.15 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 11.0 spell_power points (3.03 DPS) | yes | Mindthrust Bracers (1974, -1.51 DPS) [dungeon]; Nightsky Wristbands (6407, -1.71 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.86 DPS, sim-verified) [world_drop] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 10.6 spell_power points (2.92 DPS) | yes | Jutebraid Gloves (10654, -0.70 DPS, sim-verified) [quest]; Truefaith Gloves (7049, -0.96 DPS) [crafted]; Serpent Gloves (5970, -0.99 DPS) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 13.1 spell_power points (3.62 DPS) | yes | Belt of Arugal (6392, -0.30 DPS) [dungeon]; Warsong Sash (16975, -0.58 DPS) [quest]; Crimson Silk Belt (7055, -0.60 DPS) [crafted] |
| legs | Pearly Leggings (253995) | Tailoring [crafted] | 14.7 spell_power points (4.05 DPS) | yes | Gaze Dreamer Pants (6903, -0.28 DPS) [dungeon]; Pristine Leggings (253987, -0.50 DPS) [crafted]; Filigreed Pearly Leggings (253945, -0.77 DPS) [crafted] |
| feet | Golden Slippers (254009) | Tailoring [crafted] | 13.9 spell_power points (3.82 DPS) | yes | Gilded Slippers (254001, -0.32 DPS) [crafted]; Acidic Walkers (9454, -0.68 DPS) [dungeon]; Spidersilk Boots (4320, -1.11 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 12.2 spell_power points (3.37 DPS) | yes | Advisor's Ring (19521, -1.08 DPS) [rep]; Ring of Power Regulation (273030, -1.17 DPS) [dungeon] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 8.5 spell_power points (2.35 DPS) | yes | Ring of Power Regulation (273030, -0.15 DPS) [dungeon]; Advisor's Ring (20426, -0.61 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Soulstaff (249393, -1.30 DPS) [crafted]; Tendonscraper Dagger (273053, -1.31 DPS) [world]; Glimmering Staff (249392, -3.15 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 11.2 spell_power points (3.10 DPS) | yes | Orb of Souls (249395, -1.69 DPS) [crafted]; Witch's Finger (16887, -1.75 DPS, sim-verified) [quest]; Satyr's Rod (15962, -1.83 DPS) [world_drop] |
| ranged | Starfaller (13063) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderwood (13062, -0.39 DPS) [world_drop]; Greater Mystic Wand (217287, -0.59 DPS) [crafted]; Wand of Eventide (5214, -0.72 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Darkspear Raider's Cloak; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Defiler's Cloth Girdle; legs: Pearly Leggings; feet: Golden Slippers; finger1: Snake Hoop; finger2: Runemender's Seal; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Starfaller

No-known-source sample (15 of 318, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 8350 The 1 Ring; 9362 Brilliant Gold Ring

### Band 40 (undead, 141005213100301531-00000000000000000-000000000000000000)

Set DPS (verified): 166.3. Weights run: 3.1s. Verify run: 3.0s. 448 eligible items had no known source.

Race sweep: undead 166.3 ± 0.3, orc 164.8 ± 0.3, troll 163.4 ± 0.3 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.603 ± 0.022, crit=0.260 ± 0.006 per rating point (14 rating = 1%, 3.640 per %), hit=0.611 ± 0.034 per rating point (10 rating = 1%, 6.115 per %), spell_haste=1.979 ± 0.384, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.628 ± 0.002, spirit=not significant (0.387 ± 0.200), mp5=not significant (0.469 ± 0.230)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Miner's Hat of the Deep (9429) | Uldaman [dungeon] | 22.3 spell_power points (7.11 DPS) | yes | Papal Fez (9431, -0.16 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -0.40 DPS) [crafted]; Corpseshroud (10574, -1.21 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.8 spell_power points (3.78 DPS) | yes | Triune Amulet (7722, -0.36 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.90 DPS) [quest]; Necklace of Calisea (1714, -1.56 DPS) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.8 spell_power points (4.74 DPS) | yes | Green Silken Shoulders (7057, -0.07 DPS) [crafted]; Bloodmage Mantle (7684, -0.13 DPS) [dungeon]; Mantle of Doan (7712, -0.36 DPS) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Long Silken Cloak (4326, -0.10 DPS) [crafted]; Guardian Cloak (5965, -0.10 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.36 DPS) [vendor] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 28.8 spell_power points (9.22 DPS) | yes | Robe of Power (7054, -1.44 DPS) [crafted]; Robe of the Magi (1716, -1.70 DPS, sim-verified) [world_drop]; Elemental Raiment (9434, -2.51 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 10.4 spell_power points (3.33 DPS) | yes | Radiant Silver Bracers (4545, -0.51 DPS) [quest]; Condor Bracers (15864, -1.09 DPS) [quest]; Earthen Silk Cuffs (254019, -1.43 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 23.1 spell_power points (7.39 DPS) | yes | Red Mageweave Gloves (10018, -2.05 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -2.60 DPS) [crafted]; Golden Handwraps (254029, -3.23 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 16.4 spell_power points (5.25 DPS) | yes | Deathmage Sash (10771, -0.12 DPS) [dungeon]; Golden Cord (254045, -0.15 DPS) [crafted]; Gilded Cord (254037, -0.40 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.2 spell_power points (6.79 DPS) | yes | Crimson Silk Pantaloons (7062, -1.72 DPS) [crafted]; Gaze Dreamer Pants (6903, -1.84 DPS) [dungeon]; Stoneweaver Leggings (9407, -2.08 DPS, sim-verified) [dungeon] |
| feet | Golden Slippers (254009) | Tailoring [crafted] | 13.9 spell_power points (4.45 DPS) | yes | Gilded Slippers (254001, -0.37 DPS) [crafted]; Acidic Walkers (9454, -0.82 DPS) [dungeon]; Furen's Boots (13100, -1.12 DPS) [world_drop] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.0 spell_power points (4.16 DPS) | yes | Advisor's Ring (19520, -0.83 DPS) [rep]; Runemender's Seal (273031, -1.18 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 12.9 spell_power points (4.13 DPS) | yes | Runemender's Seal (273031, -1.15 DPS) [world]; Advisor's Ring (19521, -1.60 DPS) [rep] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Celestial Stave (9517, -10.96 DPS) [quest]; Illusionary Rod (7713, -11.15 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -11.22 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Nether Force Wand (11263) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plaguerot Sprig (10766, -0.38 DPS) [dungeon]; Starfaller (13063, -0.61 DPS) [world_drop]; Dancing Flame (6806, -0.61 DPS) [quest] |

**New at 40:** head: Miner's Hat of the Deep; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Icy Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; finger1: Philanthropist's Ring; finger2: Snake Hoop; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Staff of Jordan; ranged: Nether Force Wand

No-known-source sample (15 of 448, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (undead, 141005213100311531-30000000000000000-033000000000000000)

Set DPS (verified): 273.3. Weights run: 3.1s. Verify run: 7.6s. 590 eligible items had no known source.

Race sweep: undead 273.3 ± 0.4, orc 270.1 ± 0.4, troll 267.5 ± 0.4 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.871 ± 0.032, crit=0.406 ± 0.011 per rating point (14 rating = 1%, 5.683 per %), hit=0.830 ± 0.056 per rating point (10 rating = 1%, 8.304 per %), spell_haste=2.825 ± 0.549, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.805 ± 0.002, spirit=not significant (-0.082 ± 0.282), mp5=not significant (0.109 ± 0.321)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 36.4 spell_power points (12.95 DPS) | yes | Blood Guard's Dreadweave Hat (220907, -2.24 DPS) [vendor]; Dreamweave Circlet (10041, -2.39 DPS) [crafted]; Chief Architect's Monocle (11839, -2.56 DPS) [dungeon] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 21.2 spell_power points (7.55 DPS) | yes | Glowing Eye of Mordresh (10769, -2.57 DPS, sim-verified) [dungeon]; Jarkal's Enhancing Necklace (7888, -2.97 DPS) [quest]; Gemshard Heart (17707, -2.98 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 28.7 spell_power points (10.20 DPS) | yes | Kentic Amice (11624, -1.19 DPS) [dungeon]; Blood Guard's Dreadweave Mantle (220905, -2.55 DPS) [vendor]; Nethershine Shoulders (254057, -2.85 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Runecloth Cloak (13860, -1.16 DPS) [crafted]; Deep Woodlands Cloak (19121, -1.19 DPS, sim-verified) [quest]; Big Voodoo Cloak (8216, -2.27 DPS) [crafted] |
| chest | Robes of Insight (940) | World drop [world_drop] | 31.9 spell_power points (11.35 DPS) | yes | Robe of the Magi (1716, -1.66 DPS) [world_drop]; Acumen Robes (17775, -1.72 DPS) [quest]; Stone Guard's Dreadweave Vest (220904, -3.64 DPS, sim-verified) [vendor] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths [dungeon] | 17.1 spell_power points (6.07 DPS) | yes | Nethershine Cuffs (254069, -0.18 DPS) [crafted]; Nethergeld Cuffs (254061, -1.41 DPS) [crafted]; Bloodband Bracers (11469, -1.50 DPS) [quest] |
| hands | Golden Gloves (254103) | Tailoring [crafted] | 24.0 spell_power points (8.54 DPS) | yes | First Sergeant's Dreadweave Gloves (220908, -1.13 DPS) [vendor]; Runecloth Gloves (13863, -1.48 DPS) [crafted]; Dreamweave Gloves (10019, -3.62 DPS, sim-verified) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 29.0 spell_power points (10.32 DPS) | yes | Satyrmane Sash (17755, -2.24 DPS) [dungeon]; Dawnspire Cord (12466, -2.30 DPS) [dungeon]; Golden Waistcord (254089, -3.42 DPS, sim-verified) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | 31.7 spell_power points (11.28 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -1.27 DPS) [vendor]; Red Mageweave Pants (10009, -2.58 DPS) [crafted]; Crimson Silk Pantaloons (7062, -4.40 DPS) [crafted] |
| feet | First Sergeant's Dreadweave Boots (220909) | PvP rank 5 · First Sergeant · Horde [vendor] | 24.1 spell_power points (8.59 DPS) | yes | Golden Sandals (254115, -0.07 DPS) [crafted]; Gilded Sandals (254107, -1.89 DPS) [crafted]; Southsea Mojo Boots (20641, -2.33 DPS) [quest] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 19.7 spell_power points (7.01 DPS) | yes | Woodseed Hoop (17768, -1.78 DPS) [quest]; Mindseye Circle (10634, -1.82 DPS) [dungeon]; Philanthropist's Ring (281635, -1.90 DPS) [quest] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.1 spell_power points (5.37 DPS) | yes | Woodseed Hoop (17768, -0.14 DPS) [quest]; Mindseye Circle (10634, -0.17 DPS) [dungeon]; Philanthropist's Ring (281635, -0.26 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+5.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.52 DPS) [world_drop]; Rune of the Guard Captain (19120, -0.59 DPS) [quest]; Darkspear Voodoo Seal (272059, -3.31 DPS, sim-verified) [vendor] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Staff of Jordan (873, -7.70 DPS, sim-verified) [world_drop] |
| off_hand | Truesilver Conduit (249455) | Enchanting [crafted] | 17.0 spell_power points (6.03 DPS) | yes | Flask of Forest Mojo (19115, -0.36 DPS) [quest]; Celestial Orb (7515, -0.48 DPS) [quest]; Orb of the Forgotten Seer (7685, -1.06 DPS) [dungeon] |
| ranged | Nether Force Wand (11263) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Noxious Shooter (17745, -0.32 DPS) [dungeon]; Charged Lightning Rod (11860, -0.49 DPS) [quest]; Lesser Eternal Wand (249232, -2.35 DPS, sim-verified) [crafted] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Robes of Insight; wrist: Aristocratic Cuffs; hands: Golden Gloves; waist: Ban'thok Sash; legs: Spellshock Leggings; feet: First Sergeant's Dreadweave Boots; finger1: Brainlash; finger2: Cyclopean Band; trinket2: Frozen Heart of the Mountain; main_hand: Blade of Eternal Darkness; off_hand: Truesilver Conduit

No-known-source sample (15 of 590, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (undead, 141005213100311531-30000000000000000-053500030000000000)

Set DPS (verified): 405.4. Weights run: 3.2s. Verify run: 19.7s. 1450 eligible items had no known source.

Race sweep: undead 405.4 ± 0.5, orc 401.7 ± 0.5, troll 399.1 ± 0.5 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.807 ± 0.036, crit=0.448 ± 0.012 per rating point (14 rating = 1%, 6.266 per %), hit=0.874 ± 0.049 per rating point (10 rating = 1%, 8.737 per %), spell_haste=2.974 ± 0.617, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.728 ± 0.003, spirit=not significant (1.160 ± 0.444), mp5=not significant (1.086 ± 0.419)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | The Postmaster's Band (13390) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crimson Felt Hat (18727, +0.00 DPS) [dungeon]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Spellweaver's Turban (22267, -8.42 DPS, sim-verified) [dungeon] |
| neck | Jeweled Amulet of Cainwyn (1443) | World drop [world_drop] | 32.1 spell_power points (8.33 DPS) | yes | Lady Maye's Pendant (14558, -1.50 DPS) [world_drop]; Star of Mystaria (12103, -1.82 DPS) [dungeon]; Nacreous Shell Necklace (22403, -2.06 DPS) [dungeon] |
| shoulder | Darkspear Shoulderpads (272103) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Silk Mantle (227104, +0.00 DPS) [pvp]; Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Shroud of the Nathrezim (18720, -2.48 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -1.40 DPS) [dungeon]; Crystalline Threaded Cape (20697, -2.07 DPS) [world]; Shroud of the Exile (15421, -3.92 DPS, sim-verified) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Raiment (231596, +0.00 DPS) [pvp]; Robes of the Exalted (13346, -1.94 DPS) [dungeon]; Embrace of the Wind Serpent (12462, -14.50 DPS, sim-verified) [world] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 36.6 spell_power points (9.49 DPS) | yes | General's Silk Cuffs (16538, -2.50 DPS) [pvp]; Wyrmthalak's Shackles (13958, -3.16 DPS) [quest]; Devout Bracers (16697, -3.76 DPS) [dungeon] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 39.5 spell_power points (10.26 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Magister's Gloves (16684, -2.74 DPS, sim-verified) [dungeon] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 38.4 spell_power points (9.95 DPS) | yes | General's Silk Sash (16537, -0.62 DPS) [pvp]; Wisdom of the Timbermaw (19047, -1.46 DPS) [crafted]; Thuzadin Sash (18740, -7.00 DPS, sim-verified) [dungeon] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 54.9 spell_power points (14.25 DPS) | yes | General's Silk Trousers (231595, +0.00 DPS) [pvp]; Spiritshroud Leggings (12965, -1.16 DPS) [dungeon]; Magister's Leggings (16687, -4.61 DPS, sim-verified) [dungeon] |
| feet | The Postmaster's Treads (13391) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Wolfrunner Shoes (13101, -7.58 DPS, sim-verified) [world_drop] |
| finger1 | The Postmaster's Seal (13392) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -0.34 DPS) [vendor]; Brainlash (6440, -1.16 DPS) [dungeon]; Maiden's Circle (13001, -7.23 DPS, sim-verified) [world_drop] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-verified (+13.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Maiden's Circle (13001, +0.00 DPS) [world_drop]; Blessed Band of Light (272407, -0.34 DPS) [vendor]; Brainlash (6440, -1.15 DPS) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Weakness Analyzer (272438, -3.61 DPS, sim-verified) [vendor] |
| trinket2 | Eye of the Beast (13968) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Weakness Analyzer (272438, -2.50 DPS, sim-verified) [vendor] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Headmaster's Charge (13937, +0.00 DPS) [dungeon]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Ironbark Staff (20220, -6.07 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (7.52 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Tome of Mending (234564, -0.70 DPS) [pvp]; Penelope's Rose (15805, -4.18 DPS, sim-verified) [quest] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Oblivion's Touch (18761, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -0.70 DPS) [dungeon] |

**New at 60:** head: The Postmaster's Band; neck: Jeweled Amulet of Cainwyn; shoulder: Darkspear Shoulderpads; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Magician's Cord; legs: Outrider's Silk Leggings; feet: The Postmaster's Treads; finger1: The Postmaster's Seal; finger2: Channeler's Ring; trinket1: Burst of Knowledge; trinket2: Eye of the Beast; off_hand: Therazane's Touch; ranged: Brilliant Wand

No-known-source sample (15 of 1450, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60, raid preset (undead, 131205203100311531-52500001200100000-003000000000000000)

Set DPS (verified): 788.6. Weights run: 3.8s. Verify run: 20.9s. 1450 eligible items had no known source.

Race sweep: undead 788.6 ± 0.8, orc 786.8 ± 0.8, troll 779.3 ± 0.8 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.884 ± 0.047, crit=0.885 ± 0.040 per rating point (14 rating = 1%, 12.387 per %), hit=1.852 ± 0.114 per rating point (10 rating = 1%, 18.521 per %), spell_haste=3.621 ± 0.810, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.561 ± 0.002, spirit=not significant (0.392 ± 0.361), mp5=not significant (0.473 ± 0.411)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 56.6 spell_power points (29.93 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -2.62 DPS) [pvp]; Crimson Felt Hat (18727, -10.31 DPS, sim-verified) [dungeon] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 34.6 spell_power points (18.29 DPS) | yes | Star of Mystaria (12103, -1.85 DPS, sim-verified) [dungeon]; Lady Maye's Pendant (14558, -3.59 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -4.63 DPS) [world_drop] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 40.9 spell_power points (21.61 DPS) | yes | Warlord's Silk Amice (231594, -0.35 DPS) [pvp]; Champion's Silk Mantle (227104, -1.16 DPS) [pvp]; Soulstealer Mantle (13374, -6.47 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of Arcane Mastery (22330, -7.12 DPS, sim-verified) [dungeon]; Shroud of the Exile (15421, -9.52 DPS) [quest]; Crystalline Threaded Cape (20697, -9.55 DPS) [world] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Raiment (231596, -0.12 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -6.67 DPS) [pvp]; Alanna's Embrace (13314, -16.54 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 31.8 spell_power points (16.82 DPS) | yes | General's Silk Cuffs (16538, -5.86 DPS) [pvp]; Manacle Cuffs (11962, -6.11 DPS) [quest]; Runecloth Cuffs (254123, -6.33 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 42.6 spell_power points (22.53 DPS) | yes | General's Silk Handguards (16540, -1.61 DPS) [vendor]; General's Silk Gauntlets (231599, -1.61 DPS) [vendor]; Hands of Power (13253, -4.19 DPS, sim-verified) [dungeon] |
| waist | Magician's Cord (272393) | Pix Xizzix [vendor] | 40.2 spell_power points (21.26 DPS) | yes | Ban'thok Sash (11662, -0.46 DPS) [dungeon]; Clutch of Andros (13956, -2.59 DPS) [dungeon]; Defiler's Cloth Girdle (20163, -4.51 DPS) [rep] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 48.7 spell_power points (25.76 DPS) | yes | Legionnaire's Silk Legguards (227107, +0.00 DPS) [pvp]; General's Silk Trousers (231595, +0.00 DPS) [pvp]; Magister's Leggings (16687, -4.76 DPS, sim-verified) [dungeon] |
| feet | Kayser's Boots of Precision (22231) | Blackrock Spire: War Master Voone [dungeon] | 38.4 spell_power points (20.32 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; First Sergeant's Dreadweave Boots (220909, -2.09 DPS) [vendor]; Dragonrider Boots (18102, -3.32 DPS) [dungeon] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Julio's Band (19325, -1.65 DPS) [rep]; Maiden's Circle (13001, -3.96 DPS) [world_drop]; Naglering (11669, -18.43 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Julio's Band (19325, -0.40 DPS) [rep]; Maiden's Circle (13001, -2.70 DPS) [world_drop]; Eye of Orgrimmar (12545, -3.95 DPS) [quest] |
| trinket1 | Eye of the Beast (13968) | For The Horde! [quest] | sim-verified (+5.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+3.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -0.70 DPS) [crafted]; Second Wind (11819, -3.59 DPS, sim-verified) [dungeon] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Staff of Jordan (873, -14.11 DPS, sim-verified) [world_drop]; Headmaster's Charge (13937, -14.91 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Oblivion's Touch (18761, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -2.17 DPS) [dungeon] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Magician's Cord; legs: Outrider's Silk Leggings; feet: Kayser's Boots of Precision; finger1: Channeler's Ring; finger2: Blessed Band of Light; trinket1: Eye of the Beast; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1450, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

